// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLAttributeDescriptor] class.
var (
	_MTLAttributeDescriptorClass     MTLAttributeDescriptorClass
	_MTLAttributeDescriptorClassOnce sync.Once
)

func getMTLAttributeDescriptorClass() MTLAttributeDescriptorClass {
	_MTLAttributeDescriptorClassOnce.Do(func() {
		_MTLAttributeDescriptorClass = MTLAttributeDescriptorClass{class: objc.GetClass("MTLAttributeDescriptor")}
	})
	return _MTLAttributeDescriptorClass
}

// GetMTLAttributeDescriptorClass returns the class object for MTLAttributeDescriptor.
func GetMTLAttributeDescriptorClass() MTLAttributeDescriptorClass {
	return getMTLAttributeDescriptorClass()
}

type MTLAttributeDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLAttributeDescriptorClass) Alloc() MTLAttributeDescriptor {
	rv := objc.Send[MTLAttributeDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A descriptor of an argument’s format and where its data is in memory.
//
// # Overview
// 
// Attribute descriptors are part of an [MTLVertexDescriptor] or
// [MTLStageInputOutputDescriptor] instance to provide layout information
// about a function’s arguments. Each descriptor is for a single argument,
// containing information about the attached data, offset and stride, and data
// type.
//
// # Defining attribute location
//
//   - [MTLAttributeDescriptor.BufferIndex]: The index in the buffer argument table for the buffer that contains the data for this attribute.
//   - [MTLAttributeDescriptor.SetBufferIndex]
//   - [MTLAttributeDescriptor.Offset]: The offset, in bytes, from the start of the buffer containing the attribute data to the start of the data itself.
//   - [MTLAttributeDescriptor.SetOffset]
//   - [MTLAttributeDescriptor.Format]: The format of the attribute’s data.
//   - [MTLAttributeDescriptor.SetFormat]
//
// See: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptor
type MTLAttributeDescriptor struct {
	objectivec.Object
}

// MTLAttributeDescriptorFromID constructs a [MTLAttributeDescriptor] from an objc.ID.
//
// A descriptor of an argument’s format and where its data is in memory.
func MTLAttributeDescriptorFromID(id objc.ID) MTLAttributeDescriptor {
	return MTLAttributeDescriptor{objectivec.Object{id}}
}
// NOTE: MTLAttributeDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLAttributeDescriptor] class.
//
// # Defining attribute location
//
//   - [IMTLAttributeDescriptor.BufferIndex]: The index in the buffer argument table for the buffer that contains the data for this attribute.
//   - [IMTLAttributeDescriptor.SetBufferIndex]
//   - [IMTLAttributeDescriptor.Offset]: The offset, in bytes, from the start of the buffer containing the attribute data to the start of the data itself.
//   - [IMTLAttributeDescriptor.SetOffset]
//   - [IMTLAttributeDescriptor.Format]: The format of the attribute’s data.
//   - [IMTLAttributeDescriptor.SetFormat]
//
// See: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptor
type IMTLAttributeDescriptor interface {
	objectivec.IObject

	// Topic: Defining attribute location

	// The index in the buffer argument table for the buffer that contains the data for this attribute.
	BufferIndex() uint
	SetBufferIndex(value uint)
	// The offset, in bytes, from the start of the buffer containing the attribute data to the start of the data itself.
	Offset() uint
	SetOffset(value uint)
	// The format of the attribute’s data.
	Format() MTLAttributeFormat
	SetFormat(value MTLAttributeFormat)

	// The organization of input and output data for the next kernel call.
	StageInputDescriptor() IMTLStageInputOutputDescriptor
	SetStageInputDescriptor(value IMTLStageInputOutputDescriptor)
}





// Init initializes the instance.
func (a MTLAttributeDescriptor) Init() MTLAttributeDescriptor {
	rv := objc.Send[MTLAttributeDescriptor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MTLAttributeDescriptor) Autorelease() MTLAttributeDescriptor {
	rv := objc.Send[MTLAttributeDescriptor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLAttributeDescriptor creates a new MTLAttributeDescriptor instance.
func NewMTLAttributeDescriptor() MTLAttributeDescriptor {
	class := getMTLAttributeDescriptorClass()
	rv := objc.Send[MTLAttributeDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// The index in the buffer argument table for the buffer that contains the
// data for this attribute.
//
// See: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptor/bufferIndex
func (a MTLAttributeDescriptor) BufferIndex() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("bufferIndex"))
	return rv
}
func (a MTLAttributeDescriptor) SetBufferIndex(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setBufferIndex:"), value)
}



// The offset, in bytes, from the start of the buffer containing the attribute
// data to the start of the data itself.
//
// See: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptor/offset
func (a MTLAttributeDescriptor) Offset() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("offset"))
	return rv
}
func (a MTLAttributeDescriptor) SetOffset(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setOffset:"), value)
}



// The format of the attribute’s data.
//
// See: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptor/format
func (a MTLAttributeDescriptor) Format() MTLAttributeFormat {
	rv := objc.Send[MTLAttributeFormat](a.ID, objc.Sel("format"))
	return MTLAttributeFormat(rv)
}
func (a MTLAttributeDescriptor) SetFormat(value MTLAttributeFormat) {
	objc.Send[struct{}](a.ID, objc.Sel("setFormat:"), value)
}



// The organization of input and output data for the next kernel call.
//
// See: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/stageinputdescriptor
func (a MTLAttributeDescriptor) StageInputDescriptor() IMTLStageInputOutputDescriptor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("stageInputDescriptor"))
	return MTLStageInputOutputDescriptorFromID(objc.ID(rv))
}
func (a MTLAttributeDescriptor) SetStageInputDescriptor(value IMTLStageInputOutputDescriptor) {
	objc.Send[struct{}](a.ID, objc.Sel("setStageInputDescriptor:"), value)
}
























