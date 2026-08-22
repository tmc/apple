// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassLstmAtomizer] class.
var (
	_EspressoPassLstmAtomizerClass     EspressoPassLstmAtomizerClass
	_EspressoPassLstmAtomizerClassOnce sync.Once
)

func getEspressoPassLstmAtomizerClass() EspressoPassLstmAtomizerClass {
	_EspressoPassLstmAtomizerClassOnce.Do(func() {
		_EspressoPassLstmAtomizerClass = EspressoPassLstmAtomizerClass{class: objc.GetClass("EspressoPass_lstm_atomizer")}
	})
	return _EspressoPassLstmAtomizerClass
}

// GetEspressoPassLstmAtomizerClass returns the class object for EspressoPass_lstm_atomizer.
func GetEspressoPassLstmAtomizerClass() EspressoPassLstmAtomizerClass {
	return getEspressoPassLstmAtomizerClass()
}

type EspressoPassLstmAtomizerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassLstmAtomizerClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassLstmAtomizerClass) Alloc() EspressoPassLstmAtomizer {
	rv := objc.SendIfResponds[EspressoPassLstmAtomizer](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassLstmAtomizer struct {
	EspressoCustomPass
}

// EspressoPassLstmAtomizerFromID constructs a [EspressoPassLstmAtomizer] from an objc.ID.
func EspressoPassLstmAtomizerFromID(id objc.ID) EspressoPassLstmAtomizer {
	return EspressoPassLstmAtomizer{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_lstm_atomizerFromID is an alias for [EspressoPassLstmAtomizerFromID] for cross-framework compatibility.
func EspressoPass_lstm_atomizerFromID(id objc.ID) EspressoPassLstmAtomizer {
	return EspressoPassLstmAtomizerFromID(id)
}

// Ensure EspressoPassLstmAtomizer implements IEspressoPassLstmAtomizer.
var _ IEspressoPassLstmAtomizer = EspressoPassLstmAtomizer{}

// An interface definition for the [EspressoPassLstmAtomizer] class.
type IEspressoPassLstmAtomizer interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassLstmAtomizer) Init() EspressoPassLstmAtomizer {
	rv := objc.SendIfResponds[EspressoPassLstmAtomizer](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassLstmAtomizer) Autorelease() EspressoPassLstmAtomizer {
	rv := objc.SendIfResponds[EspressoPassLstmAtomizer](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassLstmAtomizer creates a new EspressoPassLstmAtomizer instance.
func NewEspressoPassLstmAtomizer() EspressoPassLstmAtomizer {
	class := getEspressoPassLstmAtomizerClass()
	rv := objc.SendIfResponds[EspressoPassLstmAtomizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}
