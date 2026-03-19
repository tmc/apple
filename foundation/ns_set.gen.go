// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSet] class.
var (
	_NSSetClass     NSSetClass
	_NSSetClassOnce sync.Once
)

func getNSSetClass() NSSetClass {
	_NSSetClassOnce.Do(func() {
		_NSSetClass = NSSetClass{class: objc.GetClass("NSSet")}
	})
	return _NSSetClass
}

// GetNSSetClass returns the class object for NSSet.
func GetNSSetClass() NSSetClass {
	return getNSSetClass()
}

type NSSetClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSetClass) Alloc() NSSet {
	rv := objc.Send[NSSet](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A static, unordered collection of unique objects.
//
// # Overview
// 
// The [NSSet], [NSMutableSet], and [NSCountedSet] classes declare the
// programmatic interface to an unordered collection of objects.
// 
// [NSSet] declares the programmatic interface for static sets of distinct
// objects. You establish a static set’s entries when it’s created, and
// can’t modify the entries after that. [NSMutableSet], on the other hand,
// declares a programmatic interface for dynamic sets of distinct objects. A
// dynamic — or mutable — set allows the addition and deletion of entries
// at any time, automatically allocating memory as needed.
// 
// Use sets as an alternative to arrays when the order of elements isn’t
// important and you need to consider performance in testing whether the set
// contains an object. With an array, testing for membership is slower than
// with sets.
// 
// [NSSet] is “toll-free bridged” with its Core Foundation counterpart,
// [CFSet]. See [Toll-Free Bridging] for more information on toll-free
// bridging.
// 
// In Swift, use this class instead of a [Set] constant in cases where you
// require reference semantics.
// 
// # Subclassing Notes
// 
// There should be little need of subclassing. If you need to customize
// behavior, it’s often better to consider composition instead of
// subclassing.
// 
// # Methods to Override
// 
// In a subclass, you must override all of its primitive methods:
// 
// - [NSSet.Count]
// - [NSSet.Member]
// - [NSSet.ObjectEnumerator]
// 
// # Alternatives to Subclassing
// 
// Before making a custom class of [NSSet], investigate [NSHashTable] and the
// corresponding Core Foundation type, [CFSet]. Because [NSSet] and [CFSet]
// are “toll-free bridged,” you can substitute a [CFSet] object for a
// [NSSet] object in your code (with appropriate casting). Although they’re
// corresponding types, [CFSet] and [NSSet] don’t have identical interfaces
// or implementations, and you can sometimes do things with [CFSet] that you
// can’t easily do with [NSSet].
// 
// If the behavior you want to add supplements that of the existing class, you
// could write a category on [NSSet]. Keep in mind, however, that this
// category affects all instances of [NSSet] that you use, and this might have
// unintended consequences. Alternatively, you could use composition to
// achieve the desired behavior.
//
// [CFSet]: https://developer.apple.com/documentation/CoreFoundation/CFSet
// [Set]: https://developer.apple.com/documentation/Swift/Set
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
//
// # Creating a Set
//
//   - [NSSet.SetByAddingObject]: Returns a new set formed by adding a given object to the receiving set.
//   - [NSSet.SetByAddingObjectsFromSet]: Returns a new set formed by adding the objects in a given set to the receiving set.
//   - [NSSet.SetByAddingObjectsFromArray]: Returns a new set formed by adding the objects in a given array to the receiving set.
//
// # Initializing a Set
//
//   - [NSSet.InitWithArray]: Initializes a newly allocated set with the objects that are contained in a given array.
//   - [NSSet.InitWithObjectsCount]: Initializes a newly allocated set with a specified number of objects from a given C array of objects.
//   - [NSSet.InitWithSet]: Initializes a newly allocated set and adds to it objects from another given set.
//   - [NSSet.InitWithSetCopyItems]: Initializes a newly allocated set and adds to it members of another given set.
//
// # Counting Entries
//
//   - [NSSet.Count]: The number of members in the set.
//
// # Accessing Set Members
//
//   - [NSSet.AllObjects]: An array containing the set’s members, or an empty array if the set has no members.
//   - [NSSet.AnyObject]: Returns one of the objects in the set, or `nil` if the set contains no objects.
//   - [NSSet.ContainsObject]: Returns a Boolean value that indicates whether a given object is present in the set.
//   - [NSSet.FilteredSetUsingPredicate]: Evaluates a given predicate against each object in the receiving set and returns a new set containing the objects for which the predicate returns true.
//   - [NSSet.Member]: Determines whether a given object is present in the set, and returns that object if it is.
//   - [NSSet.ObjectEnumerator]: Returns an enumerator object that lets you access each object in the set.
//   - [NSSet.EnumerateObjectsUsingBlock]: Executes a given block using each object in the set.
//   - [NSSet.EnumerateObjectsWithOptionsUsingBlock]: Executes a given block using each object in the set, using the specified enumeration options.
//   - [NSSet.ObjectsPassingTest]: Returns a set of objects that pass a test in a given block.
//   - [NSSet.ObjectsWithOptionsPassingTest]: Returns a set of objects that pass a test in a given block, using the specified enumeration options.
//
// # Comparing Sets
//
//   - [NSSet.IsSubsetOfSet]: Returns a Boolean value that indicates whether every object in the receiving set is also present in another given set.
//   - [NSSet.IntersectsSet]: Returns a Boolean value that indicates whether at least one object in the receiving set is also present in another given set.
//   - [NSSet.IsEqualToSet]: Compares the receiving set to another set.
//
// # Creating a Sorted Array
//
//   - [NSSet.SortedArrayUsingDescriptors]: Returns an array of the set’s content sorted as specified by a given array of sort descriptors.
//
// # Describing a Set
//
//   - [NSSet.Description]: A string that represents the contents of the set, formatted as a property list.
//   - [NSSet.DescriptionWithLocale]: Returns a string that represents the contents of the set, formatted as a property list.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet
type NSSet struct {
	objectivec.Object
}

// NSSetFromID constructs a [NSSet] from an objc.ID.
//
// A static, unordered collection of unique objects.
func NSSetFromID(id objc.ID) NSSet {
	return NSSet{objectivec.Object{ID: id}}
}
// NOTE: NSSet adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSet] class.
//
// # Creating a Set
//
//   - [INSSet.SetByAddingObject]: Returns a new set formed by adding a given object to the receiving set.
//   - [INSSet.SetByAddingObjectsFromSet]: Returns a new set formed by adding the objects in a given set to the receiving set.
//   - [INSSet.SetByAddingObjectsFromArray]: Returns a new set formed by adding the objects in a given array to the receiving set.
//
// # Initializing a Set
//
//   - [INSSet.InitWithArray]: Initializes a newly allocated set with the objects that are contained in a given array.
//   - [INSSet.InitWithObjectsCount]: Initializes a newly allocated set with a specified number of objects from a given C array of objects.
//   - [INSSet.InitWithSet]: Initializes a newly allocated set and adds to it objects from another given set.
//   - [INSSet.InitWithSetCopyItems]: Initializes a newly allocated set and adds to it members of another given set.
//
// # Counting Entries
//
//   - [INSSet.Count]: The number of members in the set.
//
// # Accessing Set Members
//
//   - [INSSet.AllObjects]: An array containing the set’s members, or an empty array if the set has no members.
//   - [INSSet.AnyObject]: Returns one of the objects in the set, or `nil` if the set contains no objects.
//   - [INSSet.ContainsObject]: Returns a Boolean value that indicates whether a given object is present in the set.
//   - [INSSet.FilteredSetUsingPredicate]: Evaluates a given predicate against each object in the receiving set and returns a new set containing the objects for which the predicate returns true.
//   - [INSSet.Member]: Determines whether a given object is present in the set, and returns that object if it is.
//   - [INSSet.ObjectEnumerator]: Returns an enumerator object that lets you access each object in the set.
//   - [INSSet.EnumerateObjectsUsingBlock]: Executes a given block using each object in the set.
//   - [INSSet.EnumerateObjectsWithOptionsUsingBlock]: Executes a given block using each object in the set, using the specified enumeration options.
//   - [INSSet.ObjectsPassingTest]: Returns a set of objects that pass a test in a given block.
//   - [INSSet.ObjectsWithOptionsPassingTest]: Returns a set of objects that pass a test in a given block, using the specified enumeration options.
//
// # Comparing Sets
//
//   - [INSSet.IsSubsetOfSet]: Returns a Boolean value that indicates whether every object in the receiving set is also present in another given set.
//   - [INSSet.IntersectsSet]: Returns a Boolean value that indicates whether at least one object in the receiving set is also present in another given set.
//   - [INSSet.IsEqualToSet]: Compares the receiving set to another set.
//
// # Creating a Sorted Array
//
//   - [INSSet.SortedArrayUsingDescriptors]: Returns an array of the set’s content sorted as specified by a given array of sort descriptors.
//
// # Describing a Set
//
//   - [INSSet.Description]: A string that represents the contents of the set, formatted as a property list.
//   - [INSSet.DescriptionWithLocale]: Returns a string that represents the contents of the set, formatted as a property list.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet
type INSSet interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSMutableCopying
	NSSecureCoding

	// Topic: Creating a Set

	// Returns a new set formed by adding a given object to the receiving set.
	SetByAddingObject(anObject objectivec.IObject) INSSet
	// Returns a new set formed by adding the objects in a given set to the receiving set.
	SetByAddingObjectsFromSet(other INSSet) INSSet
	// Returns a new set formed by adding the objects in a given array to the receiving set.
	SetByAddingObjectsFromArray(other []objectivec.IObject) INSSet

	// Topic: Initializing a Set

	// Initializes a newly allocated set with the objects that are contained in a given array.
	InitWithArray(array []objectivec.IObject) NSSet
	// Initializes a newly allocated set with a specified number of objects from a given C array of objects.
	InitWithObjectsCount(objects []objectivec.IObject, cnt uint) NSSet
	// Initializes a newly allocated set and adds to it objects from another given set.
	InitWithSet(set INSSet) NSSet
	// Initializes a newly allocated set and adds to it members of another given set.
	InitWithSetCopyItems(set INSSet, flag bool) NSSet

	// Topic: Counting Entries

	// The number of members in the set.
	Count() uint

	// Topic: Accessing Set Members

	// An array containing the set’s members, or an empty array if the set has no members.
	AllObjects() []objectivec.IObject
	// Returns one of the objects in the set, or `nil` if the set contains no objects.
	AnyObject() objectivec.IObject
	// Returns a Boolean value that indicates whether a given object is present in the set.
	ContainsObject(anObject objectivec.IObject) bool
	// Evaluates a given predicate against each object in the receiving set and returns a new set containing the objects for which the predicate returns true.
	FilteredSetUsingPredicate(predicate INSPredicate) INSSet
	// Determines whether a given object is present in the set, and returns that object if it is.
	Member(object objectivec.IObject) objectivec.IObject
	// Returns an enumerator object that lets you access each object in the set.
	ObjectEnumerator() INSEnumerator
	// Executes a given block using each object in the set.
	EnumerateObjectsUsingBlock(block ObjectTypeHandler)
	// Executes a given block using each object in the set, using the specified enumeration options.
	EnumerateObjectsWithOptionsUsingBlock(opts NSEnumerationOptions, block ObjectTypeHandler)
	// Returns a set of objects that pass a test in a given block.
	ObjectsPassingTest(predicate ObjectTypeHandler) INSSet
	// Returns a set of objects that pass a test in a given block, using the specified enumeration options.
	ObjectsWithOptionsPassingTest(opts NSEnumerationOptions, predicate ObjectTypeHandler) INSSet

	// Topic: Comparing Sets

	// Returns a Boolean value that indicates whether every object in the receiving set is also present in another given set.
	IsSubsetOfSet(otherSet INSSet) bool
	// Returns a Boolean value that indicates whether at least one object in the receiving set is also present in another given set.
	IntersectsSet(otherSet INSSet) bool
	// Compares the receiving set to another set.
	IsEqualToSet(otherSet INSSet) bool

	// Topic: Creating a Sorted Array

	// Returns an array of the set’s content sorted as specified by a given array of sort descriptors.
	SortedArrayUsingDescriptors(sortDescriptors []NSSortDescriptor) []objectivec.IObject

	// Topic: Describing a Set

	// A string that represents the contents of the set, formatted as a property list.
	Description() string
	// Returns a string that represents the contents of the set, formatted as a property list.
	DescriptionWithLocale(locale objectivec.IObject) string

	// Initializes a newly allocated set with members taken from the specified list of objects.
	InitWithObjects(firstObj objectivec.IObject) NSSet
	// Sends a message specified by a given selector to each object in the set.
	MakeObjectsPerformSelector(aSelector objc.SEL)
	// Sends a message specified by a given selector to each object in the set.
	MakeObjectsPerformSelectorWithObject(aSelector objc.SEL, argument objectivec.IObject)
}

// Init initializes the instance.
func (s NSSet) Init() NSSet {
	rv := objc.Send[NSSet](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSet) Autorelease() NSSet {
	rv := objc.Send[NSSet](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSet creates a new NSSet instance.
func NewNSSet() NSSet {
	class := getNSSetClass()
	rv := objc.Send[NSSet](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a newly allocated set with the objects that are contained in a
// given array.
//
// array: An array of objects to add to the new set. If the same object appears more
// than once in `array`, it is represented only once in the returned set. Each
// object receives a [retain] message as it is added to the set.
// //
// [retain]: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/retain
//
// # Return Value
// 
// An initialized set with the contents of `array`. The returned set might be
// different than the original receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/init(array:)
func NewSetWithArray(array []objectivec.IObject) NSSet {
	instance := getNSSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArray:"), objectivec.IObjectSliceToNSArray(array))
	return NSSetFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSSet/init(coder:)
func NewSetWithCoder(coder INSCoder) NSSet {
	instance := getNSSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSSetFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSSet/init(collectionViewIndexPath:)
func NewSetWithCollectionViewIndexPath(indexPath objectivec.IObject) NSSet {
	rv := objc.Send[objc.ID](objc.ID(getNSSetClass().class), objc.Sel("setWithCollectionViewIndexPath:"), indexPath)
	return NSSetFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSSet/init(collectionViewIndexPaths:)
func NewSetWithCollectionViewIndexPaths(indexPaths []objc.ID) NSSet {
	rv := objc.Send[objc.ID](objc.ID(getNSSetClass().class), objc.Sel("setWithCollectionViewIndexPaths:"), objectivec.IDSliceToNSArray(indexPaths))
	return NSSetFromID(rv)
}

// Creates and returns a set that contains a single given object.
//
// object: The object to add to the new set. `object` receives a [retain] message
// after being added to the set.
// //
// [retain]: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/retain
//
// # Return Value
// 
// A new set that contains a single member, `object`.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/init(object:)
func NewSetWithObject(object objectivec.IObject) NSSet {
	rv := objc.Send[objc.ID](objc.ID(getNSSetClass().class), objc.Sel("setWithObject:"), object)
	return NSSetFromID(rv)
}

// Initializes a newly allocated set with members taken from the specified
// list of objects.
//
// firstObj: The first object to add to the new set.
//
// # Return Value
// 
// An initialized set containing the objects specified in the parameter list.
// The returned set might be different than the original receiver.
//
// # Discussion
// 
// To add additional objects to the new set, pass a comma-separated list of
// trailing variadic arguments, ending with `nil`. If the same object appears
// more than once in the list of objects, it is added only once to the
// returned set. Each object receives a [retain] message as it is added to the
// set.
//
// [retain]: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/retain
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/initWithObjects:
func NewSetWithObjects(firstObj objectivec.IObject) NSSet {
	instance := getNSSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObjects:"), firstObj)
	return NSSetFromID(rv)
}

// Initializes a newly allocated set and adds to it objects from another given
// set.
//
// set: A set containing objects to add to the receiving set. Each object is
// retained as it is added.
//
// # Return Value
// 
// An initialized objects set containing the objects from `set`. The returned
// set might be different than the original receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/init(set:)-1xovx
func NewSetWithSet(set INSSet) NSSet {
	instance := getNSSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSet:"), set)
	return NSSetFromID(rv)
}

// Initializes a newly allocated set and adds to it members of another given
// set.
//
// set: A set containing objects to add to the new set.
//
// flag: If [true], each object in `set` receives a [copyWithZone:] message to
// create a copy of the object—objects must conform to the [NSCopying]
// protocol. In a managed memory environment, this is instead of the `retain`
// message the object would otherwise receive. The object copy is then added
// to the returned set.
// 
// If [false], then in a managed memory environment each object in `set`
// simply receives a `retain` message when it is added to the returned set.
// //
// [copyWithZone:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/copyWithZone:
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An initialized set that contains the members of `set`. The returned set
// might be different than the original receiver.
//
// # Discussion
// 
// After an immutable s has been initialized in this way, it cannot be
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
// See: https://developer.apple.com/documentation/Foundation/NSSet/init(set:copyItems:)
func NewSetWithSetCopyItems(set INSSet, flag bool) NSSet {
	instance := getNSSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSet:copyItems:"), set, flag)
	return NSSetFromID(rv)
}

// Returns a new set formed by adding a given object to the receiving set.
//
// anObject: The object to add to the set.
//
// # Return Value
// 
// A new set formed by adding `anObject` to the receiving set.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/adding(_:)
func (s NSSet) SetByAddingObject(anObject objectivec.IObject) INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("setByAddingObject:"), anObject)
	return NSSetFromID(rv)
}
// Returns a new set formed by adding the objects in a given set to the
// receiving set.
//
// other: The set of objects to add to the receiving set.
//
// # Return Value
// 
// A new set formed by adding the objects in `other` to the receiving set.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/addingObjects(from:)-2i31h
func (s NSSet) SetByAddingObjectsFromSet(other INSSet) INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("setByAddingObjectsFromSet:"), other)
	return NSSetFromID(rv)
}
// Returns a new set formed by adding the objects in a given array to the
// receiving set.
//
// other: The array of objects to add to the set.
//
// # Return Value
// 
// A new set formed by adding the objects in `other` to the receiving set.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/addingObjects(from:)-544m9
func (s NSSet) SetByAddingObjectsFromArray(other []objectivec.IObject) INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("setByAddingObjectsFromArray:"), objectivec.IObjectSliceToNSArray(other))
	return NSSetFromID(rv)
}
// Initializes a newly allocated set with the objects that are contained in a
// given array.
//
// array: An array of objects to add to the new set. If the same object appears more
// than once in `array`, it is represented only once in the returned set. Each
// object receives a [retain] message as it is added to the set.
// //
// [retain]: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/retain
//
// # Return Value
// 
// An initialized set with the contents of `array`. The returned set might be
// different than the original receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/init(array:)
func (s NSSet) InitWithArray(array []objectivec.IObject) NSSet {
	rv := objc.Send[NSSet](s.ID, objc.Sel("initWithArray:"), objectivec.IObjectSliceToNSArray(array))
	return rv
}
// Initializes a newly allocated set with a specified number of objects from a
// given C array of objects.
//
// objects: A C array of objects to add to the new set. If the same object appears more
// than once in `objects`, it is added only once to the returned set. Each
// object receives a [retain] message as it is added to the set.
// //
// [retain]: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/retain
//
// cnt: The number of objects from `objects` to add to the new set.
//
// # Return Value
// 
// An initialized set containing `cnt` objects from the list of objects
// specified by `objects`. The returned set might be different than the
// original receiver.
//
// # Discussion
// 
// This method is a designated initializer for [NSSet].
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/init(objects:count:)-7kift
func (s NSSet) InitWithObjectsCount(objects []objectivec.IObject, cnt uint) NSSet {
	rv := objc.Send[NSSet](s.ID, objc.Sel("initWithObjects:count:"), objc.CArray(objects), cnt)
	return rv
}
// Initializes a newly allocated set and adds to it objects from another given
// set.
//
// set: A set containing objects to add to the receiving set. Each object is
// retained as it is added.
//
// # Return Value
// 
// An initialized objects set containing the objects from `set`. The returned
// set might be different than the original receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/init(set:)-1xovx
func (s NSSet) InitWithSet(set INSSet) NSSet {
	rv := objc.Send[NSSet](s.ID, objc.Sel("initWithSet:"), set)
	return rv
}
// Initializes a newly allocated set and adds to it members of another given
// set.
//
// set: A set containing objects to add to the new set.
//
// flag: If [true], each object in `set` receives a [copyWithZone:] message to
// create a copy of the object—objects must conform to the [NSCopying]
// protocol. In a managed memory environment, this is instead of the `retain`
// message the object would otherwise receive. The object copy is then added
// to the returned set.
// 
// If [false], then in a managed memory environment each object in `set`
// simply receives a `retain` message when it is added to the returned set.
// //
// [copyWithZone:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/copyWithZone:
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An initialized set that contains the members of `set`. The returned set
// might be different than the original receiver.
//
// # Discussion
// 
// After an immutable s has been initialized in this way, it cannot be
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
// See: https://developer.apple.com/documentation/Foundation/NSSet/init(set:copyItems:)
func (s NSSet) InitWithSetCopyItems(set INSSet, flag bool) NSSet {
	rv := objc.Send[NSSet](s.ID, objc.Sel("initWithSet:copyItems:"), set, flag)
	return rv
}
// Returns one of the objects in the set, or `nil` if the set contains no
// objects.
//
// # Return Value
// 
// One of the objects in the set, or `nil` if the set contains no objects. The
// object returned is chosen at the set’s convenience—the selection is not
// guaranteed to be random.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/anyObject()
func (s NSSet) AnyObject() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("anyObject"))
	return objectivec.Object{ID: rv}
}
// Returns a Boolean value that indicates whether a given object is present in
// the set.
//
// anObject: An object to look for in the set.
//
// # Return Value
// 
// [true] if `anObject` is present in the set, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Each element of the set is checked for equality with `anObject` until a
// match is found or the end of the set is reached. Objects are considered
// equal if [isEqual(_:)] returns [true].
//
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/contains(_:)
func (s NSSet) ContainsObject(anObject objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("containsObject:"), anObject)
	return rv
}
// Evaluates a given predicate against each object in the receiving set and
// returns a new set containing the objects for which the predicate returns
// true.
//
// predicate: A predicate.
//
// # Return Value
// 
// A new set containing the objects in the receiving set for which `predicate`
// returns true.
//
// # Discussion
// 
// The following example illustrates the use of this method.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/filtered(using:)
func (s NSSet) FilteredSetUsingPredicate(predicate INSPredicate) INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("filteredSetUsingPredicate:"), predicate)
	return NSSetFromID(rv)
}
// Determines whether a given object is present in the set, and returns that
// object if it is.
//
// object: An object to look for in the set.
//
// # Return Value
// 
// Returns an object equal to `object` if it’s present in the set, otherwise
// `nil`.
//
// # Discussion
// 
// Each element of the set is checked for equality with `object` until a match
// is found or the end of the set is reached. Objects are considered equal if
// [isEqual(_:)] returns [true].
//
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/member(_:)
func (s NSSet) Member(object objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("member:"), object)
	return objectivec.Object{ID: rv}
}
// Returns an enumerator object that lets you access each object in the set.
//
// # Return Value
// 
// An enumerator object that lets you access each object in the set.
//
// # Discussion
// 
// The following code fragment illustrates how you can use this method.
// 
// When this method is used with mutable subclasses of [NSSet], your code
// shouldn’t modify the set during enumeration. If you intend to modify the
// set, use the [AllObjects] method to create a “snapshot” of the set’s
// members. Enumerate the snapshot, but make your modifications to the
// original set.
// 
// # Special Considerations
// 
// It is more efficient to use the fast enumeration protocol (see
// [NSFastEnumeration]). Fast enumeration is available in macOS 10.5 and later
// and iOS 2.0 and later.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/objectEnumerator()
func (s NSSet) ObjectEnumerator() INSEnumerator {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("objectEnumerator"))
	return NSEnumeratorFromID(rv)
}
// Executes a given block using each object in the set.
//
// block: The block to apply to elements in the set.
// 
// The block takes two arguments:
// 
// obj: The element in the set. stop: A reference to a Boolean value. The
// block can set the value to [true] to stop further processing of the set.
// The `stop` argument is an out-only argument. You should only ever set this
// Boolean to [true] within the block.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/enumerateObjects(_:)
func (s NSSet) EnumerateObjectsUsingBlock(block ObjectTypeHandler) {
_block0, _cleanup0 := NewObjectTypeBlock(block)
	defer _cleanup0()
	objc.Send[objc.ID](s.ID, objc.Sel("enumerateObjectsUsingBlock:"), _block0)
}
// Executes a given block using each object in the set, using the specified
// enumeration options.
//
// opts: A bitmask that specifies the options for the enumeration.
//
// block: The block to apply to elements in the set.
// 
// The block takes two arguments:
// 
// obj: The element in the set. stop: A reference to a Boolean value. The
// block can set the value to [true] to stop further processing of the set.
// The `stop` argument is an out-only argument. You should only ever set this
// Boolean to [true] within the block.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/enumerateObjects(options:using:)
func (s NSSet) EnumerateObjectsWithOptionsUsingBlock(opts NSEnumerationOptions, block ObjectTypeHandler) {
_block1, _cleanup1 := NewObjectTypeBlock(block)
	defer _cleanup1()
	objc.Send[objc.ID](s.ID, objc.Sel("enumerateObjectsWithOptions:usingBlock:"), opts, _block1)
}
// Returns a set of objects that pass a test in a given block.
//
// predicate: The block to apply to elements in the array.
// 
// The block takes two arguments:
// 
// obj: The element in the set. stop: A reference to a Boolean value. The
// block can set the value to [true] to stop further processing of the set.
// The `stop` argument is an out-only argument. You should only ever set this
// Boolean to [true] within the block.
// 
// The block returns a Boolean value that indicates whether `obj` passed the
// test.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An [NSSet] containing objects that pass the test.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/objects(passingTest:)
func (s NSSet) ObjectsPassingTest(predicate ObjectTypeHandler) INSSet {
_block0, _cleanup0 := NewObjectTypeBlock(predicate)
	defer _cleanup0()
	rv := objc.Send[objc.ID](s.ID, objc.Sel("objectsPassingTest:"), _block0)
	return NSSetFromID(rv)
}
// Returns a set of objects that pass a test in a given block, using the
// specified enumeration options.
//
// opts: A bitmask that specifies the options for the enumeration.
//
// predicate: The block to apply to elements in the set.
// 
// The block takes two arguments:
// 
// obj: The element in the set. stop: A reference to a Boolean value. The
// block can set the value to [true] to stop further processing of the set.
// The `stop` argument is an out-only argument. You should only ever set this
// Boolean to [true] within the block.
// 
// The block returns a Boolean value that indicates whether `obj` passed the
// test.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An [NSSet] containing objects that pass the test.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/objects(options:passingTest:)
func (s NSSet) ObjectsWithOptionsPassingTest(opts NSEnumerationOptions, predicate ObjectTypeHandler) INSSet {
_block1, _cleanup1 := NewObjectTypeBlock(predicate)
	defer _cleanup1()
	rv := objc.Send[objc.ID](s.ID, objc.Sel("objectsWithOptions:passingTest:"), opts, _block1)
	return NSSetFromID(rv)
}
// Returns a Boolean value that indicates whether every object in the
// receiving set is also present in another given set.
//
// otherSet: The set with which to compare the receiving set.
//
// # Return Value
// 
// [true] if every object in the receiving set is also present in `otherSet`,
// otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Object equality is tested using isEqual:.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/isSubset(of:)
func (s NSSet) IsSubsetOfSet(otherSet INSSet) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isSubsetOfSet:"), otherSet)
	return rv
}
// Returns a Boolean value that indicates whether at least one object in the
// receiving set is also present in another given set.
//
// otherSet: The set with which to compare the receiving set.
//
// # Return Value
// 
// [true] if at least one object in the receiving set is also present in
// `otherSet`, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Object equality is tested using isEqual:.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/intersects(_:)
func (s NSSet) IntersectsSet(otherSet INSSet) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("intersectsSet:"), otherSet)
	return rv
}
// Compares the receiving set to another set.
//
// otherSet: The set with which to compare the receiving set.
//
// # Return Value
// 
// [true] if the contents of `otherSet` are equal to the contents of the
// receiving set, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Two sets have equal contents if they each have the same number of members
// and if each member of one set is present in the other. Object equality is
// tested using isEqual:.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/isEqual(to:)
func (s NSSet) IsEqualToSet(otherSet INSSet) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isEqualToSet:"), otherSet)
	return rv
}
// Returns an array of the set’s content sorted as specified by a given
// array of sort descriptors.
//
// sortDescriptors: An array of [NSSortDescriptor] objects.
//
// # Return Value
// 
// An NSArray containing the set’s content sorted as specified by
// `sortDescriptors`.
//
// # Discussion
// 
// The first descriptor specifies the primary key path to be used in sorting
// the set’s contents. Any subsequent descriptors are used to further refine
// sorting of objects with duplicate values. See [NSSortDescriptor] for
// additional information.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/sortedArray(using:)
func (s NSSet) SortedArrayUsingDescriptors(sortDescriptors []NSSortDescriptor) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("sortedArrayUsingDescriptors:"), objectivec.IObjectSliceToNSArray(sortDescriptors))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
// Returns a string that represents the contents of the set, formatted as a
// property list.
//
// locale: On iOS and macOS 10.5 and later, either an instance of [NSDictionary] or an
// [NSLocale] object may be used for `locale`.In OS X v10.4 and earlier it
// must be an instance of [NSDictionary].
//
// # Return Value
// 
// A string that represents the contents of the set, formatted as a property
// list.
//
// # Discussion
// 
// This method sends each of the set’s members `` with `locale` passed as
// the sole parameter. If the set’s members do not respond to ``, this
// method sends [description] instead.
//
// [description]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/description
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/description(withLocale:)
func (s NSSet) DescriptionWithLocale(locale objectivec.IObject) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("descriptionWithLocale:"), locale)
	return NSStringFromID(rv).String()
}
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/init(coder:)
func (s NSSet) InitWithCoder(coder INSCoder) NSSet {
	rv := objc.Send[NSSet](s.ID, objc.Sel("initWithCoder:"), coder)
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
func (s NSSet) CountByEnumeratingWithStateObjectsCount(state NSFastEnumerationState, buffer []objectivec.IObject, len_ uint) uint {
	rv := objc.Send[uint](s.ID, objc.Sel("countByEnumeratingWithState:objects:count:"), state, objc.CArray(buffer), len_)
	return rv
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (s NSSet) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}
// Initializes a newly allocated set with members taken from the specified
// list of objects.
//
// firstObj: The first object to add to the new set.
//
// # Return Value
// 
// An initialized set containing the objects specified in the parameter list.
// The returned set might be different than the original receiver.
//
// # Discussion
// 
// To add additional objects to the new set, pass a comma-separated list of
// trailing variadic arguments, ending with `nil`. If the same object appears
// more than once in the list of objects, it is added only once to the
// returned set. Each object receives a [retain] message as it is added to the
// set.
//
// [retain]: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/retain
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/initWithObjects:
func (s NSSet) InitWithObjects(firstObj objectivec.IObject) NSSet {
	rv := objc.Send[NSSet](s.ID, objc.Sel("initWithObjects:"), firstObj)
	return rv
}
// Sends a message specified by a given selector to each object in the set.
//
// aSelector: A selector that specifies the message to send to the members of the set.
// The method must not take any arguments. It should not have the side effect
// of modifying the set. This value must not be [NULL].
//
// # Discussion
// 
// The message specified by `aSelector` is sent once to each member of the
// set. This method raises an [NSInvalidArgumentException] if `aSelector` is
// [NULL].
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/makeObjectsPerformSelector:
func (s NSSet) MakeObjectsPerformSelector(aSelector objc.SEL) {
	objc.Send[objc.ID](s.ID, objc.Sel("makeObjectsPerformSelector:"), aSelector)
}
// Sends a message specified by a given selector to each object in the set.
//
// aSelector: A selector that specifies the message to send to the set’s members. The
// method must take a single argument of type `id`. The method should not, as
// a side effect, modify the set. The value must not be [NULL].
//
// argument: The object to pass as an argument to the method specified by `aSelector`.
//
// # Discussion
// 
// The message specified by `aSelector` is sent, with `argument` as the
// argument, once to each member of the set. This method raises an
// [NSInvalidArgumentException] if `aSelector` is [NULL].
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/makeObjectsPerformSelector:withObject:
func (s NSSet) MakeObjectsPerformSelectorWithObject(aSelector objc.SEL, argument objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("makeObjectsPerformSelector:withObject:"), aSelector, argument)
}

// Creates and returns a set containing a specified number of objects from a
// given C array of objects.
//
// objects: A C array of objects to add to the new set. If the same object appears more
// than once in `objects`, it is added only once to the returned set. Each
// object receives a [retain] message as it is added to the set.
// //
// [retain]: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/retain
//
// cnt: The number of objects from `objects` to add to the new set.
//
// # Return Value
// 
// A new set containing `cnt` objects from the list of objects specified by
// `objects`.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/init(objects:count:)-65ni4
func (_NSSetClass NSSetClass) SetWithObjectsCount(objects []objectivec.IObject, cnt uint) NSSet {
	rv := objc.Send[objc.ID](objc.ID(_NSSetClass.class), objc.Sel("setWithObjects:count:"), objc.CArray(objects), cnt)
	return NSSetFromID(rv)
}
// Creates and returns an empty set.
//
// # Return Value
// 
// A new empty set.
//
// # Discussion
// 
// This method is declared primarily for the use of mutable subclasses of
// [NSSet].
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/set
func (_NSSetClass NSSetClass) Set() NSSet {
	rv := objc.Send[objc.ID](objc.ID(_NSSetClass.class), objc.Sel("set"))
	return NSSetFromID(rv)
}
// Creates and returns a set containing a uniqued collection of the objects
// contained in a given array.
//
// array: An array containing the objects to add to the new set. If the same object
// appears more than once in `array`, it is added only once to the returned
// set. Each object receives a [retain] message as it is added to the set.
// //
// [retain]: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/retain
//
// # Return Value
// 
// A new set containing a uniqued collection of the objects contained in
// `array`.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/setWithArray:
func (_NSSetClass NSSetClass) SetWithArray(array []objectivec.IObject) NSSet {
	rv := objc.Send[objc.ID](objc.ID(_NSSetClass.class), objc.Sel("setWithArray:"), objectivec.IObjectSliceToNSArray(array))
	return NSSetFromID(rv)
}
// Creates and returns a set containing the objects in a given argument list.
//
// firstObj: The first object to add to the new set.
//
// # Return Value
// 
// A new set containing the objects in the argument list.
//
// # Discussion
// 
// To add additional objects to the new set, pass a comma-separated list of
// trailing variadic arguments, ending with `nil`. If the same object appears
// more than once in the list of objects, it is added only once to the
// returned set. Each object receives a [retain] message as it is added to the
// set.
// 
// As an example, the following code excerpt creates a set containing three
// different types of elements (assuming `aPath` exits):
//
// [retain]: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/retain
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/setWithObjects:
func (_NSSetClass NSSetClass) SetWithObjects(firstObj objectivec.IObject) NSSet {
	rv := objc.Send[objc.ID](objc.ID(_NSSetClass.class), objc.Sel("setWithObjects:"), firstObj)
	return NSSetFromID(rv)
}
// Creates and returns a set containing the objects from another set.
//
// set: A set containing the objects to add to the new set. Each object receives a
// [retain] message as it is added to the new set.
// //
// [retain]: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/retain
//
// # Return Value
// 
// A new set containing the objects from `set`.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/setWithSet:
func (_NSSetClass NSSetClass) SetWithSet(set INSSet) NSSet {
	rv := objc.Send[objc.ID](objc.ID(_NSSetClass.class), objc.Sel("setWithSet:"), set)
	return NSSetFromID(rv)
}

// The number of members in the set.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/count
func (s NSSet) Count() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("count"))
	return rv
}
// An array containing the set’s members, or an empty array if the set has
// no members.
//
// # Discussion
// 
// The order of the objects in the array is undefined.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/allObjects
func (s NSSet) AllObjects() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("allObjects"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
// A string that represents the contents of the set, formatted as a property
// list.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/description
func (s NSSet) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return NSStringFromID(rv).String()
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSMutableCopying
			

			// Protocol methods for NSSecureCoding
			

// EnumerateObjectsUsingBlockSync is a synchronous wrapper around [NSSet.EnumerateObjectsUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s NSSet) EnumerateObjectsUsingBlockSync(ctx context.Context) (objectivec.IObject, error) {
	done := make(chan objectivec.IObject, 1)
	s.EnumerateObjectsUsingBlock(func(val objectivec.IObject) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// EnumerateObjectsWithOptionsUsingBlockSync is a synchronous wrapper around [NSSet.EnumerateObjectsWithOptionsUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s NSSet) EnumerateObjectsWithOptionsUsingBlockSync(ctx context.Context, opts NSEnumerationOptions) (objectivec.IObject, error) {
	done := make(chan objectivec.IObject, 1)
	s.EnumerateObjectsWithOptionsUsingBlock(opts, func(val objectivec.IObject) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// ObjectsPassingTestSync is a synchronous wrapper around [NSSet.ObjectsPassingTest].
// It blocks until the completion handler fires or the context is cancelled.
func (s NSSet) ObjectsPassingTestSync(ctx context.Context) (objectivec.IObject, error) {
	done := make(chan objectivec.IObject, 1)
	s.ObjectsPassingTest(func(val objectivec.IObject) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// ObjectsWithOptionsPassingTestSync is a synchronous wrapper around [NSSet.ObjectsWithOptionsPassingTest].
// It blocks until the completion handler fires or the context is cancelled.
func (s NSSet) ObjectsWithOptionsPassingTestSync(ctx context.Context, opts NSEnumerationOptions) (objectivec.IObject, error) {
	done := make(chan objectivec.IObject, 1)
	s.ObjectsWithOptionsPassingTest(opts, func(val objectivec.IObject) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

