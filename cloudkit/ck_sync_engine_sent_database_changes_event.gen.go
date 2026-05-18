// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CKSyncEngineSentDatabaseChangesEvent] class.
var (
	_CKSyncEngineSentDatabaseChangesEventClass     CKSyncEngineSentDatabaseChangesEventClass
	_CKSyncEngineSentDatabaseChangesEventClassOnce sync.Once
)

func getCKSyncEngineSentDatabaseChangesEventClass() CKSyncEngineSentDatabaseChangesEventClass {
	_CKSyncEngineSentDatabaseChangesEventClassOnce.Do(func() {
		_CKSyncEngineSentDatabaseChangesEventClass = CKSyncEngineSentDatabaseChangesEventClass{class: objc.GetClass("CKSyncEngineSentDatabaseChangesEvent")}
	})
	return _CKSyncEngineSentDatabaseChangesEventClass
}

// GetCKSyncEngineSentDatabaseChangesEventClass returns the class object for CKSyncEngineSentDatabaseChangesEvent.
func GetCKSyncEngineSentDatabaseChangesEventClass() CKSyncEngineSentDatabaseChangesEventClass {
	return getCKSyncEngineSentDatabaseChangesEventClass()
}

type CKSyncEngineSentDatabaseChangesEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineSentDatabaseChangesEventClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineSentDatabaseChangesEventClass) Alloc() CKSyncEngineSentDatabaseChangesEvent {
	rv := objc.Send[CKSyncEngineSentDatabaseChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides information about a sent batch of database changes.
//
// # Accessing successful changes
//
//   - [CKSyncEngineSentDatabaseChangesEvent.DeletedZoneIDs]: The unique identifiers of the deleted record zones.
//   - [CKSyncEngineSentDatabaseChangesEvent.SavedZones]: The modified record zones.
//
// # Accessing failed changes
//
//   - [CKSyncEngineSentDatabaseChangesEvent.FailedZoneDeletes]: The unique identifiers of the record zones CloudKit is unable to delete, and the reasons why.
//   - [CKSyncEngineSentDatabaseChangesEvent.FailedZoneSaves]: The record zones that CloudKit is unable to modify.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentDatabaseChangesEvent
type CKSyncEngineSentDatabaseChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineSentDatabaseChangesEventFromID constructs a [CKSyncEngineSentDatabaseChangesEvent] from an objc.ID.
//
// An object that provides information about a sent batch of database changes.
func CKSyncEngineSentDatabaseChangesEventFromID(id objc.ID) CKSyncEngineSentDatabaseChangesEvent {
	return CKSyncEngineSentDatabaseChangesEvent{CKSyncEngineEvent: CKSyncEngineEventFromID(id)}
}

// Ensure CKSyncEngineSentDatabaseChangesEvent implements ICKSyncEngineSentDatabaseChangesEvent.
var _ ICKSyncEngineSentDatabaseChangesEvent = CKSyncEngineSentDatabaseChangesEvent{}

// An interface definition for the [CKSyncEngineSentDatabaseChangesEvent] class.
//
// # Accessing successful changes
//
//   - [ICKSyncEngineSentDatabaseChangesEvent.DeletedZoneIDs]: The unique identifiers of the deleted record zones.
//   - [ICKSyncEngineSentDatabaseChangesEvent.SavedZones]: The modified record zones.
//
// # Accessing failed changes
//
//   - [ICKSyncEngineSentDatabaseChangesEvent.FailedZoneDeletes]: The unique identifiers of the record zones CloudKit is unable to delete, and the reasons why.
//   - [ICKSyncEngineSentDatabaseChangesEvent.FailedZoneSaves]: The record zones that CloudKit is unable to modify.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentDatabaseChangesEvent
type ICKSyncEngineSentDatabaseChangesEvent interface {
	ICKSyncEngineEvent

	// Topic: Accessing successful changes

	// The unique identifiers of the deleted record zones.
	DeletedZoneIDs() []CKRecordZoneID
	// The modified record zones.
	SavedZones() []CKRecordZone

	// Topic: Accessing failed changes

	// The unique identifiers of the record zones CloudKit is unable to delete, and the reasons why.
	FailedZoneDeletes() foundation.INSDictionary
	// The record zones that CloudKit is unable to modify.
	FailedZoneSaves() []CKSyncEngineFailedZoneSave
}

// Init initializes the instance.
func (c CKSyncEngineSentDatabaseChangesEvent) Init() CKSyncEngineSentDatabaseChangesEvent {
	rv := objc.Send[CKSyncEngineSentDatabaseChangesEvent](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineSentDatabaseChangesEvent) Autorelease() CKSyncEngineSentDatabaseChangesEvent {
	rv := objc.Send[CKSyncEngineSentDatabaseChangesEvent](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineSentDatabaseChangesEvent creates a new CKSyncEngineSentDatabaseChangesEvent instance.
func NewCKSyncEngineSentDatabaseChangesEvent() CKSyncEngineSentDatabaseChangesEvent {
	class := getCKSyncEngineSentDatabaseChangesEventClass()
	rv := objc.Send[CKSyncEngineSentDatabaseChangesEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The unique identifiers of the deleted record zones.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentDatabaseChangesEvent/deletedZoneIDs
func (c CKSyncEngineSentDatabaseChangesEvent) DeletedZoneIDs() []CKRecordZoneID {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("deletedZoneIDs"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKRecordZoneID {
		return CKRecordZoneIDFromID(id)
	})
}

// The modified record zones.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentDatabaseChangesEvent/savedZones
func (c CKSyncEngineSentDatabaseChangesEvent) SavedZones() []CKRecordZone {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("savedZones"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKRecordZone {
		return CKRecordZoneFromID(id)
	})
}

// The unique identifiers of the record zones CloudKit is unable to delete,
// and the reasons why.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentDatabaseChangesEvent/failedZoneDeletes
func (c CKSyncEngineSentDatabaseChangesEvent) FailedZoneDeletes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("failedZoneDeletes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The record zones that CloudKit is unable to modify.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentDatabaseChangesEvent/failedZoneSaves
func (c CKSyncEngineSentDatabaseChangesEvent) FailedZoneSaves() []CKSyncEngineFailedZoneSave {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("failedZoneSaves"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKSyncEngineFailedZoneSave {
		return CKSyncEngineFailedZoneSaveFromID(id)
	})
}
