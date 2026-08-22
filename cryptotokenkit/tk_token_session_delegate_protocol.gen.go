// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// The interface that a session instance delegate implements to respond to token session authentication events.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSessionDelegate
type TKTokenSessionDelegate interface {
	objectivec.IObject
}

// TKTokenSessionDelegateObject wraps an existing Objective-C object that conforms to the TKTokenSessionDelegate protocol.
type TKTokenSessionDelegateObject struct {
	objectivec.Object
}

func (o TKTokenSessionDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// TKTokenSessionDelegateObjectFromID constructs a [TKTokenSessionDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func TKTokenSessionDelegateObjectFromID(id objc.ID) TKTokenSessionDelegateObject {
	return TKTokenSessionDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the delegate whether the token session supports a given operation
// using the specified key and algorithm.
//
// session: The token session.
//
// operation: The operation to perform. For possible values, see [TKTokenOperation].
//
// keyObjectID: The identifier of the private key object.
//
// algorithm: The algorithm to be used by the operation.
//
// # Return Value
//
// true if the operation is supported; otherwise, false.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSessionDelegate/tokenSession(_:supports:keyObjectID:algorithm:)
func (o TKTokenSessionDelegateObject) TokenSessionSupportsOperationUsingKeyAlgorithm(session ITKTokenSession, operation TKTokenOperation, keyObjectID TKTokenObjectID, algorithm ITKTokenKeyAlgorithm) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("tokenSession:supportsOperation:usingKey:algorithm:"), session, operation, keyObjectID, algorithm)
	return rv
}

// Tells the delegate that authentication has begun for the specified
// operation and constraint.
//
// session: The token session.
//
// operation: The kind of operation.
//
// constraint: The constraint to be satisfied.
//
// # Return Value
//
// The resulting context of the operation, or `nil` if an error occurred.
//
// # Discussion
//
// If you return an instance of a subclass of [TKTokenAuthOperation] that is
// provided by the CryptoTokenKit framework, the system will first fill in the
// context-specific properties, such as the password, before calling the “
// method on the context.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSessionDelegate/tokenSession(_:beginAuthFor:constraint:)
func (o TKTokenSessionDelegateObject) TokenSessionBeginAuthForOperationConstraintError(session ITKTokenSession, operation TKTokenOperation, constraint TKTokenOperationConstraint) (ITKTokenAuthOperation, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("tokenSession:beginAuthForOperation:constraint:error:"), session, operation, constraint)
	if err != nil {
		return nil, err
	}
	return TKTokenAuthOperationFromID(rv), nil
}

// Tells the delegate to sign a data object using the specified key and
// algorithm.
//
// session: The token session.
//
// dataToSign: The data to sign.
//
// keyObjectID: The identifier of the private key object.
//
// algorithm: The algorithm to be used for signing.
//
// # Return Value
//
// The signed data, or `nil` if an error occurred.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSessionDelegate/tokenSession(_:sign:keyObjectID:algorithm:)
func (o TKTokenSessionDelegateObject) TokenSessionSignDataUsingKeyAlgorithmError(session ITKTokenSession, dataToSign foundation.NSData, keyObjectID TKTokenObjectID, algorithm ITKTokenKeyAlgorithm) (foundation.NSData, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("tokenSession:signData:usingKey:algorithm:error:"), session, dataToSign, keyObjectID, algorithm)
	if err != nil {
		return foundation.NSData{}, err
	}
	return foundation.NSDataFromID(rv), nil
}

// Tells the delegate to decrypt a data object using the specified key and
// algorithm.
//
// session: The token session.
//
// ciphertext: The data to decrypt.
//
// keyObjectID: The identifier of the public key object.
//
// algorithm: The algorithm to be used for decryption.
//
// # Return Value
//
// The decrypted data, or `nil` if an error occurred.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSessionDelegate/tokenSession(_:decrypt:keyObjectID:algorithm:)
func (o TKTokenSessionDelegateObject) TokenSessionDecryptDataUsingKeyAlgorithmError(session ITKTokenSession, ciphertext foundation.NSData, keyObjectID TKTokenObjectID, algorithm ITKTokenKeyAlgorithm) (foundation.NSData, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("tokenSession:decryptData:usingKey:algorithm:error:"), session, ciphertext, keyObjectID, algorithm)
	if err != nil {
		return foundation.NSData{}, err
	}
	return foundation.NSDataFromID(rv), nil
}

// Tells the delegate to perform a key exchange using the specified key and
// algorithm.
//
// session: The token session.
//
// otherPartyPublicKeyData: The public key of the other party.
//
// objectID: The identifier of the private key object.
//
// algorithm: The algorithm to be used for key exchange.
//
// parameters: Additional parameters used by `algorithm` to perform the key exchange.
//
// # Return Value
//
// The result of the key exchange, or `nil` if an error occurred.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSessionDelegate/tokenSession(_:performKeyExchange:keyObjectID:algorithm:parameters:)
func (o TKTokenSessionDelegateObject) TokenSessionPerformKeyExchangeWithPublicKeyUsingKeyAlgorithmParametersError(session ITKTokenSession, otherPartyPublicKeyData foundation.NSData, objectID TKTokenObjectID, algorithm ITKTokenKeyAlgorithm, parameters ITKTokenKeyExchangeParameters) (foundation.NSData, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("tokenSession:performKeyExchangeWithPublicKey:usingKey:algorithm:parameters:error:"), session, otherPartyPublicKeyData, objectID, algorithm, parameters)
	if err != nil {
		return foundation.NSData{}, err
	}
	return foundation.NSDataFromID(rv), nil
}

// TKTokenSessionDelegateConfig holds optional typed callbacks for [TKTokenSessionDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/cryptotokenkit/tktokensessiondelegate
type TKTokenSessionDelegateConfig struct {

	// Other Methods
	// TokenSessionSupportsOperationUsingKeyAlgorithm — Asks the delegate whether the token session supports a given operation using the specified key and algorithm.
	TokenSessionSupportsOperationUsingKeyAlgorithm func(session TKTokenSession, operation TKTokenOperation, keyObjectID TKTokenObjectID, algorithm TKTokenKeyAlgorithm) bool
	// TokenSessionBeginAuthForOperationConstraintError — Tells the delegate that authentication has begun for the specified operation and constraint.
	TokenSessionBeginAuthForOperationConstraintError func(session TKTokenSession, operation TKTokenOperation, constraint TKTokenOperationConstraint, error_ foundation.NSError) TKTokenAuthOperation
	// TokenSessionSignDataUsingKeyAlgorithmError — Tells the delegate to sign a data object using the specified key and algorithm.
	TokenSessionSignDataUsingKeyAlgorithmError func(session TKTokenSession, dataToSign foundation.NSData, keyObjectID TKTokenObjectID, algorithm TKTokenKeyAlgorithm, error_ foundation.NSError) foundation.NSData
	// TokenSessionDecryptDataUsingKeyAlgorithmError — Tells the delegate to decrypt a data object using the specified key and algorithm.
	TokenSessionDecryptDataUsingKeyAlgorithmError func(session TKTokenSession, ciphertext foundation.NSData, keyObjectID TKTokenObjectID, algorithm TKTokenKeyAlgorithm, error_ foundation.NSError) foundation.NSData
	// TokenSessionPerformKeyExchangeWithPublicKeyUsingKeyAlgorithmParametersError — Tells the delegate to perform a key exchange using the specified key and algorithm.
	TokenSessionPerformKeyExchangeWithPublicKeyUsingKeyAlgorithmParametersError func(session TKTokenSession, otherPartyPublicKeyData foundation.NSData, objectID TKTokenObjectID, algorithm TKTokenKeyAlgorithm, parameters TKTokenKeyExchangeParameters, error_ foundation.NSError) foundation.NSData
}

// NewTKTokenSessionDelegate creates an Objective-C object implementing the [TKTokenSessionDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [TKTokenSessionDelegateObject] satisfies the [TKTokenSessionDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/cryptotokenkit/tktokensessiondelegate
func NewTKTokenSessionDelegate(config TKTokenSessionDelegateConfig) TKTokenSessionDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoTKTokenSessionDelegate_%d", n)

	var methods []objc.MethodDef

	if config.TokenSessionSupportsOperationUsingKeyAlgorithm != nil {
		fn := config.TokenSessionSupportsOperationUsingKeyAlgorithm
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tokenSession:supportsOperation:usingKey:algorithm:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, operation TKTokenOperation, keyObjectID TKTokenObjectID, algorithmID objc.ID) bool {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("TKTokenSessionDelegate", "tokenSession:supportsOperation:usingKey:algorithm:")
					}
				}()
				session := TKTokenSessionFromID(sessionID)
				algorithm := TKTokenKeyAlgorithmFromID(algorithmID)
				_delegateResult := fn(session, operation, keyObjectID, algorithm)
				_delegateDone = true
				return _delegateResult
			},
		})
	}

	if config.TokenSessionBeginAuthForOperationConstraintError != nil {
		fn := config.TokenSessionBeginAuthForOperationConstraintError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tokenSession:beginAuthForOperation:constraint:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, operation TKTokenOperation, constraint TKTokenOperationConstraint, error_ID objc.ID) objc.ID {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("TKTokenSessionDelegate", "tokenSession:beginAuthForOperation:constraint:error:")
					}
				}()
				session := TKTokenSessionFromID(sessionID)
				error_ := foundation.NSErrorFromID(error_ID)
				_delegateResult := fn(session, operation, constraint, error_).GetID()
				_delegateDone = true
				return _delegateResult
			},
		})
	}

	if config.TokenSessionSignDataUsingKeyAlgorithmError != nil {
		fn := config.TokenSessionSignDataUsingKeyAlgorithmError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tokenSession:signData:usingKey:algorithm:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, dataToSignID objc.ID, keyObjectID TKTokenObjectID, algorithmID objc.ID, error_ID objc.ID) objc.ID {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("TKTokenSessionDelegate", "tokenSession:signData:usingKey:algorithm:error:")
					}
				}()
				session := TKTokenSessionFromID(sessionID)
				dataToSign := foundation.NSDataFromID(dataToSignID)
				algorithm := TKTokenKeyAlgorithmFromID(algorithmID)
				error_ := foundation.NSErrorFromID(error_ID)
				_delegateResult := fn(session, dataToSign, keyObjectID, algorithm, error_).GetID()
				_delegateDone = true
				return _delegateResult
			},
		})
	}

	if config.TokenSessionDecryptDataUsingKeyAlgorithmError != nil {
		fn := config.TokenSessionDecryptDataUsingKeyAlgorithmError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tokenSession:decryptData:usingKey:algorithm:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, ciphertextID objc.ID, keyObjectID TKTokenObjectID, algorithmID objc.ID, error_ID objc.ID) objc.ID {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("TKTokenSessionDelegate", "tokenSession:decryptData:usingKey:algorithm:error:")
					}
				}()
				session := TKTokenSessionFromID(sessionID)
				ciphertext := foundation.NSDataFromID(ciphertextID)
				algorithm := TKTokenKeyAlgorithmFromID(algorithmID)
				error_ := foundation.NSErrorFromID(error_ID)
				_delegateResult := fn(session, ciphertext, keyObjectID, algorithm, error_).GetID()
				_delegateDone = true
				return _delegateResult
			},
		})
	}

	if config.TokenSessionPerformKeyExchangeWithPublicKeyUsingKeyAlgorithmParametersError != nil {
		fn := config.TokenSessionPerformKeyExchangeWithPublicKeyUsingKeyAlgorithmParametersError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tokenSession:performKeyExchangeWithPublicKey:usingKey:algorithm:parameters:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, otherPartyPublicKeyDataID objc.ID, objectID TKTokenObjectID, algorithmID objc.ID, parametersID objc.ID, error_ID objc.ID) objc.ID {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("TKTokenSessionDelegate", "tokenSession:performKeyExchangeWithPublicKey:usingKey:algorithm:parameters:error:")
					}
				}()
				session := TKTokenSessionFromID(sessionID)
				otherPartyPublicKeyData := foundation.NSDataFromID(otherPartyPublicKeyDataID)
				algorithm := TKTokenKeyAlgorithmFromID(algorithmID)
				parameters := TKTokenKeyExchangeParametersFromID(parametersID)
				error_ := foundation.NSErrorFromID(error_ID)
				_delegateResult := fn(session, otherPartyPublicKeyData, objectID, algorithm, parameters, error_).GetID()
				_delegateDone = true
				return _delegateResult
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("TKTokenSessionDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewTKTokenSessionDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return TKTokenSessionDelegateObjectFromID(instance)
}
