// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MTL4AccelerationStructureBoundingBoxGeometryDescriptor] class.
var (
	_MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass     MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass
	_MTL4AccelerationStructureBoundingBoxGeometryDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureBoundingBoxGeometryDescriptorClass() MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass {
	_MTL4AccelerationStructureBoundingBoxGeometryDescriptorClassOnce.Do(func() {
		_MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass = MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass{class: objc.GetClass("MTL4AccelerationStructureBoundingBoxGeometryDescriptor")}
	})
	return _MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass
}

// GetMTL4AccelerationStructureBoundingBoxGeometryDescriptorClass returns the class object for MTL4AccelerationStructureBoundingBoxGeometryDescriptor.
func GetMTL4AccelerationStructureBoundingBoxGeometryDescriptorClass() MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass {
	return getMTL4AccelerationStructureBoundingBoxGeometryDescriptorClass()
}

type MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass) Alloc() MTL4AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureBoundingBoxGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// Describes bounding-box geometry suitable for ray tracing.
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
//   - [MTL4AccelerationStructureBoundingBoxGeometryDescriptor.BoundingBoxBuffer]: References a buffer containing bounding box data in [MTLAxisAlignedBoundingBoxes] format.
//   - [MTL4AccelerationStructureBoundingBoxGeometryDescriptor.SetBoundingBoxBuffer]
//   - [MTL4AccelerationStructureBoundingBoxGeometryDescriptor.BoundingBoxCount]: Describes the number of bounding boxes the `boundingBoxBuffer` contains.
//   - [MTL4AccelerationStructureBoundingBoxGeometryDescriptor.SetBoundingBoxCount]
//   - [MTL4AccelerationStructureBoundingBoxGeometryDescriptor.BoundingBoxStride]: Assigns the stride, in bytes, between bounding boxes in the bounding box buffer `boundingBoxBuffer` references.
//   - [MTL4AccelerationStructureBoundingBoxGeometryDescriptor.SetBoundingBoxStride]
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureBoundingBoxGeometryDescriptor
type MTL4AccelerationStructureBoundingBoxGeometryDescriptor struct {
	MTL4AccelerationStructureGeometryDescriptor
}

// MTL4AccelerationStructureBoundingBoxGeometryDescriptorFromID constructs a [MTL4AccelerationStructureBoundingBoxGeometryDescriptor] from an objc.ID.
//
// Describes bounding-box geometry suitable for ray tracing.
func MTL4AccelerationStructureBoundingBoxGeometryDescriptorFromID(id objc.ID) MTL4AccelerationStructureBoundingBoxGeometryDescriptor {
	return MTL4AccelerationStructureBoundingBoxGeometryDescriptor{MTL4AccelerationStructureGeometryDescriptor: MTL4AccelerationStructureGeometryDescriptorFromID(id)}
}
// NOTE: MTL4AccelerationStructureBoundingBoxGeometryDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTL4AccelerationStructureBoundingBoxGeometryDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4AccelerationStructureBoundingBoxGeometryDescriptor.BoundingBoxBuffer]: References a buffer containing bounding box data in [MTLAxisAlignedBoundingBoxes] format.
//   - [IMTL4AccelerationStructureBoundingBoxGeometryDescriptor.SetBoundingBoxBuffer]
//   - [IMTL4AccelerationStructureBoundingBoxGeometryDescriptor.BoundingBoxCount]: Describes the number of bounding boxes the `boundingBoxBuffer` contains.
//   - [IMTL4AccelerationStructureBoundingBoxGeometryDescriptor.SetBoundingBoxCount]
//   - [IMTL4AccelerationStructureBoundingBoxGeometryDescriptor.BoundingBoxStride]: Assigns the stride, in bytes, between bounding boxes in the bounding box buffer `boundingBoxBuffer` references.
//   - [IMTL4AccelerationStructureBoundingBoxGeometryDescriptor.SetBoundingBoxStride]
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureBoundingBoxGeometryDescriptor
type IMTL4AccelerationStructureBoundingBoxGeometryDescriptor interface {
	IMTL4AccelerationStructureGeometryDescriptor

	// Topic: Instance Properties

	// References a buffer containing bounding box data in [MTLAxisAlignedBoundingBoxes] format.
	BoundingBoxBuffer() MTL4BufferRange
	SetBoundingBoxBuffer(value MTL4BufferRange)
	// Describes the number of bounding boxes the `boundingBoxBuffer` contains.
	BoundingBoxCount() uint
	SetBoundingBoxCount(value uint)
	// Assigns the stride, in bytes, between bounding boxes in the bounding box buffer `boundingBoxBuffer` references.
	BoundingBoxStride() uint
	SetBoundingBoxStride(value uint)
}





// Init initializes the instance.
func (m MTL4AccelerationStructureBoundingBoxGeometryDescriptor) Init() MTL4AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureBoundingBoxGeometryDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4AccelerationStructureBoundingBoxGeometryDescriptor) Autorelease() MTL4AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureBoundingBoxGeometryDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureBoundingBoxGeometryDescriptor creates a new MTL4AccelerationStructureBoundingBoxGeometryDescriptor instance.
func NewMTL4AccelerationStructureBoundingBoxGeometryDescriptor() MTL4AccelerationStructureBoundingBoxGeometryDescriptor {
	class := getMTL4AccelerationStructureBoundingBoxGeometryDescriptorClass()
	rv := objc.Send[MTL4AccelerationStructureBoundingBoxGeometryDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// References a buffer containing bounding box data in
// [MTLAxisAlignedBoundingBoxes] format.
//
// # Discussion
// 
// You are responsible for ensuring the buffer address of the range is not
// zero.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxBuffer
func (m MTL4AccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxBuffer() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("boundingBoxBuffer"))
	return MTL4BufferRange(rv)
}
func (m MTL4AccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxBuffer(value MTL4BufferRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setBoundingBoxBuffer:"), value)
}



// Describes the number of bounding boxes the `boundingBoxBuffer` contains.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxCount
func (m MTL4AccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("boundingBoxCount"))
	return rv
}
func (m MTL4AccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setBoundingBoxCount:"), value)
}



// Assigns the stride, in bytes, between bounding boxes in the bounding box
// buffer `boundingBoxBuffer` references.
//
// # Discussion
// 
// You are responsible for ensuring this stride is at least 24 bytes and a
// multiple of 4 bytes.
// 
// This property defaults to `24` bytes.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxStride
func (m MTL4AccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxStride() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("boundingBoxStride"))
	return rv
}
func (m MTL4AccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxStride(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setBoundingBoxStride:"), value)
}
























