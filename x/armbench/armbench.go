// Package armbench measures benchmark arms with the attribution controls the
// mlxperf arena requires: interleaved execution, a mandatory identical
// duplicate-control arm, declared work equality, physics checks on reported
// byte traffic, and host qualification.
//
// The harness prevents common attribution errors; it cannot make an
// uncontrolled host or semantically unequal arms comparable. Every experiment
// it runs produces a sealed [experiment.Receipt]. A receipt with refusals has
// no performance verdict — a refusal replaces a score, never annotates one.
//
// The minimum experiment measures two identical functions plus their
// duplicate and requires that no winner be claimed:
//
//	work := func(context.Context) (armbench.Result, error) {
//		start := time.Now()
//		sum := 0
//		for i := 0; i < 100000; i++ {
//			sum += i
//		}
//		return armbench.Result{Elapsed: time.Since(start)}, nil
//	}
//	rcpt := armbench.Run(ctx, armbench.Experiment{
//		Rounds: 20,
//		Workload: experiment.Workload{ID: "spin/100k", WorkUnits: 100000},
//		Arms: []armbench.Arm{
//			{Name: "baseline", Kind: experiment.ArmBaseline, WorkUnits: 100000, Fn: work},
//			{Name: "candidate", Kind: experiment.ArmCandidate, WorkUnits: 100000, Fn: work},
//			{Name: "dup", Kind: experiment.ArmDuplicate, WorkUnits: 100000, Fn: work},
//		},
//	})
//	// rcpt.Verdict == experiment.VerdictNoWinner for identical arms.
package armbench

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/tmc/apple/x/experiment"
)

// Result is what an arm function reports for one execution. The harness
// records it verbatim; it does not trust anything the arm claims about its
// own relative speed.
type Result struct {
	Elapsed         time.Duration
	BytesByLocation map[string]int64 // byte traffic by named location
	Output          string           // output digest or identity, when declared
}

// Arm is one named competitor in an experiment.
type Arm struct {
	Name      string
	Kind      experiment.ArmKind
	WorkUnits int64 // work units this arm performs per execution
	Fn        func(ctx context.Context) (Result, error)
}

// Experiment declares a complete measurement.
type Experiment struct {
	Workload experiment.Workload
	Label    string
	Warmup   int                             // unrecorded warmup executions per arm
	Rounds   int                             // recorded rounds per arm
	Rotate   bool                            // rotate starting arm each round instead of fixed order
	Qualify  func(ctx context.Context) error // host qualification; error contaminates the run

	// MaxByteRate is the largest physically possible byte rate in bytes per
	// second on this machine. A sample reporting more than this is refused as
	// impossible. Zero disables the check.
	MaxByteRate float64

	Arms []Arm
}

// Run executes the experiment and returns its sealed receipt. It returns a
// non-nil error only for experiments that are not constructible at all (no
// arms); every other problem is recorded as a refusal in the receipt.
func Run(ctx context.Context, x Experiment) (experiment.Receipt, error) {
	rcpt := experiment.Receipt{
		Schema:   experiment.Schema,
		Workload: x.Workload,
		Label:    x.Label,
	}
	refuse := func(reason, format string, args ...any) {
		rcpt.Refusals = append(rcpt.Refusals, experiment.Refusal{Reason: reason, Detail: fmt.Sprintf(format, args...)})
	}

	if len(x.Arms) < 2 {
		return rcpt, fmt.Errorf("armbench: experiment needs at least two arms, have %d", len(x.Arms))
	}

	// Declared work must be equal across arms before anything executes.
	var units int64
	for i, arm := range x.Arms {
		if i == 0 {
			units = arm.WorkUnits
			continue
		}
		if arm.WorkUnits != units {
			refuse(experiment.RefusalUnequalWork,
				"arm %q declares %d work units, arm %q declares %d",
				x.Arms[0].Name, units, arm.Name, arm.WorkUnits)
		}
	}

	hasControl := false
	for _, arm := range x.Arms {
		if arm.Kind == experiment.ArmDuplicate {
			hasControl = true
		}
	}
	if !hasControl {
		refuse(experiment.RefusalMissingControl, "no %s arm declared", experiment.ArmDuplicate)
	}

	if x.Rounds < 1 {
		refuse(experiment.RefusalInsufficientSamples, "rounds=%d", x.Rounds)
	}

	// Host qualification gates measurement.
	if x.Qualify != nil {
		if err := x.Qualify(ctx); err != nil {
			refuse(experiment.RefusalContaminatedHost, "qualification failed: %v", err)
		}
	}

	order := 0
	for _, arm := range x.Arms {
		for w := 0; w < x.Warmup; w++ {
			if _, err := arm.Fn(ctx); err != nil {
				refuse(experiment.RefusalCorrectnessFailed, "warmup of %q: %v", arm.Name, err)
				break
			}
		}
	}
	if fatal(&rcpt) || x.Rounds < 1 {
		if len(rcpt.Samples) > 0 {
			analyze(&rcpt)
		}
		rcpt = finish(rcpt)
		if err := rcpt.Seal(); err != nil {
			return rcpt, err
		}
		return rcpt, nil
	}

	for round := 0; round < x.Rounds; round++ {
		ordering := rotatedOrder(len(x.Arms), round, x.Rotate)
		for _, idx := range ordering {
			arm := x.Arms[idx]
			res, err := arm.Fn(ctx)
			order++
			s := experiment.Sample{
				Arm:       arm.Name,
				Kind:      arm.Kind,
				Round:     round,
				Order:     order,
				WorkUnits: arm.WorkUnits,
				Output:    res.Output,
				Verdict:   experiment.VerdictPass,
			}
			if err != nil {
				s.Verdict = experiment.VerdictFail
				s.Note = err.Error()
				refuse(experiment.RefusalCorrectnessFailed, "arm %q round %d: %v", arm.Name, round, err)
			} else {
				s.Elapsed = res.Elapsed
				s.Bytes = res.BytesByLocation
				if rate := byteRate(s.Elapsed, totalBytes(s.Bytes)); x.MaxByteRate > 0 && rate > x.MaxByteRate {
					s.Verdict = experiment.VerdictFail
					s.Note = fmt.Sprintf("reported %.3g B/s exceeds physical maximum %.3g B/s", rate, x.MaxByteRate)
					refuse(experiment.RefusalImpossibleByteRate, "arm %q round %d: %s", arm.Name, round, s.Note)
				}
			}
			rcpt.Samples = append(rcpt.Samples, s)
		}
	}

	analyze(&rcpt)
	rcpt = finish(rcpt)
	if err := rcpt.Seal(); err != nil {
		return rcpt, err
	}
	return rcpt, nil
}

// fatal reports whether any refusal blocks execution entirely. Only these
// reasons prevent measurement; the others are decided from recorded samples.
func fatal(r *experiment.Receipt) bool {
	for _, ref := range r.Refusals {
		switch ref.Reason {
		case experiment.RefusalUnequalWork,
			experiment.RefusalMissingControl,
			experiment.RefusalContaminatedHost,
			experiment.RefusalCorrectnessFailed:
			return true
		}
	}
	return false
}

// rotatedOrder returns the execution order of arm indexes for a round.
func rotatedOrder(n, round int, rotate bool) []int {
	start := 0
	if rotate {
		start = round % n
	}
	out := make([]int, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, (start+i)%n)
	}
	return out
}

func totalBytes(byLoc map[string]int64) int64 {
	var total int64
	for _, b := range byLoc {
		total += b
	}
	return total
}

func byteRate(elapsed time.Duration, nbytes int64) float64 {
	if elapsed <= 0 {
		return 0
	}
	return float64(nbytes) / elapsed.Seconds()
}

// analyze assigns the receipt's verdict from its samples. The rule: compare
// candidate against baseline medians; the noise floor is the duplicate
// control's gap to the baseline plus each arm's own sampling dispersion.
// No winner is declared unless the claimed improvement exceeds that floor:
// a win inside measurement noise is not a result.
func analyze(rcpt *experiment.Receipt) {
	med := Medians(rcpt)
	base, okBase := medianOfKind(rcpt, experiment.ArmBaseline)
	cand, okCand := medianOfKind(rcpt, experiment.ArmCandidate)
	ctrl, okCtrl := medianOfKind(rcpt, experiment.ArmDuplicate)

	if !okBase || !okCand || !okCtrl || len(med) == 0 {
		rcpt.Refusals = append(rcpt.Refusals, experiment.Refusal{
			Reason: experiment.RefusalInsufficientSamples,
			Detail: "analysis needs baseline, candidate, and duplicate samples",
		})
		return
	}

	baselineName, _ := nameOfKind(rcpt, experiment.ArmBaseline)
	candidateName, _ := nameOfKind(rcpt, experiment.ArmCandidate)
	duplicateName, _ := nameOfKind(rcpt, experiment.ArmDuplicate)

	delta := base - cand // positive means the candidate is faster
	disp := Dispersions(rcpt)
	noise := absDuration(base-ctrl) + disp[baselineName] + disp[duplicateName]

	switch {
	case delta > noise:
		rcpt.Verdict = experiment.VerdictFaster
		rcpt.Winner = candidateName
	case delta < -noise:
		rcpt.Verdict = experiment.VerdictSlower
		rcpt.Winner = baselineName
	default:
		rcpt.Verdict = experiment.VerdictNoWinner
	}
}

func finish(rcpt experiment.Receipt) experiment.Receipt {
	// A refusal overrides any ranking verdict: a refusal replaces a score.
	if rcpt.Refused() {
		rcpt.Verdict = experiment.VerdictRefused
		rcpt.Winner = ""
	}
	if rcpt.Verdict == "" {
		rcpt.Verdict = experiment.VerdictRefused
	}
	return rcpt
}

func medianOfKind(r *experiment.Receipt, kind experiment.ArmKind) (time.Duration, bool) {
	name, ok := nameOfKind(r, kind)
	if !ok {
		return 0, false
	}
	m, ok := Medians(r)[name]
	return m, ok
}

func nameOfKind(r *experiment.Receipt, kind experiment.ArmKind) (string, bool) {
	for _, s := range r.Samples {
		if s.Kind == kind && s.Verdict == experiment.VerdictPass {
			return s.Arm, true
		}
	}
	// Fall back to a failing sample so the refusal names its arm.
	for _, s := range r.Samples {
		if s.Kind == kind {
			return s.Arm, true
		}
	}
	return "", false
}

// Medians returns the median elapsed time per arm name across passing
// samples. Failed samples do not contribute.
func Medians(r *experiment.Receipt) map[string]time.Duration {
	grouped := make(map[string][]time.Duration)
	for _, s := range r.Samples {
		if s.Verdict == experiment.VerdictPass {
			grouped[s.Arm] = append(grouped[s.Arm], s.Elapsed)
		}
	}
	out := make(map[string]time.Duration, len(grouped))
	for name, vals := range grouped {
		out[name] = median(vals)
	}
	return out
}

func median(vals []time.Duration) time.Duration {
	if len(vals) == 0 {
		return 0
	}
	sorted := append([]time.Duration(nil), vals...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
	return sorted[len(sorted)/2]
}

// Dispersions returns the median absolute deviation of passing elapsed times
// per arm name: a robust within-arm spread used by the noise floor.
func Dispersions(r *experiment.Receipt) map[string]time.Duration {
	grouped := make(map[string][]time.Duration)
	for _, s := range r.Samples {
		if s.Verdict == experiment.VerdictPass {
			grouped[s.Arm] = append(grouped[s.Arm], s.Elapsed)
		}
	}
	out := make(map[string]time.Duration, len(grouped))
	for name, vals := range grouped {
		m := median(vals)
		devs := make([]time.Duration, len(vals))
		for i, v := range vals {
			devs[i] = absDuration(v - m)
		}
		out[name] = median(devs)
	}
	return out
}

func absDuration(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}
	return d
}
