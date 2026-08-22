// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNBatchNormalizationGradient] class.
var (
	_MPSCNNBatchNormalizationGradientClass     MPSCNNBatchNormalizationGradientClass
	_MPSCNNBatchNormalizationGradientClassOnce sync.Once
)

func getMPSCNNBatchNormalizationGradientClass() MPSCNNBatchNormalizationGradientClass {
	_MPSCNNBatchNormalizationGradientClassOnce.Do(func() {
		_MPSCNNBatchNormalizationGradientClass = MPSCNNBatchNormalizationGradientClass{class: objc.GetClass("MPSCNNBatchNormalizationGradient")}
	})
	return _MPSCNNBatchNormalizationGradientClass
}

// GetMPSCNNBatchNormalizationGradientClass returns the class object for MPSCNNBatchNormalizationGradient.
func GetMPSCNNBatchNormalizationGradientClass() MPSCNNBatchNormalizationGradientClass {
	return getMPSCNNBatchNormalizationGradientClass()
}

type MPSCNNBatchNormalizationGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNBatchNormalizationGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNBatchNormalizationGradientClass) Alloc() MPSCNNBatchNormalizationGradient {
	rv := objc.Send[MPSCNNBatchNormalizationGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient batch normalization kernel.
//
// # Initializers
//
//   - [MPSCNNBatchNormalizationGradient.InitWithDeviceFusedNeuronDescriptor]
//
// # Instance Methods
//
//   - [MPSCNNBatchNormalizationGradient.EncodeToCommandBufferSourceGradientSourceImageBatchNormalizationState]
//   - [MPSCNNBatchNormalizationGradient.EncodeToCommandBufferSourceGradientSourceImageBatchNormalizationStateDestinationGradient]
//   - [MPSCNNBatchNormalizationGradient.EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationState]
//   - [MPSCNNBatchNormalizationGradient.EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationStateDestinationGradients]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationGradient
type MPSCNNBatchNormalizationGradient struct {
	MPSCNNGradientKernel
}

// MPSCNNBatchNormalizationGradientFromID constructs a [MPSCNNBatchNormalizationGradient] from an objc.ID.
//
// A gradient batch normalization kernel.
func MPSCNNBatchNormalizationGradientFromID(id objc.ID) MPSCNNBatchNormalizationGradient {
	return MPSCNNBatchNormalizationGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}

// NOTE: MPSCNNBatchNormalizationGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNBatchNormalizationGradient] class.
//
// # Initializers
//
//   - [IMPSCNNBatchNormalizationGradient.InitWithDeviceFusedNeuronDescriptor]
//
// # Instance Methods
//
//   - [IMPSCNNBatchNormalizationGradient.EncodeToCommandBufferSourceGradientSourceImageBatchNormalizationState]
//   - [IMPSCNNBatchNormalizationGradient.EncodeToCommandBufferSourceGradientSourceImageBatchNormalizationStateDestinationGradient]
//   - [IMPSCNNBatchNormalizationGradient.EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationState]
//   - [IMPSCNNBatchNormalizationGradient.EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationStateDestinationGradients]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationGradient
type IMPSCNNBatchNormalizationGradient interface {
	IMPSCNNGradientKernel

	// Topic: Initializers

	InitWithDeviceFusedNeuronDescriptor(device metal.MTLDevice, fusedNeuronDescriptor IMPSNNNeuronDescriptor) MPSCNNBatchNormalizationGradient

	// Topic: Instance Methods

	EncodeToCommandBufferSourceGradientSourceImageBatchNormalizationState(commandBuffer metal.MTLCommandBuffer, sourceGradient IMPSImage, sourceImage IMPSImage, batchNormalizationState IMPSCNNBatchNormalizationState) IMPSImage
	EncodeToCommandBufferSourceGradientSourceImageBatchNormalizationStateDestinationGradient(commandBuffer metal.MTLCommandBuffer, sourceGradient IMPSImage, sourceImage IMPSImage, batchNormalizationState IMPSCNNBatchNormalizationState, destinationGradient IMPSImage)
	EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationState(commandBuffer metal.MTLCommandBuffer, sourceGradients MPSImageBatch, sourceImages MPSImageBatch, batchNormalizationState IMPSCNNBatchNormalizationState) MPSImageBatch
	EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationStateDestinationGradients(commandBuffer metal.MTLCommandBuffer, sourceGradients MPSImageBatch, sourceImages MPSImageBatch, batchNormalizationState IMPSCNNBatchNormalizationState, destinationGradients MPSImageBatch)
}

// Init initializes the instance.
func (c MPSCNNBatchNormalizationGradient) Init() MPSCNNBatchNormalizationGradient {
	rv := objc.Send[MPSCNNBatchNormalizationGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNBatchNormalizationGradient) Autorelease() MPSCNNBatchNormalizationGradient {
	rv := objc.Send[MPSCNNBatchNormalizationGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNBatchNormalizationGradient creates a new MPSCNNBatchNormalizationGradient instance.
func NewMPSCNNBatchNormalizationGradient() MPSCNNBatchNormalizationGradient {
	class := getMPSCNNBatchNormalizationGradientClass()
	rv := objc.Send[MPSCNNBatchNormalizationGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNBatchNormalizationGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNBatchNormalizationGradient {
	instance := getMPSCNNBatchNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNBatchNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationGradient/init(coder:device:)
func NewCNNBatchNormalizationGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNBatchNormalizationGradient {
	instance := getMPSCNNBatchNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNBatchNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNBatchNormalizationGradientWithDevice(device metal.MTLDevice) MPSCNNBatchNormalizationGradient {
	instance := getMPSCNNBatchNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNBatchNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationGradient/init(device:fusedNeuronDescriptor:)
func NewCNNBatchNormalizationGradientWithDeviceFusedNeuronDescriptor(device metal.MTLDevice, fusedNeuronDescriptor IMPSNNNeuronDescriptor) MPSCNNBatchNormalizationGradient {
	instance := getMPSCNNBatchNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:fusedNeuronDescriptor:"), device, fusedNeuronDescriptor)
	return MPSCNNBatchNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationGradient/init(device:fusedNeuronDescriptor:)
func (c MPSCNNBatchNormalizationGradient) InitWithDeviceFusedNeuronDescriptor(device metal.MTLDevice, fusedNeuronDescriptor IMPSNNNeuronDescriptor) MPSCNNBatchNormalizationGradient {
	rv := objc.Send[MPSCNNBatchNormalizationGradient](c.ID, objc.Sel("initWithDevice:fusedNeuronDescriptor:"), device, fusedNeuronDescriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationGradient/encode(to:sourceGradient:sourceImage:batchNormalizationState:)
func (c MPSCNNBatchNormalizationGradient) EncodeToCommandBufferSourceGradientSourceImageBatchNormalizationState(commandBuffer metal.MTLCommandBuffer, sourceGradient IMPSImage, sourceImage IMPSImage, batchNormalizationState IMPSCNNBatchNormalizationState) IMPSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:sourceGradient:sourceImage:batchNormalizationState:"), commandBuffer, sourceGradient, sourceImage, batchNormalizationState)
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationGradient/encode(to:sourceGradient:sourceImage:batchNormalizationState:destinationGradient:)
func (c MPSCNNBatchNormalizationGradient) EncodeToCommandBufferSourceGradientSourceImageBatchNormalizationStateDestinationGradient(commandBuffer metal.MTLCommandBuffer, sourceGradient IMPSImage, sourceImage IMPSImage, batchNormalizationState IMPSCNNBatchNormalizationState, destinationGradient IMPSImage) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:sourceGradient:sourceImage:batchNormalizationState:destinationGradient:"), commandBuffer, sourceGradient, sourceImage, batchNormalizationState, destinationGradient)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationGradient/encodeBatch(to:sourceGradients:sourceImages:batchNormalizationState:)
func (c MPSCNNBatchNormalizationGradient) EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationState(commandBuffer metal.MTLCommandBuffer, sourceGradients MPSImageBatch, sourceImages MPSImageBatch, batchNormalizationState IMPSCNNBatchNormalizationState) MPSImageBatch {
	rv := objc.Send[MPSImageBatch](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:batchNormalizationState:"), commandBuffer, sourceGradients, sourceImages, batchNormalizationState)
	return MPSImageBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationGradient/encodeBatch(to:sourceGradients:sourceImages:batchNormalizationState:destinationGradients:)
func (c MPSCNNBatchNormalizationGradient) EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationStateDestinationGradients(commandBuffer metal.MTLCommandBuffer, sourceGradients MPSImageBatch, sourceImages MPSImageBatch, batchNormalizationState IMPSCNNBatchNormalizationState, destinationGradients MPSImageBatch) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:batchNormalizationState:destinationGradients:"), commandBuffer, sourceGradients, sourceImages, batchNormalizationState, destinationGradients)
}
