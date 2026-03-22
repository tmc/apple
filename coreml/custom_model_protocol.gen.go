// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An interface that defines the behavior of a custom model.
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomModel
type MLCustomModel interface {
	objectivec.IObject

	// Predicts output values from the given input features.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLCustomModel/prediction(from:options:)
	PredictionFromFeaturesOptionsError(input MLFeatureProvider, options IMLPredictionOptions) (MLFeatureProvider, error)
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

// Predicts output values from the given input features.
//
// input: The feature values the models needs to make its prediction.
//
// options: The options to be applied to the prediction.
//
// # Return Value
// 
// A feature provider that represents the model’s prediction.
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomModel/prediction(from:options:)
func (o MLCustomModelObject) PredictionFromFeaturesOptionsError(input MLFeatureProvider, options IMLPredictionOptions) (MLFeatureProvider, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("predictionFromFeatures:options:error:"), input, options)
	if err != nil {
		return nil, err
	}
	return MLFeatureProviderObjectFromID(rv), nil
	}
// Predicts output values from the given batch of input features.
//
// inputBatch: The batch of feature values the model needs to make its predictions.
//
// options: The options to be applied to the predictions.
//
// # Return Value
// 
// A batch provider that represents the model’s predictions for the batch of
// inputs.
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomModel/predictions(from:options:)
func (o MLCustomModelObject) PredictionsFromBatchOptionsError(inputBatch MLBatchProvider, options IMLPredictionOptions) (MLBatchProvider, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("predictionsFromBatch:options:error:"), inputBatch, options)
	if err != nil {
		return nil, err
	}
	return MLBatchProviderObjectFromID(rv), nil
	}

