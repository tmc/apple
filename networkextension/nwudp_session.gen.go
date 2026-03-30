// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NWUDPSession] class.
var (
	_NWUDPSessionClass     NWUDPSessionClass
	_NWUDPSessionClassOnce sync.Once
)

func getNWUDPSessionClass() NWUDPSessionClass {
	_NWUDPSessionClassOnce.Do(func() {
		_NWUDPSessionClass = NWUDPSessionClass{class: objc.GetClass("NWUDPSession")}
	})
	return _NWUDPSessionClass
}

// GetNWUDPSessionClass returns the class object for NWUDPSession.
func GetNWUDPSessionClass() NWUDPSessionClass {
	return getNWUDPSessionClass()
}

type NWUDPSessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NWUDPSessionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NWUDPSessionClass) Alloc() NWUDPSession {
	rv := objc.Send[NWUDPSession](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object to manage a UDP session to a network endpoint.
//
// # Overview
//
// Since UDP does not include a handshake with the remote endpoint as part of
// its protocol, it is up to the client of the UDP session to provide feedback
// on the viability of the current endpoint. If a session is opened to a
// hostname, the system will resolve that hostname into potentially several IP
// addresses. Once the session state is [NWUDPSessionStateReady], the client
// should try to write and read datagrams. If there is no response from the
// remote endpoint, the client can try the next address that was resolved
// using `tryNextResolvedEndpoint`.
//
// # Monitoring the session state
//
//   - [NWUDPSession.State]: The current state of the UDP session.
//   - [NWUDPSession.Viable]: The viability of a UDP session represents whether or not data can be transferred.
//
// # Selecting remote endpoints
//
//   - [NWUDPSession.ResolvedEndpoint]: The currently targeted remote endpoint.
//
// # Transferring data
//
//   - [NWUDPSession.MaximumDatagramLength]: The maximum size of a datagram to be written currently.
//
// # Responding to network changes
//
//   - [NWUDPSession.HasBetterPath]: If a session has a better path, new session would use a different interface.
//
// # Getting session properties
//
//   - [NWUDPSession.Endpoint]: The destination endpoint with which this session was created.
//   - [NWUDPSession.CurrentPath]: The current evaluated path for the session’s [resolvedEndpoint](<doc://com.apple.networkextension/documentation/NetworkExtension/NWUDPSession/resolvedEndpoint>) property.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWUDPSession
type NWUDPSession struct {
	objectivec.Object
}

// NWUDPSessionFromID constructs a [NWUDPSession] from an objc.ID.
//
// An object to manage a UDP session to a network endpoint.
func NWUDPSessionFromID(id objc.ID) NWUDPSession {
	return NWUDPSession{objectivec.Object{ID: id}}
}

// NOTE: NWUDPSession adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NWUDPSession] class.
//
// # Monitoring the session state
//
//   - [INWUDPSession.State]: The current state of the UDP session.
//   - [INWUDPSession.Viable]: The viability of a UDP session represents whether or not data can be transferred.
//
// # Selecting remote endpoints
//
//   - [INWUDPSession.ResolvedEndpoint]: The currently targeted remote endpoint.
//
// # Transferring data
//
//   - [INWUDPSession.MaximumDatagramLength]: The maximum size of a datagram to be written currently.
//
// # Responding to network changes
//
//   - [INWUDPSession.HasBetterPath]: If a session has a better path, new session would use a different interface.
//
// # Getting session properties
//
//   - [INWUDPSession.Endpoint]: The destination endpoint with which this session was created.
//   - [INWUDPSession.CurrentPath]: The current evaluated path for the session’s [resolvedEndpoint](<doc://com.apple.networkextension/documentation/NetworkExtension/NWUDPSession/resolvedEndpoint>) property.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWUDPSession
type INWUDPSession interface {
	objectivec.IObject

	// Topic: Monitoring the session state

	// The current state of the UDP session.
	State() NWUDPSessionState
	// The viability of a UDP session represents whether or not data can be transferred.
	Viable() bool

	// Topic: Selecting remote endpoints

	// The currently targeted remote endpoint.
	ResolvedEndpoint() INWEndpoint

	// Topic: Transferring data

	// The maximum size of a datagram to be written currently.
	MaximumDatagramLength() uint

	// Topic: Responding to network changes

	// If a session has a better path, new session would use a different interface.
	HasBetterPath() bool

	// Topic: Getting session properties

	// The destination endpoint with which this session was created.
	Endpoint() INWEndpoint
	// The current evaluated path for the session’s [resolvedEndpoint](<doc://com.apple.networkextension/documentation/NetworkExtension/NWUDPSession/resolvedEndpoint>) property.
	CurrentPath() INWPath
}

// Init initializes the instance.
func (n NWUDPSession) Init() NWUDPSession {
	rv := objc.Send[NWUDPSession](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NWUDPSession) Autorelease() NWUDPSession {
	rv := objc.Send[NWUDPSession](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWUDPSession creates a new NWUDPSession instance.
func NewNWUDPSession() NWUDPSession {
	class := getNWUDPSessionClass()
	rv := objc.Send[NWUDPSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// This convenience initializer can be used to create a new session based on
// the original session’s endpoint and parameters.
//
// # Discussion
//
// The caller should watch the `hasBetterPath` property on an existing
// [NWUDPSession] object. When `hasBetterPath` is true, the caller should call
// “ to create a new session, then start transferring data on the new session
// as soon as possible to reduce power and and avoid expensive networks. When
// the new session is ready, the application can start using the new session
// and tear down the original one.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWUDPSession/init(upgradeFor:)
func NewNWUDPSessionWithUpgradeForSession(session INWUDPSession) NWUDPSession {
	instance := getNWUDPSessionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUpgradeForSession:"), session)
	return NWUDPSessionFromID(rv)
}

// The current state of the UDP session.
//
// # Discussion
//
// Use Key-Value Observing (KVO) to monitor the state. If the state is
// [NWUDPSessionStateReady], then the connection is eligible for reading and
// writing. The state will be [NWUDPSessionStateFailed] if the endpoint could
// not be resolved, or all endpoints have been rejected. For information about
// KVO, see [Key-Value Observing Programming Guide].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWUDPSession/state
//
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
func (n NWUDPSession) State() NWUDPSessionState {
	rv := objc.Send[NWUDPSessionState](n.ID, objc.Sel("state"))
	return NWUDPSessionState(rv)
}

// The viability of a UDP session represents whether or not data can be
// transferred.
//
// # Discussion
//
// Evaluates to true if the session can read and write data, false otherwise.
// Use Key-Value Observing to watch this property.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWUDPSession/isViable
func (n NWUDPSession) Viable() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isViable"))
	return rv
}

// The currently targeted remote endpoint.
//
// # Discussion
//
// Use Key-Value Observing (KVO) to watch this property.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWUDPSession/resolvedEndpoint
func (n NWUDPSession) ResolvedEndpoint() INWEndpoint {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("resolvedEndpoint"))
	return NWEndpointFromID(objc.ID(rv))
}

// The maximum size of a datagram to be written currently.
//
// # Discussion
//
// If a datagram is written with a longer length than `maximumDatagramLength`,
// the datagram may be fragmented or encounter an error. Note that this value
// is not guaranteed to be the maximum datagram length for end-to-end
// communication across the network. Use Key-Value Observing to watch this
// property.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWUDPSession/maximumDatagramLength
func (n NWUDPSession) MaximumDatagramLength() uint {
	rv := objc.Send[uint](n.ID, objc.Sel("maximumDatagramLength"))
	return rv
}

// If a session has a better path, new session would use a different
// interface.
//
// # Discussion
//
// Evaluates to true if a new session to the remote endpoint would use a
// different and preferred path. If the current session is not viable, this
// can be used as a hint to try again. If the current session is still viable,
// this can indicate that the system or user has a preference for the newly
// available network path. For example, if the session is established over a
// cellular data network and Wi-Fi is now available, then the session has a
// better path available and this property is set to true. Use the “
// initializer to create a new session with the same parameters as the current
// session. Use Key-Value Observing to watch this property.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWUDPSession/hasBetterPath
func (n NWUDPSession) HasBetterPath() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasBetterPath"))
	return rv
}

// The destination endpoint with which this session was created.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWUDPSession/endpoint
func (n NWUDPSession) Endpoint() INWEndpoint {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("endpoint"))
	return NWEndpointFromID(objc.ID(rv))
}

// The current evaluated path for the session’s [ResolvedEndpoint] property.
//
// # Discussion
//
// Use Key-Value Observing (KVO) to watch for changes to this property. For
// information about KVO, see [Key-Value Observing Programming Guide].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWUDPSession/currentPath
//
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
func (n NWUDPSession) CurrentPath() INWPath {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("currentPath"))
	return NWPathFromID(objc.ID(rv))
}
