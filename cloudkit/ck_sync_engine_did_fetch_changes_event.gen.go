// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [CKSyncEngineDidFetchChangesEvent] class.
var (
	_CKSyncEngineDidFetchChangesEventClass     CKSyncEngineDidFetchChangesEventClass
	_CKSyncEngineDidFetchChangesEventClassOnce sync.Once
)

func getCKSyncEngineDidFetchChangesEventClass() CKSyncEngineDidFetchChangesEventClass {
	_CKSyncEngineDidFetchChangesEventClassOnce.Do(func() {
		_CKSyncEngineDidFetchChangesEventClass = CKSyncEngineDidFetchChangesEventClass{class: objc.GetClass("CKSyncEngineDidFetchChangesEvent")}
	})
	return _CKSyncEngineDidFetchChangesEventClass
}

// GetCKSyncEngineDidFetchChangesEventClass returns the class object for CKSyncEngineDidFetchChangesEvent.
func GetCKSyncEngineDidFetchChangesEventClass() CKSyncEngineDidFetchChangesEventClass {
	return getCKSyncEngineDidFetchChangesEventClass()
}

type CKSyncEngineDidFetchChangesEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineDidFetchChangesEventClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineDidFetchChangesEventClass) Alloc() CKSyncEngineDidFetchChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The sync engine finished fetching changes from the server.
//
// # Overview
//
// You should receive one [CKSyncEngineDidFetchChangesEvent] for each
// [CKSyncEngineWillFetchChangesEvent].
//
// # Instance Properties
//
//   - [CKSyncEngineDidFetchChangesEvent.Context]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidFetchChangesEvent
type CKSyncEngineDidFetchChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineDidFetchChangesEventFromID constructs a [CKSyncEngineDidFetchChangesEvent] from an objc.ID.
//
// The sync engine finished fetching changes from the server.
func CKSyncEngineDidFetchChangesEventFromID(id objc.ID) CKSyncEngineDidFetchChangesEvent {
	return CKSyncEngineDidFetchChangesEvent{CKSyncEngineEvent: CKSyncEngineEventFromID(id)}
}

// Ensure CKSyncEngineDidFetchChangesEvent implements ICKSyncEngineDidFetchChangesEvent.
var _ ICKSyncEngineDidFetchChangesEvent = CKSyncEngineDidFetchChangesEvent{}

// An interface definition for the [CKSyncEngineDidFetchChangesEvent] class.
//
// # Instance Properties
//
//   - [ICKSyncEngineDidFetchChangesEvent.Context]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidFetchChangesEvent
type ICKSyncEngineDidFetchChangesEvent interface {
	ICKSyncEngineEvent

	// Topic: Instance Properties

	Context() ICKSyncEngineFetchChangesContext
}

// Init initializes the instance.
func (c CKSyncEngineDidFetchChangesEvent) Init() CKSyncEngineDidFetchChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchChangesEvent](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineDidFetchChangesEvent) Autorelease() CKSyncEngineDidFetchChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchChangesEvent](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineDidFetchChangesEvent creates a new CKSyncEngineDidFetchChangesEvent instance.
func NewCKSyncEngineDidFetchChangesEvent() CKSyncEngineDidFetchChangesEvent {
	class := getCKSyncEngineDidFetchChangesEventClass()
	rv := objc.Send[CKSyncEngineDidFetchChangesEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidFetchChangesEvent/context
func (c CKSyncEngineDidFetchChangesEvent) Context() ICKSyncEngineFetchChangesContext {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("context"))
	return CKSyncEngineFetchChangesContextFromID(objc.ID(rv))
}
