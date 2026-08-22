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
	rv := objc.SendIfResponds[MLFeatureDescription](objc.ID(mc.class), objc.Sel("alloc"))
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
//   - [MLFeatureDescription.InitWithNameTypeOptionalContraints]
//   - [MLFeatureDescription.Optional]
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
//   - [IMLFeatureDescription.InitWithNameTypeOptionalContraints]
//   - [IMLFeatureDescription.Optional]
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
	InitWithNameTypeOptionalContraints(name objectivec.IObject, type_ int64, optional bool, contraints objectivec.IObject) MLFeatureDescription
	Optional() bool
}

// Init initializes the instance.
func (m MLFeatureDescription) Init() MLFeatureDescription {
	rv := objc.SendIfResponds[MLFeatureDescription](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLFeatureDescription) Autorelease() MLFeatureDescription {
	rv := objc.SendIfResponds[MLFeatureDescription](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLFeatureDescription creates a new MLFeatureDescription instance.
func NewMLFeatureDescription() MLFeatureDescription {
	class := getMLFeatureDescriptionClass()
	rv := objc.SendIfResponds[MLFeatureDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewFeatureDescriptionWithNameTypeOptionalContraints(name objectivec.IObject, type_ int64, optional bool, contraints objectivec.IObject) MLFeatureDescription {
	instance := getMLFeatureDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithName:type:optional:contraints:"), name, type_, optional, contraints)
	return MLFeatureDescriptionFromID(rv)
}

func (m MLFeatureDescription) AllowsValuesWithDescription(description objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("allowsValuesWithDescription:"), description)
	return rv
}
func (m MLFeatureDescription) DebugQuickLookObject() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugQuickLookObject"))
	return objectivec.Object{ID: rv}
}
func (m MLFeatureDescription) IsAllowedValueError(value objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("isAllowedValue:error:"), value, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("isAllowedValue:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLFeatureDescription) InitWithNameTypeOptionalContraints(name objectivec.IObject, type_ int64, optional bool, contraints objectivec.IObject) MLFeatureDescription {
	rv := objc.SendIfResponds[MLFeatureDescription](m.ID, objc.Sel("initWithName:type:optional:contraints:"), name, type_, optional, contraints)
	return rv
}

func (_MLFeatureDescriptionClass MLFeatureDescriptionClass) FeatureDescriptionWithNameTypeOptionalConstraints(name objectivec.IObject, type_ int64, optional bool, constraints objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLFeatureDescriptionClass.class), objc.Sel("featureDescriptionWithName:type:optional:constraints:"), name, type_, optional, constraints)
	return objectivec.Object{ID: rv}
}
func (_MLFeatureDescriptionClass MLFeatureDescriptionClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLFeatureDescriptionClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (m MLFeatureDescription) DictionaryConstraintCached() IMLDictionaryConstraint {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("dictionaryConstraintCached"))
	return MLDictionaryConstraintFromID(objc.ID(rv))
}
func (m MLFeatureDescription) ImageConstraintCached() IMLImageConstraint {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("imageConstraintCached"))
	return MLImageConstraintFromID(objc.ID(rv))
}
func (m MLFeatureDescription) MultiArrayConstraintCached() IMLMultiArrayConstraint {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("multiArrayConstraintCached"))
	return MLMultiArrayConstraintFromID(objc.ID(rv))
}
func (m MLFeatureDescription) Optional() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("optional"))
	return rv
}
func (m MLFeatureDescription) SequenceConstraintCached() IMLSequenceConstraint {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("sequenceConstraintCached"))
	return MLSequenceConstraintFromID(objc.ID(rv))
}
func (m MLFeatureDescription) ValueConstraints() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("valueConstraints"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLFeatureDescription) SetValueConstraints(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setValueConstraints:"), value)
}
