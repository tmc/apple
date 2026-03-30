// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLStatefulModelEngine protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLStatefulModelEngine
type MLStatefulModelEngine interface {
	objectivec.IObject
}

// MLStatefulModelEngineObject wraps an existing Objective-C object that conforms to the MLStatefulModelEngine protocol.
type MLStatefulModelEngineObject struct {
	objectivec.Object
}

func (o MLStatefulModelEngineObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLStatefulModelEngineObjectFromID constructs a [MLStatefulModelEngineObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLStatefulModelEngineObjectFromID(id objc.ID) MLStatefulModelEngineObject {
	return MLStatefulModelEngineObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLStatefulModelEngine/newRequestForModel:inputFeatures:usingState:options:error:
func (o MLStatefulModelEngineObject) NewRequestForModelInputFeaturesUsingStateOptionsError(model objectivec.IObject, features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newRequestForModel:inputFeatures:usingState:options:error:"), model, features, state, options)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}

// See: https://developer.apple.com/documentation/CoreML/MLStatefulModelEngine/newStateWithClientBuffers:
func (o MLStatefulModelEngineObject) NewStateWithClientBuffers(buffers objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newStateWithClientBuffers:"), buffers)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLStatefulModelEngine/predictionFromFeatures:usingState:options:error:
func (o MLStatefulModelEngineObject) PredictionFromFeaturesUsingStateOptionsError(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("predictionFromFeatures:usingState:options:error:"), features, state, options)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
