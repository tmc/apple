// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NEEthernetTunnelNetworkSettings] class.
var (
	_NEEthernetTunnelNetworkSettingsClass     NEEthernetTunnelNetworkSettingsClass
	_NEEthernetTunnelNetworkSettingsClassOnce sync.Once
)

func getNEEthernetTunnelNetworkSettingsClass() NEEthernetTunnelNetworkSettingsClass {
	_NEEthernetTunnelNetworkSettingsClassOnce.Do(func() {
		_NEEthernetTunnelNetworkSettingsClass = NEEthernetTunnelNetworkSettingsClass{class: objc.GetClass("NEEthernetTunnelNetworkSettings")}
	})
	return _NEEthernetTunnelNetworkSettingsClass
}

// GetNEEthernetTunnelNetworkSettingsClass returns the class object for NEEthernetTunnelNetworkSettings.
func GetNEEthernetTunnelNetworkSettingsClass() NEEthernetTunnelNetworkSettingsClass {
	return getNEEthernetTunnelNetworkSettingsClass()
}

type NEEthernetTunnelNetworkSettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEEthernetTunnelNetworkSettingsClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEEthernetTunnelNetworkSettingsClass) Alloc() NEEthernetTunnelNetworkSettings {
	rv := objc.Send[NEEthernetTunnelNetworkSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The network settings for an ethernet-based VPN tunnel.
//
// # Overview
// 
// You use this type with [NEEthernetTunnelProvider] instances to communicate
// the desired network settings for the packet tunnel to the framework. The
// framework takes care of applying the contained settings to the system.
// 
// Instances of this class are thread-safe.
//
// # Creating a settings instance
//
//   - [NEEthernetTunnelNetworkSettings.InitWithTunnelRemoteAddressEthernetAddressMtu]: Creates a settings object with a given tunnel remote address and MAC address.
//
// # Inspecting settings properties
//
//   - [NEEthernetTunnelNetworkSettings.EthernetAddress]: The ethernet address of the tunnel interface, as a string.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEEthernetTunnelNetworkSettings
type NEEthernetTunnelNetworkSettings struct {
	NEPacketTunnelNetworkSettings
}

// NEEthernetTunnelNetworkSettingsFromID constructs a [NEEthernetTunnelNetworkSettings] from an objc.ID.
//
// The network settings for an ethernet-based VPN tunnel.
func NEEthernetTunnelNetworkSettingsFromID(id objc.ID) NEEthernetTunnelNetworkSettings {
	return NEEthernetTunnelNetworkSettings{NEPacketTunnelNetworkSettings: NEPacketTunnelNetworkSettingsFromID(id)}
}
// NOTE: NEEthernetTunnelNetworkSettings adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEEthernetTunnelNetworkSettings] class.
//
// # Creating a settings instance
//
//   - [INEEthernetTunnelNetworkSettings.InitWithTunnelRemoteAddressEthernetAddressMtu]: Creates a settings object with a given tunnel remote address and MAC address.
//
// # Inspecting settings properties
//
//   - [INEEthernetTunnelNetworkSettings.EthernetAddress]: The ethernet address of the tunnel interface, as a string.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEEthernetTunnelNetworkSettings
type INEEthernetTunnelNetworkSettings interface {
	INEPacketTunnelNetworkSettings

	// Topic: Creating a settings instance

	// Creates a settings object with a given tunnel remote address and MAC address.
	InitWithTunnelRemoteAddressEthernetAddressMtu(address string, ethernetAddress string, mtu int) NEEthernetTunnelNetworkSettings

	// Topic: Inspecting settings properties

	// The ethernet address of the tunnel interface, as a string.
	EthernetAddress() string
}

// Init initializes the instance.
func (e NEEthernetTunnelNetworkSettings) Init() NEEthernetTunnelNetworkSettings {
	rv := objc.Send[NEEthernetTunnelNetworkSettings](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e NEEthernetTunnelNetworkSettings) Autorelease() NEEthernetTunnelNetworkSettings {
	rv := objc.Send[NEEthernetTunnelNetworkSettings](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEEthernetTunnelNetworkSettings creates a new NEEthernetTunnelNetworkSettings instance.
func NewNEEthernetTunnelNetworkSettings() NEEthernetTunnelNetworkSettings {
	class := getNEEthernetTunnelNetworkSettingsClass()
	rv := objc.Send[NEEthernetTunnelNetworkSettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initialize a [NETunnelNetworkSettings] object.
//
// address: The IP address of the tunnel server.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings/init(tunnelRemoteAddress:)
func NewEthernetTunnelNetworkSettingsWithTunnelRemoteAddress(address string) NEEthernetTunnelNetworkSettings {
	instance := getNEEthernetTunnelNetworkSettingsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTunnelRemoteAddress:"), objc.String(address))
	return NEEthernetTunnelNetworkSettingsFromID(rv)
}

// Creates a settings object with a given tunnel remote address and MAC
// address.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEEthernetTunnelNetworkSettings/init(tunnelRemoteAddress:ethernetAddress:mtu:)
func NewEthernetTunnelNetworkSettingsWithTunnelRemoteAddressEthernetAddressMtu(address string, ethernetAddress string, mtu int) NEEthernetTunnelNetworkSettings {
	instance := getNEEthernetTunnelNetworkSettingsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTunnelRemoteAddress:ethernetAddress:mtu:"), objc.String(address), objc.String(ethernetAddress), mtu)
	return NEEthernetTunnelNetworkSettingsFromID(rv)
}

// Creates a settings object with a given tunnel remote address and MAC
// address.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEEthernetTunnelNetworkSettings/init(tunnelRemoteAddress:ethernetAddress:mtu:)
func (e NEEthernetTunnelNetworkSettings) InitWithTunnelRemoteAddressEthernetAddressMtu(address string, ethernetAddress string, mtu int) NEEthernetTunnelNetworkSettings {
	rv := objc.Send[NEEthernetTunnelNetworkSettings](e.ID, objc.Sel("initWithTunnelRemoteAddress:ethernetAddress:mtu:"), objc.String(address), objc.String(ethernetAddress), mtu)
	return rv
}

// The ethernet address of the tunnel interface, as a string.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEEthernetTunnelNetworkSettings/ethernetAddress
func (e NEEthernetTunnelNetworkSettings) EthernetAddress() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("ethernetAddress"))
	return foundation.NSStringFromID(rv).String()
}

