// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [URLSessionTask] class.
var (
	_URLSessionTaskClass     URLSessionTaskClass
	_URLSessionTaskClassOnce sync.Once
)

func getURLSessionTaskClass() URLSessionTaskClass {
	_URLSessionTaskClassOnce.Do(func() {
		_URLSessionTaskClass = URLSessionTaskClass{class: objc.GetClass("NSURLSessionTask")}
	})
	return _URLSessionTaskClass
}

// GetURLSessionTaskClass returns the class object for NSURLSessionTask.
func GetURLSessionTaskClass() URLSessionTaskClass {
	return getURLSessionTaskClass()
}

type URLSessionTaskClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc URLSessionTaskClass) Alloc() URLSessionTask {
	rv := objc.Send[URLSessionTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A task, like downloading a specific resource, performed in a URL session.
//
// # Overview
// 
// The [NSURLSessionTask] class is the base class for tasks in a URL session.
// Tasks are always part of a session; you create a task by calling one of the
// task creation methods on a [NSURLSession] instance. The method you call
// determines the type of task.
// 
// - Use [NSURLSession]‘s [DataTaskWithURL] and related methods to create
// [NSURLSessionDataTask] instances. Data tasks request a resource, returning
// the server’s response as one or more [NSData] objects in memory. They are
// supported in default, ephemeral, and shared sessions, but are not supported
// in background sessions. - Use [NSURLSession]‘s
// [UploadTaskWithRequestFromData] and related methods to create
// [NSURLSessionUploadTask] instances. Upload tasks are like data tasks,
// except that they make it easier to provide a request body so you can upload
// data before retrieving the server’s response. Additionally, upload tasks
// are supported in background sessions. - Use [NSURLSession]’s
// [DownloadTaskWithURL] and related methods to create
// [NSURLSessionDownloadTask] instances. Download tasks download a resource
// directly to a file on disk. Download tasks are supported in any type of
// session. - Use [NSURLSession]’s [StreamTaskWithHostNamePort] or
// [StreamTaskWithNetService] to create [NSURLSessionStreamTask] instances.
// Stream tasks establish a TCP/IP connection from a host name and port or a
// net service object.
// 
// After you create a task, you start it by calling its [Resume] method. The
// session then maintains a strong reference to the task until the request
// finishes or fails; you don’t need to maintain a reference to the task
// unless it’s useful for your app’s internal bookkeeping.
//
// # Controlling the task state
//
//   - [URLSessionTask.Cancel]: Cancels the task.
//   - [URLSessionTask.Resume]: Resumes the task, if it is suspended.
//   - [URLSessionTask.Suspend]: Temporarily suspends a task.
//   - [URLSessionTask.State]: The current state of the task—active, suspended, in the process of being canceled, or completed.
//   - [URLSessionTask.Priority]: The relative priority at which you’d like a host to handle the task, specified as a floating point value between `0.0` (lowest priority) and `1.0` (highest priority).
//   - [URLSessionTask.SetPriority]
//
// # Obtaining task progress
//
//   - [URLSessionTask.CountOfBytesExpectedToReceive]: The number of bytes that the task expects to receive in the response body.
//   - [URLSessionTask.CountOfBytesReceived]: The number of bytes that the task has received from the server in the response body.
//   - [URLSessionTask.CountOfBytesExpectedToSend]: The number of bytes that the task expects to send in the request body.
//   - [URLSessionTask.CountOfBytesSent]: The number of bytes that the task has sent to the server in the request body.
//   - [URLSessionTask.NSURLSessionTransferSizeUnknown]: The total size of the transfer cannot be determined.
//
// # Obtaining general task information
//
//   - [URLSessionTask.CurrentRequest]: The URL request object currently being handled by the task.
//   - [URLSessionTask.OriginalRequest]: The original request object passed when the task was created.
//   - [URLSessionTask.Response]: The server’s response to the currently active request.
//   - [URLSessionTask.TaskDescription]: An app-provided string value for the current task.
//   - [URLSessionTask.SetTaskDescription]
//   - [URLSessionTask.TaskIdentifier]: An identifier uniquely identifying the task within a given session.
//   - [URLSessionTask.Error]: An error object that indicates why the task failed.
//
// # Determining task behavior
//
//   - [URLSessionTask.PrefersIncrementalDelivery]: A Boolean value that determines whether to deliver a partial response body in increments.
//   - [URLSessionTask.SetPrefersIncrementalDelivery]
//
// # Using a task-specific delegate
//
//   - [URLSessionTask.Delegate]: A delegate specific to the task.
//   - [URLSessionTask.SetDelegate]
//
// # Scheduling tasks
//
//   - [URLSessionTask.CountOfBytesClientExpectsToReceive]: A best-guess upper bound on the number of bytes the client expects to receive.
//   - [URLSessionTask.SetCountOfBytesClientExpectsToReceive]
//   - [URLSessionTask.CountOfBytesClientExpectsToSend]: A best-guess upper bound on the number of bytes the client expects to send.
//   - [URLSessionTask.SetCountOfBytesClientExpectsToSend]
//   - [URLSessionTask.EarliestBeginDate]: The earliest date at which the network load should begin.
//   - [URLSessionTask.SetEarliestBeginDate]
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask
type URLSessionTask struct {
	objectivec.Object
}

// URLSessionTaskFromID constructs a [URLSessionTask] from an objc.ID.
//
// A task, like downloading a specific resource, performed in a URL session.
func URLSessionTaskFromID(id objc.ID) URLSessionTask {
	return NSURLSessionTask{objectivec.Object{ID: id}}
}

// NSURLSessionTaskFromID is an alias for [URLSessionTaskFromID] for cross-framework compatibility.
func NSURLSessionTaskFromID(id objc.ID) URLSessionTask { return URLSessionTaskFromID(id) }
// NOTE: URLSessionTask adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [URLSessionTask] class.
//
// # Controlling the task state
//
//   - [IURLSessionTask.Cancel]: Cancels the task.
//   - [IURLSessionTask.Resume]: Resumes the task, if it is suspended.
//   - [IURLSessionTask.Suspend]: Temporarily suspends a task.
//   - [IURLSessionTask.State]: The current state of the task—active, suspended, in the process of being canceled, or completed.
//   - [IURLSessionTask.Priority]: The relative priority at which you’d like a host to handle the task, specified as a floating point value between `0.0` (lowest priority) and `1.0` (highest priority).
//   - [IURLSessionTask.SetPriority]
//
// # Obtaining task progress
//
//   - [IURLSessionTask.CountOfBytesExpectedToReceive]: The number of bytes that the task expects to receive in the response body.
//   - [IURLSessionTask.CountOfBytesReceived]: The number of bytes that the task has received from the server in the response body.
//   - [IURLSessionTask.CountOfBytesExpectedToSend]: The number of bytes that the task expects to send in the request body.
//   - [IURLSessionTask.CountOfBytesSent]: The number of bytes that the task has sent to the server in the request body.
//   - [IURLSessionTask.NSURLSessionTransferSizeUnknown]: The total size of the transfer cannot be determined.
//
// # Obtaining general task information
//
//   - [IURLSessionTask.CurrentRequest]: The URL request object currently being handled by the task.
//   - [IURLSessionTask.OriginalRequest]: The original request object passed when the task was created.
//   - [IURLSessionTask.Response]: The server’s response to the currently active request.
//   - [IURLSessionTask.TaskDescription]: An app-provided string value for the current task.
//   - [IURLSessionTask.SetTaskDescription]
//   - [IURLSessionTask.TaskIdentifier]: An identifier uniquely identifying the task within a given session.
//   - [IURLSessionTask.Error]: An error object that indicates why the task failed.
//
// # Determining task behavior
//
//   - [IURLSessionTask.PrefersIncrementalDelivery]: A Boolean value that determines whether to deliver a partial response body in increments.
//   - [IURLSessionTask.SetPrefersIncrementalDelivery]
//
// # Using a task-specific delegate
//
//   - [IURLSessionTask.Delegate]: A delegate specific to the task.
//   - [IURLSessionTask.SetDelegate]
//
// # Scheduling tasks
//
//   - [IURLSessionTask.CountOfBytesClientExpectsToReceive]: A best-guess upper bound on the number of bytes the client expects to receive.
//   - [IURLSessionTask.SetCountOfBytesClientExpectsToReceive]
//   - [IURLSessionTask.CountOfBytesClientExpectsToSend]: A best-guess upper bound on the number of bytes the client expects to send.
//   - [IURLSessionTask.SetCountOfBytesClientExpectsToSend]
//   - [IURLSessionTask.EarliestBeginDate]: The earliest date at which the network load should begin.
//   - [IURLSessionTask.SetEarliestBeginDate]
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask
type IURLSessionTask interface {
	objectivec.IObject
	NSCopying

	// Topic: Controlling the task state

	// Cancels the task.
	Cancel()
	// Resumes the task, if it is suspended.
	Resume()
	// Temporarily suspends a task.
	Suspend()
	// The current state of the task—active, suspended, in the process of being canceled, or completed.
	State() NSURLSessionTaskState
	// The relative priority at which you’d like a host to handle the task, specified as a floating point value between `0.0` (lowest priority) and `1.0` (highest priority).
	Priority() float32
	SetPriority(value float32)

	// Topic: Obtaining task progress

	// The number of bytes that the task expects to receive in the response body.
	CountOfBytesExpectedToReceive() int64
	// The number of bytes that the task has received from the server in the response body.
	CountOfBytesReceived() int64
	// The number of bytes that the task expects to send in the request body.
	CountOfBytesExpectedToSend() int64
	// The number of bytes that the task has sent to the server in the request body.
	CountOfBytesSent() int64
	// The total size of the transfer cannot be determined.
	NSURLSessionTransferSizeUnknown() objectivec.IObject

	// Topic: Obtaining general task information

	// The URL request object currently being handled by the task.
	CurrentRequest() INSURLRequest
	// The original request object passed when the task was created.
	OriginalRequest() INSURLRequest
	// The server’s response to the currently active request.
	Response() INSURLResponse
	// An app-provided string value for the current task.
	TaskDescription() string
	SetTaskDescription(value string)
	// An identifier uniquely identifying the task within a given session.
	TaskIdentifier() uint
	// An error object that indicates why the task failed.
	Error() INSError

	// Topic: Determining task behavior

	// A Boolean value that determines whether to deliver a partial response body in increments.
	PrefersIncrementalDelivery() bool
	SetPrefersIncrementalDelivery(value bool)

	// Topic: Using a task-specific delegate

	// A delegate specific to the task.
	Delegate() NSURLSessionTaskDelegate
	SetDelegate(value NSURLSessionTaskDelegate)

	// Topic: Scheduling tasks

	// A best-guess upper bound on the number of bytes the client expects to receive.
	CountOfBytesClientExpectsToReceive() int64
	SetCountOfBytesClientExpectsToReceive(value int64)
	// A best-guess upper bound on the number of bytes the client expects to send.
	CountOfBytesClientExpectsToSend() int64
	SetCountOfBytesClientExpectsToSend(value int64)
	// The earliest date at which the network load should begin.
	EarliestBeginDate() INSDate
	SetEarliestBeginDate(value INSDate)
}

// Init initializes the instance.
func (u URLSessionTask) Init() URLSessionTask {
	rv := objc.Send[URLSessionTask](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u URLSessionTask) Autorelease() URLSessionTask {
	rv := objc.Send[URLSessionTask](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionTask creates a new URLSessionTask instance.
func NewURLSessionTask() URLSessionTask {
	class := getURLSessionTaskClass()
	rv := objc.Send[URLSessionTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Cancels the task.
//
// # Discussion
// 
// This method returns immediately, marking the task as being canceled. Once a
// task is marked as being canceled, [URLSessionTaskDidCompleteWithError] will
// be sent to the task delegate, passing an error in the domain
// [NSURLErrorDomain] with the code [NSURLErrorCancelled]. A task may, under
// some circumstances, send messages to its delegate before the cancelation is
// acknowledged.
// 
// This method may be called on a task that is suspended.
//
// [NSURLErrorCancelled]: https://developer.apple.com/documentation/Foundation/NSURLErrorCancelled-swift.var
// [NSURLErrorDomain]: https://developer.apple.com/documentation/Foundation/NSURLErrorDomain
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/cancel()
func (u URLSessionTask) Cancel() {
	objc.Send[objc.ID](u.ID, objc.Sel("cancel"))
}
// Resumes the task, if it is suspended.
//
// # Discussion
// 
// Newly-initialized tasks begin in a suspended state, so you need to call
// this method to start the task.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/resume()
func (u URLSessionTask) Resume() {
	objc.Send[objc.ID](u.ID, objc.Sel("resume"))
}
// Temporarily suspends a task.
//
// # Discussion
// 
// A task, while suspended, produces no network traffic and isn’t subject to
// timeouts. Call [Resume] to resume data transfer.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/suspend()
func (u URLSessionTask) Suspend() {
	objc.Send[objc.ID](u.ID, objc.Sel("suspend"))
}

// The current state of the task—active, suspended, in the process of being
// canceled, or completed.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/state-swift.property
func (u URLSessionTask) State() NSURLSessionTaskState {
	rv := objc.Send[NSURLSessionTaskState](u.ID, objc.Sel("state"))
	return NSURLSessionTaskState(rv)
}
// The relative priority at which you’d like a host to handle the task,
// specified as a floating point value between `0.0` (lowest priority) and
// `1.0` (highest priority).
//
// # Discussion
// 
// To provide hints to a host on how to prioritize URL session tasks from your
// app, specify a priority for each task. Specifying a priority provides only
// a hint and does not guarantee performance. If you don’t specify a
// priority, a URL session task has a priority of [defaultPriority], with a
// value of `0.5`.
// 
// There are three named priorities you can employ, described in [URL session
// task priority].
// 
// You can specify or change a task’s priority at any time, but not all
// networking protocols respond to changes after a task has started. There is
// no API to let you determine the effective priority for a task from a
// host’s perspective.
//
// [URL session task priority]: https://developer.apple.com/documentation/Foundation/url-session-task-priority
// [defaultPriority]: https://developer.apple.com/documentation/Foundation/URLSessionTask/defaultPriority
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/priority
func (u URLSessionTask) Priority() float32 {
	rv := objc.Send[float32](u.ID, objc.Sel("priority"))
	return rv
}
func (u URLSessionTask) SetPriority(value float32) {
	objc.Send[struct{}](u.ID, objc.Sel("setPriority:"), value)
}
// A representation of the overall task progress.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/progress
func (u URLSessionTask) Progress() INSProgress {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("progress"))
	return NSProgressFromID(objc.ID(rv))
}
// The number of bytes that the task expects to receive in the response body.
//
// # Discussion
// 
// This value is determined based on the `Content-Length` header received from
// the server. If that header is absent, the value is
// [NSURLSessionTransferSizeUnknown].
//
// [NSURLSessionTransferSizeUnknown]: https://developer.apple.com/documentation/Foundation/NSURLSessionTransferSizeUnknown
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesExpectedToReceive
func (u URLSessionTask) CountOfBytesExpectedToReceive() int64 {
	rv := objc.Send[int64](u.ID, objc.Sel("countOfBytesExpectedToReceive"))
	return rv
}
// The number of bytes that the task has received from the server in the
// response body.
//
// # Discussion
// 
// To be notified when this value changes, implement the
// [URLSessionDataTaskDidReceiveData] delegate method (for data and upload
// tasks) or the
// [URLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite]
// method (for download tasks).
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesReceived
func (u URLSessionTask) CountOfBytesReceived() int64 {
	rv := objc.Send[int64](u.ID, objc.Sel("countOfBytesReceived"))
	return rv
}
// The number of bytes that the task expects to send in the request body.
//
// # Discussion
// 
// The URL loading system can determine the length of the upload data in three
// ways:
// 
// - From the length of the data object provided as the upload body. - From
// the length of the file on disk provided as the upload body of an upload
// task ( a download task). - From the `Content-Length` in the request object,
// if you explicitly set it.
// 
// Otherwise, the value is [NSURLSessionTransferSizeUnknown] (`-1`) if you
// provided a stream or body data object, or zero (`0`) if you did not.
//
// [NSURLSessionTransferSizeUnknown]: https://developer.apple.com/documentation/Foundation/NSURLSessionTransferSizeUnknown
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesExpectedToSend
func (u URLSessionTask) CountOfBytesExpectedToSend() int64 {
	rv := objc.Send[int64](u.ID, objc.Sel("countOfBytesExpectedToSend"))
	return rv
}
// The number of bytes that the task has sent to the server in the request
// body.
//
// # Discussion
// 
// This byte count includes the length of the request body itself, not the
// request headers.
// 
// To be notified when this value changes, implement the
// [URLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend]
// delegate method.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesSent
func (u URLSessionTask) CountOfBytesSent() int64 {
	rv := objc.Send[int64](u.ID, objc.Sel("countOfBytesSent"))
	return rv
}
// The URL request object currently being handled by the task.
//
// # Discussion
// 
// This value is typically the same as the initial request ([OriginalRequest])
// except when the server has responded to the initial request with a redirect
// to a different URL.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/currentRequest
func (u URLSessionTask) CurrentRequest() INSURLRequest {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("currentRequest"))
	return NSURLRequestFromID(objc.ID(rv))
}
// The original request object passed when the task was created.
//
// # Discussion
// 
// This value is typically the same as the currently active request
// ([CurrentRequest]) except when the server has responded to the initial
// request with a redirect to a different URL.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/originalRequest
func (u URLSessionTask) OriginalRequest() INSURLRequest {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("originalRequest"))
	return NSURLRequestFromID(objc.ID(rv))
}
// The server’s response to the currently active request.
//
// # Discussion
// 
// This object provides information about the request as provided by the
// server. This information always includes the original URL. It may also
// include an expected length, MIME type information, encoding information, a
// suggested filename, or a combination of these.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/response
func (u URLSessionTask) Response() INSURLResponse {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("response"))
	return NSURLResponseFromID(objc.ID(rv))
}
// An app-provided string value for the current task.
//
// # Discussion
// 
// The system doesn’t interpret this value; use it for whatever purpose you
// see fit. For example, you could store a description of the task for
// debugging purposes, or a key to track the task in your own data structures.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/taskDescription
func (u URLSessionTask) TaskDescription() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("taskDescription"))
	return NSStringFromID(rv).String()
}
func (u URLSessionTask) SetTaskDescription(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setTaskDescription:"), objc.String(value))
}
// An identifier uniquely identifying the task within a given session.
//
// # Discussion
// 
// This value is unique only within the context of a single session; tasks in
// other sessions may have the same [TaskIdentifier] value.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/taskIdentifier
func (u URLSessionTask) TaskIdentifier() uint {
	rv := objc.Send[uint](u.ID, objc.Sel("taskIdentifier"))
	return rv
}
// An error object that indicates why the task failed.
//
// # Discussion
// 
// This value is `nil` if the task is still active or if the transfer
// completed successfully.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/error
func (u URLSessionTask) Error() INSError {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("error"))
	return NSErrorFromID(objc.ID(rv))
}
// A Boolean value that determines whether to deliver a partial response body
// in increments.
//
// # Discussion
// 
// Set this property to `true` to tell the task that the app would benefit
// from receiving a partial response body in increments. If the app can’t
// process the response until it has all the data, set this property to
// `false`. Task performance may improve when this value is `false`, in which
// case the task only delivers data when complete.
// 
// This property defaults to `true`, except in the following cases which
// default to `false`:
// 
// - The task delivers results to a completion handler rather than to a
// delegate. - The task is a download task.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/prefersIncrementalDelivery
func (u URLSessionTask) PrefersIncrementalDelivery() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("prefersIncrementalDelivery"))
	return rv
}
func (u URLSessionTask) SetPrefersIncrementalDelivery(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setPrefersIncrementalDelivery:"), value)
}
// A delegate specific to the task.
//
// # Discussion
// 
// This task-specific delegate receives messages from the task before the
// session’s [Delegate] receives them. This is similar to the behavior of
// the `delegate` parameter used by the asychronous methods in [NSURLSession]
// like [bytes(for:delegate:)] and [data(for:delegate:)].
//
// [bytes(for:delegate:)]: https://developer.apple.com/documentation/Foundation/URLSession/bytes(for:delegate:)
// [data(for:delegate:)]: https://developer.apple.com/documentation/Foundation/URLSession/data(for:delegate:)
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/delegate
func (u URLSessionTask) Delegate() NSURLSessionTaskDelegate {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("delegate"))
	return NSURLSessionTaskDelegateObjectFromID(rv)
}
func (u URLSessionTask) SetDelegate(value NSURLSessionTaskDelegate) {
	objc.Send[struct{}](u.ID, objc.Sel("setDelegate:"), value)
}
// A best-guess upper bound on the number of bytes the client expects to
// receive.
//
// # Discussion
// 
// The value set for this property should account for the size of both HTTP
// response headers and the response body. If no value is specified, the
// system uses [NSURLSessionTransferSizeUnknown] instead. This property is
// used by the system to optimize the scheduling of URL session tasks.
// Developers are strongly encouraged to provide an approximate upper bound,
// or an exact byte count, if possible, rather than accept the default.
//
// [NSURLSessionTransferSizeUnknown]: https://developer.apple.com/documentation/Foundation/NSURLSessionTransferSizeUnknown
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesClientExpectsToReceive
func (u URLSessionTask) CountOfBytesClientExpectsToReceive() int64 {
	rv := objc.Send[int64](u.ID, objc.Sel("countOfBytesClientExpectsToReceive"))
	return rv
}
func (u URLSessionTask) SetCountOfBytesClientExpectsToReceive(value int64) {
	objc.Send[struct{}](u.ID, objc.Sel("setCountOfBytesClientExpectsToReceive:"), value)
}
// A best-guess upper bound on the number of bytes the client expects to send.
//
// # Discussion
// 
// The value set for this property should account for the size of HTTP headers
// and body data or body stream. If no value is specified, the system uses
// [NSURLSessionTransferSizeUnknown] instead. This property is used by the
// system to optimize the scheduling of URL session tasks. Developers are
// strongly encouraged to provide an approximate upper bound, or an exact byte
// count, if possible, rather than accept the default.
//
// [NSURLSessionTransferSizeUnknown]: https://developer.apple.com/documentation/Foundation/NSURLSessionTransferSizeUnknown
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesClientExpectsToSend
func (u URLSessionTask) CountOfBytesClientExpectsToSend() int64 {
	rv := objc.Send[int64](u.ID, objc.Sel("countOfBytesClientExpectsToSend"))
	return rv
}
func (u URLSessionTask) SetCountOfBytesClientExpectsToSend(value int64) {
	objc.Send[struct{}](u.ID, objc.Sel("setCountOfBytesClientExpectsToSend:"), value)
}
// The total size of the transfer cannot be determined.
//
// See: https://developer.apple.com/documentation/foundation/nsurlsessiontransfersizeunknown
func (u URLSessionTask) NSURLSessionTransferSizeUnknown() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("NSURLSessionTransferSizeUnknown"))
	return objectivec.Object{ID: rv}
}
// The earliest date at which the network load should begin.
//
// # Discussion
// 
// For tasks created from background [NSURLSession] instances, this property
// indicates that the network load should not begin any earlier than this
// date. Setting this property does not guarantee that the load will begin at
// the specified date, but only that it will not begin sooner. If not
// specified, no start delay is used.
// 
// This property has no effect for tasks created from nonbackground sessions.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/earliestBeginDate
func (u URLSessionTask) EarliestBeginDate() INSDate {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("earliestBeginDate"))
	return NSDateFromID(objc.ID(rv))
}
func (u URLSessionTask) SetEarliestBeginDate(value INSDate) {
	objc.Send[struct{}](u.ID, objc.Sel("setEarliestBeginDate:"), value)
}

			// Protocol methods for NSCopying
			

