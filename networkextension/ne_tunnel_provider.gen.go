// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NETunnelProvider] class.
var (
	_NETunnelProviderClass     NETunnelProviderClass
	_NETunnelProviderClassOnce sync.Once
)

func getNETunnelProviderClass() NETunnelProviderClass {
	_NETunnelProviderClassOnce.Do(func() {
		_NETunnelProviderClass = NETunnelProviderClass{class: objc.GetClass("NETunnelProvider")}
	})
	return _NETunnelProviderClass
}

// GetNETunnelProviderClass returns the class object for NETunnelProvider.
func GetNETunnelProviderClass() NETunnelProviderClass {
	return getNETunnelProviderClass()
}

type NETunnelProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NETunnelProviderClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NETunnelProviderClass) Alloc() NETunnelProvider {
	rv := objc.Send[NETunnelProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract base class shared by NEPacketTunnelProvider and
// NEAppProxyProvider.
//
// # Overview
//
// Each [NETunnelProvider] instance corresponds to a single tunneling session,
// with a single associated configuration.
//
// # Subclassing Notes
//
// The [NETunnelProvider] class should not be subclassed directly. Instead,
// you should create subclasses of [NETunnelProvider] subclasses.
//
// # Methods to Override
//
// - [NETunnelProvider.HandleAppMessageCompletionHandler]
//
// # Getting the tunnel configuration
//
//   - [NETunnelProvider.ProtocolConfiguration]: The configuration of the current tunneling session.
//   - [NETunnelProvider.RoutingMethod]: The method by which network traffic is routed to the tunnel.
//   - [NETunnelProvider.AppRules]: The app rules dictating which apps use the current tunneling session.
//
// # Configuring the tunnel interface
//
//   - [NETunnelProvider.SetTunnelNetworkSettingsCompletionHandler]: Specify the network settings for the current tunneling session.
//
// # Communicating with the containing app
//
//   - [NETunnelProvider.HandleAppMessageCompletionHandler]: Handle messages sent by the tunnel provider extension’s containing app.
//
// # Setting tunnel status
//
//   - [NETunnelProvider.Reasserting]: Indicate to the system that the tunnel is being re-established.
//   - [NETunnelProvider.SetReasserting]
//
// # Errors
//
//   - [NETunnelProvider.NETunnelProviderErrorDomain]: The domain used for Tunnel Provider errors.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider
type NETunnelProvider struct {
	NEProvider
}

// NETunnelProviderFromID constructs a [NETunnelProvider] from an objc.ID.
//
// An abstract base class shared by NEPacketTunnelProvider and
// NEAppProxyProvider.
func NETunnelProviderFromID(id objc.ID) NETunnelProvider {
	return NETunnelProvider{NEProvider: NEProviderFromID(id)}
}

// NOTE: NETunnelProvider adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NETunnelProvider] class.
//
// # Getting the tunnel configuration
//
//   - [INETunnelProvider.ProtocolConfiguration]: The configuration of the current tunneling session.
//   - [INETunnelProvider.RoutingMethod]: The method by which network traffic is routed to the tunnel.
//   - [INETunnelProvider.AppRules]: The app rules dictating which apps use the current tunneling session.
//
// # Configuring the tunnel interface
//
//   - [INETunnelProvider.SetTunnelNetworkSettingsCompletionHandler]: Specify the network settings for the current tunneling session.
//
// # Communicating with the containing app
//
//   - [INETunnelProvider.HandleAppMessageCompletionHandler]: Handle messages sent by the tunnel provider extension’s containing app.
//
// # Setting tunnel status
//
//   - [INETunnelProvider.Reasserting]: Indicate to the system that the tunnel is being re-established.
//   - [INETunnelProvider.SetReasserting]
//
// # Errors
//
//   - [INETunnelProvider.NETunnelProviderErrorDomain]: The domain used for Tunnel Provider errors.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider
type INETunnelProvider interface {
	INEProvider

	// Topic: Getting the tunnel configuration

	// The configuration of the current tunneling session.
	ProtocolConfiguration() INEVPNProtocol
	// The method by which network traffic is routed to the tunnel.
	RoutingMethod() NETunnelProviderRoutingMethod
	// The app rules dictating which apps use the current tunneling session.
	AppRules() []NEAppRule

	// Topic: Configuring the tunnel interface

	// Specify the network settings for the current tunneling session.
	SetTunnelNetworkSettingsCompletionHandler(tunnelNetworkSettings INETunnelNetworkSettings, completionHandler ErrorHandler)

	// Topic: Communicating with the containing app

	// Handle messages sent by the tunnel provider extension’s containing app.
	HandleAppMessageCompletionHandler(messageData foundation.INSData, completionHandler DataHandler)

	// Topic: Setting tunnel status

	// Indicate to the system that the tunnel is being re-established.
	Reasserting() bool
	SetReasserting(value bool)

	// Topic: Errors

	// The domain used for Tunnel Provider errors.
	NETunnelProviderErrorDomain() string
}

// Init initializes the instance.
func (t NETunnelProvider) Init() NETunnelProvider {
	rv := objc.Send[NETunnelProvider](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NETunnelProvider) Autorelease() NETunnelProvider {
	rv := objc.Send[NETunnelProvider](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETunnelProvider creates a new NETunnelProvider instance.
func NewNETunnelProvider() NETunnelProvider {
	class := getNETunnelProviderClass()
	rv := objc.Send[NETunnelProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Specify the network settings for the current tunneling session.
//
// tunnelNetworkSettings: The network settings to use for the tunnel. Pass nil to clear out the
// network settings for the current tunneling session.
//
// completionHandler: A block that will be executed when the operation of setting the network
// settings is complete. If the network settings could not be set due to an
// error, then the error parameter will be set to an [NSError] object
// containing more information about the error. See [NETunnelProviderError]
// for possible error codes. If the network settings were set successfully
// then the error parameter will be set to nil.
//
// # Discussion
//
// Use this method to specify the settings to be used by network communication
// that traverses the tunnel. If you are implementing a Packet Tunnel
// Provider, pass a [NEPacketTunnelNetworkSettings] object containing virtual
// IP configuration, DNS settings, proxy settings, the tunnel MTU, and IP
// routes. If you are implementing an App Proxy Provider, pass a
// [NETunnelNetworkSettings] containing DNS settings and proxy settings.
//
// This method should be called as part of the process of establishing the
// tunnel, as follows:
//
// - The system calls the appropriate “start” method on the tunnel
// provider object. - The provider obtains the network settings for the tunnel
// by some means dictated by the tunnel provider, such as by downloading them
// from the tunnel server. - The tunnel provider calls
// [SetTunnelNetworkSettingsCompletionHandler] method to apply the network
// settings to the system. - The system executes the completion handler passed
// to `completionHandler`: to indicate that the network settings have been
// set. - The tunnel provider executes the completion handler block passed to
// the “start” method to indicate that the tunnel is fully established.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/setTunnelNetworkSettings(_:completionHandler:)
//
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
func (t NETunnelProvider) SetTunnelNetworkSettingsCompletionHandler(tunnelNetworkSettings INETunnelNetworkSettings, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](t.ID, objc.Sel("setTunnelNetworkSettings:completionHandler:"), tunnelNetworkSettings, _block1)
}

// Handle messages sent by the tunnel provider extension’s containing app.
//
// messageData: The message data sent by the tunnel provider extension’s containing app.
//
// completionHandler: A block to be executed by the Tunnel Provider when it is finished handling
// the message. It may be nil, in which case the containing app does not
// expect a reply. The provider can send information back to the containing
// app via the `responseData` parameter.
//
// # Discussion
//
// Use this method to communicate information between the Tunnel Provider and
// the Tunnel Provider’s containing app.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/handleAppMessage(_:completionHandler:)
func (t NETunnelProvider) HandleAppMessageCompletionHandler(messageData foundation.INSData, completionHandler DataHandler) {
	_block1, _ := NewDataBlock(completionHandler)
	objc.Send[objc.ID](t.ID, objc.Sel("handleAppMessage:completionHandler:"), messageData, _block1)
}

// The configuration of the current tunneling session.
//
// # Discussion
//
// The configuration is created by the containing app of the Tunnel Provider
// using the [NETunnelProviderManager] class, or by the ingestion of a
// `com.AppleXCUIElementTypeVpnXCUIElementTypeManaged()` or a
// `com.AppleXCUIElementTypeVpnXCUIElementTypeManagedXCUIElementTypeApplayer()`
// configuration profile payload. See the [NETunnelProviderManager] class for
// more details.
//
// For [NEPacketTunnelProvider] subclasses and [NEAppProxyProvider]
// subclasses, this property will be set to a [NETunnelProviderProtocol]
// object.
//
// [NETunnelProvider] subclasses can observe this property using KVO to be
// notified when the configuration changes. For details see [Key-Value
// Observing Programming Guide].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/protocolConfiguration
//
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
func (t NETunnelProvider) ProtocolConfiguration() INEVPNProtocol {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("protocolConfiguration"))
	return NEVPNProtocolFromID(objc.ID(rv))
}

// The method by which network traffic is routed to the tunnel.
//
// # Discussion
//
// The default is [NETunnelProviderRoutingMethodDestinationIP].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/routingMethod
func (t NETunnelProvider) RoutingMethod() NETunnelProviderRoutingMethod {
	rv := objc.Send[NETunnelProviderRoutingMethod](t.ID, objc.Sel("routingMethod"))
	return NETunnelProviderRoutingMethod(rv)
}

// The app rules dictating which apps use the current tunneling session.
//
// # Discussion
//
// This property is only non-`nil` if the current configuration is a Per-App
// VPN configuration.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/appRules
func (t NETunnelProvider) AppRules() []NEAppRule {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("appRules"))
	return objc.ConvertSlice(rv, func(id objc.ID) NEAppRule {
		return NEAppRuleFromID(id)
	})
}

// Indicate to the system that the tunnel is being re-established.
//
// # Discussion
//
// The Tunnel Provider should set this property to true whenever it starts to
// reconnect to the tunnel server. Once the Tunnel Provider completes the
// process of reconnecting it should set this property to false.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/reasserting
func (t NETunnelProvider) Reasserting() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("reasserting"))
	return rv
}
func (t NETunnelProvider) SetReasserting(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setReasserting:"), value)
}

// The domain used for Tunnel Provider errors.
//
// See: https://developer.apple.com/documentation/networkextension/netunnelprovidererrordomain
func (t NETunnelProvider) NETunnelProviderErrorDomain() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("NETunnelProviderErrorDomain"))
	return foundation.NSStringFromID(rv).String()
}

// SetTunnelNetworkSettings is a synchronous wrapper around [NETunnelProvider.SetTunnelNetworkSettingsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t NETunnelProvider) SetTunnelNetworkSettings(ctx context.Context, tunnelNetworkSettings INETunnelNetworkSettings) error {
	done := make(chan error, 1)
	t.SetTunnelNetworkSettingsCompletionHandler(tunnelNetworkSettings, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// HandleAppMessage is a synchronous wrapper around [NETunnelProvider.HandleAppMessageCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t NETunnelProvider) HandleAppMessage(ctx context.Context, messageData foundation.INSData) (*foundation.NSData, error) {
	done := make(chan *foundation.NSData, 1)
	t.HandleAppMessageCompletionHandler(messageData, func(val *foundation.NSData) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
