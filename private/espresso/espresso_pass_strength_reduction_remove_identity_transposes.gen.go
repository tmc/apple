// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_strength_reduction_remove_identity_transposes] class.
var (
	_EspressoPass_strength_reduction_remove_identity_transposesClass     EspressoPass_strength_reduction_remove_identity_transposesClass
	_EspressoPass_strength_reduction_remove_identity_transposesClassOnce sync.Once
)

func getEspressoPass_strength_reduction_remove_identity_transposesClass() EspressoPass_strength_reduction_remove_identity_transposesClass {
	_EspressoPass_strength_reduction_remove_identity_transposesClassOnce.Do(func() {
		_EspressoPass_strength_reduction_remove_identity_transposesClass = EspressoPass_strength_reduction_remove_identity_transposesClass{class: objc.GetClass("EspressoPass_strength_reduction_remove_identity_transposes")}
	})
	return _EspressoPass_strength_reduction_remove_identity_transposesClass
}

// GetEspressoPass_strength_reduction_remove_identity_transposesClass returns the class object for EspressoPass_strength_reduction_remove_identity_transposes.
func GetEspressoPass_strength_reduction_remove_identity_transposesClass() EspressoPass_strength_reduction_remove_identity_transposesClass {
	return getEspressoPass_strength_reduction_remove_identity_transposesClass()
}

type EspressoPass_strength_reduction_remove_identity_transposesClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_strength_reduction_remove_identity_transposesClass) Alloc() EspressoPass_strength_reduction_remove_identity_transposes {
	rv := objc.Send[EspressoPass_strength_reduction_remove_identity_transposes](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/Espresso/EspressoPass_strength_reduction_remove_identity_transposes
type EspressoPass_strength_reduction_remove_identity_transposes struct {
	EspressoCustomPass
}

// EspressoPass_strength_reduction_remove_identity_transposesFromID constructs a [EspressoPass_strength_reduction_remove_identity_transposes] from an objc.ID.
func EspressoPass_strength_reduction_remove_identity_transposesFromID(id objc.ID) EspressoPass_strength_reduction_remove_identity_transposes {
	return EspressoPass_strength_reduction_remove_identity_transposes{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_strength_reduction_remove_identity_transposes implements IEspressoPass_strength_reduction_remove_identity_transposes.
var _ IEspressoPass_strength_reduction_remove_identity_transposes = EspressoPass_strength_reduction_remove_identity_transposes{}





// An interface definition for the [EspressoPass_strength_reduction_remove_identity_transposes] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_strength_reduction_remove_identity_transposes
type IEspressoPass_strength_reduction_remove_identity_transposes interface {
	IEspressoCustomPass
}





// Init initializes the instance.
func (e EspressoPass_strength_reduction_remove_identity_transposes) Init() EspressoPass_strength_reduction_remove_identity_transposes {
	rv := objc.Send[EspressoPass_strength_reduction_remove_identity_transposes](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_strength_reduction_remove_identity_transposes) Autorelease() EspressoPass_strength_reduction_remove_identity_transposes {
	rv := objc.Send[EspressoPass_strength_reduction_remove_identity_transposes](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_strength_reduction_remove_identity_transposes creates a new EspressoPass_strength_reduction_remove_identity_transposes instance.
func NewEspressoPass_strength_reduction_remove_identity_transposes() EspressoPass_strength_reduction_remove_identity_transposes {
	class := getEspressoPass_strength_reduction_remove_identity_transposesClass()
	rv := objc.Send[EspressoPass_strength_reduction_remove_identity_transposes](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































