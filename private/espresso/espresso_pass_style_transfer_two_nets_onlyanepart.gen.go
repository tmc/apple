// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_style_transfer_two_nets_onlyanepart] class.
var (
	_EspressoPass_style_transfer_two_nets_onlyanepartClass     EspressoPass_style_transfer_two_nets_onlyanepartClass
	_EspressoPass_style_transfer_two_nets_onlyanepartClassOnce sync.Once
)

func getEspressoPass_style_transfer_two_nets_onlyanepartClass() EspressoPass_style_transfer_two_nets_onlyanepartClass {
	_EspressoPass_style_transfer_two_nets_onlyanepartClassOnce.Do(func() {
		_EspressoPass_style_transfer_two_nets_onlyanepartClass = EspressoPass_style_transfer_two_nets_onlyanepartClass{class: objc.GetClass("EspressoPass_style_transfer_two_nets_onlyanepart")}
	})
	return _EspressoPass_style_transfer_two_nets_onlyanepartClass
}

// GetEspressoPass_style_transfer_two_nets_onlyanepartClass returns the class object for EspressoPass_style_transfer_two_nets_onlyanepart.
func GetEspressoPass_style_transfer_two_nets_onlyanepartClass() EspressoPass_style_transfer_two_nets_onlyanepartClass {
	return getEspressoPass_style_transfer_two_nets_onlyanepartClass()
}

type EspressoPass_style_transfer_two_nets_onlyanepartClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_style_transfer_two_nets_onlyanepartClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_style_transfer_two_nets_onlyanepartClass) Alloc() EspressoPass_style_transfer_two_nets_onlyanepart {
	rv := objc.Send[EspressoPass_style_transfer_two_nets_onlyanepart](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_style_transfer_two_nets_onlyanepart
type EspressoPass_style_transfer_two_nets_onlyanepart struct {
	EspressoCustomPass
}

// EspressoPass_style_transfer_two_nets_onlyanepartFromID constructs a [EspressoPass_style_transfer_two_nets_onlyanepart] from an objc.ID.
func EspressoPass_style_transfer_two_nets_onlyanepartFromID(id objc.ID) EspressoPass_style_transfer_two_nets_onlyanepart {
	return EspressoPass_style_transfer_two_nets_onlyanepart{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// Ensure EspressoPass_style_transfer_two_nets_onlyanepart implements IEspressoPass_style_transfer_two_nets_onlyanepart.
var _ IEspressoPass_style_transfer_two_nets_onlyanepart = EspressoPass_style_transfer_two_nets_onlyanepart{}

// An interface definition for the [EspressoPass_style_transfer_two_nets_onlyanepart] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_style_transfer_two_nets_onlyanepart
type IEspressoPass_style_transfer_two_nets_onlyanepart interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_style_transfer_two_nets_onlyanepart) Init() EspressoPass_style_transfer_two_nets_onlyanepart {
	rv := objc.Send[EspressoPass_style_transfer_two_nets_onlyanepart](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_style_transfer_two_nets_onlyanepart) Autorelease() EspressoPass_style_transfer_two_nets_onlyanepart {
	rv := objc.Send[EspressoPass_style_transfer_two_nets_onlyanepart](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_style_transfer_two_nets_onlyanepart creates a new EspressoPass_style_transfer_two_nets_onlyanepart instance.
func NewEspressoPass_style_transfer_two_nets_onlyanepart() EspressoPass_style_transfer_two_nets_onlyanepart {
	class := getEspressoPass_style_transfer_two_nets_onlyanepartClass()
	rv := objc.Send[EspressoPass_style_transfer_two_nets_onlyanepart](objc.ID(class.class), objc.Sel("new"))
	return rv
}
