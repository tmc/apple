// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNGradientState] class.
var (
	_MPSNNGradientStateClass     MPSNNGradientStateClass
	_MPSNNGradientStateClassOnce sync.Once
)

func getMPSNNGradientStateClass() MPSNNGradientStateClass {
	_MPSNNGradientStateClassOnce.Do(func() {
		_MPSNNGradientStateClass = MPSNNGradientStateClass{class: objc.GetClass("MPSNNGradientState")}
	})
	return _MPSNNGradientStateClass
}

// GetMPSNNGradientStateClass returns the class object for MPSNNGradientState.
func GetMPSNNGradientStateClass() MPSNNGradientStateClass {
	return getMPSNNGradientStateClass()
}

type MPSNNGradientStateClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNGradientStateClass) Alloc() MPSNNGradientState {
	rv := objc.Send[MPSNNGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNGradientState.DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientState
type MPSNNGradientState struct {
	MPSState
}

// MPSNNGradientStateFromID constructs a [MPSNNGradientState] from an objc.ID.
func MPSNNGradientStateFromID(id objc.ID) MPSNNGradientState {
	return MPSNNGradientState{MPSState: MPSStateFromID(id)}
}
// Ensure MPSNNGradientState implements IMPSNNGradientState.
var _ IMPSNNGradientState = MPSNNGradientState{}





// An interface definition for the [MPSNNGradientState] class.
//
// # Methods
//
//   - [IMPSNNGradientState.DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientState
type IMPSNNGradientState interface {
	IMPSState

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(images objectivec.IObject, states objectivec.IObject, kernel objectivec.IObject, descriptor objectivec.IObject) objectivec.IObject
}





// Init initializes the instance.
func (g MPSNNGradientState) Init() MPSNNGradientState {
	rv := objc.Send[MPSNNGradientState](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSNNGradientState) Autorelease() MPSNNGradientState {
	rv := objc.Send[MPSNNGradientState](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNGradientState creates a new MPSNNGradientState instance.
func NewMPSNNGradientState() MPSNNGradientState {
	class := getMPSNNGradientStateClass()
	rv := objc.Send[MPSNNGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}










//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientState/destinationImageDescriptorForSourceImages:sourceStates:forKernel:suggestedDescriptor:
func (g MPSNNGradientState) DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(images objectivec.IObject, states objectivec.IObject, kernel objectivec.IObject, descriptor objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:forKernel:suggestedDescriptor:"), images, states, kernel, descriptor)
	return objectivec.Object{ID: rv}
}


























