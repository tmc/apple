// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// WSHIDEventDispatching protocol.
type WSHIDEventDispatching interface {
	objectivec.IObject

	// PostHIDEventForTargetIdentifierToConnectionID protocol.
	PostHIDEventForTargetIdentifierToConnectionID(hIDEvent uintptr, identifier unsafe.Pointer, id uint32)
}

// WSHIDEventDispatchingObject wraps an existing Objective-C object that conforms to the WSHIDEventDispatching protocol.
type WSHIDEventDispatchingObject struct {
	objectivec.Object
}

func (o WSHIDEventDispatchingObject) BaseObject() objectivec.Object {
	return o.Object
}

// WSHIDEventDispatchingObjectFromID constructs a [WSHIDEventDispatchingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func WSHIDEventDispatchingObjectFromID(id objc.ID) WSHIDEventDispatchingObject {
	return WSHIDEventDispatchingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o WSHIDEventDispatchingObject) PostHIDEventForTargetIdentifierToConnectionID(hIDEvent uintptr, identifier unsafe.Pointer, id uint32) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("postHIDEvent:forTargetIdentifier:toConnectionID:"), hIDEvent, identifier, id)
}
