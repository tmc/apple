// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLFeatureValueConstraint protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValueConstraint
type MLFeatureValueConstraint interface {
	objectivec.IObject
}

// MLFeatureValueConstraintObject wraps an existing Objective-C object that conforms to the MLFeatureValueConstraint protocol.
type MLFeatureValueConstraintObject struct {
	objectivec.Object
}
func (o MLFeatureValueConstraintObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLFeatureValueConstraintObjectFromID constructs a [MLFeatureValueConstraintObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLFeatureValueConstraintObjectFromID(id objc.ID) MLFeatureValueConstraintObject {
	return MLFeatureValueConstraintObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValueConstraint/isAllowedValue:error:
func (o MLFeatureValueConstraintObject) IsAllowedValueError(value objectivec.IObject) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("isAllowedValue:error:"), value)
	if err != nil {
		return false, err
	}
	return rv, nil
	}

