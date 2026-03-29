// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLIOSurfaceAllocator] class.
var (
	_CoreMLIOSurfaceAllocatorClass     CoreMLIOSurfaceAllocatorClass
	_CoreMLIOSurfaceAllocatorClassOnce sync.Once
)

func getCoreMLIOSurfaceAllocatorClass() CoreMLIOSurfaceAllocatorClass {
	_CoreMLIOSurfaceAllocatorClassOnce.Do(func() {
		_CoreMLIOSurfaceAllocatorClass = CoreMLIOSurfaceAllocatorClass{class: objc.GetClass("CoreML.IOSurfaceAllocator")}
	})
	return _CoreMLIOSurfaceAllocatorClass
}

// GetCoreMLIOSurfaceAllocatorClass returns the class object for CoreML.IOSurfaceAllocator.
func GetCoreMLIOSurfaceAllocatorClass() CoreMLIOSurfaceAllocatorClass {
	return getCoreMLIOSurfaceAllocatorClass()
}

type CoreMLIOSurfaceAllocatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLIOSurfaceAllocatorClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLIOSurfaceAllocatorClass) Alloc() CoreMLIOSurfaceAllocator {
	rv := objc.Send[CoreMLIOSurfaceAllocator](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/CoreML.IOSurfaceAllocator
type CoreMLIOSurfaceAllocator struct {
	objectivec.Object
}

// CoreMLIOSurfaceAllocatorFromID constructs a [CoreMLIOSurfaceAllocator] from an objc.ID.
func CoreMLIOSurfaceAllocatorFromID(id objc.ID) CoreMLIOSurfaceAllocator {
	return CoreMLIOSurfaceAllocator{objectivec.Object{ID: id}}
}
// NOTE: CoreMLIOSurfaceAllocator struct embeds objectivec.Object (parent type unavailable) but
// ICoreMLIOSurfaceAllocator embeds the parent interface; skip compile-time assertion.

// An interface definition for the [CoreMLIOSurfaceAllocator] class.
//
// See: https://developer.apple.com/documentation/CoreML/CoreML.IOSurfaceAllocator
type ICoreMLIOSurfaceAllocator interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CoreMLIOSurfaceAllocator) Init() CoreMLIOSurfaceAllocator {
	rv := objc.Send[CoreMLIOSurfaceAllocator](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLIOSurfaceAllocator) Autorelease() CoreMLIOSurfaceAllocator {
	rv := objc.Send[CoreMLIOSurfaceAllocator](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLIOSurfaceAllocator creates a new CoreMLIOSurfaceAllocator instance.
func NewCoreMLIOSurfaceAllocator() CoreMLIOSurfaceAllocator {
	class := getCoreMLIOSurfaceAllocatorClass()
	rv := objc.Send[CoreMLIOSurfaceAllocator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

