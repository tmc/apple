// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VNDetectHorizonRequest] class.
var (
	_VNDetectHorizonRequestClass     VNDetectHorizonRequestClass
	_VNDetectHorizonRequestClassOnce sync.Once
)

func getVNDetectHorizonRequestClass() VNDetectHorizonRequestClass {
	_VNDetectHorizonRequestClassOnce.Do(func() {
		_VNDetectHorizonRequestClass = VNDetectHorizonRequestClass{class: objc.GetClass("VNDetectHorizonRequest")}
	})
	return _VNDetectHorizonRequestClass
}

// GetVNDetectHorizonRequestClass returns the class object for VNDetectHorizonRequest.
func GetVNDetectHorizonRequestClass() VNDetectHorizonRequestClass {
	return getVNDetectHorizonRequestClass()
}

type VNDetectHorizonRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNDetectHorizonRequestClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNDetectHorizonRequestClass) Alloc() VNDetectHorizonRequest {
	rv := objc.Send[VNDetectHorizonRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An image-analysis request that determines the horizon angle in an image.
//
// # Identifying Request Revisions
//
//   - [VNDetectHorizonRequest.VNDetectHorizonRequestRevision1]: A constant for specifying revision 1 of the horizon detection request.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectHorizonRequest
type VNDetectHorizonRequest struct {
	VNImageBasedRequest
}

// VNDetectHorizonRequestFromID constructs a [VNDetectHorizonRequest] from an objc.ID.
//
// An image-analysis request that determines the horizon angle in an image.
func VNDetectHorizonRequestFromID(id objc.ID) VNDetectHorizonRequest {
	return VNDetectHorizonRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}

// NOTE: VNDetectHorizonRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNDetectHorizonRequest] class.
//
// # Identifying Request Revisions
//
//   - [IVNDetectHorizonRequest.VNDetectHorizonRequestRevision1]: A constant for specifying revision 1 of the horizon detection request.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectHorizonRequest
type IVNDetectHorizonRequest interface {
	IVNImageBasedRequest

	// Topic: Identifying Request Revisions

	// A constant for specifying revision 1 of the horizon detection request.
	VNDetectHorizonRequestRevision1() int
}

// Init initializes the instance.
func (d VNDetectHorizonRequest) Init() VNDetectHorizonRequest {
	rv := objc.Send[VNDetectHorizonRequest](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VNDetectHorizonRequest) Autorelease() VNDetectHorizonRequest {
	rv := objc.Send[VNDetectHorizonRequest](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNDetectHorizonRequest creates a new VNDetectHorizonRequest instance.
func NewVNDetectHorizonRequest() VNDetectHorizonRequest {
	class := getVNDetectHorizonRequestClass()
	rv := objc.Send[VNDetectHorizonRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewDetectHorizonRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNDetectHorizonRequest {
	instance := getVNDetectHorizonRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNDetectHorizonRequestFromID(rv)
}

// A constant for specifying revision 1 of the horizon detection request.
//
// See: https://developer.apple.com/documentation/vision/vndetecthorizonrequestrevision1
func (d VNDetectHorizonRequest) VNDetectHorizonRequestRevision1() int {
	rv := objc.Send[int](d.ID, objc.Sel("VNDetectHorizonRequestRevision1"))
	return rv
}
