// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMutableAttributedString] class.
var (
	_NSMutableAttributedStringClass     NSMutableAttributedStringClass
	_NSMutableAttributedStringClassOnce sync.Once
)

func getNSMutableAttributedStringClass() NSMutableAttributedStringClass {
	_NSMutableAttributedStringClassOnce.Do(func() {
		_NSMutableAttributedStringClass = NSMutableAttributedStringClass{class: objc.GetClass("NSMutableAttributedString")}
	})
	return _NSMutableAttributedStringClass
}

// GetNSMutableAttributedStringClass returns the class object for NSMutableAttributedString.
func GetNSMutableAttributedStringClass() NSMutableAttributedStringClass {
	return getNSMutableAttributedStringClass()
}

type NSMutableAttributedStringClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMutableAttributedStringClass) Alloc() NSMutableAttributedString {
	rv := objc.Send[NSMutableAttributedString](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A mutable string with associated attributes (such as visual style,
// hyperlinks, or accessibility data) for portions of its text.
//
// # Overview
// 
// The [NSMutableAttributedString] class declares additional methods for
// mutating the content of an attributed string. You can add and remove
// characters (raw strings) and attributes separately or together as
// attributed strings. See the class description for [NSAttributedString] for
// more information about attributed strings.
// 
// [NSMutableAttributedString] adds two primitive methods to those of
// [NSAttributedString]. These primitive methods provide the basis for all the
// other methods in its class. The primitive
// [NSMutableAttributedString.ReplaceCharactersInRangeWithString] method replaces a range of characters
// with those from a string, leaving all attribute information outside that
// range intact. The primitive [NSMutableAttributedString.SetAttributesRange] method sets attributes and
// values for a given range of characters, replacing any previous attributes
// and values for that range.
// 
// In macOS, AppKit also uses [NSParagraphStyle] and its subclass
// [NSMutableParagraphStyle] to encapsulate the paragraph or ruler attributes
// used by the [NSAttributedString] classes.
// 
// Note that the default font for [NSAttributedString] objects is Helvetica
// 12-point, which may differ from the macOS system font, so you may wish to
// create the string with non-default attributes suitable for your application
// using, for example, [NSMutableAttributedString.InitWithStringAttributes].
// 
// [NSMutableAttributedString] is “toll-free bridged” with its Core
// Foundation counterpart, [CFMutableAttributedString]. See [Toll-Free
// Bridging] for more information.
//
// [CFMutableAttributedString]: https://developer.apple.com/documentation/CoreFoundation/CFMutableAttributedString
// [NSMutableParagraphStyle]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle
// [NSParagraphStyle]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
//
// # Retrieving Character Information
//
//   - [NSMutableAttributedString.MutableString]: The character contents of the receiver as a mutable string object.
//
// # Changing Characters
//
//   - [NSMutableAttributedString.ReplaceCharactersInRangeWithString]: Replaces the characters in the given range with the characters of the given string.
//   - [NSMutableAttributedString.DeleteCharactersInRange]: Deletes the characters in the given range along with their associated attributes.
//
// # Changing Attributes
//
//   - [NSMutableAttributedString.SetAttributesRange]: Sets the attributes for the characters in the specified range to the specified attributes.
//   - [NSMutableAttributedString.AddAttributeValueRange]: Adds an attribute with the given name and value to the characters in the specified range.
//   - [NSMutableAttributedString.AddAttributesRange]: Adds the given collection of attributes to the characters in the specified range.
//   - [NSMutableAttributedString.RemoveAttributeRange]: Removes the named attribute from the characters in the specified range.
//   - [NSMutableAttributedString.ApplyFontTraitsRange]: Applies the specified font-related attributes to characters in the string.
//   - [NSMutableAttributedString.SetAlignmentRange]: Sets the alignment characteristic of the paragraph style attribute for the specified range of text.
//   - [NSMutableAttributedString.SetBaseWritingDirectionRange]: Sets the base writing direction for the characters to the specified direction.
//   - [NSMutableAttributedString.SubscriptRange]: Decrements the value of the superscript attribute for characters in the specified range by one.
//   - [NSMutableAttributedString.SuperscriptRange]: Increments the value of the superscript attribute for characters in the specified range by one.
//   - [NSMutableAttributedString.UnscriptRange]: Removes the superscript attribute from the characters in the specified range.
//
// # Changing Characters and Attributes
//
//   - [NSMutableAttributedString.AppendAttributedString]: Adds the characters and attributes of a given attributed string to the end of the receiver.
//   - [NSMutableAttributedString.InsertAttributedStringAtIndex]: Inserts the characters and attributes of the given attributed string into the receiver at the given index.
//   - [NSMutableAttributedString.ReplaceCharactersInRangeWithAttributedString]: Replaces the characters and attributes in a given range with the characters and attributes of the given attributed string.
//   - [NSMutableAttributedString.SetAttributedString]: Replaces the receiver’s entire contents with the characters and attributes of the given attributed string.
//
// # Grouping Changes
//
//   - [NSMutableAttributedString.BeginEditing]: Begins the buffering of changes to the string’s characters and attributes.
//   - [NSMutableAttributedString.EndEditing]: Ends the buffering of changes to the string’s characters and attributes.
//
// # Updating Attachment Contents
//
//   - [NSMutableAttributedString.UpdateAttachmentsFromPath]: Updates all attachments based on files contained in the RTFD file package at the specified file path.
//
// # Fixing Attributes After Changes
//
//   - [NSMutableAttributedString.FixAttributesInRange]: Cleans up font, paragraph style, and attachment attributes within the given range.
//   - [NSMutableAttributedString.FixAttachmentAttributeInRange]: Cleans up attachment attributes in the specified range and removes all attachment attributes assigned to characters except the designated attachment character.
//   - [NSMutableAttributedString.FixFontAttributeInRange]: Fixes the font attribute in the specified range and assigns default fonts where appropriate.
//   - [NSMutableAttributedString.FixParagraphStyleAttributeInRange]: Fixes the paragraph style attributes in the specified range and assigns a paragraph style to all characters in the paragraph.
//
// # Reading Content
//
//   - [NSMutableAttributedString.ReadFromDataOptionsDocumentAttributesError]: Sets the contents of the attributed string using the specified data object`.`
//   - [NSMutableAttributedString.ReadFromURLOptionsDocumentAttributesError]: Sets the contents of attributed string using the contents of the specified file.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString
type NSMutableAttributedString struct {
	NSAttributedString
}

// NSMutableAttributedStringFromID constructs a [NSMutableAttributedString] from an objc.ID.
//
// A mutable string with associated attributes (such as visual style,
// hyperlinks, or accessibility data) for portions of its text.
func NSMutableAttributedStringFromID(id objc.ID) NSMutableAttributedString {
	return NSMutableAttributedString{NSAttributedString: NSAttributedStringFromID(id)}
}
// NOTE: NSMutableAttributedString adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMutableAttributedString] class.
//
// # Retrieving Character Information
//
//   - [INSMutableAttributedString.MutableString]: The character contents of the receiver as a mutable string object.
//
// # Changing Characters
//
//   - [INSMutableAttributedString.ReplaceCharactersInRangeWithString]: Replaces the characters in the given range with the characters of the given string.
//   - [INSMutableAttributedString.DeleteCharactersInRange]: Deletes the characters in the given range along with their associated attributes.
//
// # Changing Attributes
//
//   - [INSMutableAttributedString.SetAttributesRange]: Sets the attributes for the characters in the specified range to the specified attributes.
//   - [INSMutableAttributedString.AddAttributeValueRange]: Adds an attribute with the given name and value to the characters in the specified range.
//   - [INSMutableAttributedString.AddAttributesRange]: Adds the given collection of attributes to the characters in the specified range.
//   - [INSMutableAttributedString.RemoveAttributeRange]: Removes the named attribute from the characters in the specified range.
//   - [INSMutableAttributedString.ApplyFontTraitsRange]: Applies the specified font-related attributes to characters in the string.
//   - [INSMutableAttributedString.SetAlignmentRange]: Sets the alignment characteristic of the paragraph style attribute for the specified range of text.
//   - [INSMutableAttributedString.SetBaseWritingDirectionRange]: Sets the base writing direction for the characters to the specified direction.
//   - [INSMutableAttributedString.SubscriptRange]: Decrements the value of the superscript attribute for characters in the specified range by one.
//   - [INSMutableAttributedString.SuperscriptRange]: Increments the value of the superscript attribute for characters in the specified range by one.
//   - [INSMutableAttributedString.UnscriptRange]: Removes the superscript attribute from the characters in the specified range.
//
// # Changing Characters and Attributes
//
//   - [INSMutableAttributedString.AppendAttributedString]: Adds the characters and attributes of a given attributed string to the end of the receiver.
//   - [INSMutableAttributedString.InsertAttributedStringAtIndex]: Inserts the characters and attributes of the given attributed string into the receiver at the given index.
//   - [INSMutableAttributedString.ReplaceCharactersInRangeWithAttributedString]: Replaces the characters and attributes in a given range with the characters and attributes of the given attributed string.
//   - [INSMutableAttributedString.SetAttributedString]: Replaces the receiver’s entire contents with the characters and attributes of the given attributed string.
//
// # Grouping Changes
//
//   - [INSMutableAttributedString.BeginEditing]: Begins the buffering of changes to the string’s characters and attributes.
//   - [INSMutableAttributedString.EndEditing]: Ends the buffering of changes to the string’s characters and attributes.
//
// # Updating Attachment Contents
//
//   - [INSMutableAttributedString.UpdateAttachmentsFromPath]: Updates all attachments based on files contained in the RTFD file package at the specified file path.
//
// # Fixing Attributes After Changes
//
//   - [INSMutableAttributedString.FixAttributesInRange]: Cleans up font, paragraph style, and attachment attributes within the given range.
//   - [INSMutableAttributedString.FixAttachmentAttributeInRange]: Cleans up attachment attributes in the specified range and removes all attachment attributes assigned to characters except the designated attachment character.
//   - [INSMutableAttributedString.FixFontAttributeInRange]: Fixes the font attribute in the specified range and assigns default fonts where appropriate.
//   - [INSMutableAttributedString.FixParagraphStyleAttributeInRange]: Fixes the paragraph style attributes in the specified range and assigns a paragraph style to all characters in the paragraph.
//
// # Reading Content
//
//   - [INSMutableAttributedString.ReadFromDataOptionsDocumentAttributesError]: Sets the contents of the attributed string using the specified data object`.`
//   - [INSMutableAttributedString.ReadFromURLOptionsDocumentAttributesError]: Sets the contents of attributed string using the contents of the specified file.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString
type INSMutableAttributedString interface {
	INSAttributedString

	// Topic: Retrieving Character Information

	// The character contents of the receiver as a mutable string object.
	MutableString() INSMutableString

	// Topic: Changing Characters

	// Replaces the characters in the given range with the characters of the given string.
	ReplaceCharactersInRangeWithString(range_ NSRange, str string)
	// Deletes the characters in the given range along with their associated attributes.
	DeleteCharactersInRange(range_ NSRange)

	// Topic: Changing Attributes

	// Sets the attributes for the characters in the specified range to the specified attributes.
	SetAttributesRange(attrs INSDictionary, range_ NSRange)
	// Adds an attribute with the given name and value to the characters in the specified range.
	AddAttributeValueRange(name NSAttributedStringKey, value objectivec.IObject, range_ NSRange)
	// Adds the given collection of attributes to the characters in the specified range.
	AddAttributesRange(attrs INSDictionary, range_ NSRange)
	// Removes the named attribute from the characters in the specified range.
	RemoveAttributeRange(name NSAttributedStringKey, range_ NSRange)
	// Applies the specified font-related attributes to characters in the string.
	ApplyFontTraitsRange(traitMask objectivec.IObject, range_ NSRange)
	// Sets the alignment characteristic of the paragraph style attribute for the specified range of text.
	SetAlignmentRange(alignment objectivec.IObject, range_ NSRange)
	// Sets the base writing direction for the characters to the specified direction.
	SetBaseWritingDirectionRange(writingDirection objectivec.IObject, range_ NSRange)
	// Decrements the value of the superscript attribute for characters in the specified range by one.
	SubscriptRange(range_ NSRange)
	// Increments the value of the superscript attribute for characters in the specified range by one.
	SuperscriptRange(range_ NSRange)
	// Removes the superscript attribute from the characters in the specified range.
	UnscriptRange(range_ NSRange)

	// Topic: Changing Characters and Attributes

	// Adds the characters and attributes of a given attributed string to the end of the receiver.
	AppendAttributedString(attrString INSAttributedString)
	// Inserts the characters and attributes of the given attributed string into the receiver at the given index.
	InsertAttributedStringAtIndex(attrString INSAttributedString, loc uint)
	// Replaces the characters and attributes in a given range with the characters and attributes of the given attributed string.
	ReplaceCharactersInRangeWithAttributedString(range_ NSRange, attrString INSAttributedString)
	// Replaces the receiver’s entire contents with the characters and attributes of the given attributed string.
	SetAttributedString(attrString INSAttributedString)

	// Topic: Grouping Changes

	// Begins the buffering of changes to the string’s characters and attributes.
	BeginEditing()
	// Ends the buffering of changes to the string’s characters and attributes.
	EndEditing()

	// Topic: Updating Attachment Contents

	// Updates all attachments based on files contained in the RTFD file package at the specified file path.
	UpdateAttachmentsFromPath(path string)

	// Topic: Fixing Attributes After Changes

	// Cleans up font, paragraph style, and attachment attributes within the given range.
	FixAttributesInRange(range_ NSRange)
	// Cleans up attachment attributes in the specified range and removes all attachment attributes assigned to characters except the designated attachment character.
	FixAttachmentAttributeInRange(range_ NSRange)
	// Fixes the font attribute in the specified range and assigns default fonts where appropriate.
	FixFontAttributeInRange(range_ NSRange)
	// Fixes the paragraph style attributes in the specified range and assigns a paragraph style to all characters in the paragraph.
	FixParagraphStyleAttributeInRange(range_ NSRange)

	// Topic: Reading Content

	// Sets the contents of the attributed string using the specified data object`.`
	ReadFromDataOptionsDocumentAttributesError(data INSData, opts INSDictionary, dict INSDictionary) (bool, error)
	// Sets the contents of attributed string using the contents of the specified file.
	ReadFromURLOptionsDocumentAttributesError(url INSURL, opts INSDictionary, dict INSDictionary) (bool, error)

	// Formats the specified string and arguments with the current locale, then appends the result to the receiver.
	AppendLocalizedFormat(format INSAttributedString)
}

// Init initializes the instance.
func (m NSMutableAttributedString) Init() NSMutableAttributedString {
	rv := objc.Send[NSMutableAttributedString](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMutableAttributedString) Autorelease() NSMutableAttributedString {
	rv := objc.Send[NSMutableAttributedString](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMutableAttributedString creates a new NSMutableAttributedString instance.
func NewNSMutableAttributedString() NSMutableAttributedString {
	class := getNSMutableAttributedStringClass()
	rv := objc.Send[NSMutableAttributedString](objc.ID(class.class), objc.Sel("new"))
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
func NewMutableAttributedStringWithAdaptiveImageGlyphAttributes(adaptiveImageGlyph objectivec.IObject, attributes INSDictionary) NSMutableAttributedString {
	rv := objc.Send[objc.ID](objc.ID(getNSMutableAttributedStringClass().class), objc.Sel("attributedStringWithAdaptiveImageGlyph:attributes:"), adaptiveImageGlyph, attributes)
	return NSMutableAttributedStringFromID(rv)
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
func NewMutableAttributedStringWithAttachmentAttributes(attachment objectivec.IObject, attributes INSDictionary) NSMutableAttributedString {
	rv := objc.Send[objc.ID](objc.ID(getNSMutableAttributedStringClass().class), objc.Sel("attributedStringWithAttachment:attributes:"), attachment, attributes)
	return NSMutableAttributedStringFromID(rv)
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
func NewMutableAttributedStringWithAttributedString(attrStr INSAttributedString) NSMutableAttributedString {
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAttributedString:"), attrStr)
	return NSMutableAttributedStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewMutableAttributedStringWithCoder(coder INSCoder) NSMutableAttributedString {
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSMutableAttributedStringFromID(rv)
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
func NewMutableAttributedStringWithContentsOfMarkdownFileAtURLOptionsBaseURLError(markdownFile INSURL, options INSAttributedStringMarkdownParsingOptions, baseURL INSURL) (NSMutableAttributedString, error) {
	var errorPtr objc.ID
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfMarkdownFileAtURL:options:baseURL:error:"), markdownFile, options, baseURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSMutableAttributedStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSMutableAttributedStringFromID(rv), nil
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
func NewMutableAttributedStringWithDataOptionsDocumentAttributesError(data INSData, options INSDictionary, dict INSDictionary) (NSMutableAttributedString, error) {
	var errorPtr objc.ID
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:options:documentAttributes:error:"), data, options, dict, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSMutableAttributedStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSMutableAttributedStringFromID(rv), nil
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
func NewMutableAttributedStringWithDocFormatDocumentAttributes(data INSData, dict INSDictionary) NSMutableAttributedString {
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDocFormat:documentAttributes:"), data, dict)
	return NSMutableAttributedStringFromID(rv)
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
func NewMutableAttributedStringWithFileURLOptionsDocumentAttributesError(url INSURL, options INSDictionary, dict INSDictionary) (NSMutableAttributedString, error) {
	var errorPtr objc.ID
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFileURL:options:documentAttributes:error:"), url, options, dict, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSMutableAttributedStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSMutableAttributedStringFromID(rv), nil
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
func NewMutableAttributedStringWithFormatOptionsLocale(format INSAttributedString, options NSAttributedStringFormattingOptions, locale INSLocale) NSMutableAttributedString {
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:options:locale:"), format, options, locale)
	return NSMutableAttributedStringFromID(rv)
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
func NewMutableAttributedStringWithFormatOptionsLocaleArguments(format INSAttributedString, options NSAttributedStringFormattingOptions, locale INSLocale, arguments unsafe.Pointer) NSMutableAttributedString {
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:options:locale:arguments:"), format, options, locale, arguments)
	return NSMutableAttributedStringFromID(rv)
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
func NewMutableAttributedStringWithFormatOptionsLocaleContext(format INSAttributedString, options NSAttributedStringFormattingOptions, locale INSLocale, context INSDictionary) NSMutableAttributedString {
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:options:locale:context:"), format, options, locale, context)
	return NSMutableAttributedStringFromID(rv)
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
func NewMutableAttributedStringWithFormatOptionsLocaleContextArguments(format INSAttributedString, options NSAttributedStringFormattingOptions, locale INSLocale, context INSDictionary, arguments unsafe.Pointer) NSMutableAttributedString {
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:options:locale:context:arguments:"), format, options, locale, context, arguments)
	return NSMutableAttributedStringFromID(rv)
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
func NewMutableAttributedStringWithHTMLBaseURLDocumentAttributes(data INSData, base INSURL, dict INSDictionary) NSMutableAttributedString {
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithHTML:baseURL:documentAttributes:"), data, base, dict)
	return NSMutableAttributedStringFromID(rv)
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
func NewMutableAttributedStringWithHTMLDocumentAttributes(data INSData, dict INSDictionary) NSMutableAttributedString {
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithHTML:documentAttributes:"), data, dict)
	return NSMutableAttributedStringFromID(rv)
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
func NewMutableAttributedStringWithHTMLOptionsDocumentAttributes(data INSData, options INSDictionary, dict INSDictionary) NSMutableAttributedString {
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithHTML:options:documentAttributes:"), data, options, dict)
	return NSMutableAttributedStringFromID(rv)
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
func NewMutableAttributedStringWithMarkdownOptionsBaseURLError(markdown INSData, options INSAttributedStringMarkdownParsingOptions, baseURL INSURL) (NSMutableAttributedString, error) {
	var errorPtr objc.ID
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMarkdown:options:baseURL:error:"), markdown, options, baseURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSMutableAttributedStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSMutableAttributedStringFromID(rv), nil
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
func NewMutableAttributedStringWithMarkdownStringOptionsBaseURLError(markdownString string, options INSAttributedStringMarkdownParsingOptions, baseURL INSURL) (NSMutableAttributedString, error) {
	var errorPtr objc.ID
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMarkdownString:options:baseURL:error:"), objc.String(markdownString), options, baseURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSMutableAttributedStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSMutableAttributedStringFromID(rv), nil
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
func NewMutableAttributedStringWithRTFDDocumentAttributes(data INSData, dict INSDictionary) NSMutableAttributedString {
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRTFD:documentAttributes:"), data, dict)
	return NSMutableAttributedStringFromID(rv)
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
func NewMutableAttributedStringWithRTFDFileWrapperDocumentAttributes(wrapper INSFileWrapper, dict INSDictionary) NSMutableAttributedString {
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRTFDFileWrapper:documentAttributes:"), wrapper, dict)
	return NSMutableAttributedStringFromID(rv)
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
func NewMutableAttributedStringWithRTFDocumentAttributes(data INSData, dict INSDictionary) NSMutableAttributedString {
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRTF:documentAttributes:"), data, dict)
	return NSMutableAttributedStringFromID(rv)
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
func NewMutableAttributedStringWithString(str string) NSMutableAttributedString {
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:"), objc.String(str))
	return NSMutableAttributedStringFromID(rv)
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
func NewMutableAttributedStringWithStringAttributes(str string, attrs INSDictionary) NSMutableAttributedString {
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:attributes:"), objc.String(str), attrs)
	return NSMutableAttributedStringFromID(rv)
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
func NewMutableAttributedStringWithURLOptionsDocumentAttributesError(url INSURL, options INSDictionary, dict INSDictionary) (NSMutableAttributedString, error) {
	var errorPtr objc.ID
	instance := getNSMutableAttributedStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:options:documentAttributes:error:"), url, options, dict, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSMutableAttributedStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSMutableAttributedStringFromID(rv), nil
}

// Replaces the characters in the given range with the characters of the given
// string.
//
// range: A range specifying the characters to replace.
//
// str: A string specifying the characters to replace those in `range`.
//
// # Discussion
// 
// The new characters inherit the attributes of the first replaced character
// from `range`. Where the length of `range` is 0, the new characters inherit
// the attributes of the character preceding `range` if it has any, otherwise
// of the character following `range`.
// 
// Raises an [rangeException] if any part of `range` lies beyond the end of
// the receiver’s characters.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/replaceCharacters(in:with:)-6oq9r
func (m NSMutableAttributedString) ReplaceCharactersInRangeWithString(range_ NSRange, str string) {
	objc.Send[objc.ID](m.ID, objc.Sel("replaceCharactersInRange:withString:"), range_, objc.String(str))
}
// Deletes the characters in the given range along with their associated
// attributes.
//
// range: A range specifying the characters to delete.
//
// # Discussion
// 
// Raises an [rangeException] if any part of `range` lies beyond the end of
// the receiver’s characters.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/deleteCharacters(in:)
func (m NSMutableAttributedString) DeleteCharactersInRange(range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("deleteCharactersInRange:"), range_)
}
// Sets the attributes for the characters in the specified range to the
// specified attributes.
//
// attrs: A dictionary containing the attributes to set. Attribute keys can be
// supplied by another framework or can be custom ones you define. For
// information about the system-supplied attribute keys, see the Constants
// section in [NSAttributedString].
//
// range: The range of characters whose attributes are set.
//
// # Discussion
// 
// These new attributes replace any attributes previously associated with the
// characters in `range`. Raises an [rangeException] if any part of `range`
// lies beyond the end of the receiver’s characters.
// 
// To set attributes for a zero-length [NSMutableAttributedString] displayed
// in a text view, use the [NSTextView] method [typingAttributes].
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
// [typingAttributes]: https://developer.apple.com/documentation/AppKit/NSTextView/typingAttributes
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/setAttributes(_:range:)
func (m NSMutableAttributedString) SetAttributesRange(attrs INSDictionary, range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAttributes:range:"), attrs, range_)
}
// Adds an attribute with the given name and value to the characters in the
// specified range.
//
// name: A string specifying the attribute name. Attribute keys can be supplied by
// another framework or can be custom ones you define. For information about
// the system-supplied attribute keys, see the Constants section in
// [NSAttributedString].
//
// value: The attribute value associated with `name`.
//
// range: The range of characters to which the specified attribute/value pair
// applies.
//
// # Discussion
// 
// You may assign any `name`/`value` pair you wish to a range of characters.
// Raises an [invalidArgumentException] if `name` or `value` is `nil` and an
// [rangeException] if any part of `range` lies beyond the end of the
// receiver’s characters.
//
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/addAttribute(_:value:range:)
func (m NSMutableAttributedString) AddAttributeValueRange(name NSAttributedStringKey, value objectivec.IObject, range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("addAttribute:value:range:"), objc.String(string(name)), value, range_)
}
// Adds the given collection of attributes to the characters in the specified
// range.
//
// attrs: A dictionary containing the attributes to add. Attribute keys can be
// supplied by another framework or can be custom ones you define. For
// information about the system-supplied attribute keys, see the Constants
// section in [NSAttributedString].
//
// range: The range of characters to which the specified attributes apply.
//
// # Discussion
// 
// You may assign any name/value pair you wish to a range of characters.
// Raises an [invalidArgumentException] if `attrs` is `nil` and an
// [rangeException] if any part of `range` lies beyond the end of the
// receiver’s characters.
//
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/addAttributes(_:range:)
func (m NSMutableAttributedString) AddAttributesRange(attrs INSDictionary, range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("addAttributes:range:"), attrs, range_)
}
// Removes the named attribute from the characters in the specified range.
//
// name: A string specifying the attribute name to remove. Attribute keys can be
// supplied by another framework or can be custom ones you define. For
// information about where to find the system-supplied attribute keys, see the
// overview section in [NSAttributedString].
//
// range: The range of characters from which the specified attribute is removed.
//
// # Discussion
// 
// Raises an [rangeException] if any part of `range` lies beyond the end of
// the receiver’s characters.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/removeAttribute(_:range:)
func (m NSMutableAttributedString) RemoveAttributeRange(name NSAttributedStringKey, range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeAttribute:range:"), objc.String(string(name)), range_)
}
// Applies the specified font-related attributes to characters in the string.
//
// traitMask: The font attributes to apply. For information about the font traits you can
// apply, see [NSFontManager].
// //
// [NSFontManager]: https://developer.apple.com/documentation/AppKit/NSFontManager
//
// range: The range of characters.
//
// traitMask is a [appkit.NSFontTraitMask].
//
// # Discussion
// 
// Raises an [rangeException] if any part of `range` lies beyond the end of
// the receiver’s characters.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/applyFontTraits(_:range:)
func (m NSMutableAttributedString) ApplyFontTraitsRange(traitMask objectivec.IObject, range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("applyFontTraits:range:"), traitMask, range_)
}
// Sets the alignment characteristic of the paragraph style attribute for the
// specified range of text.
//
// alignment: The alignment to use.
//
// range: The range of characters.
//
// alignment is a [appkit.NSTextAlignment].
//
// # Discussion
// 
// When attribute fixing takes place, this change will affect only paragraphs
// whose first character was included in `range`. Raises an [rangeException]
// if any part of `range` lies beyond the end of the receiver’s characters.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/setAlignment(_:range:)
func (m NSMutableAttributedString) SetAlignmentRange(alignment objectivec.IObject, range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAlignment:range:"), alignment, range_)
}
// Sets the base writing direction for the characters to the specified
// direction.
//
// writingDirection: The writing direction to use.
//
// range: The range of characters.
//
// writingDirection is a [uikit.NSWritingDirection].
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/setBaseWritingDirection(_:range:)
func (m NSMutableAttributedString) SetBaseWritingDirectionRange(writingDirection objectivec.IObject, range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("setBaseWritingDirection:range:"), writingDirection, range_)
}
// Decrements the value of the superscript attribute for characters in the
// specified range by one.
//
// range: The range of characters.
//
// # Discussion
// 
// Raises an [rangeException] if any part of `range` lies beyond the end of
// the receiver’s characters.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/subscriptRange(_:)
func (m NSMutableAttributedString) SubscriptRange(range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("subscriptRange:"), range_)
}
// Increments the value of the superscript attribute for characters in the
// specified range by one.
//
// range: The range of characters.
//
// # Discussion
// 
// Raises an [rangeException] if any part of `range` lies beyond the end of
// the receiver’s characters.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/superscriptRange(_:)
func (m NSMutableAttributedString) SuperscriptRange(range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("superscriptRange:"), range_)
}
// Removes the superscript attribute from the characters in the specified
// range.
//
// range: The range of characters.
//
// # Discussion
// 
// Raises an [rangeException] if any part of `range` lies beyond the end of
// the receiver’s characters.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/unscriptRange(_:)
func (m NSMutableAttributedString) UnscriptRange(range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("unscriptRange:"), range_)
}
// Adds the characters and attributes of a given attributed string to the end
// of the receiver.
//
// attrString: The string whose characters and attributes are added.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/append(_:)
func (m NSMutableAttributedString) AppendAttributedString(attrString INSAttributedString) {
	objc.Send[objc.ID](m.ID, objc.Sel("appendAttributedString:"), attrString)
}
// Inserts the characters and attributes of the given attributed string into
// the receiver at the given index.
//
// attrString: The string whose characters and attributes are inserted.
//
// loc: The index at which the characters and attributes are inserted.
//
// # Discussion
// 
// The new characters and attributes begin at the given index and the existing
// characters and attributes from the index to the end of the receiver are
// shifted by the length of the attributed string. Raises an [rangeException]
// if `loc` lies beyond the end of the receiver’s characters.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/insert(_:at:)
func (m NSMutableAttributedString) InsertAttributedStringAtIndex(attrString INSAttributedString, loc uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("insertAttributedString:atIndex:"), attrString, loc)
}
// Replaces the characters and attributes in a given range with the characters
// and attributes of the given attributed string.
//
// range: The range of characters and attributes replaced.
//
// attrString: The attributed string whose characters and attributes replace those in the
// specified range.
//
// # Discussion
// 
// Raises an [rangeException] if any part of `range` lies beyond the end of
// the receiver’s characters.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/replaceCharacters(in:with:)-1uaw7
func (m NSMutableAttributedString) ReplaceCharactersInRangeWithAttributedString(range_ NSRange, attrString INSAttributedString) {
	objc.Send[objc.ID](m.ID, objc.Sel("replaceCharactersInRange:withAttributedString:"), range_, attrString)
}
// Replaces the receiver’s entire contents with the characters and
// attributes of the given attributed string.
//
// attrString: The attributed string whose characters and attributes replace those in the
// receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/setAttributedString(_:)
func (m NSMutableAttributedString) SetAttributedString(attrString INSAttributedString) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAttributedString:"), attrString)
}
// Begins the buffering of changes to the string’s characters and
// attributes.
//
// # Discussion
// 
// Override this method in a subclass to buffer or optimize a series of
// changes to the string’s characters or attributes. The string continues to
// buffer text until you call [EndEditing], at which time it consolidates the
// changes and notifies observers.
// 
// You can nest pairs of [BeginEditing] and [EndEditing] messages.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/beginEditing()
func (m NSMutableAttributedString) BeginEditing() {
	objc.Send[objc.ID](m.ID, objc.Sel("beginEditing"))
}
// Ends the buffering of changes to the string’s characters and attributes.
//
// # Discussion
// 
// Override this method in a subclass to consolidate changes made since a
// previous call to [BeginEditing]. When you call this method, the string
// notifies observers of the changes.
// 
// The default implementation of this method does nothing. Subclasses such as
// [NSTextStorage] override this method and use it to tell the layout manager
// to update the text layout.
//
// [NSTextStorage]: https://developer.apple.com/documentation/AppKit/NSTextStorage
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/endEditing()
func (m NSMutableAttributedString) EndEditing() {
	objc.Send[objc.ID](m.ID, objc.Sel("endEditing"))
}
// Updates all attachments based on files contained in the RTFD file package
// at the specified file path.
//
// path: The path to the file package.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/updateAttachments(fromPath:)
func (m NSMutableAttributedString) UpdateAttachmentsFromPath(path string) {
	objc.Send[objc.ID](m.ID, objc.Sel("updateAttachmentsFromPath:"), objc.String(path))
}
// Cleans up font, paragraph style, and attachment attributes within the given
// range.
//
// range: The character range within which to fix attributes. Raises an
// [rangeException] if any part of `range` lies beyond the end of the
// receiver’s characters.
// //
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// # Discussion
// 
// Removes attachment attributes assigned to characters other than
// [character], assigns default fonts to characters with illegal fonts for
// their scripts and otherwise corrects font attribute assignments, and
// assigns the first paragraph style attribute value in each paragraph to all
// characters of the paragraph.
// 
// This method extends the range as needed to cover the last paragraph
// partially contained.
// 
// Raises an [rangeException] if any part of aRange lies beyond the end of the
// receiver’s characters.
// 
// [NSTextStorage] subclasses that return [true] from the
// [fixesAttributesLazily] method should avoid directly calling
// [FixAttributesInRange] or else bracket such calls with [BeginEditing] and
// [EndEditing] messages.
//
// [character]: https://developer.apple.com/documentation/UIKit/NSTextAttachment/character
// [fixesAttributesLazily]: https://developer.apple.com/documentation/AppKit/NSTextStorage/fixesAttributesLazily
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/fixAttributes(in:)
func (m NSMutableAttributedString) FixAttributesInRange(range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("fixAttributesInRange:"), range_)
}
// Cleans up attachment attributes in the specified range and removes all
// attachment attributes assigned to characters except the designated
// attachment character.
//
// range: The range of characters.
//
// # Discussion
// 
// The method preserves the attachment attribute on the [character] special
// character. The method raises a [rangeException] if any part of `range` lies
// beyond the end of the string’s characters.
//
// [character]: https://developer.apple.com/documentation/UIKit/NSTextAttachment/character
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/fixAttachmentAttribute(in:)
func (m NSMutableAttributedString) FixAttachmentAttributeInRange(range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("fixAttachmentAttributeInRange:"), range_)
}
// Fixes the font attribute in the specified range and assigns default fonts
// where appropriate.
//
// range: The range of characters.
//
// # Discussion
// 
// This method assigns default fonts to characters with illegal fonts for
// their scripts and corrects other font attribute assignments. For example,
// Kanji characters assigned a Latin font are reassigned an appropriate Kanji
// font. Raises an [rangeException] if any part of `range` lies beyond the end
// of the receiver’s characters.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/fixFontAttribute(in:)
func (m NSMutableAttributedString) FixFontAttributeInRange(range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("fixFontAttributeInRange:"), range_)
}
// Fixes the paragraph style attributes in the specified range and assigns a
// paragraph style to all characters in the paragraph.
//
// range: The range of characters.
//
// # Discussion
// 
// This method assigns the first paragraph style attribute value in each
// paragraph to all characters of the paragraph. This method extends the range
// as needed to cover the last paragraph partially contained. A paragraph is
// delimited by any of these characters, the longest possible sequence being
// preferred to any shorter:
// 
// - U+000D (`\r` or CR) - U+000A (`\n` or LF) - U+2029 (Unicode paragraph
// separator) `\r\n`, in that order (also known as CRLF)
// 
// Raises an [rangeException] if any part of `range` lies beyond the end of
// the receiver’s characters.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/fixParagraphStyleAttribute(in:)
func (m NSMutableAttributedString) FixParagraphStyleAttributeInRange(range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("fixParagraphStyleAttributeInRange:"), range_)
}
// Sets the contents of the attributed string using the specified data
// object`.`
//
// data: The data object providing text data.
//
// opts: Keys specifying the types of documents and other document import options.
// For a list of possible values, see “Option keys for importing
// documents” in [NSAttributedString]
//
// dict: On return, the dictionary (if provided) contains keys representing various
// document-wide attributes. For a list of possible values, see “Document
// Attributes” in [NSAttributedString].
//
// error: Upon return, if an error occurs, contains an [NSError] object that
// describes the problem. If you are not interested in possible errors, pass
// in [NULL].
//
// # Discussion
// 
// `opts` can contain one of the values described in the Constants section of
// NSAttributedString Application Kit Additions Reference (“Option keys for
// importing documents”).
// 
// On return, the `documentAttributes` dictionary (if provided) contains the
// various keys described in the Constants section of NSAttributedString
// Application Kit Additions Reference. If unsuccessful, returns NO , after
// setting `error` to point to an [NSError] object that encapsulates the
// reason why the attributed string object could not be created.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/read(from:options:documentAttributes:)-5mbcx
func (m NSMutableAttributedString) ReadFromDataOptionsDocumentAttributesError(data INSData, opts INSDictionary, dict INSDictionary) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("readFromData:options:documentAttributes:error:"), data, opts, dict, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("readFromData:options:documentAttributes:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Sets the contents of attributed string using the contents of the specified
// file.
//
// url: The URL of the file to read.
//
// opts: The option keys for importing the document. For a list of possible values,
// see “Option keys for importing documents” in [NSAttributedString].
//
// dict: On return, contains the document attributes. For a list of possible values,
// see “Document Attributes” in [NSAttributedString].
//
// error: Upon return, if an error occurs, contains an [NSError] object that
// describes the problem. If you are not interested in possible errors, pass
// in [NULL].
//
// # Discussion
// 
// Filter services can be used to convert the contents of the URL into a
// format recognized by Cocoa.
// 
// For RTF formatted files, the contents of the file are appended to the
// previous string instead of replacing the previous string. Therefore, when
// using this method with existing content it’s best to clear the content
// away explicitly.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/read(from:options:documentAttributes:)-54wth
func (m NSMutableAttributedString) ReadFromURLOptionsDocumentAttributesError(url INSURL, opts INSDictionary, dict INSDictionary) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("readFromURL:options:documentAttributes:error:"), url, opts, dict, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("readFromURL:options:documentAttributes:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Formats the specified string and arguments with the current locale, then
// appends the result to the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/appendLocalizedFormat:
func (m NSMutableAttributedString) AppendLocalizedFormat(format INSAttributedString) {
	objc.Send[objc.ID](m.ID, objc.Sel("appendLocalizedFormat:"), format)
}

// The character contents of the receiver as a mutable string object.
//
// # Return Value
// 
// The mutable string object.
// 
// # Discussion
// 
// The receiver tracks changes to this string and keeps its attribute mappings
// up to date.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/mutableString
func (m NSMutableAttributedString) MutableString() INSMutableString {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mutableString"))
	return NSMutableStringFromID(objc.ID(rv))
}

