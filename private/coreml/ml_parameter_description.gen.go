// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
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

// # Methods
//
//   - [MLParameterDescription.DefaultValue]
//   - [MLParameterDescription.SetDefaultValue]
//   - [MLParameterDescription.Key]
//   - [MLParameterDescription.SetKey]
//   - [MLParameterDescription.NumericConstraint]
//   - [MLParameterDescription.SetNumericConstraint]
type MLParameterDescription struct {
	objectivec.Object
}

// MLParameterDescriptionFromID constructs a [MLParameterDescription] from an objc.ID.
func MLParameterDescriptionFromID(id objc.ID) MLParameterDescription {
	return MLParameterDescription{objectivec.Object{ID: id}}
}

// Ensure MLParameterDescription implements IMLParameterDescription.
var _ IMLParameterDescription = MLParameterDescription{}

// An interface definition for the [MLParameterDescription] class.
//
// # Methods
//
//   - [IMLParameterDescription.DefaultValue]
//   - [IMLParameterDescription.SetDefaultValue]
//   - [IMLParameterDescription.Key]
//   - [IMLParameterDescription.SetKey]
//   - [IMLParameterDescription.NumericConstraint]
//   - [IMLParameterDescription.SetNumericConstraint]
type IMLParameterDescription interface {
	objectivec.IObject

	// Topic: Methods

	DefaultValue() objectivec.IObject
	SetDefaultValue(value objectivec.IObject)
	Key() IMLParameterKey
	SetKey(value IMLParameterKey)
	NumericConstraint() IMLNumericConstraint
	SetNumericConstraint(value IMLNumericConstraint)
}

// Init initializes the instance.
func (m MLParameterDescription) Init() MLParameterDescription {
	rv := objc.Send[MLParameterDescription](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLParameterDescription) Autorelease() MLParameterDescription {
	rv := objc.Send[MLParameterDescription](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLParameterDescription creates a new MLParameterDescription instance.
func NewMLParameterDescription() MLParameterDescription {
	class := getMLParameterDescriptionClass()
	rv := objc.Send[MLParameterDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_MLParameterDescriptionClass MLParameterDescriptionClass) ParameterDescriptionForKeyBoolParameterSpec(key objectivec.IObject, spec unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterDescriptionClass.class), objc.Sel("parameterDescriptionForKey:boolParameterSpec:"), key, spec)
	return objectivec.Object{ID: rv}
}
func (_MLParameterDescriptionClass MLParameterDescriptionClass) ParameterDescriptionForKeyDefaultValueNumericConstraint(key objectivec.IObject, value objectivec.IObject, constraint objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterDescriptionClass.class), objc.Sel("parameterDescriptionForKey:defaultValue:numericConstraint:"), key, value, constraint)
	return objectivec.Object{ID: rv}
}
func (_MLParameterDescriptionClass MLParameterDescriptionClass) ParameterDescriptionForKeyDoubleParameterSpec(key objectivec.IObject, spec unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterDescriptionClass.class), objc.Sel("parameterDescriptionForKey:doubleParameterSpec:"), key, spec)
	return objectivec.Object{ID: rv}
}
func (_MLParameterDescriptionClass MLParameterDescriptionClass) ParameterDescriptionForKeyInt64ParameterSpec(key objectivec.IObject, spec unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterDescriptionClass.class), objc.Sel("parameterDescriptionForKey:int64ParameterSpec:"), key, spec)
	return objectivec.Object{ID: rv}
}
func (_MLParameterDescriptionClass MLParameterDescriptionClass) ParameterDescriptionForKeyStringParameterSpec(key objectivec.IObject, spec unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterDescriptionClass.class), objc.Sel("parameterDescriptionForKey:stringParameterSpec:"), key, spec)
	return objectivec.Object{ID: rv}
}
func (_MLParameterDescriptionClass MLParameterDescriptionClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLParameterDescriptionClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (m MLParameterDescription) DefaultValue() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("defaultValue"))
	return objectivec.Object{ID: rv}
}
func (m MLParameterDescription) SetDefaultValue(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setDefaultValue:"), value)
}
func (m MLParameterDescription) Key() IMLParameterKey {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("key"))
	return MLParameterKeyFromID(objc.ID(rv))
}
func (m MLParameterDescription) SetKey(value IMLParameterKey) {
	objc.Send[struct{}](m.ID, objc.Sel("setKey:"), value)
}
func (m MLParameterDescription) NumericConstraint() IMLNumericConstraint {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("numericConstraint"))
	return MLNumericConstraintFromID(objc.ID(rv))
}
func (m MLParameterDescription) SetNumericConstraint(value IMLNumericConstraint) {
	objc.Send[struct{}](m.ID, objc.Sel("setNumericConstraint:"), value)
}
