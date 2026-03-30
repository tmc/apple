// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VNGenerateImageFeaturePrintRequest] class.
var (
	_VNGenerateImageFeaturePrintRequestClass     VNGenerateImageFeaturePrintRequestClass
	_VNGenerateImageFeaturePrintRequestClassOnce sync.Once
)

func getVNGenerateImageFeaturePrintRequestClass() VNGenerateImageFeaturePrintRequestClass {
	_VNGenerateImageFeaturePrintRequestClassOnce.Do(func() {
		_VNGenerateImageFeaturePrintRequestClass = VNGenerateImageFeaturePrintRequestClass{class: objc.GetClass("VNGenerateImageFeaturePrintRequest")}
	})
	return _VNGenerateImageFeaturePrintRequestClass
}

// GetVNGenerateImageFeaturePrintRequestClass returns the class object for VNGenerateImageFeaturePrintRequest.
func GetVNGenerateImageFeaturePrintRequestClass() VNGenerateImageFeaturePrintRequestClass {
	return getVNGenerateImageFeaturePrintRequestClass()
}

type VNGenerateImageFeaturePrintRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNGenerateImageFeaturePrintRequestClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNGenerateImageFeaturePrintRequestClass) Alloc() VNGenerateImageFeaturePrintRequest {
	rv := objc.Send[VNGenerateImageFeaturePrintRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An image-based request to generate feature prints from an image.
//
// # Overview
//
// This request returns the feature print data it generates as an array of
// [VNFeaturePrintObservation] objects.
//
// # Scaling and Cropping Images
//
//   - [VNGenerateImageFeaturePrintRequest.ImageCropAndScaleOption]: An optional setting that tells the algorithm how to scale an input image before generating the feature print.
//   - [VNGenerateImageFeaturePrintRequest.SetImageCropAndScaleOption]
//
// # Identifying Request Revisions
//
//   - [VNGenerateImageFeaturePrintRequest.VNGenerateImageFeaturePrintRequestRevision1]: A constant for specifying the first revision of the feature-print request.
//
// See: https://developer.apple.com/documentation/Vision/VNGenerateImageFeaturePrintRequest
type VNGenerateImageFeaturePrintRequest struct {
	VNImageBasedRequest
}

// VNGenerateImageFeaturePrintRequestFromID constructs a [VNGenerateImageFeaturePrintRequest] from an objc.ID.
//
// An image-based request to generate feature prints from an image.
func VNGenerateImageFeaturePrintRequestFromID(id objc.ID) VNGenerateImageFeaturePrintRequest {
	return VNGenerateImageFeaturePrintRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}

// NOTE: VNGenerateImageFeaturePrintRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNGenerateImageFeaturePrintRequest] class.
//
// # Scaling and Cropping Images
//
//   - [IVNGenerateImageFeaturePrintRequest.ImageCropAndScaleOption]: An optional setting that tells the algorithm how to scale an input image before generating the feature print.
//   - [IVNGenerateImageFeaturePrintRequest.SetImageCropAndScaleOption]
//
// # Identifying Request Revisions
//
//   - [IVNGenerateImageFeaturePrintRequest.VNGenerateImageFeaturePrintRequestRevision1]: A constant for specifying the first revision of the feature-print request.
//
// See: https://developer.apple.com/documentation/Vision/VNGenerateImageFeaturePrintRequest
type IVNGenerateImageFeaturePrintRequest interface {
	IVNImageBasedRequest

	// Topic: Scaling and Cropping Images

	// An optional setting that tells the algorithm how to scale an input image before generating the feature print.
	ImageCropAndScaleOption() VNImageCropAndScaleOption
	SetImageCropAndScaleOption(value VNImageCropAndScaleOption)

	// Topic: Identifying Request Revisions

	// A constant for specifying the first revision of the feature-print request.
	VNGenerateImageFeaturePrintRequestRevision1() int
}

// Init initializes the instance.
func (g VNGenerateImageFeaturePrintRequest) Init() VNGenerateImageFeaturePrintRequest {
	rv := objc.Send[VNGenerateImageFeaturePrintRequest](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g VNGenerateImageFeaturePrintRequest) Autorelease() VNGenerateImageFeaturePrintRequest {
	rv := objc.Send[VNGenerateImageFeaturePrintRequest](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNGenerateImageFeaturePrintRequest creates a new VNGenerateImageFeaturePrintRequest instance.
func NewVNGenerateImageFeaturePrintRequest() VNGenerateImageFeaturePrintRequest {
	class := getVNGenerateImageFeaturePrintRequestClass()
	rv := objc.Send[VNGenerateImageFeaturePrintRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewGenerateImageFeaturePrintRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNGenerateImageFeaturePrintRequest {
	instance := getVNGenerateImageFeaturePrintRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNGenerateImageFeaturePrintRequestFromID(rv)
}

// An optional setting that tells the algorithm how to scale an input image
// before generating the feature print.
//
// # Discussion
//
// Scaling is applied before generating the feature print. The default value
// is [VNImageCropAndScaleOptionScaleFill].
//
// Scaling an image ensures that the entire image fits into the algorithm’s
// input image dimensions, which may require a change in aspect ratio. Each
// crop and scale option transforms the input image in a different way:
//
// [scale-crop-options]
//
// See: https://developer.apple.com/documentation/Vision/VNGenerateImageFeaturePrintRequest/imageCropAndScaleOption
func (g VNGenerateImageFeaturePrintRequest) ImageCropAndScaleOption() VNImageCropAndScaleOption {
	rv := objc.Send[VNImageCropAndScaleOption](g.ID, objc.Sel("imageCropAndScaleOption"))
	return VNImageCropAndScaleOption(rv)
}
func (g VNGenerateImageFeaturePrintRequest) SetImageCropAndScaleOption(value VNImageCropAndScaleOption) {
	objc.Send[struct{}](g.ID, objc.Sel("setImageCropAndScaleOption:"), value)
}

// A constant for specifying the first revision of the feature-print request.
//
// See: https://developer.apple.com/documentation/vision/vngenerateimagefeatureprintrequestrevision1
func (g VNGenerateImageFeaturePrintRequest) VNGenerateImageFeaturePrintRequestRevision1() int {
	rv := objc.Send[int](g.ID, objc.Sel("VNGenerateImageFeaturePrintRequestRevision1"))
	return rv
}
