// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TtCC6CoreML17BNNSComputeStream15ComputeFunction] class.
var (
	_TtCC6CoreML17BNNSComputeStream15ComputeFunctionClass     TtCC6CoreML17BNNSComputeStream15ComputeFunctionClass
	_TtCC6CoreML17BNNSComputeStream15ComputeFunctionClassOnce sync.Once
)

func getTtCC6CoreML17BNNSComputeStream15ComputeFunctionClass() TtCC6CoreML17BNNSComputeStream15ComputeFunctionClass {
	_TtCC6CoreML17BNNSComputeStream15ComputeFunctionClassOnce.Do(func() {
		_TtCC6CoreML17BNNSComputeStream15ComputeFunctionClass = TtCC6CoreML17BNNSComputeStream15ComputeFunctionClass{class: objc.GetClass("_TtCC6CoreML17BNNSComputeStream15ComputeFunction")}
	})
	return _TtCC6CoreML17BNNSComputeStream15ComputeFunctionClass
}

// GetTtCC6CoreML17BNNSComputeStream15ComputeFunctionClass returns the class object for _TtCC6CoreML17BNNSComputeStream15ComputeFunction.
func GetTtCC6CoreML17BNNSComputeStream15ComputeFunctionClass() TtCC6CoreML17BNNSComputeStream15ComputeFunctionClass {
	return getTtCC6CoreML17BNNSComputeStream15ComputeFunctionClass()
}

type TtCC6CoreML17BNNSComputeStream15ComputeFunctionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TtCC6CoreML17BNNSComputeStream15ComputeFunctionClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TtCC6CoreML17BNNSComputeStream15ComputeFunctionClass) Alloc() TtCC6CoreML17BNNSComputeStream15ComputeFunction {
	rv := objc.Send[TtCC6CoreML17BNNSComputeStream15ComputeFunction](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_TtCC6CoreML17BNNSComputeStream15ComputeFunction
type TtCC6CoreML17BNNSComputeStream15ComputeFunction struct {
	objectivec.Object
}

// TtCC6CoreML17BNNSComputeStream15ComputeFunctionFromID constructs a [TtCC6CoreML17BNNSComputeStream15ComputeFunction] from an objc.ID.
func TtCC6CoreML17BNNSComputeStream15ComputeFunctionFromID(id objc.ID) TtCC6CoreML17BNNSComputeStream15ComputeFunction {
	return TtCC6CoreML17BNNSComputeStream15ComputeFunction{objectivec.Object{ID: id}}
}
// NOTE: TtCC6CoreML17BNNSComputeStream15ComputeFunction struct embeds objectivec.Object (parent type unavailable) but
// ITtCC6CoreML17BNNSComputeStream15ComputeFunction embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TtCC6CoreML17BNNSComputeStream15ComputeFunction] class.
//
// See: https://developer.apple.com/documentation/CoreML/_TtCC6CoreML17BNNSComputeStream15ComputeFunction
type ITtCC6CoreML17BNNSComputeStream15ComputeFunction interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TtCC6CoreML17BNNSComputeStream15ComputeFunction) Init() TtCC6CoreML17BNNSComputeStream15ComputeFunction {
	rv := objc.Send[TtCC6CoreML17BNNSComputeStream15ComputeFunction](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TtCC6CoreML17BNNSComputeStream15ComputeFunction) Autorelease() TtCC6CoreML17BNNSComputeStream15ComputeFunction {
	rv := objc.Send[TtCC6CoreML17BNNSComputeStream15ComputeFunction](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTtCC6CoreML17BNNSComputeStream15ComputeFunction creates a new TtCC6CoreML17BNNSComputeStream15ComputeFunction instance.
func NewTtCC6CoreML17BNNSComputeStream15ComputeFunction() TtCC6CoreML17BNNSComputeStream15ComputeFunction {
	class := getTtCC6CoreML17BNNSComputeStream15ComputeFunctionClass()
	rv := objc.Send[TtCC6CoreML17BNNSComputeStream15ComputeFunction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

