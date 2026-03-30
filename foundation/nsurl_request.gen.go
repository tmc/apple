// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSURLRequest] class.
var (
	_NSURLRequestClass     NSURLRequestClass
	_NSURLRequestClassOnce sync.Once
)

func getNSURLRequestClass() NSURLRequestClass {
	_NSURLRequestClassOnce.Do(func() {
		_NSURLRequestClass = NSURLRequestClass{class: objc.GetClass("NSURLRequest")}
	})
	return _NSURLRequestClass
}

// GetNSURLRequestClass returns the class object for NSURLRequest.
func GetNSURLRequestClass() NSURLRequestClass {
	return getNSURLRequestClass()
}

type NSURLRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSURLRequestClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSURLRequestClass) Alloc() NSURLRequest {
	rv := objc.Send[NSURLRequest](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A URL load request that is independent of protocol or URL scheme.
//
// # Overview
//
// Use this type in Swift when you need reference semantics or other
// Foundation-specific behavior.
//
// [NSURLRequest] encapsulates two essential properties of a load request: the
// URL to load and the policies used to load it. In addition, for HTTP and
// HTTPS requests, [URLRequest] includes the HTTP method ([GET], [POST], and
// so on) and the HTTP headers. Finally, custom protocols can support custom
// properties as explained in [NSURLRequest].
//
// [NSURLRequest] only represents information about the request. Use other
// classes, such as [NSURLSession], to send the request to a server. See
// [Fetching website data into memory] and [Uploading data to a website] for
// an introduction to these techniques.
//
// The mutable subclass of [NSURLRequest] is [NSMutableURLRequest].
//
// # Reserved HTTP headers
//
// The URL Loading System handles various aspects of the HTTP protocol for you
// (HTTP 1.1 persistent connections, proxies, authentication, and so on). As
// part of this support, the URL Loading System takes responsibility for
// certain HTTP headers:
//
// - `Content-Length` - [Authorization] - [Connection] - [Host] -
// `Proxy-Authenticate` - `Proxy-Authorization` - `WWW-Authenticate`
//
// If you set a value for one of these reserved headers, the system may ignore
// the value you set, or overwrite it with its own value, or simply not send
// it. Moreover, the exact behavior may change over time. To avoid confusing
// problems like this, do not set these headers directly.
//
// The URL Loading System sets the `Content-Length` header based on whether
// the request body has a known length:
//
// - If so, it uses the identity transfer encoding and sets the
// `Content-Length` header to that known length. You see this when you set the
// request body to a data object. - If not, it uses the chunked transfer
// encoding and omits the `Content-Length` header. You see this when you set
// the request body to a stream.
//
// # Custom protocol properties
//
// If you implement a custom URL protocol by subclassing [NSURLProtocol], and
// it needs protocol-specific properties, extend [NSURLRequest] with accessor
// methods for those custom properties. In your accessor methods, call
// [PropertyForKeyInRequest] and [SetPropertyForKeyInRequest] to associate
// property values with the request.
//
// # Creating requests
//
//   - [NSURLRequest.InitWithURL]: Creates a URL request for a specified URL.
//   - [NSURLRequest.InitWithURLCachePolicyTimeoutInterval]: Creates a URL request with the specified URL, cache policy, and timeout values.
//
// # Working with a cache policy
//
//   - [NSURLRequest.CachePolicy]: The request’s cache policy.
//
// # Accessing request components
//
//   - [NSURLRequest.HTTPMethod]: The HTTP request method.
//   - [NSURLRequest.URL]: The URL being requested.
//   - [NSURLRequest.HTTPBody]: The request body.
//   - [NSURLRequest.HTTPBodyStream]: The request body as an input stream.
//   - [NSURLRequest.MainDocumentURL]: The main document URL associated with the request.
//
// # Getting header fields
//
//   - [NSURLRequest.AllHTTPHeaderFields]: A dictionary containing all of the HTTP header fields for a request.
//   - [NSURLRequest.ValueForHTTPHeaderField]: Returns the value of the specified HTTP header field.
//
// # Controlling request behavior
//
//   - [NSURLRequest.TimeoutInterval]: The request’s timeout interval, in seconds.
//   - [NSURLRequest.HTTPShouldHandleCookies]: A Boolean value that indicates whether the default cookie handling will be used for this request.
//   - [NSURLRequest.HTTPShouldUsePipelining]: A Boolean value that indicates whether the request should continue transmitting data before receiving a response from an earlier transmission.
//   - [NSURLRequest.AllowsCellularAccess]: A Boolean value that indicates whether the request is allowed to use the cellular radio (if present).
//
// # Supporting limited modes
//
//   - [NSURLRequest.AllowsConstrainedNetworkAccess]: A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//   - [NSURLRequest.AllowsExpensiveNetworkAccess]: A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// # Accessing the service type
//
//   - [NSURLRequest.NetworkServiceType]: The network service type of the request.
//
// # Indicating the source of the request
//
//   - [NSURLRequest.Attribution]: The entity that initiates the network request.
//
// # Instance Properties
//
//   - [NSURLRequest.AllowsPersistentDNS]
//   - [NSURLRequest.AllowsUltraConstrainedNetworkAccess]
//   - [NSURLRequest.AssumesHTTP3Capable]
//   - [NSURLRequest.CookiePartitionIdentifier]
//   - [NSURLRequest.RequiresDNSSECValidation]
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest
//
// [Fetching website data into memory]: https://developer.apple.com/documentation/Foundation/fetching-website-data-into-memory
// [URLRequest]: https://developer.apple.com/documentation/Foundation/URLRequest
// [Uploading data to a website]: https://developer.apple.com/documentation/Foundation/uploading-data-to-a-website
type NSURLRequest struct {
	objectivec.Object
}

// NSURLRequestFromID constructs a [NSURLRequest] from an objc.ID.
//
// A URL load request that is independent of protocol or URL scheme.
func NSURLRequestFromID(id objc.ID) NSURLRequest {
	return NSURLRequest{objectivec.Object{ID: id}}
}

// NOTE: NSURLRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSURLRequest] class.
//
// # Creating requests
//
//   - [INSURLRequest.InitWithURL]: Creates a URL request for a specified URL.
//   - [INSURLRequest.InitWithURLCachePolicyTimeoutInterval]: Creates a URL request with the specified URL, cache policy, and timeout values.
//
// # Working with a cache policy
//
//   - [INSURLRequest.CachePolicy]: The request’s cache policy.
//
// # Accessing request components
//
//   - [INSURLRequest.HTTPMethod]: The HTTP request method.
//   - [INSURLRequest.URL]: The URL being requested.
//   - [INSURLRequest.HTTPBody]: The request body.
//   - [INSURLRequest.HTTPBodyStream]: The request body as an input stream.
//   - [INSURLRequest.MainDocumentURL]: The main document URL associated with the request.
//
// # Getting header fields
//
//   - [INSURLRequest.AllHTTPHeaderFields]: A dictionary containing all of the HTTP header fields for a request.
//   - [INSURLRequest.ValueForHTTPHeaderField]: Returns the value of the specified HTTP header field.
//
// # Controlling request behavior
//
//   - [INSURLRequest.TimeoutInterval]: The request’s timeout interval, in seconds.
//   - [INSURLRequest.HTTPShouldHandleCookies]: A Boolean value that indicates whether the default cookie handling will be used for this request.
//   - [INSURLRequest.HTTPShouldUsePipelining]: A Boolean value that indicates whether the request should continue transmitting data before receiving a response from an earlier transmission.
//   - [INSURLRequest.AllowsCellularAccess]: A Boolean value that indicates whether the request is allowed to use the cellular radio (if present).
//
// # Supporting limited modes
//
//   - [INSURLRequest.AllowsConstrainedNetworkAccess]: A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//   - [INSURLRequest.AllowsExpensiveNetworkAccess]: A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// # Accessing the service type
//
//   - [INSURLRequest.NetworkServiceType]: The network service type of the request.
//
// # Indicating the source of the request
//
//   - [INSURLRequest.Attribution]: The entity that initiates the network request.
//
// # Instance Properties
//
//   - [INSURLRequest.AllowsPersistentDNS]
//   - [INSURLRequest.AllowsUltraConstrainedNetworkAccess]
//   - [INSURLRequest.AssumesHTTP3Capable]
//   - [INSURLRequest.CookiePartitionIdentifier]
//   - [INSURLRequest.RequiresDNSSECValidation]
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest
type INSURLRequest interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSMutableCopying
	NSSecureCoding

	// Topic: Creating requests

	// Creates a URL request for a specified URL.
	InitWithURL(URL INSURL) NSURLRequest
	// Creates a URL request with the specified URL, cache policy, and timeout values.
	InitWithURLCachePolicyTimeoutInterval(URL INSURL, cachePolicy NSURLRequestCachePolicy, timeoutInterval float64) NSURLRequest

	// Topic: Working with a cache policy

	// The request’s cache policy.
	CachePolicy() NSURLRequestCachePolicy

	// Topic: Accessing request components

	// The HTTP request method.
	HTTPMethod() string
	// The URL being requested.
	URL() INSURL
	// The request body.
	HTTPBody() INSData
	// The request body as an input stream.
	HTTPBodyStream() INSInputStream
	// The main document URL associated with the request.
	MainDocumentURL() INSURL

	// Topic: Getting header fields

	// A dictionary containing all of the HTTP header fields for a request.
	AllHTTPHeaderFields() INSDictionary
	// Returns the value of the specified HTTP header field.
	ValueForHTTPHeaderField(field string) string

	// Topic: Controlling request behavior

	// The request’s timeout interval, in seconds.
	TimeoutInterval() float64
	// A Boolean value that indicates whether the default cookie handling will be used for this request.
	HTTPShouldHandleCookies() bool
	// A Boolean value that indicates whether the request should continue transmitting data before receiving a response from an earlier transmission.
	HTTPShouldUsePipelining() bool
	// A Boolean value that indicates whether the request is allowed to use the cellular radio (if present).
	AllowsCellularAccess() bool

	// Topic: Supporting limited modes

	// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
	AllowsConstrainedNetworkAccess() bool
	// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
	AllowsExpensiveNetworkAccess() bool

	// Topic: Accessing the service type

	// The network service type of the request.
	NetworkServiceType() NSURLRequestNetworkServiceType

	// Topic: Indicating the source of the request

	// The entity that initiates the network request.
	Attribution() NSURLRequestAttribution

	// Topic: Instance Properties

	AllowsPersistentDNS() bool
	AllowsUltraConstrainedNetworkAccess() bool
	AssumesHTTP3Capable() bool
	CookiePartitionIdentifier() string
	RequiresDNSSECValidation() bool
}

// Init initializes the instance.
func (u NSURLRequest) Init() NSURLRequest {
	rv := objc.Send[NSURLRequest](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u NSURLRequest) Autorelease() NSURLRequest {
	rv := objc.Send[NSURLRequest](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSURLRequest creates a new NSURLRequest instance.
func NewNSURLRequest() NSURLRequest {
	class := getNSURLRequestClass()
	rv := objc.Send[NSURLRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewURLRequestWithCoder(coder INSCoder) NSURLRequest {
	instance := getNSURLRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSURLRequestFromID(rv)
}

// Creates a URL request for a specified URL.
//
// URL: The URL for the request.
//
// # Return Value
//
// The initialized URL request.
//
// # Discussion
//
// The request is created with the with the default cache policy
// ([NSURLRequestUseProtocolCachePolicy]), and the default timeout interval
// (60 seconds).
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/init(url:)
func NewURLRequestWithURL(URL INSURL) NSURLRequest {
	instance := getNSURLRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), URL)
	return NSURLRequestFromID(rv)
}

// Creates a URL request with the specified URL, cache policy, and timeout
// values.
//
// URL: The URL for the request.
//
// cachePolicy: The cache policy for the request.
//
// timeoutInterval: The timeout interval for the request, in seconds.
//
// # Return Value
//
// The initialized URL request.
//
// # Discussion
//
// This is the designated initializer for [NSURLRequest].
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/init(url:cachePolicy:timeoutInterval:)
func NewURLRequestWithURLCachePolicyTimeoutInterval(URL INSURL, cachePolicy NSURLRequestCachePolicy, timeoutInterval float64) NSURLRequest {
	instance := getNSURLRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:cachePolicy:timeoutInterval:"), URL, cachePolicy, timeoutInterval)
	return NSURLRequestFromID(rv)
}

// Creates a URL request for a specified URL.
//
// URL: The URL for the request.
//
// # Return Value
//
// The initialized URL request.
//
// # Discussion
//
// The request is created with the with the default cache policy
// ([NSURLRequestUseProtocolCachePolicy]), and the default timeout interval
// (60 seconds).
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/init(url:)
func (u NSURLRequest) InitWithURL(URL INSURL) NSURLRequest {
	rv := objc.Send[NSURLRequest](u.ID, objc.Sel("initWithURL:"), URL)
	return rv
}

// Creates a URL request with the specified URL, cache policy, and timeout
// values.
//
// URL: The URL for the request.
//
// cachePolicy: The cache policy for the request.
//
// timeoutInterval: The timeout interval for the request, in seconds.
//
// # Return Value
//
// The initialized URL request.
//
// # Discussion
//
// This is the designated initializer for [NSURLRequest].
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/init(url:cachePolicy:timeoutInterval:)
func (u NSURLRequest) InitWithURLCachePolicyTimeoutInterval(URL INSURL, cachePolicy NSURLRequestCachePolicy, timeoutInterval float64) NSURLRequest {
	rv := objc.Send[NSURLRequest](u.ID, objc.Sel("initWithURL:cachePolicy:timeoutInterval:"), URL, cachePolicy, timeoutInterval)
	return rv
}

// Returns the value of the specified HTTP header field.
//
// field: The name of the header field whose value is to be returned. In keeping with
// the HTTP RFC, HTTP header field names are case-insensitive.
//
// # Return Value
//
// The value associated with the header field `field`, or `nil` if there is no
// corresponding header field.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/value(forHTTPHeaderField:)
func (u NSURLRequest) ValueForHTTPHeaderField(field string) string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("valueForHTTPHeaderField:"), objc.String(field))
	return NSStringFromID(rv).String()
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (u NSURLRequest) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (u NSURLRequest) InitWithCoder(coder INSCoder) NSURLRequest {
	rv := objc.Send[NSURLRequest](u.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Creates and returns a URL request for a specified URL.
//
// URL: The URL for the new request.
//
// # Return Value
//
// The newly created URL request.
//
// # Discussion
//
// The request is created with the with the default cache policy
// ([NSURLRequestUseProtocolCachePolicy]), and the default timeout interval
// (60 seconds).
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/requestWithURL:
func (_NSURLRequestClass NSURLRequestClass) RequestWithURL(URL INSURL) NSURLRequest {
	rv := objc.Send[objc.ID](objc.ID(_NSURLRequestClass.class), objc.Sel("requestWithURL:"), URL)
	return NSURLRequestFromID(rv)
}

// Creates and returns an initialized URL request with specified URL, cache
// policy, and timeout values.
//
// URL: The URL for the new request.
//
// cachePolicy: The cache policy for the new request.
//
// timeoutInterval: The timeout interval for the new request, in seconds.
//
// # Return Value
//
// The initialized URL request.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/requestWithURL:cachePolicy:timeoutInterval:
func (_NSURLRequestClass NSURLRequestClass) RequestWithURLCachePolicyTimeoutInterval(URL INSURL, cachePolicy NSURLRequestCachePolicy, timeoutInterval float64) NSURLRequest {
	rv := objc.Send[objc.ID](objc.ID(_NSURLRequestClass.class), objc.Sel("requestWithURL:cachePolicy:timeoutInterval:"), URL, cachePolicy, timeoutInterval)
	return NSURLRequestFromID(rv)
}

// The request’s cache policy.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/cachePolicy-swift.property
func (u NSURLRequest) CachePolicy() NSURLRequestCachePolicy {
	rv := objc.Send[NSURLRequestCachePolicy](u.ID, objc.Sel("cachePolicy"))
	return NSURLRequestCachePolicy(rv)
}

// The HTTP request method.
//
// # Discussion
//
// The default HTTP method is “GET”.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/httpMethod
func (u NSURLRequest) HTTPMethod() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("HTTPMethod"))
	return NSStringFromID(rv).String()
}

// The URL being requested.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/url
func (u NSURLRequest) URL() INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("URL"))
	return NSURLFromID(objc.ID(rv))
}

// The request body.
//
// # Discussion
//
// This data is sent as the message body of a request, as in an HTTP [POST]
// request.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/httpBody
func (u NSURLRequest) HTTPBody() INSData {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("HTTPBody"))
	return NSDataFromID(objc.ID(rv))
}

// The request body as an input stream.
//
// # Discussion
//
// `nil` if the body stream has not been set. The returned stream is for
// examination only—it is not safe to manipulate the stream in any way.
//
// The receiver will have either an HTTP body or an HTTP body stream, only one
// may be set for a request. A HTTP body stream is preserved when copying an
// [NSURLRequest] object, but is lost when a request is archived using the
// [NSCoding] protocol.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/httpBodyStream
func (u NSURLRequest) HTTPBodyStream() INSInputStream {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("HTTPBodyStream"))
	return NSInputStreamFromID(objc.ID(rv))
}

// The main document URL associated with the request.
//
// # Discussion
//
// This URL is used for the cookie “same domain as main document” policy.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/mainDocumentURL
func (u NSURLRequest) MainDocumentURL() INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("mainDocumentURL"))
	return NSURLFromID(objc.ID(rv))
}

// A dictionary containing all of the HTTP header fields for a request.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/allHTTPHeaderFields
func (u NSURLRequest) AllHTTPHeaderFields() INSDictionary {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("allHTTPHeaderFields"))
	return NSDictionaryFromID(objc.ID(rv))
}

// The request’s timeout interval, in seconds.
//
// # Discussion
//
// If during a connection attempt the request remains idle for longer than the
// timeout interval, the request is considered to have timed out.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/timeoutInterval
func (u NSURLRequest) TimeoutInterval() float64 {
	rv := objc.Send[NSTimeInterval](u.ID, objc.Sel("timeoutInterval"))
	return float64(rv)
}

// A Boolean value that indicates whether the default cookie handling will be
// used for this request.
//
// # Discussion
//
// true if the default cookie handling will be used for this request, false
// otherwise. The default is true.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/httpShouldHandleCookies
func (u NSURLRequest) HTTPShouldHandleCookies() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("HTTPShouldHandleCookies"))
	return rv
}

// A Boolean value that indicates whether the request should continue
// transmitting data before receiving a response from an earlier transmission.
//
// # Discussion
//
// true if the request should continue transmitting data; otherwise, false.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/httpShouldUsePipelining
func (u NSURLRequest) HTTPShouldUsePipelining() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("HTTPShouldUsePipelining"))
	return rv
}

// A Boolean value that indicates whether the request is allowed to use the
// cellular radio (if present).
//
// # Discussion
//
// true if the cellular radio can be used; false otherwise.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/allowsCellularAccess
func (u NSURLRequest) AllowsCellularAccess() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("allowsCellularAccess"))
	return rv
}

// A Boolean value that indicates whether connections may use the network when
// the user has specified Low Data Mode.
//
// # Discussion
//
// In iOS 13 and later, users can set their device to use Low Data Mode as one
// of the Cellular Data Options in the Settings app. Users can turn on Low
// Data Mode to reduce your app’s network data usage. This property controls
// the request’s behavior when the user turns on Low Data Mode. If there are
// no nonconstrained network interfaces available and the request’s
// [AllowsConstrainedNetworkAccess] property is false, any connection created
// from the request fails. In this case, the error provided when the
// connection fails has a [networkUnavailableReason] property whose value is
// [NSURLErrorNetworkUnavailableReasonConstrained].
//
// Limit your app’s of use of constrained network access to user-initiated
// tasks, and put off discretionary tasks until a nonconstrained interface
// becomes available.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/allowsConstrainedNetworkAccess
//
// [networkUnavailableReason]: https://developer.apple.com/documentation/Foundation/URLError/networkUnavailableReason-swift.property
func (u NSURLRequest) AllowsConstrainedNetworkAccess() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("allowsConstrainedNetworkAccess"))
	return rv
}

// A Boolean value that indicates whether connections may use a network
// interface that the system considers expensive.
//
// # Discussion
//
// The system determines what constitutes “expensive” based on the nature
// of the network interface and other factors. iOS 13 considers most cellular
// networks and personal hotspots expensive. If there are no nonexpensive
// network interfaces available and the request’s
// [AllowsExpensiveNetworkAccess] property is false, any task created from the
// request fails. In this case, the error provided when the task fails has a
// [networkUnavailableReason] property whose value is
// [NSURLErrorNetworkUnavailableReasonExpensive].
//
// Limit your app’s of use of expensive network access to user-initiated
// tasks, and put off discretionary tasks until a nonexpensive interface
// becomes available.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/allowsExpensiveNetworkAccess
//
// [networkUnavailableReason]: https://developer.apple.com/documentation/Foundation/URLError/networkUnavailableReason-swift.property
func (u NSURLRequest) AllowsExpensiveNetworkAccess() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("allowsExpensiveNetworkAccess"))
	return rv
}

// The network service type of the request.
//
// # Discussion
//
// The network service type provides a hint to the operating system about what
// the underlying traffic is used for. This hint enhances the system’s
// ability to prioritize traffic, determine how quickly it needs to wake up
// the cellular or Wi-Fi radio, and so on. By providing accurate information,
// you improve the ability of the system to optimally balance battery life,
// performance, and other considerations.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/networkServiceType-swift.property
func (u NSURLRequest) NetworkServiceType() NSURLRequestNetworkServiceType {
	rv := objc.Send[NSURLRequestNetworkServiceType](u.ID, objc.Sel("networkServiceType"))
	return NSURLRequestNetworkServiceType(rv)
}

// The entity that initiates the network request.
//
// # Discussion
//
// If you don’t set this value, the system assumes a value of
// [NSURLRequestAttributionDeveloper]. Use this default value for any network
// request that your app makes that isn’t explicitly from the user. This
// includes requests that you make to your own server, even when you load user
// data. It also includes links that the user selects, but that you modify in
// any way — including by adding HTTP headers — before loading the
// content.
//
// Set this value to [NSURLRequest.Attribution.user] only for requests that
// the user explicitly makes, like when the user enters a URL or taps or
// clicks a URL that they can read, and only if your app loads and displays
// the data without altering the request.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/attribution-swift.property
//
// [NSURLRequest.Attribution.user]: https://developer.apple.com/documentation/Foundation/NSURLRequest/Attribution-swift.enum/user
func (u NSURLRequest) Attribution() NSURLRequestAttribution {
	rv := objc.Send[NSURLRequestAttribution](u.ID, objc.Sel("attribution"))
	return NSURLRequestAttribution(rv)
}

// # Discussion
//
// Allows storing and usage of DNS answers, potentially beyond TTL expiry, in
// a persistent per-process cache. This should only be set for hostnames whose
// resolutions are not expected to change across networks.
//
// YES, if the DNS lookup for this request is allowed to use a persistent
// per-process cache, NO otherwise. Defaults to NO.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/allowsPersistentDNS
func (u NSURLRequest) AllowsPersistentDNS() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("allowsPersistentDNS"))
	return rv
}

// # Return Value
//
// YES if the receiver is allowed to use an interface marked as ultra
// constrained to satisfy the request, NO otherwise.
//
// # Discussion
//
// Returns whether a connection created with this request is allowed to use
// network interfaces which have been marked as ultra constrained.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/allowsUltraConstrainedNetworkAccess
func (u NSURLRequest) AllowsUltraConstrainedNetworkAccess() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("allowsUltraConstrainedNetworkAccess"))
	return rv
}

// # Return Value
//
// YES if server endpoint is known to support HTTP/3. Defaults to NO. The
// default may be YES in a future OS update.
//
// # Discussion
//
// Returns whether we assume that server supports HTTP/3. Enables QUIC racing
// without HTTP/3 service discovery.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/assumesHTTP3Capable
func (u NSURLRequest) AssumesHTTP3Capable() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("assumesHTTP3Capable"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/cookiePartitionIdentifier
func (u NSURLRequest) CookiePartitionIdentifier() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("cookiePartitionIdentifier"))
	return NSStringFromID(rv).String()
}

// # Discussion
//
// Sets whether a request is required to do DNSSEC validation during DNS
// lookup.
//
// YES, if the DNS lookup for this request should require DNSSEC validation,
// No otherwise. Defaults to NO.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/requiresDNSSECValidation
func (u NSURLRequest) RequiresDNSSECValidation() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("requiresDNSSECValidation"))
	return rv
}

// A Boolean value indicating whether the [NSURLRequest] implements the
// [NSSecureCoding] protocol.
//
// # Return Value
//
// true to indicate that [NSURLRequest] implements the [NSSecureCoding]
// protocol.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/supportsSecureCoding
func (_NSURLRequestClass NSURLRequestClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_NSURLRequestClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// Protocol methods for NSCopying

// Protocol methods for NSMutableCopying

// Protocol methods for NSSecureCoding
