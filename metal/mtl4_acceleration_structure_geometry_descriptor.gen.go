// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4AccelerationStructureGeometryDescriptor] class.
var (
	_MTL4AccelerationStructureGeometryDescriptorClass     MTL4AccelerationStructureGeometryDescriptorClass
	_MTL4AccelerationStructureGeometryDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureGeometryDescriptorClass() MTL4AccelerationStructureGeometryDescriptorClass {
	_MTL4AccelerationStructureGeometryDescriptorClassOnce.Do(func() {
		_MTL4AccelerationStructureGeometryDescriptorClass = MTL4AccelerationStructureGeometryDescriptorClass{class: objc.GetClass("MTL4AccelerationStructureGeometryDescriptor")}
	})
	return _MTL4AccelerationStructureGeometryDescriptorClass
}

// GetMTL4AccelerationStructureGeometryDescriptorClass returns the class object for MTL4AccelerationStructureGeometryDescriptor.
func GetMTL4AccelerationStructureGeometryDescriptorClass() MTL4AccelerationStructureGeometryDescriptorClass {
	return getMTL4AccelerationStructureGeometryDescriptorClass()
}

type MTL4AccelerationStructureGeometryDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTL4AccelerationStructureGeometryDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4AccelerationStructureGeometryDescriptorClass) Alloc() MTL4AccelerationStructureGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Base class for all Metal 4 acceleration structure geometry descriptors.
//
// # Overview
//
// Don’t use this class directly. Use one of the derived classes instead.
//
// # Instance Properties
//
//   - [MTL4AccelerationStructureGeometryDescriptor.AllowDuplicateIntersectionFunctionInvocation]: A boolean value that indicates whether the ray-tracing system in Metal allows the invocation of intersection functions more than once per ray-primitive intersection.
//   - [MTL4AccelerationStructureGeometryDescriptor.SetAllowDuplicateIntersectionFunctionInvocation]
//   - [MTL4AccelerationStructureGeometryDescriptor.IntersectionFunctionTableOffset]: Sets the offset that this geometry contributes to determining the intersection function to invoke when a ray intersects it.
//   - [MTL4AccelerationStructureGeometryDescriptor.SetIntersectionFunctionTableOffset]
//   - [MTL4AccelerationStructureGeometryDescriptor.Label]: Assigns an optional label you can assign to this geometry for debugging purposes.
//   - [MTL4AccelerationStructureGeometryDescriptor.SetLabel]
//   - [MTL4AccelerationStructureGeometryDescriptor.Opaque]: Provides a hint to Metal that this geometry is opaque, potentially accelerating the ray/primitive intersection process.
//   - [MTL4AccelerationStructureGeometryDescriptor.SetOpaque]
//   - [MTL4AccelerationStructureGeometryDescriptor.PrimitiveDataBuffer]: Assigns optional buffer containing data to associate with each primitive in this geometry.
//   - [MTL4AccelerationStructureGeometryDescriptor.SetPrimitiveDataBuffer]
//   - [MTL4AccelerationStructureGeometryDescriptor.PrimitiveDataElementSize]: Sets the size, in bytes, of the data for each primitive in the primitive data buffer [primitiveDataBuffer](<doc://com.apple.metal/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/primitiveDataBuffer>) references.
//   - [MTL4AccelerationStructureGeometryDescriptor.SetPrimitiveDataElementSize]
//   - [MTL4AccelerationStructureGeometryDescriptor.PrimitiveDataStride]: Defines the stride, in bytes, between each primitive’s data in the primitive data buffer [primitiveDataBuffer](<doc://com.apple.metal/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/primitiveDataBuffer>) references.
//   - [MTL4AccelerationStructureGeometryDescriptor.SetPrimitiveDataStride]
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor
type MTL4AccelerationStructureGeometryDescriptor struct {
	objectivec.Object
}

// MTL4AccelerationStructureGeometryDescriptorFromID constructs a [MTL4AccelerationStructureGeometryDescriptor] from an objc.ID.
//
// Base class for all Metal 4 acceleration structure geometry descriptors.
func MTL4AccelerationStructureGeometryDescriptorFromID(id objc.ID) MTL4AccelerationStructureGeometryDescriptor {
	return MTL4AccelerationStructureGeometryDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MTL4AccelerationStructureGeometryDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4AccelerationStructureGeometryDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4AccelerationStructureGeometryDescriptor.AllowDuplicateIntersectionFunctionInvocation]: A boolean value that indicates whether the ray-tracing system in Metal allows the invocation of intersection functions more than once per ray-primitive intersection.
//   - [IMTL4AccelerationStructureGeometryDescriptor.SetAllowDuplicateIntersectionFunctionInvocation]
//   - [IMTL4AccelerationStructureGeometryDescriptor.IntersectionFunctionTableOffset]: Sets the offset that this geometry contributes to determining the intersection function to invoke when a ray intersects it.
//   - [IMTL4AccelerationStructureGeometryDescriptor.SetIntersectionFunctionTableOffset]
//   - [IMTL4AccelerationStructureGeometryDescriptor.Label]: Assigns an optional label you can assign to this geometry for debugging purposes.
//   - [IMTL4AccelerationStructureGeometryDescriptor.SetLabel]
//   - [IMTL4AccelerationStructureGeometryDescriptor.Opaque]: Provides a hint to Metal that this geometry is opaque, potentially accelerating the ray/primitive intersection process.
//   - [IMTL4AccelerationStructureGeometryDescriptor.SetOpaque]
//   - [IMTL4AccelerationStructureGeometryDescriptor.PrimitiveDataBuffer]: Assigns optional buffer containing data to associate with each primitive in this geometry.
//   - [IMTL4AccelerationStructureGeometryDescriptor.SetPrimitiveDataBuffer]
//   - [IMTL4AccelerationStructureGeometryDescriptor.PrimitiveDataElementSize]: Sets the size, in bytes, of the data for each primitive in the primitive data buffer [primitiveDataBuffer](<doc://com.apple.metal/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/primitiveDataBuffer>) references.
//   - [IMTL4AccelerationStructureGeometryDescriptor.SetPrimitiveDataElementSize]
//   - [IMTL4AccelerationStructureGeometryDescriptor.PrimitiveDataStride]: Defines the stride, in bytes, between each primitive’s data in the primitive data buffer [primitiveDataBuffer](<doc://com.apple.metal/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/primitiveDataBuffer>) references.
//   - [IMTL4AccelerationStructureGeometryDescriptor.SetPrimitiveDataStride]
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor
type IMTL4AccelerationStructureGeometryDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	// A boolean value that indicates whether the ray-tracing system in Metal allows the invocation of intersection functions more than once per ray-primitive intersection.
	AllowDuplicateIntersectionFunctionInvocation() bool
	SetAllowDuplicateIntersectionFunctionInvocation(value bool)
	// Sets the offset that this geometry contributes to determining the intersection function to invoke when a ray intersects it.
	IntersectionFunctionTableOffset() uint
	SetIntersectionFunctionTableOffset(value uint)
	// Assigns an optional label you can assign to this geometry for debugging purposes.
	Label() string
	SetLabel(value string)
	// Provides a hint to Metal that this geometry is opaque, potentially accelerating the ray/primitive intersection process.
	Opaque() bool
	SetOpaque(value bool)
	// Assigns optional buffer containing data to associate with each primitive in this geometry.
	PrimitiveDataBuffer() MTL4BufferRange
	SetPrimitiveDataBuffer(value MTL4BufferRange)
	// Sets the size, in bytes, of the data for each primitive in the primitive data buffer [primitiveDataBuffer](<doc://com.apple.metal/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/primitiveDataBuffer>) references.
	PrimitiveDataElementSize() uint
	SetPrimitiveDataElementSize(value uint)
	// Defines the stride, in bytes, between each primitive’s data in the primitive data buffer [primitiveDataBuffer](<doc://com.apple.metal/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/primitiveDataBuffer>) references.
	PrimitiveDataStride() uint
	SetPrimitiveDataStride(value uint)
}

// Init initializes the instance.
func (m MTL4AccelerationStructureGeometryDescriptor) Init() MTL4AccelerationStructureGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureGeometryDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4AccelerationStructureGeometryDescriptor) Autorelease() MTL4AccelerationStructureGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureGeometryDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureGeometryDescriptor creates a new MTL4AccelerationStructureGeometryDescriptor instance.
func NewMTL4AccelerationStructureGeometryDescriptor() MTL4AccelerationStructureGeometryDescriptor {
	class := getMTL4AccelerationStructureGeometryDescriptorClass()
	rv := objc.Send[MTL4AccelerationStructureGeometryDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A boolean value that indicates whether the ray-tracing system in Metal
// allows the invocation of intersection functions more than once per
// ray-primitive intersection.
//
// # Discussion
//
// The property’s default value is true.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/allowDuplicateIntersectionFunctionInvocation
func (m MTL4AccelerationStructureGeometryDescriptor) AllowDuplicateIntersectionFunctionInvocation() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("allowDuplicateIntersectionFunctionInvocation"))
	return rv
}
func (m MTL4AccelerationStructureGeometryDescriptor) SetAllowDuplicateIntersectionFunctionInvocation(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setAllowDuplicateIntersectionFunctionInvocation:"), value)
}

// Sets the offset that this geometry contributes to determining the
// intersection function to invoke when a ray intersects it.
//
// # Discussion
//
// When you perform a ray tracing operation in the Metal Shading Language, and
// provide the ray intersector object with an instance of
// [MTLIntersectionFunctionTable], Metal adds this offset to the instance
// offset from structs such as:
//
// - [MTLAccelerationStructureInstanceDescriptor] -
// [MTLAccelerationStructureUserIDInstanceDescriptor] -
// [MTLAccelerationStructureMotionInstanceDescriptor] -
// [MTLIndirectAccelerationStructureInstanceDescriptor] -
// [MTLIndirectAccelerationStructureMotionInstanceDescriptor]
//
// The sum of these offsets provides an index into the intersection function
// table that the ray tracing system uses to retrieve and invoke the function
// at this index, allowing you to customize the intersection evaluation
// process.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/intersectionFunctionTableOffset
//
// [MTLAccelerationStructureInstanceDescriptor]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureInstanceDescriptor
// [MTLAccelerationStructureMotionInstanceDescriptor]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionInstanceDescriptor
// [MTLAccelerationStructureUserIDInstanceDescriptor]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureUserIDInstanceDescriptor
// [MTLIndirectAccelerationStructureInstanceDescriptor]: https://developer.apple.com/documentation/Metal/MTLIndirectAccelerationStructureInstanceDescriptor
// [MTLIndirectAccelerationStructureMotionInstanceDescriptor]: https://developer.apple.com/documentation/Metal/MTLIndirectAccelerationStructureMotionInstanceDescriptor
func (m MTL4AccelerationStructureGeometryDescriptor) IntersectionFunctionTableOffset() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("intersectionFunctionTableOffset"))
	return rv
}
func (m MTL4AccelerationStructureGeometryDescriptor) SetIntersectionFunctionTableOffset(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setIntersectionFunctionTableOffset:"), value)
}

// Assigns an optional label you can assign to this geometry for debugging
// purposes.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/label
func (m MTL4AccelerationStructureGeometryDescriptor) Label() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (m MTL4AccelerationStructureGeometryDescriptor) SetLabel(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setLabel:"), objc.String(value))
}

// Provides a hint to Metal that this geometry is opaque, potentially
// accelerating the ray/primitive intersection process.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/opaque
func (m MTL4AccelerationStructureGeometryDescriptor) Opaque() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("opaque"))
	return rv
}
func (m MTL4AccelerationStructureGeometryDescriptor) SetOpaque(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setOpaque:"), value)
}

// Assigns optional buffer containing data to associate with each primitive in
// this geometry.
//
// # Discussion
//
// You can use zero as the buffer address in this buffer range.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/primitiveDataBuffer
func (m MTL4AccelerationStructureGeometryDescriptor) PrimitiveDataBuffer() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("primitiveDataBuffer"))
	return MTL4BufferRange(rv)
}
func (m MTL4AccelerationStructureGeometryDescriptor) SetPrimitiveDataBuffer(value MTL4BufferRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setPrimitiveDataBuffer:"), value)
}

// Sets the size, in bytes, of the data for each primitive in the primitive
// data buffer [PrimitiveDataBuffer] references.
//
// # Discussion
//
// This size needs to be at most [PrimitiveDataStride] in size and a multiple
// of 4 bytes.
//
// This property defaults to 0 bytes.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/primitiveDataElementSize
func (m MTL4AccelerationStructureGeometryDescriptor) PrimitiveDataElementSize() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("primitiveDataElementSize"))
	return rv
}
func (m MTL4AccelerationStructureGeometryDescriptor) SetPrimitiveDataElementSize(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setPrimitiveDataElementSize:"), value)
}

// Defines the stride, in bytes, between each primitive’s data in the
// primitive data buffer [PrimitiveDataBuffer] references.
//
// # Discussion
//
// You are responsible for ensuring the stride is at least
// [PrimitiveDataElementSize] in size and a multiple of 4 bytes.
//
// This property defaults to `0` bytes, which indicates the stride is equal to
// [PrimitiveDataElementSize].
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/primitiveDataStride
func (m MTL4AccelerationStructureGeometryDescriptor) PrimitiveDataStride() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("primitiveDataStride"))
	return rv
}
func (m MTL4AccelerationStructureGeometryDescriptor) SetPrimitiveDataStride(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setPrimitiveDataStride:"), value)
}
