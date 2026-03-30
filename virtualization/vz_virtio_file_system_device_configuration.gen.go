// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioFileSystemDeviceConfiguration] class.
var (
	_VZVirtioFileSystemDeviceConfigurationClass     VZVirtioFileSystemDeviceConfigurationClass
	_VZVirtioFileSystemDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioFileSystemDeviceConfigurationClass() VZVirtioFileSystemDeviceConfigurationClass {
	_VZVirtioFileSystemDeviceConfigurationClassOnce.Do(func() {
		_VZVirtioFileSystemDeviceConfigurationClass = VZVirtioFileSystemDeviceConfigurationClass{class: objc.GetClass("VZVirtioFileSystemDeviceConfiguration")}
	})
	return _VZVirtioFileSystemDeviceConfigurationClass
}

// GetVZVirtioFileSystemDeviceConfigurationClass returns the class object for VZVirtioFileSystemDeviceConfiguration.
func GetVZVirtioFileSystemDeviceConfigurationClass() VZVirtioFileSystemDeviceConfigurationClass {
	return getVZVirtioFileSystemDeviceConfigurationClass()
}

type VZVirtioFileSystemDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioFileSystemDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioFileSystemDeviceConfigurationClass) Alloc() VZVirtioFileSystemDeviceConfiguration {
	rv := objc.Send[VZVirtioFileSystemDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the configuration of a Virtio file system device.
//
// # Overview
//
// Use [VZVirtioFileSystemDeviceConfiguration] to create a Virtio file system
// device which allows the host to expose directories to a guest using a [VZVirtioFileSystemDeviceConfiguration.Tag]
// label.
//
// The example below shows the creation of a
// [VZVirtioFileSystemDeviceConfiguration] that shares a single directory that
// the user can manually mount after creating a mount point in the guest VM:
//
// A [VZVirtioFileSystemDeviceConfiguration] can also share multiple
// directories. The example below demonstrates sharing the `~/Invoices` and
// `~/Projects` directories from the user’s home directory to the guest VM:
//
// # Mounting shared directories
//
// Mounting a shared directory requires the user to execute a command in a
// terminal window, the specific mount command depends on the guest VM’s
// operating system:
//
// - In macOS guests, use `mount_virtiofs tag directory`. - In Linux guests,
// use `mount -t virtiofs tag directory`.
//
// The `tag` argument is the file system device configuration label —
// [Projects] in the single directory case and `myfiles` in the multiple
// directory case in these examples — and `directory` is the mount point in
// the guest file system to which the framework attaches the shared
// directories.
//
// # Automounting shared directories in macOS VMs
//
// In macOS 13 and later, it’s possible to specify a file system device that
// macOS 13 or later guest VMs automount using the [VZVirtioFileSystemDeviceConfiguration.MacOSGuestAutomountTag]
// property. The example below demonstrates sharing a single macOS directory,
// to a macOS guest:
//
// The macOS guest automounts the user’s `~/Projects` directory from the
// host and shares it under `/Volumes/My Shared Files/` in the macOS guest.
//
// It’s also possible to automount multiple directories in macOS guests. The
// example below demonstrates sharing the `~/Invoices` and `~/Projects`
// directories from the user’s home directory to the macOS guest VM:
//
// In the guest, the framework lists each shared directory by its specified
// name under `/Volumes/My Shared Files`. Here, the shared directories in the
// macOS guest are `/Volumes/My Shared Files/My Projects` and `/Volumes/My
// Shared Files/Invoices`.
//
// # Creating a file system device configuration
//
//   - [VZVirtioFileSystemDeviceConfiguration.InitWithTag]: Creates a configuration for a VIRTIO file system device.
//
// # Getting file system information
//
//   - [VZVirtioFileSystemDeviceConfiguration.Share]: A value that defines how the host exposes resources to the guest virtual machine.
//   - [VZVirtioFileSystemDeviceConfiguration.SetShare]
//   - [VZVirtioFileSystemDeviceConfiguration.Tag]: A label that identifies this device in the guest VM.
//   - [VZVirtioFileSystemDeviceConfiguration.SetTag]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDeviceConfiguration
type VZVirtioFileSystemDeviceConfiguration struct {
	VZDirectorySharingDeviceConfiguration
}

// VZVirtioFileSystemDeviceConfigurationFromID constructs a [VZVirtioFileSystemDeviceConfiguration] from an objc.ID.
//
// An object that represents the configuration of a Virtio file system device.
func VZVirtioFileSystemDeviceConfigurationFromID(id objc.ID) VZVirtioFileSystemDeviceConfiguration {
	return VZVirtioFileSystemDeviceConfiguration{VZDirectorySharingDeviceConfiguration: VZDirectorySharingDeviceConfigurationFromID(id)}
}

// NOTE: VZVirtioFileSystemDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtioFileSystemDeviceConfiguration] class.
//
// # Creating a file system device configuration
//
//   - [IVZVirtioFileSystemDeviceConfiguration.InitWithTag]: Creates a configuration for a VIRTIO file system device.
//
// # Getting file system information
//
//   - [IVZVirtioFileSystemDeviceConfiguration.Share]: A value that defines how the host exposes resources to the guest virtual machine.
//   - [IVZVirtioFileSystemDeviceConfiguration.SetShare]
//   - [IVZVirtioFileSystemDeviceConfiguration.Tag]: A label that identifies this device in the guest VM.
//   - [IVZVirtioFileSystemDeviceConfiguration.SetTag]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDeviceConfiguration
type IVZVirtioFileSystemDeviceConfiguration interface {
	IVZDirectorySharingDeviceConfiguration

	// Topic: Creating a file system device configuration

	// Creates a configuration for a VIRTIO file system device.
	InitWithTag(tag string) VZVirtioFileSystemDeviceConfiguration

	// Topic: Getting file system information

	// A value that defines how the host exposes resources to the guest virtual machine.
	Share() IVZDirectoryShare
	SetShare(value IVZDirectoryShare)
	// A label that identifies this device in the guest VM.
	Tag() string
	SetTag(value string)
}

// Init initializes the instance.
func (v VZVirtioFileSystemDeviceConfiguration) Init() VZVirtioFileSystemDeviceConfiguration {
	rv := objc.Send[VZVirtioFileSystemDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioFileSystemDeviceConfiguration) Autorelease() VZVirtioFileSystemDeviceConfiguration {
	rv := objc.Send[VZVirtioFileSystemDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioFileSystemDeviceConfiguration creates a new VZVirtioFileSystemDeviceConfiguration instance.
func NewVZVirtioFileSystemDeviceConfiguration() VZVirtioFileSystemDeviceConfiguration {
	class := getVZVirtioFileSystemDeviceConfigurationClass()
	rv := objc.Send[VZVirtioFileSystemDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a configuration for a VIRTIO file system device.
//
// tag: The label identifying this device in the guest.
//
// # Discussion
//
// The system presents the `tag` as a label in the guest identifying this
// device for mounting. The `tag` must be valid, which you can check with
// [ValidateTagError].
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDeviceConfiguration/init(tag:)
func NewVirtioFileSystemDeviceConfigurationWithTag(tag string) VZVirtioFileSystemDeviceConfiguration {
	instance := getVZVirtioFileSystemDeviceConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTag:"), objc.String(tag))
	return VZVirtioFileSystemDeviceConfigurationFromID(rv)
}

// Creates a configuration for a VIRTIO file system device.
//
// tag: The label identifying this device in the guest.
//
// # Discussion
//
// The system presents the `tag` as a label in the guest identifying this
// device for mounting. The `tag` must be valid, which you can check with
// [ValidateTagError].
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDeviceConfiguration/init(tag:)
func (v VZVirtioFileSystemDeviceConfiguration) InitWithTag(tag string) VZVirtioFileSystemDeviceConfiguration {
	rv := objc.Send[VZVirtioFileSystemDeviceConfiguration](v.ID, objc.Sel("initWithTag:"), objc.String(tag))
	return rv
}

// Checks to see whether a Virtio tag is valid.
//
// tag: The tag to validate.
//
// # Discussion
//
// The tag can’t be empty and must be fewer than 36 bytes when encoded in
// UTF-8.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDeviceConfiguration/validateTag(_:)
func (_VZVirtioFileSystemDeviceConfigurationClass VZVirtioFileSystemDeviceConfigurationClass) ValidateTagError(tag string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_VZVirtioFileSystemDeviceConfigurationClass.class), objc.Sel("validateTag:error:"), objc.String(tag), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateTag:error: returned NO with nil NSError")
	}
	return rv, nil

}

// A value that defines how the host exposes resources to the guest virtual
// machine.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDeviceConfiguration/share
func (v VZVirtioFileSystemDeviceConfiguration) Share() IVZDirectoryShare {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("share"))
	return VZDirectoryShareFromID(objc.ID(rv))
}
func (v VZVirtioFileSystemDeviceConfiguration) SetShare(value IVZDirectoryShare) {
	objc.Send[struct{}](v.ID, objc.Sel("setShare:"), value)
}

// A label that identifies this device in the guest VM.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDeviceConfiguration/tag
func (v VZVirtioFileSystemDeviceConfiguration) Tag() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("tag"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZVirtioFileSystemDeviceConfiguration) SetTag(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("setTag:"), objc.String(value))
}

// A value that indicates that the guest needs to automount this file system
// device in the guest VM.
//
// # Discussion
//
// A device configured with this tag is automatically mounted in a macOS
// guest.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDeviceConfiguration/macOSGuestAutomountTag
func (_VZVirtioFileSystemDeviceConfigurationClass VZVirtioFileSystemDeviceConfigurationClass) MacOSGuestAutomountTag() string {
	rv := objc.Send[objc.ID](objc.ID(_VZVirtioFileSystemDeviceConfigurationClass.class), objc.Sel("macOSGuestAutomountTag"))
	return foundation.NSStringFromID(rv).String()
}
