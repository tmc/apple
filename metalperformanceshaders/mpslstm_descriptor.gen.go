// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSLSTMDescriptor] class.
var (
	_MPSLSTMDescriptorClass     MPSLSTMDescriptorClass
	_MPSLSTMDescriptorClassOnce sync.Once
)

func getMPSLSTMDescriptorClass() MPSLSTMDescriptorClass {
	_MPSLSTMDescriptorClassOnce.Do(func() {
		_MPSLSTMDescriptorClass = MPSLSTMDescriptorClass{class: objc.GetClass("MPSLSTMDescriptor")}
	})
	return _MPSLSTMDescriptorClass
}

// GetMPSLSTMDescriptorClass returns the class object for MPSLSTMDescriptor.
func GetMPSLSTMDescriptorClass() MPSLSTMDescriptorClass {
	return getMPSLSTMDescriptorClass()
}

type MPSLSTMDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSLSTMDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSLSTMDescriptorClass) Alloc() MPSLSTMDescriptor {
	rv := objc.Send[MPSLSTMDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of a long short-term memory block or layer.
//
// # Overview
//
// The recurrent neural network (RNN) layer initialized with
// [MPSLSTMDescriptor] transforms the input data (image or matrix), the memory
// cell data, and previous output with a set of filters. Each produces one
// feature map in the output data and memory cell according to the long
// short-term memory (LSTM) formula detailed below.
//
// You may provide the LSTM unit with a single input or a sequence of inputs.
//
// # Description of Operation
//
// - Let `x_j` be the input data (at time index `t` of sequence, `j` index
// containing quadruplet: batch index, `x`,`y` and feature index (`x = y = 0`
// for matrices)). - Let `h0_j` be the recurrent input (previous output) data
// from previous time step (at time index `t-1` of sequence). - Let `h1_i` be
// the output data produced at this time step. - Let `c0_j` be the previous
// memory cell data (at time index `t-1` of sequence). - Let `c1_i` be the new
// memory cell data (at time index `t-1` of sequence). - Let `Wi_ij`, `Ui_ij`,
// `Vi_ij` be the input gate weights for input, recurrent input, and memory
// cell (peephole) data, respectively. - Let `bi_i` be the bias for the input
// gate. - Let `Wf_ij`, `Uf_ij`, `Vf_ij` be the forget gate weights for input,
// recurrent input, and memory cell data, respectively. - Let `bf_i` be the
// bias for the forget gate. - Let `Wo_ij`, `Uo_ij`, `Vo_ij` be the output
// gate weights for input, recurrent input, and memory cell data,
// respectively. - Let `bo_i` be the bias for the output gate. - Let `Wc_ij`,
// `Uc_ij`, `Vc_ij` be the memory cell gate weights for input, recurrent
// input, and memory cell data, respectively. - Let `bc_i` be the bias for the
// memory cell gate. - Let `gi(x)`, `gf(x)`, `go(x)`, `gc(x)` be the neuron
// activation function for the input, forget, output gate, and memory cell
// gate. - Let `gh(x)` be the activation function applied to result memory
// cell data.
//
// The output of the LSTM layer is computed as follows:
//
// The `*` stands for convolution (see [MPSRNNImageInferenceLayer]) or
// matrix-vector/matrix multiplication (see [MPSRNNMatrixInferenceLayer]).
//
// Summation is over index `j` (except for the batch index), but there’s no
// summation over repeated index `i` (the output index).
//
// Note that for validity, all intermediate images must be of same size, and
// all [U] and [V] matrices must be square (that is,
// [MPSRNNDescriptor.OutputFeatureChannels] ==
// [MPSRNNDescriptor.InputFeatureChannels]). Also, the bias terms are scalars
// with regard to spatial dimensions.
//
// # Instance Properties
//
//   - [MPSLSTMDescriptor.CellGateInputWeights]
//   - [MPSLSTMDescriptor.SetCellGateInputWeights]
//   - [MPSLSTMDescriptor.CellGateMemoryWeights]
//   - [MPSLSTMDescriptor.SetCellGateMemoryWeights]
//   - [MPSLSTMDescriptor.CellGateRecurrentWeights]
//   - [MPSLSTMDescriptor.SetCellGateRecurrentWeights]
//   - [MPSLSTMDescriptor.CellToOutputNeuronParamA]
//   - [MPSLSTMDescriptor.SetCellToOutputNeuronParamA]
//   - [MPSLSTMDescriptor.CellToOutputNeuronParamB]
//   - [MPSLSTMDescriptor.SetCellToOutputNeuronParamB]
//   - [MPSLSTMDescriptor.CellToOutputNeuronType]
//   - [MPSLSTMDescriptor.SetCellToOutputNeuronType]
//   - [MPSLSTMDescriptor.ForgetGateInputWeights]
//   - [MPSLSTMDescriptor.SetForgetGateInputWeights]
//   - [MPSLSTMDescriptor.ForgetGateMemoryWeights]
//   - [MPSLSTMDescriptor.SetForgetGateMemoryWeights]
//   - [MPSLSTMDescriptor.ForgetGateRecurrentWeights]
//   - [MPSLSTMDescriptor.SetForgetGateRecurrentWeights]
//   - [MPSLSTMDescriptor.InputGateInputWeights]
//   - [MPSLSTMDescriptor.SetInputGateInputWeights]
//   - [MPSLSTMDescriptor.InputGateMemoryWeights]
//   - [MPSLSTMDescriptor.SetInputGateMemoryWeights]
//   - [MPSLSTMDescriptor.InputGateRecurrentWeights]
//   - [MPSLSTMDescriptor.SetInputGateRecurrentWeights]
//   - [MPSLSTMDescriptor.MemoryWeightsAreDiagonal]
//   - [MPSLSTMDescriptor.SetMemoryWeightsAreDiagonal]
//   - [MPSLSTMDescriptor.OutputGateInputWeights]
//   - [MPSLSTMDescriptor.SetOutputGateInputWeights]
//   - [MPSLSTMDescriptor.OutputGateMemoryWeights]
//   - [MPSLSTMDescriptor.SetOutputGateMemoryWeights]
//   - [MPSLSTMDescriptor.OutputGateRecurrentWeights]
//   - [MPSLSTMDescriptor.SetOutputGateRecurrentWeights]
//   - [MPSLSTMDescriptor.CellToOutputNeuronParamC]
//   - [MPSLSTMDescriptor.SetCellToOutputNeuronParamC]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor
type MPSLSTMDescriptor struct {
	MPSRNNDescriptor
}

// MPSLSTMDescriptorFromID constructs a [MPSLSTMDescriptor] from an objc.ID.
//
// A description of a long short-term memory block or layer.
func MPSLSTMDescriptorFromID(id objc.ID) MPSLSTMDescriptor {
	return MPSLSTMDescriptor{MPSRNNDescriptor: MPSRNNDescriptorFromID(id)}
}

// NOTE: MPSLSTMDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSLSTMDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSLSTMDescriptor.CellGateInputWeights]
//   - [IMPSLSTMDescriptor.SetCellGateInputWeights]
//   - [IMPSLSTMDescriptor.CellGateMemoryWeights]
//   - [IMPSLSTMDescriptor.SetCellGateMemoryWeights]
//   - [IMPSLSTMDescriptor.CellGateRecurrentWeights]
//   - [IMPSLSTMDescriptor.SetCellGateRecurrentWeights]
//   - [IMPSLSTMDescriptor.CellToOutputNeuronParamA]
//   - [IMPSLSTMDescriptor.SetCellToOutputNeuronParamA]
//   - [IMPSLSTMDescriptor.CellToOutputNeuronParamB]
//   - [IMPSLSTMDescriptor.SetCellToOutputNeuronParamB]
//   - [IMPSLSTMDescriptor.CellToOutputNeuronType]
//   - [IMPSLSTMDescriptor.SetCellToOutputNeuronType]
//   - [IMPSLSTMDescriptor.ForgetGateInputWeights]
//   - [IMPSLSTMDescriptor.SetForgetGateInputWeights]
//   - [IMPSLSTMDescriptor.ForgetGateMemoryWeights]
//   - [IMPSLSTMDescriptor.SetForgetGateMemoryWeights]
//   - [IMPSLSTMDescriptor.ForgetGateRecurrentWeights]
//   - [IMPSLSTMDescriptor.SetForgetGateRecurrentWeights]
//   - [IMPSLSTMDescriptor.InputGateInputWeights]
//   - [IMPSLSTMDescriptor.SetInputGateInputWeights]
//   - [IMPSLSTMDescriptor.InputGateMemoryWeights]
//   - [IMPSLSTMDescriptor.SetInputGateMemoryWeights]
//   - [IMPSLSTMDescriptor.InputGateRecurrentWeights]
//   - [IMPSLSTMDescriptor.SetInputGateRecurrentWeights]
//   - [IMPSLSTMDescriptor.MemoryWeightsAreDiagonal]
//   - [IMPSLSTMDescriptor.SetMemoryWeightsAreDiagonal]
//   - [IMPSLSTMDescriptor.OutputGateInputWeights]
//   - [IMPSLSTMDescriptor.SetOutputGateInputWeights]
//   - [IMPSLSTMDescriptor.OutputGateMemoryWeights]
//   - [IMPSLSTMDescriptor.SetOutputGateMemoryWeights]
//   - [IMPSLSTMDescriptor.OutputGateRecurrentWeights]
//   - [IMPSLSTMDescriptor.SetOutputGateRecurrentWeights]
//   - [IMPSLSTMDescriptor.CellToOutputNeuronParamC]
//   - [IMPSLSTMDescriptor.SetCellToOutputNeuronParamC]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor
type IMPSLSTMDescriptor interface {
	IMPSRNNDescriptor

	// Topic: Instance Properties

	CellGateInputWeights() MPSCNNConvolutionDataSource
	SetCellGateInputWeights(value MPSCNNConvolutionDataSource)
	CellGateMemoryWeights() MPSCNNConvolutionDataSource
	SetCellGateMemoryWeights(value MPSCNNConvolutionDataSource)
	CellGateRecurrentWeights() MPSCNNConvolutionDataSource
	SetCellGateRecurrentWeights(value MPSCNNConvolutionDataSource)
	CellToOutputNeuronParamA() float32
	SetCellToOutputNeuronParamA(value float32)
	CellToOutputNeuronParamB() float32
	SetCellToOutputNeuronParamB(value float32)
	CellToOutputNeuronType() MPSCNNNeuronType
	SetCellToOutputNeuronType(value MPSCNNNeuronType)
	ForgetGateInputWeights() MPSCNNConvolutionDataSource
	SetForgetGateInputWeights(value MPSCNNConvolutionDataSource)
	ForgetGateMemoryWeights() MPSCNNConvolutionDataSource
	SetForgetGateMemoryWeights(value MPSCNNConvolutionDataSource)
	ForgetGateRecurrentWeights() MPSCNNConvolutionDataSource
	SetForgetGateRecurrentWeights(value MPSCNNConvolutionDataSource)
	InputGateInputWeights() MPSCNNConvolutionDataSource
	SetInputGateInputWeights(value MPSCNNConvolutionDataSource)
	InputGateMemoryWeights() MPSCNNConvolutionDataSource
	SetInputGateMemoryWeights(value MPSCNNConvolutionDataSource)
	InputGateRecurrentWeights() MPSCNNConvolutionDataSource
	SetInputGateRecurrentWeights(value MPSCNNConvolutionDataSource)
	MemoryWeightsAreDiagonal() bool
	SetMemoryWeightsAreDiagonal(value bool)
	OutputGateInputWeights() MPSCNNConvolutionDataSource
	SetOutputGateInputWeights(value MPSCNNConvolutionDataSource)
	OutputGateMemoryWeights() MPSCNNConvolutionDataSource
	SetOutputGateMemoryWeights(value MPSCNNConvolutionDataSource)
	OutputGateRecurrentWeights() MPSCNNConvolutionDataSource
	SetOutputGateRecurrentWeights(value MPSCNNConvolutionDataSource)
	CellToOutputNeuronParamC() float32
	SetCellToOutputNeuronParamC(value float32)
}

// Init initializes the instance.
func (l MPSLSTMDescriptor) Init() MPSLSTMDescriptor {
	rv := objc.Send[MPSLSTMDescriptor](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MPSLSTMDescriptor) Autorelease() MPSLSTMDescriptor {
	rv := objc.Send[MPSLSTMDescriptor](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSLSTMDescriptor creates a new MPSLSTMDescriptor instance.
func NewMPSLSTMDescriptor() MPSLSTMDescriptor {
	class := getMPSLSTMDescriptorClass()
	rv := objc.Send[MPSLSTMDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor/createLSTMDescriptor(withInputFeatureChannels:outputFeatureChannels:)
func (_MPSLSTMDescriptorClass MPSLSTMDescriptorClass) CreateLSTMDescriptorWithInputFeatureChannelsOutputFeatureChannels(inputFeatureChannels uint, outputFeatureChannels uint) MPSLSTMDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSLSTMDescriptorClass.class), objc.Sel("createLSTMDescriptorWithInputFeatureChannels:outputFeatureChannels:"), inputFeatureChannels, outputFeatureChannels)
	return MPSLSTMDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor/cellGateInputWeights
func (l MPSLSTMDescriptor) CellGateInputWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("cellGateInputWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (l MPSLSTMDescriptor) SetCellGateInputWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](l.ID, objc.Sel("setCellGateInputWeights:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor/cellGateMemoryWeights
func (l MPSLSTMDescriptor) CellGateMemoryWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("cellGateMemoryWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (l MPSLSTMDescriptor) SetCellGateMemoryWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](l.ID, objc.Sel("setCellGateMemoryWeights:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor/cellGateRecurrentWeights
func (l MPSLSTMDescriptor) CellGateRecurrentWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("cellGateRecurrentWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (l MPSLSTMDescriptor) SetCellGateRecurrentWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](l.ID, objc.Sel("setCellGateRecurrentWeights:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor/cellToOutputNeuronParamA
func (l MPSLSTMDescriptor) CellToOutputNeuronParamA() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("cellToOutputNeuronParamA"))
	return rv
}
func (l MPSLSTMDescriptor) SetCellToOutputNeuronParamA(value float32) {
	objc.Send[struct{}](l.ID, objc.Sel("setCellToOutputNeuronParamA:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor/cellToOutputNeuronParamB
func (l MPSLSTMDescriptor) CellToOutputNeuronParamB() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("cellToOutputNeuronParamB"))
	return rv
}
func (l MPSLSTMDescriptor) SetCellToOutputNeuronParamB(value float32) {
	objc.Send[struct{}](l.ID, objc.Sel("setCellToOutputNeuronParamB:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor/cellToOutputNeuronType
func (l MPSLSTMDescriptor) CellToOutputNeuronType() MPSCNNNeuronType {
	rv := objc.Send[MPSCNNNeuronType](l.ID, objc.Sel("cellToOutputNeuronType"))
	return MPSCNNNeuronType(rv)
}
func (l MPSLSTMDescriptor) SetCellToOutputNeuronType(value MPSCNNNeuronType) {
	objc.Send[struct{}](l.ID, objc.Sel("setCellToOutputNeuronType:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor/forgetGateInputWeights
func (l MPSLSTMDescriptor) ForgetGateInputWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("forgetGateInputWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (l MPSLSTMDescriptor) SetForgetGateInputWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](l.ID, objc.Sel("setForgetGateInputWeights:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor/forgetGateMemoryWeights
func (l MPSLSTMDescriptor) ForgetGateMemoryWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("forgetGateMemoryWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (l MPSLSTMDescriptor) SetForgetGateMemoryWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](l.ID, objc.Sel("setForgetGateMemoryWeights:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor/forgetGateRecurrentWeights
func (l MPSLSTMDescriptor) ForgetGateRecurrentWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("forgetGateRecurrentWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (l MPSLSTMDescriptor) SetForgetGateRecurrentWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](l.ID, objc.Sel("setForgetGateRecurrentWeights:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor/inputGateInputWeights
func (l MPSLSTMDescriptor) InputGateInputWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("inputGateInputWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (l MPSLSTMDescriptor) SetInputGateInputWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](l.ID, objc.Sel("setInputGateInputWeights:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor/inputGateMemoryWeights
func (l MPSLSTMDescriptor) InputGateMemoryWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("inputGateMemoryWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (l MPSLSTMDescriptor) SetInputGateMemoryWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](l.ID, objc.Sel("setInputGateMemoryWeights:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor/inputGateRecurrentWeights
func (l MPSLSTMDescriptor) InputGateRecurrentWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("inputGateRecurrentWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (l MPSLSTMDescriptor) SetInputGateRecurrentWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](l.ID, objc.Sel("setInputGateRecurrentWeights:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor/memoryWeightsAreDiagonal
func (l MPSLSTMDescriptor) MemoryWeightsAreDiagonal() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("memoryWeightsAreDiagonal"))
	return rv
}
func (l MPSLSTMDescriptor) SetMemoryWeightsAreDiagonal(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setMemoryWeightsAreDiagonal:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor/outputGateInputWeights
func (l MPSLSTMDescriptor) OutputGateInputWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("outputGateInputWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (l MPSLSTMDescriptor) SetOutputGateInputWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](l.ID, objc.Sel("setOutputGateInputWeights:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor/outputGateMemoryWeights
func (l MPSLSTMDescriptor) OutputGateMemoryWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("outputGateMemoryWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (l MPSLSTMDescriptor) SetOutputGateMemoryWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](l.ID, objc.Sel("setOutputGateMemoryWeights:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor/outputGateRecurrentWeights
func (l MPSLSTMDescriptor) OutputGateRecurrentWeights() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("outputGateRecurrentWeights"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}
func (l MPSLSTMDescriptor) SetOutputGateRecurrentWeights(value MPSCNNConvolutionDataSource) {
	objc.Send[struct{}](l.ID, objc.Sel("setOutputGateRecurrentWeights:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor/cellToOutputNeuronParamC
func (l MPSLSTMDescriptor) CellToOutputNeuronParamC() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("cellToOutputNeuronParamC"))
	return rv
}
func (l MPSLSTMDescriptor) SetCellToOutputNeuronParamC(value float32) {
	objc.Send[struct{}](l.ID, objc.Sel("setCellToOutputNeuronParamC:"), value)
}
