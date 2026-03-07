// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZUSBScreenCoordinatePointingDeviceConfiguration] class.
var (
	_VZUSBScreenCoordinatePointingDeviceConfigurationClass     VZUSBScreenCoordinatePointingDeviceConfigurationClass
	_VZUSBScreenCoordinatePointingDeviceConfigurationClassOnce sync.Once
)

func getVZUSBScreenCoordinatePointingDeviceConfigurationClass() VZUSBScreenCoordinatePointingDeviceConfigurationClass {
	_VZUSBScreenCoordinatePointingDeviceConfigurationClassOnce.Do(func() {
		_VZUSBScreenCoordinatePointingDeviceConfigurationClass = VZUSBScreenCoordinatePointingDeviceConfigurationClass{class: objc.GetClass("VZUSBScreenCoordinatePointingDeviceConfiguration")}
	})
	return _VZUSBScreenCoordinatePointingDeviceConfigurationClass
}

// GetVZUSBScreenCoordinatePointingDeviceConfigurationClass returns the class object for VZUSBScreenCoordinatePointingDeviceConfiguration.
func GetVZUSBScreenCoordinatePointingDeviceConfigurationClass() VZUSBScreenCoordinatePointingDeviceConfigurationClass {
	return getVZUSBScreenCoordinatePointingDeviceConfigurationClass()
}

type VZUSBScreenCoordinatePointingDeviceConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZUSBScreenCoordinatePointingDeviceConfigurationClass) Alloc() VZUSBScreenCoordinatePointingDeviceConfiguration {
	rv := objc.Send[VZUSBScreenCoordinatePointingDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// An object that defines the configuration for a USB pointing device that
// reports absolute coordinates.
//
// # Overview
// 
// A [VZVirtualMachineView] can use this device to send pointer events to the
// VM.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBScreenCoordinatePointingDeviceConfiguration
type VZUSBScreenCoordinatePointingDeviceConfiguration struct {
	VZPointingDeviceConfiguration
}

// VZUSBScreenCoordinatePointingDeviceConfigurationFromID constructs a [VZUSBScreenCoordinatePointingDeviceConfiguration] from an objc.ID.
//
// An object that defines the configuration for a USB pointing device that
// reports absolute coordinates.
func VZUSBScreenCoordinatePointingDeviceConfigurationFromID(id objc.ID) VZUSBScreenCoordinatePointingDeviceConfiguration {
	return VZUSBScreenCoordinatePointingDeviceConfiguration{VZPointingDeviceConfiguration: VZPointingDeviceConfigurationFromID(id)}
}
// NOTE: VZUSBScreenCoordinatePointingDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZUSBScreenCoordinatePointingDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBScreenCoordinatePointingDeviceConfiguration
type IVZUSBScreenCoordinatePointingDeviceConfiguration interface {
	IVZPointingDeviceConfiguration
}





// Init initializes the instance.
func (u VZUSBScreenCoordinatePointingDeviceConfiguration) Init() VZUSBScreenCoordinatePointingDeviceConfiguration {
	rv := objc.Send[VZUSBScreenCoordinatePointingDeviceConfiguration](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u VZUSBScreenCoordinatePointingDeviceConfiguration) Autorelease() VZUSBScreenCoordinatePointingDeviceConfiguration {
	rv := objc.Send[VZUSBScreenCoordinatePointingDeviceConfiguration](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBScreenCoordinatePointingDeviceConfiguration creates a new VZUSBScreenCoordinatePointingDeviceConfiguration instance.
func NewVZUSBScreenCoordinatePointingDeviceConfiguration() VZUSBScreenCoordinatePointingDeviceConfiguration {
	class := getVZUSBScreenCoordinatePointingDeviceConfigurationClass()
	rv := objc.Send[VZUSBScreenCoordinatePointingDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}











































