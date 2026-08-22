// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphSingleGateRNNDescriptor] class.
var (
	_MPSGraphSingleGateRNNDescriptorClass     MPSGraphSingleGateRNNDescriptorClass
	_MPSGraphSingleGateRNNDescriptorClassOnce sync.Once
)

func getMPSGraphSingleGateRNNDescriptorClass() MPSGraphSingleGateRNNDescriptorClass {
	_MPSGraphSingleGateRNNDescriptorClassOnce.Do(func() {
		_MPSGraphSingleGateRNNDescriptorClass = MPSGraphSingleGateRNNDescriptorClass{class: objc.GetClass("MPSGraphSingleGateRNNDescriptor")}
	})
	return _MPSGraphSingleGateRNNDescriptorClass
}

// GetMPSGraphSingleGateRNNDescriptorClass returns the class object for MPSGraphSingleGateRNNDescriptor.
func GetMPSGraphSingleGateRNNDescriptorClass() MPSGraphSingleGateRNNDescriptorClass {
	return getMPSGraphSingleGateRNNDescriptorClass()
}

type MPSGraphSingleGateRNNDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphSingleGateRNNDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphSingleGateRNNDescriptorClass) Alloc() MPSGraphSingleGateRNNDescriptor {
	rv := objc.Send[MPSGraphSingleGateRNNDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The class that defines the parameters for a single gate RNN operation.
//
// # Overview
//
// Use this descriptor with the following [MPSGraph] methods:
//
// -
// [MPSGraph.SingleGateRNNWithSourceTensorRecurrentWeightInitStateDescriptorName]
// -
// [MPSGraph.SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName]
// -
// [MPSGraph.SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskDescriptorName]
// -
// [MPSGraph.SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInitStateDescriptorName]
// -
// [MPSGraph.SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInputWeightBiasInitStateDescriptorName]
// -
// [MPSGraph.SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInputWeightBiasInitStateMaskDescriptorName]
// -
// [MPSGraph.SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateStateGradientInputWeightBiasInitStateMaskDescriptorName]
//
// # Instance Properties
//
//   - [MPSGraphSingleGateRNNDescriptor.Activation]: A parameter that defines the activation function to use with the RNN operation.
//   - [MPSGraphSingleGateRNNDescriptor.SetActivation]
//   - [MPSGraphSingleGateRNNDescriptor.Bidirectional]: A parameter that defines a bidirectional RNN layer.
//   - [MPSGraphSingleGateRNNDescriptor.SetBidirectional]
//   - [MPSGraphSingleGateRNNDescriptor.Reverse]: A parameter that defines time direction of the input sequence.
//   - [MPSGraphSingleGateRNNDescriptor.SetReverse]
//   - [MPSGraphSingleGateRNNDescriptor.Training]: A parameter that makes the RNN layer support training.
//   - [MPSGraphSingleGateRNNDescriptor.SetTraining]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSingleGateRNNDescriptor
type MPSGraphSingleGateRNNDescriptor struct {
	MPSGraphObject
}

// MPSGraphSingleGateRNNDescriptorFromID constructs a [MPSGraphSingleGateRNNDescriptor] from an objc.ID.
//
// The class that defines the parameters for a single gate RNN operation.
func MPSGraphSingleGateRNNDescriptorFromID(id objc.ID) MPSGraphSingleGateRNNDescriptor {
	return MPSGraphSingleGateRNNDescriptor{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphSingleGateRNNDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphSingleGateRNNDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSGraphSingleGateRNNDescriptor.Activation]: A parameter that defines the activation function to use with the RNN operation.
//   - [IMPSGraphSingleGateRNNDescriptor.SetActivation]
//   - [IMPSGraphSingleGateRNNDescriptor.Bidirectional]: A parameter that defines a bidirectional RNN layer.
//   - [IMPSGraphSingleGateRNNDescriptor.SetBidirectional]
//   - [IMPSGraphSingleGateRNNDescriptor.Reverse]: A parameter that defines time direction of the input sequence.
//   - [IMPSGraphSingleGateRNNDescriptor.SetReverse]
//   - [IMPSGraphSingleGateRNNDescriptor.Training]: A parameter that makes the RNN layer support training.
//   - [IMPSGraphSingleGateRNNDescriptor.SetTraining]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSingleGateRNNDescriptor
type IMPSGraphSingleGateRNNDescriptor interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// A parameter that defines the activation function to use with the RNN operation.
	Activation() MPSGraphRNNActivation
	SetActivation(value MPSGraphRNNActivation)
	// A parameter that defines a bidirectional RNN layer.
	Bidirectional() bool
	SetBidirectional(value bool)
	// A parameter that defines time direction of the input sequence.
	Reverse() bool
	SetReverse(value bool)
	// A parameter that makes the RNN layer support training.
	Training() bool
	SetTraining(value bool)
}

// Init initializes the instance.
func (g MPSGraphSingleGateRNNDescriptor) Init() MPSGraphSingleGateRNNDescriptor {
	rv := objc.Send[MPSGraphSingleGateRNNDescriptor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphSingleGateRNNDescriptor) Autorelease() MPSGraphSingleGateRNNDescriptor {
	rv := objc.Send[MPSGraphSingleGateRNNDescriptor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphSingleGateRNNDescriptor creates a new MPSGraphSingleGateRNNDescriptor instance.
func NewMPSGraphSingleGateRNNDescriptor() MPSGraphSingleGateRNNDescriptor {
	class := getMPSGraphSingleGateRNNDescriptorClass()
	rv := objc.Send[MPSGraphSingleGateRNNDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a single gate RNN descriptor with default values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSingleGateRNNDescriptor/descriptor
func (_MPSGraphSingleGateRNNDescriptorClass MPSGraphSingleGateRNNDescriptorClass) Descriptor() MPSGraphSingleGateRNNDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSGraphSingleGateRNNDescriptorClass.class), objc.Sel("descriptor"))
	return MPSGraphSingleGateRNNDescriptorFromID(rv)
}

// A parameter that defines the activation function to use with the RNN
// operation.
//
// # Discussion
//
// Default value: [MPSGraphRNNActivationRelu].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSingleGateRNNDescriptor/activation
func (g MPSGraphSingleGateRNNDescriptor) Activation() MPSGraphRNNActivation {
	rv := objc.Send[MPSGraphRNNActivation](g.ID, objc.Sel("activation"))
	return MPSGraphRNNActivation(rv)
}
func (g MPSGraphSingleGateRNNDescriptor) SetActivation(value MPSGraphRNNActivation) {
	objc.Send[struct{}](g.ID, objc.Sel("setActivation:"), value)
}

// A parameter that defines a bidirectional RNN layer.
//
// # Discussion
//
// If set to [YES] then the input sequence is traversed in both directions and
// the two results are concatenated together on the channel-axis. Default
// value: [NO].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSingleGateRNNDescriptor/bidirectional
func (g MPSGraphSingleGateRNNDescriptor) Bidirectional() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("bidirectional"))
	return rv
}
func (g MPSGraphSingleGateRNNDescriptor) SetBidirectional(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setBidirectional:"), value)
}

// A parameter that defines time direction of the input sequence.
//
// # Discussion
//
// If set to [YES] then the input sequence is passed in reverse time order to
// the layer. Note: Ignored when `bidirectional = YES`. Default value: [NO].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSingleGateRNNDescriptor/reverse
func (g MPSGraphSingleGateRNNDescriptor) Reverse() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("reverse"))
	return rv
}
func (g MPSGraphSingleGateRNNDescriptor) SetReverse(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setReverse:"), value)
}

// A parameter that makes the RNN layer support training.
//
// # Discussion
//
// If set to [YES] then the layer will produce training state tensor as a
// secondary output. Default value: [NO].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphSingleGateRNNDescriptor/training
func (g MPSGraphSingleGateRNNDescriptor) Training() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("training"))
	return rv
}
func (g MPSGraphSingleGateRNNDescriptor) SetTraining(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setTraining:"), value)
}
