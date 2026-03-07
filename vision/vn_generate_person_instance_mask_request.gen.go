// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNGeneratePersonInstanceMaskRequest] class.
var (
	_VNGeneratePersonInstanceMaskRequestClass     VNGeneratePersonInstanceMaskRequestClass
	_VNGeneratePersonInstanceMaskRequestClassOnce sync.Once
)

func getVNGeneratePersonInstanceMaskRequestClass() VNGeneratePersonInstanceMaskRequestClass {
	_VNGeneratePersonInstanceMaskRequestClassOnce.Do(func() {
		_VNGeneratePersonInstanceMaskRequestClass = VNGeneratePersonInstanceMaskRequestClass{class: objc.GetClass("VNGeneratePersonInstanceMaskRequest")}
	})
	return _VNGeneratePersonInstanceMaskRequestClass
}

// GetVNGeneratePersonInstanceMaskRequestClass returns the class object for VNGeneratePersonInstanceMaskRequest.
func GetVNGeneratePersonInstanceMaskRequestClass() VNGeneratePersonInstanceMaskRequestClass {
	return getVNGeneratePersonInstanceMaskRequestClass()
}

type VNGeneratePersonInstanceMaskRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNGeneratePersonInstanceMaskRequestClass) Alloc() VNGeneratePersonInstanceMaskRequest {
	rv := objc.Send[VNGeneratePersonInstanceMaskRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// An object that produces a mask of individual people it finds in the input
// image.
//
// # Identifying Request Revisions
//
//   - [VNGeneratePersonInstanceMaskRequest.VNGeneratePersonInstanceMaskRequestRevision1]: A constant for specifying revision 1 of the person instance mask request.
//
// See: https://developer.apple.com/documentation/Vision/VNGeneratePersonInstanceMaskRequest
type VNGeneratePersonInstanceMaskRequest struct {
	VNImageBasedRequest
}

// VNGeneratePersonInstanceMaskRequestFromID constructs a [VNGeneratePersonInstanceMaskRequest] from an objc.ID.
//
// An object that produces a mask of individual people it finds in the input
// image.
func VNGeneratePersonInstanceMaskRequestFromID(id objc.ID) VNGeneratePersonInstanceMaskRequest {
	return VNGeneratePersonInstanceMaskRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}
// NOTE: VNGeneratePersonInstanceMaskRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VNGeneratePersonInstanceMaskRequest] class.
//
// # Identifying Request Revisions
//
//   - [IVNGeneratePersonInstanceMaskRequest.VNGeneratePersonInstanceMaskRequestRevision1]: A constant for specifying revision 1 of the person instance mask request.
//
// See: https://developer.apple.com/documentation/Vision/VNGeneratePersonInstanceMaskRequest
type IVNGeneratePersonInstanceMaskRequest interface {
	IVNImageBasedRequest

	// Topic: Identifying Request Revisions

	// A constant for specifying revision 1 of the person instance mask request.
	VNGeneratePersonInstanceMaskRequestRevision1() int
}





// Init initializes the instance.
func (g VNGeneratePersonInstanceMaskRequest) Init() VNGeneratePersonInstanceMaskRequest {
	rv := objc.Send[VNGeneratePersonInstanceMaskRequest](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g VNGeneratePersonInstanceMaskRequest) Autorelease() VNGeneratePersonInstanceMaskRequest {
	rv := objc.Send[VNGeneratePersonInstanceMaskRequest](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNGeneratePersonInstanceMaskRequest creates a new VNGeneratePersonInstanceMaskRequest instance.
func NewVNGeneratePersonInstanceMaskRequest() VNGeneratePersonInstanceMaskRequest {
	class := getVNGeneratePersonInstanceMaskRequestClass()
	rv := objc.Send[VNGeneratePersonInstanceMaskRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewGeneratePersonInstanceMaskRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNGeneratePersonInstanceMaskRequest {
	instance := getVNGeneratePersonInstanceMaskRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNGeneratePersonInstanceMaskRequestFromID(rv)
}


















// A constant for specifying revision 1 of the person instance mask request.
//
// See: https://developer.apple.com/documentation/vision/vngeneratepersoninstancemaskrequestrevision1
func (g VNGeneratePersonInstanceMaskRequest) VNGeneratePersonInstanceMaskRequestRevision1() int {
	rv := objc.Send[int](g.ID, objc.Sel("VNGeneratePersonInstanceMaskRequestRevision1"))
	return rv
}
























