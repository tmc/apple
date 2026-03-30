// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_fix_input_underflow] class.
var (
	_EspressoPass_fix_input_underflowClass     EspressoPass_fix_input_underflowClass
	_EspressoPass_fix_input_underflowClassOnce sync.Once
)

func getEspressoPass_fix_input_underflowClass() EspressoPass_fix_input_underflowClass {
	_EspressoPass_fix_input_underflowClassOnce.Do(func() {
		_EspressoPass_fix_input_underflowClass = EspressoPass_fix_input_underflowClass{class: objc.GetClass("EspressoPass_fix_input_underflow")}
	})
	return _EspressoPass_fix_input_underflowClass
}

// GetEspressoPass_fix_input_underflowClass returns the class object for EspressoPass_fix_input_underflow.
func GetEspressoPass_fix_input_underflowClass() EspressoPass_fix_input_underflowClass {
	return getEspressoPass_fix_input_underflowClass()
}

type EspressoPass_fix_input_underflowClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_fix_input_underflowClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_fix_input_underflowClass) Alloc() EspressoPass_fix_input_underflow {
	rv := objc.Send[EspressoPass_fix_input_underflow](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fix_input_underflow
type EspressoPass_fix_input_underflow struct {
	EspressoCustomPass
}

// EspressoPass_fix_input_underflowFromID constructs a [EspressoPass_fix_input_underflow] from an objc.ID.
func EspressoPass_fix_input_underflowFromID(id objc.ID) EspressoPass_fix_input_underflow {
	return EspressoPass_fix_input_underflow{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// Ensure EspressoPass_fix_input_underflow implements IEspressoPass_fix_input_underflow.
var _ IEspressoPass_fix_input_underflow = EspressoPass_fix_input_underflow{}

// An interface definition for the [EspressoPass_fix_input_underflow] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fix_input_underflow
type IEspressoPass_fix_input_underflow interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_fix_input_underflow) Init() EspressoPass_fix_input_underflow {
	rv := objc.Send[EspressoPass_fix_input_underflow](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_fix_input_underflow) Autorelease() EspressoPass_fix_input_underflow {
	rv := objc.Send[EspressoPass_fix_input_underflow](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_fix_input_underflow creates a new EspressoPass_fix_input_underflow instance.
func NewEspressoPass_fix_input_underflow() EspressoPass_fix_input_underflow {
	class := getEspressoPass_fix_input_underflowClass()
	rv := objc.Send[EspressoPass_fix_input_underflow](objc.ID(class.class), objc.Sel("new"))
	return rv
}
