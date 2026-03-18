// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VZVirtioFileSystemDevice] class.
var (
	_VZVirtioFileSystemDeviceClass     VZVirtioFileSystemDeviceClass
	_VZVirtioFileSystemDeviceClassOnce sync.Once
)

func getVZVirtioFileSystemDeviceClass() VZVirtioFileSystemDeviceClass {
	_VZVirtioFileSystemDeviceClassOnce.Do(func() {
		_VZVirtioFileSystemDeviceClass = VZVirtioFileSystemDeviceClass{class: objc.GetClass("VZVirtioFileSystemDevice")}
	})
	return _VZVirtioFileSystemDeviceClass
}

// GetVZVirtioFileSystemDeviceClass returns the class object for VZVirtioFileSystemDevice.
func GetVZVirtioFileSystemDeviceClass() VZVirtioFileSystemDeviceClass {
	return getVZVirtioFileSystemDeviceClass()
}

type VZVirtioFileSystemDeviceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioFileSystemDeviceClass) Alloc() VZVirtioFileSystemDevice {
	rv := objc.Send[VZVirtioFileSystemDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object the defines a VIRTIO file system device.
//
// # Overview
// 
// This device exposes host resources to the guest as a file system mount. The
// directory share defines which resources the host exposes to the guest.
// 
// Create this device by instantiating a
// [VZVirtioFileSystemDeviceConfiguration] in a
// [VZVirtualMachineConfiguration]. The file system device is available in the
// [VZVirtualMachine].[VZVirtioFileSystemDevice.DirectorySharingDevices] property. The guest can use
// the [VZVirtioFileSystemDevice.Tag] label to mount and access the host resources.
// 
// With [VZVirtioFileSystemDevice], the framework enforces several permissions
// policies for shared directories:
// 
// - The framework reads and writes files using the user ID (UID) of the
// effective user, which is the UID of the current user, rather than the UID
// of the system process. - The framework doesn’t allow reading or
// overwriting of files with permissions where the file is inaccessible to the
// current user. - The framework ignores requests from guest operating systems
// to change the UID or group ID (GID) of files on the host.
//
// # Accessing directory properties
//
//   - [VZVirtioFileSystemDevice.Share]: A value that defines the directory share the host exposes to the guest VM.
//   - [VZVirtioFileSystemDevice.SetShare]
//   - [VZVirtioFileSystemDevice.Tag]: A string that identifies the device.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDevice
type VZVirtioFileSystemDevice struct {
	VZDirectorySharingDevice
}

// VZVirtioFileSystemDeviceFromID constructs a [VZVirtioFileSystemDevice] from an objc.ID.
//
// An object the defines a VIRTIO file system device.
func VZVirtioFileSystemDeviceFromID(id objc.ID) VZVirtioFileSystemDevice {
	return VZVirtioFileSystemDevice{VZDirectorySharingDevice: VZDirectorySharingDeviceFromID(id)}
}
// NOTE: VZVirtioFileSystemDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtioFileSystemDevice] class.
//
// # Accessing directory properties
//
//   - [IVZVirtioFileSystemDevice.Share]: A value that defines the directory share the host exposes to the guest VM.
//   - [IVZVirtioFileSystemDevice.SetShare]
//   - [IVZVirtioFileSystemDevice.Tag]: A string that identifies the device.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDevice
type IVZVirtioFileSystemDevice interface {
	IVZDirectorySharingDevice

	// Topic: Accessing directory properties

	// A value that defines the directory share the host exposes to the guest VM.
	Share() IVZDirectoryShare
	SetShare(value IVZDirectoryShare)
	// A string that identifies the device.
	Tag() string

	// The list of configured directory-sharing devices on the VM.
	DirectorySharingDevices() IVZDirectorySharingDevice
	SetDirectorySharingDevices(value IVZDirectorySharingDevice)
}

// Init initializes the instance.
func (v VZVirtioFileSystemDevice) Init() VZVirtioFileSystemDevice {
	rv := objc.Send[VZVirtioFileSystemDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioFileSystemDevice) Autorelease() VZVirtioFileSystemDevice {
	rv := objc.Send[VZVirtioFileSystemDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioFileSystemDevice creates a new VZVirtioFileSystemDevice instance.
func NewVZVirtioFileSystemDevice() VZVirtioFileSystemDevice {
	class := getVZVirtioFileSystemDeviceClass()
	rv := objc.Send[VZVirtioFileSystemDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A value that defines the directory share the host exposes to the guest VM.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDevice/share
func (v VZVirtioFileSystemDevice) Share() IVZDirectoryShare {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("share"))
	return VZDirectoryShareFromID(objc.ID(rv))
}
func (v VZVirtioFileSystemDevice) SetShare(value IVZDirectoryShare) {
	objc.Send[struct{}](v.ID, objc.Sel("setShare:"), value)
}

// A string that identifies the device.
//
// # Discussion
// 
// The system presents the `tag` as a label in the guest VM that identifies
// this device that’s available for mounting.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDevice/tag
func (v VZVirtioFileSystemDevice) Tag() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("tag"))
	return foundation.NSStringFromID(rv).String()
}

// The list of configured directory-sharing devices on the VM.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/directorysharingdevices
func (v VZVirtioFileSystemDevice) DirectorySharingDevices() IVZDirectorySharingDevice {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("directorySharingDevices"))
	return VZDirectorySharingDeviceFromID(objc.ID(rv))
}
func (v VZVirtioFileSystemDevice) SetDirectorySharingDevices(value IVZDirectorySharingDevice) {
	objc.Send[struct{}](v.ID, objc.Sel("setDirectorySharingDevices:"), value)
}

