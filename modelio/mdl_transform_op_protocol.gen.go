// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MDLTransformOp protocol.
//
// See: https://developer.apple.com/documentation/ModelIO/MDLTransformOp
type MDLTransformOp interface {
	objectivec.IObject
}

// MDLTransformOpObject wraps an existing Objective-C object that conforms to the MDLTransformOp protocol.
type MDLTransformOpObject struct {
	objectivec.Object
}

func (o MDLTransformOpObject) BaseObject() objectivec.Object {
	return o.Object
}

// MDLTransformOpObjectFromID constructs a [MDLTransformOpObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MDLTransformOpObjectFromID(id objc.ID) MDLTransformOpObject {
	return MDLTransformOpObject{
		Object: objectivec.ObjectFromID(id),
	}
}
