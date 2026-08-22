//go:build darwin

package coregraphics_test

import (
	"fmt"

	"github.com/tmc/apple/coregraphics"
)

func ExampleCGAffineTransformMakeScale() {
	t := coregraphics.CGAffineTransformMakeScale(2.0, 3.0)
	fmt.Printf("A: %.1f, B: %.1f, C: %.1f, D: %.1f, Tx: %.1f, Ty: %.1f\n", t.A, t.B, t.C, t.D, t.Tx, t.Ty)

	// Output:
	// A: 2.0, B: 0.0, C: 0.0, D: 3.0, Tx: 0.0, Ty: 0.0
}

func ExampleCGColorCreateGenericRGB() {
	color := coregraphics.CGColorCreateGenericRGB(1.0, 0.5, 0.0, 1.0)
	alpha := coregraphics.CGColorGetAlpha(color)
	numComp := coregraphics.CGColorGetNumberOfComponents(color)
	fmt.Printf("RGB Alpha: %.1f, Components: %d\n", alpha, numComp)

	grayColor := coregraphics.CGColorCreateGenericGray(0.75, 1.0)
	grayAlpha := coregraphics.CGColorGetAlpha(grayColor)
	grayComp := coregraphics.CGColorGetNumberOfComponents(grayColor)
	fmt.Printf("Gray Alpha: %.1f, Components: %d\n", grayAlpha, grayComp)

	// Output:
	// RGB Alpha: 1.0, Components: 4
	// Gray Alpha: 1.0, Components: 2
}
