// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLRegressor protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLRegressor
type MLRegressorProtocol interface {
	objectivec.IObject
}

// MLRegressorProtocolObject wraps an existing Objective-C object that conforms to the MLRegressorProtocol protocol.
type MLRegressorProtocolObject struct {
	objectivec.Object
}

func (o MLRegressorProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLRegressorProtocolObjectFromID constructs a [MLRegressorProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLRegressorProtocolObjectFromID(id objc.ID) MLRegressorProtocolObject {
	return MLRegressorProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLRegressor/regress:options:error:
func (o MLRegressorProtocolObject) RegressOptionsError(regress objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("regress:options:error:"), regress, options)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
