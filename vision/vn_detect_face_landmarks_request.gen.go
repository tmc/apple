// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNDetectFaceLandmarksRequest] class.
var (
	_VNDetectFaceLandmarksRequestClass     VNDetectFaceLandmarksRequestClass
	_VNDetectFaceLandmarksRequestClassOnce sync.Once
)

func getVNDetectFaceLandmarksRequestClass() VNDetectFaceLandmarksRequestClass {
	_VNDetectFaceLandmarksRequestClassOnce.Do(func() {
		_VNDetectFaceLandmarksRequestClass = VNDetectFaceLandmarksRequestClass{class: objc.GetClass("VNDetectFaceLandmarksRequest")}
	})
	return _VNDetectFaceLandmarksRequestClass
}

// GetVNDetectFaceLandmarksRequestClass returns the class object for VNDetectFaceLandmarksRequest.
func GetVNDetectFaceLandmarksRequestClass() VNDetectFaceLandmarksRequestClass {
	return getVNDetectFaceLandmarksRequestClass()
}

type VNDetectFaceLandmarksRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNDetectFaceLandmarksRequestClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNDetectFaceLandmarksRequestClass) Alloc() VNDetectFaceLandmarksRequest {
	rv := objc.Send[VNDetectFaceLandmarksRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An image-analysis request that finds facial features like eyes and mouth in
// an image.
//
// # Overview
//
// By default, a face landmarks request first locates all faces in the input
// image, then analyzes each to detect facial features.
//
// If you’ve already located all the faces in an image, or want to detect
// landmarks in only a subset of the faces in the image, set the
// [VNDetectFaceLandmarksRequest.InputFaceObservations] property to an array of [VNFaceObservation] objects
// representing the faces you want to analyze. You can either use face
// observations output by a [VNDetectFaceRectanglesRequest] or manually create
// [VNFaceObservation] instances with the bounding boxes of the faces you want
// to analyze.
//
// # Locating Face Landmarks
//
//   - [VNDetectFaceLandmarksRequest.Constellation]: A variable that describes how a face landmarks request orders or enumerates the resulting features.
//   - [VNDetectFaceLandmarksRequest.SetConstellation]
//
// See: https://developer.apple.com/documentation/Vision/VNDetectFaceLandmarksRequest
type VNDetectFaceLandmarksRequest struct {
	VNImageBasedRequest
}

// VNDetectFaceLandmarksRequestFromID constructs a [VNDetectFaceLandmarksRequest] from an objc.ID.
//
// An image-analysis request that finds facial features like eyes and mouth in
// an image.
func VNDetectFaceLandmarksRequestFromID(id objc.ID) VNDetectFaceLandmarksRequest {
	return VNDetectFaceLandmarksRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}

// NOTE: VNDetectFaceLandmarksRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNDetectFaceLandmarksRequest] class.
//
// # Locating Face Landmarks
//
//   - [IVNDetectFaceLandmarksRequest.Constellation]: A variable that describes how a face landmarks request orders or enumerates the resulting features.
//   - [IVNDetectFaceLandmarksRequest.SetConstellation]
//
// See: https://developer.apple.com/documentation/Vision/VNDetectFaceLandmarksRequest
type IVNDetectFaceLandmarksRequest interface {
	IVNImageBasedRequest

	// Topic: Locating Face Landmarks

	// A variable that describes how a face landmarks request orders or enumerates the resulting features.
	Constellation() VNRequestFaceLandmarksConstellation
	SetConstellation(value VNRequestFaceLandmarksConstellation)
}

// Init initializes the instance.
func (d VNDetectFaceLandmarksRequest) Init() VNDetectFaceLandmarksRequest {
	rv := objc.Send[VNDetectFaceLandmarksRequest](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VNDetectFaceLandmarksRequest) Autorelease() VNDetectFaceLandmarksRequest {
	rv := objc.Send[VNDetectFaceLandmarksRequest](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNDetectFaceLandmarksRequest creates a new VNDetectFaceLandmarksRequest instance.
func NewVNDetectFaceLandmarksRequest() VNDetectFaceLandmarksRequest {
	class := getVNDetectFaceLandmarksRequestClass()
	rv := objc.Send[VNDetectFaceLandmarksRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewDetectFaceLandmarksRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNDetectFaceLandmarksRequest {
	instance := getVNDetectFaceLandmarksRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNDetectFaceLandmarksRequestFromID(rv)
}

// An array of [VNFaceObservation] objects to process as part of the request.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceObservationAccepting/inputFaceObservations
func (d VNDetectFaceLandmarksRequest) InputFaceObservations() []VNFaceObservation {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("inputFaceObservations"))
	return objc.ConvertSlice(rv, func(id objc.ID) VNFaceObservation {
		return VNFaceObservationFromID(id)
	})
}

// Returns a Boolean value that indicates whether a revision supports a
// constellation.
//
// requestRevision: The revision of the request.
//
// constellation: The contellation for which to determine support.
//
// # Return Value
//
// true if the revision supports the constellation, otherwise false.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectFaceLandmarksRequest/revision(_:supportsConstellation:)
func (_VNDetectFaceLandmarksRequestClass VNDetectFaceLandmarksRequestClass) RevisionSupportsConstellation(requestRevision uint, constellation VNRequestFaceLandmarksConstellation) bool {
	rv := objc.Send[bool](objc.ID(_VNDetectFaceLandmarksRequestClass.class), objc.Sel("revision:supportsConstellation:"), requestRevision, constellation)
	return rv
}

// A variable that describes how a face landmarks request orders or enumerates
// the resulting features.
//
// # Discussion
//
// Set this variable to one of the supported constellation types detailed in
// [VNRequestFaceLandmarksConstellation]. The default value is
// [VNRequestFaceLandmarksConstellationNotDefined].
//
// See: https://developer.apple.com/documentation/Vision/VNDetectFaceLandmarksRequest/constellation
//
// [VNRequestFaceLandmarksConstellation]: https://developer.apple.com/documentation/Vision/VNRequestFaceLandmarksConstellation
func (d VNDetectFaceLandmarksRequest) Constellation() VNRequestFaceLandmarksConstellation {
	rv := objc.Send[VNRequestFaceLandmarksConstellation](d.ID, objc.Sel("constellation"))
	return VNRequestFaceLandmarksConstellation(rv)
}
func (d VNDetectFaceLandmarksRequest) SetConstellation(value VNRequestFaceLandmarksConstellation) {
	objc.Send[struct{}](d.ID, objc.Sel("setConstellation:"), value)
}

// Protocol methods for VNFaceObservationAccepting

// An array of [VNFaceObservation] objects to process as part of the request.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceObservationAccepting/inputFaceObservations
func (o VNDetectFaceLandmarksRequest) SetInputFaceObservations(value []VNFaceObservation) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputFaceObservations:"), objectivec.IObjectSliceToNSArray(value))
}
