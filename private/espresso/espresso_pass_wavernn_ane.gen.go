// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_wavernn_ane] class.
var (
	_EspressoPass_wavernn_aneClass     EspressoPass_wavernn_aneClass
	_EspressoPass_wavernn_aneClassOnce sync.Once
)

func getEspressoPass_wavernn_aneClass() EspressoPass_wavernn_aneClass {
	_EspressoPass_wavernn_aneClassOnce.Do(func() {
		_EspressoPass_wavernn_aneClass = EspressoPass_wavernn_aneClass{class: objc.GetClass("EspressoPass_wavernn_ane")}
	})
	return _EspressoPass_wavernn_aneClass
}

// GetEspressoPass_wavernn_aneClass returns the class object for EspressoPass_wavernn_ane.
func GetEspressoPass_wavernn_aneClass() EspressoPass_wavernn_aneClass {
	return getEspressoPass_wavernn_aneClass()
}

type EspressoPass_wavernn_aneClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_wavernn_aneClass) Alloc() EspressoPass_wavernn_ane {
	rv := objc.Send[EspressoPass_wavernn_ane](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/Espresso/EspressoPass_wavernn_ane
type EspressoPass_wavernn_ane struct {
	EspressoCustomPass
}

// EspressoPass_wavernn_aneFromID constructs a [EspressoPass_wavernn_ane] from an objc.ID.
func EspressoPass_wavernn_aneFromID(id objc.ID) EspressoPass_wavernn_ane {
	return EspressoPass_wavernn_ane{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_wavernn_ane implements IEspressoPass_wavernn_ane.
var _ IEspressoPass_wavernn_ane = EspressoPass_wavernn_ane{}





// An interface definition for the [EspressoPass_wavernn_ane] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_wavernn_ane
type IEspressoPass_wavernn_ane interface {
	IEspressoCustomPass
}





// Init initializes the instance.
func (e EspressoPass_wavernn_ane) Init() EspressoPass_wavernn_ane {
	rv := objc.Send[EspressoPass_wavernn_ane](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_wavernn_ane) Autorelease() EspressoPass_wavernn_ane {
	rv := objc.Send[EspressoPass_wavernn_ane](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_wavernn_ane creates a new EspressoPass_wavernn_ane instance.
func NewEspressoPass_wavernn_ane() EspressoPass_wavernn_ane {
	class := getEspressoPass_wavernn_aneClass()
	rv := objc.Send[EspressoPass_wavernn_ane](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































