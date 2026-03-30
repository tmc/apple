// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLE5ExecutionStreamOperationPool protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperationPool
type MLE5ExecutionStreamOperationPool interface {
	objectivec.IObject

	// PrepareWithInitialPoolSizeError protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperationPool/prepareWithInitialPoolSize:error:
	PrepareWithInitialPoolSizeError(size int64) (bool, error)
}

// MLE5ExecutionStreamOperationPoolObject wraps an existing Objective-C object that conforms to the MLE5ExecutionStreamOperationPool protocol.
type MLE5ExecutionStreamOperationPoolObject struct {
	objectivec.Object
}

func (o MLE5ExecutionStreamOperationPoolObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLE5ExecutionStreamOperationPoolObjectFromID constructs a [MLE5ExecutionStreamOperationPoolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLE5ExecutionStreamOperationPoolObjectFromID(id objc.ID) MLE5ExecutionStreamOperationPoolObject {
	return MLE5ExecutionStreamOperationPoolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperationPool/prepareWithInitialPoolSize:error:
func (o MLE5ExecutionStreamOperationPoolObject) PrepareWithInitialPoolSizeError(size int64) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("prepareWithInitialPoolSize:error:"), size)
	if err != nil {
		return false, err
	}
	return rv, nil
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperationPool/putBack:
func (o MLE5ExecutionStreamOperationPoolObject) PutBack(back objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("putBack:"), back)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperationPool/takeOutOperationForFeatures:error:
func (o MLE5ExecutionStreamOperationPoolObject) TakeOutOperationForFeaturesError(features objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("takeOutOperationForFeatures:error:"), features)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
