// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEDNSOverHTTPSSettings] class.
var (
	_NEDNSOverHTTPSSettingsClass     NEDNSOverHTTPSSettingsClass
	_NEDNSOverHTTPSSettingsClassOnce sync.Once
)

func getNEDNSOverHTTPSSettingsClass() NEDNSOverHTTPSSettingsClass {
	_NEDNSOverHTTPSSettingsClassOnce.Do(func() {
		_NEDNSOverHTTPSSettingsClass = NEDNSOverHTTPSSettingsClass{class: objc.GetClass("NEDNSOverHTTPSSettings")}
	})
	return _NEDNSOverHTTPSSettingsClass
}

// GetNEDNSOverHTTPSSettingsClass returns the class object for NEDNSOverHTTPSSettings.
func GetNEDNSOverHTTPSSettingsClass() NEDNSOverHTTPSSettingsClass {
	return getNEDNSOverHTTPSSettingsClass()
}

type NEDNSOverHTTPSSettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEDNSOverHTTPSSettingsClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEDNSOverHTTPSSettingsClass) Alloc() NEDNSOverHTTPSSettings {
	rv := objc.Send[NEDNSOverHTTPSSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The DNS resolver settings for a DNS-over-HTTPS server.
//
// # Configuring server properties
//
//   - [NEDNSOverHTTPSSettings.ServerURL]: The URL of a DNS-over-HTTPS server.
//   - [NEDNSOverHTTPSSettings.SetServerURL]
//
// # Configuring client properties
//
//   - [NEDNSOverHTTPSSettings.IdentityReference]: A persistent keychain reference to a keychain item containing the certificate and private key components of the DNS client credential.
//   - [NEDNSOverHTTPSSettings.SetIdentityReference]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverHTTPSSettings
type NEDNSOverHTTPSSettings struct {
	NEDNSSettings
}

// NEDNSOverHTTPSSettingsFromID constructs a [NEDNSOverHTTPSSettings] from an objc.ID.
//
// The DNS resolver settings for a DNS-over-HTTPS server.
func NEDNSOverHTTPSSettingsFromID(id objc.ID) NEDNSOverHTTPSSettings {
	return NEDNSOverHTTPSSettings{NEDNSSettings: NEDNSSettingsFromID(id)}
}

// NOTE: NEDNSOverHTTPSSettings adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEDNSOverHTTPSSettings] class.
//
// # Configuring server properties
//
//   - [INEDNSOverHTTPSSettings.ServerURL]: The URL of a DNS-over-HTTPS server.
//   - [INEDNSOverHTTPSSettings.SetServerURL]
//
// # Configuring client properties
//
//   - [INEDNSOverHTTPSSettings.IdentityReference]: A persistent keychain reference to a keychain item containing the certificate and private key components of the DNS client credential.
//   - [INEDNSOverHTTPSSettings.SetIdentityReference]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverHTTPSSettings
type INEDNSOverHTTPSSettings interface {
	INEDNSSettings

	// Topic: Configuring server properties

	// The URL of a DNS-over-HTTPS server.
	ServerURL() foundation.INSURL
	SetServerURL(value foundation.INSURL)

	// Topic: Configuring client properties

	// A persistent keychain reference to a keychain item containing the certificate and private key components of the DNS client credential.
	IdentityReference() foundation.INSData
	SetIdentityReference(value foundation.INSData)
}

// Init initializes the instance.
func (d NEDNSOverHTTPSSettings) Init() NEDNSOverHTTPSSettings {
	rv := objc.Send[NEDNSOverHTTPSSettings](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NEDNSOverHTTPSSettings) Autorelease() NEDNSOverHTTPSSettings {
	rv := objc.Send[NEDNSOverHTTPSSettings](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSOverHTTPSSettings creates a new NEDNSOverHTTPSSettings instance.
func NewNEDNSOverHTTPSSettings() NEDNSOverHTTPSSettings {
	class := getNEDNSOverHTTPSSettingsClass()
	rv := objc.Send[NEDNSOverHTTPSSettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initialize the [NEDNSSetting] object.
//
// servers: An array of DNS server IP address strings. These IP addresses can be a
// mixture of IPv4 and IPv6 addresses.
//
// # Return Value
//
// The initialized [NEDNSSettings] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/init(servers:)
func NewDNSOverHTTPSSettingsWithServers(servers []string) NEDNSOverHTTPSSettings {
	instance := getNEDNSOverHTTPSSettingsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithServers:"), objectivec.StringSliceToNSArray(servers))
	return NEDNSOverHTTPSSettingsFromID(rv)
}

// The URL of a DNS-over-HTTPS server.
//
// # Discussion
//
// The URL should use the URI template format defined by [RFC 8484], for
// example `//dnsserver.ExampleXCUIElementTypeNet()/dns-query`.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverHTTPSSettings/serverURL
//
// [RFC 8484]: https://tools.ietf.org/html/rfc8484
func (d NEDNSOverHTTPSSettings) ServerURL() foundation.INSURL {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("serverURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (d NEDNSOverHTTPSSettings) SetServerURL(value foundation.INSURL) {
	objc.Send[struct{}](d.ID, objc.Sel("setServerURL:"), value)
}

// A persistent keychain reference to a keychain item containing the
// certificate and private key components of the DNS client credential.
//
// # Discussion
//
// The keychain item must have the [kSecClassIdentity] class.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverHTTPSSettings/identityReference
//
// [kSecClassIdentity]: https://developer.apple.com/documentation/Security/kSecClassIdentity
func (d NEDNSOverHTTPSSettings) IdentityReference() foundation.INSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("identityReference"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (d NEDNSOverHTTPSSettings) SetIdentityReference(value foundation.INSData) {
	objc.Send[struct{}](d.ID, objc.Sel("setIdentityReference:"), value)
}
