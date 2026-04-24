// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXEventDispatching protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/CPXEventDispatching
type CPXEventDispatching interface {
	objectivec.IObject

	// PostBackgroundEvent protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXEventDispatching/postBackgroundEvent:
	PostBackgroundEvent(event *SLSEventRecordRef)

	// PostEventToConnectionID protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXEventDispatching/postEvent:toConnectionID:
	PostEventToConnectionID(event *SLSEventRecordRef, id uint32)
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

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDispatching/postBackgroundEvent:
func (o CPXEventDispatchingObject) PostBackgroundEvent(event *SLSEventRecordRef) {
	objc.Send[struct{}](o.ID, objc.Sel("postBackgroundEvent:"), event)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDispatching/postEvent:toConnectionID:
func (o CPXEventDispatchingObject) PostEventToConnectionID(event *SLSEventRecordRef, id uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("postEvent:toConnectionID:"), event, id)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDispatching/postEvent:toDestination:
func (o CPXEventDispatchingObject) PostEventToDestination(event *SLSEventRecordRef, destination objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("postEvent:toDestination:"), event, destination)
}
