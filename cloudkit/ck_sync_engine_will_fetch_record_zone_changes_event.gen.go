// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [CKSyncEngineWillFetchRecordZoneChangesEvent] class.
var (
	_CKSyncEngineWillFetchRecordZoneChangesEventClass     CKSyncEngineWillFetchRecordZoneChangesEventClass
	_CKSyncEngineWillFetchRecordZoneChangesEventClassOnce sync.Once
)

func getCKSyncEngineWillFetchRecordZoneChangesEventClass() CKSyncEngineWillFetchRecordZoneChangesEventClass {
	_CKSyncEngineWillFetchRecordZoneChangesEventClassOnce.Do(func() {
		_CKSyncEngineWillFetchRecordZoneChangesEventClass = CKSyncEngineWillFetchRecordZoneChangesEventClass{class: objc.GetClass("CKSyncEngineWillFetchRecordZoneChangesEvent")}
	})
	return _CKSyncEngineWillFetchRecordZoneChangesEventClass
}

// GetCKSyncEngineWillFetchRecordZoneChangesEventClass returns the class object for CKSyncEngineWillFetchRecordZoneChangesEvent.
func GetCKSyncEngineWillFetchRecordZoneChangesEventClass() CKSyncEngineWillFetchRecordZoneChangesEventClass {
	return getCKSyncEngineWillFetchRecordZoneChangesEventClass()
}

type CKSyncEngineWillFetchRecordZoneChangesEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineWillFetchRecordZoneChangesEventClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineWillFetchRecordZoneChangesEventClass) Alloc() CKSyncEngineWillFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchRecordZoneChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The sync engine finished fetching record zone changes from the server for a
// specific zone.
//
// # Overview
//
// This might be a good signal to perform any post-processing tasks on a
// per-zone basis if necessary.
//
// You should receive one [CKSyncEngineDidFetchRecordZoneChangesEvent] for
// each [CKSyncEngineWillFetchRecordZoneChangesEvent].
//
// # Identifying the record zone
//
//   - [CKSyncEngineWillFetchRecordZoneChangesEvent.ZoneID]: The associated record zone’s unique identifier.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineWillFetchRecordZoneChangesEvent
type CKSyncEngineWillFetchRecordZoneChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineWillFetchRecordZoneChangesEventFromID constructs a [CKSyncEngineWillFetchRecordZoneChangesEvent] from an objc.ID.
//
// The sync engine finished fetching record zone changes from the server for a
// specific zone.
func CKSyncEngineWillFetchRecordZoneChangesEventFromID(id objc.ID) CKSyncEngineWillFetchRecordZoneChangesEvent {
	return CKSyncEngineWillFetchRecordZoneChangesEvent{CKSyncEngineEvent: CKSyncEngineEventFromID(id)}
}

// Ensure CKSyncEngineWillFetchRecordZoneChangesEvent implements ICKSyncEngineWillFetchRecordZoneChangesEvent.
var _ ICKSyncEngineWillFetchRecordZoneChangesEvent = CKSyncEngineWillFetchRecordZoneChangesEvent{}

// An interface definition for the [CKSyncEngineWillFetchRecordZoneChangesEvent] class.
//
// # Identifying the record zone
//
//   - [ICKSyncEngineWillFetchRecordZoneChangesEvent.ZoneID]: The associated record zone’s unique identifier.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineWillFetchRecordZoneChangesEvent
type ICKSyncEngineWillFetchRecordZoneChangesEvent interface {
	ICKSyncEngineEvent

	// Topic: Identifying the record zone

	// The associated record zone’s unique identifier.
	ZoneID() ICKRecordZoneID
}

// Init initializes the instance.
func (c CKSyncEngineWillFetchRecordZoneChangesEvent) Init() CKSyncEngineWillFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchRecordZoneChangesEvent](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineWillFetchRecordZoneChangesEvent) Autorelease() CKSyncEngineWillFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchRecordZoneChangesEvent](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineWillFetchRecordZoneChangesEvent creates a new CKSyncEngineWillFetchRecordZoneChangesEvent instance.
func NewCKSyncEngineWillFetchRecordZoneChangesEvent() CKSyncEngineWillFetchRecordZoneChangesEvent {
	class := getCKSyncEngineWillFetchRecordZoneChangesEventClass()
	rv := objc.Send[CKSyncEngineWillFetchRecordZoneChangesEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The associated record zone’s unique identifier.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineWillFetchRecordZoneChangesEvent/zoneID
func (c CKSyncEngineWillFetchRecordZoneChangesEvent) ZoneID() ICKRecordZoneID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("zoneID"))
	return CKRecordZoneIDFromID(objc.ID(rv))
}
