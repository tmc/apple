// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VZUSBMassStorageDevice] class.
var (
	_VZUSBMassStorageDeviceClass     VZUSBMassStorageDeviceClass
	_VZUSBMassStorageDeviceClassOnce sync.Once
)

func getVZUSBMassStorageDeviceClass() VZUSBMassStorageDeviceClass {
	_VZUSBMassStorageDeviceClassOnce.Do(func() {
		_VZUSBMassStorageDeviceClass = VZUSBMassStorageDeviceClass{class: objc.GetClass("VZUSBMassStorageDevice")}
	})
	return _VZUSBMassStorageDeviceClass
}

// GetVZUSBMassStorageDeviceClass returns the class object for VZUSBMassStorageDevice.
func GetVZUSBMassStorageDeviceClass() VZUSBMassStorageDeviceClass {
	return getVZUSBMassStorageDeviceClass()
}

type VZUSBMassStorageDeviceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZUSBMassStorageDeviceClass) Alloc() VZUSBMassStorageDevice {
	rv := objc.Send[VZUSBMassStorageDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A class that represents a hot-pluggable USB mass storage device.
//
// # Overview
// 
// Create this device either by instantiating it directly and passing
// [VZUSBMassStorageDeviceConfiguration] to its initializer, or instantiating
// a [VZUSBMassStorageDeviceConfiguration] in a
// [VZVirtualMachineConfiguration]. Direct instantiation creates an object
// that you can pass to [AttachDeviceCompletionHandler]. Instantiation through
// [VZUSBMassStorageDeviceConfiguration] makes the device available in the
// [VZUSBMassStorageDevice.UsbDevices] property.
//
// # Creating a USB mass storage device
//
//   - [VZUSBMassStorageDevice.InitWithConfiguration]: Creates a USB mass storage device with the provided configuration.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice
type VZUSBMassStorageDevice struct {
	VZStorageDevice
}

// VZUSBMassStorageDeviceFromID constructs a [VZUSBMassStorageDevice] from an objc.ID.
//
// A class that represents a hot-pluggable USB mass storage device.
func VZUSBMassStorageDeviceFromID(id objc.ID) VZUSBMassStorageDevice {
	return VZUSBMassStorageDevice{VZStorageDevice: VZStorageDeviceFromID(id)}
}
// NOTE: VZUSBMassStorageDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZUSBMassStorageDevice] class.
//
// # Creating a USB mass storage device
//
//   - [IVZUSBMassStorageDevice.InitWithConfiguration]: Creates a USB mass storage device with the provided configuration.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice
type IVZUSBMassStorageDevice interface {
	IVZStorageDevice
	VZUSBDevice

	// Topic: Creating a USB mass storage device

	// Creates a USB mass storage device with the provided configuration.
	InitWithConfiguration(configuration IVZUSBMassStorageDeviceConfiguration) VZUSBMassStorageDevice

	// The list of attached USB devices for the controller.
	UsbDevices() VZUSBDevice
	SetUsbDevices(value VZUSBDevice)
}

// Init initializes the instance.
func (u VZUSBMassStorageDevice) Init() VZUSBMassStorageDevice {
	rv := objc.Send[VZUSBMassStorageDevice](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u VZUSBMassStorageDevice) Autorelease() VZUSBMassStorageDevice {
	rv := objc.Send[VZUSBMassStorageDevice](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBMassStorageDevice creates a new VZUSBMassStorageDevice instance.
func NewVZUSBMassStorageDevice() VZUSBMassStorageDevice {
	class := getVZUSBMassStorageDeviceClass()
	rv := objc.Send[VZUSBMassStorageDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a USB mass storage device with the provided configuration.
//
// configuration: The USB mass storage configuration.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/init(configuration:)
func NewUSBMassStorageDeviceWithConfiguration(configuration IVZUSBMassStorageDeviceConfiguration) VZUSBMassStorageDevice {
	instance := getVZUSBMassStorageDeviceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return VZUSBMassStorageDeviceFromID(rv)
}

// Creates a USB mass storage device with the provided configuration.
//
// configuration: The USB mass storage configuration.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/init(configuration:)
func (u VZUSBMassStorageDevice) InitWithConfiguration(configuration IVZUSBMassStorageDeviceConfiguration) VZUSBMassStorageDevice {
	rv := objc.Send[VZUSBMassStorageDevice](u.ID, objc.Sel("initWithConfiguration:"), configuration)
	return rv
}

// The USB controller that has an attachment to the device.
//
// # Discussion
// 
// If a USB device object that conforms to this protocol has a current
// attachment to a USB controller, this property includes a pointer to the
// device’s USB controller object. Otherwise, it’s `nil`.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBDevice/usbController
func (u VZUSBMassStorageDevice) UsbController() IVZUSBController {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("usbController"))
	return VZUSBControllerFromID(objc.ID(rv))
}

// The list of attached USB devices for the controller.
//
// See: https://developer.apple.com/documentation/virtualization/vzusbcontroller/usbdevices
func (u VZUSBMassStorageDevice) UsbDevices() VZUSBDevice {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("usbDevices"))
	return VZUSBDeviceObjectFromID(rv)
}
func (u VZUSBMassStorageDevice) SetUsbDevices(value VZUSBDevice) {
	objc.Send[struct{}](u.ID, objc.Sel("setUsbDevices:"), value)
}

// The device’s unique identifier.
//
// # Discussion
// 
// This is the identifier the system creates from device configuration objects
// that conform to [VZUSBDeviceConfiguration].
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBDevice/uuid
func (u VZUSBMassStorageDevice) Uuid() foundation.NSUUID {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("uuid"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}

			// Protocol methods for VZUSBDevice
			

