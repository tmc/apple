// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
)

// The class instance for the [VNImageBasedRequest] class.
var (
	_VNImageBasedRequestClass     VNImageBasedRequestClass
	_VNImageBasedRequestClassOnce sync.Once
)

func getVNImageBasedRequestClass() VNImageBasedRequestClass {
	_VNImageBasedRequestClassOnce.Do(func() {
		_VNImageBasedRequestClass = VNImageBasedRequestClass{class: objc.GetClass("VNImageBasedRequest")}
	})
	return _VNImageBasedRequestClass
}

// GetVNImageBasedRequestClass returns the class object for VNImageBasedRequest.
func GetVNImageBasedRequestClass() VNImageBasedRequestClass {
	return getVNImageBasedRequestClass()
}

type VNImageBasedRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNImageBasedRequestClass) Alloc() VNImageBasedRequest {
	rv := objc.Send[VNImageBasedRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// The abstract superclass for image-analysis requests that focus on a
// specific part of an image.
//
// # Overview
// 
// Other Vision request handlers that operate on still images inherit from
// this abstract base class. Don’t use it directly.
//
// # Configuring a Request
//
//   - [VNImageBasedRequest.RegionOfInterest]: The region of the image in which Vision will perform the request.
//   - [VNImageBasedRequest.SetRegionOfInterest]
//
// See: https://developer.apple.com/documentation/Vision/VNImageBasedRequest
type VNImageBasedRequest struct {
	VNRequest
}

// VNImageBasedRequestFromID constructs a [VNImageBasedRequest] from an objc.ID.
//
// The abstract superclass for image-analysis requests that focus on a
// specific part of an image.
func VNImageBasedRequestFromID(id objc.ID) VNImageBasedRequest {
	return VNImageBasedRequest{VNRequest: VNRequestFromID(id)}
}
// NOTE: VNImageBasedRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VNImageBasedRequest] class.
//
// # Configuring a Request
//
//   - [IVNImageBasedRequest.RegionOfInterest]: The region of the image in which Vision will perform the request.
//   - [IVNImageBasedRequest.SetRegionOfInterest]
//
// See: https://developer.apple.com/documentation/Vision/VNImageBasedRequest
type IVNImageBasedRequest interface {
	IVNRequest

	// Topic: Configuring a Request

	// The region of the image in which Vision will perform the request.
	RegionOfInterest() corefoundation.CGRect
	SetRegionOfInterest(value corefoundation.CGRect)
}





// Init initializes the instance.
func (i VNImageBasedRequest) Init() VNImageBasedRequest {
	rv := objc.Send[VNImageBasedRequest](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i VNImageBasedRequest) Autorelease() VNImageBasedRequest {
	rv := objc.Send[VNImageBasedRequest](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNImageBasedRequest creates a new VNImageBasedRequest instance.
func NewVNImageBasedRequest() VNImageBasedRequest {
	class := getVNImageBasedRequestClass()
	rv := objc.Send[VNImageBasedRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewImageBasedRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNImageBasedRequest {
	instance := getVNImageBasedRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNImageBasedRequestFromID(rv)
}


















// The region of the image in which Vision will perform the request.
//
// # Discussion
// 
// The rectangle is normalized to the dimensions of the processed image. Its
// origin is specified relative to the image’s lower-left corner.
// 
// The default value is `{ { 0, 0 }, { 1, 1 } }`.
//
// See: https://developer.apple.com/documentation/Vision/VNImageBasedRequest/regionOfInterest
func (i VNImageBasedRequest) RegionOfInterest() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](i.ID, objc.Sel("regionOfInterest"))
	return corefoundation.CGRect(rv)
}
func (i VNImageBasedRequest) SetRegionOfInterest(value corefoundation.CGRect) {
	objc.Send[struct{}](i.ID, objc.Sel("setRegionOfInterest:"), value)
}
























