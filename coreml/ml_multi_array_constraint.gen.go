// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLMultiArrayConstraint] class.
var (
	_MLMultiArrayConstraintClass     MLMultiArrayConstraintClass
	_MLMultiArrayConstraintClassOnce sync.Once
)

func getMLMultiArrayConstraintClass() MLMultiArrayConstraintClass {
	_MLMultiArrayConstraintClassOnce.Do(func() {
		_MLMultiArrayConstraintClass = MLMultiArrayConstraintClass{class: objc.GetClass("MLMultiArrayConstraint")}
	})
	return _MLMultiArrayConstraintClass
}

// GetMLMultiArrayConstraintClass returns the class object for MLMultiArrayConstraint.
func GetMLMultiArrayConstraintClass() MLMultiArrayConstraintClass {
	return getMLMultiArrayConstraintClass()
}

type MLMultiArrayConstraintClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLMultiArrayConstraintClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLMultiArrayConstraintClass) Alloc() MLMultiArrayConstraint {
	rv := objc.Send[MLMultiArrayConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The shape and data type constraints for a multidimensional array feature.
//
// # Accessing the Constraints
//
//   - [MLMultiArrayConstraint.Shape]: The shape of the multi array.
//   - [MLMultiArrayConstraint.DataType]: The type for the multi array.
//   - [MLMultiArrayConstraint.ShapeConstraint]: The constraint on the shape of the multiarray.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint
type MLMultiArrayConstraint struct {
	objectivec.Object
}

// MLMultiArrayConstraintFromID constructs a [MLMultiArrayConstraint] from an objc.ID.
//
// The shape and data type constraints for a multidimensional array feature.
func MLMultiArrayConstraintFromID(id objc.ID) MLMultiArrayConstraint {
	return MLMultiArrayConstraint{objectivec.Object{ID: id}}
}

// NOTE: MLMultiArrayConstraint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLMultiArrayConstraint] class.
//
// # Accessing the Constraints
//
//   - [IMLMultiArrayConstraint.Shape]: The shape of the multi array.
//   - [IMLMultiArrayConstraint.DataType]: The type for the multi array.
//   - [IMLMultiArrayConstraint.ShapeConstraint]: The constraint on the shape of the multiarray.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint
type IMLMultiArrayConstraint interface {
	objectivec.IObject

	// Topic: Accessing the Constraints

	// The shape of the multi array.
	Shape() []foundation.NSNumber
	// The type for the multi array.
	DataType() MLMultiArrayDataType
	// The constraint on the shape of the multiarray.
	ShapeConstraint() IMLMultiArrayShapeConstraint

	InitWithCoder(coder foundation.INSCoder) MLMultiArrayConstraint
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (m MLMultiArrayConstraint) Init() MLMultiArrayConstraint {
	rv := objc.Send[MLMultiArrayConstraint](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLMultiArrayConstraint) Autorelease() MLMultiArrayConstraint {
	rv := objc.Send[MLMultiArrayConstraint](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLMultiArrayConstraint creates a new MLMultiArrayConstraint instance.
func NewMLMultiArrayConstraint() MLMultiArrayConstraint {
	class := getMLMultiArrayConstraintClass()
	rv := objc.Send[MLMultiArrayConstraint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/init(coder:)
func NewMultiArrayConstraintWithCoder(coder foundation.INSCoder) MLMultiArrayConstraint {
	instance := getMLMultiArrayConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLMultiArrayConstraintFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/init(coder:)
func (m MLMultiArrayConstraint) InitWithCoder(coder foundation.INSCoder) MLMultiArrayConstraint {
	rv := objc.Send[MLMultiArrayConstraint](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (m MLMultiArrayConstraint) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The shape of the multi array.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/shape
func (m MLMultiArrayConstraint) Shape() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("shape"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// The type for the multi array.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/dataType
func (m MLMultiArrayConstraint) DataType() MLMultiArrayDataType {
	rv := objc.Send[MLMultiArrayDataType](m.ID, objc.Sel("dataType"))
	return MLMultiArrayDataType(rv)
}

// The constraint on the shape of the multiarray.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/shapeConstraint
func (m MLMultiArrayConstraint) ShapeConstraint() IMLMultiArrayShapeConstraint {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("shapeConstraint"))
	return MLMultiArrayShapeConstraintFromID(objc.ID(rv))
}
