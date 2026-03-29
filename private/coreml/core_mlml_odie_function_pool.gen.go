// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLMLOdieFunctionPool] class.
var (
	_CoreMLMLOdieFunctionPoolClass     CoreMLMLOdieFunctionPoolClass
	_CoreMLMLOdieFunctionPoolClassOnce sync.Once
)

func getCoreMLMLOdieFunctionPoolClass() CoreMLMLOdieFunctionPoolClass {
	_CoreMLMLOdieFunctionPoolClassOnce.Do(func() {
		_CoreMLMLOdieFunctionPoolClass = CoreMLMLOdieFunctionPoolClass{class: objc.GetClass("CoreML.MLOdieFunctionPool")}
	})
	return _CoreMLMLOdieFunctionPoolClass
}

// GetCoreMLMLOdieFunctionPoolClass returns the class object for CoreML.MLOdieFunctionPool.
func GetCoreMLMLOdieFunctionPoolClass() CoreMLMLOdieFunctionPoolClass {
	return getCoreMLMLOdieFunctionPoolClass()
}

type CoreMLMLOdieFunctionPoolClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLMLOdieFunctionPoolClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLMLOdieFunctionPoolClass) Alloc() CoreMLMLOdieFunctionPool {
	rv := objc.Send[CoreMLMLOdieFunctionPool](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/CoreML.MLOdieFunctionPool
type CoreMLMLOdieFunctionPool struct {
	objectivec.Object
}

// CoreMLMLOdieFunctionPoolFromID constructs a [CoreMLMLOdieFunctionPool] from an objc.ID.
func CoreMLMLOdieFunctionPoolFromID(id objc.ID) CoreMLMLOdieFunctionPool {
	return CoreMLMLOdieFunctionPool{objectivec.Object{ID: id}}
}
// NOTE: CoreMLMLOdieFunctionPool struct embeds objectivec.Object (parent type unavailable) but
// ICoreMLMLOdieFunctionPool embeds the parent interface; skip compile-time assertion.

// An interface definition for the [CoreMLMLOdieFunctionPool] class.
//
// See: https://developer.apple.com/documentation/CoreML/CoreML.MLOdieFunctionPool
type ICoreMLMLOdieFunctionPool interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CoreMLMLOdieFunctionPool) Init() CoreMLMLOdieFunctionPool {
	rv := objc.Send[CoreMLMLOdieFunctionPool](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLMLOdieFunctionPool) Autorelease() CoreMLMLOdieFunctionPool {
	rv := objc.Send[CoreMLMLOdieFunctionPool](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLMLOdieFunctionPool creates a new CoreMLMLOdieFunctionPool instance.
func NewCoreMLMLOdieFunctionPool() CoreMLMLOdieFunctionPool {
	class := getCoreMLMLOdieFunctionPoolClass()
	rv := objc.Send[CoreMLMLOdieFunctionPool](objc.ID(class.class), objc.Sel("new"))
	return rv
}

