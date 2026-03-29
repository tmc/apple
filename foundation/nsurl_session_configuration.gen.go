// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [URLSessionConfiguration] class.
var (
	_URLSessionConfigurationClass     URLSessionConfigurationClass
	_URLSessionConfigurationClassOnce sync.Once
)

func getURLSessionConfigurationClass() URLSessionConfigurationClass {
	_URLSessionConfigurationClassOnce.Do(func() {
		_URLSessionConfigurationClass = URLSessionConfigurationClass{class: objc.GetClass("NSURLSessionConfiguration")}
	})
	return _URLSessionConfigurationClass
}

// GetURLSessionConfigurationClass returns the class object for NSURLSessionConfiguration.
func GetURLSessionConfigurationClass() URLSessionConfigurationClass {
	return getURLSessionConfigurationClass()
}

type URLSessionConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc URLSessionConfigurationClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc URLSessionConfigurationClass) Alloc() URLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A configuration object that defines behavior and policies for a URL
// session.
//
// # Overview
// 
// An [NSURLSessionConfiguration] object defines the behavior and policies to
// use when uploading and downloading data using an [NSURLSession] object.
// When uploading or downloading data, creating a configuration object is
// always the first step you must take. You use this object to configure the
// timeout values, caching policies, connection requirements, and other types
// of information that you intend to use with your [NSURLSession] object.
// 
// It is important to configure your [NSURLSessionConfiguration] object
// appropriately before using it to initialize a session object. Session
// objects make a copy of the configuration settings you provide and use those
// settings to configure the session. Once configured, the session object
// ignores any changes you make to the [NSURLSessionConfiguration] object. If
// you need to modify your transfer policies, you must update the session
// configuration object and use it to create a new [NSURLSession] object.
// 
// For more information about using configuration objects to create sessions,
// see [NSURLSession].
// 
// # Types of session configurations
// 
// The behavior and capabilities of a URL session are largely determined by
// the kind of configuration used to create the session.
// 
// The singleton shared session (which has no configuration object) is for
// basic requests. It’s not as customizable as sessions that you create, but
// it serves as a good starting point if you have very limited requirements.
// You access this session by calling the shared class method. See that
// method’s discussion for more information about its limitations.
// 
// Default sessions behave much like the shared session (unless you customize
// them further), but let you obtain data incrementally using a delegate. You
// can create a default session configuration by calling the default method on
// the URLSessionConfiguration class.
// 
// Ephemeral sessions are similar to default sessions, but they don’t write
// caches, cookies, or credentials to disk. You can create an ephemeral
// session configuration by calling the ephemeral method on the
// URLSessionConfiguration class.
// 
// Background sessions let you perform uploads and downloads of content in the
// background while your app isn’t running. You can create a background
// session configuration by calling the backgroundSessionConfiguration(_:)
// method on the URLSessionConfiguration class.
//
// # Setting general properties
//
//   - [URLSessionConfiguration.Identifier]: The background session identifier of the configuration object.
//   - [URLSessionConfiguration.HTTPAdditionalHeaders]: A dictionary of additional headers to send with requests.
//   - [URLSessionConfiguration.SetHTTPAdditionalHeaders]
//   - [URLSessionConfiguration.NetworkServiceType]: The type of network service for all tasks within network sessions to enable Cellular Network Slicing.
//   - [URLSessionConfiguration.SetNetworkServiceType]
//   - [URLSessionConfiguration.AllowsCellularAccess]: A Boolean value that determines whether connections should be made over a cellular network.
//   - [URLSessionConfiguration.SetAllowsCellularAccess]
//   - [URLSessionConfiguration.TimeoutIntervalForRequest]: The timeout interval to use when waiting for additional data.
//   - [URLSessionConfiguration.SetTimeoutIntervalForRequest]
//   - [URLSessionConfiguration.TimeoutIntervalForResource]: The maximum amount of time that a resource request should be allowed to take.
//   - [URLSessionConfiguration.SetTimeoutIntervalForResource]
//   - [URLSessionConfiguration.SharedContainerIdentifier]: The identifier for the shared container into which files in background URL sessions should be downloaded.
//   - [URLSessionConfiguration.SetSharedContainerIdentifier]
//   - [URLSessionConfiguration.WaitsForConnectivity]: A Boolean value that indicates whether the session should wait for connectivity to become available, or fail immediately.
//   - [URLSessionConfiguration.SetWaitsForConnectivity]
//   - [URLSessionConfiguration.UsesClassicLoadingMode]
//   - [URLSessionConfiguration.SetUsesClassicLoadingMode]
//
// # Setting cookie policies
//
//   - [URLSessionConfiguration.HTTPCookieAcceptPolicy]: A policy constant that determines when cookies should be accepted.
//   - [URLSessionConfiguration.SetHTTPCookieAcceptPolicy]
//   - [URLSessionConfiguration.HTTPShouldSetCookies]: A Boolean value that determines whether requests should contain cookies from the cookie store.
//   - [URLSessionConfiguration.SetHTTPShouldSetCookies]
//   - [URLSessionConfiguration.HTTPCookieStorage]: The cookie store for storing cookies within this session.
//   - [URLSessionConfiguration.SetHTTPCookieStorage]
//
// # Setting security policies
//
//   - [URLSessionConfiguration.TLSMinimumSupportedProtocolVersion]: The minimum TLS protocol version that the client should accept when making connections in this session.
//   - [URLSessionConfiguration.SetTLSMinimumSupportedProtocolVersion]
//   - [URLSessionConfiguration.TLSMaximumSupportedProtocolVersion]: The maximum TLS protocol version that the client should request when making connections in this session.
//   - [URLSessionConfiguration.SetTLSMaximumSupportedProtocolVersion]
//   - [URLSessionConfiguration.URLCredentialStorage]: A credential store that provides credentials for authentication.
//   - [URLSessionConfiguration.SetURLCredentialStorage]
//   - [URLSessionConfiguration.TLSMinimumSupportedProtocol]: The minimum TLS protocol to accept during protocol negotiation.
//   - [URLSessionConfiguration.SetTLSMinimumSupportedProtocol]
//   - [URLSessionConfiguration.TLSMaximumSupportedProtocol]: The maximum TLS protocol version that the client should request when making connections in this session.
//   - [URLSessionConfiguration.SetTLSMaximumSupportedProtocol]
//   - [URLSessionConfiguration.RequiresDNSSECValidation]
//   - [URLSessionConfiguration.SetRequiresDNSSECValidation]
//
// # Setting caching policies
//
//   - [URLSessionConfiguration.URLCache]: The URL cache for providing cached responses to requests within the session.
//   - [URLSessionConfiguration.SetURLCache]
//   - [URLSessionConfiguration.RequestCachePolicy]: A predefined constant that determines when to return a response from the cache.
//   - [URLSessionConfiguration.SetRequestCachePolicy]
//
// # Supporting background transfers
//
//   - [URLSessionConfiguration.SessionSendsLaunchEvents]: A Boolean value that indicates whether the app should be resumed or launched in the background when transfers finish.
//   - [URLSessionConfiguration.SetSessionSendsLaunchEvents]
//   - [URLSessionConfiguration.Discretionary]: A Boolean value that determines whether background tasks can be scheduled at the discretion of the system for optimal performance.
//   - [URLSessionConfiguration.SetDiscretionary]
//   - [URLSessionConfiguration.ShouldUseExtendedBackgroundIdleMode]: A Boolean value that indicates whether TCP connections should be kept open when the app moves to the background.
//   - [URLSessionConfiguration.SetShouldUseExtendedBackgroundIdleMode]
//
// # Supporting custom protocols
//
//   - [URLSessionConfiguration.ProtocolClasses]: An array of extra protocol subclasses that handle requests in a session.
//   - [URLSessionConfiguration.SetProtocolClasses]
//
// # Setting HTTP policy and proxy properties
//
//   - [URLSessionConfiguration.HTTPMaximumConnectionsPerHost]: The maximum number of simultaneous connections to make to a given host.
//   - [URLSessionConfiguration.SetHTTPMaximumConnectionsPerHost]
//   - [URLSessionConfiguration.HTTPShouldUsePipelining]: A Boolean value that determines whether the session should use HTTP pipelining.
//   - [URLSessionConfiguration.SetHTTPShouldUsePipelining]
//   - [URLSessionConfiguration.ConnectionProxyDictionary]: A dictionary containing information about the proxy to use within this session.
//   - [URLSessionConfiguration.SetConnectionProxyDictionary]
//
// # Supporting limited modes
//
//   - [URLSessionConfiguration.AllowsConstrainedNetworkAccess]: A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//   - [URLSessionConfiguration.SetAllowsConstrainedNetworkAccess]
//   - [URLSessionConfiguration.AllowsExpensiveNetworkAccess]: A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//   - [URLSessionConfiguration.SetAllowsExpensiveNetworkAccess]
//
// # Instance Properties
//
//   - [URLSessionConfiguration.AllowsUltraConstrainedNetworkAccess]
//   - [URLSessionConfiguration.SetAllowsUltraConstrainedNetworkAccess]
//   - [URLSessionConfiguration.EnablesEarlyData]
//   - [URLSessionConfiguration.SetEnablesEarlyData]
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration
type URLSessionConfiguration struct {
	objectivec.Object
}

// URLSessionConfigurationFromID constructs a [URLSessionConfiguration] from an objc.ID.
//
// A configuration object that defines behavior and policies for a URL
// session.
func URLSessionConfigurationFromID(id objc.ID) URLSessionConfiguration {
	return NSURLSessionConfiguration{objectivec.Object{ID: id}}
}

// NSURLSessionConfigurationFromID is an alias for [URLSessionConfigurationFromID] for cross-framework compatibility.
func NSURLSessionConfigurationFromID(id objc.ID) URLSessionConfiguration { return URLSessionConfigurationFromID(id) }
// NOTE: URLSessionConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [URLSessionConfiguration] class.
//
// # Setting general properties
//
//   - [IURLSessionConfiguration.Identifier]: The background session identifier of the configuration object.
//   - [IURLSessionConfiguration.HTTPAdditionalHeaders]: A dictionary of additional headers to send with requests.
//   - [IURLSessionConfiguration.SetHTTPAdditionalHeaders]
//   - [IURLSessionConfiguration.NetworkServiceType]: The type of network service for all tasks within network sessions to enable Cellular Network Slicing.
//   - [IURLSessionConfiguration.SetNetworkServiceType]
//   - [IURLSessionConfiguration.AllowsCellularAccess]: A Boolean value that determines whether connections should be made over a cellular network.
//   - [IURLSessionConfiguration.SetAllowsCellularAccess]
//   - [IURLSessionConfiguration.TimeoutIntervalForRequest]: The timeout interval to use when waiting for additional data.
//   - [IURLSessionConfiguration.SetTimeoutIntervalForRequest]
//   - [IURLSessionConfiguration.TimeoutIntervalForResource]: The maximum amount of time that a resource request should be allowed to take.
//   - [IURLSessionConfiguration.SetTimeoutIntervalForResource]
//   - [IURLSessionConfiguration.SharedContainerIdentifier]: The identifier for the shared container into which files in background URL sessions should be downloaded.
//   - [IURLSessionConfiguration.SetSharedContainerIdentifier]
//   - [IURLSessionConfiguration.WaitsForConnectivity]: A Boolean value that indicates whether the session should wait for connectivity to become available, or fail immediately.
//   - [IURLSessionConfiguration.SetWaitsForConnectivity]
//   - [IURLSessionConfiguration.UsesClassicLoadingMode]
//   - [IURLSessionConfiguration.SetUsesClassicLoadingMode]
//
// # Setting cookie policies
//
//   - [IURLSessionConfiguration.HTTPCookieAcceptPolicy]: A policy constant that determines when cookies should be accepted.
//   - [IURLSessionConfiguration.SetHTTPCookieAcceptPolicy]
//   - [IURLSessionConfiguration.HTTPShouldSetCookies]: A Boolean value that determines whether requests should contain cookies from the cookie store.
//   - [IURLSessionConfiguration.SetHTTPShouldSetCookies]
//   - [IURLSessionConfiguration.HTTPCookieStorage]: The cookie store for storing cookies within this session.
//   - [IURLSessionConfiguration.SetHTTPCookieStorage]
//
// # Setting security policies
//
//   - [IURLSessionConfiguration.TLSMinimumSupportedProtocolVersion]: The minimum TLS protocol version that the client should accept when making connections in this session.
//   - [IURLSessionConfiguration.SetTLSMinimumSupportedProtocolVersion]
//   - [IURLSessionConfiguration.TLSMaximumSupportedProtocolVersion]: The maximum TLS protocol version that the client should request when making connections in this session.
//   - [IURLSessionConfiguration.SetTLSMaximumSupportedProtocolVersion]
//   - [IURLSessionConfiguration.URLCredentialStorage]: A credential store that provides credentials for authentication.
//   - [IURLSessionConfiguration.SetURLCredentialStorage]
//   - [IURLSessionConfiguration.TLSMinimumSupportedProtocol]: The minimum TLS protocol to accept during protocol negotiation.
//   - [IURLSessionConfiguration.SetTLSMinimumSupportedProtocol]
//   - [IURLSessionConfiguration.TLSMaximumSupportedProtocol]: The maximum TLS protocol version that the client should request when making connections in this session.
//   - [IURLSessionConfiguration.SetTLSMaximumSupportedProtocol]
//   - [IURLSessionConfiguration.RequiresDNSSECValidation]
//   - [IURLSessionConfiguration.SetRequiresDNSSECValidation]
//
// # Setting caching policies
//
//   - [IURLSessionConfiguration.URLCache]: The URL cache for providing cached responses to requests within the session.
//   - [IURLSessionConfiguration.SetURLCache]
//   - [IURLSessionConfiguration.RequestCachePolicy]: A predefined constant that determines when to return a response from the cache.
//   - [IURLSessionConfiguration.SetRequestCachePolicy]
//
// # Supporting background transfers
//
//   - [IURLSessionConfiguration.SessionSendsLaunchEvents]: A Boolean value that indicates whether the app should be resumed or launched in the background when transfers finish.
//   - [IURLSessionConfiguration.SetSessionSendsLaunchEvents]
//   - [IURLSessionConfiguration.Discretionary]: A Boolean value that determines whether background tasks can be scheduled at the discretion of the system for optimal performance.
//   - [IURLSessionConfiguration.SetDiscretionary]
//   - [IURLSessionConfiguration.ShouldUseExtendedBackgroundIdleMode]: A Boolean value that indicates whether TCP connections should be kept open when the app moves to the background.
//   - [IURLSessionConfiguration.SetShouldUseExtendedBackgroundIdleMode]
//
// # Supporting custom protocols
//
//   - [IURLSessionConfiguration.ProtocolClasses]: An array of extra protocol subclasses that handle requests in a session.
//   - [IURLSessionConfiguration.SetProtocolClasses]
//
// # Setting HTTP policy and proxy properties
//
//   - [IURLSessionConfiguration.HTTPMaximumConnectionsPerHost]: The maximum number of simultaneous connections to make to a given host.
//   - [IURLSessionConfiguration.SetHTTPMaximumConnectionsPerHost]
//   - [IURLSessionConfiguration.HTTPShouldUsePipelining]: A Boolean value that determines whether the session should use HTTP pipelining.
//   - [IURLSessionConfiguration.SetHTTPShouldUsePipelining]
//   - [IURLSessionConfiguration.ConnectionProxyDictionary]: A dictionary containing information about the proxy to use within this session.
//   - [IURLSessionConfiguration.SetConnectionProxyDictionary]
//
// # Supporting limited modes
//
//   - [IURLSessionConfiguration.AllowsConstrainedNetworkAccess]: A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//   - [IURLSessionConfiguration.SetAllowsConstrainedNetworkAccess]
//   - [IURLSessionConfiguration.AllowsExpensiveNetworkAccess]: A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//   - [IURLSessionConfiguration.SetAllowsExpensiveNetworkAccess]
//
// # Instance Properties
//
//   - [IURLSessionConfiguration.AllowsUltraConstrainedNetworkAccess]
//   - [IURLSessionConfiguration.SetAllowsUltraConstrainedNetworkAccess]
//   - [IURLSessionConfiguration.EnablesEarlyData]
//   - [IURLSessionConfiguration.SetEnablesEarlyData]
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration
type IURLSessionConfiguration interface {
	objectivec.IObject
	NSCopying

	// Topic: Setting general properties

	// The background session identifier of the configuration object.
	Identifier() string
	// A dictionary of additional headers to send with requests.
	HTTPAdditionalHeaders() INSDictionary
	SetHTTPAdditionalHeaders(value INSDictionary)
	// The type of network service for all tasks within network sessions to enable Cellular Network Slicing.
	NetworkServiceType() NSURLRequestNetworkServiceType
	SetNetworkServiceType(value NSURLRequestNetworkServiceType)
	// A Boolean value that determines whether connections should be made over a cellular network.
	AllowsCellularAccess() bool
	SetAllowsCellularAccess(value bool)
	// The timeout interval to use when waiting for additional data.
	TimeoutIntervalForRequest() float64
	SetTimeoutIntervalForRequest(value float64)
	// The maximum amount of time that a resource request should be allowed to take.
	TimeoutIntervalForResource() float64
	SetTimeoutIntervalForResource(value float64)
	// The identifier for the shared container into which files in background URL sessions should be downloaded.
	SharedContainerIdentifier() string
	SetSharedContainerIdentifier(value string)
	// A Boolean value that indicates whether the session should wait for connectivity to become available, or fail immediately.
	WaitsForConnectivity() bool
	SetWaitsForConnectivity(value bool)
	UsesClassicLoadingMode() bool
	SetUsesClassicLoadingMode(value bool)

	// Topic: Setting cookie policies

	// A policy constant that determines when cookies should be accepted.
	HTTPCookieAcceptPolicy() NSHTTPCookieAcceptPolicy
	SetHTTPCookieAcceptPolicy(value NSHTTPCookieAcceptPolicy)
	// A Boolean value that determines whether requests should contain cookies from the cookie store.
	HTTPShouldSetCookies() bool
	SetHTTPShouldSetCookies(value bool)
	// The cookie store for storing cookies within this session.
	HTTPCookieStorage() INSHTTPCookieStorage
	SetHTTPCookieStorage(value INSHTTPCookieStorage)

	// Topic: Setting security policies

	// The minimum TLS protocol version that the client should accept when making connections in this session.
	TLSMinimumSupportedProtocolVersion() uint16
	SetTLSMinimumSupportedProtocolVersion(value uint16)
	// The maximum TLS protocol version that the client should request when making connections in this session.
	TLSMaximumSupportedProtocolVersion() uint16
	SetTLSMaximumSupportedProtocolVersion(value uint16)
	// A credential store that provides credentials for authentication.
	URLCredentialStorage() INSURLCredentialStorage
	SetURLCredentialStorage(value INSURLCredentialStorage)
	// The minimum TLS protocol to accept during protocol negotiation.
	TLSMinimumSupportedProtocol() objectivec.IObject
	SetTLSMinimumSupportedProtocol(value objectivec.IObject)
	// The maximum TLS protocol version that the client should request when making connections in this session.
	TLSMaximumSupportedProtocol() objectivec.IObject
	SetTLSMaximumSupportedProtocol(value objectivec.IObject)
	RequiresDNSSECValidation() bool
	SetRequiresDNSSECValidation(value bool)

	// Topic: Setting caching policies

	// The URL cache for providing cached responses to requests within the session.
	URLCache() INSURLCache
	SetURLCache(value INSURLCache)
	// A predefined constant that determines when to return a response from the cache.
	RequestCachePolicy() NSURLRequestCachePolicy
	SetRequestCachePolicy(value NSURLRequestCachePolicy)

	// Topic: Supporting background transfers

	// A Boolean value that indicates whether the app should be resumed or launched in the background when transfers finish.
	SessionSendsLaunchEvents() bool
	SetSessionSendsLaunchEvents(value bool)
	// A Boolean value that determines whether background tasks can be scheduled at the discretion of the system for optimal performance.
	Discretionary() bool
	SetDiscretionary(value bool)
	// A Boolean value that indicates whether TCP connections should be kept open when the app moves to the background.
	ShouldUseExtendedBackgroundIdleMode() bool
	SetShouldUseExtendedBackgroundIdleMode(value bool)

	// Topic: Supporting custom protocols

	// An array of extra protocol subclasses that handle requests in a session.
	ProtocolClasses() []objc.Class
	SetProtocolClasses(value []objc.Class)

	// Topic: Setting HTTP policy and proxy properties

	// The maximum number of simultaneous connections to make to a given host.
	HTTPMaximumConnectionsPerHost() int
	SetHTTPMaximumConnectionsPerHost(value int)
	// A Boolean value that determines whether the session should use HTTP pipelining.
	HTTPShouldUsePipelining() bool
	SetHTTPShouldUsePipelining(value bool)
	// A dictionary containing information about the proxy to use within this session.
	ConnectionProxyDictionary() INSDictionary
	SetConnectionProxyDictionary(value INSDictionary)

	// Topic: Supporting limited modes

	// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
	AllowsConstrainedNetworkAccess() bool
	SetAllowsConstrainedNetworkAccess(value bool)
	// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
	AllowsExpensiveNetworkAccess() bool
	SetAllowsExpensiveNetworkAccess(value bool)

	// Topic: Instance Properties

	AllowsUltraConstrainedNetworkAccess() bool
	SetAllowsUltraConstrainedNetworkAccess(value bool)
	EnablesEarlyData() bool
	SetEnablesEarlyData(value bool)

	// A copy of the configuration object for this session.
	Configuration() INSURLSessionConfiguration
	SetConfiguration(value INSURLSessionConfiguration)
	// An array of proxy configuration objects containing information about the proxies to use within this session.
	ProxyConfigurations() []objectivec.Object
	SetProxyConfigurations(value []objectivec.Object)
}

// Init initializes the instance.
func (u URLSessionConfiguration) Init() URLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u URLSessionConfiguration) Autorelease() URLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionConfiguration creates a new URLSessionConfiguration instance.
func NewURLSessionConfiguration() URLSessionConfiguration {
	class := getURLSessionConfigurationClass()
	rv := objc.Send[URLSessionConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a session configuration object that allows HTTP and HTTPS uploads
// or downloads to be performed in the background.
//
// identifier: The unique identifier for the configuration object. This parameter must not
// be `nil` or an empty string.
//
// # Return Value
// 
// A configuration object that causes the system to perform upload and
// download tasks in a separate process.
//
// # Discussion
// 
// Use this method to initialize a configuration object suitable for
// transferring data files while the app runs in the background. A session
// configured with this object hands control of the transfers over to the
// system, which handles the transfers in a separate process. In iOS, this
// configuration makes it possible for transfers to continue even when the app
// itself is suspended or terminated.
// 
// If an iOS app is terminated by the system and relaunched, the app can use
// the same `identifier` to create a new configuration object and session and
// to retrieve the status of transfers that were in progress at the time of
// termination. This behavior applies only for normal termination of the app
// by the system. If the user terminates the app from the multitasking screen,
// the system cancels all of the session’s background transfers. In
// addition, the system does not automatically relaunch apps that were force
// quit by the user. The user must explicitly relaunch the app before
// transfers can begin again.
// 
// You can configure an background session to schedule transfers at the
// discretion of the system for optimal performance using the [Discretionary]
// property. When transferring large amounts of data, you are encouraged to
// set the value of this property to [true]. For an example of using the
// background configuration, see [Downloading files in the background].
//
// [Downloading files in the background]: https://developer.apple.com/documentation/Foundation/downloading-files-in-the-background
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/background(withIdentifier:)
func (_URLSessionConfigurationClass URLSessionConfigurationClass) BackgroundSessionConfigurationWithIdentifier(identifier string) URLSessionConfiguration {
	rv := objc.Send[objc.ID](objc.ID(_URLSessionConfigurationClass.class), objc.Sel("backgroundSessionConfigurationWithIdentifier:"), objc.String(identifier))
	return NSURLSessionConfigurationFromID(rv)
}

// The background session identifier of the configuration object.
//
// # Discussion
// 
// The value of this property is set only when you use the
// [BackgroundSessionConfigurationWithIdentifier] method to create the
// configuration object. The string uniquely identifies a background session
// object. In iOS, you use this string in cases where the app was terminated
// while transfers were occurring in the background. When the app relaunches,
// it uses the string to recreate the configuration and session objects
// associated with the transfers.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/identifier
func (u URLSessionConfiguration) Identifier() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("identifier"))
	return NSStringFromID(rv).String()
}
// A dictionary of additional headers to send with requests.
//
// # Discussion
// 
// This property specifies additional headers that are added to all tasks
// within sessions based on this configuration. For example, you might set the
// `User-Agent` header so that it is automatically included in every request
// your app makes through sessions based on this configuration.
// 
// An [NSURLSession] object is designed to handle various aspects of the HTTP
// protocol for you. As a result, you should not modify the following headers:
// 
// - [Authorization] - [Connection] - [Host] - `Proxy-Authenticate` -
// `Proxy-Authorization` - `WWW-Authenticate`
// 
// Additionally, if the length of your upload body data can be determined
// automatically—for example, if you provide the body content with an
// [NSData] object—the value of `Content-Length` is set for you.
// 
// If the same header appears in both this array and the request object (where
// applicable), the request object’s value takes precedence.
// 
// The default value is an empty array.
//
// [NSData]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/PropertyLists/OldStylePlists/OldStylePLists.html#//apple_ref/doc/uid/20001012-47169
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpAdditionalHeaders
func (u URLSessionConfiguration) HTTPAdditionalHeaders() INSDictionary {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("HTTPAdditionalHeaders"))
	return NSDictionaryFromID(objc.ID(rv))
}
func (u URLSessionConfiguration) SetHTTPAdditionalHeaders(value INSDictionary) {
	objc.Send[struct{}](u.ID, objc.Sel("setHTTPAdditionalHeaders:"), value)
}
// The type of network service for all tasks within network sessions to enable
// Cellular Network Slicing.
//
// # Discussion
// 
// To enable Cellular Network Slicing, you need to set the appropriate
// entitlements and properties.
// 
// Set the entitlements in your property list for [5G Network Slicing App
// Category] and [5G Network Slicing Traffic Category]. If you don’t entitle
// your app by specifying both these entitlements, your apps network
// connections won’t be using Cellular Network Slicing, even if supported by
// the carrier.
// 
// At the time of network flow creation, set this to the appropriate
// [NSURLRequest.NetworkServiceType] for your application type.
//
// [5G Network Slicing App Category]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.developer.networking.slicing.appcategory
// [5G Network Slicing Traffic Category]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.developer.networking.slicing.trafficcategory
// [NSURLRequest.NetworkServiceType]: https://developer.apple.com/documentation/Foundation/NSURLRequest/NetworkServiceType-swift.enum
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/networkServiceType
func (u URLSessionConfiguration) NetworkServiceType() NSURLRequestNetworkServiceType {
	rv := objc.Send[NSURLRequestNetworkServiceType](u.ID, objc.Sel("networkServiceType"))
	return NSURLRequestNetworkServiceType(rv)
}
func (u URLSessionConfiguration) SetNetworkServiceType(value NSURLRequestNetworkServiceType) {
	objc.Send[struct{}](u.ID, objc.Sel("setNetworkServiceType:"), value)
}
// A Boolean value that determines whether connections should be made over a
// cellular network.
//
// # Discussion
// 
// This property controls whether tasks in sessions based on this session
// configuration are allowed to make connections over a cellular network.
// 
// The default value is [true].
// 
// For more information, read [Restrict Cellular Networking Correctly].
//
// [Restrict Cellular Networking Correctly]: https://developer.apple.com/library/archive/documentation/NetworkingInternetWeb/Conceptual/NetworkingOverview/Platform-SpecificNetworkingTechnologies/Platform-SpecificNetworkingTechnologies.html#//apple_ref/doc/uid/TP40010220-CH212-SW9
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsCellularAccess
func (u URLSessionConfiguration) AllowsCellularAccess() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("allowsCellularAccess"))
	return rv
}
func (u URLSessionConfiguration) SetAllowsCellularAccess(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setAllowsCellularAccess:"), value)
}
// The timeout interval to use when waiting for additional data.
//
// # Discussion
// 
// This property determines the request timeout interval for all tasks within
// sessions based on this configuration. The request timeout interval controls
// how long (in seconds) a task should wait for additional data to arrive
// before giving up. The timer associated with this value is reset whenever
// new data arrives. When the request timer reaches the specified interval
// without receiving any new data, it triggers a timeout.
// 
// The default value is `60`.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/timeoutIntervalForRequest
func (u URLSessionConfiguration) TimeoutIntervalForRequest() float64 {
	rv := objc.Send[NSTimeInterval](u.ID, objc.Sel("timeoutIntervalForRequest"))
	return float64(rv)
}
func (u URLSessionConfiguration) SetTimeoutIntervalForRequest(value float64) {
	objc.Send[struct{}](u.ID, objc.Sel("setTimeoutIntervalForRequest:"), value)
}
// The maximum amount of time that a resource request should be allowed to
// take.
//
// # Discussion
// 
// This property determines the resource timeout interval for all tasks within
// sessions based on this configuration. The resource timeout interval
// controls how long (in seconds) to wait for an entire resource to transfer
// before giving up. The resource timer starts when the request is initiated
// and counts until either the request completes or this timeout interval is
// reached, whichever comes first.
// 
// The default value is 7 days.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/timeoutIntervalForResource
func (u URLSessionConfiguration) TimeoutIntervalForResource() float64 {
	rv := objc.Send[NSTimeInterval](u.ID, objc.Sel("timeoutIntervalForResource"))
	return float64(rv)
}
func (u URLSessionConfiguration) SetTimeoutIntervalForResource(value float64) {
	objc.Send[struct{}](u.ID, objc.Sel("setTimeoutIntervalForResource:"), value)
}
// The identifier for the shared container into which files in background URL
// sessions should be downloaded.
//
// # Discussion
// 
// To create a URL session for use by an app extension, set this property to a
// valid identifier for a container shared between the app extension and its
// containing app.
// 
// For information about app extensions, see [App Extension Programming
// Guide].
//
// [App Extension Programming Guide]: https://developer.apple.com/library/archive/documentation/General/Conceptual/ExtensibilityPG/index.html#//apple_ref/doc/uid/TP40014214
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/sharedContainerIdentifier
func (u URLSessionConfiguration) SharedContainerIdentifier() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("sharedContainerIdentifier"))
	return NSStringFromID(rv).String()
}
func (u URLSessionConfiguration) SetSharedContainerIdentifier(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setSharedContainerIdentifier:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/usesClassicLoadingMode
func (u URLSessionConfiguration) UsesClassicLoadingMode() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("usesClassicLoadingMode"))
	return rv
}
func (u URLSessionConfiguration) SetUsesClassicLoadingMode(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setUsesClassicLoadingMode:"), value)
}
// A policy constant that determines when cookies should be accepted.
//
// # Discussion
// 
// This property determines the cookie accept policy for all tasks within
// sessions based on this configuration.
// 
// The default value is [HTTPCookie.AcceptPolicy.onlyFromMainDocumentDomain].
// You can change it to any of the constants defined in the
// [HTTPCookie.AcceptPolicy] enumerated type.
// 
// If you want more direct control over what cookies are accepted, set this
// value to [HTTPCookie.AcceptPolicy.never] and then use the [AllHeaderFields]
// and [CookiesWithResponseHeaderFieldsForURL] methods to extract cookies from
// the URL response object yourself.
//
// [HTTPCookie.AcceptPolicy.never]: https://developer.apple.com/documentation/Foundation/HTTPCookie/AcceptPolicy/never
// [HTTPCookie.AcceptPolicy.onlyFromMainDocumentDomain]: https://developer.apple.com/documentation/Foundation/HTTPCookie/AcceptPolicy/onlyFromMainDocumentDomain
// [HTTPCookie.AcceptPolicy]: https://developer.apple.com/documentation/Foundation/HTTPCookie/AcceptPolicy
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpCookieAcceptPolicy
func (u URLSessionConfiguration) HTTPCookieAcceptPolicy() NSHTTPCookieAcceptPolicy {
	rv := objc.Send[NSHTTPCookieAcceptPolicy](u.ID, objc.Sel("HTTPCookieAcceptPolicy"))
	return NSHTTPCookieAcceptPolicy(rv)
}
func (u URLSessionConfiguration) SetHTTPCookieAcceptPolicy(value NSHTTPCookieAcceptPolicy) {
	objc.Send[struct{}](u.ID, objc.Sel("setHTTPCookieAcceptPolicy:"), value)
}
// A Boolean value that determines whether requests should contain cookies
// from the cookie store.
//
// # Discussion
// 
// This property controls whether tasks within sessions based on this
// configuration should automatically provide cookies from the shared cookie
// store when making requests.
// 
// If you want to provide cookies yourself, set this value to [false] and
// provide a [Cookie] header either through the session’s
// [HTTPAdditionalHeaders] property or on a per-request level using a custom
// [NSURLRequest] object.
// 
// The default value is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpShouldSetCookies
func (u URLSessionConfiguration) HTTPShouldSetCookies() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("HTTPShouldSetCookies"))
	return rv
}
func (u URLSessionConfiguration) SetHTTPShouldSetCookies(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setHTTPShouldSetCookies:"), value)
}
// The cookie store for storing cookies within this session.
//
// # Discussion
// 
// This property determines the cookie storage object used by all tasks within
// sessions based on this configuration.
// 
// To disable cookie storage, set this property to `nil`.
// 
// For default and background sessions, the default value is the
// [SharedHTTPCookieStorage] cookie storage object.
// 
// For [EphemeralSessionConfiguration] sessions, the default value is a
// private cookie storage object that stores data in memory only, and is
// destroyed when you invalidate the session.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpCookieStorage
func (u URLSessionConfiguration) HTTPCookieStorage() INSHTTPCookieStorage {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("HTTPCookieStorage"))
	return NSHTTPCookieStorageFromID(objc.ID(rv))
}
func (u URLSessionConfiguration) SetHTTPCookieStorage(value INSHTTPCookieStorage) {
	objc.Send[struct{}](u.ID, objc.Sel("setHTTPCookieStorage:"), value)
}
// The minimum TLS protocol version that the client should accept when making
// connections in this session.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/tlsMinimumSupportedProtocolVersion
func (u URLSessionConfiguration) TLSMinimumSupportedProtocolVersion() uint16 {
	rv := objc.Send[uint16](u.ID, objc.Sel("TLSMinimumSupportedProtocolVersion"))
	return rv
}
func (u URLSessionConfiguration) SetTLSMinimumSupportedProtocolVersion(value uint16) {
	objc.Send[struct{}](u.ID, objc.Sel("setTLSMinimumSupportedProtocolVersion:"), value)
}
// The maximum TLS protocol version that the client should request when making
// connections in this session.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/tlsMaximumSupportedProtocolVersion
func (u URLSessionConfiguration) TLSMaximumSupportedProtocolVersion() uint16 {
	rv := objc.Send[uint16](u.ID, objc.Sel("TLSMaximumSupportedProtocolVersion"))
	return rv
}
func (u URLSessionConfiguration) SetTLSMaximumSupportedProtocolVersion(value uint16) {
	objc.Send[struct{}](u.ID, objc.Sel("setTLSMaximumSupportedProtocolVersion:"), value)
}
// A credential store that provides credentials for authentication.
//
// # Discussion
// 
// This property determines the credential storage object used by tasks within
// sessions based on this configuration.
// 
// If you don’t want to use a credential store, set this property to `nil`.
// 
// For default and background sessions, the default value is the
// [SharedCredentialStorage] credential store object.
// 
// For [EphemeralSessionConfiguration] sessions, the default value is a
// private credential store object that stores data in memory only, and is
// destroyed when you invalidate the session.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/urlCredentialStorage
func (u URLSessionConfiguration) URLCredentialStorage() INSURLCredentialStorage {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("URLCredentialStorage"))
	return NSURLCredentialStorageFromID(objc.ID(rv))
}
func (u URLSessionConfiguration) SetURLCredentialStorage(value INSURLCredentialStorage) {
	objc.Send[struct{}](u.ID, objc.Sel("setURLCredentialStorage:"), value)
}
// The minimum TLS protocol to accept during protocol negotiation.
//
// # Discussion
// 
// This property determines the minimum supported TLS protocol version for
// tasks within sessions based on this configuration.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/tlsMinimumSupportedProtocol
func (u URLSessionConfiguration) TLSMinimumSupportedProtocol() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("TLSMinimumSupportedProtocol"))
	return objectivec.Object{ID: rv}
}
func (u URLSessionConfiguration) SetTLSMinimumSupportedProtocol(value objectivec.IObject) {
	objc.Send[struct{}](u.ID, objc.Sel("setTLSMinimumSupportedProtocol:"), value)
}
// The maximum TLS protocol version that the client should request when making
// connections in this session.
//
// # Discussion
// 
// This property determines the maximum supported TLS protocol version for
// tasks within sessions based on this configuration.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/tlsMaximumSupportedProtocol
func (u URLSessionConfiguration) TLSMaximumSupportedProtocol() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("TLSMaximumSupportedProtocol"))
	return objectivec.Object{ID: rv}
}
func (u URLSessionConfiguration) SetTLSMaximumSupportedProtocol(value objectivec.IObject) {
	objc.Send[struct{}](u.ID, objc.Sel("setTLSMaximumSupportedProtocol:"), value)
}
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/requiresDNSSECValidation
func (u URLSessionConfiguration) RequiresDNSSECValidation() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("requiresDNSSECValidation"))
	return rv
}
func (u URLSessionConfiguration) SetRequiresDNSSECValidation(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setRequiresDNSSECValidation:"), value)
}
// The URL cache for providing cached responses to requests within the
// session.
//
// # Discussion
// 
// This property determines the URL cache object used by tasks within sessions
// based on this configuration.
// 
// To disable caching, set this property to `nil`.
// 
// For default sessions, the default value is the shared URL cache object.
// 
// For background sessions, the default value is `nil`.
// 
// For ephemeral sessions, the default value is a private cache object that
// stores data in memory only, and is destroyed when you invalidate the
// session.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/urlCache
func (u URLSessionConfiguration) URLCache() INSURLCache {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("URLCache"))
	return NSURLCacheFromID(objc.ID(rv))
}
func (u URLSessionConfiguration) SetURLCache(value INSURLCache) {
	objc.Send[struct{}](u.ID, objc.Sel("setURLCache:"), value)
}
// A predefined constant that determines when to return a response from the
// cache.
//
// # Discussion
// 
// This property determines the request caching policy used by tasks within
// sessions based on this configuration.
// 
// Set this property to one of the constants defined in
// [NSURLRequest.CachePolicy] to specify whether the cache policy should
// depend on expiration dates and age, whether the cache should be disabled
// entirely, and whether the server should be contacted to determine if the
// content has changed since it was last requested.
// 
// The default value is [URLRequestUseProtocolCachePolicy].
//
// [NSURLRequest.CachePolicy]: https://developer.apple.com/documentation/Foundation/NSURLRequest/CachePolicy-swift.enum
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/requestCachePolicy
func (u URLSessionConfiguration) RequestCachePolicy() NSURLRequestCachePolicy {
	rv := objc.Send[NSURLRequestCachePolicy](u.ID, objc.Sel("requestCachePolicy"))
	return NSURLRequestCachePolicy(rv)
}
func (u URLSessionConfiguration) SetRequestCachePolicy(value NSURLRequestCachePolicy) {
	objc.Send[struct{}](u.ID, objc.Sel("setRequestCachePolicy:"), value)
}
// A Boolean value that indicates whether the app should be resumed or
// launched in the background when transfers finish.
//
// # Discussion
// 
// For configuration objects created using the
// [BackgroundSessionConfigurationWithIdentifier] method, you can use this
// property to control the launching behavior for an iOS app. This property is
// ignored for configuration objects created using other methods.
// 
// The default value of this property is [true]. When the value of this
// property is [true], the system automatically wakes up or launches the iOS
// app in the background when the session’s tasks finish or require
// authentication. At that time, the system calls the app delegate’s
// [application(_:handleEventsForBackgroundURLSession:completionHandler:)]
// method, providing it with the identifier of the session that needs
// attention. If your app had to be relaunched, you can use that identifier to
// create a new configuration and session object capable of servicing the
// tasks.
//
// [application(_:handleEventsForBackgroundURLSession:completionHandler:)]: https://developer.apple.com/documentation/UIKit/UIApplicationDelegate/application(_:handleEventsForBackgroundURLSession:completionHandler:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/sessionSendsLaunchEvents
func (u URLSessionConfiguration) SessionSendsLaunchEvents() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("sessionSendsLaunchEvents"))
	return rv
}
func (u URLSessionConfiguration) SetSessionSendsLaunchEvents(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setSessionSendsLaunchEvents:"), value)
}
// A Boolean value that determines whether background tasks can be scheduled
// at the discretion of the system for optimal performance.
//
// # Discussion
// 
// For configuration objects created using the
// [BackgroundSessionConfigurationWithIdentifier] method, use this property to
// give the system control over when transfers should occur. This property is
// ignored for configuration objects created using other methods.
// 
// When transferring large amounts of data, you are encouraged to set the
// value of this property to [true]. Doing so lets the system schedule those
// transfers at times that are more optimal for the device. For example, the
// system might delay transferring large files until the device is plugged in
// and connected to the network via Wi-Fi. The default value of this property
// is [false].
// 
// The session object applies the value of this property only to transfers
// that your app starts while it is in the foreground. For transfers started
// while your app is in the background, the system always starts transfers at
// its discretion—in other words, the system assumes this property is [true]
// and ignores any value you specified.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/isDiscretionary
func (u URLSessionConfiguration) Discretionary() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isDiscretionary"))
	return rv
}
func (u URLSessionConfiguration) SetDiscretionary(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setDiscretionary:"), value)
}
// A Boolean value that indicates whether TCP connections should be kept open
// when the app moves to the background.
//
// # Discussion
// 
// In addition to requesting that the connection be kept open, setting this
// value to [true] asks the system to delay reclaiming the connection when the
// app moves to the background.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/shouldUseExtendedBackgroundIdleMode
func (u URLSessionConfiguration) ShouldUseExtendedBackgroundIdleMode() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("shouldUseExtendedBackgroundIdleMode"))
	return rv
}
func (u URLSessionConfiguration) SetShouldUseExtendedBackgroundIdleMode(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setShouldUseExtendedBackgroundIdleMode:"), value)
}
// An array of extra protocol subclasses that handle requests in a session.
//
// # Discussion
// 
// The objects in this array are [Class] objects corresponding to custom
// [NSURLProtocol] subclasses that you define. URL session objects support a
// number of common networking protocols by default. Use this array to extend
// the default set of common networking protocols available for use by a
// session with one or more custom protocols that you define.
// 
// Prior to handling a request, the [NSURLSession] object searches the default
// protocols first and then checks your custom protocols until it finds one
// capable of handling the specified request. It uses the protocol whose
// [CanInitWithRequest] class method returns [true], indicating that the class
// is capable of handling the specified request.
// 
// The default value is an empty array.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/protocolClasses
func (u URLSessionConfiguration) ProtocolClasses() []objc.Class {
	rv := objc.Send[[]objc.ID](u.ID, objc.Sel("protocolClasses"))
	return objc.ConvertSlice(rv, func(id objc.ID) objc.Class {
		return objc.Class(id)
	})
}
func (u URLSessionConfiguration) SetProtocolClasses(value []objc.Class) {
	objc.Send[struct{}](u.ID, objc.Sel("setProtocolClasses:"), value)
}
// The maximum number of simultaneous connections to make to a given host.
//
// # Discussion
// 
// This property determines the maximum number of simultaneous connections
// made to each host by tasks within sessions based on this configuration.
// 
// This limit is per session, so if you use multiple sessions, your app as a
// whole may exceed this limit. Additionally, depending on your connection to
// the Internet, a session may use a lower limit than the one you specify.
// 
// The default value is `6`.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpMaximumConnectionsPerHost
func (u URLSessionConfiguration) HTTPMaximumConnectionsPerHost() int {
	rv := objc.Send[int](u.ID, objc.Sel("HTTPMaximumConnectionsPerHost"))
	return rv
}
func (u URLSessionConfiguration) SetHTTPMaximumConnectionsPerHost(value int) {
	objc.Send[struct{}](u.ID, objc.Sel("setHTTPMaximumConnectionsPerHost:"), value)
}
// A Boolean value that determines whether the session should use HTTP
// pipelining.
//
// # Discussion
// 
// This property determines whether tasks within sessions based on this
// configuration should use HTTP pipelining. You can also enable pipelining on
// a per-task basis by creating the task with an [NSURLRequest] object.
// 
// The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpShouldUsePipelining
func (u URLSessionConfiguration) HTTPShouldUsePipelining() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("HTTPShouldUsePipelining"))
	return rv
}
func (u URLSessionConfiguration) SetHTTPShouldUsePipelining(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setHTTPShouldUsePipelining:"), value)
}
// A dictionary containing information about the proxy to use within this
// session.
//
// # Discussion
// 
// This property controls which proxy tasks within sessions based on this
// configuration use when connecting to remote hosts.
// 
// Prefer using [proxyConfigurations], which supports secure proxy and relay
// types.
// 
// The default value is [NULL], which means that tasks use the default system
// settings.
// 
// See `Global Proxy Configuration` for more information about these
// dictionaries.
//
// [proxyConfigurations]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/proxyConfigurations
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/connectionProxyDictionary
func (u URLSessionConfiguration) ConnectionProxyDictionary() INSDictionary {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("connectionProxyDictionary"))
	return NSDictionaryFromID(objc.ID(rv))
}
func (u URLSessionConfiguration) SetConnectionProxyDictionary(value INSDictionary) {
	objc.Send[struct{}](u.ID, objc.Sel("setConnectionProxyDictionary:"), value)
}
// A Boolean value that indicates whether the session should wait for
// connectivity to become available, or fail immediately.
//
// # Discussion
// 
// Connectivity might be temporarily unavailable for several reasons. For
// example, a device might only have a cellular connection when
// [AllowsCellularAccess] is set to `false`, or the device might require a VPN
// connection but none is available. If the value of this property is `true`
// and sufficient connectivity is unavailable, the session calls the
// [URLSessionTaskIsWaitingForConnectivity] method of
// [NSURLSessionTaskDelegate] and waits for connectivity. When connectivity
// becomes available, the task begins its work and ultimately calls the
// delegate or completion handler as usual.
// 
// If the value of the property is `false` and connectivity is unavailable,
// the connection fails immediately with an error, such as
// [NSURLErrorNotConnectedToInternet].
// 
// This property is relevant only during the establishment of a connection. If
// a connection is established and then drops, the completion handler or
// delegate receives an error, such as [NSURLErrorNetworkConnectionLost]. For
// help dealing with dropped connections, see [Handling “The network
// connection was lost” Errors].
// 
// This property is ignored by background sessions, which always wait for
// connectivity.
//
// [Handling “The network connection was lost” Errors]: https://developer.apple.com/library/archive/qa/qa1941/_index.html#//apple_ref/doc/uid/DTS40017602
// [NSURLErrorNetworkConnectionLost]: https://developer.apple.com/documentation/Foundation/NSURLErrorNetworkConnectionLost-swift.var
// [NSURLErrorNotConnectedToInternet]: https://developer.apple.com/documentation/Foundation/NSURLErrorNotConnectedToInternet-swift.var
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/waitsForConnectivity
func (u URLSessionConfiguration) WaitsForConnectivity() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("waitsForConnectivity"))
	return rv
}
func (u URLSessionConfiguration) SetWaitsForConnectivity(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setWaitsForConnectivity:"), value)
}
// A Boolean value that indicates whether connections may use the network when
// the user has specified Low Data Mode.
//
// # Discussion
// 
// In iOS 13 and later, users can set their device to use Low Data Mode as one
// of the Cellular Data Options in the Settings app. Users can turn on Low
// Data Mode to reduce your app’s network data usage. This property controls
// a URL session’s behavior when the user turns on Low Data Mode. If there
// are no nonconstrained network interfaces available and the session’s
// [AllowsConstrainedNetworkAccess] property is [false], any task created from
// the session fails. In this case, the error provided when the task fails has
// a [networkUnavailableReason] property whose value is
// [URLErrorNetworkUnavailableReasonConstrained].
// 
// Limit your app’s of use of constrained network access to user-initiated
// tasks, and put off discretionary tasks until a nonconstrained interface
// becomes available. To do this, set [AllowsConstrainedNetworkAccess] (and
// [AllowsExpensiveNetworkAccess]) to [false] and [WaitsForConnectivity] to
// [true]. This way, your [NSURLSessionTask] waits for a suitable interface to
// become available before sending or receiving data.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [networkUnavailableReason]: https://developer.apple.com/documentation/Foundation/URLError/networkUnavailableReason-swift.property
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsConstrainedNetworkAccess
func (u URLSessionConfiguration) AllowsConstrainedNetworkAccess() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("allowsConstrainedNetworkAccess"))
	return rv
}
func (u URLSessionConfiguration) SetAllowsConstrainedNetworkAccess(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setAllowsConstrainedNetworkAccess:"), value)
}
// A Boolean value that indicates whether connections may use a network
// interface that the system considers expensive.
//
// # Discussion
// 
// The system determines what constitutes “expensive” based on the nature
// of the network interface and other factors. iOS 13 considers most cellular
// networks and personal hotspots expensive. If there are no nonexpensive
// network interfaces available and the session’s
// [AllowsExpensiveNetworkAccess] property is [false], any task created from
// the session fails. In this case, the error provided when the task fails has
// a [networkUnavailableReason] property whose value is
// [URLErrorNetworkUnavailableReasonExpensive].
// 
// You can limit your app’s of use of expensive network access to
// user-initiated tasks, and put off discretionary tasks until a nonexpensive
// interface becomes available. To do this, set [AllowsExpensiveNetworkAccess]
// (and [AllowsConstrainedNetworkAccess]) to [false] and
// [WaitsForConnectivity] to [true]. This way, your [NSURLSessionTask] waits
// for a suitable interface to become available before sending or receiving
// data.
// 
// To test the behavior of this property, you can override the device’s
// current values for cellular and Wi-Fi cost in Settings > Developer >
// Network Override.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [networkUnavailableReason]: https://developer.apple.com/documentation/Foundation/URLError/networkUnavailableReason-swift.property
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsExpensiveNetworkAccess
func (u URLSessionConfiguration) AllowsExpensiveNetworkAccess() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("allowsExpensiveNetworkAccess"))
	return rv
}
func (u URLSessionConfiguration) SetAllowsExpensiveNetworkAccess(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setAllowsExpensiveNetworkAccess:"), value)
}
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsUltraConstrainedNetworkAccess
func (u URLSessionConfiguration) AllowsUltraConstrainedNetworkAccess() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("allowsUltraConstrainedNetworkAccess"))
	return rv
}
func (u URLSessionConfiguration) SetAllowsUltraConstrainedNetworkAccess(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setAllowsUltraConstrainedNetworkAccess:"), value)
}
// A copy of the configuration object for this session.
//
// See: https://developer.apple.com/documentation/foundation/urlsession/configuration
func (u URLSessionConfiguration) Configuration() INSURLSessionConfiguration {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("configuration"))
	return NSURLSessionConfigurationFromID(objc.ID(rv))
}
func (u URLSessionConfiguration) SetConfiguration(value INSURLSessionConfiguration) {
	objc.Send[struct{}](u.ID, objc.Sel("setConfiguration:"), value)
}
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/enablesEarlyData
func (u URLSessionConfiguration) EnablesEarlyData() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("enablesEarlyData"))
	return rv
}
func (u URLSessionConfiguration) SetEnablesEarlyData(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setEnablesEarlyData:"), value)
}
// An array of proxy configuration objects containing information about the
// proxies to use within this session.
//
// # Discussion
// 
// This property controls which proxy tasks to use within sessions based on
// this configuration when connecting to remote hosts.
// 
// The default value is the empty array, which means that tasks use the
// default system settings.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLSessionConfiguration/proxyConfigurations
func (u URLSessionConfiguration) ProxyConfigurations() []objectivec.Object {
	rv := objc.Send[[]objc.ID](u.ID, objc.Sel("proxyConfigurations"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.Object {
		return objectivec.ObjectFromID(id)
	})
}
func (u URLSessionConfiguration) SetProxyConfigurations(value []objectivec.Object) {
	objc.Send[struct{}](u.ID, objc.Sel("setProxyConfigurations:"), objectivec.IObjectSliceToNSArray(value))
}

// A default session configuration object.
//
// # Discussion
// 
// The default session configuration uses a persistent disk-based cache
// (except when the result is downloaded to a file) and stores credentials in
// the user’s keychain. It also stores cookies (by default) in the same
// shared cookie store as the [NSURLConnection] and [NSURLDownload] classes.
// 
// Modifying the returned session configuration object does affect any
// configuration objects returned by future calls to this method, and does not
// change the default behavior for existing sessions. It is therefore always
// safe to use the returned object as a starting point for additional
// customization.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/default
func (_URLSessionConfigurationClass URLSessionConfigurationClass) DefaultSessionConfiguration() URLSessionConfiguration {
	rv := objc.Send[objc.ID](objc.ID(_URLSessionConfigurationClass.class), objc.Sel("defaultSessionConfiguration"))
	return NSURLSessionConfigurationFromID(objc.ID(rv))
}
// A session configuration that uses no persistent storage for caches,
// cookies, or credentials.
//
// # Discussion
// 
// An ephemeral session configuration object is similar to a default session
// configuration (see [DefaultSessionConfiguration]), except that the
// corresponding session object doesn’t store caches, credential stores, or
// any session-related data to disk. Instead, session-related data is stored
// in RAM. The only time an ephemeral session writes data to disk is when you
// tell it to write the contents of a URL to a file.
// 
// # Privacy and performance considerations
// 
// The main advantage to using ephemeral sessions is privacy. By not writing
// potentially sensitive data to disk, you make it less likely that the data
// will be intercepted and used later. For this reason, ephemeral sessions are
// ideal for private browsing modes in web browsers and other similar
// situations.
// 
// Because an ephemeral session doesn’t write cached data to disk, the size
// of the cache is limited by available RAM. This limitation means that
// previously fetched resources are less likely to be in the cache (and are
// guaranteed to not be there if the user quits and relaunches your app). This
// behavior may reduce perceived performance, depending on your app.
// 
// When your app invalidates the session, all ephemeral session data is purged
// automatically. Additionally, in iOS, the in-memory cache isn’t purged
// automatically when your app is suspended but may be purged when your app is
// terminated or when the system experiences memory pressure.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/ephemeral
func (_URLSessionConfigurationClass URLSessionConfigurationClass) EphemeralSessionConfiguration() URLSessionConfiguration {
	rv := objc.Send[objc.ID](objc.ID(_URLSessionConfigurationClass.class), objc.Sel("ephemeralSessionConfiguration"))
	return NSURLSessionConfigurationFromID(objc.ID(rv))
}

			// Protocol methods for NSCopying
			

