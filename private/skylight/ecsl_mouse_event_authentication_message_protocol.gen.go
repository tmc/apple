// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ECSLMouseEventAuthenticationMessage protocol.
type ECSLMouseEventAuthenticationMessage interface {
	objectivec.IObject

	// ButtonNumber protocol.
	ButtonNumber() int8
}

// ECSLMouseEventAuthenticationMessageObject wraps an existing Objective-C object that conforms to the ECSLMouseEventAuthenticationMessage protocol.
type ECSLMouseEventAuthenticationMessageObject struct {
	objectivec.Object
}

func (o ECSLMouseEventAuthenticationMessageObject) BaseObject() objectivec.Object {
	return o.Object
}

// ECSLMouseEventAuthenticationMessageObjectFromID constructs a [ECSLMouseEventAuthenticationMessageObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ECSLMouseEventAuthenticationMessageObjectFromID(id objc.ID) ECSLMouseEventAuthenticationMessageObject {
	return ECSLMouseEventAuthenticationMessageObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o ECSLMouseEventAuthenticationMessageObject) ButtonNumber() int8 {
	rv := objc.Send[int8](o.ID, objc.Sel("buttonNumber"))
	return rv
}
