// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKWebViewConfiguration] class.
var (
	_WKWebViewConfigurationClass     WKWebViewConfigurationClass
	_WKWebViewConfigurationClassOnce sync.Once
)

func getWKWebViewConfigurationClass() WKWebViewConfigurationClass {
	_WKWebViewConfigurationClassOnce.Do(func() {
		_WKWebViewConfigurationClass = WKWebViewConfigurationClass{class: objc.GetClass("WKWebViewConfiguration")}
	})
	return _WKWebViewConfigurationClass
}

// GetWKWebViewConfigurationClass returns the class object for WKWebViewConfiguration.
func GetWKWebViewConfigurationClass() WKWebViewConfigurationClass {
	return getWKWebViewConfigurationClass()
}

type WKWebViewConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKWebViewConfigurationClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKWebViewConfigurationClass) Alloc() WKWebViewConfiguration {
	rv := objc.Send[WKWebViewConfiguration](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// A collection of properties that you use to initialize a web view.
//
// # Overview
//
// A [WKWebViewConfiguration] object provides information about how to
// configure a [WKWebView] object. Use your configuration object to specify:
//
// - The initial cookies to make available to your web content - Handlers for
// any custom URL schemes your web content uses - Settings for how to handle
// media content - Information about how to manage selections within the web
// view - Custom scripts to inject into the webpage - Custom rules that
// determine how to render content
//
// You create a [WKWebViewConfiguration] object in your code, configure its
// properties, and pass it to the initializer of your [WKWebView] object. The
// web view incorporates your configuration settings only at creation time;
// you cannot change those settings dynamically later.
//
// # Configuring the web view’s behavior
//
//   - [WKWebViewConfiguration.WebsiteDataStore]: The object you use to get and set the site’s cookies and to track the cached data objects.
//   - [WKWebViewConfiguration.SetWebsiteDataStore]
//   - [WKWebViewConfiguration.UserContentController]: The object that coordinates interactions between your app’s native code and the webpage’s scripts and other content.
//   - [WKWebViewConfiguration.SetUserContentController]
//   - [WKWebViewConfiguration.ApplicationNameForUserAgent]: The app name that appears in the user agent string.
//   - [WKWebViewConfiguration.SetApplicationNameForUserAgent]
//   - [WKWebViewConfiguration.LimitsNavigationsToAppBoundDomains]: A Boolean value that indicates whether the web view limits navigation to pages within the app’s domain.
//   - [WKWebViewConfiguration.SetLimitsNavigationsToAppBoundDomains]
//   - [WKWebViewConfiguration.UpgradeKnownHostsToHTTPS]: A Boolean value that indicates whether the web view should automatically upgrade supported HTTP requests to HTTPS.
//   - [WKWebViewConfiguration.SetUpgradeKnownHostsToHTTPS]
//
// # Configuring the web view’s preferences
//
//   - [WKWebViewConfiguration.Preferences]: The object that manages the preference-related settings for the web view.
//   - [WKWebViewConfiguration.SetPreferences]
//   - [WKWebViewConfiguration.DefaultWebpagePreferences]: The default preferences to use when loading and rendering content.
//   - [WKWebViewConfiguration.SetDefaultWebpagePreferences]
//
// # Adding handlers for custom URL schemes
//
//   - [WKWebViewConfiguration.SetURLSchemeHandlerForURLScheme]: Registers an object to load resources associated with the specified URL scheme.
//   - [WKWebViewConfiguration.UrlSchemeHandlerForURLScheme]: Returns the currently registered handler object for the specified URL scheme.
//
// # Configuring the rendering behavior
//
//   - [WKWebViewConfiguration.SuppressesIncrementalRendering]: A Boolean value that indicates whether the web view suppresses content rendering until the content is fully loaded into memory.
//   - [WKWebViewConfiguration.SetSuppressesIncrementalRendering]
//
// # Setting media playback preferences
//
//   - [WKWebViewConfiguration.AllowsAirPlayForMediaPlayback]: A Boolean value that indicates whether the web view allows media playback over AirPlay.
//   - [WKWebViewConfiguration.SetAllowsAirPlayForMediaPlayback]
//   - [WKWebViewConfiguration.MediaTypesRequiringUserActionForPlayback]: The media types that require a user gesture to begin playing.
//   - [WKWebViewConfiguration.SetMediaTypesRequiringUserActionForPlayback]
//
// # Selecting user interface directionality
//
//   - [WKWebViewConfiguration.UserInterfaceDirectionPolicy]: The directionality of user interface elements.
//   - [WKWebViewConfiguration.SetUserInterfaceDirectionPolicy]
//
// # Instance Properties
//
//   - [WKWebViewConfiguration.AllowsInlinePredictions]
//   - [WKWebViewConfiguration.SetAllowsInlinePredictions]
//   - [WKWebViewConfiguration.ShowsSystemScreenTimeBlockingView]
//   - [WKWebViewConfiguration.SetShowsSystemScreenTimeBlockingView]
//   - [WKWebViewConfiguration.SupportsAdaptiveImageGlyph]
//   - [WKWebViewConfiguration.SetSupportsAdaptiveImageGlyph]
//   - [WKWebViewConfiguration.WebExtensionController]
//   - [WKWebViewConfiguration.SetWebExtensionController]
//   - [WKWebViewConfiguration.WritingToolsBehavior]
//   - [WKWebViewConfiguration.SetWritingToolsBehavior]
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration
type WKWebViewConfiguration struct {
	objectivec.Object
}

// WKWebViewConfigurationFromID constructs a [WKWebViewConfiguration] from an objc.ID.
//
// A collection of properties that you use to initialize a web view.
func WKWebViewConfigurationFromID(id objc.ID) WKWebViewConfiguration {
	return WKWebViewConfiguration{objectivec.Object{ID: id}}
}

// NOTE: WKWebViewConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKWebViewConfiguration] class.
//
// # Configuring the web view’s behavior
//
//   - [IWKWebViewConfiguration.WebsiteDataStore]: The object you use to get and set the site’s cookies and to track the cached data objects.
//   - [IWKWebViewConfiguration.SetWebsiteDataStore]
//   - [IWKWebViewConfiguration.UserContentController]: The object that coordinates interactions between your app’s native code and the webpage’s scripts and other content.
//   - [IWKWebViewConfiguration.SetUserContentController]
//   - [IWKWebViewConfiguration.ApplicationNameForUserAgent]: The app name that appears in the user agent string.
//   - [IWKWebViewConfiguration.SetApplicationNameForUserAgent]
//   - [IWKWebViewConfiguration.LimitsNavigationsToAppBoundDomains]: A Boolean value that indicates whether the web view limits navigation to pages within the app’s domain.
//   - [IWKWebViewConfiguration.SetLimitsNavigationsToAppBoundDomains]
//   - [IWKWebViewConfiguration.UpgradeKnownHostsToHTTPS]: A Boolean value that indicates whether the web view should automatically upgrade supported HTTP requests to HTTPS.
//   - [IWKWebViewConfiguration.SetUpgradeKnownHostsToHTTPS]
//
// # Configuring the web view’s preferences
//
//   - [IWKWebViewConfiguration.Preferences]: The object that manages the preference-related settings for the web view.
//   - [IWKWebViewConfiguration.SetPreferences]
//   - [IWKWebViewConfiguration.DefaultWebpagePreferences]: The default preferences to use when loading and rendering content.
//   - [IWKWebViewConfiguration.SetDefaultWebpagePreferences]
//
// # Adding handlers for custom URL schemes
//
//   - [IWKWebViewConfiguration.SetURLSchemeHandlerForURLScheme]: Registers an object to load resources associated with the specified URL scheme.
//   - [IWKWebViewConfiguration.UrlSchemeHandlerForURLScheme]: Returns the currently registered handler object for the specified URL scheme.
//
// # Configuring the rendering behavior
//
//   - [IWKWebViewConfiguration.SuppressesIncrementalRendering]: A Boolean value that indicates whether the web view suppresses content rendering until the content is fully loaded into memory.
//   - [IWKWebViewConfiguration.SetSuppressesIncrementalRendering]
//
// # Setting media playback preferences
//
//   - [IWKWebViewConfiguration.AllowsAirPlayForMediaPlayback]: A Boolean value that indicates whether the web view allows media playback over AirPlay.
//   - [IWKWebViewConfiguration.SetAllowsAirPlayForMediaPlayback]
//   - [IWKWebViewConfiguration.MediaTypesRequiringUserActionForPlayback]: The media types that require a user gesture to begin playing.
//   - [IWKWebViewConfiguration.SetMediaTypesRequiringUserActionForPlayback]
//
// # Selecting user interface directionality
//
//   - [IWKWebViewConfiguration.UserInterfaceDirectionPolicy]: The directionality of user interface elements.
//   - [IWKWebViewConfiguration.SetUserInterfaceDirectionPolicy]
//
// # Instance Properties
//
//   - [IWKWebViewConfiguration.AllowsInlinePredictions]
//   - [IWKWebViewConfiguration.SetAllowsInlinePredictions]
//   - [IWKWebViewConfiguration.ShowsSystemScreenTimeBlockingView]
//   - [IWKWebViewConfiguration.SetShowsSystemScreenTimeBlockingView]
//   - [IWKWebViewConfiguration.SupportsAdaptiveImageGlyph]
//   - [IWKWebViewConfiguration.SetSupportsAdaptiveImageGlyph]
//   - [IWKWebViewConfiguration.WebExtensionController]
//   - [IWKWebViewConfiguration.SetWebExtensionController]
//   - [IWKWebViewConfiguration.WritingToolsBehavior]
//   - [IWKWebViewConfiguration.SetWritingToolsBehavior]
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration
type IWKWebViewConfiguration interface {
	objectivec.IObject

	// Topic: Configuring the web view’s behavior

	// The object you use to get and set the site’s cookies and to track the cached data objects.
	WebsiteDataStore() IWKWebsiteDataStore
	SetWebsiteDataStore(value IWKWebsiteDataStore)
	// The object that coordinates interactions between your app’s native code and the webpage’s scripts and other content.
	UserContentController() IWKUserContentController
	SetUserContentController(value IWKUserContentController)
	// The app name that appears in the user agent string.
	ApplicationNameForUserAgent() string
	SetApplicationNameForUserAgent(value string)
	// A Boolean value that indicates whether the web view limits navigation to pages within the app’s domain.
	LimitsNavigationsToAppBoundDomains() bool
	SetLimitsNavigationsToAppBoundDomains(value bool)
	// A Boolean value that indicates whether the web view should automatically upgrade supported HTTP requests to HTTPS.
	UpgradeKnownHostsToHTTPS() bool
	SetUpgradeKnownHostsToHTTPS(value bool)

	// Topic: Configuring the web view’s preferences

	// The object that manages the preference-related settings for the web view.
	Preferences() IWKPreferences
	SetPreferences(value IWKPreferences)
	// The default preferences to use when loading and rendering content.
	DefaultWebpagePreferences() IWKWebpagePreferences
	SetDefaultWebpagePreferences(value IWKWebpagePreferences)

	// Topic: Adding handlers for custom URL schemes

	// Registers an object to load resources associated with the specified URL scheme.
	SetURLSchemeHandlerForURLScheme(urlSchemeHandler WKURLSchemeHandler, urlScheme string)
	// Returns the currently registered handler object for the specified URL scheme.
	UrlSchemeHandlerForURLScheme(urlScheme string) WKURLSchemeHandler

	// Topic: Configuring the rendering behavior

	// A Boolean value that indicates whether the web view suppresses content rendering until the content is fully loaded into memory.
	SuppressesIncrementalRendering() bool
	SetSuppressesIncrementalRendering(value bool)

	// Topic: Setting media playback preferences

	// A Boolean value that indicates whether the web view allows media playback over AirPlay.
	AllowsAirPlayForMediaPlayback() bool
	SetAllowsAirPlayForMediaPlayback(value bool)
	// The media types that require a user gesture to begin playing.
	MediaTypesRequiringUserActionForPlayback() WKAudiovisualMediaTypes
	SetMediaTypesRequiringUserActionForPlayback(value WKAudiovisualMediaTypes)

	// Topic: Selecting user interface directionality

	// The directionality of user interface elements.
	UserInterfaceDirectionPolicy() WKUserInterfaceDirectionPolicy
	SetUserInterfaceDirectionPolicy(value WKUserInterfaceDirectionPolicy)

	// Topic: Instance Properties

	AllowsInlinePredictions() bool
	SetAllowsInlinePredictions(value bool)
	ShowsSystemScreenTimeBlockingView() bool
	SetShowsSystemScreenTimeBlockingView(value bool)
	SupportsAdaptiveImageGlyph() bool
	SetSupportsAdaptiveImageGlyph(value bool)
	WebExtensionController() IWKWebExtensionController
	SetWebExtensionController(value IWKWebExtensionController)
	WritingToolsBehavior() objectivec.IObject
	SetWritingToolsBehavior(value objectivec.IObject)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (w WKWebViewConfiguration) Init() WKWebViewConfiguration {
	rv := objc.Send[WKWebViewConfiguration](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WKWebViewConfiguration) Autorelease() WKWebViewConfiguration {
	rv := objc.Send[WKWebViewConfiguration](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKWebViewConfiguration creates a new WKWebViewConfiguration instance.
func NewWKWebViewConfiguration() WKWebViewConfiguration {
	class := getWKWebViewConfigurationClass()
	rv := objc.Send[WKWebViewConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Registers an object to load resources associated with the specified URL
// scheme.
//
// urlSchemeHandler: The object to handle the URL scheme. This object must adopt the
// [WKURLSchemeHandler] protocol.
//
// urlScheme: The URL scheme to handle. Scheme names are case sensitive, must start with
// an ASCII letter, and may contain only ASCII letters, numbers, the “`+`”
// character, the “`-`” character, and the “`.`” character. This
// method raises an [invalidArgumentException] if the scheme name is an empty
// string or contains any other characters.
//
// It is a programmer error to register a handler for a scheme WebKit already
// handles, such as `https`, and this method raises an
// [invalidArgumentException] if you try to do so. To determine whether WebKit
// handles a specific scheme, call the [HandlesURLScheme] class method of
// [WKWebView].
//
// # Discussion
//
// Use this method to register any custom resource types that your web content
// uses. For example, you might register a custom URL scheme for resources
// that you provide programmatically from your app.
//
// It is a programmer error to call this method more than once for the same
// scheme.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/setURLSchemeHandler(_:forURLScheme:)
//
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
func (w WKWebViewConfiguration) SetURLSchemeHandlerForURLScheme(urlSchemeHandler WKURLSchemeHandler, urlScheme string) {
	objc.Send[objc.ID](w.ID, objc.Sel("setURLSchemeHandler:forURLScheme:"), urlSchemeHandler, objc.String(urlScheme))
}

// Returns the currently registered handler object for the specified URL
// scheme.
//
// urlScheme: The scheme to look up. Scheme names are case sensitive, must start with an
// ASCII letter, and may contain only ASCII letters, numbers, the “`+`”
// character, the “`-`” character, and the “`.`” character. If this
// parameter contains an empty string or the scheme name includes invalid
// characters, this method returns `nil`.
//
// # Return Value
//
// The handler object for the specified scheme, or `nil` if the scheme has no
// handler.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/urlSchemeHandler(forURLScheme:)
func (w WKWebViewConfiguration) UrlSchemeHandlerForURLScheme(urlScheme string) WKURLSchemeHandler {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("urlSchemeHandlerForURLScheme:"), objc.String(urlScheme))
	return WKURLSchemeHandlerObjectFromID(rv)
}
func (w WKWebViewConfiguration) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](w.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The object you use to get and set the site’s cookies and to track the
// cached data objects.
//
// # Discussion
//
// If you don’t assign a value to this property, the configuration object
// uses the default data store object to store data persistently. To create a
// private web-browsing session, create a nonpersistent data store using the
// [NonPersistentDataStore] method and assign it to this property. For more
// information, see [WKWebsiteDataStore].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/websiteDataStore
func (w WKWebViewConfiguration) WebsiteDataStore() IWKWebsiteDataStore {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("websiteDataStore"))
	return WKWebsiteDataStoreFromID(objc.ID(rv))
}
func (w WKWebViewConfiguration) SetWebsiteDataStore(value IWKWebsiteDataStore) {
	objc.Send[struct{}](w.ID, objc.Sel("setWebsiteDataStore:"), value)
}

// The object that coordinates interactions between your app’s native code
// and the webpage’s scripts and other content.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/userContentController
func (w WKWebViewConfiguration) UserContentController() IWKUserContentController {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("userContentController"))
	return WKUserContentControllerFromID(objc.ID(rv))
}
func (w WKWebViewConfiguration) SetUserContentController(value IWKUserContentController) {
	objc.Send[struct{}](w.ID, objc.Sel("setUserContentController:"), value)
}

// The app name that appears in the user agent string.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/applicationNameForUserAgent
func (w WKWebViewConfiguration) ApplicationNameForUserAgent() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("applicationNameForUserAgent"))
	return foundation.NSStringFromID(rv).String()
}
func (w WKWebViewConfiguration) SetApplicationNameForUserAgent(value string) {
	objc.Send[struct{}](w.ID, objc.Sel("setApplicationNameForUserAgent:"), objc.String(value))
}

// A Boolean value that indicates whether the web view limits navigation to
// pages within the app’s domain.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/limitsNavigationsToAppBoundDomains
func (w WKWebViewConfiguration) LimitsNavigationsToAppBoundDomains() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("limitsNavigationsToAppBoundDomains"))
	return rv
}
func (w WKWebViewConfiguration) SetLimitsNavigationsToAppBoundDomains(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setLimitsNavigationsToAppBoundDomains:"), value)
}

// A Boolean value that indicates whether the web view should automatically
// upgrade supported HTTP requests to HTTPS.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/upgradeKnownHostsToHTTPS
func (w WKWebViewConfiguration) UpgradeKnownHostsToHTTPS() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("upgradeKnownHostsToHTTPS"))
	return rv
}
func (w WKWebViewConfiguration) SetUpgradeKnownHostsToHTTPS(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setUpgradeKnownHostsToHTTPS:"), value)
}

// The object that manages the preference-related settings for the web view.
//
// # Discussion
//
// Use the preferences object in this property to customize the rendering,
// JavaScript, and other preferences related to your web view. You can also
// change the preferences by assigning a new WKPreferences object to this
// property.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/preferences
func (w WKWebViewConfiguration) Preferences() IWKPreferences {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("preferences"))
	return WKPreferencesFromID(objc.ID(rv))
}
func (w WKWebViewConfiguration) SetPreferences(value IWKPreferences) {
	objc.Send[struct{}](w.ID, objc.Sel("setPreferences:"), value)
}

// The default preferences to use when loading and rendering content.
//
// # Discussion
//
// Use this property to specify the JavaScript settings and content mode for
// new webpages. When the web view navigates to a new page, it passes the
// default preferences to its navigation delegate, which can modify the
// preferences or pass them as they are.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/defaultWebpagePreferences
func (w WKWebViewConfiguration) DefaultWebpagePreferences() IWKWebpagePreferences {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("defaultWebpagePreferences"))
	return WKWebpagePreferencesFromID(objc.ID(rv))
}
func (w WKWebViewConfiguration) SetDefaultWebpagePreferences(value IWKWebpagePreferences) {
	objc.Send[struct{}](w.ID, objc.Sel("setDefaultWebpagePreferences:"), value)
}

// A Boolean value that indicates whether the web view suppresses content
// rendering until the content is fully loaded into memory.
//
// # Discussion
//
// The default value of this property is false.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/suppressesIncrementalRendering
func (w WKWebViewConfiguration) SuppressesIncrementalRendering() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("suppressesIncrementalRendering"))
	return rv
}
func (w WKWebViewConfiguration) SetSuppressesIncrementalRendering(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setSuppressesIncrementalRendering:"), value)
}

// A Boolean value that indicates whether the web view allows media playback
// over AirPlay.
//
// # Discussion
//
// The default value of this property is true.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/allowsAirPlayForMediaPlayback
func (w WKWebViewConfiguration) AllowsAirPlayForMediaPlayback() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("allowsAirPlayForMediaPlayback"))
	return rv
}
func (w WKWebViewConfiguration) SetAllowsAirPlayForMediaPlayback(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setAllowsAirPlayForMediaPlayback:"), value)
}

// The media types that require a user gesture to begin playing.
//
// # Discussion
//
// Use [WKAudiovisualMediaTypeNone] to indicate that no user gestures are
// required to begin playing media.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/mediaTypesRequiringUserActionForPlayback
func (w WKWebViewConfiguration) MediaTypesRequiringUserActionForPlayback() WKAudiovisualMediaTypes {
	rv := objc.Send[WKAudiovisualMediaTypes](w.ID, objc.Sel("mediaTypesRequiringUserActionForPlayback"))
	return WKAudiovisualMediaTypes(rv)
}
func (w WKWebViewConfiguration) SetMediaTypesRequiringUserActionForPlayback(value WKAudiovisualMediaTypes) {
	objc.Send[struct{}](w.ID, objc.Sel("setMediaTypesRequiringUserActionForPlayback:"), value)
}

// The directionality of user interface elements.
//
// # Discussion
//
// The default value is [WKUserInterfaceDirectionPolicyContent]. For other
// possible values, see [WKUserInterfaceDirectionPolicy].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/userInterfaceDirectionPolicy
//
// [WKUserInterfaceDirectionPolicy]: https://developer.apple.com/documentation/WebKit/WKUserInterfaceDirectionPolicy
func (w WKWebViewConfiguration) UserInterfaceDirectionPolicy() WKUserInterfaceDirectionPolicy {
	rv := objc.Send[WKUserInterfaceDirectionPolicy](w.ID, objc.Sel("userInterfaceDirectionPolicy"))
	return WKUserInterfaceDirectionPolicy(rv)
}
func (w WKWebViewConfiguration) SetUserInterfaceDirectionPolicy(value WKUserInterfaceDirectionPolicy) {
	objc.Send[struct{}](w.ID, objc.Sel("setUserInterfaceDirectionPolicy:"), value)
}

// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/allowsInlinePredictions
func (w WKWebViewConfiguration) AllowsInlinePredictions() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("allowsInlinePredictions"))
	return rv
}
func (w WKWebViewConfiguration) SetAllowsInlinePredictions(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setAllowsInlinePredictions:"), value)
}

// # Discussion
//
// A Boolean value indicating whether the System Screen Time blocking view
// should be shown.
//
// The default value is YES.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/showsSystemScreenTimeBlockingView
func (w WKWebViewConfiguration) ShowsSystemScreenTimeBlockingView() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("showsSystemScreenTimeBlockingView"))
	return rv
}
func (w WKWebViewConfiguration) SetShowsSystemScreenTimeBlockingView(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setShowsSystemScreenTimeBlockingView:"), value)
}

// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/supportsAdaptiveImageGlyph
func (w WKWebViewConfiguration) SupportsAdaptiveImageGlyph() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("supportsAdaptiveImageGlyph"))
	return rv
}
func (w WKWebViewConfiguration) SetSupportsAdaptiveImageGlyph(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setSupportsAdaptiveImageGlyph:"), value)
}

// # Discussion
//
// The web extension controller to associate with the web view.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/webExtensionController
func (w WKWebViewConfiguration) WebExtensionController() IWKWebExtensionController {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("webExtensionController"))
	return WKWebExtensionControllerFromID(objc.ID(rv))
}
func (w WKWebViewConfiguration) SetWebExtensionController(value IWKWebExtensionController) {
	objc.Send[struct{}](w.ID, objc.Sel("setWebExtensionController:"), value)
}

// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/writingToolsBehavior
func (w WKWebViewConfiguration) WritingToolsBehavior() objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("writingToolsBehavior"))
	return objectivec.Object{ID: rv}
}
func (w WKWebViewConfiguration) SetWritingToolsBehavior(value objectivec.IObject) {
	objc.Send[struct{}](w.ID, objc.Sel("setWritingToolsBehavior:"), value)
}
