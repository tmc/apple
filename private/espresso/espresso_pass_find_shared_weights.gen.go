// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassFindSharedWeights] class.
var (
	_EspressoPassFindSharedWeightsClass     EspressoPassFindSharedWeightsClass
	_EspressoPassFindSharedWeightsClassOnce sync.Once
)

func getEspressoPassFindSharedWeightsClass() EspressoPassFindSharedWeightsClass {
	_EspressoPassFindSharedWeightsClassOnce.Do(func() {
		_EspressoPassFindSharedWeightsClass = EspressoPassFindSharedWeightsClass{class: objc.GetClass("EspressoPass_find_shared_weights")}
	})
	return _EspressoPassFindSharedWeightsClass
}

// GetEspressoPassFindSharedWeightsClass returns the class object for EspressoPass_find_shared_weights.
func GetEspressoPassFindSharedWeightsClass() EspressoPassFindSharedWeightsClass {
	return getEspressoPassFindSharedWeightsClass()
}

type EspressoPassFindSharedWeightsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassFindSharedWeightsClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassFindSharedWeightsClass) Alloc() EspressoPassFindSharedWeights {
	rv := objc.Send[EspressoPassFindSharedWeights](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassFindSharedWeights struct {
	EspressoCustomPass
}

// EspressoPassFindSharedWeightsFromID constructs a [EspressoPassFindSharedWeights] from an objc.ID.
func EspressoPassFindSharedWeightsFromID(id objc.ID) EspressoPassFindSharedWeights {
	return EspressoPassFindSharedWeights{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_find_shared_weightsFromID is an alias for [EspressoPassFindSharedWeightsFromID] for cross-framework compatibility.
func EspressoPass_find_shared_weightsFromID(id objc.ID) EspressoPassFindSharedWeights {
	return EspressoPassFindSharedWeightsFromID(id)
}

// Ensure EspressoPassFindSharedWeights implements IEspressoPassFindSharedWeights.
var _ IEspressoPassFindSharedWeights = EspressoPassFindSharedWeights{}

// An interface definition for the [EspressoPassFindSharedWeights] class.
type IEspressoPassFindSharedWeights interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassFindSharedWeights) Init() EspressoPassFindSharedWeights {
	rv := objc.Send[EspressoPassFindSharedWeights](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassFindSharedWeights) Autorelease() EspressoPassFindSharedWeights {
	rv := objc.Send[EspressoPassFindSharedWeights](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassFindSharedWeights creates a new EspressoPassFindSharedWeights instance.
func NewEspressoPassFindSharedWeights() EspressoPassFindSharedWeights {
	class := getEspressoPassFindSharedWeightsClass()
	rv := objc.Send[EspressoPassFindSharedWeights](objc.ID(class.class), objc.Sel("new"))
	return rv
}
