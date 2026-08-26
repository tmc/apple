// Command mlxperf is the entry point for the mlxperf optimization arena.
//
// The arena measures a candidate against an immutable baseline on one host,
// in one session, with paired interleaved arms, a mandatory identical
// duplicate control, and output-digest correctness gates. It writes sealed
// experiment receipts; it never trusts candidate-claimed scores.
//
// Usage:
//
//	mlxperf arena run -workload ID -work-units N -rounds N [-warmup N]
//	    [-rotate] -baseline-cmd argv -candidate-cmd argv
//	    [-duplicate-cmd argv] [-expect DIGEST] -out receipt.json
//
//	mlxperf arena compare a.json b.json ...
//
// The arena commands are portable: they run anywhere Go runs, including
// hosts without a GPU. A worker command that fails, prints output whose
// digest does not match -expect, or reports unequal work produces a typed
// refusal in the receipt instead of a timing verdict.
package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/tmc/apple/x/armbench"
	"github.com/tmc/apple/x/experiment"
)

func main() {
	os.Exit(run(context.Background(), os.Args[1:], os.Stdout, os.Stderr))
}

func usage() string {
	return "usage: mlxperf arena <run|compare> [flags]\n"
}

func run(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) < 2 || args[0] != "arena" {
		fmt.Fprint(stderr, usage())
		return 2
	}
	switch args[1] {
	case "run":
		fs := newRunFlags()
		fs.Parse(args[2:])
		return cmdArenaRun(ctx, fs, fs.Args(), stdout, stderr)
	case "compare":
		fs := flag.NewFlagSet("compare", flag.ContinueOnError)
		fs.SetOutput(stderr)
		fs.Usage = func() { fmt.Fprint(stderr, usage()) }
		fs.Parse(args[2:])
		return cmdArenaCompare(fs.Args(), stdout, stderr)
	default:
		fmt.Fprintf(stderr, "unknown arena subcommand %q\n%s", args[1], usage())
		return 2
	}
}

type runFlags struct {
	*flag.FlagSet
	workloadID   string
	workUnits    int64
	rounds       int
	warmup       int
	rotate       bool
	expect       string
	baselineCmd  string
	candidateCmd string
	duplicateCmd string
	out          string
}

func newRunFlags() *runFlags {
	fs := &runFlags{FlagSet: flag.NewFlagSet("arena run", flag.ContinueOnError)}
	fs.SetOutput(os.Stderr)
	fs.StringVar(&fs.workloadID, "workload", "", "workload identity (required)")
	fs.Int64Var(&fs.workUnits, "work-units", 0, "declared equal work each arm performs per execution")
	fs.IntVar(&fs.rounds, "rounds", 5, "recorded rounds per arm")
	fs.IntVar(&fs.warmup, "warmup", 1, "unrecorded warmup executions per arm")
	fs.BoolVar(&fs.rotate, "rotate", true, "rotate arm order each round")
	fs.StringVar(&fs.expect, "expect", "", "sha256 hex digest every arm's stdout must match")
	fs.StringVar(&fs.baselineCmd, "baseline-cmd", "", "baseline arm command line (required)")
	fs.StringVar(&fs.candidateCmd, "candidate-cmd", "", "candidate arm command line (required)")
	fs.StringVar(&fs.duplicateCmd, "duplicate-cmd", "", "duplicate-control command line (defaults to baseline)")
	fs.StringVar(&fs.out, "out", "", "write the sealed receipt to this path")
	return fs
}

func splitCmd(s string) []string { return strings.Fields(s) }

func digestReader(b []byte) string {
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}

func commandArm(name string, kind experiment.ArmKind, workUnits int64, expectDigest string, cmdline string) armbench.Arm {
	return armbench.Arm{
		Name:      name,
		Kind:      kind,
		WorkUnits: workUnits,
		Fn: func(ctx context.Context) (armbench.Result, error) {
			argv := splitCmd(cmdline)
			if len(argv) == 0 {
				return armbench.Result{}, fmt.Errorf("empty command")
			}
			start := time.Now()
			out, err := exec.CommandContext(ctx, argv[0], argv[1:]...).Output()
			elapsed := time.Since(start)
			res := armbench.Result{Elapsed: elapsed}
			if err != nil {
				return res, err
			}
			got := digestReader(out)
			if expectDigest != "" && got != expectDigest {
				return res, fmt.Errorf("output digest %s does not match expected %s", got[:16], expectDigest[:16])
			}
			res.Output = got
			return res, nil
		},
	}
}

func cmdArenaRun(ctx context.Context, fs *runFlags, rest []string, stdout, stderr io.Writer) int {
	if fs.workloadID == "" || fs.baselineCmd == "" || fs.candidateCmd == "" || len(rest) > 0 {
		fmt.Fprint(stderr, usage())
		return 2
	}
	dupCmd := fs.duplicateCmd
	if dupCmd == "" {
		dupCmd = fs.baselineCmd
	}

	x := armbench.Experiment{
		Warmup: fs.warmup,
		Rounds: fs.rounds,
		Rotate: fs.rotate,
		Workload: experiment.Workload{
			ID:        fs.workloadID,
			WorkUnits: fs.workUnits,
		},
		Arms: []armbench.Arm{
			commandArm("baseline", experiment.ArmBaseline, fs.workUnits, fs.expect, fs.baselineCmd),
			commandArm("candidate", experiment.ArmCandidate, fs.workUnits, fs.expect, fs.candidateCmd),
			commandArm("dup", experiment.ArmDuplicate, fs.workUnits, fs.expect, dupCmd),
		},
	}
	rcpt, err := armbench.Run(ctx, x)
	if err != nil {
		fmt.Fprintln(stderr, "mlxperf:", err)
		return 2
	}

	printReceipt(rcpt, stdout)

	b, err := rcpt.Marshal()
	if err != nil {
		fmt.Fprintln(stderr, "mlxperf:", err)
		return 1
	}
	if fs.out == "" {
		os.Stdout.Write(append(b, '\n'))
	} else if err := os.WriteFile(fs.out, append(b, '\n'), 0o644); err != nil {
		fmt.Fprintln(stderr, "mlxperf:", err)
		return 1
	}
	if rcpt.Refused() {
		return 1
	}
	return 0
}

func printReceipt(rcpt experiment.Receipt, w io.Writer) {
	med := armbench.Medians(&rcpt)
	names := make([]string, 0, len(med))
	for name := range med {
		names = append(names, name)
	}
	sortStrings(names)
	fmt.Fprintf(w, "workload %s (%d work units)\n", rcpt.Workload.ID, rcpt.Workload.WorkUnits)
	for _, name := range names {
		fmt.Fprintf(w, "  %-10s median %s\n", name, med[name])
	}
	fmt.Fprintf(w, "verdict %s\n", rcpt.Verdict)
	if rcpt.Winner != "" {
		fmt.Fprintf(w, "winner %s\n", rcpt.Winner)
	}
	for _, ref := range rcpt.Refusals {
		fmt.Fprintf(w, "refusal %s\n", ref.Error())
	}
}

func sortStrings(s []string) {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
}

func cmdArenaCompare(paths []string, stdout, stderr io.Writer) int {
	if len(paths) < 2 {
		fmt.Fprintln(stderr, "compare needs at least two receipts")
		return 2
	}
	rcpts := make([]*experiment.Receipt, 0, len(paths))
	for _, p := range paths {
		b, err := os.ReadFile(p)
		if err != nil {
			fmt.Fprintln(stderr, "mlxperf:", err)
			return 1
		}
		r, err := experiment.ParseReceipt(b)
		if err != nil {
			fmt.Fprintf(stderr, "mlxperf: %s: %v\n", p, err)
			return 1
		}
		rcpts = append(rcpts, r)
	}

	// Receipts for different workloads are never comparable.
	base := rcpts[0].Workload.ID
	for _, r := range rcpts[1:] {
		if r.Workload.ID != base {
			fmt.Fprintf(stdout, "refusal %s\n", experiment.RefusalIdentityDrift)
			fmt.Fprintf(stdout, "detail workload %q is not comparable with %q\n", base, r.Workload.ID)
			return 1
		}
	}

	for _, r := range rcpts {
		fmt.Fprintf(stdout, "%s: %s (digest %s)\n", shortPath(r), r.Verdict, r.Digest[:12])
		for _, ref := range r.Refusals {
			fmt.Fprintf(stdout, "  refusal %s\n", ref.Error())
		}
	}
	for i, r := range rcpts {
		if r.Refused() {
			continue
		}
		for j, o := range rcpts {
			if i == j || o.Refused() {
				continue
			}
			mi := armbench.Medians(r)["candidate"]
			mo := armbench.Medians(o)["candidate"]
			if mi == 0 || mo == 0 {
				continue
			}
			delta := 100 * float64(mo-mi) / float64(mo)
			label := "candidate vs baseline"
			fmt.Fprintf(stdout, "%s %s: %+.1f%%\n", shortPath(r), label, delta)
		}
		break // only the first non-refused receipt anchors the comparison
	}
	return 0
}

func shortPath(r *experiment.Receipt) string {
	if r.Label != "" {
		return r.Label
	}
	return r.CreatedAt.Format(time.RFC3339)
}
