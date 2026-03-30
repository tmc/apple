// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEIPv4Settings] class.
var (
	_NEIPv4SettingsClass     NEIPv4SettingsClass
	_NEIPv4SettingsClassOnce sync.Once
)

func getNEIPv4SettingsClass() NEIPv4SettingsClass {
	_NEIPv4SettingsClassOnce.Do(func() {
		_NEIPv4SettingsClass = NEIPv4SettingsClass{class: objc.GetClass("NEIPv4Settings")}
	})
	return _NEIPv4SettingsClass
}

// GetNEIPv4SettingsClass returns the class object for NEIPv4Settings.
func GetNEIPv4SettingsClass() NEIPv4SettingsClass {
	return getNEIPv4SettingsClass()
}

type NEIPv4SettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEIPv4SettingsClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEIPv4SettingsClass) Alloc() NEIPv4Settings {
	rv := objc.Send[NEIPv4Settings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The IPv4 settings of an IP layer network tunnel.
//
// # Overview
//
// To specify the IPv4 settings of a packet tunnel, set its
// [NEPacketTunnelNetworkSettings].[IPv4Settings] property to an instance of
// this class.
//
// # Initializing IPv4 settings
//
//   - [NEIPv4Settings.InitWithAddressesSubnetMasks]: Initializes an IPv4 settings object.
//
// # Accessing IPv4 properties
//
//   - [NEIPv4Settings.Addresses]: The IPv4 addresses to assign to the TUN interface.
//   - [NEIPv4Settings.SubnetMasks]: The IPv4 network masks to assign to the TUN interface.
//   - [NEIPv4Settings.Router]: The address of the next-hop gateway router represented as a dotted decimal string.
//   - [NEIPv4Settings.SetRouter]
//
// # Routing network traffic
//
//   - [NEIPv4Settings.IncludedRoutes]: The IPv4 network traffic that the system routes to the TUN interface.
//   - [NEIPv4Settings.SetIncludedRoutes]
//   - [NEIPv4Settings.ExcludedRoutes]: The IPv4 network traffic that the system routes to the primary physical interface, not the TUN interface.
//   - [NEIPv4Settings.SetExcludedRoutes]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings
type NEIPv4Settings struct {
	objectivec.Object
}

// NEIPv4SettingsFromID constructs a [NEIPv4Settings] from an objc.ID.
//
// The IPv4 settings of an IP layer network tunnel.
func NEIPv4SettingsFromID(id objc.ID) NEIPv4Settings {
	return NEIPv4Settings{objectivec.Object{ID: id}}
}

// NOTE: NEIPv4Settings adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEIPv4Settings] class.
//
// # Initializing IPv4 settings
//
//   - [INEIPv4Settings.InitWithAddressesSubnetMasks]: Initializes an IPv4 settings object.
//
// # Accessing IPv4 properties
//
//   - [INEIPv4Settings.Addresses]: The IPv4 addresses to assign to the TUN interface.
//   - [INEIPv4Settings.SubnetMasks]: The IPv4 network masks to assign to the TUN interface.
//   - [INEIPv4Settings.Router]: The address of the next-hop gateway router represented as a dotted decimal string.
//   - [INEIPv4Settings.SetRouter]
//
// # Routing network traffic
//
//   - [INEIPv4Settings.IncludedRoutes]: The IPv4 network traffic that the system routes to the TUN interface.
//   - [INEIPv4Settings.SetIncludedRoutes]
//   - [INEIPv4Settings.ExcludedRoutes]: The IPv4 network traffic that the system routes to the primary physical interface, not the TUN interface.
//   - [INEIPv4Settings.SetExcludedRoutes]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings
type INEIPv4Settings interface {
	objectivec.IObject

	// Topic: Initializing IPv4 settings

	// Initializes an IPv4 settings object.
	InitWithAddressesSubnetMasks(addresses []string, subnetMasks []string) NEIPv4Settings

	// Topic: Accessing IPv4 properties

	// The IPv4 addresses to assign to the TUN interface.
	Addresses() []string
	// The IPv4 network masks to assign to the TUN interface.
	SubnetMasks() []string
	// The address of the next-hop gateway router represented as a dotted decimal string.
	Router() string
	SetRouter(value string)

	// Topic: Routing network traffic

	// The IPv4 network traffic that the system routes to the TUN interface.
	IncludedRoutes() []NEIPv4Route
	SetIncludedRoutes(value []NEIPv4Route)
	// The IPv4 network traffic that the system routes to the primary physical interface, not the TUN interface.
	ExcludedRoutes() []NEIPv4Route
	SetExcludedRoutes(value []NEIPv4Route)

	// The tunnel IP version 4 settings.
	Ipv4Settings() INEIPv4Settings
	SetIpv4Settings(value INEIPv4Settings)
	// The tunnel IP version 6 settings.
	Ipv6Settings() INEIPv6Settings
	SetIpv6Settings(value INEIPv6Settings)
	// The size of the maximum trasnmission unit, in bytes.
	Mtu() foundation.NSNumber
	SetMtu(value foundation.NSNumber)
	// The number of bytes added to each tunneled packet for storing tunneling protocol headers.
	TunnelOverheadBytes() foundation.NSNumber
	SetTunnelOverheadBytes(value foundation.NSNumber)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (i NEIPv4Settings) Init() NEIPv4Settings {
	rv := objc.Send[NEIPv4Settings](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i NEIPv4Settings) Autorelease() NEIPv4Settings {
	rv := objc.Send[NEIPv4Settings](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEIPv4Settings creates a new NEIPv4Settings instance.
func NewNEIPv4Settings() NEIPv4Settings {
	class := getNEIPv4SettingsClass()
	rv := objc.Send[NEIPv4Settings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes an IPv4 settings object.
//
// addresses: An array of IPv4 address strings. These IPv4 addresses will be assigned to
// the tunnel’s TUN interface.
//
// subnetMasks: An array of IPv4 network mask strings. Each mask in this array is combined
// with the IP address in the corresponding index in `addresses` to specify an
// IPv4 network that the TUN interface is (virtually) connected to.
//
// # Return Value
//
// The initialized NEIPv4Settings object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/init(addresses:subnetMasks:)
func NewIPv4SettingsWithAddressesSubnetMasks(addresses []string, subnetMasks []string) NEIPv4Settings {
	instance := getNEIPv4SettingsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAddresses:subnetMasks:"), objectivec.StringSliceToNSArray(addresses), objectivec.StringSliceToNSArray(subnetMasks))
	return NEIPv4SettingsFromID(rv)
}

// Initializes an IPv4 settings object.
//
// addresses: An array of IPv4 address strings. These IPv4 addresses will be assigned to
// the tunnel’s TUN interface.
//
// subnetMasks: An array of IPv4 network mask strings. Each mask in this array is combined
// with the IP address in the corresponding index in `addresses` to specify an
// IPv4 network that the TUN interface is (virtually) connected to.
//
// # Return Value
//
// The initialized NEIPv4Settings object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/init(addresses:subnetMasks:)
func (i NEIPv4Settings) InitWithAddressesSubnetMasks(addresses []string, subnetMasks []string) NEIPv4Settings {
	rv := objc.Send[NEIPv4Settings](i.ID, objc.Sel("initWithAddresses:subnetMasks:"), objectivec.StringSliceToNSArray(addresses), objectivec.StringSliceToNSArray(subnetMasks))
	return rv
}
func (i NEIPv4Settings) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeWithCoder:"), coder)
}

// # Return Value
//
// The initialized object.
//
// # Discussion
//
// Create a NEIPv4Settings object that will obtain IP addresses and netmasks
// using DHCP.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/settingsWithAutomaticAddressing
func (_NEIPv4SettingsClass NEIPv4SettingsClass) SettingsWithAutomaticAddressing() NEIPv4Settings {
	rv := objc.Send[objc.ID](objc.ID(_NEIPv4SettingsClass.class), objc.Sel("settingsWithAutomaticAddressing"))
	return NEIPv4SettingsFromID(rv)
}

// The IPv4 addresses to assign to the TUN interface.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/addresses
func (i NEIPv4Settings) Addresses() []string {
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("addresses"))
	return objc.ConvertSliceToStrings(rv)
}

// The IPv4 network masks to assign to the TUN interface.
//
// # Discussion
//
// Each mask in this array is combined with the IP address in the
// corresponding index in `addresses` to specify an IPv4 network that the TUN
// interface is (virtually) connected to.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/subnetMasks
func (i NEIPv4Settings) SubnetMasks() []string {
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("subnetMasks"))
	return objc.ConvertSliceToStrings(rv)
}

// The address of the next-hop gateway router represented as a dotted decimal
// string.
//
// # Discussion
//
// The system ignores this property for TUN interfaces.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/router
func (i NEIPv4Settings) Router() string {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("router"))
	return foundation.NSStringFromID(rv).String()
}
func (i NEIPv4Settings) SetRouter(value string) {
	objc.Send[struct{}](i.ID, objc.Sel("setRouter:"), objc.String(value))
}

// The IPv4 network traffic that the system routes to the TUN interface.
//
// # Discussion
//
// If you include the default route (`0.0.0.0/0` or `::/0`) in this property,
// the system routes traffic that doesn’t match a specific rule in the
// system routing table through the VPN.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/includedRoutes
func (i NEIPv4Settings) IncludedRoutes() []NEIPv4Route {
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("includedRoutes"))
	return objc.ConvertSlice(rv, func(id objc.ID) NEIPv4Route {
		return NEIPv4RouteFromID(id)
	})
}
func (i NEIPv4Settings) SetIncludedRoutes(value []NEIPv4Route) {
	objc.Send[struct{}](i.ID, objc.Sel("setIncludedRoutes:"), objectivec.IObjectSliceToNSArray(value))
}

// The IPv4 network traffic that the system routes to the primary physical
// interface, not the TUN interface.
//
// # Discussion
//
// This property excludes routes that the system might otherwise include from
// the [IncludedRoutes] property. The system automatically excludes the IP
// address of the tunnel server.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/excludedRoutes
func (i NEIPv4Settings) ExcludedRoutes() []NEIPv4Route {
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("excludedRoutes"))
	return objc.ConvertSlice(rv, func(id objc.ID) NEIPv4Route {
		return NEIPv4RouteFromID(id)
	})
}
func (i NEIPv4Settings) SetExcludedRoutes(value []NEIPv4Route) {
	objc.Send[struct{}](i.ID, objc.Sel("setExcludedRoutes:"), objectivec.IObjectSliceToNSArray(value))
}

// The tunnel IP version 4 settings.
//
// See: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv4settings
func (i NEIPv4Settings) Ipv4Settings() INEIPv4Settings {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("IPv4Settings"))
	return NEIPv4SettingsFromID(objc.ID(rv))
}
func (i NEIPv4Settings) SetIpv4Settings(value INEIPv4Settings) {
	objc.Send[struct{}](i.ID, objc.Sel("setIPv4Settings:"), value)
}

// The tunnel IP version 6 settings.
//
// See: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv6settings
func (i NEIPv4Settings) Ipv6Settings() INEIPv6Settings {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("IPv6Settings"))
	return NEIPv6SettingsFromID(objc.ID(rv))
}
func (i NEIPv4Settings) SetIpv6Settings(value INEIPv6Settings) {
	objc.Send[struct{}](i.ID, objc.Sel("setIPv6Settings:"), value)
}

// The size of the maximum trasnmission unit, in bytes.
//
// See: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/mtu
func (i NEIPv4Settings) Mtu() foundation.NSNumber {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("MTU"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (i NEIPv4Settings) SetMtu(value foundation.NSNumber) {
	objc.Send[struct{}](i.ID, objc.Sel("setMTU:"), value)
}

// The number of bytes added to each tunneled packet for storing tunneling
// protocol headers.
//
// See: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/tunneloverheadbytes
func (i NEIPv4Settings) TunnelOverheadBytes() foundation.NSNumber {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("tunnelOverheadBytes"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (i NEIPv4Settings) SetTunnelOverheadBytes(value foundation.NSNumber) {
	objc.Send[struct{}](i.ID, objc.Sel("setTunnelOverheadBytes:"), value)
}
