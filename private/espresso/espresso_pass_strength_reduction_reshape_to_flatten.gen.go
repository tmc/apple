// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_strength_reduction_reshape_to_flatten] class.
var (
	_EspressoPass_strength_reduction_reshape_to_flattenClass     EspressoPass_strength_reduction_reshape_to_flattenClass
	_EspressoPass_strength_reduction_reshape_to_flattenClassOnce sync.Once
)

func getEspressoPass_strength_reduction_reshape_to_flattenClass() EspressoPass_strength_reduction_reshape_to_flattenClass {
	_EspressoPass_strength_reduction_reshape_to_flattenClassOnce.Do(func() {
		_EspressoPass_strength_reduction_reshape_to_flattenClass = EspressoPass_strength_reduction_reshape_to_flattenClass{class: objc.GetClass("EspressoPass_strength_reduction_reshape_to_flatten")}
	})
	return _EspressoPass_strength_reduction_reshape_to_flattenClass
}

// GetEspressoPass_strength_reduction_reshape_to_flattenClass returns the class object for EspressoPass_strength_reduction_reshape_to_flatten.
func GetEspressoPass_strength_reduction_reshape_to_flattenClass() EspressoPass_strength_reduction_reshape_to_flattenClass {
	return getEspressoPass_strength_reduction_reshape_to_flattenClass()
}

type EspressoPass_strength_reduction_reshape_to_flattenClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_strength_reduction_reshape_to_flattenClass) Alloc() EspressoPass_strength_reduction_reshape_to_flatten {
	rv := objc.Send[EspressoPass_strength_reduction_reshape_to_flatten](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_strength_reduction_reshape_to_flatten
type EspressoPass_strength_reduction_reshape_to_flatten struct {
	EspressoCustomPass
}

// EspressoPass_strength_reduction_reshape_to_flattenFromID constructs a [EspressoPass_strength_reduction_reshape_to_flatten] from an objc.ID.
func EspressoPass_strength_reduction_reshape_to_flattenFromID(id objc.ID) EspressoPass_strength_reduction_reshape_to_flatten {
	return EspressoPass_strength_reduction_reshape_to_flatten{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_strength_reduction_reshape_to_flatten implements IEspressoPass_strength_reduction_reshape_to_flatten.
var _ IEspressoPass_strength_reduction_reshape_to_flatten = EspressoPass_strength_reduction_reshape_to_flatten{}

// An interface definition for the [EspressoPass_strength_reduction_reshape_to_flatten] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_strength_reduction_reshape_to_flatten
type IEspressoPass_strength_reduction_reshape_to_flatten interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_strength_reduction_reshape_to_flatten) Init() EspressoPass_strength_reduction_reshape_to_flatten {
	rv := objc.Send[EspressoPass_strength_reduction_reshape_to_flatten](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_strength_reduction_reshape_to_flatten) Autorelease() EspressoPass_strength_reduction_reshape_to_flatten {
	rv := objc.Send[EspressoPass_strength_reduction_reshape_to_flatten](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_strength_reduction_reshape_to_flatten creates a new EspressoPass_strength_reduction_reshape_to_flatten instance.
func NewEspressoPass_strength_reduction_reshape_to_flatten() EspressoPass_strength_reduction_reshape_to_flatten {
	class := getEspressoPass_strength_reduction_reshape_to_flattenClass()
	rv := objc.Send[EspressoPass_strength_reduction_reshape_to_flatten](objc.ID(class.class), objc.Sel("new"))
	return rv
}

