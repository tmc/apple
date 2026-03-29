// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4PrimitiveAccelerationStructureDescriptor] class.
var (
	_MTL4PrimitiveAccelerationStructureDescriptorClass     MTL4PrimitiveAccelerationStructureDescriptorClass
	_MTL4PrimitiveAccelerationStructureDescriptorClassOnce sync.Once
)

func getMTL4PrimitiveAccelerationStructureDescriptorClass() MTL4PrimitiveAccelerationStructureDescriptorClass {
	_MTL4PrimitiveAccelerationStructureDescriptorClassOnce.Do(func() {
		_MTL4PrimitiveAccelerationStructureDescriptorClass = MTL4PrimitiveAccelerationStructureDescriptorClass{class: objc.GetClass("MTL4PrimitiveAccelerationStructureDescriptor")}
	})
	return _MTL4PrimitiveAccelerationStructureDescriptorClass
}

// GetMTL4PrimitiveAccelerationStructureDescriptorClass returns the class object for MTL4PrimitiveAccelerationStructureDescriptor.
func GetMTL4PrimitiveAccelerationStructureDescriptorClass() MTL4PrimitiveAccelerationStructureDescriptorClass {
	return getMTL4PrimitiveAccelerationStructureDescriptorClass()
}

type MTL4PrimitiveAccelerationStructureDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTL4PrimitiveAccelerationStructureDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4PrimitiveAccelerationStructureDescriptorClass) Alloc() MTL4PrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[MTL4PrimitiveAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Descriptor for a primitive acceleration structure that directly references
// geometric shapes, such as triangles and bounding boxes.
//
// # Instance Properties
//
//   - [MTL4PrimitiveAccelerationStructureDescriptor.GeometryDescriptors]: Associates the array of geometry descriptors that comprise this primitive acceleration structure.
//   - [MTL4PrimitiveAccelerationStructureDescriptor.SetGeometryDescriptors]
//   - [MTL4PrimitiveAccelerationStructureDescriptor.MotionEndBorderMode]: Configures the motion border mode.
//   - [MTL4PrimitiveAccelerationStructureDescriptor.SetMotionEndBorderMode]
//   - [MTL4PrimitiveAccelerationStructureDescriptor.MotionEndTime]: Configures the motion end time for this geometry.
//   - [MTL4PrimitiveAccelerationStructureDescriptor.SetMotionEndTime]
//   - [MTL4PrimitiveAccelerationStructureDescriptor.MotionKeyframeCount]: Sets the motion keyframe count.
//   - [MTL4PrimitiveAccelerationStructureDescriptor.SetMotionKeyframeCount]
//   - [MTL4PrimitiveAccelerationStructureDescriptor.MotionStartBorderMode]: Configures the behavior when the ray-tracing system samples the acceleration structure before the motion start time.
//   - [MTL4PrimitiveAccelerationStructureDescriptor.SetMotionStartBorderMode]
//   - [MTL4PrimitiveAccelerationStructureDescriptor.MotionStartTime]: Configures the motion start time for this geometry.
//   - [MTL4PrimitiveAccelerationStructureDescriptor.SetMotionStartTime]
//
// See: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor
type MTL4PrimitiveAccelerationStructureDescriptor struct {
	MTL4AccelerationStructureDescriptor
}

// MTL4PrimitiveAccelerationStructureDescriptorFromID constructs a [MTL4PrimitiveAccelerationStructureDescriptor] from an objc.ID.
//
// Descriptor for a primitive acceleration structure that directly references
// geometric shapes, such as triangles and bounding boxes.
func MTL4PrimitiveAccelerationStructureDescriptorFromID(id objc.ID) MTL4PrimitiveAccelerationStructureDescriptor {
	return MTL4PrimitiveAccelerationStructureDescriptor{MTL4AccelerationStructureDescriptor: MTL4AccelerationStructureDescriptorFromID(id)}
}
// NOTE: MTL4PrimitiveAccelerationStructureDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4PrimitiveAccelerationStructureDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4PrimitiveAccelerationStructureDescriptor.GeometryDescriptors]: Associates the array of geometry descriptors that comprise this primitive acceleration structure.
//   - [IMTL4PrimitiveAccelerationStructureDescriptor.SetGeometryDescriptors]
//   - [IMTL4PrimitiveAccelerationStructureDescriptor.MotionEndBorderMode]: Configures the motion border mode.
//   - [IMTL4PrimitiveAccelerationStructureDescriptor.SetMotionEndBorderMode]
//   - [IMTL4PrimitiveAccelerationStructureDescriptor.MotionEndTime]: Configures the motion end time for this geometry.
//   - [IMTL4PrimitiveAccelerationStructureDescriptor.SetMotionEndTime]
//   - [IMTL4PrimitiveAccelerationStructureDescriptor.MotionKeyframeCount]: Sets the motion keyframe count.
//   - [IMTL4PrimitiveAccelerationStructureDescriptor.SetMotionKeyframeCount]
//   - [IMTL4PrimitiveAccelerationStructureDescriptor.MotionStartBorderMode]: Configures the behavior when the ray-tracing system samples the acceleration structure before the motion start time.
//   - [IMTL4PrimitiveAccelerationStructureDescriptor.SetMotionStartBorderMode]
//   - [IMTL4PrimitiveAccelerationStructureDescriptor.MotionStartTime]: Configures the motion start time for this geometry.
//   - [IMTL4PrimitiveAccelerationStructureDescriptor.SetMotionStartTime]
//
// See: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor
type IMTL4PrimitiveAccelerationStructureDescriptor interface {
	IMTL4AccelerationStructureDescriptor

	// Topic: Instance Properties

	// Associates the array of geometry descriptors that comprise this primitive acceleration structure.
	GeometryDescriptors() []MTL4AccelerationStructureGeometryDescriptor
	SetGeometryDescriptors(value []MTL4AccelerationStructureGeometryDescriptor)
	// Configures the motion border mode.
	MotionEndBorderMode() MTLMotionBorderMode
	SetMotionEndBorderMode(value MTLMotionBorderMode)
	// Configures the motion end time for this geometry.
	MotionEndTime() float32
	SetMotionEndTime(value float32)
	// Sets the motion keyframe count.
	MotionKeyframeCount() uint
	SetMotionKeyframeCount(value uint)
	// Configures the behavior when the ray-tracing system samples the acceleration structure before the motion start time.
	MotionStartBorderMode() MTLMotionBorderMode
	SetMotionStartBorderMode(value MTLMotionBorderMode)
	// Configures the motion start time for this geometry.
	MotionStartTime() float32
	SetMotionStartTime(value float32)
}

// Init initializes the instance.
func (m MTL4PrimitiveAccelerationStructureDescriptor) Init() MTL4PrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[MTL4PrimitiveAccelerationStructureDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4PrimitiveAccelerationStructureDescriptor) Autorelease() MTL4PrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[MTL4PrimitiveAccelerationStructureDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4PrimitiveAccelerationStructureDescriptor creates a new MTL4PrimitiveAccelerationStructureDescriptor instance.
func NewMTL4PrimitiveAccelerationStructureDescriptor() MTL4PrimitiveAccelerationStructureDescriptor {
	class := getMTL4PrimitiveAccelerationStructureDescriptorClass()
	rv := objc.Send[MTL4PrimitiveAccelerationStructureDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Associates the array of geometry descriptors that comprise this primitive
// acceleration structure.
//
// # Discussion
// 
// If you enable keyframe motion by setting property [MotionKeyframeCount] to
// a value greater than `1`, then all geometry descriptors this array
// references need to be motion geometry descriptors and have a number of
// primitive buffers equals to [MotionKeyframeCount].
// 
// Example of motion geometry descriptors include:
// [MTL4AccelerationStructureMotionTriangleGeometryDescriptor],
// [MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor],
// [MTL4AccelerationStructureMotionCurveGeometryDescriptor].
//
// See: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor/geometryDescriptors
func (m MTL4PrimitiveAccelerationStructureDescriptor) GeometryDescriptors() []MTL4AccelerationStructureGeometryDescriptor {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("geometryDescriptors"))
	return objc.ConvertSlice(rv, func(id objc.ID) MTL4AccelerationStructureGeometryDescriptor {
		return MTL4AccelerationStructureGeometryDescriptorFromID(id)
	})
}
func (m MTL4PrimitiveAccelerationStructureDescriptor) SetGeometryDescriptors(value []MTL4AccelerationStructureGeometryDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setGeometryDescriptors:"), objectivec.IObjectSliceToNSArray(value))
}
// Configures the motion border mode.
//
// # Discussion
// 
// This property controls what happens if Metal samples the acceleration
// structure after [MotionEndTime].
// 
// Its default value is [MTLMotionBorderModeClamp].
//
// See: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor/motionEndBorderMode
func (m MTL4PrimitiveAccelerationStructureDescriptor) MotionEndBorderMode() MTLMotionBorderMode {
	rv := objc.Send[MTLMotionBorderMode](m.ID, objc.Sel("motionEndBorderMode"))
	return MTLMotionBorderMode(rv)
}
func (m MTL4PrimitiveAccelerationStructureDescriptor) SetMotionEndBorderMode(value MTLMotionBorderMode) {
	objc.Send[struct{}](m.ID, objc.Sel("setMotionEndBorderMode:"), value)
}
// Configures the motion end time for this geometry.
//
// # Discussion
// 
// The default value of this property is `1.0f`.
//
// See: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor/motionEndTime
func (m MTL4PrimitiveAccelerationStructureDescriptor) MotionEndTime() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("motionEndTime"))
	return rv
}
func (m MTL4PrimitiveAccelerationStructureDescriptor) SetMotionEndTime(value float32) {
	objc.Send[struct{}](m.ID, objc.Sel("setMotionEndTime:"), value)
}
// Sets the motion keyframe count.
//
// # Discussion
// 
// This property’s default is `1`, indicating no motion.
//
// See: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor/motionKeyframeCount
func (m MTL4PrimitiveAccelerationStructureDescriptor) MotionKeyframeCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("motionKeyframeCount"))
	return rv
}
func (m MTL4PrimitiveAccelerationStructureDescriptor) SetMotionKeyframeCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMotionKeyframeCount:"), value)
}
// Configures the behavior when the ray-tracing system samples the
// acceleration structure before the motion start time.
//
// # Discussion
// 
// Use this property to control the behavior when the ray-tracing system
// samples the acceleration structure at a time prior to the one you set for
// [MotionStartTime].
// 
// The default value of this property is [MTLMotionBorderModeClamp].
//
// See: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor/motionStartBorderMode
func (m MTL4PrimitiveAccelerationStructureDescriptor) MotionStartBorderMode() MTLMotionBorderMode {
	rv := objc.Send[MTLMotionBorderMode](m.ID, objc.Sel("motionStartBorderMode"))
	return MTLMotionBorderMode(rv)
}
func (m MTL4PrimitiveAccelerationStructureDescriptor) SetMotionStartBorderMode(value MTLMotionBorderMode) {
	objc.Send[struct{}](m.ID, objc.Sel("setMotionStartBorderMode:"), value)
}
// Configures the motion start time for this geometry.
//
// # Discussion
// 
// The default value of this property is `0.0f`.
//
// See: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor/motionStartTime
func (m MTL4PrimitiveAccelerationStructureDescriptor) MotionStartTime() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("motionStartTime"))
	return rv
}
func (m MTL4PrimitiveAccelerationStructureDescriptor) SetMotionStartTime(value float32) {
	objc.Send[struct{}](m.ID, objc.Sel("setMotionStartTime:"), value)
}

