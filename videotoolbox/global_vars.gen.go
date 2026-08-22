// Code generated from Apple documentation. DO NOT EDIT.

package videotoolbox

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

var (
	// VTFrameProcessorErrorDomain is the error domain of Video Toolbox errors.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorErrorDomain
	VTFrameProcessorErrorDomain foundation.NSErrorDomain
)

var (
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTAlphaChannelMode_PremultipliedAlpha
	KVTAlphaChannelMode_PremultipliedAlpha string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTAlphaChannelMode_StraightAlpha
	KVTAlphaChannelMode_StraightAlpha string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCameraCalibrationExtrinsicOriginSource_StereoCameraSystemBaseline
	KVTCameraCalibrationExtrinsicOriginSource_StereoCameraSystemBaseline string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCameraCalibrationLensAlgorithmKind_ParametricLens
	KVTCameraCalibrationLensAlgorithmKind_ParametricLens string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCameraCalibrationLensDomain_Color
	KVTCameraCalibrationLensDomain_Color string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCameraCalibrationLensRole_Left
	KVTCameraCalibrationLensRole_Left string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCameraCalibrationLensRole_Mono
	KVTCameraCalibrationLensRole_Mono string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCameraCalibrationLensRole_Right
	KVTCameraCalibrationLensRole_Right string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPreset_Balanced
	KVTCompressionPreset_Balanced string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPreset_HighQuality
	KVTCompressionPreset_HighQuality string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPreset_HighSpeed
	KVTCompressionPreset_HighSpeed string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPreset_VideoConferencing
	KVTCompressionPreset_VideoConferencing string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyCameraCalibrationKey_ExtrinsicOrientationQuaternion
	KVTCompressionPropertyCameraCalibrationKey_ExtrinsicOrientationQuaternion string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyCameraCalibrationKey_ExtrinsicOriginSource
	KVTCompressionPropertyCameraCalibrationKey_ExtrinsicOriginSource string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyCameraCalibrationKey_IntrinsicMatrix
	KVTCompressionPropertyCameraCalibrationKey_IntrinsicMatrix string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyCameraCalibrationKey_IntrinsicMatrixProjectionOffset
	KVTCompressionPropertyCameraCalibrationKey_IntrinsicMatrixProjectionOffset string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyCameraCalibrationKey_IntrinsicMatrixReferenceDimensions
	KVTCompressionPropertyCameraCalibrationKey_IntrinsicMatrixReferenceDimensions string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyCameraCalibrationKey_LensAlgorithmKind
	KVTCompressionPropertyCameraCalibrationKey_LensAlgorithmKind string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyCameraCalibrationKey_LensDistortions
	KVTCompressionPropertyCameraCalibrationKey_LensDistortions string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyCameraCalibrationKey_LensDomain
	KVTCompressionPropertyCameraCalibrationKey_LensDomain string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyCameraCalibrationKey_LensFrameAdjustmentsPolynomialX
	KVTCompressionPropertyCameraCalibrationKey_LensFrameAdjustmentsPolynomialX string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyCameraCalibrationKey_LensFrameAdjustmentsPolynomialY
	KVTCompressionPropertyCameraCalibrationKey_LensFrameAdjustmentsPolynomialY string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyCameraCalibrationKey_LensIdentifier
	KVTCompressionPropertyCameraCalibrationKey_LensIdentifier string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyCameraCalibrationKey_LensRole
	KVTCompressionPropertyCameraCalibrationKey_LensRole string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyCameraCalibrationKey_RadialAngleLimit
	KVTCompressionPropertyCameraCalibrationKey_RadialAngleLimit string
	// KVTCompressionPropertyKey_AllowFrameReordering is a Boolean value that indicates whether frame reordering is enabled.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_AllowFrameReordering
	KVTCompressionPropertyKey_AllowFrameReordering string
	// KVTCompressionPropertyKey_AllowOpenGOP is enables Open GOP (Group Of Pictures) encoding.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_AllowOpenGOP
	KVTCompressionPropertyKey_AllowOpenGOP string
	// KVTCompressionPropertyKey_AllowTemporalCompression is a Boolean value indicating whether temporal compression is enabled.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_AllowTemporalCompression
	KVTCompressionPropertyKey_AllowTemporalCompression string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_AlphaChannelMode
	KVTCompressionPropertyKey_AlphaChannelMode string
	// KVTCompressionPropertyKey_AspectRatio16x9 is a Boolean value indicating whether the DV video stream should have the 16x9 flag set.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_AspectRatio16x9
	KVTCompressionPropertyKey_AspectRatio16x9 string
	// KVTCompressionPropertyKey_AverageBitRate is the long-term desired average bit rate in bits per second.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_AverageBitRate
	KVTCompressionPropertyKey_AverageBitRate string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_BaseLayerBitRateFraction
	KVTCompressionPropertyKey_BaseLayerBitRateFraction string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_BaseLayerFrameRate
	KVTCompressionPropertyKey_BaseLayerFrameRate string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_BaseLayerFrameRateFraction
	KVTCompressionPropertyKey_BaseLayerFrameRateFraction string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_CalculateMeanSquaredError
	KVTCompressionPropertyKey_CalculateMeanSquaredError string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_CameraCalibrationDataLensCollection
	KVTCompressionPropertyKey_CameraCalibrationDataLensCollection string
	// KVTCompressionPropertyKey_CleanAperture is the clean aperture for encoded frames.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_CleanAperture
	KVTCompressionPropertyKey_CleanAperture string
	// KVTCompressionPropertyKey_ColorPrimaries is the color primaries for compressed content.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_ColorPrimaries
	KVTCompressionPropertyKey_ColorPrimaries string
	// KVTCompressionPropertyKey_ConstantBitRate is requires that the encoder use a Constant Bit Rate algorithm.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_ConstantBitRate
	KVTCompressionPropertyKey_ConstantBitRate string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_ConstantQualityFactor
	KVTCompressionPropertyKey_ConstantQualityFactor string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_ContentLightLevelInfo
	KVTCompressionPropertyKey_ContentLightLevelInfo string
	// KVTCompressionPropertyKey_DataRateLimits is zero, one, or two hard limits on data rate.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_DataRateLimits
	KVTCompressionPropertyKey_DataRateLimits string
	// KVTCompressionPropertyKey_Depth is the pixel depth of the encoded video.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_Depth
	KVTCompressionPropertyKey_Depth string
	// KVTCompressionPropertyKey_EnableLTR is enables Long Term Reference (LTR) frames during encoding.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_EnableLTR
	KVTCompressionPropertyKey_EnableLTR string
	// KVTCompressionPropertyKey_EncoderID is specifies a particular video encoder by its ID string.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_EncoderID
	KVTCompressionPropertyKey_EncoderID string
	// KVTCompressionPropertyKey_EstimatedAverageBytesPerFrame is an estimate of the expected size in bytes of a single encoded frame based on the current configuration.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_EstimatedAverageBytesPerFrame
	KVTCompressionPropertyKey_EstimatedAverageBytesPerFrame string
	// KVTCompressionPropertyKey_ExpectedDuration is the expected total duration of the compression session, if known.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_ExpectedDuration
	KVTCompressionPropertyKey_ExpectedDuration string
	// KVTCompressionPropertyKey_ExpectedFrameRate is the expected frame rate, if known.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_ExpectedFrameRate
	KVTCompressionPropertyKey_ExpectedFrameRate string
	// KVTCompressionPropertyKey_FieldCount is the field count indicating whether the frames should be encoded progressive (1) or interlaced (2).
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_FieldCount
	KVTCompressionPropertyKey_FieldCount string
	// KVTCompressionPropertyKey_FieldDetail is field ordering for encoded interlaced frames.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_FieldDetail
	KVTCompressionPropertyKey_FieldDetail string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_GammaLevel
	KVTCompressionPropertyKey_GammaLevel string
	// KVTCompressionPropertyKey_H264EntropyMode is the entropy encoding mode for H.264 compression.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_H264EntropyMode
	KVTCompressionPropertyKey_H264EntropyMode string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_HDRMetadataInsertionMode
	KVTCompressionPropertyKey_HDRMetadataInsertionMode string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_HasLeftStereoEyeView
	KVTCompressionPropertyKey_HasLeftStereoEyeView string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_HasRightStereoEyeView
	KVTCompressionPropertyKey_HasRightStereoEyeView string
	// KVTCompressionPropertyKey_HeroEye is a value that indicates which eye is the primary eye when rendering in 2D.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_HeroEye
	KVTCompressionPropertyKey_HeroEye string
	// KVTCompressionPropertyKey_HorizontalDisparityAdjustment is a value that indicates a relative shift of the left and right images, which changes the zero parallax plane.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_HorizontalDisparityAdjustment
	KVTCompressionPropertyKey_HorizontalDisparityAdjustment string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_HorizontalFieldOfView
	KVTCompressionPropertyKey_HorizontalFieldOfView string
	// KVTCompressionPropertyKey_ICCProfile is the ICC profile for compressed content.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_ICCProfile
	KVTCompressionPropertyKey_ICCProfile string
	// KVTCompressionPropertyKey_MVHEVCLeftAndRightViewIDs is specifies which view identifier corresponds to the left eye and right eye.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_MVHEVCLeftAndRightViewIDs
	KVTCompressionPropertyKey_MVHEVCLeftAndRightViewIDs string
	// KVTCompressionPropertyKey_MVHEVCVideoLayerIDs is the identifiers of the video layers to encode in a multiview encoding operation.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_MVHEVCVideoLayerIDs
	KVTCompressionPropertyKey_MVHEVCVideoLayerIDs string
	// KVTCompressionPropertyKey_MVHEVCViewIDs is the identifiers of the views corresponding to the video layers in a multiview encoding operation.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_MVHEVCViewIDs
	KVTCompressionPropertyKey_MVHEVCViewIDs string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_MasteringDisplayColorVolume
	KVTCompressionPropertyKey_MasteringDisplayColorVolume string
	// KVTCompressionPropertyKey_MaxAllowedFrameQP is the maximum allowed encoded frame QP (Quantization Parameter).
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_MaxAllowedFrameQP
	KVTCompressionPropertyKey_MaxAllowedFrameQP string
	// KVTCompressionPropertyKey_MaxFrameDelayCount is the maximum number of frames that a compressor is allowed to hold before it must output a compressed frame.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_MaxFrameDelayCount
	KVTCompressionPropertyKey_MaxFrameDelayCount string
	// KVTCompressionPropertyKey_MaxH264SliceBytes is the maximum slice size for H.264 encoding.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_MaxH264SliceBytes
	KVTCompressionPropertyKey_MaxH264SliceBytes string
	// KVTCompressionPropertyKey_MaxKeyFrameInterval is the maximum interval between key frames, also known as the key frame rate.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_MaxKeyFrameInterval
	KVTCompressionPropertyKey_MaxKeyFrameInterval string
	// KVTCompressionPropertyKey_MaxKeyFrameIntervalDuration is the maximum duration from one key frame to the next in seconds.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_MaxKeyFrameIntervalDuration
	KVTCompressionPropertyKey_MaxKeyFrameIntervalDuration string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_MaximizePowerEfficiency
	KVTCompressionPropertyKey_MaximizePowerEfficiency string
	// KVTCompressionPropertyKey_MaximumRealTimeFrameRate is a value that specifies the maximum real time rate at which frames can be submitted to a compression session.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_MaximumRealTimeFrameRate
	KVTCompressionPropertyKey_MaximumRealTimeFrameRate string
	// KVTCompressionPropertyKey_MinAllowedFrameQP is the minimum allowed encoded frame QP (Quantization Parameter).
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_MinAllowedFrameQP
	KVTCompressionPropertyKey_MinAllowedFrameQP string
	// KVTCompressionPropertyKey_MoreFramesAfterEnd is a Boolean value indicating whether and how a compression session concatenates frames with other compressed frames to form a longer series.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_MoreFramesAfterEnd
	KVTCompressionPropertyKey_MoreFramesAfterEnd string
	// KVTCompressionPropertyKey_MoreFramesBeforeStart is a Boolean value that indicates whether and how a compression session concatenates frames with other compressed frames to form a longer series.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_MoreFramesBeforeStart
	KVTCompressionPropertyKey_MoreFramesBeforeStart string
	// KVTCompressionPropertyKey_MultiPassStorage is a property key that enables multipass compression and provides storage for encoder private data.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_MultiPassStorage
	KVTCompressionPropertyKey_MultiPassStorage string
	// KVTCompressionPropertyKey_NumberOfPendingFrames is the number of pending frames in the compression session.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_NumberOfPendingFrames
	KVTCompressionPropertyKey_NumberOfPendingFrames string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_OutputBitDepth
	KVTCompressionPropertyKey_OutputBitDepth string
	// KVTCompressionPropertyKey_PixelAspectRatio is the pixel aspect ratio for encoded frames.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_PixelAspectRatio
	KVTCompressionPropertyKey_PixelAspectRatio string
	// KVTCompressionPropertyKey_PixelBufferPoolIsShared is a Boolean value indicating whether the common pixel buffer pool is shared between the video encoder and the session client.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_PixelBufferPoolIsShared
	KVTCompressionPropertyKey_PixelBufferPoolIsShared string
	// KVTCompressionPropertyKey_PixelTransferProperties is properties for configuring a pixel transfer session.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_PixelTransferProperties
	KVTCompressionPropertyKey_PixelTransferProperties string
	// KVTCompressionPropertyKey_PreserveAlphaChannel is a key that specifies whether to encode the alpha channel of input video frames.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_PreserveAlphaChannel
	KVTCompressionPropertyKey_PreserveAlphaChannel string
	// KVTCompressionPropertyKey_PreserveDynamicHDRMetadata is specifies whether to preserve dynamic HDR metadata on the input pixel buffer.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_PreserveDynamicHDRMetadata
	KVTCompressionPropertyKey_PreserveDynamicHDRMetadata string
	// KVTCompressionPropertyKey_PrioritizeEncodingSpeedOverQuality is a hint for the video encoder to maximize its speed during encoding, sacrificing quality if needed.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_PrioritizeEncodingSpeedOverQuality
	KVTCompressionPropertyKey_PrioritizeEncodingSpeedOverQuality string
	// KVTCompressionPropertyKey_ProfileLevel is the profile and level for the encoded bitstream.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_ProfileLevel
	KVTCompressionPropertyKey_ProfileLevel string
	// KVTCompressionPropertyKey_ProgressiveScan is a Boolean value indicating whether the DV video stream should have the progressive flag set.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_ProgressiveScan
	KVTCompressionPropertyKey_ProgressiveScan string
	// KVTCompressionPropertyKey_ProjectionKind is a value that indicates the projection kind.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_ProjectionKind
	KVTCompressionPropertyKey_ProjectionKind string
	// KVTCompressionPropertyKey_Quality is the desired compression quality.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_Quality
	KVTCompressionPropertyKey_Quality string
	// KVTCompressionPropertyKey_RealTime is a Boolean value indicating whether it’s recommended that the video encoder perform compression in real time.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_RealTime
	KVTCompressionPropertyKey_RealTime string
	// KVTCompressionPropertyKey_RecommendedParallelizationLimit is the recommended number of compression sessions to instantiate in a parallel encoding configuration.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_RecommendedParallelizationLimit
	KVTCompressionPropertyKey_RecommendedParallelizationLimit string
	// KVTCompressionPropertyKey_RecommendedParallelizedSubdivisionMinimumDuration is the recommended minimum duration for a given subdivision in a parallel encoding configuration.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_RecommendedParallelizedSubdivisionMinimumDuration
	KVTCompressionPropertyKey_RecommendedParallelizedSubdivisionMinimumDuration string
	// KVTCompressionPropertyKey_RecommendedParallelizedSubdivisionMinimumFrameCount is the recommended minimum number of video frames for a given subdivision in a parallel encoding configuration.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_RecommendedParallelizedSubdivisionMinimumFrameCount
	KVTCompressionPropertyKey_RecommendedParallelizedSubdivisionMinimumFrameCount string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_ReferenceBufferCount
	KVTCompressionPropertyKey_ReferenceBufferCount string
	// KVTCompressionPropertyKey_SourceFrameCount is the number of source frames, if known.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_SourceFrameCount
	KVTCompressionPropertyKey_SourceFrameCount string
	// KVTCompressionPropertyKey_SpatialAdaptiveQPLevel is a value that controls spatial adaptation of the quantization parameter (QP) based on per-frame statistics.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_SpatialAdaptiveQPLevel
	KVTCompressionPropertyKey_SpatialAdaptiveQPLevel string
	// KVTCompressionPropertyKey_StereoCameraBaseline is a value that specifies the distance between centers of the lenses of the camera system.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_StereoCameraBaseline
	KVTCompressionPropertyKey_StereoCameraBaseline string
	// KVTCompressionPropertyKey_SuggestedLookAheadFrameCount is a value that requests that the encoder retain the specified number of frames during encoding.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_SuggestedLookAheadFrameCount
	KVTCompressionPropertyKey_SuggestedLookAheadFrameCount string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_SupportedPresetDictionaries
	KVTCompressionPropertyKey_SupportedPresetDictionaries string
	// KVTCompressionPropertyKey_SupportsBaseFrameQP is a value that indicates whether the encoder supports base frame QP requests.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_SupportsBaseFrameQP
	KVTCompressionPropertyKey_SupportsBaseFrameQP string
	// KVTCompressionPropertyKey_TargetQualityForAlpha is the target quality to use for encoding the alpha channel.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_TargetQualityForAlpha
	KVTCompressionPropertyKey_TargetQualityForAlpha string
	// KVTCompressionPropertyKey_TransferFunction is the transfer function for compressed content.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_TransferFunction
	KVTCompressionPropertyKey_TransferFunction string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_UsingGPURegistryID
	KVTCompressionPropertyKey_UsingGPURegistryID string
	// KVTCompressionPropertyKey_UsingHardwareAcceleratedVideoEncoder is a Boolean value indicating whether a hardware-accelerated video encoder is used.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_UsingHardwareAcceleratedVideoEncoder
	KVTCompressionPropertyKey_UsingHardwareAcceleratedVideoEncoder string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_VBVBufferDuration
	KVTCompressionPropertyKey_VBVBufferDuration string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_VBVInitialDelayPercentage
	KVTCompressionPropertyKey_VBVInitialDelayPercentage string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_VBVMaxBitRate
	KVTCompressionPropertyKey_VBVMaxBitRate string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_VariableBitRate
	KVTCompressionPropertyKey_VariableBitRate string
	// KVTCompressionPropertyKey_VideoEncoderPixelBufferAttributes is the video encoder’s pixel buffer attributes for the compression session.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_VideoEncoderPixelBufferAttributes
	KVTCompressionPropertyKey_VideoEncoderPixelBufferAttributes string
	// KVTCompressionPropertyKey_ViewPackingKind is a value that indicates the view packing kind.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_ViewPackingKind
	KVTCompressionPropertyKey_ViewPackingKind string
	// KVTCompressionPropertyKey_YCbCrMatrix is the YCbCr matrix for compressed content.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTCompressionPropertyKey_YCbCrMatrix
	KVTCompressionPropertyKey_YCbCrMatrix string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecodeFrameOptionKey_ContentAnalyzerCropRectangle
	KVTDecodeFrameOptionKey_ContentAnalyzerCropRectangle string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecodeFrameOptionKey_ContentAnalyzerRotation
	KVTDecodeFrameOptionKey_ContentAnalyzerRotation string
	// KVTDecompressionPropertyKey_AllowBitstreamToChangeFrameDimensions is a Boolean value that indicates whether a decoder is allowed to output buffers matching reduced frame dimensions in the bitstream rather than under-filling them.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_AllowBitstreamToChangeFrameDimensions
	KVTDecompressionPropertyKey_AllowBitstreamToChangeFrameDimensions string
	// KVTDecompressionPropertyKey_ContentHasInterframeDependencies is an optional Boolean property indicating if the content being decoded has interframe dependencies, if the decoder knows.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_ContentHasInterframeDependencies
	KVTDecompressionPropertyKey_ContentHasInterframeDependencies string
	// KVTDecompressionPropertyKey_DecoderProducesRAWOutput is a value that indicates whether the decoder can produce RAW output requiring a RAW processing session for post-decode processing.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_DecoderProducesRAWOutput
	KVTDecompressionPropertyKey_DecoderProducesRAWOutput string
	// KVTDecompressionPropertyKey_DeinterlaceMode is modes for requesting a specific deinterlacing technique.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_DeinterlaceMode
	KVTDecompressionPropertyKey_DeinterlaceMode string
	// KVTDecompressionPropertyKey_FieldMode is modes for special handling of interlaced content (optional).
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_FieldMode
	KVTDecompressionPropertyKey_FieldMode string
	// KVTDecompressionPropertyKey_GeneratePerFrameHDRDisplayMetadata is a key that indicates to generate per frame HDR Metadata and attach it to the resulting decoded pixel buffers.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_GeneratePerFrameHDRDisplayMetadata
	KVTDecompressionPropertyKey_GeneratePerFrameHDRDisplayMetadata string
	// KVTDecompressionPropertyKey_MaxOutputPresentationTimeStampOfFramesBeingDecoded is the maximum output presentation timestamp of the frames currently being decoded.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_MaxOutputPresentationTimeStampOfFramesBeingDecoded
	KVTDecompressionPropertyKey_MaxOutputPresentationTimeStampOfFramesBeingDecoded string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_MaximizePowerEfficiency
	KVTDecompressionPropertyKey_MaximizePowerEfficiency string
	// KVTDecompressionPropertyKey_MinOutputPresentationTimeStampOfFramesBeingDecoded is the minimum output presentation timestamp of the frames currently being decoded.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_MinOutputPresentationTimeStampOfFramesBeingDecoded
	KVTDecompressionPropertyKey_MinOutputPresentationTimeStampOfFramesBeingDecoded string
	// KVTDecompressionPropertyKey_NumberOfFramesBeingDecoded is returns the number of frames currently being decoded.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_NumberOfFramesBeingDecoded
	KVTDecompressionPropertyKey_NumberOfFramesBeingDecoded string
	// KVTDecompressionPropertyKey_OnlyTheseFrames is requests that frames be filtered by type.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_OnlyTheseFrames
	KVTDecompressionPropertyKey_OnlyTheseFrames string
	// KVTDecompressionPropertyKey_OutputPoolRequestedMinimumBufferCount is the requested minimum buffer count that a decompression session should use for its output pixel buffer pool, without releasing buffers while the number in use is below this level.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_OutputPoolRequestedMinimumBufferCount
	KVTDecompressionPropertyKey_OutputPoolRequestedMinimumBufferCount string
	// KVTDecompressionPropertyKey_PixelBufferPool is a pixel buffer pool for pixel buffers being output by the decompression session.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_PixelBufferPool
	KVTDecompressionPropertyKey_PixelBufferPool string
	// KVTDecompressionPropertyKey_PixelBufferPoolIsShared is a Boolean value indicating whether a common pixel buffer pool is shared between the video decoder and the session client.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_PixelBufferPoolIsShared
	KVTDecompressionPropertyKey_PixelBufferPoolIsShared string
	// KVTDecompressionPropertyKey_PixelFormatsWithReducedResolutionSupport is pixel formats that support reduced-resolution decoding (optional).
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_PixelFormatsWithReducedResolutionSupport
	KVTDecompressionPropertyKey_PixelFormatsWithReducedResolutionSupport string
	// KVTDecompressionPropertyKey_PixelTransferProperties is specific pixel transfer features to be used during decompression.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_PixelTransferProperties
	KVTDecompressionPropertyKey_PixelTransferProperties string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_PropagatePerFrameHDRDisplayMetadata
	KVTDecompressionPropertyKey_PropagatePerFrameHDRDisplayMetadata string
	// KVTDecompressionPropertyKey_RealTime is a Boolean value indicating whether it’s recommended that the video decoder perform decompression in real time.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_RealTime
	KVTDecompressionPropertyKey_RealTime string
	// KVTDecompressionPropertyKey_ReducedCoefficientDecode is requests approximation during decoding.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_ReducedCoefficientDecode
	KVTDecompressionPropertyKey_ReducedCoefficientDecode string
	// KVTDecompressionPropertyKey_ReducedFrameDelivery is the proportion of frames that should be delivered, indicating that the rest may be dropped.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_ReducedFrameDelivery
	KVTDecompressionPropertyKey_ReducedFrameDelivery string
	// KVTDecompressionPropertyKey_ReducedResolutionDecode is request decoding at smaller resolutions than full-size (optional).
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_ReducedResolutionDecode
	KVTDecompressionPropertyKey_ReducedResolutionDecode string
	// KVTDecompressionPropertyKey_RequestRAWOutput is for decoders that produce RAW output, this property requests that the decompression session provides unprocessed output.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_RequestRAWOutput
	KVTDecompressionPropertyKey_RequestRAWOutput string
	// KVTDecompressionPropertyKey_RequestedMVHEVCVideoLayerIDs is requests multi-image decoding of specific MV-HEVC video layers.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_RequestedMVHEVCVideoLayerIDs
	KVTDecompressionPropertyKey_RequestedMVHEVCVideoLayerIDs string
	// KVTDecompressionPropertyKey_SuggestedQualityOfServiceTiers is an array of dictionaries that describe decreasing quality-of-service levels that clients can use to maintain realtime playback (optional).
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_SuggestedQualityOfServiceTiers
	KVTDecompressionPropertyKey_SuggestedQualityOfServiceTiers string
	// KVTDecompressionPropertyKey_SupportedPixelFormatsOrderedByPerformance is an array indicating speed tradeoffs between pixel formats (optional).
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_SupportedPixelFormatsOrderedByPerformance
	KVTDecompressionPropertyKey_SupportedPixelFormatsOrderedByPerformance string
	// KVTDecompressionPropertyKey_SupportedPixelFormatsOrderedByQuality is an array indicating quality levels among pixel formats.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_SupportedPixelFormatsOrderedByQuality
	KVTDecompressionPropertyKey_SupportedPixelFormatsOrderedByQuality string
	// KVTDecompressionPropertyKey_ThreadCount is the number of threads used by a codec or the suggested number of threads to use (optional).
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_ThreadCount
	KVTDecompressionPropertyKey_ThreadCount string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_UsingGPURegistryID
	KVTDecompressionPropertyKey_UsingGPURegistryID string
	// KVTDecompressionPropertyKey_UsingHardwareAcceleratedVideoDecoder is indicates if a hardware-accelerated video decoder is being used.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionPropertyKey_UsingHardwareAcceleratedVideoDecoder
	KVTDecompressionPropertyKey_UsingHardwareAcceleratedVideoDecoder string
	// KVTDecompressionProperty_DeinterlaceMode_Temporal is a temporal deinterlace mode.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionProperty_DeinterlaceMode_Temporal
	KVTDecompressionProperty_DeinterlaceMode_Temporal string
	// KVTDecompressionProperty_DeinterlaceMode_VerticalFilter is a vertical filter deinterlace mode.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionProperty_DeinterlaceMode_VerticalFilter
	KVTDecompressionProperty_DeinterlaceMode_VerticalFilter string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionProperty_FieldMode_BothFields
	KVTDecompressionProperty_FieldMode_BothFields string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionProperty_FieldMode_BottomFieldOnly
	KVTDecompressionProperty_FieldMode_BottomFieldOnly string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionProperty_FieldMode_DeinterlaceFields
	KVTDecompressionProperty_FieldMode_DeinterlaceFields string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionProperty_FieldMode_SingleField
	KVTDecompressionProperty_FieldMode_SingleField string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionProperty_FieldMode_TopFieldOnly
	KVTDecompressionProperty_FieldMode_TopFieldOnly string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionProperty_OnlyTheseFrames_AllFrames
	KVTDecompressionProperty_OnlyTheseFrames_AllFrames string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionProperty_OnlyTheseFrames_IFrames
	KVTDecompressionProperty_OnlyTheseFrames_IFrames string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionProperty_OnlyTheseFrames_KeyFrames
	KVTDecompressionProperty_OnlyTheseFrames_KeyFrames string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionProperty_OnlyTheseFrames_NonDroppableFrames
	KVTDecompressionProperty_OnlyTheseFrames_NonDroppableFrames string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionProperty_TemporalLevelLimit
	KVTDecompressionProperty_TemporalLevelLimit string
	// KVTDecompressionResolutionKey_Height is a key to specify the resolution height.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionResolutionKey_Height
	KVTDecompressionResolutionKey_Height string
	// KVTDecompressionResolutionKey_Width is a key to specify the resolution width.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDecompressionResolutionKey_Width
	KVTDecompressionResolutionKey_Width string
	// KVTDownsamplingMode_Average is average missing samples (default center).
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDownsamplingMode_Average
	KVTDownsamplingMode_Average string
	// KVTDownsamplingMode_Decimate is default, decimate extra samples.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTDownsamplingMode_Decimate
	KVTDownsamplingMode_Decimate string
	// KVTEncodeFrameOptionKey_AcknowledgedLTRTokens is enable Long Term Reference (LTR) frames during encoding.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTEncodeFrameOptionKey_AcknowledgedLTRTokens
	KVTEncodeFrameOptionKey_AcknowledgedLTRTokens string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTEncodeFrameOptionKey_BaseFrameQP
	KVTEncodeFrameOptionKey_BaseFrameQP string
	// KVTEncodeFrameOptionKey_ForceKeyFrame is boolean value indicating whether the current frame is forced to be a key frame.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTEncodeFrameOptionKey_ForceKeyFrame
	KVTEncodeFrameOptionKey_ForceKeyFrame string
	// KVTEncodeFrameOptionKey_ForceLTRRefresh is a Boolean value that indicates whether to force Long Term Reference (LTR).
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTEncodeFrameOptionKey_ForceLTRRefresh
	KVTEncodeFrameOptionKey_ForceLTRRefresh string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTH264EntropyMode_CABAC
	KVTH264EntropyMode_CABAC string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTH264EntropyMode_CAVLC
	KVTH264EntropyMode_CAVLC string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTHDRMetadataInsertionMode_Auto
	KVTHDRMetadataInsertionMode_Auto string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTHDRMetadataInsertionMode_None
	KVTHDRMetadataInsertionMode_None string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTHDRMetadataInsertionMode_RequestSDRRangePreservation
	KVTHDRMetadataInsertionMode_RequestSDRRangePreservation string
	// KVTHDRPerFrameMetadataGenerationOptionsKey_HDRFormats is specifies an array of HDR formats to generate.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTHDRPerFrameMetadataGenerationOptionsKey_HDRFormats
	KVTHDRPerFrameMetadataGenerationOptionsKey_HDRFormats string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTHeroEye_Left
	KVTHeroEye_Left string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTHeroEye_Right
	KVTHeroEye_Right string
	// KVTMotionEstimationSessionCreationOption_DetectTrueMotion is enable multi pass true motion detection.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTMotionEstimationSessionCreationOption_DetectTrueMotion
	KVTMotionEstimationSessionCreationOption_DetectTrueMotion string
	// KVTMotionEstimationSessionCreationOption_Label is a label you use to log and track resources.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTMotionEstimationSessionCreationOption_Label
	KVTMotionEstimationSessionCreationOption_Label string
	// KVTMotionEstimationSessionCreationOption_MotionVectorSize is the size of the search blocks that motion estimation session uses.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTMotionEstimationSessionCreationOption_MotionVectorSize
	KVTMotionEstimationSessionCreationOption_MotionVectorSize string
	// KVTMotionEstimationSessionCreationOption_UseMultiPassSearch is an option to use for higher quality motion estimation.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTMotionEstimationSessionCreationOption_UseMultiPassSearch
	KVTMotionEstimationSessionCreationOption_UseMultiPassSearch string
	// KVTMultiPassStorageCreationOption_DoNotDelete is indicates that the multipass storage object’s backing store should not be deleted when finalized.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTMultiPassStorageCreationOption_DoNotDelete
	KVTMultiPassStorageCreationOption_DoNotDelete string
	// KVTPixelRotationPropertyKey_FlipHorizontalOrientation is a key that specifies a horizontal flip.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPixelRotationPropertyKey_FlipHorizontalOrientation
	KVTPixelRotationPropertyKey_FlipHorizontalOrientation string
	// KVTPixelRotationPropertyKey_FlipVerticalOrientation is a key that specifies a vertical flip.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPixelRotationPropertyKey_FlipVerticalOrientation
	KVTPixelRotationPropertyKey_FlipVerticalOrientation string
	// KVTPixelRotationPropertyKey_Rotation is a key that specifies the amount of rotation in degrees.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPixelRotationPropertyKey_Rotation
	KVTPixelRotationPropertyKey_Rotation string
	// KVTPixelTransferPropertyKey_DestinationCleanAperture is the clean aperture for destination image buffers.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPixelTransferPropertyKey_DestinationCleanAperture
	KVTPixelTransferPropertyKey_DestinationCleanAperture string
	// KVTPixelTransferPropertyKey_DestinationColorPrimaries is the color primaries to be used for destination image buffers.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPixelTransferPropertyKey_DestinationColorPrimaries
	KVTPixelTransferPropertyKey_DestinationColorPrimaries string
	// KVTPixelTransferPropertyKey_DestinationICCProfile is the International Color Consortium (ICC) profile for destination image buffers.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPixelTransferPropertyKey_DestinationICCProfile
	KVTPixelTransferPropertyKey_DestinationICCProfile string
	// KVTPixelTransferPropertyKey_DestinationPixelAspectRatio is the pixel aspect ratio for destination image buffers.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPixelTransferPropertyKey_DestinationPixelAspectRatio
	KVTPixelTransferPropertyKey_DestinationPixelAspectRatio string
	// KVTPixelTransferPropertyKey_DestinationTransferFunction is the color transfer function to be used for destination image buffers.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPixelTransferPropertyKey_DestinationTransferFunction
	KVTPixelTransferPropertyKey_DestinationTransferFunction string
	// KVTPixelTransferPropertyKey_DestinationYCbCrMatrix is the color matrix to be used for YCbCr to RGB conversions involving the destination image buffers.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPixelTransferPropertyKey_DestinationYCbCrMatrix
	KVTPixelTransferPropertyKey_DestinationYCbCrMatrix string
	// KVTPixelTransferPropertyKey_DownsamplingMode is the specific chroma downsampling technique to be used.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPixelTransferPropertyKey_DownsamplingMode
	KVTPixelTransferPropertyKey_DownsamplingMode string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPixelTransferPropertyKey_RealTime
	KVTPixelTransferPropertyKey_RealTime string
	// KVTPixelTransferPropertyKey_ScalingMode is scaling mode for images during transfer between source and destination buffers.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPixelTransferPropertyKey_ScalingMode
	KVTPixelTransferPropertyKey_ScalingMode string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H263_Profile0_Level10
	KVTProfileLevel_H263_Profile0_Level10 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H263_Profile0_Level45
	KVTProfileLevel_H263_Profile0_Level45 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H263_Profile3_Level45
	KVTProfileLevel_H263_Profile3_Level45 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Baseline_1_3
	KVTProfileLevel_H264_Baseline_1_3 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Baseline_3_0
	KVTProfileLevel_H264_Baseline_3_0 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Baseline_3_1
	KVTProfileLevel_H264_Baseline_3_1 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Baseline_3_2
	KVTProfileLevel_H264_Baseline_3_2 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Baseline_4_0
	KVTProfileLevel_H264_Baseline_4_0 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Baseline_4_1
	KVTProfileLevel_H264_Baseline_4_1 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Baseline_4_2
	KVTProfileLevel_H264_Baseline_4_2 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Baseline_5_0
	KVTProfileLevel_H264_Baseline_5_0 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Baseline_5_1
	KVTProfileLevel_H264_Baseline_5_1 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Baseline_5_2
	KVTProfileLevel_H264_Baseline_5_2 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Baseline_AutoLevel
	KVTProfileLevel_H264_Baseline_AutoLevel string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_ConstrainedBaseline_AutoLevel
	KVTProfileLevel_H264_ConstrainedBaseline_AutoLevel string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_ConstrainedHigh_AutoLevel
	KVTProfileLevel_H264_ConstrainedHigh_AutoLevel string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Extended_5_0
	KVTProfileLevel_H264_Extended_5_0 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Extended_AutoLevel
	KVTProfileLevel_H264_Extended_AutoLevel string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_High_3_0
	KVTProfileLevel_H264_High_3_0 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_High_3_1
	KVTProfileLevel_H264_High_3_1 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_High_3_2
	KVTProfileLevel_H264_High_3_2 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_High_4_0
	KVTProfileLevel_H264_High_4_0 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_High_4_1
	KVTProfileLevel_H264_High_4_1 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_High_4_2
	KVTProfileLevel_H264_High_4_2 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_High_5_0
	KVTProfileLevel_H264_High_5_0 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_High_5_1
	KVTProfileLevel_H264_High_5_1 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_High_5_2
	KVTProfileLevel_H264_High_5_2 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_High_AutoLevel
	KVTProfileLevel_H264_High_AutoLevel string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Main_3_0
	KVTProfileLevel_H264_Main_3_0 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Main_3_1
	KVTProfileLevel_H264_Main_3_1 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Main_3_2
	KVTProfileLevel_H264_Main_3_2 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Main_4_0
	KVTProfileLevel_H264_Main_4_0 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Main_4_1
	KVTProfileLevel_H264_Main_4_1 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Main_4_2
	KVTProfileLevel_H264_Main_4_2 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Main_5_0
	KVTProfileLevel_H264_Main_5_0 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Main_5_1
	KVTProfileLevel_H264_Main_5_1 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Main_5_2
	KVTProfileLevel_H264_Main_5_2 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_H264_Main_AutoLevel
	KVTProfileLevel_H264_Main_AutoLevel string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_HEVC_Main10_AutoLevel
	KVTProfileLevel_HEVC_Main10_AutoLevel string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_HEVC_Main42210_AutoLevel
	KVTProfileLevel_HEVC_Main42210_AutoLevel string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_HEVC_Main_AutoLevel
	KVTProfileLevel_HEVC_Main_AutoLevel string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_HEVC_Monochrome10_AutoLevel
	KVTProfileLevel_HEVC_Monochrome10_AutoLevel string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_HEVC_Monochrome_AutoLevel
	KVTProfileLevel_HEVC_Monochrome_AutoLevel string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_MP4V_AdvancedSimple_L0
	KVTProfileLevel_MP4V_AdvancedSimple_L0 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_MP4V_AdvancedSimple_L1
	KVTProfileLevel_MP4V_AdvancedSimple_L1 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_MP4V_AdvancedSimple_L2
	KVTProfileLevel_MP4V_AdvancedSimple_L2 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_MP4V_AdvancedSimple_L3
	KVTProfileLevel_MP4V_AdvancedSimple_L3 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_MP4V_AdvancedSimple_L4
	KVTProfileLevel_MP4V_AdvancedSimple_L4 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_MP4V_Main_L2
	KVTProfileLevel_MP4V_Main_L2 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_MP4V_Main_L3
	KVTProfileLevel_MP4V_Main_L3 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_MP4V_Main_L4
	KVTProfileLevel_MP4V_Main_L4 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_MP4V_Simple_L0
	KVTProfileLevel_MP4V_Simple_L0 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_MP4V_Simple_L1
	KVTProfileLevel_MP4V_Simple_L1 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_MP4V_Simple_L2
	KVTProfileLevel_MP4V_Simple_L2 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProfileLevel_MP4V_Simple_L3
	KVTProfileLevel_MP4V_Simple_L3 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProjectionKind_Equirectangular
	KVTProjectionKind_Equirectangular string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProjectionKind_HalfEquirectangular
	KVTProjectionKind_HalfEquirectangular string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProjectionKind_ParametricImmersive
	KVTProjectionKind_ParametricImmersive string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTProjectionKind_Rectilinear
	KVTProjectionKind_Rectilinear string
	// KVTPropertyDocumentationKey is dictionary key to access any documentation intended for developers only.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPropertyDocumentationKey
	KVTPropertyDocumentationKey string
	// KVTPropertyReadWriteStatusKey is dictionary key to access the read/write status of a property.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPropertyReadWriteStatusKey
	KVTPropertyReadWriteStatusKey string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPropertyReadWriteStatus_ReadOnly
	KVTPropertyReadWriteStatus_ReadOnly string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPropertyReadWriteStatus_ReadWrite
	KVTPropertyReadWriteStatus_ReadWrite string
	// KVTPropertyShouldBeSerializedKey is dictionary key to access the serializable status of a property.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPropertyShouldBeSerializedKey
	KVTPropertyShouldBeSerializedKey string
	// KVTPropertySupportedValueListKey is dictionary key to access the array of of supported values.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPropertySupportedValueListKey
	KVTPropertySupportedValueListKey string
	// KVTPropertySupportedValueMaximumKey is dictionary key to access the maximum value of a property.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPropertySupportedValueMaximumKey
	KVTPropertySupportedValueMaximumKey string
	// KVTPropertySupportedValueMinimumKey is dictionary key to access the minimum value of a property.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPropertySupportedValueMinimumKey
	KVTPropertySupportedValueMinimumKey string
	// KVTPropertyTypeKey is dictionary key used to access the property type.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPropertyTypeKey
	KVTPropertyTypeKey string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPropertyType_Boolean
	KVTPropertyType_Boolean string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPropertyType_Enumeration
	KVTPropertyType_Enumeration string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTPropertyType_Number
	KVTPropertyType_Number string
	// KVTRAWProcessingParameterListElement_Description is the value corresponding to this key is a human-readable description for the element, suitable for displaying in a tooltip or other descriptive UI.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameterListElement_Description
	KVTRAWProcessingParameterListElement_Description string
	// KVTRAWProcessingParameterListElement_Label is the value corresponding to this key is a human-readable label for the element, suitable for displaying in a list of options.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameterListElement_Label
	KVTRAWProcessingParameterListElement_Label string
	// KVTRAWProcessingParameterListElement_ListElementID is the value corresponding to this key is a number indicating the index of a list element parameter in a list in the processing parameters dictionary.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameterListElement_ListElementID
	KVTRAWProcessingParameterListElement_ListElementID string
	// KVTRAWProcessingParameterValueType_Boolean is the value corresponding to this key is a string that indicates a Boolean parameter type.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameterValueType_Boolean
	KVTRAWProcessingParameterValueType_Boolean string
	// KVTRAWProcessingParameterValueType_Float is the value corresponding to this key is a string that indicates a floating-point parameter type.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameterValueType_Float
	KVTRAWProcessingParameterValueType_Float string
	// KVTRAWProcessingParameterValueType_Integer is the value corresponding to this key is a string that indicates an integer parameter type.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameterValueType_Integer
	KVTRAWProcessingParameterValueType_Integer string
	// KVTRAWProcessingParameterValueType_List is the value corresponding to this key is a string that indicates a list parameter type.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameterValueType_List
	KVTRAWProcessingParameterValueType_List string
	// KVTRAWProcessingParameterValueType_SubGroup is the value corresponding to this key is a string that indicates a subgroup parameter type.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameterValueType_SubGroup
	KVTRAWProcessingParameterValueType_SubGroup string
	// KVTRAWProcessingParameter_CameraValue is the value corresponding to this key is the “As Shot” value for this parameter as originally captured by the camera.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameter_CameraValue
	KVTRAWProcessingParameter_CameraValue string
	// KVTRAWProcessingParameter_CurrentValue is the value corresponding to this key is the currently configured value for this parameter.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameter_CurrentValue
	KVTRAWProcessingParameter_CurrentValue string
	// KVTRAWProcessingParameter_Description is the value corresponding to this key is a localized string with a description of the parameter suitable for display in a tooltip or other descriptive UI.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameter_Description
	KVTRAWProcessingParameter_Description string
	// KVTRAWProcessingParameter_Enabled is the value corresponding to this key is Boolean indicating whether the parameter is enabled and can be modified.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameter_Enabled
	KVTRAWProcessingParameter_Enabled string
	// KVTRAWProcessingParameter_InitialValue is the value corresponding to this key is the initial value for this parameter as defined by the container and metadata provided at creation time.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameter_InitialValue
	KVTRAWProcessingParameter_InitialValue string
	// KVTRAWProcessingParameter_Key is the value corresponding to this key must match the identifier for one of the permitted processing parameters.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameter_Key
	KVTRAWProcessingParameter_Key string
	// KVTRAWProcessingParameter_ListArray is the value corresponding to this key is an array of dictionaries describing each element in the list.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameter_ListArray
	KVTRAWProcessingParameter_ListArray string
	// KVTRAWProcessingParameter_MaximumValue is the value corresponding to this key is the maximum value allowed for this parameter.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameter_MaximumValue
	KVTRAWProcessingParameter_MaximumValue string
	// KVTRAWProcessingParameter_MinimumValue is the value corresponding to this key is the minimum value allowed for this parameter.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameter_MinimumValue
	KVTRAWProcessingParameter_MinimumValue string
	// KVTRAWProcessingParameter_Name is the value corresponding to this key is a localized string which can be displayed in UI representing this parameter.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameter_Name
	KVTRAWProcessingParameter_Name string
	// KVTRAWProcessingParameter_NeutralValue is the value corresponding to this key is a neutral setting for the processor.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameter_NeutralValue
	KVTRAWProcessingParameter_NeutralValue string
	// KVTRAWProcessingParameter_SubGroup is the value corresponding to this key is an array of dictionaries representing the individual sub-parameters in this group.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameter_SubGroup
	KVTRAWProcessingParameter_SubGroup string
	// KVTRAWProcessingParameter_ValueType is the value corresponding to this key is the type of the parameter.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingParameter_ValueType
	KVTRAWProcessingParameter_ValueType string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingPropertyKey_MetadataForSidecarFile
	KVTRAWProcessingPropertyKey_MetadataForSidecarFile string
	// KVTRAWProcessingPropertyKey_MetalDeviceRegistryID is this property requests that the Metal device corresponding to the specified registryID be used for any Metal related processing.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingPropertyKey_MetalDeviceRegistryID
	KVTRAWProcessingPropertyKey_MetalDeviceRegistryID string
	// KVTRAWProcessingPropertyKey_OutputColorAttachments is the color-related image buffer keys and values that will be attachments to the returned pixel buffers.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRAWProcessingPropertyKey_OutputColorAttachments
	KVTRAWProcessingPropertyKey_OutputColorAttachments string
	// KVTRotation_0 is a constant that indicates a rotation of 0 degrees.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRotation_0
	KVTRotation_0 string
	// KVTRotation_180 is a constant that indicates a rotation of 180 degrees.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRotation_180
	KVTRotation_180 string
	// KVTRotation_CCW90 is a constant that indicates a counterclockwise rotation of 90 degrees.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRotation_CCW90
	KVTRotation_CCW90 string
	// KVTRotation_CW90 is a constant that indicates a clockwise rotation of 90 degrees.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTRotation_CW90
	KVTRotation_CW90 string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTSampleAttachmentKey_QualityMetrics
	KVTSampleAttachmentKey_QualityMetrics string
	// KVTSampleAttachmentKey_RequireLTRAcknowledgementToken is a number value that contains a unique token for this Long Term Reference (LTR).
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTSampleAttachmentKey_RequireLTRAcknowledgementToken
	KVTSampleAttachmentKey_RequireLTRAcknowledgementToken string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTSampleAttachmentQualityMetricsKey_ChromaBlueMeanSquaredError
	KVTSampleAttachmentQualityMetricsKey_ChromaBlueMeanSquaredError string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTSampleAttachmentQualityMetricsKey_ChromaRedMeanSquaredError
	KVTSampleAttachmentQualityMetricsKey_ChromaRedMeanSquaredError string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTSampleAttachmentQualityMetricsKey_LumaMeanSquaredError
	KVTSampleAttachmentQualityMetricsKey_LumaMeanSquaredError string
	// KVTScalingMode_CropSourceToCleanAperture is the source image buffer’s clean aperture is scaled to the destination clean aperture.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTScalingMode_CropSourceToCleanAperture
	KVTScalingMode_CropSourceToCleanAperture string
	// KVTScalingMode_Letterbox is the source image buffer’s clean aperture is scaled to a rectangle fitted inside the destination clean aperture that preserves the source picture aspect ratio.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTScalingMode_Letterbox
	KVTScalingMode_Letterbox string
	// KVTScalingMode_Normal is the full width and height of the source image buffer is stretched to the full width and height of the destination image buffer.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTScalingMode_Normal
	KVTScalingMode_Normal string
	// KVTScalingMode_Trim is the source image buffer’s clean aperture is scaled to a rectangle that completely fills the destination clean aperture and preserves the source picture aspect ratio.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTScalingMode_Trim
	KVTScalingMode_Trim string
	// KVTVideoDecoderSpecification_EnableHardwareAcceleratedVideoDecoder is a Boolean value indicating whether VideoToolbox uses a hardware-accelerated video decoder, if available.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoDecoderSpecification_EnableHardwareAcceleratedVideoDecoder
	KVTVideoDecoderSpecification_EnableHardwareAcceleratedVideoDecoder string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoDecoderSpecification_PreferredDecoderGPURegistryID
	KVTVideoDecoderSpecification_PreferredDecoderGPURegistryID string
	// KVTVideoDecoderSpecification_RequireHardwareAcceleratedVideoDecoder is a Boolean value indicating whether to require hardware-accelerated decoding.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoDecoderSpecification_RequireHardwareAcceleratedVideoDecoder
	KVTVideoDecoderSpecification_RequireHardwareAcceleratedVideoDecoder string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoDecoderSpecification_RequiredDecoderGPURegistryID
	KVTVideoDecoderSpecification_RequiredDecoderGPURegistryID string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoEncoderListOption_IncludeStandardDefinitionDVEncoders
	KVTVideoEncoderListOption_IncludeStandardDefinitionDVEncoders string
	// KVTVideoEncoderList_CodecName is the encoder’s codec name key.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoEncoderList_CodecName
	KVTVideoEncoderList_CodecName string
	// KVTVideoEncoderList_CodecType is the encoder’s codec type key.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoEncoderList_CodecType
	KVTVideoEncoderList_CodecType string
	// KVTVideoEncoderList_DisplayName is the encoder’s display name key.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoEncoderList_DisplayName
	KVTVideoEncoderList_DisplayName string
	// KVTVideoEncoderList_EncoderID is a key that identifies the encoder ID.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoEncoderList_EncoderID
	KVTVideoEncoderList_EncoderID string
	// KVTVideoEncoderList_EncoderName is a key for the encoder’s name.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoEncoderList_EncoderName
	KVTVideoEncoderList_EncoderName string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoEncoderList_GPURegistryID
	KVTVideoEncoderList_GPURegistryID string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoEncoderList_InstanceLimit
	KVTVideoEncoderList_InstanceLimit string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoEncoderList_IsHardwareAccelerated
	KVTVideoEncoderList_IsHardwareAccelerated string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoEncoderList_PerformanceRating
	KVTVideoEncoderList_PerformanceRating string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoEncoderList_QualityRating
	KVTVideoEncoderList_QualityRating string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoEncoderList_SupportedSelectionProperties
	KVTVideoEncoderList_SupportedSelectionProperties string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoEncoderList_SupportsFrameReordering
	KVTVideoEncoderList_SupportsFrameReordering string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoEncoderList_SupportsMultiPass
	KVTVideoEncoderList_SupportsMultiPass string
	// KVTVideoEncoderSpecification_EnableHardwareAcceleratedVideoEncoder is a Boolean value indicating whether hardware-accelerated video encoding is allowed, if available.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoEncoderSpecification_EnableHardwareAcceleratedVideoEncoder
	KVTVideoEncoderSpecification_EnableHardwareAcceleratedVideoEncoder string
	// KVTVideoEncoderSpecification_EnableLowLatencyRateControl is specifies to select an encoder that supports low-latency operation and enables low-latency mode.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoEncoderSpecification_EnableLowLatencyRateControl
	KVTVideoEncoderSpecification_EnableLowLatencyRateControl string
	// KVTVideoEncoderSpecification_EncoderID is a key that indicates a particular video encoder to use.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoEncoderSpecification_EncoderID
	KVTVideoEncoderSpecification_EncoderID string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoEncoderSpecification_PreferredEncoderGPURegistryID
	KVTVideoEncoderSpecification_PreferredEncoderGPURegistryID string
	// KVTVideoEncoderSpecification_RequireHardwareAcceleratedVideoEncoder is a Boolean value indicating whether hardware-accelerated encoding is required.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoEncoderSpecification_RequireHardwareAcceleratedVideoEncoder
	KVTVideoEncoderSpecification_RequireHardwareAcceleratedVideoEncoder string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTVideoEncoderSpecification_RequiredEncoderGPURegistryID
	KVTVideoEncoderSpecification_RequiredEncoderGPURegistryID string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTViewPackingKind_OverUnder
	KVTViewPackingKind_OverUnder string
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTViewPackingKind_SideBySide
	KVTViewPackingKind_SideBySide string
)

var (
	// KVTExtensionProperties_CodecNameKey is a dictionary key for the user readable name string of the codec.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTExtensionProperties_CodecNameKey
	KVTExtensionProperties_CodecNameKey VTExtensionPropertiesKey
	// KVTExtensionProperties_ContainingBundleNameKey is a dictionary key for the extension host application localized name.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/VTExtensionPropertiesKey/containingBundleName
	KVTExtensionProperties_ContainingBundleNameKey VTExtensionPropertiesKey
	// KVTExtensionProperties_ContainingBundleURLKey is a dictionary key for the URL of the extension host application.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/VTExtensionPropertiesKey/containingBundleURL
	KVTExtensionProperties_ContainingBundleURLKey VTExtensionPropertiesKey
	// KVTExtensionProperties_ExtensionIdentifierKey is a dictionary key for the video decoder extension identifier.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/VTExtensionPropertiesKey/extensionIdentifier
	KVTExtensionProperties_ExtensionIdentifierKey VTExtensionPropertiesKey
	// KVTExtensionProperties_ExtensionNameKey is a dictionary key for the localized extension name.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/VTExtensionPropertiesKey/extensionName
	KVTExtensionProperties_ExtensionNameKey VTExtensionPropertiesKey
	// KVTExtensionProperties_ExtensionURLKey is a dictionary key for the URL of the extension.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/VTExtensionPropertiesKey/extensionURL
	KVTExtensionProperties_ExtensionURLKey VTExtensionPropertiesKey
)

var (
	// KVTHDRPerFrameMetadataGenerationHDRFormatType_DolbyVision is specifies that DolbyVision data should be generated and attached for each pixel buffer.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/kVTHDRPerFrameMetadataGenerationHDRFormatType_DolbyVision
	KVTHDRPerFrameMetadataGenerationHDRFormatType_DolbyVision VTHDRPerFrameMetadataGenerationHDRFormatType
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VTFrameProcessorErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VTFrameProcessorErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTAlphaChannelMode_PremultipliedAlpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTAlphaChannelMode_PremultipliedAlpha = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTAlphaChannelMode_StraightAlpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTAlphaChannelMode_StraightAlpha = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCameraCalibrationExtrinsicOriginSource_StereoCameraSystemBaseline"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCameraCalibrationExtrinsicOriginSource_StereoCameraSystemBaseline = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCameraCalibrationLensAlgorithmKind_ParametricLens"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCameraCalibrationLensAlgorithmKind_ParametricLens = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCameraCalibrationLensDomain_Color"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCameraCalibrationLensDomain_Color = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCameraCalibrationLensRole_Left"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCameraCalibrationLensRole_Left = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCameraCalibrationLensRole_Mono"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCameraCalibrationLensRole_Mono = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCameraCalibrationLensRole_Right"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCameraCalibrationLensRole_Right = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPreset_Balanced"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPreset_Balanced = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPreset_HighQuality"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPreset_HighQuality = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPreset_HighSpeed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPreset_HighSpeed = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPreset_VideoConferencing"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPreset_VideoConferencing = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyCameraCalibrationKey_ExtrinsicOrientationQuaternion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyCameraCalibrationKey_ExtrinsicOrientationQuaternion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyCameraCalibrationKey_ExtrinsicOriginSource"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyCameraCalibrationKey_ExtrinsicOriginSource = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyCameraCalibrationKey_IntrinsicMatrix"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyCameraCalibrationKey_IntrinsicMatrix = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyCameraCalibrationKey_IntrinsicMatrixProjectionOffset"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyCameraCalibrationKey_IntrinsicMatrixProjectionOffset = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyCameraCalibrationKey_IntrinsicMatrixReferenceDimensions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyCameraCalibrationKey_IntrinsicMatrixReferenceDimensions = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyCameraCalibrationKey_LensAlgorithmKind"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyCameraCalibrationKey_LensAlgorithmKind = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyCameraCalibrationKey_LensDistortions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyCameraCalibrationKey_LensDistortions = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyCameraCalibrationKey_LensDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyCameraCalibrationKey_LensDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyCameraCalibrationKey_LensFrameAdjustmentsPolynomialX"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyCameraCalibrationKey_LensFrameAdjustmentsPolynomialX = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyCameraCalibrationKey_LensFrameAdjustmentsPolynomialY"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyCameraCalibrationKey_LensFrameAdjustmentsPolynomialY = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyCameraCalibrationKey_LensIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyCameraCalibrationKey_LensIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyCameraCalibrationKey_LensRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyCameraCalibrationKey_LensRole = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyCameraCalibrationKey_RadialAngleLimit"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyCameraCalibrationKey_RadialAngleLimit = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_AllowFrameReordering"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_AllowFrameReordering = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_AllowOpenGOP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_AllowOpenGOP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_AllowTemporalCompression"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_AllowTemporalCompression = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_AlphaChannelMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_AlphaChannelMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_AspectRatio16x9"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_AspectRatio16x9 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_AverageBitRate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_AverageBitRate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_BaseLayerBitRateFraction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_BaseLayerBitRateFraction = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_BaseLayerFrameRate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_BaseLayerFrameRate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_BaseLayerFrameRateFraction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_BaseLayerFrameRateFraction = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_CalculateMeanSquaredError"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_CalculateMeanSquaredError = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_CameraCalibrationDataLensCollection"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_CameraCalibrationDataLensCollection = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_CleanAperture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_CleanAperture = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_ColorPrimaries"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_ColorPrimaries = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_ConstantBitRate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_ConstantBitRate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_ConstantQualityFactor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_ConstantQualityFactor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_ContentLightLevelInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_ContentLightLevelInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_DataRateLimits"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_DataRateLimits = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_Depth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_Depth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_EnableLTR"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_EnableLTR = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_EncoderID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_EncoderID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_EstimatedAverageBytesPerFrame"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_EstimatedAverageBytesPerFrame = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_ExpectedDuration"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_ExpectedDuration = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_ExpectedFrameRate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_ExpectedFrameRate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_FieldCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_FieldCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_FieldDetail"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_FieldDetail = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_GammaLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_GammaLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_H264EntropyMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_H264EntropyMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_HDRMetadataInsertionMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_HDRMetadataInsertionMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_HasLeftStereoEyeView"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_HasLeftStereoEyeView = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_HasRightStereoEyeView"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_HasRightStereoEyeView = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_HeroEye"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_HeroEye = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_HorizontalDisparityAdjustment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_HorizontalDisparityAdjustment = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_HorizontalFieldOfView"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_HorizontalFieldOfView = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_ICCProfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_ICCProfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_MVHEVCLeftAndRightViewIDs"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_MVHEVCLeftAndRightViewIDs = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_MVHEVCVideoLayerIDs"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_MVHEVCVideoLayerIDs = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_MVHEVCViewIDs"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_MVHEVCViewIDs = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_MasteringDisplayColorVolume"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_MasteringDisplayColorVolume = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_MaxAllowedFrameQP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_MaxAllowedFrameQP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_MaxFrameDelayCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_MaxFrameDelayCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_MaxH264SliceBytes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_MaxH264SliceBytes = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_MaxKeyFrameInterval"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_MaxKeyFrameInterval = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_MaxKeyFrameIntervalDuration"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_MaxKeyFrameIntervalDuration = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_MaximizePowerEfficiency"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_MaximizePowerEfficiency = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_MaximumRealTimeFrameRate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_MaximumRealTimeFrameRate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_MinAllowedFrameQP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_MinAllowedFrameQP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_MoreFramesAfterEnd"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_MoreFramesAfterEnd = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_MoreFramesBeforeStart"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_MoreFramesBeforeStart = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_MultiPassStorage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_MultiPassStorage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_NumberOfPendingFrames"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_NumberOfPendingFrames = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_OutputBitDepth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_OutputBitDepth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_PixelAspectRatio"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_PixelAspectRatio = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_PixelBufferPoolIsShared"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_PixelBufferPoolIsShared = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_PixelTransferProperties"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_PixelTransferProperties = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_PreserveAlphaChannel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_PreserveAlphaChannel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_PreserveDynamicHDRMetadata"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_PreserveDynamicHDRMetadata = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_PrioritizeEncodingSpeedOverQuality"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_PrioritizeEncodingSpeedOverQuality = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_ProfileLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_ProfileLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_ProgressiveScan"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_ProgressiveScan = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_ProjectionKind"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_ProjectionKind = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_Quality"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_Quality = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_RealTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_RealTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_RecommendedParallelizationLimit"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_RecommendedParallelizationLimit = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_RecommendedParallelizedSubdivisionMinimumDuration"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_RecommendedParallelizedSubdivisionMinimumDuration = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_RecommendedParallelizedSubdivisionMinimumFrameCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_RecommendedParallelizedSubdivisionMinimumFrameCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_ReferenceBufferCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_ReferenceBufferCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_SourceFrameCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_SourceFrameCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_SpatialAdaptiveQPLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_SpatialAdaptiveQPLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_StereoCameraBaseline"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_StereoCameraBaseline = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_SuggestedLookAheadFrameCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_SuggestedLookAheadFrameCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_SupportedPresetDictionaries"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_SupportedPresetDictionaries = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_SupportsBaseFrameQP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_SupportsBaseFrameQP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_TargetQualityForAlpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_TargetQualityForAlpha = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_TransferFunction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_TransferFunction = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_UsingGPURegistryID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_UsingGPURegistryID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_UsingHardwareAcceleratedVideoEncoder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_UsingHardwareAcceleratedVideoEncoder = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_VBVBufferDuration"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_VBVBufferDuration = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_VBVInitialDelayPercentage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_VBVInitialDelayPercentage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_VBVMaxBitRate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_VBVMaxBitRate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_VariableBitRate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_VariableBitRate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_VideoEncoderPixelBufferAttributes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_VideoEncoderPixelBufferAttributes = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_ViewPackingKind"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_ViewPackingKind = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTCompressionPropertyKey_YCbCrMatrix"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTCompressionPropertyKey_YCbCrMatrix = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecodeFrameOptionKey_ContentAnalyzerCropRectangle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecodeFrameOptionKey_ContentAnalyzerCropRectangle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecodeFrameOptionKey_ContentAnalyzerRotation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecodeFrameOptionKey_ContentAnalyzerRotation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_AllowBitstreamToChangeFrameDimensions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_AllowBitstreamToChangeFrameDimensions = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_ContentHasInterframeDependencies"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_ContentHasInterframeDependencies = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_DecoderProducesRAWOutput"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_DecoderProducesRAWOutput = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_DeinterlaceMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_DeinterlaceMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_FieldMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_FieldMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_GeneratePerFrameHDRDisplayMetadata"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_GeneratePerFrameHDRDisplayMetadata = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_MaxOutputPresentationTimeStampOfFramesBeingDecoded"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_MaxOutputPresentationTimeStampOfFramesBeingDecoded = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_MaximizePowerEfficiency"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_MaximizePowerEfficiency = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_MinOutputPresentationTimeStampOfFramesBeingDecoded"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_MinOutputPresentationTimeStampOfFramesBeingDecoded = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_NumberOfFramesBeingDecoded"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_NumberOfFramesBeingDecoded = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_OnlyTheseFrames"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_OnlyTheseFrames = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_OutputPoolRequestedMinimumBufferCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_OutputPoolRequestedMinimumBufferCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_PixelBufferPool"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_PixelBufferPool = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_PixelBufferPoolIsShared"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_PixelBufferPoolIsShared = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_PixelFormatsWithReducedResolutionSupport"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_PixelFormatsWithReducedResolutionSupport = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_PixelTransferProperties"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_PixelTransferProperties = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_PropagatePerFrameHDRDisplayMetadata"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_PropagatePerFrameHDRDisplayMetadata = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_RealTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_RealTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_ReducedCoefficientDecode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_ReducedCoefficientDecode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_ReducedFrameDelivery"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_ReducedFrameDelivery = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_ReducedResolutionDecode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_ReducedResolutionDecode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_RequestRAWOutput"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_RequestRAWOutput = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_RequestedMVHEVCVideoLayerIDs"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_RequestedMVHEVCVideoLayerIDs = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_SuggestedQualityOfServiceTiers"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_SuggestedQualityOfServiceTiers = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_SupportedPixelFormatsOrderedByPerformance"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_SupportedPixelFormatsOrderedByPerformance = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_SupportedPixelFormatsOrderedByQuality"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_SupportedPixelFormatsOrderedByQuality = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_ThreadCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_ThreadCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_UsingGPURegistryID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_UsingGPURegistryID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionPropertyKey_UsingHardwareAcceleratedVideoDecoder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionPropertyKey_UsingHardwareAcceleratedVideoDecoder = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionProperty_DeinterlaceMode_Temporal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionProperty_DeinterlaceMode_Temporal = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionProperty_DeinterlaceMode_VerticalFilter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionProperty_DeinterlaceMode_VerticalFilter = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionProperty_FieldMode_BothFields"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionProperty_FieldMode_BothFields = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionProperty_FieldMode_BottomFieldOnly"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionProperty_FieldMode_BottomFieldOnly = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionProperty_FieldMode_DeinterlaceFields"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionProperty_FieldMode_DeinterlaceFields = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionProperty_FieldMode_SingleField"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionProperty_FieldMode_SingleField = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionProperty_FieldMode_TopFieldOnly"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionProperty_FieldMode_TopFieldOnly = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionProperty_OnlyTheseFrames_AllFrames"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionProperty_OnlyTheseFrames_AllFrames = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionProperty_OnlyTheseFrames_IFrames"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionProperty_OnlyTheseFrames_IFrames = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionProperty_OnlyTheseFrames_KeyFrames"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionProperty_OnlyTheseFrames_KeyFrames = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionProperty_OnlyTheseFrames_NonDroppableFrames"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionProperty_OnlyTheseFrames_NonDroppableFrames = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionProperty_TemporalLevelLimit"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionProperty_TemporalLevelLimit = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionResolutionKey_Height"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionResolutionKey_Height = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDecompressionResolutionKey_Width"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDecompressionResolutionKey_Width = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDownsamplingMode_Average"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDownsamplingMode_Average = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTDownsamplingMode_Decimate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTDownsamplingMode_Decimate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTEncodeFrameOptionKey_AcknowledgedLTRTokens"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTEncodeFrameOptionKey_AcknowledgedLTRTokens = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTEncodeFrameOptionKey_BaseFrameQP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTEncodeFrameOptionKey_BaseFrameQP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTEncodeFrameOptionKey_ForceKeyFrame"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTEncodeFrameOptionKey_ForceKeyFrame = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTEncodeFrameOptionKey_ForceLTRRefresh"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTEncodeFrameOptionKey_ForceLTRRefresh = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTExtensionProperties_CodecNameKey"); err == nil && ptr != 0 {
		KVTExtensionProperties_CodecNameKey = objc.ValueAt[VTExtensionPropertiesKey](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTExtensionProperties_ContainingBundleNameKey"); err == nil && ptr != 0 {
		KVTExtensionProperties_ContainingBundleNameKey = objc.ValueAt[VTExtensionPropertiesKey](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTExtensionProperties_ContainingBundleURLKey"); err == nil && ptr != 0 {
		KVTExtensionProperties_ContainingBundleURLKey = objc.ValueAt[VTExtensionPropertiesKey](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTExtensionProperties_ExtensionIdentifierKey"); err == nil && ptr != 0 {
		KVTExtensionProperties_ExtensionIdentifierKey = objc.ValueAt[VTExtensionPropertiesKey](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTExtensionProperties_ExtensionNameKey"); err == nil && ptr != 0 {
		KVTExtensionProperties_ExtensionNameKey = objc.ValueAt[VTExtensionPropertiesKey](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTExtensionProperties_ExtensionURLKey"); err == nil && ptr != 0 {
		KVTExtensionProperties_ExtensionURLKey = objc.ValueAt[VTExtensionPropertiesKey](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTH264EntropyMode_CABAC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTH264EntropyMode_CABAC = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTH264EntropyMode_CAVLC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTH264EntropyMode_CAVLC = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTHDRMetadataInsertionMode_Auto"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTHDRMetadataInsertionMode_Auto = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTHDRMetadataInsertionMode_None"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTHDRMetadataInsertionMode_None = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTHDRMetadataInsertionMode_RequestSDRRangePreservation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTHDRMetadataInsertionMode_RequestSDRRangePreservation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTHDRPerFrameMetadataGenerationHDRFormatType_DolbyVision"); err == nil && ptr != 0 {
		KVTHDRPerFrameMetadataGenerationHDRFormatType_DolbyVision = objc.ValueAt[VTHDRPerFrameMetadataGenerationHDRFormatType](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTHDRPerFrameMetadataGenerationOptionsKey_HDRFormats"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTHDRPerFrameMetadataGenerationOptionsKey_HDRFormats = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTHeroEye_Left"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTHeroEye_Left = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTHeroEye_Right"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTHeroEye_Right = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTMotionEstimationSessionCreationOption_DetectTrueMotion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTMotionEstimationSessionCreationOption_DetectTrueMotion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTMotionEstimationSessionCreationOption_Label"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTMotionEstimationSessionCreationOption_Label = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTMotionEstimationSessionCreationOption_MotionVectorSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTMotionEstimationSessionCreationOption_MotionVectorSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTMotionEstimationSessionCreationOption_UseMultiPassSearch"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTMotionEstimationSessionCreationOption_UseMultiPassSearch = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTMultiPassStorageCreationOption_DoNotDelete"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTMultiPassStorageCreationOption_DoNotDelete = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPixelRotationPropertyKey_FlipHorizontalOrientation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPixelRotationPropertyKey_FlipHorizontalOrientation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPixelRotationPropertyKey_FlipVerticalOrientation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPixelRotationPropertyKey_FlipVerticalOrientation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPixelRotationPropertyKey_Rotation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPixelRotationPropertyKey_Rotation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPixelTransferPropertyKey_DestinationCleanAperture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPixelTransferPropertyKey_DestinationCleanAperture = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPixelTransferPropertyKey_DestinationColorPrimaries"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPixelTransferPropertyKey_DestinationColorPrimaries = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPixelTransferPropertyKey_DestinationICCProfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPixelTransferPropertyKey_DestinationICCProfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPixelTransferPropertyKey_DestinationPixelAspectRatio"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPixelTransferPropertyKey_DestinationPixelAspectRatio = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPixelTransferPropertyKey_DestinationTransferFunction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPixelTransferPropertyKey_DestinationTransferFunction = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPixelTransferPropertyKey_DestinationYCbCrMatrix"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPixelTransferPropertyKey_DestinationYCbCrMatrix = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPixelTransferPropertyKey_DownsamplingMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPixelTransferPropertyKey_DownsamplingMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPixelTransferPropertyKey_RealTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPixelTransferPropertyKey_RealTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPixelTransferPropertyKey_ScalingMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPixelTransferPropertyKey_ScalingMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H263_Profile0_Level10"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H263_Profile0_Level10 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H263_Profile0_Level45"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H263_Profile0_Level45 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H263_Profile3_Level45"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H263_Profile3_Level45 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Baseline_1_3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Baseline_1_3 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Baseline_3_0"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Baseline_3_0 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Baseline_3_1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Baseline_3_1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Baseline_3_2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Baseline_3_2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Baseline_4_0"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Baseline_4_0 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Baseline_4_1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Baseline_4_1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Baseline_4_2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Baseline_4_2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Baseline_5_0"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Baseline_5_0 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Baseline_5_1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Baseline_5_1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Baseline_5_2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Baseline_5_2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Baseline_AutoLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Baseline_AutoLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_ConstrainedBaseline_AutoLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_ConstrainedBaseline_AutoLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_ConstrainedHigh_AutoLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_ConstrainedHigh_AutoLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Extended_5_0"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Extended_5_0 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Extended_AutoLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Extended_AutoLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_High_3_0"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_High_3_0 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_High_3_1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_High_3_1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_High_3_2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_High_3_2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_High_4_0"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_High_4_0 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_High_4_1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_High_4_1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_High_4_2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_High_4_2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_High_5_0"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_High_5_0 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_High_5_1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_High_5_1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_High_5_2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_High_5_2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_High_AutoLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_High_AutoLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Main_3_0"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Main_3_0 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Main_3_1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Main_3_1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Main_3_2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Main_3_2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Main_4_0"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Main_4_0 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Main_4_1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Main_4_1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Main_4_2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Main_4_2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Main_5_0"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Main_5_0 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Main_5_1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Main_5_1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Main_5_2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Main_5_2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_H264_Main_AutoLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_H264_Main_AutoLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_HEVC_Main10_AutoLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_HEVC_Main10_AutoLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_HEVC_Main42210_AutoLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_HEVC_Main42210_AutoLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_HEVC_Main_AutoLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_HEVC_Main_AutoLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_HEVC_Monochrome10_AutoLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_HEVC_Monochrome10_AutoLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_HEVC_Monochrome_AutoLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_HEVC_Monochrome_AutoLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_MP4V_AdvancedSimple_L0"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_MP4V_AdvancedSimple_L0 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_MP4V_AdvancedSimple_L1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_MP4V_AdvancedSimple_L1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_MP4V_AdvancedSimple_L2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_MP4V_AdvancedSimple_L2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_MP4V_AdvancedSimple_L3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_MP4V_AdvancedSimple_L3 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_MP4V_AdvancedSimple_L4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_MP4V_AdvancedSimple_L4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_MP4V_Main_L2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_MP4V_Main_L2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_MP4V_Main_L3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_MP4V_Main_L3 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_MP4V_Main_L4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_MP4V_Main_L4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_MP4V_Simple_L0"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_MP4V_Simple_L0 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_MP4V_Simple_L1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_MP4V_Simple_L1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_MP4V_Simple_L2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_MP4V_Simple_L2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProfileLevel_MP4V_Simple_L3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProfileLevel_MP4V_Simple_L3 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProjectionKind_Equirectangular"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProjectionKind_Equirectangular = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProjectionKind_HalfEquirectangular"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProjectionKind_HalfEquirectangular = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProjectionKind_ParametricImmersive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProjectionKind_ParametricImmersive = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTProjectionKind_Rectilinear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTProjectionKind_Rectilinear = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPropertyDocumentationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPropertyDocumentationKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPropertyReadWriteStatusKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPropertyReadWriteStatusKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPropertyReadWriteStatus_ReadOnly"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPropertyReadWriteStatus_ReadOnly = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPropertyReadWriteStatus_ReadWrite"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPropertyReadWriteStatus_ReadWrite = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPropertyShouldBeSerializedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPropertyShouldBeSerializedKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPropertySupportedValueListKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPropertySupportedValueListKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPropertySupportedValueMaximumKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPropertySupportedValueMaximumKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPropertySupportedValueMinimumKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPropertySupportedValueMinimumKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPropertyTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPropertyTypeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPropertyType_Boolean"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPropertyType_Boolean = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPropertyType_Enumeration"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPropertyType_Enumeration = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTPropertyType_Number"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTPropertyType_Number = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameterListElement_Description"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameterListElement_Description = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameterListElement_Label"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameterListElement_Label = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameterListElement_ListElementID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameterListElement_ListElementID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameterValueType_Boolean"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameterValueType_Boolean = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameterValueType_Float"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameterValueType_Float = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameterValueType_Integer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameterValueType_Integer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameterValueType_List"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameterValueType_List = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameterValueType_SubGroup"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameterValueType_SubGroup = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameter_CameraValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameter_CameraValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameter_CurrentValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameter_CurrentValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameter_Description"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameter_Description = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameter_Enabled"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameter_Enabled = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameter_InitialValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameter_InitialValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameter_Key"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameter_Key = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameter_ListArray"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameter_ListArray = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameter_MaximumValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameter_MaximumValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameter_MinimumValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameter_MinimumValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameter_Name"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameter_Name = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameter_NeutralValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameter_NeutralValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameter_SubGroup"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameter_SubGroup = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingParameter_ValueType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingParameter_ValueType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingPropertyKey_MetadataForSidecarFile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingPropertyKey_MetadataForSidecarFile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingPropertyKey_MetalDeviceRegistryID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingPropertyKey_MetalDeviceRegistryID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRAWProcessingPropertyKey_OutputColorAttachments"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRAWProcessingPropertyKey_OutputColorAttachments = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRotation_0"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRotation_0 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRotation_180"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRotation_180 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRotation_CCW90"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRotation_CCW90 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTRotation_CW90"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTRotation_CW90 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTSampleAttachmentKey_QualityMetrics"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTSampleAttachmentKey_QualityMetrics = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTSampleAttachmentKey_RequireLTRAcknowledgementToken"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTSampleAttachmentKey_RequireLTRAcknowledgementToken = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTSampleAttachmentQualityMetricsKey_ChromaBlueMeanSquaredError"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTSampleAttachmentQualityMetricsKey_ChromaBlueMeanSquaredError = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTSampleAttachmentQualityMetricsKey_ChromaRedMeanSquaredError"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTSampleAttachmentQualityMetricsKey_ChromaRedMeanSquaredError = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTSampleAttachmentQualityMetricsKey_LumaMeanSquaredError"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTSampleAttachmentQualityMetricsKey_LumaMeanSquaredError = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTScalingMode_CropSourceToCleanAperture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTScalingMode_CropSourceToCleanAperture = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTScalingMode_Letterbox"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTScalingMode_Letterbox = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTScalingMode_Normal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTScalingMode_Normal = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTScalingMode_Trim"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTScalingMode_Trim = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoDecoderSpecification_EnableHardwareAcceleratedVideoDecoder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoDecoderSpecification_EnableHardwareAcceleratedVideoDecoder = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoDecoderSpecification_PreferredDecoderGPURegistryID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoDecoderSpecification_PreferredDecoderGPURegistryID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoDecoderSpecification_RequireHardwareAcceleratedVideoDecoder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoDecoderSpecification_RequireHardwareAcceleratedVideoDecoder = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoDecoderSpecification_RequiredDecoderGPURegistryID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoDecoderSpecification_RequiredDecoderGPURegistryID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoEncoderListOption_IncludeStandardDefinitionDVEncoders"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoEncoderListOption_IncludeStandardDefinitionDVEncoders = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoEncoderList_CodecName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoEncoderList_CodecName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoEncoderList_CodecType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoEncoderList_CodecType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoEncoderList_DisplayName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoEncoderList_DisplayName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoEncoderList_EncoderID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoEncoderList_EncoderID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoEncoderList_EncoderName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoEncoderList_EncoderName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoEncoderList_GPURegistryID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoEncoderList_GPURegistryID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoEncoderList_InstanceLimit"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoEncoderList_InstanceLimit = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoEncoderList_IsHardwareAccelerated"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoEncoderList_IsHardwareAccelerated = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoEncoderList_PerformanceRating"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoEncoderList_PerformanceRating = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoEncoderList_QualityRating"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoEncoderList_QualityRating = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoEncoderList_SupportedSelectionProperties"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoEncoderList_SupportedSelectionProperties = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoEncoderList_SupportsFrameReordering"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoEncoderList_SupportsFrameReordering = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoEncoderList_SupportsMultiPass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoEncoderList_SupportsMultiPass = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoEncoderSpecification_EnableHardwareAcceleratedVideoEncoder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoEncoderSpecification_EnableHardwareAcceleratedVideoEncoder = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoEncoderSpecification_EnableLowLatencyRateControl"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoEncoderSpecification_EnableLowLatencyRateControl = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoEncoderSpecification_EncoderID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoEncoderSpecification_EncoderID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoEncoderSpecification_PreferredEncoderGPURegistryID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoEncoderSpecification_PreferredEncoderGPURegistryID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoEncoderSpecification_RequireHardwareAcceleratedVideoEncoder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoEncoderSpecification_RequireHardwareAcceleratedVideoEncoder = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTVideoEncoderSpecification_RequiredEncoderGPURegistryID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTVideoEncoderSpecification_RequiredEncoderGPURegistryID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTViewPackingKind_OverUnder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTViewPackingKind_OverUnder = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kVTViewPackingKind_SideBySide"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KVTViewPackingKind_SideBySide = objc.GoString(cstr)
			}
		}
	}

}
