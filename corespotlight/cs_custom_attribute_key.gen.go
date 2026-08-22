// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CSCustomAttributeKey] class.
var (
	_CSCustomAttributeKeyClass     CSCustomAttributeKeyClass
	_CSCustomAttributeKeyClassOnce sync.Once
)

func getCSCustomAttributeKeyClass() CSCustomAttributeKeyClass {
	_CSCustomAttributeKeyClassOnce.Do(func() {
		_CSCustomAttributeKeyClass = CSCustomAttributeKeyClass{class: objc.GetClass("CSCustomAttributeKey")}
	})
	return _CSCustomAttributeKeyClass
}

// GetCSCustomAttributeKeyClass returns the class object for CSCustomAttributeKey.
func GetCSCustomAttributeKeyClass() CSCustomAttributeKeyClass {
	return getCSCustomAttributeKeyClass()
}

type CSCustomAttributeKeyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CSCustomAttributeKeyClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CSCustomAttributeKeyClass) Alloc() CSCustomAttributeKey {
	rv := objc.Send[CSCustomAttributeKey](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A key associated with a custom attribute for a searchable item.
//
// # Overview
//
// The [CSCustomAttributeKey] class defines a key that you can associate with
// a custom attribute for a searchable item. Item attributes provide metadata
// about the item that can be indexed and displayed to users in search
// results.
//
// Although the Core Spotlight framework provides several predefined
// attributes, such as title and description, you can create a
// [CSCustomAttributeKey] object to specify a custom attribute that makes
// sense in your domain.
//
// # Creating a custom attribute
//
//   - [CSCustomAttributeKey.InitWithKeyName]: Returns a new custom attribute key with the specified name.
//   - [CSCustomAttributeKey.InitWithKeyNameSearchableSearchableByDefaultUniqueMultiValued]: Returns a new custom attribute key with the specified name and properties.
//   - [CSCustomAttributeKey.InitWithCoder]
//
// # Getting the attribute details
//
//   - [CSCustomAttributeKey.KeyName]: The name of the custom attribute key.
//   - [CSCustomAttributeKey.IsMultiValued]: A Boolean value that indicates if the custom attribute is likely to have multiple values, such as arrays, associated with it.
//   - [CSCustomAttributeKey.IsSearchable]: A Boolean value that indicates if the custom attribute can be specified as a search term.
//   - [CSCustomAttributeKey.IsSearchableByDefault]: A Boolean value that indicates if the custom attribute should be searchable by default.
//   - [CSCustomAttributeKey.IsUnique]: A Boolean value that indicates if duplicate custom attribute values should be treated as the same value to save storage space.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey
type CSCustomAttributeKey struct {
	objectivec.Object
}

// CSCustomAttributeKeyFromID constructs a [CSCustomAttributeKey] from an objc.ID.
//
// A key associated with a custom attribute for a searchable item.
func CSCustomAttributeKeyFromID(id objc.ID) CSCustomAttributeKey {
	return CSCustomAttributeKey{objectivec.Object{ID: id}}
}

// NOTE: CSCustomAttributeKey adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CSCustomAttributeKey] class.
//
// # Creating a custom attribute
//
//   - [ICSCustomAttributeKey.InitWithKeyName]: Returns a new custom attribute key with the specified name.
//   - [ICSCustomAttributeKey.InitWithKeyNameSearchableSearchableByDefaultUniqueMultiValued]: Returns a new custom attribute key with the specified name and properties.
//   - [ICSCustomAttributeKey.InitWithCoder]
//
// # Getting the attribute details
//
//   - [ICSCustomAttributeKey.KeyName]: The name of the custom attribute key.
//   - [ICSCustomAttributeKey.IsMultiValued]: A Boolean value that indicates if the custom attribute is likely to have multiple values, such as arrays, associated with it.
//   - [ICSCustomAttributeKey.IsSearchable]: A Boolean value that indicates if the custom attribute can be specified as a search term.
//   - [ICSCustomAttributeKey.IsSearchableByDefault]: A Boolean value that indicates if the custom attribute should be searchable by default.
//   - [ICSCustomAttributeKey.IsUnique]: A Boolean value that indicates if duplicate custom attribute values should be treated as the same value to save storage space.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey
type ICSCustomAttributeKey interface {
	objectivec.IObject

	// Topic: Creating a custom attribute

	// Returns a new custom attribute key with the specified name.
	InitWithKeyName(keyName string) CSCustomAttributeKey
	// Returns a new custom attribute key with the specified name and properties.
	InitWithKeyNameSearchableSearchableByDefaultUniqueMultiValued(keyName string, searchable bool, searchableByDefault bool, unique bool, multiValued bool) CSCustomAttributeKey
	InitWithCoder(coder foundation.INSCoder) CSCustomAttributeKey

	// Topic: Getting the attribute details

	// The name of the custom attribute key.
	KeyName() string
	// A Boolean value that indicates if the custom attribute is likely to have multiple values, such as arrays, associated with it.
	IsMultiValued() bool
	// A Boolean value that indicates if the custom attribute can be specified as a search term.
	IsSearchable() bool
	// A Boolean value that indicates if the custom attribute should be searchable by default.
	IsSearchableByDefault() bool
	// A Boolean value that indicates if duplicate custom attribute values should be treated as the same value to save storage space.
	IsUnique() bool

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CSCustomAttributeKey) Init() CSCustomAttributeKey {
	rv := objc.Send[CSCustomAttributeKey](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CSCustomAttributeKey) Autorelease() CSCustomAttributeKey {
	rv := objc.Send[CSCustomAttributeKey](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSCustomAttributeKey creates a new CSCustomAttributeKey instance.
func NewCSCustomAttributeKey() CSCustomAttributeKey {
	class := getCSCustomAttributeKeyClass()
	rv := objc.Send[CSCustomAttributeKey](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/init(coder:)
func NewCSCustomAttributeKeyWithCoder(coder foundation.INSCoder) CSCustomAttributeKey {
	instance := getCSCustomAttributeKeyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CSCustomAttributeKeyFromID(rv)
}

// Returns a new custom attribute key with the specified name.
//
// keyName: The name of the custom attribute for use as a key in a
// [CSSearchableItemAttributeSet]. The key name must be a string that contains
// only ASCII characters and no punctuation other than the underscore (that is
// “_”). The prefix `kMD` is reserved.
//
// # Return Value
//
// A new custom attribute key.
//
// # Discussion
//
// To create custom attribute key names, use a reverse DNS format that
// includes your company name and does not include the period character
// (”.”). For example, a key name of the form
// `com_mycompany_myapp_mykeyname` works well.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/init(keyName:)
func NewCSCustomAttributeKeyWithKeyName(keyName string) CSCustomAttributeKey {
	instance := getCSCustomAttributeKeyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKeyName:"), objc.String(keyName))
	return CSCustomAttributeKeyFromID(rv)
}

// Returns a new custom attribute key with the specified name and properties.
//
// keyName: The name of the custom attribute for use as a key in a
// [CSSearchableItemAttributeSet]. The key name must be a string that contains
// only ASCII characters and no punctuation other than the underscore (that
// is, “_”). The prefix `kMD` is reserved.
//
// searchable: A Boolean value that indicates if the attribute can be specified as a
// search term.
//
// searchableByDefault: A Boolean value that indicates if the attribute should be searchable by
// default.
//
// unique: A Boolean value that indicates if duplicate values should be treated as the
// same value to save storage space.
//
// multiValued: A Boolean value that indicates if the attribute is likely to have multiple
// values, such as arrays, associated with it.
//
// # Return Value
//
// A new custom attribute key with the specified name and properties.
//
// # Discussion
//
// To create custom attribute key names, it’s recommended that you use a
// reverse DNS format that includes your company name and does not include the
// period character (”.”). For example, a key name of the form
// `com_mycompany_myapp_mykeyname` works well.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/init(keyName:searchable:searchableByDefault:unique:multiValued:)
func NewCSCustomAttributeKeyWithKeyNameSearchableSearchableByDefaultUniqueMultiValued(keyName string, searchable bool, searchableByDefault bool, unique bool, multiValued bool) CSCustomAttributeKey {
	instance := getCSCustomAttributeKeyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKeyName:searchable:searchableByDefault:unique:multiValued:"), objc.String(keyName), searchable, searchableByDefault, unique, multiValued)
	return CSCustomAttributeKeyFromID(rv)
}

// Returns a new custom attribute key with the specified name.
//
// keyName: The name of the custom attribute for use as a key in a
// [CSSearchableItemAttributeSet]. The key name must be a string that contains
// only ASCII characters and no punctuation other than the underscore (that is
// “_”). The prefix `kMD` is reserved.
//
// # Return Value
//
// A new custom attribute key.
//
// # Discussion
//
// To create custom attribute key names, use a reverse DNS format that
// includes your company name and does not include the period character
// (”.”). For example, a key name of the form
// `com_mycompany_myapp_mykeyname` works well.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/init(keyName:)
func (c CSCustomAttributeKey) InitWithKeyName(keyName string) CSCustomAttributeKey {
	rv := objc.Send[CSCustomAttributeKey](c.ID, objc.Sel("initWithKeyName:"), objc.String(keyName))
	return rv
}

// Returns a new custom attribute key with the specified name and properties.
//
// keyName: The name of the custom attribute for use as a key in a
// [CSSearchableItemAttributeSet]. The key name must be a string that contains
// only ASCII characters and no punctuation other than the underscore (that
// is, “_”). The prefix `kMD` is reserved.
//
// searchable: A Boolean value that indicates if the attribute can be specified as a
// search term.
//
// searchableByDefault: A Boolean value that indicates if the attribute should be searchable by
// default.
//
// unique: A Boolean value that indicates if duplicate values should be treated as the
// same value to save storage space.
//
// multiValued: A Boolean value that indicates if the attribute is likely to have multiple
// values, such as arrays, associated with it.
//
// # Return Value
//
// A new custom attribute key with the specified name and properties.
//
// # Discussion
//
// To create custom attribute key names, it’s recommended that you use a
// reverse DNS format that includes your company name and does not include the
// period character (”.”). For example, a key name of the form
// `com_mycompany_myapp_mykeyname` works well.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/init(keyName:searchable:searchableByDefault:unique:multiValued:)
func (c CSCustomAttributeKey) InitWithKeyNameSearchableSearchableByDefaultUniqueMultiValued(keyName string, searchable bool, searchableByDefault bool, unique bool, multiValued bool) CSCustomAttributeKey {
	rv := objc.Send[CSCustomAttributeKey](c.ID, objc.Sel("initWithKeyName:searchable:searchableByDefault:unique:multiValued:"), objc.String(keyName), searchable, searchableByDefault, unique, multiValued)
	return rv
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/init(coder:)
func (c CSCustomAttributeKey) InitWithCoder(coder foundation.INSCoder) CSCustomAttributeKey {
	rv := objc.Send[CSCustomAttributeKey](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (c CSCustomAttributeKey) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The name of the custom attribute key.
//
// # Discussion
//
// The key name is a string that contains only ASCII characters and no
// punctuation other than the underscore (that is “_”). The prefix `kMD`
// is reserved. To create a custom attribute key name, it’s recommended that
// you use a reverse DNS format that includes your company name and does not
// include the period character (”.”). For example, a key name of the form
// `com_mycompany_myapp_mykeyname` works well.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/keyName
func (c CSCustomAttributeKey) KeyName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("keyName"))
	return foundation.NSStringFromID(rv).String()
}

// A Boolean value that indicates if the custom attribute is likely to have
// multiple values, such as arrays, associated with it.
//
// # Discussion
//
// The default value of this property is `false`.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/isMultiValued
func (c CSCustomAttributeKey) IsMultiValued() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isMultiValued"))
	return rv
}

// A Boolean value that indicates if the custom attribute can be specified as
// a search term.
//
// # Discussion
//
// The default value of this property is `true`.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/isSearchable
func (c CSCustomAttributeKey) IsSearchable() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isSearchable"))
	return rv
}

// A Boolean value that indicates if the custom attribute should be searchable
// by default.
//
// # Discussion
//
// The default value of this property is `false`.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/isSearchableByDefault
func (c CSCustomAttributeKey) IsSearchableByDefault() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isSearchableByDefault"))
	return rv
}

// A Boolean value that indicates if duplicate custom attribute values should
// be treated as the same value to save storage space.
//
// # Discussion
//
// The default value of this property is `false`.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/isUnique
func (c CSCustomAttributeKey) IsUnique() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isUnique"))
	return rv
}
