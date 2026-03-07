// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNPermuteGradientState] class.
var (
	_MPSNNPermuteGradientStateClass     MPSNNPermuteGradientStateClass
	_MPSNNPermuteGradientStateClassOnce sync.Once
)

func getMPSNNPermuteGradientStateClass() MPSNNPermuteGradientStateClass {
	_MPSNNPermuteGradientStateClassOnce.Do(func() {
		_MPSNNPermuteGradientStateClass = MPSNNPermuteGradientStateClass{class: objc.GetClass("MPSNNPermuteGradientState")}
	})
	return _MPSNNPermuteGradientStateClass
}

// GetMPSNNPermuteGradientStateClass returns the class object for MPSNNPermuteGradientState.
func GetMPSNNPermuteGradientStateClass() MPSNNPermuteGradientStateClass {
	return getMPSNNPermuteGradientStateClass()
}

type MPSNNPermuteGradientStateClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNPermuteGradientStateClass) Alloc() MPSNNPermuteGradientState {
	rv := objc.Send[MPSNNPermuteGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNPermuteGradientState.InitWithResource]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteGradientState
type MPSNNPermuteGradientState struct {
	MPSNNGradientState
}

// MPSNNPermuteGradientStateFromID constructs a [MPSNNPermuteGradientState] from an objc.ID.
func MPSNNPermuteGradientStateFromID(id objc.ID) MPSNNPermuteGradientState {
	return MPSNNPermuteGradientState{MPSNNGradientState: MPSNNGradientStateFromID(id)}
}
// Ensure MPSNNPermuteGradientState implements IMPSNNPermuteGradientState.
var _ IMPSNNPermuteGradientState = MPSNNPermuteGradientState{}





// An interface definition for the [MPSNNPermuteGradientState] class.
//
// # Methods
//
//   - [IMPSNNPermuteGradientState.InitWithResource]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteGradientState
type IMPSNNPermuteGradientState interface {
	IMPSNNGradientState

	// Topic: Methods

	InitWithResource(resource objectivec.IObject) MPSNNPermuteGradientState
}





// Init initializes the instance.
func (p MPSNNPermuteGradientState) Init() MPSNNPermuteGradientState {
	rv := objc.Send[MPSNNPermuteGradientState](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MPSNNPermuteGradientState) Autorelease() MPSNNPermuteGradientState {
	rv := objc.Send[MPSNNPermuteGradientState](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNPermuteGradientState creates a new MPSNNPermuteGradientState instance.
func NewMPSNNPermuteGradientState() MPSNNPermuteGradientState {
	class := getMPSNNPermuteGradientStateClass()
	rv := objc.Send[MPSNNPermuteGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteGradientState/initWithResource:
func NewPermuteGradientStateWithResource(resource objectivec.IObject) MPSNNPermuteGradientState {
	instance := getMPSNNPermuteGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSNNPermuteGradientStateFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteGradientState/initWithResource:
func (p MPSNNPermuteGradientState) InitWithResource(resource objectivec.IObject) MPSNNPermuteGradientState {
	rv := objc.Send[MPSNNPermuteGradientState](p.ID, objc.Sel("initWithResource:"), resource)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteGradientState/temporaryStateWithCommandBuffer:
func (_MPSNNPermuteGradientStateClass MPSNNPermuteGradientStateClass) TemporaryStateWithCommandBuffer(buffer objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNPermuteGradientStateClass.class), objc.Sel("temporaryStateWithCommandBuffer:"), buffer)
	return objectivec.Object{ID: rv}
}






















