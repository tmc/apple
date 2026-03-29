// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVAssetResourceRenewalRequest] class.
var (
	_AVAssetResourceRenewalRequestClass     AVAssetResourceRenewalRequestClass
	_AVAssetResourceRenewalRequestClassOnce sync.Once
)

func getAVAssetResourceRenewalRequestClass() AVAssetResourceRenewalRequestClass {
	_AVAssetResourceRenewalRequestClassOnce.Do(func() {
		_AVAssetResourceRenewalRequestClass = AVAssetResourceRenewalRequestClass{class: objc.GetClass("AVAssetResourceRenewalRequest")}
	})
	return _AVAssetResourceRenewalRequestClass
}

// GetAVAssetResourceRenewalRequestClass returns the class object for AVAssetResourceRenewalRequest.
func GetAVAssetResourceRenewalRequestClass() AVAssetResourceRenewalRequestClass {
	return getAVAssetResourceRenewalRequestClass()
}

type AVAssetResourceRenewalRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetResourceRenewalRequestClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetResourceRenewalRequestClass) Alloc() AVAssetResourceRenewalRequest {
	rv := objc.Send[AVAssetResourceRenewalRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that encapsulates information about a resource request from a
// resource loader to renew a previously issued request.
//
// # Overview
// 
// When an [AVURLAsset] needs to renew a resource, because the [AVAssetResourceRenewalRequest.RenewalDate]
// has been set on a previous loading request, it asks its
// [AVAssetResourceLoader] object to assist. The resource loader encapsulates
// the request information by creating an instance of this object, which it
// then hands to its delegate for processing. The delegate uses the
// information in this object to perform the request and report on the success
// or failure of the operation.
// 
// The [AVAssetResourceRenewalRequest] class is a subclass of
// [AVAssetResourceLoadingRequest].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceRenewalRequest
type AVAssetResourceRenewalRequest struct {
	AVAssetResourceLoadingRequest
}

// AVAssetResourceRenewalRequestFromID constructs a [AVAssetResourceRenewalRequest] from an objc.ID.
//
// An object that encapsulates information about a resource request from a
// resource loader to renew a previously issued request.
func AVAssetResourceRenewalRequestFromID(id objc.ID) AVAssetResourceRenewalRequest {
	return AVAssetResourceRenewalRequest{AVAssetResourceLoadingRequest: AVAssetResourceLoadingRequestFromID(id)}
}
// NOTE: AVAssetResourceRenewalRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetResourceRenewalRequest] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceRenewalRequest
type IAVAssetResourceRenewalRequest interface {
	IAVAssetResourceLoadingRequest

	// The date at which a new resource loading request will be issued for resources that expire, if the media system still requires it.
	RenewalDate() foundation.INSDate
	SetRenewalDate(value foundation.INSDate)
}

// Init initializes the instance.
func (a AVAssetResourceRenewalRequest) Init() AVAssetResourceRenewalRequest {
	rv := objc.Send[AVAssetResourceRenewalRequest](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetResourceRenewalRequest) Autorelease() AVAssetResourceRenewalRequest {
	rv := objc.Send[AVAssetResourceRenewalRequest](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetResourceRenewalRequest creates a new AVAssetResourceRenewalRequest instance.
func NewAVAssetResourceRenewalRequest() AVAssetResourceRenewalRequest {
	class := getAVAssetResourceRenewalRequestClass()
	rv := objc.Send[AVAssetResourceRenewalRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The date at which a new resource loading request will be issued for
// resources that expire, if the media system still requires it.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingcontentinformationrequest/renewaldate
func (a AVAssetResourceRenewalRequest) RenewalDate() foundation.INSDate {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("renewalDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (a AVAssetResourceRenewalRequest) SetRenewalDate(value foundation.INSDate) {
	objc.Send[struct{}](a.ID, objc.Sel("setRenewalDate:"), value)
}

