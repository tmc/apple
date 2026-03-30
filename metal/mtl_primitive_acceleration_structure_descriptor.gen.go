// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLPrimitiveAccelerationStructureDescriptor] class.
var (
	_MTLPrimitiveAccelerationStructureDescriptorClass     MTLPrimitiveAccelerationStructureDescriptorClass
	_MTLPrimitiveAccelerationStructureDescriptorClassOnce sync.Once
)

func getMTLPrimitiveAccelerationStructureDescriptorClass() MTLPrimitiveAccelerationStructureDescriptorClass {
	_MTLPrimitiveAccelerationStructureDescriptorClassOnce.Do(func() {
		_MTLPrimitiveAccelerationStructureDescriptorClass = MTLPrimitiveAccelerationStructureDescriptorClass{class: objc.GetClass("MTLPrimitiveAccelerationStructureDescriptor")}
	})
	return _MTLPrimitiveAccelerationStructureDescriptorClass
}

// GetMTLPrimitiveAccelerationStructureDescriptorClass returns the class object for MTLPrimitiveAccelerationStructureDescriptor.
func GetMTLPrimitiveAccelerationStructureDescriptorClass() MTLPrimitiveAccelerationStructureDescriptorClass {
	return getMTLPrimitiveAccelerationStructureDescriptorClass()
}

type MTLPrimitiveAccelerationStructureDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLPrimitiveAccelerationStructureDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLPrimitiveAccelerationStructureDescriptorClass) Alloc() MTLPrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[MTLPrimitiveAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of an acceleration structure that contains geometry
// primitives.
//
// # Overview
//
// Metal provides acceleration structures with a two-level hierarchy. The
// bottom layer consists of primitive acceleration structures, which instance
// acceleration structures in the top level reference.
//
// # Specifying geometry
//
//   - [MTLPrimitiveAccelerationStructureDescriptor.GeometryDescriptors]: An array that contains the individual pieces of geometry that compose the acceleration structure.
//   - [MTLPrimitiveAccelerationStructureDescriptor.SetGeometryDescriptors]
//
// # Specifying motion behavior
//
//   - [MTLPrimitiveAccelerationStructureDescriptor.MotionKeyframeCount]: The number of keyframes in the geometry data.
//   - [MTLPrimitiveAccelerationStructureDescriptor.SetMotionKeyframeCount]
//   - [MTLPrimitiveAccelerationStructureDescriptor.MotionStartTime]: The start time for the range of motion that the keyframe data describes.
//   - [MTLPrimitiveAccelerationStructureDescriptor.SetMotionStartTime]
//   - [MTLPrimitiveAccelerationStructureDescriptor.MotionEndTime]: The end time for the range of motion that the keyframe data describes.
//   - [MTLPrimitiveAccelerationStructureDescriptor.SetMotionEndTime]
//   - [MTLPrimitiveAccelerationStructureDescriptor.MotionStartBorderMode]: The mode to use when handling timestamps before the start time.
//   - [MTLPrimitiveAccelerationStructureDescriptor.SetMotionStartBorderMode]
//   - [MTLPrimitiveAccelerationStructureDescriptor.MotionEndBorderMode]: The mode to use when handling timestamps after the end time.
//   - [MTLPrimitiveAccelerationStructureDescriptor.SetMotionEndBorderMode]
//
// See: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor
type MTLPrimitiveAccelerationStructureDescriptor struct {
	MTLAccelerationStructureDescriptor
}

// MTLPrimitiveAccelerationStructureDescriptorFromID constructs a [MTLPrimitiveAccelerationStructureDescriptor] from an objc.ID.
//
// A description of an acceleration structure that contains geometry
// primitives.
func MTLPrimitiveAccelerationStructureDescriptorFromID(id objc.ID) MTLPrimitiveAccelerationStructureDescriptor {
	return MTLPrimitiveAccelerationStructureDescriptor{MTLAccelerationStructureDescriptor: MTLAccelerationStructureDescriptorFromID(id)}
}

// NOTE: MTLPrimitiveAccelerationStructureDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLPrimitiveAccelerationStructureDescriptor] class.
//
// # Specifying geometry
//
//   - [IMTLPrimitiveAccelerationStructureDescriptor.GeometryDescriptors]: An array that contains the individual pieces of geometry that compose the acceleration structure.
//   - [IMTLPrimitiveAccelerationStructureDescriptor.SetGeometryDescriptors]
//
// # Specifying motion behavior
//
//   - [IMTLPrimitiveAccelerationStructureDescriptor.MotionKeyframeCount]: The number of keyframes in the geometry data.
//   - [IMTLPrimitiveAccelerationStructureDescriptor.SetMotionKeyframeCount]
//   - [IMTLPrimitiveAccelerationStructureDescriptor.MotionStartTime]: The start time for the range of motion that the keyframe data describes.
//   - [IMTLPrimitiveAccelerationStructureDescriptor.SetMotionStartTime]
//   - [IMTLPrimitiveAccelerationStructureDescriptor.MotionEndTime]: The end time for the range of motion that the keyframe data describes.
//   - [IMTLPrimitiveAccelerationStructureDescriptor.SetMotionEndTime]
//   - [IMTLPrimitiveAccelerationStructureDescriptor.MotionStartBorderMode]: The mode to use when handling timestamps before the start time.
//   - [IMTLPrimitiveAccelerationStructureDescriptor.SetMotionStartBorderMode]
//   - [IMTLPrimitiveAccelerationStructureDescriptor.MotionEndBorderMode]: The mode to use when handling timestamps after the end time.
//   - [IMTLPrimitiveAccelerationStructureDescriptor.SetMotionEndBorderMode]
//
// See: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor
type IMTLPrimitiveAccelerationStructureDescriptor interface {
	IMTLAccelerationStructureDescriptor

	// Topic: Specifying geometry

	// An array that contains the individual pieces of geometry that compose the acceleration structure.
	GeometryDescriptors() []MTLAccelerationStructureGeometryDescriptor
	SetGeometryDescriptors(value []MTLAccelerationStructureGeometryDescriptor)

	// Topic: Specifying motion behavior

	// The number of keyframes in the geometry data.
	MotionKeyframeCount() uint
	SetMotionKeyframeCount(value uint)
	// The start time for the range of motion that the keyframe data describes.
	MotionStartTime() float32
	SetMotionStartTime(value float32)
	// The end time for the range of motion that the keyframe data describes.
	MotionEndTime() float32
	SetMotionEndTime(value float32)
	// The mode to use when handling timestamps before the start time.
	MotionStartBorderMode() MTLMotionBorderMode
	SetMotionStartBorderMode(value MTLMotionBorderMode)
	// The mode to use when handling timestamps after the end time.
	MotionEndBorderMode() MTLMotionBorderMode
	SetMotionEndBorderMode(value MTLMotionBorderMode)
}

// Init initializes the instance.
func (p MTLPrimitiveAccelerationStructureDescriptor) Init() MTLPrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[MTLPrimitiveAccelerationStructureDescriptor](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MTLPrimitiveAccelerationStructureDescriptor) Autorelease() MTLPrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[MTLPrimitiveAccelerationStructureDescriptor](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLPrimitiveAccelerationStructureDescriptor creates a new MTLPrimitiveAccelerationStructureDescriptor instance.
func NewMTLPrimitiveAccelerationStructureDescriptor() MTLPrimitiveAccelerationStructureDescriptor {
	class := getMTLPrimitiveAccelerationStructureDescriptorClass()
	rv := objc.Send[MTLPrimitiveAccelerationStructureDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An array that contains the individual pieces of geometry that compose the
// acceleration structure.
//
// # Discussion
//
// The value of the [MotionKeyframeCount] property determines what kinds of
// geometry descriptors you can assign to this property and how you need to
// configure them.
//
// If the value of [MotionKeyframeCount] is greater than 1, then the geometry
// descriptors need to be either
// [MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor] or
// [MTLAccelerationStructureMotionTriangleGeometryDescriptor] objects.
// Further, you need to provide exactly that many keyframes of data when
// creating those geometry descriptors. If [MotionKeyframeCount] is 1, use
// [MTLAccelerationStructureBoundingBoxGeometryDescriptor] or
// [MTLAccelerationStructureTriangleGeometryDescriptor] objects instead.
//
// See: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor/geometryDescriptors
func (p MTLPrimitiveAccelerationStructureDescriptor) GeometryDescriptors() []MTLAccelerationStructureGeometryDescriptor {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("geometryDescriptors"))
	return objc.ConvertSlice(rv, func(id objc.ID) MTLAccelerationStructureGeometryDescriptor {
		return MTLAccelerationStructureGeometryDescriptorFromID(id)
	})
}
func (p MTLPrimitiveAccelerationStructureDescriptor) SetGeometryDescriptors(value []MTLAccelerationStructureGeometryDescriptor) {
	objc.Send[struct{}](p.ID, objc.Sel("setGeometryDescriptors:"), objectivec.IObjectSliceToNSArray(value))
}

// The number of keyframes in the geometry data.
//
// # Discussion
//
// The default value is `1`. If the value is greater than `1`, all geometry
// descriptors that you attach to this descriptor need to be motion
// descriptors, and each needs to have exactly that many
// [MTLMotionKeyframeData] objects.
//
// See: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor/motionKeyframeCount
func (p MTLPrimitiveAccelerationStructureDescriptor) MotionKeyframeCount() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("motionKeyframeCount"))
	return rv
}
func (p MTLPrimitiveAccelerationStructureDescriptor) SetMotionKeyframeCount(value uint) {
	objc.Send[struct{}](p.ID, objc.Sel("setMotionKeyframeCount:"), value)
}

// The start time for the range of motion that the keyframe data describes.
//
// # Discussion
//
// The default value is `0.0f`.
//
// See: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor/motionStartTime
func (p MTLPrimitiveAccelerationStructureDescriptor) MotionStartTime() float32 {
	rv := objc.Send[float32](p.ID, objc.Sel("motionStartTime"))
	return rv
}
func (p MTLPrimitiveAccelerationStructureDescriptor) SetMotionStartTime(value float32) {
	objc.Send[struct{}](p.ID, objc.Sel("setMotionStartTime:"), value)
}

// The end time for the range of motion that the keyframe data describes.
//
// # Discussion
//
// The default value is `1.0f`.
//
// See: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor/motionEndTime
func (p MTLPrimitiveAccelerationStructureDescriptor) MotionEndTime() float32 {
	rv := objc.Send[float32](p.ID, objc.Sel("motionEndTime"))
	return rv
}
func (p MTLPrimitiveAccelerationStructureDescriptor) SetMotionEndTime(value float32) {
	objc.Send[struct{}](p.ID, objc.Sel("setMotionEndTime:"), value)
}

// The mode to use when handling timestamps before the start time.
//
// # Discussion
//
// The default value is [MTLMotionBorderMode.clamp].
//
// See: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor/motionStartBorderMode
//
// [MTLMotionBorderMode.clamp]: https://developer.apple.com/documentation/Metal/MTLMotionBorderMode/clamp
func (p MTLPrimitiveAccelerationStructureDescriptor) MotionStartBorderMode() MTLMotionBorderMode {
	rv := objc.Send[MTLMotionBorderMode](p.ID, objc.Sel("motionStartBorderMode"))
	return MTLMotionBorderMode(rv)
}
func (p MTLPrimitiveAccelerationStructureDescriptor) SetMotionStartBorderMode(value MTLMotionBorderMode) {
	objc.Send[struct{}](p.ID, objc.Sel("setMotionStartBorderMode:"), value)
}

// The mode to use when handling timestamps after the end time.
//
// # Discussion
//
// The default value is [MTLMotionBorderMode.clamp].
//
// See: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor/motionEndBorderMode
//
// [MTLMotionBorderMode.clamp]: https://developer.apple.com/documentation/Metal/MTLMotionBorderMode/clamp
func (p MTLPrimitiveAccelerationStructureDescriptor) MotionEndBorderMode() MTLMotionBorderMode {
	rv := objc.Send[MTLMotionBorderMode](p.ID, objc.Sel("motionEndBorderMode"))
	return MTLMotionBorderMode(rv)
}
func (p MTLPrimitiveAccelerationStructureDescriptor) SetMotionEndBorderMode(value MTLMotionBorderMode) {
	objc.Send[struct{}](p.ID, objc.Sel("setMotionEndBorderMode:"), value)
}
