// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXEventProcessor protocol.
type CPXEventProcessor interface {
	objectivec.IObject

	// ClearEventState protocol.
	ClearEventState()

	// ProcessEventContextDispatcher protocol.
	ProcessEventContextDispatcher(event *SLSEventRecord, context *CPXEventProcessorContext, dispatcher objectivec.IObject) int64
}

// CPXEventProcessorObject wraps an existing Objective-C object that conforms to the CPXEventProcessor protocol.
type CPXEventProcessorObject struct {
	objectivec.Object
}

func (o CPXEventProcessorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CPXEventProcessorObjectFromID constructs a [CPXEventProcessorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CPXEventProcessorObjectFromID(id objc.ID) CPXEventProcessorObject {
	return CPXEventProcessorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o CPXEventProcessorObject) ClearEventState() {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("clearEventState"))
}
func (o CPXEventProcessorObject) ProcessEventContextDispatcher(event *SLSEventRecord, context *CPXEventProcessorContext, dispatcher objectivec.IObject) int64 {
	rv := objc.SendIfResponds[int64](o.ID, objc.Sel("processEvent:context:dispatcher:"), unsafe.Pointer(event), unsafe.Pointer(context), dispatcher)
	return rv
}
