// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXEventDispatching protocol.
type CPXEventDispatching interface {
	objectivec.IObject

	// PostBackgroundEvent protocol.
	PostBackgroundEvent(event SLSEventRecord)

	// PostEventToConnectionID protocol.
	PostEventToConnectionID(event SLSEventRecord, id uint32)
}

// CPXEventDispatchingObject wraps an existing Objective-C object that conforms to the CPXEventDispatching protocol.
type CPXEventDispatchingObject struct {
	objectivec.Object
}

func (o CPXEventDispatchingObject) BaseObject() objectivec.Object {
	return o.Object
}

// CPXEventDispatchingObjectFromID constructs a [CPXEventDispatchingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CPXEventDispatchingObjectFromID(id objc.ID) CPXEventDispatchingObject {
	return CPXEventDispatchingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o CPXEventDispatchingObject) PostBackgroundEvent(event SLSEventRecord) {
	objc.Send[struct{}](o.ID, objc.Sel("postBackgroundEvent:"), event)
}
func (o CPXEventDispatchingObject) PostEventToConnectionID(event SLSEventRecord, id uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("postEvent:toConnectionID:"), event, id)
}
func (o CPXEventDispatchingObject) PostEventToDestination(event SLSEventRecord, destination objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("postEvent:toDestination:"), event, destination)
}
