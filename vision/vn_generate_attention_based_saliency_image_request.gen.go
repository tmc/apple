// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VNGenerateAttentionBasedSaliencyImageRequest] class.
var (
	_VNGenerateAttentionBasedSaliencyImageRequestClass     VNGenerateAttentionBasedSaliencyImageRequestClass
	_VNGenerateAttentionBasedSaliencyImageRequestClassOnce sync.Once
)

func getVNGenerateAttentionBasedSaliencyImageRequestClass() VNGenerateAttentionBasedSaliencyImageRequestClass {
	_VNGenerateAttentionBasedSaliencyImageRequestClassOnce.Do(func() {
		_VNGenerateAttentionBasedSaliencyImageRequestClass = VNGenerateAttentionBasedSaliencyImageRequestClass{class: objc.GetClass("VNGenerateAttentionBasedSaliencyImageRequest")}
	})
	return _VNGenerateAttentionBasedSaliencyImageRequestClass
}

// GetVNGenerateAttentionBasedSaliencyImageRequestClass returns the class object for VNGenerateAttentionBasedSaliencyImageRequest.
func GetVNGenerateAttentionBasedSaliencyImageRequestClass() VNGenerateAttentionBasedSaliencyImageRequestClass {
	return getVNGenerateAttentionBasedSaliencyImageRequestClass()
}

type VNGenerateAttentionBasedSaliencyImageRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNGenerateAttentionBasedSaliencyImageRequestClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNGenerateAttentionBasedSaliencyImageRequestClass) Alloc() VNGenerateAttentionBasedSaliencyImageRequest {
	rv := objc.Send[VNGenerateAttentionBasedSaliencyImageRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that produces a heat map that identifies the parts of an image
// most likely to draw attention.
//
// See: https://developer.apple.com/documentation/Vision/VNGenerateAttentionBasedSaliencyImageRequest
type VNGenerateAttentionBasedSaliencyImageRequest struct {
	VNImageBasedRequest
}

// VNGenerateAttentionBasedSaliencyImageRequestFromID constructs a [VNGenerateAttentionBasedSaliencyImageRequest] from an objc.ID.
//
// An object that produces a heat map that identifies the parts of an image
// most likely to draw attention.
func VNGenerateAttentionBasedSaliencyImageRequestFromID(id objc.ID) VNGenerateAttentionBasedSaliencyImageRequest {
	return VNGenerateAttentionBasedSaliencyImageRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}

// NOTE: VNGenerateAttentionBasedSaliencyImageRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNGenerateAttentionBasedSaliencyImageRequest] class.
//
// See: https://developer.apple.com/documentation/Vision/VNGenerateAttentionBasedSaliencyImageRequest
type IVNGenerateAttentionBasedSaliencyImageRequest interface {
	IVNImageBasedRequest
}

// Init initializes the instance.
func (g VNGenerateAttentionBasedSaliencyImageRequest) Init() VNGenerateAttentionBasedSaliencyImageRequest {
	rv := objc.Send[VNGenerateAttentionBasedSaliencyImageRequest](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g VNGenerateAttentionBasedSaliencyImageRequest) Autorelease() VNGenerateAttentionBasedSaliencyImageRequest {
	rv := objc.Send[VNGenerateAttentionBasedSaliencyImageRequest](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNGenerateAttentionBasedSaliencyImageRequest creates a new VNGenerateAttentionBasedSaliencyImageRequest instance.
func NewVNGenerateAttentionBasedSaliencyImageRequest() VNGenerateAttentionBasedSaliencyImageRequest {
	class := getVNGenerateAttentionBasedSaliencyImageRequestClass()
	rv := objc.Send[VNGenerateAttentionBasedSaliencyImageRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new Vision request with an optional completion handler.
//
// completionHandler: The block to invoke after the request finishes processing.
//
// # Discussion
//
// Vision executes the completion handler on the same queue that it executes
// the request; however, this queue differs from the one where you called
// [VNImageRequestHandler.PerformRequestsError].
//
// See: https://developer.apple.com/documentation/Vision/VNRequest/init(completionHandler:)
func NewGenerateAttentionBasedSaliencyImageRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNGenerateAttentionBasedSaliencyImageRequest {
	instance := getVNGenerateAttentionBasedSaliencyImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNGenerateAttentionBasedSaliencyImageRequestFromID(rv)
}
