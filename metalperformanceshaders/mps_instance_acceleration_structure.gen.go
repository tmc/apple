// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSInstanceAccelerationStructure] class.
var (
	_MPSInstanceAccelerationStructureClass     MPSInstanceAccelerationStructureClass
	_MPSInstanceAccelerationStructureClassOnce sync.Once
)

func getMPSInstanceAccelerationStructureClass() MPSInstanceAccelerationStructureClass {
	_MPSInstanceAccelerationStructureClassOnce.Do(func() {
		_MPSInstanceAccelerationStructureClass = MPSInstanceAccelerationStructureClass{class: objc.GetClass("MPSInstanceAccelerationStructure")}
	})
	return _MPSInstanceAccelerationStructureClass
}

// GetMPSInstanceAccelerationStructureClass returns the class object for MPSInstanceAccelerationStructure.
func GetMPSInstanceAccelerationStructureClass() MPSInstanceAccelerationStructureClass {
	return getMPSInstanceAccelerationStructureClass()
}

type MPSInstanceAccelerationStructureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSInstanceAccelerationStructureClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSInstanceAccelerationStructureClass) Alloc() MPSInstanceAccelerationStructure {
	rv := objc.Send[MPSInstanceAccelerationStructure](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An acceleration structure built over instances of other acceleration
// structures.
//
// # Instance Properties
//
//   - [MPSInstanceAccelerationStructure.AccelerationStructures]
//   - [MPSInstanceAccelerationStructure.SetAccelerationStructures]
//   - [MPSInstanceAccelerationStructure.InstanceBuffer]
//   - [MPSInstanceAccelerationStructure.SetInstanceBuffer]
//   - [MPSInstanceAccelerationStructure.InstanceBufferOffset]
//   - [MPSInstanceAccelerationStructure.SetInstanceBufferOffset]
//   - [MPSInstanceAccelerationStructure.InstanceCount]
//   - [MPSInstanceAccelerationStructure.SetInstanceCount]
//   - [MPSInstanceAccelerationStructure.MaskBuffer]
//   - [MPSInstanceAccelerationStructure.SetMaskBuffer]
//   - [MPSInstanceAccelerationStructure.MaskBufferOffset]
//   - [MPSInstanceAccelerationStructure.SetMaskBufferOffset]
//   - [MPSInstanceAccelerationStructure.TransformBuffer]
//   - [MPSInstanceAccelerationStructure.SetTransformBuffer]
//   - [MPSInstanceAccelerationStructure.TransformBufferOffset]
//   - [MPSInstanceAccelerationStructure.SetTransformBufferOffset]
//   - [MPSInstanceAccelerationStructure.TransformType]
//   - [MPSInstanceAccelerationStructure.SetTransformType]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSInstanceAccelerationStructure
type MPSInstanceAccelerationStructure struct {
	MPSAccelerationStructure
}

// MPSInstanceAccelerationStructureFromID constructs a [MPSInstanceAccelerationStructure] from an objc.ID.
//
// An acceleration structure built over instances of other acceleration
// structures.
func MPSInstanceAccelerationStructureFromID(id objc.ID) MPSInstanceAccelerationStructure {
	return MPSInstanceAccelerationStructure{MPSAccelerationStructure: MPSAccelerationStructureFromID(id)}
}

// NOTE: MPSInstanceAccelerationStructure adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSInstanceAccelerationStructure] class.
//
// # Instance Properties
//
//   - [IMPSInstanceAccelerationStructure.AccelerationStructures]
//   - [IMPSInstanceAccelerationStructure.SetAccelerationStructures]
//   - [IMPSInstanceAccelerationStructure.InstanceBuffer]
//   - [IMPSInstanceAccelerationStructure.SetInstanceBuffer]
//   - [IMPSInstanceAccelerationStructure.InstanceBufferOffset]
//   - [IMPSInstanceAccelerationStructure.SetInstanceBufferOffset]
//   - [IMPSInstanceAccelerationStructure.InstanceCount]
//   - [IMPSInstanceAccelerationStructure.SetInstanceCount]
//   - [IMPSInstanceAccelerationStructure.MaskBuffer]
//   - [IMPSInstanceAccelerationStructure.SetMaskBuffer]
//   - [IMPSInstanceAccelerationStructure.MaskBufferOffset]
//   - [IMPSInstanceAccelerationStructure.SetMaskBufferOffset]
//   - [IMPSInstanceAccelerationStructure.TransformBuffer]
//   - [IMPSInstanceAccelerationStructure.SetTransformBuffer]
//   - [IMPSInstanceAccelerationStructure.TransformBufferOffset]
//   - [IMPSInstanceAccelerationStructure.SetTransformBufferOffset]
//   - [IMPSInstanceAccelerationStructure.TransformType]
//   - [IMPSInstanceAccelerationStructure.SetTransformType]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSInstanceAccelerationStructure
type IMPSInstanceAccelerationStructure interface {
	IMPSAccelerationStructure

	// Topic: Instance Properties

	AccelerationStructures() []MPSPolygonAccelerationStructure
	SetAccelerationStructures(value []MPSPolygonAccelerationStructure)
	InstanceBuffer() metal.MTLBuffer
	SetInstanceBuffer(value metal.MTLBuffer)
	InstanceBufferOffset() uint
	SetInstanceBufferOffset(value uint)
	InstanceCount() uint
	SetInstanceCount(value uint)
	MaskBuffer() metal.MTLBuffer
	SetMaskBuffer(value metal.MTLBuffer)
	MaskBufferOffset() uint
	SetMaskBufferOffset(value uint)
	TransformBuffer() metal.MTLBuffer
	SetTransformBuffer(value metal.MTLBuffer)
	TransformBufferOffset() uint
	SetTransformBufferOffset(value uint)
	TransformType() MPSTransformType
	SetTransformType(value MPSTransformType)
}

// Init initializes the instance.
func (i MPSInstanceAccelerationStructure) Init() MPSInstanceAccelerationStructure {
	rv := objc.Send[MPSInstanceAccelerationStructure](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSInstanceAccelerationStructure) Autorelease() MPSInstanceAccelerationStructure {
	rv := objc.Send[MPSInstanceAccelerationStructure](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSInstanceAccelerationStructure creates a new MPSInstanceAccelerationStructure instance.
func NewMPSInstanceAccelerationStructure() MPSInstanceAccelerationStructure {
	class := getMPSInstanceAccelerationStructureClass()
	rv := objc.Send[MPSInstanceAccelerationStructure](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewInstanceAccelerationStructureWithCoder(aDecoder foundation.INSCoder) MPSInstanceAccelerationStructure {
	instance := getMPSInstanceAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSInstanceAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(coder:device:)
func NewInstanceAccelerationStructureWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSInstanceAccelerationStructure {
	instance := getMPSInstanceAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSInstanceAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(coder:group:)
func NewInstanceAccelerationStructureWithCoderGroup(aDecoder foundation.INSCoder, group IMPSAccelerationStructureGroup) MPSInstanceAccelerationStructure {
	instance := getMPSInstanceAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:group:"), aDecoder, group)
	return MPSInstanceAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(device:)
func NewInstanceAccelerationStructureWithDevice(device metal.MTLDevice) MPSInstanceAccelerationStructure {
	instance := getMPSInstanceAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSInstanceAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(group:)
func NewInstanceAccelerationStructureWithGroup(group IMPSAccelerationStructureGroup) MPSInstanceAccelerationStructure {
	instance := getMPSInstanceAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGroup:"), group)
	return MPSInstanceAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSInstanceAccelerationStructure/accelerationStructures
func (i MPSInstanceAccelerationStructure) AccelerationStructures() []MPSPolygonAccelerationStructure {
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("accelerationStructures"))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSPolygonAccelerationStructure {
		return MPSPolygonAccelerationStructureFromID(id)
	})
}
func (i MPSInstanceAccelerationStructure) SetAccelerationStructures(value []MPSPolygonAccelerationStructure) {
	objc.Send[struct{}](i.ID, objc.Sel("setAccelerationStructures:"), objectivec.IObjectSliceToNSArray(value))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSInstanceAccelerationStructure/instanceBuffer
func (i MPSInstanceAccelerationStructure) InstanceBuffer() metal.MTLBuffer {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("instanceBuffer"))
	return metal.MTLBufferObjectFromID(rv)
}
func (i MPSInstanceAccelerationStructure) SetInstanceBuffer(value metal.MTLBuffer) {
	objc.Send[struct{}](i.ID, objc.Sel("setInstanceBuffer:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSInstanceAccelerationStructure/instanceBufferOffset
func (i MPSInstanceAccelerationStructure) InstanceBufferOffset() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("instanceBufferOffset"))
	return rv
}
func (i MPSInstanceAccelerationStructure) SetInstanceBufferOffset(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setInstanceBufferOffset:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSInstanceAccelerationStructure/instanceCount
func (i MPSInstanceAccelerationStructure) InstanceCount() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("instanceCount"))
	return rv
}
func (i MPSInstanceAccelerationStructure) SetInstanceCount(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setInstanceCount:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSInstanceAccelerationStructure/maskBuffer
func (i MPSInstanceAccelerationStructure) MaskBuffer() metal.MTLBuffer {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("maskBuffer"))
	return metal.MTLBufferObjectFromID(rv)
}
func (i MPSInstanceAccelerationStructure) SetMaskBuffer(value metal.MTLBuffer) {
	objc.Send[struct{}](i.ID, objc.Sel("setMaskBuffer:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSInstanceAccelerationStructure/maskBufferOffset
func (i MPSInstanceAccelerationStructure) MaskBufferOffset() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("maskBufferOffset"))
	return rv
}
func (i MPSInstanceAccelerationStructure) SetMaskBufferOffset(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setMaskBufferOffset:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSInstanceAccelerationStructure/transformBuffer
func (i MPSInstanceAccelerationStructure) TransformBuffer() metal.MTLBuffer {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("transformBuffer"))
	return metal.MTLBufferObjectFromID(rv)
}
func (i MPSInstanceAccelerationStructure) SetTransformBuffer(value metal.MTLBuffer) {
	objc.Send[struct{}](i.ID, objc.Sel("setTransformBuffer:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSInstanceAccelerationStructure/transformBufferOffset
func (i MPSInstanceAccelerationStructure) TransformBufferOffset() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("transformBufferOffset"))
	return rv
}
func (i MPSInstanceAccelerationStructure) SetTransformBufferOffset(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setTransformBufferOffset:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSInstanceAccelerationStructure/transformType
func (i MPSInstanceAccelerationStructure) TransformType() MPSTransformType {
	rv := objc.Send[MPSTransformType](i.ID, objc.Sel("transformType"))
	return MPSTransformType(rv)
}
func (i MPSInstanceAccelerationStructure) SetTransformType(value MPSTransformType) {
	objc.Send[struct{}](i.ID, objc.Sel("setTransformType:"), value)
}
