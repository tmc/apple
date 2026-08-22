//go:build darwin

package applicationservices_test

import (
	"fmt"
	"unsafe"

	"github.com/tmc/apple/applicationservices"
	"github.com/tmc/apple/corefoundation"
)

func ExampleAXUIElementCreateSystemWide() {
	elem := applicationservices.AXUIElementCreateSystemWide()
	typeID := applicationservices.AXUIElementGetTypeID()
	matches := corefoundation.CFGetTypeID(corefoundation.CFTypeRef(elem)) == typeID
	fmt.Printf("System-wide element matches AXUIElement type ID: %t\n", matches)

	// Output:
	// System-wide element matches AXUIElement type ID: true
}

func ExampleHIShapeCreateWithRect() {
	rect := corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: 10, Y: 20},
		Size:   corefoundation.CGSize{Width: 100, Height: 50},
	}
	shape := applicationservices.HIShapeCreateWithRect(unsafe.Pointer(&rect))
	empty := applicationservices.HIShapeIsEmpty(shape)
	fmt.Printf("Shape is empty: %t\n", empty)

	// Output:
	// Shape is empty: false
}
