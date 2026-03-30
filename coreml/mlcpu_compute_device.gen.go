// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLCPUComputeDevice] class.
var (
	_MLCPUComputeDeviceClass     MLCPUComputeDeviceClass
	_MLCPUComputeDeviceClassOnce sync.Once
)

func getMLCPUComputeDeviceClass() MLCPUComputeDeviceClass {
	_MLCPUComputeDeviceClassOnce.Do(func() {
		_MLCPUComputeDeviceClass = MLCPUComputeDeviceClass{class: objc.GetClass("MLCPUComputeDevice")}
	})
	return _MLCPUComputeDeviceClass
}

// GetMLCPUComputeDeviceClass returns the class object for MLCPUComputeDevice.
func GetMLCPUComputeDeviceClass() MLCPUComputeDeviceClass {
	return getMLCPUComputeDeviceClass()
}

type MLCPUComputeDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLCPUComputeDeviceClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLCPUComputeDeviceClass) Alloc() MLCPUComputeDevice {
	rv := objc.Send[MLCPUComputeDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a CPU compute device.
//
// See: https://developer.apple.com/documentation/CoreML/MLCPUComputeDevice
type MLCPUComputeDevice struct {
	objectivec.Object
}

// MLCPUComputeDeviceFromID constructs a [MLCPUComputeDevice] from an objc.ID.
//
// An object that represents a CPU compute device.
func MLCPUComputeDeviceFromID(id objc.ID) MLCPUComputeDevice {
	return MLCPUComputeDevice{objectivec.Object{ID: id}}
}

// NOTE: MLCPUComputeDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLCPUComputeDevice] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLCPUComputeDevice
type IMLCPUComputeDevice interface {
	objectivec.IObject
	MLComputeDeviceProtocol
}

// Init initializes the instance.
func (c MLCPUComputeDevice) Init() MLCPUComputeDevice {
	rv := objc.Send[MLCPUComputeDevice](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLCPUComputeDevice) Autorelease() MLCPUComputeDevice {
	rv := objc.Send[MLCPUComputeDevice](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLCPUComputeDevice creates a new MLCPUComputeDevice instance.
func NewMLCPUComputeDevice() MLCPUComputeDevice {
	class := getMLCPUComputeDeviceClass()
	rv := objc.Send[MLCPUComputeDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Protocol methods for MLComputeDeviceProtocol
