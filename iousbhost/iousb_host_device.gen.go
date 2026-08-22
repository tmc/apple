// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOUSBHostDevice] class.
var (
	_IOUSBHostDeviceClass     IOUSBHostDeviceClass
	_IOUSBHostDeviceClassOnce sync.Once
)

func getIOUSBHostDeviceClass() IOUSBHostDeviceClass {
	_IOUSBHostDeviceClassOnce.Do(func() {
		_IOUSBHostDeviceClass = IOUSBHostDeviceClass{class: objc.GetClass("IOUSBHostDevice")}
	})
	return _IOUSBHostDeviceClass
}

// GetIOUSBHostDeviceClass returns the class object for IOUSBHostDevice.
func GetIOUSBHostDeviceClass() IOUSBHostDeviceClass {
	return getIOUSBHostDeviceClass()
}

type IOUSBHostDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOUSBHostDeviceClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOUSBHostDeviceClass) Alloc() IOUSBHostDevice {
	rv := objc.Send[IOUSBHostDevice](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// The class that claims and configures devices, retrieves descriptors, and
// sends device requests.
//
// # Overview
//
// This class enables management of the device state, including sending
// control requests to the default endpoint 0, configuring the device, and
// resetting the device. The interest handler also allows monitoring of the
// device state. The client creates the class and initializes it with
// [initWithIOService:options:queue:error:interestHandler:].
//
// # Retrieving Device Descriptors
//
//   - [IOUSBHostDevice.ConfigurationDescriptor]: The currently selected configuration descriptor.
//
// # Resetting the Device
//
//   - [IOUSBHostDevice.ResetWithError]: Terminates the device and attempts to re-enumerate it.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostDevice
//
// [initWithIOService:options:queue:error:interestHandler:]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/initWithIOService:options:queue:error:interestHandler:
type IOUSBHostDevice struct {
	IOUSBHostObject
}

// IOUSBHostDeviceFromID constructs a [IOUSBHostDevice] from an objc.ID.
//
// The class that claims and configures devices, retrieves descriptors, and
// sends device requests.
func IOUSBHostDeviceFromID(id objc.ID) IOUSBHostDevice {
	return IOUSBHostDevice{IOUSBHostObject: IOUSBHostObjectFromID(id)}
}

// NOTE: IOUSBHostDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOUSBHostDevice] class.
//
// # Retrieving Device Descriptors
//
//   - [IIOUSBHostDevice.ConfigurationDescriptor]: The currently selected configuration descriptor.
//
// # Resetting the Device
//
//   - [IIOUSBHostDevice.ResetWithError]: Terminates the device and attempts to re-enumerate it.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostDevice
type IIOUSBHostDevice interface {
	IIOUSBHostObject

	// Topic: Retrieving Device Descriptors

	// The currently selected configuration descriptor.
	ConfigurationDescriptor() objectivec.IObject

	// Topic: Resetting the Device

	// Terminates the device and attempts to re-enumerate it.
	ResetWithError() (bool, error)
}

// Init initializes the instance.
func (u IOUSBHostDevice) Init() IOUSBHostDevice {
	rv := objc.Send[IOUSBHostDevice](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u IOUSBHostDevice) Autorelease() IOUSBHostDevice {
	rv := objc.Send[IOUSBHostDevice](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOUSBHostDevice creates a new IOUSBHostDevice instance.
func NewIOUSBHostDevice() IOUSBHostDevice {
	class := getIOUSBHostDeviceClass()
	rv := objc.Send[IOUSBHostDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Terminates the device and attempts to re-enumerate it.
//
// # Discussion
//
// This function resets and attempts to re-enumerate the USB device, and
// terminates the [IOUSBHostDevice] and all of its child [IOService] objects.
// If the reset is successful, it also creates and registers a new [IOService]
// object after terminating the previous object. After the call returns
// successfully, the framework [IOUSBHostDevice] no longer has a valid
// connection with the [IOService] object.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostDevice/reset()
//
// [IOService]: https://developer.apple.com/documentation/kernel/ioservice
// [IOUSBHostDevice]: https://developer.apple.com/documentation/kernel/iousbhostdevice
func (u IOUSBHostDevice) ResetWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("resetWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("resetWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// The currently selected configuration descriptor.
//
// # Return Value
//
// A pointer to the device’s configuration descriptor, or `nil` if no
// matching descriptor returns.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostDevice/configurationDescriptor
func (u IOUSBHostDevice) ConfigurationDescriptor() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("configurationDescriptor"))
	return objectivec.Object{ID: rv}
}
