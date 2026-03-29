// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLIOSurfaceStorage] class.
var (
	_CoreMLIOSurfaceStorageClass     CoreMLIOSurfaceStorageClass
	_CoreMLIOSurfaceStorageClassOnce sync.Once
)

func getCoreMLIOSurfaceStorageClass() CoreMLIOSurfaceStorageClass {
	_CoreMLIOSurfaceStorageClassOnce.Do(func() {
		_CoreMLIOSurfaceStorageClass = CoreMLIOSurfaceStorageClass{class: objc.GetClass("CoreML.IOSurfaceStorage")}
	})
	return _CoreMLIOSurfaceStorageClass
}

// GetCoreMLIOSurfaceStorageClass returns the class object for CoreML.IOSurfaceStorage.
func GetCoreMLIOSurfaceStorageClass() CoreMLIOSurfaceStorageClass {
	return getCoreMLIOSurfaceStorageClass()
}

type CoreMLIOSurfaceStorageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLIOSurfaceStorageClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLIOSurfaceStorageClass) Alloc() CoreMLIOSurfaceStorage {
	rv := objc.Send[CoreMLIOSurfaceStorage](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/CoreML.IOSurfaceStorage
type CoreMLIOSurfaceStorage struct {
	objectivec.Object
}

// CoreMLIOSurfaceStorageFromID constructs a [CoreMLIOSurfaceStorage] from an objc.ID.
func CoreMLIOSurfaceStorageFromID(id objc.ID) CoreMLIOSurfaceStorage {
	return CoreMLIOSurfaceStorage{objectivec.Object{ID: id}}
}
// NOTE: CoreMLIOSurfaceStorage struct embeds objectivec.Object (parent type unavailable) but
// ICoreMLIOSurfaceStorage embeds the parent interface; skip compile-time assertion.

// An interface definition for the [CoreMLIOSurfaceStorage] class.
//
// See: https://developer.apple.com/documentation/CoreML/CoreML.IOSurfaceStorage
type ICoreMLIOSurfaceStorage interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CoreMLIOSurfaceStorage) Init() CoreMLIOSurfaceStorage {
	rv := objc.Send[CoreMLIOSurfaceStorage](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLIOSurfaceStorage) Autorelease() CoreMLIOSurfaceStorage {
	rv := objc.Send[CoreMLIOSurfaceStorage](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLIOSurfaceStorage creates a new CoreMLIOSurfaceStorage instance.
func NewCoreMLIOSurfaceStorage() CoreMLIOSurfaceStorage {
	class := getCoreMLIOSurfaceStorageClass()
	rv := objc.Send[CoreMLIOSurfaceStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

