// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NEDNSProxyProviderProtocol] class.
var (
	_NEDNSProxyProviderProtocolClass     NEDNSProxyProviderProtocolClass
	_NEDNSProxyProviderProtocolClassOnce sync.Once
)

func getNEDNSProxyProviderProtocolClass() NEDNSProxyProviderProtocolClass {
	_NEDNSProxyProviderProtocolClassOnce.Do(func() {
		_NEDNSProxyProviderProtocolClass = NEDNSProxyProviderProtocolClass{class: objc.GetClass("NEDNSProxyProviderProtocol")}
	})
	return _NEDNSProxyProviderProtocolClass
}

// GetNEDNSProxyProviderProtocolClass returns the class object for NEDNSProxyProviderProtocol.
func GetNEDNSProxyProviderProtocolClass() NEDNSProxyProviderProtocolClass {
	return getNEDNSProxyProviderProtocolClass()
}

type NEDNSProxyProviderProtocolClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEDNSProxyProviderProtocolClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEDNSProxyProviderProtocolClass) Alloc() NEDNSProxyProviderProtocol {
	rv := objc.Send[NEDNSProxyProviderProtocol](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Configuration parameters for a DNS proxy.
//
// # Accessing the DNS proxy configuration
//
//   - [NEDNSProxyProviderProtocol.ProviderConfiguration]: A dictionary containing vendor-specific configuration parameters for a proxy provider.
//   - [NEDNSProxyProviderProtocol.SetProviderConfiguration]
//   - [NEDNSProxyProviderProtocol.ProviderBundleIdentifier]: A string containing the bundle identifier of the proxy provider to be used by this configuration.
//   - [NEDNSProxyProviderProtocol.SetProviderBundleIdentifier]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProviderProtocol
type NEDNSProxyProviderProtocol struct {
	NEVPNProtocol
}

// NEDNSProxyProviderProtocolFromID constructs a [NEDNSProxyProviderProtocol] from an objc.ID.
//
// Configuration parameters for a DNS proxy.
func NEDNSProxyProviderProtocolFromID(id objc.ID) NEDNSProxyProviderProtocol {
	return NEDNSProxyProviderProtocol{NEVPNProtocol: NEVPNProtocolFromID(id)}
}

// NOTE: NEDNSProxyProviderProtocol adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEDNSProxyProviderProtocol] class.
//
// # Accessing the DNS proxy configuration
//
//   - [INEDNSProxyProviderProtocol.ProviderConfiguration]: A dictionary containing vendor-specific configuration parameters for a proxy provider.
//   - [INEDNSProxyProviderProtocol.SetProviderConfiguration]
//   - [INEDNSProxyProviderProtocol.ProviderBundleIdentifier]: A string containing the bundle identifier of the proxy provider to be used by this configuration.
//   - [INEDNSProxyProviderProtocol.SetProviderBundleIdentifier]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProviderProtocol
type INEDNSProxyProviderProtocol interface {
	INEVPNProtocol

	// Topic: Accessing the DNS proxy configuration

	// A dictionary containing vendor-specific configuration parameters for a proxy provider.
	ProviderConfiguration() foundation.INSDictionary
	SetProviderConfiguration(value foundation.INSDictionary)
	// A string containing the bundle identifier of the proxy provider to be used by this configuration.
	ProviderBundleIdentifier() string
	SetProviderBundleIdentifier(value string)
}

// Init initializes the instance.
func (d NEDNSProxyProviderProtocol) Init() NEDNSProxyProviderProtocol {
	rv := objc.Send[NEDNSProxyProviderProtocol](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NEDNSProxyProviderProtocol) Autorelease() NEDNSProxyProviderProtocol {
	rv := objc.Send[NEDNSProxyProviderProtocol](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSProxyProviderProtocol creates a new NEDNSProxyProviderProtocol instance.
func NewNEDNSProxyProviderProtocol() NEDNSProxyProviderProtocol {
	class := getNEDNSProxyProviderProtocolClass()
	rv := objc.Send[NEDNSProxyProviderProtocol](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A dictionary containing vendor-specific configuration parameters for a
// proxy provider.
//
// # Discussion
//
// This dictionary is passed as-is through the `options` parameter when the
// framework starts a DNS proxy by calling the proxy’s
// [StartProxyWithOptionsCompletionHandler] function.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProviderProtocol/providerConfiguration
func (d NEDNSProxyProviderProtocol) ProviderConfiguration() foundation.INSDictionary {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("providerConfiguration"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (d NEDNSProxyProviderProtocol) SetProviderConfiguration(value foundation.INSDictionary) {
	objc.Send[struct{}](d.ID, objc.Sel("setProviderConfiguration:"), value)
}

// A string containing the bundle identifier of the proxy provider to be used
// by this configuration.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProviderProtocol/providerBundleIdentifier
func (d NEDNSProxyProviderProtocol) ProviderBundleIdentifier() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("providerBundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (d NEDNSProxyProviderProtocol) SetProviderBundleIdentifier(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setProviderBundleIdentifier:"), objc.String(value))
}
