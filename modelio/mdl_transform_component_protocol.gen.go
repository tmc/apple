// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The general interface for classes that manage local coordinate space transforms for 3D objects
//
// See: https://developer.apple.com/documentation/ModelIO/MDLTransformComponent
type MDLTransformComponent interface {
	objectivec.IObject
	MDLComponent
}

// MDLTransformComponentObject wraps an existing Objective-C object that conforms to the MDLTransformComponent protocol.
type MDLTransformComponentObject struct {
	objectivec.Object
}

func (o MDLTransformComponentObject) BaseObject() objectivec.Object {
	return o.Object
}

// MDLTransformComponentObjectFromID constructs a [MDLTransformComponentObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MDLTransformComponentObjectFromID(id objc.ID) MDLTransformComponentObject {
	return MDLTransformComponentObject{
		Object: objectivec.ObjectFromID(id),
	}
}
