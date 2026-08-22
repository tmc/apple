// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNDepthWiseConvolutionDescriptor] class.
var (
	_MPSCNNDepthWiseConvolutionDescriptorClass     MPSCNNDepthWiseConvolutionDescriptorClass
	_MPSCNNDepthWiseConvolutionDescriptorClassOnce sync.Once
)

func getMPSCNNDepthWiseConvolutionDescriptorClass() MPSCNNDepthWiseConvolutionDescriptorClass {
	_MPSCNNDepthWiseConvolutionDescriptorClassOnce.Do(func() {
		_MPSCNNDepthWiseConvolutionDescriptorClass = MPSCNNDepthWiseConvolutionDescriptorClass{class: objc.GetClass("MPSCNNDepthWiseConvolutionDescriptor")}
	})
	return _MPSCNNDepthWiseConvolutionDescriptorClass
}

// GetMPSCNNDepthWiseConvolutionDescriptorClass returns the class object for MPSCNNDepthWiseConvolutionDescriptor.
func GetMPSCNNDepthWiseConvolutionDescriptorClass() MPSCNNDepthWiseConvolutionDescriptorClass {
	return getMPSCNNDepthWiseConvolutionDescriptorClass()
}

type MPSCNNDepthWiseConvolutionDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNDepthWiseConvolutionDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNDepthWiseConvolutionDescriptorClass) Alloc() MPSCNNDepthWiseConvolutionDescriptor {
	rv := objc.Send[MPSCNNDepthWiseConvolutionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of a convolution object that does depthwise convolution.
//
// # Instance Properties
//
//   - [MPSCNNDepthWiseConvolutionDescriptor.ChannelMultiplier]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDepthWiseConvolutionDescriptor
type MPSCNNDepthWiseConvolutionDescriptor struct {
	MPSCNNConvolutionDescriptor
}

// MPSCNNDepthWiseConvolutionDescriptorFromID constructs a [MPSCNNDepthWiseConvolutionDescriptor] from an objc.ID.
//
// A description of a convolution object that does depthwise convolution.
func MPSCNNDepthWiseConvolutionDescriptorFromID(id objc.ID) MPSCNNDepthWiseConvolutionDescriptor {
	return MPSCNNDepthWiseConvolutionDescriptor{MPSCNNConvolutionDescriptor: MPSCNNConvolutionDescriptorFromID(id)}
}

// NOTE: MPSCNNDepthWiseConvolutionDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNDepthWiseConvolutionDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSCNNDepthWiseConvolutionDescriptor.ChannelMultiplier]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDepthWiseConvolutionDescriptor
type IMPSCNNDepthWiseConvolutionDescriptor interface {
	IMPSCNNConvolutionDescriptor

	// Topic: Instance Properties

	ChannelMultiplier() uint
}

// Init initializes the instance.
func (c MPSCNNDepthWiseConvolutionDescriptor) Init() MPSCNNDepthWiseConvolutionDescriptor {
	rv := objc.Send[MPSCNNDepthWiseConvolutionDescriptor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNDepthWiseConvolutionDescriptor) Autorelease() MPSCNNDepthWiseConvolutionDescriptor {
	rv := objc.Send[MPSCNNDepthWiseConvolutionDescriptor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNDepthWiseConvolutionDescriptor creates a new MPSCNNDepthWiseConvolutionDescriptor instance.
func NewMPSCNNDepthWiseConvolutionDescriptor() MPSCNNDepthWiseConvolutionDescriptor {
	class := getMPSCNNDepthWiseConvolutionDescriptorClass()
	rv := objc.Send[MPSCNNDepthWiseConvolutionDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/init(kernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:)
func NewCNNDepthWiseConvolutionDescriptorCnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint) MPSCNNDepthWiseConvolutionDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNDepthWiseConvolutionDescriptorClass().class), objc.Sel("cnnConvolutionDescriptorWithKernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:"), kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels)
	return MPSCNNDepthWiseConvolutionDescriptorFromID(rv)
}

// Creates a convolution descriptor with an optional neuron filter.
//
// kernelWidth: The width of the kernel window.
//
// This value must be `>0`. Larger values will take a longer time to process.
//
// kernelHeight: The height of the kernel window.
//
// The value must be `>0`. Larger values will take a longer time to process.
//
// inputFeatureChannels: The number of feature channels in the input image.
//
// This value must be `>=1`.
//
// outputFeatureChannels: The number of feature channels in the output image.
//
// This value must be `>=1`.
//
// neuronFilter: An optional neuron filter that can be applied to the output of the
// convolution operation.
//
// # Return Value
//
// A valid [MPSCNNConvolution] object or `nil`, if failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/init(kernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:neuronFilter:)
func NewCNNDepthWiseConvolutionDescriptorCnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannelsNeuronFilter(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint, neuronFilter IMPSCNNNeuron) MPSCNNDepthWiseConvolutionDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNDepthWiseConvolutionDescriptorClass().class), objc.Sel("cnnConvolutionDescriptorWithKernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:neuronFilter:"), kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels, neuronFilter)
	return MPSCNNDepthWiseConvolutionDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/init(coder:)
func NewCNNDepthWiseConvolutionDescriptorWithCoder(aDecoder foundation.INSCoder) MPSCNNDepthWiseConvolutionDescriptor {
	instance := getMPSCNNDepthWiseConvolutionDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNDepthWiseConvolutionDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDepthWiseConvolutionDescriptor/channelMultiplier
func (c MPSCNNDepthWiseConvolutionDescriptor) ChannelMultiplier() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("channelMultiplier"))
	return rv
}
