// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NETunnelProviderProtocol] class.
var (
	_NETunnelProviderProtocolClass     NETunnelProviderProtocolClass
	_NETunnelProviderProtocolClassOnce sync.Once
)

func getNETunnelProviderProtocolClass() NETunnelProviderProtocolClass {
	_NETunnelProviderProtocolClassOnce.Do(func() {
		_NETunnelProviderProtocolClass = NETunnelProviderProtocolClass{class: objc.GetClass("NETunnelProviderProtocol")}
	})
	return _NETunnelProviderProtocolClass
}

// GetNETunnelProviderProtocolClass returns the class object for NETunnelProviderProtocol.
func GetNETunnelProviderProtocolClass() NETunnelProviderProtocolClass {
	return getNETunnelProviderProtocolClass()
}

type NETunnelProviderProtocolClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NETunnelProviderProtocolClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NETunnelProviderProtocolClass) Alloc() NETunnelProviderProtocol {
	rv := objc.Send[NETunnelProviderProtocol](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Configuration parameters for a VPN tunnel.
//
// # Overview
// 
// [NETunnelProviderProtocol] objects are used to specify configuration
// parameters for Tunnel Provider extensions.
//
// # Accessing the tunnel configuration
//
//   - [NETunnelProviderProtocol.ProviderConfiguration]: A dictionary containing keys and values defined by the Tunnel Provider developer.
//   - [NETunnelProviderProtocol.SetProviderConfiguration]
//   - [NETunnelProviderProtocol.ProviderBundleIdentifier]: A string identifying the specific Tunnel Provider extension that should be used with this configuration.
//   - [NETunnelProviderProtocol.SetProviderBundleIdentifier]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderProtocol
type NETunnelProviderProtocol struct {
	NEVPNProtocol
}

// NETunnelProviderProtocolFromID constructs a [NETunnelProviderProtocol] from an objc.ID.
//
// Configuration parameters for a VPN tunnel.
func NETunnelProviderProtocolFromID(id objc.ID) NETunnelProviderProtocol {
	return NETunnelProviderProtocol{NEVPNProtocol: NEVPNProtocolFromID(id)}
}
// NOTE: NETunnelProviderProtocol adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NETunnelProviderProtocol] class.
//
// # Accessing the tunnel configuration
//
//   - [INETunnelProviderProtocol.ProviderConfiguration]: A dictionary containing keys and values defined by the Tunnel Provider developer.
//   - [INETunnelProviderProtocol.SetProviderConfiguration]
//   - [INETunnelProviderProtocol.ProviderBundleIdentifier]: A string identifying the specific Tunnel Provider extension that should be used with this configuration.
//   - [INETunnelProviderProtocol.SetProviderBundleIdentifier]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderProtocol
type INETunnelProviderProtocol interface {
	INEVPNProtocol

	// Topic: Accessing the tunnel configuration

	// A dictionary containing keys and values defined by the Tunnel Provider developer.
	ProviderConfiguration() foundation.INSDictionary
	SetProviderConfiguration(value foundation.INSDictionary)
	// A string identifying the specific Tunnel Provider extension that should be used with this configuration.
	ProviderBundleIdentifier() string
	SetProviderBundleIdentifier(value string)
}

// Init initializes the instance.
func (t NETunnelProviderProtocol) Init() NETunnelProviderProtocol {
	rv := objc.Send[NETunnelProviderProtocol](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NETunnelProviderProtocol) Autorelease() NETunnelProviderProtocol {
	rv := objc.Send[NETunnelProviderProtocol](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETunnelProviderProtocol creates a new NETunnelProviderProtocol instance.
func NewNETunnelProviderProtocol() NETunnelProviderProtocol {
	class := getNETunnelProviderProtocolClass()
	rv := objc.Send[NETunnelProviderProtocol](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A dictionary containing keys and values defined by the Tunnel Provider
// developer.
//
// # Discussion
// 
// All of the keys and values in this dictionary must conform to the
// [NSSecureCoding] and [NSCopying] protocols.
//
// [NSCopying]: https://developer.apple.com/documentation/Foundation/NSCopying
// [NSSecureCoding]: https://developer.apple.com/documentation/Foundation/NSSecureCoding
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderProtocol/providerConfiguration
func (t NETunnelProviderProtocol) ProviderConfiguration() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("providerConfiguration"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t NETunnelProviderProtocol) SetProviderConfiguration(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setProviderConfiguration:"), value)
}
// A string identifying the specific Tunnel Provider extension that should be
// used with this configuration.
//
// # Discussion
// 
// A single app may contain multiple Tunnel Provider extensions. This property
// is used to specify which Tunnel Provider extension should be used with this
// configuration.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderProtocol/providerBundleIdentifier
func (t NETunnelProviderProtocol) ProviderBundleIdentifier() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("providerBundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (t NETunnelProviderProtocol) SetProviderBundleIdentifier(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setProviderBundleIdentifier:"), objc.String(value))
}

