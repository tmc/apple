// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/vmnet"
)

// The class instance for the [VZVmnetNetworkDeviceAttachment] class.
var (
	_VZVmnetNetworkDeviceAttachmentClass     VZVmnetNetworkDeviceAttachmentClass
	_VZVmnetNetworkDeviceAttachmentClassOnce sync.Once
)

func getVZVmnetNetworkDeviceAttachmentClass() VZVmnetNetworkDeviceAttachmentClass {
	_VZVmnetNetworkDeviceAttachmentClassOnce.Do(func() {
		_VZVmnetNetworkDeviceAttachmentClass = VZVmnetNetworkDeviceAttachmentClass{class: objc.GetClass("VZVmnetNetworkDeviceAttachment")}
	})
	return _VZVmnetNetworkDeviceAttachmentClass
}

// GetVZVmnetNetworkDeviceAttachmentClass returns the class object for VZVmnetNetworkDeviceAttachment.
func GetVZVmnetNetworkDeviceAttachmentClass() VZVmnetNetworkDeviceAttachmentClass {
	return getVZVmnetNetworkDeviceAttachmentClass()
}

type VZVmnetNetworkDeviceAttachmentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVmnetNetworkDeviceAttachmentClass) Alloc() VZVmnetNetworkDeviceAttachment {
	rv := objc.Send[VZVmnetNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// A network device attachment that allows a custom network topology.
//
// # Overview
// 
// The Virtualization framework backs this attachment by a logical network
// which the client creates and customizes through the [vmnet] framework APIs
// to allow custom network topology which allows multiple virtual machines to
// appear on the same network and connect with each other.
//
// [vmnet]: https://developer.apple.com/documentation/vmnet
//
// # Creating the vmnet network device attachment
//
//   - [VZVmnetNetworkDeviceAttachment.InitWithNetwork]: Creates the attachment and configures it with the specified data.
//   - [VZVmnetNetworkDeviceAttachment.Network]: The network object that the you initialize the attachment with.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVmnetNetworkDeviceAttachment
type VZVmnetNetworkDeviceAttachment struct {
	VZNetworkDeviceAttachment
}

// VZVmnetNetworkDeviceAttachmentFromID constructs a [VZVmnetNetworkDeviceAttachment] from an objc.ID.
//
// A network device attachment that allows a custom network topology.
func VZVmnetNetworkDeviceAttachmentFromID(id objc.ID) VZVmnetNetworkDeviceAttachment {
	return VZVmnetNetworkDeviceAttachment{VZNetworkDeviceAttachment: VZNetworkDeviceAttachmentFromID(id)}
}
// NOTE: VZVmnetNetworkDeviceAttachment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZVmnetNetworkDeviceAttachment] class.
//
// # Creating the vmnet network device attachment
//
//   - [IVZVmnetNetworkDeviceAttachment.InitWithNetwork]: Creates the attachment and configures it with the specified data.
//   - [IVZVmnetNetworkDeviceAttachment.Network]: The network object that the you initialize the attachment with.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVmnetNetworkDeviceAttachment
type IVZVmnetNetworkDeviceAttachment interface {
	IVZNetworkDeviceAttachment

	// Topic: Creating the vmnet network device attachment

	// Creates the attachment and configures it with the specified data.
	InitWithNetwork(network vmnet.Vmnet_network_ref) VZVmnetNetworkDeviceAttachment
	// The network object that the you initialize the attachment with.
	Network() unsafe.Pointer
}





// Init initializes the instance.
func (v VZVmnetNetworkDeviceAttachment) Init() VZVmnetNetworkDeviceAttachment {
	rv := objc.Send[VZVmnetNetworkDeviceAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVmnetNetworkDeviceAttachment) Autorelease() VZVmnetNetworkDeviceAttachment {
	rv := objc.Send[VZVmnetNetworkDeviceAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVmnetNetworkDeviceAttachment creates a new VZVmnetNetworkDeviceAttachment instance.
func NewVZVmnetNetworkDeviceAttachment() VZVmnetNetworkDeviceAttachment {
	class := getVZVmnetNetworkDeviceAttachmentClass()
	rv := objc.Send[VZVmnetNetworkDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates the attachment and configures it with the specified data.
//
// network: The logical network object
//
// # Return Value
// 
// An initialized [VZVmnetNetworkDeviceAttachment] object, or `nil` if there
// was an error.
//
// # Discussion
// 
// To ensure proper isolation between application processes, a virtual machine
// (VM) can only use the `network` that the same application process creates.
// If an application’s VM tries to use a `network` that another
// application’s VM creates, initialization fails.
// 
// For more information on vmnet configuration requirements and restrictions,
// see [vmnet]
// 
// The following example demonstrates how to create and initialize a custom
// network using [VZVmnetNetworkDeviceAttachment].
//
// [vmnet]: https://developer.apple.com/documentation/vmnet
//
// See: https://developer.apple.com/documentation/Virtualization/VZVmnetNetworkDeviceAttachment/init(network:)
func NewVmnetNetworkDeviceAttachmentWithNetwork(network vmnet.Vmnet_network_ref) VZVmnetNetworkDeviceAttachment {
	instance := getVZVmnetNetworkDeviceAttachmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetwork:"), network)
	return VZVmnetNetworkDeviceAttachmentFromID(rv)
}







// Creates the attachment and configures it with the specified data.
//
// network: The logical network object
//
// # Return Value
// 
// An initialized [VZVmnetNetworkDeviceAttachment] object, or `nil` if there
// was an error.
//
// # Discussion
// 
// To ensure proper isolation between application processes, a virtual machine
// (VM) can only use the `network` that the same application process creates.
// If an application’s VM tries to use a `network` that another
// application’s VM creates, initialization fails.
// 
// For more information on vmnet configuration requirements and restrictions,
// see [vmnet]
// 
// The following example demonstrates how to create and initialize a custom
// network using [VZVmnetNetworkDeviceAttachment].
//
// [vmnet]: https://developer.apple.com/documentation/vmnet
//
// See: https://developer.apple.com/documentation/Virtualization/VZVmnetNetworkDeviceAttachment/init(network:)
func (v VZVmnetNetworkDeviceAttachment) InitWithNetwork(network vmnet.Vmnet_network_ref) VZVmnetNetworkDeviceAttachment {
	rv := objc.Send[VZVmnetNetworkDeviceAttachment](v.ID, objc.Sel("initWithNetwork:"), network)
	return rv
}












// The network object that the you initialize the attachment with.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVmnetNetworkDeviceAttachment/network
func (v VZVmnetNetworkDeviceAttachment) Network() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("network"))
	return rv
}























