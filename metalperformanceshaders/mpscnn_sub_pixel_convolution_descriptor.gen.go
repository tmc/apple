// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNSubPixelConvolutionDescriptor] class.
var (
	_MPSCNNSubPixelConvolutionDescriptorClass     MPSCNNSubPixelConvolutionDescriptorClass
	_MPSCNNSubPixelConvolutionDescriptorClassOnce sync.Once
)

func getMPSCNNSubPixelConvolutionDescriptorClass() MPSCNNSubPixelConvolutionDescriptorClass {
	_MPSCNNSubPixelConvolutionDescriptorClassOnce.Do(func() {
		_MPSCNNSubPixelConvolutionDescriptorClass = MPSCNNSubPixelConvolutionDescriptorClass{class: objc.GetClass("MPSCNNSubPixelConvolutionDescriptor")}
	})
	return _MPSCNNSubPixelConvolutionDescriptorClass
}

// GetMPSCNNSubPixelConvolutionDescriptorClass returns the class object for MPSCNNSubPixelConvolutionDescriptor.
func GetMPSCNNSubPixelConvolutionDescriptorClass() MPSCNNSubPixelConvolutionDescriptorClass {
	return getMPSCNNSubPixelConvolutionDescriptorClass()
}

type MPSCNNSubPixelConvolutionDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNSubPixelConvolutionDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNSubPixelConvolutionDescriptorClass) Alloc() MPSCNNSubPixelConvolutionDescriptor {
	rv := objc.Send[MPSCNNSubPixelConvolutionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of a convolution object that does subpixel upsampling and
// reshaping.
//
// # Instance Properties
//
//   - [MPSCNNSubPixelConvolutionDescriptor.SubPixelScaleFactor]
//   - [MPSCNNSubPixelConvolutionDescriptor.SetSubPixelScaleFactor]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSubPixelConvolutionDescriptor
type MPSCNNSubPixelConvolutionDescriptor struct {
	MPSCNNConvolutionDescriptor
}

// MPSCNNSubPixelConvolutionDescriptorFromID constructs a [MPSCNNSubPixelConvolutionDescriptor] from an objc.ID.
//
// A description of a convolution object that does subpixel upsampling and
// reshaping.
func MPSCNNSubPixelConvolutionDescriptorFromID(id objc.ID) MPSCNNSubPixelConvolutionDescriptor {
	return MPSCNNSubPixelConvolutionDescriptor{MPSCNNConvolutionDescriptor: MPSCNNConvolutionDescriptorFromID(id)}
}

// NOTE: MPSCNNSubPixelConvolutionDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNSubPixelConvolutionDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSCNNSubPixelConvolutionDescriptor.SubPixelScaleFactor]
//   - [IMPSCNNSubPixelConvolutionDescriptor.SetSubPixelScaleFactor]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSubPixelConvolutionDescriptor
type IMPSCNNSubPixelConvolutionDescriptor interface {
	IMPSCNNConvolutionDescriptor

	// Topic: Instance Properties

	SubPixelScaleFactor() uint
	SetSubPixelScaleFactor(value uint)
}

// Init initializes the instance.
func (c MPSCNNSubPixelConvolutionDescriptor) Init() MPSCNNSubPixelConvolutionDescriptor {
	rv := objc.Send[MPSCNNSubPixelConvolutionDescriptor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNSubPixelConvolutionDescriptor) Autorelease() MPSCNNSubPixelConvolutionDescriptor {
	rv := objc.Send[MPSCNNSubPixelConvolutionDescriptor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNSubPixelConvolutionDescriptor creates a new MPSCNNSubPixelConvolutionDescriptor instance.
func NewMPSCNNSubPixelConvolutionDescriptor() MPSCNNSubPixelConvolutionDescriptor {
	class := getMPSCNNSubPixelConvolutionDescriptorClass()
	rv := objc.Send[MPSCNNSubPixelConvolutionDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/init(kernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:)
func NewCNNSubPixelConvolutionDescriptorCnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint) MPSCNNSubPixelConvolutionDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNSubPixelConvolutionDescriptorClass().class), objc.Sel("cnnConvolutionDescriptorWithKernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:"), kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels)
	return MPSCNNSubPixelConvolutionDescriptorFromID(rv)
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
func NewCNNSubPixelConvolutionDescriptorCnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannelsNeuronFilter(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint, neuronFilter IMPSCNNNeuron) MPSCNNSubPixelConvolutionDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNSubPixelConvolutionDescriptorClass().class), objc.Sel("cnnConvolutionDescriptorWithKernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:neuronFilter:"), kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels, neuronFilter)
	return MPSCNNSubPixelConvolutionDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/init(coder:)
func NewCNNSubPixelConvolutionDescriptorWithCoder(aDecoder foundation.INSCoder) MPSCNNSubPixelConvolutionDescriptor {
	instance := getMPSCNNSubPixelConvolutionDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNSubPixelConvolutionDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSubPixelConvolutionDescriptor/subPixelScaleFactor
func (c MPSCNNSubPixelConvolutionDescriptor) SubPixelScaleFactor() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("subPixelScaleFactor"))
	return rv
}
func (c MPSCNNSubPixelConvolutionDescriptor) SetSubPixelScaleFactor(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setSubPixelScaleFactor:"), value)
}
