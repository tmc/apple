// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CBL2CAPChannel] class.
var (
	_CBL2CAPChannelClass     CBL2CAPChannelClass
	_CBL2CAPChannelClassOnce sync.Once
)

func getCBL2CAPChannelClass() CBL2CAPChannelClass {
	_CBL2CAPChannelClassOnce.Do(func() {
		_CBL2CAPChannelClass = CBL2CAPChannelClass{class: objc.GetClass("CBL2CAPChannel")}
	})
	return _CBL2CAPChannelClass
}

// GetCBL2CAPChannelClass returns the class object for CBL2CAPChannel.
func GetCBL2CAPChannelClass() CBL2CAPChannelClass {
	return getCBL2CAPChannelClass()
}

type CBL2CAPChannelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CBL2CAPChannelClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CBL2CAPChannelClass) Alloc() CBL2CAPChannel {
	rv := objc.Send[CBL2CAPChannel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A live L2CAP connection to a remote device.
//
// # Accessing Streams
//
//   - [CBL2CAPChannel.InputStream]: The stream used for reading data from the remote peer.
//   - [CBL2CAPChannel.OutputStream]: The stream used for writing data to the peer.
//
// # Accessing the Peer
//
//   - [CBL2CAPChannel.Peer]: The peer connected to the channel.
//
// # Accessing the Protocol/Service Multiplexer
//
//   - [CBL2CAPChannel.PSM]: The PSM of the channel.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBL2CAPChannel
type CBL2CAPChannel struct {
	objectivec.Object
}

// CBL2CAPChannelFromID constructs a [CBL2CAPChannel] from an objc.ID.
//
// A live L2CAP connection to a remote device.
func CBL2CAPChannelFromID(id objc.ID) CBL2CAPChannel {
	return CBL2CAPChannel{objectivec.Object{ID: id}}
}

// NOTE: CBL2CAPChannel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CBL2CAPChannel] class.
//
// # Accessing Streams
//
//   - [ICBL2CAPChannel.InputStream]: The stream used for reading data from the remote peer.
//   - [ICBL2CAPChannel.OutputStream]: The stream used for writing data to the peer.
//
// # Accessing the Peer
//
//   - [ICBL2CAPChannel.Peer]: The peer connected to the channel.
//
// # Accessing the Protocol/Service Multiplexer
//
//   - [ICBL2CAPChannel.PSM]: The PSM of the channel.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBL2CAPChannel
type ICBL2CAPChannel interface {
	objectivec.IObject

	// Topic: Accessing Streams

	// The stream used for reading data from the remote peer.
	InputStream() foundation.InputStream
	// The stream used for writing data to the peer.
	OutputStream() foundation.OutputStream

	// Topic: Accessing the Peer

	// The peer connected to the channel.
	Peer() ICBPeer

	// Topic: Accessing the Protocol/Service Multiplexer

	// The PSM of the channel.
	PSM() CBL2CAPPSM
}

// Init initializes the instance.
func (c CBL2CAPChannel) Init() CBL2CAPChannel {
	rv := objc.Send[CBL2CAPChannel](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CBL2CAPChannel) Autorelease() CBL2CAPChannel {
	rv := objc.Send[CBL2CAPChannel](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBL2CAPChannel creates a new CBL2CAPChannel instance.
func NewCBL2CAPChannel() CBL2CAPChannel {
	class := getCBL2CAPChannelClass()
	rv := objc.Send[CBL2CAPChannel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The stream used for reading data from the remote peer.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBL2CAPChannel/inputStream
func (c CBL2CAPChannel) InputStream() foundation.InputStream {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("inputStream"))
	return foundation.InputStreamFromID(objc.ID(rv))
}

// The stream used for writing data to the peer.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBL2CAPChannel/outputStream
func (c CBL2CAPChannel) OutputStream() foundation.OutputStream {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("outputStream"))
	return foundation.OutputStreamFromID(objc.ID(rv))
}

// The peer connected to the channel.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBL2CAPChannel/peer
func (c CBL2CAPChannel) Peer() ICBPeer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("peer"))
	return CBPeerFromID(objc.ID(rv))
}

// The PSM of the channel.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBL2CAPChannel/psm
func (c CBL2CAPChannel) PSM() CBL2CAPPSM {
	rv := objc.Send[CBL2CAPPSM](c.ID, objc.Sel("PSM"))
	return CBL2CAPPSM(rv)
}
