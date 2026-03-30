// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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

// The value limitations of a number.
//
// # Numeric Constraints
//
//   - [MLNumericConstraint.MinNumber]: The smallest numerical value allowed by this constraint.
//   - [MLNumericConstraint.MaxNumber]: The largest numerical value allowed by this constraint.
//   - [MLNumericConstraint.EnumeratedNumbers]: A set of the numbers allowed in this constraint.
//
// See: https://developer.apple.com/documentation/CoreML/MLNumericConstraint
type MLNumericConstraint struct {
	objectivec.Object
}

// MLNumericConstraintFromID constructs a [MLNumericConstraint] from an objc.ID.
//
// The value limitations of a number.
func MLNumericConstraintFromID(id objc.ID) MLNumericConstraint {
	return MLNumericConstraint{objectivec.Object{ID: id}}
}

// NOTE: MLNumericConstraint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLNumericConstraint] class.
//
// # Numeric Constraints
//
//   - [IMLNumericConstraint.MinNumber]: The smallest numerical value allowed by this constraint.
//   - [IMLNumericConstraint.MaxNumber]: The largest numerical value allowed by this constraint.
//   - [IMLNumericConstraint.EnumeratedNumbers]: A set of the numbers allowed in this constraint.
//
// See: https://developer.apple.com/documentation/CoreML/MLNumericConstraint
type IMLNumericConstraint interface {
	objectivec.IObject

	// Topic: Numeric Constraints

	// The smallest numerical value allowed by this constraint.
	MinNumber() foundation.NSNumber
	// The largest numerical value allowed by this constraint.
	MaxNumber() foundation.NSNumber
	// A set of the numbers allowed in this constraint.
	EnumeratedNumbers() foundation.INSSet

	// The constraints of this paramter description value, if and only if the value is numerical.
	NumericConstraint() IMLNumericConstraint
	SetNumericConstraint(value IMLNumericConstraint)
	EncodeWithCoder(coder foundation.INSCoder)
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

func (n MLNumericConstraint) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The smallest numerical value allowed by this constraint.
//
// See: https://developer.apple.com/documentation/CoreML/MLNumericConstraint/minNumber
func (n MLNumericConstraint) MinNumber() foundation.NSNumber {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("minNumber"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// The largest numerical value allowed by this constraint.
//
// See: https://developer.apple.com/documentation/CoreML/MLNumericConstraint/maxNumber
func (n MLNumericConstraint) MaxNumber() foundation.NSNumber {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("maxNumber"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// A set of the numbers allowed in this constraint.
//
// See: https://developer.apple.com/documentation/CoreML/MLNumericConstraint/enumeratedNumbers
func (n MLNumericConstraint) EnumeratedNumbers() foundation.INSSet {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("enumeratedNumbers"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The constraints of this paramter description value, if and only if the
// value is numerical.
//
// See: https://developer.apple.com/documentation/coreml/mlparameterdescription/numericconstraint
func (n MLNumericConstraint) NumericConstraint() IMLNumericConstraint {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("numericConstraint"))
	return MLNumericConstraintFromID(objc.ID(rv))
}
func (n MLNumericConstraint) SetNumericConstraint(value IMLNumericConstraint) {
	objc.Send[struct{}](n.ID, objc.Sel("setNumericConstraint:"), value)
}
