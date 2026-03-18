// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMutableString] class.
var (
	_NSMutableStringClass     NSMutableStringClass
	_NSMutableStringClassOnce sync.Once
)

func getNSMutableStringClass() NSMutableStringClass {
	_NSMutableStringClassOnce.Do(func() {
		_NSMutableStringClass = NSMutableStringClass{class: objc.GetClass("NSMutableString")}
	})
	return _NSMutableStringClass
}

// GetNSMutableStringClass returns the class object for NSMutableString.
func GetNSMutableStringClass() NSMutableStringClass {
	return getNSMutableStringClass()
}

type NSMutableStringClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMutableStringClass) Alloc() NSMutableString {
	rv := objc.Send[NSMutableString](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A dynamic plain-text Unicode string object.
//
// # Overview
// 
// In Swift, you can use this type instead of a [String] in cases that require
// reference semantics.
// 
// The [NSMutableString] class declares the programmatic interface to an
// object that manages a mutable string—that is, a string whose contents can
// be edited—that conceptually represents an array of Unicode characters. To
// construct and manage an immutable string—or a string that cannot be
// changed after it has been created—use an object of the [NSString] class.
// 
// The [NSMutableString] class adds one primitive
// method—[NSMutableString.ReplaceCharactersInRangeWithString]—to the basic
// string-handling behavior inherited from [NSString]. All other methods that
// modify a string work through this method. For example,
// [NSMutableString.InsertStringAtIndex] simply replaces the characters in a range of `0`
// length, while [NSMutableString.DeleteCharactersInRange] replaces the characters in a given
// range with no characters.
// 
// NSMutableString is “toll-free bridged” with its Core Foundation
// counterpart, [CFMutableString]. See [Toll-Free Bridging] for more
// information.
//
// [CFMutableString]: https://developer.apple.com/documentation/CoreFoundation/CFMutableString
// [String]: https://developer.apple.com/documentation/Swift/String
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
//
// # Creating and Initializing a Mutable String
//
//   - [NSMutableString.InitWithCapacity]: Returns an [NSMutableString] object initialized with initial storage for a given number of characters,
//
// # Modifying a String
//
//   - [NSMutableString.AppendString]: Adds to the end of the receiver the characters of a given string.
//   - [NSMutableString.ApplyTransformReverseRangeUpdatedRange]: Transliterates the receiver by applying a specified ICU string transform.
//   - [NSMutableString.DeleteCharactersInRange]: Removes from the receiver the characters in a given range.
//   - [NSMutableString.InsertStringAtIndex]: Inserts into the receiver the characters of a given string at a given location.
//   - [NSMutableString.ReplaceCharactersInRangeWithString]: Replaces the characters from `range` with those in `aString`.
//   - [NSMutableString.ReplaceOccurrencesOfStringWithStringOptionsRange]: Replaces all occurrences of a given string in a given range with another given string, returning the number of replacements.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableString
type NSMutableString struct {
	NSString
}

// NSMutableStringFromID constructs a [NSMutableString] from an objc.ID.
//
// A dynamic plain-text Unicode string object.
func NSMutableStringFromID(id objc.ID) NSMutableString {
	return NSMutableString{NSString: NSStringFromID(id)}
}
// NOTE: NSMutableString adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMutableString] class.
//
// # Creating and Initializing a Mutable String
//
//   - [INSMutableString.InitWithCapacity]: Returns an [NSMutableString] object initialized with initial storage for a given number of characters,
//
// # Modifying a String
//
//   - [INSMutableString.AppendString]: Adds to the end of the receiver the characters of a given string.
//   - [INSMutableString.ApplyTransformReverseRangeUpdatedRange]: Transliterates the receiver by applying a specified ICU string transform.
//   - [INSMutableString.DeleteCharactersInRange]: Removes from the receiver the characters in a given range.
//   - [INSMutableString.InsertStringAtIndex]: Inserts into the receiver the characters of a given string at a given location.
//   - [INSMutableString.ReplaceCharactersInRangeWithString]: Replaces the characters from `range` with those in `aString`.
//   - [INSMutableString.ReplaceOccurrencesOfStringWithStringOptionsRange]: Replaces all occurrences of a given string in a given range with another given string, returning the number of replacements.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableString
type INSMutableString interface {
	INSString

	// Topic: Creating and Initializing a Mutable String

	// Returns an [NSMutableString] object initialized with initial storage for a given number of characters,
	InitWithCapacity(capacity uint) NSMutableString

	// Topic: Modifying a String

	// Adds to the end of the receiver the characters of a given string.
	AppendString(aString string)
	// Transliterates the receiver by applying a specified ICU string transform.
	ApplyTransformReverseRangeUpdatedRange(transform NSStringTransform, reverse bool, range_ NSRange, resultingRange NSRangePointer) bool
	// Removes from the receiver the characters in a given range.
	DeleteCharactersInRange(range_ NSRange)
	// Inserts into the receiver the characters of a given string at a given location.
	InsertStringAtIndex(aString string, loc uint)
	// Replaces the characters from `range` with those in `aString`.
	ReplaceCharactersInRangeWithString(range_ NSRange, aString string)
	// Replaces all occurrences of a given string in a given range with another given string, returning the number of replacements.
	ReplaceOccurrencesOfStringWithStringOptionsRange(target string, replacement string, options NSStringCompareOptions, searchRange NSRange) uint

	// Adds a constructed string to the receiver.
	AppendFormat(format string)
}

// Init initializes the instance.
func (m NSMutableString) Init() NSMutableString {
	rv := objc.Send[NSMutableString](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMutableString) Autorelease() NSMutableString {
	rv := objc.Send[NSMutableString](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMutableString creates a new NSMutableString instance.
func NewNSMutableString() NSMutableString {
	class := getNSMutableStringClass()
	rv := objc.Send[NSMutableString](objc.ID(class.class), objc.Sel("new"))
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
func NewMutableStringWithBytesLengthEncoding(bytes []byte, encoding uint) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytes:length:encoding:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)), encoding)
	return NSMutableStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(bytesNoCopy:length:encoding:deallocator:)
func NewMutableStringWithBytesNoCopyLengthEncodingDeallocator(bytes unsafe.Pointer, len_ uint, encoding uint, deallocator unsafe.Pointer) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytesNoCopy:length:encoding:deallocator:"), bytes, len_, encoding, deallocator)
	return NSMutableStringFromID(rv)
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
// freeBuffer: If [true], the receiver releases the memory with `free()` when it no longer
// needs the data; if [false] it won’t.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// freed even if `flag` is [true]. In this case, the caller is responsible for
// freeing the buffer. This allows the caller to continue trying to create a
// string with the buffer, without having the buffer deallocated.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(bytesNoCopy:length:encoding:freeWhenDone:)
func NewMutableStringWithBytesNoCopyLengthEncodingFreeWhenDone(bytes unsafe.Pointer, len_ uint, encoding uint, freeBuffer bool) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytesNoCopy:length:encoding:freeWhenDone:"), bytes, len_, encoding, freeBuffer)
	return NSMutableStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(cString:)
func NewMutableStringWithCString(bytes []byte) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCString:"), unsafe.Pointer(unsafe.SliceData(bytes)))
	return NSMutableStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(cString:encoding:)
func NewMutableStringWithCStringEncoding(nullTerminatedCString []byte, encoding uint) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCString:encoding:"), unsafe.Pointer(unsafe.SliceData(nullTerminatedCString)), encoding)
	return NSMutableStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(cString:length:)
func NewMutableStringWithCStringLength(bytes []byte, length uint) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCString:length:"), unsafe.Pointer(unsafe.SliceData(bytes)), length)
	return NSMutableStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(cStringNoCopy:length:freeWhenDone:)
func NewMutableStringWithCStringNoCopyLengthFreeWhenDone(bytes []byte, length uint, freeBuffer bool) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCStringNoCopy:length:freeWhenDone:"), unsafe.Pointer(unsafe.SliceData(bytes)), length, freeBuffer)
	return NSMutableStringFromID(rv)
}

// Returns an [NSMutableString] object initialized with initial storage for a
// given number of characters,
//
// capacity: The number of characters the string is expected to initially contain.
//
// # Return Value
// 
// An initialized [NSMutableString] object with initial storage for `capacity`
// characters. The returned object might be different than the original
// receiver.
//
// # Discussion
// 
// The number of characters indicated by `capacity` is simply a hint to
// increase the efficiency of data storage. The value does limit the length of
// the string.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableString/init(capacity:)
func NewMutableStringWithCapacity(capacity uint) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCapacity:"), capacity)
	return NSMutableStringFromID(rv)
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
func NewMutableStringWithCharactersLength(characters unsafe.Pointer, length uint) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCharacters:length:"), characters, length)
	return NSMutableStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(charactersNoCopy:length:deallocator:)
func NewMutableStringWithCharactersNoCopyLengthDeallocator(chars unsafe.Pointer, len_ uint, deallocator unsafe.Pointer) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCharactersNoCopy:length:deallocator:"), chars, len_, deallocator)
	return NSMutableStringFromID(rv)
}

// Returns an initialized [NSString] object that contains a given number of
// characters from a given C array of UTF-16 code units.
//
// characters: A C array of UTF-16 code units.
//
// length: The number of characters to use from `characters`.
//
// freeBuffer: If [true], the receiver releases the memory with `free()` when it no longer
// needs the data; if [false] it won’t.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// freed even if `flag` is [true]. In this case, the caller is responsible for
// freeing the buffer. This allows the caller to continue trying to create a
// string with the buffer, without having the buffer deallocated.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(charactersNoCopy:length:freeWhenDone:)
func NewMutableStringWithCharactersNoCopyLengthFreeWhenDone(characters unsafe.Pointer, length uint, freeBuffer bool) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCharactersNoCopy:length:freeWhenDone:"), characters, length, freeBuffer)
	return NSMutableStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(coder:)
func NewMutableStringWithCoder(coder INSCoder) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSMutableStringFromID(rv)
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
func NewMutableStringWithContentsOfFile(path string) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:"), objc.String(path))
	return NSMutableStringFromID(rv)
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
func NewMutableStringWithContentsOfFileEncodingError(path string, enc uint) (NSMutableString, error) {
	var errorPtr objc.ID
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:encoding:error:"), objc.String(path), enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSMutableStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSMutableStringFromID(rv), nil
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
func NewMutableStringWithContentsOfFileUsedEncodingError(path string, enc unsafe.Pointer) (NSMutableString, error) {
	var errorPtr objc.ID
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:usedEncoding:error:"), objc.String(path), enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSMutableStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSMutableStringFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOf:)
func NewMutableStringWithContentsOfURL(url INSURL) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	return NSMutableStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOf:encoding:)
func NewMutableStringWithContentsOfURLEncodingError(url INSURL, enc uint) (NSMutableString, error) {
	var errorPtr objc.ID
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:encoding:error:"), url, enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSMutableStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSMutableStringFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOf:usedEncoding:)
func NewMutableStringWithContentsOfURLUsedEncodingError(url INSURL, enc unsafe.Pointer) (NSMutableString, error) {
	var errorPtr objc.ID
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:usedEncoding:error:"), url, enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSMutableStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSMutableStringFromID(rv), nil
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
func NewMutableStringWithDataEncoding(data INSData, encoding uint) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:encoding:"), data, encoding)
	return NSMutableStringFromID(rv)
}

// Returns an [NSString] object initialized by using a given format string as
// a template into which the remaining argument values are substituted.
//
// format: A format string. See [Formatting String Objects] for examples of how to use
// this method, and [String Format Specifiers] for a list of format
// specifiers. This value must not be `nil`.
// //
// [Formatting String Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/FormatStrings.html#//apple_ref/doc/uid/20000943
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
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
func NewMutableStringWithFormat(format string) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:"), objc.String(format))
	return NSMutableStringFromID(rv)
}

// Returns an [NSString] object initialized by using a given format string as
// a template into which the remaining argument values are substituted without
// any localization.
//
// format: A format string. See [Formatting String Objects] for examples of how to use
// this method, and [String Format Specifiers] for a list of format
// specifiers. This value must not be `nil`.
// //
// [Formatting String Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/FormatStrings.html#//apple_ref/doc/uid/20000943
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
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
func NewMutableStringWithFormatArguments(format string, argList unsafe.Pointer) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:arguments:"), objc.String(format), argList)
	return NSMutableStringFromID(rv)
}

// Returns an [NSString] object initialized by using a given format string as
// a template into which the remaining argument values are substituted
// according to given locale.
//
// format: A format string. See [Formatting String Objects] for examples of how to use
// this method, and [String Format Specifiers] for a list of format
// specifiers. This value must not be `nil`.
// //
// [Formatting String Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/FormatStrings.html#//apple_ref/doc/uid/20000943
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
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
func NewMutableStringWithFormatLocale(format string, locale objectivec.IObject) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:locale:"), objc.String(format), locale)
	return NSMutableStringFromID(rv)
}

// Returns an [NSString] object initialized by using a given format string as
// a template into which the remaining argument values are substituted
// according to given locale information. This method is meant to be called
// from within a variadic function, where the argument list will be available.
//
// format: A format string. See [Formatting String Objects] for examples of how to use
// this method, and [String Format Specifiers] for a list of format
// specifiers. This value must not be `nil`.
// //
// [Formatting String Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/FormatStrings.html#//apple_ref/doc/uid/20000943
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
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
// [String Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/introStrings.html#//apple_ref/doc/uid/10000035i
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(format:locale:arguments:)
func NewMutableStringWithFormatLocaleArguments(format string, locale objectivec.IObject, argList unsafe.Pointer) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:locale:arguments:"), objc.String(format), locale, argList)
	return NSMutableStringFromID(rv)
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
func NewMutableStringWithString(aString string) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:"), objc.String(aString))
	return NSMutableStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(utf8String:)
func NewMutableStringWithUTF8String(nullTerminatedCString []byte) NSMutableString {
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUTF8String:"), unsafe.Pointer(unsafe.SliceData(nullTerminatedCString)))
	return NSMutableStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:arguments:error:
func NewMutableStringWithValidatedFormatValidFormatSpecifiersArgumentsError(format string, validFormatSpecifiers string, argList unsafe.Pointer) (NSMutableString, error) {
	var errorPtr objc.ID
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:arguments:error:"), objc.String(format), objc.String(validFormatSpecifiers), argList, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSMutableStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSMutableStringFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:error:
func NewMutableStringWithValidatedFormatValidFormatSpecifiersError(format string, validFormatSpecifiers string) (NSMutableString, error) {
	var errorPtr objc.ID
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:error:"), objc.String(format), objc.String(validFormatSpecifiers), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSMutableStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSMutableStringFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:locale:arguments:error:
func NewMutableStringWithValidatedFormatValidFormatSpecifiersLocaleArgumentsError(format string, validFormatSpecifiers string, locale objectivec.IObject, argList unsafe.Pointer) (NSMutableString, error) {
	var errorPtr objc.ID
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:locale:arguments:error:"), objc.String(format), objc.String(validFormatSpecifiers), locale, argList, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSMutableStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSMutableStringFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:locale:error:
func NewMutableStringWithValidatedFormatValidFormatSpecifiersLocaleError(format string, validFormatSpecifiers string, locale objectivec.IObject) (NSMutableString, error) {
	var errorPtr objc.ID
	instance := getNSMutableStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:locale:error:"), objc.String(format), objc.String(validFormatSpecifiers), locale, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSMutableStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSMutableStringFromID(rv), nil
}

// Returns an [NSMutableString] object initialized with initial storage for a
// given number of characters,
//
// capacity: The number of characters the string is expected to initially contain.
//
// # Return Value
// 
// An initialized [NSMutableString] object with initial storage for `capacity`
// characters. The returned object might be different than the original
// receiver.
//
// # Discussion
// 
// The number of characters indicated by `capacity` is simply a hint to
// increase the efficiency of data storage. The value does limit the length of
// the string.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableString/init(capacity:)
func (m NSMutableString) InitWithCapacity(capacity uint) NSMutableString {
	rv := objc.Send[NSMutableString](m.ID, objc.Sel("initWithCapacity:"), capacity)
	return rv
}

// Adds to the end of the receiver the characters of a given string.
//
// aString: The string to append to the receiver. `aString` must not be `nil`
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableString/append(_:)
func (m NSMutableString) AppendString(aString string) {
	objc.Send[objc.ID](m.ID, objc.Sel("appendString:"), objc.String(aString))
}

// Transliterates the receiver by applying a specified ICU string transform.
//
// transform: The transformation to apply. For a list of possible values, see [String
// Transformations]. If the specified transform does not exist, the receiver
// is not modified, and this method returns [false].
// //
// [String Transformations]: https://developer.apple.com/documentation/Foundation/string-transformations
// [false]: https://developer.apple.com/documentation/Swift/false
//
// reverse: Whether an inverse transform should be used. If the specified transform
// does not have an inverse, the receiver is not modified, and this method
// returns [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
//
// range: The range of the string to transform. `range` must not exceed the bounds of
// the receiver.
//
// resultingRange: If the transform was successfully applied, upon return contains the range
// of the transformed string.
//
// # Return Value
// 
// [true] if the transform was successfully applied. Otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// In addition to the provided transformation constants, you may use any valid
// ICU transform ID as defined in the [ICU User Guide]. However, arbitrary ICU
// transform rules are not supported.
//
// [ICU User Guide]: http://userguide.icu-project.org/transforms/general
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableString/applyTransform(_:reverse:range:updatedRange:)
func (m NSMutableString) ApplyTransformReverseRangeUpdatedRange(transform NSStringTransform, reverse bool, range_ NSRange, resultingRange NSRangePointer) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("applyTransform:reverse:range:updatedRange:"), objc.String(string(transform)), reverse, range_, resultingRange)
	return rv
}

// Removes from the receiver the characters in a given range.
//
// range: The range of characters to delete. `range` must not exceed the bounds of
// the receiver.
//
// # Discussion
// 
// This method treats the length of the string as a valid range value that
// returns an empty string.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableString/deleteCharacters(in:)
func (m NSMutableString) DeleteCharactersInRange(range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("deleteCharactersInRange:"), range_)
}

// Inserts into the receiver the characters of a given string at a given
// location.
//
// aString: The string to insert into the receiver. `aString` must not be `nil`.
//
// loc: The location at which `aString` is inserted. The location must not exceed
// the bounds of the receiver.
//
// # Discussion
// 
// The new characters begin at `loc` and the existing characters from `loc` to
// the end are shifted by the length of `aString`.
// 
// This method treats the length of the string as a valid index value that
// returns an empty string.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableString/insert(_:at:)
func (m NSMutableString) InsertStringAtIndex(aString string, loc uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("insertString:atIndex:"), objc.String(aString), loc)
}

// Replaces the characters from `range` with those in `aString`.
//
// range: The range of characters to replace. `range` must not exceed the bounds of
// the receiver.
//
// aString: The string with which to replace the characters in `range`. `aString` must
// not be `nil`.
//
// # Discussion
// 
// This method treats the length of the string as a valid range value that
// returns an empty string.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableString/replaceCharacters(in:with:)
func (m NSMutableString) ReplaceCharactersInRangeWithString(range_ NSRange, aString string) {
	objc.Send[objc.ID](m.ID, objc.Sel("replaceCharactersInRange:withString:"), range_, objc.String(aString))
}

// Replaces all occurrences of a given string in a given range with another
// given string, returning the number of replacements.
//
// target: The string to replace.
//
// replacement: The string with which to replace `target`.
//
// options: A mask specifying search options. See [String Programming Guide] for
// details.
// 
// If `opts` is [NSBackwardsSearch], the search is done from the end of the
// range. If `opts` is [NSAnchoredSearch], only anchored (but potentially
// multiple) instances are replaced. [NSLiteralSearch] and
// [NSCaseInsensitiveSearch] also apply.
// //
// [String Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/introStrings.html#//apple_ref/doc/uid/10000035i
//
// searchRange: The range of characters to replace. `searchRange` must not exceed the
// bounds of the receiver. Specify `searchRange` as `NSMakeRange(0, [receiver
// length])` to process the entire string.
//
// # Return Value
// 
// The number of replacements made.
//
// # Discussion
// 
// This method treats the length of the string as a valid range value that
// returns an empty string.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableString/replaceOccurrences(of:with:options:range:)
func (m NSMutableString) ReplaceOccurrencesOfStringWithStringOptionsRange(target string, replacement string, options NSStringCompareOptions, searchRange NSRange) uint {
	rv := objc.Send[uint](m.ID, objc.Sel("replaceOccurrencesOfString:withString:options:range:"), objc.String(target), objc.String(replacement), options, searchRange)
	return rv
}

// Adds a constructed string to the receiver.
//
// format: A format string. See [Formatting String Objects] for more information. This
// value must not be `nil`.
// //
// [Formatting String Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/FormatStrings.html#//apple_ref/doc/uid/20000943
//
// # Discussion
// 
// Pass a comma-separated list of trailing variadic arguments to substitute
// into `format`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableString/appendFormat:
func (m NSMutableString) AppendFormat(format string) {
	objc.Send[objc.ID](m.ID, objc.Sel("appendFormat:"), objc.String(format))
}

// Returns an empty [NSMutableString] object with initial storage for a given
// number of characters.
//
// capacity: The number of characters the string is expected to initially contain.
//
// # Return Value
// 
// An empty [NSMutableString] object with initial storage for `capacity`
// characters.
//
// # Discussion
// 
// The number of characters indicated by `capacity` is simply a hint to
// increase the efficiency of data storage. The value does limit the length of
// the string.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableString/stringWithCapacity:
func (_NSMutableStringClass NSMutableStringClass) StringWithCapacity(capacity uint) NSMutableString {
	rv := objc.Send[objc.ID](objc.ID(_NSMutableStringClass.class), objc.Sel("stringWithCapacity:"), capacity)
	return NSMutableStringFromID(rv)
}

