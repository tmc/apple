// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [PersonNameComponentsFormatter] class.
var (
	_PersonNameComponentsFormatterClass     PersonNameComponentsFormatterClass
	_PersonNameComponentsFormatterClassOnce sync.Once
)

func getPersonNameComponentsFormatterClass() PersonNameComponentsFormatterClass {
	_PersonNameComponentsFormatterClassOnce.Do(func() {
		_PersonNameComponentsFormatterClass = PersonNameComponentsFormatterClass{class: objc.GetClass("NSPersonNameComponentsFormatter")}
	})
	return _PersonNameComponentsFormatterClass
}

// GetPersonNameComponentsFormatterClass returns the class object for NSPersonNameComponentsFormatter.
func GetPersonNameComponentsFormatterClass() PersonNameComponentsFormatterClass {
	return getPersonNameComponentsFormatterClass()
}

type PersonNameComponentsFormatterClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (pc PersonNameComponentsFormatterClass) Alloc() PersonNameComponentsFormatter {
	rv := objc.Send[PersonNameComponentsFormatter](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}







// A formatter that provides localized representations of the components of a
// person’s name.
//
// # Overview
// 
// Each locale has its own set of rules and conventions for how personal names
// are structured and represented. These rules vary widely across different
// locales in a several ways, including the sort and display order of given
// and family names, the use of salutations and honorifics, and other concerns
// related to the grammar, spelling, punctuation, and formatting. About the
// only thing that consistent across all locales is that personal names are
// significant and meaningful. For this reason, names deserve careful and
// respectful treatment—perhaps more than any other kind of information your
// app interacts with.
// 
// Formatters can be configured to represent names in a variety of styles,
// which are described in detail below.
// 
// - Default ([PersonNameComponentsFormatterStyleDefault]) - Short
// ([PersonNameComponentsFormatterStyleShort]) - Long
// ([PersonNameComponentsFormatterStyleLong]) - Abbreviated
// ([PersonNameComponentsFormatterStyleAbbreviated])
// 
// When determining how to represent a name in a particular style, a formatter
// takes a number of factors into consideration, in order of priority:
// 
// - Scripts may specify a strict sort or display order of given and family
// names, and the availability of styles. - Users can enable and configure the
// display of short names, as well as whether or not to display nicknames when
// available. Users can also override the default sort and display order of
// given and family names for their current locale. - Locales specify a
// default sort and display order for given and family names. - The style
// property value set for the [NSPersonNameComponentsFormatter] object.
// 
// When the behavior specified in one factor conflicts with any other factors,
// the behavior specified by the factor with the most precedence is used. For
// example, the U.S. English (`en-US`) locale specifies that names be
// displayed in “given name followed by the family name” (for
// example,“John Appleseed”). This behavior would be overridden if the
// user changed their system preferences to have names displayed as family
// name followed by given name (for example, “Appleseed, John”), because
// user-specified preferences take precedence over locale-derived defaults.
// Furthermore, if the name to be formatted were Japanese (for example, given
// name: “泰夫”, family name: “木田”), the behavior derived for
// the name’s script (CJK, for Chinese, Japanese, and Korean languages)
// would take precedence over any locale-derived defaults or user-specified
// preferences to have the name displayed as family name followed by given
// name (for example, “木田 泰夫”).
// 
// These considerations extend to the availability of certain formatter styles
// as well. Because developer-specified configurations have the lowest
// precedence in determining behavior, the value set for the formatter’s
// style property can be invalidated if it’s not supported for the locale,
// user preferences, or script. If the specified style is not available, the
// next longest valid style is used. For example, a name in Arabic script (for
// example, “أحمد الراجحي”) does not support the Abbreviated
// style, so the Short style is used instead. A name that contains more than
// one script (for example, given name: “John”, family name: “王”) is
// detected to have “Unknown” script, which has its own set of behaviors
// and characteristics.
// 
// # Styles
// 
// [NSPersonNameComponentsFormatter] can be configured to format names in the
// following styles:
// 
// [PersonNameComponentsFormatterStyleDefault]: The minimally necessary
// features for differentiation in a casual setting. Equivalent to
// [PersonNameComponentsFormatterStyleMedium].
// [PersonNameComponentsFormatterStyleShort]: Relies on user preferences and
// language defaults to display shortened form appropriate for display in
// space-constrained settings. [PersonNameComponentsFormatterStyleLong]: The
// fully qualified name complete with all known components.
// [PersonNameComponentsFormatterStyleAbbreviated]: The maximally abbreviated
// form of a name.
// 
// [Table data omitted]
// 
// # Default
// 
// The Default, or Medium, style presents names in a way that is suitable for
// most contexts. It uses the given and family names, as well as a nickname,
// if provided and enabled by the user in System Preferences.
// 
// [Table data omitted]
// 
// # Short
// 
// The Short style offers an alternative display method for names whose
// default representation may exceed a certain length constraint. It is only
// available if the user has enabled “Short Names” in System Preferences,
// and only for names with a script that supports Short style. Otherwise, a
// formatter configured to display with Short style is displayed with Medium
// style instead.
// 
// If a user has enabled the use of short names, the user can choose from one
// of four variations:
// 
// - Given Name - Family Initial - Family Name - Given Initial - Given Name
// Only - Family Name Only
// 
// Short style is not available for names in CJK script and is restricted to
// Given Name Only or Family Name Only for names in Arabic or Devanagari
// script. If the specified Short style is unavailable, the Medium style is
// used instead.
// 
// [Table data omitted]
// 
// # Long
// 
// The Long style provides the most explicit representation of names. It uses
// all available name components, with the exception of nickname.
// 
// [Table data omitted]
// 
// # Abbreviated
// 
// The Abbreviated style offers the most compact representation of names,
// similar to a monogram.
// 
// Abbreviated style is supported for names in several scripts, with the
// following general characteristics:
// 
// - For names in Cyrillic, Greek, or Latin script, the first characters of
// `givenName`, `middleName`, and `familyName` may be used. - For names in
// Chinese or Japanese script, `familyName` may be used. If `familyName` is
// too long, or if the family name is `nil`, the Short or Medium style may be
// used instead. - For names in Korean script, `givenName` may be used. If
// `givenName` is too long, the first character of `givenName` may be used. If
// `givenName` is `nil`, the `familyName` may be used instead. - For names in
// Bengali, Devanagari, Gujarati, Gurmukhi, Kannada, Malayalam, Oriya,
// Sinhala, Tamil, Telugu, Tibetan, or Thai script, the first character of
// `givenName` may be used. If `givenName` is `nil`, the first character of
// `familyName` may be used instead. - For names that contain more than one
// script, the abbreviated style may use the `familyName`, `givenName`, or the
// first characters of `givenName` and/or `familyName`.
// 
// If the Abbreviated style is unavailable, the Short style is used
// instead—unless that too is unsupported, in which case the Medium style is
// used instead.
// 
// [Table data omitted]
//
// # Configuring Formatter Behavior
//
//   - [PersonNameComponentsFormatter.Style]: The formatting style of the receiver.
//   - [PersonNameComponentsFormatter.SetStyle]
//   - [PersonNameComponentsFormatter.Phonetic]: A Boolean value that specifies whether the receiver should use only the phonetic representations of name components.
//   - [PersonNameComponentsFormatter.SetPhonetic]
//
// # Converting Between Person Name Components and Strings
//
//   - [PersonNameComponentsFormatter.StringFromPersonNameComponents]: Returns a string formatted for a given [NSPersonNameComponents] object.
//   - [PersonNameComponentsFormatter.AnnotatedStringFromPersonNameComponents]: Returns an attributed string formatted for a given [NSPersonNameComponents] object, with attribute annotations for each component.
//   - [PersonNameComponentsFormatter.PersonNameComponentsFromString]: Returns a person name components object from a given string.
//
// # Instance Properties
//
//   - [PersonNameComponentsFormatter.Locale]
//   - [PersonNameComponentsFormatter.SetLocale]
//
// See: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter
type PersonNameComponentsFormatter struct {
	NSFormatter
}

// PersonNameComponentsFormatterFromID constructs a [PersonNameComponentsFormatter] from an objc.ID.
//
// A formatter that provides localized representations of the components of a
// person’s name.
func PersonNameComponentsFormatterFromID(id objc.ID) PersonNameComponentsFormatter {
	return NSPersonNameComponentsFormatter{NSFormatter: NSFormatterFromID(id)}
}

// NSPersonNameComponentsFormatterFromID is an alias for [PersonNameComponentsFormatterFromID] for cross-framework compatibility.
func NSPersonNameComponentsFormatterFromID(id objc.ID) PersonNameComponentsFormatter { return PersonNameComponentsFormatterFromID(id) }
// NOTE: PersonNameComponentsFormatter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [PersonNameComponentsFormatter] class.
//
// # Configuring Formatter Behavior
//
//   - [IPersonNameComponentsFormatter.Style]: The formatting style of the receiver.
//   - [IPersonNameComponentsFormatter.SetStyle]
//   - [IPersonNameComponentsFormatter.Phonetic]: A Boolean value that specifies whether the receiver should use only the phonetic representations of name components.
//   - [IPersonNameComponentsFormatter.SetPhonetic]
//
// # Converting Between Person Name Components and Strings
//
//   - [IPersonNameComponentsFormatter.StringFromPersonNameComponents]: Returns a string formatted for a given [NSPersonNameComponents] object.
//   - [IPersonNameComponentsFormatter.AnnotatedStringFromPersonNameComponents]: Returns an attributed string formatted for a given [NSPersonNameComponents] object, with attribute annotations for each component.
//   - [IPersonNameComponentsFormatter.PersonNameComponentsFromString]: Returns a person name components object from a given string.
//
// # Instance Properties
//
//   - [IPersonNameComponentsFormatter.Locale]
//   - [IPersonNameComponentsFormatter.SetLocale]
//
// See: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter
type IPersonNameComponentsFormatter interface {
	INSFormatter

	// Topic: Configuring Formatter Behavior

	// The formatting style of the receiver.
	Style() NSPersonNameComponentsFormatterStyle
	SetStyle(value NSPersonNameComponentsFormatterStyle)
	// A Boolean value that specifies whether the receiver should use only the phonetic representations of name components.
	Phonetic() bool
	SetPhonetic(value bool)

	// Topic: Converting Between Person Name Components and Strings

	// Returns a string formatted for a given [NSPersonNameComponents] object.
	StringFromPersonNameComponents(components INSPersonNameComponents) string
	// Returns an attributed string formatted for a given [NSPersonNameComponents] object, with attribute annotations for each component.
	AnnotatedStringFromPersonNameComponents(components INSPersonNameComponents) INSAttributedString
	// Returns a person name components object from a given string.
	PersonNameComponentsFromString(string_ string) INSPersonNameComponents

	// Topic: Instance Properties

	Locale() INSLocale
	SetLocale(value INSLocale)
}





// Init initializes the instance.
func (p PersonNameComponentsFormatter) Init() PersonNameComponentsFormatter {
	rv := objc.Send[PersonNameComponentsFormatter](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PersonNameComponentsFormatter) Autorelease() PersonNameComponentsFormatter {
	rv := objc.Send[PersonNameComponentsFormatter](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersonNameComponentsFormatter creates a new PersonNameComponentsFormatter instance.
func NewPersonNameComponentsFormatter() PersonNameComponentsFormatter {
	class := getPersonNameComponentsFormatterClass()
	rv := objc.Send[PersonNameComponentsFormatter](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewPersonNameComponentsFormatterWithCoder(coder INSCoder) PersonNameComponentsFormatter {
	instance := getPersonNameComponentsFormatterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return PersonNameComponentsFormatterFromID(rv)
}







// Returns a string formatted for a given [NSPersonNameComponents] object.
//
// components: The name components to be formatted.
//
// # Return Value
// 
// A formatted string representation of the given name components.
//
// See: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/string(from:)
func (p PersonNameComponentsFormatter) StringFromPersonNameComponents(components INSPersonNameComponents) string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("stringFromPersonNameComponents:"), components)
	return NSStringFromID(rv).String()
}

// Returns an attributed string formatted for a given [NSPersonNameComponents]
// object, with attribute annotations for each component.
//
// components: A formatted string representation of the given name components.
//
// # Return Value
// 
// An attributed string representation of the given name components. You can
// determine the person component corresponding to a particular range of the
// attributed string by querying the [NSPersonNameComponentKey] attribute,
// providing one of the [NSPersonNameComponent] constant values defined below
// as the attribute value.
//
// # Discussion
// 
// Use this method to style individual components of a formatted name, such as
// a name in a label.
//
// See: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/annotatedString(from:)
func (p PersonNameComponentsFormatter) AnnotatedStringFromPersonNameComponents(components INSPersonNameComponents) INSAttributedString {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("annotatedStringFromPersonNameComponents:"), components)
	return NSAttributedStringFromID(rv)
}

// Returns a person name components object from a given string.
//
// string: A string that is parsed to create a person name components object.
//
// # Return Value
// 
// An [NSPersonNameComponents] object parsing string using the receiver’s
// format, or `nil` if no components could be parsed.
//
// # Discussion
// 
// This method uses a combination of locale rules and heuristics to determine
// the most likely name components for a particular string representation.
// Parsing name components from a representation created for an existing name
// components object may not produce equivalent results.
// 
// Here are some general rules that describe the name component parsing
// behavior:
// 
// - Names in Latin script have components delimited by whitespace. - Names
// with a single delimited component are parsed into their most likely name
// component. - Names in Latin script with more than two delimited components
// may include middle components in the `givenName`, `middleName`, or
// `familyName` name components. - Names in Latin script that are inverted may
// be parsed into components in a different order than they appear; names in
// CJK script that are inverted will not typically produce the correct
// results. - Names in Latin script may use a comma to indicate name
// inversion. - Names in Latin script have capitalization preserved between
// string representation and parsed components. - Text between parentheses or
// brackets, as well as extraneous characters in names is ignored.
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/personNameComponents(from:)
func (p PersonNameComponentsFormatter) PersonNameComponentsFromString(string_ string) INSPersonNameComponents {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("personNameComponentsFromString:"), objc.String(string_))
	return NSPersonNameComponentsFromID(rv)
}





// Returns a string formatted for a given [NSPersonNameComponents] object
// using the provided style and options.
//
// components: The name components to be formatted.
//
// nameFormatStyle: A format style for the name components. For possible values, see
// [PersonNameComponentsFormatter.Style].
// //
// [PersonNameComponentsFormatter.Style]: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/Style-swift.enum
//
// nameOptions: The formatting options for the name components. For possible values, see
// [PersonNameComponentsFormatter.Options].
// //
// [PersonNameComponentsFormatter.Options]: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/Options
//
// # Return Value
// 
// A formatted string representation of the given name components.
//
// # Discussion
// 
// This method is a convenience for formatting name components without
// creating an instance of [NSPersonNameComponentsFormatter]. For greater
// customizability, you can create an instance of
// [NSPersonNameComponentsFormatter] and use [StringFromPersonNameComponents]
// instead.
//
// See: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/localizedString(from:style:options:)
func (_PersonNameComponentsFormatterClass PersonNameComponentsFormatterClass) LocalizedStringFromPersonNameComponentsStyleOptions(components INSPersonNameComponents, nameFormatStyle NSPersonNameComponentsFormatterStyle, nameOptions NSPersonNameComponentsFormatterOptions) string {
	rv := objc.Send[objc.ID](objc.ID(_PersonNameComponentsFormatterClass.class), objc.Sel("localizedStringFromPersonNameComponents:style:options:"), components, nameFormatStyle, nameOptions)
	return NSStringFromID(rv).String()
}








// The formatting style of the receiver.
//
// # Discussion
// 
// Styles specify which name components are used to create a string
// representation, and how. Examples of name components formatter styles
// include `long` and `abbreviated`.
//
// See: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/style-swift.property
func (p PersonNameComponentsFormatter) Style() NSPersonNameComponentsFormatterStyle {
	rv := objc.Send[NSPersonNameComponentsFormatterStyle](p.ID, objc.Sel("style"))
	return NSPersonNameComponentsFormatterStyle(rv)
}
func (p PersonNameComponentsFormatter) SetStyle(value NSPersonNameComponentsFormatterStyle) {
	objc.Send[struct{}](p.ID, objc.Sel("setStyle:"), value)
}



// A Boolean value that specifies whether the receiver should use only the
// phonetic representations of name components.
//
// # Discussion
// 
// The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/isPhonetic
func (p PersonNameComponentsFormatter) Phonetic() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isPhonetic"))
	return rv
}
func (p PersonNameComponentsFormatter) SetPhonetic(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setPhonetic:"), value)
}



// See: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/locale
func (p PersonNameComponentsFormatter) Locale() INSLocale {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("locale"))
	return NSLocaleFromID(objc.ID(rv))
}
func (p PersonNameComponentsFormatter) SetLocale(value INSLocale) {
	objc.Send[struct{}](p.ID, objc.Sel("setLocale:"), value)
}



























