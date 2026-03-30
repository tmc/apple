// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
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

// # Methods
//
//   - [MLCPUComputeDevice.DebugDescription]
//   - [MLCPUComputeDevice.Description]
//   - [MLCPUComputeDevice.Hash]
//   - [MLCPUComputeDevice.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLCPUComputeDevice
type MLCPUComputeDevice struct {
	objectivec.Object
}

// MLCPUComputeDeviceFromID constructs a [MLCPUComputeDevice] from an objc.ID.
func MLCPUComputeDeviceFromID(id objc.ID) MLCPUComputeDevice {
	return MLCPUComputeDevice{objectivec.Object{ID: id}}
}

// Ensure MLCPUComputeDevice implements IMLCPUComputeDevice.
var _ IMLCPUComputeDevice = MLCPUComputeDevice{}

// An interface definition for the [MLCPUComputeDevice] class.
//
// # Methods
//
//   - [IMLCPUComputeDevice.DebugDescription]
//   - [IMLCPUComputeDevice.Description]
//   - [IMLCPUComputeDevice.Hash]
//   - [IMLCPUComputeDevice.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLCPUComputeDevice
type IMLCPUComputeDevice interface {
	objectivec.IObject

	// Topic: Methods

	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
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

// See: https://developer.apple.com/documentation/CoreML/MLCPUComputeDevice/physicalDevice
func (_MLCPUComputeDeviceClass MLCPUComputeDeviceClass) PhysicalDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLCPUComputeDeviceClass.class), objc.Sel("physicalDevice"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLCPUComputeDevice/debugDescription
func (c MLCPUComputeDevice) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLCPUComputeDevice/description
func (c MLCPUComputeDevice) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLCPUComputeDevice/hash
func (c MLCPUComputeDevice) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCPUComputeDevice/superclass
func (c MLCPUComputeDevice) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}
