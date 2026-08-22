// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// The interface that a smart card token driver delegate implements to respond to token creation events.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenDriverDelegate
type TKSmartCardTokenDriverDelegate interface {
	objectivec.IObject
	TKTokenDriverDelegate

	// Tells the delegate that a new Smart Card is detected.
	//
	// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenDriverDelegate/tokenDriver(_:createTokenFor:aid:)
	TokenDriverCreateTokenForSmartCardAIDError(driver ITKSmartCardTokenDriver, smartCard ITKSmartCard, AID foundation.NSData) (ITKSmartCardToken, error)
}

// TKSmartCardTokenDriverDelegateObject wraps an existing Objective-C object that conforms to the TKSmartCardTokenDriverDelegate protocol.
type TKSmartCardTokenDriverDelegateObject struct {
	objectivec.Object
}

func (o TKSmartCardTokenDriverDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// TKSmartCardTokenDriverDelegateObjectFromID constructs a [TKSmartCardTokenDriverDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func TKSmartCardTokenDriverDelegateObjectFromID(id objc.ID) TKSmartCardTokenDriverDelegateObject {
	return TKSmartCardTokenDriverDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that a new Smart Card is detected.
//
// driver: The Smart Card token driver.
//
// smartCard: The detected Smart Card.
//
// AID: The ISO 7816-4 application identifier that is selected on the Smart Card.
// If the `com.AppleXCUIElementTypeCtkXCUIElementTypeAid()` attributes is not
// present in the Smart Card token extension property list, no application is
// selected.
//
// # Return Value
//
// The token created for the Smart Card, or `nil` if an error occurs or the
// delegate decides not to provide a token.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenDriverDelegate/tokenDriver(_:createTokenFor:aid:)
func (o TKSmartCardTokenDriverDelegateObject) TokenDriverCreateTokenForSmartCardAIDError(driver ITKSmartCardTokenDriver, smartCard ITKSmartCard, AID foundation.NSData) (ITKSmartCardToken, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("tokenDriver:createTokenForSmartCard:AID:error:"), driver, smartCard, AID)
	if err != nil {
		return nil, err
	}
	return TKSmartCardTokenFromID(rv), nil
}

// Tells the delegate to terminate the token you specify.
//
// driver: The token driver.
//
// token: The token to be terminated.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriverDelegate/tokenDriver(_:terminateToken:)
func (o TKSmartCardTokenDriverDelegateObject) TokenDriverTerminateToken(driver ITKTokenDriver, token ITKToken) {
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
func (o TKSmartCardTokenDriverDelegateObject) TokenDriverTokenForConfigurationError(driver ITKTokenDriver, configuration ITKTokenConfiguration) (ITKToken, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("tokenDriver:tokenForConfiguration:error:"), driver, configuration)
	if err != nil {
		return nil, err
	}
	return TKTokenFromID(rv), nil
}

// TKSmartCardTokenDriverDelegateConfig holds optional typed callbacks for [TKSmartCardTokenDriverDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcardtokendriverdelegate
type TKSmartCardTokenDriverDelegateConfig struct {

	// Other Methods
	// TokenDriverCreateTokenForSmartCardAIDError — Tells the delegate that a new Smart Card is detected.
	TokenDriverCreateTokenForSmartCardAIDError func(driver TKSmartCardTokenDriver, smartCard TKSmartCard, AID foundation.NSData, error_ foundation.NSError) TKSmartCardToken
}

// NewTKSmartCardTokenDriverDelegate creates an Objective-C object implementing the [TKSmartCardTokenDriverDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [TKSmartCardTokenDriverDelegateObject] satisfies the [TKSmartCardTokenDriverDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcardtokendriverdelegate
func NewTKSmartCardTokenDriverDelegate(config TKSmartCardTokenDriverDelegateConfig) TKSmartCardTokenDriverDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoTKSmartCardTokenDriverDelegate_%d", n)

	var methods []objc.MethodDef

	if config.TokenDriverCreateTokenForSmartCardAIDError != nil {
		fn := config.TokenDriverCreateTokenForSmartCardAIDError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tokenDriver:createTokenForSmartCard:AID:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, driverID objc.ID, smartCardID objc.ID, AIDID objc.ID, error_ID objc.ID) objc.ID {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("TKSmartCardTokenDriverDelegate", "tokenDriver:createTokenForSmartCard:AID:error:")
					}
				}()
				driver := TKSmartCardTokenDriverFromID(driverID)
				smartCard := TKSmartCardFromID(smartCardID)
				AID := foundation.NSDataFromID(AIDID)
				error_ := foundation.NSErrorFromID(error_ID)
				_delegateResult := fn(driver, smartCard, AID, error_).GetID()
				_delegateDone = true
				return _delegateResult
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("TKSmartCardTokenDriverDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewTKSmartCardTokenDriverDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return TKSmartCardTokenDriverDelegateObjectFromID(instance)
}
