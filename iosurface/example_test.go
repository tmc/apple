//go:build darwin

package iosurface_test

import (
	"fmt"

	"github.com/tmc/apple/iosurface"
)

func ExampleIOSurfaceGetTypeID() {
	typeID := iosurface.IOSurfaceGetTypeID()
	fmt.Printf("IOSurface type ID matches: %t\n", typeID == iosurface.IOSurfaceGetTypeID())

	// Output:
	// IOSurface type ID matches: true
}

func ExampleGetIOSurfaceClass() {
	class := iosurface.GetIOSurfaceClass()
	fmt.Printf("IOSurface class consistent: %t\n", class.Class() == iosurface.GetIOSurfaceClass().Class())

	// Output:
	// IOSurface class consistent: true
}
