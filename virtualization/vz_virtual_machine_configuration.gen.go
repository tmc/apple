// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtualMachineConfiguration] class.
var (
	_VZVirtualMachineConfigurationClass     VZVirtualMachineConfigurationClass
	_VZVirtualMachineConfigurationClassOnce sync.Once
)

func getVZVirtualMachineConfigurationClass() VZVirtualMachineConfigurationClass {
	_VZVirtualMachineConfigurationClassOnce.Do(func() {
		_VZVirtualMachineConfigurationClass = VZVirtualMachineConfigurationClass{class: objc.GetClass("VZVirtualMachineConfiguration")}
	})
	return _VZVirtualMachineConfigurationClass
}

// GetVZVirtualMachineConfigurationClass returns the class object for VZVirtualMachineConfiguration.
func GetVZVirtualMachineConfigurationClass() VZVirtualMachineConfigurationClass {
	return getVZVirtualMachineConfigurationClass()
}

type VZVirtualMachineConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtualMachineConfigurationClass) Alloc() VZVirtualMachineConfiguration {
	rv := objc.Send[VZVirtualMachineConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The environment attributes and list of devices to use during the
// configuration of macOS or Linux VMs.
//
// # Overview
// 
// Use a [VZVirtualMachineConfiguration] object to configure the environment
// for a macOS or Linux VM. This configuration object contains information
// about the VM environment, including the devices that the VM exposes to the
// guest operating system. For example, use the configuration object to
// specify the network interfaces and storage devices that the operating
// system may access. For more information on the devices that macOS and Linux
// guests can support, see the Devices section on the [Virtualization]
// framework page.
// 
// You create and configure [VZVirtualMachineConfiguration] objects directly.
// After validating the configuration object, use it to initialize the
// [VZVirtualMachine] object that manages the virtual environment. The
// smallest valid configuration includes a value for the [BootLoader]
// property; you can also include more devices in the configuration depending
// on the needs of your app, such as graphics devices, shared directories, and
// so on. When you finish configuring the object, call the [VZVirtualMachineConfiguration.ValidateWithError]
// method to determine whether a VM can successfully support your
// configuration. A configuration object is invalid if your app doesn’t have
// the [com.apple.security.virtualization] entitlement.
// 
// For more information on using [VZVirtualMachineConfiguration], see
// [Installing macOS on a Virtual Machine] and [Creating and Running a Linux
// Virtual Machine].
//
// [Creating and Running a Linux Virtual Machine]: https://developer.apple.com/documentation/Virtualization/creating-and-running-a-linux-virtual-machine
// [Installing macOS on a Virtual Machine]: https://developer.apple.com/documentation/Virtualization/installing-macos-on-a-virtual-machine
// [Virtualization]: https://developer.apple.com/documentation/Virtualization
// [com.apple.security.virtualization]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.security.virtualization
//
// # Configuring the guest system
//
//   - [VZVirtualMachineConfiguration.BootLoader]: The guest system to boot when the VM starts.
//   - [VZVirtualMachineConfiguration.SetBootLoader]
//
// # Setting the number of CPUs
//
//   - [VZVirtualMachineConfiguration.CPUCount]: The number of CPUs you make available to the guest operating system.
//   - [VZVirtualMachineConfiguration.SetCPUCount]
//
// # Sizing the memory partition
//
//   - [VZVirtualMachineConfiguration.MemorySize]: The memory size in bytes for the virtual machine. Must be a multiple of 1MB and between minimumAllowedMemorySize and maximumAllowedMemorySize.
//   - [VZVirtualMachineConfiguration.SetMemorySize]
//   - [VZVirtualMachineConfiguration.MemoryBalloonDevices]: An array that you configure with a memory balloon device, used to update the memory in the VM.
//   - [VZVirtualMachineConfiguration.SetMemoryBalloonDevices]
//
// # Adding devices to the VM
//
//   - [VZVirtualMachineConfiguration.ConsoleDevices]: The array of console devices that you expose to the guest operating system.
//   - [VZVirtualMachineConfiguration.SetConsoleDevices]
//   - [VZVirtualMachineConfiguration.NetworkDevices]: The array of network devices that you expose to the guest operating system.
//   - [VZVirtualMachineConfiguration.SetNetworkDevices]
//   - [VZVirtualMachineConfiguration.SocketDevices]: The socket device that you use to implement port-based communication with the guest operating system.
//   - [VZVirtualMachineConfiguration.SetSocketDevices]
//   - [VZVirtualMachineConfiguration.SerialPorts]: The array of serial ports that you expose to the guest operating system.
//   - [VZVirtualMachineConfiguration.SetSerialPorts]
//   - [VZVirtualMachineConfiguration.StorageDevices]: The array of storage devices that you expose to the guest operating system.
//   - [VZVirtualMachineConfiguration.SetStorageDevices]
//   - [VZVirtualMachineConfiguration.EntropyDevices]: The array of randomization devices that you expose to the guest operating system.
//   - [VZVirtualMachineConfiguration.SetEntropyDevices]
//   - [VZVirtualMachineConfiguration.AudioDevices]: The list of audio devices.
//   - [VZVirtualMachineConfiguration.SetAudioDevices]
//   - [VZVirtualMachineConfiguration.DirectorySharingDevices]: The list of directory sharing devices.
//   - [VZVirtualMachineConfiguration.SetDirectorySharingDevices]
//   - [VZVirtualMachineConfiguration.GraphicsDevices]: The list of graphics devices.
//   - [VZVirtualMachineConfiguration.SetGraphicsDevices]
//   - [VZVirtualMachineConfiguration.Keyboards]: The list of keyboards.
//   - [VZVirtualMachineConfiguration.SetKeyboards]
//   - [VZVirtualMachineConfiguration.Platform]: The hardware platform to use.
//   - [VZVirtualMachineConfiguration.SetPlatform]
//   - [VZVirtualMachineConfiguration.PointingDevices]: The list of pointing devices.
//   - [VZVirtualMachineConfiguration.SetPointingDevices]
//   - [VZVirtualMachineConfiguration.UsbControllers]: The list of configured USB controllers for the VM.
//   - [VZVirtualMachineConfiguration.SetUsbControllers]
//
// # Validating the configuration
//
//   - [VZVirtualMachineConfiguration.ValidateWithError]: Validates the current configuration settings and reports any issues that might prevent the successful initialization of the VM.
//   - [VZVirtualMachineConfiguration.ValidateSaveRestoreSupportWithError]: Determines whether the framework can save or restore the VM’s current configuration.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration
type VZVirtualMachineConfiguration struct {
	objectivec.Object
}

// VZVirtualMachineConfigurationFromID constructs a [VZVirtualMachineConfiguration] from an objc.ID.
//
// The environment attributes and list of devices to use during the
// configuration of macOS or Linux VMs.
func VZVirtualMachineConfigurationFromID(id objc.ID) VZVirtualMachineConfiguration {
	return VZVirtualMachineConfiguration{objectivec.Object{ID: id}}
}
// NOTE: VZVirtualMachineConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtualMachineConfiguration] class.
//
// # Configuring the guest system
//
//   - [IVZVirtualMachineConfiguration.BootLoader]: The guest system to boot when the VM starts.
//   - [IVZVirtualMachineConfiguration.SetBootLoader]
//
// # Setting the number of CPUs
//
//   - [IVZVirtualMachineConfiguration.CPUCount]: The number of CPUs you make available to the guest operating system.
//   - [IVZVirtualMachineConfiguration.SetCPUCount]
//
// # Sizing the memory partition
//
//   - [IVZVirtualMachineConfiguration.MemorySize]: The memory size in bytes for the virtual machine. Must be a multiple of 1MB and between minimumAllowedMemorySize and maximumAllowedMemorySize.
//   - [IVZVirtualMachineConfiguration.SetMemorySize]
//   - [IVZVirtualMachineConfiguration.MemoryBalloonDevices]: An array that you configure with a memory balloon device, used to update the memory in the VM.
//   - [IVZVirtualMachineConfiguration.SetMemoryBalloonDevices]
//
// # Adding devices to the VM
//
//   - [IVZVirtualMachineConfiguration.ConsoleDevices]: The array of console devices that you expose to the guest operating system.
//   - [IVZVirtualMachineConfiguration.SetConsoleDevices]
//   - [IVZVirtualMachineConfiguration.NetworkDevices]: The array of network devices that you expose to the guest operating system.
//   - [IVZVirtualMachineConfiguration.SetNetworkDevices]
//   - [IVZVirtualMachineConfiguration.SocketDevices]: The socket device that you use to implement port-based communication with the guest operating system.
//   - [IVZVirtualMachineConfiguration.SetSocketDevices]
//   - [IVZVirtualMachineConfiguration.SerialPorts]: The array of serial ports that you expose to the guest operating system.
//   - [IVZVirtualMachineConfiguration.SetSerialPorts]
//   - [IVZVirtualMachineConfiguration.StorageDevices]: The array of storage devices that you expose to the guest operating system.
//   - [IVZVirtualMachineConfiguration.SetStorageDevices]
//   - [IVZVirtualMachineConfiguration.EntropyDevices]: The array of randomization devices that you expose to the guest operating system.
//   - [IVZVirtualMachineConfiguration.SetEntropyDevices]
//   - [IVZVirtualMachineConfiguration.AudioDevices]: The list of audio devices.
//   - [IVZVirtualMachineConfiguration.SetAudioDevices]
//   - [IVZVirtualMachineConfiguration.DirectorySharingDevices]: The list of directory sharing devices.
//   - [IVZVirtualMachineConfiguration.SetDirectorySharingDevices]
//   - [IVZVirtualMachineConfiguration.GraphicsDevices]: The list of graphics devices.
//   - [IVZVirtualMachineConfiguration.SetGraphicsDevices]
//   - [IVZVirtualMachineConfiguration.Keyboards]: The list of keyboards.
//   - [IVZVirtualMachineConfiguration.SetKeyboards]
//   - [IVZVirtualMachineConfiguration.Platform]: The hardware platform to use.
//   - [IVZVirtualMachineConfiguration.SetPlatform]
//   - [IVZVirtualMachineConfiguration.PointingDevices]: The list of pointing devices.
//   - [IVZVirtualMachineConfiguration.SetPointingDevices]
//   - [IVZVirtualMachineConfiguration.UsbControllers]: The list of configured USB controllers for the VM.
//   - [IVZVirtualMachineConfiguration.SetUsbControllers]
//
// # Validating the configuration
//
//   - [IVZVirtualMachineConfiguration.ValidateWithError]: Validates the current configuration settings and reports any issues that might prevent the successful initialization of the VM.
//   - [IVZVirtualMachineConfiguration.ValidateSaveRestoreSupportWithError]: Determines whether the framework can save or restore the VM’s current configuration.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration
type IVZVirtualMachineConfiguration interface {
	objectivec.IObject

	// Topic: Configuring the guest system

	// The guest system to boot when the VM starts.
	BootLoader() IVZBootLoader
	SetBootLoader(value IVZBootLoader)

	// Topic: Setting the number of CPUs

	// The number of CPUs you make available to the guest operating system.
	CPUCount() uint
	SetCPUCount(value uint)

	// Topic: Sizing the memory partition

	// The memory size in bytes for the virtual machine. Must be a multiple of 1MB and between minimumAllowedMemorySize and maximumAllowedMemorySize.
	MemorySize() uint64
	SetMemorySize(value uint64)
	// An array that you configure with a memory balloon device, used to update the memory in the VM.
	MemoryBalloonDevices() []VZMemoryBalloonDeviceConfiguration
	SetMemoryBalloonDevices(value []VZMemoryBalloonDeviceConfiguration)

	// Topic: Adding devices to the VM

	// The array of console devices that you expose to the guest operating system.
	ConsoleDevices() []VZConsoleDeviceConfiguration
	SetConsoleDevices(value []VZConsoleDeviceConfiguration)
	// The array of network devices that you expose to the guest operating system.
	NetworkDevices() []VZNetworkDeviceConfiguration
	SetNetworkDevices(value []VZNetworkDeviceConfiguration)
	// The socket device that you use to implement port-based communication with the guest operating system.
	SocketDevices() []VZSocketDeviceConfiguration
	SetSocketDevices(value []VZSocketDeviceConfiguration)
	// The array of serial ports that you expose to the guest operating system.
	SerialPorts() []VZSerialPortConfiguration
	SetSerialPorts(value []VZSerialPortConfiguration)
	// The array of storage devices that you expose to the guest operating system.
	StorageDevices() []VZStorageDeviceConfiguration
	SetStorageDevices(value []VZStorageDeviceConfiguration)
	// The array of randomization devices that you expose to the guest operating system.
	EntropyDevices() []VZEntropyDeviceConfiguration
	SetEntropyDevices(value []VZEntropyDeviceConfiguration)
	// The list of audio devices.
	AudioDevices() []VZAudioDeviceConfiguration
	SetAudioDevices(value []VZAudioDeviceConfiguration)
	// The list of directory sharing devices.
	DirectorySharingDevices() []VZDirectorySharingDeviceConfiguration
	SetDirectorySharingDevices(value []VZDirectorySharingDeviceConfiguration)
	// The list of graphics devices.
	GraphicsDevices() []VZGraphicsDeviceConfiguration
	SetGraphicsDevices(value []VZGraphicsDeviceConfiguration)
	// The list of keyboards.
	Keyboards() []VZKeyboardConfiguration
	SetKeyboards(value []VZKeyboardConfiguration)
	// The hardware platform to use.
	Platform() IVZPlatformConfiguration
	SetPlatform(value IVZPlatformConfiguration)
	// The list of pointing devices.
	PointingDevices() []VZPointingDeviceConfiguration
	SetPointingDevices(value []VZPointingDeviceConfiguration)
	// The list of configured USB controllers for the VM.
	UsbControllers() []VZUSBControllerConfiguration
	SetUsbControllers(value []VZUSBControllerConfiguration)

	// Topic: Validating the configuration

	// Validates the current configuration settings and reports any issues that might prevent the successful initialization of the VM.
	ValidateWithError() (bool, error)
	// Determines whether the framework can save or restore the VM’s current configuration.
	ValidateSaveRestoreSupportWithError() (bool, error)

	// The number of CPUs for the virtual machine. Must be between minimumAllowedCPUCount and maximumAllowedCPUCount.
	CpuCount() uint
	SetCpuCount(value uint)
}

// Init initializes the instance.
func (v VZVirtualMachineConfiguration) Init() VZVirtualMachineConfiguration {
	rv := objc.Send[VZVirtualMachineConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtualMachineConfiguration) Autorelease() VZVirtualMachineConfiguration {
	rv := objc.Send[VZVirtualMachineConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachineConfiguration creates a new VZVirtualMachineConfiguration instance.
func NewVZVirtualMachineConfiguration() VZVirtualMachineConfiguration {
	class := getVZVirtualMachineConfigurationClass()
	rv := objc.Send[VZVirtualMachineConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Validates the current configuration settings and reports any issues that
// might prevent the successful initialization of the VM.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/validate()
func (v VZVirtualMachineConfiguration) ValidateWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("validateWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// Determines whether the framework can save or restore the VM’s current
// configuration.
//
// # Discussion
// 
// Use this method to verify that a virtual machine with this configuration is
// savable.
// 
// - Not all configuration options can be safely saved and restored from the
// configuration file on disk. - If this method returns `false`, the caller
// should expect future calls to [SaveMachineStateToURLCompletionHandler] to
// fail.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/validateSaveRestoreSupport()
func (v VZVirtualMachineConfiguration) ValidateSaveRestoreSupportWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("validateSaveRestoreSupportWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateSaveRestoreSupportWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// The guest system to boot when the VM starts.
//
// # Discussion
// 
// Assign the boot loader object that contains information about how to load
// your guest operating system. For example, to configure your VM with a Linux
// operating system, assign a [VZLinuxBootLoader] object.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/bootLoader
func (v VZVirtualMachineConfiguration) BootLoader() IVZBootLoader {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("bootLoader"))
	return VZBootLoaderFromID(objc.ID(rv))
}
func (v VZVirtualMachineConfiguration) SetBootLoader(value IVZBootLoader) {
	objc.Send[struct{}](v.ID, objc.Sel("setBootLoader:"), value)
}

// The number of CPUs you make available to the guest operating system.
//
// # Discussion
// 
// The value of this property must be greater than or equal to the value in
// [MinimumAllowedCPUCount], and less than or equal to the value in
// [MaximumAllowedCPUCount]. The system uses the number of physical CPUs on
// the current system to determine a default value.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/cpuCount
func (v VZVirtualMachineConfiguration) CPUCount() uint {
	rv := objc.Send[uint](v.ID, objc.Sel("CPUCount"))
	return rv
}
func (v VZVirtualMachineConfiguration) SetCPUCount(value uint) {
	objc.Send[struct{}](v.ID, objc.Sel("setCPUCount:"), value)
}

// The memory size in bytes for the virtual machine. Must be a multiple of 1MB
// and between minimumAllowedMemorySize and maximumAllowedMemorySize.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/memorySize
func (v VZVirtualMachineConfiguration) MemorySize() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("memorySize"))
	return rv
}
func (v VZVirtualMachineConfiguration) SetMemorySize(value uint64) {
	objc.Send[struct{}](v.ID, objc.Sel("setMemorySize:"), value)
}

// An array that you configure with a memory balloon device, used to update
// the memory in the VM.
//
// # Discussion
// 
// Use this property to request a memory balloon device from the VM. The VM
// initially reserves the amount of memory in the [MemorySize] property for
// the guest operating system. A balloon memory device asks the guest system
// to return memory pages that it isn’t using to the VM. You might use the
// device to reclaim memory when the amount of free memory on the host system
// runs low.
// 
// The default value of this property is an empty array, which doesn’t
// result in the creation of a balloon memory device. To create a balloon
// memory device, add a single
// [VZVirtioTraditionalMemoryBalloonDeviceConfiguration] object to the array.
// In response, the VM creates a [VZVirtioTraditionalMemoryBalloonDevice]
// object and adds it to its [MemoryBalloonDevices] property. Use that object
// to change the amount of memory reserved for the guest system.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/memoryBalloonDevices
func (v VZVirtualMachineConfiguration) MemoryBalloonDevices() []VZMemoryBalloonDeviceConfiguration {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("memoryBalloonDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZMemoryBalloonDeviceConfiguration {
		return VZMemoryBalloonDeviceConfigurationFromID(id)
	})
}
func (v VZVirtualMachineConfiguration) SetMemoryBalloonDevices(value []VZMemoryBalloonDeviceConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setMemoryBalloonDevices:"), objectivec.IObjectSliceToNSArray(value))
}

// The array of console devices that you expose to the guest operating system.
//
// # Discussion
// 
// The default value of this property is an empty array.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/consoleDevices
func (v VZVirtualMachineConfiguration) ConsoleDevices() []VZConsoleDeviceConfiguration {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("consoleDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZConsoleDeviceConfiguration {
		return VZConsoleDeviceConfigurationFromID(id)
	})
}
func (v VZVirtualMachineConfiguration) SetConsoleDevices(value []VZConsoleDeviceConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setConsoleDevices:"), objectivec.IObjectSliceToNSArray(value))
}

// The array of network devices that you expose to the guest operating system.
//
// # Discussion
// 
// The default value of this property is an empty array. If your VM supports
// one or more network devices, assign an array of supported network
// configurations to this property.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/networkDevices
func (v VZVirtualMachineConfiguration) NetworkDevices() []VZNetworkDeviceConfiguration {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("networkDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZNetworkDeviceConfiguration {
		return VZNetworkDeviceConfigurationFromID(id)
	})
}
func (v VZVirtualMachineConfiguration) SetNetworkDevices(value []VZNetworkDeviceConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setNetworkDevices:"), objectivec.IObjectSliceToNSArray(value))
}

// The socket device that you use to implement port-based communication with
// the guest operating system.
//
// # Discussion
// 
// The default value of this property is an empty array. If your VM supports
// port-based communication with sockets, create a single
// [VZVirtioSocketDeviceConfiguration] object, add it to an array, and assign
// it to this property. After initializing the virtual machine, use the object
// in its [SocketDevices] property to configure the port information and
// handlers.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/socketDevices
func (v VZVirtualMachineConfiguration) SocketDevices() []VZSocketDeviceConfiguration {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("socketDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZSocketDeviceConfiguration {
		return VZSocketDeviceConfigurationFromID(id)
	})
}
func (v VZVirtualMachineConfiguration) SetSocketDevices(value []VZSocketDeviceConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setSocketDevices:"), objectivec.IObjectSliceToNSArray(value))
}

// The array of serial ports that you expose to the guest operating system.
//
// # Discussion
// 
// The default value of this property is an empty array. If your VM supports
// one or more serial communication ports, assign an array of supported serial
// configurations to this property. Don’t include more than 10
// [VZVirtioConsoleDeviceConfiguration] objects in the array.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/serialPorts
func (v VZVirtualMachineConfiguration) SerialPorts() []VZSerialPortConfiguration {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("serialPorts"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZSerialPortConfiguration {
		return VZSerialPortConfigurationFromID(id)
	})
}
func (v VZVirtualMachineConfiguration) SetSerialPorts(value []VZSerialPortConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setSerialPorts:"), objectivec.IObjectSliceToNSArray(value))
}

// The array of storage devices that you expose to the guest operating system.
//
// # Discussion
// 
// The default value of this property is an empty array. If your VM exposes
// one or more storage devices, assign an array of
// [VZVirtioBlockDeviceConfiguration] objects to this property.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/storageDevices
func (v VZVirtualMachineConfiguration) StorageDevices() []VZStorageDeviceConfiguration {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("storageDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZStorageDeviceConfiguration {
		return VZStorageDeviceConfigurationFromID(id)
	})
}
func (v VZVirtualMachineConfiguration) SetStorageDevices(value []VZStorageDeviceConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setStorageDevices:"), objectivec.IObjectSliceToNSArray(value))
}

// The array of randomization devices that you expose to the guest operating
// system.
//
// # Discussion
// 
// The default value of this property is an empty array. Add one or more
// entropy configuration objects if you want to provide the guest operating
// system with a source of entropy for its random number generator.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/entropyDevices
func (v VZVirtualMachineConfiguration) EntropyDevices() []VZEntropyDeviceConfiguration {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("entropyDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZEntropyDeviceConfiguration {
		return VZEntropyDeviceConfigurationFromID(id)
	})
}
func (v VZVirtualMachineConfiguration) SetEntropyDevices(value []VZEntropyDeviceConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setEntropyDevices:"), objectivec.IObjectSliceToNSArray(value))
}

// The list of audio devices.
//
// # Discussion
// 
// The default value of this property is an empty array. If your VM exposes
// one or more audio devices, assign an array of [VZAudioDeviceConfiguration]
// objects to this property.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/audioDevices
func (v VZVirtualMachineConfiguration) AudioDevices() []VZAudioDeviceConfiguration {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("audioDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZAudioDeviceConfiguration {
		return VZAudioDeviceConfigurationFromID(id)
	})
}
func (v VZVirtualMachineConfiguration) SetAudioDevices(value []VZAudioDeviceConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setAudioDevices:"), objectivec.IObjectSliceToNSArray(value))
}

// The list of directory sharing devices.
//
// # Discussion
// 
// The default value of this property is an empty array. If your VM exposes
// one or more directory sharing devices, assign an array of
// [VZDirectorySharingDeviceConfiguration] objects to this property.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/directorySharingDevices
func (v VZVirtualMachineConfiguration) DirectorySharingDevices() []VZDirectorySharingDeviceConfiguration {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("directorySharingDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZDirectorySharingDeviceConfiguration {
		return VZDirectorySharingDeviceConfigurationFromID(id)
	})
}
func (v VZVirtualMachineConfiguration) SetDirectorySharingDevices(value []VZDirectorySharingDeviceConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setDirectorySharingDevices:"), objectivec.IObjectSliceToNSArray(value))
}

// The list of graphics devices.
//
// # Discussion
// 
// The default value of this property is an empty array. If your VM exposes
// one or more graphics devices, assign an array of
// [VZGraphicsDeviceConfiguration] objects to this property.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/graphicsDevices
func (v VZVirtualMachineConfiguration) GraphicsDevices() []VZGraphicsDeviceConfiguration {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("graphicsDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZGraphicsDeviceConfiguration {
		return VZGraphicsDeviceConfigurationFromID(id)
	})
}
func (v VZVirtualMachineConfiguration) SetGraphicsDevices(value []VZGraphicsDeviceConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setGraphicsDevices:"), objectivec.IObjectSliceToNSArray(value))
}

// The list of keyboards.
//
// # Discussion
// 
// The default value of this property is an empty array. The default value of
// this property is an empty array. If your VM exposes one or more keyboards,
// assign an array of [VZKeyboardConfiguration] objects to this property.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/keyboards
func (v VZVirtualMachineConfiguration) Keyboards() []VZKeyboardConfiguration {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("keyboards"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZKeyboardConfiguration {
		return VZKeyboardConfigurationFromID(id)
	})
}
func (v VZVirtualMachineConfiguration) SetKeyboards(value []VZKeyboardConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setKeyboards:"), objectivec.IObjectSliceToNSArray(value))
}

// The hardware platform to use.
//
// # Discussion
// 
// This can be an instance of either [VZGenericPlatformConfiguration] or
// [VZMacPlatformConfiguration]. The default is
// [VZGenericPlatformConfiguration].
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/platform
func (v VZVirtualMachineConfiguration) Platform() IVZPlatformConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("platform"))
	return VZPlatformConfigurationFromID(objc.ID(rv))
}
func (v VZVirtualMachineConfiguration) SetPlatform(value IVZPlatformConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setPlatform:"), value)
}

// The list of pointing devices.
//
// # Discussion
// 
// The default value of this property is an empty array. If your VM exposes
// one or more pointing devices, assign an array of
// [VZPointingDeviceConfiguration] objects to this property.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/pointingDevices
func (v VZVirtualMachineConfiguration) PointingDevices() []VZPointingDeviceConfiguration {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("pointingDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZPointingDeviceConfiguration {
		return VZPointingDeviceConfigurationFromID(id)
	})
}
func (v VZVirtualMachineConfiguration) SetPointingDevices(value []VZPointingDeviceConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setPointingDevices:"), objectivec.IObjectSliceToNSArray(value))
}

// The list of configured USB controllers for the VM.
//
// # Discussion
// 
// Use this property to attach USB controllers to the VM configuration, as in
// the following example:
// 
// This property contains an empty array if the
// [VZVirtualMachineConfiguration] doesn’t have any configured USB
// controllers.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/usbControllers
func (v VZVirtualMachineConfiguration) UsbControllers() []VZUSBControllerConfiguration {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("usbControllers"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZUSBControllerConfiguration {
		return VZUSBControllerConfigurationFromID(id)
	})
}
func (v VZVirtualMachineConfiguration) SetUsbControllers(value []VZUSBControllerConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setUsbControllers:"), objectivec.IObjectSliceToNSArray(value))
}

// The number of CPUs for the virtual machine. Must be between
// minimumAllowedCPUCount and maximumAllowedCPUCount. [Full Topic]
func (v VZVirtualMachineConfiguration) CpuCount() uint {
	rv := objc.Send[uint](v.ID, objc.Sel("cpuCount"))
	return rv
}
func (v VZVirtualMachineConfiguration) SetCpuCount(value uint) {
	objc.Send[struct{}](v.ID, objc.Sel("setCpuCount:"), value)
}

// The minimum number of CPUs you may configure for the VM.
//
// # Discussion
// 
// The value in the [CPUCount] property must be greater than or equal to the
// value in this property.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/minimumAllowedCPUCount
func (_VZVirtualMachineConfigurationClass VZVirtualMachineConfigurationClass) MinimumAllowedCPUCount() uint {
	rv := objc.Send[uint](objc.ID(_VZVirtualMachineConfigurationClass.class), objc.Sel("minimumAllowedCPUCount"))
	return rv
}

// The maximum number of CPUs you may configure for the VM.
//
// # Discussion
// 
// The value in the [CPUCount] property must be less than or equal to the
// value in this property.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/maximumAllowedCPUCount
func (_VZVirtualMachineConfigurationClass VZVirtualMachineConfigurationClass) MaximumAllowedCPUCount() uint {
	rv := objc.Send[uint](objc.ID(_VZVirtualMachineConfigurationClass.class), objc.Sel("maximumAllowedCPUCount"))
	return rv
}

// The minimum amount of memory that you may configure for the VM.
//
// # Discussion
// 
// The value in the [MemorySize] property must be greater than or equal to the
// value in this property.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/minimumAllowedMemorySize
func (_VZVirtualMachineConfigurationClass VZVirtualMachineConfigurationClass) MinimumAllowedMemorySize() uint64 {
	rv := objc.Send[uint64](objc.ID(_VZVirtualMachineConfigurationClass.class), objc.Sel("minimumAllowedMemorySize"))
	return rv
}

// The maximum amount of memory that you may configure for the VM.
//
// # Discussion
// 
// The value in the [MemorySize] property must be less than or equal to the
// value in this property.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/maximumAllowedMemorySize
func (_VZVirtualMachineConfigurationClass VZVirtualMachineConfigurationClass) MaximumAllowedMemorySize() uint64 {
	rv := objc.Send[uint64](objc.ID(_VZVirtualMachineConfigurationClass.class), objc.Sel("maximumAllowedMemorySize"))
	return rv
}

