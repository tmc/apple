// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEProxySettings] class.
var (
	_NEProxySettingsClass     NEProxySettingsClass
	_NEProxySettingsClassOnce sync.Once
)

func getNEProxySettingsClass() NEProxySettingsClass {
	_NEProxySettingsClassOnce.Do(func() {
		_NEProxySettingsClass = NEProxySettingsClass{class: objc.GetClass("NEProxySettings")}
	})
	return _NEProxySettingsClass
}

// GetNEProxySettingsClass returns the class object for NEProxySettings.
func GetNEProxySettingsClass() NEProxySettingsClass {
	return getNEProxySettingsClass()
}

type NEProxySettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEProxySettingsClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEProxySettingsClass) Alloc() NEProxySettings {
	rv := objc.Send[NEProxySettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// [NEProxySettings] contains HTTP proxy settings.
//
// # Overview
// 
// [NEProxySettings] is used in the context of a VPN configuration to specify
// the proxy that should be used for network traffic when the VPN is active.
// 
// Instances of this class are thread safe.
//
// # Accessing Automatic Proxy Properties
//
//   - [NEProxySettings.AutoProxyConfigurationEnabled]: A Boolean indicating if proxy auto-configuration is enabled.
//   - [NEProxySettings.SetAutoProxyConfigurationEnabled]
//   - [NEProxySettings.ProxyAutoConfigurationURL]: A URL specifying the location from where the Proxy Auto Configuration (PAC) script should be downloaded.
//   - [NEProxySettings.SetProxyAutoConfigurationURL]
//   - [NEProxySettings.ProxyAutoConfigurationJavaScript]: A string containing the Proxy Auto Configuration (PAC) JavaScript source code.
//   - [NEProxySettings.SetProxyAutoConfigurationJavaScript]
//
// # Accessing Manual Proxy Properties
//
//   - [NEProxySettings.HTTPEnabled]: A Boolean indicating if a static HTTP proxy will be used.
//   - [NEProxySettings.SetHTTPEnabled]
//   - [NEProxySettings.HTTPServer]: An [NEProxyServer](<doc://com.apple.networkextension/documentation/NetworkExtension/NEProxyServer>) object containing the static HTTP proxy server settings.
//   - [NEProxySettings.SetHTTPServer]
//   - [NEProxySettings.HTTPSEnabled]: A Boolean indicating if a static HTTPS proxy will be used.
//   - [NEProxySettings.SetHTTPSEnabled]
//   - [NEProxySettings.HTTPSServer]: An [NEProxyServer](<doc://com.apple.networkextension/documentation/NetworkExtension/NEProxyServer>) object containing the static HTTPS proxy server settings.
//   - [NEProxySettings.SetHTTPSServer]
//
// # Accessing General Proxy Properties
//
//   - [NEProxySettings.ExcludeSimpleHostnames]: A Boolean indicating if HTTP requests using single-label host names should be excluded from using the proxy settings.
//   - [NEProxySettings.SetExcludeSimpleHostnames]
//   - [NEProxySettings.ExceptionList]: An array of domain name patterns. If the destination host name of an HTTP connection matches one of these patterns then the proxy settings will not be used for the connection.
//   - [NEProxySettings.SetExceptionList]
//   - [NEProxySettings.MatchDomains]: An array of domain strings.
//   - [NEProxySettings.SetMatchDomains]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings
type NEProxySettings struct {
	objectivec.Object
}

// NEProxySettingsFromID constructs a [NEProxySettings] from an objc.ID.
//
// [NEProxySettings] contains HTTP proxy settings.
func NEProxySettingsFromID(id objc.ID) NEProxySettings {
	return NEProxySettings{objectivec.Object{ID: id}}
}
// NOTE: NEProxySettings adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEProxySettings] class.
//
// # Accessing Automatic Proxy Properties
//
//   - [INEProxySettings.AutoProxyConfigurationEnabled]: A Boolean indicating if proxy auto-configuration is enabled.
//   - [INEProxySettings.SetAutoProxyConfigurationEnabled]
//   - [INEProxySettings.ProxyAutoConfigurationURL]: A URL specifying the location from where the Proxy Auto Configuration (PAC) script should be downloaded.
//   - [INEProxySettings.SetProxyAutoConfigurationURL]
//   - [INEProxySettings.ProxyAutoConfigurationJavaScript]: A string containing the Proxy Auto Configuration (PAC) JavaScript source code.
//   - [INEProxySettings.SetProxyAutoConfigurationJavaScript]
//
// # Accessing Manual Proxy Properties
//
//   - [INEProxySettings.HTTPEnabled]: A Boolean indicating if a static HTTP proxy will be used.
//   - [INEProxySettings.SetHTTPEnabled]
//   - [INEProxySettings.HTTPServer]: An [NEProxyServer](<doc://com.apple.networkextension/documentation/NetworkExtension/NEProxyServer>) object containing the static HTTP proxy server settings.
//   - [INEProxySettings.SetHTTPServer]
//   - [INEProxySettings.HTTPSEnabled]: A Boolean indicating if a static HTTPS proxy will be used.
//   - [INEProxySettings.SetHTTPSEnabled]
//   - [INEProxySettings.HTTPSServer]: An [NEProxyServer](<doc://com.apple.networkextension/documentation/NetworkExtension/NEProxyServer>) object containing the static HTTPS proxy server settings.
//   - [INEProxySettings.SetHTTPSServer]
//
// # Accessing General Proxy Properties
//
//   - [INEProxySettings.ExcludeSimpleHostnames]: A Boolean indicating if HTTP requests using single-label host names should be excluded from using the proxy settings.
//   - [INEProxySettings.SetExcludeSimpleHostnames]
//   - [INEProxySettings.ExceptionList]: An array of domain name patterns. If the destination host name of an HTTP connection matches one of these patterns then the proxy settings will not be used for the connection.
//   - [INEProxySettings.SetExceptionList]
//   - [INEProxySettings.MatchDomains]: An array of domain strings.
//   - [INEProxySettings.SetMatchDomains]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings
type INEProxySettings interface {
	objectivec.IObject

	// Topic: Accessing Automatic Proxy Properties

	// A Boolean indicating if proxy auto-configuration is enabled.
	AutoProxyConfigurationEnabled() bool
	SetAutoProxyConfigurationEnabled(value bool)
	// A URL specifying the location from where the Proxy Auto Configuration (PAC) script should be downloaded.
	ProxyAutoConfigurationURL() foundation.INSURL
	SetProxyAutoConfigurationURL(value foundation.INSURL)
	// A string containing the Proxy Auto Configuration (PAC) JavaScript source code.
	ProxyAutoConfigurationJavaScript() string
	SetProxyAutoConfigurationJavaScript(value string)

	// Topic: Accessing Manual Proxy Properties

	// A Boolean indicating if a static HTTP proxy will be used.
	HTTPEnabled() bool
	SetHTTPEnabled(value bool)
	// An [NEProxyServer](<doc://com.apple.networkextension/documentation/NetworkExtension/NEProxyServer>) object containing the static HTTP proxy server settings.
	HTTPServer() INEProxyServer
	SetHTTPServer(value INEProxyServer)
	// A Boolean indicating if a static HTTPS proxy will be used.
	HTTPSEnabled() bool
	SetHTTPSEnabled(value bool)
	// An [NEProxyServer](<doc://com.apple.networkextension/documentation/NetworkExtension/NEProxyServer>) object containing the static HTTPS proxy server settings.
	HTTPSServer() INEProxyServer
	SetHTTPSServer(value INEProxyServer)

	// Topic: Accessing General Proxy Properties

	// A Boolean indicating if HTTP requests using single-label host names should be excluded from using the proxy settings.
	ExcludeSimpleHostnames() bool
	SetExcludeSimpleHostnames(value bool)
	// An array of domain name patterns. If the destination host name of an HTTP connection matches one of these patterns then the proxy settings will not be used for the connection.
	ExceptionList() []string
	SetExceptionList(value []string)
	// An array of domain strings.
	MatchDomains() []string
	SetMatchDomains(value []string)

	// The tunnel DNS settings.
	DnsSettings() INEDNSSettings
	SetDnsSettings(value INEDNSSettings)
	// The tunnel HTTP proxy settings.
	ProxySettings() INEProxySettings
	SetProxySettings(value INEProxySettings)
	// The IP address of the tunnel server.
	TunnelRemoteAddress() string
	SetTunnelRemoteAddress(value string)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p NEProxySettings) Init() NEProxySettings {
	rv := objc.Send[NEProxySettings](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NEProxySettings) Autorelease() NEProxySettings {
	rv := objc.Send[NEProxySettings](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEProxySettings creates a new NEProxySettings instance.
func NewNEProxySettings() NEProxySettings {
	class := getNEProxySettingsClass()
	rv := objc.Send[NEProxySettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (p NEProxySettings) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A Boolean indicating if proxy auto-configuration is enabled.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/autoProxyConfigurationEnabled
func (p NEProxySettings) AutoProxyConfigurationEnabled() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("autoProxyConfigurationEnabled"))
	return rv
}
func (p NEProxySettings) SetAutoProxyConfigurationEnabled(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAutoProxyConfigurationEnabled:"), value)
}
// A URL specifying the location from where the Proxy Auto Configuration (PAC)
// script should be downloaded.
//
// # Discussion
// 
// If [AutoProxyConfigurationEnabled] is set to [true] and
// [ProxyAutoConfigurationJavaScript] is set to nil then the system will
// download the PAC script from this location and execute the script to
// determine what proxies to use (if any) for HTTP and HTTPS connections.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/proxyAutoConfigurationURL
func (p NEProxySettings) ProxyAutoConfigurationURL() foundation.INSURL {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("proxyAutoConfigurationURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (p NEProxySettings) SetProxyAutoConfigurationURL(value foundation.INSURL) {
	objc.Send[struct{}](p.ID, objc.Sel("setProxyAutoConfigurationURL:"), value)
}
// A string containing the Proxy Auto Configuration (PAC) JavaScript source
// code.
//
// # Discussion
// 
// If [AutoProxyConfigurationEnabled] is set to [true] then the system will
// execute the PAC script to determine what proxies to use (if any) for HTTP
// and HTTPS connections.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/proxyAutoConfigurationJavaScript
func (p NEProxySettings) ProxyAutoConfigurationJavaScript() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("proxyAutoConfigurationJavaScript"))
	return foundation.NSStringFromID(rv).String()
}
func (p NEProxySettings) SetProxyAutoConfigurationJavaScript(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setProxyAutoConfigurationJavaScript:"), objc.String(value))
}
// A Boolean indicating if a static HTTP proxy will be used.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/httpEnabled
func (p NEProxySettings) HTTPEnabled() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("HTTPEnabled"))
	return rv
}
func (p NEProxySettings) SetHTTPEnabled(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setHTTPEnabled:"), value)
}
// An [NEProxyServer] object containing the static HTTP proxy server settings.
//
// # Discussion
// 
// If [AutoProxyConfigurationEnabled] is [false] and [HTTPEnabled] is [true],
// then the proxy server specified in this property will be used for HTTP
// connections.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/httpServer
func (p NEProxySettings) HTTPServer() INEProxyServer {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("HTTPServer"))
	return NEProxyServerFromID(objc.ID(rv))
}
func (p NEProxySettings) SetHTTPServer(value INEProxyServer) {
	objc.Send[struct{}](p.ID, objc.Sel("setHTTPServer:"), value)
}
// A Boolean indicating if a static HTTPS proxy will be used.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/httpsEnabled
func (p NEProxySettings) HTTPSEnabled() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("HTTPSEnabled"))
	return rv
}
func (p NEProxySettings) SetHTTPSEnabled(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setHTTPSEnabled:"), value)
}
// An [NEProxyServer] object containing the static HTTPS proxy server
// settings.
//
// # Discussion
// 
// If [AutoProxyConfigurationEnabled] is [false] and [HTTPSEnabled] is [true],
// then the proxy server specified in this property will be used for HTTPS
// connections.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/httpsServer
func (p NEProxySettings) HTTPSServer() INEProxyServer {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("HTTPSServer"))
	return NEProxyServerFromID(objc.ID(rv))
}
func (p NEProxySettings) SetHTTPSServer(value INEProxyServer) {
	objc.Send[struct{}](p.ID, objc.Sel("setHTTPSServer:"), value)
}
// A Boolean indicating if HTTP requests using single-label host names should
// be excluded from using the proxy settings.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/excludeSimpleHostnames
func (p NEProxySettings) ExcludeSimpleHostnames() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("excludeSimpleHostnames"))
	return rv
}
func (p NEProxySettings) SetExcludeSimpleHostnames(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setExcludeSimpleHostnames:"), value)
}
// An array of domain name patterns. If the destination host name of an HTTP
// connection matches one of these patterns then the proxy settings will not
// be used for the connection.
//
// # Discussion
// 
// The pattern strings may contain ‘*’ characters as wildcards.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/exceptionList
func (p NEProxySettings) ExceptionList() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("exceptionList"))
	return objc.ConvertSliceToStrings(rv)
}
func (p NEProxySettings) SetExceptionList(value []string) {
	objc.Send[struct{}](p.ID, objc.Sel("setExceptionList:"), objectivec.StringSliceToNSArray(value))
}
// An array of domain strings.
//
// # Discussion
// 
// If the destination host name of a HTTP connection shares a suffix with one
// of these strings then the proxy settings will be used for the HTTP
// connection. Otherwise the proxy settings will not be used.
// 
// This property should be used in conjunction with a split tunnel VPN, where
// only certain networks are tunneled by the VPN. The domains of those split
// tunneling networks should be specified in this property.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/matchDomains
func (p NEProxySettings) MatchDomains() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("matchDomains"))
	return objc.ConvertSliceToStrings(rv)
}
func (p NEProxySettings) SetMatchDomains(value []string) {
	objc.Send[struct{}](p.ID, objc.Sel("setMatchDomains:"), objectivec.StringSliceToNSArray(value))
}
// The tunnel DNS settings.
//
// See: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/dnssettings
func (p NEProxySettings) DnsSettings() INEDNSSettings {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("DNSSettings"))
	return NEDNSSettingsFromID(objc.ID(rv))
}
func (p NEProxySettings) SetDnsSettings(value INEDNSSettings) {
	objc.Send[struct{}](p.ID, objc.Sel("setDNSSettings:"), value)
}
// The tunnel HTTP proxy settings.
//
// See: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/proxysettings
func (p NEProxySettings) ProxySettings() INEProxySettings {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("proxySettings"))
	return NEProxySettingsFromID(objc.ID(rv))
}
func (p NEProxySettings) SetProxySettings(value INEProxySettings) {
	objc.Send[struct{}](p.ID, objc.Sel("setProxySettings:"), value)
}
// The IP address of the tunnel server.
//
// See: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/tunnelremoteaddress
func (p NEProxySettings) TunnelRemoteAddress() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("tunnelRemoteAddress"))
	return foundation.NSStringFromID(rv).String()
}
func (p NEProxySettings) SetTunnelRemoteAddress(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setTunnelRemoteAddress:"), objc.String(value))
}

