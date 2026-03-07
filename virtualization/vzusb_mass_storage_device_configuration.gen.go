// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VZUSBMassStorageDeviceConfiguration] class.
var (
	_VZUSBMassStorageDeviceConfigurationClass     VZUSBMassStorageDeviceConfigurationClass
	_VZUSBMassStorageDeviceConfigurationClassOnce sync.Once
)

func getVZUSBMassStorageDeviceConfigurationClass() VZUSBMassStorageDeviceConfigurationClass {
	_VZUSBMassStorageDeviceConfigurationClassOnce.Do(func() {
		_VZUSBMassStorageDeviceConfigurationClass = VZUSBMassStorageDeviceConfigurationClass{class: objc.GetClass("VZUSBMassStorageDeviceConfiguration")}
	})
	return _VZUSBMassStorageDeviceConfigurationClass
}

// GetVZUSBMassStorageDeviceConfigurationClass returns the class object for VZUSBMassStorageDeviceConfiguration.
func GetVZUSBMassStorageDeviceConfigurationClass() VZUSBMassStorageDeviceConfigurationClass {
	return getVZUSBMassStorageDeviceConfigurationClass()
}

type VZUSBMassStorageDeviceConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZUSBMassStorageDeviceConfigurationClass) Alloc() VZUSBMassStorageDeviceConfiguration {
	rv := objc.Send[VZUSBMassStorageDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// The configuration object that represents a USB Mass storage device.
//
// # Creating the configuration object
//
//   - [VZUSBMassStorageDeviceConfiguration.InitWithAttachment]: Creates a new storage device configuration with the specified attachment.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDeviceConfiguration
type VZUSBMassStorageDeviceConfiguration struct {
	VZStorageDeviceConfiguration
}

// VZUSBMassStorageDeviceConfigurationFromID constructs a [VZUSBMassStorageDeviceConfiguration] from an objc.ID.
//
// The configuration object that represents a USB Mass storage device.
func VZUSBMassStorageDeviceConfigurationFromID(id objc.ID) VZUSBMassStorageDeviceConfiguration {
	return VZUSBMassStorageDeviceConfiguration{VZStorageDeviceConfiguration: VZStorageDeviceConfigurationFromID(id)}
}
// NOTE: VZUSBMassStorageDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZUSBMassStorageDeviceConfiguration] class.
//
// # Creating the configuration object
//
//   - [IVZUSBMassStorageDeviceConfiguration.InitWithAttachment]: Creates a new storage device configuration with the specified attachment.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDeviceConfiguration
type IVZUSBMassStorageDeviceConfiguration interface {
	IVZStorageDeviceConfiguration
	VZUSBDeviceConfiguration

	// Topic: Creating the configuration object

	// Creates a new storage device configuration with the specified attachment.
	InitWithAttachment(attachment IVZStorageDeviceAttachment) VZUSBMassStorageDeviceConfiguration
}





// Init initializes the instance.
func (u VZUSBMassStorageDeviceConfiguration) Init() VZUSBMassStorageDeviceConfiguration {
	rv := objc.Send[VZUSBMassStorageDeviceConfiguration](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u VZUSBMassStorageDeviceConfiguration) Autorelease() VZUSBMassStorageDeviceConfiguration {
	rv := objc.Send[VZUSBMassStorageDeviceConfiguration](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBMassStorageDeviceConfiguration creates a new VZUSBMassStorageDeviceConfiguration instance.
func NewVZUSBMassStorageDeviceConfiguration() VZUSBMassStorageDeviceConfiguration {
	class := getVZUSBMassStorageDeviceConfigurationClass()
	rv := objc.Send[VZUSBMassStorageDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a new storage device configuration with the specified attachment.
//
// attachment: A [VZStorageDeviceAttachment] object.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDeviceConfiguration/init(attachment:)
func NewUSBMassStorageDeviceConfigurationWithAttachment(attachment IVZStorageDeviceAttachment) VZUSBMassStorageDeviceConfiguration {
	instance := getVZUSBMassStorageDeviceConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAttachment:"), attachment)
	return VZUSBMassStorageDeviceConfigurationFromID(rv)
}







// Creates a new storage device configuration with the specified attachment.
//
// attachment: A [VZStorageDeviceAttachment] object.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDeviceConfiguration/init(attachment:)
func (u VZUSBMassStorageDeviceConfiguration) InitWithAttachment(attachment IVZStorageDeviceAttachment) VZUSBMassStorageDeviceConfiguration {
	rv := objc.Send[VZUSBMassStorageDeviceConfiguration](u.ID, objc.Sel("initWithAttachment:"), attachment)
	return rv
}












// The device’s unique identifier.
//
// # Discussion
// 
// The framework autogenerates the device UUID.
// 
// Before restoring the VM, you need to set the device’s UUID to the UUID of
// the device with the attachment at the time of saving the VM’s state.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBDeviceConfiguration/uuid
func (u VZUSBMassStorageDeviceConfiguration) Uuid() foundation.NSUUID {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("uuid"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (u VZUSBMassStorageDeviceConfiguration) SetUuid(value foundation.NSUUID) {
	objc.Send[struct{}](u.ID, objc.Sel("setUuid:"), value)
}















			// Protocol methods for VZUSBDeviceConfiguration
			













