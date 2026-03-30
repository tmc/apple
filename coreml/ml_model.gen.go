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

// Class returns the underlying Objective-C class pointer.
func (mc MLModelClass) Class() objc.Class {
	return mc.class
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
// [MLModel.PredictionFromFeaturesError] or [MLModel.PredictionFromFeaturesOptionsError]. -
// Make multiple predictions with your app’s custom [MLBatchProvider] by
// calling [MLModel.PredictionsFromBatchError] or [MLModel.PredictionsFromBatchOptionsError].
// - Inspect your model’s [MLModel.Metadata] and [MLFeatureDescription] instances
// through [ModelDescription].
//
// If your app downloads and compiles a model on the user’s device, you must
// use the [MLModel] class directly to make predictions. See [Downloading and
// Compiling a Model on the User’s Device].
//
// # Making predictions
//
//   - [MLModel.PredictionsFromBatchError]: Generates predictions for each input feature provider within the batch provider.
//   - [MLModel.PredictionsFromBatchOptionsError]: Generates a prediction for each input feature provider within the batch provider using the prediction options.
//
// # Inspecting a model
//
//   - [MLModel.Configuration]: The configuration of the model set during initialization.
//   - [MLModel.ModelDescription]: Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//   - [MLModel.ParameterValueForKeyError]: Returns a model parameter value for a key.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel
//
// [Downloading and Compiling a Model on the User’s Device]: https://developer.apple.com/documentation/CoreML/downloading-and-compiling-a-model-on-the-user-s-device
// [Integrating a Core ML Model into Your App]: https://developer.apple.com/documentation/CoreML/integrating-a-core-ml-model-into-your-app
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
//   - [IMLModel.PredictionsFromBatchError]: Generates predictions for each input feature provider within the batch provider.
//   - [IMLModel.PredictionsFromBatchOptionsError]: Generates a prediction for each input feature provider within the batch provider using the prediction options.
//
// # Inspecting a model
//
//   - [IMLModel.Configuration]: The configuration of the model set during initialization.
//   - [IMLModel.ModelDescription]: Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//   - [IMLModel.ParameterValueForKeyError]: Returns a model parameter value for a key.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel
type IMLModel interface {
	objectivec.IObject

	// Topic: Making predictions

	// Generates predictions for each input feature provider within the batch provider.
	PredictionsFromBatchError(inputBatch MLBatchProvider) (MLBatchProvider, error)
	// Generates a prediction for each input feature provider within the batch provider using the prediction options.
	PredictionsFromBatchOptionsError(inputBatch MLBatchProvider, options IMLPredictionOptions) (MLBatchProvider, error)

	// Topic: Inspecting a model

	// The configuration of the model set during initialization.
	Configuration() IMLModelConfiguration
	// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
	ModelDescription() IMLModelDescription
	// Returns a model parameter value for a key.
	ParameterValueForKeyError(key IMLParameterKey) (objectivec.IObject, error)

	// A dictionary of the model’s creation information, such as its description, author, version, and license.
	Metadata() MLModelMetadataKey
	SetMetadata(value MLModelMetadataKey)
	// Creates a new state object.
	NewState() IMLState
	// Generates a prediction asynchronously from the feature values within the input feature provider.
	PredictionFromFeaturesCompletionHandler(input MLFeatureProvider, completionHandler MLFeatureProviderErrorHandler)
	// Generates a prediction from the feature values within the input feature provider.
	PredictionFromFeaturesError(input MLFeatureProvider) (MLFeatureProvider, error)
	// Generates a prediction asynchronously from the feature values within the input feature provider using the prediction options.
	PredictionFromFeaturesOptionsCompletionHandler(input MLFeatureProvider, options IMLPredictionOptions, completionHandler MLFeatureProviderErrorHandler)
	// Generates a prediction from the feature values within the input feature provider using the prediction options.
	PredictionFromFeaturesOptionsError(input MLFeatureProvider, options IMLPredictionOptions) (MLFeatureProvider, error)
	// Run a stateful prediction synchronously.
	PredictionFromFeaturesUsingStateError(inputFeatures MLFeatureProvider, state IMLState) (MLFeatureProvider, error)
	// Run a stateful prediction asynchronously.
	PredictionFromFeaturesUsingStateOptionsCompletionHandler(inputFeatures MLFeatureProvider, state IMLState, options IMLPredictionOptions, completionHandler MLFeatureProviderErrorHandler)
	// Run a stateful prediction synchronously with options.
	PredictionFromFeaturesUsingStateOptionsError(inputFeatures MLFeatureProvider, state IMLState, options IMLPredictionOptions) (MLFeatureProvider, error)
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
// See: https://developer.apple.com/documentation/CoreML/MLModel/init(contentsOf:configuration:)
//
// [compileModel(at:)]: https://developer.apple.com/documentation/CoreML/MLModel/compileModel(at:)-6442s
// [Downloading and Compiling a Model on the User’s Device]: https://developer.apple.com/documentation/CoreML/downloading-and-compiling-a-model-on-the-user-s-device
// [Integrating a Core ML Model into Your App]: https://developer.apple.com/documentation/CoreML/integrating-a-core-ml-model-into-your-app
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
// See: https://developer.apple.com/documentation/CoreML/MLModel/init(contentsOf:)
//
// [compileModel(at:)]: https://developer.apple.com/documentation/CoreML/MLModel/compileModel(at:)-6442s
// [Downloading and Compiling a Model on the User’s Device]: https://developer.apple.com/documentation/CoreML/downloading-and-compiling-a-model-on-the-user-s-device
// [Integrating a Core ML Model into Your App]: https://developer.apple.com/documentation/CoreML/integrating-a-core-ml-model-into-your-app
func NewModelWithContentsOfURLError(url foundation.INSURL) (MLModel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getMLModelClass().class), objc.Sel("modelWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelFromID(rv), nil
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

// Creates a new state object.
//
// # Discussion
//
// Core ML framework will allocate the state buffers declared in the model.
//
// The allocated state buffers are initialized to zeros. To initialize with
// different values, use `XCUIElementTypeWithMultiArray()` to get the mutable
// [MLMultiArray]-view to the state buffer.
//
// It returns an empty state when the model is stateless. One can use the
// empty state with stateful prediction functions such as `prediction()` and
// those predictions will be stateless. This simplifies the call site which
// may or may not use a stateful model.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/newState
func (m MLModel) NewState() IMLState {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("newState"))
	return MLStateFromID(rv)
}

// Generates a prediction asynchronously from the feature values within the
// input feature provider.
//
// input: A feature provider that stores all the input feature values the model needs
// for a prediction.
//
// completionHandler: The callback the system invokes when it completes the prediction.
//
// output: A feature provider that contains the outputs of the prediction.
// error: If an error occurs, an error object that describes the error;
// otherwise, `nil`.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/predictionFromFeatures:completionHandler:
func (m MLModel) PredictionFromFeaturesCompletionHandler(input MLFeatureProvider, completionHandler MLFeatureProviderErrorHandler) {
	_block1, _ := NewMLFeatureProviderErrorBlock(completionHandler)
	objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:completionHandler:"), input, _block1)
}

// Generates a prediction from the feature values within the input feature
// provider.
//
// input: A feature provider that stores all the input feature values the model needs
// for a prediction.
//
// # Return Value
//
// A feature provider that contains the outputs of the prediction.
//
// # Discussion
//
// Use this method to make a single prediction.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/prediction(from:)-9y2aa
func (m MLModel) PredictionFromFeaturesError(input MLFeatureProvider) (MLFeatureProvider, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:error:"), input, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return MLFeatureProviderObjectFromID(rv), nil

}

// Generates a prediction asynchronously from the feature values within the
// input feature provider using the prediction options.
//
// input: A feature provider that stores all the input feature values the model needs
// for a prediction.
//
// options: The runtime settings the model uses as it makes a prediction.
//
// completionHandler: The callback the system invokes when it completes the prediction.
//
// output: A feature provider that contains the outputs of the prediction.
// error: If an error occurs, an error object that describes the error;
// otherwise, `nil`.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/predictionFromFeatures:options:completionHandler:
func (m MLModel) PredictionFromFeaturesOptionsCompletionHandler(input MLFeatureProvider, options IMLPredictionOptions, completionHandler MLFeatureProviderErrorHandler) {
	_block2, _ := NewMLFeatureProviderErrorBlock(completionHandler)
	objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:options:completionHandler:"), input, options, _block2)
}

// Generates a prediction from the feature values within the input feature
// provider using the prediction options.
//
// input: A feature provider that stores all the input feature values the model needs
// for a prediction.
//
// options: The runtime settings the model uses as it makes a prediction.
//
// # Return Value
//
// A feature provider that contains the outputs of the prediction.
//
// # Discussion
//
// Use this method to make a single prediction.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/prediction(from:options:)-81mr6
func (m MLModel) PredictionFromFeaturesOptionsError(input MLFeatureProvider, options IMLPredictionOptions) (MLFeatureProvider, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:options:error:"), input, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return MLFeatureProviderObjectFromID(rv), nil

}

// Run a stateful prediction synchronously.
//
// # Discussion
//
// Use this method to run predictions on a stateful model.
//
// - inputFeatures: The input features as declared in the model description. -
// state: The state object created by `newState()` method. - error: The output
// parameter to receive an error information on failure.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/prediction(from:using:)-97bu1
func (m MLModel) PredictionFromFeaturesUsingStateError(inputFeatures MLFeatureProvider, state IMLState) (MLFeatureProvider, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:usingState:error:"), inputFeatures, state, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return MLFeatureProviderObjectFromID(rv), nil

}

// Run a stateful prediction asynchronously.
//
// # Discussion
//
// Use this method to run predictions on a stateful model.
//
// Do not request a prediction while another prediction that shares the same
// state is in-flight, otherwise the behavior is undefined.
//
// - Parameters - input: The input features to make a prediction from. -
// state: The state object created by `newState()` method. - options:
// Prediction options to modify how the prediction is run. -
// completionHandler: A block that will be invoked once the prediction has
// completed successfully or unsuccessfully. On success, it is invoked with a
// valid model output. On failure, it is invoked with a nil output and NSError
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/predictionFromFeatures:usingState:options:completionHandler:
func (m MLModel) PredictionFromFeaturesUsingStateOptionsCompletionHandler(inputFeatures MLFeatureProvider, state IMLState, options IMLPredictionOptions, completionHandler MLFeatureProviderErrorHandler) {
	_block3, _ := NewMLFeatureProviderErrorBlock(completionHandler)
	objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:usingState:options:completionHandler:"), inputFeatures, state, options, _block3)
}

// Run a stateful prediction synchronously with options.
//
// # Discussion
//
// Use this method to run predictions on a stateful model.
//
// - inputFeatures: The input features as declared in the model description. -
// state: The state object created by `newState()` method. - options: The
// prediction options. - error: The output parameter to receive an error
// information on failure.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/prediction(from:using:options:)-v4wp
func (m MLModel) PredictionFromFeaturesUsingStateOptionsError(inputFeatures MLFeatureProvider, state IMLState, options IMLPredictionOptions) (MLFeatureProvider, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:usingState:options:error:"), inputFeatures, state, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return MLFeatureProviderObjectFromID(rv), nil

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
	_block2, _ := NewMLModelErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("loadModelAsset:configuration:completionHandler:"), asset, configuration, _block2)
}

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

// Compile a model for a device.
//
// modelURL: The URL to the model file.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/compileModel(at:)-3nea
func (_MLModelClass MLModelClass) CompileModelAtURLCompletionHandler(modelURL foundation.INSURL, handler URLErrorHandler) {
	_block1, _ := NewURLErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("compileModelAtURL:completionHandler:"), modelURL, _block1)
}

// Creates a Core ML model instance asynchronously from a compiled model file,
// a custom configuration, and a completion handler.
//
// url: The path to a compiled model file (`XCUIElementTypeMlmodelc`), typically
// with the [URL] that [compileModel(at:)] returns.
//
// configuration: The runtime settings for the new model instance.
//
// handler: A closure the method calls when it finishes loading the model.
//
// # Discussion
//
// Use this method to load a model asynchronously. Core ML calls your
// completion handler after it successfully loads the model, or encounters an
// error attempting to load it.
//
// In Swift, if the model loaded successfully, you can use the instance from
// the [Result.success(_:)] associated value; otherwise, use the
// [Result.failure(_:)] associated value to address the error. In Objective-C,
// you can use the [MLModel] instance in your completion hander; otherwise,
// use the [NSError] instance to address the error. See [MLModelError.Code]
// for the list of error codes.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/loadContentsOfURL:configuration:completionHandler:
//
// [compileModel(at:)]: https://developer.apple.com/documentation/CoreML/MLModel/compileModel(at:)-6442s
// [MLModelError.Code]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
// [Result.failure(_:)]: https://developer.apple.com/documentation/Swift/Result/failure(_:)
// [Result.success(_:)]: https://developer.apple.com/documentation/Swift/Result/success(_:)
func (_MLModelClass MLModelClass) LoadContentsOfURLConfigurationCompletionHandler(url foundation.INSURL, configuration IMLModelConfiguration, handler MLModelErrorHandler) {
	_block2, _ := NewMLModelErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("loadContentsOfURL:configuration:completionHandler:"), url, configuration, _block2)
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

// The list of available compute devices that the model’s prediction can
// use.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/availableComputeDevices-42uzt
func (_MLModelClass MLModelClass) AvailableComputeDevices() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](objc.ID(_MLModelClass.class), objc.Sel("availableComputeDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
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

// CompileModelAtURL is a synchronous wrapper around [MLModel.CompileModelAtURLCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (mc MLModelClass) CompileModelAtURL(ctx context.Context, modelURL foundation.INSURL) (*foundation.NSURL, error) {
	type result struct {
		val *foundation.NSURL
		err error
	}
	done := make(chan result, 1)
	mc.CompileModelAtURLCompletionHandler(modelURL, func(val *foundation.NSURL, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// LoadContentsOfURLConfiguration is a synchronous wrapper around [MLModel.LoadContentsOfURLConfigurationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (mc MLModelClass) LoadContentsOfURLConfiguration(ctx context.Context, url foundation.INSURL, configuration IMLModelConfiguration) (*MLModel, error) {
	type result struct {
		val *MLModel
		err error
	}
	done := make(chan result, 1)
	mc.LoadContentsOfURLConfigurationCompletionHandler(url, configuration, func(val *MLModel, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// PredictionFromFeatures is a synchronous wrapper around [MLModel.PredictionFromFeaturesCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModel) PredictionFromFeatures(ctx context.Context, input MLFeatureProvider) (MLFeatureProvider, error) {
	type result struct {
		val MLFeatureProvider
		err error
	}
	done := make(chan result, 1)
	m.PredictionFromFeaturesCompletionHandler(input, func(val MLFeatureProvider, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// PredictionFromFeaturesOptions is a synchronous wrapper around [MLModel.PredictionFromFeaturesOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModel) PredictionFromFeaturesOptions(ctx context.Context, input MLFeatureProvider, options IMLPredictionOptions) (MLFeatureProvider, error) {
	type result struct {
		val MLFeatureProvider
		err error
	}
	done := make(chan result, 1)
	m.PredictionFromFeaturesOptionsCompletionHandler(input, options, func(val MLFeatureProvider, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// PredictionFromFeaturesUsingStateOptions is a synchronous wrapper around [MLModel.PredictionFromFeaturesUsingStateOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModel) PredictionFromFeaturesUsingStateOptions(ctx context.Context, inputFeatures MLFeatureProvider, state IMLState, options IMLPredictionOptions) (MLFeatureProvider, error) {
	type result struct {
		val MLFeatureProvider
		err error
	}
	done := make(chan result, 1)
	m.PredictionFromFeaturesUsingStateOptionsCompletionHandler(inputFeatures, state, options, func(val MLFeatureProvider, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
