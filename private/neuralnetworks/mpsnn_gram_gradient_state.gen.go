// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNGramGradientState] class.
var (
	_MPSNNGramGradientStateClass     MPSNNGramGradientStateClass
	_MPSNNGramGradientStateClassOnce sync.Once
)

func getMPSNNGramGradientStateClass() MPSNNGramGradientStateClass {
	_MPSNNGramGradientStateClassOnce.Do(func() {
		_MPSNNGramGradientStateClass = MPSNNGramGradientStateClass{class: objc.GetClass("MPSNNGramGradientState")}
	})
	return _MPSNNGramGradientStateClass
}

// GetMPSNNGramGradientStateClass returns the class object for MPSNNGramGradientState.
func GetMPSNNGramGradientStateClass() MPSNNGramGradientStateClass {
	return getMPSNNGramGradientStateClass()
}

type MPSNNGramGradientStateClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNGramGradientStateClass) Alloc() MPSNNGramGradientState {
	rv := objc.Send[MPSNNGramGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNGramGradientState.InitWithResource]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramGradientState
type MPSNNGramGradientState struct {
	MPSNNGradientState
}

// MPSNNGramGradientStateFromID constructs a [MPSNNGramGradientState] from an objc.ID.
func MPSNNGramGradientStateFromID(id objc.ID) MPSNNGramGradientState {
	return MPSNNGramGradientState{MPSNNGradientState: MPSNNGradientStateFromID(id)}
}
// Ensure MPSNNGramGradientState implements IMPSNNGramGradientState.
var _ IMPSNNGramGradientState = MPSNNGramGradientState{}





// An interface definition for the [MPSNNGramGradientState] class.
//
// # Methods
//
//   - [IMPSNNGramGradientState.InitWithResource]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramGradientState
type IMPSNNGramGradientState interface {
	IMPSNNGradientState

	// Topic: Methods

	InitWithResource(resource objectivec.IObject) MPSNNGramGradientState
}





// Init initializes the instance.
func (g MPSNNGramGradientState) Init() MPSNNGramGradientState {
	rv := objc.Send[MPSNNGramGradientState](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSNNGramGradientState) Autorelease() MPSNNGramGradientState {
	rv := objc.Send[MPSNNGramGradientState](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNGramGradientState creates a new MPSNNGramGradientState instance.
func NewMPSNNGramGradientState() MPSNNGramGradientState {
	class := getMPSNNGramGradientStateClass()
	rv := objc.Send[MPSNNGramGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramGradientState/initWithResource:
func NewGramGradientStateWithResource(resource objectivec.IObject) MPSNNGramGradientState {
	instance := getMPSNNGramGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSNNGramGradientStateFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramGradientState/initWithResource:
func (g MPSNNGramGradientState) InitWithResource(resource objectivec.IObject) MPSNNGramGradientState {
	rv := objc.Send[MPSNNGramGradientState](g.ID, objc.Sel("initWithResource:"), resource)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramGradientState/temporaryStateWithCommandBuffer:
func (_MPSNNGramGradientStateClass MPSNNGramGradientStateClass) TemporaryStateWithCommandBuffer(buffer objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNGramGradientStateClass.class), objc.Sel("temporaryStateWithCommandBuffer:"), buffer)
	return objectivec.Object{ID: rv}
}






















