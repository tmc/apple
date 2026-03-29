// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MTL4InstanceAccelerationStructureDescriptor] class.
var (
	_MTL4InstanceAccelerationStructureDescriptorClass     MTL4InstanceAccelerationStructureDescriptorClass
	_MTL4InstanceAccelerationStructureDescriptorClassOnce sync.Once
)

func getMTL4InstanceAccelerationStructureDescriptorClass() MTL4InstanceAccelerationStructureDescriptorClass {
	_MTL4InstanceAccelerationStructureDescriptorClassOnce.Do(func() {
		_MTL4InstanceAccelerationStructureDescriptorClass = MTL4InstanceAccelerationStructureDescriptorClass{class: objc.GetClass("MTL4InstanceAccelerationStructureDescriptor")}
	})
	return _MTL4InstanceAccelerationStructureDescriptorClass
}

// GetMTL4InstanceAccelerationStructureDescriptorClass returns the class object for MTL4InstanceAccelerationStructureDescriptor.
func GetMTL4InstanceAccelerationStructureDescriptorClass() MTL4InstanceAccelerationStructureDescriptorClass {
	return getMTL4InstanceAccelerationStructureDescriptorClass()
}

type MTL4InstanceAccelerationStructureDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTL4InstanceAccelerationStructureDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4InstanceAccelerationStructureDescriptorClass) Alloc() MTL4InstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4InstanceAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Descriptor for an instance acceleration structure.
//
// # Overview
// 
// An instance acceleration structure references other acceleration
// structures, and provides the ability to “instantiate” them multiple
// times, each one with potentially a different transformation matrix.
// 
// You specify the properties of the instances in the acceleration structure
// this descriptor builds by providing a buffer of `structs` via its
// [MTL4InstanceAccelerationStructureDescriptor.InstanceDescriptorBuffer] property.
// 
// Use a [MTLResidencySet] to mark residency of all buffers and acceleration
// structures this descriptor references when you build this acceleration
// structure.
//
// # Instance Properties
//
//   - [MTL4InstanceAccelerationStructureDescriptor.InstanceCount]: Controls the number of instance descriptors in the instance descriptor buffer references.
//   - [MTL4InstanceAccelerationStructureDescriptor.SetInstanceCount]
//   - [MTL4InstanceAccelerationStructureDescriptor.InstanceDescriptorBuffer]: Assigns a reference to a buffer containing instance descriptors for acceleration structures to reference.
//   - [MTL4InstanceAccelerationStructureDescriptor.SetInstanceDescriptorBuffer]
//   - [MTL4InstanceAccelerationStructureDescriptor.InstanceDescriptorStride]: Sets the stride, in bytes, between instance descriptors the instance descriptor buffer references.
//   - [MTL4InstanceAccelerationStructureDescriptor.SetInstanceDescriptorStride]
//   - [MTL4InstanceAccelerationStructureDescriptor.InstanceDescriptorType]: The type of instance descriptor that the instance descriptor buffer references.
//   - [MTL4InstanceAccelerationStructureDescriptor.SetInstanceDescriptorType]
//   - [MTL4InstanceAccelerationStructureDescriptor.InstanceTransformationMatrixLayout]: Specifies the layout for the transformation matrices in the instance descriptor buffer and the motion transformation matrix buffer.
//   - [MTL4InstanceAccelerationStructureDescriptor.SetInstanceTransformationMatrixLayout]
//   - [MTL4InstanceAccelerationStructureDescriptor.MotionTransformBuffer]: A buffer containing transformation information for instance motion keyframes, formatted according to the motion transform type.
//   - [MTL4InstanceAccelerationStructureDescriptor.SetMotionTransformBuffer]
//   - [MTL4InstanceAccelerationStructureDescriptor.MotionTransformCount]: Controls the total number of motion transforms in the motion transform buffer.
//   - [MTL4InstanceAccelerationStructureDescriptor.SetMotionTransformCount]
//   - [MTL4InstanceAccelerationStructureDescriptor.MotionTransformStride]: Specify the stride for motion transform.
//   - [MTL4InstanceAccelerationStructureDescriptor.SetMotionTransformStride]
//   - [MTL4InstanceAccelerationStructureDescriptor.MotionTransformType]: Controls the type of motion transforms, either as a matrix or individual components.
//   - [MTL4InstanceAccelerationStructureDescriptor.SetMotionTransformType]
//
// See: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor
type MTL4InstanceAccelerationStructureDescriptor struct {
	MTL4AccelerationStructureDescriptor
}

// MTL4InstanceAccelerationStructureDescriptorFromID constructs a [MTL4InstanceAccelerationStructureDescriptor] from an objc.ID.
//
// Descriptor for an instance acceleration structure.
func MTL4InstanceAccelerationStructureDescriptorFromID(id objc.ID) MTL4InstanceAccelerationStructureDescriptor {
	return MTL4InstanceAccelerationStructureDescriptor{MTL4AccelerationStructureDescriptor: MTL4AccelerationStructureDescriptorFromID(id)}
}
// NOTE: MTL4InstanceAccelerationStructureDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4InstanceAccelerationStructureDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4InstanceAccelerationStructureDescriptor.InstanceCount]: Controls the number of instance descriptors in the instance descriptor buffer references.
//   - [IMTL4InstanceAccelerationStructureDescriptor.SetInstanceCount]
//   - [IMTL4InstanceAccelerationStructureDescriptor.InstanceDescriptorBuffer]: Assigns a reference to a buffer containing instance descriptors for acceleration structures to reference.
//   - [IMTL4InstanceAccelerationStructureDescriptor.SetInstanceDescriptorBuffer]
//   - [IMTL4InstanceAccelerationStructureDescriptor.InstanceDescriptorStride]: Sets the stride, in bytes, between instance descriptors the instance descriptor buffer references.
//   - [IMTL4InstanceAccelerationStructureDescriptor.SetInstanceDescriptorStride]
//   - [IMTL4InstanceAccelerationStructureDescriptor.InstanceDescriptorType]: The type of instance descriptor that the instance descriptor buffer references.
//   - [IMTL4InstanceAccelerationStructureDescriptor.SetInstanceDescriptorType]
//   - [IMTL4InstanceAccelerationStructureDescriptor.InstanceTransformationMatrixLayout]: Specifies the layout for the transformation matrices in the instance descriptor buffer and the motion transformation matrix buffer.
//   - [IMTL4InstanceAccelerationStructureDescriptor.SetInstanceTransformationMatrixLayout]
//   - [IMTL4InstanceAccelerationStructureDescriptor.MotionTransformBuffer]: A buffer containing transformation information for instance motion keyframes, formatted according to the motion transform type.
//   - [IMTL4InstanceAccelerationStructureDescriptor.SetMotionTransformBuffer]
//   - [IMTL4InstanceAccelerationStructureDescriptor.MotionTransformCount]: Controls the total number of motion transforms in the motion transform buffer.
//   - [IMTL4InstanceAccelerationStructureDescriptor.SetMotionTransformCount]
//   - [IMTL4InstanceAccelerationStructureDescriptor.MotionTransformStride]: Specify the stride for motion transform.
//   - [IMTL4InstanceAccelerationStructureDescriptor.SetMotionTransformStride]
//   - [IMTL4InstanceAccelerationStructureDescriptor.MotionTransformType]: Controls the type of motion transforms, either as a matrix or individual components.
//   - [IMTL4InstanceAccelerationStructureDescriptor.SetMotionTransformType]
//
// See: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor
type IMTL4InstanceAccelerationStructureDescriptor interface {
	IMTL4AccelerationStructureDescriptor

	// Topic: Instance Properties

	// Controls the number of instance descriptors in the instance descriptor buffer references.
	InstanceCount() uint
	SetInstanceCount(value uint)
	// Assigns a reference to a buffer containing instance descriptors for acceleration structures to reference.
	InstanceDescriptorBuffer() MTL4BufferRange
	SetInstanceDescriptorBuffer(value MTL4BufferRange)
	// Sets the stride, in bytes, between instance descriptors the instance descriptor buffer references.
	InstanceDescriptorStride() uint
	SetInstanceDescriptorStride(value uint)
	// The type of instance descriptor that the instance descriptor buffer references.
	InstanceDescriptorType() MTLAccelerationStructureInstanceDescriptorType
	SetInstanceDescriptorType(value MTLAccelerationStructureInstanceDescriptorType)
	// Specifies the layout for the transformation matrices in the instance descriptor buffer and the motion transformation matrix buffer.
	InstanceTransformationMatrixLayout() MTLMatrixLayout
	SetInstanceTransformationMatrixLayout(value MTLMatrixLayout)
	// A buffer containing transformation information for instance motion keyframes, formatted according to the motion transform type.
	MotionTransformBuffer() MTL4BufferRange
	SetMotionTransformBuffer(value MTL4BufferRange)
	// Controls the total number of motion transforms in the motion transform buffer.
	MotionTransformCount() uint
	SetMotionTransformCount(value uint)
	// Specify the stride for motion transform.
	MotionTransformStride() uint
	SetMotionTransformStride(value uint)
	// Controls the type of motion transforms, either as a matrix or individual components.
	MotionTransformType() MTLTransformType
	SetMotionTransformType(value MTLTransformType)
}

// Init initializes the instance.
func (m MTL4InstanceAccelerationStructureDescriptor) Init() MTL4InstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4InstanceAccelerationStructureDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4InstanceAccelerationStructureDescriptor) Autorelease() MTL4InstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4InstanceAccelerationStructureDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4InstanceAccelerationStructureDescriptor creates a new MTL4InstanceAccelerationStructureDescriptor instance.
func NewMTL4InstanceAccelerationStructureDescriptor() MTL4InstanceAccelerationStructureDescriptor {
	class := getMTL4InstanceAccelerationStructureDescriptorClass()
	rv := objc.Send[MTL4InstanceAccelerationStructureDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Controls the number of instance descriptors in the instance descriptor
// buffer references.
//
// See: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/instanceCount
func (m MTL4InstanceAccelerationStructureDescriptor) InstanceCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("instanceCount"))
	return rv
}
func (m MTL4InstanceAccelerationStructureDescriptor) SetInstanceCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setInstanceCount:"), value)
}
// Assigns a reference to a buffer containing instance descriptors for
// acceleration structures to reference.
//
// # Discussion
// 
// This buffer conceptually represents an array of instance data. The specific
// format for the structs that comprise each entry depends on the value of the
// [InstanceDescriptorType] property.
// 
// You are responsible for ensuring the buffer address the range contains is
// not zero.
//
// See: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/instanceDescriptorBuffer
func (m MTL4InstanceAccelerationStructureDescriptor) InstanceDescriptorBuffer() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("instanceDescriptorBuffer"))
	return MTL4BufferRange(rv)
}
func (m MTL4InstanceAccelerationStructureDescriptor) SetInstanceDescriptorBuffer(value MTL4BufferRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setInstanceDescriptorBuffer:"), value)
}
// Sets the stride, in bytes, between instance descriptors the instance
// descriptor buffer references.
//
// # Discussion
// 
// You are responsible for ensuring this stride is at least the size of the
// structure type corresponding to the instance descriptor type and a multiple
// of 4 bytes.
// 
// Defaults to `0`, indicating the instance descriptors are tightly packed.
//
// See: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/instanceDescriptorStride
func (m MTL4InstanceAccelerationStructureDescriptor) InstanceDescriptorStride() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("instanceDescriptorStride"))
	return rv
}
func (m MTL4InstanceAccelerationStructureDescriptor) SetInstanceDescriptorStride(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setInstanceDescriptorStride:"), value)
}
// The type of instance descriptor that the instance descriptor buffer
// references.
//
// # Discussion
// 
// This value determines the layout Metal expects for the structs the instance
// descriptor buffer contains:
// 
// - [AccelerationStructureInstanceDescriptorTypeIndirect]: Use the
// [MTLIndirectAccelerationStructureInstanceDescriptor] struct layout. -
// [AccelerationStructureInstanceDescriptorTypeIndirectMotion]: Use the
// [MTLIndirectAccelerationStructureMotionInstanceDescriptor] struct layout.
// 
// The default value is [AccelerationStructureInstanceDescriptorTypeIndirect].
//
// [MTLIndirectAccelerationStructureInstanceDescriptor]: https://developer.apple.com/documentation/Metal/MTLIndirectAccelerationStructureInstanceDescriptor
// [MTLIndirectAccelerationStructureMotionInstanceDescriptor]: https://developer.apple.com/documentation/Metal/MTLIndirectAccelerationStructureMotionInstanceDescriptor
//
// See: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/instanceDescriptorType
func (m MTL4InstanceAccelerationStructureDescriptor) InstanceDescriptorType() MTLAccelerationStructureInstanceDescriptorType {
	rv := objc.Send[MTLAccelerationStructureInstanceDescriptorType](m.ID, objc.Sel("instanceDescriptorType"))
	return MTLAccelerationStructureInstanceDescriptorType(rv)
}
func (m MTL4InstanceAccelerationStructureDescriptor) SetInstanceDescriptorType(value MTLAccelerationStructureInstanceDescriptorType) {
	objc.Send[struct{}](m.ID, objc.Sel("setInstanceDescriptorType:"), value)
}
// Specifies the layout for the transformation matrices in the instance
// descriptor buffer and the motion transformation matrix buffer.
//
// # Discussion
// 
// Metal interprets the value of this property as the layout for the buffers
// that both [InstanceDescriptorBuffer] and [MotionTransformBuffer] reference.
// 
// Defaults to [MTLMatrixLayoutColumnMajor].
//
// See: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/instanceTransformationMatrixLayout
func (m MTL4InstanceAccelerationStructureDescriptor) InstanceTransformationMatrixLayout() MTLMatrixLayout {
	rv := objc.Send[MTLMatrixLayout](m.ID, objc.Sel("instanceTransformationMatrixLayout"))
	return MTLMatrixLayout(rv)
}
func (m MTL4InstanceAccelerationStructureDescriptor) SetInstanceTransformationMatrixLayout(value MTLMatrixLayout) {
	objc.Send[struct{}](m.ID, objc.Sel("setInstanceTransformationMatrixLayout:"), value)
}
// A buffer containing transformation information for instance motion
// keyframes, formatted according to the motion transform type.
//
// # Discussion
// 
// Each instance can have a different number of keyframes that you configure
// via individual instance descriptors.
// 
// You are responsible for ensuring the buffer address the range references is
// not zero when using motion instance descriptors.
//
// See: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/motionTransformBuffer
func (m MTL4InstanceAccelerationStructureDescriptor) MotionTransformBuffer() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("motionTransformBuffer"))
	return MTL4BufferRange(rv)
}
func (m MTL4InstanceAccelerationStructureDescriptor) SetMotionTransformBuffer(value MTL4BufferRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setMotionTransformBuffer:"), value)
}
// Controls the total number of motion transforms in the motion transform
// buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/motionTransformCount
func (m MTL4InstanceAccelerationStructureDescriptor) MotionTransformCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("motionTransformCount"))
	return rv
}
func (m MTL4InstanceAccelerationStructureDescriptor) SetMotionTransformCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMotionTransformCount:"), value)
}
// Specify the stride for motion transform.
//
// # Discussion
// 
// Defaults to `0`, indicating that transforms are tightly packed according to
// the motion transform type.
//
// See: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/motionTransformStride
func (m MTL4InstanceAccelerationStructureDescriptor) MotionTransformStride() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("motionTransformStride"))
	return rv
}
func (m MTL4InstanceAccelerationStructureDescriptor) SetMotionTransformStride(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMotionTransformStride:"), value)
}
// Controls the type of motion transforms, either as a matrix or individual
// components.
//
// # Discussion
// 
// Defaults to [MTLTransformTypePackedFloat4x3]. Using a
// [MTLTransformTypeComponent] allows you to represent the rotation by a
// quaternion (instead as of part of the matrix), allowing for correct motion
// interpolation.
//
// See: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/motionTransformType
func (m MTL4InstanceAccelerationStructureDescriptor) MotionTransformType() MTLTransformType {
	rv := objc.Send[MTLTransformType](m.ID, objc.Sel("motionTransformType"))
	return MTLTransformType(rv)
}
func (m MTL4InstanceAccelerationStructureDescriptor) SetMotionTransformType(value MTLTransformType) {
	objc.Send[struct{}](m.ID, objc.Sel("setMotionTransformType:"), value)
}

