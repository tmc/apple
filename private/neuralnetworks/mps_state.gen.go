// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSState] class.
var (
	_MPSStateClass     MPSStateClass
	_MPSStateClassOnce sync.Once
)

func getMPSStateClass() MPSStateClass {
	_MPSStateClassOnce.Do(func() {
		_MPSStateClass = MPSStateClass{class: objc.GetClass("MPSState")}
	})
	return _MPSStateClass
}

// GetMPSStateClass returns the class object for MPSState.
func GetMPSStateClass() MPSStateClass {
	return getMPSStateClass()
}

type MPSStateClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSStateClass) Alloc() MPSState {
	rv := objc.Send[MPSState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A parent class referenced by other neuralnetworks classes. [Full Topic]
type MPSState struct {
	objectivec.Object
}

// MPSStateFromID constructs a [MPSState] from an objc.ID.
//
// A parent class referenced by other neuralnetworks classes.
func MPSStateFromID(id objc.ID) MPSState {
	return MPSState{objectivec.Object{id}}
}
// Ensure MPSState implements IMPSState.
var _ IMPSState = MPSState{}





// An interface definition for the [MPSState] class.
type IMPSState interface {
	objectivec.IObject
}





// Init initializes the instance.
func (s MPSState) Init() MPSState {
	rv := objc.Send[MPSState](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MPSState) Autorelease() MPSState {
	rv := objc.Send[MPSState](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSState creates a new MPSState instance.
func NewMPSState() MPSState {
	class := getMPSStateClass()
	rv := objc.Send[MPSState](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































