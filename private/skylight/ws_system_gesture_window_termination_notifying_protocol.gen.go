// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// WSSystemGestureWindowTerminationNotifying protocol.
type WSSystemGestureWindowTerminationNotifying interface {
	objectivec.IObject

	// SubscribeForWindowTermination protocol.
	SubscribeForWindowTermination(termination structHandler) objectivec.IObject
}

// WSSystemGestureWindowTerminationNotifyingObject wraps an existing Objective-C object that conforms to the WSSystemGestureWindowTerminationNotifying protocol.
type WSSystemGestureWindowTerminationNotifyingObject struct {
	objectivec.Object
}

func (o WSSystemGestureWindowTerminationNotifyingObject) BaseObject() objectivec.Object {
	return o.Object
}

// WSSystemGestureWindowTerminationNotifyingObjectFromID constructs a [WSSystemGestureWindowTerminationNotifyingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func WSSystemGestureWindowTerminationNotifyingObjectFromID(id objc.ID) WSSystemGestureWindowTerminationNotifyingObject {
	return WSSystemGestureWindowTerminationNotifyingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o WSSystemGestureWindowTerminationNotifyingObject) SubscribeForWindowTermination(termination structHandler) objectivec.IObject {
	_block0, _cleanup0 := NewstructBlock(termination)
	defer _cleanup0()
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("subscribeForWindowTermination:"), objc.ID(_block0))
	return objectivec.Object{ID: rv}
}
