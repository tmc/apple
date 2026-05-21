// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKSyncEngineState] class.
var (
	_CKSyncEngineStateClass     CKSyncEngineStateClass
	_CKSyncEngineStateClassOnce sync.Once
)

func getCKSyncEngineStateClass() CKSyncEngineStateClass {
	_CKSyncEngineStateClassOnce.Do(func() {
		_CKSyncEngineStateClass = CKSyncEngineStateClass{class: objc.GetClass("CKSyncEngineState")}
	})
	return _CKSyncEngineStateClass
}

// GetCKSyncEngineStateClass returns the class object for CKSyncEngineState.
func GetCKSyncEngineStateClass() CKSyncEngineStateClass {
	return getCKSyncEngineStateClass()
}

type CKSyncEngineStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineStateClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineStateClass) Alloc() CKSyncEngineState {
	rv := objc.Send[CKSyncEngineState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that manages the sync engine’s state.
//
// # Overview
//
// To reliably and consistently sync your app’s data, a sync engine keeps a
// record of several important pieces of data, such as server changes tokens
// (for databases and record zones), subscription identifiers, the most recent
// [CKUserIdentity.UserRecordID], and so on. This class automatically manages
// that state on behalf of your app, but there are certain elements you can
// modify. For example, you control the list of pending changes to send to the
// iCloud servers and manipulate that list using the
// [CKSyncEngineState.AddPendingDatabaseChanges] and
// [CKSyncEngineState.AddPendingRecordZoneChanges] methods. If there aren’t
// any scheduled sync operations when you invoke these methods, the engine
// automatically schedules one.
//
// An engine’s state changes periodically and, when it does, the sync engine
// dispatches an event of type [CKSyncEngineStateUpdateEvent] to your
// delegate. The event contains an instance of
// [CKSyncEngineStateSerialization] and, on receipt of such an event, it’s
// your responsibility to persist the serialized state to disk so that it’s
// available across app launches. On the next initialization of the sync
// engine, you provide the most recently persisted state as part of the
// engine’s configuration. For more information, see
// [CKSyncEngineConfiguration.InitWithDatabaseStateSerializationDelegate].
//
// # Accessing pending changes
//
//   - [CKSyncEngineState.HasPendingUntrackedChanges]: A Boolean value that indicates whether there are pending changes that the sync engine is unaware of.
//   - [CKSyncEngineState.SetHasPendingUntrackedChanges]
//   - [CKSyncEngineState.PendingDatabaseChanges]: The database changes that the sync engine has yet to send to the iCloud servers.
//   - [CKSyncEngineState.PendingRecordZoneChanges]: The record zone changes that the sync engine has yet to send to the iCloud servers.
//
// # Modifying pending changes
//
//   - [CKSyncEngineState.AddPendingDatabaseChanges]: Adds the specified database changes to the state.
//   - [CKSyncEngineState.RemovePendingDatabaseChanges]: Removes the specified database changes from the state.
//   - [CKSyncEngineState.AddPendingRecordZoneChanges]: Adds the specified record zone changes to the state.
//   - [CKSyncEngineState.RemovePendingRecordZoneChanges]: Removes the specified record zone changes from the state.
//
// # Instance Properties
//
//   - [CKSyncEngineState.ZoneIDsWithUnfetchedServerChanges]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState
type CKSyncEngineState struct {
	objectivec.Object
}

// CKSyncEngineStateFromID constructs a [CKSyncEngineState] from an objc.ID.
//
// An object that manages the sync engine’s state.
func CKSyncEngineStateFromID(id objc.ID) CKSyncEngineState {
	return CKSyncEngineState{objectivec.Object{ID: id}}
}

// Ensure CKSyncEngineState implements ICKSyncEngineState.
var _ ICKSyncEngineState = CKSyncEngineState{}

// An interface definition for the [CKSyncEngineState] class.
//
// # Accessing pending changes
//
//   - [ICKSyncEngineState.HasPendingUntrackedChanges]: A Boolean value that indicates whether there are pending changes that the sync engine is unaware of.
//   - [ICKSyncEngineState.SetHasPendingUntrackedChanges]
//   - [ICKSyncEngineState.PendingDatabaseChanges]: The database changes that the sync engine has yet to send to the iCloud servers.
//   - [ICKSyncEngineState.PendingRecordZoneChanges]: The record zone changes that the sync engine has yet to send to the iCloud servers.
//
// # Modifying pending changes
//
//   - [ICKSyncEngineState.AddPendingDatabaseChanges]: Adds the specified database changes to the state.
//   - [ICKSyncEngineState.RemovePendingDatabaseChanges]: Removes the specified database changes from the state.
//   - [ICKSyncEngineState.AddPendingRecordZoneChanges]: Adds the specified record zone changes to the state.
//   - [ICKSyncEngineState.RemovePendingRecordZoneChanges]: Removes the specified record zone changes from the state.
//
// # Instance Properties
//
//   - [ICKSyncEngineState.ZoneIDsWithUnfetchedServerChanges]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState
type ICKSyncEngineState interface {
	objectivec.IObject

	// Topic: Accessing pending changes

	// A Boolean value that indicates whether there are pending changes that the sync engine is unaware of.
	HasPendingUntrackedChanges() bool
	SetHasPendingUntrackedChanges(value bool)
	// The database changes that the sync engine has yet to send to the iCloud servers.
	PendingDatabaseChanges() []CKSyncEnginePendingDatabaseChange
	// The record zone changes that the sync engine has yet to send to the iCloud servers.
	PendingRecordZoneChanges() []CKSyncEnginePendingRecordZoneChange

	// Topic: Modifying pending changes

	// Adds the specified database changes to the state.
	AddPendingDatabaseChanges(changes []CKSyncEnginePendingDatabaseChange)
	// Removes the specified database changes from the state.
	RemovePendingDatabaseChanges(changes []CKSyncEnginePendingDatabaseChange)
	// Adds the specified record zone changes to the state.
	AddPendingRecordZoneChanges(changes []CKSyncEnginePendingRecordZoneChange)
	// Removes the specified record zone changes from the state.
	RemovePendingRecordZoneChanges(changes []CKSyncEnginePendingRecordZoneChange)

	// Topic: Instance Properties

	ZoneIDsWithUnfetchedServerChanges() []CKRecordZoneID
}

// Init initializes the instance.
func (c CKSyncEngineState) Init() CKSyncEngineState {
	rv := objc.Send[CKSyncEngineState](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineState) Autorelease() CKSyncEngineState {
	rv := objc.Send[CKSyncEngineState](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineState creates a new CKSyncEngineState instance.
func NewCKSyncEngineState() CKSyncEngineState {
	class := getCKSyncEngineStateClass()
	rv := objc.Send[CKSyncEngineState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Adds the specified database changes to the state.
//
// changes: An array of database changes.
//
// # Discussion
//
// Use this method to enable the sync engine to manage your pending database
// changes. For example, when someone makes a change that your app needs to
// send to the server, use this method to record the change. If there are no
// scheduled sync operations when you invoke this method, the sync engine
// automatically schedules one to send the changes. After the engine sends
// those changes, it notifies your app’s sync delegate with an event of type
// [CKSyncEngineSentDatabaseChangesEvent].
//
// The sync engine ensures the consistency of any pending changes it’s
// tracking, deduplicating them as necessary. The engine removes changes from
// the list as it sends them, but retains any that fail due to a recoverable
// error, such as a network issue or exceeding the rate limit.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState/addPendingDatabaseChanges:
func (c CKSyncEngineState) AddPendingDatabaseChanges(changes []CKSyncEnginePendingDatabaseChange) {
	objc.Send[objc.ID](c.ID, objc.Sel("addPendingDatabaseChanges:"), objectivec.IObjectSliceToNSArray(changes))
}

// Removes the specified database changes from the state.
//
// changes: An array of database changes.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState/removePendingDatabaseChanges:
func (c CKSyncEngineState) RemovePendingDatabaseChanges(changes []CKSyncEnginePendingDatabaseChange) {
	objc.Send[objc.ID](c.ID, objc.Sel("removePendingDatabaseChanges:"), objectivec.IObjectSliceToNSArray(changes))
}

// Adds the specified record zone changes to the state.
//
// changes: An array of record zone changes.
//
// # Discussion
//
// Use this method to enable the sync engine to manage your pending record
// zone changes. For example, when someone makes a change that your app needs
// to send to the server, use this method to record the change. Then, when
// creating the change batch for the next send operation, retrieve the pending
// changes from the [CKSyncEngineState.PendingRecordZoneChanges] property.
//
// If there are no scheduled sync operations when you invoke this method, the
// sync engine automatically schedules one to send the changes. After the
// engine sends those changes, it notifies your app’s sync delegate with an
// event of type [CKSyncEngineSentRecordZoneChangesEvent].
//
// The sync engine ensures the consistency of any pending changes it’s
// tracking, deduplicating them as necessary. The engine removes changes from
// the list as it sends them, but retains any that fail due to a recoverable
// error, such as a network issue or exceeding the rate limit.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState/addPendingRecordZoneChanges:
func (c CKSyncEngineState) AddPendingRecordZoneChanges(changes []CKSyncEnginePendingRecordZoneChange) {
	objc.Send[objc.ID](c.ID, objc.Sel("addPendingRecordZoneChanges:"), objectivec.IObjectSliceToNSArray(changes))
}

// Removes the specified record zone changes from the state.
//
// changes: An array of record zone changes.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState/removePendingRecordZoneChanges:
func (c CKSyncEngineState) RemovePendingRecordZoneChanges(changes []CKSyncEnginePendingRecordZoneChange) {
	objc.Send[objc.ID](c.ID, objc.Sel("removePendingRecordZoneChanges:"), objectivec.IObjectSliceToNSArray(changes))
}

// A Boolean value that indicates whether there are pending changes that the
// sync engine is unaware of.
//
// # Discussion
//
// Use this property to inform the sync engine that there are pending changes
// other than those available in [CKSyncEngineState.PendingRecordZoneChanges].
// After you set this property, the sync engine automatically schedules a send
// operation and, when that operation executes, asks your delegate to provide
// those changes by invoking the
// [SyncEngineNextRecordZoneChangeBatchForContext] method.
//
// Using this property is optional and is necessary only if you track pending
// changes manually, outside of the sync engine’s state.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState/hasPendingUntrackedChanges
func (c CKSyncEngineState) HasPendingUntrackedChanges() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasPendingUntrackedChanges"))
	return rv
}
func (c CKSyncEngineState) SetHasPendingUntrackedChanges(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setHasPendingUntrackedChanges:"), value)
}

// The database changes that the sync engine has yet to send to the iCloud
// servers.
//
// # Discussion
//
// This array contains any pending database changes to send to the iCloud
// servers in a subsequent send operation (scheduled or manual). After the
// sync engine sends those changes, your app’s sync delegate receives an
// event of type [CKSyncEngineSentDatabaseChangesEvent].
//
// Use the [CKSyncEngineState.AddPendingDatabaseChanges] and
// [CKSyncEngineState.RemovePendingDatabaseChanges] methods to modify the
// array’s contents.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState/pendingDatabaseChanges
func (c CKSyncEngineState) PendingDatabaseChanges() []CKSyncEnginePendingDatabaseChange {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("pendingDatabaseChanges"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKSyncEnginePendingDatabaseChange {
		return CKSyncEnginePendingDatabaseChangeFromID(id)
	})
}

// The record zone changes that the sync engine has yet to send to the iCloud
// servers.
//
// # Discussion
//
// This array contains any pending record zone changes to send to the iCloud
// servers in a subsequent send operation (scheduled or manual). After the
// sync engine sends those changes, your app’s sync delegate receives an
// event of type [CKSyncEngineSentRecordZoneChangesEvent].
//
// Use the [CKSyncEngineState.AddPendingRecordZoneChanges] and
// [CKSyncEngineState.RemovePendingRecordZoneChanges] methods to modify the
// array’s contents.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState/pendingRecordZoneChanges
func (c CKSyncEngineState) PendingRecordZoneChanges() []CKSyncEnginePendingRecordZoneChange {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("pendingRecordZoneChanges"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKSyncEnginePendingRecordZoneChange {
		return CKSyncEnginePendingRecordZoneChangeFromID(id)
	})
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineState/zoneIDsWithUnfetchedServerChanges
func (c CKSyncEngineState) ZoneIDsWithUnfetchedServerChanges() []CKRecordZoneID {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("zoneIDsWithUnfetchedServerChanges"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKRecordZoneID {
		return CKRecordZoneIDFromID(id)
	})
}
