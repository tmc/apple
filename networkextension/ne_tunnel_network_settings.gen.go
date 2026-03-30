// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NETunnelNetworkSettings] class.
var (
	_NETunnelNetworkSettingsClass     NETunnelNetworkSettingsClass
	_NETunnelNetworkSettingsClassOnce sync.Once
)

func getNETunnelNetworkSettingsClass() NETunnelNetworkSettingsClass {
	_NETunnelNetworkSettingsClassOnce.Do(func() {
		_NETunnelNetworkSettingsClass = NETunnelNetworkSettingsClass{class: objc.GetClass("NETunnelNetworkSettings")}
	})
	return _NETunnelNetworkSettingsClass
}

// GetNETunnelNetworkSettingsClass returns the class object for NETunnelNetworkSettings.
func GetNETunnelNetworkSettingsClass() NETunnelNetworkSettingsClass {
	return getNETunnelNetworkSettingsClass()
}

type NETunnelNetworkSettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NETunnelNetworkSettingsClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NETunnelNetworkSettingsClass) Alloc() NETunnelNetworkSettings {
	rv := objc.Send[NETunnelNetworkSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The configuration for a tunnel provider’s virtual interface.
//
// # Initializing tunnel network settings
//
//   - [NETunnelNetworkSettings.InitWithTunnelRemoteAddress]: Initialize a [NETunnelNetworkSettings] object.
//
// # Accessing tunnel network settings
//
//   - [NETunnelNetworkSettings.TunnelRemoteAddress]: The IP address of the tunnel server.
//   - [NETunnelNetworkSettings.DNSSettings]: The tunnel DNS settings.
//   - [NETunnelNetworkSettings.SetDNSSettings]
//   - [NETunnelNetworkSettings.ProxySettings]: The tunnel HTTP proxy settings.
//   - [NETunnelNetworkSettings.SetProxySettings]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings
type NETunnelNetworkSettings struct {
	objectivec.Object
}

// NETunnelNetworkSettingsFromID constructs a [NETunnelNetworkSettings] from an objc.ID.
//
// The configuration for a tunnel provider’s virtual interface.
func NETunnelNetworkSettingsFromID(id objc.ID) NETunnelNetworkSettings {
	return NETunnelNetworkSettings{objectivec.Object{ID: id}}
}

// NOTE: NETunnelNetworkSettings adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NETunnelNetworkSettings] class.
//
// # Initializing tunnel network settings
//
//   - [INETunnelNetworkSettings.InitWithTunnelRemoteAddress]: Initialize a [NETunnelNetworkSettings] object.
//
// # Accessing tunnel network settings
//
//   - [INETunnelNetworkSettings.TunnelRemoteAddress]: The IP address of the tunnel server.
//   - [INETunnelNetworkSettings.DNSSettings]: The tunnel DNS settings.
//   - [INETunnelNetworkSettings.SetDNSSettings]
//   - [INETunnelNetworkSettings.ProxySettings]: The tunnel HTTP proxy settings.
//   - [INETunnelNetworkSettings.SetProxySettings]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings
type INETunnelNetworkSettings interface {
	objectivec.IObject

	// Topic: Initializing tunnel network settings

	// Initialize a [NETunnelNetworkSettings] object.
	InitWithTunnelRemoteAddress(address string) NETunnelNetworkSettings

	// Topic: Accessing tunnel network settings

	// The IP address of the tunnel server.
	TunnelRemoteAddress() string
	// The tunnel DNS settings.
	DNSSettings() INEDNSSettings
	SetDNSSettings(value INEDNSSettings)
	// The tunnel HTTP proxy settings.
	ProxySettings() INEProxySettings
	SetProxySettings(value INEProxySettings)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t NETunnelNetworkSettings) Init() NETunnelNetworkSettings {
	rv := objc.Send[NETunnelNetworkSettings](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NETunnelNetworkSettings) Autorelease() NETunnelNetworkSettings {
	rv := objc.Send[NETunnelNetworkSettings](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETunnelNetworkSettings creates a new NETunnelNetworkSettings instance.
func NewNETunnelNetworkSettings() NETunnelNetworkSettings {
	class := getNETunnelNetworkSettingsClass()
	rv := objc.Send[NETunnelNetworkSettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initialize a [NETunnelNetworkSettings] object.
//
// address: The IP address of the tunnel server.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings/init(tunnelRemoteAddress:)
func NewTunnelNetworkSettingsWithTunnelRemoteAddress(address string) NETunnelNetworkSettings {
	instance := getNETunnelNetworkSettingsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTunnelRemoteAddress:"), objc.String(address))
	return NETunnelNetworkSettingsFromID(rv)
}

// Initialize a [NETunnelNetworkSettings] object.
//
// address: The IP address of the tunnel server.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings/init(tunnelRemoteAddress:)
func (t NETunnelNetworkSettings) InitWithTunnelRemoteAddress(address string) NETunnelNetworkSettings {
	rv := objc.Send[NETunnelNetworkSettings](t.ID, objc.Sel("initWithTunnelRemoteAddress:"), objc.String(address))
	return rv
}
func (t NETunnelNetworkSettings) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The IP address of the tunnel server.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings/tunnelRemoteAddress
func (t NETunnelNetworkSettings) TunnelRemoteAddress() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tunnelRemoteAddress"))
	return foundation.NSStringFromID(rv).String()
}

// The tunnel DNS settings.
//
// # Discussion
//
// Network connections to hosts in the tunnel’s internal network will use
// these DNS settings when resolving host names.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings/dnsSettings
func (t NETunnelNetworkSettings) DNSSettings() INEDNSSettings {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("DNSSettings"))
	return NEDNSSettingsFromID(objc.ID(rv))
}
func (t NETunnelNetworkSettings) SetDNSSettings(value INEDNSSettings) {
	objc.Send[struct{}](t.ID, objc.Sel("setDNSSettings:"), value)
}

// The tunnel HTTP proxy settings.
//
// # Discussion
//
// HTTP connections to hosts in the tunnel’s internal network will use these
// proxy settings.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings/proxySettings
func (t NETunnelNetworkSettings) ProxySettings() INEProxySettings {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("proxySettings"))
	return NEProxySettingsFromID(objc.ID(rv))
}
func (t NETunnelNetworkSettings) SetProxySettings(value INEProxySettings) {
	objc.Send[struct{}](t.ID, objc.Sel("setProxySettings:"), value)
}
