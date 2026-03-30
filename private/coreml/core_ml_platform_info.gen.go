// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLPlatformInfo] class.
var (
	_CoreMLPlatformInfoClass     CoreMLPlatformInfoClass
	_CoreMLPlatformInfoClassOnce sync.Once
)

func getCoreMLPlatformInfoClass() CoreMLPlatformInfoClass {
	_CoreMLPlatformInfoClassOnce.Do(func() {
		_CoreMLPlatformInfoClass = CoreMLPlatformInfoClass{class: objc.GetClass("CoreML.PlatformInfo")}
	})
	return _CoreMLPlatformInfoClass
}

// GetCoreMLPlatformInfoClass returns the class object for CoreML.PlatformInfo.
func GetCoreMLPlatformInfoClass() CoreMLPlatformInfoClass {
	return getCoreMLPlatformInfoClass()
}

type CoreMLPlatformInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLPlatformInfoClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLPlatformInfoClass) Alloc() CoreMLPlatformInfo {
	rv := objc.Send[CoreMLPlatformInfo](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/CoreML.PlatformInfo
type CoreMLPlatformInfo struct {
	objectivec.Object
}

// CoreMLPlatformInfoFromID constructs a [CoreMLPlatformInfo] from an objc.ID.
func CoreMLPlatformInfoFromID(id objc.ID) CoreMLPlatformInfo {
	return CoreMLPlatformInfo{objectivec.Object{ID: id}}
}

// NOTE: CoreMLPlatformInfo struct embeds objectivec.Object (parent type unavailable) but
// ICoreMLPlatformInfo embeds the parent interface; skip compile-time assertion.

// An interface definition for the [CoreMLPlatformInfo] class.
//
// See: https://developer.apple.com/documentation/CoreML/CoreML.PlatformInfo
type ICoreMLPlatformInfo interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CoreMLPlatformInfo) Init() CoreMLPlatformInfo {
	rv := objc.Send[CoreMLPlatformInfo](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLPlatformInfo) Autorelease() CoreMLPlatformInfo {
	rv := objc.Send[CoreMLPlatformInfo](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLPlatformInfo creates a new CoreMLPlatformInfo instance.
func NewCoreMLPlatformInfo() CoreMLPlatformInfo {
	class := getCoreMLPlatformInfoClass()
	rv := objc.Send[CoreMLPlatformInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}
