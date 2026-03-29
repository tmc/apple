// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZStorageDeviceConfiguration] class.
var (
	_VZStorageDeviceConfigurationClass     VZStorageDeviceConfigurationClass
	_VZStorageDeviceConfigurationClassOnce sync.Once
)

func getVZStorageDeviceConfigurationClass() VZStorageDeviceConfigurationClass {
	_VZStorageDeviceConfigurationClassOnce.Do(func() {
		_VZStorageDeviceConfigurationClass = VZStorageDeviceConfigurationClass{class: objc.GetClass("VZStorageDeviceConfiguration")}
	})
	return _VZStorageDeviceConfigurationClass
}

// GetVZStorageDeviceConfigurationClass returns the class object for VZStorageDeviceConfiguration.
func GetVZStorageDeviceConfigurationClass() VZStorageDeviceConfigurationClass {
	return getVZStorageDeviceConfigurationClass()
}

type VZStorageDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZStorageDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZStorageDeviceConfigurationClass) Alloc() VZStorageDeviceConfiguration {
	rv := objc.Send[VZStorageDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The common configuration traits for storage device requests.
//
// # Overview
// 
// Don’t create a [VZStorageDeviceConfiguration] object directly. Instead,
// instantiate one of its subclasses, such as
// [VZVirtioBlockDeviceConfiguration]. Use the [VZStorageDeviceConfiguration.Attachment] property of this
// class to access the device’s underlying storage.
//
// # Getting the attachment point
//
//   - [VZStorageDeviceConfiguration.Attachment]: The attachment object that provides the underlying storage for the device.
//
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceConfiguration
type VZStorageDeviceConfiguration struct {
	objectivec.Object
}

// VZStorageDeviceConfigurationFromID constructs a [VZStorageDeviceConfiguration] from an objc.ID.
//
// The common configuration traits for storage device requests.
func VZStorageDeviceConfigurationFromID(id objc.ID) VZStorageDeviceConfiguration {
	return VZStorageDeviceConfiguration{objectivec.Object{ID: id}}
}
// NOTE: VZStorageDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZStorageDeviceConfiguration] class.
//
// # Getting the attachment point
//
//   - [IVZStorageDeviceConfiguration.Attachment]: The attachment object that provides the underlying storage for the device.
//
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceConfiguration
type IVZStorageDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Getting the attachment point

	// The attachment object that provides the underlying storage for the device.
	Attachment() IVZStorageDeviceAttachment
}

// Init initializes the instance.
func (s VZStorageDeviceConfiguration) Init() VZStorageDeviceConfiguration {
	rv := objc.Send[VZStorageDeviceConfiguration](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s VZStorageDeviceConfiguration) Autorelease() VZStorageDeviceConfiguration {
	rv := objc.Send[VZStorageDeviceConfiguration](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZStorageDeviceConfiguration creates a new VZStorageDeviceConfiguration instance.
func NewVZStorageDeviceConfiguration() VZStorageDeviceConfiguration {
	class := getVZStorageDeviceConfigurationClass()
	rv := objc.Send[VZStorageDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The attachment object that provides the underlying storage for the device.
//
// # Discussion
// 
// The attachment object defines which local resource on the host computer
// appears as a disk in the virtual machine.
//
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceConfiguration/attachment
func (s VZStorageDeviceConfiguration) Attachment() IVZStorageDeviceAttachment {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("attachment"))
	return VZStorageDeviceAttachmentFromID(objc.ID(rv))
}

