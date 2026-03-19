// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLFeatureDescription] class.
var (
	_MLFeatureDescriptionClass     MLFeatureDescriptionClass
	_MLFeatureDescriptionClassOnce sync.Once
)

func getMLFeatureDescriptionClass() MLFeatureDescriptionClass {
	_MLFeatureDescriptionClassOnce.Do(func() {
		_MLFeatureDescriptionClass = MLFeatureDescriptionClass{class: objc.GetClass("MLFeatureDescription")}
	})
	return _MLFeatureDescriptionClass
}

// GetMLFeatureDescriptionClass returns the class object for MLFeatureDescription.
func GetMLFeatureDescriptionClass() MLFeatureDescriptionClass {
	return getMLFeatureDescriptionClass()
}

type MLFeatureDescriptionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLFeatureDescriptionClass) Alloc() MLFeatureDescription {
	rv := objc.Send[MLFeatureDescription](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The name, type, and constraints of an input or output feature.
//
// # Overview
// 
// In Core ML, a is a single input or output of a model. A model can have any
// number of or . Each feature has a name and a value type, which are defined
// in the feature’s [MLFeatureDescription]. Model authors use feature
// descriptions to help developers integrate their model properly. Each
// [MLFeatureDescription] instance has read-only properties that indicate the
// feature’s name, its type, and whether it’s optional.
// 
// For examples of features, see [Integrating a Core ML Model into Your App].
// Note the three input features named `solarPanels`, `greenhouses`, and
// `size`, and the output feature is named `price`. All four features are of
// type [Double].
// 
// An [MLFeatureDescription] may also include constraints, which specify the
// limitations of the model’s input and output features. For each input
// feature, the constraints describe what values the model expects from your
// app. For each output feature, the constraints describe what values your app
// should expect from the model. You can also write code to inspect these
// descriptions before using the model in your app.
//
// [Integrating a Core ML Model into Your App]: https://developer.apple.com/documentation/CoreML/integrating-a-core-ml-model-into-your-app
//
// # Inspecting a feature
//
//   - [MLFeatureDescription.Name]: The name of this feature.
//   - [MLFeatureDescription.Type]: The type of this feature.
//   - [MLFeatureDescription.Optional]: A Boolean value that indicates whether this feature is optional.
//
// # Checking for validity
//
//   - [MLFeatureDescription.IsAllowedValue]: Checks whether the model will accept an input feature value.
//
// # Accessing feature constraints
//
//   - [MLFeatureDescription.StateConstraint]: The state feature value constraint.
//   - [MLFeatureDescription.ImageConstraint]: The size and format constraints for an image feature.
//   - [MLFeatureDescription.DictionaryConstraint]: The constraint for a dictionary feature.
//   - [MLFeatureDescription.MultiArrayConstraint]: The constraints on a multidimensional array feature.
//   - [MLFeatureDescription.SequenceConstraint]: The constraints for a sequence feature.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription
type MLFeatureDescription struct {
	objectivec.Object
}

// MLFeatureDescriptionFromID constructs a [MLFeatureDescription] from an objc.ID.
//
// The name, type, and constraints of an input or output feature.
func MLFeatureDescriptionFromID(id objc.ID) MLFeatureDescription {
	return MLFeatureDescription{objectivec.Object{ID: id}}
}
// NOTE: MLFeatureDescription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLFeatureDescription] class.
//
// # Inspecting a feature
//
//   - [IMLFeatureDescription.Name]: The name of this feature.
//   - [IMLFeatureDescription.Type]: The type of this feature.
//   - [IMLFeatureDescription.Optional]: A Boolean value that indicates whether this feature is optional.
//
// # Checking for validity
//
//   - [IMLFeatureDescription.IsAllowedValue]: Checks whether the model will accept an input feature value.
//
// # Accessing feature constraints
//
//   - [IMLFeatureDescription.StateConstraint]: The state feature value constraint.
//   - [IMLFeatureDescription.ImageConstraint]: The size and format constraints for an image feature.
//   - [IMLFeatureDescription.DictionaryConstraint]: The constraint for a dictionary feature.
//   - [IMLFeatureDescription.MultiArrayConstraint]: The constraints on a multidimensional array feature.
//   - [IMLFeatureDescription.SequenceConstraint]: The constraints for a sequence feature.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription
type IMLFeatureDescription interface {
	objectivec.IObject

	// Topic: Inspecting a feature

	// The name of this feature.
	Name() string
	// The type of this feature.
	Type() MLFeatureType
	// A Boolean value that indicates whether this feature is optional.
	Optional() bool

	// Topic: Checking for validity

	// Checks whether the model will accept an input feature value.
	IsAllowedValue(value IMLFeatureValue) bool

	// Topic: Accessing feature constraints

	// The state feature value constraint.
	StateConstraint() IMLStateConstraint
	// The size and format constraints for an image feature.
	ImageConstraint() IMLImageConstraint
	// The constraint for a dictionary feature.
	DictionaryConstraint() IMLDictionaryConstraint
	// The constraints on a multidimensional array feature.
	MultiArrayConstraint() IMLMultiArrayConstraint
	// The constraints for a sequence feature.
	SequenceConstraint() IMLSequenceConstraint

	// A dictionary of input feature descriptions, which the model keys by the input’s name.
	InputDescriptionsByName() IMLFeatureDescription
	SetInputDescriptionsByName(value IMLFeatureDescription)
	// A dictionary of output feature descriptions, which the model keys by the output’s name.
	OutputDescriptionsByName() IMLFeatureDescription
	SetOutputDescriptionsByName(value IMLFeatureDescription)
	// Description of the state features.
	StateDescriptionsByName() IMLFeatureDescription
	SetStateDescriptionsByName(value IMLFeatureDescription)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f MLFeatureDescription) Init() MLFeatureDescription {
	rv := objc.Send[MLFeatureDescription](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MLFeatureDescription) Autorelease() MLFeatureDescription {
	rv := objc.Send[MLFeatureDescription](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLFeatureDescription creates a new MLFeatureDescription instance.
func NewMLFeatureDescription() MLFeatureDescription {
	class := getMLFeatureDescriptionClass()
	rv := objc.Send[MLFeatureDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Checks whether the model will accept an input feature value.
//
// value: Given the [MLFeatureValue], is it compatible with the [MLFeatureType] of
// this [MLFeatureDescription].
//
// # Return Value
// 
// [True] if the given [MLFeatureValue] is acceptable to the model’s input
// feature, `false` otherwise.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/isAllowedValue(_:)
func (f MLFeatureDescription) IsAllowedValue(value IMLFeatureValue) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isAllowedValue:"), value)
	return rv
}
func (f MLFeatureDescription) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The name of this feature.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/name
func (f MLFeatureDescription) Name() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
// The type of this feature.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/type
func (f MLFeatureDescription) Type() MLFeatureType {
	rv := objc.Send[MLFeatureType](f.ID, objc.Sel("type"))
	return MLFeatureType(rv)
}
// A Boolean value that indicates whether this feature is optional.
//
// # Discussion
// 
// Optional values can be `nil`, for models that have inputs related to
// features being present or not.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/isOptional
func (f MLFeatureDescription) Optional() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isOptional"))
	return rv
}
// The state feature value constraint.
//
// # Discussion
// 
// The property has a value when `XCUIElementTypeType == MLFeatureTypeState`.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/stateConstraint
func (f MLFeatureDescription) StateConstraint() IMLStateConstraint {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("stateConstraint"))
	return MLStateConstraintFromID(objc.ID(rv))
}
// The size and format constraints for an image feature.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/imageConstraint
func (f MLFeatureDescription) ImageConstraint() IMLImageConstraint {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("imageConstraint"))
	return MLImageConstraintFromID(objc.ID(rv))
}
// The constraint for a dictionary feature.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/dictionaryConstraint
func (f MLFeatureDescription) DictionaryConstraint() IMLDictionaryConstraint {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("dictionaryConstraint"))
	return MLDictionaryConstraintFromID(objc.ID(rv))
}
// The constraints on a multidimensional array feature.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/multiArrayConstraint
func (f MLFeatureDescription) MultiArrayConstraint() IMLMultiArrayConstraint {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("multiArrayConstraint"))
	return MLMultiArrayConstraintFromID(objc.ID(rv))
}
// The constraints for a sequence feature.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/sequenceConstraint
func (f MLFeatureDescription) SequenceConstraint() IMLSequenceConstraint {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("sequenceConstraint"))
	return MLSequenceConstraintFromID(objc.ID(rv))
}
// A dictionary of input feature descriptions, which the model keys by the
// input’s name.
//
// See: https://developer.apple.com/documentation/coreml/mlmodeldescription/inputdescriptionsbyname
func (f MLFeatureDescription) InputDescriptionsByName() IMLFeatureDescription {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("inputDescriptionsByName"))
	return MLFeatureDescriptionFromID(objc.ID(rv))
}
func (f MLFeatureDescription) SetInputDescriptionsByName(value IMLFeatureDescription) {
	objc.Send[struct{}](f.ID, objc.Sel("setInputDescriptionsByName:"), value)
}
// A dictionary of output feature descriptions, which the model keys by the
// output’s name.
//
// See: https://developer.apple.com/documentation/coreml/mlmodeldescription/outputdescriptionsbyname
func (f MLFeatureDescription) OutputDescriptionsByName() IMLFeatureDescription {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("outputDescriptionsByName"))
	return MLFeatureDescriptionFromID(objc.ID(rv))
}
func (f MLFeatureDescription) SetOutputDescriptionsByName(value IMLFeatureDescription) {
	objc.Send[struct{}](f.ID, objc.Sel("setOutputDescriptionsByName:"), value)
}
// Description of the state features.
//
// See: https://developer.apple.com/documentation/coreml/mlmodeldescription/statedescriptionsbyname
func (f MLFeatureDescription) StateDescriptionsByName() IMLFeatureDescription {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("stateDescriptionsByName"))
	return MLFeatureDescriptionFromID(objc.ID(rv))
}
func (f MLFeatureDescription) SetStateDescriptionsByName(value IMLFeatureDescription) {
	objc.Send[struct{}](f.ID, objc.Sel("setStateDescriptionsByName:"), value)
}

