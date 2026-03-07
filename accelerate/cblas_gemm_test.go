//go:build darwin

package accelerate

import "testing"

func TestCblasSgemm1x1(t *testing.T) {
	a := []float32{3}
	b := []float32{4}
	c := []float32{5}

	Cblas_sgemm(
		CblasRowMajor,
		CblasNoTrans,
		CblasNoTrans,
		1, 1, 1,
		2.0, a, 1,
		b, 1,
		1.0, c, 1,
	)

	const want = float32(29) // 2*(3*4) + 1*5
	if !almostEqual32(c[0], want) {
		t.Fatalf("c[0] = %v, want %v", c[0], want)
	}
}

func TestCblasSgemm2x3x3x2(t *testing.T) {
	a := []float32{
		1, 2, 3,
		4, 5, 6,
	}
	b := []float32{
		7, 8,
		9, 10,
		11, 12,
	}
	c := make([]float32, 4)

	Cblas_sgemm(
		CblasRowMajor,
		CblasNoTrans,
		CblasNoTrans,
		2, 2, 3,
		1.0, a, 3,
		b, 2,
		0.0, c, 2,
	)

	want := []float32{
		58, 64,
		139, 154,
	}
	for i := range want {
		if !almostEqual32(c[i], want[i]) {
			t.Fatalf("c[%d] = %v, want %v", i, c[i], want[i])
		}
	}
}

func TestCblasSgemmTransposeA(t *testing.T) {
	// A is stored as k x m (3x2); TRANSA tells BLAS to use A^T as 2x3.
	a := []float32{
		1, 4,
		2, 5,
		3, 6,
	}
	b := []float32{
		7, 8,
		9, 10,
		11, 12,
	}
	c := make([]float32, 4)

	Cblas_sgemm(
		CblasRowMajor,
		CblasTrans,
		CblasNoTrans,
		2, 2, 3,
		1.0, a, 2,
		b, 2,
		0.0, c, 2,
	)

	want := []float32{
		58, 64,
		139, 154,
	}
	for i := range want {
		if !almostEqual32(c[i], want[i]) {
			t.Fatalf("c[%d] = %v, want %v", i, c[i], want[i])
		}
	}
}

func almostEqual32(got, want float32) bool {
	d := got - want
	if d < 0 {
		d = -d
	}
	return d <= 1e-4
}
