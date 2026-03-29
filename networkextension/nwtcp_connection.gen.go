// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NWTCPConnection] class.
var (
	_NWTCPConnectionClass     NWTCPConnectionClass
	_NWTCPConnectionClassOnce sync.Once
)

func getNWTCPConnectionClass() NWTCPConnectionClass {
	_NWTCPConnectionClassOnce.Do(func() {
		_NWTCPConnectionClass = NWTCPConnectionClass{class: objc.GetClass("NWTCPConnection")}
	})
	return _NWTCPConnectionClass
}

// GetNWTCPConnectionClass returns the class object for NWTCPConnection.
func GetNWTCPConnectionClass() NWTCPConnectionClass {
	return getNWTCPConnectionClass()
}

type NWTCPConnectionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NWTCPConnectionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NWTCPConnectionClass) Alloc() NWTCPConnection {
	rv := objc.Send[NWTCPConnection](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object to manage a TCP connection, with or without TLS.
//
// # Monitoring the connection status
//
//   - [NWTCPConnection.State]: The status of the connection.
//   - [NWTCPConnection.Viable]: The viability of a TCP connection indicates whether or not data can be transferred.
//   - [NWTCPConnection.Error]: The connection-wide error property.
//
// # Responding to network changes
//
//   - [NWTCPConnection.HasBetterPath]: If a connection has a better path, new connections would use a different interface.
//
// # Getting connection properties
//
//   - [NWTCPConnection.Endpoint]: The destination endpoint with which this connection was created.
//   - [NWTCPConnection.LocalAddress]: The IP address endpoint from which the connection was established.
//   - [NWTCPConnection.RemoteAddress]: The IP address endpoint to which the connection was established.
//   - [NWTCPConnection.ConnectedPath]: The network path over which the connection was established.
//   - [NWTCPConnection.TxtRecord]: The TXT record associated with a connected Bonjour service endpoint.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection
type NWTCPConnection struct {
	objectivec.Object
}

// NWTCPConnectionFromID constructs a [NWTCPConnection] from an objc.ID.
//
// An object to manage a TCP connection, with or without TLS.
func NWTCPConnectionFromID(id objc.ID) NWTCPConnection {
	return NWTCPConnection{objectivec.Object{ID: id}}
}
// NOTE: NWTCPConnection adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NWTCPConnection] class.
//
// # Monitoring the connection status
//
//   - [INWTCPConnection.State]: The status of the connection.
//   - [INWTCPConnection.Viable]: The viability of a TCP connection indicates whether or not data can be transferred.
//   - [INWTCPConnection.Error]: The connection-wide error property.
//
// # Responding to network changes
//
//   - [INWTCPConnection.HasBetterPath]: If a connection has a better path, new connections would use a different interface.
//
// # Getting connection properties
//
//   - [INWTCPConnection.Endpoint]: The destination endpoint with which this connection was created.
//   - [INWTCPConnection.LocalAddress]: The IP address endpoint from which the connection was established.
//   - [INWTCPConnection.RemoteAddress]: The IP address endpoint to which the connection was established.
//   - [INWTCPConnection.ConnectedPath]: The network path over which the connection was established.
//   - [INWTCPConnection.TxtRecord]: The TXT record associated with a connected Bonjour service endpoint.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection
type INWTCPConnection interface {
	objectivec.IObject

	// Topic: Monitoring the connection status

	// The status of the connection.
	State() NWTCPConnectionState
	// The viability of a TCP connection indicates whether or not data can be transferred.
	Viable() bool
	// The connection-wide error property.
	Error() foundation.INSError

	// Topic: Responding to network changes

	// If a connection has a better path, new connections would use a different interface.
	HasBetterPath() bool

	// Topic: Getting connection properties

	// The destination endpoint with which this connection was created.
	Endpoint() INWEndpoint
	// The IP address endpoint from which the connection was established.
	LocalAddress() INWEndpoint
	// The IP address endpoint to which the connection was established.
	RemoteAddress() INWEndpoint
	// The network path over which the connection was established.
	ConnectedPath() INWPath
	// The TXT record associated with a connected Bonjour service endpoint.
	TxtRecord() foundation.INSData
}

// Init initializes the instance.
func (n NWTCPConnection) Init() NWTCPConnection {
	rv := objc.Send[NWTCPConnection](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NWTCPConnection) Autorelease() NWTCPConnection {
	rv := objc.Send[NWTCPConnection](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWTCPConnection creates a new NWTCPConnection instance.
func NewNWTCPConnection() NWTCPConnection {
	class := getNWTCPConnectionClass()
	rv := objc.Send[NWTCPConnection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// This convenience initializer can be used to create a new connection that
// will only be connected if there exists a better path (as determined by the
// system) to the remote endpoint of the original connection.
//
// # Discussion
// 
// An upgraded connection will be initialized using the same remote endpoint
// and set of parameters from the original connection. If the original
// connection becomes disconnected or cancelled, the new upgrade connection
// will automatically be considered better.
// 
// The caller should create an [NWTCPConnection] and watch for the
// `hasBetterPath` property. When this property is [true], the caller should
// attempt to create a new upgrade connection, with the goal to start
// transferring data on the new connection path as soon as possible to reduce
// power and avoid expensive networks. When the new connection is successfully
// connected the caller can start using the new connection and cancel the
// original one.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection/init(upgradeFor:)
func NewNWTCPConnectionWithUpgradeForConnection(connection INWTCPConnection) NWTCPConnection {
	instance := getNWTCPConnectionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUpgradeForConnection:"), connection)
	return NWTCPConnectionFromID(rv)
}

// The status of the connection.
//
// # Discussion
// 
// Use Key-Value Observing (KVO) to monitor the state. Many methods, such as
// reading and writing on the connection, are only valid when the state is
// [NWTCPConnectionStateConnected]. For information about KVO, see [Key-Value
// Observing Programming Guide].
//
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection/state
func (n NWTCPConnection) State() NWTCPConnectionState {
	rv := objc.Send[NWTCPConnectionState](n.ID, objc.Sel("state"))
	return NWTCPConnectionState(rv)
}
// The viability of a TCP connection indicates whether or not data can be
// transferred.
//
// # Discussion
// 
// Evaluates to [true] if the connection can read and write data, [false]
// otherwise. Use Key-Value Observing to watch this property.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection/isViable
func (n NWTCPConnection) Viable() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isViable"))
	return rv
}
// The connection-wide error property.
//
// # Discussion
// 
// Indicates any fatal error that occurred while processing the connection or
// performing data reading or writing. Use Key-Value Observing to watch this
// property.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection/error
func (n NWTCPConnection) Error() foundation.INSError {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
// If a connection has a better path, new connections would use a different
// interface.
//
// # Discussion
// 
// Evaluates to [true] if a new connection attempt to the remote endpoint
// would use a different and preferred path. If the current connection is not
// viable, this can be used as a hint to try again. If the current connection
// is still viable, this can indicate that the system or user has a preference
// for the newly available network path. For example, if the connection is
// established over a cellular data network and Wi-Fi is now available, then
// the connection has a better path available and this property is set to
// [true]. Use the `` initializer to create a new connection with the same
// parameters as the current connection. Use Key-Value Observing to watch this
// property.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection/hasBetterPath
func (n NWTCPConnection) HasBetterPath() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasBetterPath"))
	return rv
}
// The destination endpoint with which this connection was created.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection/endpoint
func (n NWTCPConnection) Endpoint() INWEndpoint {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("endpoint"))
	return NWEndpointFromID(objc.ID(rv))
}
// The IP address endpoint from which the connection was established.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection/localAddress
func (n NWTCPConnection) LocalAddress() INWEndpoint {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("localAddress"))
	return NWEndpointFromID(objc.ID(rv))
}
// The IP address endpoint to which the connection was established.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection/remoteAddress
func (n NWTCPConnection) RemoteAddress() INWEndpoint {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("remoteAddress"))
	return NWEndpointFromID(objc.ID(rv))
}
// The network path over which the connection was established.
//
// # Discussion
// 
// The caller can query additional properties from the [NWPath] object for
// more information. Note that this contains a snapshot of information at the
// time of connection establishment for this connection only. As a result,
// some underlying properties might change in time and might not reflect the
// path for other connections that might be established at different times.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection/connectedPath
func (n NWTCPConnection) ConnectedPath() INWPath {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("connectedPath"))
	return NWPathFromID(objc.ID(rv))
}
// The TXT record associated with a connected Bonjour service endpoint.
//
// # Discussion
// 
// When the connection is connected to a Bonjour service endpoint, the TXT
// record associated with the Bonjour service is available via this property.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection/txtRecord
func (n NWTCPConnection) TxtRecord() foundation.INSData {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("txtRecord"))
	return foundation.NSDataFromID(objc.ID(rv))
}

