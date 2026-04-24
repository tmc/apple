// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ECSLGestureEventAuthenticationMessage protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/ECSLGestureEventAuthenticationMessage
type ECSLGestureEventAuthenticationMessage interface {
	objectivec.IObject

	// GestureHidType protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/ECSLGestureEventAuthenticationMessage/gestureHidType
	GestureHidType() uint32

	// GesturePhase protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/ECSLGestureEventAuthenticationMessage/gesturePhase
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

// See: https://developer.apple.com/documentation/SkyLight/ECSLGestureEventAuthenticationMessage/gestureHidType
func (o ECSLGestureEventAuthenticationMessageObject) GestureHidType() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("gestureHidType"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECSLGestureEventAuthenticationMessage/gesturePhase
func (o ECSLGestureEventAuthenticationMessageObject) GesturePhase() byte {
	rv := objc.Send[byte](o.ID, objc.Sel("gesturePhase"))
	return rv
}
