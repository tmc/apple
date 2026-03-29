// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLBNNSDevice] class.
var (
	_CoreMLBNNSDeviceClass     CoreMLBNNSDeviceClass
	_CoreMLBNNSDeviceClassOnce sync.Once
)

func getCoreMLBNNSDeviceClass() CoreMLBNNSDeviceClass {
	_CoreMLBNNSDeviceClassOnce.Do(func() {
		_CoreMLBNNSDeviceClass = CoreMLBNNSDeviceClass{class: objc.GetClass("CoreML.BNNSDevice")}
	})
	return _CoreMLBNNSDeviceClass
}

// GetCoreMLBNNSDeviceClass returns the class object for CoreML.BNNSDevice.
func GetCoreMLBNNSDeviceClass() CoreMLBNNSDeviceClass {
	return getCoreMLBNNSDeviceClass()
}

type CoreMLBNNSDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLBNNSDeviceClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLBNNSDeviceClass) Alloc() CoreMLBNNSDevice {
	rv := objc.Send[CoreMLBNNSDevice](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/CoreML.BNNSDevice
type CoreMLBNNSDevice struct {
	objectivec.Object
}

// CoreMLBNNSDeviceFromID constructs a [CoreMLBNNSDevice] from an objc.ID.
func CoreMLBNNSDeviceFromID(id objc.ID) CoreMLBNNSDevice {
	return CoreMLBNNSDevice{objectivec.Object{ID: id}}
}
// NOTE: CoreMLBNNSDevice struct embeds objectivec.Object (parent type unavailable) but
// ICoreMLBNNSDevice embeds the parent interface; skip compile-time assertion.

// An interface definition for the [CoreMLBNNSDevice] class.
//
// See: https://developer.apple.com/documentation/CoreML/CoreML.BNNSDevice
type ICoreMLBNNSDevice interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CoreMLBNNSDevice) Init() CoreMLBNNSDevice {
	rv := objc.Send[CoreMLBNNSDevice](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLBNNSDevice) Autorelease() CoreMLBNNSDevice {
	rv := objc.Send[CoreMLBNNSDevice](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLBNNSDevice creates a new CoreMLBNNSDevice instance.
func NewCoreMLBNNSDevice() CoreMLBNNSDevice {
	class := getCoreMLBNNSDeviceClass()
	rv := objc.Send[CoreMLBNNSDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

