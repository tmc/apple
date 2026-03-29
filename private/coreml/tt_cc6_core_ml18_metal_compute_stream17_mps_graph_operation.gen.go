// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TtCC6CoreML18MetalComputeStream17MPSGraphOperation] class.
var (
	_TtCC6CoreML18MetalComputeStream17MPSGraphOperationClass     TtCC6CoreML18MetalComputeStream17MPSGraphOperationClass
	_TtCC6CoreML18MetalComputeStream17MPSGraphOperationClassOnce sync.Once
)

func getTtCC6CoreML18MetalComputeStream17MPSGraphOperationClass() TtCC6CoreML18MetalComputeStream17MPSGraphOperationClass {
	_TtCC6CoreML18MetalComputeStream17MPSGraphOperationClassOnce.Do(func() {
		_TtCC6CoreML18MetalComputeStream17MPSGraphOperationClass = TtCC6CoreML18MetalComputeStream17MPSGraphOperationClass{class: objc.GetClass("_TtCC6CoreML18MetalComputeStream17MPSGraphOperation")}
	})
	return _TtCC6CoreML18MetalComputeStream17MPSGraphOperationClass
}

// GetTtCC6CoreML18MetalComputeStream17MPSGraphOperationClass returns the class object for _TtCC6CoreML18MetalComputeStream17MPSGraphOperation.
func GetTtCC6CoreML18MetalComputeStream17MPSGraphOperationClass() TtCC6CoreML18MetalComputeStream17MPSGraphOperationClass {
	return getTtCC6CoreML18MetalComputeStream17MPSGraphOperationClass()
}

type TtCC6CoreML18MetalComputeStream17MPSGraphOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TtCC6CoreML18MetalComputeStream17MPSGraphOperationClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TtCC6CoreML18MetalComputeStream17MPSGraphOperationClass) Alloc() TtCC6CoreML18MetalComputeStream17MPSGraphOperation {
	rv := objc.Send[TtCC6CoreML18MetalComputeStream17MPSGraphOperation](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_TtCC6CoreML18MetalComputeStream17MPSGraphOperation
type TtCC6CoreML18MetalComputeStream17MPSGraphOperation struct {
	objectivec.Object
}

// TtCC6CoreML18MetalComputeStream17MPSGraphOperationFromID constructs a [TtCC6CoreML18MetalComputeStream17MPSGraphOperation] from an objc.ID.
func TtCC6CoreML18MetalComputeStream17MPSGraphOperationFromID(id objc.ID) TtCC6CoreML18MetalComputeStream17MPSGraphOperation {
	return TtCC6CoreML18MetalComputeStream17MPSGraphOperation{objectivec.Object{ID: id}}
}
// NOTE: TtCC6CoreML18MetalComputeStream17MPSGraphOperation struct embeds objectivec.Object (parent type unavailable) but
// ITtCC6CoreML18MetalComputeStream17MPSGraphOperation embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TtCC6CoreML18MetalComputeStream17MPSGraphOperation] class.
//
// See: https://developer.apple.com/documentation/CoreML/_TtCC6CoreML18MetalComputeStream17MPSGraphOperation
type ITtCC6CoreML18MetalComputeStream17MPSGraphOperation interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TtCC6CoreML18MetalComputeStream17MPSGraphOperation) Init() TtCC6CoreML18MetalComputeStream17MPSGraphOperation {
	rv := objc.Send[TtCC6CoreML18MetalComputeStream17MPSGraphOperation](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TtCC6CoreML18MetalComputeStream17MPSGraphOperation) Autorelease() TtCC6CoreML18MetalComputeStream17MPSGraphOperation {
	rv := objc.Send[TtCC6CoreML18MetalComputeStream17MPSGraphOperation](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTtCC6CoreML18MetalComputeStream17MPSGraphOperation creates a new TtCC6CoreML18MetalComputeStream17MPSGraphOperation instance.
func NewTtCC6CoreML18MetalComputeStream17MPSGraphOperation() TtCC6CoreML18MetalComputeStream17MPSGraphOperation {
	class := getTtCC6CoreML18MetalComputeStream17MPSGraphOperationClass()
	rv := objc.Send[TtCC6CoreML18MetalComputeStream17MPSGraphOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

