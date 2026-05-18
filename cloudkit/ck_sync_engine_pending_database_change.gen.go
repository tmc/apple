// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKSyncEnginePendingDatabaseChange] class.
var (
	_CKSyncEnginePendingDatabaseChangeClass     CKSyncEnginePendingDatabaseChangeClass
	_CKSyncEnginePendingDatabaseChangeClassOnce sync.Once
)

func getCKSyncEnginePendingDatabaseChangeClass() CKSyncEnginePendingDatabaseChangeClass {
	_CKSyncEnginePendingDatabaseChangeClassOnce.Do(func() {
		_CKSyncEnginePendingDatabaseChangeClass = CKSyncEnginePendingDatabaseChangeClass{class: objc.GetClass("CKSyncEnginePendingDatabaseChange")}
	})
	return _CKSyncEnginePendingDatabaseChangeClass
}

// GetCKSyncEnginePendingDatabaseChangeClass returns the class object for CKSyncEnginePendingDatabaseChange.
func GetCKSyncEnginePendingDatabaseChangeClass() CKSyncEnginePendingDatabaseChangeClass {
	return getCKSyncEnginePendingDatabaseChangeClass()
}

type CKSyncEnginePendingDatabaseChangeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEnginePendingDatabaseChangeClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEnginePendingDatabaseChangeClass) Alloc() CKSyncEnginePendingDatabaseChange {
	rv := objc.Send[CKSyncEnginePendingDatabaseChange](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes an unsent database modification.
//
// # Understanding the change
//
//   - [CKSyncEnginePendingDatabaseChange.Type]: The type of database change.
//   - [CKSyncEnginePendingDatabaseChange.ZoneID]: The identifier of the record zone to change.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingDatabaseChange
type CKSyncEnginePendingDatabaseChange struct {
	objectivec.Object
}

// CKSyncEnginePendingDatabaseChangeFromID constructs a [CKSyncEnginePendingDatabaseChange] from an objc.ID.
//
// An object that describes an unsent database modification.
func CKSyncEnginePendingDatabaseChangeFromID(id objc.ID) CKSyncEnginePendingDatabaseChange {
	return CKSyncEnginePendingDatabaseChange{objectivec.Object{ID: id}}
}

// Ensure CKSyncEnginePendingDatabaseChange implements ICKSyncEnginePendingDatabaseChange.
var _ ICKSyncEnginePendingDatabaseChange = CKSyncEnginePendingDatabaseChange{}

// An interface definition for the [CKSyncEnginePendingDatabaseChange] class.
//
// # Understanding the change
//
//   - [ICKSyncEnginePendingDatabaseChange.Type]: The type of database change.
//   - [ICKSyncEnginePendingDatabaseChange.ZoneID]: The identifier of the record zone to change.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingDatabaseChange
type ICKSyncEnginePendingDatabaseChange interface {
	objectivec.IObject

	// Topic: Understanding the change

	// The type of database change.
	Type() CKSyncEnginePendingDatabaseChangeType
	// The identifier of the record zone to change.
	ZoneID() ICKRecordZoneID
}

// Init initializes the instance.
func (c CKSyncEnginePendingDatabaseChange) Init() CKSyncEnginePendingDatabaseChange {
	rv := objc.Send[CKSyncEnginePendingDatabaseChange](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEnginePendingDatabaseChange) Autorelease() CKSyncEnginePendingDatabaseChange {
	rv := objc.Send[CKSyncEnginePendingDatabaseChange](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEnginePendingDatabaseChange creates a new CKSyncEnginePendingDatabaseChange instance.
func NewCKSyncEnginePendingDatabaseChange() CKSyncEnginePendingDatabaseChange {
	class := getCKSyncEnginePendingDatabaseChangeClass()
	rv := objc.Send[CKSyncEnginePendingDatabaseChange](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The type of database change.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingDatabaseChange/type
func (c CKSyncEnginePendingDatabaseChange) Type() CKSyncEnginePendingDatabaseChangeType {
	rv := objc.Send[CKSyncEnginePendingDatabaseChangeType](c.ID, objc.Sel("type"))
	return CKSyncEnginePendingDatabaseChangeType(rv)
}

// The identifier of the record zone to change.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingDatabaseChange/zoneID
func (c CKSyncEnginePendingDatabaseChange) ZoneID() ICKRecordZoneID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("zoneID"))
	return CKRecordZoneIDFromID(objc.ID(rv))
}
