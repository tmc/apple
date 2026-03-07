// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNPadGradientState] class.
var (
	_MPSNNPadGradientStateClass     MPSNNPadGradientStateClass
	_MPSNNPadGradientStateClassOnce sync.Once
)

func getMPSNNPadGradientStateClass() MPSNNPadGradientStateClass {
	_MPSNNPadGradientStateClassOnce.Do(func() {
		_MPSNNPadGradientStateClass = MPSNNPadGradientStateClass{class: objc.GetClass("MPSNNPadGradientState")}
	})
	return _MPSNNPadGradientStateClass
}

// GetMPSNNPadGradientStateClass returns the class object for MPSNNPadGradientState.
func GetMPSNNPadGradientStateClass() MPSNNPadGradientStateClass {
	return getMPSNNPadGradientStateClass()
}

type MPSNNPadGradientStateClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNPadGradientStateClass) Alloc() MPSNNPadGradientState {
	rv := objc.Send[MPSNNPadGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNPadGradientState.InitWithResource]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadGradientState
type MPSNNPadGradientState struct {
	MPSNNGradientState
}

// MPSNNPadGradientStateFromID constructs a [MPSNNPadGradientState] from an objc.ID.
func MPSNNPadGradientStateFromID(id objc.ID) MPSNNPadGradientState {
	return MPSNNPadGradientState{MPSNNGradientState: MPSNNGradientStateFromID(id)}
}
// Ensure MPSNNPadGradientState implements IMPSNNPadGradientState.
var _ IMPSNNPadGradientState = MPSNNPadGradientState{}





// An interface definition for the [MPSNNPadGradientState] class.
//
// # Methods
//
//   - [IMPSNNPadGradientState.InitWithResource]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadGradientState
type IMPSNNPadGradientState interface {
	IMPSNNGradientState

	// Topic: Methods

	InitWithResource(resource objectivec.IObject) MPSNNPadGradientState
}





// Init initializes the instance.
func (p MPSNNPadGradientState) Init() MPSNNPadGradientState {
	rv := objc.Send[MPSNNPadGradientState](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MPSNNPadGradientState) Autorelease() MPSNNPadGradientState {
	rv := objc.Send[MPSNNPadGradientState](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNPadGradientState creates a new MPSNNPadGradientState instance.
func NewMPSNNPadGradientState() MPSNNPadGradientState {
	class := getMPSNNPadGradientStateClass()
	rv := objc.Send[MPSNNPadGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadGradientState/initWithResource:
func NewPadGradientStateWithResource(resource objectivec.IObject) MPSNNPadGradientState {
	instance := getMPSNNPadGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSNNPadGradientStateFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadGradientState/initWithResource:
func (p MPSNNPadGradientState) InitWithResource(resource objectivec.IObject) MPSNNPadGradientState {
	rv := objc.Send[MPSNNPadGradientState](p.ID, objc.Sel("initWithResource:"), resource)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadGradientState/temporaryStateWithCommandBuffer:
func (_MPSNNPadGradientStateClass MPSNNPadGradientStateClass) TemporaryStateWithCommandBuffer(buffer objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNPadGradientStateClass.class), objc.Sel("temporaryStateWithCommandBuffer:"), buffer)
	return objectivec.Object{ID: rv}
}






















