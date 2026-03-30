// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

// Package virtualization provides Go bindings for the Virtualization framework.
//
// Create virtual machines and run macOS and Linux-based operating systems.
//
// The Virtualization framework provides high-level APIs for creating and managing virtual machines (VM) on Apple silicon and Intel-based Mac computers. Use this framework to boot and run macOS or Linux-based operating systems in custom environments that you define. The framework supports the [Virtual I/O Device (VIRTIO)](<https://docs.oasis-open.org/virtio/virtio/v1.1/csprd01/virtio-v1.1-csprd01.html>) specification, which defines standard interfaces for many device types, including network, socket, serial port, storage, entropy, and memory-balloon devices.
//
// # Essentials
//
//   - Adding the Virtualization Entitlement to Your Project: Configure your project to use the Virtualization framework.
//   - com.apple.security.virtualization: A Boolean value that indicates whether your app can use the Virtualization framework.
//   - Using iCloud with macOS virtual machines: Access iCloud from macOS guest virtual machines.
//
// # Virtual machine setup
//
//   - Running macOS in a virtual machine on Apple silicon: Install and run macOS in a virtual machine using the Virtualization framework.
//   - Running Linux in a Virtual Machine: Run a Linux operating system on your Mac using the Virtualization framework.
//   - Running GUI Linux in a virtual machine on a Mac: Install and run GUI Linux in a virtual machine using the Virtualization framework.
//   - Installing macOS on a Virtual Machine: Download a macOS restore image and install it in a new VM.
//   - Creating and Running a Linux Virtual Machine: Design and run custom Linux guests on Apple silicon or Intel-based Mac Computers.
//   - Virtualize macOS on a Mac: Configure and run macOS guests on Apple silicon. ([VZVirtualMachineConfiguration], [VZMacOSVirtualMachineStartOptions], [VZMacPlatformConfiguration], [VZPlatformConfiguration], [VZMacHardwareModel])
//   - Virtualize Linux on a Mac: Configure and run Linux guests on Apple silicon and Intel-based Mac computers. ([VZVirtualMachineConfiguration], [VZVirtualMachineStartOptions], [VZGenericPlatformConfiguration], [VZPlatformConfiguration], [VZBootLoader])
//   - Running Intel Binaries in Linux VMs with Rosetta: Run x86_64 Linux binaries under ARM Linux on Apple silicon.
//   - Accelerating the performance of Rosetta: Improve Rosetta performance by adding support for the total store ordering (TSO) memory model to your Linux kernel.
//
// # Runtime
//
//   - VZVirtualMachine: An object that manages the overall state and configuration of your VM. ([VZVirtualMachineDelegate], [VZVirtualMachineStartOptions])
//   - VZVirtualMachineView: A view that allows user interaction with a VM.
//   - VZLinuxRosettaDirectoryShare: The Linux directory share for Rosetta. ([VZLinuxRosettaAvailability])
//
// # Devices
//
//   - Audio: Configure audio devices that enable the guest operating system to perform audio playback and capture through the host’s audio devices. ([VZVirtioSoundDeviceConfiguration], [VZVirtioSoundDeviceOutputStreamConfiguration], [VZVirtioSoundDeviceInputStreamConfiguration], [VZAudioDeviceConfiguration], [VZVirtioSoundDeviceStreamConfiguration])
//   - Graphics: Configure a device for a guest to display its UI. ([VZGraphicsDisplayConfiguration], [VZMacGraphicsDeviceConfiguration], [VZMacGraphicsDisplayConfiguration], [VZGraphicsDeviceConfiguration], [VZVirtioGraphicsDeviceConfiguration])
//   - Keyboards and pointing devices: Configure devices that connect a mouse and keyboard to the guest system. ([VZKeyboardConfiguration], [VZMacKeyboardConfiguration], [VZUSBKeyboardConfiguration], [VZMacTrackpadConfiguration], [VZUSBScreenCoordinatePointingDeviceConfiguration])
//   - Memory: Configure a memory balloon device to change the allocated memory for the guest system. ([VZVirtioTraditionalMemoryBalloonDeviceConfiguration], [VZMemoryBalloonDeviceConfiguration], [VZVirtioTraditionalMemoryBalloonDevice], [VZMemoryBalloonDevice])
//   - Network: Configure the devices that connect the guest system to the network. ([VZVirtioNetworkDeviceConfiguration], [VZNetworkDeviceConfiguration], [VZMACAddress], [VZBridgedNetworkDeviceAttachment], [VZFileHandleNetworkDeviceAttachment])
//   - Randomization: Configure a device for the guest system to use to generate random numbers. ([VZVirtioEntropyDeviceConfiguration], [VZEntropyDeviceConfiguration])
//   - Serial ports: Configure the serial devices that you use to communicate with the guest system. ([VZVirtioConsoleDeviceSerialPortConfiguration], [VZSerialPortConfiguration], [VZFileHandleSerialPortAttachment], [VZFileSerialPortAttachment], [VZSerialPortAttachment])
//   - Shared directories: Configure devices that share directories from the host into the guest system. ([VZVirtioFileSystemDeviceConfiguration], [VZDirectorySharingDeviceConfiguration], [VZLinuxRosettaDirectoryShare], [VZVirtioFileSystemDevice], [VZDirectorySharingDevice])
//   - Sockets: Configure a device that manages port-based communication with the guest system. ([VZVirtioSocketDeviceConfiguration], [VZSocketDeviceConfiguration], [VZVirtioSocketDevice], [VZSocketDevice], [VZVirtioSocketListener])
//   - Storage: Configure the block-storage devices that represent the disks of the guest system. ([VZVirtioBlockDeviceConfiguration], [VZStorageDeviceConfiguration], [VZUSBMassStorageDeviceConfiguration], [VZDiskImageCachingMode], [VZDiskImageSynchronizationMode])
//   - Consoles: Configure a device that manages multiport console communication with the guest system. ([VZConsoleDeviceConfiguration], [VZConsolePortConfiguration], [VZVirtioConsoleDeviceConfiguration], [VZVirtioConsolePortConfiguration], [VZVirtioConsolePortConfigurationArray])
//   - Clipboard sharing: Share the pasteboard between the host and guest system. ([VZSpiceAgentPortAttachment])
//   - USB Devices ([VZUSBMassStorageDevice], [VZUSBControllerConfiguration], [VZXHCIControllerConfiguration], [VZUSBController], [VZXHCIController])
//
// # Enumerations
//
//   - Virtualization enumerations: Control the caching modes, disk synchronization, and macOS auxiliary storage options of VMs. ([VZDiskImageCachingMode], [VZDiskImageSynchronizationMode], [VZDiskSynchronizationMode], [VZLinuxRosettaAvailability])
//
// # Errors
//
//   - VZErrorDomain: The error domain for the Virtualization framework.
//   - VZError: Errors that you might encounter when configuring or using a VM.
//
// # Key Types
//
//   - [VZVirtualMachine] - An object that manages the overall state and configuration of your VM.
//   - [VZVirtualMachineConfiguration] - The environment attributes and list of devices to use during the configuration of macOS or Linux VMs.
//   - [VZMACAddress] - The media access control (MAC) address for a network interface in your virtual machine.
//   - [VZMacOSRestoreImage] - An object that describes a version of macOS to install on to a virtual machine.
//   - [VZNetworkBlockDeviceStorageDeviceAttachment] - A storage device attachment backed by a Network Block Device (NBD) client.
//   - [VZMacAuxiliaryStorage] - An object that contains information the boot loader needs for booting macOS as a guest operating system.
//   - [VZDiskImageStorageDeviceAttachment] - A device that stores content in a disk image.
//   - [VZGenericMachineIdentifier] - An object that represents a unique identifier for a virtual machine.
//   - [VZGraphicsDisplay] - A class that represents a graphics display in a VM.
//   - [VZUSBMassStorageDevice] - A class that represents a hot-pluggable USB mass storage device.
//
// [Virtualization Documentation]: https://developer.apple.com/documentation/Virtualization
package virtualization

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the Virtualization library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/Virtualization.framework/Virtualization",
	"/usr/lib/libVirtualization.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	fmt.Fprintf(os.Stderr, "warning: Virtualization: failed to load framework from any known path\n")
}
