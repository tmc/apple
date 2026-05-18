// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [CKSyncEngineFetchedRecordZoneChangesEvent] class.
var (
	_CKSyncEngineFetchedRecordZoneChangesEventClass     CKSyncEngineFetchedRecordZoneChangesEventClass
	_CKSyncEngineFetchedRecordZoneChangesEventClassOnce sync.Once
)

func getCKSyncEngineFetchedRecordZoneChangesEventClass() CKSyncEngineFetchedRecordZoneChangesEventClass {
	_CKSyncEngineFetchedRecordZoneChangesEventClassOnce.Do(func() {
		_CKSyncEngineFetchedRecordZoneChangesEventClass = CKSyncEngineFetchedRecordZoneChangesEventClass{class: objc.GetClass("CKSyncEngineFetchedRecordZoneChangesEvent")}
	})
	return _CKSyncEngineFetchedRecordZoneChangesEventClass
}

// GetCKSyncEngineFetchedRecordZoneChangesEventClass returns the class object for CKSyncEngineFetchedRecordZoneChangesEvent.
func GetCKSyncEngineFetchedRecordZoneChangesEventClass() CKSyncEngineFetchedRecordZoneChangesEventClass {
	return getCKSyncEngineFetchedRecordZoneChangesEventClass()
}

type CKSyncEngineFetchedRecordZoneChangesEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineFetchedRecordZoneChangesEventClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineFetchedRecordZoneChangesEventClass) Alloc() CKSyncEngineFetchedRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineFetchedRecordZoneChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides information about fetched record zone changes.
//
// # Overview
//
// # Accessing changes
//
//   - [CKSyncEngineFetchedRecordZoneChangesEvent.Deletions]: The deleted records.
//   - [CKSyncEngineFetchedRecordZoneChangesEvent.Modifications]: The modified records.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedRecordZoneChangesEvent
type CKSyncEngineFetchedRecordZoneChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineFetchedRecordZoneChangesEventFromID constructs a [CKSyncEngineFetchedRecordZoneChangesEvent] from an objc.ID.
//
// An object that provides information about fetched record zone changes.
func CKSyncEngineFetchedRecordZoneChangesEventFromID(id objc.ID) CKSyncEngineFetchedRecordZoneChangesEvent {
	return CKSyncEngineFetchedRecordZoneChangesEvent{CKSyncEngineEvent: CKSyncEngineEventFromID(id)}
}

// Ensure CKSyncEngineFetchedRecordZoneChangesEvent implements ICKSyncEngineFetchedRecordZoneChangesEvent.
var _ ICKSyncEngineFetchedRecordZoneChangesEvent = CKSyncEngineFetchedRecordZoneChangesEvent{}

// An interface definition for the [CKSyncEngineFetchedRecordZoneChangesEvent] class.
//
// # Accessing changes
//
//   - [ICKSyncEngineFetchedRecordZoneChangesEvent.Deletions]: The deleted records.
//   - [ICKSyncEngineFetchedRecordZoneChangesEvent.Modifications]: The modified records.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedRecordZoneChangesEvent
type ICKSyncEngineFetchedRecordZoneChangesEvent interface {
	ICKSyncEngineEvent

	// Topic: Accessing changes

	// The deleted records.
	Deletions() []CKSyncEngineFetchedRecordDeletion
	// The modified records.
	Modifications() []CKRecord
}

// Init initializes the instance.
func (c CKSyncEngineFetchedRecordZoneChangesEvent) Init() CKSyncEngineFetchedRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineFetchedRecordZoneChangesEvent](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineFetchedRecordZoneChangesEvent) Autorelease() CKSyncEngineFetchedRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineFetchedRecordZoneChangesEvent](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFetchedRecordZoneChangesEvent creates a new CKSyncEngineFetchedRecordZoneChangesEvent instance.
func NewCKSyncEngineFetchedRecordZoneChangesEvent() CKSyncEngineFetchedRecordZoneChangesEvent {
	class := getCKSyncEngineFetchedRecordZoneChangesEventClass()
	rv := objc.Send[CKSyncEngineFetchedRecordZoneChangesEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The deleted records.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedRecordZoneChangesEvent/deletions
func (c CKSyncEngineFetchedRecordZoneChangesEvent) Deletions() []CKSyncEngineFetchedRecordDeletion {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("deletions"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKSyncEngineFetchedRecordDeletion {
		return CKSyncEngineFetchedRecordDeletionFromID(id)
	})
}

// The modified records.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedRecordZoneChangesEvent/modifications
func (c CKSyncEngineFetchedRecordZoneChangesEvent) Modifications() []CKRecord {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("modifications"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKRecord {
		return CKRecordFromID(id)
	})
}
