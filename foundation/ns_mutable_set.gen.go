// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMutableSet] class.
var (
	_NSMutableSetClass     NSMutableSetClass
	_NSMutableSetClassOnce sync.Once
)

func getNSMutableSetClass() NSMutableSetClass {
	_NSMutableSetClassOnce.Do(func() {
		_NSMutableSetClass = NSMutableSetClass{class: objc.GetClass("NSMutableSet")}
	})
	return _NSMutableSetClass
}

// GetNSMutableSetClass returns the class object for NSMutableSet.
func GetNSMutableSetClass() NSMutableSetClass {
	return getNSMutableSetClass()
}

type NSMutableSetClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMutableSetClass) Alloc() NSMutableSet {
	rv := objc.Send[NSMutableSet](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A dynamic unordered collection of unique objects.
//
// # Overview
// 
// You can use this type in Swift instead of a [Set] in cases that require
// reference semantics.
// 
// The [NSMutableSet] class declares the programmatic interface to a mutable,
// unordered collection of distinct objects.
// 
// The [NSCountedSet] class, which is a concrete subclass of [NSMutableSet],
// supports mutable sets that can contain multiple instances of the same
// element. The [NSSet] class supports creating and managing immutable sets.
// 
// NSMutableSet is “toll-free bridged” with its Core Foundation
// counterpart, [CFMutableSet]. See [Toll-Free Bridging] for more information.
// 
// # Subclassing Notes
// 
// There should be little need of subclassing. If you need to customize
// behavior, it is often better to consider composition instead of
// subclassing.
// 
// # Methods to Override
// 
// In a subclass, you must override both of its primitive methods:
// 
// - [NSMutableSet.AddObject]
// - [NSMutableSet.RemoveObject]
// 
// You must also override the primitive methods of the [NSSet] class.
//
// [CFMutableSet]: https://developer.apple.com/documentation/CoreFoundation/CFMutableSet
// [Set]: https://developer.apple.com/documentation/Swift/Set
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
//
// # Creating a mutable set
//
//   - [NSMutableSet.InitWithCapacity]: Returns an initialized mutable set with a given initial capacity.
//
// # Adding and removing entries
//
//   - [NSMutableSet.AddObject]: Adds a given object to the set, if it is not already a member.
//   - [NSMutableSet.FilterUsingPredicate]: Evaluates a given predicate against the set’s content and removes from the set those objects for which the predicate returns false.
//   - [NSMutableSet.RemoveObject]: Removes a given object from the set.
//   - [NSMutableSet.RemoveAllObjects]: Empties the set of all of its members.
//   - [NSMutableSet.AddObjectsFromArray]: Adds to the set each object contained in a given array that is not already a member.
//
// # Combining and recombining sets
//
//   - [NSMutableSet.UnionSet]: Adds each object in another given set to the receiving set, if not present.
//   - [NSMutableSet.MinusSet]: Removes each object in another given set from the receiving set, if present.
//   - [NSMutableSet.IntersectSet]: Removes from the receiving set each object that isn’t a member of another given set.
//   - [NSMutableSet.SetSet]: Empties the receiving set, then adds each object contained in another given set.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableSet
type NSMutableSet struct {
	NSSet
}

// NSMutableSetFromID constructs a [NSMutableSet] from an objc.ID.
//
// A dynamic unordered collection of unique objects.
func NSMutableSetFromID(id objc.ID) NSMutableSet {
	return NSMutableSet{NSSet: NSSetFromID(id)}
}
// NOTE: NSMutableSet adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMutableSet] class.
//
// # Creating a mutable set
//
//   - [INSMutableSet.InitWithCapacity]: Returns an initialized mutable set with a given initial capacity.
//
// # Adding and removing entries
//
//   - [INSMutableSet.AddObject]: Adds a given object to the set, if it is not already a member.
//   - [INSMutableSet.FilterUsingPredicate]: Evaluates a given predicate against the set’s content and removes from the set those objects for which the predicate returns false.
//   - [INSMutableSet.RemoveObject]: Removes a given object from the set.
//   - [INSMutableSet.RemoveAllObjects]: Empties the set of all of its members.
//   - [INSMutableSet.AddObjectsFromArray]: Adds to the set each object contained in a given array that is not already a member.
//
// # Combining and recombining sets
//
//   - [INSMutableSet.UnionSet]: Adds each object in another given set to the receiving set, if not present.
//   - [INSMutableSet.MinusSet]: Removes each object in another given set from the receiving set, if present.
//   - [INSMutableSet.IntersectSet]: Removes from the receiving set each object that isn’t a member of another given set.
//   - [INSMutableSet.SetSet]: Empties the receiving set, then adds each object contained in another given set.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableSet
type INSMutableSet interface {
	INSSet

	// Topic: Creating a mutable set

	// Returns an initialized mutable set with a given initial capacity.
	InitWithCapacity(numItems uint) NSMutableSet

	// Topic: Adding and removing entries

	// Adds a given object to the set, if it is not already a member.
	AddObject(object objectivec.IObject)
	// Evaluates a given predicate against the set’s content and removes from the set those objects for which the predicate returns false.
	FilterUsingPredicate(predicate INSPredicate)
	// Removes a given object from the set.
	RemoveObject(object objectivec.IObject)
	// Empties the set of all of its members.
	RemoveAllObjects()
	// Adds to the set each object contained in a given array that is not already a member.
	AddObjectsFromArray(array []objectivec.IObject)

	// Topic: Combining and recombining sets

	// Adds each object in another given set to the receiving set, if not present.
	UnionSet(otherSet INSSet)
	// Removes each object in another given set from the receiving set, if present.
	MinusSet(otherSet INSSet)
	// Removes from the receiving set each object that isn’t a member of another given set.
	IntersectSet(otherSet INSSet)
	// Empties the receiving set, then adds each object contained in another given set.
	SetSet(otherSet INSSet)
}

// Init initializes the instance.
func (m NSMutableSet) Init() NSMutableSet {
	rv := objc.Send[NSMutableSet](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMutableSet) Autorelease() NSMutableSet {
	rv := objc.Send[NSMutableSet](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMutableSet creates a new NSMutableSet instance.
func NewNSMutableSet() NSMutableSet {
	class := getNSMutableSetClass()
	rv := objc.Send[NSMutableSet](objc.ID(class.class), objc.Sel("new"))
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
func NewMutableSetWithArray(array []objectivec.IObject) NSMutableSet {
	instance := getNSMutableSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArray:"), objectivec.IObjectSliceToNSArray(array))
	return NSMutableSetFromID(rv)
}

// Returns an initialized mutable set with a given initial capacity.
//
// numItems: The initial capacity of the set.
//
// # Return Value
// 
// An initialized mutable set with initial capacity to hold `numItems`
// members. The returned set might be different than the original receiver.
//
// # Discussion
// 
// Mutable sets allocate additional memory as needed, so `numItems` simply
// establishes the object’s initial capacity.
// 
// This method is a designated initializer for [NSMutableSet].
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableSet/init(capacity:)
func NewMutableSetWithCapacity(numItems uint) NSMutableSet {
	instance := getNSMutableSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCapacity:"), numItems)
	return NSMutableSetFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSMutableSet/init(coder:)
func NewMutableSetWithCoder(coder INSCoder) NSMutableSet {
	instance := getNSMutableSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSMutableSetFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSSet/init(collectionViewIndexPath:)
func NewMutableSetWithCollectionViewIndexPath(indexPath objectivec.IObject) NSMutableSet {
	rv := objc.Send[objc.ID](objc.ID(getNSMutableSetClass().class), objc.Sel("setWithCollectionViewIndexPath:"), indexPath)
	return NSMutableSetFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSSet/init(collectionViewIndexPaths:)
func NewMutableSetWithCollectionViewIndexPaths(indexPaths []objc.ID) NSMutableSet {
	rv := objc.Send[objc.ID](objc.ID(getNSMutableSetClass().class), objc.Sel("setWithCollectionViewIndexPaths:"), objectivec.IDSliceToNSArray(indexPaths))
	return NSMutableSetFromID(rv)
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
func NewMutableSetWithObject(object objectivec.IObject) NSMutableSet {
	rv := objc.Send[objc.ID](objc.ID(getNSMutableSetClass().class), objc.Sel("setWithObject:"), object)
	return NSMutableSetFromID(rv)
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
func NewMutableSetWithObjects(firstObj objectivec.IObject) NSMutableSet {
	instance := getNSMutableSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObjects:"), firstObj)
	return NSMutableSetFromID(rv)
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
func NewMutableSetWithSet(set INSSet) NSMutableSet {
	instance := getNSMutableSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSet:"), set)
	return NSMutableSetFromID(rv)
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
func NewMutableSetWithSetCopyItems(set INSSet, flag bool) NSMutableSet {
	instance := getNSMutableSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSet:copyItems:"), set, flag)
	return NSMutableSetFromID(rv)
}

// Returns an initialized mutable set with a given initial capacity.
//
// numItems: The initial capacity of the set.
//
// # Return Value
// 
// An initialized mutable set with initial capacity to hold `numItems`
// members. The returned set might be different than the original receiver.
//
// # Discussion
// 
// Mutable sets allocate additional memory as needed, so `numItems` simply
// establishes the object’s initial capacity.
// 
// This method is a designated initializer for [NSMutableSet].
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableSet/init(capacity:)
func (m NSMutableSet) InitWithCapacity(numItems uint) NSMutableSet {
	rv := objc.Send[NSMutableSet](m.ID, objc.Sel("initWithCapacity:"), numItems)
	return rv
}

// Adds a given object to the set, if it is not already a member.
//
// object: The object to add to the set.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableSet/add(_:)
func (m NSMutableSet) AddObject(object objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("addObject:"), object)
}

// Evaluates a given predicate against the set’s content and removes from
// the set those objects for which the predicate returns false.
//
// predicate: A predicate.
//
// # Discussion
// 
// The following example illustrates the use of this method.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableSet/filter(using:)
func (m NSMutableSet) FilterUsingPredicate(predicate INSPredicate) {
	objc.Send[objc.ID](m.ID, objc.Sel("filterUsingPredicate:"), predicate)
}

// Removes a given object from the set.
//
// object: The object to remove from the set.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableSet/remove(_:)
func (m NSMutableSet) RemoveObject(object objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeObject:"), object)
}

// Empties the set of all of its members.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableSet/removeAllObjects()
func (m NSMutableSet) RemoveAllObjects() {
	objc.Send[objc.ID](m.ID, objc.Sel("removeAllObjects"))
}

// Adds to the set each object contained in a given array that is not already
// a member.
//
// array: An array of objects to add to the set.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableSet/addObjects(from:)
func (m NSMutableSet) AddObjectsFromArray(array []objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("addObjectsFromArray:"), objectivec.IObjectSliceToNSArray(array))
}

// Adds each object in another given set to the receiving set, if not present.
//
// otherSet: The set of objects to add to the receiving set.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableSet/union(_:)
func (m NSMutableSet) UnionSet(otherSet INSSet) {
	objc.Send[objc.ID](m.ID, objc.Sel("unionSet:"), otherSet)
}

// Removes each object in another given set from the receiving set, if
// present.
//
// otherSet: The set of objects to remove from the receiving set.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableSet/minus(_:)
func (m NSMutableSet) MinusSet(otherSet INSSet) {
	objc.Send[objc.ID](m.ID, objc.Sel("minusSet:"), otherSet)
}

// Removes from the receiving set each object that isn’t a member of another
// given set.
//
// otherSet: The set with which to perform the intersection.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableSet/intersect(_:)
func (m NSMutableSet) IntersectSet(otherSet INSSet) {
	objc.Send[objc.ID](m.ID, objc.Sel("intersectSet:"), otherSet)
}

// Empties the receiving set, then adds each object contained in another given
// set.
//
// otherSet: The set whose members replace the receiving set’s content.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableSet/setSet(_:)
func (m NSMutableSet) SetSet(otherSet INSSet) {
	objc.Send[objc.ID](m.ID, objc.Sel("setSet:"), otherSet)
}

// Creates and returns a mutable set with a given initial capacity.
//
// numItems: The initial capacity of the new set.
//
// # Return Value
// 
// A mutable set with initial capacity to hold `numItems` members.
//
// # Discussion
// 
// Mutable sets allocate additional memory as needed, so `numItems` simply
// establishes the object’s initial capacity.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableSet/setWithCapacity:
func (_NSMutableSetClass NSMutableSetClass) SetWithCapacity(numItems uint) NSMutableSet {
	rv := objc.Send[objc.ID](objc.ID(_NSMutableSetClass.class), objc.Sel("setWithCapacity:"), numItems)
	return NSMutableSetFromID(rv)
}

