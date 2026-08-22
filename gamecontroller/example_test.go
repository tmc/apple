//go:build darwin

package gamecontroller_test

import (
	"fmt"

	"github.com/tmc/apple/gamecontroller"
)

func ExampleGCColor() {
	color := gamecontroller.NewGCColorWithRedGreenBlue(1.0, 0.5, 0.0)

	fmt.Printf("RGB: %.1f, %.1f, %.1f\n", color.Red(), color.Green(), color.Blue())

	// Output:
	// RGB: 1.0, 0.5, 0.0
}

func ExampleGCController() {
	controllers := gamecontroller.GetGCControllerClass().Controllers()

	fmt.Println("Controllers connected:", len(controllers))

	// Output:
	// Controllers connected: 0
}

func ExampleGCControllerPlayerIndex() {
	idx1 := gamecontroller.GCControllerPlayerIndex1
	idxUnset := gamecontroller.GCControllerPlayerIndexUnset

	fmt.Println("Player 1:", idx1)
	fmt.Println("Unset:", idxUnset)

	// Output:
	// Player 1: GCControllerPlayerIndex1
	// Unset: GCControllerPlayerIndexUnset
}
