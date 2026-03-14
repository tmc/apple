//go:build darwin

package linear

import (
	"context"
	"math"
	"testing"
)

func TestChannelFirstRoundTrip(t *testing.T) {
	// [batch=2, dim=3]
	in := []float32{1, 2, 3, 4, 5, 6}
	cf := make([]float32, len(in))
	toChannelFirst(cf, in, 2, 3)
	got := make([]float32, len(in))
	fromChannelFirst(got, cf, 2, 3)
	for i, v := range in {
		if got[i] != v {
			t.Errorf("round-trip mismatch at %d: %v != %v", i, got[i], v)
		}
	}
}

func TestLinearCPU(t *testing.T) {
	// x = [1, 2], w = [[1, 0], [0, 1]] -> y = [1, 2]
	x := []float32{1, 2}
	w := []float32{1, 0, 0, 1}
	y := linearCPU(x, w, 1, 2, 2)
	want := []float32{1, 2}
	for i := range want {
		if y[i] != want[i] {
			t.Errorf("linearCPU[%d] = %v, want %v", i, y[i], want[i])
		}
	}
}

func TestLinearIdentity(t *testing.T) {
	exec := New(Options{})
	defer exec.Close()

	// Identity matrix: W = I_4
	w := []float32{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	x := []float32{1, 2, 3, 4}

	y, err := exec.Linear(context.Background(), x, w, 1, 4, 4)
	if err != nil {
		t.Fatal(err)
	}

	diff := maxAbsDiff(x, y)
	if diff > 0.01 {
		t.Errorf("identity linear: maxAbsDiff = %v, want < 0.01\ngot: %v", diff, y)
	}
}

func TestLinearCaching(t *testing.T) {
	exec := New(Options{})
	defer exec.Close()

	w := []float32{1, 0, 0, 1}
	x := []float32{5, 6}

	// First call compiles.
	_, err := exec.Linear(context.Background(), x, w, 1, 2, 2)
	if err != nil {
		t.Fatal(err)
	}

	s1 := exec.Stats()
	if s1.Kernels != 1 {
		t.Errorf("cached kernels after first call = %d, want 1", s1.Kernels)
	}

	// Second call with same weights should hit cache.
	_, err = exec.Linear(context.Background(), x, w, 1, 2, 2)
	if err != nil {
		t.Fatal(err)
	}

	s2 := exec.Stats()
	if s2.Kernels != 1 {
		t.Errorf("cached kernels after second call = %d, want 1", s2.Kernels)
	}
	if s2.CacheHits < 1 {
		t.Errorf("cache hits after second call = %d, want >= 1", s2.CacheHits)
	}
}

// linearCPU computes y = x @ W^T on the CPU for comparison.
func linearCPU(x, w []float32, batch, inDim, outDim int) []float32 {
	y := make([]float32, batch*outDim)
	for b := 0; b < batch; b++ {
		for o := 0; o < outDim; o++ {
			var sum float64
			for i := 0; i < inDim; i++ {
				sum += float64(x[b*inDim+i]) * float64(w[o*inDim+i])
			}
			y[b*outDim+o] = float32(sum)
		}
	}
	return y
}

// maxAbsDiff returns the maximum absolute difference between two slices.
func maxAbsDiff(a, b []float32) float32 {
	var max float32
	for i := range a {
		d := float32(math.Abs(float64(a[i] - b[i])))
		if d > max {
			max = d
		}
	}
	return max
}
