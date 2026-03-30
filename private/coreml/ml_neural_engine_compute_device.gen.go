// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
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

// Class returns the underlying Objective-C class pointer.
func (mc MLNeuralEngineComputeDeviceClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNeuralEngineComputeDeviceClass) Alloc() MLNeuralEngineComputeDevice {
	rv := objc.Send[MLNeuralEngineComputeDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLNeuralEngineComputeDevice.InitWithTotalCoreCount]
//   - [MLNeuralEngineComputeDevice.DebugDescription]
//   - [MLNeuralEngineComputeDevice.Description]
//   - [MLNeuralEngineComputeDevice.Hash]
//   - [MLNeuralEngineComputeDevice.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDevice
type MLNeuralEngineComputeDevice struct {
	objectivec.Object
}

// MLNeuralEngineComputeDeviceFromID constructs a [MLNeuralEngineComputeDevice] from an objc.ID.
func MLNeuralEngineComputeDeviceFromID(id objc.ID) MLNeuralEngineComputeDevice {
	return MLNeuralEngineComputeDevice{objectivec.Object{ID: id}}
}

// Ensure MLNeuralEngineComputeDevice implements IMLNeuralEngineComputeDevice.
var _ IMLNeuralEngineComputeDevice = MLNeuralEngineComputeDevice{}

// An interface definition for the [MLNeuralEngineComputeDevice] class.
//
// # Methods
//
//   - [IMLNeuralEngineComputeDevice.InitWithTotalCoreCount]
//   - [IMLNeuralEngineComputeDevice.DebugDescription]
//   - [IMLNeuralEngineComputeDevice.Description]
//   - [IMLNeuralEngineComputeDevice.Hash]
//   - [IMLNeuralEngineComputeDevice.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDevice
type IMLNeuralEngineComputeDevice interface {
	objectivec.IObject

	// Topic: Methods

	InitWithTotalCoreCount(count int64) MLNeuralEngineComputeDevice
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
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

// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDevice/initWithTotalCoreCount:
func NewNeuralEngineComputeDeviceWithTotalCoreCount(count int64) MLNeuralEngineComputeDevice {
	instance := getMLNeuralEngineComputeDeviceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTotalCoreCount:"), count)
	return MLNeuralEngineComputeDeviceFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDevice/initWithTotalCoreCount:
func (n MLNeuralEngineComputeDevice) InitWithTotalCoreCount(count int64) MLNeuralEngineComputeDevice {
	rv := objc.Send[MLNeuralEngineComputeDevice](n.ID, objc.Sel("initWithTotalCoreCount:"), count)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDevice/physicalDevice
func (_MLNeuralEngineComputeDeviceClass MLNeuralEngineComputeDeviceClass) PhysicalDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralEngineComputeDeviceClass.class), objc.Sel("physicalDevice"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDevice/debugDescription
func (n MLNeuralEngineComputeDevice) DebugDescription() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDevice/description
func (n MLNeuralEngineComputeDevice) Description() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDevice/hash
func (n MLNeuralEngineComputeDevice) Hash() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDevice/superclass
func (n MLNeuralEngineComputeDevice) Superclass() objc.Class {
	rv := objc.Send[objc.Class](n.ID, objc.Sel("superclass"))
	return rv
}
