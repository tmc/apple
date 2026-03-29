// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEAppProxyFlow] class.
var (
	_NEAppProxyFlowClass     NEAppProxyFlowClass
	_NEAppProxyFlowClassOnce sync.Once
)

func getNEAppProxyFlowClass() NEAppProxyFlowClass {
	_NEAppProxyFlowClassOnce.Do(func() {
		_NEAppProxyFlowClass = NEAppProxyFlowClass{class: objc.GetClass("NEAppProxyFlow")}
	})
	return _NEAppProxyFlowClass
}

// GetNEAppProxyFlowClass returns the class object for NEAppProxyFlow.
func GetNEAppProxyFlowClass() NEAppProxyFlowClass {
	return getNEAppProxyFlowClass()
}

type NEAppProxyFlowClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEAppProxyFlowClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEAppProxyFlowClass) Alloc() NEAppProxyFlow {
	rv := objc.Send[NEAppProxyFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract base class shared by NEAppProxyTCPFlow and NEAppProxyUDPFlow.
//
// # Overview
// 
// App Proxy Providers receive network connections to be proxied in the form
// of [NEAppProxyFlow] objects, which are passed to the App Proxy Provider via
// the [HandleNewFlow] method.
// 
// [NEAppProxyFlow] objects are initially in an unopened state. Before they
// can be used to transmit network data, they must be opened using the
// [NEAppProxyFlow.OpenWithLocalEndpointCompletionHandler] method. When you are finished with
// a flow, you should call [NEAppProxyFlow.CloseReadWithError] and [NEAppProxyFlow.CloseWriteWithError], and
// then release the [NEAppProxyFlow] object.
//
// # Managing the flow life cycle
//
//   - [NEAppProxyFlow.CloseReadWithError]: Close the flow for further read operations.
//   - [NEAppProxyFlow.CloseWriteWithError]: Close the flow for further write operations.
//
// # Accessing flow information
//
//   - [NEAppProxyFlow.MetaData]: A metadata object containing information about the source app of the flow.
//   - [NEAppProxyFlow.SetMetadata]: Sets the flow’s metadata for use by proxy providers.
//   - [NEAppProxyFlow.IsBound]: A Boolean value that indicates whether the flow has a binding to a specific interface.
//   - [NEAppProxyFlow.NetworkInterface]: The network interface, if any, used by this flow.
//   - [NEAppProxyFlow.SetNetworkInterface]
//   - [NEAppProxyFlow.RemoteHostname]: The remote host name for flows created from a hostname.
//
// # Errors
//
//   - [NEAppProxyFlow.NEAppProxyErrorDomain]: The domain used for app proxy errors.
//
// # Instance Properties
//
//   - [NEAppProxyFlow.Interface]
//   - [NEAppProxyFlow.SetInterface]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow
type NEAppProxyFlow struct {
	objectivec.Object
}

// NEAppProxyFlowFromID constructs a [NEAppProxyFlow] from an objc.ID.
//
// An abstract base class shared by NEAppProxyTCPFlow and NEAppProxyUDPFlow.
func NEAppProxyFlowFromID(id objc.ID) NEAppProxyFlow {
	return NEAppProxyFlow{objectivec.Object{ID: id}}
}
// NOTE: NEAppProxyFlow adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEAppProxyFlow] class.
//
// # Managing the flow life cycle
//
//   - [INEAppProxyFlow.CloseReadWithError]: Close the flow for further read operations.
//   - [INEAppProxyFlow.CloseWriteWithError]: Close the flow for further write operations.
//
// # Accessing flow information
//
//   - [INEAppProxyFlow.MetaData]: A metadata object containing information about the source app of the flow.
//   - [INEAppProxyFlow.SetMetadata]: Sets the flow’s metadata for use by proxy providers.
//   - [INEAppProxyFlow.IsBound]: A Boolean value that indicates whether the flow has a binding to a specific interface.
//   - [INEAppProxyFlow.NetworkInterface]: The network interface, if any, used by this flow.
//   - [INEAppProxyFlow.SetNetworkInterface]
//   - [INEAppProxyFlow.RemoteHostname]: The remote host name for flows created from a hostname.
//
// # Errors
//
//   - [INEAppProxyFlow.NEAppProxyErrorDomain]: The domain used for app proxy errors.
//
// # Instance Properties
//
//   - [INEAppProxyFlow.Interface]
//   - [INEAppProxyFlow.SetInterface]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow
type INEAppProxyFlow interface {
	objectivec.IObject

	// Topic: Managing the flow life cycle

	// Close the flow for further read operations.
	CloseReadWithError(error_ foundation.INSError)
	// Close the flow for further write operations.
	CloseWriteWithError(error_ foundation.INSError)

	// Topic: Accessing flow information

	// A metadata object containing information about the source app of the flow.
	MetaData() INEFlowMetaData
	// Sets the flow’s metadata for use by proxy providers.
	SetMetadata(parameters objectivec.IObject)
	// A Boolean value that indicates whether the flow has a binding to a specific interface.
	IsBound() bool
	// The network interface, if any, used by this flow.
	NetworkInterface() objectivec.IObject
	SetNetworkInterface(value objectivec.IObject)
	// The remote host name for flows created from a hostname.
	RemoteHostname() string

	// Topic: Errors

	// The domain used for app proxy errors.
	NEAppProxyErrorDomain() string

	// Topic: Instance Properties

	Interface() objectivec.IObject
	SetInterface(value objectivec.IObject)
}

// Init initializes the instance.
func (a NEAppProxyFlow) Init() NEAppProxyFlow {
	rv := objc.Send[NEAppProxyFlow](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NEAppProxyFlow) Autorelease() NEAppProxyFlow {
	rv := objc.Send[NEAppProxyFlow](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppProxyFlow creates a new NEAppProxyFlow instance.
func NewNEAppProxyFlow() NEAppProxyFlow {
	class := getNEAppProxyFlowClass()
	rv := objc.Send[NEAppProxyFlow](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Close the flow for further read operations.
//
// error: An [NSError] object indicating to the system the error that led to the
// closure. If the flow is not being closed due to an error, this parameter
// should be set to nil. See [NEAppProxyFlowError] below for a list of
// acceptable error codes.
// //
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/closeReadWithError(_:)
func (a NEAppProxyFlow) CloseReadWithError(error_ foundation.INSError) {
	objc.Send[objc.ID](a.ID, objc.Sel("closeReadWithError:"), error_)
}
// Close the flow for further write operations.
//
// error: An NSError object indicating to the system the error that led to the
// closure. If the flow is not being closed due to an error, this parameter
// should be set to nil. See [NEAppProxyFlowError] below for a list of
// acceptable error codes.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/closeWriteWithError(_:)
func (a NEAppProxyFlow) CloseWriteWithError(error_ foundation.INSError) {
	objc.Send[objc.ID](a.ID, objc.Sel("closeWriteWithError:"), error_)
}
// Sets the flow’s metadata for use by proxy providers.
//
// parameters: A nw_parameters_t object that contains the flow metadata.
//
// parameters is a [network.nw_parameters_t].
//
// # Discussion
// 
// Use an [nw_parameters_t] object to create a connection that transparently
// proxies the flow’s data. This also provides accurate source app
// information to any subsequent [NEAppProxyProvider] instances that
// transparently proxy the flow.
//
// [nw_parameters_t]: https://developer.apple.com/documentation/Network/nw_parameters_t
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/setMetadata(_:)
// parameters is a [network.nw_parameters_t].
func (a NEAppProxyFlow) SetMetadata(parameters objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setMetadata:"), parameters)
}

// A metadata object containing information about the source app of the flow.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/metaData
func (a NEAppProxyFlow) MetaData() INEFlowMetaData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("metaData"))
	return NEFlowMetaDataFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the flow has a binding to a specific
// interface.
//
// # Discussion
// 
// When a binding exists, this value is [true], and the [NetworkInterface]
// property indicates the bound interface. If the flow isn’t bound to an
// interface, this value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/isBound
func (a NEAppProxyFlow) IsBound() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isBound"))
	return rv
}
// The network interface, if any, used by this flow.
//
// # Discussion
// 
// To transport the flow’s data over a different interface, set this
// property to that interface.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/networkInterface
func (a NEAppProxyFlow) NetworkInterface() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("networkInterface"))
	return objectivec.Object{ID: rv}
}
func (a NEAppProxyFlow) SetNetworkInterface(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setNetworkInterface:"), value)
}
// The remote host name for flows created from a hostname.
//
// # Discussion
// 
// The flow populates this property when you create the flow from a
// connect-by-name API such as [URLSession] or the [Network] framework.
//
// [Network]: https://developer.apple.com/documentation/Network
// [URLSession]: https://developer.apple.com/documentation/Foundation/URLSession
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/remoteHostname
func (a NEAppProxyFlow) RemoteHostname() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("remoteHostname"))
	return foundation.NSStringFromID(rv).String()
}
// The domain used for app proxy errors.
//
// See: https://developer.apple.com/documentation/networkextension/neappproxyerrordomain
func (a NEAppProxyFlow) NEAppProxyErrorDomain() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("NEAppProxyErrorDomain"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/networkextension/neappproxyflow/interface
func (a NEAppProxyFlow) Interface() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("interface"))
	return objectivec.Object{ID: rv}
}
func (a NEAppProxyFlow) SetInterface(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setInterface:"), value)
}

