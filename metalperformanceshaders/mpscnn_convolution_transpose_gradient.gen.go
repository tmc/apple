// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNConvolutionTransposeGradient] class.
var (
	_MPSCNNConvolutionTransposeGradientClass     MPSCNNConvolutionTransposeGradientClass
	_MPSCNNConvolutionTransposeGradientClassOnce sync.Once
)

func getMPSCNNConvolutionTransposeGradientClass() MPSCNNConvolutionTransposeGradientClass {
	_MPSCNNConvolutionTransposeGradientClassOnce.Do(func() {
		_MPSCNNConvolutionTransposeGradientClass = MPSCNNConvolutionTransposeGradientClass{class: objc.GetClass("MPSCNNConvolutionTransposeGradient")}
	})
	return _MPSCNNConvolutionTransposeGradientClass
}

// GetMPSCNNConvolutionTransposeGradientClass returns the class object for MPSCNNConvolutionTransposeGradient.
func GetMPSCNNConvolutionTransposeGradientClass() MPSCNNConvolutionTransposeGradientClass {
	return getMPSCNNConvolutionTransposeGradientClass()
}

type MPSCNNConvolutionTransposeGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNConvolutionTransposeGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNConvolutionTransposeGradientClass) Alloc() MPSCNNConvolutionTransposeGradient {
	rv := objc.Send[MPSCNNConvolutionTransposeGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSCNNConvolutionTransposeGradient.InitWithDeviceWeights]
//
// # Instance Properties
//
//   - [MPSCNNConvolutionTransposeGradient.DataSource]
//   - [MPSCNNConvolutionTransposeGradient.GradientOption]
//   - [MPSCNNConvolutionTransposeGradient.SetGradientOption]
//   - [MPSCNNConvolutionTransposeGradient.Groups]
//   - [MPSCNNConvolutionTransposeGradient.SourceGradientFeatureChannels]
//   - [MPSCNNConvolutionTransposeGradient.SourceImageFeatureChannels]
//
// # Instance Methods
//
//   - [MPSCNNConvolutionTransposeGradient.ReloadWeightsAndBiasesWithCommandBufferState]
//   - [MPSCNNConvolutionTransposeGradient.ReloadWeightsAndBiasesFromDataSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradient
type MPSCNNConvolutionTransposeGradient struct {
	MPSCNNGradientKernel
}

// MPSCNNConvolutionTransposeGradientFromID constructs a [MPSCNNConvolutionTransposeGradient] from an objc.ID.
func MPSCNNConvolutionTransposeGradientFromID(id objc.ID) MPSCNNConvolutionTransposeGradient {
	return MPSCNNConvolutionTransposeGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}

// NOTE: MPSCNNConvolutionTransposeGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNConvolutionTransposeGradient] class.
//
// # Initializers
//
//   - [IMPSCNNConvolutionTransposeGradient.InitWithDeviceWeights]
//
// # Instance Properties
//
//   - [IMPSCNNConvolutionTransposeGradient.DataSource]
//   - [IMPSCNNConvolutionTransposeGradient.GradientOption]
//   - [IMPSCNNConvolutionTransposeGradient.SetGradientOption]
//   - [IMPSCNNConvolutionTransposeGradient.Groups]
//   - [IMPSCNNConvolutionTransposeGradient.SourceGradientFeatureChannels]
//   - [IMPSCNNConvolutionTransposeGradient.SourceImageFeatureChannels]
//
// # Instance Methods
//
//   - [IMPSCNNConvolutionTransposeGradient.ReloadWeightsAndBiasesWithCommandBufferState]
//   - [IMPSCNNConvolutionTransposeGradient.ReloadWeightsAndBiasesFromDataSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradient
type IMPSCNNConvolutionTransposeGradient interface {
	IMPSCNNGradientKernel

	// Topic: Initializers

	InitWithDeviceWeights(device metal.MTLDevice, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionTransposeGradient

	// Topic: Instance Properties

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
func (c MPSCNNConvolutionTransposeGradient) Init() MPSCNNConvolutionTransposeGradient {
	rv := objc.Send[MPSCNNConvolutionTransposeGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNConvolutionTransposeGradient) Autorelease() MPSCNNConvolutionTransposeGradient {
	rv := objc.Send[MPSCNNConvolutionTransposeGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNConvolutionTransposeGradient creates a new MPSCNNConvolutionTransposeGradient instance.
func NewMPSCNNConvolutionTransposeGradient() MPSCNNConvolutionTransposeGradient {
	class := getMPSCNNConvolutionTransposeGradientClass()
	rv := objc.Send[MPSCNNConvolutionTransposeGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNConvolutionTransposeGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNConvolutionTransposeGradient {
	instance := getMPSCNNConvolutionTransposeGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNConvolutionTransposeGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradient/init(coder:device:)
func NewCNNConvolutionTransposeGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNConvolutionTransposeGradient {
	instance := getMPSCNNConvolutionTransposeGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNConvolutionTransposeGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNConvolutionTransposeGradientWithDevice(device metal.MTLDevice) MPSCNNConvolutionTransposeGradient {
	instance := getMPSCNNConvolutionTransposeGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNConvolutionTransposeGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradient/init(device:weights:)
func NewCNNConvolutionTransposeGradientWithDeviceWeights(device metal.MTLDevice, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionTransposeGradient {
	instance := getMPSCNNConvolutionTransposeGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:weights:"), device, weights)
	return MPSCNNConvolutionTransposeGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradient/init(device:weights:)
func (c MPSCNNConvolutionTransposeGradient) InitWithDeviceWeights(device metal.MTLDevice, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionTransposeGradient {
	rv := objc.Send[MPSCNNConvolutionTransposeGradient](c.ID, objc.Sel("initWithDevice:weights:"), device, weights)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradient/reloadWeightsAndBiases(with:state:)
func (c MPSCNNConvolutionTransposeGradient) ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer metal.MTLCommandBuffer, state IMPSCNNConvolutionWeightsAndBiasesState) {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadWeightsAndBiasesWithCommandBuffer:state:"), commandBuffer, state)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradient/reloadWeightsAndBiasesFromDataSource()
func (c MPSCNNConvolutionTransposeGradient) ReloadWeightsAndBiasesFromDataSource() {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadWeightsAndBiasesFromDataSource"))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradient/dataSource
func (c MPSCNNConvolutionTransposeGradient) DataSource() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dataSource"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradient/gradientOption
func (c MPSCNNConvolutionTransposeGradient) GradientOption() MPSCNNConvolutionGradientOption {
	rv := objc.Send[MPSCNNConvolutionGradientOption](c.ID, objc.Sel("gradientOption"))
	return MPSCNNConvolutionGradientOption(rv)
}
func (c MPSCNNConvolutionTransposeGradient) SetGradientOption(value MPSCNNConvolutionGradientOption) {
	objc.Send[struct{}](c.ID, objc.Sel("setGradientOption:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradient/groups
func (c MPSCNNConvolutionTransposeGradient) Groups() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("groups"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradient/sourceGradientFeatureChannels
func (c MPSCNNConvolutionTransposeGradient) SourceGradientFeatureChannels() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("sourceGradientFeatureChannels"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradient/sourceImageFeatureChannels
func (c MPSCNNConvolutionTransposeGradient) SourceImageFeatureChannels() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("sourceImageFeatureChannels"))
	return rv
}
