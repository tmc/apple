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

// The class instance for the [NSMergePolicy] class.
var (
	_NSMergePolicyClass     NSMergePolicyClass
	_NSMergePolicyClassOnce sync.Once
)

func getNSMergePolicyClass() NSMergePolicyClass {
	_NSMergePolicyClassOnce.Do(func() {
		_NSMergePolicyClass = NSMergePolicyClass{class: objc.GetClass("NSMergePolicy")}
	})
	return _NSMergePolicyClass
}

// GetNSMergePolicyClass returns the class object for NSMergePolicy.
func GetNSMergePolicyClass() NSMergePolicyClass {
	return getNSMergePolicyClass()
}

type NSMergePolicyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSMergePolicyClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMergePolicyClass) Alloc() NSMergePolicy {
	rv := objc.Send[NSMergePolicy](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A policy object that you use to resolve conflicts between the persistent
// store and in-memory versions of managed objects.
//
// # Overview
//
// A conflict is a mismatch between state held at two different layers in the
// Core Data stack. A conflict can arise when you save a managed object
// context and you have stale data at another layer. There are two places in
// which a conflict may occur:
//
// - Between the managed object context layer and its in-memory cached state
// at the persistent store coordinator layer. - Between the cached state at
// the persistent store coordinator and the external store (file, database,
// and so forth).
//
// Conflicts are represented by instances of [NSMergeConflict].
//
// # Getting a Merge Policy
//
//   - [NSMergePolicy.InitWithMergeType]: Returns a merge policy initialized with a given policy type.
//   - [NSMergePolicy.MergeType]: The merge type.
//
// # Resolving a Conflict
//
//   - [NSMergePolicy.ResolveConflictsError]: Resolves the conflicts in a given list.
//   - [NSMergePolicy.ResolveConstraintConflictsError]: Resolves the conflicts in a given list.
//   - [NSMergePolicy.ResolveOptimisticLockingVersionConflictsError]: Resolves the conflicts in a given list.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergePolicy
type NSMergePolicy struct {
	objectivec.Object
}

// NSMergePolicyFromID constructs a [NSMergePolicy] from an objc.ID.
//
// A policy object that you use to resolve conflicts between the persistent
// store and in-memory versions of managed objects.
func NSMergePolicyFromID(id objc.ID) NSMergePolicy {
	return NSMergePolicy{objectivec.Object{ID: id}}
}

// NOTE: NSMergePolicy adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMergePolicy] class.
//
// # Getting a Merge Policy
//
//   - [INSMergePolicy.InitWithMergeType]: Returns a merge policy initialized with a given policy type.
//   - [INSMergePolicy.MergeType]: The merge type.
//
// # Resolving a Conflict
//
//   - [INSMergePolicy.ResolveConflictsError]: Resolves the conflicts in a given list.
//   - [INSMergePolicy.ResolveConstraintConflictsError]: Resolves the conflicts in a given list.
//   - [INSMergePolicy.ResolveOptimisticLockingVersionConflictsError]: Resolves the conflicts in a given list.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergePolicy
type INSMergePolicy interface {
	objectivec.IObject

	// Topic: Getting a Merge Policy

	// Returns a merge policy initialized with a given policy type.
	InitWithMergeType(ty NSMergePolicyType) NSMergePolicy
	// The merge type.
	MergeType() NSMergePolicyType

	// Topic: Resolving a Conflict

	// Resolves the conflicts in a given list.
	ResolveConflictsError(list foundation.INSArray) (bool, error)
	// Resolves the conflicts in a given list.
	ResolveConstraintConflictsError(list []NSConstraintConflict) (bool, error)
	// Resolves the conflicts in a given list.
	ResolveOptimisticLockingVersionConflictsError(list []NSMergeConflict) (bool, error)
}

// Init initializes the instance.
func (m NSMergePolicy) Init() NSMergePolicy {
	rv := objc.Send[NSMergePolicy](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMergePolicy) Autorelease() NSMergePolicy {
	rv := objc.Send[NSMergePolicy](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMergePolicy creates a new NSMergePolicy instance.
func NewNSMergePolicy() NSMergePolicy {
	class := getNSMergePolicyClass()
	rv := objc.Send[NSMergePolicy](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a merge policy initialized with a given policy type.
//
// ty: A merge policy type.
//
// # Return Value
//
// A merge policy initialized with a given policy type.
//
// # Discussion
//
// If you override this method in a subclass, you should invoke the
// superclass’s implementation with the merge policy that is closest to the
// behavior you want.
//
// - This will make it easier to use the superclass’s implementation of
// [NSMergePolicy.ResolveConflictsError] and then customize the results. - Due
// to the complexity of merging to-many relationships, this class is designed
// with the expectation that you call super as the base implementation.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergePolicy/init(merge:)
func NewMergePolicyWithMergeType(ty NSMergePolicyType) NSMergePolicy {
	instance := getNSMergePolicyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMergeType:"), ty)
	return NSMergePolicyFromID(rv)
}

// Returns a merge policy initialized with a given policy type.
//
// ty: A merge policy type.
//
// # Return Value
//
// A merge policy initialized with a given policy type.
//
// # Discussion
//
// If you override this method in a subclass, you should invoke the
// superclass’s implementation with the merge policy that is closest to the
// behavior you want.
//
// - This will make it easier to use the superclass’s implementation of
// [NSMergePolicy.ResolveConflictsError] and then customize the results. - Due
// to the complexity of merging to-many relationships, this class is designed
// with the expectation that you call super as the base implementation.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergePolicy/init(merge:)
func (m NSMergePolicy) InitWithMergeType(ty NSMergePolicyType) NSMergePolicy {
	rv := objc.Send[NSMergePolicy](m.ID, objc.Sel("initWithMergeType:"), ty)
	return rv
}

// Resolves the conflicts in a given list.
//
// list: An array of merge conflicts (instances of [NSMergeConflict]).
//
// # Discussion
//
// If you override this method in a subclass, you should typically invoke the
// superclass’s implementation in addition to performing your own
// operations.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergePolicy/resolve(mergeConflicts:)
func (m NSMergePolicy) ResolveConflictsError(list foundation.INSArray) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("resolveConflicts:error:"), list, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("resolveConflicts:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Resolves the conflicts in a given list.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergePolicy/resolve(constraintConflicts:)
func (m NSMergePolicy) ResolveConstraintConflictsError(list []NSConstraintConflict) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("resolveConstraintConflicts:error:"), objectivec.IObjectSliceToNSArray(list), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("resolveConstraintConflicts:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Resolves the conflicts in a given list.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergePolicy/resolve(optimisticLockingConflicts:)
func (m NSMergePolicy) ResolveOptimisticLockingVersionConflictsError(list []NSMergeConflict) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("resolveOptimisticLockingVersionConflicts:error:"), objectivec.IObjectSliceToNSArray(list), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("resolveOptimisticLockingVersionConflicts:error: returned NO with nil NSError")
	}
	return rv, nil

}

// The merge type.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergePolicy/mergeType
func (m NSMergePolicy) MergeType() NSMergePolicyType {
	rv := objc.Send[NSMergePolicyType](m.ID, objc.Sel("mergeType"))
	return NSMergePolicyType(rv)
}

// The default merge policy for all managed object contexts.
//
// # Discussion
//
// If a save fails because of conflicting objects, you can find the IDs of
// those objects in error’s `userInfo` dictionary. Use the
// [NSInsertedObjectsKey] and [NSUpdatedObjectsKey] keys to extract the object
// IDs.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergePolicy/error
//
// [NSInsertedObjectsKey]: https://developer.apple.com/documentation/CoreData/NSInsertedObjectsKey
// [NSUpdatedObjectsKey]: https://developer.apple.com/documentation/CoreData/NSUpdatedObjectsKey
func (_NSMergePolicyClass NSMergePolicyClass) ErrorMergePolicy() NSMergePolicy {
	rv := objc.Send[objc.ID](objc.ID(_NSMergePolicyClass.class), objc.Sel("errorMergePolicy"))
	return NSMergePolicyFromID(objc.ID(rv))
}

// A property-based merge policy that applies external changes.
//
// # Discussion
//
// A policy that merges conflicts between the persistent store’s version of
// the object and the current in-memory version by individual property, with
// external changes trumping in-memory changes.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergePolicy/mergeByPropertyStoreTrump
func (_NSMergePolicyClass NSMergePolicyClass) MergeByPropertyStoreTrumpMergePolicy() NSMergePolicy {
	rv := objc.Send[objc.ID](objc.ID(_NSMergePolicyClass.class), objc.Sel("mergeByPropertyStoreTrumpMergePolicy"))
	return NSMergePolicyFromID(objc.ID(rv))
}

// A property-based merge policy that applies in-memory changes.
//
// # Discussion
//
// A policy that merges conflicts between the persistent store’s version of
// the object and the current in-memory version by individual property, with
// in-memory changes trumping external changes.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergePolicy/mergeByPropertyObjectTrump
func (_NSMergePolicyClass NSMergePolicyClass) MergeByPropertyObjectTrumpMergePolicy() NSMergePolicy {
	rv := objc.Send[objc.ID](objc.ID(_NSMergePolicyClass.class), objc.Sel("mergeByPropertyObjectTrumpMergePolicy"))
	return NSMergePolicyFromID(objc.ID(rv))
}

// A merge policy that overwrites the entire stored object.
//
// # Discussion
//
// This policy merges conflicts between the persistent store’s version of
// the object and the current in-memory version by saving the entire in-memory
// object to the persistent store.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergePolicy/overwrite
func (_NSMergePolicyClass NSMergePolicyClass) OverwriteMergePolicy() NSMergePolicy {
	rv := objc.Send[objc.ID](objc.ID(_NSMergePolicyClass.class), objc.Sel("overwriteMergePolicy"))
	return NSMergePolicyFromID(objc.ID(rv))
}

// A merge policy that discards unsaved changes.
//
// # Discussion
//
// This policy merges conflicts between the persistent store’s version of
// the object and the current in-memory version by discarding unsaved changes.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergePolicy/rollback
func (_NSMergePolicyClass NSMergePolicyClass) RollbackMergePolicy() NSMergePolicy {
	rv := objc.Send[objc.ID](objc.ID(_NSMergePolicyClass.class), objc.Sel("rollbackMergePolicy"))
	return NSMergePolicyFromID(objc.ID(rv))
}
