// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKToken] class.
var (
	_TKTokenClass     TKTokenClass
	_TKTokenClassOnce sync.Once
)

func getTKTokenClass() TKTokenClass {
	_TKTokenClassOnce.Do(func() {
		_TKTokenClass = TKTokenClass{class: objc.GetClass("TKToken")}
	})
	return _TKTokenClass
}

// GetTKTokenClass returns the class object for TKToken.
func GetTKTokenClass() TKTokenClass {
	return getTKTokenClass()
}

type TKTokenClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKTokenClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKTokenClass) Alloc() TKToken {
	rv := objc.Send[TKToken](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a hardware-based cryptographic token.
//
// # Overview
//
// # Creating Tokens
//
//   - [TKToken.InitWithTokenDriverInstanceID]: Initializes a token with the driver you specify.
//
// # Responding to Session Creation
//
//   - [TKToken.Delegate]: The token delegate.
//   - [TKToken.SetDelegate]
//
// # Accessing the Driver
//
//   - [TKToken.TokenDriver]: The token driver.
//
// # Accessing Keychain Items
//
//   - [TKToken.KeychainContents]: The contents of the keychain for this token.
//
// # Configuring the Token
//
//   - [TKToken.Configuration]: The current configuration for a token.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKToken
type TKToken struct {
	objectivec.Object
}

// TKTokenFromID constructs a [TKToken] from an objc.ID.
//
// A representation of a hardware-based cryptographic token.
func TKTokenFromID(id objc.ID) TKToken {
	return TKToken{objectivec.Object{ID: id}}
}

// NOTE: TKToken adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKToken] class.
//
// # Creating Tokens
//
//   - [ITKToken.InitWithTokenDriverInstanceID]: Initializes a token with the driver you specify.
//
// # Responding to Session Creation
//
//   - [ITKToken.Delegate]: The token delegate.
//   - [ITKToken.SetDelegate]
//
// # Accessing the Driver
//
//   - [ITKToken.TokenDriver]: The token driver.
//
// # Accessing Keychain Items
//
//   - [ITKToken.KeychainContents]: The contents of the keychain for this token.
//
// # Configuring the Token
//
//   - [ITKToken.Configuration]: The current configuration for a token.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKToken
type ITKToken interface {
	objectivec.IObject

	// Topic: Creating Tokens

	// Initializes a token with the driver you specify.
	InitWithTokenDriverInstanceID(tokenDriver ITKTokenDriver, instanceID TKTokenInstanceID) TKToken

	// Topic: Responding to Session Creation

	// The token delegate.
	Delegate() TKTokenDelegate
	SetDelegate(value TKTokenDelegate)

	// Topic: Accessing the Driver

	// The token driver.
	TokenDriver() ITKTokenDriver

	// Topic: Accessing Keychain Items

	// The contents of the keychain for this token.
	KeychainContents() ITKTokenKeychainContents

	// Topic: Configuring the Token

	// The current configuration for a token.
	Configuration() ITKTokenConfiguration
}

// Init initializes the instance.
func (t TKToken) Init() TKToken {
	rv := objc.Send[TKToken](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKToken) Autorelease() TKToken {
	rv := objc.Send[TKToken](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKToken creates a new TKToken instance.
func NewTKToken() TKToken {
	class := getTKTokenClass()
	rv := objc.Send[TKToken](objc.ID(class.class), objc.Sel("new"))
	return rv
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
func NewTKTokenWithTokenDriverInstanceID(tokenDriver ITKTokenDriver, instanceID TKTokenInstanceID) TKToken {
	instance := getTKTokenClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTokenDriver:instanceID:"), tokenDriver, objc.String(string(instanceID)))
	return TKTokenFromID(rv)
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
func (t TKToken) InitWithTokenDriverInstanceID(tokenDriver ITKTokenDriver, instanceID TKTokenInstanceID) TKToken {
	rv := objc.Send[TKToken](t.ID, objc.Sel("initWithTokenDriver:instanceID:"), tokenDriver, objc.String(string(instanceID)))
	return rv
}

// The token delegate.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/delegate
func (t TKToken) Delegate() TKTokenDelegate {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("delegate"))
	return TKTokenDelegateObjectFromID(rv)
}
func (t TKToken) SetDelegate(value TKTokenDelegate) {
	objc.Send[struct{}](t.ID, objc.Sel("setDelegate:"), value)
}

// The token driver.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/tokenDriver
func (t TKToken) TokenDriver() ITKTokenDriver {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tokenDriver"))
	return TKTokenDriverFromID(objc.ID(rv))
}

// The contents of the keychain for this token.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/keychainContents
func (t TKToken) KeychainContents() ITKTokenKeychainContents {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("keychainContents"))
	return TKTokenKeychainContentsFromID(objc.ID(rv))
}

// The current configuration for a token.
//
// # Discussion
//
// Access keychain items exported by this token with the methods
// [TKTokenConfiguration.KeyForObjectIDError] and
// [TKTokenConfiguration.CertificateForObjectIDError] provided by the
// configuration. You can access token-implementation-specific additional data
// using the [TKTokenConfiguration.ConfigurationData] property of the
// configuration.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/configuration-swift.property
func (t TKToken) Configuration() ITKTokenConfiguration {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("configuration"))
	return TKTokenConfigurationFromID(objc.ID(rv))
}
