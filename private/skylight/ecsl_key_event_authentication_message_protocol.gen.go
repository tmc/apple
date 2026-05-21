// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ECSLKeyEventAuthenticationMessage protocol.
type ECSLKeyEventAuthenticationMessage interface {
	objectivec.IObject

	// CharCode protocol.
	CharCode() uint16

	// Repeat protocol.
	Repeat() int16
}

// ECSLKeyEventAuthenticationMessageObject wraps an existing Objective-C object that conforms to the ECSLKeyEventAuthenticationMessage protocol.
type ECSLKeyEventAuthenticationMessageObject struct {
	objectivec.Object
}

func (o ECSLKeyEventAuthenticationMessageObject) BaseObject() objectivec.Object {
	return o.Object
}

// ECSLKeyEventAuthenticationMessageObjectFromID constructs a [ECSLKeyEventAuthenticationMessageObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ECSLKeyEventAuthenticationMessageObjectFromID(id objc.ID) ECSLKeyEventAuthenticationMessageObject {
	return ECSLKeyEventAuthenticationMessageObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o ECSLKeyEventAuthenticationMessageObject) CharCode() uint16 {
	rv := objc.Send[uint16](o.ID, objc.Sel("charCode"))
	return rv
}
func (o ECSLKeyEventAuthenticationMessageObject) Repeat() int16 {
	rv := objc.Send[int16](o.ID, objc.Sel("repeat"))
	return rv
}
