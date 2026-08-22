// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSConstraintConflict] class.
var (
	_NSConstraintConflictClass     NSConstraintConflictClass
	_NSConstraintConflictClassOnce sync.Once
)

func getNSConstraintConflictClass() NSConstraintConflictClass {
	_NSConstraintConflictClassOnce.Do(func() {
		_NSConstraintConflictClass = NSConstraintConflictClass{class: objc.GetClass("NSConstraintConflict")}
	})
	return _NSConstraintConflictClass
}

// GetNSConstraintConflictClass returns the class object for NSConstraintConflict.
func GetNSConstraintConflictClass() NSConstraintConflictClass {
	return getNSConstraintConflictClass()
}

type NSConstraintConflictClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSConstraintConflictClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSConstraintConflictClass) Alloc() NSConstraintConflict {
	rv := objc.Send[NSConstraintConflict](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An encapsulation of conflicts that occur during an attempt to save a
// managed object.
//
// # Overview
//
// A constraint conflict occurs when your data model is using unique
// constraints and one or more managed objects are violating that constraint.
//
// When this error occurs, the error instance can be interrogated to determine
// which instance of [NSManagedObject] is violating the constraint and which
// property on the [NSManagedObject] instance is in violation.
//
// # Initializing a Conflict
//
//   - [NSConstraintConflict.InitWithConstraintDatabaseObjectDatabaseSnapshotConflictingObjectsConflictingSnapshots]: Initializes a constraint conflict.
//
// # Inspecting a Conflict
//
//   - [NSConstraintConflict.ConflictingObjects]: The managed objects that are in conflict.
//   - [NSConstraintConflict.ConflictingSnapshots]: The original property values of objects in violation of the constraint.
//   - [NSConstraintConflict.Constraint]: The constraint that has been violated.
//   - [NSConstraintConflict.ConstraintValues]: The values that the conflicting objects had when the conflict was created.
//   - [NSConstraintConflict.DatabaseObject]: The object whose database row is using constraint values.
//   - [NSConstraintConflict.DatabaseSnapshot]: The values currently stored in the database.
//
// See: https://developer.apple.com/documentation/CoreData/NSConstraintConflict
type NSConstraintConflict struct {
	objectivec.Object
}

// NSConstraintConflictFromID constructs a [NSConstraintConflict] from an objc.ID.
//
// An encapsulation of conflicts that occur during an attempt to save a
// managed object.
func NSConstraintConflictFromID(id objc.ID) NSConstraintConflict {
	return NSConstraintConflict{objectivec.Object{ID: id}}
}

// NOTE: NSConstraintConflict adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSConstraintConflict] class.
//
// # Initializing a Conflict
//
//   - [INSConstraintConflict.InitWithConstraintDatabaseObjectDatabaseSnapshotConflictingObjectsConflictingSnapshots]: Initializes a constraint conflict.
//
// # Inspecting a Conflict
//
//   - [INSConstraintConflict.ConflictingObjects]: The managed objects that are in conflict.
//   - [INSConstraintConflict.ConflictingSnapshots]: The original property values of objects in violation of the constraint.
//   - [INSConstraintConflict.Constraint]: The constraint that has been violated.
//   - [INSConstraintConflict.ConstraintValues]: The values that the conflicting objects had when the conflict was created.
//   - [INSConstraintConflict.DatabaseObject]: The object whose database row is using constraint values.
//   - [INSConstraintConflict.DatabaseSnapshot]: The values currently stored in the database.
//
// See: https://developer.apple.com/documentation/CoreData/NSConstraintConflict
type INSConstraintConflict interface {
	objectivec.IObject

	// Topic: Initializing a Conflict

	// Initializes a constraint conflict.
	InitWithConstraintDatabaseObjectDatabaseSnapshotConflictingObjectsConflictingSnapshots(contraint []string, databaseObject INSManagedObject, databaseSnapshot foundation.INSDictionary, conflictingObjects []NSManagedObject, conflictingSnapshots foundation.INSArray) NSConstraintConflict

	// Topic: Inspecting a Conflict

	// The managed objects that are in conflict.
	ConflictingObjects() []NSManagedObject
	// The original property values of objects in violation of the constraint.
	ConflictingSnapshots() foundation.INSDictionary
	// The constraint that has been violated.
	Constraint() []string
	// The values that the conflicting objects had when the conflict was created.
	ConstraintValues() foundation.INSDictionary
	// The object whose database row is using constraint values.
	DatabaseObject() INSManagedObject
	// The values currently stored in the database.
	DatabaseSnapshot() foundation.INSDictionary
}

// Init initializes the instance.
func (c NSConstraintConflict) Init() NSConstraintConflict {
	rv := objc.Send[NSConstraintConflict](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSConstraintConflict) Autorelease() NSConstraintConflict {
	rv := objc.Send[NSConstraintConflict](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSConstraintConflict creates a new NSConstraintConflict instance.
func NewNSConstraintConflict() NSConstraintConflict {
	class := getNSConstraintConflictClass()
	rv := objc.Send[NSConstraintConflict](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a constraint conflict.
//
// See: https://developer.apple.com/documentation/CoreData/NSConstraintConflict/init(constraint:database:databaseSnapshot:conflicting:conflictingSnapshots:)
func NewConstraintConflictWithConstraintDatabaseObjectDatabaseSnapshotConflictingObjectsConflictingSnapshots(contraint []string, databaseObject INSManagedObject, databaseSnapshot foundation.INSDictionary, conflictingObjects []NSManagedObject, conflictingSnapshots foundation.INSArray) NSConstraintConflict {
	instance := getNSConstraintConflictClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConstraint:databaseObject:databaseSnapshot:conflictingObjects:conflictingSnapshots:"), objectivec.StringSliceToNSArray(contraint), databaseObject, databaseSnapshot, objectivec.IObjectSliceToNSArray(conflictingObjects), conflictingSnapshots)
	return NSConstraintConflictFromID(rv)
}

// Initializes a constraint conflict.
//
// See: https://developer.apple.com/documentation/CoreData/NSConstraintConflict/init(constraint:database:databaseSnapshot:conflicting:conflictingSnapshots:)
func (c NSConstraintConflict) InitWithConstraintDatabaseObjectDatabaseSnapshotConflictingObjectsConflictingSnapshots(contraint []string, databaseObject INSManagedObject, databaseSnapshot foundation.INSDictionary, conflictingObjects []NSManagedObject, conflictingSnapshots foundation.INSArray) NSConstraintConflict {
	rv := objc.Send[NSConstraintConflict](c.ID, objc.Sel("initWithConstraint:databaseObject:databaseSnapshot:conflictingObjects:conflictingSnapshots:"), objectivec.StringSliceToNSArray(contraint), databaseObject, databaseSnapshot, objectivec.IObjectSliceToNSArray(conflictingObjects), conflictingSnapshots)
	return rv
}

// The managed objects that are in conflict.
//
// See: https://developer.apple.com/documentation/CoreData/NSConstraintConflict/conflictingObjects
func (c NSConstraintConflict) ConflictingObjects() []NSManagedObject {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("conflictingObjects"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSManagedObject {
		return NSManagedObjectFromID(id)
	})
}

// The original property values of objects in violation of the constraint.
//
// See: https://developer.apple.com/documentation/CoreData/NSConstraintConflict/conflictingSnapshots
func (c NSConstraintConflict) ConflictingSnapshots() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("conflictingSnapshots"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The constraint that has been violated.
//
// See: https://developer.apple.com/documentation/CoreData/NSConstraintConflict/constraint
func (c NSConstraintConflict) Constraint() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("constraint"))
	return objc.ConvertSliceToStrings(rv)
}

// The values that the conflicting objects had when the conflict was created.
//
// See: https://developer.apple.com/documentation/CoreData/NSConstraintConflict/constraintValues
func (c NSConstraintConflict) ConstraintValues() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("constraintValues"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The object whose database row is using constraint values.
//
// See: https://developer.apple.com/documentation/CoreData/NSConstraintConflict/databaseObject
func (c NSConstraintConflict) DatabaseObject() INSManagedObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("databaseObject"))
	return NSManagedObjectFromID(objc.ID(rv))
}

// The values currently stored in the database.
//
// See: https://developer.apple.com/documentation/CoreData/NSConstraintConflict/databaseSnapshot
func (c NSConstraintConflict) DatabaseSnapshot() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("databaseSnapshot"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
