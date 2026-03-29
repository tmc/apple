// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NETunnelProviderManager] class.
var (
	_NETunnelProviderManagerClass     NETunnelProviderManagerClass
	_NETunnelProviderManagerClassOnce sync.Once
)

func getNETunnelProviderManagerClass() NETunnelProviderManagerClass {
	_NETunnelProviderManagerClassOnce.Do(func() {
		_NETunnelProviderManagerClass = NETunnelProviderManagerClass{class: objc.GetClass("NETunnelProviderManager")}
	})
	return _NETunnelProviderManagerClass
}

// GetNETunnelProviderManagerClass returns the class object for NETunnelProviderManager.
func GetNETunnelProviderManagerClass() NETunnelProviderManagerClass {
	return getNETunnelProviderManagerClass()
}

type NETunnelProviderManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NETunnelProviderManagerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NETunnelProviderManagerClass) Alloc() NETunnelProviderManager {
	rv := objc.Send[NETunnelProviderManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object to create and manage the tunnel provider’s VPN configuration.
//
// # Overview
// 
// Like its superclass [NEVPNManager], you use the [NETunnelProviderManager]
// class to configure and control VPN connections. The difference is that
// [NETunnelProviderManager] is used to to configure and control VPN
// connections that use a custom VPN protocol. The client side of the custom
// protocol implementation is implemented as a Packet Tunnel Provider
// extension. The Packet Tunnel Provider extension’s containing app uses
// [NETunnelProviderManager] to create and manage VPN configurations that use
// the custom protocol, and to control the VPN connections specified by the
// configurations.
// 
// The [NETunnelProviderManager] class inherits most of its functionality from
// the [NEVPNManager] class. The key differences to be aware of when using
// [NETunnelProviderManager] are:
// 
// - The [NETunnelProviderManager.ProtocolConfiguration] property can only be set to instances of the
// [NETunnelProviderProtocol] class - The [NETunnelProviderManager.Connection] read-only property is
// set to an instance of the [NETunnelProviderSession] class.
// 
// # Configuration Model
// 
// Each [NETunnelProviderManager] instance corresponds to a single VPN
// configuration stored in the Network Extension preferences. Multiple VPN
// configurations can be created and managed by creating multiple
// [NETunnelProviderManager] instances.
// 
// Each VPN configuration is associated with the app that created it. The
// app’s view of the Network Extension preferences is limited to include
// only the configurations that were created by the app.
// 
// VPN configurations created using [NETunnelProviderManager] are classified
// as regular enterprise VPN configurations (as opposed to the Personal VPN
// configurations created by [NEVPNManager]). Only one enterprise VPN
// configuration can be enabled on the system at a time. If both a Personal
// VPN and an enterprise VPN are active on the system simultaneously, the
// enterprise VPN takes precedence, meaning that if the routes for the two
// VPNs conflict then the routes for the enterprise VPN will take precedence.
// The Personal VPN will remain active and connected while the enterprise VPN
// is active and connected, and any traffic that is routed to the Personal VPN
// and is not routed to the enterprise VPN will continue to traverse the
// Personal VPN.
// 
// # Profile Configuration
// 
// It is possible to create Packet Tunnel Provider configurations using
// configuration profiles. See the
// `com.AppleXCUIElementTypeVpnXCUIElementTypeManaged()` and
// `com.AppleXCUIElementTypeVpnXCUIElementTypeManagedXCUIElementTypeApplayer()`
// payload types in [Configuration Profile Reference]. To specify that a
// configuration created via a profile payload is associated with a particular
// app (and therefore allow the app to use [NETunnelProviderManager] to manage
// the configuration), the app’s bundle identifier must be set as the value
// of the [VPNSubType] field in the profile payload.
// 
// Credential Storage
// 
// VPN credentials such as private keys and passwords that are imported into
// the system via configuration profiles are stored in the keychain in a
// special access group called
// `com.AppleXCUIElementTypeManagedXCUIElementTypeVpnXCUIElementTypeShared()`.
// In order to use these credentials the app and Packet Tunnel Provider
// extension must have the
// `com.AppleXCUIElementTypeManagedXCUIElementTypeVpnXCUIElementTypeShared()`
// keychain access group entitlement.
// 
// # Routing Network Data to the VPN
// 
// There are two ways or methods by which network data is routed to the VPN:
// 
// - By destination IP address
// - By source application (Per-App VPN)
// 
// Routing by Destination IP
// 
// This is the default routing method. The IP routes are specified by the
// Packet Tunnel Provider extension at the time that the VPN tunnel is fully
// established. See [NETunnelProvider] for more details.
// 
// Per-App VPN
// 
// The only way to configure Per-App VPN is by enrolling the device in a
// Mobile Device Management (MDM) system, and then linking apps that are
// managed by the MDM system with a VPN configuration created from a
// `com.AppleXCUIElementTypeVpnXCUIElementTypeManagedXCUIElementTypeApplayer()`
// configuration profile payload. Here are some details about how this works:
// 
// - The MDM server creates a configuration profile containing a
// `com.AppleXCUIElementTypeVpnXCUIElementTypeManagedXCUIElementTypeApplayer()`
// payload. The
// `com.AppleXCUIElementTypeVpnXCUIElementTypeManagedXCUIElementTypeApplayer()`
// payload contains all of the usual VPN configuration profile payload fields,
// and also must contain a [VPNUUID] field, containing a unique string defined
// by the MDM server. - If the VPN provider extension is a Packet Tunnel
// Provider extension, then the [ProviderType] field in the
// `com.AppleXCUIElementTypeVpnXCUIElementTypeManagedXCUIElementTypeApplayer()`
// payload should be set to `packet-tunnel`. If the VPN provider extension is
// an App Proxy Provider extension, then the [ProviderType] field in the
// `com.AppleXCUIElementTypeVpnXCUIElementTypeManagedXCUIElementTypeApplayer()`
// should be set to `app-proxy`. - The MDM server adds a [VPNUUID] key to the
// attributes dictionary of all of the managed apps that will use the VPN. The
// value of the [VPNUUID] key must be set to the same unique string contained
// in the [VPNUUID] field in the
// `com.AppleXCUIElementTypeVpnXCUIElementTypeManagedXCUIElementTypeApplayer()`
// payload. - The MDM server pushes the configuration profile and the managed
// apps to the iOS device using the MDM protocol.
// 
// The MDM client running on the device creates one app rule in the VPN
// configuration for each managed app that is linked to the VPN configuration
// via the [VPNUUID] app attribute.
// 
// Per-App VPN On Demand
// 
// The Per-App VPN app rules serve as both routing rules and VPN On Demand
// rules. This is in contrast to IP destination-based routing, where the VPN
// On Demand rules are configured separately from the routing rules. When the
// `onDemandEnabled` property is set to [true] and an app that matches the
// Per-App VPN rules attempts to communicate over the network, the VPN will be
// started automatically.
// 
// It is possible to set regular VPN On Demand rules in a Per-App VPN
// configuration via the [NETunnelProviderManager.OnDemandRules] property, but only
// [NEOnDemandRuleDisconnect] rules will be used. When a
// [NEOnDemandRuleDisconnect] rule matches, apps which match the Per-App VPN
// rules will bypass the VPN.
// 
// Testing Per-App VPN
// 
// As described above, an MDM server is required to configure Per-App VPN for
// VPN apps distributed via the App Store. To make testing Per-App VPN easier,
// it is possible to configure Per-App VPN without an MDM server during
// development by using the [NETestAppMapping] `Info.Plist()` key.
// 
// Here is what you need to do to make use of this capability:
// 
// - Create a configuration profile containing a
// `com.AppleXCUIElementTypeVpnXCUIElementTypeManagedXCUIElementTypeApplayer()`
// payload as described in [Configuration Profile Reference]. In addition to
// all of the usual VPN configuration payload fields, the payload must also
// contain a [VPNUUID] field, containing a unique string defined by you. - Add
// the [NETestAppMapping] key to your app’s `Info.Plist()`. The value of
// this key should be a dictionary that maps [VPNUUID] values to arrays of app
// bundle identifiers. Here is a sample:
// 
// - Rebuild the app. - Install the app and the configuration profile on the
// device.
// 
// The system will create one app rule in the VPN configuration for each
// bundle identifier listed in the array in the [NETestAppMapping] dictionary
// corresponding to the value of the [VPNUUID] field in the
// `com.AppleXCUIElementTypeVpnXCUIElementTypeManagedXCUIElementTypeApplayer()`
// payload.
//
// [Configuration Profile Reference]: https://developer.apple.com/library/archive/featuredarticles/iPhoneConfigurationProfileRef/Introduction/Introduction.html#//apple_ref/doc/uid/TP40010206
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Managing tunnel configurations
//
//   - [NETunnelProviderManager.CopyAppRules]: Returns a copy of the app rules currently set in the configuration.
//
// # Getting tunnel configuration properties
//
//   - [NETunnelProviderManager.RoutingMethod]: The method that the system uses to route network traffic to the tunnel.
//
// # Configuring a per-app VPN
//
//   - [NETunnelProviderManager.AppRules]: The rules for specific apps in a per-app VPN.
//   - [NETunnelProviderManager.SetAppRules]
//   - [NETunnelProviderManager.ExcludedDomains]: The domains that the system excludes from a per-app VPN.
//   - [NETunnelProviderManager.SetExcludedDomains]
//   - [NETunnelProviderManager.AssociatedDomains]: The domains that the system routes network traffic through for a per-app VPN.
//   - [NETunnelProviderManager.SetAssociatedDomains]
//   - [NETunnelProviderManager.CalendarDomains]: The calendar servers that the system routes connections from the Calendar app through for a per-app VPN.
//   - [NETunnelProviderManager.SetCalendarDomains]
//   - [NETunnelProviderManager.ContactsDomains]: The contacts servers that the system routes connections from the Contacts app through for a per-app VPN.
//   - [NETunnelProviderManager.SetContactsDomains]
//   - [NETunnelProviderManager.MailDomains]: The mail servers that the system routes connections from the Mail app through for a per-app VPN.
//   - [NETunnelProviderManager.SetMailDomains]
//   - [NETunnelProviderManager.SafariDomains]: The website domains that the system routes connections from the Safari app through a per-app VPN.
//   - [NETunnelProviderManager.SetSafariDomains]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager
type NETunnelProviderManager struct {
	NEVPNManager
}

// NETunnelProviderManagerFromID constructs a [NETunnelProviderManager] from an objc.ID.
//
// An object to create and manage the tunnel provider’s VPN configuration.
func NETunnelProviderManagerFromID(id objc.ID) NETunnelProviderManager {
	return NETunnelProviderManager{NEVPNManager: NEVPNManagerFromID(id)}
}
// NOTE: NETunnelProviderManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NETunnelProviderManager] class.
//
// # Managing tunnel configurations
//
//   - [INETunnelProviderManager.CopyAppRules]: Returns a copy of the app rules currently set in the configuration.
//
// # Getting tunnel configuration properties
//
//   - [INETunnelProviderManager.RoutingMethod]: The method that the system uses to route network traffic to the tunnel.
//
// # Configuring a per-app VPN
//
//   - [INETunnelProviderManager.AppRules]: The rules for specific apps in a per-app VPN.
//   - [INETunnelProviderManager.SetAppRules]
//   - [INETunnelProviderManager.ExcludedDomains]: The domains that the system excludes from a per-app VPN.
//   - [INETunnelProviderManager.SetExcludedDomains]
//   - [INETunnelProviderManager.AssociatedDomains]: The domains that the system routes network traffic through for a per-app VPN.
//   - [INETunnelProviderManager.SetAssociatedDomains]
//   - [INETunnelProviderManager.CalendarDomains]: The calendar servers that the system routes connections from the Calendar app through for a per-app VPN.
//   - [INETunnelProviderManager.SetCalendarDomains]
//   - [INETunnelProviderManager.ContactsDomains]: The contacts servers that the system routes connections from the Contacts app through for a per-app VPN.
//   - [INETunnelProviderManager.SetContactsDomains]
//   - [INETunnelProviderManager.MailDomains]: The mail servers that the system routes connections from the Mail app through for a per-app VPN.
//   - [INETunnelProviderManager.SetMailDomains]
//   - [INETunnelProviderManager.SafariDomains]: The website domains that the system routes connections from the Safari app through a per-app VPN.
//   - [INETunnelProviderManager.SetSafariDomains]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager
type INETunnelProviderManager interface {
	INEVPNManager

	// Topic: Managing tunnel configurations

	// Returns a copy of the app rules currently set in the configuration.
	CopyAppRules() []NEAppRule

	// Topic: Getting tunnel configuration properties

	// The method that the system uses to route network traffic to the tunnel.
	RoutingMethod() NETunnelProviderRoutingMethod

	// Topic: Configuring a per-app VPN

	// The rules for specific apps in a per-app VPN.
	AppRules() []NEAppRule
	SetAppRules(value []NEAppRule)
	// The domains that the system excludes from a per-app VPN.
	ExcludedDomains() []string
	SetExcludedDomains(value []string)
	// The domains that the system routes network traffic through for a per-app VPN.
	AssociatedDomains() []string
	SetAssociatedDomains(value []string)
	// The calendar servers that the system routes connections from the Calendar app through for a per-app VPN.
	CalendarDomains() []string
	SetCalendarDomains(value []string)
	// The contacts servers that the system routes connections from the Contacts app through for a per-app VPN.
	ContactsDomains() []string
	SetContactsDomains(value []string)
	// The mail servers that the system routes connections from the Mail app through for a per-app VPN.
	MailDomains() []string
	SetMailDomains(value []string)
	// The website domains that the system routes connections from the Safari app through a per-app VPN.
	SafariDomains() []string
	SetSafariDomains(value []string)
}

// Init initializes the instance.
func (t NETunnelProviderManager) Init() NETunnelProviderManager {
	rv := objc.Send[NETunnelProviderManager](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NETunnelProviderManager) Autorelease() NETunnelProviderManager {
	rv := objc.Send[NETunnelProviderManager](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETunnelProviderManager creates a new NETunnelProviderManager instance.
func NewNETunnelProviderManager() NETunnelProviderManager {
	class := getNETunnelProviderManagerClass()
	rv := objc.Send[NETunnelProviderManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a copy of the app rules currently set in the configuration.
//
// # Return Value
// 
// An array of [NEAppRule] objects, or `nil` if the configuration doesn’t
// have any app rules.
//
// # Discussion
// 
// This method provides read-only access to the configuration’s app rules.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/copyAppRules()
func (t NETunnelProviderManager) CopyAppRules() []NEAppRule {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("copyAppRules"))
	return objc.ConvertSlice(rv, func(id objc.ID) NEAppRule {
		return NEAppRuleFromID(id)
	})
}

// Returns a tunnel provider manager for managing a per-app VPN configuration.
//
// # Return Value
// 
// An object you use to configure a per-app VPN.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/forPerAppVPN()
func (_NETunnelProviderManagerClass NETunnelProviderManagerClass) ForPerAppVPN() NETunnelProviderManager {
	rv := objc.Send[objc.ID](objc.ID(_NETunnelProviderManagerClass.class), objc.Sel("forPerAppVPN"))
	return NETunnelProviderManagerFromID(rv)
}

// The method that the system uses to route network traffic to the tunnel.
//
// # Discussion
// 
// The default is [NETunnelProviderRoutingMethod.destinationIP].
//
// [NETunnelProviderRoutingMethod.destinationIP]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderRoutingMethod/destinationIP
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/routingMethod
func (t NETunnelProviderManager) RoutingMethod() NETunnelProviderRoutingMethod {
	rv := objc.Send[NETunnelProviderRoutingMethod](t.ID, objc.Sel("routingMethod"))
	return NETunnelProviderRoutingMethod(rv)
}
// The rules for specific apps in a per-app VPN.
//
// # Discussion
// 
// For per-app VPNs only, the system routes network traffic originating from
// an app that matches one of these rules through the VPN.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/appRules
func (t NETunnelProviderManager) AppRules() []NEAppRule {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("appRules"))
	return objc.ConvertSlice(rv, func(id objc.ID) NEAppRule {
		return NEAppRuleFromID(id)
	})
}
func (t NETunnelProviderManager) SetAppRules(value []NEAppRule) {
	objc.Send[struct{}](t.ID, objc.Sel("setAppRules:"), objectivec.IObjectSliceToNSArray(value))
}
// The domains that the system excludes from a per-app VPN.
//
// # Discussion
// 
// For per-app VPNs only, the system doesn’t route network traffic to
// servers within these domains.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/excludedDomains
func (t NETunnelProviderManager) ExcludedDomains() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("excludedDomains"))
	return objc.ConvertSliceToStrings(rv)
}
func (t NETunnelProviderManager) SetExcludedDomains(value []string) {
	objc.Send[struct{}](t.ID, objc.Sel("setExcludedDomains:"), objectivec.StringSliceToNSArray(value))
}
// The domains that the system routes network traffic through for a per-app
// VPN.
//
// # Discussion
// 
// For per-app VPNs only, the system routes HTTP requests to download the
// Apple app site association files for domains in this property through the
// VPN.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/associatedDomains
func (t NETunnelProviderManager) AssociatedDomains() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("associatedDomains"))
	return objc.ConvertSliceToStrings(rv)
}
func (t NETunnelProviderManager) SetAssociatedDomains(value []string) {
	objc.Send[struct{}](t.ID, objc.Sel("setAssociatedDomains:"), objectivec.StringSliceToNSArray(value))
}
// The calendar servers that the system routes connections from the Calendar
// app through for a per-app VPN.
//
// # Discussion
// 
// This property applies only to per-app VPNs.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/calendarDomains
func (t NETunnelProviderManager) CalendarDomains() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("calendarDomains"))
	return objc.ConvertSliceToStrings(rv)
}
func (t NETunnelProviderManager) SetCalendarDomains(value []string) {
	objc.Send[struct{}](t.ID, objc.Sel("setCalendarDomains:"), objectivec.StringSliceToNSArray(value))
}
// The contacts servers that the system routes connections from the Contacts
// app through for a per-app VPN.
//
// # Discussion
// 
// This property applies only to per-app VPNs.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/contactsDomains
func (t NETunnelProviderManager) ContactsDomains() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("contactsDomains"))
	return objc.ConvertSliceToStrings(rv)
}
func (t NETunnelProviderManager) SetContactsDomains(value []string) {
	objc.Send[struct{}](t.ID, objc.Sel("setContactsDomains:"), objectivec.StringSliceToNSArray(value))
}
// The mail servers that the system routes connections from the Mail app
// through for a per-app VPN.
//
// # Discussion
// 
// This property applies only to per-app VPNs.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/mailDomains
func (t NETunnelProviderManager) MailDomains() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("mailDomains"))
	return objc.ConvertSliceToStrings(rv)
}
func (t NETunnelProviderManager) SetMailDomains(value []string) {
	objc.Send[struct{}](t.ID, objc.Sel("setMailDomains:"), objectivec.StringSliceToNSArray(value))
}
// The website domains that the system routes connections from the Safari app
// through a per-app VPN.
//
// # Discussion
// 
// For per-app VPNs only, when the user navigates in Safari to a website
// within one of these domains, the system routes the website traffic through
// the VPN.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/safariDomains
func (t NETunnelProviderManager) SafariDomains() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("safariDomains"))
	return objc.ConvertSliceToStrings(rv)
}
func (t NETunnelProviderManager) SetSafariDomains(value []string) {
	objc.Send[struct{}](t.ID, objc.Sel("setSafariDomains:"), objectivec.StringSliceToNSArray(value))
}

