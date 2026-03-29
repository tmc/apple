// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NEAppProxyUDPFlow] class.
var (
	_NEAppProxyUDPFlowClass     NEAppProxyUDPFlowClass
	_NEAppProxyUDPFlowClassOnce sync.Once
)

func getNEAppProxyUDPFlowClass() NEAppProxyUDPFlowClass {
	_NEAppProxyUDPFlowClassOnce.Do(func() {
		_NEAppProxyUDPFlowClass = NEAppProxyUDPFlowClass{class: objc.GetClass("NEAppProxyUDPFlow")}
	})
	return _NEAppProxyUDPFlowClass
}

// GetNEAppProxyUDPFlowClass returns the class object for NEAppProxyUDPFlow.
func GetNEAppProxyUDPFlowClass() NEAppProxyUDPFlowClass {
	return getNEAppProxyUDPFlowClass()
}

type NEAppProxyUDPFlowClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEAppProxyUDPFlowClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEAppProxyUDPFlowClass) Alloc() NEAppProxyUDPFlow {
	rv := objc.Send[NEAppProxyUDPFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object for reading and writing data to and from a UDP conversation being
// proxied by the provider.
//
// # Overview
// 
// App Proxy Providers receive UDP connections to be proxied in the form of
// [NEAppProxyUDPFlow] objects.
//
// # Getting flow information
//
//   - [NEAppProxyUDPFlow.LocalEndpoint]: An [NWEndpoint](<doc://com.apple.networkextension/documentation/NetworkExtension/NWEndpoint>) object containing information about the local endpoint of the flow.
//
// # Instance Properties
//
//   - [NEAppProxyUDPFlow.LocalFlowEndpoint]
//   - [NEAppProxyUDPFlow.SetLocalFlowEndpoint]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyUDPFlow
type NEAppProxyUDPFlow struct {
	NEAppProxyFlow
}

// NEAppProxyUDPFlowFromID constructs a [NEAppProxyUDPFlow] from an objc.ID.
//
// An object for reading and writing data to and from a UDP conversation being
// proxied by the provider.
func NEAppProxyUDPFlowFromID(id objc.ID) NEAppProxyUDPFlow {
	return NEAppProxyUDPFlow{NEAppProxyFlow: NEAppProxyFlowFromID(id)}
}
// NOTE: NEAppProxyUDPFlow adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEAppProxyUDPFlow] class.
//
// # Getting flow information
//
//   - [INEAppProxyUDPFlow.LocalEndpoint]: An [NWEndpoint](<doc://com.apple.networkextension/documentation/NetworkExtension/NWEndpoint>) object containing information about the local endpoint of the flow.
//
// # Instance Properties
//
//   - [INEAppProxyUDPFlow.LocalFlowEndpoint]
//   - [INEAppProxyUDPFlow.SetLocalFlowEndpoint]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyUDPFlow
type INEAppProxyUDPFlow interface {
	INEAppProxyFlow

	// Topic: Getting flow information

	// An [NWEndpoint](<doc://com.apple.networkextension/documentation/NetworkExtension/NWEndpoint>) object containing information about the local endpoint of the flow.
	LocalEndpoint() INWEndpoint

	// Topic: Instance Properties

	LocalFlowEndpoint() INWEndpoint
	SetLocalFlowEndpoint(value INWEndpoint)
}

// Init initializes the instance.
func (a NEAppProxyUDPFlow) Init() NEAppProxyUDPFlow {
	rv := objc.Send[NEAppProxyUDPFlow](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NEAppProxyUDPFlow) Autorelease() NEAppProxyUDPFlow {
	rv := objc.Send[NEAppProxyUDPFlow](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppProxyUDPFlow creates a new NEAppProxyUDPFlow instance.
func NewNEAppProxyUDPFlow() NEAppProxyUDPFlow {
	class := getNEAppProxyUDPFlowClass()
	rv := objc.Send[NEAppProxyUDPFlow](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An [NWEndpoint] object containing information about the local endpoint of
// the flow.
//
// # Discussion
// 
// This property may be nil if the corresponding UDP socket was not bound to a
// port by the application and the App Proxy Provider did not set a local
// endpoint in [OpenWithLocalEndpointCompletionHandler].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyUDPFlow/localEndpoint
func (a NEAppProxyUDPFlow) LocalEndpoint() INWEndpoint {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("localEndpoint"))
	return NWEndpointFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/networkextension/neappproxyudpflow/localflowendpoint-7ukb6
func (a NEAppProxyUDPFlow) LocalFlowEndpoint() INWEndpoint {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("localFlowEndpoint"))
	return NWEndpointFromID(objc.ID(rv))
}
func (a NEAppProxyUDPFlow) SetLocalFlowEndpoint(value INWEndpoint) {
	objc.Send[struct{}](a.ID, objc.Sel("setLocalFlowEndpoint:"), value)
}

