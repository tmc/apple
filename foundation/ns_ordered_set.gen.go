// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSOrderedSet] class.
var (
	_NSOrderedSetClass     NSOrderedSetClass
	_NSOrderedSetClassOnce sync.Once
)

func getNSOrderedSetClass() NSOrderedSetClass {
	_NSOrderedSetClassOnce.Do(func() {
		_NSOrderedSetClass = NSOrderedSetClass{class: objc.GetClass("NSOrderedSet")}
	})
	return _NSOrderedSetClass
}

// GetNSOrderedSetClass returns the class object for NSOrderedSet.
func GetNSOrderedSetClass() NSOrderedSetClass {
	return getNSOrderedSetClass()
}

type NSOrderedSetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSOrderedSetClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSOrderedSetClass) Alloc() NSOrderedSet {
	rv := objc.Send[NSOrderedSet](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A static, ordered collection of unique objects.
//
// # Overview
//
// [NSOrderedSet] declares the programmatic interface for static sets of
// distinct objects. You establish a static set’s entries when it’s
// created, and thereafter the entries can’t be modified.
// [NSMutableOrderedSet], on the other hand, declares a programmatic interface
// for dynamic sets of distinct objects. A dynamic—or mutable—set allows
// the addition and deletion of entries at any time, automatically allocating
// memory as needed.
//
// You can use ordered sets as an alternative to arrays when the order of
// elements is important and performance in testing whether an object is
// contained in the set is a consideration—testing for membership of an
// array is slower than testing for membership of a set.
//
// # Initializing an Ordered Set
//
//   - [NSOrderedSet.InitWithArray]: Initializes a newly allocated set with the objects that are contained in a given array.
//   - [NSOrderedSet.InitWithArrayCopyItems]: Initializes a newly allocated set with the objects that are contained in a given array, optionally copying the items.
//   - [NSOrderedSet.InitWithArrayRangeCopyItems]: Initializes a newly allocated set with the objects that are contained in the specified range of an array, optionally copying the items.
//   - [NSOrderedSet.InitWithObject]: Initializes a new ordered set with the object.
//   - [NSOrderedSet.InitWithObjectsCount]: Initializes a newly allocated set with a specified number of objects from a given C array of objects.
//   - [NSOrderedSet.InitWithOrderedSet]: Initializes a new ordered set with the contents of a set.
//   - [NSOrderedSet.InitWithOrderedSetCopyItems]: Initializes a new ordered set with the contents of a set, optionally copying the items.
//   - [NSOrderedSet.InitWithOrderedSetRangeCopyItems]: Initializes a new ordered set with the contents of an ordered set, optionally copying the items.
//   - [NSOrderedSet.InitWithSet]: Initializes a new ordered set with the contents of a set.
//   - [NSOrderedSet.InitWithSetCopyItems]: Initializes a new ordered set with the contents of a set, optionally copying the objects in the set.
//
// # Counting Entries
//
//   - [NSOrderedSet.Count]: The number of members in the set.
//
// # Accessing Set Members
//
//   - [NSOrderedSet.ContainsObject]: Returns a Boolean value that indicates whether a given object is present in the ordered set.
//   - [NSOrderedSet.EnumerateObjectsAtIndexesOptionsUsingBlock]: Executes a given block using the objects in the ordered set at the specified indexes.
//   - [NSOrderedSet.EnumerateObjectsUsingBlock]: Executes a given block using each object in the ordered set.
//   - [NSOrderedSet.EnumerateObjectsWithOptionsUsingBlock]: Executes a given block using each object in the set, using the specified enumeration options.
//   - [NSOrderedSet.FirstObject]: The first object in the ordered set.
//   - [NSOrderedSet.LastObject]: The last object in the ordered set.
//   - [NSOrderedSet.ObjectAtIndex]: Returns the object at the specified index of the set.
//   - [NSOrderedSet.ObjectAtIndexedSubscript]: Returns the object at the specified index of the set.
//   - [NSOrderedSet.ObjectsAtIndexes]: Returns the objects in the ordered set at the specified indexes.
//   - [NSOrderedSet.IndexOfObject]: Returns the index of the specified object.
//   - [NSOrderedSet.IndexOfObjectInSortedRangeOptionsUsingComparator]: Returns the index, within a specified range, of an object compared with elements in the ordered set using a given NSComparator block.
//   - [NSOrderedSet.IndexOfObjectAtIndexesOptionsPassingTest]: Returns the index, from a given set of indexes, of the object in the ordered set that passes a test in a given block for a given set of enumeration options.
//   - [NSOrderedSet.IndexOfObjectPassingTest]: Returns the index of the object in the ordered set that passes a test in a given block.
//   - [NSOrderedSet.IndexOfObjectWithOptionsPassingTest]: Returns the index of an object in the ordered set that passes a test in a given block for a given set of enumeration options.
//   - [NSOrderedSet.IndexesOfObjectsAtIndexesOptionsPassingTest]: Returns the index, from a given set of indexes, of the object in the ordered set that passes a test in a given block for a given set of enumeration options.
//   - [NSOrderedSet.IndexesOfObjectsPassingTest]: Returns the index of the object in the ordered set that passes a test in a given block.
//   - [NSOrderedSet.IndexesOfObjectsWithOptionsPassingTest]: Returns the index of an object in the ordered set that passes a test in a given block for a given set of enumeration options.
//   - [NSOrderedSet.ObjectEnumerator]: Returns an enumerator object that lets you access each object in the ordered set.
//   - [NSOrderedSet.ReverseObjectEnumerator]: Returns an enumerator object that lets you access each object in the ordered set.
//   - [NSOrderedSet.ReversedOrderedSet]: An ordered set in the reverse order.
//
// # Comparing Sets
//
//   - [NSOrderedSet.IsEqualToOrderedSet]: Compares the receiving ordered set to another ordered set.
//   - [NSOrderedSet.IntersectsOrderedSet]: Returns a Boolean value that indicates whether at least one object in the receiving ordered set is also present in another given ordered set.
//   - [NSOrderedSet.IntersectsSet]: Returns a Boolean value that indicates whether at least one object in the receiving ordered set is also present in another given set.
//   - [NSOrderedSet.IsSubsetOfOrderedSet]: Returns a Boolean value that indicates whether every object in the receiving ordered set is also present in another given ordered set.
//   - [NSOrderedSet.IsSubsetOfSet]: Returns a Boolean value that indicates whether every object in the receiving ordered set is also present in another given set.
//
// # Creating a Sorted Array
//
//   - [NSOrderedSet.SortedArrayUsingDescriptors]: Returns an array of the ordered set’s elements sorted as specified by a given array of sort descriptors.
//   - [NSOrderedSet.SortedArrayUsingComparator]: Returns an array that lists the receiving ordered set’s elements in ascending order, as determined by the comparison method specified by a given [NSComparator] block
//   - [NSOrderedSet.SortedArrayWithOptionsUsingComparator]: Returns an array that lists the receiving ordered set’s elements in ascending order, as determined by the comparison method specified by a given [NSComparator] block.
//
// # Filtering Ordered Sets
//
//   - [NSOrderedSet.FilteredOrderedSetUsingPredicate]: Evaluates a given predicate against each object in the receiving ordered set and returns a new ordered set containing the objects for which the predicate returns true.
//
// # Describing a Set
//
//   - [NSOrderedSet.Description]: A string that represents the contents of the ordered set, formatted as a property list.
//   - [NSOrderedSet.DescriptionWithLocale]: Returns a string that represents the contents of the ordered set, formatted as a property list.
//   - [NSOrderedSet.DescriptionWithLocaleIndent]: Returns a string that represents the contents of the ordered set, formatted as a property list.
//
// # Converting Other Collections
//
//   - [NSOrderedSet.Array]: A representation of the ordered set as an array.
//   - [NSOrderedSet.Set]: A representation of the set containing the contents of the ordered set.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet
type NSOrderedSet struct {
	objectivec.Object
}

// NSOrderedSetFromID constructs a [NSOrderedSet] from an objc.ID.
//
// A static, ordered collection of unique objects.
func NSOrderedSetFromID(id objc.ID) NSOrderedSet {
	return NSOrderedSet{objectivec.Object{ID: id}}
}

// NOTE: NSOrderedSet adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSOrderedSet] class.
//
// # Initializing an Ordered Set
//
//   - [INSOrderedSet.InitWithArray]: Initializes a newly allocated set with the objects that are contained in a given array.
//   - [INSOrderedSet.InitWithArrayCopyItems]: Initializes a newly allocated set with the objects that are contained in a given array, optionally copying the items.
//   - [INSOrderedSet.InitWithArrayRangeCopyItems]: Initializes a newly allocated set with the objects that are contained in the specified range of an array, optionally copying the items.
//   - [INSOrderedSet.InitWithObject]: Initializes a new ordered set with the object.
//   - [INSOrderedSet.InitWithObjectsCount]: Initializes a newly allocated set with a specified number of objects from a given C array of objects.
//   - [INSOrderedSet.InitWithOrderedSet]: Initializes a new ordered set with the contents of a set.
//   - [INSOrderedSet.InitWithOrderedSetCopyItems]: Initializes a new ordered set with the contents of a set, optionally copying the items.
//   - [INSOrderedSet.InitWithOrderedSetRangeCopyItems]: Initializes a new ordered set with the contents of an ordered set, optionally copying the items.
//   - [INSOrderedSet.InitWithSet]: Initializes a new ordered set with the contents of a set.
//   - [INSOrderedSet.InitWithSetCopyItems]: Initializes a new ordered set with the contents of a set, optionally copying the objects in the set.
//
// # Counting Entries
//
//   - [INSOrderedSet.Count]: The number of members in the set.
//
// # Accessing Set Members
//
//   - [INSOrderedSet.ContainsObject]: Returns a Boolean value that indicates whether a given object is present in the ordered set.
//   - [INSOrderedSet.EnumerateObjectsAtIndexesOptionsUsingBlock]: Executes a given block using the objects in the ordered set at the specified indexes.
//   - [INSOrderedSet.EnumerateObjectsUsingBlock]: Executes a given block using each object in the ordered set.
//   - [INSOrderedSet.EnumerateObjectsWithOptionsUsingBlock]: Executes a given block using each object in the set, using the specified enumeration options.
//   - [INSOrderedSet.FirstObject]: The first object in the ordered set.
//   - [INSOrderedSet.LastObject]: The last object in the ordered set.
//   - [INSOrderedSet.ObjectAtIndex]: Returns the object at the specified index of the set.
//   - [INSOrderedSet.ObjectAtIndexedSubscript]: Returns the object at the specified index of the set.
//   - [INSOrderedSet.ObjectsAtIndexes]: Returns the objects in the ordered set at the specified indexes.
//   - [INSOrderedSet.IndexOfObject]: Returns the index of the specified object.
//   - [INSOrderedSet.IndexOfObjectInSortedRangeOptionsUsingComparator]: Returns the index, within a specified range, of an object compared with elements in the ordered set using a given NSComparator block.
//   - [INSOrderedSet.IndexOfObjectAtIndexesOptionsPassingTest]: Returns the index, from a given set of indexes, of the object in the ordered set that passes a test in a given block for a given set of enumeration options.
//   - [INSOrderedSet.IndexOfObjectPassingTest]: Returns the index of the object in the ordered set that passes a test in a given block.
//   - [INSOrderedSet.IndexOfObjectWithOptionsPassingTest]: Returns the index of an object in the ordered set that passes a test in a given block for a given set of enumeration options.
//   - [INSOrderedSet.IndexesOfObjectsAtIndexesOptionsPassingTest]: Returns the index, from a given set of indexes, of the object in the ordered set that passes a test in a given block for a given set of enumeration options.
//   - [INSOrderedSet.IndexesOfObjectsPassingTest]: Returns the index of the object in the ordered set that passes a test in a given block.
//   - [INSOrderedSet.IndexesOfObjectsWithOptionsPassingTest]: Returns the index of an object in the ordered set that passes a test in a given block for a given set of enumeration options.
//   - [INSOrderedSet.ObjectEnumerator]: Returns an enumerator object that lets you access each object in the ordered set.
//   - [INSOrderedSet.ReverseObjectEnumerator]: Returns an enumerator object that lets you access each object in the ordered set.
//   - [INSOrderedSet.ReversedOrderedSet]: An ordered set in the reverse order.
//
// # Comparing Sets
//
//   - [INSOrderedSet.IsEqualToOrderedSet]: Compares the receiving ordered set to another ordered set.
//   - [INSOrderedSet.IntersectsOrderedSet]: Returns a Boolean value that indicates whether at least one object in the receiving ordered set is also present in another given ordered set.
//   - [INSOrderedSet.IntersectsSet]: Returns a Boolean value that indicates whether at least one object in the receiving ordered set is also present in another given set.
//   - [INSOrderedSet.IsSubsetOfOrderedSet]: Returns a Boolean value that indicates whether every object in the receiving ordered set is also present in another given ordered set.
//   - [INSOrderedSet.IsSubsetOfSet]: Returns a Boolean value that indicates whether every object in the receiving ordered set is also present in another given set.
//
// # Creating a Sorted Array
//
//   - [INSOrderedSet.SortedArrayUsingDescriptors]: Returns an array of the ordered set’s elements sorted as specified by a given array of sort descriptors.
//   - [INSOrderedSet.SortedArrayUsingComparator]: Returns an array that lists the receiving ordered set’s elements in ascending order, as determined by the comparison method specified by a given [NSComparator] block
//   - [INSOrderedSet.SortedArrayWithOptionsUsingComparator]: Returns an array that lists the receiving ordered set’s elements in ascending order, as determined by the comparison method specified by a given [NSComparator] block.
//
// # Filtering Ordered Sets
//
//   - [INSOrderedSet.FilteredOrderedSetUsingPredicate]: Evaluates a given predicate against each object in the receiving ordered set and returns a new ordered set containing the objects for which the predicate returns true.
//
// # Describing a Set
//
//   - [INSOrderedSet.Description]: A string that represents the contents of the ordered set, formatted as a property list.
//   - [INSOrderedSet.DescriptionWithLocale]: Returns a string that represents the contents of the ordered set, formatted as a property list.
//   - [INSOrderedSet.DescriptionWithLocaleIndent]: Returns a string that represents the contents of the ordered set, formatted as a property list.
//
// # Converting Other Collections
//
//   - [INSOrderedSet.Array]: A representation of the ordered set as an array.
//   - [INSOrderedSet.Set]: A representation of the set containing the contents of the ordered set.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet
type INSOrderedSet interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSMutableCopying
	NSSecureCoding

	// Topic: Initializing an Ordered Set

	// Initializes a newly allocated set with the objects that are contained in a given array.
	InitWithArray(array []objectivec.IObject) NSOrderedSet
	// Initializes a newly allocated set with the objects that are contained in a given array, optionally copying the items.
	InitWithArrayCopyItems(set []objectivec.IObject, flag bool) NSOrderedSet
	// Initializes a newly allocated set with the objects that are contained in the specified range of an array, optionally copying the items.
	InitWithArrayRangeCopyItems(set []objectivec.IObject, range_ NSRange, flag bool) NSOrderedSet
	// Initializes a new ordered set with the object.
	InitWithObject(object objectivec.IObject) NSOrderedSet
	// Initializes a newly allocated set with a specified number of objects from a given C array of objects.
	InitWithObjectsCount(objects []objectivec.IObject, cnt uint) NSOrderedSet
	// Initializes a new ordered set with the contents of a set.
	InitWithOrderedSet(set INSOrderedSet) NSOrderedSet
	// Initializes a new ordered set with the contents of a set, optionally copying the items.
	InitWithOrderedSetCopyItems(set INSOrderedSet, flag bool) NSOrderedSet
	// Initializes a new ordered set with the contents of an ordered set, optionally copying the items.
	InitWithOrderedSetRangeCopyItems(set INSOrderedSet, range_ NSRange, flag bool) NSOrderedSet
	// Initializes a new ordered set with the contents of a set.
	InitWithSet(set INSSet) NSOrderedSet
	// Initializes a new ordered set with the contents of a set, optionally copying the objects in the set.
	InitWithSetCopyItems(set INSSet, flag bool) NSOrderedSet

	// Topic: Counting Entries

	// The number of members in the set.
	Count() uint

	// Topic: Accessing Set Members

	// Returns a Boolean value that indicates whether a given object is present in the ordered set.
	ContainsObject(object objectivec.IObject) bool
	// Executes a given block using the objects in the ordered set at the specified indexes.
	EnumerateObjectsAtIndexesOptionsUsingBlock(s INSIndexSet, opts NSEnumerationOptions, block ObjectTypeHandler)
	// Executes a given block using each object in the ordered set.
	EnumerateObjectsUsingBlock(block ObjectTypeHandler)
	// Executes a given block using each object in the set, using the specified enumeration options.
	EnumerateObjectsWithOptionsUsingBlock(opts NSEnumerationOptions, block ObjectTypeHandler)
	// The first object in the ordered set.
	FirstObject() objectivec.IObject
	// The last object in the ordered set.
	LastObject() objectivec.IObject
	// Returns the object at the specified index of the set.
	ObjectAtIndex(idx uint) objectivec.IObject
	// Returns the object at the specified index of the set.
	ObjectAtIndexedSubscript(idx uint) objectivec.IObject
	// Returns the objects in the ordered set at the specified indexes.
	ObjectsAtIndexes(indexes INSIndexSet) []objectivec.IObject
	// Returns the index of the specified object.
	IndexOfObject(object objectivec.IObject) uint
	// Returns the index, within a specified range, of an object compared with elements in the ordered set using a given NSComparator block.
	IndexOfObjectInSortedRangeOptionsUsingComparator(object objectivec.IObject, range_ NSRange, opts NSBinarySearchingOptions, cmp NSComparator) uint
	// Returns the index, from a given set of indexes, of the object in the ordered set that passes a test in a given block for a given set of enumeration options.
	IndexOfObjectAtIndexesOptionsPassingTest(s INSIndexSet, opts NSEnumerationOptions, predicate ObjectTypeHandler) uint
	// Returns the index of the object in the ordered set that passes a test in a given block.
	IndexOfObjectPassingTest(predicate ObjectTypeHandler) uint
	// Returns the index of an object in the ordered set that passes a test in a given block for a given set of enumeration options.
	IndexOfObjectWithOptionsPassingTest(opts NSEnumerationOptions, predicate ObjectTypeHandler) uint
	// Returns the index, from a given set of indexes, of the object in the ordered set that passes a test in a given block for a given set of enumeration options.
	IndexesOfObjectsAtIndexesOptionsPassingTest(s INSIndexSet, opts NSEnumerationOptions, predicate ObjectTypeHandler) INSIndexSet
	// Returns the index of the object in the ordered set that passes a test in a given block.
	IndexesOfObjectsPassingTest(predicate ObjectTypeHandler) INSIndexSet
	// Returns the index of an object in the ordered set that passes a test in a given block for a given set of enumeration options.
	IndexesOfObjectsWithOptionsPassingTest(opts NSEnumerationOptions, predicate ObjectTypeHandler) INSIndexSet
	// Returns an enumerator object that lets you access each object in the ordered set.
	ObjectEnumerator() INSEnumerator
	// Returns an enumerator object that lets you access each object in the ordered set.
	ReverseObjectEnumerator() INSEnumerator
	// An ordered set in the reverse order.
	ReversedOrderedSet() INSOrderedSet

	// Topic: Comparing Sets

	// Compares the receiving ordered set to another ordered set.
	IsEqualToOrderedSet(other INSOrderedSet) bool
	// Returns a Boolean value that indicates whether at least one object in the receiving ordered set is also present in another given ordered set.
	IntersectsOrderedSet(other INSOrderedSet) bool
	// Returns a Boolean value that indicates whether at least one object in the receiving ordered set is also present in another given set.
	IntersectsSet(set INSSet) bool
	// Returns a Boolean value that indicates whether every object in the receiving ordered set is also present in another given ordered set.
	IsSubsetOfOrderedSet(other INSOrderedSet) bool
	// Returns a Boolean value that indicates whether every object in the receiving ordered set is also present in another given set.
	IsSubsetOfSet(set INSSet) bool

	// Topic: Creating a Sorted Array

	// Returns an array of the ordered set’s elements sorted as specified by a given array of sort descriptors.
	SortedArrayUsingDescriptors(sortDescriptors []NSSortDescriptor) []objectivec.IObject
	// Returns an array that lists the receiving ordered set’s elements in ascending order, as determined by the comparison method specified by a given [NSComparator] block
	SortedArrayUsingComparator(cmptr NSComparator) []objectivec.IObject
	// Returns an array that lists the receiving ordered set’s elements in ascending order, as determined by the comparison method specified by a given [NSComparator] block.
	SortedArrayWithOptionsUsingComparator(opts NSSortOptions, cmptr NSComparator) []objectivec.IObject

	// Topic: Filtering Ordered Sets

	// Evaluates a given predicate against each object in the receiving ordered set and returns a new ordered set containing the objects for which the predicate returns true.
	FilteredOrderedSetUsingPredicate(p INSPredicate) INSOrderedSet

	// Topic: Describing a Set

	// A string that represents the contents of the ordered set, formatted as a property list.
	Description() string
	// Returns a string that represents the contents of the ordered set, formatted as a property list.
	DescriptionWithLocale(locale objectivec.IObject) string
	// Returns a string that represents the contents of the ordered set, formatted as a property list.
	DescriptionWithLocaleIndent(locale objectivec.IObject, level uint) string

	// Topic: Converting Other Collections

	// A representation of the ordered set as an array.
	Array() []objectivec.IObject
	// A representation of the set containing the contents of the ordered set.
	Set() INSSet

	// Compares two ordered sets to create a difference object that represents the changes between them.
	DifferenceFromOrderedSet(other INSOrderedSet) INSOrderedCollectionDifference
	// Compares two ordered sets, with options, to create a difference object that represents the changes between them.
	DifferenceFromOrderedSetWithOptions(other INSOrderedSet, options NSOrderedCollectionDifferenceCalculationOptions) INSOrderedCollectionDifference
	// Compares two ordered sets, using the provided block and with options, to create a difference object that represents the changes between them.
	DifferenceFromOrderedSetWithOptionsUsingEquivalenceTest(other INSOrderedSet, options NSOrderedCollectionDifferenceCalculationOptions, block ObjectTypeHandler) INSOrderedCollectionDifference
	// Copies the objects contained in the ordered set that fall within the specified range to `objects`.
	GetObjectsRange(objects []objectivec.IObject, range_ NSRange)
	// Initializes a newly allocated set with members taken from the specified list of objects.
	InitWithObjects(firstObj objectivec.IObject) NSOrderedSet
	// Creates a new ordered set by applying a difference object to an existing ordered set.
	OrderedSetByApplyingDifference(difference INSOrderedCollectionDifference) INSOrderedSet
}

// Init initializes the instance.
func (o NSOrderedSet) Init() NSOrderedSet {
	rv := objc.Send[NSOrderedSet](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o NSOrderedSet) Autorelease() NSOrderedSet {
	rv := objc.Send[NSOrderedSet](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSOrderedSet creates a new NSOrderedSet instance.
func NewNSOrderedSet() NSOrderedSet {
	class := getNSOrderedSetClass()
	rv := objc.Send[NSOrderedSet](objc.ID(class.class), objc.Sel("new"))
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
func NewOrderedSetWithArray(array []objectivec.IObject) NSOrderedSet {
	instance := getNSOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArray:"), objectivec.IObjectSliceToNSArray(array))
	return NSOrderedSetFromID(rv)
}

// Initializes a newly allocated set with the objects that are contained in a
// given array, optionally copying the items.
//
// set: An array of objects to add to the new set.
//
// If the same object appears more than once in array, it is represented only
// once in the returned ordered set.
//
// flag: If true the objects are copied to the ordered set; otherwise false.
//
// # Return Value
//
// An initialized ordered set containing a uniqued collection of the objects
// contained in the array.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(array:copyItems:)
func NewOrderedSetWithArrayCopyItems(set []objectivec.IObject, flag bool) NSOrderedSet {
	instance := getNSOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArray:copyItems:"), objectivec.IObjectSliceToNSArray(set), flag)
	return NSOrderedSetFromID(rv)
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
// flag: If true the objects are copied to the ordered set; otherwise false.
//
// # Return Value
//
// An initialized ordered set containing a uniqued collection of the objects
// contained in specified range of the array.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(array:range:copyItems:)
func NewOrderedSetWithArrayRangeCopyItems(set []objectivec.IObject, range_ NSRange, flag bool) NSOrderedSet {
	instance := getNSOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArray:range:copyItems:"), objectivec.IObjectSliceToNSArray(set), range_, flag)
	return NSOrderedSetFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(coder:)
func NewOrderedSetWithCoder(coder INSCoder) NSOrderedSet {
	instance := getNSOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSOrderedSetFromID(rv)
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
func NewOrderedSetWithObject(object objectivec.IObject) NSOrderedSet {
	instance := getNSOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObject:"), object)
	return NSOrderedSetFromID(rv)
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
func NewOrderedSetWithObjects(firstObj objectivec.IObject) NSOrderedSet {
	instance := getNSOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObjects:"), firstObj)
	return NSOrderedSetFromID(rv)
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
func NewOrderedSetWithOrderedSet(set INSOrderedSet) NSOrderedSet {
	instance := getNSOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOrderedSet:"), set)
	return NSOrderedSetFromID(rv)
}

// Initializes a new ordered set with the contents of a set, optionally
// copying the items.
//
// set: A set.
//
// flag: If true the objects are copied to the ordered set; otherwise false.
//
// # Return Value
//
// An initialized ordered set containing the objects in the set.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(orderedSet:copyItems:)
func NewOrderedSetWithOrderedSetCopyItems(set INSOrderedSet, flag bool) NSOrderedSet {
	instance := getNSOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOrderedSet:copyItems:"), set, flag)
	return NSOrderedSetFromID(rv)
}

// Initializes a new ordered set with the contents of an ordered set,
// optionally copying the items.
//
// set: An ordered set.
//
// range: The range of objects in `orderedSet` to add to the ordered set.
//
// flag: If true the objects are copied to the ordered set; otherwise false.
//
// # Return Value
//
// An initialized ordered set containing the objects in the ordered set.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(orderedSet:range:copyItems:)
func NewOrderedSetWithOrderedSetRangeCopyItems(set INSOrderedSet, range_ NSRange, flag bool) NSOrderedSet {
	instance := getNSOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOrderedSet:range:copyItems:"), set, range_, flag)
	return NSOrderedSetFromID(rv)
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
func NewOrderedSetWithSet(set INSSet) NSOrderedSet {
	instance := getNSOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSet:"), set)
	return NSOrderedSetFromID(rv)
}

// Initializes a new ordered set with the contents of a set, optionally
// copying the objects in the set.
//
// set: The set.
//
// flag: If true the objects are copied to the ordered set; otherwise false.
//
// # Return Value
//
// An initialized ordered set containing the objects in the set.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(set:copyItems:)
func NewOrderedSetWithSetCopyItems(set INSSet, flag bool) NSOrderedSet {
	instance := getNSOrderedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSet:copyItems:"), set, flag)
	return NSOrderedSetFromID(rv)
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
func (o NSOrderedSet) InitWithArray(array []objectivec.IObject) NSOrderedSet {
	rv := objc.Send[NSOrderedSet](o.ID, objc.Sel("initWithArray:"), objectivec.IObjectSliceToNSArray(array))
	return rv
}

// Initializes a newly allocated set with the objects that are contained in a
// given array, optionally copying the items.
//
// set: An array of objects to add to the new set.
//
// If the same object appears more than once in array, it is represented only
// once in the returned ordered set.
//
// flag: If true the objects are copied to the ordered set; otherwise false.
//
// # Return Value
//
// An initialized ordered set containing a uniqued collection of the objects
// contained in the array.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(array:copyItems:)
func (o NSOrderedSet) InitWithArrayCopyItems(set []objectivec.IObject, flag bool) NSOrderedSet {
	rv := objc.Send[NSOrderedSet](o.ID, objc.Sel("initWithArray:copyItems:"), objectivec.IObjectSliceToNSArray(set), flag)
	return rv
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
// flag: If true the objects are copied to the ordered set; otherwise false.
//
// # Return Value
//
// An initialized ordered set containing a uniqued collection of the objects
// contained in specified range of the array.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(array:range:copyItems:)
func (o NSOrderedSet) InitWithArrayRangeCopyItems(set []objectivec.IObject, range_ NSRange, flag bool) NSOrderedSet {
	rv := objc.Send[NSOrderedSet](o.ID, objc.Sel("initWithArray:range:copyItems:"), objectivec.IObjectSliceToNSArray(set), range_, flag)
	return rv
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
func (o NSOrderedSet) InitWithObject(object objectivec.IObject) NSOrderedSet {
	rv := objc.Send[NSOrderedSet](o.ID, objc.Sel("initWithObject:"), object)
	return rv
}

// Initializes a newly allocated set with a specified number of objects from a
// given C array of objects.
//
// objects: A C array of objects to add to the new set.
//
// If the same object appears more than once in objects, it is added only once
// to the returned ordered set.
//
// cnt: The number of objects from objects to add to the new ordered set.
//
// # Return Value
//
// An initialized ordered set containing cnt objects from the list of objects
// specified by objects. The returned set might be different than the original
// receiver.
//
// # Discussion
//
// This method is a designated initializer of [NSOrderedSet].
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(objects:count:)-2ai32
func (o NSOrderedSet) InitWithObjectsCount(objects []objectivec.IObject, cnt uint) NSOrderedSet {
	rv := objc.Send[NSOrderedSet](o.ID, objc.Sel("initWithObjects:count:"), objc.CArray(objects), cnt)
	return rv
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
func (o NSOrderedSet) InitWithOrderedSet(set INSOrderedSet) NSOrderedSet {
	rv := objc.Send[NSOrderedSet](o.ID, objc.Sel("initWithOrderedSet:"), set)
	return rv
}

// Initializes a new ordered set with the contents of a set, optionally
// copying the items.
//
// set: A set.
//
// flag: If true the objects are copied to the ordered set; otherwise false.
//
// # Return Value
//
// An initialized ordered set containing the objects in the set.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(orderedSet:copyItems:)
func (o NSOrderedSet) InitWithOrderedSetCopyItems(set INSOrderedSet, flag bool) NSOrderedSet {
	rv := objc.Send[NSOrderedSet](o.ID, objc.Sel("initWithOrderedSet:copyItems:"), set, flag)
	return rv
}

// Initializes a new ordered set with the contents of an ordered set,
// optionally copying the items.
//
// set: An ordered set.
//
// range: The range of objects in `orderedSet` to add to the ordered set.
//
// flag: If true the objects are copied to the ordered set; otherwise false.
//
// # Return Value
//
// An initialized ordered set containing the objects in the ordered set.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(orderedSet:range:copyItems:)
func (o NSOrderedSet) InitWithOrderedSetRangeCopyItems(set INSOrderedSet, range_ NSRange, flag bool) NSOrderedSet {
	rv := objc.Send[NSOrderedSet](o.ID, objc.Sel("initWithOrderedSet:range:copyItems:"), set, range_, flag)
	return rv
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
func (o NSOrderedSet) InitWithSet(set INSSet) NSOrderedSet {
	rv := objc.Send[NSOrderedSet](o.ID, objc.Sel("initWithSet:"), set)
	return rv
}

// Initializes a new ordered set with the contents of a set, optionally
// copying the objects in the set.
//
// set: The set.
//
// flag: If true the objects are copied to the ordered set; otherwise false.
//
// # Return Value
//
// An initialized ordered set containing the objects in the set.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(set:copyItems:)
func (o NSOrderedSet) InitWithSetCopyItems(set INSSet, flag bool) NSOrderedSet {
	rv := objc.Send[NSOrderedSet](o.ID, objc.Sel("initWithSet:copyItems:"), set, flag)
	return rv
}

// Returns a Boolean value that indicates whether a given object is present in
// the ordered set.
//
// object: The object for which to test membership of the ordered set.
//
// # Return Value
//
// YES if `object` is present in the set, otherwise false.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/contains(_:)
func (o NSOrderedSet) ContainsObject(object objectivec.IObject) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("containsObject:"), object)
	return rv
}

// Executes a given block using the objects in the ordered set at the
// specified indexes.
//
// s: The indexes of the objects over which to enumerate.
//
// opts: A bitmask that specifies the options for the enumeration (whether it should
// be performed concurrently and whether it should be performed in reverse
// order).
//
// block: The block to apply to elements in the ordered set.
//
// The block takes three arguments:
//
// obj: The element in the ordered set. idx: The index of the element in the
// ordered set. stop: A reference to a Boolean value. The block can set the
// value to true to stop further processing of the array. The `stop` argument
// is an out-only argument. You should only ever set this Boolean to true
// within the block.
//
// # Discussion
//
// By default, the enumeration starts with the first object and continues
// serially through the ordered set to the last element specified by `s`. You
// can specify [NSEnumerationConcurrent] and/or [NSEnumerationReverse] as
// enumeration options to modify this behavior.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/enumerateObjects(at:options:using:)
func (o NSOrderedSet) EnumerateObjectsAtIndexesOptionsUsingBlock(s INSIndexSet, opts NSEnumerationOptions, block ObjectTypeHandler) {
	_block2, _ := NewObjectTypeBlock(block)
	objc.Send[objc.ID](o.ID, objc.Sel("enumerateObjectsAtIndexes:options:usingBlock:"), s, opts, _block2)
}

// Executes a given block using each object in the ordered set.
//
// block: The block to apply to elements in the ordered set.
//
// The block takes three arguments:
//
// idx: The element in the set. idx: The index of the item in the set. stop: A
// reference to a Boolean value. The block can set the value to true to stop
// further processing of the set. The `stop` argument is an out-only argument.
// You should only ever set this value to true within the block.
//
// The block returns a Boolean value that indicates whether `obj` passed the
// test.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/enumerateObjects(_:)
func (o NSOrderedSet) EnumerateObjectsUsingBlock(block ObjectTypeHandler) {
	_block0, _ := NewObjectTypeBlock(block)
	objc.Send[objc.ID](o.ID, objc.Sel("enumerateObjectsUsingBlock:"), _block0)
}

// Executes a given block using each object in the set, using the specified
// enumeration options.
//
// opts: A bitmask that specifies the options for the enumeration (whether it should
// be performed concurrently and whether it should be performed in reverse
// order).
//
// block: The block to apply to elements in the ordered set.
//
// The block takes three arguments:
//
// obj: The element in the set. idx: The index of the item in the set. stop: A
// reference to a Boolean value. The block can set the value to true to stop
// further processing of the set. The `stop` argument is an out-only argument.
// You should only ever set this value to true within the block.
//
// The block returns a Boolean value that indicates whether `obj` passed the
// test.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/enumerateObjects(options:using:)
func (o NSOrderedSet) EnumerateObjectsWithOptionsUsingBlock(opts NSEnumerationOptions, block ObjectTypeHandler) {
	_block1, _ := NewObjectTypeBlock(block)
	objc.Send[objc.ID](o.ID, objc.Sel("enumerateObjectsWithOptions:usingBlock:"), opts, _block1)
}

// Returns the object at the specified index of the set.
//
// idx: The object located at index.
//
// # Return Value
//
// If `idx` is beyond the end of the ordered set (that is, if index is greater
// than or equal to the value returned by count), an [RangeException] is
// raised.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/object(at:)
func (o NSOrderedSet) ObjectAtIndex(idx uint) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("objectAtIndex:"), idx)
	return objectivec.Object{ID: rv}
}

// Returns the object at the specified index of the set.
//
// idx: The object located at index.
//
// # Return Value
//
// If `idx` is beyond the end of the ordered set (that is, if index is greater
// than or equal to the value returned by count), an [RangeException] is
// raised.
//
// # Discussion
//
// This method is the same as [ObjectAtIndex].
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/subscript(_:)
func (o NSOrderedSet) ObjectAtIndexedSubscript(idx uint) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("objectAtIndexedSubscript:"), idx)
	return objectivec.Object{ID: rv}
}

// Returns the objects in the ordered set at the specified indexes.
//
// indexes: The indexes.
//
// # Return Value
//
// The returned objects are in the ascending order of their indexes in
// indexes, so that object in returned ordered set with higher index in
// indexes will follow the object with smaller index in indexes.
//
// # Discussion
//
// Raises an [RangeException] if any location in indexes exceeds the bounds of
// the array, of if `indexes` is `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/objects(at:)
func (o NSOrderedSet) ObjectsAtIndexes(indexes INSIndexSet) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("objectsAtIndexes:"), indexes)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns the index of the specified object.
//
// object: The object.
//
// # Return Value
//
// The index whose corresponding ordered set value is equal to `object`. If
// none of the objects in the ordered set is equal to `object`, returns
// [NSNotFound].
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/index(of:)
func (o NSOrderedSet) IndexOfObject(object objectivec.IObject) uint {
	rv := objc.Send[uint](o.ID, objc.Sel("indexOfObject:"), object)
	return rv
}

// Returns the index, within a specified range, of an object compared with
// elements in the ordered set using a given NSComparator block.
//
// object: An object for which to search in the ordered set.
//
// If this value is `nil`, throws an [InvalidArgumentException].
//
// range: The range within the array to search for `object`.
//
// If r exceeds the bounds of the ordered set (if the location plus length of
// the range is greater than the count of the ordered set), throws an
// [RangeException].
//
// opts: Options for the search. For possible values, see
// [NSBinarySearchingOptions].
//
// cmp: A comparator block used to compare the object obj with elements in the
// ordered set.
//
// If this value is [NULL], throws an [InvalidArgumentException].
//
// # Return Value
//
// If the [NSBinarySearchingInsertionIndex] option is not specified:
//
// - If the `object` is found and neither [NSBinarySearchingFirstEqual] nor
// [NSBinarySearchingLastEqual] is specified, returns the matching object’s
// index. - If the [NSBinarySearchingFirstEqual] or
// [NSBinarySearchingLastEqual] option is also specified, returns the index of
// equal objects. - If the object is not found, returns [NSNotFound].
//
// If the [NSBinarySearchingInsertionIndex] option is specified, returns the
// index at which you should insert `obj` in order to maintain a sorted array:
//
// - If the `object` is found and neither [NSBinarySearchingFirstEqual] nor
// [NSBinarySearchingLastEqual] is specified, returns the matching object’s
// index. - If the [NSBinarySearchingFirstEqual] or
// [NSBinarySearchingLastEqual] option is also specified, returns the index of
// the equal objects. - If the object is not found, returns the index of the
// least greater object, or the index at the end of the array if the object is
// larger than all other elements.
//
// # Discussion
//
// The elements in the ordered set must have already been sorted using the
// comparator `cmp`. If the ordered set is not sorted, the result is
// undefined.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/index(of:inSortedRange:options:usingComparator:)
//
// [NSBinarySearchingOptions]: https://developer.apple.com/documentation/Foundation/NSBinarySearchingOptions
func (o NSOrderedSet) IndexOfObjectInSortedRangeOptionsUsingComparator(object objectivec.IObject, range_ NSRange, opts NSBinarySearchingOptions, cmp NSComparator) uint {
	rv := objc.Send[uint](o.ID, objc.Sel("indexOfObject:inSortedRange:options:usingComparator:"), object, range_, opts, cmp)
	return rv
}

// Returns the index, from a given set of indexes, of the object in the
// ordered set that passes a test in a given block for a given set of
// enumeration options.
//
// s: The indexes of the objects over which to enumerate.
//
// opts: A bitmask that specifies the options for the enumeration (whether it should
// be performed concurrently and whether it should be performed in reverse
// order).
//
// predicate: The block to apply to elements in the ordered set.
//
// The block takes three arguments:
//
// obj: The element in the ordered set. idx: The index of the element in the
// ordered set. stop: A reference to a Boolean value. The block can set the
// value to true to stop further processing of the set. The `stop` argument is
// an out-only argument. You should only ever set this value to true within
// the block.
//
// The block returns a Boolean value that indicates whether `obj` passed the
// test.
//
// # Return Value
//
// The index of the corresponding value in the ordered set passes the test
// specified by predicate. If no objects in the ordered set pass the test,
// returns [NSNotFound].
//
// # Discussion
//
// By default, the enumeration starts with the first object and continues
// serially through the ordered set to the last element specified by `s`. You
// can specify [NSEnumerationConcurrent] and/or [NSEnumerationReverse] as
// enumeration options to modify this behavior.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/index(ofObjectAt:options:passingTest:)
func (o NSOrderedSet) IndexOfObjectAtIndexesOptionsPassingTest(s INSIndexSet, opts NSEnumerationOptions, predicate ObjectTypeHandler) uint {
	_block2, _ := NewObjectTypeBlock(predicate)
	rv := objc.Send[uint](o.ID, objc.Sel("indexOfObjectAtIndexes:options:passingTest:"), s, opts, _block2)
	return rv
}

// Returns the index of the object in the ordered set that passes a test in a
// given block.
//
// predicate: The block to apply to elements in the ordered set.
//
// The block takes three arguments:
//
// obj: The element in the ordered set. Term: The index of the element in the
// ordered set. stop: A reference to a Boolean value. The block can set the
// value to true to stop further processing of the set. The `stop` argument is
// an out-only argument. You should only ever set this value to true within
// the block.
//
// # Return Value
//
// The index of the corresponding value in the ordered set that passes the
// test specified by predicate. If no objects in the ordered set pass the
// test, returns [NSNotFound].
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/index(ofObjectPassingTest:)
func (o NSOrderedSet) IndexOfObjectPassingTest(predicate ObjectTypeHandler) uint {
	_block0, _ := NewObjectTypeBlock(predicate)
	rv := objc.Send[uint](o.ID, objc.Sel("indexOfObjectPassingTest:"), _block0)
	return rv
}

// Returns the index of an object in the ordered set that passes a test in a
// given block for a given set of enumeration options.
//
// opts: A bitmask that specifies the options for the enumeration (whether it should
// be performed concurrently and whether it should be performed in reverse
// order).
//
// predicate: The block to apply to elements in the ordered set.
//
// The block takes three arguments:
//
// obj: The element in the array. idx: The index of the element in the ordered
// set. stop: A reference to a Boolean value. The block can set the value to
// true to stop further processing of the set. The `stop` argument is an
// out-only argument. You should only ever set this value to true within the
// block.
//
// The block returns a Boolean value that indicates whether obj passed the
// test.
//
// # Return Value
//
// The index whose corresponding value in the ordered set passes the test
// specified by `predicate` and `opts`. If no objects in the ordered set pass
// the test, returns [NSNotFound].
//
// # Discussion
//
// By default, the enumeration starts with the first object and continues
// serially through the ordered set to the last object. You can specify
// [NSEnumerationConcurrent] and/or [NSEnumerationReverse] as enumeration
// options to modify this behavior.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/index(_:ofObjectPassingTest:)
func (o NSOrderedSet) IndexOfObjectWithOptionsPassingTest(opts NSEnumerationOptions, predicate ObjectTypeHandler) uint {
	_block1, _ := NewObjectTypeBlock(predicate)
	rv := objc.Send[uint](o.ID, objc.Sel("indexOfObjectWithOptions:passingTest:"), opts, _block1)
	return rv
}

// Returns the index, from a given set of indexes, of the object in the
// ordered set that passes a test in a given block for a given set of
// enumeration options.
//
// s: The indexes of the objects over which to enumerate.
//
// opts: A bitmask that specifies the options for the enumeration (whether it should
// be performed concurrently and whether it should be performed in reverse
// order).
//
// predicate: The block to apply to elements in the ordered set.
//
// The block takes three arguments:
//
// obj: The element in the ordered set. idx: The index of the element in the
// ordered set. stop: A reference to a Boolean value. The block can set the
// value to true to stop further processing of the set. The `stop` argument is
// an out-only argument. You should only ever set this value to true within
// the block.
//
// The block returns a Boolean value that indicates whether `obj` passed the
// test.
//
// # Return Value
//
// The index of the corresponding value in the ordered set that passes the
// test specified by predicate. If no objects in the ordered set pass the
// test, returns NSNotFound.
//
// # Discussion
//
// By default, the enumeration starts with the first object and continues
// serially through the ordered set to the last object. You can specify
// [NSEnumerationConcurrent] and/or [NSEnumerationReverse] as enumeration
// options to modify this behavior.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/indexes(ofObjectsAt:options:passingTest:)
func (o NSOrderedSet) IndexesOfObjectsAtIndexesOptionsPassingTest(s INSIndexSet, opts NSEnumerationOptions, predicate ObjectTypeHandler) INSIndexSet {
	_block2, _ := NewObjectTypeBlock(predicate)
	rv := objc.Send[objc.ID](o.ID, objc.Sel("indexesOfObjectsAtIndexes:options:passingTest:"), s, opts, _block2)
	return NSIndexSetFromID(rv)
}

// Returns the index of the object in the ordered set that passes a test in a
// given block.
//
// predicate: The block to apply to elements in the ordered set.
//
// The block takes three arguments:
//
// obj: The element in the ordered set. Term: The index of the element in the
// ordered set. stop: A reference to a Boolean value. The block can set the
// value to true to stop further processing of the set. The `stop` argument is
// an out-only argument. You should only ever set this value to true within
// the block.
//
// # Return Value
//
// The index of the corresponding value in the ordered set that passes the
// test specified by predicate. If no objects in the ordered set pass the
// test, returns NSNotFound..
//
// # Discussion
//
// If the block parameter is `nil`, this method raises an exception.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/indexes(ofObjectsPassingTest:)
func (o NSOrderedSet) IndexesOfObjectsPassingTest(predicate ObjectTypeHandler) INSIndexSet {
	_block0, _ := NewObjectTypeBlock(predicate)
	rv := objc.Send[objc.ID](o.ID, objc.Sel("indexesOfObjectsPassingTest:"), _block0)
	return NSIndexSetFromID(rv)
}

// Returns the index of an object in the ordered set that passes a test in a
// given block for a given set of enumeration options.
//
// opts: A bitmask that specifies the options for the enumeration (whether it should
// be performed concurrently and whether it should be performed in reverse
// order).
//
// predicate: The block to apply to elements in the ordered set.
//
// The block takes three arguments:
//
// obj: The element in the ordered set. Term: The index of the element in the
// ordered set. stop: A reference to a Boolean value. The block can set the
// value to true to stop further processing of the set. The `stop` argument is
// an out-only argument. You should only ever set this value to true within
// the block.
//
// # Return Value
//
// The index whose corresponding value in the ordered set passes the test
// specified by `predicate` and `opts`. If the `opts` bitmask specifies
// reverse order, then the last item that matches is returned. Otherwise, the
// index of the first matching object is returned. If no objects in the
// ordered set pass the test, returns [NSNotFound].
//
// # Discussion
//
// By default, the enumeration starts with the first object and continues
// serially through the ordered set to the last object. You can specify
// [NSEnumerationConcurrent] and/or [NSEnumerationReverse] as enumeration
// options to modify this behavior.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/indexes(options:ofObjectsPassingTest:)
func (o NSOrderedSet) IndexesOfObjectsWithOptionsPassingTest(opts NSEnumerationOptions, predicate ObjectTypeHandler) INSIndexSet {
	_block1, _ := NewObjectTypeBlock(predicate)
	rv := objc.Send[objc.ID](o.ID, objc.Sel("indexesOfObjectsWithOptions:passingTest:"), opts, _block1)
	return NSIndexSetFromID(rv)
}

// Returns an enumerator object that lets you access each object in the
// ordered set.
//
// # Return Value
//
// An enumerator object that lets you access each object in the ordered set,
// in order, from the element at the lowest index upwards.
//
// # Discussion
//
// When you use this method with mutable subclasses of [NSOrderedSet], you
// must not modify the ordered set during enumeration.
//
// It is more efficient to use the fast enumeration protocol (see
// [NSFastEnumeration]). Fast enumeration is available in macOS 10.5 and later
// and iOS 2.0 and later.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/objectEnumerator()
func (o NSOrderedSet) ObjectEnumerator() INSEnumerator {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("objectEnumerator"))
	return NSEnumeratorFromID(rv)
}

// Returns an enumerator object that lets you access each object in the
// ordered set.
//
// # Return Value
//
// An enumerator object that lets you access each object in the ordered set,
// in order, from the element at the highest index downwards.
//
// # Discussion
//
// When you use this method with mutable subclasses of [NSOrderedSet], you
// must not modify the ordered set during enumeration.
//
// It is more efficient to use the fast enumeration protocol (see
// [NSFastEnumeration]). Fast enumeration is available in macOS 10.5 and later
// and iOS 2.0 and later.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/reverseObjectEnumerator()
func (o NSOrderedSet) ReverseObjectEnumerator() INSEnumerator {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("reverseObjectEnumerator"))
	return NSEnumeratorFromID(rv)
}

// Compares the receiving ordered set to another ordered set.
//
// other: The ordered set with which to compare the receiving ordered set.
//
// # Return Value
//
// true if the contents of `other` are equal to the contents of the receiving
// ordered set, otherwise false.
//
// # Discussion
//
// Two ordered sets have equal contents if they each have the same number of
// members, if each member of one ordered set is present in the other, and the
// members are in the same order.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/isEqual(to:)
func (o NSOrderedSet) IsEqualToOrderedSet(other INSOrderedSet) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isEqualToOrderedSet:"), other)
	return rv
}

// Returns a Boolean value that indicates whether at least one object in the
// receiving ordered set is also present in another given ordered set.
//
// other: The other ordered set.
//
// # Return Value
//
// true if at least one object in the receiving ordered set is also present in
// `other`, otherwise false.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/intersects(_:)
func (o NSOrderedSet) IntersectsOrderedSet(other INSOrderedSet) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("intersectsOrderedSet:"), other)
	return rv
}

// Returns a Boolean value that indicates whether at least one object in the
// receiving ordered set is also present in another given set.
//
// set: The set.
//
// # Return Value
//
// true if at least one object in the receiving ordered set is also present in
// `other`, otherwise false.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/intersectsSet(_:)
func (o NSOrderedSet) IntersectsSet(set INSSet) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("intersectsSet:"), set)
	return rv
}

// Returns a Boolean value that indicates whether every object in the
// receiving ordered set is also present in another given ordered set.
//
// other: The ordered set with which to compare the receiving ordered set.
//
// # Return Value
//
// true if every object in the receiving set is also present in `other`,
// otherwise false.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/isSubset(of:)-7brc
func (o NSOrderedSet) IsSubsetOfOrderedSet(other INSOrderedSet) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isSubsetOfOrderedSet:"), other)
	return rv
}

// Returns a Boolean value that indicates whether every object in the
// receiving ordered set is also present in another given set.
//
// set: The set with which to compare the receiving ordered set.
//
// # Return Value
//
// true if every object in the receiving ordered set is also present in `set`,
// otherwise false.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/isSubset(of:)-8zx9x
func (o NSOrderedSet) IsSubsetOfSet(set INSSet) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isSubsetOfSet:"), set)
	return rv
}

// Returns an array of the ordered set’s elements sorted as specified by a
// given array of sort descriptors.
//
// sortDescriptors: An array of [NSSortDescriptor] objects.
//
// # Return Value
//
// An [NSArray] containing the ordered set’s elements sorted as specified by
// `sortDescriptors`.
//
// # Discussion
//
// The first descriptor specifies the primary key path to be used in sorting
// the ordered set’s elements. Any subsequent descriptors are used to
// further refine sorting of objects with duplicate values. See
// [NSSortDescriptor] for additional information.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/sortedArray(using:)
func (o NSOrderedSet) SortedArrayUsingDescriptors(sortDescriptors []NSSortDescriptor) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("sortedArrayUsingDescriptors:"), objectivec.IObjectSliceToNSArray(sortDescriptors))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns an array that lists the receiving ordered set’s elements in
// ascending order, as determined by the comparison method specified by a
// given [NSComparator] block
//
// cmptr: A comparator block.
//
// # Return Value
//
// An array that lists the receiving ordered set’s elements in ascending
// order, as determined by the comparison method specified `cmptr`.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/sortedArray(comparator:)
func (o NSOrderedSet) SortedArrayUsingComparator(cmptr NSComparator) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("sortedArrayUsingComparator:"), cmptr)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns an array that lists the receiving ordered set’s elements in
// ascending order, as determined by the comparison method specified by a
// given [NSComparator] block.
//
// opts: A bitmask that specifies the options for the sort (whether it should be
// performed concurrently and whether it should be performed stably).
//
// cmptr: A comparator block.
//
// # Return Value
//
// An array that lists the receiving ordered set’s elements in ascending
// order, as determined by the comparison method specified `cmptr`.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/sortedArray(options:usingComparator:)
func (o NSOrderedSet) SortedArrayWithOptionsUsingComparator(opts NSSortOptions, cmptr NSComparator) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("sortedArrayWithOptions:usingComparator:"), opts, cmptr)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Evaluates a given predicate against each object in the receiving ordered
// set and returns a new ordered set containing the objects for which the
// predicate returns true.
//
// p: The predicate against which to evaluate the receiving ordered set’s
// elements.
//
// # Return Value
//
// A new ordered set containing the objects in the receiving ordered set for
// which `p` returns true.
//
// # Discussion
//
// For more details, see [Predicate Programming Guide].
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/filtered(using:)
//
// [Predicate Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Predicates/AdditionalChapters/Introduction.html#//apple_ref/doc/uid/TP40001789
func (o NSOrderedSet) FilteredOrderedSetUsingPredicate(p INSPredicate) INSOrderedSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("filteredOrderedSetUsingPredicate:"), p)
	return NSOrderedSetFromID(rv)
}

// Returns a string that represents the contents of the ordered set, formatted
// as a property list.
//
// locale: An [NSLocale] object or an [NSDictionary] object that specifies options
// used for formatting each of the ordered set’s elements (where
// recognized). Specify `nil` if you don’t want the elements formatted.
//
// # Return Value
//
// A string that represents the contents of the ordered set, formatted as a
// property list.
//
// # Discussion
//
// For a description of how locale is applied to each element in the receiving
// ordered set, see [DescriptionWithLocaleIndent].
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/description(withLocale:)
func (o NSOrderedSet) DescriptionWithLocale(locale objectivec.IObject) string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("descriptionWithLocale:"), locale)
	return NSStringFromID(rv).String()
}

// Returns a string that represents the contents of the ordered set, formatted
// as a property list.
//
// locale: An [NSLocale] object or an [NSDictionary] object that specifies options
// used for formatting each of the array’s elements (where recognized).
// Specify `nil` if you don’t want the elements formatted.
//
// level: Specifies a level of indentation, to make the output more readable: the
// indentation is (4 spaces) * `level`.
//
// # Return Value
//
// A string that represents the contents of the ordered set, formatted as a
// property list.
//
// # Discussion
//
// The returned [NSString] object contains the string representations of each
// of the ordered set’s elements, in order, from first to last. To obtain
// the string representation of a given element, “ proceeds as follows:
//
// - If the element is an [NSString] object, it is used as is. - If the
// element responds to “, that method is invoked to obtain the element’s
// string representation. - If the element responds to “, that method is
// invoked to obtain the element’s string representation. - If none of the
// above conditions is met, the element’s string representation is obtained
// by invoking its `description` method
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/description(withLocale:indent:)
func (o NSOrderedSet) DescriptionWithLocaleIndent(locale objectivec.IObject, level uint) string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("descriptionWithLocale:indent:"), locale, level)
	return NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(coder:)
func (o NSOrderedSet) InitWithCoder(coder INSCoder) NSOrderedSet {
	rv := objc.Send[NSOrderedSet](o.ID, objc.Sel("initWithCoder:"), coder)
	return rv
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
func (o NSOrderedSet) CountByEnumeratingWithStateObjectsCount(state NSFastEnumerationState, buffer []objectivec.IObject, len_ uint) uint {
	rv := objc.Send[uint](o.ID, objc.Sel("countByEnumeratingWithState:objects:count:"), state, objc.CArray(buffer), len_)
	return rv
}

// Compares two ordered sets to create a difference object that represents the
// changes between them.
//
// # Discussion
//
// The difference method creates the difference object by comparing objects
// within the ordered sets with the “ method.
//
// The following example computes the difference between two ordered sets:
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/differenceFromOrderedSet:
func (o NSOrderedSet) DifferenceFromOrderedSet(other INSOrderedSet) INSOrderedCollectionDifference {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("differenceFromOrderedSet:"), other)
	return NSOrderedCollectionDifferenceFromID(rv)
}

// Compares two ordered sets, with options, to create a difference object that
// represents the changes between them.
//
// # Discussion
//
// The difference method creates the difference object by comparing objects
// within the ordered sets with the “ method.
//
// The options allow you to choose to omit insertion or removal references to
// the change objects within the difference object. You can also choose to
// infer moves when computing the difference, which provides an
// [AssociatedIndex] within the change objects that indicates the index in the
// ordered set where the object moved from.
//
// The following example computes the difference between two ordered sets,
// inferring moves between them:
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/differenceFromOrderedSet:withOptions:
func (o NSOrderedSet) DifferenceFromOrderedSetWithOptions(other INSOrderedSet, options NSOrderedCollectionDifferenceCalculationOptions) INSOrderedCollectionDifference {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("differenceFromOrderedSet:withOptions:"), other, options)
	return NSOrderedCollectionDifferenceFromID(rv)
}

// Compares two ordered sets, using the provided block and with options, to
// create a difference object that represents the changes between them.
//
// # Discussion
//
// The options allow you to choose to omit insertion or removal references to
// the change objects within the difference object’s changes. Don’t use
// the option [NSOrderedCollectionDifferenceCalculationInferMoves] when
// providing a block for the equivalence test. The changes returned in the
// difference object don’t include valid values for [AssociatedIndex].
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/differenceFromOrderedSet:withOptions:usingEquivalenceTest:
func (o NSOrderedSet) DifferenceFromOrderedSetWithOptionsUsingEquivalenceTest(other INSOrderedSet, options NSOrderedCollectionDifferenceCalculationOptions, block ObjectTypeHandler) INSOrderedCollectionDifference {
	_block2, _ := NewObjectTypeBlock(block)
	rv := objc.Send[objc.ID](o.ID, objc.Sel("differenceFromOrderedSet:withOptions:usingEquivalenceTest:"), other, options, _block2)
	return NSOrderedCollectionDifferenceFromID(rv)
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (o NSOrderedSet) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Copies the objects contained in the ordered set that fall within the
// specified range to `objects`.
//
// objects: A C array of objects of size at least the length of the range specified by
// aRange.
//
// range: A range within the bounds of the array.
//
// If the location plus the length of the range is greater than the count of
// the array, this method raises an [RangeException].
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/getObjects:range:
func (o NSOrderedSet) GetObjectsRange(objects []objectivec.IObject, range_ NSRange) {
	objc.Send[objc.ID](o.ID, objc.Sel("getObjects:range:"), objectivec.IObjectSliceToNSArray(objects), range_)
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
func (o NSOrderedSet) InitWithObjects(firstObj objectivec.IObject) NSOrderedSet {
	rv := objc.Send[NSOrderedSet](o.ID, objc.Sel("initWithObjects:"), firstObj)
	return rv
}

// Creates a new ordered set by applying a difference object to an existing
// ordered set.
//
// # Discussion
//
// The following example computes the difference between two ordered sets,
// then applies the difference to create an ordered set that duplicates the
// original:
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/orderedSetByApplyingDifference:
func (o NSOrderedSet) OrderedSetByApplyingDifference(difference INSOrderedCollectionDifference) INSOrderedSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("orderedSetByApplyingDifference:"), difference)
	return NSOrderedSetFromID(rv)
}

// Creates and returns a set containing a specified number of objects from a
// given C array of objects.
//
// objects: A C array of objects to add to the new ordered set.
//
// If the same object appears more than once in objects, it is added only once
// to the returned ordered set. Each object receives a retain message as it is
// added to the set.
//
// cnt: The number of objects from objects to add to the new set.
//
// # Return Value
//
// A new ordered set containing cnt objects from the list of objects specified
// by `objects`.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(objects:count:)-3ny0m
func (_NSOrderedSetClass NSOrderedSetClass) OrderedSetWithObjectsCount(objects []objectivec.IObject, cnt uint) NSOrderedSet {
	rv := objc.Send[objc.ID](objc.ID(_NSOrderedSetClass.class), objc.Sel("orderedSetWithObjects:count:"), objc.CArray(objects), cnt)
	return NSOrderedSetFromID(rv)
}

// Creates and returns an empty ordered set
//
// # Return Value
//
// A new empty ordered set.
//
// # Discussion
//
// This method is declared primarily for the use of mutable subclasses of
// [NSOrderedSet].
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/orderedSet
func (_NSOrderedSetClass NSOrderedSetClass) OrderedSet() NSOrderedSet {
	rv := objc.Send[objc.ID](objc.ID(_NSOrderedSetClass.class), objc.Sel("orderedSet"))
	return NSOrderedSetFromID(rv)
}

// Creates and returns a set containing a uniqued collection of the objects
// contained in a given array.
//
// array: An array containing the objects to add to the new ordered set. If the same
// object appears more than once in `array`, it is added only once to the
// returned set.
//
// # Return Value
//
// A new ordered set containing a uniqued collection of the objects contained
// in array.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/orderedSetWithArray:
func (_NSOrderedSetClass NSOrderedSetClass) OrderedSetWithArray(array []objectivec.IObject) NSOrderedSet {
	rv := objc.Send[objc.ID](objc.ID(_NSOrderedSetClass.class), objc.Sel("orderedSetWithArray:"), objectivec.IObjectSliceToNSArray(array))
	return NSOrderedSetFromID(rv)
}

// Creates and returns a new ordered set for a specified range of objects in
// an array.
//
// array: The array
//
// range: The range of the objects to add to the ordered set.
//
// flag: If true the objects are copied to the ordered set; otherwise false.
//
// # Return Value
//
// A new ordered set containing a uniqued collection of the objects contained
// in the specified range of the array.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/orderedSetWithArray:range:copyItems:
func (_NSOrderedSetClass NSOrderedSetClass) OrderedSetWithArrayRangeCopyItems(array []objectivec.IObject, range_ NSRange, flag bool) NSOrderedSet {
	rv := objc.Send[objc.ID](objc.ID(_NSOrderedSetClass.class), objc.Sel("orderedSetWithArray:range:copyItems:"), objectivec.IObjectSliceToNSArray(array), range_, flag)
	return NSOrderedSetFromID(rv)
}

// Creates and returns a ordered set that contains a single given object.
//
// object: The object to add to the new set.
//
// # Return Value
//
// A new ordered set containing `object`.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/orderedSetWithObject:
func (_NSOrderedSetClass NSOrderedSetClass) OrderedSetWithObject(object objectivec.IObject) NSOrderedSet {
	rv := objc.Send[objc.ID](objc.ID(_NSOrderedSetClass.class), objc.Sel("orderedSetWithObject:"), object)
	return NSOrderedSetFromID(rv)
}

// Creates and returns a ordered set containing the objects in a given
// argument list.
//
// firstObj: The first object to add to the new set.
//
// If the same object appears more than once in the list of objects, it is
// added only once to the returned set. The objects are added to the ordered
// set in the order that they are listed.
//
// # Return Value
//
// A new ordered set containing the objects in the argument list.
//
// # Discussion
//
// To add additional objects to the new set, pass a comma-separated list of
// trailing variadic arguments, ending with `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/orderedSetWithObjects:
func (_NSOrderedSetClass NSOrderedSetClass) OrderedSetWithObjects(firstObj objectivec.IObject) NSOrderedSet {
	rv := objc.Send[objc.ID](objc.ID(_NSOrderedSetClass.class), objc.Sel("orderedSetWithObjects:"), firstObj)
	return NSOrderedSetFromID(rv)
}

// Creates and returns an ordered set containing the objects from another
// ordered set.
//
// set: A set containing the objects to add to the new ordered set.
//
// The objects are not copied, simply referenced.
//
// # Return Value
//
// A new ordered set containing the objects from set.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/orderedSetWithOrderedSet:
func (_NSOrderedSetClass NSOrderedSetClass) OrderedSetWithOrderedSet(set INSOrderedSet) NSOrderedSet {
	rv := objc.Send[objc.ID](objc.ID(_NSOrderedSetClass.class), objc.Sel("orderedSetWithOrderedSet:"), set)
	return NSOrderedSetFromID(rv)
}

// Creates and returns a new ordered set for a specified range of objects in
// an ordered set.
//
// set: An ordered set.
//
// range: The range of objects in `set` to add to the ordered set.
//
// flag: If true the objects are copied to the ordered set; otherwise false.
//
// # Return Value
//
// A new ordered set containing a uniqued collection of the objects contained
// in the specified range of the ordered set.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/orderedSetWithOrderedSet:range:copyItems:
func (_NSOrderedSetClass NSOrderedSetClass) OrderedSetWithOrderedSetRangeCopyItems(set INSOrderedSet, range_ NSRange, flag bool) NSOrderedSet {
	rv := objc.Send[objc.ID](objc.ID(_NSOrderedSetClass.class), objc.Sel("orderedSetWithOrderedSet:range:copyItems:"), set, range_, flag)
	return NSOrderedSetFromID(rv)
}

// Creates and returns an ordered set with the contents of a set.
//
// set: A set.
//
// # Return Value
//
// A new ordered set containing a uniqued collection of the objects contained
// in the set.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/orderedSetWithSet:
func (_NSOrderedSetClass NSOrderedSetClass) OrderedSetWithSet(set INSSet) NSOrderedSet {
	rv := objc.Send[objc.ID](objc.ID(_NSOrderedSetClass.class), objc.Sel("orderedSetWithSet:"), set)
	return NSOrderedSetFromID(rv)
}

// Creates and returns an ordered set with the contents of a set, optionally
// copying the items.
//
// set: A set.
//
// flag: If true the objects are copied to the ordered set; otherwise false.
//
// # Return Value
//
// A new ordered set containing a uniqued collection of the objects contained
// in the specified range of the set.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/orderedSetWithSet:copyItems:
func (_NSOrderedSetClass NSOrderedSetClass) OrderedSetWithSetCopyItems(set INSSet, flag bool) NSOrderedSet {
	rv := objc.Send[objc.ID](objc.ID(_NSOrderedSetClass.class), objc.Sel("orderedSetWithSet:copyItems:"), set, flag)
	return NSOrderedSetFromID(rv)
}

// The number of members in the set.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/count
func (o NSOrderedSet) Count() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("count"))
	return rv
}

// The first object in the ordered set.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/firstObject
func (o NSOrderedSet) FirstObject() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("firstObject"))
	return objectivec.Object{ID: rv}
}

// The last object in the ordered set.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/lastObject
func (o NSOrderedSet) LastObject() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("lastObject"))
	return objectivec.Object{ID: rv}
}

// An ordered set in the reverse order.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/reversed
func (o NSOrderedSet) ReversedOrderedSet() INSOrderedSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("reversedOrderedSet"))
	return NSOrderedSetFromID(objc.ID(rv))
}

// A string that represents the contents of the ordered set, formatted as a
// property list.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/description
func (o NSOrderedSet) Description() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("description"))
	return NSStringFromID(rv).String()
}

// A representation of the ordered set as an array.
//
// # Discussion
//
// This returns a proxy object for the receiving ordered set, which acts like
// an immutable array.
//
// While you cannot mutate the ordered set through this proxy, mutations to
// the original ordered set will be reflected in the proxy and it will appear
// to change spontaneously, because a copy of the ordered set is not being
// made.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/array
func (o NSOrderedSet) Array() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("array"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// A representation of the set containing the contents of the ordered set.
//
// # Discussion
//
// This returns a proxy object for the receiving ordered set, which acts like
// an immutable set.
//
// While you cannot mutate the ordered set through this proxy, mutations to
// the original ordered set will be reflected in the proxy and it will appear
// to change spontaneously, because a copy of the ordered set is not being
// made.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedSet/set
func (o NSOrderedSet) Set() INSSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("set"))
	return NSSetFromID(objc.ID(rv))
}

// Protocol methods for NSCopying

// Protocol methods for NSMutableCopying

// Protocol methods for NSSecureCoding

// EnumerateObjectsAtIndexesOptionsUsingBlockSync is a synchronous wrapper around [NSOrderedSet.EnumerateObjectsAtIndexesOptionsUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (o NSOrderedSet) EnumerateObjectsAtIndexesOptionsUsingBlockSync(ctx context.Context, s INSIndexSet, opts NSEnumerationOptions) (objectivec.IObject, error) {
	done := make(chan objectivec.IObject, 1)
	o.EnumerateObjectsAtIndexesOptionsUsingBlock(s, opts, func(val objectivec.IObject) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// EnumerateObjectsUsingBlockSync is a synchronous wrapper around [NSOrderedSet.EnumerateObjectsUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (o NSOrderedSet) EnumerateObjectsUsingBlockSync(ctx context.Context) (objectivec.IObject, error) {
	done := make(chan objectivec.IObject, 1)
	o.EnumerateObjectsUsingBlock(func(val objectivec.IObject) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// EnumerateObjectsWithOptionsUsingBlockSync is a synchronous wrapper around [NSOrderedSet.EnumerateObjectsWithOptionsUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (o NSOrderedSet) EnumerateObjectsWithOptionsUsingBlockSync(ctx context.Context, opts NSEnumerationOptions) (objectivec.IObject, error) {
	done := make(chan objectivec.IObject, 1)
	o.EnumerateObjectsWithOptionsUsingBlock(opts, func(val objectivec.IObject) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// IndexOfObjectAtIndexesOptionsPassingTestSync is a synchronous wrapper around [NSOrderedSet.IndexOfObjectAtIndexesOptionsPassingTest].
// It blocks until the completion handler fires or the context is cancelled.
func (o NSOrderedSet) IndexOfObjectAtIndexesOptionsPassingTestSync(ctx context.Context, s INSIndexSet, opts NSEnumerationOptions) (objectivec.IObject, error) {
	done := make(chan objectivec.IObject, 1)
	o.IndexOfObjectAtIndexesOptionsPassingTest(s, opts, func(val objectivec.IObject) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// IndexOfObjectPassingTestSync is a synchronous wrapper around [NSOrderedSet.IndexOfObjectPassingTest].
// It blocks until the completion handler fires or the context is cancelled.
func (o NSOrderedSet) IndexOfObjectPassingTestSync(ctx context.Context) (objectivec.IObject, error) {
	done := make(chan objectivec.IObject, 1)
	o.IndexOfObjectPassingTest(func(val objectivec.IObject) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// IndexOfObjectWithOptionsPassingTestSync is a synchronous wrapper around [NSOrderedSet.IndexOfObjectWithOptionsPassingTest].
// It blocks until the completion handler fires or the context is cancelled.
func (o NSOrderedSet) IndexOfObjectWithOptionsPassingTestSync(ctx context.Context, opts NSEnumerationOptions) (objectivec.IObject, error) {
	done := make(chan objectivec.IObject, 1)
	o.IndexOfObjectWithOptionsPassingTest(opts, func(val objectivec.IObject) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// IndexesOfObjectsAtIndexesOptionsPassingTestSync is a synchronous wrapper around [NSOrderedSet.IndexesOfObjectsAtIndexesOptionsPassingTest].
// It blocks until the completion handler fires or the context is cancelled.
func (o NSOrderedSet) IndexesOfObjectsAtIndexesOptionsPassingTestSync(ctx context.Context, s INSIndexSet, opts NSEnumerationOptions) (objectivec.IObject, error) {
	done := make(chan objectivec.IObject, 1)
	o.IndexesOfObjectsAtIndexesOptionsPassingTest(s, opts, func(val objectivec.IObject) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// IndexesOfObjectsPassingTestSync is a synchronous wrapper around [NSOrderedSet.IndexesOfObjectsPassingTest].
// It blocks until the completion handler fires or the context is cancelled.
func (o NSOrderedSet) IndexesOfObjectsPassingTestSync(ctx context.Context) (objectivec.IObject, error) {
	done := make(chan objectivec.IObject, 1)
	o.IndexesOfObjectsPassingTest(func(val objectivec.IObject) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// IndexesOfObjectsWithOptionsPassingTestSync is a synchronous wrapper around [NSOrderedSet.IndexesOfObjectsWithOptionsPassingTest].
// It blocks until the completion handler fires or the context is cancelled.
func (o NSOrderedSet) IndexesOfObjectsWithOptionsPassingTestSync(ctx context.Context, opts NSEnumerationOptions) (objectivec.IObject, error) {
	done := make(chan objectivec.IObject, 1)
	o.IndexesOfObjectsWithOptionsPassingTest(opts, func(val objectivec.IObject) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// DifferenceFromOrderedSetWithOptionsUsingEquivalenceTestSync is a synchronous wrapper around [NSOrderedSet.DifferenceFromOrderedSetWithOptionsUsingEquivalenceTest].
// It blocks until the completion handler fires or the context is cancelled.
func (o NSOrderedSet) DifferenceFromOrderedSetWithOptionsUsingEquivalenceTestSync(ctx context.Context, other INSOrderedSet, options NSOrderedCollectionDifferenceCalculationOptions) (objectivec.IObject, error) {
	done := make(chan objectivec.IObject, 1)
	o.DifferenceFromOrderedSetWithOptionsUsingEquivalenceTest(other, options, func(val objectivec.IObject) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
