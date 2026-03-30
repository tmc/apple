// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MTLAccelerationStructureBoundingBoxGeometryDescriptor] class.
var (
	_MTLAccelerationStructureBoundingBoxGeometryDescriptorClass     MTLAccelerationStructureBoundingBoxGeometryDescriptorClass
	_MTLAccelerationStructureBoundingBoxGeometryDescriptorClassOnce sync.Once
)

func getMTLAccelerationStructureBoundingBoxGeometryDescriptorClass() MTLAccelerationStructureBoundingBoxGeometryDescriptorClass {
	_MTLAccelerationStructureBoundingBoxGeometryDescriptorClassOnce.Do(func() {
		_MTLAccelerationStructureBoundingBoxGeometryDescriptorClass = MTLAccelerationStructureBoundingBoxGeometryDescriptorClass{class: objc.GetClass("MTLAccelerationStructureBoundingBoxGeometryDescriptor")}
	})
	return _MTLAccelerationStructureBoundingBoxGeometryDescriptorClass
}

// GetMTLAccelerationStructureBoundingBoxGeometryDescriptorClass returns the class object for MTLAccelerationStructureBoundingBoxGeometryDescriptor.
func GetMTLAccelerationStructureBoundingBoxGeometryDescriptorClass() MTLAccelerationStructureBoundingBoxGeometryDescriptorClass {
	return getMTLAccelerationStructureBoundingBoxGeometryDescriptorClass()
}

type MTLAccelerationStructureBoundingBoxGeometryDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLAccelerationStructureBoundingBoxGeometryDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLAccelerationStructureBoundingBoxGeometryDescriptorClass) Alloc() MTLAccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureBoundingBoxGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of a list of bounding boxes to turn into an acceleration
// structure.
//
// # Specifying the number of bounding boxes
//
//   - [MTLAccelerationStructureBoundingBoxGeometryDescriptor.BoundingBoxCount]: The number of bounding boxes in the bounding box buffer.
//   - [MTLAccelerationStructureBoundingBoxGeometryDescriptor.SetBoundingBoxCount]
//
// # Specifying bounding boxes data
//
//   - [MTLAccelerationStructureBoundingBoxGeometryDescriptor.BoundingBoxBuffer]: A buffer that contains an array of bounding box structures.
//   - [MTLAccelerationStructureBoundingBoxGeometryDescriptor.SetBoundingBoxBuffer]
//   - [MTLAccelerationStructureBoundingBoxGeometryDescriptor.BoundingBoxBufferOffset]: The offset, in bytes, to the first bounding box in the buffer.
//   - [MTLAccelerationStructureBoundingBoxGeometryDescriptor.SetBoundingBoxBufferOffset]
//   - [MTLAccelerationStructureBoundingBoxGeometryDescriptor.BoundingBoxStride]: The stride, in bytes, between bounding boxes in the buffer.
//   - [MTLAccelerationStructureBoundingBoxGeometryDescriptor.SetBoundingBoxStride]
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureBoundingBoxGeometryDescriptor
type MTLAccelerationStructureBoundingBoxGeometryDescriptor struct {
	MTLAccelerationStructureGeometryDescriptor
}

// MTLAccelerationStructureBoundingBoxGeometryDescriptorFromID constructs a [MTLAccelerationStructureBoundingBoxGeometryDescriptor] from an objc.ID.
//
// A description of a list of bounding boxes to turn into an acceleration
// structure.
func MTLAccelerationStructureBoundingBoxGeometryDescriptorFromID(id objc.ID) MTLAccelerationStructureBoundingBoxGeometryDescriptor {
	return MTLAccelerationStructureBoundingBoxGeometryDescriptor{MTLAccelerationStructureGeometryDescriptor: MTLAccelerationStructureGeometryDescriptorFromID(id)}
}

// NOTE: MTLAccelerationStructureBoundingBoxGeometryDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLAccelerationStructureBoundingBoxGeometryDescriptor] class.
//
// # Specifying the number of bounding boxes
//
//   - [IMTLAccelerationStructureBoundingBoxGeometryDescriptor.BoundingBoxCount]: The number of bounding boxes in the bounding box buffer.
//   - [IMTLAccelerationStructureBoundingBoxGeometryDescriptor.SetBoundingBoxCount]
//
// # Specifying bounding boxes data
//
//   - [IMTLAccelerationStructureBoundingBoxGeometryDescriptor.BoundingBoxBuffer]: A buffer that contains an array of bounding box structures.
//   - [IMTLAccelerationStructureBoundingBoxGeometryDescriptor.SetBoundingBoxBuffer]
//   - [IMTLAccelerationStructureBoundingBoxGeometryDescriptor.BoundingBoxBufferOffset]: The offset, in bytes, to the first bounding box in the buffer.
//   - [IMTLAccelerationStructureBoundingBoxGeometryDescriptor.SetBoundingBoxBufferOffset]
//   - [IMTLAccelerationStructureBoundingBoxGeometryDescriptor.BoundingBoxStride]: The stride, in bytes, between bounding boxes in the buffer.
//   - [IMTLAccelerationStructureBoundingBoxGeometryDescriptor.SetBoundingBoxStride]
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureBoundingBoxGeometryDescriptor
type IMTLAccelerationStructureBoundingBoxGeometryDescriptor interface {
	IMTLAccelerationStructureGeometryDescriptor

	// Topic: Specifying the number of bounding boxes

	// The number of bounding boxes in the bounding box buffer.
	BoundingBoxCount() uint
	SetBoundingBoxCount(value uint)

	// Topic: Specifying bounding boxes data

	// A buffer that contains an array of bounding box structures.
	BoundingBoxBuffer() MTLBuffer
	SetBoundingBoxBuffer(value MTLBuffer)
	// The offset, in bytes, to the first bounding box in the buffer.
	BoundingBoxBufferOffset() uint
	SetBoundingBoxBufferOffset(value uint)
	// The stride, in bytes, between bounding boxes in the buffer.
	BoundingBoxStride() uint
	SetBoundingBoxStride(value uint)
}

// Init initializes the instance.
func (a MTLAccelerationStructureBoundingBoxGeometryDescriptor) Init() MTLAccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureBoundingBoxGeometryDescriptor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MTLAccelerationStructureBoundingBoxGeometryDescriptor) Autorelease() MTLAccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureBoundingBoxGeometryDescriptor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLAccelerationStructureBoundingBoxGeometryDescriptor creates a new MTLAccelerationStructureBoundingBoxGeometryDescriptor instance.
func NewMTLAccelerationStructureBoundingBoxGeometryDescriptor() MTLAccelerationStructureBoundingBoxGeometryDescriptor {
	class := getMTLAccelerationStructureBoundingBoxGeometryDescriptorClass()
	rv := objc.Send[MTLAccelerationStructureBoundingBoxGeometryDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The number of bounding boxes in the bounding box buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxCount
func (a MTLAccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxCount() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("boundingBoxCount"))
	return rv
}
func (a MTLAccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxCount(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setBoundingBoxCount:"), value)
}

// A buffer that contains an array of bounding box structures.
//
// # Discussion
//
// The buffer contains an array of [MTLAxisAlignedBoundingBox] structures, one
// for each bounding box in the geometry.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxBuffer
//
// [MTLAxisAlignedBoundingBox]: https://developer.apple.com/documentation/Metal/MTLAxisAlignedBoundingBox-c.struct
func (a MTLAccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxBuffer() MTLBuffer {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("boundingBoxBuffer"))
	return MTLBufferObjectFromID(rv)
}
func (a MTLAccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxBuffer(value MTLBuffer) {
	objc.Send[struct{}](a.ID, objc.Sel("setBoundingBoxBuffer:"), value)
}

// The offset, in bytes, to the first bounding box in the buffer.
//
// # Discussion
//
// The offset needs to be a multiple of [BoundingBoxStride]. Check the [Metal
// feature set tables (PDF)] for potential alignment restrictions.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxBufferOffset
//
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
func (a MTLAccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxBufferOffset() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("boundingBoxBufferOffset"))
	return rv
}
func (a MTLAccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxBufferOffset(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setBoundingBoxBufferOffset:"), value)
}

// The stride, in bytes, between bounding boxes in the buffer.
//
// # Discussion
//
// The stride needs be at least 24 bytes, and be a multiple of 4 bytes. The
// default value is 24 bytes.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxStride
func (a MTLAccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxStride() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("boundingBoxStride"))
	return rv
}
func (a MTLAccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxStride(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setBoundingBoxStride:"), value)
}
