// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [CKSyncEngineDidSendChangesEvent] class.
var (
	_CKSyncEngineDidSendChangesEventClass     CKSyncEngineDidSendChangesEventClass
	_CKSyncEngineDidSendChangesEventClassOnce sync.Once
)

func getCKSyncEngineDidSendChangesEventClass() CKSyncEngineDidSendChangesEventClass {
	_CKSyncEngineDidSendChangesEventClassOnce.Do(func() {
		_CKSyncEngineDidSendChangesEventClass = CKSyncEngineDidSendChangesEventClass{class: objc.GetClass("CKSyncEngineDidSendChangesEvent")}
	})
	return _CKSyncEngineDidSendChangesEventClass
}

// GetCKSyncEngineDidSendChangesEventClass returns the class object for CKSyncEngineDidSendChangesEvent.
func GetCKSyncEngineDidSendChangesEventClass() CKSyncEngineDidSendChangesEventClass {
	return getCKSyncEngineDidSendChangesEventClass()
}

type CKSyncEngineDidSendChangesEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineDidSendChangesEventClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineDidSendChangesEventClass) Alloc() CKSyncEngineDidSendChangesEvent {
	rv := objc.Send[CKSyncEngineDidSendChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides information about a finished send operation.
//
// # Accessing the context
//
//   - [CKSyncEngineDidSendChangesEvent.Context]: The context of the finished send request.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidSendChangesEvent
type CKSyncEngineDidSendChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineDidSendChangesEventFromID constructs a [CKSyncEngineDidSendChangesEvent] from an objc.ID.
//
// An object that provides information about a finished send operation.
func CKSyncEngineDidSendChangesEventFromID(id objc.ID) CKSyncEngineDidSendChangesEvent {
	return CKSyncEngineDidSendChangesEvent{CKSyncEngineEvent: CKSyncEngineEventFromID(id)}
}

// Ensure CKSyncEngineDidSendChangesEvent implements ICKSyncEngineDidSendChangesEvent.
var _ ICKSyncEngineDidSendChangesEvent = CKSyncEngineDidSendChangesEvent{}

// An interface definition for the [CKSyncEngineDidSendChangesEvent] class.
//
// # Accessing the context
//
//   - [ICKSyncEngineDidSendChangesEvent.Context]: The context of the finished send request.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidSendChangesEvent
type ICKSyncEngineDidSendChangesEvent interface {
	ICKSyncEngineEvent

	// Topic: Accessing the context

	// The context of the finished send request.
	Context() ICKSyncEngineSendChangesContext
}

// Init initializes the instance.
func (c CKSyncEngineDidSendChangesEvent) Init() CKSyncEngineDidSendChangesEvent {
	rv := objc.Send[CKSyncEngineDidSendChangesEvent](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineDidSendChangesEvent) Autorelease() CKSyncEngineDidSendChangesEvent {
	rv := objc.Send[CKSyncEngineDidSendChangesEvent](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineDidSendChangesEvent creates a new CKSyncEngineDidSendChangesEvent instance.
func NewCKSyncEngineDidSendChangesEvent() CKSyncEngineDidSendChangesEvent {
	class := getCKSyncEngineDidSendChangesEventClass()
	rv := objc.Send[CKSyncEngineDidSendChangesEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The context of the finished send request.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidSendChangesEvent/context
func (c CKSyncEngineDidSendChangesEvent) Context() ICKSyncEngineSendChangesContext {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("context"))
	return CKSyncEngineSendChangesContextFromID(objc.ID(rv))
}
