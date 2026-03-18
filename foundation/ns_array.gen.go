// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSArray] class.
var (
	_NSArrayClass     NSArrayClass
	_NSArrayClassOnce sync.Once
)

func getNSArrayClass() NSArrayClass {
	_NSArrayClassOnce.Do(func() {
		_NSArrayClass = NSArrayClass{class: objc.GetClass("NSArray")}
	})
	return _NSArrayClass
}

// GetNSArrayClass returns the class object for NSArray.
func GetNSArrayClass() NSArrayClass {
	return getNSArrayClass()
}

type NSArrayClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSArrayClass) Alloc() NSArray {
	rv := objc.Send[NSArray](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A static ordered collection of objects.
//
// # Overview
// 
// You can use this type in Swift instead of an [Array] constant in cases that
// require reference semantics.
// 
// [NSArray] and its subclass [NSMutableArray] manage ordered collections of
// objects called . [NSArray] creates static arrays, and [NSMutableArray]
// creates dynamic arrays. You can use arrays when you need an ordered
// collection of objects.
// 
// [NSArray] is “toll-free bridged” with its Core Foundation counterpart,
// [CFArray]. See [Toll-Free Bridging] for more information on toll-free
// bridging.
// 
// # Creating NSArray Objects Using Array Literals
// 
// In addition to the provided initializers, such as [NSArray.InitWithObjects], you
// can create an [NSArray] object using an .
// 
// In Objective-C, the compiler generates code that makes an underlying call
// to the [NSArray.ArrayWithObjectsCount] method.
// 
// You should not terminate the list of objects with `nil` when using this
// literal syntax, and in fact `nil` is an invalid value. For more information
// about object literals in Objective-C, see [Working with Objects] in
// [Programming with Objective-C].
// 
// In Swift, the [NSArray] class conforms to the [ArrayLiteralConvertible]
// protocol, which allows it to be initialized with array literals. For more
// information about object literals in Swift, see [Literal Expression] in
// [The Swift Programming Language (Swift 4.1)].
// 
// # Accessing Values Using Subscripting
// 
// In addition to the provided instance methods, such as [NSArray.ObjectAtIndex], you
// can access [NSArray] values by their indexes using .
// 
// # Subclassing Notes
// 
// There is typically little reason to subclass [NSArray]. The class does well
// what it is designed to do—maintain an ordered collection of objects. But
// there are situations where a custom [NSArray] object might come in handy.
// Here are a few possibilities:
// 
// - Changing how [NSArray] stores the elements of its collection. You might
// do this for performance reasons or for better compatibility with legacy
// code. - Acquiring more information about what is happening to the
// collection (for example, statistics gathering).
// 
// # Methods to Override
// 
// Any subclass of [NSArray] override the primitive instance methods [NSArray.Count]
// and [NSArray.ObjectAtIndex]. These methods must operate on the backing store that
// you provide for the elements of the collection. For this backing store you
// can use a static array, a standard [NSArray] object, or some other data
// type or mechanism. You may also choose to override, partially or fully, any
// other [NSArray] method for which you want to provide an alternative
// implementation.
// 
// You might want to implement an initializer for your subclass that is suited
// to the backing store that the subclass is managing. If you do, your
// initializer must invoke one of the designated initializers of the [NSArray]
// class, either [NSArray.Init] or [NSArray.InitWithObjectsCount]. The [NSArray] class adopts
// the [NSCopying], [NSMutableCopying], and [NSCoding] protocols; custom
// subclasses of [NSArray] should override the methods in these protocols as
// necessary.
// 
// Remember that [NSArray] is the public interface for a class cluster and
// what this entails for your subclass. You must provide the storage for your
// subclass and implement the primitive methods that directly act on that
// storage.
// 
// # Alternatives to Subclassing
// 
// Before making a custom subclass of [NSArray], investigate [NSPointerArray]
// and the corresponding Core Foundation type, [CFArray]. Because [NSArray]
// and [CFArray] are “toll-free bridged,” you can substitute a [CFArray]
// object for a [NSArray] object in your code (with appropriate casting).
// Although they are corresponding types, [CFArray] and [NSArray] do not have
// identical interfaces or implementations, and you can sometimes do things
// with [CFArray] that you cannot easily do with [NSArray]. For example,
// [CFArray] provides a set of callbacks, some of which are for implementing
// custom retain-release behavior. If you specify [NULL] implementations for
// these callbacks, you can easily get a non-retaining array.
// 
// If the behavior you want to add supplements that of the existing class, you
// could write a category on [NSArray]. Keep in mind, however, that this
// category will be in effect for all instances of [NSArray] that you use, and
// this might have unintended consequences. Alternatively, you could use
// composition to achieve the desired behavior.
//
// [Array]: https://developer.apple.com/documentation/Swift/Array
// [CFArray]: https://developer.apple.com/documentation/CoreFoundation/CFArray
// [Literal Expression]: https://developer.apple.com/library/archive/documentation/Swift/Conceptual/Swift_Programming_Language/Expressions.html#//apple_ref/doc/uid/TP40014097-CH32-ID390
// [Programming with Objective-C]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ProgrammingWithObjectiveC/Introduction/Introduction.html#//apple_ref/doc/uid/TP40011210
// [The Swift Programming Language (Swift 4.1)]: https://developer.apple.com/library/archive/documentation/Swift/Conceptual/Swift_Programming_Language/index.html#//apple_ref/doc/uid/TP40014097
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
// [Working with Objects]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ProgrammingWithObjectiveC/WorkingwithObjects/WorkingwithObjects.html#//apple_ref/doc/uid/TP40011210-CH4
//
// # Initializing an Array
//
//   - [NSArray.InitWithArray]: Initializes a newly allocated array by placing in it the objects contained in a given array.
//   - [NSArray.InitWithArrayCopyItems]: Initializes a newly allocated array using `anArray` as the source of data objects for the array.
//   - [NSArray.InitWithObjectsCount]: Initializes a newly allocated array to include a given number of objects from a given C array.
//
// # Querying an Array
//
//   - [NSArray.ContainsObject]: Returns a Boolean value that indicates whether a given object is present in the array.
//   - [NSArray.Count]: The number of objects in the array.
//   - [NSArray.FirstObject]: The first object in the array.
//   - [NSArray.LastObject]: The last object in the array.
//   - [NSArray.ObjectAtIndex]: Returns the object located at the specified index.
//   - [NSArray.ObjectAtIndexedSubscript]: Returns the object at the specified index.
//   - [NSArray.ObjectsAtIndexes]: Returns an array containing the objects in the array at the indexes specified by a given index set.
//   - [NSArray.ObjectEnumerator]: Returns an enumerator object that lets you access each object in the array.
//   - [NSArray.ReverseObjectEnumerator]: Returns an enumerator object that lets you access each object in the array, in reverse order.
//
// # Finding Objects in an Array
//
//   - [NSArray.IndexOfObject]: Returns the lowest index whose corresponding array value is equal to a given object.
//   - [NSArray.IndexOfObjectInRange]: Returns the lowest index within a specified range whose corresponding array value is equal to a given object .
//   - [NSArray.IndexOfObjectIdenticalTo]: Returns the lowest index whose corresponding array value is identical to a given object.
//   - [NSArray.IndexOfObjectIdenticalToInRange]: Returns the lowest index within a specified range whose corresponding array value is equal to a given object .
//   - [NSArray.IndexOfObjectPassingTest]: Returns the index of the first object in the array that passes a test in a given block.
//   - [NSArray.IndexOfObjectWithOptionsPassingTest]: Returns the index of an object in the array that passes a test in a given block for a given set of enumeration options.
//   - [NSArray.IndexOfObjectAtIndexesOptionsPassingTest]: Returns the index, from a given set of indexes, of the first object in the array that passes a test in a given block for a given set of enumeration options.
//   - [NSArray.IndexesOfObjectsPassingTest]: Returns the indexes of objects in the array that pass a test in a given block.
//   - [NSArray.IndexesOfObjectsWithOptionsPassingTest]: Returns the indexes of objects in the array that pass a test in a given block for a given set of enumeration options.
//   - [NSArray.IndexesOfObjectsAtIndexesOptionsPassingTest]: Returns the indexes, from a given set of indexes, of objects in the array that pass a test in a given block for a given set of enumeration options.
//   - [NSArray.IndexOfObjectInSortedRangeOptionsUsingComparator]: Returns the index, within a specified range, of an object compared with elements in the array using a given [NSComparator] block.
//
// # Sending Messages to Elements
//
//   - [NSArray.EnumerateObjectsUsingBlock]: Executes a given closure or block using each object in the array, starting with the first object and continuing through the array to the last object.
//   - [NSArray.EnumerateObjectsWithOptionsUsingBlock]: Executes a given closure or block using each object in the array with the specified options.
//   - [NSArray.EnumerateObjectsAtIndexesOptionsUsingBlock]: Executes a given block using the objects in the array at the specified indexes.
//
// # Comparing Arrays
//
//   - [NSArray.FirstObjectCommonWithArray]: Returns the first object contained in the receiving array that’s equal to an object in another given array.
//   - [NSArray.IsEqualToArray]: Compares the receiving array to another array.
//
// # Deriving New Arrays
//
//   - [NSArray.ArrayByAddingObject]: Returns a new array that is a copy of the receiving array with a given object added to the end.
//   - [NSArray.ArrayByAddingObjectsFromArray]: Returns a new array that is a copy of the receiving array with the objects contained in another array added to the end.
//   - [NSArray.FilteredArrayUsingPredicate]: Evaluates a given predicate against each object in the receiving array and returns a new array containing the objects for which the predicate returns true.
//   - [NSArray.SubarrayWithRange]: Returns a new array containing the receiving array’s elements that fall within the limits specified by a given range.
//
// # Sorting
//
//   - [NSArray.SortedArrayHint]: Analyzes the array and returns a “hint” that speeds the sorting of the array when the hint is supplied to [sortedArray(_:context:hint:)](<doc://com.apple.foundation/documentation/Foundation/NSArray/sortedArray(_:context:hint:)>).
//   - [NSArray.SortedArrayUsingFunctionContext]: Returns a new array that lists the receiving array’s elements in ascending order as defined by the comparison function `comparator`.
//   - [NSArray.SortedArrayUsingFunctionContextHint]: Returns a new array that lists the receiving array’s elements in ascending order as defined by the comparison function `comparator`.
//   - [NSArray.SortedArrayUsingDescriptors]: Returns a copy of the receiving array sorted as specified by a given array of sort descriptors.
//   - [NSArray.SortedArrayUsingSelector]: Returns an array that lists the receiving array’s elements in ascending order, as determined by the comparison method specified by a given selector.
//   - [NSArray.SortedArrayUsingComparator]: Returns an array that lists the receiving array’s elements in ascending order, as determined by the comparison method specified by a given [NSComparator] block.
//   - [NSArray.SortedArrayWithOptionsUsingComparator]: Returns an array that lists the receiving array’s elements in ascending order, as determined by the comparison method specified by a given [NSComparator] block.
//
// # Working with String Elements
//
//   - [NSArray.ComponentsJoinedByString]: Constructs and returns an [NSString] object that is the result of interposing a given separator between the elements of the array.
//
// # Creating a Description
//
//   - [NSArray.Description]: A string that represents the contents of the array, formatted as a property list.
//   - [NSArray.DescriptionWithLocale]: Returns a string that represents the contents of the array, formatted as a property list.
//   - [NSArray.DescriptionWithLocaleIndent]: Returns a string that represents the contents of the array, formatted as a property list.
//
// # Collecting Paths
//
//   - [NSArray.PathsMatchingExtensions]: Returns an array containing all the pathname elements in the receiving array that have filename extensions from a given array.
//
// # Key-Value Observing
//
//   - [NSArray.RemoveObserverFromObjectsAtIndexesForKeyPathContext]: Raises an exception.
//   - [NSArray.AddObserverToObjectsAtIndexesForKeyPathOptionsContext]: Registers an observer to receive key value observer notifications for the specified key-path relative to the objects at the indexes.
//   - [NSArray.RemoveObserverFromObjectsAtIndexesForKeyPath]: Removes `anObserver` from all key value observer notifications associated with the specified `keyPath` relative to the array’s objects at `indexes`.
//
// # Randomly Shuffling an Array
//
//   - [NSArray.ShuffledArray]: Returns a new array that lists this array’s elements in a random order.
//   - [NSArray.ShuffledArrayWithRandomSource]: Returns a new array that lists this array’s elements in a random order, using the specified random source.
//
// # Initializers
//
//   - [NSArray.InitWithContentsOfURLError]
//
// # Instance Methods
//
//   - [NSArray.WriteToURLError]
//
// See: https://developer.apple.com/documentation/Foundation/NSArray
type NSArray struct {
	objectivec.Object
}

// NSArrayFromID constructs a [NSArray] from an objc.ID.
//
// A static ordered collection of objects.
func NSArrayFromID(id objc.ID) NSArray {
	return NSArray{objectivec.Object{ID: id}}
}
// NOTE: NSArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSArray] class.
//
// # Initializing an Array
//
//   - [INSArray.InitWithArray]: Initializes a newly allocated array by placing in it the objects contained in a given array.
//   - [INSArray.InitWithArrayCopyItems]: Initializes a newly allocated array using `anArray` as the source of data objects for the array.
//   - [INSArray.InitWithObjectsCount]: Initializes a newly allocated array to include a given number of objects from a given C array.
//
// # Querying an Array
//
//   - [INSArray.ContainsObject]: Returns a Boolean value that indicates whether a given object is present in the array.
//   - [INSArray.Count]: The number of objects in the array.
//   - [INSArray.FirstObject]: The first object in the array.
//   - [INSArray.LastObject]: The last object in the array.
//   - [INSArray.ObjectAtIndex]: Returns the object located at the specified index.
//   - [INSArray.ObjectAtIndexedSubscript]: Returns the object at the specified index.
//   - [INSArray.ObjectsAtIndexes]: Returns an array containing the objects in the array at the indexes specified by a given index set.
//   - [INSArray.ObjectEnumerator]: Returns an enumerator object that lets you access each object in the array.
//   - [INSArray.ReverseObjectEnumerator]: Returns an enumerator object that lets you access each object in the array, in reverse order.
//
// # Finding Objects in an Array
//
//   - [INSArray.IndexOfObject]: Returns the lowest index whose corresponding array value is equal to a given object.
//   - [INSArray.IndexOfObjectInRange]: Returns the lowest index within a specified range whose corresponding array value is equal to a given object .
//   - [INSArray.IndexOfObjectIdenticalTo]: Returns the lowest index whose corresponding array value is identical to a given object.
//   - [INSArray.IndexOfObjectIdenticalToInRange]: Returns the lowest index within a specified range whose corresponding array value is equal to a given object .
//   - [INSArray.IndexOfObjectPassingTest]: Returns the index of the first object in the array that passes a test in a given block.
//   - [INSArray.IndexOfObjectWithOptionsPassingTest]: Returns the index of an object in the array that passes a test in a given block for a given set of enumeration options.
//   - [INSArray.IndexOfObjectAtIndexesOptionsPassingTest]: Returns the index, from a given set of indexes, of the first object in the array that passes a test in a given block for a given set of enumeration options.
//   - [INSArray.IndexesOfObjectsPassingTest]: Returns the indexes of objects in the array that pass a test in a given block.
//   - [INSArray.IndexesOfObjectsWithOptionsPassingTest]: Returns the indexes of objects in the array that pass a test in a given block for a given set of enumeration options.
//   - [INSArray.IndexesOfObjectsAtIndexesOptionsPassingTest]: Returns the indexes, from a given set of indexes, of objects in the array that pass a test in a given block for a given set of enumeration options.
//   - [INSArray.IndexOfObjectInSortedRangeOptionsUsingComparator]: Returns the index, within a specified range, of an object compared with elements in the array using a given [NSComparator] block.
//
// # Sending Messages to Elements
//
//   - [INSArray.EnumerateObjectsUsingBlock]: Executes a given closure or block using each object in the array, starting with the first object and continuing through the array to the last object.
//   - [INSArray.EnumerateObjectsWithOptionsUsingBlock]: Executes a given closure or block using each object in the array with the specified options.
//   - [INSArray.EnumerateObjectsAtIndexesOptionsUsingBlock]: Executes a given block using the objects in the array at the specified indexes.
//
// # Comparing Arrays
//
//   - [INSArray.FirstObjectCommonWithArray]: Returns the first object contained in the receiving array that’s equal to an object in another given array.
//   - [INSArray.IsEqualToArray]: Compares the receiving array to another array.
//
// # Deriving New Arrays
//
//   - [INSArray.ArrayByAddingObject]: Returns a new array that is a copy of the receiving array with a given object added to the end.
//   - [INSArray.ArrayByAddingObjectsFromArray]: Returns a new array that is a copy of the receiving array with the objects contained in another array added to the end.
//   - [INSArray.FilteredArrayUsingPredicate]: Evaluates a given predicate against each object in the receiving array and returns a new array containing the objects for which the predicate returns true.
//   - [INSArray.SubarrayWithRange]: Returns a new array containing the receiving array’s elements that fall within the limits specified by a given range.
//
// # Sorting
//
//   - [INSArray.SortedArrayHint]: Analyzes the array and returns a “hint” that speeds the sorting of the array when the hint is supplied to [sortedArray(_:context:hint:)](<doc://com.apple.foundation/documentation/Foundation/NSArray/sortedArray(_:context:hint:)>).
//   - [INSArray.SortedArrayUsingFunctionContext]: Returns a new array that lists the receiving array’s elements in ascending order as defined by the comparison function `comparator`.
//   - [INSArray.SortedArrayUsingFunctionContextHint]: Returns a new array that lists the receiving array’s elements in ascending order as defined by the comparison function `comparator`.
//   - [INSArray.SortedArrayUsingDescriptors]: Returns a copy of the receiving array sorted as specified by a given array of sort descriptors.
//   - [INSArray.SortedArrayUsingSelector]: Returns an array that lists the receiving array’s elements in ascending order, as determined by the comparison method specified by a given selector.
//   - [INSArray.SortedArrayUsingComparator]: Returns an array that lists the receiving array’s elements in ascending order, as determined by the comparison method specified by a given [NSComparator] block.
//   - [INSArray.SortedArrayWithOptionsUsingComparator]: Returns an array that lists the receiving array’s elements in ascending order, as determined by the comparison method specified by a given [NSComparator] block.
//
// # Working with String Elements
//
//   - [INSArray.ComponentsJoinedByString]: Constructs and returns an [NSString] object that is the result of interposing a given separator between the elements of the array.
//
// # Creating a Description
//
//   - [INSArray.Description]: A string that represents the contents of the array, formatted as a property list.
//   - [INSArray.DescriptionWithLocale]: Returns a string that represents the contents of the array, formatted as a property list.
//   - [INSArray.DescriptionWithLocaleIndent]: Returns a string that represents the contents of the array, formatted as a property list.
//
// # Collecting Paths
//
//   - [INSArray.PathsMatchingExtensions]: Returns an array containing all the pathname elements in the receiving array that have filename extensions from a given array.
//
// # Key-Value Observing
//
//   - [INSArray.RemoveObserverFromObjectsAtIndexesForKeyPathContext]: Raises an exception.
//   - [INSArray.AddObserverToObjectsAtIndexesForKeyPathOptionsContext]: Registers an observer to receive key value observer notifications for the specified key-path relative to the objects at the indexes.
//   - [INSArray.RemoveObserverFromObjectsAtIndexesForKeyPath]: Removes `anObserver` from all key value observer notifications associated with the specified `keyPath` relative to the array’s objects at `indexes`.
//
// # Randomly Shuffling an Array
//
//   - [INSArray.ShuffledArray]: Returns a new array that lists this array’s elements in a random order.
//   - [INSArray.ShuffledArrayWithRandomSource]: Returns a new array that lists this array’s elements in a random order, using the specified random source.
//
// # Initializers
//
//   - [INSArray.InitWithContentsOfURLError]
//
// # Instance Methods
//
//   - [INSArray.WriteToURLError]
//
// See: https://developer.apple.com/documentation/Foundation/NSArray
type INSArray interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSMutableCopying
	NSSecureCoding

	// Topic: Initializing an Array

	// Initializes a newly allocated array by placing in it the objects contained in a given array.
	InitWithArray(array []objectivec.IObject) NSArray
	// Initializes a newly allocated array using `anArray` as the source of data objects for the array.
	InitWithArrayCopyItems(array []objectivec.IObject, flag bool) NSArray
	// Initializes a newly allocated array to include a given number of objects from a given C array.
	InitWithObjectsCount(objects []objectivec.IObject, cnt uint) NSArray

	// Topic: Querying an Array

	// Returns a Boolean value that indicates whether a given object is present in the array.
	ContainsObject(anObject objectivec.IObject) bool
	// The number of objects in the array.
	Count() uint
	// The first object in the array.
	FirstObject() objectivec.IObject
	// The last object in the array.
	LastObject() objectivec.IObject
	// Returns the object located at the specified index.
	ObjectAtIndex(index uint) objectivec.IObject
	// Returns the object at the specified index.
	ObjectAtIndexedSubscript(idx uint) objectivec.IObject
	// Returns an array containing the objects in the array at the indexes specified by a given index set.
	ObjectsAtIndexes(indexes INSIndexSet) []objectivec.IObject
	// Returns an enumerator object that lets you access each object in the array.
	ObjectEnumerator() INSEnumerator
	// Returns an enumerator object that lets you access each object in the array, in reverse order.
	ReverseObjectEnumerator() INSEnumerator

	// Topic: Finding Objects in an Array

	// Returns the lowest index whose corresponding array value is equal to a given object.
	IndexOfObject(anObject objectivec.IObject) uint
	// Returns the lowest index within a specified range whose corresponding array value is equal to a given object .
	IndexOfObjectInRange(anObject objectivec.IObject, range_ NSRange) uint
	// Returns the lowest index whose corresponding array value is identical to a given object.
	IndexOfObjectIdenticalTo(anObject objectivec.IObject) uint
	// Returns the lowest index within a specified range whose corresponding array value is equal to a given object .
	IndexOfObjectIdenticalToInRange(anObject objectivec.IObject, range_ NSRange) uint
	// Returns the index of the first object in the array that passes a test in a given block.
	IndexOfObjectPassingTest(predicate unsafe.Pointer) uint
	// Returns the index of an object in the array that passes a test in a given block for a given set of enumeration options.
	IndexOfObjectWithOptionsPassingTest(opts NSEnumerationOptions, predicate unsafe.Pointer) uint
	// Returns the index, from a given set of indexes, of the first object in the array that passes a test in a given block for a given set of enumeration options.
	IndexOfObjectAtIndexesOptionsPassingTest(s INSIndexSet, opts NSEnumerationOptions, predicate unsafe.Pointer) uint
	// Returns the indexes of objects in the array that pass a test in a given block.
	IndexesOfObjectsPassingTest(predicate unsafe.Pointer) INSIndexSet
	// Returns the indexes of objects in the array that pass a test in a given block for a given set of enumeration options.
	IndexesOfObjectsWithOptionsPassingTest(opts NSEnumerationOptions, predicate unsafe.Pointer) INSIndexSet
	// Returns the indexes, from a given set of indexes, of objects in the array that pass a test in a given block for a given set of enumeration options.
	IndexesOfObjectsAtIndexesOptionsPassingTest(s INSIndexSet, opts NSEnumerationOptions, predicate unsafe.Pointer) INSIndexSet
	// Returns the index, within a specified range, of an object compared with elements in the array using a given [NSComparator] block.
	IndexOfObjectInSortedRangeOptionsUsingComparator(obj objectivec.IObject, r NSRange, opts NSBinarySearchingOptions, cmp NSComparator) uint

	// Topic: Sending Messages to Elements

	// Executes a given closure or block using each object in the array, starting with the first object and continuing through the array to the last object.
	EnumerateObjectsUsingBlock(block unsafe.Pointer)
	// Executes a given closure or block using each object in the array with the specified options.
	EnumerateObjectsWithOptionsUsingBlock(opts NSEnumerationOptions, block unsafe.Pointer)
	// Executes a given block using the objects in the array at the specified indexes.
	EnumerateObjectsAtIndexesOptionsUsingBlock(s INSIndexSet, opts NSEnumerationOptions, block unsafe.Pointer)

	// Topic: Comparing Arrays

	// Returns the first object contained in the receiving array that’s equal to an object in another given array.
	FirstObjectCommonWithArray(otherArray []objectivec.IObject) objectivec.IObject
	// Compares the receiving array to another array.
	IsEqualToArray(otherArray []objectivec.IObject) bool

	// Topic: Deriving New Arrays

	// Returns a new array that is a copy of the receiving array with a given object added to the end.
	ArrayByAddingObject(anObject objectivec.IObject) []objectivec.IObject
	// Returns a new array that is a copy of the receiving array with the objects contained in another array added to the end.
	ArrayByAddingObjectsFromArray(otherArray []objectivec.IObject) []objectivec.IObject
	// Evaluates a given predicate against each object in the receiving array and returns a new array containing the objects for which the predicate returns true.
	FilteredArrayUsingPredicate(predicate INSPredicate) []objectivec.IObject
	// Returns a new array containing the receiving array’s elements that fall within the limits specified by a given range.
	SubarrayWithRange(range_ NSRange) []objectivec.IObject

	// Topic: Sorting

	// Analyzes the array and returns a “hint” that speeds the sorting of the array when the hint is supplied to [sortedArray(_:context:hint:)](<doc://com.apple.foundation/documentation/Foundation/NSArray/sortedArray(_:context:hint:)>).
	SortedArrayHint() INSData
	// Returns a new array that lists the receiving array’s elements in ascending order as defined by the comparison function `comparator`.
	SortedArrayUsingFunctionContext(comparator objectivec.IObject, context unsafe.Pointer) []objectivec.IObject
	// Returns a new array that lists the receiving array’s elements in ascending order as defined by the comparison function `comparator`.
	SortedArrayUsingFunctionContextHint(comparator objectivec.IObject, context unsafe.Pointer, hint INSData) []objectivec.IObject
	// Returns a copy of the receiving array sorted as specified by a given array of sort descriptors.
	SortedArrayUsingDescriptors(sortDescriptors []NSSortDescriptor) []objectivec.IObject
	// Returns an array that lists the receiving array’s elements in ascending order, as determined by the comparison method specified by a given selector.
	SortedArrayUsingSelector(comparator objc.SEL) []objectivec.IObject
	// Returns an array that lists the receiving array’s elements in ascending order, as determined by the comparison method specified by a given [NSComparator] block.
	SortedArrayUsingComparator(cmptr NSComparator) []objectivec.IObject
	// Returns an array that lists the receiving array’s elements in ascending order, as determined by the comparison method specified by a given [NSComparator] block.
	SortedArrayWithOptionsUsingComparator(opts NSSortOptions, cmptr NSComparator) []objectivec.IObject

	// Topic: Working with String Elements

	// Constructs and returns an [NSString] object that is the result of interposing a given separator between the elements of the array.
	ComponentsJoinedByString(separator string) string

	// Topic: Creating a Description

	// A string that represents the contents of the array, formatted as a property list.
	Description() string
	// Returns a string that represents the contents of the array, formatted as a property list.
	DescriptionWithLocale(locale objectivec.IObject) string
	// Returns a string that represents the contents of the array, formatted as a property list.
	DescriptionWithLocaleIndent(locale objectivec.IObject, level uint) string

	// Topic: Collecting Paths

	// Returns an array containing all the pathname elements in the receiving array that have filename extensions from a given array.
	PathsMatchingExtensions(filterTypes []string) []string

	// Topic: Key-Value Observing

	// Raises an exception.
	RemoveObserverFromObjectsAtIndexesForKeyPathContext(observer objectivec.Object, indexes INSIndexSet, keyPath string, context unsafe.Pointer)
	// Registers an observer to receive key value observer notifications for the specified key-path relative to the objects at the indexes.
	AddObserverToObjectsAtIndexesForKeyPathOptionsContext(observer objectivec.Object, indexes INSIndexSet, keyPath string, options uint, context unsafe.Pointer)
	// Removes `anObserver` from all key value observer notifications associated with the specified `keyPath` relative to the array’s objects at `indexes`.
	RemoveObserverFromObjectsAtIndexesForKeyPath(observer objectivec.Object, indexes INSIndexSet, keyPath string)

	// Topic: Randomly Shuffling an Array

	// Returns a new array that lists this array’s elements in a random order.
	ShuffledArray() []objectivec.IObject
	// Returns a new array that lists this array’s elements in a random order, using the specified random source.
	ShuffledArrayWithRandomSource(randomSource objectivec.IObject) []objectivec.IObject

	// Topic: Initializers

	InitWithContentsOfURLError(url INSURL) (NSArray, error)

	// Topic: Instance Methods

	WriteToURLError(url INSURL) (bool, error)

	// Creates a new array by applying a difference object to an existing array.
	ArrayByApplyingDifference(difference INSOrderedCollectionDifference) []objectivec.IObject
	// Compares two arrays to create a difference object that represents the changes between them.
	DifferenceFromArray(other []objectivec.IObject) INSOrderedCollectionDifference
	// Compares two arrays, with options, to create a difference object that represents the changes between them.
	DifferenceFromArrayWithOptions(other []objectivec.IObject, options NSOrderedCollectionDifferenceCalculationOptions) INSOrderedCollectionDifference
	// Compares two arrays, using the provided block and with options, to create a difference object that represents the changes between them.
	DifferenceFromArrayWithOptionsUsingEquivalenceTest(other []objectivec.IObject, options NSOrderedCollectionDifferenceCalculationOptions, block bool) INSOrderedCollectionDifference
	// Copies references to objects contained in the array that fall within the specified range to `aBuffer`.
	GetObjectsRange(objects []objectivec.IObject, range_ NSRange)
	// Initializes a newly allocated array by placing in it the objects in the argument list.
	InitWithObjects(firstObj objectivec.IObject) NSArray
	// Sends to each object in the array the message identified by a given selector, starting with the first object and continuing through the array to the last object.
	MakeObjectsPerformSelector(aSelector objc.SEL)
	// Sends the `aSelector` message to each object in the array, starting with the first object and continuing through the array to the last object.
	MakeObjectsPerformSelectorWithObject(aSelector objc.SEL, argument objectivec.IObject)
}

// Init initializes the instance.
func (a NSArray) Init() NSArray {
	rv := objc.Send[NSArray](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSArray) Autorelease() NSArray {
	rv := objc.Send[NSArray](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSArray creates a new NSArray instance.
func NewNSArray() NSArray {
	class := getNSArrayClass()
	rv := objc.Send[NSArray](objc.ID(class.class), objc.Sel("new"))
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
func NewArrayWithArray(array []objectivec.IObject) NSArray {
	instance := getNSArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArray:"), objectivec.IObjectSliceToNSArray(array))
	return NSArrayFromID(rv)
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
func NewArrayWithArrayCopyItems(array []objectivec.IObject, flag bool) NSArray {
	instance := getNSArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArray:copyItems:"), objectivec.IObjectSliceToNSArray(array), flag)
	return NSArrayFromID(rv)
}

//
// # Return Value
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/init(coder:)
func NewArrayWithCoder(coder INSCoder) NSArray {
	instance := getNSArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSArrayFromID(rv)
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
func NewArrayWithObject(anObject objectivec.IObject) NSArray {
	rv := objc.Send[objc.ID](objc.ID(getNSArrayClass().class), objc.Sel("arrayWithObject:"), anObject)
	return NSArrayFromID(rv)
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
func NewArrayWithObjects(firstObj objectivec.IObject) NSArray {
	instance := getNSArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObjects:"), firstObj)
	return NSArrayFromID(rv)
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
func (a NSArray) InitWithArray(array []objectivec.IObject) NSArray {
	rv := objc.Send[NSArray](a.ID, objc.Sel("initWithArray:"), objectivec.IObjectSliceToNSArray(array))
	return rv
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
func (a NSArray) InitWithArrayCopyItems(array []objectivec.IObject, flag bool) NSArray {
	rv := objc.Send[NSArray](a.ID, objc.Sel("initWithArray:copyItems:"), objectivec.IObjectSliceToNSArray(array), flag)
	return rv
}

// Initializes a newly allocated array to include a given number of objects
// from a given C array.
//
// objects: A C array of objects.
//
// cnt: The number of values from the `objects` C array to include in the new
// array. This number will be the count of the new array—it must not be
// negative or greater than the number of elements in `objects`.
//
// # Return Value
// 
// A newly allocated array including the first `count` objects from `objects`.
// The returned object might be different than the original receiver.
//
// # Discussion
// 
// Elements are added to the new array in the same order they appear in
// `objects`, up to but not including index `count`.
// 
// After an immutable array has been initialized in this way, it can’t be
// modified.
// 
// This method is a designated initializer.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/init(objects:count:)-5odxv
func (a NSArray) InitWithObjectsCount(objects []objectivec.IObject, cnt uint) NSArray {
	rv := objc.Send[NSArray](a.ID, objc.Sel("initWithObjects:count:"), objc.CArray(objects), cnt)
	return rv
}

// Returns a Boolean value that indicates whether a given object is present in
// the array.
//
// anObject: An object to look for in the array.
//
// # Return Value
// 
// [true] if `anObject` is present in the array, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Starting at index `0`, each element of the array is checked for equality
// with `anObject` until a match is found or the end of the array is reached.
// Objects are considered equal if [isEqual(_:)] returns [true].
// 
// To determine if the array contains a particular instance of an object, you
// can test for identity rather than equality by calling the
// [IndexOfObjectIdenticalTo] method and comparing the return value to
// [NSNotFound].
//
// [NSNotFound]: https://developer.apple.com/documentation/Foundation/NSNotFound-9t5v2
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/contains(_:)
func (a NSArray) ContainsObject(anObject objectivec.IObject) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("containsObject:"), anObject)
	return rv
}

// Returns the object located at the specified index.
//
// index: An index within the bounds of the array.
//
// # Return Value
// 
// The object located at `index`.
//
// # Discussion
// 
// If `index` is beyond the end of the array (that is, if `index` is greater
// than or equal to the value returned by `count`), an [rangeException] is
// raised.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/object(at:)
func (a NSArray) ObjectAtIndex(index uint) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("objectAtIndex:"), index)
	return objectivec.Object{ID: rv}
}

// Returns the object at the specified index.
//
// idx: An index within the bounds of the array.
//
// # Return Value
// 
// The object located at `idx`.
//
// # Discussion
// 
// This method has the same behavior as the [ObjectAtIndex] method.
// 
// If `idx` is beyond the end of the array (that is, if `idx` is greater than
// or equal to the value returned by `count`), an [rangeException] is raised.
// 
// You shouldn’t need to call this method directly. Instead, this method is
// called when accessing an object by index using subscripting.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/subscript(_:)
func (a NSArray) ObjectAtIndexedSubscript(idx uint) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("objectAtIndexedSubscript:"), idx)
	return objectivec.Object{ID: rv}
}

// Returns an array containing the objects in the array at the indexes
// specified by a given index set.
//
// # Return Value
// 
// An array containing the objects in the array at the indexes specified by
// `indexes`.
//
// # Discussion
// 
// The returned objects are in the ascending order of their indexes in
// `indexes`, so that object in returned array with higher index in indexes
// will follow the object with smaller index in `indexes`.
// 
// Raises an [rangeException] if any location in `indexes` exceeds the bounds
// of the array, `indexes` is `nil`.
//
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/objects(at:)
func (a NSArray) ObjectsAtIndexes(indexes INSIndexSet) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("objectsAtIndexes:"), indexes)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns an enumerator object that lets you access each object in the array.
//
// # Return Value
// 
// An enumerator object that lets you access each object in the array, in
// order, from the element at the lowest index upwards.
//
// # Discussion
// 
// Returns an enumerator object that lets you access each object in the array,
// in order, starting with the element at index 0, as in:
// 
// # Special Considerations
// 
// When you use this method with mutable subclasses of [NSArray], you must not
// modify the array during enumeration.
// 
// It is more efficient to use the fast enumeration protocol (see
// [NSFastEnumeration]). Fast enumeration is available in macOS 10.5 and later
// and iOS 2.0 and later.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/objectEnumerator()
func (a NSArray) ObjectEnumerator() INSEnumerator {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("objectEnumerator"))
	return NSEnumeratorFromID(rv)
}

// Returns an enumerator object that lets you access each object in the array,
// in reverse order.
//
// # Return Value
// 
// An enumerator object that lets you access each object in the array, in
// order, from the element at the highest index down to the element at index
// `0`.
//
// # Discussion
// 
// When you use this method with mutable subclasses of [NSArray], you must not
// modify the array during enumeration.
// 
// It is more efficient to use the fast enumeration protocol (see
// [NSFastEnumeration]). Fast enumeration is available in macOS 10.5 and later
// and iOS 2.0 and later.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/reverseObjectEnumerator()
func (a NSArray) ReverseObjectEnumerator() INSEnumerator {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("reverseObjectEnumerator"))
	return NSEnumeratorFromID(rv)
}

// Returns the lowest index whose corresponding array value is equal to a
// given object.
//
// anObject: An object.
//
// # Return Value
// 
// The lowest index whose corresponding array value is equal to `anObject`. If
// none of the objects in the array is equal to `anObject`, returns
// [NSNotFound].
//
// # Discussion
// 
// Starting at index `0`, each element of the array is passed as an argument
// to an [isEqual(_:)] message sent to `anObject` until a match is found or
// the end of the array is reached. Objects are considered equal if
// [isEqual(_:)] (declared in the [NSObjectProtocol] protocol) returns [true].
//
// [NSObjectProtocol]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/index(of:)
func (a NSArray) IndexOfObject(anObject objectivec.IObject) uint {
	rv := objc.Send[uint](a.ID, objc.Sel("indexOfObject:"), anObject)
	return rv
}

// Returns the lowest index within a specified range whose corresponding array
// value is equal to a given object .
//
// anObject: An object.
//
// range: The range of indexes in the array within which to search for `anObject`.
//
// # Return Value
// 
// The lowest index within `range` whose corresponding array value is equal to
// `anObject`. If none of the objects within `range` is equal to `anObject`,
// returns [NSNotFound].
//
// # Discussion
// 
// Starting at `range.Location()`, each element of the array is passed as an
// argument to an [isEqual(_:)] message sent to `anObject` until a match is
// found or the end of the `range` is reached. Objects are considered equal if
// [isEqual(_:)] returns [true].
// 
// This method raises an [rangeException] exception if the `range` parameter
// represents a range that doesn’t exist in the array.
//
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/index(of:in:)
func (a NSArray) IndexOfObjectInRange(anObject objectivec.IObject, range_ NSRange) uint {
	rv := objc.Send[uint](a.ID, objc.Sel("indexOfObject:inRange:"), anObject, range_)
	return rv
}

// Returns the lowest index whose corresponding array value is identical to a
// given object.
//
// anObject: An object.
//
// # Return Value
// 
// The lowest index whose corresponding array value is identical to
// `anObject`. If none of the objects in the array is identical to `anObject`,
// returns [NSNotFound].
//
// # Discussion
// 
// Objects are considered identical if their object addresses are the same.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/indexOfObjectIdentical(to:)
func (a NSArray) IndexOfObjectIdenticalTo(anObject objectivec.IObject) uint {
	rv := objc.Send[uint](a.ID, objc.Sel("indexOfObjectIdenticalTo:"), anObject)
	return rv
}

// Returns the lowest index within a specified range whose corresponding array
// value is equal to a given object .
//
// anObject: An object.
//
// range: The range of indexes in the array within which to search for `anObject`.
//
// # Return Value
// 
// The lowest index within `range` whose corresponding array value is
// identical to `anObject`. If none of the objects within `range` is identical
// to `anObject`, returns [NSNotFound].
//
// # Discussion
// 
// Objects are considered identical if their object addresses are the same.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/indexOfObjectIdentical(to:in:)
func (a NSArray) IndexOfObjectIdenticalToInRange(anObject objectivec.IObject, range_ NSRange) uint {
	rv := objc.Send[uint](a.ID, objc.Sel("indexOfObjectIdenticalTo:inRange:"), anObject, range_)
	return rv
}

// Returns the index of the first object in the array that passes a test in a
// given block.
//
// predicate: The block to apply to elements in the array.
// 
// The block takes three arguments:
// 
// obj: The element in the array. idx: The index of the element in the array.
// stop: A reference to a Boolean value. The block can set the value to [true]
// to stop further enumeration of the array. If a block stops further
// enumeration, that block continues to run until it’s finished. The `stop`
// argument is an out-only argument. You should only ever set this Boolean to
// [true] within the block.
// 
// The block returns a Boolean value that indicates whether `obj` passed the
// test. Returning [true] will stop further processing of the array.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The lowest index whose corresponding value in the array passes the test
// specified by `predicate`. If no objects in the array pass the test, returns
// [NSNotFound].
//
// # Discussion
// 
// If the block parameter is `nil` this method will raise an exception.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/indexOfObject(passingTest:)
func (a NSArray) IndexOfObjectPassingTest(predicate unsafe.Pointer) uint {
	rv := objc.Send[uint](a.ID, objc.Sel("indexOfObjectPassingTest:"), predicate)
	return rv
}

// Returns the index of an object in the array that passes a test in a given
// block for a given set of enumeration options.
//
// opts: A bit mask that specifies the options for the enumeration (whether it
// should be performed concurrently and whether it should be performed in
// reverse order).
//
// predicate: The block to apply to elements in the array.
// 
// The block takes three arguments:
// 
// obj: The element in the array. idx: The index of the element in the array.
// stop: A reference to a Boolean value. The block can set the value to [true]
// to stop further enumeration of the array. If a block stops further
// enumeration, that block continues to run until it’s finished. When the
// [NSEnumerationConcurrent] enumeration option is specified, enumeration
// stops after all of the currently running blocks finish. The `stop` argument
// is an out-only argument. You should only ever set this Boolean to [true]
// within the block.
// 
// The block returns a Boolean value that indicates whether `obj` passed the
// test.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The index whose corresponding value in the array passes the test specified
// by `predicate` and `opts`. If the `opts` bit mask specifies reverse order,
// then the last item that matches is returned. Otherwise, the index of the
// first matching object is returned. If no objects in the array pass the
// test, returns [NSNotFound].
//
// # Discussion
// 
// By default, the enumeration starts with the first object and continues
// serially through the array to the last object. You can specify
// [NSEnumerationConcurrent] and/or [NSEnumerationReverse] as enumeration
// options to modify this behavior.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/indexOfObject(options:passingTest:)
func (a NSArray) IndexOfObjectWithOptionsPassingTest(opts NSEnumerationOptions, predicate unsafe.Pointer) uint {
	rv := objc.Send[uint](a.ID, objc.Sel("indexOfObjectWithOptions:passingTest:"), opts, predicate)
	return rv
}

// Returns the index, from a given set of indexes, of the first object in the
// array that passes a test in a given block for a given set of enumeration
// options.
//
// s: The indexes of the objects over which to enumerate.
//
// opts: A bit mask that specifies the options for the enumeration (whether it
// should be performed concurrently and whether it should be performed in
// reverse order).
//
// predicate: The block to apply to elements in the array.
// 
// The block takes three arguments:
// 
// obj: The element in the array. idx: The index of the element in the array.
// stop: A reference to a Boolean value. The block can set the value to [true]
// to stop further enumeration of the array. If a block stops further
// enumeration, that block continues to run until it’s finished. When the
// [NSEnumerationConcurrent] enumeration option is specified, enumeration
// stops after all of the currently running blocks finish. The `stop` argument
// is an out-only argument. You should only ever set this Boolean to [true]
// within the block.
// 
// The block returns a Boolean value that indicates whether `obj` passed the
// test.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The lowest index whose corresponding value in the array passes the test
// specified by `predicate`. If no objects in the array pass the test, returns
// [NSNotFound].
//
// # Discussion
// 
// By default, the enumeration starts with the first object and continues
// serially through the array to the last element specified by `indexSet`. You
// can specify [EnumerationConcurrent] and/or [EnumerationReverse] as
// enumeration options to modify this behavior.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/indexOfObject(at:options:passingTest:)
func (a NSArray) IndexOfObjectAtIndexesOptionsPassingTest(s INSIndexSet, opts NSEnumerationOptions, predicate unsafe.Pointer) uint {
	rv := objc.Send[uint](a.ID, objc.Sel("indexOfObjectAtIndexes:options:passingTest:"), s, opts, predicate)
	return rv
}

// Returns the indexes of objects in the array that pass a test in a given
// block.
//
// predicate: The block to apply to elements in the array.
// 
// The block takes three arguments:
// 
// obj: The element in the array. idx: The index of the element in the array.
// stop: A reference to a Boolean value. The block can set the value to [true]
// to stop further enumeration of the array. If a block stops further
// enumeration, that block continues to run until it’s finished. The `stop`
// argument is an out-only argument. You should only ever set this Boolean to
// [true] within the block.
// 
// The block returns a Boolean value that indicates whether `obj` passed the
// test.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The indexes whose corresponding values in the array pass the test specified
// by `predicate`. If no objects in the array pass the test, returns an empty
// index set.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/indexesOfObjects(passingTest:)
func (a NSArray) IndexesOfObjectsPassingTest(predicate unsafe.Pointer) INSIndexSet {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("indexesOfObjectsPassingTest:"), predicate)
	return NSIndexSetFromID(rv)
}

// Returns the indexes of objects in the array that pass a test in a given
// block for a given set of enumeration options.
//
// opts: A bit mask that specifies the options for the enumeration (whether it
// should be performed concurrently and whether it should be performed in
// reverse order).
//
// predicate: The block to apply to elements in the array.
// 
// The block takes three arguments:
// 
// obj: The element in the array. idx: The index of the element in the array.
// stop: A reference to a Boolean value. The block can set the value to [true]
// to stop further enumeration of the array. If a block stops further
// enumeration, that block continues to run until it’s finished. When the
// [NSEnumerationConcurrent] enumeration option is specified, enumeration
// stops after all of the currently running blocks finish. The `stop` argument
// is an out-only argument. You should only ever set this Boolean to [true]
// within the block.
// 
// The block returns a Boolean value that indicates whether `obj` passed the
// test.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The indexes whose corresponding values in the array pass the test specified
// by `predicate`. If no objects in the array pass the test, returns an empty
// index set.
//
// # Discussion
// 
// By default, the enumeration starts with the first object and continues
// serially through the array to the last object. You can specify
// [EnumerationConcurrent] and/or [EnumerationReverse] as enumeration options
// to modify this behavior.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/indexesOfObjects(options:passingTest:)
func (a NSArray) IndexesOfObjectsWithOptionsPassingTest(opts NSEnumerationOptions, predicate unsafe.Pointer) INSIndexSet {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("indexesOfObjectsWithOptions:passingTest:"), opts, predicate)
	return NSIndexSetFromID(rv)
}

// Returns the indexes, from a given set of indexes, of objects in the array
// that pass a test in a given block for a given set of enumeration options.
//
// s: The indexes of the objects over which to enumerate.
//
// opts: A bit mask that specifies the options for the enumeration (whether it
// should be performed concurrently and whether it should be performed in
// reverse order).
//
// predicate: The block to apply to elements in the array.
// 
// The block takes three arguments:
// 
// obj: The element in the array. idx: The index of the element in the array.
// stop: A reference to a Boolean value. The block can set the value to [true]
// to stop further enumeration of the array. If a block stops further
// enumeration, that block continues to run until it’s finished. When the
// [NSEnumerationConcurrent] enumeration option is specified, enumeration
// stops after all of the currently running blocks finish. The `stop` argument
// is an out-only argument. You should only ever set this Boolean to [true]
// within the block.
// 
// The block returns a Boolean value that indicates whether `obj` passed the
// test.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The indexes whose corresponding values in the array pass the test specified
// by `predicate`. If no objects in the array pass the test, returns an empty
// index set.
//
// # Discussion
// 
// By default, the enumeration starts with the first object and continues
// serially through the array to the last element specified by `indexSet`. You
// can specify [NSEnumerationConcurrent] and/or [NSEnumerationReverse] as
// enumeration options to modify this behavior.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/indexesOfObjects(at:options:passingTest:)
func (a NSArray) IndexesOfObjectsAtIndexesOptionsPassingTest(s INSIndexSet, opts NSEnumerationOptions, predicate unsafe.Pointer) INSIndexSet {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("indexesOfObjectsAtIndexes:options:passingTest:"), s, opts, predicate)
	return NSIndexSetFromID(rv)
}

// Returns the index, within a specified range, of an object compared with
// elements in the array using a given [NSComparator] block.
//
// obj: An object for which to search in the array.
// 
// If this value is `nil`, throws an [invalidArgumentException].
// //
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
//
// r: The range within the array to search for `obj`.
// 
// If `r` exceeds the bounds of the array (if the location plus length of the
// range is greater than the count of the array), throws an [rangeException].
// //
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// opts: Options for the search. For possible values, see
// [NSBinarySearchingOptions].
// 
// If you specify both [BinarySearchingFirstEqual] and
// [BinarySearchingLastEqual], throws an [NSInvalidArgumentException].
// //
// [NSBinarySearchingOptions]: https://developer.apple.com/documentation/Foundation/NSBinarySearchingOptions
//
// cmp: A comparator block used to compare the object `obj` with elements in the
// array.
// 
// If this value is [NULL], throws an [invalidArgumentException].
// //
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
//
// # Return Value
// 
// If the [BinarySearchingInsertionIndex] option is not specified:
// 
// - If the `obj` is found and neither [BinarySearchingFirstEqual] nor
// [BinarySearchingLastEqual] is specified, returns an arbitrary matching
// object’s index. - If the [BinarySearchingFirstEqual] option is also
// specified, returns the lowest index of equal objects. - If the
// [BinarySearchingLastEqual] option is also specified, returns the highest
// index of equal objects. - If the object is not found, returns [NSNotFound].
// 
// If the [BinarySearchingInsertionIndex] option is specified, returns the
// index at which you should insert `obj` in order to maintain a sorted array:
// 
// - If the `obj` is found and neither [BinarySearchingFirstEqual] nor
// [BinarySearchingLastEqual] is specified, returns any equal or one larger
// index than any matching object’s index. - If the
// [BinarySearchingFirstEqual] option is also specified, returns the lowest
// index of equal objects. - If the [BinarySearchingLastEqual] option is also
// specified, returns the highest index of equal objects. - If the object is
// not found, returns the index of the least greater object, or the index at
// the end of the array if the object is larger than all other elements.
//
// # Discussion
// 
// The elements in the array must have already been sorted using the
// comparator `cmp`. If the array is not sorted, the result is undefined.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/index(of:inSortedRange:options:usingComparator:)
func (a NSArray) IndexOfObjectInSortedRangeOptionsUsingComparator(obj objectivec.IObject, r NSRange, opts NSBinarySearchingOptions, cmp NSComparator) uint {
	rv := objc.Send[uint](a.ID, objc.Sel("indexOfObject:inSortedRange:options:usingComparator:"), obj, r, opts, cmp)
	return rv
}

// Executes a given closure or block using each object in the array, starting
// with the first object and continuing through the array to the last object.
//
// block: A closure or block to execute for each object in the array, taking three
// arguments:
// 
// - The object. - The index of the object in the array. - A reference to a
// Boolean value, which the closure can set to [true] in order to stop further
// enumeration of the array. If a closure stops further enumeration, that
// closure continues to run until it’s finished.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method executes synchronously. Values allocated within the block are
// deallocated after the block is executed.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/enumerateObjects(_:)
func (a NSArray) EnumerateObjectsUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](a.ID, objc.Sel("enumerateObjectsUsingBlock:"), block)
}

// Executes a given closure or block using each object in the array with the
// specified options.
//
// opts: The options for the enumeration. For possible values, see
// [NSEnumerationOptions].
// //
// [NSEnumerationOptions]: https://developer.apple.com/documentation/Foundation/NSEnumerationOptions
//
// block: A closure or block to execute for each object in the array, taking three
// arguments:
// 
// - The object. - The index of the object in the array. - A reference to a
// Boolean value, which the closure can set to [true] in order to stop further
// enumeration of the array. If a closure stops further enumeration, that
// closure continues to run until it’s finished. When the
// [EnumerationConcurrent] enumeration option is specified, enumeration stops
// after all of the currently running closures finish.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method executes synchronously. By default, the enumeration starts with
// the first object and continues serially through the array to the last
// object. You can specify [EnumerationConcurrent] and/or [EnumerationReverse]
// as enumeration options to modify this behavior.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/enumerateObjects(options:using:)
func (a NSArray) EnumerateObjectsWithOptionsUsingBlock(opts NSEnumerationOptions, block unsafe.Pointer) {
	objc.Send[objc.ID](a.ID, objc.Sel("enumerateObjectsWithOptions:usingBlock:"), opts, block)
}

// Executes a given block using the objects in the array at the specified
// indexes.
//
// s: The indexes of the objects over which to enumerate.
//
// opts: A bit mask that specifies the options for the enumeration (whether it
// should be performed concurrently and whether it should be performed in
// reverse order).
//
// block: The block to apply to elements in the array.
// 
// The block takes three arguments:
// 
// obj: The element in the array. idx: The index of the element in the array.
// stop: A reference to a Boolean value. The block can set the value to [true]
// to stop further enumeration of the array. If a block stops further
// enumeration, that block continues to run until it’s finished. When the
// [NSEnumerationConcurrent] enumeration option is specified, enumeration
// stops after all of the currently running blocks finish. The `stop` argument
// is an out-only argument. You should only ever set this Boolean to [true]
// within the block.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// By default, the enumeration starts with the first object and continues
// serially through the array to the last element specified by `indexSet`. You
// can specify [NSEnumerationConcurrent] and/or [NSEnumerationReverse] as
// enumeration options to modify this behavior.
// 
// This method executes synchronously.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/enumerateObjects(at:options:using:)
func (a NSArray) EnumerateObjectsAtIndexesOptionsUsingBlock(s INSIndexSet, opts NSEnumerationOptions, block unsafe.Pointer) {
	objc.Send[objc.ID](a.ID, objc.Sel("enumerateObjectsAtIndexes:options:usingBlock:"), s, opts, block)
}

// Returns the first object contained in the receiving array that’s equal to
// an object in another given array.
//
// otherArray: An array.
//
// # Return Value
// 
// Returns the first object contained in the receiving array that’s equal to
// an object in `otherArray`. If no such object is found, returns `nil`.
//
// # Discussion
// 
// This method uses [isEqual(_:)] to check for object equality.
//
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/firstObjectCommon(with:)
func (a NSArray) FirstObjectCommonWithArray(otherArray []objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("firstObjectCommonWithArray:"), objectivec.IObjectSliceToNSArray(otherArray))
	return objectivec.Object{ID: rv}
}

// Compares the receiving array to another array.
//
// otherArray: An array.
//
// # Return Value
// 
// [true] if the contents of `otherArray` are equal to the contents of the
// receiving array, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Two arrays have equal contents if they each hold the same number of objects
// and objects at a given index in each array satisfy the [isEqual(_:)] test.
//
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/isEqual(to:)
func (a NSArray) IsEqualToArray(otherArray []objectivec.IObject) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isEqualToArray:"), objectivec.IObjectSliceToNSArray(otherArray))
	return rv
}

// Returns a new array that is a copy of the receiving array with a given
// object added to the end.
//
// anObject: An object.
//
// # Return Value
// 
// A new array that is a copy of the receiving array with `anObject` added to
// the end.
//
// # Discussion
// 
// If `anObject` is `nil`, an [NSInvalidArgumentException] is raised.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/adding(_:)
func (a NSArray) ArrayByAddingObject(anObject objectivec.IObject) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("arrayByAddingObject:"), anObject)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns a new array that is a copy of the receiving array with the objects
// contained in another array added to the end.
//
// otherArray: An array.
//
// # Return Value
// 
// A new array that is a copy of the receiving array with the objects
// contained in `otherArray` added to the end.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/addingObjects(from:)
func (a NSArray) ArrayByAddingObjectsFromArray(otherArray []objectivec.IObject) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("arrayByAddingObjectsFromArray:"), objectivec.IObjectSliceToNSArray(otherArray))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Evaluates a given predicate against each object in the receiving array and
// returns a new array containing the objects for which the predicate returns
// true.
//
// predicate: The predicate against which to evaluate the receiving array’s elements.
//
// # Return Value
// 
// A new array containing the objects in the receiving array for which
// `predicate` returns [true].
// 
// Objects in the resulting array appear in the same order as they do in the
// receiver.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// For more details, see [Predicate Programming Guide].
//
// [Predicate Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Predicates/AdditionalChapters/Introduction.html#//apple_ref/doc/uid/TP40001789
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/filtered(using:)
func (a NSArray) FilteredArrayUsingPredicate(predicate INSPredicate) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("filteredArrayUsingPredicate:"), predicate)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns a new array containing the receiving array’s elements that fall
// within the limits specified by a given range.
//
// range: A range within the receiving array’s range of elements.
//
// # Return Value
// 
// A new array containing the receiving array’s elements that fall within
// the limits specified by `range`.
//
// # Discussion
// 
// If `range` isn’t within the receiving array’s range of elements, an
// [NSRangeException] is raised.
// 
// For example, the following code example creates an array containing the
// elements found in the first half of `wholeArray` (assuming `wholeArray`
// exists).
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/subarray(with:)
func (a NSArray) SubarrayWithRange(range_ NSRange) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("subarrayWithRange:"), range_)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns a new array that lists the receiving array’s elements in
// ascending order as defined by the comparison function `comparator`.
//
// # Discussion
// 
// The new array contains references to the receiving array’s elements, not
// copies of them.
// 
// The comparison function is used to compare two elements at a time and
// should return [NSOrderedAscending] if the first element is smaller than the
// second, [NSOrderedDescending] if the first element is larger than the
// second, and [NSOrderedSame] if the elements are equal. Each time the
// comparison function is called, it’s passed `context` as its third
// argument. This allows the comparison to be based on some outside parameter,
// such as whether character sorting is case-sensitive or case-insensitive.
// 
// Given `anArray` (an array of [NSNumber] objects) and a comparison function
// of this type:
// 
// A sorted version of `anArray` is created in this way:
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/sortedArray(_:context:)
func (a NSArray) SortedArrayUsingFunctionContext(comparator objectivec.IObject, context unsafe.Pointer) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("sortedArrayUsingFunction:context:"), comparator, context)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns a new array that lists the receiving array’s elements in
// ascending order as defined by the comparison function `comparator`.
//
// # Discussion
// 
// The new array contains references to the receiving array’s elements, not
// copies of them.
// 
// This method is similar to [SortedArrayUsingFunctionContext], except that it
// uses the supplied hint to speed the sorting process. When you know the
// array is nearly sorted, this method is faster than
// [SortedArrayUsingFunctionContext]. If you sorted a large array ([N]
// entries) once, and you don’t change it much ([P] additions and deletions,
// where [P] is much smaller than [N]), then you can reuse the work you did in
// the original sort by conceptually doing a merge sort between the [N]
// “old” items and the [P] “new” items.
// 
// To obtain an appropriate hint, use [SortedArrayHint]. You should obtain
// this hint when the original array has been sorted, and keep hold of it
// until you need it, after the array has been modified. The hint is computed
// by [SortedArrayHint] in `O(N)` (where [N] is the number of items). This
// assumes that items in the array implement a `-hash` method. Given a
// suitable hint, and assuming that the hash function is a “good” hash
// function, -[SortedArrayUsingFunctionContextHint] sorts the array in
// `O(P*LOG(P)+N)` where [P] is the number of adds or deletes. This is an
// improvement over the un-hinted sort, `O(N*LOG(N))`, when [P] is small.
// 
// The hint is simply an array of size [N] containing the [N] hashes. To
// re-sort you need internally to create a map table mapping a hash to the
// index. Using this map table on the new array, you can get a first guess for
// the indexes, and then sort that. For example, a sorted array {A, B, D, E,
// F} with corresponding hash values {25, 96, 78, 32, 17}, may be subject to
// small changes that result in contents {E, A, C, B, F}. The mapping table
// maps the hashes {25, 96, 78, 32, 17} to the indexes {#0, #1, #2, #3, #4}.
// If the hashes for {E, A, C, B, F} are {32, 25, 99, 96, 17}, then by using
// the mapping table you can get a first order sort {#3, #0, ?, #1, #4}, so
// therefore create an initial semi-sorted array {A, B, E, F}, and then
// perform a cheap merge sort with {C} that yields {A, B, C, E, F}.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/sortedArray(_:context:hint:)
func (a NSArray) SortedArrayUsingFunctionContextHint(comparator objectivec.IObject, context unsafe.Pointer, hint INSData) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("sortedArrayUsingFunction:context:hint:"), comparator, context, hint)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns a copy of the receiving array sorted as specified by a given array
// of sort descriptors.
//
// sortDescriptors: An array of [NSSortDescriptor] objects.
//
// # Return Value
// 
// A copy of the receiving array sorted as specified by `sortDescriptors`.
//
// # Discussion
// 
// The first descriptor specifies the primary key path to be used in sorting
// the receiving array’s contents. Any subsequent descriptors are used to
// further refine sorting of objects with duplicate values. See
// [NSSortDescriptor] for additional information.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/sortedArray(using:)-82wi1
func (a NSArray) SortedArrayUsingDescriptors(sortDescriptors []NSSortDescriptor) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("sortedArrayUsingDescriptors:"), objectivec.IObjectSliceToNSArray(sortDescriptors))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns an array that lists the receiving array’s elements in ascending
// order, as determined by the comparison method specified by a given
// selector.
//
// comparator: A selector that identifies the method to use to compare two elements at a
// time. The method should return [NSOrderedAscending] if the receiving array
// is smaller than the argument, [NSOrderedDescending] if the receiving array
// is larger than the argument, and [NSOrderedSame] if they are equal.
//
// # Return Value
// 
// An array that lists the receiving array’s elements in ascending order, as
// determined by the comparison method specified by the selector `comparator`.
//
// # Discussion
// 
// The new array contains references to the receiving array’s elements, not
// copies of them.
// 
// The `comparator` message is sent to each object in the array and has as its
// single argument another object in the array.
// 
// For example, an array of [NSString] objects can be sorted by using the
// [CaseInsensitiveCompare] method declared in the [NSString] class. Assuming
// `anArray` exists, a sorted version of the array can be created in this way:
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/sortedArray(using:)-9nhh9
func (a NSArray) SortedArrayUsingSelector(comparator objc.SEL) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("sortedArrayUsingSelector:"), comparator)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns an array that lists the receiving array’s elements in ascending
// order, as determined by the comparison method specified by a given
// [NSComparator] block.
//
// cmptr: A comparator block.
//
// # Return Value
// 
// An array that lists the receiving array’s elements in ascending order, as
// determined by the comparison method specified `cmptr`.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/sortedArray(comparator:)
func (a NSArray) SortedArrayUsingComparator(cmptr NSComparator) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("sortedArrayUsingComparator:"), cmptr)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns an array that lists the receiving array’s elements in ascending
// order, as determined by the comparison method specified by a given
// [NSComparator] block.
//
// opts: A bit mask that specifies the options for the sort (whether it should be
// performed concurrently and whether it should be performed stably).
//
// cmptr: A comparator block.
//
// # Return Value
// 
// An array that lists the receiving array’s elements in ascending order, as
// determined by the comparison method specified `cmptr`.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/sortedArray(options:usingComparator:)
func (a NSArray) SortedArrayWithOptionsUsingComparator(opts NSSortOptions, cmptr NSComparator) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("sortedArrayWithOptions:usingComparator:"), opts, cmptr)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Constructs and returns an [NSString] object that is the result of
// interposing a given separator between the elements of the array.
//
// separator: The string to interpose between the elements of the array.
//
// # Return Value
// 
// An [NSString] object that is the result of interposing `separator` between
// the elements of the array. If the array has no elements, returns an
// [NSString] object representing an empty string.
//
// # Discussion
// 
// For example, this code excerpt writes “`here be dragons`” to the
// console:
// 
// # Special Considerations
// 
// Each element in the array must handle `description`.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/componentsJoined(by:)
func (a NSArray) ComponentsJoinedByString(separator string) string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("componentsJoinedByString:"), objc.String(separator))
	return NSStringFromID(rv).String()
}

// Returns a string that represents the contents of the array, formatted as a
// property list.
//
// locale: An [NSLocale] object or an [NSDictionary] object that specifies options
// used for formatting each of the array’s elements (where recognized).
// Specify `nil` if you don’t want the elements formatted.
//
// # Return Value
// 
// A string that represents the contents of the array, formatted as a property
// list.
//
// # Discussion
// 
// For a description of how `locale` is applied to each element in the
// receiving array, see [DescriptionWithLocaleIndent].
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/description(withLocale:)
func (a NSArray) DescriptionWithLocale(locale objectivec.IObject) string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("descriptionWithLocale:"), locale)
	return NSStringFromID(rv).String()
}

// Returns a string that represents the contents of the array, formatted as a
// property list.
//
// locale: An [NSLocale] object or an [NSDictionary] object that specifies options
// used for formatting each of the array’s elements (where recognized).
// Specify `nil` if you don’t want the elements formatted.
//
// level: A level of indent, to make the output more readable: set `level` to `0` to
// use four spaces to indent, or `1` to indent the output with a tab
// character.
//
// # Return Value
// 
// A string that represents the contents of the array, formatted as a property
// list.
//
// # Discussion
// 
// The returned [NSString] object contains the string representations of each
// of the array’s elements, in order, from first to last. To obtain the
// string representation of a given element, [DescriptionWithLocaleIndent]
// proceeds as follows:
// 
// - If the element is an [NSString] object, it is used as is. - If the
// element responds to [DescriptionWithLocaleIndent], that method is invoked
// to obtain the element’s string representation. - If the element responds
// to [DescriptionWithLocale], that method is invoked to obtain the
// element’s string representation. - If none of the above conditions is
// met, the element’s string representation is obtained by invoking its
// [Description] method.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/description(withLocale:indent:)
func (a NSArray) DescriptionWithLocaleIndent(locale objectivec.IObject, level uint) string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("descriptionWithLocale:indent:"), locale, level)
	return NSStringFromID(rv).String()
}

// Returns an array containing all the pathname elements in the receiving
// array that have filename extensions from a given array.
//
// filterTypes: An array of [NSString] objects containing filename extensions. The
// extensions should not include the dot (”.”) character.
//
// # Return Value
// 
// An array containing all the pathname elements in the receiving array that
// have filename extensions from the `filterTypes` array.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/pathsMatchingExtensions(_:)
func (a NSArray) PathsMatchingExtensions(filterTypes []string) []string {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("pathsMatchingExtensions:"), objectivec.StringSliceToNSArray(filterTypes))
	return objc.ConvertSliceToStrings(rv)
}

// Raises an exception.
//
// observer: The object to remove as an observer.
//
// indexes: The index set.
//
// keyPath: A key-path, relative to the array, for which `observer` is registered to
// receive KVO change notifications. This value must not be `nil`.
//
// context: The context passed to the notifications.
//
// # Discussion
// 
// [NSArray] objects are not observable, so this method raises an exception
// when invoked on an [NSArray] object. Instead of observing an array, observe
// the to-many relationship for which the array is the collection of related
// objects.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/removeObserver(_:fromObjectsAt:forKeyPath:context:)
func (a NSArray) RemoveObserverFromObjectsAtIndexesForKeyPathContext(observer objectivec.Object, indexes INSIndexSet, keyPath string, context unsafe.Pointer) {
	objc.Send[objc.ID](a.ID, objc.Sel("removeObserver:fromObjectsAtIndexes:forKeyPath:context:"), observer, indexes, objc.String(keyPath), context)
}

// Registers an observer to receive key value observer notifications for the
// specified key-path relative to the objects at the indexes.
//
// observer: The observer.
//
// indexes: The index set.
//
// keyPath: The key path, relative to the array, to be observed.
//
// options: The options to be included in the notification.
//
// context: The context passed to the notifications.
//
// # Discussion
// 
// The `options` determine what is included in the notifications, and the
// `context` is passed in the notifications.
// 
// This is not merely a convenience method; invoking this method is
// potentially much faster than repeatedly invoking
// [addObserver(_:forKeyPath:options:context:)].
//
// [addObserver(_:forKeyPath:options:context:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/addObserver(_:forKeyPath:options:context:)
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/addObserver(_:toObjectsAt:forKeyPath:options:context:)
func (a NSArray) AddObserverToObjectsAtIndexesForKeyPathOptionsContext(observer objectivec.Object, indexes INSIndexSet, keyPath string, options uint, context unsafe.Pointer) {
	objc.Send[objc.ID](a.ID, objc.Sel("addObserver:toObjectsAtIndexes:forKeyPath:options:context:"), observer, indexes, objc.String(keyPath), options, context)
}

// Removes `anObserver` from all key value observer notifications associated
// with the specified `keyPath` relative to the array’s objects at
// `indexes`.
//
// observer: The observer.
//
// indexes: The index set.
//
// keyPath: The key path, relative to the array, to be observed.
//
// # Discussion
// 
// This is not merely a convenience method; invoking this method is
// potentially much faster than repeatedly invoking
// [removeObserver(_:forKeyPath:)].
//
// [removeObserver(_:forKeyPath:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/removeObserver(_:forKeyPath:)
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/removeObserver(_:fromObjectsAt:forKeyPath:)
func (a NSArray) RemoveObserverFromObjectsAtIndexesForKeyPath(observer objectivec.Object, indexes INSIndexSet, keyPath string) {
	objc.Send[objc.ID](a.ID, objc.Sel("removeObserver:fromObjectsAtIndexes:forKeyPath:"), observer, indexes, objc.String(keyPath))
}

// Returns a new array that lists this array’s elements in a random order.
//
// # Return Value
// 
// A new array that lists this array’s elements in a random order.
//
// # Discussion
// 
// Calling this method is equivalent to calling the
// [ShuffledArrayWithRandomSource] method and passing the system
// [sharedRandom()] random source. To influence the random shuffling or to be
// able to deterministically reproduce a series of shuffles, create your own
// [GKRandomSource] object.
//
// [GKRandomSource]: https://developer.apple.com/documentation/GameplayKit/GKRandomSource
// [sharedRandom()]: https://developer.apple.com/documentation/GameplayKit/GKRandomSource/sharedRandom()
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/shuffled()
func (a NSArray) ShuffledArray() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("shuffledArray"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns a new array that lists this array’s elements in a random order,
// using the specified random source.
//
// randomSource: A GameplayKit random source object.
//
// randomSource is a [gameplaykit.GKRandomSource].
//
// # Return Value
// 
// A new array that lists this array’s elements in a random order.
//
// # Discussion
// 
// Use the `randomSource` parameter to influence the random shuffling. For
// example, to reproduce a series of shuffles for testing, you can create a
// [GKARC4RandomSource] object using the [seed] value of a previously used
// random source.
// 
// This method is equivalent to the [GKRandomSource] method
// [arrayByShufflingObjects(in:)], but as an [NSArray] method it preserves
// generic type parameters.
//
// [GKARC4RandomSource]: https://developer.apple.com/documentation/GameplayKit/GKARC4RandomSource
// [GKRandomSource]: https://developer.apple.com/documentation/GameplayKit/GKRandomSource
// [arrayByShufflingObjects(in:)]: https://developer.apple.com/documentation/GameplayKit/GKRandomSource/arrayByShufflingObjects(in:)
// [seed]: https://developer.apple.com/documentation/GameplayKit/GKARC4RandomSource/seed
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/shuffled(using:)
func (a NSArray) ShuffledArrayWithRandomSource(randomSource objectivec.IObject) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("shuffledArrayWithRandomSource:"), randomSource)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

//
// # Return Value
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/init(coder:)
func (a NSArray) InitWithCoder(coder INSCoder) NSArray {
	rv := objc.Send[NSArray](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSArray/init(contentsOf:error:)
func (a NSArray) InitWithContentsOfURLError(url INSURL) (NSArray, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSArray{}, NSErrorFrom(errorPtr)
	}
	return NSArrayFromID(rv), nil

}

//
// See: https://developer.apple.com/documentation/Foundation/NSArray/write(to:)
func (a NSArray) WriteToURLError(url INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("writeToURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeToURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Creates a new array by applying a difference object to an existing array.
//
// # Discussion
// 
// The following example computes the difference between two arrays, then
// applies the difference to create an array that duplicates the original:
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/arrayByApplyingDifference:
func (a NSArray) ArrayByApplyingDifference(difference INSOrderedCollectionDifference) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("arrayByApplyingDifference:"), difference)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
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
func (a NSArray) CountByEnumeratingWithStateObjectsCount(state NSFastEnumerationState, buffer []objectivec.IObject, len_ uint) uint {
	rv := objc.Send[uint](a.ID, objc.Sel("countByEnumeratingWithState:objects:count:"), state, objc.CArray(buffer), len_)
	return rv
}

// Compares two arrays to create a difference object that represents the
// changes between them.
//
// # Discussion
// 
// The difference method creates the difference object by comparing objects
// within the arrays with the `` method.
// 
// The following example computes the difference between two arrays:
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/differenceFromArray:
func (a NSArray) DifferenceFromArray(other []objectivec.IObject) INSOrderedCollectionDifference {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("differenceFromArray:"), objectivec.IObjectSliceToNSArray(other))
	return NSOrderedCollectionDifferenceFromID(rv)
}

// Compares two arrays, with options, to create a difference object that
// represents the changes between them.
//
// # Discussion
// 
// The difference method creates the difference object by comparing objects
// within the arrays with the `` method.
// 
// The options allow you to choose to omit insertion or removal references to
// the change objects within the difference object. You can also choose to
// infer moves when computing the difference, which provides an
// [AssociatedIndex] within the change objects that indicates the index in the
// array where the object moved from.
// 
// The following example computes the difference between two arrays, inferring
// moves between them:
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/differenceFromArray:withOptions:
func (a NSArray) DifferenceFromArrayWithOptions(other []objectivec.IObject, options NSOrderedCollectionDifferenceCalculationOptions) INSOrderedCollectionDifference {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("differenceFromArray:withOptions:"), objectivec.IObjectSliceToNSArray(other), options)
	return NSOrderedCollectionDifferenceFromID(rv)
}

// Compares two arrays, using the provided block and with options, to create a
// difference object that represents the changes between them.
//
// # Discussion
// 
// The options allow you to choose to omit insertion or removal references to
// the change objects within the difference object’s changes. Don’t use
// the option [OrderedCollectionDifferenceCalculationInferMoves] when
// providing a block for the equivalence test. The changes returned in the
// difference object don’t include valid values for [AssociatedIndex].
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/differenceFromArray:withOptions:usingEquivalenceTest:
func (a NSArray) DifferenceFromArrayWithOptionsUsingEquivalenceTest(other []objectivec.IObject, options NSOrderedCollectionDifferenceCalculationOptions, block bool) INSOrderedCollectionDifference {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("differenceFromArray:withOptions:usingEquivalenceTest:"), objectivec.IObjectSliceToNSArray(other), options, block)
	return NSOrderedCollectionDifferenceFromID(rv)
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (a NSArray) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Copies references to objects contained in the array that fall within the
// specified range to `aBuffer`.
//
// objects: A C array of objects of size at least the length of the range specified by
// `aRange`.
//
// range: A range within the bounds of the array.
// 
// If the location plus the length of the range is greater than the count of
// the array, this method raises an [rangeException].
// //
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// # Discussion
// 
// The method copies into `aBuffer` references to objects in the array in the
// range specified by `aRange`; the size of the buffer must therefore be at
// least the length of the range multiplied by the size of an object
// reference, as shown in the following example (this is solely for
// illustration—you should typically not create a buffer simply to iterate
// over the contents of an array):
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/getObjects:range:
func (a NSArray) GetObjectsRange(objects []objectivec.IObject, range_ NSRange) {
	objc.Send[objc.ID](a.ID, objc.Sel("getObjects:range:"), objectivec.IObjectSliceToNSArray(objects), range_)
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
func (a NSArray) InitWithObjects(firstObj objectivec.IObject) NSArray {
	rv := objc.Send[NSArray](a.ID, objc.Sel("initWithObjects:"), firstObj)
	return rv
}

// Sends to each object in the array the message identified by a given
// selector, starting with the first object and continuing through the array
// to the last object.
//
// aSelector: A selector that identifies the message to send to the objects in the array.
// The method must not take any arguments, and must not have the side effect
// of modifying the receiving array.
//
// # Discussion
// 
// This method raises an [NSInvalidArgumentException] if `aSelector` is
// [NULL].
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/makeObjectsPerformSelector:
func (a NSArray) MakeObjectsPerformSelector(aSelector objc.SEL) {
	objc.Send[objc.ID](a.ID, objc.Sel("makeObjectsPerformSelector:"), aSelector)
}

// Sends the `aSelector` message to each object in the array, starting with
// the first object and continuing through the array to the last object.
//
// aSelector: A selector that identifies the message to send to the objects in the array.
// The method must take a single argument of type id, and must not have the
// side effect of modifying the receiving array.
//
// argument: The object to send as the argument to each invocation of the `aSelector`
// method.
//
// # Discussion
// 
// This method raises an [NSInvalidArgumentException] if `aSelector` is
// [NULL].
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/makeObjectsPerformSelector:withObject:
func (a NSArray) MakeObjectsPerformSelectorWithObject(aSelector objc.SEL, argument objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("makeObjectsPerformSelector:withObject:"), aSelector, argument)
}

// Creates and returns an array that includes a given number of objects from a
// given C array.
//
// objects: A C array of objects.
//
// cnt: The number of values from the `objects` C array to include in the new
// array. This number will be the count of the new array—it must not be
// negative or greater than the number of elements in `objects`.
//
// # Return Value
// 
// A new array including the first `count` objects from `objects`.
//
// # Discussion
// 
// Elements are added to the new array in the same order they appear in
// `objects`, up to but not including index `count`. For example:
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/init(objects:count:)-7dct1
func (_NSArrayClass NSArrayClass) ArrayWithObjectsCount(objects []objectivec.IObject, cnt uint) NSArray {
	rv := objc.Send[objc.ID](objc.ID(_NSArrayClass.class), objc.Sel("arrayWithObjects:count:"), objc.CArray(objects), cnt)
	return NSArrayFromID(rv)
}

// Creates and returns an empty array.
//
// # Return Value
// 
// An empty array.
//
// # Discussion
// 
// This method is used by mutable subclasses of [NSArray].
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/array
func (_NSArrayClass NSArrayClass) Array() NSArray {
	rv := objc.Send[objc.ID](objc.ID(_NSArrayClass.class), objc.Sel("array"))
	return NSArrayFromID(rv)
}

// Creates and returns an array containing the objects in another given array.
//
// array: An array.
//
// # Return Value
// 
// An array containing the objects in `anArray`.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/arrayWithArray:
func (_NSArrayClass NSArrayClass) ArrayWithArray(array []objectivec.IObject) NSArray {
	rv := objc.Send[objc.ID](objc.ID(_NSArrayClass.class), objc.Sel("arrayWithArray:"), objectivec.IObjectSliceToNSArray(array))
	return NSArrayFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSArray/arrayWithContentsOfURL:error:
func (_NSArrayClass NSArrayClass) ArrayWithContentsOfURLError(url INSURL) ([]objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](objc.ID(_NSArrayClass.class), objc.Sel("arrayWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	}), nil

}

// Creates and returns an array containing the objects in the argument list.
//
// firstObj: The first object for the array.
//
// # Return Value
// 
// An array containing the objects in the argument list.
//
// # Discussion
// 
// Pass comma-separated list of trailing variadic arguments as additional
// objects, ending with `nil`.
// 
// The following code example creates an array containing three different
// types of element:
// 
// Alternatively, you can use array literal syntax in Objective-C or Swift to
// create an array containing given objects:
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/arrayWithObjects:
func (_NSArrayClass NSArrayClass) ArrayWithObjects(firstObj objectivec.IObject) NSArray {
	rv := objc.Send[objc.ID](objc.ID(_NSArrayClass.class), objc.Sel("arrayWithObjects:"), firstObj)
	return NSArrayFromID(rv)
}

// The number of objects in the array.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/count
func (a NSArray) Count() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("count"))
	return rv
}

// The first object in the array.
//
// # Discussion
// 
// If the array is empty, returns `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/firstObject
func (a NSArray) FirstObject() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("firstObject"))
	return objectivec.Object{ID: rv}
}

// The last object in the array.
//
// # Discussion
// 
// If the array is empty, returns `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/lastObject
func (a NSArray) LastObject() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("lastObject"))
	return objectivec.Object{ID: rv}
}

// Analyzes the array and returns a “hint” that speeds the sorting of the
// array when the hint is supplied to [SortedArrayUsingFunctionContextHint].
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/sortedArrayHint
func (a NSArray) SortedArrayHint() INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sortedArrayHint"))
	return NSDataFromID(objc.ID(rv))
}

// A string that represents the contents of the array, formatted as a property
// list.
//
// See: https://developer.apple.com/documentation/Foundation/NSArray/description
func (a NSArray) Description() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("description"))
	return NSStringFromID(rv).String()
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSMutableCopying
			

			// Protocol methods for NSSecureCoding
			

