// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MIDINetworkConnection] class.
var (
	_MIDINetworkConnectionClass     MIDINetworkConnectionClass
	_MIDINetworkConnectionClassOnce sync.Once
)

func getMIDINetworkConnectionClass() MIDINetworkConnectionClass {
	_MIDINetworkConnectionClassOnce.Do(func() {
		_MIDINetworkConnectionClass = MIDINetworkConnectionClass{class: objc.GetClass("MIDINetworkConnection")}
	})
	return _MIDINetworkConnectionClass
}

// GetMIDINetworkConnectionClass returns the class object for MIDINetworkConnection.
func GetMIDINetworkConnectionClass() MIDINetworkConnectionClass {
	return getMIDINetworkConnectionClass()
}

type MIDINetworkConnectionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MIDINetworkConnectionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MIDINetworkConnectionClass) Alloc() MIDINetworkConnection {
	rv := objc.Send[MIDINetworkConnection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that connects a session to a host.
//
// # Accessing Connections
//
//   - [MIDINetworkConnection.Host]: The host connection.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkConnection
type MIDINetworkConnection struct {
	objectivec.Object
}

// MIDINetworkConnectionFromID constructs a [MIDINetworkConnection] from an objc.ID.
//
// An object that connects a session to a host.
func MIDINetworkConnectionFromID(id objc.ID) MIDINetworkConnection {
	return MIDINetworkConnection{objectivec.Object{ID: id}}
}

// NOTE: MIDINetworkConnection adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MIDINetworkConnection] class.
//
// # Accessing Connections
//
//   - [IMIDINetworkConnection.Host]: The host connection.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkConnection
type IMIDINetworkConnection interface {
	objectivec.IObject

	// Topic: Accessing Connections

	// The host connection.
	Host() IMIDINetworkHost
}

// Init initializes the instance.
func (m MIDINetworkConnection) Init() MIDINetworkConnection {
	rv := objc.Send[MIDINetworkConnection](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MIDINetworkConnection) Autorelease() MIDINetworkConnection {
	rv := objc.Send[MIDINetworkConnection](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDINetworkConnection creates a new MIDINetworkConnection instance.
func NewMIDINetworkConnection() MIDINetworkConnection {
	class := getMIDINetworkConnectionClass()
	rv := objc.Send[MIDINetworkConnection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a connection to the specified host.
//
// host: The host with which to establish a connection.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkConnection/init(host:)
func NewMIDINetworkConnectionWithHost(host IMIDINetworkHost) MIDINetworkConnection {
	rv := objc.Send[objc.ID](objc.ID(getMIDINetworkConnectionClass().class), objc.Sel("connectionWithHost:"), host)
	return MIDINetworkConnectionFromID(rv)
}

// The host connection.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkConnection/host
func (m MIDINetworkConnection) Host() IMIDINetworkHost {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("host"))
	return MIDINetworkHostFromID(objc.ID(rv))
}
