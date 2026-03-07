// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNLossGradientState] class.
var (
	_MPSNNLossGradientStateClass     MPSNNLossGradientStateClass
	_MPSNNLossGradientStateClassOnce sync.Once
)

func getMPSNNLossGradientStateClass() MPSNNLossGradientStateClass {
	_MPSNNLossGradientStateClassOnce.Do(func() {
		_MPSNNLossGradientStateClass = MPSNNLossGradientStateClass{class: objc.GetClass("MPSNNLossGradientState")}
	})
	return _MPSNNLossGradientStateClass
}

// GetMPSNNLossGradientStateClass returns the class object for MPSNNLossGradientState.
func GetMPSNNLossGradientStateClass() MPSNNLossGradientStateClass {
	return getMPSNNLossGradientStateClass()
}

type MPSNNLossGradientStateClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNLossGradientStateClass) Alloc() MPSNNLossGradientState {
	rv := objc.Send[MPSNNLossGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNLossGradientState.InitWithResource]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientState
type MPSNNLossGradientState struct {
	MPSNNGradientState
}

// MPSNNLossGradientStateFromID constructs a [MPSNNLossGradientState] from an objc.ID.
func MPSNNLossGradientStateFromID(id objc.ID) MPSNNLossGradientState {
	return MPSNNLossGradientState{MPSNNGradientState: MPSNNGradientStateFromID(id)}
}
// Ensure MPSNNLossGradientState implements IMPSNNLossGradientState.
var _ IMPSNNLossGradientState = MPSNNLossGradientState{}





// An interface definition for the [MPSNNLossGradientState] class.
//
// # Methods
//
//   - [IMPSNNLossGradientState.InitWithResource]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientState
type IMPSNNLossGradientState interface {
	IMPSNNGradientState

	// Topic: Methods

	InitWithResource(resource objectivec.IObject) MPSNNLossGradientState
}





// Init initializes the instance.
func (l MPSNNLossGradientState) Init() MPSNNLossGradientState {
	rv := objc.Send[MPSNNLossGradientState](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MPSNNLossGradientState) Autorelease() MPSNNLossGradientState {
	rv := objc.Send[MPSNNLossGradientState](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNLossGradientState creates a new MPSNNLossGradientState instance.
func NewMPSNNLossGradientState() MPSNNLossGradientState {
	class := getMPSNNLossGradientStateClass()
	rv := objc.Send[MPSNNLossGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientState/initWithResource:
func NewLossGradientStateWithResource(resource objectivec.IObject) MPSNNLossGradientState {
	instance := getMPSNNLossGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSNNLossGradientStateFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientState/initWithResource:
func (l MPSNNLossGradientState) InitWithResource(resource objectivec.IObject) MPSNNLossGradientState {
	rv := objc.Send[MPSNNLossGradientState](l.ID, objc.Sel("initWithResource:"), resource)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientState/temporaryStateWithCommandBuffer:
func (_MPSNNLossGradientStateClass MPSNNLossGradientStateClass) TemporaryStateWithCommandBuffer(buffer objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNLossGradientStateClass.class), objc.Sel("temporaryStateWithCommandBuffer:"), buffer)
	return objectivec.Object{ID: rv}
}






















