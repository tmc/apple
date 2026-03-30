// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [URLSessionDownloadTask] class.
var (
	_URLSessionDownloadTaskClass     URLSessionDownloadTaskClass
	_URLSessionDownloadTaskClassOnce sync.Once
)

func getURLSessionDownloadTaskClass() URLSessionDownloadTaskClass {
	_URLSessionDownloadTaskClassOnce.Do(func() {
		_URLSessionDownloadTaskClass = URLSessionDownloadTaskClass{class: objc.GetClass("NSURLSessionDownloadTask")}
	})
	return _URLSessionDownloadTaskClass
}

// GetURLSessionDownloadTaskClass returns the class object for NSURLSessionDownloadTask.
func GetURLSessionDownloadTaskClass() URLSessionDownloadTaskClass {
	return getURLSessionDownloadTaskClass()
}

type URLSessionDownloadTaskClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc URLSessionDownloadTaskClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc URLSessionDownloadTaskClass) Alloc() URLSessionDownloadTask {
	rv := objc.Send[URLSessionDownloadTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A URL session task that stores downloaded data to a file.
//
// # Overview
//
// An [NSURLSessionDownloadTask] is a concrete subclass of [NSURLSessionTask],
// which provides most of the methods for this class.
//
// Download tasks directly write the server’s response data to a temporary
// file, providing your app with progress updates as data arrives from the
// server. When you use download tasks in background sessions, these downloads
// continue even when your app is in the suspended state or otherwise not
// running.
//
// You can pause (cancel) download tasks and resume them later (assuming the
// server supports doing so). You can also resume downloads that failed
// because of network connectivity problems.
//
// # Download delegate behavior
//
// When you use a download task, your delegate receives several callbacks
// unique to download scenarios.
//
// - During download, the session periodically calls the delegate’s
// [URLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite]
// method with status information. - Upon successful completion, the session
// calls the delegate’s [URLSessionDownloadTaskDidFinishDownloadingToURL]
// method or completion handler. In that method, you must either open the file
// for reading or move it to a permanent location in your app’s sandbox
// container directory. - Upon unsuccessful completion, the session calls the
// delegate’s [URLSessionTaskDidCompleteWithError] method or completion
// handler. The only errors your delegate receives through the `error`
// parameter are client-side errors, such as being unable to resolve the
// hostname or connect to the host. To check for server-side errors, inspect
// the [Response] property of the `task` parameter received by this callback.
//
// # Canceling a download
//
//   - [URLSessionDownloadTask.CancelByProducingResumeData]: Cancels a download and calls a callback with resume data for later use.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionDownloadTask
type URLSessionDownloadTask struct {
	NSURLSessionTask
}

// URLSessionDownloadTaskFromID constructs a [URLSessionDownloadTask] from an objc.ID.
//
// A URL session task that stores downloaded data to a file.
func URLSessionDownloadTaskFromID(id objc.ID) URLSessionDownloadTask {
	return NSURLSessionDownloadTask{NSURLSessionTask: NSURLSessionTaskFromID(id)}
}

// NSURLSessionDownloadTaskFromID is an alias for [URLSessionDownloadTaskFromID] for cross-framework compatibility.
func NSURLSessionDownloadTaskFromID(id objc.ID) URLSessionDownloadTask {
	return URLSessionDownloadTaskFromID(id)
}

// NOTE: URLSessionDownloadTask adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [URLSessionDownloadTask] class.
//
// # Canceling a download
//
//   - [IURLSessionDownloadTask.CancelByProducingResumeData]: Cancels a download and calls a callback with resume data for later use.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionDownloadTask
type IURLSessionDownloadTask interface {
	INSURLSessionTask

	// Topic: Canceling a download

	// Cancels a download and calls a callback with resume data for later use.
	CancelByProducingResumeData(completionHandler DataHandler)
}

// Init initializes the instance.
func (u URLSessionDownloadTask) Init() URLSessionDownloadTask {
	rv := objc.Send[URLSessionDownloadTask](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u URLSessionDownloadTask) Autorelease() URLSessionDownloadTask {
	rv := objc.Send[URLSessionDownloadTask](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionDownloadTask creates a new URLSessionDownloadTask instance.
func NewURLSessionDownloadTask() URLSessionDownloadTask {
	class := getURLSessionDownloadTaskClass()
	rv := objc.Send[URLSessionDownloadTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Cancels a download and calls a callback with resume data for later use.
//
// completionHandler: A completion handler that is called when the download has been successfully
// canceled.
//
// If the download is resumable, the completion handler is provided with a
// `resumeData` object. Your app can later pass this object to a session’s
// [DownloadTaskWithResumeData] or
// [DownloadTaskWithResumeDataCompletionHandler] method to create a new task
// that resumes the download where it left off.
//
// This block is not guaranteed to execute in a particular thread context. As
// such, you may want specify an appropriate dispatch queue in which to
// perform any work.
//
// # Discussion
//
// A download can be resumed only if the following conditions are met:
//
// - The resource has not changed since you first requested it - The task is
// an HTTP or HTTPS [GET] request - The server provides either the [ETag] or
// `Last-Modified` header (or both) in its response - The server supports
// byte-range requests - The temporary file hasn’t been deleted by the
// system in response to disk space pressure
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionDownloadTask/cancel(byProducingResumeData:)
func (u URLSessionDownloadTask) CancelByProducingResumeData(completionHandler DataHandler) {
	_block0, _ := NewDataBlock(completionHandler)
	objc.Send[objc.ID](u.ID, objc.Sel("cancelByProducingResumeData:"), _block0)
}

// CancelByProducingResumeDataSync is a synchronous wrapper around [URLSessionDownloadTask.CancelByProducingResumeData].
// It blocks until the completion handler fires or the context is cancelled.
func (u URLSessionDownloadTask) CancelByProducingResumeDataSync(ctx context.Context) (*NSData, error) {
	done := make(chan *NSData, 1)
	u.CancelByProducingResumeData(func(val *NSData) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
