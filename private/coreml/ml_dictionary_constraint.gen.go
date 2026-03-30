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
	rv := objc.Send[MLDictionaryConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLDictionaryConstraint.IsAllowedValueError]
//   - [MLDictionaryConstraint.InitWithCoder]
//   - [MLDictionaryConstraint.InitWithKeyType]
//
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryConstraint
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
//   - [IMLDictionaryConstraint.InitWithCoder]
//   - [IMLDictionaryConstraint.InitWithKeyType]
//
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryConstraint
type IMLDictionaryConstraint interface {
	objectivec.IObject

	// Topic: Methods

	IsAllowedValueError(value objectivec.IObject) (bool, error)
	InitWithCoder(coder foundation.INSCoder) MLDictionaryConstraint
	InitWithKeyType(type_ int64) MLDictionaryConstraint
}

// Init initializes the instance.
func (d MLDictionaryConstraint) Init() MLDictionaryConstraint {
	rv := objc.Send[MLDictionaryConstraint](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d MLDictionaryConstraint) Autorelease() MLDictionaryConstraint {
	rv := objc.Send[MLDictionaryConstraint](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLDictionaryConstraint creates a new MLDictionaryConstraint instance.
func NewMLDictionaryConstraint() MLDictionaryConstraint {
	class := getMLDictionaryConstraintClass()
	rv := objc.Send[MLDictionaryConstraint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLDictionaryConstraint/initWithCoder:
func NewDictionaryConstraintWithCoder(coder objectivec.IObject) MLDictionaryConstraint {
	instance := getMLDictionaryConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLDictionaryConstraintFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLDictionaryConstraint/initWithKeyType:
func NewDictionaryConstraintWithKeyType(type_ int64) MLDictionaryConstraint {
	instance := getMLDictionaryConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKeyType:"), type_)
	return MLDictionaryConstraintFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLDictionaryConstraint/isAllowedValue:error:
func (d MLDictionaryConstraint) IsAllowedValueError(value objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("isAllowedValue:error:"), value, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("isAllowedValue:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLDictionaryConstraint/initWithCoder:
func (d MLDictionaryConstraint) InitWithCoder(coder foundation.INSCoder) MLDictionaryConstraint {
	rv := objc.Send[MLDictionaryConstraint](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLDictionaryConstraint/initWithKeyType:
func (d MLDictionaryConstraint) InitWithKeyType(type_ int64) MLDictionaryConstraint {
	rv := objc.Send[MLDictionaryConstraint](d.ID, objc.Sel("initWithKeyType:"), type_)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLDictionaryConstraint/constraintWithInt64Keys
func (_MLDictionaryConstraintClass MLDictionaryConstraintClass) ConstraintWithInt64Keys() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLDictionaryConstraintClass.class), objc.Sel("constraintWithInt64Keys"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLDictionaryConstraint/constraintWithStringKeys
func (_MLDictionaryConstraintClass MLDictionaryConstraintClass) ConstraintWithStringKeys() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLDictionaryConstraintClass.class), objc.Sel("constraintWithStringKeys"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLDictionaryConstraint/supportsSecureCoding
func (_MLDictionaryConstraintClass MLDictionaryConstraintClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLDictionaryConstraintClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
