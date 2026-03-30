// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEFilterSocketFlow] class.
var (
	_NEFilterSocketFlowClass     NEFilterSocketFlowClass
	_NEFilterSocketFlowClassOnce sync.Once
)

func getNEFilterSocketFlowClass() NEFilterSocketFlowClass {
	_NEFilterSocketFlowClassOnce.Do(func() {
		_NEFilterSocketFlowClass = NEFilterSocketFlowClass{class: objc.GetClass("NEFilterSocketFlow")}
	})
	return _NEFilterSocketFlowClass
}

// GetNEFilterSocketFlowClass returns the class object for NEFilterSocketFlow.
func GetNEFilterSocketFlowClass() NEFilterSocketFlowClass {
	return getNEFilterSocketFlowClass()
}

type NEFilterSocketFlowClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEFilterSocketFlowClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEFilterSocketFlowClass) Alloc() NEFilterSocketFlow {
	rv := objc.Send[NEFilterSocketFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A flow of network data that the filter examines.
//
// # Getting socket flow properties
//
//   - [NEFilterSocketFlow.RemoteEndpoint]: An object containing details about the socket’s remote endpoint.
//   - [NEFilterSocketFlow.RemoteHostname]: The flow’s remote hostname, if applicable.
//   - [NEFilterSocketFlow.LocalEndpoint]: An object containing details about the socket’s local endpoint.
//   - [NEFilterSocketFlow.SocketFamily]: The protocol family of the socket.
//   - [NEFilterSocketFlow.SocketType]: The type of the socket.
//   - [NEFilterSocketFlow.SocketProtocol]: The protocol of the socket.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow
type NEFilterSocketFlow struct {
	NEFilterFlow
}

// NEFilterSocketFlowFromID constructs a [NEFilterSocketFlow] from an objc.ID.
//
// A flow of network data that the filter examines.
func NEFilterSocketFlowFromID(id objc.ID) NEFilterSocketFlow {
	return NEFilterSocketFlow{NEFilterFlow: NEFilterFlowFromID(id)}
}

// NOTE: NEFilterSocketFlow adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEFilterSocketFlow] class.
//
// # Getting socket flow properties
//
//   - [INEFilterSocketFlow.RemoteEndpoint]: An object containing details about the socket’s remote endpoint.
//   - [INEFilterSocketFlow.RemoteHostname]: The flow’s remote hostname, if applicable.
//   - [INEFilterSocketFlow.LocalEndpoint]: An object containing details about the socket’s local endpoint.
//   - [INEFilterSocketFlow.SocketFamily]: The protocol family of the socket.
//   - [INEFilterSocketFlow.SocketType]: The type of the socket.
//   - [INEFilterSocketFlow.SocketProtocol]: The protocol of the socket.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow
type INEFilterSocketFlow interface {
	INEFilterFlow

	// Topic: Getting socket flow properties

	// An object containing details about the socket’s remote endpoint.
	RemoteEndpoint() INWEndpoint
	// The flow’s remote hostname, if applicable.
	RemoteHostname() string
	// An object containing details about the socket’s local endpoint.
	LocalEndpoint() INWEndpoint
	// The protocol family of the socket.
	SocketFamily() int
	// The type of the socket.
	SocketType() int
	// The protocol of the socket.
	SocketProtocol() int

	LocalFlowEndpoint() objectivec.IObject
	RemoteFlowEndpoint() objectivec.IObject
}

// Init initializes the instance.
func (f NEFilterSocketFlow) Init() NEFilterSocketFlow {
	rv := objc.Send[NEFilterSocketFlow](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NEFilterSocketFlow) Autorelease() NEFilterSocketFlow {
	rv := objc.Send[NEFilterSocketFlow](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterSocketFlow creates a new NEFilterSocketFlow instance.
func NewNEFilterSocketFlow() NEFilterSocketFlow {
	class := getNEFilterSocketFlowClass()
	rv := objc.Send[NEFilterSocketFlow](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An object containing details about the socket’s remote endpoint.
//
// # Discussion
//
// This endpoint object may be `nil` when the system calls your
// [HandleNewFlow] method; if so, receiving network data populates the object.
// In such a case, the filter may still perform filtering, based on its socket
// type, socket family, or socket protocol.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/remoteEndpoint
func (f NEFilterSocketFlow) RemoteEndpoint() INWEndpoint {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("remoteEndpoint"))
	return NWEndpointFromID(objc.ID(rv))
}

// The flow’s remote hostname, if applicable.
//
// # Discussion
//
// This property is only populated for flows originating from create-by-name
// APIs like [URLSession] or [Network].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/remoteHostname
//
// [Network]: https://developer.apple.com/documentation/Network
// [URLSession]: https://developer.apple.com/documentation/Foundation/URLSession
func (f NEFilterSocketFlow) RemoteHostname() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("remoteHostname"))
	return foundation.NSStringFromID(rv).String()
}

// An object containing details about the socket’s local endpoint.
//
// # Discussion
//
// This endpoint object may be `nil` when the system calls your
// [HandleNewFlow] method; if so, receiving network data populates the object.
// In such a case, the filter may still perform filtering, based on its socket
// type, socket family, or socket protocol.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/localEndpoint
func (f NEFilterSocketFlow) LocalEndpoint() INWEndpoint {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("localEndpoint"))
	return NWEndpointFromID(objc.ID(rv))
}

// The protocol family of the socket.
//
// # Discussion
//
// Examples of protocol families include symbols like `PF_INET` and
// `PF_INET6`.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/socketFamily
func (f NEFilterSocketFlow) SocketFamily() int {
	rv := objc.Send[int](f.ID, objc.Sel("socketFamily"))
	return rv
}

// The type of the socket.
//
// # Discussion
//
// Examples of socket types include `SOCK_STREAM` and `SOCK_DGRAM`.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/socketType
func (f NEFilterSocketFlow) SocketType() int {
	rv := objc.Send[int](f.ID, objc.Sel("socketType"))
	return rv
}

// The protocol of the socket.
//
// # Discussion
//
// Examples of protocols include `IPPROTO_TCP` and `IPPROTO_IP`.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/socketProtocol
func (f NEFilterSocketFlow) SocketProtocol() int {
	rv := objc.Send[int](f.ID, objc.Sel("socketProtocol"))
	return rv
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/localFlowEndpoint-4nt54
func (f NEFilterSocketFlow) LocalFlowEndpoint() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("localFlowEndpoint"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/remoteFlowEndpoint-52dxr
func (f NEFilterSocketFlow) RemoteFlowEndpoint() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("remoteFlowEndpoint"))
	return objectivec.Object{ID: rv}
}
