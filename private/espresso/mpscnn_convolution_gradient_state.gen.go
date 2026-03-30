// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNConvolutionGradientState] class.
var (
	_MPSCNNConvolutionGradientStateClass     MPSCNNConvolutionGradientStateClass
	_MPSCNNConvolutionGradientStateClassOnce sync.Once
)

func getMPSCNNConvolutionGradientStateClass() MPSCNNConvolutionGradientStateClass {
	_MPSCNNConvolutionGradientStateClassOnce.Do(func() {
		_MPSCNNConvolutionGradientStateClass = MPSCNNConvolutionGradientStateClass{class: objc.GetClass("MPSCNNConvolutionGradientState")}
	})
	return _MPSCNNConvolutionGradientStateClass
}

// GetMPSCNNConvolutionGradientStateClass returns the class object for MPSCNNConvolutionGradientState.
func GetMPSCNNConvolutionGradientStateClass() MPSCNNConvolutionGradientStateClass {
	return getMPSCNNConvolutionGradientStateClass()
}

type MPSCNNConvolutionGradientStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNConvolutionGradientStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNConvolutionGradientStateClass) Alloc() MPSCNNConvolutionGradientState {
	rv := objc.Send[MPSCNNConvolutionGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A parent class referenced by other espresso classes. [Full Topic]
type MPSCNNConvolutionGradientState struct {
	objectivec.Object
}

// MPSCNNConvolutionGradientStateFromID constructs a [MPSCNNConvolutionGradientState] from an objc.ID.
//
// A parent class referenced by other espresso classes.
func MPSCNNConvolutionGradientStateFromID(id objc.ID) MPSCNNConvolutionGradientState {
	return MPSCNNConvolutionGradientState{objectivec.Object{ID: id}}
}

// Ensure MPSCNNConvolutionGradientState implements IMPSCNNConvolutionGradientState.
var _ IMPSCNNConvolutionGradientState = MPSCNNConvolutionGradientState{}

// An interface definition for the [MPSCNNConvolutionGradientState] class.
type IMPSCNNConvolutionGradientState interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c MPSCNNConvolutionGradientState) Init() MPSCNNConvolutionGradientState {
	rv := objc.Send[MPSCNNConvolutionGradientState](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNConvolutionGradientState) Autorelease() MPSCNNConvolutionGradientState {
	rv := objc.Send[MPSCNNConvolutionGradientState](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNConvolutionGradientState creates a new MPSCNNConvolutionGradientState instance.
func NewMPSCNNConvolutionGradientState() MPSCNNConvolutionGradientState {
	class := getMPSCNNConvolutionGradientStateClass()
	rv := objc.Send[MPSCNNConvolutionGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}
