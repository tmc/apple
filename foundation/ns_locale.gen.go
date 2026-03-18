// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSLocale] class.
var (
	_NSLocaleClass     NSLocaleClass
	_NSLocaleClassOnce sync.Once
)

func getNSLocaleClass() NSLocaleClass {
	_NSLocaleClassOnce.Do(func() {
		_NSLocaleClass = NSLocaleClass{class: objc.GetClass("NSLocale")}
	})
	return _NSLocaleClass
}

// GetNSLocaleClass returns the class object for NSLocale.
func GetNSLocaleClass() NSLocaleClass {
	return getNSLocaleClass()
}

type NSLocaleClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSLocaleClass) Alloc() NSLocale {
	rv := objc.Send[NSLocale](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Information about linguistic, cultural, and technological conventions for
// use in formatting data for presentation.
//
// # Overview
// 
// In Swift, this object bridges to [Locale]; use [NSLocale] when you need
// reference semantics or other Foundation-specific behavior.
// 
// You typically use a locale to format and interpret information about and
// according to the user’s customs and preferences.
// 
// You can initialize any number of locale instances with
// [NSLocale.InitWithLocaleIdentifier] using one of the locale identifiers found in the
// [NSLocale.AvailableLocaleIdentifiers] array. However, you usually use a locale
// configured to match the preferences of the current user.
// 
// Use the [NSLocale.CurrentLocale] property to get the locale matching the current
// user’s preferences. If you need to be alerted when the user does make
// changes to region settings, register for the
// [currentLocaleDidChangeNotification] notification. Alternatively, you can
// use the [NSLocale.AutoupdatingCurrentLocale] property to get a locale that
// automatically updates with the user’s configuration settings:
// 
// You can inspect a locale by reading its properties, as listed in Getting
// Information About a Locale. For properties containing a code or identifier,
// you can then obtain a string suitable for presentation to the user with the
// methods listed in Getting Display Information About a Locale. For example,
// you can report the user’s language as a string localized in that language
// using the autoupdating locale obtained in the previous example:
// 
// You frequently use a locale in conjunction with a formatter. For example,
// the [NSDateFormatter] class has a [Locale] property that ensures dates are
// converted to strings that match the user’s expectations about date
// formatting. By default, this property indicates the user’s current
// locale, which is usually the behavior you want, but you can instead set it
// to another locale instance to obtain a different output. See [Data
// Formatting Guide] for more information about working with formatters.
// 
// [NSLocale] is with its Core Foundation counterpart, [CFLocale]. See
// [Toll-Free Bridging] for more information on toll-free bridging.
//
// [CFLocale]: https://developer.apple.com/documentation/CoreFoundation/CFLocale
// [Data Formatting Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DataFormatting/DataFormatting.html#//apple_ref/doc/uid/10000029i
// [Locale]: https://developer.apple.com/documentation/Foundation/Locale
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
// [currentLocaleDidChangeNotification]: https://developer.apple.com/documentation/Foundation/NSLocale/currentLocaleDidChangeNotification
//
// # Initializing a Locale
//
//   - [NSLocale.InitWithLocaleIdentifier]: Initializes a locale using a given locale identifier.
//
// # Getting Information About a Locale
//
//   - [NSLocale.LocaleIdentifier]: The identifier for the locale.
//   - [NSLocale.CountryCode]: The country or region code for the locale.
//   - [NSLocale.LanguageCode]: The language code for the locale.
//   - [NSLocale.ScriptCode]: The script code for the locale.
//   - [NSLocale.VariantCode]: The variant code for the locale.
//   - [NSLocale.ExemplarCharacterSet]: The exemplar character set for the locale.
//   - [NSLocale.CollationIdentifier]: The collation identifier for the locale.
//   - [NSLocale.CollatorIdentifier]: The collator identifier for the locale.
//   - [NSLocale.UsesMetricSystem]: A Boolean value that indicates whether the locale uses the metric system.
//   - [NSLocale.DecimalSeparator]: The decimal separator for the locale.
//   - [NSLocale.GroupingSeparator]: The grouping separator for the locale.
//   - [NSLocale.CurrencyCode]: The currency code for the locale.
//   - [NSLocale.CurrencySymbol]: The currency symbol for the locale.
//   - [NSLocale.CalendarIdentifier]: The calendar identifier for the locale.
//   - [NSLocale.QuotationBeginDelimiter]: The begin quotation symbol for the locale.
//   - [NSLocale.QuotationEndDelimiter]: The end quotation symbol for the locale.
//   - [NSLocale.AlternateQuotationBeginDelimiter]: The alternate begin quotation symbol for the locale.
//   - [NSLocale.AlternateQuotationEndDelimiter]: The alternate end quotation symbol for the locale.
//
// # Getting Display Information About a Locale
//
//   - [NSLocale.LocalizedStringForLocaleIdentifier]: Returns the localized string for the specified locale identifier.
//   - [NSLocale.LocalizedStringForCountryCode]: Returns the localized string for a country or region code.
//   - [NSLocale.LocalizedStringForLanguageCode]: Returns the localized string for the specified language code.
//   - [NSLocale.LocalizedStringForScriptCode]: Returns the localized string for the specified script code.
//   - [NSLocale.LocalizedStringForVariantCode]: Returns the localized string for the specified variant code.
//   - [NSLocale.LocalizedStringForCollationIdentifier]: Returns the localized string for the specified collation identifier.
//   - [NSLocale.LocalizedStringForCollatorIdentifier]: Returns the localized string for the specified collator identifier.
//   - [NSLocale.LocalizedStringForCurrencyCode]: Returns the localized string for the specified currency code.
//   - [NSLocale.LocalizedStringForCalendarIdentifier]: Returns the localized string for the specified calendar identifier.
//
// # Accessing Locale Information by Key
//
//   - [NSLocale.ObjectForKey]: Returns the value of the component corresponding to the specified key.
//   - [NSLocale.DisplayNameForKeyValue]: Returns the display name for the given locale component value.
//
// # Instance Properties
//
//   - [NSLocale.LanguageIdentifier]: Returns the identifier for the language part of the locale. For example, returns “en-US” for “en_US@rg=gbzzzz”  locale.
//   - [NSLocale.RegionCode]: Returns the region code of the locale. If the `rg` subtag is present, the value of the subtag will be used. For example,  returns “GB” for “en_US@rg=gbzzzz” locale. If the `localeIdentifier` doesn’t contain a region, returns `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale
type NSLocale struct {
	objectivec.Object
}

// NSLocaleFromID constructs a [NSLocale] from an objc.ID.
//
// Information about linguistic, cultural, and technological conventions for
// use in formatting data for presentation.
func NSLocaleFromID(id objc.ID) NSLocale {
	return NSLocale{objectivec.Object{ID: id}}
}
// NOTE: NSLocale adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSLocale] class.
//
// # Initializing a Locale
//
//   - [INSLocale.InitWithLocaleIdentifier]: Initializes a locale using a given locale identifier.
//
// # Getting Information About a Locale
//
//   - [INSLocale.LocaleIdentifier]: The identifier for the locale.
//   - [INSLocale.CountryCode]: The country or region code for the locale.
//   - [INSLocale.LanguageCode]: The language code for the locale.
//   - [INSLocale.ScriptCode]: The script code for the locale.
//   - [INSLocale.VariantCode]: The variant code for the locale.
//   - [INSLocale.ExemplarCharacterSet]: The exemplar character set for the locale.
//   - [INSLocale.CollationIdentifier]: The collation identifier for the locale.
//   - [INSLocale.CollatorIdentifier]: The collator identifier for the locale.
//   - [INSLocale.UsesMetricSystem]: A Boolean value that indicates whether the locale uses the metric system.
//   - [INSLocale.DecimalSeparator]: The decimal separator for the locale.
//   - [INSLocale.GroupingSeparator]: The grouping separator for the locale.
//   - [INSLocale.CurrencyCode]: The currency code for the locale.
//   - [INSLocale.CurrencySymbol]: The currency symbol for the locale.
//   - [INSLocale.CalendarIdentifier]: The calendar identifier for the locale.
//   - [INSLocale.QuotationBeginDelimiter]: The begin quotation symbol for the locale.
//   - [INSLocale.QuotationEndDelimiter]: The end quotation symbol for the locale.
//   - [INSLocale.AlternateQuotationBeginDelimiter]: The alternate begin quotation symbol for the locale.
//   - [INSLocale.AlternateQuotationEndDelimiter]: The alternate end quotation symbol for the locale.
//
// # Getting Display Information About a Locale
//
//   - [INSLocale.LocalizedStringForLocaleIdentifier]: Returns the localized string for the specified locale identifier.
//   - [INSLocale.LocalizedStringForCountryCode]: Returns the localized string for a country or region code.
//   - [INSLocale.LocalizedStringForLanguageCode]: Returns the localized string for the specified language code.
//   - [INSLocale.LocalizedStringForScriptCode]: Returns the localized string for the specified script code.
//   - [INSLocale.LocalizedStringForVariantCode]: Returns the localized string for the specified variant code.
//   - [INSLocale.LocalizedStringForCollationIdentifier]: Returns the localized string for the specified collation identifier.
//   - [INSLocale.LocalizedStringForCollatorIdentifier]: Returns the localized string for the specified collator identifier.
//   - [INSLocale.LocalizedStringForCurrencyCode]: Returns the localized string for the specified currency code.
//   - [INSLocale.LocalizedStringForCalendarIdentifier]: Returns the localized string for the specified calendar identifier.
//
// # Accessing Locale Information by Key
//
//   - [INSLocale.ObjectForKey]: Returns the value of the component corresponding to the specified key.
//   - [INSLocale.DisplayNameForKeyValue]: Returns the display name for the given locale component value.
//
// # Instance Properties
//
//   - [INSLocale.LanguageIdentifier]: Returns the identifier for the language part of the locale. For example, returns “en-US” for “en_US@rg=gbzzzz”  locale.
//   - [INSLocale.RegionCode]: Returns the region code of the locale. If the `rg` subtag is present, the value of the subtag will be used. For example,  returns “GB” for “en_US@rg=gbzzzz” locale. If the `localeIdentifier` doesn’t contain a region, returns `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale
type INSLocale interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Initializing a Locale

	// Initializes a locale using a given locale identifier.
	InitWithLocaleIdentifier(string_ string) NSLocale

	// Topic: Getting Information About a Locale

	// The identifier for the locale.
	LocaleIdentifier() string
	// The country or region code for the locale.
	CountryCode() string
	// The language code for the locale.
	LanguageCode() string
	// The script code for the locale.
	ScriptCode() string
	// The variant code for the locale.
	VariantCode() string
	// The exemplar character set for the locale.
	ExemplarCharacterSet() INSCharacterSet
	// The collation identifier for the locale.
	CollationIdentifier() string
	// The collator identifier for the locale.
	CollatorIdentifier() string
	// A Boolean value that indicates whether the locale uses the metric system.
	UsesMetricSystem() bool
	// The decimal separator for the locale.
	DecimalSeparator() string
	// The grouping separator for the locale.
	GroupingSeparator() string
	// The currency code for the locale.
	CurrencyCode() string
	// The currency symbol for the locale.
	CurrencySymbol() string
	// The calendar identifier for the locale.
	CalendarIdentifier() string
	// The begin quotation symbol for the locale.
	QuotationBeginDelimiter() string
	// The end quotation symbol for the locale.
	QuotationEndDelimiter() string
	// The alternate begin quotation symbol for the locale.
	AlternateQuotationBeginDelimiter() string
	// The alternate end quotation symbol for the locale.
	AlternateQuotationEndDelimiter() string

	// Topic: Getting Display Information About a Locale

	// Returns the localized string for the specified locale identifier.
	LocalizedStringForLocaleIdentifier(localeIdentifier string) string
	// Returns the localized string for a country or region code.
	LocalizedStringForCountryCode(countryCode string) string
	// Returns the localized string for the specified language code.
	LocalizedStringForLanguageCode(languageCode string) string
	// Returns the localized string for the specified script code.
	LocalizedStringForScriptCode(scriptCode string) string
	// Returns the localized string for the specified variant code.
	LocalizedStringForVariantCode(variantCode string) string
	// Returns the localized string for the specified collation identifier.
	LocalizedStringForCollationIdentifier(collationIdentifier string) string
	// Returns the localized string for the specified collator identifier.
	LocalizedStringForCollatorIdentifier(collatorIdentifier string) string
	// Returns the localized string for the specified currency code.
	LocalizedStringForCurrencyCode(currencyCode string) string
	// Returns the localized string for the specified calendar identifier.
	LocalizedStringForCalendarIdentifier(calendarIdentifier string) string

	// Topic: Accessing Locale Information by Key

	// Returns the value of the component corresponding to the specified key.
	ObjectForKey(key NSLocaleKey) objectivec.IObject
	// Returns the display name for the given locale component value.
	DisplayNameForKeyValue(key NSLocaleKey, value objectivec.IObject) string

	// Topic: Instance Properties

	// Returns the identifier for the language part of the locale. For example, returns “en-US” for “en_US@rg=gbzzzz”  locale.
	LanguageIdentifier() string
	// Returns the region code of the locale. If the `rg` subtag is present, the value of the subtag will be used. For example,  returns “GB” for “en_US@rg=gbzzzz” locale. If the `localeIdentifier` doesn’t contain a region, returns `nil`.
	RegionCode() string

	// The locale for the receiver.
	Locale() INSLocale
	SetLocale(value INSLocale)
}

// Init initializes the instance.
func (l NSLocale) Init() NSLocale {
	rv := objc.Send[NSLocale](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l NSLocale) Autorelease() NSLocale {
	rv := objc.Send[NSLocale](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSLocale creates a new NSLocale instance.
func NewNSLocale() NSLocale {
	class := getNSLocaleClass()
	rv := objc.Send[NSLocale](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a locale initialized from data in the given unarchiver.
//
// coder: The decoder to use during initialization.
//
// # Return Value
// 
// The initialized locale.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/init(coder:)
func NewLocaleWithCoder(coder INSCoder) NSLocale {
	instance := getNSLocaleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSLocaleFromID(rv)
}

// Initializes a locale using a given locale identifier.
//
// string: The identifier for the new locale.
//
// # Return Value
// 
// The initialized locale.
//
// # Discussion
// 
// This method is the designated initializer for this class.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/init(localeIdentifier:)
func NewLocaleWithLocaleIdentifier(string_ string) NSLocale {
	instance := getNSLocaleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLocaleIdentifier:"), objc.String(string_))
	return NSLocaleFromID(rv)
}

// Initializes a locale using a given locale identifier.
//
// string: The identifier for the new locale.
//
// # Return Value
// 
// The initialized locale.
//
// # Discussion
// 
// This method is the designated initializer for this class.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/init(localeIdentifier:)
func (l NSLocale) InitWithLocaleIdentifier(string_ string) NSLocale {
	rv := objc.Send[NSLocale](l.ID, objc.Sel("initWithLocaleIdentifier:"), objc.String(string_))
	return rv
}

// Returns a locale initialized from data in the given unarchiver.
//
// coder: The decoder to use during initialization.
//
// # Return Value
// 
// The initialized locale.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/init(coder:)
func (l NSLocale) InitWithCoder(coder INSCoder) NSLocale {
	rv := objc.Send[NSLocale](l.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Returns the localized string for the specified locale identifier.
//
// localeIdentifier: The identifier for which you want the localized name.
//
// # Return Value
// 
// The localized name of the locale.
//
// # Discussion
// 
// For example, calling this method on an American English (`en_US`) locale,
// passing `"fr_FR"` for `localeIdentifier`, produces the string `"French
// (France)"`.
// 
// This method is equivalent to calling the [DisplayNameForKeyValue] method
// passing the [identifier] key and `localeIdentifier` value.
//
// [identifier]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/identifier
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forLocaleIdentifier:)
func (l NSLocale) LocalizedStringForLocaleIdentifier(localeIdentifier string) string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("localizedStringForLocaleIdentifier:"), objc.String(localeIdentifier))
	return NSStringFromID(rv).String()
}

// Returns the localized string for a country or region code.
//
// countryCode: The code of the country or region that indicates the name you want.
//
// # Return Value
// 
// The localized name of the country or region.
//
// # Discussion
// 
// For example, calling this method on an American English (`en_US`) locale,
// passing `"GB"` for `countryCode`, produces the string `"United Kingdom"`.
// 
// This method is equivalent to calling the [DisplayNameForKeyValue] method
// passing the [countryCode] key and `countryCode` value.
//
// [countryCode]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/countryCode
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forCountryCode:)
func (l NSLocale) LocalizedStringForCountryCode(countryCode string) string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("localizedStringForCountryCode:"), objc.String(countryCode))
	return NSStringFromID(rv).String()
}

// Returns the localized string for the specified language code.
//
// languageCode: The language code indicating the language whose name you want.
//
// # Return Value
// 
// The localized name of the language.
//
// # Discussion
// 
// For example, calling this method on an American English (`en_US`) locale,
// passing `"zh"` for `languageCode`, produces the string `"Chinese"`.
// 
// This method is equivalent to calling the [DisplayNameForKeyValue] method
// passing the [languageCode] key and `languageCode` value.
//
// [languageCode]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/languageCode
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forLanguageCode:)
func (l NSLocale) LocalizedStringForLanguageCode(languageCode string) string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("localizedStringForLanguageCode:"), objc.String(languageCode))
	return NSStringFromID(rv).String()
}

// Returns the localized string for the specified script code.
//
// scriptCode: The script code for the script whose name you want.
//
// # Return Value
// 
// The localized name of the script.
//
// # Discussion
// 
// For example, calling this method on an American English (`en_US`) locale,
// passing `"Hant"` for `scriptCode`, produces the string `"Traditional Han"`.
// 
// This method is equivalent to calling the [DisplayNameForKeyValue] method
// passing the [scriptCode] key and `scriptCode` value.
//
// [scriptCode]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/scriptCode
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forScriptCode:)
func (l NSLocale) LocalizedStringForScriptCode(scriptCode string) string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("localizedStringForScriptCode:"), objc.String(scriptCode))
	return NSStringFromID(rv).String()
}

// Returns the localized string for the specified variant code.
//
// variantCode: The variant code whose name you want.
//
// # Return Value
// 
// The localized name of the variant.
//
// # Discussion
// 
// For example, calling this method on an American English (`en_US`) locale,
// passing `"POSIX"` for `variantCode`, produces the string `"Computer"`.
// 
// This method is equivalent to calling the [DisplayNameForKeyValue] method
// passing the [variantCode] key and `variantCode` value.
//
// [variantCode]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/variantCode
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forVariantCode:)
func (l NSLocale) LocalizedStringForVariantCode(variantCode string) string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("localizedStringForVariantCode:"), objc.String(variantCode))
	return NSStringFromID(rv).String()
}

// Returns the localized string for the specified collation identifier.
//
// collationIdentifier: The identifier for the collation whose name you want.
//
// # Return Value
// 
// The localized name for the collation.
//
// # Discussion
// 
// For example, calling this method on an American English (`en_US`) locale,
// passing `"phonebook"` for `collationIdentifier`, produces the string
// `"Phonebook Sort Order"`.
// 
// This method is equivalent to calling the [DisplayNameForKeyValue] method
// passing the [collationIdentifier] key and `collationIdentifier` value.
//
// [collationIdentifier]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/collationIdentifier
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forCollationIdentifier:)
func (l NSLocale) LocalizedStringForCollationIdentifier(collationIdentifier string) string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("localizedStringForCollationIdentifier:"), objc.String(collationIdentifier))
	return NSStringFromID(rv).String()
}

// Returns the localized string for the specified collator identifier.
//
// collatorIdentifier: The identifier for the collator whose name you want.
//
// # Return Value
// 
// The localized name of the collator.
//
// # Discussion
// 
// This method is equivalent to calling the [DisplayNameForKeyValue] method
// passing the [collatorIdentifier] key and `collatorIdentifier` value.
//
// [collatorIdentifier]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/collatorIdentifier
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forCollatorIdentifier:)
func (l NSLocale) LocalizedStringForCollatorIdentifier(collatorIdentifier string) string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("localizedStringForCollatorIdentifier:"), objc.String(collatorIdentifier))
	return NSStringFromID(rv).String()
}

// Returns the localized string for the specified currency code.
//
// currencyCode: The code for the currency whose name you want.
//
// # Return Value
// 
// The localized name of the currency.
//
// # Discussion
// 
// For example, calling this method on an American English (`en_US`) locale,
// passing `"JPY"` for `currencyCode`, produces the string `"Japanese Yen"`.
// 
// This method is equivalent to calling the [DisplayNameForKeyValue] method
// passing the [currencyCode] key and `currencyCode` value.
//
// [currencyCode]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/currencyCode
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forCurrencyCode:)
func (l NSLocale) LocalizedStringForCurrencyCode(currencyCode string) string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("localizedStringForCurrencyCode:"), objc.String(currencyCode))
	return NSStringFromID(rv).String()
}

// Returns the localized string for the specified calendar identifier.
//
// calendarIdentifier: The calendar identifier indicating the calendar whose name you want. Use
// one of the values listed in `Calendar Identifiers`.
//
// # Return Value
// 
// A human readable string suitable for display to the user corresponding to
// the given calendar.
//
// # Discussion
// 
// For example, on an American English (`en_US`) locale, passing [gregorian]
// as the identifier, produces the string `"Gregorian Calendar"`.
//
// [gregorian]: https://developer.apple.com/documentation/Foundation/NSCalendar/Identifier/gregorian
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forCalendarIdentifier:)
func (l NSLocale) LocalizedStringForCalendarIdentifier(calendarIdentifier string) string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("localizedStringForCalendarIdentifier:"), objc.String(calendarIdentifier))
	return NSStringFromID(rv).String()
}

// Returns the value of the component corresponding to the specified key.
//
// key: The component for which to return the corresponding value. For possible
// values, see [NSLocaleKey].
//
// # Return Value
// 
// The object corresponding to `key`.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/object(forKey:)
func (l NSLocale) ObjectForKey(key NSLocaleKey) objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("objectForKey:"), objc.String(string(key)))
	return objectivec.Object{ID: rv}
}

// Returns the display name for the given locale component value.
//
// key: The locale property key of `value`. For possible values, see [NSLocaleKey].
//
// value: A value for `key`.
//
// # Return Value
// 
// The display name for `value`.
//
// # Discussion
// 
// Not all locale property keys have values with display name values.
// 
// You can use the [identifier] key to get the name of a locale in the
// language of another locale, as illustrated in the following examples.
// 
// The following example uses the `en_GB` locale.
//
// [identifier]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/identifier
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/displayName(forKey:value:)
func (l NSLocale) DisplayNameForKeyValue(key NSLocaleKey, value objectivec.IObject) string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("displayNameForKey:value:"), objc.String(string(key)), value)
	return NSStringFromID(rv).String()
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (l NSLocale) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](l.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns the canonical identifier for a given locale identification string.
//
// string: A locale identification string.
//
// # Return Value
// 
// The canonical identifier for an the locale identified by `string`.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/canonicalLocaleIdentifier(from:)
func (_NSLocaleClass NSLocaleClass) CanonicalLocaleIdentifierFromString(string_ string) string {
	rv := objc.Send[objc.ID](objc.ID(_NSLocaleClass.class), objc.Sel("canonicalLocaleIdentifierFromString:"), objc.String(string_))
	return NSStringFromID(rv).String()
}

// Returns a dictionary that is the result of parsing a locale ID.
//
// string: A locale ID, consisting of language, script, country, variant, and
// keyword/value pairs, for example, `"en_US@calendar=japanese"`.
//
// # Return Value
// 
// A dictionary that is the result of parsing `string` as a locale ID. The
// keys are the constant NSString constants corresponding to the locale ID
// components, and the values correspond to constants where available. For
// possible values, see [NSLocaleKey].
//
// # Discussion
// 
// For example, the locale identifier `"en_US@calendar=japanese"` yields a
// dictionary with three entries:
// 
// - [languageCode] = `en` - [countryCode] = [US] - [calendar] =
// [NSJapaneseCalendar]
//
// [NSJapaneseCalendar]: https://developer.apple.com/documentation/Foundation/NSJapaneseCalendar
// [calendar]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/calendar
// [countryCode]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/countryCode
// [languageCode]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/languageCode
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/components(fromLocaleIdentifier:)
func (_NSLocaleClass NSLocaleClass) ComponentsFromLocaleIdentifier(string_ string) INSDictionary {
	rv := objc.Send[objc.ID](objc.ID(_NSLocaleClass.class), objc.Sel("componentsFromLocaleIdentifier:"), objc.String(string_))
	return NSDictionaryFromID(rv)
}

// Returns a locale identifier from the components specified in a given
// dictionary.
//
// dict: A dictionary containing components that specify a locale. For possible
// values, see `NSLocale Component Keys`.
//
// # Return Value
// 
// A locale identifier created from the components specified in `dict`.
//
// # Discussion
// 
// This reverses the actions of [ComponentsFromLocaleIdentifier], so for
// example the dictionary `{NSLocaleLanguageCode="en",
// NSLocaleCountryCode="US", NSLocaleCalendar=NSJapaneseCalendar}` becomes
// `"en_US@calendar=japanese"`.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/localeIdentifier(fromComponents:)
func (_NSLocaleClass NSLocaleClass) LocaleIdentifierFromComponents(dict INSDictionary) string {
	rv := objc.Send[objc.ID](objc.ID(_NSLocaleClass.class), objc.Sel("localeIdentifierFromComponents:"), dict)
	return NSStringFromID(rv).String()
}

// Returns a canonical language identifier by mapping an arbitrary locale
// identification string to the canonical identifier.
//
// string: A string representation of an arbitrary locale identifier.
//
// # Return Value
// 
// A string that represents the canonical language identifier for the
// specified arbitrary locale identifier.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/canonicalLanguageIdentifier(from:)
func (_NSLocaleClass NSLocaleClass) CanonicalLanguageIdentifierFromString(string_ string) string {
	rv := objc.Send[objc.ID](objc.ID(_NSLocaleClass.class), objc.Sel("canonicalLanguageIdentifierFromString:"), objc.String(string_))
	return NSStringFromID(rv).String()
}

// Returns a locale identifier from a Windows locale code.
//
// lcid: The Windows locale code.
//
// # Return Value
// 
// The locale identifier.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/localeIdentifier(fromWindowsLocaleCode:)
func (_NSLocaleClass NSLocaleClass) LocaleIdentifierFromWindowsLocaleCode(lcid uint32) string {
	rv := objc.Send[objc.ID](objc.ID(_NSLocaleClass.class), objc.Sel("localeIdentifierFromWindowsLocaleCode:"), lcid)
	return NSStringFromID(rv).String()
}

// Returns a Window locale code from the locale identifier.
//
// localeIdentifier: The locale identifier.
//
// # Return Value
// 
// The Windows locale code.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/windowsLocaleCode(fromLocaleIdentifier:)
func (_NSLocaleClass NSLocaleClass) WindowsLocaleCodeFromLocaleIdentifier(localeIdentifier string) uint32 {
	rv := objc.Send[uint32](objc.ID(_NSLocaleClass.class), objc.Sel("windowsLocaleCodeFromLocaleIdentifier:"), objc.String(localeIdentifier))
	return rv
}

// Returns the direction of the sequence of characters in a line for the
// specified ISO language code.
//
// isoLangCode: The ISO language code.
//
// # Return Value
// 
// Returns the direction in which characters appear within a line in the
// specified language. See [NSLocale.LanguageDirection] for possible values.
// If the appropriate direction can’t be determined
// [LocaleLanguageDirectionUnknown] is returned.
//
// [NSLocale.LanguageDirection]: https://developer.apple.com/documentation/Foundation/NSLocale/LanguageDirection
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/characterDirection(forLanguage:)
func (_NSLocaleClass NSLocaleClass) CharacterDirectionForLanguage(isoLangCode string) NSLocaleLanguageDirection {
	rv := objc.Send[NSLocaleLanguageDirection](objc.ID(_NSLocaleClass.class), objc.Sel("characterDirectionForLanguage:"), objc.String(isoLangCode))
	return NSLocaleLanguageDirection(rv)
}

// Returns the direction of the sequence of lines for the specified ISO
// language code.
//
// isoLangCode: The ISO language code.
//
// # Return Value
// 
// Returns the direction in which lines appear in the specified language. See
// [NSLocale.LanguageDirection] for possible values. If the appropriate
// direction can’t be determined [LocaleLanguageDirectionUnknown] is
// returned.
//
// [NSLocale.LanguageDirection]: https://developer.apple.com/documentation/Foundation/NSLocale/LanguageDirection
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/lineDirection(forLanguage:)
func (_NSLocaleClass NSLocaleClass) LineDirectionForLanguage(isoLangCode string) NSLocaleLanguageDirection {
	rv := objc.Send[NSLocaleLanguageDirection](objc.ID(_NSLocaleClass.class), objc.Sel("lineDirectionForLanguage:"), objc.String(isoLangCode))
	return NSLocaleLanguageDirection(rv)
}

// Returns a locale initialized using the given locale identifier.
//
// ident: The identifier for the new locale.
//
// # Return Value
// 
// The initialized locale.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/localeWithLocaleIdentifier:
func (_NSLocaleClass NSLocaleClass) LocaleWithLocaleIdentifier(ident string) NSLocale {
	rv := objc.Send[objc.ID](objc.ID(_NSLocaleClass.class), objc.Sel("localeWithLocaleIdentifier:"), objc.String(ident))
	return NSLocaleFromID(rv)
}

// The identifier for the locale.
//
// # Discussion
// 
// Examples of locale identifiers include `"en_GB"`, `"es_ES_PREEURO"`, and
// `"zh-Hant_HK_POSIX@collation=pinyin;currency=CNY"`.
// 
// Use [localizedString(forIdentifier:)] to obtain a version of the value
// suitable for display to the user.
// 
// This property contains the same value returned by the [ObjectForKey] method
// when passing the [identifier] key.
//
// [identifier]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/identifier
// [localizedString(forIdentifier:)]: https://developer.apple.com/documentation/Foundation/Locale/localizedString(forIdentifier:)
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/localeIdentifier
func (l NSLocale) LocaleIdentifier() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("localeIdentifier"))
	return NSStringFromID(rv).String()
}

// The country or region code for the locale.
//
// # Discussion
// 
// Examples of country or region codes include `"GB"`, `"FR"`, and `"HK"`.
// 
// Use [LocalizedStringForCountryCode] to obtain a version of the value
// suitable for display to the user.
// 
// This property contains the same value returned by the [ObjectForKey] method
// when passing the [countryCode] key.
//
// [countryCode]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/countryCode
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/countryCode
func (l NSLocale) CountryCode() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("countryCode"))
	return NSStringFromID(rv).String()
}

// The language code for the locale.
//
// # Discussion
// 
// Examples of language codes include `"en"`, `"es"`, and `"zh"`.
// 
// Use [LocalizedStringForLanguageCode] to obtain a version of the value
// suitable for display to the user.
// 
// This property contains the same value returned by the [ObjectForKey] method
// when passing the [languageCode] key.
//
// [languageCode]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/languageCode
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/languageCode
func (l NSLocale) LanguageCode() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("languageCode"))
	return NSStringFromID(rv).String()
}

// The script code for the locale.
//
// # Discussion
// 
// Examples of script codes include `"Latn"` and `"Hant"`.
// 
// Use [LocalizedStringForScriptCode] to obtain a version of the value
// suitable for display to the user.
// 
// This property contains the same value returned by the [ObjectForKey] method
// when passing the [scriptCode] key.
//
// [scriptCode]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/scriptCode
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/scriptCode
func (l NSLocale) ScriptCode() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("scriptCode"))
	return NSStringFromID(rv).String()
}

// The variant code for the locale.
//
// # Discussion
// 
// Examples of variant code include `"POSIX"` and `"PREEURO"`.
// 
// Use [LocalizedStringForVariantCode] to obtain a version of the value
// suitable for display to the user.
// 
// This property contains the same value returned by the [ObjectForKey] method
// when passing the [variantCode] key.
//
// [variantCode]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/variantCode
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/variantCode
func (l NSLocale) VariantCode() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("variantCode"))
	return NSStringFromID(rv).String()
}

// The exemplar character set for the locale.
//
// # Discussion
// 
// This property contains the same value returned by the [ObjectForKey] method
// when passing the [exemplarCharacterSet] key.
//
// [exemplarCharacterSet]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/exemplarCharacterSet
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/exemplarCharacterSet
func (l NSLocale) ExemplarCharacterSet() INSCharacterSet {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("exemplarCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}

// The collation identifier for the locale.
//
// # Discussion
// 
// Example collation identifiers include `"phonebook"` and `"pinyin"`.
// 
// Use [LocalizedStringForCollationIdentifier] to obtain a version of the
// value suitable for display to the user.
// 
// This property contains the same value returned by the [ObjectForKey] method
// when passing the [collationIdentifier] key.
//
// [collationIdentifier]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/collationIdentifier
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/collationIdentifier
func (l NSLocale) CollationIdentifier() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("collationIdentifier"))
	return NSStringFromID(rv).String()
}

// The collator identifier for the locale.
//
// # Discussion
// 
// An example collator identifier is `"en"`.
// 
// Use [LocalizedStringForCollatorIdentifier] to obtain a version of the value
// suitable for display to the user.
// 
// This property contains the same value returned by the [ObjectForKey] method
// when passing the [collatorIdentifier] key.
//
// [collatorIdentifier]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/collatorIdentifier
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/collatorIdentifier
func (l NSLocale) CollatorIdentifier() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("collatorIdentifier"))
	return NSStringFromID(rv).String()
}

// A Boolean value that indicates whether the locale uses the metric system.
//
// # Discussion
// 
// This property contains the same value returned by the [ObjectForKey] method
// when passing the [usesMetricSystem] key.
//
// [usesMetricSystem]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/usesMetricSystem
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/usesMetricSystem
func (l NSLocale) UsesMetricSystem() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("usesMetricSystem"))
	return rv
}

// The decimal separator for the locale.
//
// # Discussion
// 
// Example decimal separators include `"."` and `","`.
// 
// This property contains the same value returned by the [ObjectForKey] method
// when passing the [decimalSeparator] key.
//
// [decimalSeparator]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/decimalSeparator
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/decimalSeparator
func (l NSLocale) DecimalSeparator() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("decimalSeparator"))
	return NSStringFromID(rv).String()
}

// The grouping separator for the locale.
//
// # Discussion
// 
// Example grouping separators include `","` and `" "`.
// 
// This property contains the same value returned by the [ObjectForKey] method
// when passing the [groupingSeparator] key.
//
// [groupingSeparator]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/groupingSeparator
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/groupingSeparator
func (l NSLocale) GroupingSeparator() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("groupingSeparator"))
	return NSStringFromID(rv).String()
}

// The currency code for the locale.
//
// # Discussion
// 
// Example currency codes include `"USD"`, `"EUR"`, and `"JPY"`.
// 
// Use [LocalizedStringForCurrencyCode] to obtain a version of the value
// suitable for display to the user.
// 
// This property contains the same value returned by the [ObjectForKey] method
// when passing the [currencyCode] key.
//
// [currencyCode]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/currencyCode
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/currencyCode
func (l NSLocale) CurrencyCode() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("currencyCode"))
	return NSStringFromID(rv).String()
}

// The currency symbol for the locale.
//
// # Discussion
// 
// Example currency symbols include `"$"`, `"€"`, and `"¥"`.
// 
// This property contains the same value returned by the [ObjectForKey] method
// when passing the [currencySymbol] key.
//
// [currencySymbol]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/currencySymbol
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/currencySymbol
func (l NSLocale) CurrencySymbol() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("currencySymbol"))
	return NSStringFromID(rv).String()
}

// The calendar identifier for the locale.
//
// # Discussion
// 
// Possible values are listed in `Calendar Identifiers`.
// 
// Use [LocalizedStringForCalendarIdentifier] to obtain a version of the value
// suitable for display to the user.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/calendarIdentifier
func (l NSLocale) CalendarIdentifier() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("calendarIdentifier"))
	return NSStringFromID(rv).String()
}

// The begin quotation symbol for the locale.
//
// # Discussion
// 
// Examples of begin quotation symbols include `"“"`, `"„"`, `"«"`, and
// `"「"`.
// 
// This property contains the same value returned by the [ObjectForKey] method
// when passing the [quotationBeginDelimiterKey] key.
//
// [quotationBeginDelimiterKey]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/quotationBeginDelimiterKey
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/quotationBeginDelimiter
func (l NSLocale) QuotationBeginDelimiter() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("quotationBeginDelimiter"))
	return NSStringFromID(rv).String()
}

// The end quotation symbol for the locale.
//
// # Discussion
// 
// Examples of end quotation symbols include `"”"`, `"“"`, `"»"`, and
// `"」"`.
// 
// This property contains the same value returned by the [ObjectForKey] method
// when passing the [quotationEndDelimiterKey] key.
//
// [quotationEndDelimiterKey]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/quotationEndDelimiterKey
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/quotationEndDelimiter
func (l NSLocale) QuotationEndDelimiter() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("quotationEndDelimiter"))
	return NSStringFromID(rv).String()
}

// The alternate begin quotation symbol for the locale.
//
// # Discussion
// 
// Examples of alternate begin quotation symbols include `"‘"`, `"‹"`, and
// `"『"`.
// 
// This property contains the same value returned by the [ObjectForKey] method
// when passing the [alternateQuotationBeginDelimiterKey] key.
//
// [alternateQuotationBeginDelimiterKey]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/alternateQuotationBeginDelimiterKey
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/alternateQuotationBeginDelimiter
func (l NSLocale) AlternateQuotationBeginDelimiter() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("alternateQuotationBeginDelimiter"))
	return NSStringFromID(rv).String()
}

// The alternate end quotation symbol for the locale.
//
// # Discussion
// 
// Examples of alternate end quotation symbols include `"’"`, `"›"`, and
// `"』"`.
// 
// This property contains the same value returned by the [ObjectForKey] method
// when passing the [alternateQuotationEndDelimiterKey] key.
//
// [alternateQuotationEndDelimiterKey]: https://developer.apple.com/documentation/Foundation/NSLocale/Key/alternateQuotationEndDelimiterKey
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/alternateQuotationEndDelimiter
func (l NSLocale) AlternateQuotationEndDelimiter() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("alternateQuotationEndDelimiter"))
	return NSStringFromID(rv).String()
}

// Returns the identifier for the language part of the locale. For example,
// returns “en-US” for “en_US@rg=gbzzzz” locale.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/languageIdentifier
func (l NSLocale) LanguageIdentifier() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("languageIdentifier"))
	return NSStringFromID(rv).String()
}

// Returns the region code of the locale. If the `rg` subtag is present, the
// value of the subtag will be used. For example, returns “GB” for
// “en_US@rg=gbzzzz” locale. If the `localeIdentifier` doesn’t contain a
// region, returns `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/regionCode
func (l NSLocale) RegionCode() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("regionCode"))
	return NSStringFromID(rv).String()
}

// The locale for the receiver.
//
// See: https://developer.apple.com/documentation/foundation/dateformatter/locale
func (l NSLocale) Locale() INSLocale {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("locale"))
	return NSLocaleFromID(objc.ID(rv))
}
func (l NSLocale) SetLocale(value INSLocale) {
	objc.Send[struct{}](l.ID, objc.Sel("setLocale:"), value)
}

// A locale which tracks the user’s current preferences.
//
// # Discussion
// 
// This value represents the locale currently used by the app, based on the
// following:
// 
// - The current system locale. - Any app-specific locale choice made in the
// Settings app. - The availability of the preferred locale in the app. For
// example, if the person using an app has set their device to use a
// Spanish-language locale, but the app only supports English, this value
// returns an English locale.
// 
// Use this property when you want a locale that always reflects the latest
// configuration settings. When the person using the app changes settings,
// reading properties from a locale instance obtained from this property
// provides the latest values. If you need to rely on a locale that does not
// change, use the locale given by the [CurrentLocale] property instead.
// 
// Although the locale obtained here automatically follows the latest region
// settings, it provides no indication when the settings change. To receive
// notification of locale changes, add your object as an observer of
// [currentLocaleDidChangeNotification].
//
// [currentLocaleDidChangeNotification]: https://developer.apple.com/documentation/Foundation/NSLocale/currentLocaleDidChangeNotification
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/autoupdatingCurrent
func (_NSLocaleClass NSLocaleClass) AutoupdatingCurrentLocale() NSLocale {
	rv := objc.Send[objc.ID](objc.ID(_NSLocaleClass.class), objc.Sel("autoupdatingCurrentLocale"))
	return NSLocaleFromID(objc.ID(rv))
}

// A locale that represents the user’s region settings at the time the
// property is read.
//
// # Discussion
// 
// This value represents the locale currently used by the app, based on the
// following:
// 
// - The current system locale. - Any app-specific locale choice made in the
// Settings app. - The availability of the preferred locale in the app. For
// example, if the person using an app has set their device to use a
// Spanish-language locale, but the app only supports English, this value
// returns an English locale.
// 
// Use this property when you need to rely on a consistent locale. A locale
// instance obtained this way does not change even when the person using the
// device changes region settings. If you want a locale instance that always
// reflects the current configuration, use the one provided by the
// [AutoupdatingCurrentLocale] property instead.
// 
// To receive notification of locale changes, add your object as an observer
// of [currentLocaleDidChangeNotification].
//
// [currentLocaleDidChangeNotification]: https://developer.apple.com/documentation/Foundation/NSLocale/currentLocaleDidChangeNotification
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/current
func (_NSLocaleClass NSLocaleClass) CurrentLocale() NSLocale {
	rv := objc.Send[objc.ID](objc.ID(_NSLocaleClass.class), objc.Sel("currentLocale"))
	return NSLocaleFromID(objc.ID(rv))
}

// A notification that indicates that the user’s locale changed.
//
// See: https://developer.apple.com/documentation/foundation/nslocale/currentlocaledidchangenotification
func (_NSLocaleClass NSLocaleClass) CurrentLocaleDidChangeNotification() NSNotificationName {
	rv := objc.Send[objc.ID](objc.ID(_NSLocaleClass.class), objc.Sel("NSCurrentLocaleDidChangeNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// A locale representing the generic root values with little localization.
//
// # Return Value
// 
// The generic locale that contains fixed “backstop” settings that provide
// values for otherwise undefined keys.
// 
// # Discussion
// 
// Use the system locale when you don’t want any localizations. If you want
// localizations that match the user’s region settings, use the locale given
// by the [CurrentLocale] or the [AutoupdatingCurrentLocale] property instead.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/system
func (_NSLocaleClass NSLocaleClass) SystemLocale() NSLocale {
	rv := objc.Send[objc.ID](objc.ID(_NSLocaleClass.class), objc.Sel("systemLocale"))
	return NSLocaleFromID(objc.ID(rv))
}

// The list of locale identifiers available on the system.
//
// # Discussion
// 
// A locale identifier starts with a language code, often includes a region
// code, and occasionally includes a script designator. See [Locale IDs] for
// more information about the structure of a locale identifier.
// 
// Use [LocalizedStringForLocaleIdentifier] to obtain a human readable
// description of any of the locale identifiers in this list.
//
// [Locale IDs]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPInternational/LanguageandLocaleIDs/LanguageandLocaleIDs.html#//apple_ref/doc/uid/10000171i-CH15-SW9
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/availableLocaleIdentifiers
func (_NSLocaleClass NSLocaleClass) AvailableLocaleIdentifiers() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSLocaleClass.class), objc.Sel("availableLocaleIdentifiers"))
	return objc.ConvertSliceToStrings(rv)
}

// The list of known country or region codes.
//
// # Discussion
// 
// Not all country or region codes have supporting locale data in the system.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/isoCountryCodes
func (_NSLocaleClass NSLocaleClass) ISOCountryCodes() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSLocaleClass.class), objc.Sel("ISOCountryCodes"))
	return objc.ConvertSliceToStrings(rv)
}

// The list of known language codes.
//
// # Discussion
// 
// A language code is a short string that represent a particular language. All
// languages have a three-character ISO 639-2 string, while some languages
// also have a two-character ISO 639-1 string. See [ISO 639.2 Codes for the
// Representation of Names of Languages] for the complete list of standardized
// language codes.
// 
// The array provided by this property contains a list of codes for all the
// languages the system knows about, designated by the ISO 639-1 code if
// available, or the ISO 639-2 code if not. Use the method
// [LocalizedStringForLanguageCode] to obtain a human readable string for any
// of the codes in the list.
// 
// For more information about language localization in your app, see [Language
// and Locale IDs].
// 
// Not all language codes have supporting locale data in the system.
//
// [ISO 639.2 Codes for the Representation of Names of Languages]: http://www.loc.gov/standards/iso639-2/php/English_list.php
// [Language and Locale IDs]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPInternational/LanguageandLocaleIDs/LanguageandLocaleIDs.html#//apple_ref/doc/uid/10000171i-CH15
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/isoLanguageCodes
func (_NSLocaleClass NSLocaleClass) ISOLanguageCodes() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSLocaleClass.class), objc.Sel("ISOLanguageCodes"))
	return objc.ConvertSliceToStrings(rv)
}

// The list of known currency codes.
//
// # Discussion
// 
// Not all currency codes have supporting locale data in the system.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/isoCurrencyCodes
func (_NSLocaleClass NSLocaleClass) ISOCurrencyCodes() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSLocaleClass.class), objc.Sel("ISOCurrencyCodes"))
	return objc.ConvertSliceToStrings(rv)
}

// A list of commonly encountered currency codes.
//
// # Discussion
// 
// Common codes may include, for example, `"AED"`, `"AUD"`, `"BZD"`, `"DKK"`,
// `"EUR"`, `"GBP"`, `"JPY"`, `"KES"`, `"MXN"`, `"OMR"`, `"STD"`, `"USD"`,
// `"XCD"`, and `"ZWD"`.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/commonISOCurrencyCodes
func (_NSLocaleClass NSLocaleClass) CommonISOCurrencyCodes() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSLocaleClass.class), objc.Sel("commonISOCurrencyCodes"))
	return objc.ConvertSliceToStrings(rv)
}

// An ordered list of the user’s preferred languages.
//
// # Discussion
// 
// Users choose a primary language when configuring a device, as described in
// [Reviewing Language and Region Settings]. They may also specify one or more
// secondary languages in order of preference for use when localization is
// unavailable in a higher priority language. Use this property to obtain the
// current user’s ordered list of languages, presented as an array of locale
// identifier strings.
// 
// For more information about language localization in your app, see [Language
// and Locale IDs].
//
// [Language and Locale IDs]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPInternational/LanguageandLocaleIDs/LanguageandLocaleIDs.html#//apple_ref/doc/uid/10000171i-CH15
// [Reviewing Language and Region Settings]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPInternational/SpecifyingPreferences/SpecifyingPreferences.html#//apple_ref/doc/uid/10000171i-CH12
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/preferredLanguages
func (_NSLocaleClass NSLocaleClass) PreferredLanguages() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSLocaleClass.class), objc.Sel("preferredLanguages"))
	return objc.ConvertSliceToStrings(rv)
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

