// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLModeling protocol.
type MLModeling interface {
	objectivec.IObject

	// Configuration protocol.
	Configuration() objectivec.IObject

	// EnableInstrumentsTracing protocol.
	EnableInstrumentsTracing()

	// ExecutionSchedule protocol.
	ExecutionSchedule() objectivec.IObject

	// Metadata protocol.
	Metadata() objectivec.IObject

	// ModelDescription protocol.
	ModelDescription() objectivec.IObject

	// ModelPath protocol.
	ModelPath() objectivec.IObject

	// NewRequestForModelInputFeaturesOptionsError protocol.
	NewRequestForModelInputFeaturesOptionsError(model objectivec.IObject, features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)

	// ParameterValueForKeyError protocol.
	ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error)

	// PredictionFromFeaturesError protocol.
	PredictionFromFeaturesError(features objectivec.IObject) (objectivec.IObject, error)

	// PredictionFromFeaturesOptionsError protocol.
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)

	// PredictionTypeForKTrace protocol.
	PredictionTypeForKTrace() uint64

	// PredictionsFromBatchError protocol.
	PredictionsFromBatchError(batch objectivec.IObject) (objectivec.IObject, error)

	// PredictionsFromBatchOptionsError protocol.
	PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)

	// RecordsPredictionEvent protocol.
	RecordsPredictionEvent() bool

	// SetModelPathModelName protocol.
	SetModelPathModelName(path objectivec.IObject, name objectivec.IObject)

	// SignpostID protocol.
	SignpostID() uint64

	// SupportsConcurrentSubmissions protocol.
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

func (o MLModelingObject) Configuration() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("configuration"))
	return objectivec.Object{ID: rv}
}
func (o MLModelingObject) EnableInstrumentsTracing() {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("enableInstrumentsTracing"))
}
func (o MLModelingObject) ExecutionSchedule() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("executionSchedule"))
	return objectivec.Object{ID: rv}
}
func (o MLModelingObject) Metadata() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("metadata"))
	return objectivec.Object{ID: rv}
}
func (o MLModelingObject) ModelDescription() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("modelDescription"))
	return objectivec.Object{ID: rv}
}
func (o MLModelingObject) ModelPath() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("modelPath"))
	return objectivec.Object{ID: rv}
}
func (o MLModelingObject) NewRequestForModelInputFeaturesOptionsError(model objectivec.IObject, features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newRequestForModel:inputFeatures:options:error:"), model, features, options)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
func (o MLModelingObject) ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("parameterValueForKey:error:"), key)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
func (o MLModelingObject) PredictionFromFeaturesError(features objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("predictionFromFeatures:error:"), features)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
func (o MLModelingObject) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
func (o MLModelingObject) PredictionTypeForKTrace() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("predictionTypeForKTrace"))
	return rv
}
func (o MLModelingObject) PredictionsFromBatchError(batch objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("predictionsFromBatch:error:"), batch)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
func (o MLModelingObject) PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("predictionsFromBatch:options:error:"), batch, options)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
func (o MLModelingObject) RecordsPredictionEvent() bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("recordsPredictionEvent"))
	return rv
}
func (o MLModelingObject) SetModelPathModelName(path objectivec.IObject, name objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setModelPath:modelName:"), path, name)
}
func (o MLModelingObject) SignpostID() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("signpostID"))
	return rv
}
func (o MLModelingObject) SupportsConcurrentSubmissions() bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("supportsConcurrentSubmissions"))
	return rv
}
