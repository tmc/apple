// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/network"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NEAppProxyTCPFlow] class.
var (
	_NEAppProxyTCPFlowClass     NEAppProxyTCPFlowClass
	_NEAppProxyTCPFlowClassOnce sync.Once
)

func getNEAppProxyTCPFlowClass() NEAppProxyTCPFlowClass {
	_NEAppProxyTCPFlowClassOnce.Do(func() {
		_NEAppProxyTCPFlowClass = NEAppProxyTCPFlowClass{class: objc.GetClass("NEAppProxyTCPFlow")}
	})
	return _NEAppProxyTCPFlowClass
}

// GetNEAppProxyTCPFlowClass returns the class object for NEAppProxyTCPFlow.
func GetNEAppProxyTCPFlowClass() NEAppProxyTCPFlowClass {
	return getNEAppProxyTCPFlowClass()
}

type NEAppProxyTCPFlowClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEAppProxyTCPFlowClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEAppProxyTCPFlowClass) Alloc() NEAppProxyTCPFlow {
	rv := objc.Send[NEAppProxyTCPFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object for reading and writing data to and from a TCP connection being
// proxied by the provider.
//
// # Overview
//
// App Proxy Providers receive TCP connections to be proxied in the form of
// [NEAppProxyTCPFlow] objects.
//
// # Handling flow data
//
//   - [NEAppProxyTCPFlow.WriteDataWithCompletionHandler]: Write data to the flow.
//   - [NEAppProxyTCPFlow.ReadDataWithCompletionHandler]: Read data from the flow.
//
// # Getting flow information
//
//   - [NEAppProxyTCPFlow.RemoteEndpoint]: An [NWEndpoint](<doc://com.apple.networkextension/documentation/NetworkExtension/NWEndpoint>) object containing information about the intended remote endpoint of the flow.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyTCPFlow
type NEAppProxyTCPFlow struct {
	NEAppProxyFlow
}

// NEAppProxyTCPFlowFromID constructs a [NEAppProxyTCPFlow] from an objc.ID.
//
// An object for reading and writing data to and from a TCP connection being
// proxied by the provider.
func NEAppProxyTCPFlowFromID(id objc.ID) NEAppProxyTCPFlow {
	return NEAppProxyTCPFlow{NEAppProxyFlow: NEAppProxyFlowFromID(id)}
}

// NOTE: NEAppProxyTCPFlow adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEAppProxyTCPFlow] class.
//
// # Handling flow data
//
//   - [INEAppProxyTCPFlow.WriteDataWithCompletionHandler]: Write data to the flow.
//   - [INEAppProxyTCPFlow.ReadDataWithCompletionHandler]: Read data from the flow.
//
// # Getting flow information
//
//   - [INEAppProxyTCPFlow.RemoteEndpoint]: An [NWEndpoint](<doc://com.apple.networkextension/documentation/NetworkExtension/NWEndpoint>) object containing information about the intended remote endpoint of the flow.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyTCPFlow
type INEAppProxyTCPFlow interface {
	INEAppProxyFlow

	// Topic: Handling flow data

	// Write data to the flow.
	WriteDataWithCompletionHandler(data foundation.INSData, completionHandler ErrorHandler)
	// Read data from the flow.
	ReadDataWithCompletionHandler(completionHandler DataErrorHandler)

	// Topic: Getting flow information

	// An [NWEndpoint](<doc://com.apple.networkextension/documentation/NetworkExtension/NWEndpoint>) object containing information about the intended remote endpoint of the flow.
	RemoteEndpoint() INWEndpoint

	RemoteFlowEndpoint() network.Nw_endpoint_t
}

// Init initializes the instance.
func (a NEAppProxyTCPFlow) Init() NEAppProxyTCPFlow {
	rv := objc.Send[NEAppProxyTCPFlow](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NEAppProxyTCPFlow) Autorelease() NEAppProxyTCPFlow {
	rv := objc.Send[NEAppProxyTCPFlow](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppProxyTCPFlow creates a new NEAppProxyTCPFlow instance.
func NewNEAppProxyTCPFlow() NEAppProxyTCPFlow {
	class := getNEAppProxyTCPFlowClass()
	rv := objc.Send[NEAppProxyTCPFlow](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Write data to the flow.
//
// data: An [NSData] object containing the data to write.
//
// completionHandler: A block that will be executed by the system on an internal system thread
// when the data is written into the receive buffer of the socket associated
// with the flow. The caller should use this callback as an indication that it
// is possible to write more data to the flow without using up excessive
// buffer memory. If an error occurs while writing the data then a non-nil
// [NSError] object is passed to the block. See [NEAppProxyFlowError] in
// [NEAppProxyFlow] for a list of possible errors.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyTCPFlow/write(_:withCompletionHandler:)
//
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
func (a NEAppProxyTCPFlow) WriteDataWithCompletionHandler(data foundation.INSData, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("writeData:withCompletionHandler:"), data, _block1)
}

// Read data from the flow.
//
// completionHandler: A block that will be executed by the system on an internal system thread
// when some data is read from the flow. The block is passed either the data
// that was read or a non-nil error if an error occurred. See
// [NEAppProxyFlowError] in [NEAppProxyFlow] for a list of possible errors. If
// the data parameter has a length of 0 then no data can be subsequently read
// from the flow.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyTCPFlow/readData(completionHandler:)
func (a NEAppProxyTCPFlow) ReadDataWithCompletionHandler(completionHandler DataErrorHandler) {
	_block0, _ := NewDataErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("readDataWithCompletionHandler:"), _block0)
}

// An [NWEndpoint] object containing information about the intended remote
// endpoint of the flow.
//
// # Discussion
//
// If the flow’s corresponding socket was created using one of the
// high-level networking APIs such as [URLSession] or [NSURLConnection], then
// the hostname property of the `remoteEndpoint` object contains the DNS name
// of the remote host. If the flow’s corresponding socket was created using
// the sockets API directly, then the hostname property of the
// `remoteEndpoint` object contains the IP address of the remote host.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyTCPFlow/remoteEndpoint
//
// [NSURLConnection]: https://developer.apple.com/documentation/Foundation/NSURLConnection
// [URLSession]: https://developer.apple.com/documentation/Foundation/URLSession
func (a NEAppProxyTCPFlow) RemoteEndpoint() INWEndpoint {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("remoteEndpoint"))
	return NWEndpointFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyTCPFlow/remoteFlowEndpoint-9lvob
func (a NEAppProxyTCPFlow) RemoteFlowEndpoint() network.Nw_endpoint_t {
	rv := objc.Send[network.Nw_endpoint_t](a.ID, objc.Sel("remoteFlowEndpoint"))
	return network.Nw_endpoint_t(rv)
}

// WriteData is a synchronous wrapper around [NEAppProxyTCPFlow.WriteDataWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a NEAppProxyTCPFlow) WriteData(ctx context.Context, data foundation.INSData) error {
	done := make(chan error, 1)
	a.WriteDataWithCompletionHandler(data, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ReadData is a synchronous wrapper around [NEAppProxyTCPFlow.ReadDataWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a NEAppProxyTCPFlow) ReadData(ctx context.Context) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	a.ReadDataWithCompletionHandler(func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
