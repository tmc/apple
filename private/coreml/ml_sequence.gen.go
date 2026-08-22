// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSequence] class.
var (
	_MLSequenceClass     MLSequenceClass
	_MLSequenceClassOnce sync.Once
)

func getMLSequenceClass() MLSequenceClass {
	_MLSequenceClassOnce.Do(func() {
		_MLSequenceClass = MLSequenceClass{class: objc.GetClass("MLSequence")}
	})
	return _MLSequenceClass
}

// GetMLSequenceClass returns the class object for MLSequence.
func GetMLSequenceClass() MLSequenceClass {
	return getMLSequenceClass()
}

type MLSequenceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSequenceClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSequenceClass) Alloc() MLSequence {
	rv := objc.SendIfResponds[MLSequence](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLSequence.FeatureValues]
//   - [MLSequence.Values]
//   - [MLSequence.InitWithArrayType]
type MLSequence struct {
	objectivec.Object
}

// MLSequenceFromID constructs a [MLSequence] from an objc.ID.
func MLSequenceFromID(id objc.ID) MLSequence {
	return MLSequence{objectivec.Object{ID: id}}
}

// Ensure MLSequence implements IMLSequence.
var _ IMLSequence = MLSequence{}

// An interface definition for the [MLSequence] class.
//
// # Methods
//
//   - [IMLSequence.FeatureValues]
//   - [IMLSequence.Values]
//   - [IMLSequence.InitWithArrayType]
type IMLSequence interface {
	objectivec.IObject

	// Topic: Methods

	FeatureValues() foundation.INSArray
	Values() foundation.INSArray
	InitWithArrayType(array objectivec.IObject, type_ int64) MLSequence
}

// Init initializes the instance.
func (m MLSequence) Init() MLSequence {
	rv := objc.SendIfResponds[MLSequence](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLSequence) Autorelease() MLSequence {
	rv := objc.SendIfResponds[MLSequence](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSequence creates a new MLSequence instance.
func NewMLSequence() MLSequence {
	class := getMLSequenceClass()
	rv := objc.SendIfResponds[MLSequence](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSequenceWithArrayType(array objectivec.IObject, type_ int64) MLSequence {
	instance := getMLSequenceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithArray:type:"), array, type_)
	return MLSequenceFromID(rv)
}

func (m MLSequence) InitWithArrayType(array objectivec.IObject, type_ int64) MLSequence {
	rv := objc.SendIfResponds[MLSequence](m.ID, objc.Sel("initWithArray:type:"), array, type_)
	return rv
}

func (_MLSequenceClass MLSequenceClass) EmptySequenceWithType(type_ int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLSequenceClass.class), objc.Sel("emptySequenceWithType:"), type_)
	return objectivec.Object{ID: rv}
}
func (_MLSequenceClass MLSequenceClass) SequenceFromArrayError(array objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLSequenceClass.class), objc.Sel("sequenceFromArray:error:"), array, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLSequenceClass MLSequenceClass) SequenceWithInt64Array(int64Array objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLSequenceClass.class), objc.Sel("sequenceWithInt64Array:"), int64Array)
	return objectivec.Object{ID: rv}
}
func (_MLSequenceClass MLSequenceClass) SequenceWithStringArray(array objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLSequenceClass.class), objc.Sel("sequenceWithStringArray:"), array)
	return objectivec.Object{ID: rv}
}
func (_MLSequenceClass MLSequenceClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLSequenceClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (m MLSequence) FeatureValues() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("featureValues"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSequence) Values() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("values"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
