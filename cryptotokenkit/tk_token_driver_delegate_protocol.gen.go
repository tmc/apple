// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// The interface that a token driver delegate implements to respond to token creation events.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriverDelegate
type TKTokenDriverDelegate interface {
	objectivec.IObject
}

// TKTokenDriverDelegateObject wraps an existing Objective-C object that conforms to the TKTokenDriverDelegate protocol.
type TKTokenDriverDelegateObject struct {
	objectivec.Object
}

func (o TKTokenDriverDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// TKTokenDriverDelegateObjectFromID constructs a [TKTokenDriverDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func TKTokenDriverDelegateObjectFromID(id objc.ID) TKTokenDriverDelegateObject {
	return TKTokenDriverDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate to terminate the token you specify.
//
// driver: The token driver.
//
// token: The token to be terminated.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriverDelegate/tokenDriver(_:terminateToken:)
func (o TKTokenDriverDelegateObject) TokenDriverTerminateToken(driver ITKTokenDriver, token ITKToken) {
	objc.Send[struct{}](o.ID, objc.Sel("tokenDriver:terminateToken:"), driver, token)
}

// Creates a new token for the configuration you specify.
//
// driver: The token driver.
//
// configuration: The configuration that identifies the token to create.
//
// # Return Value
//
// The created token.
//
// # Discussion
//
// The system invokes this method to request creation of a token instance,
// which the [TKTokenConfiguration.InstanceID] property of the configuration
// you specify identifies.
//
// The created token has access to its current configuration using the
// [TKTokenConfiguration.ConfigurationData] property, which can provide
// token-implementation-specific additional data. The token can access
// keychain items this token exports with the methods
// [TKTokenConfiguration.KeyForObjectIDError] and
// [TKTokenConfiguration.CertificateForObjectIDError] that the configuration
// provides.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriverDelegate/tokenDriver(_:tokenFor:)
func (o TKTokenDriverDelegateObject) TokenDriverTokenForConfigurationError(driver ITKTokenDriver, configuration ITKTokenConfiguration) (ITKToken, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("tokenDriver:tokenForConfiguration:error:"), driver, configuration)
	if err != nil {
		return nil, err
	}
	return TKTokenFromID(rv), nil
}

// TKTokenDriverDelegateConfig holds optional typed callbacks for [TKTokenDriverDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/cryptotokenkit/tktokendriverdelegate
type TKTokenDriverDelegateConfig struct {

	// Creating and Removing Tokens
	// TokenDriverTerminateToken — Tells the delegate to terminate the token you specify.
	TokenDriverTerminateToken func(driver TKTokenDriver, token TKToken)

	// Other Methods
	// TokenDriverTokenForConfigurationError — Creates a new token for the configuration you specify.
	TokenDriverTokenForConfigurationError func(driver TKTokenDriver, configuration TKTokenConfiguration, error_ foundation.NSError) TKToken
}

// NewTKTokenDriverDelegate creates an Objective-C object implementing the [TKTokenDriverDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [TKTokenDriverDelegateObject] satisfies the [TKTokenDriverDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/cryptotokenkit/tktokendriverdelegate
func NewTKTokenDriverDelegate(config TKTokenDriverDelegateConfig) TKTokenDriverDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoTKTokenDriverDelegate_%d", n)

	var methods []objc.MethodDef

	if config.TokenDriverTerminateToken != nil {
		fn := config.TokenDriverTerminateToken
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tokenDriver:terminateToken:"),
			Fn: func(self objc.ID, _cmd objc.SEL, driverID objc.ID, tokenID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("TKTokenDriverDelegate", "tokenDriver:terminateToken:")
					}
				}()
				driver := TKTokenDriverFromID(driverID)
				token := TKTokenFromID(tokenID)
				fn(driver, token)
				_delegateDone = true
			},
		})
	}

	if config.TokenDriverTokenForConfigurationError != nil {
		fn := config.TokenDriverTokenForConfigurationError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tokenDriver:tokenForConfiguration:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, driverID objc.ID, configurationID objc.ID, error_ID objc.ID) objc.ID {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("TKTokenDriverDelegate", "tokenDriver:tokenForConfiguration:error:")
					}
				}()
				driver := TKTokenDriverFromID(driverID)
				configuration := TKTokenConfigurationFromID(configurationID)
				error_ := foundation.NSErrorFromID(error_ID)
				_delegateResult := fn(driver, configuration, error_).GetID()
				_delegateDone = true
				return _delegateResult
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("TKTokenDriverDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewTKTokenDriverDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return TKTokenDriverDelegateObjectFromID(instance)
}
