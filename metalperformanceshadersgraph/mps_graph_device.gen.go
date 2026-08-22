// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphDevice] class.
var (
	_MPSGraphDeviceClass     MPSGraphDeviceClass
	_MPSGraphDeviceClassOnce sync.Once
)

func getMPSGraphDeviceClass() MPSGraphDeviceClass {
	_MPSGraphDeviceClassOnce.Do(func() {
		_MPSGraphDeviceClass = MPSGraphDeviceClass{class: objc.GetClass("MPSGraphDevice")}
	})
	return _MPSGraphDeviceClass
}

// GetMPSGraphDeviceClass returns the class object for MPSGraphDevice.
func GetMPSGraphDeviceClass() MPSGraphDeviceClass {
	return getMPSGraphDeviceClass()
}

type MPSGraphDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphDeviceClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphDeviceClass) Alloc() MPSGraphDevice {
	rv := objc.Send[MPSGraphDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that describes the compute device.
//
// # Instance Properties
//
//   - [MPSGraphDevice.MetalDevice]: If device type is Metal then returns the corresponding MTLDevice else nil.
//   - [MPSGraphDevice.Type]: Device of the MPSGraphDevice.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDevice
type MPSGraphDevice struct {
	MPSGraphObject
}

// MPSGraphDeviceFromID constructs a [MPSGraphDevice] from an objc.ID.
//
// A class that describes the compute device.
func MPSGraphDeviceFromID(id objc.ID) MPSGraphDevice {
	return MPSGraphDevice{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphDevice] class.
//
// # Instance Properties
//
//   - [IMPSGraphDevice.MetalDevice]: If device type is Metal then returns the corresponding MTLDevice else nil.
//   - [IMPSGraphDevice.Type]: Device of the MPSGraphDevice.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDevice
type IMPSGraphDevice interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// If device type is Metal then returns the corresponding MTLDevice else nil.
	MetalDevice() metal.MTLDevice
	// Device of the MPSGraphDevice.
	Type() MPSGraphDeviceType
}

// Init initializes the instance.
func (g MPSGraphDevice) Init() MPSGraphDevice {
	rv := objc.Send[MPSGraphDevice](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphDevice) Autorelease() MPSGraphDevice {
	rv := objc.Send[MPSGraphDevice](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphDevice creates a new MPSGraphDevice instance.
func NewMPSGraphDevice() MPSGraphDevice {
	class := getMPSGraphDeviceClass()
	rv := objc.Send[MPSGraphDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a device from a given Metal device.
//
// metalDevice: [MTLDevice] to create an MPSGraphDevice from.
//
// # Return Value
//
// A valid device.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDevice/init(mtlDevice:)
func NewGraphDeviceWithMTLDevice(metalDevice metal.MTLDevice) MPSGraphDevice {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphDeviceClass().class), objc.Sel("deviceWithMTLDevice:"), metalDevice)
	return MPSGraphDeviceFromID(rv)
}

// If device type is Metal then returns the corresponding MTLDevice else nil.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDevice/metalDevice
func (g MPSGraphDevice) MetalDevice() metal.MTLDevice {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("metalDevice"))
	return metal.MTLDeviceObjectFromID(rv)
}

// Device of the MPSGraphDevice.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDevice/type
func (g MPSGraphDevice) Type() MPSGraphDeviceType {
	rv := objc.Send[MPSGraphDeviceType](g.ID, objc.Sel("type"))
	return MPSGraphDeviceType(rv)
}
