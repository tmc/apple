// Code generated from Apple documentation for CoreML. DO NOT EDIT.

// Package coreml provides Go bindings for the CoreML framework.
//
// Integrate machine learning models into your app.
//
// Use [Core ML](<doc://com.apple.coreml/documentation/CoreML>) to integrate machine learning models into your app. [Core ML](<doc://com.apple.coreml/documentation/CoreML>) provides a unified representation for all models. Your app uses [Core ML](<doc://com.apple.coreml/documentation/CoreML>) APIs and user data to make predictions, and to train or fine-tune models, all on a person’s device.
//
// # Core ML models
//
//   - Getting a Core ML Model: Obtain a Core ML model to use in your app.
//   - Updating a Model File to a Model Package: Convert a Core ML model file into a model package in Xcode.
//   - Integrating a Core ML Model into Your App: Add a simple model to an app, pass input data to the model, and process the model’s predictions.
//   - MLModel: An encapsulation of all the details of your machine learning model. ([MLPredictionOptions], [MLModelDescription], [MLParameterKey], [MLModelConfiguration], [MLOptimizationHints])
//   - Model Customization: Expand and modify your model with new layers. ([MLCustomLayer], [MLCustomModel])
//   - Model Personalization: Update your model to adapt to new data. ([MLTask], [MLUpdateTask])
//
// # Model inputs and outputs
//
//   - Making Predictions with a Sequence of Inputs: Integrate a recurrent neural network model to process sequences of inputs.
//   - MLFeatureValue: A generic wrapper around an underlying value and the value’s type. ([MLImageConstraint], [MLFeatureType], [MLMultiArray], [MLSequence])
//   - MLSendableFeatureValue: A sendable feature value.
//   - MLFeatureProvider: An interface that represents a collection of values for either a model’s input or its output.
//   - MLDictionaryFeatureProvider: A convenience wrapper for the given dictionary of data.
//   - MLBatchProvider: An interface that represents a collection of feature providers.
//   - MLArrayBatchProvider: A convenience wrapper for batches of feature providers.
//   - MLModelAsset: An abstraction of a compiled Core ML model asset.
//
// # App integration
//
//   - Downloading and Compiling a Model on the User’s Device: Install Core ML models on the user’s device dynamically at runtime.
//   - Model Integration Samples: Integrate tabular, image, and text classifcation models into your app.
//
// # Model encryption
//
//   - Generating a Model Encryption Key: Create a model encryption key to encrypt a compiled model or model archive.
//   - Encrypting a Model in Your App: Encrypt your app’s built-in model at compile time by adding a compiler flag.
//
// # Compute devices
//
//   - MLComputeDevice: Compute devices for framework operations.
//   - MLCPUComputeDevice: An object that represents a CPU compute device.
//   - MLGPUComputeDevice: An object that represents a GPU compute device.
//   - MLNeuralEngineComputeDevice: An object that represents a Neural Engine compute device.
//   - MLComputeDeviceProtocol: An interface that represents a compute device type.
//
// # Compute plan
//
//   - MLComputePlan: A class representing the compute plan of a model.
//   - MLModelStructure: An enum representing the structure of a model.
//   - MLComputePolicy: The compute policy determining what compute device, or compute devices, to execute ML workloads on.
//   - withMLTensorComputePolicy(_:_:): Calls the given closure within a task-local context using the specified compute policy to influence what compute device tensor operations are executed on.
//   - withMLTensorComputePolicy(_:_:): Calls the given closure within a task-local context using the specified compute policy to influence what compute device tensor operations are executed on.
//
// # Model state
//
//   - MLState: Handle to the state buffers.
//   - MLStateConstraint: Constraint of a state feature value.
//
// # Model tensor
//
//   - MLTensor: A multi-dimensional array of numerical or Boolean scalars tailored to ML use cases, containing methods to perform transformations and mathematical operations efficiently using a ML compute device.
//   - MLTensorScalar: A type that represents the tensor scalar types supported by the framework.
//   - MLTensorRangeExpression: A type that can be used to slice a dimension of a tensor.
//   - pointwiseMin(_:_:): Computes the element-wise minimum of two tensors.
//   - pointwiseMax(_:_:): Computes the element-wise minimum between two tensors.
//   - withMLTensorComputePolicy(_:_:): Calls the given closure within a task-local context using the specified compute policy to influence what compute device tensor operations are executed on.
//
// # Model structure
//
//   - MLModelStructure: An enum representing the structure of a model.
//
// # Model errors
//
//   - MLModelError: Information about a Core ML model error.
//   - MLModelError.Code: Information about a Core ML model error.
//   - MLModelErrorDomain: The domain for Core ML errors.
//
// # Key Types
//
//   - [MLFeatureValue] - A generic wrapper around an underlying value and the value’s type.
//   - [MLMultiArray] - A machine learning collection type that stores numeric values in an array with multiple dimensions.
//   - [MLModel] - An encapsulation of all the details of your machine learning model.
//   - [MLParameterKey] - The keys for the parameter dictionary in a model configuration or a model update context.
//   - [MLModelDescription] - Information about a model, primarily the input and output format for each feature the model expects, and optional metadata.
//   - [MLFeatureDescription] - The name, type, and constraints of an input or output feature.
//   - [MLMultiArrayConstraint] - The shape and data type constraints for a multidimensional array feature.
//   - [MLImageSizeConstraint] - A list or range of sizes that augment an image constraint’s default size.
//   - [MLModelConfiguration] - The settings for creating or updating a machine learning model.
//   - [MLSequenceConstraint] - The constraints for a sequence feature.
//
// [CoreML Documentation]: https://developer.apple.com/documentation/CoreML
package coreml

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the CoreML library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/CoreML.framework/CoreML",
	"/usr/lib/libCoreML.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	fmt.Fprintf(os.Stderr, "warning: CoreML: failed to load framework from any known path\n")
}
