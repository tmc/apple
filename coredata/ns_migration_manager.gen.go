// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMigrationManager] class.
var (
	_NSMigrationManagerClass     NSMigrationManagerClass
	_NSMigrationManagerClassOnce sync.Once
)

func getNSMigrationManagerClass() NSMigrationManagerClass {
	_NSMigrationManagerClassOnce.Do(func() {
		_NSMigrationManagerClass = NSMigrationManagerClass{class: objc.GetClass("NSMigrationManager")}
	})
	return _NSMigrationManagerClass
}

// GetNSMigrationManagerClass returns the class object for NSMigrationManager.
func GetNSMigrationManagerClass() NSMigrationManagerClass {
	return getNSMigrationManagerClass()
}

type NSMigrationManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSMigrationManagerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMigrationManagerClass) Alloc() NSMigrationManager {
	rv := objc.Send[NSMigrationManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A migration manager instance that performs a migration of data from one
// persistent store to another using a given mapping model.
//
// # Creating a Migration Manager
//
//   - [NSMigrationManager.InitWithSourceModelDestinationModel]: Initializes a migration manager instance with given source and destination models.
//
// # Getting the Manager’s Configuration
//
//   - [NSMigrationManager.DestinationContext]: The managed object context the migration manager uses for writing the destination persistent store.
//   - [NSMigrationManager.DestinationModel]: The destination model for the migration manager.
//   - [NSMigrationManager.MappingModel]: The mapping model for the migration manager.
//   - [NSMigrationManager.SourceContext]: The managed object context the migration manager uses for reading the source persistent store.
//   - [NSMigrationManager.SourceModel]: The source model for the migration manager.
//   - [NSMigrationManager.DestinationEntityForEntityMapping]: Returns the entity description for the destination entity of a given entity mapping.
//   - [NSMigrationManager.SourceEntityForEntityMapping]: Returns the entity description for the source entity of a given entity mapping.
//
// # Customizing the Manager
//
//   - [NSMigrationManager.UserInfo]: The user info for the migration manager.
//   - [NSMigrationManager.SetUserInfo]
//   - [NSMigrationManager.UsesStoreSpecificMigrationManager]: A Boolean value that indicates whether the migration manager tries to use a store specific migration manager to perform the  migration.
//   - [NSMigrationManager.SetUsesStoreSpecificMigrationManager]
//
// # Managing Sources and Destinations
//
//   - [NSMigrationManager.AssociateSourceInstanceWithDestinationInstanceForEntityMapping]: Associates a given source managed object instance with an array of destination instances for a given property mapping.
//   - [NSMigrationManager.DestinationInstancesForEntityMappingNamedSourceInstances]: Returns the managed object instances created in the destination store for the named entity mapping for the given array of source instances.
//   - [NSMigrationManager.SourceInstancesForEntityMappingNamedDestinationInstances]: Returns the managed object instances in the source store used to create the given destination instances for the passed in property mapping.
//
// # Performing a Migration
//
//   - [NSMigrationManager.MigrateStoreFromURLTypeOptionsWithMappingModelToDestinationURLDestinationTypeDestinationOptionsError]: Migrates the store at a given source URL to the store at a given destination URL, performing all of the mappings specified in a given mapping model.
//
// # Monitoring a Migration’s Progress
//
//   - [NSMigrationManager.MigrationProgress]: A number between `0` and `1` that indicates the proportion of completeness of the migration.
//   - [NSMigrationManager.CurrentEntityMapping]: The entity mapping currently being processed.
//
// # Aborting a Migration
//
//   - [NSMigrationManager.CancelMigrationWithError]: Cancels the migration with a given error.
//   - [NSMigrationManager.Reset]: Resets the association tables for the migration.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager
type NSMigrationManager struct {
	objectivec.Object
}

// NSMigrationManagerFromID constructs a [NSMigrationManager] from an objc.ID.
//
// A migration manager instance that performs a migration of data from one
// persistent store to another using a given mapping model.
func NSMigrationManagerFromID(id objc.ID) NSMigrationManager {
	return NSMigrationManager{objectivec.Object{ID: id}}
}

// NOTE: NSMigrationManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMigrationManager] class.
//
// # Creating a Migration Manager
//
//   - [INSMigrationManager.InitWithSourceModelDestinationModel]: Initializes a migration manager instance with given source and destination models.
//
// # Getting the Manager’s Configuration
//
//   - [INSMigrationManager.DestinationContext]: The managed object context the migration manager uses for writing the destination persistent store.
//   - [INSMigrationManager.DestinationModel]: The destination model for the migration manager.
//   - [INSMigrationManager.MappingModel]: The mapping model for the migration manager.
//   - [INSMigrationManager.SourceContext]: The managed object context the migration manager uses for reading the source persistent store.
//   - [INSMigrationManager.SourceModel]: The source model for the migration manager.
//   - [INSMigrationManager.DestinationEntityForEntityMapping]: Returns the entity description for the destination entity of a given entity mapping.
//   - [INSMigrationManager.SourceEntityForEntityMapping]: Returns the entity description for the source entity of a given entity mapping.
//
// # Customizing the Manager
//
//   - [INSMigrationManager.UserInfo]: The user info for the migration manager.
//   - [INSMigrationManager.SetUserInfo]
//   - [INSMigrationManager.UsesStoreSpecificMigrationManager]: A Boolean value that indicates whether the migration manager tries to use a store specific migration manager to perform the  migration.
//   - [INSMigrationManager.SetUsesStoreSpecificMigrationManager]
//
// # Managing Sources and Destinations
//
//   - [INSMigrationManager.AssociateSourceInstanceWithDestinationInstanceForEntityMapping]: Associates a given source managed object instance with an array of destination instances for a given property mapping.
//   - [INSMigrationManager.DestinationInstancesForEntityMappingNamedSourceInstances]: Returns the managed object instances created in the destination store for the named entity mapping for the given array of source instances.
//   - [INSMigrationManager.SourceInstancesForEntityMappingNamedDestinationInstances]: Returns the managed object instances in the source store used to create the given destination instances for the passed in property mapping.
//
// # Performing a Migration
//
//   - [INSMigrationManager.MigrateStoreFromURLTypeOptionsWithMappingModelToDestinationURLDestinationTypeDestinationOptionsError]: Migrates the store at a given source URL to the store at a given destination URL, performing all of the mappings specified in a given mapping model.
//
// # Monitoring a Migration’s Progress
//
//   - [INSMigrationManager.MigrationProgress]: A number between `0` and `1` that indicates the proportion of completeness of the migration.
//   - [INSMigrationManager.CurrentEntityMapping]: The entity mapping currently being processed.
//
// # Aborting a Migration
//
//   - [INSMigrationManager.CancelMigrationWithError]: Cancels the migration with a given error.
//   - [INSMigrationManager.Reset]: Resets the association tables for the migration.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager
type INSMigrationManager interface {
	objectivec.IObject

	// Topic: Creating a Migration Manager

	// Initializes a migration manager instance with given source and destination models.
	InitWithSourceModelDestinationModel(sourceModel INSManagedObjectModel, destinationModel INSManagedObjectModel) NSMigrationManager

	// Topic: Getting the Manager’s Configuration

	// The managed object context the migration manager uses for writing the destination persistent store.
	DestinationContext() INSManagedObjectContext
	// The destination model for the migration manager.
	DestinationModel() INSManagedObjectModel
	// The mapping model for the migration manager.
	MappingModel() INSMappingModel
	// The managed object context the migration manager uses for reading the source persistent store.
	SourceContext() INSManagedObjectContext
	// The source model for the migration manager.
	SourceModel() INSManagedObjectModel
	// Returns the entity description for the destination entity of a given entity mapping.
	DestinationEntityForEntityMapping(mEntity INSEntityMapping) INSEntityDescription
	// Returns the entity description for the source entity of a given entity mapping.
	SourceEntityForEntityMapping(mEntity INSEntityMapping) INSEntityDescription

	// Topic: Customizing the Manager

	// The user info for the migration manager.
	UserInfo() foundation.INSDictionary
	SetUserInfo(value foundation.INSDictionary)
	// A Boolean value that indicates whether the migration manager tries to use a store specific migration manager to perform the  migration.
	UsesStoreSpecificMigrationManager() bool
	SetUsesStoreSpecificMigrationManager(value bool)

	// Topic: Managing Sources and Destinations

	// Associates a given source managed object instance with an array of destination instances for a given property mapping.
	AssociateSourceInstanceWithDestinationInstanceForEntityMapping(sourceInstance INSManagedObject, destinationInstance INSManagedObject, entityMapping INSEntityMapping)
	// Returns the managed object instances created in the destination store for the named entity mapping for the given array of source instances.
	DestinationInstancesForEntityMappingNamedSourceInstances(mappingName string, sourceInstances []NSManagedObject) []NSManagedObject
	// Returns the managed object instances in the source store used to create the given destination instances for the passed in property mapping.
	SourceInstancesForEntityMappingNamedDestinationInstances(mappingName string, destinationInstances []NSManagedObject) []NSManagedObject

	// Topic: Performing a Migration

	// Migrates the store at a given source URL to the store at a given destination URL, performing all of the mappings specified in a given mapping model.
	MigrateStoreFromURLTypeOptionsWithMappingModelToDestinationURLDestinationTypeDestinationOptionsError(sourceURL foundation.NSURL, sStoreType string, sOptions foundation.INSDictionary, mappings INSMappingModel, dURL foundation.NSURL, dStoreType string, dOptions foundation.INSDictionary) (bool, error)

	// Topic: Monitoring a Migration’s Progress

	// A number between `0` and `1` that indicates the proportion of completeness of the migration.
	MigrationProgress() float32
	// The entity mapping currently being processed.
	CurrentEntityMapping() INSEntityMapping

	// Topic: Aborting a Migration

	// Cancels the migration with a given error.
	CancelMigrationWithError(error_ foundation.NSError)
	// Resets the association tables for the migration.
	Reset()
}

// Init initializes the instance.
func (m NSMigrationManager) Init() NSMigrationManager {
	rv := objc.Send[NSMigrationManager](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMigrationManager) Autorelease() NSMigrationManager {
	rv := objc.Send[NSMigrationManager](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMigrationManager creates a new NSMigrationManager instance.
func NewNSMigrationManager() NSMigrationManager {
	class := getNSMigrationManagerClass()
	rv := objc.Send[NSMigrationManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a migration manager instance with given source and destination
// models.
//
// sourceModel: The source managed object model for the migration manager.
//
// destinationModel: The destination managed object model for the migration manager.
//
// # Return Value
//
// A migration manager instance initialized to migrate data in a store that
// uses `sourceModel` to a store that uses `destinationModel`.
//
// # Discussion
//
// You specify the mapping model in the migration method,
// [NSMigrationManager.MigrateStoreFromURLTypeOptionsWithMappingModelToDestinationURLDestinationTypeDestinationOptionsError].
//
// # Special Considerations
//
// This is the designated initializer for [NSMigrationManager].
//
// Although validation of the models is performed during
// [NSMigrationManager.MigrateStoreFromURLTypeOptionsWithMappingModelToDestinationURLDestinationTypeDestinationOptionsError],
// as with [NSPersistentStoreCoordinator] once models are added to the
// migration manager they are immutable and cannot be altered.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager/init(sourceModel:destinationModel:)
func NewMigrationManagerWithSourceModelDestinationModel(sourceModel INSManagedObjectModel, destinationModel INSManagedObjectModel) NSMigrationManager {
	instance := getNSMigrationManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceModel:destinationModel:"), sourceModel, destinationModel)
	return NSMigrationManagerFromID(rv)
}

// Initializes a migration manager instance with given source and destination
// models.
//
// sourceModel: The source managed object model for the migration manager.
//
// destinationModel: The destination managed object model for the migration manager.
//
// # Return Value
//
// A migration manager instance initialized to migrate data in a store that
// uses `sourceModel` to a store that uses `destinationModel`.
//
// # Discussion
//
// You specify the mapping model in the migration method,
// [NSMigrationManager.MigrateStoreFromURLTypeOptionsWithMappingModelToDestinationURLDestinationTypeDestinationOptionsError].
//
// # Special Considerations
//
// This is the designated initializer for [NSMigrationManager].
//
// Although validation of the models is performed during
// [NSMigrationManager.MigrateStoreFromURLTypeOptionsWithMappingModelToDestinationURLDestinationTypeDestinationOptionsError],
// as with [NSPersistentStoreCoordinator] once models are added to the
// migration manager they are immutable and cannot be altered.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager/init(sourceModel:destinationModel:)
func (m NSMigrationManager) InitWithSourceModelDestinationModel(sourceModel INSManagedObjectModel, destinationModel INSManagedObjectModel) NSMigrationManager {
	rv := objc.Send[NSMigrationManager](m.ID, objc.Sel("initWithSourceModel:destinationModel:"), sourceModel, destinationModel)
	return rv
}

// Returns the entity description for the destination entity of a given entity
// mapping.
//
// mEntity: An entity mapping.
//
// # Return Value
//
// The entity description for the destination entity of `mEntity`.
//
// # Discussion
//
// Entity mappings do not store the actual description objects, but rather the
// name and version information of the entity.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager/destinationEntity(for:)
func (m NSMigrationManager) DestinationEntityForEntityMapping(mEntity INSEntityMapping) INSEntityDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("destinationEntityForEntityMapping:"), mEntity)
	return NSEntityDescriptionFromID(rv)
}

// Returns the entity description for the source entity of a given entity
// mapping.
//
// mEntity: An entity mapping.
//
// # Return Value
//
// The entity description for the source entity of `mEntity`.
//
// # Discussion
//
// Entity mappings do not store the actual description objects, but rather the
// name and version information of the entity.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager/sourceEntity(for:)
func (m NSMigrationManager) SourceEntityForEntityMapping(mEntity INSEntityMapping) INSEntityDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("sourceEntityForEntityMapping:"), mEntity)
	return NSEntityDescriptionFromID(rv)
}

// Associates a given source managed object instance with an array of
// destination instances for a given property mapping.
//
// sourceInstance: A source managed object.
//
// destinationInstance: The destination manage object for `sourceInstance`.
//
// entityMapping: The entity mapping to use to associate `sourceInstance` with the object in
// `destinationInstances`.
//
// # Discussion
//
// Data migration is performed as a three-stage process (first create the
// data, then relate the data, then validate the data). You use this method to
// associate data between the source and destination stores, in order to allow
// for relationship creation or fix-up after the creation stage.
//
// This method is called in the default implementation of
// [NSEntityMigrationPolicy]’s
// [NSEntityMigrationPolicy.CreateDestinationInstancesForSourceInstanceEntityMappingManagerError]
// method.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager/associate(sourceInstance:withDestinationInstance:for:)
func (m NSMigrationManager) AssociateSourceInstanceWithDestinationInstanceForEntityMapping(sourceInstance INSManagedObject, destinationInstance INSManagedObject, entityMapping INSEntityMapping) {
	objc.Send[objc.ID](m.ID, objc.Sel("associateSourceInstance:withDestinationInstance:forEntityMapping:"), sourceInstance, destinationInstance, entityMapping)
}

// Returns the managed object instances created in the destination store for
// the named entity mapping for the given array of source instances.
//
// mappingName: The name of an entity mapping in use.
//
// sourceInstances: A array of managed objects in the source store.
//
// # Return Value
//
// An array containing the managed object instances created in the destination
// store for the entity mapping named `mappingName` for `sourceInstances`. If
// `sourceInstances` is `nil`, all of the destination instances created by the
// specified property mapping are returned.
//
// # Discussion
//
// This method throws an [NSInvalidArgumentException] exception if
// `mappingName` is not a valid mapping name.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager/destinationInstances(forEntityMappingName:sourceInstances:)
func (m NSMigrationManager) DestinationInstancesForEntityMappingNamedSourceInstances(mappingName string, sourceInstances []NSManagedObject) []NSManagedObject {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("destinationInstancesForEntityMappingNamed:sourceInstances:"), objc.String(mappingName), objectivec.IObjectSliceToNSArray(sourceInstances))
	return objc.ConvertSlice(rv, func(id objc.ID) NSManagedObject {
		return NSManagedObjectFromID(id)
	})
}

// Returns the managed object instances in the source store used to create the
// given destination instances for the passed in property mapping.
//
// mappingName: The name of an entity mapping in use.
//
// destinationInstances: A array of managed objects in the destination store.
//
// # Return Value
//
// An array containing the managed object instances in the source store used
// to create `destinationInstances` using the entity mapping named
// `mappingName`. If `destinationInstances` is `nil`, all of the source
// instances used to create the destination instance for this property mapping
// are returned.
//
// # Discussion
//
// This method throws an [NSInvalidArgumentException] exception if
// `mappingName` is not a valid mapping name.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager/sourceInstances(forEntityMappingName:destinationInstances:)
func (m NSMigrationManager) SourceInstancesForEntityMappingNamedDestinationInstances(mappingName string, destinationInstances []NSManagedObject) []NSManagedObject {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("sourceInstancesForEntityMappingNamed:destinationInstances:"), objc.String(mappingName), objectivec.IObjectSliceToNSArray(destinationInstances))
	return objc.ConvertSlice(rv, func(id objc.ID) NSManagedObject {
		return NSManagedObjectFromID(id)
	})
}

// Migrates the store at a given source URL to the store at a given
// destination URL, performing all of the mappings specified in a given
// mapping model.
//
// sourceURL: The location of an existing persistent store. A store must exist at this
// URL.
//
// sStoreType: The type of store at `sourceURL` (see [NSPersistentStoreCoordinator] for
// possible values).
//
// sOptions: A dictionary of options for the source (see [NSPersistentStoreCoordinator]
// for possible values).
//
// mappings: The mapping model to use to effect the migration.
//
// dURL: The location of the destination store.
//
// dStoreType: The type of store at `dURL` (see [NSPersistentStoreCoordinator] for
// possible values).
//
// dOptions: A dictionary of options for the destination (see
// [NSPersistentStoreCoordinator] for possible values).
//
// # Discussion
//
// This method performs compatibility checks on the source and destination
// models and the mapping model.
//
// # Special Considerations
//
// If a store does not exist at the destination URL (`dURL`), one is created;
// otherwise, the migration appends to the existing store.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager/migrateStore(from:sourceType:options:with:toDestinationURL:destinationType:destinationOptions:)
func (m NSMigrationManager) MigrateStoreFromURLTypeOptionsWithMappingModelToDestinationURLDestinationTypeDestinationOptionsError(sourceURL foundation.NSURL, sStoreType string, sOptions foundation.INSDictionary, mappings INSMappingModel, dURL foundation.NSURL, dStoreType string, dOptions foundation.INSDictionary) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("migrateStoreFromURL:type:options:withMappingModel:toDestinationURL:destinationType:destinationOptions:error:"), sourceURL, objc.String(sStoreType), sOptions, mappings, dURL, objc.String(dStoreType), dOptions, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("migrateStoreFromURL:type:options:withMappingModel:toDestinationURL:destinationType:destinationOptions:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Cancels the migration with a given error.
//
// error: An error object that describes the reason why the migration is canceled.
//
// # Discussion
//
// You can invoke this method from anywhere in the migration process to abort
// the migration. Calling this method causes
// [NSMigrationManager.MigrateStoreFromURLTypeOptionsWithMappingModelToDestinationURLDestinationTypeDestinationOptionsError]
// to abort the migration and return `error`—you should provide an
// appropriate error to indicate the reason for the cancellation.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager/cancelMigrationWithError(_:)
func (m NSMigrationManager) CancelMigrationWithError(error_ foundation.NSError) {
	objc.Send[objc.ID](m.ID, objc.Sel("cancelMigrationWithError:"), error_)
}

// Resets the association tables for the migration.
//
// # Discussion
//
// This method does not reset the source or destination contexts.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager/reset()
func (m NSMigrationManager) Reset() {
	objc.Send[objc.ID](m.ID, objc.Sel("reset"))
}

// The managed object context the migration manager uses for writing the
// destination persistent store.
//
// # Discussion
//
// This context is created on demand as part of the initialization of the Core
// Data stacks used for migration.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager/destinationContext
func (m NSMigrationManager) DestinationContext() INSManagedObjectContext {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("destinationContext"))
	return NSManagedObjectContextFromID(objc.ID(rv))
}

// The destination model for the migration manager.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager/destinationModel
func (m NSMigrationManager) DestinationModel() INSManagedObjectModel {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("destinationModel"))
	return NSManagedObjectModelFromID(objc.ID(rv))
}

// The mapping model for the migration manager.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager/mappingModel
func (m NSMigrationManager) MappingModel() INSMappingModel {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mappingModel"))
	return NSMappingModelFromID(objc.ID(rv))
}

// The managed object context the migration manager uses for reading the
// source persistent store.
//
// # Discussion
//
// This context is created on demand as part of the initialization of the Core
// Data stacks used for migration.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager/sourceContext
func (m NSMigrationManager) SourceContext() INSManagedObjectContext {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("sourceContext"))
	return NSManagedObjectContextFromID(objc.ID(rv))
}

// The source model for the migration manager.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager/sourceModel
func (m NSMigrationManager) SourceModel() INSManagedObjectModel {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("sourceModel"))
	return NSManagedObjectModelFromID(objc.ID(rv))
}

// The user info for the migration manager.
//
// # Discussion
//
// You can use the user info dictionary to aid the customization of your
// migration process.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager/userInfo
func (m NSMigrationManager) UserInfo() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("userInfo"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m NSMigrationManager) SetUserInfo(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setUserInfo:"), value)
}

// A Boolean value that indicates whether the migration manager tries to use a
// store specific migration manager to perform the migration.
//
// # Return Value
//
// true if the receiver uses a store-specific migration manager, otherwise
// false.
//
// # Discussion
//
// true if the receiver uses a store-specific migration manager, otherwise
// false. The default value is true.
//
// A store-specific migration manager class is not guaranteed to perform any
// of the migration manager delegate callbacks or update values for the
// observable properties.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager/usesStoreSpecificMigrationManager
func (m NSMigrationManager) UsesStoreSpecificMigrationManager() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("usesStoreSpecificMigrationManager"))
	return rv
}
func (m NSMigrationManager) SetUsesStoreSpecificMigrationManager(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setUsesStoreSpecificMigrationManager:"), value)
}

// A number between `0` and `1` that indicates the proportion of completeness
// of the migration.
//
// # Discussion
//
// If a migration is not taking place, this property is `1`. You can observe
// this value using key-value observing.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager/migrationProgress
func (m NSMigrationManager) MigrationProgress() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("migrationProgress"))
	return rv
}

// The entity mapping currently being processed.
//
// # Discussion
//
// Each entity is processed a total of three times—instance creation,
// relationship creation, and validation.
//
// # Special Considerations
//
// You can observe this value using key-value observing.
//
// See: https://developer.apple.com/documentation/CoreData/NSMigrationManager/currentEntityMapping
func (m NSMigrationManager) CurrentEntityMapping() INSEntityMapping {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("currentEntityMapping"))
	return NSEntityMappingFromID(objc.ID(rv))
}
