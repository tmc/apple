// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNArithmetic] class.
var (
	_MPSCNNArithmeticClass     MPSCNNArithmeticClass
	_MPSCNNArithmeticClassOnce sync.Once
)

func getMPSCNNArithmeticClass() MPSCNNArithmeticClass {
	_MPSCNNArithmeticClassOnce.Do(func() {
		_MPSCNNArithmeticClass = MPSCNNArithmeticClass{class: objc.GetClass("MPSCNNArithmetic")}
	})
	return _MPSCNNArithmeticClass
}

// GetMPSCNNArithmeticClass returns the class object for MPSCNNArithmetic.
func GetMPSCNNArithmeticClass() MPSCNNArithmeticClass {
	return getMPSCNNArithmeticClass()
}

type MPSCNNArithmeticClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNArithmeticClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNArithmeticClass) Alloc() MPSCNNArithmetic {
	rv := objc.Send[MPSCNNArithmetic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The base class for arithmetic operators.
//
// # Instance Properties
//
//   - [MPSCNNArithmetic.Bias]
//   - [MPSCNNArithmetic.SetBias]
//   - [MPSCNNArithmetic.MaximumValue]
//   - [MPSCNNArithmetic.SetMaximumValue]
//   - [MPSCNNArithmetic.MinimumValue]
//   - [MPSCNNArithmetic.SetMinimumValue]
//   - [MPSCNNArithmetic.PrimaryScale]
//   - [MPSCNNArithmetic.SetPrimaryScale]
//   - [MPSCNNArithmetic.PrimaryStrideInFeatureChannels]
//   - [MPSCNNArithmetic.SetPrimaryStrideInFeatureChannels]
//   - [MPSCNNArithmetic.SecondaryScale]
//   - [MPSCNNArithmetic.SetSecondaryScale]
//   - [MPSCNNArithmetic.SecondaryStrideInFeatureChannels]
//   - [MPSCNNArithmetic.SetSecondaryStrideInFeatureChannels]
//
// # Instance Methods
//
//   - [MPSCNNArithmetic.EncodeToCommandBufferPrimaryImageSecondaryImageDestinationStateDestinationImage]
//   - [MPSCNNArithmetic.EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationStatesDestinationImages]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmetic
type MPSCNNArithmetic struct {
	MPSCNNBinaryKernel
}

// MPSCNNArithmeticFromID constructs a [MPSCNNArithmetic] from an objc.ID.
//
// The base class for arithmetic operators.
func MPSCNNArithmeticFromID(id objc.ID) MPSCNNArithmetic {
	return MPSCNNArithmetic{MPSCNNBinaryKernel: MPSCNNBinaryKernelFromID(id)}
}

// NOTE: MPSCNNArithmetic adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNArithmetic] class.
//
// # Instance Properties
//
//   - [IMPSCNNArithmetic.Bias]
//   - [IMPSCNNArithmetic.SetBias]
//   - [IMPSCNNArithmetic.MaximumValue]
//   - [IMPSCNNArithmetic.SetMaximumValue]
//   - [IMPSCNNArithmetic.MinimumValue]
//   - [IMPSCNNArithmetic.SetMinimumValue]
//   - [IMPSCNNArithmetic.PrimaryScale]
//   - [IMPSCNNArithmetic.SetPrimaryScale]
//   - [IMPSCNNArithmetic.PrimaryStrideInFeatureChannels]
//   - [IMPSCNNArithmetic.SetPrimaryStrideInFeatureChannels]
//   - [IMPSCNNArithmetic.SecondaryScale]
//   - [IMPSCNNArithmetic.SetSecondaryScale]
//   - [IMPSCNNArithmetic.SecondaryStrideInFeatureChannels]
//   - [IMPSCNNArithmetic.SetSecondaryStrideInFeatureChannels]
//
// # Instance Methods
//
//   - [IMPSCNNArithmetic.EncodeToCommandBufferPrimaryImageSecondaryImageDestinationStateDestinationImage]
//   - [IMPSCNNArithmetic.EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationStatesDestinationImages]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmetic
type IMPSCNNArithmetic interface {
	IMPSCNNBinaryKernel

	// Topic: Instance Properties

	Bias() float32
	SetBias(value float32)
	MaximumValue() float32
	SetMaximumValue(value float32)
	MinimumValue() float32
	SetMinimumValue(value float32)
	PrimaryScale() float32
	SetPrimaryScale(value float32)
	PrimaryStrideInFeatureChannels() uint
	SetPrimaryStrideInFeatureChannels(value uint)
	SecondaryScale() float32
	SetSecondaryScale(value float32)
	SecondaryStrideInFeatureChannels() uint
	SetSecondaryStrideInFeatureChannels(value uint)

	// Topic: Instance Methods

	EncodeToCommandBufferPrimaryImageSecondaryImageDestinationStateDestinationImage(commandBuffer metal.MTLCommandBuffer, primaryImage IMPSImage, secondaryImage IMPSImage, destinationState IMPSCNNArithmeticGradientState, destinationImage IMPSImage)
	EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationStatesDestinationImages(commandBuffer metal.MTLCommandBuffer, primaryImages MPSImageBatch, secondaryImages MPSImageBatch, destinationStates MPSCNNArithmeticGradientStateBatch, destinationImages MPSImageBatch)
}

// Init initializes the instance.
func (c MPSCNNArithmetic) Init() MPSCNNArithmetic {
	rv := objc.Send[MPSCNNArithmetic](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNArithmetic) Autorelease() MPSCNNArithmetic {
	rv := objc.Send[MPSCNNArithmetic](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNArithmetic creates a new MPSCNNArithmetic instance.
func NewMPSCNNArithmetic() MPSCNNArithmetic {
	class := getMPSCNNArithmeticClass()
	rv := objc.Send[MPSCNNArithmetic](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNArithmeticWithCoder(aDecoder foundation.INSCoder) MPSCNNArithmetic {
	instance := getMPSCNNArithmeticClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNArithmeticFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/init(coder:device:)
func NewCNNArithmeticWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNArithmetic {
	instance := getMPSCNNArithmeticClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNArithmeticFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/init(device:)
func NewCNNArithmeticWithDevice(device metal.MTLDevice) MPSCNNArithmetic {
	instance := getMPSCNNArithmeticClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNArithmeticFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmetic/encode(commandBuffer:primaryImage:secondaryImage:destinationState:destinationImage:)
func (c MPSCNNArithmetic) EncodeToCommandBufferPrimaryImageSecondaryImageDestinationStateDestinationImage(commandBuffer metal.MTLCommandBuffer, primaryImage IMPSImage, secondaryImage IMPSImage, destinationState IMPSCNNArithmeticGradientState, destinationImage IMPSImage) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:primaryImage:secondaryImage:destinationState:destinationImage:"), commandBuffer, primaryImage, secondaryImage, destinationState, destinationImage)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmetic/encodeBatch(commandBuffer:primaryImages:secondaryImages:destinationStates:destinationImages:)
func (c MPSCNNArithmetic) EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationStatesDestinationImages(commandBuffer metal.MTLCommandBuffer, primaryImages MPSImageBatch, secondaryImages MPSImageBatch, destinationStates MPSCNNArithmeticGradientStateBatch, destinationImages MPSImageBatch) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeBatchToCommandBuffer:primaryImages:secondaryImages:destinationStates:destinationImages:"), commandBuffer, primaryImages, secondaryImages, destinationStates, destinationImages)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmetic/bias
func (c MPSCNNArithmetic) Bias() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("bias"))
	return rv
}
func (c MPSCNNArithmetic) SetBias(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setBias:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmetic/maximumValue
func (c MPSCNNArithmetic) MaximumValue() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("maximumValue"))
	return rv
}
func (c MPSCNNArithmetic) SetMaximumValue(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaximumValue:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmetic/minimumValue
func (c MPSCNNArithmetic) MinimumValue() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("minimumValue"))
	return rv
}
func (c MPSCNNArithmetic) SetMinimumValue(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setMinimumValue:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmetic/primaryScale
func (c MPSCNNArithmetic) PrimaryScale() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("primaryScale"))
	return rv
}
func (c MPSCNNArithmetic) SetPrimaryScale(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setPrimaryScale:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmetic/primaryStrideInFeatureChannels
func (c MPSCNNArithmetic) PrimaryStrideInFeatureChannels() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("primaryStrideInFeatureChannels"))
	return rv
}
func (c MPSCNNArithmetic) SetPrimaryStrideInFeatureChannels(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setPrimaryStrideInFeatureChannels:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmetic/secondaryScale
func (c MPSCNNArithmetic) SecondaryScale() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("secondaryScale"))
	return rv
}
func (c MPSCNNArithmetic) SetSecondaryScale(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setSecondaryScale:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmetic/secondaryStrideInFeatureChannels
func (c MPSCNNArithmetic) SecondaryStrideInFeatureChannels() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("secondaryStrideInFeatureChannels"))
	return rv
}
func (c MPSCNNArithmetic) SetSecondaryStrideInFeatureChannels(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setSecondaryStrideInFeatureChannels:"), value)
}
