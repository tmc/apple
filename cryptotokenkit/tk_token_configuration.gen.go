// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKTokenConfiguration] class.
var (
	_TKTokenConfigurationClass     TKTokenConfigurationClass
	_TKTokenConfigurationClassOnce sync.Once
)

func getTKTokenConfigurationClass() TKTokenConfigurationClass {
	_TKTokenConfigurationClassOnce.Do(func() {
		_TKTokenConfigurationClass = TKTokenConfigurationClass{class: objc.GetClass("TKTokenConfiguration")}
	})
	return _TKTokenConfigurationClass
}

// GetTKTokenConfigurationClass returns the class object for TKTokenConfiguration.
func GetTKTokenConfigurationClass() TKTokenConfigurationClass {
	return getTKTokenConfigurationClass()
}

type TKTokenConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKTokenConfigurationClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKTokenConfigurationClass) Alloc() TKTokenConfiguration {
	rv := objc.Send[TKTokenConfiguration](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// A token’s configuration.
//
// # Overview
//
// When you introduce a new [TKTokenConfiguration] into the system, it can
// inform the system about its identities, consisting of both private keys and
// certificates, which the [TKTokenConfiguration.KeychainItems] property
// provides. Use the [TKTokenConfiguration.ConfigurationData] property to set
// additional configuration data.
//
// You configure always-available tokens on a per-user basis. Although the
// token driver and the app hosting the token extension are shared across the
// system, the configuration for a token is stored individually for each user.
//
// # Reporting Configuration Information
//
//   - [TKTokenConfiguration.InstanceID]: The unique, persistent identifier of this token that the token implementation creates.
//   - [TKTokenConfiguration.ConfigurationData]: Additional configuration information for the token instance.
//   - [TKTokenConfiguration.SetConfigurationData]
//
// # Retrieving Keys and Certificates
//
//   - [TKTokenConfiguration.KeychainItems]: The keychain items associated with this token.
//   - [TKTokenConfiguration.SetKeychainItems]
//   - [TKTokenConfiguration.CertificateForObjectIDError]: Returns a certificate from the keychain with the object identifier you specify.
//   - [TKTokenConfiguration.KeyForObjectIDError]: Returns a key from the keychain with the object identifier you specify.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/Configuration-swift.class
type TKTokenConfiguration struct {
	objectivec.Object
}

// TKTokenConfigurationFromID constructs a [TKTokenConfiguration] from an objc.ID.
//
// A token’s configuration.
func TKTokenConfigurationFromID(id objc.ID) TKTokenConfiguration {
	return TKTokenConfiguration{objectivec.Object{ID: id}}
}

// NOTE: TKTokenConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKTokenConfiguration] class.
//
// # Reporting Configuration Information
//
//   - [ITKTokenConfiguration.InstanceID]: The unique, persistent identifier of this token that the token implementation creates.
//   - [ITKTokenConfiguration.ConfigurationData]: Additional configuration information for the token instance.
//   - [ITKTokenConfiguration.SetConfigurationData]
//
// # Retrieving Keys and Certificates
//
//   - [ITKTokenConfiguration.KeychainItems]: The keychain items associated with this token.
//   - [ITKTokenConfiguration.SetKeychainItems]
//   - [ITKTokenConfiguration.CertificateForObjectIDError]: Returns a certificate from the keychain with the object identifier you specify.
//   - [ITKTokenConfiguration.KeyForObjectIDError]: Returns a key from the keychain with the object identifier you specify.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/Configuration-swift.class
type ITKTokenConfiguration interface {
	objectivec.IObject

	// Topic: Reporting Configuration Information

	// The unique, persistent identifier of this token that the token implementation creates.
	InstanceID() TKTokenInstanceID
	// Additional configuration information for the token instance.
	ConfigurationData() foundation.NSData
	SetConfigurationData(value foundation.NSData)

	// Topic: Retrieving Keys and Certificates

	// The keychain items associated with this token.
	KeychainItems() []TKTokenKeychainItem
	SetKeychainItems(value []TKTokenKeychainItem)
	// Returns a certificate from the keychain with the object identifier you specify.
	CertificateForObjectIDError(objectID TKTokenObjectID) (ITKTokenKeychainCertificate, error)
	// Returns a key from the keychain with the object identifier you specify.
	KeyForObjectIDError(objectID TKTokenObjectID) (ITKTokenKeychainKey, error)
}

// Init initializes the instance.
func (t TKTokenConfiguration) Init() TKTokenConfiguration {
	rv := objc.Send[TKTokenConfiguration](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKTokenConfiguration) Autorelease() TKTokenConfiguration {
	rv := objc.Send[TKTokenConfiguration](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenConfiguration creates a new TKTokenConfiguration instance.
func NewTKTokenConfiguration() TKTokenConfiguration {
	class := getTKTokenConfigurationClass()
	rv := objc.Send[TKTokenConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a certificate from the keychain with the object identifier you
// specify.
//
// objectID: The identifier for the certificate within the keychain.
//
// # Return Value
//
// The certificate that the keychain stores.
//
// # Discussion
//
// If the certificate the `objectID` specifies isn’t found, the system fills
// `error` with [TKError.Code.objectNotFound].
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/Configuration-swift.class/certificate(for:)
//
// [TKError.Code.objectNotFound]: https://developer.apple.com/documentation/CryptoTokenKit/TKError/Code/objectNotFound
func (t TKTokenConfiguration) CertificateForObjectIDError(objectID TKTokenObjectID) (ITKTokenKeychainCertificate, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("certificateForObjectID:error:"), objectID, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return TKTokenKeychainCertificate{}, foundation.NSErrorFrom(errorPtr)
	}
	return TKTokenKeychainCertificateFromID(rv), nil

}

// Returns a key from the keychain with the object identifier you specify.
//
// objectID: The identifier for the key within the keychain.
//
// # Return Value
//
// The key that the keychain stores.
//
// # Discussion
//
// If the key the `objectID` specifies isn’t found, the system fills the
// `error` with [TKError.Code.objectNotFound].
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/Configuration-swift.class/key(for:)
//
// [TKError.Code.objectNotFound]: https://developer.apple.com/documentation/CryptoTokenKit/TKError/Code/objectNotFound
func (t TKTokenConfiguration) KeyForObjectIDError(objectID TKTokenObjectID) (ITKTokenKeychainKey, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("keyForObjectID:error:"), objectID, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return TKTokenKeychainKey{}, foundation.NSErrorFrom(errorPtr)
	}
	return TKTokenKeychainKeyFromID(rv), nil

}

// The unique, persistent identifier of this token that the token
// implementation creates.
//
// # Discussion
//
// The instance identifier often represents some kind of serial number of the
// target hardware.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/Configuration-swift.class/instanceID
func (t TKTokenConfiguration) InstanceID() TKTokenInstanceID {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("instanceID"))
	return TKTokenInstanceID(foundation.NSStringFromID(rv).String())
}

// Additional configuration information for the token instance.
//
// # Discussion
//
// [TKTokenConfiguration.ConfigurationData] can provide
// token-implementation-specific additional data, which the app that hosts the
// token driver extension and configures the token provides. The system
// doesn’t interpret this data in any way.
//
// For example, the network-based hardware security module (HSM) can store
// encoded target network addresses, access credentials, or the list of
// identities the HSM contains.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/Configuration-swift.class/configurationData
func (t TKTokenConfiguration) ConfigurationData() foundation.NSData {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("configurationData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (t TKTokenConfiguration) SetConfigurationData(value foundation.NSData) {
	objc.Send[struct{}](t.ID, objc.Sel("setConfigurationData:"), value)
}

// The keychain items associated with this token.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/Configuration-swift.class/keychainItems
func (t TKTokenConfiguration) KeychainItems() []TKTokenKeychainItem {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("keychainItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) TKTokenKeychainItem {
		return TKTokenKeychainItemFromID(id)
	})
}
func (t TKTokenConfiguration) SetKeychainItems(value []TKTokenKeychainItem) {
	objc.Send[struct{}](t.ID, objc.Sel("setKeychainItems:"), objectivec.IObjectSliceToNSArray(value))
}
