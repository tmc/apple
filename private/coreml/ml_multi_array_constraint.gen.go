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

// The class instance for the [MLMultiArrayConstraint] class.
var (
	_MLMultiArrayConstraintClass     MLMultiArrayConstraintClass
	_MLMultiArrayConstraintClassOnce sync.Once
)

func getMLMultiArrayConstraintClass() MLMultiArrayConstraintClass {
	_MLMultiArrayConstraintClassOnce.Do(func() {
		_MLMultiArrayConstraintClass = MLMultiArrayConstraintClass{class: objc.GetClass("MLMultiArrayConstraint")}
	})
	return _MLMultiArrayConstraintClass
}

// GetMLMultiArrayConstraintClass returns the class object for MLMultiArrayConstraint.
func GetMLMultiArrayConstraintClass() MLMultiArrayConstraintClass {
	return getMLMultiArrayConstraintClass()
}

type MLMultiArrayConstraintClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLMultiArrayConstraintClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLMultiArrayConstraintClass) Alloc() MLMultiArrayConstraint {
	rv := objc.SendIfResponds[MLMultiArrayConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLMultiArrayConstraint.DefaultOptionalValue]
//   - [MLMultiArrayConstraint.IsAllowedDataTypeError]
//   - [MLMultiArrayConstraint.IsAllowedShapeError]
//   - [MLMultiArrayConstraint.IsAllowedValueError]
//   - [MLMultiArrayConstraint.IsAllowedValueIsNeuralNetworkInputOrOutputUsingRank5MappingFeatureNameError]
//   - [MLMultiArrayConstraint.InitWithShapeDataTypeShapeConstraint]
//   - [MLMultiArrayConstraint.InitWithShapeDataTypeShapeConstraintDefaultOptionalValue]
type MLMultiArrayConstraint struct {
	objectivec.Object
}

// MLMultiArrayConstraintFromID constructs a [MLMultiArrayConstraint] from an objc.ID.
func MLMultiArrayConstraintFromID(id objc.ID) MLMultiArrayConstraint {
	return MLMultiArrayConstraint{objectivec.Object{ID: id}}
}

// Ensure MLMultiArrayConstraint implements IMLMultiArrayConstraint.
var _ IMLMultiArrayConstraint = MLMultiArrayConstraint{}

// An interface definition for the [MLMultiArrayConstraint] class.
//
// # Methods
//
//   - [IMLMultiArrayConstraint.DefaultOptionalValue]
//   - [IMLMultiArrayConstraint.IsAllowedDataTypeError]
//   - [IMLMultiArrayConstraint.IsAllowedShapeError]
//   - [IMLMultiArrayConstraint.IsAllowedValueError]
//   - [IMLMultiArrayConstraint.IsAllowedValueIsNeuralNetworkInputOrOutputUsingRank5MappingFeatureNameError]
//   - [IMLMultiArrayConstraint.InitWithShapeDataTypeShapeConstraint]
//   - [IMLMultiArrayConstraint.InitWithShapeDataTypeShapeConstraintDefaultOptionalValue]
type IMLMultiArrayConstraint interface {
	objectivec.IObject

	// Topic: Methods

	DefaultOptionalValue() objectivec.IObject
	IsAllowedDataTypeError(type_ int64) (bool, error)
	IsAllowedShapeError(shape objectivec.IObject) (bool, error)
	IsAllowedValueError(value objectivec.IObject) (bool, error)
	IsAllowedValueIsNeuralNetworkInputOrOutputUsingRank5MappingFeatureNameError(value objectivec.IObject, output bool, rank5Mapping bool, name objectivec.IObject) (bool, error)
	InitWithShapeDataTypeShapeConstraint(shape objectivec.IObject, type_ int64, constraint objectivec.IObject) MLMultiArrayConstraint
	InitWithShapeDataTypeShapeConstraintDefaultOptionalValue(shape objectivec.IObject, type_ int64, constraint objectivec.IObject, value objectivec.IObject) MLMultiArrayConstraint
}

// Init initializes the instance.
func (m MLMultiArrayConstraint) Init() MLMultiArrayConstraint {
	rv := objc.SendIfResponds[MLMultiArrayConstraint](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLMultiArrayConstraint) Autorelease() MLMultiArrayConstraint {
	rv := objc.SendIfResponds[MLMultiArrayConstraint](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLMultiArrayConstraint creates a new MLMultiArrayConstraint instance.
func NewMLMultiArrayConstraint() MLMultiArrayConstraint {
	class := getMLMultiArrayConstraintClass()
	rv := objc.SendIfResponds[MLMultiArrayConstraint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewMultiArrayConstraintWithShapeDataTypeShapeConstraint(shape objectivec.IObject, type_ int64, constraint objectivec.IObject) MLMultiArrayConstraint {
	instance := getMLMultiArrayConstraintClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithShape:dataType:shapeConstraint:"), shape, type_, constraint)
	return MLMultiArrayConstraintFromID(rv)
}

func NewMultiArrayConstraintWithShapeDataTypeShapeConstraintDefaultOptionalValue(shape objectivec.IObject, type_ int64, constraint objectivec.IObject, value objectivec.IObject) MLMultiArrayConstraint {
	instance := getMLMultiArrayConstraintClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithShape:dataType:shapeConstraint:defaultOptionalValue:"), shape, type_, constraint, value)
	return MLMultiArrayConstraintFromID(rv)
}

func (m MLMultiArrayConstraint) DefaultOptionalValue() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("defaultOptionalValue"))
	return objectivec.Object{ID: rv}
}
func (m MLMultiArrayConstraint) IsAllowedDataTypeError(type_ int64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("isAllowedDataType:error:"), type_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("isAllowedDataType:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLMultiArrayConstraint) IsAllowedShapeError(shape objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("isAllowedShape:error:"), shape, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("isAllowedShape:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLMultiArrayConstraint) IsAllowedValueError(value objectivec.IObject) (bool, error) {
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
func (m MLMultiArrayConstraint) IsAllowedValueIsNeuralNetworkInputOrOutputUsingRank5MappingFeatureNameError(value objectivec.IObject, output bool, rank5Mapping bool, name objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("isAllowedValue:isNeuralNetworkInputOrOutput:usingRank5Mapping:featureName:error:"), value, output, rank5Mapping, name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("isAllowedValue:isNeuralNetworkInputOrOutput:usingRank5Mapping:featureName:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLMultiArrayConstraint) InitWithShapeDataTypeShapeConstraint(shape objectivec.IObject, type_ int64, constraint objectivec.IObject) MLMultiArrayConstraint {
	rv := objc.SendIfResponds[MLMultiArrayConstraint](m.ID, objc.Sel("initWithShape:dataType:shapeConstraint:"), shape, type_, constraint)
	return rv
}
func (m MLMultiArrayConstraint) InitWithShapeDataTypeShapeConstraintDefaultOptionalValue(shape objectivec.IObject, type_ int64, constraint objectivec.IObject, value objectivec.IObject) MLMultiArrayConstraint {
	rv := objc.SendIfResponds[MLMultiArrayConstraint](m.ID, objc.Sel("initWithShape:dataType:shapeConstraint:defaultOptionalValue:"), shape, type_, constraint, value)
	return rv
}

func (_MLMultiArrayConstraintClass MLMultiArrayConstraintClass) ConstraintWithShapeDataType(shape objectivec.IObject, type_ int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLMultiArrayConstraintClass.class), objc.Sel("constraintWithShape:dataType:"), shape, type_)
	return objectivec.Object{ID: rv}
}
func (_MLMultiArrayConstraintClass MLMultiArrayConstraintClass) ConstraintWithShapeDataTypeShapeConstraint(shape objectivec.IObject, type_ int64, constraint objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLMultiArrayConstraintClass.class), objc.Sel("constraintWithShape:dataType:shapeConstraint:"), shape, type_, constraint)
	return objectivec.Object{ID: rv}
}
func (_MLMultiArrayConstraintClass MLMultiArrayConstraintClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLMultiArrayConstraintClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
