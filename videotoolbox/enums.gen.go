// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"fmt"
)

type KVT int32

const (
	KVTAllocationFailedErr                        KVT = -12904
	KVTColorCorrectionImageRotationFailedErr      KVT = -12219
	KVTColorCorrectionPixelTransferFailedErr      KVT = -12212
	KVTColorSyncTransformConvertFailedErr         KVT = -12919
	KVTCouldNotCreateColorCorrectionDataErr       KVT = -12918
	KVTCouldNotCreateInstanceErr                  KVT = -12907
	KVTCouldNotFindExtensionErr                   KVT = -19510
	KVTCouldNotFindTemporalFilterErr              KVT = -12217
	KVTCouldNotFindVideoDecoderErr                KVT = -12906
	KVTCouldNotFindVideoEncoderErr                KVT = -12908
	KVTCouldNotOutputTaggedBufferGroupErr         KVT = -17699
	KVTExtensionConflictErr                       KVT = -19511
	KVTExtensionDisabledErr                       KVT = -17697
	KVTFormatDescriptionChangeNotSupportedErr     KVT = -12916
	KVTFrameSiloInvalidTimeRangeErr               KVT = -12216
	KVTFrameSiloInvalidTimeStampErr               KVT = -12215
	KVTImageRotationNotSupportedErr               KVT = -12914
	KVTInsufficientSourceColorDataErr             KVT = -12917
	KVTInvalidSessionErr                          KVT = -12903
	KVTLogTransferFunctionMismatchErr             KVT = 0
	KVTMultiPassStorageIdentifierMismatchErr      KVT = -12213
	KVTMultiPassStorageInvalidErr                 KVT = -12214
	KVTParameterErr                               KVT = -12902
	KVTPixelRotationNotSupportedErr               KVT = -12914
	KVTPixelTransferNotPermittedErr               KVT = -12218
	KVTPixelTransferNotSupportedErr               KVT = -12905
	KVTPropertyNotSupportedErr                    KVT = -12900
	KVTPropertyReadOnlyErr                        KVT = -12901
	KVTSessionMalfunctionErr                      KVT = -17691
	KVTVideoDecoderAuthorizationErr               KVT = -12210
	KVTVideoDecoderBadDataErr                     KVT = -12909
	KVTVideoDecoderCallbackMessagingErr           KVT = -17695
	KVTVideoDecoderMalfunctionErr                 KVT = -12911
	KVTVideoDecoderNeedsRosettaErr                KVT = -17692
	KVTVideoDecoderNotAvailableNowErr             KVT = -12913
	KVTVideoDecoderReferenceMissingErr            KVT = -17694
	KVTVideoDecoderRemovedErr                     KVT = -17690
	KVTVideoDecoderUnknownErr                     KVT = -17696
	KVTVideoDecoderUnsupportedDataFormatErr       KVT = -12910
	KVTVideoEncoderAuthorizationErr               KVT = -12211
	KVTVideoEncoderAutoWhiteBalanceNotLockedErr   KVT = -19512
	KVTVideoEncoderMVHEVCVideoLayerIDsMismatchErr KVT = -17698
	KVTVideoEncoderMalfunctionErr                 KVT = -12912
	KVTVideoEncoderNeedsRosettaErr                KVT = -17693
	KVTVideoEncoderNotAvailableNowErr             KVT = -12915
)

func (e KVT) String() string {
	switch e {
	case KVTAllocationFailedErr:
		return "KVTAllocationFailedErr"
	case KVTColorCorrectionImageRotationFailedErr:
		return "KVTColorCorrectionImageRotationFailedErr"
	case KVTColorCorrectionPixelTransferFailedErr:
		return "KVTColorCorrectionPixelTransferFailedErr"
	case KVTColorSyncTransformConvertFailedErr:
		return "KVTColorSyncTransformConvertFailedErr"
	case KVTCouldNotCreateColorCorrectionDataErr:
		return "KVTCouldNotCreateColorCorrectionDataErr"
	case KVTCouldNotCreateInstanceErr:
		return "KVTCouldNotCreateInstanceErr"
	case KVTCouldNotFindExtensionErr:
		return "KVTCouldNotFindExtensionErr"
	case KVTCouldNotFindTemporalFilterErr:
		return "KVTCouldNotFindTemporalFilterErr"
	case KVTCouldNotFindVideoDecoderErr:
		return "KVTCouldNotFindVideoDecoderErr"
	case KVTCouldNotFindVideoEncoderErr:
		return "KVTCouldNotFindVideoEncoderErr"
	case KVTCouldNotOutputTaggedBufferGroupErr:
		return "KVTCouldNotOutputTaggedBufferGroupErr"
	case KVTExtensionConflictErr:
		return "KVTExtensionConflictErr"
	case KVTExtensionDisabledErr:
		return "KVTExtensionDisabledErr"
	case KVTFormatDescriptionChangeNotSupportedErr:
		return "KVTFormatDescriptionChangeNotSupportedErr"
	case KVTFrameSiloInvalidTimeRangeErr:
		return "KVTFrameSiloInvalidTimeRangeErr"
	case KVTFrameSiloInvalidTimeStampErr:
		return "KVTFrameSiloInvalidTimeStampErr"
	case KVTImageRotationNotSupportedErr:
		return "KVTImageRotationNotSupportedErr"
	case KVTInsufficientSourceColorDataErr:
		return "KVTInsufficientSourceColorDataErr"
	case KVTInvalidSessionErr:
		return "KVTInvalidSessionErr"
	case KVTLogTransferFunctionMismatchErr:
		return "KVTLogTransferFunctionMismatchErr"
	case KVTMultiPassStorageIdentifierMismatchErr:
		return "KVTMultiPassStorageIdentifierMismatchErr"
	case KVTMultiPassStorageInvalidErr:
		return "KVTMultiPassStorageInvalidErr"
	case KVTParameterErr:
		return "KVTParameterErr"
	case KVTPixelTransferNotPermittedErr:
		return "KVTPixelTransferNotPermittedErr"
	case KVTPixelTransferNotSupportedErr:
		return "KVTPixelTransferNotSupportedErr"
	case KVTPropertyNotSupportedErr:
		return "KVTPropertyNotSupportedErr"
	case KVTPropertyReadOnlyErr:
		return "KVTPropertyReadOnlyErr"
	case KVTSessionMalfunctionErr:
		return "KVTSessionMalfunctionErr"
	case KVTVideoDecoderAuthorizationErr:
		return "KVTVideoDecoderAuthorizationErr"
	case KVTVideoDecoderBadDataErr:
		return "KVTVideoDecoderBadDataErr"
	case KVTVideoDecoderCallbackMessagingErr:
		return "KVTVideoDecoderCallbackMessagingErr"
	case KVTVideoDecoderMalfunctionErr:
		return "KVTVideoDecoderMalfunctionErr"
	case KVTVideoDecoderNeedsRosettaErr:
		return "KVTVideoDecoderNeedsRosettaErr"
	case KVTVideoDecoderNotAvailableNowErr:
		return "KVTVideoDecoderNotAvailableNowErr"
	case KVTVideoDecoderReferenceMissingErr:
		return "KVTVideoDecoderReferenceMissingErr"
	case KVTVideoDecoderRemovedErr:
		return "KVTVideoDecoderRemovedErr"
	case KVTVideoDecoderUnknownErr:
		return "KVTVideoDecoderUnknownErr"
	case KVTVideoDecoderUnsupportedDataFormatErr:
		return "KVTVideoDecoderUnsupportedDataFormatErr"
	case KVTVideoEncoderAuthorizationErr:
		return "KVTVideoEncoderAuthorizationErr"
	case KVTVideoEncoderAutoWhiteBalanceNotLockedErr:
		return "KVTVideoEncoderAutoWhiteBalanceNotLockedErr"
	case KVTVideoEncoderMVHEVCVideoLayerIDsMismatchErr:
		return "KVTVideoEncoderMVHEVCVideoLayerIDsMismatchErr"
	case KVTVideoEncoderMalfunctionErr:
		return "KVTVideoEncoderMalfunctionErr"
	case KVTVideoEncoderNeedsRosettaErr:
		return "KVTVideoEncoderNeedsRosettaErr"
	case KVTVideoEncoderNotAvailableNowErr:
		return "KVTVideoEncoderNotAvailableNowErr"
	default:
		return fmt.Sprintf("KVT(%d)", e)
	}
}

type KVTQPModulationLevel int32

const (
	KVTQPModulationLevel_Default KVTQPModulationLevel = -1
	KVTQPModulationLevel_Disable KVTQPModulationLevel = 0
)

func (e KVTQPModulationLevel) String() string {
	switch e {
	case KVTQPModulationLevel_Default:
		return "KVTQPModulationLevel_Default"
	case KVTQPModulationLevel_Disable:
		return "KVTQPModulationLevel_Disable"
	default:
		return fmt.Sprintf("KVTQPModulationLevel(%d)", e)
	}
}

const KVTUnlimitedFrameDelayCount int32 = -1

// See: https://developer.apple.com/documentation/VideoToolbox/VTCompressionSessionOptionFlags
type VTCompressionSessionOptionFlags uint32

const (
	// KVTCompressionSessionBeginFinalPass: A flag that indicates the last pass in a multi-pass compression session.
	KVTCompressionSessionBeginFinalPass VTCompressionSessionOptionFlags = 1
)

func (e VTCompressionSessionOptionFlags) String() string {
	switch e {
	case KVTCompressionSessionBeginFinalPass:
		return "KVTCompressionSessionBeginFinalPass"
	default:
		return fmt.Sprintf("VTCompressionSessionOptionFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/VideoToolbox/VTDecodeFrameFlags
type VTDecodeFrameFlags uint32

const (
	// KVTDecodeFrame_1xRealTimePlayback: A flag that provides a hint to the video decoder that it’s ok to use a low-power mode that can’t decode faster than realtime.
	KVTDecodeFrame_1xRealTimePlayback VTDecodeFrameFlags = 4
	// KVTDecodeFrame_DoNotOutputFrame: A flag that provides a hint to the decompression session and video decoder not to return a frame.
	KVTDecodeFrame_DoNotOutputFrame VTDecodeFrameFlags = 2
	// KVTDecodeFrame_EnableAsynchronousDecompression: A flag that indicates to enable asynchronous decompression.
	KVTDecodeFrame_EnableAsynchronousDecompression VTDecodeFrameFlags = 1
	// KVTDecodeFrame_EnableTemporalProcessing: A flag that indicates to enable temporal processing.
	KVTDecodeFrame_EnableTemporalProcessing VTDecodeFrameFlags = 8
)

func (e VTDecodeFrameFlags) String() string {
	switch e {
	case KVTDecodeFrame_1xRealTimePlayback:
		return "KVTDecodeFrame_1xRealTimePlayback"
	case KVTDecodeFrame_DoNotOutputFrame:
		return "KVTDecodeFrame_DoNotOutputFrame"
	case KVTDecodeFrame_EnableAsynchronousDecompression:
		return "KVTDecodeFrame_EnableAsynchronousDecompression"
	case KVTDecodeFrame_EnableTemporalProcessing:
		return "KVTDecodeFrame_EnableTemporalProcessing"
	default:
		return fmt.Sprintf("VTDecodeFrameFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/VideoToolbox/VTDecodeInfoFlags
type VTDecodeInfoFlags uint32

const (
	// KVTDecodeInfo_Asynchronous: A flag that indicates the decode operation ran asynchronously.
	KVTDecodeInfo_Asynchronous VTDecodeInfoFlags = 1
	// KVTDecodeInfo_FrameDropped: A flag that indicates the decode operation dropped a frame.
	KVTDecodeInfo_FrameDropped     VTDecodeInfoFlags = 2
	KVTDecodeInfo_FrameInterrupted VTDecodeInfoFlags = 16
	// KVTDecodeInfo_ImageBufferModifiable: A flag that indicates the image buffer is safe to modify.
	KVTDecodeInfo_ImageBufferModifiable VTDecodeInfoFlags = 4
	// KVTDecodeInfo_SkippedLeadingFrameDropped: A flag that indicates whether the decode process skips leading frames after dropping a synchronization frame.
	KVTDecodeInfo_SkippedLeadingFrameDropped VTDecodeInfoFlags = 8
)

func (e VTDecodeInfoFlags) String() string {
	switch e {
	case KVTDecodeInfo_Asynchronous:
		return "KVTDecodeInfo_Asynchronous"
	case KVTDecodeInfo_FrameDropped:
		return "KVTDecodeInfo_FrameDropped"
	case KVTDecodeInfo_FrameInterrupted:
		return "KVTDecodeInfo_FrameInterrupted"
	case KVTDecodeInfo_ImageBufferModifiable:
		return "KVTDecodeInfo_ImageBufferModifiable"
	case KVTDecodeInfo_SkippedLeadingFrameDropped:
		return "KVTDecodeInfo_SkippedLeadingFrameDropped"
	default:
		return fmt.Sprintf("VTDecodeInfoFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/VideoToolbox/VTEncodeInfoFlags
type VTEncodeInfoFlags uint32

const (
	// KVTEncodeInfo_Asynchronous: A flag that indicates that an encode operation ran asynchronously.
	KVTEncodeInfo_Asynchronous VTEncodeInfoFlags = 1
	// KVTEncodeInfo_FrameDropped: A flag that indicates that a frame dropped during encoding.
	KVTEncodeInfo_FrameDropped VTEncodeInfoFlags = 2
)

func (e VTEncodeInfoFlags) String() string {
	switch e {
	case KVTEncodeInfo_Asynchronous:
		return "KVTEncodeInfo_Asynchronous"
	case KVTEncodeInfo_FrameDropped:
		return "KVTEncodeInfo_FrameDropped"
	default:
		return fmt.Sprintf("VTEncodeInfoFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorError-swift.struct/Code
type VTFrameProcessorError int

const (
	// VTFrameProcessorAssetDownloadFailed: Returned if download of a required model asset for the processor failed
	VTFrameProcessorAssetDownloadFailed VTFrameProcessorError = -19743
	// VTFrameProcessorFatalError: A fatal error occurred during processing.
	VTFrameProcessorFatalError VTFrameProcessorError = -19734
	// VTFrameProcessorInitializationFailed: The session failed to initialize the processing pipeline.
	VTFrameProcessorInitializationFailed VTFrameProcessorError = -19736
	// VTFrameProcessorInvalidFrameTiming: A provided frame object has a presentation time stamp which isn’t supported by the processor.
	VTFrameProcessorInvalidFrameTiming VTFrameProcessorError = -19742
	// VTFrameProcessorInvalidParameterError: A provided parameter isn’t valid.
	VTFrameProcessorInvalidParameterError VTFrameProcessorError = -19741
	// VTFrameProcessorMemoryAllocationFailure: The session or processor is unable to allocate the required memory.
	VTFrameProcessorMemoryAllocationFailure VTFrameProcessorError = -19738
	// VTFrameProcessorProcessingError: The processor encountered an issue that prevents it from processing the provided frame.
	VTFrameProcessorProcessingError VTFrameProcessorError = -19740
	// VTFrameProcessorRevisionNotSupported: The specified revision isn’t supported by the configured processor.
	VTFrameProcessorRevisionNotSupported VTFrameProcessorError = -19739
	// VTFrameProcessorSessionAlreadyActive: An attempt is made to start a session that is already started.
	VTFrameProcessorSessionAlreadyActive VTFrameProcessorError = -19733
	// VTFrameProcessorSessionLevelError: The processing failed and current session should be stopped.
	VTFrameProcessorSessionLevelError VTFrameProcessorError = -19735
	// VTFrameProcessorSessionNotStarted: The session is used to process frames without being started.
	VTFrameProcessorSessionNotStarted VTFrameProcessorError = -19732
	// VTFrameProcessorUnknownError: The processor failed for an unknown reason.
	VTFrameProcessorUnknownError VTFrameProcessorError = -19730
	// VTFrameProcessorUnsupportedInput: One or more frames is in a format which isn’t supported by the processor.
	VTFrameProcessorUnsupportedInput VTFrameProcessorError = -19737
	// VTFrameProcessorUnsupportedResolution: The processor failed due to an unsupported resolution.
	VTFrameProcessorUnsupportedResolution VTFrameProcessorError = -19731
)

func (e VTFrameProcessorError) String() string {
	switch e {
	case VTFrameProcessorAssetDownloadFailed:
		return "VTFrameProcessorAssetDownloadFailed"
	case VTFrameProcessorFatalError:
		return "VTFrameProcessorFatalError"
	case VTFrameProcessorInitializationFailed:
		return "VTFrameProcessorInitializationFailed"
	case VTFrameProcessorInvalidFrameTiming:
		return "VTFrameProcessorInvalidFrameTiming"
	case VTFrameProcessorInvalidParameterError:
		return "VTFrameProcessorInvalidParameterError"
	case VTFrameProcessorMemoryAllocationFailure:
		return "VTFrameProcessorMemoryAllocationFailure"
	case VTFrameProcessorProcessingError:
		return "VTFrameProcessorProcessingError"
	case VTFrameProcessorRevisionNotSupported:
		return "VTFrameProcessorRevisionNotSupported"
	case VTFrameProcessorSessionAlreadyActive:
		return "VTFrameProcessorSessionAlreadyActive"
	case VTFrameProcessorSessionLevelError:
		return "VTFrameProcessorSessionLevelError"
	case VTFrameProcessorSessionNotStarted:
		return "VTFrameProcessorSessionNotStarted"
	case VTFrameProcessorUnknownError:
		return "VTFrameProcessorUnknownError"
	case VTFrameProcessorUnsupportedInput:
		return "VTFrameProcessorUnsupportedInput"
	case VTFrameProcessorUnsupportedResolution:
		return "VTFrameProcessorUnsupportedResolution"
	default:
		return fmt.Sprintf("VTFrameProcessorError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionConfiguration/QualityPrioritization-swift.enum
type VTFrameRateConversionConfigurationQualityPrioritization int

const (
	// VTFrameRateConversionConfigurationQualityPrioritizationNormal: A normal quality prioritization level.
	VTFrameRateConversionConfigurationQualityPrioritizationNormal VTFrameRateConversionConfigurationQualityPrioritization = 1
	// VTFrameRateConversionConfigurationQualityPrioritizationQuality: A quality prioritization level.
	VTFrameRateConversionConfigurationQualityPrioritizationQuality VTFrameRateConversionConfigurationQualityPrioritization = 2
)

func (e VTFrameRateConversionConfigurationQualityPrioritization) String() string {
	switch e {
	case VTFrameRateConversionConfigurationQualityPrioritizationNormal:
		return "VTFrameRateConversionConfigurationQualityPrioritizationNormal"
	case VTFrameRateConversionConfigurationQualityPrioritizationQuality:
		return "VTFrameRateConversionConfigurationQualityPrioritizationQuality"
	default:
		return fmt.Sprintf("VTFrameRateConversionConfigurationQualityPrioritization(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionConfiguration/Revision-swift.enum
type VTFrameRateConversionConfigurationRevision int

const (
	// VTFrameRateConversionConfigurationRevision1: An algorithm or implementation that represents the first revision.
	VTFrameRateConversionConfigurationRevision1 VTFrameRateConversionConfigurationRevision = 1
)

func (e VTFrameRateConversionConfigurationRevision) String() string {
	switch e {
	case VTFrameRateConversionConfigurationRevision1:
		return "VTFrameRateConversionConfigurationRevision1"
	default:
		return fmt.Sprintf("VTFrameRateConversionConfigurationRevision(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionParameters/SubmissionMode-swift.enum
type VTFrameRateConversionParametersSubmissionMode int

const (
	// VTFrameRateConversionParametersSubmissionModeRandom: A submission follow presentation time order with a jump or skip in a frame sequence.
	VTFrameRateConversionParametersSubmissionModeRandom VTFrameRateConversionParametersSubmissionMode = 1
	// VTFrameRateConversionParametersSubmissionModeSequential: A submission follow presentation time order without a jump or skip when compared to a previous submission.
	VTFrameRateConversionParametersSubmissionModeSequential VTFrameRateConversionParametersSubmissionMode = 2
	// VTFrameRateConversionParametersSubmissionModeSequentialReferencesUnchanged: You are submitting frames sequentially.
	VTFrameRateConversionParametersSubmissionModeSequentialReferencesUnchanged VTFrameRateConversionParametersSubmissionMode = 3
)

func (e VTFrameRateConversionParametersSubmissionMode) String() string {
	switch e {
	case VTFrameRateConversionParametersSubmissionModeRandom:
		return "VTFrameRateConversionParametersSubmissionModeRandom"
	case VTFrameRateConversionParametersSubmissionModeSequential:
		return "VTFrameRateConversionParametersSubmissionModeSequential"
	case VTFrameRateConversionParametersSubmissionModeSequentialReferencesUnchanged:
		return "VTFrameRateConversionParametersSubmissionModeSequentialReferencesUnchanged"
	default:
		return fmt.Sprintf("VTFrameRateConversionParametersSubmissionMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurConfiguration/QualityPrioritization-swift.enum
type VTMotionBlurConfigurationQualityPrioritization int

const (
	// VTMotionBlurConfigurationQualityPrioritizationNormal: A normal quality prioritization level.
	VTMotionBlurConfigurationQualityPrioritizationNormal VTMotionBlurConfigurationQualityPrioritization = 1
	// VTMotionBlurConfigurationQualityPrioritizationQuality: A quality prioritization level.
	VTMotionBlurConfigurationQualityPrioritizationQuality VTMotionBlurConfigurationQualityPrioritization = 2
)

func (e VTMotionBlurConfigurationQualityPrioritization) String() string {
	switch e {
	case VTMotionBlurConfigurationQualityPrioritizationNormal:
		return "VTMotionBlurConfigurationQualityPrioritizationNormal"
	case VTMotionBlurConfigurationQualityPrioritizationQuality:
		return "VTMotionBlurConfigurationQualityPrioritizationQuality"
	default:
		return fmt.Sprintf("VTMotionBlurConfigurationQualityPrioritization(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurConfiguration/Revision-swift.enum
type VTMotionBlurConfigurationRevision int

const (
	// VTMotionBlurConfigurationRevision1: An algorithm or implementation that represents the first revision.
	VTMotionBlurConfigurationRevision1 VTMotionBlurConfigurationRevision = 1
)

func (e VTMotionBlurConfigurationRevision) String() string {
	switch e {
	case VTMotionBlurConfigurationRevision1:
		return "VTMotionBlurConfigurationRevision1"
	default:
		return fmt.Sprintf("VTMotionBlurConfigurationRevision(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurParameters/SubmissionMode-swift.enum
type VTMotionBlurParametersSubmissionMode int

const (
	// VTMotionBlurParametersSubmissionModeRandom: A submission follow presentation time order with a jump or skip in a frame sequence.
	VTMotionBlurParametersSubmissionModeRandom VTMotionBlurParametersSubmissionMode = 1
	// VTMotionBlurParametersSubmissionModeSequential: A submission follow presentation time order without a jump or skip when compared to a previous submission.
	VTMotionBlurParametersSubmissionModeSequential VTMotionBlurParametersSubmissionMode = 2
)

func (e VTMotionBlurParametersSubmissionMode) String() string {
	switch e {
	case VTMotionBlurParametersSubmissionModeRandom:
		return "VTMotionBlurParametersSubmissionModeRandom"
	case VTMotionBlurParametersSubmissionModeSequential:
		return "VTMotionBlurParametersSubmissionModeSequential"
	default:
		return fmt.Sprintf("VTMotionBlurParametersSubmissionMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionEstimationFrameFlags
type VTMotionEstimationFrameFlags uint32

const (
	// KVTMotionEstimationFrameFlags_CurrentBufferWillBeNextReferenceBuffer: A hint to the motion-estimation session that you are going to reuse the `currentBuffer` as `referenceBuffer` in the next call to VTMotionEstimationSessionEstimateMotionVectors.
	KVTMotionEstimationFrameFlags_CurrentBufferWillBeNextReferenceBuffer VTMotionEstimationFrameFlags = 1
)

func (e VTMotionEstimationFrameFlags) String() string {
	switch e {
	case KVTMotionEstimationFrameFlags_CurrentBufferWillBeNextReferenceBuffer:
		return "KVTMotionEstimationFrameFlags_CurrentBufferWillBeNextReferenceBuffer"
	default:
		return fmt.Sprintf("VTMotionEstimationFrameFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionEstimationInfoFlags
type VTMotionEstimationInfoFlags uint32

const (
	KVTMotionEstimationInfoFlags_Reserved0 VTMotionEstimationInfoFlags = 1
)

func (e VTMotionEstimationInfoFlags) String() string {
	switch e {
	case KVTMotionEstimationInfoFlags_Reserved0:
		return "KVTMotionEstimationInfoFlags_Reserved0"
	default:
		return fmt.Sprintf("VTMotionEstimationInfoFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowConfiguration/QualityPrioritization-swift.enum
type VTOpticalFlowConfigurationQualityPrioritization int

const (
	// VTOpticalFlowConfigurationQualityPrioritizationNormal: A normal quality prioritization level.
	VTOpticalFlowConfigurationQualityPrioritizationNormal VTOpticalFlowConfigurationQualityPrioritization = 1
	// VTOpticalFlowConfigurationQualityPrioritizationQuality: A quality prioritization level.
	VTOpticalFlowConfigurationQualityPrioritizationQuality VTOpticalFlowConfigurationQualityPrioritization = 2
)

func (e VTOpticalFlowConfigurationQualityPrioritization) String() string {
	switch e {
	case VTOpticalFlowConfigurationQualityPrioritizationNormal:
		return "VTOpticalFlowConfigurationQualityPrioritizationNormal"
	case VTOpticalFlowConfigurationQualityPrioritizationQuality:
		return "VTOpticalFlowConfigurationQualityPrioritizationQuality"
	default:
		return fmt.Sprintf("VTOpticalFlowConfigurationQualityPrioritization(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowConfiguration/Revision-swift.enum
type VTOpticalFlowConfigurationRevision int

const (
	// VTOpticalFlowConfigurationRevision1: An algorithm or implementation that represents the first revision.
	VTOpticalFlowConfigurationRevision1 VTOpticalFlowConfigurationRevision = 1
)

func (e VTOpticalFlowConfigurationRevision) String() string {
	switch e {
	case VTOpticalFlowConfigurationRevision1:
		return "VTOpticalFlowConfigurationRevision1"
	default:
		return fmt.Sprintf("VTOpticalFlowConfigurationRevision(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowParameters/SubmissionMode-swift.enum
type VTOpticalFlowParametersSubmissionMode int

const (
	// VTOpticalFlowParametersSubmissionModeRandom: A submission follow presentation time order with a jump or skip in a frame sequence.
	VTOpticalFlowParametersSubmissionModeRandom VTOpticalFlowParametersSubmissionMode = 1
	// VTOpticalFlowParametersSubmissionModeSequential: A submission follow presentation time order without a jump or skip when compared to a previous submission.
	VTOpticalFlowParametersSubmissionModeSequential VTOpticalFlowParametersSubmissionMode = 2
)

func (e VTOpticalFlowParametersSubmissionMode) String() string {
	switch e {
	case VTOpticalFlowParametersSubmissionModeRandom:
		return "VTOpticalFlowParametersSubmissionModeRandom"
	case VTOpticalFlowParametersSubmissionModeSequential:
		return "VTOpticalFlowParametersSubmissionModeSequential"
	default:
		return fmt.Sprintf("VTOpticalFlowParametersSubmissionMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/InputType-swift.enum
type VTSuperResolutionScalerConfigurationInputType int

const (
	VTSuperResolutionScalerConfigurationInputTypeImage VTSuperResolutionScalerConfigurationInputType = 2
	VTSuperResolutionScalerConfigurationInputTypeVideo VTSuperResolutionScalerConfigurationInputType = 1
)

func (e VTSuperResolutionScalerConfigurationInputType) String() string {
	switch e {
	case VTSuperResolutionScalerConfigurationInputTypeImage:
		return "VTSuperResolutionScalerConfigurationInputTypeImage"
	case VTSuperResolutionScalerConfigurationInputTypeVideo:
		return "VTSuperResolutionScalerConfigurationInputTypeVideo"
	default:
		return fmt.Sprintf("VTSuperResolutionScalerConfigurationInputType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/ModelStatus
type VTSuperResolutionScalerConfigurationModelStatus int

const (
	VTSuperResolutionScalerConfigurationModelStatusDownloadRequired VTSuperResolutionScalerConfigurationModelStatus = 0
	VTSuperResolutionScalerConfigurationModelStatusDownloading      VTSuperResolutionScalerConfigurationModelStatus = 1
	VTSuperResolutionScalerConfigurationModelStatusReady            VTSuperResolutionScalerConfigurationModelStatus = 2
)

func (e VTSuperResolutionScalerConfigurationModelStatus) String() string {
	switch e {
	case VTSuperResolutionScalerConfigurationModelStatusDownloadRequired:
		return "VTSuperResolutionScalerConfigurationModelStatusDownloadRequired"
	case VTSuperResolutionScalerConfigurationModelStatusDownloading:
		return "VTSuperResolutionScalerConfigurationModelStatusDownloading"
	case VTSuperResolutionScalerConfigurationModelStatusReady:
		return "VTSuperResolutionScalerConfigurationModelStatusReady"
	default:
		return fmt.Sprintf("VTSuperResolutionScalerConfigurationModelStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/QualityPrioritization-swift.enum
type VTSuperResolutionScalerConfigurationQualityPrioritization int

const (
	VTSuperResolutionScalerConfigurationQualityPrioritizationNormal VTSuperResolutionScalerConfigurationQualityPrioritization = 1
)

func (e VTSuperResolutionScalerConfigurationQualityPrioritization) String() string {
	switch e {
	case VTSuperResolutionScalerConfigurationQualityPrioritizationNormal:
		return "VTSuperResolutionScalerConfigurationQualityPrioritizationNormal"
	default:
		return fmt.Sprintf("VTSuperResolutionScalerConfigurationQualityPrioritization(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/Revision-swift.enum
type VTSuperResolutionScalerConfigurationRevision int

const (
	VTSuperResolutionScalerConfigurationRevision1 VTSuperResolutionScalerConfigurationRevision = 1
)

func (e VTSuperResolutionScalerConfigurationRevision) String() string {
	switch e {
	case VTSuperResolutionScalerConfigurationRevision1:
		return "VTSuperResolutionScalerConfigurationRevision1"
	default:
		return fmt.Sprintf("VTSuperResolutionScalerConfigurationRevision(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerParameters/SubmissionMode-swift.enum
type VTSuperResolutionScalerParametersSubmissionMode int

const (
	VTSuperResolutionScalerParametersSubmissionModeRandom     VTSuperResolutionScalerParametersSubmissionMode = 1
	VTSuperResolutionScalerParametersSubmissionModeSequential VTSuperResolutionScalerParametersSubmissionMode = 2
)

func (e VTSuperResolutionScalerParametersSubmissionMode) String() string {
	switch e {
	case VTSuperResolutionScalerParametersSubmissionModeRandom:
		return "VTSuperResolutionScalerParametersSubmissionModeRandom"
	case VTSuperResolutionScalerParametersSubmissionModeSequential:
		return "VTSuperResolutionScalerParametersSubmissionModeSequential"
	default:
		return fmt.Sprintf("VTSuperResolutionScalerParametersSubmissionMode(%d)", e)
	}
}
