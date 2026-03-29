// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLInstanceAccelerationStructureDescriptor] class.
var (
	_MTLInstanceAccelerationStructureDescriptorClass     MTLInstanceAccelerationStructureDescriptorClass
	_MTLInstanceAccelerationStructureDescriptorClassOnce sync.Once
)

func getMTLInstanceAccelerationStructureDescriptorClass() MTLInstanceAccelerationStructureDescriptorClass {
	_MTLInstanceAccelerationStructureDescriptorClassOnce.Do(func() {
		_MTLInstanceAccelerationStructureDescriptorClass = MTLInstanceAccelerationStructureDescriptorClass{class: objc.GetClass("MTLInstanceAccelerationStructureDescriptor")}
	})
	return _MTLInstanceAccelerationStructureDescriptorClass
}

// GetMTLInstanceAccelerationStructureDescriptorClass returns the class object for MTLInstanceAccelerationStructureDescriptor.
func GetMTLInstanceAccelerationStructureDescriptorClass() MTLInstanceAccelerationStructureDescriptorClass {
	return getMTLInstanceAccelerationStructureDescriptorClass()
}

type MTLInstanceAccelerationStructureDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLInstanceAccelerationStructureDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLInstanceAccelerationStructureDescriptorClass) Alloc() MTLInstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTLInstanceAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of an acceleration structure that derives from instances of
// primitive acceleration structures.
//
// # Overview
// 
// Metal provides acceleration structures with a two-level hierarchy. The
// bottom layer consists of primitive acceleration structures, which instance
// acceleration structures in the top level reference.
//
// # Specifying the instance structures
//
//   - [MTLInstanceAccelerationStructureDescriptor.InstanceDescriptorType]: The format of the instance data in the descriptor buffer.
//   - [MTLInstanceAccelerationStructureDescriptor.SetInstanceDescriptorType]
//   - [MTLInstanceAccelerationStructureDescriptor.InstancedAccelerationStructures]: The bottom-level acceleration structures that instances use in the instance acceleration structure .
//   - [MTLInstanceAccelerationStructureDescriptor.SetInstancedAccelerationStructures]
//
// # Specifying the list of instances
//
//   - [MTLInstanceAccelerationStructureDescriptor.InstanceCount]: The number of instances in the instance descriptor buffer.
//   - [MTLInstanceAccelerationStructureDescriptor.SetInstanceCount]
//   - [MTLInstanceAccelerationStructureDescriptor.InstanceDescriptorBuffer]: A buffer that contains descriptions of each instance in the acceleration structure.
//   - [MTLInstanceAccelerationStructureDescriptor.SetInstanceDescriptorBuffer]
//   - [MTLInstanceAccelerationStructureDescriptor.InstanceDescriptorBufferOffset]: The offset, in bytes, to the descripton of the first instance.
//   - [MTLInstanceAccelerationStructureDescriptor.SetInstanceDescriptorBufferOffset]
//   - [MTLInstanceAccelerationStructureDescriptor.InstanceDescriptorStride]: The stride, in bytes, between instance descriptions.
//   - [MTLInstanceAccelerationStructureDescriptor.SetInstanceDescriptorStride]
//
// # Specifying motion data
//
//   - [MTLInstanceAccelerationStructureDescriptor.MotionTransformCount]: The number of motion transforms in the motion transform buffer.
//   - [MTLInstanceAccelerationStructureDescriptor.SetMotionTransformCount]
//   - [MTLInstanceAccelerationStructureDescriptor.MotionTransformBuffer]: A buffer that contains descriptions of each motion transform in the acceleration structure.
//   - [MTLInstanceAccelerationStructureDescriptor.SetMotionTransformBuffer]
//   - [MTLInstanceAccelerationStructureDescriptor.MotionTransformBufferOffset]: The offset, in bytes, to the descripton of the first motion transform.
//   - [MTLInstanceAccelerationStructureDescriptor.SetMotionTransformBufferOffset]
//
// # Instance Properties
//
//   - [MTLInstanceAccelerationStructureDescriptor.InstanceTransformationMatrixLayout]
//   - [MTLInstanceAccelerationStructureDescriptor.SetInstanceTransformationMatrixLayout]
//   - [MTLInstanceAccelerationStructureDescriptor.MotionTransformStride]
//   - [MTLInstanceAccelerationStructureDescriptor.SetMotionTransformStride]
//   - [MTLInstanceAccelerationStructureDescriptor.MotionTransformType]
//   - [MTLInstanceAccelerationStructureDescriptor.SetMotionTransformType]
//
// See: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor
type MTLInstanceAccelerationStructureDescriptor struct {
	MTLAccelerationStructureDescriptor
}

// MTLInstanceAccelerationStructureDescriptorFromID constructs a [MTLInstanceAccelerationStructureDescriptor] from an objc.ID.
//
// A description of an acceleration structure that derives from instances of
// primitive acceleration structures.
func MTLInstanceAccelerationStructureDescriptorFromID(id objc.ID) MTLInstanceAccelerationStructureDescriptor {
	return MTLInstanceAccelerationStructureDescriptor{MTLAccelerationStructureDescriptor: MTLAccelerationStructureDescriptorFromID(id)}
}
// NOTE: MTLInstanceAccelerationStructureDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLInstanceAccelerationStructureDescriptor] class.
//
// # Specifying the instance structures
//
//   - [IMTLInstanceAccelerationStructureDescriptor.InstanceDescriptorType]: The format of the instance data in the descriptor buffer.
//   - [IMTLInstanceAccelerationStructureDescriptor.SetInstanceDescriptorType]
//   - [IMTLInstanceAccelerationStructureDescriptor.InstancedAccelerationStructures]: The bottom-level acceleration structures that instances use in the instance acceleration structure .
//   - [IMTLInstanceAccelerationStructureDescriptor.SetInstancedAccelerationStructures]
//
// # Specifying the list of instances
//
//   - [IMTLInstanceAccelerationStructureDescriptor.InstanceCount]: The number of instances in the instance descriptor buffer.
//   - [IMTLInstanceAccelerationStructureDescriptor.SetInstanceCount]
//   - [IMTLInstanceAccelerationStructureDescriptor.InstanceDescriptorBuffer]: A buffer that contains descriptions of each instance in the acceleration structure.
//   - [IMTLInstanceAccelerationStructureDescriptor.SetInstanceDescriptorBuffer]
//   - [IMTLInstanceAccelerationStructureDescriptor.InstanceDescriptorBufferOffset]: The offset, in bytes, to the descripton of the first instance.
//   - [IMTLInstanceAccelerationStructureDescriptor.SetInstanceDescriptorBufferOffset]
//   - [IMTLInstanceAccelerationStructureDescriptor.InstanceDescriptorStride]: The stride, in bytes, between instance descriptions.
//   - [IMTLInstanceAccelerationStructureDescriptor.SetInstanceDescriptorStride]
//
// # Specifying motion data
//
//   - [IMTLInstanceAccelerationStructureDescriptor.MotionTransformCount]: The number of motion transforms in the motion transform buffer.
//   - [IMTLInstanceAccelerationStructureDescriptor.SetMotionTransformCount]
//   - [IMTLInstanceAccelerationStructureDescriptor.MotionTransformBuffer]: A buffer that contains descriptions of each motion transform in the acceleration structure.
//   - [IMTLInstanceAccelerationStructureDescriptor.SetMotionTransformBuffer]
//   - [IMTLInstanceAccelerationStructureDescriptor.MotionTransformBufferOffset]: The offset, in bytes, to the descripton of the first motion transform.
//   - [IMTLInstanceAccelerationStructureDescriptor.SetMotionTransformBufferOffset]
//
// # Instance Properties
//
//   - [IMTLInstanceAccelerationStructureDescriptor.InstanceTransformationMatrixLayout]
//   - [IMTLInstanceAccelerationStructureDescriptor.SetInstanceTransformationMatrixLayout]
//   - [IMTLInstanceAccelerationStructureDescriptor.MotionTransformStride]
//   - [IMTLInstanceAccelerationStructureDescriptor.SetMotionTransformStride]
//   - [IMTLInstanceAccelerationStructureDescriptor.MotionTransformType]
//   - [IMTLInstanceAccelerationStructureDescriptor.SetMotionTransformType]
//
// See: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor
type IMTLInstanceAccelerationStructureDescriptor interface {
	IMTLAccelerationStructureDescriptor

	// Topic: Specifying the instance structures

	// The format of the instance data in the descriptor buffer.
	InstanceDescriptorType() MTLAccelerationStructureInstanceDescriptorType
	SetInstanceDescriptorType(value MTLAccelerationStructureInstanceDescriptorType)
	// The bottom-level acceleration structures that instances use in the instance acceleration structure .
	InstancedAccelerationStructures() []objectivec.IObject
	SetInstancedAccelerationStructures(value []objectivec.IObject)

	// Topic: Specifying the list of instances

	// The number of instances in the instance descriptor buffer.
	InstanceCount() uint
	SetInstanceCount(value uint)
	// A buffer that contains descriptions of each instance in the acceleration structure.
	InstanceDescriptorBuffer() MTLBuffer
	SetInstanceDescriptorBuffer(value MTLBuffer)
	// The offset, in bytes, to the descripton of the first instance.
	InstanceDescriptorBufferOffset() uint
	SetInstanceDescriptorBufferOffset(value uint)
	// The stride, in bytes, between instance descriptions.
	InstanceDescriptorStride() uint
	SetInstanceDescriptorStride(value uint)

	// Topic: Specifying motion data

	// The number of motion transforms in the motion transform buffer.
	MotionTransformCount() uint
	SetMotionTransformCount(value uint)
	// A buffer that contains descriptions of each motion transform in the acceleration structure.
	MotionTransformBuffer() MTLBuffer
	SetMotionTransformBuffer(value MTLBuffer)
	// The offset, in bytes, to the descripton of the first motion transform.
	MotionTransformBufferOffset() uint
	SetMotionTransformBufferOffset(value uint)

	// Topic: Instance Properties

	InstanceTransformationMatrixLayout() MTLMatrixLayout
	SetInstanceTransformationMatrixLayout(value MTLMatrixLayout)
	MotionTransformStride() uint
	SetMotionTransformStride(value uint)
	MotionTransformType() MTLTransformType
	SetMotionTransformType(value MTLTransformType)
}

// Init initializes the instance.
func (i MTLInstanceAccelerationStructureDescriptor) Init() MTLInstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTLInstanceAccelerationStructureDescriptor](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MTLInstanceAccelerationStructureDescriptor) Autorelease() MTLInstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTLInstanceAccelerationStructureDescriptor](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLInstanceAccelerationStructureDescriptor creates a new MTLInstanceAccelerationStructureDescriptor instance.
func NewMTLInstanceAccelerationStructureDescriptor() MTLInstanceAccelerationStructureDescriptor {
	class := getMTLInstanceAccelerationStructureDescriptorClass()
	rv := objc.Send[MTLInstanceAccelerationStructureDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an instance descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/descriptor
func (_MTLInstanceAccelerationStructureDescriptorClass MTLInstanceAccelerationStructureDescriptorClass) Descriptor() MTLInstanceAccelerationStructureDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MTLInstanceAccelerationStructureDescriptorClass.class), objc.Sel("descriptor"))
	return MTLInstanceAccelerationStructureDescriptorFromID(rv)
}

// The format of the instance data in the descriptor buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instanceDescriptorType
func (i MTLInstanceAccelerationStructureDescriptor) InstanceDescriptorType() MTLAccelerationStructureInstanceDescriptorType {
	rv := objc.Send[MTLAccelerationStructureInstanceDescriptorType](i.ID, objc.Sel("instanceDescriptorType"))
	return MTLAccelerationStructureInstanceDescriptorType(rv)
}
func (i MTLInstanceAccelerationStructureDescriptor) SetInstanceDescriptorType(value MTLAccelerationStructureInstanceDescriptorType) {
	objc.Send[struct{}](i.ID, objc.Sel("setInstanceDescriptorType:"), value)
}
// The bottom-level acceleration structures that instances use in the instance
// acceleration structure .
//
// # Discussion
// 
// Each instance in the instance descriptor buffer has an index into this
// array, specifying which acceleration structure to use for that instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instancedAccelerationStructures
func (i MTLInstanceAccelerationStructureDescriptor) InstancedAccelerationStructures() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("instancedAccelerationStructures"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (i MTLInstanceAccelerationStructureDescriptor) SetInstancedAccelerationStructures(value []objectivec.IObject) {
	objc.Send[struct{}](i.ID, objc.Sel("setInstancedAccelerationStructures:"), objectivec.IObjectSliceToNSArray(value))
}
// The number of instances in the instance descriptor buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instanceCount
func (i MTLInstanceAccelerationStructureDescriptor) InstanceCount() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("instanceCount"))
	return rv
}
func (i MTLInstanceAccelerationStructureDescriptor) SetInstanceCount(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setInstanceCount:"), value)
}
// A buffer that contains descriptions of each instance in the acceleration
// structure.
//
// # Discussion
// 
// You need to set a buffer before creating the instanced acceleration
// structure. The buffer needs to contain a list of instance data structures,
// each defining the characteristics of an instance. The descriptor’s
// [InstanceDescriptorType] property determines which memory layout to use for
// the instance data; see [MTLAccelerationStructureInstanceDescriptorType] for
// more information.
//
// [MTLAccelerationStructureInstanceDescriptorType]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureInstanceDescriptorType
//
// See: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instanceDescriptorBuffer
func (i MTLInstanceAccelerationStructureDescriptor) InstanceDescriptorBuffer() MTLBuffer {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("instanceDescriptorBuffer"))
	return MTLBufferObjectFromID(rv)
}
func (i MTLInstanceAccelerationStructureDescriptor) SetInstanceDescriptorBuffer(value MTLBuffer) {
	objc.Send[struct{}](i.ID, objc.Sel("setInstanceDescriptorBuffer:"), value)
}
// The offset, in bytes, to the descripton of the first instance.
//
// # Discussion
// 
// Specify an offset that is a multiple of 4 bytes and a multiple of the
// platform’s buffer offset alignment.
//
// See: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instanceDescriptorBufferOffset
func (i MTLInstanceAccelerationStructureDescriptor) InstanceDescriptorBufferOffset() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("instanceDescriptorBufferOffset"))
	return rv
}
func (i MTLInstanceAccelerationStructureDescriptor) SetInstanceDescriptorBufferOffset(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setInstanceDescriptorBufferOffset:"), value)
}
// The stride, in bytes, between instance descriptions.
//
// # Discussion
// 
// The stride needs to be at least 64 bytes and needs to be a multiple of 4
// bytes. Defaults to 64 bytes.
//
// See: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instanceDescriptorStride
func (i MTLInstanceAccelerationStructureDescriptor) InstanceDescriptorStride() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("instanceDescriptorStride"))
	return rv
}
func (i MTLInstanceAccelerationStructureDescriptor) SetInstanceDescriptorStride(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setInstanceDescriptorStride:"), value)
}
// The number of motion transforms in the motion transform buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/motionTransformCount
func (i MTLInstanceAccelerationStructureDescriptor) MotionTransformCount() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("motionTransformCount"))
	return rv
}
func (i MTLInstanceAccelerationStructureDescriptor) SetMotionTransformCount(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setMotionTransformCount:"), value)
}
// A buffer that contains descriptions of each motion transform in the
// acceleration structure.
//
// See: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/motionTransformBuffer
func (i MTLInstanceAccelerationStructureDescriptor) MotionTransformBuffer() MTLBuffer {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("motionTransformBuffer"))
	return MTLBufferObjectFromID(rv)
}
func (i MTLInstanceAccelerationStructureDescriptor) SetMotionTransformBuffer(value MTLBuffer) {
	objc.Send[struct{}](i.ID, objc.Sel("setMotionTransformBuffer:"), value)
}
// The offset, in bytes, to the descripton of the first motion transform.
//
// See: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/motionTransformBufferOffset
func (i MTLInstanceAccelerationStructureDescriptor) MotionTransformBufferOffset() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("motionTransformBufferOffset"))
	return rv
}
func (i MTLInstanceAccelerationStructureDescriptor) SetMotionTransformBufferOffset(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setMotionTransformBufferOffset:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instanceTransformationMatrixLayout
func (i MTLInstanceAccelerationStructureDescriptor) InstanceTransformationMatrixLayout() MTLMatrixLayout {
	rv := objc.Send[MTLMatrixLayout](i.ID, objc.Sel("instanceTransformationMatrixLayout"))
	return MTLMatrixLayout(rv)
}
func (i MTLInstanceAccelerationStructureDescriptor) SetInstanceTransformationMatrixLayout(value MTLMatrixLayout) {
	objc.Send[struct{}](i.ID, objc.Sel("setInstanceTransformationMatrixLayout:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/motionTransformStride
func (i MTLInstanceAccelerationStructureDescriptor) MotionTransformStride() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("motionTransformStride"))
	return rv
}
func (i MTLInstanceAccelerationStructureDescriptor) SetMotionTransformStride(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setMotionTransformStride:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/motionTransformType
func (i MTLInstanceAccelerationStructureDescriptor) MotionTransformType() MTLTransformType {
	rv := objc.Send[MTLTransformType](i.ID, objc.Sel("motionTransformType"))
	return MTLTransformType(rv)
}
func (i MTLInstanceAccelerationStructureDescriptor) SetMotionTransformType(value MTLTransformType) {
	objc.Send[struct{}](i.ID, objc.Sel("setMotionTransformType:"), value)
}

