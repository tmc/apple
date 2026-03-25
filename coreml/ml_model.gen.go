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

// The class instance for the [MLModel] class.
var (
	_MLModelClass     MLModelClass
	_MLModelClassOnce sync.Once
)

func getMLModelClass() MLModelClass {
	_MLModelClassOnce.Do(func() {
		_MLModelClass = MLModelClass{class: objc.GetClass("MLModel")}
	})
	return _MLModelClass
}

// GetMLModelClass returns the class object for MLModel.
func GetMLModelClass() MLModelClass {
	return getMLModelClass()
}

type MLModelClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelClass) Alloc() MLModel {
	rv := objc.Send[MLModel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An encapsulation of all the details of your machine learning model.
//
// # Overview
// 
// [MLModel] encapsulates a model’s prediction methods, configuration, and
// model description.
// 
// In most cases, you can use Core ML without accessing the [MLModel] class
// directly. Instead, use the programmer-friendly wrapper class that Xcode
// automatically generates when you add a model (see [Integrating a Core ML
// Model into Your App]). If your app needs the [MLModel] interface, use the
// wrapper class’s `model` property.
// 
// With the [MLModel] interface, you can:
// 
// - Make a prediction with your app’s custom [MLFeatureProvider] by calling
// [prediction(from:)] or [prediction(from:options:)]. - Make multiple
// predictions with your app’s custom [MLBatchProvider] by calling
// [MLModel.PredictionsFromBatchError] or [MLModel.PredictionsFromBatchOptionsError]. -
// Inspect your model’s [MLModel.Metadata] and [MLFeatureDescription] instances
// through [ModelDescription].
// 
// If your app downloads and compiles a model on the user’s device, you must
// use the [MLModel] class directly to make predictions. See [Downloading and
// Compiling a Model on the User’s Device].
//
// [Downloading and Compiling a Model on the User’s Device]: https://developer.apple.com/documentation/CoreML/downloading-and-compiling-a-model-on-the-user-s-device
// [Integrating a Core ML Model into Your App]: https://developer.apple.com/documentation/CoreML/integrating-a-core-ml-model-into-your-app
// [prediction(from:)]: https://developer.apple.com/documentation/CoreML/MLModel/prediction(from:)-9y2aa
// [prediction(from:options:)]: https://developer.apple.com/documentation/CoreML/MLModel/prediction(from:options:)-81mr6
//
// # Making predictions
//
//   - [MLModel.PredictionFromFeaturesError]
//   - [MLModel.PredictionFromFeaturesOptionsError]
//   - [MLModel.PredictionsFromBatchError]: Generates predictions for each input feature provider within the batch provider.
//   - [MLModel.PredictionsFromBatchOptionsError]: Generates a prediction for each input feature provider within the batch provider using the prediction options.
//   - [MLModel.PredictionFromFeaturesUsingStateError]
//   - [MLModel.PredictionFromFeaturesUsingStateOptionsError]
//
// # Inspecting a model
//
//   - [MLModel.AvailableComputeDevices]: The list of available compute devices that the model’s prediction methods use.
//   - [MLModel.SetAvailableComputeDevices]
//   - [MLModel.Configuration]: The configuration of the model set during initialization.
//   - [MLModel.ModelDescription]: Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//   - [MLModel.ParameterValueForKeyError]: Returns a model parameter value for a key.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel
type MLModel struct {
	objectivec.Object
}

// MLModelFromID constructs a [MLModel] from an objc.ID.
//
// An encapsulation of all the details of your machine learning model.
func MLModelFromID(id objc.ID) MLModel {
	return MLModel{objectivec.Object{ID: id}}
}
// NOTE: MLModel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLModel] class.
//
// # Making predictions
//
//   - [IMLModel.PredictionFromFeaturesError]
//   - [IMLModel.PredictionFromFeaturesOptionsError]
//   - [IMLModel.PredictionsFromBatchError]: Generates predictions for each input feature provider within the batch provider.
//   - [IMLModel.PredictionsFromBatchOptionsError]: Generates a prediction for each input feature provider within the batch provider using the prediction options.
//   - [IMLModel.PredictionFromFeaturesUsingStateError]
//   - [IMLModel.PredictionFromFeaturesUsingStateOptionsError]
//
// # Inspecting a model
//
//   - [IMLModel.AvailableComputeDevices]: The list of available compute devices that the model’s prediction methods use.
//   - [IMLModel.SetAvailableComputeDevices]
//   - [IMLModel.Configuration]: The configuration of the model set during initialization.
//   - [IMLModel.ModelDescription]: Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//   - [IMLModel.ParameterValueForKeyError]: Returns a model parameter value for a key.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel
type IMLModel interface {
	objectivec.IObject

	// Topic: Making predictions

	PredictionFromFeaturesError(input objectivec.IObject) (objectivec.IObject, error)
	PredictionFromFeaturesOptionsError(input objectivec.IObject, options IMLPredictionOptions) (objectivec.IObject, error)
	// Generates predictions for each input feature provider within the batch provider.
	PredictionsFromBatchError(inputBatch MLBatchProvider) (MLBatchProvider, error)
	// Generates a prediction for each input feature provider within the batch provider using the prediction options.
	PredictionsFromBatchOptionsError(inputBatch MLBatchProvider, options IMLPredictionOptions) (MLBatchProvider, error)
	PredictionFromFeaturesUsingStateError(inputFeatures objectivec.IObject, state IMLState) (objectivec.IObject, error)
	PredictionFromFeaturesUsingStateOptionsError(inputFeatures objectivec.IObject, state IMLState, options IMLPredictionOptions) (objectivec.IObject, error)

	// Topic: Inspecting a model

	// The list of available compute devices that the model’s prediction methods use.
	AvailableComputeDevices() objectivec.IObject
	SetAvailableComputeDevices(value objectivec.IObject)
	// The configuration of the model set during initialization.
	Configuration() IMLModelConfiguration
	// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
	ModelDescription() IMLModelDescription
	// Returns a model parameter value for a key.
	ParameterValueForKeyError(key IMLParameterKey) (objectivec.IObject, error)

	// A dictionary of the model’s creation information, such as its description, author, version, and license.
	Metadata() MLModelMetadataKey
	SetMetadata(value MLModelMetadataKey)
}

// Init initializes the instance.
func (m MLModel) Init() MLModel {
	rv := objc.Send[MLModel](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModel) Autorelease() MLModel {
	rv := objc.Send[MLModel](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModel creates a new MLModel instance.
func NewMLModel() MLModel {
	class := getMLModelClass()
	rv := objc.Send[MLModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a Core ML model instance from a compiled model file and a custom
// configuration.
//
// url: The path to a compiled model file (`XCUIElementTypeMlmodelc`), typically
// with the [URL] that [compileModel(at:)] returns.
// //
// [compileModel(at:)]: https://developer.apple.com/documentation/CoreML/MLModel/compileModel(at:)-6442s
//
// configuration: The runtime settings for the new model instance.
//
// # Discussion
// 
// In most cases, your app won’t need to create a model object directly.
// Consider the programmer-friendly wrapper class that Xcode automatically
// generates when you add a model to your project (see [Integrating a Core ML
// Model into Your App]).
// 
// If the wrapper class doesn’t meet your app’s needs or you need to
// customize the model’s configuration, use this initializer to create a
// model object from any compiled model file you can access. Typically, you
// use this initializer after your app has downloaded and compiled a model,
// which is one technique for saving space in your app (see [Downloading and
// Compiling a Model on the User’s Device]).
//
// [Downloading and Compiling a Model on the User’s Device]: https://developer.apple.com/documentation/CoreML/downloading-and-compiling-a-model-on-the-user-s-device
// [Integrating a Core ML Model into Your App]: https://developer.apple.com/documentation/CoreML/integrating-a-core-ml-model-into-your-app
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/init(contentsOf:configuration:)
func NewModelWithContentsOfURLConfigurationError(url foundation.INSURL, configuration IMLModelConfiguration) (MLModel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getMLModelClass().class), objc.Sel("modelWithContentsOfURL:configuration:error:"), url, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelFromID(rv), nil
}

// Creates a Core ML model instance from a compiled model file.
//
// url: The path to a compiled model file (`XCUIElementTypeMlmodelc`), typically
// with the [URL] that [compileModel(at:)] returns.
// //
// [compileModel(at:)]: https://developer.apple.com/documentation/CoreML/MLModel/compileModel(at:)-6442s
//
// # Discussion
// 
// In most cases, your app won’t need to create a model object directly.
// Consider the programmer-friendly wrapper class that Xcode automatically
// generates when you add a model to your project (see [Integrating a Core ML
// Model into Your App]).
// 
// If the wrapper class doesn’t meet your app’s needs or you need to
// customize the model’s configuration, use this initializer to create a
// model object from any compiled model file you can access. Typically, you
// use this initializer after your app has downloaded and compiled a model,
// which is one technique for saving space in your app (see [Downloading and
// Compiling a Model on the User’s Device]).
//
// [Downloading and Compiling a Model on the User’s Device]: https://developer.apple.com/documentation/CoreML/downloading-and-compiling-a-model-on-the-user-s-device
// [Integrating a Core ML Model into Your App]: https://developer.apple.com/documentation/CoreML/integrating-a-core-ml-model-into-your-app
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/init(contentsOf:)
func NewModelWithContentsOfURLError(url foundation.INSURL) (MLModel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getMLModelClass().class), objc.Sel("modelWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/prediction(from:)
func (m MLModel) PredictionFromFeaturesError(input objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:error:"), input, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/prediction(from:options:)
func (m MLModel) PredictionFromFeaturesOptionsError(input objectivec.IObject, options IMLPredictionOptions) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:options:error:"), input, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// Generates predictions for each input feature provider within the batch
// provider.
//
// inputBatch: A batch provider that contains multiple input feature providers. The model
// makes a prediction for each feature provider.
//
// # Return Value
// 
// A batch provider that contains an output feature provider for each
// prediction.
//
// # Discussion
// 
// Use this method to make more than one prediction at one time.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/predictions(fromBatch:)
func (m MLModel) PredictionsFromBatchError(inputBatch MLBatchProvider) (MLBatchProvider, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionsFromBatch:error:"), inputBatch, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return MLBatchProviderObjectFromID(rv), nil

}
// Generates a prediction for each input feature provider within the batch
// provider using the prediction options.
//
// inputBatch: A batch provider that contains multiple input feature providers. The model
// makes a prediction for each feature provider.
//
// options: The runtime settings the model uses as it makes a prediction.
//
// # Return Value
// 
// A batch provider that contains an output feature provider for each
// prediction.
//
// # Discussion
// 
// Use this method to make more than one prediction at one time.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/predictions(from:options:)
func (m MLModel) PredictionsFromBatchOptionsError(inputBatch MLBatchProvider, options IMLPredictionOptions) (MLBatchProvider, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionsFromBatch:options:error:"), inputBatch, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return MLBatchProviderObjectFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/prediction(from:using:)
func (m MLModel) PredictionFromFeaturesUsingStateError(inputFeatures objectivec.IObject, state IMLState) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:usingState:error:"), inputFeatures, state, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/prediction(from:using:options:)
func (m MLModel) PredictionFromFeaturesUsingStateOptionsError(inputFeatures objectivec.IObject, state IMLState, options IMLPredictionOptions) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:usingState:options:error:"), inputFeatures, state, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// Returns a model parameter value for a key.
//
// key: The key to a model parameter value.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/parameterValue(for:)
func (m MLModel) ParameterValueForKeyError(key IMLParameterKey) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parameterValueForKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// Construct a model asynchronously from a compiled model asset.
//
// asset: The compiled model asset derived from in-memory or on-disk Core ML model.
//
// configuration: The model configuration that holds options for loading the model.
//
// handler: The completion handler invoked when the load completes. A valid [MLModel]
// returns on success, or an error if failure.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/load(_:configuration:completionHandler:)
func (_MLModelClass MLModelClass) LoadModelAssetConfigurationCompletionHandler(asset IMLModelAsset, configuration IMLModelConfiguration, handler MLModelErrorHandler) {
_block2, _cleanup2 := NewMLModelErrorBlock(handler)
	defer _cleanup2()
	objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("loadModelAsset:configuration:completionHandler:"), asset, configuration, _block2)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/compileModel(at:)
func (_MLModelClass MLModelClass) CompileModelAtURLError(modelURL foundation.INSURL) (foundation.NSURL, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("compileModelAtURL:error:"), modelURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSURL{}, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSURLFromID(rv), nil

}

// The list of available compute devices that the model’s prediction methods
// use.
//
// See: https://developer.apple.com/documentation/coreml/mlmodel/availablecomputedevices-6klyt
func (m MLModel) AvailableComputeDevices() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("availableComputeDevices"))
	return objectivec.Object{ID: rv}
}
func (m MLModel) SetAvailableComputeDevices(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setAvailableComputeDevices:"), value)
}
// The configuration of the model set during initialization.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/configuration
func (m MLModel) Configuration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
// Model information you use at runtime during development, which Xcode also
// displays in its Core ML model editor view.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (m MLModel) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}
// A dictionary of the model’s creation information, such as its
// description, author, version, and license.
//
// See: https://developer.apple.com/documentation/coreml/mlmodeldescription/metadata
func (m MLModel) Metadata() MLModelMetadataKey {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("metadata"))
	return MLModelMetadataKey(foundation.NSStringFromID(rv).String())
}
func (m MLModel) SetMetadata(value MLModelMetadataKey) {
	objc.Send[struct{}](m.ID, objc.Sel("setMetadata:"), objc.String(string(value)))
}

// LoadModelAssetConfiguration is a synchronous wrapper around [MLModel.LoadModelAssetConfigurationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (mc MLModelClass) LoadModelAssetConfiguration(ctx context.Context, asset IMLModelAsset, configuration IMLModelConfiguration) (*MLModel, error) {
	type result struct {
		val *MLModel
		err error
	}
	done := make(chan result, 1)
	mc.LoadModelAssetConfigurationCompletionHandler(asset, configuration, func(val *MLModel, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

