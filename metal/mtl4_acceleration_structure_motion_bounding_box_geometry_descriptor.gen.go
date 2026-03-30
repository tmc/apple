// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor] class.
var (
	_MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass     MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass
	_MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass() MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass {
	_MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClassOnce.Do(func() {
		_MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass = MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass{class: objc.GetClass("MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor")}
	})
	return _MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass
}

// GetMTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass returns the class object for MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor.
func GetMTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass() MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass {
	return getMTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass()
}

type MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass) Alloc() MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Describes motion bounding box geometry, suitable for motion ray tracing.
//
// # Overview
//
// You use bounding boxes to implement procedural geometry for ray tracing,
// such as spheres or any other shape you define by using intersection
// functions.
//
// Use a [MTLResidencySet] to mark residency of all buffers this descriptor
// references when you build this acceleration structure.
//
// # Instance Properties
//
//   - [MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor.BoundingBoxBuffers]: Configures a reference to a buffer where each entry contains a reference to a buffer of bounding boxes.
//   - [MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor.SetBoundingBoxBuffers]
//   - [MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor.BoundingBoxCount]: Declares the number of bounding boxes in each buffer that `boundingBoxBuffer` references.
//   - [MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor.SetBoundingBoxCount]
//   - [MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor.BoundingBoxStride]: Declares the stride, in bytes, between bounding boxes in the bounding box buffers each entry in `boundingBoxBuffer` references.
//   - [MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor.SetBoundingBoxStride]
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor
type MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor struct {
	MTL4AccelerationStructureGeometryDescriptor
}

// MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorFromID constructs a [MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor] from an objc.ID.
//
// Describes motion bounding box geometry, suitable for motion ray tracing.
func MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorFromID(id objc.ID) MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	return MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor{MTL4AccelerationStructureGeometryDescriptor: MTL4AccelerationStructureGeometryDescriptorFromID(id)}
}

// NOTE: MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor.BoundingBoxBuffers]: Configures a reference to a buffer where each entry contains a reference to a buffer of bounding boxes.
//   - [IMTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor.SetBoundingBoxBuffers]
//   - [IMTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor.BoundingBoxCount]: Declares the number of bounding boxes in each buffer that `boundingBoxBuffer` references.
//   - [IMTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor.SetBoundingBoxCount]
//   - [IMTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor.BoundingBoxStride]: Declares the stride, in bytes, between bounding boxes in the bounding box buffers each entry in `boundingBoxBuffer` references.
//   - [IMTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor.SetBoundingBoxStride]
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor
type IMTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor interface {
	IMTL4AccelerationStructureGeometryDescriptor

	// Topic: Instance Properties

	// Configures a reference to a buffer where each entry contains a reference to a buffer of bounding boxes.
	BoundingBoxBuffers() MTL4BufferRange
	SetBoundingBoxBuffers(value MTL4BufferRange)
	// Declares the number of bounding boxes in each buffer that `boundingBoxBuffer` references.
	BoundingBoxCount() uint
	SetBoundingBoxCount(value uint)
	// Declares the stride, in bytes, between bounding boxes in the bounding box buffers each entry in `boundingBoxBuffer` references.
	BoundingBoxStride() uint
	SetBoundingBoxStride(value uint)
}

// Init initializes the instance.
func (m MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor) Init() MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor) Autorelease() MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor creates a new MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor instance.
func NewMTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor() MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	class := getMTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass()
	rv := objc.Send[MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Configures a reference to a buffer where each entry contains a reference to
// a buffer of bounding boxes.
//
// # Discussion
//
// This property references a buffer that conceptually represents an array
// with one entry for each keyframe in the motion animation. Each one of these
// entries consists of a [MTL4BufferRange] that, in turn, references a vertex
// buffer containing the bounding box data for the keyframe.
//
// You are responsible for ensuring the buffer address is not zero for the
// top-level buffer, as well as for all the vertex buffers it references.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor/boundingBoxBuffers
//
// [MTL4BufferRange]: https://developer.apple.com/documentation/Metal/MTL4BufferRange
func (m MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxBuffers() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("boundingBoxBuffers"))
	return MTL4BufferRange(rv)
}
func (m MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxBuffers(value MTL4BufferRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setBoundingBoxBuffers:"), value)
}

// Declares the number of bounding boxes in each buffer that
// `boundingBoxBuffer` references.
//
// # Discussion
//
// All keyframes share the same bounding box count.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor/boundingBoxCount
func (m MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("boundingBoxCount"))
	return rv
}
func (m MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setBoundingBoxCount:"), value)
}

// Declares the stride, in bytes, between bounding boxes in the bounding box
// buffers each entry in `boundingBoxBuffer` references.
//
// # Discussion
//
// All keyframes share the same bounding box stride. You are responsible for
// ensuring this stride is at least 24 bytes and a multiple of 4 bytes.
//
// This property defaults to `24` bytes.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor/boundingBoxStride
func (m MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxStride() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("boundingBoxStride"))
	return rv
}
func (m MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxStride(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setBoundingBoxStride:"), value)
}
