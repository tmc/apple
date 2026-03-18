// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// The protocol that delegates to the XPC listener use to accept or reject new connections.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCListenerDelegate
type NSXPCListenerDelegate interface {
	objectivec.IObject
}

// NSXPCListenerDelegateObject wraps an existing Objective-C object that conforms to the NSXPCListenerDelegate protocol.
type NSXPCListenerDelegateObject struct {
	objectivec.Object
}
func (o NSXPCListenerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSXPCListenerDelegateObjectFromID constructs a [NSXPCListenerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSXPCListenerDelegateObjectFromID(id objc.ID) NSXPCListenerDelegateObject {
	return NSXPCListenerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Accepts or rejects a new connection to the listener.
//
// # Discussion
// 
// To accept the connection, first configure the connection if desired, then
// call [Resume] on the new connection, then return [true].
// 
// To reject the connect, return a value of [false]. This causes the
// connection object to be invalidated.
// 
// In this method, you can also set up properties on the connection object,
// such as its exported object and interfaces. Be sure to call [Resume] when
// you are finished configuring the connection object and are ready for it to
// receive messages.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCListenerDelegate/listener(_:shouldAcceptNewConnection:)

func (o NSXPCListenerDelegateObject) ListenerShouldAcceptNewConnection(listener INSXPCListener, newConnection INSXPCConnection) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("listener:shouldAcceptNewConnection:"), listener, newConnection)
	return rv
	}

// NSXPCListenerDelegateConfig holds optional typed callbacks for [NSXPCListenerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsxpclistenerdelegate
type NSXPCListenerDelegateConfig struct {

	// Instance Methods
	// ListenerShouldAcceptNewConnection — Accepts or rejects a new connection to the listener.
	ListenerShouldAcceptNewConnection func(listener NSXPCListener, newConnection NSXPCConnection) bool
}

// NewNSXPCListenerDelegate creates an Objective-C object implementing the [NSXPCListenerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSXPCListenerDelegateObject] satisfies the [NSXPCListenerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsxpclistenerdelegate
func NewNSXPCListenerDelegate(config NSXPCListenerDelegateConfig) NSXPCListenerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSXPCListenerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ListenerShouldAcceptNewConnection != nil {
		fn := config.ListenerShouldAcceptNewConnection
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("listener:shouldAcceptNewConnection:"),
			Fn: func(self objc.ID, _cmd objc.SEL, listenerID objc.ID, newConnectionID objc.ID) bool {
				listener := NSXPCListenerFromID(listenerID)
				newConnection := NSXPCConnectionFromID(newConnectionID)
				return fn(listener, newConnection)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSXPCListenerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSXPCListenerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSXPCListenerDelegateObjectFromID(instance)
}

