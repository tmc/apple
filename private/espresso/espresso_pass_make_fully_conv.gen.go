// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_make_fully_conv] class.
var (
	_EspressoPass_make_fully_convClass     EspressoPass_make_fully_convClass
	_EspressoPass_make_fully_convClassOnce sync.Once
)

func getEspressoPass_make_fully_convClass() EspressoPass_make_fully_convClass {
	_EspressoPass_make_fully_convClassOnce.Do(func() {
		_EspressoPass_make_fully_convClass = EspressoPass_make_fully_convClass{class: objc.GetClass("EspressoPass_make_fully_conv")}
	})
	return _EspressoPass_make_fully_convClass
}

// GetEspressoPass_make_fully_convClass returns the class object for EspressoPass_make_fully_conv.
func GetEspressoPass_make_fully_convClass() EspressoPass_make_fully_convClass {
	return getEspressoPass_make_fully_convClass()
}

type EspressoPass_make_fully_convClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_make_fully_convClass) Alloc() EspressoPass_make_fully_conv {
	rv := objc.Send[EspressoPass_make_fully_conv](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/Espresso/EspressoPass_make_fully_conv
type EspressoPass_make_fully_conv struct {
	EspressoCustomPass
}

// EspressoPass_make_fully_convFromID constructs a [EspressoPass_make_fully_conv] from an objc.ID.
func EspressoPass_make_fully_convFromID(id objc.ID) EspressoPass_make_fully_conv {
	return EspressoPass_make_fully_conv{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_make_fully_conv implements IEspressoPass_make_fully_conv.
var _ IEspressoPass_make_fully_conv = EspressoPass_make_fully_conv{}





// An interface definition for the [EspressoPass_make_fully_conv] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_make_fully_conv
type IEspressoPass_make_fully_conv interface {
	IEspressoCustomPass
}





// Init initializes the instance.
func (e EspressoPass_make_fully_conv) Init() EspressoPass_make_fully_conv {
	rv := objc.Send[EspressoPass_make_fully_conv](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_make_fully_conv) Autorelease() EspressoPass_make_fully_conv {
	rv := objc.Send[EspressoPass_make_fully_conv](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_make_fully_conv creates a new EspressoPass_make_fully_conv instance.
func NewEspressoPass_make_fully_conv() EspressoPass_make_fully_conv {
	class := getEspressoPass_make_fully_convClass()
	rv := objc.Send[EspressoPass_make_fully_conv](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































