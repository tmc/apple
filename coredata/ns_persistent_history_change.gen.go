// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPersistentHistoryChange] class.
var (
	_NSPersistentHistoryChangeClass     NSPersistentHistoryChangeClass
	_NSPersistentHistoryChangeClassOnce sync.Once
)

func getNSPersistentHistoryChangeClass() NSPersistentHistoryChangeClass {
	_NSPersistentHistoryChangeClassOnce.Do(func() {
		_NSPersistentHistoryChangeClass = NSPersistentHistoryChangeClass{class: objc.GetClass("NSPersistentHistoryChange")}
	})
	return _NSPersistentHistoryChangeClass
}

// GetNSPersistentHistoryChangeClass returns the class object for NSPersistentHistoryChange.
func GetNSPersistentHistoryChangeClass() NSPersistentHistoryChangeClass {
	return getNSPersistentHistoryChangeClass()
}

type NSPersistentHistoryChangeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPersistentHistoryChangeClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPersistentHistoryChangeClass) Alloc() NSPersistentHistoryChange {
	rv := objc.Send[NSPersistentHistoryChange](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A change representing the insertion, update, or deletion of a managed
// object in the persistent store.
//
// # Inspecting Change Details
//
//   - [NSPersistentHistoryChange.ChangeID]: The change’s numeric identifier.
//   - [NSPersistentHistoryChange.ChangeType]: The type of change to the managed object in the persistent store.
//   - [NSPersistentHistoryChange.ChangedObjectID]: The identifier of the managed object that changed. (swift) Declaration: @property(readonly, copy) NSManagedObjectID *changedObjectID; (objc) Availability: iOS: 11.0 — iPadOS: 11.0 — Mac Catalyst: 13.1 — macOS: 10.13 — tvOS: 11.0 — visionOS: 1.0 — watchOS: 4.0 (objc,swift) }
//   - [NSPersistentHistoryChange.Tombstone]: A dictionary of attributes marked for preservation after deletion, and their values when deleted.
//   - [NSPersistentHistoryChange.Transaction]: The persistent history transaction containing this change.
//   - [NSPersistentHistoryChange.UpdatedProperties]: The set of properties that were updated on the managed object.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange
type NSPersistentHistoryChange struct {
	objectivec.Object
}

// NSPersistentHistoryChangeFromID constructs a [NSPersistentHistoryChange] from an objc.ID.
//
// A change representing the insertion, update, or deletion of a managed
// object in the persistent store.
func NSPersistentHistoryChangeFromID(id objc.ID) NSPersistentHistoryChange {
	return NSPersistentHistoryChange{objectivec.Object{ID: id}}
}

// NOTE: NSPersistentHistoryChange adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPersistentHistoryChange] class.
//
// # Inspecting Change Details
//
//   - [INSPersistentHistoryChange.ChangeID]: The change’s numeric identifier.
//   - [INSPersistentHistoryChange.ChangeType]: The type of change to the managed object in the persistent store.
//   - [INSPersistentHistoryChange.ChangedObjectID]: The identifier of the managed object that changed. (swift) Declaration: @property(readonly, copy) NSManagedObjectID *changedObjectID; (objc) Availability: iOS: 11.0 — iPadOS: 11.0 — Mac Catalyst: 13.1 — macOS: 10.13 — tvOS: 11.0 — visionOS: 1.0 — watchOS: 4.0 (objc,swift) }
//   - [INSPersistentHistoryChange.Tombstone]: A dictionary of attributes marked for preservation after deletion, and their values when deleted.
//   - [INSPersistentHistoryChange.Transaction]: The persistent history transaction containing this change.
//   - [INSPersistentHistoryChange.UpdatedProperties]: The set of properties that were updated on the managed object.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange
type INSPersistentHistoryChange interface {
	objectivec.IObject

	// Topic: Inspecting Change Details

	// The change’s numeric identifier.
	ChangeID() int64
	// The type of change to the managed object in the persistent store.
	ChangeType() NSPersistentHistoryChangeType
	// The identifier of the managed object that changed. (swift) Declaration: @property(readonly, copy) NSManagedObjectID *changedObjectID; (objc) Availability: iOS: 11.0 — iPadOS: 11.0 — Mac Catalyst: 13.1 — macOS: 10.13 — tvOS: 11.0 — visionOS: 1.0 — watchOS: 4.0 (objc,swift) }
	ChangedObjectID() INSManagedObjectID
	// A dictionary of attributes marked for preservation after deletion, and their values when deleted.
	Tombstone() foundation.INSDictionary
	// The persistent history transaction containing this change.
	Transaction() INSPersistentHistoryTransaction
	// The set of properties that were updated on the managed object.
	UpdatedProperties() foundation.INSSet
}

// Init initializes the instance.
func (p NSPersistentHistoryChange) Init() NSPersistentHistoryChange {
	rv := objc.Send[NSPersistentHistoryChange](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPersistentHistoryChange) Autorelease() NSPersistentHistoryChange {
	rv := objc.Send[NSPersistentHistoryChange](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPersistentHistoryChange creates a new NSPersistentHistoryChange instance.
func NewNSPersistentHistoryChange() NSPersistentHistoryChange {
	class := getNSPersistentHistoryChangeClass()
	rv := objc.Send[NSPersistentHistoryChange](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Requests an entity description for the managed object type affected by the
// change using the provided context.
//
// context: The managed object context for this request.
//
// # Return Value
//
// The entity description ([NSEntityDescription]) of the persistent history
// transaction entity.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange/entityDescription(with:)
func (_NSPersistentHistoryChangeClass NSPersistentHistoryChangeClass) EntityDescriptionWithContext(context INSManagedObjectContext) NSEntityDescription {
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentHistoryChangeClass.class), objc.Sel("entityDescriptionWithContext:"), context)
	return NSEntityDescriptionFromID(rv)
}

// The change’s numeric identifier.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange/changeID
func (p NSPersistentHistoryChange) ChangeID() int64 {
	rv := objc.Send[int64](p.ID, objc.Sel("changeID"))
	return rv
}

// The type of change to the managed object in the persistent store.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange/changeType
func (p NSPersistentHistoryChange) ChangeType() NSPersistentHistoryChangeType {
	rv := objc.Send[NSPersistentHistoryChangeType](p.ID, objc.Sel("changeType"))
	return NSPersistentHistoryChangeType(rv)
}

// The identifier of the managed object that changed. (swift) Declaration:
// @property(readonly, copy) NSManagedObjectID *changedObjectID; (objc)
// Availability: iOS: 11.0 — iPadOS: 11.0 — Mac Catalyst: 13.1 — macOS:
// 10.13 — tvOS: 11.0 — visionOS: 1.0 — watchOS: 4.0 (objc,swift) }
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange/changedObjectID
func (p NSPersistentHistoryChange) ChangedObjectID() INSManagedObjectID {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("changedObjectID"))
	return NSManagedObjectIDFromID(objc.ID(rv))
}

// A dictionary of attributes marked for preservation after deletion, and
// their values when deleted.
//
// # Discussion
//
// This value is expected on changes of type
// [NSPersistentHistoryChangeType.delete].
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange/tombstone
//
// [NSPersistentHistoryChangeType.delete]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeType/delete
func (p NSPersistentHistoryChange) Tombstone() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("tombstone"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The persistent history transaction containing this change.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange/transaction
func (p NSPersistentHistoryChange) Transaction() INSPersistentHistoryTransaction {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("transaction"))
	return NSPersistentHistoryTransactionFromID(objc.ID(rv))
}

// The set of properties that were updated on the managed object.
//
// # Discussion
//
// This value is expected on changes of type
// [NSPersistentHistoryChangeType.update].
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange/updatedProperties
//
// [NSPersistentHistoryChangeType.update]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeType/update
func (p NSPersistentHistoryChange) UpdatedProperties() foundation.INSSet {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("updatedProperties"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// A fetch request that has the persistent history change as the entity.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange/fetchRequest
func (_NSPersistentHistoryChangeClass NSPersistentHistoryChangeClass) FetchRequest() NSFetchRequest {
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentHistoryChangeClass.class), objc.Sel("fetchRequest"))
	return NSFetchRequestFromID(objc.ID(rv))
}

// The entity description of the persistent history change entity.
//
// # Discussion
//
// The entity description of a [NSPersistentHistoryChange], includes its
// properties, which can be useful for filtering your persistent history
// change request.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange/entityDescription
func (_NSPersistentHistoryChangeClass NSPersistentHistoryChangeClass) EntityDescription() NSEntityDescription {
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentHistoryChangeClass.class), objc.Sel("entityDescription"))
	return NSEntityDescriptionFromID(objc.ID(rv))
}
