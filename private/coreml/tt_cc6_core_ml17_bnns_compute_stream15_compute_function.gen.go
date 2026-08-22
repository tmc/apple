// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ComputeFunction] class.
var (
	_ComputeFunctionClass     ComputeFunctionClass
	_ComputeFunctionClassOnce sync.Once
)

func getComputeFunctionClass() ComputeFunctionClass {
	_ComputeFunctionClassOnce.Do(func() {
		_ComputeFunctionClass = ComputeFunctionClass{class: objc.GetClass("_TtCC6CoreML17BNNSComputeStream15ComputeFunction")}
	})
	return _ComputeFunctionClass
}

// GetComputeFunctionClass returns the class object for _TtCC6CoreML17BNNSComputeStream15ComputeFunction.
func GetComputeFunctionClass() ComputeFunctionClass {
	return getComputeFunctionClass()
}

type ComputeFunctionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc ComputeFunctionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc ComputeFunctionClass) Alloc() ComputeFunction {
	rv := objc.SendIfResponds[ComputeFunction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

type ComputeFunction struct {
	objectivec.Object
}

// ComputeFunctionFromID constructs a [ComputeFunction] from an objc.ID.
func ComputeFunctionFromID(id objc.ID) ComputeFunction {
	return ComputeFunction{objectivec.Object{ID: id}}
}

// Ensure ComputeFunction implements IComputeFunction.
var _ IComputeFunction = ComputeFunction{}

// An interface definition for the [ComputeFunction] class.
type IComputeFunction interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c ComputeFunction) Init() ComputeFunction {
	rv := objc.SendIfResponds[ComputeFunction](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c ComputeFunction) Autorelease() ComputeFunction {
	rv := objc.SendIfResponds[ComputeFunction](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewComputeFunction creates a new ComputeFunction instance.
func NewComputeFunction() ComputeFunction {
	class := getComputeFunctionClass()
	rv := objc.SendIfResponds[ComputeFunction](objc.ID(class.class), objc.Sel("new"))
	return rv
}
