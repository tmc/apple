// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLMultiArrayShapeConstraint] class.
var (
	_MLMultiArrayShapeConstraintClass     MLMultiArrayShapeConstraintClass
	_MLMultiArrayShapeConstraintClassOnce sync.Once
)

func getMLMultiArrayShapeConstraintClass() MLMultiArrayShapeConstraintClass {
	_MLMultiArrayShapeConstraintClassOnce.Do(func() {
		_MLMultiArrayShapeConstraintClass = MLMultiArrayShapeConstraintClass{class: objc.GetClass("MLMultiArrayShapeConstraint")}
	})
	return _MLMultiArrayShapeConstraintClass
}

// GetMLMultiArrayShapeConstraintClass returns the class object for MLMultiArrayShapeConstraint.
func GetMLMultiArrayShapeConstraintClass() MLMultiArrayShapeConstraintClass {
	return getMLMultiArrayShapeConstraintClass()
}

type MLMultiArrayShapeConstraintClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLMultiArrayShapeConstraintClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLMultiArrayShapeConstraintClass) Alloc() MLMultiArrayShapeConstraint {
	rv := objc.Send[MLMultiArrayShapeConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The lists of shapes or ranges of shapes that constrain a multiarray
// feature.
//
// # Accessing the Constraints
//
//   - [MLMultiArrayShapeConstraint.EnumeratedShapes]: Array of allowed shapes for a multiarray feature.
//   - [MLMultiArrayShapeConstraint.SizeRangeForDimension]: The allowable range for a dimention of the multiarray.
//   - [MLMultiArrayShapeConstraint.Type]: The type of the shape constraint.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint
type MLMultiArrayShapeConstraint struct {
	objectivec.Object
}

// MLMultiArrayShapeConstraintFromID constructs a [MLMultiArrayShapeConstraint] from an objc.ID.
//
// The lists of shapes or ranges of shapes that constrain a multiarray
// feature.
func MLMultiArrayShapeConstraintFromID(id objc.ID) MLMultiArrayShapeConstraint {
	return MLMultiArrayShapeConstraint{objectivec.Object{ID: id}}
}

// NOTE: MLMultiArrayShapeConstraint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLMultiArrayShapeConstraint] class.
//
// # Accessing the Constraints
//
//   - [IMLMultiArrayShapeConstraint.EnumeratedShapes]: Array of allowed shapes for a multiarray feature.
//   - [IMLMultiArrayShapeConstraint.SizeRangeForDimension]: The allowable range for a dimention of the multiarray.
//   - [IMLMultiArrayShapeConstraint.Type]: The type of the shape constraint.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint
type IMLMultiArrayShapeConstraint interface {
	objectivec.IObject

	// Topic: Accessing the Constraints

	// Array of allowed shapes for a multiarray feature.
	EnumeratedShapes() []foundation.NSArray
	// The allowable range for a dimention of the multiarray.
	SizeRangeForDimension() []foundation.NSValue
	// The type of the shape constraint.
	Type() MLMultiArrayShapeConstraintType

	// The type for the multi array.
	DataType() MLMultiArrayDataType
	SetDataType(value MLMultiArrayDataType)
	// The shape of the multi array.
	Shape() foundation.NSNumber
	SetShape(value foundation.NSNumber)
	// The constraint on the shape of the multiarray.
	ShapeConstraint() IMLMultiArrayShapeConstraint
	SetShapeConstraint(value IMLMultiArrayShapeConstraint)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (m MLMultiArrayShapeConstraint) Init() MLMultiArrayShapeConstraint {
	rv := objc.Send[MLMultiArrayShapeConstraint](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLMultiArrayShapeConstraint) Autorelease() MLMultiArrayShapeConstraint {
	rv := objc.Send[MLMultiArrayShapeConstraint](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLMultiArrayShapeConstraint creates a new MLMultiArrayShapeConstraint instance.
func NewMLMultiArrayShapeConstraint() MLMultiArrayShapeConstraint {
	class := getMLMultiArrayShapeConstraintClass()
	rv := objc.Send[MLMultiArrayShapeConstraint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m MLMultiArrayShapeConstraint) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Array of allowed shapes for a multiarray feature.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/enumeratedShapes
func (m MLMultiArrayShapeConstraint) EnumeratedShapes() []foundation.NSArray {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("enumeratedShapes"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSArray {
		return foundation.NSArrayFromID(id)
	})
}

// The allowable range for a dimention of the multiarray.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/sizeRangeForDimension
func (m MLMultiArrayShapeConstraint) SizeRangeForDimension() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("sizeRangeForDimension"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}

// The type of the shape constraint.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/type
func (m MLMultiArrayShapeConstraint) Type() MLMultiArrayShapeConstraintType {
	rv := objc.Send[MLMultiArrayShapeConstraintType](m.ID, objc.Sel("type"))
	return MLMultiArrayShapeConstraintType(rv)
}

// The type for the multi array.
//
// See: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/datatype
func (m MLMultiArrayShapeConstraint) DataType() MLMultiArrayDataType {
	rv := objc.Send[MLMultiArrayDataType](m.ID, objc.Sel("dataType"))
	return MLMultiArrayDataType(rv)
}
func (m MLMultiArrayShapeConstraint) SetDataType(value MLMultiArrayDataType) {
	objc.Send[struct{}](m.ID, objc.Sel("setDataType:"), value)
}

// The shape of the multi array.
//
// See: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/shape
func (m MLMultiArrayShapeConstraint) Shape() foundation.NSNumber {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("shape"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m MLMultiArrayShapeConstraint) SetShape(value foundation.NSNumber) {
	objc.Send[struct{}](m.ID, objc.Sel("setShape:"), value)
}

// The constraint on the shape of the multiarray.
//
// See: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/shapeconstraint
func (m MLMultiArrayShapeConstraint) ShapeConstraint() IMLMultiArrayShapeConstraint {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("shapeConstraint"))
	return MLMultiArrayShapeConstraintFromID(objc.ID(rv))
}
func (m MLMultiArrayShapeConstraint) SetShapeConstraint(value IMLMultiArrayShapeConstraint) {
	objc.Send[struct{}](m.ID, objc.Sel("setShapeConstraint:"), value)
}
