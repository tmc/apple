// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [CKSyncEnginePendingZoneSave] class.
var (
	_CKSyncEnginePendingZoneSaveClass     CKSyncEnginePendingZoneSaveClass
	_CKSyncEnginePendingZoneSaveClassOnce sync.Once
)

func getCKSyncEnginePendingZoneSaveClass() CKSyncEnginePendingZoneSaveClass {
	_CKSyncEnginePendingZoneSaveClassOnce.Do(func() {
		_CKSyncEnginePendingZoneSaveClass = CKSyncEnginePendingZoneSaveClass{class: objc.GetClass("CKSyncEnginePendingZoneSave")}
	})
	return _CKSyncEnginePendingZoneSaveClass
}

// GetCKSyncEnginePendingZoneSaveClass returns the class object for CKSyncEnginePendingZoneSave.
func GetCKSyncEnginePendingZoneSaveClass() CKSyncEnginePendingZoneSaveClass {
	return getCKSyncEnginePendingZoneSaveClass()
}

type CKSyncEnginePendingZoneSaveClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEnginePendingZoneSaveClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEnginePendingZoneSaveClass) Alloc() CKSyncEnginePendingZoneSave {
	rv := objc.Send[CKSyncEnginePendingZoneSave](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes an unsent record zone modification.
//
// # Creating a pending zone save
//
//   - [CKSyncEnginePendingZoneSave.InitWithZone]: Creates a pending zone save for the specified record zone.
//
// # Identifying the record zone
//
//   - [CKSyncEnginePendingZoneSave.Zone]: The record zone to save.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingZoneSave
type CKSyncEnginePendingZoneSave struct {
	CKSyncEnginePendingDatabaseChange
}

// CKSyncEnginePendingZoneSaveFromID constructs a [CKSyncEnginePendingZoneSave] from an objc.ID.
//
// An object that describes an unsent record zone modification.
func CKSyncEnginePendingZoneSaveFromID(id objc.ID) CKSyncEnginePendingZoneSave {
	return CKSyncEnginePendingZoneSave{CKSyncEnginePendingDatabaseChange: CKSyncEnginePendingDatabaseChangeFromID(id)}
}

// Ensure CKSyncEnginePendingZoneSave implements ICKSyncEnginePendingZoneSave.
var _ ICKSyncEnginePendingZoneSave = CKSyncEnginePendingZoneSave{}

// An interface definition for the [CKSyncEnginePendingZoneSave] class.
//
// # Creating a pending zone save
//
//   - [ICKSyncEnginePendingZoneSave.InitWithZone]: Creates a pending zone save for the specified record zone.
//
// # Identifying the record zone
//
//   - [ICKSyncEnginePendingZoneSave.Zone]: The record zone to save.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingZoneSave
type ICKSyncEnginePendingZoneSave interface {
	ICKSyncEnginePendingDatabaseChange

	// Topic: Creating a pending zone save

	// Creates a pending zone save for the specified record zone.
	InitWithZone(zone ICKRecordZone) CKSyncEnginePendingZoneSave

	// Topic: Identifying the record zone

	// The record zone to save.
	Zone() ICKRecordZone
}

// Init initializes the instance.
func (c CKSyncEnginePendingZoneSave) Init() CKSyncEnginePendingZoneSave {
	rv := objc.Send[CKSyncEnginePendingZoneSave](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEnginePendingZoneSave) Autorelease() CKSyncEnginePendingZoneSave {
	rv := objc.Send[CKSyncEnginePendingZoneSave](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEnginePendingZoneSave creates a new CKSyncEnginePendingZoneSave instance.
func NewCKSyncEnginePendingZoneSave() CKSyncEnginePendingZoneSave {
	class := getCKSyncEnginePendingZoneSaveClass()
	rv := objc.Send[CKSyncEnginePendingZoneSave](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a pending zone save for the specified record zone.
//
// zone: The record zone to save.
//
// # Return Value
//
// An initialized pending zone save, or `nil` if CloudKit can’t create it.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingZoneSave/initWithZone:
func NewCKSyncEnginePendingZoneSaveWithZone(zone ICKRecordZone) CKSyncEnginePendingZoneSave {
	instance := getCKSyncEnginePendingZoneSaveClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithZone:"), zone)
	return CKSyncEnginePendingZoneSaveFromID(rv)
}

// Creates a pending zone save for the specified record zone.
//
// zone: The record zone to save.
//
// # Return Value
//
// An initialized pending zone save, or `nil` if CloudKit can’t create it.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingZoneSave/initWithZone:
func (c CKSyncEnginePendingZoneSave) InitWithZone(zone ICKRecordZone) CKSyncEnginePendingZoneSave {
	rv := objc.Send[CKSyncEnginePendingZoneSave](c.ID, objc.Sel("initWithZone:"), zone)
	return rv
}

// The record zone to save.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingZoneSave/zone
func (c CKSyncEnginePendingZoneSave) Zone() ICKRecordZone {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("zone"))
	return CKRecordZoneFromID(objc.ID(rv))
}
