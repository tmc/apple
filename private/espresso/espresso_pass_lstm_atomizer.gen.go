// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_lstm_atomizer] class.
var (
	_EspressoPass_lstm_atomizerClass     EspressoPass_lstm_atomizerClass
	_EspressoPass_lstm_atomizerClassOnce sync.Once
)

func getEspressoPass_lstm_atomizerClass() EspressoPass_lstm_atomizerClass {
	_EspressoPass_lstm_atomizerClassOnce.Do(func() {
		_EspressoPass_lstm_atomizerClass = EspressoPass_lstm_atomizerClass{class: objc.GetClass("EspressoPass_lstm_atomizer")}
	})
	return _EspressoPass_lstm_atomizerClass
}

// GetEspressoPass_lstm_atomizerClass returns the class object for EspressoPass_lstm_atomizer.
func GetEspressoPass_lstm_atomizerClass() EspressoPass_lstm_atomizerClass {
	return getEspressoPass_lstm_atomizerClass()
}

type EspressoPass_lstm_atomizerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_lstm_atomizerClass) Alloc() EspressoPass_lstm_atomizer {
	rv := objc.Send[EspressoPass_lstm_atomizer](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/Espresso/EspressoPass_lstm_atomizer
type EspressoPass_lstm_atomizer struct {
	EspressoCustomPass
}

// EspressoPass_lstm_atomizerFromID constructs a [EspressoPass_lstm_atomizer] from an objc.ID.
func EspressoPass_lstm_atomizerFromID(id objc.ID) EspressoPass_lstm_atomizer {
	return EspressoPass_lstm_atomizer{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_lstm_atomizer implements IEspressoPass_lstm_atomizer.
var _ IEspressoPass_lstm_atomizer = EspressoPass_lstm_atomizer{}





// An interface definition for the [EspressoPass_lstm_atomizer] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_lstm_atomizer
type IEspressoPass_lstm_atomizer interface {
	IEspressoCustomPass
}





// Init initializes the instance.
func (e EspressoPass_lstm_atomizer) Init() EspressoPass_lstm_atomizer {
	rv := objc.Send[EspressoPass_lstm_atomizer](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_lstm_atomizer) Autorelease() EspressoPass_lstm_atomizer {
	rv := objc.Send[EspressoPass_lstm_atomizer](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_lstm_atomizer creates a new EspressoPass_lstm_atomizer instance.
func NewEspressoPass_lstm_atomizer() EspressoPass_lstm_atomizer {
	class := getEspressoPass_lstm_atomizerClass()
	rv := objc.Send[EspressoPass_lstm_atomizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































