// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// An interface you use to manage connections between the guest operating system and host computer.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketListenerDelegate
type VZVirtioSocketListenerDelegate interface {
	objectivec.IObject
}

// VZVirtioSocketListenerDelegateObject wraps an existing Objective-C object that conforms to the VZVirtioSocketListenerDelegate protocol.
type VZVirtioSocketListenerDelegateObject struct {
	objectivec.Object
}
func (o VZVirtioSocketListenerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZVirtioSocketListenerDelegateObjectFromID constructs a [VZVirtioSocketListenerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZVirtioSocketListenerDelegateObjectFromID(id objc.ID) VZVirtioSocketListenerDelegateObject {
	return VZVirtioSocketListenerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns a Boolean value that indicates whether to accept a new connection
// from the guest operating system.
//
// listener: The listener object that monitors the associated port.
//
// connection: The object that contains information about the proposed connection. Use
// this object to fetch port information.
//
// socketDevice: The Virtio socket device that requested the connection.
//
// # Return Value
// 
// [true] to establish the connection, or [false] to reject it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use your method’s implementation to quickly determine whether to accept
// or reject connection attempts. A typical implementation verifies that a
// connection between the specified ports is permissible. Return a result as
// quickly as possible, and don’t perform any long-running operations in
// this method.
// 
// If you don’t implement this method, the virtual machine refuses all
// connection requests as if this method returned [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketListenerDelegate/listener(_:shouldAcceptNewConnection:from:)
func (o VZVirtioSocketListenerDelegateObject) ListenerShouldAcceptNewConnectionFromSocketDevice(listener IVZVirtioSocketListener, connection IVZVirtioSocketConnection, socketDevice IVZVirtioSocketDevice) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("listener:shouldAcceptNewConnection:fromSocketDevice:"), listener, connection, socketDevice)
	return rv
	}

// VZVirtioSocketListenerDelegateConfig holds optional typed callbacks for [VZVirtioSocketListenerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/virtualization/vzvirtiosocketlistenerdelegate
type VZVirtioSocketListenerDelegateConfig struct {

	// Other Methods
	// ListenerShouldAcceptNewConnectionFromSocketDevice — Returns a Boolean value that indicates whether to accept a new connection from the guest operating system.
	ListenerShouldAcceptNewConnectionFromSocketDevice func(listener VZVirtioSocketListener, connection VZVirtioSocketConnection, socketDevice VZVirtioSocketDevice) bool
}

// NewVZVirtioSocketListenerDelegate creates an Objective-C object implementing the [VZVirtioSocketListenerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [VZVirtioSocketListenerDelegateObject] satisfies the [VZVirtioSocketListenerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/virtualization/vzvirtiosocketlistenerdelegate
func NewVZVirtioSocketListenerDelegate(config VZVirtioSocketListenerDelegateConfig) VZVirtioSocketListenerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoVZVirtioSocketListenerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ListenerShouldAcceptNewConnectionFromSocketDevice != nil {
		fn := config.ListenerShouldAcceptNewConnectionFromSocketDevice
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("listener:shouldAcceptNewConnection:fromSocketDevice:"),
			Fn: func(self objc.ID, _cmd objc.SEL, listenerID objc.ID, connectionID objc.ID, socketDeviceID objc.ID) bool {
				listener := VZVirtioSocketListenerFromID(listenerID)
				connection := VZVirtioSocketConnectionFromID(connectionID)
				socketDevice := VZVirtioSocketDeviceFromID(socketDeviceID)
				return fn(listener, connection, socketDevice)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("VZVirtioSocketListenerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewVZVirtioSocketListenerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return VZVirtioSocketListenerDelegateObjectFromID(instance)
}

