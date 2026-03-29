package coreml

import (
	"unsafe"

	"github.com/tmc/apple/objc"
)

// DataPointer returns a pointer to the contiguous data backing this multiarray.
//
// The pointer is valid for the lifetime of the MLMultiArray (and any parent
// prediction result that owns it). Callers must retain the MLMultiArray or its
// owner to keep the pointer valid.
//
// The memory layout depends on DataType, Shape, and Strides. For arrays
// created by CoreML prediction, the data is typically contiguous in row-major
// order, but callers should check Strides to confirm.
//
// This property is not in the generated bindings because the code generator
// does not yet emit read-only unsafe.Pointer properties.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/dataPointer
func (m MLMultiArray) DataPointer() unsafe.Pointer {
	return objc.Send[unsafe.Pointer](m.ID, objc.Sel("dataPointer"))
}
