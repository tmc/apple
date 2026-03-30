// Code generated from Apple documentation. DO NOT EDIT.

package coremedia

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionAlphaChannelMode_PremultipliedAlpha
	KCMFormatDescriptionAlphaChannelMode_PremultipliedAlpha string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionAlphaChannelMode_StraightAlpha
	KCMFormatDescriptionAlphaChannelMode_StraightAlpha string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionCameraCalibrationExtrinsicOriginSource_StereoCameraSystemBaseline
	KCMFormatDescriptionCameraCalibrationExtrinsicOriginSource_StereoCameraSystemBaseline string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionCameraCalibrationLensAlgorithmKind_ParametricLens
	KCMFormatDescriptionCameraCalibrationLensAlgorithmKind_ParametricLens string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionCameraCalibrationLensDomain_Color
	KCMFormatDescriptionCameraCalibrationLensDomain_Color string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionCameraCalibrationLensRole_Left
	KCMFormatDescriptionCameraCalibrationLensRole_Left string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionCameraCalibrationLensRole_Mono
	KCMFormatDescriptionCameraCalibrationLensRole_Mono string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionCameraCalibrationLensRole_Right
	KCMFormatDescriptionCameraCalibrationLensRole_Right string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionCameraCalibration_ExtrinsicOrientationQuaternion
	KCMFormatDescriptionCameraCalibration_ExtrinsicOrientationQuaternion string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionCameraCalibration_ExtrinsicOriginSource
	KCMFormatDescriptionCameraCalibration_ExtrinsicOriginSource string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionCameraCalibration_IntrinsicMatrix
	KCMFormatDescriptionCameraCalibration_IntrinsicMatrix string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionCameraCalibration_IntrinsicMatrixProjectionOffset
	KCMFormatDescriptionCameraCalibration_IntrinsicMatrixProjectionOffset string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionCameraCalibration_IntrinsicMatrixReferenceDimensions
	KCMFormatDescriptionCameraCalibration_IntrinsicMatrixReferenceDimensions string
	// KCMFormatDescriptionCameraCalibration_LensAlgorithmKind is the following keys are required in each kCMFormatDescriptionExtension_CameraCalibrationDataLensCollection dictionary.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionCameraCalibration_LensAlgorithmKind
	KCMFormatDescriptionCameraCalibration_LensAlgorithmKind string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionCameraCalibration_LensDistortions
	KCMFormatDescriptionCameraCalibration_LensDistortions string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionCameraCalibration_LensDomain
	KCMFormatDescriptionCameraCalibration_LensDomain string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionCameraCalibration_LensFrameAdjustmentsPolynomialX
	KCMFormatDescriptionCameraCalibration_LensFrameAdjustmentsPolynomialX string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionCameraCalibration_LensFrameAdjustmentsPolynomialY
	KCMFormatDescriptionCameraCalibration_LensFrameAdjustmentsPolynomialY string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionCameraCalibration_LensIdentifier
	KCMFormatDescriptionCameraCalibration_LensIdentifier string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionCameraCalibration_LensRole
	KCMFormatDescriptionCameraCalibration_LensRole string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionCameraCalibration_RadialAngleLimit
	KCMFormatDescriptionCameraCalibration_RadialAngleLimit string
	// KCMFormatDescriptionChromaLocation_Bottom is the chroma sample is horizontally centered, but co-sited with the bottom row of the luma samples.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionChromaLocation_Bottom-swift.var
	KCMFormatDescriptionChromaLocation_Bottom string
	// KCMFormatDescriptionChromaLocation_BottomLeft is the chroma sample is co-sited with the bottom-left luma sample.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionChromaLocation_BottomLeft-swift.var
	KCMFormatDescriptionChromaLocation_BottomLeft string
	// KCMFormatDescriptionChromaLocation_Center is the chroma sample is fully centered.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionChromaLocation_Center-swift.var
	KCMFormatDescriptionChromaLocation_Center string
	// KCMFormatDescriptionChromaLocation_DV420 is the Cr and Cb samples are alternately co-sited with the left luma samples of the same field.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionChromaLocation_DV420-swift.var
	KCMFormatDescriptionChromaLocation_DV420 string
	// KCMFormatDescriptionChromaLocation_Left is the chroma sample is horizontally co-sited with the left column of the luma samples, but centered vertically.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionChromaLocation_Left-swift.var
	KCMFormatDescriptionChromaLocation_Left string
	// KCMFormatDescriptionChromaLocation_Top is the chroma sample is horizontally centered, but co-sited with the top row of the luma samples.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionChromaLocation_Top-swift.var
	KCMFormatDescriptionChromaLocation_Top string
	// KCMFormatDescriptionChromaLocation_TopLeft is the chroma sample is co-sited with the top-left luma sample.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionChromaLocation_TopLeft-swift.var
	KCMFormatDescriptionChromaLocation_TopLeft string
	// KCMFormatDescriptionColorPrimaries_DCI_P3 is a format that describes the DCI P3 color primary.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionColorPrimaries_DCI_P3
	KCMFormatDescriptionColorPrimaries_DCI_P3 string
	// KCMFormatDescriptionColorPrimaries_EBU_3213 is a format that describes the EBU 3213 color primary.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionColorPrimaries_EBU_3213-swift.var
	KCMFormatDescriptionColorPrimaries_EBU_3213 string
	// KCMFormatDescriptionColorPrimaries_ITU_R_2020 is a format that describes the ITU R 2020 color primary.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionColorPrimaries_ITU_R_2020
	KCMFormatDescriptionColorPrimaries_ITU_R_2020 string
	// KCMFormatDescriptionColorPrimaries_ITU_R_709_2 is a format that describes the ITU R 709 2 color primary.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionColorPrimaries_ITU_R_709_2-swift.var
	KCMFormatDescriptionColorPrimaries_ITU_R_709_2 string
	// KCMFormatDescriptionColorPrimaries_P22 is a format that describes the P22 color primary.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionColorPrimaries_P22
	KCMFormatDescriptionColorPrimaries_P22 string
	// KCMFormatDescriptionColorPrimaries_P3_D65 is a format that describes the P3 D65 color primary.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionColorPrimaries_P3_D65
	KCMFormatDescriptionColorPrimaries_P3_D65 string
	// KCMFormatDescriptionColorPrimaries_SMPTE_C is a format that describes the SMPTE C color primary.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionColorPrimaries_SMPTE_C-swift.var
	KCMFormatDescriptionColorPrimaries_SMPTE_C string
	// KCMFormatDescriptionConformsToMPEG2VideoProfile is the value is a number that specifies the MPEG2 video profile.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionConformsToMPEG2VideoProfile
	KCMFormatDescriptionConformsToMPEG2VideoProfile string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtensionKey_MetadataKeyTable
	KCMFormatDescriptionExtensionKey_MetadataKeyTable string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_AlphaChannelMode
	KCMFormatDescriptionExtension_AlphaChannelMode string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_AlternativeTransferCharacteristics
	KCMFormatDescriptionExtension_AlternativeTransferCharacteristics string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_AmbientViewingEnvironment
	KCMFormatDescriptionExtension_AmbientViewingEnvironment string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_AuxiliaryTypeInfo
	KCMFormatDescriptionExtension_AuxiliaryTypeInfo string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_BitsPerComponent
	KCMFormatDescriptionExtension_BitsPerComponent string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_BytesPerRow
	KCMFormatDescriptionExtension_BytesPerRow string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_CameraCalibrationDataLensCollection
	KCMFormatDescriptionExtension_CameraCalibrationDataLensCollection string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_ChromaLocationBottomField-swift.var
	KCMFormatDescriptionExtension_ChromaLocationBottomField string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_ChromaLocationTopField-swift.var
	KCMFormatDescriptionExtension_ChromaLocationTopField string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_CleanAperture-swift.var
	KCMFormatDescriptionExtension_CleanAperture string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_ColorPrimaries-swift.var
	KCMFormatDescriptionExtension_ColorPrimaries string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_ContainsAlphaChannel
	KCMFormatDescriptionExtension_ContainsAlphaChannel string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_ContentColorVolume
	KCMFormatDescriptionExtension_ContentColorVolume string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_ContentLightLevelInfo
	KCMFormatDescriptionExtension_ContentLightLevelInfo string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_ConvertedFromExternalSphericalTags
	KCMFormatDescriptionExtension_ConvertedFromExternalSphericalTags string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_Depth
	KCMFormatDescriptionExtension_Depth string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_FieldCount-swift.var
	KCMFormatDescriptionExtension_FieldCount string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_FieldDetail-swift.var
	KCMFormatDescriptionExtension_FieldDetail string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_FormatName
	KCMFormatDescriptionExtension_FormatName string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_FullRangeVideo
	KCMFormatDescriptionExtension_FullRangeVideo string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_GammaLevel-swift.var
	KCMFormatDescriptionExtension_GammaLevel string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_HasAdditionalViews
	KCMFormatDescriptionExtension_HasAdditionalViews string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_HasLeftStereoEyeView
	KCMFormatDescriptionExtension_HasLeftStereoEyeView string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_HasRightStereoEyeView
	KCMFormatDescriptionExtension_HasRightStereoEyeView string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_HeroEye
	KCMFormatDescriptionExtension_HeroEye string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_HorizontalDisparityAdjustment
	KCMFormatDescriptionExtension_HorizontalDisparityAdjustment string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_HorizontalFieldOfView
	KCMFormatDescriptionExtension_HorizontalFieldOfView string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_ICCProfile
	KCMFormatDescriptionExtension_ICCProfile string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_LogTransferFunction
	KCMFormatDescriptionExtension_LogTransferFunction string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_MasteringDisplayColorVolume
	KCMFormatDescriptionExtension_MasteringDisplayColorVolume string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_OriginalCompressionSettings
	KCMFormatDescriptionExtension_OriginalCompressionSettings string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_PixelAspectRatio-swift.var
	KCMFormatDescriptionExtension_PixelAspectRatio string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_ProjectionKind
	KCMFormatDescriptionExtension_ProjectionKind string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_ProtectedContentOriginalFormat
	KCMFormatDescriptionExtension_ProtectedContentOriginalFormat string
	// KCMFormatDescriptionExtension_RevisionLevel is the value is a number.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_RevisionLevel
	KCMFormatDescriptionExtension_RevisionLevel string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_SampleDescriptionExtensionAtoms
	KCMFormatDescriptionExtension_SampleDescriptionExtensionAtoms string
	// KCMFormatDescriptionExtension_SpatialQuality is the value is a number.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_SpatialQuality
	KCMFormatDescriptionExtension_SpatialQuality string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_StereoCameraBaseline
	KCMFormatDescriptionExtension_StereoCameraBaseline string
	// KCMFormatDescriptionExtension_TemporalQuality is the value is a number.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_TemporalQuality
	KCMFormatDescriptionExtension_TemporalQuality string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_TransferFunction-swift.var
	KCMFormatDescriptionExtension_TransferFunction string
	// KCMFormatDescriptionExtension_Vendor is the value is a string of four character codes.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_Vendor
	KCMFormatDescriptionExtension_Vendor string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_VerbatimISOSampleEntry
	KCMFormatDescriptionExtension_VerbatimISOSampleEntry string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_VerbatimImageDescription-swift.var
	KCMFormatDescriptionExtension_VerbatimImageDescription string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_VerbatimSampleDescription
	KCMFormatDescriptionExtension_VerbatimSampleDescription string
	// KCMFormatDescriptionExtension_Version is the value is a number.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_Version
	KCMFormatDescriptionExtension_Version string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_ViewPackingKind
	KCMFormatDescriptionExtension_ViewPackingKind string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_YCbCrMatrix-swift.var
	KCMFormatDescriptionExtension_YCbCrMatrix string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionFieldDetail_SpatialFirstLineEarly-swift.var
	KCMFormatDescriptionFieldDetail_SpatialFirstLineEarly string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionFieldDetail_SpatialFirstLineLate-swift.var
	KCMFormatDescriptionFieldDetail_SpatialFirstLineLate string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionFieldDetail_TemporalBottomFirst-swift.var
	KCMFormatDescriptionFieldDetail_TemporalBottomFirst string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionFieldDetail_TemporalTopFirst-swift.var
	KCMFormatDescriptionFieldDetail_TemporalTopFirst string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionHeroEye_Left
	KCMFormatDescriptionHeroEye_Left string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionHeroEye_Right
	KCMFormatDescriptionHeroEye_Right string
	// KCMFormatDescriptionKey_CleanApertureHeight is a key that identifies the clean aperture height value.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionKey_CleanApertureHeight-swift.var
	KCMFormatDescriptionKey_CleanApertureHeight string
	// KCMFormatDescriptionKey_CleanApertureHeightRational is a key that identifies the clean aperture height rational value.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionKey_CleanApertureHeightRational
	KCMFormatDescriptionKey_CleanApertureHeightRational string
	// KCMFormatDescriptionKey_CleanApertureHorizontalOffset is a key that describes the clean aperture height horizontal offset value.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionKey_CleanApertureHorizontalOffset-swift.var
	KCMFormatDescriptionKey_CleanApertureHorizontalOffset string
	// KCMFormatDescriptionKey_CleanApertureHorizontalOffsetRational is a key that identifies the clean aperture horizontal offset rational value.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionKey_CleanApertureHorizontalOffsetRational
	KCMFormatDescriptionKey_CleanApertureHorizontalOffsetRational string
	// KCMFormatDescriptionKey_CleanApertureVerticalOffset is a key that identifies the clean aperture vertical offset value.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionKey_CleanApertureVerticalOffset-swift.var
	KCMFormatDescriptionKey_CleanApertureVerticalOffset string
	// KCMFormatDescriptionKey_CleanApertureVerticalOffsetRational is a key that identifies the clean aperture vertical offset rational value.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionKey_CleanApertureVerticalOffsetRational
	KCMFormatDescriptionKey_CleanApertureVerticalOffsetRational string
	// KCMFormatDescriptionKey_CleanApertureWidth is a key that identifies the clean aperture width value.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionKey_CleanApertureWidth-swift.var
	KCMFormatDescriptionKey_CleanApertureWidth string
	// KCMFormatDescriptionKey_CleanApertureWidthRational is a key that identifies the clean aperture width rational value.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionKey_CleanApertureWidthRational
	KCMFormatDescriptionKey_CleanApertureWidthRational string
	// KCMFormatDescriptionKey_PixelAspectRatioHorizontalSpacing is a key that identifies the horizontal spacing of the pixel aspect ratio.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionKey_PixelAspectRatioHorizontalSpacing-swift.var
	KCMFormatDescriptionKey_PixelAspectRatioHorizontalSpacing string
	// KCMFormatDescriptionKey_PixelAspectRatioVerticalSpacing is a key that identifies the vertical spacing of the pixel aspect ratio.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionKey_PixelAspectRatioVerticalSpacing-swift.var
	KCMFormatDescriptionKey_PixelAspectRatioVerticalSpacing string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionLogTransferFunction_AppleLog
	KCMFormatDescriptionLogTransferFunction_AppleLog string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionProjectionKind_AppleImmersiveVideo
	KCMFormatDescriptionProjectionKind_AppleImmersiveVideo string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionProjectionKind_Equirectangular
	KCMFormatDescriptionProjectionKind_Equirectangular string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionProjectionKind_HalfEquirectangular
	KCMFormatDescriptionProjectionKind_HalfEquirectangular string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionProjectionKind_ParametricImmersive
	KCMFormatDescriptionProjectionKind_ParametricImmersive string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionProjectionKind_Rectilinear
	KCMFormatDescriptionProjectionKind_Rectilinear string
	// KCMFormatDescriptionTransferFunction_ITU_R_2020 is a constant that describes the ITU R 2020 transfer function format.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionTransferFunction_ITU_R_2020
	KCMFormatDescriptionTransferFunction_ITU_R_2020 string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionTransferFunction_ITU_R_2100_HLG
	KCMFormatDescriptionTransferFunction_ITU_R_2100_HLG string
	// KCMFormatDescriptionTransferFunction_ITU_R_709_2 is a constant that describes the ITU R 709 2 transfer function format.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionTransferFunction_ITU_R_709_2-swift.var
	KCMFormatDescriptionTransferFunction_ITU_R_709_2 string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionTransferFunction_Linear
	KCMFormatDescriptionTransferFunction_Linear string
	// KCMFormatDescriptionTransferFunction_SMPTE_240M_1995 is a constant that describes the SMPTE 240M 1995 transfer function format.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionTransferFunction_SMPTE_240M_1995-swift.var
	KCMFormatDescriptionTransferFunction_SMPTE_240M_1995 string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionTransferFunction_SMPTE_ST_2084_PQ
	KCMFormatDescriptionTransferFunction_SMPTE_ST_2084_PQ string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionTransferFunction_SMPTE_ST_428_1
	KCMFormatDescriptionTransferFunction_SMPTE_ST_428_1 string
	// KCMFormatDescriptionTransferFunction_UseGamma is a constant that describes the gamma transfer function format.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionTransferFunction_UseGamma-swift.var
	KCMFormatDescriptionTransferFunction_UseGamma string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionTransferFunction_sRGB
	KCMFormatDescriptionTransferFunction_sRGB string
	// KCMFormatDescriptionVendor_Apple is a string that speicifies Apple as the vendor.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionVendor_Apple
	KCMFormatDescriptionVendor_Apple string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionViewPackingKind_OverUnder
	KCMFormatDescriptionViewPackingKind_OverUnder string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionViewPackingKind_SideBySide
	KCMFormatDescriptionViewPackingKind_SideBySide string
	// KCMFormatDescriptionYCbCrMatrix_ITU_R_2020 is a constant describing the YCbCrMatrix ITU R 2020 format.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionYCbCrMatrix_ITU_R_2020
	KCMFormatDescriptionYCbCrMatrix_ITU_R_2020 string
	// KCMFormatDescriptionYCbCrMatrix_ITU_R_601_4 is a constant describing the YCbCrMatrix ITU R 601 4 format.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionYCbCrMatrix_ITU_R_601_4-swift.var
	KCMFormatDescriptionYCbCrMatrix_ITU_R_601_4 string
	// KCMFormatDescriptionYCbCrMatrix_ITU_R_709_2 is a constant describing the YCbCrMatrix ITU R 709 2 format.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionYCbCrMatrix_ITU_R_709_2-swift.var
	KCMFormatDescriptionYCbCrMatrix_ITU_R_709_2 string
	// KCMFormatDescriptionYCbCrMatrix_SMPTE_240M_1995 is a constant describing the YCbCrMatrix SMPTE 240M 1995 format.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionYCbCrMatrix_SMPTE_240M_1995-swift.var
	KCMFormatDescriptionYCbCrMatrix_SMPTE_240M_1995 string
	// KCMHEVCTemporalLevelInfoKey_ConstraintIndicatorFlags is the value is a 6 byte data value.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMHEVCTemporalLevelInfoKey_ConstraintIndicatorFlags
	KCMHEVCTemporalLevelInfoKey_ConstraintIndicatorFlags string
	// KCMHEVCTemporalLevelInfoKey_LevelIndex is the value is a number.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMHEVCTemporalLevelInfoKey_LevelIndex
	KCMHEVCTemporalLevelInfoKey_LevelIndex string
	// KCMHEVCTemporalLevelInfoKey_ProfileCompatibilityFlags is the value is a 4 byte data value.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMHEVCTemporalLevelInfoKey_ProfileCompatibilityFlags
	KCMHEVCTemporalLevelInfoKey_ProfileCompatibilityFlags string
	// KCMHEVCTemporalLevelInfoKey_ProfileIndex is the value is a number.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMHEVCTemporalLevelInfoKey_ProfileIndex
	KCMHEVCTemporalLevelInfoKey_ProfileIndex string
	// KCMHEVCTemporalLevelInfoKey_ProfileSpace is the value is a number.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMHEVCTemporalLevelInfoKey_ProfileSpace
	KCMHEVCTemporalLevelInfoKey_ProfileSpace string
	// KCMHEVCTemporalLevelInfoKey_TemporalLevel is the value is a number.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMHEVCTemporalLevelInfoKey_TemporalLevel
	KCMHEVCTemporalLevelInfoKey_TemporalLevel string
	// KCMHEVCTemporalLevelInfoKey_TierFlag is the value is a number.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMHEVCTemporalLevelInfoKey_TierFlag
	KCMHEVCTemporalLevelInfoKey_TierFlag string
	// KCMMemoryPoolOption_AgeOutPeriod is the period of time before the pool recycles its memory.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMemoryPoolOption_AgeOutPeriod
	KCMMemoryPoolOption_AgeOutPeriod string
	// KCMMetadataBaseDataType_AffineTransformF64 is a type that identifies a 3x3 matrix of 64-bit big endian floating point numbers in a row-major order that specify an affine transform.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_AffineTransformF64
	KCMMetadataBaseDataType_AffineTransformF64 string
	// KCMMetadataBaseDataType_BMP is bMP image.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_BMP
	KCMMetadataBaseDataType_BMP string
	// KCMMetadataBaseDataType_DimensionsF32 is consists of a 32-bit big endian floating point x value followed by a 32-bit floating point y value.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_DimensionsF32
	KCMMetadataBaseDataType_DimensionsF32 string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_ExtendedRasterRectangleValue
	KCMMetadataBaseDataType_ExtendedRasterRectangleValue string
	// KCMMetadataBaseDataType_Float32 is 32-bit big endian floating point number.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_Float32
	KCMMetadataBaseDataType_Float32 string
	// KCMMetadataBaseDataType_Float64 is 64-bit big endian floating point number.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_Float64
	KCMMetadataBaseDataType_Float64 string
	// KCMMetadataBaseDataType_GIF is gIF image.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_GIF
	KCMMetadataBaseDataType_GIF string
	// KCMMetadataBaseDataType_JPEG is jPEG image.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_JPEG
	KCMMetadataBaseDataType_JPEG string
	// KCMMetadataBaseDataType_JSON is uTF-8 encoded JSON data.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_JSON
	KCMMetadataBaseDataType_JSON string
	// KCMMetadataBaseDataType_PNG is pNG image.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_PNG
	KCMMetadataBaseDataType_PNG string
	// KCMMetadataBaseDataType_PerspectiveTransformF64 is a 3x3 matrix of 64-bit big endian floating point numbers the system stores in row-major order that specify a perspective transform.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_PerspectiveTransformF64
	KCMMetadataBaseDataType_PerspectiveTransformF64 string
	// KCMMetadataBaseDataType_PointF32 is consists of two 32-bit big endian floating point values, the x and y values, respectively.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_PointF32
	KCMMetadataBaseDataType_PointF32 string
	// KCMMetadataBaseDataType_PolygonF32 is three or more pairs of 32-bit floating point numbers (x and y values) that define the vertices of a polygon.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_PolygonF32
	KCMMetadataBaseDataType_PolygonF32 string
	// KCMMetadataBaseDataType_PolylineF32 is two or more pairs of 32-bit floating point numbers (x and y values) that define a multi-segmented line.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_PolylineF32
	KCMMetadataBaseDataType_PolylineF32 string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_RasterRectangleValue
	KCMMetadataBaseDataType_RasterRectangleValue string
	// KCMMetadataBaseDataType_RawData is a sequence of bytes whose interpretation based upon an agreement between the reader and the writer.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_RawData
	KCMMetadataBaseDataType_RawData string
	// KCMMetadataBaseDataType_RectF32 is consists of four 32-bit big endian floating point values, the origin’s x, origin’s y, width and height values, respectively. May also be interpreted as a 32-bit floating point origin followed by a 32-bit floating point dimension.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_RectF32
	KCMMetadataBaseDataType_RectF32 string
	// KCMMetadataBaseDataType_SInt16 is 16-bit big endian signed integer.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_SInt16
	KCMMetadataBaseDataType_SInt16 string
	// KCMMetadataBaseDataType_SInt32 is 32-bit big endian signed integer.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_SInt32
	KCMMetadataBaseDataType_SInt32 string
	// KCMMetadataBaseDataType_SInt64 is 64-bit big endian signed integer.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_SInt64
	KCMMetadataBaseDataType_SInt64 string
	// KCMMetadataBaseDataType_SInt8 is 8-bit signed integer.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_SInt8
	KCMMetadataBaseDataType_SInt8 string
	// KCMMetadataBaseDataType_UInt16 is 16-bit big endian unsigned integer.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_UInt16
	KCMMetadataBaseDataType_UInt16 string
	// KCMMetadataBaseDataType_UInt32 is 32-bit big endian unsigned integer.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_UInt32
	KCMMetadataBaseDataType_UInt32 string
	// KCMMetadataBaseDataType_UInt64 is 64-bit big endian unsigned integer.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_UInt64
	KCMMetadataBaseDataType_UInt64 string
	// KCMMetadataBaseDataType_UInt8 is 8-bit unsigned integer.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_UInt8
	KCMMetadataBaseDataType_UInt8 string
	// KCMMetadataBaseDataType_UTF16 is uTF-16 string.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_UTF16
	KCMMetadataBaseDataType_UTF16 string
	// KCMMetadataBaseDataType_UTF8 is uTF-8 string.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataBaseDataType_UTF8
	KCMMetadataBaseDataType_UTF8 string
	// KCMMetadataDataType_QuickTimeMetadataDirection is a string supplying degrees offset from magnetic North.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataDataType_QuickTimeMetadataDirection
	KCMMetadataDataType_QuickTimeMetadataDirection string
	// KCMMetadataDataType_QuickTimeMetadataLocation_ISO6709 is a string supplying location information in ISO-6709 format.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataDataType_QuickTimeMetadataLocation_ISO6709
	KCMMetadataDataType_QuickTimeMetadataLocation_ISO6709 string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataDataType_QuickTimeMetadataMilliLux
	KCMMetadataDataType_QuickTimeMetadataMilliLux string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataDataType_QuickTimeMetadataUUID
	KCMMetadataDataType_QuickTimeMetadataUUID string
	// KCMMetadataFormatDescriptionKey_ConformingDataTypes is a key that identifies the conforming data types.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataFormatDescriptionKey_ConformingDataTypes
	KCMMetadataFormatDescriptionKey_ConformingDataTypes string
	// KCMMetadataFormatDescriptionKey_DataType is a key that identifies the data type.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataFormatDescriptionKey_DataType
	KCMMetadataFormatDescriptionKey_DataType string
	// KCMMetadataFormatDescriptionKey_DataTypeNamespace is a key that identifies the data type namespace.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataFormatDescriptionKey_DataTypeNamespace
	KCMMetadataFormatDescriptionKey_DataTypeNamespace string
	// KCMMetadataFormatDescriptionKey_LanguageTag is a key that identifies the language tag.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataFormatDescriptionKey_LanguageTag
	KCMMetadataFormatDescriptionKey_LanguageTag string
	// KCMMetadataFormatDescriptionKey_LocalID is a key that identifies the local identifier.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataFormatDescriptionKey_LocalID
	KCMMetadataFormatDescriptionKey_LocalID string
	// KCMMetadataFormatDescriptionKey_Namespace is a key that identifies the namespace.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataFormatDescriptionKey_Namespace
	KCMMetadataFormatDescriptionKey_Namespace string
	// KCMMetadataFormatDescriptionKey_SetupData is a key that identifies setup data.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataFormatDescriptionKey_SetupData
	KCMMetadataFormatDescriptionKey_SetupData string
	// KCMMetadataFormatDescriptionKey_StructuralDependency is a key that identifies the structural dependency.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataFormatDescriptionKey_StructuralDependency
	KCMMetadataFormatDescriptionKey_StructuralDependency string
	// KCMMetadataFormatDescriptionKey_Value is a key that identifies the value.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataFormatDescriptionKey_Value
	KCMMetadataFormatDescriptionKey_Value string
	// KCMMetadataFormatDescriptionMetadataSpecificationKey_DataType is a specification key that identifies the data type.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataFormatDescriptionMetadataSpecificationKey_DataType
	KCMMetadataFormatDescriptionMetadataSpecificationKey_DataType string
	// KCMMetadataFormatDescriptionMetadataSpecificationKey_ExtendedLanguageTag is a specification key that identifies the extended language tag.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataFormatDescriptionMetadataSpecificationKey_ExtendedLanguageTag
	KCMMetadataFormatDescriptionMetadataSpecificationKey_ExtendedLanguageTag string
	// KCMMetadataFormatDescriptionMetadataSpecificationKey_Identifier is a specification key that identifies the identifier.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataFormatDescriptionMetadataSpecificationKey_Identifier
	KCMMetadataFormatDescriptionMetadataSpecificationKey_Identifier string
	// KCMMetadataFormatDescriptionMetadataSpecificationKey_SetupData is a specification key that identifies the setup data.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataFormatDescriptionMetadataSpecificationKey_SetupData
	KCMMetadataFormatDescriptionMetadataSpecificationKey_SetupData string
	// KCMMetadataFormatDescriptionMetadataSpecificationKey_StructuralDependency is a specification key that identifies dependencies.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataFormatDescriptionMetadataSpecificationKey_StructuralDependency
	KCMMetadataFormatDescriptionMetadataSpecificationKey_StructuralDependency string
	// KCMMetadataFormatDescription_StructuralDependencyKey_DependencyIsInvalidFlag is a specification key that identifies the depencency isn’t valid.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataFormatDescription_StructuralDependencyKey_DependencyIsInvalidFlag
	KCMMetadataFormatDescription_StructuralDependencyKey_DependencyIsInvalidFlag string
	// KCMMetadataIdentifier_QuickTimeMetadataDirection_Facing is direction the observer is facing.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataIdentifier_QuickTimeMetadataDirection_Facing
	KCMMetadataIdentifier_QuickTimeMetadataDirection_Facing string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataIdentifier_QuickTimeMetadataDisplayMaskRectangleMono
	KCMMetadataIdentifier_QuickTimeMetadataDisplayMaskRectangleMono string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataIdentifier_QuickTimeMetadataDisplayMaskRectangleStereoLeft
	KCMMetadataIdentifier_QuickTimeMetadataDisplayMaskRectangleStereoLeft string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataIdentifier_QuickTimeMetadataDisplayMaskRectangleStereoRight
	KCMMetadataIdentifier_QuickTimeMetadataDisplayMaskRectangleStereoRight string
	// KCMMetadataIdentifier_QuickTimeMetadataLivePhotoStillImageTransform is a perspective transform you use to adjust a Live Photo still image to match the Live Photo movie.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataIdentifier_QuickTimeMetadataLivePhotoStillImageTransform
	KCMMetadataIdentifier_QuickTimeMetadataLivePhotoStillImageTransform string
	// KCMMetadataIdentifier_QuickTimeMetadataLivePhotoStillImageTransformReferenceDimensions is the dimensions of the live photo still image.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataIdentifier_QuickTimeMetadataLivePhotoStillImageTransformReferenceDimensions
	KCMMetadataIdentifier_QuickTimeMetadataLivePhotoStillImageTransformReferenceDimensions string
	// KCMMetadataIdentifier_QuickTimeMetadataLocation_ISO6709 is location information in ISO-6709 format.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataIdentifier_QuickTimeMetadataLocation_ISO6709
	KCMMetadataIdentifier_QuickTimeMetadataLocation_ISO6709 string
	// KCMMetadataIdentifier_QuickTimeMetadataPreferredAffineTransform is an affine transform to be applied to a video track.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataIdentifier_QuickTimeMetadataPreferredAffineTransform
	KCMMetadataIdentifier_QuickTimeMetadataPreferredAffineTransform string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataIdentifier_QuickTimeMetadataPresentationImmersiveMedia
	KCMMetadataIdentifier_QuickTimeMetadataPresentationImmersiveMedia string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataIdentifier_QuickTimeMetadataSceneIlluminance
	KCMMetadataIdentifier_QuickTimeMetadataSceneIlluminance string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataIdentifier_QuickTimeMetadataSegmentIdentifier
	KCMMetadataIdentifier_QuickTimeMetadataSegmentIdentifier string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataIdentifier_QuickTimeMetadataSpatialAudioMix
	KCMMetadataIdentifier_QuickTimeMetadataSpatialAudioMix string
	// KCMMetadataIdentifier_QuickTimeMetadataVideoOrientation is video orientation as defined by TIFF/EXIF.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataIdentifier_QuickTimeMetadataVideoOrientation
	KCMMetadataIdentifier_QuickTimeMetadataVideoOrientation string
	// KCMMetadataKeySpace_HLSDateRange is metadata keyspace for HLS DateRange tags.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataKeySpace_HLSDateRange
	KCMMetadataKeySpace_HLSDateRange string
	// KCMMetadataKeySpace_ID3 is metadata keyspace for ID3 keys.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataKeySpace_ID3
	KCMMetadataKeySpace_ID3 string
	// KCMMetadataKeySpace_ISOUserData is metadata keyspace for MPEG-4 User Data keys.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataKeySpace_ISOUserData
	KCMMetadataKeySpace_ISOUserData string
	// KCMMetadataKeySpace_Icy is metadata keyspace for ShoutCast keys.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataKeySpace_Icy
	KCMMetadataKeySpace_Icy string
	// KCMMetadataKeySpace_QuickTimeMetadata is metadata keyspace for QuickTime Metadata keys.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataKeySpace_QuickTimeMetadata
	KCMMetadataKeySpace_QuickTimeMetadata string
	// KCMMetadataKeySpace_QuickTimeUserData is metadata keyspace for QuickTime User Data keys.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataKeySpace_QuickTimeUserData
	KCMMetadataKeySpace_QuickTimeUserData string
	// KCMMetadataKeySpace_iTunes is metadata keyspace for iTunes keys.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMMetadataKeySpace_iTunes
	KCMMetadataKeySpace_iTunes string
	// KCMSampleAttachmentKey_AudioIndependentSampleDecoderRefreshCount is an attachment that’s only present if the audio sample is an independent frame or immediate playout frame.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_AudioIndependentSampleDecoderRefreshCount
	KCMSampleAttachmentKey_AudioIndependentSampleDecoderRefreshCount string
	// KCMSampleAttachmentKey_CryptorSubsampleAuxiliaryData is an attachment that describes the ranges of protected and unprotected data within a protected sample buffer.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_CryptorSubsampleAuxiliaryData
	KCMSampleAttachmentKey_CryptorSubsampleAuxiliaryData string
	// KCMSampleAttachmentKey_DependsOnOthers is indicates whether the sample depends on other samples for decoding (type [CFBoolean]).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_DependsOnOthers
	KCMSampleAttachmentKey_DependsOnOthers string
	// KCMSampleAttachmentKey_DisplayImmediately is indicates whether the sample should be displayed immediately (type [CFBoolean], default false).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_DisplayImmediately
	KCMSampleAttachmentKey_DisplayImmediately string
	// KCMSampleAttachmentKey_DoNotDisplay is indicates whether the sample should be decoded but not displayed (type [CFBoolean], default false).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_DoNotDisplay
	KCMSampleAttachmentKey_DoNotDisplay string
	// KCMSampleAttachmentKey_EarlierDisplayTimesAllowed is indicates whether later samples may have earlier display times (type [CFBoolean]).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_EarlierDisplayTimesAllowed
	KCMSampleAttachmentKey_EarlierDisplayTimesAllowed string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_HDR10PlusPerFrameData
	KCMSampleAttachmentKey_HDR10PlusPerFrameData string
	// KCMSampleAttachmentKey_HEVCStepwiseTemporalSubLayerAccess is an attachment that indicates a step-wise temporal sublayer access (STSA) sample grouping.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_HEVCStepwiseTemporalSubLayerAccess
	KCMSampleAttachmentKey_HEVCStepwiseTemporalSubLayerAccess string
	// KCMSampleAttachmentKey_HEVCSyncSampleNALUnitType is an attachment that indicates a sync sample NAL unit type.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_HEVCSyncSampleNALUnitType
	KCMSampleAttachmentKey_HEVCSyncSampleNALUnitType string
	// KCMSampleAttachmentKey_HEVCTemporalLevelInfo is an attachment that indicates a video frame’s level within a hierarchical frame dependency structure.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_HEVCTemporalLevelInfo
	KCMSampleAttachmentKey_HEVCTemporalLevelInfo string
	// KCMSampleAttachmentKey_HEVCTemporalSubLayerAccess is an attachment that indicates a temporal sublayer access grouping.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_HEVCTemporalSubLayerAccess
	KCMSampleAttachmentKey_HEVCTemporalSubLayerAccess string
	// KCMSampleAttachmentKey_HasRedundantCoding is indicates whether the sample has redundant coding (type [CFBoolean]).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_HasRedundantCoding
	KCMSampleAttachmentKey_HasRedundantCoding string
	// KCMSampleAttachmentKey_IsDependedOnByOthers is indicates whether other samples depend on this sample for decoding (type [CFBoolean]).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_IsDependedOnByOthers
	KCMSampleAttachmentKey_IsDependedOnByOthers string
	// KCMSampleAttachmentKey_NotSync is indicates whether the sample is a sync sample (type [CFBoolean], default false).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_NotSync
	KCMSampleAttachmentKey_NotSync string
	// KCMSampleAttachmentKey_PartialSync is indicates whether the sample is a partial sync sample (type [CFBoolean], default false).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_PartialSync
	KCMSampleAttachmentKey_PartialSync string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_PostDecodeProcessingMetadata
	KCMSampleAttachmentKey_PostDecodeProcessingMetadata string
	// KCMSampleBufferAttachmentKey_CameraIntrinsicMatrix is an attachment that indicates a 3x3 camera intrinsic matrix to apply to the current sample buffer.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_CameraIntrinsicMatrix
	KCMSampleBufferAttachmentKey_CameraIntrinsicMatrix string
	// KCMSampleBufferAttachmentKey_DisplayEmptyMediaImmediately is tells that the empty marker should be dequeued immediately regardless of its timestamp (type [CFBoolean], default false).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_DisplayEmptyMediaImmediately
	KCMSampleBufferAttachmentKey_DisplayEmptyMediaImmediately string
	// KCMSampleBufferAttachmentKey_DrainAfterDecoding is indicates whether the sample buffer should be drained after decoding type [CFBoolean], default false).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_DrainAfterDecoding
	KCMSampleBufferAttachmentKey_DrainAfterDecoding string
	// KCMSampleBufferAttachmentKey_DroppedFrameReason is indicates the reason the current video frame was dropped (type [CFString]).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_DroppedFrameReason
	KCMSampleBufferAttachmentKey_DroppedFrameReason string
	// KCMSampleBufferAttachmentKey_DroppedFrameReasonInfo is indicates additional information regarding the dropped video frame (type [CFString]).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_DroppedFrameReasonInfo
	KCMSampleBufferAttachmentKey_DroppedFrameReasonInfo string
	// KCMSampleBufferAttachmentKey_EmptyMedia is marks an intentionally empty interval in the sequence of samples (type [CFBoolean], default false).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_EmptyMedia
	KCMSampleBufferAttachmentKey_EmptyMedia string
	// KCMSampleBufferAttachmentKey_EndsPreviousSampleDuration is indicates that sample buffer’s decode timestamp may be used to define the previous sample buffer’s duration (type [CFBoolean], default false).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_EndsPreviousSampleDuration
	KCMSampleBufferAttachmentKey_EndsPreviousSampleDuration string
	// KCMSampleBufferAttachmentKey_FillDiscontinuitiesWithSilence is fill the difference between discontiguous sample buffers with silence (type [CFBoolean], default false).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_FillDiscontinuitiesWithSilence
	KCMSampleBufferAttachmentKey_FillDiscontinuitiesWithSilence string
	// KCMSampleBufferAttachmentKey_ForceKeyFrame is indicates that the current or next video sample buffer should be forced to be encoded as a key frame.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_ForceKeyFrame
	KCMSampleBufferAttachmentKey_ForceKeyFrame string
	// KCMSampleBufferAttachmentKey_GradualDecoderRefresh is indicates the decoder refresh count (type [CFNumber]).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_GradualDecoderRefresh
	KCMSampleBufferAttachmentKey_GradualDecoderRefresh string
	// KCMSampleBufferAttachmentKey_PermanentEmptyMedia is marks the end of the sequence of samples (type [CFBoolean], default false).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_PermanentEmptyMedia
	KCMSampleBufferAttachmentKey_PermanentEmptyMedia string
	// KCMSampleBufferAttachmentKey_PostNotificationWhenConsumed is if present, indicates that decode pipelines should post a notification when consuming the sample buffer(type [CFDictionary]).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_PostNotificationWhenConsumed
	KCMSampleBufferAttachmentKey_PostNotificationWhenConsumed string
	// KCMSampleBufferAttachmentKey_ResetDecoderBeforeDecoding is indicates whether the sample buffer should be reset before decoding (type [CFBoolean], default false).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_ResetDecoderBeforeDecoding
	KCMSampleBufferAttachmentKey_ResetDecoderBeforeDecoding string
	// KCMSampleBufferAttachmentKey_ResumeOutput is if present, indicates that output should be resumed following a discontinuity [CFBoolean], default false).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_ResumeOutput
	KCMSampleBufferAttachmentKey_ResumeOutput string
	// KCMSampleBufferAttachmentKey_Reverse is indicates that the decoded contents of the sample buffer should be reversed (type [CFBoolean], default false).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_Reverse
	KCMSampleBufferAttachmentKey_Reverse string
	// KCMSampleBufferAttachmentKey_SampleReferenceByteOffset is indicates the byte offset at which the sample data begins (type [CFNumber]).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_SampleReferenceByteOffset
	KCMSampleBufferAttachmentKey_SampleReferenceByteOffset string
	// KCMSampleBufferAttachmentKey_SampleReferenceURL is indicates the URL where the sample data is (type [CFURL]).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_SampleReferenceURL
	KCMSampleBufferAttachmentKey_SampleReferenceURL string
	// KCMSampleBufferAttachmentKey_SpeedMultiplier is the factor by which the sample buffer’s presentation should be accelerated (type [CFNumber], default 1.0).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_SpeedMultiplier
	KCMSampleBufferAttachmentKey_SpeedMultiplier string
	// KCMSampleBufferAttachmentKey_StillImageLensStabilizationInfo is indicates information about the lens stabilization applied to the current still image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_StillImageLensStabilizationInfo
	KCMSampleBufferAttachmentKey_StillImageLensStabilizationInfo string
	// KCMSampleBufferAttachmentKey_TransitionID is marks a transition from one source of buffers to another.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_TransitionID
	KCMSampleBufferAttachmentKey_TransitionID string
	// KCMSampleBufferAttachmentKey_TrimDurationAtEnd is the duration that should be removed at the end of the sample buffer, after decoding.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_TrimDurationAtEnd
	KCMSampleBufferAttachmentKey_TrimDurationAtEnd string
	// KCMSampleBufferAttachmentKey_TrimDurationAtStart is the duration that should be removed at the beginning of the sample buffer, after decoding.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_TrimDurationAtStart
	KCMSampleBufferAttachmentKey_TrimDurationAtStart string
	// KCMSampleBufferConduitNotificationParameter_MaxUpcomingOutputPTS is specifies the maximum presentation timestamp of upcoming output samples (type [CFDictionary]).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferConduitNotificationParameter_MaxUpcomingOutputPTS
	KCMSampleBufferConduitNotificationParameter_MaxUpcomingOutputPTS string
	// KCMSampleBufferConduitNotificationParameter_MinUpcomingOutputPTS is specifies the minimum presentation timestamp of upcoming output samples (type [CFDictionary]).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferConduitNotificationParameter_MinUpcomingOutputPTS
	KCMSampleBufferConduitNotificationParameter_MinUpcomingOutputPTS string
	// KCMSampleBufferConduitNotificationParameter_ResumeTag is specifies a tag to be attached to the first sample buffer following a discontinuity (type [CFNumber]).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferConduitNotificationParameter_ResumeTag
	KCMSampleBufferConduitNotificationParameter_ResumeTag string
	// KCMSampleBufferConduitNotificationParameter_UpcomingOutputPTSRangeMayOverlapQueuedOutputPTSRange is indicates that the presentation timestamps of upcoming output samples may overlap those of samples queued for output (type [CFBoolean]).
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferConduitNotificationParameter_UpcomingOutputPTSRangeMayOverlapQueuedOutputPTSRange
	KCMSampleBufferConduitNotificationParameter_UpcomingOutputPTSRangeMayOverlapQueuedOutputPTSRange string
	// KCMSampleBufferConduitNotification_InhibitOutputUntil is posted on a conduit of sample buffers to announce a coming discontinuity.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferConduitNotification_InhibitOutputUntil
	KCMSampleBufferConduitNotification_InhibitOutputUntil string
	// KCMSampleBufferConduitNotification_ResetOutput is posted on a conduit of sample buffers to request invalidation of pending output data.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferConduitNotification_ResetOutput
	KCMSampleBufferConduitNotification_ResetOutput string
	// KCMSampleBufferConduitNotification_UpcomingOutputPTSRangeChanged is posted on a conduit of video sample buffers to report information about the range of upcoming output presentation timestamps.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferConduitNotification_UpcomingOutputPTSRangeChanged
	KCMSampleBufferConduitNotification_UpcomingOutputPTSRangeChanged string
	// KCMSampleBufferConsumerNotification_BufferConsumed is optionally posted when a sample buffer is consumed.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferConsumerNotification_BufferConsumed
	KCMSampleBufferConsumerNotification_BufferConsumed string
	// KCMSampleBufferDroppedFrameReasonInfo_CameraModeSwitch is a discontinuity was caused by a camera mode switch.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferDroppedFrameReasonInfo_CameraModeSwitch
	KCMSampleBufferDroppedFrameReasonInfo_CameraModeSwitch string
	// KCMSampleBufferDroppedFrameReason_Discontinuity is an unknown number of frames were dropped.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferDroppedFrameReason_Discontinuity
	KCMSampleBufferDroppedFrameReason_Discontinuity string
	// KCMSampleBufferDroppedFrameReason_FrameWasLate is the frame was dropped because it was late.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferDroppedFrameReason_FrameWasLate
	KCMSampleBufferDroppedFrameReason_FrameWasLate string
	// KCMSampleBufferDroppedFrameReason_OutOfBuffers is the frame was dropped because the module providing frames is out of buffers.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferDroppedFrameReason_OutOfBuffers
	KCMSampleBufferDroppedFrameReason_OutOfBuffers string
	// KCMSampleBufferLensStabilizationInfo_Active is the lens stabilization module was active for the duration this buffer.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferLensStabilizationInfo_Active
	KCMSampleBufferLensStabilizationInfo_Active string
	// KCMSampleBufferLensStabilizationInfo_Off is the lens stabilization module was not used during this capture.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferLensStabilizationInfo_Off
	KCMSampleBufferLensStabilizationInfo_Off string
	// KCMSampleBufferLensStabilizationInfo_OutOfRange is the motion of the device or duration of the capture was outside of what the stabilization mechanism could support.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferLensStabilizationInfo_OutOfRange
	KCMSampleBufferLensStabilizationInfo_OutOfRange string
	// KCMSampleBufferLensStabilizationInfo_Unavailable is the lens stabilization module was unavailable for use.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferLensStabilizationInfo_Unavailable
	KCMSampleBufferLensStabilizationInfo_Unavailable string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferNotificationParameter_OSStatus
	KCMSampleBufferNotificationParameter_OSStatus string
	// KCMSampleBufferNotification_DataBecameReady is posted on a sample buffer by the [CMSampleBufferSetDataReady(_:)] function when the buffer becomes ready.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferNotification_DataBecameReady
	KCMSampleBufferNotification_DataBecameReady string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferNotification_DataFailed
	KCMSampleBufferNotification_DataFailed string
	// KCMTagCategoryKey is a constant for use as a key during tag creation from a dictionary, whose value is the tag’s category.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagCategoryKey
	KCMTagCategoryKey string
	// KCMTagCollectionTagsArrayKey is a dictionary key for assigning tag collections to.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagCollectionTagsArrayKey
	KCMTagCollectionTagsArrayKey string
	// KCMTagDataTypeKey is a constant for use as a key during tag creation from a dictionary, whose value is the tag’s data type.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagDataTypeKey
	KCMTagDataTypeKey string
	// KCMTagValueKey is a constant for use as a key during tag creation from a dictionary, whose value is the tag’s contained value.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagValueKey
	KCMTagValueKey string
	// KCMTextFormatDescriptionColor_Alpha is the color value for the alpha is a number.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionColor_Alpha
	KCMTextFormatDescriptionColor_Alpha string
	// KCMTextFormatDescriptionColor_Blue is the color value for blue is a number.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionColor_Blue
	KCMTextFormatDescriptionColor_Blue string
	// KCMTextFormatDescriptionColor_Green is the color value for green is a number.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionColor_Green
	KCMTextFormatDescriptionColor_Green string
	// KCMTextFormatDescriptionColor_Red is the color value for red is a number.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionColor_Red
	KCMTextFormatDescriptionColor_Red string
	// KCMTextFormatDescriptionExtension_BackgroundColor is the extension flag that represents the background color.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionExtension_BackgroundColor
	KCMTextFormatDescriptionExtension_BackgroundColor string
	// KCMTextFormatDescriptionExtension_DefaultFontName is the extension flag that represents the default font name.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionExtension_DefaultFontName
	KCMTextFormatDescriptionExtension_DefaultFontName string
	// KCMTextFormatDescriptionExtension_DefaultStyle is the extension flag that represents the default style.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionExtension_DefaultStyle
	KCMTextFormatDescriptionExtension_DefaultStyle string
	// KCMTextFormatDescriptionExtension_DefaultTextBox is the extension flag that represents the default text box.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionExtension_DefaultTextBox
	KCMTextFormatDescriptionExtension_DefaultTextBox string
	// KCMTextFormatDescriptionExtension_DisplayFlags is the extension flag that represents the display flags.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionExtension_DisplayFlags
	KCMTextFormatDescriptionExtension_DisplayFlags string
	// KCMTextFormatDescriptionExtension_FontTable is the extension flag that represents the font table.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionExtension_FontTable
	KCMTextFormatDescriptionExtension_FontTable string
	// KCMTextFormatDescriptionExtension_HorizontalJustification is the extension flag that represents the horizontal justification.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionExtension_HorizontalJustification
	KCMTextFormatDescriptionExtension_HorizontalJustification string
	// KCMTextFormatDescriptionExtension_TextJustification is the extension flag that represents the text justification.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionExtension_TextJustification
	KCMTextFormatDescriptionExtension_TextJustification string
	// KCMTextFormatDescriptionExtension_VerticalJustification is the extension flag that represents the vertical justification.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionExtension_VerticalJustification
	KCMTextFormatDescriptionExtension_VerticalJustification string
	// KCMTextFormatDescriptionRect_Bottom is the bottom value for the rect as a 16-bit number.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionRect_Bottom
	KCMTextFormatDescriptionRect_Bottom string
	// KCMTextFormatDescriptionRect_Left is the left value for the rect as a 16-bit number.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionRect_Left
	KCMTextFormatDescriptionRect_Left string
	// KCMTextFormatDescriptionRect_Right is the right value for the rect as a 16-bit number.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionRect_Right
	KCMTextFormatDescriptionRect_Right string
	// KCMTextFormatDescriptionRect_Top is the top value for the rect as a 16-bit number.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionRect_Top
	KCMTextFormatDescriptionRect_Top string
	// KCMTextFormatDescriptionStyle_Ascent is the style value for the ascent.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionStyle_Ascent
	KCMTextFormatDescriptionStyle_Ascent string
	// KCMTextFormatDescriptionStyle_EndChar is the style value for the ending character.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionStyle_EndChar
	KCMTextFormatDescriptionStyle_EndChar string
	// KCMTextFormatDescriptionStyle_Font is the style value for font.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionStyle_Font
	KCMTextFormatDescriptionStyle_Font string
	// KCMTextFormatDescriptionStyle_FontFace is the style value for font face.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionStyle_FontFace
	KCMTextFormatDescriptionStyle_FontFace string
	// KCMTextFormatDescriptionStyle_FontSize is the style value for the font size.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionStyle_FontSize
	KCMTextFormatDescriptionStyle_FontSize string
	// KCMTextFormatDescriptionStyle_ForegroundColor is the style value for foreground color.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionStyle_ForegroundColor
	KCMTextFormatDescriptionStyle_ForegroundColor string
	// KCMTextFormatDescriptionStyle_Height is the style value for the height.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionStyle_Height
	KCMTextFormatDescriptionStyle_Height string
	// KCMTextFormatDescriptionStyle_StartChar is the style value for the starting character.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextFormatDescriptionStyle_StartChar
	KCMTextFormatDescriptionStyle_StartChar string
	// KCMTextMarkupAlignmentType_End is an alignment type that visually aligns the text at its ending side.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAlignmentType_End
	KCMTextMarkupAlignmentType_End string
	// KCMTextMarkupAlignmentType_Left is an alignment type that visually aligns text from left-to-right.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAlignmentType_Left
	KCMTextMarkupAlignmentType_Left string
	// KCMTextMarkupAlignmentType_Middle is an alignment type that visually aligns text in the center between its starting and ending sides.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAlignmentType_Middle
	KCMTextMarkupAlignmentType_Middle string
	// KCMTextMarkupAlignmentType_Right is an alignment type that visually aligns text from right-to-left.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAlignmentType_Right
	KCMTextMarkupAlignmentType_Right string
	// KCMTextMarkupAlignmentType_Start is an alignment type that visually aligns the text at its starting side.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAlignmentType_Start
	KCMTextMarkupAlignmentType_Start string
	// KCMTextMarkupAttribute_Alignment is the text alignment in the writing direction of the first line of text.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAttribute_Alignment
	KCMTextMarkupAttribute_Alignment string
	// KCMTextMarkupAttribute_BackgroundColorARGB is a background color for the text.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAttribute_BackgroundColorARGB
	KCMTextMarkupAttribute_BackgroundColorARGB string
	// KCMTextMarkupAttribute_BaseFontSizePercentageRelativeToVideoHeight is a base font size as a percentage of the video height.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAttribute_BaseFontSizePercentageRelativeToVideoHeight
	KCMTextMarkupAttribute_BaseFontSizePercentageRelativeToVideoHeight string
	// KCMTextMarkupAttribute_BoldStyle is a bold font style.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAttribute_BoldStyle
	KCMTextMarkupAttribute_BoldStyle string
	// KCMTextMarkupAttribute_CharacterBackgroundColorARGB is a background color for individual text characters.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAttribute_CharacterBackgroundColorARGB
	KCMTextMarkupAttribute_CharacterBackgroundColorARGB string
	// KCMTextMarkupAttribute_CharacterEdgeStyle is a style for character edges.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAttribute_CharacterEdgeStyle
	KCMTextMarkupAttribute_CharacterEdgeStyle string
	// KCMTextMarkupAttribute_FontFamilyName is a name of a font family.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAttribute_FontFamilyName
	KCMTextMarkupAttribute_FontFamilyName string
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAttribute_FontFamilyNameList
	KCMTextMarkupAttribute_FontFamilyNameList string
	// KCMTextMarkupAttribute_ForegroundColorARGB is a foreground color for the text.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAttribute_ForegroundColorARGB
	KCMTextMarkupAttribute_ForegroundColorARGB string
	// KCMTextMarkupAttribute_GenericFontFamilyName is a generic font family name identifier.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAttribute_GenericFontFamilyName
	KCMTextMarkupAttribute_GenericFontFamilyName string
	// KCMTextMarkupAttribute_ItalicStyle is an italic font style.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAttribute_ItalicStyle
	KCMTextMarkupAttribute_ItalicStyle string
	// KCMTextMarkupAttribute_OrthogonalLinePositionPercentageRelativeToWritingDirection is the placement of the first line in a block of text as a percentage in the direction orthogonal to the writing direction.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAttribute_OrthogonalLinePositionPercentageRelativeToWritingDirection
	KCMTextMarkupAttribute_OrthogonalLinePositionPercentageRelativeToWritingDirection string
	// KCMTextMarkupAttribute_RelativeFontSize is a font size as a percentage of the current default font size.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAttribute_RelativeFontSize
	KCMTextMarkupAttribute_RelativeFontSize string
	// KCMTextMarkupAttribute_TextPositionPercentageRelativeToWritingDirection is the placement of the block of text as a percentage in the writing direction.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAttribute_TextPositionPercentageRelativeToWritingDirection
	KCMTextMarkupAttribute_TextPositionPercentageRelativeToWritingDirection string
	// KCMTextMarkupAttribute_UnderlineStyle is an underline font style.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAttribute_UnderlineStyle
	KCMTextMarkupAttribute_UnderlineStyle string
	// KCMTextMarkupAttribute_VerticalLayout is the vertical layout of a text block.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAttribute_VerticalLayout
	KCMTextMarkupAttribute_VerticalLayout string
	// KCMTextMarkupAttribute_WritingDirectionSizePercentage is the width or height as a percentage of the bounding box that contains the text.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupAttribute_WritingDirectionSizePercentage
	KCMTextMarkupAttribute_WritingDirectionSizePercentage string
	// KCMTextMarkupCharacterEdgeStyle_Depressed is a depressed edge style.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupCharacterEdgeStyle_Depressed
	KCMTextMarkupCharacterEdgeStyle_Depressed string
	// KCMTextMarkupCharacterEdgeStyle_DropShadow is a drop shadow style.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupCharacterEdgeStyle_DropShadow
	KCMTextMarkupCharacterEdgeStyle_DropShadow string
	// KCMTextMarkupCharacterEdgeStyle_None is no edge style.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupCharacterEdgeStyle_None
	KCMTextMarkupCharacterEdgeStyle_None string
	// KCMTextMarkupCharacterEdgeStyle_Raised is a raised edge style.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupCharacterEdgeStyle_Raised
	KCMTextMarkupCharacterEdgeStyle_Raised string
	// KCMTextMarkupCharacterEdgeStyle_Uniform is a uniform border style.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupCharacterEdgeStyle_Uniform
	KCMTextMarkupCharacterEdgeStyle_Uniform string
	// KCMTextMarkupGenericFontName_Casual is a casual font.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupGenericFontName_Casual
	KCMTextMarkupGenericFontName_Casual string
	// KCMTextMarkupGenericFontName_Cursive is a cursive font.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupGenericFontName_Cursive
	KCMTextMarkupGenericFontName_Cursive string
	// KCMTextMarkupGenericFontName_Default is the default font.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupGenericFontName_Default
	KCMTextMarkupGenericFontName_Default string
	// KCMTextMarkupGenericFontName_Fantasy is a fantasy font.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupGenericFontName_Fantasy
	KCMTextMarkupGenericFontName_Fantasy string
	// KCMTextMarkupGenericFontName_Monospace is a monospaced font with or without serifs.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupGenericFontName_Monospace
	KCMTextMarkupGenericFontName_Monospace string
	// KCMTextMarkupGenericFontName_MonospaceSansSerif is a monospaced font without serifs.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupGenericFontName_MonospaceSansSerif
	KCMTextMarkupGenericFontName_MonospaceSansSerif string
	// KCMTextMarkupGenericFontName_MonospaceSerif is a monospaced font with serifs.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupGenericFontName_MonospaceSerif
	KCMTextMarkupGenericFontName_MonospaceSerif string
	// KCMTextMarkupGenericFontName_ProportionalSansSerif is a proportional font without serifs.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupGenericFontName_ProportionalSansSerif
	KCMTextMarkupGenericFontName_ProportionalSansSerif string
	// KCMTextMarkupGenericFontName_ProportionalSerif is a proportional font with serifs.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupGenericFontName_ProportionalSerif
	KCMTextMarkupGenericFontName_ProportionalSerif string
	// KCMTextMarkupGenericFontName_SansSerif is a font without serifs.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupGenericFontName_SansSerif
	KCMTextMarkupGenericFontName_SansSerif string
	// KCMTextMarkupGenericFontName_Serif is a font with serifs.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupGenericFontName_Serif
	KCMTextMarkupGenericFontName_Serif string
	// KCMTextMarkupGenericFontName_SmallCapital is a font with lowercase letters set as small capital letters.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextMarkupGenericFontName_SmallCapital
	KCMTextMarkupGenericFontName_SmallCapital string
	// KCMTextVerticalLayout_LeftToRight is add new vertical lines from left to right.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextVerticalLayout_LeftToRight
	KCMTextVerticalLayout_LeftToRight string
	// KCMTextVerticalLayout_RightToLeft is add new vertical lines from right to left.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTextVerticalLayout_RightToLeft
	KCMTextVerticalLayout_RightToLeft string
	// KCMTimeCodeFormatDescriptionExtension_SourceReferenceName is an extension that describes the source reference name.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTimeCodeFormatDescriptionExtension_SourceReferenceName
	KCMTimeCodeFormatDescriptionExtension_SourceReferenceName string
	// KCMTimeCodeFormatDescriptionKey_LangCode is a key that describes the language code.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTimeCodeFormatDescriptionKey_LangCode
	KCMTimeCodeFormatDescriptionKey_LangCode string
	// KCMTimeCodeFormatDescriptionKey_Value is a key that describes the value.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTimeCodeFormatDescriptionKey_Value
	KCMTimeCodeFormatDescriptionKey_Value string
	// KCMTimeEpochKey is a dictionary key for a time epoch.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTimeEpochKey
	KCMTimeEpochKey string
	// KCMTimeFlagsKey is a dictionary key for time flags.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTimeFlagsKey
	KCMTimeFlagsKey string
	// KCMTimeMappingSourceKey is a dictionary key for a source time range.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTimeMappingSourceKey
	KCMTimeMappingSourceKey string
	// KCMTimeMappingTargetKey is a dictionary key for a target time range.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTimeMappingTargetKey
	KCMTimeMappingTargetKey string
	// KCMTimeRangeDurationKey is the key for the timescale of a time range.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTimeRangeDurationKey
	KCMTimeRangeDurationKey string
	// KCMTimeRangeStartKey is the key for the start field of a time range.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTimeRangeStartKey
	KCMTimeRangeStartKey string
	// KCMTimeScaleKey is a dictionary key for a timescale.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTimeScaleKey
	KCMTimeScaleKey string
	// KCMTimeValueKey is a dictionary key for a time value.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTimeValueKey
	KCMTimeValueKey string
	// KCMTimebaseNotificationKey_EventTime is a notification that a timebase posts after a discontinuous time jump.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTimebaseNotificationKey_EventTime
	KCMTimebaseNotificationKey_EventTime string
	// KCMTimebaseNotification_EffectiveRateChanged is a notification that posts by a timebase after a change in effective rate.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTimebaseNotification_EffectiveRateChanged
	KCMTimebaseNotification_EffectiveRateChanged string
	// KCMTimebaseNotification_TimeJumped is a notification that posts by a timebase after a discontinuous time jump.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTimebaseNotification_TimeJumped
	KCMTimebaseNotification_TimeJumped string
)

var (
	// KCMImageDescriptionFlavor_3GPFamily is an image description that selects the 3GP family sample format.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/CMImageDescriptionFlavor/mobile3GPFamily
	KCMImageDescriptionFlavor_3GPFamily CMImageDescriptionFlavor
	// See: https://developer.apple.com/documentation/CoreMedia/CMImageDescriptionFlavor/isoFamily
	KCMImageDescriptionFlavor_ISOFamily CMImageDescriptionFlavor
	// See: https://developer.apple.com/documentation/CoreMedia/CMImageDescriptionFlavor/isoFamilyWithAppleExtensions
	KCMImageDescriptionFlavor_ISOFamilyWithAppleExtensions CMImageDescriptionFlavor
	// KCMImageDescriptionFlavor_QuickTimeMovie is an image description that selects the QuickTime format.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/CMImageDescriptionFlavor/quickTimeMovie
	KCMImageDescriptionFlavor_QuickTimeMovie CMImageDescriptionFlavor
)

var (
	// KCMSoundDescriptionFlavor_3GPFamily is a sound description that selects the 3GP family sample format.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/CMSoundDescriptionFlavor/mobile3GPFamily
	KCMSoundDescriptionFlavor_3GPFamily CMSoundDescriptionFlavor
	// KCMSoundDescriptionFlavor_ISOFamily is a sound description that selects the ISO family sample format.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/CMSoundDescriptionFlavor/isoFamily
	KCMSoundDescriptionFlavor_ISOFamily CMSoundDescriptionFlavor
	// KCMSoundDescriptionFlavor_QuickTimeMovie is a sound description that selects the QuickTime movie format.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/CMSoundDescriptionFlavor/quickTimeMovie
	KCMSoundDescriptionFlavor_QuickTimeMovie CMSoundDescriptionFlavor
	// KCMSoundDescriptionFlavor_QuickTimeMovieV2 is a sound description that selects the second version of the QuickTime movie format.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/CMSoundDescriptionFlavor/quickTimeMovieV2
	KCMSoundDescriptionFlavor_QuickTimeMovieV2 CMSoundDescriptionFlavor
)

var (
	// KCMTagInvalid is a constant representing an invalid tag.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagInvalid
	KCMTagInvalid CMTag
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagMediaSubTypeMebx
	KCMTagMediaSubTypeMebx CMTag
	// KCMTagMediaTypeAudio is a value for associating a tag’s media type with audio.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagMediaTypeAudio
	KCMTagMediaTypeAudio CMTag
	// KCMTagMediaTypeMetadata is a value for associating a tag’s media type with metadata.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagMediaTypeMetadata
	KCMTagMediaTypeMetadata CMTag
	// KCMTagMediaTypeVideo is a value for associating a tag’s media type with video.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagMediaTypeVideo
	KCMTagMediaTypeVideo CMTag
	// KCMTagPackingTypeNone is a frame-packing tag value for video without packed frames.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagPackingTypeNone
	KCMTagPackingTypeNone CMTag
	// KCMTagPackingTypeOverUnder is a tag stating that associated video has packed frames with a left eye image on the top and right eye image on the bottom.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagPackingTypeOverUnder
	KCMTagPackingTypeOverUnder CMTag
	// KCMTagPackingTypeSideBySide is a tag stating that associated video has packed frames with a left eye image on the left and right eye image on the right.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagPackingTypeSideBySide
	KCMTagPackingTypeSideBySide CMTag
	// KCMTagProjectionTypeEquirectangular is a value for projection tags indicating that display is on a 360 degree equirectangular projection.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagProjectionTypeEquirectangular
	KCMTagProjectionTypeEquirectangular CMTag
	// KCMTagProjectionTypeFisheye is video content displays as a fisheye projection.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagProjectionTypeFisheye
	KCMTagProjectionTypeFisheye CMTag
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagProjectionTypeHalfEquirectangular
	KCMTagProjectionTypeHalfEquirectangular CMTag
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagProjectionTypeParametricImmersive
	KCMTagProjectionTypeParametricImmersive CMTag
	// KCMTagProjectionTypeRectangular is a value for projection tags indicating that display is on a flat rectangular surface.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagProjectionTypeRectangular
	KCMTagProjectionTypeRectangular CMTag
	// KCMTagStereoInterpretationOrderReversed is a value for a stereo interpretation tag indicating the video data for the left and right eyes are reversed.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagStereoInterpretationOrderReversed
	KCMTagStereoInterpretationOrderReversed CMTag
	// KCMTagStereoLeftAndRightEye is a value for a stereo tag indicating the video track has left and right eye layers.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagStereoLeftAndRightEye
	KCMTagStereoLeftAndRightEye CMTag
	// KCMTagStereoLeftEye is a value for a stereo tag indicating the video track has a left eye layer.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagStereoLeftEye
	KCMTagStereoLeftEye CMTag
	// KCMTagStereoNone is a value for a stereo tag indicating the video track has no eye layer data.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagStereoNone
	KCMTagStereoNone CMTag
	// KCMTagStereoRightEye is a value for a stereo tag indicating the video track has a right eye layer.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/kCMTagStereoRightEye
	KCMTagStereoRightEye CMTag
)

var (
	// KCMTimeIndefinite is a value that represents an indefinite time.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/CMTime/indefinite
	KCMTimeIndefinite CMTime
	// KCMTimeInvalid is a value that represents an invalid time.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
	KCMTimeInvalid CMTime
	// KCMTimeNegativeInfinity is a value that represents negative infinity.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/CMTime/negativeInfinity
	KCMTimeNegativeInfinity CMTime
	// KCMTimePositiveInfinity is a value that represents positive infinity.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/CMTime/positiveInfinity
	KCMTimePositiveInfinity CMTime
	// KCMTimeZero is a value that represents time zero.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
	KCMTimeZero CMTime
)

var (
	// KCMTimeMappingInvalid is an invalid time mapping.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMapping/invalid
	KCMTimeMappingInvalid CMTimeMapping
)

var (
	// KCMTimeRangeInvalid is a constant for generating an invalid time range.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRange/invalid
	KCMTimeRangeInvalid CMTimeRange
	// KCMTimeRangeZero is a constant for generating an empty time range at zero.
	//
	// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRange/zero
	KCMTimeRangeZero CMTimeRange
)

var (
	// See: https://developer.apple.com/documentation/CoreMedia/CMSampleTimingInfo/invalid
	KCMTimingInfoInvalid CMSampleTimingInfo
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionAlphaChannelMode_PremultipliedAlpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionAlphaChannelMode_PremultipliedAlpha = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionAlphaChannelMode_StraightAlpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionAlphaChannelMode_StraightAlpha = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionCameraCalibrationExtrinsicOriginSource_StereoCameraSystemBaseline"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionCameraCalibrationExtrinsicOriginSource_StereoCameraSystemBaseline = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionCameraCalibrationLensAlgorithmKind_ParametricLens"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionCameraCalibrationLensAlgorithmKind_ParametricLens = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionCameraCalibrationLensDomain_Color"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionCameraCalibrationLensDomain_Color = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionCameraCalibrationLensRole_Left"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionCameraCalibrationLensRole_Left = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionCameraCalibrationLensRole_Mono"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionCameraCalibrationLensRole_Mono = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionCameraCalibrationLensRole_Right"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionCameraCalibrationLensRole_Right = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionCameraCalibration_ExtrinsicOrientationQuaternion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionCameraCalibration_ExtrinsicOrientationQuaternion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionCameraCalibration_ExtrinsicOriginSource"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionCameraCalibration_ExtrinsicOriginSource = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionCameraCalibration_IntrinsicMatrix"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionCameraCalibration_IntrinsicMatrix = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionCameraCalibration_IntrinsicMatrixProjectionOffset"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionCameraCalibration_IntrinsicMatrixProjectionOffset = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionCameraCalibration_IntrinsicMatrixReferenceDimensions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionCameraCalibration_IntrinsicMatrixReferenceDimensions = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionCameraCalibration_LensAlgorithmKind"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionCameraCalibration_LensAlgorithmKind = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionCameraCalibration_LensDistortions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionCameraCalibration_LensDistortions = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionCameraCalibration_LensDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionCameraCalibration_LensDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionCameraCalibration_LensFrameAdjustmentsPolynomialX"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionCameraCalibration_LensFrameAdjustmentsPolynomialX = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionCameraCalibration_LensFrameAdjustmentsPolynomialY"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionCameraCalibration_LensFrameAdjustmentsPolynomialY = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionCameraCalibration_LensIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionCameraCalibration_LensIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionCameraCalibration_LensRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionCameraCalibration_LensRole = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionCameraCalibration_RadialAngleLimit"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionCameraCalibration_RadialAngleLimit = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionChromaLocation_Bottom"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionChromaLocation_Bottom = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionChromaLocation_BottomLeft"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionChromaLocation_BottomLeft = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionChromaLocation_Center"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionChromaLocation_Center = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionChromaLocation_DV420"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionChromaLocation_DV420 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionChromaLocation_Left"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionChromaLocation_Left = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionChromaLocation_Top"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionChromaLocation_Top = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionChromaLocation_TopLeft"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionChromaLocation_TopLeft = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionColorPrimaries_DCI_P3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionColorPrimaries_DCI_P3 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionColorPrimaries_EBU_3213"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionColorPrimaries_EBU_3213 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionColorPrimaries_ITU_R_2020"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionColorPrimaries_ITU_R_2020 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionColorPrimaries_ITU_R_709_2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionColorPrimaries_ITU_R_709_2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionColorPrimaries_P22"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionColorPrimaries_P22 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionColorPrimaries_P3_D65"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionColorPrimaries_P3_D65 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionColorPrimaries_SMPTE_C"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionColorPrimaries_SMPTE_C = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionConformsToMPEG2VideoProfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionConformsToMPEG2VideoProfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtensionKey_MetadataKeyTable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtensionKey_MetadataKeyTable = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_AlphaChannelMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_AlphaChannelMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_AlternativeTransferCharacteristics"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_AlternativeTransferCharacteristics = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_AmbientViewingEnvironment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_AmbientViewingEnvironment = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_AuxiliaryTypeInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_AuxiliaryTypeInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_BitsPerComponent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_BitsPerComponent = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_BytesPerRow"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_BytesPerRow = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_CameraCalibrationDataLensCollection"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_CameraCalibrationDataLensCollection = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_ChromaLocationBottomField"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_ChromaLocationBottomField = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_ChromaLocationTopField"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_ChromaLocationTopField = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_CleanAperture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_CleanAperture = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_ColorPrimaries"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_ColorPrimaries = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_ContainsAlphaChannel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_ContainsAlphaChannel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_ContentColorVolume"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_ContentColorVolume = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_ContentLightLevelInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_ContentLightLevelInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_ConvertedFromExternalSphericalTags"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_ConvertedFromExternalSphericalTags = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_Depth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_Depth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_FieldCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_FieldCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_FieldDetail"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_FieldDetail = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_FormatName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_FormatName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_FullRangeVideo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_FullRangeVideo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_GammaLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_GammaLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_HasAdditionalViews"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_HasAdditionalViews = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_HasLeftStereoEyeView"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_HasLeftStereoEyeView = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_HasRightStereoEyeView"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_HasRightStereoEyeView = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_HeroEye"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_HeroEye = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_HorizontalDisparityAdjustment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_HorizontalDisparityAdjustment = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_HorizontalFieldOfView"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_HorizontalFieldOfView = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_ICCProfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_ICCProfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_LogTransferFunction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_LogTransferFunction = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_MasteringDisplayColorVolume"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_MasteringDisplayColorVolume = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_OriginalCompressionSettings"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_OriginalCompressionSettings = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_PixelAspectRatio"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_PixelAspectRatio = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_ProjectionKind"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_ProjectionKind = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_ProtectedContentOriginalFormat"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_ProtectedContentOriginalFormat = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_RevisionLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_RevisionLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_SampleDescriptionExtensionAtoms"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_SampleDescriptionExtensionAtoms = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_SpatialQuality"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_SpatialQuality = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_StereoCameraBaseline"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_StereoCameraBaseline = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_TemporalQuality"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_TemporalQuality = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_TransferFunction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_TransferFunction = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_Vendor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_Vendor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_VerbatimISOSampleEntry"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_VerbatimISOSampleEntry = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_VerbatimImageDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_VerbatimImageDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_VerbatimSampleDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_VerbatimSampleDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_Version"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_Version = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_ViewPackingKind"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_ViewPackingKind = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionExtension_YCbCrMatrix"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionExtension_YCbCrMatrix = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionFieldDetail_SpatialFirstLineEarly"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionFieldDetail_SpatialFirstLineEarly = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionFieldDetail_SpatialFirstLineLate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionFieldDetail_SpatialFirstLineLate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionFieldDetail_TemporalBottomFirst"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionFieldDetail_TemporalBottomFirst = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionFieldDetail_TemporalTopFirst"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionFieldDetail_TemporalTopFirst = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionHeroEye_Left"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionHeroEye_Left = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionHeroEye_Right"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionHeroEye_Right = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionKey_CleanApertureHeight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionKey_CleanApertureHeight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionKey_CleanApertureHeightRational"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionKey_CleanApertureHeightRational = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionKey_CleanApertureHorizontalOffset"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionKey_CleanApertureHorizontalOffset = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionKey_CleanApertureHorizontalOffsetRational"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionKey_CleanApertureHorizontalOffsetRational = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionKey_CleanApertureVerticalOffset"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionKey_CleanApertureVerticalOffset = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionKey_CleanApertureVerticalOffsetRational"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionKey_CleanApertureVerticalOffsetRational = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionKey_CleanApertureWidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionKey_CleanApertureWidth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionKey_CleanApertureWidthRational"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionKey_CleanApertureWidthRational = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionKey_PixelAspectRatioHorizontalSpacing"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionKey_PixelAspectRatioHorizontalSpacing = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionKey_PixelAspectRatioVerticalSpacing"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionKey_PixelAspectRatioVerticalSpacing = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionLogTransferFunction_AppleLog"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionLogTransferFunction_AppleLog = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionProjectionKind_AppleImmersiveVideo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionProjectionKind_AppleImmersiveVideo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionProjectionKind_Equirectangular"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionProjectionKind_Equirectangular = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionProjectionKind_HalfEquirectangular"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionProjectionKind_HalfEquirectangular = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionProjectionKind_ParametricImmersive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionProjectionKind_ParametricImmersive = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionProjectionKind_Rectilinear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionProjectionKind_Rectilinear = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionTransferFunction_ITU_R_2020"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionTransferFunction_ITU_R_2020 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionTransferFunction_ITU_R_2100_HLG"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionTransferFunction_ITU_R_2100_HLG = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionTransferFunction_ITU_R_709_2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionTransferFunction_ITU_R_709_2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionTransferFunction_Linear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionTransferFunction_Linear = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionTransferFunction_SMPTE_240M_1995"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionTransferFunction_SMPTE_240M_1995 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionTransferFunction_SMPTE_ST_2084_PQ"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionTransferFunction_SMPTE_ST_2084_PQ = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionTransferFunction_SMPTE_ST_428_1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionTransferFunction_SMPTE_ST_428_1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionTransferFunction_UseGamma"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionTransferFunction_UseGamma = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionTransferFunction_sRGB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionTransferFunction_sRGB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionVendor_Apple"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionVendor_Apple = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionViewPackingKind_OverUnder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionViewPackingKind_OverUnder = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionViewPackingKind_SideBySide"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionViewPackingKind_SideBySide = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionYCbCrMatrix_ITU_R_2020"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionYCbCrMatrix_ITU_R_2020 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionYCbCrMatrix_ITU_R_601_4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionYCbCrMatrix_ITU_R_601_4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionYCbCrMatrix_ITU_R_709_2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionYCbCrMatrix_ITU_R_709_2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMFormatDescriptionYCbCrMatrix_SMPTE_240M_1995"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMFormatDescriptionYCbCrMatrix_SMPTE_240M_1995 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMHEVCTemporalLevelInfoKey_ConstraintIndicatorFlags"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMHEVCTemporalLevelInfoKey_ConstraintIndicatorFlags = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMHEVCTemporalLevelInfoKey_LevelIndex"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMHEVCTemporalLevelInfoKey_LevelIndex = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMHEVCTemporalLevelInfoKey_ProfileCompatibilityFlags"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMHEVCTemporalLevelInfoKey_ProfileCompatibilityFlags = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMHEVCTemporalLevelInfoKey_ProfileIndex"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMHEVCTemporalLevelInfoKey_ProfileIndex = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMHEVCTemporalLevelInfoKey_ProfileSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMHEVCTemporalLevelInfoKey_ProfileSpace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMHEVCTemporalLevelInfoKey_TemporalLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMHEVCTemporalLevelInfoKey_TemporalLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMHEVCTemporalLevelInfoKey_TierFlag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMHEVCTemporalLevelInfoKey_TierFlag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMImageDescriptionFlavor_3GPFamily"); err == nil && ptr != 0 {
		KCMImageDescriptionFlavor_3GPFamily = *(*CMImageDescriptionFlavor)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMImageDescriptionFlavor_ISOFamily"); err == nil && ptr != 0 {
		KCMImageDescriptionFlavor_ISOFamily = *(*CMImageDescriptionFlavor)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMImageDescriptionFlavor_ISOFamilyWithAppleExtensions"); err == nil && ptr != 0 {
		KCMImageDescriptionFlavor_ISOFamilyWithAppleExtensions = *(*CMImageDescriptionFlavor)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMImageDescriptionFlavor_QuickTimeMovie"); err == nil && ptr != 0 {
		KCMImageDescriptionFlavor_QuickTimeMovie = *(*CMImageDescriptionFlavor)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMemoryPoolOption_AgeOutPeriod"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMemoryPoolOption_AgeOutPeriod = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_AffineTransformF64"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_AffineTransformF64 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_BMP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_BMP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_DimensionsF32"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_DimensionsF32 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_ExtendedRasterRectangleValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_ExtendedRasterRectangleValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_Float32"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_Float32 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_Float64"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_Float64 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_GIF"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_GIF = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_JPEG"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_JPEG = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_JSON"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_JSON = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_PNG"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_PNG = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_PerspectiveTransformF64"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_PerspectiveTransformF64 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_PointF32"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_PointF32 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_PolygonF32"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_PolygonF32 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_PolylineF32"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_PolylineF32 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_RasterRectangleValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_RasterRectangleValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_RawData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_RawData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_RectF32"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_RectF32 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_SInt16"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_SInt16 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_SInt32"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_SInt32 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_SInt64"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_SInt64 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_SInt8"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_SInt8 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_UInt16"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_UInt16 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_UInt32"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_UInt32 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_UInt64"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_UInt64 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_UInt8"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_UInt8 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_UTF16"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_UTF16 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataBaseDataType_UTF8"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataBaseDataType_UTF8 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataDataType_QuickTimeMetadataDirection"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataDataType_QuickTimeMetadataDirection = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataDataType_QuickTimeMetadataLocation_ISO6709"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataDataType_QuickTimeMetadataLocation_ISO6709 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataDataType_QuickTimeMetadataMilliLux"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataDataType_QuickTimeMetadataMilliLux = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataDataType_QuickTimeMetadataUUID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataDataType_QuickTimeMetadataUUID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataFormatDescriptionKey_ConformingDataTypes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataFormatDescriptionKey_ConformingDataTypes = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataFormatDescriptionKey_DataType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataFormatDescriptionKey_DataType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataFormatDescriptionKey_DataTypeNamespace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataFormatDescriptionKey_DataTypeNamespace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataFormatDescriptionKey_LanguageTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataFormatDescriptionKey_LanguageTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataFormatDescriptionKey_LocalID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataFormatDescriptionKey_LocalID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataFormatDescriptionKey_Namespace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataFormatDescriptionKey_Namespace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataFormatDescriptionKey_SetupData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataFormatDescriptionKey_SetupData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataFormatDescriptionKey_StructuralDependency"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataFormatDescriptionKey_StructuralDependency = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataFormatDescriptionKey_Value"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataFormatDescriptionKey_Value = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataFormatDescriptionMetadataSpecificationKey_DataType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataFormatDescriptionMetadataSpecificationKey_DataType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataFormatDescriptionMetadataSpecificationKey_ExtendedLanguageTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataFormatDescriptionMetadataSpecificationKey_ExtendedLanguageTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataFormatDescriptionMetadataSpecificationKey_Identifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataFormatDescriptionMetadataSpecificationKey_Identifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataFormatDescriptionMetadataSpecificationKey_SetupData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataFormatDescriptionMetadataSpecificationKey_SetupData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataFormatDescriptionMetadataSpecificationKey_StructuralDependency"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataFormatDescriptionMetadataSpecificationKey_StructuralDependency = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataFormatDescription_StructuralDependencyKey_DependencyIsInvalidFlag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataFormatDescription_StructuralDependencyKey_DependencyIsInvalidFlag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataIdentifier_QuickTimeMetadataDirection_Facing"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataIdentifier_QuickTimeMetadataDirection_Facing = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataIdentifier_QuickTimeMetadataDisplayMaskRectangleMono"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataIdentifier_QuickTimeMetadataDisplayMaskRectangleMono = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataIdentifier_QuickTimeMetadataDisplayMaskRectangleStereoLeft"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataIdentifier_QuickTimeMetadataDisplayMaskRectangleStereoLeft = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataIdentifier_QuickTimeMetadataDisplayMaskRectangleStereoRight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataIdentifier_QuickTimeMetadataDisplayMaskRectangleStereoRight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataIdentifier_QuickTimeMetadataLivePhotoStillImageTransform"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataIdentifier_QuickTimeMetadataLivePhotoStillImageTransform = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataIdentifier_QuickTimeMetadataLivePhotoStillImageTransformReferenceDimensions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataIdentifier_QuickTimeMetadataLivePhotoStillImageTransformReferenceDimensions = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataIdentifier_QuickTimeMetadataLocation_ISO6709"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataIdentifier_QuickTimeMetadataLocation_ISO6709 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataIdentifier_QuickTimeMetadataPreferredAffineTransform"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataIdentifier_QuickTimeMetadataPreferredAffineTransform = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataIdentifier_QuickTimeMetadataPresentationImmersiveMedia"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataIdentifier_QuickTimeMetadataPresentationImmersiveMedia = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataIdentifier_QuickTimeMetadataSceneIlluminance"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataIdentifier_QuickTimeMetadataSceneIlluminance = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataIdentifier_QuickTimeMetadataSegmentIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataIdentifier_QuickTimeMetadataSegmentIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataIdentifier_QuickTimeMetadataSpatialAudioMix"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataIdentifier_QuickTimeMetadataSpatialAudioMix = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataIdentifier_QuickTimeMetadataVideoOrientation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataIdentifier_QuickTimeMetadataVideoOrientation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataKeySpace_HLSDateRange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataKeySpace_HLSDateRange = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataKeySpace_ID3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataKeySpace_ID3 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataKeySpace_ISOUserData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataKeySpace_ISOUserData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataKeySpace_Icy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataKeySpace_Icy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataKeySpace_QuickTimeMetadata"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataKeySpace_QuickTimeMetadata = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataKeySpace_QuickTimeUserData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataKeySpace_QuickTimeUserData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMetadataKeySpace_iTunes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMetadataKeySpace_iTunes = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleAttachmentKey_AudioIndependentSampleDecoderRefreshCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleAttachmentKey_AudioIndependentSampleDecoderRefreshCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleAttachmentKey_CryptorSubsampleAuxiliaryData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleAttachmentKey_CryptorSubsampleAuxiliaryData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleAttachmentKey_DependsOnOthers"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleAttachmentKey_DependsOnOthers = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleAttachmentKey_DisplayImmediately"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleAttachmentKey_DisplayImmediately = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleAttachmentKey_DoNotDisplay"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleAttachmentKey_DoNotDisplay = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleAttachmentKey_EarlierDisplayTimesAllowed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleAttachmentKey_EarlierDisplayTimesAllowed = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleAttachmentKey_HDR10PlusPerFrameData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleAttachmentKey_HDR10PlusPerFrameData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleAttachmentKey_HEVCStepwiseTemporalSubLayerAccess"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleAttachmentKey_HEVCStepwiseTemporalSubLayerAccess = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleAttachmentKey_HEVCSyncSampleNALUnitType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleAttachmentKey_HEVCSyncSampleNALUnitType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleAttachmentKey_HEVCTemporalLevelInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleAttachmentKey_HEVCTemporalLevelInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleAttachmentKey_HEVCTemporalSubLayerAccess"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleAttachmentKey_HEVCTemporalSubLayerAccess = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleAttachmentKey_HasRedundantCoding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleAttachmentKey_HasRedundantCoding = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleAttachmentKey_IsDependedOnByOthers"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleAttachmentKey_IsDependedOnByOthers = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleAttachmentKey_NotSync"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleAttachmentKey_NotSync = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleAttachmentKey_PartialSync"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleAttachmentKey_PartialSync = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleAttachmentKey_PostDecodeProcessingMetadata"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleAttachmentKey_PostDecodeProcessingMetadata = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_CameraIntrinsicMatrix"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_CameraIntrinsicMatrix = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_DisplayEmptyMediaImmediately"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_DisplayEmptyMediaImmediately = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_DrainAfterDecoding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_DrainAfterDecoding = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_DroppedFrameReason"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_DroppedFrameReason = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_DroppedFrameReasonInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_DroppedFrameReasonInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_EmptyMedia"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_EmptyMedia = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_EndsPreviousSampleDuration"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_EndsPreviousSampleDuration = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_FillDiscontinuitiesWithSilence"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_FillDiscontinuitiesWithSilence = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_ForceKeyFrame"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_ForceKeyFrame = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_GradualDecoderRefresh"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_GradualDecoderRefresh = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_PermanentEmptyMedia"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_PermanentEmptyMedia = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_PostNotificationWhenConsumed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_PostNotificationWhenConsumed = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_ResetDecoderBeforeDecoding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_ResetDecoderBeforeDecoding = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_ResumeOutput"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_ResumeOutput = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_Reverse"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_Reverse = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_SampleReferenceByteOffset"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_SampleReferenceByteOffset = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_SampleReferenceURL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_SampleReferenceURL = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_SpeedMultiplier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_SpeedMultiplier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_StillImageLensStabilizationInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_StillImageLensStabilizationInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_TransitionID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_TransitionID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_TrimDurationAtEnd"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_TrimDurationAtEnd = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferAttachmentKey_TrimDurationAtStart"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferAttachmentKey_TrimDurationAtStart = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferConduitNotificationParameter_MaxUpcomingOutputPTS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferConduitNotificationParameter_MaxUpcomingOutputPTS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferConduitNotificationParameter_MinUpcomingOutputPTS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferConduitNotificationParameter_MinUpcomingOutputPTS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferConduitNotificationParameter_ResumeTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferConduitNotificationParameter_ResumeTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferConduitNotificationParameter_UpcomingOutputPTSRangeMayOverlapQueuedOutputPTSRange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferConduitNotificationParameter_UpcomingOutputPTSRangeMayOverlapQueuedOutputPTSRange = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferConduitNotification_InhibitOutputUntil"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferConduitNotification_InhibitOutputUntil = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferConduitNotification_ResetOutput"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferConduitNotification_ResetOutput = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferConduitNotification_UpcomingOutputPTSRangeChanged"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferConduitNotification_UpcomingOutputPTSRangeChanged = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferConsumerNotification_BufferConsumed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferConsumerNotification_BufferConsumed = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferDroppedFrameReasonInfo_CameraModeSwitch"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferDroppedFrameReasonInfo_CameraModeSwitch = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferDroppedFrameReason_Discontinuity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferDroppedFrameReason_Discontinuity = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferDroppedFrameReason_FrameWasLate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferDroppedFrameReason_FrameWasLate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferDroppedFrameReason_OutOfBuffers"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferDroppedFrameReason_OutOfBuffers = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferLensStabilizationInfo_Active"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferLensStabilizationInfo_Active = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferLensStabilizationInfo_Off"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferLensStabilizationInfo_Off = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferLensStabilizationInfo_OutOfRange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferLensStabilizationInfo_OutOfRange = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferLensStabilizationInfo_Unavailable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferLensStabilizationInfo_Unavailable = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferNotificationParameter_OSStatus"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferNotificationParameter_OSStatus = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferNotification_DataBecameReady"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferNotification_DataBecameReady = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSampleBufferNotification_DataFailed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSampleBufferNotification_DataFailed = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSoundDescriptionFlavor_3GPFamily"); err == nil && ptr != 0 {
		KCMSoundDescriptionFlavor_3GPFamily = *(*CMSoundDescriptionFlavor)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSoundDescriptionFlavor_ISOFamily"); err == nil && ptr != 0 {
		KCMSoundDescriptionFlavor_ISOFamily = *(*CMSoundDescriptionFlavor)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSoundDescriptionFlavor_QuickTimeMovie"); err == nil && ptr != 0 {
		KCMSoundDescriptionFlavor_QuickTimeMovie = *(*CMSoundDescriptionFlavor)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSoundDescriptionFlavor_QuickTimeMovieV2"); err == nil && ptr != 0 {
		KCMSoundDescriptionFlavor_QuickTimeMovieV2 = *(*CMSoundDescriptionFlavor)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagCategoryKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTagCategoryKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagCollectionTagsArrayKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTagCollectionTagsArrayKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagDataTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTagDataTypeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagInvalid"); err == nil && ptr != 0 {
		KCMTagInvalid = *(*CMTag)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagMediaSubTypeMebx"); err == nil && ptr != 0 {
		KCMTagMediaSubTypeMebx = *(*CMTag)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagMediaTypeAudio"); err == nil && ptr != 0 {
		KCMTagMediaTypeAudio = *(*CMTag)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagMediaTypeMetadata"); err == nil && ptr != 0 {
		KCMTagMediaTypeMetadata = *(*CMTag)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagMediaTypeVideo"); err == nil && ptr != 0 {
		KCMTagMediaTypeVideo = *(*CMTag)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagPackingTypeNone"); err == nil && ptr != 0 {
		KCMTagPackingTypeNone = *(*CMTag)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagPackingTypeOverUnder"); err == nil && ptr != 0 {
		KCMTagPackingTypeOverUnder = *(*CMTag)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagPackingTypeSideBySide"); err == nil && ptr != 0 {
		KCMTagPackingTypeSideBySide = *(*CMTag)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagProjectionTypeEquirectangular"); err == nil && ptr != 0 {
		KCMTagProjectionTypeEquirectangular = *(*CMTag)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagProjectionTypeFisheye"); err == nil && ptr != 0 {
		KCMTagProjectionTypeFisheye = *(*CMTag)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagProjectionTypeHalfEquirectangular"); err == nil && ptr != 0 {
		KCMTagProjectionTypeHalfEquirectangular = *(*CMTag)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagProjectionTypeParametricImmersive"); err == nil && ptr != 0 {
		KCMTagProjectionTypeParametricImmersive = *(*CMTag)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagProjectionTypeRectangular"); err == nil && ptr != 0 {
		KCMTagProjectionTypeRectangular = *(*CMTag)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagStereoInterpretationOrderReversed"); err == nil && ptr != 0 {
		KCMTagStereoInterpretationOrderReversed = *(*CMTag)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagStereoLeftAndRightEye"); err == nil && ptr != 0 {
		KCMTagStereoLeftAndRightEye = *(*CMTag)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagStereoLeftEye"); err == nil && ptr != 0 {
		KCMTagStereoLeftEye = *(*CMTag)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagStereoNone"); err == nil && ptr != 0 {
		KCMTagStereoNone = *(*CMTag)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagStereoRightEye"); err == nil && ptr != 0 {
		KCMTagStereoRightEye = *(*CMTag)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTagValueKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTagValueKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionColor_Alpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionColor_Alpha = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionColor_Blue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionColor_Blue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionColor_Green"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionColor_Green = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionColor_Red"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionColor_Red = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionExtension_BackgroundColor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionExtension_BackgroundColor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionExtension_DefaultFontName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionExtension_DefaultFontName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionExtension_DefaultStyle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionExtension_DefaultStyle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionExtension_DefaultTextBox"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionExtension_DefaultTextBox = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionExtension_DisplayFlags"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionExtension_DisplayFlags = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionExtension_FontTable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionExtension_FontTable = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionExtension_HorizontalJustification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionExtension_HorizontalJustification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionExtension_TextJustification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionExtension_TextJustification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionExtension_VerticalJustification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionExtension_VerticalJustification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionRect_Bottom"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionRect_Bottom = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionRect_Left"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionRect_Left = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionRect_Right"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionRect_Right = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionRect_Top"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionRect_Top = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionStyle_Ascent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionStyle_Ascent = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionStyle_EndChar"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionStyle_EndChar = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionStyle_Font"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionStyle_Font = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionStyle_FontFace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionStyle_FontFace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionStyle_FontSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionStyle_FontSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionStyle_ForegroundColor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionStyle_ForegroundColor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionStyle_Height"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionStyle_Height = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextFormatDescriptionStyle_StartChar"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextFormatDescriptionStyle_StartChar = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAlignmentType_End"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAlignmentType_End = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAlignmentType_Left"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAlignmentType_Left = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAlignmentType_Middle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAlignmentType_Middle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAlignmentType_Right"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAlignmentType_Right = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAlignmentType_Start"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAlignmentType_Start = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAttribute_Alignment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAttribute_Alignment = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAttribute_BackgroundColorARGB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAttribute_BackgroundColorARGB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAttribute_BaseFontSizePercentageRelativeToVideoHeight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAttribute_BaseFontSizePercentageRelativeToVideoHeight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAttribute_BoldStyle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAttribute_BoldStyle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAttribute_CharacterBackgroundColorARGB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAttribute_CharacterBackgroundColorARGB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAttribute_CharacterEdgeStyle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAttribute_CharacterEdgeStyle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAttribute_FontFamilyName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAttribute_FontFamilyName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAttribute_FontFamilyNameList"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAttribute_FontFamilyNameList = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAttribute_ForegroundColorARGB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAttribute_ForegroundColorARGB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAttribute_GenericFontFamilyName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAttribute_GenericFontFamilyName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAttribute_ItalicStyle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAttribute_ItalicStyle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAttribute_OrthogonalLinePositionPercentageRelativeToWritingDirection"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAttribute_OrthogonalLinePositionPercentageRelativeToWritingDirection = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAttribute_RelativeFontSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAttribute_RelativeFontSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAttribute_TextPositionPercentageRelativeToWritingDirection"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAttribute_TextPositionPercentageRelativeToWritingDirection = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAttribute_UnderlineStyle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAttribute_UnderlineStyle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAttribute_VerticalLayout"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAttribute_VerticalLayout = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupAttribute_WritingDirectionSizePercentage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupAttribute_WritingDirectionSizePercentage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupCharacterEdgeStyle_Depressed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupCharacterEdgeStyle_Depressed = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupCharacterEdgeStyle_DropShadow"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupCharacterEdgeStyle_DropShadow = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupCharacterEdgeStyle_None"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupCharacterEdgeStyle_None = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupCharacterEdgeStyle_Raised"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupCharacterEdgeStyle_Raised = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupCharacterEdgeStyle_Uniform"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupCharacterEdgeStyle_Uniform = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupGenericFontName_Casual"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupGenericFontName_Casual = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupGenericFontName_Cursive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupGenericFontName_Cursive = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupGenericFontName_Default"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupGenericFontName_Default = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupGenericFontName_Fantasy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupGenericFontName_Fantasy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupGenericFontName_Monospace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupGenericFontName_Monospace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupGenericFontName_MonospaceSansSerif"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupGenericFontName_MonospaceSansSerif = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupGenericFontName_MonospaceSerif"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupGenericFontName_MonospaceSerif = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupGenericFontName_ProportionalSansSerif"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupGenericFontName_ProportionalSansSerif = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupGenericFontName_ProportionalSerif"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupGenericFontName_ProportionalSerif = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupGenericFontName_SansSerif"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupGenericFontName_SansSerif = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupGenericFontName_Serif"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupGenericFontName_Serif = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextMarkupGenericFontName_SmallCapital"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextMarkupGenericFontName_SmallCapital = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextVerticalLayout_LeftToRight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextVerticalLayout_LeftToRight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTextVerticalLayout_RightToLeft"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTextVerticalLayout_RightToLeft = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimeCodeFormatDescriptionExtension_SourceReferenceName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTimeCodeFormatDescriptionExtension_SourceReferenceName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimeCodeFormatDescriptionKey_LangCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTimeCodeFormatDescriptionKey_LangCode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimeCodeFormatDescriptionKey_Value"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTimeCodeFormatDescriptionKey_Value = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimeEpochKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTimeEpochKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimeFlagsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTimeFlagsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimeIndefinite"); err == nil && ptr != 0 {
		KCMTimeIndefinite = *(*CMTime)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimeInvalid"); err == nil && ptr != 0 {
		KCMTimeInvalid = *(*CMTime)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimeMappingInvalid"); err == nil && ptr != 0 {
		KCMTimeMappingInvalid = *(*CMTimeMapping)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimeMappingSourceKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTimeMappingSourceKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimeMappingTargetKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTimeMappingTargetKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimeNegativeInfinity"); err == nil && ptr != 0 {
		KCMTimeNegativeInfinity = *(*CMTime)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimePositiveInfinity"); err == nil && ptr != 0 {
		KCMTimePositiveInfinity = *(*CMTime)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimeRangeDurationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTimeRangeDurationKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimeRangeInvalid"); err == nil && ptr != 0 {
		KCMTimeRangeInvalid = *(*CMTimeRange)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimeRangeStartKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTimeRangeStartKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimeRangeZero"); err == nil && ptr != 0 {
		KCMTimeRangeZero = *(*CMTimeRange)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimeScaleKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTimeScaleKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimeValueKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTimeValueKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimeZero"); err == nil && ptr != 0 {
		KCMTimeZero = *(*CMTime)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimebaseNotificationKey_EventTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTimebaseNotificationKey_EventTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimebaseNotification_EffectiveRateChanged"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTimebaseNotification_EffectiveRateChanged = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimebaseNotification_TimeJumped"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMTimebaseNotification_TimeJumped = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMTimingInfoInvalid"); err == nil && ptr != 0 {
		KCMTimingInfoInvalid = *(*CMSampleTimingInfo)(unsafe.Pointer(ptr))
	}

}
