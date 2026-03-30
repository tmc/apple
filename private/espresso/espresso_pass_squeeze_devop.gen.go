// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_squeeze_devop] class.
var (
	_EspressoPass_squeeze_devopClass     EspressoPass_squeeze_devopClass
	_EspressoPass_squeeze_devopClassOnce sync.Once
)

func getEspressoPass_squeeze_devopClass() EspressoPass_squeeze_devopClass {
	_EspressoPass_squeeze_devopClassOnce.Do(func() {
		_EspressoPass_squeeze_devopClass = EspressoPass_squeeze_devopClass{class: objc.GetClass("EspressoPass_squeeze_devop")}
	})
	return _EspressoPass_squeeze_devopClass
}

// GetEspressoPass_squeeze_devopClass returns the class object for EspressoPass_squeeze_devop.
func GetEspressoPass_squeeze_devopClass() EspressoPass_squeeze_devopClass {
	return getEspressoPass_squeeze_devopClass()
}

type EspressoPass_squeeze_devopClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_squeeze_devopClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_squeeze_devopClass) Alloc() EspressoPass_squeeze_devop {
	rv := objc.Send[EspressoPass_squeeze_devop](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_squeeze_devop
type EspressoPass_squeeze_devop struct {
	EspressoCustomPass
}

// EspressoPass_squeeze_devopFromID constructs a [EspressoPass_squeeze_devop] from an objc.ID.
func EspressoPass_squeeze_devopFromID(id objc.ID) EspressoPass_squeeze_devop {
	return EspressoPass_squeeze_devop{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// Ensure EspressoPass_squeeze_devop implements IEspressoPass_squeeze_devop.
var _ IEspressoPass_squeeze_devop = EspressoPass_squeeze_devop{}

// An interface definition for the [EspressoPass_squeeze_devop] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_squeeze_devop
type IEspressoPass_squeeze_devop interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_squeeze_devop) Init() EspressoPass_squeeze_devop {
	rv := objc.Send[EspressoPass_squeeze_devop](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_squeeze_devop) Autorelease() EspressoPass_squeeze_devop {
	rv := objc.Send[EspressoPass_squeeze_devop](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_squeeze_devop creates a new EspressoPass_squeeze_devop instance.
func NewEspressoPass_squeeze_devop() EspressoPass_squeeze_devop {
	class := getEspressoPass_squeeze_devopClass()
	rv := objc.Send[EspressoPass_squeeze_devop](objc.ID(class.class), objc.Sel("new"))
	return rv
}
