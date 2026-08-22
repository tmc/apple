// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [CBCentral] class.
var (
	_CBCentralClass     CBCentralClass
	_CBCentralClassOnce sync.Once
)

func getCBCentralClass() CBCentralClass {
	_CBCentralClassOnce.Do(func() {
		_CBCentralClass = CBCentralClass{class: objc.GetClass("CBCentral")}
	})
	return _CBCentralClass
}

// GetCBCentralClass returns the class object for CBCentral.
func GetCBCentralClass() CBCentralClass {
	return getCBCentralClass()
}

type CBCentralClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CBCentralClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CBCentralClass) Alloc() CBCentral {
	rv := objc.Send[CBCentral](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A remote device connected to a local app, which is acting as a peripheral.
//
// # Overview
//
// The [CBCentral] class represents remote central devices (or centrals) that
// have connected to an app implementing the peripheral role on a local
// device. Remote centrals use universally unique identifiers (UUIDs),
// represented by [NSUUID] objects, to identify themselves.
//
// # Identifying a Remote Central
//
//   - [CBCentral.MaximumUpdateValueLength]: The maximum amount of data, in bytes, that the central can receive in a single notification or indication.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentral
//
// [NSUUID]: https://developer.apple.com/documentation/Foundation/NSUUID
type CBCentral struct {
	CBPeer
}

// CBCentralFromID constructs a [CBCentral] from an objc.ID.
//
// A remote device connected to a local app, which is acting as a peripheral.
func CBCentralFromID(id objc.ID) CBCentral {
	return CBCentral{CBPeer: CBPeerFromID(id)}
}

// NOTE: CBCentral adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CBCentral] class.
//
// # Identifying a Remote Central
//
//   - [ICBCentral.MaximumUpdateValueLength]: The maximum amount of data, in bytes, that the central can receive in a single notification or indication.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentral
type ICBCentral interface {
	ICBPeer

	// Topic: Identifying a Remote Central

	// The maximum amount of data, in bytes, that the central can receive in a single notification or indication.
	MaximumUpdateValueLength() uint
}

// Init initializes the instance.
func (c CBCentral) Init() CBCentral {
	rv := objc.Send[CBCentral](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CBCentral) Autorelease() CBCentral {
	rv := objc.Send[CBCentral](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBCentral creates a new CBCentral instance.
func NewCBCentral() CBCentral {
	class := getCBCentralClass()
	rv := objc.Send[CBCentral](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The maximum amount of data, in bytes, that the central can receive in a
// single notification or indication.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentral/maximumUpdateValueLength
func (c CBCentral) MaximumUpdateValueLength() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("maximumUpdateValueLength"))
	return rv
}
