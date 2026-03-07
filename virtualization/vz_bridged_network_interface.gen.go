// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZBridgedNetworkInterface] class.
var (
	_VZBridgedNetworkInterfaceClass     VZBridgedNetworkInterfaceClass
	_VZBridgedNetworkInterfaceClassOnce sync.Once
)

func getVZBridgedNetworkInterfaceClass() VZBridgedNetworkInterfaceClass {
	_VZBridgedNetworkInterfaceClassOnce.Do(func() {
		_VZBridgedNetworkInterfaceClass = VZBridgedNetworkInterfaceClass{class: objc.GetClass("VZBridgedNetworkInterface")}
	})
	return _VZBridgedNetworkInterfaceClass
}

// GetVZBridgedNetworkInterfaceClass returns the class object for VZBridgedNetworkInterface.
func GetVZBridgedNetworkInterfaceClass() VZBridgedNetworkInterfaceClass {
	return getVZBridgedNetworkInterfaceClass()
}

type VZBridgedNetworkInterfaceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZBridgedNetworkInterfaceClass) Alloc() VZBridgedNetworkInterface {
	rv := objc.Send[VZBridgedNetworkInterface](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// An object that identifies the supported network interfaces of the host
// computer.
//
// # Overview
// 
// Use a [VZBridgedNetworkInterface] object to retrieve the physical
// interfaces on the host computer. Use a bridged network interface to create
// a [VZBridgedNetworkDeviceAttachment] object, which maps that interface to
// one of your virtual machine’s network devices. The host computer and your
// virtual machine share access to the physical network interface, but
// communicate over it using distinct network layers.
// 
// You don’t create [VZBridgedNetworkInterface] objects directly. Instead,
// the system creates one object for each physical interface of the host
// computer and stores those objects in the [VZBridgedNetworkInterface.NetworkInterfaces] property.
// Iterate over the objects in that property to retrieve the network
// interfaces you need.
//
// # Getting the interface description
//
//   - [VZBridgedNetworkInterface.Identifier]: The unique BSD name of this network interface.
//   - [VZBridgedNetworkInterface.LocalizedDisplayName]: A user-visible name for the network interface.
//
// See: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkInterface
type VZBridgedNetworkInterface struct {
	objectivec.Object
}

// VZBridgedNetworkInterfaceFromID constructs a [VZBridgedNetworkInterface] from an objc.ID.
//
// An object that identifies the supported network interfaces of the host
// computer.
func VZBridgedNetworkInterfaceFromID(id objc.ID) VZBridgedNetworkInterface {
	return VZBridgedNetworkInterface{objectivec.Object{id}}
}
// NOTE: VZBridgedNetworkInterface adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZBridgedNetworkInterface] class.
//
// # Getting the interface description
//
//   - [IVZBridgedNetworkInterface.Identifier]: The unique BSD name of this network interface.
//   - [IVZBridgedNetworkInterface.LocalizedDisplayName]: A user-visible name for the network interface.
//
// See: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkInterface
type IVZBridgedNetworkInterface interface {
	objectivec.IObject

	// Topic: Getting the interface description

	// The unique BSD name of this network interface.
	Identifier() string
	// A user-visible name for the network interface.
	LocalizedDisplayName() string
}





// Init initializes the instance.
func (b VZBridgedNetworkInterface) Init() VZBridgedNetworkInterface {
	rv := objc.Send[VZBridgedNetworkInterface](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b VZBridgedNetworkInterface) Autorelease() VZBridgedNetworkInterface {
	rv := objc.Send[VZBridgedNetworkInterface](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZBridgedNetworkInterface creates a new VZBridgedNetworkInterface instance.
func NewVZBridgedNetworkInterface() VZBridgedNetworkInterface {
	class := getVZBridgedNetworkInterfaceClass()
	rv := objc.Send[VZBridgedNetworkInterface](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// The unique BSD name of this network interface.
//
// # Discussion
// 
// BSD names for the host computer’s Ethernet interfaces include `en0`,
// `en1`, and so on.
//
// See: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkInterface/identifier
func (b VZBridgedNetworkInterface) Identifier() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}



// A user-visible name for the network interface.
//
// # Discussion
// 
// An example interface name is [Ethernet]. Use this string when you need to
// display the name of the interface to the user.
//
// See: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkInterface/localizedDisplayName
func (b VZBridgedNetworkInterface) LocalizedDisplayName() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("localizedDisplayName"))
	return foundation.NSStringFromID(rv).String()
}







// The bridged network interfaces that you may use in your virtual machine.
//
// # Discussion
// 
// The system creates the objects in this property based on the available
// interfaces in the host machine.
//
// See: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkInterface/networkInterfaces
func (_VZBridgedNetworkInterfaceClass VZBridgedNetworkInterfaceClass) NetworkInterfaces() []VZBridgedNetworkInterface {
	rv := objc.Send[[]objc.ID](objc.ID(_VZBridgedNetworkInterfaceClass.class), objc.Sel("networkInterfaces"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZBridgedNetworkInterface {
		return VZBridgedNetworkInterfaceFromID(id)
	})
}



















