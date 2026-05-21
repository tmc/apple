// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VNDetectFaceRectanglesRequest] class.
var (
	_VNDetectFaceRectanglesRequestClass     VNDetectFaceRectanglesRequestClass
	_VNDetectFaceRectanglesRequestClassOnce sync.Once
)

func getVNDetectFaceRectanglesRequestClass() VNDetectFaceRectanglesRequestClass {
	_VNDetectFaceRectanglesRequestClassOnce.Do(func() {
		_VNDetectFaceRectanglesRequestClass = VNDetectFaceRectanglesRequestClass{class: objc.GetClass("VNDetectFaceRectanglesRequest")}
	})
	return _VNDetectFaceRectanglesRequestClass
}

// GetVNDetectFaceRectanglesRequestClass returns the class object for VNDetectFaceRectanglesRequest.
func GetVNDetectFaceRectanglesRequestClass() VNDetectFaceRectanglesRequestClass {
	return getVNDetectFaceRectanglesRequestClass()
}

type VNDetectFaceRectanglesRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNDetectFaceRectanglesRequestClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNDetectFaceRectanglesRequestClass) Alloc() VNDetectFaceRectanglesRequest {
	rv := objc.Send[VNDetectFaceRectanglesRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A request that finds faces within an image.
//
// # Overview
//
// This request returns faces as rectangular bounding boxes with origin and
// size.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectFaceRectanglesRequest
type VNDetectFaceRectanglesRequest struct {
	VNImageBasedRequest
}

// VNDetectFaceRectanglesRequestFromID constructs a [VNDetectFaceRectanglesRequest] from an objc.ID.
//
// A request that finds faces within an image.
func VNDetectFaceRectanglesRequestFromID(id objc.ID) VNDetectFaceRectanglesRequest {
	return VNDetectFaceRectanglesRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}

// NOTE: VNDetectFaceRectanglesRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNDetectFaceRectanglesRequest] class.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectFaceRectanglesRequest
type IVNDetectFaceRectanglesRequest interface {
	IVNImageBasedRequest
}

// Init initializes the instance.
func (d VNDetectFaceRectanglesRequest) Init() VNDetectFaceRectanglesRequest {
	rv := objc.Send[VNDetectFaceRectanglesRequest](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VNDetectFaceRectanglesRequest) Autorelease() VNDetectFaceRectanglesRequest {
	rv := objc.Send[VNDetectFaceRectanglesRequest](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNDetectFaceRectanglesRequest creates a new VNDetectFaceRectanglesRequest instance.
func NewVNDetectFaceRectanglesRequest() VNDetectFaceRectanglesRequest {
	class := getVNDetectFaceRectanglesRequestClass()
	rv := objc.Send[VNDetectFaceRectanglesRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewDetectFaceRectanglesRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNDetectFaceRectanglesRequest {
	instance := getVNDetectFaceRectanglesRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNDetectFaceRectanglesRequestFromID(rv)
}
