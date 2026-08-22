//go:build darwin

package objcbridge_test

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objc/objcbridge"
)

func ExampleBlockInvoker() {
	var capturedUint uint

	// Create an Objective-C block from a Go function.
	block := objc.NewBlock(func(_ objc.Block, val uint) {
		capturedUint = val
	})
	defer block.Release()

	// Invoke the block via BlockInvoker.
	invoker := objcbridge.NewBlockInvoker()
	if err := invoker.Uint(objc.ID(block), 42); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Captured value:", capturedUint)

	// Output:
	// Captured value: 42
}

func ExampleProtocolsByName() {
	// Resolve Objective-C runtime protocols by name.
	protocols := objcbridge.ProtocolsByName("NSObject")
	fmt.Println("Resolved count:", len(protocols))

	// Output:
	// Resolved count: 1
}

func ExampleRequiredProtocolsByName() {
	// Look up required protocols, returning an error if any is missing.
	protocols, err := objcbridge.RequiredProtocolsByName("NSObject")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Found required protocols:", len(protocols))

	// Output:
	// Found required protocols: 1
}
