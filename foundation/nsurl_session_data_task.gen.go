// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [URLSessionDataTask] class.
var (
	_URLSessionDataTaskClass     URLSessionDataTaskClass
	_URLSessionDataTaskClassOnce sync.Once
)

func getURLSessionDataTaskClass() URLSessionDataTaskClass {
	_URLSessionDataTaskClassOnce.Do(func() {
		_URLSessionDataTaskClass = URLSessionDataTaskClass{class: objc.GetClass("NSURLSessionDataTask")}
	})
	return _URLSessionDataTaskClass
}

// GetURLSessionDataTaskClass returns the class object for NSURLSessionDataTask.
func GetURLSessionDataTaskClass() URLSessionDataTaskClass {
	return getURLSessionDataTaskClass()
}

type URLSessionDataTaskClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc URLSessionDataTaskClass) Alloc() URLSessionDataTask {
	rv := objc.Send[URLSessionDataTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}







// A URL session task that returns downloaded data directly to the app in
// memory.
//
// # Overview
// 
// A [NSURLSessionDataTask] is a concrete subclass of [NSURLSessionTask]. The
// methods in the [NSURLSessionDataTask] class are documented in
// [NSURLSessionTask].
// 
// A data task returns data directly to the app (in memory) as one or more
// [NSData] objects. When you use a data task:
// 
// - During upload of the body data (if your app provides any), the session
// periodically calls its delegate’s
// [URLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend]
// method with status information. - After receiving an initial response, the
// session calls its delegate’s
// [URLSessionDataTaskDidReceiveResponseCompletionHandler] method to let you
// examine the status code and headers, and optionally convert the data task
// into a download task. - During the transfer, the session calls its
// delegate’s [URLSessionDataTaskDidReceiveData] method to provide your app
// with the content as it arrives. - Upon completion, the session calls its
// delegate’s [URLSessionDataTaskWillCacheResponseCompletionHandler] method
// to let you determine whether the response should be cached.
// 
// For examples of using data tasks for fetching and uploading data, see
// [Fetching website data into memory] and [Uploading data to a website].
//
// [Fetching website data into memory]: https://developer.apple.com/documentation/Foundation/fetching-website-data-into-memory
// [Uploading data to a website]: https://developer.apple.com/documentation/Foundation/uploading-data-to-a-website
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionDataTask
type URLSessionDataTask struct {
	NSURLSessionTask
}

// URLSessionDataTaskFromID constructs a [URLSessionDataTask] from an objc.ID.
//
// A URL session task that returns downloaded data directly to the app in
// memory.
func URLSessionDataTaskFromID(id objc.ID) URLSessionDataTask {
	return NSURLSessionDataTask{NSURLSessionTask: NSURLSessionTaskFromID(id)}
}

// NSURLSessionDataTaskFromID is an alias for [URLSessionDataTaskFromID] for cross-framework compatibility.
func NSURLSessionDataTaskFromID(id objc.ID) URLSessionDataTask { return URLSessionDataTaskFromID(id) }
// NOTE: URLSessionDataTask adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [URLSessionDataTask] class.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionDataTask
type IURLSessionDataTask interface {
	INSURLSessionTask
}





// Init initializes the instance.
func (u URLSessionDataTask) Init() URLSessionDataTask {
	rv := objc.Send[URLSessionDataTask](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u URLSessionDataTask) Autorelease() URLSessionDataTask {
	rv := objc.Send[URLSessionDataTask](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionDataTask creates a new URLSessionDataTask instance.
func NewURLSessionDataTask() URLSessionDataTask {
	class := getURLSessionDataTaskClass()
	rv := objc.Send[URLSessionDataTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}













































