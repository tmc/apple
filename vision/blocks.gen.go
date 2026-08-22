// Code generated from Apple documentation. DO NOT EDIT.

package vision

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// VNRequestCompletionHandler handles A type alias to encapsulate the syntax for the completion handler the system calls after the request finishes processing.

// NewVNRequestCompletionHandlerBlock wraps a Go [VNRequestCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewVNRequestCompletionHandlerBlock(handler VNRequestCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive VNRequest, extra0 foundation.NSError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// VNRequestErrorHandler handles The block to invoke after the request finishes processing.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [VNCalculateImageAestheticsScoresRequest.InitWithCompletionHandler]
//   - [VNClassifyImageRequest.InitWithCompletionHandler]
//   - [VNCoreMLRequest.InitWithCompletionHandler]
//   - [VNCoreMLRequest.InitWithModelCompletionHandler]
//   - [VNDetectAnimalBodyPoseRequest.InitWithCompletionHandler]
//   - [VNDetectBarcodesRequest.InitWithCompletionHandler]
//   - [VNDetectContoursRequest.InitWithCompletionHandler]
//   - [VNDetectDocumentSegmentationRequest.InitWithCompletionHandler]
//   - [VNDetectFaceCaptureQualityRequest.InitWithCompletionHandler]
//   - [VNDetectFaceLandmarksRequest.InitWithCompletionHandler]
//   - [VNDetectFaceRectanglesRequest.InitWithCompletionHandler]
//   - [VNDetectHorizonRequest.InitWithCompletionHandler]
//   - [VNDetectHumanBodyPose3DRequest.InitWithCompletionHandler]
//   - [VNDetectHumanBodyPose3DRequest.InitWithFrameAnalysisSpacingCompletionHandler]
//   - [VNDetectHumanBodyPoseRequest.InitWithCompletionHandler]
//   - [VNDetectHumanHandPoseRequest.InitWithCompletionHandler]
//   - [VNDetectHumanRectanglesRequest.InitWithCompletionHandler]
//   - [VNDetectRectanglesRequest.InitWithCompletionHandler]
//   - [VNDetectTextRectanglesRequest.InitWithCompletionHandler]
//   - [VNDetectTrajectoriesRequest.InitWithCompletionHandler]
//   - [VNDetectTrajectoriesRequest.InitWithFrameAnalysisSpacingCompletionHandler]
//   - [VNDetectTrajectoriesRequest.InitWithFrameAnalysisSpacingTrajectoryLengthCompletionHandler]
//   - [VNGenerateAttentionBasedSaliencyImageRequest.InitWithCompletionHandler]
//   - [VNGenerateForegroundInstanceMaskRequest.InitWithCompletionHandler]
//   - [VNGenerateImageFeaturePrintRequest.InitWithCompletionHandler]
//   - [VNGenerateObjectnessBasedSaliencyImageRequest.InitWithCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedCGImageOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedCGImageOrientationOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedCIImageOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedCIImageOrientationOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedCMSampleBufferOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedCMSampleBufferOrientationOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedCVPixelBufferOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedCVPixelBufferOrientationOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedImageDataOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedImageDataOrientationOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedImageURLOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedImageURLOrientationOptionsCompletionHandler]
//   - [VNGeneratePersonInstanceMaskRequest.InitWithCompletionHandler]
//   - [VNGeneratePersonSegmentationRequest.InitWithCompletionHandler]
//   - [VNGeneratePersonSegmentationRequest.InitWithFrameAnalysisSpacingCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedCGImageOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedCGImageOrientationOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedCIImageOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedCIImageOrientationOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedCMSampleBufferOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedCMSampleBufferOrientationOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedCVPixelBufferOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedCVPixelBufferOrientationOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedImageDataOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedImageDataOrientationOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedImageURLOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedImageURLOrientationOptionsCompletionHandler]
//   - [VNImageBasedRequest.InitWithCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedCGImageOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedCGImageOrientationOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedCIImageOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedCIImageOrientationOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedCMSampleBufferOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedCMSampleBufferOrientationOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedCVPixelBufferOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedCVPixelBufferOrientationOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedImageDataOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedImageDataOrientationOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedImageURLOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedImageURLOrientationOptionsCompletionHandler]
//   - [VNRecognizeAnimalsRequest.InitWithCompletionHandler]
//   - [VNRecognizeTextRequest.InitWithCompletionHandler]
//   - [VNRequest.InitWithCompletionHandler]
//   - [VNStatefulRequest.InitWithCompletionHandler]
//   - [VNStatefulRequest.InitWithFrameAnalysisSpacingCompletionHandler]
//   - [VNTargetedImageRequest.InitWithCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedCGImageOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedCGImageOrientationOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedCIImageOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedCIImageOrientationOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedCMSampleBufferOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedCMSampleBufferOrientationOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedCVPixelBufferOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedCVPixelBufferOrientationOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedImageDataOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedImageDataOrientationOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedImageURLOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedImageURLOrientationOptionsCompletionHandler]
//   - [VNTrackHomographicImageRegistrationRequest.InitWithCompletionHandler]
//   - [VNTrackHomographicImageRegistrationRequest.InitWithFrameAnalysisSpacingCompletionHandler]
//   - [VNTrackObjectRequest.InitWithCompletionHandler]
//   - [VNTrackObjectRequest.InitWithDetectedObjectObservationCompletionHandler]
//   - [VNTrackOpticalFlowRequest.InitWithCompletionHandler]
//   - [VNTrackOpticalFlowRequest.InitWithFrameAnalysisSpacingCompletionHandler]
//   - [VNTrackRectangleRequest.InitWithCompletionHandler]
//   - [VNTrackRectangleRequest.InitWithRectangleObservationCompletionHandler]
//   - [VNTrackTranslationalImageRegistrationRequest.InitWithCompletionHandler]
//   - [VNTrackTranslationalImageRegistrationRequest.InitWithFrameAnalysisSpacingCompletionHandler]
//   - [VNTrackingRequest.InitWithCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedCGImageOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedCGImageOrientationOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedCIImageOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedCIImageOrientationOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedCMSampleBufferOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedCMSampleBufferOrientationOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedCVPixelBufferOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedCVPixelBufferOrientationOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedImageDataOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedImageDataOrientationOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedImageURLOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedImageURLOrientationOptionsCompletionHandler]
type VNRequestErrorHandler = func(*VNRequest, error)

// NewVNRequestErrorBlock wraps a Go [VNRequestErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [VNCalculateImageAestheticsScoresRequest.InitWithCompletionHandler]
//   - [VNClassifyImageRequest.InitWithCompletionHandler]
//   - [VNCoreMLRequest.InitWithCompletionHandler]
//   - [VNCoreMLRequest.InitWithModelCompletionHandler]
//   - [VNDetectAnimalBodyPoseRequest.InitWithCompletionHandler]
//   - [VNDetectBarcodesRequest.InitWithCompletionHandler]
//   - [VNDetectContoursRequest.InitWithCompletionHandler]
//   - [VNDetectDocumentSegmentationRequest.InitWithCompletionHandler]
//   - [VNDetectFaceCaptureQualityRequest.InitWithCompletionHandler]
//   - [VNDetectFaceLandmarksRequest.InitWithCompletionHandler]
//   - [VNDetectFaceRectanglesRequest.InitWithCompletionHandler]
//   - [VNDetectHorizonRequest.InitWithCompletionHandler]
//   - [VNDetectHumanBodyPose3DRequest.InitWithCompletionHandler]
//   - [VNDetectHumanBodyPose3DRequest.InitWithFrameAnalysisSpacingCompletionHandler]
//   - [VNDetectHumanBodyPoseRequest.InitWithCompletionHandler]
//   - [VNDetectHumanHandPoseRequest.InitWithCompletionHandler]
//   - [VNDetectHumanRectanglesRequest.InitWithCompletionHandler]
//   - [VNDetectRectanglesRequest.InitWithCompletionHandler]
//   - [VNDetectTextRectanglesRequest.InitWithCompletionHandler]
//   - [VNDetectTrajectoriesRequest.InitWithCompletionHandler]
//   - [VNDetectTrajectoriesRequest.InitWithFrameAnalysisSpacingCompletionHandler]
//   - [VNDetectTrajectoriesRequest.InitWithFrameAnalysisSpacingTrajectoryLengthCompletionHandler]
//   - [VNGenerateAttentionBasedSaliencyImageRequest.InitWithCompletionHandler]
//   - [VNGenerateForegroundInstanceMaskRequest.InitWithCompletionHandler]
//   - [VNGenerateImageFeaturePrintRequest.InitWithCompletionHandler]
//   - [VNGenerateObjectnessBasedSaliencyImageRequest.InitWithCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedCGImageOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedCGImageOrientationOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedCIImageOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedCIImageOrientationOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedCMSampleBufferOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedCMSampleBufferOrientationOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedCVPixelBufferOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedCVPixelBufferOrientationOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedImageDataOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedImageDataOrientationOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedImageURLOptionsCompletionHandler]
//   - [VNGenerateOpticalFlowRequest.InitWithTargetedImageURLOrientationOptionsCompletionHandler]
//   - [VNGeneratePersonInstanceMaskRequest.InitWithCompletionHandler]
//   - [VNGeneratePersonSegmentationRequest.InitWithCompletionHandler]
//   - [VNGeneratePersonSegmentationRequest.InitWithFrameAnalysisSpacingCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedCGImageOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedCGImageOrientationOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedCIImageOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedCIImageOrientationOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedCMSampleBufferOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedCMSampleBufferOrientationOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedCVPixelBufferOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedCVPixelBufferOrientationOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedImageDataOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedImageDataOrientationOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedImageURLOptionsCompletionHandler]
//   - [VNHomographicImageRegistrationRequest.InitWithTargetedImageURLOrientationOptionsCompletionHandler]
//   - [VNImageBasedRequest.InitWithCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedCGImageOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedCGImageOrientationOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedCIImageOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedCIImageOrientationOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedCMSampleBufferOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedCMSampleBufferOrientationOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedCVPixelBufferOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedCVPixelBufferOrientationOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedImageDataOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedImageDataOrientationOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedImageURLOptionsCompletionHandler]
//   - [VNImageRegistrationRequest.InitWithTargetedImageURLOrientationOptionsCompletionHandler]
//   - [VNRecognizeAnimalsRequest.InitWithCompletionHandler]
//   - [VNRecognizeTextRequest.InitWithCompletionHandler]
//   - [VNRequest.InitWithCompletionHandler]
//   - [VNStatefulRequest.InitWithCompletionHandler]
//   - [VNStatefulRequest.InitWithFrameAnalysisSpacingCompletionHandler]
//   - [VNTargetedImageRequest.InitWithCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedCGImageOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedCGImageOrientationOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedCIImageOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedCIImageOrientationOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedCMSampleBufferOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedCMSampleBufferOrientationOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedCVPixelBufferOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedCVPixelBufferOrientationOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedImageDataOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedImageDataOrientationOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedImageURLOptionsCompletionHandler]
//   - [VNTargetedImageRequest.InitWithTargetedImageURLOrientationOptionsCompletionHandler]
//   - [VNTrackHomographicImageRegistrationRequest.InitWithCompletionHandler]
//   - [VNTrackHomographicImageRegistrationRequest.InitWithFrameAnalysisSpacingCompletionHandler]
//   - [VNTrackObjectRequest.InitWithCompletionHandler]
//   - [VNTrackObjectRequest.InitWithDetectedObjectObservationCompletionHandler]
//   - [VNTrackOpticalFlowRequest.InitWithCompletionHandler]
//   - [VNTrackOpticalFlowRequest.InitWithFrameAnalysisSpacingCompletionHandler]
//   - [VNTrackRectangleRequest.InitWithCompletionHandler]
//   - [VNTrackRectangleRequest.InitWithRectangleObservationCompletionHandler]
//   - [VNTrackTranslationalImageRegistrationRequest.InitWithCompletionHandler]
//   - [VNTrackTranslationalImageRegistrationRequest.InitWithFrameAnalysisSpacingCompletionHandler]
//   - [VNTrackingRequest.InitWithCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedCGImageOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedCGImageOrientationOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedCIImageOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedCIImageOrientationOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedCMSampleBufferOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedCMSampleBufferOrientationOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedCVPixelBufferOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedCVPixelBufferOrientationOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedImageDataOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedImageDataOrientationOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedImageURLOptionsCompletionHandler]
//   - [VNTranslationalImageRegistrationRequest.InitWithTargetedImageURLOrientationOptionsCompletionHandler]
func NewVNRequestErrorBlock(handler VNRequestErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *VNRequest
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := VNRequestFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// VNRequestFloat64ErrorHandler is the signature for a completion handler block.
type VNRequestFloat64ErrorHandler = func(*VNRequest, float64, error)

// NewVNRequestFloat64ErrorBlock wraps a Go [VNRequestFloat64ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewVNRequestFloat64ErrorBlock(handler VNRequestFloat64ErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 float64, errID objc.ID) {
		var result *VNRequest
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := VNRequestFromID(resultID)
			result = &v
		}
		handler(result, extra0, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// VNRequestProgressHandler handles A block executed at intervals during the processing of a Vision request.

// NewVNRequestProgressHandlerBlock wraps a Go [VNRequestProgressHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewVNRequestProgressHandlerBlock(handler VNRequestProgressHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive VNRequest, extra0 float64, extra1 foundation.NSError) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}
