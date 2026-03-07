// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNConcatenationGradientState] class.
var (
	_MPSNNConcatenationGradientStateClass     MPSNNConcatenationGradientStateClass
	_MPSNNConcatenationGradientStateClassOnce sync.Once
)

func getMPSNNConcatenationGradientStateClass() MPSNNConcatenationGradientStateClass {
	_MPSNNConcatenationGradientStateClassOnce.Do(func() {
		_MPSNNConcatenationGradientStateClass = MPSNNConcatenationGradientStateClass{class: objc.GetClass("MPSNNConcatenationGradientState")}
	})
	return _MPSNNConcatenationGradientStateClass
}

// GetMPSNNConcatenationGradientStateClass returns the class object for MPSNNConcatenationGradientState.
func GetMPSNNConcatenationGradientStateClass() MPSNNConcatenationGradientStateClass {
	return getMPSNNConcatenationGradientStateClass()
}

type MPSNNConcatenationGradientStateClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNConcatenationGradientStateClass) Alloc() MPSNNConcatenationGradientState {
	rv := objc.Send[MPSNNConcatenationGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNConcatenationGradientState.InitWithResource]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationGradientState
type MPSNNConcatenationGradientState struct {
	MPSNNGradientState
}

// MPSNNConcatenationGradientStateFromID constructs a [MPSNNConcatenationGradientState] from an objc.ID.
func MPSNNConcatenationGradientStateFromID(id objc.ID) MPSNNConcatenationGradientState {
	return MPSNNConcatenationGradientState{MPSNNGradientState: MPSNNGradientStateFromID(id)}
}
// Ensure MPSNNConcatenationGradientState implements IMPSNNConcatenationGradientState.
var _ IMPSNNConcatenationGradientState = MPSNNConcatenationGradientState{}





// An interface definition for the [MPSNNConcatenationGradientState] class.
//
// # Methods
//
//   - [IMPSNNConcatenationGradientState.InitWithResource]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationGradientState
type IMPSNNConcatenationGradientState interface {
	IMPSNNGradientState

	// Topic: Methods

	InitWithResource(resource objectivec.IObject) MPSNNConcatenationGradientState
}





// Init initializes the instance.
func (c MPSNNConcatenationGradientState) Init() MPSNNConcatenationGradientState {
	rv := objc.Send[MPSNNConcatenationGradientState](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSNNConcatenationGradientState) Autorelease() MPSNNConcatenationGradientState {
	rv := objc.Send[MPSNNConcatenationGradientState](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNConcatenationGradientState creates a new MPSNNConcatenationGradientState instance.
func NewMPSNNConcatenationGradientState() MPSNNConcatenationGradientState {
	class := getMPSNNConcatenationGradientStateClass()
	rv := objc.Send[MPSNNConcatenationGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationGradientState/initWithResource:
func NewConcatenationGradientStateWithResource(resource objectivec.IObject) MPSNNConcatenationGradientState {
	instance := getMPSNNConcatenationGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSNNConcatenationGradientStateFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationGradientState/initWithResource:
func (c MPSNNConcatenationGradientState) InitWithResource(resource objectivec.IObject) MPSNNConcatenationGradientState {
	rv := objc.Send[MPSNNConcatenationGradientState](c.ID, objc.Sel("initWithResource:"), resource)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationGradientState/temporaryStateWithCommandBuffer:
func (_MPSNNConcatenationGradientStateClass MPSNNConcatenationGradientStateClass) TemporaryStateWithCommandBuffer(buffer objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNConcatenationGradientStateClass.class), objc.Sel("temporaryStateWithCommandBuffer:"), buffer)
	return objectivec.Object{ID: rv}
}






















