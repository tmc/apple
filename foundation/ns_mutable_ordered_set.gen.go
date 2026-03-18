// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMutableOrderedSet] class.
var (
	_NSMutableOrderedSetClass     NSMutableOrderedSetClass
	_NSMutableOrderedSetClassOnce sync.Once
)

func getNSMutableOrderedSetClass() NSMutableOrderedSetClass {
	_NSMutableOrderedSetClassOnce.Do(func() {
		_NSMutableOrderedSetClass = NSMutableOrderedSetClass{class: objc.GetClass("NSMutableOrderedSet")}
	})
	return _NSMutableOrderedSetClass
}

// GetNSMutableOrderedSetClass returns the class object for NSMutableOrderedSet.
func GetNSMutableOrderedSetClass() NSMutableOrderedSetClass {
	return getNSMutableOrderedSetClass()
}

type NSMutableOrderedSetClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMutableOrderedSetClass) Alloc() NSMutableOrderedSet {
	rv := objc.Send[NSMutableOrderedSet](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A dynamic, ordered collection of unique objects.
//
// # Overview
// 
// [NSMutableOrderedSet] objects are not like C arrays. That is, even though
// you may specify a size when you create a mutable ordered set, the specified
// size is regarded as a “hint”; the actual size of the set is still 0.
// This means that you cannot insert an object at an index greater than the
// current count of an set. For example, if a set contains two objects, its
// size is 2, so you can add objects at indices 0, 1, or 2. Index 3 is illegal
// and out of bounds; if you try to add an object at index 3 (when the size of
// the array is 2), [NSMutableOrderedSet] raises an exception.
//
// # Creating a Mutable Ordered Set
//
//   - [NSMutableOrderedSet.InitWithCapacity]: Returns an initialized mutable ordered set with a given initial capacity.
//
// # Adding, Removing, and Reordering Entries
//
//   - [NSMutableOrderedSet.AddObject]: Appends a given object to the end of the mutable ordered set, if it is not already a member.
//   - [NSMutableOrderedSet.AddObjectsCount]: Appends the given number of objects from a given C array to the end of the mutable ordered set.
//   - [NSMutableOrderedSet.AddObjectsFromArray]: Appends to the end of the mutable ordered set each object contained in a given array that is not already a member.
//   - [NSMutableOrderedSet.InsertObjectAtIndex]: Inserts the given object at the specified index of the mutable ordered set, if it is not already a member.
//   - [NSMutableOrderedSet.InsertObjectsAtIndexes]: Inserts the objects in the array at the specified indexes.
//   - [NSMutableOrderedSet.RemoveObject]: Removes a given object from the mutable ordered set.
//   - [NSMutableOrderedSet.RemoveObjectAtIndex]: Removes a the object at the specified index from the mutable ordered set.
//   - [NSMutableOrderedSet.RemoveObjectsAtIndexes]: Removes the objects at the specified indexes from the mutable ordered set.
//   - [NSMutableOrderedSet.RemoveObjectsInArray]: Removes the objects in the array from the mutable ordered set.
//   - [NSMutableOrderedSet.RemoveObjectsInRange]: Removes from the mutable ordered set each of the objects within a given range.
//   - [NSMutableOrderedSet.RemoveAllObjects]: Removes all the objects from the mutable ordered set.
//   - [NSMutableOrderedSet.ReplaceObjectAtIndexWithObject]: Replaces the object at the specified index with the new object.
//   - [NSMutableOrderedSet.ReplaceObjectsAtIndexesWithObjects]: Replaces the objects at the specified indexes with the new objects.
//   - [NSMutableOrderedSet.ReplaceObjectsInRangeWithObjectsCount]: Replaces the objects in the receiving mutable ordered set at the range with the specified number of objects from a given C array.
//   - [NSMutableOrderedSet.SetObjectAtIndex]: Appends or replaces the object at the specified index.
//   - [NSMutableOrderedSet.MoveObjectsAtIndexesToIndex]: Moves the objects at the specified indexes to the new location.
//   - [NSMutableOrderedSet.ExchangeObjectAtIndexWithObjectAtIndex]: Exchanges the object at the specified index with the object at the other index.
//   - [NSMutableOrderedSet.FilterUsingPredicate]: Evaluates a given predicate against the mutable ordered set’s content and leaves only objects that match.
//
// # Sorting Entries
//
//   - [NSMutableOrderedSet.SortUsingDescriptors]: Sorts the receiving ordered set using a given array of sort descriptors.
//   - [NSMutableOrderedSet.SortUsingComparator]: Sorts the mutable ordered set using the comparison method specified by the comparator block.
//   - [NSMutableOrderedSet.SortWithOptionsUsingComparator]: Sorts the mutable ordered set using the specified options and the comparison method specified by a given comparator block.
//   - [NSMutableOrderedSet.SortRangeOptionsUsingComparator]: Sorts the specified range of the mutable ordered set using the specified options and the comparison method specified by a given comparator block.
//
// # Combining and Recombining Entries
//
//   - [NSMutableOrderedSet.IntersectOrderedSet]: Removes from the receiving ordered set each object that isn’t a member of another ordered set.
//   - [NSMutableOrderedSet.IntersectSet]: Removes from the receiving ordered set each object that isn’t a member of another set.
//   - [NSMutableOrderedSet.MinusOrderedSet]: Removes each object in another given ordered set from the receiving mutable ordered set, if present.
//   - [NSMutableOrderedSet.MinusSet]: Removes each object in another given set from the receiving mutable ordered set, if present.
//   - [NSMutableOrderedSet.UnionOrderedSet]: Adds each object in another given ordered set to the receiving mutable ordered set, if not present.
//   - [NSMutableOrderedSet.UnionSet]: Adds each object in another given set to the receiving mutable ordered set, if not present.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet
type NSMutableOrderedSet struct {
	NSOrderedSet
}

// NSMutableOrderedSetFromID constructs a [NSMutableOrderedSet] from an objc.ID.
//
// A dynamic, ordered collection of unique objects.
func NSMutableOrderedSetFromID(id objc.ID) NSMutableOrderedSet {
	return NSMutableOrderedSet{NSOrderedSet: NSOrderedSetFromID(id)}
}
// NOTE: NSMutableOrderedSet adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMutableOrderedSet] class.
//
// # Creating a Mutable Ordered Set
//
//   - [INSMutableOrderedSet.InitWithCapacity]: Returns an initialized mutable ordered set with a given initial capacity.
//
// # Adding, Removing, and Reordering Entries
//
//   - [INSMutableOrderedSet.AddObject]: Appends a given object to the end of the mutable ordered set, if it is not already a member.
//   - [INSMutableOrderedSet.AddObjectsCount]: Appends the given number of objects from a given C array to the end of the mutable ordered set.
//   - [INSMutableOrderedSet.AddObjectsFromArray]: Appends to the end of the mutable ordered set each object contained in a given array that is not already a member.
//   - [INSMutableOrderedSet.InsertObjectAtIndex]: Inserts the given object at the specified index of the mutable ordered set, if it is not already a member.
//   - [INSMutableOrderedSet.InsertObjectsAtIndexes]: Inserts the objects in the array at the specified indexes.
//   - [INSMutableOrderedSet.RemoveObject]: Removes a given object from the mutable ordered set.
//   - [INSMutableOrderedSet.RemoveObjectAtIndex]: Removes a the object at the specified index from the mutable ordered set.
//   - [INSMutableOrderedSet.RemoveObjectsAtIndexes]: Removes the objects at the specified indexes from the mutable ordered set.
//   - [INSMutableOrderedSet.RemoveObjectsInArray]: Removes the objects in the array from the mutable ordered set.
//   - [INSMutableOrderedSet.RemoveObjectsInRange]: Removes from the mutable ordered set each of the objects within a given range.
//   - [INSMutableOrderedSet.RemoveAllObjects]: Removes all the objects from the mutable ordered set.
//   - [INSMutableOrderedSet.ReplaceObjectAtIndexWithObject]: Replaces the object at the specified index with the new object.
//   - [INSMutableOrderedSet.ReplaceObjectsAtIndexesWithObjects]: Replaces the objects at the specified indexes with the new objects.
//   - [INSMutableOrderedSet.ReplaceObjectsInRangeWithObjectsCount]: Replaces the objects in the receiving mutable ordered set at the range with the specified number of objects from a given C array.
//   - [INSMutableOrderedSet.SetObjectAtIndex]: Appends or replaces the object at the specified index.
//   - [INSMutableOrderedSet.MoveObjectsAtIndexesToIndex]: Moves the objects at the specified indexes to the new location.
//   - [INSMutableOrderedSet.ExchangeObjectAtIndexWithObjectAtIndex]: Exchanges the object at the specified index with the object at the other index.
//   - [INSMutableOrderedSet.FilterUsingPredicate]: Evaluates a given predicate against the mutable ordered set’s content and leaves only objects that match.
//
// # Sorting Entries
//
//   - [INSMutableOrderedSet.SortUsingDescriptors]: Sorts the receiving ordered set using a given array of sort descriptors.
//   - [INSMutableOrderedSet.SortUsingComparator]: Sorts the mutable ordered set using the comparison method specified by the comparator block.
//   - [INSMutableOrderedSet.SortWithOptionsUsingComparator]: Sorts the mutable ordered set using the specified options and the comparison method specified by a given comparator block.
//   - [INSMutableOrderedSet.SortRangeOptionsUsingComparator]: Sorts the specified range of the mutable ordered set using the specified options and the comparison method specified by a given comparator block.
//
// # Combining and Recombining Entries
//
//   - [INSMutableOrderedSet.IntersectOrderedSet]: Removes from the receiving ordered set each object that isn’t a member of another ordered set.
//   - [INSMutableOrderedSet.IntersectSet]: Removes from the receiving ordered set each object that isn’t a member of another set.
//   - [INSMutableOrderedSet.MinusOrderedSet]: Removes each object in another given ordered set from the receiving mutable ordered set, if present.
//   - [INSMutableOrderedSet.MinusSet]: Removes each object in another given set from the receiving mutable ordered set, if present.
//   - [INSMutableOrderedSet.UnionOrderedSet]: Adds each object in another given ordered set to the receiving mutable ordered set, if not present.
//   - [INSMutableOrderedSet.UnionSet]: Adds each object in another given set to the receiving mutable ordered set, if not present.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet
type INSMutableOrderedSet interface {
	INSOrderedSet

	// Topic: Creating a Mutable Ordered Set

	// Returns an initialized mutable ordered set with a given initial capacity.
	InitWithCapacity(numItems uint) NSMutableOrderedSet

	// Topic: Adding, Removing, and Reordering Entries

	// Appends a given object to the end of the mutable ordered set, if it is not already a member.
	AddObject(object objectivec.IObject)
	// Appends the given number of objects from a given C array to the end of the mutable ordered set.
	AddObjectsCount(objects []objectivec.IObject, count uint)
	// Appends to the end of the mutable ordered set each object contained in a given array that is not already a member.
	AddObjectsFromArray(array []objectivec.IObject)
	// Inserts the given object at the specified index of the mutable ordered set, if it is not already a member.
	InsertObjectAtIndex(object objectivec.IObject, idx uint)
	// Inserts the objects in the array at the specified indexes.
	InsertObjectsAtIndexes(objects []objectivec.IObject, indexes INSIndexSet)
	// Removes a given object from the mutable ordered set.
	RemoveObject(object objectivec.IObject)
	// Removes a the object at the specified index from the mutable ordered set.
	RemoveObjectAtIndex(idx uint)
	// Removes the objects at the specified indexes from the mutable ordered set.
	RemoveObjectsAtIndexes(indexes INSIndexSet)
	// Removes the objects in the array from the mutable ordered set.
	RemoveObjectsInArray(array []objectivec.IObject)
	// Removes from the mutable ordered set each of the objects within a given range.
	RemoveObjectsInRange(range_ NSRange)
	// Removes all the objects from the mutable ordered set.
	RemoveAllObjects()
	// Replaces the object at the specified index with the new object.
	ReplaceObjectAtIndexWithObject(idx uint, object objectivec.IObject)
	// Replaces the objects at the specified indexes with the new objects.
	ReplaceObjectsAtIndexesWithObjects(indexes INSIndexSet, objects []objectivec.IObject)
	// Replaces the objects in the receiving mutable ordered set at the range with the specified number of objects from a given C array.
	ReplaceObjectsInRangeWithObjectsCount(range_ NSRange, objects []objectivec.IObject, count uint)
	// Appends or replaces the object at the specified index.
	SetObjectAtIndex(obj objectivec.IObject, idx uint)
	// Moves the objects at the specified indexes to the new location.
	MoveObjectsAtIndexesToIndex(indexes INSIndexSet, idx uint)
	// Exchanges the object at the specified index with the object at the other index.
	ExchangeObjectAtIndexWithObjectAtIndex(idx1 uint, idx2 uint)
	// Evaluates a given predicate against the mutable ordered set’s content and leaves only objects that match.
	FilterUsingPredicate(p INSPredicate)

	// Topic: Sorting Entries

	// Sorts the receiving ordered set using a given array of sort descriptors.
	SortUsingDescriptors(sortDescriptors []NSSortDescriptor)
	// Sorts the mutable ordered set using the comparison method specified by the comparator block.
	SortUsingComparator(cmptr NSComparator)
	// Sorts the mutable ordered set using the specified options and the comparison method specified by a given comparator block.
	SortWithOptionsUsingComparator(opts NSSortOptions, cmptr NSComparator)
	// Sorts the specified range of the mutable ordered set using the specified options and the comparison method specified by a given comparator block.
	SortRangeOptionsUsingComparator(range_ NSRange, opts NSSortOptions, cmptr NSComparator)

	// Topic: Combining and Recombining Entries

	// Removes from the receiving ordered set each object that isn’t a member of another ordered set.
	IntersectOrderedSet(other INSOrderedSet)
	// Removes from the receiving ordered set each object that isn’t a member of another set.
	IntersectSet(other INSSet)
	// Removes each object in another given ordered set from the receiving mutable ordered set, if present.
	MinusOrderedSet(other INSOrderedSet)
	// Removes each object in another given set from the receiving mutable ordered set, if present.
	MinusSet(other INSSet)
	// Adds each object in another given ordered set to the receiving mutable ordered set, if not present.
	UnionOrderedSet(other INSOrderedSet)
	// Adds each object in another given set to the receiving mutable ordered set, if not present.
	UnionSet(other INSSet)

	ApplyDifference(difference INSOrderedCollectionDifference)
	// Replaces the given object at the specified index of the mutable ordered set.
	SetObjectAtIndexedSubscript(obj objectivec.IObject, idx uint)
}

// Init initializes the instance.
func (m NSMutableOrderedSet) Init() NSMutableOrderedSet {
	rv := objc.Send[NSMutableOrderedSet](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMutableOrderedSet) Autorelease() NSMutableOrderedSet {
	rv := objc.Send[NSMutableOrderedSet](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMutableOrderedSet creates a new NSMutableOrderedSet instance.
func NewNSMutableOrderedSet() NSMutableOrderedSet {
	class := getNSMutableOrderedSetClass()
	rv := objc.Send[NSMutableOrderedSet](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a newly allocated set with the objects that are contained in a
// given array.
//
// array: An array of objects to add to the new set.
// 
// If the same object appears more than once in array, it is represented only
// once in the returned ordered set.
//
// # Return Value
// 
// An initialized ordered set with the contents of array. The returned ordered
// set might be different than the original receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(array:)
func NewMutableOrderedSetWithArray(array []objectivec.IObject) NSMutableOrderedSet {
	instance := getNSMutableOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArray:"), objectivec.IObjectSliceToNSArray(array))
	return NSMutableOrderedSetFromID(rv)
}

// Initializes a newly allocated set with the objects that are contained in a
// given array, optionally copying the items.
//
// set: An array of objects to add to the new set.
// 
// If the same object appears more than once in array, it is represented only
// once in the returned ordered set.
//
// flag: If [true] the objects are copied to the ordered set; otherwise [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An initialized ordered set containing a uniqued collection of the objects
// contained in the array.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(array:copyItems:)
func NewMutableOrderedSetWithArrayCopyItems(set []objectivec.IObject, flag bool) NSMutableOrderedSet {
	instance := getNSMutableOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArray:copyItems:"), objectivec.IObjectSliceToNSArray(set), flag)
	return NSMutableOrderedSetFromID(rv)
}

// Initializes a newly allocated set with the objects that are contained in
// the specified range of an array, optionally copying the items.
//
// set: An array of objects to add to the new set.
// 
// If the same object appears more than once in array, it is represented only
// once in the returned ordered set.
//
// range: The range of objects in `array` to add to the ordered set.
//
// flag: If [true] the objects are copied to the ordered set; otherwise [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An initialized ordered set containing a uniqued collection of the objects
// contained in specified range of the array.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(array:range:copyItems:)
func NewMutableOrderedSetWithArrayRangeCopyItems(set []objectivec.IObject, range_ NSRange, flag bool) NSMutableOrderedSet {
	instance := getNSMutableOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArray:range:copyItems:"), objectivec.IObjectSliceToNSArray(set), range_, flag)
	return NSMutableOrderedSetFromID(rv)
}

// Returns an initialized mutable ordered set with a given initial capacity.
//
// numItems: The initial capacity of the new ordered set.
//
// # Return Value
// 
// An initialized mutable ordered set with initial capacity to hold `numItems`
// members.
//
// # Discussion
// 
// Mutable ordered sets allocate additional memory as needed, so `numItems`
// simply establishes the set’s initial capacity.
// 
// This method is a designated initializer of [NSMutableOrderedSet].
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/init(capacity:)
func NewMutableOrderedSetWithCapacity(numItems uint) NSMutableOrderedSet {
	instance := getNSMutableOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCapacity:"), numItems)
	return NSMutableOrderedSetFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/init(coder:)
func NewMutableOrderedSetWithCoder(coder INSCoder) NSMutableOrderedSet {
	instance := getNSMutableOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSMutableOrderedSetFromID(rv)
}

// Initializes a new ordered set with the object.
//
// object: The object to add to the new ordered set
//
// # Return Value
// 
// A new ordered set that contains a single member, `object`.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(object:)
func NewMutableOrderedSetWithObject(object objectivec.IObject) NSMutableOrderedSet {
	instance := getNSMutableOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObject:"), object)
	return NSMutableOrderedSetFromID(rv)
}

// Initializes a newly allocated set with members taken from the specified
// list of objects.
//
// firstObj: The first object to add to the new set.
//
// # Return Value
// 
// An initialized ordered set containing the objects specified in the
// parameter list. The returned set might be different than the original
// receiver.
//
// # Discussion
// 
// To add additional objects to the new ordered set, pass comma-separated list
// of trailing variadic arguments, ending with `nil`,
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/initWithObjects:
func NewMutableOrderedSetWithObjects(firstObj objectivec.IObject) NSMutableOrderedSet {
	instance := getNSMutableOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObjects:"), firstObj)
	return NSMutableOrderedSetFromID(rv)
}

// Initializes a new ordered set with the contents of a set.
//
// set: A set.
//
// # Return Value
// 
// An initialized ordered set containing references to the objects in the set.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(orderedSet:)
func NewMutableOrderedSetWithOrderedSet(set INSOrderedSet) NSMutableOrderedSet {
	instance := getNSMutableOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOrderedSet:"), set)
	return NSMutableOrderedSetFromID(rv)
}

// Initializes a new ordered set with the contents of a set, optionally
// copying the items.
//
// set: A set.
//
// flag: If [true] the objects are copied to the ordered set; otherwise [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An initialized ordered set containing the objects in the set.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(orderedSet:copyItems:)
func NewMutableOrderedSetWithOrderedSetCopyItems(set INSOrderedSet, flag bool) NSMutableOrderedSet {
	instance := getNSMutableOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOrderedSet:copyItems:"), set, flag)
	return NSMutableOrderedSetFromID(rv)
}

// Initializes a new ordered set with the contents of an ordered set,
// optionally copying the items.
//
// set: An ordered set.
//
// range: The range of objects in `orderedSet` to add to the ordered set.
//
// flag: If [true] the objects are copied to the ordered set; otherwise [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An initialized ordered set containing the objects in the ordered set.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(orderedSet:range:copyItems:)
func NewMutableOrderedSetWithOrderedSetRangeCopyItems(set INSOrderedSet, range_ NSRange, flag bool) NSMutableOrderedSet {
	instance := getNSMutableOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOrderedSet:range:copyItems:"), set, range_, flag)
	return NSMutableOrderedSetFromID(rv)
}

// Initializes a new ordered set with the contents of a set.
//
// set: The set.
//
// # Return Value
// 
// An initialized ordered set containing the objects in the set.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(set:)
func NewMutableOrderedSetWithSet(set INSSet) NSMutableOrderedSet {
	instance := getNSMutableOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSet:"), set)
	return NSMutableOrderedSetFromID(rv)
}

// Initializes a new ordered set with the contents of a set, optionally
// copying the objects in the set.
//
// set: The set.
//
// flag: If [true] the objects are copied to the ordered set; otherwise [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An initialized ordered set containing the objects in the set.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(set:copyItems:)
func NewMutableOrderedSetWithSetCopyItems(set INSSet, flag bool) NSMutableOrderedSet {
	instance := getNSMutableOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSet:copyItems:"), set, flag)
	return NSMutableOrderedSetFromID(rv)
}

// Returns an initialized mutable ordered set with a given initial capacity.
//
// numItems: The initial capacity of the new ordered set.
//
// # Return Value
// 
// An initialized mutable ordered set with initial capacity to hold `numItems`
// members.
//
// # Discussion
// 
// Mutable ordered sets allocate additional memory as needed, so `numItems`
// simply establishes the set’s initial capacity.
// 
// This method is a designated initializer of [NSMutableOrderedSet].
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/init(capacity:)
func (m NSMutableOrderedSet) InitWithCapacity(numItems uint) NSMutableOrderedSet {
	rv := objc.Send[NSMutableOrderedSet](m.ID, objc.Sel("initWithCapacity:"), numItems)
	return rv
}

// Appends a given object to the end of the mutable ordered set, if it is not
// already a member.
//
// object: The object to add to the set.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/add(_:)
func (m NSMutableOrderedSet) AddObject(object objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("addObject:"), object)
}

// Appends the given number of objects from a given C array to the end of the
// mutable ordered set.
//
// objects: A C array of objects.
//
// count: The number of values from the objects C array to append to the mutable
// ordered set. This number will be the count of the new array—it must not
// be negative or greater than the number of elements in objects.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/add(_:count:)
func (m NSMutableOrderedSet) AddObjectsCount(objects []objectivec.IObject, count uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("addObjects:count:"), objc.CArray(objects), count)
}

// Appends to the end of the mutable ordered set each object contained in a
// given array that is not already a member.
//
// array: An array of objects to add to the set.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/addObjects(from:)
func (m NSMutableOrderedSet) AddObjectsFromArray(array []objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("addObjectsFromArray:"), objectivec.IObjectSliceToNSArray(array))
}

// Inserts the given object at the specified index of the mutable ordered set,
// if it is not already a member.
//
// object: The object to insert into the set’s content.
//
// idx: The index in the mutable ordered set at which to insert `object`. This
// value must not be greater than the count of elements in the array.
//
// # Discussion
// 
// If the object is already a member, this method has no effect. If the
// specified index is already occupied, the objects at that index and beyond
// are shifted by adding `1` to their indexes to make room for the inserted
// object.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/insert(_:at:)-7qg51
func (m NSMutableOrderedSet) InsertObjectAtIndex(object objectivec.IObject, idx uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("insertObject:atIndex:"), object, idx)
}

// Inserts the objects in the array at the specified indexes.
//
// objects: An array of objects to insert into the mutable ordered set.
//
// indexes: The indexes at which the objects in objects should be inserted. The count
// of locations in indexes must equal the count of objects.
//
// # Discussion
// 
// Each object in objects is inserted into the receiving mutable ordered set
// in turn at the corresponding location specified in indexes after earlier
// insertions have been made.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/insert(_:at:)-3ncnm
func (m NSMutableOrderedSet) InsertObjectsAtIndexes(objects []objectivec.IObject, indexes INSIndexSet) {
	objc.Send[objc.ID](m.ID, objc.Sel("insertObjects:atIndexes:"), objectivec.IObjectSliceToNSArray(objects), indexes)
}

// Removes a given object from the mutable ordered set.
//
// object: The object to remove from the mutable ordered set.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/remove(_:)
func (m NSMutableOrderedSet) RemoveObject(object objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeObject:"), object)
}

// Removes a the object at the specified index from the mutable ordered set.
//
// idx: The index of the object to remove from the mutable ordered set. The value
// must not exceed the bounds of the set.
//
// # Discussion
// 
// To fill the gap, all elements beyond index are moved by subtracting 1 from
// their index.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/removeObject(at:)
func (m NSMutableOrderedSet) RemoveObjectAtIndex(idx uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeObjectAtIndex:"), idx)
}

// Removes the objects at the specified indexes from the mutable ordered set.
//
// indexes: The indexes of the objects to remove from the mutable ordered set. The
// locations specified by indexes must lie within the bounds of the mutable
// ordered .
//
// # Discussion
// 
// This method is similar to [RemoveObjectAtIndex], but allows you to
// efficiently remove multiple objects with a single operation.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/removeObjects(at:)
func (m NSMutableOrderedSet) RemoveObjectsAtIndexes(indexes INSIndexSet) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeObjectsAtIndexes:"), indexes)
}

// Removes the objects in the array from the mutable ordered set.
//
// array: An array containing the objects to be removed from the receiving mutable
// ordered set.
//
// # Discussion
// 
// This method is similar to [RemoveObject], but allows you to efficiently
// remove large sets of objects with a single operation. If the receiving
// mutable ordered set does not contain objects in array, the method has no
// effect (although it does incur the overhead of searching the contents).
// 
// This method assumes that all elements in array respond to [Hash] and
// [isEqual(_:)].
//
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/removeObjects(in:)-8h2kh
func (m NSMutableOrderedSet) RemoveObjectsInArray(array []objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeObjectsInArray:"), objectivec.IObjectSliceToNSArray(array))
}

// Removes from the mutable ordered set each of the objects within a given
// range.
//
// range: The range of the objects to remove from the mutable ordered set.
//
// # Discussion
// 
// The objects are removed using [RemoveObjectAtIndex].
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/removeObjects(in:)-9jkis
func (m NSMutableOrderedSet) RemoveObjectsInRange(range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeObjectsInRange:"), range_)
}

// Removes all the objects from the mutable ordered set.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/removeAllObjects()
func (m NSMutableOrderedSet) RemoveAllObjects() {
	objc.Send[objc.ID](m.ID, objc.Sel("removeAllObjects"))
}

// Replaces the object at the specified index with the new object.
//
// idx: The index of the object to be replaced. This value must not exceed the
// bounds of the mutable ordered set.
//
// object: The object with which to replace the object at the index in the ordered set
// specified by `idx`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/replaceObject(at:with:)
func (m NSMutableOrderedSet) ReplaceObjectAtIndexWithObject(idx uint, object objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("replaceObjectAtIndex:withObject:"), idx, object)
}

// Replaces the objects at the specified indexes with the new objects.
//
// indexes: The indexes of the objects to be replaced.
//
// objects: The objects with which to replace the objects in the receiving mutable
// ordered set at the indexes specified by indexes.
// 
// The count of locations in `indexes` must equal the count of objects.
//
// # Discussion
// 
// The indexes in `indexes` are used in the same order as the objects in
// `objects`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/replaceObjects(at:with:)
func (m NSMutableOrderedSet) ReplaceObjectsAtIndexesWithObjects(indexes INSIndexSet, objects []objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("replaceObjectsAtIndexes:withObjects:"), indexes, objectivec.IObjectSliceToNSArray(objects))
}

// Replaces the objects in the receiving mutable ordered set at the range with
// the specified number of objects from a given C array.
//
// range: The range of the objects to replace.
//
// objects: A C array of objects.
//
// count: The number of values from the objects C array to insert in place of the
// objects in `range`. This number will be the count of the new array—it
// must not be negative or greater than the number of elements in objects.
//
// # Discussion
// 
// Elements are added to the new array in the same order they appear in
// objects, up to but not including index count.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/replaceObjects(in:with:count:)
func (m NSMutableOrderedSet) ReplaceObjectsInRangeWithObjectsCount(range_ NSRange, objects []objectivec.IObject, count uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("replaceObjectsInRange:withObjects:count:"), range_, objc.CArray(objects), count)
}

// Appends or replaces the object at the specified index.
//
// obj: The object to insert or append.
//
// idx: The index. If the index is equal to the length of the collection, then it
// inserts the object at that index, otherwise it replaces the object at that
// index with the given object.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/setObject(_:at:)
func (m NSMutableOrderedSet) SetObjectAtIndex(obj objectivec.IObject, idx uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("setObject:atIndex:"), obj, idx)
}

// Moves the objects at the specified indexes to the new location.
//
// indexes: The indexes of the objects to move.
//
// idx: The index in the mutable ordered set at which to insert the objects. The
// objects being moved are first removed from the set, then this index is used
// to find the location at which to insert the moved objects.
//
// # Discussion
// 
// For example, the following code results in the contents of `mySet` being
// equal to `["a", "c", "b", "d", "e"]:`
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/moveObjects(at:to:)
func (m NSMutableOrderedSet) MoveObjectsAtIndexesToIndex(indexes INSIndexSet, idx uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("moveObjectsAtIndexes:toIndex:"), indexes, idx)
}

// Exchanges the object at the specified index with the object at the other
// index.
//
// idx1: The index of the first object.
//
// idx2: The index of the second object.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/exchangeObject(at:withObjectAt:)
func (m NSMutableOrderedSet) ExchangeObjectAtIndexWithObjectAtIndex(idx1 uint, idx2 uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("exchangeObjectAtIndex:withObjectAtIndex:"), idx1, idx2)
}

// Evaluates a given predicate against the mutable ordered set’s content and
// leaves only objects that match.
//
// p: The predicate to evaluate against the set’s elements.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/filter(using:)
func (m NSMutableOrderedSet) FilterUsingPredicate(p INSPredicate) {
	objc.Send[objc.ID](m.ID, objc.Sel("filterUsingPredicate:"), p)
}

// Sorts the receiving ordered set using a given array of sort descriptors.
//
// sortDescriptors: An array containing the [NSSortDescriptor] objects to use to sort the
// receiving ordered set’s contents.
//
// # Discussion
// 
// See [NSSortDescriptor] for additional information.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/sort(using:)
func (m NSMutableOrderedSet) SortUsingDescriptors(sortDescriptors []NSSortDescriptor) {
	objc.Send[objc.ID](m.ID, objc.Sel("sortUsingDescriptors:"), objectivec.IObjectSliceToNSArray(sortDescriptors))
}

// Sorts the mutable ordered set using the comparison method specified by the
// comparator block.
//
// cmptr: A comparator block.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/sort(comparator:)
func (m NSMutableOrderedSet) SortUsingComparator(cmptr NSComparator) {
	objc.Send[objc.ID](m.ID, objc.Sel("sortUsingComparator:"), cmptr)
}

// Sorts the mutable ordered set using the specified options and the
// comparison method specified by a given comparator block.
//
// opts: A bitmask that specifies the options for the sort (whether it should be
// performed concurrently and whether it should be performed stably).
//
// cmptr: A comparator block.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/sort(options:usingComparator:)
func (m NSMutableOrderedSet) SortWithOptionsUsingComparator(opts NSSortOptions, cmptr NSComparator) {
	objc.Send[objc.ID](m.ID, objc.Sel("sortWithOptions:usingComparator:"), opts, cmptr)
}

// Sorts the specified range of the mutable ordered set using the specified
// options and the comparison method specified by a given comparator block.
//
// range: The range to sort.
//
// opts: A bitmask that specifies the options for the sort (whether it should be
// performed concurrently and whether it should be performed stably).
//
// cmptr: A comparator block.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/sortRange(_:options:usingComparator:)
func (m NSMutableOrderedSet) SortRangeOptionsUsingComparator(range_ NSRange, opts NSSortOptions, cmptr NSComparator) {
	objc.Send[objc.ID](m.ID, objc.Sel("sortRange:options:usingComparator:"), range_, opts, cmptr)
}

// Removes from the receiving ordered set each object that isn’t a member of
// another ordered set.
//
// other: The ordered set with which to perform the intersection.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/intersect(_:)
func (m NSMutableOrderedSet) IntersectOrderedSet(other INSOrderedSet) {
	objc.Send[objc.ID](m.ID, objc.Sel("intersectOrderedSet:"), other)
}

// Removes from the receiving ordered set each object that isn’t a member of
// another set.
//
// other: The set with which to perform the intersection.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/intersectSet(_:)
func (m NSMutableOrderedSet) IntersectSet(other INSSet) {
	objc.Send[objc.ID](m.ID, objc.Sel("intersectSet:"), other)
}

// Removes each object in another given ordered set from the receiving mutable
// ordered set, if present.
//
// other: The ordered set of objects to remove from the receiving set.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/minus(_:)
func (m NSMutableOrderedSet) MinusOrderedSet(other INSOrderedSet) {
	objc.Send[objc.ID](m.ID, objc.Sel("minusOrderedSet:"), other)
}

// Removes each object in another given set from the receiving mutable ordered
// set, if present.
//
// other: The set of objects to remove from the receiving set.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/minusSet(_:)
func (m NSMutableOrderedSet) MinusSet(other INSSet) {
	objc.Send[objc.ID](m.ID, objc.Sel("minusSet:"), other)
}

// Adds each object in another given ordered set to the receiving mutable
// ordered set, if not present.
//
// other: The set of objects to add to the receiving mutable ordered set.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/union(_:)
func (m NSMutableOrderedSet) UnionOrderedSet(other INSOrderedSet) {
	objc.Send[objc.ID](m.ID, objc.Sel("unionOrderedSet:"), other)
}

// Adds each object in another given set to the receiving mutable ordered set,
// if not present.
//
// other: The set of objects to add to the receiving mutable ordered set.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/unionSet(_:)
func (m NSMutableOrderedSet) UnionSet(other INSSet) {
	objc.Send[objc.ID](m.ID, objc.Sel("unionSet:"), other)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/applyDifference:
func (m NSMutableOrderedSet) ApplyDifference(difference INSOrderedCollectionDifference) {
	objc.Send[objc.ID](m.ID, objc.Sel("applyDifference:"), difference)
}

// Replaces the given object at the specified index of the mutable ordered
// set.
//
// obj: The object to replace the set’s content.
//
// idx: The index in the mutable ordered set at which to insert `obj`. This value
// must not be greater than the count of elements in the array.
//
// # Discussion
// 
// If the index is already occupied, the objects at index and beyond are
// shifted by adding `1` to their indices to make room.
// 
// This method is identical to [InsertObjectAtIndex].
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/setObject:atIndexedSubscript:
func (m NSMutableOrderedSet) SetObjectAtIndexedSubscript(obj objectivec.IObject, idx uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("setObject:atIndexedSubscript:"), obj, idx)
}

// Creates and returns an mutable ordered set with a given initial capacity.
//
// numItems: The initial capacity of the new ordered set.
//
// # Return Value
// 
// A mutable ordered set with initial capacity to hold `numItems` members.
//
// # Discussion
// 
// Mutable ordered sets allocate additional memory as needed, so `numItems`
// simply establishes the set’s initial capacity.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/orderedSetWithCapacity:
func (_NSMutableOrderedSetClass NSMutableOrderedSetClass) OrderedSetWithCapacity(numItems uint) NSMutableOrderedSet {
	rv := objc.Send[objc.ID](objc.ID(_NSMutableOrderedSetClass.class), objc.Sel("orderedSetWithCapacity:"), numItems)
	return NSMutableOrderedSetFromID(rv)
}

