package armbench_test

import (
	"context"
	"fmt"
	"time"

	"github.com/tmc/apple/x/armbench"
	"github.com/tmc/apple/x/experiment"
)

// ExampleRun measures two identical arms plus their duplicate control. The
// duplicate-control gap is the noise floor, so identical arms produce no
// winner.
func ExampleRun() {
	work := func(context.Context) (armbench.Result, error) {
		start := time.Now()
		sum := 0
		for i := 0; i < 100000; i++ {
			sum += i
		}
		return armbench.Result{Elapsed: time.Since(start)}, nil
	}
	rcpt, err := armbench.Run(context.Background(), armbench.Experiment{
		Warmup:   2,
		Rounds:   30,
		Workload: experiment.Workload{ID: "spin/100k", WorkUnits: 100000},
		Arms: []armbench.Arm{
			{Name: "baseline", Kind: experiment.ArmBaseline, WorkUnits: 100000, Fn: work},
			{Name: "candidate", Kind: experiment.ArmCandidate, WorkUnits: 100000, Fn: work},
			{Name: "dup", Kind: experiment.ArmDuplicate, WorkUnits: 100000, Fn: work},
		},
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(rcpt.Verdict)
	if rcpt.Refused() {
		for _, ref := range rcpt.Refusals {
			fmt.Println(ref.Reason)
		}
	}
	// Output:
	// no-winner
}
