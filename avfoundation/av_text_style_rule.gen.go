// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVTextStyleRule] class.
var (
	_AVTextStyleRuleClass     AVTextStyleRuleClass
	_AVTextStyleRuleClassOnce sync.Once
)

func getAVTextStyleRuleClass() AVTextStyleRuleClass {
	_AVTextStyleRuleClassOnce.Do(func() {
		_AVTextStyleRuleClass = AVTextStyleRuleClass{class: objc.GetClass("AVTextStyleRule")}
	})
	return _AVTextStyleRuleClass
}

// GetAVTextStyleRuleClass returns the class object for AVTextStyleRule.
func GetAVTextStyleRuleClass() AVTextStyleRuleClass {
	return getAVTextStyleRuleClass()
}

type AVTextStyleRuleClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVTextStyleRuleClass) Alloc() AVTextStyleRule {
	rv := objc.Send[AVTextStyleRule](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the text styling rules to apply to a media
// item’s textual content.
//
// # Overview
// 
// You use text style objects to format subtitles, closed captions, and other
// text-related content of the item. The system applies these rules to all or
// part of the text of the media item.
//
// # Creating and initializing style rules
//
//   - [AVTextStyleRule.InitWithTextMarkupAttributes]: Creates a text style rule object with the specified style attributes.
//   - [AVTextStyleRule.InitWithTextMarkupAttributesTextSelector]: Creates a text style rule object with the specified style attributes and text range information.
//
// # Accessing the style attributes
//
//   - [AVTextStyleRule.TextMarkupAttributes]: A dictionary of text style attributes to apply to the text.
//   - [AVTextStyleRule.TextSelector]: A string that identifies the text to which the attributes should apply.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule
type AVTextStyleRule struct {
	objectivec.Object
}

// AVTextStyleRuleFromID constructs a [AVTextStyleRule] from an objc.ID.
//
// An object that represents the text styling rules to apply to a media
// item’s textual content.
func AVTextStyleRuleFromID(id objc.ID) AVTextStyleRule {
	return AVTextStyleRule{objectivec.Object{ID: id}}
}
// NOTE: AVTextStyleRule adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVTextStyleRule] class.
//
// # Creating and initializing style rules
//
//   - [IAVTextStyleRule.InitWithTextMarkupAttributes]: Creates a text style rule object with the specified style attributes.
//   - [IAVTextStyleRule.InitWithTextMarkupAttributesTextSelector]: Creates a text style rule object with the specified style attributes and text range information.
//
// # Accessing the style attributes
//
//   - [IAVTextStyleRule.TextMarkupAttributes]: A dictionary of text style attributes to apply to the text.
//   - [IAVTextStyleRule.TextSelector]: A string that identifies the text to which the attributes should apply.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule
type IAVTextStyleRule interface {
	objectivec.IObject

	// Topic: Creating and initializing style rules

	// Creates a text style rule object with the specified style attributes.
	InitWithTextMarkupAttributes(textMarkupAttributes foundation.INSDictionary) AVTextStyleRule
	// Creates a text style rule object with the specified style attributes and text range information.
	InitWithTextMarkupAttributesTextSelector(textMarkupAttributes foundation.INSDictionary, textSelector string) AVTextStyleRule

	// Topic: Accessing the style attributes

	// A dictionary of text style attributes to apply to the text.
	TextMarkupAttributes() foundation.INSDictionary
	// A string that identifies the text to which the attributes should apply.
	TextSelector() string

	// An array of text style rules that specify the formatting and presentation of Web Video Text Tracks (WebVTT) subtitles.
	TextStyleRules() IAVTextStyleRule
	SetTextStyleRules(value IAVTextStyleRule)
}

// Init initializes the instance.
func (t AVTextStyleRule) Init() AVTextStyleRule {
	rv := objc.Send[AVTextStyleRule](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t AVTextStyleRule) Autorelease() AVTextStyleRule {
	rv := objc.Send[AVTextStyleRule](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVTextStyleRule creates a new AVTextStyleRule instance.
func NewAVTextStyleRule() AVTextStyleRule {
	class := getAVTextStyleRuleClass()
	rv := objc.Send[AVTextStyleRule](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a text style rule object with the specified style attributes.
//
// textMarkupAttributes: A dictionary of style attributes. For a list of supported keys and values
// that you can include in this dictionary, see `CMTextMarkup.H()`.
//
// # Return Value
// 
// A text style rule object initialized with the specified attributes.
//
// # Discussion
// 
// This method sets the [TextSelector] property of the style object to `nil`,
// which causes the rules to be applied to all of the text in the media item.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule/init(textMarkupAttributes:)
func NewTextStyleRuleWithTextMarkupAttributes(textMarkupAttributes foundation.INSDictionary) AVTextStyleRule {
	instance := getAVTextStyleRuleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTextMarkupAttributes:"), textMarkupAttributes)
	return AVTextStyleRuleFromID(rv)
}

// Creates a text style rule object with the specified style attributes and
// text range information.
//
// textMarkupAttributes: A dictionary of style attributes. For a list of supported keys and values
// that you can include in this dictionary, see `CMTextMarkup.H()`.
//
// textSelector: A string contains an identifier for the ranges of text to which the style
// attributes should be applied. Eligible identifiers are determined by the
// media format and its corresponding text content. For example, the string
// could contain the CSS selectors used by the corresponding text in Web Video
// Text Tracks (WebVTT) markup. Specify `nil` if you want the style attributes
// to apply to all text in the item.
//
// # Return Value
// 
// A text style rule object initialized with the specified attributes and
// range information.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule/init(textMarkupAttributes:textSelector:)
func NewTextStyleRuleWithTextMarkupAttributesTextSelector(textMarkupAttributes foundation.INSDictionary, textSelector string) AVTextStyleRule {
	instance := getAVTextStyleRuleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTextMarkupAttributes:textSelector:"), textMarkupAttributes, objc.String(textSelector))
	return AVTextStyleRuleFromID(rv)
}

// Creates a text style rule object with the specified style attributes.
//
// textMarkupAttributes: A dictionary of style attributes. For a list of supported keys and values
// that you can include in this dictionary, see `CMTextMarkup.H()`.
//
// # Return Value
// 
// A text style rule object initialized with the specified attributes.
//
// # Discussion
// 
// This method sets the [TextSelector] property of the style object to `nil`,
// which causes the rules to be applied to all of the text in the media item.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule/init(textMarkupAttributes:)
func (t AVTextStyleRule) InitWithTextMarkupAttributes(textMarkupAttributes foundation.INSDictionary) AVTextStyleRule {
	rv := objc.Send[AVTextStyleRule](t.ID, objc.Sel("initWithTextMarkupAttributes:"), textMarkupAttributes)
	return rv
}
// Creates a text style rule object with the specified style attributes and
// text range information.
//
// textMarkupAttributes: A dictionary of style attributes. For a list of supported keys and values
// that you can include in this dictionary, see `CMTextMarkup.H()`.
//
// textSelector: A string contains an identifier for the ranges of text to which the style
// attributes should be applied. Eligible identifiers are determined by the
// media format and its corresponding text content. For example, the string
// could contain the CSS selectors used by the corresponding text in Web Video
// Text Tracks (WebVTT) markup. Specify `nil` if you want the style attributes
// to apply to all text in the item.
//
// # Return Value
// 
// A text style rule object initialized with the specified attributes and
// range information.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule/init(textMarkupAttributes:textSelector:)
func (t AVTextStyleRule) InitWithTextMarkupAttributesTextSelector(textMarkupAttributes foundation.INSDictionary, textSelector string) AVTextStyleRule {
	rv := objc.Send[AVTextStyleRule](t.ID, objc.Sel("initWithTextMarkupAttributes:textSelector:"), textMarkupAttributes, objc.String(textSelector))
	return rv
}

// Creates an array of text style rule objects from the specified
// property-list object.
//
// plist: A property-list object containing the text style data.
//
// # Return Value
// 
// An array of [AVTextStyleRule] objects corresponding to the style
// information in the property-list object.
//
// # Discussion
// 
// Use this method to create new text style rule objects based on data you
// previously converted to a property-list format using the
// [PropertyListForTextStyleRules] class method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule/textStyleRules(fromPropertyList:)
func (_AVTextStyleRuleClass AVTextStyleRuleClass) TextStyleRulesFromPropertyList(plist objectivec.IObject) []AVTextStyleRule {
	rv := objc.Send[[]objc.ID](objc.ID(_AVTextStyleRuleClass.class), objc.Sel("textStyleRulesFromPropertyList:"), plist)
	return objc.ConvertSlice(rv, func(id objc.ID) AVTextStyleRule {
		return AVTextStyleRuleFromID(id)
	})
}
// Converts one or more text style rules into a serializable property list
// object.
//
// textStyleRules: An array of [AVTextStyleRule] objects to write to the property list.
//
// # Return Value
// 
// A property-list object that you can pass to the [PropertyListSerialization]
// serialization routines.
//
// [PropertyListSerialization]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization
//
// # Discussion
// 
// The property-list object returned by this method can be written to disk and
// stored persistently.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule/propertyList(for:)
func (_AVTextStyleRuleClass AVTextStyleRuleClass) PropertyListForTextStyleRules(textStyleRules []AVTextStyleRule) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVTextStyleRuleClass.class), objc.Sel("propertyListForTextStyleRules:"), objectivec.IObjectSliceToNSArray(textStyleRules))
	return objectivec.Object{ID: rv}
}
// Creates a new text style rule object using the style attributes in the
// specified dictionary.
//
// textMarkupAttributes: A dictionary of style attributes. For a list of supported keys and values
// that you can include in this dictionary, see `CMTextMarkup.H()`.
//
// # Return Value
// 
// A new text style rule object with the specified attributes.
//
// # Discussion
// 
// This method sets the [TextSelector] property of the style object to `nil`,
// which causes the rules to be applied to all of the text in the media item.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule/textStyleRuleWithTextMarkupAttributes:
func (_AVTextStyleRuleClass AVTextStyleRuleClass) TextStyleRuleWithTextMarkupAttributes(textMarkupAttributes foundation.INSDictionary) AVTextStyleRule {
	rv := objc.Send[objc.ID](objc.ID(_AVTextStyleRuleClass.class), objc.Sel("textStyleRuleWithTextMarkupAttributes:"), textMarkupAttributes)
	return AVTextStyleRuleFromID(rv)
}
// Creates a new text style rule object using the specified style attributes
// and text range information.
//
// textMarkupAttributes: A dictionary of style attributes. For a list of supported keys and values
// that you can include in this dictionary, see `CMTextMarkup.H()`.
//
// textSelector: A string contains an identifier for the ranges of text to which the style
// attributes should be applied. Eligible identifiers are determined by the
// media format and its corresponding text content. For example, the string
// could contain the CSS selectors used by the corresponding text in Web Video
// Text Tracks (WebVTT) markup. Specify `nil` if you want the style attributes
// to apply to all text in the item.
//
// # Return Value
// 
// A new text style rule object with the specified attributes and range
// information.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule/textStyleRuleWithTextMarkupAttributes:textSelector:
func (_AVTextStyleRuleClass AVTextStyleRuleClass) TextStyleRuleWithTextMarkupAttributesTextSelector(textMarkupAttributes foundation.INSDictionary, textSelector string) AVTextStyleRule {
	rv := objc.Send[objc.ID](objc.ID(_AVTextStyleRuleClass.class), objc.Sel("textStyleRuleWithTextMarkupAttributes:textSelector:"), textMarkupAttributes, objc.String(textSelector))
	return AVTextStyleRuleFromID(rv)
}

// A dictionary of text style attributes to apply to the text.
//
// # Discussion
// 
// The supported keys for this dictionary are defined in `CMTextMarkup.H()`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule/textMarkupAttributes
func (t AVTextStyleRule) TextMarkupAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textMarkupAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// A string that identifies the text to which the attributes should apply.
//
// # Discussion
// 
// The contents of the string are determined by the format of the legible
// media. For example, the string could contain the CSS selectors used by the
// corresponding text in Web Video Text Tracks (WebVTT) markup.
// 
// If the value of this property is `nil`, the text style attributes apply to
// all text in the media item.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVTextStyleRule/textSelector
func (t AVTextStyleRule) TextSelector() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textSelector"))
	return foundation.NSStringFromID(rv).String()
}
// An array of text style rules that specify the formatting and presentation
// of Web Video Text Tracks (WebVTT) subtitles.
//
// See: https://developer.apple.com/documentation/avfoundation/avplayeritem/textstylerules
func (t AVTextStyleRule) TextStyleRules() IAVTextStyleRule {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textStyleRules"))
	return AVTextStyleRuleFromID(objc.ID(rv))
}
func (t AVTextStyleRule) SetTextStyleRules(value IAVTextStyleRule) {
	objc.Send[struct{}](t.ID, objc.Sel("setTextStyleRules:"), value)
}

