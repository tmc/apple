// Code generated from Apple documentation. DO NOT EDIT.

package coredata

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var (
	// AddedPersistentStoresKey is key for the array of stores that were added.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSAddedPersistentStoresKey
	AddedPersistentStoresKey string
	// AffectedObjectsErrorKey is the key for objects prompting an error.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSAffectedObjectsErrorKey
	AffectedObjectsErrorKey string
	// AffectedStoresErrorKey is the key for stores prompting an error.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSAffectedStoresErrorKey
	AffectedStoresErrorKey string
	// BinaryStoreInsecureDecodingCompatibilityOption is a flag that indicates Core Data decodes the binary store insecurely.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSBinaryStoreInsecureDecodingCompatibilityOption
	BinaryStoreInsecureDecodingCompatibilityOption string
	// BinaryStoreSecureDecodingClasses is an additional set of classes to use while decoding a binary store.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSBinaryStoreSecureDecodingClasses
	BinaryStoreSecureDecodingClasses string
	// BinaryStoreType is the binary store type.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSBinaryStoreType
	BinaryStoreType string
	// CoreDataCoreSpotlightExporter is the key you use to specify your Core Spotlight delegate.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightExporter
	CoreDataCoreSpotlightExporter string
	// DeletedObjectIDsKey is a user info key to identify deleted object identifiers in notifications after saving a managed object context.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSDeletedObjectIDsKey
	DeletedObjectIDsKey string
	// DeletedObjectsKey is a key for the set of objects that were marked for deletion during the previous event.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSDeletedObjectsKey
	DeletedObjectsKey string
	// DetailedErrorsKey is if multiple validation errors occur in one operation, they are collected in an array and added with this key to the “top-level error” of the operation.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSDetailedErrorsKey
	DetailedErrorsKey string
	// IgnorePersistentStoreVersioningOption is key to ignore the built-in versioning provided by Core Data.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSIgnorePersistentStoreVersioningOption
	IgnorePersistentStoreVersioningOption string
	// InMemoryStoreType is the in-memory store type.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSInMemoryStoreType
	InMemoryStoreType string
	// InferMappingModelAutomaticallyOption is key to attempt to create the mapping model automatically.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSInferMappingModelAutomaticallyOption
	InferMappingModelAutomaticallyOption string
	// InsertedObjectIDsKey is a user info key to identify inserted object identifiers in notifications after saving a managed object context.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSInsertedObjectIDsKey
	InsertedObjectIDsKey string
	// InsertedObjectsKey is a key for the set of objects that were inserted into the context.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSInsertedObjectsKey
	InsertedObjectsKey string
	// InvalidatedAllObjectsKey is a key that specifies that all objects in the context have been invalidated.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSInvalidatedAllObjectsKey
	InvalidatedAllObjectsKey string
	// InvalidatedObjectIDsKey is a user info key to identify invalidated object identifiers in notifications after saving a managed object context.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSInvalidatedObjectIDsKey
	InvalidatedObjectIDsKey string
	// InvalidatedObjectsKey is a key for the set of objects that were invalidated.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSInvalidatedObjectsKey
	InvalidatedObjectsKey string
	// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContextDidMergeChangesObjectIDsNotification
	ManagedObjectContextDidMergeChangesObjectIDsNotification string
	// ManagedObjectContextDidSaveNotification is a notification that posts after a context finishes writing unsaved changes.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContextDidSaveNotification
	ManagedObjectContextDidSaveNotification string
	// ManagedObjectContextDidSaveObjectIDsNotification is a notification that posts after a context finishes writing changes.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContextDidSaveObjectIDsNotification
	ManagedObjectContextDidSaveObjectIDsNotification string
	// ManagedObjectContextObjectsDidChangeNotification is a notification that posts when there are changes to context’s registered managed objects.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContextObjectsDidChangeNotification
	ManagedObjectContextObjectsDidChangeNotification string
	// ManagedObjectContextQueryGenerationKey is constant used to reference the query generation token.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContextQueryGenerationKey
	ManagedObjectContextQueryGenerationKey string
	// ManagedObjectContextWillSaveNotification is a notification that posts before a context writes unsaved changes.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContextWillSaveNotification
	ManagedObjectContextWillSaveNotification string
	// MigratePersistentStoresAutomaticallyOption is key to automatically attempt to migrate versioned stores.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSMigratePersistentStoresAutomaticallyOption
	MigratePersistentStoresAutomaticallyOption string
	// MigrationDestinationObjectKey is key for the destination object.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSMigrationDestinationObjectKey
	MigrationDestinationObjectKey string
	// MigrationEntityMappingKey is key for the entity mapping object.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSMigrationEntityMappingKey
	MigrationEntityMappingKey string
	// MigrationEntityPolicyKey is key for the entity migration policy object.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSMigrationEntityPolicyKey
	MigrationEntityPolicyKey string
	// MigrationManagerKey is key for the migration manager.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSMigrationManagerKey
	MigrationManagerKey string
	// MigrationPropertyMappingKey is key for the property mapping object.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSMigrationPropertyMappingKey
	MigrationPropertyMappingKey string
	// MigrationSourceObjectKey is key for the source object.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSMigrationSourceObjectKey
	MigrationSourceObjectKey string
	// PersistentCloudKitContainerEventUserInfoKey is the user info dictionary key for the persistent CloudKit container event.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/eventNotificationUserInfoKey
	PersistentCloudKitContainerEventUserInfoKey string
	// PersistentHistoryTokenKey is a user info key to identify the history token in persistent store remote change notifications.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTokenKey
	PersistentHistoryTokenKey string
	// PersistentHistoryTrackingKey is the key you use to enable persistent history tracking.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTrackingKey
	PersistentHistoryTrackingKey string
	// PersistentStoreConnectionPoolMaxSizeKey is the maximum connection pool size to use on a store that supports concurrent request handling.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreConnectionPoolMaxSizeKey
	PersistentStoreConnectionPoolMaxSizeKey string
	// PersistentStoreCoordinatorStoresDidChangeNotification is a notification that the coordinator posts after its registered stores change.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinatorStoresDidChangeNotification
	PersistentStoreCoordinatorStoresDidChangeNotification string
	// PersistentStoreCoordinatorStoresWillChangeNotification is a notification that posts before a coordinator changes its registered stores.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinatorStoresWillChangeNotification
	PersistentStoreCoordinatorStoresWillChangeNotification string
	// PersistentStoreCoordinatorWillRemoveStoreNotification is a notification that posts before a coordinator removes a store.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinatorWillRemoveStoreNotification
	PersistentStoreCoordinatorWillRemoveStoreNotification string
	// PersistentStoreDeferredLightweightMigrationOptionKey is the key for enabling deferred lightweight migrations.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDeferredLightweightMigrationOptionKey
	PersistentStoreDeferredLightweightMigrationOptionKey string
	// PersistentStoreForceDestroyOption is a flag that indicates the coordinator destroys the store file even if the operation might be unsafe, overriding locks, if necessary.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreForceDestroyOption
	PersistentStoreForceDestroyOption string
	// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreModelVersionChecksumKey
	PersistentStoreModelVersionChecksumKey string
	// PersistentStoreOSCompatibility is key to represent the earliest version of the operation system that the persistent store supports.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreOSCompatibility
	PersistentStoreOSCompatibility string
	// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreRemoteChangeNotification
	PersistentStoreRemoteChangeNotification string
	// PersistentStoreRemoteChangeNotificationPostOptionKey is a key that indicates a persistent store posts a remote change notification for every write to the store, including writes by other processes.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreRemoteChangeNotificationPostOptionKey
	PersistentStoreRemoteChangeNotificationPostOptionKey string
	// PersistentStoreSaveConflictsErrorKey is the key for the array of merge conflict objects (instances of [NSMergeConflict]).
	//
	// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreSaveConflictsErrorKey
	PersistentStoreSaveConflictsErrorKey string
	// PersistentStoreStagedMigrationManagerOptionKey is the key for specifying your staged migration manager.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreStagedMigrationManagerOptionKey
	PersistentStoreStagedMigrationManagerOptionKey string
	// PersistentStoreTimeoutOption is options key that specifies the connection timeout for Core Data stores.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreTimeoutOption
	PersistentStoreTimeoutOption string
	// PersistentStoreURLKey is a user info key to identify the store URL in persistent store remote change notifications.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreURLKey
	PersistentStoreURLKey string
	// ReadOnlyPersistentStoreOption is a flag that indicates whether a store is treated as read-only or not.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSReadOnlyPersistentStoreOption
	ReadOnlyPersistentStoreOption string
	// RefreshedObjectIDsKey is a user info key to identify refreshed object identifiers in notifications after saving a managed object context.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSRefreshedObjectIDsKey
	RefreshedObjectIDsKey string
	// RefreshedObjectsKey is a key for the set of objects that were refreshed but were not dirtied in the scope of this context.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSRefreshedObjectsKey
	RefreshedObjectsKey string
	// RemovedPersistentStoresKey is key for the array of stores that were removed.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSRemovedPersistentStoresKey
	RemovedPersistentStoresKey string
	// SQLiteAnalyzeOption is option key to run an analysis of the store data to optimize indices based on statistical information when the store is added to the coordinator.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSSQLiteAnalyzeOption
	SQLiteAnalyzeOption string
	// SQLiteErrorDomain is domain for SQLite errors.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSSQLiteErrorDomain
	SQLiteErrorDomain string
	// SQLiteManualVacuumOption is option key to rebuild the store file, forcing a database wide defragmentation when the store is added to the coordinator.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSSQLiteManualVacuumOption
	SQLiteManualVacuumOption string
	// SQLitePragmasOption is options key for a dictionary of SQLite pragma settings with pragma values indexed by pragma names as keys.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSSQLitePragmasOption
	SQLitePragmasOption string
	// SQLiteStoreType is the SQLite database store type.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSSQLiteStoreType
	SQLiteStoreType string
	// StoreModelVersionHashesKey is key to represent the version hash information for the model used to create the store.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSStoreModelVersionHashesKey
	StoreModelVersionHashesKey string
	// StoreModelVersionIdentifiersKey is key to represent the version identifiers for the model used to create the store.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSStoreModelVersionIdentifiersKey
	StoreModelVersionIdentifiersKey string
	// StoreTypeKey is a key that identifies the store type.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSStoreTypeKey
	StoreTypeKey string
	// StoreUUIDKey is a key that provides the store’s UUID.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSStoreUUIDKey
	StoreUUIDKey string
	// UUIDChangedPersistentStoresKey is key for an array containing the old and new stores.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSUUIDChangedPersistentStoresKey
	UUIDChangedPersistentStoresKey string
	// UpdatedObjectIDsKey is a user info key to identify updated object identifiers in notifications after saving a managed object context.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSUpdatedObjectIDsKey
	UpdatedObjectIDsKey string
	// UpdatedObjectsKey is a key for the set of objects that were updated.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSUpdatedObjectsKey
	UpdatedObjectsKey string
	// ValidateXMLStoreOption is a flag that indicates whether an XML file should be validated with the DTD while opening.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSValidateXMLStoreOption
	ValidateXMLStoreOption string
	// ValidationKeyErrorKey is the error key for the attribute that failed to validate.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSValidationKeyErrorKey
	ValidationKeyErrorKey string
	// ValidationObjectErrorKey is the error key for the object that failed to validate.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSValidationObjectErrorKey
	ValidationObjectErrorKey string
	// ValidationPredicateErrorKey is the error key for the predicate that failed to validate.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSValidationPredicateErrorKey
	ValidationPredicateErrorKey string
	// ValidationValueErrorKey is the error key for the value that failed to validate.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSValidationValueErrorKey
	ValidationValueErrorKey string
	// XMLStoreType is the XML store type.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSXMLStoreType
	XMLStoreType string
)

var (
	// See: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegateIndexDidUpdateNotification
	CoreDataCoreSpotlightDelegateIndexDidUpdateNotification foundation.NSNotificationName
	// PersistentCloudKitContainerEventChangedNotification is a notification that contains details about an event in a persistent CloudKit container.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/eventChangedNotification
	PersistentCloudKitContainerEventChangedNotification foundation.NSNotificationName
)

var (
	// CoreDataVersionNumber is the version of Core Data available in the current process.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSCoreDataVersionNumber
	CoreDataVersionNumber float64
)

var (
	// ErrorMergePolicy is the default merge policy for all managed object contexts.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSErrorMergePolicy
	ErrorMergePolicy objectivec.Object
	// MergeByPropertyObjectTrumpMergePolicy is a property-based merge policy that applies in-memory changes.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSMergeByPropertyObjectTrumpMergePolicy
	MergeByPropertyObjectTrumpMergePolicy objectivec.Object
	// MergeByPropertyStoreTrumpMergePolicy is a property-based merge policy that applies external changes.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSMergeByPropertyStoreTrumpMergePolicy
	MergeByPropertyStoreTrumpMergePolicy objectivec.Object
	// OverwriteMergePolicy is a merge policy that overwrites the entire stored object.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSOverwriteMergePolicy
	OverwriteMergePolicy objectivec.Object
	// RollbackMergePolicy is a merge policy that discards unsaved changes.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSRollbackMergePolicy
	RollbackMergePolicy objectivec.Object
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAddedPersistentStoresKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AddedPersistentStoresKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAffectedObjectsErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AffectedObjectsErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAffectedStoresErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AffectedStoresErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBinaryStoreInsecureDecodingCompatibilityOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				BinaryStoreInsecureDecodingCompatibilityOption = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBinaryStoreSecureDecodingClasses"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				BinaryStoreSecureDecodingClasses = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBinaryStoreType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				BinaryStoreType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCoreDataCoreSpotlightDelegateIndexDidUpdateNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CoreDataCoreSpotlightDelegateIndexDidUpdateNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCoreDataCoreSpotlightExporter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CoreDataCoreSpotlightExporter = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCoreDataVersionNumber"); err == nil && ptr != 0 {
		CoreDataVersionNumber = objc.ValueAt[float64](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDeletedObjectIDsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DeletedObjectIDsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDeletedObjectsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DeletedObjectsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDetailedErrorsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DetailedErrorsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSErrorMergePolicy"); err == nil && ptr != 0 {
		ErrorMergePolicy = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSIgnorePersistentStoreVersioningOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IgnorePersistentStoreVersioningOption = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInMemoryStoreType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InMemoryStoreType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInferMappingModelAutomaticallyOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InferMappingModelAutomaticallyOption = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInsertedObjectIDsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InsertedObjectIDsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInsertedObjectsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InsertedObjectsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInvalidatedAllObjectsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InvalidatedAllObjectsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInvalidatedObjectIDsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InvalidatedObjectIDsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInvalidatedObjectsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InvalidatedObjectsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSManagedObjectContextDidMergeChangesObjectIDsNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ManagedObjectContextDidMergeChangesObjectIDsNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSManagedObjectContextDidSaveNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ManagedObjectContextDidSaveNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSManagedObjectContextDidSaveObjectIDsNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ManagedObjectContextDidSaveObjectIDsNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSManagedObjectContextObjectsDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ManagedObjectContextObjectsDidChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSManagedObjectContextQueryGenerationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ManagedObjectContextQueryGenerationKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSManagedObjectContextWillSaveNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ManagedObjectContextWillSaveNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMergeByPropertyObjectTrumpMergePolicy"); err == nil && ptr != 0 {
		MergeByPropertyObjectTrumpMergePolicy = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMergeByPropertyStoreTrumpMergePolicy"); err == nil && ptr != 0 {
		MergeByPropertyStoreTrumpMergePolicy = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMigratePersistentStoresAutomaticallyOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MigratePersistentStoresAutomaticallyOption = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMigrationDestinationObjectKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MigrationDestinationObjectKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMigrationEntityMappingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MigrationEntityMappingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMigrationEntityPolicyKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MigrationEntityPolicyKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMigrationManagerKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MigrationManagerKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMigrationPropertyMappingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MigrationPropertyMappingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMigrationSourceObjectKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MigrationSourceObjectKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOverwriteMergePolicy"); err == nil && ptr != 0 {
		OverwriteMergePolicy = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersistentCloudKitContainerEventChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersistentCloudKitContainerEventChangedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersistentCloudKitContainerEventUserInfoKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersistentCloudKitContainerEventUserInfoKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersistentHistoryTokenKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersistentHistoryTokenKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersistentHistoryTrackingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersistentHistoryTrackingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersistentStoreConnectionPoolMaxSizeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersistentStoreConnectionPoolMaxSizeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersistentStoreCoordinatorStoresDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersistentStoreCoordinatorStoresDidChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersistentStoreCoordinatorStoresWillChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersistentStoreCoordinatorStoresWillChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersistentStoreCoordinatorWillRemoveStoreNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersistentStoreCoordinatorWillRemoveStoreNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersistentStoreDeferredLightweightMigrationOptionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersistentStoreDeferredLightweightMigrationOptionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersistentStoreForceDestroyOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersistentStoreForceDestroyOption = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersistentStoreModelVersionChecksumKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersistentStoreModelVersionChecksumKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersistentStoreOSCompatibility"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersistentStoreOSCompatibility = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersistentStoreRemoteChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersistentStoreRemoteChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersistentStoreRemoteChangeNotificationPostOptionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersistentStoreRemoteChangeNotificationPostOptionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersistentStoreSaveConflictsErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersistentStoreSaveConflictsErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersistentStoreStagedMigrationManagerOptionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersistentStoreStagedMigrationManagerOptionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersistentStoreTimeoutOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersistentStoreTimeoutOption = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersistentStoreURLKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersistentStoreURLKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSReadOnlyPersistentStoreOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ReadOnlyPersistentStoreOption = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRefreshedObjectIDsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RefreshedObjectIDsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRefreshedObjectsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RefreshedObjectsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRemovedPersistentStoresKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RemovedPersistentStoresKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRollbackMergePolicy"); err == nil && ptr != 0 {
		RollbackMergePolicy = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSQLiteAnalyzeOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SQLiteAnalyzeOption = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSQLiteErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SQLiteErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSQLiteManualVacuumOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SQLiteManualVacuumOption = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSQLitePragmasOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SQLitePragmasOption = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSQLiteStoreType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SQLiteStoreType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStoreModelVersionHashesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StoreModelVersionHashesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStoreModelVersionIdentifiersKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StoreModelVersionIdentifiersKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStoreTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StoreTypeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStoreUUIDKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StoreUUIDKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUUIDChangedPersistentStoresKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UUIDChangedPersistentStoresKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUpdatedObjectIDsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UpdatedObjectIDsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUpdatedObjectsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UpdatedObjectsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSValidateXMLStoreOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ValidateXMLStoreOption = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSValidationKeyErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ValidationKeyErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSValidationObjectErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ValidationObjectErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSValidationPredicateErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ValidationPredicateErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSValidationValueErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ValidationValueErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSXMLStoreType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				XMLStoreType = objc.GoString(cstr)
			}
		}
	}

}
