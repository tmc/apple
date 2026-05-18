// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassRemoveSqueeze] class.
var (
	_EspressoPassRemoveSqueezeClass     EspressoPassRemoveSqueezeClass
	_EspressoPassRemoveSqueezeClassOnce sync.Once
)

func getEspressoPassRemoveSqueezeClass() EspressoPassRemoveSqueezeClass {
	_EspressoPassRemoveSqueezeClassOnce.Do(func() {
		_EspressoPassRemoveSqueezeClass = EspressoPassRemoveSqueezeClass{class: objc.GetClass("EspressoPass_remove_squeeze")}
	})
	return _EspressoPassRemoveSqueezeClass
}

// GetEspressoPassRemoveSqueezeClass returns the class object for EspressoPass_remove_squeeze.
func GetEspressoPassRemoveSqueezeClass() EspressoPassRemoveSqueezeClass {
	return getEspressoPassRemoveSqueezeClass()
}

type EspressoPassRemoveSqueezeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassRemoveSqueezeClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassRemoveSqueezeClass) Alloc() EspressoPassRemoveSqueeze {
	rv := objc.Send[EspressoPassRemoveSqueeze](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_remove_squeeze
type EspressoPassRemoveSqueeze struct {
	EspressoCustomPass
}

// EspressoPassRemoveSqueezeFromID constructs a [EspressoPassRemoveSqueeze] from an objc.ID.
func EspressoPassRemoveSqueezeFromID(id objc.ID) EspressoPassRemoveSqueeze {
	return EspressoPassRemoveSqueeze{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_remove_squeezeFromID is an alias for [EspressoPassRemoveSqueezeFromID] for cross-framework compatibility.
func EspressoPass_remove_squeezeFromID(id objc.ID) EspressoPassRemoveSqueeze {
	return EspressoPassRemoveSqueezeFromID(id)
}

// Ensure EspressoPassRemoveSqueeze implements IEspressoPassRemoveSqueeze.
var _ IEspressoPassRemoveSqueeze = EspressoPassRemoveSqueeze{}

// An interface definition for the [EspressoPassRemoveSqueeze] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_remove_squeeze
type IEspressoPassRemoveSqueeze interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassRemoveSqueeze) Init() EspressoPassRemoveSqueeze {
	rv := objc.Send[EspressoPassRemoveSqueeze](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassRemoveSqueeze) Autorelease() EspressoPassRemoveSqueeze {
	rv := objc.Send[EspressoPassRemoveSqueeze](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassRemoveSqueeze creates a new EspressoPassRemoveSqueeze instance.
func NewEspressoPassRemoveSqueeze() EspressoPassRemoveSqueeze {
	class := getEspressoPassRemoveSqueezeClass()
	rv := objc.Send[EspressoPassRemoveSqueeze](objc.ID(class.class), objc.Sel("new"))
	return rv
}
