// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZBridgedNetworkDeviceAttachment] class.
var (
	_VZBridgedNetworkDeviceAttachmentClass     VZBridgedNetworkDeviceAttachmentClass
	_VZBridgedNetworkDeviceAttachmentClassOnce sync.Once
)

func getVZBridgedNetworkDeviceAttachmentClass() VZBridgedNetworkDeviceAttachmentClass {
	_VZBridgedNetworkDeviceAttachmentClassOnce.Do(func() {
		_VZBridgedNetworkDeviceAttachmentClass = VZBridgedNetworkDeviceAttachmentClass{class: objc.GetClass("VZBridgedNetworkDeviceAttachment")}
	})
	return _VZBridgedNetworkDeviceAttachmentClass
}

// GetVZBridgedNetworkDeviceAttachmentClass returns the class object for VZBridgedNetworkDeviceAttachment.
func GetVZBridgedNetworkDeviceAttachmentClass() VZBridgedNetworkDeviceAttachmentClass {
	return getVZBridgedNetworkDeviceAttachmentClass()
}

type VZBridgedNetworkDeviceAttachmentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZBridgedNetworkDeviceAttachmentClass) Alloc() VZBridgedNetworkDeviceAttachment {
	rv := objc.Send[VZBridgedNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// A network device that interacts directly with a physical network interface
// on the host computer.
//
// # Overview
// 
// A [VZBridgedNetworkDeviceAttachment] object represents a physical interface
// on the host computer. Use this object when configuring a network interface
// for your virtual machine. A bridged network device sends and receives
// packets on the same physical interface as the host computer, but does so
// using a different network layer.
// 
// To configure a network device with a bridged network interface:
// 
// - Obtain a reference to one of the host’s physical network interfaces
// from the [VZBridgedNetworkDeviceAttachment.NetworkInterfaces] property of [VZBridgedNetworkInterface]. -
// Create the [VZBridgedNetworkDeviceAttachment] object using the network
// interface. - Assign the attachment object to the [VZBridgedNetworkDeviceAttachment.Attachment] property of a
// [VZVirtioNetworkDeviceConfiguration] object. - Add the
// [VZVirtioNetworkDeviceConfiguration] object to the [VZBridgedNetworkDeviceAttachment.NetworkDevices]
// property of your [VZVirtualMachineConfiguration].
//
// # Creating the attachment point
//
//   - [VZBridgedNetworkDeviceAttachment.InitWithInterface]: Creates the attachment from a bridged network interface object.
//
// # Getting the network interface
//
//   - [VZBridgedNetworkDeviceAttachment.Interface]: The network interface assigned to this attachment.
//
// See: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkDeviceAttachment
type VZBridgedNetworkDeviceAttachment struct {
	VZNetworkDeviceAttachment
}

// VZBridgedNetworkDeviceAttachmentFromID constructs a [VZBridgedNetworkDeviceAttachment] from an objc.ID.
//
// A network device that interacts directly with a physical network interface
// on the host computer.
func VZBridgedNetworkDeviceAttachmentFromID(id objc.ID) VZBridgedNetworkDeviceAttachment {
	return VZBridgedNetworkDeviceAttachment{VZNetworkDeviceAttachment: VZNetworkDeviceAttachmentFromID(id)}
}
// NOTE: VZBridgedNetworkDeviceAttachment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZBridgedNetworkDeviceAttachment] class.
//
// # Creating the attachment point
//
//   - [IVZBridgedNetworkDeviceAttachment.InitWithInterface]: Creates the attachment from a bridged network interface object.
//
// # Getting the network interface
//
//   - [IVZBridgedNetworkDeviceAttachment.Interface]: The network interface assigned to this attachment.
//
// See: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkDeviceAttachment
type IVZBridgedNetworkDeviceAttachment interface {
	IVZNetworkDeviceAttachment

	// Topic: Creating the attachment point

	// Creates the attachment from a bridged network interface object.
	InitWithInterface(interface_ IVZBridgedNetworkInterface) VZBridgedNetworkDeviceAttachment

	// Topic: Getting the network interface

	// The network interface assigned to this attachment.
	Interface() IVZBridgedNetworkInterface

	// The object that defines how the virtual network device communicates with the host system.
	Attachment() IVZNetworkDeviceAttachment
	SetAttachment(value IVZNetworkDeviceAttachment)
	// The array of network devices that you expose to the guest operating system.
	NetworkDevices() IVZNetworkDeviceConfiguration
	SetNetworkDevices(value IVZNetworkDeviceConfiguration)
}





// Init initializes the instance.
func (b VZBridgedNetworkDeviceAttachment) Init() VZBridgedNetworkDeviceAttachment {
	rv := objc.Send[VZBridgedNetworkDeviceAttachment](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b VZBridgedNetworkDeviceAttachment) Autorelease() VZBridgedNetworkDeviceAttachment {
	rv := objc.Send[VZBridgedNetworkDeviceAttachment](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZBridgedNetworkDeviceAttachment creates a new VZBridgedNetworkDeviceAttachment instance.
func NewVZBridgedNetworkDeviceAttachment() VZBridgedNetworkDeviceAttachment {
	class := getVZBridgedNetworkDeviceAttachmentClass()
	rv := objc.Send[VZBridgedNetworkDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates the attachment from a bridged network interface object.
//
// interface: An existing network interface of the host computer. Get a list of available
// interfaces from the [NetworkInterfaces] property of
// [VZBridgedNetworkInterface].
//
// # Return Value
// 
// An attachment object for the specified network interface.
//
// See: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkDeviceAttachment/init(interface:)
func NewBridgedNetworkDeviceAttachmentWithInterface(interface_ IVZBridgedNetworkInterface) VZBridgedNetworkDeviceAttachment {
	instance := getVZBridgedNetworkDeviceAttachmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInterface:"), interface_)
	return VZBridgedNetworkDeviceAttachmentFromID(rv)
}







// Creates the attachment from a bridged network interface object.
//
// interface: An existing network interface of the host computer. Get a list of available
// interfaces from the [NetworkInterfaces] property of
// [VZBridgedNetworkInterface].
//
// # Return Value
// 
// An attachment object for the specified network interface.
//
// See: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkDeviceAttachment/init(interface:)
func (b VZBridgedNetworkDeviceAttachment) InitWithInterface(interface_ IVZBridgedNetworkInterface) VZBridgedNetworkDeviceAttachment {
	rv := objc.Send[VZBridgedNetworkDeviceAttachment](b.ID, objc.Sel("initWithInterface:"), interface_)
	return rv
}












// The network interface assigned to this attachment.
//
// See: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkDeviceAttachment/interface
func (b VZBridgedNetworkDeviceAttachment) Interface() IVZBridgedNetworkInterface {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("interface"))
	return VZBridgedNetworkInterfaceFromID(objc.ID(rv))
}



// The object that defines how the virtual network device communicates with
// the host system.
//
// See: https://developer.apple.com/documentation/virtualization/vznetworkdeviceconfiguration/attachment
func (b VZBridgedNetworkDeviceAttachment) Attachment() IVZNetworkDeviceAttachment {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("attachment"))
	return VZNetworkDeviceAttachmentFromID(objc.ID(rv))
}
func (b VZBridgedNetworkDeviceAttachment) SetAttachment(value IVZNetworkDeviceAttachment) {
	objc.Send[struct{}](b.ID, objc.Sel("setAttachment:"), value)
}



// The array of network devices that you expose to the guest operating system.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/networkdevices
func (b VZBridgedNetworkDeviceAttachment) NetworkDevices() IVZNetworkDeviceConfiguration {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("networkDevices"))
	return VZNetworkDeviceConfigurationFromID(objc.ID(rv))
}
func (b VZBridgedNetworkDeviceAttachment) SetNetworkDevices(value IVZNetworkDeviceConfiguration) {
	objc.Send[struct{}](b.ID, objc.Sel("setNetworkDevices:"), value)
}







// The bridged network interfaces that you may use in your virtual machine.
//
// See: https://developer.apple.com/documentation/virtualization/vzbridgednetworkinterface/networkinterfaces
func (_VZBridgedNetworkDeviceAttachmentClass VZBridgedNetworkDeviceAttachmentClass) NetworkInterfaces() VZBridgedNetworkInterface {
	rv := objc.Send[objc.ID](objc.ID(_VZBridgedNetworkDeviceAttachmentClass.class), objc.Sel("networkInterfaces"))
	return VZBridgedNetworkInterfaceFromID(objc.ID(rv))
}
func (_VZBridgedNetworkDeviceAttachmentClass VZBridgedNetworkDeviceAttachmentClass) SetNetworkInterfaces(value VZBridgedNetworkInterface) {
	objc.Send[struct{}](objc.ID(_VZBridgedNetworkDeviceAttachmentClass.class), objc.Sel("setNetworkInterfaces:"), value)
}



















