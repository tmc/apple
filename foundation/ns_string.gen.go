// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSString] class.
var (
	_NSStringClass     NSStringClass
	_NSStringClassOnce sync.Once
)

func getNSStringClass() NSStringClass {
	_NSStringClassOnce.Do(func() {
		_NSStringClass = NSStringClass{class: objc.GetClass("NSString")}
	})
	return _NSStringClass
}

// GetNSStringClass returns the class object for NSString.
func GetNSStringClass() NSStringClass {
	return getNSStringClass()
}

type NSStringClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSStringClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSStringClass) Alloc() NSString {
	rv := objc.Send[NSString](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A static, plain-text Unicode string object.
//
// # Overview
//
// You can use this type in Swift when you need reference semantics or other
// Foundation-specific behavior.
//
// The [NSString] class and its mutable subclass, [NSMutableString], provide
// an extensive set of APIs for working with strings, including methods for
// comparing, searching, and modifying strings. [NSString] objects are used
// throughout Foundation and other Cocoa frameworks, serving as the basis for
// all textual and linguistic functionality on the platform.
//
// [NSString] is with its Core Foundation counterpart, [CFString]. See
// [Toll-Free Bridging] for more information.
//
// # String Objects
//
// An [NSString] object encodes a Unicode-compliant text string, represented
// as a sequence of UTF–16 code units. All lengths, character indexes, and
// ranges are expressed in terms of 16-bit platform-endian values, with index
// values starting at `0`.
//
// An [NSString] object can be initialized from or written to a C buffer, an
// [NSData] object, or the contents of an [NSURL]. It can also be encoded and
// decoded to and from ASCII, UTF–8, UTF–16, UTF–32, or any other string
// encoding represented by [NSStringEncoding].
//
// The objects you create using [NSString] and [NSMutableString] are referred
// to as string objects (or, when no confusion will result, merely as
// strings). The term C string refers to the standard `char *` type.
//
// Because of the nature of class clusters, string objects aren’t actual
// instances of the [NSString] or [NSMutableString] classes but of one of
// their private subclasses. Although a string object’s class is private,
// its interface is public, as declared by these abstract superclasses,
// [NSString] and [NSMutableString]. The string classes adopt the [NSCopying]
// and [NSMutableCopying] protocols, making it convenient to convert a string
// of one type to the other.
//
// # Understanding Characters
//
// A string object presents itself as a sequence of UTF–16 code units. You
// can determine how many UTF-16 code units a string object contains with the
// [NSString.Length] method and can retrieve a specific UTF-16 code unit with the
// [NSString.CharacterAtIndex] method. These two “primitive” methods provide basic
// access to a string object.
//
// Most use of strings, however, is at a higher level, with the strings being
// treated as single entities: You compare strings against one another, search
// them for substrings, combine them into new strings, and so on. If you need
// to access string objects character by character, you must understand the
// Unicode character encoding, specifically issues related to composed
// character sequences. For details see (The Unicode Consortium, Boston:
// Addison-Wesley, 2003, ISBN 0-321-18578-1) and the Unicode Consortium web
// site: [http://www.unicode.org/]. See also [Characters and Grapheme
// Clusters] in [String Programming Guide].
//
// Localized string comparisons are based on the Unicode Collation Algorithm,
// as tailored for different languages by CLDR (Common Locale Data
// Repository). Both are projects of the Unicode Consortium. Unicode is a
// registered trademark of Unicode, Inc.
//
// # Interpreting UTF-16-Encoded Data
//
// When creating an [NSString] object from a UTF-16-encoded string (or a byte
// stream interpreted as UTF-16), if the byte order is not otherwise
// specified, [NSString] assumes that the UTF-16 characters are big-endian,
// unless there is a BOM (byte-order mark), in which case the BOM dictates the
// byte order. When creating an [NSString] object from an array of `unichar`
// values, the returned string is always native-endian, since the array always
// contains UTF–16 code units in native byte order.
//
// # Subclassing Notes
//
// It is possible to subclass [NSString] (and [NSMutableString]), but doing so
// requires providing storage facilities for the string (which is not
// inherited by subclasses) and implementing two primitive methods. The
// abstract [NSString] and [NSMutableString] classes are the public interface
// of a class cluster consisting mostly of private, concrete classes that
// create and return a string object appropriate for a given situation. Making
// your own concrete subclass of this cluster imposes certain requirements
// (discussed in [NSString]).
//
// Make sure your reasons for subclassing [NSString] are valid. Instances of
// your subclass should represent a string and not something else. Thus the
// only attributes the subclass should have are the length of the character
// buffer it’s managing and access to individual characters in the buffer.
// Valid reasons for making a subclass of [NSString] include providing a
// different backing store (perhaps for better performance) or implementing
// some aspect of object behavior differently, such as memory management. If
// your purpose is to add non-essential attributes or metadata to your
// subclass of [NSString], a better alternative would be object composition
// (see [NSString]). Cocoa already provides an example of this with the
// [NSAttributedString] class.
//
// # Methods to Override
//
// Any subclass of [NSString] override the primitive instance methods [NSString.Length]
// and [NSString.CharacterAtIndex]. These methods must operate on the backing store
// that you provide for the characters of the string. For this backing store
// you can use a static array, a dynamically allocated buffer, a standard
// [NSString] object, or some other data type or mechanism. You may also
// choose to override, partially or fully, any other [NSString] method for
// which you want to provide an alternative implementation. For example, for
// better performance it is recommended that you override [NSString.GetCharactersRange]
// and give it a faster implementation.
//
// You might want to implement an initializer for your subclass that is suited
// to the backing store that the subclass is managing. The [NSString] class
// does not have a designated initializer, so your initializer need only
// invoke the [init()] method of `super`. The [NSString] class adopts the
// [NSCopying], [NSMutableCopying], and [NSCoding] protocols; if you want
// instances of your own custom subclass created from copying or coding,
// override the methods in these protocols.
//
// # Alternatives to Subclassing
//
// Often a better and easier alternative to making a subclass of
// [NSString]—or of any other abstract, public class of a class cluster, for
// that matter—is object composition. This is especially the case when your
// intent is to add to the subclass metadata or some other attribute that is
// not essential to a string object. In object composition, you would have an
// [NSString] object as one instance variable of your custom class (typically
// a subclass of [NSObject]) and one or more instance variables that store the
// metadata that you want for the custom object. Then just design your
// subclass interface to include accessor methods for the embedded string
// object and the metadata.
//
// If the behavior you want to add supplements that of the existing class, you
// could write a category on [NSString]. Keep in mind, however, that this
// category will be in effect for all instances of [NSString] that you use,
// and this might have unintended consequences.
//
// # Creating and Initializing Strings
//
//   - [NSString.InitWithBytesLengthEncoding]: Returns an initialized [NSString] object containing a given number of bytes from a given buffer of bytes interpreted in a given encoding.
//   - [NSString.InitWithBytesNoCopyLengthEncodingFreeWhenDone]: Returns an initialized [NSString] object that contains a given number of bytes from a given buffer of bytes interpreted in a given encoding, and optionally frees the buffer.
//   - [NSString.InitWithCharactersLength]: Returns an initialized [NSString] object that contains a given number of characters from a given C array of UTF-16 code units.
//   - [NSString.InitWithCharactersNoCopyLengthFreeWhenDone]: Returns an initialized [NSString] object that contains a given number of characters from a given C array of UTF-16 code units.
//   - [NSString.InitWithString]: Returns an [NSString] object initialized by copying the characters from another given string.
//   - [NSString.InitWithFormatArguments]: Returns an [NSString] object initialized by using a given format string as a template into which the remaining argument values are substituted without any localization.
//   - [NSString.InitWithFormatLocaleArguments]: Returns an [NSString] object initialized by using a given format string as a template into which the remaining argument values are substituted according to given locale information. This method is meant to be called from within a variadic function, where the argument list will be available.
//   - [NSString.InitWithDataEncoding]: Returns an [NSString] object initialized by converting given data into UTF-16 code units using a given encoding.
//
// # Creating and Initializing a String from a File
//
//   - [NSString.InitWithContentsOfFileEncodingError]: Returns an [NSString] object initialized by reading data from the file at a given path using a given encoding.
//   - [NSString.InitWithContentsOfFileUsedEncodingError]: Returns an [NSString] object initialized by reading data from the file at a given path and returns by reference the encoding used to interpret the characters.
//
// # Getting a String’s Length
//
//   - [NSString.Length]: The number of UTF-16 code units in the receiver.
//   - [NSString.LengthOfBytesUsingEncoding]: Returns the number of bytes required to store the receiver in a given encoding.
//   - [NSString.MaximumLengthOfBytesUsingEncoding]: Returns the maximum number of bytes needed to store the receiver in a given encoding.
//
// # Getting Characters and Bytes
//
//   - [NSString.CharacterAtIndex]: Returns the character at a given UTF-16 code unit index.
//   - [NSString.GetCharactersRange]: Copies characters from a given range in the receiver into a given buffer.
//   - [NSString.GetBytesMaxLengthUsedLengthEncodingOptionsRangeRemainingRange]: Gets a given range of characters as bytes in a specified encoding.
//
// # Getting C Strings
//
//   - [NSString.CStringUsingEncoding]: Returns a representation of the string as a C string using a given encoding.
//   - [NSString.GetCStringMaxLengthEncoding]: Converts the string to a given encoding and stores it in a buffer.
//   - [NSString.UTF8String]: A null-terminated UTF8 representation of the string.
//
// # Identifying and Comparing Strings
//
//   - [NSString.CaseInsensitiveCompare]: Returns the result of invoking [compare(_:options:)](<doc://com.apple.foundation/documentation/Foundation/NSString/compare(_:options:)>) with [NSCaseInsensitiveSearch] as the only option.
//   - [NSString.LocalizedCaseInsensitiveCompare]: Compares the string with a given string using a case-insensitive, localized, comparison.
//   - [NSString.Compare]: Returns the result of invoking [compare(_:options:range:)](<doc://com.apple.foundation/documentation/Foundation/NSString/compare(_:options:range:)>) with no options and the receiver’s full extent as the range.
//   - [NSString.LocalizedCompare]: Compares the string and a given string using a localized comparison.
//   - [NSString.CompareOptions]: Compares the string with the specified string using the given options.
//   - [NSString.CompareOptionsRange]: Returns the result of invoking [compare(_:options:range:locale:)](<doc://com.apple.foundation/documentation/Foundation/NSString/compare(_:options:range:locale:)>) with a `nil` locale.
//   - [NSString.CompareOptionsRangeLocale]: Compares the string using the specified options and returns the lexical ordering for the range.
//   - [NSString.LocalizedStandardCompare]: Compares strings as sorted by the Finder.
//   - [NSString.HasPrefix]: Returns a Boolean value that indicates whether a given string matches the beginning characters of the receiver.
//   - [NSString.HasSuffix]: Returns a Boolean value that indicates whether a given string matches the ending characters of the receiver.
//   - [NSString.IsEqualToString]: Returns a Boolean value that indicates whether a given string is equal to the receiver using a literal Unicode-based comparison.
//   - [NSString.Hash]: An unsigned integer that can be used as a hash table address.
//
// # Combining Strings
//
//   - [NSString.StringByAppendingString]: Returns a new string made by appending a given string to the receiver.
//   - [NSString.StringByPaddingToLengthWithStringStartingAtIndex]: Returns a new string formed from the receiver by either removing characters from the end, or by appending as many occurrences as necessary of a given pad string.
//
// # Changing Case
//
//   - [NSString.LowercaseString]: A lowercase representation of the string.
//   - [NSString.LocalizedLowercaseString]: Returns a version of the string with all letters converted to lowercase, taking into account the current locale.
//   - [NSString.LowercaseStringWithLocale]: Returns a version of the string with all letters converted to lowercase, taking into account the specified locale.
//   - [NSString.UppercaseString]: An uppercase representation of the string.
//   - [NSString.LocalizedUppercaseString]: Returns a version of the string with all letters converted to uppercase, taking into account the current locale.
//   - [NSString.UppercaseStringWithLocale]: Returns a version of the string with all letters converted to uppercase, taking into account the specified locale.
//   - [NSString.CapitalizedString]: A capitalized representation of the string.
//   - [NSString.LocalizedCapitalizedString]: Returns a capitalized representation of the receiver using the current locale.
//   - [NSString.CapitalizedStringWithLocale]: Returns a capitalized representation of the receiver using the specified locale.
//
// # Dividing Strings
//
//   - [NSString.ComponentsSeparatedByString]: Returns an array containing substrings from the receiver that have been divided by a given separator.
//   - [NSString.ComponentsSeparatedByCharactersInSet]: Returns an array containing substrings from the receiver that have been divided by characters in a given set.
//   - [NSString.StringByTrimmingCharactersInSet]: Returns a new string made by removing from both ends of the receiver characters contained in a given character set.
//   - [NSString.SubstringFromIndex]: Returns a new string containing the characters of the receiver from the one at a given index to the end.
//   - [NSString.SubstringWithRange]: Returns a string object containing the characters of the receiver that lie within a given range.
//   - [NSString.SubstringToIndex]: Returns a new string containing the characters of the receiver up to, but not including, the one at a given index.
//
// # Normalizing Strings
//
//   - [NSString.DecomposedStringWithCanonicalMapping]: A string made by normalizing the string’s contents using the Unicode Normalization Form D.
//   - [NSString.DecomposedStringWithCompatibilityMapping]: A string made by normalizing the receiver’s contents using the Unicode Normalization Form KD.
//   - [NSString.PrecomposedStringWithCanonicalMapping]: A string made by normalizing the string’s contents using the Unicode Normalization Form C.
//   - [NSString.PrecomposedStringWithCompatibilityMapping]: A string made by normalizing the receiver’s contents using the Unicode Normalization Form KC.
//
// # Folding Strings
//
//   - [NSString.StringByFoldingWithOptionsLocale]: Creates a string suitable for comparison by removing the specified character distinctions from a string.
//
// # Transforming Strings
//
//   - [NSString.StringByApplyingTransformReverse]: Returns a new string by applying a specified transform to the string.
//
// # Finding Characters and Substrings
//
//   - [NSString.ContainsString]: Returns a Boolean value indicating whether the string contains a given string by performing a case-sensitive, locale-unaware search.
//   - [NSString.LocalizedCaseInsensitiveContainsString]: Returns a Boolean value indicating whether the string contains a given string by performing a case-insensitive, locale-aware search.
//   - [NSString.LocalizedStandardContainsString]: Returns a Boolean value indicating whether the string contains a given string by performing a case and diacritic insensitive, locale-aware search.
//   - [NSString.RangeOfCharacterFromSet]: Finds and returns the range in the string of the first character from a given character set.
//   - [NSString.RangeOfCharacterFromSetOptions]: Finds and returns the range in the string of the first character, using given options, from a given character set.
//   - [NSString.RangeOfCharacterFromSetOptionsRange]: Finds and returns the range in the string of the first character from a given character set found in a given range with given options.
//   - [NSString.RangeOfString]: Finds and returns the range of the first occurrence of a given string within the string.
//   - [NSString.RangeOfStringOptions]: Finds and returns the range of the first occurrence of a given string within the string, subject to given options.
//   - [NSString.RangeOfStringOptionsRange]: Finds and returns the range of the first occurrence of a given string, within the given range of the string, subject to given options.
//   - [NSString.RangeOfStringOptionsRangeLocale]: Finds and returns the range of the first occurrence of a given string within a given range of the string, subject to given options, using the specified locale, if any.
//   - [NSString.LocalizedStandardRangeOfString]: Finds and returns the range of the first occurrence of a given string within the string by performing a case and diacritic insensitive, locale-aware search.
//
// # Replacing Substrings
//
//   - [NSString.StringByReplacingOccurrencesOfStringWithString]: Returns a new string in which all occurrences of a target string in the receiver are replaced by another given string.
//   - [NSString.StringByReplacingOccurrencesOfStringWithStringOptionsRange]: Returns a new string in which all occurrences of a target string in a specified range of the receiver are replaced by another given string.
//   - [NSString.StringByReplacingCharactersInRangeWithString]: Returns a new string in which the characters in a specified range of the receiver are replaced by a given string.
//
// # Getting a Shared Prefix
//
//   - [NSString.CommonPrefixWithStringOptions]: Returns a string containing characters the receiver and a given string have in common, starting from the beginning of each up to the first characters that aren’t equivalent.
//
// # Determining Line and Paragraph Ranges
//
//   - [NSString.GetLineStartEndContentsEndForRange]: Returns by reference the beginning of the first line and the end of the last line touched by the given range.
//   - [NSString.LineRangeForRange]: Returns the range of characters representing the line or lines containing a given range.
//   - [NSString.GetParagraphStartEndContentsEndForRange]: Returns by reference the beginning of the first paragraph and the end of the last paragraph touched by the given range.
//   - [NSString.ParagraphRangeForRange]: Returns the range of characters representing the paragraph or paragraphs containing a given range.
//
// # Determining Composed Character Sequences
//
//   - [NSString.RangeOfComposedCharacterSequenceAtIndex]: Returns the range in the receiver of the composed character sequence located at a given index.
//   - [NSString.RangeOfComposedCharacterSequencesForRange]: Returns the range in the string of the composed character sequences for a given range.
//
// # Writing to a File or URL
//
//   - [NSString.WriteToFileAtomicallyEncodingError]: Writes the contents of the receiver to a file at a given path using a given encoding.
//   - [NSString.WriteToURLAtomicallyEncodingError]: Writes the contents of the receiver to the URL specified by `url` using the specified encoding.
//
// # Converting String Contents Into a Property List
//
//   - [NSString.PropertyList]: Parses the receiver as a text representation of a property list, returning an [NSString], [NSData], [NSArray], or [NSDictionary] object, according to the topmost element.
//   - [NSString.PropertyListFromStringsFileFormat]: Returns a dictionary object initialized with the keys and values found in the receiver.
//
// # Sizing and Drawing Strings
//
//   - [NSString.DrawAtPointWithAttributes]: Draws the receiver with the font and other display characteristics of the given attributes, at the specified point in the current graphics context.
//   - [NSString.DrawInRectWithAttributes]: Draws the attributed string inside the specified bounding rectangle.
//   - [NSString.DrawWithRectOptionsAttributesContext]: Draws the attributed string in the specified bounding rectangle using the provided options.
//   - [NSString.BoundingRectWithSizeOptionsAttributesContext]: Calculates and returns the bounding rect for the receiver drawn using the given options and display characteristics, within the specified rectangle in the current graphics context.
//   - [NSString.SizeWithAttributes]: Returns the bounding box size the receiver occupies when drawn with the given attributes.
//   - [NSString.VariantFittingPresentationWidth]: Returns a string variation suitable for the specified presentation width.
//
// # Getting Numeric Values
//
//   - [NSString.DoubleValue]: The floating-point value of the string as a `double`.
//   - [NSString.FloatValue]: The floating-point value of the string as a `float`.
//   - [NSString.IntValue]: The integer value of the string.
//   - [NSString.IntegerValue]: The [NSInteger] value of the string.
//   - [NSString.LongLongValue]: The `long long` value of the string.
//   - [NSString.BoolValue]: The Boolean value of the string.
//
// # Working with Encodings
//
//   - [NSString.CanBeConvertedToEncoding]: Returns a Boolean value that indicates whether the receiver can be converted to a given encoding without loss of information.
//   - [NSString.DataUsingEncoding]: Returns an [NSData] object containing a representation of the receiver encoded using a given encoding.
//   - [NSString.DataUsingEncodingAllowLossyConversion]: Returns an [NSData] object containing a representation of the receiver encoded using a given encoding.
//   - [NSString.Description]
//   - [NSString.FastestEncoding]: The fastest encoding to which the receiver may be converted without loss of information.
//   - [NSString.SmallestEncoding]: The smallest encoding to which the receiver can be converted without loss of information.
//
// # Working with Paths
//
//   - [NSString.PathComponents]: The file-system path components of the receiver.
//   - [NSString.CompletePathIntoStringCaseSensitiveMatchesIntoArrayFilterTypes]: Interprets the receiver as a path in the file system and attempts to perform filename completion, returning a numeric value that indicates whether a match was possible, and by reference the longest path that matches the receiver.
//   - [NSString.FileSystemRepresentation]: A file system-specific representation of the receiver.
//   - [NSString.GetFileSystemRepresentationMaxLength]: Interprets the receiver as a system-independent path and fills a buffer with a C-string in a format and encoding suitable for use with file-system calls.
//   - [NSString.AbsolutePath]: A Boolean value that indicates whether the receiver represents an absolute path.
//   - [NSString.LastPathComponent]: The last path component of the receiver.
//   - [NSString.PathExtension]: The path extension, if any, of the string as interpreted as a path.
//   - [NSString.StringByAbbreviatingWithTildeInPath]: A new string that replaces the current home directory portion of the current path with a tilde (`~`) character.
//   - [NSString.StringByAppendingPathComponent]: Returns a new string made by appending to the receiver a given string.
//   - [NSString.StringByAppendingPathExtension]: Returns a new string made by appending to the receiver an extension separator followed by a given extension.
//   - [NSString.StringByDeletingLastPathComponent]: A new string made by deleting the last path component from the receiver, along with any final path separator.
//   - [NSString.StringByDeletingPathExtension]: A new string made by deleting the extension (if any, and only the last) from the receiver.
//   - [NSString.StringByExpandingTildeInPath]: A new string made by expanding the initial component of the receiver to its full path value.
//   - [NSString.StringByResolvingSymlinksInPath]: A new string made from the receiver by resolving all symbolic links and standardizing path.
//   - [NSString.StringByStandardizingPath]: A new string made by removing extraneous path components from the receiver.
//   - [NSString.StringsByAppendingPaths]: Returns an array of strings made by separately appending to the receiver each string in a given array.
//
// # Working with URL Strings
//
//   - [NSString.StringByAddingPercentEncodingWithAllowedCharacters]: Returns a new string made from the receiver by replacing all characters not in the specified set with percent-encoded characters.
//   - [NSString.StringByRemovingPercentEncoding]: Returns a new string made from the receiver by replacing all percent encoded sequences with the matching UTF-8 characters.
//
// # Deprecated
//
//   - [NSString.GetCharacters]: Copies all characters from the receiver into a given buffer.
//   - [NSString.DrawWithRectOptionsAttributes]: Draws the receiver with the specified options and other display characteristics of the given attributes, within the specified rectangle in the current graphics context.
//   - [NSString.BoundingRectWithSizeOptionsAttributes]: Calculates and returns the bounding rect for the receiver drawn using the given options and display characteristics, within the specified rectangle in the current graphics context.
//
// # Initializers
//
//   - [NSString.InitWithBytesNoCopyLengthEncodingDeallocator]
//   - [NSString.InitWithCStringEncoding]
//   - [NSString.InitWithContentsOfURLEncodingError]
//   - [NSString.InitWithContentsOfURLUsedEncodingError]
//   - [NSString.InitWithUTF8String]
//
// # Instance Properties
//
//   - [NSString.CustomPlaygroundQuickLook]: A custom playground Quick Look for this instance.
//   - [NSString.SetCustomPlaygroundQuickLook]
//
// # Instance Methods
//
//   - [NSString.StringByAppendingPathComponentConformingToType]
//   - [NSString.StringByAppendingPathExtensionForType]
//
// See: https://developer.apple.com/documentation/Foundation/NSString
//
// [CFString]: https://developer.apple.com/documentation/CoreFoundation/CFString
// [Characters and Grapheme Clusters]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/stringsClusters.html#//apple_ref/doc/uid/TP40008025
// [String Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/introStrings.html#//apple_ref/doc/uid/10000035i
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
// [http://www.unicode.org/]: http://www.unicode.org/
// [init()]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/init()
type NSString struct {
	objectivec.Object
}

// NSStringFromID constructs a [NSString] from an objc.ID.
//
// A static, plain-text Unicode string object.
func NSStringFromID(id objc.ID) NSString {
	return NSString{objectivec.Object{ID: id}}
}

// NOTE: NSString adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSString] class.
//
// # Creating and Initializing Strings
//
//   - [INSString.InitWithBytesLengthEncoding]: Returns an initialized [NSString] object containing a given number of bytes from a given buffer of bytes interpreted in a given encoding.
//   - [INSString.InitWithBytesNoCopyLengthEncodingFreeWhenDone]: Returns an initialized [NSString] object that contains a given number of bytes from a given buffer of bytes interpreted in a given encoding, and optionally frees the buffer.
//   - [INSString.InitWithCharactersLength]: Returns an initialized [NSString] object that contains a given number of characters from a given C array of UTF-16 code units.
//   - [INSString.InitWithCharactersNoCopyLengthFreeWhenDone]: Returns an initialized [NSString] object that contains a given number of characters from a given C array of UTF-16 code units.
//   - [INSString.InitWithString]: Returns an [NSString] object initialized by copying the characters from another given string.
//   - [INSString.InitWithFormatArguments]: Returns an [NSString] object initialized by using a given format string as a template into which the remaining argument values are substituted without any localization.
//   - [INSString.InitWithFormatLocaleArguments]: Returns an [NSString] object initialized by using a given format string as a template into which the remaining argument values are substituted according to given locale information. This method is meant to be called from within a variadic function, where the argument list will be available.
//   - [INSString.InitWithDataEncoding]: Returns an [NSString] object initialized by converting given data into UTF-16 code units using a given encoding.
//
// # Creating and Initializing a String from a File
//
//   - [INSString.InitWithContentsOfFileEncodingError]: Returns an [NSString] object initialized by reading data from the file at a given path using a given encoding.
//   - [INSString.InitWithContentsOfFileUsedEncodingError]: Returns an [NSString] object initialized by reading data from the file at a given path and returns by reference the encoding used to interpret the characters.
//
// # Getting a String’s Length
//
//   - [INSString.Length]: The number of UTF-16 code units in the receiver.
//   - [INSString.LengthOfBytesUsingEncoding]: Returns the number of bytes required to store the receiver in a given encoding.
//   - [INSString.MaximumLengthOfBytesUsingEncoding]: Returns the maximum number of bytes needed to store the receiver in a given encoding.
//
// # Getting Characters and Bytes
//
//   - [INSString.CharacterAtIndex]: Returns the character at a given UTF-16 code unit index.
//   - [INSString.GetCharactersRange]: Copies characters from a given range in the receiver into a given buffer.
//   - [INSString.GetBytesMaxLengthUsedLengthEncodingOptionsRangeRemainingRange]: Gets a given range of characters as bytes in a specified encoding.
//
// # Getting C Strings
//
//   - [INSString.CStringUsingEncoding]: Returns a representation of the string as a C string using a given encoding.
//   - [INSString.GetCStringMaxLengthEncoding]: Converts the string to a given encoding and stores it in a buffer.
//   - [INSString.UTF8String]: A null-terminated UTF8 representation of the string.
//
// # Identifying and Comparing Strings
//
//   - [INSString.CaseInsensitiveCompare]: Returns the result of invoking [compare(_:options:)](<doc://com.apple.foundation/documentation/Foundation/NSString/compare(_:options:)>) with [NSCaseInsensitiveSearch] as the only option.
//   - [INSString.LocalizedCaseInsensitiveCompare]: Compares the string with a given string using a case-insensitive, localized, comparison.
//   - [INSString.Compare]: Returns the result of invoking [compare(_:options:range:)](<doc://com.apple.foundation/documentation/Foundation/NSString/compare(_:options:range:)>) with no options and the receiver’s full extent as the range.
//   - [INSString.LocalizedCompare]: Compares the string and a given string using a localized comparison.
//   - [INSString.CompareOptions]: Compares the string with the specified string using the given options.
//   - [INSString.CompareOptionsRange]: Returns the result of invoking [compare(_:options:range:locale:)](<doc://com.apple.foundation/documentation/Foundation/NSString/compare(_:options:range:locale:)>) with a `nil` locale.
//   - [INSString.CompareOptionsRangeLocale]: Compares the string using the specified options and returns the lexical ordering for the range.
//   - [INSString.LocalizedStandardCompare]: Compares strings as sorted by the Finder.
//   - [INSString.HasPrefix]: Returns a Boolean value that indicates whether a given string matches the beginning characters of the receiver.
//   - [INSString.HasSuffix]: Returns a Boolean value that indicates whether a given string matches the ending characters of the receiver.
//   - [INSString.IsEqualToString]: Returns a Boolean value that indicates whether a given string is equal to the receiver using a literal Unicode-based comparison.
//   - [INSString.Hash]: An unsigned integer that can be used as a hash table address.
//
// # Combining Strings
//
//   - [INSString.StringByAppendingString]: Returns a new string made by appending a given string to the receiver.
//   - [INSString.StringByPaddingToLengthWithStringStartingAtIndex]: Returns a new string formed from the receiver by either removing characters from the end, or by appending as many occurrences as necessary of a given pad string.
//
// # Changing Case
//
//   - [INSString.LowercaseString]: A lowercase representation of the string.
//   - [INSString.LocalizedLowercaseString]: Returns a version of the string with all letters converted to lowercase, taking into account the current locale.
//   - [INSString.LowercaseStringWithLocale]: Returns a version of the string with all letters converted to lowercase, taking into account the specified locale.
//   - [INSString.UppercaseString]: An uppercase representation of the string.
//   - [INSString.LocalizedUppercaseString]: Returns a version of the string with all letters converted to uppercase, taking into account the current locale.
//   - [INSString.UppercaseStringWithLocale]: Returns a version of the string with all letters converted to uppercase, taking into account the specified locale.
//   - [INSString.CapitalizedString]: A capitalized representation of the string.
//   - [INSString.LocalizedCapitalizedString]: Returns a capitalized representation of the receiver using the current locale.
//   - [INSString.CapitalizedStringWithLocale]: Returns a capitalized representation of the receiver using the specified locale.
//
// # Dividing Strings
//
//   - [INSString.ComponentsSeparatedByString]: Returns an array containing substrings from the receiver that have been divided by a given separator.
//   - [INSString.ComponentsSeparatedByCharactersInSet]: Returns an array containing substrings from the receiver that have been divided by characters in a given set.
//   - [INSString.StringByTrimmingCharactersInSet]: Returns a new string made by removing from both ends of the receiver characters contained in a given character set.
//   - [INSString.SubstringFromIndex]: Returns a new string containing the characters of the receiver from the one at a given index to the end.
//   - [INSString.SubstringWithRange]: Returns a string object containing the characters of the receiver that lie within a given range.
//   - [INSString.SubstringToIndex]: Returns a new string containing the characters of the receiver up to, but not including, the one at a given index.
//
// # Normalizing Strings
//
//   - [INSString.DecomposedStringWithCanonicalMapping]: A string made by normalizing the string’s contents using the Unicode Normalization Form D.
//   - [INSString.DecomposedStringWithCompatibilityMapping]: A string made by normalizing the receiver’s contents using the Unicode Normalization Form KD.
//   - [INSString.PrecomposedStringWithCanonicalMapping]: A string made by normalizing the string’s contents using the Unicode Normalization Form C.
//   - [INSString.PrecomposedStringWithCompatibilityMapping]: A string made by normalizing the receiver’s contents using the Unicode Normalization Form KC.
//
// # Folding Strings
//
//   - [INSString.StringByFoldingWithOptionsLocale]: Creates a string suitable for comparison by removing the specified character distinctions from a string.
//
// # Transforming Strings
//
//   - [INSString.StringByApplyingTransformReverse]: Returns a new string by applying a specified transform to the string.
//
// # Finding Characters and Substrings
//
//   - [INSString.ContainsString]: Returns a Boolean value indicating whether the string contains a given string by performing a case-sensitive, locale-unaware search.
//   - [INSString.LocalizedCaseInsensitiveContainsString]: Returns a Boolean value indicating whether the string contains a given string by performing a case-insensitive, locale-aware search.
//   - [INSString.LocalizedStandardContainsString]: Returns a Boolean value indicating whether the string contains a given string by performing a case and diacritic insensitive, locale-aware search.
//   - [INSString.RangeOfCharacterFromSet]: Finds and returns the range in the string of the first character from a given character set.
//   - [INSString.RangeOfCharacterFromSetOptions]: Finds and returns the range in the string of the first character, using given options, from a given character set.
//   - [INSString.RangeOfCharacterFromSetOptionsRange]: Finds and returns the range in the string of the first character from a given character set found in a given range with given options.
//   - [INSString.RangeOfString]: Finds and returns the range of the first occurrence of a given string within the string.
//   - [INSString.RangeOfStringOptions]: Finds and returns the range of the first occurrence of a given string within the string, subject to given options.
//   - [INSString.RangeOfStringOptionsRange]: Finds and returns the range of the first occurrence of a given string, within the given range of the string, subject to given options.
//   - [INSString.RangeOfStringOptionsRangeLocale]: Finds and returns the range of the first occurrence of a given string within a given range of the string, subject to given options, using the specified locale, if any.
//   - [INSString.LocalizedStandardRangeOfString]: Finds and returns the range of the first occurrence of a given string within the string by performing a case and diacritic insensitive, locale-aware search.
//
// # Replacing Substrings
//
//   - [INSString.StringByReplacingOccurrencesOfStringWithString]: Returns a new string in which all occurrences of a target string in the receiver are replaced by another given string.
//   - [INSString.StringByReplacingOccurrencesOfStringWithStringOptionsRange]: Returns a new string in which all occurrences of a target string in a specified range of the receiver are replaced by another given string.
//   - [INSString.StringByReplacingCharactersInRangeWithString]: Returns a new string in which the characters in a specified range of the receiver are replaced by a given string.
//
// # Getting a Shared Prefix
//
//   - [INSString.CommonPrefixWithStringOptions]: Returns a string containing characters the receiver and a given string have in common, starting from the beginning of each up to the first characters that aren’t equivalent.
//
// # Determining Line and Paragraph Ranges
//
//   - [INSString.GetLineStartEndContentsEndForRange]: Returns by reference the beginning of the first line and the end of the last line touched by the given range.
//   - [INSString.LineRangeForRange]: Returns the range of characters representing the line or lines containing a given range.
//   - [INSString.GetParagraphStartEndContentsEndForRange]: Returns by reference the beginning of the first paragraph and the end of the last paragraph touched by the given range.
//   - [INSString.ParagraphRangeForRange]: Returns the range of characters representing the paragraph or paragraphs containing a given range.
//
// # Determining Composed Character Sequences
//
//   - [INSString.RangeOfComposedCharacterSequenceAtIndex]: Returns the range in the receiver of the composed character sequence located at a given index.
//   - [INSString.RangeOfComposedCharacterSequencesForRange]: Returns the range in the string of the composed character sequences for a given range.
//
// # Writing to a File or URL
//
//   - [INSString.WriteToFileAtomicallyEncodingError]: Writes the contents of the receiver to a file at a given path using a given encoding.
//   - [INSString.WriteToURLAtomicallyEncodingError]: Writes the contents of the receiver to the URL specified by `url` using the specified encoding.
//
// # Converting String Contents Into a Property List
//
//   - [INSString.PropertyList]: Parses the receiver as a text representation of a property list, returning an [NSString], [NSData], [NSArray], or [NSDictionary] object, according to the topmost element.
//   - [INSString.PropertyListFromStringsFileFormat]: Returns a dictionary object initialized with the keys and values found in the receiver.
//
// # Sizing and Drawing Strings
//
//   - [INSString.DrawAtPointWithAttributes]: Draws the receiver with the font and other display characteristics of the given attributes, at the specified point in the current graphics context.
//   - [INSString.DrawInRectWithAttributes]: Draws the attributed string inside the specified bounding rectangle.
//   - [INSString.DrawWithRectOptionsAttributesContext]: Draws the attributed string in the specified bounding rectangle using the provided options.
//   - [INSString.BoundingRectWithSizeOptionsAttributesContext]: Calculates and returns the bounding rect for the receiver drawn using the given options and display characteristics, within the specified rectangle in the current graphics context.
//   - [INSString.SizeWithAttributes]: Returns the bounding box size the receiver occupies when drawn with the given attributes.
//   - [INSString.VariantFittingPresentationWidth]: Returns a string variation suitable for the specified presentation width.
//
// # Getting Numeric Values
//
//   - [INSString.DoubleValue]: The floating-point value of the string as a `double`.
//   - [INSString.FloatValue]: The floating-point value of the string as a `float`.
//   - [INSString.IntValue]: The integer value of the string.
//   - [INSString.IntegerValue]: The [NSInteger] value of the string.
//   - [INSString.LongLongValue]: The `long long` value of the string.
//   - [INSString.BoolValue]: The Boolean value of the string.
//
// # Working with Encodings
//
//   - [INSString.CanBeConvertedToEncoding]: Returns a Boolean value that indicates whether the receiver can be converted to a given encoding without loss of information.
//   - [INSString.DataUsingEncoding]: Returns an [NSData] object containing a representation of the receiver encoded using a given encoding.
//   - [INSString.DataUsingEncodingAllowLossyConversion]: Returns an [NSData] object containing a representation of the receiver encoded using a given encoding.
//   - [INSString.Description]
//   - [INSString.FastestEncoding]: The fastest encoding to which the receiver may be converted without loss of information.
//   - [INSString.SmallestEncoding]: The smallest encoding to which the receiver can be converted without loss of information.
//
// # Working with Paths
//
//   - [INSString.PathComponents]: The file-system path components of the receiver.
//   - [INSString.CompletePathIntoStringCaseSensitiveMatchesIntoArrayFilterTypes]: Interprets the receiver as a path in the file system and attempts to perform filename completion, returning a numeric value that indicates whether a match was possible, and by reference the longest path that matches the receiver.
//   - [INSString.FileSystemRepresentation]: A file system-specific representation of the receiver.
//   - [INSString.GetFileSystemRepresentationMaxLength]: Interprets the receiver as a system-independent path and fills a buffer with a C-string in a format and encoding suitable for use with file-system calls.
//   - [INSString.AbsolutePath]: A Boolean value that indicates whether the receiver represents an absolute path.
//   - [INSString.LastPathComponent]: The last path component of the receiver.
//   - [INSString.PathExtension]: The path extension, if any, of the string as interpreted as a path.
//   - [INSString.StringByAbbreviatingWithTildeInPath]: A new string that replaces the current home directory portion of the current path with a tilde (`~`) character.
//   - [INSString.StringByAppendingPathComponent]: Returns a new string made by appending to the receiver a given string.
//   - [INSString.StringByAppendingPathExtension]: Returns a new string made by appending to the receiver an extension separator followed by a given extension.
//   - [INSString.StringByDeletingLastPathComponent]: A new string made by deleting the last path component from the receiver, along with any final path separator.
//   - [INSString.StringByDeletingPathExtension]: A new string made by deleting the extension (if any, and only the last) from the receiver.
//   - [INSString.StringByExpandingTildeInPath]: A new string made by expanding the initial component of the receiver to its full path value.
//   - [INSString.StringByResolvingSymlinksInPath]: A new string made from the receiver by resolving all symbolic links and standardizing path.
//   - [INSString.StringByStandardizingPath]: A new string made by removing extraneous path components from the receiver.
//   - [INSString.StringsByAppendingPaths]: Returns an array of strings made by separately appending to the receiver each string in a given array.
//
// # Working with URL Strings
//
//   - [INSString.StringByAddingPercentEncodingWithAllowedCharacters]: Returns a new string made from the receiver by replacing all characters not in the specified set with percent-encoded characters.
//   - [INSString.StringByRemovingPercentEncoding]: Returns a new string made from the receiver by replacing all percent encoded sequences with the matching UTF-8 characters.
//
// # Deprecated
//
//   - [INSString.GetCharacters]: Copies all characters from the receiver into a given buffer.
//   - [INSString.DrawWithRectOptionsAttributes]: Draws the receiver with the specified options and other display characteristics of the given attributes, within the specified rectangle in the current graphics context.
//   - [INSString.BoundingRectWithSizeOptionsAttributes]: Calculates and returns the bounding rect for the receiver drawn using the given options and display characteristics, within the specified rectangle in the current graphics context.
//
// # Initializers
//
//   - [INSString.InitWithBytesNoCopyLengthEncodingDeallocator]
//   - [INSString.InitWithCStringEncoding]
//   - [INSString.InitWithContentsOfURLEncodingError]
//   - [INSString.InitWithContentsOfURLUsedEncodingError]
//   - [INSString.InitWithUTF8String]
//
// # Instance Properties
//
//   - [INSString.CustomPlaygroundQuickLook]: A custom playground Quick Look for this instance.
//   - [INSString.SetCustomPlaygroundQuickLook]
//
// # Instance Methods
//
//   - [INSString.StringByAppendingPathComponentConformingToType]
//   - [INSString.StringByAppendingPathExtensionForType]
//
// See: https://developer.apple.com/documentation/Foundation/NSString
type INSString interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSItemProviderReading
	NSItemProviderWriting
	NSMutableCopying
	NSSecureCoding

	// Topic: Creating and Initializing Strings

	// Returns an initialized [NSString] object containing a given number of bytes from a given buffer of bytes interpreted in a given encoding.
	InitWithBytesLengthEncoding(bytes []byte, encoding uint) NSString
	// Returns an initialized [NSString] object that contains a given number of bytes from a given buffer of bytes interpreted in a given encoding, and optionally frees the buffer.
	InitWithBytesNoCopyLengthEncodingFreeWhenDone(bytes unsafe.Pointer, len_ uint, encoding uint, freeBuffer bool) NSString
	// Returns an initialized [NSString] object that contains a given number of characters from a given C array of UTF-16 code units.
	InitWithCharactersLength(characters unsafe.Pointer, length uint) NSString
	// Returns an initialized [NSString] object that contains a given number of characters from a given C array of UTF-16 code units.
	InitWithCharactersNoCopyLengthFreeWhenDone(characters unsafe.Pointer, length uint, freeBuffer bool) NSString
	// Returns an [NSString] object initialized by copying the characters from another given string.
	InitWithString(aString string) NSString
	// Returns an [NSString] object initialized by using a given format string as a template into which the remaining argument values are substituted without any localization.
	InitWithFormatArguments(format string, argList unsafe.Pointer) NSString
	// Returns an [NSString] object initialized by using a given format string as a template into which the remaining argument values are substituted according to given locale information. This method is meant to be called from within a variadic function, where the argument list will be available.
	InitWithFormatLocaleArguments(format string, locale objectivec.IObject, argList unsafe.Pointer) NSString
	// Returns an [NSString] object initialized by converting given data into UTF-16 code units using a given encoding.
	InitWithDataEncoding(data INSData, encoding uint) NSString

	// Topic: Creating and Initializing a String from a File

	// Returns an [NSString] object initialized by reading data from the file at a given path using a given encoding.
	InitWithContentsOfFileEncodingError(path string, enc uint) (NSString, error)
	// Returns an [NSString] object initialized by reading data from the file at a given path and returns by reference the encoding used to interpret the characters.
	InitWithContentsOfFileUsedEncodingError(path string, enc unsafe.Pointer) (NSString, error)

	// Topic: Getting a String’s Length

	// The number of UTF-16 code units in the receiver.
	Length() uint
	// Returns the number of bytes required to store the receiver in a given encoding.
	LengthOfBytesUsingEncoding(enc uint) uint
	// Returns the maximum number of bytes needed to store the receiver in a given encoding.
	MaximumLengthOfBytesUsingEncoding(enc uint) uint

	// Topic: Getting Characters and Bytes

	// Returns the character at a given UTF-16 code unit index.
	CharacterAtIndex(index uint) Unichar
	// Copies characters from a given range in the receiver into a given buffer.
	GetCharactersRange(buffer unsafe.Pointer, range_ NSRange)
	// Gets a given range of characters as bytes in a specified encoding.
	GetBytesMaxLengthUsedLengthEncodingOptionsRangeRemainingRange(buffer unsafe.Pointer, maxBufferCount uint, encoding NSStringEncoding, options NSStringEncodingConversionOptions, range_ NSRange, leftover NSRangePointer) (uint, bool)

	// Topic: Getting C Strings

	// Returns a representation of the string as a C string using a given encoding.
	CStringUsingEncoding(encoding uint) string
	// Converts the string to a given encoding and stores it in a buffer.
	GetCStringMaxLengthEncoding(buffer string, maxBufferCount uint, encoding uint) bool
	// A null-terminated UTF8 representation of the string.
	UTF8String() string

	// Topic: Identifying and Comparing Strings

	// Returns the result of invoking [compare(_:options:)](<doc://com.apple.foundation/documentation/Foundation/NSString/compare(_:options:)>) with [NSCaseInsensitiveSearch] as the only option.
	CaseInsensitiveCompare(string_ string) ComparisonResult
	// Compares the string with a given string using a case-insensitive, localized, comparison.
	LocalizedCaseInsensitiveCompare(string_ string) ComparisonResult
	// Returns the result of invoking [compare(_:options:range:)](<doc://com.apple.foundation/documentation/Foundation/NSString/compare(_:options:range:)>) with no options and the receiver’s full extent as the range.
	Compare(string_ string) ComparisonResult
	// Compares the string and a given string using a localized comparison.
	LocalizedCompare(string_ string) ComparisonResult
	// Compares the string with the specified string using the given options.
	CompareOptions(string_ string, mask NSStringCompareOptions) ComparisonResult
	// Returns the result of invoking [compare(_:options:range:locale:)](<doc://com.apple.foundation/documentation/Foundation/NSString/compare(_:options:range:locale:)>) with a `nil` locale.
	CompareOptionsRange(string_ string, mask NSStringCompareOptions, rangeOfReceiverToCompare NSRange) ComparisonResult
	// Compares the string using the specified options and returns the lexical ordering for the range.
	CompareOptionsRangeLocale(string_ string, mask NSStringCompareOptions, rangeOfReceiverToCompare NSRange, locale objectivec.IObject) ComparisonResult
	// Compares strings as sorted by the Finder.
	LocalizedStandardCompare(string_ string) ComparisonResult
	// Returns a Boolean value that indicates whether a given string matches the beginning characters of the receiver.
	HasPrefix(str string) bool
	// Returns a Boolean value that indicates whether a given string matches the ending characters of the receiver.
	HasSuffix(str string) bool
	// Returns a Boolean value that indicates whether a given string is equal to the receiver using a literal Unicode-based comparison.
	IsEqualToString(aString string) bool
	// An unsigned integer that can be used as a hash table address.
	Hash() uint

	// Topic: Combining Strings

	// Returns a new string made by appending a given string to the receiver.
	StringByAppendingString(aString string) string
	// Returns a new string formed from the receiver by either removing characters from the end, or by appending as many occurrences as necessary of a given pad string.
	StringByPaddingToLengthWithStringStartingAtIndex(newLength uint, padString string, padIndex uint) string

	// Topic: Changing Case

	// A lowercase representation of the string.
	LowercaseString() string
	// Returns a version of the string with all letters converted to lowercase, taking into account the current locale.
	LocalizedLowercaseString() string
	// Returns a version of the string with all letters converted to lowercase, taking into account the specified locale.
	LowercaseStringWithLocale(locale INSLocale) string
	// An uppercase representation of the string.
	UppercaseString() string
	// Returns a version of the string with all letters converted to uppercase, taking into account the current locale.
	LocalizedUppercaseString() string
	// Returns a version of the string with all letters converted to uppercase, taking into account the specified locale.
	UppercaseStringWithLocale(locale INSLocale) string
	// A capitalized representation of the string.
	CapitalizedString() string
	// Returns a capitalized representation of the receiver using the current locale.
	LocalizedCapitalizedString() string
	// Returns a capitalized representation of the receiver using the specified locale.
	CapitalizedStringWithLocale(locale INSLocale) string

	// Topic: Dividing Strings

	// Returns an array containing substrings from the receiver that have been divided by a given separator.
	ComponentsSeparatedByString(separator string) []string
	// Returns an array containing substrings from the receiver that have been divided by characters in a given set.
	ComponentsSeparatedByCharactersInSet(separator INSCharacterSet) []string
	// Returns a new string made by removing from both ends of the receiver characters contained in a given character set.
	StringByTrimmingCharactersInSet(set INSCharacterSet) string
	// Returns a new string containing the characters of the receiver from the one at a given index to the end.
	SubstringFromIndex(from uint) string
	// Returns a string object containing the characters of the receiver that lie within a given range.
	SubstringWithRange(range_ NSRange) string
	// Returns a new string containing the characters of the receiver up to, but not including, the one at a given index.
	SubstringToIndex(to uint) string

	// Topic: Normalizing Strings

	// A string made by normalizing the string’s contents using the Unicode Normalization Form D.
	DecomposedStringWithCanonicalMapping() string
	// A string made by normalizing the receiver’s contents using the Unicode Normalization Form KD.
	DecomposedStringWithCompatibilityMapping() string
	// A string made by normalizing the string’s contents using the Unicode Normalization Form C.
	PrecomposedStringWithCanonicalMapping() string
	// A string made by normalizing the receiver’s contents using the Unicode Normalization Form KC.
	PrecomposedStringWithCompatibilityMapping() string

	// Topic: Folding Strings

	// Creates a string suitable for comparison by removing the specified character distinctions from a string.
	StringByFoldingWithOptionsLocale(options NSStringCompareOptions, locale INSLocale) string

	// Topic: Transforming Strings

	// Returns a new string by applying a specified transform to the string.
	StringByApplyingTransformReverse(transform NSStringTransform, reverse bool) string

	// Topic: Finding Characters and Substrings

	// Returns a Boolean value indicating whether the string contains a given string by performing a case-sensitive, locale-unaware search.
	ContainsString(str string) bool
	// Returns a Boolean value indicating whether the string contains a given string by performing a case-insensitive, locale-aware search.
	LocalizedCaseInsensitiveContainsString(str string) bool
	// Returns a Boolean value indicating whether the string contains a given string by performing a case and diacritic insensitive, locale-aware search.
	LocalizedStandardContainsString(str string) bool
	// Finds and returns the range in the string of the first character from a given character set.
	RangeOfCharacterFromSet(searchSet INSCharacterSet) NSRange
	// Finds and returns the range in the string of the first character, using given options, from a given character set.
	RangeOfCharacterFromSetOptions(searchSet INSCharacterSet, mask NSStringCompareOptions) NSRange
	// Finds and returns the range in the string of the first character from a given character set found in a given range with given options.
	RangeOfCharacterFromSetOptionsRange(searchSet INSCharacterSet, mask NSStringCompareOptions, rangeOfReceiverToSearch NSRange) NSRange
	// Finds and returns the range of the first occurrence of a given string within the string.
	RangeOfString(searchString string) NSRange
	// Finds and returns the range of the first occurrence of a given string within the string, subject to given options.
	RangeOfStringOptions(searchString string, mask NSStringCompareOptions) NSRange
	// Finds and returns the range of the first occurrence of a given string, within the given range of the string, subject to given options.
	RangeOfStringOptionsRange(searchString string, mask NSStringCompareOptions, rangeOfReceiverToSearch NSRange) NSRange
	// Finds and returns the range of the first occurrence of a given string within a given range of the string, subject to given options, using the specified locale, if any.
	RangeOfStringOptionsRangeLocale(searchString string, mask NSStringCompareOptions, rangeOfReceiverToSearch NSRange, locale INSLocale) NSRange
	// Finds and returns the range of the first occurrence of a given string within the string by performing a case and diacritic insensitive, locale-aware search.
	LocalizedStandardRangeOfString(str string) NSRange

	// Topic: Replacing Substrings

	// Returns a new string in which all occurrences of a target string in the receiver are replaced by another given string.
	StringByReplacingOccurrencesOfStringWithString(target string, replacement string) string
	// Returns a new string in which all occurrences of a target string in a specified range of the receiver are replaced by another given string.
	StringByReplacingOccurrencesOfStringWithStringOptionsRange(target string, replacement string, options NSStringCompareOptions, searchRange NSRange) string
	// Returns a new string in which the characters in a specified range of the receiver are replaced by a given string.
	StringByReplacingCharactersInRangeWithString(range_ NSRange, replacement string) string

	// Topic: Getting a Shared Prefix

	// Returns a string containing characters the receiver and a given string have in common, starting from the beginning of each up to the first characters that aren’t equivalent.
	CommonPrefixWithStringOptions(str string, mask NSStringCompareOptions) string

	// Topic: Determining Line and Paragraph Ranges

	// Returns by reference the beginning of the first line and the end of the last line touched by the given range.
	GetLineStartEndContentsEndForRange(startPtr unsafe.Pointer, lineEndPtr unsafe.Pointer, contentsEndPtr unsafe.Pointer, range_ NSRange)
	// Returns the range of characters representing the line or lines containing a given range.
	LineRangeForRange(range_ NSRange) NSRange
	// Returns by reference the beginning of the first paragraph and the end of the last paragraph touched by the given range.
	GetParagraphStartEndContentsEndForRange(startPtr unsafe.Pointer, parEndPtr unsafe.Pointer, contentsEndPtr unsafe.Pointer, range_ NSRange)
	// Returns the range of characters representing the paragraph or paragraphs containing a given range.
	ParagraphRangeForRange(range_ NSRange) NSRange

	// Topic: Determining Composed Character Sequences

	// Returns the range in the receiver of the composed character sequence located at a given index.
	RangeOfComposedCharacterSequenceAtIndex(index uint) NSRange
	// Returns the range in the string of the composed character sequences for a given range.
	RangeOfComposedCharacterSequencesForRange(range_ NSRange) NSRange

	// Topic: Writing to a File or URL

	// Writes the contents of the receiver to a file at a given path using a given encoding.
	WriteToFileAtomicallyEncodingError(path string, useAuxiliaryFile bool, enc uint) (bool, error)
	// Writes the contents of the receiver to the URL specified by `url` using the specified encoding.
	WriteToURLAtomicallyEncodingError(url INSURL, useAuxiliaryFile bool, enc uint) (bool, error)

	// Topic: Converting String Contents Into a Property List

	// Parses the receiver as a text representation of a property list, returning an [NSString], [NSData], [NSArray], or [NSDictionary] object, according to the topmost element.
	PropertyList() objectivec.IObject
	// Returns a dictionary object initialized with the keys and values found in the receiver.
	PropertyListFromStringsFileFormat() INSDictionary

	// Topic: Sizing and Drawing Strings

	// Draws the receiver with the font and other display characteristics of the given attributes, at the specified point in the current graphics context.
	DrawAtPointWithAttributes(point corefoundation.CGPoint, attrs INSDictionary)
	// Draws the attributed string inside the specified bounding rectangle.
	DrawInRectWithAttributes(rect corefoundation.CGRect, attrs INSDictionary)
	// Draws the attributed string in the specified bounding rectangle using the provided options.
	DrawWithRectOptionsAttributesContext(rect corefoundation.CGRect, options NSStringDrawingOptions, attributes INSDictionary, context objectivec.IObject)
	// Calculates and returns the bounding rect for the receiver drawn using the given options and display characteristics, within the specified rectangle in the current graphics context.
	BoundingRectWithSizeOptionsAttributesContext(size corefoundation.CGSize, options NSStringDrawingOptions, attributes INSDictionary, context objectivec.IObject) corefoundation.CGRect
	// Returns the bounding box size the receiver occupies when drawn with the given attributes.
	SizeWithAttributes(attrs INSDictionary) corefoundation.CGSize
	// Returns a string variation suitable for the specified presentation width.
	VariantFittingPresentationWidth(width int) string

	// Topic: Getting Numeric Values

	// The floating-point value of the string as a `double`.
	DoubleValue() float64
	// The floating-point value of the string as a `float`.
	FloatValue() float32
	// The integer value of the string.
	IntValue() int
	// The [NSInteger] value of the string.
	IntegerValue() int
	// The `long long` value of the string.
	LongLongValue() int64
	// The Boolean value of the string.
	BoolValue() bool

	// Topic: Working with Encodings

	// Returns a Boolean value that indicates whether the receiver can be converted to a given encoding without loss of information.
	CanBeConvertedToEncoding(encoding uint) bool
	// Returns an [NSData] object containing a representation of the receiver encoded using a given encoding.
	DataUsingEncoding(encoding uint) INSData
	// Returns an [NSData] object containing a representation of the receiver encoded using a given encoding.
	DataUsingEncodingAllowLossyConversion(encoding uint, lossy bool) INSData
	Description() string
	// The fastest encoding to which the receiver may be converted without loss of information.
	FastestEncoding() NSStringEncoding
	// The smallest encoding to which the receiver can be converted without loss of information.
	SmallestEncoding() NSStringEncoding

	// Topic: Working with Paths

	// The file-system path components of the receiver.
	PathComponents() []string
	// Interprets the receiver as a path in the file system and attempts to perform filename completion, returning a numeric value that indicates whether a match was possible, and by reference the longest path that matches the receiver.
	CompletePathIntoStringCaseSensitiveMatchesIntoArrayFilterTypes(outputName string, flag bool, outputArray []string, filterTypes []string) uint
	// A file system-specific representation of the receiver.
	FileSystemRepresentation() string
	// Interprets the receiver as a system-independent path and fills a buffer with a C-string in a format and encoding suitable for use with file-system calls.
	GetFileSystemRepresentationMaxLength(cname string, max uint) bool
	// A Boolean value that indicates whether the receiver represents an absolute path.
	AbsolutePath() bool
	// The last path component of the receiver.
	LastPathComponent() string
	// The path extension, if any, of the string as interpreted as a path.
	PathExtension() string
	// A new string that replaces the current home directory portion of the current path with a tilde (`~`) character.
	StringByAbbreviatingWithTildeInPath() string
	// Returns a new string made by appending to the receiver a given string.
	StringByAppendingPathComponent(str string) string
	// Returns a new string made by appending to the receiver an extension separator followed by a given extension.
	StringByAppendingPathExtension(str string) string
	// A new string made by deleting the last path component from the receiver, along with any final path separator.
	StringByDeletingLastPathComponent() string
	// A new string made by deleting the extension (if any, and only the last) from the receiver.
	StringByDeletingPathExtension() string
	// A new string made by expanding the initial component of the receiver to its full path value.
	StringByExpandingTildeInPath() string
	// A new string made from the receiver by resolving all symbolic links and standardizing path.
	StringByResolvingSymlinksInPath() string
	// A new string made by removing extraneous path components from the receiver.
	StringByStandardizingPath() string
	// Returns an array of strings made by separately appending to the receiver each string in a given array.
	StringsByAppendingPaths(paths []string) []string

	// Topic: Working with URL Strings

	// Returns a new string made from the receiver by replacing all characters not in the specified set with percent-encoded characters.
	StringByAddingPercentEncodingWithAllowedCharacters(allowedCharacters INSCharacterSet) string
	// Returns a new string made from the receiver by replacing all percent encoded sequences with the matching UTF-8 characters.
	StringByRemovingPercentEncoding() string

	// Topic: Deprecated

	// Copies all characters from the receiver into a given buffer.
	GetCharacters(buffer unsafe.Pointer)
	// Draws the receiver with the specified options and other display characteristics of the given attributes, within the specified rectangle in the current graphics context.
	DrawWithRectOptionsAttributes(rect corefoundation.CGRect, options NSStringDrawingOptions, attributes INSDictionary)
	// Calculates and returns the bounding rect for the receiver drawn using the given options and display characteristics, within the specified rectangle in the current graphics context.
	BoundingRectWithSizeOptionsAttributes(size corefoundation.CGSize, options NSStringDrawingOptions, attributes INSDictionary) NSRect

	// Topic: Initializers

	InitWithBytesNoCopyLengthEncodingDeallocator(bytes unsafe.Pointer, len_ uint, encoding uint, deallocator func(unsafe.Pointer, uint64)) NSString
	InitWithCStringEncoding(nullTerminatedCString string, encoding uint) NSString
	InitWithContentsOfURLEncodingError(url INSURL, enc uint) (NSString, error)
	InitWithContentsOfURLUsedEncodingError(url INSURL, enc unsafe.Pointer) (NSString, error)
	InitWithUTF8String(nullTerminatedCString string) NSString

	// Topic: Instance Properties

	// A custom playground Quick Look for this instance.
	CustomPlaygroundQuickLook() objectivec.IObject
	SetCustomPlaygroundQuickLook(value objectivec.IObject)

	// Topic: Instance Methods

	StringByAppendingPathComponentConformingToType(partialName string, contentType objectivec.IObject) string
	StringByAppendingPathExtensionForType(contentType objectivec.IObject) string

	// Returns an [NSString] object initialized by using a given format string as a template into which the remaining argument values are substituted.
	InitWithFormat(format string) NSString
	// Returns an [NSString] object initialized by using a given format string as a template into which the remaining argument values are substituted according to given locale.
	InitWithFormatLocale(format string, locale objectivec.IObject) NSString
	InitWithValidatedFormatValidFormatSpecifiersArgumentsError(format string, validFormatSpecifiers string, argList unsafe.Pointer) (NSString, error)
	InitWithValidatedFormatValidFormatSpecifiersError(format string, validFormatSpecifiers string) (NSString, error)
	InitWithValidatedFormatValidFormatSpecifiersLocaleArgumentsError(format string, validFormatSpecifiers string, locale objectivec.IObject, argList unsafe.Pointer) (NSString, error)
	InitWithValidatedFormatValidFormatSpecifiersLocaleError(format string, validFormatSpecifiers string, locale objectivec.IObject) (NSString, error)
	// Returns a string made by appending to the receiver a string constructed from a given format string and the following arguments.
	StringByAppendingFormat(format string) string
}

// Init initializes the instance.
func (s NSString) Init() NSString {
	rv := objc.Send[NSString](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSString) Autorelease() NSString {
	rv := objc.Send[NSString](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSString creates a new NSString instance.
func NewNSString() NSString {
	class := getNSStringClass()
	rv := objc.Send[NSString](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an initialized [NSString] object containing a given number of bytes
// from a given buffer of bytes interpreted in a given encoding.
//
// bytes: A buffer of bytes interpreted in the encoding specified by `encoding`.
//
// len: The number of bytes to use from `bytes`.
//
// encoding: The character encoding applied to `bytes`. For possible values, see
// [NSStringEncoding].
//
// # Return Value
//
// An initialized [NSString] object containing `length` bytes from `bytes`
// interpreted using the encoding `encoding`. The returned object may be
// different from the original receiver. The return byte strings are allowed
// to be unterminated. If the length of the byte string is greater than the
// specified length a `nil` value is returned.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(bytes:length:encoding:)
func NewStringWithBytesLengthEncoding(bytes []byte, encoding uint) NSString {
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytes:length:encoding:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)), encoding)
	return NSStringFromID(rv)
}

// Returns an initialized [NSString] object that contains a given number of
// bytes from a given buffer of bytes interpreted in a given encoding, and
// optionally frees the buffer.
//
// bytes: A buffer of bytes interpreted in the encoding specified by `encoding`.
//
// len: The number of bytes to use from `bytes`.
//
// encoding: The character encoding of `bytes`. For possible values, see
// [NSStringEncoding].
//
// freeBuffer: If true, the receiver releases the memory with `free()` when it no longer
// needs the data; if false it won’t.
//
// # Return Value
//
// An initialized [NSString] object containing `length` bytes from `bytes`
// interpreted using the encoding `encoding`. The returned object may be
// different from the original receiver.
//
// # Discussion
//
// If an error occurs during the creation of the string, then `bytes` isn’t
// freed even if `flag` is true. In this case, the caller is responsible for
// freeing the buffer. This allows the caller to continue trying to create a
// string with the buffer, without having the buffer deallocated.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(bytesNoCopy:length:encoding:freeWhenDone:)
func NewStringWithBytesNoCopyLengthEncodingFreeWhenDone(bytes unsafe.Pointer, len_ uint, encoding uint, freeBuffer bool) NSString {
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytesNoCopy:length:encoding:freeWhenDone:"), bytes, len_, encoding, freeBuffer)
	return NSStringFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSString/init(cString:)
func NewStringWithCString(bytes string) NSString {
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCString:"), unsafe.Pointer(unsafe.StringData(bytes+"\x00")))
	return NSStringFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSString/init(cString:encoding:)
func NewStringWithCStringEncoding(nullTerminatedCString string, encoding uint) NSString {
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCString:encoding:"), unsafe.Pointer(unsafe.StringData(nullTerminatedCString+"\x00")), encoding)
	return NSStringFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSString/init(cString:length:)
func NewStringWithCStringLength(bytes string, length uint) NSString {
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCString:length:"), unsafe.Pointer(unsafe.StringData(bytes+"\x00")), length)
	return NSStringFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSString/init(cStringNoCopy:length:freeWhenDone:)
func NewStringWithCStringNoCopyLengthFreeWhenDone(bytes string, length uint, freeBuffer bool) NSString {
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCStringNoCopy:length:freeWhenDone:"), unsafe.Pointer(unsafe.StringData(bytes+"\x00")), length, freeBuffer)
	return NSStringFromID(rv)
}

// Returns an initialized [NSString] object that contains a given number of
// characters from a given C array of UTF-16 code units.
//
// characters: A C array of UTF-16 code units; the value must not be [NULL].
//
// length: The number of characters to use from `characters`.
//
// # Return Value
//
// An initialized [NSString] object containing `length` characters taken from
// `characters`. The returned object may be different from the original
// receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(characters:length:)
func NewStringWithCharactersLength(characters unsafe.Pointer, length uint) NSString {
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCharacters:length:"), characters, length)
	return NSStringFromID(rv)
}

// Returns an initialized [NSString] object that contains a given number of
// characters from a given C array of UTF-16 code units.
//
// characters: A C array of UTF-16 code units.
//
// length: The number of characters to use from `characters`.
//
// freeBuffer: If true, the receiver releases the memory with `free()` when it no longer
// needs the data; if false it won’t.
//
// # Return Value
//
// An initialized [NSString] object that contains `length` characters from
// `characters`. The returned object may be different from the original
// receiver.
//
// # Discussion
//
// If an error occurs during the creation of the string, then `bytes` is not
// freed even if `flag` is true. In this case, the caller is responsible for
// freeing the buffer. This allows the caller to continue trying to create a
// string with the buffer, without having the buffer deallocated.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(charactersNoCopy:length:freeWhenDone:)
func NewStringWithCharactersNoCopyLengthFreeWhenDone(characters unsafe.Pointer, length uint, freeBuffer bool) NSString {
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCharactersNoCopy:length:freeWhenDone:"), characters, length, freeBuffer)
	return NSStringFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSString/init(coder:)
func NewStringWithCoder(coder INSCoder) NSString {
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSStringFromID(rv)
}

// Initializes the receiver, a newly allocated [NSString] object, by reading
// data from the file named by `path`.
//
// # Discussion
//
// Initializes the receiver, a newly allocated [NSString] object, by reading
// data from the file named by `path`. If the contents begin with a byte-order
// mark (`U+FEFF` or `U+FFFE`), interprets the contents as UTF-16 code units;
// otherwise interprets the contents as data in the default C string encoding.
// Returns an initialized object, which might be different from the original
// receiver, or `nil` if the file can’t be opened.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfFile:)
func NewStringWithContentsOfFile(path string) NSString {
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:"), objc.String(path))
	return NSStringFromID(rv)
}

// Returns an [NSString] object initialized by reading data from the file at a
// given path using a given encoding.
//
// path: A path to a file.
//
// enc: The encoding of the file at `path`. For possible values, see
// [NSStringEncoding].
//
// # Return Value
//
// An [NSString] object initialized by reading data from the file named by
// `path` using the encoding, `enc`. The returned object may be different from
// the original receiver. If the file can’t be opened or there is an
// encoding error, returns `nil`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfFile:encoding:)
func NewStringWithContentsOfFileEncodingError(path string, enc uint) (NSString, error) {
	var errorPtr objc.ID
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:encoding:error:"), objc.String(path), enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil
}

// Returns an [NSString] object initialized by reading data from the file at a
// given path and returns by reference the encoding used to interpret the
// characters.
//
// path: A path to a file.
//
// enc: Upon return, if the file is read successfully, contains the encoding used
// to interpret the file at `path`. For possible values, see
// [NSStringEncoding].
//
// # Return Value
//
// An [NSString] object initialized by reading data from the file named by
// `path`. The returned object may be different from the original receiver. If
// the file can’t be opened or there is an encoding error, returns `nil`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfFile:usedEncoding:)
func NewStringWithContentsOfFileUsedEncodingError(path string, enc unsafe.Pointer) (NSString, error) {
	var errorPtr objc.ID
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:usedEncoding:error:"), objc.String(path), enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOf:)
func NewStringWithContentsOfURL(url INSURL) NSString {
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	return NSStringFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOf:encoding:)
func NewStringWithContentsOfURLEncodingError(url INSURL, enc uint) (NSString, error) {
	var errorPtr objc.ID
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:encoding:error:"), url, enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOf:usedEncoding:)
func NewStringWithContentsOfURLUsedEncodingError(url INSURL, enc unsafe.Pointer) (NSString, error) {
	var errorPtr objc.ID
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:usedEncoding:error:"), url, enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil
}

// Returns an [NSString] object initialized by converting given data into
// UTF-16 code units using a given encoding.
//
// data: An [NSData] object containing bytes in `encoding` and the default plain
// text format (that is, pure content with no attributes or other markups) for
// that encoding.
//
// encoding: The encoding used by `data`. For possible values, see [NSStringEncoding].
//
// # Return Value
//
// An [NSString] object initialized by converting the bytes in `data` into
// UTF-16 code units using `encoding`. The returned object may be different
// from the original receiver. Returns `nil` if the initialization fails for
// some reason (for example if `data` does not represent valid data for
// `encoding`).
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(data:encoding:)
func NewStringWithDataEncoding(data INSData, encoding uint) NSString {
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:encoding:"), data, encoding)
	return NSStringFromID(rv)
}

// Returns an [NSString] object initialized by using a given format string as
// a template into which the remaining argument values are substituted.
//
// format: A format string. See [Formatting String Objects] for examples of how to use
// this method, and [String Format Specifiers] for a list of format
// specifiers. This value must not be `nil`.
//
// # Return Value
//
// An [NSString] object initialized by using `format` as a template into which
// the remaining argument values are substituted according to the system
// locale. The returned object may be different from the original receiver.
//
// # Discussion
//
// Pass a comma-separated list of variadic arguments to substitute into
// `format`.
//
// This method invokes [InitWithFormatLocaleArguments] without applying any
// localization. This is useful, for example, when working with fixed-format
// representations of information that is written out and read back in at a
// later time.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/initWithFormat:
//
// [Formatting String Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/FormatStrings.html#//apple_ref/doc/uid/20000943
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
func NewStringWithFormat(format string) NSString {
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:"), objc.String(format))
	return NSStringFromID(rv)
}

// Returns an [NSString] object initialized by using a given format string as
// a template into which the remaining argument values are substituted without
// any localization.
//
// format: A format string. See [Formatting String Objects] for examples of how to use
// this method, and [String Format Specifiers] for a list of format
// specifiers. This value must not be `nil`.
//
// argList: A list of arguments to substitute into `format`.
//
// # Return Value
//
// An [NSString] object initialized by using `format` as a template into which
// the values in `argList` are substituted according to the current locale.
// The returned object may be different from the original receiver.
//
// # Discussion
//
// This method is meant to be called from within a variadic function, where
// the argument list will be available.
//
// This method invokes [InitWithFormatLocaleArguments] without applying any
// localization. This is useful, for example, when working with fixed-format
// representations of information that is written out and read back in at a
// later time.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(format:arguments:)
//
// [Formatting String Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/FormatStrings.html#//apple_ref/doc/uid/20000943
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
func NewStringWithFormatArguments(format string, argList unsafe.Pointer) NSString {
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:arguments:"), objc.String(format), argList)
	return NSStringFromID(rv)
}

// Returns an [NSString] object initialized by using a given format string as
// a template into which the remaining argument values are substituted
// according to given locale.
//
// format: A format string. See [Formatting String Objects] for examples of how to use
// this method, and [String Format Specifiers] for a list of format
// specifiers. This value must not be `nil`.
//
// locale: An [NSLocale] object specifying the locale to use. To use the current
// locale, pass `[NSLocale currentLocale]`. To use the system locale, pass
// `nil`.
//
// For legacy support, this may be an instance of [NSDictionary] containing
// locale information.
//
// # Discussion
//
// Pass comma-separated list of trailing variadic arguments to substitute into
// `format`.
//
// Invokes [InitWithFormatLocaleArguments] with `locale` as the locale.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/initWithFormat:locale:
//
// [Formatting String Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/FormatStrings.html#//apple_ref/doc/uid/20000943
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
func NewStringWithFormatLocale(format string, locale objectivec.IObject) NSString {
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:locale:"), objc.String(format), locale)
	return NSStringFromID(rv)
}

// Returns an [NSString] object initialized by using a given format string as
// a template into which the remaining argument values are substituted
// according to given locale information. This method is meant to be called
// from within a variadic function, where the argument list will be available.
//
// format: A format string. See [Formatting String Objects] for examples of how to use
// this method, and [String Format Specifiers] for a list of format
// specifiers. This value must not be `nil`.
//
// locale: An [NSLocale] object specifying the locale to use. To use the current
// locale (specified by user preferences), pass [NSLocale] [CurrentLocale]].
// To use the system locale, pass `nil`.
//
// For legacy support, this may be an instance of [NSDictionary] containing
// locale information.
//
// argList: A list of arguments to substitute into `format`.
//
// # Return Value
//
// An [NSString] object initialized by using `format` as a template into which
// values in `argList` are substituted according the locale information in
// `locale`. The returned object may be different from the original receiver.
//
// # Discussion
//
// The following Objective-C code fragment illustrates how to create a string
// from `myArgs`, which is derived from a string object with the value
// “Cost:” and an int with the value 32:
//
// The resulting string has the value “`Cost: 32\n`”.
//
// See [String Programming Guide] for more information.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(format:locale:arguments:)
//
// [Formatting String Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/FormatStrings.html#//apple_ref/doc/uid/20000943
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
// [String Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/introStrings.html#//apple_ref/doc/uid/10000035i
func NewStringWithFormatLocaleArguments(format string, locale objectivec.IObject, argList unsafe.Pointer) NSString {
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:locale:arguments:"), objc.String(format), locale, argList)
	return NSStringFromID(rv)
}

// Returns an [NSString] object initialized by copying the characters from
// another given string.
//
// aString: The string from which to copy characters. This value must not be `nil`.
//
// # Return Value
//
// An [NSString] object initialized by copying the characters from `aString`.
// The returned object may be different from the original receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(string:)-210xa
func NewStringWithString(aString string) NSString {
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:"), objc.String(aString))
	return NSStringFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSString/init(utf8String:)
func NewStringWithUTF8String(nullTerminatedCString string) NSString {
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUTF8String:"), unsafe.Pointer(unsafe.StringData(nullTerminatedCString+"\x00")))
	return NSStringFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:arguments:error:
func NewStringWithValidatedFormatValidFormatSpecifiersArgumentsError(format string, validFormatSpecifiers string, argList unsafe.Pointer) (NSString, error) {
	var errorPtr objc.ID
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:arguments:error:"), objc.String(format), objc.String(validFormatSpecifiers), argList, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:error:
func NewStringWithValidatedFormatValidFormatSpecifiersError(format string, validFormatSpecifiers string) (NSString, error) {
	var errorPtr objc.ID
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:error:"), objc.String(format), objc.String(validFormatSpecifiers), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:locale:arguments:error:
func NewStringWithValidatedFormatValidFormatSpecifiersLocaleArgumentsError(format string, validFormatSpecifiers string, locale objectivec.IObject, argList unsafe.Pointer) (NSString, error) {
	var errorPtr objc.ID
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:locale:arguments:error:"), objc.String(format), objc.String(validFormatSpecifiers), locale, argList, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:locale:error:
func NewStringWithValidatedFormatValidFormatSpecifiersLocaleError(format string, validFormatSpecifiers string, locale objectivec.IObject) (NSString, error) {
	var errorPtr objc.ID
	instance := getNSStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:locale:error:"), objc.String(format), objc.String(validFormatSpecifiers), locale, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil
}

// Returns an initialized [NSString] object containing a given number of bytes
// from a given buffer of bytes interpreted in a given encoding.
//
// bytes: A buffer of bytes interpreted in the encoding specified by `encoding`.
//
// len: The number of bytes to use from `bytes`.
//
// encoding: The character encoding applied to `bytes`. For possible values, see
// [NSStringEncoding].
//
// # Return Value
//
// An initialized [NSString] object containing `length` bytes from `bytes`
// interpreted using the encoding `encoding`. The returned object may be
// different from the original receiver. The return byte strings are allowed
// to be unterminated. If the length of the byte string is greater than the
// specified length a `nil` value is returned.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(bytes:length:encoding:)
func (s NSString) InitWithBytesLengthEncoding(bytes []byte, encoding uint) NSString {
	rv := objc.Send[NSString](s.ID, objc.Sel("initWithBytes:length:encoding:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)), encoding)
	return rv
}

// Returns an initialized [NSString] object that contains a given number of
// bytes from a given buffer of bytes interpreted in a given encoding, and
// optionally frees the buffer.
//
// bytes: A buffer of bytes interpreted in the encoding specified by `encoding`.
//
// len: The number of bytes to use from `bytes`.
//
// encoding: The character encoding of `bytes`. For possible values, see
// [NSStringEncoding].
//
// freeBuffer: If true, the receiver releases the memory with `free()` when it no longer
// needs the data; if false it won’t.
//
// # Return Value
//
// An initialized [NSString] object containing `length` bytes from `bytes`
// interpreted using the encoding `encoding`. The returned object may be
// different from the original receiver.
//
// # Discussion
//
// If an error occurs during the creation of the string, then `bytes` isn’t
// freed even if `flag` is true. In this case, the caller is responsible for
// freeing the buffer. This allows the caller to continue trying to create a
// string with the buffer, without having the buffer deallocated.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(bytesNoCopy:length:encoding:freeWhenDone:)
func (s NSString) InitWithBytesNoCopyLengthEncodingFreeWhenDone(bytes unsafe.Pointer, len_ uint, encoding uint, freeBuffer bool) NSString {
	rv := objc.Send[NSString](s.ID, objc.Sel("initWithBytesNoCopy:length:encoding:freeWhenDone:"), bytes, len_, encoding, freeBuffer)
	return rv
}

// Returns an initialized [NSString] object that contains a given number of
// characters from a given C array of UTF-16 code units.
//
// characters: A C array of UTF-16 code units; the value must not be [NULL].
//
// length: The number of characters to use from `characters`.
//
// # Return Value
//
// An initialized [NSString] object containing `length` characters taken from
// `characters`. The returned object may be different from the original
// receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(characters:length:)
func (s NSString) InitWithCharactersLength(characters unsafe.Pointer, length uint) NSString {
	rv := objc.Send[NSString](s.ID, objc.Sel("initWithCharacters:length:"), characters, length)
	return rv
}

// Returns an initialized [NSString] object that contains a given number of
// characters from a given C array of UTF-16 code units.
//
// characters: A C array of UTF-16 code units.
//
// length: The number of characters to use from `characters`.
//
// freeBuffer: If true, the receiver releases the memory with `free()` when it no longer
// needs the data; if false it won’t.
//
// # Return Value
//
// An initialized [NSString] object that contains `length` characters from
// `characters`. The returned object may be different from the original
// receiver.
//
// # Discussion
//
// If an error occurs during the creation of the string, then `bytes` is not
// freed even if `flag` is true. In this case, the caller is responsible for
// freeing the buffer. This allows the caller to continue trying to create a
// string with the buffer, without having the buffer deallocated.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(charactersNoCopy:length:freeWhenDone:)
func (s NSString) InitWithCharactersNoCopyLengthFreeWhenDone(characters unsafe.Pointer, length uint, freeBuffer bool) NSString {
	rv := objc.Send[NSString](s.ID, objc.Sel("initWithCharactersNoCopy:length:freeWhenDone:"), characters, length, freeBuffer)
	return rv
}

// Returns an [NSString] object initialized by copying the characters from
// another given string.
//
// aString: The string from which to copy characters. This value must not be `nil`.
//
// # Return Value
//
// An [NSString] object initialized by copying the characters from `aString`.
// The returned object may be different from the original receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(string:)-210xa
func (s NSString) InitWithString(aString string) NSString {
	rv := objc.Send[NSString](s.ID, objc.Sel("initWithString:"), objc.String(aString))
	return rv
}

// Returns an [NSString] object initialized by using a given format string as
// a template into which the remaining argument values are substituted without
// any localization.
//
// format: A format string. See [Formatting String Objects] for examples of how to use
// this method, and [String Format Specifiers] for a list of format
// specifiers. This value must not be `nil`.
//
// argList: A list of arguments to substitute into `format`.
//
// # Return Value
//
// An [NSString] object initialized by using `format` as a template into which
// the values in `argList` are substituted according to the current locale.
// The returned object may be different from the original receiver.
//
// # Discussion
//
// This method is meant to be called from within a variadic function, where
// the argument list will be available.
//
// This method invokes [InitWithFormatLocaleArguments] without applying any
// localization. This is useful, for example, when working with fixed-format
// representations of information that is written out and read back in at a
// later time.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(format:arguments:)
//
// [Formatting String Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/FormatStrings.html#//apple_ref/doc/uid/20000943
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
func (s NSString) InitWithFormatArguments(format string, argList unsafe.Pointer) NSString {
	rv := objc.Send[NSString](s.ID, objc.Sel("initWithFormat:arguments:"), objc.String(format), argList)
	return rv
}

// Returns an [NSString] object initialized by using a given format string as
// a template into which the remaining argument values are substituted
// according to given locale information. This method is meant to be called
// from within a variadic function, where the argument list will be available.
//
// format: A format string. See [Formatting String Objects] for examples of how to use
// this method, and [String Format Specifiers] for a list of format
// specifiers. This value must not be `nil`.
//
// locale: An [NSLocale] object specifying the locale to use. To use the current
// locale (specified by user preferences), pass [NSLocale] [CurrentLocale]].
// To use the system locale, pass `nil`.
//
// For legacy support, this may be an instance of [NSDictionary] containing
// locale information.
//
// argList: A list of arguments to substitute into `format`.
//
// # Return Value
//
// An [NSString] object initialized by using `format` as a template into which
// values in `argList` are substituted according the locale information in
// `locale`. The returned object may be different from the original receiver.
//
// # Discussion
//
// The following Objective-C code fragment illustrates how to create a string
// from `myArgs`, which is derived from a string object with the value
// “Cost:” and an int with the value 32:
//
// The resulting string has the value “`Cost: 32\n`”.
//
// See [String Programming Guide] for more information.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(format:locale:arguments:)
//
// [Formatting String Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/FormatStrings.html#//apple_ref/doc/uid/20000943
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
// [String Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/introStrings.html#//apple_ref/doc/uid/10000035i
func (s NSString) InitWithFormatLocaleArguments(format string, locale objectivec.IObject, argList unsafe.Pointer) NSString {
	rv := objc.Send[NSString](s.ID, objc.Sel("initWithFormat:locale:arguments:"), objc.String(format), locale, argList)
	return rv
}

// Returns an [NSString] object initialized by converting given data into
// UTF-16 code units using a given encoding.
//
// data: An [NSData] object containing bytes in `encoding` and the default plain
// text format (that is, pure content with no attributes or other markups) for
// that encoding.
//
// encoding: The encoding used by `data`. For possible values, see [NSStringEncoding].
//
// # Return Value
//
// An [NSString] object initialized by converting the bytes in `data` into
// UTF-16 code units using `encoding`. The returned object may be different
// from the original receiver. Returns `nil` if the initialization fails for
// some reason (for example if `data` does not represent valid data for
// `encoding`).
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(data:encoding:)
func (s NSString) InitWithDataEncoding(data INSData, encoding uint) NSString {
	rv := objc.Send[NSString](s.ID, objc.Sel("initWithData:encoding:"), data, encoding)
	return rv
}

// Returns an [NSString] object initialized by reading data from the file at a
// given path using a given encoding.
//
// path: A path to a file.
//
// enc: The encoding of the file at `path`. For possible values, see
// [NSStringEncoding].
//
// # Return Value
//
// An [NSString] object initialized by reading data from the file named by
// `path` using the encoding, `enc`. The returned object may be different from
// the original receiver. If the file can’t be opened or there is an
// encoding error, returns `nil`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfFile:encoding:)
func (s NSString) InitWithContentsOfFileEncodingError(path string, enc uint) (NSString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithContentsOfFile:encoding:error:"), objc.String(path), enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil

}

// Returns an [NSString] object initialized by reading data from the file at a
// given path and returns by reference the encoding used to interpret the
// characters.
//
// path: A path to a file.
//
// enc: Upon return, if the file is read successfully, contains the encoding used
// to interpret the file at `path`. For possible values, see
// [NSStringEncoding].
//
// # Return Value
//
// An [NSString] object initialized by reading data from the file named by
// `path`. The returned object may be different from the original receiver. If
// the file can’t be opened or there is an encoding error, returns `nil`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfFile:usedEncoding:)
func (s NSString) InitWithContentsOfFileUsedEncodingError(path string, enc unsafe.Pointer) (NSString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithContentsOfFile:usedEncoding:error:"), objc.String(path), enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil

}

// Returns the number of bytes required to store the receiver in a given
// encoding.
//
// enc: The encoding for which to determine the receiver’s length.
//
// # Return Value
//
// The number of bytes required to store the receiver in the encoding `enc` in
// a non-external representation. The length does not include space for a
// terminating [NULL] character. Returns `0` if the specified encoding cannot
// be used to convert the receiver or if the amount of memory required for
// storing the results of the encoding conversion would exceed [NSIntegerMax].
//
// # Discussion
//
// The result is exact and is returned in `O(n)` time.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/lengthOfBytes(using:)
//
// [NSIntegerMax]: https://developer.apple.com/documentation/ObjectiveC/NSIntegerMax
func (s NSString) LengthOfBytesUsingEncoding(enc uint) uint {
	rv := objc.Send[uint](s.ID, objc.Sel("lengthOfBytesUsingEncoding:"), enc)
	return rv
}

// Returns the maximum number of bytes needed to store the receiver in a given
// encoding.
//
// enc: The encoding for which to determine the receiver’s length.
//
// # Return Value
//
// The maximum number of bytes needed to store the receiver in `encoding` in a
// non-external representation. The length does not include space for a
// terminating [NULL] character. Returns `0` if the amount of memory required
// for storing the results of the encoding conversion would exceed
// [NSIntegerMax].
//
// # Discussion
//
// The result is an estimate and is returned in `O(1)` time; the estimate may
// be considerably greater than the actual length needed.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/maximumLengthOfBytes(using:)
//
// [NSIntegerMax]: https://developer.apple.com/documentation/ObjectiveC/NSIntegerMax
func (s NSString) MaximumLengthOfBytesUsingEncoding(enc uint) uint {
	rv := objc.Send[uint](s.ID, objc.Sel("maximumLengthOfBytesUsingEncoding:"), enc)
	return rv
}

// Returns the character at a given UTF-16 code unit index.
//
// index: The index of the character to retrieve.
//
// # Return Value
//
// The character at the array position given by `index`.
//
// # Discussion
//
// You should always use the [RangeOfComposedCharacterSequenceAtIndex] or
// [RangeOfComposedCharacterSequencesForRange] method to determine character
// boundaries, so that any surrogate pairs or character clusters are handled
// correctly.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/character(at:)
func (s NSString) CharacterAtIndex(index uint) Unichar {
	rv := objc.Send[Unichar](s.ID, objc.Sel("characterAtIndex:"), index)
	return Unichar(rv)
}

// Copies characters from a given range in the receiver into a given buffer.
//
// buffer: Upon return, contains the characters from the receiver. `buffer` must be
// large enough to contain the characters in the range `aRange`
// (`aRange.Length()*sizeof(unichar)`).
//
// range: The range of characters to retrieve. The range must not exceed the bounds
// of the receiver.
//
// # Discussion
//
// This method does not add a [NULL] character.
//
// The abstract implementation of this method uses [CharacterAtIndex]
// repeatedly, correctly extracting the characters, though very inefficiently.
// Subclasses should override it to provide a fast implementation.
//
// You should always use the [RangeOfComposedCharacterSequenceAtIndex] or
// [RangeOfComposedCharacterSequencesForRange] method to determine character
// boundaries, so that any surrogate pairs or character clusters are handled
// correctly.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/getCharacters(_:range:)
func (s NSString) GetCharactersRange(buffer unsafe.Pointer, range_ NSRange) {
	objc.Send[objc.ID](s.ID, objc.Sel("getCharacters:range:"), buffer, range_)
}

// Gets a given range of characters as bytes in a specified encoding.
//
// buffer: A buffer into which to store the bytes from the receiver. The returned
// bytes are [NULL]-terminated.
//
// maxBufferCount: The maximum number of bytes to write to `buffer`.
//
// usedBufferCount: The number of bytes used from `buffer`. Pass [NULL] if you do not need this
// value.
//
// encoding: The encoding to use for the returned bytes. For possible values, see
// [NSStringEncoding].
//
// options: A mask to specify options to use for converting the receiver’s contents
// to `encoding` (if conversion is necessary).
//
// range: The range of characters in the receiver to get.
//
// leftover: The remaining range. Pass [NULL] If you do not need this value.
//
// # Return Value
//
// true if some characters were converted, otherwise false.
//
// # Discussion
//
// Conversion might stop when the buffer fills, but it might also stop when
// the conversion isn’t possible due to the chosen encoding.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/getBytes(_:maxLength:usedLength:encoding:options:range:remaining:)
func (s NSString) GetBytesMaxLengthUsedLengthEncodingOptionsRangeRemainingRange(buffer unsafe.Pointer, maxBufferCount uint, encoding NSStringEncoding, options NSStringEncodingConversionOptions, range_ NSRange, leftover NSRangePointer) (uint, bool) {
	var usedBufferCount uint
	rv := objc.Send[bool](s.ID, objc.Sel("getBytes:maxLength:usedLength:encoding:options:range:remainingRange:"), buffer, maxBufferCount, unsafe.Pointer(&usedBufferCount), encoding, options, range_, leftover)
	return usedBufferCount, rv
}

// Returns a representation of the string as a C string using a given
// encoding.
//
// encoding: The encoding for the returned C string. For possible values, see
// [NSStringEncoding].
//
// # Return Value
//
// A C string representation of the receiver using the encoding specified by
// `encoding`. Returns [NULL] if the receiver cannot be losslessly converted
// to `encoding`.
//
// # Discussion
//
// The returned C string is guaranteed to be valid only until either the
// receiver is freed, or until the current memory is emptied, whichever occurs
// first. You should copy the C string or use [GetCStringMaxLengthEncoding] if
// it needs to store the C string beyond this time.
//
// You can use [CanBeConvertedToEncoding] to check whether a string can be
// losslessly converted to `encoding`. If it can’t, you can use
// [DataUsingEncodingAllowLossyConversion] to get a C-string representation
// using `encoding`, allowing some loss of information (note that the data
// returned by [DataUsingEncodingAllowLossyConversion] is not a strict
// C-string since it does not have a [NULL] terminator).
//
// # Special Considerations
//
// UTF-16 and UTF-32 are not considered to be C string encodings, and should
// not be used with this method—the results of passing
// [NSUTF16StringEncoding], [NSUTF32StringEncoding], or any of their variants
// are undefined.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/cString(using:)
func (s NSString) CStringUsingEncoding(encoding uint) string {
	rv := objc.Send[*byte](s.ID, objc.Sel("cStringUsingEncoding:"), encoding)
	return objc.GoString(rv)
}

// Converts the string to a given encoding and stores it in a buffer.
//
// buffer: Upon return, contains the converted C-string plus the [NULL] termination
// byte. The buffer must include room for `maxBufferCount` bytes.
//
// maxBufferCount: The maximum number of bytes in the string to return in buffer ( the [NULL]
// termination byte).
//
// encoding: The encoding for the returned C string. For possible values, see
// [NSStringEncoding].
//
// # Return Value
//
// true if the operation was successful, otherwise false. Returns false if
// conversion is not possible due to encoding errors or if `buffer` is too
// small.
//
// # Discussion
//
// Note that in the treatment of the `maxBufferCount` argument, this method
// differs from the deprecated [GetCStringMaxLength] method which it replaces.
// (The buffer should include room for `maxBufferCount` bytes; this number
// should accommodate the expected size of the return value plus the [NULL]
// termination byte, which this method adds.)
//
// You can use [CanBeConvertedToEncoding] to check whether a string can be
// losslessly converted to `encoding`. If it can’t, you can use
// [DataUsingEncodingAllowLossyConversion] to get a C-string representation
// using `encoding`, allowing some loss of information (note that the data
// returned by [DataUsingEncodingAllowLossyConversion] is not a strict
// C-string since it does not have a [NULL] terminator).
//
// See: https://developer.apple.com/documentation/Foundation/NSString/getCString(_:maxLength:encoding:)
func (s NSString) GetCStringMaxLengthEncoding(buffer string, maxBufferCount uint, encoding uint) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("getCString:maxLength:encoding:"), unsafe.Pointer(unsafe.StringData(buffer+"\x00")), maxBufferCount, encoding)
	return rv
}

// Returns the result of invoking [CompareOptions] with
// [NSCaseInsensitiveSearch] as the only option.
//
// string: The string with which to compare the receiver.
//
// # Return Value
//
// Returns an [ComparisonResult] value that indicates the lexical ordering.
// [NSOrderedAscending] the receiver precedes `aString` in lexical ordering,
// [NSOrderedSame] the receiver and `aString` are equivalent in lexical value,
// and [NSOrderedDescending] if the receiver follows `aString`.
//
// # Discussion
//
// This method is the equivalent of invoking [CompareOptions] with
// [NSCaseInsensitiveSearch] as the only option.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/caseInsensitiveCompare(_:)
//
// [ComparisonResult]: https://developer.apple.com/documentation/Foundation/ComparisonResult
func (s NSString) CaseInsensitiveCompare(string_ string) ComparisonResult {
	rv := objc.Send[ComparisonResult](s.ID, objc.Sel("caseInsensitiveCompare:"), objc.String(string_))
	return ComparisonResult(rv)
}

// Compares the string with a given string using a case-insensitive,
// localized, comparison.
//
// string: The string with which to compare the receiver.
//
// This value must not be `nil`. If this value is `nil`, the behavior is
// undefined and may change in future versions of macOS.
//
// # Return Value
//
// Returns an [ComparisonResult] value that indicates the lexical ordering.
// [NSOrderedAscending] the receiver precedes `aString` in lexical ordering,
// [NSOrderedSame] the receiver and `aString` are equivalent in lexical value,
// and [NSOrderedDescending] if the receiver follows `aString`.
//
// # Discussion
//
// This method uses the current locale.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/localizedCaseInsensitiveCompare(_:)
//
// [ComparisonResult]: https://developer.apple.com/documentation/Foundation/ComparisonResult
func (s NSString) LocalizedCaseInsensitiveCompare(string_ string) ComparisonResult {
	rv := objc.Send[ComparisonResult](s.ID, objc.Sel("localizedCaseInsensitiveCompare:"), objc.String(string_))
	return ComparisonResult(rv)
}

// Returns the result of invoking [CompareOptionsRange] with no options and
// the receiver’s full extent as the range.
//
// string: The string with which to compare the receiver.
//
// This value must not be `nil`. If this value is `nil`, the behavior is
// undefined and may change in future versions of macOS.
//
// # Return Value
//
// Returns an [ComparisonResult] value that indicates the lexical ordering.
// [NSOrderedAscending] the receiver precedes `aString` in lexical ordering,
// [NSOrderedSame] the receiver and `aString` are equivalent in lexical value,
// and [NSOrderedDescending] if the receiver follows `aString`.
//
// # Discussion
//
// This method is equivalent to invoking [CompareOptionsRange] with no options
// and the receiver’s full extent as the range.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/compare(_:)
//
// [ComparisonResult]: https://developer.apple.com/documentation/Foundation/ComparisonResult
func (s NSString) Compare(string_ string) ComparisonResult {
	rv := objc.Send[ComparisonResult](s.ID, objc.Sel("compare:"), objc.String(string_))
	return ComparisonResult(rv)
}

// Compares the string and a given string using a localized comparison.
//
// string: The string with which to compare the receiver.
//
// This value must not be `nil`. If this value is `nil`, the behavior is
// undefined and may change in future versions of macOS.
//
// # Return Value
//
// Returns an [ComparisonResult] value that indicates the lexical ordering.
// [NSOrderedAscending] the receiver precedes `aString` in lexical ordering,
// [NSOrderedSame] the receiver and `aString` are equivalent in lexical value,
// and [NSOrderedDescending] if the receiver follows `aString`.
//
// # Discussion
//
// This method uses the current locale.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/localizedCompare(_:)
//
// [ComparisonResult]: https://developer.apple.com/documentation/Foundation/ComparisonResult
func (s NSString) LocalizedCompare(string_ string) ComparisonResult {
	rv := objc.Send[ComparisonResult](s.ID, objc.Sel("localizedCompare:"), objc.String(string_))
	return ComparisonResult(rv)
}

// Compares the string with the specified string using the given options.
//
// string: The string with which to compare the receiver.
//
// This value must not be `nil`. If this value is `nil`, the behavior is
// undefined and may change in future versions of macOS.
//
// mask: Options for the search—you can combine any of the following using a C
// bitwise OR operator: [NSCaseInsensitiveSearch], [NSLiteralSearch],
// [NSNumericSearch]. See [String Programming Guide] for details on these
// options.
//
// # Return Value
//
// Returns an [ComparisonResult] value that indicates the lexical ordering.
// [NSOrderedAscending] the receiver precedes `aString` in lexical ordering,
// [NSOrderedSame] the receiver and `aString` are equivalent in lexical value,
// and [NSOrderedDescending] if the receiver follows `aString`.
//
// # Discussion
//
// This method is equivalent to invoking [CompareOptionsRange] with a given
// mask as the options and the receiver’s full extent as the range.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/compare(_:options:)
//
// [String Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/introStrings.html#//apple_ref/doc/uid/10000035i
// [ComparisonResult]: https://developer.apple.com/documentation/Foundation/ComparisonResult
func (s NSString) CompareOptions(string_ string, mask NSStringCompareOptions) ComparisonResult {
	rv := objc.Send[ComparisonResult](s.ID, objc.Sel("compare:options:"), objc.String(string_), mask)
	return ComparisonResult(rv)
}

// Returns the result of invoking [CompareOptionsRangeLocale] with a `nil`
// locale.
//
// string: The string with which to compare the range of the receiver specified by
// `range`.
//
// This value must not be `nil`. If this value is `nil`, the behavior is
// undefined and may change in future versions of macOS.
//
// mask: Options for the search—you can combine any of the following using a C
// bitwise OR operator: [NSCaseInsensitiveSearch], [NSLiteralSearch],
// [NSNumericSearch].
//
// See [String Programming Guide] for details on these options.
//
// rangeOfReceiverToCompare: The range of the receiver over which to perform the comparison. The range
// must not exceed the bounds of the receiver.
//
// # Return Value
//
// Returns an [ComparisonResult] value that indicates the lexical ordering.
// [NSOrderedAscending] the receiver precedes `aString` in lexical ordering,
// [NSOrderedSame] the receiver and `aString` are equivalent in lexical value,
// and [NSOrderedDescending] if the receiver follows `aString`.
//
// # Discussion
//
// This method is equivalent to invoking [CompareOptionsRangeLocale] with a
// `nil` locale.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/compare(_:options:range:)
//
// [String Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/introStrings.html#//apple_ref/doc/uid/10000035i
// [ComparisonResult]: https://developer.apple.com/documentation/Foundation/ComparisonResult
func (s NSString) CompareOptionsRange(string_ string, mask NSStringCompareOptions, rangeOfReceiverToCompare NSRange) ComparisonResult {
	rv := objc.Send[ComparisonResult](s.ID, objc.Sel("compare:options:range:"), objc.String(string_), mask, rangeOfReceiverToCompare)
	return ComparisonResult(rv)
}

// Compares the string using the specified options and returns the lexical
// ordering for the range.
//
// string: The string with which to compare the range of the receiver specified by
// `range`.
//
// This value must not be `nil`. If this value is `nil`, the behavior is
// undefined and may change in future versions of macOS.
//
// mask: Options for the search—you can combine any of the following using a C
// bitwise OR operator: [NSCaseInsensitiveSearch], [NSLiteralSearch],
// [NSNumericSearch].
//
// See [String Programming Guide] for details on these options.
//
// rangeOfReceiverToCompare: The range of the receiver over which to perform the comparison. The range
// must not exceed the bounds of the receiver.
//
// locale: An instance of [NSLocale]. To use the current locale, pass [NSLocale]
// [CurrentLocale]]. For example, if you are comparing strings to present to
// the end-user, use the current locale. To use the system locale, pass `nil`.
//
// # Return Value
//
// Returns an [ComparisonResult] value that indicates the lexical ordering of
// a specified range within the receiver and a given string.
// [NSOrderedAscending] if the substring of the receiver given by `range`
// precedes `aString` in lexical ordering for the locale given in `dict`,
// [NSOrderedSame] if the substring of the receiver and `aString` are
// equivalent in lexical value, and [NSOrderedDescending] if the substring of
// the receiver follows `aString`.
//
// # Discussion
//
// The `locale` argument affects both equality and ordering algorithms. For
// example, in some locales, accented characters are ordered immediately after
// the base; other locales order them after “z”.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/compare(_:options:range:locale:)
//
// [String Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/introStrings.html#//apple_ref/doc/uid/10000035i
// [ComparisonResult]: https://developer.apple.com/documentation/Foundation/ComparisonResult
func (s NSString) CompareOptionsRangeLocale(string_ string, mask NSStringCompareOptions, rangeOfReceiverToCompare NSRange, locale objectivec.IObject) ComparisonResult {
	rv := objc.Send[ComparisonResult](s.ID, objc.Sel("compare:options:range:locale:"), objc.String(string_), mask, rangeOfReceiverToCompare, locale)
	return ComparisonResult(rv)
}

// Compares strings as sorted by the Finder.
//
// string: The string to compare with the receiver.
//
// # Return Value
//
// The result of the comparison.
//
// # Discussion
//
// This method should be used whenever file names or other strings are
// presented in lists and tables where Finder-like sorting is appropriate. The
// exact sorting behavior of this method is different under different locales
// and may be changed in future releases. This method uses the current locale.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/localizedStandardCompare(_:)
func (s NSString) LocalizedStandardCompare(string_ string) ComparisonResult {
	rv := objc.Send[ComparisonResult](s.ID, objc.Sel("localizedStandardCompare:"), objc.String(string_))
	return ComparisonResult(rv)
}

// Returns a Boolean value that indicates whether a given string matches the
// beginning characters of the receiver.
//
// str: A string.
//
// # Return Value
//
// true if `aString` matches the beginning characters of the receiver,
// otherwise false. Returns false if `aString` is empty.
//
// # Discussion
//
// This method is a convenience for comparing strings using the
// [NSAnchoredSearch] option. See [String Programming Guide] for more
// information.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/hasPrefix(_:)
//
// [String Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/introStrings.html#//apple_ref/doc/uid/10000035i
func (s NSString) HasPrefix(str string) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasPrefix:"), objc.String(str))
	return rv
}

// Returns a Boolean value that indicates whether a given string matches the
// ending characters of the receiver.
//
// str: A string.
//
// # Return Value
//
// true if `aString` matches the ending characters of the receiver, otherwise
// false. Returns false if `aString` is empty.
//
// # Discussion
//
// This method is a convenience for comparing strings using the
// [NSAnchoredSearch] and [NSBackwardsSearch] options. See [String Programming
// Guide] for more information.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/hasSuffix(_:)
//
// [String Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/introStrings.html#//apple_ref/doc/uid/10000035i
func (s NSString) HasSuffix(str string) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasSuffix:"), objc.String(str))
	return rv
}

// Returns a Boolean value that indicates whether a given string is equal to
// the receiver using a literal Unicode-based comparison.
//
// aString: The string with which to compare the receiver.
//
// # Return Value
//
// true if `aString` is equivalent to the receiver (if they have the same id
// or if they are [NSOrderedSame] in a literal comparison), otherwise false.
//
// # Discussion
//
// The comparison uses the canonical representation of strings, which for a
// particular string is the length of the string plus the UTF-16 code units
// that make up the string. When this method compares two strings, if the
// individual Unicodes are the same, then the strings are equal, regardless of
// the backing store. “Literal” when applied to string comparison means
// that various Unicode decomposition rules are not applied and UTF-16 code
// units are individually compared. So, for instance, “Ö” represented as
// the composed character sequence “O” (`U+004F LATIN CAPITAL LETTER O`)
// and a combining diaeresis “¨” (`U+0308 COMBINING DIAERESIS`) would not
// compare equal to “Ö” represented as a single Unicode character
// (`U+00D6 LATIN CAPITAL LETTER O WITH DIAERESIS`).
//
// # Special Considerations
//
// When you know both objects are strings, this method is a faster way to
// check equality than [isEqual(_:)].
//
// See: https://developer.apple.com/documentation/Foundation/NSString/isEqual(to:)
//
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
func (s NSString) IsEqualToString(aString string) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isEqualToString:"), objc.String(aString))
	return rv
}

// Returns a new string made by appending a given string to the receiver.
//
// aString: The string to append to the receiver. This value must not be `nil`.
//
// # Return Value
//
// A new string made by appending `aString` to the receiver.
//
// # Discussion
//
// This code excerpt, for example:
//
// produces the string “`Error: premature end of file.`”.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/appending(_:)
func (s NSString) StringByAppendingString(aString string) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByAppendingString:"), objc.String(aString))
	return NSStringFromID(rv).String()
}

// Returns a new string formed from the receiver by either removing characters
// from the end, or by appending as many occurrences as necessary of a given
// pad string.
//
// newLength: The new length for the receiver.
//
// padString: The string with which to extend the receiver.
//
// padIndex: The index in `padString` from which to start padding.
//
// # Return Value
//
// A new string formed from the receiver by either removing characters from
// the end, or by appending as many occurrences of `padString` as necessary.
//
// # Discussion
//
// Here are some examples of usage:
//
// See: https://developer.apple.com/documentation/Foundation/NSString/padding(toLength:withPad:startingAt:)
func (s NSString) StringByPaddingToLengthWithStringStartingAtIndex(newLength uint, padString string, padIndex uint) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByPaddingToLength:withString:startingAtIndex:"), newLength, objc.String(padString), padIndex)
	return NSStringFromID(rv).String()
}

// Returns a version of the string with all letters converted to lowercase,
// taking into account the specified locale.
//
// locale: The locale. For strings presented to users, pass the current locale
// ([NSLocale] [CurrentLocale]]). To use the system locale, pass `nil`.
//
// # Return Value
//
// A lowercase string using the `locale`.
//
// # Discussion
//
// Case transformations aren’t guaranteed to be symmetrical or to produce
// strings of the same lengths as the originals. See [LowercaseString] for an
// example.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/lowercased(with:)
func (s NSString) LowercaseStringWithLocale(locale INSLocale) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("lowercaseStringWithLocale:"), locale)
	return NSStringFromID(rv).String()
}

// Returns a version of the string with all letters converted to uppercase,
// taking into account the specified locale.
//
// locale: The locale. For strings presented to users, pass the current locale
// ([NSLocale] [CurrentLocale]]). To use the system locale, pass `nil`.
//
// # Return Value
//
// An uppercase string using the `locale`.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/uppercased(with:)
func (s NSString) UppercaseStringWithLocale(locale INSLocale) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("uppercaseStringWithLocale:"), locale)
	return NSStringFromID(rv).String()
}

// Returns a capitalized representation of the receiver using the specified
// locale.
//
// locale: The locale. For strings presented to users, pass the current locale
// ([NSLocale] [CurrentLocale]]). To use the system locale, pass `nil`.
//
// # Return Value
//
// A string with the first character from each word in the receiver changed to
// its corresponding uppercase value, and all remaining characters set to
// their corresponding lowercase values.
//
// # Discussion
//
// A capitalized string is a string with the first character in each word
// changed to its corresponding uppercase value, and all remaining characters
// set to their corresponding lowercase values. A “word” is any sequence
// of characters delimited by spaces, tabs, or line terminators (listed under
// [GetLineStartEndContentsEndForRange]). Some common word delimiting
// punctuation isn’t considered, so this property may not generally produce
// the desired results for multiword strings.
//
// Case transformations aren’t guaranteed to be symmetrical or to produce
// strings of the same lengths as the originals. See [LowercaseString] for an
// example.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/capitalized(with:)
func (s NSString) CapitalizedStringWithLocale(locale INSLocale) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("capitalizedStringWithLocale:"), locale)
	return NSStringFromID(rv).String()
}

// Returns an array containing substrings from the receiver that have been
// divided by a given separator.
//
// separator: The separator string.
//
// # Return Value
//
// An [NSArray] object containing substrings from the receiver that have been
// divided by `separator`.
//
// # Discussion
//
// The substrings in the array appear in the order they did in the receiver.
// Adjacent occurrences of the separator string produce empty strings in the
// result. Similarly, if the string begins or ends with the separator, the
// first or last substring, respectively, is empty. For example, this code
// fragment:
//
// produces the array `@[@"Karin", @"Carrie", @"David"]`.
//
// If `list` begins with a comma and space—for example, `@", Norman,
// Stanley, Fletcher"`—the array has these contents: `@[@"", @"Norman",
// @"Stanley", @"Fletcher"]`.
//
// If `list` has no separators—for example, `@"Karin"`—the array contains
// the string itself, in this case `@[@"Karin"]`.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/components(separatedBy:)-238fy
func (s NSString) ComponentsSeparatedByString(separator string) []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("componentsSeparatedByString:"), objc.String(separator))
	return objc.ConvertSliceToStrings(rv)
}

// Returns an array containing substrings from the receiver that have been
// divided by characters in a given set.
//
// separator: A character set containing the characters to use to split the receiver.
// Must not be `nil`.
//
// # Return Value
//
// An [NSArray] object containing substrings from the receiver that have been
// divided by characters in `separator`.
//
// # Discussion
//
// The substrings in the array appear in the order they did in the receiver.
// Adjacent occurrences of the separator characters produce empty strings in
// the result. Similarly, if the string begins or ends with separator
// characters, the first or last substring, respectively, is empty.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/components(separatedBy:)-27x9g
func (s NSString) ComponentsSeparatedByCharactersInSet(separator INSCharacterSet) []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("componentsSeparatedByCharactersInSet:"), separator)
	return objc.ConvertSliceToStrings(rv)
}

// Returns a new string made by removing from both ends of the receiver
// characters contained in a given character set.
//
// set: A character set containing the characters to remove from the receiver.
// `set` must not be `nil`.
//
// # Return Value
//
// A new string made by removing from both ends of the receiver characters
// contained in `set`. If the receiver is composed entirely of characters from
// `set`, the empty string is returned.
//
// # Discussion
//
// Use [WhitespaceCharacterSet] or [WhitespaceAndNewlineCharacterSet] to
// remove whitespace around strings.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/trimmingCharacters(in:)
func (s NSString) StringByTrimmingCharactersInSet(set INSCharacterSet) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByTrimmingCharactersInSet:"), set)
	return NSStringFromID(rv).String()
}

// Returns a new string containing the characters of the receiver from the one
// at a given index to the end.
//
// from: An index. The value must lie within the bounds of the receiver, or be equal
// to the length of the receiver.
//
// Raises an [RangeException] if (`anIndex` - 1) lies beyond the end of the
// receiver.
//
// # Return Value
//
// A new string containing the characters of the receiver from the one at
// `anIndex` to the end. If `anIndex` is equal to the length of the string,
// returns an empty string.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/substring(from:)
func (s NSString) SubstringFromIndex(from uint) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("substringFromIndex:"), from)
	return NSStringFromID(rv).String()
}

// Returns a string object containing the characters of the receiver that lie
// within a given range.
//
// range: A range. The range must not exceed the bounds of the receiver.
//
// Raises an [RangeException] if (`aRange.Location()` - 1) or
// (`aRange.Location()` + `aRange.Length()` - 1) lies beyond the end of the
// receiver.
//
// # Return Value
//
// A string object containing the characters of the receiver that lie within
// `aRange`.
//
// # Discussion
//
// This method detects all invalid ranges (including those with negative
// lengths). For applications linked against macOS 10.6 and later, this error
// causes an exception; for applications linked against earlier releases, this
// error causes a warning, which is displayed just once per application
// execution.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/substring(with:)
func (s NSString) SubstringWithRange(range_ NSRange) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("substringWithRange:"), range_)
	return NSStringFromID(rv).String()
}

// Returns a new string containing the characters of the receiver up to, but
// not including, the one at a given index.
//
// to: An index. The value must lie within the bounds of the receiver, or be equal
// to the length of the receiver.
//
// Raises an [RangeException] if (`anIndex` - 1) lies beyond the end of the
// receiver.
//
// # Return Value
//
// A new string containing the characters of the receiver up to, but not
// including, the one at `anIndex`. If `anIndex` is equal to the length of the
// string, returns a copy of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/substring(to:)
func (s NSString) SubstringToIndex(to uint) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("substringToIndex:"), to)
	return NSStringFromID(rv).String()
}

// Creates a string suitable for comparison by removing the specified
// character distinctions from a string.
//
// options: Any combination of the [NSCaseInsensitiveSearch],
// [NSWidthInsensitiveSearch], and [NSDiacriticInsensitiveSearch] comparison
// options.
//
// locale: The locale to use for the folding operation. Pass `nil` to use the system
// locale.
//
// # Return Value
//
// A string created by performing a character folding operation with the
// specified options and locale.
//
// # Discussion
//
// When working with text—especially text in Latin script—it’s often
// useful to compare strings in such a way that ignores differences in case
// (uppercase or lowercase), width (full-width or half-width), and/or
// diacritics (accents and other marks).
//
// To accomplish this, you may use one of the methods described in Identifying
// and Comparing Strings, passing one or more options specified by the
// [NSString.CompareOptions] type as appropriate. If you’re performing many
// such comparisons, you may instead, as an optimization, perform a single
// folding operation on the string and store the result, passing that into a
// more efficient comparison method. For example, if you were determining the
// case-insensitive (that is, “HELLO”, “hello”, and “Hello” are
// all considered equal) intersection of two sets of strings, instead of
// calling the [LocalizedCaseInsensitiveCompare] method for each pair of
// strings, you might first normalize both sets of strings by calling the
// [StringByFoldingWithOptionsLocale] method, passing the
// [NSCaseInsensitiveSearch] option and the current locale, and then compare
// each pair of folded strings using the [IsEqualToString] method.
//
// Rules for how characters are folded may vary, depending on the locale. For
// example, when folding a string containing the character “I” (`U+0049`
// `LATIN CAPITAL LETTER I`) and specifying the [NSCaseInsensitiveSearch]
// comparison option, an English-speaking locale returns “i” (`U+0069`
// `LATIN SMALL LETTER I`), and a Turkish-speaking locale returns “ı”
// (`U+0131 LATIN SMALL DOTLESS I`).
//
// See: https://developer.apple.com/documentation/Foundation/NSString/folding(options:locale:)
//
// [NSString.CompareOptions]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions
func (s NSString) StringByFoldingWithOptionsLocale(options NSStringCompareOptions, locale INSLocale) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByFoldingWithOptions:locale:"), options, locale)
	return NSStringFromID(rv).String()
}

// Returns a new string by applying a specified transform to the string.
//
// # Discussion
//
// You can use this method to, for example, transliterate text from one script
// to another, strip diacritics or combining marks, and get the unicode names
// of characters.
//
// .
//
// See: https://developer.apple.com/documentation/Foundation/NSString/applyingTransform(_:reverse:)
func (s NSString) StringByApplyingTransformReverse(transform NSStringTransform, reverse bool) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByApplyingTransform:reverse:"), objc.String(string(transform)), reverse)
	return NSStringFromID(rv).String()
}

// Returns a Boolean value indicating whether the string contains a given
// string by performing a case-sensitive, locale-unaware search.
//
// str: The string to search for. This value must not be `nil`.
//
// # Return Value
//
// true if the receiver contains `str`, otherwise false.
//
// # Discussion
//
// Calling this method is equivalent to calling [RangeOfStringOptions] with no
// options.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/contains(_:)
func (s NSString) ContainsString(str string) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("containsString:"), objc.String(str))
	return rv
}

// Returns a Boolean value indicating whether the string contains a given
// string by performing a case-insensitive, locale-aware search.
//
// str: The string to search for. This value must not be `nil`.
//
// # Return Value
//
// true if the receiver contains `str`, otherwise false.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/localizedCaseInsensitiveContains(_:)
func (s NSString) LocalizedCaseInsensitiveContainsString(str string) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("localizedCaseInsensitiveContainsString:"), objc.String(str))
	return rv
}

// Returns a Boolean value indicating whether the string contains a given
// string by performing a case and diacritic insensitive, locale-aware search.
//
// str: The string to search for. This value must not be `nil`.
//
// # Return Value
//
// true if the receiver contains `str`, otherwise false.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/localizedStandardContains(_:)
func (s NSString) LocalizedStandardContainsString(str string) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("localizedStandardContainsString:"), objc.String(str))
	return rv
}

// Finds and returns the range in the string of the first character from a
// given character set.
//
// searchSet: A character set. This value must not be `nil`.
//
// Raises an [InvalidArgumentException] if `aSet` is `nil`.
//
// # Return Value
//
// The range in the receiver of the first character found from `aSet`. Returns
// a range of `{“NSNotFound“, 0}` if none of the characters in `aSet` are
// found.
//
// # Discussion
//
// Invokes [RangeOfCharacterFromSetOptions] with no options.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/rangeOfCharacter(from:)
func (s NSString) RangeOfCharacterFromSet(searchSet INSCharacterSet) NSRange {
	rv := objc.Send[NSRange](s.ID, objc.Sel("rangeOfCharacterFromSet:"), searchSet)
	return NSRange(rv)
}

// Finds and returns the range in the string of the first character, using
// given options, from a given character set.
//
// searchSet: A character set. This value must not be `nil`.
//
// Raises an [NSInvalidArgumentException] if `aSet` is `nil`.
//
// mask: A mask specifying search options. The following options may be specified by
// combining them with the C bitwise [OR] operator: [NSAnchoredSearch],
// [NSBackwardsSearch].
//
// # Return Value
//
// The range in the receiver of the first character found from `aSet`. Returns
// a range of `{“NSNotFound“, 0}` if none of the characters in `aSet` are
// found.
//
// # Discussion
//
// Invokes [RangeOfCharacterFromSetOptionsRange] with `mask` for the options
// and the entire extent of the receiver for the range.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/rangeOfCharacter(from:options:)
func (s NSString) RangeOfCharacterFromSetOptions(searchSet INSCharacterSet, mask NSStringCompareOptions) NSRange {
	rv := objc.Send[NSRange](s.ID, objc.Sel("rangeOfCharacterFromSet:options:"), searchSet, mask)
	return NSRange(rv)
}

// Finds and returns the range in the string of the first character from a
// given character set found in a given range with given options.
//
// searchSet: A character set. This value must not be `nil`.
//
// Raises an [NSInvalidArgumentException] if `aSet` is `nil`.
//
// mask: A mask specifying search options. The following options may be specified by
// combining them with the C bitwise [OR] operator: [NSAnchoredSearch],
// [NSBackwardsSearch].
//
// rangeOfReceiverToSearch: The range in which to search. `aRange` must not exceed the bounds of the
// receiver.
//
// Raises an [RangeException] if `aRange` is invalid.
//
// # Return Value
//
// The range in the receiver of the first character found from `aSet` within
// `aRange`. Returns a range of `{“NSNotFound“, 0}` if none of the
// characters in `aSet` are found.
//
// # Discussion
//
// This method does not perform any Unicode normalization on the receiver, so
// canonically equivalent forms will not be matched. For example, searching
// the string “strüdel”—containing the decomposed characters “`u`”
// (`U+0075 LATIN SMALL LETTER U`) and “`¨`” (`U+0308 COMBINING
// DIAERESIS`)—with a character set containing the precomposed character
// “`ü`” (`U+00FC LATIN SMALL LETTER U WITH DIAERESIS`) would return the
// range `{“NSNotFound“, 0}`, because none of the characters in the set are
// found.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/rangeOfCharacter(from:options:range:)
func (s NSString) RangeOfCharacterFromSetOptionsRange(searchSet INSCharacterSet, mask NSStringCompareOptions, rangeOfReceiverToSearch NSRange) NSRange {
	rv := objc.Send[NSRange](s.ID, objc.Sel("rangeOfCharacterFromSet:options:range:"), searchSet, mask, rangeOfReceiverToSearch)
	return NSRange(rv)
}

// Finds and returns the range of the first occurrence of a given string
// within the string.
//
// searchString: The string to search for.
//
// # Return Value
//
// An [NSRange] structure giving the location and length in the receiver of
// the first occurrence of `searchString`. Returns `{“NSNotFound“, 0}` if
// `searchString` is not found or is empty (`""`).
//
// # Discussion
//
// Invokes [RangeOfStringOptions] with no options.
//
// [NSString] objects are compared by checking the Unicode canonical
// equivalence of their code point sequences. The length of the returned range
// and that of `searchString` may differ if equivalent composed character
// sequences are matched.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/range(of:)
func (s NSString) RangeOfString(searchString string) NSRange {
	rv := objc.Send[NSRange](s.ID, objc.Sel("rangeOfString:"), objc.String(searchString))
	return NSRange(rv)
}

// Finds and returns the range of the first occurrence of a given string
// within the string, subject to given options.
//
// searchString: The string to search for.
//
// mask: A mask specifying search options. For possible values, see
// [NSString.CompareOptions].
//
// # Return Value
//
// An [NSRange] structure giving the location and length in the receiver of
// the first occurrence of
//
// # Discussion
//
// `searchString`,
//
// modulo the options in `mask`. Returns `{“NSNotFound“, 0}` if
//
// `searchString`
//
// is not found or is empty (`""`).
//
// # Discussion
//
// Invokes [RangeOfStringOptionsRange] with the options specified by `mask`
// and the entire extent of the receiver as the range.
//
// [NSString] objects are compared by checking the Unicode canonical
// equivalence of their code point sequences. The length of the returned range
// and that of
//
// `searchString`
//
// may differ if equivalent composed character sequences are matched.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/range(of:options:)
//
// [NSString.CompareOptions]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions
func (s NSString) RangeOfStringOptions(searchString string, mask NSStringCompareOptions) NSRange {
	rv := objc.Send[NSRange](s.ID, objc.Sel("rangeOfString:options:"), objc.String(searchString), mask)
	return NSRange(rv)
}

// Finds and returns the range of the first occurrence of a given string,
// within the given range of the string, subject to given options.
//
// searchString: The string for which to search.
//
// mask: A mask specifying search options. The following options may be specified by
// combining them with the C bitwise [OR] operator: [NSCaseInsensitiveSearch],
// [NSLiteralSearch], [NSBackwardsSearch], and [NSAnchoredSearch]. See [String
// Programming Guide] for details on these options.
//
// rangeOfReceiverToSearch: The range within the receiver for which to search for `aString`.
//
// Raises an [NSRangeException] if
//
// `rangeOfReceiverToSearch`
//
// is invalid.
//
// # Return Value
//
// An [NSRange] structure giving the location and length in the receiver of
//
// # Discussion
//
// `searchString`
//
// within
//
// `rangeOfReceiverToSearch`
//
// in the receiver, modulo the options in `mask`. The range returned is
// relative to the start of the string, not to the passed-in range. Returns
// `{“NSNotFound“, 0}` if
//
// `searchString`
//
// is not found or is empty (`""`).
//
// # Discussion
//
// [NSString] objects are compared by checking the Unicode canonical
// equivalence of their code point sequences. T
//
// he length of the returned range and that of
//
// `searchString`
//
// may differ if equivalent composed character sequences are matched.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/range(of:options:range:)
//
// [String Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/introStrings.html#//apple_ref/doc/uid/10000035i
func (s NSString) RangeOfStringOptionsRange(searchString string, mask NSStringCompareOptions, rangeOfReceiverToSearch NSRange) NSRange {
	rv := objc.Send[NSRange](s.ID, objc.Sel("rangeOfString:options:range:"), objc.String(searchString), mask, rangeOfReceiverToSearch)
	return NSRange(rv)
}

// Finds and returns the range of the first occurrence of a given string
// within a given range of the string, subject to given options, using the
// specified locale, if any.
//
// searchString: The string for which to search.
//
// mask: A mask specifying search options. The following options may be specified by
// combining them with the C bitwise [OR] operator: [NSCaseInsensitiveSearch],
// [NSLiteralSearch], [NSBackwardsSearch], and [NSAnchoredSearch]. See [String
// Programming Guide] for details on these options.
//
// rangeOfReceiverToSearch: The range within the receiver for which to search for `aString`.
//
// Raises an [RangeException] if `aRange` is invalid.
//
// locale: The locale to use when comparing the receiver with `aString`. To use the
// current locale, pass [NSLocale] [CurrentLocale]]. To use the system locale,
// pass `nil`.
//
// The locale argument affects the equality checking algorithm. For example,
// for the Turkish locale, case-insensitive compare matches “I” to
// “ı” (`U+0131 LATIN SMALL DOTLESS I`), not the normal “i”
// character.
//
// # Return Value
//
// An [NSRange] structure giving the location and length in the receiver of
// `aString` within `aRange` in the receiver, modulo the options in `mask`.
// The range returned is relative to the start of the string, not to the
// passed-in range. Returns `{“NSNotFound“, 0}` if `aString` is not found or
// is empty (`""`).
//
// # Discussion
//
// [NSString] objects are compared by checking the Unicode canonical
// equivalence of their code point sequences. The length of the returned range
// and that of `aString` may differ if equivalent composed character sequences
// are matched.
//
// # Special Considerations
//
// This method detects all invalid ranges (including those with negative
// lengths). For applications linked against macOS 10.6 and later, this error
// causes an exception; for applications linked against earlier releases, this
// error causes a warning, which is displayed just once per application
// execution.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/range(of:options:range:locale:)
//
// [String Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/introStrings.html#//apple_ref/doc/uid/10000035i
// [NSRange]: https://developer.apple.com/documentation/Foundation/NSRange-c.struct
func (s NSString) RangeOfStringOptionsRangeLocale(searchString string, mask NSStringCompareOptions, rangeOfReceiverToSearch NSRange, locale INSLocale) NSRange {
	rv := objc.Send[NSRange](s.ID, objc.Sel("rangeOfString:options:range:locale:"), objc.String(searchString), mask, rangeOfReceiverToSearch, locale)
	return NSRange(rv)
}

// Finds and returns the range of the first occurrence of a given string
// within the string by performing a case and diacritic insensitive,
// locale-aware search.
//
// str: The string to search for. This value must not be `nil`.
//
// # Return Value
//
// The range of the first occurrence of `str` in the receiver. Returns
// `{“NSNotFound“, 0}` if `str` is not found.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/localizedStandardRange(of:)
func (s NSString) LocalizedStandardRangeOfString(str string) NSRange {
	rv := objc.Send[NSRange](s.ID, objc.Sel("localizedStandardRangeOfString:"), objc.String(str))
	return NSRange(rv)
}

// Returns a new string in which all occurrences of a target string in the
// receiver are replaced by another given string.
//
// target: The string to replace.
//
// replacement: The string with which to replace `target`.
//
// # Return Value
//
// A new string in which all occurrences of `target` in the receiver are
// replaced by `replacement`.
//
// # Discussion
//
// Invokes [StringByReplacingOccurrencesOfStringWithStringOptionsRange]with
// `0` options and range of the whole string.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/replacingOccurrences(of:with:)
func (s NSString) StringByReplacingOccurrencesOfStringWithString(target string, replacement string) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByReplacingOccurrencesOfString:withString:"), objc.String(target), objc.String(replacement))
	return NSStringFromID(rv).String()
}

// Returns a new string in which all occurrences of a target string in a
// specified range of the receiver are replaced by another given string.
//
// target: The string to replace.
//
// replacement: The string with which to replace `target`.
//
// options: A mask of options to use when comparing `target` with the receiver. Pass
// `0` to specify no options.
//
// searchRange: The range in the receiver in which to search for `target`.
//
// # Return Value
//
// A new string in which all occurrences of `target`, matched using `options`,
// in `searchRange` of the receiver are replaced by `replacement`.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/replacingOccurrences(of:with:options:range:)
func (s NSString) StringByReplacingOccurrencesOfStringWithStringOptionsRange(target string, replacement string, options NSStringCompareOptions, searchRange NSRange) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByReplacingOccurrencesOfString:withString:options:range:"), objc.String(target), objc.String(replacement), options, searchRange)
	return NSStringFromID(rv).String()
}

// Returns a new string in which the characters in a specified range of the
// receiver are replaced by a given string.
//
// range: A range of characters in the receiver.
//
// replacement: The string with which to replace the characters in `range`.
//
// # Return Value
//
// A new string in which the characters in `range` of the receiver are
// replaced by `replacement`.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/replacingCharacters(in:with:)
func (s NSString) StringByReplacingCharactersInRangeWithString(range_ NSRange, replacement string) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByReplacingCharactersInRange:withString:"), range_, objc.String(replacement))
	return NSStringFromID(rv).String()
}

// Returns a string containing characters the receiver and a given string have
// in common, starting from the beginning of each up to the first characters
// that aren’t equivalent.
//
// str: The string with which to compare the receiver.
//
// mask: Options for the comparison. The following search options may be specified
// by combining them with the C bitwise [OR] operator:
// [NSCaseInsensitiveSearch], [NSLiteralSearch]. See [String Programming
// Guide] for details on these options.
//
// # Return Value
//
// A string containing characters the receiver and `aString` have in common,
// starting from the beginning of each up to the first characters that
// aren’t equivalent.
//
// # Discussion
//
// The returned string is based on the characters of the receiver. For
// example, if the receiver is “Ma¨dchen” and `aString` is
// “Mädchenschule”, the string returned is “Ma¨dchen”, not
// “Mädchen”.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/commonPrefix(with:options:)
//
// [String Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/introStrings.html#//apple_ref/doc/uid/10000035i
func (s NSString) CommonPrefixWithStringOptions(str string, mask NSStringCompareOptions) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("commonPrefixWithString:options:"), objc.String(str), mask)
	return NSStringFromID(rv).String()
}

// Returns by reference the beginning of the first line and the end of the
// last line touched by the given range.
//
// startPtr: Upon return, contains the index of the first character of the line
// containing the beginning of `aRange`. Pass [NULL] if you do not need this
// value (in which case the work to compute the value isn’t performed).
//
// lineEndPtr: Upon return, contains the index of the first character past the terminator
// of the line containing the end of `aRange`. Pass [NULL] if you do not need
// this value (in which case the work to compute the value isn’t performed).
//
// contentsEndPtr: Upon return, contains the index of the first character of the terminator of
// the line containing the end of `aRange`. Pass [NULL] if you do not need
// this value (in which case the work to compute the value isn’t performed).
//
// range: A range within the receiver. The value must not exceed the bounds of the
// receiver.
//
// Raises an [NSRangeException] if `aRange` is invalid.
//
// # Discussion
//
// A line is delimited by any of these characters, the longest possible
// sequence being preferred to any shorter:
//
// - `U+000A` Unicode Character ‘LINE FEED (LF)’ (`\n`) - `U+000D` Unicode
// Character ‘CARRIAGE RETURN (CR)’ (`\r`) - `U+0085` Unicode Character
// ‘NEXT LINE (NEL)’ - `U+2028` Unicode Character ‘LINE SEPARATOR’ -
// `U+2029` Unicode Character ‘PARAGRAPH SEPARATOR’ - `\r\n`, in that
// order (also known as [CRLF])
//
// If `aRange` is contained with a single line, of course, the returned
// indexes all belong to that line. You can use the results of this method to
// construct ranges for lines by using the start index as the range’s
// location and the difference between the end index and the start index as
// the range’s length.
//
// # Special Considerations
//
// This method detects all invalid ranges (including those with negative
// lengths). For applications linked against macOS 10.6 and later, this error
// causes an exception; for applications linked against earlier releases, this
// error causes a warning, which is displayed just once per application
// execution.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/getLineStart(_:end:contentsEnd:for:)
func (s NSString) GetLineStartEndContentsEndForRange(startPtr unsafe.Pointer, lineEndPtr unsafe.Pointer, contentsEndPtr unsafe.Pointer, range_ NSRange) {
	objc.Send[objc.ID](s.ID, objc.Sel("getLineStart:end:contentsEnd:forRange:"), startPtr, lineEndPtr, contentsEndPtr, range_)
}

// Returns the range of characters representing the line or lines containing a
// given range.
//
// range: A range within the receiver. The value must not exceed the bounds of the
// receiver.
//
// # Return Value
//
// The range of characters representing the line or lines containing `aRange`,
// including the line termination characters. See
// [GetLineStartEndContentsEndForRange] for a discussion of line terminators.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/lineRange(for:)
func (s NSString) LineRangeForRange(range_ NSRange) NSRange {
	rv := objc.Send[NSRange](s.ID, objc.Sel("lineRangeForRange:"), range_)
	return NSRange(rv)
}

// Returns by reference the beginning of the first paragraph and the end of
// the last paragraph touched by the given range.
//
// startPtr: Upon return, contains the index of the first character of the paragraph
// containing the beginning of `aRange`. Pass [NULL] if you do not need this
// value (in which case the work to compute the value isn’t performed).
//
// parEndPtr: Upon return, contains the index of the first character past the terminator
// of the paragraph containing the end of `aRange`. Pass [NULL] if you do not
// need this value (in which case the work to compute the value isn’t
// performed).
//
// contentsEndPtr: Upon return, contains the index of the first character of the terminator of
// the paragraph containing the end of `aRange`. Pass [NULL] if you do not
// need this value (in which case the work to compute the value isn’t
// performed).
//
// range: A range within the receiver. The value must not exceed the bounds of the
// receiver.
//
// # Discussion
//
// A paragraph is any segment of text delimited by a carriage return
// (`U+000D`), newline (`U+000A`), or paragraph separator (`U+2029`).
//
// If `aRange` is contained with a single paragraph, of course, the returned
// indexes all belong to that paragraph. Similar to
// [GetLineStartEndContentsEndForRange], you can use the results of this
// method to construct the ranges for paragraphs.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/getParagraphStart(_:end:contentsEnd:for:)
func (s NSString) GetParagraphStartEndContentsEndForRange(startPtr unsafe.Pointer, parEndPtr unsafe.Pointer, contentsEndPtr unsafe.Pointer, range_ NSRange) {
	objc.Send[objc.ID](s.ID, objc.Sel("getParagraphStart:end:contentsEnd:forRange:"), startPtr, parEndPtr, contentsEndPtr, range_)
}

// Returns the range of characters representing the paragraph or paragraphs
// containing a given range.
//
// range: A range within the receiver. The range must not exceed the bounds of the
// receiver.
//
// # Return Value
//
// The range of characters representing the paragraph or paragraphs containing
// `aRange`, including the paragraph termination characters.
//
// # Discussion
//
// A paragraph is any segment of text delimited by a carriage return
// (`U+000D`), newline (`U+000A`), or paragraph separator (`U+2029`).
//
// See: https://developer.apple.com/documentation/Foundation/NSString/paragraphRange(for:)
func (s NSString) ParagraphRangeForRange(range_ NSRange) NSRange {
	rv := objc.Send[NSRange](s.ID, objc.Sel("paragraphRangeForRange:"), range_)
	return NSRange(rv)
}

// Returns the range in the receiver of the composed character sequence
// located at a given index.
//
// index: The index of a character in the receiver. The value must not exceed the
// bounds of the receiver.
//
// # Return Value
//
// The range in the receiver of the composed character sequence located at
// `anIndex`.
//
// # Discussion
//
// The composed character sequence includes the first decomposed base letter
// found at or before `anIndex`, and its length includes the decomposed base
// letter and all combining characters that follow.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/rangeOfComposedCharacterSequence(at:)
func (s NSString) RangeOfComposedCharacterSequenceAtIndex(index uint) NSRange {
	rv := objc.Send[NSRange](s.ID, objc.Sel("rangeOfComposedCharacterSequenceAtIndex:"), index)
	return NSRange(rv)
}

// Returns the range in the string of the composed character sequences for a
// given range.
//
// range: A range in the receiver. The range must not exceed the bounds of the
// receiver.
//
// # Return Value
//
// The range in the receiver that includes the composed character sequences in
// `range`.
//
// # Discussion
//
// This method provides a convenient way to grow a range to include all
// composed character sequences it overlaps.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/rangeOfComposedCharacterSequences(for:)
func (s NSString) RangeOfComposedCharacterSequencesForRange(range_ NSRange) NSRange {
	rv := objc.Send[NSRange](s.ID, objc.Sel("rangeOfComposedCharacterSequencesForRange:"), range_)
	return NSRange(rv)
}

// Writes the contents of the receiver to a file at a given path using a given
// encoding.
//
// path: The file to which to write the receiver. If `path` contains a tilde (`~`)
// character, you must expand it with [StringByExpandingTildeInPath] before
// invoking this method.
//
// useAuxiliaryFile: If true, the receiver is written to an auxiliary file, and then the
// auxiliary file is renamed to `path`. If false, the receiver is written
// directly to `path`. The true option guarantees that `path`, if it exists at
// all, won’t be corrupted even if the system should crash during writing.
//
// enc: The encoding to use for the output. For possible values, see
// [NSStringEncoding].
//
// # Discussion
//
// This method overwrites any existing file at `path`.
//
// This method stores the specified encoding with the file in an extended
// attribute under the name `com.Apple().TextEncoding`. The value contains the
// IANA name for the encoding and the [CFStringEncoding] value for the
// encoding, separated by a semicolon. The [CFStringEncoding] value is written
// as an ASCII string containing an unsigned 32-bit decimal integer and is not
// terminated by a null character. One or both of these values may be missing.
// Examples of the value written include the following:
//
// - `MACINTOSH;0`
// - `UTF-8;134217984`
// - `UTF-8;`
// - `;3071`
//
// The methods [InitWithContentsOfFileUsedEncodingError],
// `NSString/init()-2c72d`, [StringWithContentsOfFileUsedEncodingError], and
// `NSString/init()-9jrum` use this information to open the file using the
// right encoding.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/write(toFile:atomically:encoding:)
//
// [CFStringEncoding]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncoding
func (s NSString) WriteToFileAtomicallyEncodingError(path string, useAuxiliaryFile bool, enc uint) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("writeToFile:atomically:encoding:error:"), objc.String(path), useAuxiliaryFile, enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeToFile:atomically:encoding:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Writes the contents of the receiver to the URL specified by `url` using the
// specified encoding.
//
// url: The URL to which to write the receiver. Only file URLs are supported.
//
// useAuxiliaryFile: If true, the receiver is written to an auxiliary file, and then the
// auxiliary file is renamed to `url`. If false, the receiver is written
// directly to `url`. The true option guarantees that `url`, if it exists at
// all, won’t be corrupted even if the system should crash during writing.
//
// The `useAuxiliaryFile` parameter is ignored if `url` is not of a type that
// can be accessed atomically.
//
// enc: The encoding to use for the output.
//
// # Discussion
//
// This method stores the specified encoding with the file in an extended
// attribute under the name `com.Apple().TextEncoding`. The value contains the
// IANA name for the encoding and the [CFStringEncoding] value for the
// encoding, separated by a semicolon. The [CFStringEncoding] value is written
// as an ASCII string containing an unsigned 32-bit decimal integer and is not
// terminated by a null character. One or both of these values may be missing.
// Examples of the value written include the following:
//
// - `MACINTOSH;0`
// - UTF-8;134217984
// - `UTF-8;`
// - `;3071`
//
// The methods [InitWithContentsOfFileUsedEncodingError],
// `NSString/init()-2c72d`, [StringWithContentsOfFileUsedEncodingError], and
// `NSString/init()-9jrum` use this information to open the file using the
// right encoding.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/write(to:atomically:encoding:)
//
// [CFStringEncoding]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncoding
func (s NSString) WriteToURLAtomicallyEncodingError(url INSURL, useAuxiliaryFile bool, enc uint) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("writeToURL:atomically:encoding:error:"), url, useAuxiliaryFile, enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeToURL:atomically:encoding:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Parses the receiver as a text representation of a property list, returning
// an [NSString], [NSData], [NSArray], or [NSDictionary] object, according to
// the topmost element.
//
// # Return Value
//
// A property list representation of returning an [NSString], [NSData],
// [NSArray], or [NSDictionary] object, according to the topmost element.
//
// # Discussion
//
// The receiver must contain a string in a property list format. For a
// discussion of property list formats, see [Property List Programming Guide].
//
// See: https://developer.apple.com/documentation/Foundation/NSString/propertyList()
//
// [Property List Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/PropertyLists/Introduction/Introduction.html#//apple_ref/doc/uid/10000048i
func (s NSString) PropertyList() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("propertyList"))
	return objectivec.Object{ID: rv}
}

// Returns a dictionary object initialized with the keys and values found in
// the receiver.
//
// # Return Value
//
// A dictionary object initialized with the keys and values found in the
// receiver
//
// # Discussion
//
// The receiver must contain text in the format used for
// `XCUIElementTypeStrings` files. In this format, keys and values are
// separated by an equal sign, and each key-value pair is terminated with a
// semicolon. The value is optional—if not present, the equal sign is also
// omitted. The keys and values themselves are always strings enclosed in
// straight quotation marks. Comments may be included, delimited by `/*` and
// `*/` as for ANSI C comments. Here’s a short example of a strings file:
//
// See: https://developer.apple.com/documentation/Foundation/NSString/propertyListFromStringsFileFormat()
func (s NSString) PropertyListFromStringsFileFormat() INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("propertyListFromStringsFileFormat"))
	return NSDictionaryFromID(rv)
}

// Draws the receiver with the font and other display characteristics of the
// given attributes, at the specified point in the current graphics context.
//
// point: The point in the current graphics context where you want to start drawing
// the string. The coordinate system of the graphics context is usually
// defined by the view in which you are drawing. In AppKit, the origin is
// normally in the lower-left corner of the drawing area, but the origin is in
// the upper-left corner if the focused view is flipped.
//
// attrs: A dictionary of text attributes to be applied to the string. These are the
// same attributes that can be applied to an [NSAttributedString] object, but
// in the case of [NSString] objects, the attributes apply to the entire
// string, rather than ranges within the string.
//
// # Discussion
//
// The width (height for vertical layout) of the rendering area is unlimited,
// unlike [DrawInRectWithAttributes], which uses a bounding rectangle. As a
// result, this method renders the text in a single line. However, if newline
// characters are present in the string, those characters are honored and
// cause subsequent text to be placed on the next line underneath the starting
// point.
//
// There must be either a focused view or an active graphics context when you
// call this method.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/draw(at:withAttributes:)
func (s NSString) DrawAtPointWithAttributes(point corefoundation.CGPoint, attrs INSDictionary) {
	objc.Send[objc.ID](s.ID, objc.Sel("drawAtPoint:withAttributes:"), point, attrs)
}

// Draws the attributed string inside the specified bounding rectangle.
//
// rect: The bounding rectangle in which to draw the string. In AppKit, the origin
// of the bounding box is normally in the lower-left corner, but the origin is
// in the upper-left corner if the focused view is flipped.
//
// attrs: The text attributes with which to draw the string. These are the same
// attributes that can be applied to an [NSAttributedString] object, but in
// the case of [NSString] objects, the attributes apply to the entire string,
// rather than ranges within the string.
//
// # Discussion
//
// This method draws as much of the string as it can inside the specified
// rectangle, wrapping the string text as needed to make it fit. If the string
// is too long to fit inside the rectangle, the method renders as much as
// possible and clips the rest.
//
// If newline characters are present in the string, those characters are
// honored and cause subsequent text to be placed on the next line underneath
// the starting point.
//
// There must be either a focused view or an active graphics context when you
// call this method.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/draw(in:withAttributes:)
func (s NSString) DrawInRectWithAttributes(rect corefoundation.CGRect, attrs INSDictionary) {
	objc.Send[objc.ID](s.ID, objc.Sel("drawInRect:withAttributes:"), rect, attrs)
}

// Draws the attributed string in the specified bounding rectangle using the
// provided options.
//
// rect: The bounding rectangle in which to draw the string.
//
// options: Additional drawing options to apply to the string during rendering. For a
// list of possible values, see [NSStringDrawingOptions].
//
// attributes: The text attributes with which to draw the string. These are the same
// attributes that can be applied to an [NSAttributedString] object, but in
// the case of [NSString] objects, the attributes apply to the entire string,
// rather than ranges within the string.
//
// context: A context object with information about how to adjust the font tracking and
// scaling information. On return, the specified object contains information
// about the actual values used to render the string. This parameter may be
// `nil`.
//
// context is a [*appkit.NSStringDrawingContext].
//
// # Discussion
//
// This method draws as much of the string as it can inside the specified
// rectangle, wrapping the string text as needed to make it fit. If the string
// is too big to fit completely inside the rectangle, the method scales the
// font or adjusts the letter spacing to make the string fit within the given
// bounds.
//
// If newline characters are present in the string, those characters are
// honored and cause subsequent text to be placed on the next line underneath
// the starting point. To correctly draw and size multi-line text, pass
// [usesLineFragmentOrigin] in the options parameter.
//
// # Special Considerations
//
// This method uses the baseline origin by default.
//
// If [usesLineFragmentOrigin] is not specified, the rectangle’s height will
// be ignored and the operation considered to be single-line rendering.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/draw(with:options:attributes:context:)
//
// [NSStringDrawingOptions]: https://developer.apple.com/documentation/UIKit/NSStringDrawingOptions
// [usesLineFragmentOrigin]: https://developer.apple.com/documentation/UIKit/NSStringDrawingOptions/usesLineFragmentOrigin
func (s NSString) DrawWithRectOptionsAttributesContext(rect corefoundation.CGRect, options NSStringDrawingOptions, attributes INSDictionary, context objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("drawWithRect:options:attributes:context:"), rect, options, attributes, context)
}

// Calculates and returns the bounding rect for the receiver drawn using the
// given options and display characteristics, within the specified rectangle
// in the current graphics context.
//
// size: The size of the rectangle to draw in.
//
// options: String drawing options.
//
// attributes: A dictionary of text attributes to be applied to the string. These are the
// same attributes that can be applied to an [NSAttributedString] object, but
// in the case of [NSString] objects, the attributes apply to the entire
// string, rather than ranges within the string.
//
// context: The string drawing context to use for the receiver, specifying minimum
// scale factor and tracking adjustments.
//
// context is a [*appkit.NSStringDrawingContext].
//
// # Return Value
//
// The bounding rect for the receiver drawn using the given options and
// display characteristics. The rect origin returned from this method is the
// first glyph origin.
//
// # Discussion
//
// To correctly draw and size multi-line text, pass [usesLineFragmentOrigin]
// in the options parameter.
//
// This method returns fractional sizes (in the `size` component of the
// returned [CGRect]); to use a returned size to size views, you must raise
// its value to the nearest higher integer using the [ceil] function.
//
// This method returns the actual bounds of the glyphs in the string. Some of
// the glyphs (spaces, for example) are allowed to overlap the layout
// constraints specified by the size passed in, so in some cases the width
// value of the size component of the returned [CGRect] can exceed the width
// value of the `size` parameter.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/boundingRect(with:options:attributes:context:)
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
// [ceil]: https://developer.apple.com/documentation/kernel/1557272-ceil
// [usesLineFragmentOrigin]: https://developer.apple.com/documentation/UIKit/NSStringDrawingOptions/usesLineFragmentOrigin
func (s NSString) BoundingRectWithSizeOptionsAttributesContext(size corefoundation.CGSize, options NSStringDrawingOptions, attributes INSDictionary, context objectivec.IObject) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("boundingRectWithSize:options:attributes:context:"), size, options, attributes, context)
	return corefoundation.CGRect(rv)
}

// Returns the bounding box size the receiver occupies when drawn with the
// given attributes.
//
// attrs: A dictionary of text attributes to be applied to the string. These are the
// same attributes that can be applied to an [NSAttributedString] object, but
// in the case of [NSString] objects, the attributes apply to the entire
// string, rather than ranges within the string.
//
// # Return Value
//
// The bounding box size the receiver occupies when drawn with the specified
// attributes.
//
// # Discussion
//
// This method returns fractional sizes; to use a returned size to size views,
// you must raise its value to the nearest higher integer using the [ceil]
// function.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/size(withAttributes:)
//
// [ceil]: https://developer.apple.com/documentation/kernel/1557272-ceil
func (s NSString) SizeWithAttributes(attrs INSDictionary) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s.ID, objc.Sel("sizeWithAttributes:"), attrs)
	return corefoundation.CGSize(rv)
}

// Returns a string variation suitable for the specified presentation width.
//
// width: The desired width of the string variation.
//
// # Return Value
//
// A string variation, or the original string if no variations exist for the
// specified width.
//
// # Discussion
//
// You can use this method to provide adaptive strings for the user’s
// device—that is, text that avoids truncation and maximizes available
// space. For example, an app running on an iPad Pro in Landscape orientation
// might welcome a user with the message “Greetings and Salutations!”,
// whereas the same app running on an iPhone SE in Portrait orientation might
// instead show an abbreviated welcome message, like “Hello!”.
//
// Call this method on a string with one or more width variations. You define
// width variations for a localized string in a Stringsdict file using the
// [NSStringVariableWidthRuleType] key, and then retrieve a string with
// variations using [NSLocalizedString].
//
// For example, consider the `stringsdict.Plist()` file in the first listing
// and the corresponding code in the second listing.
//
// # Understanding How Width Variants Are Selected
//
// This method selects a variation for a specified width according to the
// following behavior:
//
// - If no variations exist for the string, the original string is returned. -
// If a variation exists for the specified width, that string is returned. -
// If no variation is found with a width less than the specified value, the
// variation with the smallest width is returned. - Otherwise, the variation
// with the next smallest width value is returned.
//
// # Specifying Width Values
//
// Smaller width values correspond to shorter strings, whereas larger values
// correspond to longer strings. By default, width values do not have an
// associated unit.
//
// In iOS, width contexts are measured by the number of em units that fit
// across the app window. This value depends on several factors, including the
// size and orientation of the device, the user’s preferred content size
// category, whether the Zoom accessibility feature is enabled, and whether
// the app is displayed in a multitasking context, such as in Split View or
// Slide Over.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/variantFittingPresentationWidth(_:)
//
// [NSLocalizedString]: https://developer.apple.com/documentation/Foundation/NSLocalizedString
func (s NSString) VariantFittingPresentationWidth(width int) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("variantFittingPresentationWidth:"), width)
	return NSStringFromID(rv).String()
}

// Returns a Boolean value that indicates whether the receiver can be
// converted to a given encoding without loss of information.
//
// encoding: A string encoding. For possible values, see [NSStringEncoding].
//
// # Return Value
//
// true if the receiver can be converted to `encoding` without loss of
// information. Returns false if characters would have to be changed or
// deleted in the process of changing encodings.
//
// # Discussion
//
// If you plan to actually convert a string, the `...` methods return `nil` on
// failure, so you can avoid the overhead of invoking this method yourself by
// simply trying to convert the string.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/canBeConverted(to:)
func (s NSString) CanBeConvertedToEncoding(encoding uint) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("canBeConvertedToEncoding:"), encoding)
	return rv
}

// Returns an [NSData] object containing a representation of the receiver
// encoded using a given encoding.
//
// encoding: A string encoding. For possible values, see [NSStringEncoding].
//
// # Return Value
//
// The result of invoking [DataUsingEncodingAllowLossyConversion] with false
// as the second argument (that is, requiring lossless conversion).
//
// See: https://developer.apple.com/documentation/Foundation/NSString/data(using:)
func (s NSString) DataUsingEncoding(encoding uint) INSData {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("dataUsingEncoding:"), encoding)
	return NSDataFromID(rv)
}

// Returns an [NSData] object containing a representation of the receiver
// encoded using a given encoding.
//
// encoding: A string encoding. For possible values, see [NSStringEncoding].
//
// lossy: If true, then allows characters to be removed or altered in conversion.
//
// # Return Value
//
// An [NSData] object containing a representation of the receiver encoded
// using `encoding`. Returns `nil` if `flag` is false and the receiver can’t
// be converted without losing some information (such as accents or case).
//
// # Discussion
//
// If `flag` is true and the receiver can’t be converted without losing some
// information, some characters may be removed or altered in conversion. For
// example, in converting a character from [NSUnicodeStringEncoding] to
// [NSASCIIStringEncoding], the character ‘Á’ becomes ‘A’, losing the
// accent.
//
// This method creates an external representation (with a byte order marker,
// if necessary, to indicate endianness) to ensure that the resulting [NSData]
// object can be written out to a file safely. The result of this method, when
// lossless conversion is made, is the default “plain text” format for
// encoding and is the recommended way to save or transmit a string object.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/data(using:allowLossyConversion:)
func (s NSString) DataUsingEncodingAllowLossyConversion(encoding uint, lossy bool) INSData {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("dataUsingEncoding:allowLossyConversion:"), encoding, lossy)
	return NSDataFromID(rv)
}

// Interprets the receiver as a path in the file system and attempts to
// perform filename completion, returning a numeric value that indicates
// whether a match was possible, and by reference the longest path that
// matches the receiver.
//
// outputName: Upon return, contains the longest path that matches the receiver.
//
// flag: If true, the method considers case for possible completions.
//
// outputArray: Upon return, contains all matching filenames.
//
// filterTypes: An array of [NSString] objects specifying path extensions to consider for
// completion. Only paths whose extensions (not including the extension
// separator) match one of these strings are included in `outputArray`. Pass
// `nil` if you don’t want to filter the output.
//
// # Return Value
//
// `0` if no matches are found and `1` if exactly one match is found. In the
// case of multiple matches, returns the actual number of matching paths if
// `outputArray` is provided, or simply a positive value if `outputArray` is
// [NULL].
//
// # Discussion
//
// You can check for the existence of matches without retrieving by passing
// [NULL] as `outputArray`.
//
// Note that this method only works with file paths (not, for example, string
// representations of URLs).
//
// See: https://developer.apple.com/documentation/Foundation/NSString/completePath(into:caseSensitive:matchesInto:filterTypes:)
func (s NSString) CompletePathIntoStringCaseSensitiveMatchesIntoArrayFilterTypes(outputName string, flag bool, outputArray []string, filterTypes []string) uint {
	rv := objc.Send[uint](s.ID, objc.Sel("completePathIntoString:caseSensitive:matchesIntoArray:filterTypes:"), objc.String(outputName), flag, objectivec.StringSliceToNSArray(outputArray), objectivec.StringSliceToNSArray(filterTypes))
	return rv
}

// Interprets the receiver as a system-independent path and fills a buffer
// with a C-string in a format and encoding suitable for use with file-system
// calls.
//
// cname: Upon return, contains a C-string that represent the receiver as a
// system-independent path, plus the [NULL] termination byte. The size of
// `buffer` must be large enough to contain `maxLength` bytes.
//
// max: The maximum number of bytes in the string to return in `buffer` (including
// a terminating [NULL] character, which this method adds).
//
// # Return Value
//
// true if `buffer` is successfully filled with a file-system representation,
// otherwise false (for example, if `maxLength` would be exceeded or if the
// receiver can’t be represented in the file system’s encoding).
//
// # Discussion
//
// This method operates by replacing the abstract path and extension separator
// characters (’`/`’ and ‘`.`’ respectively) with their equivalents
// for the operating system. If the system-specific path or extension
// separator appears in the abstract representation, the characters it is
// converted to depend on the system (unless they’re identical to the
// abstract separators).
//
// Note that this method only works with file paths (not, for example, string
// representations of URLs).
//
// The following example illustrates the use of the `maxLength` argument. The
// first method invocation returns failure as the file representation of the
// string (`@"/mach_kernel"`) is 12 bytes long and the value passed as the
// `maxLength` argument (`12`) does not allow for the addition of a [NULL]
// termination byte.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/getFileSystemRepresentation(_:maxLength:)
func (s NSString) GetFileSystemRepresentationMaxLength(cname string, max uint) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("getFileSystemRepresentation:maxLength:"), unsafe.Pointer(unsafe.StringData(cname+"\x00")), max)
	return rv
}

// Returns a new string made by appending to the receiver a given string.
//
// str: The path component to append to the receiver.
//
// # Return Value
//
// A new string made by appending `aString` to the receiver, preceded if
// necessary by a path separator.
//
// # Discussion
//
// The following table illustrates the effect of this method on a variety of
// different paths, assuming that `aString` is supplied as
// “`scratch.Tiff()`”:
//
// [Table data omitted]
//
// Note that this method only works with file paths (not, for example, string
// representations of URLs).
//
// See: https://developer.apple.com/documentation/Foundation/NSString/appendingPathComponent(_:)
func (s NSString) StringByAppendingPathComponent(str string) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByAppendingPathComponent:"), objc.String(str))
	return NSStringFromID(rv).String()
}

// Returns a new string made by appending to the receiver an extension
// separator followed by a given extension.
//
// str: The extension to append to the receiver.
//
// # Return Value
//
// A new string made by appending to the receiver an extension separator
// followed by `ext`.
//
// # Discussion
//
// The following table illustrates the effect of this method on a variety of
// different paths, assuming that `ext` is supplied as `@"tiff"`:
//
// [Table data omitted]
//
// Note that adding an extension to `@"/tmp/"` causes the result to be
// `@"/tmp.Tiff()"` instead of `@"/tmp/XCUIElementTypeTiff"`. This difference
// is because a file named `@"XCUIElementTypeTiff"` is not considered to have
// an extension, so the string is appended to the last nonempty path
// component.
//
// Note that this method only works with file paths (not, for example, string
// representations of URLs).
//
// # Special Considerations
//
// Prior to OS X v10.9 this method did not allow you to append file extensions
// to filenames starting with the tilde character (`~`).
//
// See: https://developer.apple.com/documentation/Foundation/NSString/appendingPathExtension(_:)
func (s NSString) StringByAppendingPathExtension(str string) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByAppendingPathExtension:"), objc.String(str))
	return NSStringFromID(rv).String()
}

// Returns an array of strings made by separately appending to the receiver
// each string in a given array.
//
// paths: An array of [NSString] objects specifying paths to add to the receiver.
//
// # Return Value
//
// An array of string objects made by separately appending each string in
// `paths` to the receiver, preceded if necessary by a path separator.
//
// # Discussion
//
// Note that this method only works with file paths (not, for example, string
// representations of URLs). See [StringByAppendingPathComponent] for an
// individual example.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/strings(byAppendingPaths:)
func (s NSString) StringsByAppendingPaths(paths []string) []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("stringsByAppendingPaths:"), objectivec.StringSliceToNSArray(paths))
	return objc.ConvertSliceToStrings(rv)
}

// Returns a new string made from the receiver by replacing all characters not
// in the specified set with percent-encoded characters.
//
// allowedCharacters: The characters not replaced in the string. Typically, you specify one of
// the predefined character sets for a particular URL component, such as
// [URLPathAllowedCharacterSet] or [URLQueryAllowedCharacterSet].
//
// # Return Value
//
// Returns the encoded string, or `nil` if the transformation is not possible.
//
// # Discussion
//
// Entire URL strings cannot be percent-encoded, because each URL component
// specifies a different set of allowed characters. For example, the query
// component of a URL allows the “`@`” character, but that character must
// be percent-encoded in the password component.
//
// UTF-8 encoding is used to determine the correct percent-encoded characters.
// Any characters in `allowedCharacters` outside of the 7-bit ASCII range are
// ignored.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/addingPercentEncoding(withAllowedCharacters:)
func (s NSString) StringByAddingPercentEncodingWithAllowedCharacters(allowedCharacters INSCharacterSet) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByAddingPercentEncodingWithAllowedCharacters:"), allowedCharacters)
	return NSStringFromID(rv).String()
}

// Copies all characters from the receiver into a given buffer.
//
// buffer: Upon return, contains the characters from the receiver. `buffer` must be
// large enough to contain all characters in the string (`[string
// length]*sizeof(unichar)`).
//
// # Discussion
//
// Invokes [GetCharactersRange] with `buffer` and the entire extent of the
// receiver as the range.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/getCharacters(_:)
func (s NSString) GetCharacters(buffer unsafe.Pointer) {
	objc.Send[objc.ID](s.ID, objc.Sel("getCharacters:"), buffer)
}

// Draws the receiver with the specified options and other display
// characteristics of the given attributes, within the specified rectangle in
// the current graphics context.
//
// rect: The rectangle in which to draw the string.
//
// options: String drawing options.
//
// attributes: A dictionary of text attributes to be applied to the string. These are the
// same attributes that can be applied to an [NSAttributedString] object, but
// in the case of [NSString] objects, the attributes apply to the entire
// string, rather than ranges within the string.
//
// # Discussion
//
// This method works in single-line, baseline rendering configuration by
// default. That is, the `rect` argument’s `origin` field specifies the
// rendering origin, and that point is interpreted as the baseline origin by
// default. If the string drawing option
// [NSStringDrawingUsesLineFragmentOrigin] is specified, `origin` is
// interpreted as the upper left corner of the line fragment rectangle, and
// the method behaves in multiline configuration.
//
// The `size` field specifies the text container size. The `width` part of the
// size field specifies the maximum line fragment width if larger than `0.0`.
// The `height` defines the maximum size that can be occupied with text if
// larger than `0.0` and [NSStringDrawingUsesLineFragmentOrigin] is specified.
// If [NSStringDrawingUsesLineFragmentOrigin] is not specified, height is
// ignored and considered to be single-line rendering
// ([NSLineBreakByWordWrapping] and [NSLineBreakByCharWrapping] are treated as
// [NSLineBreakByClipping]).
//
// You should only invoke this method when there is a current graphics
// context.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/draw(with:options:attributes:)
func (s NSString) DrawWithRectOptionsAttributes(rect corefoundation.CGRect, options NSStringDrawingOptions, attributes INSDictionary) {
	objc.Send[objc.ID](s.ID, objc.Sel("drawWithRect:options:attributes:"), rect, options, attributes)
}

// Calculates and returns the bounding rect for the receiver drawn using the
// given options and display characteristics, within the specified rectangle
// in the current graphics context.
//
// size: The size of the rectangle to draw in.
//
// options: String drawing options.
//
// attributes: A dictionary of text attributes to be applied to the string. These are the
// same attributes that can be applied to an [NSAttributedString] object, but
// in the case of [NSString] objects, the attributes apply to the entire
// string, rather than ranges within the string.
//
// # Return Value
//
// The bounding rect for the receiver drawn using the given options and
// display characteristics. The rect origin returned from this method is the
// first glyph origin.
//
// # Discussion
//
// This method works in single-line, baseline rendering configuration by
// default. If the string drawing option
// [NSStringDrawingUsesLineFragmentOrigin] is specified, the method behaves in
// multiline configuration.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/boundingRect(with:options:attributes:)
func (s NSString) BoundingRectWithSizeOptionsAttributes(size corefoundation.CGSize, options NSStringDrawingOptions, attributes INSDictionary) NSRect {
	rv := objc.Send[NSRect](s.ID, objc.Sel("boundingRectWithSize:options:attributes:"), size, options, attributes)
	return NSRect(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSString/init(bytesNoCopy:length:encoding:deallocator:)
func (s NSString) InitWithBytesNoCopyLengthEncodingDeallocator(bytes unsafe.Pointer, len_ uint, encoding uint, deallocator func(unsafe.Pointer, uint64)) NSString {
	_block3 := objc.NewBlock(func(_ objc.Block, arg0 unsafe.Pointer, arg1 uint64) { deallocator(arg0, arg1) })
	defer _block3.Release()
	rv := objc.Send[NSString](s.ID, objc.Sel("initWithBytesNoCopy:length:encoding:deallocator:"), bytes, len_, encoding, objc.ID(_block3))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSString/init(cString:encoding:)
func (s NSString) InitWithCStringEncoding(nullTerminatedCString string, encoding uint) NSString {
	rv := objc.Send[NSString](s.ID, objc.Sel("initWithCString:encoding:"), unsafe.Pointer(unsafe.StringData(nullTerminatedCString+"\x00")), encoding)
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSString/init(coder:)
func (s NSString) InitWithCoder(coder INSCoder) NSString {
	rv := objc.Send[NSString](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOf:encoding:)
func (s NSString) InitWithContentsOfURLEncodingError(url INSURL, enc uint) (NSString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithContentsOfURL:encoding:error:"), url, enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOf:usedEncoding:)
func (s NSString) InitWithContentsOfURLUsedEncodingError(url INSURL, enc unsafe.Pointer) (NSString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithContentsOfURL:usedEncoding:error:"), url, enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Foundation/NSString/init(utf8String:)
func (s NSString) InitWithUTF8String(nullTerminatedCString string) NSString {
	rv := objc.Send[NSString](s.ID, objc.Sel("initWithUTF8String:"), unsafe.Pointer(unsafe.StringData(nullTerminatedCString+"\x00")))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSString/appendingPathComponent(_:conformingTo:)
func (s NSString) StringByAppendingPathComponentConformingToType(partialName string, contentType objectivec.IObject) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByAppendingPathComponent:conformingToType:"), objc.String(partialName), contentType)
	return NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Foundation/NSString/appendingPathExtension(for:)
func (s NSString) StringByAppendingPathExtensionForType(contentType objectivec.IObject) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByAppendingPathExtensionForType:"), contentType)
	return NSStringFromID(rv).String()
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (s NSString) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns an [NSString] object initialized by using a given format string as
// a template into which the remaining argument values are substituted.
//
// format: A format string. See [Formatting String Objects] for examples of how to use
// this method, and [String Format Specifiers] for a list of format
// specifiers. This value must not be `nil`.
//
// # Return Value
//
// An [NSString] object initialized by using `format` as a template into which
// the remaining argument values are substituted according to the system
// locale. The returned object may be different from the original receiver.
//
// # Discussion
//
// Pass a comma-separated list of variadic arguments to substitute into
// `format`.
//
// This method invokes [InitWithFormatLocaleArguments] without applying any
// localization. This is useful, for example, when working with fixed-format
// representations of information that is written out and read back in at a
// later time.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/initWithFormat:
//
// [Formatting String Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/FormatStrings.html#//apple_ref/doc/uid/20000943
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
func (s NSString) InitWithFormat(format string) NSString {
	rv := objc.Send[NSString](s.ID, objc.Sel("initWithFormat:"), objc.String(format))
	return rv
}

// Returns an [NSString] object initialized by using a given format string as
// a template into which the remaining argument values are substituted
// according to given locale.
//
// format: A format string. See [Formatting String Objects] for examples of how to use
// this method, and [String Format Specifiers] for a list of format
// specifiers. This value must not be `nil`.
//
// locale: An [NSLocale] object specifying the locale to use. To use the current
// locale, pass `[NSLocale currentLocale]`. To use the system locale, pass
// `nil`.
//
// For legacy support, this may be an instance of [NSDictionary] containing
// locale information.
//
// # Discussion
//
// Pass comma-separated list of trailing variadic arguments to substitute into
// `format`.
//
// Invokes [InitWithFormatLocaleArguments] with `locale` as the locale.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/initWithFormat:locale:
//
// [Formatting String Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/FormatStrings.html#//apple_ref/doc/uid/20000943
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
func (s NSString) InitWithFormatLocale(format string, locale objectivec.IObject) NSString {
	rv := objc.Send[NSString](s.ID, objc.Sel("initWithFormat:locale:"), objc.String(format), locale)
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:arguments:error:
func (s NSString) InitWithValidatedFormatValidFormatSpecifiersArgumentsError(format string, validFormatSpecifiers string, argList unsafe.Pointer) (NSString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:arguments:error:"), objc.String(format), objc.String(validFormatSpecifiers), argList, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:error:
func (s NSString) InitWithValidatedFormatValidFormatSpecifiersError(format string, validFormatSpecifiers string) (NSString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:error:"), objc.String(format), objc.String(validFormatSpecifiers), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:locale:arguments:error:
func (s NSString) InitWithValidatedFormatValidFormatSpecifiersLocaleArgumentsError(format string, validFormatSpecifiers string, locale objectivec.IObject, argList unsafe.Pointer) (NSString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:locale:arguments:error:"), objc.String(format), objc.String(validFormatSpecifiers), locale, argList, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:locale:error:
func (s NSString) InitWithValidatedFormatValidFormatSpecifiersLocaleError(format string, validFormatSpecifiers string, locale objectivec.IObject) (NSString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:locale:error:"), objc.String(format), objc.String(validFormatSpecifiers), locale, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil

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
func (s NSString) ItemProviderVisibilityForRepresentationWithTypeIdentifier(typeIdentifier string) NSItemProviderRepresentationVisibility {
	rv := objc.Send[NSItemProviderRepresentationVisibility](s.ID, objc.Sel("itemProviderVisibilityForRepresentationWithTypeIdentifier:"), objc.String(typeIdentifier))
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
func (s NSString) LoadDataWithTypeIdentifierForItemProviderCompletionHandler(typeIdentifier string, completionHandler DataErrorHandler) INSProgress {
	_block1, _ := NewDataErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](s.ID, objc.Sel("loadDataWithTypeIdentifier:forItemProviderCompletionHandler:"), objc.String(typeIdentifier), _block1)
	return NSProgressFromID(rv)
}

// Returns a string made by appending to the receiver a string constructed
// from a given format string and the following arguments.
//
// format: A format string. See [Formatting String Objects] for more information. This
// value must not be `nil`.
//
// # Return Value
//
// A string made by appending to the receiver a string constructed from
// `format` and the following arguments, in the manner of [StringWithFormat].
//
// # Discussion
//
// Pass a comma-separated list of variadic arguments to substitute into
// `format`.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/stringByAppendingFormat:
//
// [Formatting String Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/FormatStrings.html#//apple_ref/doc/uid/20000943
func (s NSString) StringByAppendingFormat(format string) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByAppendingFormat:"), objc.String(format))
	return NSStringFromID(rv).String()
}

// Returns a localized string intended for display in a notification alert.
//
// key: The key to use when looking up the string in the app’s
// `Localizable.Strings()` file.
//
// arguments: An array of values to substitute for escaped characters in the string.
//
// # Return Value
//
// A string whose value is created dynamically from a localized string
// resource. If a string resource corresponding to the specified `key` cannot
// be found, this method returns `key`.
//
// # Discussion
//
// When configuring the content of a local notification using the User
// Notifications framework, use this method to create strings whose contents
// are stored in your app’s `Localizable.Strings()` file. When the
// notification is about to be displayed, the string object uses the key and
// arguments you specify to load the appropriate localized version of the
// string. If the localized string has any escaped character sequences—that
// is, special characters proceeded by a percent (%) sign—those character
// sequences are replaced by the values in the `arguments` parameter.
//
// For information about how strings are formatted, see [String Resources] in
// [Resource Programming Guide].
//
// See: https://developer.apple.com/documentation/Foundation/NSString/localizedUserNotificationString(forKey:arguments:)
//
// [Resource Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/LoadingResources/Introduction/Introduction.html#//apple_ref/doc/uid/10000051i
// [String Resources]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/LoadingResources/Strings/Strings.html#//apple_ref/doc/uid/10000051i-CH6
func (_NSStringClass NSStringClass) LocalizedUserNotificationStringForKeyArguments(key string, arguments INSArray) string {
	rv := objc.Send[objc.ID](objc.ID(_NSStringClass.class), objc.Sel("localizedUserNotificationStringForKey:arguments:"), objc.String(key), arguments)
	return NSStringFromID(rv).String()
}

// Returns the string encoding for the given data as detected by attempting to
// create a string according to the specified encoding options.
//
// data: An [NSData] object containing bytes in an encoding to be determined.
//
// opts: Options to use when attempting to determine the string encoding. See
// `String Encoding Detection Options` for a full list of supported options.
//
// string: If a string encoding could be determined, upon return contains an
// [NSString] object constructed from data using the determined string
// encoding.
//
// usedLossyConversion: If a string encoding could be determined, upon return contains a [BOOL]
// value corresponding to whether lossy conversion was used.
//
// # Return Value
//
// An [NSStringEncoding] value, or `0` if a string encoding could not be
// determined.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/stringEncoding(for:encodingOptions:convertedString:usedLossyConversion:)
func (_NSStringClass NSStringClass) StringEncodingForDataEncodingOptionsConvertedStringUsedLossyConversion(data INSData, opts INSDictionary, string_ string, usedLossyConversion unsafe.Pointer) NSStringEncoding {
	rv := objc.Send[NSStringEncoding](objc.ID(_NSStringClass.class), objc.Sel("stringEncodingForData:encodingOptions:convertedString:usedLossyConversion:"), data, opts, objc.String(string_), usedLossyConversion)
	return NSStringEncoding(rv)
}

// Returns a human-readable string giving the name of a given encoding.
//
// encoding: A string encoding. For possible values, see [NSStringEncoding].
//
// # Return Value
//
// A human-readable string giving the name of `encoding` in the current
// locale.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/localizedName(of:)
func (_NSStringClass NSStringClass) LocalizedNameOfStringEncoding(encoding uint) string {
	rv := objc.Send[objc.ID](objc.ID(_NSStringClass.class), objc.Sel("localizedNameOfStringEncoding:"), encoding)
	return NSStringFromID(rv).String()
}

// Returns a string built from the strings in a given array by concatenating
// them with a path separator between each pair.
//
// components: An array of [NSString] objects representing a file path. To create an
// absolute path, use a slash mark (”`/`”) as the first component. To
// include a trailing path divider, use an empty string as the last component.
//
// # Return Value
//
// A string built from the strings in `components` by concatenating them (in
// the order they appear in the array) with a path separator between each
// pair.
//
// # Discussion
//
// This method doesn’t clean up the path created; use
// [StringByStandardizingPath] to resolve empty components, references to the
// parent directory, and so on.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/path(withComponents:)
func (_NSStringClass NSStringClass) PathWithComponents(components []string) string {
	rv := objc.Send[objc.ID](objc.ID(_NSStringClass.class), objc.Sel("pathWithComponents:"), objectivec.StringSliceToNSArray(components))
	return NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Foundation/NSString/deferredLocalizedIntentsStringWithFormat:
func (_NSStringClass NSStringClass) DeferredLocalizedIntentsStringWithFormat(format string) string {
	rv := objc.Send[objc.ID](objc.ID(_NSStringClass.class), objc.Sel("deferredLocalizedIntentsStringWithFormat:"), objc.String(format))
	return NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Foundation/NSString/deferredLocalizedIntentsStringWithFormat:fromTable:
func (_NSStringClass NSStringClass) DeferredLocalizedIntentsStringWithFormatFromTable(format string, table string) string {
	rv := objc.Send[objc.ID](objc.ID(_NSStringClass.class), objc.Sel("deferredLocalizedIntentsStringWithFormat:fromTable:"), objc.String(format), objc.String(table))
	return NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Foundation/NSString/deferredLocalizedIntentsStringWithFormat:fromTable:arguments:
func (_NSStringClass NSStringClass) DeferredLocalizedIntentsStringWithFormatFromTableArguments(format string, table string, arguments unsafe.Pointer) string {
	rv := objc.Send[objc.ID](objc.ID(_NSStringClass.class), objc.Sel("deferredLocalizedIntentsStringWithFormat:fromTable:arguments:"), objc.String(format), objc.String(table), arguments)
	return NSStringFromID(rv).String()
}

// Returns a string created by using a given format string as a template into
// which the remaining argument values are substituted according to the
// current locale.
//
// format: A format string. See [Formatting String Objects] for examples of how to use
// this method, and [String Format Specifiers] for a list of format
// specifiers. This value must not be `nil`.
//
// Raises an [NSInvalidArgumentException] if `format` is `nil`.
//
// # Return Value
//
// A string created by using `format` as a template into which the following
// argument values are substituted according to the formatting information in
// the current locale.
//
// # Discussion
//
// Pass a comma-separated list of variadic arguments to substitute into
// `format`.
//
// This method is equivalent to using [InitWithFormatLocale] and passing the
// current locale as the locale argument.
//
// As an example of formatting, this method replaces the decimal according to
// the locale in `%f` and `%d` substitutions, and calls
// [DescriptionWithLocale] instead of [description()] where necessary.
//
// This code excerpt creates a string from another string and a float:
//
// The resulting string has the value “`Cost: 1234.560000\n`” if the
// locale is `en_US`, and “`Cost: 1234,560000\n`” if the locale is
// `fr_FR`.
//
// See [Formatting String Objects] for more information.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/localizedStringWithFormat:
//
// [Formatting String Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/FormatStrings.html#//apple_ref/doc/uid/20000943
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
// [description()]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/description()
//
// [Formatting String Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/FormatStrings.html#//apple_ref/doc/uid/20000943
func (_NSStringClass NSStringClass) LocalizedStringWithFormat(format string) NSString {
	rv := objc.Send[objc.ID](objc.ID(_NSStringClass.class), objc.Sel("localizedStringWithFormat:"), objc.String(format))
	return NSStringFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSString/localizedStringWithValidatedFormat:validFormatSpecifiers:error:
func (_NSStringClass NSStringClass) LocalizedStringWithValidatedFormatValidFormatSpecifiersError(format string, validFormatSpecifiers string) (NSString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSStringClass.class), objc.Sel("localizedStringWithValidatedFormat:validFormatSpecifiers:error:"), objc.String(format), objc.String(validFormatSpecifiers), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil

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
func (_NSStringClass NSStringClass) ObjectWithItemProviderDataTypeIdentifierError(data INSData, typeIdentifier string) (NSString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSStringClass.class), objc.Sel("objectWithItemProviderData:typeIdentifier:error:"), data, objc.String(typeIdentifier), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil

}

// Returns an empty string.
//
// # Return Value
//
// An empty string.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/string
func (_NSStringClass NSStringClass) String() NSString {
	rv := objc.Send[objc.ID](objc.ID(_NSStringClass.class), objc.Sel("string"))
	return NSStringFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSString/stringWithCString:encoding:
func (_NSStringClass NSStringClass) StringWithCStringEncoding(cString string, enc uint) NSString {
	rv := objc.Send[objc.ID](objc.ID(_NSStringClass.class), objc.Sel("stringWithCString:encoding:"), unsafe.Pointer(unsafe.StringData(cString+"\x00")), enc)
	return NSStringFromID(rv)
}

// Returns a string containing a given number of characters taken from a given
// C array of UTF-16 code units.
//
// characters: A C array of UTF-16 code units; the value must not be [NULL].
//
// length: The number of characters to use from `chars`.
//
// # Return Value
//
// A string containing `length` UTF-16 code units taken (starting with the
// first) from `chars`.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/stringWithCharacters:length:
func (_NSStringClass NSStringClass) StringWithCharactersLength(characters unsafe.Pointer, length uint) NSString {
	rv := objc.Send[objc.ID](objc.ID(_NSStringClass.class), objc.Sel("stringWithCharacters:length:"), characters, length)
	return NSStringFromID(rv)
}

// Returns a string created by reading data from the file at a given path
// interpreted using a given encoding.
//
// path: A path to a file.
//
// enc: The encoding of the file at `path`. For possible values, see
// [NSStringEncoding].
//
// error: If an error occurs, upon returns contains an [NSError] object that
// describes the problem. If you are not interested in possible errors, pass
// in [NULL].
//
// # Return Value
//
// A string created by reading data from the file named by `path` using the
// encoding, `enc`. If the file can’t be opened or there is an encoding
// error, returns `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/stringWithContentsOfFile:encoding:error:
func (_NSStringClass NSStringClass) StringWithContentsOfFileEncodingError(path string, enc uint) (NSString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSStringClass.class), objc.Sel("stringWithContentsOfFile:encoding:error:"), objc.String(path), enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil

}

// Returns a string created by reading data from the file at a given path and
// returns by reference the encoding used to interpret the file.
//
// path: A path to a file.
//
// enc: Upon return, if the file is read successfully, contains the encoding used
// to interpret the file at `path`. For possible values, see
// [NSStringEncoding].
//
// error: If an error occurs, upon returns contains an [NSError] object that
// describes the problem. If you are not interested in possible errors, you
// may pass in [NULL].
//
// # Return Value
//
// A string created by reading data from the file named by `path`. If the file
// can’t be opened or there is an encoding error, returns `nil`.
//
// # Discussion
//
// This method attempts to determine the encoding of the file at `path`.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/stringWithContentsOfFile:usedEncoding:error:
func (_NSStringClass NSStringClass) StringWithContentsOfFileUsedEncodingError(path string, enc unsafe.Pointer) (NSString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSStringClass.class), objc.Sel("stringWithContentsOfFile:usedEncoding:error:"), objc.String(path), enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Foundation/NSString/stringWithContentsOfURL:encoding:error:
func (_NSStringClass NSStringClass) StringWithContentsOfURLEncodingError(url INSURL, enc uint) (NSString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSStringClass.class), objc.Sel("stringWithContentsOfURL:encoding:error:"), url, enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Foundation/NSString/stringWithContentsOfURL:usedEncoding:error:
func (_NSStringClass NSStringClass) StringWithContentsOfURLUsedEncodingError(url INSURL, enc unsafe.Pointer) (NSString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSStringClass.class), objc.Sel("stringWithContentsOfURL:usedEncoding:error:"), url, enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil

}

// Returns a string created by using a given format string as a template into
// which the remaining argument values are substituted.
//
// format: A format string. See [Formatting String Objects] for examples of how to use
// this method, and [String Format Specifiers] for a list of format
// specifiers. This value must not be `nil`.
//
// # Return Value
//
// A string created by using `format` as a template into which the remaining
// argument values are substituted without any localization.
//
// # Discussion
//
// Pass a comma-separated list of trailing variadic arguments to substitute
// into `format`.
//
// This method invokes [InitWithFormatLocaleArguments] without applying any
// localization. This is useful, for example, when working with fixed-format
// representations of information that is written out and read back in at a
// later time.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/stringWithFormat:
//
// [Formatting String Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/FormatStrings.html#//apple_ref/doc/uid/20000943
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
func (_NSStringClass NSStringClass) StringWithFormat(format string) NSString {
	rv := objc.Send[objc.ID](objc.ID(_NSStringClass.class), objc.Sel("stringWithFormat:"), objc.String(format))
	return NSStringFromID(rv)
}

// Returns a string created by copying the characters from another given
// string.
//
// string: The string from which to copy characters. This value must not be `nil`.
//
// # Return Value
//
// A string created by copying the characters from `aString`.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/stringWithString:
func (_NSStringClass NSStringClass) StringWithString(string_ string) NSString {
	rv := objc.Send[objc.ID](objc.ID(_NSStringClass.class), objc.Sel("stringWithString:"), objc.String(string_))
	return NSStringFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSString/stringWithUTF8String:
func (_NSStringClass NSStringClass) StringWithUTF8String(nullTerminatedCString string) NSString {
	rv := objc.Send[objc.ID](objc.ID(_NSStringClass.class), objc.Sel("stringWithUTF8String:"), unsafe.Pointer(unsafe.StringData(nullTerminatedCString+"\x00")))
	return NSStringFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSString/stringWithValidatedFormat:validFormatSpecifiers:error:
func (_NSStringClass NSStringClass) StringWithValidatedFormatValidFormatSpecifiersError(format string, validFormatSpecifiers string) (NSString, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSStringClass.class), objc.Sel("stringWithValidatedFormat:validFormatSpecifiers:error:"), objc.String(format), objc.String(validFormatSpecifiers), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSString{}, NSErrorFrom(errorPtr)
	}
	return NSStringFromID(rv), nil

}

// The number of UTF-16 code units in the receiver.
//
// # Discussion
//
// This number includes the individual characters of composed character
// sequences, so you cannot use this property to determine if a string will be
// visible when printed or how long it will appear.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/length
func (s NSString) Length() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("length"))
	return rv
}

// A null-terminated UTF8 representation of the string.
//
// # Discussion
//
// This C string is a pointer to a structure inside the string object, which
// may have a lifetime shorter than the string object and will certainly not
// have a longer lifetime. Therefore, you should copy the C string if it needs
// to be stored outside of the memory context in which you use this property.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/utf8String
func (s NSString) UTF8String() string {
	rv := objc.Send[*byte](s.ID, objc.Sel("UTF8String"))
	return objc.GoString(rv)
}

// An unsigned integer that can be used as a hash table address.
//
// # Discussion
//
// If two string objects are equal (as determined by the [IsEqualToString]
// method), they must have the same hash value. This property fulfills this
// requirement.
//
// You should not rely on this property having the same hash value across
// releases of macOS.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/hash
func (s NSString) Hash() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("hash"))
	return rv
}

// A lowercase representation of the string.
//
// # Discussion
//
// This property performs the canonical (non-localized) mapping. It is
// suitable for programming operations that require stable results not
// depending on the current locale.
//
// Case transformations aren’t guaranteed to be symmetrical or to produce
// strings of the same lengths as the originals. That is, the result of this
// statement:
//
// …might not be equal to this statement:
//
// For example, the uppercase form of “ß” in German is “SS”, so
// converting “Straße” to uppercase, then lowercase, produces this
// sequence of strings:
//
// - “Straße”
// - “STRASSE”
// - “strasse”
//
// See: https://developer.apple.com/documentation/Foundation/NSString/lowercased
func (s NSString) LowercaseString() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("lowercaseString"))
	return NSStringFromID(rv).String()
}

// Returns a version of the string with all letters converted to lowercase,
// taking into account the current locale.
//
// # Discussion
//
// Case transformations aren’t guaranteed to be symmetrical or to produce
// strings of the same lengths as the originals. See [LowercaseString] for an
// example.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/localizedLowercase
func (s NSString) LocalizedLowercaseString() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("localizedLowercaseString"))
	return NSStringFromID(rv).String()
}

// An uppercase representation of the string.
//
// # Discussion
//
// This property performs the canonical (non-localized) mapping. It is
// suitable for programming operations that require stable results not
// depending on the current locale.
//
// Case transformations aren’t guaranteed to be symmetrical or to produce
// strings of the same lengths as the originals. See [LowercaseString] for an
// example.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/uppercased
func (s NSString) UppercaseString() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("uppercaseString"))
	return NSStringFromID(rv).String()
}

// Returns a version of the string with all letters converted to uppercase,
// taking into account the current locale.
//
// # Discussion
//
// Case transformations aren’t guaranteed to be symmetrical or to produce
// strings of the same lengths as the originals. See [LowercaseString] for an
// example.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/localizedUppercase
func (s NSString) LocalizedUppercaseString() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("localizedUppercaseString"))
	return NSStringFromID(rv).String()
}

// A capitalized representation of the string.
//
// # Discussion
//
// A capitalized string is a string with the first character in each word
// changed to its corresponding uppercase value, and all remaining characters
// set to their corresponding lowercase values. A word is any sequence of
// characters delimited by spaces, tabs, or line terminators (listed under
// [GetLineStartEndContentsEndForRange]). Some common word delimiting
// punctuation isn’t considered, so this property may not generally produce
// the desired results for multiword strings.
//
// Case transformations aren’t guaranteed to be symmetrical or to produce
// strings of the same lengths as the originals. See [LowercaseString] for an
// example.
//
// This property performs the canonical (non-localized) mapping. It is
// suitable for programming operations that require stable results not
// depending on the current locale.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/capitalized
func (s NSString) CapitalizedString() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("capitalizedString"))
	return NSStringFromID(rv).String()
}

// Returns a capitalized representation of the receiver using the current
// locale.
//
// # Discussion
//
// A capitalized string is a string with the first character in each word
// changed to its corresponding uppercase value, and all remaining characters
// set to their corresponding lowercase values. A “word” is any sequence
// of characters delimited by spaces, tabs, or line terminators (listed under
// [GetLineStartEndContentsEndForRange]). Some common word delimiting
// punctuation isn’t considered, so this property may not generally produce
// the desired results for multiword strings.
//
// Case transformations aren’t guaranteed to be symmetrical or to produce
// strings of the same lengths as the originals. See [LowercaseString] for an
// example.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/localizedCapitalized
func (s NSString) LocalizedCapitalizedString() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("localizedCapitalizedString"))
	return NSStringFromID(rv).String()
}

// A string made by normalizing the string’s contents using the Unicode
// Normalization Form D.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/decomposedStringWithCanonicalMapping
func (s NSString) DecomposedStringWithCanonicalMapping() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("decomposedStringWithCanonicalMapping"))
	return NSStringFromID(rv).String()
}

// A string made by normalizing the receiver’s contents using the Unicode
// Normalization Form KD.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/decomposedStringWithCompatibilityMapping
func (s NSString) DecomposedStringWithCompatibilityMapping() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("decomposedStringWithCompatibilityMapping"))
	return NSStringFromID(rv).String()
}

// A string made by normalizing the string’s contents using the Unicode
// Normalization Form C.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/precomposedStringWithCanonicalMapping
func (s NSString) PrecomposedStringWithCanonicalMapping() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("precomposedStringWithCanonicalMapping"))
	return NSStringFromID(rv).String()
}

// A string made by normalizing the receiver’s contents using the Unicode
// Normalization Form KC.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/precomposedStringWithCompatibilityMapping
func (s NSString) PrecomposedStringWithCompatibilityMapping() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("precomposedStringWithCompatibilityMapping"))
	return NSStringFromID(rv).String()
}

// The floating-point value of the string as a `double`.
//
// # Discussion
//
// This property doesn’t include any whitespace at the beginning of the
// string. This property is `HUGE_VAL` or `–HUGE_VAL` on overflow, `0.0` on
// underflow. This property is `0.0` if the string doesn’t begin with a
// valid text representation of a floating-point number.
//
// This property uses formatting information stored in the non-localized
// value; use an [NSScanner] object for localized scanning of numeric values
// from a string.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/doubleValue
func (s NSString) DoubleValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("doubleValue"))
	return rv
}

// The floating-point value of the string as a `float`.
//
// # Discussion
//
// This property doesn’t include whitespace at the beginning of the string.
// This property is `HUGE_VAL` or `–HUGE_VAL` on overflow, `0.0` on
// underflow. This property is `0.0` if the string doesn’t begin with a
// valid text representation of a floating-point number.
//
// This method uses formatting information stored in the non-localized value;
// use an [NSScanner] object for localized scanning of numeric values from a
// string.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/floatValue
func (s NSString) FloatValue() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("floatValue"))
	return rv
}

// The integer value of the string.
//
// # Discussion
//
// The integer value of the string, assuming a decimal representation and
// skipping whitespace at the beginning of the string. This property is
// `INT_MAX` or `INT_MIN` on overflow. This property is `0` if the string
// doesn’t begin with a valid decimal text representation of a number.
//
// This property uses formatting information stored in the non-localized
// value; use an [NSScanner] object for localized scanning of numeric values
// from a string.
//
// # Special Considerations
//
// In macOS 10.5 and later, use [IntegerValue] instead.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/intValue
func (s NSString) IntValue() int {
	rv := objc.Send[int](s.ID, objc.Sel("intValue"))
	return rv
}

// The [NSInteger] value of the string.
//
// # Discussion
//
// The [NSInteger] value of the string, assuming a decimal representation and
// skipping whitespace at the beginning of the string. This property is `0` if
// the string doesn’t begin with a valid decimal text representation of a
// number.
//
// This property uses formatting information stored in the non-localized
// value; use an [NSScanner] object for localized scanning of numeric values
// from a string.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/integerValue
func (s NSString) IntegerValue() int {
	rv := objc.Send[int](s.ID, objc.Sel("integerValue"))
	return rv
}

// The `long long` value of the string.
//
// # Discussion
//
// The `long long` value of the string, assuming a decimal representation and
// skipping whitespace at the beginning of the string. This property is
// `LLONG_MAX` or `LLONG_MIN` on overflow. This property is `0` if the
// receiver doesn’t begin with a valid decimal text representation of a
// number.
//
// This property uses formatting information stored in the non-localized
// value; use an [NSScanner] object for localized scanning of numeric values
// from a string.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/longLongValue
func (s NSString) LongLongValue() int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("longLongValue"))
	return rv
}

// The Boolean value of the string.
//
// # Discussion
//
// This property is true on encountering one of “Y”, “y”, “T”,
// “t”, or a digit 1-9—the method ignores any trailing characters. This
// property is false if the receiver doesn’t begin with a valid decimal text
// representation of a number.
//
// The property assumes a decimal representation and skips whitespace at the
// beginning of the string. It also skips initial whitespace characters, or
// optional -/+ sign followed by zeroes.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/boolValue
func (s NSString) BoolValue() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("boolValue"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSString/description
func (s NSString) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return NSStringFromID(rv).String()
}

// The fastest encoding to which the receiver may be converted without loss of
// information.
//
// # Discussion
//
// “Fastest” applies to retrieval of characters from the string. This
// encoding may not be space efficient.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/fastestEncoding
func (s NSString) FastestEncoding() NSStringEncoding {
	rv := objc.Send[NSStringEncoding](s.ID, objc.Sel("fastestEncoding"))
	return NSStringEncoding(rv)
}

// The smallest encoding to which the receiver can be converted without loss
// of information.
//
// # Discussion
//
// This encoding may not be the fastest for accessing characters, but is
// space-efficient. This property may take some time to access.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/smallestEncoding
func (s NSString) SmallestEncoding() NSStringEncoding {
	rv := objc.Send[NSStringEncoding](s.ID, objc.Sel("smallestEncoding"))
	return NSStringEncoding(rv)
}

// The file-system path components of the receiver.
//
// # Discussion
//
// The strings in the array appear in the order they did in the receiver. If
// the string begins or ends with the path separator, then the first or last
// component, respectively, will contain the separator. Empty components
// (caused by consecutive path separators) are deleted. For example, this code
// excerpt:
//
// produces an array with these contents:
//
// [Table data omitted]
//
// If the receiver begins with a slash—for example,
// “`/tmp/scratch`”—the array has these contents:
//
// [Table data omitted]
//
// If the receiver has no separators—for example, “`scratch`”—the
// array contains the string itself, in this case “`scratch`”.
//
// Note that this method only works with file paths—not, for example, string
// representations of URLs.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/pathComponents
func (s NSString) PathComponents() []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("pathComponents"))
	return objc.ConvertSliceToStrings(rv)
}

// A file system-specific representation of the receiver.
//
// # Discussion
//
// The returned C string will be automatically freed just as a returned object
// would be released; your code should copy the representation or use
// [GetFileSystemRepresentationMaxLength] if it needs to store the
// representation outside of the memory context in which the representation
// was created.
//
// Raises an [characterConversionException] if the receiver can’t be
// represented in the file system’s encoding. It also raises an exception if
// the receiver contains no characters.
//
// Note that this method only works with file paths (not, for example, string
// representations of URLs).
//
// To convert a `char *` path (such as you might get from a C library routine)
// to an [NSString] object, use the [StringWithFileSystemRepresentationLength]
// method on [NSFileManager].
//
// See: https://developer.apple.com/documentation/Foundation/NSString/fileSystemRepresentation
//
// [characterConversionException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/characterConversionException
func (s NSString) FileSystemRepresentation() string {
	rv := objc.Send[*byte](s.ID, objc.Sel("fileSystemRepresentation"))
	return objc.GoString(rv)
}

// A Boolean value that indicates whether the receiver represents an absolute
// path.
//
// # Discussion
//
// true if the receiver (if interpreted as a path) represents an absolute
// path, otherwise false.
//
// See [String Programming Guide] for more information on paths.
//
// Note that this method only works with file paths (not, for example, string
// representations of URLs). The method does not check the filesystem for the
// existence of the path (use [FileExistsAtPath] or similar methods in
// [NSFileManager] for that task).
//
// See: https://developer.apple.com/documentation/Foundation/NSString/isAbsolutePath
//
// [String Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/introStrings.html#//apple_ref/doc/uid/10000035i
func (s NSString) AbsolutePath() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isAbsolutePath"))
	return rv
}

// The last path component of the receiver.
//
// # Discussion
//
// Path components are alphanumeric strings delineated by the path separator
// (slash “/”) or the beginning or end of the path string. Multiple path
// separators at the end of the string are stripped.
//
// The following table illustrates the effect of [LastPathComponent] on a
// variety of different paths:
//
// [Table data omitted]
//
// Note that this method only works with file paths (not, for example, string
// representations of URLs).
//
// See: https://developer.apple.com/documentation/Foundation/NSString/lastPathComponent
func (s NSString) LastPathComponent() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("lastPathComponent"))
	return NSStringFromID(rv).String()
}

// The path extension, if any, of the string as interpreted as a path.
//
// # Discussion
//
// The path extension is the portion of the last path component which follows
// the final period, if there is one. The extension divider is not included.
// The following table illustrates the effect of [PathExtension] on a variety
// of different paths:
//
// [Table data omitted]
//
// Note that this method only works with file paths (not, for example, string
// representations of URLs).
//
// See: https://developer.apple.com/documentation/Foundation/NSString/pathExtension
func (s NSString) PathExtension() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("pathExtension"))
	return NSStringFromID(rv).String()
}

// A new string that replaces the current home directory portion of the
// current path with a tilde (`~`) character.
//
// # Discussion
//
// A new string based on the current string object. If the new string
// specifies a file in the current home directory, the home directory portion
// of the path is replaced with a tilde (`~`) character. If the string does
// not specify a file in the current home directory, this method returns a new
// string object whose path is unchanged from the path in the current string.
//
// Note that this method only works with file paths. It does not work for
// string representations of URLs.
//
// For sandboxed apps in macOS, the current home directory is not the same as
// the user’s home directory. For a sandboxed app, the home directory is the
// app’s home directory. So if you specified a path of
// `/Users/“/file.Txt()` for a sandboxed app, the returned path would be
// unchanged from the original. However, if you specified the same path for an
// app not in a sandbox, this method would replace the `/Users/` portion of
// the path with a tilde.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/abbreviatingWithTildeInPath
func (s NSString) StringByAbbreviatingWithTildeInPath() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByAbbreviatingWithTildeInPath"))
	return NSStringFromID(rv).String()
}

// A new string made by deleting the last path component from the receiver,
// along with any final path separator.
//
// # Discussion
//
// A new string made by deleting the last path component from the receiver,
// along with any final path separator. If the receiver represents the root
// path it is returned unaltered.
//
// The following table illustrates the effect of this method on a variety of
// different paths:
//
// [Table data omitted]
//
// Note that this method only works with file paths (not, for example, string
// representations of URLs).
//
// See: https://developer.apple.com/documentation/Foundation/NSString/deletingLastPathComponent
func (s NSString) StringByDeletingLastPathComponent() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByDeletingLastPathComponent"))
	return NSStringFromID(rv).String()
}

// A new string made by deleting the extension (if any, and only the last)
// from the receiver.
//
// # Discussion
//
// A new string made by deleting the extension (if any, and only the last)
// from the receiver. Strips any trailing path separator before checking for
// an extension. If the receiver represents the root path, it is returned
// unaltered.
//
// The following table illustrates the effect of this method on a variety of
// different paths:
//
// [Table data omitted]
//
// Note that attempting to delete an extension from `@"XCUIElementTypeTiff"`
// causes the result to be `@"XCUIElementTypeTiff"` instead of an empty
// string. This difference is because a file named `@"XCUIElementTypeTiff"` is
// not considered to have an extension, so nothing is deleted. Note also that
// this method only works with file paths (not, for example, string
// representations of URLs).
//
// See: https://developer.apple.com/documentation/Foundation/NSString/deletingPathExtension
func (s NSString) StringByDeletingPathExtension() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByDeletingPathExtension"))
	return NSStringFromID(rv).String()
}

// A new string made by expanding the initial component of the receiver to its
// full path value.
//
// # Discussion
//
// A new string made by expanding the initial component of the receiver, if it
// begins with “`~`” or “`~user`”, to its full path value. Returns a
// new string matching the receiver if the receiver’s initial component
// can’t be expanded.
//
// Note that this method only works with file paths (not, for example, string
// representations of URLs).
//
// See: https://developer.apple.com/documentation/Foundation/NSString/expandingTildeInPath
func (s NSString) StringByExpandingTildeInPath() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByExpandingTildeInPath"))
	return NSStringFromID(rv).String()
}

// A new string made from the receiver by resolving all symbolic links and
// standardizing path.
//
// # Discussion
//
// A new string made by resolving all symbolic links, then removing extraneous
// path components. For absolute paths, all symbolic links are guaranteed to
// be removed. For relative paths, symbolic links that can’t be resolved are
// left unresolved in the returned string.
//
// Returns `self` if an error occurs.
//
// Note that this method only works with file paths (not, for example, string
// representations of URLs).
//
// See: https://developer.apple.com/documentation/Foundation/NSString/resolvingSymlinksInPath
func (s NSString) StringByResolvingSymlinksInPath() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByResolvingSymlinksInPath"))
	return NSStringFromID(rv).String()
}

// A new string made by removing extraneous path components from the receiver.
//
// # Discussion
//
// A new string made by performing the following operations:
//
// - Expanding an initial tilde expression using
// [StringByExpandingTildeInPath]. - Removing an initial component of
// “`/private/var/automount`”, “`/var/automount`”, or “`/private`”
// from the path, if the result still indicates an existing file or directory
// (checked by consulting the file system). - Reducing empty components and
// references to the current directory (that is, the sequences “//” and
// “/./”) to single path separators. - Removing a trailing slash from the
// last component. - For absolute paths only, resolving references to the
// parent directory (that is, the component “..”) to the real parent
// directory if possible using [StringByResolvingSymlinksInPath]. For relative
// paths, references to the parent directory are left in place.
//
// Returns `self` if an error occurs.
//
// Note that the path returned by this method may still have symbolic link
// components in it. Note also that this method only works with file paths
// (not, for example, string representations of URLs).
//
// See: https://developer.apple.com/documentation/Foundation/NSString/standardizingPath
func (s NSString) StringByStandardizingPath() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByStandardizingPath"))
	return NSStringFromID(rv).String()
}

// Returns a new string made from the receiver by replacing all percent
// encoded sequences with the matching UTF-8 characters.
//
// # Return Value
//
// A new string with the percent-encoded sequences removed, or `nil` if the
// receiver contains an invalid percent-encoding sequence.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/NSString/removingPercentEncoding
func (s NSString) StringByRemovingPercentEncoding() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stringByRemovingPercentEncoding"))
	return NSStringFromID(rv).String()
}

// A custom playground Quick Look for this instance.
//
// See: https://developer.apple.com/documentation/foundation/nsstring/customplaygroundquicklook
func (s NSString) CustomPlaygroundQuickLook() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("customPlaygroundQuickLook"))
	return objectivec.Object{ID: rv}
}
func (s NSString) SetCustomPlaygroundQuickLook(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setCustomPlaygroundQuickLook:"), value)
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
func (s NSString) WritableTypeIdentifiersForItemProvider() []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("writableTypeIdentifiersForItemProvider"))
	return objc.ConvertSliceToStrings(rv)
}

// Returns a zero-terminated list of the encodings string objects support in
// the application’s environment.
//
// # Return Value
//
// A zero-terminated list of the encodings string objects support in the
// application’s environment.
//
// # Discussion
//
// Among the more commonly used encodings are:
//
// - [NSASCIIStringEncoding] - [NSUnicodeStringEncoding] -
// [NSISOLatin1StringEncoding] - [NSISOLatin2StringEncoding] -
// [NSSymbolStringEncoding]
//
// See the [NSStringEncoding] type for a larger list and descriptions of many
// supported encodings. In addition to those encodings listed here, you can
// also use the encodings defined for CFString in Core Foundation; you just
// need to call the [CFStringConvertEncodingToNSStringEncoding(_:)] function
// to convert them to a usable format.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/availableStringEncodings
//
// [CFStringConvertEncodingToNSStringEncoding(_:)]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertEncodingToNSStringEncoding(_:)
func (_NSStringClass NSStringClass) AvailableStringEncodings() NSStringEncoding {
	rv := objc.Send[NSStringEncoding](objc.ID(_NSStringClass.class), objc.Sel("availableStringEncodings"))
	return NSStringEncoding(rv)
}

// Returns the C-string encoding assumed for any method accepting a C string
// as an argument.
//
// # Return Value
//
// The C-string encoding assumed for any method accepting a C string as an
// argument.
//
// # Discussion
//
// This method returns a user-dependent encoding who value is derived from
// user’s default language and potentially other factors. You might
// sometimes need to use this encoding when interpreting user documents with
// unknown encodings, in the absence of other hints, but in general this
// encoding should be used rarely, if at all. Note that some potential values
// might result in unexpected encoding conversions of even fairly
// straightforward [NSString] content—for example, punctuation characters
// with a bidirectional encoding.
//
// Methods that accept a C string as an argument use `...CString...` in the
// keywords for such arguments: for example, [StringWithCString]—note,
// though, that these are deprecated. The default C-string encoding is
// determined from system information and can’t be changed programmatically
// for an individual process. See [NSStringEncoding] for a full list of
// supported encodings.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/defaultCStringEncoding
func (_NSStringClass NSStringClass) DefaultCStringEncoding() NSStringEncoding {
	rv := objc.Send[NSStringEncoding](objc.ID(_NSStringClass.class), objc.Sel("defaultCStringEncoding"))
	return NSStringEncoding(rv)
}

// Protocol methods for NSCopying

// Protocol methods for NSItemProviderReading

// Protocol methods for NSItemProviderWriting

// Protocol methods for NSMutableCopying

// Protocol methods for NSSecureCoding

// LoadDataWithTypeIdentifierForItemProvider is a synchronous wrapper around [NSString.LoadDataWithTypeIdentifierForItemProviderCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s NSString) LoadDataWithTypeIdentifierForItemProvider(ctx context.Context, typeIdentifier string) (*NSData, error) {
	type result struct {
		val *NSData
		err error
	}
	done := make(chan result, 1)
	s.LoadDataWithTypeIdentifierForItemProviderCompletionHandler(typeIdentifier, func(val *NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
