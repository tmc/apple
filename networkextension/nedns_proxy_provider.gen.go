// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NEDNSProxyProvider] class.
var (
	_NEDNSProxyProviderClass     NEDNSProxyProviderClass
	_NEDNSProxyProviderClassOnce sync.Once
)

func getNEDNSProxyProviderClass() NEDNSProxyProviderClass {
	_NEDNSProxyProviderClassOnce.Do(func() {
		_NEDNSProxyProviderClass = NEDNSProxyProviderClass{class: objc.GetClass("NEDNSProxyProvider")}
	})
	return _NEDNSProxyProviderClass
}

// GetNEDNSProxyProviderClass returns the class object for NEDNSProxyProvider.
func GetNEDNSProxyProviderClass() NEDNSProxyProviderClass {
	return getNEDNSProxyProviderClass()
}

type NEDNSProxyProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEDNSProxyProviderClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEDNSProxyProviderClass) Alloc() NEDNSProxyProvider {
	rv := objc.Send[NEDNSProxyProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The principal class for a DNS proxy provider app extension.
//
// # Overview
// 
// A DNS proxy allows your app to intercept all DNS traffic generated on a
// device. You can use this capability to provide services like DNS traffic
// encryption, typically by redirecting DNS traffic to your own server. You
// usually do this in the context of managed devices, such as those owned by a
// school or an enterprise.
// 
// You create a DNS proxy as an app extension based on a custom subclass of
// the [NEDNSProxyProvider] class. Once active, the proxy receives access to
// flows of DNS traffic in the form of [NEAppProxyFlow] instances. Each flow
// corresponds to a socket opened by an app to UDP port 53 or TCP port 53.
// Your DNS proxy provider acts as a transparent DNS proxy for the flows of
// network data that it receives.
// 
// When you subclass [NEDNSProxyProvider], you must provide implementations
// for the following methods:
// 
// - [NEDNSProxyProvider.StartProxyWithOptionsCompletionHandler] -
// [NEDNSProxyProvider.StopProxyWithReasonCompletionHandler] - [NEDNSProxyProvider.HandleNewFlow]
//
// # Managing the DNS proxy life cycle
//
//   - [NEDNSProxyProvider.StartProxyWithOptionsCompletionHandler]: Starts the DNS proxy.
//   - [NEDNSProxyProvider.StopProxyWithReasonCompletionHandler]: Stops the DNS proxy.
//   - [NEDNSProxyProvider.CancelProxyWithError]: Cancels the DNS proxy.
//
// # Handling proxied DNS flow
//
//   - [NEDNSProxyProvider.HandleNewFlow]: Handles a new flow of DNS traffic.
//
// # Getting system DNS settings
//
//   - [NEDNSProxyProvider.SystemDNSSettings]: The current system DNS settings.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProvider
type NEDNSProxyProvider struct {
	NEProvider
}

// NEDNSProxyProviderFromID constructs a [NEDNSProxyProvider] from an objc.ID.
//
// The principal class for a DNS proxy provider app extension.
func NEDNSProxyProviderFromID(id objc.ID) NEDNSProxyProvider {
	return NEDNSProxyProvider{NEProvider: NEProviderFromID(id)}
}
// NOTE: NEDNSProxyProvider adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEDNSProxyProvider] class.
//
// # Managing the DNS proxy life cycle
//
//   - [INEDNSProxyProvider.StartProxyWithOptionsCompletionHandler]: Starts the DNS proxy.
//   - [INEDNSProxyProvider.StopProxyWithReasonCompletionHandler]: Stops the DNS proxy.
//   - [INEDNSProxyProvider.CancelProxyWithError]: Cancels the DNS proxy.
//
// # Handling proxied DNS flow
//
//   - [INEDNSProxyProvider.HandleNewFlow]: Handles a new flow of DNS traffic.
//
// # Getting system DNS settings
//
//   - [INEDNSProxyProvider.SystemDNSSettings]: The current system DNS settings.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProvider
type INEDNSProxyProvider interface {
	INEProvider

	// Topic: Managing the DNS proxy life cycle

	// Starts the DNS proxy.
	StartProxyWithOptionsCompletionHandler(options foundation.INSDictionary, completionHandler ErrorHandler)
	// Stops the DNS proxy.
	StopProxyWithReasonCompletionHandler(reason NEProviderStopReason, completionHandler VoidHandler)
	// Cancels the DNS proxy.
	CancelProxyWithError(error_ foundation.INSError)

	// Topic: Handling proxied DNS flow

	// Handles a new flow of DNS traffic.
	HandleNewFlow(flow INEAppProxyFlow) bool

	// Topic: Getting system DNS settings

	// The current system DNS settings.
	SystemDNSSettings() foundation.INSSet
}

// Init initializes the instance.
func (d NEDNSProxyProvider) Init() NEDNSProxyProvider {
	rv := objc.Send[NEDNSProxyProvider](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NEDNSProxyProvider) Autorelease() NEDNSProxyProvider {
	rv := objc.Send[NEDNSProxyProvider](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSProxyProvider creates a new NEDNSProxyProvider instance.
func NewNEDNSProxyProvider() NEDNSProxyProvider {
	class := getNEDNSProxyProviderClass()
	rv := objc.Send[NEDNSProxyProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Starts the DNS proxy.
//
// options: A dictionary that you define as part of a device configuration profile. You
// can also modify the contents of this dictionary from your app using the
// shared instance of [NEDNSProxyManager]. The dictionary appears as the
// [ProviderConfiguration] component of the manager’s [ProviderProtocol]
// property.
//
// completionHandler: A block that you must execute when the proxy is fully established, or when
// the proxy cannot be started due to an error. If the proxy is successfully
// established, the error parameter should be set to `nil`. Otherwise, the
// error parameter passed to this block indicates the reason for failure.
//
// # Discussion
// 
// Subclasses of [NEDNSProxyProvider] must override this method to perform any
// necessary steps to ready the proxy for handling flows of network data.
// 
// The framework calls this method when a new proxy instance is created. You
// indicate that setup is complete by calling the completion handler with a
// `nil` error parameter, or that setup failed by calling the completion
// handler with an error instance. You define the error domain and code.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProvider/startProxy(options:completionHandler:)
func (d NEDNSProxyProvider) StartProxyWithOptionsCompletionHandler(options foundation.INSDictionary, completionHandler ErrorHandler) {
_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](d.ID, objc.Sel("startProxyWithOptions:completionHandler:"), options, _block1)
}
// Stops the DNS proxy.
//
// reason: A code indicating why the proxy is being stopped.
//
// completionHandler: A block that must be called when the proxy is completely stopped.
//
// # Discussion
// 
// Subclasses of [NEDNSProxyProvider] must override this method to perform
// whatever steps are necessary to stop the proxy.
// 
// The system calls this method to stop the proxy. You indicate that the proxy
// is fully stopped by calling the completion handler.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProvider/stopProxy(with:completionHandler:)
func (d NEDNSProxyProvider) StopProxyWithReasonCompletionHandler(reason NEProviderStopReason, completionHandler VoidHandler) {
_block1, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](d.ID, objc.Sel("stopProxyWithReason:completionHandler:"), reason, _block1)
}
// Cancels the DNS proxy.
//
// error: An error instance containing details about the problem that the proxy
// provider implementation encountered.
//
// # Discussion
// 
// Call this method from within the proxy provider when you need to stop the
// proxy due to a network error that renders the proxy no longer viable.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProvider/cancelProxyWithError(_:)
func (d NEDNSProxyProvider) CancelProxyWithError(error_ foundation.INSError) {
	objc.Send[objc.ID](d.ID, objc.Sel("cancelProxyWithError:"), error_)
}
// Handles a new flow of DNS traffic.
//
// flow: The flow representing the DNS traffic that the proxy should handle.
//
// # Return Value
// 
// A Boolean value set to [true] if the proxy implementation decides to handle
// the flow, or [false] if it instead decides to terminate the flow.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The system calls this method to deliver a new network data flow to the
// proxy provider implementation. Subclasses must override this method to
// perform whatever steps are necessary to ready the proxy to receive data
// from the flow.
// 
// The proxy provider indicates that the proxy is ready to handle flow data by
// calling the flow’s [OpenWithLocalEndpointCompletionHandler] method.
// 
// If the proxy implementation decides to handle the flow, it’s responsible
// for retaining a reference to the flow instance.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProvider/handleNewFlow(_:)
func (d NEDNSProxyProvider) HandleNewFlow(flow INEAppProxyFlow) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("handleNewFlow:"), flow)
	return rv
}

// The current system DNS settings.
//
// # Discussion
// 
// You can use key-value observing to watch for changes on this parameter.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProvider/systemDNSSettings
func (d NEDNSProxyProvider) SystemDNSSettings() foundation.INSSet {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("systemDNSSettings"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// StartProxyWithOptions is a synchronous wrapper around [NEDNSProxyProvider.StartProxyWithOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NEDNSProxyProvider) StartProxyWithOptions(ctx context.Context, options foundation.INSDictionary) error {
	done := make(chan error, 1)
	d.StartProxyWithOptionsCompletionHandler(options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// StopProxyWithReason is a synchronous wrapper around [NEDNSProxyProvider.StopProxyWithReasonCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NEDNSProxyProvider) StopProxyWithReason(ctx context.Context, reason NEProviderStopReason) error {
	done := make(chan struct{}, 1)
	d.StopProxyWithReasonCompletionHandler(reason, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

