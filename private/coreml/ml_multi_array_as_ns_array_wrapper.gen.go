// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLMultiArrayAsNSArrayWrapper] class.
var (
	_MLMultiArrayAsNSArrayWrapperClass     MLMultiArrayAsNSArrayWrapperClass
	_MLMultiArrayAsNSArrayWrapperClassOnce sync.Once
)

func getMLMultiArrayAsNSArrayWrapperClass() MLMultiArrayAsNSArrayWrapperClass {
	_MLMultiArrayAsNSArrayWrapperClassOnce.Do(func() {
		_MLMultiArrayAsNSArrayWrapperClass = MLMultiArrayAsNSArrayWrapperClass{class: objc.GetClass("MLMultiArrayAsNSArrayWrapper")}
	})
	return _MLMultiArrayAsNSArrayWrapperClass
}

// GetMLMultiArrayAsNSArrayWrapperClass returns the class object for MLMultiArrayAsNSArrayWrapper.
func GetMLMultiArrayAsNSArrayWrapperClass() MLMultiArrayAsNSArrayWrapperClass {
	return getMLMultiArrayAsNSArrayWrapperClass()
}

type MLMultiArrayAsNSArrayWrapperClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLMultiArrayAsNSArrayWrapperClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLMultiArrayAsNSArrayWrapperClass) Alloc() MLMultiArrayAsNSArrayWrapper {
	rv := objc.Send[MLMultiArrayAsNSArrayWrapper](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLMultiArrayAsNSArrayWrapper.MultiArray]
//   - [MLMultiArrayAsNSArrayWrapper.SetMultiArray]
//   - [MLMultiArrayAsNSArrayWrapper.InitWrappingMultiArray]
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayAsNSArrayWrapper
type MLMultiArrayAsNSArrayWrapper struct {
	foundation.NSArray
}

// MLMultiArrayAsNSArrayWrapperFromID constructs a [MLMultiArrayAsNSArrayWrapper] from an objc.ID.
func MLMultiArrayAsNSArrayWrapperFromID(id objc.ID) MLMultiArrayAsNSArrayWrapper {
	return MLMultiArrayAsNSArrayWrapper{NSArray: foundation.NSArrayFromID(id)}
}

// Ensure MLMultiArrayAsNSArrayWrapper implements IMLMultiArrayAsNSArrayWrapper.
var _ IMLMultiArrayAsNSArrayWrapper = MLMultiArrayAsNSArrayWrapper{}

// An interface definition for the [MLMultiArrayAsNSArrayWrapper] class.
//
// # Methods
//
//   - [IMLMultiArrayAsNSArrayWrapper.MultiArray]
//   - [IMLMultiArrayAsNSArrayWrapper.SetMultiArray]
//   - [IMLMultiArrayAsNSArrayWrapper.InitWrappingMultiArray]
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayAsNSArrayWrapper
type IMLMultiArrayAsNSArrayWrapper interface {
	foundation.INSArray

	// Topic: Methods

	MultiArray() IMLMultiArray
	SetMultiArray(value IMLMultiArray)
	InitWrappingMultiArray(array objectivec.IObject) MLMultiArrayAsNSArrayWrapper
}

// Init initializes the instance.
func (m MLMultiArrayAsNSArrayWrapper) Init() MLMultiArrayAsNSArrayWrapper {
	rv := objc.Send[MLMultiArrayAsNSArrayWrapper](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLMultiArrayAsNSArrayWrapper) Autorelease() MLMultiArrayAsNSArrayWrapper {
	rv := objc.Send[MLMultiArrayAsNSArrayWrapper](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLMultiArrayAsNSArrayWrapper creates a new MLMultiArrayAsNSArrayWrapper instance.
func NewMLMultiArrayAsNSArrayWrapper() MLMultiArrayAsNSArrayWrapper {
	class := getMLMultiArrayAsNSArrayWrapperClass()
	rv := objc.Send[MLMultiArrayAsNSArrayWrapper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayAsNSArrayWrapper/initWrappingMultiArray:
func NewMultiArrayAsNSArrayWrapperWrappingMultiArray(array objectivec.IObject) MLMultiArrayAsNSArrayWrapper {
	instance := getMLMultiArrayAsNSArrayWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWrappingMultiArray:"), array)
	return MLMultiArrayAsNSArrayWrapperFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayAsNSArrayWrapper/initWrappingMultiArray:
func (m MLMultiArrayAsNSArrayWrapper) InitWrappingMultiArray(array objectivec.IObject) MLMultiArrayAsNSArrayWrapper {
	rv := objc.Send[MLMultiArrayAsNSArrayWrapper](m.ID, objc.Sel("initWrappingMultiArray:"), array)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayAsNSArrayWrapper/multiArray
func (m MLMultiArrayAsNSArrayWrapper) MultiArray() IMLMultiArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("multiArray"))
	return MLMultiArrayFromID(objc.ID(rv))
}
func (m MLMultiArrayAsNSArrayWrapper) SetMultiArray(value IMLMultiArray) {
	objc.Send[struct{}](m.ID, objc.Sel("setMultiArray:"), value)
}
