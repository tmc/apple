// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
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

// Class returns the underlying Objective-C class pointer.
func (mc MLGPUComputeDeviceClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLGPUComputeDeviceClass) Alloc() MLGPUComputeDevice {
	rv := objc.Send[MLGPUComputeDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLGPUComputeDevice.InitWithMetalDevice]
//   - [MLGPUComputeDevice.DebugDescription]
//   - [MLGPUComputeDevice.Description]
//   - [MLGPUComputeDevice.Hash]
//   - [MLGPUComputeDevice.Superclass]
type MLGPUComputeDevice struct {
	objectivec.Object
}

// MLGPUComputeDeviceFromID constructs a [MLGPUComputeDevice] from an objc.ID.
func MLGPUComputeDeviceFromID(id objc.ID) MLGPUComputeDevice {
	return MLGPUComputeDevice{objectivec.Object{ID: id}}
}

// Ensure MLGPUComputeDevice implements IMLGPUComputeDevice.
var _ IMLGPUComputeDevice = MLGPUComputeDevice{}

// An interface definition for the [MLGPUComputeDevice] class.
//
// # Methods
//
//   - [IMLGPUComputeDevice.InitWithMetalDevice]
//   - [IMLGPUComputeDevice.DebugDescription]
//   - [IMLGPUComputeDevice.Description]
//   - [IMLGPUComputeDevice.Hash]
//   - [IMLGPUComputeDevice.Superclass]
type IMLGPUComputeDevice interface {
	objectivec.IObject

	// Topic: Methods

	InitWithMetalDevice(device objectivec.IObject) MLGPUComputeDevice
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLGPUComputeDevice) Init() MLGPUComputeDevice {
	rv := objc.Send[MLGPUComputeDevice](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLGPUComputeDevice) Autorelease() MLGPUComputeDevice {
	rv := objc.Send[MLGPUComputeDevice](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLGPUComputeDevice creates a new MLGPUComputeDevice instance.
func NewMLGPUComputeDevice() MLGPUComputeDevice {
	class := getMLGPUComputeDeviceClass()
	rv := objc.Send[MLGPUComputeDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUComputeDeviceWithMetalDevice(device objectivec.IObject) MLGPUComputeDevice {
	instance := getMLGPUComputeDeviceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMetalDevice:"), device)
	return MLGPUComputeDeviceFromID(rv)
}

func (m MLGPUComputeDevice) InitWithMetalDevice(device objectivec.IObject) MLGPUComputeDevice {
	rv := objc.Send[MLGPUComputeDevice](m.ID, objc.Sel("initWithMetalDevice:"), device)
	return rv
}

func (_MLGPUComputeDeviceClass MLGPUComputeDeviceClass) DeviceWithMetalDevice(device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLGPUComputeDeviceClass.class), objc.Sel("deviceWithMetalDevice:"), device)
	return objectivec.Object{ID: rv}
}

func (m MLGPUComputeDevice) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLGPUComputeDevice) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLGPUComputeDevice) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLGPUComputeDevice) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
