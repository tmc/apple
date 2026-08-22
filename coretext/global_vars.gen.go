// Code generated from Apple documentation. DO NOT EDIT.

package coretext

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// See: https://developer.apple.com/documentation/CoreText/kCTAdaptiveImageProviderAttributeName
	KCTAdaptiveImageProviderAttributeName string
	// See: https://developer.apple.com/documentation/CoreText/kCTBackgroundColorAttributeName
	KCTBackgroundColorAttributeName string
	// See: https://developer.apple.com/documentation/CoreText/kCTBaselineClassAttributeName
	KCTBaselineClassAttributeName string
	// See: https://developer.apple.com/documentation/CoreText/kCTBaselineClassHanging
	KCTBaselineClassHanging string
	// See: https://developer.apple.com/documentation/CoreText/kCTBaselineClassIdeographicCentered
	KCTBaselineClassIdeographicCentered string
	// See: https://developer.apple.com/documentation/CoreText/kCTBaselineClassIdeographicHigh
	KCTBaselineClassIdeographicHigh string
	// See: https://developer.apple.com/documentation/CoreText/kCTBaselineClassIdeographicLow
	KCTBaselineClassIdeographicLow string
	// See: https://developer.apple.com/documentation/CoreText/kCTBaselineClassMath
	KCTBaselineClassMath string
	// See: https://developer.apple.com/documentation/CoreText/kCTBaselineClassRoman
	KCTBaselineClassRoman string
	// See: https://developer.apple.com/documentation/CoreText/kCTBaselineInfoAttributeName
	KCTBaselineInfoAttributeName string
	// KCTBaselineOffsetAttributeName is vertical offset for text position.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTBaselineOffsetAttributeName
	KCTBaselineOffsetAttributeName string
	// See: https://developer.apple.com/documentation/CoreText/kCTBaselineOriginalFont
	KCTBaselineOriginalFont string
	// See: https://developer.apple.com/documentation/CoreText/kCTBaselineReferenceFont
	KCTBaselineReferenceFont string
	// See: https://developer.apple.com/documentation/CoreText/kCTBaselineReferenceInfoAttributeName
	KCTBaselineReferenceInfoAttributeName string
	// KCTFontAttributeName is the font of the text to which this attribute applies.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontAttributeName
	KCTFontAttributeName string
	// KCTFontBaselineAdjustAttribute is the baseline adjustment for a font reference.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontBaselineAdjustAttribute
	KCTFontBaselineAdjustAttribute string
	// KCTFontCascadeListAttribute is the cascade list used for a font reference.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontCascadeListAttribute
	KCTFontCascadeListAttribute string
	// KCTFontCharacterSetAttribute is the Unicode character coverage set for a font reference.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontCharacterSetAttribute
	KCTFontCharacterSetAttribute string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontCollectionDisallowAutoActivationOption
	KCTFontCollectionDisallowAutoActivationOption string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontCollectionIncludeDisabledFontsOption
	KCTFontCollectionIncludeDisabledFontsOption string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontCollectionRemoveDuplicatesOption
	KCTFontCollectionRemoveDuplicatesOption string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontCopyrightNameKey
	KCTFontCopyrightNameKey string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontDescriptionNameKey
	KCTFontDescriptionNameKey string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontDescriptorLanguageAttribute
	KCTFontDescriptorLanguageAttribute string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontDescriptorMatchingCurrentAssetSize
	KCTFontDescriptorMatchingCurrentAssetSize string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontDescriptorMatchingDescriptors
	KCTFontDescriptorMatchingDescriptors string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontDescriptorMatchingError
	KCTFontDescriptorMatchingError string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontDescriptorMatchingPercentage
	KCTFontDescriptorMatchingPercentage string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontDescriptorMatchingResult
	KCTFontDescriptorMatchingResult string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontDescriptorMatchingSourceDescriptor
	KCTFontDescriptorMatchingSourceDescriptor string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontDescriptorMatchingTotalAssetSize
	KCTFontDescriptorMatchingTotalAssetSize string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontDescriptorMatchingTotalDownloadedSize
	KCTFontDescriptorMatchingTotalDownloadedSize string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontDesignerNameKey
	KCTFontDesignerNameKey string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontDesignerURLNameKey
	KCTFontDesignerURLNameKey string
	// KCTFontDisplayNameAttribute is the name used to display the font.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontDisplayNameAttribute
	KCTFontDisplayNameAttribute string
	// KCTFontDownloadableAttribute is the font downloadable state.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontDownloadableAttribute
	KCTFontDownloadableAttribute string
	// KCTFontDownloadedAttribute is the download state.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontDownloadedAttribute
	KCTFontDownloadedAttribute string
	// KCTFontEnabledAttribute is the font enabled state.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontEnabledAttribute
	KCTFontEnabledAttribute string
	// KCTFontFamilyNameAttribute is the font family name from the font descriptor.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontFamilyNameAttribute
	KCTFontFamilyNameAttribute string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontFamilyNameKey
	KCTFontFamilyNameKey string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontFeatureSampleTextKey
	KCTFontFeatureSampleTextKey string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontFeatureSelectorDefaultKey
	KCTFontFeatureSelectorDefaultKey string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontFeatureSelectorIdentifierKey
	KCTFontFeatureSelectorIdentifierKey string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontFeatureSelectorNameKey
	KCTFontFeatureSelectorNameKey string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontFeatureSelectorSettingKey
	KCTFontFeatureSelectorSettingKey string
	// KCTFontFeatureSettingsAttribute is the font features settings for a font reference.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontFeatureSettingsAttribute
	KCTFontFeatureSettingsAttribute string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontFeatureTooltipTextKey
	KCTFontFeatureTooltipTextKey string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontFeatureTypeExclusiveKey
	KCTFontFeatureTypeExclusiveKey string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontFeatureTypeIdentifierKey
	KCTFontFeatureTypeIdentifierKey string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontFeatureTypeNameKey
	KCTFontFeatureTypeNameKey string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontFeatureTypeSelectorsKey
	KCTFontFeatureTypeSelectorsKey string
	// KCTFontFeaturesAttribute is the font features for a font reference.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontFeaturesAttribute
	KCTFontFeaturesAttribute string
	// KCTFontFixedAdvanceAttribute is a fixed advance to be used for a font reference.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontFixedAdvanceAttribute
	KCTFontFixedAdvanceAttribute string
	// KCTFontFormatAttribute is the recognized format of the font.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontFormatAttribute
	KCTFontFormatAttribute string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontFullNameKey
	KCTFontFullNameKey string
	// KCTFontLanguagesAttribute is a list of covered languages for a font reference.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontLanguagesAttribute
	KCTFontLanguagesAttribute string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontLicenseNameKey
	KCTFontLicenseNameKey string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontLicenseURLNameKey
	KCTFontLicenseURLNameKey string
	// KCTFontMacintoshEncodingsAttribute is the Macintosh encodings for a font reference.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontMacintoshEncodingsAttribute
	KCTFontMacintoshEncodingsAttribute string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontManagerBundleIdentifier
	KCTFontManagerBundleIdentifier string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontManagerErrorDomain
	KCTFontManagerErrorDomain string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontManagerErrorFontURLsKey
	KCTFontManagerErrorFontURLsKey string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontManagerRegisteredFontsChangedNotification
	KCTFontManagerRegisteredFontsChangedNotification string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontManufacturerNameKey
	KCTFontManufacturerNameKey string
	// KCTFontMatrixAttribute is the font transformation matrix when creating a font.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontMatrixAttribute
	KCTFontMatrixAttribute string
	// KCTFontNameAttribute is the PostScript name from the font descriptor.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontNameAttribute
	KCTFontNameAttribute string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontOpenTypeFeatureTag
	KCTFontOpenTypeFeatureTag string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontOpenTypeFeatureValue
	KCTFontOpenTypeFeatureValue string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontOpticalSizeAttribute
	KCTFontOpticalSizeAttribute string
	// KCTFontOrientationAttribute is the orientation for the glyphs of the font.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontOrientationAttribute
	KCTFontOrientationAttribute string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontPostScriptCIDNameKey
	KCTFontPostScriptCIDNameKey string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontPostScriptNameKey
	KCTFontPostScriptNameKey string
	// KCTFontPriorityAttribute is the font priority used by font descriptors when resolving duplicates and sorting match results.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontPriorityAttribute
	KCTFontPriorityAttribute string
	// KCTFontRegistrationScopeAttribute is the font descriptor’s registration scope.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontRegistrationScopeAttribute
	KCTFontRegistrationScopeAttribute string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontSampleTextNameKey
	KCTFontSampleTextNameKey string
	// KCTFontSizeAttribute is the font point size.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontSizeAttribute
	KCTFontSizeAttribute string
	// KCTFontSlantTrait is the normalized slant angle from the font traits dictionary.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontSlantTrait
	KCTFontSlantTrait string
	// KCTFontStyleNameAttribute is the style name of the font.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontStyleNameAttribute
	KCTFontStyleNameAttribute string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontStyleNameKey
	KCTFontStyleNameKey string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontSubFamilyNameKey
	KCTFontSubFamilyNameKey string
	// KCTFontSymbolicTrait is the symbolic traits value from the font traits dictionary.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontSymbolicTrait
	KCTFontSymbolicTrait string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontTrademarkNameKey
	KCTFontTrademarkNameKey string
	// KCTFontTraitsAttribute is the dictionary of font traits for stylistic information.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontTraitsAttribute
	KCTFontTraitsAttribute string
	// KCTFontURLAttribute is the font URL from the font descriptor.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontURLAttribute
	KCTFontURLAttribute string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontUniqueNameKey
	KCTFontUniqueNameKey string
	// KCTFontVariationAttribute is the dictionary of font variation.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontVariationAttribute
	KCTFontVariationAttribute string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontVariationAxesAttribute
	KCTFontVariationAxesAttribute string
	// KCTFontVariationAxisDefaultValueKey is key to get the variation axis default value.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontVariationAxisDefaultValueKey
	KCTFontVariationAxisDefaultValueKey string
	// KCTFontVariationAxisHiddenKey is the key to find out if the axis is hidden.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontVariationAxisHiddenKey
	KCTFontVariationAxisHiddenKey string
	// KCTFontVariationAxisIdentifierKey is key to get the variation axis identifier.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontVariationAxisIdentifierKey
	KCTFontVariationAxisIdentifierKey string
	// KCTFontVariationAxisMaximumValueKey is key to get the variation axis maximum value.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontVariationAxisMaximumValueKey
	KCTFontVariationAxisMaximumValueKey string
	// KCTFontVariationAxisMinimumValueKey is key to get the variation axis minimum value.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontVariationAxisMinimumValueKey
	KCTFontVariationAxisMinimumValueKey string
	// KCTFontVariationAxisNameKey is key to get the localized variation axis name string.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontVariationAxisNameKey
	KCTFontVariationAxisNameKey string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontVendorURLNameKey
	KCTFontVendorURLNameKey string
	// See: https://developer.apple.com/documentation/CoreText/kCTFontVersionNameKey
	KCTFontVersionNameKey string
	// KCTFontWeightTrait is the normalized weight trait from the font traits dictionary.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontWeightTrait
	KCTFontWeightTrait string
	// KCTFontWidthTrait is the normalized proportion (width condense or expand) trait from the font traits dictionary.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFontWidthTrait
	KCTFontWidthTrait string
	// KCTForegroundColorAttributeName is the foreground color of the text to which this attribute applies.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTForegroundColorAttributeName
	KCTForegroundColorAttributeName string
	// KCTForegroundColorFromContextAttributeName is sets a foreground color using the context’s fill color.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTForegroundColorFromContextAttributeName
	KCTForegroundColorFromContextAttributeName string
	// KCTFrameClippingPathsAttributeName is specifies array of paths to clip frame.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFrameClippingPathsAttributeName
	KCTFrameClippingPathsAttributeName string
	// KCTFramePathClippingPathAttributeName is specifies clipping path.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFramePathClippingPathAttributeName
	KCTFramePathClippingPathAttributeName string
	// KCTFramePathFillRuleAttributeName is the key used to specify the fill rule for a frame.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFramePathFillRuleAttributeName
	KCTFramePathFillRuleAttributeName string
	// KCTFramePathWidthAttributeName is the key used to specify the frame width.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFramePathWidthAttributeName
	KCTFramePathWidthAttributeName string
	// KCTFrameProgressionAttributeName is specifies progression for a frame.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTFrameProgressionAttributeName
	KCTFrameProgressionAttributeName string
	// KCTGlyphInfoAttributeName is the glyph info object to apply to the text associated with this attribute.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTGlyphInfoAttributeName
	KCTGlyphInfoAttributeName string
	// KCTHorizontalInVerticalFormsAttributeName is setting text in tate-chu-yoko form (horizontal numerals in vertical text).
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTHorizontalInVerticalFormsAttributeName
	KCTHorizontalInVerticalFormsAttributeName string
	// KCTKernAttributeName is the amount to kern the next character.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTKernAttributeName
	KCTKernAttributeName string
	// KCTLanguageAttributeName is the name of the text language.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTLanguageAttributeName
	KCTLanguageAttributeName string
	// KCTLigatureAttributeName is the type of ligatures to use.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTLigatureAttributeName
	KCTLigatureAttributeName string
	// KCTParagraphStyleAttributeName is the paragraph style of the text to which this attribute applies.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTParagraphStyleAttributeName
	KCTParagraphStyleAttributeName string
	// See: https://developer.apple.com/documentation/CoreText/kCTRubyAnnotationAttributeName
	KCTRubyAnnotationAttributeName string
	// See: https://developer.apple.com/documentation/CoreText/kCTRubyAnnotationScaleToFitAttributeName
	KCTRubyAnnotationScaleToFitAttributeName string
	// See: https://developer.apple.com/documentation/CoreText/kCTRubyAnnotationSizeFactorAttributeName
	KCTRubyAnnotationSizeFactorAttributeName string
	// KCTRunDelegateAttributeName is the run-delegate object to apply to an attribute range of the string.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTRunDelegateAttributeName
	KCTRunDelegateAttributeName string
	// KCTStrokeColorAttributeName is the stroke color.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTStrokeColorAttributeName
	KCTStrokeColorAttributeName string
	// KCTStrokeWidthAttributeName is the stroke width.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTStrokeWidthAttributeName
	KCTStrokeWidthAttributeName string
	// KCTSuperscriptAttributeName is controls vertical text positioning.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTSuperscriptAttributeName
	KCTSuperscriptAttributeName string
	// See: https://developer.apple.com/documentation/CoreText/kCTTabColumnTerminatorsAttributeName
	KCTTabColumnTerminatorsAttributeName string
	// KCTTrackingAttributeName is the tracking for the text.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTTrackingAttributeName
	KCTTrackingAttributeName string
	// KCTTypesetterOptionAllowUnboundedLayout is a key that specifies whether the text system lays out text that requires unreasonable effort.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTTypesetterOptionAllowUnboundedLayout
	KCTTypesetterOptionAllowUnboundedLayout string
	// KCTTypesetterOptionForcedEmbeddingLevel is a key that specifies the embedding level of the typesetter’s text.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTTypesetterOptionForcedEmbeddingLevel
	KCTTypesetterOptionForcedEmbeddingLevel string
	// KCTUnderlineColorAttributeName is the underline color.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTUnderlineColorAttributeName
	KCTUnderlineColorAttributeName string
	// KCTUnderlineStyleAttributeName is the style of underlining, to be applied at render time, for the text to which this attribute applies.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTUnderlineStyleAttributeName
	KCTUnderlineStyleAttributeName string
	// KCTVerticalFormsAttributeName is the orientation of the glyphs in the text to which this attribute applies.
	//
	// See: https://developer.apple.com/documentation/CoreText/kCTVerticalFormsAttributeName
	KCTVerticalFormsAttributeName string
	// See: https://developer.apple.com/documentation/CoreText/kCTWritingDirectionAttributeName
	KCTWritingDirectionAttributeName string
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTAdaptiveImageProviderAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTAdaptiveImageProviderAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTBackgroundColorAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTBackgroundColorAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTBaselineClassAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTBaselineClassAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTBaselineClassHanging"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTBaselineClassHanging = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTBaselineClassIdeographicCentered"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTBaselineClassIdeographicCentered = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTBaselineClassIdeographicHigh"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTBaselineClassIdeographicHigh = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTBaselineClassIdeographicLow"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTBaselineClassIdeographicLow = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTBaselineClassMath"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTBaselineClassMath = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTBaselineClassRoman"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTBaselineClassRoman = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTBaselineInfoAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTBaselineInfoAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTBaselineOffsetAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTBaselineOffsetAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTBaselineOriginalFont"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTBaselineOriginalFont = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTBaselineReferenceFont"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTBaselineReferenceFont = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTBaselineReferenceInfoAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTBaselineReferenceInfoAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontBaselineAdjustAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontBaselineAdjustAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontCascadeListAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontCascadeListAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontCharacterSetAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontCharacterSetAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontCollectionDisallowAutoActivationOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontCollectionDisallowAutoActivationOption = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontCollectionIncludeDisabledFontsOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontCollectionIncludeDisabledFontsOption = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontCollectionRemoveDuplicatesOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontCollectionRemoveDuplicatesOption = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontCopyrightNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontCopyrightNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontDescriptionNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontDescriptionNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontDescriptorLanguageAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontDescriptorLanguageAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontDescriptorMatchingCurrentAssetSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontDescriptorMatchingCurrentAssetSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontDescriptorMatchingDescriptors"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontDescriptorMatchingDescriptors = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontDescriptorMatchingError"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontDescriptorMatchingError = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontDescriptorMatchingPercentage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontDescriptorMatchingPercentage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontDescriptorMatchingResult"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontDescriptorMatchingResult = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontDescriptorMatchingSourceDescriptor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontDescriptorMatchingSourceDescriptor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontDescriptorMatchingTotalAssetSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontDescriptorMatchingTotalAssetSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontDescriptorMatchingTotalDownloadedSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontDescriptorMatchingTotalDownloadedSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontDesignerNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontDesignerNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontDesignerURLNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontDesignerURLNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontDisplayNameAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontDisplayNameAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontDownloadableAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontDownloadableAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontDownloadedAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontDownloadedAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontEnabledAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontEnabledAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontFamilyNameAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontFamilyNameAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontFamilyNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontFamilyNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontFeatureSampleTextKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontFeatureSampleTextKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontFeatureSelectorDefaultKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontFeatureSelectorDefaultKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontFeatureSelectorIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontFeatureSelectorIdentifierKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontFeatureSelectorNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontFeatureSelectorNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontFeatureSelectorSettingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontFeatureSelectorSettingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontFeatureSettingsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontFeatureSettingsAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontFeatureTooltipTextKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontFeatureTooltipTextKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontFeatureTypeExclusiveKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontFeatureTypeExclusiveKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontFeatureTypeIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontFeatureTypeIdentifierKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontFeatureTypeNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontFeatureTypeNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontFeatureTypeSelectorsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontFeatureTypeSelectorsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontFeaturesAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontFeaturesAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontFixedAdvanceAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontFixedAdvanceAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontFormatAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontFormatAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontFullNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontFullNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontLanguagesAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontLanguagesAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontLicenseNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontLicenseNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontLicenseURLNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontLicenseURLNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontMacintoshEncodingsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontMacintoshEncodingsAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontManagerBundleIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontManagerBundleIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontManagerErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontManagerErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontManagerErrorFontURLsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontManagerErrorFontURLsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontManagerRegisteredFontsChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontManagerRegisteredFontsChangedNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontManufacturerNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontManufacturerNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontMatrixAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontMatrixAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontNameAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontNameAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontOpenTypeFeatureTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontOpenTypeFeatureTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontOpenTypeFeatureValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontOpenTypeFeatureValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontOpticalSizeAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontOpticalSizeAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontOrientationAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontOrientationAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontPostScriptCIDNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontPostScriptCIDNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontPostScriptNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontPostScriptNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontPriorityAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontPriorityAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontRegistrationScopeAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontRegistrationScopeAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontSampleTextNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontSampleTextNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontSizeAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontSizeAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontSlantTrait"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontSlantTrait = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontStyleNameAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontStyleNameAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontStyleNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontStyleNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontSubFamilyNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontSubFamilyNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontSymbolicTrait"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontSymbolicTrait = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontTrademarkNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontTrademarkNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontTraitsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontTraitsAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontURLAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontURLAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontUniqueNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontUniqueNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontVariationAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontVariationAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontVariationAxesAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontVariationAxesAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontVariationAxisDefaultValueKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontVariationAxisDefaultValueKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontVariationAxisHiddenKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontVariationAxisHiddenKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontVariationAxisIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontVariationAxisIdentifierKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontVariationAxisMaximumValueKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontVariationAxisMaximumValueKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontVariationAxisMinimumValueKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontVariationAxisMinimumValueKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontVariationAxisNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontVariationAxisNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontVendorURLNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontVendorURLNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontVersionNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontVersionNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontWeightTrait"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontWeightTrait = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFontWidthTrait"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFontWidthTrait = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTForegroundColorAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTForegroundColorAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTForegroundColorFromContextAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTForegroundColorFromContextAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFrameClippingPathsAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFrameClippingPathsAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFramePathClippingPathAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFramePathClippingPathAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFramePathFillRuleAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFramePathFillRuleAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFramePathWidthAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFramePathWidthAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTFrameProgressionAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTFrameProgressionAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTGlyphInfoAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTGlyphInfoAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTHorizontalInVerticalFormsAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTHorizontalInVerticalFormsAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTKernAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTKernAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTLanguageAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTLanguageAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTLigatureAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTLigatureAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTParagraphStyleAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTParagraphStyleAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTRubyAnnotationAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTRubyAnnotationAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTRubyAnnotationScaleToFitAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTRubyAnnotationScaleToFitAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTRubyAnnotationSizeFactorAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTRubyAnnotationSizeFactorAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTRunDelegateAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTRunDelegateAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTStrokeColorAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTStrokeColorAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTStrokeWidthAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTStrokeWidthAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTSuperscriptAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTSuperscriptAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTTabColumnTerminatorsAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTTabColumnTerminatorsAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTTrackingAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTTrackingAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTTypesetterOptionAllowUnboundedLayout"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTTypesetterOptionAllowUnboundedLayout = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTTypesetterOptionForcedEmbeddingLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTTypesetterOptionForcedEmbeddingLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTUnderlineColorAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTUnderlineColorAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTUnderlineStyleAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTUnderlineStyleAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTVerticalFormsAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTVerticalFormsAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCTWritingDirectionAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCTWritingDirectionAttributeName = objc.GoString(cstr)
			}
		}
	}

}
