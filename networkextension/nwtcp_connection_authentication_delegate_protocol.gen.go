// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A delegate protocol to customize the TLS authentication done by a connection.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnectionAuthenticationDelegate
type NWTCPConnectionAuthenticationDelegate interface {
	objectivec.IObject
}

// NWTCPConnectionAuthenticationDelegateObject wraps an existing Objective-C object that conforms to the NWTCPConnectionAuthenticationDelegate protocol.
type NWTCPConnectionAuthenticationDelegateObject struct {
	objectivec.Object
}
func (o NWTCPConnectionAuthenticationDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NWTCPConnectionAuthenticationDelegateObjectFromID constructs a [NWTCPConnectionAuthenticationDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NWTCPConnectionAuthenticationDelegateObjectFromID(id objc.ID) NWTCPConnectionAuthenticationDelegateObject {
	return NWTCPConnectionAuthenticationDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// NWTCPConnectionAuthenticationDelegateConfig holds optional typed callbacks for [NWTCPConnectionAuthenticationDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/networkextension/nwtcpconnectionauthenticationdelegate
type NWTCPConnectionAuthenticationDelegateConfig struct {

	// Other Methods
	// ShouldEvaluateTrustForConnection — Indicate that the delegate should override the default trust evaluation for the connection.
	ShouldEvaluateTrustForConnection func(connection NWTCPConnection) bool
	// ShouldProvideIdentityForConnection — Indicate that the delegate can provide an identity for the connection authentication.
	ShouldProvideIdentityForConnection func(connection NWTCPConnection) bool
}

// NewNWTCPConnectionAuthenticationDelegate creates an Objective-C object implementing the [NWTCPConnectionAuthenticationDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NWTCPConnectionAuthenticationDelegateObject] satisfies the [NWTCPConnectionAuthenticationDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/networkextension/nwtcpconnectionauthenticationdelegate
func NewNWTCPConnectionAuthenticationDelegate(config NWTCPConnectionAuthenticationDelegateConfig) NWTCPConnectionAuthenticationDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNWTCPConnectionAuthenticationDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ShouldEvaluateTrustForConnection != nil {
		fn := config.ShouldEvaluateTrustForConnection
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("shouldEvaluateTrustForConnection:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connectionID objc.ID) bool {
				connection := NWTCPConnectionFromID(connectionID)
				return fn(connection)
			},
		})
	}

	if config.ShouldProvideIdentityForConnection != nil {
		fn := config.ShouldProvideIdentityForConnection
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("shouldProvideIdentityForConnection:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connectionID objc.ID) bool {
				connection := NWTCPConnectionFromID(connectionID)
				return fn(connection)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NWTCPConnectionAuthenticationDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNWTCPConnectionAuthenticationDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NWTCPConnectionAuthenticationDelegateObjectFromID(instance)
}

