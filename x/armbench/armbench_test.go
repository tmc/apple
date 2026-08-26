package armbench

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/tmc/apple/x/experiment"
)

func spin(d time.Duration) func(context.Context) (Result, error) {
	return func(context.Context) (Result, error) {
		time.Sleep(d)
		return Result{Elapsed: d}, nil
	}
}

func failingArm(context.Context) (Result, error) {
	return Result{}, errors.New("boom")
}

func experiment3(base, cand time.Duration, ctrl time.Duration) Experiment {
	return Experiment{
		Rounds:   5,
		Workload: experiment.Workload{ID: "test/spin", WorkUnits: 100},
		Arms: []Arm{
			{Name: "baseline", Kind: experiment.ArmBaseline, WorkUnits: 100, Fn: spin(base)},
			{Name: "candidate", Kind: experiment.ArmCandidate, WorkUnits: 100, Fn: spin(cand)},
			{Name: "dup", Kind: experiment.ArmDuplicate, WorkUnits: 100, Fn: spin(ctrl)},
		},
	}
}

func TestIdenticalArmsProduceNoWinner(t *testing.T) {
	const d = time.Millisecond
	rcpt, err := Run(context.Background(), experiment3(d, d, d))
	if err != nil {
		t.Fatal(err)
	}
	if rcpt.Refused() {
		t.Fatalf("identical arms refused: %+v", rcpt.Refusals)
	}
	if rcpt.Verdict != experiment.VerdictNoWinner {
		t.Fatalf("identical arms verdict %q, want no-winner", rcpt.Verdict)
	}
}

func TestPlantedSlowdownDetected(t *testing.T) {
	rcpt, err := Run(context.Background(), experiment3(time.Millisecond, 20*time.Millisecond, time.Millisecond))
	if err != nil {
		t.Fatal(err)
	}
	if rcpt.Verdict != experiment.VerdictSlower || rcpt.Winner != "baseline" {
		t.Fatalf("verdict %q winner %q; want slower/baseline", rcpt.Verdict, rcpt.Winner)
	}
}

func TestRealImprovementDeclared(t *testing.T) {
	// Candidate is faster by far more than the control gap.
	rcpt, err := Run(context.Background(), experiment3(20*time.Millisecond, time.Millisecond, 21*time.Millisecond))
	if err != nil {
		t.Fatal(err)
	}
	if rcpt.Verdict != experiment.VerdictFaster || rcpt.Winner != "candidate" {
		t.Fatalf("verdict %q winner %q; want faster/candidate", rcpt.Verdict, rcpt.Winner)
	}
}

// Improvement smaller than the duplicate-control gap must not win even when
// the candidate is nominally faster every round.
func TestImprovementInsideControlGapRefused(t *testing.T) {
	rcpt, err := Run(context.Background(),
		experiment3(10*time.Millisecond, 9*time.Millisecond+800*time.Microsecond, 11*time.Millisecond))
	if err != nil {
		t.Fatal(err)
	}
	if rcpt.Verdict == experiment.VerdictFaster {
		t.Fatalf("claimed a win inside the control gap (noise=%s): %+v", Medians(&rcpt)["baseline"], Medians(&rcpt))
	}
}

func TestMissingDuplicateControlRefused(t *testing.T) {
	x := experiment3(time.Millisecond, time.Millisecond, time.Millisecond)
	x.Arms = x.Arms[:2]
	rcpt, err := Run(context.Background(), x)
	if err != nil {
		t.Fatal(err)
	}
	if !rcpt.HasRefusal(experiment.RefusalMissingControl) {
		t.Fatalf("missing control not refused: %+v", rcpt.Refusals)
	}
	if rcpt.Verdict != experiment.VerdictRefused || len(rcpt.Samples) != 0 {
		t.Fatal("experiment with missing control measured anyway")
	}
}

func TestUnequalWorkRefused(t *testing.T) {
	x := experiment3(time.Millisecond, time.Millisecond, time.Millisecond)
	x.Arms[1].WorkUnits = 50
	rcpt, err := Run(context.Background(), x)
	if err != nil {
		t.Fatal(err)
	}
	if !rcpt.HasRefusal(experiment.RefusalUnequalWork) {
		t.Fatalf("unequal work not refused: %+v", rcpt.Refusals)
	}
}

func TestCorrectnessFailureBlocksConclusion(t *testing.T) {
	x := experiment3(time.Millisecond, time.Millisecond, time.Millisecond)
	x.Arms[1].Fn = failingArm
	rcpt, err := Run(context.Background(), x)
	if err != nil {
		t.Fatal(err)
	}
	if !rcpt.HasRefusal(experiment.RefusalCorrectnessFailed) {
		t.Fatalf("correctness failure not refused: %+v", rcpt.Refusals)
	}
	for _, s := range rcpt.Samples {
		if s.Arm == "candidate" && s.Verdict != experiment.VerdictFail {
			t.Fatalf("failed arm recorded as %q", s.Verdict)
		}
	}
	if rcpt.Verdict != experiment.VerdictRefused {
		t.Fatalf("verdict %q with failed correctness", rcpt.Verdict)
	}
}

func TestImpossibleByteRateRefused(t *testing.T) {
	x := experiment3(time.Millisecond, time.Millisecond, time.Millisecond)
	x.MaxByteRate = 1000 // one kilobyte per second is impossible on real hardware
	for i := range x.Arms {
		arm := x.Arms[i]
		arm.Fn = func(context.Context) (Result, error) {
			return Result{
				Elapsed:         time.Millisecond,
				BytesByLocation: map[string]int64{"device": 1 << 30},
			}, nil
		}
		x.Arms[i] = arm
	}
	rcpt, err := Run(context.Background(), x)
	if err != nil {
		t.Fatal(err)
	}
	if !rcpt.HasRefusal(experiment.RefusalImpossibleByteRate) {
		t.Fatalf("impossible byte rate not refused: %+v", rcpt.Refusals)
	}
}

func TestHostQualificationGatesRun(t *testing.T) {
	x := experiment3(time.Millisecond, time.Millisecond, time.Millisecond)
	x.Qualify = func(context.Context) error { return errors.New("thermal") }
	rcpt, err := Run(context.Background(), x)
	if err != nil {
		t.Fatal(err)
	}
	if !rcpt.HasRefusal(experiment.RefusalContaminatedHost) {
		t.Fatalf("contaminated host not refused: %+v", rcpt.Refusals)
	}
	if len(rcpt.Samples) != 0 {
		t.Fatal("measured on an unqualified host")
	}
}

func TestRotationInterleavesOrder(t *testing.T) {
	x := experiment3(time.Microsecond, time.Microsecond, time.Microsecond)
	x.Rotate = true
	rcpt, err := Run(context.Background(), x)
	if err != nil {
		t.Fatal(err)
	}
	var firstRound []string
	for _, s := range rcpt.Samples {
		if s.Round == 0 {
			firstRound = append(firstRound, s.Arm)
		}
	}
	want := []string{"baseline", "candidate", "dup"}
	for i, name := range want {
		if firstRound[i] != name {
			t.Fatalf("round 0 order %v, want %v", firstRound, want)
		}
	}
	// Round 1 starts from the next arm.
	var secondRound []string
	for _, s := range rcpt.Samples {
		if s.Round == 1 {
			secondRound = append(secondRound, s.Arm)
		}
	}
	want2 := []string{"candidate", "dup", "baseline"}
	for i, name := range want2 {
		if secondRound[i] != name {
			t.Fatalf("round 1 order %v, want %v", secondRound, want2)
		}
	}
}

func TestReceiptSealsAndSamplesOrdered(t *testing.T) {
	rcpt, err := Run(context.Background(), experiment3(time.Microsecond, time.Microsecond, time.Microsecond))
	if err != nil {
		t.Fatal(err)
	}
	if err := rcpt.Verify(); err != nil {
		t.Fatalf("receipt does not verify: %v", err)
	}
	last := -1
	for _, s := range rcpt.Samples {
		if s.Order <= last {
			t.Fatalf("samples out of execution order: %d after %d", s.Order, last)
		}
		last = s.Order
	}
}

func TestTooFewArmsIsAnError(t *testing.T) {
	x := experiment3(time.Millisecond, time.Millisecond, time.Millisecond)
	x.Arms = x.Arms[:1]
	if _, err := Run(context.Background(), x); err == nil {
		t.Fatal("single-arm experiment accepted")
	}
}

func TestWarmupNotRecorded(t *testing.T) {
	x := experiment3(time.Microsecond, time.Microsecond, time.Microsecond)
	x.Warmup = 2
	x.Rounds = 3
	rcpt, err := Run(context.Background(), x)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(rcpt.Samples), 3*3; got != want {
		t.Fatalf("recorded %d samples, want %d", got, want)
	}
}
