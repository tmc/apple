// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZNetworkDeviceConfiguration] class.
var (
	_VZNetworkDeviceConfigurationClass     VZNetworkDeviceConfigurationClass
	_VZNetworkDeviceConfigurationClassOnce sync.Once
)

func getVZNetworkDeviceConfigurationClass() VZNetworkDeviceConfigurationClass {
	_VZNetworkDeviceConfigurationClassOnce.Do(func() {
		_VZNetworkDeviceConfigurationClass = VZNetworkDeviceConfigurationClass{class: objc.GetClass("VZNetworkDeviceConfiguration")}
	})
	return _VZNetworkDeviceConfigurationClass
}

// GetVZNetworkDeviceConfigurationClass returns the class object for VZNetworkDeviceConfiguration.
func GetVZNetworkDeviceConfigurationClass() VZNetworkDeviceConfigurationClass {
	return getVZNetworkDeviceConfigurationClass()
}

type VZNetworkDeviceConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZNetworkDeviceConfigurationClass) Alloc() VZNetworkDeviceConfiguration {
	rv := objc.Send[VZNetworkDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The common configuration traits for network devices.
//
// # Overview
// 
// Don’t instantiate the [VZNetworkDeviceConfiguration] class directly.
// Instead, instantiate one of its subclasses, such as
// [VZVirtioNetworkDeviceConfiguration]. Then use the properties of this class
// to configure the network device.
//
// # Setting configuration attributes
//
//   - [VZNetworkDeviceConfiguration.Attachment]: The object that defines how the virtual network device communicates with the host system.
//   - [VZNetworkDeviceConfiguration.SetAttachment]
//   - [VZNetworkDeviceConfiguration.MACAddress]: The media access control (MAC) address to assign to the network device.
//   - [VZNetworkDeviceConfiguration.SetMACAddress]
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration
type VZNetworkDeviceConfiguration struct {
	objectivec.Object
}

// VZNetworkDeviceConfigurationFromID constructs a [VZNetworkDeviceConfiguration] from an objc.ID.
//
// The common configuration traits for network devices.
func VZNetworkDeviceConfigurationFromID(id objc.ID) VZNetworkDeviceConfiguration {
	return VZNetworkDeviceConfiguration{objectivec.Object{ID: id}}
}
// NOTE: VZNetworkDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZNetworkDeviceConfiguration] class.
//
// # Setting configuration attributes
//
//   - [IVZNetworkDeviceConfiguration.Attachment]: The object that defines how the virtual network device communicates with the host system.
//   - [IVZNetworkDeviceConfiguration.SetAttachment]
//   - [IVZNetworkDeviceConfiguration.MACAddress]: The media access control (MAC) address to assign to the network device.
//   - [IVZNetworkDeviceConfiguration.SetMACAddress]
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration
type IVZNetworkDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Setting configuration attributes

	// The object that defines how the virtual network device communicates with the host system.
	Attachment() IVZNetworkDeviceAttachment
	SetAttachment(value IVZNetworkDeviceAttachment)
	// The media access control (MAC) address to assign to the network device.
	MACAddress() IVZMACAddress
	SetMACAddress(value IVZMACAddress)
}

// Init initializes the instance.
func (n VZNetworkDeviceConfiguration) Init() VZNetworkDeviceConfiguration {
	rv := objc.Send[VZNetworkDeviceConfiguration](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n VZNetworkDeviceConfiguration) Autorelease() VZNetworkDeviceConfiguration {
	rv := objc.Send[VZNetworkDeviceConfiguration](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNetworkDeviceConfiguration creates a new VZNetworkDeviceConfiguration instance.
func NewVZNetworkDeviceConfiguration() VZNetworkDeviceConfiguration {
	class := getVZNetworkDeviceConfigurationClass()
	rv := objc.Send[VZNetworkDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The object that defines how the virtual network device communicates with
// the host system.
//
// # Discussion
// 
// The default value of this property is `nil`. Assign an appropriate value to
// specify the type of network interface you want to make available to the
// guest operating system. For example, assign a
// [VZBridgedNetworkDeviceAttachment] object to grant access to one of the
// host computer’s physical network interfaces.
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration/attachment
func (n VZNetworkDeviceConfiguration) Attachment() IVZNetworkDeviceAttachment {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("attachment"))
	return VZNetworkDeviceAttachmentFromID(objc.ID(rv))
}
func (n VZNetworkDeviceConfiguration) SetAttachment(value IVZNetworkDeviceAttachment) {
	objc.Send[struct{}](n.ID, objc.Sel("setAttachment:"), value)
}

// The media access control (MAC) address to assign to the network device.
//
// # Discussion
// 
// The default value of this property is a random, locally administered,
// unicast address. Assign a custom value to this property if you want your
// network device to have a specific MAC address.
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration/macAddress
func (n VZNetworkDeviceConfiguration) MACAddress() IVZMACAddress {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("MACAddress"))
	return VZMACAddressFromID(objc.ID(rv))
}
func (n VZNetworkDeviceConfiguration) SetMACAddress(value IVZMACAddress) {
	objc.Send[struct{}](n.ID, objc.Sel("setMACAddress:"), value)
}

