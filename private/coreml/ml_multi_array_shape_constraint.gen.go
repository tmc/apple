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

// # Methods
//
//   - [MLMultiArrayShapeConstraint.FindAvailableShape]
//   - [MLMultiArrayShapeConstraint.IsAllowedShapeError]
//   - [MLMultiArrayShapeConstraint.ShapeSet]
//   - [MLMultiArrayShapeConstraint.InitUnspecified]
//   - [MLMultiArrayShapeConstraint.InitWithCoder]
//   - [MLMultiArrayShapeConstraint.InitWithEnumeratedShapes]
//   - [MLMultiArrayShapeConstraint.InitWithSizeRangeForDimension]
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint
type MLMultiArrayShapeConstraint struct {
	objectivec.Object
}

// MLMultiArrayShapeConstraintFromID constructs a [MLMultiArrayShapeConstraint] from an objc.ID.
func MLMultiArrayShapeConstraintFromID(id objc.ID) MLMultiArrayShapeConstraint {
	return MLMultiArrayShapeConstraint{objectivec.Object{ID: id}}
}

// Ensure MLMultiArrayShapeConstraint implements IMLMultiArrayShapeConstraint.
var _ IMLMultiArrayShapeConstraint = MLMultiArrayShapeConstraint{}

// An interface definition for the [MLMultiArrayShapeConstraint] class.
//
// # Methods
//
//   - [IMLMultiArrayShapeConstraint.FindAvailableShape]
//   - [IMLMultiArrayShapeConstraint.IsAllowedShapeError]
//   - [IMLMultiArrayShapeConstraint.ShapeSet]
//   - [IMLMultiArrayShapeConstraint.InitUnspecified]
//   - [IMLMultiArrayShapeConstraint.InitWithCoder]
//   - [IMLMultiArrayShapeConstraint.InitWithEnumeratedShapes]
//   - [IMLMultiArrayShapeConstraint.InitWithSizeRangeForDimension]
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint
type IMLMultiArrayShapeConstraint interface {
	objectivec.IObject

	// Topic: Methods

	FindAvailableShape(shape objectivec.IObject) objectivec.IObject
	IsAllowedShapeError(shape objectivec.IObject) (bool, error)
	ShapeSet() foundation.INSOrderedSet
	InitUnspecified() MLMultiArrayShapeConstraint
	InitWithCoder(coder foundation.INSCoder) MLMultiArrayShapeConstraint
	InitWithEnumeratedShapes(shapes objectivec.IObject) MLMultiArrayShapeConstraint
	InitWithSizeRangeForDimension(dimension objectivec.IObject) MLMultiArrayShapeConstraint
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

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/initUnspecified
func NewMultiArrayShapeConstraintUnspecified() MLMultiArrayShapeConstraint {
	instance := getMLMultiArrayShapeConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initUnspecified"))
	return MLMultiArrayShapeConstraintFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/initWithCoder:
func NewMultiArrayShapeConstraintWithCoder(coder objectivec.IObject) MLMultiArrayShapeConstraint {
	instance := getMLMultiArrayShapeConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLMultiArrayShapeConstraintFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/initWithEnumeratedShapes:
func NewMultiArrayShapeConstraintWithEnumeratedShapes(shapes objectivec.IObject) MLMultiArrayShapeConstraint {
	instance := getMLMultiArrayShapeConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEnumeratedShapes:"), shapes)
	return MLMultiArrayShapeConstraintFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/initWithSizeRangeForDimension:
func NewMultiArrayShapeConstraintWithSizeRangeForDimension(dimension objectivec.IObject) MLMultiArrayShapeConstraint {
	instance := getMLMultiArrayShapeConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSizeRangeForDimension:"), dimension)
	return MLMultiArrayShapeConstraintFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/findAvailableShape:
func (m MLMultiArrayShapeConstraint) FindAvailableShape(shape objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("findAvailableShape:"), shape)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/isAllowedShape:error:
func (m MLMultiArrayShapeConstraint) IsAllowedShapeError(shape objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("isAllowedShape:error:"), shape, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("isAllowedShape:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/initUnspecified
func (m MLMultiArrayShapeConstraint) InitUnspecified() MLMultiArrayShapeConstraint {
	rv := objc.Send[MLMultiArrayShapeConstraint](m.ID, objc.Sel("initUnspecified"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/initWithCoder:
func (m MLMultiArrayShapeConstraint) InitWithCoder(coder foundation.INSCoder) MLMultiArrayShapeConstraint {
	rv := objc.Send[MLMultiArrayShapeConstraint](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/initWithEnumeratedShapes:
func (m MLMultiArrayShapeConstraint) InitWithEnumeratedShapes(shapes objectivec.IObject) MLMultiArrayShapeConstraint {
	rv := objc.Send[MLMultiArrayShapeConstraint](m.ID, objc.Sel("initWithEnumeratedShapes:"), shapes)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/initWithSizeRangeForDimension:
func (m MLMultiArrayShapeConstraint) InitWithSizeRangeForDimension(dimension objectivec.IObject) MLMultiArrayShapeConstraint {
	rv := objc.Send[MLMultiArrayShapeConstraint](m.ID, objc.Sel("initWithSizeRangeForDimension:"), dimension)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/supportsSecureCoding
func (_MLMultiArrayShapeConstraintClass MLMultiArrayShapeConstraintClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLMultiArrayShapeConstraintClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/shapeSet
func (m MLMultiArrayShapeConstraint) ShapeSet() foundation.INSOrderedSet {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("shapeSet"))
	return foundation.NSOrderedSetFromID(objc.ID(rv))
}
