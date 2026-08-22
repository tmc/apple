// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MDLJointAnimation protocol.
//
// See: https://developer.apple.com/documentation/ModelIO/MDLJointAnimation
type MDLJointAnimation interface {
	objectivec.IObject
}

// MDLJointAnimationObject wraps an existing Objective-C object that conforms to the MDLJointAnimation protocol.
type MDLJointAnimationObject struct {
	objectivec.Object
}

func (o MDLJointAnimationObject) BaseObject() objectivec.Object {
	return o.Object
}

// MDLJointAnimationObjectFromID constructs a [MDLJointAnimationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MDLJointAnimationObjectFromID(id objc.ID) MDLJointAnimationObject {
	return MDLJointAnimationObject{
		Object: objectivec.ObjectFromID(id),
	}
}
