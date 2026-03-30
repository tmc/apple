// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKPreferences] class.
var (
	_WKPreferencesClass     WKPreferencesClass
	_WKPreferencesClassOnce sync.Once
)

func getWKPreferencesClass() WKPreferencesClass {
	_WKPreferencesClassOnce.Do(func() {
		_WKPreferencesClass = WKPreferencesClass{class: objc.GetClass("WKPreferences")}
	})
	return _WKPreferencesClass
}

// GetWKPreferencesClass returns the class object for WKPreferences.
func GetWKPreferencesClass() WKPreferencesClass {
	return getWKPreferencesClass()
}

type WKPreferencesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKPreferencesClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKPreferencesClass) Alloc() WKPreferences {
	rv := objc.Send[WKPreferences](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that encapsulates the standard behaviors to apply to websites.
//
// # Overview
//
// Use a [WKPreferences] object to specify the preferences for your website,
// including the minimum font size, the JavaScript behavior, and the behavior
// for handling fraudulent websites. Create this object and assign it to the
// [Preferences] property of the [WKWebViewConfiguration] object you use to
// create your web view.
//
// # Setting Rendering Preferences
//
//   - [WKPreferences.MinimumFontSize]: The minimum font size, in points.
//   - [WKPreferences.SetMinimumFontSize]
//   - [WKPreferences.ShouldPrintBackgrounds]: A Boolean value that indicates whether to include any background color or graphics when printing content.
//   - [WKPreferences.SetShouldPrintBackgrounds]
//
// # Setting Behavior Preferences
//
//   - [WKPreferences.TabFocusesLinks]: A Boolean value that indicates whether pressing the tab key changes the focus to links and form controls.
//   - [WKPreferences.SetTabFocusesLinks]
//   - [WKPreferences.TextInteractionEnabled]: A Boolean value that indicates whether to allow people to select or otherwise interact with text.
//   - [WKPreferences.SetTextInteractionEnabled]
//   - [WKPreferences.ElementFullscreenEnabled]: A Boolean value that indicates whether a web view can display content full screen.
//   - [WKPreferences.SetElementFullscreenEnabled]
//   - [WKPreferences.InactiveSchedulingPolicy]: A policy you set to specify how a web view that’s not in a window handles tasks.
//   - [WKPreferences.SetInactiveSchedulingPolicy]
//
// # Setting Java and JavaScript Preferences
//
//   - [WKPreferences.JavaScriptCanOpenWindowsAutomatically]: A Boolean value that indicates whether JavaScript can open windows without user interaction.
//   - [WKPreferences.SetJavaScriptCanOpenWindowsAutomatically]
//   - [WKPreferences.SiteSpecificQuirksModeEnabled]: A Boolean that indicates whether to apply site-specific compatibility workarounds.
//   - [WKPreferences.SetSiteSpecificQuirksModeEnabled]
//
// # Setting Fraud Warning Preferences
//
//   - [WKPreferences.FraudulentWebsiteWarningEnabled]: A Boolean value that indicates whether the web view shows warnings for suspected fraudulent content, such as malware or phishing attemps.
//   - [WKPreferences.SetFraudulentWebsiteWarningEnabled]
//
// See: https://developer.apple.com/documentation/WebKit/WKPreferences
type WKPreferences struct {
	objectivec.Object
}

// WKPreferencesFromID constructs a [WKPreferences] from an objc.ID.
//
// An object that encapsulates the standard behaviors to apply to websites.
func WKPreferencesFromID(id objc.ID) WKPreferences {
	return WKPreferences{objectivec.Object{ID: id}}
}

// NOTE: WKPreferences adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKPreferences] class.
//
// # Setting Rendering Preferences
//
//   - [IWKPreferences.MinimumFontSize]: The minimum font size, in points.
//   - [IWKPreferences.SetMinimumFontSize]
//   - [IWKPreferences.ShouldPrintBackgrounds]: A Boolean value that indicates whether to include any background color or graphics when printing content.
//   - [IWKPreferences.SetShouldPrintBackgrounds]
//
// # Setting Behavior Preferences
//
//   - [IWKPreferences.TabFocusesLinks]: A Boolean value that indicates whether pressing the tab key changes the focus to links and form controls.
//   - [IWKPreferences.SetTabFocusesLinks]
//   - [IWKPreferences.TextInteractionEnabled]: A Boolean value that indicates whether to allow people to select or otherwise interact with text.
//   - [IWKPreferences.SetTextInteractionEnabled]
//   - [IWKPreferences.ElementFullscreenEnabled]: A Boolean value that indicates whether a web view can display content full screen.
//   - [IWKPreferences.SetElementFullscreenEnabled]
//   - [IWKPreferences.InactiveSchedulingPolicy]: A policy you set to specify how a web view that’s not in a window handles tasks.
//   - [IWKPreferences.SetInactiveSchedulingPolicy]
//
// # Setting Java and JavaScript Preferences
//
//   - [IWKPreferences.JavaScriptCanOpenWindowsAutomatically]: A Boolean value that indicates whether JavaScript can open windows without user interaction.
//   - [IWKPreferences.SetJavaScriptCanOpenWindowsAutomatically]
//   - [IWKPreferences.SiteSpecificQuirksModeEnabled]: A Boolean that indicates whether to apply site-specific compatibility workarounds.
//   - [IWKPreferences.SetSiteSpecificQuirksModeEnabled]
//
// # Setting Fraud Warning Preferences
//
//   - [IWKPreferences.FraudulentWebsiteWarningEnabled]: A Boolean value that indicates whether the web view shows warnings for suspected fraudulent content, such as malware or phishing attemps.
//   - [IWKPreferences.SetFraudulentWebsiteWarningEnabled]
//
// See: https://developer.apple.com/documentation/WebKit/WKPreferences
type IWKPreferences interface {
	objectivec.IObject

	// Topic: Setting Rendering Preferences

	// The minimum font size, in points.
	MinimumFontSize() float64
	SetMinimumFontSize(value float64)
	// A Boolean value that indicates whether to include any background color or graphics when printing content.
	ShouldPrintBackgrounds() bool
	SetShouldPrintBackgrounds(value bool)

	// Topic: Setting Behavior Preferences

	// A Boolean value that indicates whether pressing the tab key changes the focus to links and form controls.
	TabFocusesLinks() bool
	SetTabFocusesLinks(value bool)
	// A Boolean value that indicates whether to allow people to select or otherwise interact with text.
	TextInteractionEnabled() bool
	SetTextInteractionEnabled(value bool)
	// A Boolean value that indicates whether a web view can display content full screen.
	ElementFullscreenEnabled() bool
	SetElementFullscreenEnabled(value bool)
	// A policy you set to specify how a web view that’s not in a window handles tasks.
	InactiveSchedulingPolicy() WKInactiveSchedulingPolicy
	SetInactiveSchedulingPolicy(value WKInactiveSchedulingPolicy)

	// Topic: Setting Java and JavaScript Preferences

	// A Boolean value that indicates whether JavaScript can open windows without user interaction.
	JavaScriptCanOpenWindowsAutomatically() bool
	SetJavaScriptCanOpenWindowsAutomatically(value bool)
	// A Boolean that indicates whether to apply site-specific compatibility workarounds.
	SiteSpecificQuirksModeEnabled() bool
	SetSiteSpecificQuirksModeEnabled(value bool)

	// Topic: Setting Fraud Warning Preferences

	// A Boolean value that indicates whether the web view shows warnings for suspected fraudulent content, such as malware or phishing attemps.
	FraudulentWebsiteWarningEnabled() bool
	SetFraudulentWebsiteWarningEnabled(value bool)

	// The object that manages the preference-related settings for the web view.
	Preferences() IWKPreferences
	SetPreferences(value IWKPreferences)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p WKPreferences) Init() WKPreferences {
	rv := objc.Send[WKPreferences](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p WKPreferences) Autorelease() WKPreferences {
	rv := objc.Send[WKPreferences](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKPreferences creates a new WKPreferences instance.
func NewWKPreferences() WKPreferences {
	class := getWKPreferencesClass()
	rv := objc.Send[WKPreferences](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (p WKPreferences) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The minimum font size, in points.
//
// # Discussion
//
// The default value is `0`.
//
// See: https://developer.apple.com/documentation/WebKit/WKPreferences/minimumFontSize
func (p WKPreferences) MinimumFontSize() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("minimumFontSize"))
	return rv
}
func (p WKPreferences) SetMinimumFontSize(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setMinimumFontSize:"), value)
}

// A Boolean value that indicates whether to include any background color or
// graphics when printing content.
//
// # Discussion
//
// The default value for this preference is false.
//
// See: https://developer.apple.com/documentation/WebKit/WKPreferences/shouldPrintBackgrounds
func (p WKPreferences) ShouldPrintBackgrounds() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("shouldPrintBackgrounds"))
	return rv
}
func (p WKPreferences) SetShouldPrintBackgrounds(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setShouldPrintBackgrounds:"), value)
}

// A Boolean value that indicates whether pressing the tab key changes the
// focus to links and form controls.
//
// # Discussion
//
// When the value of this property is true, the web view includes links and
// form controls in the set of items that may receive focus. Pressing the
// Option key temporarily reverses this preference.
//
// See: https://developer.apple.com/documentation/WebKit/WKPreferences/tabFocusesLinks
func (p WKPreferences) TabFocusesLinks() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("tabFocusesLinks"))
	return rv
}
func (p WKPreferences) SetTabFocusesLinks(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setTabFocusesLinks:"), value)
}

// A Boolean value that indicates whether to allow people to select or
// otherwise interact with text.
//
// # Discussion
//
// The default value for this preference is true on macOS and iOS. On watchOS,
// the default value is false.
//
// See: https://developer.apple.com/documentation/WebKit/WKPreferences/isTextInteractionEnabled
func (p WKPreferences) TextInteractionEnabled() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isTextInteractionEnabled"))
	return rv
}
func (p WKPreferences) SetTextInteractionEnabled(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setTextInteractionEnabled:"), value)
}

// A Boolean value that indicates whether a web view can display content full
// screen.
//
// # Discussion
//
// The default value for this preference is false.
//
// See: https://developer.apple.com/documentation/WebKit/WKPreferences/isElementFullscreenEnabled
func (p WKPreferences) ElementFullscreenEnabled() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isElementFullscreenEnabled"))
	return rv
}
func (p WKPreferences) SetElementFullscreenEnabled(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setElementFullscreenEnabled:"), value)
}

// A policy you set to specify how a web view that’s not in a window handles
// tasks.
//
// # Discussion
//
// Set this to indicate how a web view that’s not in a window handles tasks;
// for example, when the web view is in a background tab in a browser. The
// default value is [WKInactiveSchedulingPolicySuspend].
//
// A web view that’s not in a window is exempted from this policy if it is
// playing media, performing media capture, or performing another
// user-interactive activity.
//
// See: https://developer.apple.com/documentation/WebKit/WKPreferences/inactiveSchedulingPolicy-swift.property
func (p WKPreferences) InactiveSchedulingPolicy() WKInactiveSchedulingPolicy {
	rv := objc.Send[WKInactiveSchedulingPolicy](p.ID, objc.Sel("inactiveSchedulingPolicy"))
	return WKInactiveSchedulingPolicy(rv)
}
func (p WKPreferences) SetInactiveSchedulingPolicy(value WKInactiveSchedulingPolicy) {
	objc.Send[struct{}](p.ID, objc.Sel("setInactiveSchedulingPolicy:"), value)
}

// A Boolean value that indicates whether JavaScript can open windows without
// user interaction.
//
// # Discussion
//
// The default value is false in iOS and true in macOS.
//
// See: https://developer.apple.com/documentation/WebKit/WKPreferences/javaScriptCanOpenWindowsAutomatically
func (p WKPreferences) JavaScriptCanOpenWindowsAutomatically() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("javaScriptCanOpenWindowsAutomatically"))
	return rv
}
func (p WKPreferences) SetJavaScriptCanOpenWindowsAutomatically(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setJavaScriptCanOpenWindowsAutomatically:"), value)
}

// A Boolean that indicates whether to apply site-specific compatibility
// workarounds.
//
// # Discussion
//
// The default value for this preference is true.
//
// See: https://developer.apple.com/documentation/WebKit/WKPreferences/isSiteSpecificQuirksModeEnabled
func (p WKPreferences) SiteSpecificQuirksModeEnabled() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isSiteSpecificQuirksModeEnabled"))
	return rv
}
func (p WKPreferences) SetSiteSpecificQuirksModeEnabled(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setSiteSpecificQuirksModeEnabled:"), value)
}

// A Boolean value that indicates whether the web view shows warnings for
// suspected fraudulent content, such as malware or phishing attemps.
//
// # Discussion
//
// The default value of this property is true.
//
// See: https://developer.apple.com/documentation/WebKit/WKPreferences/isFraudulentWebsiteWarningEnabled
func (p WKPreferences) FraudulentWebsiteWarningEnabled() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isFraudulentWebsiteWarningEnabled"))
	return rv
}
func (p WKPreferences) SetFraudulentWebsiteWarningEnabled(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setFraudulentWebsiteWarningEnabled:"), value)
}

// The object that manages the preference-related settings for the web view.
//
// See: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/preferences
func (p WKPreferences) Preferences() IWKPreferences {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("preferences"))
	return WKPreferencesFromID(objc.ID(rv))
}
func (p WKPreferences) SetPreferences(value IWKPreferences) {
	objc.Send[struct{}](p.ID, objc.Sel("setPreferences:"), value)
}
