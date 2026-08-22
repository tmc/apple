// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSPersistentCloudKitContainer] class.
var (
	_NSPersistentCloudKitContainerClass     NSPersistentCloudKitContainerClass
	_NSPersistentCloudKitContainerClassOnce sync.Once
)

func getNSPersistentCloudKitContainerClass() NSPersistentCloudKitContainerClass {
	_NSPersistentCloudKitContainerClassOnce.Do(func() {
		_NSPersistentCloudKitContainerClass = NSPersistentCloudKitContainerClass{class: objc.GetClass("NSPersistentCloudKitContainer")}
	})
	return _NSPersistentCloudKitContainerClass
}

// GetNSPersistentCloudKitContainerClass returns the class object for NSPersistentCloudKitContainer.
func GetNSPersistentCloudKitContainerClass() NSPersistentCloudKitContainerClass {
	return getNSPersistentCloudKitContainerClass()
}

type NSPersistentCloudKitContainerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPersistentCloudKitContainerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPersistentCloudKitContainerClass) Alloc() NSPersistentCloudKitContainer {
	rv := objc.Send[NSPersistentCloudKitContainer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A container that encapsulates the Core Data stack in your app, and mirrors
// select persistent stores to a CloudKit private database.
//
// # Overview
//
// [NSPersistentCloudKitContainer] is a subclass of [NSPersistentContainer]
// capable of managing both CloudKit-backed and noncloud stores.
//
// By default, [NSPersistentCloudKitContainer] contains a single store
// description, which Core Data assigns to the first CloudKit container
// identifier in an app’s entitlements. Use
// [NSPersistentCloudKitContainerOptions] to customize this behavior or create
// additional store descriptions with backing by different containers.
//
// For more information about setting up multiple stores, see [Setting Up Core
// Data with CloudKit].
//
// # Checking Permissions
//
//   - [NSPersistentCloudKitContainer.CanUpdateRecordForManagedObjectWithID]: Returns a Boolean value that indicates whether the user can modify the managed object’s underlying CloudKit record.
//   - [NSPersistentCloudKitContainer.CanDeleteRecordForManagedObjectWithID]: Returns a Boolean value that indicates whether the user can delete the managed object’s underlying CloudKit record.
//   - [NSPersistentCloudKitContainer.CanModifyManagedObjectsInStore]: Returns a Boolean value that indicates whether the user can modify the specified persistent store.
//
// # Promoting Your Schema
//
//   - [NSPersistentCloudKitContainer.InitializeCloudKitSchemaWithOptionsError]: Creates the CloudKit schema for all stores in the container that manage a CloudKit database.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer
//
// [Setting Up Core Data with CloudKit]: https://developer.apple.com/documentation/CoreData/setting-up-core-data-with-cloudkit
type NSPersistentCloudKitContainer struct {
	NSPersistentContainer
}

// NSPersistentCloudKitContainerFromID constructs a [NSPersistentCloudKitContainer] from an objc.ID.
//
// A container that encapsulates the Core Data stack in your app, and mirrors
// select persistent stores to a CloudKit private database.
func NSPersistentCloudKitContainerFromID(id objc.ID) NSPersistentCloudKitContainer {
	return NSPersistentCloudKitContainer{NSPersistentContainer: NSPersistentContainerFromID(id)}
}

// NOTE: NSPersistentCloudKitContainer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPersistentCloudKitContainer] class.
//
// # Checking Permissions
//
//   - [INSPersistentCloudKitContainer.CanUpdateRecordForManagedObjectWithID]: Returns a Boolean value that indicates whether the user can modify the managed object’s underlying CloudKit record.
//   - [INSPersistentCloudKitContainer.CanDeleteRecordForManagedObjectWithID]: Returns a Boolean value that indicates whether the user can delete the managed object’s underlying CloudKit record.
//   - [INSPersistentCloudKitContainer.CanModifyManagedObjectsInStore]: Returns a Boolean value that indicates whether the user can modify the specified persistent store.
//
// # Promoting Your Schema
//
//   - [INSPersistentCloudKitContainer.InitializeCloudKitSchemaWithOptionsError]: Creates the CloudKit schema for all stores in the container that manage a CloudKit database.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer
type INSPersistentCloudKitContainer interface {
	INSPersistentContainer

	// Topic: Checking Permissions

	// Returns a Boolean value that indicates whether the user can modify the managed object’s underlying CloudKit record.
	CanUpdateRecordForManagedObjectWithID(objectID INSManagedObjectID) bool
	// Returns a Boolean value that indicates whether the user can delete the managed object’s underlying CloudKit record.
	CanDeleteRecordForManagedObjectWithID(objectID INSManagedObjectID) bool
	// Returns a Boolean value that indicates whether the user can modify the specified persistent store.
	CanModifyManagedObjectsInStore(store INSPersistentStore) bool

	// Topic: Promoting Your Schema

	// Creates the CloudKit schema for all stores in the container that manage a CloudKit database.
	InitializeCloudKitSchemaWithOptionsError(options NSPersistentCloudKitContainerSchemaInitializationOptions) (NSPersistentCloudKitContainer, error)
}

// Init initializes the instance.
func (p NSPersistentCloudKitContainer) Init() NSPersistentCloudKitContainer {
	rv := objc.Send[NSPersistentCloudKitContainer](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPersistentCloudKitContainer) Autorelease() NSPersistentCloudKitContainer {
	rv := objc.Send[NSPersistentCloudKitContainer](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPersistentCloudKitContainer creates a new NSPersistentCloudKitContainer instance.
func NewNSPersistentCloudKitContainer() NSPersistentCloudKitContainer {
	class := getNSPersistentCloudKitContainerClass()
	rv := objc.Send[NSPersistentCloudKitContainer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a container with the specified name.
//
// name: The name of the [NSPersistentContainer] object.
//
// # Return Value
//
// A persistent container initialized with the given name.
//
// # Discussion
//
// By default, the provided name value is used to name the persistent store
// and is used to look up the name of the [NSManagedObjectModel] object to be
// used with the [NSPersistentContainer] object.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/init(name:)
func NewPersistentCloudKitContainerWithName(name string) NSPersistentCloudKitContainer {
	instance := getNSPersistentCloudKitContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:"), objc.String(name))
	return NSPersistentCloudKitContainerFromID(rv)
}

// Create a container with the specified name and managed object model.
//
// name: The name used by the persistent container.
//
// model: The managed object model to be used by the persistent container.
//
// # Return Value
//
// A persistent container initialized with the given name and model.
//
// # Discussion
//
// By default, the provided name value of the container is used as the name of
// the persisent store associated with the container. Passing in the
// [NSManagedObjectModel] object overrides the lookup of the model by the
// provided name value.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/init(name:managedObjectModel:)
func NewPersistentCloudKitContainerWithNameManagedObjectModel(name string, model INSManagedObjectModel) NSPersistentCloudKitContainer {
	instance := getNSPersistentCloudKitContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:managedObjectModel:"), objc.String(name), model)
	return NSPersistentCloudKitContainerFromID(rv)
}

// Returns a Boolean value that indicates whether the user can modify the
// managed object’s underlying CloudKit record.
//
// objectID: The ID of the managed object.
//
// # Return Value
//
// true if the user can modify the CloudKit record; otherwise, false.
//
// # Discussion
//
// This method returns true if
// [NSPersistentCloudKitContainer.CanModifyManagedObjectsInStore] returns true
// and any of the following conditions are true:
//
// - `objectID` is a temporary object identifier. - The persistent store that
// contains the managed object isn’t using CloudKit. - The persistent store
// manages the user’s private database. - The persistent store manages the
// public database, and the user owns the underlying record or Core Data has
// yet to save the managed object to iCloud. - The persistent store manages
// the shared database, and the user has the necessary permissions to update
// the managed object’s underlying record. For more information, see
// [CKShare.ParticipantPermission].
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/canUpdateRecord(forManagedObjectWith:)
//
// [CKShare.ParticipantPermission]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantPermission
func (p NSPersistentCloudKitContainer) CanUpdateRecordForManagedObjectWithID(objectID INSManagedObjectID) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canUpdateRecordForManagedObjectWithID:"), objectID)
	return rv
}

// Returns a Boolean value that indicates whether the user can delete the
// managed object’s underlying CloudKit record.
//
// objectID: The ID of the managed object.
//
// # Return Value
//
// true if the user can delete the CloudKit record; otherwise, false.
//
// # Discussion
//
// This method returns true if
// [NSPersistentCloudKitContainer.CanModifyManagedObjectsInStore] returns true
// and any of the following conditions are true:
//
// - `objectID` is a temporary object identifier. - The persistent store that
// contains the managed object isn’t using CloudKit. - The persistent store
// manages the user’s private database. - The persistent store manages the
// public database, and the user owns the underlying record or Core Data has
// yet to save the managed object to iCloud. - The persistent store manages
// the shared database, and the user has the necessary permissions to delete
// the managed object’s underlying record. For more information, see
// [CKShare.ParticipantPermission].
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/canDeleteRecord(forManagedObjectWith:)
//
// [CKShare.ParticipantPermission]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantPermission
func (p NSPersistentCloudKitContainer) CanDeleteRecordForManagedObjectWithID(objectID INSManagedObjectID) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canDeleteRecordForManagedObjectWithID:"), objectID)
	return rv
}

// Returns a Boolean value that indicates whether the user can modify the
// specified persistent store.
//
// store: The persistent store.
//
// # Return Value
//
// true if the user can modify records in the persistent store’s CloudKit
// database; otherwise, false.
//
// # Discussion
//
// Use this method to determine whether the user is able to write any records
// to the CloudKit database. To find out if the user can modify a specific
// object, use the
// [NSPersistentCloudKitContainer.CanUpdateRecordForManagedObjectWithID] and
// [NSPersistentCloudKitContainer.CanDeleteRecordForManagedObjectWithID]
// methods instead.
//
// This method always returns true for persistent stores that manage the
// user’s private CloudKit database.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/canModifyManagedObjects(in:)
func (p NSPersistentCloudKitContainer) CanModifyManagedObjectsInStore(store INSPersistentStore) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canModifyManagedObjectsInStore:"), store)
	return rv
}

// Creates the CloudKit schema for all stores in the container that manage a
// CloudKit database.
//
// options: The options to use when creating the CloudKit schema.
//
// # Discussion
//
// To create the schema, this method creates a set of representative
// [CKRecord] instances for all stores in the container that use Core Data
// with CloudKit, and uploads them to CloudKit. These records have a
// representative value for every field Core Data might serialize for the
// specified managed object model. After successfully uploading the records,
// the schema is visible in the CloudKit Dashboard and the container deletes
// the representative records.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/initializeCloudKitSchema(options:)
//
// [CKRecord]: https://developer.apple.com/documentation/CloudKit/CKRecord
func (p NSPersistentCloudKitContainer) InitializeCloudKitSchemaWithOptionsError(options NSPersistentCloudKitContainerSchemaInitializationOptions) (NSPersistentCloudKitContainer, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](p.ID, objc.Sel("initializeCloudKitSchemaWithOptions:error:"), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSPersistentCloudKitContainer{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSPersistentCloudKitContainerFromID(rv), nil

}
