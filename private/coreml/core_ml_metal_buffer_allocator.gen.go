// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLMetalBufferAllocator] class.
var (
	_CoreMLMetalBufferAllocatorClass     CoreMLMetalBufferAllocatorClass
	_CoreMLMetalBufferAllocatorClassOnce sync.Once
)

func getCoreMLMetalBufferAllocatorClass() CoreMLMetalBufferAllocatorClass {
	_CoreMLMetalBufferAllocatorClassOnce.Do(func() {
		_CoreMLMetalBufferAllocatorClass = CoreMLMetalBufferAllocatorClass{class: objc.GetClass("CoreML.MetalBufferAllocator")}
	})
	return _CoreMLMetalBufferAllocatorClass
}

// GetCoreMLMetalBufferAllocatorClass returns the class object for CoreML.MetalBufferAllocator.
func GetCoreMLMetalBufferAllocatorClass() CoreMLMetalBufferAllocatorClass {
	return getCoreMLMetalBufferAllocatorClass()
}

type CoreMLMetalBufferAllocatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLMetalBufferAllocatorClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLMetalBufferAllocatorClass) Alloc() CoreMLMetalBufferAllocator {
	rv := objc.Send[CoreMLMetalBufferAllocator](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/CoreML.MetalBufferAllocator
type CoreMLMetalBufferAllocator struct {
	objectivec.Object
}

// CoreMLMetalBufferAllocatorFromID constructs a [CoreMLMetalBufferAllocator] from an objc.ID.
func CoreMLMetalBufferAllocatorFromID(id objc.ID) CoreMLMetalBufferAllocator {
	return CoreMLMetalBufferAllocator{objectivec.Object{ID: id}}
}
// NOTE: CoreMLMetalBufferAllocator struct embeds objectivec.Object (parent type unavailable) but
// ICoreMLMetalBufferAllocator embeds the parent interface; skip compile-time assertion.

// An interface definition for the [CoreMLMetalBufferAllocator] class.
//
// See: https://developer.apple.com/documentation/CoreML/CoreML.MetalBufferAllocator
type ICoreMLMetalBufferAllocator interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CoreMLMetalBufferAllocator) Init() CoreMLMetalBufferAllocator {
	rv := objc.Send[CoreMLMetalBufferAllocator](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLMetalBufferAllocator) Autorelease() CoreMLMetalBufferAllocator {
	rv := objc.Send[CoreMLMetalBufferAllocator](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLMetalBufferAllocator creates a new CoreMLMetalBufferAllocator instance.
func NewCoreMLMetalBufferAllocator() CoreMLMetalBufferAllocator {
	class := getCoreMLMetalBufferAllocatorClass()
	rv := objc.Send[CoreMLMetalBufferAllocator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

