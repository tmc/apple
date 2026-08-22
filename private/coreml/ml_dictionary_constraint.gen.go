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

// The class instance for the [MLDictionaryConstraint] class.
var (
	_MLDictionaryConstraintClass     MLDictionaryConstraintClass
	_MLDictionaryConstraintClassOnce sync.Once
)

func getMLDictionaryConstraintClass() MLDictionaryConstraintClass {
	_MLDictionaryConstraintClassOnce.Do(func() {
		_MLDictionaryConstraintClass = MLDictionaryConstraintClass{class: objc.GetClass("MLDictionaryConstraint")}
	})
	return _MLDictionaryConstraintClass
}

// GetMLDictionaryConstraintClass returns the class object for MLDictionaryConstraint.
func GetMLDictionaryConstraintClass() MLDictionaryConstraintClass {
	return getMLDictionaryConstraintClass()
}

type MLDictionaryConstraintClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLDictionaryConstraintClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLDictionaryConstraintClass) Alloc() MLDictionaryConstraint {
	rv := objc.SendIfResponds[MLDictionaryConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLDictionaryConstraint.IsAllowedValueError]
//   - [MLDictionaryConstraint.InitWithKeyType]
type MLDictionaryConstraint struct {
	objectivec.Object
}

// MLDictionaryConstraintFromID constructs a [MLDictionaryConstraint] from an objc.ID.
func MLDictionaryConstraintFromID(id objc.ID) MLDictionaryConstraint {
	return MLDictionaryConstraint{objectivec.Object{ID: id}}
}

// Ensure MLDictionaryConstraint implements IMLDictionaryConstraint.
var _ IMLDictionaryConstraint = MLDictionaryConstraint{}

// An interface definition for the [MLDictionaryConstraint] class.
//
// # Methods
//
//   - [IMLDictionaryConstraint.IsAllowedValueError]
//   - [IMLDictionaryConstraint.InitWithKeyType]
type IMLDictionaryConstraint interface {
	objectivec.IObject

	// Topic: Methods

	IsAllowedValueError(value objectivec.IObject) (bool, error)
	InitWithKeyType(type_ int64) MLDictionaryConstraint
}

// Init initializes the instance.
func (m MLDictionaryConstraint) Init() MLDictionaryConstraint {
	rv := objc.SendIfResponds[MLDictionaryConstraint](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLDictionaryConstraint) Autorelease() MLDictionaryConstraint {
	rv := objc.SendIfResponds[MLDictionaryConstraint](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLDictionaryConstraint creates a new MLDictionaryConstraint instance.
func NewMLDictionaryConstraint() MLDictionaryConstraint {
	class := getMLDictionaryConstraintClass()
	rv := objc.SendIfResponds[MLDictionaryConstraint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDictionaryConstraintWithKeyType(type_ int64) MLDictionaryConstraint {
	instance := getMLDictionaryConstraintClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithKeyType:"), type_)
	return MLDictionaryConstraintFromID(rv)
}

func (m MLDictionaryConstraint) IsAllowedValueError(value objectivec.IObject) (bool, error) {
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
func (m MLDictionaryConstraint) InitWithKeyType(type_ int64) MLDictionaryConstraint {
	rv := objc.SendIfResponds[MLDictionaryConstraint](m.ID, objc.Sel("initWithKeyType:"), type_)
	return rv
}

func (_MLDictionaryConstraintClass MLDictionaryConstraintClass) ConstraintWithInt64Keys() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLDictionaryConstraintClass.class), objc.Sel("constraintWithInt64Keys"))
	return objectivec.Object{ID: rv}
}
func (_MLDictionaryConstraintClass MLDictionaryConstraintClass) ConstraintWithStringKeys() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLDictionaryConstraintClass.class), objc.Sel("constraintWithStringKeys"))
	return objectivec.Object{ID: rv}
}
func (_MLDictionaryConstraintClass MLDictionaryConstraintClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLDictionaryConstraintClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
