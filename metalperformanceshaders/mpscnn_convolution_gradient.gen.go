// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNConvolutionGradient] class.
var (
	_MPSCNNConvolutionGradientClass     MPSCNNConvolutionGradientClass
	_MPSCNNConvolutionGradientClassOnce sync.Once
)

func getMPSCNNConvolutionGradientClass() MPSCNNConvolutionGradientClass {
	_MPSCNNConvolutionGradientClassOnce.Do(func() {
		_MPSCNNConvolutionGradientClass = MPSCNNConvolutionGradientClass{class: objc.GetClass("MPSCNNConvolutionGradient")}
	})
	return _MPSCNNConvolutionGradientClass
}

// GetMPSCNNConvolutionGradientClass returns the class object for MPSCNNConvolutionGradient.
func GetMPSCNNConvolutionGradientClass() MPSCNNConvolutionGradientClass {
	return getMPSCNNConvolutionGradientClass()
}

type MPSCNNConvolutionGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNConvolutionGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNConvolutionGradientClass) Alloc() MPSCNNConvolutionGradient {
	rv := objc.Send[MPSCNNConvolutionGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient convolution kernel.
//
// # Initializers
//
//   - [MPSCNNConvolutionGradient.InitWithDeviceWeights]
//
// # Instance Properties
//
//   - [MPSCNNConvolutionGradient.ChannelMultiplier]
//   - [MPSCNNConvolutionGradient.DataSource]
//   - [MPSCNNConvolutionGradient.GradientOption]
//   - [MPSCNNConvolutionGradient.SetGradientOption]
//   - [MPSCNNConvolutionGradient.Groups]
//   - [MPSCNNConvolutionGradient.SourceGradientFeatureChannels]
//   - [MPSCNNConvolutionGradient.SourceImageFeatureChannels]
//
// # Instance Methods
//
//   - [MPSCNNConvolutionGradient.ReloadWeightsAndBiasesWithCommandBufferState]
//   - [MPSCNNConvolutionGradient.ReloadWeightsAndBiasesFromDataSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradient
type MPSCNNConvolutionGradient struct {
	MPSCNNGradientKernel
}

// MPSCNNConvolutionGradientFromID constructs a [MPSCNNConvolutionGradient] from an objc.ID.
//
// A gradient convolution kernel.
func MPSCNNConvolutionGradientFromID(id objc.ID) MPSCNNConvolutionGradient {
	return MPSCNNConvolutionGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}

// NOTE: MPSCNNConvolutionGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNConvolutionGradient] class.
//
// # Initializers
//
//   - [IMPSCNNConvolutionGradient.InitWithDeviceWeights]
//
// # Instance Properties
//
//   - [IMPSCNNConvolutionGradient.ChannelMultiplier]
//   - [IMPSCNNConvolutionGradient.DataSource]
//   - [IMPSCNNConvolutionGradient.GradientOption]
//   - [IMPSCNNConvolutionGradient.SetGradientOption]
//   - [IMPSCNNConvolutionGradient.Groups]
//   - [IMPSCNNConvolutionGradient.SourceGradientFeatureChannels]
//   - [IMPSCNNConvolutionGradient.SourceImageFeatureChannels]
//
// # Instance Methods
//
//   - [IMPSCNNConvolutionGradient.ReloadWeightsAndBiasesWithCommandBufferState]
//   - [IMPSCNNConvolutionGradient.ReloadWeightsAndBiasesFromDataSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradient
type IMPSCNNConvolutionGradient interface {
	IMPSCNNGradientKernel

	// Topic: Initializers

	InitWithDeviceWeights(device metal.MTLDevice, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionGradient

	// Topic: Instance Properties

	ChannelMultiplier() uint
	DataSource() MPSCNNConvolutionDataSource
	GradientOption() MPSCNNConvolutionGradientOption
	SetGradientOption(value MPSCNNConvolutionGradientOption)
	Groups() uint
	SourceGradientFeatureChannels() uint
	SourceImageFeatureChannels() uint

	// Topic: Instance Methods

	ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer metal.MTLCommandBuffer, state IMPSCNNConvolutionWeightsAndBiasesState)
	ReloadWeightsAndBiasesFromDataSource()
}

// Init initializes the instance.
func (c MPSCNNConvolutionGradient) Init() MPSCNNConvolutionGradient {
	rv := objc.Send[MPSCNNConvolutionGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNConvolutionGradient) Autorelease() MPSCNNConvolutionGradient {
	rv := objc.Send[MPSCNNConvolutionGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNConvolutionGradient creates a new MPSCNNConvolutionGradient instance.
func NewMPSCNNConvolutionGradient() MPSCNNConvolutionGradient {
	class := getMPSCNNConvolutionGradientClass()
	rv := objc.Send[MPSCNNConvolutionGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNConvolutionGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNConvolutionGradient {
	instance := getMPSCNNConvolutionGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNConvolutionGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradient/init(coder:device:)
func NewCNNConvolutionGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNConvolutionGradient {
	instance := getMPSCNNConvolutionGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNConvolutionGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNConvolutionGradientWithDevice(device metal.MTLDevice) MPSCNNConvolutionGradient {
	instance := getMPSCNNConvolutionGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNConvolutionGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradient/init(device:weights:)
func NewCNNConvolutionGradientWithDeviceWeights(device metal.MTLDevice, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionGradient {
	instance := getMPSCNNConvolutionGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:weights:"), device, weights)
	return MPSCNNConvolutionGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradient/init(device:weights:)
func (c MPSCNNConvolutionGradient) InitWithDeviceWeights(device metal.MTLDevice, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionGradient {
	rv := objc.Send[MPSCNNConvolutionGradient](c.ID, objc.Sel("initWithDevice:weights:"), device, weights)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradient/reloadWeightsAndBiases(with:state:)
func (c MPSCNNConvolutionGradient) ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer metal.MTLCommandBuffer, state IMPSCNNConvolutionWeightsAndBiasesState) {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadWeightsAndBiasesWithCommandBuffer:state:"), commandBuffer, state)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradient/reloadWeightsAndBiasesFromDataSource()
func (c MPSCNNConvolutionGradient) ReloadWeightsAndBiasesFromDataSource() {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadWeightsAndBiasesFromDataSource"))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradient/channelMultiplier
func (c MPSCNNConvolutionGradient) ChannelMultiplier() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("channelMultiplier"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradient/dataSource
func (c MPSCNNConvolutionGradient) DataSource() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dataSource"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradient/gradientOption
func (c MPSCNNConvolutionGradient) GradientOption() MPSCNNConvolutionGradientOption {
	rv := objc.Send[MPSCNNConvolutionGradientOption](c.ID, objc.Sel("gradientOption"))
	return MPSCNNConvolutionGradientOption(rv)
}
func (c MPSCNNConvolutionGradient) SetGradientOption(value MPSCNNConvolutionGradientOption) {
	objc.Send[struct{}](c.ID, objc.Sel("setGradientOption:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradient/groups
func (c MPSCNNConvolutionGradient) Groups() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("groups"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradient/sourceGradientFeatureChannels
func (c MPSCNNConvolutionGradient) SourceGradientFeatureChannels() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("sourceGradientFeatureChannels"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradient/sourceImageFeatureChannels
func (c MPSCNNConvolutionGradient) SourceImageFeatureChannels() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("sourceImageFeatureChannels"))
	return rv
}
