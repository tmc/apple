// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLBufferLayoutDescriptorArray] class.
var (
	_MTLBufferLayoutDescriptorArrayClass     MTLBufferLayoutDescriptorArrayClass
	_MTLBufferLayoutDescriptorArrayClassOnce sync.Once
)

func getMTLBufferLayoutDescriptorArrayClass() MTLBufferLayoutDescriptorArrayClass {
	_MTLBufferLayoutDescriptorArrayClassOnce.Do(func() {
		_MTLBufferLayoutDescriptorArrayClass = MTLBufferLayoutDescriptorArrayClass{class: objc.GetClass("MTLBufferLayoutDescriptorArray")}
	})
	return _MTLBufferLayoutDescriptorArrayClass
}

// GetMTLBufferLayoutDescriptorArrayClass returns the class object for MTLBufferLayoutDescriptorArray.
func GetMTLBufferLayoutDescriptorArrayClass() MTLBufferLayoutDescriptorArrayClass {
	return getMTLBufferLayoutDescriptorArrayClass()
}

type MTLBufferLayoutDescriptorArrayClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLBufferLayoutDescriptorArrayClass) Alloc() MTLBufferLayoutDescriptorArray {
	rv := objc.Send[MTLBufferLayoutDescriptorArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// An array of buffer layout descriptor objects.
//
// # Overview
// 
// An [MTLBufferLayoutDescriptorArray] defines the data layout and loading for
// compute data, using [MTLBufferLayoutDescriptor] instances.
//
// # Array accessors
//
//   - [MTLBufferLayoutDescriptorArray.ObjectAtIndexedSubscript]: Returns the state of the specified buffer layout.
//
// See: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptorArray
type MTLBufferLayoutDescriptorArray struct {
	objectivec.Object
}

// MTLBufferLayoutDescriptorArrayFromID constructs a [MTLBufferLayoutDescriptorArray] from an objc.ID.
//
// An array of buffer layout descriptor objects.
func MTLBufferLayoutDescriptorArrayFromID(id objc.ID) MTLBufferLayoutDescriptorArray {
	return MTLBufferLayoutDescriptorArray{objectivec.Object{id}}
}
// NOTE: MTLBufferLayoutDescriptorArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLBufferLayoutDescriptorArray] class.
//
// # Array accessors
//
//   - [IMTLBufferLayoutDescriptorArray.ObjectAtIndexedSubscript]: Returns the state of the specified buffer layout.
//
// See: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptorArray
type IMTLBufferLayoutDescriptorArray interface {
	objectivec.IObject

	// Topic: Array accessors

	// Returns the state of the specified buffer layout.
	ObjectAtIndexedSubscript(index uint) IMTLBufferLayoutDescriptor

	// The organization of input and output data for the next kernel call.
	StageInputDescriptor() IMTLStageInputOutputDescriptor
	SetStageInputDescriptor(value IMTLStageInputOutputDescriptor)
	// Sets the state of the specified buffer layout.
	SetObjectAtIndexedSubscript(bufferDesc IMTLBufferLayoutDescriptor, index uint)
}





// Init initializes the instance.
func (b MTLBufferLayoutDescriptorArray) Init() MTLBufferLayoutDescriptorArray {
	rv := objc.Send[MTLBufferLayoutDescriptorArray](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MTLBufferLayoutDescriptorArray) Autorelease() MTLBufferLayoutDescriptorArray {
	rv := objc.Send[MTLBufferLayoutDescriptorArray](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLBufferLayoutDescriptorArray creates a new MTLBufferLayoutDescriptorArray instance.
func NewMTLBufferLayoutDescriptorArray() MTLBufferLayoutDescriptorArray {
	class := getMTLBufferLayoutDescriptorArrayClass()
	rv := objc.Send[MTLBufferLayoutDescriptorArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Returns the state of the specified buffer layout.
//
// index: A specified index in the array of buffer layouts.
//
// # Return Value
// 
// The buffer layout descriptor for the buffer bound to the given attribute
// table index.
//
// See: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptorArray/subscript(_:)
func (b MTLBufferLayoutDescriptorArray) ObjectAtIndexedSubscript(index uint) IMTLBufferLayoutDescriptor {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("objectAtIndexedSubscript:"), index)
	return MTLBufferLayoutDescriptorFromID(rv)
}

// Sets the state of the specified buffer layout.
//
// bufferDesc: A descriptor that contains buffer layout state.
//
// index: An index in the array of buffer layouts.
//
// See: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptorArray/setObject:atIndexedSubscript:
func (b MTLBufferLayoutDescriptorArray) SetObjectAtIndexedSubscript(bufferDesc IMTLBufferLayoutDescriptor, index uint) {
	objc.Send[objc.ID](b.ID, objc.Sel("setObject:atIndexedSubscript:"), bufferDesc, index)
}












// The organization of input and output data for the next kernel call.
//
// See: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/stageinputdescriptor
func (b MTLBufferLayoutDescriptorArray) StageInputDescriptor() IMTLStageInputOutputDescriptor {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("stageInputDescriptor"))
	return MTLStageInputOutputDescriptorFromID(objc.ID(rv))
}
func (b MTLBufferLayoutDescriptorArray) SetStageInputDescriptor(value IMTLStageInputOutputDescriptor) {
	objc.Send[struct{}](b.ID, objc.Sel("setStageInputDescriptor:"), value)
}























