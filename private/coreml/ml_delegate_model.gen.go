// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLDelegateModel] class.
var (
	_MLDelegateModelClass     MLDelegateModelClass
	_MLDelegateModelClassOnce sync.Once
)

func getMLDelegateModelClass() MLDelegateModelClass {
	_MLDelegateModelClassOnce.Do(func() {
		_MLDelegateModelClass = MLDelegateModelClass{class: objc.GetClass("MLDelegateModel")}
	})
	return _MLDelegateModelClass
}

// GetMLDelegateModelClass returns the class object for MLDelegateModel.
func GetMLDelegateModelClass() MLDelegateModelClass {
	return getMLDelegateModelClass()
}

type MLDelegateModelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLDelegateModelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLDelegateModelClass) Alloc() MLDelegateModel {
	rv := objc.Send[MLDelegateModel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLDelegateModel._finishPredictionAndDispatchPendingPredictions]
//   - [MLDelegateModel._predictionFromFeaturesUsingStateOptionsCompletionHandler]
//   - [MLDelegateModel._predictionFromFeaturesUsingStateOptionsError]
//   - [MLDelegateModel._predictionsFromBatchOptionsError]
//   - [MLDelegateModel._schedulePredictionRequestCompletionHandler]
//   - [MLDelegateModel._submitPredictionRequestCompletionHandler]
//   - [MLDelegateModel._validateStateFeatureNamedBackingMultiArray]
//   - [MLDelegateModel.Engine]
//   - [MLDelegateModel.MaxAsyncPredictionsInFlight]
//   - [MLDelegateModel.ParameterValueForKeyError]
//   - [MLDelegateModel.PendingPredictionQueue]
//   - [MLDelegateModel.PredictionFromFeaturesOptionsError]
//   - [MLDelegateModel.PredictionFromFeaturesUsingStateOptionsError]
//   - [MLDelegateModel.PredictionsFromBatchOptionsError]
//   - [MLDelegateModel.InitWithEngineError]
// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel
type MLDelegateModel struct {
	MLModel
}

// MLDelegateModelFromID constructs a [MLDelegateModel] from an objc.ID.
func MLDelegateModelFromID(id objc.ID) MLDelegateModel {
	return MLDelegateModel{MLModel: MLModelFromID(id)}
}
// Ensure MLDelegateModel implements IMLDelegateModel.
var _ IMLDelegateModel = MLDelegateModel{}

// An interface definition for the [MLDelegateModel] class.
//
// # Methods
//
//   - [IMLDelegateModel._finishPredictionAndDispatchPendingPredictions]
//   - [IMLDelegateModel._predictionFromFeaturesUsingStateOptionsCompletionHandler]
//   - [IMLDelegateModel._predictionFromFeaturesUsingStateOptionsError]
//   - [IMLDelegateModel._predictionsFromBatchOptionsError]
//   - [IMLDelegateModel._schedulePredictionRequestCompletionHandler]
//   - [IMLDelegateModel._submitPredictionRequestCompletionHandler]
//   - [IMLDelegateModel._validateStateFeatureNamedBackingMultiArray]
//   - [IMLDelegateModel.Engine]
//   - [IMLDelegateModel.MaxAsyncPredictionsInFlight]
//   - [IMLDelegateModel.ParameterValueForKeyError]
//   - [IMLDelegateModel.PendingPredictionQueue]
//   - [IMLDelegateModel.PredictionFromFeaturesOptionsError]
//   - [IMLDelegateModel.PredictionFromFeaturesUsingStateOptionsError]
//   - [IMLDelegateModel.PredictionsFromBatchOptionsError]
//   - [IMLDelegateModel.InitWithEngineError]
//
// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel
type IMLDelegateModel interface {
	IMLModel

	// Topic: Methods

	_finishPredictionAndDispatchPendingPredictions()
	_predictionFromFeaturesUsingStateOptionsCompletionHandler(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject, handler ErrorHandler)
	_predictionFromFeaturesUsingStateOptionsError(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	_predictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	_schedulePredictionRequestCompletionHandler(request objectivec.IObject, handler ErrorHandler)
	_submitPredictionRequestCompletionHandler(request objectivec.IObject, handler ErrorHandler)
	_validateStateFeatureNamedBackingMultiArray(named objectivec.IObject, array objectivec.IObject)
	Engine() int
	MaxAsyncPredictionsInFlight() uint64
	ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error)
	PendingPredictionQueue() foundation.INSArray
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	PredictionFromFeaturesUsingStateOptionsError(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	InitWithEngineError(engine objectivec.IObject) (MLDelegateModel, error)
}

// Init initializes the instance.
func (d MLDelegateModel) Init() MLDelegateModel {
	rv := objc.Send[MLDelegateModel](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d MLDelegateModel) Autorelease() MLDelegateModel {
	rv := objc.Send[MLDelegateModel](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLDelegateModel creates a new MLDelegateModel instance.
func NewMLDelegateModel() MLDelegateModel {
	class := getMLDelegateModelClass()
	rv := objc.Send[MLDelegateModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func NewDelegateModelDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLDelegateModel, error) {
	var errorPtr objc.ID
	instance := getMLDelegateModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLDelegateModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLDelegateModelFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func NewDelegateModelInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLDelegateModel, error) {
	var errorPtr objc.ID
	instance := getMLDelegateModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLDelegateModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLDelegateModelFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewDelegateModelWithConfiguration(configuration objectivec.IObject) MLDelegateModel {
	instance := getMLDelegateModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLDelegateModelFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewDelegateModelWithDescription(description objectivec.IObject) MLDelegateModel {
	instance := getMLDelegateModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLDelegateModelFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewDelegateModelWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLDelegateModel {
	instance := getMLDelegateModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLDelegateModelFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/initWithEngine:error:
func NewDelegateModelWithEngineError(engine objectivec.IObject) (MLDelegateModel, error) {
	var errorPtr objc.ID
	instance := getMLDelegateModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEngine:error:"), engine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLDelegateModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLDelegateModelFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewDelegateModelWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLDelegateModel {
	instance := getMLDelegateModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLDelegateModelFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/_finishPredictionAndDispatchPendingPredictions
func (d MLDelegateModel) _finishPredictionAndDispatchPendingPredictions() {
	objc.Send[objc.ID](d.ID, objc.Sel("_finishPredictionAndDispatchPendingPredictions"))
}

// FinishPredictionAndDispatchPendingPredictions is an exported wrapper for the private method _finishPredictionAndDispatchPendingPredictions.
func (d MLDelegateModel) FinishPredictionAndDispatchPendingPredictions() {
	d._finishPredictionAndDispatchPendingPredictions()
}
//
// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/_predictionFromFeatures:usingState:options:completionHandler:
func (d MLDelegateModel) _predictionFromFeaturesUsingStateOptionsCompletionHandler(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject, handler ErrorHandler) {
_block3, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](d.ID, objc.Sel("_predictionFromFeatures:usingState:options:completionHandler:"), features, state, options, _block3)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/_predictionFromFeatures:usingState:options:error:
func (d MLDelegateModel) _predictionFromFeaturesUsingStateOptionsError(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("_predictionFromFeatures:usingState:options:error:"), features, state, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/_predictionsFromBatch:options:error:
func (d MLDelegateModel) _predictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("_predictionsFromBatch:options:error:"), batch, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/_schedulePredictionRequest:completionHandler:
func (d MLDelegateModel) _schedulePredictionRequestCompletionHandler(request objectivec.IObject, handler ErrorHandler) {
_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](d.ID, objc.Sel("_schedulePredictionRequest:completionHandler:"), request, _block1)
}

// SchedulePredictionRequestCompletionHandler is an exported wrapper for the private method _schedulePredictionRequestCompletionHandler.
func (d MLDelegateModel) SchedulePredictionRequestCompletionHandler(request objectivec.IObject, handler ErrorHandler) {
	d._schedulePredictionRequestCompletionHandler(request, handler)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/_submitPredictionRequest:completionHandler:
func (d MLDelegateModel) _submitPredictionRequestCompletionHandler(request objectivec.IObject, handler ErrorHandler) {
_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](d.ID, objc.Sel("_submitPredictionRequest:completionHandler:"), request, _block1)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/_validateStateFeatureNamed:backingMultiArray:
func (d MLDelegateModel) _validateStateFeatureNamedBackingMultiArray(named objectivec.IObject, array objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("_validateStateFeatureNamed:backingMultiArray:"), named, array)
}

// ValidateStateFeatureNamedBackingMultiArray is an exported wrapper for the private method _validateStateFeatureNamedBackingMultiArray.
func (d MLDelegateModel) ValidateStateFeatureNamedBackingMultiArray(named objectivec.IObject, array objectivec.IObject) {
	d._validateStateFeatureNamedBackingMultiArray(named, array)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/parameterValueForKey:error:
func (d MLDelegateModel) ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("parameterValueForKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/predictionFromFeatures:options:error:
func (d MLDelegateModel) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/predictionFromFeatures:usingState:options:error:
func (d MLDelegateModel) PredictionFromFeaturesUsingStateOptionsError(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("predictionFromFeatures:usingState:options:error:"), features, state, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/predictionsFromBatch:options:error:
func (d MLDelegateModel) PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("predictionsFromBatch:options:error:"), batch, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/initWithEngine:error:
func (d MLDelegateModel) InitWithEngineError(engine objectivec.IObject) (MLDelegateModel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithEngine:error:"), engine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLDelegateModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLDelegateModelFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/engine
func (d MLDelegateModel) Engine() int {
	rv := objc.Send[int](d.ID, objc.Sel("engine"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/maxAsyncPredictionsInFlight
func (d MLDelegateModel) MaxAsyncPredictionsInFlight() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("maxAsyncPredictionsInFlight"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/pendingPredictionQueue
func (d MLDelegateModel) PendingPredictionQueue() foundation.INSArray {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("pendingPredictionQueue"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// _predictionFromFeaturesUsingStateOptions is a synchronous wrapper around [MLDelegateModel._predictionFromFeaturesUsingStateOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d MLDelegateModel) _predictionFromFeaturesUsingStateOptions(ctx context.Context, features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) error {
	done := make(chan error, 1)
	d._predictionFromFeaturesUsingStateOptionsCompletionHandler(features, state, options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// _schedulePredictionRequest is a synchronous wrapper around [MLDelegateModel._schedulePredictionRequestCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d MLDelegateModel) _schedulePredictionRequest(ctx context.Context, request objectivec.IObject) error {
	done := make(chan error, 1)
	d._schedulePredictionRequestCompletionHandler(request, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// _submitPredictionRequest is a synchronous wrapper around [MLDelegateModel._submitPredictionRequestCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d MLDelegateModel) _submitPredictionRequest(ctx context.Context, request objectivec.IObject) error {
	done := make(chan error, 1)
	d._submitPredictionRequestCompletionHandler(request, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

