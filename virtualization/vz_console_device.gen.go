// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZConsoleDevice] class.
var (
	_VZConsoleDeviceClass     VZConsoleDeviceClass
	_VZConsoleDeviceClassOnce sync.Once
)

func getVZConsoleDeviceClass() VZConsoleDeviceClass {
	_VZConsoleDeviceClassOnce.Do(func() {
		_VZConsoleDeviceClass = VZConsoleDeviceClass{class: objc.GetClass("VZConsoleDevice")}
	})
	return _VZConsoleDeviceClass
}

// GetVZConsoleDeviceClass returns the class object for VZConsoleDevice.
func GetVZConsoleDeviceClass() VZConsoleDeviceClass {
	return getVZConsoleDeviceClass()
}

type VZConsoleDeviceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZConsoleDeviceClass) Alloc() VZConsoleDevice {
	rv := objc.Send[VZConsoleDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A class that represents a console device in a VM.
//
// # Overview
// 
// Don’t instantiate a [VZConsoleDevice] directly: You first configure
// console devices on the [VZVirtualMachineConfiguration] through a subclass
// of [VZConsoleDeviceConfiguration]. After you create [VZVirtualMachine] from
// the configuration, the console devices are available through the
// [VZConsoleDevice.ConsoleDevices] property.
// 
// The actual type of [VZConsoleDevice] corresponds to the type that the
// configuration uses. For example, a [VZVirtioConsoleDeviceConfiguration] is
// a device of type [VZVirtioConsoleDevice].
//
// See: https://developer.apple.com/documentation/Virtualization/VZConsoleDevice
type VZConsoleDevice struct {
	objectivec.Object
}

// VZConsoleDeviceFromID constructs a [VZConsoleDevice] from an objc.ID.
//
// A class that represents a console device in a VM.
func VZConsoleDeviceFromID(id objc.ID) VZConsoleDevice {
	return VZConsoleDevice{objectivec.Object{ID: id}}
}
// NOTE: VZConsoleDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZConsoleDevice] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZConsoleDevice
type IVZConsoleDevice interface {
	objectivec.IObject

	// The list of configured console devices on the VM.
	ConsoleDevices() IVZConsoleDevice
	SetConsoleDevices(value IVZConsoleDevice)
}

// Init initializes the instance.
func (c VZConsoleDevice) Init() VZConsoleDevice {
	rv := objc.Send[VZConsoleDevice](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c VZConsoleDevice) Autorelease() VZConsoleDevice {
	rv := objc.Send[VZConsoleDevice](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZConsoleDevice creates a new VZConsoleDevice instance.
func NewVZConsoleDevice() VZConsoleDevice {
	class := getVZConsoleDeviceClass()
	rv := objc.Send[VZConsoleDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The list of configured console devices on the VM.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/consoledevices
func (c VZConsoleDevice) ConsoleDevices() IVZConsoleDevice {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("consoleDevices"))
	return VZConsoleDeviceFromID(objc.ID(rv))
}
func (c VZConsoleDevice) SetConsoleDevices(value IVZConsoleDevice) {
	objc.Send[struct{}](c.ID, objc.Sel("setConsoleDevices:"), value)
}

