// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The names of the button elements.
//
// See: https://developer.apple.com/documentation/GameController/GCButtonElementName-c.protocol
type GCButtonElementName interface {
	objectivec.IObject
	GCPhysicalInputElementName
}

// GCButtonElementNameObject wraps an existing Objective-C object that conforms to the GCButtonElementName protocol.
type GCButtonElementNameObject struct {
	objectivec.Object
}

func (o GCButtonElementNameObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCButtonElementNameObjectFromID constructs a [GCButtonElementNameObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCButtonElementNameObjectFromID(id objc.ID) GCButtonElementNameObject {
	return GCButtonElementNameObject{
		Object: objectivec.ObjectFromID(id),
	}
}
