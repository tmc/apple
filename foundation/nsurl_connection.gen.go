// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSURLConnection] class.
var (
	_NSURLConnectionClass     NSURLConnectionClass
	_NSURLConnectionClassOnce sync.Once
)

func getNSURLConnectionClass() NSURLConnectionClass {
	_NSURLConnectionClassOnce.Do(func() {
		_NSURLConnectionClass = NSURLConnectionClass{class: objc.GetClass("NSURLConnection")}
	})
	return _NSURLConnectionClass
}

// GetNSURLConnectionClass returns the class object for NSURLConnection.
func GetNSURLConnectionClass() NSURLConnectionClass {
	return getNSURLConnectionClass()
}

type NSURLConnectionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSURLConnectionClass) Alloc() NSURLConnection {
	rv := objc.Send[NSURLConnection](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that enables you to start and stop URL requests.
//
// # Overview
// 
// An [NSURLConnection] object lets you load the contents of a URL by
// providing a URL request object. The interface for [NSURLConnection] is
// sparse, providing only the controls to start and cancel asynchronous loads
// of a URL request. You perform most of your configuration on the URL request
// object itself.
// 
// The [NSURLConnection] class provides convenience class methods to load URL
// requests both asynchronously using a callback block and synchronously.
// 
// For greater control, you can create a URL connection object with a delegate
// object that conforms to the [NSURLConnectionDelegate] and
// [NSURLConnectionDataDelegate] protocols. The connection calls methods on
// that delegate to provide you with progress and status as the URL request is
// loaded asynchronously. The connection also calls delegate methods to let
// you override the connection’s default behavior (for example, specifying
// how a particular redirect should be handled). These delegate methods are
// called on the thread that initiated the asynchronous load operation.
// 
// For more information about errors, see the `NSURLError.H()` header,
// [Foundation Constants], and URL Loading System Error Codes in [Error
// Handling Programming Guide].
// 
// # NSURLConnection Protocols
// 
// The [NSURLConnection] class works in tandem with three formal protocols:
// [NSURLConnectionDelegate], [NSURLConnectionDataDelegate], and
// [NSURLConnectionDownloadDelegate]. To use these protocols, you write a
// class that conforms to them and implement any methods that are appropriate,
// then provide an instance of that class as the delegate when you create a
// connection object.
// 
// The [NSURLConnectionDelegate] protocol is primarily used for credential
// handling, but also handles connection completion. Because it handles
// connection failure during data transfers, all connection delegates must
// typically implement this protocol.
// 
// In addition, unless you’re using Newsstand Kit, your delegate must also
// conform to the [NSURLConnectionDataDelegate] protocol, because this
// protocol provides methods that the [NSURLConnection] class calls with
// progress information during an upload, with fragments of the response data
// during a download, and to provide a new upload body stream if the
// server’s response necessitates a second connection attempt—for example,
// if [NSURLConnection] must retry the request with different credentials.
// 
// Finally, if you’re using Newsstand Kit, your delegate can conform to the
// [NSURLConnectionDownloadDelegate] protocol. This protocol provides support
// for continuing interrupted file downloads and receiving a notification
// whenever a download finishes. This protocol is solely for use with
// [NSURLConnection] objects created using Newsstand Kit’s `download()`
// method.
//
// [Error Handling Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ErrorHandlingCocoa/ErrorHandling/ErrorHandling.html#//apple_ref/doc/uid/TP40001806
// [Foundation Constants]: https://developer.apple.com/documentation/Foundation/foundation-constants
//
// # Connection URL Information
//
//   - [NSURLConnection.OriginalRequest]: A deep copy of the original connection request.
//   - [NSURLConnection.CurrentRequest]: The current connection request.
//
// # Loading Data Asynchronously
//
//   - [NSURLConnection.Start]: Causes the connection to begin loading data, if it has not already.
//
// # Stopping a Connection
//
//   - [NSURLConnection.Cancel]: Cancels an asynchronous load of a request.
//
// # Scheduling Delegate Method Calls
//
//   - [NSURLConnection.ScheduleInRunLoopForMode]: Determines the run loop and mode that the connection uses to call methods on its delegate.
//   - [NSURLConnection.SetDelegateQueue]: Determines the operation queue that is used to call methods on the connection’s delegate.
//   - [NSURLConnection.UnscheduleFromRunLoopForMode]: Causes the connection to stop calling delegate methods in the specified run loop and mode.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnection
type NSURLConnection struct {
	objectivec.Object
}

// NSURLConnectionFromID constructs a [NSURLConnection] from an objc.ID.
//
// An object that enables you to start and stop URL requests.
func NSURLConnectionFromID(id objc.ID) NSURLConnection {
	return NSURLConnection{objectivec.Object{ID: id}}
}
// NOTE: NSURLConnection adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSURLConnection] class.
//
// # Connection URL Information
//
//   - [INSURLConnection.OriginalRequest]: A deep copy of the original connection request.
//   - [INSURLConnection.CurrentRequest]: The current connection request.
//
// # Loading Data Asynchronously
//
//   - [INSURLConnection.Start]: Causes the connection to begin loading data, if it has not already.
//
// # Stopping a Connection
//
//   - [INSURLConnection.Cancel]: Cancels an asynchronous load of a request.
//
// # Scheduling Delegate Method Calls
//
//   - [INSURLConnection.ScheduleInRunLoopForMode]: Determines the run loop and mode that the connection uses to call methods on its delegate.
//   - [INSURLConnection.SetDelegateQueue]: Determines the operation queue that is used to call methods on the connection’s delegate.
//   - [INSURLConnection.UnscheduleFromRunLoopForMode]: Causes the connection to stop calling delegate methods in the specified run loop and mode.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnection
type INSURLConnection interface {
	objectivec.IObject

	// Topic: Connection URL Information

	// A deep copy of the original connection request.
	OriginalRequest() INSURLRequest
	// The current connection request.
	CurrentRequest() INSURLRequest

	// Topic: Loading Data Asynchronously

	// Causes the connection to begin loading data, if it has not already.
	Start()

	// Topic: Stopping a Connection

	// Cancels an asynchronous load of a request.
	Cancel()

	// Topic: Scheduling Delegate Method Calls

	// Determines the run loop and mode that the connection uses to call methods on its delegate.
	ScheduleInRunLoopForMode(aRunLoop INSRunLoop, mode NSRunLoopMode)
	// Determines the operation queue that is used to call methods on the connection’s delegate.
	SetDelegateQueue(queue INSOperationQueue)
	// Causes the connection to stop calling delegate methods in the specified run loop and mode.
	UnscheduleFromRunLoopForMode(aRunLoop INSRunLoop, mode NSRunLoopMode)
}

// Init initializes the instance.
func (u NSURLConnection) Init() NSURLConnection {
	rv := objc.Send[NSURLConnection](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u NSURLConnection) Autorelease() NSURLConnection {
	rv := objc.Send[NSURLConnection](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSURLConnection creates a new NSURLConnection instance.
func NewNSURLConnection() NSURLConnection {
	class := getNSURLConnectionClass()
	rv := objc.Send[NSURLConnection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Causes the connection to begin loading data, if it has not already.
//
// # Discussion
// 
// Calling this method is necessary only if you create a connection with the
// [init(request:delegate:startImmediately:)] method and provide [false] for
// the `startImmediately` parameter. If you don’t schedule the connection in
// a run loop or an operation queue before calling this method, the connection
// is scheduled in the current run loop in the default mode.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [init(request:delegate:startImmediately:)]: https://developer.apple.com/documentation/Foundation/NSURLConnection/init(request:delegate:startImmediately:)
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnection/start()
func (u NSURLConnection) Start() {
	objc.Send[objc.ID](u.ID, objc.Sel("start"))
}

// Cancels an asynchronous load of a request.
//
// # Discussion
// 
// After this method is called, the connection makes no further delegate
// method calls. If you want to reattempt the connection, you should create a
// new connection object.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnection/cancel()
func (u NSURLConnection) Cancel() {
	objc.Send[objc.ID](u.ID, objc.Sel("cancel"))
}

// Determines the run loop and mode that the connection uses to call methods
// on its delegate.
//
// aRunLoop: The [NSRunLoop] instance to use when calling delegate methods.
//
// mode: The mode in which to call delegate methods.
//
// # Discussion
// 
// By default, a connection is scheduled on the current thread in the default
// mode when it is created. If you create a connection with the
// [init(request:delegate:startImmediately:)] method and provide [false] for
// the `startImmediately` parameter, you can schedule the connection on a
// different run loop or mode before starting it with the [Start] method. You
// can schedule a connection on multiple run loops and modes, or on the same
// run loop in multiple modes.
// 
// You cannot reschedule a connection after it has started.
// 
// It is an error to schedule delegate method calls with both this method and
// the [SetDelegateQueue] method.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [init(request:delegate:startImmediately:)]: https://developer.apple.com/documentation/Foundation/NSURLConnection/init(request:delegate:startImmediately:)
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnection/schedule(in:forMode:)
func (u NSURLConnection) ScheduleInRunLoopForMode(aRunLoop INSRunLoop, mode NSRunLoopMode) {
	objc.Send[objc.ID](u.ID, objc.Sel("scheduleInRunLoop:forMode:"), aRunLoop, objc.String(string(mode)))
}

// Determines the operation queue that is used to call methods on the
// connection’s delegate.
//
// queue: The operation queue to use when calling delegate methods.
//
// # Discussion
// 
// By default, a connection is scheduled on the current thread in the default
// mode when it is created. If you create a connection with the
// [init(request:delegate:startImmediately:)] method and provide [false] for
// the `startImmediately` parameter, you can instead schedule the connection
// on an operation queue before starting it with the [Start] method.
// 
// You cannot reschedule a connection after it has started.
// 
// It is an error to schedule delegate method calls with both this method and
// the [ScheduleInRunLoopForMode] method.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [init(request:delegate:startImmediately:)]: https://developer.apple.com/documentation/Foundation/NSURLConnection/init(request:delegate:startImmediately:)
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnection/setDelegateQueue(_:)
func (u NSURLConnection) SetDelegateQueue(queue INSOperationQueue) {
	objc.Send[objc.ID](u.ID, objc.Sel("setDelegateQueue:"), queue)
}

// Causes the connection to stop calling delegate methods in the specified run
// loop and mode.
//
// aRunLoop: The run loop instance to unschedule.
//
// mode: The mode to unschedule.
//
// # Discussion
// 
// By default, a connection is scheduled on the current thread in the default
// mode when it is created. If you create a connection with the
// [init(request:delegate:startImmediately:)] method and provide [false] for
// the `startImmediately` parameter, you can instead schedule connection on a
// different run loop or mode before starting it with the [Start] method. You
// can schedule a connection on multiple run loops and modes, or on the same
// run loop in multiple modes. Use this method to unschedule the connection
// from an undesired run loop and mode before starting the connection.
// 
// You cannot reschedule a connection after it has started. It is not
// necessary to unschedule a connection after it has finished.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [init(request:delegate:startImmediately:)]: https://developer.apple.com/documentation/Foundation/NSURLConnection/init(request:delegate:startImmediately:)
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnection/unschedule(from:forMode:)
func (u NSURLConnection) UnscheduleFromRunLoopForMode(aRunLoop INSRunLoop, mode NSRunLoopMode) {
	objc.Send[objc.ID](u.ID, objc.Sel("unscheduleFromRunLoop:forMode:"), aRunLoop, objc.String(string(mode)))
}

// Returns whether a request can be handled based on a preflight evaluation.
//
// request: The request to evaluate. The connection deep-copies the request on
// creation.
//
// # Return Value
// 
// [true] if a preflight operation determines that a connection with `request`
// can be created and the associated I/O can be started, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The result of this method is valid as long as no [NSURLProtocol] classes
// are registered or unregistered, and `request` remains unchanged.
// Applications should be prepared to handle failures even if they have
// performed request preflighting by calling this method.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnection/canHandle(_:)
func (_NSURLConnectionClass NSURLConnectionClass) CanHandleRequest(request INSURLRequest) bool {
	rv := objc.Send[bool](objc.ID(_NSURLConnectionClass.class), objc.Sel("canHandleRequest:"), request)
	return rv
}

// A deep copy of the original connection request.
//
// # Discussion
// 
// As the connection performs the load, the request may change as a result of
// protocol canonicalization or due to following redirects. [CurrentRequest]
// can be used to retrieve this value.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnection/originalRequest
func (u NSURLConnection) OriginalRequest() INSURLRequest {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("originalRequest"))
	return NSURLRequestFromID(objc.ID(rv))
}

// The current connection request.
//
// # Discussion
// 
// As the connection performs the load, the request may change as a result of
// protocol canonicalization or due to following redirects. This property
// provides the current value of the request.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnection/currentRequest
func (u NSURLConnection) CurrentRequest() INSURLRequest {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("currentRequest"))
	return NSURLRequestFromID(objc.ID(rv))
}

