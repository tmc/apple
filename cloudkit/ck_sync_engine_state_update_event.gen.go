// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [CKSyncEngineStateUpdateEvent] class.
var (
	_CKSyncEngineStateUpdateEventClass     CKSyncEngineStateUpdateEventClass
	_CKSyncEngineStateUpdateEventClassOnce sync.Once
)

func getCKSyncEngineStateUpdateEventClass() CKSyncEngineStateUpdateEventClass {
	_CKSyncEngineStateUpdateEventClassOnce.Do(func() {
		_CKSyncEngineStateUpdateEventClass = CKSyncEngineStateUpdateEventClass{class: objc.GetClass("CKSyncEngineStateUpdateEvent")}
	})
	return _CKSyncEngineStateUpdateEventClass
}

// GetCKSyncEngineStateUpdateEventClass returns the class object for CKSyncEngineStateUpdateEvent.
func GetCKSyncEngineStateUpdateEventClass() CKSyncEngineStateUpdateEventClass {
	return getCKSyncEngineStateUpdateEventClass()
}

type CKSyncEngineStateUpdateEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineStateUpdateEventClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineStateUpdateEventClass) Alloc() CKSyncEngineStateUpdateEvent {
	rv := objc.Send[CKSyncEngineStateUpdateEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The sync engine state was updated, and you should persist it locally.
//
// # Overview
//
// In order to function properly and efficiently, [CKSyncEngine] tracks some
// state internally. When the sync engine state changes, it gives you the
// latest serialized version in a [CKSyncEngine.Event.StateUpdate]. This event
// happens occasionally when the sync engine modifies the state internally
// during normal sync operation. This event also happens when you change the
// state yourself.
//
// The sync engine does not persist this state to disk, so you need to persist
// it in alongside your own local data. The next time your process launches,
// use this latest state serialization in [CKSyncEngineStateUpdateEvent.StateSerialization] to initialize
// your sync engine.
//
// This state is directly tied to the changes you fetch and send with the sync
// engine. You should persist this state alongside any changes fetched prior
// to receiving this state.
//
// # Accessing the state
//
//   - [CKSyncEngineStateUpdateEvent.StateSerialization]: The current state of the sync engine.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineStateUpdateEvent
//
// [CKSyncEngine.Event.StateUpdate]: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-5sie5/Event/StateUpdate
type CKSyncEngineStateUpdateEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineStateUpdateEventFromID constructs a [CKSyncEngineStateUpdateEvent] from an objc.ID.
//
// The sync engine state was updated, and you should persist it locally.
func CKSyncEngineStateUpdateEventFromID(id objc.ID) CKSyncEngineStateUpdateEvent {
	return CKSyncEngineStateUpdateEvent{CKSyncEngineEvent: CKSyncEngineEventFromID(id)}
}

// Ensure CKSyncEngineStateUpdateEvent implements ICKSyncEngineStateUpdateEvent.
var _ ICKSyncEngineStateUpdateEvent = CKSyncEngineStateUpdateEvent{}

// An interface definition for the [CKSyncEngineStateUpdateEvent] class.
//
// # Accessing the state
//
//   - [ICKSyncEngineStateUpdateEvent.StateSerialization]: The current state of the sync engine.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineStateUpdateEvent
type ICKSyncEngineStateUpdateEvent interface {
	ICKSyncEngineEvent

	// Topic: Accessing the state

	// The current state of the sync engine.
	StateSerialization() ICKSyncEngineStateSerialization
}

// Init initializes the instance.
func (c CKSyncEngineStateUpdateEvent) Init() CKSyncEngineStateUpdateEvent {
	rv := objc.Send[CKSyncEngineStateUpdateEvent](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineStateUpdateEvent) Autorelease() CKSyncEngineStateUpdateEvent {
	rv := objc.Send[CKSyncEngineStateUpdateEvent](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineStateUpdateEvent creates a new CKSyncEngineStateUpdateEvent instance.
func NewCKSyncEngineStateUpdateEvent() CKSyncEngineStateUpdateEvent {
	class := getCKSyncEngineStateUpdateEventClass()
	rv := objc.Send[CKSyncEngineStateUpdateEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The current state of the sync engine.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineStateUpdateEvent/stateSerialization
func (c CKSyncEngineStateUpdateEvent) StateSerialization() ICKSyncEngineStateSerialization {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("stateSerialization"))
	return CKSyncEngineStateSerializationFromID(objc.ID(rv))
}
