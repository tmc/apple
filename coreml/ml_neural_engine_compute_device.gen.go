// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNeuralEngineComputeDevice] class.
var (
	_MLNeuralEngineComputeDeviceClass     MLNeuralEngineComputeDeviceClass
	_MLNeuralEngineComputeDeviceClassOnce sync.Once
)

func getMLNeuralEngineComputeDeviceClass() MLNeuralEngineComputeDeviceClass {
	_MLNeuralEngineComputeDeviceClassOnce.Do(func() {
		_MLNeuralEngineComputeDeviceClass = MLNeuralEngineComputeDeviceClass{class: objc.GetClass("MLNeuralEngineComputeDevice")}
	})
	return _MLNeuralEngineComputeDeviceClass
}

// GetMLNeuralEngineComputeDeviceClass returns the class object for MLNeuralEngineComputeDevice.
func GetMLNeuralEngineComputeDeviceClass() MLNeuralEngineComputeDeviceClass {
	return getMLNeuralEngineComputeDeviceClass()
}

type MLNeuralEngineComputeDeviceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNeuralEngineComputeDeviceClass) Alloc() MLNeuralEngineComputeDevice {
	rv := objc.Send[MLNeuralEngineComputeDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a Neural Engine compute device.
//
// # Getting the Total Core Count
//
//   - [MLNeuralEngineComputeDevice.TotalCoreCount]: The total number of cores in the Neural Engine.
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDevice
type MLNeuralEngineComputeDevice struct {
	objectivec.Object
}

// MLNeuralEngineComputeDeviceFromID constructs a [MLNeuralEngineComputeDevice] from an objc.ID.
//
// An object that represents a Neural Engine compute device.
func MLNeuralEngineComputeDeviceFromID(id objc.ID) MLNeuralEngineComputeDevice {
	return MLNeuralEngineComputeDevice{objectivec.Object{ID: id}}
}
// NOTE: MLNeuralEngineComputeDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLNeuralEngineComputeDevice] class.
//
// # Getting the Total Core Count
//
//   - [IMLNeuralEngineComputeDevice.TotalCoreCount]: The total number of cores in the Neural Engine.
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDevice
type IMLNeuralEngineComputeDevice interface {
	objectivec.IObject
	MLComputeDeviceProtocol

	// Topic: Getting the Total Core Count

	// The total number of cores in the Neural Engine.
	TotalCoreCount() int
}

// Init initializes the instance.
func (n MLNeuralEngineComputeDevice) Init() MLNeuralEngineComputeDevice {
	rv := objc.Send[MLNeuralEngineComputeDevice](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MLNeuralEngineComputeDevice) Autorelease() MLNeuralEngineComputeDevice {
	rv := objc.Send[MLNeuralEngineComputeDevice](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNeuralEngineComputeDevice creates a new MLNeuralEngineComputeDevice instance.
func NewMLNeuralEngineComputeDevice() MLNeuralEngineComputeDevice {
	class := getMLNeuralEngineComputeDeviceClass()
	rv := objc.Send[MLNeuralEngineComputeDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The total number of cores in the Neural Engine.
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDevice/totalCoreCount
func (n MLNeuralEngineComputeDevice) TotalCoreCount() int {
	rv := objc.Send[int](n.ID, objc.Sel("totalCoreCount"))
	return rv
}

			// Protocol methods for MLComputeDeviceProtocol
			

