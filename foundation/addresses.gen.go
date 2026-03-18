// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Addresses] class.
var (
	_AddressesClass     AddressesClass
	_AddressesClassOnce sync.Once
)

func getAddressesClass() AddressesClass {
	_AddressesClassOnce.Do(func() {
		_AddressesClass = AddressesClass{class: objc.GetClass("addresses")}
	})
	return _AddressesClass
}

// GetAddressesClass returns the class object for addresses.
func GetAddressesClass() AddressesClass {
	return getAddressesClass()
}

type AddressesClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AddressesClass) Alloc() Addresses {
	rv := objc.Send[Addresses](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSHost/addresses-c.ivar
type Addresses struct {
	objectivec.Object
}

// AddressesFromID constructs a [Addresses] from an objc.ID.
func AddressesFromID(id objc.ID) Addresses {
	return Addresses{objectivec.Object{ID: id}}
}
// Ensure Addresses implements IAddresses.
var _ IAddresses = Addresses{}

// An interface definition for the [Addresses] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSHost/addresses-c.ivar
type IAddresses interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a Addresses) Init() Addresses {
	rv := objc.Send[Addresses](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a Addresses) Autorelease() Addresses {
	rv := objc.Send[Addresses](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAddresses creates a new Addresses instance.
func NewAddresses() Addresses {
	class := getAddressesClass()
	rv := objc.Send[Addresses](objc.ID(class.class), objc.Sel("new"))
	return rv
}

