// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_find_shared_weights] class.
var (
	_EspressoPass_find_shared_weightsClass     EspressoPass_find_shared_weightsClass
	_EspressoPass_find_shared_weightsClassOnce sync.Once
)

func getEspressoPass_find_shared_weightsClass() EspressoPass_find_shared_weightsClass {
	_EspressoPass_find_shared_weightsClassOnce.Do(func() {
		_EspressoPass_find_shared_weightsClass = EspressoPass_find_shared_weightsClass{class: objc.GetClass("EspressoPass_find_shared_weights")}
	})
	return _EspressoPass_find_shared_weightsClass
}

// GetEspressoPass_find_shared_weightsClass returns the class object for EspressoPass_find_shared_weights.
func GetEspressoPass_find_shared_weightsClass() EspressoPass_find_shared_weightsClass {
	return getEspressoPass_find_shared_weightsClass()
}

type EspressoPass_find_shared_weightsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_find_shared_weightsClass) Alloc() EspressoPass_find_shared_weights {
	rv := objc.Send[EspressoPass_find_shared_weights](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_find_shared_weights
type EspressoPass_find_shared_weights struct {
	EspressoCustomPass
}

// EspressoPass_find_shared_weightsFromID constructs a [EspressoPass_find_shared_weights] from an objc.ID.
func EspressoPass_find_shared_weightsFromID(id objc.ID) EspressoPass_find_shared_weights {
	return EspressoPass_find_shared_weights{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_find_shared_weights implements IEspressoPass_find_shared_weights.
var _ IEspressoPass_find_shared_weights = EspressoPass_find_shared_weights{}

// An interface definition for the [EspressoPass_find_shared_weights] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_find_shared_weights
type IEspressoPass_find_shared_weights interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_find_shared_weights) Init() EspressoPass_find_shared_weights {
	rv := objc.Send[EspressoPass_find_shared_weights](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_find_shared_weights) Autorelease() EspressoPass_find_shared_weights {
	rv := objc.Send[EspressoPass_find_shared_weights](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_find_shared_weights creates a new EspressoPass_find_shared_weights instance.
func NewEspressoPass_find_shared_weights() EspressoPass_find_shared_weights {
	class := getEspressoPass_find_shared_weightsClass()
	rv := objc.Send[EspressoPass_find_shared_weights](objc.ID(class.class), objc.Sel("new"))
	return rv
}

