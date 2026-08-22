// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/CoreData/NSAttributeType
type NSAttributeType uint

const (
	// NSBinaryDataAttributeType: An attribute that stores binary data.
	NSBinaryDataAttributeType NSAttributeType = 1000
	// NSBooleanAttributeType: An attribute that stores a Boolean value.
	NSBooleanAttributeType NSAttributeType = 800
	// NSCompositeAttributeType: An attribute that derives its value by composing other attributes.
	NSCompositeAttributeType NSAttributeType = 2100
	// NSDateAttributeType: An attribute that stores a date.
	NSDateAttributeType NSAttributeType = 900
	// NSDecimalAttributeType: An attribute that stores a decimal value.
	NSDecimalAttributeType NSAttributeType = 400
	// NSDoubleAttributeType: An attribute that stores a double value.
	NSDoubleAttributeType NSAttributeType = 500
	// NSFloatAttributeType: An attribute that stores a float value.
	NSFloatAttributeType NSAttributeType = 600
	// NSInteger16AttributeType: An attribute that stores a 16-bit signed integer value.
	NSInteger16AttributeType NSAttributeType = 100
	// NSInteger32AttributeType: An attribute that stores a 32-bit signed integer value.
	NSInteger32AttributeType NSAttributeType = 200
	// NSInteger64AttributeType: An attribute that stores a 64-bit signed integer value.
	NSInteger64AttributeType NSAttributeType = 300
	// NSObjectIDAttributeType: An attribute that stores a managed object’s ID.
	NSObjectIDAttributeType NSAttributeType = 2000
	// NSStringAttributeType: An attribute that stores a string.
	NSStringAttributeType NSAttributeType = 700
	// NSTransformableAttributeType: An attribute that uses a value transformer to derive its value.
	NSTransformableAttributeType NSAttributeType = 1800
	// NSURIAttributeType: An attribute that stores a uniform resource identifier.
	NSURIAttributeType NSAttributeType = 1200
	// NSUUIDAttributeType: An attribute that stores a universally unique identifier.
	NSUUIDAttributeType NSAttributeType = 1100
	// NSUndefinedAttributeType: An attribute that doesn’t have an explicit type.
	NSUndefinedAttributeType NSAttributeType = 0
)

func (e NSAttributeType) String() string {
	switch e {
	case NSBinaryDataAttributeType:
		return "NSBinaryDataAttributeType"
	case NSBooleanAttributeType:
		return "NSBooleanAttributeType"
	case NSCompositeAttributeType:
		return "NSCompositeAttributeType"
	case NSDateAttributeType:
		return "NSDateAttributeType"
	case NSDecimalAttributeType:
		return "NSDecimalAttributeType"
	case NSDoubleAttributeType:
		return "NSDoubleAttributeType"
	case NSFloatAttributeType:
		return "NSFloatAttributeType"
	case NSInteger16AttributeType:
		return "NSInteger16AttributeType"
	case NSInteger32AttributeType:
		return "NSInteger32AttributeType"
	case NSInteger64AttributeType:
		return "NSInteger64AttributeType"
	case NSObjectIDAttributeType:
		return "NSObjectIDAttributeType"
	case NSStringAttributeType:
		return "NSStringAttributeType"
	case NSTransformableAttributeType:
		return "NSTransformableAttributeType"
	case NSURIAttributeType:
		return "NSURIAttributeType"
	case NSUUIDAttributeType:
		return "NSUUIDAttributeType"
	case NSUndefinedAttributeType:
		return "NSUndefinedAttributeType"
	default:
		return fmt.Sprintf("NSAttributeType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreData/NSBatchDeleteRequestResultType
type NSBatchDeleteRequestResultType uint

const (
	// NSBatchDeleteResultTypeCount: Returns the number of managed objects the request deletes.
	NSBatchDeleteResultTypeCount NSBatchDeleteRequestResultType = 0x2
	// NSBatchDeleteResultTypeObjectIDs: Returns an array of the deleted managed objects’ identifiers.
	NSBatchDeleteResultTypeObjectIDs NSBatchDeleteRequestResultType = 0x1
	// NSBatchDeleteResultTypeStatusOnly: Returns a Boolean value that indicates if the request succeeds.
	NSBatchDeleteResultTypeStatusOnly NSBatchDeleteRequestResultType = 0
)

func (e NSBatchDeleteRequestResultType) String() string {
	switch e {
	case NSBatchDeleteResultTypeCount:
		return "NSBatchDeleteResultTypeCount"
	case NSBatchDeleteResultTypeObjectIDs:
		return "NSBatchDeleteResultTypeObjectIDs"
	case NSBatchDeleteResultTypeStatusOnly:
		return "NSBatchDeleteResultTypeStatusOnly"
	default:
		return fmt.Sprintf("NSBatchDeleteRequestResultType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequestResultType
type NSBatchInsertRequestResultType uint

const (
	// NSBatchInsertRequestResultTypeCount: A value that indicates that the return type is the number of inserted rows.
	NSBatchInsertRequestResultTypeCount NSBatchInsertRequestResultType = 0x2
	// NSBatchInsertRequestResultTypeObjectIDs: A value that indicates the return type is an array of object IDs that corresponds to the inserted rows.
	NSBatchInsertRequestResultTypeObjectIDs NSBatchInsertRequestResultType = 0x1
	// NSBatchInsertRequestResultTypeStatusOnly: A value that indicates that the return type is a Boolean value representing whether the batch-insertion request succeeded.
	NSBatchInsertRequestResultTypeStatusOnly NSBatchInsertRequestResultType = 0
)

func (e NSBatchInsertRequestResultType) String() string {
	switch e {
	case NSBatchInsertRequestResultTypeCount:
		return "NSBatchInsertRequestResultTypeCount"
	case NSBatchInsertRequestResultTypeObjectIDs:
		return "NSBatchInsertRequestResultTypeObjectIDs"
	case NSBatchInsertRequestResultTypeStatusOnly:
		return "NSBatchInsertRequestResultTypeStatusOnly"
	default:
		return fmt.Sprintf("NSBatchInsertRequestResultType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequestResultType
type NSBatchUpdateRequestResultType uint

const (
	// NSStatusOnlyResultType: A value that indicates the return type is a Boolean value representing whether the batch-update request succeeds.
	NSStatusOnlyResultType NSBatchUpdateRequestResultType = 0
	// NSUpdatedObjectIDsResultType: A value that indicates the return type is an array of object IDs that corresponds to the updated rows.
	NSUpdatedObjectIDsResultType NSBatchUpdateRequestResultType = 0x1
	// NSUpdatedObjectsCountResultType: A value that indicates the return type is the number of updated rows.
	NSUpdatedObjectsCountResultType NSBatchUpdateRequestResultType = 0x2
)

func (e NSBatchUpdateRequestResultType) String() string {
	switch e {
	case NSStatusOnlyResultType:
		return "NSStatusOnlyResultType"
	case NSUpdatedObjectIDsResultType:
		return "NSUpdatedObjectIDsResultType"
	case NSUpdatedObjectsCountResultType:
		return "NSUpdatedObjectsCountResultType"
	default:
		return fmt.Sprintf("NSBatchUpdateRequestResultType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreData/NSDeleteRule
type NSDeleteRule uint

const (
	// NSCascadeDeleteRule: A rule that deletes the referenced managed objects.
	NSCascadeDeleteRule NSDeleteRule = 2
	// NSDenyDeleteRule: A rule that prevents the deletion of the owning managed object if the relationship has references to other objects.
	NSDenyDeleteRule NSDeleteRule = 3
	// NSNoActionDeleteRule: A rule that prevents modification of the referenced managed objects.
	NSNoActionDeleteRule NSDeleteRule = 0
	// NSNullifyDeleteRule: A rule that nullifies the inverse relationship of the referenced managed objects.
	NSNullifyDeleteRule NSDeleteRule = 1
)

func (e NSDeleteRule) String() string {
	switch e {
	case NSCascadeDeleteRule:
		return "NSCascadeDeleteRule"
	case NSDenyDeleteRule:
		return "NSDenyDeleteRule"
	case NSNoActionDeleteRule:
		return "NSNoActionDeleteRule"
	case NSNullifyDeleteRule:
		return "NSNullifyDeleteRule"
	default:
		return fmt.Sprintf("NSDeleteRule(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreData/NSEntityMappingType
type NSEntityMappingType uint

const (
	// NSAddEntityMappingType: Specifies that this is a new entity in the destination model.
	NSAddEntityMappingType NSEntityMappingType = 0x2
	// NSCopyEntityMappingType: Specifies that source instances are migrated as-is.
	NSCopyEntityMappingType NSEntityMappingType = 0x4
	// NSCustomEntityMappingType: Specifies a custom mapping.
	NSCustomEntityMappingType NSEntityMappingType = 0x1
	// NSRemoveEntityMappingType: Specifies that this entity is not present in the destination model.
	NSRemoveEntityMappingType NSEntityMappingType = 0x3
	// NSTransformEntityMappingType: Specifies that entity exists in source and destination and is mapped.
	NSTransformEntityMappingType NSEntityMappingType = 0x5
	// NSUndefinedEntityMappingType: Specifies that the developer handles destination instance creation.
	NSUndefinedEntityMappingType NSEntityMappingType = 0
)

func (e NSEntityMappingType) String() string {
	switch e {
	case NSAddEntityMappingType:
		return "NSAddEntityMappingType"
	case NSCopyEntityMappingType:
		return "NSCopyEntityMappingType"
	case NSCustomEntityMappingType:
		return "NSCustomEntityMappingType"
	case NSRemoveEntityMappingType:
		return "NSRemoveEntityMappingType"
	case NSTransformEntityMappingType:
		return "NSTransformEntityMappingType"
	case NSUndefinedEntityMappingType:
		return "NSUndefinedEntityMappingType"
	default:
		return fmt.Sprintf("NSEntityMappingType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreData/NSFetchIndexElementType
type NSFetchIndexElementType uint

const (
	// NSFetchIndexElementTypeBinary: A binary index type.
	NSFetchIndexElementTypeBinary NSFetchIndexElementType = 0
	// NSFetchIndexElementTypeRTree: An R-tree index type.
	NSFetchIndexElementTypeRTree NSFetchIndexElementType = 1
)

func (e NSFetchIndexElementType) String() string {
	switch e {
	case NSFetchIndexElementTypeBinary:
		return "NSFetchIndexElementTypeBinary"
	case NSFetchIndexElementTypeRTree:
		return "NSFetchIndexElementTypeRTree"
	default:
		return fmt.Sprintf("NSFetchIndexElementType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreData/NSFetchRequestResultType
type NSFetchRequestResultType uint

const (
	// NSCountResultType: The request returns the count of the objects that match the request.
	NSCountResultType NSFetchRequestResultType = 0x4
	// NSDictionaryResultType: The request returns dictionaries.
	NSDictionaryResultType NSFetchRequestResultType = 0x2
	// NSManagedObjectIDResultType: The request returns managed object IDs.
	NSManagedObjectIDResultType NSFetchRequestResultType = 0x1
	// NSManagedObjectResultType: The request returns managed objects.
	NSManagedObjectResultType NSFetchRequestResultType = 0
)

func (e NSFetchRequestResultType) String() string {
	switch e {
	case NSCountResultType:
		return "NSCountResultType"
	case NSDictionaryResultType:
		return "NSDictionaryResultType"
	case NSManagedObjectIDResultType:
		return "NSManagedObjectIDResultType"
	case NSManagedObjectResultType:
		return "NSManagedObjectResultType"
	default:
		return fmt.Sprintf("NSFetchRequestResultType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsChangeType
type NSFetchedResultsChangeType uint

const (
	// NSFetchedResultsChangeDelete: Specifies that an object was deleted.
	NSFetchedResultsChangeDelete NSFetchedResultsChangeType = 2
	// NSFetchedResultsChangeInsert: Specifies that an object was inserted.
	NSFetchedResultsChangeInsert NSFetchedResultsChangeType = 1
	// NSFetchedResultsChangeMove: Specifies that an object was moved.
	NSFetchedResultsChangeMove NSFetchedResultsChangeType = 3
	// NSFetchedResultsChangeUpdate: Specifies that an object was changed.
	NSFetchedResultsChangeUpdate NSFetchedResultsChangeType = 4
)

func (e NSFetchedResultsChangeType) String() string {
	switch e {
	case NSFetchedResultsChangeDelete:
		return "NSFetchedResultsChangeDelete"
	case NSFetchedResultsChangeInsert:
		return "NSFetchedResultsChangeInsert"
	case NSFetchedResultsChangeMove:
		return "NSFetchedResultsChangeMove"
	case NSFetchedResultsChangeUpdate:
		return "NSFetchedResultsChangeUpdate"
	default:
		return fmt.Sprintf("NSFetchedResultsChangeType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectContextConcurrencyType
type NSManagedObjectContextConcurrencyType uint

const (
	// NSConfinementConcurrencyType: Specifies that the context will use the thread confinement pattern.
	NSConfinementConcurrencyType NSManagedObjectContextConcurrencyType = 0
	// NSMainQueueConcurrencyType: Specifies that the context will be associated with the main queue.
	NSMainQueueConcurrencyType NSManagedObjectContextConcurrencyType = 0x2
	// NSPrivateQueueConcurrencyType: Specifies that the context will be associated with a private dispatch queue.
	NSPrivateQueueConcurrencyType NSManagedObjectContextConcurrencyType = 0x1
)

func (e NSManagedObjectContextConcurrencyType) String() string {
	switch e {
	case NSConfinementConcurrencyType:
		return "NSConfinementConcurrencyType"
	case NSMainQueueConcurrencyType:
		return "NSMainQueueConcurrencyType"
	case NSPrivateQueueConcurrencyType:
		return "NSPrivateQueueConcurrencyType"
	default:
		return fmt.Sprintf("NSManagedObjectContextConcurrencyType(%d)", e)
	}
}

type NSManagedObjectValidationErrorConstants int

const (
	// NSCoreDataError: An error code that indicates a nonspecific Core Data error.
	NSCoreDataError NSManagedObjectValidationErrorConstants = 134060
	// NSEntityMigrationPolicyError: An error code that indicates a migration failure during processing of an entity migration policy.
	NSEntityMigrationPolicyError NSManagedObjectValidationErrorConstants = 134170
	// NSExternalRecordImportError: Error code to denote a general error encountered while importing external records.
	NSExternalRecordImportError NSManagedObjectValidationErrorConstants = 134200
	// NSInferredMappingModelError: Error code to denote a problem with the creation of an inferred mapping model.
	NSInferredMappingModelError NSManagedObjectValidationErrorConstants = 134190
	// NSManagedObjectConstraintMergeError: Error code to denote a problem with the merging of instances of a managed object.
	NSManagedObjectConstraintMergeError NSManagedObjectValidationErrorConstants = 133021
	// NSManagedObjectConstraintValidationError: Error code to denote a problem with the validation of a managed object.
	NSManagedObjectConstraintValidationError NSManagedObjectValidationErrorConstants = 1551
	// NSManagedObjectContextLockingError: Error code to denote an inability to acquire a lock in a managed object context.
	NSManagedObjectContextLockingError NSManagedObjectValidationErrorConstants = 132000
	// NSManagedObjectExternalRelationshipError: Error code to denote that an object being saved has a relationship containing an object from another store.
	NSManagedObjectExternalRelationshipError NSManagedObjectValidationErrorConstants = 133010
	// NSManagedObjectMergeError: Error code to denote that a merge policy failed—Core Data is unable to complete merging.
	NSManagedObjectMergeError NSManagedObjectValidationErrorConstants = 133020
	// NSManagedObjectModelReferenceNotFoundError: An error code that indicates Core Data isn’t able to find or instantiate the referenced object model.
	NSManagedObjectModelReferenceNotFoundError NSManagedObjectValidationErrorConstants = 134504
	// NSManagedObjectReferentialIntegrityError: Error code to denote an attempt to fire a fault pointing to an object that does not exist.
	NSManagedObjectReferentialIntegrityError NSManagedObjectValidationErrorConstants = 133000
	// NSManagedObjectValidationError: Error code to denote a generic validation error.
	NSManagedObjectValidationError NSManagedObjectValidationErrorConstants = 1550
	// NSMigrationCancelledError: Error code to denote that migration failed due to manual cancellation.
	NSMigrationCancelledError NSManagedObjectValidationErrorConstants = 134120
	// NSMigrationConstraintViolationError: Error code to denote a problem with the validation of a managed object during a migration.
	NSMigrationConstraintViolationError NSManagedObjectValidationErrorConstants = 134111
	// NSMigrationError: Error code to denote a general migration error.
	NSMigrationError NSManagedObjectValidationErrorConstants = 134110
	// NSMigrationManagerDestinationStoreError: Error code to denote that migration failed due to a problem with the destination data store.
	NSMigrationManagerDestinationStoreError NSManagedObjectValidationErrorConstants = 134160
	// NSMigrationManagerSourceStoreError: Error code to denote that migration failed due to a problem with the source data store.
	NSMigrationManagerSourceStoreError NSManagedObjectValidationErrorConstants = 134150
	// NSMigrationMissingMappingModelError: Error code to denote that migration failed due to a missing mapping model.
	NSMigrationMissingMappingModelError NSManagedObjectValidationErrorConstants = 134140
	// NSMigrationMissingSourceModelError: Error code to denote that migration failed due to a missing source data model.
	NSMigrationMissingSourceModelError NSManagedObjectValidationErrorConstants = 134130
	// NSPersistentHistoryTokenExpiredError: Error code to denote that the persistent history token has expired.
	NSPersistentHistoryTokenExpiredError NSManagedObjectValidationErrorConstants = 134301
	// NSPersistentStoreCoordinatorLockingError: Error code to denote an inability to acquire a lock in a persistent store.
	NSPersistentStoreCoordinatorLockingError NSManagedObjectValidationErrorConstants = 132010
	// NSPersistentStoreIncompatibleSchemaError: Error code to denote that a persistent store returned an error for a save operation.
	NSPersistentStoreIncompatibleSchemaError NSManagedObjectValidationErrorConstants = 134020
	// NSPersistentStoreIncompatibleVersionHashError: Error code to denote that entity version hashes in the store are incompatible with the current managed object model.
	NSPersistentStoreIncompatibleVersionHashError NSManagedObjectValidationErrorConstants = 134100
	// NSPersistentStoreIncompleteSaveError: Error code to denote that one or more of the stores returned an error during a save operations.
	NSPersistentStoreIncompleteSaveError NSManagedObjectValidationErrorConstants = 134040
	// NSPersistentStoreInvalidTypeError: Error code to denote an unknown persistent store type/format/version.
	NSPersistentStoreInvalidTypeError NSManagedObjectValidationErrorConstants = 134000
	// NSPersistentStoreOpenError: Error code to denote an error occurred while attempting to open a persistent store.
	NSPersistentStoreOpenError NSManagedObjectValidationErrorConstants = 134080
	// NSPersistentStoreOperationError: Error code to denote that a persistent store operation failed.
	NSPersistentStoreOperationError NSManagedObjectValidationErrorConstants = 134070
	// NSPersistentStoreSaveConflictsError: Error code to denote that an unresolved merge conflict was encountered during a save.
	NSPersistentStoreSaveConflictsError NSManagedObjectValidationErrorConstants = 134050
	// NSPersistentStoreSaveError: Error code to denote that a persistent store returned an error for a save operation.
	NSPersistentStoreSaveError NSManagedObjectValidationErrorConstants = 134030
	// NSPersistentStoreTimeoutError: Error code to denote that Core Data failed to connect to a persistent store within the time specified by [NSPersistentStoreTimeoutOption].
	NSPersistentStoreTimeoutError NSManagedObjectValidationErrorConstants = 134090
	// NSPersistentStoreTypeMismatchError: Error code returned by a persistent store coordinator if a store is accessed that does not match the specified type.
	NSPersistentStoreTypeMismatchError NSManagedObjectValidationErrorConstants = 134010
	// NSPersistentStoreUnsupportedRequestTypeError: Error code to denote that an [NSPersistentStore] subclass was passed a request (an instance of NSPersistentStoreRequest) that it did not understand.
	NSPersistentStoreUnsupportedRequestTypeError NSManagedObjectValidationErrorConstants = 134091
	// NSSQLiteError: Error code to denote a general SQLite error.
	NSSQLiteError NSManagedObjectValidationErrorConstants = 134180
	// NSStagedMigrationBackwardMigrationError: An error code that indicates a failed migration because of an attempt to migrate backward.
	NSStagedMigrationBackwardMigrationError NSManagedObjectValidationErrorConstants = 134506
	// NSStagedMigrationFrameworkVersionMismatchError: An error code that indicates a failed migration because the persistent store’s metadata doesn’t support staged lightweight migrations.
	NSStagedMigrationFrameworkVersionMismatchError NSManagedObjectValidationErrorConstants = 134505
	// NSValidationDateTooLateError: Error code to denote some date value is too late.
	NSValidationDateTooLateError NSManagedObjectValidationErrorConstants = 1630
	// NSValidationDateTooSoonError: Error code to denote some date value is too soon.
	NSValidationDateTooSoonError NSManagedObjectValidationErrorConstants = 1640
	// NSValidationInvalidDateError: Error code to denote some date value fails to match date pattern.
	NSValidationInvalidDateError NSManagedObjectValidationErrorConstants = 1650
	// NSValidationInvalidURIError: Error code to denote a problem with the validation of a URI property.
	NSValidationInvalidURIError NSManagedObjectValidationErrorConstants = 1690
	// NSValidationMissingMandatoryPropertyError: Error code for a non-optional property with a nil value.
	NSValidationMissingMandatoryPropertyError NSManagedObjectValidationErrorConstants = 1570
	// NSValidationMultipleErrorsError: Error code to denote an error containing multiple validation errors.
	NSValidationMultipleErrorsError NSManagedObjectValidationErrorConstants = 1560
	// NSValidationNumberTooLargeError: Error code to denote some numerical value is too large.
	NSValidationNumberTooLargeError NSManagedObjectValidationErrorConstants = 1610
	// NSValidationNumberTooSmallError: Error code to denote some numerical value is too small.
	NSValidationNumberTooSmallError NSManagedObjectValidationErrorConstants = 1620
	// NSValidationRelationshipDeniedDeleteError: Error code to denote some relationship with delete rule [NSDeleteRuleDeny] is non-empty.
	NSValidationRelationshipDeniedDeleteError NSManagedObjectValidationErrorConstants = 1600
	// NSValidationRelationshipExceedsMaximumCountError: Error code to denote a bounded to-many relationship with too many destination objects.
	NSValidationRelationshipExceedsMaximumCountError NSManagedObjectValidationErrorConstants = 1590
	// NSValidationRelationshipLacksMinimumCountError: Error code to denote a to-many relationship with too few destination objects.
	NSValidationRelationshipLacksMinimumCountError NSManagedObjectValidationErrorConstants = 1580
	// NSValidationStringPatternMatchingError: Error code to denote some string value fails to match some pattern.
	NSValidationStringPatternMatchingError NSManagedObjectValidationErrorConstants = 1680
	// NSValidationStringTooLongError: Error code to denote some string value is too long.
	NSValidationStringTooLongError NSManagedObjectValidationErrorConstants = 1660
	// NSValidationStringTooShortError: Error code to denote some string value is too short.
	NSValidationStringTooShortError NSManagedObjectValidationErrorConstants = 1670
)

func (e NSManagedObjectValidationErrorConstants) String() string {
	switch e {
	case NSCoreDataError:
		return "NSCoreDataError"
	case NSEntityMigrationPolicyError:
		return "NSEntityMigrationPolicyError"
	case NSExternalRecordImportError:
		return "NSExternalRecordImportError"
	case NSInferredMappingModelError:
		return "NSInferredMappingModelError"
	case NSManagedObjectConstraintMergeError:
		return "NSManagedObjectConstraintMergeError"
	case NSManagedObjectConstraintValidationError:
		return "NSManagedObjectConstraintValidationError"
	case NSManagedObjectContextLockingError:
		return "NSManagedObjectContextLockingError"
	case NSManagedObjectExternalRelationshipError:
		return "NSManagedObjectExternalRelationshipError"
	case NSManagedObjectMergeError:
		return "NSManagedObjectMergeError"
	case NSManagedObjectModelReferenceNotFoundError:
		return "NSManagedObjectModelReferenceNotFoundError"
	case NSManagedObjectReferentialIntegrityError:
		return "NSManagedObjectReferentialIntegrityError"
	case NSManagedObjectValidationError:
		return "NSManagedObjectValidationError"
	case NSMigrationCancelledError:
		return "NSMigrationCancelledError"
	case NSMigrationConstraintViolationError:
		return "NSMigrationConstraintViolationError"
	case NSMigrationError:
		return "NSMigrationError"
	case NSMigrationManagerDestinationStoreError:
		return "NSMigrationManagerDestinationStoreError"
	case NSMigrationManagerSourceStoreError:
		return "NSMigrationManagerSourceStoreError"
	case NSMigrationMissingMappingModelError:
		return "NSMigrationMissingMappingModelError"
	case NSMigrationMissingSourceModelError:
		return "NSMigrationMissingSourceModelError"
	case NSPersistentHistoryTokenExpiredError:
		return "NSPersistentHistoryTokenExpiredError"
	case NSPersistentStoreCoordinatorLockingError:
		return "NSPersistentStoreCoordinatorLockingError"
	case NSPersistentStoreIncompatibleSchemaError:
		return "NSPersistentStoreIncompatibleSchemaError"
	case NSPersistentStoreIncompatibleVersionHashError:
		return "NSPersistentStoreIncompatibleVersionHashError"
	case NSPersistentStoreIncompleteSaveError:
		return "NSPersistentStoreIncompleteSaveError"
	case NSPersistentStoreInvalidTypeError:
		return "NSPersistentStoreInvalidTypeError"
	case NSPersistentStoreOpenError:
		return "NSPersistentStoreOpenError"
	case NSPersistentStoreOperationError:
		return "NSPersistentStoreOperationError"
	case NSPersistentStoreSaveConflictsError:
		return "NSPersistentStoreSaveConflictsError"
	case NSPersistentStoreSaveError:
		return "NSPersistentStoreSaveError"
	case NSPersistentStoreTimeoutError:
		return "NSPersistentStoreTimeoutError"
	case NSPersistentStoreTypeMismatchError:
		return "NSPersistentStoreTypeMismatchError"
	case NSPersistentStoreUnsupportedRequestTypeError:
		return "NSPersistentStoreUnsupportedRequestTypeError"
	case NSSQLiteError:
		return "NSSQLiteError"
	case NSStagedMigrationBackwardMigrationError:
		return "NSStagedMigrationBackwardMigrationError"
	case NSStagedMigrationFrameworkVersionMismatchError:
		return "NSStagedMigrationFrameworkVersionMismatchError"
	case NSValidationDateTooLateError:
		return "NSValidationDateTooLateError"
	case NSValidationDateTooSoonError:
		return "NSValidationDateTooSoonError"
	case NSValidationInvalidDateError:
		return "NSValidationInvalidDateError"
	case NSValidationInvalidURIError:
		return "NSValidationInvalidURIError"
	case NSValidationMissingMandatoryPropertyError:
		return "NSValidationMissingMandatoryPropertyError"
	case NSValidationMultipleErrorsError:
		return "NSValidationMultipleErrorsError"
	case NSValidationNumberTooLargeError:
		return "NSValidationNumberTooLargeError"
	case NSValidationNumberTooSmallError:
		return "NSValidationNumberTooSmallError"
	case NSValidationRelationshipDeniedDeleteError:
		return "NSValidationRelationshipDeniedDeleteError"
	case NSValidationRelationshipExceedsMaximumCountError:
		return "NSValidationRelationshipExceedsMaximumCountError"
	case NSValidationRelationshipLacksMinimumCountError:
		return "NSValidationRelationshipLacksMinimumCountError"
	case NSValidationStringPatternMatchingError:
		return "NSValidationStringPatternMatchingError"
	case NSValidationStringTooLongError:
		return "NSValidationStringTooLongError"
	case NSValidationStringTooShortError:
		return "NSValidationStringTooShortError"
	default:
		return fmt.Sprintf("NSManagedObjectValidationErrorConstants(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreData/NSMergePolicyType
type NSMergePolicyType uint

const (
	// NSErrorMergePolicyType: The default merge policy for all managed object contexts.
	NSErrorMergePolicyType NSMergePolicyType = 0
	// NSMergeByPropertyObjectTrumpMergePolicyType: A property-based merge policy that applies in-memory changes.
	NSMergeByPropertyObjectTrumpMergePolicyType NSMergePolicyType = 0x2
	// NSMergeByPropertyStoreTrumpMergePolicyType: A property-based merge policy that applies external changes.
	NSMergeByPropertyStoreTrumpMergePolicyType NSMergePolicyType = 0x1
	// NSOverwriteMergePolicyType: A merge policy type that overwrites the entire stored object.
	NSOverwriteMergePolicyType NSMergePolicyType = 0x3
	// NSRollbackMergePolicyType: A merge policy that discards unsaved changes.
	NSRollbackMergePolicyType NSMergePolicyType = 0x4
)

func (e NSMergePolicyType) String() string {
	switch e {
	case NSErrorMergePolicyType:
		return "NSErrorMergePolicyType"
	case NSMergeByPropertyObjectTrumpMergePolicyType:
		return "NSMergeByPropertyObjectTrumpMergePolicyType"
	case NSMergeByPropertyStoreTrumpMergePolicyType:
		return "NSMergeByPropertyStoreTrumpMergePolicyType"
	case NSOverwriteMergePolicyType:
		return "NSOverwriteMergePolicyType"
	case NSRollbackMergePolicyType:
		return "NSRollbackMergePolicyType"
	default:
		return fmt.Sprintf("NSMergePolicyType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventResult/ResultType-swift.enum
type NSPersistentCloudKitContainerEventResultType int

const (
	// NSPersistentCloudKitContainerEventResultTypeCountEvents: The number of CloudKit container events that match the event request.
	NSPersistentCloudKitContainerEventResultTypeCountEvents NSPersistentCloudKitContainerEventResultType = 1
	// NSPersistentCloudKitContainerEventResultTypeEvents: The persistent CloudKit container events that match the event request.
	NSPersistentCloudKitContainerEventResultTypeEvents NSPersistentCloudKitContainerEventResultType = 0
)

func (e NSPersistentCloudKitContainerEventResultType) String() string {
	switch e {
	case NSPersistentCloudKitContainerEventResultTypeCountEvents:
		return "NSPersistentCloudKitContainerEventResultTypeCountEvents"
	case NSPersistentCloudKitContainerEventResultTypeEvents:
		return "NSPersistentCloudKitContainerEventResultTypeEvents"
	default:
		return fmt.Sprintf("NSPersistentCloudKitContainerEventResultType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/EventType
type NSPersistentCloudKitContainerEventType int

const (
	// NSPersistentCloudKitContainerEventTypeExport: An event the persistent CloudKit container generates when exporting managed objects from a store.
	NSPersistentCloudKitContainerEventTypeExport NSPersistentCloudKitContainerEventType = 2
	// NSPersistentCloudKitContainerEventTypeImport: An event the persistent CloudKit container generates when importing records into a store.
	NSPersistentCloudKitContainerEventTypeImport NSPersistentCloudKitContainerEventType = 1
	// NSPersistentCloudKitContainerEventTypeSetup: An event the persistent CloudKit container generates when setting up a store.
	NSPersistentCloudKitContainerEventTypeSetup NSPersistentCloudKitContainerEventType = 0
)

func (e NSPersistentCloudKitContainerEventType) String() string {
	switch e {
	case NSPersistentCloudKitContainerEventTypeExport:
		return "NSPersistentCloudKitContainerEventTypeExport"
	case NSPersistentCloudKitContainerEventTypeImport:
		return "NSPersistentCloudKitContainerEventTypeImport"
	case NSPersistentCloudKitContainerEventTypeSetup:
		return "NSPersistentCloudKitContainerEventTypeSetup"
	default:
		return fmt.Sprintf("NSPersistentCloudKitContainerEventType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerSchemaInitializationOptions
type NSPersistentCloudKitContainerSchemaInitializationOptions uint

const (
	// NSPersistentCloudKitContainerSchemaInitializationOptionsDryRun: A flag that indicates the container validates the model and generates the records, but doesn’t upload them to CloudKit.
	NSPersistentCloudKitContainerSchemaInitializationOptionsDryRun NSPersistentCloudKitContainerSchemaInitializationOptions = 2
	// NSPersistentCloudKitContainerSchemaInitializationOptionsNone: Indicates there are no specified schema options.
	NSPersistentCloudKitContainerSchemaInitializationOptionsNone NSPersistentCloudKitContainerSchemaInitializationOptions = 0
	// NSPersistentCloudKitContainerSchemaInitializationOptionsPrintSchema: Prints the generated records to the console.
	NSPersistentCloudKitContainerSchemaInitializationOptionsPrintSchema NSPersistentCloudKitContainerSchemaInitializationOptions = 4
)

func (e NSPersistentCloudKitContainerSchemaInitializationOptions) String() string {
	switch e {
	case NSPersistentCloudKitContainerSchemaInitializationOptionsDryRun:
		return "NSPersistentCloudKitContainerSchemaInitializationOptionsDryRun"
	case NSPersistentCloudKitContainerSchemaInitializationOptionsNone:
		return "NSPersistentCloudKitContainerSchemaInitializationOptionsNone"
	case NSPersistentCloudKitContainerSchemaInitializationOptionsPrintSchema:
		return "NSPersistentCloudKitContainerSchemaInitializationOptionsPrintSchema"
	default:
		return fmt.Sprintf("NSPersistentCloudKitContainerSchemaInitializationOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeType
type NSPersistentHistoryChangeType int

const (
	// NSPersistentHistoryChangeTypeDelete: The deletion of a managed object from the persistent store.
	NSPersistentHistoryChangeTypeDelete NSPersistentHistoryChangeType = 2
	// NSPersistentHistoryChangeTypeInsert: The insertion of a managed object into the persistent store.
	NSPersistentHistoryChangeTypeInsert NSPersistentHistoryChangeType = 0
	// NSPersistentHistoryChangeTypeUpdate: An update to a managed object’s properties in the persistent store.
	NSPersistentHistoryChangeTypeUpdate NSPersistentHistoryChangeType = 1
)

func (e NSPersistentHistoryChangeType) String() string {
	switch e {
	case NSPersistentHistoryChangeTypeDelete:
		return "NSPersistentHistoryChangeTypeDelete"
	case NSPersistentHistoryChangeTypeInsert:
		return "NSPersistentHistoryChangeTypeInsert"
	case NSPersistentHistoryChangeTypeUpdate:
		return "NSPersistentHistoryChangeTypeUpdate"
	default:
		return fmt.Sprintf("NSPersistentHistoryChangeType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryResultType
type NSPersistentHistoryResultType int

const (
	// NSPersistentHistoryResultTypeChangesOnly: The persistent history changes since the requested date, token, or transaction.
	NSPersistentHistoryResultTypeChangesOnly NSPersistentHistoryResultType = 0x4
	// NSPersistentHistoryResultTypeCount: The number of persistent history changes since the requested date, token, or transaction.
	NSPersistentHistoryResultTypeCount NSPersistentHistoryResultType = 0x2
	// NSPersistentHistoryResultTypeObjectIDs: The identifiers of managed objects changed since the requested date, token, or transaction.
	NSPersistentHistoryResultTypeObjectIDs NSPersistentHistoryResultType = 0x1
	// NSPersistentHistoryResultTypeStatusOnly: The status of the persistent history change request.
	NSPersistentHistoryResultTypeStatusOnly NSPersistentHistoryResultType = 0
	// NSPersistentHistoryResultTypeTransactionsAndChanges: The persistent history transactions and changes since the requested date, token, or transaction.
	NSPersistentHistoryResultTypeTransactionsAndChanges NSPersistentHistoryResultType = 0x5
	// NSPersistentHistoryResultTypeTransactionsOnly: The persistent history transactions since the requested date, token, or transaction.
	NSPersistentHistoryResultTypeTransactionsOnly NSPersistentHistoryResultType = 0x3
)

func (e NSPersistentHistoryResultType) String() string {
	switch e {
	case NSPersistentHistoryResultTypeChangesOnly:
		return "NSPersistentHistoryResultTypeChangesOnly"
	case NSPersistentHistoryResultTypeCount:
		return "NSPersistentHistoryResultTypeCount"
	case NSPersistentHistoryResultTypeObjectIDs:
		return "NSPersistentHistoryResultTypeObjectIDs"
	case NSPersistentHistoryResultTypeStatusOnly:
		return "NSPersistentHistoryResultTypeStatusOnly"
	case NSPersistentHistoryResultTypeTransactionsAndChanges:
		return "NSPersistentHistoryResultTypeTransactionsAndChanges"
	case NSPersistentHistoryResultTypeTransactionsOnly:
		return "NSPersistentHistoryResultTypeTransactionsOnly"
	default:
		return fmt.Sprintf("NSPersistentHistoryResultType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreRequestType
type NSPersistentStoreRequestType uint

const (
	// NSBatchDeleteRequestType: A request that deletes data for multiple managed objects from a persistent store.
	NSBatchDeleteRequestType NSPersistentStoreRequestType = 7
	// NSBatchInsertRequestType: A request that inserts data into a persistent store using a batch of managed objects or dictionaries.
	NSBatchInsertRequestType NSPersistentStoreRequestType = 5
	// NSBatchUpdateRequestType: A request that updates data for multiple managed objects in a persistent store.
	NSBatchUpdateRequestType NSPersistentStoreRequestType = 6
	// NSFetchRequestType: Specifies that the request returns managed objects.
	NSFetchRequestType NSPersistentStoreRequestType = 1
	// NSSaveRequestType: Specifies that the request saves managed objects.
	NSSaveRequestType NSPersistentStoreRequestType = 2
)

func (e NSPersistentStoreRequestType) String() string {
	switch e {
	case NSBatchDeleteRequestType:
		return "NSBatchDeleteRequestType"
	case NSBatchInsertRequestType:
		return "NSBatchInsertRequestType"
	case NSBatchUpdateRequestType:
		return "NSBatchUpdateRequestType"
	case NSFetchRequestType:
		return "NSFetchRequestType"
	case NSSaveRequestType:
		return "NSSaveRequestType"
	default:
		return fmt.Sprintf("NSPersistentStoreRequestType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreUbiquitousTransitionType
type NSPersistentStoreUbiquitousTransitionType uint

const (
	// NSPersistentStoreUbiquitousTransitionTypeAccountAdded: This value indicates that a new iCloud account is available, and the persistent store in use will or did transition to the new account.
	NSPersistentStoreUbiquitousTransitionTypeAccountAdded NSPersistentStoreUbiquitousTransitionType = 1
	// NSPersistentStoreUbiquitousTransitionTypeAccountRemoved: This value indicates that no iCloud account is available, and the persistent store in use will or did transition to the “local” store.
	NSPersistentStoreUbiquitousTransitionTypeAccountRemoved NSPersistentStoreUbiquitousTransitionType = 2
	// NSPersistentStoreUbiquitousTransitionTypeContentRemoved: This value indicates that the user has wiped the contents of the iCloud account, usually using Delete All from Documents & Data in Settings.
	NSPersistentStoreUbiquitousTransitionTypeContentRemoved NSPersistentStoreUbiquitousTransitionType = 3
	// NSPersistentStoreUbiquitousTransitionTypeInitialImportCompleted: This value indicates that the Core Data integration has finished building a store file that is consistent with the contents of the iCloud account, and is ready to replace the fallback store with that file.
	NSPersistentStoreUbiquitousTransitionTypeInitialImportCompleted NSPersistentStoreUbiquitousTransitionType = 4
)

func (e NSPersistentStoreUbiquitousTransitionType) String() string {
	switch e {
	case NSPersistentStoreUbiquitousTransitionTypeAccountAdded:
		return "NSPersistentStoreUbiquitousTransitionTypeAccountAdded"
	case NSPersistentStoreUbiquitousTransitionTypeAccountRemoved:
		return "NSPersistentStoreUbiquitousTransitionTypeAccountRemoved"
	case NSPersistentStoreUbiquitousTransitionTypeContentRemoved:
		return "NSPersistentStoreUbiquitousTransitionTypeContentRemoved"
	case NSPersistentStoreUbiquitousTransitionTypeInitialImportCompleted:
		return "NSPersistentStoreUbiquitousTransitionTypeInitialImportCompleted"
	default:
		return fmt.Sprintf("NSPersistentStoreUbiquitousTransitionType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreData/NSSnapshotEventType
type NSSnapshotEventType uint

const (
	// NSSnapshotEventMergePolicy: Specifies a change due to conflict resolution during a save operation.
	NSSnapshotEventMergePolicy NSSnapshotEventType = 64
	// NSSnapshotEventRefresh: Specifies a change due to the managed object being refreshed.
	NSSnapshotEventRefresh NSSnapshotEventType = 32
	// NSSnapshotEventRollback: Specifies a change due to the managed object context being rolled back.
	NSSnapshotEventRollback NSSnapshotEventType = 16
	// NSSnapshotEventUndoDeletion: Specifies a change due to undo from deletion.
	NSSnapshotEventUndoDeletion NSSnapshotEventType = 4
	// NSSnapshotEventUndoInsertion: Specifies a change due to undo from insertion.
	NSSnapshotEventUndoInsertion NSSnapshotEventType = 2
	// NSSnapshotEventUndoUpdate: Specifies a change due to a property-level undo.
	NSSnapshotEventUndoUpdate NSSnapshotEventType = 8
)

func (e NSSnapshotEventType) String() string {
	switch e {
	case NSSnapshotEventMergePolicy:
		return "NSSnapshotEventMergePolicy"
	case NSSnapshotEventRefresh:
		return "NSSnapshotEventRefresh"
	case NSSnapshotEventRollback:
		return "NSSnapshotEventRollback"
	case NSSnapshotEventUndoDeletion:
		return "NSSnapshotEventUndoDeletion"
	case NSSnapshotEventUndoInsertion:
		return "NSSnapshotEventUndoInsertion"
	case NSSnapshotEventUndoUpdate:
		return "NSSnapshotEventUndoUpdate"
	default:
		return fmt.Sprintf("NSSnapshotEventType(%d)", e)
	}
}
