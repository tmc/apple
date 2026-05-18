// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassStyleTransferTwoNetsOnlyanepart] class.
var (
	_EspressoPassStyleTransferTwoNetsOnlyanepartClass     EspressoPassStyleTransferTwoNetsOnlyanepartClass
	_EspressoPassStyleTransferTwoNetsOnlyanepartClassOnce sync.Once
)

func getEspressoPassStyleTransferTwoNetsOnlyanepartClass() EspressoPassStyleTransferTwoNetsOnlyanepartClass {
	_EspressoPassStyleTransferTwoNetsOnlyanepartClassOnce.Do(func() {
		_EspressoPassStyleTransferTwoNetsOnlyanepartClass = EspressoPassStyleTransferTwoNetsOnlyanepartClass{class: objc.GetClass("EspressoPass_style_transfer_two_nets_onlyanepart")}
	})
	return _EspressoPassStyleTransferTwoNetsOnlyanepartClass
}

// GetEspressoPassStyleTransferTwoNetsOnlyanepartClass returns the class object for EspressoPass_style_transfer_two_nets_onlyanepart.
func GetEspressoPassStyleTransferTwoNetsOnlyanepartClass() EspressoPassStyleTransferTwoNetsOnlyanepartClass {
	return getEspressoPassStyleTransferTwoNetsOnlyanepartClass()
}

type EspressoPassStyleTransferTwoNetsOnlyanepartClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassStyleTransferTwoNetsOnlyanepartClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassStyleTransferTwoNetsOnlyanepartClass) Alloc() EspressoPassStyleTransferTwoNetsOnlyanepart {
	rv := objc.Send[EspressoPassStyleTransferTwoNetsOnlyanepart](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_style_transfer_two_nets_onlyanepart
type EspressoPassStyleTransferTwoNetsOnlyanepart struct {
	EspressoCustomPass
}

// EspressoPassStyleTransferTwoNetsOnlyanepartFromID constructs a [EspressoPassStyleTransferTwoNetsOnlyanepart] from an objc.ID.
func EspressoPassStyleTransferTwoNetsOnlyanepartFromID(id objc.ID) EspressoPassStyleTransferTwoNetsOnlyanepart {
	return EspressoPassStyleTransferTwoNetsOnlyanepart{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_style_transfer_two_nets_onlyanepartFromID is an alias for [EspressoPassStyleTransferTwoNetsOnlyanepartFromID] for cross-framework compatibility.
func EspressoPass_style_transfer_two_nets_onlyanepartFromID(id objc.ID) EspressoPassStyleTransferTwoNetsOnlyanepart {
	return EspressoPassStyleTransferTwoNetsOnlyanepartFromID(id)
}

// Ensure EspressoPassStyleTransferTwoNetsOnlyanepart implements IEspressoPassStyleTransferTwoNetsOnlyanepart.
var _ IEspressoPassStyleTransferTwoNetsOnlyanepart = EspressoPassStyleTransferTwoNetsOnlyanepart{}

// An interface definition for the [EspressoPassStyleTransferTwoNetsOnlyanepart] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_style_transfer_two_nets_onlyanepart
type IEspressoPassStyleTransferTwoNetsOnlyanepart interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassStyleTransferTwoNetsOnlyanepart) Init() EspressoPassStyleTransferTwoNetsOnlyanepart {
	rv := objc.Send[EspressoPassStyleTransferTwoNetsOnlyanepart](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassStyleTransferTwoNetsOnlyanepart) Autorelease() EspressoPassStyleTransferTwoNetsOnlyanepart {
	rv := objc.Send[EspressoPassStyleTransferTwoNetsOnlyanepart](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassStyleTransferTwoNetsOnlyanepart creates a new EspressoPassStyleTransferTwoNetsOnlyanepart instance.
func NewEspressoPassStyleTransferTwoNetsOnlyanepart() EspressoPassStyleTransferTwoNetsOnlyanepart {
	class := getEspressoPassStyleTransferTwoNetsOnlyanepartClass()
	rv := objc.Send[EspressoPassStyleTransferTwoNetsOnlyanepart](objc.ID(class.class), objc.Sel("new"))
	return rv
}
