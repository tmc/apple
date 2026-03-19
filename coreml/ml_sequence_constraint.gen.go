// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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

	// The constraint for a dictionary feature.
	DictionaryConstraint() IMLDictionaryConstraint
	SetDictionaryConstraint(value IMLDictionaryConstraint)
	// The size and format constraints for an image feature.
	ImageConstraint() IMLImageConstraint
	SetImageConstraint(value IMLImageConstraint)
	// The constraints on a multidimensional array feature.
	MultiArrayConstraint() IMLMultiArrayConstraint
	SetMultiArrayConstraint(value IMLMultiArrayConstraint)
	// The constraints for a sequence feature.
	SequenceConstraint() IMLSequenceConstraint
	SetSequenceConstraint(value IMLSequenceConstraint)
	// The state feature value constraint.
	StateConstraint() IMLStateConstraint
	SetStateConstraint(value IMLStateConstraint)
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
// The constraint for a dictionary feature.
//
// See: https://developer.apple.com/documentation/coreml/mlfeaturedescription/dictionaryconstraint
func (s MLSequenceConstraint) DictionaryConstraint() IMLDictionaryConstraint {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("dictionaryConstraint"))
	return MLDictionaryConstraintFromID(objc.ID(rv))
}
func (s MLSequenceConstraint) SetDictionaryConstraint(value IMLDictionaryConstraint) {
	objc.Send[struct{}](s.ID, objc.Sel("setDictionaryConstraint:"), value)
}
// The size and format constraints for an image feature.
//
// See: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (s MLSequenceConstraint) ImageConstraint() IMLImageConstraint {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("imageConstraint"))
	return MLImageConstraintFromID(objc.ID(rv))
}
func (s MLSequenceConstraint) SetImageConstraint(value IMLImageConstraint) {
	objc.Send[struct{}](s.ID, objc.Sel("setImageConstraint:"), value)
}
// The constraints on a multidimensional array feature.
//
// See: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint
func (s MLSequenceConstraint) MultiArrayConstraint() IMLMultiArrayConstraint {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("multiArrayConstraint"))
	return MLMultiArrayConstraintFromID(objc.ID(rv))
}
func (s MLSequenceConstraint) SetMultiArrayConstraint(value IMLMultiArrayConstraint) {
	objc.Send[struct{}](s.ID, objc.Sel("setMultiArrayConstraint:"), value)
}
// The constraints for a sequence feature.
//
// See: https://developer.apple.com/documentation/coreml/mlfeaturedescription/sequenceconstraint
func (s MLSequenceConstraint) SequenceConstraint() IMLSequenceConstraint {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("sequenceConstraint"))
	return MLSequenceConstraintFromID(objc.ID(rv))
}
func (s MLSequenceConstraint) SetSequenceConstraint(value IMLSequenceConstraint) {
	objc.Send[struct{}](s.ID, objc.Sel("setSequenceConstraint:"), value)
}
// The state feature value constraint.
//
// See: https://developer.apple.com/documentation/coreml/mlfeaturedescription/stateconstraint
func (s MLSequenceConstraint) StateConstraint() IMLStateConstraint {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("stateConstraint"))
	return MLStateConstraintFromID(objc.ID(rv))
}
func (s MLSequenceConstraint) SetStateConstraint(value IMLStateConstraint) {
	objc.Send[struct{}](s.ID, objc.Sel("setStateConstraint:"), value)
}

