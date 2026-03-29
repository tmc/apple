// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [URLSessionUploadTask] class.
var (
	_URLSessionUploadTaskClass     URLSessionUploadTaskClass
	_URLSessionUploadTaskClassOnce sync.Once
)

func getURLSessionUploadTaskClass() URLSessionUploadTaskClass {
	_URLSessionUploadTaskClassOnce.Do(func() {
		_URLSessionUploadTaskClass = URLSessionUploadTaskClass{class: objc.GetClass("NSURLSessionUploadTask")}
	})
	return _URLSessionUploadTaskClass
}

// GetURLSessionUploadTaskClass returns the class object for NSURLSessionUploadTask.
func GetURLSessionUploadTaskClass() URLSessionUploadTaskClass {
	return getURLSessionUploadTaskClass()
}

type URLSessionUploadTaskClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc URLSessionUploadTaskClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc URLSessionUploadTaskClass) Alloc() URLSessionUploadTask {
	rv := objc.Send[URLSessionUploadTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A URL session task that uploads data to the network in a request body.
//
// # Overview
// 
// The [NSURLSessionUploadTask] class is a subclass of [NSURLSessionDataTask],
// which in turn is a concrete subclass of [NSURLSessionTask]. The methods
// associated with the [NSURLSessionUploadTask] class are documented in
// [NSURLSessionTask].
// 
// Upload tasks are used for making HTTP requests that require a request body
// (such as [POST] or [PUT]). They behave similarly to data tasks, but you
// create them by calling different methods on the session that are designed
// to make it easier to provide the content to upload. As with data tasks, if
// the server provides a response, upload tasks return that response as one or
// more [NSData] objects in memory.
// 
// When you create an upload task, you provide a [URLRequest] instance that
// contains any additional headers that you might need to send alongside the
// upload, such as the content type, content disposition, and so on. In iOS,
// when you create an upload task for a file in a background session, the
// system copies that file to a temporary location and streams data from
// there.
// 
// While the upload is in progress, the task calls the session delegate’s
// [URLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend]
// method periodically to provide you with status information.
// 
// When the upload phase of the request finishes, the task behaves like a data
// task, calling methods on the session delegate to provide you with the
// server’s response—headers, status code, content data, and so on.
//
// [URLRequest]: https://developer.apple.com/documentation/Foundation/URLRequest
//
// # Instance Methods
//
//   - [URLSessionUploadTask.CancelByProducingResumeData]: Cancels an upload and calls the completion handler with resume data for later use. resumeData will be nil if the server does not support the latest resumable uploads Internet-Draft from the HTTP Working Group, found at https://datatracker.ietf.org/doc/draft-ietf-httpbis-resumable-upload/
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionUploadTask
type URLSessionUploadTask struct {
	NSURLSessionDataTask
}

// URLSessionUploadTaskFromID constructs a [URLSessionUploadTask] from an objc.ID.
//
// A URL session task that uploads data to the network in a request body.
func URLSessionUploadTaskFromID(id objc.ID) URLSessionUploadTask {
	return NSURLSessionUploadTask{NSURLSessionDataTask: NSURLSessionDataTaskFromID(id)}
}

// NSURLSessionUploadTaskFromID is an alias for [URLSessionUploadTaskFromID] for cross-framework compatibility.
func NSURLSessionUploadTaskFromID(id objc.ID) URLSessionUploadTask { return URLSessionUploadTaskFromID(id) }
// NOTE: URLSessionUploadTask adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [URLSessionUploadTask] class.
//
// # Instance Methods
//
//   - [IURLSessionUploadTask.CancelByProducingResumeData]: Cancels an upload and calls the completion handler with resume data for later use. resumeData will be nil if the server does not support the latest resumable uploads Internet-Draft from the HTTP Working Group, found at https://datatracker.ietf.org/doc/draft-ietf-httpbis-resumable-upload/
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionUploadTask
type IURLSessionUploadTask interface {
	INSURLSessionDataTask

	// Topic: Instance Methods

	// Cancels an upload and calls the completion handler with resume data for later use. resumeData will be nil if the server does not support the latest resumable uploads Internet-Draft from the HTTP Working Group, found at https://datatracker.ietf.org/doc/draft-ietf-httpbis-resumable-upload/
	CancelByProducingResumeData(completionHandler DataHandler)
}

// Init initializes the instance.
func (u URLSessionUploadTask) Init() URLSessionUploadTask {
	rv := objc.Send[URLSessionUploadTask](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u URLSessionUploadTask) Autorelease() URLSessionUploadTask {
	rv := objc.Send[URLSessionUploadTask](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionUploadTask creates a new URLSessionUploadTask instance.
func NewURLSessionUploadTask() URLSessionUploadTask {
	class := getURLSessionUploadTaskClass()
	rv := objc.Send[URLSessionUploadTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Cancels an upload and calls the completion handler with resume data for
// later use. resumeData will be nil if the server does not support the latest
// resumable uploads Internet-Draft from the HTTP Working Group, found at
// https://datatracker.ietf.org/doc/draft-ietf-httpbis-resumable-upload/
//
// completionHandler: The completion handler to call when the upload has been successfully
// canceled.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionUploadTask/cancel(byProducingResumeData:)
func (u URLSessionUploadTask) CancelByProducingResumeData(completionHandler DataHandler) {
_block0, _ := NewDataBlock(completionHandler)
	objc.Send[objc.ID](u.ID, objc.Sel("cancelByProducingResumeData:"), _block0)
}

// CancelByProducingResumeDataSync is a synchronous wrapper around [URLSessionUploadTask.CancelByProducingResumeData].
// It blocks until the completion handler fires or the context is cancelled.
func (u URLSessionUploadTask) CancelByProducingResumeDataSync(ctx context.Context) (*NSData, error) {
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

