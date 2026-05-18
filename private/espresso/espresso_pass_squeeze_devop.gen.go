// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassSqueezeDevop] class.
var (
	_EspressoPassSqueezeDevopClass     EspressoPassSqueezeDevopClass
	_EspressoPassSqueezeDevopClassOnce sync.Once
)

func getEspressoPassSqueezeDevopClass() EspressoPassSqueezeDevopClass {
	_EspressoPassSqueezeDevopClassOnce.Do(func() {
		_EspressoPassSqueezeDevopClass = EspressoPassSqueezeDevopClass{class: objc.GetClass("EspressoPass_squeeze_devop")}
	})
	return _EspressoPassSqueezeDevopClass
}

// GetEspressoPassSqueezeDevopClass returns the class object for EspressoPass_squeeze_devop.
func GetEspressoPassSqueezeDevopClass() EspressoPassSqueezeDevopClass {
	return getEspressoPassSqueezeDevopClass()
}

type EspressoPassSqueezeDevopClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassSqueezeDevopClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassSqueezeDevopClass) Alloc() EspressoPassSqueezeDevop {
	rv := objc.Send[EspressoPassSqueezeDevop](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_squeeze_devop
type EspressoPassSqueezeDevop struct {
	EspressoCustomPass
}

// EspressoPassSqueezeDevopFromID constructs a [EspressoPassSqueezeDevop] from an objc.ID.
func EspressoPassSqueezeDevopFromID(id objc.ID) EspressoPassSqueezeDevop {
	return EspressoPassSqueezeDevop{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_squeeze_devopFromID is an alias for [EspressoPassSqueezeDevopFromID] for cross-framework compatibility.
func EspressoPass_squeeze_devopFromID(id objc.ID) EspressoPassSqueezeDevop {
	return EspressoPassSqueezeDevopFromID(id)
}

// Ensure EspressoPassSqueezeDevop implements IEspressoPassSqueezeDevop.
var _ IEspressoPassSqueezeDevop = EspressoPassSqueezeDevop{}

// An interface definition for the [EspressoPassSqueezeDevop] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_squeeze_devop
type IEspressoPassSqueezeDevop interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassSqueezeDevop) Init() EspressoPassSqueezeDevop {
	rv := objc.Send[EspressoPassSqueezeDevop](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassSqueezeDevop) Autorelease() EspressoPassSqueezeDevop {
	rv := objc.Send[EspressoPassSqueezeDevop](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassSqueezeDevop creates a new EspressoPassSqueezeDevop instance.
func NewEspressoPassSqueezeDevop() EspressoPassSqueezeDevop {
	class := getEspressoPassSqueezeDevopClass()
	rv := objc.Send[EspressoPassSqueezeDevop](objc.ID(class.class), objc.Sel("new"))
	return rv
}
