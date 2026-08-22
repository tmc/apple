// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/security"
)

// The class instance for the [TKTokenKeychainCertificate] class.
var (
	_TKTokenKeychainCertificateClass     TKTokenKeychainCertificateClass
	_TKTokenKeychainCertificateClassOnce sync.Once
)

func getTKTokenKeychainCertificateClass() TKTokenKeychainCertificateClass {
	_TKTokenKeychainCertificateClassOnce.Do(func() {
		_TKTokenKeychainCertificateClass = TKTokenKeychainCertificateClass{class: objc.GetClass("TKTokenKeychainCertificate")}
	})
	return _TKTokenKeychainCertificateClass
}

// GetTKTokenKeychainCertificateClass returns the class object for TKTokenKeychainCertificate.
func GetTKTokenKeychainCertificateClass() TKTokenKeychainCertificateClass {
	return getTKTokenKeychainCertificateClass()
}

type TKTokenKeychainCertificateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKTokenKeychainCertificateClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKTokenKeychainCertificateClass) Alloc() TKTokenKeychainCertificate {
	rv := objc.Send[TKTokenKeychainCertificate](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// A token’s certificate as stored in the keychain.
//
// # Creating Token Keychain Certificates
//
//   - [TKTokenKeychainCertificate.InitWithCertificateObjectID]: Initializes a token keychain certificate with data from the specified certificate reference and a given object ID.
//
// # Accessing Certificate Data
//
//   - [TKTokenKeychainCertificate.Data]: Returns a DER-encoded representation of an X.509 certificate.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainCertificate
type TKTokenKeychainCertificate struct {
	TKTokenKeychainItem
}

// TKTokenKeychainCertificateFromID constructs a [TKTokenKeychainCertificate] from an objc.ID.
//
// A token’s certificate as stored in the keychain.
func TKTokenKeychainCertificateFromID(id objc.ID) TKTokenKeychainCertificate {
	return TKTokenKeychainCertificate{TKTokenKeychainItem: TKTokenKeychainItemFromID(id)}
}

// NOTE: TKTokenKeychainCertificate adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKTokenKeychainCertificate] class.
//
// # Creating Token Keychain Certificates
//
//   - [ITKTokenKeychainCertificate.InitWithCertificateObjectID]: Initializes a token keychain certificate with data from the specified certificate reference and a given object ID.
//
// # Accessing Certificate Data
//
//   - [ITKTokenKeychainCertificate.Data]: Returns a DER-encoded representation of an X.509 certificate.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainCertificate
type ITKTokenKeychainCertificate interface {
	ITKTokenKeychainItem

	// Topic: Creating Token Keychain Certificates

	// Initializes a token keychain certificate with data from the specified certificate reference and a given object ID.
	InitWithCertificateObjectID(certificateRef security.SecCertificateRef, objectID TKTokenObjectID) TKTokenKeychainCertificate

	// Topic: Accessing Certificate Data

	// Returns a DER-encoded representation of an X.509 certificate.
	Data() foundation.NSData
}

// Init initializes the instance.
func (t TKTokenKeychainCertificate) Init() TKTokenKeychainCertificate {
	rv := objc.Send[TKTokenKeychainCertificate](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKTokenKeychainCertificate) Autorelease() TKTokenKeychainCertificate {
	rv := objc.Send[TKTokenKeychainCertificate](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenKeychainCertificate creates a new TKTokenKeychainCertificate instance.
func NewTKTokenKeychainCertificate() TKTokenKeychainCertificate {
	class := getTKTokenKeychainCertificateClass()
	rv := objc.Send[TKTokenKeychainCertificate](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a token keychain certificate with data from the specified
// certificate reference and a given object ID.
//
// certificateRef: The certificate reference.
//
// You can create a [SecCertificate] value from a data object using the
// [SecCertificateCreateWithData(_:_:)] function.
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
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainCertificate/init(certificate:objectID:)
//
// [SecCertificateCreateWithData(_:_:)]: https://developer.apple.com/documentation/Security/SecCertificateCreateWithData(_:_:)
// [SecCertificate]: https://developer.apple.com/documentation/Security/SecCertificate
func NewTKTokenKeychainCertificateWithCertificateObjectID(certificateRef security.SecCertificateRef, objectID TKTokenObjectID) TKTokenKeychainCertificate {
	instance := getTKTokenKeychainCertificateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCertificate:objectID:"), certificateRef, objectID)
	return TKTokenKeychainCertificateFromID(rv)
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
func NewTKTokenKeychainCertificateWithObjectID(objectID TKTokenObjectID) TKTokenKeychainCertificate {
	instance := getTKTokenKeychainCertificateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObjectID:"), objectID)
	return TKTokenKeychainCertificateFromID(rv)
}

// Initializes a token keychain certificate with data from the specified
// certificate reference and a given object ID.
//
// certificateRef: The certificate reference.
//
// You can create a [SecCertificate] value from a data object using the
// [SecCertificateCreateWithData(_:_:)] function.
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
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainCertificate/init(certificate:objectID:)
//
// [SecCertificateCreateWithData(_:_:)]: https://developer.apple.com/documentation/Security/SecCertificateCreateWithData(_:_:)
// [SecCertificate]: https://developer.apple.com/documentation/Security/SecCertificate
func (t TKTokenKeychainCertificate) InitWithCertificateObjectID(certificateRef security.SecCertificateRef, objectID TKTokenObjectID) TKTokenKeychainCertificate {
	rv := objc.Send[TKTokenKeychainCertificate](t.ID, objc.Sel("initWithCertificate:objectID:"), certificateRef, objectID)
	return rv
}

// Returns a DER-encoded representation of an X.509 certificate.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainCertificate/data
func (t TKTokenKeychainCertificate) Data() foundation.NSData {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("data"))
	return foundation.NSDataFromID(objc.ID(rv))
}
