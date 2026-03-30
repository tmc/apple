// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEDNSOverTLSSettings] class.
var (
	_NEDNSOverTLSSettingsClass     NEDNSOverTLSSettingsClass
	_NEDNSOverTLSSettingsClassOnce sync.Once
)

func getNEDNSOverTLSSettingsClass() NEDNSOverTLSSettingsClass {
	_NEDNSOverTLSSettingsClassOnce.Do(func() {
		_NEDNSOverTLSSettingsClass = NEDNSOverTLSSettingsClass{class: objc.GetClass("NEDNSOverTLSSettings")}
	})
	return _NEDNSOverTLSSettingsClass
}

// GetNEDNSOverTLSSettingsClass returns the class object for NEDNSOverTLSSettings.
func GetNEDNSOverTLSSettingsClass() NEDNSOverTLSSettingsClass {
	return getNEDNSOverTLSSettingsClass()
}

type NEDNSOverTLSSettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEDNSOverTLSSettingsClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEDNSOverTLSSettingsClass) Alloc() NEDNSOverTLSSettings {
	rv := objc.Send[NEDNSOverTLSSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The DNS resolver settings for a DNS-over-TLS server.
//
// # Configuring server properties
//
//   - [NEDNSOverTLSSettings.ServerName]: The TLS name of a DNS-over-TLS server.
//   - [NEDNSOverTLSSettings.SetServerName]
//
// # Configuring client properties
//
//   - [NEDNSOverTLSSettings.IdentityReference]: A persistent keychain reference to a keychain item containing the certificate and private key components of the DNS client credential.
//   - [NEDNSOverTLSSettings.SetIdentityReference]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverTLSSettings
type NEDNSOverTLSSettings struct {
	NEDNSSettings
}

// NEDNSOverTLSSettingsFromID constructs a [NEDNSOverTLSSettings] from an objc.ID.
//
// The DNS resolver settings for a DNS-over-TLS server.
func NEDNSOverTLSSettingsFromID(id objc.ID) NEDNSOverTLSSettings {
	return NEDNSOverTLSSettings{NEDNSSettings: NEDNSSettingsFromID(id)}
}

// NOTE: NEDNSOverTLSSettings adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEDNSOverTLSSettings] class.
//
// # Configuring server properties
//
//   - [INEDNSOverTLSSettings.ServerName]: The TLS name of a DNS-over-TLS server.
//   - [INEDNSOverTLSSettings.SetServerName]
//
// # Configuring client properties
//
//   - [INEDNSOverTLSSettings.IdentityReference]: A persistent keychain reference to a keychain item containing the certificate and private key components of the DNS client credential.
//   - [INEDNSOverTLSSettings.SetIdentityReference]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverTLSSettings
type INEDNSOverTLSSettings interface {
	INEDNSSettings

	// Topic: Configuring server properties

	// The TLS name of a DNS-over-TLS server.
	ServerName() string
	SetServerName(value string)

	// Topic: Configuring client properties

	// A persistent keychain reference to a keychain item containing the certificate and private key components of the DNS client credential.
	IdentityReference() foundation.INSData
	SetIdentityReference(value foundation.INSData)
}

// Init initializes the instance.
func (d NEDNSOverTLSSettings) Init() NEDNSOverTLSSettings {
	rv := objc.Send[NEDNSOverTLSSettings](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NEDNSOverTLSSettings) Autorelease() NEDNSOverTLSSettings {
	rv := objc.Send[NEDNSOverTLSSettings](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSOverTLSSettings creates a new NEDNSOverTLSSettings instance.
func NewNEDNSOverTLSSettings() NEDNSOverTLSSettings {
	class := getNEDNSOverTLSSettingsClass()
	rv := objc.Send[NEDNSOverTLSSettings](objc.ID(class.class), objc.Sel("new"))
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
func NewDNSOverTLSSettingsWithServers(servers []string) NEDNSOverTLSSettings {
	instance := getNEDNSOverTLSSettingsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithServers:"), objectivec.StringSliceToNSArray(servers))
	return NEDNSOverTLSSettingsFromID(rv)
}

// The TLS name of a DNS-over-TLS server.
//
// # Discussion
//
// The server will be accessed over TCP port 853, as defined in [RFC 7858].
// The server name is used for TLS validation.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverTLSSettings/serverName
//
// [RFC 7858]: https://tools.ietf.org/html/rfc7858
func (d NEDNSOverTLSSettings) ServerName() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("serverName"))
	return foundation.NSStringFromID(rv).String()
}
func (d NEDNSOverTLSSettings) SetServerName(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setServerName:"), objc.String(value))
}

// A persistent keychain reference to a keychain item containing the
// certificate and private key components of the DNS client credential.
//
// # Discussion
//
// The keychain item must have the [kSecClassIdentity] class.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverTLSSettings/identityReference
//
// [kSecClassIdentity]: https://developer.apple.com/documentation/Security/kSecClassIdentity
func (d NEDNSOverTLSSettings) IdentityReference() foundation.INSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("identityReference"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (d NEDNSOverTLSSettings) SetIdentityReference(value foundation.INSData) {
	objc.Send[struct{}](d.ID, objc.Sel("setIdentityReference:"), value)
}
