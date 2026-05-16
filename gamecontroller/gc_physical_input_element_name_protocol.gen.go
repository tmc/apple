// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The name of a physical input element.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElementName-c.protocol
type GCPhysicalInputElementName interface {
	objectivec.IObject
}

// GCPhysicalInputElementNameObject wraps an existing Objective-C object that conforms to the GCPhysicalInputElementName protocol.
type GCPhysicalInputElementNameObject struct {
	objectivec.Object
}

func (o GCPhysicalInputElementNameObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCPhysicalInputElementNameObjectFromID constructs a [GCPhysicalInputElementNameObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCPhysicalInputElementNameObjectFromID(id objc.ID) GCPhysicalInputElementNameObject {
	return GCPhysicalInputElementNameObject{
		Object: objectivec.ObjectFromID(id),
	}
}
