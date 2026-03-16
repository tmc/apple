// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_style_transfer_two_nets] class.
var (
	_EspressoPass_style_transfer_two_netsClass     EspressoPass_style_transfer_two_netsClass
	_EspressoPass_style_transfer_two_netsClassOnce sync.Once
)

func getEspressoPass_style_transfer_two_netsClass() EspressoPass_style_transfer_two_netsClass {
	_EspressoPass_style_transfer_two_netsClassOnce.Do(func() {
		_EspressoPass_style_transfer_two_netsClass = EspressoPass_style_transfer_two_netsClass{class: objc.GetClass("EspressoPass_style_transfer_two_nets")}
	})
	return _EspressoPass_style_transfer_two_netsClass
}

// GetEspressoPass_style_transfer_two_netsClass returns the class object for EspressoPass_style_transfer_two_nets.
func GetEspressoPass_style_transfer_two_netsClass() EspressoPass_style_transfer_two_netsClass {
	return getEspressoPass_style_transfer_two_netsClass()
}

type EspressoPass_style_transfer_two_netsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_style_transfer_two_netsClass) Alloc() EspressoPass_style_transfer_two_nets {
	rv := objc.Send[EspressoPass_style_transfer_two_nets](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_style_transfer_two_nets
type EspressoPass_style_transfer_two_nets struct {
	EspressoCustomPass
}

// EspressoPass_style_transfer_two_netsFromID constructs a [EspressoPass_style_transfer_two_nets] from an objc.ID.
func EspressoPass_style_transfer_two_netsFromID(id objc.ID) EspressoPass_style_transfer_two_nets {
	return EspressoPass_style_transfer_two_nets{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_style_transfer_two_nets implements IEspressoPass_style_transfer_two_nets.
var _ IEspressoPass_style_transfer_two_nets = EspressoPass_style_transfer_two_nets{}

// An interface definition for the [EspressoPass_style_transfer_two_nets] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_style_transfer_two_nets
type IEspressoPass_style_transfer_two_nets interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_style_transfer_two_nets) Init() EspressoPass_style_transfer_two_nets {
	rv := objc.Send[EspressoPass_style_transfer_two_nets](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_style_transfer_two_nets) Autorelease() EspressoPass_style_transfer_two_nets {
	rv := objc.Send[EspressoPass_style_transfer_two_nets](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_style_transfer_two_nets creates a new EspressoPass_style_transfer_two_nets instance.
func NewEspressoPass_style_transfer_two_nets() EspressoPass_style_transfer_two_nets {
	class := getEspressoPass_style_transfer_two_netsClass()
	rv := objc.Send[EspressoPass_style_transfer_two_nets](objc.ID(class.class), objc.Sel("new"))
	return rv
}

