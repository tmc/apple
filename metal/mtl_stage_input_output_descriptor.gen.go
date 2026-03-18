// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLStageInputOutputDescriptor] class.
var (
	_MTLStageInputOutputDescriptorClass     MTLStageInputOutputDescriptorClass
	_MTLStageInputOutputDescriptorClassOnce sync.Once
)

func getMTLStageInputOutputDescriptorClass() MTLStageInputOutputDescriptorClass {
	_MTLStageInputOutputDescriptorClassOnce.Do(func() {
		_MTLStageInputOutputDescriptorClass = MTLStageInputOutputDescriptorClass{class: objc.GetClass("MTLStageInputOutputDescriptor")}
	})
	return _MTLStageInputOutputDescriptorClass
}

// GetMTLStageInputOutputDescriptorClass returns the class object for MTLStageInputOutputDescriptor.
func GetMTLStageInputOutputDescriptorClass() MTLStageInputOutputDescriptorClass {
	return getMTLStageInputOutputDescriptorClass()
}

type MTLStageInputOutputDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLStageInputOutputDescriptorClass) Alloc() MTLStageInputOutputDescriptor {
	rv := objc.Send[MTLStageInputOutputDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of the input and output data of a function.
//
// # Describing argument layouts
//
//   - [MTLStageInputOutputDescriptor.Attributes]: An array that describes where and how to fetch data for the function.
//   - [MTLStageInputOutputDescriptor.Layouts]: An array that describes how the function fetches data.
//
// # Declaring index buffers for indirect compute commands
//
//   - [MTLStageInputOutputDescriptor.IndexBufferIndex]: The location of the index buffer for a compute function using indexed thread addressing.
//   - [MTLStageInputOutputDescriptor.SetIndexBufferIndex]
//   - [MTLStageInputOutputDescriptor.IndexType]: The data type of the indices stored in the index buffer.
//   - [MTLStageInputOutputDescriptor.SetIndexType]
//
// # Resetting the descriptor
//
//   - [MTLStageInputOutputDescriptor.Reset]: Resets the default state for the descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor
type MTLStageInputOutputDescriptor struct {
	objectivec.Object
}

// MTLStageInputOutputDescriptorFromID constructs a [MTLStageInputOutputDescriptor] from an objc.ID.
//
// A description of the input and output data of a function.
func MTLStageInputOutputDescriptorFromID(id objc.ID) MTLStageInputOutputDescriptor {
	return MTLStageInputOutputDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLStageInputOutputDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLStageInputOutputDescriptor] class.
//
// # Describing argument layouts
//
//   - [IMTLStageInputOutputDescriptor.Attributes]: An array that describes where and how to fetch data for the function.
//   - [IMTLStageInputOutputDescriptor.Layouts]: An array that describes how the function fetches data.
//
// # Declaring index buffers for indirect compute commands
//
//   - [IMTLStageInputOutputDescriptor.IndexBufferIndex]: The location of the index buffer for a compute function using indexed thread addressing.
//   - [IMTLStageInputOutputDescriptor.SetIndexBufferIndex]
//   - [IMTLStageInputOutputDescriptor.IndexType]: The data type of the indices stored in the index buffer.
//   - [IMTLStageInputOutputDescriptor.SetIndexType]
//
// # Resetting the descriptor
//
//   - [IMTLStageInputOutputDescriptor.Reset]: Resets the default state for the descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor
type IMTLStageInputOutputDescriptor interface {
	objectivec.IObject

	// Topic: Describing argument layouts

	// An array that describes where and how to fetch data for the function.
	Attributes() IMTLAttributeDescriptorArray
	// An array that describes how the function fetches data.
	Layouts() IMTLBufferLayoutDescriptorArray

	// Topic: Declaring index buffers for indirect compute commands

	// The location of the index buffer for a compute function using indexed thread addressing.
	IndexBufferIndex() uint
	SetIndexBufferIndex(value uint)
	// The data type of the indices stored in the index buffer.
	IndexType() MTLIndexType
	SetIndexType(value MTLIndexType)

	// Topic: Resetting the descriptor

	// Resets the default state for the descriptor.
	Reset()
}

// Init initializes the instance.
func (s MTLStageInputOutputDescriptor) Init() MTLStageInputOutputDescriptor {
	rv := objc.Send[MTLStageInputOutputDescriptor](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MTLStageInputOutputDescriptor) Autorelease() MTLStageInputOutputDescriptor {
	rv := objc.Send[MTLStageInputOutputDescriptor](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLStageInputOutputDescriptor creates a new MTLStageInputOutputDescriptor instance.
func NewMTLStageInputOutputDescriptor() MTLStageInputOutputDescriptor {
	class := getMTLStageInputOutputDescriptorClass()
	rv := objc.Send[MTLStageInputOutputDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Resets the default state for the descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/reset()
func (s MTLStageInputOutputDescriptor) Reset() {
	objc.Send[objc.ID](s.ID, objc.Sel("reset"))
}

// See: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/stageInputOutputDescriptor
func (_MTLStageInputOutputDescriptorClass MTLStageInputOutputDescriptorClass) StageInputOutputDescriptor() MTLStageInputOutputDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MTLStageInputOutputDescriptorClass.class), objc.Sel("stageInputOutputDescriptor"))
	return MTLStageInputOutputDescriptorFromID(rv)
}

// An array that describes where and how to fetch data for the function.
//
// See: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/attributes
func (s MTLStageInputOutputDescriptor) Attributes() IMTLAttributeDescriptorArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("attributes"))
	return MTLAttributeDescriptorArrayFromID(objc.ID(rv))
}

// An array that describes how the function fetches data.
//
// See: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/layouts
func (s MTLStageInputOutputDescriptor) Layouts() IMTLBufferLayoutDescriptorArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("layouts"))
	return MTLBufferLayoutDescriptorArrayFromID(objc.ID(rv))
}

// The location of the index buffer for a compute function using indexed
// thread addressing.
//
// See: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/indexBufferIndex
func (s MTLStageInputOutputDescriptor) IndexBufferIndex() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("indexBufferIndex"))
	return rv
}
func (s MTLStageInputOutputDescriptor) SetIndexBufferIndex(value uint) {
	objc.Send[struct{}](s.ID, objc.Sel("setIndexBufferIndex:"), value)
}

// The data type of the indices stored in the index buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/indexType
func (s MTLStageInputOutputDescriptor) IndexType() MTLIndexType {
	rv := objc.Send[MTLIndexType](s.ID, objc.Sel("indexType"))
	return MTLIndexType(rv)
}
func (s MTLStageInputOutputDescriptor) SetIndexType(value MTLIndexType) {
	objc.Send[struct{}](s.ID, objc.Sel("setIndexType:"), value)
}

