// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGRUDescriptor] class.
var (
	_MPSGRUDescriptorClass     MPSGRUDescriptorClass
	_MPSGRUDescriptorClassOnce sync.Once
)

func getMPSGRUDescriptorClass() MPSGRUDescriptorClass {
	_MPSGRUDescriptorClassOnce.Do(func() {
		_MPSGRUDescriptorClass = MPSGRUDescriptorClass{class: objc.GetClass("MPSGRUDescriptor")}
	})
	return _MPSGRUDescriptorClass
}

// GetMPSGRUDescriptorClass returns the class object for MPSGRUDescriptor.
func GetMPSGRUDescriptorClass() MPSGRUDescriptorClass {
	return getMPSGRUDescriptorClass()
}

type MPSGRUDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGRUDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGRUDescriptorClass) Alloc() MPSGRUDescriptor {
	rv := objc.Send[MPSGRUDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of a gated recurrent unit block or layer.
//
// # Overview
//
// The recurrent neural network (RNN) layer initialized with a
// [MPSGRUDescriptor] transforms the input data (image or matrix) and previous
// output with a set of filters. Each produces one feature map in the output
// data according to the gated recurrent unit (GRU) unit formula detailed
// below.
//
// You may provide the GRU unit with a single input or a sequence of inputs.
// The layer also supports p-norm gating.
//
// # Description of Operation
//
// - Let `x_j` be the input data (at time index `t` of sequence, `j` index
// containing quadruplet: batch index, `x,y` and feature index (`x = y = 0`
// for matrices)). - Let `h0_`j be the recurrent input (previous output) data
// from previous time step (at time index `t-1` of sequence). - Let `h_i` be
// the proposed new output. - Let `h1_i` be the output data produced at this
// time step. - Let `Wz_ij`, `Uz_ij` be the input gate weights for input and
// recurrent input data, respectively. - Let `bi_i` be the bias for the input
// gate. - Let `Wr_ij`, `Ur_ij` be the recurrent gate weights for input and
// recurrent input data, respectively. - Let `br_i` be the bias for the
// recurrent gate. - Let `Wh_ij`, `Uh_ij`, `Vh_ij` be the output gate weights
// for input, recurrent gate, and input gate, respectively. - Let `bh_i` be
// the bias for the output gate. - Let `gz(x“)`, `gr(x)`, `gh(x)` be the
// neuron activation function for the input, recurrent, and output gates. -
// Let `p > 0` be a scalar variable (typical `p >= 1.0`) that defines the
// p-norm gating norm value.
//
// The output of the GRU layer is computed as follows:
//
// The `*` stands for convolution (see [MPSRNNImageInferenceLayer]) or
// matrix-vector/matrix multiplication (see [MPSRNNMatrixInferenceLayer]).
//
// Summation is over index `j` (except for the batch index), but there’s no
// summation over repeated index `i`,` `the output index.
//
// Note that for validity, all intermediate images must be of same size, and
// all [U] and [V] matrices must be square (that is,
// [MPSRNNDescriptor.OutputFeatureChannels] `==`
// [MPSRNNDescriptor.InputFeatureChannels]). Also, the bias terms are scalars
// with regard to spatial dimensions. The conventional GRU block is achieved
// by setting `Vh = 0` (nil), and the Minimal Gated Unit is achieved with `Uh
// = 0`.
//
// # Instance Properties
//
//   - [MPSGRUDescriptor.FlipOutputGates]
//   - [MPSGRUDescriptor.SetFlipOutputGates]
//   - [MPSGRUDescriptor.GatePnormValue]
//   - [MPSGRUDescriptor.SetGatePnormValue]
//   - [MPSGRUDescriptor.InputGateInputWeights]
//   - [MPSGRUDescriptor.SetInputGateInputWeights]
//   - [MPSGRUDescriptor.InputGateRecurrentWeights]
//   - [MPSGRUDescriptor.SetInputGateRecurrentWeights]
//   - [MPSGRUDescriptor.OutputGateInputGateWeights]
//   - [MPSGRUDescriptor.SetOutputGateInputGateWeights]
//   - [MPSGRUDescriptor.OutputGateInputWeights]
//   - [MPSGRUDescriptor.SetOutputGateInputWeights]
//   - [MPSGRUDescriptor.OutputGateRecurrentWeights]
//   - [MPSGRUDescriptor.SetOutputGateRecurrentWeights]
//   - [MPSGRUDescriptor.RecurrentGateInputWeights]
//   - [MPSGRUDescriptor.SetRecurrentGateInputWeights]
//   - [MPSGRUDescriptor.RecurrentGateRecurrentWeights]
//   - [MPSGRUDescriptor.SetRecurrentGateRecurrentWeights]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSGRUDescriptor
type MPSGRUDescriptor struct {
	MPSRNNDescriptor
}

// MPSGRUDescriptorFromID constructs a [MPSGRUDescriptor] from an objc.ID.
//
// A description of a gated recurrent unit block or layer.
func MPSGRUDescriptorFromID(id objc.ID) MPSGRUDescriptor {
	return MPSGRUDescriptor{MPSRNNDescriptor: MPSRNNDescriptorFromID(id)}
}

// NOTE: MPSGRUDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGRUDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSGRUDescriptor.FlipOutputGates]
//   - [IMPSGRUDescriptor.SetFlipOutputGates]
//   - [IMPSGRUDescriptor.GatePnormValue]
//   - [IMPSGRUDescriptor.SetGatePnormValue]
//   - [IMPSGRUDescriptor.InputGateInputWeights]
//   - [IMPSGRUDescriptor.SetInputGateInputWeights]
//   - [IMPSGRUDescriptor.InputGateRecurrentWeights]
//   - [IMPSGRUDescriptor.SetInputGateRecurrentWeights]
//   - [IMPSGRUDescriptor.OutputGateInputGateWeights]
//   - [IMPSGRUDescriptor.SetOutputGateInputGateWeights]
//   - [IMPSGRUDescriptor.OutputGateInputWeights]
//   - [IMPSGRUDescriptor.SetOutputGateInputWeights]
//   - [IMPSGRUDescriptor.OutputGateRecurrentWeights]
//   - [IMPSGRUDescriptor.SetOutputGateRecurrentWeights]
//   - [IMPSGRUDescriptor.RecurrentGateInputWeights]
//   - [IMPSGRUDescriptor.SetRecurrentGateInputWeights]
//   - [IMPSGRUDescriptor.RecurrentGateRecurrentWeights]
//   - [IMPSGRUDescriptor.SetRecurrentGateRecurrentWeights]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSGRUDescriptor
type IMPSGRUDescriptor interface {
	IMPSRNNDescriptor

	// Topic: Instance Properties

	FlipOutputGates() bool
	SetFlipOutputGates(value bool)
	GatePnormValue() float32
	SetGatePnormValue(value float32)
	InputGateInputWeights() MPSCNNConvolutionDataSource
	SetInputGateInputWeights(value MPSCNNConvolutionDataSource)
	InputGateRecurrentWeights() MPSCNNConvolutionDataSource
	SetInputGateRecurrentWeights(value MPSCNNConvolutionDataSource)
	OutputGateInputGateWeights() MPSCNNConvolutionDataSource
	SetOutputGateInputGateWeights(value MPSCNNConvolutionDataSource)
	OutputGateInputWeights() MPSCNNConvolutionDataSource
	SetOutputGateInputWeights(value MPSCNNConvolutionDataSource)
	OutputGateRecurrentWeights() MPSCNNConvolutionDataSource
	SetOutputGateRecurrentWeights(value MPSCNNConvolutionDataSource)
	RecurrentGateInputWeights() MPSCNNConvolutionDataSource
	SetRecurrentGateInputWeights(value MPSCNNConvolutionDataSource)
	RecurrentGateRecurrentWeights() MPSCNNConvolutionDataSource
	SetRecurrentGateRecurrentWeights(value MPSCNNConvolutionDataSource)
}

// Init initializes the instance.
func (r MPSGRUDescriptor) Init() MPSGRUDescriptor {
	rv := objc.Send[MPSGRUDescriptor](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSGRUDescriptor) Autorelease() MPSGRUDescriptor {
	rv := objc.Send[MPSGRUDescriptor](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGRUDescriptor creates a new MPSGRUDescriptor instance.
func NewMPSGRUDescriptor() MPSGRUDescriptor {
	class := getMPSGRUDescriptorClass()
	rv := objc.Send[MPSGRUDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSGRUDescriptor/createGRUDescriptor(withInputFeatureChannels:outputFeatureChannels:)
func (_MPSGRUDescriptorClass MPSGRUDescriptorClass) CreateGRUDescriptorWithInputFeatureChannelsOutputFeatureChannels(inputFeatureChannels uint, outputFeatureChannels uint) MPSGRUDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSGRUDescriptorClass.class), objc.Sel("createGRUDescriptorWithInputFeatureChannels:outputFeatureChannels:"), inputFeatureChannels, outputFeatureChannels)
	return MPSGRUDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSGRUDescriptor/flipOutputGates
func (r MPSGRUDescriptor) FlipOutputGates() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("flipOutputGates"))
	return rv
}
func (r MPSGRUDescriptor) SetFlipOutputGates(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setFlipOutputGates:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSGRUDescriptor/gatePnormValue
func (r MPSGRUDescriptor) GatePnormValue() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("gatePnormValue"))
	return rv
}
func (r MPSGRUDescriptor) SetGatePnormValue(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setGatePnormValue:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSGRUDescriptor/inputGateInputWeights
func (r MPSGRUDescriptor) InputGateInputWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("inputGateInputWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (r MPSGRUDescriptor) SetInputGateInputWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](r.ID, objc.Sel("setInputGateInputWeights:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSGRUDescriptor/inputGateRecurrentWeights
func (r MPSGRUDescriptor) InputGateRecurrentWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("inputGateRecurrentWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (r MPSGRUDescriptor) SetInputGateRecurrentWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](r.ID, objc.Sel("setInputGateRecurrentWeights:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSGRUDescriptor/outputGateInputGateWeights
func (r MPSGRUDescriptor) OutputGateInputGateWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("outputGateInputGateWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (r MPSGRUDescriptor) SetOutputGateInputGateWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](r.ID, objc.Sel("setOutputGateInputGateWeights:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSGRUDescriptor/outputGateInputWeights
func (r MPSGRUDescriptor) OutputGateInputWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("outputGateInputWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (r MPSGRUDescriptor) SetOutputGateInputWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](r.ID, objc.Sel("setOutputGateInputWeights:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSGRUDescriptor/outputGateRecurrentWeights
func (r MPSGRUDescriptor) OutputGateRecurrentWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("outputGateRecurrentWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (r MPSGRUDescriptor) SetOutputGateRecurrentWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](r.ID, objc.Sel("setOutputGateRecurrentWeights:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSGRUDescriptor/recurrentGateInputWeights
func (r MPSGRUDescriptor) RecurrentGateInputWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("recurrentGateInputWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (r MPSGRUDescriptor) SetRecurrentGateInputWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](r.ID, objc.Sel("setRecurrentGateInputWeights:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSGRUDescriptor/recurrentGateRecurrentWeights
func (r MPSGRUDescriptor) RecurrentGateRecurrentWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("recurrentGateRecurrentWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (r MPSGRUDescriptor) SetRecurrentGateRecurrentWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](r.ID, objc.Sel("setRecurrentGateRecurrentWeights:"), value)
}
