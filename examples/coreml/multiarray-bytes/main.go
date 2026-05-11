// Command multiarray-bytes creates an MLMultiArray and reads its backing bytes.
//
// It demonstrates MLMultiArray.GetBytesWithHandler, whose callback receives both
// the buffer pointer and the byte count from CoreML.
package main

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/tmc/apple/coreml"
	"github.com/tmc/apple/foundation"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	shape := []foundation.NSNumber{
		foundation.NewNumberWithInteger(2),
		foundation.NewNumberWithInteger(3),
	}
	array, err := coreml.NewMultiArrayWithShapeDataTypeError(shape, coreml.MLMultiArrayDataTypeFloat32)
	if err != nil {
		panic(err)
	}

	ptr := array.DataPointer()
	values := unsafe.Slice((*float32)(ptr), array.Count())
	for i := range values {
		values[i] = float32(i) + 0.5
	}

	array.GetBytesWithHandler(func(bytes unsafe.Pointer, size int64) {
		fmt.Printf("count=%d bytes=%d first=%.1f last=%.1f\n",
			array.Count(), size, values[0], values[len(values)-1])
		_ = bytes
	})
}
