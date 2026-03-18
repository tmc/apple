// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNDetectHumanRectanglesRequest] class.
var (
	_VNDetectHumanRectanglesRequestClass     VNDetectHumanRectanglesRequestClass
	_VNDetectHumanRectanglesRequestClassOnce sync.Once
)

func getVNDetectHumanRectanglesRequestClass() VNDetectHumanRectanglesRequestClass {
	_VNDetectHumanRectanglesRequestClassOnce.Do(func() {
		_VNDetectHumanRectanglesRequestClass = VNDetectHumanRectanglesRequestClass{class: objc.GetClass("VNDetectHumanRectanglesRequest")}
	})
	return _VNDetectHumanRectanglesRequestClass
}

// GetVNDetectHumanRectanglesRequestClass returns the class object for VNDetectHumanRectanglesRequest.
func GetVNDetectHumanRectanglesRequestClass() VNDetectHumanRectanglesRequestClass {
	return getVNDetectHumanRectanglesRequestClass()
}

type VNDetectHumanRectanglesRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNDetectHumanRectanglesRequestClass) Alloc() VNDetectHumanRectanglesRequest {
	rv := objc.Send[VNDetectHumanRectanglesRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A request that finds rectangular regions that contain people in an image.
//
// # Configuring the Request
//
//   - [VNDetectHumanRectanglesRequest.UpperBodyOnly]: A Boolean value that indicates whether the request requires detecting a full body or upper body only to produce a result.
//   - [VNDetectHumanRectanglesRequest.SetUpperBodyOnly]
//
// # Identifying Request Revisions
//
//   - [VNDetectHumanRectanglesRequest.VNDetectHumanRectanglesRequestRevision2]: A constant for specifying revision 2 of the human rectangles detection request.
//   - [VNDetectHumanRectanglesRequest.VNDetectHumanRectanglesRequestRevision1]: A constant for specifying revision 1 of the human rectangles detection request.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectHumanRectanglesRequest
type VNDetectHumanRectanglesRequest struct {
	VNImageBasedRequest
}

// VNDetectHumanRectanglesRequestFromID constructs a [VNDetectHumanRectanglesRequest] from an objc.ID.
//
// A request that finds rectangular regions that contain people in an image.
func VNDetectHumanRectanglesRequestFromID(id objc.ID) VNDetectHumanRectanglesRequest {
	return VNDetectHumanRectanglesRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}
// NOTE: VNDetectHumanRectanglesRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNDetectHumanRectanglesRequest] class.
//
// # Configuring the Request
//
//   - [IVNDetectHumanRectanglesRequest.UpperBodyOnly]: A Boolean value that indicates whether the request requires detecting a full body or upper body only to produce a result.
//   - [IVNDetectHumanRectanglesRequest.SetUpperBodyOnly]
//
// # Identifying Request Revisions
//
//   - [IVNDetectHumanRectanglesRequest.VNDetectHumanRectanglesRequestRevision2]: A constant for specifying revision 2 of the human rectangles detection request.
//   - [IVNDetectHumanRectanglesRequest.VNDetectHumanRectanglesRequestRevision1]: A constant for specifying revision 1 of the human rectangles detection request.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectHumanRectanglesRequest
type IVNDetectHumanRectanglesRequest interface {
	IVNImageBasedRequest

	// Topic: Configuring the Request

	// A Boolean value that indicates whether the request requires detecting a full body or upper body only to produce a result.
	UpperBodyOnly() bool
	SetUpperBodyOnly(value bool)

	// Topic: Identifying Request Revisions

	// A constant for specifying revision 2 of the human rectangles detection request.
	VNDetectHumanRectanglesRequestRevision2() int
	// A constant for specifying revision 1 of the human rectangles detection request.
	VNDetectHumanRectanglesRequestRevision1() int
}

// Init initializes the instance.
func (d VNDetectHumanRectanglesRequest) Init() VNDetectHumanRectanglesRequest {
	rv := objc.Send[VNDetectHumanRectanglesRequest](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VNDetectHumanRectanglesRequest) Autorelease() VNDetectHumanRectanglesRequest {
	rv := objc.Send[VNDetectHumanRectanglesRequest](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNDetectHumanRectanglesRequest creates a new VNDetectHumanRectanglesRequest instance.
func NewVNDetectHumanRectanglesRequest() VNDetectHumanRectanglesRequest {
	class := getVNDetectHumanRectanglesRequestClass()
	rv := objc.Send[VNDetectHumanRectanglesRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewDetectHumanRectanglesRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNDetectHumanRectanglesRequest {
	instance := getVNDetectHumanRectanglesRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNDetectHumanRectanglesRequestFromID(rv)
}

// A Boolean value that indicates whether the request requires detecting a
// full body or upper body only to produce a result.
//
// # Discussion
// 
// The default value of [true] indicates that the request requires detecting a
// person’s upper body only to find the bound box around it.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Vision/VNDetectHumanRectanglesRequest/upperBodyOnly
func (d VNDetectHumanRectanglesRequest) UpperBodyOnly() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("upperBodyOnly"))
	return rv
}
func (d VNDetectHumanRectanglesRequest) SetUpperBodyOnly(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setUpperBodyOnly:"), value)
}

// A constant for specifying revision 2 of the human rectangles detection
// request.
//
// See: https://developer.apple.com/documentation/vision/vndetecthumanrectanglesrequestrevision2
func (d VNDetectHumanRectanglesRequest) VNDetectHumanRectanglesRequestRevision2() int {
	rv := objc.Send[int](d.ID, objc.Sel("VNDetectHumanRectanglesRequestRevision2"))
	return rv
}

// A constant for specifying revision 1 of the human rectangles detection
// request.
//
// See: https://developer.apple.com/documentation/vision/vndetecthumanrectanglesrequestrevision1
func (d VNDetectHumanRectanglesRequest) VNDetectHumanRectanglesRequestRevision1() int {
	rv := objc.Send[int](d.ID, objc.Sel("VNDetectHumanRectanglesRequestRevision1"))
	return rv
}

