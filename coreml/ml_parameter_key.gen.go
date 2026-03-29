// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLParameterKey] class.
var (
	_MLParameterKeyClass     MLParameterKeyClass
	_MLParameterKeyClassOnce sync.Once
)

func getMLParameterKeyClass() MLParameterKeyClass {
	_MLParameterKeyClassOnce.Do(func() {
		_MLParameterKeyClass = MLParameterKeyClass{class: objc.GetClass("MLParameterKey")}
	})
	return _MLParameterKeyClass
}

// GetMLParameterKeyClass returns the class object for MLParameterKey.
func GetMLParameterKeyClass() MLParameterKeyClass {
	return getMLParameterKeyClass()
}

type MLParameterKeyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLParameterKeyClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLParameterKeyClass) Alloc() MLParameterKey {
	rv := objc.Send[MLParameterKey](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The keys for the parameter dictionary in a model configuration or a model
// update context.
//
// # Overview
// 
// Use an [MLParameterKey] to retrieve a model’s parameter value using:
// 
// - The model’s [ParameterValueForKeyError] method - The [MLParameterKey.Parameters]
// dictionary of an [MLModelConfiguration] - The [MLParameterKey.Parameters] dictionary of an
// [MLUpdateContext]
// 
// # Overriding model and layer parameters
// 
// To override a model’s default parameter values:
// 
// - Create an [MLModelConfiguration] instance. - Use an [MLParameterKey] for
// each parameter to set its value in the model configuration’s [MLParameterKey.Parameters]
// dictionary. - Create a new model instance using
// [ModelWithContentsOfURLConfigurationError] with your custom model
// configuration.
// 
// # Configuring update parameters
// 
// To configure the update parameters for an [MLUpdateTask]:
// 
// - Create an [MLModelConfiguration] instance. - Use an [MLParameterKey] for
// each parameter to set its value in the model configuration’s [MLParameterKey.Parameters]
// dictionary. - Create a new update task with your custom model
// configuration.
// 
// See [Personalizing a Model with On-Device Updates].
//
// [Personalizing a Model with On-Device Updates]: https://developer.apple.com/documentation/CoreML/personalizing-a-model-with-on-device-updates
//
// # Scoping parameter keys
//
//   - [MLParameterKey.ScopedTo]: Creates a copy of a parameter key and adds the scope to it.
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey
type MLParameterKey struct {
	MLKey
}

// MLParameterKeyFromID constructs a [MLParameterKey] from an objc.ID.
//
// The keys for the parameter dictionary in a model configuration or a model
// update context.
func MLParameterKeyFromID(id objc.ID) MLParameterKey {
	return MLParameterKey{MLKey: MLKeyFromID(id)}
}
// NOTE: MLParameterKey adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLParameterKey] class.
//
// # Scoping parameter keys
//
//   - [IMLParameterKey.ScopedTo]: Creates a copy of a parameter key and adds the scope to it.
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey
type IMLParameterKey interface {
	IMLKey

	// Topic: Scoping parameter keys

	// Creates a copy of a parameter key and adds the scope to it.
	ScopedTo(scope string) IMLParameterKey

	// The list of available compute devices that the model’s prediction methods use.
	AvailableComputeDevices() objectivec.IObject
	SetAvailableComputeDevices(value objectivec.IObject)
	// The configuration of the model set during initialization.
	Configuration() IMLModelConfiguration
	SetConfiguration(value IMLModelConfiguration)
	// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
	ModelDescription() IMLModelDescription
	SetModelDescription(value IMLModelDescription)
	// A dictionary of configuration settings your app can override when loading a model.
	Parameters() IMLParameterKey
	SetParameters(value IMLParameterKey)
}

// Init initializes the instance.
func (p MLParameterKey) Init() MLParameterKey {
	rv := objc.Send[MLParameterKey](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLParameterKey) Autorelease() MLParameterKey {
	rv := objc.Send[MLParameterKey](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLParameterKey creates a new MLParameterKey instance.
func NewMLParameterKey() MLParameterKey {
	class := getMLParameterKeyClass()
	rv := objc.Send[MLParameterKey](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a copy of a parameter key and adds the scope to it.
//
// scope: A scope that targets the key to an internal component of a model.
//
// # Return Value
// 
// A new parameter key.
//
// # Discussion
// 
// Use this method to target:
// 
// - A specific layer of a model - A specific model within a pipeline model -
// A specific layer of a model within a pipeline model
// 
// For example, to target an [MLParameterKey] to a layer in a model, scope the
// key with the layer’s name.
// 
// To target a model within a pipeline model, scope the [MLParameterKey] with
// the model’s name.
// 
// To target a layer of a model within a pipeline model, start with an
// [MLParameterKey] you’ve already scoped to the model, and add an
// additional scope with the layer’s name.
// 
// By default, a pipeline model names its individual models in the pipeline as
// `model0`, `model1`, and so on.
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/scoped(to:)
func (p MLParameterKey) ScopedTo(scope string) IMLParameterKey {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("scopedTo:"), objc.String(scope))
	return MLParameterKeyFromID(rv)
}

// The list of available compute devices that the model’s prediction methods
// use.
//
// See: https://developer.apple.com/documentation/coreml/mlmodel/availablecomputedevices-6klyt
func (p MLParameterKey) AvailableComputeDevices() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("availableComputeDevices"))
	return objectivec.Object{ID: rv}
}
func (p MLParameterKey) SetAvailableComputeDevices(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setAvailableComputeDevices:"), value)
}
// The configuration of the model set during initialization.
//
// See: https://developer.apple.com/documentation/coreml/mlmodel/configuration
func (p MLParameterKey) Configuration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
func (p MLParameterKey) SetConfiguration(value IMLModelConfiguration) {
	objc.Send[struct{}](p.ID, objc.Sel("setConfiguration:"), value)
}
// Model information you use at runtime during development, which Xcode also
// displays in its Core ML model editor view.
//
// See: https://developer.apple.com/documentation/coreml/mlmodel/modeldescription
func (p MLParameterKey) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}
func (p MLParameterKey) SetModelDescription(value IMLModelDescription) {
	objc.Send[struct{}](p.ID, objc.Sel("setModelDescription:"), value)
}
// A dictionary of configuration settings your app can override when loading a
// model.
//
// See: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/parameters
func (p MLParameterKey) Parameters() IMLParameterKey {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("parameters"))
	return MLParameterKeyFromID(objc.ID(rv))
}
func (p MLParameterKey) SetParameters(value IMLParameterKey) {
	objc.Send[struct{}](p.ID, objc.Sel("setParameters:"), value)
}

// The key you use to access the number of neighbors that adjusts the affinity
// of a k-nearest-neighbor model.
//
// # Discussion
// 
// The value type for the [NumberOfNeighbors] key is an [Int64].
//
// [Int64]: https://developer.apple.com/documentation/Swift/Int64
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/numberOfNeighbors
func (_MLParameterKeyClass MLParameterKeyClass) NumberOfNeighbors() MLParameterKey {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("numberOfNeighbors"))
	return MLParameterKeyFromID(objc.ID(rv))
}
// The key you use to access the linked model’s filename.
//
// # Discussion
// 
// The value type for the [LinkedModelFileName] key is a [String].
//
// [String]: https://developer.apple.com/documentation/Swift/String
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/linkedModelFileName
func (_MLParameterKeyClass MLParameterKeyClass) LinkedModelFileName() MLParameterKey {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("linkedModelFileName"))
	return MLParameterKeyFromID(objc.ID(rv))
}
// The key you use to access the linked model’s search path.
//
// # Discussion
// 
// The value type for the [LinkedModelSearchPath] key is a [String].
//
// [String]: https://developer.apple.com/documentation/Swift/String
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/linkedModelSearchPath
func (_MLParameterKeyClass MLParameterKeyClass) LinkedModelSearchPath() MLParameterKey {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("linkedModelSearchPath"))
	return MLParameterKeyFromID(objc.ID(rv))
}
// The key you use to access the weights of a layer in a neural network model.
//
// # Discussion
// 
// The value type for the [Weights] key is an [MLMultiArray]. You must scope
// this key with the name of the specific neural network layer whose weights
// you’d like to access. See [ScopedTo].
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/weights
func (_MLParameterKeyClass MLParameterKeyClass) Weights() MLParameterKey {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("weights"))
	return MLParameterKeyFromID(objc.ID(rv))
}
// The key you use to access the biases of a layer in a neural network model.
//
// # Discussion
// 
// The value type for the [Biases] key is an [MLMultiArray]. You must scope
// this key with the name of the specific neural network layer whose biases
// you’d like to access. See [ScopedTo].
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/biases
func (_MLParameterKeyClass MLParameterKeyClass) Biases() MLParameterKey {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("biases"))
	return MLParameterKeyFromID(objc.ID(rv))
}
// The key you use to access the optimizer’s learning rate parameter.
//
// # Discussion
// 
// The value type for the [LearningRate] key is a [Double].
// 
// To modify a model’s learning rate midway through an [MLUpdateTask], use
// its [ResumeWithParameters] method to set a new value for the model’s
// learning rate. You do this in the progress handler that you specified in
// the [MLUpdateProgressHandlers] instance when you created the update task
// using
// [UpdateTaskForModelAtURLTrainingDataConfigurationProgressHandlersError].
// 
// See [Personalizing a Model with On-Device Updates].
//
// [Double]: https://developer.apple.com/documentation/Swift/Double
// [Personalizing a Model with On-Device Updates]: https://developer.apple.com/documentation/CoreML/personalizing-a-model-with-on-device-updates
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/learningRate
func (_MLParameterKeyClass MLParameterKeyClass) LearningRate() MLParameterKey {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("learningRate"))
	return MLParameterKeyFromID(objc.ID(rv))
}
// The key you use to access the stochastic gradient descent (SGD)
// optimizer’s momentum parameter.
//
// # Discussion
// 
// The value type for the [Momentum] key is a [Double].
//
// [Double]: https://developer.apple.com/documentation/Swift/Double
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/momentum
func (_MLParameterKeyClass MLParameterKeyClass) Momentum() MLParameterKey {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("momentum"))
	return MLParameterKeyFromID(objc.ID(rv))
}
// The key you use to access the optimizer’s mini batch-size parameter.
//
// # Discussion
// 
// The value type for the [MiniBatchSize] key is an [Int64].
//
// [Int64]: https://developer.apple.com/documentation/Swift/Int64
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/miniBatchSize
func (_MLParameterKeyClass MLParameterKeyClass) MiniBatchSize() MLParameterKey {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("miniBatchSize"))
	return MLParameterKeyFromID(objc.ID(rv))
}
// The key you use to access the Adam optimizer’s first beta parameter.
//
// # Discussion
// 
// The value type for the [Beta1] key is a [Double].
//
// [Double]: https://developer.apple.com/documentation/Swift/Double
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/beta1
func (_MLParameterKeyClass MLParameterKeyClass) Beta1() MLParameterKey {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("beta1"))
	return MLParameterKeyFromID(objc.ID(rv))
}
// The key you use to access the Adam optimizer’s second beta parameter.
//
// # Discussion
// 
// The value type for the [Beta2] key is a [Double].
//
// [Double]: https://developer.apple.com/documentation/Swift/Double
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/beta2
func (_MLParameterKeyClass MLParameterKeyClass) Beta2() MLParameterKey {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("beta2"))
	return MLParameterKeyFromID(objc.ID(rv))
}
// The key you use to access the Adam optimizer’s epsilon parameter.
//
// # Discussion
// 
// The value type for the [Eps] key is a [Double].
//
// [Double]: https://developer.apple.com/documentation/Swift/Double
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/eps
func (_MLParameterKeyClass MLParameterKeyClass) Eps() MLParameterKey {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("eps"))
	return MLParameterKeyFromID(objc.ID(rv))
}
// The key you use to access the optimizer’s epochs parameter.
//
// # Discussion
// 
// The value type for the [Epochs] key is an [Int64].
//
// [Int64]: https://developer.apple.com/documentation/Swift/Int64
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/epochs
func (_MLParameterKeyClass MLParameterKeyClass) Epochs() MLParameterKey {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("epochs"))
	return MLParameterKeyFromID(objc.ID(rv))
}
// The key you use to access the shuffle parameter, a Boolean value that
// determines whether the model randomizes the data between epochs.
//
// # Discussion
// 
// The value type for the [Shuffle] key is an [Bool].
//
// [Bool]: https://developer.apple.com/documentation/Swift/Bool
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/shuffle
func (_MLParameterKeyClass MLParameterKeyClass) Shuffle() MLParameterKey {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("shuffle"))
	return MLParameterKeyFromID(objc.ID(rv))
}
// The key you use to access the seed parameter that initializes the random
// number generator for the shuffle option.
//
// # Discussion
// 
// The value type for the [Seed] key is an [Int64].
//
// [Int64]: https://developer.apple.com/documentation/Swift/Int64
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterKey/seed
func (_MLParameterKeyClass MLParameterKeyClass) Seed() MLParameterKey {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterKeyClass.class), objc.Sel("seed"))
	return MLParameterKeyFromID(objc.ID(rv))
}

