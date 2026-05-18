// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKSyncEngineEvent] class.
var (
	_CKSyncEngineEventClass     CKSyncEngineEventClass
	_CKSyncEngineEventClassOnce sync.Once
)

func getCKSyncEngineEventClass() CKSyncEngineEventClass {
	_CKSyncEngineEventClassOnce.Do(func() {
		_CKSyncEngineEventClass = CKSyncEngineEventClass{class: objc.GetClass("CKSyncEngineEvent")}
	})
	return _CKSyncEngineEventClass
}

// GetCKSyncEngineEventClass returns the class object for CKSyncEngineEvent.
func GetCKSyncEngineEventClass() CKSyncEngineEventClass {
	return getCKSyncEngineEventClass()
}

type CKSyncEngineEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineEventClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineEventClass) Alloc() CKSyncEngineEvent {
	rv := objc.Send[CKSyncEngineEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An event that occurs during a sync operation.
//
// # Overview
//
// All sync operation events descend from this base class, and as such you
// don’t create instances of it directly. Instead, the sync engine
// dispatches them to your app’s delegate, periodically, throughout a sync
// operation.
//
// See the documentation for each event class for more details about when and
// why an event is posted.
//
// Use the [CKSyncEngineEvent.Type] property to determine the event’s proper type, and then
// use the corresponding convenience property to retrieve a reference to the
// event that’s downcast to the appropriate subclass. For example, when
// CloudKit vends an event with [CKSyncEngineEvent.Type] set to
// [CKSyncEngineEventType.stateUpdate], use the [CKSyncEngineEvent.StateUpdateEvent] property to
// get the downcast reference.
//
// # Determining the type
//
//   - [CKSyncEngineEvent.Type]: The type of event.
//
// # Accessing account changes
//
//   - [CKSyncEngineEvent.AccountChangeEvent]: The event downcast to the subclass that represents a change to the device’s iCloud account.
//
// # Accessing fetch events
//
//   - [CKSyncEngineEvent.WillFetchChangesEvent]: The event downcast to the subclass that represents an imminent database fetch.
//   - [CKSyncEngineEvent.WillFetchRecordZoneChangesEvent]: The event downcast to the subclass that represents an imminent fetch of record zone changes.
//   - [CKSyncEngineEvent.FetchedDatabaseChangesEvent]: The event downcast to the subclass that represents a set of fetched database changes to process.
//   - [CKSyncEngineEvent.FetchedRecordZoneChangesEvent]: The event downcast to the subclass that represents a set of fetched record zone changes to process.
//   - [CKSyncEngineEvent.DidFetchRecordZoneChangesEvent]: The event downcast to the subclass that represents a completed record zone fetch.
//   - [CKSyncEngineEvent.DidFetchChangesEvent]: The event downcast to the subclass that represents a completed database fetch.
//
// # Accessing send events
//
//   - [CKSyncEngineEvent.WillSendChangesEvent]: The event downcast to the subclass that represents an imminent send operation.
//   - [CKSyncEngineEvent.SentDatabaseChangesEvent]: The event downcast to the subclass that represents a sent batch of database changes.
//   - [CKSyncEngineEvent.SentRecordZoneChangesEvent]: The event downcast to the subclass that represents a sent batch of record zone changes.
//   - [CKSyncEngineEvent.DidSendChangesEvent]: The event downcast to the subclass that represents a completed send operation.
//
// # Accessing state updates
//
//   - [CKSyncEngineEvent.StateUpdateEvent]: The event downcast to the subclass that represents an update to the sync engine’s state.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent
//
// [CKSyncEngineEventType.stateUpdate]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEventType/stateUpdate
type CKSyncEngineEvent struct {
	objectivec.Object
}

// CKSyncEngineEventFromID constructs a [CKSyncEngineEvent] from an objc.ID.
//
// An event that occurs during a sync operation.
func CKSyncEngineEventFromID(id objc.ID) CKSyncEngineEvent {
	return CKSyncEngineEvent{objectivec.Object{ID: id}}
}

// Ensure CKSyncEngineEvent implements ICKSyncEngineEvent.
var _ ICKSyncEngineEvent = CKSyncEngineEvent{}

// An interface definition for the [CKSyncEngineEvent] class.
//
// # Determining the type
//
//   - [ICKSyncEngineEvent.Type]: The type of event.
//
// # Accessing account changes
//
//   - [ICKSyncEngineEvent.AccountChangeEvent]: The event downcast to the subclass that represents a change to the device’s iCloud account.
//
// # Accessing fetch events
//
//   - [ICKSyncEngineEvent.WillFetchChangesEvent]: The event downcast to the subclass that represents an imminent database fetch.
//   - [ICKSyncEngineEvent.WillFetchRecordZoneChangesEvent]: The event downcast to the subclass that represents an imminent fetch of record zone changes.
//   - [ICKSyncEngineEvent.FetchedDatabaseChangesEvent]: The event downcast to the subclass that represents a set of fetched database changes to process.
//   - [ICKSyncEngineEvent.FetchedRecordZoneChangesEvent]: The event downcast to the subclass that represents a set of fetched record zone changes to process.
//   - [ICKSyncEngineEvent.DidFetchRecordZoneChangesEvent]: The event downcast to the subclass that represents a completed record zone fetch.
//   - [ICKSyncEngineEvent.DidFetchChangesEvent]: The event downcast to the subclass that represents a completed database fetch.
//
// # Accessing send events
//
//   - [ICKSyncEngineEvent.WillSendChangesEvent]: The event downcast to the subclass that represents an imminent send operation.
//   - [ICKSyncEngineEvent.SentDatabaseChangesEvent]: The event downcast to the subclass that represents a sent batch of database changes.
//   - [ICKSyncEngineEvent.SentRecordZoneChangesEvent]: The event downcast to the subclass that represents a sent batch of record zone changes.
//   - [ICKSyncEngineEvent.DidSendChangesEvent]: The event downcast to the subclass that represents a completed send operation.
//
// # Accessing state updates
//
//   - [ICKSyncEngineEvent.StateUpdateEvent]: The event downcast to the subclass that represents an update to the sync engine’s state.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent
type ICKSyncEngineEvent interface {
	objectivec.IObject

	// Topic: Determining the type

	// The type of event.
	Type() CKSyncEngineEventType

	// Topic: Accessing account changes

	// The event downcast to the subclass that represents a change to the device’s iCloud account.
	AccountChangeEvent() ICKSyncEngineAccountChangeEvent

	// Topic: Accessing fetch events

	// The event downcast to the subclass that represents an imminent database fetch.
	WillFetchChangesEvent() ICKSyncEngineWillFetchChangesEvent
	// The event downcast to the subclass that represents an imminent fetch of record zone changes.
	WillFetchRecordZoneChangesEvent() ICKSyncEngineWillFetchRecordZoneChangesEvent
	// The event downcast to the subclass that represents a set of fetched database changes to process.
	FetchedDatabaseChangesEvent() ICKSyncEngineFetchedDatabaseChangesEvent
	// The event downcast to the subclass that represents a set of fetched record zone changes to process.
	FetchedRecordZoneChangesEvent() ICKSyncEngineFetchedRecordZoneChangesEvent
	// The event downcast to the subclass that represents a completed record zone fetch.
	DidFetchRecordZoneChangesEvent() ICKSyncEngineDidFetchRecordZoneChangesEvent
	// The event downcast to the subclass that represents a completed database fetch.
	DidFetchChangesEvent() ICKSyncEngineDidFetchChangesEvent

	// Topic: Accessing send events

	// The event downcast to the subclass that represents an imminent send operation.
	WillSendChangesEvent() ICKSyncEngineWillSendChangesEvent
	// The event downcast to the subclass that represents a sent batch of database changes.
	SentDatabaseChangesEvent() ICKSyncEngineSentDatabaseChangesEvent
	// The event downcast to the subclass that represents a sent batch of record zone changes.
	SentRecordZoneChangesEvent() ICKSyncEngineSentRecordZoneChangesEvent
	// The event downcast to the subclass that represents a completed send operation.
	DidSendChangesEvent() ICKSyncEngineDidSendChangesEvent

	// Topic: Accessing state updates

	// The event downcast to the subclass that represents an update to the sync engine’s state.
	StateUpdateEvent() ICKSyncEngineStateUpdateEvent
}

// Init initializes the instance.
func (c CKSyncEngineEvent) Init() CKSyncEngineEvent {
	rv := objc.Send[CKSyncEngineEvent](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineEvent) Autorelease() CKSyncEngineEvent {
	rv := objc.Send[CKSyncEngineEvent](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineEvent creates a new CKSyncEngineEvent instance.
func NewCKSyncEngineEvent() CKSyncEngineEvent {
	class := getCKSyncEngineEventClass()
	rv := objc.Send[CKSyncEngineEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The type of event.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/type
func (c CKSyncEngineEvent) Type() CKSyncEngineEventType {
	rv := objc.Send[CKSyncEngineEventType](c.ID, objc.Sel("type"))
	return CKSyncEngineEventType(rv)
}

// The event downcast to the subclass that represents a change to the
// device’s iCloud account.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/accountChangeEvent
func (c CKSyncEngineEvent) AccountChangeEvent() ICKSyncEngineAccountChangeEvent {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("accountChangeEvent"))
	return CKSyncEngineAccountChangeEventFromID(objc.ID(rv))
}

// The event downcast to the subclass that represents an imminent database
// fetch.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/willFetchChangesEvent
func (c CKSyncEngineEvent) WillFetchChangesEvent() ICKSyncEngineWillFetchChangesEvent {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("willFetchChangesEvent"))
	return CKSyncEngineWillFetchChangesEventFromID(objc.ID(rv))
}

// The event downcast to the subclass that represents an imminent fetch of
// record zone changes.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/willFetchRecordZoneChangesEvent
func (c CKSyncEngineEvent) WillFetchRecordZoneChangesEvent() ICKSyncEngineWillFetchRecordZoneChangesEvent {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("willFetchRecordZoneChangesEvent"))
	return CKSyncEngineWillFetchRecordZoneChangesEventFromID(objc.ID(rv))
}

// The event downcast to the subclass that represents a set of fetched
// database changes to process.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/fetchedDatabaseChangesEvent
func (c CKSyncEngineEvent) FetchedDatabaseChangesEvent() ICKSyncEngineFetchedDatabaseChangesEvent {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("fetchedDatabaseChangesEvent"))
	return CKSyncEngineFetchedDatabaseChangesEventFromID(objc.ID(rv))
}

// The event downcast to the subclass that represents a set of fetched record
// zone changes to process.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/fetchedRecordZoneChangesEvent
func (c CKSyncEngineEvent) FetchedRecordZoneChangesEvent() ICKSyncEngineFetchedRecordZoneChangesEvent {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("fetchedRecordZoneChangesEvent"))
	return CKSyncEngineFetchedRecordZoneChangesEventFromID(objc.ID(rv))
}

// The event downcast to the subclass that represents a completed record zone
// fetch.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/didFetchRecordZoneChangesEvent
func (c CKSyncEngineEvent) DidFetchRecordZoneChangesEvent() ICKSyncEngineDidFetchRecordZoneChangesEvent {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("didFetchRecordZoneChangesEvent"))
	return CKSyncEngineDidFetchRecordZoneChangesEventFromID(objc.ID(rv))
}

// The event downcast to the subclass that represents a completed database
// fetch.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/didFetchChangesEvent
func (c CKSyncEngineEvent) DidFetchChangesEvent() ICKSyncEngineDidFetchChangesEvent {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("didFetchChangesEvent"))
	return CKSyncEngineDidFetchChangesEventFromID(objc.ID(rv))
}

// The event downcast to the subclass that represents an imminent send
// operation.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/willSendChangesEvent
func (c CKSyncEngineEvent) WillSendChangesEvent() ICKSyncEngineWillSendChangesEvent {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("willSendChangesEvent"))
	return CKSyncEngineWillSendChangesEventFromID(objc.ID(rv))
}

// The event downcast to the subclass that represents a sent batch of database
// changes.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/sentDatabaseChangesEvent
func (c CKSyncEngineEvent) SentDatabaseChangesEvent() ICKSyncEngineSentDatabaseChangesEvent {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("sentDatabaseChangesEvent"))
	return CKSyncEngineSentDatabaseChangesEventFromID(objc.ID(rv))
}

// The event downcast to the subclass that represents a sent batch of record
// zone changes.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/sentRecordZoneChangesEvent
func (c CKSyncEngineEvent) SentRecordZoneChangesEvent() ICKSyncEngineSentRecordZoneChangesEvent {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("sentRecordZoneChangesEvent"))
	return CKSyncEngineSentRecordZoneChangesEventFromID(objc.ID(rv))
}

// The event downcast to the subclass that represents a completed send
// operation.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/didSendChangesEvent
func (c CKSyncEngineEvent) DidSendChangesEvent() ICKSyncEngineDidSendChangesEvent {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("didSendChangesEvent"))
	return CKSyncEngineDidSendChangesEventFromID(objc.ID(rv))
}

// The event downcast to the subclass that represents an update to the sync
// engine’s state.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/stateUpdateEvent
func (c CKSyncEngineEvent) StateUpdateEvent() ICKSyncEngineStateUpdateEvent {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("stateUpdateEvent"))
	return CKSyncEngineStateUpdateEventFromID(objc.ID(rv))
}
