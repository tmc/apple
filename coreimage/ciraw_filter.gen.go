// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CIRAWFilter] class.
var (
	_CIRAWFilterClass     CIRAWFilterClass
	_CIRAWFilterClassOnce sync.Once
)

func getCIRAWFilterClass() CIRAWFilterClass {
	_CIRAWFilterClassOnce.Do(func() {
		_CIRAWFilterClass = CIRAWFilterClass{class: objc.GetClass("CIRAWFilter")}
	})
	return _CIRAWFilterClass
}

// GetCIRAWFilterClass returns the class object for CIRAWFilter.
func GetCIRAWFilterClass() CIRAWFilterClass {
	return getCIRAWFilterClass()
}

type CIRAWFilterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CIRAWFilterClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIRAWFilterClass) Alloc() CIRAWFilter {
	rv := objc.Send[CIRAWFilter](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A filter subclass that produces an image by manipulating RAW image sensor
// data from a digital camera or scanner.
//
// # Overview
// 
// Use this class to generate a [CIImage] object based on the configuration
// parameters you provide.
// 
// You can use this object in conjunction with other Core Image classes—such
// as [CIFilter] and [CIContext]—to take advantage of the built-in Core
// Image filters when processing images or writing custom filters.
// 
// You can also query this object to find out about the supported camera
// models, decoders, and filters.
//
// # Inspecting supported camera models, decoders, and filters
//
//   - [CIRAWFilter.SupportedDecoderVersions]: An array of all supported decoder versions for the given image type.
//   - [CIRAWFilter.ColorNoiseReductionSupported]: A Boolean that indicates if the current image supports color noise reduction adjustments.
//   - [CIRAWFilter.ContrastSupported]: A Boolean that indicates if the current image supports contrast adjustments.
//   - [CIRAWFilter.DetailSupported]: A Boolean that indicates if the current image supports detail enhancement adjustments.
//   - [CIRAWFilter.LensCorrectionSupported]: A Boolean that indicates if you can enable lens correction for the current image.
//   - [CIRAWFilter.LocalToneMapSupported]: A Boolean that indicates if the current image supports local tone curve adjustments.
//   - [CIRAWFilter.LuminanceNoiseReductionSupported]: A Boolean that indicates if the current image supports luminance noise reduction adjustments.
//   - [CIRAWFilter.MoireReductionSupported]: A Boolean that indicates if the current image supports moire artifact reduction adjustments.
//   - [CIRAWFilter.SharpnessSupported]: A Boolean that indicates if the current image supports sharpness adjustments.
//   - [CIRAWFilter.NativeSize]: The full native size of the unscaled image.
//
// # Configuring a filter
//
//   - [CIRAWFilter.BaselineExposure]: A value that indicates the baseline exposure to apply to the image.
//   - [CIRAWFilter.SetBaselineExposure]
//   - [CIRAWFilter.BoostAmount]: A value that indicates the amount of global tone curve to apply to the image.
//   - [CIRAWFilter.SetBoostAmount]
//   - [CIRAWFilter.BoostShadowAmount]: A value that indicates the amount to boost the shadow areas of the image.
//   - [CIRAWFilter.SetBoostShadowAmount]
//   - [CIRAWFilter.ColorNoiseReductionAmount]: A value that indicates the amount of chroma noise reduction to apply to the image.
//   - [CIRAWFilter.SetColorNoiseReductionAmount]
//   - [CIRAWFilter.ContrastAmount]: A value that indicates the amount of local contrast to apply to the edges of the image.
//   - [CIRAWFilter.SetContrastAmount]
//   - [CIRAWFilter.DecoderVersion]: A value that indicates the decoder version to use.
//   - [CIRAWFilter.SetDecoderVersion]
//   - [CIRAWFilter.DetailAmount]: A value that indicates the amount of detail enhancement to apply to the edges of the image.
//   - [CIRAWFilter.SetDetailAmount]
//   - [CIRAWFilter.Exposure]: A value that indicates the amount of exposure to apply to the image.
//   - [CIRAWFilter.SetExposure]
//   - [CIRAWFilter.ExtendedDynamicRangeAmount]: A value that indicates the amount of extended dynamic range (EDR) to apply to the image.
//   - [CIRAWFilter.SetExtendedDynamicRangeAmount]
//   - [CIRAWFilter.DraftModeEnabled]: A Boolean that indicates whether to enable draft mode.
//   - [CIRAWFilter.SetDraftModeEnabled]
//   - [CIRAWFilter.GamutMappingEnabled]: A Boolean that indicates whether to enable gamut mapping.
//   - [CIRAWFilter.SetGamutMappingEnabled]
//   - [CIRAWFilter.LensCorrectionEnabled]: A Boolean that indicates whether to enable lens correction.
//   - [CIRAWFilter.SetLensCorrectionEnabled]
//   - [CIRAWFilter.LinearSpaceFilter]: An optional filter you can apply to the RAW image while it’s in linear space.
//   - [CIRAWFilter.SetLinearSpaceFilter]
//   - [CIRAWFilter.LocalToneMapAmount]: A value that indicates the amount of local tone curve to apply to the image.
//   - [CIRAWFilter.SetLocalToneMapAmount]
//   - [CIRAWFilter.LuminanceNoiseReductionAmount]: A value that indicates the amount of luminance noise reduction to apply to the image.
//   - [CIRAWFilter.SetLuminanceNoiseReductionAmount]
//   - [CIRAWFilter.MoireReductionAmount]: A value that indicates the amount of moire artifact reduction to apply to high frequency areas of the image.
//   - [CIRAWFilter.SetMoireReductionAmount]
//   - [CIRAWFilter.NeutralChromaticity]: A value that indicates the amount of white balance based on chromaticity values to apply to the image.
//   - [CIRAWFilter.SetNeutralChromaticity]
//   - [CIRAWFilter.NeutralLocation]: A value that indicates the amount of white balance based on pixel coordinates to apply to the image.
//   - [CIRAWFilter.SetNeutralLocation]
//   - [CIRAWFilter.NeutralTemperature]: A value that indicates the amount of white balance based on temperature values to apply to the image.
//   - [CIRAWFilter.SetNeutralTemperature]
//   - [CIRAWFilter.NeutralTint]: A value that indicates the amount of white balance based on tint values to apply to the image.
//   - [CIRAWFilter.SetNeutralTint]
//   - [CIRAWFilter.Orientation]: A value that indicates the orientation of the image.
//   - [CIRAWFilter.SetOrientation]
//   - [CIRAWFilter.PortraitEffectsMatte]: An optional auxiliary image that represents the portrait effects matte of the image.
//   - [CIRAWFilter.PreviewImage]: An optional auxiliary image that represents a preview of the original image.
//   - [CIRAWFilter.Properties]: A dictionary that contains properties of the image source.
//   - [CIRAWFilter.ScaleFactor]: A value that indicates the desired scale factor to draw the output image.
//   - [CIRAWFilter.SetScaleFactor]
//   - [CIRAWFilter.SemanticSegmentationGlassesMatte]: An optional auxiliary image that represents the semantic segmentation glasses matte of the image.
//   - [CIRAWFilter.SemanticSegmentationHairMatte]: An optional auxiliary image that represents the semantic segmentation hair matte of the image.
//   - [CIRAWFilter.SemanticSegmentationSkinMatte]: An optional auxiliary image that represents the semantic segmentation skin matte of the image.
//   - [CIRAWFilter.SemanticSegmentationSkyMatte]: An optional auxiliary image that represents the semantic segmentation sky matte of the image.
//   - [CIRAWFilter.SemanticSegmentationTeethMatte]: An optional auxiliary image that represents the semantic segmentation teeth matte of the image.
//   - [CIRAWFilter.ShadowBias]: A value that indicates the amount to subtract from the shadows in the image.
//   - [CIRAWFilter.SetShadowBias]
//   - [CIRAWFilter.SharpnessAmount]: A value that indicates the amount of sharpness to apply to the edges of the image.
//   - [CIRAWFilter.SetSharpnessAmount]
//
// # Instance Properties
//
//   - [CIRAWFilter.HighlightRecoveryEnabled]
//   - [CIRAWFilter.SetHighlightRecoveryEnabled]
//   - [CIRAWFilter.HighlightRecoverySupported]
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter
type CIRAWFilter struct {
	CIFilter
}

// CIRAWFilterFromID constructs a [CIRAWFilter] from an objc.ID.
//
// A filter subclass that produces an image by manipulating RAW image sensor
// data from a digital camera or scanner.
func CIRAWFilterFromID(id objc.ID) CIRAWFilter {
	return CIRAWFilter{CIFilter: CIFilterFromID(id)}
}
// NOTE: CIRAWFilter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIRAWFilter] class.
//
// # Inspecting supported camera models, decoders, and filters
//
//   - [ICIRAWFilter.SupportedDecoderVersions]: An array of all supported decoder versions for the given image type.
//   - [ICIRAWFilter.ColorNoiseReductionSupported]: A Boolean that indicates if the current image supports color noise reduction adjustments.
//   - [ICIRAWFilter.ContrastSupported]: A Boolean that indicates if the current image supports contrast adjustments.
//   - [ICIRAWFilter.DetailSupported]: A Boolean that indicates if the current image supports detail enhancement adjustments.
//   - [ICIRAWFilter.LensCorrectionSupported]: A Boolean that indicates if you can enable lens correction for the current image.
//   - [ICIRAWFilter.LocalToneMapSupported]: A Boolean that indicates if the current image supports local tone curve adjustments.
//   - [ICIRAWFilter.LuminanceNoiseReductionSupported]: A Boolean that indicates if the current image supports luminance noise reduction adjustments.
//   - [ICIRAWFilter.MoireReductionSupported]: A Boolean that indicates if the current image supports moire artifact reduction adjustments.
//   - [ICIRAWFilter.SharpnessSupported]: A Boolean that indicates if the current image supports sharpness adjustments.
//   - [ICIRAWFilter.NativeSize]: The full native size of the unscaled image.
//
// # Configuring a filter
//
//   - [ICIRAWFilter.BaselineExposure]: A value that indicates the baseline exposure to apply to the image.
//   - [ICIRAWFilter.SetBaselineExposure]
//   - [ICIRAWFilter.BoostAmount]: A value that indicates the amount of global tone curve to apply to the image.
//   - [ICIRAWFilter.SetBoostAmount]
//   - [ICIRAWFilter.BoostShadowAmount]: A value that indicates the amount to boost the shadow areas of the image.
//   - [ICIRAWFilter.SetBoostShadowAmount]
//   - [ICIRAWFilter.ColorNoiseReductionAmount]: A value that indicates the amount of chroma noise reduction to apply to the image.
//   - [ICIRAWFilter.SetColorNoiseReductionAmount]
//   - [ICIRAWFilter.ContrastAmount]: A value that indicates the amount of local contrast to apply to the edges of the image.
//   - [ICIRAWFilter.SetContrastAmount]
//   - [ICIRAWFilter.DecoderVersion]: A value that indicates the decoder version to use.
//   - [ICIRAWFilter.SetDecoderVersion]
//   - [ICIRAWFilter.DetailAmount]: A value that indicates the amount of detail enhancement to apply to the edges of the image.
//   - [ICIRAWFilter.SetDetailAmount]
//   - [ICIRAWFilter.Exposure]: A value that indicates the amount of exposure to apply to the image.
//   - [ICIRAWFilter.SetExposure]
//   - [ICIRAWFilter.ExtendedDynamicRangeAmount]: A value that indicates the amount of extended dynamic range (EDR) to apply to the image.
//   - [ICIRAWFilter.SetExtendedDynamicRangeAmount]
//   - [ICIRAWFilter.DraftModeEnabled]: A Boolean that indicates whether to enable draft mode.
//   - [ICIRAWFilter.SetDraftModeEnabled]
//   - [ICIRAWFilter.GamutMappingEnabled]: A Boolean that indicates whether to enable gamut mapping.
//   - [ICIRAWFilter.SetGamutMappingEnabled]
//   - [ICIRAWFilter.LensCorrectionEnabled]: A Boolean that indicates whether to enable lens correction.
//   - [ICIRAWFilter.SetLensCorrectionEnabled]
//   - [ICIRAWFilter.LinearSpaceFilter]: An optional filter you can apply to the RAW image while it’s in linear space.
//   - [ICIRAWFilter.SetLinearSpaceFilter]
//   - [ICIRAWFilter.LocalToneMapAmount]: A value that indicates the amount of local tone curve to apply to the image.
//   - [ICIRAWFilter.SetLocalToneMapAmount]
//   - [ICIRAWFilter.LuminanceNoiseReductionAmount]: A value that indicates the amount of luminance noise reduction to apply to the image.
//   - [ICIRAWFilter.SetLuminanceNoiseReductionAmount]
//   - [ICIRAWFilter.MoireReductionAmount]: A value that indicates the amount of moire artifact reduction to apply to high frequency areas of the image.
//   - [ICIRAWFilter.SetMoireReductionAmount]
//   - [ICIRAWFilter.NeutralChromaticity]: A value that indicates the amount of white balance based on chromaticity values to apply to the image.
//   - [ICIRAWFilter.SetNeutralChromaticity]
//   - [ICIRAWFilter.NeutralLocation]: A value that indicates the amount of white balance based on pixel coordinates to apply to the image.
//   - [ICIRAWFilter.SetNeutralLocation]
//   - [ICIRAWFilter.NeutralTemperature]: A value that indicates the amount of white balance based on temperature values to apply to the image.
//   - [ICIRAWFilter.SetNeutralTemperature]
//   - [ICIRAWFilter.NeutralTint]: A value that indicates the amount of white balance based on tint values to apply to the image.
//   - [ICIRAWFilter.SetNeutralTint]
//   - [ICIRAWFilter.Orientation]: A value that indicates the orientation of the image.
//   - [ICIRAWFilter.SetOrientation]
//   - [ICIRAWFilter.PortraitEffectsMatte]: An optional auxiliary image that represents the portrait effects matte of the image.
//   - [ICIRAWFilter.PreviewImage]: An optional auxiliary image that represents a preview of the original image.
//   - [ICIRAWFilter.Properties]: A dictionary that contains properties of the image source.
//   - [ICIRAWFilter.ScaleFactor]: A value that indicates the desired scale factor to draw the output image.
//   - [ICIRAWFilter.SetScaleFactor]
//   - [ICIRAWFilter.SemanticSegmentationGlassesMatte]: An optional auxiliary image that represents the semantic segmentation glasses matte of the image.
//   - [ICIRAWFilter.SemanticSegmentationHairMatte]: An optional auxiliary image that represents the semantic segmentation hair matte of the image.
//   - [ICIRAWFilter.SemanticSegmentationSkinMatte]: An optional auxiliary image that represents the semantic segmentation skin matte of the image.
//   - [ICIRAWFilter.SemanticSegmentationSkyMatte]: An optional auxiliary image that represents the semantic segmentation sky matte of the image.
//   - [ICIRAWFilter.SemanticSegmentationTeethMatte]: An optional auxiliary image that represents the semantic segmentation teeth matte of the image.
//   - [ICIRAWFilter.ShadowBias]: A value that indicates the amount to subtract from the shadows in the image.
//   - [ICIRAWFilter.SetShadowBias]
//   - [ICIRAWFilter.SharpnessAmount]: A value that indicates the amount of sharpness to apply to the edges of the image.
//   - [ICIRAWFilter.SetSharpnessAmount]
//
// # Instance Properties
//
//   - [ICIRAWFilter.HighlightRecoveryEnabled]
//   - [ICIRAWFilter.SetHighlightRecoveryEnabled]
//   - [ICIRAWFilter.HighlightRecoverySupported]
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter
type ICIRAWFilter interface {
	ICIFilter

	// Topic: Inspecting supported camera models, decoders, and filters

	// An array of all supported decoder versions for the given image type.
	SupportedDecoderVersions() []string
	// A Boolean that indicates if the current image supports color noise reduction adjustments.
	ColorNoiseReductionSupported() bool
	// A Boolean that indicates if the current image supports contrast adjustments.
	ContrastSupported() bool
	// A Boolean that indicates if the current image supports detail enhancement adjustments.
	DetailSupported() bool
	// A Boolean that indicates if you can enable lens correction for the current image.
	LensCorrectionSupported() bool
	// A Boolean that indicates if the current image supports local tone curve adjustments.
	LocalToneMapSupported() bool
	// A Boolean that indicates if the current image supports luminance noise reduction adjustments.
	LuminanceNoiseReductionSupported() bool
	// A Boolean that indicates if the current image supports moire artifact reduction adjustments.
	MoireReductionSupported() bool
	// A Boolean that indicates if the current image supports sharpness adjustments.
	SharpnessSupported() bool
	// The full native size of the unscaled image.
	NativeSize() corefoundation.CGSize

	// Topic: Configuring a filter

	// A value that indicates the baseline exposure to apply to the image.
	BaselineExposure() float32
	SetBaselineExposure(value float32)
	// A value that indicates the amount of global tone curve to apply to the image.
	BoostAmount() float32
	SetBoostAmount(value float32)
	// A value that indicates the amount to boost the shadow areas of the image.
	BoostShadowAmount() float32
	SetBoostShadowAmount(value float32)
	// A value that indicates the amount of chroma noise reduction to apply to the image.
	ColorNoiseReductionAmount() float32
	SetColorNoiseReductionAmount(value float32)
	// A value that indicates the amount of local contrast to apply to the edges of the image.
	ContrastAmount() float32
	SetContrastAmount(value float32)
	// A value that indicates the decoder version to use.
	DecoderVersion() CIRAWDecoderVersion
	SetDecoderVersion(value CIRAWDecoderVersion)
	// A value that indicates the amount of detail enhancement to apply to the edges of the image.
	DetailAmount() float32
	SetDetailAmount(value float32)
	// A value that indicates the amount of exposure to apply to the image.
	Exposure() float32
	SetExposure(value float32)
	// A value that indicates the amount of extended dynamic range (EDR) to apply to the image.
	ExtendedDynamicRangeAmount() float32
	SetExtendedDynamicRangeAmount(value float32)
	// A Boolean that indicates whether to enable draft mode.
	DraftModeEnabled() bool
	SetDraftModeEnabled(value bool)
	// A Boolean that indicates whether to enable gamut mapping.
	GamutMappingEnabled() bool
	SetGamutMappingEnabled(value bool)
	// A Boolean that indicates whether to enable lens correction.
	LensCorrectionEnabled() bool
	SetLensCorrectionEnabled(value bool)
	// An optional filter you can apply to the RAW image while it’s in linear space.
	LinearSpaceFilter() CIFilter
	SetLinearSpaceFilter(value CIFilter)
	// A value that indicates the amount of local tone curve to apply to the image.
	LocalToneMapAmount() float32
	SetLocalToneMapAmount(value float32)
	// A value that indicates the amount of luminance noise reduction to apply to the image.
	LuminanceNoiseReductionAmount() float32
	SetLuminanceNoiseReductionAmount(value float32)
	// A value that indicates the amount of moire artifact reduction to apply to high frequency areas of the image.
	MoireReductionAmount() float32
	SetMoireReductionAmount(value float32)
	// A value that indicates the amount of white balance based on chromaticity values to apply to the image.
	NeutralChromaticity() corefoundation.CGPoint
	SetNeutralChromaticity(value corefoundation.CGPoint)
	// A value that indicates the amount of white balance based on pixel coordinates to apply to the image.
	NeutralLocation() corefoundation.CGPoint
	SetNeutralLocation(value corefoundation.CGPoint)
	// A value that indicates the amount of white balance based on temperature values to apply to the image.
	NeutralTemperature() float32
	SetNeutralTemperature(value float32)
	// A value that indicates the amount of white balance based on tint values to apply to the image.
	NeutralTint() float32
	SetNeutralTint(value float32)
	// A value that indicates the orientation of the image.
	Orientation() objectivec.IObject
	SetOrientation(value objectivec.IObject)
	// An optional auxiliary image that represents the portrait effects matte of the image.
	PortraitEffectsMatte() ICIImage
	// An optional auxiliary image that represents a preview of the original image.
	PreviewImage() ICIImage
	// A dictionary that contains properties of the image source.
	Properties() foundation.INSDictionary
	// A value that indicates the desired scale factor to draw the output image.
	ScaleFactor() float32
	SetScaleFactor(value float32)
	// An optional auxiliary image that represents the semantic segmentation glasses matte of the image.
	SemanticSegmentationGlassesMatte() ICIImage
	// An optional auxiliary image that represents the semantic segmentation hair matte of the image.
	SemanticSegmentationHairMatte() ICIImage
	// An optional auxiliary image that represents the semantic segmentation skin matte of the image.
	SemanticSegmentationSkinMatte() ICIImage
	// An optional auxiliary image that represents the semantic segmentation sky matte of the image.
	SemanticSegmentationSkyMatte() ICIImage
	// An optional auxiliary image that represents the semantic segmentation teeth matte of the image.
	SemanticSegmentationTeethMatte() ICIImage
	// A value that indicates the amount to subtract from the shadows in the image.
	ShadowBias() float32
	SetShadowBias(value float32)
	// A value that indicates the amount of sharpness to apply to the edges of the image.
	SharpnessAmount() float32
	SetSharpnessAmount(value float32)

	// Topic: Instance Properties

	HighlightRecoveryEnabled() bool
	SetHighlightRecoveryEnabled(value bool)
	HighlightRecoverySupported() bool
}

// Init initializes the instance.
func (r CIRAWFilter) Init() CIRAWFilter {
	rv := objc.Send[CIRAWFilter](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r CIRAWFilter) Autorelease() CIRAWFilter {
	rv := objc.Send[CIRAWFilter](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIRAWFilter creates a new CIRAWFilter instance.
func NewCIRAWFilter() CIRAWFilter {
	class := getCIRAWFilterClass()
	rv := objc.Send[CIRAWFilter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a RAW filter from the pixel buffer and its properties that you
// specify.
//
// buffer: A Core Video pixel buffer.
//
// properties: A dictionary that defines the properties of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(cvPixelBuffer:properties:)
func NewRAWFilterWithCVPixelBufferProperties(buffer corevideo.CVImageBufferRef, properties foundation.INSDictionary) CIRAWFilter {
	rv := objc.Send[objc.ID](objc.ID(getCIRAWFilterClass().class), objc.Sel("filterWithCVPixelBuffer:properties:"), buffer, properties)
	return CIRAWFilterFromID(rv)
}

// Creates a RAW filter from the image data and type hint that you specify.
//
// data: The image data.
//
// identifierHint: A string that identifies the image type.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(imageData:identifierHint:)
func NewRAWFilterWithImageDataIdentifierHint(data foundation.INSData, identifierHint string) CIRAWFilter {
	rv := objc.Send[objc.ID](objc.ID(getCIRAWFilterClass().class), objc.Sel("filterWithImageData:identifierHint:"), data, objc.String(identifierHint))
	return CIRAWFilterFromID(rv)
}

// Creates a RAW filter from the image at the URL location that you specify.
//
// url: The URL location of the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(imageURL:)
func NewRAWFilterWithImageURL(url foundation.INSURL) CIRAWFilter {
	rv := objc.Send[objc.ID](objc.ID(getCIRAWFilterClass().class), objc.Sel("filterWithImageURL:"), url)
	return CIRAWFilterFromID(rv)
}

// An array of all supported decoder versions for the given image type.
//
// # Discussion
// 
// This array is sorted in reverse chronological order. All entries represent
// a valid version identifier set to [DecoderVersion].
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/supportedDecoderVersions
func (r CIRAWFilter) SupportedDecoderVersions() []string {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("supportedDecoderVersions"))
	return objc.ConvertSliceToStrings(rv)
}
// A Boolean that indicates if the current image supports color noise
// reduction adjustments.
//
// # Discussion
// 
// If this value is [true], you can adjust the amount of color noise reduction
// to apply to the image by setting [ColorNoiseReductionAmount].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isColorNoiseReductionSupported
func (r CIRAWFilter) ColorNoiseReductionSupported() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isColorNoiseReductionSupported"))
	return rv
}
// A Boolean that indicates if the current image supports contrast
// adjustments.
//
// # Discussion
// 
// If this value is [true], you can adjust the amount of contrast to apply to
// the image by setting [ContrastAmount].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isContrastSupported
func (r CIRAWFilter) ContrastSupported() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isContrastSupported"))
	return rv
}
// A Boolean that indicates if the current image supports detail enhancement
// adjustments.
//
// # Discussion
// 
// If this value is [true], you can adjust the amount of detail enhancement to
// apply to the image by setting [DetailAmount].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isDetailSupported
func (r CIRAWFilter) DetailSupported() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isDetailSupported"))
	return rv
}
// A Boolean that indicates if you can enable lens correction for the current
// image.
//
// # Discussion
// 
// If this value is [true], you can enable lens correction on the image by
// setting [LensCorrectionEnabled] to [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isLensCorrectionSupported
func (r CIRAWFilter) LensCorrectionSupported() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isLensCorrectionSupported"))
	return rv
}
// A Boolean that indicates if the current image supports local tone curve
// adjustments.
//
// # Discussion
// 
// If this value is [true], you can adjust the amount of local tone curve to
// apply to the image by setting [LocalToneMapAmount].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isLocalToneMapSupported
func (r CIRAWFilter) LocalToneMapSupported() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isLocalToneMapSupported"))
	return rv
}
// A Boolean that indicates if the current image supports luminance noise
// reduction adjustments.
//
// # Discussion
// 
// If this value is [true], you can adjust the amount of luminance noise
// reduction to apply to the image by setting [LuminanceNoiseReductionAmount].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isLuminanceNoiseReductionSupported
func (r CIRAWFilter) LuminanceNoiseReductionSupported() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isLuminanceNoiseReductionSupported"))
	return rv
}
// A Boolean that indicates if the current image supports moire artifact
// reduction adjustments.
//
// # Discussion
// 
// If this value is [true], you can adjust the amount of moire artifact
// reduction to apply to the image by setting [MoireReductionAmount].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isMoireReductionSupported
func (r CIRAWFilter) MoireReductionSupported() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isMoireReductionSupported"))
	return rv
}
// A Boolean that indicates if the current image supports sharpness
// adjustments.
//
// # Discussion
// 
// If this value is [true], you can adjust the amount of sharpness to apply to
// the image by setting [SharpnessAmount].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isSharpnessSupported
func (r CIRAWFilter) SharpnessSupported() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isSharpnessSupported"))
	return rv
}
// The full native size of the unscaled image.
//
// # Discussion
// 
// This value isn’t affected by orientation changes.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/nativeSize
func (r CIRAWFilter) NativeSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](r.ID, objc.Sel("nativeSize"))
	return corefoundation.CGSize(rv)
}
// A value that indicates the baseline exposure to apply to the image.
//
// # Discussion
// 
// The default value varies with camera settings. A value of `0` indicates
// linear response.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/baselineExposure
func (r CIRAWFilter) BaselineExposure() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("baselineExposure"))
	return rv
}
func (r CIRAWFilter) SetBaselineExposure(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setBaselineExposure:"), value)
}
// A value that indicates the amount of global tone curve to apply to the
// image.
//
// # Discussion
// 
// The value should be in the range of `0...1`. A value of `0` indicates no
// global tone curve (linear response), and a value of `1` indicates full
// global tone curve. The default value is `1`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/boostAmount
func (r CIRAWFilter) BoostAmount() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("boostAmount"))
	return rv
}
func (r CIRAWFilter) SetBoostAmount(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setBoostAmount:"), value)
}
// A value that indicates the amount to boost the shadow areas of the image.
//
// # Discussion
// 
// Use this value to lighten details in shadows. The value should be in the
// range of `0...2`. The default value is `1`. A value less than `1` darkens
// the shadows, and a value greater than `1` lightens the shadows.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/boostShadowAmount
func (r CIRAWFilter) BoostShadowAmount() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("boostShadowAmount"))
	return rv
}
func (r CIRAWFilter) SetBoostShadowAmount(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setBoostShadowAmount:"), value)
}
// A value that indicates the amount of chroma noise reduction to apply to the
// image.
//
// # Discussion
// 
// The value should be in the range of `0...1`. The default value varies by
// image. A value of `0` indicates no chroma noise reduction, and a value of
// `1` indicates maximum chroma noise reduction.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/colorNoiseReductionAmount
func (r CIRAWFilter) ColorNoiseReductionAmount() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("colorNoiseReductionAmount"))
	return rv
}
func (r CIRAWFilter) SetColorNoiseReductionAmount(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setColorNoiseReductionAmount:"), value)
}
// A value that indicates the amount of local contrast to apply to the edges
// of the image.
//
// # Discussion
// 
// The value should be in the range of `0...1`. The default value varies by
// image. A value of `0` indicates no contrast, and a value of `1` indicates
// maximum contrast.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/contrastAmount
func (r CIRAWFilter) ContrastAmount() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("contrastAmount"))
	return rv
}
func (r CIRAWFilter) SetContrastAmount(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setContrastAmount:"), value)
}
// A value that indicates the decoder version to use.
//
// # Discussion
// 
// The value should be in the range of `0` to the current decoder version.
// 
// A newly initialized object defaults to the newest available decoder version
// for the given image type. You can request an older version to maintain
// compatibility with older releases. However, the version you request needs
// to be a member of [SupportedDecoderVersions], otherwise the system
// generates a `nil` output image.
// 
// When you request a specific version of the decoder, Core Image produces an
// image that looks the same across different versions. However, Core Image
// doesn’t guarantee that the bits are the same across versions, because the
// rounding behavior of floating-point arithmetic varies due to differences in
// compilers or hardware.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/decoderVersion
func (r CIRAWFilter) DecoderVersion() CIRAWDecoderVersion {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("decoderVersion"))
	return CIRAWDecoderVersion(foundation.NSStringFromID(rv).String())
}
func (r CIRAWFilter) SetDecoderVersion(value CIRAWDecoderVersion) {
	objc.Send[struct{}](r.ID, objc.Sel("setDecoderVersion:"), objc.String(string(value)))
}
// A value that indicates the amount of detail enhancement to apply to the
// edges of the image.
//
// # Discussion
// 
// The value should be in the range of `0...3`. The default value varies by
// image. A value of `0` indicates no detail enhancement, and a value of `3`
// indicates maximum detail enhancement.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/detailAmount
func (r CIRAWFilter) DetailAmount() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("detailAmount"))
	return rv
}
func (r CIRAWFilter) SetDetailAmount(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setDetailAmount:"), value)
}
// A value that indicates the amount of exposure to apply to the image.
//
// # Discussion
// 
// The default value is `0`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/exposure
func (r CIRAWFilter) Exposure() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("exposure"))
	return rv
}
func (r CIRAWFilter) SetExposure(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setExposure:"), value)
}
// A value that indicates the amount of extended dynamic range (EDR) to apply
// to the image.
//
// # Discussion
// 
// The value should be in the range of `0...2`. The default value is `0`. A
// value of `0` indicates no EDR. A value of `1` indicates default EDR, and a
// value of `2` indicates maximum EDR.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/extendedDynamicRangeAmount
func (r CIRAWFilter) ExtendedDynamicRangeAmount() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("extendedDynamicRangeAmount"))
	return rv
}
func (r CIRAWFilter) SetExtendedDynamicRangeAmount(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setExtendedDynamicRangeAmount:"), value)
}
// A Boolean that indicates whether to enable draft mode.
//
// # Discussion
// 
// Setting this value to [true] can improve image decoding speed with minimal
// loss of quality. The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isDraftModeEnabled
func (r CIRAWFilter) DraftModeEnabled() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isDraftModeEnabled"))
	return rv
}
func (r CIRAWFilter) SetDraftModeEnabled(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setDraftModeEnabled:"), value)
}
// A Boolean that indicates whether to enable gamut mapping.
//
// # Discussion
// 
// The default value is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isGamutMappingEnabled
func (r CIRAWFilter) GamutMappingEnabled() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isGamutMappingEnabled"))
	return rv
}
func (r CIRAWFilter) SetGamutMappingEnabled(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setGamutMappingEnabled:"), value)
}
// A Boolean that indicates whether to enable lens correction.
//
// # Discussion
// 
// The default value varies by image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isLensCorrectionEnabled
func (r CIRAWFilter) LensCorrectionEnabled() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isLensCorrectionEnabled"))
	return rv
}
func (r CIRAWFilter) SetLensCorrectionEnabled(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setLensCorrectionEnabled:"), value)
}
// An optional filter you can apply to the RAW image while it’s in linear
// space.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/linearSpaceFilter
func (r CIRAWFilter) LinearSpaceFilter() CIFilter {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("linearSpaceFilter"))
	return CIFilterFromID(objc.ID(rv))
}
func (r CIRAWFilter) SetLinearSpaceFilter(value CIFilter) {
	objc.Send[struct{}](r.ID, objc.Sel("setLinearSpaceFilter:"), value)
}
// A value that indicates the amount of local tone curve to apply to the
// image.
//
// # Discussion
// 
// The value should be in the range of `0...1`. The default value varies by
// image. A value of `0` indicates no local tone curve (linear response), and
// a value of `1` indicates full global tone curve.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/localToneMapAmount
func (r CIRAWFilter) LocalToneMapAmount() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("localToneMapAmount"))
	return rv
}
func (r CIRAWFilter) SetLocalToneMapAmount(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setLocalToneMapAmount:"), value)
}
// A value that indicates the amount of luminance noise reduction to apply to
// the image.
//
// # Discussion
// 
// The value should be in the range of `0...1`. The default value varies by
// image. A value of `0` indicates no luminance noise reduction, and a value
// of `1` indicates maximum luminance noise reduction.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/luminanceNoiseReductionAmount
func (r CIRAWFilter) LuminanceNoiseReductionAmount() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("luminanceNoiseReductionAmount"))
	return rv
}
func (r CIRAWFilter) SetLuminanceNoiseReductionAmount(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setLuminanceNoiseReductionAmount:"), value)
}
// A value that indicates the amount of moire artifact reduction to apply to
// high frequency areas of the image.
//
// # Discussion
// 
// The value should be in the range of `0...1`. The default value varies by
// image. A value of `0` indicates no moire reduction, and a value of `1`
// indicates maximum moire reduction.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/moireReductionAmount
func (r CIRAWFilter) MoireReductionAmount() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("moireReductionAmount"))
	return rv
}
func (r CIRAWFilter) SetMoireReductionAmount(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setMoireReductionAmount:"), value)
}
// A value that indicates the amount of white balance based on chromaticity
// values to apply to the image.
//
// # Discussion
// 
// Set this value to change the level of white balance to apply to the image,
// based on `x,y` chromaticity values in the range `0...1`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralChromaticity
func (r CIRAWFilter) NeutralChromaticity() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r.ID, objc.Sel("neutralChromaticity"))
	return corefoundation.CGPoint(rv)
}
func (r CIRAWFilter) SetNeutralChromaticity(value corefoundation.CGPoint) {
	objc.Send[struct{}](r.ID, objc.Sel("setNeutralChromaticity:"), value)
}
// A value that indicates the amount of white balance based on pixel
// coordinates to apply to the image.
//
// # Discussion
// 
// Set this value based on `x,y` pixel coordinates in the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralLocation
func (r CIRAWFilter) NeutralLocation() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r.ID, objc.Sel("neutralLocation"))
	return corefoundation.CGPoint(rv)
}
func (r CIRAWFilter) SetNeutralLocation(value corefoundation.CGPoint) {
	objc.Send[struct{}](r.ID, objc.Sel("setNeutralLocation:"), value)
}
// A value that indicates the amount of white balance based on temperature
// values to apply to the image.
//
// # Discussion
// 
// Set this value based on temperature values in the range `2000K...50000K`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralTemperature
func (r CIRAWFilter) NeutralTemperature() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("neutralTemperature"))
	return rv
}
func (r CIRAWFilter) SetNeutralTemperature(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setNeutralTemperature:"), value)
}
// A value that indicates the amount of white balance based on tint values to
// apply to the image.
//
// # Discussion
// 
// Set this value based on tint values in the range `-150...150`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralTint
func (r CIRAWFilter) NeutralTint() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("neutralTint"))
	return rv
}
func (r CIRAWFilter) SetNeutralTint(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setNeutralTint:"), value)
}
// A value that indicates the orientation of the image.
//
// # Discussion
// 
// The value should be in the range of `1...8` and follow the EXIF
// specification.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/orientation
func (r CIRAWFilter) Orientation() objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("orientation"))
	return objectivec.Object{ID: rv}
}
func (r CIRAWFilter) SetOrientation(value objectivec.IObject) {
	objc.Send[struct{}](r.ID, objc.Sel("setOrientation:"), value)
}
// An optional auxiliary image that represents the portrait effects matte of
// the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/portraitEffectsMatte
func (r CIRAWFilter) PortraitEffectsMatte() ICIImage {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("portraitEffectsMatte"))
	return CIImageFromID(objc.ID(rv))
}
// An optional auxiliary image that represents a preview of the original
// image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/previewImage
func (r CIRAWFilter) PreviewImage() ICIImage {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("previewImage"))
	return CIImageFromID(objc.ID(rv))
}
// A dictionary that contains properties of the image source.
//
// # Discussion
// 
// This dictionary contains the same contents as the [Image Properties]
// accessed using [CGImageSourceCopyProperties(_:_:)].
//
// [CGImageSourceCopyProperties(_:_:)]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyProperties(_:_:)
// [Image Properties]: https://developer.apple.com/documentation/ImageIO/image-properties
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/properties
func (r CIRAWFilter) Properties() foundation.INSDictionary {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("properties"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// A value that indicates the desired scale factor to draw the output image.
//
// # Discussion
// 
// The value should be in the range of `0...1`. The default value is `1`. A
// value of `1` results in a full-size output image, and a value less than `1`
// results in a smaller output image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/scaleFactor
func (r CIRAWFilter) ScaleFactor() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("scaleFactor"))
	return rv
}
func (r CIRAWFilter) SetScaleFactor(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setScaleFactor:"), value)
}
// An optional auxiliary image that represents the semantic segmentation
// glasses matte of the image.
//
// # Discussion
// 
// This matting image segments eyeglasses and sunglasses from all people in
// the visible field of view of the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/semanticSegmentationGlassesMatte
func (r CIRAWFilter) SemanticSegmentationGlassesMatte() ICIImage {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("semanticSegmentationGlassesMatte"))
	return CIImageFromID(objc.ID(rv))
}
// An optional auxiliary image that represents the semantic segmentation hair
// matte of the image.
//
// # Discussion
// 
// This matting image segments hair from all people in the visible field of
// view of the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/semanticSegmentationHairMatte
func (r CIRAWFilter) SemanticSegmentationHairMatte() ICIImage {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("semanticSegmentationHairMatte"))
	return CIImageFromID(objc.ID(rv))
}
// An optional auxiliary image that represents the semantic segmentation skin
// matte of the image.
//
// # Discussion
// 
// This matting image segments skin from all people in the visible field of
// view of the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/semanticSegmentationSkinMatte
func (r CIRAWFilter) SemanticSegmentationSkinMatte() ICIImage {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("semanticSegmentationSkinMatte"))
	return CIImageFromID(objc.ID(rv))
}
// An optional auxiliary image that represents the semantic segmentation sky
// matte of the image.
//
// # Discussion
// 
// This matting image segments the sky from the visible field of view of the
// image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/semanticSegmentationSkyMatte
func (r CIRAWFilter) SemanticSegmentationSkyMatte() ICIImage {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("semanticSegmentationSkyMatte"))
	return CIImageFromID(objc.ID(rv))
}
// An optional auxiliary image that represents the semantic segmentation teeth
// matte of the image.
//
// # Discussion
// 
// This matting image segments the teeth from all people in the visible field
// of view of the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/semanticSegmentationTeethMatte
func (r CIRAWFilter) SemanticSegmentationTeethMatte() ICIImage {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("semanticSegmentationTeethMatte"))
	return CIImageFromID(objc.ID(rv))
}
// A value that indicates the amount to subtract from the shadows in the
// image.
//
// # Discussion
// 
// The default value varies by image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/shadowBias
func (r CIRAWFilter) ShadowBias() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("shadowBias"))
	return rv
}
func (r CIRAWFilter) SetShadowBias(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setShadowBias:"), value)
}
// A value that indicates the amount of sharpness to apply to the edges of the
// image.
//
// # Discussion
// 
// The value should be in the range of `0...1`. The default value varies by
// image. A value of `0` indicates no sharpness, and a value of `1` indicates
// maximum sharpness.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/sharpnessAmount
func (r CIRAWFilter) SharpnessAmount() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("sharpnessAmount"))
	return rv
}
func (r CIRAWFilter) SetSharpnessAmount(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setSharpnessAmount:"), value)
}
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isHighlightRecoveryEnabled
func (r CIRAWFilter) HighlightRecoveryEnabled() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isHighlightRecoveryEnabled"))
	return rv
}
func (r CIRAWFilter) SetHighlightRecoveryEnabled(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setHighlightRecoveryEnabled:"), value)
}
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isHighlightRecoverySupported
func (r CIRAWFilter) HighlightRecoverySupported() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isHighlightRecoverySupported"))
	return rv
}

// An array containing the names of all supported camera models.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/supportedCameraModels
func (_CIRAWFilterClass CIRAWFilterClass) SupportedCameraModels() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_CIRAWFilterClass.class), objc.Sel("supportedCameraModels"))
	return objc.ConvertSliceToStrings(rv)
}

