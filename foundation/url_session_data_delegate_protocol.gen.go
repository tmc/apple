// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A protocol that defines methods that URL session instances call on their delegates to handle task-level events specific to data and upload tasks.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionDataDelegate
type NSURLSessionDataDelegate interface {
	objectivec.IObject
	NSURLSessionDelegate
	NSURLSessionTaskDelegate
}

// NSURLSessionDataDelegateObject wraps an existing Objective-C object that conforms to the NSURLSessionDataDelegate protocol.
type NSURLSessionDataDelegateObject struct {
	objectivec.Object
}
func (o NSURLSessionDataDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSURLSessionDataDelegateObjectFromID constructs a [NSURLSessionDataDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSURLSessionDataDelegateObjectFromID(id objc.ID) NSURLSessionDataDelegateObject {
	return NSURLSessionDataDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that the data task received the initial reply (headers)
// from the server.
//
// session: The session containing the data task that received an initial reply.
//
// dataTask: The data task that received an initial reply.
//
// response: A URL response object populated with headers.
//
// completionHandler: A completion handler that your code calls to continue a transfer, passing a
// [URLSession.ResponseDisposition] constant to indicate whether the transfer
// should continue as a data task or should become a download task.
// 
// - If you pass [URLSession.ResponseDisposition.allow], the task continues as
// a data task. - If you pass [URLSession.ResponseDisposition.cancel], the
// task is canceled. - If you pass
// [URLSession.ResponseDisposition.becomeDownload], your delegate’s
// [URLSessionDataTaskDidBecomeDownloadTask] method is called to provide the
// new download task that supersedes the current task.
// //
// [URLSession.ResponseDisposition.allow]: https://developer.apple.com/documentation/Foundation/URLSession/ResponseDisposition/allow
// [URLSession.ResponseDisposition.becomeDownload]: https://developer.apple.com/documentation/Foundation/URLSession/ResponseDisposition/becomeDownload
// [URLSession.ResponseDisposition.cancel]: https://developer.apple.com/documentation/Foundation/URLSession/ResponseDisposition/cancel
// [URLSession.ResponseDisposition]: https://developer.apple.com/documentation/Foundation/URLSession/ResponseDisposition
//
// # Discussion
// 
// Implementing this method is optional unless you need to cancel the transfer
// or convert it to a download task when the response headers are first
// received. If you don’t provide this delegate method, the session always
// allows the task to continue.
// 
// You also implement this method if you need to support the fairly obscure
// `multipart/x-mixed-replace` content type. With that content type, the
// server sends a series of parts, each of which is intended to replace the
// previous part. The session calls this method at the beginning of each part,
// followed by one or more calls to [URLSessionDataTaskDidReceiveData] with
// the contents of that part.
// 
// Each time the [URLSessionDataTaskDidReceiveResponseCompletionHandler]
// method is called for a part, collect the data received for the previous
// part (if any) and process the data as needed for your application. This
// processing can include storing the data to the filesystem, parsing it into
// custom types, or displaying it to the user. Next, begin receiving the next
// part by calling the completion handler with the
// [URLSession.ResponseDisposition.allow] constant. Finally, if you have also
// implemented [URLSessionTaskDidCompleteWithError], the session will call it
// after sending all the data for the last part.
//
// [URLSession.ResponseDisposition.allow]: https://developer.apple.com/documentation/Foundation/URLSession/ResponseDisposition/allow
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionDataDelegate/urlSession(_:dataTask:didReceive:completionHandler:)

func (o NSURLSessionDataDelegateObject) URLSessionDataTaskDidReceiveResponseCompletionHandler(session INSURLSession, dataTask INSURLSessionDataTask, response INSURLResponse, completionHandler ErrorHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:dataTask:didReceiveResponse:completionHandler:"), session, dataTask, response, completionHandler)
	}

// Tells the delegate that the data task was changed to a download task.
//
// session: The session containing the task that was replaced by a download task.
//
// dataTask: The data task that was replaced by a download task.
//
// downloadTask: The new download task that replaced the data task.
//
// # Discussion
// 
// When your [URLSessionDataTaskDidReceiveResponseCompletionHandler] delegate
// method uses the [URLSession.ResponseDisposition.becomeDownload] disposition
// to convert the request to use a download, the session calls this delegate
// method to provide you with the new download task. After this call, the
// session delegate receives no further delegate method calls related to the
// original data task.
//
// [URLSession.ResponseDisposition.becomeDownload]: https://developer.apple.com/documentation/Foundation/URLSession/ResponseDisposition/becomeDownload
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionDataDelegate/urlSession(_:dataTask:didBecome:)-60op5

func (o NSURLSessionDataDelegateObject) URLSessionDataTaskDidBecomeDownloadTask(session INSURLSession, dataTask INSURLSessionDataTask, downloadTask INSURLSessionDownloadTask) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:dataTask:didBecomeDownloadTask:"), session, dataTask, downloadTask)
	}

// Tells the delegate that the data task was changed to a stream task.
//
// session: The session containing the task that was replaced by a stream task.
//
// dataTask: The data task that was replaced by a stream task.
//
// streamTask: The new stream task that replaced the data task.
//
// # Discussion
// 
// When your [URLSessionDataTaskDidReceiveResponseCompletionHandler] delegate
// method uses the [URLSession.ResponseDisposition.becomeStream] disposition
// to convert the request to use a stream, the session calls this delegate
// method to provide you with the new stream task. After this call, the
// session delegate receives no further delegate method calls related to the
// original data task.
// 
// For requests that were pipelined, the stream task allows only reading, and
// the object immediately sends the delegate message
// [URLSessionWriteClosedForStreamTask]. You can disable pipelining for all
// requests in a session by setting the [HTTPShouldUsePipelining] property on
// its [NSURLSessionConfiguration] object, or for individual requests by
// setting the [HTTPShouldUsePipelining] property on an [NSURLRequest] object.
//
// [URLSession.ResponseDisposition.becomeStream]: https://developer.apple.com/documentation/Foundation/URLSession/ResponseDisposition/becomeStream
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionDataDelegate/urlSession(_:dataTask:didBecome:)-7nqzu

func (o NSURLSessionDataDelegateObject) URLSessionDataTaskDidBecomeStreamTask(session INSURLSession, dataTask INSURLSessionDataTask, streamTask INSURLSessionStreamTask) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:dataTask:didBecomeStreamTask:"), session, dataTask, streamTask)
	}

// Tells the delegate that the data task has received some of the expected
// data.
//
// session: The session containing the data task that provided data.
//
// dataTask: The data task that provided data.
//
// data: A data object containing the transferred data.
//
// # Discussion
// 
// Because the data object parameter is often pieced together from a number of
// different data objects, whenever possible, use the
// [EnumerateByteRangesUsingBlock] method to iterate through the data rather
// than using the [Bytes] method (which flattens the data object into a single
// memory block).
// 
// This delegate method may be called more than once, and each call provides
// only data received since the previous call. The app is responsible for
// accumulating this data if needed.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionDataDelegate/urlSession(_:dataTask:didReceive:)

func (o NSURLSessionDataDelegateObject) URLSessionDataTaskDidReceiveData(session INSURLSession, dataTask INSURLSessionDataTask, data INSData) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:dataTask:didReceiveData:"), session, dataTask, data)
	}

// Asks the delegate whether the data (or upload) task should store the
// response in the cache.
//
// session: The session containing the data (or upload) task.
//
// dataTask: The data (or upload) task.
//
// proposedResponse: The default caching behavior. This behavior is determined based on the
// current caching policy and the values of certain received headers, such as
// the [Pragma] and `Cache-Control` headers.
//
// completionHandler: A block that your handler must call, providing either the original proposed
// response, a modified version of that response, or [NULL] to prevent caching
// the response. If your delegate implements this method, it call this
// completion handler; otherwise, your app leaks memory.
//
// # Discussion
// 
// The session calls this delegate method after the task finishes receiving
// all of the expected data. If you don’t implement this method, the default
// behavior is to use the caching policy specified in the session’s
// configuration object. The primary purpose of this method is to prevent
// caching of specific URLs or to modify the `userInfo` dictionary associated
// with the URL response.
// 
// This method is called only if the [NSURLProtocol] handling the request
// decides to cache the response. As a rule, responses are cached only when
// all of the following are true:
// 
// - The request is for an HTTP or HTTPS URL (or your own custom networking
// protocol that supports caching). - The request was successful (with a
// status code in the `200–299` range). - The provided response came from
// the server, rather than out of the cache. - The session configuration’s
// cache policy allows caching. - The provided [URLRequest] object’s cache
// policy (if applicable) allows caching. - The cache-related headers in the
// server’s response (if present) allow caching. - The response size is
// small enough to reasonably fit within the cache. (For example, if you
// provide a disk cache, the response must be no larger than about 5% of the
// disk cache size.)
//
// [URLRequest]: https://developer.apple.com/documentation/Foundation/URLRequest
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionDataDelegate/urlSession(_:dataTask:willCacheResponse:completionHandler:)

func (o NSURLSessionDataDelegateObject) URLSessionDataTaskWillCacheResponseCompletionHandler(session INSURLSession, dataTask INSURLSessionDataTask, proposedResponse INSCachedURLResponse, completionHandler CachedURLResponseHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:dataTask:willCacheResponse:completionHandler:"), session, dataTask, proposedResponse, completionHandler)
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

func (o NSURLSessionDataDelegateObject) URLSessionDidBecomeInvalidWithError(session INSURLSession, error_ INSError) {
	
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

func (o NSURLSessionDataDelegateObject) URLSessionDidFinishEventsForBackgroundURLSession(session INSURLSession) {
	
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

func (o NSURLSessionDataDelegateObject) URLSessionDidReceiveChallengeCompletionHandler(session INSURLSession, challenge INSURLAuthenticationChallenge, completionHandler URLSessionAuthChallengeDispositionURLCredentialHandler) {
	
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

func (o NSURLSessionDataDelegateObject) URLSessionTaskDidCompleteWithError(session INSURLSession, task INSURLSessionTask, error_ INSError) {
	
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

func (o NSURLSessionDataDelegateObject) URLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler(session INSURLSession, task INSURLSessionTask, response INSHTTPURLResponse, request INSURLRequest, completionHandler URLRequestHandler) {
	
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

func (o NSURLSessionDataDelegateObject) URLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend(session INSURLSession, task INSURLSessionTask, bytesSent int64, totalBytesSent int64, totalBytesExpectedToSend int64) {
	
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

func (o NSURLSessionDataDelegateObject) URLSessionTaskNeedNewBodyStream(session INSURLSession, task INSURLSessionTask, completionHandler InputStreamHandler) {
	
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

func (o NSURLSessionDataDelegateObject) URLSessionTaskDidReceiveChallengeCompletionHandler(session INSURLSession, task INSURLSessionTask, challenge INSURLAuthenticationChallenge, completionHandler URLSessionAuthChallengeDispositionURLCredentialHandler) {
	
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

func (o NSURLSessionDataDelegateObject) URLSessionTaskWillBeginDelayedRequestCompletionHandler(session INSURLSession, task INSURLSessionTask, request INSURLRequest, completionHandler URLSessionDelayedRequestDispositionURLRequestHandler) {
	
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

func (o NSURLSessionDataDelegateObject) URLSessionTaskIsWaitingForConnectivity(session INSURLSession, task INSURLSessionTask) {
	
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

func (o NSURLSessionDataDelegateObject) URLSessionTaskDidFinishCollectingMetrics(session INSURLSession, task INSURLSessionTask, metrics INSURLSessionTaskMetrics) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:task:didFinishCollectingMetrics:"), session, task, metrics)
	}

//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskDelegate/urlSession(_:didCreateTask:)

func (o NSURLSessionDataDelegateObject) URLSessionDidCreateTask(session INSURLSession, task INSURLSessionTask) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:didCreateTask:"), session, task)
	}

//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskDelegate/urlSession(_:task:didReceiveInformationalResponse:)

func (o NSURLSessionDataDelegateObject) URLSessionTaskDidReceiveInformationalResponse(session INSURLSession, task INSURLSessionTask, response INSHTTPURLResponse) {
	
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

func (o NSURLSessionDataDelegateObject) URLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler(session INSURLSession, task INSURLSessionTask, offset int64, completionHandler InputStreamHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:task:needNewBodyStreamFromOffset:completionHandler:"), session, task, offset, completionHandler)
	}

// NSURLSessionDataDelegateConfig holds optional typed callbacks for [NSURLSessionDataDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsurlsessiondatadelegate
type NSURLSessionDataDelegateConfig struct {

	// Other Methods
	// URLSessionDataTaskDidReceiveResponseCompletionHandler — Tells the delegate that the data task received the initial reply (headers) from the server.
	URLSessionDataTaskDidReceiveResponseCompletionHandler func(session NSURLSession, dataTask NSURLSessionDataTask, response NSURLResponse, completionHandler NSURLSessionResponseDisposition)
	// URLSessionDataTaskDidBecomeDownloadTask — Tells the delegate that the data task was changed to a download task.
	URLSessionDataTaskDidBecomeDownloadTask func(session NSURLSession, dataTask NSURLSessionDataTask, downloadTask NSURLSessionDownloadTask)
	// URLSessionDataTaskDidBecomeStreamTask — Tells the delegate that the data task was changed to a stream task.
	URLSessionDataTaskDidBecomeStreamTask func(session NSURLSession, dataTask NSURLSessionDataTask, streamTask NSURLSessionStreamTask)
	// URLSessionDataTaskDidReceiveData — Tells the delegate that the data task has received some of the expected data.
	URLSessionDataTaskDidReceiveData func(session NSURLSession, dataTask NSURLSessionDataTask, data objectivec.Object)
	// URLSessionDataTaskWillCacheResponseCompletionHandler — Asks the delegate whether the data (or upload) task should store the response in the cache.
	URLSessionDataTaskWillCacheResponseCompletionHandler func(session NSURLSession, dataTask NSURLSessionDataTask, proposedResponse NSCachedURLResponse, completionHandler objc.ID)
}

// NewNSURLSessionDataDelegate creates an Objective-C object implementing the [NSURLSessionDataDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSURLSessionDataDelegateObject] satisfies the [NSURLSessionDataDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsurlsessiondatadelegate
func NewNSURLSessionDataDelegate(config NSURLSessionDataDelegateConfig) NSURLSessionDataDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSURLSessionDataDelegate_%d", n)

	var methods []objc.MethodDef

	if config.URLSessionDataTaskDidReceiveResponseCompletionHandler != nil {
		fn := config.URLSessionDataTaskDidReceiveResponseCompletionHandler
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("URLSession:dataTask:didReceiveResponse:completionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, dataTaskID objc.ID, responseID objc.ID, completionHandler NSURLSessionResponseDisposition) {
				session := NSURLSessionFromID(sessionID)
				dataTask := NSURLSessionDataTaskFromID(dataTaskID)
				response := NSURLResponseFromID(responseID)
				fn(session, dataTask, response, completionHandler)
			},
		})
	}

	if config.URLSessionDataTaskDidBecomeDownloadTask != nil {
		fn := config.URLSessionDataTaskDidBecomeDownloadTask
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("URLSession:dataTask:didBecomeDownloadTask:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, dataTaskID objc.ID, downloadTaskID objc.ID) {
				session := NSURLSessionFromID(sessionID)
				dataTask := NSURLSessionDataTaskFromID(dataTaskID)
				downloadTask := NSURLSessionDownloadTaskFromID(downloadTaskID)
				fn(session, dataTask, downloadTask)
			},
		})
	}

	if config.URLSessionDataTaskDidBecomeStreamTask != nil {
		fn := config.URLSessionDataTaskDidBecomeStreamTask
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("URLSession:dataTask:didBecomeStreamTask:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, dataTaskID objc.ID, streamTaskID objc.ID) {
				session := NSURLSessionFromID(sessionID)
				dataTask := NSURLSessionDataTaskFromID(dataTaskID)
				streamTask := NSURLSessionStreamTaskFromID(streamTaskID)
				fn(session, dataTask, streamTask)
			},
		})
	}

	if config.URLSessionDataTaskDidReceiveData != nil {
		fn := config.URLSessionDataTaskDidReceiveData
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("URLSession:dataTask:didReceiveData:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, dataTaskID objc.ID, dataID objc.ID) {
				session := NSURLSessionFromID(sessionID)
				dataTask := NSURLSessionDataTaskFromID(dataTaskID)
				data := objectivec.ObjectFromID(dataID)
				fn(session, dataTask, data)
			},
		})
	}

	if config.URLSessionDataTaskWillCacheResponseCompletionHandler != nil {
		fn := config.URLSessionDataTaskWillCacheResponseCompletionHandler
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("URLSession:dataTask:willCacheResponse:completionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, dataTaskID objc.ID, proposedResponseID objc.ID, completionHandlerID objc.ID) {
				session := NSURLSessionFromID(sessionID)
				dataTask := NSURLSessionDataTaskFromID(dataTaskID)
				proposedResponse := NSCachedURLResponseFromID(proposedResponseID)
				completionHandler := completionHandlerID
				fn(session, dataTask, proposedResponse, completionHandler)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSURLSessionDataDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSURLSessionDataDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSURLSessionDataDelegateObjectFromID(instance)
}

