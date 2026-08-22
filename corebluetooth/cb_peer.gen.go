// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CBPeer] class.
var (
	_CBPeerClass     CBPeerClass
	_CBPeerClassOnce sync.Once
)

func getCBPeerClass() CBPeerClass {
	_CBPeerClassOnce.Do(func() {
		_CBPeerClass = CBPeerClass{class: objc.GetClass("CBPeer")}
	})
	return _CBPeerClass
}

// GetCBPeerClass returns the class object for CBPeer.
func GetCBPeerClass() CBPeerClass {
	return getCBPeerClass()
}

type CBPeerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CBPeerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CBPeerClass) Alloc() CBPeer {
	rv := objc.Send[CBPeer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a remote device.
//
// # Overview
//
// The [CBPeer] class is an abstract base class that defines common behavior
// for objects representing remote devices. You typically don’t create
// instances of either [CBPeer] or its concrete subclasses. Instead, the
// system creates them for you during the process of peer discovery.
//
// Your app takes the role of either a central (by creating an instance of
// [CBCentralManager]) or a peripheral (by creating an instance of
// [CBPeripheralManager]), and interacts through the manager with remote
// devices in the opposite role. During the process of peer discovery, where a
// central device scans for peripherals advertising services, the system
// creates objects from the concrete subclasses of [CBPeer] to represent
// discovered remote devices. The concrete subclasses of [CBPeer] are
// [CBPeripheral] and [CBCentral].
//
// # Identifying a Peer
//
//   - [CBPeer.Identifier]: The UUID associated with the peer.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeer
type CBPeer struct {
	objectivec.Object
}

// CBPeerFromID constructs a [CBPeer] from an objc.ID.
//
// An object that represents a remote device.
func CBPeerFromID(id objc.ID) CBPeer {
	return CBPeer{objectivec.Object{ID: id}}
}

// NOTE: CBPeer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CBPeer] class.
//
// # Identifying a Peer
//
//   - [ICBPeer.Identifier]: The UUID associated with the peer.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeer
type ICBPeer interface {
	objectivec.IObject

	// Topic: Identifying a Peer

	// The UUID associated with the peer.
	Identifier() foundation.NSUUID
}

// Init initializes the instance.
func (c CBPeer) Init() CBPeer {
	rv := objc.Send[CBPeer](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CBPeer) Autorelease() CBPeer {
	rv := objc.Send[CBPeer](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBPeer creates a new CBPeer instance.
func NewCBPeer() CBPeer {
	class := getCBPeerClass()
	rv := objc.Send[CBPeer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The UUID associated with the peer.
//
// # Discussion
//
// The value of this property represents the unique identifier of the peer.
// The first time a local manager encounters a peer, the system assigns the
// peer a UUID, represented by a new [UUID] object. Peers use [UUID] instances
// to identify themselves, instead of by the [CBUUID] objects that identify a
// peripheral’s services, characteristics, and descriptors.
//
// For more information, see [Core Bluetooth Programming Guide].
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeer/identifier
//
// [Core Bluetooth Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/CoreBluetooth_concepts/AboutCoreBluetooth/Introduction.html#//apple_ref/doc/uid/TP40013257
// [UUID]: https://developer.apple.com/documentation/Foundation/UUID
func (c CBPeer) Identifier() foundation.NSUUID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("identifier"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
