// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassMakeFullyConv] class.
var (
	_EspressoPassMakeFullyConvClass     EspressoPassMakeFullyConvClass
	_EspressoPassMakeFullyConvClassOnce sync.Once
)

func getEspressoPassMakeFullyConvClass() EspressoPassMakeFullyConvClass {
	_EspressoPassMakeFullyConvClassOnce.Do(func() {
		_EspressoPassMakeFullyConvClass = EspressoPassMakeFullyConvClass{class: objc.GetClass("EspressoPass_make_fully_conv")}
	})
	return _EspressoPassMakeFullyConvClass
}

// GetEspressoPassMakeFullyConvClass returns the class object for EspressoPass_make_fully_conv.
func GetEspressoPassMakeFullyConvClass() EspressoPassMakeFullyConvClass {
	return getEspressoPassMakeFullyConvClass()
}

type EspressoPassMakeFullyConvClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassMakeFullyConvClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassMakeFullyConvClass) Alloc() EspressoPassMakeFullyConv {
	rv := objc.Send[EspressoPassMakeFullyConv](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_make_fully_conv
type EspressoPassMakeFullyConv struct {
	EspressoCustomPass
}

// EspressoPassMakeFullyConvFromID constructs a [EspressoPassMakeFullyConv] from an objc.ID.
func EspressoPassMakeFullyConvFromID(id objc.ID) EspressoPassMakeFullyConv {
	return EspressoPassMakeFullyConv{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_make_fully_convFromID is an alias for [EspressoPassMakeFullyConvFromID] for cross-framework compatibility.
func EspressoPass_make_fully_convFromID(id objc.ID) EspressoPassMakeFullyConv {
	return EspressoPassMakeFullyConvFromID(id)
}

// Ensure EspressoPassMakeFullyConv implements IEspressoPassMakeFullyConv.
var _ IEspressoPassMakeFullyConv = EspressoPassMakeFullyConv{}

// An interface definition for the [EspressoPassMakeFullyConv] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_make_fully_conv
type IEspressoPassMakeFullyConv interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassMakeFullyConv) Init() EspressoPassMakeFullyConv {
	rv := objc.Send[EspressoPassMakeFullyConv](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassMakeFullyConv) Autorelease() EspressoPassMakeFullyConv {
	rv := objc.Send[EspressoPassMakeFullyConv](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassMakeFullyConv creates a new EspressoPassMakeFullyConv instance.
func NewEspressoPassMakeFullyConv() EspressoPassMakeFullyConv {
	class := getEspressoPassMakeFullyConvClass()
	rv := objc.Send[EspressoPassMakeFullyConv](objc.ID(class.class), objc.Sel("new"))
	return rv
}
