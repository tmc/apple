// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The names for the elements that provide values along an axis.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisElementName-c.protocol
type GCAxisElementName interface {
	objectivec.IObject
	GCPhysicalInputElementName
}

// GCAxisElementNameObject wraps an existing Objective-C object that conforms to the GCAxisElementName protocol.
type GCAxisElementNameObject struct {
	objectivec.Object
}

func (o GCAxisElementNameObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCAxisElementNameObjectFromID constructs a [GCAxisElementNameObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCAxisElementNameObjectFromID(id objc.ID) GCAxisElementNameObject {
	return GCAxisElementNameObject{
		Object: objectivec.ObjectFromID(id),
	}
}
