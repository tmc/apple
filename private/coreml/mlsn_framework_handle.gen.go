// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSNFrameworkHandle] class.
var (
	_MLSNFrameworkHandleClass     MLSNFrameworkHandleClass
	_MLSNFrameworkHandleClassOnce sync.Once
)

func getMLSNFrameworkHandleClass() MLSNFrameworkHandleClass {
	_MLSNFrameworkHandleClassOnce.Do(func() {
		_MLSNFrameworkHandleClass = MLSNFrameworkHandleClass{class: objc.GetClass("_MLSNFrameworkHandle")}
	})
	return _MLSNFrameworkHandleClass
}

// GetMLSNFrameworkHandleClass returns the class object for _MLSNFrameworkHandle.
func GetMLSNFrameworkHandleClass() MLSNFrameworkHandleClass {
	return getMLSNFrameworkHandleClass()
}

type MLSNFrameworkHandleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSNFrameworkHandleClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSNFrameworkHandleClass) Alloc() MLSNFrameworkHandle {
	rv := objc.Send[MLSNFrameworkHandle](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLSNFrameworkHandle
type MLSNFrameworkHandle struct {
	objectivec.Object
}

// MLSNFrameworkHandleFromID constructs a [MLSNFrameworkHandle] from an objc.ID.
func MLSNFrameworkHandleFromID(id objc.ID) MLSNFrameworkHandle {
	return MLSNFrameworkHandle{objectivec.Object{ID: id}}
}

// Ensure MLSNFrameworkHandle implements IMLSNFrameworkHandle.
var _ IMLSNFrameworkHandle = MLSNFrameworkHandle{}

// An interface definition for the [MLSNFrameworkHandle] class.
//
// See: https://developer.apple.com/documentation/CoreML/_MLSNFrameworkHandle
type IMLSNFrameworkHandle interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MLSNFrameworkHandle) Init() MLSNFrameworkHandle {
	rv := objc.Send[MLSNFrameworkHandle](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLSNFrameworkHandle) Autorelease() MLSNFrameworkHandle {
	rv := objc.Send[MLSNFrameworkHandle](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSNFrameworkHandle creates a new MLSNFrameworkHandle instance.
func NewMLSNFrameworkHandle() MLSNFrameworkHandle {
	class := getMLSNFrameworkHandleClass()
	rv := objc.Send[MLSNFrameworkHandle](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLSNFrameworkHandle/sharedHandle
func (_MLSNFrameworkHandleClass MLSNFrameworkHandleClass) SharedHandle() *MLSNFrameworkHandle {
	rv := objc.Send[objc.ID](objc.ID(_MLSNFrameworkHandleClass.class), objc.Sel("sharedHandle"))
	if rv == 0 {
		return nil
	}
	val := MLSNFrameworkHandleFromID(rv)
	return &val
}
