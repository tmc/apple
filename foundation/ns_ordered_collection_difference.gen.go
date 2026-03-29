// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSOrderedCollectionDifference] class.
var (
	_NSOrderedCollectionDifferenceClass     NSOrderedCollectionDifferenceClass
	_NSOrderedCollectionDifferenceClassOnce sync.Once
)

func getNSOrderedCollectionDifferenceClass() NSOrderedCollectionDifferenceClass {
	_NSOrderedCollectionDifferenceClassOnce.Do(func() {
		_NSOrderedCollectionDifferenceClass = NSOrderedCollectionDifferenceClass{class: objc.GetClass("NSOrderedCollectionDifference")}
	})
	return _NSOrderedCollectionDifferenceClass
}

// GetNSOrderedCollectionDifferenceClass returns the class object for NSOrderedCollectionDifference.
func GetNSOrderedCollectionDifferenceClass() NSOrderedCollectionDifferenceClass {
	return getNSOrderedCollectionDifferenceClass()
}

type NSOrderedCollectionDifferenceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSOrderedCollectionDifferenceClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSOrderedCollectionDifferenceClass) Alloc() NSOrderedCollectionDifference {
	rv := objc.Send[NSOrderedCollectionDifference](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object representing the difference between two ordered collections.
//
// # Overview
// 
// Use [DifferenceFromArray] or one of its variations to get an instance of
// [NSOrderedCollectionDifference], which represents the difference between
// two ordered collections.
// 
// For example, the following sample compares two arrays of strings to create
// a difference that represents the changes:
//
// # Accessing Changes
//
//   - [NSOrderedCollectionDifference.HasChanges]: A Boolean value that indicates if the difference has changes.
//   - [NSOrderedCollectionDifference.Insertions]: A collection of insertion change objects.
//   - [NSOrderedCollectionDifference.Removals]: A collection of removal change objects.
//
// # Inverting a Difference Object
//
//   - [NSOrderedCollectionDifference.InverseDifference]: Calculate the difference between two objects in the reverse direction of comparison.
//
// # Creating a Collection Difference Object
//
//   - [NSOrderedCollectionDifference.InitWithChanges]: Creates an ordered collection difference using an array of ordered collection changes.
//   - [NSOrderedCollectionDifference.InitWithInsertIndexesInsertedObjectsRemoveIndexesRemovedObjects]: Creates an ordered collection difference from arrays of inserted and removed objects with corresponding sets of indices.
//   - [NSOrderedCollectionDifference.InitWithInsertIndexesInsertedObjectsRemoveIndexesRemovedObjectsAdditionalChanges]: Creates an ordered collection difference from arrays of inserted and removed objects with corresponding sets of indices, in addition to an array of ordered collection changes.
//
// # Updating Changes from a Difference Object
//
//   - [NSOrderedCollectionDifference.DifferenceByTransformingChangesWithBlock]: Create a new ordered collection difference by mapping over this difference’s members, processing the change objects with the block provided.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference
type NSOrderedCollectionDifference struct {
	objectivec.Object
}

// NSOrderedCollectionDifferenceFromID constructs a [NSOrderedCollectionDifference] from an objc.ID.
//
// An object representing the difference between two ordered collections.
func NSOrderedCollectionDifferenceFromID(id objc.ID) NSOrderedCollectionDifference {
	return NSOrderedCollectionDifference{objectivec.Object{ID: id}}
}
// NOTE: NSOrderedCollectionDifference adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSOrderedCollectionDifference] class.
//
// # Accessing Changes
//
//   - [INSOrderedCollectionDifference.HasChanges]: A Boolean value that indicates if the difference has changes.
//   - [INSOrderedCollectionDifference.Insertions]: A collection of insertion change objects.
//   - [INSOrderedCollectionDifference.Removals]: A collection of removal change objects.
//
// # Inverting a Difference Object
//
//   - [INSOrderedCollectionDifference.InverseDifference]: Calculate the difference between two objects in the reverse direction of comparison.
//
// # Creating a Collection Difference Object
//
//   - [INSOrderedCollectionDifference.InitWithChanges]: Creates an ordered collection difference using an array of ordered collection changes.
//   - [INSOrderedCollectionDifference.InitWithInsertIndexesInsertedObjectsRemoveIndexesRemovedObjects]: Creates an ordered collection difference from arrays of inserted and removed objects with corresponding sets of indices.
//   - [INSOrderedCollectionDifference.InitWithInsertIndexesInsertedObjectsRemoveIndexesRemovedObjectsAdditionalChanges]: Creates an ordered collection difference from arrays of inserted and removed objects with corresponding sets of indices, in addition to an array of ordered collection changes.
//
// # Updating Changes from a Difference Object
//
//   - [INSOrderedCollectionDifference.DifferenceByTransformingChangesWithBlock]: Create a new ordered collection difference by mapping over this difference’s members, processing the change objects with the block provided.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference
type INSOrderedCollectionDifference interface {
	objectivec.IObject

	// Topic: Accessing Changes

	// A Boolean value that indicates if the difference has changes.
	HasChanges() bool
	// A collection of insertion change objects.
	Insertions() []NSOrderedCollectionChange
	// A collection of removal change objects.
	Removals() []NSOrderedCollectionChange

	// Topic: Inverting a Difference Object

	// Calculate the difference between two objects in the reverse direction of comparison.
	InverseDifference() INSOrderedCollectionDifference

	// Topic: Creating a Collection Difference Object

	// Creates an ordered collection difference using an array of ordered collection changes.
	InitWithChanges(changes []NSOrderedCollectionChange) NSOrderedCollectionDifference
	// Creates an ordered collection difference from arrays of inserted and removed objects with corresponding sets of indices.
	InitWithInsertIndexesInsertedObjectsRemoveIndexesRemovedObjects(inserts INSIndexSet, insertedObjects []objectivec.IObject, removes INSIndexSet, removedObjects []objectivec.IObject) NSOrderedCollectionDifference
	// Creates an ordered collection difference from arrays of inserted and removed objects with corresponding sets of indices, in addition to an array of ordered collection changes.
	InitWithInsertIndexesInsertedObjectsRemoveIndexesRemovedObjectsAdditionalChanges(inserts INSIndexSet, insertedObjects []objectivec.IObject, removes INSIndexSet, removedObjects []objectivec.IObject, changes []NSOrderedCollectionChange) NSOrderedCollectionDifference

	// Topic: Updating Changes from a Difference Object

	// Create a new ordered collection difference by mapping over this difference’s members, processing the change objects with the block provided.
	DifferenceByTransformingChangesWithBlock(block OrderedCollectionChangeHandler) INSOrderedCollectionDifference
}

// Init initializes the instance.
func (o NSOrderedCollectionDifference) Init() NSOrderedCollectionDifference {
	rv := objc.Send[NSOrderedCollectionDifference](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o NSOrderedCollectionDifference) Autorelease() NSOrderedCollectionDifference {
	rv := objc.Send[NSOrderedCollectionDifference](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSOrderedCollectionDifference creates a new NSOrderedCollectionDifference instance.
func NewNSOrderedCollectionDifference() NSOrderedCollectionDifference {
	class := getNSOrderedCollectionDifferenceClass()
	rv := objc.Send[NSOrderedCollectionDifference](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an ordered collection difference using an array of ordered
// collection changes.
//
// changes: An array of ordered collection changes.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/init(changes:)
func NewOrderedCollectionDifferenceWithChanges(changes []NSOrderedCollectionChange) NSOrderedCollectionDifference {
	instance := getNSOrderedCollectionDifferenceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithChanges:"), objectivec.IObjectSliceToNSArray(changes))
	return NSOrderedCollectionDifferenceFromID(rv)
}

// Creates an ordered collection difference from arrays of inserted and
// removed objects with corresponding sets of indices.
//
// inserts: An index set that represents the index values to associate with the objects
// in the provided array of inserted objects.
//
// insertedObjects: An array of objects the ordered collection difference will insert.
//
// removes: An index set that represents the index values to associate with the objects
// in the provided array of removed objects.
//
// removedObjects: An array of objects the ordered collection difference will remove.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/init(insert:insertedObjects:remove:removedObjects:)
func NewOrderedCollectionDifferenceWithInsertIndexesInsertedObjectsRemoveIndexesRemovedObjects(inserts INSIndexSet, insertedObjects []objectivec.IObject, removes INSIndexSet, removedObjects []objectivec.IObject) NSOrderedCollectionDifference {
	instance := getNSOrderedCollectionDifferenceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInsertIndexes:insertedObjects:removeIndexes:removedObjects:"), inserts, objectivec.IObjectSliceToNSArray(insertedObjects), removes, objectivec.IObjectSliceToNSArray(removedObjects))
	return NSOrderedCollectionDifferenceFromID(rv)
}

// Creates an ordered collection difference from arrays of inserted and
// removed objects with corresponding sets of indices, in addition to an array
// of ordered collection changes.
//
// inserts: An index set that represents the index values to associate with the objects
// in the provided array of inserted objects.
//
// insertedObjects: An array of objects the ordered collection difference will insert.
//
// removes: An index set that represents the index values to associate with the objects
// in the provided array of removed objects.
//
// removedObjects: An array of objects the ordered collection difference will remove.
//
// changes: An array of ordered collection changes.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/init(insert:insertedObjects:remove:removedObjects:additionalChanges:)
func NewOrderedCollectionDifferenceWithInsertIndexesInsertedObjectsRemoveIndexesRemovedObjectsAdditionalChanges(inserts INSIndexSet, insertedObjects []objectivec.IObject, removes INSIndexSet, removedObjects []objectivec.IObject, changes []NSOrderedCollectionChange) NSOrderedCollectionDifference {
	instance := getNSOrderedCollectionDifferenceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInsertIndexes:insertedObjects:removeIndexes:removedObjects:additionalChanges:"), inserts, objectivec.IObjectSliceToNSArray(insertedObjects), removes, objectivec.IObjectSliceToNSArray(removedObjects), objectivec.IObjectSliceToNSArray(changes))
	return NSOrderedCollectionDifferenceFromID(rv)
}

// Calculate the difference between two objects in the reverse direction of
// comparison.
//
// # Return Value
// 
// A copy of the receiver with all removals changed to insertions (and vice
// versa).
//
// # Discussion
// 
// Applying a difference to an ordered collection and then applying the
// inverse difference results in the original ordered collection:
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/inverse()
func (o NSOrderedCollectionDifference) InverseDifference() INSOrderedCollectionDifference {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inverseDifference"))
	return NSOrderedCollectionDifferenceFromID(rv)
}
// Creates an ordered collection difference using an array of ordered
// collection changes.
//
// changes: An array of ordered collection changes.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/init(changes:)
func (o NSOrderedCollectionDifference) InitWithChanges(changes []NSOrderedCollectionChange) NSOrderedCollectionDifference {
	rv := objc.Send[NSOrderedCollectionDifference](o.ID, objc.Sel("initWithChanges:"), objectivec.IObjectSliceToNSArray(changes))
	return rv
}
// Creates an ordered collection difference from arrays of inserted and
// removed objects with corresponding sets of indices.
//
// inserts: An index set that represents the index values to associate with the objects
// in the provided array of inserted objects.
//
// insertedObjects: An array of objects the ordered collection difference will insert.
//
// removes: An index set that represents the index values to associate with the objects
// in the provided array of removed objects.
//
// removedObjects: An array of objects the ordered collection difference will remove.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/init(insert:insertedObjects:remove:removedObjects:)
func (o NSOrderedCollectionDifference) InitWithInsertIndexesInsertedObjectsRemoveIndexesRemovedObjects(inserts INSIndexSet, insertedObjects []objectivec.IObject, removes INSIndexSet, removedObjects []objectivec.IObject) NSOrderedCollectionDifference {
	rv := objc.Send[NSOrderedCollectionDifference](o.ID, objc.Sel("initWithInsertIndexes:insertedObjects:removeIndexes:removedObjects:"), inserts, objectivec.IObjectSliceToNSArray(insertedObjects), removes, objectivec.IObjectSliceToNSArray(removedObjects))
	return rv
}
// Creates an ordered collection difference from arrays of inserted and
// removed objects with corresponding sets of indices, in addition to an array
// of ordered collection changes.
//
// inserts: An index set that represents the index values to associate with the objects
// in the provided array of inserted objects.
//
// insertedObjects: An array of objects the ordered collection difference will insert.
//
// removes: An index set that represents the index values to associate with the objects
// in the provided array of removed objects.
//
// removedObjects: An array of objects the ordered collection difference will remove.
//
// changes: An array of ordered collection changes.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/init(insert:insertedObjects:remove:removedObjects:additionalChanges:)
func (o NSOrderedCollectionDifference) InitWithInsertIndexesInsertedObjectsRemoveIndexesRemovedObjectsAdditionalChanges(inserts INSIndexSet, insertedObjects []objectivec.IObject, removes INSIndexSet, removedObjects []objectivec.IObject, changes []NSOrderedCollectionChange) NSOrderedCollectionDifference {
	rv := objc.Send[NSOrderedCollectionDifference](o.ID, objc.Sel("initWithInsertIndexes:insertedObjects:removeIndexes:removedObjects:additionalChanges:"), inserts, objectivec.IObjectSliceToNSArray(insertedObjects), removes, objectivec.IObjectSliceToNSArray(removedObjects), objectivec.IObjectSliceToNSArray(changes))
	return rv
}
// Create a new ordered collection difference by mapping over this
// difference’s members, processing the change objects with the block
// provided.
//
// block: A block receives an ordered collection change and returns an updated
// change.
//
// # Return Value
// 
// A new ordered collection difference.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/transformingChanges(_:)
func (o NSOrderedCollectionDifference) DifferenceByTransformingChangesWithBlock(block OrderedCollectionChangeHandler) INSOrderedCollectionDifference {
_block0, _ := NewOrderedCollectionChangeBlock(block)
	rv := objc.Send[objc.ID](o.ID, objc.Sel("differenceByTransformingChangesWithBlock:"), _block0)
	return NSOrderedCollectionDifferenceFromID(rv)
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
func (o NSOrderedCollectionDifference) CountByEnumeratingWithStateObjectsCount(state NSFastEnumerationState, buffer []objectivec.IObject, len_ uint) uint {
	rv := objc.Send[uint](o.ID, objc.Sel("countByEnumeratingWithState:objects:count:"), state, objc.CArray(buffer), len_)
	return rv
}

// A Boolean value that indicates if the difference has changes.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/hasChanges
func (o NSOrderedCollectionDifference) HasChanges() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("hasChanges"))
	return rv
}
// A collection of insertion change objects.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/insertions
func (o NSOrderedCollectionDifference) Insertions() []NSOrderedCollectionChange {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("insertions"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSOrderedCollectionChange {
		return NSOrderedCollectionChangeFromID(id)
	})
}
// A collection of removal change objects.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/removals
func (o NSOrderedCollectionDifference) Removals() []NSOrderedCollectionChange {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("removals"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSOrderedCollectionChange {
		return NSOrderedCollectionChangeFromID(id)
	})
}

// DifferenceByTransformingChangesWithBlockSync is a synchronous wrapper around [NSOrderedCollectionDifference.DifferenceByTransformingChangesWithBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (o NSOrderedCollectionDifference) DifferenceByTransformingChangesWithBlockSync(ctx context.Context) (*NSOrderedCollectionChange, error) {
	done := make(chan *NSOrderedCollectionChange, 1)
	o.DifferenceByTransformingChangesWithBlock(func(val *NSOrderedCollectionChange) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

