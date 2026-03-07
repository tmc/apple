// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZNVMExpressControllerDeviceConfiguration] class.
var (
	_VZNVMExpressControllerDeviceConfigurationClass     VZNVMExpressControllerDeviceConfigurationClass
	_VZNVMExpressControllerDeviceConfigurationClassOnce sync.Once
)

func getVZNVMExpressControllerDeviceConfigurationClass() VZNVMExpressControllerDeviceConfigurationClass {
	_VZNVMExpressControllerDeviceConfigurationClassOnce.Do(func() {
		_VZNVMExpressControllerDeviceConfigurationClass = VZNVMExpressControllerDeviceConfigurationClass{class: objc.GetClass("VZNVMExpressControllerDeviceConfiguration")}
	})
	return _VZNVMExpressControllerDeviceConfigurationClass
}

// GetVZNVMExpressControllerDeviceConfigurationClass returns the class object for VZNVMExpressControllerDeviceConfiguration.
func GetVZNVMExpressControllerDeviceConfigurationClass() VZNVMExpressControllerDeviceConfigurationClass {
	return getVZNVMExpressControllerDeviceConfigurationClass()
}

type VZNVMExpressControllerDeviceConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZNVMExpressControllerDeviceConfigurationClass) Alloc() VZNVMExpressControllerDeviceConfiguration {
	rv := objc.Send[VZNVMExpressControllerDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// The configuration object that represents an NVM Express Controller storage
// device.
//
// # Overview
// 
// This device configuration creates a storage device that conforms to the
// [NVM Express specification revision 1.1b].
// 
// The device configuration is valid only if used with
// [VZGenericPlatformConfiguration].
//
// [NVM Express specification revision 1.1b]: https://nvmexpress.org/wp-content/uploads/NVM-Express-1_1b-1.pdf
//
// # Creating a new device configuration
//
//   - [VZNVMExpressControllerDeviceConfiguration.InitWithAttachment]: Creates a new NVM Express controller configuration with the storage device attachment you provide.
//
// See: https://developer.apple.com/documentation/Virtualization/VZNVMExpressControllerDeviceConfiguration
type VZNVMExpressControllerDeviceConfiguration struct {
	VZStorageDeviceConfiguration
}

// VZNVMExpressControllerDeviceConfigurationFromID constructs a [VZNVMExpressControllerDeviceConfiguration] from an objc.ID.
//
// The configuration object that represents an NVM Express Controller storage
// device.
func VZNVMExpressControllerDeviceConfigurationFromID(id objc.ID) VZNVMExpressControllerDeviceConfiguration {
	return VZNVMExpressControllerDeviceConfiguration{VZStorageDeviceConfiguration: VZStorageDeviceConfigurationFromID(id)}
}
// NOTE: VZNVMExpressControllerDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZNVMExpressControllerDeviceConfiguration] class.
//
// # Creating a new device configuration
//
//   - [IVZNVMExpressControllerDeviceConfiguration.InitWithAttachment]: Creates a new NVM Express controller configuration with the storage device attachment you provide.
//
// See: https://developer.apple.com/documentation/Virtualization/VZNVMExpressControllerDeviceConfiguration
type IVZNVMExpressControllerDeviceConfiguration interface {
	IVZStorageDeviceConfiguration

	// Topic: Creating a new device configuration

	// Creates a new NVM Express controller configuration with the storage device attachment you provide.
	InitWithAttachment(attachment IVZStorageDeviceAttachment) VZNVMExpressControllerDeviceConfiguration
}





// Init initializes the instance.
func (n VZNVMExpressControllerDeviceConfiguration) Init() VZNVMExpressControllerDeviceConfiguration {
	rv := objc.Send[VZNVMExpressControllerDeviceConfiguration](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n VZNVMExpressControllerDeviceConfiguration) Autorelease() VZNVMExpressControllerDeviceConfiguration {
	rv := objc.Send[VZNVMExpressControllerDeviceConfiguration](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNVMExpressControllerDeviceConfiguration creates a new VZNVMExpressControllerDeviceConfiguration instance.
func NewVZNVMExpressControllerDeviceConfiguration() VZNVMExpressControllerDeviceConfiguration {
	class := getVZNVMExpressControllerDeviceConfigurationClass()
	rv := objc.Send[VZNVMExpressControllerDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a new NVM Express controller configuration with the storage device
// attachment you provide.
//
// attachment: The storage device attachment. This defines how the virtualized device
// operates on the host side.
//
// See: https://developer.apple.com/documentation/Virtualization/VZNVMExpressControllerDeviceConfiguration/init(attachment:)
func NewNVMExpressControllerDeviceConfigurationWithAttachment(attachment IVZStorageDeviceAttachment) VZNVMExpressControllerDeviceConfiguration {
	instance := getVZNVMExpressControllerDeviceConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAttachment:"), attachment)
	return VZNVMExpressControllerDeviceConfigurationFromID(rv)
}







// Creates a new NVM Express controller configuration with the storage device
// attachment you provide.
//
// attachment: The storage device attachment. This defines how the virtualized device
// operates on the host side.
//
// See: https://developer.apple.com/documentation/Virtualization/VZNVMExpressControllerDeviceConfiguration/init(attachment:)
func (n VZNVMExpressControllerDeviceConfiguration) InitWithAttachment(attachment IVZStorageDeviceAttachment) VZNVMExpressControllerDeviceConfiguration {
	rv := objc.Send[VZNVMExpressControllerDeviceConfiguration](n.ID, objc.Sel("initWithAttachment:"), attachment)
	return rv
}

































