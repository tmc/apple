// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKTokenKeychainContents] class.
var (
	_TKTokenKeychainContentsClass     TKTokenKeychainContentsClass
	_TKTokenKeychainContentsClassOnce sync.Once
)

func getTKTokenKeychainContentsClass() TKTokenKeychainContentsClass {
	_TKTokenKeychainContentsClassOnce.Do(func() {
		_TKTokenKeychainContentsClass = TKTokenKeychainContentsClass{class: objc.GetClass("TKTokenKeychainContents")}
	})
	return _TKTokenKeychainContentsClass
}

// GetTKTokenKeychainContentsClass returns the class object for TKTokenKeychainContents.
func GetTKTokenKeychainContentsClass() TKTokenKeychainContentsClass {
	return getTKTokenKeychainContentsClass()
}

type TKTokenKeychainContentsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKTokenKeychainContentsClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKTokenKeychainContentsClass) Alloc() TKTokenKeychainContents {
	rv := objc.Send[TKTokenKeychainContents](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// A representation of the state of the keychain for a particular token.
//
// # Adding Keychain Items
//
//   - [TKTokenKeychainContents.FillWithItems]: Fills the keychain with the specified items.
//
// # Accessing Keychain Items
//
//   - [TKTokenKeychainContents.Items]: Returns all items for token in the keychain.
//   - [TKTokenKeychainContents.KeyForObjectIDError]: Returns the key for a specified object identifier.
//   - [TKTokenKeychainContents.CertificateForObjectIDError]: Returns the key for a specified object identifier.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainContents
type TKTokenKeychainContents struct {
	objectivec.Object
}

// TKTokenKeychainContentsFromID constructs a [TKTokenKeychainContents] from an objc.ID.
//
// A representation of the state of the keychain for a particular token.
func TKTokenKeychainContentsFromID(id objc.ID) TKTokenKeychainContents {
	return TKTokenKeychainContents{objectivec.Object{ID: id}}
}

// NOTE: TKTokenKeychainContents adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKTokenKeychainContents] class.
//
// # Adding Keychain Items
//
//   - [ITKTokenKeychainContents.FillWithItems]: Fills the keychain with the specified items.
//
// # Accessing Keychain Items
//
//   - [ITKTokenKeychainContents.Items]: Returns all items for token in the keychain.
//   - [ITKTokenKeychainContents.KeyForObjectIDError]: Returns the key for a specified object identifier.
//   - [ITKTokenKeychainContents.CertificateForObjectIDError]: Returns the key for a specified object identifier.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainContents
type ITKTokenKeychainContents interface {
	objectivec.IObject

	// Topic: Adding Keychain Items

	// Fills the keychain with the specified items.
	FillWithItems(items []TKTokenKeychainItem)

	// Topic: Accessing Keychain Items

	// Returns all items for token in the keychain.
	Items() []TKTokenKeychainItem
	// Returns the key for a specified object identifier.
	KeyForObjectIDError(objectID TKTokenObjectID) (ITKTokenKeychainKey, error)
	// Returns the key for a specified object identifier.
	CertificateForObjectIDError(objectID TKTokenObjectID) (ITKTokenKeychainCertificate, error)
}

// Init initializes the instance.
func (t TKTokenKeychainContents) Init() TKTokenKeychainContents {
	rv := objc.Send[TKTokenKeychainContents](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKTokenKeychainContents) Autorelease() TKTokenKeychainContents {
	rv := objc.Send[TKTokenKeychainContents](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenKeychainContents creates a new TKTokenKeychainContents instance.
func NewTKTokenKeychainContents() TKTokenKeychainContents {
	class := getTKTokenKeychainContentsClass()
	rv := objc.Send[TKTokenKeychainContents](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Fills the keychain with the specified items.
//
// items: The items to be added to the keychain.
//
// # Discussion
//
// All existing items for the token are first removed from the keychain before
// filling the keychain with `items`.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainContents/fill(with:)
func (t TKTokenKeychainContents) FillWithItems(items []TKTokenKeychainItem) {
	objc.Send[objc.ID](t.ID, objc.Sel("fillWithItems:"), objectivec.IObjectSliceToNSArray(items))
}

// Returns the key for a specified object identifier.
//
// objectID: The object identifier for the keychain item.
//
// # Return Value
//
// The key, or `nil` if no key exists.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainContents/key(forObjectID:)
func (t TKTokenKeychainContents) KeyForObjectIDError(objectID TKTokenObjectID) (ITKTokenKeychainKey, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("keyForObjectID:error:"), objectID, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return TKTokenKeychainKey{}, foundation.NSErrorFrom(errorPtr)
	}
	return TKTokenKeychainKeyFromID(rv), nil

}

// Returns the key for a specified object identifier.
//
// objectID: The object identifier for the keychain item.
//
// # Return Value
//
// The certificate, or `nil` if no certificate exists.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainContents/certificate(forObjectID:)
func (t TKTokenKeychainContents) CertificateForObjectIDError(objectID TKTokenObjectID) (ITKTokenKeychainCertificate, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("certificateForObjectID:error:"), objectID, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return TKTokenKeychainCertificate{}, foundation.NSErrorFrom(errorPtr)
	}
	return TKTokenKeychainCertificateFromID(rv), nil

}

// Returns all items for token in the keychain.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainContents/items
func (t TKTokenKeychainContents) Items() []TKTokenKeychainItem {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("items"))
	return objc.ConvertSlice(rv, func(id objc.ID) TKTokenKeychainItem {
		return TKTokenKeychainItemFromID(id)
	})
}
