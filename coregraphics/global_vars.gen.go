// Code generated from Apple documentation. DO NOT EDIT.

package coregraphics

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
)

var (
	// CGAffineTransformIdentity is the identity transform.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformIdentity
	CGAffineTransformIdentity corefoundation.CGAffineTransform
)

var (
	// CGPointZero is a point constant with location `(0,0)`. The zero point is equivalent to `CGPointMake(0,0)`.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGPointZero
	CGPointZero corefoundation.CGPoint
)

var (
	// CGRectInfinite is a rectangle that has infinite extent.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGRectInfinite
	CGRectInfinite corefoundation.CGRect
	// CGRectNull is the null rectangle, representing an invalid value.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGRectNull
	CGRectNull corefoundation.CGRect
	// CGRectZero is a rectangle constant with location `(0,0)`, and width and height of 0. The zero rectangle is equivalent to `CGRectMake(0,0,0,0)`.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGRectZero
	CGRectZero corefoundation.CGRect
)

var (
	// CGSizeZero is a size constant with width and height of `0`. The zero size is equivalent to `CGSizeMake(0,0)`.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGSizeZero
	CGSizeZero corefoundation.CGSize
)

var (
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGAdaptiveMaximumBitDepth
	KCGAdaptiveMaximumBitDepth string
	// KCGColorBlack is the black color in the Generic gray color space.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGColorBlack
	KCGColorBlack string
	// KCGColorClear is the clear color in the Generic gray color space.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGColorClear
	KCGColorClear string
	// KCGColorConversionBlackPointCompensation is an option for whether to apply black point compensation when converting between color profiles.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/conversionBlackPointCompensation
	KCGColorConversionBlackPointCompensation string
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/conversionTRCSize
	KCGColorConversionTRCSize string
	// KCGColorSpaceACESCGLinear is the ACEScg color space.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/acescgLinear
	KCGColorSpaceACESCGLinear string
	// KCGColorSpaceAdobeRGB1998 is the Adobe RGB (1998) color space.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/adobeRGB1998
	KCGColorSpaceAdobeRGB1998 string
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/coreMedia709
	KCGColorSpaceCoreMedia709 string
	// KCGColorSpaceDCIP3 is the DCI P3 color space, which is the digital cinema standard.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/dcip3
	KCGColorSpaceDCIP3 string
	// KCGColorSpaceDisplayP3 is the Display P3 color space, created by Apple.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/displayP3
	KCGColorSpaceDisplayP3 string
	// KCGColorSpaceDisplayP3_HLG is the Display P3 color space, using the HLG transfer function.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/displayP3_HLG
	KCGColorSpaceDisplayP3_HLG string
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/displayP3_PQ
	KCGColorSpaceDisplayP3_PQ string
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/extendedDisplayP3
	KCGColorSpaceExtendedDisplayP3 string
	// KCGColorSpaceExtendedGray is the extended gray color space.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/extendedGray
	KCGColorSpaceExtendedGray string
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/extendedITUR_2020
	KCGColorSpaceExtendedITUR_2020 string
	// KCGColorSpaceExtendedLinearDisplayP3 is the Display P3 color space with a linear transfer function and extended-range values.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/extendedLinearDisplayP3
	KCGColorSpaceExtendedLinearDisplayP3 string
	// KCGColorSpaceExtendedLinearGray is the extended gray color space with a linear transfer function.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/extendedLinearGray
	KCGColorSpaceExtendedLinearGray string
	// KCGColorSpaceExtendedLinearITUR_2020 is the recommendation of the International Telecommunication Union (ITU) Radiocommunication sector for the BT.2020 color space, with a linear transfer function and extended range values.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/extendedLinearITUR_2020
	KCGColorSpaceExtendedLinearITUR_2020 string
	// KCGColorSpaceExtendedLinearSRGB is the sRGB color space with a linear transfer function and extended-range values.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/extendedLinearSRGB
	KCGColorSpaceExtendedLinearSRGB string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGColorSpaceExtendedRange
	KCGColorSpaceExtendedRange string
	// KCGColorSpaceExtendedSRGB is the extended sRGB color space.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/extendedSRGB
	KCGColorSpaceExtendedSRGB string
	// KCGColorSpaceGenericCMYK is the generic CMYK color space.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/genericCMYK
	KCGColorSpaceGenericCMYK string
	// KCGColorSpaceGenericGrayGamma2_2 is the generic gray color space that has an exponential transfer function with a power of 2.2.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/genericGrayGamma2_2
	KCGColorSpaceGenericGrayGamma2_2 string
	// KCGColorSpaceGenericLab is the generic LAB color space.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/genericLab
	KCGColorSpaceGenericLab string
	// KCGColorSpaceGenericRGBLinear is the generic RGB color space with a linear transfer function.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/genericRGBLinear
	KCGColorSpaceGenericRGBLinear string
	// KCGColorSpaceGenericXYZ is the XYZ color space, as defined by the CIE 1931 standard.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/genericXYZ
	KCGColorSpaceGenericXYZ string
	// KCGColorSpaceITUR_2020 is the recommendation of the International Telecommunication Union (ITU) Radiocommunication sector for the BT.2020 color space.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/itur_2020
	KCGColorSpaceITUR_2020 string
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/itur_2020_sRGBGamma
	KCGColorSpaceITUR_2020_sRGBGamma string
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/itur_2100_HLG
	KCGColorSpaceITUR_2100_HLG string
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/itur_2100_PQ
	KCGColorSpaceITUR_2100_PQ string
	// KCGColorSpaceITUR_709 is the recommendation of the International Telecommunication Union (ITU) Radiocommunication sector for the BT.709 color space.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/itur_709
	KCGColorSpaceITUR_709 string
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/itur_709_HLG
	KCGColorSpaceITUR_709_HLG string
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/itur_709_PQ
	KCGColorSpaceITUR_709_PQ string
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/linearDisplayP3
	KCGColorSpaceLinearDisplayP3 string
	// KCGColorSpaceLinearGray is the gray color space using a linear transfer function.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/linearGray
	KCGColorSpaceLinearGray string
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/linearITUR_2020
	KCGColorSpaceLinearITUR_2020 string
	// KCGColorSpaceLinearSRGB is the sRGB color space with a linear transfer function.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/linearSRGB
	KCGColorSpaceLinearSRGB string
	// KCGColorSpaceROMMRGB is the Reference Output Medium Metric (ROMM) RGB color space.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/rommrgb
	KCGColorSpaceROMMRGB string
	// KCGColorSpaceSRGB is the standard Red Green Blue (sRGB) color space.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/sRGB
	KCGColorSpaceSRGB string
	// KCGColorWhite is the white color in the Generic gray color space.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGColorWhite
	KCGColorWhite string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGContentAverageLightLevel
	KCGContentAverageLightLevel string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGContentAverageLightLevelNits
	KCGContentAverageLightLevelNits string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGDisplayShowDuplicateLowResolutionModes
	KCGDisplayShowDuplicateLowResolutionModes string
	// KCGDisplayStreamYCbCrMatrix_ITU_R_601_4 is specifies the YCbCr to RGB conversion matrix for standard digital television (ITU R 601) images.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStream/yCbCrMatrix_ITU_R_601_4
	KCGDisplayStreamYCbCrMatrix_ITU_R_601_4 string
	// KCGDisplayStreamYCbCrMatrix_ITU_R_709_2 is specifies the YCbCr to RGB conversion matrix for HDTV digital television (ITU R 709) images.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStream/yCbCrMatrix_ITU_R_709_2
	KCGDisplayStreamYCbCrMatrix_ITU_R_709_2 string
	// KCGDisplayStreamYCbCrMatrix_SMPTE_240M_1995 is specifies the YCbCR to RGB conversion matrix for 1920 x 1135 HDTV (SMPTE 240M 1995).
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStream/yCbCrMatrix_SMPTE_240M_1995
	KCGDisplayStreamYCbCrMatrix_SMPTE_240M_1995 string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGDynamicRangeConstrained
	KCGDynamicRangeConstrained string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGDynamicRangeHigh
	KCGDynamicRangeHigh string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGDynamicRangeStandard
	KCGDynamicRangeStandard string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGEXRToneMappingGammaDefog
	KCGEXRToneMappingGammaDefog string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGEXRToneMappingGammaExposure
	KCGEXRToneMappingGammaExposure string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGEXRToneMappingGammaKneeHigh
	KCGEXRToneMappingGammaKneeHigh string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGEXRToneMappingGammaKneeLow
	KCGEXRToneMappingGammaKneeLow string
	// KCGFontVariationAxisDefaultValue is the key used to obtain the default variation axis value from a variation axis dictionary.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/variationAxisDefaultValue
	KCGFontVariationAxisDefaultValue string
	// KCGFontVariationAxisMaxValue is the key used to obtain the maximum variation axis value from a variation axis dictionary.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/variationAxisMaxValue
	KCGFontVariationAxisMaxValue string
	// KCGFontVariationAxisMinValue is the key used to obtain the minimum variation axis value from a variation axis dictionary.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/variationAxisMinValue
	KCGFontVariationAxisMinValue string
	// KCGFontVariationAxisName is the key used to obtain the variation axis name from a variation axis dictionary.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/variationAxisName
	KCGFontVariationAxisName string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFContextAccessPermissions
	KCGPDFContextAccessPermissions string
	// KCGPDFContextAllowsCopying is whether the document allows copying when unlocked with the user password.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFContextAllowsCopying
	KCGPDFContextAllowsCopying string
	// KCGPDFContextAllowsPrinting is whether the document allows printing when unlocked with the user password.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFContextAllowsPrinting
	KCGPDFContextAllowsPrinting string
	// KCGPDFContextArtBox is the art box for the document or for a given page.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFContextArtBox
	KCGPDFContextArtBox string
	// KCGPDFContextAuthor is the corresponding value is a string that represents the name of the person who created the document. This key is optional.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFContextAuthor
	KCGPDFContextAuthor string
	// KCGPDFContextBleedBox is the bleed box for the document or for a given page.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFContextBleedBox
	KCGPDFContextBleedBox string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFContextCreateLinearizedPDF
	KCGPDFContextCreateLinearizedPDF string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFContextCreatePDFA
	KCGPDFContextCreatePDFA string
	// KCGPDFContextCreator is the corresponding value is a string that represents the name of the application used to produce the document. This key is optional.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFContextCreator
	KCGPDFContextCreator string
	// KCGPDFContextCropBox is the crop box for the document or for a given page.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFContextCropBox
	KCGPDFContextCropBox string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFContextEncryptionKeyLength
	KCGPDFContextEncryptionKeyLength string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFContextKeywords
	KCGPDFContextKeywords string
	// KCGPDFContextMediaBox is the media box for the document or for a given page.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFContextMediaBox
	KCGPDFContextMediaBox string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFContextOutputIntent
	KCGPDFContextOutputIntent string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFContextOutputIntents
	KCGPDFContextOutputIntents string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFContextOwnerPassword
	KCGPDFContextOwnerPassword string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFContextSubject
	KCGPDFContextSubject string
	// KCGPDFContextTitle is the corresponding value is a string that represents the title of the document. This key is optional.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFContextTitle
	KCGPDFContextTitle string
	// KCGPDFContextTrimBox is the trim box for the document or for a given page.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFContextTrimBox
	KCGPDFContextTrimBox string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFContextUserPassword
	KCGPDFContextUserPassword string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFOutlineChildren
	KCGPDFOutlineChildren string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFOutlineDestination
	KCGPDFOutlineDestination string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFOutlineDestinationRect
	KCGPDFOutlineDestinationRect string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFOutlineTitle
	KCGPDFOutlineTitle string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFXDestinationOutputProfile
	KCGPDFXDestinationOutputProfile string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFXInfo
	KCGPDFXInfo string
	// KCGPDFXOutputCondition is a text string identifying the intended output device or production condition in a human-readable form.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFXOutputCondition
	KCGPDFXOutputCondition string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFXOutputConditionIdentifier
	KCGPDFXOutputConditionIdentifier string
	// KCGPDFXOutputIntentSubtype is the output intent subtype. This key is required.
	//
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFXOutputIntentSubtype
	KCGPDFXOutputIntentSubtype string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPDFXRegistryName
	KCGPDFXRegistryName string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGPreferredDynamicRange
	KCGPreferredDynamicRange string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGSkipBoostToHDR
	KCGSkipBoostToHDR string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGUse100nitsHLGOOTF
	KCGUse100nitsHLGOOTF string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGUseBT1886ForCoreVideoGamma
	KCGUseBT1886ForCoreVideoGamma string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGUseLegacyHDREcosystem
	KCGUseLegacyHDREcosystem string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGWindowAlpha
	KCGWindowAlpha string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGWindowBackingLocationVideoMemory
	KCGWindowBackingLocationVideoMemory string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGWindowBounds
	KCGWindowBounds string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGWindowIsOnscreen
	KCGWindowIsOnscreen string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGWindowLayer
	KCGWindowLayer string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGWindowMemoryUsage
	KCGWindowMemoryUsage string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGWindowName
	KCGWindowName string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGWindowNumber
	KCGWindowNumber string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGWindowOwnerName
	KCGWindowOwnerName string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGWindowOwnerPID
	KCGWindowOwnerPID string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGWindowSharingState
	KCGWindowSharingState string
	// See: https://developer.apple.com/documentation/CoreGraphics/kCGWindowStoreType
	KCGWindowStoreType string
)

var (
	// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagProperty/actualText
	KCGPDFTagPropertyActualText CGPDFTagProperty
	// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagProperty/alternativeText
	KCGPDFTagPropertyAlternativeText CGPDFTagProperty
	// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagProperty/languageText
	KCGPDFTagPropertyLanguageText CGPDFTagProperty
	// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagProperty/titleText
	KCGPDFTagPropertyTitleText CGPDFTagProperty
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CGAffineTransformIdentity"); err == nil && ptr != 0 {
		CGAffineTransformIdentity = *(*corefoundation.CGAffineTransform)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CGPointZero"); err == nil && ptr != 0 {
		CGPointZero = *(*corefoundation.CGPoint)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CGRectInfinite"); err == nil && ptr != 0 {
		CGRectInfinite = *(*corefoundation.CGRect)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CGRectNull"); err == nil && ptr != 0 {
		CGRectNull = *(*corefoundation.CGRect)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CGRectZero"); err == nil && ptr != 0 {
		CGRectZero = *(*corefoundation.CGRect)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CGSizeZero"); err == nil && ptr != 0 {
		CGSizeZero = *(*corefoundation.CGSize)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGAdaptiveMaximumBitDepth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGAdaptiveMaximumBitDepth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorBlack"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorBlack = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorClear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorClear = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorConversionBlackPointCompensation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorConversionBlackPointCompensation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorConversionTRCSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorConversionTRCSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceACESCGLinear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceACESCGLinear = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceAdobeRGB1998"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceAdobeRGB1998 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceCoreMedia709"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceCoreMedia709 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceDCIP3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceDCIP3 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceDisplayP3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceDisplayP3 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceDisplayP3_HLG"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceDisplayP3_HLG = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceDisplayP3_PQ"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceDisplayP3_PQ = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceExtendedDisplayP3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceExtendedDisplayP3 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceExtendedGray"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceExtendedGray = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceExtendedITUR_2020"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceExtendedITUR_2020 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceExtendedLinearDisplayP3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceExtendedLinearDisplayP3 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceExtendedLinearGray"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceExtendedLinearGray = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceExtendedLinearITUR_2020"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceExtendedLinearITUR_2020 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceExtendedLinearSRGB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceExtendedLinearSRGB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceExtendedRange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceExtendedRange = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceExtendedSRGB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceExtendedSRGB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceGenericCMYK"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceGenericCMYK = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceGenericGrayGamma2_2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceGenericGrayGamma2_2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceGenericLab"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceGenericLab = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceGenericRGBLinear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceGenericRGBLinear = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceGenericXYZ"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceGenericXYZ = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceITUR_2020"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceITUR_2020 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceITUR_2020_sRGBGamma"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceITUR_2020_sRGBGamma = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceITUR_2100_HLG"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceITUR_2100_HLG = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceITUR_2100_PQ"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceITUR_2100_PQ = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceITUR_709"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceITUR_709 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceITUR_709_HLG"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceITUR_709_HLG = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceITUR_709_PQ"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceITUR_709_PQ = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceLinearDisplayP3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceLinearDisplayP3 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceLinearGray"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceLinearGray = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceLinearITUR_2020"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceLinearITUR_2020 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceLinearSRGB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceLinearSRGB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceROMMRGB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceROMMRGB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorSpaceSRGB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorSpaceSRGB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGColorWhite"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGColorWhite = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGContentAverageLightLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGContentAverageLightLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGContentAverageLightLevelNits"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGContentAverageLightLevelNits = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGDisplayShowDuplicateLowResolutionModes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGDisplayShowDuplicateLowResolutionModes = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGDisplayStreamYCbCrMatrix_ITU_R_601_4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGDisplayStreamYCbCrMatrix_ITU_R_601_4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGDisplayStreamYCbCrMatrix_ITU_R_709_2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGDisplayStreamYCbCrMatrix_ITU_R_709_2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGDisplayStreamYCbCrMatrix_SMPTE_240M_1995"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGDisplayStreamYCbCrMatrix_SMPTE_240M_1995 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGDynamicRangeConstrained"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGDynamicRangeConstrained = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGDynamicRangeHigh"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGDynamicRangeHigh = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGDynamicRangeStandard"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGDynamicRangeStandard = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGEXRToneMappingGammaDefog"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGEXRToneMappingGammaDefog = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGEXRToneMappingGammaExposure"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGEXRToneMappingGammaExposure = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGEXRToneMappingGammaKneeHigh"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGEXRToneMappingGammaKneeHigh = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGEXRToneMappingGammaKneeLow"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGEXRToneMappingGammaKneeLow = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGFontVariationAxisDefaultValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGFontVariationAxisDefaultValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGFontVariationAxisMaxValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGFontVariationAxisMaxValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGFontVariationAxisMinValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGFontVariationAxisMinValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGFontVariationAxisName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGFontVariationAxisName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFContextAccessPermissions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFContextAccessPermissions = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFContextAllowsCopying"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFContextAllowsCopying = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFContextAllowsPrinting"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFContextAllowsPrinting = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFContextArtBox"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFContextArtBox = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFContextAuthor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFContextAuthor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFContextBleedBox"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFContextBleedBox = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFContextCreateLinearizedPDF"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFContextCreateLinearizedPDF = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFContextCreatePDFA"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFContextCreatePDFA = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFContextCreator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFContextCreator = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFContextCropBox"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFContextCropBox = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFContextEncryptionKeyLength"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFContextEncryptionKeyLength = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFContextKeywords"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFContextKeywords = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFContextMediaBox"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFContextMediaBox = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFContextOutputIntent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFContextOutputIntent = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFContextOutputIntents"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFContextOutputIntents = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFContextOwnerPassword"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFContextOwnerPassword = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFContextSubject"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFContextSubject = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFContextTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFContextTitle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFContextTrimBox"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFContextTrimBox = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFContextUserPassword"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFContextUserPassword = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFOutlineChildren"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFOutlineChildren = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFOutlineDestination"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFOutlineDestination = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFOutlineDestinationRect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFOutlineDestinationRect = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFOutlineTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFOutlineTitle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFTagPropertyActualText"); err == nil && ptr != 0 {
		KCGPDFTagPropertyActualText = *(*CGPDFTagProperty)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFTagPropertyAlternativeText"); err == nil && ptr != 0 {
		KCGPDFTagPropertyAlternativeText = *(*CGPDFTagProperty)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFTagPropertyLanguageText"); err == nil && ptr != 0 {
		KCGPDFTagPropertyLanguageText = *(*CGPDFTagProperty)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFTagPropertyTitleText"); err == nil && ptr != 0 {
		KCGPDFTagPropertyTitleText = *(*CGPDFTagProperty)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFXDestinationOutputProfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFXDestinationOutputProfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFXInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFXInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFXOutputCondition"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFXOutputCondition = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFXOutputConditionIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFXOutputConditionIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFXOutputIntentSubtype"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFXOutputIntentSubtype = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPDFXRegistryName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPDFXRegistryName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGPreferredDynamicRange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGPreferredDynamicRange = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGSkipBoostToHDR"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGSkipBoostToHDR = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGUse100nitsHLGOOTF"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGUse100nitsHLGOOTF = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGUseBT1886ForCoreVideoGamma"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGUseBT1886ForCoreVideoGamma = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGUseLegacyHDREcosystem"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGUseLegacyHDREcosystem = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGWindowAlpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGWindowAlpha = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGWindowBackingLocationVideoMemory"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGWindowBackingLocationVideoMemory = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGWindowBounds"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGWindowBounds = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGWindowIsOnscreen"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGWindowIsOnscreen = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGWindowLayer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGWindowLayer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGWindowMemoryUsage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGWindowMemoryUsage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGWindowName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGWindowName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGWindowNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGWindowNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGWindowOwnerName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGWindowOwnerName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGWindowOwnerPID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGWindowOwnerPID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGWindowSharingState"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGWindowSharingState = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCGWindowStoreType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCGWindowStoreType = objc.GoString(cstr)
			}
		}
	}

}
