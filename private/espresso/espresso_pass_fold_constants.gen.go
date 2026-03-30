// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_fold_constants] class.
var (
	_EspressoPass_fold_constantsClass     EspressoPass_fold_constantsClass
	_EspressoPass_fold_constantsClassOnce sync.Once
)

func getEspressoPass_fold_constantsClass() EspressoPass_fold_constantsClass {
	_EspressoPass_fold_constantsClassOnce.Do(func() {
		_EspressoPass_fold_constantsClass = EspressoPass_fold_constantsClass{class: objc.GetClass("EspressoPass_fold_constants")}
	})
	return _EspressoPass_fold_constantsClass
}

// GetEspressoPass_fold_constantsClass returns the class object for EspressoPass_fold_constants.
func GetEspressoPass_fold_constantsClass() EspressoPass_fold_constantsClass {
	return getEspressoPass_fold_constantsClass()
}

type EspressoPass_fold_constantsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_fold_constantsClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_fold_constantsClass) Alloc() EspressoPass_fold_constants {
	rv := objc.Send[EspressoPass_fold_constants](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fold_constants
type EspressoPass_fold_constants struct {
	EspressoCustomPass
}

// EspressoPass_fold_constantsFromID constructs a [EspressoPass_fold_constants] from an objc.ID.
func EspressoPass_fold_constantsFromID(id objc.ID) EspressoPass_fold_constants {
	return EspressoPass_fold_constants{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// Ensure EspressoPass_fold_constants implements IEspressoPass_fold_constants.
var _ IEspressoPass_fold_constants = EspressoPass_fold_constants{}

// An interface definition for the [EspressoPass_fold_constants] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fold_constants
type IEspressoPass_fold_constants interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_fold_constants) Init() EspressoPass_fold_constants {
	rv := objc.Send[EspressoPass_fold_constants](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_fold_constants) Autorelease() EspressoPass_fold_constants {
	rv := objc.Send[EspressoPass_fold_constants](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_fold_constants creates a new EspressoPass_fold_constants instance.
func NewEspressoPass_fold_constants() EspressoPass_fold_constants {
	class := getEspressoPass_fold_constantsClass()
	rv := objc.Send[EspressoPass_fold_constants](objc.ID(class.class), objc.Sel("new"))
	return rv
}
