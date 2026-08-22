// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSRNNSingleGateDescriptor] class.
var (
	_MPSRNNSingleGateDescriptorClass     MPSRNNSingleGateDescriptorClass
	_MPSRNNSingleGateDescriptorClassOnce sync.Once
)

func getMPSRNNSingleGateDescriptorClass() MPSRNNSingleGateDescriptorClass {
	_MPSRNNSingleGateDescriptorClassOnce.Do(func() {
		_MPSRNNSingleGateDescriptorClass = MPSRNNSingleGateDescriptorClass{class: objc.GetClass("MPSRNNSingleGateDescriptor")}
	})
	return _MPSRNNSingleGateDescriptorClass
}

// GetMPSRNNSingleGateDescriptorClass returns the class object for MPSRNNSingleGateDescriptor.
func GetMPSRNNSingleGateDescriptorClass() MPSRNNSingleGateDescriptorClass {
	return getMPSRNNSingleGateDescriptorClass()
}

type MPSRNNSingleGateDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSRNNSingleGateDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSRNNSingleGateDescriptorClass) Alloc() MPSRNNSingleGateDescriptor {
	rv := objc.Send[MPSRNNSingleGateDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of a simple recurrent block or layer.
//
// # Overview
//
// The recurrent neural network (RNN) layer initialized with a
// [MPSRNNSingleGateDescriptor] transforms the input data (image or matrix)
// and previous output with a set of filters. Each produces one feature map in
// the new output data.
//
// You may provide the RNN unit with a single input or a sequence of inputs.
//
// # Description of Operation
//
// - Let `x_j` be the input data (at time index `t` of sequence, `j` index
// containing quadruplet: batch index, `x,y` and feature index (`x = y = 0`
// for matrices)). - Let `h0_j` be the recurrent input (previous output) data
// from previous time step (at time index `t-1` of sequence). - Let `h1_i` be
// the output data produced at this time step. - Let `W_ij, U_ij` be the
// weights for input and recurrent input data, respectively. - Let `b_i` be a
// bias term. - Let `gi(x)` be a neuron activation function.
//
// The new output image `h1_i` data is computed as follows:
//
// The `*` stands for convolution (see [MPSRNNImageInferenceLayer]) or
// matrix-vector/matrix multiplication (see [MPSRNNMatrixInferenceLayer]).
//
// Summation is over index `j` (except for the batch index), but there’s no
// summation over repeated index `i` (the output index).
//
// Note that for validity, all intermediate images must be of same size, and
// the [U] matrix must be square (that is,
// [MPSRNNDescriptor.OutputFeatureChannels] `==`
// [MPSRNNDescriptor.InputFeatureChannels]). Also, the bias terms are scalars
// with regard to spatial dimensions.
//
// # Instance Properties
//
//   - [MPSRNNSingleGateDescriptor.InputWeights]
//   - [MPSRNNSingleGateDescriptor.SetInputWeights]
//   - [MPSRNNSingleGateDescriptor.RecurrentWeights]
//   - [MPSRNNSingleGateDescriptor.SetRecurrentWeights]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNSingleGateDescriptor
type MPSRNNSingleGateDescriptor struct {
	MPSRNNDescriptor
}

// MPSRNNSingleGateDescriptorFromID constructs a [MPSRNNSingleGateDescriptor] from an objc.ID.
//
// A description of a simple recurrent block or layer.
func MPSRNNSingleGateDescriptorFromID(id objc.ID) MPSRNNSingleGateDescriptor {
	return MPSRNNSingleGateDescriptor{MPSRNNDescriptor: MPSRNNDescriptorFromID(id)}
}

// NOTE: MPSRNNSingleGateDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSRNNSingleGateDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSRNNSingleGateDescriptor.InputWeights]
//   - [IMPSRNNSingleGateDescriptor.SetInputWeights]
//   - [IMPSRNNSingleGateDescriptor.RecurrentWeights]
//   - [IMPSRNNSingleGateDescriptor.SetRecurrentWeights]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNSingleGateDescriptor
type IMPSRNNSingleGateDescriptor interface {
	IMPSRNNDescriptor

	// Topic: Instance Properties

	InputWeights() MPSCNNConvolutionDataSource
	SetInputWeights(value MPSCNNConvolutionDataSource)
	RecurrentWeights() MPSCNNConvolutionDataSource
	SetRecurrentWeights(value MPSCNNConvolutionDataSource)
}

// Init initializes the instance.
func (r MPSRNNSingleGateDescriptor) Init() MPSRNNSingleGateDescriptor {
	rv := objc.Send[MPSRNNSingleGateDescriptor](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSRNNSingleGateDescriptor) Autorelease() MPSRNNSingleGateDescriptor {
	rv := objc.Send[MPSRNNSingleGateDescriptor](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSRNNSingleGateDescriptor creates a new MPSRNNSingleGateDescriptor instance.
func NewMPSRNNSingleGateDescriptor() MPSRNNSingleGateDescriptor {
	class := getMPSRNNSingleGateDescriptorClass()
	rv := objc.Send[MPSRNNSingleGateDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNSingleGateDescriptor/createRNNSingleGateDescriptor(withInputFeatureChannels:outputFeatureChannels:)
func (_MPSRNNSingleGateDescriptorClass MPSRNNSingleGateDescriptorClass) CreateRNNSingleGateDescriptorWithInputFeatureChannelsOutputFeatureChannels(inputFeatureChannels uint, outputFeatureChannels uint) MPSRNNSingleGateDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSRNNSingleGateDescriptorClass.class), objc.Sel("createRNNSingleGateDescriptorWithInputFeatureChannels:outputFeatureChannels:"), inputFeatureChannels, outputFeatureChannels)
	return MPSRNNSingleGateDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNSingleGateDescriptor/inputWeights
func (r MPSRNNSingleGateDescriptor) InputWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("inputWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (r MPSRNNSingleGateDescriptor) SetInputWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](r.ID, objc.Sel("setInputWeights:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNSingleGateDescriptor/recurrentWeights
func (r MPSRNNSingleGateDescriptor) RecurrentWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("recurrentWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (r MPSRNNSingleGateDescriptor) SetRecurrentWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](r.ID, objc.Sel("setRecurrentWeights:"), value)
}
