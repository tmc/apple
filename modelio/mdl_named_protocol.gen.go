// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The common interface for Model I/O objects that expose a human-readable name.
//
// See: https://developer.apple.com/documentation/ModelIO/MDLNamed
type MDLNamed interface {
	objectivec.IObject
}

// MDLNamedObject wraps an existing Objective-C object that conforms to the MDLNamed protocol.
type MDLNamedObject struct {
	objectivec.Object
}

func (o MDLNamedObject) BaseObject() objectivec.Object {
	return o.Object
}

// MDLNamedObjectFromID constructs a [MDLNamedObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MDLNamedObjectFromID(id objc.ID) MDLNamedObject {
	return MDLNamedObject{
		Object: objectivec.ObjectFromID(id),
	}
}
