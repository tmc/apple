// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetResourceLoadingContentInformationRequest] class.
var (
	_AVAssetResourceLoadingContentInformationRequestClass     AVAssetResourceLoadingContentInformationRequestClass
	_AVAssetResourceLoadingContentInformationRequestClassOnce sync.Once
)

func getAVAssetResourceLoadingContentInformationRequestClass() AVAssetResourceLoadingContentInformationRequestClass {
	_AVAssetResourceLoadingContentInformationRequestClassOnce.Do(func() {
		_AVAssetResourceLoadingContentInformationRequestClass = AVAssetResourceLoadingContentInformationRequestClass{class: objc.GetClass("AVAssetResourceLoadingContentInformationRequest")}
	})
	return _AVAssetResourceLoadingContentInformationRequestClass
}

// GetAVAssetResourceLoadingContentInformationRequestClass returns the class object for AVAssetResourceLoadingContentInformationRequest.
func GetAVAssetResourceLoadingContentInformationRequestClass() AVAssetResourceLoadingContentInformationRequestClass {
	return getAVAssetResourceLoadingContentInformationRequestClass()
}

type AVAssetResourceLoadingContentInformationRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetResourceLoadingContentInformationRequestClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetResourceLoadingContentInformationRequestClass) Alloc() AVAssetResourceLoadingContentInformationRequest {
	rv := objc.Send[AVAssetResourceLoadingContentInformationRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A query for retrieving essential information about a resource that an asset
// resource-loading request references.
//
// # Overview
// 
// When a resource loading delegate, which must implement the
// [AVAssetResourceLoaderDelegate] protocol, receives an instance of
// [AVAssetResourceLoadingRequest] when the
// [ResourceLoaderShouldWaitForLoadingOfRequestedResource] is invoked and
// accepts responsibility for loading the resource, it must check whether the
// [AVAssetResourceLoadingContentInformationRequest.ContentInformationRequest] property of the [AVAssetResourceLoadingRequest]
// is not `nil`. Whenever the value is not `nil`, the request includes a query
// for the information that [AVAssetResourceLoadingContentInformationRequest]
// encapsulates. In response to such queries, the resource loading delegate
// should set the values of the content information request’s properties
// appropriately before invoking the [AVAssetResourceLoadingRequest] method
// [FinishLoading].
// 
// When [FinishLoading] is invoked, the values of the properties of its
// [AVAssetResourceLoadingContentInformationRequest.ContentInformationRequest] property will, in part, determine how the
// requested resource is processed. For example, if the requested resource’s
// URL is the URL of an [AVURLAsset] and [AVAssetResourceLoadingContentInformationRequest.ContentType] is set by the resource
// loading delegate to a value that the underlying media system doesn’t
// recognize as a supported media file type, operations on the [AVURLAsset],
// such as playback, are likely to fail.
//
// # Configuring content information
//
//   - [AVAssetResourceLoadingContentInformationRequest.AllowedContentTypes]: The types of data that are accepted as a valid response for the requested resource.
//   - [AVAssetResourceLoadingContentInformationRequest.ContentType]: The UTI that specifies the type of data contained by the requested resource.
//   - [AVAssetResourceLoadingContentInformationRequest.SetContentType]
//   - [AVAssetResourceLoadingContentInformationRequest.ContentLength]: The length, in bytes, of the requested resource.
//   - [AVAssetResourceLoadingContentInformationRequest.SetContentLength]
//   - [AVAssetResourceLoadingContentInformationRequest.ByteRangeAccessSupported]: A Boolean value that indicates whether random access to arbitrary ranges of bytes of the resource is supported.
//   - [AVAssetResourceLoadingContentInformationRequest.SetByteRangeAccessSupported]
//   - [AVAssetResourceLoadingContentInformationRequest.RenewalDate]: The date at which a new resource loading request will be issued for resources that expire, if the media system still requires it.
//   - [AVAssetResourceLoadingContentInformationRequest.SetRenewalDate]
//   - [AVAssetResourceLoadingContentInformationRequest.EntireLengthAvailableOnDemand]: A Boolean value that indicates whether asset data loading can expect data immediately.
//   - [AVAssetResourceLoadingContentInformationRequest.SetEntireLengthAvailableOnDemand]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingContentInformationRequest
type AVAssetResourceLoadingContentInformationRequest struct {
	objectivec.Object
}

// AVAssetResourceLoadingContentInformationRequestFromID constructs a [AVAssetResourceLoadingContentInformationRequest] from an objc.ID.
//
// A query for retrieving essential information about a resource that an asset
// resource-loading request references.
func AVAssetResourceLoadingContentInformationRequestFromID(id objc.ID) AVAssetResourceLoadingContentInformationRequest {
	return AVAssetResourceLoadingContentInformationRequest{objectivec.Object{ID: id}}
}
// NOTE: AVAssetResourceLoadingContentInformationRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetResourceLoadingContentInformationRequest] class.
//
// # Configuring content information
//
//   - [IAVAssetResourceLoadingContentInformationRequest.AllowedContentTypes]: The types of data that are accepted as a valid response for the requested resource.
//   - [IAVAssetResourceLoadingContentInformationRequest.ContentType]: The UTI that specifies the type of data contained by the requested resource.
//   - [IAVAssetResourceLoadingContentInformationRequest.SetContentType]
//   - [IAVAssetResourceLoadingContentInformationRequest.ContentLength]: The length, in bytes, of the requested resource.
//   - [IAVAssetResourceLoadingContentInformationRequest.SetContentLength]
//   - [IAVAssetResourceLoadingContentInformationRequest.ByteRangeAccessSupported]: A Boolean value that indicates whether random access to arbitrary ranges of bytes of the resource is supported.
//   - [IAVAssetResourceLoadingContentInformationRequest.SetByteRangeAccessSupported]
//   - [IAVAssetResourceLoadingContentInformationRequest.RenewalDate]: The date at which a new resource loading request will be issued for resources that expire, if the media system still requires it.
//   - [IAVAssetResourceLoadingContentInformationRequest.SetRenewalDate]
//   - [IAVAssetResourceLoadingContentInformationRequest.EntireLengthAvailableOnDemand]: A Boolean value that indicates whether asset data loading can expect data immediately.
//   - [IAVAssetResourceLoadingContentInformationRequest.SetEntireLengthAvailableOnDemand]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingContentInformationRequest
type IAVAssetResourceLoadingContentInformationRequest interface {
	objectivec.IObject

	// Topic: Configuring content information

	// The types of data that are accepted as a valid response for the requested resource.
	AllowedContentTypes() []string
	// The UTI that specifies the type of data contained by the requested resource.
	ContentType() string
	SetContentType(value string)
	// The length, in bytes, of the requested resource.
	ContentLength() int64
	SetContentLength(value int64)
	// A Boolean value that indicates whether random access to arbitrary ranges of bytes of the resource is supported.
	ByteRangeAccessSupported() bool
	SetByteRangeAccessSupported(value bool)
	// The date at which a new resource loading request will be issued for resources that expire, if the media system still requires it.
	RenewalDate() foundation.INSDate
	SetRenewalDate(value foundation.INSDate)
	// A Boolean value that indicates whether asset data loading can expect data immediately.
	EntireLengthAvailableOnDemand() bool
	SetEntireLengthAvailableOnDemand(value bool)

	// The information for a requested resource.
	ContentInformationRequest() IAVAssetResourceLoadingContentInformationRequest
	SetContentInformationRequest(value IAVAssetResourceLoadingContentInformationRequest)
}

// Init initializes the instance.
func (a AVAssetResourceLoadingContentInformationRequest) Init() AVAssetResourceLoadingContentInformationRequest {
	rv := objc.Send[AVAssetResourceLoadingContentInformationRequest](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetResourceLoadingContentInformationRequest) Autorelease() AVAssetResourceLoadingContentInformationRequest {
	rv := objc.Send[AVAssetResourceLoadingContentInformationRequest](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetResourceLoadingContentInformationRequest creates a new AVAssetResourceLoadingContentInformationRequest instance.
func NewAVAssetResourceLoadingContentInformationRequest() AVAssetResourceLoadingContentInformationRequest {
	class := getAVAssetResourceLoadingContentInformationRequestClass()
	rv := objc.Send[AVAssetResourceLoadingContentInformationRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The types of data that are accepted as a valid response for the requested
// resource.
//
// # Discussion
// 
// This property contains an array of file format UTIs. When
// `allowedContentTypes` is non-nil, the value of [ContentType] must be set to
// a value contained in `allowedContentTypes` or `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingContentInformationRequest/allowedContentTypes
func (a AVAssetResourceLoadingContentInformationRequest) AllowedContentTypes() []string {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("allowedContentTypes"))
	return objc.ConvertSliceToStrings(rv)
}
// The UTI that specifies the type of data contained by the requested
// resource.
//
// # Discussion
// 
// Before finishing loading an [AVAssetResourceLoadingRequest] instance, if
// its [ContentInformationRequest] property is not `nil`, set the value of
// this property to a UTI indicating the type of data contained by the
// requested resource.
// 
// When responding to an [AVAssetResourceLoadingRequest] for a FairPlay
// Streaming key, only set `contentType` to
// [AVStreamingKeyDeliveryContentKeyType],
// [AVStreamingKeyDeliveryPersistentContentKeyType], or `nil`. The value of
// contentType must be contained in the [AllowedContentTypes] property or
// `nil`.
//
// [AVStreamingKeyDeliveryContentKeyType]: https://developer.apple.com/documentation/AVFoundation/AVStreamingKeyDeliveryContentKeyType
// [AVStreamingKeyDeliveryPersistentContentKeyType]: https://developer.apple.com/documentation/AVFoundation/AVStreamingKeyDeliveryPersistentContentKeyType
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingContentInformationRequest/contentType
func (a AVAssetResourceLoadingContentInformationRequest) ContentType() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("contentType"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVAssetResourceLoadingContentInformationRequest) SetContentType(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setContentType:"), objc.String(value))
}
// The length, in bytes, of the requested resource.
//
// # Discussion
// 
// Before finishing loading an [AVAssetResourceLoadingRequest] instance, if
// its [ContentInformationRequest] property is not `nil`, set the value of the
// `contentLength` property to the number of bytes contained by the requested
// resource.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingContentInformationRequest/contentLength
func (a AVAssetResourceLoadingContentInformationRequest) ContentLength() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("contentLength"))
	return rv
}
func (a AVAssetResourceLoadingContentInformationRequest) SetContentLength(value int64) {
	objc.Send[struct{}](a.ID, objc.Sel("setContentLength:"), value)
}
// A Boolean value that indicates whether random access to arbitrary ranges of
// bytes of the resource is supported.
//
// # Discussion
// 
// Before finishing loading an [AVAssetResourceLoadingRequest] instance, if
// its [ContentInformationRequest] property is not `nil`, set the value of
// this property to [true] if it supports random access to arbitrary ranges of
// bytes of the resource.
// 
// If this property is not [true] for resources that must be loaded
// incrementally, loading of the resource may fail. Such resources include
// anything that contains media data.
// 
// If byte range access is supported portions of the resource can be requested
// more than once.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingContentInformationRequest/isByteRangeAccessSupported
func (a AVAssetResourceLoadingContentInformationRequest) ByteRangeAccessSupported() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isByteRangeAccessSupported"))
	return rv
}
func (a AVAssetResourceLoadingContentInformationRequest) SetByteRangeAccessSupported(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setByteRangeAccessSupported:"), value)
}
// The date at which a new resource loading request will be issued for
// resources that expire, if the media system still requires it.
//
// # Discussion
// 
// If the asset resource is prone to expiry set the value of this property to
// the date at which a renewal should be triggered. You must do this before
// you finish loading an [AVAssetResourceLoadingRequest] object. This value
// must be set sufficiently early enough to allow an
// [AVAssetResourceRenewalRequest], delivered to the delegate’s `` method ß
// to finish before the actual expiry time, otherwise media playback may fail.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingContentInformationRequest/renewalDate
func (a AVAssetResourceLoadingContentInformationRequest) RenewalDate() foundation.INSDate {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("renewalDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (a AVAssetResourceLoadingContentInformationRequest) SetRenewalDate(value foundation.INSDate) {
	objc.Send[struct{}](a.ID, objc.Sel("setRenewalDate:"), value)
}
// A Boolean value that indicates whether asset data loading can expect data
// immediately.
//
// # Discussion
// 
// Before you finish loading an [AVAssetResourceLoadingRequest], if its
// [ContentInformationRequest] isn’t `nil`, set the value to [true] to
// indicate that all asset data is available. This may be [true] because the
// data is fully cached, or because the custom URL scheme ultimately refers to
// files on local storage, which allows for significant data flow
// optimizations.
// 
// For backward compatibility, this property defaults to [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingContentInformationRequest/isEntireLengthAvailableOnDemand
func (a AVAssetResourceLoadingContentInformationRequest) EntireLengthAvailableOnDemand() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isEntireLengthAvailableOnDemand"))
	return rv
}
func (a AVAssetResourceLoadingContentInformationRequest) SetEntireLengthAvailableOnDemand(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setEntireLengthAvailableOnDemand:"), value)
}
// The information for a requested resource.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/contentinformationrequest
func (a AVAssetResourceLoadingContentInformationRequest) ContentInformationRequest() IAVAssetResourceLoadingContentInformationRequest {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("contentInformationRequest"))
	return AVAssetResourceLoadingContentInformationRequestFromID(objc.ID(rv))
}
func (a AVAssetResourceLoadingContentInformationRequest) SetContentInformationRequest(value IAVAssetResourceLoadingContentInformationRequest) {
	objc.Send[struct{}](a.ID, objc.Sel("setContentInformationRequest:"), value)
}

