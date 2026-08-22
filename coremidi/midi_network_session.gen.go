// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MIDINetworkSession] class.
var (
	_MIDINetworkSessionClass     MIDINetworkSessionClass
	_MIDINetworkSessionClassOnce sync.Once
)

func getMIDINetworkSessionClass() MIDINetworkSessionClass {
	_MIDINetworkSessionClassOnce.Do(func() {
		_MIDINetworkSessionClass = MIDINetworkSessionClass{class: objc.GetClass("MIDINetworkSession")}
	})
	return _MIDINetworkSessionClass
}

// GetMIDINetworkSessionClass returns the class object for MIDINetworkSession.
func GetMIDINetworkSessionClass() MIDINetworkSessionClass {
	return getMIDINetworkSessionClass()
}

type MIDINetworkSessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MIDINetworkSessionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MIDINetworkSessionClass) Alloc() MIDINetworkSession {
	rv := objc.Send[MIDINetworkSession](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a pairing of a source and destination.
//
// # Overview
//
// A session can have any number of connections. The system broadcasts output
// to all connections, and merges input from multiple connections.
//
// # Configuring a Session
//
//   - [MIDINetworkSession.IsEnabled]: A Boolean value that determines whether the session is enabled.
//   - [MIDINetworkSession.SetEnabled]
//   - [MIDINetworkSession.ConnectionPolicy]: The policy that determines who can connect to this session.
//   - [MIDINetworkSession.SetConnectionPolicy]
//
// # Inspecting a Sessions
//
//   - [MIDINetworkSession.LocalName]: The name of this session’s entity.
//   - [MIDINetworkSession.NetworkName]: The name with which this session advertises itself over Bonjour.
//   - [MIDINetworkSession.NetworkPort]: The session’s UDP port.
//   - [MIDINetworkSession.SourceEndpoint]: Returns the session’s source endpoint.
//   - [MIDINetworkSession.DestinationEndpoint]: Returns the session’s destination endpoint.
//
// # Managing Connections
//
//   - [MIDINetworkSession.Connections]: Returns the session’s set of MIDI network connections.
//   - [MIDINetworkSession.AddConnection]: Adds a new connection to this session.
//   - [MIDINetworkSession.RemoveConnection]: Removes a connection from this session.
//
// # Managing Contacts
//
//   - [MIDINetworkSession.Contacts]: Returns the array of network hosts.
//   - [MIDINetworkSession.AddContact]: Adds a host as a contact.
//   - [MIDINetworkSession.RemoveContact]: Removes a host as a contact.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession
type MIDINetworkSession struct {
	objectivec.Object
}

// MIDINetworkSessionFromID constructs a [MIDINetworkSession] from an objc.ID.
//
// An object that represents a pairing of a source and destination.
func MIDINetworkSessionFromID(id objc.ID) MIDINetworkSession {
	return MIDINetworkSession{objectivec.Object{ID: id}}
}

// NOTE: MIDINetworkSession adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MIDINetworkSession] class.
//
// # Configuring a Session
//
//   - [IMIDINetworkSession.IsEnabled]: A Boolean value that determines whether the session is enabled.
//   - [IMIDINetworkSession.SetEnabled]
//   - [IMIDINetworkSession.ConnectionPolicy]: The policy that determines who can connect to this session.
//   - [IMIDINetworkSession.SetConnectionPolicy]
//
// # Inspecting a Sessions
//
//   - [IMIDINetworkSession.LocalName]: The name of this session’s entity.
//   - [IMIDINetworkSession.NetworkName]: The name with which this session advertises itself over Bonjour.
//   - [IMIDINetworkSession.NetworkPort]: The session’s UDP port.
//   - [IMIDINetworkSession.SourceEndpoint]: Returns the session’s source endpoint.
//   - [IMIDINetworkSession.DestinationEndpoint]: Returns the session’s destination endpoint.
//
// # Managing Connections
//
//   - [IMIDINetworkSession.Connections]: Returns the session’s set of MIDI network connections.
//   - [IMIDINetworkSession.AddConnection]: Adds a new connection to this session.
//   - [IMIDINetworkSession.RemoveConnection]: Removes a connection from this session.
//
// # Managing Contacts
//
//   - [IMIDINetworkSession.Contacts]: Returns the array of network hosts.
//   - [IMIDINetworkSession.AddContact]: Adds a host as a contact.
//   - [IMIDINetworkSession.RemoveContact]: Removes a host as a contact.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession
type IMIDINetworkSession interface {
	objectivec.IObject

	// Topic: Configuring a Session

	// A Boolean value that determines whether the session is enabled.
	IsEnabled() bool
	SetEnabled(value bool)
	// The policy that determines who can connect to this session.
	ConnectionPolicy() MIDINetworkConnectionPolicy
	SetConnectionPolicy(value MIDINetworkConnectionPolicy)

	// Topic: Inspecting a Sessions

	// The name of this session’s entity.
	LocalName() string
	// The name with which this session advertises itself over Bonjour.
	NetworkName() string
	// The session’s UDP port.
	NetworkPort() uint
	// Returns the session’s source endpoint.
	SourceEndpoint() MIDIEndpointRef
	// Returns the session’s destination endpoint.
	DestinationEndpoint() MIDIEndpointRef

	// Topic: Managing Connections

	// Returns the session’s set of MIDI network connections.
	Connections() foundation.INSSet
	// Adds a new connection to this session.
	AddConnection(connection IMIDINetworkConnection) bool
	// Removes a connection from this session.
	RemoveConnection(connection IMIDINetworkConnection) bool

	// Topic: Managing Contacts

	// Returns the array of network hosts.
	Contacts() foundation.INSSet
	// Adds a host as a contact.
	AddContact(contact IMIDINetworkHost) bool
	// Removes a host as a contact.
	RemoveContact(contact IMIDINetworkHost) bool
}

// Init initializes the instance.
func (m MIDINetworkSession) Init() MIDINetworkSession {
	rv := objc.Send[MIDINetworkSession](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MIDINetworkSession) Autorelease() MIDINetworkSession {
	rv := objc.Send[MIDINetworkSession](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDINetworkSession creates a new MIDINetworkSession instance.
func NewMIDINetworkSession() MIDINetworkSession {
	class := getMIDINetworkSessionClass()
	rv := objc.Send[MIDINetworkSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the session’s source endpoint.
//
// # Return Value
//
// The source endpoint.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/sourceEndpoint()
func (m MIDINetworkSession) SourceEndpoint() MIDIEndpointRef {
	rv := objc.Send[MIDIEndpointRef](m.ID, objc.Sel("sourceEndpoint"))
	return MIDIEndpointRef(rv)
}

// Returns the session’s destination endpoint.
//
// # Return Value
//
// The destination endpoint.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/destinationEndpoint()
func (m MIDINetworkSession) DestinationEndpoint() MIDIEndpointRef {
	rv := objc.Send[MIDIEndpointRef](m.ID, objc.Sel("destinationEndpoint"))
	return MIDIEndpointRef(rv)
}

// Returns the session’s set of MIDI network connections.
//
// # Return Value
//
// The set of connections.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/connections()
func (m MIDINetworkSession) Connections() foundation.INSSet {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("connections"))
	return foundation.NSSetFromID(rv)
}

// Adds a new connection to this session.
//
// connection: The connection to add.
//
// # Return Value
//
// true if the session successfully added the connection, otherwise false.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/addConnection(_:)
func (m MIDINetworkSession) AddConnection(connection IMIDINetworkConnection) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("addConnection:"), connection)
	return rv
}

// Removes a connection from this session.
//
// connection: The connection to remove.
//
// # Return Value
//
// true if the session successfully removed the connection, otherwise false.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/removeConnection(_:)
func (m MIDINetworkSession) RemoveConnection(connection IMIDINetworkConnection) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("removeConnection:"), connection)
	return rv
}

// Returns the array of network hosts.
//
// # Return Value
//
// The set of MIDI network host objects.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/contacts()
func (m MIDINetworkSession) Contacts() foundation.INSSet {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("contacts"))
	return foundation.NSSetFromID(rv)
}

// Adds a host as a contact.
//
// contact: The MIDI network host to add.
//
// # Return Value
//
// true if the session successfully added the host, otherwise false.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/addContact(_:)
func (m MIDINetworkSession) AddContact(contact IMIDINetworkHost) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("addContact:"), contact)
	return rv
}

// Removes a host as a contact.
//
// contact: The host to remove.
//
// # Return Value
//
// true if the session successfully removed the host, otherwise false.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/removeContact(_:)
func (m MIDINetworkSession) RemoveContact(contact IMIDINetworkHost) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("removeContact:"), contact)
	return rv
}

// Returns the default singleton session.
//
// # Return Value
//
// The default session.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/default()
func (_MIDINetworkSessionClass MIDINetworkSessionClass) DefaultSession() MIDINetworkSession {
	rv := objc.Send[objc.ID](objc.ID(_MIDINetworkSessionClass.class), objc.Sel("defaultSession"))
	return MIDINetworkSessionFromID(rv)
}

// A Boolean value that determines whether the session is enabled.
//
// # Discussion
//
// The default value is false. Disabled sessions don’t appear on the network
// and can’t initiate or receive connections.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/isEnabled
func (m MIDINetworkSession) IsEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isEnabled"))
	return rv
}
func (m MIDINetworkSession) SetEnabled(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setEnabled:"), value)
}

// The policy that determines who can connect to this session.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/connectionPolicy
func (m MIDINetworkSession) ConnectionPolicy() MIDINetworkConnectionPolicy {
	rv := objc.Send[MIDINetworkConnectionPolicy](m.ID, objc.Sel("connectionPolicy"))
	return MIDINetworkConnectionPolicy(rv)
}
func (m MIDINetworkSession) SetConnectionPolicy(value MIDINetworkConnectionPolicy) {
	objc.Send[struct{}](m.ID, objc.Sel("setConnectionPolicy:"), value)
}

// The name of this session’s entity.
//
// # Discussion
//
// The session’s endpoints inherit this value.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/localName
func (m MIDINetworkSession) LocalName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("localName"))
	return foundation.NSStringFromID(rv).String()
}

// The name with which this session advertises itself over Bonjour.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/networkName
func (m MIDINetworkSession) NetworkName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("networkName"))
	return foundation.NSStringFromID(rv).String()
}

// The session’s UDP port.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/networkPort
func (m MIDINetworkSession) NetworkPort() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("networkPort"))
	return rv
}
