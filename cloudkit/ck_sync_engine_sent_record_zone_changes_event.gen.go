// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CKSyncEngineSentRecordZoneChangesEvent] class.
var (
	_CKSyncEngineSentRecordZoneChangesEventClass     CKSyncEngineSentRecordZoneChangesEventClass
	_CKSyncEngineSentRecordZoneChangesEventClassOnce sync.Once
)

func getCKSyncEngineSentRecordZoneChangesEventClass() CKSyncEngineSentRecordZoneChangesEventClass {
	_CKSyncEngineSentRecordZoneChangesEventClassOnce.Do(func() {
		_CKSyncEngineSentRecordZoneChangesEventClass = CKSyncEngineSentRecordZoneChangesEventClass{class: objc.GetClass("CKSyncEngineSentRecordZoneChangesEvent")}
	})
	return _CKSyncEngineSentRecordZoneChangesEventClass
}

// GetCKSyncEngineSentRecordZoneChangesEventClass returns the class object for CKSyncEngineSentRecordZoneChangesEvent.
func GetCKSyncEngineSentRecordZoneChangesEventClass() CKSyncEngineSentRecordZoneChangesEventClass {
	return getCKSyncEngineSentRecordZoneChangesEventClass()
}

type CKSyncEngineSentRecordZoneChangesEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineSentRecordZoneChangesEventClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineSentRecordZoneChangesEventClass) Alloc() CKSyncEngineSentRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineSentRecordZoneChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The sync engine finished sending a batch of record zone changes to the
// server.
//
// # Overview
//
// If a record save succeeded, you should encode the system fields of this
// record to use the next time you save. See
// [CKRecord.EncodeSystemFieldsWithCoder].
//
// If a record deletion succeeded, you should remove any local system fields
// for that record.
//
// If the record change failed, try to resolve the issue causing the error and
// save the record again if necessary.
//
// # Accessing successful changes
//
//   - [CKSyncEngineSentRecordZoneChangesEvent.DeletedRecordIDs]: The unique identifiers of the deleted records.
//   - [CKSyncEngineSentRecordZoneChangesEvent.SavedRecords]: The modified records.
//
// # Accessing failed changes
//
//   - [CKSyncEngineSentRecordZoneChangesEvent.FailedRecordDeletes]: The unique identifiers of the records CloudKit is unable to delete, and the reasons why.
//   - [CKSyncEngineSentRecordZoneChangesEvent.FailedRecordSaves]: The records that CloudKit is unable to modify.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentRecordZoneChangesEvent
type CKSyncEngineSentRecordZoneChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineSentRecordZoneChangesEventFromID constructs a [CKSyncEngineSentRecordZoneChangesEvent] from an objc.ID.
//
// The sync engine finished sending a batch of record zone changes to the
// server.
func CKSyncEngineSentRecordZoneChangesEventFromID(id objc.ID) CKSyncEngineSentRecordZoneChangesEvent {
	return CKSyncEngineSentRecordZoneChangesEvent{CKSyncEngineEvent: CKSyncEngineEventFromID(id)}
}

// Ensure CKSyncEngineSentRecordZoneChangesEvent implements ICKSyncEngineSentRecordZoneChangesEvent.
var _ ICKSyncEngineSentRecordZoneChangesEvent = CKSyncEngineSentRecordZoneChangesEvent{}

// An interface definition for the [CKSyncEngineSentRecordZoneChangesEvent] class.
//
// # Accessing successful changes
//
//   - [ICKSyncEngineSentRecordZoneChangesEvent.DeletedRecordIDs]: The unique identifiers of the deleted records.
//   - [ICKSyncEngineSentRecordZoneChangesEvent.SavedRecords]: The modified records.
//
// # Accessing failed changes
//
//   - [ICKSyncEngineSentRecordZoneChangesEvent.FailedRecordDeletes]: The unique identifiers of the records CloudKit is unable to delete, and the reasons why.
//   - [ICKSyncEngineSentRecordZoneChangesEvent.FailedRecordSaves]: The records that CloudKit is unable to modify.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentRecordZoneChangesEvent
type ICKSyncEngineSentRecordZoneChangesEvent interface {
	ICKSyncEngineEvent

	// Topic: Accessing successful changes

	// The unique identifiers of the deleted records.
	DeletedRecordIDs() []CKRecordID
	// The modified records.
	SavedRecords() []CKRecord

	// Topic: Accessing failed changes

	// The unique identifiers of the records CloudKit is unable to delete, and the reasons why.
	FailedRecordDeletes() foundation.INSDictionary
	// The records that CloudKit is unable to modify.
	FailedRecordSaves() []CKSyncEngineFailedRecordSave
}

// Init initializes the instance.
func (c CKSyncEngineSentRecordZoneChangesEvent) Init() CKSyncEngineSentRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineSentRecordZoneChangesEvent](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineSentRecordZoneChangesEvent) Autorelease() CKSyncEngineSentRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineSentRecordZoneChangesEvent](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineSentRecordZoneChangesEvent creates a new CKSyncEngineSentRecordZoneChangesEvent instance.
func NewCKSyncEngineSentRecordZoneChangesEvent() CKSyncEngineSentRecordZoneChangesEvent {
	class := getCKSyncEngineSentRecordZoneChangesEventClass()
	rv := objc.Send[CKSyncEngineSentRecordZoneChangesEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The unique identifiers of the deleted records.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentRecordZoneChangesEvent/deletedRecordIDs
func (c CKSyncEngineSentRecordZoneChangesEvent) DeletedRecordIDs() []CKRecordID {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("deletedRecordIDs"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKRecordID {
		return CKRecordIDFromID(id)
	})
}

// The modified records.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentRecordZoneChangesEvent/savedRecords
func (c CKSyncEngineSentRecordZoneChangesEvent) SavedRecords() []CKRecord {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("savedRecords"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKRecord {
		return CKRecordFromID(id)
	})
}

// The unique identifiers of the records CloudKit is unable to delete, and the
// reasons why.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentRecordZoneChangesEvent/failedRecordDeletes
func (c CKSyncEngineSentRecordZoneChangesEvent) FailedRecordDeletes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("failedRecordDeletes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The records that CloudKit is unable to modify.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentRecordZoneChangesEvent/failedRecordSaves
func (c CKSyncEngineSentRecordZoneChangesEvent) FailedRecordSaves() []CKSyncEngineFailedRecordSave {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("failedRecordSaves"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKSyncEngineFailedRecordSave {
		return CKSyncEngineFailedRecordSaveFromID(id)
	})
}
