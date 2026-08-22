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

// The class instance for the [MLStateConstraint] class.
var (
	_MLStateConstraintClass     MLStateConstraintClass
	_MLStateConstraintClassOnce sync.Once
)

func getMLStateConstraintClass() MLStateConstraintClass {
	_MLStateConstraintClassOnce.Do(func() {
		_MLStateConstraintClass = MLStateConstraintClass{class: objc.GetClass("MLStateConstraint")}
	})
	return _MLStateConstraintClass
}

// GetMLStateConstraintClass returns the class object for MLStateConstraint.
func GetMLStateConstraintClass() MLStateConstraintClass {
	return getMLStateConstraintClass()
}

type MLStateConstraintClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLStateConstraintClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLStateConstraintClass) Alloc() MLStateConstraint {
	rv := objc.SendIfResponds[MLStateConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLStateConstraint.IsAllowedDataTypeError]
//   - [MLStateConstraint.IsAllowedShapeError]
//   - [MLStateConstraint.IsAllowedValueError]
//   - [MLStateConstraint.Shape]
//   - [MLStateConstraint.InitWithShapeDataTypeShapeConstraintDefaultOptionalValue]
type MLStateConstraint struct {
	objectivec.Object
}

// MLStateConstraintFromID constructs a [MLStateConstraint] from an objc.ID.
func MLStateConstraintFromID(id objc.ID) MLStateConstraint {
	return MLStateConstraint{objectivec.Object{ID: id}}
}

// Ensure MLStateConstraint implements IMLStateConstraint.
var _ IMLStateConstraint = MLStateConstraint{}

// An interface definition for the [MLStateConstraint] class.
//
// # Methods
//
//   - [IMLStateConstraint.IsAllowedDataTypeError]
//   - [IMLStateConstraint.IsAllowedShapeError]
//   - [IMLStateConstraint.IsAllowedValueError]
//   - [IMLStateConstraint.Shape]
//   - [IMLStateConstraint.InitWithShapeDataTypeShapeConstraintDefaultOptionalValue]
type IMLStateConstraint interface {
	objectivec.IObject

	// Topic: Methods

	IsAllowedDataTypeError(type_ int64) (bool, error)
	IsAllowedShapeError(shape objectivec.IObject) (bool, error)
	IsAllowedValueError(value objectivec.IObject) (bool, error)
	Shape() objectivec.IObject
	InitWithShapeDataTypeShapeConstraintDefaultOptionalValue(shape objectivec.IObject, type_ int64, constraint objectivec.IObject, value objectivec.IObject) MLStateConstraint
}

// Init initializes the instance.
func (m MLStateConstraint) Init() MLStateConstraint {
	rv := objc.SendIfResponds[MLStateConstraint](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLStateConstraint) Autorelease() MLStateConstraint {
	rv := objc.SendIfResponds[MLStateConstraint](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLStateConstraint creates a new MLStateConstraint instance.
func NewMLStateConstraint() MLStateConstraint {
	class := getMLStateConstraintClass()
	rv := objc.SendIfResponds[MLStateConstraint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewStateConstraintWithShapeDataTypeShapeConstraintDefaultOptionalValue(shape objectivec.IObject, type_ int64, constraint objectivec.IObject, value objectivec.IObject) MLStateConstraint {
	instance := getMLStateConstraintClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithShape:dataType:shapeConstraint:defaultOptionalValue:"), shape, type_, constraint, value)
	return MLStateConstraintFromID(rv)
}

func (m MLStateConstraint) IsAllowedDataTypeError(type_ int64) (bool, error) {
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
func (m MLStateConstraint) IsAllowedShapeError(shape objectivec.IObject) (bool, error) {
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
func (m MLStateConstraint) IsAllowedValueError(value objectivec.IObject) (bool, error) {
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
func (m MLStateConstraint) Shape() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("shape"))
	return objectivec.Object{ID: rv}
}
func (m MLStateConstraint) InitWithShapeDataTypeShapeConstraintDefaultOptionalValue(shape objectivec.IObject, type_ int64, constraint objectivec.IObject, value objectivec.IObject) MLStateConstraint {
	rv := objc.SendIfResponds[MLStateConstraint](m.ID, objc.Sel("initWithShape:dataType:shapeConstraint:defaultOptionalValue:"), shape, type_, constraint, value)
	return rv
}

func (_MLStateConstraintClass MLStateConstraintClass) ConstraintWithBufferShapeDataType(shape objectivec.IObject, type_ int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLStateConstraintClass.class), objc.Sel("constraintWithBufferShape:dataType:"), shape, type_)
	return objectivec.Object{ID: rv}
}
func (_MLStateConstraintClass MLStateConstraintClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLStateConstraintClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
