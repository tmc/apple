// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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

// Class returns the underlying Objective-C class pointer.
func (mc MLFeatureDescriptionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLFeatureDescriptionClass) Alloc() MLFeatureDescription {
	rv := objc.Send[MLFeatureDescription](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLFeatureDescription.AllowsValuesWithDescription]
//   - [MLFeatureDescription.DebugQuickLookObject]
//   - [MLFeatureDescription.DictionaryConstraintCached]
//   - [MLFeatureDescription.ImageConstraintCached]
//   - [MLFeatureDescription.IsAllowedValueError]
//   - [MLFeatureDescription.MultiArrayConstraintCached]
//   - [MLFeatureDescription.SequenceConstraintCached]
//   - [MLFeatureDescription.ValueConstraints]
//   - [MLFeatureDescription.SetValueConstraints]
//   - [MLFeatureDescription.InitWithCoder]
//   - [MLFeatureDescription.InitWithNameTypeOptionalContraints]
//   - [MLFeatureDescription.Optional]
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription
type MLFeatureDescription struct {
	objectivec.Object
}

// MLFeatureDescriptionFromID constructs a [MLFeatureDescription] from an objc.ID.
func MLFeatureDescriptionFromID(id objc.ID) MLFeatureDescription {
	return MLFeatureDescription{objectivec.Object{ID: id}}
}

// Ensure MLFeatureDescription implements IMLFeatureDescription.
var _ IMLFeatureDescription = MLFeatureDescription{}

// An interface definition for the [MLFeatureDescription] class.
//
// # Methods
//
//   - [IMLFeatureDescription.AllowsValuesWithDescription]
//   - [IMLFeatureDescription.DebugQuickLookObject]
//   - [IMLFeatureDescription.DictionaryConstraintCached]
//   - [IMLFeatureDescription.ImageConstraintCached]
//   - [IMLFeatureDescription.IsAllowedValueError]
//   - [IMLFeatureDescription.MultiArrayConstraintCached]
//   - [IMLFeatureDescription.SequenceConstraintCached]
//   - [IMLFeatureDescription.ValueConstraints]
//   - [IMLFeatureDescription.SetValueConstraints]
//   - [IMLFeatureDescription.InitWithCoder]
//   - [IMLFeatureDescription.InitWithNameTypeOptionalContraints]
//   - [IMLFeatureDescription.Optional]
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription
type IMLFeatureDescription interface {
	objectivec.IObject

	// Topic: Methods

	AllowsValuesWithDescription(description objectivec.IObject) bool
	DebugQuickLookObject() objectivec.IObject
	DictionaryConstraintCached() IMLDictionaryConstraint
	ImageConstraintCached() IMLImageConstraint
	IsAllowedValueError(value objectivec.IObject) (bool, error)
	MultiArrayConstraintCached() IMLMultiArrayConstraint
	SequenceConstraintCached() IMLSequenceConstraint
	ValueConstraints() foundation.INSDictionary
	SetValueConstraints(value foundation.INSDictionary)
	InitWithCoder(coder foundation.INSCoder) MLFeatureDescription
	InitWithNameTypeOptionalContraints(name objectivec.IObject, type_ int64, optional bool, contraints objectivec.IObject) MLFeatureDescription
	Optional() bool
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

// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/initWithCoder:
func NewFeatureDescriptionWithCoder(coder objectivec.IObject) MLFeatureDescription {
	instance := getMLFeatureDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLFeatureDescriptionFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/initWithName:type:optional:contraints:
func NewFeatureDescriptionWithNameTypeOptionalContraints(name objectivec.IObject, type_ int64, optional bool, contraints objectivec.IObject) MLFeatureDescription {
	instance := getMLFeatureDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:type:optional:contraints:"), name, type_, optional, contraints)
	return MLFeatureDescriptionFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/allowsValuesWithDescription:
func (f MLFeatureDescription) AllowsValuesWithDescription(description objectivec.IObject) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("allowsValuesWithDescription:"), description)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/debugQuickLookObject
func (f MLFeatureDescription) DebugQuickLookObject() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("debugQuickLookObject"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/isAllowedValue:error:
func (f MLFeatureDescription) IsAllowedValueError(value objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("isAllowedValue:error:"), value, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("isAllowedValue:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/initWithCoder:
func (f MLFeatureDescription) InitWithCoder(coder foundation.INSCoder) MLFeatureDescription {
	rv := objc.Send[MLFeatureDescription](f.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/initWithName:type:optional:contraints:
func (f MLFeatureDescription) InitWithNameTypeOptionalContraints(name objectivec.IObject, type_ int64, optional bool, contraints objectivec.IObject) MLFeatureDescription {
	rv := objc.Send[MLFeatureDescription](f.ID, objc.Sel("initWithName:type:optional:contraints:"), name, type_, optional, contraints)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/featureDescriptionWithName:type:optional:constraints:
func (_MLFeatureDescriptionClass MLFeatureDescriptionClass) FeatureDescriptionWithNameTypeOptionalConstraints(name objectivec.IObject, type_ int64, optional bool, constraints objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureDescriptionClass.class), objc.Sel("featureDescriptionWithName:type:optional:constraints:"), name, type_, optional, constraints)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/supportsSecureCoding
func (_MLFeatureDescriptionClass MLFeatureDescriptionClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLFeatureDescriptionClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/dictionaryConstraintCached
func (f MLFeatureDescription) DictionaryConstraintCached() IMLDictionaryConstraint {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("dictionaryConstraintCached"))
	return MLDictionaryConstraintFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/imageConstraintCached
func (f MLFeatureDescription) ImageConstraintCached() IMLImageConstraint {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("imageConstraintCached"))
	return MLImageConstraintFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/multiArrayConstraintCached
func (f MLFeatureDescription) MultiArrayConstraintCached() IMLMultiArrayConstraint {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("multiArrayConstraintCached"))
	return MLMultiArrayConstraintFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/optional
func (f MLFeatureDescription) Optional() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("optional"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/sequenceConstraintCached
func (f MLFeatureDescription) SequenceConstraintCached() IMLSequenceConstraint {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("sequenceConstraintCached"))
	return MLSequenceConstraintFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/valueConstraints
func (f MLFeatureDescription) ValueConstraints() foundation.INSDictionary {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("valueConstraints"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (f MLFeatureDescription) SetValueConstraints(value foundation.INSDictionary) {
	objc.Send[struct{}](f.ID, objc.Sel("setValueConstraints:"), value)
}
