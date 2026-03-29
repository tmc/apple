// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VNGeneratePersonSegmentationRequest] class.
var (
	_VNGeneratePersonSegmentationRequestClass     VNGeneratePersonSegmentationRequestClass
	_VNGeneratePersonSegmentationRequestClassOnce sync.Once
)

func getVNGeneratePersonSegmentationRequestClass() VNGeneratePersonSegmentationRequestClass {
	_VNGeneratePersonSegmentationRequestClassOnce.Do(func() {
		_VNGeneratePersonSegmentationRequestClass = VNGeneratePersonSegmentationRequestClass{class: objc.GetClass("VNGeneratePersonSegmentationRequest")}
	})
	return _VNGeneratePersonSegmentationRequestClass
}

// GetVNGeneratePersonSegmentationRequestClass returns the class object for VNGeneratePersonSegmentationRequest.
func GetVNGeneratePersonSegmentationRequestClass() VNGeneratePersonSegmentationRequestClass {
	return getVNGeneratePersonSegmentationRequestClass()
}

type VNGeneratePersonSegmentationRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNGeneratePersonSegmentationRequestClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNGeneratePersonSegmentationRequestClass) Alloc() VNGeneratePersonSegmentationRequest {
	rv := objc.Send[VNGeneratePersonSegmentationRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that produces a matte image for a person it finds in the input
// image.
//
// # Overview
// 
// Perform this request to detect and generate an image mask for a person in
// an image. The request returns the resulting image mask in an instance of
// [VNPixelBufferObservation].
//
// # Configuring the Request
//
//   - [VNGeneratePersonSegmentationRequest.OutputPixelFormat]: The pixel format of the output image.
//   - [VNGeneratePersonSegmentationRequest.SetOutputPixelFormat]
//   - [VNGeneratePersonSegmentationRequest.QualityLevel]: A value that indicates how the request balances accuracy and performance.
//   - [VNGeneratePersonSegmentationRequest.SetQualityLevel]
//
// # Getting the supported output pixel formats
//
//   - [VNGeneratePersonSegmentationRequest.SupportedOutputPixelFormatsAndReturnError]: Returns a list of output pixel formats that the request supports.
//
// # Identifying Request Revisions
//
//   - [VNGeneratePersonSegmentationRequest.VNGeneratePersonSegmentationRequestRevision1]: A constant for specifying revision 1 of the person segmentation generation request.
//
// See: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest
type VNGeneratePersonSegmentationRequest struct {
	VNStatefulRequest
}

// VNGeneratePersonSegmentationRequestFromID constructs a [VNGeneratePersonSegmentationRequest] from an objc.ID.
//
// An object that produces a matte image for a person it finds in the input
// image.
func VNGeneratePersonSegmentationRequestFromID(id objc.ID) VNGeneratePersonSegmentationRequest {
	return VNGeneratePersonSegmentationRequest{VNStatefulRequest: VNStatefulRequestFromID(id)}
}
// NOTE: VNGeneratePersonSegmentationRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNGeneratePersonSegmentationRequest] class.
//
// # Configuring the Request
//
//   - [IVNGeneratePersonSegmentationRequest.OutputPixelFormat]: The pixel format of the output image.
//   - [IVNGeneratePersonSegmentationRequest.SetOutputPixelFormat]
//   - [IVNGeneratePersonSegmentationRequest.QualityLevel]: A value that indicates how the request balances accuracy and performance.
//   - [IVNGeneratePersonSegmentationRequest.SetQualityLevel]
//
// # Getting the supported output pixel formats
//
//   - [IVNGeneratePersonSegmentationRequest.SupportedOutputPixelFormatsAndReturnError]: Returns a list of output pixel formats that the request supports.
//
// # Identifying Request Revisions
//
//   - [IVNGeneratePersonSegmentationRequest.VNGeneratePersonSegmentationRequestRevision1]: A constant for specifying revision 1 of the person segmentation generation request.
//
// See: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest
type IVNGeneratePersonSegmentationRequest interface {
	IVNStatefulRequest

	// Topic: Configuring the Request

	// The pixel format of the output image.
	OutputPixelFormat() uint32
	SetOutputPixelFormat(value uint32)
	// A value that indicates how the request balances accuracy and performance.
	QualityLevel() VNGeneratePersonSegmentationRequestQualityLevel
	SetQualityLevel(value VNGeneratePersonSegmentationRequestQualityLevel)

	// Topic: Getting the supported output pixel formats

	// Returns a list of output pixel formats that the request supports.
	SupportedOutputPixelFormatsAndReturnError() ([]foundation.NSNumber, error)

	// Topic: Identifying Request Revisions

	// A constant for specifying revision 1 of the person segmentation generation request.
	VNGeneratePersonSegmentationRequestRevision1() int
}

// Init initializes the instance.
func (g VNGeneratePersonSegmentationRequest) Init() VNGeneratePersonSegmentationRequest {
	rv := objc.Send[VNGeneratePersonSegmentationRequest](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g VNGeneratePersonSegmentationRequest) Autorelease() VNGeneratePersonSegmentationRequest {
	rv := objc.Send[VNGeneratePersonSegmentationRequest](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNGeneratePersonSegmentationRequest creates a new VNGeneratePersonSegmentationRequest instance.
func NewVNGeneratePersonSegmentationRequest() VNGeneratePersonSegmentationRequest {
	class := getVNGeneratePersonSegmentationRequestClass()
	rv := objc.Send[VNGeneratePersonSegmentationRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a generate person segmentation request with a completion handler.
//
// completionHandler: A completion handler that processes the resuts of the request.
//
// See: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest/init(completionHandler:)
func NewGeneratePersonSegmentationRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNGeneratePersonSegmentationRequest {
	instance := getVNGeneratePersonSegmentationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNGeneratePersonSegmentationRequestFromID(rv)
}

// Initializes a video-based request.
//
// frameAnalysisSpacing: A [CMTime] value that indicates the duration between analysis operations.
// Increase this value to reduce the number of frames analyzed on slower
// devices. Set this argument to [zero] to analyze all frames.
// //
// [CMTime]: https://developer.apple.com/documentation/CoreMedia/CMTime
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
//
// completionHandler: A closure that’s invoked after the request has completed its processing.
// The system invokes the completion handler on the same dispatch queue as the
// request performs its processing.
//
// See: https://developer.apple.com/documentation/Vision/VNStatefulRequest/init(frameAnalysisSpacing:completionHandler:)
func NewGeneratePersonSegmentationRequestWithFrameAnalysisSpacingCompletionHandler(frameAnalysisSpacing coremedia.CMTime, completionHandler VNRequestCompletionHandler) VNGeneratePersonSegmentationRequest {
	instance := getVNGeneratePersonSegmentationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrameAnalysisSpacing:completionHandler:"), frameAnalysisSpacing, completionHandler)
	return VNGeneratePersonSegmentationRequestFromID(rv)
}

// Returns a list of output pixel formats that the request supports.
//
// See: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest/supportedOutputPixelFormats()
func (g VNGeneratePersonSegmentationRequest) SupportedOutputPixelFormatsAndReturnError() ([]foundation.NSNumber, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("supportedOutputPixelFormatsAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	}), nil

}

// The pixel format of the output image.
//
// # Discussion
// 
// The property supports the following values:
// 
// - [kCVPixelFormatType_OneComponent8] -
// [kCVPixelFormatType_OneComponent16Half] -
// [kCVPixelFormatType_OneComponent32Float]
// 
// The default value is [kCVPixelFormatType_OneComponent8].
//
// [kCVPixelFormatType_OneComponent16Half]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_OneComponent16Half
// [kCVPixelFormatType_OneComponent32Float]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_OneComponent32Float
// [kCVPixelFormatType_OneComponent8]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_OneComponent8
//
// See: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest/outputPixelFormat
func (g VNGeneratePersonSegmentationRequest) OutputPixelFormat() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("outputPixelFormat"))
	return rv
}
func (g VNGeneratePersonSegmentationRequest) SetOutputPixelFormat(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setOutputPixelFormat:"), value)
}
// A value that indicates how the request balances accuracy and performance.
//
// See: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest/qualityLevel-swift.property
func (g VNGeneratePersonSegmentationRequest) QualityLevel() VNGeneratePersonSegmentationRequestQualityLevel {
	rv := objc.Send[VNGeneratePersonSegmentationRequestQualityLevel](g.ID, objc.Sel("qualityLevel"))
	return VNGeneratePersonSegmentationRequestQualityLevel(rv)
}
func (g VNGeneratePersonSegmentationRequest) SetQualityLevel(value VNGeneratePersonSegmentationRequestQualityLevel) {
	objc.Send[struct{}](g.ID, objc.Sel("setQualityLevel:"), value)
}
// A constant for specifying revision 1 of the person segmentation generation
// request.
//
// See: https://developer.apple.com/documentation/vision/vngeneratepersonsegmentationrequestrevision1
func (g VNGeneratePersonSegmentationRequest) VNGeneratePersonSegmentationRequestRevision1() int {
	rv := objc.Send[int](g.ID, objc.Sel("VNGeneratePersonSegmentationRequestRevision1"))
	return rv
}

