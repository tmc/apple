// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coreml"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VNCoreMLRequest] class.
var (
	_VNCoreMLRequestClass     VNCoreMLRequestClass
	_VNCoreMLRequestClassOnce sync.Once
)

func getVNCoreMLRequestClass() VNCoreMLRequestClass {
	_VNCoreMLRequestClassOnce.Do(func() {
		_VNCoreMLRequestClass = VNCoreMLRequestClass{class: objc.GetClass("VNCoreMLRequest")}
	})
	return _VNCoreMLRequestClass
}

// GetVNCoreMLRequestClass returns the class object for VNCoreMLRequest.
func GetVNCoreMLRequestClass() VNCoreMLRequestClass {
	return getVNCoreMLRequestClass()
}

type VNCoreMLRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNCoreMLRequestClass) Alloc() VNCoreMLRequest {
	rv := objc.Send[VNCoreMLRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An image-analysis request that uses a Core ML model to process images.
//
// # Overview
// 
// The results array of a Core ML-based image analysis request contains a
// different observation type, depending on the kind of [MLModel] object you
// use:
// 
// - If the model predicts a single feature, the model’s [VNCoreMLRequest.ModelDescription]
// object has a non-`nil` value for [VNCoreMLRequest.PredictedFeatureName] and Vision treats
// the model as a classifier. The results are [VNClassificationObservation]
// objects. - If the model’s outputs include at least one output with a
// feature type of [MLFeatureType.image], Vision treats that model as an
// image-to-image model. The results are [VNPixelBufferObservation] objects. -
// Otherwise, Vision treats the model as a general predictor model. The
// results are [VNCoreMLFeatureValueObservation] objects.
//
// [MLFeatureType.image]: https://developer.apple.com/documentation/CoreML/MLFeatureType/image
// [MLModel]: https://developer.apple.com/documentation/CoreML/MLModel
//
// # Initializing with a Core ML Model
//
//   - [VNCoreMLRequest.InitWithModel]: Creates a model container to use with an image analysis request based on the model you provide.
//   - [VNCoreMLRequest.InitWithModelCompletionHandler]: Creates a model container to use with an image analysis request based on the model you provide, with an optional completion handler.
//   - [VNCoreMLRequest.Model]: The model to base the image analysis request on.
//
// # Configuring Image Options
//
//   - [VNCoreMLRequest.ImageCropAndScaleOption]: An optional setting that tells the Vision algorithm how to scale an input image.
//   - [VNCoreMLRequest.SetImageCropAndScaleOption]
//
// # Identifying Request Revisions
//
//   - [VNCoreMLRequest.VNCoreMLRequestRevision1]: A constant for specifying revision 1 of a Core ML request.
//
// See: https://developer.apple.com/documentation/Vision/VNCoreMLRequest
type VNCoreMLRequest struct {
	VNImageBasedRequest
}

// VNCoreMLRequestFromID constructs a [VNCoreMLRequest] from an objc.ID.
//
// An image-analysis request that uses a Core ML model to process images.
func VNCoreMLRequestFromID(id objc.ID) VNCoreMLRequest {
	return VNCoreMLRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}
// NOTE: VNCoreMLRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNCoreMLRequest] class.
//
// # Initializing with a Core ML Model
//
//   - [IVNCoreMLRequest.InitWithModel]: Creates a model container to use with an image analysis request based on the model you provide.
//   - [IVNCoreMLRequest.InitWithModelCompletionHandler]: Creates a model container to use with an image analysis request based on the model you provide, with an optional completion handler.
//   - [IVNCoreMLRequest.Model]: The model to base the image analysis request on.
//
// # Configuring Image Options
//
//   - [IVNCoreMLRequest.ImageCropAndScaleOption]: An optional setting that tells the Vision algorithm how to scale an input image.
//   - [IVNCoreMLRequest.SetImageCropAndScaleOption]
//
// # Identifying Request Revisions
//
//   - [IVNCoreMLRequest.VNCoreMLRequestRevision1]: A constant for specifying revision 1 of a Core ML request.
//
// See: https://developer.apple.com/documentation/Vision/VNCoreMLRequest
type IVNCoreMLRequest interface {
	IVNImageBasedRequest

	// Topic: Initializing with a Core ML Model

	// Creates a model container to use with an image analysis request based on the model you provide.
	InitWithModel(model IVNCoreMLModel) VNCoreMLRequest
	// Creates a model container to use with an image analysis request based on the model you provide, with an optional completion handler.
	InitWithModelCompletionHandler(model IVNCoreMLModel, completionHandler ErrorHandler) VNCoreMLRequest
	// The model to base the image analysis request on.
	Model() IVNCoreMLModel

	// Topic: Configuring Image Options

	// An optional setting that tells the Vision algorithm how to scale an input image.
	ImageCropAndScaleOption() VNImageCropAndScaleOption
	SetImageCropAndScaleOption(value VNImageCropAndScaleOption)

	// Topic: Identifying Request Revisions

	// A constant for specifying revision 1 of a Core ML request.
	VNCoreMLRequestRevision1() int

	// The level of confidence in the observation’s accuracy.
	Confidence() VNConfidence
	SetConfidence(value VNConfidence)
	// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
	ModelDescription() coreml.MLModelDescription
	SetModelDescription(value coreml.MLModelDescription)
	// The name of the primary prediction feature output description.
	PredictedFeatureName() string
	SetPredictedFeatureName(value string)
}

// Init initializes the instance.
func (c VNCoreMLRequest) Init() VNCoreMLRequest {
	rv := objc.Send[VNCoreMLRequest](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c VNCoreMLRequest) Autorelease() VNCoreMLRequest {
	rv := objc.Send[VNCoreMLRequest](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNCoreMLRequest creates a new VNCoreMLRequest instance.
func NewVNCoreMLRequest() VNCoreMLRequest {
	class := getVNCoreMLRequestClass()
	rv := objc.Send[VNCoreMLRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewCoreMLRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNCoreMLRequest {
	instance := getVNCoreMLRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNCoreMLRequestFromID(rv)
}

// Creates a model container to use with an image analysis request based on
// the model you provide.
//
// model: The [Core ML] model on which to base the Vision request.
// //
// [Core ML]: https://developer.apple.com/documentation/CoreML
//
// # Discussion
// 
// Initialization can fail if the [Core ML] model you provide isn’t
// supported in Vision, such as if the model doesn’t accept an image as
// input.
//
// [Core ML]: https://developer.apple.com/documentation/CoreML
//
// See: https://developer.apple.com/documentation/Vision/VNCoreMLRequest/init(model:)
func NewCoreMLRequestWithModel(model IVNCoreMLModel) VNCoreMLRequest {
	instance := getVNCoreMLRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModel:"), model)
	return VNCoreMLRequestFromID(rv)
}

// Creates a model container to use with an image analysis request based on
// the model you provide, with an optional completion handler.
//
// model: The [Core ML] model on which to base the Vision request.
// //
// [Core ML]: https://developer.apple.com/documentation/CoreML
//
// completionHandler: An optional block of code to execute after model initialization.
//
// # Discussion
// 
// Initialization can fail if the [Core ML] model you provide isn’t
// supported in Vision, such as if the model doesn’t accept an image as
// input.
//
// [Core ML]: https://developer.apple.com/documentation/CoreML
//
// See: https://developer.apple.com/documentation/Vision/VNCoreMLRequest/init(model:completionHandler:)
func NewCoreMLRequestWithModelCompletionHandler(model IVNCoreMLModel, completionHandler VNRequestCompletionHandler) VNCoreMLRequest {
	instance := getVNCoreMLRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModel:completionHandler:"), model, completionHandler)
	return VNCoreMLRequestFromID(rv)
}

// Creates a model container to use with an image analysis request based on
// the model you provide.
//
// model: The [Core ML] model on which to base the Vision request.
// //
// [Core ML]: https://developer.apple.com/documentation/CoreML
//
// # Discussion
// 
// Initialization can fail if the [Core ML] model you provide isn’t
// supported in Vision, such as if the model doesn’t accept an image as
// input.
//
// [Core ML]: https://developer.apple.com/documentation/CoreML
//
// See: https://developer.apple.com/documentation/Vision/VNCoreMLRequest/init(model:)
func (c VNCoreMLRequest) InitWithModel(model IVNCoreMLModel) VNCoreMLRequest {
	rv := objc.Send[VNCoreMLRequest](c.ID, objc.Sel("initWithModel:"), model)
	return rv
}

// Creates a model container to use with an image analysis request based on
// the model you provide, with an optional completion handler.
//
// model: The [Core ML] model on which to base the Vision request.
// //
// [Core ML]: https://developer.apple.com/documentation/CoreML
//
// completionHandler: An optional block of code to execute after model initialization.
//
// # Discussion
// 
// Initialization can fail if the [Core ML] model you provide isn’t
// supported in Vision, such as if the model doesn’t accept an image as
// input.
//
// [Core ML]: https://developer.apple.com/documentation/CoreML
//
// See: https://developer.apple.com/documentation/Vision/VNCoreMLRequest/init(model:completionHandler:)
func (c VNCoreMLRequest) InitWithModelCompletionHandler(model IVNCoreMLModel, completionHandler ErrorHandler) VNCoreMLRequest {
_block1, _cleanup1 := NewErrorBlock(completionHandler)
	defer _cleanup1()
	rv := objc.Send[objc.ID](c.ID, objc.Sel("initWithModel:completionHandler:"), model, _block1)
	return VNCoreMLRequestFromID(rv)
}

// The model to base the image analysis request on.
//
// # Discussion
// 
// This object wraps a [Core ML] model.
//
// [Core ML]: https://developer.apple.com/documentation/CoreML
//
// See: https://developer.apple.com/documentation/Vision/VNCoreMLRequest/model
func (c VNCoreMLRequest) Model() IVNCoreMLModel {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("model"))
	return VNCoreMLModelFromID(objc.ID(rv))
}

// An optional setting that tells the Vision algorithm how to scale an input
// image.
//
// # Discussion
// 
// Scaling an image ensures that the entire image fits into the algorithm’s
// input image dimensions, which may require a change in aspect ratio. Each
// crop-and-scale option transforms the input image in a different way.
// 
// [scale-crop-options]
//
// See: https://developer.apple.com/documentation/Vision/VNCoreMLRequest/imageCropAndScaleOption
func (c VNCoreMLRequest) ImageCropAndScaleOption() VNImageCropAndScaleOption {
	rv := objc.Send[VNImageCropAndScaleOption](c.ID, objc.Sel("imageCropAndScaleOption"))
	return VNImageCropAndScaleOption(rv)
}
func (c VNCoreMLRequest) SetImageCropAndScaleOption(value VNImageCropAndScaleOption) {
	objc.Send[struct{}](c.ID, objc.Sel("setImageCropAndScaleOption:"), value)
}

// A constant for specifying revision 1 of a Core ML request.
//
// See: https://developer.apple.com/documentation/vision/vncoremlrequestrevision1
func (c VNCoreMLRequest) VNCoreMLRequestRevision1() int {
	rv := objc.Send[int](c.ID, objc.Sel("VNCoreMLRequestRevision1"))
	return rv
}

// The level of confidence in the observation’s accuracy.
//
// See: https://developer.apple.com/documentation/vision/vnobservation/confidence
func (c VNCoreMLRequest) Confidence() VNConfidence {
	rv := objc.Send[VNConfidence](c.ID, objc.Sel("confidence"))
	return VNConfidence(rv)
}
func (c VNCoreMLRequest) SetConfidence(value VNConfidence) {
	objc.Send[struct{}](c.ID, objc.Sel("setConfidence:"), value)
}

// Model information you use at runtime during development, which Xcode also
// displays in its Core ML model editor view.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (c VNCoreMLRequest) ModelDescription() coreml.MLModelDescription {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("modelDescription"))
	return coreml.MLModelDescriptionFromID(objc.ID(rv))
}
func (c VNCoreMLRequest) SetModelDescription(value coreml.MLModelDescription) {
	objc.Send[struct{}](c.ID, objc.Sel("setModelDescription:"), value)
}

// The name of the primary prediction feature output description.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (c VNCoreMLRequest) PredictedFeatureName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("predictedFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (c VNCoreMLRequest) SetPredictedFeatureName(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setPredictedFeatureName:"), objc.String(value))
}

