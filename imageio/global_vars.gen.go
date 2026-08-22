// Code generated from Apple documentation. DO NOT EDIT.

package imageio

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// KCFErrorDomainCGImageMetadata is the domain for metadata-related errors that originate in the Image I/O framework.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCFErrorDomainCGImageMetadata
	KCFErrorDomainCGImageMetadata string
	// See: https://developer.apple.com/documentation/ImageIO/kCGComputeHDRStats
	KCGComputeHDRStats string
	// KCGImageAnimationDelayTime is the number of seconds to wait before displaying the next image in an animated sequence.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageAnimationDelayTime
	KCGImageAnimationDelayTime string
	// KCGImageAnimationLoopCount is the number of times to repeat the animated sequence.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageAnimationLoopCount
	KCGImageAnimationLoopCount string
	// KCGImageAnimationStartIndex is a property that specifies the index of the first frame of an animation.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageAnimationStartIndex
	KCGImageAnimationStartIndex string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageAuxiliaryDataInfoColorSpace
	KCGImageAuxiliaryDataInfoColorSpace string
	// KCGImageAuxiliaryDataInfoData is the auxiliary data for the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageAuxiliaryDataInfoData
	KCGImageAuxiliaryDataInfoData string
	// KCGImageAuxiliaryDataInfoDataDescription is a dictionary of keys that describe the auxiliary data.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageAuxiliaryDataInfoDataDescription
	KCGImageAuxiliaryDataInfoDataDescription string
	// KCGImageAuxiliaryDataInfoMetadata is the metadata for any auxiliary data.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageAuxiliaryDataInfoMetadata
	KCGImageAuxiliaryDataInfoMetadata string
	// KCGImageAuxiliaryDataTypeDepth is the type for depth map information.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageAuxiliaryDataTypeDepth
	KCGImageAuxiliaryDataTypeDepth string
	// KCGImageAuxiliaryDataTypeDisparity is the type for image disparity information.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageAuxiliaryDataTypeDisparity
	KCGImageAuxiliaryDataTypeDisparity string
	// KCGImageAuxiliaryDataTypeHDRGainMap is the type for High Dynamic Range (HDR) gain map information.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageAuxiliaryDataTypeHDRGainMap
	KCGImageAuxiliaryDataTypeHDRGainMap string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageAuxiliaryDataTypeISOGainMap
	KCGImageAuxiliaryDataTypeISOGainMap string
	// KCGImageAuxiliaryDataTypePortraitEffectsMatte is the type for portrait effects matte information.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageAuxiliaryDataTypePortraitEffectsMatte
	KCGImageAuxiliaryDataTypePortraitEffectsMatte string
	// KCGImageAuxiliaryDataTypeSemanticSegmentationGlassesMatte is the type for glasses matte informaton.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageAuxiliaryDataTypeSemanticSegmentationGlassesMatte
	KCGImageAuxiliaryDataTypeSemanticSegmentationGlassesMatte string
	// KCGImageAuxiliaryDataTypeSemanticSegmentationHairMatte is the type for hair matte information.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageAuxiliaryDataTypeSemanticSegmentationHairMatte
	KCGImageAuxiliaryDataTypeSemanticSegmentationHairMatte string
	// KCGImageAuxiliaryDataTypeSemanticSegmentationSkinMatte is the type for skin matte informaton.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageAuxiliaryDataTypeSemanticSegmentationSkinMatte
	KCGImageAuxiliaryDataTypeSemanticSegmentationSkinMatte string
	// KCGImageAuxiliaryDataTypeSemanticSegmentationSkyMatte is the type for sky matte information.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageAuxiliaryDataTypeSemanticSegmentationSkyMatte
	KCGImageAuxiliaryDataTypeSemanticSegmentationSkyMatte string
	// KCGImageAuxiliaryDataTypeSemanticSegmentationTeethMatte is the type for teeth matte information.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageAuxiliaryDataTypeSemanticSegmentationTeethMatte
	KCGImageAuxiliaryDataTypeSemanticSegmentationTeethMatte string
	// KCGImageDestinationBackgroundColor is the background color to use when the image has an alpha component, but the destination format doesn’t support alpha.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationBackgroundColor
	KCGImageDestinationBackgroundColor string
	// KCGImageDestinationDateTime is the date and time information to associate with the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationDateTime
	KCGImageDestinationDateTime string
	// KCGImageDestinationEmbedThumbnail is a Boolean value that indicates whether to embed a thumbnail for JPEG and HEIF images.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationEmbedThumbnail
	KCGImageDestinationEmbedThumbnail string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationEncodeAlternateColorSpace
	KCGImageDestinationEncodeAlternateColorSpace string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationEncodeBaseColorSpace
	KCGImageDestinationEncodeBaseColorSpace string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationEncodeBaseIsSDR
	KCGImageDestinationEncodeBaseIsSDR string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationEncodeBasePixelFormatRequest
	KCGImageDestinationEncodeBasePixelFormatRequest string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationEncodeGainMapPixelFormatRequest
	KCGImageDestinationEncodeGainMapPixelFormatRequest string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationEncodeGainMapSubsampleFactor
	KCGImageDestinationEncodeGainMapSubsampleFactor string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationEncodeGenerateGainMapWithBaseImage
	KCGImageDestinationEncodeGenerateGainMapWithBaseImage string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationEncodeIsBaseImage
	KCGImageDestinationEncodeIsBaseImage string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationEncodeRequest
	KCGImageDestinationEncodeRequest string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationEncodeRequestOptions
	KCGImageDestinationEncodeRequestOptions string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationEncodeToISOGainmap
	KCGImageDestinationEncodeToISOGainmap string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationEncodeToISOHDR
	KCGImageDestinationEncodeToISOHDR string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationEncodeToSDR
	KCGImageDestinationEncodeToSDR string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationEncodeTonemapMode
	KCGImageDestinationEncodeTonemapMode string
	// KCGImageDestinationImageMaxPixelSize is the maximum width and height of the image, in pixels.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationImageMaxPixelSize
	KCGImageDestinationImageMaxPixelSize string
	// KCGImageDestinationLossyCompressionQuality is the desired compression quality to use when writing the image data.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationLossyCompressionQuality
	KCGImageDestinationLossyCompressionQuality string
	// KCGImageDestinationMergeMetadata is a Boolean value that indicates whether to merge new metadata with the image’s existing metadata.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationMergeMetadata
	KCGImageDestinationMergeMetadata string
	// KCGImageDestinationMetadata is the metadata tags to include with the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationMetadata
	KCGImageDestinationMetadata string
	// KCGImageDestinationOptimizeColorForSharing is a Boolean value that indicates whether to create the image using a colorspace.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationOptimizeColorForSharing
	KCGImageDestinationOptimizeColorForSharing string
	// KCGImageDestinationOrientation is the orientation of the image, specified as an EXIF value in the range 1 to 8.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationOrientation
	KCGImageDestinationOrientation string
	// KCGImageDestinationPreserveGainMap is a Boolean value that indicates whether to include a HEIF-embedded gain map in the image data.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationPreserveGainMap
	KCGImageDestinationPreserveGainMap string
	// KCGImageMetadataEnumerateRecursively is an option to enumerate recursively through a set of metadata tags.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataEnumerateRecursively
	KCGImageMetadataEnumerateRecursively string
	// KCGImageMetadataNamespaceDublinCore is the namespace for the Dublin Core Metadata Element Set.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataNamespaceDublinCore
	KCGImageMetadataNamespaceDublinCore string
	// KCGImageMetadataNamespaceExif is the namespace for the Exchangeable Image File (EXIF) format.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataNamespaceExif
	KCGImageMetadataNamespaceExif string
	// KCGImageMetadataNamespaceExifAux is the namespace for EXIF auxiliary keys.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataNamespaceExifAux
	KCGImageMetadataNamespaceExifAux string
	// KCGImageMetadataNamespaceExifEX is the namespace for the exifEX format.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataNamespaceExifEX
	KCGImageMetadataNamespaceExifEX string
	// KCGImageMetadataNamespaceIPTCCore is the namespace for the IPTC format.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataNamespaceIPTCCore
	KCGImageMetadataNamespaceIPTCCore string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataNamespaceIPTCExtension
	KCGImageMetadataNamespaceIPTCExtension string
	// KCGImageMetadataNamespacePhotoshop is the namespace for Photoshop image metadata.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataNamespacePhotoshop
	KCGImageMetadataNamespacePhotoshop string
	// KCGImageMetadataNamespaceTIFF is the namespace for TIFF image metadata.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataNamespaceTIFF
	KCGImageMetadataNamespaceTIFF string
	// KCGImageMetadataNamespaceXMPBasic is the namespace for the Extensible Metadata Platform (XMP) format.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataNamespaceXMPBasic
	KCGImageMetadataNamespaceXMPBasic string
	// KCGImageMetadataNamespaceXMPRights is the namespace for XMP metadata that conveys legal restrictions associated with a resource.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataNamespaceXMPRights
	KCGImageMetadataNamespaceXMPRights string
	// KCGImageMetadataPrefixDublinCore is the prefix string for tags in the Dublin Core Metadata Element Set.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataPrefixDublinCore
	KCGImageMetadataPrefixDublinCore string
	// KCGImageMetadataPrefixExif is the prefix string for tags in the Exchangeable Image File (EXIF) metadata.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataPrefixExif
	KCGImageMetadataPrefixExif string
	// KCGImageMetadataPrefixExifAux is the prefix string for tags in the EXIF auxiliary metadata collection.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataPrefixExifAux
	KCGImageMetadataPrefixExifAux string
	// KCGImageMetadataPrefixExifEX is the prefix string for tags in the exifEX metadata.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataPrefixExifEX
	KCGImageMetadataPrefixExifEX string
	// KCGImageMetadataPrefixIPTCCore is the prefix string for tags in the IPTC metadata.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataPrefixIPTCCore
	KCGImageMetadataPrefixIPTCCore string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataPrefixIPTCExtension
	KCGImageMetadataPrefixIPTCExtension string
	// KCGImageMetadataPrefixPhotoshop is the prefix string for tags in the Photoshop image metadata.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataPrefixPhotoshop
	KCGImageMetadataPrefixPhotoshop string
	// KCGImageMetadataPrefixTIFF is the prefix string for tags in the TIFF image metadata.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataPrefixTIFF
	KCGImageMetadataPrefixTIFF string
	// KCGImageMetadataPrefixXMPBasic is the prefix string for tags in the XMP metadata.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataPrefixXMPBasic
	KCGImageMetadataPrefixXMPBasic string
	// KCGImageMetadataPrefixXMPRights is the prefix string for tags in the XMP metadata that convey legal restrictions for the resource.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataPrefixXMPRights
	KCGImageMetadataPrefixXMPRights string
	// KCGImageMetadataShouldExcludeGPS is a Boolean value that indicates whether to exclude GPS metadata from EXIF data or the corresponding XMP tags.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataShouldExcludeGPS
	KCGImageMetadataShouldExcludeGPS string
	// KCGImageMetadataShouldExcludeXMP is a Boolean value that indicates whether to exclude XMP data from the destination.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageMetadataShouldExcludeXMP
	KCGImageMetadataShouldExcludeXMP string
	// KCGImageProperty8BIMDictionary is a dictionary of key-value pairs for an Adobe Photoshop image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageProperty8BIMDictionary
	KCGImageProperty8BIMDictionary string
	// KCGImageProperty8BIMLayerNames is the layer names for an Adobe Photoshop file.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageProperty8BIMLayerNames
	KCGImageProperty8BIMLayerNames string
	// KCGImageProperty8BIMVersion is the Adobe Photoshop file version.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageProperty8BIMVersion
	KCGImageProperty8BIMVersion string
	// KCGImagePropertyAPNGCanvasPixelHeight is the height of the main image, in pixels.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyAPNGCanvasPixelHeight
	KCGImagePropertyAPNGCanvasPixelHeight string
	// KCGImagePropertyAPNGCanvasPixelWidth is the width of the main image, in pixels.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyAPNGCanvasPixelWidth
	KCGImagePropertyAPNGCanvasPixelWidth string
	// KCGImagePropertyAPNGDelayTime is the number of seconds to wait before displaying the next image in an animated sequence.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyAPNGDelayTime
	KCGImagePropertyAPNGDelayTime string
	// KCGImagePropertyAPNGFrameInfoArray is an array of dictionaries that contain timing information for the image sequence.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyAPNGFrameInfoArray
	KCGImagePropertyAPNGFrameInfoArray string
	// KCGImagePropertyAPNGLoopCount is the number of times that an animated PNG should play through its frames before stopping.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyAPNGLoopCount
	KCGImagePropertyAPNGLoopCount string
	// KCGImagePropertyAPNGUnclampedDelayTime is the number of seconds to wait before displaying the next image in an animated sequence.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyAPNGUnclampedDelayTime
	KCGImagePropertyAPNGUnclampedDelayTime string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyASTCBlockSize
	KCGImagePropertyASTCBlockSize string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyASTCBlockSize4x4
	KCGImagePropertyASTCBlockSize4x4 string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyASTCBlockSize8x8
	KCGImagePropertyASTCBlockSize8x8 string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyASTCEncoder
	KCGImagePropertyASTCEncoder string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyAVISDictionary
	KCGImagePropertyAVISDictionary string
	// KCGImagePropertyAuxiliaryData is an array of dictionaries that contain auxiliary data for the images.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyAuxiliaryData
	KCGImagePropertyAuxiliaryData string
	// KCGImagePropertyAuxiliaryDataType is the type of the auxiliary data.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyAuxiliaryDataType
	KCGImagePropertyAuxiliaryDataType string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyBCEncoder
	KCGImagePropertyBCEncoder string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyBCFormat
	KCGImagePropertyBCFormat string
	// KCGImagePropertyBytesPerRow is the total number of bytes in each row of the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyBytesPerRow
	KCGImagePropertyBytesPerRow string
	// KCGImagePropertyCIFFCameraSerialNumber is the camera serial number.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFCameraSerialNumber
	KCGImagePropertyCIFFCameraSerialNumber string
	// KCGImagePropertyCIFFContinuousDrive is the continuous drive mode.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFContinuousDrive
	KCGImagePropertyCIFFContinuousDrive string
	// KCGImagePropertyCIFFDescription is the camera description.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFDescription
	KCGImagePropertyCIFFDescription string
	// KCGImagePropertyCIFFDictionary is a dictionary of key-value pairs for an image that uses Camera Image File Format (CIFF).
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFDictionary
	KCGImagePropertyCIFFDictionary string
	// KCGImagePropertyCIFFFirmware is the firmware version of the camera.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFFirmware
	KCGImagePropertyCIFFFirmware string
	// KCGImagePropertyCIFFFlashExposureComp is the flash exposure compensation.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFFlashExposureComp
	KCGImagePropertyCIFFFlashExposureComp string
	// KCGImagePropertyCIFFFocusMode is the focus mode.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFFocusMode
	KCGImagePropertyCIFFFocusMode string
	// KCGImagePropertyCIFFImageFileName is the image file name.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFImageFileName
	KCGImagePropertyCIFFImageFileName string
	// KCGImagePropertyCIFFImageName is the image name.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFImageName
	KCGImagePropertyCIFFImageName string
	// KCGImagePropertyCIFFImageSerialNumber is the image serial number.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFImageSerialNumber
	KCGImagePropertyCIFFImageSerialNumber string
	// KCGImagePropertyCIFFLensMaxMM is the maximum lens length in millimeters.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFLensMaxMM
	KCGImagePropertyCIFFLensMaxMM string
	// KCGImagePropertyCIFFLensMinMM is the minimum lens length in millimeters.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFLensMinMM
	KCGImagePropertyCIFFLensMinMM string
	// KCGImagePropertyCIFFLensModel is the lens model.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFLensModel
	KCGImagePropertyCIFFLensModel string
	// KCGImagePropertyCIFFMeasuredEV is the measured exposure value.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFMeasuredEV
	KCGImagePropertyCIFFMeasuredEV string
	// KCGImagePropertyCIFFMeteringMode is the metering mode.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFMeteringMode
	KCGImagePropertyCIFFMeteringMode string
	// KCGImagePropertyCIFFOwnerName is the name of the camera’s owner.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFOwnerName
	KCGImagePropertyCIFFOwnerName string
	// KCGImagePropertyCIFFRecordID is the number of images taken since the camera shipped.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFRecordID
	KCGImagePropertyCIFFRecordID string
	// KCGImagePropertyCIFFReleaseMethod is the method of shutter release—single-shot or continuous.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFReleaseMethod
	KCGImagePropertyCIFFReleaseMethod string
	// KCGImagePropertyCIFFReleaseTiming is the priority for shutter release timing—shutter or focus.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFReleaseTiming
	KCGImagePropertyCIFFReleaseTiming string
	// KCGImagePropertyCIFFSelfTimingTime is the time in milliseconds until shutter release when using the self-timer.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFSelfTimingTime
	KCGImagePropertyCIFFSelfTimingTime string
	// KCGImagePropertyCIFFShootingMode is the shooting mode.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFShootingMode
	KCGImagePropertyCIFFShootingMode string
	// KCGImagePropertyCIFFWhiteBalanceIndex is the white balance index.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyCIFFWhiteBalanceIndex
	KCGImagePropertyCIFFWhiteBalanceIndex string
	// KCGImagePropertyColorModel is the color model of the image, such as RGB, CMYK, grayscale, or Lab.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyColorModel
	KCGImagePropertyColorModel string
	// KCGImagePropertyColorModelCMYK is a Cyan Magenta Yellow Black (CMYK) color model.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyColorModelCMYK
	KCGImagePropertyColorModelCMYK string
	// KCGImagePropertyColorModelGray is a grayscale color model.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyColorModelGray
	KCGImagePropertyColorModelGray string
	// KCGImagePropertyColorModelLab is a Lab color model, where color values contain the amount of light and the amounts of four human-perceivable colors.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyColorModelLab
	KCGImagePropertyColorModelLab string
	// KCGImagePropertyColorModelRGB is a Red Green Blue (RGB) color model.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyColorModelRGB
	KCGImagePropertyColorModelRGB string
	// KCGImagePropertyDNGActiveArea is the rectangle that defines the non-masked pixels of the sensor.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGActiveArea
	KCGImagePropertyDNGActiveArea string
	// KCGImagePropertyDNGAnalogBalance is the analog or digital gain that applies to the stored raw values.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGAnalogBalance
	KCGImagePropertyDNGAnalogBalance string
	// KCGImagePropertyDNGAntiAliasStrength is a hint to the DNG reader about how strong the camera’s antialias filter is.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGAntiAliasStrength
	KCGImagePropertyDNGAntiAliasStrength string
	// KCGImagePropertyDNGAsShotICCProfile is a profile that specifies default color rendering from camera color-space coordinates into the ICC profile space.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGAsShotICCProfile
	KCGImagePropertyDNGAsShotICCProfile string
	// KCGImagePropertyDNGAsShotNeutral is the selected white balance at the time of capture, encoded as the coordinates of a neutral color in linear reference space values.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGAsShotNeutral
	KCGImagePropertyDNGAsShotNeutral string
	// KCGImagePropertyDNGAsShotPreProfileMatrix is a matrix to apply to the camera color-space coordinates before processing values through the ICC profile.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGAsShotPreProfileMatrix
	KCGImagePropertyDNGAsShotPreProfileMatrix string
	// KCGImagePropertyDNGAsShotProfileName is a string containing the name of the “as shot” camera profile, if any.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGAsShotProfileName
	KCGImagePropertyDNGAsShotProfileName string
	// KCGImagePropertyDNGAsShotWhiteXY is the selected white balance at the time of capture, encoded as x-y chromaticity coordinates.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGAsShotWhiteXY
	KCGImagePropertyDNGAsShotWhiteXY string
	// KCGImagePropertyDNGBackwardVersion is the oldest version for which a file is compatible.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGBackwardVersion
	KCGImagePropertyDNGBackwardVersion string
	// KCGImagePropertyDNGBaselineExposure is the amount by which to adjust the zero point of the exposure, specified in EV units.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGBaselineExposure
	KCGImagePropertyDNGBaselineExposure string
	// KCGImagePropertyDNGBaselineExposureOffset is the amount of EV units to add to the baseline exposure during image rendering.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGBaselineExposureOffset
	KCGImagePropertyDNGBaselineExposureOffset string
	// KCGImagePropertyDNGBaselineNoise is the relative noise level of the camera model at an ISO of 100.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGBaselineNoise
	KCGImagePropertyDNGBaselineNoise string
	// KCGImagePropertyDNGBaselineSharpness is the amount of sharpening required for this camera model.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGBaselineSharpness
	KCGImagePropertyDNGBaselineSharpness string
	// KCGImagePropertyDNGBayerGreenSplit is a value that specifies how closely green pixels in the blue/green rows track the green pixels in red/green rows.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGBayerGreenSplit
	KCGImagePropertyDNGBayerGreenSplit string
	// KCGImagePropertyDNGBestQualityScale is the scale factor to apply to the default scale to achieve the best quality image size.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGBestQualityScale
	KCGImagePropertyDNGBestQualityScale string
	// KCGImagePropertyDNGBlackLevel is the zero light encoding level, specified as a repeating pattern.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGBlackLevel
	KCGImagePropertyDNGBlackLevel string
	// KCGImagePropertyDNGBlackLevelDeltaH is the difference between the zero-light encoding level for each column and the baseline zero-light encoding level.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGBlackLevelDeltaH
	KCGImagePropertyDNGBlackLevelDeltaH string
	// KCGImagePropertyDNGBlackLevelDeltaV is the difference between the zero-light encodoing level for each row and the baseline zero-light encoding level.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGBlackLevelDeltaV
	KCGImagePropertyDNGBlackLevelDeltaV string
	// KCGImagePropertyDNGBlackLevelRepeatDim is the repeat pattern size for the black level tag.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGBlackLevelRepeatDim
	KCGImagePropertyDNGBlackLevelRepeatDim string
	// KCGImagePropertyDNGCFALayout is the spatial layout of the CFA.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGCFALayout
	KCGImagePropertyDNGCFALayout string
	// KCGImagePropertyDNGCFAPlaneColor is a mapping between the values in the CFA pattern tag and the plane numbers in linear raw space.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGCFAPlaneColor
	KCGImagePropertyDNGCFAPlaneColor string
	// KCGImagePropertyDNGCalibrationIlluminant1 is the illuminant for the first set of color calibration tags.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGCalibrationIlluminant1
	KCGImagePropertyDNGCalibrationIlluminant1 string
	// KCGImagePropertyDNGCalibrationIlluminant2 is the illuminant for an optional second set of color calibration tags.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGCalibrationIlluminant2
	KCGImagePropertyDNGCalibrationIlluminant2 string
	// KCGImagePropertyDNGCameraCalibration1 is a matrix that transforms reference camera native space values to camera-native space values under the first calibration illuminant.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGCameraCalibration1
	KCGImagePropertyDNGCameraCalibration1 string
	// KCGImagePropertyDNGCameraCalibration2 is a matrix that transforms reference camera native space values to camera-native space values under the second calibration illuminant.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGCameraCalibration2
	KCGImagePropertyDNGCameraCalibration2 string
	// KCGImagePropertyDNGCameraCalibrationSignature is a string to match against the profile calibration signature for the selected camera profile.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGCameraCalibrationSignature
	KCGImagePropertyDNGCameraCalibrationSignature string
	// KCGImagePropertyDNGCameraSerialNumber is the camera serial number.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGCameraSerialNumber
	KCGImagePropertyDNGCameraSerialNumber string
	// KCGImagePropertyDNGChromaBlurRadius is a hint to the DNG reader about how much chroma blur to apply to the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGChromaBlurRadius
	KCGImagePropertyDNGChromaBlurRadius string
	// KCGImagePropertyDNGColorMatrix1 is a transformation matrix that converts XYZ values to reference camera native color spaces, under the first calibration illuminant.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGColorMatrix1
	KCGImagePropertyDNGColorMatrix1 string
	// KCGImagePropertyDNGColorMatrix2 is a transformation matrix that converts XYZ values to reference camera native color spaces, under the second calibration illuminant.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGColorMatrix2
	KCGImagePropertyDNGColorMatrix2 string
	// KCGImagePropertyDNGColorimetricReference is the colorimetric reference for the CIE XYZ values.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGColorimetricReference
	KCGImagePropertyDNGColorimetricReference string
	// KCGImagePropertyDNGCurrentICCProfile is a profile that specifies default color rendering from camera color-space coordinates into the ICC profile space.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGCurrentICCProfile
	KCGImagePropertyDNGCurrentICCProfile string
	// KCGImagePropertyDNGCurrentPreProfileMatrix is a matrix to apply to the current camera color-space coordinates before processing values through the ICC profile.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGCurrentPreProfileMatrix
	KCGImagePropertyDNGCurrentPreProfileMatrix string
	// KCGImagePropertyDNGDefaultBlackRender is a hint to the raw converter about how to handle the black point during rendering.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGDefaultBlackRender
	KCGImagePropertyDNGDefaultBlackRender string
	// KCGImagePropertyDNGDefaultCropOrigin is the origin of the final image area, relative to the top-left corner of the active area rectangle.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGDefaultCropOrigin
	KCGImagePropertyDNGDefaultCropOrigin string
	// KCGImagePropertyDNGDefaultCropSize is the size of the final image area, in raw image coordinates.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGDefaultCropSize
	KCGImagePropertyDNGDefaultCropSize string
	// KCGImagePropertyDNGDefaultScale is the default scale factors for each direction to convert the image to square pixels.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGDefaultScale
	KCGImagePropertyDNGDefaultScale string
	// KCGImagePropertyDNGDefaultUserCrop is a default user-crop rectangle in relative coordinates.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGDefaultUserCrop
	KCGImagePropertyDNGDefaultUserCrop string
	// KCGImagePropertyDNGDictionary is a dictionary of key-value pairs for an image that uses the Digital Negative (DNG) archival format.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGDictionary
	KCGImagePropertyDNGDictionary string
	// KCGImagePropertyDNGExtraCameraProfiles is a list of file offsets to extra camera profiles.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGExtraCameraProfiles
	KCGImagePropertyDNGExtraCameraProfiles string
	// KCGImagePropertyDNGFixVignetteRadial is an opcode to apply a gain function to an image to correct vignetting.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGFixVignetteRadial
	KCGImagePropertyDNGFixVignetteRadial string
	// KCGImagePropertyDNGForwardMatrix1 is a matrix that maps white balanced camera colors to XYZ D50 colors.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGForwardMatrix1
	KCGImagePropertyDNGForwardMatrix1 string
	// KCGImagePropertyDNGForwardMatrix2 is a matrix that maps white balanced camera colors to XYZ D50 colors.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGForwardMatrix2
	KCGImagePropertyDNGForwardMatrix2 string
	// KCGImagePropertyDNGLensInfo is information about the lens used for the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGLensInfo
	KCGImagePropertyDNGLensInfo string
	// KCGImagePropertyDNGLinearResponseLimit is the fraction of the encoding range, above which the response may become significantly non-linear.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGLinearResponseLimit
	KCGImagePropertyDNGLinearResponseLimit string
	// KCGImagePropertyDNGLinearizationTable is a lookup table that maps stored values into linear values.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGLinearizationTable
	KCGImagePropertyDNGLinearizationTable string
	// KCGImagePropertyDNGLocalizedCameraModel is the localized camera model name.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGLocalizedCameraModel
	KCGImagePropertyDNGLocalizedCameraModel string
	// KCGImagePropertyDNGMakerNoteSafety is a Boolean value that tells the DNG reader whether the EXIF MakerNote tag is safe to preserve.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGMakerNoteSafety
	KCGImagePropertyDNGMakerNoteSafety string
	// KCGImagePropertyDNGMaskedAreas is a list of non-overlapping rectangles that contain fully masked pixels in the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGMaskedAreas
	KCGImagePropertyDNGMaskedAreas string
	// KCGImagePropertyDNGNewRawImageDigest is an MD5 digest of the raw image data.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGNewRawImageDigest
	KCGImagePropertyDNGNewRawImageDigest string
	// KCGImagePropertyDNGNoiseProfile is the amount of noise in the raw image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGNoiseProfile
	KCGImagePropertyDNGNoiseProfile string
	// KCGImagePropertyDNGNoiseReductionApplied is the amount of noise reduction applied to the raw data on a scale of 0.0 to 1.0.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGNoiseReductionApplied
	KCGImagePropertyDNGNoiseReductionApplied string
	// KCGImagePropertyDNGOpcodeList1 is the list of opcodes to apply to the raw image, as read directly from the file.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGOpcodeList1
	KCGImagePropertyDNGOpcodeList1 string
	// KCGImagePropertyDNGOpcodeList2 is tHe list of opcodes to apply to the raw image, after mapping it to linear reference values.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGOpcodeList2
	KCGImagePropertyDNGOpcodeList2 string
	// KCGImagePropertyDNGOpcodeList3 is the list of opcodes to apply to the raw image, after demosaicing it.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGOpcodeList3
	KCGImagePropertyDNGOpcodeList3 string
	// KCGImagePropertyDNGOriginalBestQualityFinalSize is the best-quality final size of the larger original file that was the source of this proxy.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGOriginalBestQualityFinalSize
	KCGImagePropertyDNGOriginalBestQualityFinalSize string
	// KCGImagePropertyDNGOriginalDefaultCropSize is the default crop size of the larger original file that was the source of this proxy.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGOriginalDefaultCropSize
	KCGImagePropertyDNGOriginalDefaultCropSize string
	// KCGImagePropertyDNGOriginalDefaultFinalSize is tHe default final size of the larger original file that was the source of this proxy.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGOriginalDefaultFinalSize
	KCGImagePropertyDNGOriginalDefaultFinalSize string
	// KCGImagePropertyDNGOriginalRawFileData is the compressed contents of the original raw file.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGOriginalRawFileData
	KCGImagePropertyDNGOriginalRawFileData string
	// KCGImagePropertyDNGOriginalRawFileDigest is an MD5 digest of the data stored for the original raw file data.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGOriginalRawFileDigest
	KCGImagePropertyDNGOriginalRawFileDigest string
	// KCGImagePropertyDNGOriginalRawFileName is the file name of the original raw file.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGOriginalRawFileName
	KCGImagePropertyDNGOriginalRawFileName string
	// KCGImagePropertyDNGPreviewApplicationName is the name of the app that created the preview stored in the IFD.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGPreviewApplicationName
	KCGImagePropertyDNGPreviewApplicationName string
	// KCGImagePropertyDNGPreviewApplicationVersion is the version number of the app that created the preview stored in the IFD.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGPreviewApplicationVersion
	KCGImagePropertyDNGPreviewApplicationVersion string
	// KCGImagePropertyDNGPreviewColorSpace is the color space associated with the rendered preview.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGPreviewColorSpace
	KCGImagePropertyDNGPreviewColorSpace string
	// KCGImagePropertyDNGPreviewDateTime is the date and time for the render of the preview.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGPreviewDateTime
	KCGImagePropertyDNGPreviewDateTime string
	// KCGImagePropertyDNGPreviewSettingsDigest is a unique ID of the conversion settings used to render the preview.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGPreviewSettingsDigest
	KCGImagePropertyDNGPreviewSettingsDigest string
	// KCGImagePropertyDNGPreviewSettingsName is the name of the conversion settings for the preview.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGPreviewSettingsName
	KCGImagePropertyDNGPreviewSettingsName string
	// KCGImagePropertyDNGPrivateData is private data that manufacturers may store with an image and use in their own converters.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGPrivateData
	KCGImagePropertyDNGPrivateData string
	// KCGImagePropertyDNGProfileCalibrationSignature is a string that describes the calibration for the current profile.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGProfileCalibrationSignature
	KCGImagePropertyDNGProfileCalibrationSignature string
	// KCGImagePropertyDNGProfileCopyright is the copyright information for the camera profile.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGProfileCopyright
	KCGImagePropertyDNGProfileCopyright string
	// KCGImagePropertyDNGProfileEmbedPolicy is the usage rules for the camera profile.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGProfileEmbedPolicy
	KCGImagePropertyDNGProfileEmbedPolicy string
	// KCGImagePropertyDNGProfileHueSatMapData1 is the data for the first hue/saturation/value mapping table.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGProfileHueSatMapData1
	KCGImagePropertyDNGProfileHueSatMapData1 string
	// KCGImagePropertyDNGProfileHueSatMapData2 is the data for the second hue/saturation/value mapping table.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGProfileHueSatMapData2
	KCGImagePropertyDNGProfileHueSatMapData2 string
	// KCGImagePropertyDNGProfileHueSatMapDims is the number of input samples in each dimension of the hue/saturation/value mapping tables.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGProfileHueSatMapDims
	KCGImagePropertyDNGProfileHueSatMapDims string
	// KCGImagePropertyDNGProfileHueSatMapEncoding is the encoding option to use when indexing into a 3D look table during raw conversion.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGProfileHueSatMapEncoding
	KCGImagePropertyDNGProfileHueSatMapEncoding string
	// KCGImagePropertyDNGProfileLookTableData is the default “look” table to apply when processing the image as a starting point for user adjustment.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGProfileLookTableData
	KCGImagePropertyDNGProfileLookTableData string
	// KCGImagePropertyDNGProfileLookTableDims is the number of input samples in each dimentsion of a default “look” table.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGProfileLookTableDims
	KCGImagePropertyDNGProfileLookTableDims string
	// KCGImagePropertyDNGProfileLookTableEncoding is the encoding option to use when indexing into a 3D look table during raw conversion.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGProfileLookTableEncoding
	KCGImagePropertyDNGProfileLookTableEncoding string
	// KCGImagePropertyDNGProfileName is a string containing the name of the camera profile.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGProfileName
	KCGImagePropertyDNGProfileName string
	// KCGImagePropertyDNGProfileToneCurve is the default tone curve to apply when processing the image as a starting point for user adjustments.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGProfileToneCurve
	KCGImagePropertyDNGProfileToneCurve string
	// KCGImagePropertyDNGRawDataUniqueID is a 16-byte unique identifier for the raw image data.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGRawDataUniqueID
	KCGImagePropertyDNGRawDataUniqueID string
	// KCGImagePropertyDNGRawImageDigest is a modified MD5 digest of the raw image data.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGRawImageDigest
	KCGImagePropertyDNGRawImageDigest string
	// KCGImagePropertyDNGRawToPreviewGain is the gain between the main raw IFD and the preview IFD that contains this tag.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGRawToPreviewGain
	KCGImagePropertyDNGRawToPreviewGain string
	// KCGImagePropertyDNGReductionMatrix1 is a reduction matrix that converts color camera-native space values to XYZ values, under the first calibration illuminant.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGReductionMatrix1
	KCGImagePropertyDNGReductionMatrix1 string
	// KCGImagePropertyDNGReductionMatrix2 is a reduction matrix that converts color camera-native space values to XYZ values, under the second calibration illuminant.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGReductionMatrix2
	KCGImagePropertyDNGReductionMatrix2 string
	// KCGImagePropertyDNGRowInterleaveFactor is the number of interleaved fields for the rows of the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGRowInterleaveFactor
	KCGImagePropertyDNGRowInterleaveFactor string
	// KCGImagePropertyDNGShadowScale is a tag that Adobe Camera Raw uses to control the sensitivity of its Shadows slider.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGShadowScale
	KCGImagePropertyDNGShadowScale string
	// KCGImagePropertyDNGSubTileBlockSize is the size of rectangular blocks that tiles use to group pixels.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGSubTileBlockSize
	KCGImagePropertyDNGSubTileBlockSize string
	// KCGImagePropertyDNGUniqueCameraModel is a unique, nonlocalized name for the camera model.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGUniqueCameraModel
	KCGImagePropertyDNGUniqueCameraModel string
	// KCGImagePropertyDNGVersion is an encoding of the four-tier version number.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGVersion
	KCGImagePropertyDNGVersion string
	// KCGImagePropertyDNGWarpFisheye is an opcode to unwrap an image captued with a fisheye lens and map it to a perspective projection.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGWarpFisheye
	KCGImagePropertyDNGWarpFisheye string
	// KCGImagePropertyDNGWarpRectilinear is an opcode to apply a warp to an image to correct for geometric distortion and lateral chromatic aberration for rectilinear lenses.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGWarpRectilinear
	KCGImagePropertyDNGWarpRectilinear string
	// KCGImagePropertyDNGWhiteLevel is the saturated encoding level for the raw sample values.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDNGWhiteLevel
	KCGImagePropertyDNGWhiteLevel string
	// KCGImagePropertyDPIHeight is the resolution, in dots per inch, in the y dimension.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDPIHeight
	KCGImagePropertyDPIHeight string
	// KCGImagePropertyDPIWidth is the resolution, in dots per inch, in the x dimension.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDPIWidth
	KCGImagePropertyDPIWidth string
	// KCGImagePropertyDepth is the number of bits in the color sample of a pixel.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyDepth
	KCGImagePropertyDepth string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyEncoder
	KCGImagePropertyEncoder string
	// KCGImagePropertyExifApertureValue is the aperture value.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifApertureValue
	KCGImagePropertyExifApertureValue string
	// KCGImagePropertyExifAuxDictionary is an auxiliary dictionary of key-value pairs for an image that uses Exchangeable Image File Format (EXIF).
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifAuxDictionary
	KCGImagePropertyExifAuxDictionary string
	// KCGImagePropertyExifAuxFirmware is firmware information.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifAuxFirmware
	KCGImagePropertyExifAuxFirmware string
	// KCGImagePropertyExifAuxFlashCompensation is flash compensation.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifAuxFlashCompensation
	KCGImagePropertyExifAuxFlashCompensation string
	// KCGImagePropertyExifAuxImageNumber is the image number.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifAuxImageNumber
	KCGImagePropertyExifAuxImageNumber string
	// KCGImagePropertyExifAuxLensID is the lens ID.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifAuxLensID
	KCGImagePropertyExifAuxLensID string
	// KCGImagePropertyExifAuxLensInfo is lens information.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifAuxLensInfo
	KCGImagePropertyExifAuxLensInfo string
	// KCGImagePropertyExifAuxLensModel is the lens model.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifAuxLensModel
	KCGImagePropertyExifAuxLensModel string
	// KCGImagePropertyExifAuxLensSerialNumber is the lens serial number.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifAuxLensSerialNumber
	KCGImagePropertyExifAuxLensSerialNumber string
	// KCGImagePropertyExifAuxOwnerName is the owner name.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifAuxOwnerName
	KCGImagePropertyExifAuxOwnerName string
	// KCGImagePropertyExifAuxSerialNumber is the serial number.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifAuxSerialNumber
	KCGImagePropertyExifAuxSerialNumber string
	// KCGImagePropertyExifBodySerialNumber is a string with the serial number of the camera.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifBodySerialNumber
	KCGImagePropertyExifBodySerialNumber string
	// KCGImagePropertyExifBrightnessValue is the brightness value.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifBrightnessValue
	KCGImagePropertyExifBrightnessValue string
	// KCGImagePropertyExifCFAPattern is the color filter array (CFA) pattern, which is the geometric pattern of the image sensor for a 1-chip color sensor area.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifCFAPattern
	KCGImagePropertyExifCFAPattern string
	// KCGImagePropertyExifCameraOwnerName is a string with the name of the camera’s owner.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifCameraOwnerName
	KCGImagePropertyExifCameraOwnerName string
	// KCGImagePropertyExifColorSpace is the color space.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifColorSpace
	KCGImagePropertyExifColorSpace string
	// KCGImagePropertyExifComponentsConfiguration is the components configuration for compressed data.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifComponentsConfiguration
	KCGImagePropertyExifComponentsConfiguration string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifCompositeImage
	KCGImagePropertyExifCompositeImage string
	// KCGImagePropertyExifCompressedBitsPerPixel is the bits per pixel of the compression mode.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifCompressedBitsPerPixel
	KCGImagePropertyExifCompressedBitsPerPixel string
	// KCGImagePropertyExifContrast is the contrast setting.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifContrast
	KCGImagePropertyExifContrast string
	// KCGImagePropertyExifCustomRendered is special rendering performed on the image data.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifCustomRendered
	KCGImagePropertyExifCustomRendered string
	// KCGImagePropertyExifDateTimeDigitized is the digitized date and time.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifDateTimeDigitized
	KCGImagePropertyExifDateTimeDigitized string
	// KCGImagePropertyExifDateTimeOriginal is the original date and time.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifDateTimeOriginal
	KCGImagePropertyExifDateTimeOriginal string
	// KCGImagePropertyExifDeviceSettingDescription is for a particular camera mode, indicates the conditions for taking the picture.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifDeviceSettingDescription
	KCGImagePropertyExifDeviceSettingDescription string
	// KCGImagePropertyExifDictionary is a dictionary of key-value pairs for an image that uses Exchangeable Image File Format (EXIF).
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifDictionary
	KCGImagePropertyExifDictionary string
	// KCGImagePropertyExifDigitalZoomRatio is the digital zoom ratio.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifDigitalZoomRatio
	KCGImagePropertyExifDigitalZoomRatio string
	// KCGImagePropertyExifExposureBiasValue is the exposure bias value.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifExposureBiasValue
	KCGImagePropertyExifExposureBiasValue string
	// KCGImagePropertyExifExposureIndex is the selected exposure index.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifExposureIndex
	KCGImagePropertyExifExposureIndex string
	// KCGImagePropertyExifExposureMode is the exposure mode setting.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifExposureMode
	KCGImagePropertyExifExposureMode string
	// KCGImagePropertyExifExposureProgram is the exposure program.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifExposureProgram
	KCGImagePropertyExifExposureProgram string
	// KCGImagePropertyExifExposureTime is the exposure time.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifExposureTime
	KCGImagePropertyExifExposureTime string
	// KCGImagePropertyExifFNumber is the F-number.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifFNumber
	KCGImagePropertyExifFNumber string
	// KCGImagePropertyExifFileSource is the image source.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifFileSource
	KCGImagePropertyExifFileSource string
	// KCGImagePropertyExifFlash is the flash status when the image was shot.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifFlash
	KCGImagePropertyExifFlash string
	// KCGImagePropertyExifFlashEnergy is the strobe energy when the image was captured, in beam candle power seconds.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifFlashEnergy
	KCGImagePropertyExifFlashEnergy string
	// KCGImagePropertyExifFlashPixVersion is the FlashPix version supported by an FPXR file.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifFlashPixVersion
	KCGImagePropertyExifFlashPixVersion string
	// KCGImagePropertyExifFocalLenIn35mmFilm is the equivalent focal length in 35 mm film.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifFocalLenIn35mmFilm
	KCGImagePropertyExifFocalLenIn35mmFilm string
	// KCGImagePropertyExifFocalLength is the focal length.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifFocalLength
	KCGImagePropertyExifFocalLength string
	// KCGImagePropertyExifFocalPlaneResolutionUnit is the unit of measurement for the focal plane x and y resolutions.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifFocalPlaneResolutionUnit
	KCGImagePropertyExifFocalPlaneResolutionUnit string
	// KCGImagePropertyExifFocalPlaneXResolution is the number of image-width pixels (x-axis) per focal plane resolution unit.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifFocalPlaneXResolution
	KCGImagePropertyExifFocalPlaneXResolution string
	// KCGImagePropertyExifFocalPlaneYResolution is the number of image-height pixels (y-axis) per focal plane resolution unit.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifFocalPlaneYResolution
	KCGImagePropertyExifFocalPlaneYResolution string
	// KCGImagePropertyExifGainControl is the gain adjustment setting.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifGainControl
	KCGImagePropertyExifGainControl string
	// KCGImagePropertyExifGamma is the gamma setting.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifGamma
	KCGImagePropertyExifGamma string
	// KCGImagePropertyExifISOSpeed is the ISO speed setting used to capture the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifISOSpeed
	KCGImagePropertyExifISOSpeed string
	// KCGImagePropertyExifISOSpeedLatitudeyyy is the ISO speed latitude yyy value.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifISOSpeedLatitudeyyy
	KCGImagePropertyExifISOSpeedLatitudeyyy string
	// KCGImagePropertyExifISOSpeedLatitudezzz is the ISO speed latitude zzz value.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifISOSpeedLatitudezzz
	KCGImagePropertyExifISOSpeedLatitudezzz string
	// KCGImagePropertyExifISOSpeedRatings is the ISO speed ratings.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifISOSpeedRatings
	KCGImagePropertyExifISOSpeedRatings string
	// KCGImagePropertyExifImageUniqueID is the unique ID of the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifImageUniqueID
	KCGImagePropertyExifImageUniqueID string
	// KCGImagePropertyExifLensMake is a string with the name of the lens manufacturer.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifLensMake
	KCGImagePropertyExifLensMake string
	// KCGImagePropertyExifLensModel is a string with the lens model information.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifLensModel
	KCGImagePropertyExifLensModel string
	// KCGImagePropertyExifLensSerialNumber is a string with the lens’s serial number.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifLensSerialNumber
	KCGImagePropertyExifLensSerialNumber string
	// KCGImagePropertyExifLensSpecification is the specification information for the camera lens.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifLensSpecification
	KCGImagePropertyExifLensSpecification string
	// KCGImagePropertyExifLightSource is the light source.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifLightSource
	KCGImagePropertyExifLightSource string
	// KCGImagePropertyExifMakerNote is information specified by the camera manufacturer.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifMakerNote
	KCGImagePropertyExifMakerNote string
	// KCGImagePropertyExifMaxApertureValue is the maximum aperture value.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifMaxApertureValue
	KCGImagePropertyExifMaxApertureValue string
	// KCGImagePropertyExifMeteringMode is the metering mode.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifMeteringMode
	KCGImagePropertyExifMeteringMode string
	// KCGImagePropertyExifOECF is the opto-electric conversion function (OECF) that defines the relationship between the optical input of the camera and the resulting image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifOECF
	KCGImagePropertyExifOECF string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifOffsetTime
	KCGImagePropertyExifOffsetTime string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifOffsetTimeDigitized
	KCGImagePropertyExifOffsetTimeDigitized string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifOffsetTimeOriginal
	KCGImagePropertyExifOffsetTimeOriginal string
	// KCGImagePropertyExifPixelXDimension is the x dimension of a pixel.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifPixelXDimension
	KCGImagePropertyExifPixelXDimension string
	// KCGImagePropertyExifPixelYDimension is the y dimension of a pixel.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifPixelYDimension
	KCGImagePropertyExifPixelYDimension string
	// KCGImagePropertyExifRecommendedExposureIndex is the recommended exposure index.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifRecommendedExposureIndex
	KCGImagePropertyExifRecommendedExposureIndex string
	// KCGImagePropertyExifRelatedSoundFile is a sound file related to the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifRelatedSoundFile
	KCGImagePropertyExifRelatedSoundFile string
	// KCGImagePropertyExifSaturation is the saturation setting.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifSaturation
	KCGImagePropertyExifSaturation string
	// KCGImagePropertyExifSceneCaptureType is the scene capture type; for example, standard, landscape, portrait, or night.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifSceneCaptureType
	KCGImagePropertyExifSceneCaptureType string
	// KCGImagePropertyExifSceneType is the scene type.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifSceneType
	KCGImagePropertyExifSceneType string
	// KCGImagePropertyExifSensingMethod is the sensor type of the camera or input device.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifSensingMethod
	KCGImagePropertyExifSensingMethod string
	// KCGImagePropertyExifSensitivityType is the type of sensitivity data stored for the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifSensitivityType
	KCGImagePropertyExifSensitivityType string
	// KCGImagePropertyExifSharpness is the sharpness setting.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifSharpness
	KCGImagePropertyExifSharpness string
	// KCGImagePropertyExifShutterSpeedValue is the shutter speed value.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifShutterSpeedValue
	KCGImagePropertyExifShutterSpeedValue string
	// KCGImagePropertyExifSourceExposureTimesOfCompositeImage is the exposure times for composite images.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifSourceExposureTimesOfCompositeImage
	KCGImagePropertyExifSourceExposureTimesOfCompositeImage string
	// KCGImagePropertyExifSourceImageNumberOfCompositeImage is the number of images that make up a composite image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifSourceImageNumberOfCompositeImage
	KCGImagePropertyExifSourceImageNumberOfCompositeImage string
	// KCGImagePropertyExifSpatialFrequencyResponse is the spatial frequency table and spatial frequency response values in the width, height, and diagonal directions.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifSpatialFrequencyResponse
	KCGImagePropertyExifSpatialFrequencyResponse string
	// KCGImagePropertyExifSpectralSensitivity is the spectral sensitivity of each channel.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifSpectralSensitivity
	KCGImagePropertyExifSpectralSensitivity string
	// KCGImagePropertyExifStandardOutputSensitivity is the sensitivity data for the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifStandardOutputSensitivity
	KCGImagePropertyExifStandardOutputSensitivity string
	// KCGImagePropertyExifSubjectArea is the subject area.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifSubjectArea
	KCGImagePropertyExifSubjectArea string
	// KCGImagePropertyExifSubjectDistRange is the distance to the subject.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifSubjectDistRange
	KCGImagePropertyExifSubjectDistRange string
	// KCGImagePropertyExifSubjectDistance is the distance to the subject, in meters.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifSubjectDistance
	KCGImagePropertyExifSubjectDistance string
	// KCGImagePropertyExifSubjectLocation is the location of the image’s primary subject.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifSubjectLocation
	KCGImagePropertyExifSubjectLocation string
	// KCGImagePropertyExifSubsecTime is the fraction of seconds for the date and time tag.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifSubsecTime
	KCGImagePropertyExifSubsecTime string
	// KCGImagePropertyExifSubsecTimeDigitized is the fraction of seconds for the digitized date and time tag.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifSubsecTimeDigitized
	KCGImagePropertyExifSubsecTimeDigitized string
	// KCGImagePropertyExifSubsecTimeOriginal is the fraction of seconds for the original date and time tag.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifSubsecTimeOriginal
	KCGImagePropertyExifSubsecTimeOriginal string
	// KCGImagePropertyExifUserComment is a user comment.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifUserComment
	KCGImagePropertyExifUserComment string
	// KCGImagePropertyExifVersion is the EXIF version.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifVersion
	KCGImagePropertyExifVersion string
	// KCGImagePropertyExifWhiteBalance is the white balance mode.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifWhiteBalance
	KCGImagePropertyExifWhiteBalance string
	// KCGImagePropertyFileContentsDictionary is a dictionary of properties related to the image’s on-disk file.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyFileContentsDictionary
	KCGImagePropertyFileContentsDictionary string
	// KCGImagePropertyFileSize is the size of the image file in bytes, if known.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyFileSize
	KCGImagePropertyFileSize string
	// KCGImagePropertyGIFCanvasPixelHeight is the height of the main image, in pixels.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGIFCanvasPixelHeight
	KCGImagePropertyGIFCanvasPixelHeight string
	// KCGImagePropertyGIFCanvasPixelWidth is the width of the main image, in pixels.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGIFCanvasPixelWidth
	KCGImagePropertyGIFCanvasPixelWidth string
	// KCGImagePropertyGIFDelayTime is the number of seconds to wait before displaying the next image in an animated sequence, clamped to a minimum of 100 milliseconds.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGIFDelayTime
	KCGImagePropertyGIFDelayTime string
	// KCGImagePropertyGIFDictionary is a dictionary of key-value pairs for an image that uses Graphics Interchange Format (GIF).
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGIFDictionary
	KCGImagePropertyGIFDictionary string
	// KCGImagePropertyGIFFrameInfoArray is an array of dictionaries that contain timing information for the image sequence.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGIFFrameInfoArray
	KCGImagePropertyGIFFrameInfoArray string
	// KCGImagePropertyGIFHasGlobalColorMap is a Boolean value that indicates whether the GIF has a global color map.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGIFHasGlobalColorMap
	KCGImagePropertyGIFHasGlobalColorMap string
	// KCGImagePropertyGIFImageColorMap is the image color map.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGIFImageColorMap
	KCGImagePropertyGIFImageColorMap string
	// KCGImagePropertyGIFLoopCount is the number of times to repeat an animated sequence.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGIFLoopCount
	KCGImagePropertyGIFLoopCount string
	// KCGImagePropertyGIFUnclampedDelayTime is the number of seconds to wait before displaying the next image in an animated sequence.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGIFUnclampedDelayTime
	KCGImagePropertyGIFUnclampedDelayTime string
	// KCGImagePropertyGPSAltitude is the altitude.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSAltitude
	KCGImagePropertyGPSAltitude string
	// KCGImagePropertyGPSAltitudeRef is the altitude point of reference.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSAltitudeRef
	KCGImagePropertyGPSAltitudeRef string
	// KCGImagePropertyGPSAreaInformation is the name of the GPS area.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSAreaInformation
	KCGImagePropertyGPSAreaInformation string
	// KCGImagePropertyGPSDOP is the degree of precision (DOP) of the data.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSDOP
	KCGImagePropertyGPSDOP string
	// KCGImagePropertyGPSDateStamp is the date and time information relative to Coordinated Universal Time (UTC).
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSDateStamp
	KCGImagePropertyGPSDateStamp string
	// KCGImagePropertyGPSDestBearing is the bearing to the destination point.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSDestBearing
	KCGImagePropertyGPSDestBearing string
	// KCGImagePropertyGPSDestBearingRef is the reference for giving the bearing to the destination point.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSDestBearingRef
	KCGImagePropertyGPSDestBearingRef string
	// KCGImagePropertyGPSDestDistance is the distance to the destination point.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSDestDistance
	KCGImagePropertyGPSDestDistance string
	// KCGImagePropertyGPSDestDistanceRef is the units for expressing the distance to the destination point.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSDestDistanceRef
	KCGImagePropertyGPSDestDistanceRef string
	// KCGImagePropertyGPSDestLatitude is the latitude of the destination point.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSDestLatitude
	KCGImagePropertyGPSDestLatitude string
	// KCGImagePropertyGPSDestLatitudeRef is an indication of whether the latitude of the destination point is northern or southern.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSDestLatitudeRef
	KCGImagePropertyGPSDestLatitudeRef string
	// KCGImagePropertyGPSDestLongitude is the longitude of the destination point.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSDestLongitude
	KCGImagePropertyGPSDestLongitude string
	// KCGImagePropertyGPSDestLongitudeRef is an indication of whether the longitude of the destination point is east or west.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSDestLongitudeRef
	KCGImagePropertyGPSDestLongitudeRef string
	// KCGImagePropertyGPSDictionary is a dictionary of key-value pairs for an image that has Global Positioning System (GPS) information.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSDictionary
	KCGImagePropertyGPSDictionary string
	// KCGImagePropertyGPSDifferental is an indication of whether differential correction is applied to the GPS receiver.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSDifferental
	KCGImagePropertyGPSDifferental string
	// KCGImagePropertyGPSHPositioningError is the horizontal error in the GPS position.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSHPositioningError
	KCGImagePropertyGPSHPositioningError string
	// KCGImagePropertyGPSImgDirection is the direction of the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSImgDirection
	KCGImagePropertyGPSImgDirection string
	// KCGImagePropertyGPSImgDirectionRef is the reference for the direction of the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSImgDirectionRef
	KCGImagePropertyGPSImgDirectionRef string
	// KCGImagePropertyGPSLatitude is the latitude.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSLatitude
	KCGImagePropertyGPSLatitude string
	// KCGImagePropertyGPSLatitudeRef is an indication of whether the latitude is north or south.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSLatitudeRef
	KCGImagePropertyGPSLatitudeRef string
	// KCGImagePropertyGPSLongitude is the longitude.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSLongitude
	KCGImagePropertyGPSLongitude string
	// KCGImagePropertyGPSLongitudeRef is an indication of whether the longitude is east or west.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSLongitudeRef
	KCGImagePropertyGPSLongitudeRef string
	// KCGImagePropertyGPSMapDatum is the geodetic survey data used by the GPS receiver.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSMapDatum
	KCGImagePropertyGPSMapDatum string
	// KCGImagePropertyGPSMeasureMode is the measurement mode.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSMeasureMode
	KCGImagePropertyGPSMeasureMode string
	// KCGImagePropertyGPSProcessingMethod is the name of the method used to find a location.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSProcessingMethod
	KCGImagePropertyGPSProcessingMethod string
	// KCGImagePropertyGPSSatellites is the satellites used for GPS measurements.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSSatellites
	KCGImagePropertyGPSSatellites string
	// KCGImagePropertyGPSSpeed is the GPS receiver’s speed of movement.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSSpeed
	KCGImagePropertyGPSSpeed string
	// KCGImagePropertyGPSSpeedRef is the unit for expressing the GPS receiver’s speed of movement.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSSpeedRef
	KCGImagePropertyGPSSpeedRef string
	// KCGImagePropertyGPSStatus is the status of the GPS receiver.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSStatus
	KCGImagePropertyGPSStatus string
	// KCGImagePropertyGPSTimeStamp is the time in UTC (Coordinated Universal Time).
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSTimeStamp
	KCGImagePropertyGPSTimeStamp string
	// KCGImagePropertyGPSTrack is the direction of GPS receiver’s movement.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSTrack
	KCGImagePropertyGPSTrack string
	// KCGImagePropertyGPSTrackRef is the reference for the direction of GPS receiver’s movement.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSTrackRef
	KCGImagePropertyGPSTrackRef string
	// KCGImagePropertyGPSVersion is the GPS version information.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGPSVersion
	KCGImagePropertyGPSVersion string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGroupImageBaseline
	KCGImagePropertyGroupImageBaseline string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGroupImageDisparityAdjustment
	KCGImagePropertyGroupImageDisparityAdjustment string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGroupImageIndexLeft
	KCGImagePropertyGroupImageIndexLeft string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGroupImageIndexMonoscopic
	KCGImagePropertyGroupImageIndexMonoscopic string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGroupImageIndexRight
	KCGImagePropertyGroupImageIndexRight string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGroupImageIsAlternateImage
	KCGImagePropertyGroupImageIsAlternateImage string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGroupImageIsLeftImage
	KCGImagePropertyGroupImageIsLeftImage string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGroupImageIsMonoscopicImage
	KCGImagePropertyGroupImageIsMonoscopicImage string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGroupImageIsRightImage
	KCGImagePropertyGroupImageIsRightImage string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGroupImageStereoAggressors
	KCGImagePropertyGroupImageStereoAggressors string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGroupImagesAlternate
	KCGImagePropertyGroupImagesAlternate string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGroupIndex
	KCGImagePropertyGroupIndex string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGroupMonoscopicImageLocation
	KCGImagePropertyGroupMonoscopicImageLocation string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGroupType
	KCGImagePropertyGroupType string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGroupTypeAlternate
	KCGImagePropertyGroupTypeAlternate string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGroupTypeStereoPair
	KCGImagePropertyGroupTypeStereoPair string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyGroups
	KCGImagePropertyGroups string
	// KCGImagePropertyHEICSCanvasPixelHeight is the height of the main image, in pixels.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyHEICSCanvasPixelHeight
	KCGImagePropertyHEICSCanvasPixelHeight string
	// KCGImagePropertyHEICSCanvasPixelWidth is the width of the main image, in pixels.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyHEICSCanvasPixelWidth
	KCGImagePropertyHEICSCanvasPixelWidth string
	// KCGImagePropertyHEICSDelayTime is the number of seconds to wait before displaying the next image in the sequence, clamped to a minimum of `0.1` seconds.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyHEICSDelayTime
	KCGImagePropertyHEICSDelayTime string
	// KCGImagePropertyHEICSDictionary is a dictionary of properties related to an HEIC container.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyHEICSDictionary
	KCGImagePropertyHEICSDictionary string
	// KCGImagePropertyHEICSFrameInfoArray is an array of dictionaries that contain timing information for the image sequence.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyHEICSFrameInfoArray
	KCGImagePropertyHEICSFrameInfoArray string
	// KCGImagePropertyHEICSLoopCount is the number of times to play the sequence.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyHEICSLoopCount
	KCGImagePropertyHEICSLoopCount string
	// KCGImagePropertyHEICSUnclampedDelayTime is the unclamped number of seconds to wait before displaying the next image in the sequence.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyHEICSUnclampedDelayTime
	KCGImagePropertyHEICSUnclampedDelayTime string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyHEIFDictionary
	KCGImagePropertyHEIFDictionary string
	// KCGImagePropertyHasAlpha is a Boolean value that indicates whether the image has an alpha channel.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyHasAlpha
	KCGImagePropertyHasAlpha string
	// KCGImagePropertyHeight is the height of the image, in the image’s coordinate space.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyHeight
	KCGImagePropertyHeight string
	// KCGImagePropertyIPTCActionAdvised is the advised action.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCActionAdvised
	KCGImagePropertyIPTCActionAdvised string
	// KCGImagePropertyIPTCByline is the name of the person who created the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCByline
	KCGImagePropertyIPTCByline string
	// KCGImagePropertyIPTCBylineTitle is the title of the person who created the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCBylineTitle
	KCGImagePropertyIPTCBylineTitle string
	// KCGImagePropertyIPTCCaptionAbstract is the description of the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCCaptionAbstract
	KCGImagePropertyIPTCCaptionAbstract string
	// KCGImagePropertyIPTCCategory is the category.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCCategory
	KCGImagePropertyIPTCCategory string
	// KCGImagePropertyIPTCCity is the city where the image was created.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCCity
	KCGImagePropertyIPTCCity string
	// KCGImagePropertyIPTCContact is the contact information for getting details about the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCContact
	KCGImagePropertyIPTCContact string
	// KCGImagePropertyIPTCContactInfoAddress is the address portion of the contact information.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCContactInfoAddress
	KCGImagePropertyIPTCContactInfoAddress string
	// KCGImagePropertyIPTCContactInfoCity is the city portion of the contact information.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCContactInfoCity
	KCGImagePropertyIPTCContactInfoCity string
	// KCGImagePropertyIPTCContactInfoCountry is the country or region portion of the contact information.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCContactInfoCountry
	KCGImagePropertyIPTCContactInfoCountry string
	// KCGImagePropertyIPTCContactInfoEmails is email addresses for the contact.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCContactInfoEmails
	KCGImagePropertyIPTCContactInfoEmails string
	// KCGImagePropertyIPTCContactInfoPhones is phone numbers for the contact.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCContactInfoPhones
	KCGImagePropertyIPTCContactInfoPhones string
	// KCGImagePropertyIPTCContactInfoPostalCode is the postal code portion of the contact.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCContactInfoPostalCode
	KCGImagePropertyIPTCContactInfoPostalCode string
	// KCGImagePropertyIPTCContactInfoStateProvince is the state or province of the contact.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCContactInfoStateProvince
	KCGImagePropertyIPTCContactInfoStateProvince string
	// KCGImagePropertyIPTCContactInfoWebURLs is web addresses for the contact.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCContactInfoWebURLs
	KCGImagePropertyIPTCContactInfoWebURLs string
	// KCGImagePropertyIPTCContentLocationCode is the content location code.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCContentLocationCode
	KCGImagePropertyIPTCContentLocationCode string
	// KCGImagePropertyIPTCContentLocationName is the content location name.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCContentLocationName
	KCGImagePropertyIPTCContentLocationName string
	// KCGImagePropertyIPTCCopyrightNotice is the copyright notice.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCCopyrightNotice
	KCGImagePropertyIPTCCopyrightNotice string
	// KCGImagePropertyIPTCCountryPrimaryLocationCode is the primary country code, a three-letter code defined by ISO 3166-1.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCCountryPrimaryLocationCode
	KCGImagePropertyIPTCCountryPrimaryLocationCode string
	// KCGImagePropertyIPTCCountryPrimaryLocationName is the primary country name.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCCountryPrimaryLocationName
	KCGImagePropertyIPTCCountryPrimaryLocationName string
	// KCGImagePropertyIPTCCreatorContactInfo is the creator’s contact info.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCCreatorContactInfo
	KCGImagePropertyIPTCCreatorContactInfo string
	// KCGImagePropertyIPTCCredit is the name of the service that provided the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCCredit
	KCGImagePropertyIPTCCredit string
	// KCGImagePropertyIPTCDateCreated is the creation date.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCDateCreated
	KCGImagePropertyIPTCDateCreated string
	// KCGImagePropertyIPTCDictionary is a dictionary of key-value pairs for an image that uses International Press Telecommunications Council (IPTC) metadata.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCDictionary
	KCGImagePropertyIPTCDictionary string
	// KCGImagePropertyIPTCDigitalCreationDate is the digital creation date.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCDigitalCreationDate
	KCGImagePropertyIPTCDigitalCreationDate string
	// KCGImagePropertyIPTCDigitalCreationTime is the digital creation time.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCDigitalCreationTime
	KCGImagePropertyIPTCDigitalCreationTime string
	// KCGImagePropertyIPTCEditStatus is the edit status.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCEditStatus
	KCGImagePropertyIPTCEditStatus string
	// KCGImagePropertyIPTCEditorialUpdate is an editorial update.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCEditorialUpdate
	KCGImagePropertyIPTCEditorialUpdate string
	// KCGImagePropertyIPTCExpirationDate is the latest date you can use the image, in the form CCYYMMDD.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExpirationDate
	KCGImagePropertyIPTCExpirationDate string
	// KCGImagePropertyIPTCExpirationTime is the latest time on the expiration date you can use the image, in the form HHMMSS.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExpirationTime
	KCGImagePropertyIPTCExpirationTime string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtAboutCvTerm
	KCGImagePropertyIPTCExtAboutCvTerm string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtAboutCvTermCvId
	KCGImagePropertyIPTCExtAboutCvTermCvId string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtAboutCvTermId
	KCGImagePropertyIPTCExtAboutCvTermId string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtAboutCvTermName
	KCGImagePropertyIPTCExtAboutCvTermName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtAboutCvTermRefinedAbout
	KCGImagePropertyIPTCExtAboutCvTermRefinedAbout string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtAddlModelInfo
	KCGImagePropertyIPTCExtAddlModelInfo string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtArtworkCircaDateCreated
	KCGImagePropertyIPTCExtArtworkCircaDateCreated string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtArtworkContentDescription
	KCGImagePropertyIPTCExtArtworkContentDescription string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtArtworkContributionDescription
	KCGImagePropertyIPTCExtArtworkContributionDescription string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtArtworkCopyrightNotice
	KCGImagePropertyIPTCExtArtworkCopyrightNotice string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtArtworkCopyrightOwnerID
	KCGImagePropertyIPTCExtArtworkCopyrightOwnerID string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtArtworkCopyrightOwnerName
	KCGImagePropertyIPTCExtArtworkCopyrightOwnerName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtArtworkCreator
	KCGImagePropertyIPTCExtArtworkCreator string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtArtworkCreatorID
	KCGImagePropertyIPTCExtArtworkCreatorID string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtArtworkDateCreated
	KCGImagePropertyIPTCExtArtworkDateCreated string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtArtworkLicensorID
	KCGImagePropertyIPTCExtArtworkLicensorID string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtArtworkLicensorName
	KCGImagePropertyIPTCExtArtworkLicensorName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtArtworkOrObject
	KCGImagePropertyIPTCExtArtworkOrObject string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtArtworkPhysicalDescription
	KCGImagePropertyIPTCExtArtworkPhysicalDescription string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtArtworkSource
	KCGImagePropertyIPTCExtArtworkSource string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtArtworkSourceInvURL
	KCGImagePropertyIPTCExtArtworkSourceInvURL string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtArtworkSourceInventoryNo
	KCGImagePropertyIPTCExtArtworkSourceInventoryNo string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtArtworkStylePeriod
	KCGImagePropertyIPTCExtArtworkStylePeriod string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtArtworkTitle
	KCGImagePropertyIPTCExtArtworkTitle string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtAudioBitrate
	KCGImagePropertyIPTCExtAudioBitrate string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtAudioBitrateMode
	KCGImagePropertyIPTCExtAudioBitrateMode string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtAudioChannelCount
	KCGImagePropertyIPTCExtAudioChannelCount string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtCircaDateCreated
	KCGImagePropertyIPTCExtCircaDateCreated string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtContainerFormat
	KCGImagePropertyIPTCExtContainerFormat string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtContainerFormatIdentifier
	KCGImagePropertyIPTCExtContainerFormatIdentifier string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtContainerFormatName
	KCGImagePropertyIPTCExtContainerFormatName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtContributor
	KCGImagePropertyIPTCExtContributor string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtContributorIdentifier
	KCGImagePropertyIPTCExtContributorIdentifier string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtContributorName
	KCGImagePropertyIPTCExtContributorName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtContributorRole
	KCGImagePropertyIPTCExtContributorRole string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtControlledVocabularyTerm
	KCGImagePropertyIPTCExtControlledVocabularyTerm string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtCopyrightYear
	KCGImagePropertyIPTCExtCopyrightYear string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtCreator
	KCGImagePropertyIPTCExtCreator string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtCreatorIdentifier
	KCGImagePropertyIPTCExtCreatorIdentifier string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtCreatorName
	KCGImagePropertyIPTCExtCreatorName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtCreatorRole
	KCGImagePropertyIPTCExtCreatorRole string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtDataOnScreen
	KCGImagePropertyIPTCExtDataOnScreen string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtDataOnScreenRegion
	KCGImagePropertyIPTCExtDataOnScreenRegion string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtDataOnScreenRegionD
	KCGImagePropertyIPTCExtDataOnScreenRegionD string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtDataOnScreenRegionH
	KCGImagePropertyIPTCExtDataOnScreenRegionH string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtDataOnScreenRegionText
	KCGImagePropertyIPTCExtDataOnScreenRegionText string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtDataOnScreenRegionUnit
	KCGImagePropertyIPTCExtDataOnScreenRegionUnit string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtDataOnScreenRegionW
	KCGImagePropertyIPTCExtDataOnScreenRegionW string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtDataOnScreenRegionX
	KCGImagePropertyIPTCExtDataOnScreenRegionX string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtDataOnScreenRegionY
	KCGImagePropertyIPTCExtDataOnScreenRegionY string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtDigitalImageGUID
	KCGImagePropertyIPTCExtDigitalImageGUID string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtDigitalSourceFileType
	KCGImagePropertyIPTCExtDigitalSourceFileType string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtDigitalSourceType
	KCGImagePropertyIPTCExtDigitalSourceType string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtDopesheet
	KCGImagePropertyIPTCExtDopesheet string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtDopesheetLink
	KCGImagePropertyIPTCExtDopesheetLink string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtDopesheetLinkLink
	KCGImagePropertyIPTCExtDopesheetLinkLink string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtDopesheetLinkLinkQualifier
	KCGImagePropertyIPTCExtDopesheetLinkLinkQualifier string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtEmbdEncRightsExpr
	KCGImagePropertyIPTCExtEmbdEncRightsExpr string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtEmbeddedEncodedRightsExpr
	KCGImagePropertyIPTCExtEmbeddedEncodedRightsExpr string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtEmbeddedEncodedRightsExprLangID
	KCGImagePropertyIPTCExtEmbeddedEncodedRightsExprLangID string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtEmbeddedEncodedRightsExprType
	KCGImagePropertyIPTCExtEmbeddedEncodedRightsExprType string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtEpisode
	KCGImagePropertyIPTCExtEpisode string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtEpisodeIdentifier
	KCGImagePropertyIPTCExtEpisodeIdentifier string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtEpisodeName
	KCGImagePropertyIPTCExtEpisodeName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtEpisodeNumber
	KCGImagePropertyIPTCExtEpisodeNumber string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtEvent
	KCGImagePropertyIPTCExtEvent string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtExternalMetadataLink
	KCGImagePropertyIPTCExtExternalMetadataLink string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtFeedIdentifier
	KCGImagePropertyIPTCExtFeedIdentifier string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtGenre
	KCGImagePropertyIPTCExtGenre string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtGenreCvId
	KCGImagePropertyIPTCExtGenreCvId string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtGenreCvTermId
	KCGImagePropertyIPTCExtGenreCvTermId string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtGenreCvTermName
	KCGImagePropertyIPTCExtGenreCvTermName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtGenreCvTermRefinedAbout
	KCGImagePropertyIPTCExtGenreCvTermRefinedAbout string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtHeadline
	KCGImagePropertyIPTCExtHeadline string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtIPTCLastEdited
	KCGImagePropertyIPTCExtIPTCLastEdited string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtLinkedEncRightsExpr
	KCGImagePropertyIPTCExtLinkedEncRightsExpr string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtLinkedEncodedRightsExpr
	KCGImagePropertyIPTCExtLinkedEncodedRightsExpr string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtLinkedEncodedRightsExprLangID
	KCGImagePropertyIPTCExtLinkedEncodedRightsExprLangID string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtLinkedEncodedRightsExprType
	KCGImagePropertyIPTCExtLinkedEncodedRightsExprType string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtLocationCity
	KCGImagePropertyIPTCExtLocationCity string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtLocationCountryCode
	KCGImagePropertyIPTCExtLocationCountryCode string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtLocationCountryName
	KCGImagePropertyIPTCExtLocationCountryName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtLocationCreated
	KCGImagePropertyIPTCExtLocationCreated string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtLocationGPSAltitude
	KCGImagePropertyIPTCExtLocationGPSAltitude string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtLocationGPSLatitude
	KCGImagePropertyIPTCExtLocationGPSLatitude string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtLocationGPSLongitude
	KCGImagePropertyIPTCExtLocationGPSLongitude string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtLocationIdentifier
	KCGImagePropertyIPTCExtLocationIdentifier string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtLocationLocationId
	KCGImagePropertyIPTCExtLocationLocationId string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtLocationLocationName
	KCGImagePropertyIPTCExtLocationLocationName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtLocationProvinceState
	KCGImagePropertyIPTCExtLocationProvinceState string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtLocationShown
	KCGImagePropertyIPTCExtLocationShown string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtLocationSublocation
	KCGImagePropertyIPTCExtLocationSublocation string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtLocationWorldRegion
	KCGImagePropertyIPTCExtLocationWorldRegion string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtMaxAvailHeight
	KCGImagePropertyIPTCExtMaxAvailHeight string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtMaxAvailWidth
	KCGImagePropertyIPTCExtMaxAvailWidth string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtModelAge
	KCGImagePropertyIPTCExtModelAge string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtOrganisationInImageCode
	KCGImagePropertyIPTCExtOrganisationInImageCode string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtOrganisationInImageName
	KCGImagePropertyIPTCExtOrganisationInImageName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtPersonHeard
	KCGImagePropertyIPTCExtPersonHeard string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtPersonHeardIdentifier
	KCGImagePropertyIPTCExtPersonHeardIdentifier string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtPersonHeardName
	KCGImagePropertyIPTCExtPersonHeardName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtPersonInImage
	KCGImagePropertyIPTCExtPersonInImage string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtPersonInImageCharacteristic
	KCGImagePropertyIPTCExtPersonInImageCharacteristic string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtPersonInImageCvTermCvId
	KCGImagePropertyIPTCExtPersonInImageCvTermCvId string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtPersonInImageCvTermId
	KCGImagePropertyIPTCExtPersonInImageCvTermId string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtPersonInImageCvTermName
	KCGImagePropertyIPTCExtPersonInImageCvTermName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtPersonInImageCvTermRefinedAbout
	KCGImagePropertyIPTCExtPersonInImageCvTermRefinedAbout string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtPersonInImageDescription
	KCGImagePropertyIPTCExtPersonInImageDescription string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtPersonInImageId
	KCGImagePropertyIPTCExtPersonInImageId string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtPersonInImageName
	KCGImagePropertyIPTCExtPersonInImageName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtPersonInImageWDetails
	KCGImagePropertyIPTCExtPersonInImageWDetails string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtProductInImage
	KCGImagePropertyIPTCExtProductInImage string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtProductInImageDescription
	KCGImagePropertyIPTCExtProductInImageDescription string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtProductInImageGTIN
	KCGImagePropertyIPTCExtProductInImageGTIN string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtProductInImageName
	KCGImagePropertyIPTCExtProductInImageName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtPublicationEvent
	KCGImagePropertyIPTCExtPublicationEvent string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtPublicationEventDate
	KCGImagePropertyIPTCExtPublicationEventDate string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtPublicationEventIdentifier
	KCGImagePropertyIPTCExtPublicationEventIdentifier string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtPublicationEventName
	KCGImagePropertyIPTCExtPublicationEventName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRating
	KCGImagePropertyIPTCExtRating string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRatingRatingRegion
	KCGImagePropertyIPTCExtRatingRatingRegion string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRatingRegionCity
	KCGImagePropertyIPTCExtRatingRegionCity string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRatingRegionCountryCode
	KCGImagePropertyIPTCExtRatingRegionCountryCode string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRatingRegionCountryName
	KCGImagePropertyIPTCExtRatingRegionCountryName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRatingRegionGPSAltitude
	KCGImagePropertyIPTCExtRatingRegionGPSAltitude string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRatingRegionGPSLatitude
	KCGImagePropertyIPTCExtRatingRegionGPSLatitude string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRatingRegionGPSLongitude
	KCGImagePropertyIPTCExtRatingRegionGPSLongitude string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRatingRegionIdentifier
	KCGImagePropertyIPTCExtRatingRegionIdentifier string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRatingRegionLocationId
	KCGImagePropertyIPTCExtRatingRegionLocationId string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRatingRegionLocationName
	KCGImagePropertyIPTCExtRatingRegionLocationName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRatingRegionProvinceState
	KCGImagePropertyIPTCExtRatingRegionProvinceState string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRatingRegionSublocation
	KCGImagePropertyIPTCExtRatingRegionSublocation string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRatingRegionWorldRegion
	KCGImagePropertyIPTCExtRatingRegionWorldRegion string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRatingScaleMaxValue
	KCGImagePropertyIPTCExtRatingScaleMaxValue string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRatingScaleMinValue
	KCGImagePropertyIPTCExtRatingScaleMinValue string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRatingSourceLink
	KCGImagePropertyIPTCExtRatingSourceLink string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRatingValue
	KCGImagePropertyIPTCExtRatingValue string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRatingValueLogoLink
	KCGImagePropertyIPTCExtRatingValueLogoLink string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRegistryEntryRole
	KCGImagePropertyIPTCExtRegistryEntryRole string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRegistryID
	KCGImagePropertyIPTCExtRegistryID string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRegistryItemID
	KCGImagePropertyIPTCExtRegistryItemID string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtRegistryOrganisationID
	KCGImagePropertyIPTCExtRegistryOrganisationID string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtReleaseReady
	KCGImagePropertyIPTCExtReleaseReady string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtSeason
	KCGImagePropertyIPTCExtSeason string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtSeasonIdentifier
	KCGImagePropertyIPTCExtSeasonIdentifier string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtSeasonName
	KCGImagePropertyIPTCExtSeasonName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtSeasonNumber
	KCGImagePropertyIPTCExtSeasonNumber string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtSeries
	KCGImagePropertyIPTCExtSeries string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtSeriesIdentifier
	KCGImagePropertyIPTCExtSeriesIdentifier string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtSeriesName
	KCGImagePropertyIPTCExtSeriesName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtShownEvent
	KCGImagePropertyIPTCExtShownEvent string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtShownEventIdentifier
	KCGImagePropertyIPTCExtShownEventIdentifier string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtShownEventName
	KCGImagePropertyIPTCExtShownEventName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtStorylineIdentifier
	KCGImagePropertyIPTCExtStorylineIdentifier string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtStreamReady
	KCGImagePropertyIPTCExtStreamReady string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtStylePeriod
	KCGImagePropertyIPTCExtStylePeriod string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtSupplyChainSource
	KCGImagePropertyIPTCExtSupplyChainSource string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtSupplyChainSourceIdentifier
	KCGImagePropertyIPTCExtSupplyChainSourceIdentifier string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtSupplyChainSourceName
	KCGImagePropertyIPTCExtSupplyChainSourceName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtTemporalCoverage
	KCGImagePropertyIPTCExtTemporalCoverage string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtTemporalCoverageFrom
	KCGImagePropertyIPTCExtTemporalCoverageFrom string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtTemporalCoverageTo
	KCGImagePropertyIPTCExtTemporalCoverageTo string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtTranscript
	KCGImagePropertyIPTCExtTranscript string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtTranscriptLink
	KCGImagePropertyIPTCExtTranscriptLink string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtTranscriptLinkLink
	KCGImagePropertyIPTCExtTranscriptLinkLink string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtTranscriptLinkLinkQualifier
	KCGImagePropertyIPTCExtTranscriptLinkLinkQualifier string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtVideoBitrate
	KCGImagePropertyIPTCExtVideoBitrate string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtVideoBitrateMode
	KCGImagePropertyIPTCExtVideoBitrateMode string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtVideoDisplayAspectRatio
	KCGImagePropertyIPTCExtVideoDisplayAspectRatio string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtVideoEncodingProfile
	KCGImagePropertyIPTCExtVideoEncodingProfile string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtVideoShotType
	KCGImagePropertyIPTCExtVideoShotType string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtVideoShotTypeIdentifier
	KCGImagePropertyIPTCExtVideoShotTypeIdentifier string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtVideoShotTypeName
	KCGImagePropertyIPTCExtVideoShotTypeName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtVideoStreamsCount
	KCGImagePropertyIPTCExtVideoStreamsCount string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtVisualColor
	KCGImagePropertyIPTCExtVisualColor string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtWorkflowTag
	KCGImagePropertyIPTCExtWorkflowTag string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtWorkflowTagCvId
	KCGImagePropertyIPTCExtWorkflowTagCvId string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtWorkflowTagCvTermId
	KCGImagePropertyIPTCExtWorkflowTagCvTermId string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtWorkflowTagCvTermName
	KCGImagePropertyIPTCExtWorkflowTagCvTermName string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCExtWorkflowTagCvTermRefinedAbout
	KCGImagePropertyIPTCExtWorkflowTagCvTermRefinedAbout string
	// KCGImagePropertyIPTCFixtureIdentifier is a fixture identifier.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCFixtureIdentifier
	KCGImagePropertyIPTCFixtureIdentifier string
	// KCGImagePropertyIPTCHeadline is a summary of the contents of the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCHeadline
	KCGImagePropertyIPTCHeadline string
	// KCGImagePropertyIPTCImageOrientation is the image orientation (portrait, landscape, or square).
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCImageOrientation
	KCGImagePropertyIPTCImageOrientation string
	// KCGImagePropertyIPTCImageType is the image type.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCImageType
	KCGImagePropertyIPTCImageType string
	// KCGImagePropertyIPTCKeywords is keywords relevant to the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCKeywords
	KCGImagePropertyIPTCKeywords string
	// KCGImagePropertyIPTCLanguageIdentifier is the language identifier, a two-letter code defined by ISO 639:1988.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCLanguageIdentifier
	KCGImagePropertyIPTCLanguageIdentifier string
	// KCGImagePropertyIPTCObjectAttributeReference is the object attribute.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCObjectAttributeReference
	KCGImagePropertyIPTCObjectAttributeReference string
	// KCGImagePropertyIPTCObjectCycle is the editorial cycle (morning, evening, or both) of the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCObjectCycle
	KCGImagePropertyIPTCObjectCycle string
	// KCGImagePropertyIPTCObjectName is the object name.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCObjectName
	KCGImagePropertyIPTCObjectName string
	// KCGImagePropertyIPTCObjectTypeReference is the object type.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCObjectTypeReference
	KCGImagePropertyIPTCObjectTypeReference string
	// KCGImagePropertyIPTCOriginalTransmissionReference is the call letter or number combination associated with the originating point of an image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCOriginalTransmissionReference
	KCGImagePropertyIPTCOriginalTransmissionReference string
	// KCGImagePropertyIPTCOriginatingProgram is the originating application.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCOriginatingProgram
	KCGImagePropertyIPTCOriginatingProgram string
	// KCGImagePropertyIPTCProgramVersion is the application version.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCProgramVersion
	KCGImagePropertyIPTCProgramVersion string
	// KCGImagePropertyIPTCProvinceState is the province or state.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCProvinceState
	KCGImagePropertyIPTCProvinceState string
	// KCGImagePropertyIPTCReferenceDate is the reference date.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCReferenceDate
	KCGImagePropertyIPTCReferenceDate string
	// KCGImagePropertyIPTCReferenceNumber is the reference number.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCReferenceNumber
	KCGImagePropertyIPTCReferenceNumber string
	// KCGImagePropertyIPTCReferenceService is the reference service.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCReferenceService
	KCGImagePropertyIPTCReferenceService string
	// KCGImagePropertyIPTCReleaseDate is the earliest day on which you can use the image, in the form CCYYMMDD.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCReleaseDate
	KCGImagePropertyIPTCReleaseDate string
	// KCGImagePropertyIPTCReleaseTime is the earliest time at which you can use the image, in the form HHMMSS.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCReleaseTime
	KCGImagePropertyIPTCReleaseTime string
	// KCGImagePropertyIPTCRightsUsageTerms is the usage rights for the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCRightsUsageTerms
	KCGImagePropertyIPTCRightsUsageTerms string
	// KCGImagePropertyIPTCScene is the scene codes for the image; a scene code is a six-digit string.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCScene
	KCGImagePropertyIPTCScene string
	// KCGImagePropertyIPTCSource is the original owner of the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCSource
	KCGImagePropertyIPTCSource string
	// KCGImagePropertyIPTCSpecialInstructions is special instructions about the use of the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCSpecialInstructions
	KCGImagePropertyIPTCSpecialInstructions string
	// KCGImagePropertyIPTCStarRating is the star rating.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCStarRating
	KCGImagePropertyIPTCStarRating string
	// KCGImagePropertyIPTCSubLocation is the location within the city where the image was created.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCSubLocation
	KCGImagePropertyIPTCSubLocation string
	// KCGImagePropertyIPTCSubjectReference is the subject.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCSubjectReference
	KCGImagePropertyIPTCSubjectReference string
	// KCGImagePropertyIPTCSupplementalCategory is a supplemental category.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCSupplementalCategory
	KCGImagePropertyIPTCSupplementalCategory string
	// KCGImagePropertyIPTCTimeCreated is the creation time.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCTimeCreated
	KCGImagePropertyIPTCTimeCreated string
	// KCGImagePropertyIPTCUrgency is the urgency level.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCUrgency
	KCGImagePropertyIPTCUrgency string
	// KCGImagePropertyIPTCWriterEditor is the name of the person who wrote or edited the description of the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIPTCWriterEditor
	KCGImagePropertyIPTCWriterEditor string
	// KCGImagePropertyImageCount is the number of images in the file.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyImageCount
	KCGImagePropertyImageCount string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyImageIndex
	KCGImagePropertyImageIndex string
	// KCGImagePropertyImages is an array of dictionaries, each of which contains metadata for one of the images in the file.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyImages
	KCGImagePropertyImages string
	// KCGImagePropertyIsFloat is a Boolean value that indicates whether the image contains floating-point pixel samples.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIsFloat
	KCGImagePropertyIsFloat string
	// KCGImagePropertyIsIndexed is a Boolean value that indicates whether the image contains indexed pixel samples.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyIsIndexed
	KCGImagePropertyIsIndexed string
	// KCGImagePropertyJFIFDensityUnit is the units for the x and y density fields.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyJFIFDensityUnit
	KCGImagePropertyJFIFDensityUnit string
	// KCGImagePropertyJFIFDictionary is a dictionary of key-value pairs for an image that uses JPEG File Interchange Format (JFIF).
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyJFIFDictionary
	KCGImagePropertyJFIFDictionary string
	// KCGImagePropertyJFIFIsProgressive is whether there are versions of the image of increasing quality.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyJFIFIsProgressive
	KCGImagePropertyJFIFIsProgressive string
	// KCGImagePropertyJFIFVersion is the version of JFIF.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyJFIFVersion
	KCGImagePropertyJFIFVersion string
	// KCGImagePropertyJFIFXDensity is the x pixel density.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyJFIFXDensity
	KCGImagePropertyJFIFXDensity string
	// KCGImagePropertyJFIFYDensity is the y pixel density.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyJFIFYDensity
	KCGImagePropertyJFIFYDensity string
	// KCGImagePropertyMakerAppleDictionary is a dictionary of key-value pairs for an image from an Apple camera.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerAppleDictionary
	KCGImagePropertyMakerAppleDictionary string
	// KCGImagePropertyMakerCanonAspectRatioInfo is the image aspect ratio.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerCanonAspectRatioInfo
	KCGImagePropertyMakerCanonAspectRatioInfo string
	// KCGImagePropertyMakerCanonCameraSerialNumber is the camera serial number.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerCanonCameraSerialNumber
	KCGImagePropertyMakerCanonCameraSerialNumber string
	// KCGImagePropertyMakerCanonContinuousDrive is the presence of a continuous drive.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerCanonContinuousDrive
	KCGImagePropertyMakerCanonContinuousDrive string
	// KCGImagePropertyMakerCanonDictionary is a dictionary of key-value pairs for an image from a Canon camera.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerCanonDictionary
	KCGImagePropertyMakerCanonDictionary string
	// KCGImagePropertyMakerCanonFirmware is the firmware version.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerCanonFirmware
	KCGImagePropertyMakerCanonFirmware string
	// KCGImagePropertyMakerCanonFlashExposureComp is the flash exposure compensation setting.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerCanonFlashExposureComp
	KCGImagePropertyMakerCanonFlashExposureComp string
	// KCGImagePropertyMakerCanonImageSerialNumber is the image serial number.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerCanonImageSerialNumber
	KCGImagePropertyMakerCanonImageSerialNumber string
	// KCGImagePropertyMakerCanonLensModel is the lens model.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerCanonLensModel
	KCGImagePropertyMakerCanonLensModel string
	// KCGImagePropertyMakerCanonOwnerName is the name of the camera’s owner.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerCanonOwnerName
	KCGImagePropertyMakerCanonOwnerName string
	// KCGImagePropertyMakerFujiDictionary is a dictionary of key-value pairs for an image from a Fuji camera.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerFujiDictionary
	KCGImagePropertyMakerFujiDictionary string
	// KCGImagePropertyMakerMinoltaDictionary is a dictionary of key-value pairs for an image from a Minolta camera.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerMinoltaDictionary
	KCGImagePropertyMakerMinoltaDictionary string
	// KCGImagePropertyMakerNikonCameraSerialNumber is the camera serial number.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerNikonCameraSerialNumber
	KCGImagePropertyMakerNikonCameraSerialNumber string
	// KCGImagePropertyMakerNikonColorMode is the color mode.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerNikonColorMode
	KCGImagePropertyMakerNikonColorMode string
	// KCGImagePropertyMakerNikonDictionary is a dictionary of key-value pairs for an image from a Nikon camera.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerNikonDictionary
	KCGImagePropertyMakerNikonDictionary string
	// KCGImagePropertyMakerNikonDigitalZoom is the digital zoom setting.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerNikonDigitalZoom
	KCGImagePropertyMakerNikonDigitalZoom string
	// KCGImagePropertyMakerNikonFlashExposureComp is the flash exposure compensation.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerNikonFlashExposureComp
	KCGImagePropertyMakerNikonFlashExposureComp string
	// KCGImagePropertyMakerNikonFlashSetting is the flash setting.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerNikonFlashSetting
	KCGImagePropertyMakerNikonFlashSetting string
	// KCGImagePropertyMakerNikonFocusDistance is the focus distance.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerNikonFocusDistance
	KCGImagePropertyMakerNikonFocusDistance string
	// KCGImagePropertyMakerNikonFocusMode is the focus mode.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerNikonFocusMode
	KCGImagePropertyMakerNikonFocusMode string
	// KCGImagePropertyMakerNikonISOSelection is the ISO selection.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerNikonISOSelection
	KCGImagePropertyMakerNikonISOSelection string
	// KCGImagePropertyMakerNikonISOSetting is the ISO setting.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerNikonISOSetting
	KCGImagePropertyMakerNikonISOSetting string
	// KCGImagePropertyMakerNikonImageAdjustment is the image adjustment setting.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerNikonImageAdjustment
	KCGImagePropertyMakerNikonImageAdjustment string
	// KCGImagePropertyMakerNikonLensAdapter is the lens adapter.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerNikonLensAdapter
	KCGImagePropertyMakerNikonLensAdapter string
	// KCGImagePropertyMakerNikonLensInfo is lens information.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerNikonLensInfo
	KCGImagePropertyMakerNikonLensInfo string
	// KCGImagePropertyMakerNikonLensType is the lens type.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerNikonLensType
	KCGImagePropertyMakerNikonLensType string
	// KCGImagePropertyMakerNikonQuality is the quality setting.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerNikonQuality
	KCGImagePropertyMakerNikonQuality string
	// KCGImagePropertyMakerNikonSharpenMode is the sharpening mode.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerNikonSharpenMode
	KCGImagePropertyMakerNikonSharpenMode string
	// KCGImagePropertyMakerNikonShootingMode is the shooting mode.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerNikonShootingMode
	KCGImagePropertyMakerNikonShootingMode string
	// KCGImagePropertyMakerNikonShutterCount is the number of times the shutter has been actuated.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerNikonShutterCount
	KCGImagePropertyMakerNikonShutterCount string
	// KCGImagePropertyMakerNikonWhiteBalanceMode is the white balance mode.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerNikonWhiteBalanceMode
	KCGImagePropertyMakerNikonWhiteBalanceMode string
	// KCGImagePropertyMakerOlympusDictionary is a dictionary of key-value pairs for an image from a Olympus camera.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerOlympusDictionary
	KCGImagePropertyMakerOlympusDictionary string
	// KCGImagePropertyMakerPentaxDictionary is a dictionary of key-value pairs for an image from a Pentax camera.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyMakerPentaxDictionary
	KCGImagePropertyMakerPentaxDictionary string
	// KCGImagePropertyNamedColorSpace is the name of the image’s color space.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyNamedColorSpace
	KCGImagePropertyNamedColorSpace string
	// KCGImagePropertyOpenEXRAspectRatio is the aspect ratio of the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOpenEXRAspectRatio
	KCGImagePropertyOpenEXRAspectRatio string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOpenEXRCompression
	KCGImagePropertyOpenEXRCompression string
	// KCGImagePropertyOpenEXRDictionary is a dictionary of properties specific to the OpenEXR metadata standard.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOpenEXRDictionary
	KCGImagePropertyOpenEXRDictionary string
	// KCGImagePropertyOrientation is the intended display orientation of the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
	KCGImagePropertyOrientation string
	// KCGImagePropertyPNGAuthor is a string that identifies the author of the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGAuthor
	KCGImagePropertyPNGAuthor string
	// KCGImagePropertyPNGChromaticities is the chromaticities.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGChromaticities
	KCGImagePropertyPNGChromaticities string
	// KCGImagePropertyPNGComment is a string that contains image comments.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGComment
	KCGImagePropertyPNGComment string
	// KCGImagePropertyPNGCompressionFilter is the PNG filter to apply prior to compression.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGCompressionFilter
	KCGImagePropertyPNGCompressionFilter string
	// KCGImagePropertyPNGCopyright is a string that identifies the copyright of the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGCopyright
	KCGImagePropertyPNGCopyright string
	// KCGImagePropertyPNGCreationTime is a string that identifies the date and time the image was created.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGCreationTime
	KCGImagePropertyPNGCreationTime string
	// KCGImagePropertyPNGDescription is a string that describes the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGDescription
	KCGImagePropertyPNGDescription string
	// KCGImagePropertyPNGDictionary is a dictionary of key-value pairs for an image that uses Portable Network Graphics (PNG) format.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGDictionary
	KCGImagePropertyPNGDictionary string
	// KCGImagePropertyPNGDisclaimer is a disclaimer string.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGDisclaimer
	KCGImagePropertyPNGDisclaimer string
	// KCGImagePropertyPNGGamma is the gamma value.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGGamma
	KCGImagePropertyPNGGamma string
	// KCGImagePropertyPNGInterlaceType is the interlace type.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGInterlaceType
	KCGImagePropertyPNGInterlaceType string
	// KCGImagePropertyPNGModificationTime is a string that identifies the last date and time the image was modified.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGModificationTime
	KCGImagePropertyPNGModificationTime string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGPixelsAspectRatio
	KCGImagePropertyPNGPixelsAspectRatio string
	// KCGImagePropertyPNGSoftware is a string that identifies the software used to create the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGSoftware
	KCGImagePropertyPNGSoftware string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGSource
	KCGImagePropertyPNGSource string
	// KCGImagePropertyPNGTitle is a string that holds the image’s title.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGTitle
	KCGImagePropertyPNGTitle string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGTransparency
	KCGImagePropertyPNGTransparency string
	// KCGImagePropertyPNGWarning is a warning string.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGWarning
	KCGImagePropertyPNGWarning string
	// KCGImagePropertyPNGXPixelsPerMeter is the number of x pixels per meter.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGXPixelsPerMeter
	KCGImagePropertyPNGXPixelsPerMeter string
	// KCGImagePropertyPNGYPixelsPerMeter is the number of y pixels per meter.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGYPixelsPerMeter
	KCGImagePropertyPNGYPixelsPerMeter string
	// KCGImagePropertyPNGsRGBIntent is the sRGB intent.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPNGsRGBIntent
	KCGImagePropertyPNGsRGBIntent string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPVREncoder
	KCGImagePropertyPVREncoder string
	// KCGImagePropertyPixelFormat is the format of the image’s individual pixels.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPixelFormat
	KCGImagePropertyPixelFormat string
	// KCGImagePropertyPixelHeight is the number of pixels along the y-axis of the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPixelHeight
	KCGImagePropertyPixelHeight string
	// KCGImagePropertyPixelWidth is the number of pixels along the x-axis of the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPixelWidth
	KCGImagePropertyPixelWidth string
	// KCGImagePropertyPrimaryImage is the index of the primary image in the file.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyPrimaryImage
	KCGImagePropertyPrimaryImage string
	// KCGImagePropertyProfileName is the name of the optional International Color Consortium (ICC) profile embedded in the image, if known.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyProfileName
	KCGImagePropertyProfileName string
	// KCGImagePropertyRawDictionary is a dictionary of key-value pairs for an image that contains minimally processed, or raw, data.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyRawDictionary
	KCGImagePropertyRawDictionary string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTGACompression
	KCGImagePropertyTGACompression string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTGADictionary
	KCGImagePropertyTGADictionary string
	// KCGImagePropertyTIFFArtist is the artist who created the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFArtist
	KCGImagePropertyTIFFArtist string
	// KCGImagePropertyTIFFCompression is the compression scheme used on the image data.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFCompression
	KCGImagePropertyTIFFCompression string
	// KCGImagePropertyTIFFCopyright is copyright information.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFCopyright
	KCGImagePropertyTIFFCopyright string
	// KCGImagePropertyTIFFDateTime is the date and time that the image was created.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFDateTime
	KCGImagePropertyTIFFDateTime string
	// KCGImagePropertyTIFFDictionary is a dictionary of key-value pairs for an image that uses Tagged Image File Format (TIFF).
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFDictionary
	KCGImagePropertyTIFFDictionary string
	// KCGImagePropertyTIFFDocumentName is the document name.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFDocumentName
	KCGImagePropertyTIFFDocumentName string
	// KCGImagePropertyTIFFHostComputer is the computer or operating system used when the image was created.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFHostComputer
	KCGImagePropertyTIFFHostComputer string
	// KCGImagePropertyTIFFImageDescription is the image description.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFImageDescription
	KCGImagePropertyTIFFImageDescription string
	// KCGImagePropertyTIFFMake is the name of the manufacturer of the camera or input device.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFMake
	KCGImagePropertyTIFFMake string
	// KCGImagePropertyTIFFModel is the camera or input device model.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFModel
	KCGImagePropertyTIFFModel string
	// KCGImagePropertyTIFFOrientation is the image orientation.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFOrientation
	KCGImagePropertyTIFFOrientation string
	// KCGImagePropertyTIFFPhotometricInterpretation is the color space of the image data.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFPhotometricInterpretation
	KCGImagePropertyTIFFPhotometricInterpretation string
	// KCGImagePropertyTIFFPrimaryChromaticities is the chromaticities of the primaries of the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFPrimaryChromaticities
	KCGImagePropertyTIFFPrimaryChromaticities string
	// KCGImagePropertyTIFFResolutionUnit is the units of resolution.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFResolutionUnit
	KCGImagePropertyTIFFResolutionUnit string
	// KCGImagePropertyTIFFSoftware is the name and version of the software used for image creation.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFSoftware
	KCGImagePropertyTIFFSoftware string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFTileLength
	KCGImagePropertyTIFFTileLength string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFTileWidth
	KCGImagePropertyTIFFTileWidth string
	// KCGImagePropertyTIFFTransferFunction is the transfer function, in tabular format, used to map pixel components from a nonlinear form into a linear form.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFTransferFunction
	KCGImagePropertyTIFFTransferFunction string
	// KCGImagePropertyTIFFWhitePoint is the white point of the image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFWhitePoint
	KCGImagePropertyTIFFWhitePoint string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFXPosition
	KCGImagePropertyTIFFXPosition string
	// KCGImagePropertyTIFFXResolution is the number of pixels per resolution unit in the image width direction.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFXResolution
	KCGImagePropertyTIFFXResolution string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFYPosition
	KCGImagePropertyTIFFYPosition string
	// KCGImagePropertyTIFFYResolution is the number of pixels per resolution unit in the image height direction.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyTIFFYResolution
	KCGImagePropertyTIFFYResolution string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyThumbnailImages
	KCGImagePropertyThumbnailImages string
	// KCGImagePropertyWebPCanvasPixelHeight is the height of the main image, in pixels.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyWebPCanvasPixelHeight
	KCGImagePropertyWebPCanvasPixelHeight string
	// KCGImagePropertyWebPCanvasPixelWidth is the width of the main image, in pixels.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyWebPCanvasPixelWidth
	KCGImagePropertyWebPCanvasPixelWidth string
	// KCGImagePropertyWebPDelayTime is the number of seconds to wait before displaying the next image in the sequence.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyWebPDelayTime
	KCGImagePropertyWebPDelayTime string
	// KCGImagePropertyWebPDictionary is a dictionary of properties related to a WebP container.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyWebPDictionary
	KCGImagePropertyWebPDictionary string
	// KCGImagePropertyWebPFrameInfoArray is an array of dictionaries that contain timing information for the image sequence.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyWebPFrameInfoArray
	KCGImagePropertyWebPFrameInfoArray string
	// KCGImagePropertyWebPLoopCount is the number of times to play the sequence.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyWebPLoopCount
	KCGImagePropertyWebPLoopCount string
	// KCGImagePropertyWebPUnclampedDelayTime is the unadjusted number of seconds to wait before displaying the next image in the sequence.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyWebPUnclampedDelayTime
	KCGImagePropertyWebPUnclampedDelayTime string
	// KCGImagePropertyWidth is the width of the image, in the image’s coordinate space.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyWidth
	KCGImagePropertyWidth string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageProviderPreferredTileHeight
	KCGImageProviderPreferredTileHeight string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageProviderPreferredTileWidth
	KCGImageProviderPreferredTileWidth string
	// KCGImageSourceCreateThumbnailFromImageAlways is a Boolean value that indicates whether to always create a thumbnail image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageSourceCreateThumbnailFromImageAlways
	KCGImageSourceCreateThumbnailFromImageAlways string
	// KCGImageSourceCreateThumbnailFromImageIfAbsent is a Boolean value that indicates whether to create a thumbnail image automatically if the data source doesn’t contain one.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageSourceCreateThumbnailFromImageIfAbsent
	KCGImageSourceCreateThumbnailFromImageIfAbsent string
	// KCGImageSourceCreateThumbnailWithTransform is a Boolean value that indicates whether to rotate and scale the thumbnail image to match the image’s orientation and aspect ratio.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageSourceCreateThumbnailWithTransform
	KCGImageSourceCreateThumbnailWithTransform string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageSourceDecodeRequest
	KCGImageSourceDecodeRequest string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageSourceDecodeRequestOptions
	KCGImageSourceDecodeRequestOptions string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageSourceDecodeToHDR
	KCGImageSourceDecodeToHDR string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageSourceDecodeToSDR
	KCGImageSourceDecodeToSDR string
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageSourceGenerateImageSpecificLumaScaling
	KCGImageSourceGenerateImageSpecificLumaScaling string
	// KCGImageSourceShouldAllowFloat is a Boolean that indicates whether to use floating-point values in returned images.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageSourceShouldAllowFloat
	KCGImageSourceShouldAllowFloat string
	// KCGImageSourceShouldCache is a Boolean value that indicates whether to cache the decoded image.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageSourceShouldCache
	KCGImageSourceShouldCache string
	// KCGImageSourceShouldCacheImmediately is a Boolean value that indicates whether image decoding and caching happens at image creation time.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageSourceShouldCacheImmediately
	KCGImageSourceShouldCacheImmediately string
	// KCGImageSourceSubsampleFactor is the factor by which to scale down any returned images.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageSourceSubsampleFactor
	KCGImageSourceSubsampleFactor string
	// KCGImageSourceThumbnailMaxPixelSize is the maximum width and height of a thumbnail image, specified in pixels.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageSourceThumbnailMaxPixelSize
	KCGImageSourceThumbnailMaxPixelSize string
	// KCGImageSourceTypeIdentifierHint is the uniform type identifier that represents your best guess for the image’s type.
	//
	// See: https://developer.apple.com/documentation/ImageIO/kCGImageSourceTypeIdentifierHint
	KCGImageSourceTypeIdentifierHint string
	// See: https://developer.apple.com/documentation/ImageIO/kIIOCameraExtrinsics_CoordinateSystemID
	KIIOCameraExtrinsics_CoordinateSystemID string
	// See: https://developer.apple.com/documentation/ImageIO/kIIOCameraExtrinsics_Position
	KIIOCameraExtrinsics_Position string
	// See: https://developer.apple.com/documentation/ImageIO/kIIOCameraExtrinsics_Rotation
	KIIOCameraExtrinsics_Rotation string
	// See: https://developer.apple.com/documentation/ImageIO/kIIOCameraModelType_GenericPinhole
	KIIOCameraModelType_GenericPinhole string
	// See: https://developer.apple.com/documentation/ImageIO/kIIOCameraModelType_SimplifiedPinhole
	KIIOCameraModelType_SimplifiedPinhole string
	// See: https://developer.apple.com/documentation/ImageIO/kIIOCameraModel_Intrinsics
	KIIOCameraModel_Intrinsics string
	// See: https://developer.apple.com/documentation/ImageIO/kIIOCameraModel_ModelType
	KIIOCameraModel_ModelType string
	// See: https://developer.apple.com/documentation/ImageIO/kIIOMetadata_CameraExtrinsicsKey
	KIIOMetadata_CameraExtrinsicsKey string
	// See: https://developer.apple.com/documentation/ImageIO/kIIOMetadata_CameraModelKey
	KIIOMetadata_CameraModelKey string
	// See: https://developer.apple.com/documentation/ImageIO/kIIOMonoscopicImageLocation_Center
	KIIOMonoscopicImageLocation_Center string
	// See: https://developer.apple.com/documentation/ImageIO/kIIOMonoscopicImageLocation_Left
	KIIOMonoscopicImageLocation_Left string
	// See: https://developer.apple.com/documentation/ImageIO/kIIOMonoscopicImageLocation_Right
	KIIOMonoscopicImageLocation_Right string
	// See: https://developer.apple.com/documentation/ImageIO/kIIOMonoscopicImageLocation_Unspecified
	KIIOMonoscopicImageLocation_Unspecified string
	// See: https://developer.apple.com/documentation/ImageIO/kIIOStereoAggressors_Severity
	KIIOStereoAggressors_Severity string
	// See: https://developer.apple.com/documentation/ImageIO/kIIOStereoAggressors_SubTypeURI
	KIIOStereoAggressors_SubTypeURI string
	// See: https://developer.apple.com/documentation/ImageIO/kIIOStereoAggressors_Type
	KIIOStereoAggressors_Type string
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFErrorDomainCGImageMetadata"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFErrorDomainCGImageMetadata = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGComputeHDRStats"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGComputeHDRStats = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageAnimationDelayTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageAnimationDelayTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageAnimationLoopCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageAnimationLoopCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageAnimationStartIndex"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageAnimationStartIndex = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageAuxiliaryDataInfoColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageAuxiliaryDataInfoColorSpace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageAuxiliaryDataInfoData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageAuxiliaryDataInfoData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageAuxiliaryDataInfoDataDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageAuxiliaryDataInfoDataDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageAuxiliaryDataInfoMetadata"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageAuxiliaryDataInfoMetadata = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageAuxiliaryDataTypeDepth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageAuxiliaryDataTypeDepth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageAuxiliaryDataTypeDisparity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageAuxiliaryDataTypeDisparity = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageAuxiliaryDataTypeHDRGainMap"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageAuxiliaryDataTypeHDRGainMap = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageAuxiliaryDataTypeISOGainMap"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageAuxiliaryDataTypeISOGainMap = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageAuxiliaryDataTypePortraitEffectsMatte"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageAuxiliaryDataTypePortraitEffectsMatte = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageAuxiliaryDataTypeSemanticSegmentationGlassesMatte"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageAuxiliaryDataTypeSemanticSegmentationGlassesMatte = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageAuxiliaryDataTypeSemanticSegmentationHairMatte"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageAuxiliaryDataTypeSemanticSegmentationHairMatte = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageAuxiliaryDataTypeSemanticSegmentationSkinMatte"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageAuxiliaryDataTypeSemanticSegmentationSkinMatte = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageAuxiliaryDataTypeSemanticSegmentationSkyMatte"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageAuxiliaryDataTypeSemanticSegmentationSkyMatte = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageAuxiliaryDataTypeSemanticSegmentationTeethMatte"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageAuxiliaryDataTypeSemanticSegmentationTeethMatte = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationBackgroundColor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationBackgroundColor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationDateTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationDateTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationEmbedThumbnail"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationEmbedThumbnail = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationEncodeAlternateColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationEncodeAlternateColorSpace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationEncodeBaseColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationEncodeBaseColorSpace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationEncodeBaseIsSDR"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationEncodeBaseIsSDR = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationEncodeBasePixelFormatRequest"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationEncodeBasePixelFormatRequest = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationEncodeGainMapPixelFormatRequest"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationEncodeGainMapPixelFormatRequest = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationEncodeGainMapSubsampleFactor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationEncodeGainMapSubsampleFactor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationEncodeGenerateGainMapWithBaseImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationEncodeGenerateGainMapWithBaseImage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationEncodeIsBaseImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationEncodeIsBaseImage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationEncodeRequest"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationEncodeRequest = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationEncodeRequestOptions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationEncodeRequestOptions = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationEncodeToISOGainmap"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationEncodeToISOGainmap = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationEncodeToISOHDR"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationEncodeToISOHDR = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationEncodeToSDR"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationEncodeToSDR = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationEncodeTonemapMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationEncodeTonemapMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationImageMaxPixelSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationImageMaxPixelSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationLossyCompressionQuality"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationLossyCompressionQuality = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationMergeMetadata"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationMergeMetadata = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationMetadata"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationMetadata = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationOptimizeColorForSharing"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationOptimizeColorForSharing = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationOrientation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationOrientation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageDestinationPreserveGainMap"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageDestinationPreserveGainMap = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataEnumerateRecursively"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataEnumerateRecursively = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataNamespaceDublinCore"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataNamespaceDublinCore = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataNamespaceExif"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataNamespaceExif = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataNamespaceExifAux"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataNamespaceExifAux = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataNamespaceExifEX"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataNamespaceExifEX = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataNamespaceIPTCCore"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataNamespaceIPTCCore = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataNamespaceIPTCExtension"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataNamespaceIPTCExtension = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataNamespacePhotoshop"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataNamespacePhotoshop = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataNamespaceTIFF"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataNamespaceTIFF = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataNamespaceXMPBasic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataNamespaceXMPBasic = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataNamespaceXMPRights"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataNamespaceXMPRights = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataPrefixDublinCore"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataPrefixDublinCore = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataPrefixExif"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataPrefixExif = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataPrefixExifAux"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataPrefixExifAux = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataPrefixExifEX"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataPrefixExifEX = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataPrefixIPTCCore"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataPrefixIPTCCore = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataPrefixIPTCExtension"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataPrefixIPTCExtension = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataPrefixPhotoshop"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataPrefixPhotoshop = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataPrefixTIFF"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataPrefixTIFF = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataPrefixXMPBasic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataPrefixXMPBasic = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataPrefixXMPRights"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataPrefixXMPRights = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataShouldExcludeGPS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataShouldExcludeGPS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageMetadataShouldExcludeXMP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageMetadataShouldExcludeXMP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageProperty8BIMDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageProperty8BIMDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageProperty8BIMLayerNames"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageProperty8BIMLayerNames = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageProperty8BIMVersion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageProperty8BIMVersion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyAPNGCanvasPixelHeight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyAPNGCanvasPixelHeight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyAPNGCanvasPixelWidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyAPNGCanvasPixelWidth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyAPNGDelayTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyAPNGDelayTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyAPNGFrameInfoArray"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyAPNGFrameInfoArray = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyAPNGLoopCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyAPNGLoopCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyAPNGUnclampedDelayTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyAPNGUnclampedDelayTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyASTCBlockSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyASTCBlockSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyASTCBlockSize4x4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyASTCBlockSize4x4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyASTCBlockSize8x8"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyASTCBlockSize8x8 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyASTCEncoder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyASTCEncoder = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyAVISDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyAVISDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyAuxiliaryData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyAuxiliaryData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyAuxiliaryDataType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyAuxiliaryDataType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyBCEncoder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyBCEncoder = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyBCFormat"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyBCFormat = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyBytesPerRow"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyBytesPerRow = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFCameraSerialNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFCameraSerialNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFContinuousDrive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFContinuousDrive = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFFirmware"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFFirmware = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFFlashExposureComp"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFFlashExposureComp = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFFocusMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFFocusMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFImageFileName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFImageFileName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFImageName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFImageName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFImageSerialNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFImageSerialNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFLensMaxMM"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFLensMaxMM = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFLensMinMM"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFLensMinMM = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFLensModel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFLensModel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFMeasuredEV"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFMeasuredEV = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFMeteringMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFMeteringMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFOwnerName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFOwnerName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFRecordID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFRecordID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFReleaseMethod"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFReleaseMethod = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFReleaseTiming"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFReleaseTiming = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFSelfTimingTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFSelfTimingTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFShootingMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFShootingMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyCIFFWhiteBalanceIndex"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyCIFFWhiteBalanceIndex = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyColorModel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyColorModel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyColorModelCMYK"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyColorModelCMYK = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyColorModelGray"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyColorModelGray = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyColorModelLab"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyColorModelLab = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyColorModelRGB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyColorModelRGB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGActiveArea"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGActiveArea = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGAnalogBalance"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGAnalogBalance = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGAntiAliasStrength"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGAntiAliasStrength = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGAsShotICCProfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGAsShotICCProfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGAsShotNeutral"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGAsShotNeutral = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGAsShotPreProfileMatrix"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGAsShotPreProfileMatrix = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGAsShotProfileName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGAsShotProfileName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGAsShotWhiteXY"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGAsShotWhiteXY = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGBackwardVersion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGBackwardVersion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGBaselineExposure"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGBaselineExposure = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGBaselineExposureOffset"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGBaselineExposureOffset = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGBaselineNoise"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGBaselineNoise = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGBaselineSharpness"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGBaselineSharpness = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGBayerGreenSplit"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGBayerGreenSplit = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGBestQualityScale"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGBestQualityScale = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGBlackLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGBlackLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGBlackLevelDeltaH"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGBlackLevelDeltaH = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGBlackLevelDeltaV"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGBlackLevelDeltaV = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGBlackLevelRepeatDim"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGBlackLevelRepeatDim = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGCFALayout"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGCFALayout = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGCFAPlaneColor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGCFAPlaneColor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGCalibrationIlluminant1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGCalibrationIlluminant1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGCalibrationIlluminant2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGCalibrationIlluminant2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGCameraCalibration1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGCameraCalibration1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGCameraCalibration2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGCameraCalibration2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGCameraCalibrationSignature"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGCameraCalibrationSignature = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGCameraSerialNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGCameraSerialNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGChromaBlurRadius"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGChromaBlurRadius = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGColorMatrix1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGColorMatrix1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGColorMatrix2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGColorMatrix2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGColorimetricReference"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGColorimetricReference = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGCurrentICCProfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGCurrentICCProfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGCurrentPreProfileMatrix"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGCurrentPreProfileMatrix = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGDefaultBlackRender"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGDefaultBlackRender = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGDefaultCropOrigin"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGDefaultCropOrigin = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGDefaultCropSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGDefaultCropSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGDefaultScale"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGDefaultScale = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGDefaultUserCrop"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGDefaultUserCrop = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGExtraCameraProfiles"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGExtraCameraProfiles = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGFixVignetteRadial"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGFixVignetteRadial = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGForwardMatrix1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGForwardMatrix1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGForwardMatrix2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGForwardMatrix2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGLensInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGLensInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGLinearResponseLimit"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGLinearResponseLimit = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGLinearizationTable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGLinearizationTable = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGLocalizedCameraModel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGLocalizedCameraModel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGMakerNoteSafety"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGMakerNoteSafety = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGMaskedAreas"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGMaskedAreas = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGNewRawImageDigest"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGNewRawImageDigest = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGNoiseProfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGNoiseProfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGNoiseReductionApplied"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGNoiseReductionApplied = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGOpcodeList1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGOpcodeList1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGOpcodeList2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGOpcodeList2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGOpcodeList3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGOpcodeList3 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGOriginalBestQualityFinalSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGOriginalBestQualityFinalSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGOriginalDefaultCropSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGOriginalDefaultCropSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGOriginalDefaultFinalSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGOriginalDefaultFinalSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGOriginalRawFileData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGOriginalRawFileData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGOriginalRawFileDigest"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGOriginalRawFileDigest = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGOriginalRawFileName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGOriginalRawFileName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGPreviewApplicationName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGPreviewApplicationName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGPreviewApplicationVersion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGPreviewApplicationVersion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGPreviewColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGPreviewColorSpace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGPreviewDateTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGPreviewDateTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGPreviewSettingsDigest"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGPreviewSettingsDigest = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGPreviewSettingsName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGPreviewSettingsName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGPrivateData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGPrivateData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGProfileCalibrationSignature"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGProfileCalibrationSignature = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGProfileCopyright"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGProfileCopyright = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGProfileEmbedPolicy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGProfileEmbedPolicy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGProfileHueSatMapData1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGProfileHueSatMapData1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGProfileHueSatMapData2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGProfileHueSatMapData2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGProfileHueSatMapDims"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGProfileHueSatMapDims = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGProfileHueSatMapEncoding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGProfileHueSatMapEncoding = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGProfileLookTableData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGProfileLookTableData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGProfileLookTableDims"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGProfileLookTableDims = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGProfileLookTableEncoding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGProfileLookTableEncoding = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGProfileName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGProfileName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGProfileToneCurve"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGProfileToneCurve = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGRawDataUniqueID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGRawDataUniqueID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGRawImageDigest"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGRawImageDigest = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGRawToPreviewGain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGRawToPreviewGain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGReductionMatrix1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGReductionMatrix1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGReductionMatrix2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGReductionMatrix2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGRowInterleaveFactor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGRowInterleaveFactor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGShadowScale"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGShadowScale = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGSubTileBlockSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGSubTileBlockSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGUniqueCameraModel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGUniqueCameraModel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGVersion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGVersion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGWarpFisheye"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGWarpFisheye = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGWarpRectilinear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGWarpRectilinear = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDNGWhiteLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDNGWhiteLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDPIHeight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDPIHeight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDPIWidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDPIWidth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyDepth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyDepth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyEncoder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyEncoder = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifApertureValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifApertureValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifAuxDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifAuxDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifAuxFirmware"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifAuxFirmware = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifAuxFlashCompensation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifAuxFlashCompensation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifAuxImageNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifAuxImageNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifAuxLensID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifAuxLensID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifAuxLensInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifAuxLensInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifAuxLensModel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifAuxLensModel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifAuxLensSerialNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifAuxLensSerialNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifAuxOwnerName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifAuxOwnerName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifAuxSerialNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifAuxSerialNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifBodySerialNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifBodySerialNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifBrightnessValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifBrightnessValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifCFAPattern"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifCFAPattern = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifCameraOwnerName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifCameraOwnerName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifColorSpace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifComponentsConfiguration"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifComponentsConfiguration = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifCompositeImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifCompositeImage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifCompressedBitsPerPixel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifCompressedBitsPerPixel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifContrast"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifContrast = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifCustomRendered"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifCustomRendered = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifDateTimeDigitized"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifDateTimeDigitized = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifDateTimeOriginal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifDateTimeOriginal = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifDeviceSettingDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifDeviceSettingDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifDigitalZoomRatio"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifDigitalZoomRatio = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifExposureBiasValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifExposureBiasValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifExposureIndex"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifExposureIndex = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifExposureMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifExposureMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifExposureProgram"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifExposureProgram = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifExposureTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifExposureTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifFNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifFNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifFileSource"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifFileSource = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifFlash"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifFlash = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifFlashEnergy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifFlashEnergy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifFlashPixVersion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifFlashPixVersion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifFocalLenIn35mmFilm"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifFocalLenIn35mmFilm = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifFocalLength"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifFocalLength = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifFocalPlaneResolutionUnit"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifFocalPlaneResolutionUnit = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifFocalPlaneXResolution"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifFocalPlaneXResolution = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifFocalPlaneYResolution"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifFocalPlaneYResolution = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifGainControl"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifGainControl = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifGamma"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifGamma = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifISOSpeed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifISOSpeed = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifISOSpeedLatitudeyyy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifISOSpeedLatitudeyyy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifISOSpeedLatitudezzz"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifISOSpeedLatitudezzz = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifISOSpeedRatings"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifISOSpeedRatings = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifImageUniqueID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifImageUniqueID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifLensMake"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifLensMake = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifLensModel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifLensModel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifLensSerialNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifLensSerialNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifLensSpecification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifLensSpecification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifLightSource"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifLightSource = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifMakerNote"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifMakerNote = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifMaxApertureValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifMaxApertureValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifMeteringMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifMeteringMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifOECF"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifOECF = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifOffsetTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifOffsetTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifOffsetTimeDigitized"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifOffsetTimeDigitized = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifOffsetTimeOriginal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifOffsetTimeOriginal = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifPixelXDimension"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifPixelXDimension = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifPixelYDimension"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifPixelYDimension = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifRecommendedExposureIndex"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifRecommendedExposureIndex = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifRelatedSoundFile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifRelatedSoundFile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifSaturation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifSaturation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifSceneCaptureType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifSceneCaptureType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifSceneType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifSceneType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifSensingMethod"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifSensingMethod = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifSensitivityType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifSensitivityType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifSharpness"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifSharpness = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifShutterSpeedValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifShutterSpeedValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifSourceExposureTimesOfCompositeImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifSourceExposureTimesOfCompositeImage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifSourceImageNumberOfCompositeImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifSourceImageNumberOfCompositeImage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifSpatialFrequencyResponse"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifSpatialFrequencyResponse = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifSpectralSensitivity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifSpectralSensitivity = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifStandardOutputSensitivity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifStandardOutputSensitivity = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifSubjectArea"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifSubjectArea = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifSubjectDistRange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifSubjectDistRange = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifSubjectDistance"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifSubjectDistance = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifSubjectLocation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifSubjectLocation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifSubsecTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifSubsecTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifSubsecTimeDigitized"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifSubsecTimeDigitized = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifSubsecTimeOriginal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifSubsecTimeOriginal = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifUserComment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifUserComment = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifVersion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifVersion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyExifWhiteBalance"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyExifWhiteBalance = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyFileContentsDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyFileContentsDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyFileSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyFileSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGIFCanvasPixelHeight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGIFCanvasPixelHeight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGIFCanvasPixelWidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGIFCanvasPixelWidth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGIFDelayTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGIFDelayTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGIFDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGIFDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGIFFrameInfoArray"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGIFFrameInfoArray = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGIFHasGlobalColorMap"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGIFHasGlobalColorMap = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGIFImageColorMap"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGIFImageColorMap = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGIFLoopCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGIFLoopCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGIFUnclampedDelayTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGIFUnclampedDelayTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSAltitude"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSAltitude = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSAltitudeRef"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSAltitudeRef = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSAreaInformation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSAreaInformation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSDOP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSDOP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSDateStamp"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSDateStamp = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSDestBearing"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSDestBearing = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSDestBearingRef"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSDestBearingRef = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSDestDistance"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSDestDistance = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSDestDistanceRef"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSDestDistanceRef = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSDestLatitude"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSDestLatitude = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSDestLatitudeRef"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSDestLatitudeRef = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSDestLongitude"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSDestLongitude = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSDestLongitudeRef"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSDestLongitudeRef = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSDifferental"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSDifferental = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSHPositioningError"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSHPositioningError = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSImgDirection"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSImgDirection = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSImgDirectionRef"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSImgDirectionRef = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSLatitude"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSLatitude = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSLatitudeRef"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSLatitudeRef = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSLongitude"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSLongitude = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSLongitudeRef"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSLongitudeRef = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSMapDatum"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSMapDatum = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSMeasureMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSMeasureMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSProcessingMethod"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSProcessingMethod = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSSatellites"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSSatellites = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSSpeed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSSpeed = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSSpeedRef"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSSpeedRef = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSStatus"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSStatus = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSTimeStamp"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSTimeStamp = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSTrack"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSTrack = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSTrackRef"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSTrackRef = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGPSVersion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGPSVersion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGroupImageBaseline"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGroupImageBaseline = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGroupImageDisparityAdjustment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGroupImageDisparityAdjustment = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGroupImageIndexLeft"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGroupImageIndexLeft = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGroupImageIndexMonoscopic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGroupImageIndexMonoscopic = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGroupImageIndexRight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGroupImageIndexRight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGroupImageIsAlternateImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGroupImageIsAlternateImage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGroupImageIsLeftImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGroupImageIsLeftImage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGroupImageIsMonoscopicImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGroupImageIsMonoscopicImage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGroupImageIsRightImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGroupImageIsRightImage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGroupImageStereoAggressors"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGroupImageStereoAggressors = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGroupImagesAlternate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGroupImagesAlternate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGroupIndex"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGroupIndex = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGroupMonoscopicImageLocation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGroupMonoscopicImageLocation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGroupType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGroupType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGroupTypeAlternate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGroupTypeAlternate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGroupTypeStereoPair"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGroupTypeStereoPair = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyGroups"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyGroups = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyHEICSCanvasPixelHeight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyHEICSCanvasPixelHeight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyHEICSCanvasPixelWidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyHEICSCanvasPixelWidth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyHEICSDelayTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyHEICSDelayTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyHEICSDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyHEICSDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyHEICSFrameInfoArray"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyHEICSFrameInfoArray = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyHEICSLoopCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyHEICSLoopCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyHEICSUnclampedDelayTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyHEICSUnclampedDelayTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyHEIFDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyHEIFDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyHasAlpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyHasAlpha = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyHeight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyHeight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCActionAdvised"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCActionAdvised = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCByline"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCByline = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCBylineTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCBylineTitle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCCaptionAbstract"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCCaptionAbstract = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCCategory"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCCategory = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCCity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCCity = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCContact"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCContact = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCContactInfoAddress"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCContactInfoAddress = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCContactInfoCity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCContactInfoCity = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCContactInfoCountry"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCContactInfoCountry = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCContactInfoEmails"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCContactInfoEmails = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCContactInfoPhones"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCContactInfoPhones = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCContactInfoPostalCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCContactInfoPostalCode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCContactInfoStateProvince"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCContactInfoStateProvince = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCContactInfoWebURLs"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCContactInfoWebURLs = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCContentLocationCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCContentLocationCode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCContentLocationName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCContentLocationName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCCopyrightNotice"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCCopyrightNotice = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCCountryPrimaryLocationCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCCountryPrimaryLocationCode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCCountryPrimaryLocationName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCCountryPrimaryLocationName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCCreatorContactInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCCreatorContactInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCCredit"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCCredit = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCDateCreated"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCDateCreated = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCDigitalCreationDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCDigitalCreationDate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCDigitalCreationTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCDigitalCreationTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCEditStatus"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCEditStatus = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCEditorialUpdate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCEditorialUpdate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExpirationDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExpirationDate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExpirationTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExpirationTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtAboutCvTerm"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtAboutCvTerm = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtAboutCvTermCvId"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtAboutCvTermCvId = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtAboutCvTermId"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtAboutCvTermId = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtAboutCvTermName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtAboutCvTermName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtAboutCvTermRefinedAbout"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtAboutCvTermRefinedAbout = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtAddlModelInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtAddlModelInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtArtworkCircaDateCreated"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtArtworkCircaDateCreated = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtArtworkContentDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtArtworkContentDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtArtworkContributionDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtArtworkContributionDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtArtworkCopyrightNotice"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtArtworkCopyrightNotice = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtArtworkCopyrightOwnerID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtArtworkCopyrightOwnerID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtArtworkCopyrightOwnerName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtArtworkCopyrightOwnerName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtArtworkCreator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtArtworkCreator = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtArtworkCreatorID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtArtworkCreatorID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtArtworkDateCreated"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtArtworkDateCreated = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtArtworkLicensorID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtArtworkLicensorID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtArtworkLicensorName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtArtworkLicensorName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtArtworkOrObject"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtArtworkOrObject = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtArtworkPhysicalDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtArtworkPhysicalDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtArtworkSource"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtArtworkSource = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtArtworkSourceInvURL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtArtworkSourceInvURL = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtArtworkSourceInventoryNo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtArtworkSourceInventoryNo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtArtworkStylePeriod"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtArtworkStylePeriod = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtArtworkTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtArtworkTitle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtAudioBitrate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtAudioBitrate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtAudioBitrateMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtAudioBitrateMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtAudioChannelCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtAudioChannelCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtCircaDateCreated"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtCircaDateCreated = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtContainerFormat"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtContainerFormat = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtContainerFormatIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtContainerFormatIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtContainerFormatName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtContainerFormatName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtContributor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtContributor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtContributorIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtContributorIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtContributorName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtContributorName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtContributorRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtContributorRole = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtControlledVocabularyTerm"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtControlledVocabularyTerm = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtCopyrightYear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtCopyrightYear = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtCreator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtCreator = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtCreatorIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtCreatorIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtCreatorName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtCreatorName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtCreatorRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtCreatorRole = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtDataOnScreen"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtDataOnScreen = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtDataOnScreenRegion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtDataOnScreenRegion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtDataOnScreenRegionD"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtDataOnScreenRegionD = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtDataOnScreenRegionH"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtDataOnScreenRegionH = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtDataOnScreenRegionText"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtDataOnScreenRegionText = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtDataOnScreenRegionUnit"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtDataOnScreenRegionUnit = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtDataOnScreenRegionW"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtDataOnScreenRegionW = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtDataOnScreenRegionX"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtDataOnScreenRegionX = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtDataOnScreenRegionY"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtDataOnScreenRegionY = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtDigitalImageGUID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtDigitalImageGUID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtDigitalSourceFileType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtDigitalSourceFileType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtDigitalSourceType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtDigitalSourceType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtDopesheet"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtDopesheet = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtDopesheetLink"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtDopesheetLink = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtDopesheetLinkLink"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtDopesheetLinkLink = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtDopesheetLinkLinkQualifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtDopesheetLinkLinkQualifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtEmbdEncRightsExpr"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtEmbdEncRightsExpr = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtEmbeddedEncodedRightsExpr"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtEmbeddedEncodedRightsExpr = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtEmbeddedEncodedRightsExprLangID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtEmbeddedEncodedRightsExprLangID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtEmbeddedEncodedRightsExprType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtEmbeddedEncodedRightsExprType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtEpisode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtEpisode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtEpisodeIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtEpisodeIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtEpisodeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtEpisodeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtEpisodeNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtEpisodeNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtEvent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtEvent = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtExternalMetadataLink"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtExternalMetadataLink = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtFeedIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtFeedIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtGenre"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtGenre = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtGenreCvId"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtGenreCvId = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtGenreCvTermId"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtGenreCvTermId = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtGenreCvTermName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtGenreCvTermName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtGenreCvTermRefinedAbout"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtGenreCvTermRefinedAbout = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtHeadline"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtHeadline = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtIPTCLastEdited"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtIPTCLastEdited = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtLinkedEncRightsExpr"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtLinkedEncRightsExpr = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtLinkedEncodedRightsExpr"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtLinkedEncodedRightsExpr = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtLinkedEncodedRightsExprLangID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtLinkedEncodedRightsExprLangID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtLinkedEncodedRightsExprType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtLinkedEncodedRightsExprType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtLocationCity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtLocationCity = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtLocationCountryCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtLocationCountryCode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtLocationCountryName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtLocationCountryName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtLocationCreated"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtLocationCreated = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtLocationGPSAltitude"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtLocationGPSAltitude = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtLocationGPSLatitude"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtLocationGPSLatitude = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtLocationGPSLongitude"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtLocationGPSLongitude = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtLocationIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtLocationIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtLocationLocationId"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtLocationLocationId = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtLocationLocationName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtLocationLocationName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtLocationProvinceState"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtLocationProvinceState = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtLocationShown"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtLocationShown = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtLocationSublocation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtLocationSublocation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtLocationWorldRegion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtLocationWorldRegion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtMaxAvailHeight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtMaxAvailHeight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtMaxAvailWidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtMaxAvailWidth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtModelAge"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtModelAge = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtOrganisationInImageCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtOrganisationInImageCode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtOrganisationInImageName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtOrganisationInImageName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtPersonHeard"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtPersonHeard = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtPersonHeardIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtPersonHeardIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtPersonHeardName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtPersonHeardName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtPersonInImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtPersonInImage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtPersonInImageCharacteristic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtPersonInImageCharacteristic = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtPersonInImageCvTermCvId"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtPersonInImageCvTermCvId = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtPersonInImageCvTermId"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtPersonInImageCvTermId = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtPersonInImageCvTermName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtPersonInImageCvTermName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtPersonInImageCvTermRefinedAbout"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtPersonInImageCvTermRefinedAbout = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtPersonInImageDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtPersonInImageDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtPersonInImageId"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtPersonInImageId = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtPersonInImageName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtPersonInImageName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtPersonInImageWDetails"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtPersonInImageWDetails = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtProductInImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtProductInImage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtProductInImageDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtProductInImageDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtProductInImageGTIN"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtProductInImageGTIN = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtProductInImageName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtProductInImageName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtPublicationEvent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtPublicationEvent = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtPublicationEventDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtPublicationEventDate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtPublicationEventIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtPublicationEventIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtPublicationEventName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtPublicationEventName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRating"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRating = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRatingRatingRegion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRatingRatingRegion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRatingRegionCity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRatingRegionCity = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRatingRegionCountryCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRatingRegionCountryCode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRatingRegionCountryName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRatingRegionCountryName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRatingRegionGPSAltitude"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRatingRegionGPSAltitude = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRatingRegionGPSLatitude"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRatingRegionGPSLatitude = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRatingRegionGPSLongitude"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRatingRegionGPSLongitude = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRatingRegionIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRatingRegionIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRatingRegionLocationId"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRatingRegionLocationId = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRatingRegionLocationName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRatingRegionLocationName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRatingRegionProvinceState"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRatingRegionProvinceState = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRatingRegionSublocation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRatingRegionSublocation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRatingRegionWorldRegion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRatingRegionWorldRegion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRatingScaleMaxValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRatingScaleMaxValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRatingScaleMinValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRatingScaleMinValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRatingSourceLink"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRatingSourceLink = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRatingValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRatingValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRatingValueLogoLink"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRatingValueLogoLink = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRegistryEntryRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRegistryEntryRole = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRegistryID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRegistryID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRegistryItemID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRegistryItemID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtRegistryOrganisationID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtRegistryOrganisationID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtReleaseReady"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtReleaseReady = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtSeason"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtSeason = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtSeasonIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtSeasonIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtSeasonName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtSeasonName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtSeasonNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtSeasonNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtSeries"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtSeries = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtSeriesIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtSeriesIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtSeriesName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtSeriesName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtShownEvent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtShownEvent = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtShownEventIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtShownEventIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtShownEventName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtShownEventName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtStorylineIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtStorylineIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtStreamReady"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtStreamReady = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtStylePeriod"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtStylePeriod = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtSupplyChainSource"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtSupplyChainSource = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtSupplyChainSourceIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtSupplyChainSourceIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtSupplyChainSourceName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtSupplyChainSourceName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtTemporalCoverage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtTemporalCoverage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtTemporalCoverageFrom"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtTemporalCoverageFrom = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtTemporalCoverageTo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtTemporalCoverageTo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtTranscript"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtTranscript = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtTranscriptLink"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtTranscriptLink = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtTranscriptLinkLink"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtTranscriptLinkLink = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtTranscriptLinkLinkQualifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtTranscriptLinkLinkQualifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtVideoBitrate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtVideoBitrate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtVideoBitrateMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtVideoBitrateMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtVideoDisplayAspectRatio"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtVideoDisplayAspectRatio = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtVideoEncodingProfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtVideoEncodingProfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtVideoShotType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtVideoShotType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtVideoShotTypeIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtVideoShotTypeIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtVideoShotTypeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtVideoShotTypeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtVideoStreamsCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtVideoStreamsCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtVisualColor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtVisualColor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtWorkflowTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtWorkflowTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtWorkflowTagCvId"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtWorkflowTagCvId = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtWorkflowTagCvTermId"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtWorkflowTagCvTermId = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtWorkflowTagCvTermName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtWorkflowTagCvTermName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCExtWorkflowTagCvTermRefinedAbout"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCExtWorkflowTagCvTermRefinedAbout = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCFixtureIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCFixtureIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCHeadline"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCHeadline = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCImageOrientation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCImageOrientation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCImageType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCImageType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCKeywords"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCKeywords = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCLanguageIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCLanguageIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCObjectAttributeReference"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCObjectAttributeReference = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCObjectCycle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCObjectCycle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCObjectName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCObjectName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCObjectTypeReference"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCObjectTypeReference = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCOriginalTransmissionReference"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCOriginalTransmissionReference = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCOriginatingProgram"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCOriginatingProgram = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCProgramVersion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCProgramVersion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCProvinceState"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCProvinceState = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCReferenceDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCReferenceDate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCReferenceNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCReferenceNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCReferenceService"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCReferenceService = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCReleaseDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCReleaseDate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCReleaseTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCReleaseTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCRightsUsageTerms"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCRightsUsageTerms = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCScene"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCScene = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCSource"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCSource = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCSpecialInstructions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCSpecialInstructions = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCStarRating"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCStarRating = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCSubLocation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCSubLocation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCSubjectReference"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCSubjectReference = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCSupplementalCategory"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCSupplementalCategory = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCTimeCreated"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCTimeCreated = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCUrgency"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCUrgency = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIPTCWriterEditor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIPTCWriterEditor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyImageCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyImageCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyImageIndex"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyImageIndex = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyImages"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyImages = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIsFloat"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIsFloat = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyIsIndexed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyIsIndexed = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyJFIFDensityUnit"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyJFIFDensityUnit = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyJFIFDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyJFIFDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyJFIFIsProgressive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyJFIFIsProgressive = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyJFIFVersion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyJFIFVersion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyJFIFXDensity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyJFIFXDensity = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyJFIFYDensity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyJFIFYDensity = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerAppleDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerAppleDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerCanonAspectRatioInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerCanonAspectRatioInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerCanonCameraSerialNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerCanonCameraSerialNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerCanonContinuousDrive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerCanonContinuousDrive = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerCanonDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerCanonDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerCanonFirmware"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerCanonFirmware = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerCanonFlashExposureComp"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerCanonFlashExposureComp = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerCanonImageSerialNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerCanonImageSerialNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerCanonLensModel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerCanonLensModel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerCanonOwnerName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerCanonOwnerName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerFujiDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerFujiDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerMinoltaDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerMinoltaDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerNikonCameraSerialNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerNikonCameraSerialNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerNikonColorMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerNikonColorMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerNikonDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerNikonDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerNikonDigitalZoom"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerNikonDigitalZoom = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerNikonFlashExposureComp"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerNikonFlashExposureComp = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerNikonFlashSetting"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerNikonFlashSetting = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerNikonFocusDistance"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerNikonFocusDistance = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerNikonFocusMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerNikonFocusMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerNikonISOSelection"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerNikonISOSelection = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerNikonISOSetting"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerNikonISOSetting = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerNikonImageAdjustment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerNikonImageAdjustment = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerNikonLensAdapter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerNikonLensAdapter = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerNikonLensInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerNikonLensInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerNikonLensType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerNikonLensType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerNikonQuality"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerNikonQuality = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerNikonSharpenMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerNikonSharpenMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerNikonShootingMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerNikonShootingMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerNikonShutterCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerNikonShutterCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerNikonWhiteBalanceMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerNikonWhiteBalanceMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerOlympusDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerOlympusDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyMakerPentaxDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyMakerPentaxDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyNamedColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyNamedColorSpace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyOpenEXRAspectRatio"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyOpenEXRAspectRatio = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyOpenEXRCompression"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyOpenEXRCompression = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyOpenEXRDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyOpenEXRDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyOrientation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyOrientation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGAuthor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGAuthor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGChromaticities"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGChromaticities = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGComment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGComment = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGCompressionFilter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGCompressionFilter = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGCopyright"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGCopyright = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGCreationTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGCreationTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGDisclaimer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGDisclaimer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGGamma"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGGamma = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGInterlaceType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGInterlaceType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGModificationTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGModificationTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGPixelsAspectRatio"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGPixelsAspectRatio = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGSoftware"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGSoftware = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGSource"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGSource = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGTitle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGTransparency"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGTransparency = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGWarning"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGWarning = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGXPixelsPerMeter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGXPixelsPerMeter = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGYPixelsPerMeter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGYPixelsPerMeter = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPNGsRGBIntent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPNGsRGBIntent = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPVREncoder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPVREncoder = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPixelFormat"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPixelFormat = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPixelHeight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPixelHeight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPixelWidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPixelWidth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyPrimaryImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyPrimaryImage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyProfileName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyProfileName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyRawDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyRawDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTGACompression"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTGACompression = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTGADictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTGADictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFArtist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFArtist = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFCompression"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFCompression = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFCopyright"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFCopyright = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFDateTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFDateTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFDocumentName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFDocumentName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFHostComputer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFHostComputer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFImageDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFImageDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFMake"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFMake = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFModel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFModel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFOrientation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFOrientation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFPhotometricInterpretation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFPhotometricInterpretation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFPrimaryChromaticities"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFPrimaryChromaticities = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFResolutionUnit"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFResolutionUnit = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFSoftware"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFSoftware = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFTileLength"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFTileLength = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFTileWidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFTileWidth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFTransferFunction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFTransferFunction = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFWhitePoint"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFWhitePoint = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFXPosition"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFXPosition = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFXResolution"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFXResolution = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFYPosition"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFYPosition = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyTIFFYResolution"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyTIFFYResolution = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyThumbnailImages"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyThumbnailImages = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyWebPCanvasPixelHeight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyWebPCanvasPixelHeight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyWebPCanvasPixelWidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyWebPCanvasPixelWidth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyWebPDelayTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyWebPDelayTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyWebPDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyWebPDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyWebPFrameInfoArray"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyWebPFrameInfoArray = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyWebPLoopCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyWebPLoopCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyWebPUnclampedDelayTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyWebPUnclampedDelayTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImagePropertyWidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImagePropertyWidth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageProviderPreferredTileHeight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageProviderPreferredTileHeight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageProviderPreferredTileWidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageProviderPreferredTileWidth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageSourceCreateThumbnailFromImageAlways"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageSourceCreateThumbnailFromImageAlways = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageSourceCreateThumbnailFromImageIfAbsent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageSourceCreateThumbnailFromImageIfAbsent = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageSourceCreateThumbnailWithTransform"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageSourceCreateThumbnailWithTransform = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageSourceDecodeRequest"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageSourceDecodeRequest = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageSourceDecodeRequestOptions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageSourceDecodeRequestOptions = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageSourceDecodeToHDR"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageSourceDecodeToHDR = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageSourceDecodeToSDR"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageSourceDecodeToSDR = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageSourceGenerateImageSpecificLumaScaling"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageSourceGenerateImageSpecificLumaScaling = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageSourceShouldAllowFloat"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageSourceShouldAllowFloat = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageSourceShouldCache"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageSourceShouldCache = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageSourceShouldCacheImmediately"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageSourceShouldCacheImmediately = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageSourceSubsampleFactor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageSourceSubsampleFactor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageSourceThumbnailMaxPixelSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageSourceThumbnailMaxPixelSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGImageSourceTypeIdentifierHint"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGImageSourceTypeIdentifierHint = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIIOCameraExtrinsics_CoordinateSystemID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIIOCameraExtrinsics_CoordinateSystemID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIIOCameraExtrinsics_Position"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIIOCameraExtrinsics_Position = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIIOCameraExtrinsics_Rotation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIIOCameraExtrinsics_Rotation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIIOCameraModelType_GenericPinhole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIIOCameraModelType_GenericPinhole = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIIOCameraModelType_SimplifiedPinhole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIIOCameraModelType_SimplifiedPinhole = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIIOCameraModel_Intrinsics"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIIOCameraModel_Intrinsics = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIIOCameraModel_ModelType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIIOCameraModel_ModelType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIIOMetadata_CameraExtrinsicsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIIOMetadata_CameraExtrinsicsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIIOMetadata_CameraModelKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIIOMetadata_CameraModelKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIIOMonoscopicImageLocation_Center"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIIOMonoscopicImageLocation_Center = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIIOMonoscopicImageLocation_Left"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIIOMonoscopicImageLocation_Left = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIIOMonoscopicImageLocation_Right"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIIOMonoscopicImageLocation_Right = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIIOMonoscopicImageLocation_Unspecified"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIIOMonoscopicImageLocation_Unspecified = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIIOStereoAggressors_Severity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIIOStereoAggressors_Severity = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIIOStereoAggressors_SubTypeURI"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIIOStereoAggressors_SubTypeURI = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIIOStereoAggressors_Type"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIIOStereoAggressors_Type = objc.GoString(cstr)
			}
		}
	}

}
