// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNConvolution] class.
var (
	_MPSCNNConvolutionClass     MPSCNNConvolutionClass
	_MPSCNNConvolutionClassOnce sync.Once
)

func getMPSCNNConvolutionClass() MPSCNNConvolutionClass {
	_MPSCNNConvolutionClassOnce.Do(func() {
		_MPSCNNConvolutionClass = MPSCNNConvolutionClass{class: objc.GetClass("MPSCNNConvolution")}
	})
	return _MPSCNNConvolutionClass
}

// GetMPSCNNConvolutionClass returns the class object for MPSCNNConvolution.
func GetMPSCNNConvolutionClass() MPSCNNConvolutionClass {
	return getMPSCNNConvolutionClass()
}

type MPSCNNConvolutionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNConvolutionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNConvolutionClass) Alloc() MPSCNNConvolution {
	rv := objc.Send[MPSCNNConvolution](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A convolution kernel that convolves the input image with a set of filters,
// with each producing one feature map in the output image.
//
// # Overview
//
// The attributes of a convolution operation are described by an
// [MPSCNNConvolutionDescriptor] object.
//
// # Initializers
//
//   - [MPSCNNConvolution.InitWithDeviceWeights]
//
// # Instance Properties
//
//   - [MPSCNNConvolution.InputFeatureChannels]: The number of feature channels per pixel in the input image.
//   - [MPSCNNConvolution.OutputFeatureChannels]: The number of feature channels per pixel in the output image.
//   - [MPSCNNConvolution.Groups]: The number of groups that the input and output channels are divided into.
//   - [MPSCNNConvolution.SubPixelScaleFactor]
//   - [MPSCNNConvolution.Neuron]: The neuron filter to be applied as part of the convolution operation.
//   - [MPSCNNConvolution.AccumulatorPrecisionOption]
//   - [MPSCNNConvolution.SetAccumulatorPrecisionOption]
//   - [MPSCNNConvolution.ChannelMultiplier]
//   - [MPSCNNConvolution.DataSource]
//   - [MPSCNNConvolution.FusedNeuronDescriptor]
//
// # Instance Methods
//
//   - [MPSCNNConvolution.ExportWeightsAndBiasesWithCommandBufferResultStateCanBeTemporary]
//   - [MPSCNNConvolution.ReloadWeightsAndBiasesWithCommandBufferState]
//   - [MPSCNNConvolution.ReloadWeightsAndBiasesFromDataSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution
type MPSCNNConvolution struct {
	MPSCNNKernel
}

// MPSCNNConvolutionFromID constructs a [MPSCNNConvolution] from an objc.ID.
//
// A convolution kernel that convolves the input image with a set of filters,
// with each producing one feature map in the output image.
func MPSCNNConvolutionFromID(id objc.ID) MPSCNNConvolution {
	return MPSCNNConvolution{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSCNNConvolution adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNConvolution] class.
//
// # Initializers
//
//   - [IMPSCNNConvolution.InitWithDeviceWeights]
//
// # Instance Properties
//
//   - [IMPSCNNConvolution.InputFeatureChannels]: The number of feature channels per pixel in the input image.
//   - [IMPSCNNConvolution.OutputFeatureChannels]: The number of feature channels per pixel in the output image.
//   - [IMPSCNNConvolution.Groups]: The number of groups that the input and output channels are divided into.
//   - [IMPSCNNConvolution.SubPixelScaleFactor]
//   - [IMPSCNNConvolution.Neuron]: The neuron filter to be applied as part of the convolution operation.
//   - [IMPSCNNConvolution.AccumulatorPrecisionOption]
//   - [IMPSCNNConvolution.SetAccumulatorPrecisionOption]
//   - [IMPSCNNConvolution.ChannelMultiplier]
//   - [IMPSCNNConvolution.DataSource]
//   - [IMPSCNNConvolution.FusedNeuronDescriptor]
//
// # Instance Methods
//
//   - [IMPSCNNConvolution.ExportWeightsAndBiasesWithCommandBufferResultStateCanBeTemporary]
//   - [IMPSCNNConvolution.ReloadWeightsAndBiasesWithCommandBufferState]
//   - [IMPSCNNConvolution.ReloadWeightsAndBiasesFromDataSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution
type IMPSCNNConvolution interface {
	IMPSCNNKernel

	// Topic: Initializers

	InitWithDeviceWeights(device metal.MTLDevice, weights MPSCNNConvolutionDataSource) MPSCNNConvolution

	// Topic: Instance Properties

	// The number of feature channels per pixel in the input image.
	InputFeatureChannels() uint
	// The number of feature channels per pixel in the output image.
	OutputFeatureChannels() uint
	// The number of groups that the input and output channels are divided into.
	Groups() uint
	SubPixelScaleFactor() uint
	// The neuron filter to be applied as part of the convolution operation.
	Neuron() IMPSCNNNeuron
	AccumulatorPrecisionOption() MPSNNConvolutionAccumulatorPrecisionOption
	SetAccumulatorPrecisionOption(value MPSNNConvolutionAccumulatorPrecisionOption)
	ChannelMultiplier() uint
	DataSource() MPSCNNConvolutionDataSource
	FusedNeuronDescriptor() IMPSNNNeuronDescriptor

	// Topic: Instance Methods

	ExportWeightsAndBiasesWithCommandBufferResultStateCanBeTemporary(commandBuffer metal.MTLCommandBuffer, resultStateCanBeTemporary bool) IMPSCNNConvolutionWeightsAndBiasesState
	ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer metal.MTLCommandBuffer, state IMPSCNNConvolutionWeightsAndBiasesState)
	ReloadWeightsAndBiasesFromDataSource()
}

// Init initializes the instance.
func (c MPSCNNConvolution) Init() MPSCNNConvolution {
	rv := objc.Send[MPSCNNConvolution](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNConvolution) Autorelease() MPSCNNConvolution {
	rv := objc.Send[MPSCNNConvolution](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNConvolution creates a new MPSCNNConvolution instance.
func NewMPSCNNConvolution() MPSCNNConvolution {
	class := getMPSCNNConvolutionClass()
	rv := objc.Send[MPSCNNConvolution](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNConvolutionWithCoder(aDecoder foundation.INSCoder) MPSCNNConvolution {
	instance := getMPSCNNConvolutionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNConvolutionFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution/init(coder:device:)
func NewCNNConvolutionWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNConvolution {
	instance := getMPSCNNConvolutionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNConvolutionFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNConvolutionWithDevice(device metal.MTLDevice) MPSCNNConvolution {
	instance := getMPSCNNConvolutionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNConvolutionFromID(rv)
}

// Initializes a convolution kernel.
//
// device: The device on which this kernel will run.
//
// convolutionDescriptor: A pointer to a valid convolution descriptor.
//
// kernelWeights: A pointer to a weights array.
//
// Each entry is a float value. The number of entries is equal to
// [MPSCNNConvolutionDescriptor.InputFeatureChannels] `*`
// [MPSCNNConvolutionDescriptor.OutputFeatureChannels] `*`
// [MPSCNNConvolutionDescriptor.KernelHeight] `*`
// [MPSCNNConvolutionDescriptor.KernelWidth].
//
// The layout of the filter weight is arranged so that it can be reinterpreted
// as a 4D tensor (array)
// `weight[outputChannels][kernelHeight][kernelWidth][inputChannels/groups]`
//
// Weights are converted to half float precision (`fp16`) internally for best
// performance.
//
// biasTerms: A pointer to bias terms to be applied to the convolution output.
//
// Each entry is a float value. The number of entries is the number of output
// feature maps.
//
// flags: Currently unused.
//
// This value must be [MPSCNNConvolutionFlagsNone].
//
// # Return Value
//
// A valid [MPSCNNConvolution] object or `nil`, if failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution/init(device:convolutionDescriptor:kernelWeights:biasTerms:flags:)
func NewCNNConvolutionWithDeviceConvolutionDescriptorKernelWeightsBiasTermsFlags(device metal.MTLDevice, convolutionDescriptor IMPSCNNConvolutionDescriptor, kernelWeights *float32, biasTerms *float32, flags MPSCNNConvolutionFlags) MPSCNNConvolution {
	instance := getMPSCNNConvolutionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:convolutionDescriptor:kernelWeights:biasTerms:flags:"), device, convolutionDescriptor, kernelWeights, biasTerms, flags)
	return MPSCNNConvolutionFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution/init(device:weights:)
func NewCNNConvolutionWithDeviceWeights(device metal.MTLDevice, weights MPSCNNConvolutionDataSource) MPSCNNConvolution {
	instance := getMPSCNNConvolutionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:weights:"), device, weights)
	return MPSCNNConvolutionFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution/init(device:weights:)
func (c MPSCNNConvolution) InitWithDeviceWeights(device metal.MTLDevice, weights MPSCNNConvolutionDataSource) MPSCNNConvolution {
	rv := objc.Send[MPSCNNConvolution](c.ID, objc.Sel("initWithDevice:weights:"), device, weights)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution/exportWeightsAndBiases(with:resultStateCanBeTemporary:)
func (c MPSCNNConvolution) ExportWeightsAndBiasesWithCommandBufferResultStateCanBeTemporary(commandBuffer metal.MTLCommandBuffer, resultStateCanBeTemporary bool) IMPSCNNConvolutionWeightsAndBiasesState {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("exportWeightsAndBiasesWithCommandBuffer:resultStateCanBeTemporary:"), commandBuffer, resultStateCanBeTemporary)
	return MPSCNNConvolutionWeightsAndBiasesStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution/reloadWeightsAndBiases(with:state:)
func (c MPSCNNConvolution) ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer metal.MTLCommandBuffer, state IMPSCNNConvolutionWeightsAndBiasesState) {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadWeightsAndBiasesWithCommandBuffer:state:"), commandBuffer, state)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution/reloadWeightsAndBiasesFromDataSource()
func (c MPSCNNConvolution) ReloadWeightsAndBiasesFromDataSource() {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadWeightsAndBiasesFromDataSource"))
}

// The number of feature channels per pixel in the input image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution/inputFeatureChannels
func (c MPSCNNConvolution) InputFeatureChannels() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("inputFeatureChannels"))
	return rv
}

// The number of feature channels per pixel in the output image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution/outputFeatureChannels
func (c MPSCNNConvolution) OutputFeatureChannels() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("outputFeatureChannels"))
	return rv
}

// The number of groups that the input and output channels are divided into.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution/groups
func (c MPSCNNConvolution) Groups() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("groups"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution/subPixelScaleFactor
func (c MPSCNNConvolution) SubPixelScaleFactor() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("subPixelScaleFactor"))
	return rv
}

// The neuron filter to be applied as part of the convolution operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution/neuron
func (c MPSCNNConvolution) Neuron() IMPSCNNNeuron {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("neuron"))
	return MPSCNNNeuronFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution/accumulatorPrecisionOption
func (c MPSCNNConvolution) AccumulatorPrecisionOption() MPSNNConvolutionAccumulatorPrecisionOption {
	rv := objc.Send[MPSNNConvolutionAccumulatorPrecisionOption](c.ID, objc.Sel("accumulatorPrecisionOption"))
	return MPSNNConvolutionAccumulatorPrecisionOption(rv)
}
func (c MPSCNNConvolution) SetAccumulatorPrecisionOption(value MPSNNConvolutionAccumulatorPrecisionOption) {
	objc.Send[struct{}](c.ID, objc.Sel("setAccumulatorPrecisionOption:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution/channelMultiplier
func (c MPSCNNConvolution) ChannelMultiplier() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("channelMultiplier"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution/dataSource
func (c MPSCNNConvolution) DataSource() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dataSource"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution/fusedNeuronDescriptor
func (c MPSCNNConvolution) FusedNeuronDescriptor() IMPSNNNeuronDescriptor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("fusedNeuronDescriptor"))
	return MPSNNNeuronDescriptorFromID(objc.ID(rv))
}
