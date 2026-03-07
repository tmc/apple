// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNBinaryGradientState] class.
var (
	_MPSNNBinaryGradientStateClass     MPSNNBinaryGradientStateClass
	_MPSNNBinaryGradientStateClassOnce sync.Once
)

func getMPSNNBinaryGradientStateClass() MPSNNBinaryGradientStateClass {
	_MPSNNBinaryGradientStateClassOnce.Do(func() {
		_MPSNNBinaryGradientStateClass = MPSNNBinaryGradientStateClass{class: objc.GetClass("MPSNNBinaryGradientState")}
	})
	return _MPSNNBinaryGradientStateClass
}

// GetMPSNNBinaryGradientStateClass returns the class object for MPSNNBinaryGradientState.
func GetMPSNNBinaryGradientStateClass() MPSNNBinaryGradientStateClass {
	return getMPSNNBinaryGradientStateClass()
}

type MPSNNBinaryGradientStateClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNBinaryGradientStateClass) Alloc() MPSNNBinaryGradientState {
	rv := objc.Send[MPSNNBinaryGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNBinaryGradientState.DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryGradientState
type MPSNNBinaryGradientState struct {
	MPSState
}

// MPSNNBinaryGradientStateFromID constructs a [MPSNNBinaryGradientState] from an objc.ID.
func MPSNNBinaryGradientStateFromID(id objc.ID) MPSNNBinaryGradientState {
	return MPSNNBinaryGradientState{MPSState: MPSStateFromID(id)}
}
// Ensure MPSNNBinaryGradientState implements IMPSNNBinaryGradientState.
var _ IMPSNNBinaryGradientState = MPSNNBinaryGradientState{}





// An interface definition for the [MPSNNBinaryGradientState] class.
//
// # Methods
//
//   - [IMPSNNBinaryGradientState.DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryGradientState
type IMPSNNBinaryGradientState interface {
	IMPSState

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(images objectivec.IObject, states objectivec.IObject, kernel objectivec.IObject, descriptor objectivec.IObject) objectivec.IObject
}





// Init initializes the instance.
func (b MPSNNBinaryGradientState) Init() MPSNNBinaryGradientState {
	rv := objc.Send[MPSNNBinaryGradientState](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MPSNNBinaryGradientState) Autorelease() MPSNNBinaryGradientState {
	rv := objc.Send[MPSNNBinaryGradientState](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNBinaryGradientState creates a new MPSNNBinaryGradientState instance.
func NewMPSNNBinaryGradientState() MPSNNBinaryGradientState {
	class := getMPSNNBinaryGradientStateClass()
	rv := objc.Send[MPSNNBinaryGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}










//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryGradientState/destinationImageDescriptorForSourceImages:sourceStates:forKernel:suggestedDescriptor:
func (b MPSNNBinaryGradientState) DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(images objectivec.IObject, states objectivec.IObject, kernel objectivec.IObject, descriptor objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:forKernel:suggestedDescriptor:"), images, states, kernel, descriptor)
	return objectivec.Object{ID: rv}
}


























