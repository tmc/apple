//go:build darwin

package mil_test

import (
	"math"
	"testing"

	"github.com/tmc/apple/x/ane"
	"github.com/tmc/apple/x/ane/mil"
)

// TestGenRMSNormValues checks that GenRMSNorm computes RMS normalization and not
// merely that it compiles. The weights differ per channel and the input differs
// per spatial position, so a program that broadcast the weight along the wrong
// axis, or reduced over the wrong axis, produces a different answer rather than
// the same one by symmetry.
func TestGenRMSNormValues(t *testing.T) {
	c := openOrSkip(t)
	defer c.Close()

	const channels, spatial = 4, 2
	const eps = 1e-6

	weights := []float32{1, 2, 3, 4}
	// Channel-first: x[c*spatial+s]. Column 0 varies by channel, column 1 is
	// constant, so the two columns have different norms.
	input := []float32{
		1, 1,
		2, 1,
		3, 1,
		4, 1,
	}

	blob, err := mil.BuildWeightBlobV1(weights)
	if err != nil {
		t.Fatal(err)
	}
	m, err := c.Compile(ane.CompileOptions{
		ModelType:  ane.ModelTypeMIL,
		MILText:    []byte(mil.GenRMSNorm(channels, spatial, eps)),
		WeightBlob: blob,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer m.Close()

	if err := m.WriteInputFP16(0, input); err != nil {
		t.Fatal(err)
	}
	if err := m.Eval(); err != nil {
		t.Fatal(err)
	}
	got := make([]float32, channels*spatial)
	if err := m.ReadOutputFP16(0, got); err != nil {
		t.Fatal(err)
	}

	want := rmsNormReference(input, weights, channels, spatial, eps)

	// The reference must not be constant along either axis, or a wrong
	// reduction or broadcast axis would match it anyway.
	if want[0*spatial+0] == want[1*spatial+0] {
		t.Fatal("reference is constant across channels; the test cannot detect a wrong broadcast axis")
	}
	if want[0*spatial+0] == want[0*spatial+1] {
		t.Fatal("reference is constant across spatial positions; the test cannot detect a wrong reduction axis")
	}

	const tol = 0.01
	for i := range want {
		if d := math.Abs(float64(got[i] - want[i])); d > tol {
			t.Errorf("out[%d] = %v, want %v (diff %v)", i, got[i], want[i], d)
		}
	}
	t.Logf("got  %v", got)
	t.Logf("want %v", want)
}

// rmsNormReference computes RMS normalization over the channel axis in float64.
func rmsNormReference(x, w []float32, channels, spatial int, eps float64) []float32 {
	out := make([]float32, channels*spatial)
	for s := range spatial {
		var sum float64
		for c := range channels {
			v := float64(x[c*spatial+s])
			sum += v * v
		}
		rms := math.Sqrt(sum/float64(channels) + eps)
		for c := range channels {
			out[c*spatial+s] = float32(float64(x[c*spatial+s]) / rms * float64(w[c]))
		}
	}
	return out
}
