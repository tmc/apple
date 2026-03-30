// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_style_transfer_parameterize_transplant] class.
var (
	_EspressoPass_style_transfer_parameterize_transplantClass     EspressoPass_style_transfer_parameterize_transplantClass
	_EspressoPass_style_transfer_parameterize_transplantClassOnce sync.Once
)

func getEspressoPass_style_transfer_parameterize_transplantClass() EspressoPass_style_transfer_parameterize_transplantClass {
	_EspressoPass_style_transfer_parameterize_transplantClassOnce.Do(func() {
		_EspressoPass_style_transfer_parameterize_transplantClass = EspressoPass_style_transfer_parameterize_transplantClass{class: objc.GetClass("EspressoPass_style_transfer_parameterize_transplant")}
	})
	return _EspressoPass_style_transfer_parameterize_transplantClass
}

// GetEspressoPass_style_transfer_parameterize_transplantClass returns the class object for EspressoPass_style_transfer_parameterize_transplant.
func GetEspressoPass_style_transfer_parameterize_transplantClass() EspressoPass_style_transfer_parameterize_transplantClass {
	return getEspressoPass_style_transfer_parameterize_transplantClass()
}

type EspressoPass_style_transfer_parameterize_transplantClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_style_transfer_parameterize_transplantClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_style_transfer_parameterize_transplantClass) Alloc() EspressoPass_style_transfer_parameterize_transplant {
	rv := objc.Send[EspressoPass_style_transfer_parameterize_transplant](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_style_transfer_parameterize_transplant
type EspressoPass_style_transfer_parameterize_transplant struct {
	EspressoCustomPass
}

// EspressoPass_style_transfer_parameterize_transplantFromID constructs a [EspressoPass_style_transfer_parameterize_transplant] from an objc.ID.
func EspressoPass_style_transfer_parameterize_transplantFromID(id objc.ID) EspressoPass_style_transfer_parameterize_transplant {
	return EspressoPass_style_transfer_parameterize_transplant{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// Ensure EspressoPass_style_transfer_parameterize_transplant implements IEspressoPass_style_transfer_parameterize_transplant.
var _ IEspressoPass_style_transfer_parameterize_transplant = EspressoPass_style_transfer_parameterize_transplant{}

// An interface definition for the [EspressoPass_style_transfer_parameterize_transplant] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_style_transfer_parameterize_transplant
type IEspressoPass_style_transfer_parameterize_transplant interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_style_transfer_parameterize_transplant) Init() EspressoPass_style_transfer_parameterize_transplant {
	rv := objc.Send[EspressoPass_style_transfer_parameterize_transplant](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_style_transfer_parameterize_transplant) Autorelease() EspressoPass_style_transfer_parameterize_transplant {
	rv := objc.Send[EspressoPass_style_transfer_parameterize_transplant](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_style_transfer_parameterize_transplant creates a new EspressoPass_style_transfer_parameterize_transplant instance.
func NewEspressoPass_style_transfer_parameterize_transplant() EspressoPass_style_transfer_parameterize_transplant {
	class := getEspressoPass_style_transfer_parameterize_transplantClass()
	rv := objc.Send[EspressoPass_style_transfer_parameterize_transplant](objc.ID(class.class), objc.Sel("new"))
	return rv
}
