// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/security"
)

// The class instance for the [TKTokenKeychainKey] class.
var (
	_TKTokenKeychainKeyClass     TKTokenKeychainKeyClass
	_TKTokenKeychainKeyClassOnce sync.Once
)

func getTKTokenKeychainKeyClass() TKTokenKeychainKeyClass {
	_TKTokenKeychainKeyClassOnce.Do(func() {
		_TKTokenKeychainKeyClass = TKTokenKeychainKeyClass{class: objc.GetClass("TKTokenKeychainKey")}
	})
	return _TKTokenKeychainKeyClass
}

// GetTKTokenKeychainKeyClass returns the class object for TKTokenKeychainKey.
func GetTKTokenKeychainKeyClass() TKTokenKeychainKeyClass {
	return getTKTokenKeychainKeyClass()
}

type TKTokenKeychainKeyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKTokenKeychainKeyClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKTokenKeychainKeyClass) Alloc() TKTokenKeychainKey {
	rv := objc.Send[TKTokenKeychainKey](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// A token’s key as stored in the keychain.
//
// # Creating Token Keychain Keys
//
//   - [TKTokenKeychainKey.InitWithCertificateObjectID]: Initializes a token keychain key with data from the specified certificate reference and a given object ID.
//
// # Accessing Key Attributes
//
//   - [TKTokenKeychainKey.KeyType]: The type of the key. Currently, only [kSecAttrKeyTypeRSA](<https://developer.apple.com/documentation/Security/kSecAttrKeyTypeRSA>) and `kSecAttrKeyTypeECSECPrimeRandom` are supported values.
//   - [TKTokenKeychainKey.SetKeyType]
//   - [TKTokenKeychainKey.KeySizeInBits]
//   - [TKTokenKeychainKey.SetKeySizeInBits]
//   - [TKTokenKeychainKey.ApplicationTag]: The private tag data.
//   - [TKTokenKeychainKey.SetApplicationTag]
//   - [TKTokenKeychainKey.PublicKeyData]: The public key data.
//   - [TKTokenKeychainKey.SetPublicKeyData]
//   - [TKTokenKeychainKey.PublicKeyHash]: The SHA1 hash of the raw public key.
//   - [TKTokenKeychainKey.SetPublicKeyHash]
//   - [TKTokenKeychainKey.CanDecrypt]: Whether the key can be used to decrypt data.
//   - [TKTokenKeychainKey.SetCanDecrypt]
//   - [TKTokenKeychainKey.CanSign]: Whether the key can be used to sign data.
//   - [TKTokenKeychainKey.SetCanSign]
//   - [TKTokenKeychainKey.CanPerformKeyExchange]: Whether the key can be used to perform Diffie-Hellman style cryptographic key exchange.
//   - [TKTokenKeychainKey.SetCanPerformKeyExchange]
//   - [TKTokenKeychainKey.IsSuitableForLogin]: Whether the key can be used for system login.
//   - [TKTokenKeychainKey.SetSuitableForLogin]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey
type TKTokenKeychainKey struct {
	TKTokenKeychainItem
}

// TKTokenKeychainKeyFromID constructs a [TKTokenKeychainKey] from an objc.ID.
//
// A token’s key as stored in the keychain.
func TKTokenKeychainKeyFromID(id objc.ID) TKTokenKeychainKey {
	return TKTokenKeychainKey{TKTokenKeychainItem: TKTokenKeychainItemFromID(id)}
}

// NOTE: TKTokenKeychainKey adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKTokenKeychainKey] class.
//
// # Creating Token Keychain Keys
//
//   - [ITKTokenKeychainKey.InitWithCertificateObjectID]: Initializes a token keychain key with data from the specified certificate reference and a given object ID.
//
// # Accessing Key Attributes
//
//   - [ITKTokenKeychainKey.KeyType]: The type of the key. Currently, only [kSecAttrKeyTypeRSA](<https://developer.apple.com/documentation/Security/kSecAttrKeyTypeRSA>) and `kSecAttrKeyTypeECSECPrimeRandom` are supported values.
//   - [ITKTokenKeychainKey.SetKeyType]
//   - [ITKTokenKeychainKey.KeySizeInBits]
//   - [ITKTokenKeychainKey.SetKeySizeInBits]
//   - [ITKTokenKeychainKey.ApplicationTag]: The private tag data.
//   - [ITKTokenKeychainKey.SetApplicationTag]
//   - [ITKTokenKeychainKey.PublicKeyData]: The public key data.
//   - [ITKTokenKeychainKey.SetPublicKeyData]
//   - [ITKTokenKeychainKey.PublicKeyHash]: The SHA1 hash of the raw public key.
//   - [ITKTokenKeychainKey.SetPublicKeyHash]
//   - [ITKTokenKeychainKey.CanDecrypt]: Whether the key can be used to decrypt data.
//   - [ITKTokenKeychainKey.SetCanDecrypt]
//   - [ITKTokenKeychainKey.CanSign]: Whether the key can be used to sign data.
//   - [ITKTokenKeychainKey.SetCanSign]
//   - [ITKTokenKeychainKey.CanPerformKeyExchange]: Whether the key can be used to perform Diffie-Hellman style cryptographic key exchange.
//   - [ITKTokenKeychainKey.SetCanPerformKeyExchange]
//   - [ITKTokenKeychainKey.IsSuitableForLogin]: Whether the key can be used for system login.
//   - [ITKTokenKeychainKey.SetSuitableForLogin]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey
type ITKTokenKeychainKey interface {
	ITKTokenKeychainItem

	// Topic: Creating Token Keychain Keys

	// Initializes a token keychain key with data from the specified certificate reference and a given object ID.
	InitWithCertificateObjectID(certificateRef security.SecCertificateRef, objectID TKTokenObjectID) TKTokenKeychainKey

	// Topic: Accessing Key Attributes

	// The type of the key. Currently, only [kSecAttrKeyTypeRSA](<https://developer.apple.com/documentation/Security/kSecAttrKeyTypeRSA>) and `kSecAttrKeyTypeECSECPrimeRandom` are supported values.
	KeyType() string
	SetKeyType(value string)
	KeySizeInBits() int
	SetKeySizeInBits(value int)
	// The private tag data.
	ApplicationTag() foundation.NSData
	SetApplicationTag(value foundation.NSData)
	// The public key data.
	PublicKeyData() foundation.NSData
	SetPublicKeyData(value foundation.NSData)
	// The SHA1 hash of the raw public key.
	PublicKeyHash() foundation.NSData
	SetPublicKeyHash(value foundation.NSData)
	// Whether the key can be used to decrypt data.
	CanDecrypt() bool
	SetCanDecrypt(value bool)
	// Whether the key can be used to sign data.
	CanSign() bool
	SetCanSign(value bool)
	// Whether the key can be used to perform Diffie-Hellman style cryptographic key exchange.
	CanPerformKeyExchange() bool
	SetCanPerformKeyExchange(value bool)
	// Whether the key can be used for system login.
	IsSuitableForLogin() bool
	SetSuitableForLogin(value bool)
}

// Init initializes the instance.
func (t TKTokenKeychainKey) Init() TKTokenKeychainKey {
	rv := objc.Send[TKTokenKeychainKey](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKTokenKeychainKey) Autorelease() TKTokenKeychainKey {
	rv := objc.Send[TKTokenKeychainKey](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenKeychainKey creates a new TKTokenKeychainKey instance.
func NewTKTokenKeychainKey() TKTokenKeychainKey {
	class := getTKTokenKeychainKeyClass()
	rv := objc.Send[TKTokenKeychainKey](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a token keychain key with data from the specified certificate
// reference and a given object ID.
//
// certificateRef: The certificate reference.
//
// You can create a [SecCertificateRef] value from a data object using the
// [SecCertificateCreateWithData] function.
//
// objectID: The object ID.
//
// # Return Value
//
// A new token keychain certificate.
//
// # Discussion
//
// This is the designated initializer.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/init(certificate:objectID:)
func NewTKTokenKeychainKeyWithCertificateObjectID(certificateRef security.SecCertificateRef, objectID TKTokenObjectID) TKTokenKeychainKey {
	instance := getTKTokenKeychainKeyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCertificate:objectID:"), certificateRef, objectID)
	return TKTokenKeychainKeyFromID(rv)
}

// Initializes a token keychain item with the specified object ID.
//
// objectID: The object ID.
//
// # Return Value
//
// A new keychain item.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainItem/init(objectID:)
func NewTKTokenKeychainKeyWithObjectID(objectID TKTokenObjectID) TKTokenKeychainKey {
	instance := getTKTokenKeychainKeyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObjectID:"), objectID)
	return TKTokenKeychainKeyFromID(rv)
}

// Initializes a token keychain key with data from the specified certificate
// reference and a given object ID.
//
// certificateRef: The certificate reference.
//
// You can create a [SecCertificateRef] value from a data object using the
// [SecCertificateCreateWithData] function.
//
// objectID: The object ID.
//
// # Return Value
//
// A new token keychain certificate.
//
// # Discussion
//
// This is the designated initializer.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/init(certificate:objectID:)
func (t TKTokenKeychainKey) InitWithCertificateObjectID(certificateRef security.SecCertificateRef, objectID TKTokenObjectID) TKTokenKeychainKey {
	rv := objc.Send[TKTokenKeychainKey](t.ID, objc.Sel("initWithCertificate:objectID:"), certificateRef, objectID)
	return rv
}

// The type of the key. Currently, only [kSecAttrKeyTypeRSA] and
// `kSecAttrKeyTypeECSECPrimeRandom` are supported values.
//
// # Discussion
//
// This property is equivalent to the `kSecAttrKeyType` type attribute.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/keyType
//
// [kSecAttrKeyTypeRSA]: https://developer.apple.com/documentation/Security/kSecAttrKeyTypeRSA
func (t TKTokenKeychainKey) KeyType() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("keyType"))
	return foundation.NSStringFromID(rv).String()
}
func (t TKTokenKeychainKey) SetKeyType(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setKeyType:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/keySizeInBits
func (t TKTokenKeychainKey) KeySizeInBits() int {
	rv := objc.Send[int](t.ID, objc.Sel("keySizeInBits"))
	return rv
}
func (t TKTokenKeychainKey) SetKeySizeInBits(value int) {
	objc.Send[struct{}](t.ID, objc.Sel("setKeySizeInBits:"), value)
}

// The private tag data.
//
// # Discussion
//
// This property is equivalent to the `kSecAttrApplicationTag` type attribute.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/applicationTag
func (t TKTokenKeychainKey) ApplicationTag() foundation.NSData {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("applicationTag"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (t TKTokenKeychainKey) SetApplicationTag(value foundation.NSData) {
	objc.Send[struct{}](t.ID, objc.Sel("setApplicationTag:"), value)
}

// The public key data.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/publicKeyData
func (t TKTokenKeychainKey) PublicKeyData() foundation.NSData {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("publicKeyData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (t TKTokenKeychainKey) SetPublicKeyData(value foundation.NSData) {
	objc.Send[struct{}](t.ID, objc.Sel("setPublicKeyData:"), value)
}

// The SHA1 hash of the raw public key.
//
// # Discussion
//
// This property is equivalent to the `kSecAttrApplicationLabel` type
// attribute.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/publicKeyHash
func (t TKTokenKeychainKey) PublicKeyHash() foundation.NSData {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("publicKeyHash"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (t TKTokenKeychainKey) SetPublicKeyHash(value foundation.NSData) {
	objc.Send[struct{}](t.ID, objc.Sel("setPublicKeyHash:"), value)
}

// Whether the key can be used to decrypt data.
//
// # Discussion
//
// This property is equivalent to the `kSecAttrCanDecrypt` type attribute.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/canDecrypt
func (t TKTokenKeychainKey) CanDecrypt() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("canDecrypt"))
	return rv
}
func (t TKTokenKeychainKey) SetCanDecrypt(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setCanDecrypt:"), value)
}

// Whether the key can be used to sign data.
//
// # Discussion
//
// This property is equivalent to the `kSecAttrCanSign` type attribute.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/canSign
func (t TKTokenKeychainKey) CanSign() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("canSign"))
	return rv
}
func (t TKTokenKeychainKey) SetCanSign(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setCanSign:"), value)
}

// Whether the key can be used to perform Diffie-Hellman style cryptographic
// key exchange.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/canPerformKeyExchange
func (t TKTokenKeychainKey) CanPerformKeyExchange() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("canPerformKeyExchange"))
	return rv
}
func (t TKTokenKeychainKey) SetCanPerformKeyExchange(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setCanPerformKeyExchange:"), value)
}

// Whether the key can be used for system login.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/isSuitableForLogin
func (t TKTokenKeychainKey) IsSuitableForLogin() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isSuitableForLogin"))
	return rv
}
func (t TKTokenKeychainKey) SetSuitableForLogin(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setSuitableForLogin:"), value)
}
