// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

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

// The constraints for a sequence feature.
//
// # Accessing the constraints
//
//   - [MLSequenceConstraint.ValueDescription]: The description that all sequence elements must match.
//   - [MLSequenceConstraint.CountRange]: The range of values allowed for the sequence’s length.
//
// See: https://developer.apple.com/documentation/CoreML/MLSequenceConstraint
type MLSequenceConstraint struct {
	objectivec.Object
}

// MLSequenceConstraintFromID constructs a [MLSequenceConstraint] from an objc.ID.
//
// The constraints for a sequence feature.
func MLSequenceConstraintFromID(id objc.ID) MLSequenceConstraint {
	return MLSequenceConstraint{objectivec.Object{ID: id}}
}

// NOTE: MLSequenceConstraint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLSequenceConstraint] class.
//
// # Accessing the constraints
//
//   - [IMLSequenceConstraint.ValueDescription]: The description that all sequence elements must match.
//   - [IMLSequenceConstraint.CountRange]: The range of values allowed for the sequence’s length.
//
// See: https://developer.apple.com/documentation/CoreML/MLSequenceConstraint
type IMLSequenceConstraint interface {
	objectivec.IObject

	// Topic: Accessing the constraints

	// The description that all sequence elements must match.
	ValueDescription() IMLFeatureDescription
	// The range of values allowed for the sequence’s length.
	CountRange() foundation.NSRange

	InitWithCoder(coder foundation.INSCoder) MLSequenceConstraint
	EncodeWithCoder(coder foundation.INSCoder)
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

// See: https://developer.apple.com/documentation/CoreML/MLSequenceConstraint/init(coder:)
func NewSequenceConstraintWithCoder(coder foundation.INSCoder) MLSequenceConstraint {
	instance := getMLSequenceConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLSequenceConstraintFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLSequenceConstraint/init(coder:)
func (s MLSequenceConstraint) InitWithCoder(coder foundation.INSCoder) MLSequenceConstraint {
	rv := objc.Send[MLSequenceConstraint](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (s MLSequenceConstraint) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The description that all sequence elements must match.
//
// See: https://developer.apple.com/documentation/CoreML/MLSequenceConstraint/valueDescription
func (s MLSequenceConstraint) ValueDescription() IMLFeatureDescription {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("valueDescription"))
	return MLFeatureDescriptionFromID(objc.ID(rv))
}

// The range of values allowed for the sequence’s length.
//
// See: https://developer.apple.com/documentation/CoreML/MLSequenceConstraint/countRange
func (s MLSequenceConstraint) CountRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](s.ID, objc.Sel("countRange"))
	return foundation.NSRange(rv)
}
