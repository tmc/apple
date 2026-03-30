// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLPipelineBufferDescriptorArray] class.
var (
	_MTLPipelineBufferDescriptorArrayClass     MTLPipelineBufferDescriptorArrayClass
	_MTLPipelineBufferDescriptorArrayClassOnce sync.Once
)

func getMTLPipelineBufferDescriptorArrayClass() MTLPipelineBufferDescriptorArrayClass {
	_MTLPipelineBufferDescriptorArrayClassOnce.Do(func() {
		_MTLPipelineBufferDescriptorArrayClass = MTLPipelineBufferDescriptorArrayClass{class: objc.GetClass("MTLPipelineBufferDescriptorArray")}
	})
	return _MTLPipelineBufferDescriptorArrayClass
}

// GetMTLPipelineBufferDescriptorArrayClass returns the class object for MTLPipelineBufferDescriptorArray.
func GetMTLPipelineBufferDescriptorArrayClass() MTLPipelineBufferDescriptorArrayClass {
	return getMTLPipelineBufferDescriptorArrayClass()
}

type MTLPipelineBufferDescriptorArrayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLPipelineBufferDescriptorArrayClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLPipelineBufferDescriptorArrayClass) Alloc() MTLPipelineBufferDescriptorArray {
	rv := objc.Send[MTLPipelineBufferDescriptorArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An array of pipeline buffer descriptors.
//
// # Accessing array elements
//
//   - [MTLPipelineBufferDescriptorArray.ObjectAtIndexedSubscript]: Returns the pipeline buffer descriptor at the specified array index.
//
// See: https://developer.apple.com/documentation/Metal/MTLPipelineBufferDescriptorArray
type MTLPipelineBufferDescriptorArray struct {
	objectivec.Object
}

// MTLPipelineBufferDescriptorArrayFromID constructs a [MTLPipelineBufferDescriptorArray] from an objc.ID.
//
// An array of pipeline buffer descriptors.
func MTLPipelineBufferDescriptorArrayFromID(id objc.ID) MTLPipelineBufferDescriptorArray {
	return MTLPipelineBufferDescriptorArray{objectivec.Object{ID: id}}
}

// NOTE: MTLPipelineBufferDescriptorArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLPipelineBufferDescriptorArray] class.
//
// # Accessing array elements
//
//   - [IMTLPipelineBufferDescriptorArray.ObjectAtIndexedSubscript]: Returns the pipeline buffer descriptor at the specified array index.
//
// See: https://developer.apple.com/documentation/Metal/MTLPipelineBufferDescriptorArray
type IMTLPipelineBufferDescriptorArray interface {
	objectivec.IObject

	// Topic: Accessing array elements

	// Returns the pipeline buffer descriptor at the specified array index.
	ObjectAtIndexedSubscript(bufferIndex uint) IMTLPipelineBufferDescriptor

	// Sets a pipeline buffer descriptor at the specified array index.
	SetObjectAtIndexedSubscript(buffer IMTLPipelineBufferDescriptor, bufferIndex uint)
}

// Init initializes the instance.
func (p MTLPipelineBufferDescriptorArray) Init() MTLPipelineBufferDescriptorArray {
	rv := objc.Send[MTLPipelineBufferDescriptorArray](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MTLPipelineBufferDescriptorArray) Autorelease() MTLPipelineBufferDescriptorArray {
	rv := objc.Send[MTLPipelineBufferDescriptorArray](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLPipelineBufferDescriptorArray creates a new MTLPipelineBufferDescriptorArray instance.
func NewMTLPipelineBufferDescriptorArray() MTLPipelineBufferDescriptorArray {
	class := getMTLPipelineBufferDescriptorArrayClass()
	rv := objc.Send[MTLPipelineBufferDescriptorArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the pipeline buffer descriptor at the specified array index.
//
// bufferIndex: The array index of the requested pipeline buffer descriptor.
//
// # Return Value
//
// The descriptor for the buffer bound at this index.
//
// See: https://developer.apple.com/documentation/Metal/MTLPipelineBufferDescriptorArray/subscript(_:)
func (p MTLPipelineBufferDescriptorArray) ObjectAtIndexedSubscript(bufferIndex uint) IMTLPipelineBufferDescriptor {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("objectAtIndexedSubscript:"), bufferIndex)
	return MTLPipelineBufferDescriptorFromID(rv)
}

// Sets a pipeline buffer descriptor at the specified array index.
//
// buffer: The pipeline buffer descriptor to set in the array.
//
// bufferIndex: The array index in which to set the given pipeline buffer descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTLPipelineBufferDescriptorArray/setObject:atIndexedSubscript:
func (p MTLPipelineBufferDescriptorArray) SetObjectAtIndexedSubscript(buffer IMTLPipelineBufferDescriptor, bufferIndex uint) {
	objc.Send[objc.ID](p.ID, objc.Sel("setObject:atIndexedSubscript:"), buffer, bufferIndex)
}
