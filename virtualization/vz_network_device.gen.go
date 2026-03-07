// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZNetworkDevice] class.
var (
	_VZNetworkDeviceClass     VZNetworkDeviceClass
	_VZNetworkDeviceClassOnce sync.Once
)

func getVZNetworkDeviceClass() VZNetworkDeviceClass {
	_VZNetworkDeviceClassOnce.Do(func() {
		_VZNetworkDeviceClass = VZNetworkDeviceClass{class: objc.GetClass("VZNetworkDevice")}
	})
	return _VZNetworkDeviceClass
}

// GetVZNetworkDeviceClass returns the class object for VZNetworkDevice.
func GetVZNetworkDeviceClass() VZNetworkDeviceClass {
	return getVZNetworkDeviceClass()
}

type VZNetworkDeviceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZNetworkDeviceClass) Alloc() VZNetworkDevice {
	rv := objc.Send[VZNetworkDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// A base class that represents a network device in a virtual machine.
//
// # Overview
// 
// Don’t instantiate a [VZNetworkDevice] directly. When you create a
// [VZVirtualMachineConfiguration] instance with a
// [VZVirtualMachineConfiguration] the system creates the number of network
// devices based on the number of [VZVirtioNetworkDeviceConfiguration] objects
// you specify in the VM configuration. Before initializing the virtual
// machine (VM), validate the configuration using [ValidateWithError] to
// ensure the user’s computer supports the number of network and other
// devices you’ve specified.
// 
// For many purposes, a single network that uses a Network Address Translation
// (NAT) attachment and connects the VM to the host computer’s network is
// sufficient. You can use additional network interfaces for purposes of your
// own design, such as:
// 
// - Bridging several physical interfaces to connect to multiple networks. -
// Using the file descriptor attachment to create specialized connections for
// different purposes.
// 
// You access the network devices through the
// [VZVirtualMachine].[VZNetworkDevice.NetworkDevices] property. The network devices map to
// their respective configurations in a one to one relationship, where index
// `i` of `VZVirtualMachine.NetworkDevices()` corresponds to the network
// device configuration at index `i` set on
// [VZVirtualMachineConfiguration].[VZNetworkDevice.NetworkDevices].
//
// # Getting the network attachment point
//
//   - [VZNetworkDevice.Attachment]: The network attachment that’s connected to this network device.
//   - [VZNetworkDevice.SetAttachment]
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDevice
type VZNetworkDevice struct {
	objectivec.Object
}

// VZNetworkDeviceFromID constructs a [VZNetworkDevice] from an objc.ID.
//
// A base class that represents a network device in a virtual machine.
func VZNetworkDeviceFromID(id objc.ID) VZNetworkDevice {
	return VZNetworkDevice{objectivec.Object{id}}
}
// NOTE: VZNetworkDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZNetworkDevice] class.
//
// # Getting the network attachment point
//
//   - [IVZNetworkDevice.Attachment]: The network attachment that’s connected to this network device.
//   - [IVZNetworkDevice.SetAttachment]
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDevice
type IVZNetworkDevice interface {
	objectivec.IObject

	// Topic: Getting the network attachment point

	// The network attachment that’s connected to this network device.
	Attachment() IVZNetworkDeviceAttachment
	SetAttachment(value IVZNetworkDeviceAttachment)

	// The list of configured network devices on the VM.
	NetworkDevices() IVZNetworkDevice
	SetNetworkDevices(value IVZNetworkDevice)
}





// Init initializes the instance.
func (n VZNetworkDevice) Init() VZNetworkDevice {
	rv := objc.Send[VZNetworkDevice](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n VZNetworkDevice) Autorelease() VZNetworkDevice {
	rv := objc.Send[VZNetworkDevice](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNetworkDevice creates a new VZNetworkDevice instance.
func NewVZNetworkDevice() VZNetworkDevice {
	class := getVZNetworkDeviceClass()
	rv := objc.Send[VZNetworkDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// The network attachment that’s connected to this network device.
//
// # Discussion
// 
// Setting this property results in an attempt to change the network device
// attachment which may fail. If the devices fails to attach, the system
// invokes [VirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError] and
// sets this property to `nil`. This property may change at any time while the
// VM is running based on the state of the host network.
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDevice/attachment
func (n VZNetworkDevice) Attachment() IVZNetworkDeviceAttachment {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("attachment"))
	return VZNetworkDeviceAttachmentFromID(objc.ID(rv))
}
func (n VZNetworkDevice) SetAttachment(value IVZNetworkDeviceAttachment) {
	objc.Send[struct{}](n.ID, objc.Sel("setAttachment:"), value)
}



// The list of configured network devices on the VM.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/networkdevices
func (n VZNetworkDevice) NetworkDevices() IVZNetworkDevice {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("networkDevices"))
	return VZNetworkDeviceFromID(objc.ID(rv))
}
func (n VZNetworkDevice) SetNetworkDevices(value IVZNetworkDevice) {
	objc.Send[struct{}](n.ID, objc.Sel("setNetworkDevices:"), value)
}























