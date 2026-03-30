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

// The class instance for the [MLSequenceConstraint] class.
var (
	_MLSequenceConstraintClass     MLSequenceConstraintClass
	_MLSequenceConstraintClassOnce sync.Once
)

func getMLSequenceConstraintClass() MLSequenceConstraintClass {
	_MLSequenceConstraintClassOnce.Do(func() {
		_MLSequenceConstraintClass = MLSequenceConstraintClass{class: objc.GetClass("MLSequenceConstraint")}
	})
	return _MLSequenceConstraintClass
}

// GetMLSequenceConstraintClass returns the class object for MLSequenceConstraint.
func GetMLSequenceConstraintClass() MLSequenceConstraintClass {
	return getMLSequenceConstraintClass()
}

type MLSequenceConstraintClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSequenceConstraintClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSequenceConstraintClass) Alloc() MLSequenceConstraint {
	rv := objc.Send[MLSequenceConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLSequenceConstraint.IsAllowedValueError]
//   - [MLSequenceConstraint.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreML/MLSequenceConstraint
type MLSequenceConstraint struct {
	objectivec.Object
}

// MLSequenceConstraintFromID constructs a [MLSequenceConstraint] from an objc.ID.
func MLSequenceConstraintFromID(id objc.ID) MLSequenceConstraint {
	return MLSequenceConstraint{objectivec.Object{ID: id}}
}

// Ensure MLSequenceConstraint implements IMLSequenceConstraint.
var _ IMLSequenceConstraint = MLSequenceConstraint{}

// An interface definition for the [MLSequenceConstraint] class.
//
// # Methods
//
//   - [IMLSequenceConstraint.IsAllowedValueError]
//   - [IMLSequenceConstraint.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreML/MLSequenceConstraint
type IMLSequenceConstraint interface {
	objectivec.IObject

	// Topic: Methods

	IsAllowedValueError(value objectivec.IObject) (bool, error)
	InitWithCoder(coder foundation.INSCoder) MLSequenceConstraint
}

// Init initializes the instance.
func (s MLSequenceConstraint) Init() MLSequenceConstraint {
	rv := objc.Send[MLSequenceConstraint](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MLSequenceConstraint) Autorelease() MLSequenceConstraint {
	rv := objc.Send[MLSequenceConstraint](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSequenceConstraint creates a new MLSequenceConstraint instance.
func NewMLSequenceConstraint() MLSequenceConstraint {
	class := getMLSequenceConstraintClass()
	rv := objc.Send[MLSequenceConstraint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSequenceConstraint/initWithCoder:
func NewSequenceConstraintWithCoder(coder objectivec.IObject) MLSequenceConstraint {
	instance := getMLSequenceConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLSequenceConstraintFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLSequenceConstraint/isAllowedValue:error:
func (s MLSequenceConstraint) IsAllowedValueError(value objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("isAllowedValue:error:"), value, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("isAllowedValue:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLSequenceConstraint/initWithCoder:
func (s MLSequenceConstraint) InitWithCoder(coder foundation.INSCoder) MLSequenceConstraint {
	rv := objc.Send[MLSequenceConstraint](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSequenceConstraint/supportsSecureCoding
func (_MLSequenceConstraintClass MLSequenceConstraintClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLSequenceConstraintClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
