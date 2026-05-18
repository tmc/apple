// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [CKSyncEngineFetchedDatabaseChangesEvent] class.
var (
	_CKSyncEngineFetchedDatabaseChangesEventClass     CKSyncEngineFetchedDatabaseChangesEventClass
	_CKSyncEngineFetchedDatabaseChangesEventClassOnce sync.Once
)

func getCKSyncEngineFetchedDatabaseChangesEventClass() CKSyncEngineFetchedDatabaseChangesEventClass {
	_CKSyncEngineFetchedDatabaseChangesEventClassOnce.Do(func() {
		_CKSyncEngineFetchedDatabaseChangesEventClass = CKSyncEngineFetchedDatabaseChangesEventClass{class: objc.GetClass("CKSyncEngineFetchedDatabaseChangesEvent")}
	})
	return _CKSyncEngineFetchedDatabaseChangesEventClass
}

// GetCKSyncEngineFetchedDatabaseChangesEventClass returns the class object for CKSyncEngineFetchedDatabaseChangesEvent.
func GetCKSyncEngineFetchedDatabaseChangesEventClass() CKSyncEngineFetchedDatabaseChangesEventClass {
	return getCKSyncEngineFetchedDatabaseChangesEventClass()
}

type CKSyncEngineFetchedDatabaseChangesEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineFetchedDatabaseChangesEventClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineFetchedDatabaseChangesEventClass) Alloc() CKSyncEngineFetchedDatabaseChangesEvent {
	rv := objc.Send[CKSyncEngineFetchedDatabaseChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides information about fetched database changes.
//
// # Overview
//
// # Accessing changes
//
//   - [CKSyncEngineFetchedDatabaseChangesEvent.Deletions]: The deleted record zones.
//   - [CKSyncEngineFetchedDatabaseChangesEvent.Modifications]: The modified record zones.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedDatabaseChangesEvent
type CKSyncEngineFetchedDatabaseChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineFetchedDatabaseChangesEventFromID constructs a [CKSyncEngineFetchedDatabaseChangesEvent] from an objc.ID.
//
// An object that provides information about fetched database changes.
func CKSyncEngineFetchedDatabaseChangesEventFromID(id objc.ID) CKSyncEngineFetchedDatabaseChangesEvent {
	return CKSyncEngineFetchedDatabaseChangesEvent{CKSyncEngineEvent: CKSyncEngineEventFromID(id)}
}

// Ensure CKSyncEngineFetchedDatabaseChangesEvent implements ICKSyncEngineFetchedDatabaseChangesEvent.
var _ ICKSyncEngineFetchedDatabaseChangesEvent = CKSyncEngineFetchedDatabaseChangesEvent{}

// An interface definition for the [CKSyncEngineFetchedDatabaseChangesEvent] class.
//
// # Accessing changes
//
//   - [ICKSyncEngineFetchedDatabaseChangesEvent.Deletions]: The deleted record zones.
//   - [ICKSyncEngineFetchedDatabaseChangesEvent.Modifications]: The modified record zones.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedDatabaseChangesEvent
type ICKSyncEngineFetchedDatabaseChangesEvent interface {
	ICKSyncEngineEvent

	// Topic: Accessing changes

	// The deleted record zones.
	Deletions() []CKSyncEngineFetchedZoneDeletion
	// The modified record zones.
	Modifications() []CKRecordZone
}

// Init initializes the instance.
func (c CKSyncEngineFetchedDatabaseChangesEvent) Init() CKSyncEngineFetchedDatabaseChangesEvent {
	rv := objc.Send[CKSyncEngineFetchedDatabaseChangesEvent](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineFetchedDatabaseChangesEvent) Autorelease() CKSyncEngineFetchedDatabaseChangesEvent {
	rv := objc.Send[CKSyncEngineFetchedDatabaseChangesEvent](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFetchedDatabaseChangesEvent creates a new CKSyncEngineFetchedDatabaseChangesEvent instance.
func NewCKSyncEngineFetchedDatabaseChangesEvent() CKSyncEngineFetchedDatabaseChangesEvent {
	class := getCKSyncEngineFetchedDatabaseChangesEventClass()
	rv := objc.Send[CKSyncEngineFetchedDatabaseChangesEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The deleted record zones.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedDatabaseChangesEvent/deletions
func (c CKSyncEngineFetchedDatabaseChangesEvent) Deletions() []CKSyncEngineFetchedZoneDeletion {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("deletions"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKSyncEngineFetchedZoneDeletion {
		return CKSyncEngineFetchedZoneDeletionFromID(id)
	})
}

// The modified record zones.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedDatabaseChangesEvent/modifications
func (c CKSyncEngineFetchedDatabaseChangesEvent) Modifications() []CKRecordZone {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("modifications"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKRecordZone {
		return CKRecordZoneFromID(id)
	})
}
