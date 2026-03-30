// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCountedSet] class.
var (
	_NSCountedSetClass     NSCountedSetClass
	_NSCountedSetClassOnce sync.Once
)

func getNSCountedSetClass() NSCountedSetClass {
	_NSCountedSetClassOnce.Do(func() {
		_NSCountedSetClass = NSCountedSetClass{class: objc.GetClass("NSCountedSet")}
	})
	return _NSCountedSetClass
}

// GetNSCountedSetClass returns the class object for NSCountedSet.
func GetNSCountedSetClass() NSCountedSetClass {
	return getNSCountedSetClass()
}

type NSCountedSetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSCountedSetClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCountedSetClass) Alloc() NSCountedSet {
	rv := objc.Send[NSCountedSet](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A mutable, unordered collection of distinct objects that may appear more
// than once in the collection.
//
// # Overview
//
// Each distinct object inserted into an [NSCountedSet] object has a counter
// associated with it. [NSCountedSet] keeps track of the number of times
// objects are inserted and requires that objects be removed the same number
// of times. Thus, there is only one instance of an object in an [NSSet]
// object even if the object has been added to the set multiple times. The
// [NSCountedSet.Count] method defined by the superclass [NSSet] has special significance;
// it returns the number of distinct objects, not the total number of times
// objects are represented in the set. The [NSSet] and [NSMutableSet] classes
// are provided for static and dynamic sets, respectively, whose elements are
// distinct.
//
// While [NSCountedSet] and [CFBag] are not toll-free bridged, they provide
// similar functionality. For more information about [CFBag], see the [CFBag].
//
// # Subclassing Notes
//
// Because [NSCountedSet] is not a class cluster, it does not have primitive
// methods that provide the basis for its implementation. In general, there
// should be little need for subclassing.
//
// # Methods to Override
//
// If you subclass [NSCountedSet], you must override any method of which you
// want to change the behavior.
//
// If you change the primitive behavior of an [NSCountedSet], for instance if
// you change how objects are stored, you must override all of the affected
// methods. These include:
//
// - [NSCountedSet.AddObject]
// - [NSCountedSet.RemoveObject]
// - [NSCountedSet.ObjectEnumerator]
// - [NSCountedSet.CountForObject]
//
// If you change the primitive behavior, you must also override the primitive
// methods of [NSSet] and [NSMutableSet].
//
// # Examining a Counted Set
//
//   - [NSCountedSet.CountForObject]: Returns the count associated with a given object in the set.
//
// See: https://developer.apple.com/documentation/Foundation/NSCountedSet
//
// [CFBag]: https://developer.apple.com/documentation/CoreFoundation/CFBag
type NSCountedSet struct {
	NSMutableSet
}

// NSCountedSetFromID constructs a [NSCountedSet] from an objc.ID.
//
// A mutable, unordered collection of distinct objects that may appear more
// than once in the collection.
func NSCountedSetFromID(id objc.ID) NSCountedSet {
	return NSCountedSet{NSMutableSet: NSMutableSetFromID(id)}
}

// NOTE: NSCountedSet adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCountedSet] class.
//
// # Examining a Counted Set
//
//   - [INSCountedSet.CountForObject]: Returns the count associated with a given object in the set.
//
// See: https://developer.apple.com/documentation/Foundation/NSCountedSet
type INSCountedSet interface {
	INSMutableSet

	// Topic: Examining a Counted Set

	// Returns the count associated with a given object in the set.
	CountForObject(object objectivec.IObject) uint
}

// Init initializes the instance.
func (c NSCountedSet) Init() NSCountedSet {
	rv := objc.Send[NSCountedSet](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCountedSet) Autorelease() NSCountedSet {
	rv := objc.Send[NSCountedSet](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCountedSet creates a new NSCountedSet instance.
func NewNSCountedSet() NSCountedSet {
	class := getNSCountedSetClass()
	rv := objc.Send[NSCountedSet](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a counted set object initialized with the contents of a given
// array.
//
// array: An array of objects to add to the new set.
//
// # Return Value
//
// An initialized counted set object with the contents of `array`. The
// returned object might be different than the original receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSCountedSet/init(array:)
func NewCountedSetWithArray(array []objectivec.IObject) NSCountedSet {
	instance := getNSCountedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArray:"), objectivec.IObjectSliceToNSArray(array))
	return NSCountedSetFromID(rv)
}

// Returns a counted set object initialized with enough memory to hold a given
// number of objects.
//
// numItems: The initial capacity of the new counted set.
//
// # Return Value
//
// A counted set object initialized with enough memory to hold `numItems`
// objects
//
// # Discussion
//
// The method is the designated initializer for [NSCountedSet].
//
// Note that the capacity is simply a hint to help initial memory
// allocation—the initial count of the object is `0`, and the set still
// grows and shrinks as you add and remove objects. The hint is typically
// useful if the set will become large.
//
// See: https://developer.apple.com/documentation/Foundation/NSCountedSet/init(capacity:)
func NewCountedSetWithCapacity(numItems uint) NSCountedSet {
	instance := getNSCountedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCapacity:"), numItems)
	return NSCountedSetFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSMutableSet/init(coder:)
func NewCountedSetWithCoder(coder INSCoder) NSCountedSet {
	instance := getNSCountedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSCountedSetFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSSet/init(collectionViewIndexPath:)
func NewCountedSetWithCollectionViewIndexPath(indexPath objectivec.IObject) NSCountedSet {
	rv := objc.Send[objc.ID](objc.ID(getNSCountedSetClass().class), objc.Sel("setWithCollectionViewIndexPath:"), indexPath)
	return NSCountedSetFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSSet/init(collectionViewIndexPaths:)
func NewCountedSetWithCollectionViewIndexPaths(indexPaths []objc.ID) NSCountedSet {
	rv := objc.Send[objc.ID](objc.ID(getNSCountedSetClass().class), objc.Sel("setWithCollectionViewIndexPaths:"), objectivec.IDSliceToNSArray(indexPaths))
	return NSCountedSetFromID(rv)
}

// Creates and returns a set that contains a single given object.
//
// object: The object to add to the new set. `object` receives a [retain] message
// after being added to the set.
//
// # Return Value
//
// A new set that contains a single member, `object`.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/init(object:)
//
// [retain]: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/retain
func NewCountedSetWithObject(object objectivec.IObject) NSCountedSet {
	rv := objc.Send[objc.ID](objc.ID(getNSCountedSetClass().class), objc.Sel("setWithObject:"), object)
	return NSCountedSetFromID(rv)
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
// See: https://developer.apple.com/documentation/Foundation/NSSet/initWithObjects:
//
// [retain]: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/retain
func NewCountedSetWithObjects(firstObj objectivec.IObject) NSCountedSet {
	instance := getNSCountedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObjects:"), firstObj)
	return NSCountedSetFromID(rv)
}

// Returns a counted set object initialized with the contents of a given set.
//
// set: An set of objects to add to the new set.
//
// # Return Value
//
// An initialized counted set object with the contents of `set`. The returned
// object might be different than the original receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSCountedSet/init(set:)
func NewCountedSetWithSet(set INSSet) NSCountedSet {
	instance := getNSCountedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSet:"), set)
	return NSCountedSetFromID(rv)
}

// Initializes a newly allocated set and adds to it members of another given
// set.
//
// set: A set containing objects to add to the new set.
//
// flag: If true, each object in `set` receives a [copyWithZone:] message to create
// a copy of the object—objects must conform to the [NSCopying] protocol. In
// a managed memory environment, this is instead of the `retain` message the
// object would otherwise receive. The object copy is then added to the
// returned set.
//
// If false, then in a managed memory environment each object in `set` simply
// receives a `retain` message when it is added to the returned set.
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
// of arbitrary depth, passing true for the `flag` parameter will perform an
// immutable copy of the first level below the surface. If you pass false the
// mutability of the first level is unaffected. In either case, the mutability
// of all deeper levels is unaffected.
//
// See: https://developer.apple.com/documentation/Foundation/NSSet/init(set:copyItems:)
//
// [copyWithZone:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/copyWithZone:
func NewCountedSetWithSetCopyItems(set INSSet, flag bool) NSCountedSet {
	instance := getNSCountedSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSet:copyItems:"), set, flag)
	return NSCountedSetFromID(rv)
}

// Returns the count associated with a given object in the set.
//
// object: The object for which to return the count.
//
// # Return Value
//
// The count associated with `object` in the set, which can be thought of as
// the number of occurrences of `object` present in the set.
//
// See: https://developer.apple.com/documentation/Foundation/NSCountedSet/count(for:)
func (c NSCountedSet) CountForObject(object objectivec.IObject) uint {
	rv := objc.Send[uint](c.ID, objc.Sel("countForObject:"), object)
	return rv
}
