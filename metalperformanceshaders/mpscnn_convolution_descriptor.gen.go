// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNConvolutionDescriptor] class.
var (
	_MPSCNNConvolutionDescriptorClass     MPSCNNConvolutionDescriptorClass
	_MPSCNNConvolutionDescriptorClassOnce sync.Once
)

func getMPSCNNConvolutionDescriptorClass() MPSCNNConvolutionDescriptorClass {
	_MPSCNNConvolutionDescriptorClassOnce.Do(func() {
		_MPSCNNConvolutionDescriptorClass = MPSCNNConvolutionDescriptorClass{class: objc.GetClass("MPSCNNConvolutionDescriptor")}
	})
	return _MPSCNNConvolutionDescriptorClass
}

// GetMPSCNNConvolutionDescriptorClass returns the class object for MPSCNNConvolutionDescriptor.
func GetMPSCNNConvolutionDescriptorClass() MPSCNNConvolutionDescriptorClass {
	return getMPSCNNConvolutionDescriptorClass()
}

type MPSCNNConvolutionDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNConvolutionDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNConvolutionDescriptorClass) Alloc() MPSCNNConvolutionDescriptor {
	rv := objc.Send[MPSCNNConvolutionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of the attributes of a convolution kernel.
//
// # Overview
//
// You use an [MPSCNNConvolutionDescriptor] object to describe the properties
// of an [MPSCNNConvolution] kernel such as its size, pixel format and CPU
// cache mode.
//
// # Type Methods
//
//   - [MPSCNNConvolutionDescriptor.InitWithCoder]
//
// # Instance Properties
//
//   - [MPSCNNConvolutionDescriptor.Groups]: The number of groups that the input and output channels are divided into.
//   - [MPSCNNConvolutionDescriptor.SetGroups]
//   - [MPSCNNConvolutionDescriptor.InputFeatureChannels]: The number of feature channels per pixel in the input image.
//   - [MPSCNNConvolutionDescriptor.SetInputFeatureChannels]
//   - [MPSCNNConvolutionDescriptor.KernelHeight]: The height of the kernel window.
//   - [MPSCNNConvolutionDescriptor.SetKernelHeight]
//   - [MPSCNNConvolutionDescriptor.KernelWidth]: The width of the kernel window.
//   - [MPSCNNConvolutionDescriptor.SetKernelWidth]
//   - [MPSCNNConvolutionDescriptor.OutputFeatureChannels]: The number of feature channels per pixel in the output image.
//   - [MPSCNNConvolutionDescriptor.SetOutputFeatureChannels]
//   - [MPSCNNConvolutionDescriptor.StrideInPixelsX]: The output stride (downsampling factor) in the x dimension.
//   - [MPSCNNConvolutionDescriptor.SetStrideInPixelsX]
//   - [MPSCNNConvolutionDescriptor.StrideInPixelsY]: The output stride (downsampling factor) in the y dimension.
//   - [MPSCNNConvolutionDescriptor.SetStrideInPixelsY]
//   - [MPSCNNConvolutionDescriptor.Neuron]: The neuron filter to be applied as part of the convolution operation.
//   - [MPSCNNConvolutionDescriptor.SetNeuron]
//   - [MPSCNNConvolutionDescriptor.DilationRateX]
//   - [MPSCNNConvolutionDescriptor.SetDilationRateX]
//   - [MPSCNNConvolutionDescriptor.DilationRateY]
//   - [MPSCNNConvolutionDescriptor.SetDilationRateY]
//   - [MPSCNNConvolutionDescriptor.FusedNeuronDescriptor]
//   - [MPSCNNConvolutionDescriptor.SetFusedNeuronDescriptor]
//
// # Instance Methods
//
//   - [MPSCNNConvolutionDescriptor.EncodeWithCoder]
//   - [MPSCNNConvolutionDescriptor.SetBatchNormalizationParametersForInferenceWithMeanVarianceGammaBetaEpsilon]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor
type MPSCNNConvolutionDescriptor struct {
	objectivec.Object
}

// MPSCNNConvolutionDescriptorFromID constructs a [MPSCNNConvolutionDescriptor] from an objc.ID.
//
// A description of the attributes of a convolution kernel.
func MPSCNNConvolutionDescriptorFromID(id objc.ID) MPSCNNConvolutionDescriptor {
	return MPSCNNConvolutionDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MPSCNNConvolutionDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNConvolutionDescriptor] class.
//
// # Type Methods
//
//   - [IMPSCNNConvolutionDescriptor.InitWithCoder]
//
// # Instance Properties
//
//   - [IMPSCNNConvolutionDescriptor.Groups]: The number of groups that the input and output channels are divided into.
//   - [IMPSCNNConvolutionDescriptor.SetGroups]
//   - [IMPSCNNConvolutionDescriptor.InputFeatureChannels]: The number of feature channels per pixel in the input image.
//   - [IMPSCNNConvolutionDescriptor.SetInputFeatureChannels]
//   - [IMPSCNNConvolutionDescriptor.KernelHeight]: The height of the kernel window.
//   - [IMPSCNNConvolutionDescriptor.SetKernelHeight]
//   - [IMPSCNNConvolutionDescriptor.KernelWidth]: The width of the kernel window.
//   - [IMPSCNNConvolutionDescriptor.SetKernelWidth]
//   - [IMPSCNNConvolutionDescriptor.OutputFeatureChannels]: The number of feature channels per pixel in the output image.
//   - [IMPSCNNConvolutionDescriptor.SetOutputFeatureChannels]
//   - [IMPSCNNConvolutionDescriptor.StrideInPixelsX]: The output stride (downsampling factor) in the x dimension.
//   - [IMPSCNNConvolutionDescriptor.SetStrideInPixelsX]
//   - [IMPSCNNConvolutionDescriptor.StrideInPixelsY]: The output stride (downsampling factor) in the y dimension.
//   - [IMPSCNNConvolutionDescriptor.SetStrideInPixelsY]
//   - [IMPSCNNConvolutionDescriptor.Neuron]: The neuron filter to be applied as part of the convolution operation.
//   - [IMPSCNNConvolutionDescriptor.SetNeuron]
//   - [IMPSCNNConvolutionDescriptor.DilationRateX]
//   - [IMPSCNNConvolutionDescriptor.SetDilationRateX]
//   - [IMPSCNNConvolutionDescriptor.DilationRateY]
//   - [IMPSCNNConvolutionDescriptor.SetDilationRateY]
//   - [IMPSCNNConvolutionDescriptor.FusedNeuronDescriptor]
//   - [IMPSCNNConvolutionDescriptor.SetFusedNeuronDescriptor]
//
// # Instance Methods
//
//   - [IMPSCNNConvolutionDescriptor.EncodeWithCoder]
//   - [IMPSCNNConvolutionDescriptor.SetBatchNormalizationParametersForInferenceWithMeanVarianceGammaBetaEpsilon]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor
type IMPSCNNConvolutionDescriptor interface {
	objectivec.IObject

	// Topic: Type Methods

	InitWithCoder(aDecoder foundation.INSCoder) MPSCNNConvolutionDescriptor

	// Topic: Instance Properties

	// The number of groups that the input and output channels are divided into.
	Groups() uint
	SetGroups(value uint)
	// The number of feature channels per pixel in the input image.
	InputFeatureChannels() uint
	SetInputFeatureChannels(value uint)
	// The height of the kernel window.
	KernelHeight() uint
	SetKernelHeight(value uint)
	// The width of the kernel window.
	KernelWidth() uint
	SetKernelWidth(value uint)
	// The number of feature channels per pixel in the output image.
	OutputFeatureChannels() uint
	SetOutputFeatureChannels(value uint)
	// The output stride (downsampling factor) in the x dimension.
	StrideInPixelsX() uint
	SetStrideInPixelsX(value uint)
	// The output stride (downsampling factor) in the y dimension.
	StrideInPixelsY() uint
	SetStrideInPixelsY(value uint)
	// The neuron filter to be applied as part of the convolution operation.
	Neuron() IMPSCNNNeuron
	SetNeuron(value IMPSCNNNeuron)
	DilationRateX() uint
	SetDilationRateX(value uint)
	DilationRateY() uint
	SetDilationRateY(value uint)
	FusedNeuronDescriptor() IMPSNNNeuronDescriptor
	SetFusedNeuronDescriptor(value IMPSNNNeuronDescriptor)

	// Topic: Instance Methods

	EncodeWithCoder(aCoder foundation.INSCoder)
	SetBatchNormalizationParametersForInferenceWithMeanVarianceGammaBetaEpsilon(mean *float32, variance *float32, gamma *float32, beta *float32, epsilon float32)
}

// Init initializes the instance.
func (c MPSCNNConvolutionDescriptor) Init() MPSCNNConvolutionDescriptor {
	rv := objc.Send[MPSCNNConvolutionDescriptor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNConvolutionDescriptor) Autorelease() MPSCNNConvolutionDescriptor {
	rv := objc.Send[MPSCNNConvolutionDescriptor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNConvolutionDescriptor creates a new MPSCNNConvolutionDescriptor instance.
func NewMPSCNNConvolutionDescriptor() MPSCNNConvolutionDescriptor {
	class := getMPSCNNConvolutionDescriptorClass()
	rv := objc.Send[MPSCNNConvolutionDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/init(kernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:)
func NewCNNConvolutionDescriptorCnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint) MPSCNNConvolutionDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNConvolutionDescriptorClass().class), objc.Sel("cnnConvolutionDescriptorWithKernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:"), kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels)
	return MPSCNNConvolutionDescriptorFromID(rv)
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
func NewCNNConvolutionDescriptorCnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannelsNeuronFilter(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint, neuronFilter IMPSCNNNeuron) MPSCNNConvolutionDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNConvolutionDescriptorClass().class), objc.Sel("cnnConvolutionDescriptorWithKernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:neuronFilter:"), kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels, neuronFilter)
	return MPSCNNConvolutionDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/init(coder:)
func NewCNNConvolutionDescriptorWithCoder(aDecoder foundation.INSCoder) MPSCNNConvolutionDescriptor {
	instance := getMPSCNNConvolutionDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNConvolutionDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/init(coder:)
func (c MPSCNNConvolutionDescriptor) InitWithCoder(aDecoder foundation.INSCoder) MPSCNNConvolutionDescriptor {
	rv := objc.Send[MPSCNNConvolutionDescriptor](c.ID, objc.Sel("initWithCoder:"), aDecoder)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/encode(with:)
func (c MPSCNNConvolutionDescriptor) EncodeWithCoder(aCoder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), aCoder)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/setBatchNormalizationParametersForInferenceWithMean(_:variance:gamma:beta:epsilon:)
func (c MPSCNNConvolutionDescriptor) SetBatchNormalizationParametersForInferenceWithMeanVarianceGammaBetaEpsilon(mean *float32, variance *float32, gamma *float32, beta *float32, epsilon float32) {
	objc.Send[objc.ID](c.ID, objc.Sel("setBatchNormalizationParametersForInferenceWithMean:variance:gamma:beta:epsilon:"), mean, variance, gamma, beta, epsilon)
}

// The number of groups that the input and output channels are divided into.
//
// # Discussion
//
// The default value is `1`.
//
// Groups let you reduce parametrization. If the value of this property is set
// to `n`, the input is divided into `n` groups with
// `inputFeatureChannels“/n` channels in each group. Similarly, the output is
// divided into `n` groups with `outputFeatureChannels/n` channels in each
// group. The `ith` group in the input is only connected to the `ith` group in
// the output, so the number of weights (parameters) needed is reduced by a
// factor of `n`. Both the value of the
// [MPSCNNConvolutionDescriptor.InputFeatureChannels] and
// [MPSCNNConvolutionDescriptor.OutputFeatureChannels] properties must be
// divisible by `n` and the number of channels in each group must be a
// multiple of `4`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/groups
func (c MPSCNNConvolutionDescriptor) Groups() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("groups"))
	return rv
}
func (c MPSCNNConvolutionDescriptor) SetGroups(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setGroups:"), value)
}

// The number of feature channels per pixel in the input image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/inputFeatureChannels
func (c MPSCNNConvolutionDescriptor) InputFeatureChannels() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("inputFeatureChannels"))
	return rv
}
func (c MPSCNNConvolutionDescriptor) SetInputFeatureChannels(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setInputFeatureChannels:"), value)
}

// The height of the kernel window.
//
// # Discussion
//
// The default value is `3`.
//
// Any positive non-zero value is valid, including even values. The position
// of the top edge of the kernel window is given by `offset.Y() -
// (kernelHeight>>1)`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/kernelHeight
func (c MPSCNNConvolutionDescriptor) KernelHeight() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelHeight"))
	return rv
}
func (c MPSCNNConvolutionDescriptor) SetKernelHeight(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setKernelHeight:"), value)
}

// The width of the kernel window.
//
// # Discussion
//
// The default value is `3`.
//
// Any positive non-zero value is valid, including even values. The position
// of the left edge of the kernel window is given by `offset.X() -
// (kernelWidth>>1)`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/kernelWidth
func (c MPSCNNConvolutionDescriptor) KernelWidth() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelWidth"))
	return rv
}
func (c MPSCNNConvolutionDescriptor) SetKernelWidth(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setKernelWidth:"), value)
}

// The number of feature channels per pixel in the output image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/outputFeatureChannels
func (c MPSCNNConvolutionDescriptor) OutputFeatureChannels() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("outputFeatureChannels"))
	return rv
}
func (c MPSCNNConvolutionDescriptor) SetOutputFeatureChannels(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setOutputFeatureChannels:"), value)
}

// The output stride (downsampling factor) in the x dimension.
//
// # Discussion
//
// The default value is `1`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/strideInPixelsX
func (c MPSCNNConvolutionDescriptor) StrideInPixelsX() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("strideInPixelsX"))
	return rv
}
func (c MPSCNNConvolutionDescriptor) SetStrideInPixelsX(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setStrideInPixelsX:"), value)
}

// The output stride (downsampling factor) in the y dimension.
//
// # Discussion
//
// The default value is `1`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/strideInPixelsY
func (c MPSCNNConvolutionDescriptor) StrideInPixelsY() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("strideInPixelsY"))
	return rv
}
func (c MPSCNNConvolutionDescriptor) SetStrideInPixelsY(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setStrideInPixelsY:"), value)
}

// The neuron filter to be applied as part of the convolution operation.
//
// # Discussion
//
// The default value is `nil`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/neuron
func (c MPSCNNConvolutionDescriptor) Neuron() IMPSCNNNeuron {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("neuron"))
	return MPSCNNNeuronFromID(objc.ID(rv))
}
func (c MPSCNNConvolutionDescriptor) SetNeuron(value IMPSCNNNeuron) {
	objc.Send[struct{}](c.ID, objc.Sel("setNeuron:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/dilationRateX
func (c MPSCNNConvolutionDescriptor) DilationRateX() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("dilationRateX"))
	return rv
}
func (c MPSCNNConvolutionDescriptor) SetDilationRateX(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setDilationRateX:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/dilationRateY
func (c MPSCNNConvolutionDescriptor) DilationRateY() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("dilationRateY"))
	return rv
}
func (c MPSCNNConvolutionDescriptor) SetDilationRateY(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setDilationRateY:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/fusedNeuronDescriptor
func (c MPSCNNConvolutionDescriptor) FusedNeuronDescriptor() IMPSNNNeuronDescriptor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("fusedNeuronDescriptor"))
	return MPSNNNeuronDescriptorFromID(objc.ID(rv))
}
func (c MPSCNNConvolutionDescriptor) SetFusedNeuronDescriptor(value IMPSNNNeuronDescriptor) {
	objc.Send[struct{}](c.ID, objc.Sel("setFusedNeuronDescriptor:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/supportsSecureCoding
func (_MPSCNNConvolutionDescriptorClass MPSCNNConvolutionDescriptorClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MPSCNNConvolutionDescriptorClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
