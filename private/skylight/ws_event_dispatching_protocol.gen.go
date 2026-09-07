// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// WSEventDispatching protocol.
type WSEventDispatching interface {
	objectivec.IObject

	// PostBackgroundEvent protocol.
	PostBackgroundEvent(event *SLSEventRecord)

	// PostEventToConnectionID protocol.
	PostEventToConnectionID(event *SLSEventRecord, id uint32)

	// PostEventToDestination protocol.
	PostEventToDestination(event *SLSEventRecord, destination objectivec.IObject)
}

// WSEventDispatchingObject wraps an existing Objective-C object that conforms to the WSEventDispatching protocol.
type WSEventDispatchingObject struct {
	objectivec.Object
}

func (o WSEventDispatchingObject) BaseObject() objectivec.Object {
	return o.Object
}

// WSEventDispatchingObjectFromID constructs a [WSEventDispatchingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func WSEventDispatchingObjectFromID(id objc.ID) WSEventDispatchingObject {
	return WSEventDispatchingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o WSEventDispatchingObject) PostBackgroundEvent(event *SLSEventRecord) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("postBackgroundEvent:"), unsafe.Pointer(event))
}
func (o WSEventDispatchingObject) PostEventToConnectionID(event *SLSEventRecord, id uint32) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("postEvent:toConnectionID:"), unsafe.Pointer(event), id)
}
func (o WSEventDispatchingObject) PostEventToDestination(event *SLSEventRecord, destination objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("postEvent:toDestination:"), unsafe.Pointer(event), destination)
}
