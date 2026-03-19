// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [URLSessionTaskTransactionMetrics] class.
var (
	_URLSessionTaskTransactionMetricsClass     URLSessionTaskTransactionMetricsClass
	_URLSessionTaskTransactionMetricsClassOnce sync.Once
)

func getURLSessionTaskTransactionMetricsClass() URLSessionTaskTransactionMetricsClass {
	_URLSessionTaskTransactionMetricsClassOnce.Do(func() {
		_URLSessionTaskTransactionMetricsClass = URLSessionTaskTransactionMetricsClass{class: objc.GetClass("NSURLSessionTaskTransactionMetrics")}
	})
	return _URLSessionTaskTransactionMetricsClass
}

// GetURLSessionTaskTransactionMetricsClass returns the class object for NSURLSessionTaskTransactionMetrics.
func GetURLSessionTaskTransactionMetricsClass() URLSessionTaskTransactionMetricsClass {
	return getURLSessionTaskTransactionMetricsClass()
}

type URLSessionTaskTransactionMetricsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc URLSessionTaskTransactionMetricsClass) Alloc() URLSessionTaskTransactionMetrics {
	rv := objc.Send[URLSessionTaskTransactionMetrics](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// An object that encapsualtes the performance metrics collected by the URL
// Loading System during the execution of a session task.
//
// # Overview
// 
// Each [NSURLSessionTaskTransactionMetrics] object consists of a [Request]
// and [Response] property, corresponding to the request and response of the
// corresponding task. It also contains temporal metrics, starting with
// [FetchStartDate] and ending with [ResponseEndDate], as well as other
// characteristics like [NetworkProtocolName] and [ResourceFetchType].
// 
// # Understanding temporal metrics
// 
// The figure below shows the sequence of events for a URL session task, which
// correspond to the temporal metrics captured by
// [NSURLSessionTaskTransactionMetrics].
// 
// [media-3162616]
// 
// For all metrics with a start and end date, if an aspect of the task was not
// completed, then its corresponding end date metric is `nil`. This can happen
// if name lookup begins, but the operation either times out, fails, or the
// client cancels the task before the name can be resolved. In this case, the
// [DomainLookupEndDate] property is `nil`, along with all metrics for aspects
// that occurred afterwards.
// 
// # Measuring tasks using iCloud Private Relay
// 
// iCloud Private Relay can change the timing and sequence of events for your
// tasks by sending requests through a set of privacy proxies. All tasks that
// use iCloud Private Relay set the [ProxyConnection] property in their
// transaction metrics. In this case, the [RemoteAddress] property contains
// the address of the proxy, and not the origin server.
// 
// Tasks to different hosts can reuse the same transport connection, just like
// how tasks can already share a connection when using HTTP/2. In these cases,
// a proxied task may not report any [SecureConnectionStartDate] or
// [SecureConnectionEndDate].
//
// # Accessing request and response
//
//   - [URLSessionTaskTransactionMetrics.Request]: The transaction request.
//   - [URLSessionTaskTransactionMetrics.Response]: The transaction response.
//
// # Accessing temporal metrics
//
//   - [URLSessionTaskTransactionMetrics.FetchStartDate]: The time when the task started fetching the resource, from the server or locally.
//   - [URLSessionTaskTransactionMetrics.DomainLookupStartDate]: The time immediately before the task started the name lookup for the resource.
//   - [URLSessionTaskTransactionMetrics.DomainLookupEndDate]: The time after the name lookup was completed.
//   - [URLSessionTaskTransactionMetrics.ConnectStartDate]: The time immediately before the task started establishing a TCP connection to the server.
//   - [URLSessionTaskTransactionMetrics.SecureConnectionStartDate]: The time immediately before the task started the TLS security handshake to secure the current connection.
//   - [URLSessionTaskTransactionMetrics.SecureConnectionEndDate]: The time immediately after the security handshake completed.
//   - [URLSessionTaskTransactionMetrics.ConnectEndDate]: The time immediately after the task finished establishing the connection to the server.
//   - [URLSessionTaskTransactionMetrics.RequestStartDate]: The time immediately before the task started requesting the resource, regardless of whether it is retrieved from the server or local resources.
//   - [URLSessionTaskTransactionMetrics.RequestEndDate]: The time immediately after the task finished requesting the resource, regardless of whether it was retrieved from the server or local resources.
//   - [URLSessionTaskTransactionMetrics.ResponseStartDate]: The time immediately after the task received the first byte of the response from the server or from local resources.
//   - [URLSessionTaskTransactionMetrics.ResponseEndDate]: The time immediately after the task received the last byte of the resource.
//
// # Accessing data transfer metrics
//
//   - [URLSessionTaskTransactionMetrics.CountOfRequestBodyBytesBeforeEncoding]: The size of the upload body data, file, or stream, in bytes.
//   - [URLSessionTaskTransactionMetrics.CountOfRequestBodyBytesSent]: The number of bytes transferred for the request body.
//   - [URLSessionTaskTransactionMetrics.CountOfRequestHeaderBytesSent]: The number of bytes transferred for the request header.
//   - [URLSessionTaskTransactionMetrics.CountOfResponseBodyBytesAfterDecoding]: The size of data delivered to your delegate or completion handler.
//   - [URLSessionTaskTransactionMetrics.CountOfResponseBodyBytesReceived]: The number of bytes transferred for the response body.
//   - [URLSessionTaskTransactionMetrics.CountOfResponseHeaderBytesReceived]: The number of bytes transferred for the response header.
//
// # Accessing transaction characteristics
//
//   - [URLSessionTaskTransactionMetrics.NetworkProtocolName]: The network protocol used to fetch the resource.
//   - [URLSessionTaskTransactionMetrics.RemoteAddress]: The IP address string of the remote interface for the connection.
//   - [URLSessionTaskTransactionMetrics.LocalAddress]: The IP address string of the local interface for the connection.
//   - [URLSessionTaskTransactionMetrics.NegotiatedTLSProtocolVersion]: The TLS protocol version the task negotiated with the endpoint for the connection.
//   - [URLSessionTaskTransactionMetrics.SetNegotiatedTLSProtocolVersion]
//   - [URLSessionTaskTransactionMetrics.Cellular]: A Boolean value that indicates whether the connection operates over a cellular interface.
//   - [URLSessionTaskTransactionMetrics.Expensive]: A Boolean value that indicates whether the connection operates over an expensive interface.
//   - [URLSessionTaskTransactionMetrics.Constrained]: A Boolean value that indicates whether the connection operates over an interface marked as constrained.
//   - [URLSessionTaskTransactionMetrics.ProxyConnection]: A Boolean value that indicastes whether the task used a proxy connection to fetch the resource.
//   - [URLSessionTaskTransactionMetrics.ReusedConnection]: A Boolean value that indicates whether the task used a persistent connection to fetch the resource.
//   - [URLSessionTaskTransactionMetrics.Multipath]: A Boolean value that indicates whether the connection uses a successfully negotiated multipath protocol.
//   - [URLSessionTaskTransactionMetrics.ResourceFetchType]: A value that indicates whether the resource was loaded, pushed, or retrieved from the local cache.
//   - [URLSessionTaskTransactionMetrics.DomainResolutionProtocol]
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics
type URLSessionTaskTransactionMetrics struct {
	objectivec.Object
}

// URLSessionTaskTransactionMetricsFromID constructs a [URLSessionTaskTransactionMetrics] from an objc.ID.
//
// An object that encapsualtes the performance metrics collected by the URL
// Loading System during the execution of a session task.
func URLSessionTaskTransactionMetricsFromID(id objc.ID) URLSessionTaskTransactionMetrics {
	return NSURLSessionTaskTransactionMetrics{objectivec.Object{ID: id}}
}

// NSURLSessionTaskTransactionMetricsFromID is an alias for [URLSessionTaskTransactionMetricsFromID] for cross-framework compatibility.
func NSURLSessionTaskTransactionMetricsFromID(id objc.ID) URLSessionTaskTransactionMetrics { return URLSessionTaskTransactionMetricsFromID(id) }
// NOTE: URLSessionTaskTransactionMetrics adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [URLSessionTaskTransactionMetrics] class.
//
// # Accessing request and response
//
//   - [IURLSessionTaskTransactionMetrics.Request]: The transaction request.
//   - [IURLSessionTaskTransactionMetrics.Response]: The transaction response.
//
// # Accessing temporal metrics
//
//   - [IURLSessionTaskTransactionMetrics.FetchStartDate]: The time when the task started fetching the resource, from the server or locally.
//   - [IURLSessionTaskTransactionMetrics.DomainLookupStartDate]: The time immediately before the task started the name lookup for the resource.
//   - [IURLSessionTaskTransactionMetrics.DomainLookupEndDate]: The time after the name lookup was completed.
//   - [IURLSessionTaskTransactionMetrics.ConnectStartDate]: The time immediately before the task started establishing a TCP connection to the server.
//   - [IURLSessionTaskTransactionMetrics.SecureConnectionStartDate]: The time immediately before the task started the TLS security handshake to secure the current connection.
//   - [IURLSessionTaskTransactionMetrics.SecureConnectionEndDate]: The time immediately after the security handshake completed.
//   - [IURLSessionTaskTransactionMetrics.ConnectEndDate]: The time immediately after the task finished establishing the connection to the server.
//   - [IURLSessionTaskTransactionMetrics.RequestStartDate]: The time immediately before the task started requesting the resource, regardless of whether it is retrieved from the server or local resources.
//   - [IURLSessionTaskTransactionMetrics.RequestEndDate]: The time immediately after the task finished requesting the resource, regardless of whether it was retrieved from the server or local resources.
//   - [IURLSessionTaskTransactionMetrics.ResponseStartDate]: The time immediately after the task received the first byte of the response from the server or from local resources.
//   - [IURLSessionTaskTransactionMetrics.ResponseEndDate]: The time immediately after the task received the last byte of the resource.
//
// # Accessing data transfer metrics
//
//   - [IURLSessionTaskTransactionMetrics.CountOfRequestBodyBytesBeforeEncoding]: The size of the upload body data, file, or stream, in bytes.
//   - [IURLSessionTaskTransactionMetrics.CountOfRequestBodyBytesSent]: The number of bytes transferred for the request body.
//   - [IURLSessionTaskTransactionMetrics.CountOfRequestHeaderBytesSent]: The number of bytes transferred for the request header.
//   - [IURLSessionTaskTransactionMetrics.CountOfResponseBodyBytesAfterDecoding]: The size of data delivered to your delegate or completion handler.
//   - [IURLSessionTaskTransactionMetrics.CountOfResponseBodyBytesReceived]: The number of bytes transferred for the response body.
//   - [IURLSessionTaskTransactionMetrics.CountOfResponseHeaderBytesReceived]: The number of bytes transferred for the response header.
//
// # Accessing transaction characteristics
//
//   - [IURLSessionTaskTransactionMetrics.NetworkProtocolName]: The network protocol used to fetch the resource.
//   - [IURLSessionTaskTransactionMetrics.RemoteAddress]: The IP address string of the remote interface for the connection.
//   - [IURLSessionTaskTransactionMetrics.LocalAddress]: The IP address string of the local interface for the connection.
//   - [IURLSessionTaskTransactionMetrics.NegotiatedTLSProtocolVersion]: The TLS protocol version the task negotiated with the endpoint for the connection.
//   - [IURLSessionTaskTransactionMetrics.SetNegotiatedTLSProtocolVersion]
//   - [IURLSessionTaskTransactionMetrics.Cellular]: A Boolean value that indicates whether the connection operates over a cellular interface.
//   - [IURLSessionTaskTransactionMetrics.Expensive]: A Boolean value that indicates whether the connection operates over an expensive interface.
//   - [IURLSessionTaskTransactionMetrics.Constrained]: A Boolean value that indicates whether the connection operates over an interface marked as constrained.
//   - [IURLSessionTaskTransactionMetrics.ProxyConnection]: A Boolean value that indicastes whether the task used a proxy connection to fetch the resource.
//   - [IURLSessionTaskTransactionMetrics.ReusedConnection]: A Boolean value that indicates whether the task used a persistent connection to fetch the resource.
//   - [IURLSessionTaskTransactionMetrics.Multipath]: A Boolean value that indicates whether the connection uses a successfully negotiated multipath protocol.
//   - [IURLSessionTaskTransactionMetrics.ResourceFetchType]: A value that indicates whether the resource was loaded, pushed, or retrieved from the local cache.
//   - [IURLSessionTaskTransactionMetrics.DomainResolutionProtocol]
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics
type IURLSessionTaskTransactionMetrics interface {
	objectivec.IObject

	// Topic: Accessing request and response

	// The transaction request.
	Request() INSURLRequest
	// The transaction response.
	Response() INSURLResponse

	// Topic: Accessing temporal metrics

	// The time when the task started fetching the resource, from the server or locally.
	FetchStartDate() INSDate
	// The time immediately before the task started the name lookup for the resource.
	DomainLookupStartDate() INSDate
	// The time after the name lookup was completed.
	DomainLookupEndDate() INSDate
	// The time immediately before the task started establishing a TCP connection to the server.
	ConnectStartDate() INSDate
	// The time immediately before the task started the TLS security handshake to secure the current connection.
	SecureConnectionStartDate() INSDate
	// The time immediately after the security handshake completed.
	SecureConnectionEndDate() INSDate
	// The time immediately after the task finished establishing the connection to the server.
	ConnectEndDate() INSDate
	// The time immediately before the task started requesting the resource, regardless of whether it is retrieved from the server or local resources.
	RequestStartDate() INSDate
	// The time immediately after the task finished requesting the resource, regardless of whether it was retrieved from the server or local resources.
	RequestEndDate() INSDate
	// The time immediately after the task received the first byte of the response from the server or from local resources.
	ResponseStartDate() INSDate
	// The time immediately after the task received the last byte of the resource.
	ResponseEndDate() INSDate

	// Topic: Accessing data transfer metrics

	// The size of the upload body data, file, or stream, in bytes.
	CountOfRequestBodyBytesBeforeEncoding() int64
	// The number of bytes transferred for the request body.
	CountOfRequestBodyBytesSent() int64
	// The number of bytes transferred for the request header.
	CountOfRequestHeaderBytesSent() int64
	// The size of data delivered to your delegate or completion handler.
	CountOfResponseBodyBytesAfterDecoding() int64
	// The number of bytes transferred for the response body.
	CountOfResponseBodyBytesReceived() int64
	// The number of bytes transferred for the response header.
	CountOfResponseHeaderBytesReceived() int64

	// Topic: Accessing transaction characteristics

	// The network protocol used to fetch the resource.
	NetworkProtocolName() string
	// The IP address string of the remote interface for the connection.
	RemoteAddress() string
	// The IP address string of the local interface for the connection.
	LocalAddress() string
	// The TLS protocol version the task negotiated with the endpoint for the connection.
	NegotiatedTLSProtocolVersion() uint16
	SetNegotiatedTLSProtocolVersion(value uint16)
	// A Boolean value that indicates whether the connection operates over a cellular interface.
	Cellular() bool
	// A Boolean value that indicates whether the connection operates over an expensive interface.
	Expensive() bool
	// A Boolean value that indicates whether the connection operates over an interface marked as constrained.
	Constrained() bool
	// A Boolean value that indicastes whether the task used a proxy connection to fetch the resource.
	ProxyConnection() bool
	// A Boolean value that indicates whether the task used a persistent connection to fetch the resource.
	ReusedConnection() bool
	// A Boolean value that indicates whether the connection uses a successfully negotiated multipath protocol.
	Multipath() bool
	// A value that indicates whether the resource was loaded, pushed, or retrieved from the local cache.
	ResourceFetchType() NSURLSessionTaskMetricsResourceFetchType
	DomainResolutionProtocol() NSURLSessionTaskMetricsDomainResolutionProtocol

	// The port number of the local interface for the connection.
	LocalPort() INSNumber
	// The TLS cipher suite the task negotiated with the endpoint for the connection.
	NegotiatedTLSCipherSuite() INSNumber
	// The number of redirects that occurred during the execution of the task.
	RedirectCount() int
	SetRedirectCount(value int)
	// The port number of the remote interface for the connection.
	RemotePort() INSNumber
	// The time interval between when a task is instantiated and when the task is completed.
	TaskInterval() INSDateInterval
	SetTaskInterval(value INSDateInterval)
	// An array of metrics for each individual request-response transaction made during the execution of the task.
	TransactionMetrics() INSURLSessionTaskTransactionMetrics
	SetTransactionMetrics(value INSURLSessionTaskTransactionMetrics)
}

// Init initializes the instance.
func (u URLSessionTaskTransactionMetrics) Init() URLSessionTaskTransactionMetrics {
	rv := objc.Send[URLSessionTaskTransactionMetrics](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u URLSessionTaskTransactionMetrics) Autorelease() URLSessionTaskTransactionMetrics {
	rv := objc.Send[URLSessionTaskTransactionMetrics](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionTaskTransactionMetrics creates a new URLSessionTaskTransactionMetrics instance.
func NewURLSessionTaskTransactionMetrics() URLSessionTaskTransactionMetrics {
	class := getURLSessionTaskTransactionMetricsClass()
	rv := objc.Send[URLSessionTaskTransactionMetrics](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The transaction request.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/request
func (u URLSessionTaskTransactionMetrics) Request() INSURLRequest {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("request"))
	return NSURLRequestFromID(objc.ID(rv))
}
// The transaction response.
//
// # Discussion
// 
// This property is `nil` if an error occurred and no response was generated.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/response
func (u URLSessionTaskTransactionMetrics) Response() INSURLResponse {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("response"))
	return NSURLResponseFromID(objc.ID(rv))
}
// The time when the task started fetching the resource, from the server or
// locally.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/fetchStartDate
func (u URLSessionTaskTransactionMetrics) FetchStartDate() INSDate {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("fetchStartDate"))
	return NSDateFromID(objc.ID(rv))
}
// The time immediately before the task started the name lookup for the
// resource.
//
// # Discussion
// 
// This value will be `nil` if a persistent connection was used, or if the
// resource was retrieved from local resources.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/domainLookupStartDate
func (u URLSessionTaskTransactionMetrics) DomainLookupStartDate() INSDate {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("domainLookupStartDate"))
	return NSDateFromID(objc.ID(rv))
}
// The time after the name lookup was completed.
//
// # Discussion
// 
// This value will be `nil` if a persistent connection was used, or if the
// resource was retrieved from local resources.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/domainLookupEndDate
func (u URLSessionTaskTransactionMetrics) DomainLookupEndDate() INSDate {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("domainLookupEndDate"))
	return NSDateFromID(objc.ID(rv))
}
// The time immediately before the task started establishing a TCP connection
// to the server.
//
// # Discussion
// 
// This value will be `nil` if a persistent connection is used, or if the
// resource is retrieved from local resources.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/connectStartDate
func (u URLSessionTaskTransactionMetrics) ConnectStartDate() INSDate {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("connectStartDate"))
	return NSDateFromID(objc.ID(rv))
}
// The time immediately before the task started the TLS security handshake to
// secure the current connection.
//
// # Discussion
// 
// This value is `nil` if an encrypted connection isn’t used, if a
// persistent connection is used, or if the resource is retrieved from local
// resources.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/secureConnectionStartDate
func (u URLSessionTaskTransactionMetrics) SecureConnectionStartDate() INSDate {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("secureConnectionStartDate"))
	return NSDateFromID(objc.ID(rv))
}
// The time immediately after the security handshake completed.
//
// # Discussion
// 
// This value is `nil` if an encrypted connection isn’t used, if a
// persistent connection is used, or if the resource is retrieved from local
// resources.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/secureConnectionEndDate
func (u URLSessionTaskTransactionMetrics) SecureConnectionEndDate() INSDate {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("secureConnectionEndDate"))
	return NSDateFromID(objc.ID(rv))
}
// The time immediately after the task finished establishing the connection to
// the server.
//
// # Discussion
// 
// This value accounts for completion of security-related and other
// handshakes. The value will be `nil` if a persistent connection is used, or
// if the resource is retrieved from local resources.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/connectEndDate
func (u URLSessionTaskTransactionMetrics) ConnectEndDate() INSDate {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("connectEndDate"))
	return NSDateFromID(objc.ID(rv))
}
// The time immediately before the task started requesting the resource,
// regardless of whether it is retrieved from the server or local resources.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/requestStartDate
func (u URLSessionTaskTransactionMetrics) RequestStartDate() INSDate {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("requestStartDate"))
	return NSDateFromID(objc.ID(rv))
}
// The time immediately after the task finished requesting the resource,
// regardless of whether it was retrieved from the server or local resources.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/requestEndDate
func (u URLSessionTaskTransactionMetrics) RequestEndDate() INSDate {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("requestEndDate"))
	return NSDateFromID(objc.ID(rv))
}
// The time immediately after the task received the first byte of the response
// from the server or from local resources.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/responseStartDate
func (u URLSessionTaskTransactionMetrics) ResponseStartDate() INSDate {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("responseStartDate"))
	return NSDateFromID(objc.ID(rv))
}
// The time immediately after the task received the last byte of the resource.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/responseEndDate
func (u URLSessionTaskTransactionMetrics) ResponseEndDate() INSDate {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("responseEndDate"))
	return NSDateFromID(objc.ID(rv))
}
// The size of the upload body data, file, or stream, in bytes.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/countOfRequestBodyBytesBeforeEncoding
func (u URLSessionTaskTransactionMetrics) CountOfRequestBodyBytesBeforeEncoding() int64 {
	rv := objc.Send[int64](u.ID, objc.Sel("countOfRequestBodyBytesBeforeEncoding"))
	return rv
}
// The number of bytes transferred for the request body.
//
// # Discussion
// 
// This value includes protocol-specific framing, transfer encoding, and
// content encoding.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/countOfRequestBodyBytesSent
func (u URLSessionTaskTransactionMetrics) CountOfRequestBodyBytesSent() int64 {
	rv := objc.Send[int64](u.ID, objc.Sel("countOfRequestBodyBytesSent"))
	return rv
}
// The number of bytes transferred for the request header.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/countOfRequestHeaderBytesSent
func (u URLSessionTaskTransactionMetrics) CountOfRequestHeaderBytesSent() int64 {
	rv := objc.Send[int64](u.ID, objc.Sel("countOfRequestHeaderBytesSent"))
	return rv
}
// The size of data delivered to your delegate or completion handler.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/countOfResponseBodyBytesAfterDecoding
func (u URLSessionTaskTransactionMetrics) CountOfResponseBodyBytesAfterDecoding() int64 {
	rv := objc.Send[int64](u.ID, objc.Sel("countOfResponseBodyBytesAfterDecoding"))
	return rv
}
// The number of bytes transferred for the response body.
//
// # Discussion
// 
// This value includes protocol-specific framing, transfer encoding, and
// content encoding.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/countOfResponseBodyBytesReceived
func (u URLSessionTaskTransactionMetrics) CountOfResponseBodyBytesReceived() int64 {
	rv := objc.Send[int64](u.ID, objc.Sel("countOfResponseBodyBytesReceived"))
	return rv
}
// The number of bytes transferred for the response header.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/countOfResponseHeaderBytesReceived
func (u URLSessionTaskTransactionMetrics) CountOfResponseHeaderBytesReceived() int64 {
	rv := objc.Send[int64](u.ID, objc.Sel("countOfResponseHeaderBytesReceived"))
	return rv
}
// The network protocol used to fetch the resource.
//
// # Discussion
// 
// When a proxy is configured and a tunnel connection is established, this
// attribute returns the value for the tunneled protocol, which is identified
// by the ALPN Protocol ID Identification Sequence, as per [RFC 7310]. For
// example:
// 
// - If no proxy is used, and HTTP/2 is negotiated, then `h2` is returned. -
// If HTTP/1.1 is used with the proxy, and the tunneled connection is HTTP/2,
// then `h2` is returned. - If HTTP/1.1 is used with the proxy, and there’s
// no tunnel, then `http/1.1` is returned.
//
// [RFC 7310]: https://tools.ietf.org/html/rfc7301
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/networkProtocolName
func (u URLSessionTaskTransactionMetrics) NetworkProtocolName() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("networkProtocolName"))
	return NSStringFromID(rv).String()
}
// The IP address string of the remote interface for the connection.
//
// # Discussion
// 
// For multipath protocols, this is the remote address of the initial flow. If
// the app didn’t use the connection, this value is `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/remoteAddress
func (u URLSessionTaskTransactionMetrics) RemoteAddress() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("remoteAddress"))
	return NSStringFromID(rv).String()
}
// The IP address string of the local interface for the connection.
//
// # Discussion
// 
// For multipath protocols, this is the local address of the initial flow. If
// the app didn’t use the connection, this value is `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/localAddress
func (u URLSessionTaskTransactionMetrics) LocalAddress() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("localAddress"))
	return NSStringFromID(rv).String()
}
// The TLS protocol version the task negotiated with the endpoint for the
// connection.
//
// See: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/negotiatedtlsprotocolversion
func (u URLSessionTaskTransactionMetrics) NegotiatedTLSProtocolVersion() uint16 {
	rv := objc.Send[uint16](u.ID, objc.Sel("negotiatedTLSProtocolVersion"))
	return rv
}
func (u URLSessionTaskTransactionMetrics) SetNegotiatedTLSProtocolVersion(value uint16) {
	objc.Send[struct{}](u.ID, objc.Sel("setNegotiatedTLSProtocolVersion:"), value)
}
// A Boolean value that indicates whether the connection operates over a
// cellular interface.
//
// # Discussion
// 
// You permit or deny use of cellular interfaces with the
// [AllowsCellularAccess] property on [NSURLSessionConfiguration] or
// [allowsCellularAccess] on [URLRequest].
//
// [URLRequest]: https://developer.apple.com/documentation/Foundation/URLRequest
// [allowsCellularAccess]: https://developer.apple.com/documentation/Foundation/URLRequest/allowsCellularAccess
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/isCellular
func (u URLSessionTaskTransactionMetrics) Cellular() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isCellular"))
	return rv
}
// A Boolean value that indicates whether the connection operates over an
// expensive interface.
//
// # Discussion
// 
// The system considers an interface expensive if it’s more costly or
// consumes more power, such as 3G or LTE as compared to ethernet or Wi-Fi.
// You permit or deny use of expensive interfaces with the
// [AllowsExpensiveNetworkAccess] property on [NSURLSessionConfiguration] or
// [allowsExpensiveNetworkAccess] on [URLRequest].
//
// [URLRequest]: https://developer.apple.com/documentation/Foundation/URLRequest
// [allowsExpensiveNetworkAccess]: https://developer.apple.com/documentation/Foundation/URLRequest/allowsExpensiveNetworkAccess
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/isExpensive
func (u URLSessionTaskTransactionMetrics) Expensive() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isExpensive"))
	return rv
}
// A Boolean value that indicates whether the connection operates over an
// interface marked as constrained.
//
// # Discussion
// 
// A constrained interface is one the user marks as constrained by selecting
// “Low Data Mode” in the Settings app. You permit or deny use of
// constrained interfaces with the [AllowsConstrainedNetworkAccess] property
// on [NSURLSessionConfiguration] or [allowsConstrainedNetworkAccess] on
// [URLRequest].
//
// [URLRequest]: https://developer.apple.com/documentation/Foundation/URLRequest
// [allowsConstrainedNetworkAccess]: https://developer.apple.com/documentation/Foundation/URLRequest/allowsConstrainedNetworkAccess
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/isConstrained
func (u URLSessionTaskTransactionMetrics) Constrained() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isConstrained"))
	return rv
}
// A Boolean value that indicastes whether the task used a proxy connection to
// fetch the resource.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/isProxyConnection
func (u URLSessionTaskTransactionMetrics) ProxyConnection() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isProxyConnection"))
	return rv
}
// A Boolean value that indicates whether the task used a persistent
// connection to fetch the resource.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/isReusedConnection
func (u URLSessionTaskTransactionMetrics) ReusedConnection() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isReusedConnection"))
	return rv
}
// A Boolean value that indicates whether the connection uses a successfully
// negotiated multipath protocol.
//
// # Discussion
// 
// You configure the use of multipath protocols with the
// [MultipathServiceType] property on [NSURLSessionConfiguration].
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/isMultipath
func (u URLSessionTaskTransactionMetrics) Multipath() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isMultipath"))
	return rv
}
// A value that indicates whether the resource was loaded, pushed, or
// retrieved from the local cache.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/resourceFetchType
func (u URLSessionTaskTransactionMetrics) ResourceFetchType() NSURLSessionTaskMetricsResourceFetchType {
	rv := objc.Send[NSURLSessionTaskMetricsResourceFetchType](u.ID, objc.Sel("resourceFetchType"))
	return NSURLSessionTaskMetricsResourceFetchType(rv)
}
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/domainResolutionProtocol
func (u URLSessionTaskTransactionMetrics) DomainResolutionProtocol() NSURLSessionTaskMetricsDomainResolutionProtocol {
	rv := objc.Send[NSURLSessionTaskMetricsDomainResolutionProtocol](u.ID, objc.Sel("domainResolutionProtocol"))
	return NSURLSessionTaskMetricsDomainResolutionProtocol(rv)
}
// The port number of the local interface for the connection.
//
// # Discussion
// 
// For multipath protocols, this is the local port of the initial flow. If the
// app didn’t use the connection, this value is `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLSessionTaskTransactionMetrics/localPort
func (u URLSessionTaskTransactionMetrics) LocalPort() INSNumber {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("localPort"))
	return NSNumberFromID(objc.ID(rv))
}
// The TLS cipher suite the task negotiated with the endpoint for the
// connection.
//
// # Discussion
// 
// This value is a 2-byte sequence in host byte order. See [tls_ciphersuite_t]
// for possible values. If the task didn’t negotiate an encrypted
// connection, this value is `nil`.
//
// [tls_ciphersuite_t]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t
//
// See: https://developer.apple.com/documentation/Foundation/NSURLSessionTaskTransactionMetrics/negotiatedTLSCipherSuite
func (u URLSessionTaskTransactionMetrics) NegotiatedTLSCipherSuite() INSNumber {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("negotiatedTLSCipherSuite"))
	return NSNumberFromID(objc.ID(rv))
}
// The number of redirects that occurred during the execution of the task.
//
// See: https://developer.apple.com/documentation/foundation/urlsessiontaskmetrics/redirectcount
func (u URLSessionTaskTransactionMetrics) RedirectCount() int {
	rv := objc.Send[int](u.ID, objc.Sel("redirectCount"))
	return rv
}
func (u URLSessionTaskTransactionMetrics) SetRedirectCount(value int) {
	objc.Send[struct{}](u.ID, objc.Sel("setRedirectCount:"), value)
}
// The port number of the remote interface for the connection.
//
// # Discussion
// 
// For multipath protocols, this is the remote port of the initial flow. If
// the app didn’t use the connection, this value is `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLSessionTaskTransactionMetrics/remotePort
func (u URLSessionTaskTransactionMetrics) RemotePort() INSNumber {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("remotePort"))
	return NSNumberFromID(objc.ID(rv))
}
// The time interval between when a task is instantiated and when the task is
// completed.
//
// See: https://developer.apple.com/documentation/foundation/urlsessiontaskmetrics/taskinterval
func (u URLSessionTaskTransactionMetrics) TaskInterval() INSDateInterval {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("taskInterval"))
	return NSDateIntervalFromID(objc.ID(rv))
}
func (u URLSessionTaskTransactionMetrics) SetTaskInterval(value INSDateInterval) {
	objc.Send[struct{}](u.ID, objc.Sel("setTaskInterval:"), value)
}
// An array of metrics for each individual request-response transaction made
// during the execution of the task.
//
// See: https://developer.apple.com/documentation/foundation/urlsessiontaskmetrics/transactionmetrics
func (u URLSessionTaskTransactionMetrics) TransactionMetrics() INSURLSessionTaskTransactionMetrics {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("transactionMetrics"))
	return NSURLSessionTaskTransactionMetricsFromID(objc.ID(rv))
}
func (u URLSessionTaskTransactionMetrics) SetTransactionMetrics(value INSURLSessionTaskTransactionMetrics) {
	objc.Send[struct{}](u.ID, objc.Sel("setTransactionMetrics:"), value)
}

