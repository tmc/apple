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

// The class instance for the [NSEntityMigrationPolicy] class.
var (
	_NSEntityMigrationPolicyClass     NSEntityMigrationPolicyClass
	_NSEntityMigrationPolicyClassOnce sync.Once
)

func getNSEntityMigrationPolicyClass() NSEntityMigrationPolicyClass {
	_NSEntityMigrationPolicyClassOnce.Do(func() {
		_NSEntityMigrationPolicyClass = NSEntityMigrationPolicyClass{class: objc.GetClass("NSEntityMigrationPolicy")}
	})
	return _NSEntityMigrationPolicyClass
}

// GetNSEntityMigrationPolicyClass returns the class object for NSEntityMigrationPolicy.
func GetNSEntityMigrationPolicyClass() NSEntityMigrationPolicyClass {
	return getNSEntityMigrationPolicyClass()
}

type NSEntityMigrationPolicyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSEntityMigrationPolicyClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSEntityMigrationPolicyClass) Alloc() NSEntityMigrationPolicy {
	rv := objc.Send[NSEntityMigrationPolicy](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A policy instance that customizes the migration process for an entity
// mapping.
//
// # Overview
//
// You set the policy for an entity mapping by passing the name of the
// migration policy class as the argument to
// [NSEntityMapping.EntityMigrationPolicyClassName]. Typically, you specify
// the name in the Xcode mapping model editor.
//
// # Customizing Stages of the Mapping Life Cycle
//
//   - [NSEntityMigrationPolicy.BeginEntityMappingManagerError]: Sets up state information before the start of a given entity mapping.
//   - [NSEntityMigrationPolicy.CreateDestinationInstancesForSourceInstanceEntityMappingManagerError]: Creates the destination instance(s) for a given source instance.
//   - [NSEntityMigrationPolicy.EndInstanceCreationForEntityMappingManagerError]: Indicates the end of the instance creation stage for the specified entity mapping, and the precursor to the next migration stage.
//   - [NSEntityMigrationPolicy.CreateRelationshipsForDestinationInstanceEntityMappingManagerError]: Constructs the relationships between the newly-created destination instances.
//   - [NSEntityMigrationPolicy.EndRelationshipCreationForEntityMappingManagerError]: Indicates the end of the relationship creation stage for the specified entity mapping.
//   - [NSEntityMigrationPolicy.PerformCustomValidationForEntityMappingManagerError]: Provides the option to perform custom validation on migrated objects during the validation stage of the entity migration policy.
//   - [NSEntityMigrationPolicy.EndEntityMappingManagerError]: Performs cleanup at the end of the migration, from any phase of the mapping.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMigrationPolicy
type NSEntityMigrationPolicy struct {
	objectivec.Object
}

// NSEntityMigrationPolicyFromID constructs a [NSEntityMigrationPolicy] from an objc.ID.
//
// A policy instance that customizes the migration process for an entity
// mapping.
func NSEntityMigrationPolicyFromID(id objc.ID) NSEntityMigrationPolicy {
	return NSEntityMigrationPolicy{objectivec.Object{ID: id}}
}

// NOTE: NSEntityMigrationPolicy adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSEntityMigrationPolicy] class.
//
// # Customizing Stages of the Mapping Life Cycle
//
//   - [INSEntityMigrationPolicy.BeginEntityMappingManagerError]: Sets up state information before the start of a given entity mapping.
//   - [INSEntityMigrationPolicy.CreateDestinationInstancesForSourceInstanceEntityMappingManagerError]: Creates the destination instance(s) for a given source instance.
//   - [INSEntityMigrationPolicy.EndInstanceCreationForEntityMappingManagerError]: Indicates the end of the instance creation stage for the specified entity mapping, and the precursor to the next migration stage.
//   - [INSEntityMigrationPolicy.CreateRelationshipsForDestinationInstanceEntityMappingManagerError]: Constructs the relationships between the newly-created destination instances.
//   - [INSEntityMigrationPolicy.EndRelationshipCreationForEntityMappingManagerError]: Indicates the end of the relationship creation stage for the specified entity mapping.
//   - [INSEntityMigrationPolicy.PerformCustomValidationForEntityMappingManagerError]: Provides the option to perform custom validation on migrated objects during the validation stage of the entity migration policy.
//   - [INSEntityMigrationPolicy.EndEntityMappingManagerError]: Performs cleanup at the end of the migration, from any phase of the mapping.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMigrationPolicy
type INSEntityMigrationPolicy interface {
	objectivec.IObject

	// Topic: Customizing Stages of the Mapping Life Cycle

	// Sets up state information before the start of a given entity mapping.
	BeginEntityMappingManagerError(mapping INSEntityMapping, manager INSMigrationManager) (bool, error)
	// Creates the destination instance(s) for a given source instance.
	CreateDestinationInstancesForSourceInstanceEntityMappingManagerError(sInstance INSManagedObject, mapping INSEntityMapping, manager INSMigrationManager) (bool, error)
	// Indicates the end of the instance creation stage for the specified entity mapping, and the precursor to the next migration stage.
	EndInstanceCreationForEntityMappingManagerError(mapping INSEntityMapping, manager INSMigrationManager) (bool, error)
	// Constructs the relationships between the newly-created destination instances.
	CreateRelationshipsForDestinationInstanceEntityMappingManagerError(dInstance INSManagedObject, mapping INSEntityMapping, manager INSMigrationManager) (bool, error)
	// Indicates the end of the relationship creation stage for the specified entity mapping.
	EndRelationshipCreationForEntityMappingManagerError(mapping INSEntityMapping, manager INSMigrationManager) (bool, error)
	// Provides the option to perform custom validation on migrated objects during the validation stage of the entity migration policy.
	PerformCustomValidationForEntityMappingManagerError(mapping INSEntityMapping, manager INSMigrationManager) (bool, error)
	// Performs cleanup at the end of the migration, from any phase of the mapping.
	EndEntityMappingManagerError(mapping INSEntityMapping, manager INSMigrationManager) (bool, error)
}

// Init initializes the instance.
func (e NSEntityMigrationPolicy) Init() NSEntityMigrationPolicy {
	rv := objc.Send[NSEntityMigrationPolicy](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e NSEntityMigrationPolicy) Autorelease() NSEntityMigrationPolicy {
	rv := objc.Send[NSEntityMigrationPolicy](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSEntityMigrationPolicy creates a new NSEntityMigrationPolicy instance.
func NewNSEntityMigrationPolicy() NSEntityMigrationPolicy {
	class := getNSEntityMigrationPolicyClass()
	rv := objc.Send[NSEntityMigrationPolicy](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Sets up state information before the start of a given entity mapping.
//
// mapping: The mapping object in use.
//
// manager: The migration manager performing the migration.
//
// # Discussion
//
// This method is the precursor to the creation stage. In a custom class, you
// can implement this method to set up any state information that will be
// useful for the duration of the migration.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMigrationPolicy/begin(_:with:)
func (e NSEntityMigrationPolicy) BeginEntityMappingManagerError(mapping INSEntityMapping, manager INSMigrationManager) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("beginEntityMapping:manager:error:"), mapping, manager, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("beginEntityMapping:manager:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Creates the destination instance(s) for a given source instance.
//
// sInstance: The source instance for which to create destination instances.
//
// mapping: The mapping object in use.
//
// manager: The migration manager performing the migration.
//
// # Discussion
//
// This method is invoked by the migration manager on each source instance (as
// specified by the [NSEntityMapping.SourceExpression] in the mapping) to
// create the corresponding destination instance(s). It also associates the
// source and destination instances by calling [NSMigrationManager]’s
// [NSMigrationManager.AssociateSourceInstanceWithDestinationInstanceForEntityMapping]
// method.
//
// # Special Considerations
//
// If you override this method and do not invoke `super`, you must invoke
// [NSMigrationManager]’s
// [NSMigrationManager.AssociateSourceInstanceWithDestinationInstanceForEntityMapping]
// to associate the source and destination instances as required. .
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMigrationPolicy/createDestinationInstances(forSource:in:manager:)
func (e NSEntityMigrationPolicy) CreateDestinationInstancesForSourceInstanceEntityMappingManagerError(sInstance INSManagedObject, mapping INSEntityMapping, manager INSMigrationManager) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("createDestinationInstancesForSourceInstance:entityMapping:manager:error:"), sInstance, mapping, manager, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("createDestinationInstancesForSourceInstance:entityMapping:manager:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Indicates the end of the instance creation stage for the specified entity
// mapping, and the precursor to the next migration stage.
//
// mapping: The mapping object in use.
//
// manager: The migration manager performing the migration.
//
// # Discussion
//
// You can override this method to clean up state from the creation of
// destination or to prepare state for the creation of relationships.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMigrationPolicy/endInstanceCreation(forMapping:manager:)
func (e NSEntityMigrationPolicy) EndInstanceCreationForEntityMappingManagerError(mapping INSEntityMapping, manager INSMigrationManager) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("endInstanceCreationForEntityMapping:manager:error:"), mapping, manager, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("endInstanceCreationForEntityMapping:manager:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Constructs the relationships between the newly-created destination
// instances.
//
// dInstance: The destination instance for which to create relationships.
//
// mapping: The mapping object in use.
//
// manager: The migration manager performing the migration.
//
// # Discussion
//
// You can use this stage to (re)create relationships between migrated
// objects—you use the association lookup methods on the
// [NSMigrationManager] instance to determine the appropriate relationship
// targets.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMigrationPolicy/createRelationships(forDestination:in:manager:)
func (e NSEntityMigrationPolicy) CreateRelationshipsForDestinationInstanceEntityMappingManagerError(dInstance INSManagedObject, mapping INSEntityMapping, manager INSMigrationManager) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("createRelationshipsForDestinationInstance:entityMapping:manager:error:"), dInstance, mapping, manager, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("createRelationshipsForDestinationInstance:entityMapping:manager:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Indicates the end of the relationship creation stage for the specified
// entity mapping.
//
// mapping: The mapping object in use.
//
// manager: The migration manager performing the migration.
//
// # Discussion
//
// This method is invoked after
// [NSEntityMigrationPolicy.CreateRelationshipsForDestinationInstanceEntityMappingManagerError];
// you can override it to clean up state from the creation of relationships,
// or prepare state for custom validation in
// [NSEntityMigrationPolicy.PerformCustomValidationForEntityMappingManagerError].
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMigrationPolicy/endRelationshipCreation(forMapping:manager:)
func (e NSEntityMigrationPolicy) EndRelationshipCreationForEntityMappingManagerError(mapping INSEntityMapping, manager INSMigrationManager) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("endRelationshipCreationForEntityMapping:manager:error:"), mapping, manager, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("endRelationshipCreationForEntityMapping:manager:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Provides the option to perform custom validation on migrated objects during
// the validation stage of the entity migration policy.
//
// mapping: The mapping object in use.
//
// manager: The migration manager performing the migration.
//
// # Discussion
//
// This method is called before the default save validation is performed by
// the framework.
//
// If you implement this method, you must manually obtain the collection of
// objects you are interested in validating.
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMigrationPolicy/performCustomValidation(forMapping:manager:)
func (e NSEntityMigrationPolicy) PerformCustomValidationForEntityMappingManagerError(mapping INSEntityMapping, manager INSMigrationManager) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("performCustomValidationForEntityMapping:manager:error:"), mapping, manager, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("performCustomValidationForEntityMapping:manager:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Performs cleanup at the end of the migration, from any phase of the
// mapping.
//
// mapping: The mapping object in use.
//
// manager: The migration manager performing the migration.
//
// # Discussion
//
// This is the end to the given entity mapping. You can implement this method
// to perform any clean-up at the end of the migration (from any of the three
// phases of the mapping).
//
// See: https://developer.apple.com/documentation/CoreData/NSEntityMigrationPolicy/end(_:manager:)
func (e NSEntityMigrationPolicy) EndEntityMappingManagerError(mapping INSEntityMapping, manager INSMigrationManager) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("endEntityMapping:manager:error:"), mapping, manager, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("endEntityMapping:manager:error: returned NO with nil NSError")
	}
	return rv, nil

}
