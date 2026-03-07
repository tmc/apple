// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/Vision/VNBarcodeCompositeType
type VNBarcodeCompositeType int

const (
	// VNBarcodeCompositeTypeGS1TypeA: A type that represents trade items in bulk.
	VNBarcodeCompositeTypeGS1TypeA VNBarcodeCompositeType = 2
	// VNBarcodeCompositeTypeGS1TypeB: A type that represents trade items by piece.
	VNBarcodeCompositeTypeGS1TypeB VNBarcodeCompositeType = 3
	// VNBarcodeCompositeTypeGS1TypeC: A type that represents trade items in varying quantity.
	VNBarcodeCompositeTypeGS1TypeC VNBarcodeCompositeType = 4
	// VNBarcodeCompositeTypeLinked: A type that represents a linked composite type.
	VNBarcodeCompositeTypeLinked VNBarcodeCompositeType = 1
	// VNBarcodeCompositeTypeNone: A type that represents no composite type.
	VNBarcodeCompositeTypeNone VNBarcodeCompositeType = 0
)


func (e VNBarcodeCompositeType) String() string {
	switch e {
	case VNBarcodeCompositeTypeGS1TypeA:
		return "VNBarcodeCompositeTypeGS1TypeA"
	case VNBarcodeCompositeTypeGS1TypeB:
		return "VNBarcodeCompositeTypeGS1TypeB"
	case VNBarcodeCompositeTypeGS1TypeC:
		return "VNBarcodeCompositeTypeGS1TypeC"
	case VNBarcodeCompositeTypeLinked:
		return "VNBarcodeCompositeTypeLinked"
	case VNBarcodeCompositeTypeNone:
		return "VNBarcodeCompositeTypeNone"
	default:
		return fmt.Sprintf("VNBarcodeCompositeType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/Vision/VNChirality
type VNChirality int

const (
	// VNChiralityLeft: Indicates a left-handed pose.
	VNChiralityLeft VNChirality = -1
	// VNChiralityRight: Indicates a right-handed pose.
	VNChiralityRight VNChirality = 1
	// VNChiralityUnknown: Indicates that the pose chirality is unknown.
	VNChiralityUnknown VNChirality = 0
)


func (e VNChirality) String() string {
	switch e {
	case VNChiralityLeft:
		return "VNChiralityLeft"
	case VNChiralityRight:
		return "VNChiralityRight"
	case VNChiralityUnknown:
		return "VNChiralityUnknown"
	default:
		return fmt.Sprintf("VNChirality(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/Vision/VNElementType
type VNElementType int

const (
	// VNElementTypeDouble: The elements are double-precision floating-point numbers.
	VNElementTypeDouble VNElementType = 2
	// VNElementTypeFloat: The elements are floating-point numbers.
	VNElementTypeFloat VNElementType = 1
	// VNElementTypeUnknown: The element type isn’t known.
	VNElementTypeUnknown VNElementType = 0
)


func (e VNElementType) String() string {
	switch e {
	case VNElementTypeDouble:
		return "VNElementTypeDouble"
	case VNElementTypeFloat:
		return "VNElementTypeFloat"
	case VNElementTypeUnknown:
		return "VNElementTypeUnknown"
	default:
		return fmt.Sprintf("VNElementType(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/Vision/VNErrorCode
type VNErrorCode int

const (
	// VNErrorDataUnavailable: The data isn’t available.
	VNErrorDataUnavailable VNErrorCode = 17
	// VNErrorIOError: An I/O error for an image, image sequence, or Core ML model.
	VNErrorIOError VNErrorCode = 6
	// VNErrorInternalError: An internal error occurred within the framework.
	VNErrorInternalError VNErrorCode = 9
	// VNErrorInvalidArgument: An app passed an invalid parameter to a request.
	VNErrorInvalidArgument VNErrorCode = 14
	// VNErrorInvalidFormat: The format of the image is invalid.
	VNErrorInvalidFormat VNErrorCode = 2
	// VNErrorInvalidImage: The image is invalid.
	VNErrorInvalidImage VNErrorCode = 13
	// VNErrorInvalidModel: The Core ML model is incompatible with the request.
	VNErrorInvalidModel VNErrorCode = 15
	// VNErrorInvalidOperation: An app requested an unsupported operation.
	VNErrorInvalidOperation VNErrorCode = 12
	// VNErrorInvalidOption: An app specified an invalid option on a request.
	VNErrorInvalidOption VNErrorCode = 5
	// VNErrorMissingOption: A request is missing a required option.
	VNErrorMissingOption VNErrorCode = 7
	// VNErrorNotImplemented: The method isn’t implemented in the underlying model.
	VNErrorNotImplemented VNErrorCode = 8
	// VNErrorOK: The operation finished without error.
	VNErrorOK VNErrorCode = 0
	// VNErrorOperationFailed: The requested operation failed.
	VNErrorOperationFailed VNErrorCode = 3
	// VNErrorOutOfBoundsError: An app attempted to access data that’s out-of-bounds.
	VNErrorOutOfBoundsError VNErrorCode = 4
	// VNErrorOutOfMemory: The system doesn’t have enough memory to complete the request.
	VNErrorOutOfMemory VNErrorCode = 10
	// VNErrorRequestCancelled: An app canceled the request.
	VNErrorRequestCancelled VNErrorCode = 1
	// VNErrorTimeStampNotFound: The system can’t find a timestamp.
	VNErrorTimeStampNotFound VNErrorCode = 18
	// VNErrorTimeout: The requested operation timed out.
	VNErrorTimeout VNErrorCode = 20
	// VNErrorTuriCoreErrorCode: An error occurred during Create ML training due to an invalid transformation or image.
	VNErrorTuriCoreErrorCode VNErrorCode = -1
	// VNErrorUnknownError: An unidentified error occurred.
	VNErrorUnknownError VNErrorCode = 11
	// VNErrorUnsupportedComputeDevice: An app requested an unsupported compute device.
	VNErrorUnsupportedComputeDevice VNErrorCode = 22
	// VNErrorUnsupportedComputeStage: An app requested an unsupported compute stage.
	VNErrorUnsupportedComputeStage VNErrorCode = 21
	// VNErrorUnsupportedRequest: An app attempted an unsupported request.
	VNErrorUnsupportedRequest VNErrorCode = 19
	// VNErrorUnsupportedRevision: An app specified an unsupported request revision.
	VNErrorUnsupportedRevision VNErrorCode = 16
)


func (e VNErrorCode) String() string {
	switch e {
	case VNErrorDataUnavailable:
		return "VNErrorDataUnavailable"
	case VNErrorIOError:
		return "VNErrorIOError"
	case VNErrorInternalError:
		return "VNErrorInternalError"
	case VNErrorInvalidArgument:
		return "VNErrorInvalidArgument"
	case VNErrorInvalidFormat:
		return "VNErrorInvalidFormat"
	case VNErrorInvalidImage:
		return "VNErrorInvalidImage"
	case VNErrorInvalidModel:
		return "VNErrorInvalidModel"
	case VNErrorInvalidOperation:
		return "VNErrorInvalidOperation"
	case VNErrorInvalidOption:
		return "VNErrorInvalidOption"
	case VNErrorMissingOption:
		return "VNErrorMissingOption"
	case VNErrorNotImplemented:
		return "VNErrorNotImplemented"
	case VNErrorOK:
		return "VNErrorOK"
	case VNErrorOperationFailed:
		return "VNErrorOperationFailed"
	case VNErrorOutOfBoundsError:
		return "VNErrorOutOfBoundsError"
	case VNErrorOutOfMemory:
		return "VNErrorOutOfMemory"
	case VNErrorRequestCancelled:
		return "VNErrorRequestCancelled"
	case VNErrorTimeStampNotFound:
		return "VNErrorTimeStampNotFound"
	case VNErrorTimeout:
		return "VNErrorTimeout"
	case VNErrorTuriCoreErrorCode:
		return "VNErrorTuriCoreErrorCode"
	case VNErrorUnknownError:
		return "VNErrorUnknownError"
	case VNErrorUnsupportedComputeDevice:
		return "VNErrorUnsupportedComputeDevice"
	case VNErrorUnsupportedComputeStage:
		return "VNErrorUnsupportedComputeStage"
	case VNErrorUnsupportedRequest:
		return "VNErrorUnsupportedRequest"
	case VNErrorUnsupportedRevision:
		return "VNErrorUnsupportedRevision"
	default:
		return fmt.Sprintf("VNErrorCode(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/Vision/VNGenerateOpticalFlowRequest/ComputationAccuracy-swift.enum
type VNGenerateOpticalFlowRequestComputationAccuracy int

const (
	// VNGenerateOpticalFlowRequestComputationAccuracyHigh: High accuracy.
	VNGenerateOpticalFlowRequestComputationAccuracyHigh VNGenerateOpticalFlowRequestComputationAccuracy = 2
	// VNGenerateOpticalFlowRequestComputationAccuracyLow: Low accuracy.
	VNGenerateOpticalFlowRequestComputationAccuracyLow VNGenerateOpticalFlowRequestComputationAccuracy = 0
	// VNGenerateOpticalFlowRequestComputationAccuracyMedium: Medium accuracy.
	VNGenerateOpticalFlowRequestComputationAccuracyMedium VNGenerateOpticalFlowRequestComputationAccuracy = 1
	// VNGenerateOpticalFlowRequestComputationAccuracyVeryHigh: Very high accuracy.
	VNGenerateOpticalFlowRequestComputationAccuracyVeryHigh VNGenerateOpticalFlowRequestComputationAccuracy = 3
)


func (e VNGenerateOpticalFlowRequestComputationAccuracy) String() string {
	switch e {
	case VNGenerateOpticalFlowRequestComputationAccuracyHigh:
		return "VNGenerateOpticalFlowRequestComputationAccuracyHigh"
	case VNGenerateOpticalFlowRequestComputationAccuracyLow:
		return "VNGenerateOpticalFlowRequestComputationAccuracyLow"
	case VNGenerateOpticalFlowRequestComputationAccuracyMedium:
		return "VNGenerateOpticalFlowRequestComputationAccuracyMedium"
	case VNGenerateOpticalFlowRequestComputationAccuracyVeryHigh:
		return "VNGenerateOpticalFlowRequestComputationAccuracyVeryHigh"
	default:
		return fmt.Sprintf("VNGenerateOpticalFlowRequestComputationAccuracy(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest/QualityLevel-swift.enum
type VNGeneratePersonSegmentationRequestQualityLevel int

const (
	// VNGeneratePersonSegmentationRequestQualityLevelAccurate: Prefers image quality over performance.
	VNGeneratePersonSegmentationRequestQualityLevelAccurate VNGeneratePersonSegmentationRequestQualityLevel = 0
	// VNGeneratePersonSegmentationRequestQualityLevelBalanced: Prefers processing that balances image quality and performance.
	VNGeneratePersonSegmentationRequestQualityLevelBalanced VNGeneratePersonSegmentationRequestQualityLevel = 1
	// VNGeneratePersonSegmentationRequestQualityLevelFast: Prefers performance over image quality.
	VNGeneratePersonSegmentationRequestQualityLevelFast VNGeneratePersonSegmentationRequestQualityLevel = 2
)


func (e VNGeneratePersonSegmentationRequestQualityLevel) String() string {
	switch e {
	case VNGeneratePersonSegmentationRequestQualityLevelAccurate:
		return "VNGeneratePersonSegmentationRequestQualityLevelAccurate"
	case VNGeneratePersonSegmentationRequestQualityLevelBalanced:
		return "VNGeneratePersonSegmentationRequestQualityLevelBalanced"
	case VNGeneratePersonSegmentationRequestQualityLevelFast:
		return "VNGeneratePersonSegmentationRequestQualityLevelFast"
	default:
		return fmt.Sprintf("VNGeneratePersonSegmentationRequestQualityLevel(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/HeightEstimation-swift.enum
type VNHumanBodyPose3DObservationHeightEstimation int

const (
	// VNHumanBodyPose3DObservationHeightEstimationMeasured: A technique that uses LiDAR depth data to measure body height, in meters.
	VNHumanBodyPose3DObservationHeightEstimationMeasured VNHumanBodyPose3DObservationHeightEstimation = 1
	// VNHumanBodyPose3DObservationHeightEstimationReference: A technique that uses a reference height.
	VNHumanBodyPose3DObservationHeightEstimationReference VNHumanBodyPose3DObservationHeightEstimation = 0
)


func (e VNHumanBodyPose3DObservationHeightEstimation) String() string {
	switch e {
	case VNHumanBodyPose3DObservationHeightEstimationMeasured:
		return "VNHumanBodyPose3DObservationHeightEstimationMeasured"
	case VNHumanBodyPose3DObservationHeightEstimationReference:
		return "VNHumanBodyPose3DObservationHeightEstimationReference"
	default:
		return fmt.Sprintf("VNHumanBodyPose3DObservationHeightEstimation(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption
type VNImageCropAndScaleOption int

const (
	// VNImageCropAndScaleOptionCenterCrop: An option that scales the image to fit its shorter side within the input dimensions, while preserving its aspect ratio, and center-crops the image.
	VNImageCropAndScaleOptionCenterCrop VNImageCropAndScaleOption = 0
	// VNImageCropAndScaleOptionScaleFill: An option that scales the image to fill the input dimensions, resizing it if necessary.
	VNImageCropAndScaleOptionScaleFill VNImageCropAndScaleOption = 2
	// VNImageCropAndScaleOptionScaleFillRotate90CCW: An option that rotates the image 90 degrees counterclockwise and then scales it to fill the input dimensions.
	VNImageCropAndScaleOptionScaleFillRotate90CCW VNImageCropAndScaleOption = 256
	// VNImageCropAndScaleOptionScaleFit: An option that scales the image to fit its longer side within the input dimensions, while preserving its aspect ratio, and center-crops the image.
	VNImageCropAndScaleOptionScaleFit VNImageCropAndScaleOption = 1
	// VNImageCropAndScaleOptionScaleFitRotate90CCW: An option that rotates the image 90 degrees counterclockwise and then scales it, while preserving its aspect ratio, to fit on the long side.
	VNImageCropAndScaleOptionScaleFitRotate90CCW VNImageCropAndScaleOption = 256
)


func (e VNImageCropAndScaleOption) String() string {
	switch e {
	case VNImageCropAndScaleOptionCenterCrop:
		return "VNImageCropAndScaleOptionCenterCrop"
	case VNImageCropAndScaleOptionScaleFill:
		return "VNImageCropAndScaleOptionScaleFill"
	case VNImageCropAndScaleOptionScaleFillRotate90CCW:
		return "VNImageCropAndScaleOptionScaleFillRotate90CCW"
	case VNImageCropAndScaleOptionScaleFit:
		return "VNImageCropAndScaleOptionScaleFit"
	default:
		return fmt.Sprintf("VNImageCropAndScaleOption(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/Vision/VNPointsClassification
type VNPointsClassification int

const (
	VNPointsClassificationClosedPath VNPointsClassification = 2
	VNPointsClassificationDisconnected VNPointsClassification = 0
	VNPointsClassificationOpenPath VNPointsClassification = 1
)


func (e VNPointsClassification) String() string {
	switch e {
	case VNPointsClassificationClosedPath:
		return "VNPointsClassificationClosedPath"
	case VNPointsClassificationDisconnected:
		return "VNPointsClassificationDisconnected"
	case VNPointsClassificationOpenPath:
		return "VNPointsClassificationOpenPath"
	default:
		return fmt.Sprintf("VNPointsClassification(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/Vision/VNRequestFaceLandmarksConstellation
type VNRequestFaceLandmarksConstellation int

const (
	// VNRequestFaceLandmarksConstellation65Points: A constellation with 65 points.
	VNRequestFaceLandmarksConstellation65Points VNRequestFaceLandmarksConstellation = 1
	// VNRequestFaceLandmarksConstellation76Points: A constellation with 76 points.
	VNRequestFaceLandmarksConstellation76Points VNRequestFaceLandmarksConstellation = 2
	// VNRequestFaceLandmarksConstellationNotDefined: An undefined constellation.
	VNRequestFaceLandmarksConstellationNotDefined VNRequestFaceLandmarksConstellation = 0
)


func (e VNRequestFaceLandmarksConstellation) String() string {
	switch e {
	case VNRequestFaceLandmarksConstellation65Points:
		return "VNRequestFaceLandmarksConstellation65Points"
	case VNRequestFaceLandmarksConstellation76Points:
		return "VNRequestFaceLandmarksConstellation76Points"
	case VNRequestFaceLandmarksConstellationNotDefined:
		return "VNRequestFaceLandmarksConstellationNotDefined"
	default:
		return fmt.Sprintf("VNRequestFaceLandmarksConstellation(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/Vision/VNRequestTextRecognitionLevel
type VNRequestTextRecognitionLevel int

const (
	// VNRequestTextRecognitionLevelAccurate: Accurate text recognition takes more time to produce a more comprehensive result.
	VNRequestTextRecognitionLevelAccurate VNRequestTextRecognitionLevel = 0
	// VNRequestTextRecognitionLevelFast: Fast text recognition returns results more quickly at the expense of accuracy.
	VNRequestTextRecognitionLevelFast VNRequestTextRecognitionLevel = 1
)


func (e VNRequestTextRecognitionLevel) String() string {
	switch e {
	case VNRequestTextRecognitionLevelAccurate:
		return "VNRequestTextRecognitionLevelAccurate"
	case VNRequestTextRecognitionLevelFast:
		return "VNRequestTextRecognitionLevelFast"
	default:
		return fmt.Sprintf("VNRequestTextRecognitionLevel(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/Vision/VNRequestTrackingLevel
type VNRequestTrackingLevel int

const (
	// VNRequestTrackingLevelAccurate: Tracking level that favors location accuracy over speed.
	VNRequestTrackingLevelAccurate VNRequestTrackingLevel = 0
	// VNRequestTrackingLevelFast: Tracking level that favors speed over location accuracy.
	VNRequestTrackingLevelFast VNRequestTrackingLevel = 1
)


func (e VNRequestTrackingLevel) String() string {
	switch e {
	case VNRequestTrackingLevelAccurate:
		return "VNRequestTrackingLevelAccurate"
	case VNRequestTrackingLevelFast:
		return "VNRequestTrackingLevelFast"
	default:
		return fmt.Sprintf("VNRequestTrackingLevel(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/Vision/VNTrackOpticalFlowRequest/ComputationAccuracy-swift.enum
type VNTrackOpticalFlowRequestComputationAccuracy int

const (
	// VNTrackOpticalFlowRequestComputationAccuracyHigh: An option that indicates a high level of computational accuracy.
	VNTrackOpticalFlowRequestComputationAccuracyHigh VNTrackOpticalFlowRequestComputationAccuracy = 2
	// VNTrackOpticalFlowRequestComputationAccuracyLow: An option that indicates a low level of computational accuracy.
	VNTrackOpticalFlowRequestComputationAccuracyLow VNTrackOpticalFlowRequestComputationAccuracy = 0
	// VNTrackOpticalFlowRequestComputationAccuracyMedium: An option that indicates a moderate level of computational accuracy.
	VNTrackOpticalFlowRequestComputationAccuracyMedium VNTrackOpticalFlowRequestComputationAccuracy = 1
	// VNTrackOpticalFlowRequestComputationAccuracyVeryHigh: An option that indicates a very high level of computational accuracy.
	VNTrackOpticalFlowRequestComputationAccuracyVeryHigh VNTrackOpticalFlowRequestComputationAccuracy = 3
)


func (e VNTrackOpticalFlowRequestComputationAccuracy) String() string {
	switch e {
	case VNTrackOpticalFlowRequestComputationAccuracyHigh:
		return "VNTrackOpticalFlowRequestComputationAccuracyHigh"
	case VNTrackOpticalFlowRequestComputationAccuracyLow:
		return "VNTrackOpticalFlowRequestComputationAccuracyLow"
	case VNTrackOpticalFlowRequestComputationAccuracyMedium:
		return "VNTrackOpticalFlowRequestComputationAccuracyMedium"
	case VNTrackOpticalFlowRequestComputationAccuracyVeryHigh:
		return "VNTrackOpticalFlowRequestComputationAccuracyVeryHigh"
	default:
		return fmt.Sprintf("VNTrackOpticalFlowRequestComputationAccuracy(%d)", e)
	}
}





