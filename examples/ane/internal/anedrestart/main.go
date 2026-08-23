// Command anedrestart reopens an E5RT bundle compiled under a previous aned
// instance and checks the answer, to measure whether bundle reuse survives a
// restart of the daemon.
//
// bundlereuse leaves that arm explicitly UNMEASURED. This is the arm.
//
// # Result
//
// Reuse SURVIVES an aned restart. A bundle compiled under one daemon instance
// reopens and computes correctly under the next, with the compiling instance
// provably gone: the launchd runs counter had incremented and the pid differed.
// Measured on macOS 26.x (build 25G76) with two bundles, against a float64 CPU
// reference and a mutation control.
//
// The same run found something the restart question was hiding. Two bundles
// roughly a day old failed to create an operation at all, status 13, while
// bundles minutes old succeeded in the same processes, with aned both up and
// down. So bundles are NOT durable, and what bounds their life is not the
// daemon's lifetime. The boundary is UNMEASURED; the observation is two bundles
// at about 24 hours, which is one point, not a half-life.
//
// The two bundles are structurally identical -- same files, same 3888-byte
// payload differing in seven bytes -- but name different artifacts in
// model.anehash. That is consistent with a bundle being a reference into a
// store this program cannot inspect, and it is inference, not a finding.
//
// # Why openability and correctness are reported separately
//
// OperationCreatePrecompiled runs before any port is bound, so it cannot fail
// from a wrong -in/-out/-spatial. Openability is therefore the restart answer
// on its own, and the CPU reference is a separate, stricter claim that only
// applies when the bundle really is a single convolution of the given shape.
// Reporting them together would let a shape guess masquerade as a reuse
// failure, which is precisely what the two-convolution bundle in this family
// did on the first attempt. Pass -check=false for bundles this reference does
// not model.
package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"os/exec"
	"strings"

	"github.com/tmc/apple/x/ane/e5rt"
)

var (
	bundle  = flag.String("bundle", "", "path to a bundle compiled under an earlier aned instance")
	inCh    = flag.Int("in", 16, "input channels the bundle was compiled for")
	outCh   = flag.Int("out", 16, "output channels the bundle was compiled for")
	spatial = flag.Int("spatial", 8, "spatial width the bundle was compiled for")
	check   = flag.Bool("check", true, "compare the result against a single-convolution CPU reference")
)

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "FAIL:", err)
		os.Exit(1)
	}
	fmt.Println("\nOK")
}

func run() error {
	if *bundle == "" {
		return fmt.Errorf("-bundle is required")
	}
	info, err := os.Stat(*bundle)
	if err != nil {
		return err
	}
	fmt.Printf("bundle:  %s\n", *bundle)
	fmt.Printf("  compiled (mtime): %s\n", info.ModTime().Format("2006-01-02 15:04:05"))
	fmt.Printf("  aned before open: %s\n", anedState())

	// The weights are the ones e5rtdispatch generates, reconstructed rather
	// than read: the model directory that produced this bundle is long gone,
	// which is the point.
	weights := make([]float32, *outCh**inCh)
	for i := range weights {
		weights[i] = float32(i%7-3) * 0.25
	}
	input := make([]float32, *inCh**spatial)
	for i := range input {
		input[i] = float32(i%5-2) * 0.25
	}
	want := convReference(weights, input, *outCh, *inCh, *spatial)

	p, err := e5rt.OpenBundle(e5rt.BundleOptions{
		BundlePath: *bundle,
		Inputs:     []e5rt.Port{{Name: "x", Size: *inCh * *spatial * 2}},
		Outputs:    []e5rt.Port{{Name: "y", Size: *outCh * *spatial * 2}},
	})
	if err != nil {
		return fmt.Errorf("open the bundle: %w", err)
	}
	defer p.Close()
	fmt.Printf("  aned after open:  %s\n", anedState())
	fmt.Println("  OPENED: the bundle produced a usable operation")

	// Openability and correctness are reported separately and on purpose. The
	// restart question is answered by whether the operation can be created at
	// all -- OperationCreatePrecompiled runs before any port is bound, so it
	// cannot fail from a wrong -in/-out/-spatial. Mixing the two would let a
	// shape guess masquerade as a reuse failure, which is exactly what the
	// two-convolution bundle in this family would have done.
	if !*check {
		fmt.Println("  correctness: SKIPPED (-check=false); this run measures openability only")
		return nil
	}

	in, err := p.Input("x")
	if err != nil {
		return err
	}
	if err := in.WriteFP16(input); err != nil {
		return err
	}
	if err := p.Execute(); err != nil {
		return fmt.Errorf("execute: %w", err)
	}
	got := make([]float32, *outCh**spatial)
	out, err := p.Output("y")
	if err != nil {
		return err
	}
	if err := out.ReadFP16(got); err != nil {
		return err
	}
	diff, err := compare(got, want)
	if err != nil {
		return err
	}
	fmt.Printf("  reopened vs float64 CPU reference: max diff %.4g over %d values\n", diff, len(got))
	if diff > 0.05 {
		return fmt.Errorf("the reopened bundle opened but disagrees with the single-convolution reference by %g; "+
			"this is a statement about the reference, not about reuse -- the bundle already opened. Re-run with "+
			"-check=false if this bundle is not a single convolution of the given shape", diff)
	}

	// Mutation control: a second execution with a different input must move
	// and must match its own reference. Without it, a stale or constant
	// buffer would satisfy the check above.
	for i := range input {
		input[i] = -input[i]
	}
	want2 := convReference(weights, input, *outCh, *inCh, *spatial)
	if err := in.WriteFP16(input); err != nil {
		return err
	}
	if err := p.Execute(); err != nil {
		return err
	}
	got2 := make([]float32, len(got))
	if err := out.ReadFP16(got2); err != nil {
		return err
	}
	diff2, err := compare(got2, want2)
	if err != nil {
		return err
	}
	moved, err := compare(got2, got)
	if err != nil {
		return err
	}
	if moved == 0 {
		return fmt.Errorf("negating the input did not move the output; the result may be stale")
	}
	fmt.Printf("  mutation control: negated input moved the output by %.4g and matches its own reference (%.4g)\n", moved, diff2)
	if diff2 > 0.05 {
		return fmt.Errorf("the moved output does not match its own reference (%g)", diff2)
	}
	fmt.Printf("  aned at end:      %s\n", anedState())
	return nil
}

// anedState reports whether the daemon is running, so the evidence for a
// restart claim does not depend on a note taken by hand elsewhere.
func anedState() string {
	out, err := exec.Command("launchctl", "print", "system/com.apple.aned").CombinedOutput()
	if err != nil {
		return fmt.Sprintf("UNKNOWN (%v)", err)
	}
	state, runs, pid := "?", "?", ""
	for line := range strings.SplitSeq(string(out), "\n") {
		f := strings.TrimSpace(line)
		if s, ok := strings.CutPrefix(f, "state = "); ok && state == "?" {
			state = s
		}
		if s, ok := strings.CutPrefix(f, "runs = "); ok && runs == "?" {
			runs = s
		}
		if s, ok := strings.CutPrefix(f, "pid = "); ok && pid == "" {
			pid = s
		}
	}
	return fmt.Sprintf("state=%s runs=%s pid=%s", state, runs, pid)
}

func convReference(w, x []float32, outCh, inCh, spatial int) []float32 {
	y := make([]float32, outCh*spatial)
	for o := range outCh {
		for s := range spatial {
			var acc float64
			for i := range inCh {
				acc += float64(w[o*inCh+i]) * float64(x[i*spatial+s])
			}
			y[o*spatial+s] = float32(acc)
		}
	}
	return y
}

func compare(got, want []float32) (float64, error) {
	if len(got) != len(want) {
		return 0, fmt.Errorf("comparing %d values against %d", len(got), len(want))
	}
	if len(got) == 0 {
		return 0, fmt.Errorf("nothing to compare")
	}
	var maxDiff float64
	for i := range got {
		g, w := float64(got[i]), float64(want[i])
		if math.IsNaN(g) || math.IsInf(g, 0) {
			return 0, fmt.Errorf("result element %d is %v", i, g)
		}
		if d := math.Abs(g - w); d > maxDiff {
			maxDiff = d
		}
	}
	return maxDiff, nil
}
