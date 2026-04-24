// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXSpecialKeyEventProcessing protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessing
type CPXSpecialKeyEventProcessing interface {
	objectivec.IObject

	// CleanupForProcessDeath protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessing/cleanupForProcessDeath:
	CleanupForProcessDeath(death *CPSProcessRecRef)

	// ExitSpecialKeyModeForProcess protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessing/exitSpecialKeyMode:forProcess:
	ExitSpecialKeyModeForProcess(mode uint32, process *CPSProcessRecRef)

	// HotKeyChanged protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessing/hotKeyChanged:
	HotKeyChanged(changed objectivec.IObject)

	// RegisterSpecialKeyConnectionForProcess protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessing/registerSpecialKey:connection:forProcess:
	RegisterSpecialKeyConnectionForProcess(key uint32, connection unsafe.Pointer, process *CPSProcessRecRef) int

	// UnregisterSpecialKeyForProcess protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessing/unregisterSpecialKey:forProcess:
	UnregisterSpecialKeyForProcess(key uint32, process *CPSProcessRecRef) int
}

// CPXSpecialKeyEventProcessingObject wraps an existing Objective-C object that conforms to the CPXSpecialKeyEventProcessing protocol.
type CPXSpecialKeyEventProcessingObject struct {
	objectivec.Object
}

func (o CPXSpecialKeyEventProcessingObject) BaseObject() objectivec.Object {
	return o.Object
}

// CPXSpecialKeyEventProcessingObjectFromID constructs a [CPXSpecialKeyEventProcessingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CPXSpecialKeyEventProcessingObjectFromID(id objc.ID) CPXSpecialKeyEventProcessingObject {
	return CPXSpecialKeyEventProcessingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessing/cleanupForProcessDeath:
func (o CPXSpecialKeyEventProcessingObject) CleanupForProcessDeath(death *CPSProcessRecRef) {
	objc.Send[struct{}](o.ID, objc.Sel("cleanupForProcessDeath:"), death)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessing/exitSpecialKeyMode:forProcess:
func (o CPXSpecialKeyEventProcessingObject) ExitSpecialKeyModeForProcess(mode uint32, process *CPSProcessRecRef) {
	objc.Send[struct{}](o.ID, objc.Sel("exitSpecialKeyMode:forProcess:"), mode, process)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessing/hotKeyChanged:
func (o CPXSpecialKeyEventProcessingObject) HotKeyChanged(changed objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("hotKeyChanged:"), changed)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessing/processHotKeyEvent:hotKeyID:isDown:context:dispatcher:
func (o CPXSpecialKeyEventProcessingObject) ProcessHotKeyEventHotKeyIDIsDownContextDispatcher(event *SLSEventRecordRef, id uint64, down bool, context unsafe.Pointer, dispatcher objectivec.IObject) int64 {
	rv := objc.Send[int64](o.ID, objc.Sel("processHotKeyEvent:hotKeyID:isDown:context:dispatcher:"), event, id, down, context, dispatcher)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessing/registerSpecialKey:connection:forProcess:
func (o CPXSpecialKeyEventProcessingObject) RegisterSpecialKeyConnectionForProcess(key uint32, connection unsafe.Pointer, process *CPSProcessRecRef) int {
	rv := objc.Send[int](o.ID, objc.Sel("registerSpecialKey:connection:forProcess:"), key, connection, process)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessing/unregisterSpecialKey:forProcess:
func (o CPXSpecialKeyEventProcessingObject) UnregisterSpecialKeyForProcess(key uint32, process *CPSProcessRecRef) int {
	rv := objc.Send[int](o.ID, objc.Sel("unregisterSpecialKey:forProcess:"), key, process)
	return rv
}
