// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMergeConflict] class.
var (
	_NSMergeConflictClass     NSMergeConflictClass
	_NSMergeConflictClassOnce sync.Once
)

func getNSMergeConflictClass() NSMergeConflictClass {
	_NSMergeConflictClassOnce.Do(func() {
		_NSMergeConflictClass = NSMergeConflictClass{class: objc.GetClass("NSMergeConflict")}
	})
	return _NSMergeConflictClass
}

// GetNSMergeConflictClass returns the class object for NSMergeConflict.
func GetNSMergeConflictClass() NSMergeConflictClass {
	return getNSMergeConflictClass()
}

type NSMergeConflictClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSMergeConflictClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMergeConflictClass) Alloc() NSMergeConflict {
	rv := objc.Send[NSMergeConflict](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An encapsulation of conflicts that occur during an attempt to save changes
// in a managed object context.
//
// # Overview
//
// A conflict can occur in two situations:
//
// - Between the managed object context and its in-memory cached state at the
// persistent store coordinator layer. - Between the cached state at the
// persistent store coordinator layer and the external store (file, database,
// and so forth). In this case, the merge conflict has a cached snapshot and a
// persisted snapshot. The source object is also provided as a convenience,
// but it is not directly involved in the conflict.
//
// Snapshot dictionaries include values for all attributes and to-one
// relationships, but not to-many relationships. Relationship values are
// [NSManagedObjectID] references. To-many relationships must be pulled from
// the persistent store as needed.
//
// # Initializing a Merge Conflict
//
//   - [NSMergeConflict.InitWithSourceNewVersionOldVersionCachedSnapshotPersistedSnapshot]: Initializes a merge conflict.
//
// # Accessing Merge Conflict Details
//
//   - [NSMergeConflict.SourceObject]: The source object for the conflict.
//   - [NSMergeConflict.ObjectSnapshot]: A dictionary containing the values of the source object.
//   - [NSMergeConflict.CachedSnapshot]: A dictionary containing the values of the source object held in the persistent store coordinator layer.
//   - [NSMergeConflict.PersistedSnapshot]: A dictionary containing the values of the source object held in the persistent store.
//   - [NSMergeConflict.NewVersionNumber]: The new version number for the change.
//   - [NSMergeConflict.OldVersionNumber]: The old version number for the change.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergeConflict
type NSMergeConflict struct {
	objectivec.Object
}

// NSMergeConflictFromID constructs a [NSMergeConflict] from an objc.ID.
//
// An encapsulation of conflicts that occur during an attempt to save changes
// in a managed object context.
func NSMergeConflictFromID(id objc.ID) NSMergeConflict {
	return NSMergeConflict{objectivec.Object{ID: id}}
}

// NOTE: NSMergeConflict adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMergeConflict] class.
//
// # Initializing a Merge Conflict
//
//   - [INSMergeConflict.InitWithSourceNewVersionOldVersionCachedSnapshotPersistedSnapshot]: Initializes a merge conflict.
//
// # Accessing Merge Conflict Details
//
//   - [INSMergeConflict.SourceObject]: The source object for the conflict.
//   - [INSMergeConflict.ObjectSnapshot]: A dictionary containing the values of the source object.
//   - [INSMergeConflict.CachedSnapshot]: A dictionary containing the values of the source object held in the persistent store coordinator layer.
//   - [INSMergeConflict.PersistedSnapshot]: A dictionary containing the values of the source object held in the persistent store.
//   - [INSMergeConflict.NewVersionNumber]: The new version number for the change.
//   - [INSMergeConflict.OldVersionNumber]: The old version number for the change.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergeConflict
type INSMergeConflict interface {
	objectivec.IObject

	// Topic: Initializing a Merge Conflict

	// Initializes a merge conflict.
	InitWithSourceNewVersionOldVersionCachedSnapshotPersistedSnapshot(srcObject INSManagedObject, newvers uint, oldvers uint, cachesnap foundation.INSDictionary, persnap foundation.INSDictionary) NSMergeConflict

	// Topic: Accessing Merge Conflict Details

	// The source object for the conflict.
	SourceObject() INSManagedObject
	// A dictionary containing the values of the source object.
	ObjectSnapshot() foundation.INSDictionary
	// A dictionary containing the values of the source object held in the persistent store coordinator layer.
	CachedSnapshot() foundation.INSDictionary
	// A dictionary containing the values of the source object held in the persistent store.
	PersistedSnapshot() foundation.INSDictionary
	// The new version number for the change.
	NewVersionNumber() uint
	// The old version number for the change.
	OldVersionNumber() uint
}

// Init initializes the instance.
func (m NSMergeConflict) Init() NSMergeConflict {
	rv := objc.Send[NSMergeConflict](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMergeConflict) Autorelease() NSMergeConflict {
	rv := objc.Send[NSMergeConflict](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMergeConflict creates a new NSMergeConflict instance.
func NewNSMergeConflict() NSMergeConflict {
	class := getNSMergeConflictClass()
	rv := objc.Send[NSMergeConflict](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a merge conflict.
//
// srcObject: The source object for the conflict.
//
// newvers: The new version number for the change.
//
// A value of 0 means the object was deleted and the corresponding snapshot is
// `nil`.
//
// oldvers: The old version number for the change.
//
// cachesnap: A dictionary containing the values of `srcObject` held in the persistent
// store coordinator layer.
//
// persnap: A dictionary containing the values of `srcObject` held in the persistent
// store.
//
// # Return Value
//
// A merge conflict object initialized with the given parameters.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergeConflict/init(source:newVersion:oldVersion:cachedSnapshot:persistedSnapshot:)
func NewMergeConflictWithSourceNewVersionOldVersionCachedSnapshotPersistedSnapshot(srcObject INSManagedObject, newvers uint, oldvers uint, cachesnap foundation.INSDictionary, persnap foundation.INSDictionary) NSMergeConflict {
	instance := getNSMergeConflictClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:newVersion:oldVersion:cachedSnapshot:persistedSnapshot:"), srcObject, newvers, oldvers, cachesnap, persnap)
	return NSMergeConflictFromID(rv)
}

// Initializes a merge conflict.
//
// srcObject: The source object for the conflict.
//
// newvers: The new version number for the change.
//
// A value of 0 means the object was deleted and the corresponding snapshot is
// `nil`.
//
// oldvers: The old version number for the change.
//
// cachesnap: A dictionary containing the values of `srcObject` held in the persistent
// store coordinator layer.
//
// persnap: A dictionary containing the values of `srcObject` held in the persistent
// store.
//
// # Return Value
//
// A merge conflict object initialized with the given parameters.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergeConflict/init(source:newVersion:oldVersion:cachedSnapshot:persistedSnapshot:)
func (m NSMergeConflict) InitWithSourceNewVersionOldVersionCachedSnapshotPersistedSnapshot(srcObject INSManagedObject, newvers uint, oldvers uint, cachesnap foundation.INSDictionary, persnap foundation.INSDictionary) NSMergeConflict {
	rv := objc.Send[NSMergeConflict](m.ID, objc.Sel("initWithSource:newVersion:oldVersion:cachedSnapshot:persistedSnapshot:"), srcObject, newvers, oldvers, cachesnap, persnap)
	return rv
}

// The source object for the conflict.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergeConflict/sourceObject
func (m NSMergeConflict) SourceObject() INSManagedObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("sourceObject"))
	return NSManagedObjectFromID(objc.ID(rv))
}

// A dictionary containing the values of the source object.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergeConflict/objectSnapshot
func (m NSMergeConflict) ObjectSnapshot() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("objectSnapshot"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// A dictionary containing the values of the source object held in the
// persistent store coordinator layer.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergeConflict/cachedSnapshot
func (m NSMergeConflict) CachedSnapshot() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("cachedSnapshot"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// A dictionary containing the values of the source object held in the
// persistent store.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergeConflict/persistedSnapshot
func (m NSMergeConflict) PersistedSnapshot() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("persistedSnapshot"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The new version number for the change.
//
// # Discussion
//
// A new version number of 0 means the object was deleted and the
// corresponding snapshot is `nil`.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergeConflict/newVersionNumber
func (m NSMergeConflict) NewVersionNumber() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("newVersionNumber"))
	return rv
}

// The old version number for the change.
//
// See: https://developer.apple.com/documentation/CoreData/NSMergeConflict/oldVersionNumber
func (m NSMergeConflict) OldVersionNumber() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("oldVersionNumber"))
	return rv
}
