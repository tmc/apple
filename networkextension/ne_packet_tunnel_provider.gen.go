// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEPacketTunnelProvider] class.
var (
	_NEPacketTunnelProviderClass     NEPacketTunnelProviderClass
	_NEPacketTunnelProviderClassOnce sync.Once
)

func getNEPacketTunnelProviderClass() NEPacketTunnelProviderClass {
	_NEPacketTunnelProviderClassOnce.Do(func() {
		_NEPacketTunnelProviderClass = NEPacketTunnelProviderClass{class: objc.GetClass("NEPacketTunnelProvider")}
	})
	return _NEPacketTunnelProviderClass
}

// GetNEPacketTunnelProviderClass returns the class object for NEPacketTunnelProvider.
func GetNEPacketTunnelProviderClass() NEPacketTunnelProviderClass {
	return getNEPacketTunnelProviderClass()
}

type NEPacketTunnelProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEPacketTunnelProviderClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEPacketTunnelProviderClass) Alloc() NEPacketTunnelProvider {
	rv := objc.Send[NEPacketTunnelProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The principal class for a packet tunnel provider app extension.
//
// # Overview
// 
// The [NEPacketTunnelProvider] class gives its subclasses access to a virtual
// network interface via the [NEPacketTunnelProvider.PacketFlow] property. Use the
// [SetTunnelNetworkSettingsCompletionHandler] method in the Packet Tunnel
// Provider to specify that the following network settings be associated with
// the virtual interface:
// 
// - Virtual IP address - DNS resolver configuration - HTTP proxy
// configuration - IP destination networks to be routed through the tunnel -
// IP destination networks to be routed outside the tunnel - Interface MTU
// 
// By specifying IP destination networks, the Packet Tunnel Provider can
// dictate what IP destinations will be routed to the virtual interface. IP
// packets with matching destination addresses will then be diverted to Packet
// Tunnel Provider and can be read using the [NEPacketTunnelProvider.PacketFlow] property. The Packet
// Tunnel Provider can then encapsulate the IP packets per a custom tunneling
// protocol and send them to a tunnel server. When the Packet Tunnel Provider
// decapsulates IP packets received from the tunnel server, it can use the
// [NEPacketTunnelProvider.PacketFlow] property to inject the packets into the networking stack.
// 
// # Creating a Packet Tunnel Provider Extension
// 
// Packet Tunnel Providers run as App Extensions for the
// `com.AppleXCUIElementTypeNetworkextensionXCUIElementTypePacket()-tunnel`
// extension point.
// 
// To create a Packet Tunnel Provider extension, first create a new App
// Extension target in your project.
// 
// For an example of an Xcode build target for this app extension, see the
// [SimpleTunnel: Customized Networking Using the NetworkExtension Framework]
// sample code project.
// 
// Once you have a Packet Tunnel Provider extension target, create a subclass
// of NEPacketTunnelProvider. Then, set the [NSExtensionPrincipalClass] key in
// the the extension’s `Info.Plist()` to the name of your subclass.
// 
// If it is not already, set the [NSExtensionPointIdentifier] key in the
// extension’s `Info.Plist()` to
// `com.AppleXCUIElementTypeNetworkextensionXCUIElementTypePacket()-tunnel`.
// 
// Here is an example of the NSExtension dictionary in a Packet Tunnel
// Provider extension’s `Info.Plist()`:
// 
// Finally, add the Packet Tunnel Provider extension target to your app’s
// Embed App Extensions build phase.
// 
// # Subclassing Notes
// 
// In order to create a Packet Tunnel Provider extension, you must create a
// subclass of [NEPacketTunnelProvider] and override the methods listed below.
// 
// # Methods to Override
// 
// - [NEPacketTunnelProvider.StartTunnelWithOptionsCompletionHandler] -
// [NEPacketTunnelProvider.StopTunnelWithReasonCompletionHandler]
//
// [SimpleTunnel: Customized Networking Using the NetworkExtension Framework]: https://developer.apple.com/library/archive/samplecode/SimpleTunnel/Introduction/Intro.html#//apple_ref/doc/uid/TP40016140
//
// # Managing the tunnel life cycle
//
//   - [NEPacketTunnelProvider.StartTunnelWithOptionsCompletionHandler]: Start the network tunnel.
//   - [NEPacketTunnelProvider.StopTunnelWithReasonCompletionHandler]: Stop the network tunnel.
//   - [NEPacketTunnelProvider.CancelTunnelWithError]: Stop the network tunnel from the Packet Tunnel Provider.
//
// # Handling IP packets
//
//   - [NEPacketTunnelProvider.PacketFlow]: A [NEPacketTunnelFlow](<doc://com.apple.networkextension/documentation/NetworkExtension/NEPacketTunnelFlow>) object which is used to receive IP packets routed to the tunnel’s virtual interface and inject IP packets into the networking stack via the tunnel’s virtual interface.
//
// # Instance Properties
//
//   - [NEPacketTunnelProvider.VirtualInterface]
//   - [NEPacketTunnelProvider.SetVirtualInterface]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelProvider
type NEPacketTunnelProvider struct {
	NETunnelProvider
}

// NEPacketTunnelProviderFromID constructs a [NEPacketTunnelProvider] from an objc.ID.
//
// The principal class for a packet tunnel provider app extension.
func NEPacketTunnelProviderFromID(id objc.ID) NEPacketTunnelProvider {
	return NEPacketTunnelProvider{NETunnelProvider: NETunnelProviderFromID(id)}
}
// NOTE: NEPacketTunnelProvider adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEPacketTunnelProvider] class.
//
// # Managing the tunnel life cycle
//
//   - [INEPacketTunnelProvider.StartTunnelWithOptionsCompletionHandler]: Start the network tunnel.
//   - [INEPacketTunnelProvider.StopTunnelWithReasonCompletionHandler]: Stop the network tunnel.
//   - [INEPacketTunnelProvider.CancelTunnelWithError]: Stop the network tunnel from the Packet Tunnel Provider.
//
// # Handling IP packets
//
//   - [INEPacketTunnelProvider.PacketFlow]: A [NEPacketTunnelFlow](<doc://com.apple.networkextension/documentation/NetworkExtension/NEPacketTunnelFlow>) object which is used to receive IP packets routed to the tunnel’s virtual interface and inject IP packets into the networking stack via the tunnel’s virtual interface.
//
// # Instance Properties
//
//   - [INEPacketTunnelProvider.VirtualInterface]
//   - [INEPacketTunnelProvider.SetVirtualInterface]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelProvider
type INEPacketTunnelProvider interface {
	INETunnelProvider

	// Topic: Managing the tunnel life cycle

	// Start the network tunnel.
	StartTunnelWithOptionsCompletionHandler(options foundation.INSDictionary, completionHandler ErrorHandler)
	// Stop the network tunnel.
	StopTunnelWithReasonCompletionHandler(reason NEProviderStopReason, completionHandler VoidHandler)
	// Stop the network tunnel from the Packet Tunnel Provider.
	CancelTunnelWithError(error_ foundation.INSError)

	// Topic: Handling IP packets

	// A [NEPacketTunnelFlow](<doc://com.apple.networkextension/documentation/NetworkExtension/NEPacketTunnelFlow>) object which is used to receive IP packets routed to the tunnel’s virtual interface and inject IP packets into the networking stack via the tunnel’s virtual interface.
	PacketFlow() INEPacketTunnelFlow

	// Topic: Instance Properties

	VirtualInterface() objectivec.IObject
	SetVirtualInterface(value objectivec.IObject)
}

// Init initializes the instance.
func (p NEPacketTunnelProvider) Init() NEPacketTunnelProvider {
	rv := objc.Send[NEPacketTunnelProvider](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NEPacketTunnelProvider) Autorelease() NEPacketTunnelProvider {
	rv := objc.Send[NEPacketTunnelProvider](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEPacketTunnelProvider creates a new NEPacketTunnelProvider instance.
func NewNEPacketTunnelProvider() NEPacketTunnelProvider {
	class := getNEPacketTunnelProviderClass()
	rv := objc.Send[NEPacketTunnelProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Start the network tunnel.
//
// options: A dictionary passed by the app that requested that the tunnel be started.
// If the starting app did not specify a dictionary of options then this
// parameter will be nil. If the tunnel was started via Connect On Demand,
// then this parameter will be nil.
//
// completionHandler: A block that must be executed when the tunnel is fully established, or when
// the tunnel cannot be started due to an error. If the tunnel was
// successfully established, then the error parameter must be set to nil. If
// an error occurred, the error parameter passed to this block must be set to
// a non-nil [NSError] object.
// //
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// # Discussion
// 
// This method is called by the system to start the network tunnel.
// 
// [NEPacketTunnelProvider] subclasses must override this method.
// 
// When the Packet Tunnel Provider executes the completionHandler block with a
// nil error parameter, it signals to the system that it is ready to begin
// handling network data. Therefore, the Packet Tunnel Provider should call
// [SetTunnelNetworkSettingsCompletionHandler] and wait for it to complete
// before executing the completionHandler block.
// 
// The domain and code of the [NSError] object passed to the
// `completionHandler` block are defined by the Packet Tunnel Provider.
//
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelProvider/startTunnel(options:completionHandler:)
func (p NEPacketTunnelProvider) StartTunnelWithOptionsCompletionHandler(options foundation.INSDictionary, completionHandler ErrorHandler) {
_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](p.ID, objc.Sel("startTunnelWithOptions:completionHandler:"), options, _block1)
}
// Stop the network tunnel.
//
// reason: An [NEProviderStopReason] code indicating why the tunnel is being stopped.
// Possible codes are listed in [NEProvider].
//
// completionHandler: A block that must be executed when the tunnel is fully stopped.
//
// # Discussion
// 
// This method is called by the system to stop the network tunnel.
// 
// NEPacketTunnelProvider subclasses must override this method.
// 
// Do not use this method to stop the tunnel from the Packet Tunnel Provider.
// Use `cancelTunnelWithError`: instead.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelProvider/stopTunnel(with:completionHandler:)
func (p NEPacketTunnelProvider) StopTunnelWithReasonCompletionHandler(reason NEProviderStopReason, completionHandler VoidHandler) {
_block1, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](p.ID, objc.Sel("stopTunnelWithReason:completionHandler:"), reason, _block1)
}
// Stop the network tunnel from the Packet Tunnel Provider.
//
// error: An [NSError] object containing the error that caused the tunnel to be
// stopped. The domain and code of this NSError object is defined by the
// caller.
// //
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// # Discussion
// 
// The Packet Tunnel Provider should call this method when an unrecoverable
// error occurs, such as the tunnel server going down or the VPN
// authentication session expiring.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelProvider/cancelTunnelWithError(_:)
func (p NEPacketTunnelProvider) CancelTunnelWithError(error_ foundation.INSError) {
	objc.Send[objc.ID](p.ID, objc.Sel("cancelTunnelWithError:"), error_)
}

// A [NEPacketTunnelFlow] object which is used to receive IP packets routed to
// the tunnel’s virtual interface and inject IP packets into the networking
// stack via the tunnel’s virtual interface.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelProvider/packetFlow
func (p NEPacketTunnelProvider) PacketFlow() INEPacketTunnelFlow {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("packetFlow"))
	return NEPacketTunnelFlowFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/networkextension/nepackettunnelprovider/virtualinterface-7l3ol
func (p NEPacketTunnelProvider) VirtualInterface() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("virtualInterface"))
	return objectivec.Object{ID: rv}
}
func (p NEPacketTunnelProvider) SetVirtualInterface(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setVirtualInterface:"), value)
}

// StartTunnelWithOptions is a synchronous wrapper around [NEPacketTunnelProvider.StartTunnelWithOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (p NEPacketTunnelProvider) StartTunnelWithOptions(ctx context.Context, options foundation.INSDictionary) error {
	done := make(chan error, 1)
	p.StartTunnelWithOptionsCompletionHandler(options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// StopTunnelWithReason is a synchronous wrapper around [NEPacketTunnelProvider.StopTunnelWithReasonCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (p NEPacketTunnelProvider) StopTunnelWithReason(ctx context.Context, reason NEProviderStopReason) error {
	done := make(chan struct{}, 1)
	p.StopTunnelWithReasonCompletionHandler(reason, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

