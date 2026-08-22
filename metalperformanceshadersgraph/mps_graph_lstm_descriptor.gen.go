// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphLSTMDescriptor] class.
var (
	_MPSGraphLSTMDescriptorClass     MPSGraphLSTMDescriptorClass
	_MPSGraphLSTMDescriptorClassOnce sync.Once
)

func getMPSGraphLSTMDescriptorClass() MPSGraphLSTMDescriptorClass {
	_MPSGraphLSTMDescriptorClassOnce.Do(func() {
		_MPSGraphLSTMDescriptorClass = MPSGraphLSTMDescriptorClass{class: objc.GetClass("MPSGraphLSTMDescriptor")}
	})
	return _MPSGraphLSTMDescriptorClass
}

// GetMPSGraphLSTMDescriptorClass returns the class object for MPSGraphLSTMDescriptor.
func GetMPSGraphLSTMDescriptorClass() MPSGraphLSTMDescriptorClass {
	return getMPSGraphLSTMDescriptorClass()
}

type MPSGraphLSTMDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphLSTMDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphLSTMDescriptorClass) Alloc() MPSGraphLSTMDescriptor {
	rv := objc.Send[MPSGraphLSTMDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The class that defines the parameters for a long short-term memory (LSTM)
// operation.
//
// # Overview
//
// Use this descriptor with the following [MPSGraph] methods:
//
// -
// [MPSGraph.LSTMWithSourceTensorRecurrentWeightInitStateInitCellDescriptorName]
// -
// [MPSGraph.LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellDescriptorName]
// -
// [MPSGraph.LSTMWithSourceTensorRecurrentWeightInputWeightBiasInitStateInitCellMaskPeepholeDescriptorName]
// -
// [MPSGraph.LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdDescriptorName]
// -
// [MPSGraph.LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdInputWeightBiasInitStateInitCellDescriptorName]
// -
// [MPSGraph.LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdInputWeightBiasInitStateInitCellMaskDescriptorName]
// -
// [MPSGraph.LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdStateGradientCellGradientInputWeightBiasInitStateInitCellMaskPeepholeDescriptorName]
//
// # Instance Properties
//
//   - [MPSGraphLSTMDescriptor.Activation]: A parameter that defines the activation function used with the current cell value of the LSTM operation.
//   - [MPSGraphLSTMDescriptor.SetActivation]
//   - [MPSGraphLSTMDescriptor.Bidirectional]: A parameter that defines a bidirectional LSTM layer.
//   - [MPSGraphLSTMDescriptor.SetBidirectional]
//   - [MPSGraphLSTMDescriptor.CellGateActivation]: A parameter that defines the activation function used with the cell gate of the LSTM operation.
//   - [MPSGraphLSTMDescriptor.SetCellGateActivation]
//   - [MPSGraphLSTMDescriptor.ForgetGateActivation]: A parameter that defines the activation function used with the forget gate of the LSTM operation.
//   - [MPSGraphLSTMDescriptor.SetForgetGateActivation]
//   - [MPSGraphLSTMDescriptor.ForgetGateLast]: A parameter that controls the internal order of the LSTM gates.
//   - [MPSGraphLSTMDescriptor.SetForgetGateLast]
//   - [MPSGraphLSTMDescriptor.InputGateActivation]: A parameter that defines the activation function used with the input gate of the LSTM operation.
//   - [MPSGraphLSTMDescriptor.SetInputGateActivation]
//   - [MPSGraphLSTMDescriptor.OutputGateActivation]: A parameter that defines the activation function used with the output gate of the LSTM operation.
//   - [MPSGraphLSTMDescriptor.SetOutputGateActivation]
//   - [MPSGraphLSTMDescriptor.ProduceCell]: A parameter that controls whether or not to return the output cell from the LSTM layer.
//   - [MPSGraphLSTMDescriptor.SetProduceCell]
//   - [MPSGraphLSTMDescriptor.Reverse]: A parameter that defines time direction of the input sequence.
//   - [MPSGraphLSTMDescriptor.SetReverse]
//   - [MPSGraphLSTMDescriptor.Training]: A parameter that enables the LSTM layer to support training.
//   - [MPSGraphLSTMDescriptor.SetTraining]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor
type MPSGraphLSTMDescriptor struct {
	MPSGraphObject
}

// MPSGraphLSTMDescriptorFromID constructs a [MPSGraphLSTMDescriptor] from an objc.ID.
//
// The class that defines the parameters for a long short-term memory (LSTM)
// operation.
func MPSGraphLSTMDescriptorFromID(id objc.ID) MPSGraphLSTMDescriptor {
	return MPSGraphLSTMDescriptor{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphLSTMDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphLSTMDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSGraphLSTMDescriptor.Activation]: A parameter that defines the activation function used with the current cell value of the LSTM operation.
//   - [IMPSGraphLSTMDescriptor.SetActivation]
//   - [IMPSGraphLSTMDescriptor.Bidirectional]: A parameter that defines a bidirectional LSTM layer.
//   - [IMPSGraphLSTMDescriptor.SetBidirectional]
//   - [IMPSGraphLSTMDescriptor.CellGateActivation]: A parameter that defines the activation function used with the cell gate of the LSTM operation.
//   - [IMPSGraphLSTMDescriptor.SetCellGateActivation]
//   - [IMPSGraphLSTMDescriptor.ForgetGateActivation]: A parameter that defines the activation function used with the forget gate of the LSTM operation.
//   - [IMPSGraphLSTMDescriptor.SetForgetGateActivation]
//   - [IMPSGraphLSTMDescriptor.ForgetGateLast]: A parameter that controls the internal order of the LSTM gates.
//   - [IMPSGraphLSTMDescriptor.SetForgetGateLast]
//   - [IMPSGraphLSTMDescriptor.InputGateActivation]: A parameter that defines the activation function used with the input gate of the LSTM operation.
//   - [IMPSGraphLSTMDescriptor.SetInputGateActivation]
//   - [IMPSGraphLSTMDescriptor.OutputGateActivation]: A parameter that defines the activation function used with the output gate of the LSTM operation.
//   - [IMPSGraphLSTMDescriptor.SetOutputGateActivation]
//   - [IMPSGraphLSTMDescriptor.ProduceCell]: A parameter that controls whether or not to return the output cell from the LSTM layer.
//   - [IMPSGraphLSTMDescriptor.SetProduceCell]
//   - [IMPSGraphLSTMDescriptor.Reverse]: A parameter that defines time direction of the input sequence.
//   - [IMPSGraphLSTMDescriptor.SetReverse]
//   - [IMPSGraphLSTMDescriptor.Training]: A parameter that enables the LSTM layer to support training.
//   - [IMPSGraphLSTMDescriptor.SetTraining]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor
type IMPSGraphLSTMDescriptor interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// A parameter that defines the activation function used with the current cell value of the LSTM operation.
	Activation() MPSGraphRNNActivation
	SetActivation(value MPSGraphRNNActivation)
	// A parameter that defines a bidirectional LSTM layer.
	Bidirectional() bool
	SetBidirectional(value bool)
	// A parameter that defines the activation function used with the cell gate of the LSTM operation.
	CellGateActivation() MPSGraphRNNActivation
	SetCellGateActivation(value MPSGraphRNNActivation)
	// A parameter that defines the activation function used with the forget gate of the LSTM operation.
	ForgetGateActivation() MPSGraphRNNActivation
	SetForgetGateActivation(value MPSGraphRNNActivation)
	// A parameter that controls the internal order of the LSTM gates.
	ForgetGateLast() bool
	SetForgetGateLast(value bool)
	// A parameter that defines the activation function used with the input gate of the LSTM operation.
	InputGateActivation() MPSGraphRNNActivation
	SetInputGateActivation(value MPSGraphRNNActivation)
	// A parameter that defines the activation function used with the output gate of the LSTM operation.
	OutputGateActivation() MPSGraphRNNActivation
	SetOutputGateActivation(value MPSGraphRNNActivation)
	// A parameter that controls whether or not to return the output cell from the LSTM layer.
	ProduceCell() bool
	SetProduceCell(value bool)
	// A parameter that defines time direction of the input sequence.
	Reverse() bool
	SetReverse(value bool)
	// A parameter that enables the LSTM layer to support training.
	Training() bool
	SetTraining(value bool)
}

// Init initializes the instance.
func (g MPSGraphLSTMDescriptor) Init() MPSGraphLSTMDescriptor {
	rv := objc.Send[MPSGraphLSTMDescriptor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphLSTMDescriptor) Autorelease() MPSGraphLSTMDescriptor {
	rv := objc.Send[MPSGraphLSTMDescriptor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphLSTMDescriptor creates a new MPSGraphLSTMDescriptor instance.
func NewMPSGraphLSTMDescriptor() MPSGraphLSTMDescriptor {
	class := getMPSGraphLSTMDescriptorClass()
	rv := objc.Send[MPSGraphLSTMDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an LSTM descriptor with default values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/descriptor
func (_MPSGraphLSTMDescriptorClass MPSGraphLSTMDescriptorClass) Descriptor() MPSGraphLSTMDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSGraphLSTMDescriptorClass.class), objc.Sel("descriptor"))
	return MPSGraphLSTMDescriptorFromID(rv)
}

// A parameter that defines the activation function used with the current cell
// value of the LSTM operation.
//
// # Discussion
//
// Default value: [MPSGraphRNNActivationTanh].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/activation
func (g MPSGraphLSTMDescriptor) Activation() MPSGraphRNNActivation {
	rv := objc.Send[MPSGraphRNNActivation](g.ID, objc.Sel("activation"))
	return MPSGraphRNNActivation(rv)
}
func (g MPSGraphLSTMDescriptor) SetActivation(value MPSGraphRNNActivation) {
	objc.Send[struct{}](g.ID, objc.Sel("setActivation:"), value)
}

// A parameter that defines a bidirectional LSTM layer.
//
// # Discussion
//
// If set to [YES] then the input sequence is traversed in both directions and
// the two results are concatenated together on the channel-axis. Default
// value: [NO].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/bidirectional
func (g MPSGraphLSTMDescriptor) Bidirectional() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("bidirectional"))
	return rv
}
func (g MPSGraphLSTMDescriptor) SetBidirectional(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setBidirectional:"), value)
}

// A parameter that defines the activation function used with the cell gate of
// the LSTM operation.
//
// # Discussion
//
// Default value: [MPSGraphRNNActivationTanh].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/cellGateActivation
func (g MPSGraphLSTMDescriptor) CellGateActivation() MPSGraphRNNActivation {
	rv := objc.Send[MPSGraphRNNActivation](g.ID, objc.Sel("cellGateActivation"))
	return MPSGraphRNNActivation(rv)
}
func (g MPSGraphLSTMDescriptor) SetCellGateActivation(value MPSGraphRNNActivation) {
	objc.Send[struct{}](g.ID, objc.Sel("setCellGateActivation:"), value)
}

// A parameter that defines the activation function used with the forget gate
// of the LSTM operation.
//
// # Discussion
//
// Default value: [MPSGraphRNNActivationSigmoid].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/forgetGateActivation
func (g MPSGraphLSTMDescriptor) ForgetGateActivation() MPSGraphRNNActivation {
	rv := objc.Send[MPSGraphRNNActivation](g.ID, objc.Sel("forgetGateActivation"))
	return MPSGraphRNNActivation(rv)
}
func (g MPSGraphLSTMDescriptor) SetForgetGateActivation(value MPSGraphRNNActivation) {
	objc.Send[struct{}](g.ID, objc.Sel("setForgetGateActivation:"), value)
}

// A parameter that controls the internal order of the LSTM gates.
//
// # Discussion
//
// If set to [YES] then the layer will use the gate-ordering `[ i, z, f, o ]`
// instead of default `[ i, f, z, o ]`. Default value: [NO]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/forgetGateLast
func (g MPSGraphLSTMDescriptor) ForgetGateLast() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("forgetGateLast"))
	return rv
}
func (g MPSGraphLSTMDescriptor) SetForgetGateLast(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setForgetGateLast:"), value)
}

// A parameter that defines the activation function used with the input gate
// of the LSTM operation.
//
// # Discussion
//
// Default value: [MPSGraphRNNActivationSigmoid].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/inputGateActivation
func (g MPSGraphLSTMDescriptor) InputGateActivation() MPSGraphRNNActivation {
	rv := objc.Send[MPSGraphRNNActivation](g.ID, objc.Sel("inputGateActivation"))
	return MPSGraphRNNActivation(rv)
}
func (g MPSGraphLSTMDescriptor) SetInputGateActivation(value MPSGraphRNNActivation) {
	objc.Send[struct{}](g.ID, objc.Sel("setInputGateActivation:"), value)
}

// A parameter that defines the activation function used with the output gate
// of the LSTM operation.
//
// # Discussion
//
// Default value: [MPSGraphRNNActivationSigmoid].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/outputGateActivation
func (g MPSGraphLSTMDescriptor) OutputGateActivation() MPSGraphRNNActivation {
	rv := objc.Send[MPSGraphRNNActivation](g.ID, objc.Sel("outputGateActivation"))
	return MPSGraphRNNActivation(rv)
}
func (g MPSGraphLSTMDescriptor) SetOutputGateActivation(value MPSGraphRNNActivation) {
	objc.Send[struct{}](g.ID, objc.Sel("setOutputGateActivation:"), value)
}

// A parameter that controls whether or not to return the output cell from the
// LSTM layer.
//
// # Discussion
//
// If set to [YES] then this layer will produce the internal cell of the LSTM
// unit as secondary output. Default value: [NO].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/produceCell
func (g MPSGraphLSTMDescriptor) ProduceCell() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("produceCell"))
	return rv
}
func (g MPSGraphLSTMDescriptor) SetProduceCell(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setProduceCell:"), value)
}

// A parameter that defines time direction of the input sequence.
//
// # Discussion
//
// If set to [YES] then the input sequence is passed in reverse time order to
// the layer. Note: Ignored when `bidirectional = YES`. Default value: [NO].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/reverse
func (g MPSGraphLSTMDescriptor) Reverse() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("reverse"))
	return rv
}
func (g MPSGraphLSTMDescriptor) SetReverse(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setReverse:"), value)
}

// A parameter that enables the LSTM layer to support training.
//
// # Discussion
//
// If set to [YES] then the layer will produce training state tensor as a
// secondary output. Default value: [NO].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/training
func (g MPSGraphLSTMDescriptor) Training() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("training"))
	return rv
}
func (g MPSGraphLSTMDescriptor) SetTraining(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setTraining:"), value)
}
