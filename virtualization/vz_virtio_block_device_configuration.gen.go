// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VZVirtioBlockDeviceConfiguration] class.
var (
	_VZVirtioBlockDeviceConfigurationClass     VZVirtioBlockDeviceConfigurationClass
	_VZVirtioBlockDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioBlockDeviceConfigurationClass() VZVirtioBlockDeviceConfigurationClass {
	_VZVirtioBlockDeviceConfigurationClassOnce.Do(func() {
		_VZVirtioBlockDeviceConfigurationClass = VZVirtioBlockDeviceConfigurationClass{class: objc.GetClass("VZVirtioBlockDeviceConfiguration")}
	})
	return _VZVirtioBlockDeviceConfigurationClass
}

// GetVZVirtioBlockDeviceConfigurationClass returns the class object for VZVirtioBlockDeviceConfiguration.
func GetVZVirtioBlockDeviceConfigurationClass() VZVirtioBlockDeviceConfigurationClass {
	return getVZVirtioBlockDeviceConfigurationClass()
}

type VZVirtioBlockDeviceConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioBlockDeviceConfigurationClass) Alloc() VZVirtioBlockDeviceConfiguration {
	rv := objc.Send[VZVirtioBlockDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// The configuration object that requests the creation of a virtual storage
// device in the guest system.
//
// # Overview
// 
// Use a [VZVirtioBlockDeviceConfiguration] object to create an emulated
// storage device in your virtual machine. When you add this object to your
// virtual machine configuration, the virtual machine creates an emulated disk
// for the guest operating system to use to read and write files. The emulated
// storage device conforms to the Virtio Block Device specification.
// 
// When you create a [VZVirtioBlockDeviceConfiguration] object, specify the
// attachment object that implements the underlying storage. For example,
// specify a [VZDiskImageStorageDeviceAttachment] object to configure the
// storage device using a disk image in the local file system. Assign your
// configuration object to the [VZVirtioBlockDeviceConfiguration.StorageDevices] property of your
// [VZVirtualMachineConfiguration] object before creating your virtual
// machine.
//
// # Creating the configuration object
//
//   - [VZVirtioBlockDeviceConfiguration.InitWithAttachment]: Creates a block device configuration object that uses the specified storage medium.
//
// # Identifying a block device
//
//   - [VZVirtioBlockDeviceConfiguration.BlockDeviceIdentifier]: The string that identifies the VIRTIO block device.
//   - [VZVirtioBlockDeviceConfiguration.SetBlockDeviceIdentifier]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioBlockDeviceConfiguration
type VZVirtioBlockDeviceConfiguration struct {
	VZStorageDeviceConfiguration
}

// VZVirtioBlockDeviceConfigurationFromID constructs a [VZVirtioBlockDeviceConfiguration] from an objc.ID.
//
// The configuration object that requests the creation of a virtual storage
// device in the guest system.
func VZVirtioBlockDeviceConfigurationFromID(id objc.ID) VZVirtioBlockDeviceConfiguration {
	return VZVirtioBlockDeviceConfiguration{VZStorageDeviceConfiguration: VZStorageDeviceConfigurationFromID(id)}
}
// NOTE: VZVirtioBlockDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZVirtioBlockDeviceConfiguration] class.
//
// # Creating the configuration object
//
//   - [IVZVirtioBlockDeviceConfiguration.InitWithAttachment]: Creates a block device configuration object that uses the specified storage medium.
//
// # Identifying a block device
//
//   - [IVZVirtioBlockDeviceConfiguration.BlockDeviceIdentifier]: The string that identifies the VIRTIO block device.
//   - [IVZVirtioBlockDeviceConfiguration.SetBlockDeviceIdentifier]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioBlockDeviceConfiguration
type IVZVirtioBlockDeviceConfiguration interface {
	IVZStorageDeviceConfiguration

	// Topic: Creating the configuration object

	// Creates a block device configuration object that uses the specified storage medium.
	InitWithAttachment(attachment IVZStorageDeviceAttachment) VZVirtioBlockDeviceConfiguration

	// Topic: Identifying a block device

	// The string that identifies the VIRTIO block device.
	BlockDeviceIdentifier() string
	SetBlockDeviceIdentifier(value string)

	// The array of storage devices that you expose to the guest operating system.
	StorageDevices() IVZStorageDeviceConfiguration
	SetStorageDevices(value IVZStorageDeviceConfiguration)
}





// Init initializes the instance.
func (v VZVirtioBlockDeviceConfiguration) Init() VZVirtioBlockDeviceConfiguration {
	rv := objc.Send[VZVirtioBlockDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioBlockDeviceConfiguration) Autorelease() VZVirtioBlockDeviceConfiguration {
	rv := objc.Send[VZVirtioBlockDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioBlockDeviceConfiguration creates a new VZVirtioBlockDeviceConfiguration instance.
func NewVZVirtioBlockDeviceConfiguration() VZVirtioBlockDeviceConfiguration {
	class := getVZVirtioBlockDeviceConfigurationClass()
	rv := objc.Send[VZVirtioBlockDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a block device configuration object that uses the specified storage
// medium.
//
// attachment: The attachment object that provides the storage for the device. For
// example, specify a [VZDiskImageStorageDeviceAttachment] object to implement
// the storage using a local disk image on the host computer.
//
// # Return Value
// 
// A block storage device configuration object to include in your virtual
// machine’s configuration.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioBlockDeviceConfiguration/init(attachment:)
func NewVirtioBlockDeviceConfigurationWithAttachment(attachment IVZStorageDeviceAttachment) VZVirtioBlockDeviceConfiguration {
	instance := getVZVirtioBlockDeviceConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAttachment:"), attachment)
	return VZVirtioBlockDeviceConfigurationFromID(rv)
}







// Creates a block device configuration object that uses the specified storage
// medium.
//
// attachment: The attachment object that provides the storage for the device. For
// example, specify a [VZDiskImageStorageDeviceAttachment] object to implement
// the storage using a local disk image on the host computer.
//
// # Return Value
// 
// A block storage device configuration object to include in your virtual
// machine’s configuration.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioBlockDeviceConfiguration/init(attachment:)
func (v VZVirtioBlockDeviceConfiguration) InitWithAttachment(attachment IVZStorageDeviceAttachment) VZVirtioBlockDeviceConfiguration {
	rv := objc.Send[VZVirtioBlockDeviceConfiguration](v.ID, objc.Sel("initWithAttachment:"), attachment)
	return rv
}





// Checks the validity of a block device identifier.
//
// blockDeviceIdentifier: The device identifier to validate. In the case of an invalid identifier
// string, the method throws an error that describes why the device identifier
// isn’t valid.
//
// # Discussion
// 
// Use [ValidateBlockDeviceIdentifierError] to validate an identifier string
// before trying to set the [BlockDeviceIdentifier] property. The identifier
// must be an ASCII encodable string of 20 bytes or less.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioBlockDeviceConfiguration/validateBlockDeviceIdentifier(_:)
func (_VZVirtioBlockDeviceConfigurationClass VZVirtioBlockDeviceConfigurationClass) ValidateBlockDeviceIdentifierError(blockDeviceIdentifier string) (bool, error) {
			var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_VZVirtioBlockDeviceConfigurationClass.class), objc.Sel("validateBlockDeviceIdentifier:error:"), objc.String(blockDeviceIdentifier), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateBlockDeviceIdentifier:error: returned NO with nil NSError")
	}
	return rv, nil

}








// The string that identifies the VIRTIO block device.
//
// # Discussion
// 
// Use `blockDeviceIdentifier` to name devices so they’re more discoverable
// in the Linux guest. The identifier must be an ASCII encodable string of 20
// bytes or less.
// 
// Validate the identifier string using [ValidateBlockDeviceIdentifierError]
// before attempting to set this property, for example:
// 
// In a Linux guest, device identifiers are visible in the `/dev/disk/by-id/`
// directory.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioBlockDeviceConfiguration/blockDeviceIdentifier
func (v VZVirtioBlockDeviceConfiguration) BlockDeviceIdentifier() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("blockDeviceIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZVirtioBlockDeviceConfiguration) SetBlockDeviceIdentifier(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("setBlockDeviceIdentifier:"), objc.String(value))
}



// The array of storage devices that you expose to the guest operating system.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/storagedevices
func (v VZVirtioBlockDeviceConfiguration) StorageDevices() IVZStorageDeviceConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("storageDevices"))
	return VZStorageDeviceConfigurationFromID(objc.ID(rv))
}
func (v VZVirtioBlockDeviceConfiguration) SetStorageDevices(value IVZStorageDeviceConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setStorageDevices:"), value)
}
























