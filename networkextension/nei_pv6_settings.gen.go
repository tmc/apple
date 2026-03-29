// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEIPv6Settings] class.
var (
	_NEIPv6SettingsClass     NEIPv6SettingsClass
	_NEIPv6SettingsClassOnce sync.Once
)

func getNEIPv6SettingsClass() NEIPv6SettingsClass {
	_NEIPv6SettingsClassOnce.Do(func() {
		_NEIPv6SettingsClass = NEIPv6SettingsClass{class: objc.GetClass("NEIPv6Settings")}
	})
	return _NEIPv6SettingsClass
}

// GetNEIPv6SettingsClass returns the class object for NEIPv6Settings.
func GetNEIPv6SettingsClass() NEIPv6SettingsClass {
	return getNEIPv6SettingsClass()
}

type NEIPv6SettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEIPv6SettingsClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEIPv6SettingsClass) Alloc() NEIPv6Settings {
	rv := objc.Send[NEIPv6Settings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The IPv6 settings of an IP layer network tunnel.
//
// # Overview
// 
// To specify the IPv6 settings of a packet tunnel, set its
// [NEPacketTunnelNetworkSettings].[IPv6Settings] property to an instance of
// this class.
//
// # Initializing IPv6 settings
//
//   - [NEIPv6Settings.InitWithAddressesNetworkPrefixLengths]: Initializes the IPv6 settings object.
//
// # Accessing IPv6 properties
//
//   - [NEIPv6Settings.Addresses]: The IPv6 addresses to assign to the TUN interface.
//   - [NEIPv6Settings.NetworkPrefixLengths]: The IPv6 network prefix lengths to assign to the TUN interface.
//
// # Routing network traffic
//
//   - [NEIPv6Settings.IncludedRoutes]: The IPv6 network traffic that the system routes to the TUN interface.
//   - [NEIPv6Settings.SetIncludedRoutes]
//   - [NEIPv6Settings.ExcludedRoutes]: The IPv6 network traffic that the system routes to the primary physical interface, not the TUN interface.
//   - [NEIPv6Settings.SetExcludedRoutes]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings
type NEIPv6Settings struct {
	objectivec.Object
}

// NEIPv6SettingsFromID constructs a [NEIPv6Settings] from an objc.ID.
//
// The IPv6 settings of an IP layer network tunnel.
func NEIPv6SettingsFromID(id objc.ID) NEIPv6Settings {
	return NEIPv6Settings{objectivec.Object{ID: id}}
}
// NOTE: NEIPv6Settings adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEIPv6Settings] class.
//
// # Initializing IPv6 settings
//
//   - [INEIPv6Settings.InitWithAddressesNetworkPrefixLengths]: Initializes the IPv6 settings object.
//
// # Accessing IPv6 properties
//
//   - [INEIPv6Settings.Addresses]: The IPv6 addresses to assign to the TUN interface.
//   - [INEIPv6Settings.NetworkPrefixLengths]: The IPv6 network prefix lengths to assign to the TUN interface.
//
// # Routing network traffic
//
//   - [INEIPv6Settings.IncludedRoutes]: The IPv6 network traffic that the system routes to the TUN interface.
//   - [INEIPv6Settings.SetIncludedRoutes]
//   - [INEIPv6Settings.ExcludedRoutes]: The IPv6 network traffic that the system routes to the primary physical interface, not the TUN interface.
//   - [INEIPv6Settings.SetExcludedRoutes]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings
type INEIPv6Settings interface {
	objectivec.IObject

	// Topic: Initializing IPv6 settings

	// Initializes the IPv6 settings object.
	InitWithAddressesNetworkPrefixLengths(addresses []string, networkPrefixLengths []foundation.NSNumber) NEIPv6Settings

	// Topic: Accessing IPv6 properties

	// The IPv6 addresses to assign to the TUN interface.
	Addresses() []string
	// The IPv6 network prefix lengths to assign to the TUN interface.
	NetworkPrefixLengths() []foundation.NSNumber

	// Topic: Routing network traffic

	// The IPv6 network traffic that the system routes to the TUN interface.
	IncludedRoutes() []NEIPv6Route
	SetIncludedRoutes(value []NEIPv6Route)
	// The IPv6 network traffic that the system routes to the primary physical interface, not the TUN interface.
	ExcludedRoutes() []NEIPv6Route
	SetExcludedRoutes(value []NEIPv6Route)

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
func (i NEIPv6Settings) Init() NEIPv6Settings {
	rv := objc.Send[NEIPv6Settings](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i NEIPv6Settings) Autorelease() NEIPv6Settings {
	rv := objc.Send[NEIPv6Settings](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEIPv6Settings creates a new NEIPv6Settings instance.
func NewNEIPv6Settings() NEIPv6Settings {
	class := getNEIPv6SettingsClass()
	rv := objc.Send[NEIPv6Settings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes the IPv6 settings object.
//
// addresses: An array of IPv6 address strings. These IPv6 addresses will be assigned to
// the tunnel’s TUN interface.
//
// networkPrefixLengths: An array of IPv6 network prefix lengths. Each prefix length in this array
// is combined with the IP address in the corresponding index in `addresses`
// to specify an IPv6 network that the TUN interface is (virtually) connected
// to. Each prefix length must be set to an integer between 0 and 128.
//
// # Return Value
// 
// The initialized [NEIPv6Settings] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings/init(addresses:networkPrefixLengths:)
func NewIPv6SettingsWithAddressesNetworkPrefixLengths(addresses []string, networkPrefixLengths []foundation.NSNumber) NEIPv6Settings {
	instance := getNEIPv6SettingsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAddresses:networkPrefixLengths:"), objectivec.StringSliceToNSArray(addresses), objectivec.IObjectSliceToNSArray(networkPrefixLengths))
	return NEIPv6SettingsFromID(rv)
}

// Initializes the IPv6 settings object.
//
// addresses: An array of IPv6 address strings. These IPv6 addresses will be assigned to
// the tunnel’s TUN interface.
//
// networkPrefixLengths: An array of IPv6 network prefix lengths. Each prefix length in this array
// is combined with the IP address in the corresponding index in `addresses`
// to specify an IPv6 network that the TUN interface is (virtually) connected
// to. Each prefix length must be set to an integer between 0 and 128.
//
// # Return Value
// 
// The initialized [NEIPv6Settings] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings/init(addresses:networkPrefixLengths:)
func (i NEIPv6Settings) InitWithAddressesNetworkPrefixLengths(addresses []string, networkPrefixLengths []foundation.NSNumber) NEIPv6Settings {
	rv := objc.Send[NEIPv6Settings](i.ID, objc.Sel("initWithAddresses:networkPrefixLengths:"), objectivec.StringSliceToNSArray(addresses), objectivec.IObjectSliceToNSArray(networkPrefixLengths))
	return rv
}
func (i NEIPv6Settings) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The IPv6 addresses to assign to the TUN interface.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings/addresses
func (i NEIPv6Settings) Addresses() []string {
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("addresses"))
	return objc.ConvertSliceToStrings(rv)
}
// The IPv6 network prefix lengths to assign to the TUN interface.
//
// # Discussion
// 
// Each network prefix length in this array is combined with the IP address in
// the corresponding index in `addresses` to specify an IPv6 network that the
// TUN interface is (virtually) connected to.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings/networkPrefixLengths
func (i NEIPv6Settings) NetworkPrefixLengths() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("networkPrefixLengths"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
// The IPv6 network traffic that the system routes to the TUN interface.
//
// # Discussion
// 
// If you include the default route (`0.0.0.0/0` or `::/0`) in this property,
// the system routes traffic that doesn’t match a specific rule in the
// system routing table through the VPN.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings/includedRoutes
func (i NEIPv6Settings) IncludedRoutes() []NEIPv6Route {
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("includedRoutes"))
	return objc.ConvertSlice(rv, func(id objc.ID) NEIPv6Route {
		return NEIPv6RouteFromID(id)
	})
}
func (i NEIPv6Settings) SetIncludedRoutes(value []NEIPv6Route) {
	objc.Send[struct{}](i.ID, objc.Sel("setIncludedRoutes:"), objectivec.IObjectSliceToNSArray(value))
}
// The IPv6 network traffic that the system routes to the primary physical
// interface, not the TUN interface.
//
// # Discussion
// 
// This property excludes routes that the system might otherwise include from
// the [IncludedRoutes] property. The system automatically excludes the IP
// address of the tunnel server.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings/excludedRoutes
func (i NEIPv6Settings) ExcludedRoutes() []NEIPv6Route {
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("excludedRoutes"))
	return objc.ConvertSlice(rv, func(id objc.ID) NEIPv6Route {
		return NEIPv6RouteFromID(id)
	})
}
func (i NEIPv6Settings) SetExcludedRoutes(value []NEIPv6Route) {
	objc.Send[struct{}](i.ID, objc.Sel("setExcludedRoutes:"), objectivec.IObjectSliceToNSArray(value))
}
// The tunnel IP version 4 settings.
//
// See: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv4settings
func (i NEIPv6Settings) Ipv4Settings() INEIPv4Settings {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("IPv4Settings"))
	return NEIPv4SettingsFromID(objc.ID(rv))
}
func (i NEIPv6Settings) SetIpv4Settings(value INEIPv4Settings) {
	objc.Send[struct{}](i.ID, objc.Sel("setIPv4Settings:"), value)
}
// The tunnel IP version 6 settings.
//
// See: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv6settings
func (i NEIPv6Settings) Ipv6Settings() INEIPv6Settings {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("IPv6Settings"))
	return NEIPv6SettingsFromID(objc.ID(rv))
}
func (i NEIPv6Settings) SetIpv6Settings(value INEIPv6Settings) {
	objc.Send[struct{}](i.ID, objc.Sel("setIPv6Settings:"), value)
}
// The size of the maximum trasnmission unit, in bytes.
//
// See: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/mtu
func (i NEIPv6Settings) Mtu() foundation.NSNumber {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("MTU"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (i NEIPv6Settings) SetMtu(value foundation.NSNumber) {
	objc.Send[struct{}](i.ID, objc.Sel("setMTU:"), value)
}
// The number of bytes added to each tunneled packet for storing tunneling
// protocol headers.
//
// See: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/tunneloverheadbytes
func (i NEIPv6Settings) TunnelOverheadBytes() foundation.NSNumber {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("tunnelOverheadBytes"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (i NEIPv6Settings) SetTunnelOverheadBytes(value foundation.NSNumber) {
	objc.Send[struct{}](i.ID, objc.Sel("setTunnelOverheadBytes:"), value)
}

