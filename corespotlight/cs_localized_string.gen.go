// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CSLocalizedString] class.
var (
	_CSLocalizedStringClass     CSLocalizedStringClass
	_CSLocalizedStringClassOnce sync.Once
)

func getCSLocalizedStringClass() CSLocalizedStringClass {
	_CSLocalizedStringClassOnce.Do(func() {
		_CSLocalizedStringClass = CSLocalizedStringClass{class: objc.GetClass("CSLocalizedString")}
	})
	return _CSLocalizedStringClass
}

// GetCSLocalizedStringClass returns the class object for CSLocalizedString.
func GetCSLocalizedStringClass() CSLocalizedStringClass {
	return getCSLocalizedStringClass()
}

type CSLocalizedStringClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CSLocalizedStringClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CSLocalizedStringClass) Alloc() CSLocalizedString {
	rv := objc.Send[CSLocalizedString](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that displays localized text in search results related to your
// app.
//
// # Overview
//
// The [CSLocalizedString] class helps you localize text in searchable items.
// You can use a [CSLocalizedString] object in place of an [NSString] object
// to display localized text in search results related to your app.
//
// For example, you might use the following code to define a
// [CSLocalizedString] object for a searchable item you want to identify as
// “Song” in English:
//
// # Specifying localized strings
//
//   - [CSLocalizedString.InitWithLocalizedStrings]: Initializes a [CSLocalizedString] object with the specified dictionary of localized strings.
//
// # Getting a localized string
//
//   - [CSLocalizedString.LocalizedString]: Returns the localized string for the current language.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSLocalizedString
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
type CSLocalizedString struct {
	foundation.NSString
}

// CSLocalizedStringFromID constructs a [CSLocalizedString] from an objc.ID.
//
// An object that displays localized text in search results related to your
// app.
func CSLocalizedStringFromID(id objc.ID) CSLocalizedString {
	return CSLocalizedString{NSString: foundation.NSStringFromID(id)}
}

// NOTE: CSLocalizedString adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CSLocalizedString] class.
//
// # Specifying localized strings
//
//   - [ICSLocalizedString.InitWithLocalizedStrings]: Initializes a [CSLocalizedString] object with the specified dictionary of localized strings.
//
// # Getting a localized string
//
//   - [ICSLocalizedString.LocalizedString]: Returns the localized string for the current language.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSLocalizedString
type ICSLocalizedString interface {
	foundation.INSString

	// Topic: Specifying localized strings

	// Initializes a [CSLocalizedString] object with the specified dictionary of localized strings.
	InitWithLocalizedStrings(localizedStrings foundation.INSDictionary) CSLocalizedString

	// Topic: Getting a localized string

	// Returns the localized string for the current language.
	LocalizedString() string
}

// Init initializes the instance.
func (c CSLocalizedString) Init() CSLocalizedString {
	rv := objc.Send[CSLocalizedString](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CSLocalizedString) Autorelease() CSLocalizedString {
	rv := objc.Send[CSLocalizedString](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSLocalizedString creates a new CSLocalizedString instance.
func NewCSLocalizedString() CSLocalizedString {
	class := getCSLocalizedStringClass()
	rv := objc.Send[CSLocalizedString](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a [CSLocalizedString] object with the specified dictionary of
// localized strings.
//
// localizedStrings: A dictionary in which each key-value pair consists of a language designator
// and a localized string. For example, you might pass in a dictionary like
// `@{@"en":@"Email Message"}`.
//
// # Return Value
//
// An object that contains the localized versions for a specific string.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSLocalizedString/init(localizedStrings:)
func NewCSLocalizedStringWithLocalizedStrings(localizedStrings foundation.INSDictionary) CSLocalizedString {
	instance := getCSLocalizedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLocalizedStrings:"), localizedStrings)
	return CSLocalizedStringFromID(rv)
}

// Initializes a [CSLocalizedString] object with the specified dictionary of
// localized strings.
//
// localizedStrings: A dictionary in which each key-value pair consists of a language designator
// and a localized string. For example, you might pass in a dictionary like
// `@{@"en":@"Email Message"}`.
//
// # Return Value
//
// An object that contains the localized versions for a specific string.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSLocalizedString/init(localizedStrings:)
func (c CSLocalizedString) InitWithLocalizedStrings(localizedStrings foundation.INSDictionary) CSLocalizedString {
	rv := objc.Send[CSLocalizedString](c.ID, objc.Sel("initWithLocalizedStrings:"), localizedStrings)
	return rv
}

// Returns the localized string for the current language.
//
// # Return Value
//
// The localized string for the current language.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSLocalizedString/localizedString()
func (c CSLocalizedString) LocalizedString() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("localizedString"))
	return foundation.NSStringFromID(rv).String()
}
