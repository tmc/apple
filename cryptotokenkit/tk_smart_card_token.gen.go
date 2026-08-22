// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [TKSmartCardToken] class.
var (
	_TKSmartCardTokenClass     TKSmartCardTokenClass
	_TKSmartCardTokenClassOnce sync.Once
)

func getTKSmartCardTokenClass() TKSmartCardTokenClass {
	_TKSmartCardTokenClassOnce.Do(func() {
		_TKSmartCardTokenClass = TKSmartCardTokenClass{class: objc.GetClass("TKSmartCardToken")}
	})
	return _TKSmartCardTokenClass
}

// GetTKSmartCardTokenClass returns the class object for TKSmartCardToken.
func GetTKSmartCardTokenClass() TKSmartCardTokenClass {
	return getTKSmartCardTokenClass()
}

type TKSmartCardTokenClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKSmartCardTokenClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKSmartCardTokenClass) Alloc() TKSmartCardToken {
	rv := objc.Send[TKSmartCardToken](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a smart card based cryptographic token.
//
// # Creating Smart Card Tokens
//
//   - [TKSmartCardToken.InitWithSmartCardAIDInstanceIDTokenDriver]: Initializes a smart card token with the specified smart card, application identifier, and token driver.
//
// # Accessing the Application Identifier
//
//   - [TKSmartCardToken.AID]: The ISO 7816-4 application identifiers of the Smart Card.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardToken
type TKSmartCardToken struct {
	TKToken
}

// TKSmartCardTokenFromID constructs a [TKSmartCardToken] from an objc.ID.
//
// A representation of a smart card based cryptographic token.
func TKSmartCardTokenFromID(id objc.ID) TKSmartCardToken {
	return TKSmartCardToken{TKToken: TKTokenFromID(id)}
}

// NOTE: TKSmartCardToken adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKSmartCardToken] class.
//
// # Creating Smart Card Tokens
//
//   - [ITKSmartCardToken.InitWithSmartCardAIDInstanceIDTokenDriver]: Initializes a smart card token with the specified smart card, application identifier, and token driver.
//
// # Accessing the Application Identifier
//
//   - [ITKSmartCardToken.AID]: The ISO 7816-4 application identifiers of the Smart Card.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardToken
type ITKSmartCardToken interface {
	ITKToken

	// Topic: Creating Smart Card Tokens

	// Initializes a smart card token with the specified smart card, application identifier, and token driver.
	InitWithSmartCardAIDInstanceIDTokenDriver(smartCard ITKSmartCard, AID foundation.NSData, instanceID string, tokenDriver ITKSmartCardTokenDriver) TKSmartCardToken

	// Topic: Accessing the Application Identifier

	// The ISO 7816-4 application identifiers of the Smart Card.
	AID() foundation.NSData
}

// Init initializes the instance.
func (t TKSmartCardToken) Init() TKSmartCardToken {
	rv := objc.Send[TKSmartCardToken](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKSmartCardToken) Autorelease() TKSmartCardToken {
	rv := objc.Send[TKSmartCardToken](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardToken creates a new TKSmartCardToken instance.
func NewTKSmartCardToken() TKSmartCardToken {
	class := getTKSmartCardTokenClass()
	rv := objc.Send[TKSmartCardToken](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a smart card token with the specified smart card, application
// identifier, and token driver.
//
// smartCard: The smart card on which the created token should operate.
//
// AID: The ISO 7816-4 application identifier for the smart card.
//
// instanceID: A unique, persistent identifier for this token. This value is typically
// generated from the serial number of the target hardware.
//
// tokenDriver: The driver associated with the created token.
//
// # Discussion
//
// This is the designated initializer.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardToken/init(smartCard:aid:instanceID:tokenDriver:)
func NewTKSmartCardTokenWithSmartCardAIDInstanceIDTokenDriver(smartCard ITKSmartCard, AID foundation.NSData, instanceID string, tokenDriver ITKSmartCardTokenDriver) TKSmartCardToken {
	instance := getTKSmartCardTokenClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSmartCard:AID:instanceID:tokenDriver:"), smartCard, AID, objc.String(instanceID), tokenDriver)
	return TKSmartCardTokenFromID(rv)
}

// Initializes a token with the driver you specify.
//
// tokenDriver: The driver of the token.
//
// instanceID: A unique, persistent identifier for this token. This value is typically
// generated from the serial number of the target hardware.
//
// # Return Value
//
// A new token object.
//
// # Discussion
//
// This is the designated initializer.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/init(tokenDriver:instanceID:)
func NewTKSmartCardTokenWithTokenDriverInstanceID(tokenDriver ITKTokenDriver, instanceID TKTokenInstanceID) TKSmartCardToken {
	instance := getTKSmartCardTokenClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTokenDriver:instanceID:"), tokenDriver, objc.String(string(instanceID)))
	return TKSmartCardTokenFromID(rv)
}

// Initializes a smart card token with the specified smart card, application
// identifier, and token driver.
//
// smartCard: The smart card on which the created token should operate.
//
// AID: The ISO 7816-4 application identifier for the smart card.
//
// instanceID: A unique, persistent identifier for this token. This value is typically
// generated from the serial number of the target hardware.
//
// tokenDriver: The driver associated with the created token.
//
// # Discussion
//
// This is the designated initializer.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardToken/init(smartCard:aid:instanceID:tokenDriver:)
func (t TKSmartCardToken) InitWithSmartCardAIDInstanceIDTokenDriver(smartCard ITKSmartCard, AID foundation.NSData, instanceID string, tokenDriver ITKSmartCardTokenDriver) TKSmartCardToken {
	rv := objc.Send[TKSmartCardToken](t.ID, objc.Sel("initWithSmartCard:AID:instanceID:tokenDriver:"), smartCard, AID, objc.String(instanceID), tokenDriver)
	return rv
}

// The ISO 7816-4 application identifiers of the Smart Card.
//
// # Discussion
//
// This value is specified in the Smart Card token extension’s
// [NSExtensionAttributes] property list by the
// `com.AppleXCUIElementTypeCtkXCUIElementTypeAid()` attribute. If this
// attribute specifies multiple AIDs, this parameter represents the
// application identifier found on the card that is already preselected. If
// the `com.AppleXCUIElementTypeCtkXCUIElementTypeAid()` attribute is not
// present, no application is automatically preselected and the value of this
// property is `nil`.
//
// For more information, see [App Extension Keys] in [Information Property
// List Key Reference].
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardToken/aid
//
// [App Extension Keys]: https://developer.apple.com/library/archive/documentation/General/Reference/InfoPlistKeyReference/Articles/AppExtensionKeys.html#//apple_ref/doc/uid/TP40014212
// [Information Property List Key Reference]: https://developer.apple.com/library/archive/documentation/General/Reference/InfoPlistKeyReference/Introduction/Introduction.html#//apple_ref/doc/uid/TP40009247
func (t TKSmartCardToken) AID() foundation.NSData {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("AID"))
	return foundation.NSDataFromID(objc.ID(rv))
}
