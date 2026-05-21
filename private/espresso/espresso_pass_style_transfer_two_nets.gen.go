// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassStyleTransferTwoNets] class.
var (
	_EspressoPassStyleTransferTwoNetsClass     EspressoPassStyleTransferTwoNetsClass
	_EspressoPassStyleTransferTwoNetsClassOnce sync.Once
)

func getEspressoPassStyleTransferTwoNetsClass() EspressoPassStyleTransferTwoNetsClass {
	_EspressoPassStyleTransferTwoNetsClassOnce.Do(func() {
		_EspressoPassStyleTransferTwoNetsClass = EspressoPassStyleTransferTwoNetsClass{class: objc.GetClass("EspressoPass_style_transfer_two_nets")}
	})
	return _EspressoPassStyleTransferTwoNetsClass
}

// GetEspressoPassStyleTransferTwoNetsClass returns the class object for EspressoPass_style_transfer_two_nets.
func GetEspressoPassStyleTransferTwoNetsClass() EspressoPassStyleTransferTwoNetsClass {
	return getEspressoPassStyleTransferTwoNetsClass()
}

type EspressoPassStyleTransferTwoNetsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassStyleTransferTwoNetsClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassStyleTransferTwoNetsClass) Alloc() EspressoPassStyleTransferTwoNets {
	rv := objc.Send[EspressoPassStyleTransferTwoNets](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassStyleTransferTwoNets struct {
	EspressoCustomPass
}

// EspressoPassStyleTransferTwoNetsFromID constructs a [EspressoPassStyleTransferTwoNets] from an objc.ID.
func EspressoPassStyleTransferTwoNetsFromID(id objc.ID) EspressoPassStyleTransferTwoNets {
	return EspressoPassStyleTransferTwoNets{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_style_transfer_two_netsFromID is an alias for [EspressoPassStyleTransferTwoNetsFromID] for cross-framework compatibility.
func EspressoPass_style_transfer_two_netsFromID(id objc.ID) EspressoPassStyleTransferTwoNets {
	return EspressoPassStyleTransferTwoNetsFromID(id)
}

// Ensure EspressoPassStyleTransferTwoNets implements IEspressoPassStyleTransferTwoNets.
var _ IEspressoPassStyleTransferTwoNets = EspressoPassStyleTransferTwoNets{}

// An interface definition for the [EspressoPassStyleTransferTwoNets] class.
type IEspressoPassStyleTransferTwoNets interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassStyleTransferTwoNets) Init() EspressoPassStyleTransferTwoNets {
	rv := objc.Send[EspressoPassStyleTransferTwoNets](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassStyleTransferTwoNets) Autorelease() EspressoPassStyleTransferTwoNets {
	rv := objc.Send[EspressoPassStyleTransferTwoNets](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassStyleTransferTwoNets creates a new EspressoPassStyleTransferTwoNets instance.
func NewEspressoPassStyleTransferTwoNets() EspressoPassStyleTransferTwoNets {
	class := getEspressoPassStyleTransferTwoNetsClass()
	rv := objc.Send[EspressoPassStyleTransferTwoNets](objc.ID(class.class), objc.Sel("new"))
	return rv
}
