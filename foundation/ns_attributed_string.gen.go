// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAttributedString] class.
var (
	_NSAttributedStringClass     NSAttributedStringClass
	_NSAttributedStringClassOnce sync.Once
)

func getNSAttributedStringClass() NSAttributedStringClass {
	_NSAttributedStringClassOnce.Do(func() {
		_NSAttributedStringClass = NSAttributedStringClass{class: objc.GetClass("NSAttributedString")}
	})
	return _NSAttributedStringClass
}

// GetNSAttributedStringClass returns the class object for NSAttributedString.
func GetNSAttributedStringClass() NSAttributedStringClass {
	return getNSAttributedStringClass()
}

type NSAttributedStringClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAttributedStringClass) Alloc() NSAttributedString {
	rv := objc.Send[NSAttributedString](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A string of text that manages data, layout, and stylistic information for
// ranges of characters to support rendering.
//
// # Overview
// 
// [NSAttributedString] is a type you use to manage strings of stylized
// Unicode text. In addition to text, an attributed string contains key-value
// pairs known as that specify additional information to apply to ranges of
// characters within the string. Attributed strings support many different
// kinds of attributes, including:
// 
// - Rendering attributes that specify font, color, kern, ligature, and other
// details - Attributes for attachments and adaptive image glyphs - Semantic
// attributes such as link URLs or tool-tip information - Language attributes
// to support automatic gender agreement and text layout - Accessibility
// attributes that provide information for assistive technologies - Attributes
// that summarize details of the Markdown import process - Custom attributes
// you define for your app
// 
// Use attributed strings anywhere you need styled text, or when you need to
// associate additional information with your text. Because
// [NSAttributedString] is an immutable type, you specify all of the text and
// attributes for it at creation time and can’t change them later. You can
// create attributed strings directly from a string of characters and a
// dictionary of attributes. You can also create attributed strings from the
// contents of a file, including files that contain RTF, RTFD, HTML, Markdown,
// or other file formats. If you need to modify the contents of an attributed
// string later, use the [NSMutableAttributedString] type instead.
// 
// If you create an [NSAttributedString] without any font information, the
// string’s default font is Helvetica 12-point, which might differ from the
// default system font for the platform. To change the font, specify a font
// attribute at creation time.
// 
// # Persistence
// 
// Be aware of how you persist attributed strings to and from the disk. RTF
// and RTFD are the preferred format for attributed strings because they offer
// the best fidelity for reading and writing attribute data. The RTF formats
// support a large number of standard attributes, and Apple extends the
// formats to support many Apple-specific attributes. If you define custom
// attributes for ranges of characters, store them separately alongside the
// RTF file for your text.
// 
// If you work extensively with HTML content, validate the results and
// performance of import and export operations during testing. WebKit handles
// the conversion between HTML markup and attributed strings. If an HTML file
// contains tags or constructs that attributed strings don’t support, the
// import process ignores them and imports what it can.
// 
// When you create an attributed string from Markdown, the system adds
// presentation intent attributes with information about the original Markdown
// content. The system doesn’t add style attributes to match the Markdown
// elements, but the system applies default style information when it renders
// a string with intent attributes. To change the rendering behavior of your
// Markdown content, remove the intent attributes and add the style attributes
// you prefer.
// 
// The methods for reading and writing common file formats also support
// document attributes. Document attributes aren’t part of the attributed
// string itself, but accompany the text when you save it to a file. When you
// read a file, the system returns any document attributes that it finds.
// Similarly, when you write an attributed string to a file, you can specify
// the attributes to include. For more information about document attributes,
// see [NSAttributedString.DocumentAttributeKey] and
// [NSAttributedString.DocumentReadingOptionKey].
// 
// # System framework interoperability
// 
// [TextKit] and [Core Text] use attributed strings extensively during the
// layout and rendering processes. These technologies use the string’s text
// and rendering-related attributes to calculate the text metrics needed
// during layout. Similarly, these technologies apply those same attributes
// during rendering to give the text its styled appearance. The technologies
// use only attributes that directly affect the appearance of the text, and
// ignore most other attributes. For some attributes, the text system adds
// attributes during rendering as needed. For example, the text system
// provides default style attributes for text with the [link] attribute.
// 
// [AppKit] and [UIKit] also support attributed strings in several ways. Some
// views and controls in these frameworks have APIs that accept attributed
// strings, and render the string with its style information. The frameworks
// also add methods to the [NSAttributedString] class that let you draw a
// styled string directly in one of your custom views. Because these methods
// use TextKit to draw the string, they recognize the same rendering-related
// attributes as that technology.
// 
// The [NSAttributedString] class and its Core Foundation counterpart,
// [CFAttributedString], are toll-free bridged, which means you can use the
// two types interchangeably in your code without losing any text or attribute
// information.
//
// [AppKit]: https://developer.apple.com/documentation/AppKit
// [CFAttributedString]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedString
// [Core Text]: https://developer.apple.com/documentation/CoreText
// [NSAttributedString.DocumentAttributeKey]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey
// [NSAttributedString.DocumentReadingOptionKey]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentReadingOptionKey
// [TextKit]: https://developer.apple.com/documentation/UIKit/textkit
// [UIKit]: https://developer.apple.com/documentation/UIKit
// [link]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/link
//
// # Exporting the string as data
//
//   - [NSAttributedString.DataFromRangeDocumentAttributesError]: Returns a data object that contains a text stream corresponding to the characters and attributes within the specified range.
//   - [NSAttributedString.FileWrapperFromRangeDocumentAttributesError]: Returns a file wrapper object that contains a text stream corresponding to the characters and attributes within the specified range.
//   - [NSAttributedString.DocFormatFromRangeDocumentAttributes]: Returns a data object that contains a Microsoft Word–format stream corresponding to the characters and attributes within the specified range.
//   - [NSAttributedString.RTFFromRangeDocumentAttributes]: Returns a data object that contains an RTF stream corresponding to the characters and attributes within the specified range, omitting all attachment attributes.
//   - [NSAttributedString.RTFDFromRangeDocumentAttributes]: Returns a data object that contains an RTFD stream corresponding to the characters and attributes within the specified range.
//   - [NSAttributedString.RTFDFileWrapperFromRangeDocumentAttributes]: Returns a file wrapper object that contains an RTFD document corresponding to the characters and attributes within the specified range.
//
// # Getting the characters
//
//   - [NSAttributedString.String]: The character contents of the attributed string as a string.
//   - [NSAttributedString.Length]: The length of the attributed string.
//   - [NSAttributedString.AttributedSubstringFromRange]: Returns an attributed string consisting of the characters and attributes within the specified range in the attributed string.
//
// # Getting font attribute information
//
//   - [NSAttributedString.FontAttributesInRange]: Returns the font attributes in effect for the character at the specified location.
//   - [NSAttributedString.RulerAttributesInRange]: Returns the ruler (paragraph) attributes in effect for the characters within the specified range.
//
// # Getting attributes for a range of text
//
//   - [NSAttributedString.AttributesAtIndexEffectiveRange]: Returns the attributes for the character at the specified index.
//   - [NSAttributedString.AttributesAtIndexLongestEffectiveRangeInRange]: Returns the attributes for the character at the specified index and, by reference, the range where the attributes apply.
//   - [NSAttributedString.AttributeAtIndexEffectiveRange]: Returns the value for an attribute with the specified name of the character at the specified index and, by reference, the range where the attribute applies.
//   - [NSAttributedString.AttributeAtIndexLongestEffectiveRangeInRange]: Returns the value for the attribute with the specified name of the character at the specified index and, by reference, the range where the attribute applies.
//   - [NSAttributedString.EnumerateAttributeInRangeOptionsUsingBlock]: Executes the specified closure or block for each range of a particular attribute in the attributed string.
//
// # Comparing strings
//
//   - [NSAttributedString.IsEqualToAttributedString]: Returns a Boolean value that indicates whether the attributed string is equal to the specified string.
//
// # Getting the supported text-file formats
//
//   - [NSAttributedString.PrefersRTFDInRange]: Returns a Boolean value that indicates whether the specified range of text prefers RTFD formatting.
//
// # Calculating linguistic units
//
//   - [NSAttributedString.DoubleClickAtIndex]: Returns the range of characters that form a word (or other linguistic unit) surrounding the specified index, taking language characteristics into account.
//   - [NSAttributedString.LineBreakBeforeIndexWithinRange]: Returns the appropriate line break when the character at the index doesn’t fit on the same line as the character at the beginning of the range.
//   - [NSAttributedString.LineBreakByHyphenatingBeforeIndexWithinRange]: Returns the index of the closest character before the specified index, and within the specified range, that can fit on a new line by hyphenating.
//   - [NSAttributedString.NextWordFromIndexForward]: Returns the index of the first character of the word after or before the specified index.
//
// # Performing automatic grammar agreement
//
//   - [NSAttributedString.AttributedStringByInflectingString]: If the string has portions tagged with NSInflectionRuleAttributeName that have no format specifiers, create a new string with those portions inflected by following the rule in the attribute.
//
// # Calculating ranges for common elements
//
//   - [NSAttributedString.ItemNumberInTextListAtIndex]: Returns the index of the item at the specified location within the list.
//   - [NSAttributedString.RangeOfTextBlockAtIndex]: Returns the range of the individual text block that contains the specified location.
//   - [NSAttributedString.RangeOfTextListAtIndex]: Returns the range of the specified text list that contains the specified location.
//   - [NSAttributedString.RangeOfTextTableAtIndex]: Returns the range of the specified text table that contains the specified location.
//
// # Drawing the attributed string
//
//   - [NSAttributedString.DrawAtPoint]: Draws the attributed string starting at the specified point in the current graphics context.
//   - [NSAttributedString.DrawInRect]: Draws the attributed string inside the specified bounding rectangle in the current graphics context.
//   - [NSAttributedString.DrawWithRectOptionsContext]: Draws the attributed string in the specified bounding rectangle using the provided options.
//
// # Getting metrics for the string
//
//   - [NSAttributedString.Size]: Returns the size necessary to draw the string.
//   - [NSAttributedString.BoundingRectWithSizeOptionsContext]: Returns the bounding rectangle necessary to draw the string.
//   - [NSAttributedString.ContainsAttachmentsInRange]: Returns a Boolean value that indicates if the attributed string contains an attachment in the specified range.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString
type NSAttributedString struct {
	objectivec.Object
}

// NSAttributedStringFromID constructs a [NSAttributedString] from an objc.ID.
//
// A string of text that manages data, layout, and stylistic information for
// ranges of characters to support rendering.
func NSAttributedStringFromID(id objc.ID) NSAttributedString {
	return NSAttributedString{objectivec.Object{ID: id}}
}
// NOTE: NSAttributedString adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSAttributedString] class.
//
// # Exporting the string as data
//
//   - [INSAttributedString.DataFromRangeDocumentAttributesError]: Returns a data object that contains a text stream corresponding to the characters and attributes within the specified range.
//   - [INSAttributedString.FileWrapperFromRangeDocumentAttributesError]: Returns a file wrapper object that contains a text stream corresponding to the characters and attributes within the specified range.
//   - [INSAttributedString.DocFormatFromRangeDocumentAttributes]: Returns a data object that contains a Microsoft Word–format stream corresponding to the characters and attributes within the specified range.
//   - [INSAttributedString.RTFFromRangeDocumentAttributes]: Returns a data object that contains an RTF stream corresponding to the characters and attributes within the specified range, omitting all attachment attributes.
//   - [INSAttributedString.RTFDFromRangeDocumentAttributes]: Returns a data object that contains an RTFD stream corresponding to the characters and attributes within the specified range.
//   - [INSAttributedString.RTFDFileWrapperFromRangeDocumentAttributes]: Returns a file wrapper object that contains an RTFD document corresponding to the characters and attributes within the specified range.
//
// # Getting the characters
//
//   - [INSAttributedString.String]: The character contents of the attributed string as a string.
//   - [INSAttributedString.Length]: The length of the attributed string.
//   - [INSAttributedString.AttributedSubstringFromRange]: Returns an attributed string consisting of the characters and attributes within the specified range in the attributed string.
//
// # Getting font attribute information
//
//   - [INSAttributedString.FontAttributesInRange]: Returns the font attributes in effect for the character at the specified location.
//   - [INSAttributedString.RulerAttributesInRange]: Returns the ruler (paragraph) attributes in effect for the characters within the specified range.
//
// # Getting attributes for a range of text
//
//   - [INSAttributedString.AttributesAtIndexEffectiveRange]: Returns the attributes for the character at the specified index.
//   - [INSAttributedString.AttributesAtIndexLongestEffectiveRangeInRange]: Returns the attributes for the character at the specified index and, by reference, the range where the attributes apply.
//   - [INSAttributedString.AttributeAtIndexEffectiveRange]: Returns the value for an attribute with the specified name of the character at the specified index and, by reference, the range where the attribute applies.
//   - [INSAttributedString.AttributeAtIndexLongestEffectiveRangeInRange]: Returns the value for the attribute with the specified name of the character at the specified index and, by reference, the range where the attribute applies.
//   - [INSAttributedString.EnumerateAttributeInRangeOptionsUsingBlock]: Executes the specified closure or block for each range of a particular attribute in the attributed string.
//
// # Comparing strings
//
//   - [INSAttributedString.IsEqualToAttributedString]: Returns a Boolean value that indicates whether the attributed string is equal to the specified string.
//
// # Getting the supported text-file formats
//
//   - [INSAttributedString.PrefersRTFDInRange]: Returns a Boolean value that indicates whether the specified range of text prefers RTFD formatting.
//
// # Calculating linguistic units
//
//   - [INSAttributedString.DoubleClickAtIndex]: Returns the range of characters that form a word (or other linguistic unit) surrounding the specified index, taking language characteristics into account.
//   - [INSAttributedString.LineBreakBeforeIndexWithinRange]: Returns the appropriate line break when the character at the index doesn’t fit on the same line as the character at the beginning of the range.
//   - [INSAttributedString.LineBreakByHyphenatingBeforeIndexWithinRange]: Returns the index of the closest character before the specified index, and within the specified range, that can fit on a new line by hyphenating.
//   - [INSAttributedString.NextWordFromIndexForward]: Returns the index of the first character of the word after or before the specified index.
//
// # Performing automatic grammar agreement
//
//   - [INSAttributedString.AttributedStringByInflectingString]: If the string has portions tagged with NSInflectionRuleAttributeName that have no format specifiers, create a new string with those portions inflected by following the rule in the attribute.
//
// # Calculating ranges for common elements
//
//   - [INSAttributedString.ItemNumberInTextListAtIndex]: Returns the index of the item at the specified location within the list.
//   - [INSAttributedString.RangeOfTextBlockAtIndex]: Returns the range of the individual text block that contains the specified location.
//   - [INSAttributedString.RangeOfTextListAtIndex]: Returns the range of the specified text list that contains the specified location.
//   - [INSAttributedString.RangeOfTextTableAtIndex]: Returns the range of the specified text table that contains the specified location.
//
// # Drawing the attributed string
//
//   - [INSAttributedString.DrawAtPoint]: Draws the attributed string starting at the specified point in the current graphics context.
//   - [INSAttributedString.DrawInRect]: Draws the attributed string inside the specified bounding rectangle in the current graphics context.
//   - [INSAttributedString.DrawWithRectOptionsContext]: Draws the attributed string in the specified bounding rectangle using the provided options.
//
// # Getting metrics for the string
//
//   - [INSAttributedString.Size]: Returns the size necessary to draw the string.
//   - [INSAttributedString.BoundingRectWithSizeOptionsContext]: Returns the bounding rectangle necessary to draw the string.
//   - [INSAttributedString.ContainsAttachmentsInRange]: Returns a Boolean value that indicates if the attributed string contains an attachment in the specified range.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString
type INSAttributedString interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSItemProviderReading
	NSItemProviderWriting
	NSMutableCopying
	NSSecureCoding

	// Topic: Exporting the string as data

	// Returns a data object that contains a text stream corresponding to the characters and attributes within the specified range.
	DataFromRangeDocumentAttributesError(range_ NSRange, dict INSDictionary) (INSData, error)
	// Returns a file wrapper object that contains a text stream corresponding to the characters and attributes within the specified range.
	FileWrapperFromRangeDocumentAttributesError(range_ NSRange, dict INSDictionary) (INSFileWrapper, error)
	// Returns a data object that contains a Microsoft Word–format stream corresponding to the characters and attributes within the specified range.
	DocFormatFromRangeDocumentAttributes(range_ NSRange, dict INSDictionary) INSData
	// Returns a data object that contains an RTF stream corresponding to the characters and attributes within the specified range, omitting all attachment attributes.
	RTFFromRangeDocumentAttributes(range_ NSRange, dict INSDictionary) INSData
	// Returns a data object that contains an RTFD stream corresponding to the characters and attributes within the specified range.
	RTFDFromRangeDocumentAttributes(range_ NSRange, dict INSDictionary) INSData
	// Returns a file wrapper object that contains an RTFD document corresponding to the characters and attributes within the specified range.
	RTFDFileWrapperFromRangeDocumentAttributes(range_ NSRange, dict INSDictionary) INSFileWrapper

	// Topic: Getting the characters

	// The character contents of the attributed string as a string.
	String() string
	// The length of the attributed string.
	Length() uint
	// Returns an attributed string consisting of the characters and attributes within the specified range in the attributed string.
	AttributedSubstringFromRange(range_ NSRange) INSAttributedString

	// Topic: Getting font attribute information

	// Returns the font attributes in effect for the character at the specified location.
	FontAttributesInRange(range_ NSRange) INSDictionary
	// Returns the ruler (paragraph) attributes in effect for the characters within the specified range.
	RulerAttributesInRange(range_ NSRange) INSDictionary

	// Topic: Getting attributes for a range of text

	// Returns the attributes for the character at the specified index.
	AttributesAtIndexEffectiveRange(location uint, range_ NSRangePointer) INSDictionary
	// Returns the attributes for the character at the specified index and, by reference, the range where the attributes apply.
	AttributesAtIndexLongestEffectiveRangeInRange(location uint, range_ NSRangePointer, rangeLimit NSRange) INSDictionary
	// Returns the value for an attribute with the specified name of the character at the specified index and, by reference, the range where the attribute applies.
	AttributeAtIndexEffectiveRange(attrName NSAttributedStringKey, location uint, range_ NSRangePointer) objectivec.IObject
	// Returns the value for the attribute with the specified name of the character at the specified index and, by reference, the range where the attribute applies.
	AttributeAtIndexLongestEffectiveRangeInRange(attrName NSAttributedStringKey, location uint, range_ NSRangePointer, rangeLimit NSRange) objectivec.IObject
	// Executes the specified closure or block for each range of a particular attribute in the attributed string.
	EnumerateAttributeInRangeOptionsUsingBlock(attrName NSAttributedStringKey, enumerationRange NSRange, opts NSAttributedStringEnumerationOptions, block ObjectHandler)

	// Topic: Comparing strings

	// Returns a Boolean value that indicates whether the attributed string is equal to the specified string.
	IsEqualToAttributedString(other INSAttributedString) bool

	// Topic: Getting the supported text-file formats

	// Returns a Boolean value that indicates whether the specified range of text prefers RTFD formatting.
	PrefersRTFDInRange(range_ NSRange) bool

	// Topic: Calculating linguistic units

	// Returns the range of characters that form a word (or other linguistic unit) surrounding the specified index, taking language characteristics into account.
	DoubleClickAtIndex(location uint) NSRange
	// Returns the appropriate line break when the character at the index doesn’t fit on the same line as the character at the beginning of the range.
	LineBreakBeforeIndexWithinRange(location uint, aRange NSRange) uint
	// Returns the index of the closest character before the specified index, and within the specified range, that can fit on a new line by hyphenating.
	LineBreakByHyphenatingBeforeIndexWithinRange(location uint, aRange NSRange) uint
	// Returns the index of the first character of the word after or before the specified index.
	NextWordFromIndexForward(location uint, isForward bool) uint

	// Topic: Performing automatic grammar agreement

	// If the string has portions tagged with NSInflectionRuleAttributeName that have no format specifiers, create a new string with those portions inflected by following the rule in the attribute.
	AttributedStringByInflectingString() INSAttributedString

	// Topic: Calculating ranges for common elements

	// Returns the index of the item at the specified location within the list.
	ItemNumberInTextListAtIndex(list objectivec.IObject, location uint) int
	// Returns the range of the individual text block that contains the specified location.
	RangeOfTextBlockAtIndex(block objectivec.IObject, location uint) NSRange
	// Returns the range of the specified text list that contains the specified location.
	RangeOfTextListAtIndex(list objectivec.IObject, location uint) NSRange
	// Returns the range of the specified text table that contains the specified location.
	RangeOfTextTableAtIndex(table objectivec.IObject, location uint) NSRange

	// Topic: Drawing the attributed string

	// Draws the attributed string starting at the specified point in the current graphics context.
	DrawAtPoint(point corefoundation.CGPoint)
	// Draws the attributed string inside the specified bounding rectangle in the current graphics context.
	DrawInRect(rect corefoundation.CGRect)
	// Draws the attributed string in the specified bounding rectangle using the provided options.
	DrawWithRectOptionsContext(rect corefoundation.CGRect, options NSStringDrawingOptions, context objectivec.IObject)

	// Topic: Getting metrics for the string

	// Returns the size necessary to draw the string.
	Size() corefoundation.CGSize
	// Returns the bounding rectangle necessary to draw the string.
	BoundingRectWithSizeOptionsContext(size corefoundation.CGSize, options NSStringDrawingOptions, context objectivec.IObject) corefoundation.CGRect
	// Returns a Boolean value that indicates if the attributed string contains an attachment in the specified range.
	ContainsAttachmentsInRange(range_ NSRange) bool

	// The string encoding for the document.
	CharacterEncoding() INSString
	// A Boolean value that indicates whether the attribute string contains any attachment attributes.
	ContainsAttachments() bool
	// The HTML elements to exclude in generated HTML.
	ExcludedElements() INSString
	// The link for the text.
	Link() NSAttributedStringKey
	// The number of spaces for indenting nested HTML elements.
	PrefixSpaces() INSString
	// The name of the text encoding to use.
	TextEncodingName() INSString
	// Calculates and returns a bounding rectangle for the attributed string using the options specified within the specified rectangle in the current graphics context.
	BoundingRectWithSizeOptions(size corefoundation.CGSize, options NSStringDrawingOptions) NSRect
	// Draws the attributed string with the specified options within the specified rectangle in the current graphics context.
	DrawWithRectOptions(rect corefoundation.CGRect, options NSStringDrawingOptions)
	// Creates a new attributed string from the contents of another attributed string.
	InitWithAttributedString(attrStr INSAttributedString) NSAttributedString
	// Creates an attributed string from the contents of a specified URL that contains Markdown-formatted data using the provided options.
	InitWithContentsOfMarkdownFileAtURLOptionsBaseURLError(markdownFile INSURL, options INSAttributedStringMarkdownParsingOptions, baseURL INSURL) (NSAttributedString, error)
	// Creates an attributed string from the contents of the specified data object.
	InitWithDataOptionsDocumentAttributesError(data INSData, options INSDictionary, dict INSDictionary) (NSAttributedString, error)
	// Creates an attributed string from Microsoft Word format data in the specified data object.
	InitWithDocFormatDocumentAttributes(data INSData, dict INSDictionary) NSAttributedString
	// Initializes an attributed string by substituting arguments into a specially formatted string.
	InitWithFormatOptionsLocale(format INSAttributedString, options NSAttributedStringFormattingOptions, locale INSLocale) NSAttributedString
	// Initializes an attributed string by substituting a list of function arguments into a specially formatted string.
	InitWithFormatOptionsLocaleArguments(format INSAttributedString, options NSAttributedStringFormattingOptions, locale INSLocale, arguments unsafe.Pointer) NSAttributedString
	// Initializes an attributed string by substituting arguments into a specially formatted string and applying additional contextual information.
	InitWithFormatOptionsLocaleContext(format INSAttributedString, options NSAttributedStringFormattingOptions, locale INSLocale, context INSDictionary) NSAttributedString
	// Initializes an attributed string by substituting a list of function arguments into a specially formatted string and applying additional contextual information.
	InitWithFormatOptionsLocaleContextArguments(format INSAttributedString, options NSAttributedStringFormattingOptions, locale INSLocale, context INSDictionary, arguments unsafe.Pointer) NSAttributedString
	// Creates an attributed string from the HTML in the specified data object and base URL.
	InitWithHTMLBaseURLDocumentAttributes(data INSData, base INSURL, dict INSDictionary) NSAttributedString
	// Creates an attributed string from the HTML in the specified data object.
	InitWithHTMLDocumentAttributes(data INSData, dict INSDictionary) NSAttributedString
	// Creates an attributed string from the HTML in the specified data object.
	InitWithHTMLOptionsDocumentAttributes(data INSData, options INSDictionary, dict INSDictionary) NSAttributedString
	// Creates an attributed string from Markdown-formatted data using the provided options.
	InitWithMarkdownOptionsBaseURLError(markdown INSData, options INSAttributedStringMarkdownParsingOptions, baseURL INSURL) (NSAttributedString, error)
	// Creates an attributed string from a Markdown-formatted string using the provided options.
	InitWithMarkdownStringOptionsBaseURLError(markdownString string, options INSAttributedStringMarkdownParsingOptions, baseURL INSURL) (NSAttributedString, error)
	// Creates an attributed string by decoding the stream of RTFD commands and data in the specified data object.
	InitWithRTFDDocumentAttributes(data INSData, dict INSDictionary) NSAttributedString
	// Creates an attributed string from the specified file wrapper that contains an RTFD document.
	InitWithRTFDFileWrapperDocumentAttributes(wrapper INSFileWrapper, dict INSDictionary) NSAttributedString
	// Creates an attributed string by decoding the stream of RTF commands and data in the specified data object.
	InitWithRTFDocumentAttributes(data INSData, dict INSDictionary) NSAttributedString
	// Creates an attributed string with the specified text and no attribute information.
	InitWithString(str string) NSAttributedString
	// Creates an attributed string with the specified text and attributes.
	InitWithStringAttributes(str string, attrs INSDictionary) NSAttributedString
	// Creates an attributed string from the contents of the specified URL.
	InitWithURLOptionsDocumentAttributesError(url INSURL, options INSDictionary, dict INSDictionary) (NSAttributedString, error)
}

// Init initializes the instance.
func (a NSAttributedString) Init() NSAttributedString {
	rv := objc.Send[NSAttributedString](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAttributedString) Autorelease() NSAttributedString {
	rv := objc.Send[NSAttributedString](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAttributedString creates a new NSAttributedString instance.
func NewNSAttributedString() NSAttributedString {
	class := getNSAttributedStringClass()
	rv := objc.Send[NSAttributedString](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an attributed string with an adaptive image glyph and applies the
// specified attributes to it.
//
// adaptiveImageGlyph: The adaptive image glyph to place in the string. Typically, you get this
// type from the text input system.
//
// attributes: The attributes to apply to the adaptive image glyph. Specify an empty
// dictionary to create the string without any extra attributes.
//
// # Return Value
// 
// An attributed string containing the adaptive image glyph.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(adaptiveImageGlyph:attributes:)
// adaptiveImageGlyph is a [appkit.NSAdaptiveImageGlyph].
func NewAttributedStringWithAdaptiveImageGlyphAttributes(adaptiveImageGlyph objectivec.IObject, attributes INSDictionary) NSAttributedString {
	rv := objc.Send[objc.ID](objc.ID(getNSAttributedStringClass().class), objc.Sel("attributedStringWithAdaptiveImageGlyph:attributes:"), adaptiveImageGlyph, attributes)
	return NSAttributedStringFromID(rv)
}

// Creates an attributed string with an attachment.
//
// attachment: The attrachment to place in the string.
//
// # Return Value
// 
// An attributed string containing the attachment.
//
// # Discussion
// 
// This is a convenience method for creating an attributed string containing
// an attachment using [character] as the base character.
//
// [character]: https://developer.apple.com/documentation/UIKit/NSTextAttachment/character
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(attachment:)
// attachment is a [appkit.NSTextAttachment].
func NewAttributedStringWithAttachment(attachment objectivec.IObject) NSAttributedString {
	rv := objc.Send[objc.ID](objc.ID(getNSAttributedStringClass().class), objc.Sel("attributedStringWithAttachment:"), attachment)
	return NSAttributedStringFromID(rv)
}

// Creates an attributed string with an attachment and applies the specified
// attributes to it.
//
// attachment: The attrachment to place in the string.
//
// attributes: The attributes to apply to the attachment. Specify an empty dictionary to
// create the string without any extra attributes.
//
// # Return Value
// 
// An attributed string containing the attachment.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(attachment:attributes:)
// attachment is a [appkit.NSTextAttachment].
func NewAttributedStringWithAttachmentAttributes(attachment objectivec.IObject, attributes INSDictionary) NSAttributedString {
	rv := objc.Send[objc.ID](objc.ID(getNSAttributedStringClass().class), objc.Sel("attributedStringWithAttachment:attributes:"), attachment, attributes)
	return NSAttributedStringFromID(rv)
}

// Creates a new attributed string from the contents of another attributed
// string.
//
// attrStr: An attributed string.
//
// # Return Value
// 
// An [NSAttributedString] object initialized with the characters and
// attributes of `attrStr`.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(attributedString:)
func NewAttributedStringWithAttributedString(attrStr INSAttributedString) NSAttributedString {
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAttributedString:"), attrStr)
	return NSAttributedStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewAttributedStringWithCoder(coder INSCoder) NSAttributedString {
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSAttributedStringFromID(rv)
}

// Creates an attributed string from the contents of a specified URL that
// contains Markdown-formatted data using the provided options.
//
// markdownFile: The URL to load Markdown-formatted data from.
//
// options: Options that affect how the initializer interprets formatting in the
// Markdown string. This parameter defaults to no options.
//
// baseURL: The base URL to use when resolving Markdown URLs. The initializer treats
// URLs as being relative to this URL. If this value is `nil`, the initializer
// doesn’t resolve URLs. The default is `nil`.
//
// error: On input, a pointer to an error object. On return, if an error occurs, this
// pointer contains an actual error object with the error information. You may
// specify `nil` for this parameter if you don’t want the error information.
//
// # Return Value
// 
// An attributed string with the parsed Markdown text and styling, or `nil` if
// parsing the data fails.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithContentsOfMarkdownFileAtURL:options:baseURL:error:
func NewAttributedStringWithContentsOfMarkdownFileAtURLOptionsBaseURLError(markdownFile INSURL, options INSAttributedStringMarkdownParsingOptions, baseURL INSURL) (NSAttributedString, error) {
	var errorPtr objc.ID
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfMarkdownFileAtURL:options:baseURL:error:"), markdownFile, options, baseURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSAttributedStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSAttributedStringFromID(rv), nil
}

// Creates an attributed string from the contents of the specified data
// object.
//
// data: The data from which to create the string.
//
// options: Attributes for interpreting the document contents. Specify the
// [documentType] or [fileType] option to interpret the data as a specific
// type. When sharing files between different platforms, specify the
// [sourceTextScaling] or [targetTextScaling] options for any required text
// scaling behaviors. Specify the [characterEncoding] attribute for plain-text
// files. Specify the [defaultAttributes] key to apply document attributes to
// the returned string. If you specify an empty dictionary, the method
// identifies the data format from the data itself.
// //
// [characterEncoding]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/characterEncoding
// [defaultAttributes]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/defaultAttributes
// [documentType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/documentType
// [fileType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentReadingOptionKey/fileType
// [sourceTextScaling]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentReadingOptionKey/sourceTextScaling
// [targetTextScaling]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentReadingOptionKey/targetTextScaling
//
// dict: An in-out dictionary containing document-level attributes. On output, this
// method updates the dictionary to contain any document-specific keys found
// in the data. Specify `nil` if you don’t want the document attributes.
//
// # Return Value
// 
// Returns an initialized attributed string object, or `nil` if the method
// can’t decode the data.
//
// # Discussion
// 
// Don’t call this method from a background thread if the `options`
// dictionary includes the [documentType] attribute with a value of [html]. If
// you do, the method tries to synchronize with the main thread, fails, and
// times out. Calling it from the main thread works, but can still time out if
// the HTML contains references to external resources. The HTML import
// mechanism is meant for implementing something like markdown (that is, text
// styles, colors, and so on), not for general HTML import.
//
// [documentType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/documentType
// [html]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentType/html
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(data:options:documentAttributes:)
func NewAttributedStringWithDataOptionsDocumentAttributesError(data INSData, options INSDictionary, dict INSDictionary) (NSAttributedString, error) {
	var errorPtr objc.ID
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:options:documentAttributes:error:"), data, options, dict, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSAttributedStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSAttributedStringFromID(rv), nil
}

// Creates an attributed string from Microsoft Word format data in the
// specified data object.
//
// data: The data from which to create the string.
//
// dict: An in-out dictionary containing document-level attributes. On output, this
// method updates the dictionary to contain any document-specific keys found
// in the data. Specify `nil` if you don’t want the document attributes.
//
// # Return Value
// 
// Returns an initialized attributed string object, or `nil` if the method
// can’t decode the data.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(docFormat:documentAttributes:)
func NewAttributedStringWithDocFormatDocumentAttributes(data INSData, dict INSDictionary) NSAttributedString {
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDocFormat:documentAttributes:"), data, dict)
	return NSAttributedStringFromID(rv)
}

// Initializes a new attributed string object from the data at the specified
// URL.
//
// url: An [NSURL] object specifying the document to load.
//
// options: Document attributes for interpreting the document contents. [documentType],
// [characterEncoding], and [defaultAttributes] are supported option keys. If
// not specified, the method examines the data to attempt to determine the
// appropriate attributes.
// //
// [characterEncoding]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/characterEncoding
// [defaultAttributes]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/defaultAttributes
// [documentType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/documentType
//
// dict: If non-[NULL], returns a dictionary with various document-wide attributes
// accessible via document attribute keys.
//
// # Return Value
// 
// Returns an initialized attributed string object, or `nil` if the data
// can’t be decoded.
//
// # Discussion
// 
// The HTML importer should not be called from a background thread (that is,
// the `options` dictionary includes [documentType] with a value of [html]).
// It will try to synchronize with the main thread, fail, and time out.
// Calling it from the main thread works (but can still time out if the HTML
// contains references to external resources, which should be avoided at all
// costs). The HTML import mechanism is meant for implementing something like
// markdown (that is, text styles, colors, and so on), not for general HTML
// import.
//
// [documentType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/documentType
// [html]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentType/html
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(fileURL:options:documentAttributes:)
func NewAttributedStringWithFileURLOptionsDocumentAttributesError(url INSURL, options INSDictionary, dict INSDictionary) (NSAttributedString, error) {
	var errorPtr objc.ID
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFileURL:options:documentAttributes:error:"), url, options, dict, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSAttributedStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSAttributedStringFromID(rv), nil
}

// Initializes an attributed string by substituting arguments into a specially
// formatted string.
//
// format: The format string to use to create the final string. For a list of format
// specifiers you can include in this string, see [String Format Specifiers].
// //
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
//
// options: Options for how to apply attributes to the string’s content.
//
// locale: The locale to use for formatting the string. The locale controls the
// formatting of region-sensitive values such as numbers and currencies.
//
// # Return Value
// 
// An initialized attributed string that combines the format string with the
// provided arguments and other information.
//
// # Discussion
// 
// Pass an optional list of trailing variadic arguments to substitute into the
// `format` string.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithFormat:options:locale:
func NewAttributedStringWithFormatOptionsLocale(format INSAttributedString, options NSAttributedStringFormattingOptions, locale INSLocale) NSAttributedString {
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:options:locale:"), format, options, locale)
	return NSAttributedStringFromID(rv)
}

// Initializes an attributed string by substituting a list of function
// arguments into a specially formatted string.
//
// format: The format string to use to create the final string. For a list of format
// specifiers you can include in this string, see [String Format Specifiers].
// //
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
//
// options: Options for how to apply attributes to the string’s content.
//
// locale: The locale to use for formatting the string. The locale controls the
// formatting of region-sensitive values such as numbers and currencies.
//
// arguments: A list of arguments to substitute into the `format` string.
//
// # Return Value
// 
// An initialized attributed string that combines the format string with the
// provided arguments and other information.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithFormat:options:locale:arguments:
func NewAttributedStringWithFormatOptionsLocaleArguments(format INSAttributedString, options NSAttributedStringFormattingOptions, locale INSLocale, arguments unsafe.Pointer) NSAttributedString {
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:options:locale:arguments:"), format, options, locale, arguments)
	return NSAttributedStringFromID(rv)
}

// Initializes an attributed string by substituting arguments into a specially
// formatted string and applying additional contextual information.
//
// format: The format string to use to create the final string. For a list of format
// specifiers you can include in this string, see [String Format Specifiers].
// //
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
//
// options: Options for how to apply attributes to the string’s content.
//
// locale: The locale to use for formatting the string. The locale controls the
// formatting of region-sensitive values such as numbers and currencies.
//
// context: Additional options to apply to the string.
//
// # Return Value
// 
// An initialized attributed string that combines the format string with the
// provided arguments and other information.
//
// # Discussion
// 
// Pass an optional list of trailing variadic arguments to substitute into the
// `format` string.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithFormat:options:locale:context:
func NewAttributedStringWithFormatOptionsLocaleContext(format INSAttributedString, options NSAttributedStringFormattingOptions, locale INSLocale, context INSDictionary) NSAttributedString {
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:options:locale:context:"), format, options, locale, context)
	return NSAttributedStringFromID(rv)
}

// Initializes an attributed string by substituting a list of function
// arguments into a specially formatted string and applying additional
// contextual information.
//
// format: The format string to use to create the final string. For a list of format
// specifiers you can include in this string, see [String Format Specifiers].
// //
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
//
// options: Options for how to apply attributes to the string’s content.
//
// locale: The locale to use for formatting the string. The locale controls the
// formatting of region-sensitive values such as numbers and currencies.
//
// context: Additional options to apply to the string.
//
// arguments: A list of arguments to substitute into the `format` string.
//
// # Return Value
// 
// An initialized attributed string that combines the format string with the
// provided arguments and other information.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithFormat:options:locale:context:arguments:
func NewAttributedStringWithFormatOptionsLocaleContextArguments(format INSAttributedString, options NSAttributedStringFormattingOptions, locale INSLocale, context INSDictionary, arguments unsafe.Pointer) NSAttributedString {
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:options:locale:context:arguments:"), format, options, locale, context, arguments)
	return NSAttributedStringFromID(rv)
}

// Creates an attributed string from the HTML in the specified data object and
// base URL.
//
// data: A data object with text in HTML format. The method uses this data to create
// the attributed string.
//
// base: An [NSURL] that represents the base URL for all links within the HTML.
//
// dict: An in-out dictionary containing document-level attributes. On output, this
// method updates the dictionary to contain any document-specific keys found
// in the data. Specify `nil` if you don’t want the document attributes.
//
// # Return Value
// 
// Returns an initialized object, or `nil` if the data can’t be decoded.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(HTML:baseURL:documentAttributes:)
func NewAttributedStringWithHTMLBaseURLDocumentAttributes(data INSData, base INSURL, dict INSDictionary) NSAttributedString {
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithHTML:baseURL:documentAttributes:"), data, base, dict)
	return NSAttributedStringFromID(rv)
}

// Creates an attributed string from the HTML in the specified data object.
//
// data: A data object with text in HTML format. The method uses this data to create
// the attributed string.
//
// dict: An in-out dictionary containing document-level attributes. On output, this
// method updates the dictionary to contain any document-specific keys found
// in the data. Specify `nil` if you don’t want the document attributes
//
// # Return Value
// 
// Returns an initialized object, or `nil` if the data can’t be decoded.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(HTML:documentAttributes:)
func NewAttributedStringWithHTMLDocumentAttributes(data INSData, dict INSDictionary) NSAttributedString {
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithHTML:documentAttributes:"), data, dict)
	return NSAttributedStringFromID(rv)
}

// Creates an attributed string from the HTML in the specified data object.
//
// data: A data object with text in HTML format. The method uses this data to create
// the attributed string.
//
// options: Specifies additional options for loading the document. For a list of
// possible keys, see [NSAttributedStringDocumentReadingOptionKey].
// //
// [NSAttributedStringDocumentReadingOptionKey]: https://developer.apple.com/documentation/UIKit/NSAttributedStringDocumentReadingOptionKey
//
// dict: An in-out dictionary containing document-level attributes. On output, this
// method updates the dictionary to contain any document-specific keys found
// in the data. Specify `nil` if you don’t want the document attributes.
//
// # Return Value
// 
// Returns an initialized object, or `nil` if the data can’t be decoded.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(HTML:options:documentAttributes:)
func NewAttributedStringWithHTMLOptionsDocumentAttributes(data INSData, options INSDictionary, dict INSDictionary) NSAttributedString {
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithHTML:options:documentAttributes:"), data, options, dict)
	return NSAttributedStringFromID(rv)
}

// Creates an attributed string from Markdown-formatted data using the
// provided options.
//
// markdown: The [NSData] instance that contains the Markdown formatting.
//
// options: Options that affect how the initializer interprets formatting in the
// Markdown string. This parameter defaults to no options.
//
// baseURL: The base URL to use when resolving Markdown URLs. The initializer treats
// URLs as being relative to this URL. If this value is `nil`, the initializer
// doesn’t resolve URLs. The default is `nil`.
//
// error: On input, a pointer to an error object. On return, if an error occurs, this
// pointer contains an actual error object with the error information. You may
// specify `nil` for this parameter if you don’t want the error information.
//
// # Return Value
// 
// An attributed string with the parsed Markdown text and styling, or `nil` if
// parsing the data fails.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithMarkdown:options:baseURL:error:
func NewAttributedStringWithMarkdownOptionsBaseURLError(markdown INSData, options INSAttributedStringMarkdownParsingOptions, baseURL INSURL) (NSAttributedString, error) {
	var errorPtr objc.ID
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMarkdown:options:baseURL:error:"), markdown, options, baseURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSAttributedStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSAttributedStringFromID(rv), nil
}

// Creates an attributed string from a Markdown-formatted string using the
// provided options.
//
// markdownString: The string that contains the Markdown formatting.
//
// options: Options that affect how the initializer interprets formatting in the
// Markdown string. This parameter defaults to no options.
//
// baseURL: The base URL to use when resolving Markdown URLs. The initializer treats
// URLs as being relative to this URL. If this value is `nil`, the initializer
// doesn’t resolve URLs. The default is `nil`.
//
// error: On input, a pointer to an error object. On return, if an error occurs, this
// pointer contains an actual error object with the error information. You may
// specify `nil` for this parameter if you don’t want the error information.
//
// # Return Value
// 
// An attributed string with the parsed Markdown text and styling, or `nil` if
// parsing the data fails.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithMarkdownString:options:baseURL:error:
func NewAttributedStringWithMarkdownStringOptionsBaseURLError(markdownString string, options INSAttributedStringMarkdownParsingOptions, baseURL INSURL) (NSAttributedString, error) {
	var errorPtr objc.ID
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMarkdownString:options:baseURL:error:"), objc.String(markdownString), options, baseURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSAttributedStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSAttributedStringFromID(rv), nil
}

// Creates an attributed string by decoding the stream of RTFD commands and
// data in the specified data object.
//
// data: The data containing the RTFD content.
//
// dict: An in-out dictionary containing document-level attributes. On output, this
// method updates the dictionary to contain any document-specific keys found
// in the data. Specify `nil` if you don’t want the document attributes.
//
// # Return Value
// 
// Returns an initialized attributed string object, or `nil` if the method
// can’t decode the data.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(RTFD:documentAttributes:)
func NewAttributedStringWithRTFDDocumentAttributes(data INSData, dict INSDictionary) NSAttributedString {
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRTFD:documentAttributes:"), data, dict)
	return NSAttributedStringFromID(rv)
}

// Creates an attributed string from the specified file wrapper that contains
// an RTFD document.
//
// wrapper: The [NSFileWrapper] containing the RTFD document.
//
// dict: An in-out dictionary containing document-level attributes. On output, this
// method updates the dictionary to contain any document-specific keys found
// in the data. Specify `nil` if you don’t want the document attributes.
//
// # Return Value
// 
// Returns an initialized attributed string object, or `nil` if the method
// can’t decode the data.
//
// # Discussion
// 
// Also returns by reference in `dict` a dictionary containing document-level
// attributes described in [NSAttributedString.DocumentAttributeKey]. `dict`
// may be [NULL], in which case no document attributes are returned. Returns
// an initialized object, or `nil` if `wrapper` can’t be interpreted as an
// RTFD document.
//
// [NSAttributedString.DocumentAttributeKey]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(RTFDFileWrapper:documentAttributes:)
func NewAttributedStringWithRTFDFileWrapperDocumentAttributes(wrapper INSFileWrapper, dict INSDictionary) NSAttributedString {
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRTFDFileWrapper:documentAttributes:"), wrapper, dict)
	return NSAttributedStringFromID(rv)
}

// Creates an attributed string by decoding the stream of RTF commands and
// data in the specified data object.
//
// data: The data containing RTF content.
//
// dict: An in-out dictionary containing document-level attributes. On output, this
// method updates the dictionary to contain any document-specific keys found
// in the data. Specify `nil` if you don’t want the document attributes.
//
// # Return Value
// 
// Returns an initialized attributed string object, or `nil` if the method
// can’t decode the data.
//
// # Discussion
// 
// Also returns by reference in `dict` a dictionary containing document-level
// attributes described in [NSAttributedString.DocumentAttributeKey]. `dict`
// may be [NULL], in which case no document attributes are returned. Returns
// an initialized object, or `nil` if `data` can’t be decoded.
//
// [NSAttributedString.DocumentAttributeKey]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(RTF:documentAttributes:)
func NewAttributedStringWithRTFDocumentAttributes(data INSData, dict INSDictionary) NSAttributedString {
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRTF:documentAttributes:"), data, dict)
	return NSAttributedStringFromID(rv)
}

// Creates an attributed string with the specified text and no attribute
// information.
//
// str: The text for the new attributed string.
//
// # Return Value
// 
// An [NSAttributedString] object initialized with the characters of `str` and
// no attribute information.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(string:)
func NewAttributedStringWithString(str string) NSAttributedString {
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:"), objc.String(str))
	return NSAttributedStringFromID(rv)
}

// Creates an attributed string with the specified text and attributes.
//
// str: The text for the new attributed string.
//
// attrs: The attributes for the new attributed string. This method applies the
// attributes to the entire string. For a list of attributes that you can
// include in this dictionary, see [NSAttributedStringKey].
//
// # Discussion
// 
// Returns an [NSAttributedString] object initialized with the characters of
// `str` and the attributes of `attrs`.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(string:attributes:)
func NewAttributedStringWithStringAttributes(str string, attrs INSDictionary) NSAttributedString {
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:attributes:"), objc.String(str), attrs)
	return NSAttributedStringFromID(rv)
}

// Creates an attributed string from the contents of the specified URL.
//
// url: An [NSURL] object specifying the document to load.
//
// options: Attributes for interpreting the document contents. Specify the
// [documentType] or [fileType] option to interpret the data as a specific
// type. When sharing files between different platforms, specify the
// [sourceTextScaling] or [targetTextScaling] options for any required text
// scaling behaviors. Specify the [characterEncoding] attribute for plain-text
// files. Specify the [defaultAttributes] key to apply document attributes to
// the returned string. If you specify an empty dictionary, the method
// identifies the data format from the data itself.
// //
// [characterEncoding]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/characterEncoding
// [defaultAttributes]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/defaultAttributes
// [documentType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/documentType
// [fileType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentReadingOptionKey/fileType
// [sourceTextScaling]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentReadingOptionKey/sourceTextScaling
// [targetTextScaling]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentReadingOptionKey/targetTextScaling
//
// dict: An in-out dictionary containing document-level attributes. On output, this
// method updates the dictionary to contain any document-specific keys found
// in the data. Specify `nil` if you don’t want the document attributes.
//
// # Return Value
// 
// Returns an initialized attributed string object, or `nil` if the method
// can’t decode the data.
//
// # Discussion
// 
// Filter services can be used to convert the file into a format recognized by
// Cocoa. The `options` dictionary specifies how the document should be loaded
// and can contain the values described in
// [NSAttributedStringDocumentReadingOptionKey]. If you specify the
// [documentType] or [fileType] attribute, this method treats the data as if
// it is in the specified format. If you don’t specify one of these options,
// the method examines the document and loads it using whatever format it
// seems to contain.
// 
// If an error occurs, the method returns `nil` and sets the `error` parameter
// to an [NSError] object with information about why it couldn’t create the
// attributed string object.
//
// [NSAttributedStringDocumentReadingOptionKey]: https://developer.apple.com/documentation/UIKit/NSAttributedStringDocumentReadingOptionKey
// [documentType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentReadingOptionKey/documentType
// [fileType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentReadingOptionKey/fileType
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(URL:options:documentAttributes:)
func NewAttributedStringWithURLOptionsDocumentAttributesError(url INSURL, options INSDictionary, dict INSDictionary) (NSAttributedString, error) {
	var errorPtr objc.ID
	instance := getNSAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:options:documentAttributes:error:"), url, options, dict, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSAttributedStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSAttributedStringFromID(rv), nil
}

// Returns a data object that contains a text stream corresponding to the
// characters and attributes within the specified range.
//
// range: The range.
//
// dict: A required dictionary specifying the document attributes. The dictionary
// contains values from `Document Types` and must at least contain
// [documentType].
// //
// [documentType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/documentType
//
// # Return Value
// 
// Returns the data for the attributed string, or `nil` if failure. When
// `nil`, `error` encapsulates the error information.
//
// # Discussion
// 
// Raises an [rangeException] if any part of `range` lies beyond the end of
// the receiver’s characters.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/data(from:documentAttributes:)
func (a NSAttributedString) DataFromRangeDocumentAttributesError(range_ NSRange, dict INSDictionary) (INSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("dataFromRange:documentAttributes:error:"), range_, dict, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSData{}, NSErrorFrom(errorPtr)
	}
	return NSDataFromID(rv), nil

}
// Returns a file wrapper object that contains a text stream corresponding to
// the characters and attributes within the specified range.
//
// range: The range.
//
// dict: A required dictionary specifying the document attributes. The dictionary
// contains values from `Document Types` and must at least contain
// [documentType].
// //
// [documentType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/documentType
//
// # Return Value
// 
// Returns a file wrapper for the appropriate document type, or `nil` if
// failure. When `nil`, `error` encapsulates the error information.
//
// # Discussion
// 
// Raises an [rangeException] if any part of `range` lies beyond the end of
// the receiver’s characters.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/fileWrapper(from:documentAttributes:)
func (a NSAttributedString) FileWrapperFromRangeDocumentAttributesError(range_ NSRange, dict INSDictionary) (INSFileWrapper, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("fileWrapperFromRange:documentAttributes:error:"), range_, dict, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return FileWrapper{}, NSErrorFrom(errorPtr)
	}
	return NSFileWrapperFromID(rv), nil

}
// Returns a data object that contains a Microsoft Word–format stream
// corresponding to the characters and attributes within the specified range.
//
// range: The range.
//
// dict: A required dictionary specifying the document attributes. The dictionary
// contains values from `Document Types` and must at least contain
// [documentType].
// //
// [documentType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/documentType
//
// # Return Value
// 
// Returns a data object containing the attributed string as a Microsoft Word
// doc file.
//
// # Discussion
// 
// Raises an [rangeException] if any part of `range` lies beyond the end of
// the receiver’s characters.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/docFormat(from:documentAttributes:)
func (a NSAttributedString) DocFormatFromRangeDocumentAttributes(range_ NSRange, dict INSDictionary) INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("docFormatFromRange:documentAttributes:"), range_, dict)
	return NSDataFromID(rv)
}
// Returns a data object that contains an RTF stream corresponding to the
// characters and attributes within the specified range, omitting all
// attachment attributes.
//
// range: The range.
//
// dict: A required dictionary specifying the document attributes. The dictionary
// contains values from `Document Types` and must at least contain
// [documentType].
// //
// [documentType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/documentType
//
// # Return Value
// 
// A data object containing an RTF stream for the attributed string.
//
// # Discussion
// 
// Writes the document-level attributes in `docAttributes`, as explained in
// [RTF Files and Attributed Strings].
// 
// Raises an [rangeException] if any part of `aRange` lies beyond the end of
// the receiver’s characters.
// 
// When writing data to the pasteboard, you can use the [NSData] object as the
// first argument to the [NSPasteboard] method [setData(_:forType:)], with a
// second argument of [NSRTFPboardType]. Although this method strips
// attachments, it leaves the attachment characters in the text itself. The
// [NSText] method [rtf(from:)], on the other hand, does strip attachment
// characters when extracting RTF.
//
// [NSPasteboard]: https://developer.apple.com/documentation/AppKit/NSPasteboard
// [RTF Files and Attributed Strings]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/TextAttributes/RTFAndAttrStrings.html#//apple_ref/doc/uid/20000164
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
// [rtf(from:)]: https://developer.apple.com/documentation/AppKit/NSText/rtf(from:)
// [setData(_:forType:)]: https://developer.apple.com/documentation/AppKit/NSPasteboard/setData(_:forType:)
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/rtf(from:documentAttributes:)
func (a NSAttributedString) RTFFromRangeDocumentAttributes(range_ NSRange, dict INSDictionary) INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("RTFFromRange:documentAttributes:"), range_, dict)
	return NSDataFromID(rv)
}
// Returns a data object that contains an RTFD stream corresponding to the
// characters and attributes within the specified range.
//
// range: The range.
//
// dict: A required dictionary specifying the document attributes. The dictionary
// contains values from `Document Types` and must at least contain
// [documentType].
// //
// [documentType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/documentType
//
// # Return Value
// 
// A data object containing the RTFD stream containing the characters and
// attributes.
//
// # Discussion
// 
// Writes the document-level attributes in `docAttributes`, as explained in
// [RTF Files and Attributed Strings].
// 
// Raises an [rangeException] if any part of `aRange` lies beyond the end of
// the receiver’s characters.
// 
// When writing data to the pasteboard, you can use the [NSData] object as the
// first argument to the [NSPasteboard] method [setData(_:forType:)], with a
// second argument of [NSRTFPboardType].
//
// [NSPasteboard]: https://developer.apple.com/documentation/AppKit/NSPasteboard
// [RTF Files and Attributed Strings]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/TextAttributes/RTFAndAttrStrings.html#//apple_ref/doc/uid/20000164
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
// [setData(_:forType:)]: https://developer.apple.com/documentation/AppKit/NSPasteboard/setData(_:forType:)
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/rtfd(from:documentAttributes:)
func (a NSAttributedString) RTFDFromRangeDocumentAttributes(range_ NSRange, dict INSDictionary) INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("RTFDFromRange:documentAttributes:"), range_, dict)
	return NSDataFromID(rv)
}
// Returns a file wrapper object that contains an RTFD document corresponding
// to the characters and attributes within the specified range.
//
// range: The range.
//
// dict: A required dictionary specifying the document attributes. The dictionary
// contains values from `Document Types` and must at least contain
// [NSDocumentTypeDocumentAttribute].
//
// # Return Value
// 
// A file wrapper containing the RTFD data.
//
// # Discussion
// 
// The file wrapper also includes the document-level attributes in
// `docAttributes`, as explained in [RTF Files and Attributed Strings].
// 
// Raises an [rangeException] if any part of `aRange` lies beyond the end of
// the receiver’s characters.
// 
// You can save the file wrapper using the
// [write(toFile:atomically:updateFilenames:)] method of [NSFileWrapper].
//
// [RTF Files and Attributed Strings]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/TextAttributes/RTFAndAttrStrings.html#//apple_ref/doc/uid/20000164
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
// [write(toFile:atomically:updateFilenames:)]: https://developer.apple.com/documentation/Foundation/FileWrapper/write(toFile:atomically:updateFilenames:)
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/rtfdFileWrapper(from:documentAttributes:)
func (a NSAttributedString) RTFDFileWrapperFromRangeDocumentAttributes(range_ NSRange, dict INSDictionary) INSFileWrapper {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("RTFDFileWrapperFromRange:documentAttributes:"), range_, dict)
	return NSFileWrapperFromID(rv)
}
// Returns an attributed string consisting of the characters and attributes
// within the specified range in the attributed string.
//
// range: The range from which to create a new attributed string. `aRange` must lie
// within the bounds of the receiver.
//
// # Return Value
// 
// An [NSAttributedString] object consisting of the characters and attributes
// within `aRange` in the receiver.
//
// # Discussion
// 
// Raises an [rangeException] if any part of `aRange` lies beyond the end of
// the receiver’s characters. This method treats the length of the string as
// a valid range value that returns an empty string.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/attributedSubstring(from:)
func (a NSAttributedString) AttributedSubstringFromRange(range_ NSRange) INSAttributedString {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("attributedSubstringFromRange:"), range_)
	return NSAttributedStringFromID(rv)
}
// Returns the font attributes in effect for the character at the specified
// location.
//
// range: The range.
//
// # Return Value
// 
// A dictionary containing the font attributes for the range.
//
// # Discussion
// 
// The dictionary attributes are all those listed in `Character Attributes`,
// except [link], [paragraphStyle], and [attachment].
// 
// Use this method to obtain font attributes that are to be copied or pasted
// with “copy font” operations.
// 
// Raises an [NSRangeException] if any part of `aRange` lies beyond the end of
// the receiver’s characters.
//
// [attachment]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/attachment
// [link]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/link
// [paragraphStyle]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/paragraphStyle
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/fontAttributes(in:)
func (a NSAttributedString) FontAttributesInRange(range_ NSRange) INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("fontAttributesInRange:"), range_)
	return NSDictionaryFromID(rv)
}
// Returns the ruler (paragraph) attributes in effect for the characters
// within the specified range.
//
// range: The range.
//
// # Return Value
// 
// A dictionary containing the ruler attributes in the range.
//
// # Discussion
// 
// The only ruler attribute currently defined is that named by
// [paragraphStyle]. Use this method to obtain attributes that are to be
// copied or pasted with “copy ruler” operations.
// 
// Raises an [rangeException] if any part of `aRange` lies beyond the end of
// the receiver’s characters.
//
// [paragraphStyle]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/paragraphStyle
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/rulerAttributes(in:)
func (a NSAttributedString) RulerAttributesInRange(range_ NSRange) INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("rulerAttributesInRange:"), range_)
	return NSDictionaryFromID(rv)
}
// Returns the attributes for the character at the specified index.
//
// location: The index for which to return attributes. This value must lie within the
// bounds of the receiver.
//
// range: Upon return, the range over which the attributes and values are the same as
// those at `index`. This range isn’t necessarily the maximum range covered,
// and its extent is implementation-dependent. If you need the maximum range,
// use [AttributesAtIndexLongestEffectiveRangeInRange]. If you don’t need
// this value, pass [NULL].
//
// # Return Value
// 
// The attributes for the character at `index`.
//
// # Discussion
// 
// Raises an [rangeException] if `index` lies beyond the end of the
// receiver’s characters.
// 
// For a list of possible attributes, see [NSAttributedStringKey].
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/attributes(at:effectiveRange:)
func (a NSAttributedString) AttributesAtIndexEffectiveRange(location uint, range_ NSRangePointer) INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("attributesAtIndex:effectiveRange:"), location, range_)
	return NSDictionaryFromID(rv)
}
// Returns the attributes for the character at the specified index and, by
// reference, the range where the attributes apply.
//
// location: The index for which to return attributes. This value must not exceed the
// bounds of the receiver.
//
// range: If non-[NULL], upon return contains the maximum range over which the
// attributes and values are the same as those at `index`, clipped to
// `rangeLimit`.
//
// rangeLimit: The range over which to search for continuous presence of the attributes at
// `index`. This value must not exceed the bounds of the receiver.
//
// # Discussion
// 
// Raises an [rangeException] if `index` or any part of `rangeLimit` lies
// beyond the end of the receiver’s characters.
// 
// If you don’t need the range information, it’s far more efficient to use
// the [AttributesAtIndexEffectiveRange] method to retrieve the attribute
// value.
// 
// For a list of possible attributes, see [NSAttributedStringKey].
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/attributes(at:longestEffectiveRange:in:)
func (a NSAttributedString) AttributesAtIndexLongestEffectiveRangeInRange(location uint, range_ NSRangePointer, rangeLimit NSRange) INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("attributesAtIndex:longestEffectiveRange:inRange:"), location, range_, rangeLimit)
	return NSDictionaryFromID(rv)
}
// Returns the value for an attribute with the specified name of the character
// at the specified index and, by reference, the range where the attribute
// applies.
//
// attrName: The name of an attribute.
//
// location: The index for which to return attributes. This value must not exceed the
// bounds of the receiver.
//
// range: If non-[NULL]:
// 
// - If the named attribute exists at `index`, upon return `aRange` contains a
// range over which the named attribute’s value applies. - If the named
// attribute does not exist at `index`, upon return `aRange` contains the
// range over which the attribute does not exist.
// 
// The range isn’t necessarily the maximum range covered by `attributeName`,
// and its extent is implementation-dependent. If you need the maximum range,
// use [AttributeAtIndexLongestEffectiveRangeInRange]. If you don’t need
// this value, pass [NULL].
//
// # Return Value
// 
// The value for the attribute named `attrName` of the character at
// `location`, or `nil` if there is no such attribute.
//
// # Discussion
// 
// For a list of possible attributes, see [NSAttributedStringKey].
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/attribute(_:at:effectiveRange:)
func (a NSAttributedString) AttributeAtIndexEffectiveRange(attrName NSAttributedStringKey, location uint, range_ NSRangePointer) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("attribute:atIndex:effectiveRange:"), objc.String(string(attrName)), location, range_)
	return objectivec.Object{ID: rv}
}
// Returns the value for the attribute with the specified name of the
// character at the specified index and, by reference, the range where the
// attribute applies.
//
// attrName: The name of an attribute.
//
// location: The index at which to test for `attributeName`.
//
// range: If non-[NULL]:
// 
// - If the named attribute exists at `index`, upon return `aRange` contains
// the full range over which the value of the named attribute is the same as
// that at `index`, clipped to `rangeLimit`. - If the named attribute does not
// exist at `index`, upon return `aRange` contains the full range over which
// the attribute does not exist, clipped to `rangeLimit`.
// 
// If you don’t need this value, pass [NULL].
//
// rangeLimit: The range over which to search for continuous presence of `attributeName`.
// This value must not exceed the bounds of the receiver.
//
// # Return Value
// 
// The value for the attribute named `attributeName` of the character at
// `index`, or `nil` if there is no such attribute.
//
// # Discussion
// 
// Raises an [rangeException] if `index` or any part of `rangeLimit` lies
// beyond the end of the receiver’s characters.
// 
// If you don’t need the longest effective range, it’s far more efficient
// to use the [AttributeAtIndexEffectiveRange] method to retrieve the
// attribute value.
// 
// For a list of possible attributes, see [NSAttributedStringKey].
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/attribute(_:at:longestEffectiveRange:in:)
func (a NSAttributedString) AttributeAtIndexLongestEffectiveRangeInRange(attrName NSAttributedStringKey, location uint, range_ NSRangePointer, rangeLimit NSRange) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("attribute:atIndex:longestEffectiveRange:inRange:"), objc.String(string(attrName)), location, range_, rangeLimit)
	return objectivec.Object{ID: rv}
}
// Executes the specified closure or block for each range of a particular
// attribute in the attributed string.
//
// attrName: The name of the attribute to enumerate.
//
// enumerationRange: The range over which the attribute values are enumerated.
//
// opts: The options used by the enumeration. For possible values, see
// [NSAttributedString.EnumerationOptions].
// //
// [NSAttributedString.EnumerationOptions]: https://developer.apple.com/documentation/Foundation/NSAttributedString/EnumerationOptions
//
// block: A closure or block to apply to ranges of the specified attribute in the
// attributed string, taking three arguments:
// 
// - The value for the specified attribute. - The range of the attribute value
// in the attributed string. - A reference to a Boolean value, which you can
// set to [true] within the closure to stop further processing of the
// attributed string.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If this method is called by an instance of [NSMutableAttributedString],
// mutation (deletion, addition, or change) is allowed only if the mutation is
// within the range provided to the block. After a mutation, the enumeration
// continues with the range immediately following the processed range,
// adjusting for any change in length caused by the mutation. For example, if
// `block` is called with a range starting at location [N], and the block
// deletes all the characters in the provided range, the next call will also
// pass [N] as the location of the range.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/enumerateAttribute(_:in:options:using:)
func (a NSAttributedString) EnumerateAttributeInRangeOptionsUsingBlock(attrName NSAttributedStringKey, enumerationRange NSRange, opts NSAttributedStringEnumerationOptions, block ObjectHandler) {
_block3, _cleanup3 := NewObjectBlock(block)
	defer _cleanup3()
	objc.Send[objc.ID](a.ID, objc.Sel("enumerateAttribute:inRange:options:usingBlock:"), objc.String(string(attrName)), enumerationRange, opts, _block3)
}
// Returns a Boolean value that indicates whether the attributed string is
// equal to the specified string.
//
// other: The attributed string with which to compare the receiver.
//
// # Return Value
// 
// [true] if the text and attributes in the current string and `otherString`
// are the same, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method performs a character-by-character comparison of the string and
// its attributes. The character and its attributes must be the same in both
// strings for the method to return [true]. In attributed strings with many
// attributes, such a comparison is unlikely to yield an exact match [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/isEqual(to:)
func (a NSAttributedString) IsEqualToAttributedString(other INSAttributedString) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isEqualToAttributedString:"), other)
	return rv
}
// Returns a Boolean value that indicates whether the specified range of text
// prefers RTFD formatting.
//
// range: The range of text to test.
//
// # Return Value
// 
// [true] if the range of text prefers RTFD formatting, or [false] if you can
// use the RTF format instead.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// When an attributed string contains attachments, you must save it using the
// RTFD file format to preserve the attached files.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/prefersRTFD(in:)
func (a NSAttributedString) PrefersRTFDInRange(range_ NSRange) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("prefersRTFDInRange:"), range_)
	return rv
}
// Returns the range of characters that form a word (or other linguistic unit)
// surrounding the specified index, taking language characteristics into
// account.
//
// location: The index in the attributed string.
//
// # Return Value
// 
// Returns the range of characters that form a word (or other linguistic unit)
// surrounding the given index, taking language characteristics into account.
//
// # Discussion
// 
// Raises an [rangeException] if `index` lies beyond the end of the
// receiver’s characters.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/doubleClick(at:)
func (a NSAttributedString) DoubleClickAtIndex(location uint) NSRange {
	rv := objc.Send[NSRange](a.ID, objc.Sel("doubleClickAtIndex:"), location)
	return NSRange(rv)
}
// Returns the appropriate line break when the character at the index
// doesn’t fit on the same line as the character at the beginning of the
// range.
//
// location: The index in the attributed string.
//
// aRange: The range.
//
// # Return Value
// 
// Returns the index of the closest character before `index` within `aRange`,
// that can be placed on a new line when laying out text. Returns [NSNotFound]
// if no line break is possible before `index`.
//
// # Discussion
// 
// Raises an [rangeException] if `index` or any part of `aRange` lies beyond
// the end of the receiver’s characters.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/lineBreak(before:within:)
func (a NSAttributedString) LineBreakBeforeIndexWithinRange(location uint, aRange NSRange) uint {
	rv := objc.Send[uint](a.ID, objc.Sel("lineBreakBeforeIndex:withinRange:"), location, aRange)
	return rv
}
// Returns the index of the closest character before the specified index, and
// within the specified range, that can fit on a new line by hyphenating.
//
// location: The location in the attributed string.
//
// aRange: The range.
//
// # Return Value
// 
// Returns the index of the closest character before `index` within `aRange`,
// that can be placed on a new line by hyphenating. Returns [NSNotFound] if no
// line break by hyphenation is possible before `index`.
//
// # Discussion
// 
// In other words, during text layout, finds the appropriate line break by
// hyphenation (the character index at which the hyphen glyph should be
// inserted) when the character at `index` won’t fit on the same line as the
// character at the beginning of `aRange`.
// 
// Raises an [rangeException] if `index` or any part of `aRange` lies beyond
// the end of the receiver’s characters.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/lineBreakByHyphenating(before:within:)
func (a NSAttributedString) LineBreakByHyphenatingBeforeIndexWithinRange(location uint, aRange NSRange) uint {
	rv := objc.Send[uint](a.ID, objc.Sel("lineBreakByHyphenatingBeforeIndex:withinRange:"), location, aRange)
	return rv
}
// Returns the index of the first character of the word after or before the
// specified index.
//
// location: The index in the attribute string.
//
// isForward: [true] if the search should be forward, otherwise [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// [true] if this is the first character after `index` that begins a word; if
// `flag` is [false], it’s the first character before `index` that begins a
// word, whether `index` is located within a word or not.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If `index` lies at either end of the string and the search direction would
// progress past that end, it’s returned unchanged.
// 
// This method is intended for moving the insertion point during editing, not
// for linguistic analysis or parsing of text.
// 
// Raises an [rangeException] if `index` lies beyond the end of the
// receiver’s characters.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/nextWord(from:forward:)
func (a NSAttributedString) NextWordFromIndexForward(location uint, isForward bool) uint {
	rv := objc.Send[uint](a.ID, objc.Sel("nextWordFromIndex:forward:"), location, isForward)
	return rv
}
// If the string has portions tagged with NSInflectionRuleAttributeName that
// have no format specifiers, create a new string with those portions
// inflected by following the rule in the attribute.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/inflecting()
func (a NSAttributedString) AttributedStringByInflectingString() INSAttributedString {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("attributedStringByInflectingString"))
	return NSAttributedStringFromID(rv)
}
// Returns the index of the item at the specified location within the list.
//
// list: The text list.
//
// location: The location of the item.
//
// list is a [appkit.NSTextList].
//
// # Return Value
// 
// Returns the index within the list.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/itemNumber(in:at:)
func (a NSAttributedString) ItemNumberInTextListAtIndex(list objectivec.IObject, location uint) int {
	rv := objc.Send[int](a.ID, objc.Sel("itemNumberInTextList:atIndex:"), list, location)
	return rv
}
// Returns the range of the individual text block that contains the specified
// location.
//
// block: The text block.
//
// location: The location in the text block.
//
// block is a [appkit.NSTextBlock].
//
// # Return Value
// 
// The range of the text block containing the location.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/range(of:at:)-1wrcp
func (a NSAttributedString) RangeOfTextBlockAtIndex(block objectivec.IObject, location uint) NSRange {
	rv := objc.Send[NSRange](a.ID, objc.Sel("rangeOfTextBlock:atIndex:"), block, location)
	return NSRange(rv)
}
// Returns the range of the specified text list that contains the specified
// location.
//
// list: The text list.
//
// location: The location in the text list.
//
// list is a [appkit.NSTextList].
//
// # Return Value
// 
// The range of the given text list containing the location.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/range(of:at:)-6um0x
func (a NSAttributedString) RangeOfTextListAtIndex(list objectivec.IObject, location uint) NSRange {
	rv := objc.Send[NSRange](a.ID, objc.Sel("rangeOfTextList:atIndex:"), list, location)
	return NSRange(rv)
}
// Returns the range of the specified text table that contains the specified
// location.
//
// table: The text table.
//
// location: The location.
//
// table is a [appkit.NSTextTable].
//
// # Return Value
// 
// Returns the range of `table` that contains `location`.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/range(of:at:)-3fevu
func (a NSAttributedString) RangeOfTextTableAtIndex(table objectivec.IObject, location uint) NSRange {
	rv := objc.Send[NSRange](a.ID, objc.Sel("rangeOfTextTable:atIndex:"), table, location)
	return NSRange(rv)
}
// Draws the attributed string starting at the specified point in the current
// graphics context.
//
// point: The point in the current graphics context where you want to start drawing
// the string. The coordinate system of the graphics context is usually
// defined by the view in which you are drawing.
//
// # Discussion
// 
// This method draws the entire string starting at the specified point. This
// method draws the line using the attributes specified in the attributed
// string itself. If newline characters are present in the string, those
// characters are honored and cause subsequent text to be placed on the next
// line underneath the starting point.
// 
// There must be either a focused view or an active graphics context when you
// call this method.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/draw(at:)
func (a NSAttributedString) DrawAtPoint(point corefoundation.CGPoint) {
	objc.Send[objc.ID](a.ID, objc.Sel("drawAtPoint:"), point)
}
// Draws the attributed string inside the specified bounding rectangle in the
// current graphics context.
//
// rect: The bounding rectangle in which to draw the string. In AppKit, the origin
// is normally in the lower-left corner of the drawing area, but the origin is
// in the upper-left corner if the focused view is flipped.
//
// # Discussion
// 
// This method draws as much of the string as it can inside the specified
// rectangle, wrapping the string text as needed to make it fit. If the string
// is too long to fit inside the rectangle, the method renders as much as
// possible and clips the rest. It is possible for a portion of a glyph to
// appear outside the area of `rect` if the image bounding box of the
// particular glyph exceeds its typographic bounding box. Text is drawn
// according to its line sweep direction; for example, Arabic text begins at
// the right edge and is potentially clipped on the left.
// 
// Layout always occurs from top to bottom. AppKit automatically adjusts the
// initial drawing point whether or not the view is flipped. For example, if
// the `rect` argument is `{0.0, 0.0, 100.0, 100.0}`, the text origin is {0.0,
// 0.0} when the view coordinates are flipped and {0.0, 100.0} when the view
// is not flipped.
// 
// This method draws the line using the attributes specified in the attributed
// string itself. If newline characters are present in the string, those
// characters are honored and cause subsequent text to be placed on the next
// line underneath the starting point.
// 
// There must be either a focused view or an active graphics context when you
// call this method.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/draw(in:)
func (a NSAttributedString) DrawInRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](a.ID, objc.Sel("drawInRect:"), rect)
}
// Draws the attributed string in the specified bounding rectangle using the
// provided options.
//
// rect: The bounding rectangle in which to draw the string.
//
// options: Additional drawing options to apply to the string during rendering. For a
// list of possible values, see [NSStringDrawingOptions].
// //
// [NSStringDrawingOptions]: https://developer.apple.com/documentation/UIKit/NSStringDrawingOptions
//
// context: A context object with information about how to adjust the font tracking and
// scaling information. On return, the specified object contains information
// about the actual values used to render the string. This parameter may be
// `nil`.
//
// context is a [appkit.NSStringDrawingContext].
//
// # Discussion
// 
// If [usesLineFragmentOrigin] is specified in `options`, it wraps the string
// text as needed to make it fit. If the string is too big to fit completely
// inside the rectangle, the method scales the font or adjusts the letter
// spacing to make the string fit within the given bounds.
// 
// If [usesLineFragmentOrigin] is not specified in `options`, the origin of
// the rectangle is the baseline of the only line. The text will be displayed
// above the rectangle and not inside of it. For example, if you specify a
// rectangle starting at 0,0 and draw the string ‘juxtaposed’, only the
// descenders of the ‘j’ and ‘p’ will be seen. The rest of the text
// will be on the top edge of the rectangle.
// 
// This method draws the line using the attributes specified in the attributed
// string itself. If newline characters are present in the string, those
// characters are honored and cause subsequent text to be placed on the next
// line underneath the starting point.
// 
// There must be either an active graphics context when you call this method.
// 
// # Special Considerations
// 
// This method uses the baseline origin by default, so it renders the string
// as a single line. To render the string in multiple lines, specify
// [usesLineFragmentOrigin] in `options`.
//
// [usesLineFragmentOrigin]: https://developer.apple.com/documentation/UIKit/NSStringDrawingOptions/usesLineFragmentOrigin
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/draw(with:options:context:)
func (a NSAttributedString) DrawWithRectOptionsContext(rect corefoundation.CGRect, options NSStringDrawingOptions, context objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("drawWithRect:options:context:"), rect, options, context)
}
// Returns the size necessary to draw the string.
//
// # Return Value
// 
// The minimum size required to draw the entire contents of the string.
//
// # Discussion
// 
// You can use this method prior to drawing to compute how much space is
// required to draw the string.
// 
// This method may return fractional sizes. When setting the size of your
// view, use the [ceil] function to round fractional values up to the nearest
// whole number.
//
// [ceil]: https://developer.apple.com/documentation/kernel/1557272-ceil
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/size()
func (a NSAttributedString) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](a.ID, objc.Sel("size"))
	return corefoundation.CGSize(rv)
}
// Returns the bounding rectangle necessary to draw the string.
//
// size: The width and height constraints to apply when computing the string’s
// bounding rectangle.
//
// options: Additional drawing options to apply to the string during rendering. For a
// list of possible values, see [NSStringDrawingOptions].
// //
// [NSStringDrawingOptions]: https://developer.apple.com/documentation/UIKit/NSStringDrawingOptions
//
// context: A context object with information about how to adjust the font tracking and
// scaling information. On return, the specified object contains information
// about the actual values used to render the string. This parameter may be
// `nil`.
//
// context is a [appkit.NSStringDrawingContext].
//
// # Return Value
// 
// A rectangle whose size component indicates the width and height required to
// draw the entire contents of the string.
//
// # Discussion
// 
// You can use this method to compute the space required to draw the string.
// The constraints you specify in the size parameter are a guide for the
// renderer for how to size the string. However, the actual bounding rectangle
// returned by this method can be larger than the constraints if additional
// space is needed to render the entire string. Typically, the renderer
// preserves the width constraint and adjusts the height constraint as needed.
// 
// In iOS 7 and later, this method returns fractional sizes (in the `size`
// component of the returned rectangle); to use a returned size to size views,
// you must use raise its value to the nearest higher integer using the [ceil]
// function.
// 
// # Special Considerations
// 
// To calculate the bounding rectangle, this method uses the baseline origin
// by default, so it behaves as a single line. To render the string in
// multiple lines, specify [usesLineFragmentOrigin] in `options`.
//
// [ceil]: https://developer.apple.com/documentation/kernel/1557272-ceil
// [usesLineFragmentOrigin]: https://developer.apple.com/documentation/UIKit/NSStringDrawingOptions/usesLineFragmentOrigin
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/boundingRect(with:options:context:)
func (a NSAttributedString) BoundingRectWithSizeOptionsContext(size corefoundation.CGSize, options NSStringDrawingOptions, context objectivec.IObject) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](a.ID, objc.Sel("boundingRectWithSize:options:context:"), size, options, context)
	return corefoundation.CGRect(rv)
}
// Returns a Boolean value that indicates if the attributed string contains an
// attachment in the specified range.
//
// range: The range.
//
// # Return Value
// 
// [true] if the attributed string contains a property configured as
// [attachment] with [character] in `range`; otherwise, [false].
//
// [attachment]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/attachment
// [character]: https://developer.apple.com/documentation/UIKit/NSTextAttachment/character
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/containsAttachments(in:)
func (a NSAttributedString) ContainsAttachmentsInRange(range_ NSRange) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("containsAttachmentsInRange:"), range_)
	return rv
}
// Calculates and returns a bounding rectangle for the attributed string using
// the options specified within the specified rectangle in the current
// graphics context.
//
// size: The size of the rectangle to draw in.
//
// options: The string drawing options. See [NSStringDrawingOptions] for the possible
// values.
// //
// [NSStringDrawingOptions]: https://developer.apple.com/documentation/UIKit/NSStringDrawingOptions
//
// # Return Value
// 
// The bounding rectangle in the current graphics context.
//
// # Discussion
// 
// The origin of the rectangle returned from this method is the first glyph
// origin.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/boundingRect(with:options:)
func (a NSAttributedString) BoundingRectWithSizeOptions(size corefoundation.CGSize, options NSStringDrawingOptions) NSRect {
	rv := objc.Send[NSRect](a.ID, objc.Sel("boundingRectWithSize:options:"), size, options)
	return NSRect(rv)
}
// Draws the attributed string with the specified options within the specified
// rectangle in the current graphics context.
//
// rect: The rectangle specifies the rendering origin in the current graphics
// context.
//
// options: The string drawing options. See [NSStringDrawingOptions] for the available
// options.
// //
// [NSStringDrawingOptions]: https://developer.apple.com/documentation/UIKit/NSStringDrawingOptions
//
// # Discussion
// 
// The `rect` argument’s origin field specifies the rendering origin. The
// point is interpreted as the baseline origin by default. With
// [NSStringDrawingUsesLineFragmentOrigin], it is interpreted as the upper
// left corner of the line fragment rect. The size field specifies the text
// container size. The width part of the size field specifies the maximum line
// fragment width if larger than `0.0`. The height defines the maximum size
// that can be occupied with text if larger than `0.0` and
// [NSStringDrawingUsesLineFragmentOrigin] is specified. If
// [NSStringDrawingUsesLineFragmentOrigin] is not specified, height is ignored
// and considered to be single-line rendering ([NSLineBreakByWordWrapping] and
// [NSLineBreakByCharWrapping] are treated as [NSLineBreakByClipping]).
// 
// You should only invoke this method when there is a current graphics
// context.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/draw(with:options:)
func (a NSAttributedString) DrawWithRectOptions(rect corefoundation.CGRect, options NSStringDrawingOptions) {
	objc.Send[objc.ID](a.ID, objc.Sel("drawWithRect:options:"), rect, options)
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (a NSAttributedString) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}
// Creates a new attributed string from the contents of another attributed
// string.
//
// attrStr: An attributed string.
//
// # Return Value
// 
// An [NSAttributedString] object initialized with the characters and
// attributes of `attrStr`.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(attributedString:)
func (a NSAttributedString) InitWithAttributedString(attrStr INSAttributedString) NSAttributedString {
	rv := objc.Send[NSAttributedString](a.ID, objc.Sel("initWithAttributedString:"), attrStr)
	return rv
}
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (a NSAttributedString) InitWithCoder(coder INSCoder) NSAttributedString {
	rv := objc.Send[NSAttributedString](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Creates an attributed string from the contents of a specified URL that
// contains Markdown-formatted data using the provided options.
//
// markdownFile: The URL to load Markdown-formatted data from.
//
// options: Options that affect how the initializer interprets formatting in the
// Markdown string. This parameter defaults to no options.
//
// baseURL: The base URL to use when resolving Markdown URLs. The initializer treats
// URLs as being relative to this URL. If this value is `nil`, the initializer
// doesn’t resolve URLs. The default is `nil`.
//
// error: On input, a pointer to an error object. On return, if an error occurs, this
// pointer contains an actual error object with the error information. You may
// specify `nil` for this parameter if you don’t want the error information.
//
// # Return Value
// 
// An attributed string with the parsed Markdown text and styling, or `nil` if
// parsing the data fails.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithContentsOfMarkdownFileAtURL:options:baseURL:error:
func (a NSAttributedString) InitWithContentsOfMarkdownFileAtURLOptionsBaseURLError(markdownFile INSURL, options INSAttributedStringMarkdownParsingOptions, baseURL INSURL) (NSAttributedString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithContentsOfMarkdownFileAtURL:options:baseURL:error:"), markdownFile, options, baseURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSAttributedString{}, NSErrorFrom(errorPtr)
	}
	return NSAttributedStringFromID(rv), nil

}
// Creates an attributed string from the contents of the specified data
// object.
//
// data: The data from which to create the string.
//
// options: Attributes for interpreting the document contents. Specify the
// [documentType] or [fileType] option to interpret the data as a specific
// type. When sharing files between different platforms, specify the
// [sourceTextScaling] or [targetTextScaling] options for any required text
// scaling behaviors. Specify the [characterEncoding] attribute for plain-text
// files. Specify the [defaultAttributes] key to apply document attributes to
// the returned string. If you specify an empty dictionary, the method
// identifies the data format from the data itself.
// //
// [characterEncoding]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/characterEncoding
// [defaultAttributes]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/defaultAttributes
// [documentType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/documentType
// [fileType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentReadingOptionKey/fileType
// [sourceTextScaling]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentReadingOptionKey/sourceTextScaling
// [targetTextScaling]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentReadingOptionKey/targetTextScaling
//
// dict: An in-out dictionary containing document-level attributes. On output, this
// method updates the dictionary to contain any document-specific keys found
// in the data. Specify `nil` if you don’t want the document attributes.
//
// # Return Value
// 
// Returns an initialized attributed string object, or `nil` if the method
// can’t decode the data.
//
// # Discussion
// 
// Don’t call this method from a background thread if the `options`
// dictionary includes the [documentType] attribute with a value of [html]. If
// you do, the method tries to synchronize with the main thread, fails, and
// times out. Calling it from the main thread works, but can still time out if
// the HTML contains references to external resources. The HTML import
// mechanism is meant for implementing something like markdown (that is, text
// styles, colors, and so on), not for general HTML import.
//
// [documentType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/documentType
// [html]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentType/html
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(data:options:documentAttributes:)
func (a NSAttributedString) InitWithDataOptionsDocumentAttributesError(data INSData, options INSDictionary, dict INSDictionary) (NSAttributedString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithData:options:documentAttributes:error:"), data, options, dict, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSAttributedString{}, NSErrorFrom(errorPtr)
	}
	return NSAttributedStringFromID(rv), nil

}
// Creates an attributed string from Microsoft Word format data in the
// specified data object.
//
// data: The data from which to create the string.
//
// dict: An in-out dictionary containing document-level attributes. On output, this
// method updates the dictionary to contain any document-specific keys found
// in the data. Specify `nil` if you don’t want the document attributes.
//
// # Return Value
// 
// Returns an initialized attributed string object, or `nil` if the method
// can’t decode the data.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(docFormat:documentAttributes:)
func (a NSAttributedString) InitWithDocFormatDocumentAttributes(data INSData, dict INSDictionary) NSAttributedString {
	rv := objc.Send[NSAttributedString](a.ID, objc.Sel("initWithDocFormat:documentAttributes:"), data, dict)
	return rv
}
// Initializes an attributed string by substituting arguments into a specially
// formatted string.
//
// format: The format string to use to create the final string. For a list of format
// specifiers you can include in this string, see [String Format Specifiers].
// //
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
//
// options: Options for how to apply attributes to the string’s content.
//
// locale: The locale to use for formatting the string. The locale controls the
// formatting of region-sensitive values such as numbers and currencies.
//
// # Return Value
// 
// An initialized attributed string that combines the format string with the
// provided arguments and other information.
//
// # Discussion
// 
// Pass an optional list of trailing variadic arguments to substitute into the
// `format` string.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithFormat:options:locale:
func (a NSAttributedString) InitWithFormatOptionsLocale(format INSAttributedString, options NSAttributedStringFormattingOptions, locale INSLocale) NSAttributedString {
	rv := objc.Send[NSAttributedString](a.ID, objc.Sel("initWithFormat:options:locale:"), format, options, locale)
	return rv
}
// Initializes an attributed string by substituting a list of function
// arguments into a specially formatted string.
//
// format: The format string to use to create the final string. For a list of format
// specifiers you can include in this string, see [String Format Specifiers].
// //
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
//
// options: Options for how to apply attributes to the string’s content.
//
// locale: The locale to use for formatting the string. The locale controls the
// formatting of region-sensitive values such as numbers and currencies.
//
// arguments: A list of arguments to substitute into the `format` string.
//
// # Return Value
// 
// An initialized attributed string that combines the format string with the
// provided arguments and other information.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithFormat:options:locale:arguments:
func (a NSAttributedString) InitWithFormatOptionsLocaleArguments(format INSAttributedString, options NSAttributedStringFormattingOptions, locale INSLocale, arguments unsafe.Pointer) NSAttributedString {
	rv := objc.Send[NSAttributedString](a.ID, objc.Sel("initWithFormat:options:locale:arguments:"), format, options, locale, arguments)
	return rv
}
// Initializes an attributed string by substituting arguments into a specially
// formatted string and applying additional contextual information.
//
// format: The format string to use to create the final string. For a list of format
// specifiers you can include in this string, see [String Format Specifiers].
// //
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
//
// options: Options for how to apply attributes to the string’s content.
//
// locale: The locale to use for formatting the string. The locale controls the
// formatting of region-sensitive values such as numbers and currencies.
//
// context: Additional options to apply to the string.
//
// # Return Value
// 
// An initialized attributed string that combines the format string with the
// provided arguments and other information.
//
// # Discussion
// 
// Pass an optional list of trailing variadic arguments to substitute into the
// `format` string.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithFormat:options:locale:context:
func (a NSAttributedString) InitWithFormatOptionsLocaleContext(format INSAttributedString, options NSAttributedStringFormattingOptions, locale INSLocale, context INSDictionary) NSAttributedString {
	rv := objc.Send[NSAttributedString](a.ID, objc.Sel("initWithFormat:options:locale:context:"), format, options, locale, context)
	return rv
}
// Initializes an attributed string by substituting a list of function
// arguments into a specially formatted string and applying additional
// contextual information.
//
// format: The format string to use to create the final string. For a list of format
// specifiers you can include in this string, see [String Format Specifiers].
// //
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
//
// options: Options for how to apply attributes to the string’s content.
//
// locale: The locale to use for formatting the string. The locale controls the
// formatting of region-sensitive values such as numbers and currencies.
//
// context: Additional options to apply to the string.
//
// arguments: A list of arguments to substitute into the `format` string.
//
// # Return Value
// 
// An initialized attributed string that combines the format string with the
// provided arguments and other information.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithFormat:options:locale:context:arguments:
func (a NSAttributedString) InitWithFormatOptionsLocaleContextArguments(format INSAttributedString, options NSAttributedStringFormattingOptions, locale INSLocale, context INSDictionary, arguments unsafe.Pointer) NSAttributedString {
	rv := objc.Send[NSAttributedString](a.ID, objc.Sel("initWithFormat:options:locale:context:arguments:"), format, options, locale, context, arguments)
	return rv
}
// Creates an attributed string from the HTML in the specified data object and
// base URL.
//
// data: A data object with text in HTML format. The method uses this data to create
// the attributed string.
//
// base: An [NSURL] that represents the base URL for all links within the HTML.
//
// dict: An in-out dictionary containing document-level attributes. On output, this
// method updates the dictionary to contain any document-specific keys found
// in the data. Specify `nil` if you don’t want the document attributes.
//
// # Return Value
// 
// Returns an initialized object, or `nil` if the data can’t be decoded.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(HTML:baseURL:documentAttributes:)
func (a NSAttributedString) InitWithHTMLBaseURLDocumentAttributes(data INSData, base INSURL, dict INSDictionary) NSAttributedString {
	rv := objc.Send[NSAttributedString](a.ID, objc.Sel("initWithHTML:baseURL:documentAttributes:"), data, base, dict)
	return rv
}
// Creates an attributed string from the HTML in the specified data object.
//
// data: A data object with text in HTML format. The method uses this data to create
// the attributed string.
//
// dict: An in-out dictionary containing document-level attributes. On output, this
// method updates the dictionary to contain any document-specific keys found
// in the data. Specify `nil` if you don’t want the document attributes
//
// # Return Value
// 
// Returns an initialized object, or `nil` if the data can’t be decoded.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(HTML:documentAttributes:)
func (a NSAttributedString) InitWithHTMLDocumentAttributes(data INSData, dict INSDictionary) NSAttributedString {
	rv := objc.Send[NSAttributedString](a.ID, objc.Sel("initWithHTML:documentAttributes:"), data, dict)
	return rv
}
// Creates an attributed string from the HTML in the specified data object.
//
// data: A data object with text in HTML format. The method uses this data to create
// the attributed string.
//
// options: Specifies additional options for loading the document. For a list of
// possible keys, see [NSAttributedStringDocumentReadingOptionKey].
// //
// [NSAttributedStringDocumentReadingOptionKey]: https://developer.apple.com/documentation/UIKit/NSAttributedStringDocumentReadingOptionKey
//
// dict: An in-out dictionary containing document-level attributes. On output, this
// method updates the dictionary to contain any document-specific keys found
// in the data. Specify `nil` if you don’t want the document attributes.
//
// # Return Value
// 
// Returns an initialized object, or `nil` if the data can’t be decoded.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(HTML:options:documentAttributes:)
func (a NSAttributedString) InitWithHTMLOptionsDocumentAttributes(data INSData, options INSDictionary, dict INSDictionary) NSAttributedString {
	rv := objc.Send[NSAttributedString](a.ID, objc.Sel("initWithHTML:options:documentAttributes:"), data, options, dict)
	return rv
}
// Creates an attributed string from Markdown-formatted data using the
// provided options.
//
// markdown: The [NSData] instance that contains the Markdown formatting.
//
// options: Options that affect how the initializer interprets formatting in the
// Markdown string. This parameter defaults to no options.
//
// baseURL: The base URL to use when resolving Markdown URLs. The initializer treats
// URLs as being relative to this URL. If this value is `nil`, the initializer
// doesn’t resolve URLs. The default is `nil`.
//
// error: On input, a pointer to an error object. On return, if an error occurs, this
// pointer contains an actual error object with the error information. You may
// specify `nil` for this parameter if you don’t want the error information.
//
// # Return Value
// 
// An attributed string with the parsed Markdown text and styling, or `nil` if
// parsing the data fails.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithMarkdown:options:baseURL:error:
func (a NSAttributedString) InitWithMarkdownOptionsBaseURLError(markdown INSData, options INSAttributedStringMarkdownParsingOptions, baseURL INSURL) (NSAttributedString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithMarkdown:options:baseURL:error:"), markdown, options, baseURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSAttributedString{}, NSErrorFrom(errorPtr)
	}
	return NSAttributedStringFromID(rv), nil

}
// Creates an attributed string from a Markdown-formatted string using the
// provided options.
//
// markdownString: The string that contains the Markdown formatting.
//
// options: Options that affect how the initializer interprets formatting in the
// Markdown string. This parameter defaults to no options.
//
// baseURL: The base URL to use when resolving Markdown URLs. The initializer treats
// URLs as being relative to this URL. If this value is `nil`, the initializer
// doesn’t resolve URLs. The default is `nil`.
//
// error: On input, a pointer to an error object. On return, if an error occurs, this
// pointer contains an actual error object with the error information. You may
// specify `nil` for this parameter if you don’t want the error information.
//
// # Return Value
// 
// An attributed string with the parsed Markdown text and styling, or `nil` if
// parsing the data fails.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/initWithMarkdownString:options:baseURL:error:
func (a NSAttributedString) InitWithMarkdownStringOptionsBaseURLError(markdownString string, options INSAttributedStringMarkdownParsingOptions, baseURL INSURL) (NSAttributedString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithMarkdownString:options:baseURL:error:"), objc.String(markdownString), options, baseURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSAttributedString{}, NSErrorFrom(errorPtr)
	}
	return NSAttributedStringFromID(rv), nil

}
// Creates an attributed string by decoding the stream of RTFD commands and
// data in the specified data object.
//
// data: The data containing the RTFD content.
//
// dict: An in-out dictionary containing document-level attributes. On output, this
// method updates the dictionary to contain any document-specific keys found
// in the data. Specify `nil` if you don’t want the document attributes.
//
// # Return Value
// 
// Returns an initialized attributed string object, or `nil` if the method
// can’t decode the data.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(RTFD:documentAttributes:)
func (a NSAttributedString) InitWithRTFDDocumentAttributes(data INSData, dict INSDictionary) NSAttributedString {
	rv := objc.Send[NSAttributedString](a.ID, objc.Sel("initWithRTFD:documentAttributes:"), data, dict)
	return rv
}
// Creates an attributed string from the specified file wrapper that contains
// an RTFD document.
//
// wrapper: The [NSFileWrapper] containing the RTFD document.
//
// dict: An in-out dictionary containing document-level attributes. On output, this
// method updates the dictionary to contain any document-specific keys found
// in the data. Specify `nil` if you don’t want the document attributes.
//
// # Return Value
// 
// Returns an initialized attributed string object, or `nil` if the method
// can’t decode the data.
//
// # Discussion
// 
// Also returns by reference in `dict` a dictionary containing document-level
// attributes described in [NSAttributedString.DocumentAttributeKey]. `dict`
// may be [NULL], in which case no document attributes are returned. Returns
// an initialized object, or `nil` if `wrapper` can’t be interpreted as an
// RTFD document.
//
// [NSAttributedString.DocumentAttributeKey]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(RTFDFileWrapper:documentAttributes:)
func (a NSAttributedString) InitWithRTFDFileWrapperDocumentAttributes(wrapper INSFileWrapper, dict INSDictionary) NSAttributedString {
	rv := objc.Send[NSAttributedString](a.ID, objc.Sel("initWithRTFDFileWrapper:documentAttributes:"), wrapper, dict)
	return rv
}
// Creates an attributed string by decoding the stream of RTF commands and
// data in the specified data object.
//
// data: The data containing RTF content.
//
// dict: An in-out dictionary containing document-level attributes. On output, this
// method updates the dictionary to contain any document-specific keys found
// in the data. Specify `nil` if you don’t want the document attributes.
//
// # Return Value
// 
// Returns an initialized attributed string object, or `nil` if the method
// can’t decode the data.
//
// # Discussion
// 
// Also returns by reference in `dict` a dictionary containing document-level
// attributes described in [NSAttributedString.DocumentAttributeKey]. `dict`
// may be [NULL], in which case no document attributes are returned. Returns
// an initialized object, or `nil` if `data` can’t be decoded.
//
// [NSAttributedString.DocumentAttributeKey]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(RTF:documentAttributes:)
func (a NSAttributedString) InitWithRTFDocumentAttributes(data INSData, dict INSDictionary) NSAttributedString {
	rv := objc.Send[NSAttributedString](a.ID, objc.Sel("initWithRTF:documentAttributes:"), data, dict)
	return rv
}
// Creates an attributed string with the specified text and no attribute
// information.
//
// str: The text for the new attributed string.
//
// # Return Value
// 
// An [NSAttributedString] object initialized with the characters of `str` and
// no attribute information.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(string:)
func (a NSAttributedString) InitWithString(str string) NSAttributedString {
	rv := objc.Send[NSAttributedString](a.ID, objc.Sel("initWithString:"), objc.String(str))
	return rv
}
// Creates an attributed string with the specified text and attributes.
//
// str: The text for the new attributed string.
//
// attrs: The attributes for the new attributed string. This method applies the
// attributes to the entire string. For a list of attributes that you can
// include in this dictionary, see [NSAttributedStringKey].
//
// # Discussion
// 
// Returns an [NSAttributedString] object initialized with the characters of
// `str` and the attributes of `attrs`.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(string:attributes:)
func (a NSAttributedString) InitWithStringAttributes(str string, attrs INSDictionary) NSAttributedString {
	rv := objc.Send[NSAttributedString](a.ID, objc.Sel("initWithString:attributes:"), objc.String(str), attrs)
	return rv
}
// Creates an attributed string from the contents of the specified URL.
//
// url: An [NSURL] object specifying the document to load.
//
// options: Attributes for interpreting the document contents. Specify the
// [documentType] or [fileType] option to interpret the data as a specific
// type. When sharing files between different platforms, specify the
// [sourceTextScaling] or [targetTextScaling] options for any required text
// scaling behaviors. Specify the [characterEncoding] attribute for plain-text
// files. Specify the [defaultAttributes] key to apply document attributes to
// the returned string. If you specify an empty dictionary, the method
// identifies the data format from the data itself.
// //
// [characterEncoding]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/characterEncoding
// [defaultAttributes]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/defaultAttributes
// [documentType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentAttributeKey/documentType
// [fileType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentReadingOptionKey/fileType
// [sourceTextScaling]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentReadingOptionKey/sourceTextScaling
// [targetTextScaling]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentReadingOptionKey/targetTextScaling
//
// dict: An in-out dictionary containing document-level attributes. On output, this
// method updates the dictionary to contain any document-specific keys found
// in the data. Specify `nil` if you don’t want the document attributes.
//
// # Return Value
// 
// Returns an initialized attributed string object, or `nil` if the method
// can’t decode the data.
//
// # Discussion
// 
// Filter services can be used to convert the file into a format recognized by
// Cocoa. The `options` dictionary specifies how the document should be loaded
// and can contain the values described in
// [NSAttributedStringDocumentReadingOptionKey]. If you specify the
// [documentType] or [fileType] attribute, this method treats the data as if
// it is in the specified format. If you don’t specify one of these options,
// the method examines the document and loads it using whatever format it
// seems to contain.
// 
// If an error occurs, the method returns `nil` and sets the `error` parameter
// to an [NSError] object with information about why it couldn’t create the
// attributed string object.
//
// [NSAttributedStringDocumentReadingOptionKey]: https://developer.apple.com/documentation/UIKit/NSAttributedStringDocumentReadingOptionKey
// [documentType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentReadingOptionKey/documentType
// [fileType]: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentReadingOptionKey/fileType
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(URL:options:documentAttributes:)
func (a NSAttributedString) InitWithURLOptionsDocumentAttributesError(url INSURL, options INSDictionary, dict INSDictionary) (NSAttributedString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithURL:options:documentAttributes:error:"), url, options, dict, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSAttributedString{}, NSErrorFrom(errorPtr)
	}
	return NSAttributedStringFromID(rv), nil

}
// Asks the item provider for the representation visibility specification for
// the given UTI.
//
// typeIdentifier: A uniform type identifier (UTI).
//
// # Return Value
// 
// A representation visibility specification for the given UTI.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProviderWriting/itemProviderVisibilityForRepresentation(withTypeIdentifier:)-swift.method
func (a NSAttributedString) ItemProviderVisibilityForRepresentationWithTypeIdentifier(typeIdentifier string) NSItemProviderRepresentationVisibility {
	rv := objc.Send[NSItemProviderRepresentationVisibility](a.ID, objc.Sel("itemProviderVisibilityForRepresentationWithTypeIdentifier:"), objc.String(typeIdentifier))
	return NSItemProviderRepresentationVisibility(rv)
}
// Loads data of a particular type, identified by the given UTI.
//
// typeIdentifier: The uniform type identifier (UTI) identifying the type of data to load.
//
// completionHandler: The handler that’s called after the data is loaded.
//
// # Return Value
// 
// The progress of the data load process.
//
// # Discussion
// 
// When the system calls this method, the `typeIdentifier` parameter is set to
// one of the elements in the `writableTypeIdentifiersForItemProvider` array.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProviderWriting/loadData(withTypeIdentifier:forItemProviderCompletionHandler:)
func (a NSAttributedString) LoadDataWithTypeIdentifierForItemProviderCompletionHandler(typeIdentifier string, completionHandler DataErrorHandler) INSProgress {
_block1, _cleanup1 := NewDataErrorBlock(completionHandler)
	defer _cleanup1()
	rv := objc.Send[objc.ID](a.ID, objc.Sel("loadDataWithTypeIdentifier:forItemProviderCompletionHandler:"), objc.String(typeIdentifier), _block1)
	return NSProgressFromID(rv)
}

// Creates an attributed string from the specified HTML data.
//
// data: A data object with text in HTML format. The method uses this data to create
// the attributed string.
//
// options: Specifies additional options for loading the document. For a list of
// possible keys, see [NSAttributedStringDocumentReadingOptionKey].
// //
// [NSAttributedStringDocumentReadingOptionKey]: https://developer.apple.com/documentation/UIKit/NSAttributedStringDocumentReadingOptionKey
//
// completionHandler: A completion handler to execute with the results.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/loadFromHTML(data:options:completionHandler:)
func (_NSAttributedStringClass NSAttributedStringClass) LoadFromHTMLWithDataOptionsCompletionHandler(data INSData, options INSDictionary, completionHandler ErrorHandler) {
_block2, _cleanup2 := NewErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](objc.ID(_NSAttributedStringClass.class), objc.Sel("loadFromHTMLWithData:options:completionHandler:"), data, options, _block2)
}
// Creates an attributed string by converting the content of a local HTML file
// at the specified URL.
//
// fileURL: A URL that specifies the file to load.
//
// options: Specifies additional options for loading the document. For a list of
// possible keys, see [NSAttributedStringDocumentReadingOptionKey].
// //
// [NSAttributedStringDocumentReadingOptionKey]: https://developer.apple.com/documentation/UIKit/NSAttributedStringDocumentReadingOptionKey
//
// completionHandler: A completion handler to execute with the results.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/loadFromHTML(fileURL:options:completionHandler:)
func (_NSAttributedStringClass NSAttributedStringClass) LoadFromHTMLWithFileURLOptionsCompletionHandler(fileURL INSURL, options INSDictionary, completionHandler ErrorHandler) {
_block2, _cleanup2 := NewErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](objc.ID(_NSAttributedStringClass.class), objc.Sel("loadFromHTMLWithFileURL:options:completionHandler:"), fileURL, options, _block2)
}
// Creates an attributed string by converting the contents of the specified
// HTML URL request.
//
// request: A request to fetch an HTML document.
//
// options: Specifies additional options for loading the document. For a list of
// possible keys, see [NSAttributedStringDocumentReadingOptionKey].
// //
// [NSAttributedStringDocumentReadingOptionKey]: https://developer.apple.com/documentation/UIKit/NSAttributedStringDocumentReadingOptionKey
//
// completionHandler: A completion handler to execute with the results.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/loadFromHTML(request:options:completionHandler:)
func (_NSAttributedStringClass NSAttributedStringClass) LoadFromHTMLWithRequestOptionsCompletionHandler(request INSURLRequest, options INSDictionary, completionHandler ErrorHandler) {
_block2, _cleanup2 := NewErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](objc.ID(_NSAttributedStringClass.class), objc.Sel("loadFromHTMLWithRequest:options:completionHandler:"), request, options, _block2)
}
// Creates an attributed string from the specified HTML string.
//
// string: A string that contains the HTML to convert to an attributed string.
//
// options: Specifies additional options for loading the document. For a list of
// possible keys, see [NSAttributedStringDocumentReadingOptionKey].
// //
// [NSAttributedStringDocumentReadingOptionKey]: https://developer.apple.com/documentation/UIKit/NSAttributedStringDocumentReadingOptionKey
//
// completionHandler: A completion handler to execute with the results.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/loadFromHTML(string:options:completionHandler:)
func (_NSAttributedStringClass NSAttributedStringClass) LoadFromHTMLWithStringOptionsCompletionHandler(string_ string, options INSDictionary, completionHandler ErrorHandler) {
_block2, _cleanup2 := NewErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](objc.ID(_NSAttributedStringClass.class), objc.Sel("loadFromHTMLWithString:options:completionHandler:"), objc.String(string_), options, _block2)
}
// Creates an attributed string by substituting arguments into a specially
// formatted string.
//
// format: The format string to use to create the final string. For a list of format
// specifiers you can include in this string, see [String Format Specifiers].
// //
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
//
// # Return Value
// 
// A new attributed string that combines the format string with the provided
// arguments and other information.
//
// # Discussion
// 
// Pass an optional list of trailing variadic arguments to substitute into the
// `format` string.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/localizedAttributedStringWithFormat:
func (_NSAttributedStringClass NSAttributedStringClass) LocalizedAttributedStringWithFormat(format INSAttributedString) NSAttributedString {
	rv := objc.Send[objc.ID](objc.ID(_NSAttributedStringClass.class), objc.Sel("localizedAttributedStringWithFormat:"), format)
	return NSAttributedStringFromID(rv)
}
// Creates an attributed string by substituting arguments into a specially
// formatted string and applying additional contextual information.
//
// format: The format string to use to create the final string. For a list of format
// specifiers you can include in this string, see [String Format Specifiers].
// //
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
//
// context: Additional options to apply to the string.
//
// # Return Value
// 
// A new attributed string that combines the format string with the provided
// arguments and other information.
//
// # Discussion
// 
// Pass an optional list of trailing variadic arguments to substitute into the
// `format` string.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/localizedAttributedStringWithFormat:context:
func (_NSAttributedStringClass NSAttributedStringClass) LocalizedAttributedStringWithFormatContext(format INSAttributedString, context INSDictionary) NSAttributedString {
	rv := objc.Send[objc.ID](objc.ID(_NSAttributedStringClass.class), objc.Sel("localizedAttributedStringWithFormat:context:"), format, context)
	return NSAttributedStringFromID(rv)
}
// Creates an attributed string by substituting a list of function arguments
// into a specially formatted string.
//
// format: The format string to use to create the final string. For a list of format
// specifiers you can include in this string, see [String Format Specifiers].
// //
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
//
// options: Options for how to apply attributes to the string’s content.
//
// # Return Value
// 
// A new attributed string that combines the format string with the provided
// arguments and other information.
//
// # Discussion
// 
// Pass an optional list of trailing variadic arguments to substitute into the
// `format` string.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/localizedAttributedStringWithFormat:options:
func (_NSAttributedStringClass NSAttributedStringClass) LocalizedAttributedStringWithFormatOptions(format INSAttributedString, options NSAttributedStringFormattingOptions) NSAttributedString {
	rv := objc.Send[objc.ID](objc.ID(_NSAttributedStringClass.class), objc.Sel("localizedAttributedStringWithFormat:options:"), format, options)
	return NSAttributedStringFromID(rv)
}
// Creates an attributed string by substituting a list of function arguments
// into a specially formatted string and applying additional contextual
// information.
//
// format: The format string to use to create the final string. For a list of format
// specifiers you can include in this string, see [String Format Specifiers].
// //
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
//
// options: Options for how to apply attributes to the string’s content.
//
// context: Additional options to apply to the string.
//
// # Return Value
// 
// A new attributed string that combines the format string with the provided
// arguments and other information.
//
// # Discussion
// 
// Pass an optional list of trailing variadic arguments to substitute into the
// `format` string.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/localizedAttributedStringWithFormat:options:context:
func (_NSAttributedStringClass NSAttributedStringClass) LocalizedAttributedStringWithFormatOptionsContext(format INSAttributedString, options NSAttributedStringFormattingOptions, context INSDictionary) NSAttributedString {
	rv := objc.Send[objc.ID](objc.ID(_NSAttributedStringClass.class), objc.Sel("localizedAttributedStringWithFormat:options:context:"), format, options, context)
	return NSAttributedStringFromID(rv)
}
// Creates a new instance of a class using the given data and UTI string.
//
// data: The data used to create the object.
//
// typeIdentifier: The uniform type identifier (UTI) representing the data type of `data`.
//
// # Return Value
// 
// An object created from the given data.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProviderReading/object(withItemProviderData:typeIdentifier:)
func (_NSAttributedStringClass NSAttributedStringClass) ObjectWithItemProviderDataTypeIdentifierError(data INSData, typeIdentifier string) (NSAttributedString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSAttributedStringClass.class), objc.Sel("objectWithItemProviderData:typeIdentifier:error:"), data, objc.String(typeIdentifier), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSAttributedString{}, NSErrorFrom(errorPtr)
	}
	return NSAttributedStringFromID(rv), nil

}

// The character contents of the attributed string as a string.
//
// # Discussion
// 
// Attachment characters are not removed from the value of this property.
// 
// For performance reasons, this property returns the current backing store of
// the attributed string object. If you want to maintain a snapshot of this as
// you manipulate the returned string, you should make a copy of the
// appropriate substring.
// 
// This primitive property must guarantee efficient access to an attributed
// string’s characters; subclasses should implement it to execute in O(1)
// time.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/string
func (a NSAttributedString) String() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("string"))
	return NSStringFromID(rv).String()
}
// The length of the attributed string.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/length
func (a NSAttributedString) Length() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("length"))
	return rv
}
// The string encoding for the document.
//
// See: https://developer.apple.com/documentation/foundation/nsattributedstring/documentattributekey/characterencoding
func (a NSAttributedString) CharacterEncoding() INSString {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("characterEncoding"))
	return NSStringFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the attribute string contains any
// attachment attributes.
//
// # Return Value
// 
// YES if the attributed string contains any attachment attributes, otherwise
// NO.
// 
// # Discussion
// 
// This method checks only for attachment attributes, not for
// [NSAttachmentCharacter].
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/containsAttachments
func (a NSAttributedString) ContainsAttachments() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("containsAttachments"))
	return rv
}
// The HTML elements to exclude in generated HTML.
//
// See: https://developer.apple.com/documentation/foundation/nsattributedstring/documentattributekey/excludedelements
func (a NSAttributedString) ExcludedElements() INSString {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("excludedElements"))
	return NSStringFromID(objc.ID(rv))
}
// The link for the text.
//
// See: https://developer.apple.com/documentation/foundation/nsattributedstring/key/link
func (a NSAttributedString) Link() NSAttributedStringKey {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("link"))
	return NSAttributedStringKey(NSStringFromID(rv).String())
}
// The number of spaces for indenting nested HTML elements.
//
// See: https://developer.apple.com/documentation/foundation/nsattributedstring/documentattributekey/prefixspaces
func (a NSAttributedString) PrefixSpaces() INSString {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("prefixSpaces"))
	return NSStringFromID(objc.ID(rv))
}
// The name of the text encoding to use.
//
// See: https://developer.apple.com/documentation/foundation/nsattributedstring/documentattributekey/textencodingname
func (a NSAttributedString) TextEncodingName() INSString {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("textEncodingName"))
	return NSStringFromID(objc.ID(rv))
}
// An array of UTI strings representing the types of data that can be loaded
// for an item provider.
//
// # Discussion
// 
// Provide uniform type identifiers (UTIs) in order from highest fidelity to
// lowest. If your app employs a native data representation, place that first
// in the array.
// 
// Use the instance version of this property when you initialize an item
// provider with an object. As possible, implement this property to provide an
// extended array of UTIs based on the object. For example, for an [NSURL]
// object, your implementation could offer the `public.File()-url` UTI, in
// addition to the `public.Url()` UTI, if your implementation detects that the
// stored URL uses the `//` scheme.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProviderWriting/writableTypeIdentifiersForItemProvider-swift.property
func (a NSAttributedString) WritableTypeIdentifiersForItemProvider() []string {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("writableTypeIdentifiersForItemProvider"))
	return objc.ConvertSliceToStrings(rv)
}

// An array of UTI strings that identify the file types that attributed
// strings support, either directly or through a user-installed filter
// service.
//
// # Return Value
// 
// An array of [NSString] objects, each of which contains a UTI identifying a
// supported file type.
// 
// # Discussion
// 
// The returned list includes UTIs all file types supported by the receiver
// plus those that can be opened by the receiver after being converted by a
// user-installed filter service. You can use the returned UTI strings with
// any method that supports UTIs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/textTypes
func (_NSAttributedStringClass NSAttributedStringClass) TextTypes() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSAttributedStringClass.class), objc.Sel("textTypes"))
	return objc.ConvertSliceToStrings(rv)
}
// An array of UTI strings that identify the file types that attributed
// strings support directly.
//
// # Return Value
// 
// An array of [NSString] objects, each of which contains a UTI identifying a
// supported file type.
// 
// # Discussion
// 
// The returned list includes UTI strings only for those file types that are
// supported directly by the receiver. It does not include types that are
// supported through user-installed filter services. You can use the returned
// UTI strings with any method that supports UTIs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/textUnfilteredTypes
func (_NSAttributedStringClass NSAttributedStringClass) TextUnfilteredTypes() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSAttributedStringClass.class), objc.Sel("textUnfilteredTypes"))
	return objc.ConvertSliceToStrings(rv)
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSItemProviderReading
			

			// Protocol methods for NSItemProviderWriting
			

			// Protocol methods for NSMutableCopying
			

			// Protocol methods for NSSecureCoding
			

// EnumerateAttributeInRangeOptionsUsingBlockSync is a synchronous wrapper around [NSAttributedString.EnumerateAttributeInRangeOptionsUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (a NSAttributedString) EnumerateAttributeInRangeOptionsUsingBlockSync(ctx context.Context, attrName NSAttributedStringKey, enumerationRange NSRange, opts NSAttributedStringEnumerationOptions) (objectivec.IObject, error) {
	done := make(chan objectivec.IObject, 1)
	a.EnumerateAttributeInRangeOptionsUsingBlock(attrName, enumerationRange, opts, func(val objectivec.IObject) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// LoadDataWithTypeIdentifierForItemProvider is a synchronous wrapper around [NSAttributedString.LoadDataWithTypeIdentifierForItemProviderCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a NSAttributedString) LoadDataWithTypeIdentifierForItemProvider(ctx context.Context, typeIdentifier string) (*NSData, error) {
	type result struct {
		val *NSData
		err error
	}
	done := make(chan result, 1)
	a.LoadDataWithTypeIdentifierForItemProviderCompletionHandler(typeIdentifier, func(val *NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

