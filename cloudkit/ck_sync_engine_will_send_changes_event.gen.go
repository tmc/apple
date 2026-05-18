// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [CKSyncEngineWillSendChangesEvent] class.
var (
	_CKSyncEngineWillSendChangesEventClass     CKSyncEngineWillSendChangesEventClass
	_CKSyncEngineWillSendChangesEventClassOnce sync.Once
)

func getCKSyncEngineWillSendChangesEventClass() CKSyncEngineWillSendChangesEventClass {
	_CKSyncEngineWillSendChangesEventClassOnce.Do(func() {
		_CKSyncEngineWillSendChangesEventClass = CKSyncEngineWillSendChangesEventClass{class: objc.GetClass("CKSyncEngineWillSendChangesEvent")}
	})
	return _CKSyncEngineWillSendChangesEventClass
}

// GetCKSyncEngineWillSendChangesEventClass returns the class object for CKSyncEngineWillSendChangesEvent.
func GetCKSyncEngineWillSendChangesEventClass() CKSyncEngineWillSendChangesEventClass {
	return getCKSyncEngineWillSendChangesEventClass()
}

type CKSyncEngineWillSendChangesEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineWillSendChangesEventClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineWillSendChangesEventClass) Alloc() CKSyncEngineWillSendChangesEvent {
	rv := objc.Send[CKSyncEngineWillSendChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides information about an imminent send of local
// changes.
//
// # Accessing the context
//
//   - [CKSyncEngineWillSendChangesEvent.Context]: The context of the imminent send request.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineWillSendChangesEvent
type CKSyncEngineWillSendChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineWillSendChangesEventFromID constructs a [CKSyncEngineWillSendChangesEvent] from an objc.ID.
//
// An object that provides information about an imminent send of local
// changes.
func CKSyncEngineWillSendChangesEventFromID(id objc.ID) CKSyncEngineWillSendChangesEvent {
	return CKSyncEngineWillSendChangesEvent{CKSyncEngineEvent: CKSyncEngineEventFromID(id)}
}

// Ensure CKSyncEngineWillSendChangesEvent implements ICKSyncEngineWillSendChangesEvent.
var _ ICKSyncEngineWillSendChangesEvent = CKSyncEngineWillSendChangesEvent{}

// An interface definition for the [CKSyncEngineWillSendChangesEvent] class.
//
// # Accessing the context
//
//   - [ICKSyncEngineWillSendChangesEvent.Context]: The context of the imminent send request.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineWillSendChangesEvent
type ICKSyncEngineWillSendChangesEvent interface {
	ICKSyncEngineEvent

	// Topic: Accessing the context

	// The context of the imminent send request.
	Context() ICKSyncEngineSendChangesContext
}

// Init initializes the instance.
func (c CKSyncEngineWillSendChangesEvent) Init() CKSyncEngineWillSendChangesEvent {
	rv := objc.Send[CKSyncEngineWillSendChangesEvent](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineWillSendChangesEvent) Autorelease() CKSyncEngineWillSendChangesEvent {
	rv := objc.Send[CKSyncEngineWillSendChangesEvent](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineWillSendChangesEvent creates a new CKSyncEngineWillSendChangesEvent instance.
func NewCKSyncEngineWillSendChangesEvent() CKSyncEngineWillSendChangesEvent {
	class := getCKSyncEngineWillSendChangesEventClass()
	rv := objc.Send[CKSyncEngineWillSendChangesEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The context of the imminent send request.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineWillSendChangesEvent/context
func (c CKSyncEngineWillSendChangesEvent) Context() ICKSyncEngineSendChangesContext {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("context"))
	return CKSyncEngineSendChangesContextFromID(objc.ID(rv))
}
