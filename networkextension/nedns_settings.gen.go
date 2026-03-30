// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEDNSSettings] class.
var (
	_NEDNSSettingsClass     NEDNSSettingsClass
	_NEDNSSettingsClassOnce sync.Once
)

func getNEDNSSettingsClass() NEDNSSettingsClass {
	_NEDNSSettingsClassOnce.Do(func() {
		_NEDNSSettingsClass = NEDNSSettingsClass{class: objc.GetClass("NEDNSSettings")}
	})
	return _NEDNSSettingsClass
}

// GetNEDNSSettingsClass returns the class object for NEDNSSettings.
func GetNEDNSSettingsClass() NEDNSSettingsClass {
	return getNEDNSSettingsClass()
}

type NEDNSSettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEDNSSettingsClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEDNSSettingsClass) Alloc() NEDNSSettings {
	rv := objc.Send[NEDNSSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The DNS resolver settings of a network tunnel or a system-wide
// configuration.
//
// # Initializing DNS settings
//
//   - [NEDNSSettings.InitWithServers]: Initialize the [NEDNSSetting] object.
//
// # Accessing DNS properties
//
//   - [NEDNSSettings.Servers]: The DNS server IP addresses.
//   - [NEDNSSettings.SearchDomains]: A list of domain strings used to fully qualify single-label host names.
//   - [NEDNSSettings.SetSearchDomains]
//   - [NEDNSSettings.DomainName]: The primary domain of the tunnel.
//   - [NEDNSSettings.SetDomainName]
//   - [NEDNSSettings.MatchDomains]: A list of domain strings used to determine which DNS queries will use the DNS resolver settings contained in this object.
//   - [NEDNSSettings.SetMatchDomains]
//   - [NEDNSSettings.MatchDomainsNoSearch]: A Boolean that specifies if the domains in the `matchDomains` list should not be appended to the resolver’s list of search domains.
//   - [NEDNSSettings.SetMatchDomainsNoSearch]
//   - [NEDNSSettings.DnsProtocol]: The DNS protocol used by the server, such as HTTPS or TLS.
//
// # Instance Properties
//
//   - [NEDNSSettings.AllowFailover]
//   - [NEDNSSettings.SetAllowFailover]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings
type NEDNSSettings struct {
	objectivec.Object
}

// NEDNSSettingsFromID constructs a [NEDNSSettings] from an objc.ID.
//
// The DNS resolver settings of a network tunnel or a system-wide
// configuration.
func NEDNSSettingsFromID(id objc.ID) NEDNSSettings {
	return NEDNSSettings{objectivec.Object{ID: id}}
}

// NOTE: NEDNSSettings adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEDNSSettings] class.
//
// # Initializing DNS settings
//
//   - [INEDNSSettings.InitWithServers]: Initialize the [NEDNSSetting] object.
//
// # Accessing DNS properties
//
//   - [INEDNSSettings.Servers]: The DNS server IP addresses.
//   - [INEDNSSettings.SearchDomains]: A list of domain strings used to fully qualify single-label host names.
//   - [INEDNSSettings.SetSearchDomains]
//   - [INEDNSSettings.DomainName]: The primary domain of the tunnel.
//   - [INEDNSSettings.SetDomainName]
//   - [INEDNSSettings.MatchDomains]: A list of domain strings used to determine which DNS queries will use the DNS resolver settings contained in this object.
//   - [INEDNSSettings.SetMatchDomains]
//   - [INEDNSSettings.MatchDomainsNoSearch]: A Boolean that specifies if the domains in the `matchDomains` list should not be appended to the resolver’s list of search domains.
//   - [INEDNSSettings.SetMatchDomainsNoSearch]
//   - [INEDNSSettings.DnsProtocol]: The DNS protocol used by the server, such as HTTPS or TLS.
//
// # Instance Properties
//
//   - [INEDNSSettings.AllowFailover]
//   - [INEDNSSettings.SetAllowFailover]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings
type INEDNSSettings interface {
	objectivec.IObject

	// Topic: Initializing DNS settings

	// Initialize the [NEDNSSetting] object.
	InitWithServers(servers []string) NEDNSSettings

	// Topic: Accessing DNS properties

	// The DNS server IP addresses.
	Servers() []string
	// A list of domain strings used to fully qualify single-label host names.
	SearchDomains() []string
	SetSearchDomains(value []string)
	// The primary domain of the tunnel.
	DomainName() string
	SetDomainName(value string)
	// A list of domain strings used to determine which DNS queries will use the DNS resolver settings contained in this object.
	MatchDomains() []string
	SetMatchDomains(value []string)
	// A Boolean that specifies if the domains in the `matchDomains` list should not be appended to the resolver’s list of search domains.
	MatchDomainsNoSearch() bool
	SetMatchDomainsNoSearch(value bool)
	// The DNS protocol used by the server, such as HTTPS or TLS.
	DnsProtocol() NEDNSProtocol

	// Topic: Instance Properties

	AllowFailover() bool
	SetAllowFailover(value bool)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (d NEDNSSettings) Init() NEDNSSettings {
	rv := objc.Send[NEDNSSettings](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NEDNSSettings) Autorelease() NEDNSSettings {
	rv := objc.Send[NEDNSSettings](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSSettings creates a new NEDNSSettings instance.
func NewNEDNSSettings() NEDNSSettings {
	class := getNEDNSSettingsClass()
	rv := objc.Send[NEDNSSettings](objc.ID(class.class), objc.Sel("new"))
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
func NewDNSSettingsWithServers(servers []string) NEDNSSettings {
	instance := getNEDNSSettingsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithServers:"), objectivec.StringSliceToNSArray(servers))
	return NEDNSSettingsFromID(rv)
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
func (d NEDNSSettings) InitWithServers(servers []string) NEDNSSettings {
	rv := objc.Send[NEDNSSettings](d.ID, objc.Sel("initWithServers:"), objectivec.StringSliceToNSArray(servers))
	return rv
}
func (d NEDNSSettings) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The DNS server IP addresses.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/servers
func (d NEDNSSettings) Servers() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("servers"))
	return objc.ConvertSliceToStrings(rv)
}

// A list of domain strings used to fully qualify single-label host names.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/searchDomains
func (d NEDNSSettings) SearchDomains() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("searchDomains"))
	return objc.ConvertSliceToStrings(rv)
}
func (d NEDNSSettings) SetSearchDomains(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setSearchDomains:"), objectivec.StringSliceToNSArray(value))
}

// The primary domain of the tunnel.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/domainName
func (d NEDNSSettings) DomainName() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("domainName"))
	return foundation.NSStringFromID(rv).String()
}
func (d NEDNSSettings) SetDomainName(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setDomainName:"), objc.String(value))
}

// A list of domain strings used to determine which DNS queries will use the
// DNS resolver settings contained in this object.
//
// # Discussion
//
// This property is used to create a “split DNS” configuration, where only
// hosts in certain domains are resolved using the tunnel’s DNS resolver
// settings. Hosts not in one of the domains in this list are resolved using
// the system’s default resolver.
//
// If `matchDomains` contains the empty string it becomes the default domain.
// This is how a split-tunnel configuration can direct all DNS queries first
// to the VPN DNS servers before the primary DNS servers.
//
// If the VPN tunnel becomes the network’s default route, the servers listed
// earlier by [NEDNSSettings] become the default resolver and the
// `matchDomains` list is ignored.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/matchDomains
func (d NEDNSSettings) MatchDomains() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("matchDomains"))
	return objc.ConvertSliceToStrings(rv)
}
func (d NEDNSSettings) SetMatchDomains(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setMatchDomains:"), objectivec.StringSliceToNSArray(value))
}

// A Boolean that specifies if the domains in the `matchDomains` list should
// not be appended to the resolver’s list of search domains.
//
// # Discussion
//
// The default value is false.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/matchDomainsNoSearch
func (d NEDNSSettings) MatchDomainsNoSearch() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("matchDomainsNoSearch"))
	return rv
}
func (d NEDNSSettings) SetMatchDomainsNoSearch(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setMatchDomainsNoSearch:"), value)
}

// The DNS protocol used by the server, such as HTTPS or TLS.
//
// # Discussion
//
// By default, an [NEDNSSettings] object will use [NEDNSProtocolCleartext]. In
// order to use encryption, create an [NEDNSOverHTTPSSettings] or
// [NEDNSOverTLSSettings] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/dnsProtocol
func (d NEDNSSettings) DnsProtocol() NEDNSProtocol {
	rv := objc.Send[NEDNSProtocol](d.ID, objc.Sel("dnsProtocol"))
	return NEDNSProtocol(rv)
}

// # Discussion
//
// A boolean indicating if failover to the default system resolver is
// permitted on resolution failure.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/allowFailover
func (d NEDNSSettings) AllowFailover() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("allowFailover"))
	return rv
}
func (d NEDNSSettings) SetAllowFailover(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setAllowFailover:"), value)
}
