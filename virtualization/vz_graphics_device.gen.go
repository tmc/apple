// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZGraphicsDevice] class.
var (
	_VZGraphicsDeviceClass     VZGraphicsDeviceClass
	_VZGraphicsDeviceClassOnce sync.Once
)

func getVZGraphicsDeviceClass() VZGraphicsDeviceClass {
	_VZGraphicsDeviceClassOnce.Do(func() {
		_VZGraphicsDeviceClass = VZGraphicsDeviceClass{class: objc.GetClass("VZGraphicsDevice")}
	})
	return _VZGraphicsDeviceClass
}

// GetVZGraphicsDeviceClass returns the class object for VZGraphicsDevice.
func GetVZGraphicsDeviceClass() VZGraphicsDeviceClass {
	return getVZGraphicsDeviceClass()
}

type VZGraphicsDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZGraphicsDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZGraphicsDeviceClass) Alloc() VZGraphicsDevice {
	rv := objc.Send[VZGraphicsDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A class that represents a graphics device in a VM.
//
// # Overview
// 
// You don’t instantiate a [VZGraphicsDevice] directly. Graphics devices are
// first configured on the [VZVirtualMachineConfiguration] through a subclass
// of [VZGraphicsDeviceConfiguration].
// 
// When the framework creates a [VZVirtualMachine] from the configuration, the
// graphics devices are available through the [VZGraphicsDevice.GraphicsDevices] property.
// 
// The real type of [VZGraphicsDevice] corresponds to the type used by the
// configuration.
// 
// For example, a [VZVirtioGraphicsDeviceConfiguration] leads to a device of
// type [VZVirtioGraphicsDevice] and a [VZMacGraphicsDeviceConfiguration]
// leads to a device of type [VZMacGraphicsDevice].
//
// # Getting the device’s displays
//
//   - [VZGraphicsDevice.Displays]: The list of graphics displays configured for this graphics device.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDevice
type VZGraphicsDevice struct {
	objectivec.Object
}

// VZGraphicsDeviceFromID constructs a [VZGraphicsDevice] from an objc.ID.
//
// A class that represents a graphics device in a VM.
func VZGraphicsDeviceFromID(id objc.ID) VZGraphicsDevice {
	return VZGraphicsDevice{objectivec.Object{ID: id}}
}
// NOTE: VZGraphicsDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZGraphicsDevice] class.
//
// # Getting the device’s displays
//
//   - [IVZGraphicsDevice.Displays]: The list of graphics displays configured for this graphics device.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDevice
type IVZGraphicsDevice interface {
	objectivec.IObject

	// Topic: Getting the device’s displays

	// The list of graphics displays configured for this graphics device.
	Displays() []VZGraphicsDisplay

	// The list of configured graphics devices on the virtual machine.
	GraphicsDevices() IVZGraphicsDevice
	SetGraphicsDevices(value IVZGraphicsDevice)
}

// Init initializes the instance.
func (g VZGraphicsDevice) Init() VZGraphicsDevice {
	rv := objc.Send[VZGraphicsDevice](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g VZGraphicsDevice) Autorelease() VZGraphicsDevice {
	rv := objc.Send[VZGraphicsDevice](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGraphicsDevice creates a new VZGraphicsDevice instance.
func NewVZGraphicsDevice() VZGraphicsDevice {
	class := getVZGraphicsDeviceClass()
	rv := objc.Send[VZGraphicsDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The list of graphics displays configured for this graphics device.
//
// # Discussion
// 
// This is a list of the graphics displays configured on the graphics device
// configuration.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDevice/displays
func (g VZGraphicsDevice) Displays() []VZGraphicsDisplay {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("displays"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZGraphicsDisplay {
		return VZGraphicsDisplayFromID(id)
	})
}
// The list of configured graphics devices on the virtual machine.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/graphicsdevices
func (g VZGraphicsDevice) GraphicsDevices() IVZGraphicsDevice {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("graphicsDevices"))
	return VZGraphicsDeviceFromID(objc.ID(rv))
}
func (g VZGraphicsDevice) SetGraphicsDevices(value IVZGraphicsDevice) {
	objc.Send[struct{}](g.ID, objc.Sel("setGraphicsDevices:"), value)
}

