// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKWebpagePreferences] class.
var (
	_WKWebpagePreferencesClass     WKWebpagePreferencesClass
	_WKWebpagePreferencesClassOnce sync.Once
)

func getWKWebpagePreferencesClass() WKWebpagePreferencesClass {
	_WKWebpagePreferencesClassOnce.Do(func() {
		_WKWebpagePreferencesClass = WKWebpagePreferencesClass{class: objc.GetClass("WKWebpagePreferences")}
	})
	return _WKWebpagePreferencesClass
}

// GetWKWebpagePreferencesClass returns the class object for WKWebpagePreferences.
func GetWKWebpagePreferencesClass() WKWebpagePreferencesClass {
	return getWKWebpagePreferencesClass()
}

type WKWebpagePreferencesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKWebpagePreferencesClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKWebpagePreferencesClass) Alloc() WKWebpagePreferences {
	rv := objc.Send[WKWebpagePreferences](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that specifies the behaviors to use when loading and rendering
// page content.
//
// # Overview
//
// Create a [WKWebpagePreferences] object when you want to change the default
// rendering behavior of your web view. Typically, iOS devices render web
// content for a mobile experience, and Mac devices render content for a
// desktop experience.
//
// # Setting the JavaScript preferences
//
//   - [WKWebpagePreferences.AllowsContentJavaScript]: A Boolean value that indicates whether JavaScript from web content is allowed to run.
//   - [WKWebpagePreferences.SetAllowsContentJavaScript]
//
// # Setting the preferred content mode
//
//   - [WKWebpagePreferences.PreferredContentMode]: The content mode for the web view to use when it loads and renders a webpage.
//   - [WKWebpagePreferences.SetPreferredContentMode]
//
// # Getting Lockdown Mode info
//
//   - [WKWebpagePreferences.LockdownModeEnabled]: A Boolean value that indicates whether to use Lockdown Mode in the web view.
//   - [WKWebpagePreferences.SetLockdownModeEnabled]
//
// # Instance Properties
//
//   - [WKWebpagePreferences.PreferredHTTPSNavigationPolicy]
//   - [WKWebpagePreferences.SetPreferredHTTPSNavigationPolicy]
//
// See: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences
type WKWebpagePreferences struct {
	objectivec.Object
}

// WKWebpagePreferencesFromID constructs a [WKWebpagePreferences] from an objc.ID.
//
// An object that specifies the behaviors to use when loading and rendering
// page content.
func WKWebpagePreferencesFromID(id objc.ID) WKWebpagePreferences {
	return WKWebpagePreferences{objectivec.Object{ID: id}}
}

// NOTE: WKWebpagePreferences adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKWebpagePreferences] class.
//
// # Setting the JavaScript preferences
//
//   - [IWKWebpagePreferences.AllowsContentJavaScript]: A Boolean value that indicates whether JavaScript from web content is allowed to run.
//   - [IWKWebpagePreferences.SetAllowsContentJavaScript]
//
// # Setting the preferred content mode
//
//   - [IWKWebpagePreferences.PreferredContentMode]: The content mode for the web view to use when it loads and renders a webpage.
//   - [IWKWebpagePreferences.SetPreferredContentMode]
//
// # Getting Lockdown Mode info
//
//   - [IWKWebpagePreferences.LockdownModeEnabled]: A Boolean value that indicates whether to use Lockdown Mode in the web view.
//   - [IWKWebpagePreferences.SetLockdownModeEnabled]
//
// # Instance Properties
//
//   - [IWKWebpagePreferences.PreferredHTTPSNavigationPolicy]
//   - [IWKWebpagePreferences.SetPreferredHTTPSNavigationPolicy]
//
// See: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences
type IWKWebpagePreferences interface {
	objectivec.IObject

	// Topic: Setting the JavaScript preferences

	// A Boolean value that indicates whether JavaScript from web content is allowed to run.
	AllowsContentJavaScript() bool
	SetAllowsContentJavaScript(value bool)

	// Topic: Setting the preferred content mode

	// The content mode for the web view to use when it loads and renders a webpage.
	PreferredContentMode() WKContentMode
	SetPreferredContentMode(value WKContentMode)

	// Topic: Getting Lockdown Mode info

	// A Boolean value that indicates whether to use Lockdown Mode in the web view.
	LockdownModeEnabled() bool
	SetLockdownModeEnabled(value bool)

	// Topic: Instance Properties

	PreferredHTTPSNavigationPolicy() WKWebpagePreferencesUpgradeToHTTPSPolicy
	SetPreferredHTTPSNavigationPolicy(value WKWebpagePreferencesUpgradeToHTTPSPolicy)

	SecurityRestrictionMode() WKSecurityRestrictionMode
	SetSecurityRestrictionMode(value WKSecurityRestrictionMode)
}

// Init initializes the instance.
func (w WKWebpagePreferences) Init() WKWebpagePreferences {
	rv := objc.Send[WKWebpagePreferences](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WKWebpagePreferences) Autorelease() WKWebpagePreferences {
	rv := objc.Send[WKWebpagePreferences](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKWebpagePreferences creates a new WKWebpagePreferences instance.
func NewWKWebpagePreferences() WKWebpagePreferences {
	class := getWKWebpagePreferencesClass()
	rv := objc.Send[WKWebpagePreferences](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether JavaScript from web content is
// allowed to run.
//
// # Discussion
//
// The default value of this property is true. If you change the value to
// false, the web view doesn’t execute JavaScript code referenced by the web
// content. That includes JavaScript code found in inline “ elements, “
// URLs, and all other referenced JavaScript content.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/allowsContentJavaScript
func (w WKWebpagePreferences) AllowsContentJavaScript() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("allowsContentJavaScript"))
	return rv
}
func (w WKWebpagePreferences) SetAllowsContentJavaScript(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setAllowsContentJavaScript:"), value)
}

// The content mode for the web view to use when it loads and renders a
// webpage.
//
// # Discussion
//
// The default value of this property is [WKContentModeRecommended]. The web
// view ignores this preference for subframe navigation.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/preferredContentMode
func (w WKWebpagePreferences) PreferredContentMode() WKContentMode {
	rv := objc.Send[WKContentMode](w.ID, objc.Sel("preferredContentMode"))
	return WKContentMode(rv)
}
func (w WKWebpagePreferences) SetPreferredContentMode(value WKContentMode) {
	objc.Send[struct{}](w.ID, objc.Sel("setPreferredContentMode:"), value)
}

// A Boolean value that indicates whether to use Lockdown Mode in the web
// view.
//
// # Discussion
//
// By default, this reflects whether the user has enabled Lockdown Mode on the
// device. Update this preference to override the device setting when you
// implement a per-website or similar setting.
//
// For more information about Lockdown Mode, see [About Lockdown Mode].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/isLockdownModeEnabled
//
// [About Lockdown Mode]: https://support.apple.com/en-us/HT212650
func (w WKWebpagePreferences) LockdownModeEnabled() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isLockdownModeEnabled"))
	return rv
}
func (w WKWebpagePreferences) SetLockdownModeEnabled(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setLockdownModeEnabled:"), value)
}

// # Discussion
//
// A WKWebpagePreferencesUpgradeToHTTPSPolicy indicating the desired mode used
// when performing a top-level navigation to a webpage.
//
// The default value is
// WKWebpagePreferencesUpgradeToHTTPSPolicyKeepAsRequested. The stated
// preference is ignored on subframe navigation, and it may be ignored based
// on system configuration. The upgradeKnownHostsToHTTPS property on
// WKWebViewConfiguration supercedes this policy for known hosts.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/preferredHTTPSNavigationPolicy
func (w WKWebpagePreferences) PreferredHTTPSNavigationPolicy() WKWebpagePreferencesUpgradeToHTTPSPolicy {
	rv := objc.Send[WKWebpagePreferencesUpgradeToHTTPSPolicy](w.ID, objc.Sel("preferredHTTPSNavigationPolicy"))
	return WKWebpagePreferencesUpgradeToHTTPSPolicy(rv)
}
func (w WKWebpagePreferences) SetPreferredHTTPSNavigationPolicy(value WKWebpagePreferencesUpgradeToHTTPSPolicy) {
	objc.Send[struct{}](w.ID, objc.Sel("setPreferredHTTPSNavigationPolicy:"), value)
}

// # Discussion
//
// Security restriction mode for this navigation.
//
// Security restriction modes provide different levels of security hardening
// for high-risk browsing contexts.
// WKSecurityRestrictionModeMaximizeCompatibility provides additional
// hardening while maintaining full web compatibility:
//
// - JavaScript JIT compilation disabled (interpreter-only execution) -
// Increased Memory Tagging Extension (MTE) coverage across allocations in the
// WebContent process Setting a security restriction mode creates separate,
// isolated WebContent processes for the specified protection level. This
// preference only applies to main frame navigations and will be ignored for
// subframe navigations. When set for a main frame, all subframe content and
// opened windows inherit the same security restrictions. When the system has
// chosen WKSecurityRestrictionModeLockdown (e.g., in Lockdown Mode), attempts
// to set a less restrictive mode will fail silently. The default value is
// WKSecurityRestrictionModeNone.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/securityRestrictionMode
func (w WKWebpagePreferences) SecurityRestrictionMode() WKSecurityRestrictionMode {
	rv := objc.Send[WKSecurityRestrictionMode](w.ID, objc.Sel("securityRestrictionMode"))
	return WKSecurityRestrictionMode(rv)
}
func (w WKWebpagePreferences) SetSecurityRestrictionMode(value WKSecurityRestrictionMode) {
	objc.Send[struct{}](w.ID, objc.Sel("setSecurityRestrictionMode:"), value)
}
