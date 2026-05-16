// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The names for directional pad elements.
//
// See: https://developer.apple.com/documentation/GameController/GCDirectionPadElementName-c.protocol
type GCDirectionPadElementName interface {
	objectivec.IObject
	GCPhysicalInputElementName
}

// GCDirectionPadElementNameObject wraps an existing Objective-C object that conforms to the GCDirectionPadElementName protocol.
type GCDirectionPadElementNameObject struct {
	objectivec.Object
}

func (o GCDirectionPadElementNameObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCDirectionPadElementNameObjectFromID constructs a [GCDirectionPadElementNameObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCDirectionPadElementNameObjectFromID(id objc.ID) GCDirectionPadElementNameObject {
	return GCDirectionPadElementNameObject{
		Object: objectivec.ObjectFromID(id),
	}
}
