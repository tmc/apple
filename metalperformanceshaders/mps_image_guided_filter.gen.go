// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageGuidedFilter] class.
var (
	_MPSImageGuidedFilterClass     MPSImageGuidedFilterClass
	_MPSImageGuidedFilterClassOnce sync.Once
)

func getMPSImageGuidedFilterClass() MPSImageGuidedFilterClass {
	_MPSImageGuidedFilterClassOnce.Do(func() {
		_MPSImageGuidedFilterClass = MPSImageGuidedFilterClass{class: objc.GetClass("MPSImageGuidedFilter")}
	})
	return _MPSImageGuidedFilterClass
}

// GetMPSImageGuidedFilterClass returns the class object for MPSImageGuidedFilter.
func GetMPSImageGuidedFilterClass() MPSImageGuidedFilterClass {
	return getMPSImageGuidedFilterClass()
}

type MPSImageGuidedFilterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageGuidedFilterClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageGuidedFilterClass) Alloc() MPSImageGuidedFilter {
	rv := objc.Send[MPSImageGuidedFilter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that performs edge-aware filtering on an image.
//
// # Initializers
//
//   - [MPSImageGuidedFilter.InitWithDeviceKernelDiameter]
//
// # Instance Properties
//
//   - [MPSImageGuidedFilter.Epsilon]
//   - [MPSImageGuidedFilter.SetEpsilon]
//   - [MPSImageGuidedFilter.KernelDiameter]
//   - [MPSImageGuidedFilter.ReconstructOffset]
//   - [MPSImageGuidedFilter.SetReconstructOffset]
//   - [MPSImageGuidedFilter.ReconstructScale]
//   - [MPSImageGuidedFilter.SetReconstructScale]
//
// # Instance Methods
//
//   - [MPSImageGuidedFilter.EncodeReconstructionToCommandBufferGuidanceTextureCoefficientsTextureACoefficientsTextureBDestinationTexture]
//   - [MPSImageGuidedFilter.EncodeReconstructionToCommandBufferGuidanceTextureCoefficientsTextureDestinationTexture]
//   - [MPSImageGuidedFilter.EncodeRegressionToCommandBufferSourceTextureGuidanceTextureWeightsTextureDestinationCoefficientsTextureADestinationCoefficientsTextureB]
//   - [MPSImageGuidedFilter.EncodeRegressionToCommandBufferSourceTextureGuidanceTextureWeightsTextureDestinationCoefficientsTexture]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGuidedFilter
type MPSImageGuidedFilter struct {
	MPSKernel
}

// MPSImageGuidedFilterFromID constructs a [MPSImageGuidedFilter] from an objc.ID.
//
// A filter that performs edge-aware filtering on an image.
func MPSImageGuidedFilterFromID(id objc.ID) MPSImageGuidedFilter {
	return MPSImageGuidedFilter{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSImageGuidedFilter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageGuidedFilter] class.
//
// # Initializers
//
//   - [IMPSImageGuidedFilter.InitWithDeviceKernelDiameter]
//
// # Instance Properties
//
//   - [IMPSImageGuidedFilter.Epsilon]
//   - [IMPSImageGuidedFilter.SetEpsilon]
//   - [IMPSImageGuidedFilter.KernelDiameter]
//   - [IMPSImageGuidedFilter.ReconstructOffset]
//   - [IMPSImageGuidedFilter.SetReconstructOffset]
//   - [IMPSImageGuidedFilter.ReconstructScale]
//   - [IMPSImageGuidedFilter.SetReconstructScale]
//
// # Instance Methods
//
//   - [IMPSImageGuidedFilter.EncodeReconstructionToCommandBufferGuidanceTextureCoefficientsTextureACoefficientsTextureBDestinationTexture]
//   - [IMPSImageGuidedFilter.EncodeReconstructionToCommandBufferGuidanceTextureCoefficientsTextureDestinationTexture]
//   - [IMPSImageGuidedFilter.EncodeRegressionToCommandBufferSourceTextureGuidanceTextureWeightsTextureDestinationCoefficientsTextureADestinationCoefficientsTextureB]
//   - [IMPSImageGuidedFilter.EncodeRegressionToCommandBufferSourceTextureGuidanceTextureWeightsTextureDestinationCoefficientsTexture]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGuidedFilter
type IMPSImageGuidedFilter interface {
	IMPSKernel

	// Topic: Initializers

	InitWithDeviceKernelDiameter(device metal.MTLDevice, kernelDiameter uint) MPSImageGuidedFilter

	// Topic: Instance Properties

	Epsilon() float32
	SetEpsilon(value float32)
	KernelDiameter() uint
	ReconstructOffset() float32
	SetReconstructOffset(value float32)
	ReconstructScale() float32
	SetReconstructScale(value float32)

	// Topic: Instance Methods

	EncodeReconstructionToCommandBufferGuidanceTextureCoefficientsTextureACoefficientsTextureBDestinationTexture(commandBuffer metal.MTLCommandBuffer, guidanceTexture metal.MTLTexture, coefficientsTextureA metal.MTLTexture, coefficientsTextureB metal.MTLTexture, destinationTexture metal.MTLTexture)
	EncodeReconstructionToCommandBufferGuidanceTextureCoefficientsTextureDestinationTexture(commandBuffer metal.MTLCommandBuffer, guidanceTexture metal.MTLTexture, coefficientsTexture metal.MTLTexture, destinationTexture metal.MTLTexture)
	EncodeRegressionToCommandBufferSourceTextureGuidanceTextureWeightsTextureDestinationCoefficientsTextureADestinationCoefficientsTextureB(commandBuffer metal.MTLCommandBuffer, sourceTexture metal.MTLTexture, guidanceTexture metal.MTLTexture, weightsTexture metal.MTLTexture, destinationCoefficientsTextureA metal.MTLTexture, destinationCoefficientsTextureB metal.MTLTexture)
	EncodeRegressionToCommandBufferSourceTextureGuidanceTextureWeightsTextureDestinationCoefficientsTexture(commandBuffer metal.MTLCommandBuffer, sourceTexture metal.MTLTexture, guidanceTexture metal.MTLTexture, weightsTexture metal.MTLTexture, destinationCoefficientsTexture metal.MTLTexture)
}

// Init initializes the instance.
func (i MPSImageGuidedFilter) Init() MPSImageGuidedFilter {
	rv := objc.Send[MPSImageGuidedFilter](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageGuidedFilter) Autorelease() MPSImageGuidedFilter {
	rv := objc.Send[MPSImageGuidedFilter](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageGuidedFilter creates a new MPSImageGuidedFilter instance.
func NewMPSImageGuidedFilter() MPSImageGuidedFilter {
	class := getMPSImageGuidedFilterClass()
	rv := objc.Send[MPSImageGuidedFilter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageGuidedFilterWithCoder(aDecoder foundation.INSCoder) MPSImageGuidedFilter {
	instance := getMPSImageGuidedFilterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageGuidedFilterFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGuidedFilter/init(coder:device:)
func NewImageGuidedFilterWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageGuidedFilter {
	instance := getMPSImageGuidedFilterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageGuidedFilterFromID(rv)
}

// Initializes a new kernel object.
//
// device: The Metal device on which the kernel will be used.
//
// # Return Value
//
// An initialized kernel object.
//
// # Discussion
//
// This method fails if the device is not supported. Query the
// [MPSSupportsMTLDevice] function to determine whether the device is
// supported.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(device:)
func NewImageGuidedFilterWithDevice(device metal.MTLDevice) MPSImageGuidedFilter {
	instance := getMPSImageGuidedFilterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageGuidedFilterFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGuidedFilter/init(device:kernelDiameter:)
func NewImageGuidedFilterWithDeviceKernelDiameter(device metal.MTLDevice, kernelDiameter uint) MPSImageGuidedFilter {
	instance := getMPSImageGuidedFilterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelDiameter:"), device, kernelDiameter)
	return MPSImageGuidedFilterFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGuidedFilter/init(device:kernelDiameter:)
func (i MPSImageGuidedFilter) InitWithDeviceKernelDiameter(device metal.MTLDevice, kernelDiameter uint) MPSImageGuidedFilter {
	rv := objc.Send[MPSImageGuidedFilter](i.ID, objc.Sel("initWithDevice:kernelDiameter:"), device, kernelDiameter)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGuidedFilter/encodeReconstruction(commandBuffer:guidance:coefficientsA:coefficientsB:destination:)
func (i MPSImageGuidedFilter) EncodeReconstructionToCommandBufferGuidanceTextureCoefficientsTextureACoefficientsTextureBDestinationTexture(commandBuffer metal.MTLCommandBuffer, guidanceTexture metal.MTLTexture, coefficientsTextureA metal.MTLTexture, coefficientsTextureB metal.MTLTexture, destinationTexture metal.MTLTexture) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeReconstructionToCommandBuffer:guidanceTexture:coefficientsTextureA:coefficientsTextureB:destinationTexture:"), commandBuffer, guidanceTexture, coefficientsTextureA, coefficientsTextureB, destinationTexture)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGuidedFilter/encodeReconstruction(to:guidanceTexture:coefficientsTexture:destinationTexture:)
func (i MPSImageGuidedFilter) EncodeReconstructionToCommandBufferGuidanceTextureCoefficientsTextureDestinationTexture(commandBuffer metal.MTLCommandBuffer, guidanceTexture metal.MTLTexture, coefficientsTexture metal.MTLTexture, destinationTexture metal.MTLTexture) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeReconstructionToCommandBuffer:guidanceTexture:coefficientsTexture:destinationTexture:"), commandBuffer, guidanceTexture, coefficientsTexture, destinationTexture)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGuidedFilter/encodeRegression(commandBuffer:source:guidance:weights:destinationCoefficientsA:destinationCoefficientsB:)
func (i MPSImageGuidedFilter) EncodeRegressionToCommandBufferSourceTextureGuidanceTextureWeightsTextureDestinationCoefficientsTextureADestinationCoefficientsTextureB(commandBuffer metal.MTLCommandBuffer, sourceTexture metal.MTLTexture, guidanceTexture metal.MTLTexture, weightsTexture metal.MTLTexture, destinationCoefficientsTextureA metal.MTLTexture, destinationCoefficientsTextureB metal.MTLTexture) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeRegressionToCommandBuffer:sourceTexture:guidanceTexture:weightsTexture:destinationCoefficientsTextureA:destinationCoefficientsTextureB:"), commandBuffer, sourceTexture, guidanceTexture, weightsTexture, destinationCoefficientsTextureA, destinationCoefficientsTextureB)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGuidedFilter/encodeRegression(to:sourceTexture:guidanceTexture:weightsTexture:destinationCoefficientsTexture:)
func (i MPSImageGuidedFilter) EncodeRegressionToCommandBufferSourceTextureGuidanceTextureWeightsTextureDestinationCoefficientsTexture(commandBuffer metal.MTLCommandBuffer, sourceTexture metal.MTLTexture, guidanceTexture metal.MTLTexture, weightsTexture metal.MTLTexture, destinationCoefficientsTexture metal.MTLTexture) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeRegressionToCommandBuffer:sourceTexture:guidanceTexture:weightsTexture:destinationCoefficientsTexture:"), commandBuffer, sourceTexture, guidanceTexture, weightsTexture, destinationCoefficientsTexture)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGuidedFilter/epsilon
func (i MPSImageGuidedFilter) Epsilon() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("epsilon"))
	return rv
}
func (i MPSImageGuidedFilter) SetEpsilon(value float32) {
	objc.Send[struct{}](i.ID, objc.Sel("setEpsilon:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGuidedFilter/kernelDiameter
func (i MPSImageGuidedFilter) KernelDiameter() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("kernelDiameter"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGuidedFilter/reconstructOffset
func (i MPSImageGuidedFilter) ReconstructOffset() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("reconstructOffset"))
	return rv
}
func (i MPSImageGuidedFilter) SetReconstructOffset(value float32) {
	objc.Send[struct{}](i.ID, objc.Sel("setReconstructOffset:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGuidedFilter/reconstructScale
func (i MPSImageGuidedFilter) ReconstructScale() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("reconstructScale"))
	return rv
}
func (i MPSImageGuidedFilter) SetReconstructScale(value float32) {
	objc.Send[struct{}](i.ID, objc.Sel("setReconstructScale:"), value)
}
