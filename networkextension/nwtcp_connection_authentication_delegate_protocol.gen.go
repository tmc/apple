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

// Indicate that the delegate should override the default trust evaluation for
// the connection.
//
// connection: The connection sending this message
//
// # Return Value
//
// Return true to take over the default trust evaluation, in which case the
// delegate method `completionHandler`: will be called.
//
// # Discussion
//
// The caller can implement this optional protocol method to decide whether it
// wants to take over the default trust evaluation for this connection. If
// this delegate method is not implemented, the return value will default to
// YES if “ is implemented.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnectionAuthenticationDelegate/shouldEvaluateTrust(for:)
func (o NWTCPConnectionAuthenticationDelegateObject) ShouldEvaluateTrustForConnection(connection INWTCPConnection) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("shouldEvaluateTrustForConnection:"), connection)
	return rv
}

// Override the default trust evaluation for the connection.
//
// connection: The connection sending this message
//
// peerCertificateChain: The connection peer’s certificate chain
//
// completion: The completion handler for passing the [SecTrust] object to the connection.
// The [SecTrustRef] object `trust` is required and must not be `nil`. It will
// be evaluated using [SecTrustEvaluate(_:_:)] if necessary.
//
// The caller is responsible for keeping the argument object valid for the
// duration of the completion handler invocation.
//
// # Discussion
//
// The caller can implement this optional protocol method to set up custom
// policies for peer certificate trust evaluation. If the delegate method is
// implemented, the caller is responsible for creating and setting up the
// [SecTrust] object and passing it to the completion handler. Otherwise, the
// default trust evaluation policy is used for the connection.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnectionAuthenticationDelegate/evaluateTrust(for:peerCertificateChain:completionHandler:)
//
// [SecTrustEvaluate(_:_:)]: https://developer.apple.com/documentation/Security/SecTrustEvaluate(_:_:)
// [SecTrust]: https://developer.apple.com/documentation/Security/SecTrust
//
// [SecTrust]: https://developer.apple.com/documentation/Security/SecTrust
func (o NWTCPConnectionAuthenticationDelegateObject) EvaluateTrustForConnectionPeerCertificateChainCompletionHandler(connection INWTCPConnection, peerCertificateChain []objectivec.IObject, completion SecTrustRefHandler) {
	_block2, _ := NewSecTrustRefBlock(completion)
	objc.Send[struct{}](o.ID, objc.Sel("evaluateTrustForConnection:peerCertificateChain:completionHandler:"), connection, peerCertificateChain, _block2)
}

// Indicate that the delegate can provide an identity for the connection
// authentication.
//
// connection: The connection sending this message.
//
// # Return Value
//
// Return true to provide the identity for this connection, in which case the
// delegate method [ProvideIdentityForConnectionCompletionHandler] will be
// called.
//
// # Discussion
//
// The caller can implement this optional protocol method to decide whether it
// wants to provide the identity for this connection for authentication. If
// this delegate method is not implemented, the return value will default to
// YES if [ProvideIdentityForConnectionCompletionHandler] is implemented.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnectionAuthenticationDelegate/shouldProvideIdentity(for:)
func (o NWTCPConnectionAuthenticationDelegateObject) ShouldProvideIdentityForConnection(connection INWTCPConnection) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("shouldProvideIdentityForConnection:"), connection)
	return rv
}

// Provide the identity and an optional certificate chain to be used for
// authentication.
//
// connection: The connection sending this message
//
// completion: The completion handler for passing an identity and certificate chain to the
// connection. The `identity` is required and must not be `nil`. The
// `certificateChain` argument is optional, and is an array of one or more
// [SecCertificate] objects. The certificate chain must contain objects of
// type [SecCertificateRef] only. If the certificate chain is set, it will be
// used. Otherwise, the leaf certificate will be extracted from the
// [SecIdentity] object and will be used for authentication.
//
// The caller is responsible for keeping the argument object(s) valid for the
// duration of the completion handler invocation.
//
// # Discussion
//
// Optional. If this method is not implemented, the default certificate
// evaluation will be used.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnectionAuthenticationDelegate/provideIdentity(for:completionHandler:)
//
// [SecCertificate]: https://developer.apple.com/documentation/Security/SecCertificate
// [SecIdentity]: https://developer.apple.com/documentation/Security/SecIdentity
func (o NWTCPConnectionAuthenticationDelegateObject) ProvideIdentityForConnectionCompletionHandler(connection INWTCPConnection, completion SecIdentityRefArrayHandler) {
	_block1, _ := NewSecIdentityRefArrayBlock(completion)
	objc.Send[struct{}](o.ID, objc.Sel("provideIdentityForConnection:completionHandler:"), connection, _block1)
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
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("NWTCPConnectionAuthenticationDelegate", "shouldEvaluateTrustForConnection:")
					}
				}()
				connection := NWTCPConnectionFromID(connectionID)
				_delegateResult := fn(connection)
				_delegateDone = true
				return _delegateResult
			},
		})
	}

	if config.ShouldProvideIdentityForConnection != nil {
		fn := config.ShouldProvideIdentityForConnection
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("shouldProvideIdentityForConnection:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connectionID objc.ID) bool {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("NWTCPConnectionAuthenticationDelegate", "shouldProvideIdentityForConnection:")
					}
				}()
				connection := NWTCPConnectionFromID(connectionID)
				_delegateResult := fn(connection)
				_delegateDone = true
				return _delegateResult
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
