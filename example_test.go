//go:build darwin

package apple_test

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

func Example() {
	// Demonstrates using the github.com/tmc/apple framework packages
	// to interact with Objective-C objects without cgo.
	nsStr := objc.String("Apple Go Bindings")
	obj := objectivec.ObjectFromID(nsStr)

	fmt.Printf("Description: %s\n", obj.Description())

	// Output:
	// Description: Apple Go Bindings
}
