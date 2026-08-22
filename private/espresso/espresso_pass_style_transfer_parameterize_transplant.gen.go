// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassStyleTransferParameterizeTransplant] class.
var (
	_EspressoPassStyleTransferParameterizeTransplantClass     EspressoPassStyleTransferParameterizeTransplantClass
	_EspressoPassStyleTransferParameterizeTransplantClassOnce sync.Once
)

func getEspressoPassStyleTransferParameterizeTransplantClass() EspressoPassStyleTransferParameterizeTransplantClass {
	_EspressoPassStyleTransferParameterizeTransplantClassOnce.Do(func() {
		_EspressoPassStyleTransferParameterizeTransplantClass = EspressoPassStyleTransferParameterizeTransplantClass{class: objc.GetClass("EspressoPass_style_transfer_parameterize_transplant")}
	})
	return _EspressoPassStyleTransferParameterizeTransplantClass
}

// GetEspressoPassStyleTransferParameterizeTransplantClass returns the class object for EspressoPass_style_transfer_parameterize_transplant.
func GetEspressoPassStyleTransferParameterizeTransplantClass() EspressoPassStyleTransferParameterizeTransplantClass {
	return getEspressoPassStyleTransferParameterizeTransplantClass()
}

type EspressoPassStyleTransferParameterizeTransplantClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassStyleTransferParameterizeTransplantClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassStyleTransferParameterizeTransplantClass) Alloc() EspressoPassStyleTransferParameterizeTransplant {
	rv := objc.SendIfResponds[EspressoPassStyleTransferParameterizeTransplant](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassStyleTransferParameterizeTransplant struct {
	EspressoCustomPass
}

// EspressoPassStyleTransferParameterizeTransplantFromID constructs a [EspressoPassStyleTransferParameterizeTransplant] from an objc.ID.
func EspressoPassStyleTransferParameterizeTransplantFromID(id objc.ID) EspressoPassStyleTransferParameterizeTransplant {
	return EspressoPassStyleTransferParameterizeTransplant{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_style_transfer_parameterize_transplantFromID is an alias for [EspressoPassStyleTransferParameterizeTransplantFromID] for cross-framework compatibility.
func EspressoPass_style_transfer_parameterize_transplantFromID(id objc.ID) EspressoPassStyleTransferParameterizeTransplant {
	return EspressoPassStyleTransferParameterizeTransplantFromID(id)
}

// Ensure EspressoPassStyleTransferParameterizeTransplant implements IEspressoPassStyleTransferParameterizeTransplant.
var _ IEspressoPassStyleTransferParameterizeTransplant = EspressoPassStyleTransferParameterizeTransplant{}

// An interface definition for the [EspressoPassStyleTransferParameterizeTransplant] class.
type IEspressoPassStyleTransferParameterizeTransplant interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassStyleTransferParameterizeTransplant) Init() EspressoPassStyleTransferParameterizeTransplant {
	rv := objc.SendIfResponds[EspressoPassStyleTransferParameterizeTransplant](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassStyleTransferParameterizeTransplant) Autorelease() EspressoPassStyleTransferParameterizeTransplant {
	rv := objc.SendIfResponds[EspressoPassStyleTransferParameterizeTransplant](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassStyleTransferParameterizeTransplant creates a new EspressoPassStyleTransferParameterizeTransplant instance.
func NewEspressoPassStyleTransferParameterizeTransplant() EspressoPassStyleTransferParameterizeTransplant {
	class := getEspressoPassStyleTransferParameterizeTransplantClass()
	rv := objc.SendIfResponds[EspressoPassStyleTransferParameterizeTransplant](objc.ID(class.class), objc.Sel("new"))
	return rv
}
