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
	rv := objc.Send[MLMultiArrayConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLMultiArrayConstraint.DefaultOptionalValue]
//   - [MLMultiArrayConstraint.IsAllowedDataTypeError]
//   - [MLMultiArrayConstraint.IsAllowedShapeError]
//   - [MLMultiArrayConstraint.IsAllowedValueError]
//   - [MLMultiArrayConstraint.IsAllowedValueIsNeuralNetworkInputOrOutputUsingRank5MappingFeatureNameError]
//   - [MLMultiArrayConstraint.InitWithCoder]
//   - [MLMultiArrayConstraint.InitWithShapeDataTypeShapeConstraint]
//   - [MLMultiArrayConstraint.InitWithShapeDataTypeShapeConstraintDefaultOptionalValue]
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint
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
//   - [IMLMultiArrayConstraint.InitWithCoder]
//   - [IMLMultiArrayConstraint.InitWithShapeDataTypeShapeConstraint]
//   - [IMLMultiArrayConstraint.InitWithShapeDataTypeShapeConstraintDefaultOptionalValue]
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint
type IMLMultiArrayConstraint interface {
	objectivec.IObject

	// Topic: Methods

	DefaultOptionalValue() objectivec.IObject
	IsAllowedDataTypeError(type_ int64) (bool, error)
	IsAllowedShapeError(shape objectivec.IObject) (bool, error)
	IsAllowedValueError(value objectivec.IObject) (bool, error)
	IsAllowedValueIsNeuralNetworkInputOrOutputUsingRank5MappingFeatureNameError(value objectivec.IObject, output bool, rank5Mapping bool, name objectivec.IObject) (bool, error)
	InitWithCoder(coder foundation.INSCoder) MLMultiArrayConstraint
	InitWithShapeDataTypeShapeConstraint(shape objectivec.IObject, type_ int64, constraint objectivec.IObject) MLMultiArrayConstraint
	InitWithShapeDataTypeShapeConstraintDefaultOptionalValue(shape objectivec.IObject, type_ int64, constraint objectivec.IObject, value objectivec.IObject) MLMultiArrayConstraint
}

// Init initializes the instance.
func (m MLMultiArrayConstraint) Init() MLMultiArrayConstraint {
	rv := objc.Send[MLMultiArrayConstraint](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLMultiArrayConstraint) Autorelease() MLMultiArrayConstraint {
	rv := objc.Send[MLMultiArrayConstraint](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLMultiArrayConstraint creates a new MLMultiArrayConstraint instance.
func NewMLMultiArrayConstraint() MLMultiArrayConstraint {
	class := getMLMultiArrayConstraintClass()
	rv := objc.Send[MLMultiArrayConstraint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/initWithCoder:
func NewMultiArrayConstraintWithCoder(coder objectivec.IObject) MLMultiArrayConstraint {
	instance := getMLMultiArrayConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLMultiArrayConstraintFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/initWithShape:dataType:shapeConstraint:
func NewMultiArrayConstraintWithShapeDataTypeShapeConstraint(shape objectivec.IObject, type_ int64, constraint objectivec.IObject) MLMultiArrayConstraint {
	instance := getMLMultiArrayConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithShape:dataType:shapeConstraint:"), shape, type_, constraint)
	return MLMultiArrayConstraintFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/initWithShape:dataType:shapeConstraint:defaultOptionalValue:
func NewMultiArrayConstraintWithShapeDataTypeShapeConstraintDefaultOptionalValue(shape objectivec.IObject, type_ int64, constraint objectivec.IObject, value objectivec.IObject) MLMultiArrayConstraint {
	instance := getMLMultiArrayConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithShape:dataType:shapeConstraint:defaultOptionalValue:"), shape, type_, constraint, value)
	return MLMultiArrayConstraintFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/defaultOptionalValue
func (m MLMultiArrayConstraint) DefaultOptionalValue() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("defaultOptionalValue"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/isAllowedDataType:error:
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

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/isAllowedShape:error:
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

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/isAllowedValue:error:
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

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/isAllowedValue:isNeuralNetworkInputOrOutput:usingRank5Mapping:featureName:error:
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

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/initWithCoder:
func (m MLMultiArrayConstraint) InitWithCoder(coder foundation.INSCoder) MLMultiArrayConstraint {
	rv := objc.Send[MLMultiArrayConstraint](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/initWithShape:dataType:shapeConstraint:
func (m MLMultiArrayConstraint) InitWithShapeDataTypeShapeConstraint(shape objectivec.IObject, type_ int64, constraint objectivec.IObject) MLMultiArrayConstraint {
	rv := objc.Send[MLMultiArrayConstraint](m.ID, objc.Sel("initWithShape:dataType:shapeConstraint:"), shape, type_, constraint)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/initWithShape:dataType:shapeConstraint:defaultOptionalValue:
func (m MLMultiArrayConstraint) InitWithShapeDataTypeShapeConstraintDefaultOptionalValue(shape objectivec.IObject, type_ int64, constraint objectivec.IObject, value objectivec.IObject) MLMultiArrayConstraint {
	rv := objc.Send[MLMultiArrayConstraint](m.ID, objc.Sel("initWithShape:dataType:shapeConstraint:defaultOptionalValue:"), shape, type_, constraint, value)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/constraintWithShape:dataType:
func (_MLMultiArrayConstraintClass MLMultiArrayConstraintClass) ConstraintWithShapeDataType(shape objectivec.IObject, type_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLMultiArrayConstraintClass.class), objc.Sel("constraintWithShape:dataType:"), shape, type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/constraintWithShape:dataType:shapeConstraint:
func (_MLMultiArrayConstraintClass MLMultiArrayConstraintClass) ConstraintWithShapeDataTypeShapeConstraint(shape objectivec.IObject, type_ int64, constraint objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLMultiArrayConstraintClass.class), objc.Sel("constraintWithShape:dataType:shapeConstraint:"), shape, type_, constraint)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/supportsSecureCoding
func (_MLMultiArrayConstraintClass MLMultiArrayConstraintClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLMultiArrayConstraintClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
