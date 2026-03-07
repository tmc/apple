// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLGPUComputeDevice] class.
var (
	_MLGPUComputeDeviceClass     MLGPUComputeDeviceClass
	_MLGPUComputeDeviceClassOnce sync.Once
)

func getMLGPUComputeDeviceClass() MLGPUComputeDeviceClass {
	_MLGPUComputeDeviceClassOnce.Do(func() {
		_MLGPUComputeDeviceClass = MLGPUComputeDeviceClass{class: objc.GetClass("MLGPUComputeDevice")}
	})
	return _MLGPUComputeDeviceClass
}

// GetMLGPUComputeDeviceClass returns the class object for MLGPUComputeDevice.
func GetMLGPUComputeDeviceClass() MLGPUComputeDeviceClass {
	return getMLGPUComputeDeviceClass()
}

type MLGPUComputeDeviceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLGPUComputeDeviceClass) Alloc() MLGPUComputeDevice {
	rv := objc.Send[MLGPUComputeDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// An object that represents a GPU compute device.
//
// # Getting the metal device
//
//   - [MLGPUComputeDevice.MetalDevice]: The device that represents the underlying metal device.
//
// See: https://developer.apple.com/documentation/CoreML/MLGPUComputeDevice
type MLGPUComputeDevice struct {
	objectivec.Object
}

// MLGPUComputeDeviceFromID constructs a [MLGPUComputeDevice] from an objc.ID.
//
// An object that represents a GPU compute device.
func MLGPUComputeDeviceFromID(id objc.ID) MLGPUComputeDevice {
	return MLGPUComputeDevice{objectivec.Object{id}}
}
// NOTE: MLGPUComputeDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MLGPUComputeDevice] class.
//
// # Getting the metal device
//
//   - [IMLGPUComputeDevice.MetalDevice]: The device that represents the underlying metal device.
//
// See: https://developer.apple.com/documentation/CoreML/MLGPUComputeDevice
type IMLGPUComputeDevice interface {
	objectivec.IObject
	MLComputeDeviceProtocol

	// Topic: Getting the metal device

	// The device that represents the underlying metal device.
	MetalDevice() objectivec.IObject
}





// Init initializes the instance.
func (g MLGPUComputeDevice) Init() MLGPUComputeDevice {
	rv := objc.Send[MLGPUComputeDevice](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MLGPUComputeDevice) Autorelease() MLGPUComputeDevice {
	rv := objc.Send[MLGPUComputeDevice](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLGPUComputeDevice creates a new MLGPUComputeDevice instance.
func NewMLGPUComputeDevice() MLGPUComputeDevice {
	class := getMLGPUComputeDeviceClass()
	rv := objc.Send[MLGPUComputeDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// The device that represents the underlying metal device.
//
// See: https://developer.apple.com/documentation/CoreML/MLGPUComputeDevice/metalDevice
func (g MLGPUComputeDevice) MetalDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("metalDevice"))
	return objectivec.Object{ID: rv}
}













			// Protocol methods for MLComputeDeviceProtocol
			
















