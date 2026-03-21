// Code generated from Apple documentation. DO NOT EDIT.

package corevideo

import (
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// KCVBufferMovieTimeKey is the movie time associated with the buffer. Generally only available for frames emitted by QuickTime (type [CFDictionary] containing the [kCVBufferTimeValueKey] and [kCVBufferTimeScaleKey] keys).
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVBufferMovieTimeKey
	KCVBufferMovieTimeKey string
	// KCVBufferNonPropagatedAttachmentsKey is attachments that should not be copied when using the [CVBufferPropagateAttachments(_:_:)] function (type [CFDictionary], containing a list of attachments as key-value pairs).
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVBufferNonPropagatedAttachmentsKey
	KCVBufferNonPropagatedAttachmentsKey string
	// KCVBufferPropagatedAttachmentsKey is attachments that should be copied when using the [CVBufferPropagateAttachments(_:_:)] function (type [CFDictionary], containing a list of attachments as key-value pairs).
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVBufferPropagatedAttachmentsKey
	KCVBufferPropagatedAttachmentsKey string
	// KCVBufferTimeScaleKey is the time scale associated with the movie.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVBufferTimeScaleKey
	KCVBufferTimeScaleKey string
	// KCVBufferTimeValueKey is the time value associated with the movie.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVBufferTimeValueKey
	KCVBufferTimeValueKey string
	// KCVImageBufferAlphaChannelIsOpaque is a key to indicate whether the alpha channel is fully opaque.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferAlphaChannelIsOpaque
	KCVImageBufferAlphaChannelIsOpaque string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferAlphaChannelModeKey
	KCVImageBufferAlphaChannelModeKey string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferAlphaChannelMode_PremultipliedAlpha
	KCVImageBufferAlphaChannelMode_PremultipliedAlpha string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferAlphaChannelMode_StraightAlpha
	KCVImageBufferAlphaChannelMode_StraightAlpha string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferAmbientViewingEnvironmentKey
	KCVImageBufferAmbientViewingEnvironmentKey string
	// KCVImageBufferCGColorSpaceKey is a key to the color space of the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferCGColorSpaceKey
	KCVImageBufferCGColorSpaceKey string
	// KCVImageBufferChromaLocationBottomFieldKey is a key to the location of chroma bottom field information in the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferChromaLocationBottomFieldKey
	KCVImageBufferChromaLocationBottomFieldKey string
	// KCVImageBufferChromaLocationTopFieldKey is a key to the location of chroma top field information in the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferChromaLocationTopFieldKey
	KCVImageBufferChromaLocationTopFieldKey string
	// KCVImageBufferChromaLocation_Bottom is a key that indicates the chroma sample is horizontally centered, but is co-sited with the bottom row of luma samples.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferChromaLocation_Bottom
	KCVImageBufferChromaLocation_Bottom string
	// KCVImageBufferChromaLocation_BottomLeft is a key that indicates the chroma sample is co-sited with the bottom-left luma sample.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferChromaLocation_BottomLeft
	KCVImageBufferChromaLocation_BottomLeft string
	// KCVImageBufferChromaLocation_Center is a key that indicates the chroma sample is fully centered.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferChromaLocation_Center
	KCVImageBufferChromaLocation_Center string
	// KCVImageBufferChromaLocation_DV420 is a key that indicates the Cr and Cb samples are alternatingly co-sited with the left luma samples of the same field.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferChromaLocation_DV420
	KCVImageBufferChromaLocation_DV420 string
	// KCVImageBufferChromaLocation_Left is a key that indicates the chroma sample is horizontally co-sited with the left column of luma samples, but centered vertically.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferChromaLocation_Left
	KCVImageBufferChromaLocation_Left string
	// KCVImageBufferChromaLocation_Top is a key that indicates the chroma sample is horizontally centered, but is co-sited with the top row of luma samples.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferChromaLocation_Top
	KCVImageBufferChromaLocation_Top string
	// KCVImageBufferChromaLocation_TopLeft is a key that indicates the chroma sample is co-sited with the top-left luma sample.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferChromaLocation_TopLeft
	KCVImageBufferChromaLocation_TopLeft string
	// KCVImageBufferChromaSubsamplingKey is a key to the original format of subsampled data in the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferChromaSubsamplingKey
	KCVImageBufferChromaSubsamplingKey string
	// KCVImageBufferChromaSubsampling_411 is a key that indicates the original chroma-subsampled data used 4:1:1 formatting.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferChromaSubsampling_411
	KCVImageBufferChromaSubsampling_411 string
	// KCVImageBufferChromaSubsampling_420 is a key that indicates the original chroma-subsampled data used 4:2:0 formatting.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferChromaSubsampling_420
	KCVImageBufferChromaSubsampling_420 string
	// KCVImageBufferChromaSubsampling_422 is a key that indicates the original chroma-subsampled data used 4:2:2 formatting.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferChromaSubsampling_422
	KCVImageBufferChromaSubsampling_422 string
	// KCVImageBufferCleanApertureHeightKey is a key to the clean aperture height of the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferCleanApertureHeightKey
	KCVImageBufferCleanApertureHeightKey string
	// KCVImageBufferCleanApertureHorizontalOffsetKey is a key to the clean aperture horizontal offset value from the center of the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferCleanApertureHorizontalOffsetKey
	KCVImageBufferCleanApertureHorizontalOffsetKey string
	// KCVImageBufferCleanApertureKey is a key to the dictionary describing the clean aperture for the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferCleanApertureKey
	KCVImageBufferCleanApertureKey string
	// KCVImageBufferCleanApertureVerticalOffsetKey is a key to the clean aperture vertical offset value from the center of the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferCleanApertureVerticalOffsetKey
	KCVImageBufferCleanApertureVerticalOffsetKey string
	// KCVImageBufferCleanApertureWidthKey is a key to the clean aperture width of the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferCleanApertureWidthKey
	KCVImageBufferCleanApertureWidthKey string
	// KCVImageBufferColorPrimariesKey is a key to the color primaries gamut for the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferColorPrimariesKey
	KCVImageBufferColorPrimariesKey string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferColorPrimaries_DCI_P3
	KCVImageBufferColorPrimaries_DCI_P3 string
	// KCVImageBufferColorPrimaries_EBU_3213 is a key to the color primaries gamut for PAL video.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferColorPrimaries_EBU_3213
	KCVImageBufferColorPrimaries_EBU_3213 string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferColorPrimaries_ITU_R_2020
	KCVImageBufferColorPrimaries_ITU_R_2020 string
	// KCVImageBufferColorPrimaries_ITU_R_709_2 is a key to the color primaries gamut for HD video.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferColorPrimaries_ITU_R_709_2
	KCVImageBufferColorPrimaries_ITU_R_709_2 string
	// KCVImageBufferColorPrimaries_P22 is a key to the color primaries gamut for sRGB video.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferColorPrimaries_P22
	KCVImageBufferColorPrimaries_P22 string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferColorPrimaries_P3_D65
	KCVImageBufferColorPrimaries_P3_D65 string
	// KCVImageBufferColorPrimaries_SMPTE_C is a key to the color primaries gamut for standard-definition video.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferColorPrimaries_SMPTE_C
	KCVImageBufferColorPrimaries_SMPTE_C string
	// KCVImageBufferContentLightLevelInfoKey is a key to the content light level information.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferContentLightLevelInfoKey
	KCVImageBufferContentLightLevelInfoKey string
	// KCVImageBufferDisplayDimensionsKey is a key to the dictionary describing the display dimensions for the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferDisplayDimensionsKey
	KCVImageBufferDisplayDimensionsKey string
	// KCVImageBufferDisplayHeightKey is a key to the display height of the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferDisplayHeightKey
	KCVImageBufferDisplayHeightKey string
	// KCVImageBufferDisplayMaskRectangleKey is specifies the rectangular display area within the image.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferDisplayMaskRectangleKey
	KCVImageBufferDisplayMaskRectangleKey string
	// KCVImageBufferDisplayMaskRectangleStereoLeftKey is specifies the rectangular display area within the left-eye view of stereo images, using the same keys as `kCVImageBufferDisplayMaskRectangleKey`.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferDisplayMaskRectangleStereoLeftKey
	KCVImageBufferDisplayMaskRectangleStereoLeftKey string
	// KCVImageBufferDisplayMaskRectangleStereoRightKey is specifies the rectangular display area within the right-eye view of stereo images, using the same keys as `kCVImageBufferDisplayMaskRectangleKey`.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferDisplayMaskRectangleStereoRightKey
	KCVImageBufferDisplayMaskRectangleStereoRightKey string
	// KCVImageBufferDisplayMaskRectangle_LeftEdgePointsKey is specifies inset points on the left vertical edge of the rectangle.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferDisplayMaskRectangle_LeftEdgePointsKey
	KCVImageBufferDisplayMaskRectangle_LeftEdgePointsKey string
	// KCVImageBufferDisplayMaskRectangle_RectangleHeightKey is specifies the height of the rectangle starting at the rectangle’s top offset toward the rectangle’s bottom edge.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferDisplayMaskRectangle_RectangleHeightKey
	KCVImageBufferDisplayMaskRectangle_RectangleHeightKey string
	// KCVImageBufferDisplayMaskRectangle_RectangleLeftKey is specifies the horizontal pixel offset of the rectangle from the left of the bounding raster.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferDisplayMaskRectangle_RectangleLeftKey
	KCVImageBufferDisplayMaskRectangle_RectangleLeftKey string
	// KCVImageBufferDisplayMaskRectangle_RectangleTopKey is specifies the vertical pixel offset of the rectangle from the top of the bounding raster.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferDisplayMaskRectangle_RectangleTopKey
	KCVImageBufferDisplayMaskRectangle_RectangleTopKey string
	// KCVImageBufferDisplayMaskRectangle_RectangleWidthKey is specifies the width of the rectangle starting at the rectangle’s left offset toward the rectangle’s right edge.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferDisplayMaskRectangle_RectangleWidthKey
	KCVImageBufferDisplayMaskRectangle_RectangleWidthKey string
	// KCVImageBufferDisplayMaskRectangle_ReferenceRasterHeightKey is specifies the height in pixels of the 2D coordinate system to define the rectangle.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferDisplayMaskRectangle_ReferenceRasterHeightKey
	KCVImageBufferDisplayMaskRectangle_ReferenceRasterHeightKey string
	// KCVImageBufferDisplayMaskRectangle_ReferenceRasterWidthKey is specifies the width in pixels of the 2D coordinate system to define the rectangle.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferDisplayMaskRectangle_ReferenceRasterWidthKey
	KCVImageBufferDisplayMaskRectangle_ReferenceRasterWidthKey string
	// KCVImageBufferDisplayMaskRectangle_RightEdgePointsKey is specifies inset points on the right vertical edge of the rectangle.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferDisplayMaskRectangle_RightEdgePointsKey
	KCVImageBufferDisplayMaskRectangle_RightEdgePointsKey string
	// KCVImageBufferDisplayWidthKey is a key to the display width of the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferDisplayWidthKey
	KCVImageBufferDisplayWidthKey string
	// KCVImageBufferFieldCountKey is a key to the field count for the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferFieldCountKey
	KCVImageBufferFieldCountKey string
	// KCVImageBufferFieldDetailKey is a key to the field detail for an image buffer that indicates the order of interlaced video data in the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferFieldDetailKey
	KCVImageBufferFieldDetailKey string
	// KCVImageBufferFieldDetailSpatialFirstLineEarly is a key to the spatial first line early detail field of the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferFieldDetailSpatialFirstLineEarly
	KCVImageBufferFieldDetailSpatialFirstLineEarly string
	// KCVImageBufferFieldDetailSpatialFirstLineLate is a key to the spatial first line late detail field of the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferFieldDetailSpatialFirstLineLate
	KCVImageBufferFieldDetailSpatialFirstLineLate string
	// KCVImageBufferFieldDetailTemporalBottomFirst is a key to the temporal bottom first detail field of the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferFieldDetailTemporalBottomFirst
	KCVImageBufferFieldDetailTemporalBottomFirst string
	// KCVImageBufferFieldDetailTemporalTopFirst is a key to the temporal top first detail field of the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferFieldDetailTemporalTopFirst
	KCVImageBufferFieldDetailTemporalTopFirst string
	// KCVImageBufferGammaLevelKey is a key to the gamma level for the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferGammaLevelKey
	KCVImageBufferGammaLevelKey string
	// KCVImageBufferICCProfileKey is a key to the ICC color profile for the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferICCProfileKey
	KCVImageBufferICCProfileKey string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferLogTransferFunctionKey
	KCVImageBufferLogTransferFunctionKey string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferLogTransferFunction_AppleLog
	KCVImageBufferLogTransferFunction_AppleLog string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferLogTransferFunction_AppleLog2
	KCVImageBufferLogTransferFunction_AppleLog2 string
	// KCVImageBufferMasteringDisplayColorVolumeKey is a key to the mastering display color volume.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferMasteringDisplayColorVolumeKey
	KCVImageBufferMasteringDisplayColorVolumeKey string
	// KCVImageBufferPixelAspectRatioHorizontalSpacingKey is a key to the horizontal component of the image buffer aspect ratio.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferPixelAspectRatioHorizontalSpacingKey
	KCVImageBufferPixelAspectRatioHorizontalSpacingKey string
	// KCVImageBufferPixelAspectRatioKey is a key to the dictionary describing the pixel aspect ratio for the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferPixelAspectRatioKey
	KCVImageBufferPixelAspectRatioKey string
	// KCVImageBufferPixelAspectRatioVerticalSpacingKey is a key to the vertical component of the image buffer aspect ratio.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferPixelAspectRatioVerticalSpacingKey
	KCVImageBufferPixelAspectRatioVerticalSpacingKey string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferPostDecodeProcessingFrameMetadataKey
	KCVImageBufferPostDecodeProcessingFrameMetadataKey string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferPostDecodeProcessingSequenceMetadataKey
	KCVImageBufferPostDecodeProcessingSequenceMetadataKey string
	// KCVImageBufferPreferredCleanApertureKey is a key to the dictionary describing the preferred clean aperture for the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferPreferredCleanApertureKey
	KCVImageBufferPreferredCleanApertureKey string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferRegionOfInterestKey
	KCVImageBufferRegionOfInterestKey string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferSceneIlluminationKey
	KCVImageBufferSceneIlluminationKey string
	// KCVImageBufferTransferFunctionKey is a key to the transfer function for the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferTransferFunctionKey
	KCVImageBufferTransferFunctionKey string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferTransferFunction_ITU_R_2020
	KCVImageBufferTransferFunction_ITU_R_2020 string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferTransferFunction_ITU_R_2100_HLG
	KCVImageBufferTransferFunction_ITU_R_2100_HLG string
	// KCVImageBufferTransferFunction_ITU_R_709_2 is a key to the transfer function for high-definition and standard-definition video.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferTransferFunction_ITU_R_709_2
	KCVImageBufferTransferFunction_ITU_R_709_2 string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferTransferFunction_Linear
	KCVImageBufferTransferFunction_Linear string
	// KCVImageBufferTransferFunction_SMPTE_240M_1995 is a key to the transfer function for HDTV interim video.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferTransferFunction_SMPTE_240M_1995
	KCVImageBufferTransferFunction_SMPTE_240M_1995 string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferTransferFunction_SMPTE_ST_2084_PQ
	KCVImageBufferTransferFunction_SMPTE_ST_2084_PQ string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferTransferFunction_SMPTE_ST_428_1
	KCVImageBufferTransferFunction_SMPTE_ST_428_1 string
	// KCVImageBufferTransferFunction_UseGamma is a key to the transfer function that’s defined by the gamma level value of the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferTransferFunction_UseGamma
	KCVImageBufferTransferFunction_UseGamma string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferTransferFunction_sRGB
	KCVImageBufferTransferFunction_sRGB string
	// KCVImageBufferYCbCrMatrixKey is a key to the YCbCr to RGB color conversion matrix for the image buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferYCbCrMatrixKey
	KCVImageBufferYCbCrMatrixKey string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferYCbCrMatrix_ITU_R_2020
	KCVImageBufferYCbCrMatrix_ITU_R_2020 string
	// KCVImageBufferYCbCrMatrix_ITU_R_601_4 is a key to the conversion matrix for standard digital television images, that follows the ITU R 601 standard.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferYCbCrMatrix_ITU_R_601_4
	KCVImageBufferYCbCrMatrix_ITU_R_601_4 string
	// KCVImageBufferYCbCrMatrix_ITU_R_709_2 is a key to the conversion matrix for HDTV digital television images, that follows the ITU R 709 standard.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferYCbCrMatrix_ITU_R_709_2
	KCVImageBufferYCbCrMatrix_ITU_R_709_2 string
	// KCVImageBufferYCbCrMatrix_SMPTE_240M_1995 is a key to the conversion matrix for 1920 x 1135 HDTV images, that follows the SMPTE 240M 1995 standard.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferYCbCrMatrix_SMPTE_240M_1995
	KCVImageBufferYCbCrMatrix_SMPTE_240M_1995 string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVMetalBufferCacheMaximumBufferAgeKey
	KCVMetalBufferCacheMaximumBufferAgeKey string
	// KCVMetalTextureCacheMaximumTextureAgeKey is the length of time, in seconds, before the cache is automatically evicted.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVMetalTextureCacheMaximumTextureAgeKey
	KCVMetalTextureCacheMaximumTextureAgeKey string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVMetalTextureStorageMode
	KCVMetalTextureStorageMode string
	// KCVMetalTextureUsage is the set of options that define how you can use a texture on the GPU.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVMetalTextureUsage
	KCVMetalTextureUsage string
	// KCVPixelBufferBytesPerRowAlignmentKey is a key to a number that specifies the alignment of number of bytes per row in the pixel buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferBytesPerRowAlignmentKey
	KCVPixelBufferBytesPerRowAlignmentKey string
	// KCVPixelBufferCGBitmapContextCompatibilityKey is a key to a Boolean value that indicates whether the pixel buffer is compatible with Core Graphics bitmap contexts.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferCGBitmapContextCompatibilityKey
	KCVPixelBufferCGBitmapContextCompatibilityKey string
	// KCVPixelBufferCGImageCompatibilityKey is a key to a Boolean value that indicates whether the pixel buffer is compatible with Core Graphics bitmap image types.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferCGImageCompatibilityKey
	KCVPixelBufferCGImageCompatibilityKey string
	// KCVPixelBufferExtendedPixelsBottomKey is a key to the number of pixels padding the bottom of the image.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferExtendedPixelsBottomKey
	KCVPixelBufferExtendedPixelsBottomKey string
	// KCVPixelBufferExtendedPixelsLeftKey is a key to the number of pixels padding the left of the image.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferExtendedPixelsLeftKey
	KCVPixelBufferExtendedPixelsLeftKey string
	// KCVPixelBufferExtendedPixelsRightKey is a key to the number of pixels padding the right of the image.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferExtendedPixelsRightKey
	KCVPixelBufferExtendedPixelsRightKey string
	// KCVPixelBufferExtendedPixelsTopKey is a key to the number of pixels padding the top of the image.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferExtendedPixelsTopKey
	KCVPixelBufferExtendedPixelsTopKey string
	// KCVPixelBufferHeightKey is a key to the height of the pixel buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferHeightKey
	KCVPixelBufferHeightKey string
	// KCVPixelBufferIOSurfaceCoreAnimationCompatibilityKey is a key to a Boolean value that indicates whether Core Animation can display the pixel buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferIOSurfaceCoreAnimationCompatibilityKey
	KCVPixelBufferIOSurfaceCoreAnimationCompatibilityKey string
	// KCVPixelBufferIOSurfaceOpenGLFBOCompatibilityKey is a key to a Boolean value that indicates whether OpenGL can create a valid texture for use as a color buffer attachment.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferIOSurfaceOpenGLFBOCompatibilityKey
	KCVPixelBufferIOSurfaceOpenGLFBOCompatibilityKey string
	// KCVPixelBufferIOSurfaceOpenGLTextureCompatibilityKey is a key to a Boolean value that indicates whether OpenGL can create a valid texture object from the IOSurface-backed pixel buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferIOSurfaceOpenGLTextureCompatibilityKey
	KCVPixelBufferIOSurfaceOpenGLTextureCompatibilityKey string
	// KCVPixelBufferIOSurfacePropertiesKey is a key to the dictionary containing optional properties for the IOSurface framework.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferIOSurfacePropertiesKey
	KCVPixelBufferIOSurfacePropertiesKey string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferIOSurfacePurgeableKey
	KCVPixelBufferIOSurfacePurgeableKey string
	// KCVPixelBufferMemoryAllocatorKey is a key to the allocator that the system uses to create the pixel buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferMemoryAllocatorKey
	KCVPixelBufferMemoryAllocatorKey string
	// KCVPixelBufferMetalCompatibilityKey is a key to a Boolean value that indicates whether the pixel buffer is compatible with the Metal framework.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferMetalCompatibilityKey
	KCVPixelBufferMetalCompatibilityKey string
	// KCVPixelBufferOpenGLCompatibilityKey is a key to a Boolean value that indicates whether the pixel buffer is compatible with OpenGL contexts.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferOpenGLCompatibilityKey
	KCVPixelBufferOpenGLCompatibilityKey string
	// KCVPixelBufferOpenGLTextureCacheCompatibilityKey is a key to a Boolean value that indicates whether OpenGL performs format conversions of the texture-cache data in a shader.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferOpenGLTextureCacheCompatibilityKey
	KCVPixelBufferOpenGLTextureCacheCompatibilityKey string
	// KCVPixelBufferPixelFormatTypeKey is a key to one or more pixel buffer format types.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferPixelFormatTypeKey
	KCVPixelBufferPixelFormatTypeKey string
	// KCVPixelBufferPlaneAlignmentKey is a key to a number that specifies the alignment of the planes in the pixel buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferPlaneAlignmentKey
	KCVPixelBufferPlaneAlignmentKey string
	// KCVPixelBufferPoolAllocationThresholdKey is the key you use to set the auxiliary attributes dictionary.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferPoolAllocationThresholdKey
	KCVPixelBufferPoolAllocationThresholdKey string
	// KCVPixelBufferPoolFreeBufferNotification is a notification that the system posts if a buffer becomes available after it fails to create a pixel buffer with auxiliary attributes because it exceeded the threshold you specified.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferPoolFreeBufferNotification
	KCVPixelBufferPoolFreeBufferNotification string
	// KCVPixelBufferPoolMaximumBufferAgeKey is the key you use to set the maximum allowable age for a buffer in the pixel buffer pool.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferPoolMaximumBufferAgeKey
	KCVPixelBufferPoolMaximumBufferAgeKey string
	// KCVPixelBufferPoolMinimumBufferCountKey is the minimum number of buffers allowed in the pixel buffer pool.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferPoolMinimumBufferCountKey
	KCVPixelBufferPoolMinimumBufferCountKey string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferProResRAWKey_BlackLevel
	KCVPixelBufferProResRAWKey_BlackLevel string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferProResRAWKey_ColorMatrix
	KCVPixelBufferProResRAWKey_ColorMatrix string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferProResRAWKey_GainFactor
	KCVPixelBufferProResRAWKey_GainFactor string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferProResRAWKey_MetadataExtension
	KCVPixelBufferProResRAWKey_MetadataExtension string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferProResRAWKey_RecommendedCrop
	KCVPixelBufferProResRAWKey_RecommendedCrop string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferProResRAWKey_SenselSitingOffsets
	KCVPixelBufferProResRAWKey_SenselSitingOffsets string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferProResRAWKey_WhiteBalanceBlueFactor
	KCVPixelBufferProResRAWKey_WhiteBalanceBlueFactor string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferProResRAWKey_WhiteBalanceCCT
	KCVPixelBufferProResRAWKey_WhiteBalanceCCT string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferProResRAWKey_WhiteBalanceRedFactor
	KCVPixelBufferProResRAWKey_WhiteBalanceRedFactor string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferProResRAWKey_WhiteLevel
	KCVPixelBufferProResRAWKey_WhiteLevel string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferVersatileBayerKey_BayerPattern
	KCVPixelBufferVersatileBayerKey_BayerPattern string
	// KCVPixelBufferWidthKey is a key to the width of the pixel buffer.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferWidthKey
	KCVPixelBufferWidthKey string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatBitsPerBlock
	KCVPixelFormatBitsPerBlock string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatBitsPerComponent
	KCVPixelFormatBitsPerComponent string
	// KCVPixelFormatBlackBlock is the bit pattern for a block of black pixels (type [CFData]. If this key is absent, black is assumed to be all zeros. If present, this should be `bitsPerPixel` bits long; if `bitsPerPixel` is less than a byte, repeat the bit pattern for the full byte.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatBlackBlock
	KCVPixelFormatBlackBlock string
	// KCVPixelFormatBlockHeight is the height, in pixels, of the smallest byte-addressable group of pixels (type [CFNumber]). Assumed to be 1 if this key is not present.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatBlockHeight
	KCVPixelFormatBlockHeight string
	// KCVPixelFormatBlockHorizontalAlignment is the horizontal alignment requirements of this format (type [CFNumber]). For example,the alignment for v210 would be 8 here for the horizontal case to match the standard v210 row alignment value of 48. Assumed to be 1 if this key is not present.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatBlockHorizontalAlignment
	KCVPixelFormatBlockHorizontalAlignment string
	// KCVPixelFormatBlockVerticalAlignment is the vertical alignment requirements of this format (type [CFNumber]). Assumed to be 1 if this key is not present.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatBlockVerticalAlignment
	KCVPixelFormatBlockVerticalAlignment string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatBlockWidth
	KCVPixelFormatBlockWidth string
	// KCVPixelFormatCGBitmapContextCompatibility is if true, this format is compatible with Core Graphics bitmap contexts(type [CFBoolean]).
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatCGBitmapContextCompatibility
	KCVPixelFormatCGBitmapContextCompatibility string
	// KCVPixelFormatCGBitmapInfo is the Core Graphics bitmap information for this pixel format (if applicable).
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatCGBitmapInfo
	KCVPixelFormatCGBitmapInfo string
	// KCVPixelFormatCGImageCompatibility is if true, this format is compatible with the [CGImage] type (type [CFBoolean]).
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatCGImageCompatibility
	KCVPixelFormatCGImageCompatibility string
	// KCVPixelFormatCodecType is the codec type (type [CFString]). For example, `'2vuy'` or `k422YpCbCr8CodecType`.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatCodecType
	KCVPixelFormatCodecType string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatComponentRange
	KCVPixelFormatComponentRange string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatComponentRange_FullRange
	KCVPixelFormatComponentRange_FullRange string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatComponentRange_VideoRange
	KCVPixelFormatComponentRange_VideoRange string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatComponentRange_WideRange
	KCVPixelFormatComponentRange_WideRange string
	// KCVPixelFormatConstant is the pixel format constant for QuickTime.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatConstant
	KCVPixelFormatConstant string
	// KCVPixelFormatContainsAlpha is a Boolean value where [kCFBooleanTrue] indicates that the format contains alpha and some images may be considered transparent; [kCFBooleanFalse] indicates that there is no alpha and images are always opaque.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatContainsAlpha
	KCVPixelFormatContainsAlpha string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatContainsGrayscale
	KCVPixelFormatContainsGrayscale string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatContainsRGB
	KCVPixelFormatContainsRGB string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatContainsSenselArray
	KCVPixelFormatContainsSenselArray string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatContainsYCbCr
	KCVPixelFormatContainsYCbCr string
	// KCVPixelFormatFillExtendedPixelsCallback is a custom extended pixel fill algorithm (type [CFData]). See [CVFillExtendedPixelsCallBack] and [CVFillExtendedPixelsCallBackData] for more information.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatFillExtendedPixelsCallback
	KCVPixelFormatFillExtendedPixelsCallback string
	// KCVPixelFormatFourCC is the Microsoft FourCC equivalent code for this pixel format (type [CFString]).
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatFourCC
	KCVPixelFormatFourCC string
	// KCVPixelFormatHorizontalSubsampling is horizontal subsampling information for this plane (type [CFNumber]). Assumed to be 1 if this key is not present.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatHorizontalSubsampling
	KCVPixelFormatHorizontalSubsampling string
	// KCVPixelFormatName is the name of the pixel format (type [CFString]). This should be the same as the codec name you would use in QuickTime.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatName
	KCVPixelFormatName string
	// KCVPixelFormatOpenGLCompatibility is if true, this format is compatible with OpenGL (type [CFBoolean]).
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatOpenGLCompatibility
	KCVPixelFormatOpenGLCompatibility string
	// KCVPixelFormatOpenGLFormat is the OpenGL format used to describe this image plane (if applicable). See the [OpenGL specification] for possible values.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatOpenGLFormat
	KCVPixelFormatOpenGLFormat string
	// KCVPixelFormatOpenGLInternalFormat is the OpenGL internal format for this pixel format (if applicable). See the [OpenGL specification] for possible values.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatOpenGLInternalFormat
	KCVPixelFormatOpenGLInternalFormat string
	// KCVPixelFormatOpenGLType is the OpenGL type to describe this image plane (if applicable). See the [OpenGL specification] for possible values.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatOpenGLType
	KCVPixelFormatOpenGLType string
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatPlanes
	KCVPixelFormatPlanes string
	// KCVPixelFormatQDCompatibility is if true, this format is compatible with QuickDraw (type [CFBoolean]).
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatQDCompatibility
	KCVPixelFormatQDCompatibility string
	// KCVPixelFormatVerticalSubsampling is vertical subsampling information for this plane (type [CFNumber]). Assumed to be 1 if this key is not present.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatVerticalSubsampling
	KCVPixelFormatVerticalSubsampling string
)

var (
	// KCVIndefiniteTime is an unknown or indefinite time. For example, [CVDisplayLinkGetNominalOutputVideoRefreshPeriod(_:)] returns `kCVIndefiniteTime` if the display link specified is not valid.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVIndefiniteTime
	KCVIndefiniteTime CVTime
	// KCVZeroTime is zero time or duration. For example, [CVDisplayLinkGetOutputVideoLatency(_:)] returns `kCVZeroTime` for zero video latency.
	//
	// See: https://developer.apple.com/documentation/CoreVideo/kCVZeroTime
	KCVZeroTime CVTime
)
func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVBufferMovieTimeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVBufferMovieTimeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVBufferNonPropagatedAttachmentsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVBufferNonPropagatedAttachmentsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVBufferPropagatedAttachmentsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVBufferPropagatedAttachmentsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVBufferTimeScaleKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVBufferTimeScaleKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVBufferTimeValueKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVBufferTimeValueKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferAlphaChannelIsOpaque"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferAlphaChannelIsOpaque = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferAlphaChannelModeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferAlphaChannelModeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferAlphaChannelMode_PremultipliedAlpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferAlphaChannelMode_PremultipliedAlpha = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferAlphaChannelMode_StraightAlpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferAlphaChannelMode_StraightAlpha = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferAmbientViewingEnvironmentKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferAmbientViewingEnvironmentKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferCGColorSpaceKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferCGColorSpaceKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferChromaLocationBottomFieldKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferChromaLocationBottomFieldKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferChromaLocationTopFieldKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferChromaLocationTopFieldKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferChromaLocation_Bottom"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferChromaLocation_Bottom = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferChromaLocation_BottomLeft"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferChromaLocation_BottomLeft = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferChromaLocation_Center"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferChromaLocation_Center = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferChromaLocation_DV420"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferChromaLocation_DV420 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferChromaLocation_Left"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferChromaLocation_Left = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferChromaLocation_Top"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferChromaLocation_Top = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferChromaLocation_TopLeft"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferChromaLocation_TopLeft = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferChromaSubsamplingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferChromaSubsamplingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferChromaSubsampling_411"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferChromaSubsampling_411 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferChromaSubsampling_420"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferChromaSubsampling_420 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferChromaSubsampling_422"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferChromaSubsampling_422 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferCleanApertureHeightKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferCleanApertureHeightKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferCleanApertureHorizontalOffsetKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferCleanApertureHorizontalOffsetKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferCleanApertureKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferCleanApertureKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferCleanApertureVerticalOffsetKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferCleanApertureVerticalOffsetKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferCleanApertureWidthKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferCleanApertureWidthKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferColorPrimariesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferColorPrimariesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferColorPrimaries_DCI_P3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferColorPrimaries_DCI_P3 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferColorPrimaries_EBU_3213"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferColorPrimaries_EBU_3213 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferColorPrimaries_ITU_R_2020"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferColorPrimaries_ITU_R_2020 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferColorPrimaries_ITU_R_709_2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferColorPrimaries_ITU_R_709_2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferColorPrimaries_P22"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferColorPrimaries_P22 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferColorPrimaries_P3_D65"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferColorPrimaries_P3_D65 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferColorPrimaries_SMPTE_C"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferColorPrimaries_SMPTE_C = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferContentLightLevelInfoKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferContentLightLevelInfoKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferDisplayDimensionsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferDisplayDimensionsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferDisplayHeightKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferDisplayHeightKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferDisplayMaskRectangleKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferDisplayMaskRectangleKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferDisplayMaskRectangleStereoLeftKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferDisplayMaskRectangleStereoLeftKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferDisplayMaskRectangleStereoRightKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferDisplayMaskRectangleStereoRightKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferDisplayMaskRectangle_LeftEdgePointsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferDisplayMaskRectangle_LeftEdgePointsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferDisplayMaskRectangle_RectangleHeightKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferDisplayMaskRectangle_RectangleHeightKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferDisplayMaskRectangle_RectangleLeftKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferDisplayMaskRectangle_RectangleLeftKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferDisplayMaskRectangle_RectangleTopKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferDisplayMaskRectangle_RectangleTopKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferDisplayMaskRectangle_RectangleWidthKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferDisplayMaskRectangle_RectangleWidthKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferDisplayMaskRectangle_ReferenceRasterHeightKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferDisplayMaskRectangle_ReferenceRasterHeightKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferDisplayMaskRectangle_ReferenceRasterWidthKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferDisplayMaskRectangle_ReferenceRasterWidthKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferDisplayMaskRectangle_RightEdgePointsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferDisplayMaskRectangle_RightEdgePointsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferDisplayWidthKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferDisplayWidthKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferFieldCountKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferFieldCountKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferFieldDetailKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferFieldDetailKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferFieldDetailSpatialFirstLineEarly"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferFieldDetailSpatialFirstLineEarly = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferFieldDetailSpatialFirstLineLate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferFieldDetailSpatialFirstLineLate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferFieldDetailTemporalBottomFirst"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferFieldDetailTemporalBottomFirst = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferFieldDetailTemporalTopFirst"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferFieldDetailTemporalTopFirst = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferGammaLevelKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferGammaLevelKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferICCProfileKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferICCProfileKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferLogTransferFunctionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferLogTransferFunctionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferLogTransferFunction_AppleLog"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferLogTransferFunction_AppleLog = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferLogTransferFunction_AppleLog2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferLogTransferFunction_AppleLog2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferMasteringDisplayColorVolumeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferMasteringDisplayColorVolumeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferPixelAspectRatioHorizontalSpacingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferPixelAspectRatioHorizontalSpacingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferPixelAspectRatioKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferPixelAspectRatioKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferPixelAspectRatioVerticalSpacingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferPixelAspectRatioVerticalSpacingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferPostDecodeProcessingFrameMetadataKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferPostDecodeProcessingFrameMetadataKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferPostDecodeProcessingSequenceMetadataKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferPostDecodeProcessingSequenceMetadataKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferPreferredCleanApertureKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferPreferredCleanApertureKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferRegionOfInterestKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferRegionOfInterestKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferSceneIlluminationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferSceneIlluminationKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferTransferFunctionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferTransferFunctionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferTransferFunction_ITU_R_2020"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferTransferFunction_ITU_R_2020 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferTransferFunction_ITU_R_2100_HLG"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferTransferFunction_ITU_R_2100_HLG = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferTransferFunction_ITU_R_709_2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferTransferFunction_ITU_R_709_2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferTransferFunction_Linear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferTransferFunction_Linear = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferTransferFunction_SMPTE_240M_1995"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferTransferFunction_SMPTE_240M_1995 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferTransferFunction_SMPTE_ST_2084_PQ"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferTransferFunction_SMPTE_ST_2084_PQ = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferTransferFunction_SMPTE_ST_428_1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferTransferFunction_SMPTE_ST_428_1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferTransferFunction_UseGamma"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferTransferFunction_UseGamma = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferTransferFunction_sRGB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferTransferFunction_sRGB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferYCbCrMatrixKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferYCbCrMatrixKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferYCbCrMatrix_ITU_R_2020"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferYCbCrMatrix_ITU_R_2020 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferYCbCrMatrix_ITU_R_601_4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferYCbCrMatrix_ITU_R_601_4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferYCbCrMatrix_ITU_R_709_2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferYCbCrMatrix_ITU_R_709_2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVImageBufferYCbCrMatrix_SMPTE_240M_1995"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVImageBufferYCbCrMatrix_SMPTE_240M_1995 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVIndefiniteTime"); err == nil && ptr != 0 {
		KCVIndefiniteTime = *(*CVTime)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVMetalBufferCacheMaximumBufferAgeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVMetalBufferCacheMaximumBufferAgeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVMetalTextureCacheMaximumTextureAgeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVMetalTextureCacheMaximumTextureAgeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVMetalTextureStorageMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVMetalTextureStorageMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVMetalTextureUsage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVMetalTextureUsage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferBytesPerRowAlignmentKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferBytesPerRowAlignmentKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferCGBitmapContextCompatibilityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferCGBitmapContextCompatibilityKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferCGImageCompatibilityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferCGImageCompatibilityKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferExtendedPixelsBottomKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferExtendedPixelsBottomKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferExtendedPixelsLeftKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferExtendedPixelsLeftKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferExtendedPixelsRightKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferExtendedPixelsRightKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferExtendedPixelsTopKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferExtendedPixelsTopKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferHeightKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferHeightKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferIOSurfaceCoreAnimationCompatibilityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferIOSurfaceCoreAnimationCompatibilityKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferIOSurfaceOpenGLFBOCompatibilityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferIOSurfaceOpenGLFBOCompatibilityKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferIOSurfaceOpenGLTextureCompatibilityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferIOSurfaceOpenGLTextureCompatibilityKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferIOSurfacePropertiesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferIOSurfacePropertiesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferIOSurfacePurgeableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferIOSurfacePurgeableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferMemoryAllocatorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferMemoryAllocatorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferMetalCompatibilityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferMetalCompatibilityKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferOpenGLCompatibilityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferOpenGLCompatibilityKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferOpenGLTextureCacheCompatibilityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferOpenGLTextureCacheCompatibilityKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferPixelFormatTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferPixelFormatTypeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferPlaneAlignmentKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferPlaneAlignmentKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferPoolAllocationThresholdKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferPoolAllocationThresholdKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferPoolFreeBufferNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferPoolFreeBufferNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferPoolMaximumBufferAgeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferPoolMaximumBufferAgeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferPoolMinimumBufferCountKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferPoolMinimumBufferCountKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferProResRAWKey_BlackLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferProResRAWKey_BlackLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferProResRAWKey_ColorMatrix"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferProResRAWKey_ColorMatrix = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferProResRAWKey_GainFactor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferProResRAWKey_GainFactor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferProResRAWKey_MetadataExtension"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferProResRAWKey_MetadataExtension = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferProResRAWKey_RecommendedCrop"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferProResRAWKey_RecommendedCrop = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferProResRAWKey_SenselSitingOffsets"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferProResRAWKey_SenselSitingOffsets = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferProResRAWKey_WhiteBalanceBlueFactor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferProResRAWKey_WhiteBalanceBlueFactor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferProResRAWKey_WhiteBalanceCCT"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferProResRAWKey_WhiteBalanceCCT = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferProResRAWKey_WhiteBalanceRedFactor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferProResRAWKey_WhiteBalanceRedFactor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferProResRAWKey_WhiteLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferProResRAWKey_WhiteLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferVersatileBayerKey_BayerPattern"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferVersatileBayerKey_BayerPattern = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelBufferWidthKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelBufferWidthKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatBitsPerBlock"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatBitsPerBlock = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatBitsPerComponent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatBitsPerComponent = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatBlackBlock"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatBlackBlock = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatBlockHeight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatBlockHeight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatBlockHorizontalAlignment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatBlockHorizontalAlignment = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatBlockVerticalAlignment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatBlockVerticalAlignment = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatBlockWidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatBlockWidth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatCGBitmapContextCompatibility"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatCGBitmapContextCompatibility = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatCGBitmapInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatCGBitmapInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatCGImageCompatibility"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatCGImageCompatibility = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatCodecType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatCodecType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatComponentRange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatComponentRange = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatComponentRange_FullRange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatComponentRange_FullRange = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatComponentRange_VideoRange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatComponentRange_VideoRange = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatComponentRange_WideRange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatComponentRange_WideRange = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatConstant"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatConstant = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatContainsAlpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatContainsAlpha = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatContainsGrayscale"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatContainsGrayscale = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatContainsRGB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatContainsRGB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatContainsSenselArray"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatContainsSenselArray = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatContainsYCbCr"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatContainsYCbCr = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatFillExtendedPixelsCallback"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatFillExtendedPixelsCallback = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatFourCC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatFourCC = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatHorizontalSubsampling"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatHorizontalSubsampling = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatOpenGLCompatibility"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatOpenGLCompatibility = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatOpenGLFormat"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatOpenGLFormat = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatOpenGLInternalFormat"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatOpenGLInternalFormat = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatOpenGLType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatOpenGLType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatPlanes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatPlanes = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatQDCompatibility"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatQDCompatibility = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVPixelFormatVerticalSubsampling"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCVPixelFormatVerticalSubsampling = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCVZeroTime"); err == nil && ptr != 0 {
		KCVZeroTime = *(*CVTime)(unsafe.Pointer(ptr))
	}

}

