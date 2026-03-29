// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLMetalBufferStorage] class.
var (
	_CoreMLMetalBufferStorageClass     CoreMLMetalBufferStorageClass
	_CoreMLMetalBufferStorageClassOnce sync.Once
)

func getCoreMLMetalBufferStorageClass() CoreMLMetalBufferStorageClass {
	_CoreMLMetalBufferStorageClassOnce.Do(func() {
		_CoreMLMetalBufferStorageClass = CoreMLMetalBufferStorageClass{class: objc.GetClass("CoreML.MetalBufferStorage")}
	})
	return _CoreMLMetalBufferStorageClass
}

// GetCoreMLMetalBufferStorageClass returns the class object for CoreML.MetalBufferStorage.
func GetCoreMLMetalBufferStorageClass() CoreMLMetalBufferStorageClass {
	return getCoreMLMetalBufferStorageClass()
}

type CoreMLMetalBufferStorageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLMetalBufferStorageClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLMetalBufferStorageClass) Alloc() CoreMLMetalBufferStorage {
	rv := objc.Send[CoreMLMetalBufferStorage](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/CoreML.MetalBufferStorage
type CoreMLMetalBufferStorage struct {
	objectivec.Object
}

// CoreMLMetalBufferStorageFromID constructs a [CoreMLMetalBufferStorage] from an objc.ID.
func CoreMLMetalBufferStorageFromID(id objc.ID) CoreMLMetalBufferStorage {
	return CoreMLMetalBufferStorage{objectivec.Object{ID: id}}
}
// NOTE: CoreMLMetalBufferStorage struct embeds objectivec.Object (parent type unavailable) but
// ICoreMLMetalBufferStorage embeds the parent interface; skip compile-time assertion.

// An interface definition for the [CoreMLMetalBufferStorage] class.
//
// See: https://developer.apple.com/documentation/CoreML/CoreML.MetalBufferStorage
type ICoreMLMetalBufferStorage interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CoreMLMetalBufferStorage) Init() CoreMLMetalBufferStorage {
	rv := objc.Send[CoreMLMetalBufferStorage](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLMetalBufferStorage) Autorelease() CoreMLMetalBufferStorage {
	rv := objc.Send[CoreMLMetalBufferStorage](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLMetalBufferStorage creates a new CoreMLMetalBufferStorage instance.
func NewCoreMLMetalBufferStorage() CoreMLMetalBufferStorage {
	class := getCoreMLMetalBufferStorageClass()
	rv := objc.Send[CoreMLMetalBufferStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

