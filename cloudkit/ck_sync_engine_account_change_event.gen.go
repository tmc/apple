// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [CKSyncEngineAccountChangeEvent] class.
var (
	_CKSyncEngineAccountChangeEventClass     CKSyncEngineAccountChangeEventClass
	_CKSyncEngineAccountChangeEventClassOnce sync.Once
)

func getCKSyncEngineAccountChangeEventClass() CKSyncEngineAccountChangeEventClass {
	_CKSyncEngineAccountChangeEventClassOnce.Do(func() {
		_CKSyncEngineAccountChangeEventClass = CKSyncEngineAccountChangeEventClass{class: objc.GetClass("CKSyncEngineAccountChangeEvent")}
	})
	return _CKSyncEngineAccountChangeEventClass
}

// GetCKSyncEngineAccountChangeEventClass returns the class object for CKSyncEngineAccountChangeEvent.
func GetCKSyncEngineAccountChangeEventClass() CKSyncEngineAccountChangeEventClass {
	return getCKSyncEngineAccountChangeEventClass()
}

type CKSyncEngineAccountChangeEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineAccountChangeEventClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineAccountChangeEventClass) Alloc() CKSyncEngineAccountChangeEvent {
	rv := objc.Send[CKSyncEngineAccountChangeEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The user signed in or out of their account.
//
// # Overview
//
// The sync engine automatically listens for account changes, and it sends
// this event when the user signs in or out. It’s your responsibility to
// react appropriately to this change and update your local persistence.
//
// When the logged-in account changes, the sync engine resets its internal
// state. This means that it clears any pending database or record zone
// changes that you may have added.
//
// Note that it’s possible the account changes multiple times while your app
// is quit. If this happens, you only receive one account change event
// representing the transition between the last known state and the current
// state.
//
// # Understanding the change
//
//   - [CKSyncEngineAccountChangeEvent.ChangeType]: The iCloud account’s change type.
//   - [CKSyncEngineAccountChangeEvent.PreviousUser]: The previous iCloud account’s record identifier.
//   - [CKSyncEngineAccountChangeEvent.CurrentUser]: The current iCloud account’s record identifier.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineAccountChangeEvent
type CKSyncEngineAccountChangeEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineAccountChangeEventFromID constructs a [CKSyncEngineAccountChangeEvent] from an objc.ID.
//
// The user signed in or out of their account.
func CKSyncEngineAccountChangeEventFromID(id objc.ID) CKSyncEngineAccountChangeEvent {
	return CKSyncEngineAccountChangeEvent{CKSyncEngineEvent: CKSyncEngineEventFromID(id)}
}

// Ensure CKSyncEngineAccountChangeEvent implements ICKSyncEngineAccountChangeEvent.
var _ ICKSyncEngineAccountChangeEvent = CKSyncEngineAccountChangeEvent{}

// An interface definition for the [CKSyncEngineAccountChangeEvent] class.
//
// # Understanding the change
//
//   - [ICKSyncEngineAccountChangeEvent.ChangeType]: The iCloud account’s change type.
//   - [ICKSyncEngineAccountChangeEvent.PreviousUser]: The previous iCloud account’s record identifier.
//   - [ICKSyncEngineAccountChangeEvent.CurrentUser]: The current iCloud account’s record identifier.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineAccountChangeEvent
type ICKSyncEngineAccountChangeEvent interface {
	ICKSyncEngineEvent

	// Topic: Understanding the change

	// The iCloud account’s change type.
	ChangeType() CKSyncEngineAccountChangeType
	// The previous iCloud account’s record identifier.
	PreviousUser() ICKRecordID
	// The current iCloud account’s record identifier.
	CurrentUser() ICKRecordID
}

// Init initializes the instance.
func (c CKSyncEngineAccountChangeEvent) Init() CKSyncEngineAccountChangeEvent {
	rv := objc.Send[CKSyncEngineAccountChangeEvent](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineAccountChangeEvent) Autorelease() CKSyncEngineAccountChangeEvent {
	rv := objc.Send[CKSyncEngineAccountChangeEvent](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineAccountChangeEvent creates a new CKSyncEngineAccountChangeEvent instance.
func NewCKSyncEngineAccountChangeEvent() CKSyncEngineAccountChangeEvent {
	class := getCKSyncEngineAccountChangeEventClass()
	rv := objc.Send[CKSyncEngineAccountChangeEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The iCloud account’s change type.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineAccountChangeEvent/changeType
func (c CKSyncEngineAccountChangeEvent) ChangeType() CKSyncEngineAccountChangeType {
	rv := objc.Send[CKSyncEngineAccountChangeType](c.ID, objc.Sel("changeType"))
	return CKSyncEngineAccountChangeType(rv)
}

// The previous iCloud account’s record identifier.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineAccountChangeEvent/previousUser
func (c CKSyncEngineAccountChangeEvent) PreviousUser() ICKRecordID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("previousUser"))
	return CKRecordIDFromID(objc.ID(rv))
}

// The current iCloud account’s record identifier.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineAccountChangeEvent/currentUser
func (c CKSyncEngineAccountChangeEvent) CurrentUser() ICKRecordID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("currentUser"))
	return CKRecordIDFromID(objc.ID(rv))
}
