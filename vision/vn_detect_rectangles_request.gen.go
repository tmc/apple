// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNDetectRectanglesRequest] class.
var (
	_VNDetectRectanglesRequestClass     VNDetectRectanglesRequestClass
	_VNDetectRectanglesRequestClassOnce sync.Once
)

func getVNDetectRectanglesRequestClass() VNDetectRectanglesRequestClass {
	_VNDetectRectanglesRequestClassOnce.Do(func() {
		_VNDetectRectanglesRequestClass = VNDetectRectanglesRequestClass{class: objc.GetClass("VNDetectRectanglesRequest")}
	})
	return _VNDetectRectanglesRequestClass
}

// GetVNDetectRectanglesRequestClass returns the class object for VNDetectRectanglesRequest.
func GetVNDetectRectanglesRequestClass() VNDetectRectanglesRequestClass {
	return getVNDetectRectanglesRequestClass()
}

type VNDetectRectanglesRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNDetectRectanglesRequestClass) Alloc() VNDetectRectanglesRequest {
	rv := objc.Send[VNDetectRectanglesRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An image-analysis request that finds projected rectangular regions in an
// image.
//
// # Overview
// 
// A rectangle detection request locates regions of an image with rectangular
// shape, like credit cards, business cards, documents, and signs. The request
// returns its observations in the form of [VNRectangleObservation] objects,
// which contain normalized coordinates of bounding boxes containing the
// rectangle.
// 
// Use this type of request to find the bounding boxes of rectangles in an
// image. Vision returns observations for rectangles found in all orientations
// and sizes, along with a confidence level to indicate how likely it’s that
// the observation contains an actual rectangle.
// 
// To further configure or restrict the types of rectangles found, set
// properties on the request specifying a range of aspect ratios, sizes, and
// quadrature tolerance.
//
// # Configuring Detection
//
//   - [VNDetectRectanglesRequest.MinimumAspectRatio]: A `float` specifying the minimum aspect ratio of the rectangle to detect, defined as the shorter dimension over the longer dimension.
//   - [VNDetectRectanglesRequest.SetMinimumAspectRatio]
//   - [VNDetectRectanglesRequest.MaximumAspectRatio]: A `float` specifying the maximum aspect ratio of the rectangle to detect, defined as the shorter dimension over the longer dimension.
//   - [VNDetectRectanglesRequest.SetMaximumAspectRatio]
//   - [VNDetectRectanglesRequest.QuadratureTolerance]: A float specifying the number of degrees a rectangle corner angle can deviate from 90°.
//   - [VNDetectRectanglesRequest.SetQuadratureTolerance]
//   - [VNDetectRectanglesRequest.MinimumSize]: The minimum size of a rectangle to detect, as a proportion of the smallest dimension.
//   - [VNDetectRectanglesRequest.SetMinimumSize]
//   - [VNDetectRectanglesRequest.MinimumConfidence]: A value specifying the minimum acceptable confidence level.
//   - [VNDetectRectanglesRequest.SetMinimumConfidence]
//   - [VNDetectRectanglesRequest.MaximumObservations]: An integer specifying the maximum number of rectangles Vision returns.
//   - [VNDetectRectanglesRequest.SetMaximumObservations]
//
// # Identifying Request Revisions
//
//   - [VNDetectRectanglesRequest.VNDetectRectanglesRequestRevision1]: A constant for specifying revision 1 of the rectangle detection request.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest
type VNDetectRectanglesRequest struct {
	VNImageBasedRequest
}

// VNDetectRectanglesRequestFromID constructs a [VNDetectRectanglesRequest] from an objc.ID.
//
// An image-analysis request that finds projected rectangular regions in an
// image.
func VNDetectRectanglesRequestFromID(id objc.ID) VNDetectRectanglesRequest {
	return VNDetectRectanglesRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}
// NOTE: VNDetectRectanglesRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNDetectRectanglesRequest] class.
//
// # Configuring Detection
//
//   - [IVNDetectRectanglesRequest.MinimumAspectRatio]: A `float` specifying the minimum aspect ratio of the rectangle to detect, defined as the shorter dimension over the longer dimension.
//   - [IVNDetectRectanglesRequest.SetMinimumAspectRatio]
//   - [IVNDetectRectanglesRequest.MaximumAspectRatio]: A `float` specifying the maximum aspect ratio of the rectangle to detect, defined as the shorter dimension over the longer dimension.
//   - [IVNDetectRectanglesRequest.SetMaximumAspectRatio]
//   - [IVNDetectRectanglesRequest.QuadratureTolerance]: A float specifying the number of degrees a rectangle corner angle can deviate from 90°.
//   - [IVNDetectRectanglesRequest.SetQuadratureTolerance]
//   - [IVNDetectRectanglesRequest.MinimumSize]: The minimum size of a rectangle to detect, as a proportion of the smallest dimension.
//   - [IVNDetectRectanglesRequest.SetMinimumSize]
//   - [IVNDetectRectanglesRequest.MinimumConfidence]: A value specifying the minimum acceptable confidence level.
//   - [IVNDetectRectanglesRequest.SetMinimumConfidence]
//   - [IVNDetectRectanglesRequest.MaximumObservations]: An integer specifying the maximum number of rectangles Vision returns.
//   - [IVNDetectRectanglesRequest.SetMaximumObservations]
//
// # Identifying Request Revisions
//
//   - [IVNDetectRectanglesRequest.VNDetectRectanglesRequestRevision1]: A constant for specifying revision 1 of the rectangle detection request.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest
type IVNDetectRectanglesRequest interface {
	IVNImageBasedRequest

	// Topic: Configuring Detection

	// A `float` specifying the minimum aspect ratio of the rectangle to detect, defined as the shorter dimension over the longer dimension.
	MinimumAspectRatio() VNAspectRatio
	SetMinimumAspectRatio(value VNAspectRatio)
	// A `float` specifying the maximum aspect ratio of the rectangle to detect, defined as the shorter dimension over the longer dimension.
	MaximumAspectRatio() VNAspectRatio
	SetMaximumAspectRatio(value VNAspectRatio)
	// A float specifying the number of degrees a rectangle corner angle can deviate from 90°.
	QuadratureTolerance() VNDegrees
	SetQuadratureTolerance(value VNDegrees)
	// The minimum size of a rectangle to detect, as a proportion of the smallest dimension.
	MinimumSize() float32
	SetMinimumSize(value float32)
	// A value specifying the minimum acceptable confidence level.
	MinimumConfidence() VNConfidence
	SetMinimumConfidence(value VNConfidence)
	// An integer specifying the maximum number of rectangles Vision returns.
	MaximumObservations() uint
	SetMaximumObservations(value uint)

	// Topic: Identifying Request Revisions

	// A constant for specifying revision 1 of the rectangle detection request.
	VNDetectRectanglesRequestRevision1() int
}

// Init initializes the instance.
func (d VNDetectRectanglesRequest) Init() VNDetectRectanglesRequest {
	rv := objc.Send[VNDetectRectanglesRequest](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VNDetectRectanglesRequest) Autorelease() VNDetectRectanglesRequest {
	rv := objc.Send[VNDetectRectanglesRequest](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNDetectRectanglesRequest creates a new VNDetectRectanglesRequest instance.
func NewVNDetectRectanglesRequest() VNDetectRectanglesRequest {
	class := getVNDetectRectanglesRequestClass()
	rv := objc.Send[VNDetectRectanglesRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewDetectRectanglesRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNDetectRectanglesRequest {
	instance := getVNDetectRectanglesRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNDetectRectanglesRequestFromID(rv)
}

// A `float` specifying the minimum aspect ratio of the rectangle to detect,
// defined as the shorter dimension over the longer dimension.
//
// # Discussion
// 
// The value should range from `0.0` to `1.0`, inclusive. The default value is
// `0.5`.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/minimumAspectRatio
func (d VNDetectRectanglesRequest) MinimumAspectRatio() VNAspectRatio {
	rv := objc.Send[VNAspectRatio](d.ID, objc.Sel("minimumAspectRatio"))
	return VNAspectRatio(rv)
}
func (d VNDetectRectanglesRequest) SetMinimumAspectRatio(value VNAspectRatio) {
	objc.Send[struct{}](d.ID, objc.Sel("setMinimumAspectRatio:"), value)
}

// A `float` specifying the maximum aspect ratio of the rectangle to detect,
// defined as the shorter dimension over the longer dimension.
//
// # Discussion
// 
// The value should range from `0.0` to `1.0`, inclusive. The default value is
// `0.5`.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/maximumAspectRatio
func (d VNDetectRectanglesRequest) MaximumAspectRatio() VNAspectRatio {
	rv := objc.Send[VNAspectRatio](d.ID, objc.Sel("maximumAspectRatio"))
	return VNAspectRatio(rv)
}
func (d VNDetectRectanglesRequest) SetMaximumAspectRatio(value VNAspectRatio) {
	objc.Send[struct{}](d.ID, objc.Sel("setMaximumAspectRatio:"), value)
}

// A float specifying the number of degrees a rectangle corner angle can
// deviate from 90°.
//
// # Discussion
// 
// The tolerance value should range from `0` to `45`, inclusive. The default
// tolerance is `30`.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/quadratureTolerance
func (d VNDetectRectanglesRequest) QuadratureTolerance() VNDegrees {
	rv := objc.Send[VNDegrees](d.ID, objc.Sel("quadratureTolerance"))
	return VNDegrees(rv)
}
func (d VNDetectRectanglesRequest) SetQuadratureTolerance(value VNDegrees) {
	objc.Send[struct{}](d.ID, objc.Sel("setQuadratureTolerance:"), value)
}

// The minimum size of a rectangle to detect, as a proportion of the smallest
// dimension.
//
// # Discussion
// 
// The value should range from `0.0` to `1.0` inclusive. The default minimum
// size is `0.2`.
// 
// Any smaller rectangles that Vision may have detected aren’t returned.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/minimumSize
func (d VNDetectRectanglesRequest) MinimumSize() float32 {
	rv := objc.Send[float32](d.ID, objc.Sel("minimumSize"))
	return rv
}
func (d VNDetectRectanglesRequest) SetMinimumSize(value float32) {
	objc.Send[struct{}](d.ID, objc.Sel("setMinimumSize:"), value)
}

// A value specifying the minimum acceptable confidence level.
//
// # Discussion
// 
// Vision won’t return rectangles with a confidence score lower than the
// specified minimum.
// 
// The confidence score ranges from `0.0` to `1.0`, inclusive, where `0.0`
// represents no confidence, and `1.0` represents full confidence.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/minimumConfidence
func (d VNDetectRectanglesRequest) MinimumConfidence() VNConfidence {
	rv := objc.Send[VNConfidence](d.ID, objc.Sel("minimumConfidence"))
	return VNConfidence(rv)
}
func (d VNDetectRectanglesRequest) SetMinimumConfidence(value VNConfidence) {
	objc.Send[struct{}](d.ID, objc.Sel("setMinimumConfidence:"), value)
}

// An integer specifying the maximum number of rectangles Vision returns.
//
// # Discussion
// 
// The default value is `1`.
// 
// Setting this property to `0` allows Vision algorithms to return an
// unlimited number of observations.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/maximumObservations
func (d VNDetectRectanglesRequest) MaximumObservations() uint {
	rv := objc.Send[uint](d.ID, objc.Sel("maximumObservations"))
	return rv
}
func (d VNDetectRectanglesRequest) SetMaximumObservations(value uint) {
	objc.Send[struct{}](d.ID, objc.Sel("setMaximumObservations:"), value)
}

// A constant for specifying revision 1 of the rectangle detection request.
//
// See: https://developer.apple.com/documentation/vision/vndetectrectanglesrequestrevision1
func (d VNDetectRectanglesRequest) VNDetectRectanglesRequestRevision1() int {
	rv := objc.Send[int](d.ID, objc.Sel("VNDetectRectanglesRequestRevision1"))
	return rv
}

