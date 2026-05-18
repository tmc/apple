// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
//   - [MLDelegateModel.PredictionFromFeaturesOptionsCompletionHandler]
//   - [MLDelegateModel.PredictionFromFeaturesOptionsError]
//   - [MLDelegateModel.PredictionFromFeaturesUsingStateOptionsCompletionHandler]
//   - [MLDelegateModel.PredictionFromFeaturesUsingStateOptionsError]
//   - [MLDelegateModel.PredictionsFromBatchOptionsError]
//   - [MLDelegateModel.InitWithEngineError]
//
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
//   - [IMLDelegateModel.PredictionFromFeaturesOptionsCompletionHandler]
//   - [IMLDelegateModel.PredictionFromFeaturesOptionsError]
//   - [IMLDelegateModel.PredictionFromFeaturesUsingStateOptionsCompletionHandler]
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
	PredictionFromFeaturesOptionsCompletionHandler(features objectivec.IObject, options objectivec.IObject, handler ErrorHandler)
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	PredictionFromFeaturesUsingStateOptionsCompletionHandler(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject, handler ErrorHandler)
	PredictionFromFeaturesUsingStateOptionsError(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	InitWithEngineError(engine objectivec.IObject) (MLDelegateModel, error)
}

// Init initializes the instance.
func (m MLDelegateModel) Init() MLDelegateModel {
	rv := objc.Send[MLDelegateModel](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLDelegateModel) Autorelease() MLDelegateModel {
	rv := objc.Send[MLDelegateModel](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLDelegateModel creates a new MLDelegateModel instance.
func NewMLDelegateModel() MLDelegateModel {
	class := getMLDelegateModelClass()
	rv := objc.Send[MLDelegateModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

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

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewDelegateModelWithConfiguration(configuration objectivec.IObject) MLDelegateModel {
	instance := getMLDelegateModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLDelegateModelFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewDelegateModelWithDescription(description objectivec.IObject) MLDelegateModel {
	instance := getMLDelegateModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLDelegateModelFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewDelegateModelWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLDelegateModel {
	instance := getMLDelegateModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLDelegateModelFromID(rv)
}

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

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewDelegateModelWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLDelegateModel {
	instance := getMLDelegateModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLDelegateModelFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/_finishPredictionAndDispatchPendingPredictions
func (m MLDelegateModel) _finishPredictionAndDispatchPendingPredictions() {
	objc.Send[objc.ID](m.ID, objc.Sel("_finishPredictionAndDispatchPendingPredictions"))
}

// FinishPredictionAndDispatchPendingPredictions is an exported wrapper for the private method _finishPredictionAndDispatchPendingPredictions.
func (m MLDelegateModel) FinishPredictionAndDispatchPendingPredictions() error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_finishPredictionAndDispatchPendingPredictions")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_finishPredictionAndDispatchPendingPredictions"}
		return err
	}
	m._finishPredictionAndDispatchPendingPredictions()
	return nil
}

// CanFinishPredictionAndDispatchPendingPredictions reports whether the receiver responds to the private selector _finishPredictionAndDispatchPendingPredictions.
func (m MLDelegateModel) CanFinishPredictionAndDispatchPendingPredictions() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_finishPredictionAndDispatchPendingPredictions"))
}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/_predictionFromFeatures:usingState:options:completionHandler:
func (m MLDelegateModel) _predictionFromFeaturesUsingStateOptionsCompletionHandler(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject, handler ErrorHandler) {
	_block3, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("_predictionFromFeatures:usingState:options:completionHandler:"), features, state, options, _block3)
}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/_predictionFromFeatures:usingState:options:error:
func (m MLDelegateModel) _predictionFromFeaturesUsingStateOptionsError(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_predictionFromFeatures:usingState:options:error:"), features, state, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/_predictionsFromBatch:options:error:
func (m MLDelegateModel) _predictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_predictionsFromBatch:options:error:"), batch, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/_schedulePredictionRequest:completionHandler:
func (m MLDelegateModel) _schedulePredictionRequestCompletionHandler(request objectivec.IObject, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("_schedulePredictionRequest:completionHandler:"), request, _block1)
}

// SchedulePredictionRequestCompletionHandler is an exported wrapper for the private method _schedulePredictionRequestCompletionHandler.
func (m MLDelegateModel) SchedulePredictionRequestCompletionHandler(request objectivec.IObject, handler ErrorHandler) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_schedulePredictionRequest:completionHandler:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_schedulePredictionRequest:completionHandler:"}
		return err
	}
	m._schedulePredictionRequestCompletionHandler(request, handler)
	return nil
}

// CanSchedulePredictionRequestCompletionHandler reports whether the receiver responds to the private selector _schedulePredictionRequest:completionHandler:.
func (m MLDelegateModel) CanSchedulePredictionRequestCompletionHandler() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_schedulePredictionRequest:completionHandler:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/_submitPredictionRequest:completionHandler:
func (m MLDelegateModel) _submitPredictionRequestCompletionHandler(request objectivec.IObject, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("_submitPredictionRequest:completionHandler:"), request, _block1)
}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/_validateStateFeatureNamed:backingMultiArray:
func (m MLDelegateModel) _validateStateFeatureNamedBackingMultiArray(named objectivec.IObject, array objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("_validateStateFeatureNamed:backingMultiArray:"), named, array)
}

// ValidateStateFeatureNamedBackingMultiArray is an exported wrapper for the private method _validateStateFeatureNamedBackingMultiArray.
func (m MLDelegateModel) ValidateStateFeatureNamedBackingMultiArray(named objectivec.IObject, array objectivec.IObject) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_validateStateFeatureNamed:backingMultiArray:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_validateStateFeatureNamed:backingMultiArray:"}
		return err
	}
	m._validateStateFeatureNamedBackingMultiArray(named, array)
	return nil
}

// CanValidateStateFeatureNamedBackingMultiArray reports whether the receiver responds to the private selector _validateStateFeatureNamed:backingMultiArray:.
func (m MLDelegateModel) CanValidateStateFeatureNamedBackingMultiArray() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_validateStateFeatureNamed:backingMultiArray:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/parameterValueForKey:error:
func (m MLDelegateModel) ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parameterValueForKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/predictionFromFeatures:options:completionHandler:
func (m MLDelegateModel) PredictionFromFeaturesOptionsCompletionHandler(features objectivec.IObject, options objectivec.IObject, handler ErrorHandler) {
	_block2, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:options:completionHandler:"), features, options, _block2)
}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/predictionFromFeatures:options:error:
func (m MLDelegateModel) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/predictionFromFeatures:usingState:options:completionHandler:
func (m MLDelegateModel) PredictionFromFeaturesUsingStateOptionsCompletionHandler(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject, handler ErrorHandler) {
	_block3, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:usingState:options:completionHandler:"), features, state, options, _block3)
}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/predictionFromFeatures:usingState:options:error:
func (m MLDelegateModel) PredictionFromFeaturesUsingStateOptionsError(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:usingState:options:error:"), features, state, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/predictionsFromBatch:options:error:
func (m MLDelegateModel) PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionsFromBatch:options:error:"), batch, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/initWithEngine:error:
func (m MLDelegateModel) InitWithEngineError(engine objectivec.IObject) (MLDelegateModel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithEngine:error:"), engine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLDelegateModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLDelegateModelFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/engine
func (m MLDelegateModel) Engine() int {
	rv := objc.Send[int](m.ID, objc.Sel("engine"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/maxAsyncPredictionsInFlight
func (m MLDelegateModel) MaxAsyncPredictionsInFlight() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("maxAsyncPredictionsInFlight"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateModel/pendingPredictionQueue
func (m MLDelegateModel) PendingPredictionQueue() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("pendingPredictionQueue"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// _predictionFromFeaturesUsingStateOptions is a synchronous wrapper around [MLDelegateModel._predictionFromFeaturesUsingStateOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLDelegateModel) _predictionFromFeaturesUsingStateOptions(ctx context.Context, features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) error {
	done := make(chan error, 1)
	m._predictionFromFeaturesUsingStateOptionsCompletionHandler(features, state, options, func(err error) {
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
func (m MLDelegateModel) _schedulePredictionRequest(ctx context.Context, request objectivec.IObject) error {
	done := make(chan error, 1)
	m._schedulePredictionRequestCompletionHandler(request, func(err error) {
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
func (m MLDelegateModel) _submitPredictionRequest(ctx context.Context, request objectivec.IObject) error {
	done := make(chan error, 1)
	m._submitPredictionRequestCompletionHandler(request, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PredictionFromFeaturesOptions is a synchronous wrapper around [MLDelegateModel.PredictionFromFeaturesOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLDelegateModel) PredictionFromFeaturesOptions(ctx context.Context, features objectivec.IObject, options objectivec.IObject) error {
	done := make(chan error, 1)
	m.PredictionFromFeaturesOptionsCompletionHandler(features, options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PredictionFromFeaturesUsingStateOptions is a synchronous wrapper around [MLDelegateModel.PredictionFromFeaturesUsingStateOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLDelegateModel) PredictionFromFeaturesUsingStateOptions(ctx context.Context, features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) error {
	done := make(chan error, 1)
	m.PredictionFromFeaturesUsingStateOptionsCompletionHandler(features, state, options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
