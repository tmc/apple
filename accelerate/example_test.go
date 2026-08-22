//go:build darwin

package accelerate_test

import (
	"fmt"

	"github.com/tmc/apple/accelerate"
)

func Example() {
	a := []float32{3}
	b := []float32{4}
	c := []float32{5}

	accelerate.Cblas_sgemm(
		accelerate.CblasRowMajor,
		accelerate.CblasNoTrans,
		accelerate.CblasNoTrans,
		1, 1, 1,
		2.0, a, 1,
		b, 1,
		1.0, c, 1,
	)

	fmt.Printf("sgemm result: %.0f\n", c[0])
	fmt.Println("order:", accelerate.CblasRowMajor)

	// Output:
	// sgemm result: 29
	// order: CblasRowMajor
}
