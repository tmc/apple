// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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

// Alloc allocates memory for a new instance of the class.
func (mc MLDictionaryConstraintClass) Alloc() MLDictionaryConstraint {
	rv := objc.Send[MLDictionaryConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// The constraint on the keys for a dictionary feature.
//
// # Accessing the constraint
//
//   - [MLDictionaryConstraint.KeyType]: The key type for the dictionary.
//
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryConstraint
type MLDictionaryConstraint struct {
	objectivec.Object
}

// MLDictionaryConstraintFromID constructs a [MLDictionaryConstraint] from an objc.ID.
//
// The constraint on the keys for a dictionary feature.
func MLDictionaryConstraintFromID(id objc.ID) MLDictionaryConstraint {
	return MLDictionaryConstraint{objectivec.Object{id}}
}
// NOTE: MLDictionaryConstraint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MLDictionaryConstraint] class.
//
// # Accessing the constraint
//
//   - [IMLDictionaryConstraint.KeyType]: The key type for the dictionary.
//
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryConstraint
type IMLDictionaryConstraint interface {
	objectivec.IObject

	// Topic: Accessing the constraint

	// The key type for the dictionary.
	KeyType() MLFeatureType

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









func (d MLDictionaryConstraint) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The key type for the dictionary.
//
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryConstraint/keyType
func (d MLDictionaryConstraint) KeyType() MLFeatureType {
	rv := objc.Send[MLFeatureType](d.ID, objc.Sel("keyType"))
	return MLFeatureType(rv)
}



// The constraint for a dictionary feature.
//
// See: https://developer.apple.com/documentation/coreml/mlfeaturedescription/dictionaryconstraint
func (d MLDictionaryConstraint) DictionaryConstraint() IMLDictionaryConstraint {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("dictionaryConstraint"))
	return MLDictionaryConstraintFromID(objc.ID(rv))
}
func (d MLDictionaryConstraint) SetDictionaryConstraint(value IMLDictionaryConstraint) {
	objc.Send[struct{}](d.ID, objc.Sel("setDictionaryConstraint:"), value)
}



// The size and format constraints for an image feature.
//
// See: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (d MLDictionaryConstraint) ImageConstraint() IMLImageConstraint {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("imageConstraint"))
	return MLImageConstraintFromID(objc.ID(rv))
}
func (d MLDictionaryConstraint) SetImageConstraint(value IMLImageConstraint) {
	objc.Send[struct{}](d.ID, objc.Sel("setImageConstraint:"), value)
}



// The constraints on a multidimensional array feature.
//
// See: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint
func (d MLDictionaryConstraint) MultiArrayConstraint() IMLMultiArrayConstraint {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("multiArrayConstraint"))
	return MLMultiArrayConstraintFromID(objc.ID(rv))
}
func (d MLDictionaryConstraint) SetMultiArrayConstraint(value IMLMultiArrayConstraint) {
	objc.Send[struct{}](d.ID, objc.Sel("setMultiArrayConstraint:"), value)
}



// The constraints for a sequence feature.
//
// See: https://developer.apple.com/documentation/coreml/mlfeaturedescription/sequenceconstraint
func (d MLDictionaryConstraint) SequenceConstraint() IMLSequenceConstraint {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("sequenceConstraint"))
	return MLSequenceConstraintFromID(objc.ID(rv))
}
func (d MLDictionaryConstraint) SetSequenceConstraint(value IMLSequenceConstraint) {
	objc.Send[struct{}](d.ID, objc.Sel("setSequenceConstraint:"), value)
}



// The state feature value constraint.
//
// See: https://developer.apple.com/documentation/coreml/mlfeaturedescription/stateconstraint
func (d MLDictionaryConstraint) StateConstraint() IMLStateConstraint {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("stateConstraint"))
	return MLStateConstraintFromID(objc.ID(rv))
}
func (d MLDictionaryConstraint) SetStateConstraint(value IMLStateConstraint) {
	objc.Send[struct{}](d.ID, objc.Sel("setStateConstraint:"), value)
}

























