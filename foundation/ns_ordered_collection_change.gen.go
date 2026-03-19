// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSOrderedCollectionChange] class.
var (
	_NSOrderedCollectionChangeClass     NSOrderedCollectionChangeClass
	_NSOrderedCollectionChangeClassOnce sync.Once
)

func getNSOrderedCollectionChangeClass() NSOrderedCollectionChangeClass {
	_NSOrderedCollectionChangeClassOnce.Do(func() {
		_NSOrderedCollectionChangeClass = NSOrderedCollectionChangeClass{class: objc.GetClass("NSOrderedCollectionChange")}
	})
	return _NSOrderedCollectionChangeClass
}

// GetNSOrderedCollectionChangeClass returns the class object for NSOrderedCollectionChange.
func GetNSOrderedCollectionChangeClass() NSOrderedCollectionChangeClass {
	return getNSOrderedCollectionChangeClass()
}

type NSOrderedCollectionChangeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSOrderedCollectionChangeClass) Alloc() NSOrderedCollectionChange {
	rv := objc.Send[NSOrderedCollectionChange](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents an indexed change within an ordered collection.
//
// # Overview
// 
// An ordered collection change represents changes by adding, removing, or
// moving objects within an ordered collection. Changes with an associated
// index indicate a move within the collection.
//
// # Creating a Change
//
//   - [NSOrderedCollectionChange.InitWithObjectTypeIndex]: Creates a change object that represents inserting or removing an object from an ordered collection at a specific index.
//   - [NSOrderedCollectionChange.InitWithObjectTypeIndexAssociatedIndex]: Creates a change object that represents inserting, removing, or moving an object from an ordered collection at a specific index.
//
// # Accessing the Change
//
//   - [NSOrderedCollectionChange.ChangeType]: The type of change.
//   - [NSOrderedCollectionChange.Index]: The index location of the change.
//   - [NSOrderedCollectionChange.GetObject]: An object the change inserts or removes.
//   - [NSOrderedCollectionChange.AssociatedIndex]: When this property is set to a value other than [NSNotFound](<doc://com.apple.foundation/documentation/Foundation/NSNotFound-9t5v2>), the receiver is one half of a move, and this value is the index of the change’s counterpart of the opposite type in the diff.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange
type NSOrderedCollectionChange struct {
	objectivec.Object
}

// NSOrderedCollectionChangeFromID constructs a [NSOrderedCollectionChange] from an objc.ID.
//
// An object that represents an indexed change within an ordered collection.
func NSOrderedCollectionChangeFromID(id objc.ID) NSOrderedCollectionChange {
	return NSOrderedCollectionChange{objectivec.Object{ID: id}}
}
// NOTE: NSOrderedCollectionChange adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSOrderedCollectionChange] class.
//
// # Creating a Change
//
//   - [INSOrderedCollectionChange.InitWithObjectTypeIndex]: Creates a change object that represents inserting or removing an object from an ordered collection at a specific index.
//   - [INSOrderedCollectionChange.InitWithObjectTypeIndexAssociatedIndex]: Creates a change object that represents inserting, removing, or moving an object from an ordered collection at a specific index.
//
// # Accessing the Change
//
//   - [INSOrderedCollectionChange.ChangeType]: The type of change.
//   - [INSOrderedCollectionChange.Index]: The index location of the change.
//   - [INSOrderedCollectionChange.GetObject]: An object the change inserts or removes.
//   - [INSOrderedCollectionChange.AssociatedIndex]: When this property is set to a value other than [NSNotFound](<doc://com.apple.foundation/documentation/Foundation/NSNotFound-9t5v2>), the receiver is one half of a move, and this value is the index of the change’s counterpart of the opposite type in the diff.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange
type INSOrderedCollectionChange interface {
	objectivec.IObject

	// Topic: Creating a Change

	// Creates a change object that represents inserting or removing an object from an ordered collection at a specific index.
	InitWithObjectTypeIndex(anObject objectivec.IObject, type_ NSCollectionChangeType, index uint) NSOrderedCollectionChange
	// Creates a change object that represents inserting, removing, or moving an object from an ordered collection at a specific index.
	InitWithObjectTypeIndexAssociatedIndex(anObject objectivec.IObject, type_ NSCollectionChangeType, index uint, associatedIndex uint) NSOrderedCollectionChange

	// Topic: Accessing the Change

	// The type of change.
	ChangeType() NSCollectionChangeType
	// The index location of the change.
	Index() uint
	// An object the change inserts or removes.
	GetObject() objectivec.IObject
	// When this property is set to a value other than [NSNotFound](<doc://com.apple.foundation/documentation/Foundation/NSNotFound-9t5v2>), the receiver is one half of a move, and this value is the index of the change’s counterpart of the opposite type in the diff.
	AssociatedIndex() uint

	NSNotFound() int
	// A Boolean value that indicates if the difference has changes.
	HasChanges() bool
	SetHasChanges(value bool)
	// A collection of insertion change objects.
	Insertions() INSOrderedCollectionChange
	SetInsertions(value INSOrderedCollectionChange)
	// A collection of removal change objects.
	Removals() INSOrderedCollectionChange
	SetRemovals(value INSOrderedCollectionChange)
}

// Init initializes the instance.
func (o NSOrderedCollectionChange) Init() NSOrderedCollectionChange {
	rv := objc.Send[NSOrderedCollectionChange](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o NSOrderedCollectionChange) Autorelease() NSOrderedCollectionChange {
	rv := objc.Send[NSOrderedCollectionChange](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSOrderedCollectionChange creates a new NSOrderedCollectionChange instance.
func NewNSOrderedCollectionChange() NSOrderedCollectionChange {
	class := getNSOrderedCollectionChangeClass()
	rv := objc.Send[NSOrderedCollectionChange](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a change object that represents inserting or removing an object
// from an ordered collection at a specific index.
//
// anObject: An optional object the change will remove or insert.
//
// type: The type of change.
//
// index: The index location within an ordered collection where the change applies.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange/init(object:type:index:)
func NewOrderedCollectionChangeWithObjectTypeIndex(anObject objectivec.IObject, type_ NSCollectionChangeType, index uint) NSOrderedCollectionChange {
	instance := getNSOrderedCollectionChangeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObject:type:index:"), anObject, type_, index)
	return NSOrderedCollectionChangeFromID(rv)
}

// Creates a change object that represents inserting, removing, or moving an
// object from an ordered collection at a specific index.
//
// anObject: An optional object the change will remove or insert.
//
// type: The type of change.
//
// index: The index location within an ordered collection where the change applies.
//
// associatedIndex: The index of the change’s counterpart of the opposite type in the diff.
//
// # Discussion
// 
// Pairs of changes with opposite types that refer to each other represent the
// index location of their counterpart with the [AssociatedIndex] property.
// Initializing an [NSOrderedCollectionDifference] with broken associations
// (or associations that aren’t reflexive) generates an exception. The
// following example creates a diff where the object `@”Red”` moves from
// index `8` to index `3`:
// 
// A move pair can have a different `object` in its removal and insertion
// changes, which can imply that the change represents moving and changing or
// replacing an element. Diffs that [controller(_:didChangeContentWith:)]
// passes to delegates of [NSFetchedResultsController] communicate that an
// object changed even when its position in the results is unaffected.
//
// [NSFetchedResultsController]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController
// [controller(_:didChangeContentWith:)]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsControllerDelegate/controller(_:didChangeContentWith:)-5ullb
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange/init(object:type:index:associatedIndex:)
func NewOrderedCollectionChangeWithObjectTypeIndexAssociatedIndex(anObject objectivec.IObject, type_ NSCollectionChangeType, index uint, associatedIndex uint) NSOrderedCollectionChange {
	instance := getNSOrderedCollectionChangeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObject:type:index:associatedIndex:"), anObject, type_, index, associatedIndex)
	return NSOrderedCollectionChangeFromID(rv)
}

// Creates a change object that represents inserting or removing an object
// from an ordered collection at a specific index.
//
// anObject: An optional object the change will remove or insert.
//
// type: The type of change.
//
// index: The index location within an ordered collection where the change applies.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange/init(object:type:index:)
func (o NSOrderedCollectionChange) InitWithObjectTypeIndex(anObject objectivec.IObject, type_ NSCollectionChangeType, index uint) NSOrderedCollectionChange {
	rv := objc.Send[NSOrderedCollectionChange](o.ID, objc.Sel("initWithObject:type:index:"), anObject, type_, index)
	return rv
}
// Creates a change object that represents inserting, removing, or moving an
// object from an ordered collection at a specific index.
//
// anObject: An optional object the change will remove or insert.
//
// type: The type of change.
//
// index: The index location within an ordered collection where the change applies.
//
// associatedIndex: The index of the change’s counterpart of the opposite type in the diff.
//
// # Discussion
// 
// Pairs of changes with opposite types that refer to each other represent the
// index location of their counterpart with the [AssociatedIndex] property.
// Initializing an [NSOrderedCollectionDifference] with broken associations
// (or associations that aren’t reflexive) generates an exception. The
// following example creates a diff where the object `@”Red”` moves from
// index `8` to index `3`:
// 
// A move pair can have a different `object` in its removal and insertion
// changes, which can imply that the change represents moving and changing or
// replacing an element. Diffs that [controller(_:didChangeContentWith:)]
// passes to delegates of [NSFetchedResultsController] communicate that an
// object changed even when its position in the results is unaffected.
//
// [NSFetchedResultsController]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController
// [controller(_:didChangeContentWith:)]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsControllerDelegate/controller(_:didChangeContentWith:)-5ullb
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange/init(object:type:index:associatedIndex:)
func (o NSOrderedCollectionChange) InitWithObjectTypeIndexAssociatedIndex(anObject objectivec.IObject, type_ NSCollectionChangeType, index uint, associatedIndex uint) NSOrderedCollectionChange {
	rv := objc.Send[NSOrderedCollectionChange](o.ID, objc.Sel("initWithObject:type:index:associatedIndex:"), anObject, type_, index, associatedIndex)
	return rv
}

// Creates an change object that represents inserting or removing an object
// from an ordered collection at a specific index.
//
// anObject: An object to be removed or inserted by the change.
//
// type: The type of change
//
// index: The index location within an ordered collection where the change applies.
//
// # Return Value
// 
// An object that represents an indexed change to an ordered collection and
// references the object to be inserted or removed.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange/changeWithObject:type:index:
func (_NSOrderedCollectionChangeClass NSOrderedCollectionChangeClass) ChangeWithObjectTypeIndex(anObject objectivec.IObject, type_ NSCollectionChangeType, index uint) NSOrderedCollectionChange {
	rv := objc.Send[objc.ID](objc.ID(_NSOrderedCollectionChangeClass.class), objc.Sel("changeWithObject:type:index:"), anObject, type_, index)
	return NSOrderedCollectionChangeFromID(rv)
}
// Creates an change object that represents inserting or removing an object
// from an ordered collection at a specific index, matched with an associated
// location that infers a move within the collection.
//
// anObject: An object to be removed or inserted by the change.
//
// type: The type of change
//
// index: The index location within an ordered collection where the change applies.
//
// associatedIndex: The index of the change’s counterpart of the opposite type in the diff.
//
// # Return Value
// 
// An object that represents an indexed change to an ordered collection and
// references the object to be inserted or removed with an associated index
// that infers a move within the collection.
//
// # Discussion
// 
// Pairs of changes with opposite types that refer to each other represent the
// index location of their counterpart with the [AssociatedIndex] property.
// Initializing a [NSOrderedCollectionDifference] with broken associations (or
// associations that aren’t reflexive) will generate an exception. The
// following example creates a diff where the object `@”Red”` moves from
// index `8` to index `3`:
// 
// A move pair can have a different `object` in its removal and insertion
// changes, which can imply that the change represents moving and changing or
// replacing an element. Diffs that [controller(_:didChangeContentWith:)]
// passes to delegates of [NSFetchedResultsController] communicate that an
// object changed even when its position in the results is unaffected.
//
// [NSFetchedResultsController]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController
// [controller(_:didChangeContentWith:)]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsControllerDelegate/controller(_:didChangeContentWith:)-5ullb
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange/changeWithObject:type:index:associatedIndex:
func (_NSOrderedCollectionChangeClass NSOrderedCollectionChangeClass) ChangeWithObjectTypeIndexAssociatedIndex(anObject objectivec.IObject, type_ NSCollectionChangeType, index uint, associatedIndex uint) NSOrderedCollectionChange {
	rv := objc.Send[objc.ID](objc.ID(_NSOrderedCollectionChangeClass.class), objc.Sel("changeWithObject:type:index:associatedIndex:"), anObject, type_, index, associatedIndex)
	return NSOrderedCollectionChangeFromID(rv)
}

// The type of change.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange/changeType
func (o NSOrderedCollectionChange) ChangeType() NSCollectionChangeType {
	rv := objc.Send[NSCollectionChangeType](o.ID, objc.Sel("changeType"))
	return NSCollectionChangeType(rv)
}
// The index location of the change.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange/index
func (o NSOrderedCollectionChange) Index() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("index"))
	return rv
}
// An object the change inserts or removes.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange/object
func (o NSOrderedCollectionChange) GetObject() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("object"))
	return objectivec.Object{ID: rv}
}
// When this property is set to a value other than [NSNotFound], the receiver
// is one half of a move, and this value is the index of the change’s
// counterpart of the opposite type in the diff.
//
// [NSNotFound]: https://developer.apple.com/documentation/Foundation/NSNotFound-9t5v2
//
// # Discussion
// 
// Pairs of changes with opposite types that refer to each other represent the
// index location of their counterpart with the [AssociatedIndex] property.
// The following example creates a diff where the object `@”Red”` moves
// from index `8` to index `3`:
// 
// A move pair can have a different `object` in its removal and insertion
// changes, which can imply that the change represents moving and changing or
// replacing an element. Diffs that [controller(_:didChangeContentWith:)]
// passes to delegates of [NSFetchedResultsController] communicate that an
// object changed even when its position in the results is unaffected.
//
// [NSFetchedResultsController]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsController
// [controller(_:didChangeContentWith:)]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsControllerDelegate/controller(_:didChangeContentWith:)-5ullb
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange/associatedIndex
func (o NSOrderedCollectionChange) AssociatedIndex() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("associatedIndex"))
	return rv
}
// See: https://developer.apple.com/documentation/foundation/nsnotfound-9t5v2
func (o NSOrderedCollectionChange) NSNotFound() int {
	rv := objc.Send[int](o.ID, objc.Sel("NSNotFound"))
	return rv
}
// A Boolean value that indicates if the difference has changes.
//
// See: https://developer.apple.com/documentation/foundation/nsorderedcollectiondifference/haschanges
func (o NSOrderedCollectionChange) HasChanges() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("hasChanges"))
	return rv
}
func (o NSOrderedCollectionChange) SetHasChanges(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setHasChanges:"), value)
}
// A collection of insertion change objects.
//
// See: https://developer.apple.com/documentation/foundation/nsorderedcollectiondifference/insertions
func (o NSOrderedCollectionChange) Insertions() INSOrderedCollectionChange {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("insertions"))
	return NSOrderedCollectionChangeFromID(objc.ID(rv))
}
func (o NSOrderedCollectionChange) SetInsertions(value INSOrderedCollectionChange) {
	objc.Send[struct{}](o.ID, objc.Sel("setInsertions:"), value)
}
// A collection of removal change objects.
//
// See: https://developer.apple.com/documentation/foundation/nsorderedcollectiondifference/removals
func (o NSOrderedCollectionChange) Removals() INSOrderedCollectionChange {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("removals"))
	return NSOrderedCollectionChangeFromID(objc.ID(rv))
}
func (o NSOrderedCollectionChange) SetRemovals(value INSOrderedCollectionChange) {
	objc.Send[struct{}](o.ID, objc.Sel("setRemovals:"), value)
}

