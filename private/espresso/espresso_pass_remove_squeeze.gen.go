// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_remove_squeeze] class.
var (
	_EspressoPass_remove_squeezeClass     EspressoPass_remove_squeezeClass
	_EspressoPass_remove_squeezeClassOnce sync.Once
)

func getEspressoPass_remove_squeezeClass() EspressoPass_remove_squeezeClass {
	_EspressoPass_remove_squeezeClassOnce.Do(func() {
		_EspressoPass_remove_squeezeClass = EspressoPass_remove_squeezeClass{class: objc.GetClass("EspressoPass_remove_squeeze")}
	})
	return _EspressoPass_remove_squeezeClass
}

// GetEspressoPass_remove_squeezeClass returns the class object for EspressoPass_remove_squeeze.
func GetEspressoPass_remove_squeezeClass() EspressoPass_remove_squeezeClass {
	return getEspressoPass_remove_squeezeClass()
}

type EspressoPass_remove_squeezeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_remove_squeezeClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_remove_squeezeClass) Alloc() EspressoPass_remove_squeeze {
	rv := objc.Send[EspressoPass_remove_squeeze](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_remove_squeeze
type EspressoPass_remove_squeeze struct {
	EspressoCustomPass
}

// EspressoPass_remove_squeezeFromID constructs a [EspressoPass_remove_squeeze] from an objc.ID.
func EspressoPass_remove_squeezeFromID(id objc.ID) EspressoPass_remove_squeeze {
	return EspressoPass_remove_squeeze{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// Ensure EspressoPass_remove_squeeze implements IEspressoPass_remove_squeeze.
var _ IEspressoPass_remove_squeeze = EspressoPass_remove_squeeze{}

// An interface definition for the [EspressoPass_remove_squeeze] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_remove_squeeze
type IEspressoPass_remove_squeeze interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_remove_squeeze) Init() EspressoPass_remove_squeeze {
	rv := objc.Send[EspressoPass_remove_squeeze](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_remove_squeeze) Autorelease() EspressoPass_remove_squeeze {
	rv := objc.Send[EspressoPass_remove_squeeze](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_remove_squeeze creates a new EspressoPass_remove_squeeze instance.
func NewEspressoPass_remove_squeeze() EspressoPass_remove_squeeze {
	class := getEspressoPass_remove_squeezeClass()
	rv := objc.Send[EspressoPass_remove_squeeze](objc.ID(class.class), objc.Sel("new"))
	return rv
}
