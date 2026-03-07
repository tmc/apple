// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNMultiaryGradientState] class.
var (
	_MPSNNMultiaryGradientStateClass     MPSNNMultiaryGradientStateClass
	_MPSNNMultiaryGradientStateClassOnce sync.Once
)

func getMPSNNMultiaryGradientStateClass() MPSNNMultiaryGradientStateClass {
	_MPSNNMultiaryGradientStateClassOnce.Do(func() {
		_MPSNNMultiaryGradientStateClass = MPSNNMultiaryGradientStateClass{class: objc.GetClass("MPSNNMultiaryGradientState")}
	})
	return _MPSNNMultiaryGradientStateClass
}

// GetMPSNNMultiaryGradientStateClass returns the class object for MPSNNMultiaryGradientState.
func GetMPSNNMultiaryGradientStateClass() MPSNNMultiaryGradientStateClass {
	return getMPSNNMultiaryGradientStateClass()
}

type MPSNNMultiaryGradientStateClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNMultiaryGradientStateClass) Alloc() MPSNNMultiaryGradientState {
	rv := objc.Send[MPSNNMultiaryGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNMultiaryGradientState.InitWithSourceCount]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNMultiaryGradientState
type MPSNNMultiaryGradientState struct {
	MPSState
}

// MPSNNMultiaryGradientStateFromID constructs a [MPSNNMultiaryGradientState] from an objc.ID.
func MPSNNMultiaryGradientStateFromID(id objc.ID) MPSNNMultiaryGradientState {
	return MPSNNMultiaryGradientState{MPSState: MPSStateFromID(id)}
}
// Ensure MPSNNMultiaryGradientState implements IMPSNNMultiaryGradientState.
var _ IMPSNNMultiaryGradientState = MPSNNMultiaryGradientState{}





// An interface definition for the [MPSNNMultiaryGradientState] class.
//
// # Methods
//
//   - [IMPSNNMultiaryGradientState.InitWithSourceCount]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNMultiaryGradientState
type IMPSNNMultiaryGradientState interface {
	IMPSState

	// Topic: Methods

	InitWithSourceCount(count uint64) MPSNNMultiaryGradientState
}





// Init initializes the instance.
func (m MPSNNMultiaryGradientState) Init() MPSNNMultiaryGradientState {
	rv := objc.Send[MPSNNMultiaryGradientState](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSNNMultiaryGradientState) Autorelease() MPSNNMultiaryGradientState {
	rv := objc.Send[MPSNNMultiaryGradientState](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNMultiaryGradientState creates a new MPSNNMultiaryGradientState instance.
func NewMPSNNMultiaryGradientState() MPSNNMultiaryGradientState {
	class := getMPSNNMultiaryGradientStateClass()
	rv := objc.Send[MPSNNMultiaryGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNMultiaryGradientState/initWithSourceCount:
func NewMultiaryGradientStateWithSourceCount(count uint64) MPSNNMultiaryGradientState {
	instance := getMPSNNMultiaryGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceCount:"), count)
	return MPSNNMultiaryGradientStateFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNMultiaryGradientState/initWithSourceCount:
func (m MPSNNMultiaryGradientState) InitWithSourceCount(count uint64) MPSNNMultiaryGradientState {
	rv := objc.Send[MPSNNMultiaryGradientState](m.ID, objc.Sel("initWithSourceCount:"), count)
	return rv
}


























