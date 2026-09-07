// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLSDisplaySyncSessionControl protocol.
type SLSDisplaySyncSessionControl interface {
	objectivec.IObject

	// RegisterForNotificationsWithQueueBlock protocol.
	RegisterForNotificationsWithQueueBlock(queue objectivec.IObject, block IObjectINSDictionaryHandler)

	// UnregisterNotification protocol.
	UnregisterNotification()
}

// SLSDisplaySyncSessionControlObject wraps an existing Objective-C object that conforms to the SLSDisplaySyncSessionControl protocol.
type SLSDisplaySyncSessionControlObject struct {
	objectivec.Object
}

func (o SLSDisplaySyncSessionControlObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLSDisplaySyncSessionControlObjectFromID constructs a [SLSDisplaySyncSessionControlObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLSDisplaySyncSessionControlObjectFromID(id objc.ID) SLSDisplaySyncSessionControlObject {
	return SLSDisplaySyncSessionControlObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o SLSDisplaySyncSessionControlObject) RegisterForNotificationsWithQueueBlock(queue objectivec.IObject, block IObjectINSDictionaryHandler) {
	_block1, _cleanup1 := NewIObjectINSDictionaryBlock(block)
	defer _cleanup1()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("registerForNotificationsWithQueue:block:"), queue, objc.ID(_block1))
}
func (o SLSDisplaySyncSessionControlObject) UnregisterNotification() {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("unregisterNotification"))
}
