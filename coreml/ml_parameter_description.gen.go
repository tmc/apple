// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLParameterDescription] class.
var (
	_MLParameterDescriptionClass     MLParameterDescriptionClass
	_MLParameterDescriptionClassOnce sync.Once
)

func getMLParameterDescriptionClass() MLParameterDescriptionClass {
	_MLParameterDescriptionClassOnce.Do(func() {
		_MLParameterDescriptionClass = MLParameterDescriptionClass{class: objc.GetClass("MLParameterDescription")}
	})
	return _MLParameterDescriptionClass
}

// GetMLParameterDescriptionClass returns the class object for MLParameterDescription.
func GetMLParameterDescriptionClass() MLParameterDescriptionClass {
	return getMLParameterDescriptionClass()
}

type MLParameterDescriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLParameterDescriptionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLParameterDescriptionClass) Alloc() MLParameterDescription {
	rv := objc.Send[MLParameterDescription](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of a model parameter that includes a default value and a
// constraint, if applicable.
//
// # Describing the model parameter
//
//   - [MLParameterDescription.DefaultValue]: The default value for the parameter.
//   - [MLParameterDescription.Key]: The key for this parameter description value.
//
// # Constraining numeric values
//
//   - [MLParameterDescription.NumericConstraint]: The constraints of this paramter description value, if and only if the value is numerical.
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterDescription
type MLParameterDescription struct {
	objectivec.Object
}

// MLParameterDescriptionFromID constructs a [MLParameterDescription] from an objc.ID.
//
// A description of a model parameter that includes a default value and a
// constraint, if applicable.
func MLParameterDescriptionFromID(id objc.ID) MLParameterDescription {
	return MLParameterDescription{objectivec.Object{ID: id}}
}
// NOTE: MLParameterDescription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLParameterDescription] class.
//
// # Describing the model parameter
//
//   - [IMLParameterDescription.DefaultValue]: The default value for the parameter.
//   - [IMLParameterDescription.Key]: The key for this parameter description value.
//
// # Constraining numeric values
//
//   - [IMLParameterDescription.NumericConstraint]: The constraints of this paramter description value, if and only if the value is numerical.
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterDescription
type IMLParameterDescription interface {
	objectivec.IObject

	// Topic: Describing the model parameter

	// The default value for the parameter.
	DefaultValue() objectivec.IObject
	// The key for this parameter description value.
	Key() IMLParameterKey

	// Topic: Constraining numeric values

	// The constraints of this paramter description value, if and only if the value is numerical.
	NumericConstraint() IMLNumericConstraint

	// A Boolean value that indicates whether you can update the model with additional training.
	IsUpdatable() bool
	SetIsUpdatable(value bool)
	// A dictionary of the descriptions for the model’s parameters.
	ParameterDescriptionsByKey() IMLParameterDescription
	SetParameterDescriptionsByKey(value IMLParameterDescription)
	// A dictionary of the training input feature descriptions, which the model keys by the input’s name.
	TrainingInputDescriptionsByName() IMLFeatureDescription
	SetTrainingInputDescriptionsByName(value IMLFeatureDescription)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p MLParameterDescription) Init() MLParameterDescription {
	rv := objc.Send[MLParameterDescription](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLParameterDescription) Autorelease() MLParameterDescription {
	rv := objc.Send[MLParameterDescription](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLParameterDescription creates a new MLParameterDescription instance.
func NewMLParameterDescription() MLParameterDescription {
	class := getMLParameterDescriptionClass()
	rv := objc.Send[MLParameterDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (p MLParameterDescription) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The default value for the parameter.
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterDescription/defaultValue
func (p MLParameterDescription) DefaultValue() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("defaultValue"))
	return objectivec.Object{ID: rv}
}
// The key for this parameter description value.
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterDescription/key
func (p MLParameterDescription) Key() IMLParameterKey {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("key"))
	return MLParameterKeyFromID(objc.ID(rv))
}
// The constraints of this paramter description value, if and only if the
// value is numerical.
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterDescription/numericConstraint
func (p MLParameterDescription) NumericConstraint() IMLNumericConstraint {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("numericConstraint"))
	return MLNumericConstraintFromID(objc.ID(rv))
}
// A Boolean value that indicates whether you can update the model with
// additional training.
//
// See: https://developer.apple.com/documentation/coreml/mlmodeldescription/isupdatable
func (p MLParameterDescription) IsUpdatable() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isUpdatable"))
	return rv
}
func (p MLParameterDescription) SetIsUpdatable(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setIsUpdatable:"), value)
}
// A dictionary of the descriptions for the model’s parameters.
//
// See: https://developer.apple.com/documentation/coreml/mlmodeldescription/parameterdescriptionsbykey
func (p MLParameterDescription) ParameterDescriptionsByKey() IMLParameterDescription {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("parameterDescriptionsByKey"))
	return MLParameterDescriptionFromID(objc.ID(rv))
}
func (p MLParameterDescription) SetParameterDescriptionsByKey(value IMLParameterDescription) {
	objc.Send[struct{}](p.ID, objc.Sel("setParameterDescriptionsByKey:"), value)
}
// A dictionary of the training input feature descriptions, which the model
// keys by the input’s name.
//
// See: https://developer.apple.com/documentation/coreml/mlmodeldescription/traininginputdescriptionsbyname
func (p MLParameterDescription) TrainingInputDescriptionsByName() IMLFeatureDescription {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("trainingInputDescriptionsByName"))
	return MLFeatureDescriptionFromID(objc.ID(rv))
}
func (p MLParameterDescription) SetTrainingInputDescriptionsByName(value IMLFeatureDescription) {
	objc.Send[struct{}](p.ID, objc.Sel("setTrainingInputDescriptionsByName:"), value)
}

