// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// WSEventProcessor protocol.
type WSEventProcessor interface {
	objectivec.IObject

	// ClearEventState protocol.
	ClearEventState()

	// ProcessEventDispatcher protocol.
	ProcessEventDispatcher(event *SLSEventRecord, dispatcher objectivec.IObject) int64
}

// WSEventProcessorObject wraps an existing Objective-C object that conforms to the WSEventProcessor protocol.
type WSEventProcessorObject struct {
	objectivec.Object
}

func (o WSEventProcessorObject) BaseObject() objectivec.Object {
	return o.Object
}

// WSEventProcessorObjectFromID constructs a [WSEventProcessorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func WSEventProcessorObjectFromID(id objc.ID) WSEventProcessorObject {
	return WSEventProcessorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o WSEventProcessorObject) ClearEventState() {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("clearEventState"))
}
func (o WSEventProcessorObject) ProcessEventDispatcher(event *SLSEventRecord, dispatcher objectivec.IObject) int64 {
	rv := objc.SendIfResponds[int64](o.ID, objc.Sel("processEvent:dispatcher:"), unsafe.Pointer(event), dispatcher)
	return rv
}
