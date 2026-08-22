//go:build darwin

package metal_test

import (
	"fmt"

	"github.com/tmc/apple/metal"
)

func ExampleMTLCreateSystemDefaultDevice() {
	dev := metal.MTLCreateSystemDefaultDevice()
	if dev.ID != 0 {
		buf := dev.NewBufferWithLengthOptions(1024, metal.MTLResourceStorageModeShared)
		fmt.Printf("Buffer length: %d\n", buf.Length())
		fmt.Printf("Has device name: %t\n", len(dev.Name()) > 0)
	}

	// Output:
	// Buffer length: 1024
	// Has device name: true
}

func ExampleNewMTLCompileOptions() {
	opts := metal.NewMTLCompileOptions()
	opts.SetFastMathEnabled(true)
	opts.SetLanguageVersion(metal.MTLLanguageVersion3_0)

	fmt.Printf("FastMathEnabled: %t\n", opts.FastMathEnabled())
	fmt.Printf("LanguageVersion: %s\n", opts.LanguageVersion())

	// Output:
	// FastMathEnabled: true
	// LanguageVersion: MTLLanguageVersion3_0
}
