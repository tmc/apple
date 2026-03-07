// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf


// A protocol that defines methods that URL session instances call on their delegates to handle task-level events specific to WebSocket tasks.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketDelegate
type NSURLSessionWebSocketDelegate interface {
	objectivec.IObject
	NSURLSessionDelegate
	NSURLSessionTaskDelegate
}



// NSURLSessionWebSocketDelegateObject wraps an existing Objective-C object that conforms to the NSURLSessionWebSocketDelegate protocol.
type NSURLSessionWebSocketDelegateObject struct {
	objectivec.Object
}
func (o NSURLSessionWebSocketDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSURLSessionWebSocketDelegateObjectFromID constructs a [NSURLSessionWebSocketDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSURLSessionWebSocketDelegateObjectFromID(id objc.ID) NSURLSessionWebSocketDelegateObject {
	return NSURLSessionWebSocketDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Tells the delegate that the WebSocket task successfully negotiated the
// handshake with the endpoint, indicating the negotiated protocol.
//
// session: The session of the WebSocket task that opened.
//
// webSocketTask: The WebSocket task that opened.
//
// protocol: The protocol picked during the handshake phase. This parameter is `nil` if
// the server did not pick a protocol, or if the client did not advertise
// protocols when creating the task.
//
// # Discussion
// 
// If the handshake fails, the task doesn’t call this delegate method.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketDelegate/urlSession(_:webSocketTask:didOpenWithProtocol:)

func (o NSURLSessionWebSocketDelegateObject) URLSessionWebSocketTaskDidOpenWithProtocol(session INSURLSession, webSocketTask INSURLSessionWebSocketTask, protocol_ string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:webSocketTask:didOpenWithProtocol:"), session, webSocketTask, objc.String(protocol_))
	}

// Tells the delegate that the WebSocket task received a close frame from the
// server endpoint, optionally including a close code and reason from the
// server.
//
// session: The session of the WebSocket task that closed.
//
// webSocketTask: The WebSocket task that closed.
//
// closeCode: The close code provided by the server. If the close frame didn’t include
// a close code, this value is `nil`.
//
// reason: The close reason provided by the server. If the close frame didn’t
// include a reason, this value is `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketDelegate/urlSession(_:webSocketTask:didCloseWith:reason:)

func (o NSURLSessionWebSocketDelegateObject) URLSessionWebSocketTaskDidCloseWithCodeReason(session INSURLSession, webSocketTask INSURLSessionWebSocketTask, closeCode NSURLSessionWebSocketCloseCode, reason INSData) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:webSocketTask:didCloseWithCode:reason:"), session, webSocketTask, closeCode, reason)
	}

// Tells the URL session that the session has been invalidated.
//
// session: The session object that was invalidated.
//
// error: The error that caused invalidation, or `nil` if the invalidation was
// explicit.
//
// # Discussion
// 
// If you invalidate a session by calling its [FinishTasksAndInvalidate]
// method, the session waits until after the final task in the session
// finishes or fails before calling this delegate method. If you call the
// [InvalidateAndCancel] method, the session calls this delegate method
// immediately.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionDelegate/urlSession(_:didBecomeInvalidWithError:)

func (o NSURLSessionWebSocketDelegateObject) URLSessionDidBecomeInvalidWithError(session INSURLSession, error_ INSError) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:didBecomeInvalidWithError:"), session, error_)
	}

// Tells the delegate that all messages enqueued for a session have been
// delivered.
//
// session: The session that no longer has any outstanding requests.
//
// # Discussion
// 
// In iOS, when a background transfer completes or requires credentials, if
// your app is no longer running, your app is automatically relaunched in the
// background, and the app’s [UIApplicationDelegate] is sent an
// [application(_:handleEventsForBackgroundURLSession:completionHandler:)]
// message. This call contains the identifier of the session that caused your
// app to be launched. You should then store that completion handler before
// creating a background configuration object with the same identifier, and
// creating a session with that configuration. The newly created session is
// automatically reassociated with ongoing background activity.
// 
// When your app later receives a
// [URLSessionDidFinishEventsForBackgroundURLSession] message, this indicates
// that all messages previously enqueued for this session have been delivered,
// and that it is now safe to invoke the previously stored completion handler
// or to begin any internal updates that may result in invoking the completion
// handler.
//
// [application(_:handleEventsForBackgroundURLSession:completionHandler:)]: https://developer.apple.com/documentation/UIKit/UIApplicationDelegate/application(_:handleEventsForBackgroundURLSession:completionHandler:)
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionDelegate/urlSessionDidFinishEvents(forBackgroundURLSession:)

func (o NSURLSessionWebSocketDelegateObject) URLSessionDidFinishEventsForBackgroundURLSession(session INSURLSession) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSessionDidFinishEventsForBackgroundURLSession:"), session)
	}

// Requests credentials from the delegate in response to a session-level
// authentication request from the remote server.
//
// session: The session containing the task that requested authentication.
//
// challenge: An object that contains the request for authentication.
//
// completionHandler: A handler that your delegate method must call. This completion handler
// takes the following parameters::
// 
// - `disposition`—One of several constants that describes how the challenge
// should be handled. - `credential`—The credential that should be used for
// authentication if disposition is [NSURLSessionAuthChallengeUseCredential],
// otherwise [NULL].
//
// # Discussion
// 
// This method is called in two situations:
// 
// - When a remote server asks for client certificates or Windows NT LAN
// Manager (NTLM) authentication, to allow your app to provide appropriate
// credentials - When a session first establishes a connection to a remote
// server that uses SSL or TLS, to allow your app to verify the server’s
// certificate chain
// 
// If you do not implement this method, the session calls its delegate’s
// [URLSessionTaskDidReceiveChallengeCompletionHandler] method instead.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionDelegate/urlSession(_:didReceive:completionHandler:)

func (o NSURLSessionWebSocketDelegateObject) URLSessionDidReceiveChallengeCompletionHandler(session INSURLSession, challenge INSURLAuthenticationChallenge, completionHandler URLSessionAuthChallengeDispositionURLCredentialHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:didReceiveChallenge:completionHandler:"), session, challenge, completionHandler)
	}

// Tells the delegate that the task finished transferring data.
//
// session: The session containing the task that has finished transferring data.
//
// task: The task that has finished transferring data.
//
// error: If an error occurred, an error object indicating how the transfer failed,
// otherwise [NULL].
//
// # Discussion
// 
// The only errors your delegate receives through the `error` parameter are
// client-side errors, such as being unable to resolve the hostname or connect
// to the host. To check for server-side errors, inspect the [Response]
// property of the `task` parameter received by this callback.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskDelegate/urlSession(_:task:didCompleteWithError:)

func (o NSURLSessionWebSocketDelegateObject) URLSessionTaskDidCompleteWithError(session INSURLSession, task INSURLSessionTask, error_ INSError) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:task:didCompleteWithError:"), session, task, error_)
	}

// Tells the delegate that the remote server requested an HTTP redirect.
//
// session: The session containing the task whose request resulted in a redirect.
//
// task: The task whose request resulted in a redirect.
//
// response: An object containing the server’s response to the original request.
//
// request: A URL request object filled out with the new location.
//
// completionHandler: A block that your handler should call with either the value of the
// `request` parameter, a modified URL request object, or [NULL] to refuse the
// redirect and return the body of the redirect response.
//
// # Discussion
// 
// This method is called for tasks in default and ephemeral sessions. Tasks in
// background sessions automatically follow redirects.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskDelegate/urlSession(_:task:willPerformHTTPRedirection:newRequest:completionHandler:)

func (o NSURLSessionWebSocketDelegateObject) URLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler(session INSURLSession, task INSURLSessionTask, response INSHTTPURLResponse, request INSURLRequest, completionHandler URLRequestHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:task:willPerformHTTPRedirection:newRequest:completionHandler:"), session, task, response, request, completionHandler)
	}

// Periodically informs the delegate of the progress of sending body content
// to the server.
//
// session: The session containing the data task.
//
// task: The data task.
//
// bytesSent: The number of bytes sent since the last time this delegate method was
// called.
//
// totalBytesSent: The total number of bytes sent so far.
//
// totalBytesExpectedToSend: The expected length of the body data. The URL loading system can determine
// the length of the upload data in three ways:
// 
// - From the length of the [NSData] object provided as the upload body. -
// From the length of the file on disk provided as the upload body of an
// upload task ( a download task). - From the `Content-Length` in the request
// object, if you explicitly set it.
// 
// Otherwise, the value is [NSURLSessionTransferSizeUnknown] (`-1`) if you
// provided a stream or body data object, or zero (`0`) if you did not.
// //
// [NSURLSessionTransferSizeUnknown]: https://developer.apple.com/documentation/Foundation/NSURLSessionTransferSizeUnknown
//
// # Discussion
// 
// The `totalBytesSent` and `totalBytesExpectedToSend` parameters are also
// available as [NSURLSessionTask] properties [CountOfBytesSent] and
// [CountOfBytesExpectedToSend]. Or, since [NSURLSessionTask] supports
// [NSProgressReporting], you can use the task’s [Progress] property
// instead, which may be more convenient.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskDelegate/urlSession(_:task:didSendBodyData:totalBytesSent:totalBytesExpectedToSend:)

func (o NSURLSessionWebSocketDelegateObject) URLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend(session INSURLSession, task INSURLSessionTask, bytesSent int64, totalBytesSent int64, totalBytesExpectedToSend int64) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:task:didSendBodyData:totalBytesSent:totalBytesExpectedToSend:"), session, task, bytesSent, totalBytesSent, totalBytesExpectedToSend)
	}

// Tells the delegate when a task requires a new request body stream to send
// to the remote server.
//
// session: The session containing the task that needs a new body stream.
//
// task: The task that needs a new body stream.
//
// completionHandler: A completion handler that your delegate method should call with the new
// body stream.
//
// # Discussion
// 
// The task calls this delegate method under two circumstances:
// 
// - To provide the initial request body stream if the task was created with
// [UploadTaskWithStreamedRequest] - To provide a replacement request body
// stream if the task needs to resend a request that has a body stream because
// of an authentication challenge or other recoverable server error.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskDelegate/urlSession(_:task:needNewBodyStream:)

func (o NSURLSessionWebSocketDelegateObject) URLSessionTaskNeedNewBodyStream(session INSURLSession, task INSURLSessionTask, completionHandler InputStreamHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:task:needNewBodyStream:"), session, task, completionHandler)
	}

// Requests credentials from the delegate in response to an authentication
// request from the remote server.
//
// session: The session containing the task whose request requires authentication.
//
// task: The task whose request requires authentication.
//
// challenge: An object that contains the request for authentication.
//
// completionHandler: A handler that your delegate method must call. Its parameters are:
// 
// - `disposition`—One of several constants that describes how the challenge
// should be handled. - `credential`—The credential that should be used for
// authentication if disposition is [NSURLSessionAuthChallengeUseCredential];
// otherwise, [NULL].
//
// # Discussion
// 
// This method handles task-level authentication challenges. The
// [NSURLSessionDelegate] protocol also provides a session-level
// authentication delegate method. The method called depends on the type of
// authentication challenge:
// 
// - For session-level challenges—[NSURLAuthenticationMethodNTLM],
// [NSURLAuthenticationMethodNegotiate],
// [NSURLAuthenticationMethodClientCertificate], or
// [NSURLAuthenticationMethodServerTrust]—the [NSURLSession] object calls
// the session delegate’s [URLSessionDidReceiveChallengeCompletionHandler]
// method. If your app does not provide a session delegate method, the
// [NSURLSession] object calls the task delegate’s
// [URLSessionTaskDidReceiveChallengeCompletionHandler] method to handle the
// challenge. - For non-session-level challenges (all others), the
// [NSURLSession] object calls the session delegate’s
// [URLSessionTaskDidReceiveChallengeCompletionHandler] method to handle the
// challenge. If your app provides a session delegate and you need to handle
// authentication, then you must either handle the authentication at the task
// level or provide a task-level handler that calls the per-session handler
// explicitly. The session delegate’s
// [URLSessionDidReceiveChallengeCompletionHandler] method is called for
// non-session-level challenges.
//
// [NSURLAuthenticationMethodClientCertificate]: https://developer.apple.com/documentation/Foundation/NSURLAuthenticationMethodClientCertificate
// [NSURLAuthenticationMethodNTLM]: https://developer.apple.com/documentation/Foundation/NSURLAuthenticationMethodNTLM
// [NSURLAuthenticationMethodNegotiate]: https://developer.apple.com/documentation/Foundation/NSURLAuthenticationMethodNegotiate
// [NSURLAuthenticationMethodServerTrust]: https://developer.apple.com/documentation/Foundation/NSURLAuthenticationMethodServerTrust
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskDelegate/urlSession(_:task:didReceive:completionHandler:)

func (o NSURLSessionWebSocketDelegateObject) URLSessionTaskDidReceiveChallengeCompletionHandler(session INSURLSession, task INSURLSessionTask, challenge INSURLAuthenticationChallenge, completionHandler URLSessionAuthChallengeDispositionURLCredentialHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:task:didReceiveChallenge:completionHandler:"), session, task, challenge, completionHandler)
	}

// Tells the delegate that a delayed URL session task will now begin loading.
//
// session: The session containing the delayed request.
//
// task: The task handling the delayed request.
//
// request: The request that was delayed.
//
// completionHandler: A completion handler to perform the request. The completion handler takes
// two parameters: a disposition that tells the task how to proceed, and a new
// request object that is only used if the disposition is
// [URLSession.DelayedRequestDisposition.useNewRequest].
// //
// [URLSession.DelayedRequestDisposition.useNewRequest]: https://developer.apple.com/documentation/Foundation/URLSession/DelayedRequestDisposition/useNewRequest
//
// # Discussion
// 
// This method is called when a background session task with a delayed start
// time (as set with the [EarliestBeginDate] property) is ready to start. This
// delegate method should only be implemented if the request might become
// stale while waiting for the network load and needs to be replaced by a new
// request.
// 
// For loading to continue, the delegate must call the completion handler,
// passing in a disposition that indicates how the task should proceed.
// Passing the [URLSession.DelayedRequestDisposition.cancel] disposition is
// equivalent to calling [Cancel] on the task directly.
//
// [URLSession.DelayedRequestDisposition.cancel]: https://developer.apple.com/documentation/Foundation/URLSession/DelayedRequestDisposition/cancel
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskDelegate/urlSession(_:task:willBeginDelayedRequest:completionHandler:)

func (o NSURLSessionWebSocketDelegateObject) URLSessionTaskWillBeginDelayedRequestCompletionHandler(session INSURLSession, task INSURLSessionTask, request INSURLRequest, completionHandler URLSessionDelayedRequestDispositionURLRequestHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:task:willBeginDelayedRequest:completionHandler:"), session, task, request, completionHandler)
	}

// Tells the delegate that the task is waiting until suitable connectivity is
// available before beginning the network load.
//
// session: The session that contains the waiting task.
//
// task: The task that is waiting for a change in connectivity.
//
// # Discussion
// 
// This method is called if the [WaitsForConnectivity] property of
// [NSURLSessionConfiguration] is `true`, and sufficient connectivity is
// unavailable. The delegate can use this opportunity to update the user
// interface; for example, by presenting an offline mode or a cellular-only
// mode.
// 
// This method is called, at most, once per task, and only if connectivity is
// initially unavailable. It is never called for background sessions because
// `waitsForConnectivity` is ignored for those sessions.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskDelegate/urlSession(_:taskIsWaitingForConnectivity:)

func (o NSURLSessionWebSocketDelegateObject) URLSessionTaskIsWaitingForConnectivity(session INSURLSession, task INSURLSessionTask) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:taskIsWaitingForConnectivity:"), session, task)
	}

// Tells the delegate that the session finished collecting metrics for the
// task.
//
// session: The session collecting the metrics.
//
// task: The task whose metrics have been collected.
//
// metrics: The collected metrics.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskDelegate/urlSession(_:task:didFinishCollecting:)

func (o NSURLSessionWebSocketDelegateObject) URLSessionTaskDidFinishCollectingMetrics(session INSURLSession, task INSURLSessionTask, metrics INSURLSessionTaskMetrics) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:task:didFinishCollectingMetrics:"), session, task, metrics)
	}

//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskDelegate/urlSession(_:didCreateTask:)

func (o NSURLSessionWebSocketDelegateObject) URLSessionDidCreateTask(session INSURLSession, task INSURLSessionTask) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:didCreateTask:"), session, task)
	}

//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskDelegate/urlSession(_:task:didReceiveInformationalResponse:)

func (o NSURLSessionWebSocketDelegateObject) URLSessionTaskDidReceiveInformationalResponse(session INSURLSession, task INSURLSessionTask, response INSHTTPURLResponse) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:task:didReceiveInformationalResponse:"), session, task, response)
	}

// Tells the delegate if a task requires a new body stream starting from the
// given offset. This may be necessary when resuming a failed upload task.
//
// session: The session containing the task that needs a new body stream from the given
// offset.
//
// task: The task that needs a new body stream.
//
// offset: The starting offset required for the body stream.
//
// completionHandler: A completion handler that your delegate method should call with the new
// body stream.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskDelegate/urlSession(_:task:needNewBodyStreamFrom:completionHandler:)

func (o NSURLSessionWebSocketDelegateObject) URLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler(session INSURLSession, task INSURLSessionTask, offset int64, completionHandler InputStreamHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:task:needNewBodyStreamFromOffset:completionHandler:"), session, task, offset, completionHandler)
	}





// NSURLSessionWebSocketDelegateConfig holds optional typed callbacks for [NSURLSessionWebSocketDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsocketdelegate
type NSURLSessionWebSocketDelegateConfig struct {

	// Other Methods
	// URLSessionWebSocketTaskDidCloseWithCodeReason — Tells the delegate that the WebSocket task received a close frame from the server endpoint, optionally including a close code and reason from the server.
	URLSessionWebSocketTaskDidCloseWithCodeReason func(session NSURLSession, webSocketTask NSURLSessionWebSocketTask, closeCode NSURLSessionWebSocketCloseCode, reason objectivec.Object)
}

// NewNSURLSessionWebSocketDelegate creates an Objective-C object implementing the [NSURLSessionWebSocketDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSURLSessionWebSocketDelegateObject] satisfies the [NSURLSessionWebSocketDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsocketdelegate
func NewNSURLSessionWebSocketDelegate(config NSURLSessionWebSocketDelegateConfig) NSURLSessionWebSocketDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSURLSessionWebSocketDelegate_%d", n)

	var methods []objc.MethodDef

	if config.URLSessionWebSocketTaskDidCloseWithCodeReason != nil {
		fn := config.URLSessionWebSocketTaskDidCloseWithCodeReason
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("URLSession:webSocketTask:didCloseWithCode:reason:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, webSocketTaskID objc.ID, closeCode NSURLSessionWebSocketCloseCode, reasonID objc.ID) {
				session := NSURLSessionFromID(sessionID)
				webSocketTask := NSURLSessionWebSocketTaskFromID(webSocketTaskID)
				reason := objectivec.ObjectFromID(reasonID)
				fn(session, webSocketTask, closeCode, reason)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSURLSessionWebSocketDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSURLSessionWebSocketDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSURLSessionWebSocketDelegateObjectFromID(instance)
}






