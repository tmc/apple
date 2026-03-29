// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLCustomModel protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomModel
type MLCustomModel interface {
	objectivec.IObject
}

// MLCustomModelObject wraps an existing Objective-C object that conforms to the MLCustomModel protocol.
type MLCustomModelObject struct {
	objectivec.Object
}
func (o MLCustomModelObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLCustomModelObjectFromID constructs a [MLCustomModelObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLCustomModelObjectFromID(id objc.ID) MLCustomModelObject {
	return MLCustomModelObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/CoreML/MLCustomModel/predictionFromFeatures:options:error:
func (o MLCustomModelObject) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
	}
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomModel/predictionsFromBatch:options:error:
func (o MLCustomModelObject) PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("predictionsFromBatch:options:error:"), batch, options)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
	}

