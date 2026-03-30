// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLModeling protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLModeling
type MLModeling interface {
	objectivec.IObject

	// EnableInstrumentsTracing protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLModeling/enableInstrumentsTracing
	EnableInstrumentsTracing()

	// PredictionTypeForKTrace protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLModeling/predictionTypeForKTrace
	PredictionTypeForKTrace() uint64

	// RecordsPredictionEvent protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLModeling/recordsPredictionEvent
	RecordsPredictionEvent() bool

	// SignpostID protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLModeling/signpostID
	SignpostID() uint64

	// SupportsConcurrentSubmissions protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLModeling/supportsConcurrentSubmissions
	SupportsConcurrentSubmissions() bool
}

// MLModelingObject wraps an existing Objective-C object that conforms to the MLModeling protocol.
type MLModelingObject struct {
	objectivec.Object
}

func (o MLModelingObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLModelingObjectFromID constructs a [MLModelingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLModelingObjectFromID(id objc.ID) MLModelingObject {
	return MLModelingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLModeling/configuration
func (o MLModelingObject) Configuration() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("configuration"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLModeling/enableInstrumentsTracing
func (o MLModelingObject) EnableInstrumentsTracing() {
	objc.Send[struct{}](o.ID, objc.Sel("enableInstrumentsTracing"))
}

// See: https://developer.apple.com/documentation/CoreML/MLModeling/executionSchedule
func (o MLModelingObject) ExecutionSchedule() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("executionSchedule"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLModeling/metadata
func (o MLModelingObject) Metadata() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("metadata"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLModeling/modelDescription
func (o MLModelingObject) ModelDescription() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("modelDescription"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLModeling/modelPath
func (o MLModelingObject) ModelPath() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("modelPath"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLModeling/newRequestForModel:inputFeatures:options:error:
func (o MLModelingObject) NewRequestForModelInputFeaturesOptionsError(model objectivec.IObject, features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newRequestForModel:inputFeatures:options:error:"), model, features, options)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModeling/parameterValueForKey:error:
func (o MLModelingObject) ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("parameterValueForKey:error:"), key)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModeling/predictionFromFeatures:error:
func (o MLModelingObject) PredictionFromFeaturesError(features objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("predictionFromFeatures:error:"), features)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModeling/predictionFromFeatures:options:error:
func (o MLModelingObject) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModeling/predictionTypeForKTrace
func (o MLModelingObject) PredictionTypeForKTrace() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("predictionTypeForKTrace"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModeling/predictionsFromBatch:error:
func (o MLModelingObject) PredictionsFromBatchError(batch objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("predictionsFromBatch:error:"), batch)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModeling/predictionsFromBatch:options:error:
func (o MLModelingObject) PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("predictionsFromBatch:options:error:"), batch, options)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModeling/recordsPredictionEvent
func (o MLModelingObject) RecordsPredictionEvent() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("recordsPredictionEvent"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModeling/setModelPath:modelName:
func (o MLModelingObject) SetModelPathModelName(path objectivec.IObject, name objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setModelPath:modelName:"), path, name)
}

// See: https://developer.apple.com/documentation/CoreML/MLModeling/signpostID
func (o MLModelingObject) SignpostID() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("signpostID"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModeling/submitPredictionRequest:completionHandler:
func (o MLModelingObject) SubmitPredictionRequestCompletionHandler(request objectivec.IObject, handler MLFeatureProviderErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("submitPredictionRequest:completionHandler:"), request, handler)
}

// See: https://developer.apple.com/documentation/CoreML/MLModeling/supportsConcurrentSubmissions
func (o MLModelingObject) SupportsConcurrentSubmissions() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportsConcurrentSubmissions"))
	return rv
}
