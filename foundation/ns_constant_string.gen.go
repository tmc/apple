// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSConstantString] class.
var (
	_NSConstantStringClass     NSConstantStringClass
	_NSConstantStringClassOnce sync.Once
)

func getNSConstantStringClass() NSConstantStringClass {
	_NSConstantStringClassOnce.Do(func() {
		_NSConstantStringClass = NSConstantStringClass{class: objc.GetClass("NSConstantString")}
	})
	return _NSConstantStringClass
}

// GetNSConstantStringClass returns the class object for NSConstantString.
func GetNSConstantStringClass() NSConstantStringClass {
	return getNSConstantStringClass()
}

type NSConstantStringClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSConstantStringClass) Alloc() NSConstantString {
	rv := objc.Send[NSConstantString](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSConstantString
type NSConstantString struct {
	NSSimpleCString
}

// NSConstantStringFromID constructs a [NSConstantString] from an objc.ID.
func NSConstantStringFromID(id objc.ID) NSConstantString {
	return NSConstantString{NSSimpleCString: NSSimpleCStringFromID(id)}
}
// Ensure NSConstantString implements INSConstantString.
var _ INSConstantString = NSConstantString{}

// An interface definition for the [NSConstantString] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSConstantString
type INSConstantString interface {
	INSSimpleCString
}

// Init initializes the instance.
func (c NSConstantString) Init() NSConstantString {
	rv := objc.Send[NSConstantString](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSConstantString) Autorelease() NSConstantString {
	rv := objc.Send[NSConstantString](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSConstantString creates a new NSConstantString instance.
func NewNSConstantString() NSConstantString {
	class := getNSConstantStringClass()
	rv := objc.Send[NSConstantString](objc.ID(class.class), objc.Sel("new"))
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
func NewConstantStringWithBytesLengthEncoding(bytes []byte, encoding uint) NSConstantString {
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytes:length:encoding:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)), encoding)
	return NSConstantStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(bytesNoCopy:length:encoding:deallocator:)
func NewConstantStringWithBytesNoCopyLengthEncodingDeallocator(bytes unsafe.Pointer, len_ uint, encoding uint, deallocator unsafe.Pointer) NSConstantString {
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytesNoCopy:length:encoding:deallocator:"), bytes, len_, encoding, deallocator)
	return NSConstantStringFromID(rv)
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
func NewConstantStringWithBytesNoCopyLengthEncodingFreeWhenDone(bytes unsafe.Pointer, len_ uint, encoding uint, freeBuffer bool) NSConstantString {
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytesNoCopy:length:encoding:freeWhenDone:"), bytes, len_, encoding, freeBuffer)
	return NSConstantStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(cString:)
func NewConstantStringWithCString(bytes []byte) NSConstantString {
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCString:"), unsafe.Pointer(unsafe.SliceData(bytes)))
	return NSConstantStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(cString:encoding:)
func NewConstantStringWithCStringEncoding(nullTerminatedCString []byte, encoding uint) NSConstantString {
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCString:encoding:"), unsafe.Pointer(unsafe.SliceData(nullTerminatedCString)), encoding)
	return NSConstantStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(cString:length:)
func NewConstantStringWithCStringLength(bytes []byte, length uint) NSConstantString {
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCString:length:"), unsafe.Pointer(unsafe.SliceData(bytes)), length)
	return NSConstantStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(cStringNoCopy:length:freeWhenDone:)
func NewConstantStringWithCStringNoCopyLengthFreeWhenDone(bytes []byte, length uint, freeBuffer bool) NSConstantString {
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCStringNoCopy:length:freeWhenDone:"), unsafe.Pointer(unsafe.SliceData(bytes)), length, freeBuffer)
	return NSConstantStringFromID(rv)
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
func NewConstantStringWithCharactersLength(characters unsafe.Pointer, length uint) NSConstantString {
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCharacters:length:"), characters, length)
	return NSConstantStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(charactersNoCopy:length:deallocator:)
func NewConstantStringWithCharactersNoCopyLengthDeallocator(chars unsafe.Pointer, len_ uint, deallocator unsafe.Pointer) NSConstantString {
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCharactersNoCopy:length:deallocator:"), chars, len_, deallocator)
	return NSConstantStringFromID(rv)
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
func NewConstantStringWithCharactersNoCopyLengthFreeWhenDone(characters unsafe.Pointer, length uint, freeBuffer bool) NSConstantString {
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCharactersNoCopy:length:freeWhenDone:"), characters, length, freeBuffer)
	return NSConstantStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(coder:)
func NewConstantStringWithCoder(coder INSCoder) NSConstantString {
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSConstantStringFromID(rv)
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
func NewConstantStringWithContentsOfFile(path string) NSConstantString {
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:"), objc.String(path))
	return NSConstantStringFromID(rv)
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
func NewConstantStringWithContentsOfFileEncodingError(path string, enc uint) (NSConstantString, error) {
	var errorPtr objc.ID
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:encoding:error:"), objc.String(path), enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSConstantStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSConstantStringFromID(rv), nil
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
func NewConstantStringWithContentsOfFileUsedEncodingError(path string, enc unsafe.Pointer) (NSConstantString, error) {
	var errorPtr objc.ID
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:usedEncoding:error:"), objc.String(path), enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSConstantStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSConstantStringFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOf:)
func NewConstantStringWithContentsOfURL(url INSURL) NSConstantString {
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	return NSConstantStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOf:encoding:)
func NewConstantStringWithContentsOfURLEncodingError(url INSURL, enc uint) (NSConstantString, error) {
	var errorPtr objc.ID
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:encoding:error:"), url, enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSConstantStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSConstantStringFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOf:usedEncoding:)
func NewConstantStringWithContentsOfURLUsedEncodingError(url INSURL, enc unsafe.Pointer) (NSConstantString, error) {
	var errorPtr objc.ID
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:usedEncoding:error:"), url, enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSConstantStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSConstantStringFromID(rv), nil
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
func NewConstantStringWithDataEncoding(data INSData, encoding uint) NSConstantString {
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:encoding:"), data, encoding)
	return NSConstantStringFromID(rv)
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
func NewConstantStringWithFormat(format string) NSConstantString {
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:"), objc.String(format))
	return NSConstantStringFromID(rv)
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
func NewConstantStringWithFormatArguments(format string, argList unsafe.Pointer) NSConstantString {
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:arguments:"), objc.String(format), argList)
	return NSConstantStringFromID(rv)
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
func NewConstantStringWithFormatLocale(format string, locale objectivec.IObject) NSConstantString {
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:locale:"), objc.String(format), locale)
	return NSConstantStringFromID(rv)
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
func NewConstantStringWithFormatLocaleArguments(format string, locale objectivec.IObject, argList unsafe.Pointer) NSConstantString {
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:locale:arguments:"), objc.String(format), locale, argList)
	return NSConstantStringFromID(rv)
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
func NewConstantStringWithString(aString string) NSConstantString {
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:"), objc.String(aString))
	return NSConstantStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(utf8String:)
func NewConstantStringWithUTF8String(nullTerminatedCString []byte) NSConstantString {
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUTF8String:"), unsafe.Pointer(unsafe.SliceData(nullTerminatedCString)))
	return NSConstantStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:arguments:error:
func NewConstantStringWithValidatedFormatValidFormatSpecifiersArgumentsError(format string, validFormatSpecifiers string, argList unsafe.Pointer) (NSConstantString, error) {
	var errorPtr objc.ID
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:arguments:error:"), objc.String(format), objc.String(validFormatSpecifiers), argList, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSConstantStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSConstantStringFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:error:
func NewConstantStringWithValidatedFormatValidFormatSpecifiersError(format string, validFormatSpecifiers string) (NSConstantString, error) {
	var errorPtr objc.ID
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:error:"), objc.String(format), objc.String(validFormatSpecifiers), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSConstantStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSConstantStringFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:locale:arguments:error:
func NewConstantStringWithValidatedFormatValidFormatSpecifiersLocaleArgumentsError(format string, validFormatSpecifiers string, locale objectivec.IObject, argList unsafe.Pointer) (NSConstantString, error) {
	var errorPtr objc.ID
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:locale:arguments:error:"), objc.String(format), objc.String(validFormatSpecifiers), locale, argList, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSConstantStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSConstantStringFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:locale:error:
func NewConstantStringWithValidatedFormatValidFormatSpecifiersLocaleError(format string, validFormatSpecifiers string, locale objectivec.IObject) (NSConstantString, error) {
	var errorPtr objc.ID
	instance := getNSConstantStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:locale:error:"), objc.String(format), objc.String(validFormatSpecifiers), locale, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSConstantStringFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSConstantStringFromID(rv), nil
}

