// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioConsoleDevice] class.
var (
	_VZVirtioConsoleDeviceClass     VZVirtioConsoleDeviceClass
	_VZVirtioConsoleDeviceClassOnce sync.Once
)

func getVZVirtioConsoleDeviceClass() VZVirtioConsoleDeviceClass {
	_VZVirtioConsoleDeviceClassOnce.Do(func() {
		_VZVirtioConsoleDeviceClass = VZVirtioConsoleDeviceClass{class: objc.GetClass("VZVirtioConsoleDevice")}
	})
	return _VZVirtioConsoleDeviceClass
}

// GetVZVirtioConsoleDeviceClass returns the class object for VZVirtioConsoleDevice.
func GetVZVirtioConsoleDeviceClass() VZVirtioConsoleDeviceClass {
	return getVZVirtioConsoleDeviceClass()
}

type VZVirtioConsoleDeviceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioConsoleDeviceClass) Alloc() VZVirtioConsoleDevice {
	rv := objc.Send[VZVirtioConsoleDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// A class that represents a Virtio console device in a virtual machine.
//
// # Configuring the console
//
//   - [VZVirtioConsoleDevice.Ports]: The array of console ports that a specific device uses.
//   - [VZVirtioConsoleDevice.Delegate]: The delegate object for the console device.
//   - [VZVirtioConsoleDevice.SetDelegate]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDevice
type VZVirtioConsoleDevice struct {
	VZConsoleDevice
}

// VZVirtioConsoleDeviceFromID constructs a [VZVirtioConsoleDevice] from an objc.ID.
//
// A class that represents a Virtio console device in a virtual machine.
func VZVirtioConsoleDeviceFromID(id objc.ID) VZVirtioConsoleDevice {
	return VZVirtioConsoleDevice{VZConsoleDevice: VZConsoleDeviceFromID(id)}
}
// NOTE: VZVirtioConsoleDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZVirtioConsoleDevice] class.
//
// # Configuring the console
//
//   - [IVZVirtioConsoleDevice.Ports]: The array of console ports that a specific device uses.
//   - [IVZVirtioConsoleDevice.Delegate]: The delegate object for the console device.
//   - [IVZVirtioConsoleDevice.SetDelegate]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDevice
type IVZVirtioConsoleDevice interface {
	IVZConsoleDevice

	// Topic: Configuring the console

	// The array of console ports that a specific device uses.
	Ports() IVZVirtioConsolePortArray
	// The delegate object for the console device.
	Delegate() VZVirtioConsoleDeviceDelegate
	SetDelegate(value VZVirtioConsoleDeviceDelegate)
}





// Init initializes the instance.
func (v VZVirtioConsoleDevice) Init() VZVirtioConsoleDevice {
	rv := objc.Send[VZVirtioConsoleDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioConsoleDevice) Autorelease() VZVirtioConsoleDevice {
	rv := objc.Send[VZVirtioConsoleDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsoleDevice creates a new VZVirtioConsoleDevice instance.
func NewVZVirtioConsoleDevice() VZVirtioConsoleDevice {
	class := getVZVirtioConsoleDeviceClass()
	rv := objc.Send[VZVirtioConsoleDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// The array of console ports that a specific device uses.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDevice/ports
func (v VZVirtioConsoleDevice) Ports() IVZVirtioConsolePortArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("ports"))
	return VZVirtioConsolePortArrayFromID(objc.ID(rv))
}



// The delegate object for the console device.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDevice/delegate
func (v VZVirtioConsoleDevice) Delegate() VZVirtioConsoleDeviceDelegate {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("delegate"))
	return VZVirtioConsoleDeviceDelegateObjectFromID(rv)
}
func (v VZVirtioConsoleDevice) SetDelegate(value VZVirtioConsoleDeviceDelegate) {
	objc.Send[struct{}](v.ID, objc.Sel("setDelegate:"), value)
}























