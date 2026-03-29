// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLStateConstraint] class.
var (
	_MLStateConstraintClass     MLStateConstraintClass
	_MLStateConstraintClassOnce sync.Once
)

func getMLStateConstraintClass() MLStateConstraintClass {
	_MLStateConstraintClassOnce.Do(func() {
		_MLStateConstraintClass = MLStateConstraintClass{class: objc.GetClass("MLStateConstraint")}
	})
	return _MLStateConstraintClass
}

// GetMLStateConstraintClass returns the class object for MLStateConstraint.
func GetMLStateConstraintClass() MLStateConstraintClass {
	return getMLStateConstraintClass()
}

type MLStateConstraintClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLStateConstraintClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLStateConstraintClass) Alloc() MLStateConstraint {
	rv := objc.Send[MLStateConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Constraint of a state feature value.
//
// # Inspecting a state constraint
//
//   - [MLStateConstraint.BufferShape]: The shape of the state buffer.
//   - [MLStateConstraint.SetBufferShape]
//   - [MLStateConstraint.DataType]: The data type of scalars in the state buffer.
//
// See: https://developer.apple.com/documentation/CoreML/MLStateConstraint
type MLStateConstraint struct {
	objectivec.Object
}

// MLStateConstraintFromID constructs a [MLStateConstraint] from an objc.ID.
//
// Constraint of a state feature value.
func MLStateConstraintFromID(id objc.ID) MLStateConstraint {
	return MLStateConstraint{objectivec.Object{ID: id}}
}
// NOTE: MLStateConstraint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLStateConstraint] class.
//
// # Inspecting a state constraint
//
//   - [IMLStateConstraint.BufferShape]: The shape of the state buffer.
//   - [IMLStateConstraint.SetBufferShape]
//   - [IMLStateConstraint.DataType]: The data type of scalars in the state buffer.
//
// See: https://developer.apple.com/documentation/CoreML/MLStateConstraint
type IMLStateConstraint interface {
	objectivec.IObject

	// Topic: Inspecting a state constraint

	// The shape of the state buffer.
	BufferShape() int
	SetBufferShape(value int)
	// The data type of scalars in the state buffer.
	DataType() MLMultiArrayDataType

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s MLStateConstraint) Init() MLStateConstraint {
	rv := objc.Send[MLStateConstraint](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MLStateConstraint) Autorelease() MLStateConstraint {
	rv := objc.Send[MLStateConstraint](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLStateConstraint creates a new MLStateConstraint instance.
func NewMLStateConstraint() MLStateConstraint {
	class := getMLStateConstraintClass()
	rv := objc.Send[MLStateConstraint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (s MLStateConstraint) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The shape of the state buffer.
//
// See: https://developer.apple.com/documentation/coreml/mlstateconstraint/buffershape-4zb3w
func (s MLStateConstraint) BufferShape() int {
	rv := objc.Send[int](s.ID, objc.Sel("bufferShape"))
	return rv
}
func (s MLStateConstraint) SetBufferShape(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setBufferShape:"), value)
}
// The data type of scalars in the state buffer.
//
// See: https://developer.apple.com/documentation/CoreML/MLStateConstraint/dataType
func (s MLStateConstraint) DataType() MLMultiArrayDataType {
	rv := objc.Send[MLMultiArrayDataType](s.ID, objc.Sel("dataType"))
	return MLMultiArrayDataType(rv)
}

