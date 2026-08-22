//go:build darwin

package corefoundation_test

import (
	"fmt"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
)

func ExampleCFRange() {
	r := corefoundation.CFRange{Location: 5, Length: 20}
	fmt.Printf("Location: %d, Length: %d\n", r.Location, r.Length)

	// Output:
	// Location: 5, Length: 20
}

func ExampleCFArrayGetCount() {
	arr := corefoundation.CFArrayCreateMutable(0, 0, nil)
	fmt.Println(corefoundation.CFArrayGetCount(corefoundation.CFArrayRef(arr)))

	val := uintptr(42)
	corefoundation.CFArrayAppendValue(arr, unsafe.Pointer(val))
	fmt.Println(corefoundation.CFArrayGetCount(corefoundation.CFArrayRef(arr)))

	// Output:
	// 0
	// 1
}
