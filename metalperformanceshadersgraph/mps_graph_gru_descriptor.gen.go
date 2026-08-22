// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphGRUDescriptor] class.
var (
	_MPSGraphGRUDescriptorClass     MPSGraphGRUDescriptorClass
	_MPSGraphGRUDescriptorClassOnce sync.Once
)

func getMPSGraphGRUDescriptorClass() MPSGraphGRUDescriptorClass {
	_MPSGraphGRUDescriptorClassOnce.Do(func() {
		_MPSGraphGRUDescriptorClass = MPSGraphGRUDescriptorClass{class: objc.GetClass("MPSGraphGRUDescriptor")}
	})
	return _MPSGraphGRUDescriptorClass
}

// GetMPSGraphGRUDescriptorClass returns the class object for MPSGraphGRUDescriptor.
func GetMPSGraphGRUDescriptorClass() MPSGraphGRUDescriptorClass {
	return getMPSGraphGRUDescriptorClass()
}

type MPSGraphGRUDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphGRUDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphGRUDescriptorClass) Alloc() MPSGraphGRUDescriptor {
	rv := objc.Send[MPSGraphGRUDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The class that defines the parameters for a gated recurrent unit (GRU)
// operation.
//
// # Overview
//
// Use this descriptor with the following [MPSGraph] methods:
//
// -
// [MPSGraph.GRUWithSourceTensorRecurrentWeightInputWeightBiasDescriptorName]
// -
// [MPSGraph.GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateDescriptorName]
// -
// [MPSGraph.GRUWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskSecondaryBiasDescriptorName]
// -
// [MPSGraph.GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdInputWeightBiasDescriptorName]
// -
// [MPSGraph.GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdInputWeightBiasInitStateDescriptorName]
// -
// [MPSGraph.GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdStateGradientInputWeightBiasInitStateMaskSecondaryBiasDescriptorName]
//
// # Instance Properties
//
//   - [MPSGraphGRUDescriptor.Bidirectional]: A parameter that defines a bidirectional GRU layer.
//   - [MPSGraphGRUDescriptor.SetBidirectional]
//   - [MPSGraphGRUDescriptor.FlipZ]: A parameter that chooses between two variants for the final output computation.
//   - [MPSGraphGRUDescriptor.SetFlipZ]
//   - [MPSGraphGRUDescriptor.OutputGateActivation]: A parameter that defines the activation function to use with the output-gate of the GRU operation.
//   - [MPSGraphGRUDescriptor.SetOutputGateActivation]
//   - [MPSGraphGRUDescriptor.ResetAfter]: A parameter that chooses between two variants for the reset gate computation.
//   - [MPSGraphGRUDescriptor.SetResetAfter]
//   - [MPSGraphGRUDescriptor.ResetGateActivation]: A parameter that defines the activation function to use with the reset-gate of the GRU operation.
//   - [MPSGraphGRUDescriptor.SetResetGateActivation]
//   - [MPSGraphGRUDescriptor.ResetGateFirst]: A parameter that controls the internal order of the GRU gates.
//   - [MPSGraphGRUDescriptor.SetResetGateFirst]
//   - [MPSGraphGRUDescriptor.Reverse]: A parameter that defines the time direction of the input sequence.
//   - [MPSGraphGRUDescriptor.SetReverse]
//   - [MPSGraphGRUDescriptor.Training]: A parameter that enables the GRU layer to support training.
//   - [MPSGraphGRUDescriptor.SetTraining]
//   - [MPSGraphGRUDescriptor.UpdateGateActivation]: A parameter that defines the activation function to use with the update-gate of the GRU operation.
//   - [MPSGraphGRUDescriptor.SetUpdateGateActivation]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor
type MPSGraphGRUDescriptor struct {
	MPSGraphObject
}

// MPSGraphGRUDescriptorFromID constructs a [MPSGraphGRUDescriptor] from an objc.ID.
//
// The class that defines the parameters for a gated recurrent unit (GRU)
// operation.
func MPSGraphGRUDescriptorFromID(id objc.ID) MPSGraphGRUDescriptor {
	return MPSGraphGRUDescriptor{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphGRUDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphGRUDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSGraphGRUDescriptor.Bidirectional]: A parameter that defines a bidirectional GRU layer.
//   - [IMPSGraphGRUDescriptor.SetBidirectional]
//   - [IMPSGraphGRUDescriptor.FlipZ]: A parameter that chooses between two variants for the final output computation.
//   - [IMPSGraphGRUDescriptor.SetFlipZ]
//   - [IMPSGraphGRUDescriptor.OutputGateActivation]: A parameter that defines the activation function to use with the output-gate of the GRU operation.
//   - [IMPSGraphGRUDescriptor.SetOutputGateActivation]
//   - [IMPSGraphGRUDescriptor.ResetAfter]: A parameter that chooses between two variants for the reset gate computation.
//   - [IMPSGraphGRUDescriptor.SetResetAfter]
//   - [IMPSGraphGRUDescriptor.ResetGateActivation]: A parameter that defines the activation function to use with the reset-gate of the GRU operation.
//   - [IMPSGraphGRUDescriptor.SetResetGateActivation]
//   - [IMPSGraphGRUDescriptor.ResetGateFirst]: A parameter that controls the internal order of the GRU gates.
//   - [IMPSGraphGRUDescriptor.SetResetGateFirst]
//   - [IMPSGraphGRUDescriptor.Reverse]: A parameter that defines the time direction of the input sequence.
//   - [IMPSGraphGRUDescriptor.SetReverse]
//   - [IMPSGraphGRUDescriptor.Training]: A parameter that enables the GRU layer to support training.
//   - [IMPSGraphGRUDescriptor.SetTraining]
//   - [IMPSGraphGRUDescriptor.UpdateGateActivation]: A parameter that defines the activation function to use with the update-gate of the GRU operation.
//   - [IMPSGraphGRUDescriptor.SetUpdateGateActivation]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor
type IMPSGraphGRUDescriptor interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// A parameter that defines a bidirectional GRU layer.
	Bidirectional() bool
	SetBidirectional(value bool)
	// A parameter that chooses between two variants for the final output computation.
	FlipZ() bool
	SetFlipZ(value bool)
	// A parameter that defines the activation function to use with the output-gate of the GRU operation.
	OutputGateActivation() MPSGraphRNNActivation
	SetOutputGateActivation(value MPSGraphRNNActivation)
	// A parameter that chooses between two variants for the reset gate computation.
	ResetAfter() bool
	SetResetAfter(value bool)
	// A parameter that defines the activation function to use with the reset-gate of the GRU operation.
	ResetGateActivation() MPSGraphRNNActivation
	SetResetGateActivation(value MPSGraphRNNActivation)
	// A parameter that controls the internal order of the GRU gates.
	ResetGateFirst() bool
	SetResetGateFirst(value bool)
	// A parameter that defines the time direction of the input sequence.
	Reverse() bool
	SetReverse(value bool)
	// A parameter that enables the GRU layer to support training.
	Training() bool
	SetTraining(value bool)
	// A parameter that defines the activation function to use with the update-gate of the GRU operation.
	UpdateGateActivation() MPSGraphRNNActivation
	SetUpdateGateActivation(value MPSGraphRNNActivation)
}

// Init initializes the instance.
func (g MPSGraphGRUDescriptor) Init() MPSGraphGRUDescriptor {
	rv := objc.Send[MPSGraphGRUDescriptor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphGRUDescriptor) Autorelease() MPSGraphGRUDescriptor {
	rv := objc.Send[MPSGraphGRUDescriptor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphGRUDescriptor creates a new MPSGraphGRUDescriptor instance.
func NewMPSGraphGRUDescriptor() MPSGraphGRUDescriptor {
	class := getMPSGraphGRUDescriptorClass()
	rv := objc.Send[MPSGraphGRUDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an GRU descriptor with default values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/descriptor
func (_MPSGraphGRUDescriptorClass MPSGraphGRUDescriptorClass) Descriptor() MPSGraphGRUDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSGraphGRUDescriptorClass.class), objc.Sel("descriptor"))
	return MPSGraphGRUDescriptorFromID(rv)
}

// A parameter that defines a bidirectional GRU layer.
//
// # Discussion
//
// If set to [YES] then the input sequence is traversed in both directions and
// the two results are concatenated together on the channel-axis. Default
// value: [NO].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/bidirectional
func (g MPSGraphGRUDescriptor) Bidirectional() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("bidirectional"))
	return rv
}
func (g MPSGraphGRUDescriptor) SetBidirectional(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setBidirectional:"), value)
}

// A parameter that chooses between two variants for the final output
// computation.
//
// # Discussion
//
// If set to [YES] then the layer will compute the final value as `h[t] = z[t]
// h[t-1] + (1-z[t]) o[t]`. Otherwise it’s computed as `h[t] = (1-z[t])
// h[t-1] + z[t] o[t]`. Default value: [NO].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/flipZ
func (g MPSGraphGRUDescriptor) FlipZ() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("flipZ"))
	return rv
}
func (g MPSGraphGRUDescriptor) SetFlipZ(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setFlipZ:"), value)
}

// A parameter that defines the activation function to use with the
// output-gate of the GRU operation.
//
// # Discussion
//
// Default value: [MPSGraphRNNActivationTanh].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/outputGateActivation
func (g MPSGraphGRUDescriptor) OutputGateActivation() MPSGraphRNNActivation {
	rv := objc.Send[MPSGraphRNNActivation](g.ID, objc.Sel("outputGateActivation"))
	return MPSGraphRNNActivation(rv)
}
func (g MPSGraphGRUDescriptor) SetOutputGateActivation(value MPSGraphRNNActivation) {
	objc.Send[struct{}](g.ID, objc.Sel("setOutputGateActivation:"), value)
}

// A parameter that chooses between two variants for the reset gate
// computation.
//
// # Discussion
//
// If set to [YES] then the layer will compute the intermediate value as `c[t]
// = ( b + (h[t-1] m ) R^T) r[t]`. Otherwise it’s computed as `c[t] =
// (h[t-1] r[t] m) R^T`. Default value: [NO].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/resetAfter
func (g MPSGraphGRUDescriptor) ResetAfter() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("resetAfter"))
	return rv
}
func (g MPSGraphGRUDescriptor) SetResetAfter(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setResetAfter:"), value)
}

// A parameter that defines the activation function to use with the reset-gate
// of the GRU operation.
//
// # Discussion
//
// Default value: [MPSGraphRNNActivationSigmoid].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/resetGateActivation
func (g MPSGraphGRUDescriptor) ResetGateActivation() MPSGraphRNNActivation {
	rv := objc.Send[MPSGraphRNNActivation](g.ID, objc.Sel("resetGateActivation"))
	return MPSGraphRNNActivation(rv)
}
func (g MPSGraphGRUDescriptor) SetResetGateActivation(value MPSGraphRNNActivation) {
	objc.Send[struct{}](g.ID, objc.Sel("setResetGateActivation:"), value)
}

// A parameter that controls the internal order of the GRU gates.
//
// # Discussion
//
// If set to [YES] then the layer will use the gate-ordering `[ r, z, o ]`
// instead of default `[ z, r, o ]`. Default value: [NO].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/resetGateFirst
func (g MPSGraphGRUDescriptor) ResetGateFirst() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("resetGateFirst"))
	return rv
}
func (g MPSGraphGRUDescriptor) SetResetGateFirst(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setResetGateFirst:"), value)
}

// A parameter that defines the time direction of the input sequence.
//
// # Discussion
//
// If set to [YES] then the input sequence is passed in reverse time order to
// the layer. Note: Ignored when `bidirectional = YES`. Default value: [NO].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/reverse
func (g MPSGraphGRUDescriptor) Reverse() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("reverse"))
	return rv
}
func (g MPSGraphGRUDescriptor) SetReverse(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setReverse:"), value)
}

// A parameter that enables the GRU layer to support training.
//
// # Discussion
//
// If set to [YES] then the layer will produce training state tensor as a
// secondary output. Default value: [NO].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/training
func (g MPSGraphGRUDescriptor) Training() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("training"))
	return rv
}
func (g MPSGraphGRUDescriptor) SetTraining(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setTraining:"), value)
}

// A parameter that defines the activation function to use with the
// update-gate of the GRU operation.
//
// # Discussion
//
// Default value: [MPSGraphRNNActivationSigmoid].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/updateGateActivation
func (g MPSGraphGRUDescriptor) UpdateGateActivation() MPSGraphRNNActivation {
	rv := objc.Send[MPSGraphRNNActivation](g.ID, objc.Sel("updateGateActivation"))
	return MPSGraphRNNActivation(rv)
}
func (g MPSGraphGRUDescriptor) SetUpdateGateActivation(value MPSGraphRNNActivation) {
	objc.Send[struct{}](g.ID, objc.Sel("setUpdateGateActivation:"), value)
}
