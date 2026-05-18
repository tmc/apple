// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CKSyncEngineDidFetchRecordZoneChangesEvent] class.
var (
	_CKSyncEngineDidFetchRecordZoneChangesEventClass     CKSyncEngineDidFetchRecordZoneChangesEventClass
	_CKSyncEngineDidFetchRecordZoneChangesEventClassOnce sync.Once
)

func getCKSyncEngineDidFetchRecordZoneChangesEventClass() CKSyncEngineDidFetchRecordZoneChangesEventClass {
	_CKSyncEngineDidFetchRecordZoneChangesEventClassOnce.Do(func() {
		_CKSyncEngineDidFetchRecordZoneChangesEventClass = CKSyncEngineDidFetchRecordZoneChangesEventClass{class: objc.GetClass("CKSyncEngineDidFetchRecordZoneChangesEvent")}
	})
	return _CKSyncEngineDidFetchRecordZoneChangesEventClass
}

// GetCKSyncEngineDidFetchRecordZoneChangesEventClass returns the class object for CKSyncEngineDidFetchRecordZoneChangesEvent.
func GetCKSyncEngineDidFetchRecordZoneChangesEventClass() CKSyncEngineDidFetchRecordZoneChangesEventClass {
	return getCKSyncEngineDidFetchRecordZoneChangesEventClass()
}

type CKSyncEngineDidFetchRecordZoneChangesEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineDidFetchRecordZoneChangesEventClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineDidFetchRecordZoneChangesEventClass) Alloc() CKSyncEngineDidFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchRecordZoneChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides information about a finished record zone fetch.
//
// # Identifying the record zone
//
//   - [CKSyncEngineDidFetchRecordZoneChangesEvent.ZoneID]: The associated record zone’s unique identifier.
//
// # Handling errors
//
//   - [CKSyncEngineDidFetchRecordZoneChangesEvent.Error]: An error that describes the cause of a failed fetch operation.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidFetchRecordZoneChangesEvent
type CKSyncEngineDidFetchRecordZoneChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineDidFetchRecordZoneChangesEventFromID constructs a [CKSyncEngineDidFetchRecordZoneChangesEvent] from an objc.ID.
//
// An object that provides information about a finished record zone fetch.
func CKSyncEngineDidFetchRecordZoneChangesEventFromID(id objc.ID) CKSyncEngineDidFetchRecordZoneChangesEvent {
	return CKSyncEngineDidFetchRecordZoneChangesEvent{CKSyncEngineEvent: CKSyncEngineEventFromID(id)}
}

// Ensure CKSyncEngineDidFetchRecordZoneChangesEvent implements ICKSyncEngineDidFetchRecordZoneChangesEvent.
var _ ICKSyncEngineDidFetchRecordZoneChangesEvent = CKSyncEngineDidFetchRecordZoneChangesEvent{}

// An interface definition for the [CKSyncEngineDidFetchRecordZoneChangesEvent] class.
//
// # Identifying the record zone
//
//   - [ICKSyncEngineDidFetchRecordZoneChangesEvent.ZoneID]: The associated record zone’s unique identifier.
//
// # Handling errors
//
//   - [ICKSyncEngineDidFetchRecordZoneChangesEvent.Error]: An error that describes the cause of a failed fetch operation.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidFetchRecordZoneChangesEvent
type ICKSyncEngineDidFetchRecordZoneChangesEvent interface {
	ICKSyncEngineEvent

	// Topic: Identifying the record zone

	// The associated record zone’s unique identifier.
	ZoneID() ICKRecordZoneID

	// Topic: Handling errors

	// An error that describes the cause of a failed fetch operation.
	Error() foundation.INSError
}

// Init initializes the instance.
func (c CKSyncEngineDidFetchRecordZoneChangesEvent) Init() CKSyncEngineDidFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchRecordZoneChangesEvent](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineDidFetchRecordZoneChangesEvent) Autorelease() CKSyncEngineDidFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchRecordZoneChangesEvent](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineDidFetchRecordZoneChangesEvent creates a new CKSyncEngineDidFetchRecordZoneChangesEvent instance.
func NewCKSyncEngineDidFetchRecordZoneChangesEvent() CKSyncEngineDidFetchRecordZoneChangesEvent {
	class := getCKSyncEngineDidFetchRecordZoneChangesEventClass()
	rv := objc.Send[CKSyncEngineDidFetchRecordZoneChangesEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The associated record zone’s unique identifier.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidFetchRecordZoneChangesEvent/zoneID
func (c CKSyncEngineDidFetchRecordZoneChangesEvent) ZoneID() ICKRecordZoneID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("zoneID"))
	return CKRecordZoneIDFromID(objc.ID(rv))
}

// An error that describes the cause of a failed fetch operation.
//
// # Discussion
//
// A `nil` value indicates a successful fetch.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidFetchRecordZoneChangesEvent/error
func (c CKSyncEngineDidFetchRecordZoneChangesEvent) Error() foundation.INSError {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
