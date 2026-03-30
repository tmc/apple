// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLSupervisedOnlineUpdatable protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLSupervisedOnlineUpdatable
type MLSupervisedOnlineUpdatable interface {
	objectivec.IObject
}

// MLSupervisedOnlineUpdatableObject wraps an existing Objective-C object that conforms to the MLSupervisedOnlineUpdatable protocol.
type MLSupervisedOnlineUpdatableObject struct {
	objectivec.Object
}

func (o MLSupervisedOnlineUpdatableObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLSupervisedOnlineUpdatableObjectFromID constructs a [MLSupervisedOnlineUpdatableObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLSupervisedOnlineUpdatableObjectFromID(id objc.ID) MLSupervisedOnlineUpdatableObject {
	return MLSupervisedOnlineUpdatableObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLSupervisedOnlineUpdatable/updateModelFromFeatures:toTarget:options:error:
func (o MLSupervisedOnlineUpdatableObject) UpdateModelFromFeaturesToTargetOptionsError(features objectivec.IObject, target objectivec.IObject, options objectivec.IObject) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("updateModelFromFeatures:toTarget:options:error:"), features, target, options)
	if err != nil {
		return false, err
	}
	return rv, nil
}
