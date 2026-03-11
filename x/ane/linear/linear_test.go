//go:build darwin

package linear

import (
	"testing"

	"github.com/tmc/apple/x/ane"
)

func TestChannelFirstRoundTrip(t *testing.T) {
	// [batch=2, dim=3]
	in := []float32{1, 2, 3, 4, 5, 6}
	cf := toChannelFirst(in, 2, 3)
	got := fromChannelFirst(cf, 2, 3)
	for i, v := range in {
		if got[i] != v {
			t.Errorf("round-trip mismatch at %d: %v != %v", i, got[i], v)
		}
	}
}

func TestLinearCPU(t *testing.T) {
	// x = [1, 2], w = [[1, 0], [0, 1]] → y = [1, 2]
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
	rt, err := ane.Open()
	if err != nil {
		t.Skip("no ANE:", err)
	}
	defer rt.Close()

	exec := New(rt)
	defer exec.Close()

	// Identity matrix: W = I_4
	w := []float32{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	x := []float32{1, 2, 3, 4}

	y, err := exec.Linear(x, w, 1, 4, 4)
	if err != nil {
		t.Fatal(err)
	}

	diff := maxAbsDiff(x, y)
	if diff > 0.01 {
		t.Errorf("identity linear: maxAbsDiff = %v, want < 0.01\ngot: %v", diff, y)
	}
}

func TestLinearCaching(t *testing.T) {
	rt, err := ane.Open()
	if err != nil {
		t.Skip("no ANE:", err)
	}
	defer rt.Close()

	exec := New(rt)
	defer exec.Close()

	w := []float32{1, 0, 0, 1}
	x := []float32{5, 6}

	// First call compiles.
	_, err = exec.Linear(x, w, 1, 2, 2)
	if err != nil {
		t.Fatal(err)
	}

	s1 := exec.Stats()
	if s1.CachedKernels != 1 {
		t.Errorf("cached kernels after first call = %d, want 1", s1.CachedKernels)
	}

	// Second call with same weights should hit cache.
	_, err = exec.Linear(x, w, 1, 2, 2)
	if err != nil {
		t.Fatal(err)
	}

	s2 := exec.Stats()
	if s2.CachedKernels != 1 {
		t.Errorf("cached kernels after second call = %d, want 1", s2.CachedKernels)
	}
}
