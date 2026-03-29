// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZUSBControllerConfiguration] class.
var (
	_VZUSBControllerConfigurationClass     VZUSBControllerConfigurationClass
	_VZUSBControllerConfigurationClassOnce sync.Once
)

func getVZUSBControllerConfigurationClass() VZUSBControllerConfigurationClass {
	_VZUSBControllerConfigurationClassOnce.Do(func() {
		_VZUSBControllerConfigurationClass = VZUSBControllerConfigurationClass{class: objc.GetClass("VZUSBControllerConfiguration")}
	})
	return _VZUSBControllerConfigurationClass
}

// GetVZUSBControllerConfigurationClass returns the class object for VZUSBControllerConfiguration.
func GetVZUSBControllerConfigurationClass() VZUSBControllerConfigurationClass {
	return getVZUSBControllerConfigurationClass()
}

type VZUSBControllerConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZUSBControllerConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZUSBControllerConfigurationClass) Alloc() VZUSBControllerConfiguration {
	rv := objc.Send[VZUSBControllerConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The base class for a USB controller configuration.
//
// # Overview
// 
// Don’t create [VZUSBControllerConfiguration] objects directly. Use one of
// its subclasses, such as [VZXHCIControllerConfiguration], instead.
//
// # Instance properties
//
//   - [VZUSBControllerConfiguration.UsbDevices]: The list of USB devices.
//   - [VZUSBControllerConfiguration.SetUsbDevices]
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration
type VZUSBControllerConfiguration struct {
	objectivec.Object
}

// VZUSBControllerConfigurationFromID constructs a [VZUSBControllerConfiguration] from an objc.ID.
//
// The base class for a USB controller configuration.
func VZUSBControllerConfigurationFromID(id objc.ID) VZUSBControllerConfiguration {
	return VZUSBControllerConfiguration{objectivec.Object{ID: id}}
}
// NOTE: VZUSBControllerConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZUSBControllerConfiguration] class.
//
// # Instance properties
//
//   - [IVZUSBControllerConfiguration.UsbDevices]: The list of USB devices.
//   - [IVZUSBControllerConfiguration.SetUsbDevices]
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration
type IVZUSBControllerConfiguration interface {
	objectivec.IObject

	// Topic: Instance properties

	// The list of USB devices.
	UsbDevices() []objectivec.IObject
	SetUsbDevices(value []objectivec.IObject)
}

// Init initializes the instance.
func (u VZUSBControllerConfiguration) Init() VZUSBControllerConfiguration {
	rv := objc.Send[VZUSBControllerConfiguration](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u VZUSBControllerConfiguration) Autorelease() VZUSBControllerConfiguration {
	rv := objc.Send[VZUSBControllerConfiguration](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBControllerConfiguration creates a new VZUSBControllerConfiguration instance.
func NewVZUSBControllerConfiguration() VZUSBControllerConfiguration {
	class := getVZUSBControllerConfigurationClass()
	rv := objc.Send[VZUSBControllerConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The list of USB devices.
//
// # Discussion
// 
// This list represents a set of USB devices that a VM starts with. For each
// entry in the list, the system creates a corresponding runtime object in the
// [UsbDevices] property.
// 
// The list is empty by default.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration/usbDevices
func (u VZUSBControllerConfiguration) UsbDevices() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](u.ID, objc.Sel("usbDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (u VZUSBControllerConfiguration) SetUsbDevices(value []objectivec.IObject) {
	objc.Send[struct{}](u.ID, objc.Sel("setUsbDevices:"), objectivec.IObjectSliceToNSArray(value))
}

