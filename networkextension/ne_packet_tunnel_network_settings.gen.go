// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NEPacketTunnelNetworkSettings] class.
var (
	_NEPacketTunnelNetworkSettingsClass     NEPacketTunnelNetworkSettingsClass
	_NEPacketTunnelNetworkSettingsClassOnce sync.Once
)

func getNEPacketTunnelNetworkSettingsClass() NEPacketTunnelNetworkSettingsClass {
	_NEPacketTunnelNetworkSettingsClassOnce.Do(func() {
		_NEPacketTunnelNetworkSettingsClass = NEPacketTunnelNetworkSettingsClass{class: objc.GetClass("NEPacketTunnelNetworkSettings")}
	})
	return _NEPacketTunnelNetworkSettingsClass
}

// GetNEPacketTunnelNetworkSettingsClass returns the class object for NEPacketTunnelNetworkSettings.
func GetNEPacketTunnelNetworkSettingsClass() NEPacketTunnelNetworkSettingsClass {
	return getNEPacketTunnelNetworkSettingsClass()
}

type NEPacketTunnelNetworkSettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEPacketTunnelNetworkSettingsClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEPacketTunnelNetworkSettingsClass) Alloc() NEPacketTunnelNetworkSettings {
	rv := objc.Send[NEPacketTunnelNetworkSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The configuration for a packet tunnel provider’s virtual interface.
//
// # Accessing network properties
//
//   - [NEPacketTunnelNetworkSettings.IPv4Settings]: The tunnel IP version 4 settings.
//   - [NEPacketTunnelNetworkSettings.SetIPv4Settings]
//   - [NEPacketTunnelNetworkSettings.IPv6Settings]: The tunnel IP version 6 settings.
//   - [NEPacketTunnelNetworkSettings.SetIPv6Settings]
//   - [NEPacketTunnelNetworkSettings.TunnelOverheadBytes]: The number of bytes added to each tunneled packet for storing tunneling protocol headers.
//   - [NEPacketTunnelNetworkSettings.SetTunnelOverheadBytes]
//   - [NEPacketTunnelNetworkSettings.MTU]: The size of the maximum trasnmission unit, in bytes.
//   - [NEPacketTunnelNetworkSettings.SetMTU]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings
type NEPacketTunnelNetworkSettings struct {
	NETunnelNetworkSettings
}

// NEPacketTunnelNetworkSettingsFromID constructs a [NEPacketTunnelNetworkSettings] from an objc.ID.
//
// The configuration for a packet tunnel provider’s virtual interface.
func NEPacketTunnelNetworkSettingsFromID(id objc.ID) NEPacketTunnelNetworkSettings {
	return NEPacketTunnelNetworkSettings{NETunnelNetworkSettings: NETunnelNetworkSettingsFromID(id)}
}
// NOTE: NEPacketTunnelNetworkSettings adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEPacketTunnelNetworkSettings] class.
//
// # Accessing network properties
//
//   - [INEPacketTunnelNetworkSettings.IPv4Settings]: The tunnel IP version 4 settings.
//   - [INEPacketTunnelNetworkSettings.SetIPv4Settings]
//   - [INEPacketTunnelNetworkSettings.IPv6Settings]: The tunnel IP version 6 settings.
//   - [INEPacketTunnelNetworkSettings.SetIPv6Settings]
//   - [INEPacketTunnelNetworkSettings.TunnelOverheadBytes]: The number of bytes added to each tunneled packet for storing tunneling protocol headers.
//   - [INEPacketTunnelNetworkSettings.SetTunnelOverheadBytes]
//   - [INEPacketTunnelNetworkSettings.MTU]: The size of the maximum trasnmission unit, in bytes.
//   - [INEPacketTunnelNetworkSettings.SetMTU]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings
type INEPacketTunnelNetworkSettings interface {
	INETunnelNetworkSettings

	// Topic: Accessing network properties

	// The tunnel IP version 4 settings.
	IPv4Settings() INEIPv4Settings
	SetIPv4Settings(value INEIPv4Settings)
	// The tunnel IP version 6 settings.
	IPv6Settings() INEIPv6Settings
	SetIPv6Settings(value INEIPv6Settings)
	// The number of bytes added to each tunneled packet for storing tunneling protocol headers.
	TunnelOverheadBytes() foundation.NSNumber
	SetTunnelOverheadBytes(value foundation.NSNumber)
	// The size of the maximum trasnmission unit, in bytes.
	MTU() foundation.NSNumber
	SetMTU(value foundation.NSNumber)
}

// Init initializes the instance.
func (p NEPacketTunnelNetworkSettings) Init() NEPacketTunnelNetworkSettings {
	rv := objc.Send[NEPacketTunnelNetworkSettings](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NEPacketTunnelNetworkSettings) Autorelease() NEPacketTunnelNetworkSettings {
	rv := objc.Send[NEPacketTunnelNetworkSettings](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEPacketTunnelNetworkSettings creates a new NEPacketTunnelNetworkSettings instance.
func NewNEPacketTunnelNetworkSettings() NEPacketTunnelNetworkSettings {
	class := getNEPacketTunnelNetworkSettingsClass()
	rv := objc.Send[NEPacketTunnelNetworkSettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initialize a [NETunnelNetworkSettings] object.
//
// address: The IP address of the tunnel server.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings/init(tunnelRemoteAddress:)
func NewPacketTunnelNetworkSettingsWithTunnelRemoteAddress(address string) NEPacketTunnelNetworkSettings {
	instance := getNEPacketTunnelNetworkSettingsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTunnelRemoteAddress:"), objc.String(address))
	return NEPacketTunnelNetworkSettingsFromID(rv)
}

// The tunnel IP version 4 settings.
//
// # Discussion
// 
// This property contains the IPv4 routes specifying what IPv4 traffic to
// route to the tunnel, as well as the IPv4 address and netmask to assign to
// the TUN interface.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/ipv4Settings
func (p NEPacketTunnelNetworkSettings) IPv4Settings() INEIPv4Settings {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("IPv4Settings"))
	return NEIPv4SettingsFromID(objc.ID(rv))
}
func (p NEPacketTunnelNetworkSettings) SetIPv4Settings(value INEIPv4Settings) {
	objc.Send[struct{}](p.ID, objc.Sel("setIPv4Settings:"), value)
}
// The tunnel IP version 6 settings.
//
// # Discussion
// 
// This property contains the IPv6 routes specifying what IPv6 traffic to
// route to the tunnel, as well as the IPv6 address and network prefix to
// assign to the TUN interface.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/ipv6Settings
func (p NEPacketTunnelNetworkSettings) IPv6Settings() INEIPv6Settings {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("IPv6Settings"))
	return NEIPv6SettingsFromID(objc.ID(rv))
}
func (p NEPacketTunnelNetworkSettings) SetIPv6Settings(value INEIPv6Settings) {
	objc.Send[struct{}](p.ID, objc.Sel("setIPv6Settings:"), value)
}
// The number of bytes added to each tunneled packet for storing tunneling
// protocol headers.
//
// # Discussion
// 
// The value of this property is subtracted from the Maximum Transmission Unit
// (MTU) of the tunnel’s underlying physical network interface to compute
// the MTU of the TUN interface.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/tunnelOverheadBytes
func (p NEPacketTunnelNetworkSettings) TunnelOverheadBytes() foundation.NSNumber {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("tunnelOverheadBytes"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (p NEPacketTunnelNetworkSettings) SetTunnelOverheadBytes(value foundation.NSNumber) {
	objc.Send[struct{}](p.ID, objc.Sel("setTunnelOverheadBytes:"), value)
}
// The size of the maximum trasnmission unit, in bytes.
//
// # Discussion
// 
// The maximum transmission unit (MTU) size represents the largest number of
// bytes that anything can assign to the TUN interface.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/mtu
func (p NEPacketTunnelNetworkSettings) MTU() foundation.NSNumber {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("MTU"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (p NEPacketTunnelNetworkSettings) SetMTU(value foundation.NSNumber) {
	objc.Send[struct{}](p.ID, objc.Sel("setMTU:"), value)
}

