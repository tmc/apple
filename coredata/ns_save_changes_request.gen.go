// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSaveChangesRequest] class.
var (
	_NSSaveChangesRequestClass     NSSaveChangesRequestClass
	_NSSaveChangesRequestClassOnce sync.Once
)

func getNSSaveChangesRequestClass() NSSaveChangesRequestClass {
	_NSSaveChangesRequestClassOnce.Do(func() {
		_NSSaveChangesRequestClass = NSSaveChangesRequestClass{class: objc.GetClass("NSSaveChangesRequest")}
	})
	return _NSSaveChangesRequestClass
}

// GetNSSaveChangesRequestClass returns the class object for NSSaveChangesRequest.
func GetNSSaveChangesRequestClass() NSSaveChangesRequestClass {
	return getNSSaveChangesRequestClass()
}

type NSSaveChangesRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSaveChangesRequestClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSaveChangesRequestClass) Alloc() NSSaveChangesRequest {
	rv := objc.Send[NSSaveChangesRequest](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An encapsulation of a collection of changes to be made by an object store
// in response to a save operation on a managed object context.
//
// # Initializing a Request
//
//   - [NSSaveChangesRequest.InitWithInsertedObjectsUpdatedObjectsDeletedObjectsLockedObjects]: Initializes a save changes request with collections of given changes.
//
// # Getting Information about a Request
//
//   - [NSSaveChangesRequest.InsertedObjects]: The objects that were inserted into the calling context.
//   - [NSSaveChangesRequest.UpdatedObjects]: The objects that were modified in the calling context.
//   - [NSSaveChangesRequest.DeletedObjects]: The objects that were deleted in the calling context.
//   - [NSSaveChangesRequest.LockedObjects]: The objects that were flagged for optimistic locking on the calling context.
//
// See: https://developer.apple.com/documentation/CoreData/NSSaveChangesRequest
type NSSaveChangesRequest struct {
	NSPersistentStoreRequest
}

// NSSaveChangesRequestFromID constructs a [NSSaveChangesRequest] from an objc.ID.
//
// An encapsulation of a collection of changes to be made by an object store
// in response to a save operation on a managed object context.
func NSSaveChangesRequestFromID(id objc.ID) NSSaveChangesRequest {
	return NSSaveChangesRequest{NSPersistentStoreRequest: NSPersistentStoreRequestFromID(id)}
}

// NOTE: NSSaveChangesRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSaveChangesRequest] class.
//
// # Initializing a Request
//
//   - [INSSaveChangesRequest.InitWithInsertedObjectsUpdatedObjectsDeletedObjectsLockedObjects]: Initializes a save changes request with collections of given changes.
//
// # Getting Information about a Request
//
//   - [INSSaveChangesRequest.InsertedObjects]: The objects that were inserted into the calling context.
//   - [INSSaveChangesRequest.UpdatedObjects]: The objects that were modified in the calling context.
//   - [INSSaveChangesRequest.DeletedObjects]: The objects that were deleted in the calling context.
//   - [INSSaveChangesRequest.LockedObjects]: The objects that were flagged for optimistic locking on the calling context.
//
// See: https://developer.apple.com/documentation/CoreData/NSSaveChangesRequest
type INSSaveChangesRequest interface {
	INSPersistentStoreRequest

	// Topic: Initializing a Request

	// Initializes a save changes request with collections of given changes.
	InitWithInsertedObjectsUpdatedObjectsDeletedObjectsLockedObjects(insertedObjects foundation.INSSet, updatedObjects foundation.INSSet, deletedObjects foundation.INSSet, lockedObjects foundation.INSSet) NSSaveChangesRequest

	// Topic: Getting Information about a Request

	// The objects that were inserted into the calling context.
	InsertedObjects() foundation.INSSet
	// The objects that were modified in the calling context.
	UpdatedObjects() foundation.INSSet
	// The objects that were deleted in the calling context.
	DeletedObjects() foundation.INSSet
	// The objects that were flagged for optimistic locking on the calling context.
	LockedObjects() foundation.INSSet
}

// Init initializes the instance.
func (s NSSaveChangesRequest) Init() NSSaveChangesRequest {
	rv := objc.Send[NSSaveChangesRequest](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSaveChangesRequest) Autorelease() NSSaveChangesRequest {
	rv := objc.Send[NSSaveChangesRequest](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSaveChangesRequest creates a new NSSaveChangesRequest instance.
func NewNSSaveChangesRequest() NSSaveChangesRequest {
	class := getNSSaveChangesRequestClass()
	rv := objc.Send[NSSaveChangesRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a save changes request with collections of given changes.
//
// insertedObjects: Objects that were inserted into the calling context.
//
// updatedObjects: Objects that were updated in the calling context.
//
// deletedObjects: Objects that were deleted in the calling context.
//
// lockedObjects: Objects that were flagged for optimistic locking on the calling context.
//
// # Return Value
//
// A save changes request initialized with the given changes.
//
// See: https://developer.apple.com/documentation/CoreData/NSSaveChangesRequest/init(inserted:updated:deleted:locked:)
func NewSaveChangesRequestWithInsertedObjectsUpdatedObjectsDeletedObjectsLockedObjects(insertedObjects foundation.INSSet, updatedObjects foundation.INSSet, deletedObjects foundation.INSSet, lockedObjects foundation.INSSet) NSSaveChangesRequest {
	instance := getNSSaveChangesRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInsertedObjects:updatedObjects:deletedObjects:lockedObjects:"), insertedObjects, updatedObjects, deletedObjects, lockedObjects)
	return NSSaveChangesRequestFromID(rv)
}

// Initializes a save changes request with collections of given changes.
//
// insertedObjects: Objects that were inserted into the calling context.
//
// updatedObjects: Objects that were updated in the calling context.
//
// deletedObjects: Objects that were deleted in the calling context.
//
// lockedObjects: Objects that were flagged for optimistic locking on the calling context.
//
// # Return Value
//
// A save changes request initialized with the given changes.
//
// See: https://developer.apple.com/documentation/CoreData/NSSaveChangesRequest/init(inserted:updated:deleted:locked:)
func (s NSSaveChangesRequest) InitWithInsertedObjectsUpdatedObjectsDeletedObjectsLockedObjects(insertedObjects foundation.INSSet, updatedObjects foundation.INSSet, deletedObjects foundation.INSSet, lockedObjects foundation.INSSet) NSSaveChangesRequest {
	rv := objc.Send[NSSaveChangesRequest](s.ID, objc.Sel("initWithInsertedObjects:updatedObjects:deletedObjects:lockedObjects:"), insertedObjects, updatedObjects, deletedObjects, lockedObjects)
	return rv
}

// The objects that were inserted into the calling context.
//
// See: https://developer.apple.com/documentation/CoreData/NSSaveChangesRequest/insertedObjects
func (s NSSaveChangesRequest) InsertedObjects() foundation.INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("insertedObjects"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The objects that were modified in the calling context.
//
// See: https://developer.apple.com/documentation/CoreData/NSSaveChangesRequest/updatedObjects
func (s NSSaveChangesRequest) UpdatedObjects() foundation.INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("updatedObjects"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The objects that were deleted in the calling context.
//
// See: https://developer.apple.com/documentation/CoreData/NSSaveChangesRequest/deletedObjects
func (s NSSaveChangesRequest) DeletedObjects() foundation.INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("deletedObjects"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The objects that were flagged for optimistic locking on the calling
// context.
//
// # Discussion
//
// Objects are flagged for optimistic locking with
// [NSManagedObjectContext.DetectConflictsForObject].
//
// See: https://developer.apple.com/documentation/CoreData/NSSaveChangesRequest/lockedObjects
func (s NSSaveChangesRequest) LockedObjects() foundation.INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("lockedObjects"))
	return foundation.NSSetFromID(objc.ID(rv))
}
