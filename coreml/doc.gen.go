// Code generated from Apple documentation for CoreML. DO NOT EDIT.

// Package coreml provides Go bindings for the CoreML framework.
//
// Integrate machine learning models into your app.
//
// Use Core ML to integrate machine learning models into your app. Core ML
// provides a unified representation for all models. Your app uses Core ML
// APIs and user data to make predictions, and to train or fine-tune models,
// all on a person’s device.
//
// # Core ML models
//
//   - [Getting a Core ML Model]: Obtain a Core ML model to use in your app.
//   - [Updating a Model File to a Model Package]: Convert a Core ML model file into a model package in Xcode.
//   - [Integrating a Core ML Model into Your App]: Add a simple model to an app, pass input data to the model, and process the model’s predictions.
//   - [MLModel]: An encapsulation of all the details of your machine learning model. ([MLPredictionOptions], [MLModelDescription], [MLParameterKey], [MLModelConfiguration], [MLOptimizationHints])
//   - [Model Customization]: Expand and modify your model with new layers. ([MLCustomLayer], [MLCustomModel])
//   - [Model Personalization]: Update your model to adapt to new data. ([MLTask], [MLUpdateTask])
//
// # Model inputs and outputs
//
//   - [Making Predictions with a Sequence of Inputs]: Integrate a recurrent neural network model to process sequences of inputs.
//   - [MLFeatureValue]: A generic wrapper around an underlying value and the value’s type. ([MLImageConstraint], [MLFeatureType], [MLMultiArray], [MLSequence])
//   - [MLFeatureProvider]: An interface that represents a collection of values for either a model’s input or its output.
//   - [MLDictionaryFeatureProvider]: A convenience wrapper for the given dictionary of data.
//   - [MLBatchProvider]: An interface that represents a collection of feature providers.
//   - [MLArrayBatchProvider]: A convenience wrapper for batches of feature providers.
//   - [MLModelAsset]: An abstraction of a compiled Core ML model asset.
//
// # App integration
//
//   - [Downloading and Compiling a Model on the User’s Device]: Install Core ML models on the user’s device dynamically at runtime.
//   - [Model Integration Samples]: Integrate tabular, image, and text classifcation models into your app.
//
// # Model encryption
//
//   - [Generating a Model Encryption Key]: Create a model encryption key to encrypt a compiled model or model archive.
//   - [Encrypting a Model in Your App]: Encrypt your app’s built-in model at compile time by adding a compiler flag.
//
// # Compute devices
//
//   - [MLCPUComputeDevice]: An object that represents a CPU compute device.
//   - [MLGPUComputeDevice]: An object that represents a GPU compute device.
//   - [MLNeuralEngineComputeDevice]: An object that represents a Neural Engine compute device.
//   - [MLComputeDeviceProtocol]: An interface that represents a compute device type.
//   - [MLAllComputeDevices]: Returns an array that contains all of the compute devices that are accessible.
//
// # Compute plan
//
//   - [MLComputePlan]: A class describing the plan for executing a model. ([MLComputePlanDeviceUsage], [MLComputePlanCost])
//   - [MLComputePlanCost]: A class that represents the estimated cost of executing a layer or operation.
//   - [MLComputePlanDeviceUsage]: The anticipated compute devices to use for executing a layer or operation.
//
// # Model state
//
//   - [MLState]: Handle to the state buffers.
//   - [MLStateConstraint]: Constraint of a state feature value.
//
// # Model structure
//
//   - [MLModelStructure]: A class representing the structure of a model.
//   - [MLModelStructureNeuralNetwork]: A class representing the structure of a NeuralNetwork model.
//   - [MLModelStructureNeuralNetworkLayer]: A class representing a layer in a NeuralNetwork.
//   - [MLModelStructurePipeline]: A class representing the structure of a Pipeline model.
//   - [MLModelStructureProgram]: A class representing the structure of an ML Program model.
//   - [MLModelStructureProgramArgument]: A class representing an argument in the Program.
//   - [MLModelStructureProgramBinding]: A class representing a binding in the Program
//   - [MLModelStructureProgramBlock]: A class representing a block in the Program.
//   - [MLModelStructureProgramFunction]: A class representing a function in the Program.
//   - [MLModelStructureProgramNamedValueType]: A class representing a named value type in a Program.
//   - [MLModelStructureProgramOperation]: A class representing an Operation in a Program.
//   - [MLModelStructureProgramValue]: A class representing a constant value in the Program.
//   - [MLModelStructureProgramValueType]: A class representing the type of a value or a variable in the Program.
//
// # Model errors
//
//   - [MLModelError]: Information about a Core ML model error.
//   - [MLModelErrorDomain]: The domain for Core ML errors.
//
// # Optimization
//
//   - [MLOptimizationHints]: MLOptimizationHints//
//
// # Key Types
//
//   - [MLFeatureValue] - A generic wrapper around an underlying value and the value’s type.
//   - [MLModel] - An encapsulation of all the details of your machine learning model.
//   - [MLMultiArray] - A machine learning collection type that stores numeric values in an array with multiple dimensions.
//   - [MLParameterKey] - The keys for the parameter dictionary in a model configuration or a model update context.
//   - [MLModelDescription] - Information about a model, primarily the input and output format for each feature the model expects, and optional metadata.
//   - [MLFeatureDescription] - The name, type, and constraints of an input or output feature.
//   - [MLModelConfiguration] - The settings for creating or updating a machine learning model.
//   - [MLDictionaryFeatureProvider] - A convenience wrapper for the given dictionary of data.
//   - [MLSequence] - A machine learning collection type that stores a series of strings or integers.
//   - [MLArrayBatchProvider] - A convenience wrapper for batches of feature providers.
//
// [Downloading and Compiling a Model on the User’s Device]: https://developer.apple.com/documentation/coreml/downloading-and-compiling-a-model-on-the-user-s-device
// [Encrypting a Model in Your App]: https://developer.apple.com/documentation/coreml/encrypting-a-model-in-your-app
// [Generating a Model Encryption Key]: https://developer.apple.com/documentation/coreml/generating-a-model-encryption-key
// [Getting a Core ML Model]: https://developer.apple.com/documentation/coreml/getting-a-core-ml-model
// [Integrating a Core ML Model into Your App]: https://developer.apple.com/documentation/coreml/integrating-a-core-ml-model-into-your-app
// [Making Predictions with a Sequence of Inputs]: https://developer.apple.com/documentation/coreml/making-predictions-with-a-sequence-of-inputs
// [Model Customization]: https://developer.apple.com/documentation/coreml/model-customization
// [Model Integration Samples]: https://developer.apple.com/documentation/coreml/model-integration-samples
// [Model Personalization]: https://developer.apple.com/documentation/coreml/model-personalization
// [Updating a Model File to a Model Package]: https://developer.apple.com/documentation/coreml/updating-a-model-file-to-a-model-package
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
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: CoreML: failed to load framework from any known path\n")
	}
}
