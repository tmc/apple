// Code generated from Apple documentation. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// CIDetectorAccuracy is a key used to specify the desired accuracy for the detector.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDetectorAccuracy
	CIDetectorAccuracy string
	// CIDetectorAccuracyHigh is indicates that the detector should choose techniques that are higher in accuracy, even if it requires more processing time.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDetectorAccuracyHigh
	CIDetectorAccuracyHigh string
	// CIDetectorAccuracyLow is indicates that the detector should choose techniques that are lower in accuracy, but can be processed more quickly.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDetectorAccuracyLow
	CIDetectorAccuracyLow string
	// CIDetectorAspectRatio is an option specifying the aspect ratio (width divided by height) of rectangles to search for.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDetectorAspectRatio
	CIDetectorAspectRatio string
	// CIDetectorEyeBlink is an option for whether Core Image will perform additional processing to recognize closed eyes in detected faces.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDetectorEyeBlink
	CIDetectorEyeBlink string
	// CIDetectorFocalLength is an option identifying the focal length in pixels used in capturing images to be processed by the detector.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDetectorFocalLength
	CIDetectorFocalLength string
	// CIDetectorImageOrientation is an option for the display orientation of the image whose features you want to detect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDetectorImageOrientation
	CIDetectorImageOrientation string
	// CIDetectorMaxFeatureCount is the key to the configuration dictionary whose value represents the maximum number of features the detector should return.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDetectorMaxFeatureCount
	CIDetectorMaxFeatureCount string
	// CIDetectorMinFeatureSize is a key used to specify the minimum size that the detector will recognize as a feature.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDetectorMinFeatureSize
	CIDetectorMinFeatureSize string
	// CIDetectorNumberOfAngles is the number of perspectives to use for detecting a face in video input.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDetectorNumberOfAngles
	CIDetectorNumberOfAngles string
	// CIDetectorReturnSubFeatures is an option specifying whether to return feature information for components of detected features.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDetectorReturnSubFeatures
	CIDetectorReturnSubFeatures string
	// CIDetectorSmile is an option for whether Core Image will perform additional processing to recognize smiles in detected faces.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDetectorSmile
	CIDetectorSmile string
	// CIDetectorTracking is a key used to enable or disable face tracking for the detector. Use this option when you want to track faces across frames in a video.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDetectorTracking
	CIDetectorTracking string
	// CIDetectorTypeFace is a detector that searches for faces in a still image or video, returning [CIFaceFeature] objects that provide information about detected faces.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDetectorTypeFace
	CIDetectorTypeFace string
	// CIDetectorTypeQRCode is a detector that searches for Quick Response codes (a type of 2D barcode) in a still image or video, returning [CIQRCodeFeature] objects that provide information about detected barcodes.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDetectorTypeQRCode
	CIDetectorTypeQRCode string
	// CIDetectorTypeRectangle is a detector that searches for rectangular areas in a still image or video, returning [CIRectangleFeature] objects that provide information about detected regions.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDetectorTypeRectangle
	CIDetectorTypeRectangle string
	// CIDetectorTypeText is a detector that searches for text in a still image or video, returning [CITextFeature] objects that provide information about detected regions.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDetectorTypeText
	CIDetectorTypeText string
	// CIFeatureTypeFace is a Core Image feature type for person’s face.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFeatureTypeFace
	CIFeatureTypeFace string
	// CIFeatureTypeQRCode is a Core Image feature type for QR code object.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFeatureTypeQRCode
	CIFeatureTypeQRCode string
	// CIFeatureTypeRectangle is a Core Image feature type for rectangular object.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFeatureTypeRectangle
	CIFeatureTypeRectangle string
	// CIFeatureTypeText is a Core Image feature type for text.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFeatureTypeText
	CIFeatureTypeText string
	// KCIApplyOptionColorSpace is the color space of the produced image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIApplyOptionColorSpace
	KCIApplyOptionColorSpace string
	// KCIApplyOptionDefinition is the domain of definition (DOD) of the produced image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIApplyOptionDefinition
	KCIApplyOptionDefinition string
	// KCIApplyOptionExtent is the extent of the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIApplyOptionExtent
	KCIApplyOptionExtent string
	// KCIApplyOptionUserInfo is information needed by a callback. The associated value is an object that Core Image will pass to any callbacks invoked for that filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIApplyOptionUserInfo
	KCIApplyOptionUserInfo string
	// KCIAttributeClass is the class name of the filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeClass
	KCIAttributeClass string
	// KCIAttributeDefault is the default value, specified as a floating-point value, for a filter parameter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeDefault
	KCIAttributeDefault string
	// KCIAttributeDescription is the localized description of the filter. This description should inform the user what the filter does and be short enough to display in the user interface for the filter. It is not intended to be technically detailed.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeDescription
	KCIAttributeDescription string
	// KCIAttributeDisplayName is the localized display name of the attribute.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeDisplayName
	KCIAttributeDisplayName string
	// KCIAttributeFilterAvailable_Mac is the macOS version in which the filter first became available.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeFilterAvailable_Mac
	KCIAttributeFilterAvailable_Mac string
	// KCIAttributeFilterAvailable_iOS is the iOS version in which the filter first became available.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeFilterAvailable_iOS
	KCIAttributeFilterAvailable_iOS string
	// KCIAttributeFilterCategories is an array of filter category keys that specifies all the categories in which the filter is a member.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeFilterCategories
	KCIAttributeFilterCategories string
	// KCIAttributeFilterDisplayName is the localized version of the filter name that is displayed in the user interface.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeFilterDisplayName
	KCIAttributeFilterDisplayName string
	// KCIAttributeFilterName is the filter name.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeFilterName
	KCIAttributeFilterName string
	// KCIAttributeIdentity is if supplied as a value for a parameter, the parameter has no effect on the input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeIdentity
	KCIAttributeIdentity string
	// KCIAttributeMax is the maximum value for a filter parameter, specified as a floating-point value.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeMax
	KCIAttributeMax string
	// KCIAttributeMin is the minimum value for a filter parameter, specified as a floating-point value.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeMin
	KCIAttributeMin string
	// KCIAttributeName is the name of the attribute.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeName
	KCIAttributeName string
	// KCIAttributeReferenceDocumentation is the localized reference documentation for the filter. The reference should provide developers with technical details.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeReferenceDocumentation
	KCIAttributeReferenceDocumentation string
	// KCIAttributeSliderMax is the maximum value, specified as a floating-point value, to use for a slider that controls input values for a filter parameter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeSliderMax
	KCIAttributeSliderMax string
	// KCIAttributeSliderMin is the minimum value, specified as a floating-point value, to use for a slider that controls input values for a filter parameter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeSliderMin
	KCIAttributeSliderMin string
	// KCIAttributeType is the type of an attribute.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeType
	KCIAttributeType string
	// KCIAttributeTypeAngle is an angle.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeTypeAngle
	KCIAttributeTypeAngle string
	// KCIAttributeTypeBoolean is a Boolean value.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeTypeBoolean
	KCIAttributeTypeBoolean string
	// KCIAttributeTypeColor is a Core Image color ([CIColor] object) that specifies red, green, and blue component values.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeTypeColor
	KCIAttributeTypeColor string
	// KCIAttributeTypeCount is a positive integer value.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeTypeCount
	KCIAttributeTypeCount string
	// KCIAttributeTypeDistance is a distance.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeTypeDistance
	KCIAttributeTypeDistance string
	// KCIAttributeTypeGradient is an n-by-1 gradient image used to describe a color ramp.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeTypeGradient
	KCIAttributeTypeGradient string
	// KCIAttributeTypeImage is a [CIImage] object.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeTypeImage
	KCIAttributeTypeImage string
	// KCIAttributeTypeInteger is an integer value.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeTypeInteger
	KCIAttributeTypeInteger string
	// KCIAttributeTypeOffset is an offset. (A 2-element vector type.).
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeTypeOffset
	KCIAttributeTypeOffset string
	// KCIAttributeTypeOpaqueColor is a Core Image color ([CIColor] object) that specifies red, green, and blue component values. Use this key for colors with no alpha component. If the key is not present, Core Image assumes color with alpha.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeTypeOpaqueColor
	KCIAttributeTypeOpaqueColor string
	// KCIAttributeTypePosition is a two-dimensional location in the working coordinate space. (A 2-element vector type.).
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeTypePosition
	KCIAttributeTypePosition string
	// KCIAttributeTypePosition3 is a three-dimensional location in the working coordinate space. (A 3-element vector type.).
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeTypePosition3
	KCIAttributeTypePosition3 string
	// KCIAttributeTypeRectangle is a Core Image vector that specifies the and values of the rectangle origin, and the width () and height () of the rectangle. The vector takes the form [, , , ]. (A 4-element vector type.).
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeTypeRectangle
	KCIAttributeTypeRectangle string
	// KCIAttributeTypeScalar is a scalar value.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeTypeScalar
	KCIAttributeTypeScalar string
	// KCIAttributeTypeTime is a parametric time for transitions, specified as a floating-point value in the range of `0.0` to `1.0`.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeTypeTime
	KCIAttributeTypeTime string
	// KCIAttributeTypeTransform is the transform type of an attribute.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIAttributeTypeTransform
	KCIAttributeTypeTransform string
	// KCICategoryBlur is a filter that softens images, decreasing the contrast between the edges in an image. Examples of blur filters are Gaussian blur and zoom blur.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategoryBlur
	KCICategoryBlur string
	// KCICategoryBuiltIn is a filter provided by Core Image. This distinguishes built-in filters from plug-in filters.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategoryBuiltIn
	KCICategoryBuiltIn string
	// KCICategoryColorAdjustment is the category for color adjustment filters.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategoryColorAdjustment
	KCICategoryColorAdjustment string
	// KCICategoryColorEffect is a filter that modifies the color of an image to achieve an artistic effect. Examples of color effect filters include filters that change a color image to a sepia image or a monochrome image or that produces such effects as posterizing.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategoryColorEffect
	KCICategoryColorEffect string
	// KCICategoryCompositeOperation is a filter operates on two image sources, using the color values of one image to operate on the other. Composite filters perform computations such as computing maximum values, minimum values, and multiplying values between input images. You can use compositing filters to add effects to an image, crop an image, and achieve a variety of other effects.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategoryCompositeOperation
	KCICategoryCompositeOperation string
	// KCICategoryDistortionEffect is a filter that reshapes an image by altering its geometry to create a 3D effect. Using distortion filters, you can displace portions of an image, apply lens effects, make a bulge in an image, and perform other operation to achieve an artistic effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategoryDistortionEffect
	KCICategoryDistortionEffect string
	// KCICategoryFilterGenerator is a filter created by chaining several filters together and then packaged as a [CIFilterGenerator] object.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategoryFilterGenerator
	KCICategoryFilterGenerator string
	// KCICategoryGenerator is a filter that generates a pattern, such as a solid color, a checkerboard, or a star shine. The generated output is typically used as input to another filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategoryGenerator
	KCICategoryGenerator string
	// KCICategoryGeometryAdjustment is a filter that changes the geometry of an image. Some of these filters are used to warp an image to achieve an artistic effects, but these filters can also be used to correct problems in the source image. For example, you can apply an affine transform to straighten an image that is rotated with respect to the horizon.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategoryGeometryAdjustment
	KCICategoryGeometryAdjustment string
	// KCICategoryGradient is a filter that generates a fill whose color varies smoothly. Exactly how color varies depends on the type of gradient—linear, radial, or Gaussian.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategoryGradient
	KCICategoryGradient string
	// KCICategoryHalftoneEffect is a filter that simulates a variety of halftone screens, to mimic the halftone process used in print media. The output of these filters has the familiar “newspaper” look of the various dot patterns. Filters are typically named after the pattern created by the virtual halftone screen, such as circular screen or hatched screen.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategoryHalftoneEffect
	KCICategoryHalftoneEffect string
	// KCICategoryHighDynamicRange is a filter that works on high dynamic range pixels.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategoryHighDynamicRange
	KCICategoryHighDynamicRange string
	// KCICategoryInterlaced is a filter that works on interlaced images.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategoryInterlaced
	KCICategoryInterlaced string
	// KCICategoryNonSquarePixels is a filter that works on non-square pixels.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategoryNonSquarePixels
	KCICategoryNonSquarePixels string
	// KCICategoryReduction is a filter that reduces image data. These filters are used to solve image analysis problems.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategoryReduction
	KCICategoryReduction string
	// KCICategorySharpen is a filter that sharpens images, increasing the contrast between the edges in an image. Examples of sharpen filters are unsharp mask and sharpen luminance.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategorySharpen
	KCICategorySharpen string
	// KCICategoryStillImage is a filter that works on still images.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategoryStillImage
	KCICategoryStillImage string
	// KCICategoryStylize is a filter that makes a photographic image look as if it was painted or sketched. These filters are typically used alone or in combination with other filters to achieve artistic effects.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategoryStylize
	KCICategoryStylize string
	// KCICategoryTileEffect is a filter that typically applies an effect to an image and then create smaller versions of the image (tiles), which are then laid out to create a pattern that’s infinite in extent.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategoryTileEffect
	KCICategoryTileEffect string
	// KCICategoryTransition is a filter that provides a bridge between two or more images by applying a motion effect that defines how the pixels of a source image yield to that of the destination image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategoryTransition
	KCICategoryTransition string
	// KCICategoryVideo is a filter that works on video images.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCICategoryVideo
	KCICategoryVideo string
	// KCIFilterGeneratorExportedKey is the key ([CIFilterGeneratorExportedKey]) for the exported parameter. The associated value is the key name of the parameter you are exporting, such as `inputRadius`.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIFilterGeneratorExportedKey
	KCIFilterGeneratorExportedKey string
	// KCIFilterGeneratorExportedKeyName is the key ([CIFilterGeneratorExportedKeyName]) for the name used to export the [CIFilterGenerator] object. The associated value is a string that specifies a unique name for the filter generator object.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIFilterGeneratorExportedKeyName
	KCIFilterGeneratorExportedKeyName string
	// KCIFilterGeneratorExportedKeyTargetObject is the target object ([CIFilterGeneratorExportedKeyTargetObject]) for the exported key. The associated value is the name of the object, such as [CIMotionBlur].
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIFilterGeneratorExportedKeyTargetObject
	KCIFilterGeneratorExportedKeyTargetObject string
	// KCIImageAutoAdjustCrop is a key used to specify whether to return a filter that crops the image to focus on detected features.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageAutoAdjustmentOption/crop
	KCIImageAutoAdjustCrop string
	// KCIImageAutoAdjustEnhance is a key used to specify whether to return enhancement filters.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageAutoAdjustmentOption/enhance
	KCIImageAutoAdjustEnhance string
	// KCIImageAutoAdjustFeatures is a key used to specify an array of features that you want to apply enhancement and red eye filters to.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageAutoAdjustmentOption/features
	KCIImageAutoAdjustFeatures string
	// KCIImageAutoAdjustLevel is a key used to specify whether to return a filter that rotates the image to keep a level perspective.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageAutoAdjustmentOption/level
	KCIImageAutoAdjustLevel string
	// KCIImageAutoAdjustRedEye is a key used to specify whether to return a red eye filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageAutoAdjustmentOption/redEye
	KCIImageAutoAdjustRedEye string
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputAmountKey
	KCIInputAmountKey string
	// KCIInputAngleKey is the angle.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputAngleKey
	KCIInputAngleKey string
	// KCIInputAspectRatioKey is aspect Ratio.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputAspectRatioKey
	KCIInputAspectRatioKey string
	// KCIInputBackgroundImageKey is a key for the [CIImage] object to use as a background image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputBackgroundImageKey
	KCIInputBackgroundImageKey string
	// KCIInputBacksideImageKey is a key to get or set the backside image for a transition Core Image filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputBacksideImageKey
	KCIInputBacksideImageKey string
	// KCIInputBiasKey is simple bias value.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputBiasKey
	KCIInputBiasKey string
	// KCIInputBiasVectorKey is a key to get or set the vector bias value of a Core Image filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputBiasVectorKey
	KCIInputBiasVectorKey string
	// KCIInputBrightnessKey is brightness level.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputBrightnessKey
	KCIInputBrightnessKey string
	// KCIInputCenterKey is a key for a [CIVector] object that specifies the center of the area, as and - coordinates, to be filtered.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputCenterKey
	KCIInputCenterKey string
	// KCIInputColor0Key is a key to get or set a color value of a Core Image filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputColor0Key
	KCIInputColor0Key string
	// KCIInputColor1Key is a key to get or set a color value of a Core Image filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputColor1Key
	KCIInputColor1Key string
	// KCIInputColorKey is a key for a [CIColor] object that specifies a color value.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputColorKey
	KCIInputColorKey string
	// KCIInputColorSpaceKey is a key to get or set a color space value of a Core Image filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputColorSpaceKey
	KCIInputColorSpaceKey string
	// KCIInputContrastKey is a contrast level.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputContrastKey
	KCIInputContrastKey string
	// KCIInputCountKey is a key to get or set the scalar count value of a Core Image filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputCountKey
	KCIInputCountKey string
	// KCIInputDepthImageKey is a key for an image with depth values.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputDepthImageKey
	KCIInputDepthImageKey string
	// KCIInputDisparityImageKey is a key for an image with disparity values.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputDisparityImageKey
	KCIInputDisparityImageKey string
	// KCIInputEVKey is how many F-stops brighter or darker the image should be.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputEVKey
	KCIInputEVKey string
	// KCIInputExtentKey is a key for a [CIVector] object that specifies a rectangle that defines the extent of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputExtentKey
	KCIInputExtentKey string
	// KCIInputExtrapolateKey is a key to get or set the boolean behavior of a Core Image filter that specifies if the filter should extrapolate a table beyond the defined range.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputExtrapolateKey
	KCIInputExtrapolateKey string
	// KCIInputGradientImageKey is a key for a [CIImage] object that specifies an environment map with alpha. Typically, this image contains highlight and shadow.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputGradientImageKey
	KCIInputGradientImageKey string
	// KCIInputImageKey is a key for the [CIImage] object to use as an input image. For filters that also use a background image, this key refers to the foreground image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputImageKey
	KCIInputImageKey string
	// KCIInputIntensityKey is an intensity value.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputIntensityKey
	KCIInputIntensityKey string
	// KCIInputMaskImageKey is a key for a [CIImage] object to use as a mask.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputMaskImageKey
	KCIInputMaskImageKey string
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputMatteImageKey
	KCIInputMatteImageKey string
	// KCIInputPaletteImageKey is a key to get or set the palette image for a Core Image filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputPaletteImageKey
	KCIInputPaletteImageKey string
	// KCIInputPerceptualKey is a key to get or set the boolean behavior of a Core Image filter that specifies if the filter should operate in linear or perceptual colors.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputPerceptualKey
	KCIInputPerceptualKey string
	// KCIInputPoint0Key is a key to get or set the coordinate value of a Core Image filter. The value for this key needs to be a [CIVector] instance containing the `x,y` coordinate.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputPoint0Key
	KCIInputPoint0Key string
	// KCIInputPoint1Key is a key to get or set a coordinate value of a Core Image filter. The value for this key needs to be a [CIVector] instance containing the `x,y` coordinate.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputPoint1Key
	KCIInputPoint1Key string
	// KCIInputRadius0Key is a key to get or set the geometric radius value of a Core Image filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputRadius0Key
	KCIInputRadius0Key string
	// KCIInputRadius1Key is a key to get or set the geometric radius value of a Core Image filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputRadius1Key
	KCIInputRadius1Key string
	// KCIInputRadiusKey is the distance from the center of an effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputRadiusKey
	KCIInputRadiusKey string
	// KCIInputRefractionKey is the index of refraction to use.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputRefractionKey
	KCIInputRefractionKey string
	// KCIInputSaturationKey is the amount to adjust the saturation.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputSaturationKey
	KCIInputSaturationKey string
	// KCIInputScaleKey is the amount of scale to apply.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputScaleKey
	KCIInputScaleKey string
	// KCIInputShadingImageKey is a key for a [CIImage] object that specifies an environment map with alpha values. Typically this image contains highlight and shadow.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputShadingImageKey
	KCIInputShadingImageKey string
	// KCIInputSharpnessKey is amount of sharpening to apply.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputSharpnessKey
	KCIInputSharpnessKey string
	// KCIInputTargetImageKey is a key for a [CIImage] object that is the target image for a transition.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputTargetImageKey
	KCIInputTargetImageKey string
	// KCIInputThresholdKey is a key to get or set the scalar threshold value of a Core Image filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputThresholdKey
	KCIInputThresholdKey string
	// KCIInputTimeKey is specify a time.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputTimeKey
	KCIInputTimeKey string
	// KCIInputTransformKey is transformation to apply.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputTransformKey
	KCIInputTransformKey string
	// KCIInputVersionKey is version Key.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputVersionKey
	KCIInputVersionKey string
	// KCIInputWeightsKey is a key for a [CIVector] object that describes a weight matrix for use with a convolution filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputWeightsKey
	KCIInputWeightsKey string
	// KCIInputWidthKey is a key for a scalar value that specifies the width of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIInputWidthKey
	KCIInputWidthKey string
	// KCIOutputImageKey is a key for the [CIImage] object produced by a filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIOutputImageKey
	KCIOutputImageKey string
	// KCISamplerAffineMatrix is the key for an affine matrix. The associated value is an [NSArray] object ([]) that defines the transformation to apply to the sampler.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCISamplerAffineMatrix
	KCISamplerAffineMatrix string
	// KCISamplerColorSpace is the key for the color space to use when sampling the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCISamplerColorSpace
	KCISamplerColorSpace string
	// KCISamplerFilterLinear is bilinear interpolation.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCISamplerFilterLinear
	KCISamplerFilterLinear string
	// KCISamplerFilterMode is the key for the filtering to use when sampling the image. Possible values are [kCISamplerFilterNearest] and [kCISamplerFilterLinear].
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCISamplerFilterMode
	KCISamplerFilterMode string
	// KCISamplerFilterNearest is nearest neighbor sampling.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCISamplerFilterNearest
	KCISamplerFilterNearest string
	// KCISamplerWrapBlack is pixels are transparent black.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCISamplerWrapBlack
	KCISamplerWrapBlack string
	// KCISamplerWrapClamp is coordinates are clamped to the extent.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCISamplerWrapClamp
	KCISamplerWrapClamp string
	// KCISamplerWrapMode is the key for the sampler wrap mode. The wrap mode specifies how Core Image produces pixels that are outside the extent of the sample. Possible values are [kCISamplerWrapBlack] and [kCISamplerWrapClamp].
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCISamplerWrapMode
	KCISamplerWrapMode string
	// KCIUIParameterSet is the set of input parameters to use. The associated value can be [kCIUISetBasic], [kCIUISetIntermediate], [kCIUISetAdvanced], or [kCIUISetDevelopment].
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIUIParameterSet
	KCIUIParameterSet string
	// KCIUISetAdvanced is controls that are appropriate for an advanced user scenario.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIUISetAdvanced
	KCIUISetAdvanced string
	// KCIUISetBasic is controls that are appropriate for a basic user scenario, that is, the minimum of settings to control the filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIUISetBasic
	KCIUISetBasic string
	// KCIUISetDevelopment is controls that should be visible only for development purposes.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIUISetDevelopment
	KCIUISetDevelopment string
	// KCIUISetIntermediate is controls that are appropriate for an intermediate user scenario.
	//
	// See: https://developer.apple.com/documentation/CoreImage/kCIUISetIntermediate
	KCIUISetIntermediate string
)

var (
	// See: https://developer.apple.com/documentation/CoreImage/CIRAWDecoderVersion/version6
	CIRAWDecoderVersion6 CIRAWDecoderVersion
	// See: https://developer.apple.com/documentation/CoreImage/CIRAWDecoderVersion/version6DNG
	CIRAWDecoderVersion6DNG CIRAWDecoderVersion
	// See: https://developer.apple.com/documentation/CoreImage/CIRAWDecoderVersion/version7
	CIRAWDecoderVersion7 CIRAWDecoderVersion
	// See: https://developer.apple.com/documentation/CoreImage/CIRAWDecoderVersion/version7DNG
	CIRAWDecoderVersion7DNG CIRAWDecoderVersion
	// See: https://developer.apple.com/documentation/CoreImage/CIRAWDecoderVersion/version8
	CIRAWDecoderVersion8 CIRAWDecoderVersion
	// See: https://developer.apple.com/documentation/CoreImage/CIRAWDecoderVersion/version8DNG
	CIRAWDecoderVersion8DNG CIRAWDecoderVersion
	// See: https://developer.apple.com/documentation/CoreImage/CIRAWDecoderVersion/version9
	CIRAWDecoderVersion9 CIRAWDecoderVersion
	// See: https://developer.apple.com/documentation/CoreImage/CIRAWDecoderVersion/version9DNG
	CIRAWDecoderVersion9DNG CIRAWDecoderVersion
	// See: https://developer.apple.com/documentation/CoreImage/CIRAWDecoderVersion/none
	CIRAWDecoderVersionNone CIRAWDecoderVersion
)

var (
	// KCIContextAllowLowPower is a Boolean value to control the power level of Core Image context renders.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIContextOption/allowLowPower
	KCIContextAllowLowPower CIContextOption
	// KCIContextCVMetalTextureCache is a Core Video Metal texture cache object to improve the performance of Core Image context renders that use Core Video pixel buffers.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIContextOption/cvMetalTextureCache
	KCIContextCVMetalTextureCache CIContextOption
	// KCIContextCacheIntermediates is a Boolean value to control how a Core Image context caches the contents of any intermediate image buffers it uses during rendering.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIContextOption/cacheIntermediates
	KCIContextCacheIntermediates CIContextOption
	// KCIContextHighQualityDownsample is a Boolean value to control the quality of image downsampling operations performed by the Core Image context.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIContextOption/highQualityDownsample
	KCIContextHighQualityDownsample CIContextOption
	// KCIContextMemoryLimit is a number value to control the maximum memory in megabytes that the context allocates for render tasks.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIContextOption/memoryTarget
	KCIContextMemoryLimit CIContextOption
	// KCIContextName is a Boolean value to specify a client-provided name for a context.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIContextOption/name
	KCIContextName CIContextOption
	// KCIContextOutputColorSpace is a Core Image context option key to specify the default destination color space for rendering.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIContextOption/outputColorSpace
	KCIContextOutputColorSpace CIContextOption
	// KCIContextOutputPremultiplied is a Boolean value to control how a Core Image context render produces alpha-premultiplied pixels.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIContextOption/outputPremultiplied
	KCIContextOutputPremultiplied CIContextOption
	// KCIContextPriorityRequestLow is a Boolean value to control the priority Core Image context renders.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIContextOption/priorityRequestLow
	KCIContextPriorityRequestLow CIContextOption
	// KCIContextUseSoftwareRenderer is a Boolean value to control if a Core Image context will use a software renderer.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIContextOption/useSoftwareRenderer
	KCIContextUseSoftwareRenderer CIContextOption
	// KCIContextWorkingColorSpace is a Core Image context option key to specify the working color space for rendering.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIContextOption/workingColorSpace
	KCIContextWorkingColorSpace CIContextOption
	// KCIContextWorkingFormat is a Core Image context option key to specify the pixel format to for intermediate results when rendering.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIContextOption/workingFormat
	KCIContextWorkingFormat CIContextOption
)

var (
	// KCIDynamicRangeConstrainedHigh is use extended dynamic range, but brightness is modulated to optimize for co-existence with other composited content.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDynamicRangeOption/constrainedHigh
	KCIDynamicRangeConstrainedHigh CIDynamicRangeOption
	// KCIDynamicRangeHigh is use High dynamic range.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDynamicRangeOption/high
	KCIDynamicRangeHigh CIDynamicRangeOption
	// KCIDynamicRangeStandard is use Standard dynamic range.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDynamicRangeOption/standard
	KCIDynamicRangeStandard CIDynamicRangeOption
)

var (
	// KCIFormatA16 is a 16-bit-per-pixel, fixed-point pixel format in which the sole component is alpha.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/A16
	KCIFormatA16 int
	// KCIFormatA8 is an 8-bit-per-pixel, fixed-point pixel format in which the sole component is alpha.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/A8
	KCIFormatA8 int
	// KCIFormatABGR8 is a 32-bit-per-pixel, fixed-point pixel format in which the alpha value precedes the blue, green, and red color components.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/ABGR8
	KCIFormatABGR8 int
	// KCIFormatARGB8 is a 32-bit-per-pixel, fixed-point pixel format in which the alpha value precedes the red, green, and blue color components.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/ARGB8
	KCIFormatARGB8 int
	// KCIFormatAf is a 32-bit-per-pixel, full-width floating-point pixel format in which the sole component is alpha.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/Af
	KCIFormatAf int
	// KCIFormatAh is a 16-bit-per-pixel, half-width floating-point pixel format in which the sole component is alpha.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/Ah
	KCIFormatAh int
	// KCIFormatBGRA8 is a 32-bit-per-pixel, fixed-point pixel format in which the blue, green, and red color components precede the alpha value.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/BGRA8
	KCIFormatBGRA8 int
	// KCIFormatL16 is a 16-bit-per-pixel, fixed-point pixel format in which the sole component is luminance.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/L16
	KCIFormatL16 int
	// KCIFormatL8 is an 8-bit-per-pixel, fixed-point pixel format in which the sole component is luminance.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/L8
	KCIFormatL8 int
	// KCIFormatLA16 is a 32-bit-per-pixel, fixed-point pixel format with only 16-bit luminance and alpha components.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/LA16
	KCIFormatLA16 int
	// KCIFormatLA8 is a 16-bit-per-pixel, fixed-point pixel format with only 8-bit luminance and alpha components.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/LA8
	KCIFormatLA8 int
	// KCIFormatLAf is a 64-bit-per-pixel, full-width floating-point pixel format with 32-bit luminance and alpha components.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/LAf
	KCIFormatLAf int
	// KCIFormatLAh is a 32-bit-per-pixel, half-width floating-point pixel format with 16-bit luminance and alpha components.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/LAh
	KCIFormatLAh int
	// KCIFormatLf is a 32-bit-per-pixel, full-width floating-point pixel format in which the sole component is luminance.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/Lf
	KCIFormatLf int
	// KCIFormatLh is a 16-bit-per-pixel, half-width floating-point pixel format in which the sole component is luminance.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/Lh
	KCIFormatLh int
	// KCIFormatR16 is a 16-bit-per-pixel, fixed-point pixel format in which the sole component is a red color value.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/R16
	KCIFormatR16 int
	// KCIFormatR8 is an 8-bit-per-pixel, fixed-point pixel format in which the sole component is a red color value.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/R8
	KCIFormatR8 int
	// KCIFormatRG16 is a 32-bit-per-pixel, fixed-point pixel format with only red and green color components.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/RG16
	KCIFormatRG16 int
	// KCIFormatRG8 is a 16-bit-per-pixel, fixed-point pixel format with only red and green color components.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/RG8
	KCIFormatRG8 int
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/RGB10
	KCIFormatRGB10 int
	// KCIFormatRGBA16 is a 64-bit-per-pixel, fixed-point pixel format.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/RGBA16
	KCIFormatRGBA16 int
	// KCIFormatRGBA8 is a 32-bit-per-pixel, fixed-point pixel format in which the red, green, and blue color components precede the alpha value.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/RGBA8
	KCIFormatRGBA8 int
	// KCIFormatRGBAf is a 128-bit-per-pixel, floating-point pixel format.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/RGBAf
	KCIFormatRGBAf int
	// KCIFormatRGBAh is a 64-bit-per-pixel, floating-point pixel format.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/RGBAh
	KCIFormatRGBAh int
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/RGBX16
	KCIFormatRGBX16 int
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/RGBX8
	KCIFormatRGBX8 int
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/rgbXf
	KCIFormatRGBXf int
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/rgbXh
	KCIFormatRGBXh int
	// KCIFormatRGf is a 64-bit-per-pixel, floating-point pixel format with only red and green color components.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/RGf
	KCIFormatRGf int
	// KCIFormatRGh is a 32-bit-per-pixel, floating-point pixel format with only red and green color components.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/RGh
	KCIFormatRGh int
	// KCIFormatRf is a 32-bit-per-pixel, floating-point pixel format in which the sole component is a red color value.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/Rf
	KCIFormatRf int
	// KCIFormatRh is a 16-bit-per-pixel, floating-point pixel format in which the sole component is a red color value.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFormat/Rh
	KCIFormatRh int
)

var (
	// KCIImageApplyCleanAperture is a Boolean value to control whether an image created with a CVPixelBuffer or an IOSurface should be cropped and offset according clean aperture attachments.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/applyCleanAperture
	KCIImageApplyCleanAperture CIImageOption
	// KCIImageApplyOrientationProperty is the key for transforming an image according to orientation metadata.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/applyOrientationProperty
	KCIImageApplyOrientationProperty CIImageOption
	// KCIImageAuxiliaryDepth is the key into the properties dictionary indicating whether to return an auxiliary depth image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliaryDepth
	KCIImageAuxiliaryDepth CIImageOption
	// KCIImageAuxiliaryDisparity is the key into the properties dictionary indicating whether to return an auxiliary disparity image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliaryDisparity
	KCIImageAuxiliaryDisparity CIImageOption
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliaryHDRGainMap
	KCIImageAuxiliaryHDRGainMap CIImageOption
	// KCIImageAuxiliaryPortraitEffectsMatte is the key into the properties dictionary indicating whether to return auxiliary portrait effects matte.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliaryPortraitEffectsMatte
	KCIImageAuxiliaryPortraitEffectsMatte CIImageOption
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliarySemanticSegmentationGlassesMatte
	KCIImageAuxiliarySemanticSegmentationGlassesMatte CIImageOption
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliarySemanticSegmentationHairMatte
	KCIImageAuxiliarySemanticSegmentationHairMatte CIImageOption
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliarySemanticSegmentationSkinMatte
	KCIImageAuxiliarySemanticSegmentationSkinMatte CIImageOption
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliarySemanticSegmentationSkyMatte
	KCIImageAuxiliarySemanticSegmentationSkyMatte CIImageOption
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliarySemanticSegmentationTeethMatte
	KCIImageAuxiliarySemanticSegmentationTeethMatte CIImageOption
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/cacheImmediately
	KCIImageCacheImmediately CIImageOption
	// KCIImageColorSpace is the key for a color space.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/colorSpace
	KCIImageColorSpace CIImageOption
	// KCIImageContentAverageLightLevel is a value for overriding the automatic behavior of the Content Average Light Level property when creating an image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/contentAverageLightLevel
	KCIImageContentAverageLightLevel CIImageOption
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/contentHeadroom
	KCIImageContentHeadroom CIImageOption
	// KCIImageExpandToHDR is a Boolean value that indicates whether to read Gain Map HDR images as HDR.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/expandToHDR
	KCIImageExpandToHDR CIImageOption
	// KCIImageNearestSampling is the key into the properties dictionary to indicate whether to use nearest-neighbor sampling.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/nearestSampling
	KCIImageNearestSampling CIImageOption
	// KCIImageProperties is the key for image metadata properties.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/properties
	KCIImageProperties CIImageOption
	// KCIImageProviderTileSize is a key for the image tiles size. The associated value is an [NSArray] that contains[NSNumber] objects for the dimensions of the image tiles requested from the image provider.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/providerTileSize
	KCIImageProviderTileSize CIImageOption
	// KCIImageProviderUserInfo is a key for data needed by the image provider. The associated value is an object that contains the needed data.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/providerUserInfo
	KCIImageProviderUserInfo CIImageOption
	// See: https://developer.apple.com/documentation/CoreImage/CIImageOption/toneMapHDRtoSDR
	KCIImageToneMapHDRtoSDR CIImageOption
)

var (
	// KCIImageRepresentationAVDepthData is the depth data representation of an image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption/avDepthData
	KCIImageRepresentationAVDepthData CIImageRepresentationOption
	// See: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption/avPortraitEffectsMatte
	KCIImageRepresentationAVPortraitEffectsMatte CIImageRepresentationOption
	// See: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption/avSemanticSegmentationMattes
	KCIImageRepresentationAVSemanticSegmentationMattes CIImageRepresentationOption
	// KCIImageRepresentationDepthImage is `options` dictionary key for image export methods to output depth data.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption/depthImage
	KCIImageRepresentationDepthImage CIImageRepresentationOption
	// KCIImageRepresentationDisparityImage is `options` dictionary key for image export methods to output disparity data.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption/disparityImage
	KCIImageRepresentationDisparityImage CIImageRepresentationOption
	// KCIImageRepresentationHDRGainMapAsRGB is an optional key and value to request the gain map channel to be color instead of monochrome.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption/hdrGainMapAsRGB
	KCIImageRepresentationHDRGainMapAsRGB CIImageRepresentationOption
	// KCIImageRepresentationHDRGainMapImage is an optional key and value to save a gain map channel to a JPEG or HEIF.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption/hdrGainMapImage
	KCIImageRepresentationHDRGainMapImage CIImageRepresentationOption
	// See: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption/hdrImage
	KCIImageRepresentationHDRImage CIImageRepresentationOption
	// See: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption/portraitEffectsMatteImage
	KCIImageRepresentationPortraitEffectsMatteImage CIImageRepresentationOption
	// See: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption/semanticSegmentationGlassesMatteImage
	KCIImageRepresentationSemanticSegmentationGlassesMatteImage CIImageRepresentationOption
	// See: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption/semanticSegmentationHairMatteImage
	KCIImageRepresentationSemanticSegmentationHairMatteImage CIImageRepresentationOption
	// See: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption/semanticSegmentationSkinMatteImage
	KCIImageRepresentationSemanticSegmentationSkinMatteImage CIImageRepresentationOption
	// See: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption/semanticSegmentationSkyMatteImage
	KCIImageRepresentationSemanticSegmentationSkyMatteImage CIImageRepresentationOption
	// See: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption/semanticSegmentationTeethMatteImage
	KCIImageRepresentationSemanticSegmentationTeethMatteImage CIImageRepresentationOption
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIDetectorAccuracy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIDetectorAccuracy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIDetectorAccuracyHigh"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIDetectorAccuracyHigh = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIDetectorAccuracyLow"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIDetectorAccuracyLow = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIDetectorAspectRatio"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIDetectorAspectRatio = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIDetectorEyeBlink"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIDetectorEyeBlink = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIDetectorFocalLength"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIDetectorFocalLength = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIDetectorImageOrientation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIDetectorImageOrientation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIDetectorMaxFeatureCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIDetectorMaxFeatureCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIDetectorMinFeatureSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIDetectorMinFeatureSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIDetectorNumberOfAngles"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIDetectorNumberOfAngles = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIDetectorReturnSubFeatures"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIDetectorReturnSubFeatures = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIDetectorSmile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIDetectorSmile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIDetectorTracking"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIDetectorTracking = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIDetectorTypeFace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIDetectorTypeFace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIDetectorTypeQRCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIDetectorTypeQRCode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIDetectorTypeRectangle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIDetectorTypeRectangle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIDetectorTypeText"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIDetectorTypeText = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIFeatureTypeFace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIFeatureTypeFace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIFeatureTypeQRCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIFeatureTypeQRCode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIFeatureTypeRectangle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIFeatureTypeRectangle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIFeatureTypeText"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIFeatureTypeText = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIRAWDecoderVersion6"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIRAWDecoderVersion6 = CIRAWDecoderVersion(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIRAWDecoderVersion6DNG"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIRAWDecoderVersion6DNG = CIRAWDecoderVersion(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIRAWDecoderVersion7"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIRAWDecoderVersion7 = CIRAWDecoderVersion(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIRAWDecoderVersion7DNG"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIRAWDecoderVersion7DNG = CIRAWDecoderVersion(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIRAWDecoderVersion8"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIRAWDecoderVersion8 = CIRAWDecoderVersion(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIRAWDecoderVersion8DNG"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIRAWDecoderVersion8DNG = CIRAWDecoderVersion(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIRAWDecoderVersion9"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIRAWDecoderVersion9 = CIRAWDecoderVersion(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIRAWDecoderVersion9DNG"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIRAWDecoderVersion9DNG = CIRAWDecoderVersion(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CIRAWDecoderVersionNone"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CIRAWDecoderVersionNone = CIRAWDecoderVersion(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIApplyOptionColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIApplyOptionColorSpace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIApplyOptionDefinition"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIApplyOptionDefinition = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIApplyOptionExtent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIApplyOptionExtent = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIApplyOptionUserInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIApplyOptionUserInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeClass = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeDefault"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeDefault = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeDisplayName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeDisplayName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeFilterAvailable_Mac"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeFilterAvailable_Mac = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeFilterAvailable_iOS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeFilterAvailable_iOS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeFilterCategories"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeFilterCategories = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeFilterDisplayName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeFilterDisplayName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeFilterName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeFilterName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeIdentity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeIdentity = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeMax"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeMax = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeMin"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeMin = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeReferenceDocumentation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeReferenceDocumentation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeSliderMax"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeSliderMax = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeSliderMin"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeSliderMin = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeTypeAngle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeTypeAngle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeTypeBoolean"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeTypeBoolean = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeTypeColor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeTypeColor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeTypeCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeTypeCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeTypeDistance"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeTypeDistance = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeTypeGradient"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeTypeGradient = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeTypeImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeTypeImage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeTypeInteger"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeTypeInteger = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeTypeOffset"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeTypeOffset = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeTypeOpaqueColor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeTypeOpaqueColor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeTypePosition"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeTypePosition = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeTypePosition3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeTypePosition3 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeTypeRectangle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeTypeRectangle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeTypeScalar"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeTypeScalar = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeTypeTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeTypeTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIAttributeTypeTransform"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIAttributeTypeTransform = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategoryBlur"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategoryBlur = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategoryBuiltIn"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategoryBuiltIn = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategoryColorAdjustment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategoryColorAdjustment = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategoryColorEffect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategoryColorEffect = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategoryCompositeOperation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategoryCompositeOperation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategoryDistortionEffect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategoryDistortionEffect = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategoryFilterGenerator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategoryFilterGenerator = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategoryGenerator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategoryGenerator = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategoryGeometryAdjustment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategoryGeometryAdjustment = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategoryGradient"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategoryGradient = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategoryHalftoneEffect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategoryHalftoneEffect = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategoryHighDynamicRange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategoryHighDynamicRange = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategoryInterlaced"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategoryInterlaced = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategoryNonSquarePixels"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategoryNonSquarePixels = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategoryReduction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategoryReduction = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategorySharpen"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategorySharpen = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategoryStillImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategoryStillImage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategoryStylize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategoryStylize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategoryTileEffect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategoryTileEffect = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategoryTransition"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategoryTransition = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCICategoryVideo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCICategoryVideo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIContextAllowLowPower"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIContextAllowLowPower = CIContextOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIContextCVMetalTextureCache"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIContextCVMetalTextureCache = CIContextOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIContextCacheIntermediates"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIContextCacheIntermediates = CIContextOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIContextHighQualityDownsample"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIContextHighQualityDownsample = CIContextOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIContextMemoryLimit"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIContextMemoryLimit = CIContextOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIContextName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIContextName = CIContextOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIContextOutputColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIContextOutputColorSpace = CIContextOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIContextOutputPremultiplied"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIContextOutputPremultiplied = CIContextOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIContextPriorityRequestLow"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIContextPriorityRequestLow = CIContextOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIContextUseSoftwareRenderer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIContextUseSoftwareRenderer = CIContextOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIContextWorkingColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIContextWorkingColorSpace = CIContextOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIContextWorkingFormat"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIContextWorkingFormat = CIContextOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIDynamicRangeConstrainedHigh"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIDynamicRangeConstrainedHigh = CIDynamicRangeOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIDynamicRangeHigh"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIDynamicRangeHigh = CIDynamicRangeOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIDynamicRangeStandard"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIDynamicRangeStandard = CIDynamicRangeOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFilterGeneratorExportedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIFilterGeneratorExportedKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFilterGeneratorExportedKeyName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIFilterGeneratorExportedKeyName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFilterGeneratorExportedKeyTargetObject"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIFilterGeneratorExportedKeyTargetObject = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatA16"); err == nil && ptr != 0 {
		KCIFormatA16 = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatA8"); err == nil && ptr != 0 {
		KCIFormatA8 = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatABGR8"); err == nil && ptr != 0 {
		KCIFormatABGR8 = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatARGB8"); err == nil && ptr != 0 {
		KCIFormatARGB8 = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatAf"); err == nil && ptr != 0 {
		KCIFormatAf = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatAh"); err == nil && ptr != 0 {
		KCIFormatAh = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatBGRA8"); err == nil && ptr != 0 {
		KCIFormatBGRA8 = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatL16"); err == nil && ptr != 0 {
		KCIFormatL16 = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatL8"); err == nil && ptr != 0 {
		KCIFormatL8 = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatLA16"); err == nil && ptr != 0 {
		KCIFormatLA16 = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatLA8"); err == nil && ptr != 0 {
		KCIFormatLA8 = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatLAf"); err == nil && ptr != 0 {
		KCIFormatLAf = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatLAh"); err == nil && ptr != 0 {
		KCIFormatLAh = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatLf"); err == nil && ptr != 0 {
		KCIFormatLf = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatLh"); err == nil && ptr != 0 {
		KCIFormatLh = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatR16"); err == nil && ptr != 0 {
		KCIFormatR16 = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatR8"); err == nil && ptr != 0 {
		KCIFormatR8 = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatRG16"); err == nil && ptr != 0 {
		KCIFormatRG16 = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatRG8"); err == nil && ptr != 0 {
		KCIFormatRG8 = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatRGB10"); err == nil && ptr != 0 {
		KCIFormatRGB10 = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatRGBA16"); err == nil && ptr != 0 {
		KCIFormatRGBA16 = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatRGBA8"); err == nil && ptr != 0 {
		KCIFormatRGBA8 = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatRGBAf"); err == nil && ptr != 0 {
		KCIFormatRGBAf = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatRGBAh"); err == nil && ptr != 0 {
		KCIFormatRGBAh = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatRGBX16"); err == nil && ptr != 0 {
		KCIFormatRGBX16 = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatRGBX8"); err == nil && ptr != 0 {
		KCIFormatRGBX8 = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatRGBXf"); err == nil && ptr != 0 {
		KCIFormatRGBXf = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatRGBXh"); err == nil && ptr != 0 {
		KCIFormatRGBXh = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatRGf"); err == nil && ptr != 0 {
		KCIFormatRGf = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatRGh"); err == nil && ptr != 0 {
		KCIFormatRGh = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatRf"); err == nil && ptr != 0 {
		KCIFormatRf = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIFormatRh"); err == nil && ptr != 0 {
		KCIFormatRh = *(*int)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageApplyCleanAperture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageApplyCleanAperture = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageApplyOrientationProperty"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageApplyOrientationProperty = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageAutoAdjustCrop"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageAutoAdjustCrop = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageAutoAdjustEnhance"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageAutoAdjustEnhance = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageAutoAdjustFeatures"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageAutoAdjustFeatures = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageAutoAdjustLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageAutoAdjustLevel = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageAutoAdjustRedEye"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageAutoAdjustRedEye = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageAuxiliaryDepth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageAuxiliaryDepth = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageAuxiliaryDisparity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageAuxiliaryDisparity = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageAuxiliaryHDRGainMap"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageAuxiliaryHDRGainMap = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageAuxiliaryPortraitEffectsMatte"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageAuxiliaryPortraitEffectsMatte = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageAuxiliarySemanticSegmentationGlassesMatte"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageAuxiliarySemanticSegmentationGlassesMatte = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageAuxiliarySemanticSegmentationHairMatte"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageAuxiliarySemanticSegmentationHairMatte = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageAuxiliarySemanticSegmentationSkinMatte"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageAuxiliarySemanticSegmentationSkinMatte = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageAuxiliarySemanticSegmentationSkyMatte"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageAuxiliarySemanticSegmentationSkyMatte = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageAuxiliarySemanticSegmentationTeethMatte"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageAuxiliarySemanticSegmentationTeethMatte = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageCacheImmediately"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageCacheImmediately = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageColorSpace = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageContentAverageLightLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageContentAverageLightLevel = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageContentHeadroom"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageContentHeadroom = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageExpandToHDR"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageExpandToHDR = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageNearestSampling"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageNearestSampling = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageProperties"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageProperties = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageProviderTileSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageProviderTileSize = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageProviderUserInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageProviderUserInfo = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageRepresentationAVDepthData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageRepresentationAVDepthData = CIImageRepresentationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageRepresentationAVPortraitEffectsMatte"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageRepresentationAVPortraitEffectsMatte = CIImageRepresentationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageRepresentationAVSemanticSegmentationMattes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageRepresentationAVSemanticSegmentationMattes = CIImageRepresentationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageRepresentationDepthImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageRepresentationDepthImage = CIImageRepresentationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageRepresentationDisparityImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageRepresentationDisparityImage = CIImageRepresentationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageRepresentationHDRGainMapAsRGB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageRepresentationHDRGainMapAsRGB = CIImageRepresentationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageRepresentationHDRGainMapImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageRepresentationHDRGainMapImage = CIImageRepresentationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageRepresentationHDRImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageRepresentationHDRImage = CIImageRepresentationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageRepresentationPortraitEffectsMatteImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageRepresentationPortraitEffectsMatteImage = CIImageRepresentationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageRepresentationSemanticSegmentationGlassesMatteImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageRepresentationSemanticSegmentationGlassesMatteImage = CIImageRepresentationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageRepresentationSemanticSegmentationHairMatteImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageRepresentationSemanticSegmentationHairMatteImage = CIImageRepresentationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageRepresentationSemanticSegmentationSkinMatteImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageRepresentationSemanticSegmentationSkinMatteImage = CIImageRepresentationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageRepresentationSemanticSegmentationSkyMatteImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageRepresentationSemanticSegmentationSkyMatteImage = CIImageRepresentationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageRepresentationSemanticSegmentationTeethMatteImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageRepresentationSemanticSegmentationTeethMatteImage = CIImageRepresentationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIImageToneMapHDRtoSDR"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIImageToneMapHDRtoSDR = CIImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputAmountKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputAmountKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputAngleKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputAngleKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputAspectRatioKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputAspectRatioKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputBackgroundImageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputBackgroundImageKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputBacksideImageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputBacksideImageKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputBiasKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputBiasKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputBiasVectorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputBiasVectorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputBrightnessKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputBrightnessKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputCenterKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputCenterKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputColor0Key"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputColor0Key = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputColor1Key"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputColor1Key = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputColorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputColorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputColorSpaceKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputColorSpaceKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputContrastKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputContrastKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputCountKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputCountKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputDepthImageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputDepthImageKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputDisparityImageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputDisparityImageKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputEVKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputEVKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputExtentKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputExtentKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputExtrapolateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputExtrapolateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputGradientImageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputGradientImageKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputImageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputImageKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputIntensityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputIntensityKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputMaskImageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputMaskImageKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputMatteImageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputMatteImageKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputPaletteImageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputPaletteImageKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputPerceptualKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputPerceptualKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputPoint0Key"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputPoint0Key = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputPoint1Key"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputPoint1Key = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputRadius0Key"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputRadius0Key = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputRadius1Key"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputRadius1Key = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputRadiusKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputRadiusKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputRefractionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputRefractionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputSaturationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputSaturationKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputScaleKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputScaleKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputShadingImageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputShadingImageKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputSharpnessKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputSharpnessKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputTargetImageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputTargetImageKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputThresholdKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputThresholdKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputTimeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputTimeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputTransformKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputTransformKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputVersionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputVersionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputWeightsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputWeightsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIInputWidthKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIInputWidthKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIOutputImageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIOutputImageKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCISamplerAffineMatrix"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCISamplerAffineMatrix = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCISamplerColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCISamplerColorSpace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCISamplerFilterLinear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCISamplerFilterLinear = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCISamplerFilterMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCISamplerFilterMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCISamplerFilterNearest"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCISamplerFilterNearest = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCISamplerWrapBlack"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCISamplerWrapBlack = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCISamplerWrapClamp"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCISamplerWrapClamp = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCISamplerWrapMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCISamplerWrapMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIUIParameterSet"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIUIParameterSet = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIUISetAdvanced"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIUISetAdvanced = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIUISetBasic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIUISetBasic = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIUISetDevelopment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIUISetDevelopment = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCIUISetIntermediate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCIUISetIntermediate = objc.GoString(cstr)
			}
		}
	}

}
