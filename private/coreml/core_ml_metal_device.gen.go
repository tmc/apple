// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLMetalDevice] class.
var (
	_CoreMLMetalDeviceClass     CoreMLMetalDeviceClass
	_CoreMLMetalDeviceClassOnce sync.Once
)

func getCoreMLMetalDeviceClass() CoreMLMetalDeviceClass {
	_CoreMLMetalDeviceClassOnce.Do(func() {
		_CoreMLMetalDeviceClass = CoreMLMetalDeviceClass{class: objc.GetClass("CoreML.MetalDevice")}
	})
	return _CoreMLMetalDeviceClass
}

// GetCoreMLMetalDeviceClass returns the class object for CoreML.MetalDevice.
func GetCoreMLMetalDeviceClass() CoreMLMetalDeviceClass {
	return getCoreMLMetalDeviceClass()
}

type CoreMLMetalDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLMetalDeviceClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLMetalDeviceClass) Alloc() CoreMLMetalDevice {
	rv := objc.Send[CoreMLMetalDevice](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/CoreML.MetalDevice
type CoreMLMetalDevice struct {
	objectivec.Object
}

// CoreMLMetalDeviceFromID constructs a [CoreMLMetalDevice] from an objc.ID.
func CoreMLMetalDeviceFromID(id objc.ID) CoreMLMetalDevice {
	return CoreMLMetalDevice{objectivec.Object{ID: id}}
}
// NOTE: CoreMLMetalDevice struct embeds objectivec.Object (parent type unavailable) but
// ICoreMLMetalDevice embeds the parent interface; skip compile-time assertion.

// An interface definition for the [CoreMLMetalDevice] class.
//
// See: https://developer.apple.com/documentation/CoreML/CoreML.MetalDevice
type ICoreMLMetalDevice interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CoreMLMetalDevice) Init() CoreMLMetalDevice {
	rv := objc.Send[CoreMLMetalDevice](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLMetalDevice) Autorelease() CoreMLMetalDevice {
	rv := objc.Send[CoreMLMetalDevice](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLMetalDevice creates a new CoreMLMetalDevice instance.
func NewCoreMLMetalDevice() CoreMLMetalDevice {
	class := getCoreMLMetalDeviceClass()
	rv := objc.Send[CoreMLMetalDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

