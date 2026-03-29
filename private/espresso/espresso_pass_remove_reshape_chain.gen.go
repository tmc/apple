// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_remove_reshape_chain] class.
var (
	_EspressoPass_remove_reshape_chainClass     EspressoPass_remove_reshape_chainClass
	_EspressoPass_remove_reshape_chainClassOnce sync.Once
)

func getEspressoPass_remove_reshape_chainClass() EspressoPass_remove_reshape_chainClass {
	_EspressoPass_remove_reshape_chainClassOnce.Do(func() {
		_EspressoPass_remove_reshape_chainClass = EspressoPass_remove_reshape_chainClass{class: objc.GetClass("EspressoPass_remove_reshape_chain")}
	})
	return _EspressoPass_remove_reshape_chainClass
}

// GetEspressoPass_remove_reshape_chainClass returns the class object for EspressoPass_remove_reshape_chain.
func GetEspressoPass_remove_reshape_chainClass() EspressoPass_remove_reshape_chainClass {
	return getEspressoPass_remove_reshape_chainClass()
}

type EspressoPass_remove_reshape_chainClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_remove_reshape_chainClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_remove_reshape_chainClass) Alloc() EspressoPass_remove_reshape_chain {
	rv := objc.Send[EspressoPass_remove_reshape_chain](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_remove_reshape_chain
type EspressoPass_remove_reshape_chain struct {
	EspressoCustomPass
}

// EspressoPass_remove_reshape_chainFromID constructs a [EspressoPass_remove_reshape_chain] from an objc.ID.
func EspressoPass_remove_reshape_chainFromID(id objc.ID) EspressoPass_remove_reshape_chain {
	return EspressoPass_remove_reshape_chain{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_remove_reshape_chain implements IEspressoPass_remove_reshape_chain.
var _ IEspressoPass_remove_reshape_chain = EspressoPass_remove_reshape_chain{}

// An interface definition for the [EspressoPass_remove_reshape_chain] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_remove_reshape_chain
type IEspressoPass_remove_reshape_chain interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_remove_reshape_chain) Init() EspressoPass_remove_reshape_chain {
	rv := objc.Send[EspressoPass_remove_reshape_chain](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_remove_reshape_chain) Autorelease() EspressoPass_remove_reshape_chain {
	rv := objc.Send[EspressoPass_remove_reshape_chain](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_remove_reshape_chain creates a new EspressoPass_remove_reshape_chain instance.
func NewEspressoPass_remove_reshape_chain() EspressoPass_remove_reshape_chain {
	class := getEspressoPass_remove_reshape_chainClass()
	rv := objc.Send[EspressoPass_remove_reshape_chain](objc.ID(class.class), objc.Sel("new"))
	return rv
}

