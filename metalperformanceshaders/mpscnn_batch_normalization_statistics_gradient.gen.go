// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNBatchNormalizationStatisticsGradient] class.
var (
	_MPSCNNBatchNormalizationStatisticsGradientClass     MPSCNNBatchNormalizationStatisticsGradientClass
	_MPSCNNBatchNormalizationStatisticsGradientClassOnce sync.Once
)

func getMPSCNNBatchNormalizationStatisticsGradientClass() MPSCNNBatchNormalizationStatisticsGradientClass {
	_MPSCNNBatchNormalizationStatisticsGradientClassOnce.Do(func() {
		_MPSCNNBatchNormalizationStatisticsGradientClass = MPSCNNBatchNormalizationStatisticsGradientClass{class: objc.GetClass("MPSCNNBatchNormalizationStatisticsGradient")}
	})
	return _MPSCNNBatchNormalizationStatisticsGradientClass
}

// GetMPSCNNBatchNormalizationStatisticsGradientClass returns the class object for MPSCNNBatchNormalizationStatisticsGradient.
func GetMPSCNNBatchNormalizationStatisticsGradientClass() MPSCNNBatchNormalizationStatisticsGradientClass {
	return getMPSCNNBatchNormalizationStatisticsGradientClass()
}

type MPSCNNBatchNormalizationStatisticsGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNBatchNormalizationStatisticsGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNBatchNormalizationStatisticsGradientClass) Alloc() MPSCNNBatchNormalizationStatisticsGradient {
	rv := objc.Send[MPSCNNBatchNormalizationStatisticsGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that stores the gradient of the loss function with respect to the
// batch statistics and batch normalization weights.
//
// # Initializers
//
//   - [MPSCNNBatchNormalizationStatisticsGradient.InitWithDeviceFusedNeuronDescriptor]
//
// # Instance Methods
//
//   - [MPSCNNBatchNormalizationStatisticsGradient.EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationStatisticsGradient
type MPSCNNBatchNormalizationStatisticsGradient struct {
	MPSCNNGradientKernel
}

// MPSCNNBatchNormalizationStatisticsGradientFromID constructs a [MPSCNNBatchNormalizationStatisticsGradient] from an objc.ID.
//
// An object that stores the gradient of the loss function with respect to the
// batch statistics and batch normalization weights.
func MPSCNNBatchNormalizationStatisticsGradientFromID(id objc.ID) MPSCNNBatchNormalizationStatisticsGradient {
	return MPSCNNBatchNormalizationStatisticsGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}

// NOTE: MPSCNNBatchNormalizationStatisticsGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNBatchNormalizationStatisticsGradient] class.
//
// # Initializers
//
//   - [IMPSCNNBatchNormalizationStatisticsGradient.InitWithDeviceFusedNeuronDescriptor]
//
// # Instance Methods
//
//   - [IMPSCNNBatchNormalizationStatisticsGradient.EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationStatisticsGradient
type IMPSCNNBatchNormalizationStatisticsGradient interface {
	IMPSCNNGradientKernel

	// Topic: Initializers

	InitWithDeviceFusedNeuronDescriptor(device metal.MTLDevice, fusedNeuronDescriptor IMPSNNNeuronDescriptor) MPSCNNBatchNormalizationStatisticsGradient

	// Topic: Instance Methods

	EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationState(commandBuffer metal.MTLCommandBuffer, sourceGradients MPSImageBatch, sourceImages MPSImageBatch, batchNormalizationState IMPSCNNBatchNormalizationState)
}

// Init initializes the instance.
func (c MPSCNNBatchNormalizationStatisticsGradient) Init() MPSCNNBatchNormalizationStatisticsGradient {
	rv := objc.Send[MPSCNNBatchNormalizationStatisticsGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNBatchNormalizationStatisticsGradient) Autorelease() MPSCNNBatchNormalizationStatisticsGradient {
	rv := objc.Send[MPSCNNBatchNormalizationStatisticsGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNBatchNormalizationStatisticsGradient creates a new MPSCNNBatchNormalizationStatisticsGradient instance.
func NewMPSCNNBatchNormalizationStatisticsGradient() MPSCNNBatchNormalizationStatisticsGradient {
	class := getMPSCNNBatchNormalizationStatisticsGradientClass()
	rv := objc.Send[MPSCNNBatchNormalizationStatisticsGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNBatchNormalizationStatisticsGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNBatchNormalizationStatisticsGradient {
	instance := getMPSCNNBatchNormalizationStatisticsGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNBatchNormalizationStatisticsGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationStatisticsGradient/init(coder:device:)
func NewCNNBatchNormalizationStatisticsGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNBatchNormalizationStatisticsGradient {
	instance := getMPSCNNBatchNormalizationStatisticsGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNBatchNormalizationStatisticsGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNBatchNormalizationStatisticsGradientWithDevice(device metal.MTLDevice) MPSCNNBatchNormalizationStatisticsGradient {
	instance := getMPSCNNBatchNormalizationStatisticsGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNBatchNormalizationStatisticsGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationStatisticsGradient/init(device:fusedNeuronDescriptor:)
func NewCNNBatchNormalizationStatisticsGradientWithDeviceFusedNeuronDescriptor(device metal.MTLDevice, fusedNeuronDescriptor IMPSNNNeuronDescriptor) MPSCNNBatchNormalizationStatisticsGradient {
	instance := getMPSCNNBatchNormalizationStatisticsGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:fusedNeuronDescriptor:"), device, fusedNeuronDescriptor)
	return MPSCNNBatchNormalizationStatisticsGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationStatisticsGradient/init(device:fusedNeuronDescriptor:)
func (c MPSCNNBatchNormalizationStatisticsGradient) InitWithDeviceFusedNeuronDescriptor(device metal.MTLDevice, fusedNeuronDescriptor IMPSNNNeuronDescriptor) MPSCNNBatchNormalizationStatisticsGradient {
	rv := objc.Send[MPSCNNBatchNormalizationStatisticsGradient](c.ID, objc.Sel("initWithDevice:fusedNeuronDescriptor:"), device, fusedNeuronDescriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationStatisticsGradient/encodeBatch(to:sourceGradients:sourceImages:batchNormalizationState:)
func (c MPSCNNBatchNormalizationStatisticsGradient) EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationState(commandBuffer metal.MTLCommandBuffer, sourceGradients MPSImageBatch, sourceImages MPSImageBatch, batchNormalizationState IMPSCNNBatchNormalizationState) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:batchNormalizationState:"), commandBuffer, sourceGradients, sourceImages, batchNormalizationState)
}
