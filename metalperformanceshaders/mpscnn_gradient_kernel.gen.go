// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNGradientKernel] class.
var (
	_MPSCNNGradientKernelClass     MPSCNNGradientKernelClass
	_MPSCNNGradientKernelClassOnce sync.Once
)

func getMPSCNNGradientKernelClass() MPSCNNGradientKernelClass {
	_MPSCNNGradientKernelClassOnce.Do(func() {
		_MPSCNNGradientKernelClass = MPSCNNGradientKernelClass{class: objc.GetClass("MPSCNNGradientKernel")}
	})
	return _MPSCNNGradientKernelClass
}

// GetMPSCNNGradientKernelClass returns the class object for MPSCNNGradientKernel.
func GetMPSCNNGradientKernelClass() MPSCNNGradientKernelClass {
	return getMPSCNNGradientKernelClass()
}

type MPSCNNGradientKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNGradientKernelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNGradientKernelClass) Alloc() MPSCNNGradientKernel {
	rv := objc.Send[MPSCNNGradientKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The base class for gradient layers.
//
// # Instance Properties
//
//   - [MPSCNNGradientKernel.KernelOffsetX]
//   - [MPSCNNGradientKernel.SetKernelOffsetX]
//   - [MPSCNNGradientKernel.KernelOffsetY]
//   - [MPSCNNGradientKernel.SetKernelOffsetY]
//
// # Instance Methods
//
//   - [MPSCNNGradientKernel.EncodeToCommandBufferSourceGradientSourceImageGradientState]
//   - [MPSCNNGradientKernel.EncodeToCommandBufferSourceGradientSourceImageGradientStateDestinationGradient]
//   - [MPSCNNGradientKernel.EncodeBatchToCommandBufferSourceGradientsSourceImagesGradientStates]
//   - [MPSCNNGradientKernel.EncodeBatchToCommandBufferSourceGradientsSourceImagesGradientStatesDestinationGradients]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel
type MPSCNNGradientKernel struct {
	MPSCNNBinaryKernel
}

// MPSCNNGradientKernelFromID constructs a [MPSCNNGradientKernel] from an objc.ID.
//
// The base class for gradient layers.
func MPSCNNGradientKernelFromID(id objc.ID) MPSCNNGradientKernel {
	return MPSCNNGradientKernel{MPSCNNBinaryKernel: MPSCNNBinaryKernelFromID(id)}
}

// NOTE: MPSCNNGradientKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNGradientKernel] class.
//
// # Instance Properties
//
//   - [IMPSCNNGradientKernel.KernelOffsetX]
//   - [IMPSCNNGradientKernel.SetKernelOffsetX]
//   - [IMPSCNNGradientKernel.KernelOffsetY]
//   - [IMPSCNNGradientKernel.SetKernelOffsetY]
//
// # Instance Methods
//
//   - [IMPSCNNGradientKernel.EncodeToCommandBufferSourceGradientSourceImageGradientState]
//   - [IMPSCNNGradientKernel.EncodeToCommandBufferSourceGradientSourceImageGradientStateDestinationGradient]
//   - [IMPSCNNGradientKernel.EncodeBatchToCommandBufferSourceGradientsSourceImagesGradientStates]
//   - [IMPSCNNGradientKernel.EncodeBatchToCommandBufferSourceGradientsSourceImagesGradientStatesDestinationGradients]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel
type IMPSCNNGradientKernel interface {
	IMPSCNNBinaryKernel

	// Topic: Instance Properties

	KernelOffsetX() int
	SetKernelOffsetX(value int)
	KernelOffsetY() int
	SetKernelOffsetY(value int)

	// Topic: Instance Methods

	EncodeToCommandBufferSourceGradientSourceImageGradientState(commandBuffer metal.MTLCommandBuffer, sourceGradient IMPSImage, sourceImage IMPSImage, gradientState IMPSState) IMPSImage
	EncodeToCommandBufferSourceGradientSourceImageGradientStateDestinationGradient(commandBuffer metal.MTLCommandBuffer, sourceGradient IMPSImage, sourceImage IMPSImage, gradientState IMPSState, destinationGradient IMPSImage)
	EncodeBatchToCommandBufferSourceGradientsSourceImagesGradientStates(commandBuffer metal.MTLCommandBuffer, sourceGradients MPSImageBatch, sourceImages MPSImageBatch, gradientStates MPSStateBatch) MPSImageBatch
	EncodeBatchToCommandBufferSourceGradientsSourceImagesGradientStatesDestinationGradients(commandBuffer metal.MTLCommandBuffer, sourceGradients MPSImageBatch, sourceImages MPSImageBatch, gradientStates MPSStateBatch, destinationGradients MPSImageBatch)
}

// Init initializes the instance.
func (c MPSCNNGradientKernel) Init() MPSCNNGradientKernel {
	rv := objc.Send[MPSCNNGradientKernel](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNGradientKernel) Autorelease() MPSCNNGradientKernel {
	rv := objc.Send[MPSCNNGradientKernel](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNGradientKernel creates a new MPSCNNGradientKernel instance.
func NewMPSCNNGradientKernel() MPSCNNGradientKernel {
	class := getMPSCNNGradientKernelClass()
	rv := objc.Send[MPSCNNGradientKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNGradientKernelWithCoder(aDecoder foundation.INSCoder) MPSCNNGradientKernel {
	instance := getMPSCNNGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(coder:device:)
func NewCNNGradientKernelWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNGradientKernel {
	instance := getMPSCNNGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNGradientKernelWithDevice(device metal.MTLDevice) MPSCNNGradientKernel {
	instance := getMPSCNNGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/encode(commandBuffer:sourceGradient:sourceImage:gradientState:)
func (c MPSCNNGradientKernel) EncodeToCommandBufferSourceGradientSourceImageGradientState(commandBuffer metal.MTLCommandBuffer, sourceGradient IMPSImage, sourceImage IMPSImage, gradientState IMPSState) IMPSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:sourceGradient:sourceImage:gradientState:"), commandBuffer, sourceGradient, sourceImage, gradientState)
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/encode(commandBuffer:sourceGradient:sourceImage:gradientState:destinationGradient:)
func (c MPSCNNGradientKernel) EncodeToCommandBufferSourceGradientSourceImageGradientStateDestinationGradient(commandBuffer metal.MTLCommandBuffer, sourceGradient IMPSImage, sourceImage IMPSImage, gradientState IMPSState, destinationGradient IMPSImage) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:sourceGradient:sourceImage:gradientState:destinationGradient:"), commandBuffer, sourceGradient, sourceImage, gradientState, destinationGradient)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/encodeBatch(commandBuffer:sourceGradients:sourceImages:gradientStates:)
func (c MPSCNNGradientKernel) EncodeBatchToCommandBufferSourceGradientsSourceImagesGradientStates(commandBuffer metal.MTLCommandBuffer, sourceGradients MPSImageBatch, sourceImages MPSImageBatch, gradientStates MPSStateBatch) MPSImageBatch {
	rv := objc.Send[MPSImageBatch](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:gradientStates:"), commandBuffer, sourceGradients, sourceImages, gradientStates)
	return MPSImageBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/encodeBatch(commandBuffer:sourceGradients:sourceImages:gradientStates:destinationGradients:)
func (c MPSCNNGradientKernel) EncodeBatchToCommandBufferSourceGradientsSourceImagesGradientStatesDestinationGradients(commandBuffer metal.MTLCommandBuffer, sourceGradients MPSImageBatch, sourceImages MPSImageBatch, gradientStates MPSStateBatch, destinationGradients MPSImageBatch) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:gradientStates:destinationGradients:"), commandBuffer, sourceGradients, sourceImages, gradientStates, destinationGradients)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/kernelOffsetX
func (c MPSCNNGradientKernel) KernelOffsetX() int {
	rv := objc.Send[int](c.ID, objc.Sel("kernelOffsetX"))
	return rv
}
func (c MPSCNNGradientKernel) SetKernelOffsetX(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setKernelOffsetX:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/kernelOffsetY
func (c MPSCNNGradientKernel) KernelOffsetY() int {
	rv := objc.Send[int](c.ID, objc.Sel("kernelOffsetY"))
	return rv
}
func (c MPSCNNGradientKernel) SetKernelOffsetY(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setKernelOffsetY:"), value)
}
