// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Scanner] class.
var (
	_ScannerClass     ScannerClass
	_ScannerClassOnce sync.Once
)

func getScannerClass() ScannerClass {
	_ScannerClassOnce.Do(func() {
		_ScannerClass = ScannerClass{class: objc.GetClass("NSScanner")}
	})
	return _ScannerClass
}

// GetScannerClass returns the class object for NSScanner.
func GetScannerClass() ScannerClass {
	return getScannerClass()
}

type ScannerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc ScannerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc ScannerClass) Alloc() Scanner {
	rv := objc.Send[Scanner](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A string parser that scans for substrings or characters in a character set,
// and for numeric values from decimal, hexadecimal, and floating-point
// representations.
//
// # Overview
// 
// A [NSScanner] object interprets and converts the characters of a [String]
// into number and string values. You assign the scanner’s string when you
// create the scanner, and the scanner progresses through the characters of
// that string from beginning to end as you request items.
// 
// Because of the nature of class clusters, a scanner object isn’t an actual
// instance of the [NSScanner] class, but is one of its private subclasses.
// Although a scanner object’s class is private, its interface is public, as
// declared by this abstract superclass, [NSScanner]. The objects you create
// using this class are referred to as scanner objects (and when no confusion
// will result, merely as scanners).
// 
// To set a [NSScanner] object to ignore a set of characters as it scans the
// string, use the [CharactersToBeSkipped] property. Characters in the skip
// set are skipped over before the target is scanned. The default set of
// characters to skip is the whitespace and newline character set.
// 
// To retrieve the unscanned remainder of the string, use
// `scanner.StringXCUIElementTypeSubstring(scanner.ScanLocation())`.
//
// [String]: https://developer.apple.com/documentation/Swift/String
//
// # Creating a Scanner
//
//   - [Scanner.InitWithString]: Returns an [NSScanner] object initialized to scan a given string.
//
// # Getting a Scanner’s String
//
//   - [Scanner.String]: The string the scanner will scan.
//
// # Configuring a Scanner
//
//   - [Scanner.CaseSensitive]: Flag that indicates whether the receiver distinguishes case in the characters it scans.
//   - [Scanner.SetCaseSensitive]
//   - [Scanner.CharactersToBeSkipped]: Character set containing the characters the scanner ignores when looking for a scannable element.
//   - [Scanner.SetCharactersToBeSkipped]
//   - [Scanner.Locale]: The locale to use when scanning.
//   - [Scanner.SetLocale]
//
// # Scanning Numeric Values
//
//   - [Scanner.ScanHexDouble]: Scans for a double value from a hexadecimal representation, returning a found value by reference.
//   - [Scanner.ScanHexFloat]: Scans for a double value from a hexadecimal representation, returning a found value by reference.
//   - [Scanner.ScanHexLongLong]: Scans for a long long value from a hexadecimal representation, returning a found value by reference.
//   - [Scanner.ScanInteger]: Scans for an NSInteger value from a decimal representation, returning a found value by reference
//   - [Scanner.ScanLongLong]: Scans for a long long value from a decimal representation, returning a found value by reference.
//   - [Scanner.ScanUnsignedLongLong]: Scans for an unsigned long long value from a decimal representation, returning a found value by reference.
//
// # Monitoring Scanner Progress
//
//   - [Scanner.AtEnd]: Flag that indicates whether the receiver has exhausted all significant characters.
//
// # Instance Properties
//
//   - [Scanner.CurrentIndex]
//   - [Scanner.SetCurrentIndex]
//
// See: https://developer.apple.com/documentation/Foundation/Scanner
type Scanner struct {
	objectivec.Object
}

// ScannerFromID constructs a [Scanner] from an objc.ID.
//
// A string parser that scans for substrings or characters in a character set,
// and for numeric values from decimal, hexadecimal, and floating-point
// representations.
func ScannerFromID(id objc.ID) Scanner {
	return NSScanner{objectivec.Object{ID: id}}
}

// NSScannerFromID is an alias for [ScannerFromID] for cross-framework compatibility.
func NSScannerFromID(id objc.ID) Scanner { return ScannerFromID(id) }
// NOTE: Scanner adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [Scanner] class.
//
// # Creating a Scanner
//
//   - [IScanner.InitWithString]: Returns an [NSScanner] object initialized to scan a given string.
//
// # Getting a Scanner’s String
//
//   - [IScanner.String]: The string the scanner will scan.
//
// # Configuring a Scanner
//
//   - [IScanner.CaseSensitive]: Flag that indicates whether the receiver distinguishes case in the characters it scans.
//   - [IScanner.SetCaseSensitive]
//   - [IScanner.CharactersToBeSkipped]: Character set containing the characters the scanner ignores when looking for a scannable element.
//   - [IScanner.SetCharactersToBeSkipped]
//   - [IScanner.Locale]: The locale to use when scanning.
//   - [IScanner.SetLocale]
//
// # Scanning Numeric Values
//
//   - [IScanner.ScanHexDouble]: Scans for a double value from a hexadecimal representation, returning a found value by reference.
//   - [IScanner.ScanHexFloat]: Scans for a double value from a hexadecimal representation, returning a found value by reference.
//   - [IScanner.ScanHexLongLong]: Scans for a long long value from a hexadecimal representation, returning a found value by reference.
//   - [IScanner.ScanInteger]: Scans for an NSInteger value from a decimal representation, returning a found value by reference
//   - [IScanner.ScanLongLong]: Scans for a long long value from a decimal representation, returning a found value by reference.
//   - [IScanner.ScanUnsignedLongLong]: Scans for an unsigned long long value from a decimal representation, returning a found value by reference.
//
// # Monitoring Scanner Progress
//
//   - [IScanner.AtEnd]: Flag that indicates whether the receiver has exhausted all significant characters.
//
// # Instance Properties
//
//   - [IScanner.CurrentIndex]
//   - [IScanner.SetCurrentIndex]
//
// See: https://developer.apple.com/documentation/Foundation/Scanner
type IScanner interface {
	objectivec.IObject
	NSCopying

	// Topic: Creating a Scanner

	// Returns an [NSScanner] object initialized to scan a given string.
	InitWithString(string_ string) Scanner

	// Topic: Getting a Scanner’s String

	// The string the scanner will scan.
	String() string

	// Topic: Configuring a Scanner

	// Flag that indicates whether the receiver distinguishes case in the characters it scans.
	CaseSensitive() bool
	SetCaseSensitive(value bool)
	// Character set containing the characters the scanner ignores when looking for a scannable element.
	CharactersToBeSkipped() INSCharacterSet
	SetCharactersToBeSkipped(value INSCharacterSet)
	// The locale to use when scanning.
	Locale() objectivec.IObject
	SetLocale(value objectivec.IObject)

	// Topic: Scanning Numeric Values

	// Scans for a double value from a hexadecimal representation, returning a found value by reference.
	ScanHexDouble() (float64, bool)
	// Scans for a double value from a hexadecimal representation, returning a found value by reference.
	ScanHexFloat() (float32, bool)
	// Scans for a long long value from a hexadecimal representation, returning a found value by reference.
	ScanHexLongLong() (uint64, bool)
	// Scans for an NSInteger value from a decimal representation, returning a found value by reference
	ScanInteger() (int, bool)
	// Scans for a long long value from a decimal representation, returning a found value by reference.
	ScanLongLong() (int64, bool)
	// Scans for an unsigned long long value from a decimal representation, returning a found value by reference.
	ScanUnsignedLongLong() (uint64, bool)

	// Topic: Monitoring Scanner Progress

	// Flag that indicates whether the receiver has exhausted all significant characters.
	AtEnd() bool

	// Topic: Instance Properties

	CurrentIndex() int
	SetCurrentIndex(value int)

	// A value indicating that a requested item couldn’t be found or doesn’t exist.
	NSNotFound() int
	SetNSNotFound(value int)
}

// Init initializes the instance.
func (s Scanner) Init() Scanner {
	rv := objc.Send[Scanner](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s Scanner) Autorelease() Scanner {
	rv := objc.Send[Scanner](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewScanner creates a new Scanner instance.
func NewScanner() Scanner {
	class := getScannerClass()
	rv := objc.Send[Scanner](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an [NSScanner] object initialized to scan a given string.
//
// string: The string to scan.
//
// # Return Value
// 
// An [NSScanner] object initialized to scan `aString` from the beginning. The
// returned object might be different than the original receiver.
//
// See: https://developer.apple.com/documentation/Foundation/Scanner/init(string:)
func NewScannerWithString(string_ string) Scanner {
	instance := getScannerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:"), objc.String(string_))
	return ScannerFromID(rv)
}

// Returns an [NSScanner] object initialized to scan a given string.
//
// string: The string to scan.
//
// # Return Value
// 
// An [NSScanner] object initialized to scan `aString` from the beginning. The
// returned object might be different than the original receiver.
//
// See: https://developer.apple.com/documentation/Foundation/Scanner/init(string:)
func (s Scanner) InitWithString(string_ string) Scanner {
	rv := objc.Send[Scanner](s.ID, objc.Sel("initWithString:"), objc.String(string_))
	return rv
}
// Scans for a double value from a hexadecimal representation, returning a
// found value by reference.
//
// result: Upon return, contains the scanned value. Contains `HUGE_VAL` or
// `–HUGE_VAL` on overflow, or `0.0` on underflow.
//
// # Return Value
// 
// [true] if the receiver finds a valid double-point representation, otherwise
// [false]. Overflow or underflow are both considered valid floating-point
// representations.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This corresponds to `%a` or `%A` formatting. The hexadecimal double
// representation must be preceded by `0x` or `0X`.
// 
// Skips past excess digits in the case of overflow, so the scanner’s
// position is past the entire floating-point representation.
// 
// Invoke this method with [NULL] as `result` to simply scan past a
// hexadecimal double representation.
//
// See: https://developer.apple.com/documentation/Foundation/Scanner/scanHexDouble(_:)
func (s Scanner) ScanHexDouble() (float64, bool) {
	var result float64
	rv := objc.Send[bool](s.ID, objc.Sel("scanHexDouble:"), unsafe.Pointer(&result))
	return result, rv
}
// Scans for a double value from a hexadecimal representation, returning a
// found value by reference.
//
// result: Upon return, contains the scanned value. Contains `HUGE_VAL` or
// `–HUGE_VAL` on overflow, or `0.0` on underflow.
//
// # Return Value
// 
// [true] if the receiver finds a valid float-point representation, otherwise
// [false]. Overflow or underflow are both considered valid floating-point
// representations.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This corresponds to `%a` or `%A` formatting. The hexadecimal float
// representation must be preceded by `0x` or `0X`.
// 
// Skips past excess digits in the case of overflow, so the scanner’s
// position is past the entire floating-point representation.
// 
// Invoke this method with [NULL] as `result` to simply scan past a
// hexadecimal float representation.
//
// See: https://developer.apple.com/documentation/Foundation/Scanner/scanHexFloat(_:)
func (s Scanner) ScanHexFloat() (float32, bool) {
	var result float32
	rv := objc.Send[bool](s.ID, objc.Sel("scanHexFloat:"), unsafe.Pointer(&result))
	return result, rv
}
// Scans for a long long value from a hexadecimal representation, returning a
// found value by reference.
//
// result: Upon return, contains the scanned value. Contains `HUGE_VAL` or
// `–HUGE_VAL` on overflow.
//
// # Return Value
// 
// [true] if the receiver finds a valid hexadecimal long long representation,
// otherwise [false]. Overflow is considered a valid hexadecimal long long
// representation.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The hexadecimal integer representation may optionally be preceded by `0x`
// or `0X`. Skips past excess digits in the case of overflow, so the
// receiver’s position is past the entire hexadecimal representation.
// 
// Invoke this method with [NULL] as `result` to simply scan past a
// hexadecimal long long representation.
//
// See: https://developer.apple.com/documentation/Foundation/Scanner/scanHexInt64(_:)
func (s Scanner) ScanHexLongLong() (uint64, bool) {
	var result uint64
	rv := objc.Send[bool](s.ID, objc.Sel("scanHexLongLong:"), unsafe.Pointer(&result))
	return result, rv
}
// Scans for an NSInteger value from a decimal representation, returning a
// found value by reference
//
// result: Upon return, contains the scanned value. Contains `INT_MAX` or `INT_MIN` on
// overflow.
//
// # Return Value
// 
// [true] if the receiver finds a valid integer representation, otherwise
// [false]. Overflow is considered a valid integer representation.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Skips past excess digits in the case of overflow, so the receiver’s
// position is past the entire integer representation.
// 
// Invoke this method with [NULL] as `value` to simply scan past a decimal
// integer representation.
//
// See: https://developer.apple.com/documentation/Foundation/Scanner/scanInt(_:)
func (s Scanner) ScanInteger() (int, bool) {
	var result int
	rv := objc.Send[bool](s.ID, objc.Sel("scanInteger:"), unsafe.Pointer(&result))
	return result, rv
}
// Scans for a long long value from a decimal representation, returning a
// found value by reference.
//
// result: Upon return, contains the scanned value. Contains `LLONG_MAX` or
// `LLONG_MIN` on overflow.
//
// # Return Value
// 
// [true] if the receiver finds a valid decimal integer representation,
// otherwise [false]. Overflow is considered a valid decimal integer
// representation.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// All overflow digits are skipped. Skips past excess digits in the case of
// overflow, so the receiver’s position is past the entire decimal
// representation.
// 
// Invoke this method with [NULL] as `longLongValue` to simply scan past a
// long decimal integer representation.
//
// See: https://developer.apple.com/documentation/Foundation/Scanner/scanInt64(_:)
func (s Scanner) ScanLongLong() (int64, bool) {
	var result int64
	rv := objc.Send[bool](s.ID, objc.Sel("scanLongLong:"), unsafe.Pointer(&result))
	return result, rv
}
// Scans for an unsigned long long value from a decimal representation,
// returning a found value by reference.
//
// result: Upon return, contains the scanned value. Contains `ULLONG_MAX` on overflow.
//
// # Return Value
// 
// [true] if the receiver finds a valid decimal integer representation,
// otherwise [false]. Overflow is considered a valid decimal integer
// representation.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// All overflow digits are skipped. Skips past excess digits in the case of
// overflow, so the receiver’s position is past the entire decimal
// representation.
// 
// Invoke this method with [NULL] as `unsignedLongLongValue` to simply scan
// past an unsigned long decimal integer representation.
//
// See: https://developer.apple.com/documentation/Foundation/Scanner/scanUnsignedLongLong(_:)
func (s Scanner) ScanUnsignedLongLong() (uint64, bool) {
	var result uint64
	rv := objc.Send[bool](s.ID, objc.Sel("scanUnsignedLongLong:"), unsafe.Pointer(&result))
	return result, rv
}

// Returns an [NSScanner] object that scans a given string according to the
// user’s default locale.
//
// string: The string to scan.
//
// # Return Value
// 
// An [NSScanner] object that scans `aString` according to the user’s
// default locale.
//
// # Discussion
// 
// Sets the string to scan by invoking [InitWithString] with `aString`. The
// locale is set with [NSScanner].
//
// See: https://developer.apple.com/documentation/Foundation/Scanner/localizedScanner(with:)
func (_ScannerClass ScannerClass) LocalizedScannerWithString(string_ string) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ScannerClass.class), objc.Sel("localizedScannerWithString:"), objc.String(string_))
	return objectivec.Object{ID: rv}
}
// Returns an [NSScanner] object that scans a given string.
//
// string: The string to scan.
//
// # Return Value
// 
// An [NSScanner] object that scans `aString`.
//
// # Discussion
// 
// Sets the string to scan by invoking [InitWithString] with `aString`.
//
// See: https://developer.apple.com/documentation/Foundation/NSScanner/scannerWithString:
func (_ScannerClass ScannerClass) ScannerWithString(string_ string) Scanner {
	rv := objc.Send[objc.ID](objc.ID(_ScannerClass.class), objc.Sel("scannerWithString:"), objc.String(string_))
	return NSScannerFromID(rv)
}

// The string the scanner will scan.
//
// See: https://developer.apple.com/documentation/Foundation/Scanner/string
func (s Scanner) String() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("string"))
	return NSStringFromID(rv).String()
}
// Flag that indicates whether the receiver distinguishes case in the
// characters it scans.
//
// # Discussion
// 
// [true] if the receiver distinguishes case in the characters it scans,
// otherwise [false]. The default value is [false]. Note that case sensitivity
// doesn’t apply to the characters to be skipped.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Scanner/caseSensitive
func (s Scanner) CaseSensitive() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("caseSensitive"))
	return rv
}
func (s Scanner) SetCaseSensitive(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setCaseSensitive:"), value)
}
// Character set containing the characters the scanner ignores when looking
// for a scannable element.
//
// # Discussion
// 
// Characters to be skipped are skipped prior to the scanner examining the
// target. For example, if a scanner ignores spaces and you send it a
// [scanInt32(_:)] message, it skips spaces until it finds a decimal digit or
// other character. While an element is being scanned, no characters are
// skipped. If you scan for something made of characters in the set to be
// skipped (for example, using [scanInt32(_:)] when the set of characters to
// be skipped is the decimal digits), the result is undefined.
// 
// The characters to be skipped are treated as single values. A scanner
// doesn’t apply its case sensitivity setting to these characters and
// doesn’t attempt to match composed character sequences with anything in
// the set of characters to be skipped (though it does match pre-composed
// characters individually). If you want to skip all vowels while scanning a
// string, for example, you can set the characters to be skipped to those in
// the string “AEIOUaeiou” (plus any accented variants with pre-composed
// characters).
// 
// The default set to skip is the whitespace and newline character set.
//
// [scanInt32(_:)]: https://developer.apple.com/documentation/Foundation/Scanner/scanInt32(_:)
//
// See: https://developer.apple.com/documentation/Foundation/Scanner/charactersToBeSkipped
func (s Scanner) CharactersToBeSkipped() INSCharacterSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("charactersToBeSkipped"))
	return NSCharacterSetFromID(objc.ID(rv))
}
func (s Scanner) SetCharactersToBeSkipped(value INSCharacterSet) {
	objc.Send[struct{}](s.ID, objc.Sel("setCharactersToBeSkipped:"), value)
}
// The locale to use when scanning.
//
// # Discussion
// 
// A scanner’s locale affects the way it interprets numeric values from the
// string. In particular, a scanner uses the locale’s decimal separator to
// distinguish the integer and fractional parts of floating-point
// representations. A scanner with no locale set uses non-localized values.
// New scanners have no locale by default.
//
// See: https://developer.apple.com/documentation/Foundation/Scanner/locale
func (s Scanner) Locale() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("locale"))
	return objectivec.Object{ID: rv}
}
func (s Scanner) SetLocale(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setLocale:"), value)
}
// Flag that indicates whether the receiver has exhausted all significant
// characters.
//
// # Discussion
// 
// [true] if the receiver has exhausted all significant characters in its
// string, otherwise [false].
// 
// If only characters from the set to be skipped remain, returns [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Scanner/isAtEnd
func (s Scanner) AtEnd() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isAtEnd"))
	return rv
}
// See: https://developer.apple.com/documentation/foundation/scanner/currentindex
func (s Scanner) CurrentIndex() int {
	rv := objc.Send[int](s.ID, objc.Sel("currentIndex"))
	return rv
}
func (s Scanner) SetCurrentIndex(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setCurrentIndex:"), value)
}
// A value indicating that a requested item couldn’t be found or doesn’t
// exist.
//
// See: https://developer.apple.com/documentation/foundation/nsnotfound-4qp9h
func (s Scanner) NSNotFound() int {
	rv := objc.Send[int](s.ID, objc.Sel("NSNotFound"))
	return rv
}
func (s Scanner) SetNSNotFound(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setNSNotFound:"), value)
}

			// Protocol methods for NSCopying
			

