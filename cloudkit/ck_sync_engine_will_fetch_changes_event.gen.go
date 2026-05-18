// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [CKSyncEngineWillFetchChangesEvent] class.
var (
	_CKSyncEngineWillFetchChangesEventClass     CKSyncEngineWillFetchChangesEventClass
	_CKSyncEngineWillFetchChangesEventClassOnce sync.Once
)

func getCKSyncEngineWillFetchChangesEventClass() CKSyncEngineWillFetchChangesEventClass {
	_CKSyncEngineWillFetchChangesEventClassOnce.Do(func() {
		_CKSyncEngineWillFetchChangesEventClass = CKSyncEngineWillFetchChangesEventClass{class: objc.GetClass("CKSyncEngineWillFetchChangesEvent")}
	})
	return _CKSyncEngineWillFetchChangesEventClass
}

// GetCKSyncEngineWillFetchChangesEventClass returns the class object for CKSyncEngineWillFetchChangesEvent.
func GetCKSyncEngineWillFetchChangesEventClass() CKSyncEngineWillFetchChangesEventClass {
	return getCKSyncEngineWillFetchChangesEventClass()
}

type CKSyncEngineWillFetchChangesEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineWillFetchChangesEventClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineWillFetchChangesEventClass) Alloc() CKSyncEngineWillFetchChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The sync engine is about to fetch changes from the server.
//
// # Overview
//
// The sync engine delivers the changes themselves via
// [CKSyncEngineFetchedDatabaseChangesEvent] and
// [CKSyncEngineFetchedRecordZoneChangesEvent].
//
// Note that this event might not always occur every time you call
// [FetchChangesWithCompletionHandler]. For example, if you call
// [FetchChangesWithCompletionHandler] concurrently while the engine is
// already fetching changes, this event might not be sent. Similarly, if
// there’s no logged-in account, the engine might short-circuit the call to
// `fetchChanges`, and this event won’t be sent.
//
// # Instance Properties
//
//   - [CKSyncEngineWillFetchChangesEvent.Context]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineWillFetchChangesEvent
type CKSyncEngineWillFetchChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineWillFetchChangesEventFromID constructs a [CKSyncEngineWillFetchChangesEvent] from an objc.ID.
//
// The sync engine is about to fetch changes from the server.
func CKSyncEngineWillFetchChangesEventFromID(id objc.ID) CKSyncEngineWillFetchChangesEvent {
	return CKSyncEngineWillFetchChangesEvent{CKSyncEngineEvent: CKSyncEngineEventFromID(id)}
}

// Ensure CKSyncEngineWillFetchChangesEvent implements ICKSyncEngineWillFetchChangesEvent.
var _ ICKSyncEngineWillFetchChangesEvent = CKSyncEngineWillFetchChangesEvent{}

// An interface definition for the [CKSyncEngineWillFetchChangesEvent] class.
//
// # Instance Properties
//
//   - [ICKSyncEngineWillFetchChangesEvent.Context]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineWillFetchChangesEvent
type ICKSyncEngineWillFetchChangesEvent interface {
	ICKSyncEngineEvent

	// Topic: Instance Properties

	Context() ICKSyncEngineFetchChangesContext
}

// Init initializes the instance.
func (c CKSyncEngineWillFetchChangesEvent) Init() CKSyncEngineWillFetchChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchChangesEvent](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineWillFetchChangesEvent) Autorelease() CKSyncEngineWillFetchChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchChangesEvent](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineWillFetchChangesEvent creates a new CKSyncEngineWillFetchChangesEvent instance.
func NewCKSyncEngineWillFetchChangesEvent() CKSyncEngineWillFetchChangesEvent {
	class := getCKSyncEngineWillFetchChangesEventClass()
	rv := objc.Send[CKSyncEngineWillFetchChangesEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineWillFetchChangesEvent/context
func (c CKSyncEngineWillFetchChangesEvent) Context() ICKSyncEngineFetchChangesContext {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("context"))
	return CKSyncEngineFetchChangesContextFromID(objc.ID(rv))
}
