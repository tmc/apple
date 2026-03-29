// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLAccelerationStructureGeometryDescriptor] class.
var (
	_MTLAccelerationStructureGeometryDescriptorClass     MTLAccelerationStructureGeometryDescriptorClass
	_MTLAccelerationStructureGeometryDescriptorClassOnce sync.Once
)

func getMTLAccelerationStructureGeometryDescriptorClass() MTLAccelerationStructureGeometryDescriptorClass {
	_MTLAccelerationStructureGeometryDescriptorClassOnce.Do(func() {
		_MTLAccelerationStructureGeometryDescriptorClass = MTLAccelerationStructureGeometryDescriptorClass{class: objc.GetClass("MTLAccelerationStructureGeometryDescriptor")}
	})
	return _MTLAccelerationStructureGeometryDescriptorClass
}

// GetMTLAccelerationStructureGeometryDescriptorClass returns the class object for MTLAccelerationStructureGeometryDescriptor.
func GetMTLAccelerationStructureGeometryDescriptorClass() MTLAccelerationStructureGeometryDescriptorClass {
	return getMTLAccelerationStructureGeometryDescriptorClass()
}

type MTLAccelerationStructureGeometryDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLAccelerationStructureGeometryDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLAccelerationStructureGeometryDescriptorClass) Alloc() MTLAccelerationStructureGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A base class for descriptors that contain geometry data to convert into a
// ray-tracing acceleration structure.
//
// # Overview
// 
// Don’t use this base class directly. Use one of the derived classes
// instead, as [MTLAccelerationStructure] describes.
//
// # Specifying base geometry properties
//
//   - [MTLAccelerationStructureGeometryDescriptor.Label]: A label for the geometry structure, suitable for debugging.
//   - [MTLAccelerationStructureGeometryDescriptor.SetLabel]
//   - [MTLAccelerationStructureGeometryDescriptor.IntersectionFunctionTableOffset]: An index into the intersection table for determining which intersection function Metal calls when it intersects a ray with the acceleration structure.
//   - [MTLAccelerationStructureGeometryDescriptor.SetIntersectionFunctionTableOffset]
//   - [MTLAccelerationStructureGeometryDescriptor.Opaque]: A Boolean value that determines whether the geometry data in the acceleration structure needs to skip triangle-intersection tests.
//   - [MTLAccelerationStructureGeometryDescriptor.SetOpaque]
//   - [MTLAccelerationStructureGeometryDescriptor.AllowDuplicateIntersectionFunctionInvocation]: A Boolean value that indicates whether Metal calls the ray-intersection test more than once per primitive on the structure.
//   - [MTLAccelerationStructureGeometryDescriptor.SetAllowDuplicateIntersectionFunctionInvocation]
//
// # Instance Properties
//
//   - [MTLAccelerationStructureGeometryDescriptor.PrimitiveDataBuffer]
//   - [MTLAccelerationStructureGeometryDescriptor.SetPrimitiveDataBuffer]
//   - [MTLAccelerationStructureGeometryDescriptor.PrimitiveDataBufferOffset]
//   - [MTLAccelerationStructureGeometryDescriptor.SetPrimitiveDataBufferOffset]
//   - [MTLAccelerationStructureGeometryDescriptor.PrimitiveDataElementSize]
//   - [MTLAccelerationStructureGeometryDescriptor.SetPrimitiveDataElementSize]
//   - [MTLAccelerationStructureGeometryDescriptor.PrimitiveDataStride]
//   - [MTLAccelerationStructureGeometryDescriptor.SetPrimitiveDataStride]
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor
type MTLAccelerationStructureGeometryDescriptor struct {
	objectivec.Object
}

// MTLAccelerationStructureGeometryDescriptorFromID constructs a [MTLAccelerationStructureGeometryDescriptor] from an objc.ID.
//
// A base class for descriptors that contain geometry data to convert into a
// ray-tracing acceleration structure.
func MTLAccelerationStructureGeometryDescriptorFromID(id objc.ID) MTLAccelerationStructureGeometryDescriptor {
	return MTLAccelerationStructureGeometryDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLAccelerationStructureGeometryDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLAccelerationStructureGeometryDescriptor] class.
//
// # Specifying base geometry properties
//
//   - [IMTLAccelerationStructureGeometryDescriptor.Label]: A label for the geometry structure, suitable for debugging.
//   - [IMTLAccelerationStructureGeometryDescriptor.SetLabel]
//   - [IMTLAccelerationStructureGeometryDescriptor.IntersectionFunctionTableOffset]: An index into the intersection table for determining which intersection function Metal calls when it intersects a ray with the acceleration structure.
//   - [IMTLAccelerationStructureGeometryDescriptor.SetIntersectionFunctionTableOffset]
//   - [IMTLAccelerationStructureGeometryDescriptor.Opaque]: A Boolean value that determines whether the geometry data in the acceleration structure needs to skip triangle-intersection tests.
//   - [IMTLAccelerationStructureGeometryDescriptor.SetOpaque]
//   - [IMTLAccelerationStructureGeometryDescriptor.AllowDuplicateIntersectionFunctionInvocation]: A Boolean value that indicates whether Metal calls the ray-intersection test more than once per primitive on the structure.
//   - [IMTLAccelerationStructureGeometryDescriptor.SetAllowDuplicateIntersectionFunctionInvocation]
//
// # Instance Properties
//
//   - [IMTLAccelerationStructureGeometryDescriptor.PrimitiveDataBuffer]
//   - [IMTLAccelerationStructureGeometryDescriptor.SetPrimitiveDataBuffer]
//   - [IMTLAccelerationStructureGeometryDescriptor.PrimitiveDataBufferOffset]
//   - [IMTLAccelerationStructureGeometryDescriptor.SetPrimitiveDataBufferOffset]
//   - [IMTLAccelerationStructureGeometryDescriptor.PrimitiveDataElementSize]
//   - [IMTLAccelerationStructureGeometryDescriptor.SetPrimitiveDataElementSize]
//   - [IMTLAccelerationStructureGeometryDescriptor.PrimitiveDataStride]
//   - [IMTLAccelerationStructureGeometryDescriptor.SetPrimitiveDataStride]
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor
type IMTLAccelerationStructureGeometryDescriptor interface {
	objectivec.IObject

	// Topic: Specifying base geometry properties

	// A label for the geometry structure, suitable for debugging.
	Label() string
	SetLabel(value string)
	// An index into the intersection table for determining which intersection function Metal calls when it intersects a ray with the acceleration structure.
	IntersectionFunctionTableOffset() uint
	SetIntersectionFunctionTableOffset(value uint)
	// A Boolean value that determines whether the geometry data in the acceleration structure needs to skip triangle-intersection tests.
	Opaque() bool
	SetOpaque(value bool)
	// A Boolean value that indicates whether Metal calls the ray-intersection test more than once per primitive on the structure.
	AllowDuplicateIntersectionFunctionInvocation() bool
	SetAllowDuplicateIntersectionFunctionInvocation(value bool)

	// Topic: Instance Properties

	PrimitiveDataBuffer() MTLBuffer
	SetPrimitiveDataBuffer(value MTLBuffer)
	PrimitiveDataBufferOffset() uint
	SetPrimitiveDataBufferOffset(value uint)
	PrimitiveDataElementSize() uint
	SetPrimitiveDataElementSize(value uint)
	PrimitiveDataStride() uint
	SetPrimitiveDataStride(value uint)
}

// Init initializes the instance.
func (a MTLAccelerationStructureGeometryDescriptor) Init() MTLAccelerationStructureGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureGeometryDescriptor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MTLAccelerationStructureGeometryDescriptor) Autorelease() MTLAccelerationStructureGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureGeometryDescriptor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLAccelerationStructureGeometryDescriptor creates a new MTLAccelerationStructureGeometryDescriptor instance.
func NewMTLAccelerationStructureGeometryDescriptor() MTLAccelerationStructureGeometryDescriptor {
	class := getMTLAccelerationStructureGeometryDescriptorClass()
	rv := objc.Send[MTLAccelerationStructureGeometryDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A label for the geometry structure, suitable for debugging.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/label
func (a MTLAccelerationStructureGeometryDescriptor) Label() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (a MTLAccelerationStructureGeometryDescriptor) SetLabel(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setLabel:"), objc.String(value))
}
// An index into the intersection table for determining which intersection
// function Metal calls when it intersects a ray with the acceleration
// structure.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/intersectionFunctionTableOffset
func (a MTLAccelerationStructureGeometryDescriptor) IntersectionFunctionTableOffset() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("intersectionFunctionTableOffset"))
	return rv
}
func (a MTLAccelerationStructureGeometryDescriptor) SetIntersectionFunctionTableOffset(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setIntersectionFunctionTableOffset:"), value)
}
// A Boolean value that determines whether the geometry data in the
// acceleration structure needs to skip triangle-intersection tests.
//
// # Discussion
// 
// By default, after Metal finds an intersection between a ray and a
// primitive, it runs your specified intersection function to determine
// whether the ray actually hit the primitive. If you specify that triangle
// geometry is opaque, Metal skips the intersection function and processes any
// intersection as a hit.
// 
// If you are using bounding box geometry, Metal calls your intersection
// function, passing a Boolean value that indicates that the bounding box that
// the ray intersected with is opaque.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/opaque
func (a MTLAccelerationStructureGeometryDescriptor) Opaque() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("opaque"))
	return rv
}
func (a MTLAccelerationStructureGeometryDescriptor) SetOpaque(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setOpaque:"), value)
}
// A Boolean value that indicates whether Metal calls the ray-intersection
// test more than once per primitive on the structure.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/allowDuplicateIntersectionFunctionInvocation
func (a MTLAccelerationStructureGeometryDescriptor) AllowDuplicateIntersectionFunctionInvocation() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("allowDuplicateIntersectionFunctionInvocation"))
	return rv
}
func (a MTLAccelerationStructureGeometryDescriptor) SetAllowDuplicateIntersectionFunctionInvocation(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setAllowDuplicateIntersectionFunctionInvocation:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/primitiveDataBuffer
func (a MTLAccelerationStructureGeometryDescriptor) PrimitiveDataBuffer() MTLBuffer {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("primitiveDataBuffer"))
	return MTLBufferObjectFromID(rv)
}
func (a MTLAccelerationStructureGeometryDescriptor) SetPrimitiveDataBuffer(value MTLBuffer) {
	objc.Send[struct{}](a.ID, objc.Sel("setPrimitiveDataBuffer:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/primitiveDataBufferOffset
func (a MTLAccelerationStructureGeometryDescriptor) PrimitiveDataBufferOffset() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("primitiveDataBufferOffset"))
	return rv
}
func (a MTLAccelerationStructureGeometryDescriptor) SetPrimitiveDataBufferOffset(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setPrimitiveDataBufferOffset:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/primitiveDataElementSize
func (a MTLAccelerationStructureGeometryDescriptor) PrimitiveDataElementSize() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("primitiveDataElementSize"))
	return rv
}
func (a MTLAccelerationStructureGeometryDescriptor) SetPrimitiveDataElementSize(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setPrimitiveDataElementSize:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/primitiveDataStride
func (a MTLAccelerationStructureGeometryDescriptor) PrimitiveDataStride() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("primitiveDataStride"))
	return rv
}
func (a MTLAccelerationStructureGeometryDescriptor) SetPrimitiveDataStride(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setPrimitiveDataStride:"), value)
}

