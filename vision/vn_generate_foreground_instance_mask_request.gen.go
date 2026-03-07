// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNGenerateForegroundInstanceMaskRequest] class.
var (
	_VNGenerateForegroundInstanceMaskRequestClass     VNGenerateForegroundInstanceMaskRequestClass
	_VNGenerateForegroundInstanceMaskRequestClassOnce sync.Once
)

func getVNGenerateForegroundInstanceMaskRequestClass() VNGenerateForegroundInstanceMaskRequestClass {
	_VNGenerateForegroundInstanceMaskRequestClassOnce.Do(func() {
		_VNGenerateForegroundInstanceMaskRequestClass = VNGenerateForegroundInstanceMaskRequestClass{class: objc.GetClass("VNGenerateForegroundInstanceMaskRequest")}
	})
	return _VNGenerateForegroundInstanceMaskRequestClass
}

// GetVNGenerateForegroundInstanceMaskRequestClass returns the class object for VNGenerateForegroundInstanceMaskRequest.
func GetVNGenerateForegroundInstanceMaskRequestClass() VNGenerateForegroundInstanceMaskRequestClass {
	return getVNGenerateForegroundInstanceMaskRequestClass()
}

type VNGenerateForegroundInstanceMaskRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNGenerateForegroundInstanceMaskRequestClass) Alloc() VNGenerateForegroundInstanceMaskRequest {
	rv := objc.Send[VNGenerateForegroundInstanceMaskRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// A request that generates an instance mask of noticable objects to separate
// from the background.
//
// See: https://developer.apple.com/documentation/Vision/VNGenerateForegroundInstanceMaskRequest
type VNGenerateForegroundInstanceMaskRequest struct {
	VNImageBasedRequest
}

// VNGenerateForegroundInstanceMaskRequestFromID constructs a [VNGenerateForegroundInstanceMaskRequest] from an objc.ID.
//
// A request that generates an instance mask of noticable objects to separate
// from the background.
func VNGenerateForegroundInstanceMaskRequestFromID(id objc.ID) VNGenerateForegroundInstanceMaskRequest {
	return VNGenerateForegroundInstanceMaskRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}
// NOTE: VNGenerateForegroundInstanceMaskRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VNGenerateForegroundInstanceMaskRequest] class.
//
// See: https://developer.apple.com/documentation/Vision/VNGenerateForegroundInstanceMaskRequest
type IVNGenerateForegroundInstanceMaskRequest interface {
	IVNImageBasedRequest

	// A constant for specifying the first revision of the foreground instance mask request.
	VNGenerateForegroundInstanceMaskRequestRevision1() int
}





// Init initializes the instance.
func (g VNGenerateForegroundInstanceMaskRequest) Init() VNGenerateForegroundInstanceMaskRequest {
	rv := objc.Send[VNGenerateForegroundInstanceMaskRequest](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g VNGenerateForegroundInstanceMaskRequest) Autorelease() VNGenerateForegroundInstanceMaskRequest {
	rv := objc.Send[VNGenerateForegroundInstanceMaskRequest](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNGenerateForegroundInstanceMaskRequest creates a new VNGenerateForegroundInstanceMaskRequest instance.
func NewVNGenerateForegroundInstanceMaskRequest() VNGenerateForegroundInstanceMaskRequest {
	class := getVNGenerateForegroundInstanceMaskRequestClass()
	rv := objc.Send[VNGenerateForegroundInstanceMaskRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewGenerateForegroundInstanceMaskRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNGenerateForegroundInstanceMaskRequest {
	instance := getVNGenerateForegroundInstanceMaskRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNGenerateForegroundInstanceMaskRequestFromID(rv)
}


















// A constant for specifying the first revision of the foreground instance
// mask request.
//
// See: https://developer.apple.com/documentation/vision/vngenerateforegroundinstancemaskrequestrevision1
func (g VNGenerateForegroundInstanceMaskRequest) VNGenerateForegroundInstanceMaskRequestRevision1() int {
	rv := objc.Send[int](g.ID, objc.Sel("VNGenerateForegroundInstanceMaskRequestRevision1"))
	return rv
}
























