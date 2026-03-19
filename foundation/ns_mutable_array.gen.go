// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMutableArray] class.
var (
	_NSMutableArrayClass     NSMutableArrayClass
	_NSMutableArrayClassOnce sync.Once
)

func getNSMutableArrayClass() NSMutableArrayClass {
	_NSMutableArrayClassOnce.Do(func() {
		_NSMutableArrayClass = NSMutableArrayClass{class: objc.GetClass("NSMutableArray")}
	})
	return _NSMutableArrayClass
}

// GetNSMutableArrayClass returns the class object for NSMutableArray.
func GetNSMutableArrayClass() NSMutableArrayClass {
	return getNSMutableArrayClass()
}

type NSMutableArrayClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMutableArrayClass) Alloc() NSMutableArray {
	rv := objc.Send[NSMutableArray](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A dynamic ordered collection of objects.
//
// # Overview
// 
// You can use this type in Swift instead of an [Array] variable in cases that
// require reference semantics.
// 
// The [NSMutableArray] class declares the programmatic interface to objects
// that manage a modifiable array of objects. This class adds insertion and
// deletion operations to the basic array-handling behavior inherited from
// [NSArray].
// 
// NSMutableArray is “toll-free bridged” with its Core Foundation
// counterpart, [CFMutableArray]. See [Toll-Free Bridging] for more
// information.
// 
// # Accessing Values Using Subscripting
// 
// In addition to the provided instance methods, such as
// [NSMutableArray.ReplaceObjectAtIndexWithObject], you can access [NSArray] values by their
// indexes using .
// 
// # Subclassing Notes
// 
// There is typically little reason to subclass [NSMutableArray]. The class
// does well what it is designed to do—maintain a mutable, ordered
// collection of objects. But there are situations where a custom [NSArray]
// object might come in handy. Here are a few possibilities:
// 
// - Changing how [NSMutableArray] stores the elements of its collection. You
// might do this for performance reasons or for better compatibility with
// legacy code. - Acquiring more information about what is happening to the
// collection (for example, statistics gathering).
// 
// # Methods to Override
// 
// [NSMutableArray] defines five primitive methods:
// 
// - [NSMutableArray.InsertObjectAtIndex] - [NSMutableArray.RemoveObjectAtIndex] - [NSMutableArray.AddObject] -
// [NSMutableArray.RemoveLastObject] - [NSMutableArray.ReplaceObjectAtIndexWithObject]
// 
// In a subclass, you must override all these methods. You must also override
// the primitive methods of the [NSArray] class.
//
// [Array]: https://developer.apple.com/documentation/Swift/Array
// [CFMutableArray]: https://developer.apple.com/documentation/CoreFoundation/CFMutableArray
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
//
// # Creating and Initializing a Mutable Array
//
//   - [NSMutableArray.InitWithCapacity]: Returns an array, initialized with enough memory to initially hold a given number of objects.
//
// # Adding Objects
//
//   - [NSMutableArray.AddObject]: Inserts a given object at the end of the array.
//   - [NSMutableArray.AddObjectsFromArray]: Adds the objects contained in another given array to the end of the receiving array’s content.
//   - [NSMutableArray.InsertObjectAtIndex]: Inserts a given object into the array’s contents at a given index.
//   - [NSMutableArray.InsertObjectsAtIndexes]: Inserts the objects in the provided array into the receiving array at the specified indexes.
//
// # Removing Objects
//
//   - [NSMutableArray.RemoveAllObjects]: Empties the array of all its elements.
//   - [NSMutableArray.RemoveLastObject]: Removes the object with the highest-valued index in the array
//   - [NSMutableArray.RemoveObject]: Removes all occurrences in the array of a given object.
//   - [NSMutableArray.RemoveObjectInRange]: Removes all occurrences within a specified range in the array of a given object.
//   - [NSMutableArray.RemoveObjectAtIndex]: Removes the object at `index` .
//   - [NSMutableArray.RemoveObjectsAtIndexes]: Removes the objects at the specified indexes from the array.
//   - [NSMutableArray.RemoveObjectIdenticalTo]: Removes all occurrences of a given object in the array.
//   - [NSMutableArray.RemoveObjectIdenticalToInRange]: Removes all occurrences of `anObject` within the specified range in the array.
//   - [NSMutableArray.RemoveObjectsInArray]: Removes from the receiving array the objects in another given array.
//   - [NSMutableArray.RemoveObjectsInRange]: Removes from the array each of the objects within a given range.
//
// # Replacing Objects
//
//   - [NSMutableArray.ReplaceObjectAtIndexWithObject]: Replaces the object at `index` with `anObject`.
//   - [NSMutableArray.ReplaceObjectsAtIndexesWithObjects]: Replaces the objects in the receiving array at locations specified with the objects from a given array.
//   - [NSMutableArray.ReplaceObjectsInRangeWithObjectsFromArrayRange]: Replaces the objects in the receiving array specified by one given range with the objects in another array specified by another range.
//   - [NSMutableArray.ReplaceObjectsInRangeWithObjectsFromArray]: Replaces the objects in the receiving array specified by a given range with all of the objects from a given array.
//   - [NSMutableArray.SetArray]: Sets the receiving array’s elements to those in another given array.
//
// # Filtering Content
//
//   - [NSMutableArray.FilterUsingPredicate]: Evaluates a given predicate against the array’s content and leaves only objects that match.
//
// # Rearranging Content
//
//   - [NSMutableArray.ExchangeObjectAtIndexWithObjectAtIndex]: Exchanges the objects in the array at given indexes.
//   - [NSMutableArray.SortUsingDescriptors]: Sorts the receiver using a given array of sort descriptors.
//   - [NSMutableArray.SortUsingComparator]: Sorts the receiver in ascending order using the comparison method specified by a given [Comparator](<doc://com.apple.foundation/documentation/Foundation/Comparator>) block.
//   - [NSMutableArray.SortWithOptionsUsingComparator]: Sorts the receiver in ascending order using the specified options and the comparison method specified by a given [Comparator](<doc://com.apple.foundation/documentation/Foundation/Comparator>) block.
//   - [NSMutableArray.SortUsingFunctionContext]: Sorts the receiver in ascending order as defined by the comparison function `compare`.
//   - [NSMutableArray.SortUsingSelector]: Sorts the receiver in ascending order, as determined by the comparison method specified by a given selector.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray
type NSMutableArray struct {
	NSArray
}

// NSMutableArrayFromID constructs a [NSMutableArray] from an objc.ID.
//
// A dynamic ordered collection of objects.
func NSMutableArrayFromID(id objc.ID) NSMutableArray {
	return NSMutableArray{NSArray: NSArrayFromID(id)}
}
// NOTE: NSMutableArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMutableArray] class.
//
// # Creating and Initializing a Mutable Array
//
//   - [INSMutableArray.InitWithCapacity]: Returns an array, initialized with enough memory to initially hold a given number of objects.
//
// # Adding Objects
//
//   - [INSMutableArray.AddObject]: Inserts a given object at the end of the array.
//   - [INSMutableArray.AddObjectsFromArray]: Adds the objects contained in another given array to the end of the receiving array’s content.
//   - [INSMutableArray.InsertObjectAtIndex]: Inserts a given object into the array’s contents at a given index.
//   - [INSMutableArray.InsertObjectsAtIndexes]: Inserts the objects in the provided array into the receiving array at the specified indexes.
//
// # Removing Objects
//
//   - [INSMutableArray.RemoveAllObjects]: Empties the array of all its elements.
//   - [INSMutableArray.RemoveLastObject]: Removes the object with the highest-valued index in the array
//   - [INSMutableArray.RemoveObject]: Removes all occurrences in the array of a given object.
//   - [INSMutableArray.RemoveObjectInRange]: Removes all occurrences within a specified range in the array of a given object.
//   - [INSMutableArray.RemoveObjectAtIndex]: Removes the object at `index` .
//   - [INSMutableArray.RemoveObjectsAtIndexes]: Removes the objects at the specified indexes from the array.
//   - [INSMutableArray.RemoveObjectIdenticalTo]: Removes all occurrences of a given object in the array.
//   - [INSMutableArray.RemoveObjectIdenticalToInRange]: Removes all occurrences of `anObject` within the specified range in the array.
//   - [INSMutableArray.RemoveObjectsInArray]: Removes from the receiving array the objects in another given array.
//   - [INSMutableArray.RemoveObjectsInRange]: Removes from the array each of the objects within a given range.
//
// # Replacing Objects
//
//   - [INSMutableArray.ReplaceObjectAtIndexWithObject]: Replaces the object at `index` with `anObject`.
//   - [INSMutableArray.ReplaceObjectsAtIndexesWithObjects]: Replaces the objects in the receiving array at locations specified with the objects from a given array.
//   - [INSMutableArray.ReplaceObjectsInRangeWithObjectsFromArrayRange]: Replaces the objects in the receiving array specified by one given range with the objects in another array specified by another range.
//   - [INSMutableArray.ReplaceObjectsInRangeWithObjectsFromArray]: Replaces the objects in the receiving array specified by a given range with all of the objects from a given array.
//   - [INSMutableArray.SetArray]: Sets the receiving array’s elements to those in another given array.
//
// # Filtering Content
//
//   - [INSMutableArray.FilterUsingPredicate]: Evaluates a given predicate against the array’s content and leaves only objects that match.
//
// # Rearranging Content
//
//   - [INSMutableArray.ExchangeObjectAtIndexWithObjectAtIndex]: Exchanges the objects in the array at given indexes.
//   - [INSMutableArray.SortUsingDescriptors]: Sorts the receiver using a given array of sort descriptors.
//   - [INSMutableArray.SortUsingComparator]: Sorts the receiver in ascending order using the comparison method specified by a given [Comparator](<doc://com.apple.foundation/documentation/Foundation/Comparator>) block.
//   - [INSMutableArray.SortWithOptionsUsingComparator]: Sorts the receiver in ascending order using the specified options and the comparison method specified by a given [Comparator](<doc://com.apple.foundation/documentation/Foundation/Comparator>) block.
//   - [INSMutableArray.SortUsingFunctionContext]: Sorts the receiver in ascending order as defined by the comparison function `compare`.
//   - [INSMutableArray.SortUsingSelector]: Sorts the receiver in ascending order, as determined by the comparison method specified by a given selector.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray
type INSMutableArray interface {
	INSArray

	// Topic: Creating and Initializing a Mutable Array

	// Returns an array, initialized with enough memory to initially hold a given number of objects.
	InitWithCapacity(numItems uint) NSMutableArray

	// Topic: Adding Objects

	// Inserts a given object at the end of the array.
	AddObject(anObject objectivec.IObject)
	// Adds the objects contained in another given array to the end of the receiving array’s content.
	AddObjectsFromArray(otherArray []objectivec.IObject)
	// Inserts a given object into the array’s contents at a given index.
	InsertObjectAtIndex(anObject objectivec.IObject, index uint)
	// Inserts the objects in the provided array into the receiving array at the specified indexes.
	InsertObjectsAtIndexes(objects []objectivec.IObject, indexes INSIndexSet)

	// Topic: Removing Objects

	// Empties the array of all its elements.
	RemoveAllObjects()
	// Removes the object with the highest-valued index in the array
	RemoveLastObject()
	// Removes all occurrences in the array of a given object.
	RemoveObject(anObject objectivec.IObject)
	// Removes all occurrences within a specified range in the array of a given object.
	RemoveObjectInRange(anObject objectivec.IObject, range_ NSRange)
	// Removes the object at `index` .
	RemoveObjectAtIndex(index uint)
	// Removes the objects at the specified indexes from the array.
	RemoveObjectsAtIndexes(indexes INSIndexSet)
	// Removes all occurrences of a given object in the array.
	RemoveObjectIdenticalTo(anObject objectivec.IObject)
	// Removes all occurrences of `anObject` within the specified range in the array.
	RemoveObjectIdenticalToInRange(anObject objectivec.IObject, range_ NSRange)
	// Removes from the receiving array the objects in another given array.
	RemoveObjectsInArray(otherArray []objectivec.IObject)
	// Removes from the array each of the objects within a given range.
	RemoveObjectsInRange(range_ NSRange)

	// Topic: Replacing Objects

	// Replaces the object at `index` with `anObject`.
	ReplaceObjectAtIndexWithObject(index uint, anObject objectivec.IObject)
	// Replaces the objects in the receiving array at locations specified with the objects from a given array.
	ReplaceObjectsAtIndexesWithObjects(indexes INSIndexSet, objects []objectivec.IObject)
	// Replaces the objects in the receiving array specified by one given range with the objects in another array specified by another range.
	ReplaceObjectsInRangeWithObjectsFromArrayRange(range_ NSRange, otherArray []objectivec.IObject, otherRange NSRange)
	// Replaces the objects in the receiving array specified by a given range with all of the objects from a given array.
	ReplaceObjectsInRangeWithObjectsFromArray(range_ NSRange, otherArray []objectivec.IObject)
	// Sets the receiving array’s elements to those in another given array.
	SetArray(otherArray []objectivec.IObject)

	// Topic: Filtering Content

	// Evaluates a given predicate against the array’s content and leaves only objects that match.
	FilterUsingPredicate(predicate INSPredicate)

	// Topic: Rearranging Content

	// Exchanges the objects in the array at given indexes.
	ExchangeObjectAtIndexWithObjectAtIndex(idx1 uint, idx2 uint)
	// Sorts the receiver using a given array of sort descriptors.
	SortUsingDescriptors(sortDescriptors []NSSortDescriptor)
	// Sorts the receiver in ascending order using the comparison method specified by a given [Comparator](<doc://com.apple.foundation/documentation/Foundation/Comparator>) block.
	SortUsingComparator(cmptr NSComparator)
	// Sorts the receiver in ascending order using the specified options and the comparison method specified by a given [Comparator](<doc://com.apple.foundation/documentation/Foundation/Comparator>) block.
	SortWithOptionsUsingComparator(opts NSSortOptions, cmptr NSComparator)
	// Sorts the receiver in ascending order as defined by the comparison function `compare`.
	SortUsingFunctionContext(compare objectivec.IObject, context unsafe.Pointer)
	// Sorts the receiver in ascending order, as determined by the comparison method specified by a given selector.
	SortUsingSelector(comparator objc.SEL)

	ApplyDifference(difference INSOrderedCollectionDifference)
	// Replaces the object at the index with the new object, possibly adding the object.
	SetObjectAtIndexedSubscript(obj objectivec.IObject, idx uint)
}

// Init initializes the instance.
func (m NSMutableArray) Init() NSMutableArray {
	rv := objc.Send[NSMutableArray](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMutableArray) Autorelease() NSMutableArray {
	rv := objc.Send[NSMutableArray](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMutableArray creates a new NSMutableArray instance.
func NewNSMutableArray() NSMutableArray {
	class := getNSMutableArrayClass()
	rv := objc.Send[NSMutableArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a newly allocated array by placing in it the objects contained
// in a given array.
//
// array: An array.
//
// # Return Value
// 
// An array initialized to contain the objects in `anArray`. The returned
// object might be different than the original receiver.
//
// # Discussion
// 
// After an immutable array has been initialized in this way, it cannot be
// modified.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/init(array:)-o72h
func NewMutableArrayWithArray(array []objectivec.IObject) NSMutableArray {
	instance := getNSMutableArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArray:"), objectivec.IObjectSliceToNSArray(array))
	return NSMutableArrayFromID(rv)
}

// Initializes a newly allocated array using `anArray` as the source of data
// objects for the array.
//
// array: An array containing the objects with which to initialize the new array.
//
// flag: If [true], each object in `array` receives a [copyWithZone:] message to
// create a copy of the object—objects must conform to the [NSCopying]
// protocol. In a managed memory environment, this is instead of the `retain`
// message the object would otherwise receive. The object copy is then added
// to the returned array.
// 
// If [false], then in a managed memory environment each object in `array`
// simply receives a `retain` message when it is added to the returned array.
// //
// [copyWithZone:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/copyWithZone:
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An array initialized to contain the objects—or if `flag` is [true],
// copies of the objects—in `array`. The returned object might be different
// than the original receiver.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// After an immutable array has been initialized in this way, it cannot be
// modified.
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
// See: https://developer.apple.com/documentation/Foundation/NSArray/init(array:copyItems:)
func NewMutableArrayWithArrayCopyItems(array []objectivec.IObject, flag bool) NSMutableArray {
	instance := getNSMutableArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArray:copyItems:"), objectivec.IObjectSliceToNSArray(array), flag)
	return NSMutableArrayFromID(rv)
}

// Returns an array, initialized with enough memory to initially hold a given
// number of objects.
//
// numItems: The initial capacity of the new array.
//
// # Return Value
// 
// An array initialized with enough memory to hold `numItems` objects. The
// returned object might be different than the original receiver.
//
// # Discussion
// 
// Mutable arrays expand as needed; `numItems` simply establishes the
// object’s initial capacity.
// 
// This method is a designated initializer.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/init(capacity:)
func NewMutableArrayWithCapacity(numItems uint) NSMutableArray {
	instance := getNSMutableArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCapacity:"), numItems)
	return NSMutableArrayFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/init(coder:)
func NewMutableArrayWithCoder(coder INSCoder) NSMutableArray {
	instance := getNSMutableArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSMutableArrayFromID(rv)
}

// Creates and returns an array containing a given object.
//
// anObject: An object.
//
// # Return Value
// 
// An array containing the single element `anObject`.
//
// # Discussion
// 
// Alternatively, you can use array literal syntax in Objective-C or Swift to
// create an array containing a given object:
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/init(object:)
func NewMutableArrayWithObject(anObject objectivec.IObject) NSMutableArray {
	rv := objc.Send[objc.ID](objc.ID(getNSMutableArrayClass().class), objc.Sel("arrayWithObject:"), anObject)
	return NSMutableArrayFromID(rv)
}

// Initializes a newly allocated array by placing in it the objects in the
// argument list.
//
// firstObj: The first object for the array.
//
// # Return Value
// 
// An array initialized to include the objects in the argument list. The
// returned object might be different than the original receiver.
//
// # Discussion
// 
// Pass comma-separated list of trailing variadic arguments as additional
// objects, ending with `nil`.
// 
// After an immutable array has been initialized in this way, it can’t be
// modified.
// 
// This method is a designated initializer.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/initWithObjects:
func NewMutableArrayWithObjects(firstObj objectivec.IObject) NSMutableArray {
	instance := getNSMutableArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObjects:"), firstObj)
	return NSMutableArrayFromID(rv)
}

// Returns an array, initialized with enough memory to initially hold a given
// number of objects.
//
// numItems: The initial capacity of the new array.
//
// # Return Value
// 
// An array initialized with enough memory to hold `numItems` objects. The
// returned object might be different than the original receiver.
//
// # Discussion
// 
// Mutable arrays expand as needed; `numItems` simply establishes the
// object’s initial capacity.
// 
// This method is a designated initializer.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/init(capacity:)
func (m NSMutableArray) InitWithCapacity(numItems uint) NSMutableArray {
	rv := objc.Send[NSMutableArray](m.ID, objc.Sel("initWithCapacity:"), numItems)
	return rv
}
// Inserts a given object at the end of the array.
//
// anObject: The object to add to the end of the array’s content. This value must not
// be `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/add(_:)
func (m NSMutableArray) AddObject(anObject objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("addObject:"), anObject)
}
// Adds the objects contained in another given array to the end of the
// receiving array’s content.
//
// otherArray: An array of objects to add to the end of the receiving array’s content.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/addObjects(from:)
func (m NSMutableArray) AddObjectsFromArray(otherArray []objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("addObjectsFromArray:"), objectivec.IObjectSliceToNSArray(otherArray))
}
// Inserts a given object into the array’s contents at a given index.
//
// anObject: The object to add to the array’s content. This value must not be `nil`.
//
// index: The index in the array at which to insert `anObject`. This value must not
// be greater than the count of elements in the array.
//
// # Discussion
// 
// If `index` is already occupied, the objects at `index` and beyond are
// shifted by adding `1` to their indices to make room.
// 
// Note that [NSArray] objects are not like C arrays. That is, even though you
// specify a size when you create an array, the specified size is regarded as
// a “hint”; the actual size of the array is still 0. This means that you
// cannot insert an object at an index greater than the current count of an
// array. For example, if an array contains two objects, its size is 2, so you
// can add objects at indices 0, 1, or 2. Index 3 is illegal and out of
// bounds; if you try to add an object at index 3 (when the size of the array
// is 2), [NSMutableArray] raises an exception.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/insert(_:at:)-5dbx5
func (m NSMutableArray) InsertObjectAtIndex(anObject objectivec.IObject, index uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("insertObject:atIndex:"), anObject, index)
}
// Inserts the objects in the provided array into the receiving array at the
// specified indexes.
//
// objects: An array of objects to insert into the receiving array.
//
// indexes: The indexes at which the objects in `objects` should be inserted. The count
// of locations in `indexes` must equal the count of `objects`. For more
// details, see the Discussion.
//
// # Discussion
// 
// Each object in `objects` is inserted into the receiving array in turn at
// the corresponding location specified in `indexes` after earlier insertions
// have been made. The implementation is conceptually similar to that
// illustrated in the following example:
// 
// The resulting behavior is illustrated by the following example:
// 
// The locations specified by `indexes` may therefore only exceed the bounds
// of the receiving array if one location specifies the count of the array or
// the count of the array after preceding insertions, and other locations
// exceeding the bounds do so in a contiguous fashion from that location, as
// illustrated in the following examples.
// 
// In this example, both new objects are appended to the end of the array.
// 
// If you replace `[indexes 4]` with `[indexes 6]` (so that the indexes are 5
// and 6), then the application will fail with an out of bounds exception.
// 
// In this example, two objects are added into the middle of the array, and
// another at the current end of the array (index 4) which means that it is
// third from the end of the modified array.
// 
// If you replace `[indexes 4]` with `[indexes 6]` (so that the indexes are 1,
// 2, and 6), then the output is `(one, a, b, two, three, four, c)`.
// 
// If `objects` or `indexes` is `nil`, this method will raises an exception.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/insert(_:at:)-73pln
func (m NSMutableArray) InsertObjectsAtIndexes(objects []objectivec.IObject, indexes INSIndexSet) {
	objc.Send[objc.ID](m.ID, objc.Sel("insertObjects:atIndexes:"), objectivec.IObjectSliceToNSArray(objects), indexes)
}
// Empties the array of all its elements.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeAllObjects()
func (m NSMutableArray) RemoveAllObjects() {
	objc.Send[objc.ID](m.ID, objc.Sel("removeAllObjects"))
}
// Removes the object with the highest-valued index in the array
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeLastObject()
func (m NSMutableArray) RemoveLastObject() {
	objc.Send[objc.ID](m.ID, objc.Sel("removeLastObject"))
}
// Removes all occurrences in the array of a given object.
//
// anObject: The object to remove from the array.
//
// # Discussion
// 
// This method determines a match by comparing `anObject` to the objects in
// the receiver using the `` method. If the array does not contain `anObject`,
// the method has no effect (although it does incur the overhead of searching
// the contents).
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/remove(_:)
func (m NSMutableArray) RemoveObject(anObject objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeObject:"), anObject)
}
// Removes all occurrences within a specified range in the array of a given
// object.
//
// anObject: The object to be removed from the array’s content.
//
// range: The range from which to remove `anObject`.
//
// # Discussion
// 
// Matches are determined by comparing `anObject` to the objects in the
// receiver using the `` method. If the array does not contain `anObject`
// within `aRange`, the method has no effect (although it does incur the
// overhead of searching the contents).
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/remove(_:in:)
func (m NSMutableArray) RemoveObjectInRange(anObject objectivec.IObject, range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeObject:inRange:"), anObject, range_)
}
// Removes the object at `index` .
//
// index: The index from which to remove the object in the array. The value must not
// exceed the bounds of the array.
//
// # Discussion
// 
// To fill the gap, all elements beyond `index` are moved by subtracting 1
// from their index.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObject(at:)
func (m NSMutableArray) RemoveObjectAtIndex(index uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeObjectAtIndex:"), index)
}
// Removes the objects at the specified indexes from the array.
//
// indexes: The indexes of the objects to remove from the array. The locations
// specified by `indexes` must lie within the bounds of the array.
//
// # Discussion
// 
// This method is similar to [RemoveObjectAtIndex], but allows you to
// efficiently remove multiple objects with a single operation. `indexes`
// specifies the locations of objects to be removed given the state of the
// array when the method is invoked, as illustrated in the following example:
// 
// If `indexes` is `nil`, this method raises an exception.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObjects(at:)
func (m NSMutableArray) RemoveObjectsAtIndexes(indexes INSIndexSet) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeObjectsAtIndexes:"), indexes)
}
// Removes all occurrences of a given object in the array.
//
// anObject: The object to remove from the array.
//
// # Discussion
// 
// This method determines a match by comparing the address of `anObject` to
// the addresses of objects in the receiver. If the array does not contain
// `anObject`, the method has no effect (although it does incur the overhead
// of searching the contents).
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObject(identicalTo:)
func (m NSMutableArray) RemoveObjectIdenticalTo(anObject objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeObjectIdenticalTo:"), anObject)
}
// Removes all occurrences of `anObject` within the specified range in the
// array.
//
// anObject: The object to remove from the array within `aRange`.
//
// range: The range in the array from which to remove `anObject`.
//
// # Discussion
// 
// This method determines a match by comparing the address of `anObject` to
// the addresses of objects in the receiver. If the array does not contain
// `anObject` within `aRange`, the method has no effect (although it does
// incur the overhead of searching the contents).
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObject(identicalTo:in:)
func (m NSMutableArray) RemoveObjectIdenticalToInRange(anObject objectivec.IObject, range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeObjectIdenticalTo:inRange:"), anObject, range_)
}
// Removes from the receiving array the objects in another given array.
//
// otherArray: An array containing the objects to be removed from the receiving array.
//
// # Discussion
// 
// This method is similar to [RemoveObject], but it allows you to efficiently
// remove large sets of objects with a single operation. If the receiving
// array does not contain objects in `otherArray`, the method has no effect
// (although it does incur the overhead of searching the contents).
// 
// This method assumes that all elements in `otherArray` respond to `hash` and
// ``.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObjects(in:)-4yb26
func (m NSMutableArray) RemoveObjectsInArray(otherArray []objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeObjectsInArray:"), objectivec.IObjectSliceToNSArray(otherArray))
}
// Removes from the array each of the objects within a given range.
//
// range: The range of the objects to be removed from the array.
//
// # Discussion
// 
// The objects are removed using [RemoveObjectAtIndex].
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObjects(in:)-1udmn
func (m NSMutableArray) RemoveObjectsInRange(range_ NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeObjectsInRange:"), range_)
}
// Replaces the object at `index` with `anObject`.
//
// index: The index of the object to be replaced. This value must not exceed the
// bounds of the array.
//
// anObject: The object with which to replace the object at index `index` in the array.
// This value must not be `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/replaceObject(at:with:)
func (m NSMutableArray) ReplaceObjectAtIndexWithObject(index uint, anObject objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("replaceObjectAtIndex:withObject:"), index, anObject)
}
// Replaces the objects in the receiving array at locations specified with the
// objects from a given array.
//
// indexes: The indexes of the objects to be replaced.
//
// objects: The objects with which to replace the objects in the receiving array at the
// indexes specified by `indexes`. The count of locations in `indexes` must
// equal the count of `objects`.
//
// # Discussion
// 
// The indexes in `indexes` are used in the same order as the objects in
// `objects`.
// 
// If `objects` or `indexes` is `nil`, this method raises an exception.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/replaceObjects(at:with:)
func (m NSMutableArray) ReplaceObjectsAtIndexesWithObjects(indexes INSIndexSet, objects []objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("replaceObjectsAtIndexes:withObjects:"), indexes, objectivec.IObjectSliceToNSArray(objects))
}
// Replaces the objects in the receiving array specified by one given range
// with the objects in another array specified by another range.
//
// range: The range of objects to be replaced in (or removed from) the receiving
// array.
//
// otherArray: The array of objects from which to select replacements for the objects in
// `aRange`.
//
// otherRange: The range of objects be selected from `otherArray` as replacements for the
// objects in `aRange`.
//
// # Discussion
// 
// The lengths of `aRange` and `otherRange` don’t have to be equal: If
// `aRange` is longer than `otherRange`, the extra objects in the receiving
// array are removed; if `otherRange` is longer than `aRange`, the extra
// objects from `otherArray` are inserted into the receiving array.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/replaceObjects(in:withObjectsFrom:range:)
func (m NSMutableArray) ReplaceObjectsInRangeWithObjectsFromArrayRange(range_ NSRange, otherArray []objectivec.IObject, otherRange NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("replaceObjectsInRange:withObjectsFromArray:range:"), range_, objectivec.IObjectSliceToNSArray(otherArray), otherRange)
}
// Replaces the objects in the receiving array specified by a given range with
// all of the objects from a given array.
//
// range: The range of objects to be replaced in (or removed from) the receiving
// array.
//
// otherArray: The array of objects from which to select replacements for the objects in
// `aRange`.
//
// # Discussion
// 
// If `otherArray` has fewer objects than are specified by `aRange`, the extra
// objects in the receiving array are removed. If `otherArray` has more
// objects than are specified by `aRange`, the extra objects from `otherArray`
// are inserted into the receiving array.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/replaceObjects(in:withObjectsFrom:)
func (m NSMutableArray) ReplaceObjectsInRangeWithObjectsFromArray(range_ NSRange, otherArray []objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("replaceObjectsInRange:withObjectsFromArray:"), range_, objectivec.IObjectSliceToNSArray(otherArray))
}
// Sets the receiving array’s elements to those in another given array.
//
// otherArray: The array of objects with which to replace the receiving array’s content.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/setArray(_:)
func (m NSMutableArray) SetArray(otherArray []objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setArray:"), objectivec.IObjectSliceToNSArray(otherArray))
}
// Evaluates a given predicate against the array’s content and leaves only
// objects that match.
//
// predicate: The predicate to evaluate against the array’s elements.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/filter(using:)
func (m NSMutableArray) FilterUsingPredicate(predicate INSPredicate) {
	objc.Send[objc.ID](m.ID, objc.Sel("filterUsingPredicate:"), predicate)
}
// Exchanges the objects in the array at given indexes.
//
// idx1: The index of the object with which to replace the object at index `idx2`.
//
// idx2: The index of the object with which to replace the object at index `idx1`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/exchangeObject(at:withObjectAt:)
func (m NSMutableArray) ExchangeObjectAtIndexWithObjectAtIndex(idx1 uint, idx2 uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("exchangeObjectAtIndex:withObjectAtIndex:"), idx1, idx2)
}
// Sorts the receiver using a given array of sort descriptors.
//
// sortDescriptors: An array containing the [NSSortDescriptor] objects to use to sort the
// receiving array’s contents.
//
// # Discussion
// 
// See [NSSortDescriptor] for additional information.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/sort(using:)-4eh07
func (m NSMutableArray) SortUsingDescriptors(sortDescriptors []NSSortDescriptor) {
	objc.Send[objc.ID](m.ID, objc.Sel("sortUsingDescriptors:"), objectivec.IObjectSliceToNSArray(sortDescriptors))
}
// Sorts the receiver in ascending order using the comparison method specified
// by a given [NSComparator] block.
//
// cmptr: A comparator block.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/sort(comparator:)
func (m NSMutableArray) SortUsingComparator(cmptr NSComparator) {
	objc.Send[objc.ID](m.ID, objc.Sel("sortUsingComparator:"), cmptr)
}
// Sorts the receiver in ascending order using the specified options and the
// comparison method specified by a given [NSComparator] block.
//
// opts: A bitmask that specifies the options for the sort (whether it should be
// performed concurrently and whether it should be performed stably).
//
// cmptr: A comparator block.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/sort(options:usingComparator:)
func (m NSMutableArray) SortWithOptionsUsingComparator(opts NSSortOptions, cmptr NSComparator) {
	objc.Send[objc.ID](m.ID, objc.Sel("sortWithOptions:usingComparator:"), opts, cmptr)
}
// Sorts the receiver in ascending order as defined by the comparison function
// `compare`.
//
// compare: The comparison function to use to compare two elements at a time.
// 
// The function’s parameters are two objects to compare and the context
// parameter, `context`. The function should return [NSOrderedAscending] if
// the first element is smaller than the second, [NSOrderedDescending] if the
// first element is larger than the second, and [NSOrderedSame] if the
// elements are equal.
//
// context: The context argument to be passed to the compare function.
//
// # Discussion
// 
// This approach allows the comparison to be based on some outside parameter,
// such as whether character sorting is case sensitive or case insensitive.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/sort(_:context:)
func (m NSMutableArray) SortUsingFunctionContext(compare objectivec.IObject, context unsafe.Pointer) {
	objc.Send[objc.ID](m.ID, objc.Sel("sortUsingFunction:context:"), compare, context)
}
// Sorts the receiver in ascending order, as determined by the comparison
// method specified by a given selector.
//
// comparator: A selector that specifies the comparison method to use to compare elements
// in the array.
// 
// The `comparator` message is sent to each object in the array and has as its
// single argument another object in the array. The `comparator` method should
// return [NSOrderedAscending] if the array is smaller than the argument,
// [NSOrderedDescending] if the array is larger than the argument, and
// [NSOrderedSame] if they are equal.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/sort(using:)-537vs
func (m NSMutableArray) SortUsingSelector(comparator objc.SEL) {
	objc.Send[objc.ID](m.ID, objc.Sel("sortUsingSelector:"), comparator)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/applyDifference:
func (m NSMutableArray) ApplyDifference(difference INSOrderedCollectionDifference) {
	objc.Send[objc.ID](m.ID, objc.Sel("applyDifference:"), difference)
}
// Replaces the object at the index with the new object, possibly adding the
// object.
//
// obj: The object with which to replace the object at index `idx` in the array.
// This value must not be `nil`.
//
// idx: The index of the object to be replaced. This value must not exceed the
// bounds of the array.
//
// # Discussion
// 
// This method has the same behavior as the [ReplaceObjectAtIndexWithObject]
// method.
// 
// If `idx` is beyond the end of the array (that is, if `idx` is greater than
// the value returned by `count`), an [rangeException] is raised.
// 
// You shouldn’t need to call this method directly. Instead, this method is
// called when setting an object by index using subscripting.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/setObject:atIndexedSubscript:
func (m NSMutableArray) SetObjectAtIndexedSubscript(obj objectivec.IObject, idx uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("setObject:atIndexedSubscript:"), obj, idx)
}

// Creates and returns an [NSMutableArray] object with enough allocated memory
// to initially hold a given number of objects.
//
// numItems: The initial capacity of the new array.
//
// # Return Value
// 
// A new [NSMutableArray] object with enough allocated memory to hold
// `numItems` objects.
//
// # Discussion
// 
// Mutable arrays expand as needed; `numItems` simply establishes the
// object’s initial capacity.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableArray/arrayWithCapacity:
func (_NSMutableArrayClass NSMutableArrayClass) ArrayWithCapacity(numItems uint) NSMutableArray {
	rv := objc.Send[objc.ID](objc.ID(_NSMutableArrayClass.class), objc.Sel("arrayWithCapacity:"), numItems)
	return NSMutableArrayFromID(rv)
}

