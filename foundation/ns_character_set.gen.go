// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCharacterSet] class.
var (
	_NSCharacterSetClass     NSCharacterSetClass
	_NSCharacterSetClassOnce sync.Once
)

func getNSCharacterSetClass() NSCharacterSetClass {
	_NSCharacterSetClassOnce.Do(func() {
		_NSCharacterSetClass = NSCharacterSetClass{class: objc.GetClass("NSCharacterSet")}
	})
	return _NSCharacterSetClass
}

// GetNSCharacterSetClass returns the class object for NSCharacterSet.
func GetNSCharacterSetClass() NSCharacterSetClass {
	return getNSCharacterSetClass()
}

type NSCharacterSetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSCharacterSetClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCharacterSetClass) Alloc() NSCharacterSet {
	rv := objc.Send[NSCharacterSet](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object representing a fixed set of Unicode character values for use in
// search operations.
//
// # Overview
// 
// In Swift, this bridges to a [CharacterSet]; use [NSCharacterSet] when you
// need reference semantics or other Foundation-specific behavior.
// 
// An [NSCharacterSet] object represents a set of Unicode-compliant
// characters. [NSString] and [NSScanner] objects use [NSCharacterSet] objects
// to group characters together for searching operations, so that they can
// find any of a particular set of characters during a search. The cluster’s
// two public classes, [NSCharacterSet] and [NSMutableCharacterSet], declare
// the programmatic interface for static and dynamic character sets,
// respectively.
// 
// The objects you create using these classes are referred to as character set
// objects (and when no confusion will result, merely as character sets).
// Because of the nature of class clusters, character set objects aren’t
// actual instances of the [NSCharacterSet] or [NSMutableCharacterSet] classes
// but of one of their private subclasses. Although a character set object’s
// class is private, its interface is public, as declared by these abstract
// superclasses, [NSCharacterSet] and [NSMutableCharacterSet]. The character
// set classes adopt the [NSCopying] and [NSMutableCopying] protocols, making
// it convenient to convert a character set of one type to the other.
// 
// The [NSCharacterSet] class declares the programmatic interface for an
// object that manages a set of Unicode characters (see the [NSString] class
// cluster specification for information on Unicode). [NSCharacterSet]’s
// principal primitive method, [NSCharacterSet.CharacterIsMember], provides the basis for all
// other instance methods in its interface. A subclass of [NSCharacterSet]
// needs only to implement this method, plus [NSCharacterSet.MutableCopyWithZone], for proper
// behavior. For optimal performance, a subclass should also override
// [NSCharacterSet.BitmapRepresentation], which otherwise works by invoking
// [NSCharacterSet.CharacterIsMember] for every possible Unicode value.
// 
// [NSCharacterSet] is “toll-free bridged” with its Core Foundation
// counterpart, [CFCharacterSet]. See [Toll-Free Bridging] for more
// information on toll-free bridging.
//
// [CFCharacterSet]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSet
// [CharacterSet]: https://developer.apple.com/documentation/Foundation/CharacterSet
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
//
// # Creating and Managing Character Sets as Bitmap Representations
//
//   - [NSCharacterSet.BitmapRepresentation]: An [NSData] object encoding the receiver in binary format.
//
// # Inverting a Character Set
//
//   - [NSCharacterSet.InvertedSet]: A character set containing only characters that don’t exist in the receiver.
//
// # Testing Set Membership
//
//   - [NSCharacterSet.CharacterIsMember]: Returns a Boolean value that indicates whether a given character is in the receiver.
//   - [NSCharacterSet.HasMemberInPlane]: Returns a Boolean value that indicates whether the receiver has at least one member in a given character plane.
//   - [NSCharacterSet.IsSupersetOfSet]: Returns a Boolean value that indicates whether the receiver is a superset of another given character set.
//   - [NSCharacterSet.LongCharacterIsMember]: Returns a Boolean value that indicates whether a given long character is a member of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet
type NSCharacterSet struct {
	objectivec.Object
}

// NSCharacterSetFromID constructs a [NSCharacterSet] from an objc.ID.
//
// An object representing a fixed set of Unicode character values for use in
// search operations.
func NSCharacterSetFromID(id objc.ID) NSCharacterSet {
	return NSCharacterSet{objectivec.Object{ID: id}}
}
// NOTE: NSCharacterSet adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCharacterSet] class.
//
// # Creating and Managing Character Sets as Bitmap Representations
//
//   - [INSCharacterSet.BitmapRepresentation]: An [NSData] object encoding the receiver in binary format.
//
// # Inverting a Character Set
//
//   - [INSCharacterSet.InvertedSet]: A character set containing only characters that don’t exist in the receiver.
//
// # Testing Set Membership
//
//   - [INSCharacterSet.CharacterIsMember]: Returns a Boolean value that indicates whether a given character is in the receiver.
//   - [INSCharacterSet.HasMemberInPlane]: Returns a Boolean value that indicates whether the receiver has at least one member in a given character plane.
//   - [INSCharacterSet.IsSupersetOfSet]: Returns a Boolean value that indicates whether the receiver is a superset of another given character set.
//   - [INSCharacterSet.LongCharacterIsMember]: Returns a Boolean value that indicates whether a given long character is a member of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet
type INSCharacterSet interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSMutableCopying
	NSSecureCoding

	// Topic: Creating and Managing Character Sets as Bitmap Representations

	// An [NSData] object encoding the receiver in binary format.
	BitmapRepresentation() INSData

	// Topic: Inverting a Character Set

	// A character set containing only characters that don’t exist in the receiver.
	InvertedSet() INSCharacterSet

	// Topic: Testing Set Membership

	// Returns a Boolean value that indicates whether a given character is in the receiver.
	CharacterIsMember(aCharacter uint16) bool
	// Returns a Boolean value that indicates whether the receiver has at least one member in a given character plane.
	HasMemberInPlane(thePlane uint8) bool
	// Returns a Boolean value that indicates whether the receiver is a superset of another given character set.
	IsSupersetOfSet(theOtherSet INSCharacterSet) bool
	// Returns a Boolean value that indicates whether a given long character is a member of the receiver.
	LongCharacterIsMember(theLongChar uint32) bool
}

// Init initializes the instance.
func (c NSCharacterSet) Init() NSCharacterSet {
	rv := objc.Send[NSCharacterSet](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCharacterSet) Autorelease() NSCharacterSet {
	rv := objc.Send[NSCharacterSet](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCharacterSet creates a new NSCharacterSet instance.
func NewNSCharacterSet() NSCharacterSet {
	class := getNSCharacterSetClass()
	rv := objc.Send[NSCharacterSet](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a character set containing characters determined by a given bitmap
// representation.
//
// data: A bitmap representation of a character set.
//
// # Return Value
// 
// A character set containing characters determined by `data`.
//
// # Discussion
// 
// This method is useful for creating a character set object with data from a
// file or other external data source.
// 
// A raw bitmap representation of a character set is a byte array with the
// first 2^16 bits (that is, 8192 bytes) representing the code point range of
// the the Basic Multilingual Plane (BMP), such that the value of the bit at
// position n represents the presence in the character set of the character
// with decimal Unicode value n. A bitmap representation may contain zero to
// sixteen additional 8192 byte segments to for each additional Unicode plane
// containing a character in a character set, with each 8192 byte segment
// prepended with a single plane index byte.
// 
// To add a character in the Basic Multilingual Plane (BMP) with decimal
// Unicode value n to a raw bitmap representation, you might do the following:
// 
// To remove that character:
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/init(bitmapRepresentation:)
func NewCharacterSetWithBitmapRepresentation(data INSData) NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(getNSCharacterSetClass().class), objc.Sel("characterSetWithBitmapRepresentation:"), data)
	return NSCharacterSetFromID(rv)
}

// Returns a character set containing the characters in a given string.
//
// aString: A string containing characters for the new character set.
//
// # Return Value
// 
// A character set containing the characters in `aString`. Returns an empty
// character set if `aString` is empty.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/init(charactersIn:)
func NewCharacterSetWithCharactersInString(aString string) NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(getNSCharacterSetClass().class), objc.Sel("characterSetWithCharactersInString:"), objc.String(aString))
	return NSCharacterSetFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/init(coder:)
func NewCharacterSetWithCoder(coder INSCoder) NSCharacterSet {
	instance := getNSCharacterSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSCharacterSetFromID(rv)
}

// Returns a character set read from the bitmap representation stored in the
// file a given path.
//
// fName: A path to a file containing a bitmap representation of a character set. The
// path name must end with the extension `XCUIElementTypeBitmap`.
//
// # Return Value
// 
// A character set read from the bitmap representation stored in the file at
// `path`.
//
// # Discussion
// 
// This method doesn’t use filenames to check for the uniqueness of the
// character sets it creates. To prevent duplication of character sets in
// memory, cache them and make them available through an API that checks
// whether the requested set has already been loaded.
// 
// To read a bitmap representation from any file, use the [NSData]
// method[DataWithContentsOfFileOptionsError] and pass the result to
// [CharacterSetWithBitmapRepresentation].
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/init(contentsOfFile:)
func NewCharacterSetWithContentsOfFile(fName string) NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(getNSCharacterSetClass().class), objc.Sel("characterSetWithContentsOfFile:"), objc.String(fName))
	return NSCharacterSetFromID(rv)
}

// Returns a character set containing characters with Unicode values in a
// given range.
//
// aRange: A range of Unicode values. `aRange.Location()` is the value of the first
// character to return; `aRange.Location() + aRange.Length() – 1` is the
// value of the last.
//
// # Return Value
// 
// A character set containing characters whose Unicode values are given by
// `aRange`. If `aRange.Length()` is `0`, returns an empty character set.
//
// # Discussion
// 
// This code excerpt creates a character set object containing the lowercase
// English alphabetic characters:
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/init(range:)
func NewCharacterSetWithRange(aRange NSRange) NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(getNSCharacterSetClass().class), objc.Sel("characterSetWithRange:"), aRange)
	return NSCharacterSetFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/init(coder:)
func (c NSCharacterSet) InitWithCoder(coder INSCoder) NSCharacterSet {
	rv := objc.Send[NSCharacterSet](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Returns a Boolean value that indicates whether a given character is in the
// receiver.
//
// aCharacter: The character to test for membership of the receiver.
//
// # Return Value
// 
// [true] if `aCharacter` is in the receiving character set, otherwise
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/characterIsMember(_:)
func (c NSCharacterSet) CharacterIsMember(aCharacter uint16) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("characterIsMember:"), aCharacter)
	return rv
}
// Returns a Boolean value that indicates whether the receiver has at least
// one member in a given character plane.
//
// thePlane: A character plane.
//
// # Return Value
// 
// [true] if the receiver has at least one member in `thePlane`, otherwise
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method makes it easier to find the plane containing the members of the
// current character set. The Basic Multilingual Plane (BMP) is plane `0`.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/hasMemberInPlane(_:)
func (c NSCharacterSet) HasMemberInPlane(thePlane uint8) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasMemberInPlane:"), thePlane)
	return rv
}
// Returns a Boolean value that indicates whether the receiver is a superset
// of another given character set.
//
// theOtherSet: A character set.
//
// # Return Value
// 
// [true] if the receiver is a superset of `theOtherSet`, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/isSuperset(of:)
func (c NSCharacterSet) IsSupersetOfSet(theOtherSet INSCharacterSet) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isSupersetOfSet:"), theOtherSet)
	return rv
}
// Returns a Boolean value that indicates whether a given long character is a
// member of the receiver.
//
// theLongChar: A UTF32 character.
//
// # Return Value
// 
// [true] if `theLongChar` is in the receiver, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method supports the specification of 32-bit characters.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/longCharacterIsMember(_:)
func (c NSCharacterSet) LongCharacterIsMember(theLongChar uint32) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("longCharacterIsMember:"), theLongChar)
	return rv
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (c NSCharacterSet) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// An [NSData] object encoding the receiver in binary format.
//
// # Discussion
// 
// This format is suitable for saving to a file or otherwise transmitting or
// archiving.
// 
// A raw bitmap representation of a character set is a byte array with the
// first 2^16 bits (that is, 8192 bytes) representing the code point range of
// the the Basic Multilingual Plane (BMP), such that the value of the bit at
// position n represents the presence in the character set of the character
// with decimal Unicode value n. A bitmap representation may contain zero to
// sixteen additional 8192 byte segments to for each additional Unicode plane
// containing a character in a character set, with each 8192 byte segment
// prepended with a single plane index byte.
// 
// For example, a character set containing only Basic Latin (ASCII)
// characters, which are contained by the Basic Multilingual Plane (BMP, plane
// 0), has a bitmap representation with a size of 8192 bytes, whereas a
// character set containing both Basic Latin (ASCII) characters and emoji
// characters, which are contained by the Supplementary Multilingual Plane
// (SMP, plane 1), has a bitmap representation with a size of 16385 bytes
// (8192 bytes for BMP, followed by the byte `0x01` for the plane index of
// SMP, followed by 8192 bytes for SMP).
// 
// To test for the presence of a character in the Basic Multilingual Plane
// (BMP) with decimal Unicode value n in a raw bitmap representation, you
// might do the following:
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/bitmapRepresentation
func (c NSCharacterSet) BitmapRepresentation() INSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("bitmapRepresentation"))
	return NSDataFromID(objc.ID(rv))
}
// A character set containing only characters that don’t exist in the
// receiver.
//
// # Discussion
// 
// Using the inverse of an immutable character set is much more efficient than
// inverting a mutable character set.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/inverted
func (c NSCharacterSet) InvertedSet() INSCharacterSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("invertedSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}

// A character set containing the characters in Unicode General Categories L*,
// M*, and N*.
//
// # Return Value
// 
// A character set containing all the alphanumeric characters.
// 
// # Discussion
// 
// Informally, this set is the set of all characters used as basic units of
// alphabets, syllabaries, ideographs, and digits.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/alphanumerics
func (_NSCharacterSetClass NSCharacterSetClass) AlphanumericCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("alphanumericCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}
// A character set containing the characters in Unicode General Category Lt.
//
// # Return Value
// 
// A character set containing all the capitalized letter characters.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/capitalizedLetters
func (_NSCharacterSetClass NSCharacterSetClass) CapitalizedLetterCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("capitalizedLetterCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}
// A character set containing the characters in Unicode General Category Cc
// and Cf.
//
// # Return Value
// 
// A character set containing all the control characters.
// 
// # Discussion
// 
// These characters include, for example, the soft hyphen (`U+00AD`), control
// characters to support bi-directional text, and IETF language tag
// characters.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/controlCharacters
func (_NSCharacterSetClass NSCharacterSetClass) ControlCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("controlCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}
// A character set containing the characters in the category of Decimal
// Numbers.
//
// # Return Value
// 
// A character set containing all the decimal digit characters.
// 
// # Discussion
// 
// Informally, this set is the set of all characters used to represent the
// decimal values `0` through `9`. These characters include, for example, the
// decimal digits of the Indic scripts and Arabic.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/decimalDigits
func (_NSCharacterSetClass NSCharacterSetClass) DecimalDigitCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("decimalDigitCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}
// A character set containing individual Unicode characters that can also be
// represented as composed character sequences (such as for letters with
// accents), by the definition of “standard decomposition” in version 3.2
// of the Unicode character encoding standard.
//
// # Return Value
// 
// A character set containing all the decomposable characters.
// 
// # Discussion
// 
// These characters include compatibility characters as well as pre-composed
// characters.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/decomposables
func (_NSCharacterSetClass NSCharacterSetClass) DecomposableCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("decomposableCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}
// A character set containing values in the category of Non-Characters or that
// have not yet been defined in version 3.2 of the Unicode standard.
//
// # Return Value
// 
// A character set containing all the illegal characters.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/illegalCharacters
func (_NSCharacterSetClass NSCharacterSetClass) IllegalCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("illegalCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}
// A character set containing the characters in Unicode General Category L* &
// M*.
//
// # Return Value
// 
// A character set containing all the letter characters.
// 
// # Discussion
// 
// Informally, this set is the set of all characters used as letters of
// alphabets and ideographs.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/letters
func (_NSCharacterSetClass NSCharacterSetClass) LetterCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("letterCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}
// A character set containing the characters in Unicode General Category Ll.
//
// # Return Value
// 
// A character set containing all the lowercase letter characters.
// 
// # Discussion
// 
// Informally, this set is the set of all characters used as lowercase letters
// in alphabets that make case distinctions.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/lowercaseLetters
func (_NSCharacterSetClass NSCharacterSetClass) LowercaseLetterCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("lowercaseLetterCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}
// A character set containing the newline characters (`U+000A` ~ `U+000D`,
// `U+0085`, `U+2028`, and `U+2029`).
//
// # Return Value
// 
// A character set containing all the newline characters.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/newlines
func (_NSCharacterSetClass NSCharacterSetClass) NewlineCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("newlineCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}
// A character set containing the characters in Unicode General Category M*.
//
// # Return Value
// 
// A character set containing all the non-base characters.
// 
// # Discussion
// 
// This set is also defined as all legal Unicode characters with a non-spacing
// priority greater than `0`. Informally, this set is the set of all
// characters used as modifiers of base characters.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/nonBaseCharacters
func (_NSCharacterSetClass NSCharacterSetClass) NonBaseCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("nonBaseCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}
// A character set containing the characters in Unicode General Category P*.
//
// # Return Value
// 
// A character set containing all the punctuation characters.
// 
// # Discussion
// 
// Informally, this set is the set of all non-whitespace characters used to
// separate linguistic units in scripts, such as periods, dashes, parentheses,
// and so on.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/punctuationCharacters
func (_NSCharacterSetClass NSCharacterSetClass) PunctuationCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("punctuationCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}
// A character set containing the characters in Unicode General Category S*.
//
// # Return Value
// 
// A character set containing all the symbol characters.
// 
// # Discussion
// 
// These characters include, for example, the dollar sign (`$`) and the plus
// (`+`) sign.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/symbols
func (_NSCharacterSetClass NSCharacterSetClass) SymbolCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("symbolCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}
// A character set containing the characters in Unicode General Category Lu
// and Lt.
//
// # Return Value
// 
// A character set containing all the uppercase letter characters.
// 
// # Discussion
// 
// Informally, this set is the set of all characters used as uppercase letters
// in alphabets that make case distinctions.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/uppercaseLetters
func (_NSCharacterSetClass NSCharacterSetClass) UppercaseLetterCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("uppercaseLetterCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}
// A character set containing characters in Unicode General Category Z*,
// `U+000A` ~ `U+000D`, and `U+0085`.
//
// # Return Value
// 
// A character set containing all the whitespace and newline characters.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/whitespacesAndNewlines
func (_NSCharacterSetClass NSCharacterSetClass) WhitespaceAndNewlineCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("whitespaceAndNewlineCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}
// A character set containing the characters in Unicode General Category Zs
// and `CHARACTER TABULATION` (`U+0009`).
//
// # Return Value
// 
// A character set containing all the whitespace characters.
// 
// # Discussion
// 
// This set doesn’t contain the newline or carriage return characters.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/whitespaces
func (_NSCharacterSetClass NSCharacterSetClass) WhitespaceCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("whitespaceCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}
// Returns the character set for characters allowed in a fragment URL
// component.
//
// # Discussion
// 
// The fragment component of a URL is the component after a `#` symbol. For
// example, in the URL
// `//www.ExampleXCUIElementTypeCom()/index.Html()#jumpLocation`, the fragment
// is `jumpLocation`.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlFragmentAllowed
func (_NSCharacterSetClass NSCharacterSetClass) URLFragmentAllowedCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("URLFragmentAllowedCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}
// Returns the character set for characters allowed in a host URL
// subcomponent.
//
// # Discussion
// 
// The host component of a URL is usually the component immediately after the
// first two leading slashes. If the URL contains a username and password, the
// host component is the component after the `@` sign. For example, in the URL
// `//password@www.ExampleXCUIElementTypeCom()/index.Html()`, the host
// component is `www.ExampleXCUIElementTypeCom()`.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlHostAllowed
func (_NSCharacterSetClass NSCharacterSetClass) URLHostAllowedCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("URLHostAllowedCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}
// Returns the character set for characters allowed in a password URL
// subcomponent.
//
// # Discussion
// 
// The password component of a URL is the component immediately following the
// colon after the username component of the URL, and ends at the `@` sign.
// For example, in the URL
// `//password@www.ExampleXCUIElementTypeCom()/index.Html()`, the pass
// component is `password`.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlPasswordAllowed
func (_NSCharacterSetClass NSCharacterSetClass) URLPasswordAllowedCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("URLPasswordAllowedCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}
// Returns the character set for characters allowed in a path URL component.
//
// # Discussion
// 
// The path component of a URL is the component immediately following the host
// component (if present). It ends wherever the query or fragment component
// begins. For example, in the URL
// `//www.ExampleXCUIElementTypeCom()/index.Php()?key1=value1`, the path
// component is `/index.Php()`.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlPathAllowed
func (_NSCharacterSetClass NSCharacterSetClass) URLPathAllowedCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("URLPathAllowedCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}
// Returns the character set for characters allowed in a query URL component.
//
// # Discussion
// 
// The query component of a URL is the component immediately following a
// question mark (`?`). For example, in the URL
// `//www.ExampleXCUIElementTypeCom()/index.Php()?key1=value1#jumpLink`, the
// query component is `key1=value1`.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlQueryAllowed
func (_NSCharacterSetClass NSCharacterSetClass) URLQueryAllowedCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("URLQueryAllowedCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}
// Returns the character set for characters allowed in a user URL
// subcomponent.
//
// # Discussion
// 
// The user component of a URL is an optional component that precedes the host
// component, and ends at either a colon (if a password is specified) or an
// `@` sign (if no password is specified). For example, in the URL
// `//password@www.ExampleXCUIElementTypeCom()/index.Html()`, the user
// component is `username`.
//
// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlUserAllowed
func (_NSCharacterSetClass NSCharacterSetClass) URLUserAllowedCharacterSet() NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSCharacterSetClass.class), objc.Sel("URLUserAllowedCharacterSet"))
	return NSCharacterSetFromID(objc.ID(rv))
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSMutableCopying
			

			// Protocol methods for NSSecureCoding
			

