// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXSpecialKeyEventProcessing protocol.
type CPXSpecialKeyEventProcessing interface {
	objectivec.IObject

	// CleanupForProcessDeath protocol.
	CleanupForProcessDeath(death *CPSProcessRec)

	// ExitSpecialKeyModeForProcess protocol.
	ExitSpecialKeyModeForProcess(mode uint32, process *CPSProcessRec)

	// HotKeyChanged protocol.
	HotKeyChanged(changed unsafe.Pointer)

	// ProcessHotKeyEventHotKeyIDIsDownContextDispatcher protocol.
	ProcessHotKeyEventHotKeyIDIsDownContextDispatcher(event *SLSEventRecord, id uint64, down bool, context *CPXEventProcessorContext, dispatcher objectivec.IObject) int64

	// RegisterSpecialKeyConnectionForProcess protocol.
	RegisterSpecialKeyConnectionForProcess(key uint32, connection *CGXConnection, process *CPSProcessRec) int

	// UnregisterSpecialKeyForProcess protocol.
	UnregisterSpecialKeyForProcess(key uint32, process *CPSProcessRec) int
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

func (o CPXSpecialKeyEventProcessingObject) CleanupForProcessDeath(death *CPSProcessRec) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("cleanupForProcessDeath:"), unsafe.Pointer(death))
}
func (o CPXSpecialKeyEventProcessingObject) ExitSpecialKeyModeForProcess(mode uint32, process *CPSProcessRec) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("exitSpecialKeyMode:forProcess:"), mode, unsafe.Pointer(process))
}
func (o CPXSpecialKeyEventProcessingObject) HotKeyChanged(changed unsafe.Pointer) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("hotKeyChanged:"), changed)
}
func (o CPXSpecialKeyEventProcessingObject) ProcessHotKeyEventHotKeyIDIsDownContextDispatcher(event *SLSEventRecord, id uint64, down bool, context *CPXEventProcessorContext, dispatcher objectivec.IObject) int64 {
	rv := objc.SendIfResponds[int64](o.ID, objc.Sel("processHotKeyEvent:hotKeyID:isDown:context:dispatcher:"), unsafe.Pointer(event), id, down, unsafe.Pointer(context), dispatcher)
	return rv
}
func (o CPXSpecialKeyEventProcessingObject) RegisterSpecialKeyConnectionForProcess(key uint32, connection *CGXConnection, process *CPSProcessRec) int {
	rv := objc.SendIfResponds[int](o.ID, objc.Sel("registerSpecialKey:connection:forProcess:"), key, unsafe.Pointer(connection), unsafe.Pointer(process))
	return rv
}
func (o CPXSpecialKeyEventProcessingObject) UnregisterSpecialKeyForProcess(key uint32, process *CPSProcessRec) int {
	rv := objc.SendIfResponds[int](o.ID, objc.Sel("unregisterSpecialKey:forProcess:"), key, unsafe.Pointer(process))
	return rv
}
