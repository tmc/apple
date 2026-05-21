// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

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
	return MLDictionaryConstraint{objectivec.Object{ID: id}}
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

	InitWithCoder(coder foundation.INSCoder) MLDictionaryConstraint
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

// See: https://developer.apple.com/documentation/CoreML/MLDictionaryConstraint/init(coder:)
func NewDictionaryConstraintWithCoder(coder foundation.INSCoder) MLDictionaryConstraint {
	instance := getMLDictionaryConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLDictionaryConstraintFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLDictionaryConstraint/init(coder:)
func (d MLDictionaryConstraint) InitWithCoder(coder foundation.INSCoder) MLDictionaryConstraint {
	rv := objc.Send[MLDictionaryConstraint](d.ID, objc.Sel("initWithCoder:"), coder)
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
