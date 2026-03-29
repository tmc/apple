// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MTL4IndirectInstanceAccelerationStructureDescriptor] class.
var (
	_MTL4IndirectInstanceAccelerationStructureDescriptorClass     MTL4IndirectInstanceAccelerationStructureDescriptorClass
	_MTL4IndirectInstanceAccelerationStructureDescriptorClassOnce sync.Once
)

func getMTL4IndirectInstanceAccelerationStructureDescriptorClass() MTL4IndirectInstanceAccelerationStructureDescriptorClass {
	_MTL4IndirectInstanceAccelerationStructureDescriptorClassOnce.Do(func() {
		_MTL4IndirectInstanceAccelerationStructureDescriptorClass = MTL4IndirectInstanceAccelerationStructureDescriptorClass{class: objc.GetClass("MTL4IndirectInstanceAccelerationStructureDescriptor")}
	})
	return _MTL4IndirectInstanceAccelerationStructureDescriptorClass
}

// GetMTL4IndirectInstanceAccelerationStructureDescriptorClass returns the class object for MTL4IndirectInstanceAccelerationStructureDescriptor.
func GetMTL4IndirectInstanceAccelerationStructureDescriptorClass() MTL4IndirectInstanceAccelerationStructureDescriptorClass {
	return getMTL4IndirectInstanceAccelerationStructureDescriptorClass()
}

type MTL4IndirectInstanceAccelerationStructureDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTL4IndirectInstanceAccelerationStructureDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4IndirectInstanceAccelerationStructureDescriptorClass) Alloc() MTL4IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4IndirectInstanceAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Descriptor for an “indirect” instance acceleration structure that
// allows providing the instance count and motion transform count indirectly,
// through buffer references.
//
// # Overview
// 
// An instance acceleration structure references other acceleration
// structures, and provides the ability to “instantiate” them multiple
// times, each one with potentially a different transformation matrix.
// 
// You specify the properties of the instances in the acceleration structure
// this descriptor builds by providing a buffer of `structs` via its
// [MTL4IndirectInstanceAccelerationStructureDescriptor.InstanceDescriptorBuffer] property.
// 
// Compared to [MTL4InstanceAccelerationStructureDescriptor], this descriptor
// allows you to provide the number of instances it references indirectly
// through a buffer reference, as well as the number of motion transforms.
// 
// This enables you to determine these counts indirectly in the GPU timeline
// via a compute pipeline. Metal needs only to know the maximum possible
// number of instances and motion transforms to support, which you specify via
// the [MTL4IndirectInstanceAccelerationStructureDescriptor.MaxInstanceCount] and [MTL4IndirectInstanceAccelerationStructureDescriptor.MaxMotionTransformCount] properties.
// 
// Use a [MTLResidencySet] to mark residency of all buffers and acceleration
// structures this descriptor references when you build this acceleration
// structure.
//
// # Instance Properties
//
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.InstanceCountBuffer]: Provides a reference to a buffer containing the number of instances in the instance descriptor buffer, formatted as a 32-bit unsigned integer.
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.SetInstanceCountBuffer]
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.InstanceDescriptorBuffer]: Assigns a reference to a buffer containing instance descriptors for acceleration structures to reference.
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.SetInstanceDescriptorBuffer]
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.InstanceDescriptorStride]: Sets the stride, in bytes, between instance descriptors in the instance descriptor buffer.
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.SetInstanceDescriptorStride]
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.InstanceDescriptorType]: Controls the type of instance descriptor that the instance descriptor buffer references.
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.SetInstanceDescriptorType]
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.InstanceTransformationMatrixLayout]: Specifies the layout for the transformation matrices in the instance descriptor buffer and the motion transformation matrix buffer.
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.SetInstanceTransformationMatrixLayout]
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.MaxInstanceCount]: Controls the maximum number of instance descriptors the instance descriptor buffer can reference.
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.SetMaxInstanceCount]
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.MaxMotionTransformCount]: Controls the maximum number of motion transforms in the motion transform buffer.
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.SetMaxMotionTransformCount]
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.MotionTransformBuffer]: A buffer containing transformation information for instance motion keyframes, formatted according to the motion transform type.
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.SetMotionTransformBuffer]
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.MotionTransformCountBuffer]: Associates a buffer reference containing the number of motion transforms in the motion transform buffer, formatted as a 32-bit unsigned integer.
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.SetMotionTransformCountBuffer]
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.MotionTransformStride]: Sets the stride for motion transform.
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.SetMotionTransformStride]
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.MotionTransformType]: Sets the type of motion transforms, either as a matrix or individual components.
//   - [MTL4IndirectInstanceAccelerationStructureDescriptor.SetMotionTransformType]
//
// See: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor
type MTL4IndirectInstanceAccelerationStructureDescriptor struct {
	MTL4AccelerationStructureDescriptor
}

// MTL4IndirectInstanceAccelerationStructureDescriptorFromID constructs a [MTL4IndirectInstanceAccelerationStructureDescriptor] from an objc.ID.
//
// Descriptor for an “indirect” instance acceleration structure that
// allows providing the instance count and motion transform count indirectly,
// through buffer references.
func MTL4IndirectInstanceAccelerationStructureDescriptorFromID(id objc.ID) MTL4IndirectInstanceAccelerationStructureDescriptor {
	return MTL4IndirectInstanceAccelerationStructureDescriptor{MTL4AccelerationStructureDescriptor: MTL4AccelerationStructureDescriptorFromID(id)}
}
// NOTE: MTL4IndirectInstanceAccelerationStructureDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4IndirectInstanceAccelerationStructureDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.InstanceCountBuffer]: Provides a reference to a buffer containing the number of instances in the instance descriptor buffer, formatted as a 32-bit unsigned integer.
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.SetInstanceCountBuffer]
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.InstanceDescriptorBuffer]: Assigns a reference to a buffer containing instance descriptors for acceleration structures to reference.
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.SetInstanceDescriptorBuffer]
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.InstanceDescriptorStride]: Sets the stride, in bytes, between instance descriptors in the instance descriptor buffer.
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.SetInstanceDescriptorStride]
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.InstanceDescriptorType]: Controls the type of instance descriptor that the instance descriptor buffer references.
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.SetInstanceDescriptorType]
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.InstanceTransformationMatrixLayout]: Specifies the layout for the transformation matrices in the instance descriptor buffer and the motion transformation matrix buffer.
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.SetInstanceTransformationMatrixLayout]
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.MaxInstanceCount]: Controls the maximum number of instance descriptors the instance descriptor buffer can reference.
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.SetMaxInstanceCount]
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.MaxMotionTransformCount]: Controls the maximum number of motion transforms in the motion transform buffer.
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.SetMaxMotionTransformCount]
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.MotionTransformBuffer]: A buffer containing transformation information for instance motion keyframes, formatted according to the motion transform type.
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.SetMotionTransformBuffer]
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.MotionTransformCountBuffer]: Associates a buffer reference containing the number of motion transforms in the motion transform buffer, formatted as a 32-bit unsigned integer.
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.SetMotionTransformCountBuffer]
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.MotionTransformStride]: Sets the stride for motion transform.
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.SetMotionTransformStride]
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.MotionTransformType]: Sets the type of motion transforms, either as a matrix or individual components.
//   - [IMTL4IndirectInstanceAccelerationStructureDescriptor.SetMotionTransformType]
//
// See: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor
type IMTL4IndirectInstanceAccelerationStructureDescriptor interface {
	IMTL4AccelerationStructureDescriptor

	// Topic: Instance Properties

	// Provides a reference to a buffer containing the number of instances in the instance descriptor buffer, formatted as a 32-bit unsigned integer.
	InstanceCountBuffer() MTL4BufferRange
	SetInstanceCountBuffer(value MTL4BufferRange)
	// Assigns a reference to a buffer containing instance descriptors for acceleration structures to reference.
	InstanceDescriptorBuffer() MTL4BufferRange
	SetInstanceDescriptorBuffer(value MTL4BufferRange)
	// Sets the stride, in bytes, between instance descriptors in the instance descriptor buffer.
	InstanceDescriptorStride() uint
	SetInstanceDescriptorStride(value uint)
	// Controls the type of instance descriptor that the instance descriptor buffer references.
	InstanceDescriptorType() MTLAccelerationStructureInstanceDescriptorType
	SetInstanceDescriptorType(value MTLAccelerationStructureInstanceDescriptorType)
	// Specifies the layout for the transformation matrices in the instance descriptor buffer and the motion transformation matrix buffer.
	InstanceTransformationMatrixLayout() MTLMatrixLayout
	SetInstanceTransformationMatrixLayout(value MTLMatrixLayout)
	// Controls the maximum number of instance descriptors the instance descriptor buffer can reference.
	MaxInstanceCount() uint
	SetMaxInstanceCount(value uint)
	// Controls the maximum number of motion transforms in the motion transform buffer.
	MaxMotionTransformCount() uint
	SetMaxMotionTransformCount(value uint)
	// A buffer containing transformation information for instance motion keyframes, formatted according to the motion transform type.
	MotionTransformBuffer() MTL4BufferRange
	SetMotionTransformBuffer(value MTL4BufferRange)
	// Associates a buffer reference containing the number of motion transforms in the motion transform buffer, formatted as a 32-bit unsigned integer.
	MotionTransformCountBuffer() MTL4BufferRange
	SetMotionTransformCountBuffer(value MTL4BufferRange)
	// Sets the stride for motion transform.
	MotionTransformStride() uint
	SetMotionTransformStride(value uint)
	// Sets the type of motion transforms, either as a matrix or individual components.
	MotionTransformType() MTLTransformType
	SetMotionTransformType(value MTLTransformType)
}

// Init initializes the instance.
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) Init() MTL4IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4IndirectInstanceAccelerationStructureDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) Autorelease() MTL4IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4IndirectInstanceAccelerationStructureDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4IndirectInstanceAccelerationStructureDescriptor creates a new MTL4IndirectInstanceAccelerationStructureDescriptor instance.
func NewMTL4IndirectInstanceAccelerationStructureDescriptor() MTL4IndirectInstanceAccelerationStructureDescriptor {
	class := getMTL4IndirectInstanceAccelerationStructureDescriptorClass()
	rv := objc.Send[MTL4IndirectInstanceAccelerationStructureDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Provides a reference to a buffer containing the number of instances in the
// instance descriptor buffer, formatted as a 32-bit unsigned integer.
//
// # Discussion
// 
// You are responsible for ensuring that the final number of instances at
// build time, which you provide indirectly via this buffer reference , is
// less than or equal to the value of property [MaxInstanceCount].
//
// See: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceCountBuffer
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) InstanceCountBuffer() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("instanceCountBuffer"))
	return MTL4BufferRange(rv)
}
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) SetInstanceCountBuffer(value MTL4BufferRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setInstanceCountBuffer:"), value)
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
// See: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceDescriptorBuffer
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorBuffer() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("instanceDescriptorBuffer"))
	return MTL4BufferRange(rv)
}
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorBuffer(value MTL4BufferRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setInstanceDescriptorBuffer:"), value)
}
// Sets the stride, in bytes, between instance descriptors in the instance
// descriptor buffer.
//
// # Discussion
// 
// You are responsible for ensuring this stride is at least the size of the
// structure type corresponding to the instance descriptor type and a multiple
// of 4 bytes.
// 
// Defaults to `0`, indicating the instance descriptors are tightly packed.
//
// See: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceDescriptorStride
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorStride() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("instanceDescriptorStride"))
	return rv
}
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorStride(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setInstanceDescriptorStride:"), value)
}
// Controls the type of instance descriptor that the instance descriptor
// buffer references.
//
// # Discussion
// 
// This value determines the layout Metal expects for the structs the instance
// descriptor buffer contains.
// 
// Defaults to [MTLAccelerationStructureInstanceDescriptorTypeIndirect]. Valid
// values for this property are
// [MTLAccelerationStructureInstanceDescriptorTypeIndirect] or
// [MTLAccelerationStructureInstanceDescriptorTypeIndirectMotion].
//
// See: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceDescriptorType
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorType() MTLAccelerationStructureInstanceDescriptorType {
	rv := objc.Send[MTLAccelerationStructureInstanceDescriptorType](m.ID, objc.Sel("instanceDescriptorType"))
	return MTLAccelerationStructureInstanceDescriptorType(rv)
}
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorType(value MTLAccelerationStructureInstanceDescriptorType) {
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
// See: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceTransformationMatrixLayout
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) InstanceTransformationMatrixLayout() MTLMatrixLayout {
	rv := objc.Send[MTLMatrixLayout](m.ID, objc.Sel("instanceTransformationMatrixLayout"))
	return MTLMatrixLayout(rv)
}
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) SetInstanceTransformationMatrixLayout(value MTLMatrixLayout) {
	objc.Send[struct{}](m.ID, objc.Sel("setInstanceTransformationMatrixLayout:"), value)
}
// Controls the maximum number of instance descriptors the instance descriptor
// buffer can reference.
//
// # Discussion
// 
// You are responsible for ensuring that the final number of instances at
// build time, which you provide indirectly via a buffer reference in
// [InstanceCountBuffer], is less than or equal to this number.
//
// See: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/maxInstanceCount
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) MaxInstanceCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("maxInstanceCount"))
	return rv
}
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) SetMaxInstanceCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaxInstanceCount:"), value)
}
// Controls the maximum number of motion transforms in the motion transform
// buffer.
//
// # Discussion
// 
// You are responsible for ensuring that final number of motion transforms at
// build time that the buffer [MotionTransformCountBuffer] references is less
// than or equal to this number.
//
// See: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/maxMotionTransformCount
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) MaxMotionTransformCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("maxMotionTransformCount"))
	return rv
}
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) SetMaxMotionTransformCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaxMotionTransformCount:"), value)
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
// See: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/motionTransformBuffer
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) MotionTransformBuffer() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("motionTransformBuffer"))
	return MTL4BufferRange(rv)
}
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformBuffer(value MTL4BufferRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setMotionTransformBuffer:"), value)
}
// Associates a buffer reference containing the number of motion transforms in
// the motion transform buffer, formatted as a 32-bit unsigned integer.
//
// # Discussion
// 
// You are responsible for ensuring that the final number of motion transforms
// at build time in the buffer this property references is less than or equal
// to the value of property [MaxMotionTransformCount].
//
// See: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/motionTransformCountBuffer
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) MotionTransformCountBuffer() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("motionTransformCountBuffer"))
	return MTL4BufferRange(rv)
}
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformCountBuffer(value MTL4BufferRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setMotionTransformCountBuffer:"), value)
}
// Sets the stride for motion transform.
//
// # Discussion
// 
// Defaults to `0`, indicating that transforms are tightly packed according to
// the motion transform type.
//
// See: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/motionTransformStride
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) MotionTransformStride() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("motionTransformStride"))
	return rv
}
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformStride(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMotionTransformStride:"), value)
}
// Sets the type of motion transforms, either as a matrix or individual
// components.
//
// # Discussion
// 
// Defaults to [MTLTransformTypePackedFloat4x3]. Using a
// [MTLTransformTypeComponent] allows you to represent the rotation by a
// quaternion (instead as of part of the matrix), allowing for correct motion
// interpolation.
//
// See: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/motionTransformType
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) MotionTransformType() MTLTransformType {
	rv := objc.Send[MTLTransformType](m.ID, objc.Sel("motionTransformType"))
	return MTLTransformType(rv)
}
func (m MTL4IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformType(value MTLTransformType) {
	objc.Send[struct{}](m.ID, objc.Sel("setMotionTransformType:"), value)
}

