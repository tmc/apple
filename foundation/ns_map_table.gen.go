// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMapTable] class.
var (
	_NSMapTableClass     NSMapTableClass
	_NSMapTableClassOnce sync.Once
)

func getNSMapTableClass() NSMapTableClass {
	_NSMapTableClassOnce.Do(func() {
		_NSMapTableClass = NSMapTableClass{class: objc.GetClass("NSMapTable")}
	})
	return _NSMapTableClass
}

// GetNSMapTableClass returns the class object for NSMapTable.
func GetNSMapTableClass() NSMapTableClass {
	return getNSMapTableClass()
}

type NSMapTableClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMapTableClass) Alloc() NSMapTable {
	rv := objc.Send[NSMapTable](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A collection similar to a dictionary, but with a broader range of available
// memory semantics.
//
// # Overview
// 
// The map table is modeled after [NSDictionary] with the following
// differences:
// 
// - Keys and/or values are optionally held “weakly” such that entries are
// removed when one of the objects is reclaimed. - Its keys or values may be
// copied on input or may use pointer identity for equality and hashing. - It
// can contain arbitrary pointers (its contents are not constrained to being
// objects).
// 
// You can configure an [NSMapTable] instance to operate on arbitrary pointers
// and not just objects, although typically you are encouraged to use the C
// function API for void * pointers. The object-based API (such as
// [NSMapTable.SetObjectForKey]) will not work for non-object pointers without
// type-casting.
// 
// When configuring map tables, note that only the options listed in
// [NSMapTableOptions] guarantee that the rest of the API will work
// correctly—including copying, archiving, and fast enumeration. While other
// [NSPointerFunctions] options are used for certain configurations, such as
// to hold arbitrary pointers, not all combinations of the options are valid.
// With some combinations the map table may not work correctly, or may not
// even be initialized correctly.
// 
// # Subclassing Notes
// 
// [NSMapTable] is not suitable for subclassing.
//
// # Creating and Initializing a Map Table
//
//   - [NSMapTable.InitWithKeyOptionsValueOptionsCapacity]: Returns a map table, initialized with the given options.
//   - [NSMapTable.InitWithKeyPointerFunctionsValuePointerFunctionsCapacity]: Returns a map table, initialized with the given functions.
//
// # Accessing Content
//
//   - [NSMapTable.ObjectForKey]: Returns a the value associated with a given key.
//   - [NSMapTable.KeyEnumerator]: Returns an enumerator object that lets you access each key in the map table.
//   - [NSMapTable.ObjectEnumerator]: Returns an enumerator object that lets you access each value in the map table.
//   - [NSMapTable.Count]: The number of key-value pairs in the map table.
//
// # Manipulating Content
//
//   - [NSMapTable.SetObjectForKey]: Adds a given key-value pair to the map table.
//   - [NSMapTable.RemoveObjectForKey]: Removes a given key and its associated value from the map table.
//   - [NSMapTable.RemoveAllObjects]: Empties the map table of its entries.
//
// # Creating a Dictionary Representation
//
//   - [NSMapTable.DictionaryRepresentation]: Returns a dictionary representation of the map table.
//
// # Accessing Pointer Functions
//
//   - [NSMapTable.KeyPointerFunctions]: The pointer functions the map table uses to manage keys.
//   - [NSMapTable.ValuePointerFunctions]: The pointer functions the map table uses to manage values.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable
type NSMapTable struct {
	objectivec.Object
}

// NSMapTableFromID constructs a [NSMapTable] from an objc.ID.
//
// A collection similar to a dictionary, but with a broader range of available
// memory semantics.
func NSMapTableFromID(id objc.ID) NSMapTable {
	return NSMapTable{objectivec.Object{ID: id}}
}
// NOTE: NSMapTable adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMapTable] class.
//
// # Creating and Initializing a Map Table
//
//   - [INSMapTable.InitWithKeyOptionsValueOptionsCapacity]: Returns a map table, initialized with the given options.
//   - [INSMapTable.InitWithKeyPointerFunctionsValuePointerFunctionsCapacity]: Returns a map table, initialized with the given functions.
//
// # Accessing Content
//
//   - [INSMapTable.ObjectForKey]: Returns a the value associated with a given key.
//   - [INSMapTable.KeyEnumerator]: Returns an enumerator object that lets you access each key in the map table.
//   - [INSMapTable.ObjectEnumerator]: Returns an enumerator object that lets you access each value in the map table.
//   - [INSMapTable.Count]: The number of key-value pairs in the map table.
//
// # Manipulating Content
//
//   - [INSMapTable.SetObjectForKey]: Adds a given key-value pair to the map table.
//   - [INSMapTable.RemoveObjectForKey]: Removes a given key and its associated value from the map table.
//   - [INSMapTable.RemoveAllObjects]: Empties the map table of its entries.
//
// # Creating a Dictionary Representation
//
//   - [INSMapTable.DictionaryRepresentation]: Returns a dictionary representation of the map table.
//
// # Accessing Pointer Functions
//
//   - [INSMapTable.KeyPointerFunctions]: The pointer functions the map table uses to manage keys.
//   - [INSMapTable.ValuePointerFunctions]: The pointer functions the map table uses to manage values.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable
type INSMapTable interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Creating and Initializing a Map Table

	// Returns a map table, initialized with the given options.
	InitWithKeyOptionsValueOptionsCapacity(keyOptions NSPointerFunctionsOptions, valueOptions NSPointerFunctionsOptions, initialCapacity uint) NSMapTable
	// Returns a map table, initialized with the given functions.
	InitWithKeyPointerFunctionsValuePointerFunctionsCapacity(keyFunctions INSPointerFunctions, valueFunctions INSPointerFunctions, initialCapacity uint) NSMapTable

	// Topic: Accessing Content

	// Returns a the value associated with a given key.
	ObjectForKey(aKey objectivec.IObject) objectivec.IObject
	// Returns an enumerator object that lets you access each key in the map table.
	KeyEnumerator() INSEnumerator
	// Returns an enumerator object that lets you access each value in the map table.
	ObjectEnumerator() INSEnumerator
	// The number of key-value pairs in the map table.
	Count() uint

	// Topic: Manipulating Content

	// Adds a given key-value pair to the map table.
	SetObjectForKey(anObject objectivec.IObject, aKey objectivec.IObject)
	// Removes a given key and its associated value from the map table.
	RemoveObjectForKey(aKey objectivec.IObject)
	// Empties the map table of its entries.
	RemoveAllObjects()

	// Topic: Creating a Dictionary Representation

	// Returns a dictionary representation of the map table.
	DictionaryRepresentation() INSDictionary

	// Topic: Accessing Pointer Functions

	// The pointer functions the map table uses to manage keys.
	KeyPointerFunctions() INSPointerFunctions
	// The pointer functions the map table uses to manage values.
	ValuePointerFunctions() INSPointerFunctions
}

// Init initializes the instance.
func (m NSMapTable) Init() NSMapTable {
	rv := objc.Send[NSMapTable](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMapTable) Autorelease() NSMapTable {
	rv := objc.Send[NSMapTable](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMapTable creates a new NSMapTable instance.
func NewNSMapTable() NSMapTable {
	class := getNSMapTableClass()
	rv := objc.Send[NSMapTable](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewMapTableWithCoder(coder INSCoder) NSMapTable {
	instance := getNSMapTableClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSMapTableFromID(rv)
}

// Returns a map table, initialized with the given options.
//
// keyOptions: A bit field that specifies the options for the keys in the map table. For
// possible values, see [NSMapTableOptions].
//
// valueOptions: A bit field that specifies the options for the values in the map table. For
// possible values, see [NSMapTableOptions].
//
// initialCapacity: The initial capacity of the map table. This is just a hint; the map table
// may subsequently grow and shrink as required.
//
// # Return Value
// 
// A map table initialized using the given options.
//
// # Discussion
// 
// `values` must contain entries at all the indexes specified in `keys`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable/init(keyOptions:valueOptions:capacity:)
func NewMapTableWithKeyOptionsValueOptionsCapacity(keyOptions NSPointerFunctionsOptions, valueOptions NSPointerFunctionsOptions, initialCapacity uint) NSMapTable {
	instance := getNSMapTableClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKeyOptions:valueOptions:capacity:"), keyOptions, valueOptions, initialCapacity)
	return NSMapTableFromID(rv)
}

// Returns a map table, initialized with the given functions.
//
// keyFunctions: The functions the map table uses to manage keys.
//
// valueFunctions: The functions the map table uses to manage values.
//
// initialCapacity: The initial capacity of the map table. This is just a hint; the map table
// may subsequently grow and shrink as required.
//
// # Return Value
// 
// A map table, initialized with the given functions.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable/init(keyPointerFunctions:valuePointerFunctions:capacity:)
func NewMapTableWithKeyPointerFunctionsValuePointerFunctionsCapacity(keyFunctions INSPointerFunctions, valueFunctions INSPointerFunctions, initialCapacity uint) NSMapTable {
	instance := getNSMapTableClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKeyPointerFunctions:valuePointerFunctions:capacity:"), keyFunctions, valueFunctions, initialCapacity)
	return NSMapTableFromID(rv)
}

// Returns a map table, initialized with the given options.
//
// keyOptions: A bit field that specifies the options for the keys in the map table. For
// possible values, see [NSMapTableOptions].
//
// valueOptions: A bit field that specifies the options for the values in the map table. For
// possible values, see [NSMapTableOptions].
//
// initialCapacity: The initial capacity of the map table. This is just a hint; the map table
// may subsequently grow and shrink as required.
//
// # Return Value
// 
// A map table initialized using the given options.
//
// # Discussion
// 
// `values` must contain entries at all the indexes specified in `keys`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable/init(keyOptions:valueOptions:capacity:)
func (m NSMapTable) InitWithKeyOptionsValueOptionsCapacity(keyOptions NSPointerFunctionsOptions, valueOptions NSPointerFunctionsOptions, initialCapacity uint) NSMapTable {
	rv := objc.Send[NSMapTable](m.ID, objc.Sel("initWithKeyOptions:valueOptions:capacity:"), keyOptions, valueOptions, initialCapacity)
	return rv
}
// Returns a map table, initialized with the given functions.
//
// keyFunctions: The functions the map table uses to manage keys.
//
// valueFunctions: The functions the map table uses to manage values.
//
// initialCapacity: The initial capacity of the map table. This is just a hint; the map table
// may subsequently grow and shrink as required.
//
// # Return Value
// 
// A map table, initialized with the given functions.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable/init(keyPointerFunctions:valuePointerFunctions:capacity:)
func (m NSMapTable) InitWithKeyPointerFunctionsValuePointerFunctionsCapacity(keyFunctions INSPointerFunctions, valueFunctions INSPointerFunctions, initialCapacity uint) NSMapTable {
	rv := objc.Send[NSMapTable](m.ID, objc.Sel("initWithKeyPointerFunctions:valuePointerFunctions:capacity:"), keyFunctions, valueFunctions, initialCapacity)
	return rv
}
// Returns a the value associated with a given key.
//
// aKey: The key for which to return the corresponding value.
//
// # Return Value
// 
// The value associated with `aKey`, or `nil` if no value is associated with
// `aKey`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable/object(forKey:)
func (m NSMapTable) ObjectForKey(aKey objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("objectForKey:"), aKey)
	return objectivec.Object{ID: rv}
}
// Returns an enumerator object that lets you access each key in the map
// table.
//
// # Return Value
// 
// An enumerator object that lets you access each key in the map table.
//
// # Discussion
// 
// The following code fragment illustrates how you might use the method.
// 
// # Special Considerations
// 
// It is more efficient to use the fast enumeration protocol (see
// [NSFastEnumeration]).
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable/keyEnumerator()
func (m NSMapTable) KeyEnumerator() INSEnumerator {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("keyEnumerator"))
	return NSEnumeratorFromID(rv)
}
// Returns an enumerator object that lets you access each value in the map
// table.
//
// # Return Value
// 
// An enumerator object that lets you access each value in the map table.
//
// # Discussion
// 
// The following code fragment illustrates how you might use the method.
// 
// # Special Considerations
// 
// It is more efficient to use the fast enumeration protocol (see
// [NSFastEnumeration]).
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable/objectEnumerator()
func (m NSMapTable) ObjectEnumerator() INSEnumerator {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("objectEnumerator"))
	return NSEnumeratorFromID(rv)
}
// Adds a given key-value pair to the map table.
//
// anObject: The value for `aKey`.
//
// aKey: The key for `anObject`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable/setObject(_:forKey:)
func (m NSMapTable) SetObjectForKey(anObject objectivec.IObject, aKey objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setObject:forKey:"), anObject, aKey)
}
// Removes a given key and its associated value from the map table.
//
// aKey: The key to remove.
//
// # Discussion
// 
// Does nothing if `aKey` does not exist.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable/removeObject(forKey:)
func (m NSMapTable) RemoveObjectForKey(aKey objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeObjectForKey:"), aKey)
}
// Empties the map table of its entries.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable/removeAllObjects()
func (m NSMapTable) RemoveAllObjects() {
	objc.Send[objc.ID](m.ID, objc.Sel("removeAllObjects"))
}
// Returns a dictionary representation of the map table.
//
// # Return Value
// 
// A dictionary representation of the map table.
//
// # Discussion
// 
// The map table’s values and keys must conform to all the requirements
// specified in [SetObjectForKey] in [NSMutableDictionary].
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable/dictionaryRepresentation()
func (m NSMapTable) DictionaryRepresentation() INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("dictionaryRepresentation"))
	return NSDictionaryFromID(rv)
}
// Returns by reference a C array of objects over which the sender should
// iterate, and as the return value the number of objects in the array.
//
// state: Context information that is used in the enumeration to, in addition to
// other possibilities, ensure that the collection has not been mutated.
//
// buffer: A C array of objects over which the sender is to iterate.
//
// len: The maximum number of objects to return in `stackbuf`.
//
// # Return Value
// 
// The number of objects returned in `stackbuf`. Returns `0` when the
// iteration is finished.
//
// # Discussion
// 
// The state structure is assumed to be of stack local memory, so you can
// recast the passed in state structure to one more suitable for your
// iteration.
//
// See: https://developer.apple.com/documentation/Foundation/NSFastEnumeration/countByEnumerating(with:objects:count:)
func (m NSMapTable) CountByEnumeratingWithStateObjectsCount(state NSFastEnumerationState, buffer []objectivec.IObject, len_ uint) uint {
	rv := objc.Send[uint](m.ID, objc.Sel("countByEnumeratingWithState:objects:count:"), state, objc.CArray(buffer), len_)
	return rv
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (m NSMapTable) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (m NSMapTable) InitWithCoder(coder INSCoder) NSMapTable {
	rv := objc.Send[NSMapTable](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Returns a new map table, initialized with the given options
//
// keyOptions: A bit field that specifies the options for the keys in the map table. For
// possible values, see [NSMapTableOptions].
//
// valueOptions: A bit field that specifies the options for the values in the map table. For
// possible values, see [NSMapTableOptions].
//
// # Return Value
// 
// A new map table, initialized with the given options.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable/init(keyOptions:valueOptions:)
func (_NSMapTableClass NSMapTableClass) MapTableWithKeyOptionsValueOptions(keyOptions NSPointerFunctionsOptions, valueOptions NSPointerFunctionsOptions) NSMapTable {
	rv := objc.Send[objc.ID](objc.ID(_NSMapTableClass.class), objc.Sel("mapTableWithKeyOptions:valueOptions:"), keyOptions, valueOptions)
	return NSMapTableFromID(rv)
}
// Returns a new map table object which has strong references to the keys and
// values.
//
// # Return Value
// 
// A new map table object which has strong references to the keys and values.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable/strongToStrongObjects()
func (_NSMapTableClass NSMapTableClass) StrongToStrongObjectsMapTable() NSMapTable {
	rv := objc.Send[objc.ID](objc.ID(_NSMapTableClass.class), objc.Sel("strongToStrongObjectsMapTable"))
	return NSMapTableFromID(rv)
}
// Returns a new map table object which has weak references to the keys and
// strong references to the values.
//
// # Return Value
// 
// A new map table object which has weak references to the keys and strong
// references to the values.
//
// # Discussion
// 
// Use of weak-to-strong map tables is not recommended. The strong values for
// weak keys which get zeroed out continue to be maintained until the map
// table resizes itself.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable/weakToStrongObjects()
func (_NSMapTableClass NSMapTableClass) WeakToStrongObjectsMapTable() NSMapTable {
	rv := objc.Send[objc.ID](objc.ID(_NSMapTableClass.class), objc.Sel("weakToStrongObjectsMapTable"))
	return NSMapTableFromID(rv)
}
// Returns a new map table object which has strong references to the keys and
// weak references to the values.
//
// # Return Value
// 
// A new map table object which has strong references to the keys and weak
// references to the values.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable/strongToWeakObjects()
func (_NSMapTableClass NSMapTableClass) StrongToWeakObjectsMapTable() NSMapTable {
	rv := objc.Send[objc.ID](objc.ID(_NSMapTableClass.class), objc.Sel("strongToWeakObjectsMapTable"))
	return NSMapTableFromID(rv)
}
// Returns a new map table object which has weak references to the keys and
// values.
//
// # Return Value
// 
// A new map table object which has weak references to the keys and values.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable/weakToWeakObjects()
func (_NSMapTableClass NSMapTableClass) WeakToWeakObjectsMapTable() NSMapTable {
	rv := objc.Send[objc.ID](objc.ID(_NSMapTableClass.class), objc.Sel("weakToWeakObjectsMapTable"))
	return NSMapTableFromID(rv)
}

// The number of key-value pairs in the map table.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable/count
func (m NSMapTable) Count() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("count"))
	return rv
}
// The pointer functions the map table uses to manage keys.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable/keyPointerFunctions
func (m NSMapTable) KeyPointerFunctions() INSPointerFunctions {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("keyPointerFunctions"))
	return NSPointerFunctionsFromID(objc.ID(rv))
}
// The pointer functions the map table uses to manage values.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTable/valuePointerFunctions
func (m NSMapTable) ValuePointerFunctions() INSPointerFunctions {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("valuePointerFunctions"))
	return NSPointerFunctionsFromID(objc.ID(rv))
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

