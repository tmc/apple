// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_strength_reduction_gather_to_lookup] class.
var (
	_EspressoPass_strength_reduction_gather_to_lookupClass     EspressoPass_strength_reduction_gather_to_lookupClass
	_EspressoPass_strength_reduction_gather_to_lookupClassOnce sync.Once
)

func getEspressoPass_strength_reduction_gather_to_lookupClass() EspressoPass_strength_reduction_gather_to_lookupClass {
	_EspressoPass_strength_reduction_gather_to_lookupClassOnce.Do(func() {
		_EspressoPass_strength_reduction_gather_to_lookupClass = EspressoPass_strength_reduction_gather_to_lookupClass{class: objc.GetClass("EspressoPass_strength_reduction_gather_to_lookup")}
	})
	return _EspressoPass_strength_reduction_gather_to_lookupClass
}

// GetEspressoPass_strength_reduction_gather_to_lookupClass returns the class object for EspressoPass_strength_reduction_gather_to_lookup.
func GetEspressoPass_strength_reduction_gather_to_lookupClass() EspressoPass_strength_reduction_gather_to_lookupClass {
	return getEspressoPass_strength_reduction_gather_to_lookupClass()
}

type EspressoPass_strength_reduction_gather_to_lookupClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_strength_reduction_gather_to_lookupClass) Alloc() EspressoPass_strength_reduction_gather_to_lookup {
	rv := objc.Send[EspressoPass_strength_reduction_gather_to_lookup](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_strength_reduction_gather_to_lookup
type EspressoPass_strength_reduction_gather_to_lookup struct {
	EspressoCustomPass
}

// EspressoPass_strength_reduction_gather_to_lookupFromID constructs a [EspressoPass_strength_reduction_gather_to_lookup] from an objc.ID.
func EspressoPass_strength_reduction_gather_to_lookupFromID(id objc.ID) EspressoPass_strength_reduction_gather_to_lookup {
	return EspressoPass_strength_reduction_gather_to_lookup{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_strength_reduction_gather_to_lookup implements IEspressoPass_strength_reduction_gather_to_lookup.
var _ IEspressoPass_strength_reduction_gather_to_lookup = EspressoPass_strength_reduction_gather_to_lookup{}

// An interface definition for the [EspressoPass_strength_reduction_gather_to_lookup] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_strength_reduction_gather_to_lookup
type IEspressoPass_strength_reduction_gather_to_lookup interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_strength_reduction_gather_to_lookup) Init() EspressoPass_strength_reduction_gather_to_lookup {
	rv := objc.Send[EspressoPass_strength_reduction_gather_to_lookup](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_strength_reduction_gather_to_lookup) Autorelease() EspressoPass_strength_reduction_gather_to_lookup {
	rv := objc.Send[EspressoPass_strength_reduction_gather_to_lookup](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_strength_reduction_gather_to_lookup creates a new EspressoPass_strength_reduction_gather_to_lookup instance.
func NewEspressoPass_strength_reduction_gather_to_lookup() EspressoPass_strength_reduction_gather_to_lookup {
	class := getEspressoPass_strength_reduction_gather_to_lookupClass()
	rv := objc.Send[EspressoPass_strength_reduction_gather_to_lookup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

