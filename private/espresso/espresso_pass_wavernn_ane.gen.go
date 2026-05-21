// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassWavernnAne] class.
var (
	_EspressoPassWavernnAneClass     EspressoPassWavernnAneClass
	_EspressoPassWavernnAneClassOnce sync.Once
)

func getEspressoPassWavernnAneClass() EspressoPassWavernnAneClass {
	_EspressoPassWavernnAneClassOnce.Do(func() {
		_EspressoPassWavernnAneClass = EspressoPassWavernnAneClass{class: objc.GetClass("EspressoPass_wavernn_ane")}
	})
	return _EspressoPassWavernnAneClass
}

// GetEspressoPassWavernnAneClass returns the class object for EspressoPass_wavernn_ane.
func GetEspressoPassWavernnAneClass() EspressoPassWavernnAneClass {
	return getEspressoPassWavernnAneClass()
}

type EspressoPassWavernnAneClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassWavernnAneClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassWavernnAneClass) Alloc() EspressoPassWavernnAne {
	rv := objc.Send[EspressoPassWavernnAne](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassWavernnAne struct {
	EspressoCustomPass
}

// EspressoPassWavernnAneFromID constructs a [EspressoPassWavernnAne] from an objc.ID.
func EspressoPassWavernnAneFromID(id objc.ID) EspressoPassWavernnAne {
	return EspressoPassWavernnAne{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_wavernn_aneFromID is an alias for [EspressoPassWavernnAneFromID] for cross-framework compatibility.
func EspressoPass_wavernn_aneFromID(id objc.ID) EspressoPassWavernnAne {
	return EspressoPassWavernnAneFromID(id)
}

// Ensure EspressoPassWavernnAne implements IEspressoPassWavernnAne.
var _ IEspressoPassWavernnAne = EspressoPassWavernnAne{}

// An interface definition for the [EspressoPassWavernnAne] class.
type IEspressoPassWavernnAne interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassWavernnAne) Init() EspressoPassWavernnAne {
	rv := objc.Send[EspressoPassWavernnAne](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassWavernnAne) Autorelease() EspressoPassWavernnAne {
	rv := objc.Send[EspressoPassWavernnAne](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassWavernnAne creates a new EspressoPassWavernnAne instance.
func NewEspressoPassWavernnAne() EspressoPassWavernnAne {
	class := getEspressoPassWavernnAneClass()
	rv := objc.Send[EspressoPassWavernnAne](objc.ID(class.class), objc.Sel("new"))
	return rv
}
