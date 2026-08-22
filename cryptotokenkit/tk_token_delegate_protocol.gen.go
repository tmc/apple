// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// The interface that a token delegate implements to respond to session creation events.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDelegate
type TKTokenDelegate interface {
	objectivec.IObject

	// Tells the delegate to create a session for the specified token.
	//
	// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDelegate/createSession(_:)
	TokenCreateSessionWithError(token ITKToken) (ITKTokenSession, error)
}

// TKTokenDelegateObject wraps an existing Objective-C object that conforms to the TKTokenDelegate protocol.
type TKTokenDelegateObject struct {
	objectivec.Object
}

func (o TKTokenDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// TKTokenDelegateObjectFromID constructs a [TKTokenDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func TKTokenDelegateObjectFromID(id objc.ID) TKTokenDelegateObject {
	return TKTokenDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate to create a session for the specified token.
//
// token: The token.
//
// # Return Value
//
// A new token session, or `nil` if an error occurred.
//
// # Discussion
//
// All operations for a token are performed within a session representing an
// authentication context. This delegate method is called whenever new
// authentication context is needed. For example, a client may want to perform
// a token operation using a keychain object that has an associated
// [LAContext].
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDelegate/createSession(_:)
func (o TKTokenDelegateObject) TokenCreateSessionWithError(token ITKToken) (ITKTokenSession, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("token:createSessionWithError:"), token)
	if err != nil {
		return nil, err
	}
	return TKTokenSessionFromID(rv), nil
}

// Tells the delegate to terminate the specified token session.
//
// token: The token.
//
// session: The token session to be terminated.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDelegate/token(_:terminateSession:)
func (o TKTokenDelegateObject) TokenTerminateSession(token ITKToken, session ITKTokenSession) {
	objc.Send[struct{}](o.ID, objc.Sel("token:terminateSession:"), token, session)
}

// TKTokenDelegateConfig holds optional typed callbacks for [TKTokenDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/cryptotokenkit/tktokendelegate
type TKTokenDelegateConfig struct {

	// Delegate Methods
	// TokenTerminateSession — Tells the delegate to terminate the specified token session.
	TokenTerminateSession func(token TKToken, session TKTokenSession)

	// Other Methods
	// TokenCreateSessionWithError — Tells the delegate to create a session for the specified token.
	TokenCreateSessionWithError func(token TKToken, error_ foundation.NSError) TKTokenSession
}

// NewTKTokenDelegate creates an Objective-C object implementing the [TKTokenDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [TKTokenDelegateObject] satisfies the [TKTokenDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/cryptotokenkit/tktokendelegate
func NewTKTokenDelegate(config TKTokenDelegateConfig) TKTokenDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoTKTokenDelegate_%d", n)

	var methods []objc.MethodDef

	if config.TokenCreateSessionWithError != nil {
		fn := config.TokenCreateSessionWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("token:createSessionWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tokenID objc.ID, error_ID objc.ID) objc.ID {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("TKTokenDelegate", "token:createSessionWithError:")
					}
				}()
				token := TKTokenFromID(tokenID)
				error_ := foundation.NSErrorFromID(error_ID)
				_delegateResult := fn(token, error_).GetID()
				_delegateDone = true
				return _delegateResult
			},
		})
	}

	if config.TokenTerminateSession != nil {
		fn := config.TokenTerminateSession
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("token:terminateSession:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tokenID objc.ID, sessionID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("TKTokenDelegate", "token:terminateSession:")
					}
				}()
				token := TKTokenFromID(tokenID)
				session := TKTokenSessionFromID(sessionID)
				fn(token, session)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("TKTokenDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewTKTokenDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return TKTokenDelegateObjectFromID(instance)
}
