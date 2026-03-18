// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNCalculateImageAestheticsScoresRequest] class.
var (
	_VNCalculateImageAestheticsScoresRequestClass     VNCalculateImageAestheticsScoresRequestClass
	_VNCalculateImageAestheticsScoresRequestClassOnce sync.Once
)

func getVNCalculateImageAestheticsScoresRequestClass() VNCalculateImageAestheticsScoresRequestClass {
	_VNCalculateImageAestheticsScoresRequestClassOnce.Do(func() {
		_VNCalculateImageAestheticsScoresRequestClass = VNCalculateImageAestheticsScoresRequestClass{class: objc.GetClass("VNCalculateImageAestheticsScoresRequest")}
	})
	return _VNCalculateImageAestheticsScoresRequestClass
}

// GetVNCalculateImageAestheticsScoresRequestClass returns the class object for VNCalculateImageAestheticsScoresRequest.
func GetVNCalculateImageAestheticsScoresRequestClass() VNCalculateImageAestheticsScoresRequestClass {
	return getVNCalculateImageAestheticsScoresRequestClass()
}

type VNCalculateImageAestheticsScoresRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNCalculateImageAestheticsScoresRequestClass) Alloc() VNCalculateImageAestheticsScoresRequest {
	rv := objc.Send[VNCalculateImageAestheticsScoresRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that analyzes an image for aesthetically pleasing attributes.
//
// See: https://developer.apple.com/documentation/Vision/VNCalculateImageAestheticsScoresRequest
type VNCalculateImageAestheticsScoresRequest struct {
	VNImageBasedRequest
}

// VNCalculateImageAestheticsScoresRequestFromID constructs a [VNCalculateImageAestheticsScoresRequest] from an objc.ID.
//
// An object that analyzes an image for aesthetically pleasing attributes.
func VNCalculateImageAestheticsScoresRequestFromID(id objc.ID) VNCalculateImageAestheticsScoresRequest {
	return VNCalculateImageAestheticsScoresRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}
// NOTE: VNCalculateImageAestheticsScoresRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNCalculateImageAestheticsScoresRequest] class.
//
// See: https://developer.apple.com/documentation/Vision/VNCalculateImageAestheticsScoresRequest
type IVNCalculateImageAestheticsScoresRequest interface {
	IVNImageBasedRequest
}

// Init initializes the instance.
func (c VNCalculateImageAestheticsScoresRequest) Init() VNCalculateImageAestheticsScoresRequest {
	rv := objc.Send[VNCalculateImageAestheticsScoresRequest](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c VNCalculateImageAestheticsScoresRequest) Autorelease() VNCalculateImageAestheticsScoresRequest {
	rv := objc.Send[VNCalculateImageAestheticsScoresRequest](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNCalculateImageAestheticsScoresRequest creates a new VNCalculateImageAestheticsScoresRequest instance.
func NewVNCalculateImageAestheticsScoresRequest() VNCalculateImageAestheticsScoresRequest {
	class := getVNCalculateImageAestheticsScoresRequestClass()
	rv := objc.Send[VNCalculateImageAestheticsScoresRequest](objc.ID(class.class), objc.Sel("new"))
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
// [PerformRequestsError].
//
// See: https://developer.apple.com/documentation/Vision/VNRequest/init(completionHandler:)
func NewCalculateImageAestheticsScoresRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNCalculateImageAestheticsScoresRequest {
	instance := getVNCalculateImageAestheticsScoresRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNCalculateImageAestheticsScoresRequestFromID(rv)
}

