// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSimpleCString] class.
var (
	_NSSimpleCStringClass     NSSimpleCStringClass
	_NSSimpleCStringClassOnce sync.Once
)

func getNSSimpleCStringClass() NSSimpleCStringClass {
	_NSSimpleCStringClassOnce.Do(func() {
		_NSSimpleCStringClass = NSSimpleCStringClass{class: objc.GetClass("NSSimpleCString")}
	})
	return _NSSimpleCStringClass
}

// GetNSSimpleCStringClass returns the class object for NSSimpleCString.
func GetNSSimpleCStringClass() NSSimpleCStringClass {
	return getNSSimpleCStringClass()
}

type NSSimpleCStringClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSimpleCStringClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSimpleCStringClass) Alloc() NSSimpleCString {
	rv := objc.Send[NSSimpleCString](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSSimpleCString
type NSSimpleCString struct {
	NSString
}

// NSSimpleCStringFromID constructs a [NSSimpleCString] from an objc.ID.
func NSSimpleCStringFromID(id objc.ID) NSSimpleCString {
	return NSSimpleCString{NSString: NSStringFromID(id)}
}

// Ensure NSSimpleCString implements INSSimpleCString.
var _ INSSimpleCString = NSSimpleCString{}

// An interface definition for the [NSSimpleCString] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSSimpleCString
type INSSimpleCString interface {
	INSString
}

// Init initializes the instance.
func (s NSSimpleCString) Init() NSSimpleCString {
	rv := objc.Send[NSSimpleCString](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSimpleCString) Autorelease() NSSimpleCString {
	rv := objc.Send[NSSimpleCString](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSimpleCString creates a new NSSimpleCString instance.
func NewNSSimpleCString() NSSimpleCString {
	class := getNSSimpleCStringClass()
	rv := objc.Send[NSSimpleCString](objc.ID(class.class), objc.Sel("new"))
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
func NewSimpleCStringWithBytesLengthEncoding(bytes []byte, encoding uint) NSSimpleCString {
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytes:length:encoding:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)), encoding)
	return NSSimpleCStringFromID(rv)
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
func NewSimpleCStringWithBytesNoCopyLengthEncodingFreeWhenDone(bytes unsafe.Pointer, len_ uint, encoding uint, freeBuffer bool) NSSimpleCString {
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytesNoCopy:length:encoding:freeWhenDone:"), bytes, len_, encoding, freeBuffer)
	return NSSimpleCStringFromID(rv)
}

// Initializes the receiver, a newly allocated [NSString] object, by
// converting the data in a given C-string from the default C-string encoding
// into the Unicode character encoding.
//
// # Discussion
//
// `cString` must be a zero-terminated C string in the default C string
// encoding, and may not be [NULL]. Returns an initialized object, which might
// be different from the original receiver.
//
// To create an immutable string from an immutable C string buffer, do not
// attempt to use this method. Instead, use
// [NSString.InitWithCStringNoCopyLengthFreeWhenDone].
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(CString:)-vkuo
func NewSimpleCStringWithCString(bytes string) NSSimpleCString {
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCString:"), unsafe.Pointer(unsafe.StringData(bytes+"\x00")))
	return NSSimpleCStringFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSString/init(cString:encoding:)-20f9h
func NewSimpleCStringWithCStringEncoding(nullTerminatedCString string, encoding uint) NSSimpleCString {
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCString:encoding:"), unsafe.Pointer(unsafe.StringData(nullTerminatedCString+"\x00")), encoding)
	return NSSimpleCStringFromID(rv)
}

// Initializes the receiver, a newly allocated [NSString] object, by
// converting the data in a given C-string from the default C-string encoding
// into the Unicode character encoding.
//
// # Discussion
//
// This method converts `length` * `sizeof(char)` bytes from `cString` and
// doesn’t stop short at a zero character. `cString` must contain bytes in
// the default C-string encoding and may not be [NULL]. Returns an initialized
// object, which might be different from the original receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(CString:length:)-5ure3
func NewSimpleCStringWithCStringLength(bytes unsafe.Pointer, length int) NSSimpleCString {
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCString:length:"), bytes, length)
	return NSSimpleCStringFromID(rv)
}

// Initializes the receiver, a newly allocated [NSString] object, by
// converting the data in a given C-string from the default C-string encoding
// into the Unicode character encoding.
//
// # Discussion
//
// This method converts `length` * `sizeof(char)` bytes from `cString` and
// doesn’t stop short at a zero character. `cString` must contain data in
// the default C-string encoding and may not be [NULL]. The receiver becomes
// the owner of `cString`; if `flag` is true it will free the memory when it
// no longer needs it, but if `flag` is false it won’t. Returns an
// initialized object, which might be different from the original receiver.
//
// You can use this method to create an immutable string from an immutable
// (`const char *`) C-string buffer. If you receive a warning message, you can
// disregard it; its purpose is simply to warn you that the C string passed as
// the method’s first argument may be modified. If you make certain the
// `freeWhenDone` argument to `initWithStringNoCopy` is false, the C string
// passed as the method’s first argument cannot be modified, so you can
// safely use `initWithStringNoCopy` to create an immutable string from an
// immutable (`const char *`) C-string buffer.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(CStringNoCopy:length:freeWhenDone:)-86dm2
func NewSimpleCStringWithCStringNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length int, freeBuffer bool) NSSimpleCString {
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCStringNoCopy:length:freeWhenDone:"), bytes, length, freeBuffer)
	return NSSimpleCStringFromID(rv)
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
func NewSimpleCStringWithCharactersLength(characters Unichar, length uint) NSSimpleCString {
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCharacters:length:"), characters, length)
	return NSSimpleCStringFromID(rv)
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
func NewSimpleCStringWithCharactersNoCopyLengthFreeWhenDone(characters Unichar, length uint, freeBuffer bool) NSSimpleCString {
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCharactersNoCopy:length:freeWhenDone:"), characters, length, freeBuffer)
	return NSSimpleCStringFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSString/init(coder:)
func NewSimpleCStringWithCoder(coder INSCoder) NSSimpleCString {
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSSimpleCStringFromID(rv)
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
func NewSimpleCStringWithContentsOfFile(path string) NSSimpleCString {
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:"), objc.String(path))
	return NSSimpleCStringFromID(rv)
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
func NewSimpleCStringWithContentsOfFileEncodingError(path string, enc uint) (NSSimpleCString, error) {
	var errorPtr objc.ID
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:encoding:error:"), objc.String(path), enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSSimpleCString{}, NSErrorFrom(errorPtr)
	}
	return NSSimpleCStringFromID(rv), nil
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
func NewSimpleCStringWithContentsOfFileUsedEncodingError(path string, enc uint) (NSSimpleCString, error) {
	var errorPtr objc.ID
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:usedEncoding:error:"), objc.String(path), enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSSimpleCString{}, NSErrorFrom(errorPtr)
	}
	return NSSimpleCStringFromID(rv), nil
}

// Initializes the receiver, a newly allocated [NSString] object, by reading
// data from the location named by a given URL.
//
// # Discussion
//
// Initializes the receiver, a newly allocated [NSString] object, by reading
// data from the location named by `aURL`. If the contents begin with a
// byte-order mark (`U+FEFF` or `U+FFFE`), interprets the contents as UTF-16
// code units; otherwise interprets the contents as data in the default C
// string encoding. Returns an initialized object, which might be different
// from the original receiver, or `nil` if the location can’t be opened.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOfURL:)
func NewSimpleCStringWithContentsOfURL(url INSURL) NSSimpleCString {
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	return NSSimpleCStringFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOf:encoding:)
func NewSimpleCStringWithContentsOfURLEncodingError(url INSURL, enc uint) (NSSimpleCString, error) {
	var errorPtr objc.ID
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:encoding:error:"), url, enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSSimpleCString{}, NSErrorFrom(errorPtr)
	}
	return NSSimpleCStringFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Foundation/NSString/init(contentsOf:usedEncoding:)
func NewSimpleCStringWithContentsOfURLUsedEncodingError(url INSURL, enc uint) (NSSimpleCString, error) {
	var errorPtr objc.ID
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:usedEncoding:error:"), url, enc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSSimpleCString{}, NSErrorFrom(errorPtr)
	}
	return NSSimpleCStringFromID(rv), nil
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
func NewSimpleCStringWithDataEncoding(data INSData, encoding uint) NSSimpleCString {
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:encoding:"), data, encoding)
	return NSSimpleCStringFromID(rv)
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
// This method invokes [NSString.InitWithFormatLocaleArguments] without
// applying any localization. This is useful, for example, when working with
// fixed-format representations of information that is written out and read
// back in at a later time.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/initWithFormat:
//
// [Formatting String Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/FormatStrings.html#//apple_ref/doc/uid/20000943
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
func NewSimpleCStringWithFormat(format string) NSSimpleCString {
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:"), objc.String(format))
	return NSSimpleCStringFromID(rv)
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
// This method invokes [NSString.InitWithFormatLocaleArguments] without
// applying any localization. This is useful, for example, when working with
// fixed-format representations of information that is written out and read
// back in at a later time.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/init(format:arguments:)
//
// [Formatting String Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/FormatStrings.html#//apple_ref/doc/uid/20000943
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
func NewSimpleCStringWithFormatArguments(format string, argList kernel.VaList) NSSimpleCString {
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:arguments:"), objc.String(format), argList)
	return NSSimpleCStringFromID(rv)
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
// Invokes [NSString.InitWithFormatLocaleArguments] with `locale` as the
// locale.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/initWithFormat:locale:
//
// [Formatting String Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/FormatStrings.html#//apple_ref/doc/uid/20000943
// [String Format Specifiers]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFStrings/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265
func NewSimpleCStringWithFormatLocale(format string, locale objectivec.IObject) NSSimpleCString {
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:locale:"), objc.String(format), locale)
	return NSSimpleCStringFromID(rv)
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
// locale (specified by user preferences), pass [NSLocale]
// [NSLocaleClass.CurrentLocale]]. To use the system locale, pass `nil`.
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
func NewSimpleCStringWithFormatLocaleArguments(format string, locale objectivec.IObject, argList kernel.VaList) NSSimpleCString {
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:locale:arguments:"), objc.String(format), locale, argList)
	return NSSimpleCStringFromID(rv)
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
func NewSimpleCStringWithString(aString string) NSSimpleCString {
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:"), objc.String(aString))
	return NSSimpleCStringFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSString/init(utf8String:)-vg2b
func NewSimpleCStringWithUTF8String(nullTerminatedCString string) NSSimpleCString {
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUTF8String:"), unsafe.Pointer(unsafe.StringData(nullTerminatedCString+"\x00")))
	return NSSimpleCStringFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:arguments:error:
func NewSimpleCStringWithValidatedFormatValidFormatSpecifiersArgumentsError(format string, validFormatSpecifiers string, argList kernel.VaList) (NSSimpleCString, error) {
	var errorPtr objc.ID
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:arguments:error:"), objc.String(format), objc.String(validFormatSpecifiers), argList, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSSimpleCString{}, NSErrorFrom(errorPtr)
	}
	return NSSimpleCStringFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:error:
func NewSimpleCStringWithValidatedFormatValidFormatSpecifiersError(format string, validFormatSpecifiers string) (NSSimpleCString, error) {
	var errorPtr objc.ID
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:error:"), objc.String(format), objc.String(validFormatSpecifiers), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSSimpleCString{}, NSErrorFrom(errorPtr)
	}
	return NSSimpleCStringFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:locale:arguments:error:
func NewSimpleCStringWithValidatedFormatValidFormatSpecifiersLocaleArgumentsError(format string, validFormatSpecifiers string, locale objectivec.IObject, argList kernel.VaList) (NSSimpleCString, error) {
	var errorPtr objc.ID
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:locale:arguments:error:"), objc.String(format), objc.String(validFormatSpecifiers), locale, argList, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSSimpleCString{}, NSErrorFrom(errorPtr)
	}
	return NSSimpleCStringFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Foundation/NSString/initWithValidatedFormat:validFormatSpecifiers:locale:error:
func NewSimpleCStringWithValidatedFormatValidFormatSpecifiersLocaleError(format string, validFormatSpecifiers string, locale objectivec.IObject) (NSSimpleCString, error) {
	var errorPtr objc.ID
	instance := getNSSimpleCStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValidatedFormat:validFormatSpecifiers:locale:error:"), objc.String(format), objc.String(validFormatSpecifiers), locale, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSSimpleCString{}, NSErrorFrom(errorPtr)
	}
	return NSSimpleCStringFromID(rv), nil
}
