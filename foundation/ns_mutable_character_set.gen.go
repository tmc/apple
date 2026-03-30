// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [NSMutableCharacterSet] class.
var (
	_NSMutableCharacterSetClass     NSMutableCharacterSetClass
	_NSMutableCharacterSetClassOnce sync.Once
)

func getNSMutableCharacterSetClass() NSMutableCharacterSetClass {
	_NSMutableCharacterSetClassOnce.Do(func() {
		_NSMutableCharacterSetClass = NSMutableCharacterSetClass{class: objc.GetClass("NSMutableCharacterSet")}
	})
	return _NSMutableCharacterSetClass
}

// GetNSMutableCharacterSetClass returns the class object for NSMutableCharacterSet.
func GetNSMutableCharacterSetClass() NSMutableCharacterSetClass {
	return getNSMutableCharacterSetClass()
}

type NSMutableCharacterSetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSMutableCharacterSetClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMutableCharacterSetClass) Alloc() NSMutableCharacterSet {
	rv := objc.Send[NSMutableCharacterSet](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object representing a mutable set of Unicode character values for use in
// search operations.
//
// # Overview
//
// In Swift, this object bridges to [CharacterSet]; use
// [NSMutableCharacterSet] when you need reference semantics or other
// Foundation-specific behavior.
//
// The [NSMutableCharacterSet] class declares the programmatic interface to
// objects that manage a modifiable set of Unicode characters. You can add or
// remove characters from a mutable character set as numeric values in
// [NSRange] structures or as character values in strings, combine character
// sets by union or intersection, and invert a character set.
//
// Mutable character sets are less efficient to use than immutable character
// sets. If you don’t need to change a character set after creating it,
// create an immutable copy with `copy` and use that.
//
// [NSMutableCharacterSet] defines no primitive methods. Subclasses must
// implement all methods declared by this class in addition to the primitives
// of [NSCharacterSet]. They must also implement [NSMutableCharacterSet.MutableCopyWithZone].
//
// [NSMutableCharacterSet] is “toll-free bridged” with its Core Foundation
// counterpart, [CFMutableCharacterSet]. See [Toll-Free Bridging] for more
// information.
//
// # Adding and Removing Characters
//
//   - [NSMutableCharacterSet.AddCharactersInRange]: Adds to the receiver the characters whose Unicode values are in a given range.
//   - [NSMutableCharacterSet.RemoveCharactersInRange]: Removes from the receiver the characters whose Unicode values are in a given range.
//   - [NSMutableCharacterSet.AddCharactersInString]: Adds to the receiver the characters in a given string.
//   - [NSMutableCharacterSet.RemoveCharactersInString]: Removes from the receiver the characters in a given string.
//
// # Combining Character Sets
//
//   - [NSMutableCharacterSet.FormIntersectionWithCharacterSet]: Modifies the receiver so it contains only characters that exist in both the receiver and another set.
//   - [NSMutableCharacterSet.FormUnionWithCharacterSet]: Modifies the receiver so it contains all characters that exist in either the receiver or another set.
//
// # Inverting a Character Set
//
//   - [NSMutableCharacterSet.Invert]: Replaces all the characters in the receiver with all the characters it didn’t previously contain.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet
//
// [CFMutableCharacterSet]: https://developer.apple.com/documentation/CoreFoundation/CFMutableCharacterSet
// [CharacterSet]: https://developer.apple.com/documentation/Foundation/CharacterSet
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
type NSMutableCharacterSet struct {
	NSCharacterSet
}

// NSMutableCharacterSetFromID constructs a [NSMutableCharacterSet] from an objc.ID.
//
// An object representing a mutable set of Unicode character values for use in
// search operations.
func NSMutableCharacterSetFromID(id objc.ID) NSMutableCharacterSet {
	return NSMutableCharacterSet{NSCharacterSet: NSCharacterSetFromID(id)}
}

// NOTE: NSMutableCharacterSet adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMutableCharacterSet] class.
//
// # Adding and Removing Characters
//
//   - [INSMutableCharacterSet.AddCharactersInRange]: Adds to the receiver the characters whose Unicode values are in a given range.
//   - [INSMutableCharacterSet.RemoveCharactersInRange]: Removes from the receiver the characters whose Unicode values are in a given range.
//   - [INSMutableCharacterSet.AddCharactersInString]: Adds to the receiver the characters in a given string.
//   - [INSMutableCharacterSet.RemoveCharactersInString]: Removes from the receiver the characters in a given string.
//
// # Combining Character Sets
//
//   - [INSMutableCharacterSet.FormIntersectionWithCharacterSet]: Modifies the receiver so it contains only characters that exist in both the receiver and another set.
//   - [INSMutableCharacterSet.FormUnionWithCharacterSet]: Modifies the receiver so it contains all characters that exist in either the receiver or another set.
//
// # Inverting a Character Set
//
//   - [INSMutableCharacterSet.Invert]: Replaces all the characters in the receiver with all the characters it didn’t previously contain.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet
type INSMutableCharacterSet interface {
	INSCharacterSet

	// Topic: Adding and Removing Characters

	// Adds to the receiver the characters whose Unicode values are in a given range.
	AddCharactersInRange(aRange NSRange)
	// Removes from the receiver the characters whose Unicode values are in a given range.
	RemoveCharactersInRange(aRange NSRange)
	// Adds to the receiver the characters in a given string.
	AddCharactersInString(aString string)
	// Removes from the receiver the characters in a given string.
	RemoveCharactersInString(aString string)

	// Topic: Combining Character Sets

	// Modifies the receiver so it contains only characters that exist in both the receiver and another set.
	FormIntersectionWithCharacterSet(otherSet INSCharacterSet)
	// Modifies the receiver so it contains all characters that exist in either the receiver or another set.
	FormUnionWithCharacterSet(otherSet INSCharacterSet)

	// Topic: Inverting a Character Set

	// Replaces all the characters in the receiver with all the characters it didn’t previously contain.
	Invert()
}

// Init initializes the instance.
func (m NSMutableCharacterSet) Init() NSMutableCharacterSet {
	rv := objc.Send[NSMutableCharacterSet](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMutableCharacterSet) Autorelease() NSMutableCharacterSet {
	rv := objc.Send[NSMutableCharacterSet](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMutableCharacterSet creates a new NSMutableCharacterSet instance.
func NewNSMutableCharacterSet() NSMutableCharacterSet {
	class := getNSMutableCharacterSetClass()
	rv := objc.Send[NSMutableCharacterSet](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a character set containing characters determined by a given bitmap
// representation.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/init(bitmapRepresentation:)
func NewMutableCharacterSetWithBitmapRepresentation(data INSData) NSMutableCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(getNSMutableCharacterSetClass().class), objc.Sel("characterSetWithBitmapRepresentation:"), data)
	return NSMutableCharacterSetFromID(rv)
}

// Returns a character set containing the characters in a given string.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/init(charactersIn:)
func NewMutableCharacterSetWithCharactersInString(aString string) NSMutableCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(getNSMutableCharacterSetClass().class), objc.Sel("characterSetWithCharactersInString:"), objc.String(aString))
	return NSMutableCharacterSetFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSCharacterSet/init(coder:)
func NewMutableCharacterSetWithCoder(coder INSCoder) NSMutableCharacterSet {
	instance := getNSMutableCharacterSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSMutableCharacterSetFromID(rv)
}

// Returns a character set read from the bitmap representation stored in the
// file a given path.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/init(contentsOfFile:)
func NewMutableCharacterSetWithContentsOfFile(fName string) NSMutableCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(getNSMutableCharacterSetClass().class), objc.Sel("characterSetWithContentsOfFile:"), objc.String(fName))
	return NSMutableCharacterSetFromID(rv)
}

// Returns a character set containing characters with Unicode values in a
// given range.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/init(range:)
func NewMutableCharacterSetWithRange(aRange NSRange) NSMutableCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(getNSMutableCharacterSetClass().class), objc.Sel("characterSetWithRange:"), aRange)
	return NSMutableCharacterSetFromID(rv)
}

// Adds to the receiver the characters whose Unicode values are in a given
// range.
//
// aRange: The range of characters to add. `aRange.Location()` is the value of the
// first character to add; `aRange.Location() + aRange.Length() – 1` is the
// value of the last. If `aRange.Length()` is `0`, this method has no effect.
//
// # Discussion
//
// This code excerpt adds to a character set the lowercase English alphabetic
// characters:
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/addCharacters(in:)-4ppyw
func (m NSMutableCharacterSet) AddCharactersInRange(aRange NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("addCharactersInRange:"), aRange)
}

// Removes from the receiver the characters whose Unicode values are in a
// given range.
//
// aRange: The range of characters to remove. `aRange.Location()` is the value of the
// first character to remove; `aRange.Location() + aRange.Length() – 1` is
// the value of the last. If `aRange.Length()` is `0`, this method has no
// effect.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/removeCharacters(in:)-70nqp
func (m NSMutableCharacterSet) RemoveCharactersInRange(aRange NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeCharactersInRange:"), aRange)
}

// Adds to the receiver the characters in a given string.
//
// aString: The characters to add to the receiver.
//
// # Discussion
//
// This method has no effect if `aString` is empty.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/addCharacters(in:)-7q02
func (m NSMutableCharacterSet) AddCharactersInString(aString string) {
	objc.Send[objc.ID](m.ID, objc.Sel("addCharactersInString:"), objc.String(aString))
}

// Removes from the receiver the characters in a given string.
//
// aString: The characters to remove from the receiver.
//
// # Discussion
//
// This method has no effect if `aString` is empty.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/removeCharacters(in:)-762gt
func (m NSMutableCharacterSet) RemoveCharactersInString(aString string) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeCharactersInString:"), objc.String(aString))
}

// Modifies the receiver so it contains only characters that exist in both the
// receiver and another set.
//
// otherSet: The character set with which to perform the intersection.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/formIntersection(with:)
func (m NSMutableCharacterSet) FormIntersectionWithCharacterSet(otherSet INSCharacterSet) {
	objc.Send[objc.ID](m.ID, objc.Sel("formIntersectionWithCharacterSet:"), otherSet)
}

// Modifies the receiver so it contains all characters that exist in either
// the receiver or another set.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/formUnion(with:)
func (m NSMutableCharacterSet) FormUnionWithCharacterSet(otherSet INSCharacterSet) {
	objc.Send[objc.ID](m.ID, objc.Sel("formUnionWithCharacterSet:"), otherSet)
}

// Replaces all the characters in the receiver with all the characters it
// didn’t previously contain.
//
// # Discussion
//
// Inverting a mutable character set, whether by [Invert] or by [InvertedSet],
// is much less efficient than inverting an immutable character set with
// [InvertedSet].
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/invert()
func (m NSMutableCharacterSet) Invert() {
	objc.Send[objc.ID](m.ID, objc.Sel("invert"))
}
