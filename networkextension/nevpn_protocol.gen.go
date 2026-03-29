// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEVPNProtocol] class.
var (
	_NEVPNProtocolClass     NEVPNProtocolClass
	_NEVPNProtocolClassOnce sync.Once
)

func getNEVPNProtocolClass() NEVPNProtocolClass {
	_NEVPNProtocolClassOnce.Do(func() {
		_NEVPNProtocolClass = NEVPNProtocolClass{class: objc.GetClass("NEVPNProtocol")}
	})
	return _NEVPNProtocolClass
}

// GetNEVPNProtocolClass returns the class object for NEVPNProtocol.
func GetNEVPNProtocolClass() NEVPNProtocolClass {
	return getNEVPNProtocolClass()
}

type NEVPNProtocolClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEVPNProtocolClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEVPNProtocolClass) Alloc() NEVPNProtocol {
	rv := objc.Send[NEVPNProtocol](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Settings common to both IKEv2 and IPsec VPN configurations.
//
// # Overview
// 
// The [NEVPNProtocol] class is an abstract base class with one subclass for
// each type of supported VPN configuration. This class provides properties
// for configuring the VPN, authenticating network connections, and routing
// network traffic. You can include all network traffic, with some exceptions,
// and selectively exclude types of network traffic.
// 
// Instances of this class are thread-safe.
//
// # Configuring the VPN
//
//   - [NEVPNProtocol.ServerAddress]: The address of the VPN server.
//   - [NEVPNProtocol.SetServerAddress]
//   - [NEVPNProtocol.DisconnectOnSleep]: A Boolean value that indicates whether the VPN disconnects when the device sleeps.
//   - [NEVPNProtocol.SetDisconnectOnSleep]
//   - [NEVPNProtocol.ProxySettings]: The proxy settings to use for HTTP and HTTPS connections that route through the VPN.
//   - [NEVPNProtocol.SetProxySettings]
//
// # Authenticating the user
//
//   - [NEVPNProtocol.Username]: The user name component of the tunneling protocol authentication credential.
//   - [NEVPNProtocol.SetUsername]
//   - [NEVPNProtocol.PasswordReference]: A persistent keychain reference to a keychain item containing the password component of the tunneling protocol authentication credential.
//   - [NEVPNProtocol.SetPasswordReference]
//   - [NEVPNProtocol.IdentityReference]: A persistent keychain reference to a keychain item containing the certificate and private key components of the tunneling protocol authentication credential.
//   - [NEVPNProtocol.SetIdentityReference]
//   - [NEVPNProtocol.IdentityData]: The certificate and private key components of the tunneling protocol authentication credential, in PKCS12 format.
//   - [NEVPNProtocol.SetIdentityData]
//   - [NEVPNProtocol.IdentityDataPassword]: The password for the PKCS12 tunneling protocol authentication credentials.
//   - [NEVPNProtocol.SetIdentityDataPassword]
//
// # Routing network traffic
//
//   - [NEVPNProtocol.IncludeAllNetworks]: A Boolean value that indicates whether the system sends most network traffic over the tunnel.
//   - [NEVPNProtocol.SetIncludeAllNetworks]
//   - [NEVPNProtocol.ExcludeAPNs]: A Boolean value that indicates whether the system excludes all APNs network traffic from the tunnel.
//   - [NEVPNProtocol.SetExcludeAPNs]
//   - [NEVPNProtocol.ExcludeCellularServices]: A Boolean value that indicates whether the system excludes all cellular services network traffic from the tunnel.
//   - [NEVPNProtocol.SetExcludeCellularServices]
//   - [NEVPNProtocol.ExcludeLocalNetworks]: A Boolean value that indicates whether the system excludes all traffic destined for local networks from the tunnel.
//   - [NEVPNProtocol.SetExcludeLocalNetworks]
//   - [NEVPNProtocol.EnforceRoutes]: A Boolean value that indicates whether route rules for the tunnel take precedence over any locally defined routes.
//   - [NEVPNProtocol.SetEnforceRoutes]
//
// # Instance Properties
//
//   - [NEVPNProtocol.ExcludeDeviceCommunication]
//   - [NEVPNProtocol.SetExcludeDeviceCommunication]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol
type NEVPNProtocol struct {
	objectivec.Object
}

// NEVPNProtocolFromID constructs a [NEVPNProtocol] from an objc.ID.
//
// Settings common to both IKEv2 and IPsec VPN configurations.
func NEVPNProtocolFromID(id objc.ID) NEVPNProtocol {
	return NEVPNProtocol{objectivec.Object{ID: id}}
}
// NOTE: NEVPNProtocol adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEVPNProtocol] class.
//
// # Configuring the VPN
//
//   - [INEVPNProtocol.ServerAddress]: The address of the VPN server.
//   - [INEVPNProtocol.SetServerAddress]
//   - [INEVPNProtocol.DisconnectOnSleep]: A Boolean value that indicates whether the VPN disconnects when the device sleeps.
//   - [INEVPNProtocol.SetDisconnectOnSleep]
//   - [INEVPNProtocol.ProxySettings]: The proxy settings to use for HTTP and HTTPS connections that route through the VPN.
//   - [INEVPNProtocol.SetProxySettings]
//
// # Authenticating the user
//
//   - [INEVPNProtocol.Username]: The user name component of the tunneling protocol authentication credential.
//   - [INEVPNProtocol.SetUsername]
//   - [INEVPNProtocol.PasswordReference]: A persistent keychain reference to a keychain item containing the password component of the tunneling protocol authentication credential.
//   - [INEVPNProtocol.SetPasswordReference]
//   - [INEVPNProtocol.IdentityReference]: A persistent keychain reference to a keychain item containing the certificate and private key components of the tunneling protocol authentication credential.
//   - [INEVPNProtocol.SetIdentityReference]
//   - [INEVPNProtocol.IdentityData]: The certificate and private key components of the tunneling protocol authentication credential, in PKCS12 format.
//   - [INEVPNProtocol.SetIdentityData]
//   - [INEVPNProtocol.IdentityDataPassword]: The password for the PKCS12 tunneling protocol authentication credentials.
//   - [INEVPNProtocol.SetIdentityDataPassword]
//
// # Routing network traffic
//
//   - [INEVPNProtocol.IncludeAllNetworks]: A Boolean value that indicates whether the system sends most network traffic over the tunnel.
//   - [INEVPNProtocol.SetIncludeAllNetworks]
//   - [INEVPNProtocol.ExcludeAPNs]: A Boolean value that indicates whether the system excludes all APNs network traffic from the tunnel.
//   - [INEVPNProtocol.SetExcludeAPNs]
//   - [INEVPNProtocol.ExcludeCellularServices]: A Boolean value that indicates whether the system excludes all cellular services network traffic from the tunnel.
//   - [INEVPNProtocol.SetExcludeCellularServices]
//   - [INEVPNProtocol.ExcludeLocalNetworks]: A Boolean value that indicates whether the system excludes all traffic destined for local networks from the tunnel.
//   - [INEVPNProtocol.SetExcludeLocalNetworks]
//   - [INEVPNProtocol.EnforceRoutes]: A Boolean value that indicates whether route rules for the tunnel take precedence over any locally defined routes.
//   - [INEVPNProtocol.SetEnforceRoutes]
//
// # Instance Properties
//
//   - [INEVPNProtocol.ExcludeDeviceCommunication]
//   - [INEVPNProtocol.SetExcludeDeviceCommunication]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol
type INEVPNProtocol interface {
	objectivec.IObject

	// Topic: Configuring the VPN

	// The address of the VPN server.
	ServerAddress() string
	SetServerAddress(value string)
	// A Boolean value that indicates whether the VPN disconnects when the device sleeps.
	DisconnectOnSleep() bool
	SetDisconnectOnSleep(value bool)
	// The proxy settings to use for HTTP and HTTPS connections that route through the VPN.
	ProxySettings() INEProxySettings
	SetProxySettings(value INEProxySettings)

	// Topic: Authenticating the user

	// The user name component of the tunneling protocol authentication credential.
	Username() string
	SetUsername(value string)
	// A persistent keychain reference to a keychain item containing the password component of the tunneling protocol authentication credential.
	PasswordReference() foundation.INSData
	SetPasswordReference(value foundation.INSData)
	// A persistent keychain reference to a keychain item containing the certificate and private key components of the tunneling protocol authentication credential.
	IdentityReference() foundation.INSData
	SetIdentityReference(value foundation.INSData)
	// The certificate and private key components of the tunneling protocol authentication credential, in PKCS12 format.
	IdentityData() foundation.INSData
	SetIdentityData(value foundation.INSData)
	// The password for the PKCS12 tunneling protocol authentication credentials.
	IdentityDataPassword() string
	SetIdentityDataPassword(value string)

	// Topic: Routing network traffic

	// A Boolean value that indicates whether the system sends most network traffic over the tunnel.
	IncludeAllNetworks() bool
	SetIncludeAllNetworks(value bool)
	// A Boolean value that indicates whether the system excludes all APNs network traffic from the tunnel.
	ExcludeAPNs() bool
	SetExcludeAPNs(value bool)
	// A Boolean value that indicates whether the system excludes all cellular services network traffic from the tunnel.
	ExcludeCellularServices() bool
	SetExcludeCellularServices(value bool)
	// A Boolean value that indicates whether the system excludes all traffic destined for local networks from the tunnel.
	ExcludeLocalNetworks() bool
	SetExcludeLocalNetworks(value bool)
	// A Boolean value that indicates whether route rules for the tunnel take precedence over any locally defined routes.
	EnforceRoutes() bool
	SetEnforceRoutes(value bool)

	// Topic: Instance Properties

	ExcludeDeviceCommunication() bool
	SetExcludeDeviceCommunication(value bool)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (v NEVPNProtocol) Init() NEVPNProtocol {
	rv := objc.Send[NEVPNProtocol](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v NEVPNProtocol) Autorelease() NEVPNProtocol {
	rv := objc.Send[NEVPNProtocol](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEVPNProtocol creates a new NEVPNProtocol instance.
func NewNEVPNProtocol() NEVPNProtocol {
	class := getNEVPNProtocolClass()
	rv := objc.Send[NEVPNProtocol](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v NEVPNProtocol) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](v.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The address of the VPN server.
//
// # Discussion
// 
// The format of the value of this property depends on the type of VPN
// protocol in use. For example, for IPSec the value should be a hostname or
// an IP address. For a custom SSL-VPN protocol the value may be a URL. The
// only requirement imposed by the Network Extension framework is that this
// property must have a non-`nil` string value for the protocol configuration
// to be valid.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/serverAddress
func (v NEVPNProtocol) ServerAddress() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("serverAddress"))
	return foundation.NSStringFromID(rv).String()
}
func (v NEVPNProtocol) SetServerAddress(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("setServerAddress:"), objc.String(value))
}
// A Boolean value that indicates whether the VPN disconnects when the device
// sleeps.
//
// # Discussion
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/disconnectOnSleep
func (v NEVPNProtocol) DisconnectOnSleep() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("disconnectOnSleep"))
	return rv
}
func (v NEVPNProtocol) SetDisconnectOnSleep(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setDisconnectOnSleep:"), value)
}
// The proxy settings to use for HTTP and HTTPS connections that route through
// the VPN.
//
// # Discussion
// 
// While operating under an established VPN tunnel, HTTP and HTTPS connections
// inside the tunnel use the given proxy settings.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/proxySettings
func (v NEVPNProtocol) ProxySettings() INEProxySettings {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("proxySettings"))
	return NEProxySettingsFromID(objc.ID(rv))
}
func (v NEVPNProtocol) SetProxySettings(value INEProxySettings) {
	objc.Send[struct{}](v.ID, objc.Sel("setProxySettings:"), value)
}
// The user name component of the tunneling protocol authentication
// credential.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/username
func (v NEVPNProtocol) Username() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("username"))
	return foundation.NSStringFromID(rv).String()
}
func (v NEVPNProtocol) SetUsername(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("setUsername:"), objc.String(value))
}
// A persistent keychain reference to a keychain item containing the password
// component of the tunneling protocol authentication credential.
//
// # Discussion
// 
// The keychain item must have the kSecClassGenericPassword class.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/passwordReference
func (v NEVPNProtocol) PasswordReference() foundation.INSData {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("passwordReference"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (v NEVPNProtocol) SetPasswordReference(value foundation.INSData) {
	objc.Send[struct{}](v.ID, objc.Sel("setPasswordReference:"), value)
}
// A persistent keychain reference to a keychain item containing the
// certificate and private key components of the tunneling protocol
// authentication credential.
//
// # Discussion
// 
// The keychain item must have the [kSecClassIdentity] class. In macOS, the
// system ignores this property for [NEVPNProtocolIPSec] objects. On iOS, the
// system ignores this property for [NEVPNProtocolIPSec] and
// [NEVPNProtocolIKEv2] objects. In these cases where the system ingores this
// property, set the identity using the [IdentityData] property.
//
// [kSecClassIdentity]: https://developer.apple.com/documentation/Security/kSecClassIdentity
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/identityReference
func (v NEVPNProtocol) IdentityReference() foundation.INSData {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("identityReference"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (v NEVPNProtocol) SetIdentityReference(value foundation.INSData) {
	objc.Send[struct{}](v.ID, objc.Sel("setIdentityReference:"), value)
}
// The certificate and private key components of the tunneling protocol
// authentication credential, in PKCS12 format.
//
// # Discussion
// 
// In macOS, the system ignores this property for [NEVPNProtocolIKEv2] and
// [NETunnelProviderProtocol] objects. On iOS, the system ignores this
// property for [NETunnelProviderProtocol] objects. In cases where the system
// ignores this property, set the identity using the [IdentityReference]
// property.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/identityData
func (v NEVPNProtocol) IdentityData() foundation.INSData {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("identityData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (v NEVPNProtocol) SetIdentityData(value foundation.INSData) {
	objc.Send[struct{}](v.ID, objc.Sel("setIdentityData:"), value)
}
// The password for the PKCS12 tunneling protocol authentication credentials.
//
// # Discussion
// 
// If the PKCS12 data set in the [IdentityData] property uses a password for
// encryption, you must specify the password here.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/identityDataPassword
func (v NEVPNProtocol) IdentityDataPassword() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("identityDataPassword"))
	return foundation.NSStringFromID(rv).String()
}
func (v NEVPNProtocol) SetIdentityDataPassword(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("setIdentityDataPassword:"), objc.String(value))
}
// A Boolean value that indicates whether the system sends most network
// traffic over the tunnel.
//
// # Discussion
// 
// If this property is [true], the system routes network traffic through the
// tunnel except traffic for designated system services necessary for
// maintaining expected device functionality.
// 
// You can exclude some types of traffic using the [ExcludeAPNs],
// [ExcludeLocalNetworks], and [ExcludeCellularServices] properties in
// combination with this property. The system always excludes the following
// network traffic from the tunnel regardless of this property value:
// 
// - Network control plane traffic that maintains a device’s connection to
// the local network, such as DHCP. - Captive portal negotiation traffic that
// authorizes a device with a Wi-Fi hotspot. - Certain cellular services
// traffic that uses the cellular network only, such as VoLTE. - Traffic that
// communicates with a companion device, such as an Apple Watch.
// 
// [NETransparentProxyManager] doesn’t support this property. The default
// value for this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/includeAllNetworks
func (v NEVPNProtocol) IncludeAllNetworks() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("includeAllNetworks"))
	return rv
}
func (v NEVPNProtocol) SetIncludeAllNetworks(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setIncludeAllNetworks:"), value)
}
// A Boolean value that indicates whether the system excludes all APNs network
// traffic from the tunnel.
//
// # Discussion
// 
// If this property is [true], the system excludes Apple Push Notification
// services (APNs) traffic, but only when the [IncludeAllNetworks] property is
// also [true]. [NETransparentProxyManager] doesn’t support this property.
// 
// The default value for this property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/excludeAPNs
func (v NEVPNProtocol) ExcludeAPNs() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("excludeAPNs"))
	return rv
}
func (v NEVPNProtocol) SetExcludeAPNs(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setExcludeAPNs:"), value)
}
// A Boolean value that indicates whether the system excludes all cellular
// services network traffic from the tunnel.
//
// # Discussion
// 
// If this property is [true], the system excludes cellular services — such
// as Wi-Fi Calling, MMS, SMS, and Visual Voicemail — but only when the
// [IncludeAllNetworks] property is also [true]. This property doesn’t
// impact services that use the cellular network only — such as VoLTE —
// which the system automatically excludes. [NETransparentProxyManager]
// doesn’t support this property.
// 
// The default value for this property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/excludeCellularServices
func (v NEVPNProtocol) ExcludeCellularServices() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("excludeCellularServices"))
	return rv
}
func (v NEVPNProtocol) SetExcludeCellularServices(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setExcludeCellularServices:"), value)
}
// A Boolean value that indicates whether the system excludes all traffic
// destined for local networks from the tunnel.
//
// # Discussion
// 
// If this property is [true], the system excludes network connections to
// hosts on the local network — such as AirPlay, AirDrop, and CarPlay —
// but only when the [IncludeAllNetworks] or [EnforceRoutes] property is also
// [true]. [NETransparentProxyManager] doesn’t support this property.
// 
// The default value for this property is [false] in macOS and [true] in
// iOS`.`
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/excludeLocalNetworks
func (v NEVPNProtocol) ExcludeLocalNetworks() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("excludeLocalNetworks"))
	return rv
}
func (v NEVPNProtocol) SetExcludeLocalNetworks(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setExcludeLocalNetworks:"), value)
}
// A Boolean value that indicates whether route rules for the tunnel take
// precedence over any locally defined routes.
//
// # Discussion
// 
// If this property is [true] when the [IncludeAllNetworks] property is
// [false], the system scopes the included routes to the VPN and the excluded
// routes to the current primary network interface. This property supersedes
// the system routing table and scoping operations by apps.
// 
// If you set both the [EnforceRoutes] and [ExcludeLocalNetworks] properties
// to [true], the system excludes network connections to hosts on the local
// network.
// 
// [NETransparentProxyManager] doesn’t support this property. The default
// value for this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/enforceRoutes
func (v NEVPNProtocol) EnforceRoutes() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("enforceRoutes"))
	return rv
}
func (v NEVPNProtocol) SetEnforceRoutes(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setEnforceRoutes:"), value)
}
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/excludeDeviceCommunication
func (v NEVPNProtocol) ExcludeDeviceCommunication() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("excludeDeviceCommunication"))
	return rv
}
func (v NEVPNProtocol) SetExcludeDeviceCommunication(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setExcludeDeviceCommunication:"), value)
}

