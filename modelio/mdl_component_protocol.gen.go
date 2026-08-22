// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The base protocol for extensible file format support in Model I/O.
//
// See: https://developer.apple.com/documentation/ModelIO/MDLComponent
type MDLComponent interface {
	objectivec.IObject
}

// MDLComponentObject wraps an existing Objective-C object that conforms to the MDLComponent protocol.
type MDLComponentObject struct {
	objectivec.Object
}

func (o MDLComponentObject) BaseObject() objectivec.Object {
	return o.Object
}

// MDLComponentObjectFromID constructs a [MDLComponentObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MDLComponentObjectFromID(id objc.ID) MDLComponentObject {
	return MDLComponentObject{
		Object: objectivec.ObjectFromID(id),
	}
}
