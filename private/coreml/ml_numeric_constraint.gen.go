// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNumericConstraint] class.
var (
	_MLNumericConstraintClass     MLNumericConstraintClass
	_MLNumericConstraintClassOnce sync.Once
)

func getMLNumericConstraintClass() MLNumericConstraintClass {
	_MLNumericConstraintClassOnce.Do(func() {
		_MLNumericConstraintClass = MLNumericConstraintClass{class: objc.GetClass("MLNumericConstraint")}
	})
	return _MLNumericConstraintClass
}

// GetMLNumericConstraintClass returns the class object for MLNumericConstraint.
func GetMLNumericConstraintClass() MLNumericConstraintClass {
	return getMLNumericConstraintClass()
}

type MLNumericConstraintClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNumericConstraintClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNumericConstraintClass) Alloc() MLNumericConstraint {
	rv := objc.Send[MLNumericConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLNumericConstraint.InitWithCoder]
//   - [MLNumericConstraint.EnumeratedNumbers]
//   - [MLNumericConstraint.SetEnumeratedNumbers]
//   - [MLNumericConstraint.MaxNumber]
//   - [MLNumericConstraint.SetMaxNumber]
//   - [MLNumericConstraint.MinNumber]
//   - [MLNumericConstraint.SetMinNumber]
// See: https://developer.apple.com/documentation/CoreML/MLNumericConstraint
type MLNumericConstraint struct {
	objectivec.Object
}

// MLNumericConstraintFromID constructs a [MLNumericConstraint] from an objc.ID.
func MLNumericConstraintFromID(id objc.ID) MLNumericConstraint {
	return MLNumericConstraint{objectivec.Object{ID: id}}
}
// Ensure MLNumericConstraint implements IMLNumericConstraint.
var _ IMLNumericConstraint = MLNumericConstraint{}

// An interface definition for the [MLNumericConstraint] class.
//
// # Methods
//
//   - [IMLNumericConstraint.InitWithCoder]
//   - [IMLNumericConstraint.EnumeratedNumbers]
//   - [IMLNumericConstraint.SetEnumeratedNumbers]
//   - [IMLNumericConstraint.MaxNumber]
//   - [IMLNumericConstraint.SetMaxNumber]
//   - [IMLNumericConstraint.MinNumber]
//   - [IMLNumericConstraint.SetMinNumber]
//
// See: https://developer.apple.com/documentation/CoreML/MLNumericConstraint
type IMLNumericConstraint interface {
	objectivec.IObject

	// Topic: Methods

	InitWithCoder(coder foundation.INSCoder) MLNumericConstraint
	EnumeratedNumbers() foundation.INSSet
	SetEnumeratedNumbers(value foundation.INSSet)
	MaxNumber() foundation.NSNumber
	SetMaxNumber(value foundation.NSNumber)
	MinNumber() foundation.NSNumber
	SetMinNumber(value foundation.NSNumber)
}

// Init initializes the instance.
func (n MLNumericConstraint) Init() MLNumericConstraint {
	rv := objc.Send[MLNumericConstraint](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MLNumericConstraint) Autorelease() MLNumericConstraint {
	rv := objc.Send[MLNumericConstraint](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNumericConstraint creates a new MLNumericConstraint instance.
func NewMLNumericConstraint() MLNumericConstraint {
	class := getMLNumericConstraintClass()
	rv := objc.Send[MLNumericConstraint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNumericConstraint/initWithCoder:
func NewNumericConstraintWithCoder(coder objectivec.IObject) MLNumericConstraint {
	instance := getMLNumericConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLNumericConstraintFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNumericConstraint/initWithCoder:
func (n MLNumericConstraint) InitWithCoder(coder foundation.INSCoder) MLNumericConstraint {
	rv := objc.Send[MLNumericConstraint](n.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNumericConstraint/numericConstraintWithEnumeratedNumbers:
func (_MLNumericConstraintClass MLNumericConstraintClass) NumericConstraintWithEnumeratedNumbers(numbers objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNumericConstraintClass.class), objc.Sel("numericConstraintWithEnumeratedNumbers:"), numbers)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNumericConstraint/numericConstraintWithMinNumber:maxNumber:
func (_MLNumericConstraintClass MLNumericConstraintClass) NumericConstraintWithMinNumberMaxNumber(number objectivec.IObject, number2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNumericConstraintClass.class), objc.Sel("numericConstraintWithMinNumber:maxNumber:"), number, number2)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLNumericConstraint/supportsSecureCoding
func (_MLNumericConstraintClass MLNumericConstraintClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLNumericConstraintClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNumericConstraint/enumeratedNumbers
func (n MLNumericConstraint) EnumeratedNumbers() foundation.INSSet {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("enumeratedNumbers"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (n MLNumericConstraint) SetEnumeratedNumbers(value foundation.INSSet) {
	objc.Send[struct{}](n.ID, objc.Sel("setEnumeratedNumbers:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNumericConstraint/maxNumber
func (n MLNumericConstraint) MaxNumber() foundation.NSNumber {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("maxNumber"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (n MLNumericConstraint) SetMaxNumber(value foundation.NSNumber) {
	objc.Send[struct{}](n.ID, objc.Sel("setMaxNumber:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNumericConstraint/minNumber
func (n MLNumericConstraint) MinNumber() foundation.NSNumber {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("minNumber"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (n MLNumericConstraint) SetMinNumber(value foundation.NSNumber) {
	objc.Send[struct{}](n.ID, objc.Sel("setMinNumber:"), value)
}

