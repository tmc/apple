// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextAlternatives] class.
var (
	_NSTextAlternativesClass     NSTextAlternativesClass
	_NSTextAlternativesClassOnce sync.Once
)

func getNSTextAlternativesClass() NSTextAlternativesClass {
	_NSTextAlternativesClassOnce.Do(func() {
		_NSTextAlternativesClass = NSTextAlternativesClass{class: objc.GetClass("NSTextAlternatives")}
	})
	return _NSTextAlternativesClass
}

// GetNSTextAlternativesClass returns the class object for NSTextAlternatives.
func GetNSTextAlternativesClass() NSTextAlternativesClass {
	return getNSTextAlternativesClass()
}

type NSTextAlternativesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTextAlternativesClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextAlternativesClass) Alloc() NSTextAlternatives {
	rv := objc.Send[NSTextAlternatives](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A list of alternative strings for a piece of text.
//
// # Overview
// 
// [NSTextAlternatives] is an immutable value class that stores a list of
// alternatives for a piece of text and communicates the user’s selection of
// an alternative via a notification to your app. To support dictation, for
// example, you might use [NSTextAlternatives] to present a list of
// alternative interpretations for a word or phrase the user speaks. If the
// user chooses to replace the initial interpretation with an alternative,
// [NSTextAlternatives] notifies you of the choice so that you can update the
// text appropriately.
// 
// [NSTextAlternatives] instances are attached to attributed strings as the
// value of a text attribute, [NSTextAlternativesAttributeName].
//
// [NSTextAlternativesAttributeName]: https://developer.apple.com/documentation/AppKit/NSTextAlternativesAttributeName
//
// # Initializing a Text Alternatives Object
//
//   - [NSTextAlternatives.InitWithPrimaryStringAlternativeStrings]: Initializes an [NSTextAlternatives] instance.
//
// # Storing Alternative Text Strings
//
//   - [NSTextAlternatives.PrimaryString]: The text that was initially chosen as the input string.
//   - [NSTextAlternatives.AlternativeStrings]: An array of alternative possible interpretations that the user might select.
//
// # Selecting an Alternative String
//
//   - [NSTextAlternatives.NoteSelectedAlternativeString]: Sent to the [NSTextAlternatives] object by the text view when the user chooses one of the alternative strings.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAlternatives
type NSTextAlternatives struct {
	objectivec.Object
}

// NSTextAlternativesFromID constructs a [NSTextAlternatives] from an objc.ID.
//
// A list of alternative strings for a piece of text.
func NSTextAlternativesFromID(id objc.ID) NSTextAlternatives {
	return NSTextAlternatives{objectivec.Object{ID: id}}
}
// NOTE: NSTextAlternatives adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextAlternatives] class.
//
// # Initializing a Text Alternatives Object
//
//   - [INSTextAlternatives.InitWithPrimaryStringAlternativeStrings]: Initializes an [NSTextAlternatives] instance.
//
// # Storing Alternative Text Strings
//
//   - [INSTextAlternatives.PrimaryString]: The text that was initially chosen as the input string.
//   - [INSTextAlternatives.AlternativeStrings]: An array of alternative possible interpretations that the user might select.
//
// # Selecting an Alternative String
//
//   - [INSTextAlternatives.NoteSelectedAlternativeString]: Sent to the [NSTextAlternatives] object by the text view when the user chooses one of the alternative strings.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAlternatives
type INSTextAlternatives interface {
	objectivec.IObject

	// Topic: Initializing a Text Alternatives Object

	// Initializes an [NSTextAlternatives] instance.
	InitWithPrimaryStringAlternativeStrings(primaryString string, alternativeStrings []string) NSTextAlternatives

	// Topic: Storing Alternative Text Strings

	// The text that was initially chosen as the input string.
	PrimaryString() string
	// An array of alternative possible interpretations that the user might select.
	AlternativeStrings() []string

	// Topic: Selecting an Alternative String

	// Sent to the [NSTextAlternatives] object by the text view when the user chooses one of the alternative strings.
	NoteSelectedAlternativeString(alternativeString string)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t NSTextAlternatives) Init() NSTextAlternatives {
	rv := objc.Send[NSTextAlternatives](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextAlternatives) Autorelease() NSTextAlternatives {
	rv := objc.Send[NSTextAlternatives](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextAlternatives creates a new NSTextAlternatives instance.
func NewNSTextAlternatives() NSTextAlternatives {
	class := getNSTextAlternativesClass()
	rv := objc.Send[NSTextAlternatives](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes an [NSTextAlternatives] instance.
//
// primaryString: The string that is initially chosen as the input string.
//
// alternativeStrings: An array of alternative possible interpretations that the user might
// select.
//
// # Return Value
// 
// An initialized [NSTextAlternatives] instance.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAlternatives/init(primaryString:alternativeStrings:)
func NewTextAlternativesWithPrimaryStringAlternativeStrings(primaryString string, alternativeStrings []string) NSTextAlternatives {
	instance := getNSTextAlternativesClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPrimaryString:alternativeStrings:"), objc.String(primaryString), objectivec.StringSliceToNSArray(alternativeStrings))
	return NSTextAlternativesFromID(rv)
}

// Initializes an [NSTextAlternatives] instance.
//
// primaryString: The string that is initially chosen as the input string.
//
// alternativeStrings: An array of alternative possible interpretations that the user might
// select.
//
// # Return Value
// 
// An initialized [NSTextAlternatives] instance.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAlternatives/init(primaryString:alternativeStrings:)
func (t NSTextAlternatives) InitWithPrimaryStringAlternativeStrings(primaryString string, alternativeStrings []string) NSTextAlternatives {
	rv := objc.Send[NSTextAlternatives](t.ID, objc.Sel("initWithPrimaryString:alternativeStrings:"), objc.String(primaryString), objectivec.StringSliceToNSArray(alternativeStrings))
	return rv
}
// Sent to the [NSTextAlternatives] object by the text view when the user
// chooses one of the alternative strings.
//
// alternativeString: The alternative string chosen by the user.
//
// # Discussion
// 
// The base class implementation sends a notification,
// [NSTextAlternativesSelectedAlternativeStringNotification], with the
// selected alternative string in the user info under the key
// `@"NSAlternativeString"`. Using this mechanism, arbitrary objects can
// listen for user selections of alternative strings.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAlternatives/noteSelectedAlternativeString(_:)
func (t NSTextAlternatives) NoteSelectedAlternativeString(alternativeString string) {
	objc.Send[objc.ID](t.ID, objc.Sel("noteSelectedAlternativeString:"), objc.String(alternativeString))
}
func (t NSTextAlternatives) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The text that was initially chosen as the input string.
//
// # Discussion
// 
// The text system uses the `primaryString` property to make sure that the
// text is still in the same state as when it was entered.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAlternatives/primaryString
func (t NSTextAlternatives) PrimaryString() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("primaryString"))
	return foundation.NSStringFromID(rv).String()
}
// An array of alternative possible interpretations that the user might
// select.
//
// # Discussion
// 
// The text system presents the alternative strings via a user interface
// similar to that used for spelling correction alternatives.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAlternatives/alternativeStrings
func (t NSTextAlternatives) AlternativeStrings() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("alternativeStrings"))
	return objc.ConvertSliceToStrings(rv)
}

