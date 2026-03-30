// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLVertexAttributeDescriptor] class.
var (
	_MTLVertexAttributeDescriptorClass     MTLVertexAttributeDescriptorClass
	_MTLVertexAttributeDescriptorClassOnce sync.Once
)

func getMTLVertexAttributeDescriptorClass() MTLVertexAttributeDescriptorClass {
	_MTLVertexAttributeDescriptorClassOnce.Do(func() {
		_MTLVertexAttributeDescriptorClass = MTLVertexAttributeDescriptorClass{class: objc.GetClass("MTLVertexAttributeDescriptor")}
	})
	return _MTLVertexAttributeDescriptorClass
}

// GetMTLVertexAttributeDescriptorClass returns the class object for MTLVertexAttributeDescriptor.
func GetMTLVertexAttributeDescriptorClass() MTLVertexAttributeDescriptorClass {
	return getMTLVertexAttributeDescriptorClass()
}

type MTLVertexAttributeDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLVertexAttributeDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLVertexAttributeDescriptorClass) Alloc() MTLVertexAttributeDescriptor {
	rv := objc.Send[MTLVertexAttributeDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that determines how to store attribute data in memory and map it
// to the arguments of a vertex function.
//
// # Overview
//
// A vertex attribute descriptor provides organization information so a vertex
// shader function can locate and load data into its arguments. The descriptor
// maps memory locations to attribute locations. It supports access to
// multiple attributes (such as vertex coordinates, surface normals, and
// texture coordinates) that are interleaved within the same buffer.
//
// # Organizing the vertex attribute
//
//   - [MTLVertexAttributeDescriptor.Format]: The format of the vertex attribute.
//   - [MTLVertexAttributeDescriptor.SetFormat]
//   - [MTLVertexAttributeDescriptor.Offset]: The location of an attribute in vertex data, determined by the byte offset from the start of the vertex data.
//   - [MTLVertexAttributeDescriptor.SetOffset]
//   - [MTLVertexAttributeDescriptor.BufferIndex]: The index in the argument table for the associated vertex buffer.
//   - [MTLVertexAttributeDescriptor.SetBufferIndex]
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptor
type MTLVertexAttributeDescriptor struct {
	objectivec.Object
}

// MTLVertexAttributeDescriptorFromID constructs a [MTLVertexAttributeDescriptor] from an objc.ID.
//
// An object that determines how to store attribute data in memory and map it
// to the arguments of a vertex function.
func MTLVertexAttributeDescriptorFromID(id objc.ID) MTLVertexAttributeDescriptor {
	return MTLVertexAttributeDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MTLVertexAttributeDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLVertexAttributeDescriptor] class.
//
// # Organizing the vertex attribute
//
//   - [IMTLVertexAttributeDescriptor.Format]: The format of the vertex attribute.
//   - [IMTLVertexAttributeDescriptor.SetFormat]
//   - [IMTLVertexAttributeDescriptor.Offset]: The location of an attribute in vertex data, determined by the byte offset from the start of the vertex data.
//   - [IMTLVertexAttributeDescriptor.SetOffset]
//   - [IMTLVertexAttributeDescriptor.BufferIndex]: The index in the argument table for the associated vertex buffer.
//   - [IMTLVertexAttributeDescriptor.SetBufferIndex]
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptor
type IMTLVertexAttributeDescriptor interface {
	objectivec.IObject

	// Topic: Organizing the vertex attribute

	// The format of the vertex attribute.
	Format() MTLVertexFormat
	SetFormat(value MTLVertexFormat)
	// The location of an attribute in vertex data, determined by the byte offset from the start of the vertex data.
	Offset() uint
	SetOffset(value uint)
	// The index in the argument table for the associated vertex buffer.
	BufferIndex() uint
	SetBufferIndex(value uint)

	MTLBufferLayoutStrideDynamic() int
}

// Init initializes the instance.
func (v MTLVertexAttributeDescriptor) Init() MTLVertexAttributeDescriptor {
	rv := objc.Send[MTLVertexAttributeDescriptor](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v MTLVertexAttributeDescriptor) Autorelease() MTLVertexAttributeDescriptor {
	rv := objc.Send[MTLVertexAttributeDescriptor](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLVertexAttributeDescriptor creates a new MTLVertexAttributeDescriptor instance.
func NewMTLVertexAttributeDescriptor() MTLVertexAttributeDescriptor {
	class := getMTLVertexAttributeDescriptorClass()
	rv := objc.Send[MTLVertexAttributeDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The format of the vertex attribute.
//
// # Discussion
//
// This property specifies the data type of the vertex attribute that
// corresponds to an input argument of a shading language function. The
// [MTLVertexFormat] may be converted to the data type in the shading function
// argument with the following specified limitations. Invalid type conversion
// causes a compilation error.
//
// Conversion of vectors of different lengths is valid. The length of vectors
// can be reduced. For example, [MTLVertexFormatInt4] data can be reduced to a
// single `int` shader argument is valid, and the last three values of the
// vector are discarded. Vectors can also be expanded; for example, expanding
// [MTLVertexFormatInt] to an `int4` vector shader argument is valid. When
// expanding, the extra components are filled with the corresponding
// components of (0,0,0,1).
//
// The sign of an integer [MTLVertexFormat] can not be cast to a shader
// argument with an integer type of a different sign. For example, casting the
// signed format [MTLVertexFormatInt] to an `uint` shader argument is invalid.
// Casting [MTLVertexFormatUInt] to an `int` argument is also invalid.
//
// Integer truncation is not supported. For example, casting the
// [MTLVertexFormatInt] to a `short` is invalid. However, casting
// [MTLVertexFormatShort2] to a vector of `int` values is valid.
//
// Casting any [MTLVertexFormat] to a `float` or `half` is valid. Casting
// normalized [MTLVertexFormat] types (such as
// [MTLVertexFormatShort2Normalized]) are only valid to `float` or `half`.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptor/format
//
// [MTLVertexFormat]: https://developer.apple.com/documentation/Metal/MTLVertexFormat
func (v MTLVertexAttributeDescriptor) Format() MTLVertexFormat {
	rv := objc.Send[MTLVertexFormat](v.ID, objc.Sel("format"))
	return MTLVertexFormat(rv)
}
func (v MTLVertexAttributeDescriptor) SetFormat(value MTLVertexFormat) {
	objc.Send[struct{}](v.ID, objc.Sel("setFormat:"), value)
}

// The location of an attribute in vertex data, determined by the byte offset
// from the start of the vertex data.
//
// # Discussion
//
// The `offset` value needs to be a multiple of `4` bytes.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptor/offset
func (v MTLVertexAttributeDescriptor) Offset() uint {
	rv := objc.Send[uint](v.ID, objc.Sel("offset"))
	return rv
}
func (v MTLVertexAttributeDescriptor) SetOffset(value uint) {
	objc.Send[struct{}](v.ID, objc.Sel("setOffset:"), value)
}

// The index in the argument table for the associated vertex buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptor/bufferIndex
func (v MTLVertexAttributeDescriptor) BufferIndex() uint {
	rv := objc.Send[uint](v.ID, objc.Sel("bufferIndex"))
	return rv
}
func (v MTLVertexAttributeDescriptor) SetBufferIndex(value uint) {
	objc.Send[struct{}](v.ID, objc.Sel("setBufferIndex:"), value)
}

// See: https://developer.apple.com/documentation/metal/mtlbufferlayoutstridedynamic
func (v MTLVertexAttributeDescriptor) MTLBufferLayoutStrideDynamic() int {
	rv := objc.Send[int](v.ID, objc.Sel("MTLBufferLayoutStrideDynamic"))
	return rv
}
