// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ECSLGestureEventAuthenticationMessage protocol.
type ECSLGestureEventAuthenticationMessage interface {
	objectivec.IObject

	// GestureHidType protocol.
	GestureHidType() uint32

	// GesturePhase protocol.
	GesturePhase() byte
}

// ECSLGestureEventAuthenticationMessageObject wraps an existing Objective-C object that conforms to the ECSLGestureEventAuthenticationMessage protocol.
type ECSLGestureEventAuthenticationMessageObject struct {
	objectivec.Object
}

func (o ECSLGestureEventAuthenticationMessageObject) BaseObject() objectivec.Object {
	return o.Object
}

// ECSLGestureEventAuthenticationMessageObjectFromID constructs a [ECSLGestureEventAuthenticationMessageObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ECSLGestureEventAuthenticationMessageObjectFromID(id objc.ID) ECSLGestureEventAuthenticationMessageObject {
	return ECSLGestureEventAuthenticationMessageObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o ECSLGestureEventAuthenticationMessageObject) GestureHidType() uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("gestureHidType"))
	return rv
}
func (o ECSLGestureEventAuthenticationMessageObject) GesturePhase() byte {
	rv := objc.SendIfResponds[byte](o.ID, objc.Sel("gesturePhase"))
	return rv
}
