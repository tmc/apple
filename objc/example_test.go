//go:build darwin

package objc_test

import (
	"fmt"

	"github.com/tmc/apple/objc"
)

func Example() {
	// Convert a Go string to an Objective-C NSString object.
	strObj := objc.String("Hello, Apple Runtime")

	// Convert the NSString back to a Go string.
	str := objc.IDToString(strObj)
	fmt.Println(str)

	// Send selector 'length' to the NSString object.
	length := objc.Send[uint](strObj, objc.Sel("length"))
	fmt.Println("Length:", length)

	// Output:
	// Hello, Apple Runtime
	// Length: 20
}

func ExampleAutoreleasePool() {
	// Execute memory-intensive Objective-C allocations within an autorelease pool.
	objc.AutoreleasePool(func() {
		strObj := objc.String("Autoreleased String")
		fmt.Println(objc.IDToString(strObj))
	})

	// Output:
	// Autoreleased String
}

func ExampleSafeSend() {
	strObj := objc.String("Apple")
	lengthSel := objc.Sel("length")

	// Safely send a selector after verifying the receiver responds to it.
	length, err := objc.SafeSend[uint](strObj, lengthSel)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Length:", length)

	// Output:
	// Length: 5
}
