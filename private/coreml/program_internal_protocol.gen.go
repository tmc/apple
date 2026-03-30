// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLProgramInternal protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLProgramInternal
type MLProgramInternal interface {
	objectivec.IObject
}

// MLProgramInternalObject wraps an existing Objective-C object that conforms to the MLProgramInternal protocol.
type MLProgramInternalObject struct {
	objectivec.Object
}

func (o MLProgramInternalObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLProgramInternalObjectFromID constructs a [MLProgramInternalObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLProgramInternalObjectFromID(id objc.ID) MLProgramInternalObject {
	return MLProgramInternalObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramInternal/evaluateFunction:arguments:error:
func (o MLProgramInternalObject) EvaluateFunctionArgumentsError(function objectivec.IObject, arguments objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("evaluateFunction:arguments:error:"), function, arguments)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramInternal/newContextAndReturnError:
func (o MLProgramInternalObject) NewContextAndReturnError() (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newContextAndReturnError:"))
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
