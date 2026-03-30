// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetResourceLoadingRequest] class.
var (
	_AVAssetResourceLoadingRequestClass     AVAssetResourceLoadingRequestClass
	_AVAssetResourceLoadingRequestClassOnce sync.Once
)

func getAVAssetResourceLoadingRequestClass() AVAssetResourceLoadingRequestClass {
	_AVAssetResourceLoadingRequestClassOnce.Do(func() {
		_AVAssetResourceLoadingRequestClass = AVAssetResourceLoadingRequestClass{class: objc.GetClass("AVAssetResourceLoadingRequest")}
	})
	return _AVAssetResourceLoadingRequestClass
}

// GetAVAssetResourceLoadingRequestClass returns the class object for AVAssetResourceLoadingRequest.
func GetAVAssetResourceLoadingRequestClass() AVAssetResourceLoadingRequestClass {
	return getAVAssetResourceLoadingRequestClass()
}

type AVAssetResourceLoadingRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetResourceLoadingRequestClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetResourceLoadingRequestClass) Alloc() AVAssetResourceLoadingRequest {
	rv := objc.Send[AVAssetResourceLoadingRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that encapsulates information about a resource request from a
// resource loader object.
//
// # Overview
//
// When an [AVURLAsset] object needs help loading a resource, it asks its
// [AVAssetResourceLoader] object to assist. The resource loader encapsulates
// the request information by creating an instance of this object, which it
// then hands to its delegate object for processing. The delegate uses the
// information in this object to perform the request and report on the success
// or failure of the operation.
//
// # Accessing the request data
//
//   - [AVAssetResourceLoadingRequest.Request]: The URL request object for the resource.
//   - [AVAssetResourceLoadingRequest.Requestor]: The asset resource requestor that made the request.
//   - [AVAssetResourceLoadingRequest.ContentInformationRequest]: The information for a requested resource.
//   - [AVAssetResourceLoadingRequest.DataRequest]: The range of requested resource data.
//   - [AVAssetResourceLoadingRequest.Redirect]: An URL request instance if the loading request was redirected.
//   - [AVAssetResourceLoadingRequest.SetRedirect]
//   - [AVAssetResourceLoadingRequest.AVAssetResourceLoadingRequestStreamingContentKeyRequestRequiresPersistentKey]: Specifies whether the content key request requires a persistable key to be returned from the key vendor.
//
// # Reporting the result of the request
//
//   - [AVAssetResourceLoadingRequest.Response]: The URL response for the loading request.
//   - [AVAssetResourceLoadingRequest.SetResponse]
//   - [AVAssetResourceLoadingRequest.FinishLoading]: Causes the receiver to treat the processing of the request as complete.
//   - [AVAssetResourceLoadingRequest.Cancelled]: A Boolean value that indicates whether the request has been cancelled.
//   - [AVAssetResourceLoadingRequest.FinishLoadingWithError]: Causes the receiver to handle the failure to load a resource for which a resource loader’s delegate took responsibility.
//   - [AVAssetResourceLoadingRequest.Finished]: A Boolean value that indicates whether loading of the resource has finished.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest
type AVAssetResourceLoadingRequest struct {
	objectivec.Object
}

// AVAssetResourceLoadingRequestFromID constructs a [AVAssetResourceLoadingRequest] from an objc.ID.
//
// An object that encapsulates information about a resource request from a
// resource loader object.
func AVAssetResourceLoadingRequestFromID(id objc.ID) AVAssetResourceLoadingRequest {
	return AVAssetResourceLoadingRequest{objectivec.Object{ID: id}}
}

// NOTE: AVAssetResourceLoadingRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetResourceLoadingRequest] class.
//
// # Accessing the request data
//
//   - [IAVAssetResourceLoadingRequest.Request]: The URL request object for the resource.
//   - [IAVAssetResourceLoadingRequest.Requestor]: The asset resource requestor that made the request.
//   - [IAVAssetResourceLoadingRequest.ContentInformationRequest]: The information for a requested resource.
//   - [IAVAssetResourceLoadingRequest.DataRequest]: The range of requested resource data.
//   - [IAVAssetResourceLoadingRequest.Redirect]: An URL request instance if the loading request was redirected.
//   - [IAVAssetResourceLoadingRequest.SetRedirect]
//   - [IAVAssetResourceLoadingRequest.AVAssetResourceLoadingRequestStreamingContentKeyRequestRequiresPersistentKey]: Specifies whether the content key request requires a persistable key to be returned from the key vendor.
//
// # Reporting the result of the request
//
//   - [IAVAssetResourceLoadingRequest.Response]: The URL response for the loading request.
//   - [IAVAssetResourceLoadingRequest.SetResponse]
//   - [IAVAssetResourceLoadingRequest.FinishLoading]: Causes the receiver to treat the processing of the request as complete.
//   - [IAVAssetResourceLoadingRequest.Cancelled]: A Boolean value that indicates whether the request has been cancelled.
//   - [IAVAssetResourceLoadingRequest.FinishLoadingWithError]: Causes the receiver to handle the failure to load a resource for which a resource loader’s delegate took responsibility.
//   - [IAVAssetResourceLoadingRequest.Finished]: A Boolean value that indicates whether loading of the resource has finished.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest
type IAVAssetResourceLoadingRequest interface {
	objectivec.IObject

	// Topic: Accessing the request data

	// The URL request object for the resource.
	Request() foundation.NSURLRequest
	// The asset resource requestor that made the request.
	Requestor() IAVAssetResourceLoadingRequestor
	// The information for a requested resource.
	ContentInformationRequest() IAVAssetResourceLoadingContentInformationRequest
	// The range of requested resource data.
	DataRequest() IAVAssetResourceLoadingDataRequest
	// An URL request instance if the loading request was redirected.
	Redirect() foundation.NSURLRequest
	SetRedirect(value foundation.NSURLRequest)
	// Specifies whether the content key request requires a persistable key to be returned from the key vendor.
	AVAssetResourceLoadingRequestStreamingContentKeyRequestRequiresPersistentKey() string

	// Topic: Reporting the result of the request

	// The URL response for the loading request.
	Response() foundation.NSURLResponse
	SetResponse(value foundation.NSURLResponse)
	// Causes the receiver to treat the processing of the request as complete.
	FinishLoading()
	// A Boolean value that indicates whether the request has been cancelled.
	Cancelled() bool
	// Causes the receiver to handle the failure to load a resource for which a resource loader’s delegate took responsibility.
	FinishLoadingWithError(error_ foundation.INSError)
	// A Boolean value that indicates whether loading of the resource has finished.
	Finished() bool
}

// Init initializes the instance.
func (a AVAssetResourceLoadingRequest) Init() AVAssetResourceLoadingRequest {
	rv := objc.Send[AVAssetResourceLoadingRequest](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetResourceLoadingRequest) Autorelease() AVAssetResourceLoadingRequest {
	rv := objc.Send[AVAssetResourceLoadingRequest](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetResourceLoadingRequest creates a new AVAssetResourceLoadingRequest instance.
func NewAVAssetResourceLoadingRequest() AVAssetResourceLoadingRequest {
	class := getAVAssetResourceLoadingRequestClass()
	rv := objc.Send[AVAssetResourceLoadingRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Causes the receiver to treat the processing of the request as complete.
//
// # Discussion
//
// If a [DataRequest] is present and the resource does not contain the full
// extent of the data that has been requested according to the values of the
// [RequestedOffset] and [RequestedLength] properties of the request, invoke
// `finishLoading` after providing as much of the requested data as the
// resource contains.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/finishLoading()
func (a AVAssetResourceLoadingRequest) FinishLoading() {
	objc.Send[objc.ID](a.ID, objc.Sel("finishLoading"))
}

// Causes the receiver to handle the failure to load a resource for which a
// resource loader’s delegate took responsibility.
//
// error: An error object indicating the reason for the failure.
//
// # Discussion
//
// When a resource loader’s delegate takes responsibility for loading a
// resource, it calls this method when a failure occurred when loading the
// resource. This method marks the loading request as finished and notifies
// the resource loader object that the resource could not be loaded.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/finishLoading(with:)
func (a AVAssetResourceLoadingRequest) FinishLoadingWithError(error_ foundation.INSError) {
	objc.Send[objc.ID](a.ID, objc.Sel("finishLoadingWithError:"), error_)
}

// The URL request object for the resource.
//
// # Discussion
//
// Use the value in this property to identify the requested resource and to
// formulate an appropriate response object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/request
func (a AVAssetResourceLoadingRequest) Request() foundation.NSURLRequest {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("request"))
	return foundation.NSURLRequestFromID(objc.ID(rv))
}

// The asset resource requestor that made the request.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/requestor
func (a AVAssetResourceLoadingRequest) Requestor() IAVAssetResourceLoadingRequestor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("requestor"))
	return AVAssetResourceLoadingRequestorFromID(objc.ID(rv))
}

// The information for a requested resource.
//
// # Discussion
//
// An instance of [AVAssetResourceLoadingContentInformationRequest] that you
// populate with information about the resource. The value of this property is
// `nil` if no such information is being requested.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/contentInformationRequest
func (a AVAssetResourceLoadingRequest) ContentInformationRequest() IAVAssetResourceLoadingContentInformationRequest {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("contentInformationRequest"))
	return AVAssetResourceLoadingContentInformationRequestFromID(objc.ID(rv))
}

// The range of requested resource data.
//
// # Discussion
//
// An instance of [AVAssetResourceLoadingDataRequest] that indicates the range
// of resource data that’s being requested. The value of this property is
// `nil` if no data is being requested.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/dataRequest
func (a AVAssetResourceLoadingRequest) DataRequest() IAVAssetResourceLoadingDataRequest {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("dataRequest"))
	return AVAssetResourceLoadingDataRequestFromID(objc.ID(rv))
}

// An URL request instance if the loading request was redirected.
//
// # Discussion
//
// Set this property to an instance of [NSURLRequest] indicating a redirection
// of the loading request to another URL.
//
// If no redirection is needed, the value of this property must be `nil`,
// which is the default.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/redirect
//
// [NSURLRequest]: https://developer.apple.com/documentation/Foundation/NSURLRequest
func (a AVAssetResourceLoadingRequest) Redirect() foundation.NSURLRequest {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("redirect"))
	return foundation.NSURLRequestFromID(objc.ID(rv))
}
func (a AVAssetResourceLoadingRequest) SetRedirect(value foundation.NSURLRequest) {
	objc.Send[struct{}](a.ID, objc.Sel("setRedirect:"), value)
}

// Specifies whether the content key request requires a persistable key to be
// returned from the key vendor.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequeststreamingcontentkeyrequestrequirespersistentkey
func (a AVAssetResourceLoadingRequest) AVAssetResourceLoadingRequestStreamingContentKeyRequestRequiresPersistentKey() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVAssetResourceLoadingRequestStreamingContentKeyRequestRequiresPersistentKey"))
	return foundation.NSStringFromID(rv).String()
}

// The URL response for the loading request.
//
// # Discussion
//
// The value of this property is an instance of [URLResponse], indicating a
// response to the loading request. If no response is needed, the value of
// this property is `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/response
//
// [URLResponse]: https://developer.apple.com/documentation/Foundation/URLResponse
func (a AVAssetResourceLoadingRequest) Response() foundation.NSURLResponse {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("response"))
	return foundation.NSURLResponseFromID(objc.ID(rv))
}
func (a AVAssetResourceLoadingRequest) SetResponse(value foundation.NSURLResponse) {
	objc.Send[struct{}](a.ID, objc.Sel("setResponse:"), value)
}

// A Boolean value that indicates whether the request has been cancelled.
//
// # Discussion
//
// true when the resource loader cancels the loading of a request, just prior
// to sending the message [ResourceLoaderDidCancelLoadingRequest] to the
// delegate.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/isCancelled
func (a AVAssetResourceLoadingRequest) Cancelled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isCancelled"))
	return rv
}

// A Boolean value that indicates whether loading of the resource has
// finished.
//
// # Discussion
//
// The value of this property is false initially. The value changes to true
// when the delegate object handling the request calls the
// [finishLoading(with:data:redirect:)] or [FinishLoadingWithError] method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/isFinished
//
// [finishLoading(with:data:redirect:)]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/finishLoading(with:data:redirect:)
func (a AVAssetResourceLoadingRequest) Finished() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isFinished"))
	return rv
}
