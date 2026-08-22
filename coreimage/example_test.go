//go:build darwin

package coreimage_test

import (
	"fmt"

	"github.com/tmc/apple/coreimage"
)

func ExampleCIColor() {
	c := coreimage.NewColorWithRedGreenBlueAlpha(1.0, 0.5, 0.25, 1.0)
	fmt.Printf("R: %.1f, G: %.1f, B: %.2f, A: %.1f\n", c.Red(), c.Green(), c.Blue(), c.Alpha())
	fmt.Println("Components:", c.NumberOfComponents())

	// Output:
	// R: 1.0, G: 0.5, B: 0.25, A: 1.0
	// Components: 4
}

func ExampleCIVector() {
	v := coreimage.NewVectorWithXY(10.5, 20.25)
	fmt.Println("Count:", v.Count())
	fmt.Printf("X: %.1f, Y: %.2f\n", v.X(), v.Y())

	// Output:
	// Count: 2
	// X: 10.5, Y: 20.25
}
