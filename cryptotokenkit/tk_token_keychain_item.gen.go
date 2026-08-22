// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKTokenKeychainItem] class.
var (
	_TKTokenKeychainItemClass     TKTokenKeychainItemClass
	_TKTokenKeychainItemClassOnce sync.Once
)

func getTKTokenKeychainItemClass() TKTokenKeychainItemClass {
	_TKTokenKeychainItemClassOnce.Do(func() {
		_TKTokenKeychainItemClass = TKTokenKeychainItemClass{class: objc.GetClass("TKTokenKeychainItem")}
	})
	return _TKTokenKeychainItemClass
}

// GetTKTokenKeychainItemClass returns the class object for TKTokenKeychainItem.
func GetTKTokenKeychainItemClass() TKTokenKeychainItemClass {
	return getTKTokenKeychainItemClass()
}

type TKTokenKeychainItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKTokenKeychainItemClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKTokenKeychainItemClass) Alloc() TKTokenKeychainItem {
	rv := objc.Send[TKTokenKeychainItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// An abstract base class for managing a token’s contents as keychain items.
//
// # Overview
//
// Don’t use this base class directly. Instead, use one of its subclasses,
// such as [TKTokenKeychainCertificate] for managing certificates or
// [TKTokenKeychainKey] for managing cryptographic keys.
//
// # Creating Token Keychain Items
//
//   - [TKTokenKeychainItem.InitWithObjectID]: Initializes a token keychain item with the specified object ID.
//
// # Accessing Keychain Item Attributes
//
//   - [TKTokenKeychainItem.ObjectID]: Returns the object ID used for keychain item identification.
//   - [TKTokenKeychainItem.Label]: The user-visible label for the keychain item.
//   - [TKTokenKeychainItem.SetLabel]
//   - [TKTokenKeychainItem.Constraints]: Access constraints for the keychain item, keyed by [TKTokenOperation](<https://developer.apple.com/documentation/CryptoTokenKit/TKTokenOperation>) values wrapped in [NSNumber](<https://developer.apple.com/documentation/Foundation/NSNumber>) objects.
//   - [TKTokenKeychainItem.SetConstraints]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainItem
type TKTokenKeychainItem struct {
	objectivec.Object
}

// TKTokenKeychainItemFromID constructs a [TKTokenKeychainItem] from an objc.ID.
//
// An abstract base class for managing a token’s contents as keychain items.
func TKTokenKeychainItemFromID(id objc.ID) TKTokenKeychainItem {
	return TKTokenKeychainItem{objectivec.Object{ID: id}}
}

// NOTE: TKTokenKeychainItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKTokenKeychainItem] class.
//
// # Creating Token Keychain Items
//
//   - [ITKTokenKeychainItem.InitWithObjectID]: Initializes a token keychain item with the specified object ID.
//
// # Accessing Keychain Item Attributes
//
//   - [ITKTokenKeychainItem.ObjectID]: Returns the object ID used for keychain item identification.
//   - [ITKTokenKeychainItem.Label]: The user-visible label for the keychain item.
//   - [ITKTokenKeychainItem.SetLabel]
//   - [ITKTokenKeychainItem.Constraints]: Access constraints for the keychain item, keyed by [TKTokenOperation](<https://developer.apple.com/documentation/CryptoTokenKit/TKTokenOperation>) values wrapped in [NSNumber](<https://developer.apple.com/documentation/Foundation/NSNumber>) objects.
//   - [ITKTokenKeychainItem.SetConstraints]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainItem
type ITKTokenKeychainItem interface {
	objectivec.IObject

	// Topic: Creating Token Keychain Items

	// Initializes a token keychain item with the specified object ID.
	InitWithObjectID(objectID TKTokenObjectID) TKTokenKeychainItem

	// Topic: Accessing Keychain Item Attributes

	// Returns the object ID used for keychain item identification.
	ObjectID() TKTokenObjectID
	// The user-visible label for the keychain item.
	Label() string
	SetLabel(value string)
	// Access constraints for the keychain item, keyed by [TKTokenOperation](<https://developer.apple.com/documentation/CryptoTokenKit/TKTokenOperation>) values wrapped in [NSNumber](<https://developer.apple.com/documentation/Foundation/NSNumber>) objects.
	Constraints() foundation.INSDictionary
	SetConstraints(value foundation.INSDictionary)
}

// Init initializes the instance.
func (t TKTokenKeychainItem) Init() TKTokenKeychainItem {
	rv := objc.Send[TKTokenKeychainItem](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKTokenKeychainItem) Autorelease() TKTokenKeychainItem {
	rv := objc.Send[TKTokenKeychainItem](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenKeychainItem creates a new TKTokenKeychainItem instance.
func NewTKTokenKeychainItem() TKTokenKeychainItem {
	class := getTKTokenKeychainItemClass()
	rv := objc.Send[TKTokenKeychainItem](objc.ID(class.class), objc.Sel("new"))
	return rv
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
func NewTKTokenKeychainItemWithObjectID(objectID TKTokenObjectID) TKTokenKeychainItem {
	instance := getTKTokenKeychainItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObjectID:"), objectID)
	return TKTokenKeychainItemFromID(rv)
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
func (t TKTokenKeychainItem) InitWithObjectID(objectID TKTokenObjectID) TKTokenKeychainItem {
	rv := objc.Send[TKTokenKeychainItem](t.ID, objc.Sel("initWithObjectID:"), objectID)
	return rv
}

// Returns the object ID used for keychain item identification.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainItem/objectID
func (t TKTokenKeychainItem) ObjectID() TKTokenObjectID {
	rv := objc.Send[TKTokenObjectID](t.ID, objc.Sel("objectID"))
	return TKTokenObjectID(rv)
}

// The user-visible label for the keychain item.
//
// # Discussion
//
// This property is equivalent to the `kSecAttrLabel` attribute type.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainItem/label
func (t TKTokenKeychainItem) Label() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (t TKTokenKeychainItem) SetLabel(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setLabel:"), objc.String(value))
}

// Access constraints for the keychain item, keyed by [TKTokenOperation]
// values wrapped in [NSNumber] objects.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainItem/constraints
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
// [TKTokenOperation]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenOperation
func (t TKTokenKeychainItem) Constraints() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("constraints"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t TKTokenKeychainItem) SetConstraints(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setConstraints:"), value)
}
