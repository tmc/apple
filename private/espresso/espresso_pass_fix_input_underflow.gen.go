// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassFixInputUnderflow] class.
var (
	_EspressoPassFixInputUnderflowClass     EspressoPassFixInputUnderflowClass
	_EspressoPassFixInputUnderflowClassOnce sync.Once
)

func getEspressoPassFixInputUnderflowClass() EspressoPassFixInputUnderflowClass {
	_EspressoPassFixInputUnderflowClassOnce.Do(func() {
		_EspressoPassFixInputUnderflowClass = EspressoPassFixInputUnderflowClass{class: objc.GetClass("EspressoPass_fix_input_underflow")}
	})
	return _EspressoPassFixInputUnderflowClass
}

// GetEspressoPassFixInputUnderflowClass returns the class object for EspressoPass_fix_input_underflow.
func GetEspressoPassFixInputUnderflowClass() EspressoPassFixInputUnderflowClass {
	return getEspressoPassFixInputUnderflowClass()
}

type EspressoPassFixInputUnderflowClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassFixInputUnderflowClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassFixInputUnderflowClass) Alloc() EspressoPassFixInputUnderflow {
	rv := objc.Send[EspressoPassFixInputUnderflow](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fix_input_underflow
type EspressoPassFixInputUnderflow struct {
	EspressoCustomPass
}

// EspressoPassFixInputUnderflowFromID constructs a [EspressoPassFixInputUnderflow] from an objc.ID.
func EspressoPassFixInputUnderflowFromID(id objc.ID) EspressoPassFixInputUnderflow {
	return EspressoPassFixInputUnderflow{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_fix_input_underflowFromID is an alias for [EspressoPassFixInputUnderflowFromID] for cross-framework compatibility.
func EspressoPass_fix_input_underflowFromID(id objc.ID) EspressoPassFixInputUnderflow {
	return EspressoPassFixInputUnderflowFromID(id)
}

// Ensure EspressoPassFixInputUnderflow implements IEspressoPassFixInputUnderflow.
var _ IEspressoPassFixInputUnderflow = EspressoPassFixInputUnderflow{}

// An interface definition for the [EspressoPassFixInputUnderflow] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fix_input_underflow
type IEspressoPassFixInputUnderflow interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassFixInputUnderflow) Init() EspressoPassFixInputUnderflow {
	rv := objc.Send[EspressoPassFixInputUnderflow](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassFixInputUnderflow) Autorelease() EspressoPassFixInputUnderflow {
	rv := objc.Send[EspressoPassFixInputUnderflow](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassFixInputUnderflow creates a new EspressoPassFixInputUnderflow instance.
func NewEspressoPassFixInputUnderflow() EspressoPassFixInputUnderflow {
	class := getEspressoPassFixInputUnderflowClass()
	rv := objc.Send[EspressoPassFixInputUnderflow](objc.ID(class.class), objc.Sel("new"))
	return rv
}
