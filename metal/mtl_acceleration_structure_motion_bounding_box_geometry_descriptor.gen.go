// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor] class.
var (
	_MTLAccelerationStructureMotionBoundingBoxGeometryDescriptorClass     MTLAccelerationStructureMotionBoundingBoxGeometryDescriptorClass
	_MTLAccelerationStructureMotionBoundingBoxGeometryDescriptorClassOnce sync.Once
)

func getMTLAccelerationStructureMotionBoundingBoxGeometryDescriptorClass() MTLAccelerationStructureMotionBoundingBoxGeometryDescriptorClass {
	_MTLAccelerationStructureMotionBoundingBoxGeometryDescriptorClassOnce.Do(func() {
		_MTLAccelerationStructureMotionBoundingBoxGeometryDescriptorClass = MTLAccelerationStructureMotionBoundingBoxGeometryDescriptorClass{class: objc.GetClass("MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor")}
	})
	return _MTLAccelerationStructureMotionBoundingBoxGeometryDescriptorClass
}

// GetMTLAccelerationStructureMotionBoundingBoxGeometryDescriptorClass returns the class object for MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor.
func GetMTLAccelerationStructureMotionBoundingBoxGeometryDescriptorClass() MTLAccelerationStructureMotionBoundingBoxGeometryDescriptorClass {
	return getMTLAccelerationStructureMotionBoundingBoxGeometryDescriptorClass()
}

type MTLAccelerationStructureMotionBoundingBoxGeometryDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLAccelerationStructureMotionBoundingBoxGeometryDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLAccelerationStructureMotionBoundingBoxGeometryDescriptorClass) Alloc() MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of a list of bounding boxes, as motion keyframe data, to turn
// into an acceleration structure.
//
// # Specifying the number of bounding boxes
//
//   - [MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor.BoundingBoxCount]: The number of bounding boxes in each bounding box buffer.
//   - [MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor.SetBoundingBoxCount]
//
// # Specifying bounding boxes data
//
//   - [MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor.BoundingBoxBuffers]: A array of motion keyframes, each containing bounding box data.
//   - [MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor.SetBoundingBoxBuffers]
//   - [MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor.BoundingBoxStride]: The stride, in bytes, between bounding boxes in each buffer.
//   - [MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor.SetBoundingBoxStride]
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor
type MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor struct {
	MTLAccelerationStructureGeometryDescriptor
}

// MTLAccelerationStructureMotionBoundingBoxGeometryDescriptorFromID constructs a [MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor] from an objc.ID.
//
// A description of a list of bounding boxes, as motion keyframe data, to turn
// into an acceleration structure.
func MTLAccelerationStructureMotionBoundingBoxGeometryDescriptorFromID(id objc.ID) MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor {
	return MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor{MTLAccelerationStructureGeometryDescriptor: MTLAccelerationStructureGeometryDescriptorFromID(id)}
}

// NOTE: MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor] class.
//
// # Specifying the number of bounding boxes
//
//   - [IMTLAccelerationStructureMotionBoundingBoxGeometryDescriptor.BoundingBoxCount]: The number of bounding boxes in each bounding box buffer.
//   - [IMTLAccelerationStructureMotionBoundingBoxGeometryDescriptor.SetBoundingBoxCount]
//
// # Specifying bounding boxes data
//
//   - [IMTLAccelerationStructureMotionBoundingBoxGeometryDescriptor.BoundingBoxBuffers]: A array of motion keyframes, each containing bounding box data.
//   - [IMTLAccelerationStructureMotionBoundingBoxGeometryDescriptor.SetBoundingBoxBuffers]
//   - [IMTLAccelerationStructureMotionBoundingBoxGeometryDescriptor.BoundingBoxStride]: The stride, in bytes, between bounding boxes in each buffer.
//   - [IMTLAccelerationStructureMotionBoundingBoxGeometryDescriptor.SetBoundingBoxStride]
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor
type IMTLAccelerationStructureMotionBoundingBoxGeometryDescriptor interface {
	IMTLAccelerationStructureGeometryDescriptor

	// Topic: Specifying the number of bounding boxes

	// The number of bounding boxes in each bounding box buffer.
	BoundingBoxCount() uint
	SetBoundingBoxCount(value uint)

	// Topic: Specifying bounding boxes data

	// A array of motion keyframes, each containing bounding box data.
	BoundingBoxBuffers() []MTLMotionKeyframeData
	SetBoundingBoxBuffers(value []MTLMotionKeyframeData)
	// The stride, in bytes, between bounding boxes in each buffer.
	BoundingBoxStride() uint
	SetBoundingBoxStride(value uint)
}

// Init initializes the instance.
func (a MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) Init() MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) Autorelease() MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLAccelerationStructureMotionBoundingBoxGeometryDescriptor creates a new MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor instance.
func NewMTLAccelerationStructureMotionBoundingBoxGeometryDescriptor() MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor {
	class := getMTLAccelerationStructureMotionBoundingBoxGeometryDescriptorClass()
	rv := objc.Send[MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The number of bounding boxes in each bounding box buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor/boundingBoxCount
func (a MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxCount() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("boundingBoxCount"))
	return rv
}
func (a MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxCount(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setBoundingBoxCount:"), value)
}

// A array of motion keyframes, each containing bounding box data.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor/boundingBoxBuffers
func (a MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxBuffers() []MTLMotionKeyframeData {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("boundingBoxBuffers"))
	return objc.ConvertSlice(rv, func(id objc.ID) MTLMotionKeyframeData {
		return MTLMotionKeyframeDataFromID(id)
	})
}
func (a MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxBuffers(value []MTLMotionKeyframeData) {
	objc.Send[struct{}](a.ID, objc.Sel("setBoundingBoxBuffers:"), objectivec.IObjectSliceToNSArray(value))
}

// The stride, in bytes, between bounding boxes in each buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor/boundingBoxStride
func (a MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxStride() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("boundingBoxStride"))
	return rv
}
func (a MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxStride(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setBoundingBoxStride:"), value)
}
