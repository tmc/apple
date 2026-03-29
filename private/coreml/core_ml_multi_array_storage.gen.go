// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLMultiArrayStorage] class.
var (
	_CoreMLMultiArrayStorageClass     CoreMLMultiArrayStorageClass
	_CoreMLMultiArrayStorageClassOnce sync.Once
)

func getCoreMLMultiArrayStorageClass() CoreMLMultiArrayStorageClass {
	_CoreMLMultiArrayStorageClassOnce.Do(func() {
		_CoreMLMultiArrayStorageClass = CoreMLMultiArrayStorageClass{class: objc.GetClass("CoreML.MultiArrayStorage")}
	})
	return _CoreMLMultiArrayStorageClass
}

// GetCoreMLMultiArrayStorageClass returns the class object for CoreML.MultiArrayStorage.
func GetCoreMLMultiArrayStorageClass() CoreMLMultiArrayStorageClass {
	return getCoreMLMultiArrayStorageClass()
}

type CoreMLMultiArrayStorageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLMultiArrayStorageClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLMultiArrayStorageClass) Alloc() CoreMLMultiArrayStorage {
	rv := objc.Send[CoreMLMultiArrayStorage](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/CoreML.MultiArrayStorage
type CoreMLMultiArrayStorage struct {
	objectivec.Object
}

// CoreMLMultiArrayStorageFromID constructs a [CoreMLMultiArrayStorage] from an objc.ID.
func CoreMLMultiArrayStorageFromID(id objc.ID) CoreMLMultiArrayStorage {
	return CoreMLMultiArrayStorage{objectivec.Object{ID: id}}
}
// NOTE: CoreMLMultiArrayStorage struct embeds objectivec.Object (parent type unavailable) but
// ICoreMLMultiArrayStorage embeds the parent interface; skip compile-time assertion.

// An interface definition for the [CoreMLMultiArrayStorage] class.
//
// See: https://developer.apple.com/documentation/CoreML/CoreML.MultiArrayStorage
type ICoreMLMultiArrayStorage interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CoreMLMultiArrayStorage) Init() CoreMLMultiArrayStorage {
	rv := objc.Send[CoreMLMultiArrayStorage](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLMultiArrayStorage) Autorelease() CoreMLMultiArrayStorage {
	rv := objc.Send[CoreMLMultiArrayStorage](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLMultiArrayStorage creates a new CoreMLMultiArrayStorage instance.
func NewCoreMLMultiArrayStorage() CoreMLMultiArrayStorage {
	class := getCoreMLMultiArrayStorageClass()
	rv := objc.Send[CoreMLMultiArrayStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

