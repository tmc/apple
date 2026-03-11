// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSDictionary] class.
var (
	_NSDictionaryClass     NSDictionaryClass
	_NSDictionaryClassOnce sync.Once
)

func getNSDictionaryClass() NSDictionaryClass {
	_NSDictionaryClassOnce.Do(func() {
		_NSDictionaryClass = NSDictionaryClass{class: objc.GetClass("NSDictionary")}
	})
	return _NSDictionaryClass
}

// GetNSDictionaryClass returns the class object for NSDictionary.
func GetNSDictionaryClass() NSDictionaryClass {
	return getNSDictionaryClass()
}

type NSDictionaryClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDictionaryClass) Alloc() NSDictionary {
	rv := objc.Send[NSDictionary](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A static collection of objects associated with unique keys.
//
// # Overview
// 
// You can use this type in Swift instead of a [Dictionary] in cases that
// require reference semantics.
// 
// The [NSDictionary] class declares the programmatic interface to objects
// that manage immutable associations of keys and values. For example, an
// interactive form could be represented as a dictionary, with the field names
// as keys, corresponding to user-entered values.
// 
// Use this class or its subclass [NSMutableDictionary] when you need a
// convenient and efficient way to retrieve data associated with an arbitrary
// key. [NSDictionary] creates static dictionaries, and [NSMutableDictionary]
// creates dynamic dictionaries. (For convenience, the term refers to any
// instance of one of these classes without specifying its exact class
// membership.)
// 
// A key-value pair within a dictionary is called an entry. Each entry
// consists of one object that represents the key and a second object that is
// that key’s value. Within a dictionary, the keys are unique. That is, no
// two keys in a single dictionary are equal (as determined by [isEqual(_:)]).
// In general, a key can be any object (provided that it conforms to the
// [NSCopying] protocol—see below), but note that when using key-value
// coding the key must be a string (see [Accessing Object Properties]).
// Neither a key nor a value can be `nil`; if you need to represent a null
// value in a dictionary, you should use [NSNull].
// 
// [NSDictionary] is “toll-free bridged” with its Core Foundation
// counterpart, [CFDictionary]. See [Toll-Free Bridging] for more information
// on toll-free bridging.
// 
// # Creating NSDictionary Objects Using Dictionary Literals
// 
// In addition to the provided initializers, such as [NSDictionary.InitWithObjectsForKeys],
// you can create an [NSDictionary] object using a .
// 
// In Objective-C, the compiler generates code that makes an underlying call
// to the [NSDictionary.DictionaryWithObjectsForKeysCount] method.
// 
// Unlike [NSDictionary.DictionaryWithObjectsAndKeys] and other initializers, dictionary
// literals specify entries in key-value order. You should not terminate the
// list of objects with `nil` when using this literal syntax, and in fact
// `nil` is an invalid value. For more information about object literals in
// Objective-C, see [Working with Objects] in [Programming with Objective-C].
// 
// In Swift, the [NSDictionary] class conforms to the
// [DictionaryLiteralConvertible] protocol, which allows it to be initialized
// with dictionary literals. For more information about object literals in
// Swift, see [Literal Expression] in [The Swift Programming Language (Swift
// 4.1)].
// 
// # Accessing Values Using Subscripting
// 
// In addition to the provided instance methods, such as [NSDictionary.ObjectForKey], you
// can access [NSDictionary] values by their keys using .
// 
// # Enumerating Entries Using for-in Loops
// 
// In addition to the provided instance methods, such as
// [NSDictionary.EnumerateKeysAndObjectsUsingBlock], you can enumerate [NSDictionary]
// entries using .
// 
// In Objective-C, [NSDictionary] conforms to the [NSFastEnumeration]
// protocol.
// 
// In Swift, [NSDictionary] conforms to the [SequenceType] protocol.
// 
// # Subclassing Notes
// 
// You generally shouldn’t need to subclass [NSDictionary]. Custom behavior
// can usually be achieved through composition rather than subclassing.
// 
// # Methods to Override
// 
// If you do need to subclass [NSDictionary], take into account that it is a
// [Class cluster]. Any subclass must override the following primitive
// methods:
// 
// - [NSDictionary.InitWithObjectsForKeysCount] - [NSDictionary.Count] - [NSDictionary.ObjectForKey] -
// [NSDictionary.KeyEnumerator]
// 
// The other methods of [NSDictionary] operate by invoking one or more of
// these primitives. The non-primitive methods provide convenient ways of
// accessing multiple entries at once.
// 
// # Alternatives to Subclassing
// 
// Before making a custom class of [NSDictionary], investigate [NSMapTable]
// and the corresponding Core Foundation type, [CFDictionary]. Because
// [NSDictionary] and [CFDictionary] are “toll-free bridged,” you can
// substitute a [CFDictionary] object for a [NSDictionary] object in your code
// (with appropriate casting). Although they are corresponding types,
// [CFDictionary] and [NSDictionary] do not have identical interfaces or
// implementations, and you can sometimes do things with [CFDictionary] that
// you cannot easily do with [NSDictionary].
// 
// If the behavior you want to add supplements that of the existing class, you
// could write a category on [NSDictionary]. Keep in mind, however, that this
// category will be in effect for all instances of [NSDictionary] that you
// use, and this might have unintended consequences. Alternatively, you could
// use composition to achieve the desired behavior.
//
// [Accessing Object Properties]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueCoding/BasicPrinciples.html#//apple_ref/doc/uid/20002170
// [CFDictionary]: https://developer.apple.com/documentation/CoreFoundation/CFDictionary
// [Class cluster]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/ClassCluster.html#//apple_ref/doc/uid/TP40008195-CH7
// [Dictionary]: https://developer.apple.com/documentation/Swift/Dictionary
// [Literal Expression]: https://developer.apple.com/library/archive/documentation/Swift/Conceptual/Swift_Programming_Language/Expressions.html#//apple_ref/doc/uid/TP40014097-CH32-ID390
// [Programming with Objective-C]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ProgrammingWithObjectiveC/Introduction/Introduction.html#//apple_ref/doc/uid/TP40011210
// [The Swift Programming Language (Swift 4.1)]: https://developer.apple.com/library/archive/documentation/Swift/Conceptual/Swift_Programming_Language/index.html#//apple_ref/doc/uid/TP40014097
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
// [Working with Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ProgrammingWithObjectiveC/WorkingwithObjects/WorkingwithObjects.html#//apple_ref/doc/uid/TP40011210-CH4
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
//
// # Creating a Dictionary from Objects and Keys
//
//   - [NSDictionary.InitWithObjectsForKeys]: Initializes a newly allocated dictionary with key-value pairs constructed from the provided arrays of keys and objects.
//   - [NSDictionary.InitWithObjectsForKeysCount]: Initializes a newly allocated dictionary with the specified number of key-value pairs constructed from the provided C arrays of keys and objects.
//
// # Creating a Dictionary from Another Dictionary
//
//   - [NSDictionary.InitWithDictionary]: Initializes a newly allocated dictionary by placing in it the keys and values contained in another given dictionary.
//   - [NSDictionary.InitWithDictionaryCopyItems]: Initializes a newly allocated dictionary using the objects contained in another given dictionary.
//
// # Creating a Dictionary from an NSCoder
//
//   - [NSDictionary.InitWithCoder]: Creates a dictionary initialized from data in the provided unarchiver.
//
// # Counting Entries
//
//   - [NSDictionary.Count]: The number of entries in the dictionary.
//
// # Comparing Dictionaries
//
//   - [NSDictionary.IsEqualToDictionary]: Returns a Boolean value that indicates whether the contents of the receiving dictionary are equal to the contents of another given dictionary.
//
// # Accessing Keys and Values
//
//   - [NSDictionary.AllKeys]: A new array containing the dictionary’s keys, or an empty array if the dictionary has no entries.
//   - [NSDictionary.AllKeysForObject]: Returns a new array containing the keys corresponding to all occurrences of a given object in the dictionary.
//   - [NSDictionary.AllValues]: A new array containing the dictionary’s values, or an empty array if the dictionary has no entries.
//   - [NSDictionary.ObjectsForKeysNotFoundMarker]: Returns as a static array the set of objects from the dictionary that corresponds to the specified keys.
//   - [NSDictionary.ObjectForKey]: Returns the value associated with a given key.
//   - [NSDictionary.ObjectForKeyedSubscript]: Returns the value associated with a given key.
//
// # Enumerating Dictionaries
//
//   - [NSDictionary.KeyEnumerator]: Provides an enumerator to access the keys in the dictionary.
//   - [NSDictionary.ObjectEnumerator]: Returns an enumerator object that lets you access each value in the dictionary.
//   - [NSDictionary.EnumerateKeysAndObjectsUsingBlock]: Applies a given block object to the entries of the dictionary.
//   - [NSDictionary.EnumerateKeysAndObjectsWithOptionsUsingBlock]: Applies a given block object to the entries of the dictionary, with options specifying how the enumeration is performed.
//
// # Sorting Dictionaries
//
//   - [NSDictionary.KeysSortedByValueUsingSelector]: Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values.
//   - [NSDictionary.KeysSortedByValueUsingComparator]: Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values using a given comparator block.
//   - [NSDictionary.KeysSortedByValueWithOptionsUsingComparator]: Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values using a given comparator block and a specified set of options.
//
// # Filtering Dictionaries
//
//   - [NSDictionary.KeysOfEntriesPassingTest]: Returns the set of keys whose corresponding value satisfies a constraint described by a block object.
//   - [NSDictionary.KeysOfEntriesWithOptionsPassingTest]: Returns the set of keys whose corresponding value satisfies a constraint described by a block object.
//
// # Storing Dictionaries
//
//   - [NSDictionary.WriteToURLError]: Writes a property list representation of the contents of the dictionary to a given URL.
//
// # Accessing File Attributes
//
//   - [NSDictionary.FileSize]: Returns the file’s size, in bytes.
//   - [NSDictionary.FileType]: Returns the file type.
//   - [NSDictionary.FileCreationDate]: Returns the file’s creation date.
//   - [NSDictionary.FileModificationDate]: Returns file’s modification date.
//   - [NSDictionary.FilePosixPermissions]: Returns the file’s POSIX permissions.
//   - [NSDictionary.FileOwnerAccountID]: Returns the file’s owner account ID.
//   - [NSDictionary.FileOwnerAccountName]: Returns the file’s owner account name.
//   - [NSDictionary.FileGroupOwnerAccountID]: Returns file’s group owner account ID.
//   - [NSDictionary.FileGroupOwnerAccountName]: Returns the file’s group owner account name.
//   - [NSDictionary.FileExtensionHidden]: Returns a Boolean value indicating whether the file hides its extension.
//   - [NSDictionary.FileIsImmutable]: Returns a Boolean value indicating whether the file is immutable.
//   - [NSDictionary.FileIsAppendOnly]: Returns a Boolean value indicating whether the file is append only.
//   - [NSDictionary.FileSystemFileNumber]: Returns the filesystem file number.
//   - [NSDictionary.FileSystemNumber]: Returns the filesystem number.
//   - [NSDictionary.FileHFSTypeCode]: Returns file’s HFS type code.
//   - [NSDictionary.FileHFSCreatorCode]: Returns the file’s HFS creator code.
//
// # Describing a Dictionary
//
//   - [NSDictionary.Description]: A string that represents the contents of the dictionary, formatted as a property list.
//   - [NSDictionary.DescriptionInStringsFileFormat]: A string that represents the contents of the dictionary, formatted in `XCUIElementTypeStrings` file format.
//   - [NSDictionary.DescriptionWithLocale]: Returns a string object that represents the contents of the dictionary, formatted as a property list.
//   - [NSDictionary.DescriptionWithLocaleIndent]: Returns a string object that represents the contents of the dictionary, formatted as a property list.
//
// # Initializers
//
//   - [NSDictionary.InitWithContentsOfURLError]
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary
type NSDictionary struct {
	objectivec.Object
}

// NSDictionaryFromID constructs a [NSDictionary] from an objc.ID.
//
// A static collection of objects associated with unique keys.
func NSDictionaryFromID(id objc.ID) NSDictionary {
	return NSDictionary{objectivec.Object{id}}
}
// NOTE: NSDictionary adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSDictionary] class.
//
// # Creating a Dictionary from Objects and Keys
//
//   - [INSDictionary.InitWithObjectsForKeys]: Initializes a newly allocated dictionary with key-value pairs constructed from the provided arrays of keys and objects.
//   - [INSDictionary.InitWithObjectsForKeysCount]: Initializes a newly allocated dictionary with the specified number of key-value pairs constructed from the provided C arrays of keys and objects.
//
// # Creating a Dictionary from Another Dictionary
//
//   - [INSDictionary.InitWithDictionary]: Initializes a newly allocated dictionary by placing in it the keys and values contained in another given dictionary.
//   - [INSDictionary.InitWithDictionaryCopyItems]: Initializes a newly allocated dictionary using the objects contained in another given dictionary.
//
// # Creating a Dictionary from an NSCoder
//
//   - [INSDictionary.InitWithCoder]: Creates a dictionary initialized from data in the provided unarchiver.
//
// # Counting Entries
//
//   - [INSDictionary.Count]: The number of entries in the dictionary.
//
// # Comparing Dictionaries
//
//   - [INSDictionary.IsEqualToDictionary]: Returns a Boolean value that indicates whether the contents of the receiving dictionary are equal to the contents of another given dictionary.
//
// # Accessing Keys and Values
//
//   - [INSDictionary.AllKeys]: A new array containing the dictionary’s keys, or an empty array if the dictionary has no entries.
//   - [INSDictionary.AllKeysForObject]: Returns a new array containing the keys corresponding to all occurrences of a given object in the dictionary.
//   - [INSDictionary.AllValues]: A new array containing the dictionary’s values, or an empty array if the dictionary has no entries.
//   - [INSDictionary.ObjectsForKeysNotFoundMarker]: Returns as a static array the set of objects from the dictionary that corresponds to the specified keys.
//   - [INSDictionary.ObjectForKey]: Returns the value associated with a given key.
//   - [INSDictionary.ObjectForKeyedSubscript]: Returns the value associated with a given key.
//
// # Enumerating Dictionaries
//
//   - [INSDictionary.KeyEnumerator]: Provides an enumerator to access the keys in the dictionary.
//   - [INSDictionary.ObjectEnumerator]: Returns an enumerator object that lets you access each value in the dictionary.
//   - [INSDictionary.EnumerateKeysAndObjectsUsingBlock]: Applies a given block object to the entries of the dictionary.
//   - [INSDictionary.EnumerateKeysAndObjectsWithOptionsUsingBlock]: Applies a given block object to the entries of the dictionary, with options specifying how the enumeration is performed.
//
// # Sorting Dictionaries
//
//   - [INSDictionary.KeysSortedByValueUsingSelector]: Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values.
//   - [INSDictionary.KeysSortedByValueUsingComparator]: Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values using a given comparator block.
//   - [INSDictionary.KeysSortedByValueWithOptionsUsingComparator]: Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values using a given comparator block and a specified set of options.
//
// # Filtering Dictionaries
//
//   - [INSDictionary.KeysOfEntriesPassingTest]: Returns the set of keys whose corresponding value satisfies a constraint described by a block object.
//   - [INSDictionary.KeysOfEntriesWithOptionsPassingTest]: Returns the set of keys whose corresponding value satisfies a constraint described by a block object.
//
// # Storing Dictionaries
//
//   - [INSDictionary.WriteToURLError]: Writes a property list representation of the contents of the dictionary to a given URL.
//
// # Accessing File Attributes
//
//   - [INSDictionary.FileSize]: Returns the file’s size, in bytes.
//   - [INSDictionary.FileType]: Returns the file type.
//   - [INSDictionary.FileCreationDate]: Returns the file’s creation date.
//   - [INSDictionary.FileModificationDate]: Returns file’s modification date.
//   - [INSDictionary.FilePosixPermissions]: Returns the file’s POSIX permissions.
//   - [INSDictionary.FileOwnerAccountID]: Returns the file’s owner account ID.
//   - [INSDictionary.FileOwnerAccountName]: Returns the file’s owner account name.
//   - [INSDictionary.FileGroupOwnerAccountID]: Returns file’s group owner account ID.
//   - [INSDictionary.FileGroupOwnerAccountName]: Returns the file’s group owner account name.
//   - [INSDictionary.FileExtensionHidden]: Returns a Boolean value indicating whether the file hides its extension.
//   - [INSDictionary.FileIsImmutable]: Returns a Boolean value indicating whether the file is immutable.
//   - [INSDictionary.FileIsAppendOnly]: Returns a Boolean value indicating whether the file is append only.
//   - [INSDictionary.FileSystemFileNumber]: Returns the filesystem file number.
//   - [INSDictionary.FileSystemNumber]: Returns the filesystem number.
//   - [INSDictionary.FileHFSTypeCode]: Returns file’s HFS type code.
//   - [INSDictionary.FileHFSCreatorCode]: Returns the file’s HFS creator code.
//
// # Describing a Dictionary
//
//   - [INSDictionary.Description]: A string that represents the contents of the dictionary, formatted as a property list.
//   - [INSDictionary.DescriptionInStringsFileFormat]: A string that represents the contents of the dictionary, formatted in `XCUIElementTypeStrings` file format.
//   - [INSDictionary.DescriptionWithLocale]: Returns a string object that represents the contents of the dictionary, formatted as a property list.
//   - [INSDictionary.DescriptionWithLocaleIndent]: Returns a string object that represents the contents of the dictionary, formatted as a property list.
//
// # Initializers
//
//   - [INSDictionary.InitWithContentsOfURLError]
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary
type INSDictionary interface {
	objectivec.IObject
	NSCopying
	NSMutableCopying

	// Topic: Creating a Dictionary from Objects and Keys

	// Initializes a newly allocated dictionary with key-value pairs constructed from the provided arrays of keys and objects.
	InitWithObjectsForKeys(objects []objectivec.IObject, keys []objectivec.IObject) NSDictionary
	// Initializes a newly allocated dictionary with the specified number of key-value pairs constructed from the provided C arrays of keys and objects.
	InitWithObjectsForKeysCount(objects []objectivec.IObject, keys []NSCopying, cnt uint) NSDictionary

	// Topic: Creating a Dictionary from Another Dictionary

	// Initializes a newly allocated dictionary by placing in it the keys and values contained in another given dictionary.
	InitWithDictionary(otherDictionary INSDictionary) NSDictionary
	// Initializes a newly allocated dictionary using the objects contained in another given dictionary.
	InitWithDictionaryCopyItems(otherDictionary INSDictionary, flag bool) NSDictionary

	// Topic: Creating a Dictionary from an NSCoder

	// Creates a dictionary initialized from data in the provided unarchiver.
	InitWithCoder(coder INSCoder) NSDictionary

	// Topic: Counting Entries

	// The number of entries in the dictionary.
	Count() uint

	// Topic: Comparing Dictionaries

	// Returns a Boolean value that indicates whether the contents of the receiving dictionary are equal to the contents of another given dictionary.
	IsEqualToDictionary(otherDictionary INSDictionary) bool

	// Topic: Accessing Keys and Values

	// A new array containing the dictionary’s keys, or an empty array if the dictionary has no entries.
	AllKeys() []objectivec.IObject
	// Returns a new array containing the keys corresponding to all occurrences of a given object in the dictionary.
	AllKeysForObject(anObject objectivec.IObject) []objectivec.IObject
	// A new array containing the dictionary’s values, or an empty array if the dictionary has no entries.
	AllValues() []objectivec.IObject
	// Returns as a static array the set of objects from the dictionary that corresponds to the specified keys.
	ObjectsForKeysNotFoundMarker(keys []objectivec.IObject, marker objectivec.IObject) []objectivec.IObject
	// Returns the value associated with a given key.
	ObjectForKey(aKey objectivec.IObject) objectivec.IObject
	// Returns the value associated with a given key.
	ObjectForKeyedSubscript(key objectivec.IObject) objectivec.IObject

	// Topic: Enumerating Dictionaries

	// Provides an enumerator to access the keys in the dictionary.
	KeyEnumerator() INSEnumerator
	// Returns an enumerator object that lets you access each value in the dictionary.
	ObjectEnumerator() INSEnumerator
	// Applies a given block object to the entries of the dictionary.
	EnumerateKeysAndObjectsUsingBlock(block bool)
	// Applies a given block object to the entries of the dictionary, with options specifying how the enumeration is performed.
	EnumerateKeysAndObjectsWithOptionsUsingBlock(opts NSEnumerationOptions, block bool)

	// Topic: Sorting Dictionaries

	// Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values.
	KeysSortedByValueUsingSelector(comparator objc.SEL) []objectivec.IObject
	// Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values using a given comparator block.
	KeysSortedByValueUsingComparator(cmptr NSComparator) []objectivec.IObject
	// Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values using a given comparator block and a specified set of options.
	KeysSortedByValueWithOptionsUsingComparator(opts NSSortOptions, cmptr NSComparator) []objectivec.IObject

	// Topic: Filtering Dictionaries

	// Returns the set of keys whose corresponding value satisfies a constraint described by a block object.
	KeysOfEntriesPassingTest(predicate bool) INSSet
	// Returns the set of keys whose corresponding value satisfies a constraint described by a block object.
	KeysOfEntriesWithOptionsPassingTest(opts NSEnumerationOptions, predicate bool) INSSet

	// Topic: Storing Dictionaries

	// Writes a property list representation of the contents of the dictionary to a given URL.
	WriteToURLError(url INSURL) (bool, error)

	// Topic: Accessing File Attributes

	// Returns the file’s size, in bytes.
	FileSize() uint64
	// Returns the file type.
	FileType() string
	// Returns the file’s creation date.
	FileCreationDate() INSDate
	// Returns file’s modification date.
	FileModificationDate() INSDate
	// Returns the file’s POSIX permissions.
	FilePosixPermissions() uint
	// Returns the file’s owner account ID.
	FileOwnerAccountID() INSNumber
	// Returns the file’s owner account name.
	FileOwnerAccountName() string
	// Returns file’s group owner account ID.
	FileGroupOwnerAccountID() INSNumber
	// Returns the file’s group owner account name.
	FileGroupOwnerAccountName() string
	// Returns a Boolean value indicating whether the file hides its extension.
	FileExtensionHidden() bool
	// Returns a Boolean value indicating whether the file is immutable.
	FileIsImmutable() bool
	// Returns a Boolean value indicating whether the file is append only.
	FileIsAppendOnly() bool
	// Returns the filesystem file number.
	FileSystemFileNumber() uint
	// Returns the filesystem number.
	FileSystemNumber() int
	// Returns file’s HFS type code.
	FileHFSTypeCode() uint32
	// Returns the file’s HFS creator code.
	FileHFSCreatorCode() uint32

	// Topic: Describing a Dictionary

	// A string that represents the contents of the dictionary, formatted as a property list.
	Description() string
	// A string that represents the contents of the dictionary, formatted in `XCUIElementTypeStrings` file format.
	DescriptionInStringsFileFormat() string
	// Returns a string object that represents the contents of the dictionary, formatted as a property list.
	DescriptionWithLocale(locale objectivec.IObject) string
	// Returns a string object that represents the contents of the dictionary, formatted as a property list.
	DescriptionWithLocaleIndent(locale objectivec.IObject, level uint) string

	// Topic: Initializers

	InitWithContentsOfURLError(url INSURL) (NSDictionary, error)

	// Encodes the receiver using a given archiver.
	EncodeWithCoder(coder INSCoder)
	// Returns by reference C arrays of the keys and values in the dictionary.
	GetObjectsAndKeysCount(objects []objectivec.IObject, keys []objectivec.IObject, count uint)
	// Initializes a newly allocated dictionary with entries constructed from the specified set of values and keys.
	InitWithObjectsAndKeys(firstObject objectivec.IObject) NSDictionary
}





// Init initializes the instance.
func (d NSDictionary) Init() NSDictionary {
	rv := objc.Send[NSDictionary](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDictionary) Autorelease() NSDictionary {
	rv := objc.Send[NSDictionary](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDictionary creates a new NSDictionary instance.
func NewNSDictionary() NSDictionary {
	class := getNSDictionaryClass()
	rv := objc.Send[NSDictionary](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a dictionary initialized from data in the provided unarchiver.
//
// coder: An unarchiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/init(coder:)
func NewDictionaryWithCoder(coder INSCoder) NSDictionary {
	instance := getNSDictionaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSDictionaryFromID(rv)
}


// Initializes a newly allocated dictionary by placing in it the keys and
// values contained in another given dictionary.
//
// otherDictionary: A dictionary containing the keys and values with which to initialize the
// new dictionary.
//
// # Return Value
// 
// An initialized dictionary—which might be different than the original
// receiver—containing the keys and values found in `otherDictionary`.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/init(dictionary:)-9fw1u
func NewDictionaryWithDictionary(otherDictionary INSDictionary) NSDictionary {
	instance := getNSDictionaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDictionary:"), otherDictionary)
	return NSDictionaryFromID(rv)
}


// Initializes a newly allocated dictionary using the objects contained in
// another given dictionary.
//
// otherDictionary: A dictionary containing the keys and values with which to initialize the
// new dictionary.
//
// flag: If [true], each object in `otherDictionary` receives a [copyWithZone:]
// message to create a copy of the object—objects must conform to the
// [NSCopying] protocol. In a managed memory environment, this is instead of
// the `retain` message the object would otherwise receive. The object copy is
// then added to the returned dictionary.
// 
// If [false], then in a managed memory environment each object in
// `otherDictionary` simply receives a `retain` message when it is added to
// the returned dictionary.
// //
// [copyWithZone:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/copyWithZone:
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An initialized object—which might be different than the original
// receiver—containing the keys and values found in `otherDictionary`.
//
// # Discussion
// 
// After an immutable dictionary has been initialized in this way, it cannot
// be modified.
// 
// The [CopyWithZone] method performs a shallow copy. If you have a collection
// of arbitrary depth, passing [true] for the `flag` parameter will perform an
// immutable copy of the first level below the surface. If you pass [false]
// the mutability of the first level is unaffected. In either case, the
// mutability of all deeper levels is unaffected.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/init(dictionary:copyItems:)
func NewDictionaryWithDictionaryCopyItems(otherDictionary INSDictionary, flag bool) NSDictionary {
	instance := getNSDictionaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDictionary:copyItems:"), otherDictionary, flag)
	return NSDictionaryFromID(rv)
}


// Creates a dictionary containing a given key and value.
//
// object: The value corresponding to `aKey`.
// 
// If this value is `nil`, an [invalidArgumentException] is raised.
// //
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
//
// key: The key for `anObject`.
// 
// If this value is `nil`, an [invalidArgumentException] is raised.
// //
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
//
// # Return Value
// 
// A new dictionary containing a single object, `object`, for a single key,
// `aKey`.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/init(object:forKey:)
func NewDictionaryWithObjectForKey(object objectivec.IObject, key NSCopying) NSDictionary {
	rv := objc.Send[objc.ID](objc.ID(getNSDictionaryClass().class), objc.Sel("dictionaryWithObject:forKey:"), object, key)
	return NSDictionaryFromID(rv)
}


// Initializes a newly allocated dictionary with entries constructed from the
// specified set of values and keys.
//
// firstObject: The first value to add to the new dictionary.
//
// # Discussion
// 
// After the `firstObject` value, pass the key for `firstObject`, then a
// null-terminated list of alternating values and keys. If any key is `nil`,
// an [invalidArgumentException] is raised.
// 
// This method is similar to [InitWithObjectsForKeys], differing only in the
// way in which the key-value pairs are specified.
// 
// For example:
//
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/initWithObjectsAndKeys:
func NewDictionaryWithObjectsAndKeys(firstObject objectivec.IObject) NSDictionary {
	instance := getNSDictionaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObjectsAndKeys:"), firstObject)
	return NSDictionaryFromID(rv)
}


// Initializes a newly allocated dictionary with key-value pairs constructed
// from the provided arrays of keys and objects.
//
// objects: An array containing the values for the new dictionary.
//
// keys: An array containing the keys for the new dictionary. Each key is copied
// (using [CopyWithZone]; keys must conform to the [NSCopying] protocol), and
// the copy is added to the new dictionary.
//
// # Discussion
// 
// This method steps through the `objects` and `keys` arrays, creating entries
// in the new dictionary as it goes. An [NSInvalidArgumentException] is raised
// if the objects and keys arrays do not have the same number of elements.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/init(objects:forKeys:)
func NewDictionaryWithObjectsForKeys(objects []objectivec.IObject, keys []objectivec.IObject) NSDictionary {
	instance := getNSDictionaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObjects:forKeys:"), objectivec.IObjectSliceToNSArray(objects), objectivec.IObjectSliceToNSArray(keys))
	return NSDictionaryFromID(rv)
}







// Initializes a newly allocated dictionary with key-value pairs constructed
// from the provided arrays of keys and objects.
//
// objects: An array containing the values for the new dictionary.
//
// keys: An array containing the keys for the new dictionary. Each key is copied
// (using [CopyWithZone]; keys must conform to the [NSCopying] protocol), and
// the copy is added to the new dictionary.
//
// # Discussion
// 
// This method steps through the `objects` and `keys` arrays, creating entries
// in the new dictionary as it goes. An [NSInvalidArgumentException] is raised
// if the objects and keys arrays do not have the same number of elements.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/init(objects:forKeys:)
func (d NSDictionary) InitWithObjectsForKeys(objects []objectivec.IObject, keys []objectivec.IObject) NSDictionary {
	rv := objc.Send[NSDictionary](d.ID, objc.Sel("initWithObjects:forKeys:"), objectivec.IObjectSliceToNSArray(objects), objectivec.IObjectSliceToNSArray(keys))
	return rv
}

// Initializes a newly allocated dictionary with the specified number of
// key-value pairs constructed from the provided C arrays of keys and objects.
//
// objects: A C array of values for the new dictionary.
//
// keys: A C array of keys for the new dictionary. Each key is copied (using
// [CopyWithZone]; keys must conform to the [NSCopying] protocol), and the
// copy is added to the new dictionary.
//
// cnt: The number of elements to use from the `keys` and `objects` arrays. `count`
// must not exceed the number of elements in `objects` or `keys`.
//
// # Discussion
// 
// This method steps through the `objects` and `keys` arrays, creating entries
// in the new dictionary as it goes. An [NSInvalidArgumentException] is raised
// if a key or value object is `nil`.
// 
// This method is a designated initializer of [NSDictionary].
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/init(objects:forKeys:count:)
func (d NSDictionary) InitWithObjectsForKeysCount(objects []objectivec.IObject, keys []NSCopying, cnt uint) NSDictionary {
	rv := objc.Send[NSDictionary](d.ID, objc.Sel("initWithObjects:forKeys:count:"), objc.CArray(objects), objc.CArray(keys), cnt)
	return rv
}

// Initializes a newly allocated dictionary by placing in it the keys and
// values contained in another given dictionary.
//
// otherDictionary: A dictionary containing the keys and values with which to initialize the
// new dictionary.
//
// # Return Value
// 
// An initialized dictionary—which might be different than the original
// receiver—containing the keys and values found in `otherDictionary`.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/init(dictionary:)-9fw1u
func (d NSDictionary) InitWithDictionary(otherDictionary INSDictionary) NSDictionary {
	rv := objc.Send[NSDictionary](d.ID, objc.Sel("initWithDictionary:"), otherDictionary)
	return rv
}

// Initializes a newly allocated dictionary using the objects contained in
// another given dictionary.
//
// otherDictionary: A dictionary containing the keys and values with which to initialize the
// new dictionary.
//
// flag: If [true], each object in `otherDictionary` receives a [copyWithZone:]
// message to create a copy of the object—objects must conform to the
// [NSCopying] protocol. In a managed memory environment, this is instead of
// the `retain` message the object would otherwise receive. The object copy is
// then added to the returned dictionary.
// 
// If [false], then in a managed memory environment each object in
// `otherDictionary` simply receives a `retain` message when it is added to
// the returned dictionary.
// //
// [copyWithZone:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/copyWithZone:
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An initialized object—which might be different than the original
// receiver—containing the keys and values found in `otherDictionary`.
//
// # Discussion
// 
// After an immutable dictionary has been initialized in this way, it cannot
// be modified.
// 
// The [CopyWithZone] method performs a shallow copy. If you have a collection
// of arbitrary depth, passing [true] for the `flag` parameter will perform an
// immutable copy of the first level below the surface. If you pass [false]
// the mutability of the first level is unaffected. In either case, the
// mutability of all deeper levels is unaffected.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/init(dictionary:copyItems:)
func (d NSDictionary) InitWithDictionaryCopyItems(otherDictionary INSDictionary, flag bool) NSDictionary {
	rv := objc.Send[NSDictionary](d.ID, objc.Sel("initWithDictionary:copyItems:"), otherDictionary, flag)
	return rv
}

// Creates a dictionary initialized from data in the provided unarchiver.
//
// coder: An unarchiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/init(coder:)
func (d NSDictionary) InitWithCoder(coder INSCoder) NSDictionary {
	rv := objc.Send[NSDictionary](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Returns a Boolean value that indicates whether the contents of the
// receiving dictionary are equal to the contents of another given dictionary.
//
// otherDictionary: The dictionary with which to compare the receiving dictionary.
//
// # Return Value
// 
// [true] if the contents of `otherDictionary` are equal to the contents of
// the receiving dictionary, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Two dictionaries have equal contents if they each hold the same number of
// entries and, for a given key, the corresponding value objects in each
// dictionary satisfy the [isEqual(_:)] test.
//
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/isEqual(to:)
func (d NSDictionary) IsEqualToDictionary(otherDictionary INSDictionary) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isEqualToDictionary:"), otherDictionary)
	return rv
}

// Returns a new array containing the keys corresponding to all occurrences of
// a given object in the dictionary.
//
// anObject: The value to look for in the dictionary.
//
// # Return Value
// 
// A new array containing the keys corresponding to all occurrences of
// `anObject` in the dictionary. If no object matching `anObject` is found,
// returns an empty array.
//
// # Discussion
// 
// Each object in the dictionary is sent an [isEqual(_:)] message to determine
// if it’s equal to `anObject`.
//
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/allKeys(for:)
func (d NSDictionary) AllKeysForObject(anObject objectivec.IObject) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("allKeysForObject:"), anObject)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns as a static array the set of objects from the dictionary that
// corresponds to the specified keys.
//
// keys: An [NSArray] containing the keys for which to return corresponding values.
//
// marker: The marker object to place in the corresponding element of the returned
// array if an object isn’t found in the dictionary to correspond to a given
// key.
//
// # Discussion
// 
// The objects in the returned array and the `keys` array have a one-for-one
// correspondence, so that the nthe object in the returned array corresponds
// to the nthe key in `keys`.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/objects(forKeys:notFoundMarker:)
func (d NSDictionary) ObjectsForKeysNotFoundMarker(keys []objectivec.IObject, marker objectivec.IObject) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("objectsForKeys:notFoundMarker:"), objectivec.IObjectSliceToNSArray(keys), marker)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns the value associated with a given key.
//
// aKey: The key for which to return the corresponding value.
//
// # Return Value
// 
// The value associated with `aKey`, or `nil` if no value is associated with
// `aKey`.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/object(forKey:)
func (d NSDictionary) ObjectForKey(aKey objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("objectForKey:"), aKey)
	return objectivec.Object{ID: rv}
}

// Returns the value associated with a given key.
//
// key: The key for which to return the corresponding value.
//
// # Return Value
// 
// The value associated with `key`, or `nil` if no value is associated with
// `aKey`.
//
// # Discussion
// 
// This method has the same behavior as the [ObjectForKey] method.
// 
// You shouldn’t need to call this method directly. Instead, this method is
// called when accessing an object by key using subscripting.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/subscript(_:)-52n56
func (d NSDictionary) ObjectForKeyedSubscript(key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("objectForKeyedSubscript:"), key)
	return objectivec.Object{ID: rv}
}

// Provides an enumerator to access the keys in the dictionary.
//
// # Return Value
// 
// An enumerator object that lets you access each key in the dictionary.
//
// # Discussion
// 
// Here’s how you might use this method.
// 
// If you use this method with instances of mutable subclasses of
// [NSDictionary], your code should not modify the entries during enumeration.
// If you intend to modify the entries, use the [AllKeys] property to create a
// snapshot of the dictionary’s keys. Then use this snapshot to traverse the
// entries, modifying them along the way.
// 
// If you want to enumerate the dictionary’s values rather than its keys,
// use the [ObjectEnumerator] method.
// 
// # Special Considerations
// 
// It is more efficient to use the fast enumeration protocol (see
// [NSFastEnumeration]) than this method. Fast enumeration is available in
// macOS 10.5 and later and iOS 2.0 and later.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/keyEnumerator()
func (d NSDictionary) KeyEnumerator() INSEnumerator {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("keyEnumerator"))
	return NSEnumeratorFromID(rv)
}

// Returns an enumerator object that lets you access each value in the
// dictionary.
//
// # Return Value
// 
// An enumerator object that lets you access each value in the dictionary.
//
// # Discussion
// 
// The following code fragment illustrates how you might use the method.
// 
// If you use this method with instances of mutable subclasses of
// [NSDictionary], your code should not modify the entries during enumeration.
// If you intend to modify the entries, use the [AllValues] method to create a
// “snapshot” of the dictionary’s values. Work from this snapshot to
// modify the values.
// 
// # Special Considerations
// 
// It is more efficient to use the fast enumeration protocol (see
// [NSFastEnumeration]). Fast enumeration is available in macOS 10.5 and later
// and iOS 2.0 and later.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/objectEnumerator()
func (d NSDictionary) ObjectEnumerator() INSEnumerator {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("objectEnumerator"))
	return NSEnumeratorFromID(rv)
}

// Applies a given block object to the entries of the dictionary.
//
// block: A block object to operate on entries in the dictionary.
//
// # Discussion
// 
// If the block sets `*stop` to [true], the enumeration stops.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/enumerateKeysAndObjects(_:)
func (d NSDictionary) EnumerateKeysAndObjectsUsingBlock(block bool) {
	objc.Send[objc.ID](d.ID, objc.Sel("enumerateKeysAndObjectsUsingBlock:"), block)
}

// Applies a given block object to the entries of the dictionary, with options
// specifying how the enumeration is performed.
//
// opts: Enumeration options.
//
// block: A block object to operate on entries in the dictionary.
//
// # Discussion
// 
// If the block sets `*stop` to [true], the enumeration stops.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/enumerateKeysAndObjects(options:using:)
func (d NSDictionary) EnumerateKeysAndObjectsWithOptionsUsingBlock(opts NSEnumerationOptions, block bool) {
	objc.Send[objc.ID](d.ID, objc.Sel("enumerateKeysAndObjectsWithOptions:usingBlock:"), opts, block)
}

// Returns an array of the dictionary’s keys, in the order they would be in
// if the dictionary were sorted by its values.
//
// comparator: A selector that specifies the method to use to compare the values in the
// dictionary.
// 
// The `comparator` method should return [NSOrderedAscending] if the
// dictionary value is smaller than the argument, [NSOrderedDescending] if the
// dictionary value is larger than the argument, and [NSOrderedSame] if they
// are equal.
//
// # Return Value
// 
// An array of the dictionary’s keys, in the order they would be in if the
// dictionary were sorted by its values.
//
// # Discussion
// 
// Pairs of dictionary values are compared using the comparison method
// specified by `comparator`; the `comparator` message is sent to one of the
// values and has as its single argument the other value from the dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/keysSortedByValue(using:)
func (d NSDictionary) KeysSortedByValueUsingSelector(comparator objc.SEL) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("keysSortedByValueUsingSelector:"), comparator)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns an array of the dictionary’s keys, in the order they would be in
// if the dictionary were sorted by its values using a given comparator block.
//
// cmptr: A comparator block.
//
// # Return Value
// 
// An array of the dictionary’s keys, in the order they would be in if the
// dictionary were sorted by its values using `cmptr`.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/keysSortedByValue(comparator:)
func (d NSDictionary) KeysSortedByValueUsingComparator(cmptr NSComparator) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("keysSortedByValueUsingComparator:"), cmptr)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns an array of the dictionary’s keys, in the order they would be in
// if the dictionary were sorted by its values using a given comparator block
// and a specified set of options.
//
// opts: A bitmask of sort options.
//
// cmptr: A comparator block.
//
// # Return Value
// 
// An array of the dictionary’s keys, in the order they would be in if the
// dictionary were sorted by its values using `cmptr` with the options given
// in `opts`.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/keysSortedByValue(options:usingComparator:)
func (d NSDictionary) KeysSortedByValueWithOptionsUsingComparator(opts NSSortOptions, cmptr NSComparator) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("keysSortedByValueWithOptions:usingComparator:"), opts, cmptr)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns the set of keys whose corresponding value satisfies a constraint
// described by a block object.
//
// predicate: A block object that specifies constraints for values in the dictionary.
//
// # Return Value
// 
// The set of keys whose corresponding value satisfies `predicate`.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/keysOfEntries(passingTest:)
func (d NSDictionary) KeysOfEntriesPassingTest(predicate bool) INSSet {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("keysOfEntriesPassingTest:"), predicate)
	return NSSetFromID(rv)
}

// Returns the set of keys whose corresponding value satisfies a constraint
// described by a block object.
//
// opts: A bit mask of enumeration options.
//
// predicate: A block object that specifies constraints for values in the dictionary.
//
// # Return Value
// 
// The set of keys whose corresponding value satisfies `predicate`.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/keysOfEntries(options:passingTest:)
func (d NSDictionary) KeysOfEntriesWithOptionsPassingTest(opts NSEnumerationOptions, predicate bool) INSSet {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("keysOfEntriesWithOptions:passingTest:"), opts, predicate)
	return NSSetFromID(rv)
}

// Writes a property list representation of the contents of the dictionary to
// a given URL.
//
// url: The URL to which to write the dictionary.
//
// # Discussion
// 
// This method recursively validates that all the contained objects are
// property list objects (instances of [NSData], [NSDate], [NSNumber],
// [NSString], [NSArray], or [NSDictionary]) before writing out the file. The
// method throws an error if all the objects are not property list objects,
// because the resulting output wouldn’t be a valid property list.
// 
// If the dictionary’s contents are all property list objects, you can use
// the location written by this method to initialize a new dictionary with the
// instance method `NSDictionary/init()-4pv16`.
// 
// If you need greater control over the property list representation, use
// [NSPropertyListSerialization] instead.
// 
// For more information about property lists, see [Property List Programming
// Guide].
//
// [Property List Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/PropertyLists/Introduction/Introduction.html#//apple_ref/doc/uid/10000048i
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/write(to:)
func (d NSDictionary) WriteToURLError(url INSURL) (bool, error) {
			var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("writeToURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeToURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns the file’s size, in bytes.
//
// # Return Value
// 
// The value associated with the [size] key, as a [UInt64] in Swift (`unsigned
// long long` in Objective-C), or `0` if the file attributes dictionary has no
// entry for the key.
//
// [size]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/size
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/fileSize()
func (d NSDictionary) FileSize() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("fileSize"))
	return rv
}

// Returns the file type.
//
// # Return Value
// 
// The value associated with the [type] file attributes key, or `nil` if the
// file attributes dictionary has no entry for the key. For possible values,
// see [NSFileAttributeType].
//
// [type]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/type
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/fileType()
func (d NSDictionary) FileType() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("fileType"))
	return NSStringFromID(rv).String()
}

// Returns the file’s creation date.
//
// # Return Value
// 
// The value associated with the [creationDate] file attributes key, or `nil`
// if the file attributes dictionary has no entry for the key.
//
// [creationDate]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/creationDate
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/fileCreationDate()
func (d NSDictionary) FileCreationDate() INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("fileCreationDate"))
	return NSDateFromID(rv)
}

// Returns file’s modification date.
//
// # Return Value
// 
// The value associated with the [modificationDate] file attributes key, or
// `nil` if the file attributes dictionary has no entry for the key.
//
// [modificationDate]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/modificationDate
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/fileModificationDate()
func (d NSDictionary) FileModificationDate() INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("fileModificationDate"))
	return NSDateFromID(rv)
}

// Returns the file’s POSIX permissions.
//
// # Return Value
// 
// The value associated with the [posixPermissions] file attributes key as an
// [Int] (`unsigned long` in Objective-C), or `0` if the file attributes
// dictionary has no entry for the key.
//
// [posixPermissions]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/posixPermissions
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/filePosixPermissions()
func (d NSDictionary) FilePosixPermissions() uint {
	rv := objc.Send[uint](d.ID, objc.Sel("filePosixPermissions"))
	return rv
}

// Returns the file’s owner account ID.
//
// # Return Value
// 
// The value associated with the [ownerAccountID] file attributes key, or
// `nil` if the file attributes dictionary has no entry for the key.
//
// [ownerAccountID]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/ownerAccountID
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/fileOwnerAccountID()
func (d NSDictionary) FileOwnerAccountID() INSNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("fileOwnerAccountID"))
	return NSNumberFromID(rv)
}

// Returns the file’s owner account name.
//
// # Return Value
// 
// The value associated with the [ownerAccountName] file attributes key, or
// `nil` if the file attributes dictionary has no entry for the key.
//
// [ownerAccountName]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/ownerAccountName
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/fileOwnerAccountName()
func (d NSDictionary) FileOwnerAccountName() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("fileOwnerAccountName"))
	return NSStringFromID(rv).String()
}

// Returns file’s group owner account ID.
//
// # Return Value
// 
// The value associated with the [groupOwnerAccountID] file attributes key, or
// `nil` if the file attributes dictionary has no entry for the key.
//
// [groupOwnerAccountID]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/groupOwnerAccountID
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/fileGroupOwnerAccountID()
func (d NSDictionary) FileGroupOwnerAccountID() INSNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("fileGroupOwnerAccountID"))
	return NSNumberFromID(rv)
}

// Returns the file’s group owner account name.
//
// # Return Value
// 
// The value associated with the [groupOwnerAccountName] file attributes key,
// or `nil` if the file attributes dictionary has no entry for the key.
//
// [groupOwnerAccountName]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/groupOwnerAccountName
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/fileGroupOwnerAccountName()
func (d NSDictionary) FileGroupOwnerAccountName() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("fileGroupOwnerAccountName"))
	return NSStringFromID(rv).String()
}

// Returns a Boolean value indicating whether the file hides its extension.
//
// # Return Value
// 
// The value associated with the [extensionHidden] file attributes key, or
// [false] if the file attributes dictionary has no entry for the key.
//
// [extensionHidden]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/extensionHidden
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/fileExtensionHidden()
func (d NSDictionary) FileExtensionHidden() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("fileExtensionHidden"))
	return rv
}

// Returns a Boolean value indicating whether the file is immutable.
//
// # Return Value
// 
// The value associated with the [immutable] file attributes key, or [false]
// if the file attributes dictionary has no entry for the key.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [immutable]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/immutable
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/fileIsImmutable()
func (d NSDictionary) FileIsImmutable() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("fileIsImmutable"))
	return rv
}

// Returns a Boolean value indicating whether the file is append only.
//
// # Return Value
// 
// The value associated with the [appendOnly] file attributes key, or [false]
// if the file attributes dictionary has no entry for the key.
//
// [appendOnly]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/appendOnly
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/fileIsAppendOnly()
func (d NSDictionary) FileIsAppendOnly() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("fileIsAppendOnly"))
	return rv
}

// Returns the filesystem file number.
//
// # Return Value
// 
// The value associated with the [systemFileNumber] file attributes key as an
// [Int] (`unsigned long` in Objective-C), or `0` if the file attributes
// dictionary has no entry for the key.
//
// [systemFileNumber]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/systemFileNumber
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/fileSystemFileNumber()
func (d NSDictionary) FileSystemFileNumber() uint {
	rv := objc.Send[uint](d.ID, objc.Sel("fileSystemFileNumber"))
	return rv
}

// Returns the filesystem number.
//
// # Return Value
// 
// The value associated with the [systemNumber] file attributes key as an
// [Int] (`unsigned long` in Objective-C), or `0` if the file attributes
// dictionary has no entry for the key
//
// [systemNumber]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/systemNumber
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/fileSystemNumber()
func (d NSDictionary) FileSystemNumber() int {
	rv := objc.Send[int](d.ID, objc.Sel("fileSystemNumber"))
	return rv
}

// Returns file’s HFS type code.
//
// # Return Value
// 
// The value associated with the [hfsTypeCode] file attributes key, or `0` if
// the file attributes dictionary has no entry for the key.
//
// [hfsTypeCode]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/hfsTypeCode
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/fileHFSTypeCode()
func (d NSDictionary) FileHFSTypeCode() uint32 {
	rv := objc.Send[uint32](d.ID, objc.Sel("fileHFSTypeCode"))
	return rv
}

// Returns the file’s HFS creator code.
//
// # Return Value
// 
// The value associated with the [hfsCreatorCode] file attributes key, or `0`
// if the file attributes dictionary has no entry for the key.
//
// [hfsCreatorCode]: https://developer.apple.com/documentation/Foundation/FileAttributeKey/hfsCreatorCode
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/fileHFSCreatorCode()
func (d NSDictionary) FileHFSCreatorCode() uint32 {
	rv := objc.Send[uint32](d.ID, objc.Sel("fileHFSCreatorCode"))
	return rv
}

// Returns a string object that represents the contents of the dictionary,
// formatted as a property list.
//
// locale: An object that specifies options used for formatting each of the
// dictionary’s keys and values; pass `nil` if you don’t want them
// formatted.
// 
// On iOS and macOS 10.5 and later, either an instance of [NSDictionary] or an
// [NSLocale] object may be used for `locale`. In OS X v10.4 and earlier it
// must be an instance of [NSDictionary].
//
// # Discussion
// 
// For a description of how `locale` is applied to each element in the
// dictionary, see [DescriptionWithLocaleIndent].
// 
// If each key in the dictionary responds to ``, the entries are listed in
// ascending order by key, otherwise the order in which the entries are listed
// is undefined.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/description(withLocale:)
func (d NSDictionary) DescriptionWithLocale(locale objectivec.IObject) string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("descriptionWithLocale:"), locale)
	return NSStringFromID(rv).String()
}

// Returns a string object that represents the contents of the dictionary,
// formatted as a property list.
//
// locale: An object that specifies options used for formatting each of the
// dictionary’s keys and values; pass `nil` if you don’t want them
// formatted.
// 
// On iOS and macOS 10.5 and later, either an instance of [NSDictionary] or an
// [NSLocale] object may be used for `locale`. In OS X v10.4 and earlier it
// must be an instance of [NSDictionary].
//
// level: Specifies a level of indentation, to make the output more readable: the
// indentation is (4 spaces) * `level`.
//
// # Return Value
// 
// A string object that represents the contents of the dictionary, formatted
// as a property list.
//
// # Discussion
// 
// The returned [NSString] object contains the string representations of each
// of the dictionary’s entries. [DescriptionWithLocaleIndent] obtains the
// string representation of a given key or value as follows:
// 
// - If the object is an [NSString] object, it is used as is. - If the object
// responds to ``, that method is invoked to obtain the object’s string
// representation. - If the object responds to ``, that method is invoked to
// obtain the object’s string representation. - If none of the above
// conditions is met, the object’s string representation is obtained by
// through its `description` property.
// 
// If each key in the dictionary responds to ``, the entries are listed in
// ascending order, by key. Otherwise, the order in which the entries are
// listed is undefined.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/description(withLocale:indent:)
func (d NSDictionary) DescriptionWithLocaleIndent(locale objectivec.IObject, level uint) string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("descriptionWithLocale:indent:"), locale, level)
	return NSStringFromID(rv).String()
}

//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/init(contentsOf:error:)
func (d NSDictionary) InitWithContentsOfURLError(url INSURL) (NSDictionary, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSDictionary{}, NSErrorFrom(errorPtr)
	}
	return NSDictionaryFromID(rv), nil

}

// Returns by reference a C array of objects over which the sender should
// iterate.
//
// state: Context information that is used in the enumeration to, in addition to
// other possibilities, ensure that the collection has not been mutated.
//
// buffer: A C array of objects over which the sender is to iterate.
//
// len: The maximum number of objects to return in `buffer`.
//
// # Return Value
// 
// The number of objects returned in `buffer`. Returns `0` when the iteration
// is finished.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/countByEnumeratingWithState:objects:count:
func (d NSDictionary) CountByEnumeratingWithStateObjectsCount(state NSFastEnumerationState, buffer []objectivec.IObject, len_ uint) uint {
	rv := objc.Send[uint](d.ID, objc.Sel("countByEnumeratingWithState:objects:count:"), state, objc.CArray(buffer), len_)
	return rv
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (d NSDictionary) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns by reference C arrays of the keys and values in the dictionary.
//
// objects: Upon return, contains a C array of the values in the dictionary.
//
// keys: Upon return, contains a C array of the keys in the dictionary.
//
// count: The maximum number of objects to return.
//
// # Discussion
// 
// The elements in the returned array and the keys array have a one-for-one
// correspondence, so that the nth object in the returned array corresponds to
// the the key in keys.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/getObjects:andKeys:count:
func (d NSDictionary) GetObjectsAndKeysCount(objects []objectivec.IObject, keys []objectivec.IObject, count uint) {
	objc.Send[objc.ID](d.ID, objc.Sel("getObjects:andKeys:count:"), objc.CArray(objects), objc.CArray(keys), count)
}

// Initializes a newly allocated dictionary with entries constructed from the
// specified set of values and keys.
//
// firstObject: The first value to add to the new dictionary.
//
// # Discussion
// 
// After the `firstObject` value, pass the key for `firstObject`, then a
// null-terminated list of alternating values and keys. If any key is `nil`,
// an [invalidArgumentException] is raised.
// 
// This method is similar to [InitWithObjectsForKeys], differing only in the
// way in which the key-value pairs are specified.
// 
// For example:
//
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/initWithObjectsAndKeys:
func (d NSDictionary) InitWithObjectsAndKeys(firstObject objectivec.IObject) NSDictionary {
	rv := objc.Send[NSDictionary](d.ID, objc.Sel("initWithObjectsAndKeys:"), firstObject)
	return rv
}





// Creates a shared key set object for the specified keys.
//
// keys: The array of keys. If the parameter is nil, an exception is thrown. If the
// array of keys is empty, an empty key set is returned.
//
// # Return Value
// 
// A shared key set object.
//
// # Discussion
// 
// The array of `keys` may contain duplicates which are quietly ignored.
// Duplicate hash values of the keys are quietly allowed, but may cause lower
// performance and increase memory usage.
// 
// Typically you would create a shared key set for a given set of keys once,
// before creating shared key dictionaries, and retain and save the result of
// this method for use with the [NSMutableDictionary] class method `.`
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/sharedKeySet(forKeys:)
func (_NSDictionaryClass NSDictionaryClass) SharedKeySetForKeys(keys []objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NSDictionaryClass.class), objc.Sel("sharedKeySetForKeys:"), objectivec.IObjectSliceToNSArray(keys))
	return objectivec.Object{ID: rv}
}

// Creates an empty dictionary.
//
// # Return Value
// 
// A new empty dictionary.
//
// # Discussion
// 
// This method is declared primarily for use with mutable subclasses of
// [NSDictionary].
// 
// If you don’t want a temporary object, you can also create an empty
// dictionary using `alloc` and [Init].
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionary
func (_NSDictionaryClass NSDictionaryClass) Dictionary() NSDictionary {
	rv := objc.Send[objc.ID](objc.ID(_NSDictionaryClass.class), objc.Sel("dictionary"))
	return NSDictionaryFromID(rv)
}

// Creates a dictionary using the keys and values found in a resource
// specified by a given URL.
//
// url: A URL that identifies a resource containing a string representation of a
// property list whose root object is a dictionary.
//
// # Return Value
// 
// A new dictionary that contains the dictionary at `url`, or `nil` if there
// is an error or if the contents of the resource are an invalid
// representation of a dictionary.
//
// # Discussion
// 
// The dictionary representation in the file identified by path must contain
// only property list objects ([NSString], [NSData], [NSDate], [NSNumber],
// [NSArray], or [NSDictionary] objects). For more details, see [Property List
// Programming Guide]. The objects contained by this dictionary are immutable,
// even if the dictionary is mutable.
//
// [Property List Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/PropertyLists/Introduction/Introduction.html#//apple_ref/doc/uid/10000048i
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionaryWithContentsOfURL:error:
func (_NSDictionaryClass NSDictionaryClass) DictionaryWithContentsOfURLError(url INSURL) (INSDictionary, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSDictionaryClass.class), objc.Sel("dictionaryWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return NSDictionaryFromID(rv), nil

}

// Creates a dictionary containing the keys and values from another given
// dictionary.
//
// dict: A dictionary containing the keys and values with which to initialize the
// new dictionary.
//
// # Return Value
// 
// A new dictionary containing the keys and values found in `dict`.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionaryWithDictionary:
func (_NSDictionaryClass NSDictionaryClass) DictionaryWithDictionary(dict INSDictionary) NSDictionary {
	rv := objc.Send[objc.ID](objc.ID(_NSDictionaryClass.class), objc.Sel("dictionaryWithDictionary:"), dict)
	return NSDictionaryFromID(rv)
}

// Creates a dictionary containing entries constructed from the specified set
// of values and keys.
//
// firstObject: The first value to add to the new dictionary.
//
// # Discussion
// 
// After passing `firstObj`, pass a null-terminated list of alternating values
// and keys as variadic arguments. If any key is `nil`, an
// [invalidArgumentException] is raised.
// 
// This method is similar to [DictionaryWithObjectsForKeys], differing only in
// the way key-value pairs are specified.
// 
// For example:
//
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionaryWithObjectsAndKeys:
func (_NSDictionaryClass NSDictionaryClass) DictionaryWithObjectsAndKeys(firstObject objectivec.IObject) NSDictionary {
	rv := objc.Send[objc.ID](objc.ID(_NSDictionaryClass.class), objc.Sel("dictionaryWithObjectsAndKeys:"), firstObject)
	return NSDictionaryFromID(rv)
}

// Creates a dictionary containing entries constructed from the contents of an
// array of keys and an array of values.
//
// objects: An array containing the values for the new dictionary.
//
// keys: An array containing the keys for the new dictionary. Each key is copied
// (using [CopyWithZone]; keys must conform to the [NSCopying] protocol), and
// the copy is added to the dictionary.
//
// # Return Value
// 
// A new dictionary containing entries constructed from the contents of
// `objects` and `keys`.
//
// # Discussion
// 
// This method steps through the `objects` and `keys` arrays, creating entries
// in the new dictionary as it goes. An [NSInvalidArgumentException] is raised
// if objects and keys don’t have the same number of elements.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionaryWithObjects:forKeys:
func (_NSDictionaryClass NSDictionaryClass) DictionaryWithObjectsForKeys(objects []objectivec.IObject, keys []objectivec.IObject) NSDictionary {
	rv := objc.Send[objc.ID](objc.ID(_NSDictionaryClass.class), objc.Sel("dictionaryWithObjects:forKeys:"), objectivec.IObjectSliceToNSArray(objects), objectivec.IObjectSliceToNSArray(keys))
	return NSDictionaryFromID(rv)
}

// Creates a dictionary containing a specified number of objects from a C
// array.
//
// objects: A C array of values for the new dictionary.
//
// keys: A C array of keys for the new dictionary. Each key is copied (using
// [CopyWithZone]; keys must conform to the [NSCopying] protocol), and the
// copy is added to the new dictionary.
//
// cnt: The number of elements to use from the `keys` and `objects` arrays. `cnt`
// must not exceed the number of elements in `objects` or `keys`.
//
// # Discussion
// 
// This method steps through the `objects` and `keys` arrays, creating entries
// in the new dictionary as it goes. An [NSInvalidArgumentException] is raised
// if a key or value object is `nil`.
// 
// The following code fragment illustrates how to create a dictionary that
// associates the alphabetic characters with their ASCII values:
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionaryWithObjects:forKeys:count:
func (_NSDictionaryClass NSDictionaryClass) DictionaryWithObjectsForKeysCount(objects []objectivec.IObject, keys []NSCopying, cnt uint) NSDictionary {
	rv := objc.Send[objc.ID](objc.ID(_NSDictionaryClass.class), objc.Sel("dictionaryWithObjects:forKeys:count:"), objc.CArray(objects), objc.CArray(keys), cnt)
	return NSDictionaryFromID(rv)
}








// The number of entries in the dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/count
func (d NSDictionary) Count() uint {
	rv := objc.Send[uint](d.ID, objc.Sel("count"))
	return rv
}



// A new array containing the dictionary’s keys, or an empty array if the
// dictionary has no entries.
//
// # Discussion
// 
// The order of the elements in the array is not defined.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/allKeys
func (d NSDictionary) AllKeys() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("allKeys"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}



// A new array containing the dictionary’s values, or an empty array if the
// dictionary has no entries.
//
// # Discussion
// 
// The order of the values in the array isn’t defined.
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/allValues
func (d NSDictionary) AllValues() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("allValues"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}



// A string that represents the contents of the dictionary, formatted as a
// property list.
//
// # Discussion
// 
// If each key in the dictionary is an [NSString] object, the entries are
// listed in ascending order by key, otherwise the order in which the entries
// are listed is undefined. This property is intended to produce readable
// output for debugging purposes, not for serializing data. If you want to
// store dictionary data for later retrieval, see [Property List Programming
// Guide] and [Archives and Serializations Programming Guide].
//
// [Archives and Serializations Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Archiving/Archiving.html#//apple_ref/doc/uid/10000047i
// [Property List Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/PropertyLists/Introduction/Introduction.html#//apple_ref/doc/uid/10000048i
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/description
func (d NSDictionary) Description() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("description"))
	return NSStringFromID(rv).String()
}



// A string that represents the contents of the dictionary, formatted in
// `XCUIElementTypeStrings` file format.
//
// # Discussion
// 
// The order in which the entries are listed is undefined.
// 
// This method fails unless the dictionary can be represented by a strings
// resource file. For details, see [String Resources] in [Resource Programming
// Guide].
//
// [Resource Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/LoadingResources/Introduction/Introduction.html#//apple_ref/doc/uid/10000051i
// [String Resources]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/LoadingResources/Strings/Strings.html#//apple_ref/doc/uid/10000051i-CH6
//
// See: https://developer.apple.com/documentation/Foundation/NSDictionary/descriptionInStringsFileFormat
func (d NSDictionary) DescriptionInStringsFileFormat() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("descriptionInStringsFileFormat"))
	return NSStringFromID(rv).String()
}


















			// Protocol methods for NSCopying
			






			// Protocol methods for NSMutableCopying
			
















