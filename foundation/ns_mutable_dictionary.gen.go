// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMutableDictionary] class.
var (
	_NSMutableDictionaryClass     NSMutableDictionaryClass
	_NSMutableDictionaryClassOnce sync.Once
)

func getNSMutableDictionaryClass() NSMutableDictionaryClass {
	_NSMutableDictionaryClassOnce.Do(func() {
		_NSMutableDictionaryClass = NSMutableDictionaryClass{class: objc.GetClass("NSMutableDictionary")}
	})
	return _NSMutableDictionaryClass
}

// GetNSMutableDictionaryClass returns the class object for NSMutableDictionary.
func GetNSMutableDictionaryClass() NSMutableDictionaryClass {
	return getNSMutableDictionaryClass()
}

type NSMutableDictionaryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSMutableDictionaryClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMutableDictionaryClass) Alloc() NSMutableDictionary {
	rv := objc.Send[NSMutableDictionary](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A dynamic collection of objects associated with unique keys.
//
// # Overview
// 
// In Swift, you can use this type instead of a [Dictionary] variable in cases
// that require reference semantics.
// 
// The [NSMutableDictionary] class declares the programmatic interface to
// objects that manage mutable associations of keys and values. It adds
// modification operations to the basic operations it inherits from
// [NSDictionary].
// 
// [NSMutableDictionary] is “toll-free bridged” with its Core Foundation
// counterpart, [CFMutableDictionary]. See [Toll-Free Bridging] for more
// information on toll-free bridging.
// 
// # Setting Values Using Subscripting
// 
// In addition to the provided instance methods, such as [NSMutableDictionary.SetObjectForKey],
// you can access [NSDictionary] values by their keys using .
// 
// # Subclassing Notes
// 
// There should typically be little need to subclass [NSMutableDictionary]. If
// you do need to customize behavior, it is often better to consider
// composition rather than subclassing.
// 
// # Methods to Override
// 
// In a subclass, you must override both of its primitive methods:
// 
// - [NSMutableDictionary.SetObjectForKey]
// - [NSMutableDictionary.RemoveObjectForKey]
// 
// You must also override the primitive methods of the [NSDictionary] class.
//
// [CFMutableDictionary]: https://developer.apple.com/documentation/CoreFoundation/CFMutableDictionary
// [Dictionary]: https://developer.apple.com/documentation/Swift/Dictionary
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
//
// # Creating and Initializing a Mutable Dictionary
//
//   - [NSMutableDictionary.InitWithCapacity]: Initializes a newly allocated mutable dictionary, allocating enough memory to hold `numItems` entries.
//
// # Adding Entries to a Mutable Dictionary
//
//   - [NSMutableDictionary.SetObjectForKey]: Adds a given key-value pair to the dictionary.
//   - [NSMutableDictionary.AddEntriesFromDictionary]: Adds to the receiving dictionary the entries from another dictionary.
//   - [NSMutableDictionary.SetDictionary]: Sets the contents of the receiving dictionary to entries in a given dictionary.
//
// # Removing Entries From a Mutable Dictionary
//
//   - [NSMutableDictionary.RemoveObjectForKey]: Removes a given key and its associated value from the dictionary.
//   - [NSMutableDictionary.RemoveAllObjects]: Empties the dictionary of its entries.
//   - [NSMutableDictionary.RemoveObjectsForKeys]: Removes from the dictionary entries specified by elements in a given array.
//
// # Instance Methods
//
//   - [NSMutableDictionary.AddApplicationParameterHeaderLength]
//   - [NSMutableDictionary.AddAuthorizationChallengeHeaderLength]
//   - [NSMutableDictionary.AddAuthorizationResponseHeaderLength]
//   - [NSMutableDictionary.AddBodyHeaderLengthEndOfBody]
//   - [NSMutableDictionary.AddByteSequenceHeaderLength]
//   - [NSMutableDictionary.AddConnectionIDHeaderLength]
//   - [NSMutableDictionary.AddCountHeader]
//   - [NSMutableDictionary.AddDescriptionHeader]
//   - [NSMutableDictionary.AddHTTPHeaderLength]
//   - [NSMutableDictionary.AddImageDescriptorHeaderLength]
//   - [NSMutableDictionary.AddImageHandleHeader]
//   - [NSMutableDictionary.AddLengthHeader]
//   - [NSMutableDictionary.AddNameHeader]
//   - [NSMutableDictionary.AddObjectClassHeaderLength]
//   - [NSMutableDictionary.AddTargetHeaderLength]
//   - [NSMutableDictionary.AddTime4ByteHeader]
//   - [NSMutableDictionary.AddTimeISOHeaderLength]
//   - [NSMutableDictionary.AddTypeHeader]
//   - [NSMutableDictionary.AddUserDefinedHeaderLength]
//   - [NSMutableDictionary.AddWhoHeaderLength]
//   - [NSMutableDictionary.GetHeaderBytes]
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary
type NSMutableDictionary struct {
	NSDictionary
}

// NSMutableDictionaryFromID constructs a [NSMutableDictionary] from an objc.ID.
//
// A dynamic collection of objects associated with unique keys.
func NSMutableDictionaryFromID(id objc.ID) NSMutableDictionary {
	return NSMutableDictionary{NSDictionary: NSDictionaryFromID(id)}
}
// NOTE: NSMutableDictionary adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMutableDictionary] class.
//
// # Creating and Initializing a Mutable Dictionary
//
//   - [INSMutableDictionary.InitWithCapacity]: Initializes a newly allocated mutable dictionary, allocating enough memory to hold `numItems` entries.
//
// # Adding Entries to a Mutable Dictionary
//
//   - [INSMutableDictionary.SetObjectForKey]: Adds a given key-value pair to the dictionary.
//   - [INSMutableDictionary.AddEntriesFromDictionary]: Adds to the receiving dictionary the entries from another dictionary.
//   - [INSMutableDictionary.SetDictionary]: Sets the contents of the receiving dictionary to entries in a given dictionary.
//
// # Removing Entries From a Mutable Dictionary
//
//   - [INSMutableDictionary.RemoveObjectForKey]: Removes a given key and its associated value from the dictionary.
//   - [INSMutableDictionary.RemoveAllObjects]: Empties the dictionary of its entries.
//   - [INSMutableDictionary.RemoveObjectsForKeys]: Removes from the dictionary entries specified by elements in a given array.
//
// # Instance Methods
//
//   - [INSMutableDictionary.AddApplicationParameterHeaderLength]
//   - [INSMutableDictionary.AddAuthorizationChallengeHeaderLength]
//   - [INSMutableDictionary.AddAuthorizationResponseHeaderLength]
//   - [INSMutableDictionary.AddBodyHeaderLengthEndOfBody]
//   - [INSMutableDictionary.AddByteSequenceHeaderLength]
//   - [INSMutableDictionary.AddConnectionIDHeaderLength]
//   - [INSMutableDictionary.AddCountHeader]
//   - [INSMutableDictionary.AddDescriptionHeader]
//   - [INSMutableDictionary.AddHTTPHeaderLength]
//   - [INSMutableDictionary.AddImageDescriptorHeaderLength]
//   - [INSMutableDictionary.AddImageHandleHeader]
//   - [INSMutableDictionary.AddLengthHeader]
//   - [INSMutableDictionary.AddNameHeader]
//   - [INSMutableDictionary.AddObjectClassHeaderLength]
//   - [INSMutableDictionary.AddTargetHeaderLength]
//   - [INSMutableDictionary.AddTime4ByteHeader]
//   - [INSMutableDictionary.AddTimeISOHeaderLength]
//   - [INSMutableDictionary.AddTypeHeader]
//   - [INSMutableDictionary.AddUserDefinedHeaderLength]
//   - [INSMutableDictionary.AddWhoHeaderLength]
//   - [INSMutableDictionary.GetHeaderBytes]
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary
type INSMutableDictionary interface {
	INSDictionary

	// Topic: Creating and Initializing a Mutable Dictionary

	// Initializes a newly allocated mutable dictionary, allocating enough memory to hold `numItems` entries.
	InitWithCapacity(numItems uint) NSMutableDictionary

	// Topic: Adding Entries to a Mutable Dictionary

	// Adds a given key-value pair to the dictionary.
	SetObjectForKey(anObject objectivec.IObject, aKey NSCopying)
	// Adds to the receiving dictionary the entries from another dictionary.
	AddEntriesFromDictionary(otherDictionary INSDictionary)
	// Sets the contents of the receiving dictionary to entries in a given dictionary.
	SetDictionary(otherDictionary INSDictionary)

	// Topic: Removing Entries From a Mutable Dictionary

	// Removes a given key and its associated value from the dictionary.
	RemoveObjectForKey(aKey objectivec.IObject)
	// Empties the dictionary of its entries.
	RemoveAllObjects()
	// Removes from the dictionary entries specified by elements in a given array.
	RemoveObjectsForKeys(keyArray []objectivec.IObject)

	// Topic: Instance Methods

	AddApplicationParameterHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject
	AddAuthorizationChallengeHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject
	AddAuthorizationResponseHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject
	AddBodyHeaderLengthEndOfBody(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, isEndOfBody bool) objectivec.IObject
	AddByteSequenceHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject
	AddConnectionIDHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject
	AddCountHeader(inCount uint32) objectivec.IObject
	AddDescriptionHeader(inDescriptionString string) objectivec.IObject
	AddHTTPHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject
	AddImageDescriptorHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject
	AddImageHandleHeader(type_ string) objectivec.IObject
	AddLengthHeader(length uint32) objectivec.IObject
	AddNameHeader(inNameString string) objectivec.IObject
	AddObjectClassHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject
	AddTargetHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject
	AddTime4ByteHeader(time4Byte uint32) objectivec.IObject
	AddTimeISOHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject
	AddTypeHeader(type_ string) objectivec.IObject
	AddUserDefinedHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject
	AddWhoHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject
	GetHeaderBytes() INSMutableData

	// Adds a given key-value pair to the dictionary.
	SetObjectForKeyedSubscript(obj objectivec.IObject, key NSCopying)
}

// Init initializes the instance.
func (m NSMutableDictionary) Init() NSMutableDictionary {
	rv := objc.Send[NSMutableDictionary](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMutableDictionary) Autorelease() NSMutableDictionary {
	rv := objc.Send[NSMutableDictionary](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMutableDictionary creates a new NSMutableDictionary instance.
func NewNSMutableDictionary() NSMutableDictionary {
	class := getNSMutableDictionaryClass()
	rv := objc.Send[NSMutableDictionary](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a newly allocated mutable dictionary, allocating enough memory
// to hold `numItems` entries.
//
// numItems: The initial capacity of the initialized dictionary.
//
// # Return Value
// 
// An initialized mutable dictionary, which might be different than the
// original receiver.
//
// # Discussion
// 
// Mutable dictionaries allocate additional memory as needed, so `numItems`
// simply establishes the object’s initial capacity.
// 
// This method is a designated initializer of [NSMutableDictionary].
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/init(capacity:)
func NewMutableDictionaryWithCapacity(numItems uint) NSMutableDictionary {
	instance := getNSMutableDictionaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCapacity:"), numItems)
	return NSMutableDictionaryFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/init(coder:)
func NewMutableDictionaryWithCoder(coder INSCoder) NSMutableDictionary {
	instance := getNSMutableDictionaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSMutableDictionaryFromID(rv)
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
func NewMutableDictionaryWithDictionary(otherDictionary INSDictionary) NSMutableDictionary {
	instance := getNSMutableDictionaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDictionary:"), otherDictionary)
	return NSMutableDictionaryFromID(rv)
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
func NewMutableDictionaryWithDictionaryCopyItems(otherDictionary INSDictionary, flag bool) NSMutableDictionary {
	instance := getNSMutableDictionaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDictionary:copyItems:"), otherDictionary, flag)
	return NSMutableDictionaryFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/init(OBEXHeadersData:)
func NewMutableDictionaryWithOBEXHeadersData(inHeadersData INSData) NSMutableDictionary {
	rv := objc.Send[objc.ID](objc.ID(getNSMutableDictionaryClass().class), objc.Sel("dictionaryWithOBEXHeadersData:"), inHeadersData)
	return NSMutableDictionaryFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/init(OBEXHeadersData:headersDataSize:)
func NewMutableDictionaryWithOBEXHeadersDataHeadersDataSize(inHeadersData unsafe.Pointer, inDataSize uintptr) NSMutableDictionary {
	rv := objc.Send[objc.ID](objc.ID(getNSMutableDictionaryClass().class), objc.Sel("dictionaryWithOBEXHeadersData:headersDataSize:"), inHeadersData, inDataSize)
	return NSMutableDictionaryFromID(rv)
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
func NewMutableDictionaryWithObjectForKey(object objectivec.IObject, key NSCopying) NSMutableDictionary {
	rv := objc.Send[objc.ID](objc.ID(getNSMutableDictionaryClass().class), objc.Sel("dictionaryWithObject:forKey:"), object, key)
	return NSMutableDictionaryFromID(rv)
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
func NewMutableDictionaryWithObjectsAndKeys(firstObject objectivec.IObject) NSMutableDictionary {
	instance := getNSMutableDictionaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObjectsAndKeys:"), firstObject)
	return NSMutableDictionaryFromID(rv)
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
func NewMutableDictionaryWithObjectsForKeys(objects []objectivec.IObject, keys []objectivec.IObject) NSMutableDictionary {
	instance := getNSMutableDictionaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObjects:forKeys:"), objectivec.IObjectSliceToNSArray(objects), objectivec.IObjectSliceToNSArray(keys))
	return NSMutableDictionaryFromID(rv)
}

// Initializes a newly allocated mutable dictionary, allocating enough memory
// to hold `numItems` entries.
//
// numItems: The initial capacity of the initialized dictionary.
//
// # Return Value
// 
// An initialized mutable dictionary, which might be different than the
// original receiver.
//
// # Discussion
// 
// Mutable dictionaries allocate additional memory as needed, so `numItems`
// simply establishes the object’s initial capacity.
// 
// This method is a designated initializer of [NSMutableDictionary].
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/init(capacity:)
func (m NSMutableDictionary) InitWithCapacity(numItems uint) NSMutableDictionary {
	rv := objc.Send[NSMutableDictionary](m.ID, objc.Sel("initWithCapacity:"), numItems)
	return rv
}
// Adds a given key-value pair to the dictionary.
//
// anObject: The value for `aKey`. A strong reference to the object is maintained by the
// dictionary.
//
// aKey: The key for `value`. The key is copied (using [CopyWithZone]; keys must
// conform to the [NSCopying] protocol). If `aKey` already exists in the
// dictionary, `anObject` takes its place.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/setObject(_:forKey:)
func (m NSMutableDictionary) SetObjectForKey(anObject objectivec.IObject, aKey NSCopying) {
	objc.Send[objc.ID](m.ID, objc.Sel("setObject:forKey:"), anObject, aKey)
}
// Adds to the receiving dictionary the entries from another dictionary.
//
// otherDictionary: The dictionary from which to add entries
//
// # Discussion
// 
// Each value object from `otherDictionary` is sent a [retain] message before
// being added to the receiving dictionary. In contrast, each key object is
// copied (using [CopyWithZone]—keys must conform to the [NSCopying]
// protocol), and the copy is added to the receiving dictionary.
// 
// If both dictionaries contain the same key, the receiving dictionary’s
// previous value object for that key is sent a `release` message, and the new
// value object takes its place.
//
// [retain]: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/retain
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addEntries(from:)
func (m NSMutableDictionary) AddEntriesFromDictionary(otherDictionary INSDictionary) {
	objc.Send[objc.ID](m.ID, objc.Sel("addEntriesFromDictionary:"), otherDictionary)
}
// Sets the contents of the receiving dictionary to entries in a given
// dictionary.
//
// otherDictionary: A dictionary containing the new entries.
//
// # Discussion
// 
// All entries are removed from the receiving dictionary (with
// [RemoveAllObjects]), then each entry from `otherDictionary` added into the
// receiving dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/setDictionary(_:)
func (m NSMutableDictionary) SetDictionary(otherDictionary INSDictionary) {
	objc.Send[objc.ID](m.ID, objc.Sel("setDictionary:"), otherDictionary)
}
// Removes a given key and its associated value from the dictionary.
//
// aKey: The key to remove.
//
// # Discussion
// 
// Does nothing if `aKey` does not exist.
// 
// For example, assume you have an archived dictionary that records the call
// letters and associated frequencies of radio stations. To remove an entry
// for a defunct station, you could write code similar to the following:
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/removeObject(forKey:)
func (m NSMutableDictionary) RemoveObjectForKey(aKey objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeObjectForKey:"), aKey)
}
// Empties the dictionary of its entries.
//
// # Discussion
// 
// Each key and corresponding value object is sent a [release] message.
//
// [release]: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/release
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/removeAllObjects()
func (m NSMutableDictionary) RemoveAllObjects() {
	objc.Send[objc.ID](m.ID, objc.Sel("removeAllObjects"))
}
// Removes from the dictionary entries specified by elements in a given array.
//
// keyArray: An array of objects specifying the keys to remove.
//
// # Discussion
// 
// If a key in `keyArray` does not exist, the entry is ignored.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/removeObjects(forKeys:)
func (m NSMutableDictionary) RemoveObjectsForKeys(keyArray []objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeObjectsForKeys:"), objectivec.IObjectSliceToNSArray(keyArray))
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addApplicationParameterHeader(_:length:)
func (m NSMutableDictionary) AddApplicationParameterHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addApplicationParameterHeader:length:"), inHeaderData, inHeaderDataLength)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addAuthorizationChallengeHeader(_:length:)
func (m NSMutableDictionary) AddAuthorizationChallengeHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addAuthorizationChallengeHeader:length:"), inHeaderData, inHeaderDataLength)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addAuthorizationResponseHeader(_:length:)
func (m NSMutableDictionary) AddAuthorizationResponseHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addAuthorizationResponseHeader:length:"), inHeaderData, inHeaderDataLength)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addBodyHeader(_:length:endOfBody:)
func (m NSMutableDictionary) AddBodyHeaderLengthEndOfBody(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, isEndOfBody bool) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addBodyHeader:length:endOfBody:"), inHeaderData, inHeaderDataLength, isEndOfBody)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addByteSequenceHeader(_:length:)
func (m NSMutableDictionary) AddByteSequenceHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addByteSequenceHeader:length:"), inHeaderData, inHeaderDataLength)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addConnectionIDHeader(_:length:)
func (m NSMutableDictionary) AddConnectionIDHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addConnectionIDHeader:length:"), inHeaderData, inHeaderDataLength)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addCountHeader(_:)
func (m NSMutableDictionary) AddCountHeader(inCount uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addCountHeader:"), inCount)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addDescriptionHeader(_:)
func (m NSMutableDictionary) AddDescriptionHeader(inDescriptionString string) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addDescriptionHeader:"), objc.String(inDescriptionString))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addHTTPHeader(_:length:)
func (m NSMutableDictionary) AddHTTPHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addHTTPHeader:length:"), inHeaderData, inHeaderDataLength)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addImageDescriptorHeader(_:length:)
func (m NSMutableDictionary) AddImageDescriptorHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addImageDescriptorHeader:length:"), inHeaderData, inHeaderDataLength)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addImageHandleHeader(_:)
func (m NSMutableDictionary) AddImageHandleHeader(type_ string) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addImageHandleHeader:"), objc.String(type_))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addLengthHeader(_:)
func (m NSMutableDictionary) AddLengthHeader(length uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addLengthHeader:"), length)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addNameHeader(_:)
func (m NSMutableDictionary) AddNameHeader(inNameString string) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addNameHeader:"), objc.String(inNameString))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addObjectClassHeader(_:length:)
func (m NSMutableDictionary) AddObjectClassHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addObjectClassHeader:length:"), inHeaderData, inHeaderDataLength)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addTargetHeader(_:length:)
func (m NSMutableDictionary) AddTargetHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addTargetHeader:length:"), inHeaderData, inHeaderDataLength)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addTime4ByteHeader(_:)
func (m NSMutableDictionary) AddTime4ByteHeader(time4Byte uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addTime4ByteHeader:"), time4Byte)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addTimeISOHeader(_:length:)
func (m NSMutableDictionary) AddTimeISOHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addTimeISOHeader:length:"), inHeaderData, inHeaderDataLength)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addTypeHeader(_:)
func (m NSMutableDictionary) AddTypeHeader(type_ string) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addTypeHeader:"), objc.String(type_))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addUserDefinedHeader(_:length:)
func (m NSMutableDictionary) AddUserDefinedHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addUserDefinedHeader:length:"), inHeaderData, inHeaderDataLength)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addWhoHeader(_:length:)
func (m NSMutableDictionary) AddWhoHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addWhoHeader:length:"), inHeaderData, inHeaderDataLength)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/getHeaderBytes()
func (m NSMutableDictionary) GetHeaderBytes() INSMutableData {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("getHeaderBytes"))
	return NSMutableDataFromID(rv)
}
// Adds a given key-value pair to the dictionary.
//
// obj: The value for `key`. A strong reference to the object is maintained by the
// dictionary.
// 
// Passing `nil` will cause any object corresponding to `key` to be removed
// from the dictionary.
//
// key: The key for `obj`. The key is copied (using [CopyWithZone]; keys must
// conform to the [NSCopying] protocol). If `key` already exists in the
// dictionary, `anObject` takes its place.
//
// # Discussion
// 
// This method has the same behavior as the [SetObjectForKey] method.
// 
// You shouldn’t need to call this method directly. Instead, this method is
// called when setting an object for a key using subscripting.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/setObject:forKeyedSubscript:
func (m NSMutableDictionary) SetObjectForKeyedSubscript(obj objectivec.IObject, key NSCopying) {
	objc.Send[objc.ID](m.ID, objc.Sel("setObject:forKeyedSubscript:"), obj, key)
}

// Creates a mutable dictionary which is optimized for dealing with a known
// set of keys.
//
// keyset: The `keyset`, created by the [NSDictionary] class method
// [SharedKeySetForKeys].
//
// # Return Value
// 
// A new mutable dictionary optimized for a known set of keys.
//
// # Discussion
// 
// Keys that are not in the key set can still be set in the dictionary, but
// that usage is not optimal.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/init(sharedKeySet:)
func (_NSMutableDictionaryClass NSMutableDictionaryClass) DictionaryWithSharedKeySet(keyset objectivec.IObject) INSDictionary {
	rv := objc.Send[objc.ID](objc.ID(_NSMutableDictionaryClass.class), objc.Sel("dictionaryWithSharedKeySet:"), keyset)
	return NSDictionaryFromID(rv)
}
// Creates and returns a mutable dictionary, initially giving it enough
// allocated memory to hold a given number of entries.
//
// numItems: The initial capacity of the new dictionary.
//
// # Return Value
// 
// A new mutable dictionary with enough allocated memory to hold `numItems`
// entries.
//
// # Discussion
// 
// Mutable dictionaries allocate additional memory as needed, so `numItems`
// simply establishes the object’s initial capacity.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/dictionaryWithCapacity:
func (_NSMutableDictionaryClass NSMutableDictionaryClass) DictionaryWithCapacity(numItems uint) NSMutableDictionary {
	rv := objc.Send[objc.ID](objc.ID(_NSMutableDictionaryClass.class), objc.Sel("dictionaryWithCapacity:"), numItems)
	return NSMutableDictionaryFromID(rv)
}

