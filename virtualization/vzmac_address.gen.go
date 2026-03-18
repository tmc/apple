// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMACAddress] class.
var (
	_VZMACAddressClass     VZMACAddressClass
	_VZMACAddressClassOnce sync.Once
)

func getVZMACAddressClass() VZMACAddressClass {
	_VZMACAddressClassOnce.Do(func() {
		_VZMACAddressClass = VZMACAddressClass{class: objc.GetClass("VZMACAddress")}
	})
	return _VZMACAddressClass
}

// GetVZMACAddressClass returns the class object for VZMACAddress.
func GetVZMACAddressClass() VZMACAddressClass {
	return getVZMACAddressClass()
}

type VZMACAddressClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMACAddressClass) Alloc() VZMACAddress {
	rv := objc.Send[VZMACAddress](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The media access control (MAC) address for a network interface in your
// virtual machine.
//
// # Overview
// 
// A [VZMACAddress] object contains the hardware address of your network
// interface. Every network device has a unique 48-bit MAC address that the
// system uses to route network packets to that device.
// 
// Call the [VZMACAddress.RandomLocallyAdministeredAddress] method to get a local MAC
// address suitable for use with your network interfaces. Alternatively, you
// can create a [VZMACAddress] object yourself from a string or `ether_addr_t`
// structure.
//
// # Creating a MAC address
//
//   - [VZMACAddress.InitWithString]: Creates a MAC address object from a specially formatted string.
//   - [VZMACAddress.InitWithEthernetAddress]: Creates a MAC address from the specified 48-bit Ethernet address.
//
// # Getting the address
//
//   - [VZMACAddress.String]: The MAC address as a formatted string.
//   - [VZMACAddress.EthernetAddress]: The MAC address as an Ethernet data structure.
//
// # Getting address attributes
//
//   - [VZMACAddress.IsBroadcastAddress]: A Boolean value that indicates whether the address is a broadcast address.
//   - [VZMACAddress.IsMulticastAddress]: A Boolean value that indicates whether the address is a multicast address.
//   - [VZMACAddress.IsUnicastAddress]: A Boolean value that indicates whether the address is a unicast address.
//   - [VZMACAddress.IsLocallyAdministeredAddress]: A Boolean value that indicates whether the address is a locally administered address (LAA).
//   - [VZMACAddress.IsUniversallyAdministeredAddress]: A Boolean value that indicates whether the address is a universally adminstered address (UAA).
//
// See: https://developer.apple.com/documentation/Virtualization/VZMACAddress
type VZMACAddress struct {
	objectivec.Object
}

// VZMACAddressFromID constructs a [VZMACAddress] from an objc.ID.
//
// The media access control (MAC) address for a network interface in your
// virtual machine.
func VZMACAddressFromID(id objc.ID) VZMACAddress {
	return VZMACAddress{objectivec.Object{ID: id}}
}
// NOTE: VZMACAddress adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZMACAddress] class.
//
// # Creating a MAC address
//
//   - [IVZMACAddress.InitWithString]: Creates a MAC address object from a specially formatted string.
//   - [IVZMACAddress.InitWithEthernetAddress]: Creates a MAC address from the specified 48-bit Ethernet address.
//
// # Getting the address
//
//   - [IVZMACAddress.String]: The MAC address as a formatted string.
//   - [IVZMACAddress.EthernetAddress]: The MAC address as an Ethernet data structure.
//
// # Getting address attributes
//
//   - [IVZMACAddress.IsBroadcastAddress]: A Boolean value that indicates whether the address is a broadcast address.
//   - [IVZMACAddress.IsMulticastAddress]: A Boolean value that indicates whether the address is a multicast address.
//   - [IVZMACAddress.IsUnicastAddress]: A Boolean value that indicates whether the address is a unicast address.
//   - [IVZMACAddress.IsLocallyAdministeredAddress]: A Boolean value that indicates whether the address is a locally administered address (LAA).
//   - [IVZMACAddress.IsUniversallyAdministeredAddress]: A Boolean value that indicates whether the address is a universally adminstered address (UAA).
//
// See: https://developer.apple.com/documentation/Virtualization/VZMACAddress
type IVZMACAddress interface {
	objectivec.IObject

	// Topic: Creating a MAC address

	// Creates a MAC address object from a specially formatted string.
	InitWithString(string_ string) VZMACAddress
	// Creates a MAC address from the specified 48-bit Ethernet address.
	InitWithEthernetAddress(ethernetAddress unsafe.Pointer) VZMACAddress

	// Topic: Getting the address

	// The MAC address as a formatted string.
	String() string
	// The MAC address as an Ethernet data structure.
	EthernetAddress() [6]byte

	// Topic: Getting address attributes

	// A Boolean value that indicates whether the address is a broadcast address.
	IsBroadcastAddress() bool
	// A Boolean value that indicates whether the address is a multicast address.
	IsMulticastAddress() bool
	// A Boolean value that indicates whether the address is a unicast address.
	IsUnicastAddress() bool
	// A Boolean value that indicates whether the address is a locally administered address (LAA).
	IsLocallyAdministeredAddress() bool
	// A Boolean value that indicates whether the address is a universally adminstered address (UAA).
	IsUniversallyAdministeredAddress() bool
}

// Init initializes the instance.
func (m VZMACAddress) Init() VZMACAddress {
	rv := objc.Send[VZMACAddress](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMACAddress) Autorelease() VZMACAddress {
	rv := objc.Send[VZMACAddress](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMACAddress creates a new VZMACAddress instance.
func NewVZMACAddress() VZMACAddress {
	class := getVZMACAddressClass()
	rv := objc.Send[VZMACAddress](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a MAC address from the specified 48-bit Ethernet address.
//
// ethernetAddress: A 48-bit Ethernet address.
//
// # Return Value
// 
// A MAC address object with the specified Ethernet address.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMACAddress/init(ethernetAddress:)
func NewMACAddressWithEthernetAddress(ethernetAddress unsafe.Pointer) VZMACAddress {
	instance := getVZMACAddressClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEthernetAddress:"), ethernetAddress)
	return VZMACAddressFromID(rv)
}

// Creates a MAC address object from a specially formatted string.
//
// string: A string that contains the 6 hexadecimal bytes of the MAC address separated
// by colon characters. An example string is `01:23:45:ef`. The string is
// case-insensitive, so you may use uppercase or lowercase for alphabetical
// characters.
//
// # Return Value
// 
// A MAC address object with the specified value, or `nil` if the string is
// formatted incorrectly.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMACAddress/init(string:)
func NewMACAddressWithString(string_ string) VZMACAddress {
	instance := getVZMACAddressClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:"), objc.String(string_))
	return VZMACAddressFromID(rv)
}

// Creates a MAC address object from a specially formatted string.
//
// string: A string that contains the 6 hexadecimal bytes of the MAC address separated
// by colon characters. An example string is `01:23:45:ef`. The string is
// case-insensitive, so you may use uppercase or lowercase for alphabetical
// characters.
//
// # Return Value
// 
// A MAC address object with the specified value, or `nil` if the string is
// formatted incorrectly.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMACAddress/init(string:)
func (m VZMACAddress) InitWithString(string_ string) VZMACAddress {
	rv := objc.Send[VZMACAddress](m.ID, objc.Sel("initWithString:"), objc.String(string_))
	return rv
}

// Creates a MAC address from the specified 48-bit Ethernet address.
//
// ethernetAddress: A 48-bit Ethernet address.
//
// # Return Value
// 
// A MAC address object with the specified Ethernet address.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMACAddress/init(ethernetAddress:)
func (m VZMACAddress) InitWithEthernetAddress(ethernetAddress unsafe.Pointer) VZMACAddress {
	rv := objc.Send[VZMACAddress](m.ID, objc.Sel("initWithEthernetAddress:"), ethernetAddress)
	return rv
}

// Returns a valid, random, locally administered, unicast MAC address.
//
// # Return Value
// 
// A MAC address suitable for use in your network devices.
//
// # Discussion
// 
// This method doesn’t guarantee the generation of a unique MAC address.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMACAddress/randomLocallyAdministered()
func (_VZMACAddressClass VZMACAddressClass) RandomLocallyAdministeredAddress() VZMACAddress {
	rv := objc.Send[objc.ID](objc.ID(_VZMACAddressClass.class), objc.Sel("randomLocallyAdministeredAddress"))
	return VZMACAddressFromID(rv)
}

// The MAC address as a formatted string.
//
// # Discussion
// 
// The string contains the 6 hexadecimal bytes of the MAC address, separated
// by colon characters. Alphabetical characters are lowercase in the string.
// An example string is `01:23:45:ef`.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMACAddress/string
func (m VZMACAddress) String() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("string"))
	return foundation.NSStringFromID(rv).String()
}

// The MAC address as an Ethernet data structure.
//
// # Discussion
// 
// Use this property to obtain the individual octets of the Ethernet address.
// For more information, see `ether_addr_t`.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMACAddress/ethernetAddress
func (m VZMACAddress) EthernetAddress() [6]byte {
	rv := objc.Send[[6]byte](m.ID, objc.Sel("ethernetAddress"))
	return [6]byte(rv)
}

// A Boolean value that indicates whether the address is a broadcast address.
//
// # Discussion
// 
// The value of this property is [true] if the address is a broadcast address,
// or [false] if it isn’t.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Virtualization/VZMACAddress/isBroadcastAddress
func (m VZMACAddress) IsBroadcastAddress() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isBroadcastAddress"))
	return rv
}

// A Boolean value that indicates whether the address is a multicast address.
//
// # Discussion
// 
// The value of this property is [true] if the address is a multicast address,
// or [false] if it isn’t.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Virtualization/VZMACAddress/isMulticastAddress
func (m VZMACAddress) IsMulticastAddress() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isMulticastAddress"))
	return rv
}

// A Boolean value that indicates whether the address is a unicast address.
//
// # Discussion
// 
// The value of this property is [true] if the address is a unicast address,
// or [false] if it isn’t.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Virtualization/VZMACAddress/isUnicastAddress
func (m VZMACAddress) IsUnicastAddress() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isUnicastAddress"))
	return rv
}

// A Boolean value that indicates whether the address is a locally
// administered address (LAA).
//
// # Discussion
// 
// The value of this property is [true] if the address is locally
// administered, or [false] if it isn’t. A locally administered address
// different than the address burned in to the physical network interface.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Virtualization/VZMACAddress/isLocallyAdministeredAddress
func (m VZMACAddress) IsLocallyAdministeredAddress() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isLocallyAdministeredAddress"))
	return rv
}

// A Boolean value that indicates whether the address is a universally
// adminstered address (UAA).
//
// # Discussion
// 
// The value of this property is [true] if the address is universally
// administered, or [false] if it isn’t. The manufacturer of a device
// assigns an address of this type, and the address includes the
// organization’s unique identification code.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Virtualization/VZMACAddress/isUniversallyAdministeredAddress
func (m VZMACAddress) IsUniversallyAdministeredAddress() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isUniversallyAdministeredAddress"))
	return rv
}

