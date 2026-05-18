// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKSyncEnginePendingRecordZoneChange] class.
var (
	_CKSyncEnginePendingRecordZoneChangeClass     CKSyncEnginePendingRecordZoneChangeClass
	_CKSyncEnginePendingRecordZoneChangeClassOnce sync.Once
)

func getCKSyncEnginePendingRecordZoneChangeClass() CKSyncEnginePendingRecordZoneChangeClass {
	_CKSyncEnginePendingRecordZoneChangeClassOnce.Do(func() {
		_CKSyncEnginePendingRecordZoneChangeClass = CKSyncEnginePendingRecordZoneChangeClass{class: objc.GetClass("CKSyncEnginePendingRecordZoneChange")}
	})
	return _CKSyncEnginePendingRecordZoneChangeClass
}

// GetCKSyncEnginePendingRecordZoneChangeClass returns the class object for CKSyncEnginePendingRecordZoneChange.
func GetCKSyncEnginePendingRecordZoneChangeClass() CKSyncEnginePendingRecordZoneChangeClass {
	return getCKSyncEnginePendingRecordZoneChangeClass()
}

type CKSyncEnginePendingRecordZoneChangeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEnginePendingRecordZoneChangeClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEnginePendingRecordZoneChangeClass) Alloc() CKSyncEnginePendingRecordZoneChange {
	rv := objc.Send[CKSyncEnginePendingRecordZoneChange](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes an unsent record modification.
//
// # Creating a record zone change
//
//   - [CKSyncEnginePendingRecordZoneChange.InitWithRecordIDType]: Creates a record zone change of the specified type for the given record.
//
// # Accessing the modified record
//
//   - [CKSyncEnginePendingRecordZoneChange.RecordID]: The identifier of the modified record.
//   - [CKSyncEnginePendingRecordZoneChange.Type]: The type of change to make.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingRecordZoneChange
type CKSyncEnginePendingRecordZoneChange struct {
	objectivec.Object
}

// CKSyncEnginePendingRecordZoneChangeFromID constructs a [CKSyncEnginePendingRecordZoneChange] from an objc.ID.
//
// An object that describes an unsent record modification.
func CKSyncEnginePendingRecordZoneChangeFromID(id objc.ID) CKSyncEnginePendingRecordZoneChange {
	return CKSyncEnginePendingRecordZoneChange{objectivec.Object{ID: id}}
}

// Ensure CKSyncEnginePendingRecordZoneChange implements ICKSyncEnginePendingRecordZoneChange.
var _ ICKSyncEnginePendingRecordZoneChange = CKSyncEnginePendingRecordZoneChange{}

// An interface definition for the [CKSyncEnginePendingRecordZoneChange] class.
//
// # Creating a record zone change
//
//   - [ICKSyncEnginePendingRecordZoneChange.InitWithRecordIDType]: Creates a record zone change of the specified type for the given record.
//
// # Accessing the modified record
//
//   - [ICKSyncEnginePendingRecordZoneChange.RecordID]: The identifier of the modified record.
//   - [ICKSyncEnginePendingRecordZoneChange.Type]: The type of change to make.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingRecordZoneChange
type ICKSyncEnginePendingRecordZoneChange interface {
	objectivec.IObject

	// Topic: Creating a record zone change

	// Creates a record zone change of the specified type for the given record.
	InitWithRecordIDType(recordID ICKRecordID, type_ CKSyncEnginePendingRecordZoneChangeType) CKSyncEnginePendingRecordZoneChange

	// Topic: Accessing the modified record

	// The identifier of the modified record.
	RecordID() ICKRecordID
	// The type of change to make.
	Type() CKSyncEnginePendingRecordZoneChangeType
}

// Init initializes the instance.
func (c CKSyncEnginePendingRecordZoneChange) Init() CKSyncEnginePendingRecordZoneChange {
	rv := objc.Send[CKSyncEnginePendingRecordZoneChange](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEnginePendingRecordZoneChange) Autorelease() CKSyncEnginePendingRecordZoneChange {
	rv := objc.Send[CKSyncEnginePendingRecordZoneChange](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEnginePendingRecordZoneChange creates a new CKSyncEnginePendingRecordZoneChange instance.
func NewCKSyncEnginePendingRecordZoneChange() CKSyncEnginePendingRecordZoneChange {
	class := getCKSyncEnginePendingRecordZoneChangeClass()
	rv := objc.Send[CKSyncEnginePendingRecordZoneChange](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a record zone change of the specified type for the given record.
//
// recordID: The identifier of the record to change.
//
// type: The type of change to make.
//
// # Return Value
//
// An initialized record zone change, or `nil` if CloudKit can’t create one.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingRecordZoneChange/initWithRecordID:type:
func NewCKSyncEnginePendingRecordZoneChangeWithRecordIDType(recordID ICKRecordID, type_ CKSyncEnginePendingRecordZoneChangeType) CKSyncEnginePendingRecordZoneChange {
	instance := getCKSyncEnginePendingRecordZoneChangeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRecordID:type:"), recordID, type_)
	return CKSyncEnginePendingRecordZoneChangeFromID(rv)
}

// Creates a record zone change of the specified type for the given record.
//
// recordID: The identifier of the record to change.
//
// type: The type of change to make.
//
// # Return Value
//
// An initialized record zone change, or `nil` if CloudKit can’t create one.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingRecordZoneChange/initWithRecordID:type:
func (c CKSyncEnginePendingRecordZoneChange) InitWithRecordIDType(recordID ICKRecordID, type_ CKSyncEnginePendingRecordZoneChangeType) CKSyncEnginePendingRecordZoneChange {
	rv := objc.Send[CKSyncEnginePendingRecordZoneChange](c.ID, objc.Sel("initWithRecordID:type:"), recordID, type_)
	return rv
}

// The identifier of the modified record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingRecordZoneChange/recordID
func (c CKSyncEnginePendingRecordZoneChange) RecordID() ICKRecordID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordID"))
	return CKRecordIDFromID(objc.ID(rv))
}

// The type of change to make.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingRecordZoneChange/type
func (c CKSyncEnginePendingRecordZoneChange) Type() CKSyncEnginePendingRecordZoneChangeType {
	rv := objc.Send[CKSyncEnginePendingRecordZoneChangeType](c.ID, objc.Sel("type"))
	return CKSyncEnginePendingRecordZoneChangeType(rv)
}
