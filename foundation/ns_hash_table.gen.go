// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSHashTable] class.
var (
	_NSHashTableClass     NSHashTableClass
	_NSHashTableClassOnce sync.Once
)

func getNSHashTableClass() NSHashTableClass {
	_NSHashTableClassOnce.Do(func() {
		_NSHashTableClass = NSHashTableClass{class: objc.GetClass("NSHashTable")}
	})
	return _NSHashTableClass
}

// GetNSHashTableClass returns the class object for NSHashTable.
func GetNSHashTableClass() NSHashTableClass {
	return getNSHashTableClass()
}

type NSHashTableClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSHashTableClass) Alloc() NSHashTable {
	rv := objc.Send[NSHashTable](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A collection similar to a set, but with broader range of available memory
// semantics.
//
// # Overview
// 
// The hash table is modeled after [NSSet] with the following differences:
// 
// - It can hold weak references to its members. - Its members may be copied
// on input or may use pointer identity for equality and hashing. - It can
// contain arbitrary pointers (its members are not constrained to being
// objects).
// 
// You can configure an [NSHashTable] instance to operate on arbitrary
// pointers and not just objects, although typically you are encouraged to use
// the C function API for void * pointers. The object-based API (such as
// [NSHashTable.AddObject]) will not work for non-object pointers without type-casting.
// 
// Because of its options, [NSHashTable] is not a set because it can behave
// differently (for example, if pointer equality is specified two `` strings
// will both be entered).
// 
// When configuring hash tables, note that only the options listed in
// [NSHashTableOptions] guarantee that the rest of the API will work
// correctly—including copying, archiving, and fast enumeration. While other
// [NSPointerFunctions] options are used for certain configurations, such as
// to hold arbitrary pointers, not all combinations of the options are valid.
// With some combinations the hash table may not work correctly, or may not
// even be initialized correctly.
// 
// # Subclassing Notes
// 
// [NSHashTable] is not suitable for subclassing.
//
// # Initialization
//
//   - [NSHashTable.InitWithOptionsCapacity]: Returns a hash table initialized with the given attributes.
//   - [NSHashTable.InitWithPointerFunctionsCapacity]: Returns a hash table initialized with the given functions and capacity.
//
// # Accessing Content
//
//   - [NSHashTable.AnyObject]: One of the objects in the hash table.
//   - [NSHashTable.AllObjects]: The hash table’s members.
//   - [NSHashTable.SetRepresentation]: A set that contains the hash table’s members.
//   - [NSHashTable.Count]: The number of elements in the hash table.
//   - [NSHashTable.ContainsObject]: Returns a Boolean value that indicates whether the hash table contains a given object.
//   - [NSHashTable.Member]: Determines whether the hash table contains a given object, and returns that object if it is present
//   - [NSHashTable.ObjectEnumerator]: Returns an enumerator object that lets you access each object in the hash table.
//
// # Manipulating Membership
//
//   - [NSHashTable.AddObject]: Adds a given object to the hash table.
//   - [NSHashTable.RemoveObject]: Removes a given object from the hash table.
//   - [NSHashTable.RemoveAllObjects]: Removes all objects from the hash table.
//
// # Comparing Hash Tables
//
//   - [NSHashTable.IntersectHashTable]: Removes from the receiving hash table each element that isn’t a member of another given hash table.
//   - [NSHashTable.IntersectsHashTable]: Returns a Boolean value that indicates whether a given hash table intersects with the receiving hash table.
//   - [NSHashTable.IsSubsetOfHashTable]: Returns a Boolean value that indicates whether every element in the receiving hash table is also present in another given hash table.
//   - [NSHashTable.IsEqualToHashTable]: Returns a Boolean value that indicates whether a given hash table is equal to the receiving hash table.
//
// # Set Functions
//
//   - [NSHashTable.MinusHashTable]: Removes each element in another given hash table from the receiving hash table, if present.
//   - [NSHashTable.UnionHashTable]: Adds each element in another given hash table to the receiving hash table, if not present.
//
// # Accessing Pointer Functions
//
//   - [NSHashTable.PointerFunctions]: The pointer functions for the hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable
type NSHashTable struct {
	objectivec.Object
}

// NSHashTableFromID constructs a [NSHashTable] from an objc.ID.
//
// A collection similar to a set, but with broader range of available memory
// semantics.
func NSHashTableFromID(id objc.ID) NSHashTable {
	return NSHashTable{objectivec.Object{ID: id}}
}
// NOTE: NSHashTable adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSHashTable] class.
//
// # Initialization
//
//   - [INSHashTable.InitWithOptionsCapacity]: Returns a hash table initialized with the given attributes.
//   - [INSHashTable.InitWithPointerFunctionsCapacity]: Returns a hash table initialized with the given functions and capacity.
//
// # Accessing Content
//
//   - [INSHashTable.AnyObject]: One of the objects in the hash table.
//   - [INSHashTable.AllObjects]: The hash table’s members.
//   - [INSHashTable.SetRepresentation]: A set that contains the hash table’s members.
//   - [INSHashTable.Count]: The number of elements in the hash table.
//   - [INSHashTable.ContainsObject]: Returns a Boolean value that indicates whether the hash table contains a given object.
//   - [INSHashTable.Member]: Determines whether the hash table contains a given object, and returns that object if it is present
//   - [INSHashTable.ObjectEnumerator]: Returns an enumerator object that lets you access each object in the hash table.
//
// # Manipulating Membership
//
//   - [INSHashTable.AddObject]: Adds a given object to the hash table.
//   - [INSHashTable.RemoveObject]: Removes a given object from the hash table.
//   - [INSHashTable.RemoveAllObjects]: Removes all objects from the hash table.
//
// # Comparing Hash Tables
//
//   - [INSHashTable.IntersectHashTable]: Removes from the receiving hash table each element that isn’t a member of another given hash table.
//   - [INSHashTable.IntersectsHashTable]: Returns a Boolean value that indicates whether a given hash table intersects with the receiving hash table.
//   - [INSHashTable.IsSubsetOfHashTable]: Returns a Boolean value that indicates whether every element in the receiving hash table is also present in another given hash table.
//   - [INSHashTable.IsEqualToHashTable]: Returns a Boolean value that indicates whether a given hash table is equal to the receiving hash table.
//
// # Set Functions
//
//   - [INSHashTable.MinusHashTable]: Removes each element in another given hash table from the receiving hash table, if present.
//   - [INSHashTable.UnionHashTable]: Adds each element in another given hash table to the receiving hash table, if not present.
//
// # Accessing Pointer Functions
//
//   - [INSHashTable.PointerFunctions]: The pointer functions for the hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable
type INSHashTable interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Initialization

	// Returns a hash table initialized with the given attributes.
	InitWithOptionsCapacity(options NSPointerFunctionsOptions, initialCapacity uint) NSHashTable
	// Returns a hash table initialized with the given functions and capacity.
	InitWithPointerFunctionsCapacity(functions INSPointerFunctions, initialCapacity uint) NSHashTable

	// Topic: Accessing Content

	// One of the objects in the hash table.
	AnyObject() objectivec.IObject
	// The hash table’s members.
	AllObjects() []objectivec.IObject
	// A set that contains the hash table’s members.
	SetRepresentation() INSSet
	// The number of elements in the hash table.
	Count() uint
	// Returns a Boolean value that indicates whether the hash table contains a given object.
	ContainsObject(anObject objectivec.IObject) bool
	// Determines whether the hash table contains a given object, and returns that object if it is present
	Member(object objectivec.IObject) objectivec.IObject
	// Returns an enumerator object that lets you access each object in the hash table.
	ObjectEnumerator() INSEnumerator

	// Topic: Manipulating Membership

	// Adds a given object to the hash table.
	AddObject(object objectivec.IObject)
	// Removes a given object from the hash table.
	RemoveObject(object objectivec.IObject)
	// Removes all objects from the hash table.
	RemoveAllObjects()

	// Topic: Comparing Hash Tables

	// Removes from the receiving hash table each element that isn’t a member of another given hash table.
	IntersectHashTable(other INSHashTable)
	// Returns a Boolean value that indicates whether a given hash table intersects with the receiving hash table.
	IntersectsHashTable(other INSHashTable) bool
	// Returns a Boolean value that indicates whether every element in the receiving hash table is also present in another given hash table.
	IsSubsetOfHashTable(other INSHashTable) bool
	// Returns a Boolean value that indicates whether a given hash table is equal to the receiving hash table.
	IsEqualToHashTable(other INSHashTable) bool

	// Topic: Set Functions

	// Removes each element in another given hash table from the receiving hash table, if present.
	MinusHashTable(other INSHashTable)
	// Adds each element in another given hash table to the receiving hash table, if not present.
	UnionHashTable(other INSHashTable)

	// Topic: Accessing Pointer Functions

	// The pointer functions for the hash table.
	PointerFunctions() INSPointerFunctions
}

// Init initializes the instance.
func (h NSHashTable) Init() NSHashTable {
	rv := objc.Send[NSHashTable](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h NSHashTable) Autorelease() NSHashTable {
	rv := objc.Send[NSHashTable](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSHashTable creates a new NSHashTable instance.
func NewNSHashTable() NSHashTable {
	class := getNSHashTableClass()
	rv := objc.Send[NSHashTable](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewHashTableWithCoder(coder INSCoder) NSHashTable {
	instance := getNSHashTableClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSHashTableFromID(rv)
}

// Returns a hash table initialized with the given attributes.
//
// options: A bit field that specifies the options for the elements in the hash table.
// For possible values, see [NSHashTableOptions].
//
// initialCapacity: The initial number of elements the hash table can hold.
//
// # Return Value
// 
// A hash table initialized with options specified by `options` and initial
// capacity of `capacity`.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/init(options:capacity:)
func NewHashTableWithOptionsCapacity(options NSPointerFunctionsOptions, initialCapacity uint) NSHashTable {
	instance := getNSHashTableClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOptions:capacity:"), options, initialCapacity)
	return NSHashTableFromID(rv)
}

// Returns a hash table initialized with the given functions and capacity.
//
// functions: The pointer functions for the new hash table.
//
// initialCapacity: The initial capacity of the hash table.
//
// # Return Value
// 
// A hash table initialized with the given functions and capacity.
//
// # Discussion
// 
// Hash tables allocate additional memory as needed, so `initialCapacity`
// simply establishes the object’s initial capacity.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/init(pointerFunctions:capacity:)
func NewHashTableWithPointerFunctionsCapacity(functions INSPointerFunctions, initialCapacity uint) NSHashTable {
	instance := getNSHashTableClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPointerFunctions:capacity:"), functions, initialCapacity)
	return NSHashTableFromID(rv)
}

// Returns a hash table initialized with the given attributes.
//
// options: A bit field that specifies the options for the elements in the hash table.
// For possible values, see [NSHashTableOptions].
//
// initialCapacity: The initial number of elements the hash table can hold.
//
// # Return Value
// 
// A hash table initialized with options specified by `options` and initial
// capacity of `capacity`.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/init(options:capacity:)
func (h NSHashTable) InitWithOptionsCapacity(options NSPointerFunctionsOptions, initialCapacity uint) NSHashTable {
	rv := objc.Send[NSHashTable](h.ID, objc.Sel("initWithOptions:capacity:"), options, initialCapacity)
	return rv
}

// Returns a hash table initialized with the given functions and capacity.
//
// functions: The pointer functions for the new hash table.
//
// initialCapacity: The initial capacity of the hash table.
//
// # Return Value
// 
// A hash table initialized with the given functions and capacity.
//
// # Discussion
// 
// Hash tables allocate additional memory as needed, so `initialCapacity`
// simply establishes the object’s initial capacity.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/init(pointerFunctions:capacity:)
func (h NSHashTable) InitWithPointerFunctionsCapacity(functions INSPointerFunctions, initialCapacity uint) NSHashTable {
	rv := objc.Send[NSHashTable](h.ID, objc.Sel("initWithPointerFunctions:capacity:"), functions, initialCapacity)
	return rv
}

// Returns a Boolean value that indicates whether the hash table contains a
// given object.
//
// anObject: The object to test for membership in the hash table.
//
// # Return Value
// 
// [true] if the hash table contains `anObject`, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The equality test used depends on the personality option selected. For
// instance, choosing the [PointerFunctionsObjectPersonality] option will use
// `` to determine equality. See [NSPointerFunctions.Options] for more
// information on personality options and their corresponding equality tests.
//
// [NSPointerFunctions.Options]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/contains(_:)
func (h NSHashTable) ContainsObject(anObject objectivec.IObject) bool {
	rv := objc.Send[bool](h.ID, objc.Sel("containsObject:"), anObject)
	return rv
}

// Determines whether the hash table contains a given object, and returns that
// object if it is present
//
// object: The object to test for membership in the hash table.
//
// # Return Value
// 
// If `object` is a member of the hash table, returns `object`, otherwise
// returns `nil`.
//
// # Discussion
// 
// The equality test used depends on the personality option selected. For
// instance, choosing the [PointerFunctionsObjectPersonality] option will use
// `` to determine equality. See [NSPointerFunctions.Options] for more
// information on personality options and their corresponding equality tests.
//
// [NSPointerFunctions.Options]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/member(_:)
func (h NSHashTable) Member(object objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("member:"), object)
	return objectivec.Object{ID: rv}
}

// Returns an enumerator object that lets you access each object in the hash
// table.
//
// # Return Value
// 
// An enumerator object that lets you access each object in the hash table.
//
// # Discussion
// 
// The following code fragment illustrates how you can use this method.
// 
// # Special Considerations
// 
// It is more efficient to use the fast enumeration protocol (see
// [NSFastEnumeration]).
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/objectEnumerator()
func (h NSHashTable) ObjectEnumerator() INSEnumerator {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("objectEnumerator"))
	return NSEnumeratorFromID(rv)
}

// Adds a given object to the hash table.
//
// object: The object to add to the hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/add(_:)
func (h NSHashTable) AddObject(object objectivec.IObject) {
	objc.Send[objc.ID](h.ID, objc.Sel("addObject:"), object)
}

// Removes a given object from the hash table.
//
// object: The object to remove from the hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/remove(_:)
func (h NSHashTable) RemoveObject(object objectivec.IObject) {
	objc.Send[objc.ID](h.ID, objc.Sel("removeObject:"), object)
}

// Removes all objects from the hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/removeAllObjects()
func (h NSHashTable) RemoveAllObjects() {
	objc.Send[objc.ID](h.ID, objc.Sel("removeAllObjects"))
}

// Removes from the receiving hash table each element that isn’t a member of
// another given hash table.
//
// other: The hash table with which to perform the intersection.
//
// # Discussion
// 
// The equality test used for members depends on the personality option
// selected. For instance, choosing the [PointerFunctionsObjectPersonality]
// option will use `` to determine equality. See [NSPointerFunctions.Options]
// for more information on personality options and their corresponding
// equality tests.
//
// [NSPointerFunctions.Options]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/intersect(_:)
func (h NSHashTable) IntersectHashTable(other INSHashTable) {
	objc.Send[objc.ID](h.ID, objc.Sel("intersectHashTable:"), other)
}

// Returns a Boolean value that indicates whether a given hash table
// intersects with the receiving hash table.
//
// other: The hash table with which to compare the receiving hash table.
//
// # Return Value
// 
// [true] if `other` intersects with the receiving hash table, otherwise
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The equality test used for members depends on the personality option
// selected. For instance, choosing the [PointerFunctionsObjectPersonality]
// option will use `` to determine equality. See [NSPointerFunctions.Options]
// for more information on personality options and their corresponding
// equality tests.
//
// [NSPointerFunctions.Options]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/intersects(_:)
func (h NSHashTable) IntersectsHashTable(other INSHashTable) bool {
	rv := objc.Send[bool](h.ID, objc.Sel("intersectsHashTable:"), other)
	return rv
}

// Returns a Boolean value that indicates whether every element in the
// receiving hash table is also present in another given hash table.
//
// other: The hash table with which to compare the receiving hash table.
//
// # Return Value
// 
// [true] if every element in the receiving hash table is also present in
// `other`, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The equality test used for members depends on the personality option
// selected. For instance, choosing the [PointerFunctionsObjectPersonality]
// option will use `` to determine equality. See [NSPointerFunctions.Options]
// for more information on personality options and their corresponding
// equality tests.
//
// [NSPointerFunctions.Options]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/isSubset(of:)
func (h NSHashTable) IsSubsetOfHashTable(other INSHashTable) bool {
	rv := objc.Send[bool](h.ID, objc.Sel("isSubsetOfHashTable:"), other)
	return rv
}

// Returns a Boolean value that indicates whether a given hash table is equal
// to the receiving hash table.
//
// other: The hash table with which to compare the receiving hash table.
//
// # Return Value
// 
// [true] if the contents of `other` are equal to the contents of the
// receiving hash table, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Two hash tables have equal contents if they each have the same number of
// members and if each member of one hash table is present in the other.
// 
// The equality test used for members depends on the personality option
// selected. For instance, choosing the [PointerFunctionsObjectPersonality]
// option will use `` to determine equality. See [NSPointerFunctions.Options]
// for more information on personality options and their corresponding
// equality tests.
//
// [NSPointerFunctions.Options]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/isEqual(to:)
func (h NSHashTable) IsEqualToHashTable(other INSHashTable) bool {
	rv := objc.Send[bool](h.ID, objc.Sel("isEqualToHashTable:"), other)
	return rv
}

// Removes each element in another given hash table from the receiving hash
// table, if present.
//
// other: The hash table of elements to remove from the receiving hash table.
//
// # Discussion
// 
// The equality test used for members depends on the personality option
// selected. For instance, choosing the [PointerFunctionsObjectPersonality]
// option will use `` to determine equality. See [NSPointerFunctions.Options]
// for more information on personality options and their corresponding
// equality tests.
//
// [NSPointerFunctions.Options]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/minus(_:)
func (h NSHashTable) MinusHashTable(other INSHashTable) {
	objc.Send[objc.ID](h.ID, objc.Sel("minusHashTable:"), other)
}

// Adds each element in another given hash table to the receiving hash table,
// if not present.
//
// other: The hash table of elements to add to the receiving hash table.
//
// # Discussion
// 
// The equality test used for members depends on the personality option
// selected. For instance, choosing the [PointerFunctionsObjectPersonality]
// option will use `` to determine equality. See [NSPointerFunctions.Options]
// for more information on personality options and their corresponding
// equality tests.
//
// [NSPointerFunctions.Options]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/union(_:)
func (h NSHashTable) UnionHashTable(other INSHashTable) {
	objc.Send[objc.ID](h.ID, objc.Sel("unionHashTable:"), other)
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
func (h NSHashTable) CountByEnumeratingWithStateObjectsCount(state NSFastEnumerationState, buffer []objectivec.IObject, len_ uint) uint {
	rv := objc.Send[uint](h.ID, objc.Sel("countByEnumeratingWithState:objects:count:"), state, objc.CArray(buffer), len_)
	return rv
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (h NSHashTable) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](h.ID, objc.Sel("encodeWithCoder:"), coder)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (h NSHashTable) InitWithCoder(coder INSCoder) NSHashTable {
	rv := objc.Send[NSHashTable](h.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Returns a new hash table for storing weak references to its contents.
//
// # Return Value
// 
// A new hash table that uses the [PointerFunctionsWeakMemory] options and
// [PointerFunctionsObjectPersonality] and has an initial capacity of `0`.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/weakObjects()
func (_NSHashTableClass NSHashTableClass) WeakObjectsHashTable() NSHashTable {
	rv := objc.Send[objc.ID](objc.ID(_NSHashTableClass.class), objc.Sel("weakObjectsHashTable"))
	return NSHashTableFromID(rv)
}

// Returns a hash table with given pointer functions options.
//
// options: A bit field that specifies the options for the elements in the hash table.
// For possible values, see [NSHashTableOptions].
//
// # Return Value
// 
// A hash table with given pointer functions options.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/init(options:)
func (_NSHashTableClass NSHashTableClass) HashTableWithOptions(options NSPointerFunctionsOptions) NSHashTable {
	rv := objc.Send[objc.ID](objc.ID(_NSHashTableClass.class), objc.Sel("hashTableWithOptions:"), options)
	return NSHashTableFromID(rv)
}

// One of the objects in the hash table.
//
// # Discussion
// 
// One of the objects in the hash table, or `nil` if the hash table contains
// no objects.
// 
// The object returned is chosen at the hash table’s convenience—the
// selection is not guaranteed to be random.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/anyObject
func (h NSHashTable) AnyObject() objectivec.IObject {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("anyObject"))
	return objectivec.Object{ID: rv}
}

// The hash table’s members.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/allObjects
func (h NSHashTable) AllObjects() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](h.ID, objc.Sel("allObjects"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// A set that contains the hash table’s members.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/setRepresentation
func (h NSHashTable) SetRepresentation() INSSet {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("setRepresentation"))
	return NSSetFromID(objc.ID(rv))
}

// The number of elements in the hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/count
func (h NSHashTable) Count() uint {
	rv := objc.Send[uint](h.ID, objc.Sel("count"))
	return rv
}

// The pointer functions for the hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTable/pointerFunctions
func (h NSHashTable) PointerFunctions() INSPointerFunctions {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("pointerFunctions"))
	return NSPointerFunctionsFromID(objc.ID(rv))
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

