// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEAppProxyProvider] class.
var (
	_NEAppProxyProviderClass     NEAppProxyProviderClass
	_NEAppProxyProviderClassOnce sync.Once
)

func getNEAppProxyProviderClass() NEAppProxyProviderClass {
	_NEAppProxyProviderClassOnce.Do(func() {
		_NEAppProxyProviderClass = NEAppProxyProviderClass{class: objc.GetClass("NEAppProxyProvider")}
	})
	return _NEAppProxyProviderClass
}

// GetNEAppProxyProviderClass returns the class object for NEAppProxyProvider.
func GetNEAppProxyProviderClass() NEAppProxyProviderClass {
	return getNEAppProxyProviderClass()
}

type NEAppProxyProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEAppProxyProviderClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEAppProxyProviderClass) Alloc() NEAppProxyProvider {
	rv := objc.Send[NEAppProxyProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The principal class for an app proxy provider app extension.
//
// # Overview
//
// The [NEAppProxyProvider] class provides access to flows of network data in
// the form of [NEAppProxyFlow] objects. Each [NEAppProxyFlow] object
// corresponds to a socket opened by an app that matches the app rules
// specified in the current App Proxy configuration. Your App Proxy Provider
// acts as a transparent network proxy for the flows of network data that it
// receives.
//
// # DNS Handling
//
// In addition to flows of raw network data from applications, the App Proxy
// Provider also receives flows of DNS queries in the form of
// [NEAppProxyUDPFlow] objects. DNS query flows are received only for
// applications that use low-level DNS resolution APIs such as
// [DNSServiceGetAddrInfo(_:_:_:_:_:_:_:)](). The App Proxy Provider can
// specify the DNS resolver configuration that will be used by these
// applications using the [SetTunnelNetworkSettingsCompletionHandler] method.
//
// Applications that use higher-level networking APIs such as [URLSession] and
// [NSURLConnection] do not generate DNS queries. Instead the destination
// hostname for the connection is included in the endpoint information of the
// [NEAppProxyFlow] object.
//
// # Creating an App Proxy Provider Extension
//
// App Proxy Providers run as App Extensions for the
// `com.AppleXCUIElementTypeNetworkextensionXCUIElementTypeApp()-proxy`
// extension point.
//
// To create a App Proxy Provider extension, first create a new App Extension
// target in your project.
//
// For an example of an Xcode build target for this app extension, see the
// [SimpleTunnel: Customized Networking Using the NetworkExtension Framework]
// sample code project.
//
// Once you have a App Proxy Provider extension target, create a sub-class of
// [NEAppProxyProvider]. Then, set the [NSExtensionPrincipalClass] key in the
// the extension’s `Info.Plist()` to the name of your sub-class.
//
// If it is not already done, set the [NSExtensionPointIdentifier] key in the
// extension’s `Info.Plist()` to
// `com.AppleXCUIElementTypeNetworkextensionXCUIElementTypeApp()-proxy`.
//
// Here is an example of the NSExtension dictionary in a App Proxy Provider
// extension’s `Info.Plist()`:
//
// Finally, add your App Proxy Provider extension target to your app’s Embed
// App Extensions build phase.
//
// # Subclassing Notes
//
// In order to create a App Proxy Provider extension, you must create a
// subclass of [NEAppProxyProvider] and override the methods listed below.
//
// # Methods to Override
//
// - [NEAppProxyProvider.StartProxyWithOptionsCompletionHandler] -
// [NEAppProxyProvider.StopProxyWithReasonCompletionHandler] - [NEAppProxyProvider.HandleNewFlow]
//
// # Managing the app proxy life cycle
//
//   - [NEAppProxyProvider.StartProxyWithOptionsCompletionHandler]: Start the network proxy.
//   - [NEAppProxyProvider.StopProxyWithReasonCompletionHandler]: Stop the network proxy.
//   - [NEAppProxyProvider.CancelProxyWithError]: Stop the network proxy from the App Proxy Provider.
//
// # Handling proxied flows
//
//   - [NEAppProxyProvider.HandleNewFlow]: Handle a new flow of network data.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyProvider
//
// [DNSServiceGetAddrInfo(_:_:_:_:_:_:_:)]: https://developer.apple.com/documentation/dnssd/DNSServiceGetAddrInfo(_:_:_:_:_:_:_:)
// [NSURLConnection]: https://developer.apple.com/documentation/Foundation/NSURLConnection
// [SimpleTunnel: Customized Networking Using the NetworkExtension Framework]: https://developer.apple.com/library/archive/samplecode/SimpleTunnel/Introduction/Intro.html#//apple_ref/doc/uid/TP40016140
// [URLSession]: https://developer.apple.com/documentation/Foundation/URLSession
type NEAppProxyProvider struct {
	NETunnelProvider
}

// NEAppProxyProviderFromID constructs a [NEAppProxyProvider] from an objc.ID.
//
// The principal class for an app proxy provider app extension.
func NEAppProxyProviderFromID(id objc.ID) NEAppProxyProvider {
	return NEAppProxyProvider{NETunnelProvider: NETunnelProviderFromID(id)}
}

// NOTE: NEAppProxyProvider adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEAppProxyProvider] class.
//
// # Managing the app proxy life cycle
//
//   - [INEAppProxyProvider.StartProxyWithOptionsCompletionHandler]: Start the network proxy.
//   - [INEAppProxyProvider.StopProxyWithReasonCompletionHandler]: Stop the network proxy.
//   - [INEAppProxyProvider.CancelProxyWithError]: Stop the network proxy from the App Proxy Provider.
//
// # Handling proxied flows
//
//   - [INEAppProxyProvider.HandleNewFlow]: Handle a new flow of network data.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyProvider
type INEAppProxyProvider interface {
	INETunnelProvider

	// Topic: Managing the app proxy life cycle

	// Start the network proxy.
	StartProxyWithOptionsCompletionHandler(options foundation.INSDictionary, completionHandler ErrorHandler)
	// Stop the network proxy.
	StopProxyWithReasonCompletionHandler(reason NEProviderStopReason, completionHandler VoidHandler)
	// Stop the network proxy from the App Proxy Provider.
	CancelProxyWithError(error_ foundation.INSError)

	// Topic: Handling proxied flows

	// Handle a new flow of network data.
	HandleNewFlow(flow INEAppProxyFlow) bool

	HandleNewUDPFlowInitialRemoteFlowEndpoint(flow INEAppProxyUDPFlow, remoteEndpoint objectivec.IObject) bool
}

// Init initializes the instance.
func (a NEAppProxyProvider) Init() NEAppProxyProvider {
	rv := objc.Send[NEAppProxyProvider](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NEAppProxyProvider) Autorelease() NEAppProxyProvider {
	rv := objc.Send[NEAppProxyProvider](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppProxyProvider creates a new NEAppProxyProvider instance.
func NewNEAppProxyProvider() NEAppProxyProvider {
	class := getNEAppProxyProviderClass()
	rv := objc.Send[NEAppProxyProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Start the network proxy.
//
// options: A dictionary passed by the app that requested that the proxy be started. If
// the starting app did not specify a dictionary of options then this
// parameter will be nil. If the proxy was started via Connect On Demand, then
// this parameter will be nil.
//
// completionHandler: A block that must be executed when the proxy is fully established, or when
// the proxy cannot be started due to an error. If the proxy was successfully
// established, then the error parameter must be set to nil. If an error
// occurred, the error parameter passed to this block must be set to a non-nil
// [NSError] object.
//
// # Discussion
//
// This method is called by the system to start the network proxy.
//
// [NEAppProxyProvider] subclasses must override this method.
//
// When the App Proxy Provider executes the `completionHandler` block with a
// nil error parameter, it signals to the system that it is ready to begin
// handling network data.
//
// The domain and code of the [NSError] object passed to the
// `completionHandler` block are defined by the App Proxy Provider.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyProvider/startProxy(options:completionHandler:)
//
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
func (a NEAppProxyProvider) StartProxyWithOptionsCompletionHandler(options foundation.INSDictionary, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("startProxyWithOptions:completionHandler:"), options, _block1)
}

// Stop the network proxy.
//
// reason: A [NEProviderStopReason] code indicating why the proxy is being stopped.
// For a list of possible codes, see [NEProvider].
//
// completionHandler: A block that must be executed when the proxy is fully stopped.
//
// # Discussion
//
// This method is called by the system to stop the network proxy.
//
// [NEAppProxyProvider] subclasses must override this method.
//
// Do not use this method to stop the proxy from the App Proxy Provider. Use
// “ instead.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyProvider/stopProxy(with:completionHandler:)
func (a NEAppProxyProvider) StopProxyWithReasonCompletionHandler(reason NEProviderStopReason, completionHandler VoidHandler) {
	_block1, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("stopProxyWithReason:completionHandler:"), reason, _block1)
}

// Stop the network proxy from the App Proxy Provider.
//
// error: An [NSError] object containing the error that caused the proxy to be
// stopped. The domain and code of this [NSError] object is defined by the
// caller.
//
// # Discussion
//
// The App Proxy Provider should call this method when an unrecoverable error
// occurs that makes the proxy no longer viable.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyProvider/cancelProxyWithError(_:)
//
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
func (a NEAppProxyProvider) CancelProxyWithError(error_ foundation.INSError) {
	objc.Send[objc.ID](a.ID, objc.Sel("cancelProxyWithError:"), error_)
}

// Handle a new flow of network data.
//
// flow: The new [NEAppProxyFlow] object. If the App Proxy Provider decides to proxy
// the flow, it should create a reference to the flow in its data structures.
//
// # Return Value
//
// Return true to indicate that the App Proxy Provider will handle the flow.
// Return false to indicate that the flow should be closed.
//
// # Discussion
//
// This method is called by the system whenever an app which matches the
// current App Proxy configuration’s app rules opens a new network
// connection.
//
// [NEAppProxyProvider] subclasses must override this method.
//
// New flows are initially in an unopened state. The App Proxy Provider should
// take whatever steps are necessary to ready itself to handle the flow data
// and then open the flow.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyProvider/handleNewFlow(_:)
func (a NEAppProxyProvider) HandleNewFlow(flow INEAppProxyFlow) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("handleNewFlow:"), flow)
	return rv
}

// remoteEndpoint is a [network.nw_endpoint_t].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyProvider/handleNewUDPFlow:initialRemoteFlowEndpoint:
// remoteEndpoint is a [network.nw_endpoint_t].
func (a NEAppProxyProvider) HandleNewUDPFlowInitialRemoteFlowEndpoint(flow INEAppProxyUDPFlow, remoteEndpoint objectivec.IObject) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("handleNewUDPFlow:initialRemoteFlowEndpoint:"), flow, remoteEndpoint)
	return rv
}

// StartProxyWithOptions is a synchronous wrapper around [NEAppProxyProvider.StartProxyWithOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a NEAppProxyProvider) StartProxyWithOptions(ctx context.Context, options foundation.INSDictionary) error {
	done := make(chan error, 1)
	a.StartProxyWithOptionsCompletionHandler(options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// StopProxyWithReason is a synchronous wrapper around [NEAppProxyProvider.StopProxyWithReasonCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a NEAppProxyProvider) StopProxyWithReason(ctx context.Context, reason NEProviderStopReason) error {
	done := make(chan struct{}, 1)
	a.StopProxyWithReasonCompletionHandler(reason, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
