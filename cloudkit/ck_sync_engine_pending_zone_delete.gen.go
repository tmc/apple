// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [CKSyncEnginePendingZoneDelete] class.
var (
	_CKSyncEnginePendingZoneDeleteClass     CKSyncEnginePendingZoneDeleteClass
	_CKSyncEnginePendingZoneDeleteClassOnce sync.Once
)

func getCKSyncEnginePendingZoneDeleteClass() CKSyncEnginePendingZoneDeleteClass {
	_CKSyncEnginePendingZoneDeleteClassOnce.Do(func() {
		_CKSyncEnginePendingZoneDeleteClass = CKSyncEnginePendingZoneDeleteClass{class: objc.GetClass("CKSyncEnginePendingZoneDelete")}
	})
	return _CKSyncEnginePendingZoneDeleteClass
}

// GetCKSyncEnginePendingZoneDeleteClass returns the class object for CKSyncEnginePendingZoneDelete.
func GetCKSyncEnginePendingZoneDeleteClass() CKSyncEnginePendingZoneDeleteClass {
	return getCKSyncEnginePendingZoneDeleteClass()
}

type CKSyncEnginePendingZoneDeleteClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEnginePendingZoneDeleteClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEnginePendingZoneDeleteClass) Alloc() CKSyncEnginePendingZoneDelete {
	rv := objc.Send[CKSyncEnginePendingZoneDelete](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes an unsent record zone deletion.
//
// # Creating a pending zone delete
//
//   - [CKSyncEnginePendingZoneDelete.InitWithZoneID]: Creates a pending zone delete for the specified record zone identifier.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingZoneDelete
type CKSyncEnginePendingZoneDelete struct {
	CKSyncEnginePendingDatabaseChange
}

// CKSyncEnginePendingZoneDeleteFromID constructs a [CKSyncEnginePendingZoneDelete] from an objc.ID.
//
// An object that describes an unsent record zone deletion.
func CKSyncEnginePendingZoneDeleteFromID(id objc.ID) CKSyncEnginePendingZoneDelete {
	return CKSyncEnginePendingZoneDelete{CKSyncEnginePendingDatabaseChange: CKSyncEnginePendingDatabaseChangeFromID(id)}
}

// Ensure CKSyncEnginePendingZoneDelete implements ICKSyncEnginePendingZoneDelete.
var _ ICKSyncEnginePendingZoneDelete = CKSyncEnginePendingZoneDelete{}

// An interface definition for the [CKSyncEnginePendingZoneDelete] class.
//
// # Creating a pending zone delete
//
//   - [ICKSyncEnginePendingZoneDelete.InitWithZoneID]: Creates a pending zone delete for the specified record zone identifier.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingZoneDelete
type ICKSyncEnginePendingZoneDelete interface {
	ICKSyncEnginePendingDatabaseChange

	// Topic: Creating a pending zone delete

	// Creates a pending zone delete for the specified record zone identifier.
	InitWithZoneID(zoneID ICKRecordZoneID) CKSyncEnginePendingZoneDelete
}

// Init initializes the instance.
func (c CKSyncEnginePendingZoneDelete) Init() CKSyncEnginePendingZoneDelete {
	rv := objc.Send[CKSyncEnginePendingZoneDelete](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEnginePendingZoneDelete) Autorelease() CKSyncEnginePendingZoneDelete {
	rv := objc.Send[CKSyncEnginePendingZoneDelete](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEnginePendingZoneDelete creates a new CKSyncEnginePendingZoneDelete instance.
func NewCKSyncEnginePendingZoneDelete() CKSyncEnginePendingZoneDelete {
	class := getCKSyncEnginePendingZoneDeleteClass()
	rv := objc.Send[CKSyncEnginePendingZoneDelete](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a pending zone delete for the specified record zone identifier.
//
// zoneID: The unique identifier of the record zone to delete.
//
// # Return Value
//
// An initialized pending zone delete, or `nil` if CloudKit can’t create it.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingZoneDelete/initWithZoneID:
func NewCKSyncEnginePendingZoneDeleteWithZoneID(zoneID ICKRecordZoneID) CKSyncEnginePendingZoneDelete {
	instance := getCKSyncEnginePendingZoneDeleteClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithZoneID:"), zoneID)
	return CKSyncEnginePendingZoneDeleteFromID(rv)
}

// Creates a pending zone delete for the specified record zone identifier.
//
// zoneID: The unique identifier of the record zone to delete.
//
// # Return Value
//
// An initialized pending zone delete, or `nil` if CloudKit can’t create it.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingZoneDelete/initWithZoneID:
func (c CKSyncEnginePendingZoneDelete) InitWithZoneID(zoneID ICKRecordZoneID) CKSyncEnginePendingZoneDelete {
	rv := objc.Send[CKSyncEnginePendingZoneDelete](c.ID, objc.Sel("initWithZoneID:"), zoneID)
	return rv
}
