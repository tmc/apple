// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNForwardLoss] class.
var (
	_MPSNNForwardLossClass     MPSNNForwardLossClass
	_MPSNNForwardLossClassOnce sync.Once
)

func getMPSNNForwardLossClass() MPSNNForwardLossClass {
	_MPSNNForwardLossClassOnce.Do(func() {
		_MPSNNForwardLossClass = MPSNNForwardLossClass{class: objc.GetClass("MPSNNForwardLoss")}
	})
	return _MPSNNForwardLossClass
}

// GetMPSNNForwardLossClass returns the class object for MPSNNForwardLoss.
func GetMPSNNForwardLossClass() MPSNNForwardLossClass {
	return getMPSNNForwardLossClass()
}

type MPSNNForwardLossClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNForwardLossClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNForwardLossClass) Alloc() MPSNNForwardLoss {
	rv := objc.Send[MPSNNForwardLoss](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNNForwardLoss.InitWithDeviceLossDescriptor]
//
// # Instance Properties
//
//   - [MPSNNForwardLoss.Delta]
//   - [MPSNNForwardLoss.SetDelta]
//   - [MPSNNForwardLoss.Epsilon]
//   - [MPSNNForwardLoss.SetEpsilon]
//   - [MPSNNForwardLoss.LabelSmoothing]
//   - [MPSNNForwardLoss.SetLabelSmoothing]
//   - [MPSNNForwardLoss.LossType]
//   - [MPSNNForwardLoss.NumberOfClasses]
//   - [MPSNNForwardLoss.ReduceAcrossBatch]
//   - [MPSNNForwardLoss.ReductionType]
//   - [MPSNNForwardLoss.Weight]
//   - [MPSNNForwardLoss.SetWeight]
//
// # Instance Methods
//
//   - [MPSNNForwardLoss.EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationImages]
//   - [MPSNNForwardLoss.EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationStateIsTemporary]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss
type MPSNNForwardLoss struct {
	MPSCNNKernel
}

// MPSNNForwardLossFromID constructs a [MPSNNForwardLoss] from an objc.ID.
func MPSNNForwardLossFromID(id objc.ID) MPSNNForwardLoss {
	return MPSNNForwardLoss{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSNNForwardLoss adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNForwardLoss] class.
//
// # Initializers
//
//   - [IMPSNNForwardLoss.InitWithDeviceLossDescriptor]
//
// # Instance Properties
//
//   - [IMPSNNForwardLoss.Delta]
//   - [IMPSNNForwardLoss.SetDelta]
//   - [IMPSNNForwardLoss.Epsilon]
//   - [IMPSNNForwardLoss.SetEpsilon]
//   - [IMPSNNForwardLoss.LabelSmoothing]
//   - [IMPSNNForwardLoss.SetLabelSmoothing]
//   - [IMPSNNForwardLoss.LossType]
//   - [IMPSNNForwardLoss.NumberOfClasses]
//   - [IMPSNNForwardLoss.ReduceAcrossBatch]
//   - [IMPSNNForwardLoss.ReductionType]
//   - [IMPSNNForwardLoss.Weight]
//   - [IMPSNNForwardLoss.SetWeight]
//
// # Instance Methods
//
//   - [IMPSNNForwardLoss.EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationImages]
//   - [IMPSNNForwardLoss.EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationStateIsTemporary]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss
type IMPSNNForwardLoss interface {
	IMPSCNNKernel

	// Topic: Initializers

	InitWithDeviceLossDescriptor(device metal.MTLDevice, lossDescriptor IMPSCNNLossDescriptor) MPSNNForwardLoss

	// Topic: Instance Properties

	Delta() float32
	SetDelta(value float32)
	Epsilon() float32
	SetEpsilon(value float32)
	LabelSmoothing() float32
	SetLabelSmoothing(value float32)
	LossType() MPSCNNLossType
	NumberOfClasses() uint
	ReduceAcrossBatch() bool
	ReductionType() MPSCNNReductionType
	Weight() float32
	SetWeight(value float32)

	// Topic: Instance Methods

	EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationImages(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, labels MPSImageBatch, weights MPSImageBatch, destinationStates MPSStateBatch, destinationImages MPSImageBatch)
	EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationStateIsTemporary(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, labels MPSImageBatch, weights MPSImageBatch, outStates MPSStateBatch, isTemporary bool) MPSImageBatch
}

// Init initializes the instance.
func (f MPSNNForwardLoss) Init() MPSNNForwardLoss {
	rv := objc.Send[MPSNNForwardLoss](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MPSNNForwardLoss) Autorelease() MPSNNForwardLoss {
	rv := objc.Send[MPSNNForwardLoss](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNForwardLoss creates a new MPSNNForwardLoss instance.
func NewMPSNNForwardLoss() MPSNNForwardLoss {
	class := getMPSNNForwardLossClass()
	rv := objc.Send[MPSNNForwardLoss](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewForwardLossWithCoder(aDecoder foundation.INSCoder) MPSNNForwardLoss {
	instance := getMPSNNForwardLossClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNForwardLossFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/init(coder:device:)
func NewForwardLossWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNForwardLoss {
	instance := getMPSNNForwardLossClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNForwardLossFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewForwardLossWithDevice(device metal.MTLDevice) MPSNNForwardLoss {
	instance := getMPSNNForwardLossClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNForwardLossFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/init(device:lossDescriptor:)
func NewForwardLossWithDeviceLossDescriptor(device metal.MTLDevice, lossDescriptor IMPSCNNLossDescriptor) MPSNNForwardLoss {
	instance := getMPSNNForwardLossClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, lossDescriptor)
	return MPSNNForwardLossFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/init(device:lossDescriptor:)
func (f MPSNNForwardLoss) InitWithDeviceLossDescriptor(device metal.MTLDevice, lossDescriptor IMPSCNNLossDescriptor) MPSNNForwardLoss {
	rv := objc.Send[MPSNNForwardLoss](f.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, lossDescriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/encodeBatch(commandBuffer:sourceImages:labels:weights:destinationStates:destinationImages:)
func (f MPSNNForwardLoss) EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationImages(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, labels MPSImageBatch, weights MPSImageBatch, destinationStates MPSStateBatch, destinationImages MPSImageBatch) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:weights:destinationStates:destinationImages:"), commandBuffer, sourceImages, labels, weights, destinationStates, destinationImages)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/encodeBatch(commandBuffer:sourceImages:labels:weights:outStates:isTemporary:)
func (f MPSNNForwardLoss) EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationStateIsTemporary(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, labels MPSImageBatch, weights MPSImageBatch, outStates MPSStateBatch, isTemporary bool) MPSImageBatch {
	rv := objc.Send[MPSImageBatch](f.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:weights:destinationStates:destinationStateIsTemporary:"), commandBuffer, sourceImages, labels, weights, outStates, isTemporary)
	return MPSImageBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/delta
func (f MPSNNForwardLoss) Delta() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("delta"))
	return rv
}
func (f MPSNNForwardLoss) SetDelta(value float32) {
	objc.Send[struct{}](f.ID, objc.Sel("setDelta:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/epsilon
func (f MPSNNForwardLoss) Epsilon() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("epsilon"))
	return rv
}
func (f MPSNNForwardLoss) SetEpsilon(value float32) {
	objc.Send[struct{}](f.ID, objc.Sel("setEpsilon:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/labelSmoothing
func (f MPSNNForwardLoss) LabelSmoothing() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("labelSmoothing"))
	return rv
}
func (f MPSNNForwardLoss) SetLabelSmoothing(value float32) {
	objc.Send[struct{}](f.ID, objc.Sel("setLabelSmoothing:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/lossType
func (f MPSNNForwardLoss) LossType() MPSCNNLossType {
	rv := objc.Send[MPSCNNLossType](f.ID, objc.Sel("lossType"))
	return MPSCNNLossType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/numberOfClasses
func (f MPSNNForwardLoss) NumberOfClasses() uint {
	rv := objc.Send[uint](f.ID, objc.Sel("numberOfClasses"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/reduceAcrossBatch
func (f MPSNNForwardLoss) ReduceAcrossBatch() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/reductionType
func (f MPSNNForwardLoss) ReductionType() MPSCNNReductionType {
	rv := objc.Send[MPSCNNReductionType](f.ID, objc.Sel("reductionType"))
	return MPSCNNReductionType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/weight
func (f MPSNNForwardLoss) Weight() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("weight"))
	return rv
}
func (f MPSNNForwardLoss) SetWeight(value float32) {
	objc.Send[struct{}](f.ID, objc.Sel("setWeight:"), value)
}
