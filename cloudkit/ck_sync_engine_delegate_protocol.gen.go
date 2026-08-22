// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// An interface for providing record data to a sync engine and customizing that engine’s behavior.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDelegate-3c38p
type CKSyncEngineDelegate interface {
	objectivec.IObject

	// Tells the delegate to handle the specified sync event.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDelegate-3c38p/syncEngine:handleEvent:
	SyncEngineHandleEvent(syncEngine ICKSyncEngine, event ICKSyncEngineEvent)

	// Asks the delegate to provide the next set of record changes to send to the server.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDelegate-3c38p/syncEngine:nextRecordZoneChangeBatchForContext:
	SyncEngineNextRecordZoneChangeBatchForContext(syncEngine ICKSyncEngine, context ICKSyncEngineSendChangesContext) ICKSyncEngineRecordZoneChangeBatch

	// SyncEngineNextFetchChangesOptionsForContext protocol.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDelegate-3c38p/syncEngine:nextFetchChangesOptionsForContext:
	SyncEngineNextFetchChangesOptionsForContext(syncEngine ICKSyncEngine, context ICKSyncEngineFetchChangesContext) ICKSyncEngineFetchChangesOptions
}

// CKSyncEngineDelegateObject wraps an existing Objective-C object that conforms to the CKSyncEngineDelegate protocol.
type CKSyncEngineDelegateObject struct {
	objectivec.Object
}

func (o CKSyncEngineDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// CKSyncEngineDelegateObjectFromID constructs a [CKSyncEngineDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CKSyncEngineDelegateObjectFromID(id objc.ID) CKSyncEngineDelegateObject {
	return CKSyncEngineDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate to handle the specified sync event.
//
// syncEngine: Information about the event. An event may occur for a number of reasons,
// such as when new data is available or when the device’s iCloud account
// changes. For more information, see [CKSyncEngineEvent].
//
// event: The sync engine that generates the event.
//
// # Discussion
//
// The sync engines provides events serially; your delegate won’t receive
// the subsequent event until it finishes processing the current one and
// returns from this method.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDelegate-3c38p/syncEngine:handleEvent:
func (o CKSyncEngineDelegateObject) SyncEngineHandleEvent(syncEngine ICKSyncEngine, event ICKSyncEngineEvent) {
	objc.Send[struct{}](o.ID, objc.Sel("syncEngine:handleEvent:"), syncEngine, event)
}

// Asks the delegate to provide the next set of record changes to send to the
// server.
//
// syncEngine: The sync engine requesting changes.
//
// context: The reason for the sync engine’s request, and any additional options that
// request is using.
//
// # Return Value
//
// If there are pending record changes, a batch of those changes for the sync
// engine to process; otherwise, `nil` to indicate there are no changes to
// send.
//
// # Discussion
//
// In your implementation, ask the sync engine’s state for any pending
// record zone changes and then return a change batch containing an instance
// of [CKRecord] for each record identifier the state provides. For both
// scheduled and manual send operations, the sync engine calls this method
// repeatedly until your app has no more changes and returns `nil`.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDelegate-3c38p/syncEngine:nextRecordZoneChangeBatchForContext:
func (o CKSyncEngineDelegateObject) SyncEngineNextRecordZoneChangeBatchForContext(syncEngine ICKSyncEngine, context ICKSyncEngineSendChangesContext) ICKSyncEngineRecordZoneChangeBatch {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("syncEngine:nextRecordZoneChangeBatchForContext:"), syncEngine, context)
	return CKSyncEngineRecordZoneChangeBatchFromID(rv)
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDelegate-3c38p/syncEngine:nextFetchChangesOptionsForContext:
func (o CKSyncEngineDelegateObject) SyncEngineNextFetchChangesOptionsForContext(syncEngine ICKSyncEngine, context ICKSyncEngineFetchChangesContext) ICKSyncEngineFetchChangesOptions {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("syncEngine:nextFetchChangesOptionsForContext:"), syncEngine, context)
	return CKSyncEngineFetchChangesOptionsFromID(rv)
}

// CKSyncEngineDelegateConfig holds optional typed callbacks for [CKSyncEngineDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/cloudkit/cksyncenginedelegate
type CKSyncEngineDelegateConfig struct {

	// Other Methods
	// SyncEngineHandleEvent — Tells the delegate to handle the specified sync event.
	SyncEngineHandleEvent func(syncEngine CKSyncEngine, event CKSyncEngineEvent)
	// SyncEngineNextRecordZoneChangeBatchForContext — Asks the delegate to provide the next set of record changes to send to the server.
	SyncEngineNextRecordZoneChangeBatchForContext func(syncEngine CKSyncEngine, context CKSyncEngineSendChangesContext) CKSyncEngineRecordZoneChangeBatch
	SyncEngineNextFetchChangesOptionsForContext   func(syncEngine CKSyncEngine, context CKSyncEngineFetchChangesContext) CKSyncEngineFetchChangesOptions
}

// NewCKSyncEngineDelegate creates an Objective-C object implementing the [CKSyncEngineDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [CKSyncEngineDelegateObject] satisfies the [CKSyncEngineDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/cloudkit/cksyncenginedelegate
func NewCKSyncEngineDelegate(config CKSyncEngineDelegateConfig) CKSyncEngineDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoCKSyncEngineDelegate_%d", n)

	var methods []objc.MethodDef

	if config.SyncEngineHandleEvent != nil {
		fn := config.SyncEngineHandleEvent
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("syncEngine:handleEvent:"),
			Fn: func(self objc.ID, _cmd objc.SEL, syncEngineID objc.ID, eventID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CKSyncEngineDelegate", "syncEngine:handleEvent:")
					}
				}()
				syncEngine := CKSyncEngineFromID(syncEngineID)
				event := CKSyncEngineEventFromID(eventID)
				fn(syncEngine, event)
				_delegateDone = true
			},
		})
	}

	if config.SyncEngineNextRecordZoneChangeBatchForContext != nil {
		fn := config.SyncEngineNextRecordZoneChangeBatchForContext
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("syncEngine:nextRecordZoneChangeBatchForContext:"),
			Fn: func(self objc.ID, _cmd objc.SEL, syncEngineID objc.ID, contextID objc.ID) objc.ID {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CKSyncEngineDelegate", "syncEngine:nextRecordZoneChangeBatchForContext:")
					}
				}()
				syncEngine := CKSyncEngineFromID(syncEngineID)
				context := CKSyncEngineSendChangesContextFromID(contextID)
				_delegateResult := fn(syncEngine, context).GetID()
				_delegateDone = true
				return _delegateResult
			},
		})
	}

	if config.SyncEngineNextFetchChangesOptionsForContext != nil {
		fn := config.SyncEngineNextFetchChangesOptionsForContext
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("syncEngine:nextFetchChangesOptionsForContext:"),
			Fn: func(self objc.ID, _cmd objc.SEL, syncEngineID objc.ID, contextID objc.ID) objc.ID {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CKSyncEngineDelegate", "syncEngine:nextFetchChangesOptionsForContext:")
					}
				}()
				syncEngine := CKSyncEngineFromID(syncEngineID)
				context := CKSyncEngineFetchChangesContextFromID(contextID)
				_delegateResult := fn(syncEngine, context).GetID()
				_delegateDone = true
				return _delegateResult
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("CKSyncEngineDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewCKSyncEngineDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return CKSyncEngineDelegateObjectFromID(instance)
}
