//go:build darwin

package ane

import (
	"testing"

	"github.com/tmc/apple/x/ane/mil"
)

// TestShareSurfaceReachesEval checks that a surface shared between two models
// is the one the second model's Eval actually reads.
//
// This is not the same as checking that ShareSurface assigned something. A
// Model's ANERequest holds ANEIOSurfaceObject wrappers built at compile time,
// so an implementation that only rewrote the Go slice would pass every
// structural check — InputSurface would report the shared surface, and the
// engine would still read the one it was compiled with. The test therefore
// looks at the numbers: the second model must produce the first model's result,
// which it can only do by having read the first model's output surface.
func TestShareSurfaceReachesEval(t *testing.T) {
	c, err := Open()
	if err != nil {
		t.Skipf("no ANE client: %v", err)
	}
	defer c.Close()

	const channels, spatial = 8, 4
	identity := func() *Model {
		t.Helper()
		blob, err := mil.BuildIdentityWeightBlob(channels)
		if err != nil {
			t.Fatal(err)
		}
		m, err := c.Compile(CompileOptions{
			ModelType:  ModelTypeMIL,
			MILText:    []byte(mil.GenIdentityFP16IO(channels, spatial)),
			WeightBlob: blob,
		})
		if err != nil {
			t.Fatal(err)
		}
		return m
	}

	src := identity()
	defer src.Close()
	dst := identity()
	defer dst.Close()

	// Distinct, non-zero, and not equal to whatever an unwritten surface holds.
	input := make([]float32, channels*spatial)
	for i := range input {
		input[i] = float32(i) + 1
	}
	if err := src.WriteInputFP16(0, input); err != nil {
		t.Fatal(err)
	}

	// Negative control: before sharing, dst has its own untouched input, so
	// running the pair must not reproduce the input. Without this the positive
	// result below could come from dst having held the right values already.
	if err := src.Eval(); err != nil {
		t.Fatal(err)
	}
	if err := dst.Eval(); err != nil {
		t.Fatal(err)
	}
	unshared := make([]float32, len(input))
	if err := dst.ReadOutputFP16(0, unshared); err != nil {
		t.Fatal(err)
	}
	if equalFP16(unshared, input) {
		t.Fatal("the second model reproduced the input before the surfaces were shared, so this test cannot tell whether sharing did anything")
	}

	if err := ShareSurface(src, 0, dst, 0); err != nil {
		t.Fatal(err)
	}
	if dst.InputSurface(0) != src.OutputSurface(0) {
		t.Fatal("ShareSurface did not point the destination input at the source output")
	}

	if err := src.Eval(); err != nil {
		t.Fatal(err)
	}
	if err := dst.Eval(); err != nil {
		t.Fatal(err)
	}
	got := make([]float32, len(input))
	if err := dst.ReadOutputFP16(0, got); err != nil {
		t.Fatal(err)
	}
	if !equalFP16(got, input) {
		t.Errorf("after ShareSurface the second model did not read the first model's output\ngot  %v\nwant %v", got, input)
	}

	// Rebinding repeatedly must stay correct. Each call builds a fresh request
	// and hands the Model's single retain over to it; an implementation that
	// only retained would leak every request but the last, and one that
	// released the wrong one would fault here rather than merely leak.
	for i := range 8 {
		if err := ShareSurface(src, 0, dst, 0); err != nil {
			t.Fatalf("rebind %d: %v", i, err)
		}
		if err := src.Eval(); err != nil {
			t.Fatalf("rebind %d: eval src: %v", i, err)
		}
		if err := dst.Eval(); err != nil {
			t.Fatalf("rebind %d: eval dst: %v", i, err)
		}
	}
	if err := dst.ReadOutputFP16(0, got); err != nil {
		t.Fatal(err)
	}
	if !equalFP16(got, input) {
		t.Errorf("after repeated ShareSurface calls the second model no longer reads the first model's output\ngot  %v\nwant %v", got, input)
	}
}

// equalFP16 reports whether two activations agree to fp16 precision. Both
// models here are identities, so the only loss is the round trip through fp16.
func equalFP16(got, want []float32) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		diff := got[i] - want[i]
		if diff < 0 {
			diff = -diff
		}
		if diff > 0.01*abs32(want[i])+0.01 {
			return false
		}
	}
	return true
}

func abs32(v float32) float32 {
	if v < 0 {
		return -v
	}
	return v
}
