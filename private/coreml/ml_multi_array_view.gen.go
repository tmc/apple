// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLMultiArrayView] class.
var (
	_MLMultiArrayViewClass     MLMultiArrayViewClass
	_MLMultiArrayViewClassOnce sync.Once
)

func getMLMultiArrayViewClass() MLMultiArrayViewClass {
	_MLMultiArrayViewClassOnce.Do(func() {
		_MLMultiArrayViewClass = MLMultiArrayViewClass{class: objc.GetClass("MLMultiArrayView")}
	})
	return _MLMultiArrayViewClass
}

// GetMLMultiArrayViewClass returns the class object for MLMultiArrayView.
func GetMLMultiArrayViewClass() MLMultiArrayViewClass {
	return getMLMultiArrayViewClass()
}

type MLMultiArrayViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLMultiArrayViewClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLMultiArrayViewClass) Alloc() MLMultiArrayView {
	rv := objc.Send[MLMultiArrayView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLMultiArrayView.Parent]
//   - [MLMultiArrayView.InitExpandingDimensionsOfMultiArrayAxis]
//   - [MLMultiArrayView.InitSlicingMultiArrayOriginShapeSqueezeError]
//   - [MLMultiArrayView.InitSqueezingMultiArrayDimensionsError]
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayView
type MLMultiArrayView struct {
	MLMultiArray
}

// MLMultiArrayViewFromID constructs a [MLMultiArrayView] from an objc.ID.
func MLMultiArrayViewFromID(id objc.ID) MLMultiArrayView {
	return MLMultiArrayView{MLMultiArray: MLMultiArrayFromID(id)}
}

// Ensure MLMultiArrayView implements IMLMultiArrayView.
var _ IMLMultiArrayView = MLMultiArrayView{}

// An interface definition for the [MLMultiArrayView] class.
//
// # Methods
//
//   - [IMLMultiArrayView.Parent]
//   - [IMLMultiArrayView.InitExpandingDimensionsOfMultiArrayAxis]
//   - [IMLMultiArrayView.InitSlicingMultiArrayOriginShapeSqueezeError]
//   - [IMLMultiArrayView.InitSqueezingMultiArrayDimensionsError]
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayView
type IMLMultiArrayView interface {
	IMLMultiArray

	// Topic: Methods

	Parent() IMLMultiArray
	InitExpandingDimensionsOfMultiArrayAxis(array objectivec.IObject, axis int64) MLMultiArrayView
	InitSlicingMultiArrayOriginShapeSqueezeError(array objectivec.IObject, origin objectivec.IObject, shape objectivec.IObject, squeeze bool) (MLMultiArrayView, error)
	InitSqueezingMultiArrayDimensionsError(array objectivec.IObject, dimensions objectivec.IObject) (MLMultiArrayView, error)
}

// Init initializes the instance.
func (m MLMultiArrayView) Init() MLMultiArrayView {
	rv := objc.Send[MLMultiArrayView](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLMultiArrayView) Autorelease() MLMultiArrayView {
	rv := objc.Send[MLMultiArrayView](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLMultiArrayView creates a new MLMultiArrayView instance.
func NewMLMultiArrayView() MLMultiArrayView {
	class := getMLMultiArrayViewClass()
	rv := objc.Send[MLMultiArrayView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayView/initExpandingDimensionsOfMultiArray:axis:
func NewMultiArrayViewExpandingDimensionsOfMultiArrayAxis(array objectivec.IObject, axis int64) MLMultiArrayView {
	instance := getMLMultiArrayViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initExpandingDimensionsOfMultiArray:axis:"), array, axis)
	return MLMultiArrayViewFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayView/initSlicingMultiArray:origin:shape:squeeze:error:
func NewMultiArrayViewSlicingMultiArrayOriginShapeSqueezeError(array objectivec.IObject, origin objectivec.IObject, shape objectivec.IObject, squeeze bool) (MLMultiArrayView, error) {
	var errorPtr objc.ID
	instance := getMLMultiArrayViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initSlicingMultiArray:origin:shape:squeeze:error:"), array, origin, shape, squeeze, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLMultiArrayView{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLMultiArrayViewFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayView/initSqueezingMultiArray:dimensions:error:
func NewMultiArrayViewSqueezingMultiArrayDimensionsError(array objectivec.IObject, dimensions objectivec.IObject) (MLMultiArrayView, error) {
	var errorPtr objc.ID
	instance := getMLMultiArrayViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initSqueezingMultiArray:dimensions:error:"), array, dimensions, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLMultiArrayView{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLMultiArrayViewFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithArray:dataType:
func NewMultiArrayViewWithArrayDataType(array objectivec.IObject, type_ int64) MLMultiArrayView {
	instance := getMLMultiArrayViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArray:dataType:"), array, type_)
	return MLMultiArrayViewFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithCoder:
func NewMultiArrayViewWithCoder(coder objectivec.IObject) MLMultiArrayView {
	instance := getMLMultiArrayViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLMultiArrayViewFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithScalars:shape:dataType:
func NewMultiArrayViewWithScalarsShapeDataType(scalars objectivec.IObject, shape objectivec.IObject, type_ int64) MLMultiArrayView {
	instance := getMLMultiArrayViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithScalars:shape:dataType:"), scalars, shape, type_)
	return MLMultiArrayViewFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithShape:dataType:storageOrder:bufferAlignment:
func NewMultiArrayViewWithShapeDataTypeStorageOrderBufferAlignment(shape objectivec.IObject, type_ int64, order int64, alignment uint64) MLMultiArrayView {
	instance := getMLMultiArrayViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithShape:dataType:storageOrder:bufferAlignment:"), shape, type_, order, alignment)
	return MLMultiArrayViewFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithShape:dataType:storageOrder:error:
func NewMultiArrayViewWithShapeDataTypeStorageOrderError(shape objectivec.IObject, type_ int64, order int64) (MLMultiArrayView, error) {
	var errorPtr objc.ID
	instance := getMLMultiArrayViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithShape:dataType:storageOrder:error:"), shape, type_, order, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLMultiArrayView{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLMultiArrayViewFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayView/initExpandingDimensionsOfMultiArray:axis:
func (m MLMultiArrayView) InitExpandingDimensionsOfMultiArrayAxis(array objectivec.IObject, axis int64) MLMultiArrayView {
	rv := objc.Send[MLMultiArrayView](m.ID, objc.Sel("initExpandingDimensionsOfMultiArray:axis:"), array, axis)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayView/initSlicingMultiArray:origin:shape:squeeze:error:
func (m MLMultiArrayView) InitSlicingMultiArrayOriginShapeSqueezeError(array objectivec.IObject, origin objectivec.IObject, shape objectivec.IObject, squeeze bool) (MLMultiArrayView, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initSlicingMultiArray:origin:shape:squeeze:error:"), array, origin, shape, squeeze, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLMultiArrayView{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLMultiArrayViewFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayView/initSqueezingMultiArray:dimensions:error:
func (m MLMultiArrayView) InitSqueezingMultiArrayDimensionsError(array objectivec.IObject, dimensions objectivec.IObject) (MLMultiArrayView, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initSqueezingMultiArray:dimensions:error:"), array, dimensions, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLMultiArrayView{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLMultiArrayViewFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayView/isSqueezableShape:
func (_MLMultiArrayViewClass MLMultiArrayViewClass) IsSqueezableShape(shape objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_MLMultiArrayViewClass.class), objc.Sel("isSqueezableShape:"), shape)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayView/isSqueezableShape:dimensions:
func (_MLMultiArrayViewClass MLMultiArrayViewClass) IsSqueezableShapeDimensions(shape objectivec.IObject, dimensions objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_MLMultiArrayViewClass.class), objc.Sel("isSqueezableShape:dimensions:"), shape, dimensions)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayView/squeezeShape:strides:resultingShape:resultingStrides:
func (_MLMultiArrayViewClass MLMultiArrayViewClass) SqueezeShapeStridesResultingShapeResultingStrides(shape objectivec.IObject, strides objectivec.IObject, shape2 []objectivec.IObject, strides2 []objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_MLMultiArrayViewClass.class), objc.Sel("squeezeShape:strides:resultingShape:resultingStrides:"), shape, strides, objectivec.IObjectSliceToNSArray(shape2), objectivec.IObjectSliceToNSArray(strides2))
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayView/parent
func (m MLMultiArrayView) Parent() IMLMultiArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parent"))
	return MLMultiArrayFromID(objc.ID(rv))
}
