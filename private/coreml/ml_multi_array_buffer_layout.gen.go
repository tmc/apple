// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLMultiArrayBufferLayout] class.
var (
	_MLMultiArrayBufferLayoutClass     MLMultiArrayBufferLayoutClass
	_MLMultiArrayBufferLayoutClassOnce sync.Once
)

func getMLMultiArrayBufferLayoutClass() MLMultiArrayBufferLayoutClass {
	_MLMultiArrayBufferLayoutClassOnce.Do(func() {
		_MLMultiArrayBufferLayoutClass = MLMultiArrayBufferLayoutClass{class: objc.GetClass("MLMultiArrayBufferLayout")}
	})
	return _MLMultiArrayBufferLayoutClass
}

// GetMLMultiArrayBufferLayoutClass returns the class object for MLMultiArrayBufferLayout.
func GetMLMultiArrayBufferLayoutClass() MLMultiArrayBufferLayoutClass {
	return getMLMultiArrayBufferLayoutClass()
}

type MLMultiArrayBufferLayoutClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLMultiArrayBufferLayoutClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLMultiArrayBufferLayoutClass) Alloc() MLMultiArrayBufferLayout {
	rv := objc.Send[MLMultiArrayBufferLayout](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLMultiArrayBufferLayout.Count]
//   - [MLMultiArrayBufferLayout.IsSubspaceOfBufferLayout]
//   - [MLMultiArrayBufferLayout.OffsetOfScalarAtIndexContiguousScalars]
//   - [MLMultiArrayBufferLayout.Shape]
//   - [MLMultiArrayBufferLayout.Strides]
//   - [MLMultiArrayBufferLayout.InitWithShapeStrides]
type MLMultiArrayBufferLayout struct {
	objectivec.Object
}

// MLMultiArrayBufferLayoutFromID constructs a [MLMultiArrayBufferLayout] from an objc.ID.
func MLMultiArrayBufferLayoutFromID(id objc.ID) MLMultiArrayBufferLayout {
	return MLMultiArrayBufferLayout{objectivec.Object{ID: id}}
}

// Ensure MLMultiArrayBufferLayout implements IMLMultiArrayBufferLayout.
var _ IMLMultiArrayBufferLayout = MLMultiArrayBufferLayout{}

// An interface definition for the [MLMultiArrayBufferLayout] class.
//
// # Methods
//
//   - [IMLMultiArrayBufferLayout.Count]
//   - [IMLMultiArrayBufferLayout.IsSubspaceOfBufferLayout]
//   - [IMLMultiArrayBufferLayout.OffsetOfScalarAtIndexContiguousScalars]
//   - [IMLMultiArrayBufferLayout.Shape]
//   - [IMLMultiArrayBufferLayout.Strides]
//   - [IMLMultiArrayBufferLayout.InitWithShapeStrides]
type IMLMultiArrayBufferLayout interface {
	objectivec.IObject

	// Topic: Methods

	Count() int64
	IsSubspaceOfBufferLayout(layout objectivec.IObject) bool
	OffsetOfScalarAtIndexContiguousScalars(index int64, scalars *int64) int64
	Shape() foundation.INSArray
	Strides() foundation.INSArray
	InitWithShapeStrides(shape objectivec.IObject, strides objectivec.IObject) MLMultiArrayBufferLayout
}

// Init initializes the instance.
func (m MLMultiArrayBufferLayout) Init() MLMultiArrayBufferLayout {
	rv := objc.Send[MLMultiArrayBufferLayout](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLMultiArrayBufferLayout) Autorelease() MLMultiArrayBufferLayout {
	rv := objc.Send[MLMultiArrayBufferLayout](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLMultiArrayBufferLayout creates a new MLMultiArrayBufferLayout instance.
func NewMLMultiArrayBufferLayout() MLMultiArrayBufferLayout {
	class := getMLMultiArrayBufferLayoutClass()
	rv := objc.Send[MLMultiArrayBufferLayout](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewMultiArrayBufferLayoutWithShapeStrides(shape objectivec.IObject, strides objectivec.IObject) MLMultiArrayBufferLayout {
	instance := getMLMultiArrayBufferLayoutClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithShape:strides:"), shape, strides)
	return MLMultiArrayBufferLayoutFromID(rv)
}

func (m MLMultiArrayBufferLayout) IsSubspaceOfBufferLayout(layout objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isSubspaceOfBufferLayout:"), layout)
	return rv
}
func (m MLMultiArrayBufferLayout) OffsetOfScalarAtIndexContiguousScalars(index int64, scalars *int64) int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("offsetOfScalarAtIndex:contiguousScalars:"), index, scalars)
	return rv
}
func (m MLMultiArrayBufferLayout) InitWithShapeStrides(shape objectivec.IObject, strides objectivec.IObject) MLMultiArrayBufferLayout {
	rv := objc.Send[MLMultiArrayBufferLayout](m.ID, objc.Sel("initWithShape:strides:"), shape, strides)
	return rv
}

func (m MLMultiArrayBufferLayout) Count() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("count"))
	return rv
}
func (m MLMultiArrayBufferLayout) Shape() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("shape"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLMultiArrayBufferLayout) Strides() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("strides"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
