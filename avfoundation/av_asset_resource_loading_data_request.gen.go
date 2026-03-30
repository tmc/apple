// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetResourceLoadingDataRequest] class.
var (
	_AVAssetResourceLoadingDataRequestClass     AVAssetResourceLoadingDataRequestClass
	_AVAssetResourceLoadingDataRequestClassOnce sync.Once
)

func getAVAssetResourceLoadingDataRequestClass() AVAssetResourceLoadingDataRequestClass {
	_AVAssetResourceLoadingDataRequestClassOnce.Do(func() {
		_AVAssetResourceLoadingDataRequestClass = AVAssetResourceLoadingDataRequestClass{class: objc.GetClass("AVAssetResourceLoadingDataRequest")}
	})
	return _AVAssetResourceLoadingDataRequestClass
}

// GetAVAssetResourceLoadingDataRequestClass returns the class object for AVAssetResourceLoadingDataRequest.
func GetAVAssetResourceLoadingDataRequestClass() AVAssetResourceLoadingDataRequestClass {
	return getAVAssetResourceLoadingDataRequestClass()
}

type AVAssetResourceLoadingDataRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetResourceLoadingDataRequestClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetResourceLoadingDataRequestClass) Alloc() AVAssetResourceLoadingDataRequest {
	rv := objc.Send[AVAssetResourceLoadingDataRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object for requesting data from a resource that an asset
// resource-loading request references.
//
// # Overview
//
// The [AVAssetResourceLoaderDelegate] uses the
// [AVAssetResourceLoadingDataRequest] class to do the actual data reading,
// and its methods will be invoked, as necessary, to acquire data for the
// [AVAssetResourceLoadingRequest] instance.
//
// When the resource loading delegate, which implements the
// [AVAssetResourceLoaderDelegate] protocol, receives an instance of
// [AVAssetResourceLoadingRequest] as the second parameter of the delegate’s
// [ResourceLoaderShouldWaitForLoadingOfRequestedResource] method, it has the
// option of accepting responsibility for loading the referenced resource. If
// it accepts that responsibility, by returning true, it must check whether
// the [AVAssetResourceLoadingDataRequest.DataRequest] property of the [AVAssetResourceLoadingRequest] instance
// is not `nil`. If it is not `nil`, the resource loading delegate is informed
// of the range of bytes within the resource that are required by the
// underlying media system. In response, the data is provided by one or more
// invocations of [AVAssetResourceLoadingDataRequest.RespondWithData] as required to provide the requested data.
// The data can be provided in increments determined by the resource loading
// delegate according to convenience or efficiency.
//
// When the [AVAssetResourceLoadingRequest] method [FinishLoading] is invoked,
// the data request is considered fully satisfied. If the entire range of
// bytes requested has not yet been provided, the underlying media system
// assumes that the resource’s length is limited to the provided content.
//
// # Providing data to a request
//
//   - [AVAssetResourceLoadingDataRequest.RespondWithData]: Provides data to the loading request.
//   - [AVAssetResourceLoadingDataRequest.RequestedLength]: The length, in bytes, of the data requested.
//   - [AVAssetResourceLoadingDataRequest.RequestedOffset]: The position within the resource of the first byte requested.
//   - [AVAssetResourceLoadingDataRequest.CurrentOffset]: The position within the resource of the next byte.
//   - [AVAssetResourceLoadingDataRequest.RequestsAllDataToEndOfResource]: A Boolean value that indicates the entire remaining length of the resource from the offest to the end of the resource is being requested.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingDataRequest
type AVAssetResourceLoadingDataRequest struct {
	objectivec.Object
}

// AVAssetResourceLoadingDataRequestFromID constructs a [AVAssetResourceLoadingDataRequest] from an objc.ID.
//
// An object for requesting data from a resource that an asset
// resource-loading request references.
func AVAssetResourceLoadingDataRequestFromID(id objc.ID) AVAssetResourceLoadingDataRequest {
	return AVAssetResourceLoadingDataRequest{objectivec.Object{ID: id}}
}

// NOTE: AVAssetResourceLoadingDataRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetResourceLoadingDataRequest] class.
//
// # Providing data to a request
//
//   - [IAVAssetResourceLoadingDataRequest.RespondWithData]: Provides data to the loading request.
//   - [IAVAssetResourceLoadingDataRequest.RequestedLength]: The length, in bytes, of the data requested.
//   - [IAVAssetResourceLoadingDataRequest.RequestedOffset]: The position within the resource of the first byte requested.
//   - [IAVAssetResourceLoadingDataRequest.CurrentOffset]: The position within the resource of the next byte.
//   - [IAVAssetResourceLoadingDataRequest.RequestsAllDataToEndOfResource]: A Boolean value that indicates the entire remaining length of the resource from the offest to the end of the resource is being requested.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingDataRequest
type IAVAssetResourceLoadingDataRequest interface {
	objectivec.IObject

	// Topic: Providing data to a request

	// Provides data to the loading request.
	RespondWithData(data foundation.INSData)
	// The length, in bytes, of the data requested.
	RequestedLength() int
	// The position within the resource of the first byte requested.
	RequestedOffset() int64
	// The position within the resource of the next byte.
	CurrentOffset() int64
	// A Boolean value that indicates the entire remaining length of the resource from the offest to the end of the resource is being requested.
	RequestsAllDataToEndOfResource() bool

	// The range of requested resource data.
	DataRequest() IAVAssetResourceLoadingDataRequest
	SetDataRequest(value IAVAssetResourceLoadingDataRequest)
}

// Init initializes the instance.
func (a AVAssetResourceLoadingDataRequest) Init() AVAssetResourceLoadingDataRequest {
	rv := objc.Send[AVAssetResourceLoadingDataRequest](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetResourceLoadingDataRequest) Autorelease() AVAssetResourceLoadingDataRequest {
	rv := objc.Send[AVAssetResourceLoadingDataRequest](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetResourceLoadingDataRequest creates a new AVAssetResourceLoadingDataRequest instance.
func NewAVAssetResourceLoadingDataRequest() AVAssetResourceLoadingDataRequest {
	class := getAVAssetResourceLoadingDataRequestClass()
	rv := objc.Send[AVAssetResourceLoadingDataRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Provides data to the loading request.
//
// data: An instance of NSData containing some or all of the requested bytes.
//
// # Discussion
//
// This method may be invoked multiple times on the same instance of
// [AVAssetResourceLoadingDataRequest] to provide the full range of requested
// data incrementally. Upon each invocation, the value of the [CurrentOffset]
// property is updated to match the amount of data provided.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingDataRequest/respond(with:)
func (a AVAssetResourceLoadingDataRequest) RespondWithData(data foundation.INSData) {
	objc.Send[objc.ID](a.ID, objc.Sel("respondWithData:"), data)
}

// The length, in bytes, of the data requested.
//
// # Discussion
//
// If the content length of the resource is unknown, the sum of the
// [RequestedLength] and [RequestedOffset] properties may be greater than the
// actual content length. When this situation occurs, an application must
// attempt to provide as much of the requested data beginning at the
// [RequestedOffset] property as the resource contains. The application must
// then invoke either the [AVAssetResourceLoadingRequest] instance’s
// [FinishLoading] method upon success, or the [FinishLoadingWithError] method
// if an error is encountered during the loading.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingDataRequest/requestedLength
func (a AVAssetResourceLoadingDataRequest) RequestedLength() int {
	rv := objc.Send[int](a.ID, objc.Sel("requestedLength"))
	return rv
}

// The position within the resource of the first byte requested.
//
// # Discussion
//
// When all of the requested bytes that can be provided have been
// loaded—including the possible [ContentInformationRequest] data in the
// [AVAssetResourceLoadingRequest] instance that contains the receiver—the
// delegate should respond by invoking [FinishLoading].
//
// If the `requestedOffset` value is beyond the content length of the
// resource, the [AVAssetResourceLoadingRequest] instance is sent a
// [FinishLoading] message without any prior invocations of [RespondWithData].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingDataRequest/requestedOffset
func (a AVAssetResourceLoadingDataRequest) RequestedOffset() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("requestedOffset"))
	return rv
}

// The position within the resource of the next byte.
//
// # Discussion
//
// When incrementally loading data you should begin loading at this offset,
// returning the data by invoking the [RespondWithData] method. Bytes previous
// to this value have already been provided.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingDataRequest/currentOffset
func (a AVAssetResourceLoadingDataRequest) CurrentOffset() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("currentOffset"))
	return rv
}

// A Boolean value that indicates the entire remaining length of the resource
// from the offest to the end of the resource is being requested.
//
// # Discussion
//
// When this property is true, you should disregard the value of
// requestedLength and incrementally provide as much data, starting from the
// requested offset, as the resource contains. Continue until all available
// data was successfully loaded, the request was cancelled, or an error
// occurs.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingDataRequest/requestsAllDataToEndOfResource
func (a AVAssetResourceLoadingDataRequest) RequestsAllDataToEndOfResource() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("requestsAllDataToEndOfResource"))
	return rv
}

// The range of requested resource data.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/datarequest
func (a AVAssetResourceLoadingDataRequest) DataRequest() IAVAssetResourceLoadingDataRequest {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("dataRequest"))
	return AVAssetResourceLoadingDataRequestFromID(objc.ID(rv))
}
func (a AVAssetResourceLoadingDataRequest) SetDataRequest(value IAVAssetResourceLoadingDataRequest) {
	objc.Send[struct{}](a.ID, objc.Sel("setDataRequest:"), value)
}
