// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNGenerateObjectnessBasedSaliencyImageRequest] class.
var (
	_VNGenerateObjectnessBasedSaliencyImageRequestClass     VNGenerateObjectnessBasedSaliencyImageRequestClass
	_VNGenerateObjectnessBasedSaliencyImageRequestClassOnce sync.Once
)

func getVNGenerateObjectnessBasedSaliencyImageRequestClass() VNGenerateObjectnessBasedSaliencyImageRequestClass {
	_VNGenerateObjectnessBasedSaliencyImageRequestClassOnce.Do(func() {
		_VNGenerateObjectnessBasedSaliencyImageRequestClass = VNGenerateObjectnessBasedSaliencyImageRequestClass{class: objc.GetClass("VNGenerateObjectnessBasedSaliencyImageRequest")}
	})
	return _VNGenerateObjectnessBasedSaliencyImageRequestClass
}

// GetVNGenerateObjectnessBasedSaliencyImageRequestClass returns the class object for VNGenerateObjectnessBasedSaliencyImageRequest.
func GetVNGenerateObjectnessBasedSaliencyImageRequestClass() VNGenerateObjectnessBasedSaliencyImageRequestClass {
	return getVNGenerateObjectnessBasedSaliencyImageRequestClass()
}

type VNGenerateObjectnessBasedSaliencyImageRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNGenerateObjectnessBasedSaliencyImageRequestClass) Alloc() VNGenerateObjectnessBasedSaliencyImageRequest {
	rv := objc.Send[VNGenerateObjectnessBasedSaliencyImageRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// A request that generates a heat map that identifies the parts of an image
// most likely to represent objects.
//
// # Overview
// 
// The resulting observation, [VNSaliencyImageObservation], encodes this data
// as a heat map, which you can use to highlight regions of interest.
//
// # Identifying Request Revisions
//
//   - [VNGenerateObjectnessBasedSaliencyImageRequest.VNGenerateObjectnessBasedSaliencyImageRequestRevision1]: A constant for specifying revision 1 of the image saliency request.
//
// See: https://developer.apple.com/documentation/Vision/VNGenerateObjectnessBasedSaliencyImageRequest
type VNGenerateObjectnessBasedSaliencyImageRequest struct {
	VNImageBasedRequest
}

// VNGenerateObjectnessBasedSaliencyImageRequestFromID constructs a [VNGenerateObjectnessBasedSaliencyImageRequest] from an objc.ID.
//
// A request that generates a heat map that identifies the parts of an image
// most likely to represent objects.
func VNGenerateObjectnessBasedSaliencyImageRequestFromID(id objc.ID) VNGenerateObjectnessBasedSaliencyImageRequest {
	return VNGenerateObjectnessBasedSaliencyImageRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}
// NOTE: VNGenerateObjectnessBasedSaliencyImageRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VNGenerateObjectnessBasedSaliencyImageRequest] class.
//
// # Identifying Request Revisions
//
//   - [IVNGenerateObjectnessBasedSaliencyImageRequest.VNGenerateObjectnessBasedSaliencyImageRequestRevision1]: A constant for specifying revision 1 of the image saliency request.
//
// See: https://developer.apple.com/documentation/Vision/VNGenerateObjectnessBasedSaliencyImageRequest
type IVNGenerateObjectnessBasedSaliencyImageRequest interface {
	IVNImageBasedRequest

	// Topic: Identifying Request Revisions

	// A constant for specifying revision 1 of the image saliency request.
	VNGenerateObjectnessBasedSaliencyImageRequestRevision1() int
}





// Init initializes the instance.
func (g VNGenerateObjectnessBasedSaliencyImageRequest) Init() VNGenerateObjectnessBasedSaliencyImageRequest {
	rv := objc.Send[VNGenerateObjectnessBasedSaliencyImageRequest](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g VNGenerateObjectnessBasedSaliencyImageRequest) Autorelease() VNGenerateObjectnessBasedSaliencyImageRequest {
	rv := objc.Send[VNGenerateObjectnessBasedSaliencyImageRequest](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNGenerateObjectnessBasedSaliencyImageRequest creates a new VNGenerateObjectnessBasedSaliencyImageRequest instance.
func NewVNGenerateObjectnessBasedSaliencyImageRequest() VNGenerateObjectnessBasedSaliencyImageRequest {
	class := getVNGenerateObjectnessBasedSaliencyImageRequestClass()
	rv := objc.Send[VNGenerateObjectnessBasedSaliencyImageRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewGenerateObjectnessBasedSaliencyImageRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNGenerateObjectnessBasedSaliencyImageRequest {
	instance := getVNGenerateObjectnessBasedSaliencyImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNGenerateObjectnessBasedSaliencyImageRequestFromID(rv)
}


















// A constant for specifying revision 1 of the image saliency request.
//
// See: https://developer.apple.com/documentation/vision/vngenerateobjectnessbasedsaliencyimagerequestrevision1
func (g VNGenerateObjectnessBasedSaliencyImageRequest) VNGenerateObjectnessBasedSaliencyImageRequestRevision1() int {
	rv := objc.Send[int](g.ID, objc.Sel("VNGenerateObjectnessBasedSaliencyImageRequestRevision1"))
	return rv
}
























