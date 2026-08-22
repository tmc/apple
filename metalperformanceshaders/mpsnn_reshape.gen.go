// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReshape] class.
var (
	_MPSNNReshapeClass     MPSNNReshapeClass
	_MPSNNReshapeClassOnce sync.Once
)

func getMPSNNReshapeClass() MPSNNReshapeClass {
	_MPSNNReshapeClassOnce.Do(func() {
		_MPSNNReshapeClass = MPSNNReshapeClass{class: objc.GetClass("MPSNNReshape")}
	})
	return _MPSNNReshapeClass
}

// GetMPSNNReshapeClass returns the class object for MPSNNReshape.
func GetMPSNNReshapeClass() MPSNNReshapeClass {
	return getMPSNNReshapeClass()
}

type MPSNNReshapeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReshapeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReshapeClass) Alloc() MPSNNReshape {
	rv := objc.Send[MPSNNReshape](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The base class for reshape operations.
//
// # Instance Methods
//
//   - [MPSNNReshape.EncodeToCommandBufferSourceImageDestinationStateDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels]
//   - [MPSNNReshape.EncodeToCommandBufferSourceImageReshapedWidthReshapedHeightReshapedFeatureChannels]
//   - [MPSNNReshape.EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels]
//   - [MPSNNReshape.EncodeBatchToCommandBufferSourceImagesReshapedWidthReshapedHeightReshapedFeatureChannels]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshape
type MPSNNReshape struct {
	MPSCNNKernel
}

// MPSNNReshapeFromID constructs a [MPSNNReshape] from an objc.ID.
//
// The base class for reshape operations.
func MPSNNReshapeFromID(id objc.ID) MPSNNReshape {
	return MPSNNReshape{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSNNReshape adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReshape] class.
//
// # Instance Methods
//
//   - [IMPSNNReshape.EncodeToCommandBufferSourceImageDestinationStateDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels]
//   - [IMPSNNReshape.EncodeToCommandBufferSourceImageReshapedWidthReshapedHeightReshapedFeatureChannels]
//   - [IMPSNNReshape.EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels]
//   - [IMPSNNReshape.EncodeBatchToCommandBufferSourceImagesReshapedWidthReshapedHeightReshapedFeatureChannels]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshape
type IMPSNNReshape interface {
	IMPSCNNKernel

	// Topic: Instance Methods

	EncodeToCommandBufferSourceImageDestinationStateDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, outState IMPSState, isTemporary bool, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) IMPSImage
	EncodeToCommandBufferSourceImageReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) IMPSImage
	EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, outStates MPSStateBatch, isTemporary bool, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) MPSImageBatch
	EncodeBatchToCommandBufferSourceImagesReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) MPSImageBatch
}

// Init initializes the instance.
func (r MPSNNReshape) Init() MPSNNReshape {
	rv := objc.Send[MPSNNReshape](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReshape) Autorelease() MPSNNReshape {
	rv := objc.Send[MPSNNReshape](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReshape creates a new MPSNNReshape instance.
func NewMPSNNReshape() MPSNNReshape {
	class := getMPSNNReshapeClass()
	rv := objc.Send[MPSNNReshape](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewReshapeWithCoder(aDecoder foundation.INSCoder) MPSNNReshape {
	instance := getMPSNNReshapeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNReshapeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshape/init(coder:device:)
func NewReshapeWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNReshape {
	instance := getMPSNNReshapeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNReshapeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshape/init(device:)
func NewReshapeWithDevice(device metal.MTLDevice) MPSNNReshape {
	instance := getMPSNNReshapeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReshapeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshape/encode(commandBuffer:sourceImage:destinationState:destinationStateIsTemporary:reshapedWidth:reshapedHeight:reshapedFeatureChannels:)
func (r MPSNNReshape) EncodeToCommandBufferSourceImageDestinationStateDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, outState IMPSState, isTemporary bool, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) IMPSImage {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("encodeToCommandBuffer:sourceImage:destinationState:destinationStateIsTemporary:reshapedWidth:reshapedHeight:reshapedFeatureChannels:"), commandBuffer, sourceImage, outState, isTemporary, reshapedWidth, reshapedHeight, reshapedFeatureChannels)
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshape/encode(commandBuffer:sourceImage:reshapedWidth:reshapedHeight:reshapedFeatureChannels:)
func (r MPSNNReshape) EncodeToCommandBufferSourceImageReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) IMPSImage {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("encodeToCommandBuffer:sourceImage:reshapedWidth:reshapedHeight:reshapedFeatureChannels:"), commandBuffer, sourceImage, reshapedWidth, reshapedHeight, reshapedFeatureChannels)
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshape/encodeBatch(commandBuffer:sourceImages:destinationStates:destinationStateIsTemporary:reshapedWidth:reshapedHeight:reshapedFeatureChannels:)
func (r MPSNNReshape) EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, outStates MPSStateBatch, isTemporary bool, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) MPSImageBatch {
	rv := objc.Send[MPSImageBatch](r.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:destinationStates:destinationStateIsTemporary:reshapedWidth:reshapedHeight:reshapedFeatureChannels:"), commandBuffer, sourceImages, outStates, isTemporary, reshapedWidth, reshapedHeight, reshapedFeatureChannels)
	return MPSImageBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshape/encodeBatch(commandBuffer:sourceImages:reshapedWidth:reshapedHeight:reshapedFeatureChannels:)
func (r MPSNNReshape) EncodeBatchToCommandBufferSourceImagesReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) MPSImageBatch {
	rv := objc.Send[MPSImageBatch](r.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:reshapedWidth:reshapedHeight:reshapedFeatureChannels:"), commandBuffer, sourceImages, reshapedWidth, reshapedHeight, reshapedFeatureChannels)
	return MPSImageBatch(rv)
}
