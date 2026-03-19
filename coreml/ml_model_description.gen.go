// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelDescription] class.
var (
	_MLModelDescriptionClass     MLModelDescriptionClass
	_MLModelDescriptionClassOnce sync.Once
)

func getMLModelDescriptionClass() MLModelDescriptionClass {
	_MLModelDescriptionClassOnce.Do(func() {
		_MLModelDescriptionClass = MLModelDescriptionClass{class: objc.GetClass("MLModelDescription")}
	})
	return _MLModelDescriptionClass
}

// GetMLModelDescriptionClass returns the class object for MLModelDescription.
func GetMLModelDescriptionClass() MLModelDescriptionClass {
	return getMLModelDescriptionClass()
}

type MLModelDescriptionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelDescriptionClass) Alloc() MLModelDescription {
	rv := objc.Send[MLModelDescription](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Information about a model, primarily the input and output format for each
// feature the model expects, and optional metadata.
//
// # Accessing feature descriptions
//
//   - [MLModelDescription.StateDescriptionsByName]: Description of the state features.
//   - [MLModelDescription.InputDescriptionsByName]: A dictionary of input feature descriptions, which the model keys by the input’s name.
//   - [MLModelDescription.OutputDescriptionsByName]: A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// # Accessing metadata
//
//   - [MLModelDescription.ClassLabels]: An array of labels, which can be either strings or a numbers, for classifier models.
//   - [MLModelDescription.Metadata]: A dictionary of the model’s creation information, such as its description, author, version, and license.
//
// # Accessing prediction names
//
//   - [MLModelDescription.PredictedFeatureName]: The name of the primary prediction feature output description.
//   - [MLModelDescription.PredictedProbabilitiesName]: The name of the feature output description for all probabilities of a prediction.
//
// # Accessing update descriptions
//
//   - [MLModelDescription.IsUpdatable]: A Boolean value that indicates whether you can update the model with additional training.
//   - [MLModelDescription.TrainingInputDescriptionsByName]: A dictionary of the training input feature descriptions, which the model keys by the input’s name.
//   - [MLModelDescription.ParameterDescriptionsByKey]: A dictionary of the descriptions for the model’s parameters.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelDescription
type MLModelDescription struct {
	objectivec.Object
}

// MLModelDescriptionFromID constructs a [MLModelDescription] from an objc.ID.
//
// Information about a model, primarily the input and output format for each
// feature the model expects, and optional metadata.
func MLModelDescriptionFromID(id objc.ID) MLModelDescription {
	return MLModelDescription{objectivec.Object{ID: id}}
}
// NOTE: MLModelDescription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLModelDescription] class.
//
// # Accessing feature descriptions
//
//   - [IMLModelDescription.StateDescriptionsByName]: Description of the state features.
//   - [IMLModelDescription.InputDescriptionsByName]: A dictionary of input feature descriptions, which the model keys by the input’s name.
//   - [IMLModelDescription.OutputDescriptionsByName]: A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// # Accessing metadata
//
//   - [IMLModelDescription.ClassLabels]: An array of labels, which can be either strings or a numbers, for classifier models.
//   - [IMLModelDescription.Metadata]: A dictionary of the model’s creation information, such as its description, author, version, and license.
//
// # Accessing prediction names
//
//   - [IMLModelDescription.PredictedFeatureName]: The name of the primary prediction feature output description.
//   - [IMLModelDescription.PredictedProbabilitiesName]: The name of the feature output description for all probabilities of a prediction.
//
// # Accessing update descriptions
//
//   - [IMLModelDescription.IsUpdatable]: A Boolean value that indicates whether you can update the model with additional training.
//   - [IMLModelDescription.TrainingInputDescriptionsByName]: A dictionary of the training input feature descriptions, which the model keys by the input’s name.
//   - [IMLModelDescription.ParameterDescriptionsByKey]: A dictionary of the descriptions for the model’s parameters.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelDescription
type IMLModelDescription interface {
	objectivec.IObject

	// Topic: Accessing feature descriptions

	// Description of the state features.
	StateDescriptionsByName() foundation.INSDictionary
	// A dictionary of input feature descriptions, which the model keys by the input’s name.
	InputDescriptionsByName() foundation.INSDictionary
	// A dictionary of output feature descriptions, which the model keys by the output’s name.
	OutputDescriptionsByName() foundation.INSDictionary

	// Topic: Accessing metadata

	// An array of labels, which can be either strings or a numbers, for classifier models.
	ClassLabels() []objectivec.IObject
	// A dictionary of the model’s creation information, such as its description, author, version, and license.
	Metadata() foundation.INSDictionary

	// Topic: Accessing prediction names

	// The name of the primary prediction feature output description.
	PredictedFeatureName() string
	// The name of the feature output description for all probabilities of a prediction.
	PredictedProbabilitiesName() string

	// Topic: Accessing update descriptions

	// A Boolean value that indicates whether you can update the model with additional training.
	IsUpdatable() bool
	// A dictionary of the training input feature descriptions, which the model keys by the input’s name.
	TrainingInputDescriptionsByName() foundation.INSDictionary
	// A dictionary of the descriptions for the model’s parameters.
	ParameterDescriptionsByKey() foundation.INSDictionary

	// The list of available compute devices that the model’s prediction methods use.
	AvailableComputeDevices() objectivec.IObject
	SetAvailableComputeDevices(value objectivec.IObject)
	// The configuration of the model set during initialization.
	Configuration() IMLModelConfiguration
	SetConfiguration(value IMLModelConfiguration)
	// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
	ModelDescription() IMLModelDescription
	SetModelDescription(value IMLModelDescription)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (m MLModelDescription) Init() MLModelDescription {
	rv := objc.Send[MLModelDescription](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelDescription) Autorelease() MLModelDescription {
	rv := objc.Send[MLModelDescription](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelDescription creates a new MLModelDescription instance.
func NewMLModelDescription() MLModelDescription {
	class := getMLModelDescriptionClass()
	rv := objc.Send[MLModelDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m MLModelDescription) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Description of the state features.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/stateDescriptionsByName
func (m MLModelDescription) StateDescriptionsByName() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("stateDescriptionsByName"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// A dictionary of input feature descriptions, which the model keys by the
// input’s name.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/inputDescriptionsByName
func (m MLModelDescription) InputDescriptionsByName() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputDescriptionsByName"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// A dictionary of output feature descriptions, which the model keys by the
// output’s name.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
func (m MLModelDescription) OutputDescriptionsByName() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputDescriptionsByName"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// An array of labels, which can be either strings or a numbers, for
// classifier models.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/classLabels
func (m MLModelDescription) ClassLabels() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("classLabels"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
// A dictionary of the model’s creation information, such as its
// description, author, version, and license.
//
// # Discussion
// 
// Use the keys defined by [MLModelMetadataKey] to retrieve the dictionary’s
// entries.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/metadata
func (m MLModelDescription) Metadata() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("metadata"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// The name of the primary prediction feature output description.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (m MLModelDescription) PredictedFeatureName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictedFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
// The name of the feature output description for all probabilities of a
// prediction.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedProbabilitiesName
func (m MLModelDescription) PredictedProbabilitiesName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictedProbabilitiesName"))
	return foundation.NSStringFromID(rv).String()
}
// A Boolean value that indicates whether you can update the model with
// additional training.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/isUpdatable
func (m MLModelDescription) IsUpdatable() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isUpdatable"))
	return rv
}
// A dictionary of the training input feature descriptions, which the model
// keys by the input’s name.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/trainingInputDescriptionsByName
func (m MLModelDescription) TrainingInputDescriptionsByName() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("trainingInputDescriptionsByName"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// A dictionary of the descriptions for the model’s parameters.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/parameterDescriptionsByKey
func (m MLModelDescription) ParameterDescriptionsByKey() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parameterDescriptionsByKey"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// The list of available compute devices that the model’s prediction methods
// use.
//
// See: https://developer.apple.com/documentation/coreml/mlmodel/availablecomputedevices-6klyt
func (m MLModelDescription) AvailableComputeDevices() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("availableComputeDevices"))
	return objectivec.Object{ID: rv}
}
func (m MLModelDescription) SetAvailableComputeDevices(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setAvailableComputeDevices:"), value)
}
// The configuration of the model set during initialization.
//
// See: https://developer.apple.com/documentation/coreml/mlmodel/configuration
func (m MLModelDescription) Configuration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
func (m MLModelDescription) SetConfiguration(value IMLModelConfiguration) {
	objc.Send[struct{}](m.ID, objc.Sel("setConfiguration:"), value)
}
// Model information you use at runtime during development, which Xcode also
// displays in its Core ML model editor view.
//
// See: https://developer.apple.com/documentation/coreml/mlmodel/modeldescription
func (m MLModelDescription) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}
func (m MLModelDescription) SetModelDescription(value IMLModelDescription) {
	objc.Send[struct{}](m.ID, objc.Sel("setModelDescription:"), value)
}

