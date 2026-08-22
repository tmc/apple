// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNLossGradient] class.
var (
	_MPSNNLossGradientClass     MPSNNLossGradientClass
	_MPSNNLossGradientClassOnce sync.Once
)

func getMPSNNLossGradientClass() MPSNNLossGradientClass {
	_MPSNNLossGradientClassOnce.Do(func() {
		_MPSNNLossGradientClass = MPSNNLossGradientClass{class: objc.GetClass("MPSNNLossGradient")}
	})
	return _MPSNNLossGradientClass
}

// GetMPSNNLossGradientClass returns the class object for MPSNNLossGradient.
func GetMPSNNLossGradientClass() MPSNNLossGradientClass {
	return getMPSNNLossGradientClass()
}

type MPSNNLossGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNLossGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNLossGradientClass) Alloc() MPSNNLossGradient {
	rv := objc.Send[MPSNNLossGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNNLossGradient.InitWithDeviceLossDescriptor]
//
// # Instance Properties
//
//   - [MPSNNLossGradient.ComputeLabelGradients]
//   - [MPSNNLossGradient.SetComputeLabelGradients]
//   - [MPSNNLossGradient.Delta]
//   - [MPSNNLossGradient.SetDelta]
//   - [MPSNNLossGradient.Epsilon]
//   - [MPSNNLossGradient.SetEpsilon]
//   - [MPSNNLossGradient.LabelSmoothing]
//   - [MPSNNLossGradient.SetLabelSmoothing]
//   - [MPSNNLossGradient.LossType]
//   - [MPSNNLossGradient.NumberOfClasses]
//   - [MPSNNLossGradient.ReduceAcrossBatch]
//   - [MPSNNLossGradient.ReductionType]
//   - [MPSNNLossGradient.Weight]
//   - [MPSNNLossGradient.SetWeight]
//
// # Instance Methods
//
//   - [MPSNNLossGradient.EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStates]
//   - [MPSNNLossGradient.EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStatesDestinationGradients]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient
type MPSNNLossGradient struct {
	MPSCNNBinaryKernel
}

// MPSNNLossGradientFromID constructs a [MPSNNLossGradient] from an objc.ID.
func MPSNNLossGradientFromID(id objc.ID) MPSNNLossGradient {
	return MPSNNLossGradient{MPSCNNBinaryKernel: MPSCNNBinaryKernelFromID(id)}
}

// NOTE: MPSNNLossGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNLossGradient] class.
//
// # Initializers
//
//   - [IMPSNNLossGradient.InitWithDeviceLossDescriptor]
//
// # Instance Properties
//
//   - [IMPSNNLossGradient.ComputeLabelGradients]
//   - [IMPSNNLossGradient.SetComputeLabelGradients]
//   - [IMPSNNLossGradient.Delta]
//   - [IMPSNNLossGradient.SetDelta]
//   - [IMPSNNLossGradient.Epsilon]
//   - [IMPSNNLossGradient.SetEpsilon]
//   - [IMPSNNLossGradient.LabelSmoothing]
//   - [IMPSNNLossGradient.SetLabelSmoothing]
//   - [IMPSNNLossGradient.LossType]
//   - [IMPSNNLossGradient.NumberOfClasses]
//   - [IMPSNNLossGradient.ReduceAcrossBatch]
//   - [IMPSNNLossGradient.ReductionType]
//   - [IMPSNNLossGradient.Weight]
//   - [IMPSNNLossGradient.SetWeight]
//
// # Instance Methods
//
//   - [IMPSNNLossGradient.EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStates]
//   - [IMPSNNLossGradient.EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStatesDestinationGradients]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient
type IMPSNNLossGradient interface {
	IMPSCNNBinaryKernel

	// Topic: Initializers

	InitWithDeviceLossDescriptor(device metal.MTLDevice, lossDescriptor IMPSCNNLossDescriptor) MPSNNLossGradient

	// Topic: Instance Properties

	ComputeLabelGradients() bool
	SetComputeLabelGradients(value bool)
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

	EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStates(commandBuffer metal.MTLCommandBuffer, sourceGradients MPSImageBatch, sourceImages MPSImageBatch, labels MPSImageBatch, weights MPSImageBatch, sourceStates MPSStateBatch) MPSImageBatch
	EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStatesDestinationGradients(commandBuffer metal.MTLCommandBuffer, sourceGradients MPSImageBatch, sourceImages MPSImageBatch, labels MPSImageBatch, weights MPSImageBatch, sourceStates MPSStateBatch, destinationGradients MPSImageBatch)
}

// Init initializes the instance.
func (l MPSNNLossGradient) Init() MPSNNLossGradient {
	rv := objc.Send[MPSNNLossGradient](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MPSNNLossGradient) Autorelease() MPSNNLossGradient {
	rv := objc.Send[MPSNNLossGradient](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNLossGradient creates a new MPSNNLossGradient instance.
func NewMPSNNLossGradient() MPSNNLossGradient {
	class := getMPSNNLossGradientClass()
	rv := objc.Send[MPSNNLossGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewLossGradientWithCoder(aDecoder foundation.INSCoder) MPSNNLossGradient {
	instance := getMPSNNLossGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNLossGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient/init(coder:device:)
func NewLossGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNLossGradient {
	instance := getMPSNNLossGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNLossGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/init(device:)
func NewLossGradientWithDevice(device metal.MTLDevice) MPSNNLossGradient {
	instance := getMPSNNLossGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNLossGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient/init(device:lossDescriptor:)
func NewLossGradientWithDeviceLossDescriptor(device metal.MTLDevice, lossDescriptor IMPSCNNLossDescriptor) MPSNNLossGradient {
	instance := getMPSNNLossGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, lossDescriptor)
	return MPSNNLossGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient/init(device:lossDescriptor:)
func (l MPSNNLossGradient) InitWithDeviceLossDescriptor(device metal.MTLDevice, lossDescriptor IMPSCNNLossDescriptor) MPSNNLossGradient {
	rv := objc.Send[MPSNNLossGradient](l.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, lossDescriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient/encodeBatch(commandBuffer:sourceGradients:sourceImages:labels:weights:sourceStates:)
func (l MPSNNLossGradient) EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStates(commandBuffer metal.MTLCommandBuffer, sourceGradients MPSImageBatch, sourceImages MPSImageBatch, labels MPSImageBatch, weights MPSImageBatch, sourceStates MPSStateBatch) MPSImageBatch {
	rv := objc.Send[MPSImageBatch](l.ID, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:labels:weights:sourceStates:"), commandBuffer, sourceGradients, sourceImages, labels, weights, sourceStates)
	return MPSImageBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient/encodeBatch(commandBuffer:sourceGradients:sourceImages:labels:weights:sourceStates:destinationGradients:)
func (l MPSNNLossGradient) EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStatesDestinationGradients(commandBuffer metal.MTLCommandBuffer, sourceGradients MPSImageBatch, sourceImages MPSImageBatch, labels MPSImageBatch, weights MPSImageBatch, sourceStates MPSStateBatch, destinationGradients MPSImageBatch) {
	objc.Send[objc.ID](l.ID, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:labels:weights:sourceStates:destinationGradients:"), commandBuffer, sourceGradients, sourceImages, labels, weights, sourceStates, destinationGradients)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient/computeLabelGradients
func (l MPSNNLossGradient) ComputeLabelGradients() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("computeLabelGradients"))
	return rv
}
func (l MPSNNLossGradient) SetComputeLabelGradients(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setComputeLabelGradients:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient/delta
func (l MPSNNLossGradient) Delta() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("delta"))
	return rv
}
func (l MPSNNLossGradient) SetDelta(value float32) {
	objc.Send[struct{}](l.ID, objc.Sel("setDelta:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient/epsilon
func (l MPSNNLossGradient) Epsilon() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("epsilon"))
	return rv
}
func (l MPSNNLossGradient) SetEpsilon(value float32) {
	objc.Send[struct{}](l.ID, objc.Sel("setEpsilon:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient/labelSmoothing
func (l MPSNNLossGradient) LabelSmoothing() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("labelSmoothing"))
	return rv
}
func (l MPSNNLossGradient) SetLabelSmoothing(value float32) {
	objc.Send[struct{}](l.ID, objc.Sel("setLabelSmoothing:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient/lossType
func (l MPSNNLossGradient) LossType() MPSCNNLossType {
	rv := objc.Send[MPSCNNLossType](l.ID, objc.Sel("lossType"))
	return MPSCNNLossType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient/numberOfClasses
func (l MPSNNLossGradient) NumberOfClasses() uint {
	rv := objc.Send[uint](l.ID, objc.Sel("numberOfClasses"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient/reduceAcrossBatch
func (l MPSNNLossGradient) ReduceAcrossBatch() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient/reductionType
func (l MPSNNLossGradient) ReductionType() MPSCNNReductionType {
	rv := objc.Send[MPSCNNReductionType](l.ID, objc.Sel("reductionType"))
	return MPSCNNReductionType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient/weight
func (l MPSNNLossGradient) Weight() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("weight"))
	return rv
}
func (l MPSNNLossGradient) SetWeight(value float32) {
	objc.Send[struct{}](l.ID, objc.Sel("setWeight:"), value)
}
