// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [URLSession] class.
var (
	_URLSessionClass     URLSessionClass
	_URLSessionClassOnce sync.Once
)

func getURLSessionClass() URLSessionClass {
	_URLSessionClassOnce.Do(func() {
		_URLSessionClass = URLSessionClass{class: objc.GetClass("NSURLSession")}
	})
	return _URLSessionClass
}

// GetURLSessionClass returns the class object for NSURLSession.
func GetURLSessionClass() URLSessionClass {
	return getURLSessionClass()
}

type URLSessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc URLSessionClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc URLSessionClass) Alloc() URLSession {
	rv := objc.Send[URLSession](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// An object that coordinates a group of related, network data transfer tasks.
//
// # Overview
// 
// The [NSURLSession] class and related classes provide an API for downloading
// data from and uploading data to endpoints indicated by URLs. Your app can
// also use this API to perform background downloads when your app isn’t
// running or, in iOS, while your app is suspended. You can use the related
// [NSURLSessionDelegate] and [NSURLSessionTaskDelegate] to support
// authentication and receive events like redirection and task completion.
// 
// Your app creates one or more [NSURLSession] instances, each of which
// coordinates a group of related data-transfer tasks. For example, if
// you’re creating a web browser, your app might create one session per tab
// or window, or one session for interactive use and another for background
// downloads. Within each session, your app adds a series of tasks, each of
// which represents a request for a specific URL (following HTTP redirects, if
// necessary).
// 
// # Types of URL sessions
// 
// The tasks within a given URL session share a common session configuration
// object, which defines connection behavior, like the maximum number of
// simultaneous connections to make to a single host, whether connections can
// use the cellular network, and so on.
// 
// [NSURLSession] has a singleton [SharedSession] session (which doesn’t
// have a configuration object) for basic requests. It’s not as customizable
// as sessions you create, but it serves as a good starting point if you have
// very limited requirements. You access this session by calling the shared
// class method. For other kinds of sessions, you create a [NSURLSession] with
// one of three kinds of configurations:
// 
// - A default session behaves much like the shared session, but lets you
// configure it. You can also assign a delegate to the default session to
// obtain data incrementally. - Ephemeral sessions are similar to shared
// sessions, but don’t write caches, cookies, or credentials to disk. -
// Background sessions let you perform uploads and downloads of content in the
// background while your app isn’t running.
// 
// See Creating a session configuration object in the
// [NSURLSessionConfiguration] class for details on creating each type of
// configuration.
// 
// # Types of URL session tasks
// 
// Within a session, you create tasks that optionally upload data to a server
// and then retrieve data from the server either as a file on disk or as one
// or more [NSData] objects in memory. The [NSURLSession] API provides four
// types of tasks:
// 
// - Data tasks send and receive data using [NSData] objects. Data tasks are
// intended for short, often interactive requests to a server. - Upload tasks
// are similar to data tasks, but they also send data (often in the form of a
// file), and support background uploads while the app isn’t running. -
// Download tasks retrieve data in the form of a file, and support background
// downloads and uploads while the app isn’t running. - WebSocket tasks
// exchange messages over TCP and TLS, using the WebSocket protocol defined in
// [RFC 6455].
// 
// # Using a session delegate
// 
// Tasks in a session also share a common delegate object. You implement this
// delegate to provide and obtain information when various events occur,
// including when:
// 
// - Authentication fails. - Data arrives from the server. - Data becomes
// available for caching.
// 
// If you don’t need the features provided by a delegate, you can use this
// API without providing one by passing `nil` when you create a session.
// 
// Each task you create with the session calls back to the session’s
// delegate, using the methods defined in [NSURLSessionTaskDelegate]. You can
// also intercept these callbacks before they reach the session delegate by
// populating a separate [Delegate] that’s specific to the task.
// 
// # Asynchronicity and URL sessions
// 
// Like most networking APIs, the [NSURLSession] API is highly asynchronous.
// It returns data to your app in one of three ways, depending on the methods
// you call:
// 
// - If you’re using Swift, you can use the methods marked with the `async`
// keyword to perform common tasks. For example, [data(from:delegate:)]
// fetches data, while [download(from:delegate:)] downloads files. Your call
// point uses the `await` keyword to suspend running until the transfer
// completes. You can also use the [bytes(from:delegate:)] method to receive
// data as an [AsyncSequence]. With this approach, you use the
// `for`-`await`-`in` syntax to iterate over the data as your app receives it.
// The [URL] type also offers covenience methods to fetch bytes or lines from
// the shared URL session. - In Swift or Objective-C, you can provide a
// completion handler block, which runs when the transfer completes. - In
// Swift or Objective-C, you can receive callbacks to a delegate method as the
// transfer progresses and immediately after it completes.
// 
// In addition to delivering this information to delegates, the [NSURLSession]
// provides status and progress properties. Query these properties if you need
// to make programmatic decisions based on the current state of the task (with
// the caveat that its state can change at any time).
// 
// # Protocol support
// 
// The [NSURLSession] class natively supports the `data`, `file`, `ftp`,
// `http`, and `https` URL schemes, with transparent support for proxy servers
// and SOCKS gateways, as configured in the user’s system preferences.
// 
// [NSURLSession] supports the HTTP/1.1, HTTP/2, and HTTP/3 protocols. HTTP/2
// support, as described by [RFC 7540], requires a server that supports
// Application-Layer Protocol Negotiation (ALPN).
// 
// You can also add support for your own custom networking protocols and URL
// schemes (for your app’s private use) by subclassing [NSURLProtocol].
// 
// # App Transport Security (ATS)
// 
// iOS 9.0 and macOS 10.11 and later use App Transport Security (ATS) for all
// HTTP connections made with [NSURLSession]. ATS requires that HTTP
// connections use HTTPS ([RFC 2818]).
// 
// For more information, see [NSAppTransportSecurity].
// 
// # Foundation copying behavior
// 
// Session and task objects conform to the [NSCopying] protocol as follows:
// 
// - When your app copies a session or task object, you get the same object
// back. - When your app copies a configuration object, you get a new copy you
// can independently modify.
// 
// # Thread safety
// 
// The URL session API is thread-safe. You can freely create sessions and
// tasks in any thread context. When your delegate methods call the provided
// completion handlers, the work is automatically scheduled on the correct
// delegate queue.
//
// [AsyncSequence]: https://developer.apple.com/documentation/Swift/AsyncSequence
// [NSAppTransportSecurity]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/NSAppTransportSecurity
// [RFC 2818]: https://tools.ietf.org/html/rfc2818
// [RFC 6455]: https://tools.ietf.org/html/rfc6455
// [RFC 7540]: https://tools.ietf.org/html/rfc7540
// [URL]: https://developer.apple.com/documentation/Foundation/URL
// [bytes(from:delegate:)]: https://developer.apple.com/documentation/Foundation/URLSession/bytes(from:delegate:)
// [data(from:delegate:)]: https://developer.apple.com/documentation/Foundation/URLSession/data(from:delegate:)
// [download(from:delegate:)]: https://developer.apple.com/documentation/Foundation/URLSession/download(from:delegate:)
//
// # Creating a session
//
//   - [URLSession.Configuration]: A copy of the configuration object for this session.
//
// # Working with a delegate
//
//   - [URLSession.Delegate]: The delegate assigned when this object was created.
//   - [URLSession.DelegateQueue]: The operation queue provided when this object was created.
//
// # Adding data tasks to a session
//
//   - [URLSession.DataTaskWithURL]: Creates a task that retrieves the contents of the specified URL.
//   - [URLSession.DataTaskWithURLCompletionHandler]: Creates a task that retrieves the contents of the specified URL, then calls a handler upon completion.
//   - [URLSession.DataTaskWithRequest]: Creates a task that retrieves the contents of a URL based on the specified URL request object.
//   - [URLSession.DataTaskWithRequestCompletionHandler]: Creates a task that retrieves the contents of a URL based on the specified URL request object, and calls a handler upon completion.
//
// # Adding download tasks to a session
//
//   - [URLSession.DownloadTaskWithURL]: Creates a download task that retrieves the contents of the specified URL and saves the results to a file.
//   - [URLSession.DownloadTaskWithURLCompletionHandler]: Creates a download task that retrieves the contents of the specified URL, saves the results to a file, and calls a handler upon completion.
//   - [URLSession.DownloadTaskWithRequest]: Creates a download task that retrieves the contents of a URL based on the specified URL request object and saves the results to a file.
//   - [URLSession.DownloadTaskWithRequestCompletionHandler]: Creates a download task that retrieves the contents of a URL based on the specified URL request object, saves the results to a file, and calls a handler upon completion.
//   - [URLSession.DownloadTaskWithResumeData]: Creates a download task to resume a previously canceled or failed download.
//   - [URLSession.DownloadTaskWithResumeDataCompletionHandler]: Creates a download task to resume a previously canceled or failed download and calls a handler upon completion.
//
// # Adding upload tasks to a session
//
//   - [URLSession.UploadTaskWithRequestFromData]: Creates a task that performs an HTTP request for the specified URL request object and uploads the provided data.
//   - [URLSession.UploadTaskWithRequestFromDataCompletionHandler]: Creates a task that performs an HTTP request for the specified URL request object, uploads the provided data, and calls a handler upon completion.
//   - [URLSession.UploadTaskWithRequestFromFile]: Creates a task that performs an HTTP request for uploading the specified file.
//   - [URLSession.UploadTaskWithRequestFromFileCompletionHandler]: Creates a task that performs an HTTP request for uploading the specified file, then calls a handler upon completion.
//   - [URLSession.UploadTaskWithStreamedRequest]: Creates a task that performs an HTTP request for uploading data based on the specified URL request.
//   - [URLSession.UploadTaskWithResumeData]: Creates an upload task from a resume data blob. Requires the server to support the latest resumable uploads Internet-Draft from the HTTP Working Group, found at https://datatracker.ietf.org/doc/draft-ietf-httpbis-resumable-upload/ If resuming from an upload file, the file must still exist and be unmodified. If the upload cannot be successfully resumed, URLSession:task:didCompleteWithError: will be called.
//   - [URLSession.UploadTaskWithResumeDataCompletionHandler]: Creates a URLSessionUploadTask from a resume data blob. If resuming from an upload file, the file must still exist and be unmodified.
//
// # Adding stream tasks to a session
//
//   - [URLSession.StreamTaskWithHostNamePort]: Creates a task that establishes a bidirectional TCP/IP connection to a specified hostname and port.
//
// # Adding WebSocket tasks to a session
//
//   - [URLSession.WebSocketTaskWithURL]: Creates a WebSocket task for the provided URL.
//   - [URLSession.WebSocketTaskWithRequest]: Creates a WebSocket task for the provided URL request.
//   - [URLSession.WebSocketTaskWithURLProtocols]: Creates a WebSocket task given a URL and an array of protocols.
//
// # Managing the session
//
//   - [URLSession.FinishTasksAndInvalidate]: Invalidates the session, allowing any outstanding tasks to finish.
//   - [URLSession.FlushWithCompletionHandler]: Flushes cookies and credentials to disk, clears transient caches, and ensures that future requests occur on a new TCP connection.
//   - [URLSession.GetTasksWithCompletionHandler]: Asynchronously calls a completion callback with all data, upload, and download tasks in a session.
//   - [URLSession.InvalidateAndCancel]: Cancels all outstanding tasks and then invalidates the session.
//   - [URLSession.ResetWithCompletionHandler]: Empties all cookies, caches and credential stores, removes disk files, flushes in-progress downloads to disk, and ensures that future requests occur on a new socket.
//   - [URLSession.SessionDescription]: An app-defined descriptive label for the session.
//   - [URLSession.SetSessionDescription]
//
// See: https://developer.apple.com/documentation/Foundation/URLSession
type URLSession struct {
	objectivec.Object
}

// URLSessionFromID constructs a [URLSession] from an objc.ID.
//
// An object that coordinates a group of related, network data transfer tasks.
func URLSessionFromID(id objc.ID) URLSession {
	return NSURLSession{objectivec.Object{ID: id}}
}

// NSURLSessionFromID is an alias for [URLSessionFromID] for cross-framework compatibility.
func NSURLSessionFromID(id objc.ID) URLSession { return URLSessionFromID(id) }
// NOTE: URLSession adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [URLSession] class.
//
// # Creating a session
//
//   - [IURLSession.Configuration]: A copy of the configuration object for this session.
//
// # Working with a delegate
//
//   - [IURLSession.Delegate]: The delegate assigned when this object was created.
//   - [IURLSession.DelegateQueue]: The operation queue provided when this object was created.
//
// # Adding data tasks to a session
//
//   - [IURLSession.DataTaskWithURL]: Creates a task that retrieves the contents of the specified URL.
//   - [IURLSession.DataTaskWithURLCompletionHandler]: Creates a task that retrieves the contents of the specified URL, then calls a handler upon completion.
//   - [IURLSession.DataTaskWithRequest]: Creates a task that retrieves the contents of a URL based on the specified URL request object.
//   - [IURLSession.DataTaskWithRequestCompletionHandler]: Creates a task that retrieves the contents of a URL based on the specified URL request object, and calls a handler upon completion.
//
// # Adding download tasks to a session
//
//   - [IURLSession.DownloadTaskWithURL]: Creates a download task that retrieves the contents of the specified URL and saves the results to a file.
//   - [IURLSession.DownloadTaskWithURLCompletionHandler]: Creates a download task that retrieves the contents of the specified URL, saves the results to a file, and calls a handler upon completion.
//   - [IURLSession.DownloadTaskWithRequest]: Creates a download task that retrieves the contents of a URL based on the specified URL request object and saves the results to a file.
//   - [IURLSession.DownloadTaskWithRequestCompletionHandler]: Creates a download task that retrieves the contents of a URL based on the specified URL request object, saves the results to a file, and calls a handler upon completion.
//   - [IURLSession.DownloadTaskWithResumeData]: Creates a download task to resume a previously canceled or failed download.
//   - [IURLSession.DownloadTaskWithResumeDataCompletionHandler]: Creates a download task to resume a previously canceled or failed download and calls a handler upon completion.
//
// # Adding upload tasks to a session
//
//   - [IURLSession.UploadTaskWithRequestFromData]: Creates a task that performs an HTTP request for the specified URL request object and uploads the provided data.
//   - [IURLSession.UploadTaskWithRequestFromDataCompletionHandler]: Creates a task that performs an HTTP request for the specified URL request object, uploads the provided data, and calls a handler upon completion.
//   - [IURLSession.UploadTaskWithRequestFromFile]: Creates a task that performs an HTTP request for uploading the specified file.
//   - [IURLSession.UploadTaskWithRequestFromFileCompletionHandler]: Creates a task that performs an HTTP request for uploading the specified file, then calls a handler upon completion.
//   - [IURLSession.UploadTaskWithStreamedRequest]: Creates a task that performs an HTTP request for uploading data based on the specified URL request.
//   - [IURLSession.UploadTaskWithResumeData]: Creates an upload task from a resume data blob. Requires the server to support the latest resumable uploads Internet-Draft from the HTTP Working Group, found at https://datatracker.ietf.org/doc/draft-ietf-httpbis-resumable-upload/ If resuming from an upload file, the file must still exist and be unmodified. If the upload cannot be successfully resumed, URLSession:task:didCompleteWithError: will be called.
//   - [IURLSession.UploadTaskWithResumeDataCompletionHandler]: Creates a URLSessionUploadTask from a resume data blob. If resuming from an upload file, the file must still exist and be unmodified.
//
// # Adding stream tasks to a session
//
//   - [IURLSession.StreamTaskWithHostNamePort]: Creates a task that establishes a bidirectional TCP/IP connection to a specified hostname and port.
//
// # Adding WebSocket tasks to a session
//
//   - [IURLSession.WebSocketTaskWithURL]: Creates a WebSocket task for the provided URL.
//   - [IURLSession.WebSocketTaskWithRequest]: Creates a WebSocket task for the provided URL request.
//   - [IURLSession.WebSocketTaskWithURLProtocols]: Creates a WebSocket task given a URL and an array of protocols.
//
// # Managing the session
//
//   - [IURLSession.FinishTasksAndInvalidate]: Invalidates the session, allowing any outstanding tasks to finish.
//   - [IURLSession.FlushWithCompletionHandler]: Flushes cookies and credentials to disk, clears transient caches, and ensures that future requests occur on a new TCP connection.
//   - [IURLSession.GetTasksWithCompletionHandler]: Asynchronously calls a completion callback with all data, upload, and download tasks in a session.
//   - [IURLSession.InvalidateAndCancel]: Cancels all outstanding tasks and then invalidates the session.
//   - [IURLSession.ResetWithCompletionHandler]: Empties all cookies, caches and credential stores, removes disk files, flushes in-progress downloads to disk, and ensures that future requests occur on a new socket.
//   - [IURLSession.SessionDescription]: An app-defined descriptive label for the session.
//   - [IURLSession.SetSessionDescription]
//
// See: https://developer.apple.com/documentation/Foundation/URLSession
type IURLSession interface {
	objectivec.IObject

	// Topic: Creating a session

	// A copy of the configuration object for this session.
	Configuration() INSURLSessionConfiguration

	// Topic: Working with a delegate

	// The delegate assigned when this object was created.
	Delegate() NSURLSessionDelegate
	// The operation queue provided when this object was created.
	DelegateQueue() INSOperationQueue

	// Topic: Adding data tasks to a session

	// Creates a task that retrieves the contents of the specified URL.
	DataTaskWithURL(url INSURL) INSURLSessionDataTask
	// Creates a task that retrieves the contents of the specified URL, then calls a handler upon completion.
	DataTaskWithURLCompletionHandler(url INSURL, completionHandler DataURLResponseErrorHandler) INSURLSessionDataTask
	// Creates a task that retrieves the contents of a URL based on the specified URL request object.
	DataTaskWithRequest(request INSURLRequest) INSURLSessionDataTask
	// Creates a task that retrieves the contents of a URL based on the specified URL request object, and calls a handler upon completion.
	DataTaskWithRequestCompletionHandler(request INSURLRequest, completionHandler DataURLResponseErrorHandler) INSURLSessionDataTask

	// Topic: Adding download tasks to a session

	// Creates a download task that retrieves the contents of the specified URL and saves the results to a file.
	DownloadTaskWithURL(url INSURL) INSURLSessionDownloadTask
	// Creates a download task that retrieves the contents of the specified URL, saves the results to a file, and calls a handler upon completion.
	DownloadTaskWithURLCompletionHandler(url INSURL, completionHandler URLURLResponseErrorHandler) INSURLSessionDownloadTask
	// Creates a download task that retrieves the contents of a URL based on the specified URL request object and saves the results to a file.
	DownloadTaskWithRequest(request INSURLRequest) INSURLSessionDownloadTask
	// Creates a download task that retrieves the contents of a URL based on the specified URL request object, saves the results to a file, and calls a handler upon completion.
	DownloadTaskWithRequestCompletionHandler(request INSURLRequest, completionHandler URLURLResponseErrorHandler) INSURLSessionDownloadTask
	// Creates a download task to resume a previously canceled or failed download.
	DownloadTaskWithResumeData(resumeData INSData) INSURLSessionDownloadTask
	// Creates a download task to resume a previously canceled or failed download and calls a handler upon completion.
	DownloadTaskWithResumeDataCompletionHandler(resumeData INSData, completionHandler URLURLResponseErrorHandler) INSURLSessionDownloadTask

	// Topic: Adding upload tasks to a session

	// Creates a task that performs an HTTP request for the specified URL request object and uploads the provided data.
	UploadTaskWithRequestFromData(request INSURLRequest, bodyData INSData) INSURLSessionUploadTask
	// Creates a task that performs an HTTP request for the specified URL request object, uploads the provided data, and calls a handler upon completion.
	UploadTaskWithRequestFromDataCompletionHandler(request INSURLRequest, bodyData INSData, completionHandler DataURLResponseErrorHandler) INSURLSessionUploadTask
	// Creates a task that performs an HTTP request for uploading the specified file.
	UploadTaskWithRequestFromFile(request INSURLRequest, fileURL INSURL) INSURLSessionUploadTask
	// Creates a task that performs an HTTP request for uploading the specified file, then calls a handler upon completion.
	UploadTaskWithRequestFromFileCompletionHandler(request INSURLRequest, fileURL INSURL, completionHandler DataURLResponseErrorHandler) INSURLSessionUploadTask
	// Creates a task that performs an HTTP request for uploading data based on the specified URL request.
	UploadTaskWithStreamedRequest(request INSURLRequest) INSURLSessionUploadTask
	// Creates an upload task from a resume data blob. Requires the server to support the latest resumable uploads Internet-Draft from the HTTP Working Group, found at https://datatracker.ietf.org/doc/draft-ietf-httpbis-resumable-upload/ If resuming from an upload file, the file must still exist and be unmodified. If the upload cannot be successfully resumed, URLSession:task:didCompleteWithError: will be called.
	UploadTaskWithResumeData(resumeData INSData) INSURLSessionUploadTask
	// Creates a URLSessionUploadTask from a resume data blob. If resuming from an upload file, the file must still exist and be unmodified.
	UploadTaskWithResumeDataCompletionHandler(resumeData INSData, completionHandler DataURLResponseErrorHandler) INSURLSessionUploadTask

	// Topic: Adding stream tasks to a session

	// Creates a task that establishes a bidirectional TCP/IP connection to a specified hostname and port.
	StreamTaskWithHostNamePort(hostname string, port int) INSURLSessionStreamTask

	// Topic: Adding WebSocket tasks to a session

	// Creates a WebSocket task for the provided URL.
	WebSocketTaskWithURL(url INSURL) INSURLSessionWebSocketTask
	// Creates a WebSocket task for the provided URL request.
	WebSocketTaskWithRequest(request INSURLRequest) INSURLSessionWebSocketTask
	// Creates a WebSocket task given a URL and an array of protocols.
	WebSocketTaskWithURLProtocols(url INSURL, protocols []string) INSURLSessionWebSocketTask

	// Topic: Managing the session

	// Invalidates the session, allowing any outstanding tasks to finish.
	FinishTasksAndInvalidate()
	// Flushes cookies and credentials to disk, clears transient caches, and ensures that future requests occur on a new TCP connection.
	FlushWithCompletionHandler(completionHandler VoidHandler)
	// Asynchronously calls a completion callback with all data, upload, and download tasks in a session.
	GetTasksWithCompletionHandler(completionHandler VoidHandler)
	// Cancels all outstanding tasks and then invalidates the session.
	InvalidateAndCancel()
	// Empties all cookies, caches and credential stores, removes disk files, flushes in-progress downloads to disk, and ensures that future requests occur on a new socket.
	ResetWithCompletionHandler(completionHandler VoidHandler)
	// An app-defined descriptive label for the session.
	SessionDescription() string
	SetSessionDescription(value string)
}

// Init initializes the instance.
func (u URLSession) Init() URLSession {
	rv := objc.Send[URLSession](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u URLSession) Autorelease() URLSession {
	rv := objc.Send[URLSession](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSession creates a new URLSession instance.
func NewURLSession() URLSession {
	class := getURLSessionClass()
	rv := objc.Send[URLSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a session with the specified session configuration.
//
// configuration: A configuration object that specifies certain behaviors, such as caching
// policies, timeouts, proxies, pipelining, TLS versions to support, cookie
// policies, credential storage, and so on.
// 
// See [NSURLSessionConfiguration] for more information.
//
// # Discussion
// 
// Calling this method is equivalent to calling
// [SessionWithConfigurationDelegateDelegateQueue] with a `nil` delegate and
// queue.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/init(configuration:)
func NewURLSessionWithConfiguration(configuration INSURLSessionConfiguration) URLSession {
	rv := objc.Send[objc.ID](objc.ID(getURLSessionClass().class), objc.Sel("sessionWithConfiguration:"), configuration)
	return URLSessionFromID(rv)
}

// Creates a session with the specified session configuration, delegate, and
// operation queue.
//
// configuration: A configuration object that specifies certain behaviors, such as caching
// policies, timeouts, proxies, pipelining, TLS versions to support, cookie
// policies, and credential storage.
// 
// See [NSURLSessionConfiguration] for more information.
//
// delegate: A session delegate object that handles requests for authentication and
// other session-related events.
// 
// This delegate object is responsible for handling authentication challenges,
// for making caching decisions, and for handling other session-related
// events. If `nil`, the class should be used only with methods that take
// completion handlers.
//
// queue: An operation queue for scheduling the delegate calls and completion
// handlers. The queue should be a serial queue, in order to ensure the
// correct ordering of callbacks. If `nil`, the session creates a serial
// operation queue for performing all delegate method calls and completion
// handler calls.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/init(configuration:delegate:delegateQueue:)
func NewURLSessionWithConfigurationDelegateDelegateQueue(configuration INSURLSessionConfiguration, delegate NSURLSessionDelegate, queue INSOperationQueue) URLSession {
	rv := objc.Send[objc.ID](objc.ID(getURLSessionClass().class), objc.Sel("sessionWithConfiguration:delegate:delegateQueue:"), configuration, delegate, queue)
	return URLSessionFromID(rv)
}

// Creates a task that retrieves the contents of the specified URL.
//
// url: The URL to be retrieved.
//
// # Return Value
// 
// The new session data task.
//
// # Discussion
// 
// After you create the task, you must start it by calling its [Resume]
// method. The task calls methods on the session’s delegate to provide you
// with the response metadata, response data, and so on.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/dataTask(with:)-10dy7
func (u URLSession) DataTaskWithURL(url INSURL) INSURLSessionDataTask {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("dataTaskWithURL:"), url)
	return NSURLSessionDataTaskFromID(rv)
}
// Creates a task that retrieves the contents of the specified URL, then calls
// a handler upon completion.
//
// url: The URL to be retrieved.
//
// completionHandler: The completion handler to call when the load request is complete. This
// handler is executed on the delegate queue.
// 
// If you pass `nil`, only the session delegate methods are called when the
// task completes, making this method equivalent to the [DataTaskWithRequest]
// method.
// 
// This completion handler takes the following parameters:
// 
// `data`: The data returned by the server. `response`: An object that
// provides response metadata, such as HTTP headers and status code. If you
// are making an HTTP or HTTPS request, the returned object is actually an
// [NSHTTPURLResponse] object. `error`: An error object that indicates why the
// request failed, or `nil` if the request was successful.
//
// # Return Value
// 
// The new session data task.
//
// # Discussion
// 
// After you create the task, you must start it by calling its [Resume]
// method.
// 
// By using the completion handler, the task bypasses calls to delegate
// methods for response and data delivery, and instead provides any resulting
// [NSData], [NSURLResponse], and [NSError] objects inside the completion
// handler. Delegate methods for handling authentication challenges, however,
// are still called.
// 
// You should pass a `nil` completion handler when creating tasks in sessions
// whose delegates include a [URLSessionDataTaskDidReceiveData] method.
// 
// If the request completes successfully, the `data` parameter of the
// completion handler block contains the resource data, and the `error`
// parameter is `nil`. If the request fails, the `data` parameter is `nil` and
// the `error` parameter contain information about the failure. If a response
// from the server is received, regardless of whether the request completes
// successfully or fails, the `response` parameter contains that information.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/dataTask(with:completionHandler:)-52wk8
func (u URLSession) DataTaskWithURLCompletionHandler(url INSURL, completionHandler DataURLResponseErrorHandler) INSURLSessionDataTask {
_block1, _ := NewDataURLResponseErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](u.ID, objc.Sel("dataTaskWithURL:completionHandler:"), url, _block1)
	return NSURLSessionDataTaskFromID(rv)
}
// Creates a task that retrieves the contents of a URL based on the specified
// URL request object.
//
// request: A URL request object that provides request-specific information such as the
// URL, cache policy, request type, and body data or body stream.
//
// # Return Value
// 
// The new session data task.
//
// # Discussion
// 
// By creating a task based on a request object, you can tune various aspects
// of the task’s behavior, including the cache policy and timeout interval.
// 
// After you create the task, you must start it by calling its [Resume]
// method.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/dataTask(with:)-7jpys
func (u URLSession) DataTaskWithRequest(request INSURLRequest) INSURLSessionDataTask {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("dataTaskWithRequest:"), request)
	return NSURLSessionDataTaskFromID(rv)
}
// Creates a task that retrieves the contents of a URL based on the specified
// URL request object, and calls a handler upon completion.
//
// request: A URL request object that provides the URL, cache policy, request type,
// body data or body stream, and so on.
//
// completionHandler: The completion handler to call when the load request is complete. This
// handler is executed on the delegate queue.
// 
// If you pass `nil`, only the session delegate methods are called when the
// task completes, making this method equivalent to the [DataTaskWithRequest]
// method.
// 
// This completion handler takes the following parameters:
// 
// `data`: The data returned by the server. `response`: An object that
// provides response metadata, such as HTTP headers and status code. If you
// are making an HTTP or HTTPS request, the returned object is actually an
// [NSHTTPURLResponse] object. `error`: An error object that indicates why the
// request failed, or `nil` if the request was successful.
//
// # Return Value
// 
// The new session data task.
//
// # Discussion
// 
// By creating a task based on a request object, you can tune various aspects
// of the task’s behavior, including the cache policy and timeout interval.
// 
// By using the completion handler, the task bypasses calls to delegate
// methods for response and data delivery, and instead provides any resulting
// [NSData], [NSURLResponse], and [NSError] objects inside the completion
// handler. Delegate methods for handling authentication challenges, however,
// are still called.
// 
// You should pass a `nil` completion handler when creating tasks in sessions
// whose delegates include a [URLSessionDataTaskDidReceiveData] method.
// 
// After you create the task, you must start it by calling its [Resume]
// method.
// 
// If the request completes successfully, the `data` parameter of the
// completion handler block contains the resource data, and the `error`
// parameter is `nil`. If the request fails, the `data` parameter is `nil` and
// the `error` parameter contain information about the failure. If a response
// from the server is received, regardless of whether the request completes
// successfully or fails, the `response` parameter contains that information.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/dataTask(with:completionHandler:)-e6xv
func (u URLSession) DataTaskWithRequestCompletionHandler(request INSURLRequest, completionHandler DataURLResponseErrorHandler) INSURLSessionDataTask {
_block1, _ := NewDataURLResponseErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](u.ID, objc.Sel("dataTaskWithRequest:completionHandler:"), request, _block1)
	return NSURLSessionDataTaskFromID(rv)
}
// Creates a download task that retrieves the contents of the specified URL
// and saves the results to a file.
//
// url: The URL to download.
//
// # Return Value
// 
// The new session download task.
//
// # Discussion
// 
// After you create the task, you must start it by calling its [Resume]
// method.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/downloadTask(with:)-1onj
func (u URLSession) DownloadTaskWithURL(url INSURL) INSURLSessionDownloadTask {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("downloadTaskWithURL:"), url)
	return NSURLSessionDownloadTaskFromID(rv)
}
// Creates a download task that retrieves the contents of the specified URL,
// saves the results to a file, and calls a handler upon completion.
//
// url: The URL to download.
//
// completionHandler: The completion handler to call when the load request is complete. This
// handler is executed on the delegate queue.
// 
// If you pass `nil`, only the session delegate methods are called when the
// task completes, making this method equivalent to the [DownloadTaskWithURL]
// method.
// 
// This completion handler takes the following parameters:
// 
// `location`: The location of a temporary file where the server’s response
// is stored. You must move this file or open it for reading before your
// completion handler returns. Otherwise, the file is deleted, and the data is
// lost. `response`: An object that provides response metadata, such as HTTP
// headers and status code. If you are making an HTTP or HTTPS request, the
// returned object is actually an [NSHTTPURLResponse] object. `error`: An
// error object that indicates why the request failed, or `nil` if the request
// was successful.
//
// # Return Value
// 
// The new session download task.
//
// # Discussion
// 
// By using the completion handler, the task bypasses calls to delegate
// methods for response and data delivery, and instead provides any resulting
// [NSData], [NSURLResponse], and [NSError] objects inside the completion
// handler. Delegate methods for handling authentication challenges, however,
// are still called.
// 
// You should pass a `nil` completion handler when creating tasks in sessions
// whose delegates include a [URLSessionDownloadTaskDidFinishDownloadingToURL]
// method.
// 
// After you create the task, you must start it by calling its [Resume]
// method.
// 
// If the request completes successfully, the `location` parameter of the
// completion handler block contains the location of the temporary file, and
// the `error` parameter is `nil`. If the request fails, the `location`
// parameter is `nil` and the `error` parameter contain information about the
// failure. If a response from the server is received, regardless of whether
// the request completes successfully or fails, the `response` parameter
// contains that information.
//
// [NSData]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/PropertyLists/OldStylePlists/OldStylePLists.html#//apple_ref/doc/uid/20001012-47169
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/downloadTask(with:completionHandler:)-7cuje
func (u URLSession) DownloadTaskWithURLCompletionHandler(url INSURL, completionHandler URLURLResponseErrorHandler) INSURLSessionDownloadTask {
_block1, _ := NewURLURLResponseErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](u.ID, objc.Sel("downloadTaskWithURL:completionHandler:"), url, _block1)
	return NSURLSessionDownloadTaskFromID(rv)
}
// Creates a download task that retrieves the contents of a URL based on the
// specified URL request object and saves the results to a file.
//
// request: A URL request object that provides the URL, cache policy, request type,
// body data or body stream, and so on.
//
// # Return Value
// 
// The new session download task.
//
// # Discussion
// 
// By creating a task based on a request object, you can tune various aspects
// of the task’s behavior, including the cache policy and timeout interval.
// 
// After you create the task, you must start it by calling its [Resume]
// method. The task calls methods on the session’s delegate to provide you
// with progress notifications, the location of the resulting temporary file,
// and so on.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/downloadTask(with:)-3fb7s
func (u URLSession) DownloadTaskWithRequest(request INSURLRequest) INSURLSessionDownloadTask {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("downloadTaskWithRequest:"), request)
	return NSURLSessionDownloadTaskFromID(rv)
}
// Creates a download task that retrieves the contents of a URL based on the
// specified URL request object, saves the results to a file, and calls a
// handler upon completion.
//
// request: A URL request object that provides the URL, cache policy, request type,
// body data or body stream, and so on.
//
// completionHandler: The completion handler to call when the load request is complete. This
// handler is executed on the delegate queue.
// 
// If you pass `nil`, only the session delegate methods are called when the
// task completes, making this method equivalent to the
// [DownloadTaskWithRequest] method.
// 
// `location`: The location of a temporary file where the server’s response
// is stored. You must move this file or open it for reading before your
// completion handler returns. Otherwise, the file is deleted, and the data is
// lost. `response`: An object that provides response metadata, such as HTTP
// headers and status code. If you are making an HTTP or HTTPS request, the
// returned object is actually an [NSHTTPURLResponse] object. `error`: An
// error object that indicates why the request failed, or `nil` if the request
// was successful.
//
// # Return Value
// 
// The new session download task.
//
// # Discussion
// 
// By creating a task based on a request object, you can tune various aspects
// of the task’s behavior, including the cache policy and timeout interval.
// 
// By using a completion handler, the task bypasses calls to delegate methods
// for response and data delivery, and instead provides any resulting [NSURL],
// [NSURLResponse], and [NSError] objects inside the completion handler.
// Delegate methods for handling authentication challenges, however, are still
// called.
// 
// You should pass a `nil` completion handler when creating tasks in sessions
// whose delegates include a [URLSessionDownloadTaskDidFinishDownloadingToURL]
// method.
// 
// After you create the task, you must start it by calling its [Resume]
// method.
// 
// If the request completes successfully, the `location` parameter of the
// completion handler block contains the location of the temporary file, and
// the `error` parameter is `nil`. If the request fails, the `location`
// parameter is `nil` and the `error` parameter contain information about the
// failure. If a response from the server is received, regardless of whether
// the request completes successfully or fails, the `response` parameter
// contains that information.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/downloadTask(with:completionHandler:)-4a84s
func (u URLSession) DownloadTaskWithRequestCompletionHandler(request INSURLRequest, completionHandler URLURLResponseErrorHandler) INSURLSessionDownloadTask {
_block1, _ := NewURLURLResponseErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](u.ID, objc.Sel("downloadTaskWithRequest:completionHandler:"), request, _block1)
	return NSURLSessionDownloadTaskFromID(rv)
}
// Creates a download task to resume a previously canceled or failed download.
//
// resumeData: A data object that provides the data necessary to resume a download.
//
// # Return Value
// 
// The new session download task.
//
// # Discussion
// 
// After you create the task, you must start it by calling its [Resume]
// method.
// 
// This method is equivalent to the
// [DownloadTaskWithResumeDataCompletionHandler] with a `nil` completion
// handler. For detailed usage information, including ways to obtain a resume
// data object, see that method.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/downloadTask(withResumeData:)
func (u URLSession) DownloadTaskWithResumeData(resumeData INSData) INSURLSessionDownloadTask {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("downloadTaskWithResumeData:"), resumeData)
	return NSURLSessionDownloadTaskFromID(rv)
}
// Creates a download task to resume a previously canceled or failed download
// and calls a handler upon completion.
//
// resumeData: A data object that provides the data necessary to resume the download.
//
// completionHandler: The completion handler to call when the load request is complete. This
// handler is executed on the delegate queue.
// 
// If you pass `nil`, only the session delegate methods are called when the
// task completes, making this method equivalent to the
// [DownloadTaskWithResumeData] method.
// 
// `location`: The location of a temporary file where the server’s response
// is stored. You must move this file or open it for reading before your
// completion handler returns. Otherwise, the file is deleted, and the data is
// lost. `response`: An object that provides response metadata, such as HTTP
// headers and status code. If you are making an HTTP or HTTPS request, the
// returned object is actually an [NSHTTPURLResponse] object. `error`: An
// error object that indicates why the request failed, or `nil` if the request
// was successful.
//
// # Return Value
// 
// The new session download task.
//
// # Discussion
// 
// By using a completion handler, the task bypasses calls to delegate methods
// for response and data delivery, and instead provides any resulting data,
// response, or error inside the completion handler. Delegate methods for
// handling authentication challenges, however, are still called.
// 
// You should pass a `nil` completion handler when creating tasks in sessions
// whose delegates include a [URLSessionDownloadTaskDidFinishDownloadingToURL]
// method.
// 
// Your app can obtain a `resumeData` object in two ways:
// 
// - If your app cancels an existing transfer by calling
// [CancelByProducingResumeData], the session object passes a `resumeData`
// object to the completion handler that you provided in that call. - If a
// transfer fails, the session object provides an [NSError] object either to
// its delegate or to the task’s completion handler. In that object, the
// [NSURLSessionDownloadTaskResumeData] key in the `userInfo` dictionary
// contains a `resumeData` object.
// 
// After you create the task, you must start it by calling its [Resume]
// method.
// 
// If the request completes successfully, the `location` parameter of the
// completion handler block contains the location of the temporary file, and
// the `error` parameter is `nil`. If the request fails, the `location`
// parameter is `nil` and the `error` parameter contain information about the
// failure. If a response from the server is received, regardless of whether
// the request completes successfully or fails, the `response` parameter
// contains that information.
//
// [NSURLSessionDownloadTaskResumeData]: https://developer.apple.com/documentation/Foundation/NSURLSessionDownloadTaskResumeData
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/downloadTask(withResumeData:completionHandler:)
func (u URLSession) DownloadTaskWithResumeDataCompletionHandler(resumeData INSData, completionHandler URLURLResponseErrorHandler) INSURLSessionDownloadTask {
_block1, _ := NewURLURLResponseErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](u.ID, objc.Sel("downloadTaskWithResumeData:completionHandler:"), resumeData, _block1)
	return NSURLSessionDownloadTaskFromID(rv)
}
// Creates a task that performs an HTTP request for the specified URL request
// object and uploads the provided data.
//
// request: A URL request object that provides the URL, cache policy, request type, and
// so on. The body stream and body data in this request object are ignored.
//
// bodyData: The body data for the request.
//
// # Return Value
// 
// The new session upload task.
//
// # Discussion
// 
// An HTTP upload request is any request that contains a request body, such as
// a [POST] or [PUT] request. Upload tasks require you to create a request
// object so that you can provide metadata for the upload, like HTTP request
// headers.
// 
// After you create the task, you must start it by calling its [Resume]
// method. The task calls methods on the session’s delegate to provide you
// with the upload’s progress, response metadata, response data, and so on.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(with:from:)
func (u URLSession) UploadTaskWithRequestFromData(request INSURLRequest, bodyData INSData) INSURLSessionUploadTask {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("uploadTaskWithRequest:fromData:"), request, bodyData)
	return NSURLSessionUploadTaskFromID(rv)
}
// Creates a task that performs an HTTP request for the specified URL request
// object, uploads the provided data, and calls a handler upon completion.
//
// request: A URL request object that provides the URL, cache policy, request type, and
// so on. The body stream and body data in this request object are ignored.
//
// bodyData: The body data for the request.
//
// completionHandler: The completion handler to call when the load request is complete. This
// handler is executed on the delegate queue.
// 
// If you pass `nil`, only the session delegate methods are called when the
// task completes, making this method equivalent to the
// [UploadTaskWithRequestFromData] method.
// 
// This completion handler takes the following parameters:
// 
// `data`: The data returned by the server. `response`: An object that
// provides response metadata, such as HTTP headers and status code. If you
// are making an HTTP or HTTPS request, the returned object is actually an
// [NSHTTPURLResponse] object. `error`: An error object that indicates why the
// request failed, or `nil` if the request was successful.
//
// # Return Value
// 
// The new session upload task.
//
// # Discussion
// 
// An HTTP upload request is any request that contains a request body, such as
// a [POST] or [PUT] request. Upload tasks require you to create a request
// object so that you can provide metadata for the upload, like HTTP request
// headers.
// 
// Unlike [UploadTaskWithRequestFromData], this method returns the response
// body after it has been received in full, and does not require you to write
// a custom delegate to obtain the response body.
// 
// By using a completion handler, the task bypasses calls to delegate methods
// for response and data delivery, and instead provides any resulting data,
// response, or error inside the completion handler. Delegate methods for
// handling authentication challenges, however, are still called.
// 
// Typically you pass a `nil` completion handler only when creating tasks in
// sessions whose delegates include a [URLSessionDataTaskDidReceiveData]
// method. However, if you do not need the response data, use key-value
// observing to watch for changes to the task’s status to determine when it
// completes.
// 
// After you create the task, you must start it by calling its [Resume]
// method.
// 
// If the request completes successfully, the `data` parameter of the
// completion handler block contains the resource data, and the `error`
// parameter is `nil`. If the request fails, the `data` parameter is `nil,`
// and the `error` parameter contains information about the failure. If a
// response from the server is received, regardless of whether the request
// completes successfully or fails, the `response` parameter contains that
// information.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(with:from:completionHandler:)
func (u URLSession) UploadTaskWithRequestFromDataCompletionHandler(request INSURLRequest, bodyData INSData, completionHandler DataURLResponseErrorHandler) INSURLSessionUploadTask {
_block2, _ := NewDataURLResponseErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](u.ID, objc.Sel("uploadTaskWithRequest:fromData:completionHandler:"), request, bodyData, _block2)
	return NSURLSessionUploadTaskFromID(rv)
}
// Creates a task that performs an HTTP request for uploading the specified
// file.
//
// request: A URL request object that provides the URL, cache policy, request type, and
// so on. The body stream and body data in this request object are ignored.
//
// fileURL: The URL of the file to upload.
//
// # Return Value
// 
// The new session upload task.
//
// # Discussion
// 
// An HTTP upload request is any request that contains a request body, such as
// a [POST] or [PUT] request. Upload tasks require you to create a request
// object so that you can provide metadata for the upload, like HTTP request
// headers.
// 
// After you create the task, you must start it by calling its [Resume]
// method. The task calls methods on the session’s delegate to provide you
// with the upload’s progress, response metadata, response data, and so on.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(with:fromFile:)
func (u URLSession) UploadTaskWithRequestFromFile(request INSURLRequest, fileURL INSURL) INSURLSessionUploadTask {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("uploadTaskWithRequest:fromFile:"), request, fileURL)
	return NSURLSessionUploadTaskFromID(rv)
}
// Creates a task that performs an HTTP request for uploading the specified
// file, then calls a handler upon completion.
//
// request: A URL request object that provides the URL, cache policy, request type, and
// so on. The body stream and body data in this request object are ignored.
//
// fileURL: The URL of the file to upload.
//
// completionHandler: The completion handler to call when the load request is complete. This
// handler is executed on the delegate queue.
// 
// If you pass `nil`, only the session delegate methods are called when the
// task completes, making this method equivalent to the
// [UploadTaskWithRequestFromFile] method.
// 
// This completion handler takes the following parameters:
// 
// `data`: The data returned by the server. `response`: An object that
// provides response metadata, such as HTTP headers and status code. If you
// are making an HTTP or HTTPS request, the returned object is actually an
// [NSHTTPURLResponse] object. `error`: An error object that indicates why the
// request failed, or `nil` if the request was successful.
//
// # Return Value
// 
// The new session upload task.
//
// # Discussion
// 
// An HTTP upload request is any request that contains a request body, such as
// a [POST] or [PUT] request. Upload tasks require you to create a request
// object so that you can provide metadata for the upload, like HTTP request
// headers.
// 
// Unlike [UploadTaskWithRequestFromFile], this method returns the response
// body after it has been received in full, and does not require you to write
// a custom delegate to obtain the response body.
// 
// By using a completion handler, the task bypasses calls to delegate methods
// for response and data delivery, and instead provides any resulting data,
// response, or error inside the completion handler. Delegate methods for
// handling authentication challenges, however, are still called.
// 
// Typically you usually pass a `nil` completion handler only when creating
// tasks in sessions whose delegates include a
// [URLSessionDataTaskDidReceiveData] method. However, if you do not need the
// response data, use key-value observing to watch for changes to the task’s
// status to determine when it completes.
// 
// After you create the task, you must start it by calling its [Resume]
// method.
// 
// If the request completes successfully, the `data` parameter of the
// completion handler block contains the resource data, and the `error`
// parameter is `nil`. If the request fails, the `data` parameter is `nil,`
// and the `error` parameter contains information about the failure. If a
// response from the server is received, regardless of whether the request
// completes successfully or fails, the `response` parameter contains that
// information.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(with:fromFile:completionHandler:)
func (u URLSession) UploadTaskWithRequestFromFileCompletionHandler(request INSURLRequest, fileURL INSURL, completionHandler DataURLResponseErrorHandler) INSURLSessionUploadTask {
_block2, _ := NewDataURLResponseErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](u.ID, objc.Sel("uploadTaskWithRequest:fromFile:completionHandler:"), request, fileURL, _block2)
	return NSURLSessionUploadTaskFromID(rv)
}
// Creates a task that performs an HTTP request for uploading data based on
// the specified URL request.
//
// request: A URL request object that provides the URL, cache policy, request type, and
// so on. The body stream and body data in this request object are ignored,
// and the session calls its delegate’s [URLSessionTaskNeedNewBodyStream]
// method to provide the body data.
//
// # Return Value
// 
// The new session upload task.
//
// # Discussion
// 
// An HTTP upload request is any request that contains a request body, such as
// a [POST] or [PUT] request. Upload tasks require you to provide a request
// object so that you can provide metadata for the upload, such as HTTP
// request headers.
// 
// After you create the task, you must start it by calling its [Resume]
// method. The task calls methods on the session’s delegate to provide you
// with the upload’s progress, response metadata, response data, and so on.
// The session’s delegate must have a [URLSessionTaskNeedNewBodyStream]
// method that provides the body data to upload.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(withStreamedRequest:)
func (u URLSession) UploadTaskWithStreamedRequest(request INSURLRequest) INSURLSessionUploadTask {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("uploadTaskWithStreamedRequest:"), request)
	return NSURLSessionUploadTaskFromID(rv)
}
// Creates an upload task from a resume data blob. Requires the server to
// support the latest resumable uploads Internet-Draft from the HTTP Working
// Group, found at
// https://datatracker.ietf.org/doc/draft-ietf-httpbis-resumable-upload/ If
// resuming from an upload file, the file must still exist and be unmodified.
// If the upload cannot be successfully resumed,
// URLSession:task:didCompleteWithError: will be called.
//
// resumeData: Resume data blob from an incomplete upload, such as data returned by the
// cancelByProducingResumeData: method.
//
// # Return Value
// 
// A new session upload task, or nil if the resumeData is invalid.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(withResumeData:)
func (u URLSession) UploadTaskWithResumeData(resumeData INSData) INSURLSessionUploadTask {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("uploadTaskWithResumeData:"), resumeData)
	return NSURLSessionUploadTaskFromID(rv)
}
// Creates a URLSessionUploadTask from a resume data blob. If resuming from an
// upload file, the file must still exist and be unmodified.
//
// resumeData: Resume data blob from an incomplete upload, such as data returned by the
// cancelByProducingResumeData: method.
//
// completionHandler: The completion handler to call when the load request is complete.
//
// # Return Value
// 
// A new session upload task, or nil if the resumeData is invalid.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(withResumeData:completionHandler:)
func (u URLSession) UploadTaskWithResumeDataCompletionHandler(resumeData INSData, completionHandler DataURLResponseErrorHandler) INSURLSessionUploadTask {
_block1, _ := NewDataURLResponseErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](u.ID, objc.Sel("uploadTaskWithResumeData:completionHandler:"), resumeData, _block1)
	return NSURLSessionUploadTaskFromID(rv)
}
// Creates a task that establishes a bidirectional TCP/IP connection to a
// specified hostname and port.
//
// hostname: The hostname of the connection endpoint.
//
// port: The port of the connection endpoint.
//
// # Return Value
// 
// The new session stream task.
//
// # Discussion
// 
// After you create the task, you must start it by calling its [Resume]
// method.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/streamTask(withHostName:port:)
func (u URLSession) StreamTaskWithHostNamePort(hostname string, port int) INSURLSessionStreamTask {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("streamTaskWithHostName:port:"), objc.String(hostname), port)
	return NSURLSessionStreamTaskFromID(rv)
}
// Creates a WebSocket task for the provided URL.
//
// url: The WebSocket URL with which to connect.
//
// # Discussion
// 
// The provided URL must have a `ws` or `wss` scheme.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/webSocketTask(with:)-87ipz
func (u URLSession) WebSocketTaskWithURL(url INSURL) INSURLSessionWebSocketTask {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("webSocketTaskWithURL:"), url)
	return NSURLSessionWebSocketTaskFromID(rv)
}
// Creates a WebSocket task for the provided URL request.
//
// request: A URL request that indicates a WebSockets endpoint with which to connect.
//
// # Discussion
// 
// You can modify the request’s properties prior to calling [Resume] on the
// task. The task uses these properties during the HTTP handshake phase.
// 
// To add custom protocols, add a header with the key
// `Sec-WebSocket-Protocol`, and a comma-separated list of protocols you want
// to negotiate with the server. The custom HTTP headers provided by the
// client remain unchanged for the handshake with the server.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/webSocketTask(with:)-mtks
func (u URLSession) WebSocketTaskWithRequest(request INSURLRequest) INSURLSessionWebSocketTask {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("webSocketTaskWithRequest:"), request)
	return NSURLSessionWebSocketTaskFromID(rv)
}
// Creates a WebSocket task given a URL and an array of protocols.
//
// url: The WebSocket URL with which to connect.
//
// protocols: An array of protocols to negotiate with the server.
//
// # Discussion
// 
// During the WebSocket handshake, the task uses the provided protocols to
// negotiate a preferred protocol with the server.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/webSocketTask(with:protocols:)
func (u URLSession) WebSocketTaskWithURLProtocols(url INSURL, protocols []string) INSURLSessionWebSocketTask {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("webSocketTaskWithURL:protocols:"), url, objectivec.StringSliceToNSArray(protocols))
	return NSURLSessionWebSocketTaskFromID(rv)
}
// Invalidates the session, allowing any outstanding tasks to finish.
//
// # Discussion
// 
// This method returns immediately without waiting for tasks to finish. Once a
// session is invalidated, new tasks cannot be created in the session, but
// existing tasks continue until completion. After the last task finishes and
// the session makes the last delegate call related to those tasks, the
// session calls the [URLSessionDidBecomeInvalidWithError] method on its
// delegate, then breaks references to the delegate and callback objects.
// After invalidation, session objects cannot be reused.
// 
// To cancel all outstanding tasks, call [InvalidateAndCancel] instead.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/finishTasksAndInvalidate()
func (u URLSession) FinishTasksAndInvalidate() {
	objc.Send[objc.ID](u.ID, objc.Sel("finishTasksAndInvalidate"))
}
// Flushes cookies and credentials to disk, clears transient caches, and
// ensures that future requests occur on a new TCP connection.
//
// completionHandler: The completion handler to call when the flush operation is complete. This
// handler is executed on the delegate queue.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/flush(completionHandler:)
func (u URLSession) FlushWithCompletionHandler(completionHandler VoidHandler) {
_block0, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](u.ID, objc.Sel("flushWithCompletionHandler:"), _block0)
}
// Asynchronously calls a completion callback with all data, upload, and
// download tasks in a session.
//
// completionHandler: The completion handler to call with the list of tasks. This handler is
// executed on the delegate queue.
//
// # Discussion
// 
// The arrays passed to the completion handler contain any tasks that you have
// created within the session, not including any tasks that have been
// invalidated after completing, failing, or being cancelled.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/getTasksWithCompletionHandler(_:)
func (u URLSession) GetTasksWithCompletionHandler(completionHandler VoidHandler) {
_block0, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](u.ID, objc.Sel("getTasksWithCompletionHandler:"), _block0)
}
// Cancels all outstanding tasks and then invalidates the session.
//
// # Discussion
// 
// Once invalidated, references to the delegate and callback objects are
// broken. After invalidation, session objects cannot be reused.
// 
// To allow outstanding tasks to run until completion, call
// [FinishTasksAndInvalidate] instead.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/invalidateAndCancel()
func (u URLSession) InvalidateAndCancel() {
	objc.Send[objc.ID](u.ID, objc.Sel("invalidateAndCancel"))
}
// Empties all cookies, caches and credential stores, removes disk files,
// flushes in-progress downloads to disk, and ensures that future requests
// occur on a new socket.
//
// completionHandler: The completion handler to call when the reset operation is complete. This
// handler is executed on the delegate queue.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/reset(completionHandler:)
func (u URLSession) ResetWithCompletionHandler(completionHandler VoidHandler) {
_block0, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](u.ID, objc.Sel("resetWithCompletionHandler:"), _block0)
}

// A copy of the configuration object for this session.
//
// # Discussion
// 
// Beginning in iOS 9 and OS X 10.11, [NSURLSession] objects store a copy of
// the [NSURLSessionConfiguration] object passed to their initializers, such
// that a session’s configuration is immutable after initialization. Any
// further changes to mutable properties on the configuration object passed to
// a session’s initializer or the value returned from a session’s
// configuration property do not affect the behavior of that session. However,
// you can create a new session with the modified configuration object.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/configuration
func (u URLSession) Configuration() INSURLSessionConfiguration {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("configuration"))
	return NSURLSessionConfigurationFromID(objc.ID(rv))
}
// The delegate assigned when this object was created.
//
// # Discussion
// 
// This delegate object is responsible for handling authentication challenges,
// for making caching decisions, and for handling other session-related
// events. The session object keeps a strong reference to this delegate until
// your app exits or explicitly invalidates the session. If you do not
// invalidate the session, your app leaks memory until it exits.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/delegate
func (u URLSession) Delegate() NSURLSessionDelegate {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("delegate"))
	return NSURLSessionDelegateObjectFromID(rv)
}
// The operation queue provided when this object was created.
//
// # Discussion
// 
// All delegate method calls and completion handlers related to the session
// are performed on this queue. The session object keeps a strong reference to
// this queue until your app exits or the session object is deallocated. If
// you do not invalidate the session, your app leaks memory until it exits.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/delegateQueue
func (u URLSession) DelegateQueue() INSOperationQueue {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("delegateQueue"))
	return NSOperationQueueFromID(objc.ID(rv))
}
// An app-defined descriptive label for the session.
//
// # Discussion
// 
// This property contains a human-readable string that you can use for
// debugging purposes. This value may be `nil` and defaults to `nil`. The
// value is ignored by the session.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/sessionDescription
func (u URLSession) SessionDescription() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("sessionDescription"))
	return NSStringFromID(rv).String()
}
func (u URLSession) SetSessionDescription(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setSessionDescription:"), objc.String(value))
}

// The shared singleton session object.
//
// # Discussion
// 
// For basic requests, the [NSURLSession] class provides a shared singleton
// session object that gives you a reasonable default behavior for creating
// tasks. Use the shared session to fetch the contents of a URL to memory with
// just a few lines of code.
// 
// Unlike the other session types, you don’t create the shared session; you
// merely access it by using this property directly. As a result, you don’t
// provide a delegate or a configuration object.
// 
// # Limitations of the shared session
// 
// Because the shared session has neither a delegate nor a customizable
// configuration object, the shared session has important limitations:
// 
// - You can’t obtain data incrementally as it arrives from the server. -
// You can’t significantly customize the default connection behavior. - Your
// ability to perform authentication is limited. - You can’t perform
// background downloads or uploads when your app isn’t running.
// 
// The shared session uses the shared [NSURLCache], [NSHTTPCookieStorage], and
// [NSURLCredentialStorage] objects, uses a shared custom networking protocol
// list (configured with [RegisterClass] and [UnregisterClass]), and is based
// on a default configuration.
// 
// In general, when working with a shared session, you should avoid
// customizing the cache, cookie storage, or credential storage (unless you
// are already doing so with [NSURLConnection]). There’s a good chance that
// you’ll outgrow the capabilities of the default session, at which point
// you’ll have to rewrite all of those customizations to work with your
// custom URL sessions.
// 
// In other words, if you’re doing with caches, cookies, authentication, or
// custom networking protocols, you should probably be using a default session
// instead of the shared session.
//
// See: https://developer.apple.com/documentation/Foundation/URLSession/shared
func (_URLSessionClass URLSessionClass) SharedSession() URLSession {
	rv := objc.Send[objc.ID](objc.ID(_URLSessionClass.class), objc.Sel("sharedSession"))
	return NSURLSessionFromID(objc.ID(rv))
}

// Flush is a synchronous wrapper around [URLSession.FlushWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u URLSession) Flush(ctx context.Context) error {
	done := make(chan struct{}, 1)
	u.FlushWithCompletionHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Reset is a synchronous wrapper around [URLSession.ResetWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u URLSession) Reset(ctx context.Context) error {
	done := make(chan struct{}, 1)
	u.ResetWithCompletionHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

