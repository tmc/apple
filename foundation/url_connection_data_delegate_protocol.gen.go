// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A protocol that most delegates of a URL connection implement to receive data associated with the connection.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDataDelegate
type NSURLConnectionDataDelegate interface {
	objectivec.IObject
	NSURLConnectionDelegate
}

// NSURLConnectionDataDelegateObject wraps an existing Objective-C object that conforms to the NSURLConnectionDataDelegate protocol.
type NSURLConnectionDataDelegateObject struct {
	objectivec.Object
}
func (o NSURLConnectionDataDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSURLConnectionDataDelegateObjectFromID constructs a [NSURLConnectionDataDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSURLConnectionDataDelegateObjectFromID(id objc.ID) NSURLConnectionDataDelegateObject {
	return NSURLConnectionDataDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Sent when the connection has received sufficient data to construct the URL
// response for its request.
//
// connection: The connection sending the message.
//
// response: The URL response for the connection’s request. This object is immutable
// and will not be modified by the URL loading system once it is presented to
// the delegate.
//
// # Discussion
// 
// In rare cases, for example in the case of an HTTP load where the content
// type of the load data is `multipart/x-mixed-replace`, the delegate will
// receive more than one `` message. When this happens, discard (or process)
// all data previously delivered by ``, and prepare to handle the next part
// (which could potentially have a different MIME type).
// 
// The only case where this message is not sent to the delegate is when the
// protocol implementation encounters an error before a response could be
// created.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDataDelegate/connection(_:didReceive:)-8t66w

func (o NSURLConnectionDataDelegateObject) ConnectionDidReceiveResponse(connection INSURLConnection, response INSURLResponse) {
	
	objc.Send[struct{}](o.ID, objc.Sel("connection:didReceiveResponse:"), connection, response)
	}

// Sent as a connection loads data incrementally.
//
// connection: The connection sending the message.
//
// data: The newly available data. The delegate should concatenate the contents of
// each `data` object delivered to build up the complete data for a URL load.
//
// # Discussion
// 
// This method provides the only way for an asynchronous delegate to retrieve
// the loaded data. It is the responsibility of the delegate to retain or copy
// this data as it is delivered.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDataDelegate/connection(_:didReceive:)-8p5vg

func (o NSURLConnectionDataDelegateObject) ConnectionDidReceiveData(connection INSURLConnection, data INSData) {
	
	objc.Send[struct{}](o.ID, objc.Sel("connection:didReceiveData:"), connection, data)
	}

// Sent as the body (message data) of a request is transmitted (such as in an
// HTTP POST request).
//
// connection: The connection sending the message.
//
// bytesWritten: The number of bytes written in the latest write.
//
// totalBytesWritten: The total number of bytes written for this connection.
//
// totalBytesExpectedToWrite: The number of bytes the connection expects to write.
//
// # Discussion
// 
// This method provides an estimate of the progress of a URL upload.
// 
// The value of `totalBytesExpectedToWrite` may change during the upload if
// the request needs to be retransmitted due to a lost connection or an
// authentication challenge from the server.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDataDelegate/connection(_:didSendBodyData:totalBytesWritten:totalBytesExpectedToWrite:)

func (o NSURLConnectionDataDelegateObject) ConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite(connection INSURLConnection, bytesWritten int, totalBytesWritten int, totalBytesExpectedToWrite int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("connection:didSendBodyData:totalBytesWritten:totalBytesExpectedToWrite:"), connection, bytesWritten, totalBytesWritten, totalBytesExpectedToWrite)
	}

// Sent when a connection has finished loading successfully.
//
// connection: The connection sending the message.
//
// # Discussion
// 
// The delegate will receive no further messages for `connection`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDataDelegate/connectionDidFinishLoading(_:)

func (o NSURLConnectionDataDelegateObject) ConnectionDidFinishLoading(connection INSURLConnection) {
	
	objc.Send[struct{}](o.ID, objc.Sel("connectionDidFinishLoading:"), connection)
	}

// Sent when the connection determines that it must change URLs in order to
// continue loading a request.
//
// connection: The connection sending the message.
//
// request: The proposed redirected request. The delegate should inspect the redirected
// request to verify that it meets its needs, and create a copy with new
// attributes to return to the connection if necessary.
//
// response: The URL response that caused the redirect. May be `nil` in cases where this
// method is called because of URL canonicalization.
//
// # Return Value
// 
// The actual URL request to use in light of the redirection response. The
// delegate may return `request` unmodified to allow the redirect, return a
// new request, or return `nil` to reject the redirect and continue processing
// the connection.
//
// # Discussion
// 
// If `redirectResponse` is `nil`, the URL was canonicalized (rewritten into
// its standard form) by the [NSURLProtocol] object handling the request.
// Update your user interface to show the standardized form of the URL, then
// return the original request unmodified.
// 
// Otherwise, to cancel the redirect, call the `connection` object’s
// `cancel` method, then return the provided request object.
// 
// To receive the body of the redirect response itself, return `nil` to cancel
// the redirect. The connection continues to process, eventually sending your
// delegate a `connectionDidFinishLoading` or `` message, as appropriate.
// 
// To redirect the request to a different URL, create a new request object and
// return it.
// 
// The delegate should be prepared to receive this message multiple times.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDataDelegate/connection(_:willSend:redirectResponse:)

func (o NSURLConnectionDataDelegateObject) ConnectionWillSendRequestRedirectResponse(connection INSURLConnection, request INSURLRequest, response INSURLResponse) INSURLRequest {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("connection:willSendRequest:redirectResponse:"), connection, request, response)
	return NSURLRequestFromID(rv)
	}

// Called when an [NSURLConnection] needs to retransmit a request that has a
// body stream to provide a new, unopened stream.
//
// connection: The NSURLConnection that is requesting a new body stream.
//
// # Return Value
// 
// This delegate method should return a new, unopened stream that provides the
// body contents for the request.
// 
// If this delegate method returns [NULL], the connection fails.
//
// # Discussion
// 
// In macOS, if this method is not implemented, body stream data is spooled to
// disk in case retransmission is required. This spooling may not be desirable
// for large data sets.
// 
// By implementing this delegate method, the client opts out of automatic
// spooling, and must provide a new, unopened stream for each retransmission.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDataDelegate/connection(_:needNewBodyStream:)

func (o NSURLConnectionDataDelegateObject) ConnectionNeedNewBodyStream(connection INSURLConnection, request INSURLRequest) INSInputStream {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("connection:needNewBodyStream:"), connection, request)
	return NSInputStreamFromID(rv)
	}

// Sent before the connection stores a cached response in the cache, to give
// the delegate an opportunity to alter it.
//
// connection: The connection sending the message.
//
// cachedResponse: The proposed cached response to store in the cache.
//
// # Return Value
// 
// The actual cached response to store in the cache. The delegate may return
// `cachedResponse` unmodified, return a modified cached response, or return
// `nil` if no cached response should be stored for the connection.
//
// # Discussion
// 
// This method is called only if the [NSURLProtocol] handling the request
// decides to cache the response. As a rule, responses are cached only when
// all of the following are true:
// 
// - The request is for an HTTP or HTTPS URL (or your own custom networking
// protocol that supports caching). - The request was successful (with a
// status code in the `200–299` range). - The provided response came from
// the server, rather than out of the cache. - The [NSURLRequest] object’s
// cache policy allows caching. - The cache-related headers in the server’s
// response (if present) allow caching. - The response size is small enough to
// reasonably fit within the cache. (For example, if you provide a disk cache,
// the response must be no larger than about 5% of the disk cache size.)
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDataDelegate/connection(_:willCacheResponse:)

func (o NSURLConnectionDataDelegateObject) ConnectionWillCacheResponse(connection INSURLConnection, cachedResponse INSCachedURLResponse) INSCachedURLResponse {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("connection:willCacheResponse:"), connection, cachedResponse)
	return NSCachedURLResponseFromID(rv)
	}

// Tells the delegate that the connection will send a request for an
// authentication challenge.
//
// connection: The connection sending the message.
//
// challenge: The authentication challenge for which a request is being sent.
//
// # Discussion
// 
// This method allows the delegate to make an informed decision about
// connection authentication at once. If the delegate implements this method,
// it has no need to implement
// [connection(_:canAuthenticateAgainstProtectionSpace:)] or
// [connection(_:didReceive:)]. In fact, those other methods are not invoked
// (except on older operating systems, where applicable).
// 
// In this method,you invoke one of the challenge-responder methods
// ([NSURLAuthenticationChallengeSender] protocol):
// 
// - [UseCredentialForAuthenticationChallenge] -
// [ContinueWithoutCredentialForAuthenticationChallenge] -
// [CancelAuthenticationChallenge] -
// [PerformDefaultHandlingForAuthenticationChallenge] -
// [RejectProtectionSpaceAndContinueWithChallenge]
// 
// You might also want to analyze `challenge` for the authentication scheme
// and the proposed credential before calling a
// [NSURLAuthenticationChallengeSender] method. You should never assume that a
// proposed credential is present. You can either create your own credential
// and respond with that, or you can send the proposed credential back.
// (Because this object is immutable, if you want to change it you must copy
// it and then modify the copy.)
//
// [connection(_:canAuthenticateAgainstProtectionSpace:)]: https://developer.apple.com/documentation/Foundation/NSURLConnectionDelegate/connection(_:canAuthenticateAgainstProtectionSpace:)
// [connection(_:didReceive:)]: https://developer.apple.com/documentation/Foundation/NSURLConnectionDelegate/connection(_:didReceive:)
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDelegate/connection(_:willSendRequestFor:)

func (o NSURLConnectionDataDelegateObject) ConnectionWillSendRequestForAuthenticationChallenge(connection INSURLConnection, challenge INSURLAuthenticationChallenge) {
	
	objc.Send[struct{}](o.ID, objc.Sel("connection:willSendRequestForAuthenticationChallenge:"), connection, challenge)
	}

// Sent to determine whether the URL loader should use the credential storage
// for authenticating the connection.
//
// connection: The connection sending the message.
//
// # Discussion
// 
// This method is called before any attempt to authenticate is made.
// 
// If you return [false], the connection does not consult the credential
// storage automatically, and does not store credentials. However, in your
// connection:didReceiveAuthenticationChallenge: method, you can consult the
// credential storage yourself and store credentials yourself, as needed.
// 
// Not implementing this method is the same as returning [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDelegate/connectionShouldUseCredentialStorage(_:)

func (o NSURLConnectionDataDelegateObject) ConnectionShouldUseCredentialStorage(connection INSURLConnection) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("connectionShouldUseCredentialStorage:"), connection)
	return rv
	}

// Sent when a connection fails to load its request successfully.
//
// connection: The connection sending the message.
//
// error: An error object containing details of why the connection failed to load the
// request successfully.
//
// # Discussion
// 
// Once the delegate receives this message, it will receive no further
// messages for `connection`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDelegate/connection(_:didFailWithError:)

func (o NSURLConnectionDataDelegateObject) ConnectionDidFailWithError(connection INSURLConnection, error_ INSError) {
	
	objc.Send[struct{}](o.ID, objc.Sel("connection:didFailWithError:"), connection, error_)
	}

// NSURLConnectionDataDelegateConfig holds optional typed callbacks for [NSURLConnectionDataDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsurlconnectiondatadelegate
type NSURLConnectionDataDelegateConfig struct {

	// Receiving Connection Progress
	// ConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite — Sent as the body (message data) of a request is transmitted (such as in an HTTP POST request).
	ConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite func(connection NSURLConnection, bytesWritten int, totalBytesWritten int, totalBytesExpectedToWrite int)
	// ConnectionDidFinishLoading — Sent when a connection has finished loading successfully.
	ConnectionDidFinishLoading func(connection NSURLConnection)

	// Handling Redirects
	// ConnectionNeedNewBodyStream — Called when an [NSURLConnection] needs to retransmit a request that has a body stream to provide a new, unopened stream.
	ConnectionNeedNewBodyStream func(connection NSURLConnection, request NSURLRequest) NSInputStream

	// Overriding Caching Behavior
	// ConnectionWillCacheResponse — Sent before the connection stores a cached response in the cache, to give the delegate an opportunity to alter it.
	ConnectionWillCacheResponse func(connection NSURLConnection, cachedResponse NSCachedURLResponse) NSCachedURLResponse

	// Other Methods
	// ConnectionDidReceiveResponse — Sent when the connection has received sufficient data to construct the URL response for its request.
	ConnectionDidReceiveResponse func(connection NSURLConnection, response NSURLResponse)
	// ConnectionDidReceiveData — Sent as a connection loads data incrementally.
	ConnectionDidReceiveData func(connection NSURLConnection, data objectivec.Object)
	// ConnectionWillSendRequestRedirectResponse — Sent when the connection determines that it must change URLs in order to continue loading a request.
	ConnectionWillSendRequestRedirectResponse func(connection NSURLConnection, request NSURLRequest, response NSURLResponse) NSURLRequest
}

// NewNSURLConnectionDataDelegate creates an Objective-C object implementing the [NSURLConnectionDataDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSURLConnectionDataDelegateObject] satisfies the [NSURLConnectionDataDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsurlconnectiondatadelegate
func NewNSURLConnectionDataDelegate(config NSURLConnectionDataDelegateConfig) NSURLConnectionDataDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSURLConnectionDataDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ConnectionDidReceiveResponse != nil {
		fn := config.ConnectionDidReceiveResponse
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("connection:didReceiveResponse:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connectionID objc.ID, responseID objc.ID) {
				connection := NSURLConnectionFromID(connectionID)
				response := NSURLResponseFromID(responseID)
				fn(connection, response)
			},
		})
	}

	if config.ConnectionDidReceiveData != nil {
		fn := config.ConnectionDidReceiveData
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("connection:didReceiveData:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connectionID objc.ID, dataID objc.ID) {
				connection := NSURLConnectionFromID(connectionID)
				data := objectivec.ObjectFromID(dataID)
				fn(connection, data)
			},
		})
	}

	if config.ConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite != nil {
		fn := config.ConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("connection:didSendBodyData:totalBytesWritten:totalBytesExpectedToWrite:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connectionID objc.ID, bytesWritten int, totalBytesWritten int, totalBytesExpectedToWrite int) {
				connection := NSURLConnectionFromID(connectionID)
				fn(connection, bytesWritten, totalBytesWritten, totalBytesExpectedToWrite)
			},
		})
	}

	if config.ConnectionDidFinishLoading != nil {
		fn := config.ConnectionDidFinishLoading
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("connectionDidFinishLoading:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connectionID objc.ID) {
				connection := NSURLConnectionFromID(connectionID)
				fn(connection)
			},
		})
	}

	if config.ConnectionWillSendRequestRedirectResponse != nil {
		fn := config.ConnectionWillSendRequestRedirectResponse
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("connection:willSendRequest:redirectResponse:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connectionID objc.ID, requestID objc.ID, responseID objc.ID) objc.ID {
				connection := NSURLConnectionFromID(connectionID)
				request := NSURLRequestFromID(requestID)
				response := NSURLResponseFromID(responseID)
				return fn(connection, request, response).GetID()
			},
		})
	}

	if config.ConnectionNeedNewBodyStream != nil {
		fn := config.ConnectionNeedNewBodyStream
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("connection:needNewBodyStream:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connectionID objc.ID, requestID objc.ID) objc.ID {
				connection := NSURLConnectionFromID(connectionID)
				request := NSURLRequestFromID(requestID)
				return fn(connection, request).GetID()
			},
		})
	}

	if config.ConnectionWillCacheResponse != nil {
		fn := config.ConnectionWillCacheResponse
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("connection:willCacheResponse:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connectionID objc.ID, cachedResponseID objc.ID) objc.ID {
				connection := NSURLConnectionFromID(connectionID)
				cachedResponse := NSCachedURLResponseFromID(cachedResponseID)
				return fn(connection, cachedResponse).GetID()
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSURLConnectionDataDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSURLConnectionDataDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSURLConnectionDataDelegateObjectFromID(instance)
}

