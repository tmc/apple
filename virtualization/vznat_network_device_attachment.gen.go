// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZNATNetworkDeviceAttachment] class.
var (
	_VZNATNetworkDeviceAttachmentClass     VZNATNetworkDeviceAttachmentClass
	_VZNATNetworkDeviceAttachmentClassOnce sync.Once
)

func getVZNATNetworkDeviceAttachmentClass() VZNATNetworkDeviceAttachmentClass {
	_VZNATNetworkDeviceAttachmentClassOnce.Do(func() {
		_VZNATNetworkDeviceAttachmentClass = VZNATNetworkDeviceAttachmentClass{class: objc.GetClass("VZNATNetworkDeviceAttachment")}
	})
	return _VZNATNetworkDeviceAttachmentClass
}

// GetVZNATNetworkDeviceAttachmentClass returns the class object for VZNATNetworkDeviceAttachment.
func GetVZNATNetworkDeviceAttachmentClass() VZNATNetworkDeviceAttachmentClass {
	return getVZNATNetworkDeviceAttachmentClass()
}

type VZNATNetworkDeviceAttachmentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZNATNetworkDeviceAttachmentClass) Alloc() VZNATNetworkDeviceAttachment {
	rv := objc.Send[VZNATNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// A device that routes network requests through the host computer and
// performs network address translation on the resulting packets.
//
// # Overview
// 
// A [VZNATNetworkDeviceAttachment] works with the host computer to perform
// network address translation (NAT) on the guest system’s network packets,
// and then route those packets to outside networks. Use this attachment to
// give the guest system indirect access to external networks, instead of
// direct access through a shared physical network interface.
// 
// To configure a network device with a NAT attachment:
// 
// - Create the [VZNATNetworkDeviceAttachment] object. - Assign the attachment
// object to the [VZNATNetworkDeviceAttachment.Attachment] property of a
// [VZVirtioNetworkDeviceConfiguration] object. - Add the
// [VZVirtioNetworkDeviceConfiguration] object to the [VZNATNetworkDeviceAttachment.NetworkDevices]
// property of your [VZVirtualMachineConfiguration].
// 
// This attachment doesn’t require your app to have the
// [com.apple.vm.networking] entitlement.
//
// [com.apple.vm.networking]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.vm.networking
//
// See: https://developer.apple.com/documentation/Virtualization/VZNATNetworkDeviceAttachment
type VZNATNetworkDeviceAttachment struct {
	VZNetworkDeviceAttachment
}

// VZNATNetworkDeviceAttachmentFromID constructs a [VZNATNetworkDeviceAttachment] from an objc.ID.
//
// A device that routes network requests through the host computer and
// performs network address translation on the resulting packets.
func VZNATNetworkDeviceAttachmentFromID(id objc.ID) VZNATNetworkDeviceAttachment {
	return VZNATNetworkDeviceAttachment{VZNetworkDeviceAttachment: VZNetworkDeviceAttachmentFromID(id)}
}
// NOTE: VZNATNetworkDeviceAttachment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZNATNetworkDeviceAttachment] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZNATNetworkDeviceAttachment
type IVZNATNetworkDeviceAttachment interface {
	IVZNetworkDeviceAttachment

	// The object that defines how the virtual network device communicates with the host system.
	Attachment() IVZNetworkDeviceAttachment
	SetAttachment(value IVZNetworkDeviceAttachment)
	// The array of network devices that you expose to the guest operating system.
	NetworkDevices() IVZNetworkDeviceConfiguration
	SetNetworkDevices(value IVZNetworkDeviceConfiguration)
}





// Init initializes the instance.
func (n VZNATNetworkDeviceAttachment) Init() VZNATNetworkDeviceAttachment {
	rv := objc.Send[VZNATNetworkDeviceAttachment](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n VZNATNetworkDeviceAttachment) Autorelease() VZNATNetworkDeviceAttachment {
	rv := objc.Send[VZNATNetworkDeviceAttachment](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNATNetworkDeviceAttachment creates a new VZNATNetworkDeviceAttachment instance.
func NewVZNATNetworkDeviceAttachment() VZNATNetworkDeviceAttachment {
	class := getVZNATNetworkDeviceAttachmentClass()
	rv := objc.Send[VZNATNetworkDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}






















// The object that defines how the virtual network device communicates with
// the host system.
//
// See: https://developer.apple.com/documentation/virtualization/vznetworkdeviceconfiguration/attachment
func (n VZNATNetworkDeviceAttachment) Attachment() IVZNetworkDeviceAttachment {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("attachment"))
	return VZNetworkDeviceAttachmentFromID(objc.ID(rv))
}
func (n VZNATNetworkDeviceAttachment) SetAttachment(value IVZNetworkDeviceAttachment) {
	objc.Send[struct{}](n.ID, objc.Sel("setAttachment:"), value)
}



// The array of network devices that you expose to the guest operating system.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/networkdevices
func (n VZNATNetworkDeviceAttachment) NetworkDevices() IVZNetworkDeviceConfiguration {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("networkDevices"))
	return VZNetworkDeviceConfigurationFromID(objc.ID(rv))
}
func (n VZNATNetworkDeviceAttachment) SetNetworkDevices(value IVZNetworkDeviceConfiguration) {
	objc.Send[struct{}](n.ID, objc.Sel("setNetworkDevices:"), value)
}























