// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MTLIndirectInstanceAccelerationStructureDescriptor] class.
var (
	_MTLIndirectInstanceAccelerationStructureDescriptorClass     MTLIndirectInstanceAccelerationStructureDescriptorClass
	_MTLIndirectInstanceAccelerationStructureDescriptorClassOnce sync.Once
)

func getMTLIndirectInstanceAccelerationStructureDescriptorClass() MTLIndirectInstanceAccelerationStructureDescriptorClass {
	_MTLIndirectInstanceAccelerationStructureDescriptorClassOnce.Do(func() {
		_MTLIndirectInstanceAccelerationStructureDescriptorClass = MTLIndirectInstanceAccelerationStructureDescriptorClass{class: objc.GetClass("MTLIndirectInstanceAccelerationStructureDescriptor")}
	})
	return _MTLIndirectInstanceAccelerationStructureDescriptorClass
}

// GetMTLIndirectInstanceAccelerationStructureDescriptorClass returns the class object for MTLIndirectInstanceAccelerationStructureDescriptor.
func GetMTLIndirectInstanceAccelerationStructureDescriptorClass() MTLIndirectInstanceAccelerationStructureDescriptorClass {
	return getMTLIndirectInstanceAccelerationStructureDescriptorClass()
}

type MTLIndirectInstanceAccelerationStructureDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLIndirectInstanceAccelerationStructureDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLIndirectInstanceAccelerationStructureDescriptorClass) Alloc() MTLIndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTLIndirectInstanceAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of an acceleration structure that Metal derives from
// instances of primitive acceleration structures that the GPU can populate.
//
// # Instance Properties
//
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.InstanceCountBuffer]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.SetInstanceCountBuffer]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.InstanceCountBufferOffset]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.SetInstanceCountBufferOffset]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.InstanceDescriptorBuffer]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.SetInstanceDescriptorBuffer]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.InstanceDescriptorBufferOffset]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.SetInstanceDescriptorBufferOffset]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.InstanceDescriptorStride]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.SetInstanceDescriptorStride]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.InstanceDescriptorType]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.SetInstanceDescriptorType]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.InstanceTransformationMatrixLayout]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.SetInstanceTransformationMatrixLayout]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.MaxInstanceCount]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.SetMaxInstanceCount]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.MaxMotionTransformCount]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.SetMaxMotionTransformCount]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.MotionTransformBuffer]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.SetMotionTransformBuffer]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.MotionTransformBufferOffset]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.SetMotionTransformBufferOffset]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.MotionTransformCountBuffer]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.SetMotionTransformCountBuffer]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.MotionTransformCountBufferOffset]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.SetMotionTransformCountBufferOffset]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.MotionTransformStride]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.SetMotionTransformStride]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.MotionTransformType]
//   - [MTLIndirectInstanceAccelerationStructureDescriptor.SetMotionTransformType]
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor
type MTLIndirectInstanceAccelerationStructureDescriptor struct {
	MTLAccelerationStructureDescriptor
}

// MTLIndirectInstanceAccelerationStructureDescriptorFromID constructs a [MTLIndirectInstanceAccelerationStructureDescriptor] from an objc.ID.
//
// A description of an acceleration structure that Metal derives from
// instances of primitive acceleration structures that the GPU can populate.
func MTLIndirectInstanceAccelerationStructureDescriptorFromID(id objc.ID) MTLIndirectInstanceAccelerationStructureDescriptor {
	return MTLIndirectInstanceAccelerationStructureDescriptor{MTLAccelerationStructureDescriptor: MTLAccelerationStructureDescriptorFromID(id)}
}

// NOTE: MTLIndirectInstanceAccelerationStructureDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLIndirectInstanceAccelerationStructureDescriptor] class.
//
// # Instance Properties
//
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.InstanceCountBuffer]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.SetInstanceCountBuffer]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.InstanceCountBufferOffset]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.SetInstanceCountBufferOffset]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.InstanceDescriptorBuffer]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.SetInstanceDescriptorBuffer]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.InstanceDescriptorBufferOffset]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.SetInstanceDescriptorBufferOffset]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.InstanceDescriptorStride]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.SetInstanceDescriptorStride]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.InstanceDescriptorType]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.SetInstanceDescriptorType]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.InstanceTransformationMatrixLayout]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.SetInstanceTransformationMatrixLayout]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.MaxInstanceCount]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.SetMaxInstanceCount]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.MaxMotionTransformCount]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.SetMaxMotionTransformCount]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.MotionTransformBuffer]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.SetMotionTransformBuffer]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.MotionTransformBufferOffset]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.SetMotionTransformBufferOffset]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.MotionTransformCountBuffer]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.SetMotionTransformCountBuffer]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.MotionTransformCountBufferOffset]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.SetMotionTransformCountBufferOffset]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.MotionTransformStride]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.SetMotionTransformStride]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.MotionTransformType]
//   - [IMTLIndirectInstanceAccelerationStructureDescriptor.SetMotionTransformType]
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor
type IMTLIndirectInstanceAccelerationStructureDescriptor interface {
	IMTLAccelerationStructureDescriptor

	// Topic: Instance Properties

	InstanceCountBuffer() MTLBuffer
	SetInstanceCountBuffer(value MTLBuffer)
	InstanceCountBufferOffset() uint
	SetInstanceCountBufferOffset(value uint)
	InstanceDescriptorBuffer() MTLBuffer
	SetInstanceDescriptorBuffer(value MTLBuffer)
	InstanceDescriptorBufferOffset() uint
	SetInstanceDescriptorBufferOffset(value uint)
	InstanceDescriptorStride() uint
	SetInstanceDescriptorStride(value uint)
	InstanceDescriptorType() MTLAccelerationStructureInstanceDescriptorType
	SetInstanceDescriptorType(value MTLAccelerationStructureInstanceDescriptorType)
	InstanceTransformationMatrixLayout() MTLMatrixLayout
	SetInstanceTransformationMatrixLayout(value MTLMatrixLayout)
	MaxInstanceCount() uint
	SetMaxInstanceCount(value uint)
	MaxMotionTransformCount() uint
	SetMaxMotionTransformCount(value uint)
	MotionTransformBuffer() MTLBuffer
	SetMotionTransformBuffer(value MTLBuffer)
	MotionTransformBufferOffset() uint
	SetMotionTransformBufferOffset(value uint)
	MotionTransformCountBuffer() MTLBuffer
	SetMotionTransformCountBuffer(value MTLBuffer)
	MotionTransformCountBufferOffset() uint
	SetMotionTransformCountBufferOffset(value uint)
	MotionTransformStride() uint
	SetMotionTransformStride(value uint)
	MotionTransformType() MTLTransformType
	SetMotionTransformType(value MTLTransformType)
}

// Init initializes the instance.
func (i MTLIndirectInstanceAccelerationStructureDescriptor) Init() MTLIndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTLIndirectInstanceAccelerationStructureDescriptor](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MTLIndirectInstanceAccelerationStructureDescriptor) Autorelease() MTLIndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTLIndirectInstanceAccelerationStructureDescriptor](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLIndirectInstanceAccelerationStructureDescriptor creates a new MTLIndirectInstanceAccelerationStructureDescriptor instance.
func NewMTLIndirectInstanceAccelerationStructureDescriptor() MTLIndirectInstanceAccelerationStructureDescriptor {
	class := getMTLIndirectInstanceAccelerationStructureDescriptorClass()
	rv := objc.Send[MTLIndirectInstanceAccelerationStructureDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceCountBuffer
func (i MTLIndirectInstanceAccelerationStructureDescriptor) InstanceCountBuffer() MTLBuffer {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("instanceCountBuffer"))
	return MTLBufferObjectFromID(rv)
}
func (i MTLIndirectInstanceAccelerationStructureDescriptor) SetInstanceCountBuffer(value MTLBuffer) {
	objc.Send[struct{}](i.ID, objc.Sel("setInstanceCountBuffer:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceCountBufferOffset
func (i MTLIndirectInstanceAccelerationStructureDescriptor) InstanceCountBufferOffset() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("instanceCountBufferOffset"))
	return rv
}
func (i MTLIndirectInstanceAccelerationStructureDescriptor) SetInstanceCountBufferOffset(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setInstanceCountBufferOffset:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceDescriptorBuffer
func (i MTLIndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorBuffer() MTLBuffer {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("instanceDescriptorBuffer"))
	return MTLBufferObjectFromID(rv)
}
func (i MTLIndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorBuffer(value MTLBuffer) {
	objc.Send[struct{}](i.ID, objc.Sel("setInstanceDescriptorBuffer:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceDescriptorBufferOffset
func (i MTLIndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorBufferOffset() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("instanceDescriptorBufferOffset"))
	return rv
}
func (i MTLIndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorBufferOffset(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setInstanceDescriptorBufferOffset:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceDescriptorStride
func (i MTLIndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorStride() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("instanceDescriptorStride"))
	return rv
}
func (i MTLIndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorStride(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setInstanceDescriptorStride:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceDescriptorType
func (i MTLIndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorType() MTLAccelerationStructureInstanceDescriptorType {
	rv := objc.Send[MTLAccelerationStructureInstanceDescriptorType](i.ID, objc.Sel("instanceDescriptorType"))
	return MTLAccelerationStructureInstanceDescriptorType(rv)
}
func (i MTLIndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorType(value MTLAccelerationStructureInstanceDescriptorType) {
	objc.Send[struct{}](i.ID, objc.Sel("setInstanceDescriptorType:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceTransformationMatrixLayout
func (i MTLIndirectInstanceAccelerationStructureDescriptor) InstanceTransformationMatrixLayout() MTLMatrixLayout {
	rv := objc.Send[MTLMatrixLayout](i.ID, objc.Sel("instanceTransformationMatrixLayout"))
	return MTLMatrixLayout(rv)
}
func (i MTLIndirectInstanceAccelerationStructureDescriptor) SetInstanceTransformationMatrixLayout(value MTLMatrixLayout) {
	objc.Send[struct{}](i.ID, objc.Sel("setInstanceTransformationMatrixLayout:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/maxInstanceCount
func (i MTLIndirectInstanceAccelerationStructureDescriptor) MaxInstanceCount() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("maxInstanceCount"))
	return rv
}
func (i MTLIndirectInstanceAccelerationStructureDescriptor) SetMaxInstanceCount(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setMaxInstanceCount:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/maxMotionTransformCount
func (i MTLIndirectInstanceAccelerationStructureDescriptor) MaxMotionTransformCount() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("maxMotionTransformCount"))
	return rv
}
func (i MTLIndirectInstanceAccelerationStructureDescriptor) SetMaxMotionTransformCount(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setMaxMotionTransformCount:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/motionTransformBuffer
func (i MTLIndirectInstanceAccelerationStructureDescriptor) MotionTransformBuffer() MTLBuffer {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("motionTransformBuffer"))
	return MTLBufferObjectFromID(rv)
}
func (i MTLIndirectInstanceAccelerationStructureDescriptor) SetMotionTransformBuffer(value MTLBuffer) {
	objc.Send[struct{}](i.ID, objc.Sel("setMotionTransformBuffer:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/motionTransformBufferOffset
func (i MTLIndirectInstanceAccelerationStructureDescriptor) MotionTransformBufferOffset() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("motionTransformBufferOffset"))
	return rv
}
func (i MTLIndirectInstanceAccelerationStructureDescriptor) SetMotionTransformBufferOffset(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setMotionTransformBufferOffset:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/motionTransformCountBuffer
func (i MTLIndirectInstanceAccelerationStructureDescriptor) MotionTransformCountBuffer() MTLBuffer {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("motionTransformCountBuffer"))
	return MTLBufferObjectFromID(rv)
}
func (i MTLIndirectInstanceAccelerationStructureDescriptor) SetMotionTransformCountBuffer(value MTLBuffer) {
	objc.Send[struct{}](i.ID, objc.Sel("setMotionTransformCountBuffer:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/motionTransformCountBufferOffset
func (i MTLIndirectInstanceAccelerationStructureDescriptor) MotionTransformCountBufferOffset() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("motionTransformCountBufferOffset"))
	return rv
}
func (i MTLIndirectInstanceAccelerationStructureDescriptor) SetMotionTransformCountBufferOffset(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setMotionTransformCountBufferOffset:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/motionTransformStride
func (i MTLIndirectInstanceAccelerationStructureDescriptor) MotionTransformStride() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("motionTransformStride"))
	return rv
}
func (i MTLIndirectInstanceAccelerationStructureDescriptor) SetMotionTransformStride(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setMotionTransformStride:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/motionTransformType
func (i MTLIndirectInstanceAccelerationStructureDescriptor) MotionTransformType() MTLTransformType {
	rv := objc.Send[MTLTransformType](i.ID, objc.Sel("motionTransformType"))
	return MTLTransformType(rv)
}
func (i MTLIndirectInstanceAccelerationStructureDescriptor) SetMotionTransformType(value MTLTransformType) {
	objc.Send[struct{}](i.ID, objc.Sel("setMotionTransformType:"), value)
}
