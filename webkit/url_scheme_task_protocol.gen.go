// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An interface that WebKit uses to request custom resources from your app.
//
// See: https://developer.apple.com/documentation/WebKit/WKURLSchemeTask
type WKURLSchemeTask interface {
	objectivec.IObject

	// Information about the resource to load.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKURLSchemeTask/request
	Request() foundation.NSURLRequest

	// Returns a URL response to WebKit with information about the requested resource.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKURLSchemeTask/didReceive(_:)-2u23r
	DidReceiveResponse(response foundation.NSURLResponse)

	// Sends some or all of the resource data to WebKit.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKURLSchemeTask/didReceive(_:)-8t5f8
	DidReceiveData(data foundation.INSData)

	// Signals the successful completion of the task.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKURLSchemeTask/didFinish()
	DidFinish()

	// Completes the task and reports the specified error back to WebKit.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKURLSchemeTask/didFailWithError(_:)
	DidFailWithError(error_ foundation.INSError)
}

// WKURLSchemeTaskObject wraps an existing Objective-C object that conforms to the WKURLSchemeTask protocol.
type WKURLSchemeTaskObject struct {
	objectivec.Object
}

func (o WKURLSchemeTaskObject) BaseObject() objectivec.Object {
	return o.Object
}

// WKURLSchemeTaskObjectFromID constructs a [WKURLSchemeTaskObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func WKURLSchemeTaskObjectFromID(id objc.ID) WKURLSchemeTaskObject {
	return WKURLSchemeTaskObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Information about the resource to load.
//
// See: https://developer.apple.com/documentation/WebKit/WKURLSchemeTask/request
func (o WKURLSchemeTaskObject) Request() foundation.NSURLRequest {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("request"))
	return foundation.NSURLRequestFromID(rv)
}

// Returns a URL response to WebKit with information about the requested
// resource.
//
// response: The response to return to WebKit. Your response object must include the
// MIME type of the requested resource.
//
// # Discussion
//
// Call this method to provide WebKit with the MIME type of the requested
// resource and its expected size. You must call this method at least once for
// each task, and you may call it multiple times if needed. Always call it
// before sending any data back to WebKit using the [DidReceiveData] method.
//
// It is a programmer error to call this method after calling the [DidFinish]
// method of the same task object. It is also a programmer error to call this
// method after WebKit calls the [WebViewStopURLSchemeTask] method of the
// corresponding handler object. If you do, this method raises an exception in
// both cases.
//
// See: https://developer.apple.com/documentation/WebKit/WKURLSchemeTask/didReceive(_:)-2u23r
func (o WKURLSchemeTaskObject) DidReceiveResponse(response foundation.NSURLResponse) {
	objc.Send[struct{}](o.ID, objc.Sel("didReceiveResponse:"), response)
}

// Sends some or all of the resource data to WebKit.
//
// data: Data for the resource. This object may contain all of the data or only some
// of it.
//
// # Discussion
//
// Call this method to deliver any resource data back to WebKit. If you load
// the data incrementally, you may call this method multiple times to deliver
// each new portion of the data. Each time you call this method, WebKit
// appends the new data to any previously received data.
//
// This method raises an exception if you call it before the
// [DidReceiveResponse] method, or after the [DidFinish] method. It also
// raises an exception if you call it after WebKit calls the
// [WebViewStopURLSchemeTask] method of the corresponding handler object.
//
// See: https://developer.apple.com/documentation/WebKit/WKURLSchemeTask/didReceive(_:)-8t5f8
func (o WKURLSchemeTaskObject) DidReceiveData(data foundation.INSData) {
	objc.Send[struct{}](o.ID, objc.Sel("didReceiveData:"), data)
}

// Signals the successful completion of the task.
//
// # Discussion
//
// This method signals to WebKit that it has all of the resource’s data and
// the task is now complete. Call this method after sending a response and the
// resource data to WebKit using the [DidReceiveResponse] and [DidReceiveData]
// methods.
//
// This method raises an exception if you call it before the
// [DidReceiveResponse] method, or if the task is already complete. It also
// raises an exception if you call it after WebKit calls the
// [WebViewStopURLSchemeTask] method of the corresponding handler object.
//
// See: https://developer.apple.com/documentation/WebKit/WKURLSchemeTask/didFinish()
func (o WKURLSchemeTaskObject) DidFinish() {
	objc.Send[struct{}](o.ID, objc.Sel("didFinish"))
}

// Completes the task and reports the specified error back to WebKit.
//
// error: The error that caused the task to fail.
//
// # Discussion
//
// This method signals to WebKit that the task is complete, but failed with
// the specified error. Use this method to report any errors during the load
// process.
//
// After calling this method, it’s a programmer error to call other methods
// of the task object, and those methods raise an exception if you do. This
// method raises an exception if you call it after WebKit calls the
// [WebViewStopURLSchemeTask] method of the corresponding handler object.
//
// See: https://developer.apple.com/documentation/WebKit/WKURLSchemeTask/didFailWithError(_:)
func (o WKURLSchemeTaskObject) DidFailWithError(error_ foundation.INSError) {
	objc.Send[struct{}](o.ID, objc.Sel("didFailWithError:"), error_)
}
