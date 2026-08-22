// Code generated from Apple documentation for CoreText. DO NOT EDIT.

package coretext

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/CoreText/CTCharacterCollection
type CTCharacterCollection uint16

const (
	// KCTCharacterCollectionAdobeCNS1: The Adobe-CNS1 mapping.
	KCTCharacterCollectionAdobeCNS1 CTCharacterCollection = 1
	// KCTCharacterCollectionAdobeGB1: The Adobe-GB1 mapping.
	KCTCharacterCollectionAdobeGB1 CTCharacterCollection = 2
	// KCTCharacterCollectionAdobeJapan1: The Adobe-Japan1 mapping.
	KCTCharacterCollectionAdobeJapan1 CTCharacterCollection = 3
	// KCTCharacterCollectionAdobeJapan2: The Adobe-Japan2 mapping.
	KCTCharacterCollectionAdobeJapan2 CTCharacterCollection = 4
	// KCTCharacterCollectionAdobeKorea1: The Adobe-Korea1 mapping.
	KCTCharacterCollectionAdobeKorea1 CTCharacterCollection = 5
	// KCTCharacterCollectionIdentityMapping: The character identifier is equal to the glyph index.
	KCTCharacterCollectionIdentityMapping CTCharacterCollection = 0
)

func (e CTCharacterCollection) String() string {
	switch e {
	case KCTCharacterCollectionAdobeCNS1:
		return "KCTCharacterCollectionAdobeCNS1"
	case KCTCharacterCollectionAdobeGB1:
		return "KCTCharacterCollectionAdobeGB1"
	case KCTCharacterCollectionAdobeJapan1:
		return "KCTCharacterCollectionAdobeJapan1"
	case KCTCharacterCollectionAdobeJapan2:
		return "KCTCharacterCollectionAdobeJapan2"
	case KCTCharacterCollectionAdobeKorea1:
		return "KCTCharacterCollectionAdobeKorea1"
	case KCTCharacterCollectionIdentityMapping:
		return "KCTCharacterCollectionIdentityMapping"
	default:
		return fmt.Sprintf("CTCharacterCollection(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyOptions
type CTFontCollectionCopyOptions uint32

const (
	// KCTFontCollectionCopyDefaultOptions: Passing this option indicates that defaults are to be used.
	KCTFontCollectionCopyDefaultOptions CTFontCollectionCopyOptions = 0
	// KCTFontCollectionCopyStandardSort: Passing this option indicates that the return values should be sorted in standard UI order, suitable for display to the user.
	KCTFontCollectionCopyStandardSort CTFontCollectionCopyOptions = 2
	// KCTFontCollectionCopyUnique: Passing this option indicates that duplicate values should be removed from the results.
	KCTFontCollectionCopyUnique CTFontCollectionCopyOptions = 1
)

func (e CTFontCollectionCopyOptions) String() string {
	switch e {
	case KCTFontCollectionCopyDefaultOptions:
		return "KCTFontCollectionCopyDefaultOptions"
	case KCTFontCollectionCopyStandardSort:
		return "KCTFontCollectionCopyStandardSort"
	case KCTFontCollectionCopyUnique:
		return "KCTFontCollectionCopyUnique"
	default:
		return fmt.Sprintf("CTFontCollectionCopyOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchingState
type CTFontDescriptorMatchingState uint32

const (
	// KCTFontDescriptorMatchingDidBegin: A state that indicates matching is about to begin.
	KCTFontDescriptorMatchingDidBegin CTFontDescriptorMatchingState = 0
	// KCTFontDescriptorMatchingDidFailWithError: A state that indicates an error.
	KCTFontDescriptorMatchingDidFailWithError CTFontDescriptorMatchingState = 8
	// KCTFontDescriptorMatchingDidFinish: A state that indicates matching is done.
	KCTFontDescriptorMatchingDidFinish CTFontDescriptorMatchingState = 1
	// KCTFontDescriptorMatchingDidFinishDownloading: A state that indicates downloading is done.
	KCTFontDescriptorMatchingDidFinishDownloading CTFontDescriptorMatchingState = 6
	// KCTFontDescriptorMatchingDidMatch: A state that indicates the font descriptor match is successful.
	KCTFontDescriptorMatchingDidMatch CTFontDescriptorMatchingState = 7
	// KCTFontDescriptorMatchingDownloading: A state that indicates downloading is in progress.
	KCTFontDescriptorMatchingDownloading CTFontDescriptorMatchingState = 5
	// KCTFontDescriptorMatchingStalled: A state that indicates that matching is stalled, such as while waiting for a server response.
	KCTFontDescriptorMatchingStalled CTFontDescriptorMatchingState = 3
	// KCTFontDescriptorMatchingWillBeginDownloading: A state that indicates downloading is about to begin.
	KCTFontDescriptorMatchingWillBeginDownloading CTFontDescriptorMatchingState = 4
	// KCTFontDescriptorMatchingWillBeginQuerying: A state that indicates communication with the server is about to begin.
	KCTFontDescriptorMatchingWillBeginQuerying CTFontDescriptorMatchingState = 2
)

func (e CTFontDescriptorMatchingState) String() string {
	switch e {
	case KCTFontDescriptorMatchingDidBegin:
		return "KCTFontDescriptorMatchingDidBegin"
	case KCTFontDescriptorMatchingDidFailWithError:
		return "KCTFontDescriptorMatchingDidFailWithError"
	case KCTFontDescriptorMatchingDidFinish:
		return "KCTFontDescriptorMatchingDidFinish"
	case KCTFontDescriptorMatchingDidFinishDownloading:
		return "KCTFontDescriptorMatchingDidFinishDownloading"
	case KCTFontDescriptorMatchingDidMatch:
		return "KCTFontDescriptorMatchingDidMatch"
	case KCTFontDescriptorMatchingDownloading:
		return "KCTFontDescriptorMatchingDownloading"
	case KCTFontDescriptorMatchingStalled:
		return "KCTFontDescriptorMatchingStalled"
	case KCTFontDescriptorMatchingWillBeginDownloading:
		return "KCTFontDescriptorMatchingWillBeginDownloading"
	case KCTFontDescriptorMatchingWillBeginQuerying:
		return "KCTFontDescriptorMatchingWillBeginQuerying"
	default:
		return fmt.Sprintf("CTFontDescriptorMatchingState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTFontFormat
type CTFontFormat uint32

const (
	// KCTFontFormatBitmap: The font is a bitmap-only format.
	KCTFontFormatBitmap CTFontFormat = 5
	// KCTFontFormatOpenTypePostScript: The font is an OpenType format containing PostScript data.
	KCTFontFormatOpenTypePostScript CTFontFormat = 1
	// KCTFontFormatOpenTypeTrueType: The font is an OpenType format containing TrueType data.
	KCTFontFormatOpenTypeTrueType CTFontFormat = 2
	// KCTFontFormatPostScript: The font is a recognized PostScript format.
	KCTFontFormatPostScript CTFontFormat = 4
	// KCTFontFormatTrueType: The font is a recognized TrueType format.
	KCTFontFormatTrueType CTFontFormat = 3
	// KCTFontFormatUnrecognized: The font is not a recognized format.
	KCTFontFormatUnrecognized CTFontFormat = 0
)

func (e CTFontFormat) String() string {
	switch e {
	case KCTFontFormatBitmap:
		return "KCTFontFormatBitmap"
	case KCTFontFormatOpenTypePostScript:
		return "KCTFontFormatOpenTypePostScript"
	case KCTFontFormatOpenTypeTrueType:
		return "KCTFontFormatOpenTypeTrueType"
	case KCTFontFormatPostScript:
		return "KCTFontFormatPostScript"
	case KCTFontFormatTrueType:
		return "KCTFontFormatTrueType"
	case KCTFontFormatUnrecognized:
		return "KCTFontFormatUnrecognized"
	default:
		return fmt.Sprintf("CTFontFormat(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTFontManagerAutoActivationSetting
type CTFontManagerAutoActivationSetting uint32

const ()

// See: https://developer.apple.com/documentation/CoreText/CTFontManagerError
type CTFontManagerError int

const (
	// KCTFontManagerErrorAlreadyRegistered: An error that indicates the file is already registered in the specified scope.
	KCTFontManagerErrorAlreadyRegistered CTFontManagerError = 105
	// KCTFontManagerErrorAssetNotFound: An error that indicates the asset isn’t found.
	KCTFontManagerErrorAssetNotFound CTFontManagerError = 107
	// KCTFontManagerErrorCancelledByUser: An error that indicates the user cancelled the operation.
	KCTFontManagerErrorCancelledByUser CTFontManagerError = 304
	// KCTFontManagerErrorDuplicatedName: An error that indicates the file can’t register because of a duplicate font name.
	KCTFontManagerErrorDuplicatedName CTFontManagerError = 305
	// KCTFontManagerErrorExceededResourceLimit: An error that indicates an operation failure due to a system limitation.
	KCTFontManagerErrorExceededResourceLimit CTFontManagerError = 106
	// KCTFontManagerErrorFileNotFound: An error that indicates the file doesn’t exist at the specified URL.
	KCTFontManagerErrorFileNotFound CTFontManagerError = 101
	// KCTFontManagerErrorInUse: An error that indicates the font file is actively in use and can’t be unregistered.
	KCTFontManagerErrorInUse CTFontManagerError = 202
	// KCTFontManagerErrorInsufficientInfo: An error that indicates the font descriptor doesn’t have the necessary information to specify a font file.
	KCTFontManagerErrorInsufficientInfo CTFontManagerError = 303
	// KCTFontManagerErrorInsufficientPermissions: An error that indicates insufficient permissions to access the file.
	KCTFontManagerErrorInsufficientPermissions CTFontManagerError = 102
	// KCTFontManagerErrorInvalidFilePath: An error that indicates the file isn’t in an allowed location, which must be either in the app’s bundle or an on-demand resource.
	KCTFontManagerErrorInvalidFilePath CTFontManagerError = 306
	// KCTFontManagerErrorInvalidFontData: An error that indicates the file contains invalid font data that could cause system problems.
	KCTFontManagerErrorInvalidFontData CTFontManagerError = 104
	// KCTFontManagerErrorMissingEntitlement: An error that indicates the file can’t be processed because the provider doesn’t have a necessary entitlement.
	KCTFontManagerErrorMissingEntitlement CTFontManagerError = 302
	// KCTFontManagerErrorNotRegistered: An error that indicates the file isn’t registered in the specified scope.
	KCTFontManagerErrorNotRegistered CTFontManagerError = 201
	// KCTFontManagerErrorRegistrationFailed: An error that indicates the file can’t be processed due to an unexpected FontProvider error.
	KCTFontManagerErrorRegistrationFailed CTFontManagerError = 301
	// KCTFontManagerErrorSystemRequired: An error that indicates the file is required by the system and can’t be unregistered.
	KCTFontManagerErrorSystemRequired CTFontManagerError = 203
	// KCTFontManagerErrorUnrecognizedFormat: An error that indicates the file’s format is unrecognized or unsupported.
	KCTFontManagerErrorUnrecognizedFormat CTFontManagerError = 103
	// KCTFontManagerErrorUnsupportedScope: An error that indicates the specified scope isn’t supported.
	KCTFontManagerErrorUnsupportedScope CTFontManagerError = 307
)

func (e CTFontManagerError) String() string {
	switch e {
	case KCTFontManagerErrorAlreadyRegistered:
		return "KCTFontManagerErrorAlreadyRegistered"
	case KCTFontManagerErrorAssetNotFound:
		return "KCTFontManagerErrorAssetNotFound"
	case KCTFontManagerErrorCancelledByUser:
		return "KCTFontManagerErrorCancelledByUser"
	case KCTFontManagerErrorDuplicatedName:
		return "KCTFontManagerErrorDuplicatedName"
	case KCTFontManagerErrorExceededResourceLimit:
		return "KCTFontManagerErrorExceededResourceLimit"
	case KCTFontManagerErrorFileNotFound:
		return "KCTFontManagerErrorFileNotFound"
	case KCTFontManagerErrorInUse:
		return "KCTFontManagerErrorInUse"
	case KCTFontManagerErrorInsufficientInfo:
		return "KCTFontManagerErrorInsufficientInfo"
	case KCTFontManagerErrorInsufficientPermissions:
		return "KCTFontManagerErrorInsufficientPermissions"
	case KCTFontManagerErrorInvalidFilePath:
		return "KCTFontManagerErrorInvalidFilePath"
	case KCTFontManagerErrorInvalidFontData:
		return "KCTFontManagerErrorInvalidFontData"
	case KCTFontManagerErrorMissingEntitlement:
		return "KCTFontManagerErrorMissingEntitlement"
	case KCTFontManagerErrorNotRegistered:
		return "KCTFontManagerErrorNotRegistered"
	case KCTFontManagerErrorRegistrationFailed:
		return "KCTFontManagerErrorRegistrationFailed"
	case KCTFontManagerErrorSystemRequired:
		return "KCTFontManagerErrorSystemRequired"
	case KCTFontManagerErrorUnrecognizedFormat:
		return "KCTFontManagerErrorUnrecognizedFormat"
	case KCTFontManagerErrorUnsupportedScope:
		return "KCTFontManagerErrorUnsupportedScope"
	default:
		return fmt.Sprintf("CTFontManagerError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTFontManagerScope
type CTFontManagerScope uint32

const (
	// KCTFontManagerScopeNone: No scope is defined.
	KCTFontManagerScopeNone CTFontManagerScope = 0
	// KCTFontManagerScopePersistent: The font is available to all processes for the current user session and will be available in subsequent sessions unless unregistered.
	KCTFontManagerScopePersistent CTFontManagerScope = 2
	// KCTFontManagerScopeProcess: The font is available to the current process for the duration of the process unless directly unregistered.
	KCTFontManagerScopeProcess CTFontManagerScope = 1
	// KCTFontManagerScopeSession: The font is available to the current user session but won’t be available in subsequent sessions.
	KCTFontManagerScopeSession CTFontManagerScope = 3
)

func (e CTFontManagerScope) String() string {
	switch e {
	case KCTFontManagerScopeNone:
		return "KCTFontManagerScopeNone"
	case KCTFontManagerScopePersistent:
		return "KCTFontManagerScopePersistent"
	case KCTFontManagerScopeProcess:
		return "KCTFontManagerScopeProcess"
	case KCTFontManagerScopeSession:
		return "KCTFontManagerScopeSession"
	default:
		return fmt.Sprintf("CTFontManagerScope(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTFontOptions
type CTFontOptions uint

const (
	// KCTFontOptionsDefault: Default options are used.
	KCTFontOptionsDefault CTFontOptions = 0
	// KCTFontOptionsPreferSystemFont: Font matching prefers to match Apple system fonts.
	KCTFontOptionsPreferSystemFont CTFontOptions = 4
	// KCTFontOptionsPreventAutoActivation: Prevents automatic font activation.
	KCTFontOptionsPreventAutoActivation CTFontOptions = 1
	KCTFontOptionsPreventAutoDownload   CTFontOptions = 2
)

func (e CTFontOptions) String() string {
	switch e {
	case KCTFontOptionsDefault:
		return "KCTFontOptionsDefault"
	case KCTFontOptionsPreferSystemFont:
		return "KCTFontOptionsPreferSystemFont"
	case KCTFontOptionsPreventAutoActivation:
		return "KCTFontOptionsPreventAutoActivation"
	case KCTFontOptionsPreventAutoDownload:
		return "KCTFontOptionsPreventAutoDownload"
	default:
		return fmt.Sprintf("CTFontOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTFontOrientation
type CTFontOrientation uint32

const (
	// KCTFontOrientationDefault: The native orientation of the font.
	KCTFontOrientationDefault CTFontOrientation = 0
	// KCTFontOrientationHorizontal: The horizontal orientation.
	KCTFontOrientationHorizontal CTFontOrientation = 1
	// KCTFontOrientationVertical: The vertical orientation.
	KCTFontOrientationVertical CTFontOrientation = 2
)

func (e CTFontOrientation) String() string {
	switch e {
	case KCTFontOrientationDefault:
		return "KCTFontOrientationDefault"
	case KCTFontOrientationHorizontal:
		return "KCTFontOrientationHorizontal"
	case KCTFontOrientationVertical:
		return "KCTFontOrientationVertical"
	default:
		return fmt.Sprintf("CTFontOrientation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass
type CTFontStylisticClass uint32

const (
	// KCTFontClarendonSerifsClass: The font’s style is a variation of the Oldstyle Serifs and the Transitional Serifs.
	KCTFontClarendonSerifsClass CTFontStylisticClass = 1073741824
	// KCTFontClassClarendonSerifs: A font style variation of the Oldstyle Serifs and the Transitional Serifs.
	KCTFontClassClarendonSerifs CTFontStylisticClass = 1073741824
	// KCTFontClassFreeformSerifs: A font style that includes serifs but expresses a design freedom that doesn’t generally fit within the other serif design classifications.
	KCTFontClassFreeformSerifs CTFontStylisticClass = 1879048192
	// KCTFontClassModernSerifs: A font style based on the Latin printing style of the 20th century.
	KCTFontClassModernSerifs CTFontStylisticClass = 805306368
	// KCTFontClassOldStyleSerifs: A font style based on the Latin printing style of the 15th to 17th century.
	KCTFontClassOldStyleSerifs CTFontStylisticClass = 268435456
	// KCTFontClassOrnamentals: A font style that includes highly decorated or stylized character shapes such as those typically used in headlines.
	KCTFontClassOrnamentals CTFontStylisticClass = 2415919104
	// KCTFontClassSansSerif: A font style that includes most basic letter forms (excluding Scripts and Ornamentals) that do not have serifs on the strokes.
	KCTFontClassSansSerif CTFontStylisticClass = 2147483648
	// KCTFontClassScripts: A font style among those typefaces designed to simulate handwriting.
	KCTFontClassScripts CTFontStylisticClass = 2684354560
	// KCTFontClassSlabSerifs: A font style characterized by serifs with a square transition between the strokes and the serifs (no brackets).
	KCTFontClassSlabSerifs CTFontStylisticClass = 1342177280
	// KCTFontClassSymbolic: A generally design-independent font style.
	KCTFontClassSymbolic CTFontStylisticClass = 3221225472
	// KCTFontClassTransitionalSerifs: A font style based on the Latin printing style of the 18th to 19th century.
	KCTFontClassTransitionalSerifs CTFontStylisticClass = 536870912
	// KCTFontClassUnknown: The font has no design classification.
	KCTFontClassUnknown CTFontStylisticClass = 0
	// KCTFontFreeformSerifsClass: The font’s style includes serifs but expresses a design freedom that doesn’t generally fit within the other serif design classifications.
	KCTFontFreeformSerifsClass CTFontStylisticClass = 1879048192
	// KCTFontModernSerifsClass: The font’s style is based on the Latin printing style of the 20th century.
	KCTFontModernSerifsClass CTFontStylisticClass = 805306368
	// KCTFontOldStyleSerifsClass: The font’s style is based on the Latin printing style of the 15th to 17th century.
	KCTFontOldStyleSerifsClass CTFontStylisticClass = 268435456
	// KCTFontOrnamentalsClass: The font’s style includes highly decorated or stylized character shapes such as those typically used in headlines.
	KCTFontOrnamentalsClass CTFontStylisticClass = 2415919104
	// KCTFontSansSerifClass: The font’s style includes most basic letter forms (excluding Scripts and Ornamentals) that do not have serifs on the strokes.
	KCTFontSansSerifClass CTFontStylisticClass = 2147483648
	// KCTFontScriptsClass: The font’s style is among those typefaces designed to simulate handwriting.
	KCTFontScriptsClass CTFontStylisticClass = 2684354560
	// KCTFontSlabSerifsClass: The font’s style is characterized by serifs with a square transition between the strokes and the serifs (no brackets).
	KCTFontSlabSerifsClass CTFontStylisticClass = 1342177280
	// KCTFontSymbolicClass: The font’s style is generally design independent.
	KCTFontSymbolicClass CTFontStylisticClass = 3221225472
	// KCTFontTransitionalSerifsClass: The font’s style is based on the Latin printing style of the 18th to 19th century.
	KCTFontTransitionalSerifsClass CTFontStylisticClass = 536870912
	// KCTFontUnknownClass: The font has no design classification.
	KCTFontUnknownClass CTFontStylisticClass = 0
)

func (e CTFontStylisticClass) String() string {
	switch e {
	case KCTFontClarendonSerifsClass:
		return "KCTFontClarendonSerifsClass"
	case KCTFontClassFreeformSerifs:
		return "KCTFontClassFreeformSerifs"
	case KCTFontClassModernSerifs:
		return "KCTFontClassModernSerifs"
	case KCTFontClassOldStyleSerifs:
		return "KCTFontClassOldStyleSerifs"
	case KCTFontClassOrnamentals:
		return "KCTFontClassOrnamentals"
	case KCTFontClassSansSerif:
		return "KCTFontClassSansSerif"
	case KCTFontClassScripts:
		return "KCTFontClassScripts"
	case KCTFontClassSlabSerifs:
		return "KCTFontClassSlabSerifs"
	case KCTFontClassSymbolic:
		return "KCTFontClassSymbolic"
	case KCTFontClassTransitionalSerifs:
		return "KCTFontClassTransitionalSerifs"
	case KCTFontClassUnknown:
		return "KCTFontClassUnknown"
	default:
		return fmt.Sprintf("CTFontStylisticClass(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits
type CTFontSymbolicTraits uint32

const (
	// KCTFontBoldTrait: The font typestyle is boldface.
	KCTFontBoldTrait CTFontSymbolicTraits = 2
	// KCTFontClassMaskTrait: Mask for the font class.
	KCTFontClassMaskTrait CTFontSymbolicTraits = 4026531840
	// KCTFontColorGlyphsTrait: The font contains color glyphs.
	KCTFontColorGlyphsTrait CTFontSymbolicTraits = 8192
	// KCTFontCompositeTrait: The font is in Composite Font Reference format.
	KCTFontCompositeTrait CTFontSymbolicTraits = 16384
	// KCTFontCondensedTrait: The font typestyle is condensed.
	KCTFontCondensedTrait CTFontSymbolicTraits = 64
	// KCTFontExpandedTrait: The font typestyle is expanded.
	KCTFontExpandedTrait CTFontSymbolicTraits = 32
	// KCTFontItalicTrait: The font typestyle is italic.
	KCTFontItalicTrait CTFontSymbolicTraits = 1
	// KCTFontMonoSpaceTrait: The font uses fixed-pitch glyphs if available.
	KCTFontMonoSpaceTrait CTFontSymbolicTraits = 1024
	// KCTFontTraitBold: The font typestyle is boldface.
	KCTFontTraitBold CTFontSymbolicTraits = 2
	// KCTFontTraitClassMask: Mask for the font class.
	KCTFontTraitClassMask CTFontSymbolicTraits = 4026531840
	// KCTFontTraitColorGlyphs: The font contains color glyphs.
	KCTFontTraitColorGlyphs CTFontSymbolicTraits = 8192
	// KCTFontTraitComposite: The font is in Composite Font Reference format.
	KCTFontTraitComposite CTFontSymbolicTraits = 16384
	// KCTFontTraitCondensed: The font typestyle is condensed.
	KCTFontTraitCondensed CTFontSymbolicTraits = 64
	// KCTFontTraitExpanded: The font typestyle is expanded.
	KCTFontTraitExpanded CTFontSymbolicTraits = 32
	// KCTFontTraitItalic: The font typestyle is italic.
	KCTFontTraitItalic CTFontSymbolicTraits = 1
	// KCTFontTraitMonoSpace: The font uses fixed-pitch glyphs if available.
	KCTFontTraitMonoSpace CTFontSymbolicTraits = 1024
	// KCTFontTraitUIOptimized: The font synthesizes appropriate attributes for user interface rendering, such as control titles, if necessary.
	KCTFontTraitUIOptimized CTFontSymbolicTraits = 4096
	// KCTFontTraitVertical: The font uses vertical glyph variants and metrics.
	KCTFontTraitVertical CTFontSymbolicTraits = 2048
	// KCTFontUIOptimizedTrait: The font synthesizes appropriate attributes for user interface rendering, such as control titles, if necessary.
	KCTFontUIOptimizedTrait CTFontSymbolicTraits = 4096
	// KCTFontVerticalTrait: The font uses vertical glyph variants and metrics.
	KCTFontVerticalTrait CTFontSymbolicTraits = 2048
)

func (e CTFontSymbolicTraits) String() string {
	switch e {
	case KCTFontBoldTrait:
		return "KCTFontBoldTrait"
	case KCTFontClassMaskTrait:
		return "KCTFontClassMaskTrait"
	case KCTFontColorGlyphsTrait:
		return "KCTFontColorGlyphsTrait"
	case KCTFontCompositeTrait:
		return "KCTFontCompositeTrait"
	case KCTFontCondensedTrait:
		return "KCTFontCondensedTrait"
	case KCTFontExpandedTrait:
		return "KCTFontExpandedTrait"
	case KCTFontItalicTrait:
		return "KCTFontItalicTrait"
	case KCTFontMonoSpaceTrait:
		return "KCTFontMonoSpaceTrait"
	case KCTFontTraitUIOptimized:
		return "KCTFontTraitUIOptimized"
	case KCTFontTraitVertical:
		return "KCTFontTraitVertical"
	default:
		return fmt.Sprintf("CTFontSymbolicTraits(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTFontTableOptions
type CTFontTableOptions uint32

const (
	// KCTFontTableOptionNoOptions: No font table options are specified.
	KCTFontTableOptionNoOptions CTFontTableOptions = 0
	// Deprecated.
	KCTFontTableOptionExcludeSynthetic CTFontTableOptions = 1
)

func (e CTFontTableOptions) String() string {
	switch e {
	case KCTFontTableOptionNoOptions:
		return "KCTFontTableOptionNoOptions"
	case KCTFontTableOptionExcludeSynthetic:
		return "KCTFontTableOptionExcludeSynthetic"
	default:
		return fmt.Sprintf("CTFontTableOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTFontUIFontType
type CTFontUIFontType uint32

const (
	// KCTFontUIFontAlertHeader: The font for alert headers.
	KCTFontUIFontAlertHeader CTFontUIFontType = 18
	// KCTFontUIFontApplication: The default font for text documents.
	KCTFontUIFontApplication CTFontUIFontType = 9
	// KCTFontUIFontControlContent: The font for contents of user-interface controls.
	KCTFontUIFontControlContent CTFontUIFontType = 26
	// KCTFontUIFontEmphasizedSystem: The system font for emphasis in alerts.
	KCTFontUIFontEmphasizedSystem CTFontUIFontType = 3
	// KCTFontUIFontEmphasizedSystemDetail: The system font for emphasis in details.
	KCTFontUIFontEmphasizedSystemDetail CTFontUIFontType = 20
	// KCTFontUIFontLabel: The font for labels and tick marks on full-size sliders.
	KCTFontUIFontLabel CTFontUIFontType = 10
	// KCTFontUIFontMenuItem: The font for menu items.
	KCTFontUIFontMenuItem CTFontUIFontType = 12
	// KCTFontUIFontMenuItemCmdKey: The font for menu-item command-key equivalents.
	KCTFontUIFontMenuItemCmdKey CTFontUIFontType = 14
	// KCTFontUIFontMenuItemMark: The font to draw menu-item marks.
	KCTFontUIFontMenuItemMark CTFontUIFontType = 13
	// KCTFontUIFontMenuTitle: The font for menu titles.
	KCTFontUIFontMenuTitle CTFontUIFontType = 11
	// KCTFontUIFontMessage: The font for standard interface items, such as button labels and menu items.
	KCTFontUIFontMessage CTFontUIFontType = 23
	// KCTFontUIFontMiniEmphasizedSystem: The miniature system font for emphasis.
	KCTFontUIFontMiniEmphasizedSystem CTFontUIFontType = 7
	// KCTFontUIFontMiniSystem: The standard miniature system font for mini controls and utility window labels and text.
	KCTFontUIFontMiniSystem CTFontUIFontType = 6
	// KCTFontUIFontNone: The user-interface font type isn’t specified.
	KCTFontUIFontNone CTFontUIFontType = 4294967295
	// KCTFontUIFontPalette: The font in tool palettes.
	KCTFontUIFontPalette CTFontUIFontType = 24
	// KCTFontUIFontPushButton: The font for a push button, a rounded rectangular button with a text label on it.
	KCTFontUIFontPushButton CTFontUIFontType = 16
	// KCTFontUIFontSmallEmphasizedSystem: The small system font for emphasis.
	KCTFontUIFontSmallEmphasizedSystem CTFontUIFontType = 5
	// KCTFontUIFontSmallSystem: The standard small system font for informative text in alerts, column headings in lists, help tags, and small controls.
	KCTFontUIFontSmallSystem CTFontUIFontType = 4
	// KCTFontUIFontSmallToolbar: The small font for labels of toolbar items.
	KCTFontUIFontSmallToolbar CTFontUIFontType = 22
	// KCTFontUIFontSystem: The system font for standard user-interface items, such as button labels and menu items.
	KCTFontUIFontSystem CTFontUIFontType = 2
	// KCTFontUIFontSystemDetail: The standard system font for details.
	KCTFontUIFontSystemDetail CTFontUIFontType = 19
	// KCTFontUIFontToolTip: The font for tool tips.
	KCTFontUIFontToolTip CTFontUIFontType = 25
	// KCTFontUIFontToolbar: The font used for labels of toolbar items.
	KCTFontUIFontToolbar CTFontUIFontType = 21
	// KCTFontUIFontUser: The default font for documents and other text whose font the user can typically change.
	KCTFontUIFontUser CTFontUIFontType = 0
	// KCTFontUIFontUserFixedPitch: The default font for documents and other text under the user’s control when that font is fixed-pitch.
	KCTFontUIFontUserFixedPitch CTFontUIFontType = 1
	// KCTFontUIFontUtilityWindowTitle: The font for utility window titles.
	KCTFontUIFontUtilityWindowTitle CTFontUIFontType = 17
	// KCTFontUIFontViews: The default view font for text in lists and tables.
	KCTFontUIFontViews CTFontUIFontType = 8
	// KCTFontUIFontWindowTitle: The font for window titles.
	KCTFontUIFontWindowTitle CTFontUIFontType = 15
)

func (e CTFontUIFontType) String() string {
	switch e {
	case KCTFontUIFontAlertHeader:
		return "KCTFontUIFontAlertHeader"
	case KCTFontUIFontApplication:
		return "KCTFontUIFontApplication"
	case KCTFontUIFontControlContent:
		return "KCTFontUIFontControlContent"
	case KCTFontUIFontEmphasizedSystem:
		return "KCTFontUIFontEmphasizedSystem"
	case KCTFontUIFontEmphasizedSystemDetail:
		return "KCTFontUIFontEmphasizedSystemDetail"
	case KCTFontUIFontLabel:
		return "KCTFontUIFontLabel"
	case KCTFontUIFontMenuItem:
		return "KCTFontUIFontMenuItem"
	case KCTFontUIFontMenuItemCmdKey:
		return "KCTFontUIFontMenuItemCmdKey"
	case KCTFontUIFontMenuItemMark:
		return "KCTFontUIFontMenuItemMark"
	case KCTFontUIFontMenuTitle:
		return "KCTFontUIFontMenuTitle"
	case KCTFontUIFontMessage:
		return "KCTFontUIFontMessage"
	case KCTFontUIFontMiniEmphasizedSystem:
		return "KCTFontUIFontMiniEmphasizedSystem"
	case KCTFontUIFontMiniSystem:
		return "KCTFontUIFontMiniSystem"
	case KCTFontUIFontNone:
		return "KCTFontUIFontNone"
	case KCTFontUIFontPalette:
		return "KCTFontUIFontPalette"
	case KCTFontUIFontPushButton:
		return "KCTFontUIFontPushButton"
	case KCTFontUIFontSmallEmphasizedSystem:
		return "KCTFontUIFontSmallEmphasizedSystem"
	case KCTFontUIFontSmallSystem:
		return "KCTFontUIFontSmallSystem"
	case KCTFontUIFontSmallToolbar:
		return "KCTFontUIFontSmallToolbar"
	case KCTFontUIFontSystem:
		return "KCTFontUIFontSystem"
	case KCTFontUIFontSystemDetail:
		return "KCTFontUIFontSystemDetail"
	case KCTFontUIFontToolTip:
		return "KCTFontUIFontToolTip"
	case KCTFontUIFontToolbar:
		return "KCTFontUIFontToolbar"
	case KCTFontUIFontUser:
		return "KCTFontUIFontUser"
	case KCTFontUIFontUserFixedPitch:
		return "KCTFontUIFontUserFixedPitch"
	case KCTFontUIFontUtilityWindowTitle:
		return "KCTFontUIFontUtilityWindowTitle"
	case KCTFontUIFontViews:
		return "KCTFontUIFontViews"
	case KCTFontUIFontWindowTitle:
		return "KCTFontUIFontWindowTitle"
	default:
		return fmt.Sprintf("CTFontUIFontType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTFramePathFillRule
type CTFramePathFillRule uint32

const (
	// KCTFramePathFillEvenOdd: Paints the area using the even-odd fill rule.
	KCTFramePathFillEvenOdd CTFramePathFillRule = 0
	// KCTFramePathFillWindingNumber: Paints the area using the nonzero winding number rule.
	KCTFramePathFillWindingNumber CTFramePathFillRule = 1
)

func (e CTFramePathFillRule) String() string {
	switch e {
	case KCTFramePathFillEvenOdd:
		return "KCTFramePathFillEvenOdd"
	case KCTFramePathFillWindingNumber:
		return "KCTFramePathFillWindingNumber"
	default:
		return fmt.Sprintf("CTFramePathFillRule(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTFrameProgression
type CTFrameProgression uint32

const (
	// KCTFrameProgressionLeftToRight: Lines stack left to right for vertical text.
	KCTFrameProgressionLeftToRight CTFrameProgression = 2
	// KCTFrameProgressionRightToLeft: Lines stack right to left for vertical text.
	KCTFrameProgressionRightToLeft CTFrameProgression = 1
	// KCTFrameProgressionTopToBottom: Lines stack top to bottom for horizontal text.
	KCTFrameProgressionTopToBottom CTFrameProgression = 0
)

func (e CTFrameProgression) String() string {
	switch e {
	case KCTFrameProgressionLeftToRight:
		return "KCTFrameProgressionLeftToRight"
	case KCTFrameProgressionRightToLeft:
		return "KCTFrameProgressionRightToLeft"
	case KCTFrameProgressionTopToBottom:
		return "KCTFrameProgressionTopToBottom"
	default:
		return fmt.Sprintf("CTFrameProgression(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTLineBoundsOptions
type CTLineBoundsOptions uint

const (
	// KCTLineBoundsExcludeTypographicLeading: An option to exclude typographic leading.
	KCTLineBoundsExcludeTypographicLeading CTLineBoundsOptions = 1
	// KCTLineBoundsExcludeTypographicShifts: An option to ignore cross-stream shifts due to positioning, such as kerning or baseline alignment.
	KCTLineBoundsExcludeTypographicShifts CTLineBoundsOptions = 2
	// KCTLineBoundsIncludeLanguageExtents: An option to include additional space based on common glyph sequences for various languages.
	KCTLineBoundsIncludeLanguageExtents CTLineBoundsOptions = 32
	// KCTLineBoundsUseGlyphPathBounds: An option to use glyph path bounds rather than the default typographic bounds.
	KCTLineBoundsUseGlyphPathBounds CTLineBoundsOptions = 8
	// KCTLineBoundsUseHangingPunctuation: An option to enable hanging punctuation.
	KCTLineBoundsUseHangingPunctuation CTLineBoundsOptions = 4
	// KCTLineBoundsUseOpticalBounds: An option to use optical bounds.
	KCTLineBoundsUseOpticalBounds CTLineBoundsOptions = 16
)

func (e CTLineBoundsOptions) String() string {
	switch e {
	case KCTLineBoundsExcludeTypographicLeading:
		return "KCTLineBoundsExcludeTypographicLeading"
	case KCTLineBoundsExcludeTypographicShifts:
		return "KCTLineBoundsExcludeTypographicShifts"
	case KCTLineBoundsIncludeLanguageExtents:
		return "KCTLineBoundsIncludeLanguageExtents"
	case KCTLineBoundsUseGlyphPathBounds:
		return "KCTLineBoundsUseGlyphPathBounds"
	case KCTLineBoundsUseHangingPunctuation:
		return "KCTLineBoundsUseHangingPunctuation"
	case KCTLineBoundsUseOpticalBounds:
		return "KCTLineBoundsUseOpticalBounds"
	default:
		return fmt.Sprintf("CTLineBoundsOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTLineBreakMode
type CTLineBreakMode uint8

const ()

// See: https://developer.apple.com/documentation/CoreText/CTLineTruncationType
type CTLineTruncationType uint32

const ()

// See: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier
type CTParagraphStyleSpecifier uint32

const (
	// KCTParagraphStyleSpecifierAlignment: The text alignment.
	KCTParagraphStyleSpecifierAlignment CTParagraphStyleSpecifier = 0
	// KCTParagraphStyleSpecifierBaseWritingDirection: The base writing direction of the lines.
	KCTParagraphStyleSpecifierBaseWritingDirection CTParagraphStyleSpecifier = 13
	// KCTParagraphStyleSpecifierCount: The number of style specifiers.
	KCTParagraphStyleSpecifierCount CTParagraphStyleSpecifier = 18
	// KCTParagraphStyleSpecifierDefaultTabInterval: The document-wide default tab interval.
	KCTParagraphStyleSpecifierDefaultTabInterval CTParagraphStyleSpecifier = 5
	// KCTParagraphStyleSpecifierFirstLineHeadIndent: The distance, in points, from the leading margin of a frame to the beginning of the paragraph’s first line.
	KCTParagraphStyleSpecifierFirstLineHeadIndent CTParagraphStyleSpecifier = 1
	// KCTParagraphStyleSpecifierHeadIndent: The distance, in points, from the leading margin of a text container to the beginning of lines other than the first.
	KCTParagraphStyleSpecifierHeadIndent CTParagraphStyleSpecifier = 2
	// KCTParagraphStyleSpecifierLineBoundsOptions: Options that control the alignment of the line edges with the leading and trailing margins.
	KCTParagraphStyleSpecifierLineBoundsOptions CTParagraphStyleSpecifier = 17
	// KCTParagraphStyleSpecifierLineBreakMode: The mode that should be used to break lines when laying out the paragraph’s text.
	KCTParagraphStyleSpecifierLineBreakMode CTParagraphStyleSpecifier = 6
	// KCTParagraphStyleSpecifierLineHeightMultiple: The line height multiple.
	KCTParagraphStyleSpecifierLineHeightMultiple CTParagraphStyleSpecifier = 7
	// KCTParagraphStyleSpecifierLineSpacing: The space in points added between lines within the paragraph (commonly known as leading).
	KCTParagraphStyleSpecifierLineSpacing CTParagraphStyleSpecifier = 10
	// KCTParagraphStyleSpecifierLineSpacingAdjustment: The space in points added between lines within the paragraph (commonly known as leading).
	KCTParagraphStyleSpecifierLineSpacingAdjustment CTParagraphStyleSpecifier = 16
	// KCTParagraphStyleSpecifierMaximumLineHeight: The maximum height that any line in the frame will occupy, regardless of the font size or size of any attached graphic.
	KCTParagraphStyleSpecifierMaximumLineHeight CTParagraphStyleSpecifier = 8
	// KCTParagraphStyleSpecifierMaximumLineSpacing: The maximum space in points between lines within the paragraph (commonly known as leading).
	KCTParagraphStyleSpecifierMaximumLineSpacing CTParagraphStyleSpecifier = 14
	// KCTParagraphStyleSpecifierMinimumLineHeight: The minimum height that any line in the frame will occupy, regardless of the font size or size of any attached graphic.
	KCTParagraphStyleSpecifierMinimumLineHeight CTParagraphStyleSpecifier = 9
	// KCTParagraphStyleSpecifierMinimumLineSpacing: The minimum space in points between lines within the paragraph (commonly known as leading).
	KCTParagraphStyleSpecifierMinimumLineSpacing CTParagraphStyleSpecifier = 15
	// KCTParagraphStyleSpecifierParagraphSpacing: The space added at the end of the paragraph to separate it from the following paragraph.
	KCTParagraphStyleSpecifierParagraphSpacing CTParagraphStyleSpecifier = 11
	// KCTParagraphStyleSpecifierParagraphSpacingBefore: The distance between the paragraph’s top and the beginning of its text content.
	KCTParagraphStyleSpecifierParagraphSpacingBefore CTParagraphStyleSpecifier = 12
	// KCTParagraphStyleSpecifierTabStops: The text tab objects, sorted by location, that define the tab stops for the paragraph style.
	KCTParagraphStyleSpecifierTabStops CTParagraphStyleSpecifier = 4
	// KCTParagraphStyleSpecifierTailIndent: The distance, in points, from the margin of a frame to the end of lines.
	KCTParagraphStyleSpecifierTailIndent CTParagraphStyleSpecifier = 3
)

func (e CTParagraphStyleSpecifier) String() string {
	switch e {
	case KCTParagraphStyleSpecifierAlignment:
		return "KCTParagraphStyleSpecifierAlignment"
	case KCTParagraphStyleSpecifierBaseWritingDirection:
		return "KCTParagraphStyleSpecifierBaseWritingDirection"
	case KCTParagraphStyleSpecifierCount:
		return "KCTParagraphStyleSpecifierCount"
	case KCTParagraphStyleSpecifierDefaultTabInterval:
		return "KCTParagraphStyleSpecifierDefaultTabInterval"
	case KCTParagraphStyleSpecifierFirstLineHeadIndent:
		return "KCTParagraphStyleSpecifierFirstLineHeadIndent"
	case KCTParagraphStyleSpecifierHeadIndent:
		return "KCTParagraphStyleSpecifierHeadIndent"
	case KCTParagraphStyleSpecifierLineBoundsOptions:
		return "KCTParagraphStyleSpecifierLineBoundsOptions"
	case KCTParagraphStyleSpecifierLineBreakMode:
		return "KCTParagraphStyleSpecifierLineBreakMode"
	case KCTParagraphStyleSpecifierLineHeightMultiple:
		return "KCTParagraphStyleSpecifierLineHeightMultiple"
	case KCTParagraphStyleSpecifierLineSpacing:
		return "KCTParagraphStyleSpecifierLineSpacing"
	case KCTParagraphStyleSpecifierLineSpacingAdjustment:
		return "KCTParagraphStyleSpecifierLineSpacingAdjustment"
	case KCTParagraphStyleSpecifierMaximumLineHeight:
		return "KCTParagraphStyleSpecifierMaximumLineHeight"
	case KCTParagraphStyleSpecifierMaximumLineSpacing:
		return "KCTParagraphStyleSpecifierMaximumLineSpacing"
	case KCTParagraphStyleSpecifierMinimumLineHeight:
		return "KCTParagraphStyleSpecifierMinimumLineHeight"
	case KCTParagraphStyleSpecifierMinimumLineSpacing:
		return "KCTParagraphStyleSpecifierMinimumLineSpacing"
	case KCTParagraphStyleSpecifierParagraphSpacing:
		return "KCTParagraphStyleSpecifierParagraphSpacing"
	case KCTParagraphStyleSpecifierParagraphSpacingBefore:
		return "KCTParagraphStyleSpecifierParagraphSpacingBefore"
	case KCTParagraphStyleSpecifierTabStops:
		return "KCTParagraphStyleSpecifierTabStops"
	case KCTParagraphStyleSpecifierTailIndent:
		return "KCTParagraphStyleSpecifierTailIndent"
	default:
		return fmt.Sprintf("CTParagraphStyleSpecifier(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTRubyAlignment
type CTRubyAlignment uint8

const (
	// KCTRubyAlignmentAuto: Core Text automatically determines the alignment.
	KCTRubyAlignmentAuto CTRubyAlignment = 0
	// KCTRubyAlignmentCenter: Centers the ruby text within the width of the base text.
	KCTRubyAlignmentCenter CTRubyAlignment = 2
	// KCTRubyAlignmentDistributeLetter: Distributes the ruby text evenly over the width of the base text, aligning the first and last characters of the ruby text with the first and last characters of the base text.
	KCTRubyAlignmentDistributeLetter CTRubyAlignment = 4
	// KCTRubyAlignmentDistributeSpace: Distributes the ruby text evenly over the width of the base text, adding space before the first and after the last character.
	KCTRubyAlignmentDistributeSpace CTRubyAlignment = 5
	// KCTRubyAlignmentEnd: Aligns the ruby text with the ending edge of the base text.
	KCTRubyAlignmentEnd CTRubyAlignment = 3
	// KCTRubyAlignmentInvalid: The alignment is invalid.
	KCTRubyAlignmentInvalid CTRubyAlignment = 255
	// KCTRubyAlignmentLineEdge: Aligns the ruby text to an adjacent line edge.
	KCTRubyAlignmentLineEdge CTRubyAlignment = 6
	// KCTRubyAlignmentStart: Aligns the ruby text with the starting edge of the base text.
	KCTRubyAlignmentStart CTRubyAlignment = 1
)

func (e CTRubyAlignment) String() string {
	switch e {
	case KCTRubyAlignmentAuto:
		return "KCTRubyAlignmentAuto"
	case KCTRubyAlignmentCenter:
		return "KCTRubyAlignmentCenter"
	case KCTRubyAlignmentDistributeLetter:
		return "KCTRubyAlignmentDistributeLetter"
	case KCTRubyAlignmentDistributeSpace:
		return "KCTRubyAlignmentDistributeSpace"
	case KCTRubyAlignmentEnd:
		return "KCTRubyAlignmentEnd"
	case KCTRubyAlignmentInvalid:
		return "KCTRubyAlignmentInvalid"
	case KCTRubyAlignmentLineEdge:
		return "KCTRubyAlignmentLineEdge"
	case KCTRubyAlignmentStart:
		return "KCTRubyAlignmentStart"
	default:
		return fmt.Sprintf("CTRubyAlignment(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTRubyOverhang
type CTRubyOverhang uint8

const (
	// KCTRubyOverhangAuto: The ruby text can overhang adjacent text on both sides.
	KCTRubyOverhangAuto CTRubyOverhang = 0
	// KCTRubyOverhangEnd: The ruby text can overhang the text that follows it.
	KCTRubyOverhangEnd CTRubyOverhang = 2
	// KCTRubyOverhangInvalid: The overhang specification is invalid.
	KCTRubyOverhangInvalid CTRubyOverhang = 255
	// KCTRubyOverhangNone: The ruby text can’t overhang the preceding or following text.
	KCTRubyOverhangNone CTRubyOverhang = 3
	// KCTRubyOverhangStart: The ruby text can overhang the text that precedes it.
	KCTRubyOverhangStart CTRubyOverhang = 1
)

func (e CTRubyOverhang) String() string {
	switch e {
	case KCTRubyOverhangAuto:
		return "KCTRubyOverhangAuto"
	case KCTRubyOverhangEnd:
		return "KCTRubyOverhangEnd"
	case KCTRubyOverhangInvalid:
		return "KCTRubyOverhangInvalid"
	case KCTRubyOverhangNone:
		return "KCTRubyOverhangNone"
	case KCTRubyOverhangStart:
		return "KCTRubyOverhangStart"
	default:
		return fmt.Sprintf("CTRubyOverhang(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTRubyPosition
type CTRubyPosition uint8

const (
	// KCTRubyPositionAfter: The ruby text is positioned after the base text, appearing below horizontal text and to the left of vertical text.
	KCTRubyPositionAfter CTRubyPosition = 1
	// KCTRubyPositionBefore: The ruby text is positioned before the base text, appearing above horizontal text and to the right of vertical text.
	KCTRubyPositionBefore CTRubyPosition = 0
	// KCTRubyPositionCount: A constant that accounts for all ruby positions during ruby annotation creation.
	KCTRubyPositionCount CTRubyPosition = 4
	// KCTRubyPositionInline: The ruby text follows the base text with no special styling.
	KCTRubyPositionInline CTRubyPosition = 3
	// KCTRubyPositionInterCharacter: The ruby text is positioned to the right of the base text, regardless of whether it’s horizontal or vertical.
	KCTRubyPositionInterCharacter CTRubyPosition = 2
)

func (e CTRubyPosition) String() string {
	switch e {
	case KCTRubyPositionAfter:
		return "KCTRubyPositionAfter"
	case KCTRubyPositionBefore:
		return "KCTRubyPositionBefore"
	case KCTRubyPositionCount:
		return "KCTRubyPositionCount"
	case KCTRubyPositionInline:
		return "KCTRubyPositionInline"
	case KCTRubyPositionInterCharacter:
		return "KCTRubyPositionInterCharacter"
	default:
		return fmt.Sprintf("CTRubyPosition(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTRunStatus
type CTRunStatus uint32

const (
	// KCTRunStatusHasNonIdentityMatrix: The run requires a specific text matrix to be set in the current Core Graphics context for proper drawing.
	KCTRunStatusHasNonIdentityMatrix CTRunStatus = 4
	// KCTRunStatusNoStatus: The run has no special attributes.
	KCTRunStatusNoStatus CTRunStatus = 0
	// KCTRunStatusNonMonotonic: The run isn’t in strictly increasing or decreasing order.
	KCTRunStatusNonMonotonic CTRunStatus = 2
	// KCTRunStatusRightToLeft: The run proceeds from right to left.
	KCTRunStatusRightToLeft CTRunStatus = 1
)

func (e CTRunStatus) String() string {
	switch e {
	case KCTRunStatusHasNonIdentityMatrix:
		return "KCTRunStatusHasNonIdentityMatrix"
	case KCTRunStatusNoStatus:
		return "KCTRunStatusNoStatus"
	case KCTRunStatusNonMonotonic:
		return "KCTRunStatusNonMonotonic"
	case KCTRunStatusRightToLeft:
		return "KCTRunStatusRightToLeft"
	default:
		return fmt.Sprintf("CTRunStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTTextAlignment
type CTTextAlignment uint8

const (
	// KCTTextAlignmentCenter: Text is visually center-aligned.
	KCTTextAlignmentCenter CTTextAlignment = 2
	// KCTTextAlignmentJustified: Text is fully justified.
	KCTTextAlignmentJustified CTTextAlignment = 3
	// KCTTextAlignmentLeft: Text is visually left-aligned.
	KCTTextAlignmentLeft CTTextAlignment = 0
	// KCTTextAlignmentNatural: Text uses the natural alignment of the text’s script.
	KCTTextAlignmentNatural CTTextAlignment = 4
	// KCTTextAlignmentRight: Text is visually right-aligned.
	KCTTextAlignmentRight CTTextAlignment = 1
)

func (e CTTextAlignment) String() string {
	switch e {
	case KCTTextAlignmentCenter:
		return "KCTTextAlignmentCenter"
	case KCTTextAlignmentJustified:
		return "KCTTextAlignmentJustified"
	case KCTTextAlignmentLeft:
		return "KCTTextAlignmentLeft"
	case KCTTextAlignmentNatural:
		return "KCTTextAlignmentNatural"
	case KCTTextAlignmentRight:
		return "KCTTextAlignmentRight"
	default:
		return fmt.Sprintf("CTTextAlignment(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTUnderlineStyle
type CTUnderlineStyle int32

const (
	// KCTUnderlineStyleDouble: A specifier that indicates to draw an underline consisting of a double line.
	KCTUnderlineStyleDouble CTUnderlineStyle = 0x9
	// KCTUnderlineStyleNone: A specifier that indicates not to draw an underline.
	KCTUnderlineStyleNone CTUnderlineStyle = 0
	// KCTUnderlineStyleSingle: A specifier that indicates to draw an underline consisting of a single line.
	KCTUnderlineStyleSingle CTUnderlineStyle = 0x1
	// KCTUnderlineStyleThick: A specifier that indicates to draw an underline consisting of a thick line.
	KCTUnderlineStyleThick CTUnderlineStyle = 0x2
)

func (e CTUnderlineStyle) String() string {
	switch e {
	case KCTUnderlineStyleDouble:
		return "KCTUnderlineStyleDouble"
	case KCTUnderlineStyleNone:
		return "KCTUnderlineStyleNone"
	case KCTUnderlineStyleSingle:
		return "KCTUnderlineStyleSingle"
	case KCTUnderlineStyleThick:
		return "KCTUnderlineStyleThick"
	default:
		return fmt.Sprintf("CTUnderlineStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTUnderlineStyleModifiers
type CTUnderlineStyleModifiers int32

const (
	// KCTUnderlinePatternDash: A modifier that indicates to draw an underline using a pattern of dashes.
	KCTUnderlinePatternDash CTUnderlineStyleModifiers = 0x200
	// KCTUnderlinePatternDashDot: A modifier that indicates to draw an underline using a pattern of alternating dashes and dots.
	KCTUnderlinePatternDashDot CTUnderlineStyleModifiers = 0x300
	// KCTUnderlinePatternDashDotDot: A modifier that indicates to draw an underline using a pattern of a dash followed by two dots.
	KCTUnderlinePatternDashDotDot CTUnderlineStyleModifiers = 0x400
	// KCTUnderlinePatternDot: A modifier that indicates to draw an underline using a pattern of dots.
	KCTUnderlinePatternDot CTUnderlineStyleModifiers = 0x100
	// KCTUnderlinePatternSolid: A modifier that indicates to draw a solid underline.
	KCTUnderlinePatternSolid CTUnderlineStyleModifiers = 0
)

func (e CTUnderlineStyleModifiers) String() string {
	switch e {
	case KCTUnderlinePatternDash:
		return "KCTUnderlinePatternDash"
	case KCTUnderlinePatternDashDot:
		return "KCTUnderlinePatternDashDot"
	case KCTUnderlinePatternDashDotDot:
		return "KCTUnderlinePatternDashDotDot"
	case KCTUnderlinePatternDot:
		return "KCTUnderlinePatternDot"
	case KCTUnderlinePatternSolid:
		return "KCTUnderlinePatternSolid"
	default:
		return fmt.Sprintf("CTUnderlineStyleModifiers(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreText/CTWritingDirection
type CTWritingDirection int8

const ()

const CmapFontTableTag uint32 = 0x636d6170

const DescriptorFontTableTag uint32 = 0x66647363

const FeatureFontTableTag uint32 = 0x66656174

const KANKRCurrentVersion uint32 = 0

type KAllTypeFeatures uint32

const (
	KAllTypeFeaturesOffSelector KAllTypeFeatures = 1
	KAllTypeFeaturesOnSelector  KAllTypeFeatures = 0
)

func (e KAllTypeFeatures) String() string {
	switch e {
	case KAllTypeFeaturesOffSelector:
		return "KAllTypeFeaturesOffSelector"
	case KAllTypeFeaturesOnSelector:
		return "KAllTypeFeaturesOnSelector"
	default:
		return fmt.Sprintf("KAllTypeFeatures(%d)", e)
	}
}

type KAllTypographicFeaturesType int32

const (
	KAllTypographicFeaturesTypeValue KAllTypographicFeaturesType = 0
	KAlternateKanaType               KAllTypographicFeaturesType = 34
	KAnnotationType                  KAllTypographicFeaturesType = 24
	KCJKRomanSpacingType             KAllTypographicFeaturesType = 103
	KCJKSymbolAlternativesType       KAllTypographicFeaturesType = 29
	KCJKVerticalRomanPlacementType   KAllTypographicFeaturesType = 31
	KCaseSensitiveLayoutType         KAllTypographicFeaturesType = 33
	KCharacterAlternativesType       KAllTypographicFeaturesType = 17
	KCharacterShapeType              KAllTypographicFeaturesType = 20
	KContextualAlternatesType        KAllTypographicFeaturesType = 36
	KCursiveConnectionType           KAllTypographicFeaturesType = 2
	KDesignComplexityType            KAllTypographicFeaturesType = 18
	KDiacriticsType                  KAllTypographicFeaturesType = 9
	KFractionsType                   KAllTypographicFeaturesType = 11
	KIdeographicAlternativesType     KAllTypographicFeaturesType = 30
	KIdeographicSpacingType          KAllTypographicFeaturesType = 26
	KItalicCJKRomanType              KAllTypographicFeaturesType = 32
	KKanaSpacingType                 KAllTypographicFeaturesType = 25
	KLanguageTagType                 KAllTypographicFeaturesType = 39
	KLastFeatureType                 KAllTypographicFeaturesType = -1
	KLetterCaseType                  KAllTypographicFeaturesType = 3
	KLigaturesType                   KAllTypographicFeaturesType = 1
	KLinguisticRearrangementType     KAllTypographicFeaturesType = 5
	KLowerCaseType                   KAllTypographicFeaturesType = 37
	KMathematicalExtrasType          KAllTypographicFeaturesType = 15
	KNumberCaseType                  KAllTypographicFeaturesType = 21
	KNumberSpacingType               KAllTypographicFeaturesType = 6
	KOrnamentSetsType                KAllTypographicFeaturesType = 16
	KOverlappingCharactersType       KAllTypographicFeaturesType = 13
	KRubyKanaType                    KAllTypographicFeaturesType = 28
	KSmartSwashType                  KAllTypographicFeaturesType = 8
	KStyleOptionsType                KAllTypographicFeaturesType = 19
	KStylisticAlternativesType       KAllTypographicFeaturesType = 35
	KTextSpacingType                 KAllTypographicFeaturesType = 22
	KTransliterationType             KAllTypographicFeaturesType = 23
	KTypographicExtrasType           KAllTypographicFeaturesType = 14
	KUnicodeDecompositionType        KAllTypographicFeaturesType = 27
	KUpperCaseType                   KAllTypographicFeaturesType = 38
	KVerticalPositionType            KAllTypographicFeaturesType = 10
	KVerticalSubstitutionType        KAllTypographicFeaturesType = 4
)

func (e KAllTypographicFeaturesType) String() string {
	switch e {
	case KAllTypographicFeaturesTypeValue:
		return "KAllTypographicFeaturesTypeValue"
	case KAlternateKanaType:
		return "KAlternateKanaType"
	case KAnnotationType:
		return "KAnnotationType"
	case KCJKRomanSpacingType:
		return "KCJKRomanSpacingType"
	case KCJKSymbolAlternativesType:
		return "KCJKSymbolAlternativesType"
	case KCJKVerticalRomanPlacementType:
		return "KCJKVerticalRomanPlacementType"
	case KCaseSensitiveLayoutType:
		return "KCaseSensitiveLayoutType"
	case KCharacterAlternativesType:
		return "KCharacterAlternativesType"
	case KCharacterShapeType:
		return "KCharacterShapeType"
	case KContextualAlternatesType:
		return "KContextualAlternatesType"
	case KCursiveConnectionType:
		return "KCursiveConnectionType"
	case KDesignComplexityType:
		return "KDesignComplexityType"
	case KDiacriticsType:
		return "KDiacriticsType"
	case KFractionsType:
		return "KFractionsType"
	case KIdeographicAlternativesType:
		return "KIdeographicAlternativesType"
	case KIdeographicSpacingType:
		return "KIdeographicSpacingType"
	case KItalicCJKRomanType:
		return "KItalicCJKRomanType"
	case KKanaSpacingType:
		return "KKanaSpacingType"
	case KLanguageTagType:
		return "KLanguageTagType"
	case KLastFeatureType:
		return "KLastFeatureType"
	case KLetterCaseType:
		return "KLetterCaseType"
	case KLigaturesType:
		return "KLigaturesType"
	case KLinguisticRearrangementType:
		return "KLinguisticRearrangementType"
	case KLowerCaseType:
		return "KLowerCaseType"
	case KMathematicalExtrasType:
		return "KMathematicalExtrasType"
	case KNumberCaseType:
		return "KNumberCaseType"
	case KNumberSpacingType:
		return "KNumberSpacingType"
	case KOrnamentSetsType:
		return "KOrnamentSetsType"
	case KOverlappingCharactersType:
		return "KOverlappingCharactersType"
	case KRubyKanaType:
		return "KRubyKanaType"
	case KSmartSwashType:
		return "KSmartSwashType"
	case KStyleOptionsType:
		return "KStyleOptionsType"
	case KStylisticAlternativesType:
		return "KStylisticAlternativesType"
	case KTextSpacingType:
		return "KTextSpacingType"
	case KTransliterationType:
		return "KTransliterationType"
	case KTypographicExtrasType:
		return "KTypographicExtrasType"
	case KUnicodeDecompositionType:
		return "KUnicodeDecompositionType"
	case KUpperCaseType:
		return "KUpperCaseType"
	case KVerticalPositionType:
		return "KVerticalPositionType"
	case KVerticalSubstitutionType:
		return "KVerticalSubstitutionType"
	default:
		return fmt.Sprintf("KAllTypographicFeaturesType(%d)", e)
	}
}

type KAlternate uint32

const (
	KAlternateHorizKanaOffSelector KAlternate = 1
	KAlternateHorizKanaOnSelector  KAlternate = 0
	KAlternateVertKanaOffSelector  KAlternate = 3
	KAlternateVertKanaOnSelector   KAlternate = 2
)

func (e KAlternate) String() string {
	switch e {
	case KAlternateHorizKanaOffSelector:
		return "KAlternateHorizKanaOffSelector"
	case KAlternateHorizKanaOnSelector:
		return "KAlternateHorizKanaOnSelector"
	case KAlternateVertKanaOffSelector:
		return "KAlternateVertKanaOffSelector"
	case KAlternateVertKanaOnSelector:
		return "KAlternateVertKanaOnSelector"
	default:
		return fmt.Sprintf("KAlternate(%d)", e)
	}
}

type KBSLNRomanBaseline uint32

const (
	KBSLNHangingBaseline           KBSLNRomanBaseline = 3
	KBSLNIdeographicCenterBaseline KBSLNRomanBaseline = 1
	KBSLNIdeographicHighBaseline   KBSLNRomanBaseline = 5
	KBSLNIdeographicLowBaseline    KBSLNRomanBaseline = 2
	KBSLNLastBaseline              KBSLNRomanBaseline = 31
	KBSLNMathBaseline              KBSLNRomanBaseline = 4
	KBSLNNoBaseline                KBSLNRomanBaseline = 255
	KBSLNNoBaselineOverride        KBSLNRomanBaseline = 255
	KBSLNNumBaselineClasses        KBSLNRomanBaseline = 32
	KBSLNRomanBaselineValue        KBSLNRomanBaseline = 0
)

func (e KBSLNRomanBaseline) String() string {
	switch e {
	case KBSLNHangingBaseline:
		return "KBSLNHangingBaseline"
	case KBSLNIdeographicCenterBaseline:
		return "KBSLNIdeographicCenterBaseline"
	case KBSLNIdeographicHighBaseline:
		return "KBSLNIdeographicHighBaseline"
	case KBSLNIdeographicLowBaseline:
		return "KBSLNIdeographicLowBaseline"
	case KBSLNLastBaseline:
		return "KBSLNLastBaseline"
	case KBSLNMathBaseline:
		return "KBSLNMathBaseline"
	case KBSLNNoBaseline:
		return "KBSLNNoBaseline"
	case KBSLNNumBaselineClasses:
		return "KBSLNNumBaselineClasses"
	case KBSLNRomanBaselineValue:
		return "KBSLNRomanBaselineValue"
	default:
		return fmt.Sprintf("KBSLNRomanBaseline(%d)", e)
	}
}

type KBSLNTag uint32

const (
	KBSLNControlPointFormatNoMap   KBSLNTag = 2
	KBSLNControlPointFormatWithMap KBSLNTag = 3
	KBSLNCurrentVersion            KBSLNTag = 0x10000
	KBSLNDistanceFormatNoMap       KBSLNTag = 0
	KBSLNDistanceFormatWithMap     KBSLNTag = 1
	KBSLNTagValue                  KBSLNTag = 0x62736c6e
)

func (e KBSLNTag) String() string {
	switch e {
	case KBSLNControlPointFormatNoMap:
		return "KBSLNControlPointFormatNoMap"
	case KBSLNControlPointFormatWithMap:
		return "KBSLNControlPointFormatWithMap"
	case KBSLNCurrentVersion:
		return "KBSLNCurrentVersion"
	case KBSLNDistanceFormatNoMap:
		return "KBSLNDistanceFormatNoMap"
	case KBSLNDistanceFormatWithMap:
		return "KBSLNDistanceFormatWithMap"
	case KBSLNTagValue:
		return "KBSLNTagValue"
	default:
		return fmt.Sprintf("KBSLNTag(%d)", e)
	}
}

type KCJKVerticalRoman uint32

const (
	KCJKVerticalRomanCenteredSelector  KCJKVerticalRoman = 0
	KCJKVerticalRomanHBaselineSelector KCJKVerticalRoman = 1
)

func (e KCJKVerticalRoman) String() string {
	switch e {
	case KCJKVerticalRomanCenteredSelector:
		return "KCJKVerticalRomanCenteredSelector"
	case KCJKVerticalRomanHBaselineSelector:
		return "KCJKVerticalRomanHBaselineSelector"
	default:
		return fmt.Sprintf("KCJKVerticalRoman(%d)", e)
	}
}

const KCTFontClassMaskShift uint32 = 28

type KCTFontPriority uint32

const (
	// KCTFontPriorityComputer: Priority of computer local fonts.
	KCTFontPriorityComputer KCTFontPriority = 30000
	// KCTFontPriorityDynamic: Priority of fonts registered dynamically, not located in a standard location.
	KCTFontPriorityDynamic KCTFontPriority = 50000
	// KCTFontPriorityNetwork: Priority of network fonts.
	KCTFontPriorityNetwork KCTFontPriority = 20000
	// KCTFontPriorityProcess: Priority of fonts registered for the process.
	KCTFontPriorityProcess KCTFontPriority = 60000
	// KCTFontPrioritySystem: Priority of system fonts.
	KCTFontPrioritySystem KCTFontPriority = 10000
	// KCTFontPriorityUser: Priority of local fonts.
	KCTFontPriorityUser KCTFontPriority = 40000
)

func (e KCTFontPriority) String() string {
	switch e {
	case KCTFontPriorityComputer:
		return "KCTFontPriorityComputer"
	case KCTFontPriorityDynamic:
		return "KCTFontPriorityDynamic"
	case KCTFontPriorityNetwork:
		return "KCTFontPriorityNetwork"
	case KCTFontPriorityProcess:
		return "KCTFontPriorityProcess"
	case KCTFontPrioritySystem:
		return "KCTFontPrioritySystem"
	case KCTFontPriorityUser:
		return "KCTFontPriorityUser"
	default:
		return fmt.Sprintf("KCTFontPriority(%d)", e)
	}
}

type KCTFontTable uint32

const (
	// KCTFontTableAcnt: # Discussion
	KCTFontTableAcnt KCTFontTable = 'a'<<24 | 'c'<<16 | 'n'<<8 | 't' // 'acnt'
	KCTFontTableAnkr KCTFontTable = 'a'<<24 | 'n'<<16 | 'k'<<8 | 'r' // 'ankr'
	// KCTFontTableAvar: # Discussion
	KCTFontTableAvar KCTFontTable = 'a'<<24 | 'v'<<16 | 'a'<<8 | 'r' // 'avar'
	// KCTFontTableBASE: # Discussion
	KCTFontTableBASE KCTFontTable = 'B'<<24 | 'A'<<16 | 'S'<<8 | 'E' // 'BASE'
	// KCTFontTableBdat: # Discussion
	KCTFontTableBdat KCTFontTable = 'b'<<24 | 'd'<<16 | 'a'<<8 | 't' // 'bdat'
	// KCTFontTableBhed: # Discussion
	KCTFontTableBhed KCTFontTable = 'b'<<24 | 'h'<<16 | 'e'<<8 | 'd' // 'bhed'
	// KCTFontTableBloc: # Discussion
	KCTFontTableBloc KCTFontTable = 'b'<<24 | 'l'<<16 | 'o'<<8 | 'c' // 'bloc'
	// KCTFontTableBsln: # Discussion
	KCTFontTableBsln KCTFontTable = 'b'<<24 | 's'<<16 | 'l'<<8 | 'n' // 'bsln'
	KCTFontTableCBDT KCTFontTable = 'C'<<24 | 'B'<<16 | 'D'<<8 | 'T' // 'CBDT'
	KCTFontTableCBLC KCTFontTable = 'C'<<24 | 'B'<<16 | 'L'<<8 | 'C' // 'CBLC'
	// KCTFontTableCFF: # Discussion
	KCTFontTableCFF  KCTFontTable = 'C'<<24 | 'F'<<16 | 'F'<<8 | ' ' // 'CFF '
	KCTFontTableCFF2 KCTFontTable = 'C'<<24 | 'F'<<16 | 'F'<<8 | '2' // 'CFF2'
	KCTFontTableCOLR KCTFontTable = 'C'<<24 | 'O'<<16 | 'L'<<8 | 'R' // 'COLR'
	KCTFontTableCPAL KCTFontTable = 'C'<<24 | 'P'<<16 | 'A'<<8 | 'L' // 'CPAL'
	KCTFontTableCidg KCTFontTable = 'c'<<24 | 'i'<<16 | 'd'<<8 | 'g' // 'cidg'
	// KCTFontTableCmap: # Discussion
	KCTFontTableCmap KCTFontTable = 'c'<<24 | 'm'<<16 | 'a'<<8 | 'p' // 'cmap'
	// KCTFontTableCvar: # Discussion
	KCTFontTableCvar KCTFontTable = 'c'<<24 | 'v'<<16 | 'a'<<8 | 'r' // 'cvar'
	// KCTFontTableCvt: # Discussion
	KCTFontTableCvt KCTFontTable = 'c'<<24 | 'v'<<16 | 't'<<8 | ' ' // 'cvt '
	// KCTFontTableDSIG: # Discussion
	KCTFontTableDSIG KCTFontTable = 'D'<<24 | 'S'<<16 | 'I'<<8 | 'G' // 'DSIG'
	// KCTFontTableEBDT: # Discussion
	KCTFontTableEBDT KCTFontTable = 'E'<<24 | 'B'<<16 | 'D'<<8 | 'T' // 'EBDT'
	// KCTFontTableEBLC: # Discussion
	KCTFontTableEBLC KCTFontTable = 'E'<<24 | 'B'<<16 | 'L'<<8 | 'C' // 'EBLC'
	// KCTFontTableEBSC: # Discussion
	KCTFontTableEBSC KCTFontTable = 'E'<<24 | 'B'<<16 | 'S'<<8 | 'C' // 'EBSC'
	// KCTFontTableFdsc: # Discussion
	KCTFontTableFdsc KCTFontTable = 'f'<<24 | 'd'<<16 | 's'<<8 | 'c' // 'fdsc'
	// KCTFontTableFeat: # Discussion
	KCTFontTableFeat KCTFontTable = 'f'<<24 | 'e'<<16 | 'a'<<8 | 't' // 'feat'
	// KCTFontTableFmtx: # Discussion
	KCTFontTableFmtx KCTFontTable = 'f'<<24 | 'm'<<16 | 't'<<8 | 'x' // 'fmtx'
	KCTFontTableFond KCTFontTable = 'f'<<24 | 'o'<<16 | 'n'<<8 | 'd' // 'fond'
	// KCTFontTableFpgm: # Discussion
	KCTFontTableFpgm KCTFontTable = 'f'<<24 | 'p'<<16 | 'g'<<8 | 'm' // 'fpgm'
	// KCTFontTableFvar: # Discussion
	KCTFontTableFvar KCTFontTable = 'f'<<24 | 'v'<<16 | 'a'<<8 | 'r' // 'fvar'
	// KCTFontTableGDEF: # Discussion
	KCTFontTableGDEF KCTFontTable = 'G'<<24 | 'D'<<16 | 'E'<<8 | 'F' // 'GDEF'
	// KCTFontTableGPOS: # Discussion
	KCTFontTableGPOS KCTFontTable = 'G'<<24 | 'P'<<16 | 'O'<<8 | 'S' // 'GPOS'
	// KCTFontTableGSUB: # Discussion
	KCTFontTableGSUB KCTFontTable = 'G'<<24 | 'S'<<16 | 'U'<<8 | 'B' // 'GSUB'
	// KCTFontTableGasp: # Discussion
	KCTFontTableGasp KCTFontTable = 'g'<<24 | 'a'<<16 | 's'<<8 | 'p' // 'gasp'
	// KCTFontTableGlyf: # Discussion
	KCTFontTableGlyf KCTFontTable = 'g'<<24 | 'l'<<16 | 'y'<<8 | 'f' // 'glyf'
	// KCTFontTableGvar: # Discussion
	KCTFontTableGvar KCTFontTable = 'g'<<24 | 'v'<<16 | 'a'<<8 | 'r' // 'gvar'
	KCTFontTableHVAR KCTFontTable = 'H'<<24 | 'V'<<16 | 'A'<<8 | 'R' // 'HVAR'
	// KCTFontTableHdmx: # Discussion
	KCTFontTableHdmx KCTFontTable = 'h'<<24 | 'd'<<16 | 'm'<<8 | 'x' // 'hdmx'
	// KCTFontTableHead: # Discussion
	KCTFontTableHead KCTFontTable = 'h'<<24 | 'e'<<16 | 'a'<<8 | 'd' // 'head'
	// KCTFontTableHhea: # Discussion
	KCTFontTableHhea KCTFontTable = 'h'<<24 | 'h'<<16 | 'e'<<8 | 'a' // 'hhea'
	// KCTFontTableHmtx: # Discussion
	KCTFontTableHmtx KCTFontTable = 'h'<<24 | 'm'<<16 | 't'<<8 | 'x' // 'hmtx'
	// KCTFontTableHsty: # Discussion
	KCTFontTableHsty KCTFontTable = 'h'<<24 | 's'<<16 | 't'<<8 | 'y' // 'hsty'
	// KCTFontTableJSTF: # Discussion
	KCTFontTableJSTF KCTFontTable = 'J'<<24 | 'S'<<16 | 'T'<<8 | 'F' // 'JSTF'
	// KCTFontTableJust: # Discussion
	KCTFontTableJust KCTFontTable = 'j'<<24 | 'u'<<16 | 's'<<8 | 't' // 'just'
	// KCTFontTableKern: # Discussion
	KCTFontTableKern KCTFontTable = 'k'<<24 | 'e'<<16 | 'r'<<8 | 'n' // 'kern'
	// KCTFontTableKerx: # Discussion
	KCTFontTableKerx KCTFontTable = 'k'<<24 | 'e'<<16 | 'r'<<8 | 'x' // 'kerx'
	// KCTFontTableLTSH: # Discussion
	KCTFontTableLTSH KCTFontTable = 'L'<<24 | 'T'<<16 | 'S'<<8 | 'H' // 'LTSH'
	// KCTFontTableLcar: # Discussion
	KCTFontTableLcar KCTFontTable = 'l'<<24 | 'c'<<16 | 'a'<<8 | 'r' // 'lcar'
	// KCTFontTableLoca: # Discussion
	KCTFontTableLoca KCTFontTable = 'l'<<24 | 'o'<<16 | 'c'<<8 | 'a' // 'loca'
	KCTFontTableLtag KCTFontTable = 'l'<<24 | 't'<<16 | 'a'<<8 | 'g' // 'ltag'
	KCTFontTableMATH KCTFontTable = 'M'<<24 | 'A'<<16 | 'T'<<8 | 'H' // 'MATH'
	KCTFontTableMERG KCTFontTable = 'M'<<24 | 'E'<<16 | 'R'<<8 | 'G' // 'MERG'
	KCTFontTableMVAR KCTFontTable = 'M'<<24 | 'V'<<16 | 'A'<<8 | 'R' // 'MVAR'
	// KCTFontTableMaxp: # Discussion
	KCTFontTableMaxp KCTFontTable = 'm'<<24 | 'a'<<16 | 'x'<<8 | 'p' // 'maxp'
	KCTFontTableMeta KCTFontTable = 'm'<<24 | 'e'<<16 | 't'<<8 | 'a' // 'meta'
	// KCTFontTableMort: # Discussion
	KCTFontTableMort KCTFontTable = 'm'<<24 | 'o'<<16 | 'r'<<8 | 't' // 'mort'
	// KCTFontTableMorx: # Discussion
	KCTFontTableMorx KCTFontTable = 'm'<<24 | 'o'<<16 | 'r'<<8 | 'x' // 'morx'
	// KCTFontTableName: # Discussion
	KCTFontTableName KCTFontTable = 'n'<<24 | 'a'<<16 | 'm'<<8 | 'e' // 'name'
	// KCTFontTableOS2: # Discussion
	KCTFontTableOS2 KCTFontTable = 'O'<<24 | 'S'<<16 | '/'<<8 | '2' // 'OS/2'
	// KCTFontTableOpbd: # Discussion
	KCTFontTableOpbd KCTFontTable = 'o'<<24 | 'p'<<16 | 'b'<<8 | 'd' // 'opbd'
	// KCTFontTablePCLT: # Discussion
	KCTFontTablePCLT KCTFontTable = 'P'<<24 | 'C'<<16 | 'L'<<8 | 'T' // 'PCLT'
	// KCTFontTablePost: # Discussion
	KCTFontTablePost KCTFontTable = 'p'<<24 | 'o'<<16 | 's'<<8 | 't' // 'post'
	// KCTFontTablePrep: # Discussion
	KCTFontTablePrep KCTFontTable = 'p'<<24 | 'r'<<16 | 'e'<<8 | 'p' // 'prep'
	// KCTFontTableProp: # Discussion
	KCTFontTableProp KCTFontTable = 'p'<<24 | 'r'<<16 | 'o'<<8 | 'p' // 'prop'
	KCTFontTableSTAT KCTFontTable = 'S'<<24 | 'T'<<16 | 'A'<<8 | 'T' // 'STAT'
	KCTFontTableSVG  KCTFontTable = 'S'<<24 | 'V'<<16 | 'G'<<8 | ' ' // 'SVG '
	// KCTFontTableSbit: Font table tag for bitmap data.
	KCTFontTableSbit KCTFontTable = 's'<<24 | 'b'<<16 | 'i'<<8 | 't' // 'sbit'
	// KCTFontTableSbix: Font table tag for extended bitmap data.
	KCTFontTableSbix KCTFontTable = 's'<<24 | 'b'<<16 | 'i'<<8 | 'x' // 'sbix'
	// KCTFontTableTrak: # Discussion
	KCTFontTableTrak KCTFontTable = 't'<<24 | 'r'<<16 | 'a'<<8 | 'k' // 'trak'
	// KCTFontTableVDMX: # Discussion
	KCTFontTableVDMX KCTFontTable = 'V'<<24 | 'D'<<16 | 'M'<<8 | 'X' // 'VDMX'
	// KCTFontTableVORG: # Discussion
	KCTFontTableVORG KCTFontTable = 'V'<<24 | 'O'<<16 | 'R'<<8 | 'G' // 'VORG'
	KCTFontTableVVAR KCTFontTable = 'V'<<24 | 'V'<<16 | 'A'<<8 | 'R' // 'VVAR'
	// KCTFontTableVhea: # Discussion
	KCTFontTableVhea KCTFontTable = 'v'<<24 | 'h'<<16 | 'e'<<8 | 'a' // 'vhea'
	// KCTFontTableVmtx: # Discussion
	KCTFontTableVmtx KCTFontTable = 'v'<<24 | 'm'<<16 | 't'<<8 | 'x' // 'vmtx'
	KCTFontTableXref KCTFontTable = 'x'<<24 | 'r'<<16 | 'e'<<8 | 'f' // 'xref'
	// KCTFontTableZapf: # Discussion
	KCTFontTableZapf KCTFontTable = 'Z'<<24 | 'a'<<16 | 'p'<<8 | 'f' // 'Zapf'
)

func (e KCTFontTable) String() string {
	switch e {
	case KCTFontTableAcnt:
		return "KCTFontTableAcnt"
	case KCTFontTableAnkr:
		return "KCTFontTableAnkr"
	case KCTFontTableAvar:
		return "KCTFontTableAvar"
	case KCTFontTableBASE:
		return "KCTFontTableBASE"
	case KCTFontTableBdat:
		return "KCTFontTableBdat"
	case KCTFontTableBhed:
		return "KCTFontTableBhed"
	case KCTFontTableBloc:
		return "KCTFontTableBloc"
	case KCTFontTableBsln:
		return "KCTFontTableBsln"
	case KCTFontTableCBDT:
		return "KCTFontTableCBDT"
	case KCTFontTableCBLC:
		return "KCTFontTableCBLC"
	case KCTFontTableCFF:
		return "KCTFontTableCFF"
	case KCTFontTableCFF2:
		return "KCTFontTableCFF2"
	case KCTFontTableCOLR:
		return "KCTFontTableCOLR"
	case KCTFontTableCPAL:
		return "KCTFontTableCPAL"
	case KCTFontTableCidg:
		return "KCTFontTableCidg"
	case KCTFontTableCmap:
		return "KCTFontTableCmap"
	case KCTFontTableCvar:
		return "KCTFontTableCvar"
	case KCTFontTableCvt:
		return "KCTFontTableCvt"
	case KCTFontTableDSIG:
		return "KCTFontTableDSIG"
	case KCTFontTableEBDT:
		return "KCTFontTableEBDT"
	case KCTFontTableEBLC:
		return "KCTFontTableEBLC"
	case KCTFontTableEBSC:
		return "KCTFontTableEBSC"
	case KCTFontTableFdsc:
		return "KCTFontTableFdsc"
	case KCTFontTableFeat:
		return "KCTFontTableFeat"
	case KCTFontTableFmtx:
		return "KCTFontTableFmtx"
	case KCTFontTableFond:
		return "KCTFontTableFond"
	case KCTFontTableFpgm:
		return "KCTFontTableFpgm"
	case KCTFontTableFvar:
		return "KCTFontTableFvar"
	case KCTFontTableGDEF:
		return "KCTFontTableGDEF"
	case KCTFontTableGPOS:
		return "KCTFontTableGPOS"
	case KCTFontTableGSUB:
		return "KCTFontTableGSUB"
	case KCTFontTableGasp:
		return "KCTFontTableGasp"
	case KCTFontTableGlyf:
		return "KCTFontTableGlyf"
	case KCTFontTableGvar:
		return "KCTFontTableGvar"
	case KCTFontTableHVAR:
		return "KCTFontTableHVAR"
	case KCTFontTableHdmx:
		return "KCTFontTableHdmx"
	case KCTFontTableHead:
		return "KCTFontTableHead"
	case KCTFontTableHhea:
		return "KCTFontTableHhea"
	case KCTFontTableHmtx:
		return "KCTFontTableHmtx"
	case KCTFontTableHsty:
		return "KCTFontTableHsty"
	case KCTFontTableJSTF:
		return "KCTFontTableJSTF"
	case KCTFontTableJust:
		return "KCTFontTableJust"
	case KCTFontTableKern:
		return "KCTFontTableKern"
	case KCTFontTableKerx:
		return "KCTFontTableKerx"
	case KCTFontTableLTSH:
		return "KCTFontTableLTSH"
	case KCTFontTableLcar:
		return "KCTFontTableLcar"
	case KCTFontTableLoca:
		return "KCTFontTableLoca"
	case KCTFontTableLtag:
		return "KCTFontTableLtag"
	case KCTFontTableMATH:
		return "KCTFontTableMATH"
	case KCTFontTableMERG:
		return "KCTFontTableMERG"
	case KCTFontTableMVAR:
		return "KCTFontTableMVAR"
	case KCTFontTableMaxp:
		return "KCTFontTableMaxp"
	case KCTFontTableMeta:
		return "KCTFontTableMeta"
	case KCTFontTableMort:
		return "KCTFontTableMort"
	case KCTFontTableMorx:
		return "KCTFontTableMorx"
	case KCTFontTableName:
		return "KCTFontTableName"
	case KCTFontTableOS2:
		return "KCTFontTableOS2"
	case KCTFontTableOpbd:
		return "KCTFontTableOpbd"
	case KCTFontTablePCLT:
		return "KCTFontTablePCLT"
	case KCTFontTablePost:
		return "KCTFontTablePost"
	case KCTFontTablePrep:
		return "KCTFontTablePrep"
	case KCTFontTableProp:
		return "KCTFontTableProp"
	case KCTFontTableSTAT:
		return "KCTFontTableSTAT"
	case KCTFontTableSVG:
		return "KCTFontTableSVG"
	case KCTFontTableSbit:
		return "KCTFontTableSbit"
	case KCTFontTableSbix:
		return "KCTFontTableSbix"
	case KCTFontTableTrak:
		return "KCTFontTableTrak"
	case KCTFontTableVDMX:
		return "KCTFontTableVDMX"
	case KCTFontTableVORG:
		return "KCTFontTableVORG"
	case KCTFontTableVVAR:
		return "KCTFontTableVVAR"
	case KCTFontTableVhea:
		return "KCTFontTableVhea"
	case KCTFontTableVmtx:
		return "KCTFontTableVmtx"
	case KCTFontTableXref:
		return "KCTFontTableXref"
	case KCTFontTableZapf:
		return "KCTFontTableZapf"
	default:
		return fmt.Sprintf("KCTFontTable(%d)", e)
	}
}

type KCTRunDelegate uint32

const (
	// KCTRunDelegateCurrentVersion: The current version of the run delegate.
	KCTRunDelegateCurrentVersion KCTRunDelegate = 1
	// KCTRunDelegateVersion1: Version 1 of the run delegate.
	KCTRunDelegateVersion1 KCTRunDelegate = 1
)

func (e KCTRunDelegate) String() string {
	switch e {
	case KCTRunDelegateCurrentVersion:
		return "KCTRunDelegateCurrentVersion"
	default:
		return fmt.Sprintf("KCTRunDelegate(%d)", e)
	}
}

type KCTWritingDirection uint32

const (
	KCTWritingDirectionEmbedding KCTWritingDirection = 0
	KCTWritingDirectionOverride  KCTWritingDirection = 2
)

func (e KCTWritingDirection) String() string {
	switch e {
	case KCTWritingDirectionEmbedding:
		return "KCTWritingDirectionEmbedding"
	case KCTWritingDirectionOverride:
		return "KCTWritingDirectionOverride"
	default:
		return fmt.Sprintf("KCTWritingDirection(%d)", e)
	}
}

type KCanonicalCompositionOnSelector uint32

const (
	KCanonicalCompositionOffSelector     KCanonicalCompositionOnSelector = 1
	KCanonicalCompositionOnSelectorValue KCanonicalCompositionOnSelector = 0
	KCompatibilityCompositionOffSelector KCanonicalCompositionOnSelector = 3
	KCompatibilityCompositionOnSelector  KCanonicalCompositionOnSelector = 2
	KTranscodingCompositionOffSelector   KCanonicalCompositionOnSelector = 5
	KTranscodingCompositionOnSelector    KCanonicalCompositionOnSelector = 4
)

func (e KCanonicalCompositionOnSelector) String() string {
	switch e {
	case KCanonicalCompositionOffSelector:
		return "KCanonicalCompositionOffSelector"
	case KCanonicalCompositionOnSelectorValue:
		return "KCanonicalCompositionOnSelectorValue"
	case KCompatibilityCompositionOffSelector:
		return "KCompatibilityCompositionOffSelector"
	case KCompatibilityCompositionOnSelector:
		return "KCompatibilityCompositionOnSelector"
	case KTranscodingCompositionOffSelector:
		return "KTranscodingCompositionOffSelector"
	case KTranscodingCompositionOnSelector:
		return "KTranscodingCompositionOnSelector"
	default:
		return fmt.Sprintf("KCanonicalCompositionOnSelector(%d)", e)
	}
}

type KCaseSensitive uint32

const (
	KCaseSensitiveLayoutOffSelector  KCaseSensitive = 1
	KCaseSensitiveLayoutOnSelector   KCaseSensitive = 0
	KCaseSensitiveSpacingOffSelector KCaseSensitive = 3
	KCaseSensitiveSpacingOnSelector  KCaseSensitive = 2
)

func (e KCaseSensitive) String() string {
	switch e {
	case KCaseSensitiveLayoutOffSelector:
		return "KCaseSensitiveLayoutOffSelector"
	case KCaseSensitiveLayoutOnSelector:
		return "KCaseSensitiveLayoutOnSelector"
	case KCaseSensitiveSpacingOffSelector:
		return "KCaseSensitiveSpacingOffSelector"
	case KCaseSensitiveSpacingOnSelector:
		return "KCaseSensitiveSpacingOnSelector"
	default:
		return fmt.Sprintf("KCaseSensitive(%d)", e)
	}
}

type KContextualAlternatesOnSelector uint32

const (
	KContextualAlternatesOffSelector      KContextualAlternatesOnSelector = 1
	KContextualAlternatesOnSelectorValue  KContextualAlternatesOnSelector = 0
	KContextualSwashAlternatesOffSelector KContextualAlternatesOnSelector = 5
	KContextualSwashAlternatesOnSelector  KContextualAlternatesOnSelector = 4
	KSwashAlternatesOffSelector           KContextualAlternatesOnSelector = 3
	KSwashAlternatesOnSelector            KContextualAlternatesOnSelector = 2
)

func (e KContextualAlternatesOnSelector) String() string {
	switch e {
	case KContextualAlternatesOffSelector:
		return "KContextualAlternatesOffSelector"
	case KContextualAlternatesOnSelectorValue:
		return "KContextualAlternatesOnSelectorValue"
	case KContextualSwashAlternatesOffSelector:
		return "KContextualSwashAlternatesOffSelector"
	case KContextualSwashAlternatesOnSelector:
		return "KContextualSwashAlternatesOnSelector"
	case KSwashAlternatesOffSelector:
		return "KSwashAlternatesOffSelector"
	case KSwashAlternatesOnSelector:
		return "KSwashAlternatesOnSelector"
	default:
		return fmt.Sprintf("KContextualAlternatesOnSelector(%d)", e)
	}
}

type KDefaultLowerCaseSelector uint32

const (
	KDefaultLowerCaseSelectorValue KDefaultLowerCaseSelector = 0
	KLowerCasePetiteCapsSelector   KDefaultLowerCaseSelector = 2
	KLowerCaseSmallCapsSelector    KDefaultLowerCaseSelector = 1
)

func (e KDefaultLowerCaseSelector) String() string {
	switch e {
	case KDefaultLowerCaseSelectorValue:
		return "KDefaultLowerCaseSelectorValue"
	case KLowerCasePetiteCapsSelector:
		return "KLowerCasePetiteCapsSelector"
	case KLowerCaseSmallCapsSelector:
		return "KLowerCaseSmallCapsSelector"
	default:
		return fmt.Sprintf("KDefaultLowerCaseSelector(%d)", e)
	}
}

type KDefaultUpperCaseSelector uint32

const (
	KDefaultUpperCaseSelectorValue KDefaultUpperCaseSelector = 0
	KUpperCasePetiteCapsSelector   KDefaultUpperCaseSelector = 2
	KUpperCaseSmallCapsSelector    KDefaultUpperCaseSelector = 1
)

func (e KDefaultUpperCaseSelector) String() string {
	switch e {
	case KDefaultUpperCaseSelectorValue:
		return "KDefaultUpperCaseSelectorValue"
	case KUpperCasePetiteCapsSelector:
		return "KUpperCasePetiteCapsSelector"
	case KUpperCaseSmallCapsSelector:
		return "KUpperCaseSmallCapsSelector"
	default:
		return fmt.Sprintf("KDefaultUpperCaseSelector(%d)", e)
	}
}

type KDesign uint32

const (
	KDesignLevel1Selector KDesign = 0
	KDesignLevel2Selector KDesign = 1
	KDesignLevel3Selector KDesign = 2
	KDesignLevel4Selector KDesign = 3
	KDesignLevel5Selector KDesign = 4
)

func (e KDesign) String() string {
	switch e {
	case KDesignLevel1Selector:
		return "KDesignLevel1Selector"
	case KDesignLevel2Selector:
		return "KDesignLevel2Selector"
	case KDesignLevel3Selector:
		return "KDesignLevel3Selector"
	case KDesignLevel4Selector:
		return "KDesignLevel4Selector"
	case KDesignLevel5Selector:
		return "KDesignLevel5Selector"
	default:
		return fmt.Sprintf("KDesign(%d)", e)
	}
}

type KFontCopyrightName uint32

const (
	KFontCopyrightNameValue     KFontCopyrightName = 0
	KFontDescriptionName        KFontCopyrightName = 10
	KFontDesignerName           KFontCopyrightName = 9
	KFontDesignerURLName        KFontCopyrightName = 12
	KFontFamilyName             KFontCopyrightName = 1
	KFontFullName               KFontCopyrightName = 4
	KFontLastReservedName       KFontCopyrightName = 255
	KFontLicenseDescriptionName KFontCopyrightName = 13
	KFontLicenseInfoURLName     KFontCopyrightName = 14
	KFontMacCompatibleFullName  KFontCopyrightName = 18
	KFontManufacturerName       KFontCopyrightName = 8
	KFontPostScriptCIDName      KFontCopyrightName = 20
	KFontPostscriptName         KFontCopyrightName = 6
	KFontPreferredFamilyName    KFontCopyrightName = 16
	KFontPreferredSubfamilyName KFontCopyrightName = 17
	KFontSampleTextName         KFontCopyrightName = 19
	KFontStyleName              KFontCopyrightName = 2
	KFontTrademarkName          KFontCopyrightName = 7
	KFontUniqueName             KFontCopyrightName = 3
	KFontVendorURLName          KFontCopyrightName = 11
	KFontVersionName            KFontCopyrightName = 5
)

func (e KFontCopyrightName) String() string {
	switch e {
	case KFontCopyrightNameValue:
		return "KFontCopyrightNameValue"
	case KFontDescriptionName:
		return "KFontDescriptionName"
	case KFontDesignerName:
		return "KFontDesignerName"
	case KFontDesignerURLName:
		return "KFontDesignerURLName"
	case KFontFamilyName:
		return "KFontFamilyName"
	case KFontFullName:
		return "KFontFullName"
	case KFontLastReservedName:
		return "KFontLastReservedName"
	case KFontLicenseDescriptionName:
		return "KFontLicenseDescriptionName"
	case KFontLicenseInfoURLName:
		return "KFontLicenseInfoURLName"
	case KFontMacCompatibleFullName:
		return "KFontMacCompatibleFullName"
	case KFontManufacturerName:
		return "KFontManufacturerName"
	case KFontPostScriptCIDName:
		return "KFontPostScriptCIDName"
	case KFontPostscriptName:
		return "KFontPostscriptName"
	case KFontPreferredFamilyName:
		return "KFontPreferredFamilyName"
	case KFontPreferredSubfamilyName:
		return "KFontPreferredSubfamilyName"
	case KFontSampleTextName:
		return "KFontSampleTextName"
	case KFontStyleName:
		return "KFontStyleName"
	case KFontTrademarkName:
		return "KFontTrademarkName"
	case KFontUniqueName:
		return "KFontUniqueName"
	case KFontVendorURLName:
		return "KFontVendorURLName"
	case KFontVersionName:
		return "KFontVersionName"
	default:
		return fmt.Sprintf("KFontCopyrightName(%d)", e)
	}
}

type KFontCustom8BitScript uint32

const (
	KFontCustom16BitScript     KFontCustom8BitScript = 2
	KFontCustom816BitScript    KFontCustom8BitScript = 1
	KFontCustom8BitScriptValue KFontCustom8BitScript = 0
)

func (e KFontCustom8BitScript) String() string {
	switch e {
	case KFontCustom16BitScript:
		return "KFontCustom16BitScript"
	case KFontCustom816BitScript:
		return "KFontCustom816BitScript"
	case KFontCustom8BitScriptValue:
		return "KFontCustom8BitScriptValue"
	default:
		return fmt.Sprintf("KFontCustom8BitScript(%d)", e)
	}
}

type KFontEnglishLanguage uint32

const (
	KFontAlbanianLanguage     KFontEnglishLanguage = 36
	KFontAmharicLanguage      KFontEnglishLanguage = 85
	KFontArabicLanguage       KFontEnglishLanguage = 12
	KFontArmenianLanguage     KFontEnglishLanguage = 51
	KFontAssameseLanguage     KFontEnglishLanguage = 68
	KFontAymaraLanguage       KFontEnglishLanguage = 134
	KFontAzerbaijanArLanguage KFontEnglishLanguage = 50
	KFontAzerbaijaniLanguage  KFontEnglishLanguage = 49
	KFontBasqueLanguage       KFontEnglishLanguage = 129
	KFontBengaliLanguage      KFontEnglishLanguage = 67
	KFontBulgarianLanguage    KFontEnglishLanguage = 44
	KFontBurmeseLanguage      KFontEnglishLanguage = 77
	KFontByelorussianLanguage KFontEnglishLanguage = 46
	KFontCatalanLanguage      KFontEnglishLanguage = 130
	KFontChewaLanguage        KFontEnglishLanguage = 92
	KFontCroatianLanguage     KFontEnglishLanguage = 18
	KFontCzechLanguage        KFontEnglishLanguage = 38
	KFontDanishLanguage       KFontEnglishLanguage = 7
	KFontDutchLanguage        KFontEnglishLanguage = 4
	KFontDzongkhaLanguage     KFontEnglishLanguage = 137
	KFontEnglishLanguageValue KFontEnglishLanguage = 0
	KFontEsperantoLanguage    KFontEnglishLanguage = 94
	KFontEstonianLanguage     KFontEnglishLanguage = 27
	KFontFaeroeseLanguage     KFontEnglishLanguage = 30
	KFontFarsiLanguage        KFontEnglishLanguage = 31
	KFontFinnishLanguage      KFontEnglishLanguage = 13
	KFontFlemishLanguage      KFontEnglishLanguage = 34
	KFontFrenchLanguage       KFontEnglishLanguage = 1
	KFontGallaLanguage        KFontEnglishLanguage = 87
	KFontGeorgianLanguage     KFontEnglishLanguage = 52
	KFontGermanLanguage       KFontEnglishLanguage = 2
	KFontGreekLanguage        KFontEnglishLanguage = 14
	KFontGuaraniLanguage      KFontEnglishLanguage = 133
	KFontGujaratiLanguage     KFontEnglishLanguage = 69
	KFontHebrewLanguage       KFontEnglishLanguage = 10
	KFontHindiLanguage        KFontEnglishLanguage = 21
	KFontHungarianLanguage    KFontEnglishLanguage = 26
	KFontIcelandicLanguage    KFontEnglishLanguage = 15
	KFontIndonesianLanguage   KFontEnglishLanguage = 81
	KFontIrishLanguage        KFontEnglishLanguage = 35
	KFontItalianLanguage      KFontEnglishLanguage = 3
	KFontJapaneseLanguage     KFontEnglishLanguage = 11
	KFontJavaneseRomLanguage  KFontEnglishLanguage = 138
	KFontKannadaLanguage      KFontEnglishLanguage = 73
	KFontKashmiriLanguage     KFontEnglishLanguage = 61
	KFontKazakhLanguage       KFontEnglishLanguage = 48
	KFontKhmerLanguage        KFontEnglishLanguage = 78
	KFontKirghizLanguage      KFontEnglishLanguage = 54
	KFontKoreanLanguage       KFontEnglishLanguage = 23
	KFontKurdishLanguage      KFontEnglishLanguage = 60
	KFontLaoLanguage          KFontEnglishLanguage = 79
	KFontLappishLanguage      KFontEnglishLanguage = 29
	KFontLatinLanguage        KFontEnglishLanguage = 131
	KFontLatvianLanguage      KFontEnglishLanguage = 28
	KFontLettishLanguage      KFontEnglishLanguage = 28
	KFontLithuanianLanguage   KFontEnglishLanguage = 24
	KFontMacedonianLanguage   KFontEnglishLanguage = 43
	KFontMalagasyLanguage     KFontEnglishLanguage = 93
	KFontMalayArabicLanguage  KFontEnglishLanguage = 84
	KFontMalayRomanLanguage   KFontEnglishLanguage = 83
	KFontMalayalamLanguage    KFontEnglishLanguage = 72
	KFontMalteseLanguage      KFontEnglishLanguage = 16
	KFontMarathiLanguage      KFontEnglishLanguage = 66
	KFontMoldavianLanguage    KFontEnglishLanguage = 53
	KFontMongolianCyrLanguage KFontEnglishLanguage = 58
	KFontMongolianLanguage    KFontEnglishLanguage = 57
	KFontNepaliLanguage       KFontEnglishLanguage = 64
	KFontNorwegianLanguage    KFontEnglishLanguage = 9
	KFontOriyaLanguage        KFontEnglishLanguage = 71
	KFontOromoLanguage        KFontEnglishLanguage = 87
	KFontPashtoLanguage       KFontEnglishLanguage = 59
	KFontPersianLanguage      KFontEnglishLanguage = 31
	KFontPolishLanguage       KFontEnglishLanguage = 25
	KFontPortugueseLanguage   KFontEnglishLanguage = 8
	KFontPunjabiLanguage      KFontEnglishLanguage = 70
	KFontQuechuaLanguage      KFontEnglishLanguage = 132
	KFontRomanianLanguage     KFontEnglishLanguage = 37
	KFontRuandaLanguage       KFontEnglishLanguage = 90
	KFontRundiLanguage        KFontEnglishLanguage = 91
	KFontRussianLanguage      KFontEnglishLanguage = 32
	KFontSaamiskLanguage      KFontEnglishLanguage = 29
	KFontSanskritLanguage     KFontEnglishLanguage = 65
	KFontSerbianLanguage      KFontEnglishLanguage = 42
	KFontSimpChineseLanguage  KFontEnglishLanguage = 33
	KFontSindhiLanguage       KFontEnglishLanguage = 62
	KFontSinhaleseLanguage    KFontEnglishLanguage = 76
	KFontSlovakLanguage       KFontEnglishLanguage = 39
	KFontSlovenianLanguage    KFontEnglishLanguage = 40
	KFontSomaliLanguage       KFontEnglishLanguage = 88
	KFontSpanishLanguage      KFontEnglishLanguage = 6
	KFontSundaneseRomLanguage KFontEnglishLanguage = 139
	KFontSwahiliLanguage      KFontEnglishLanguage = 89
	KFontSwedishLanguage      KFontEnglishLanguage = 5
	KFontTagalogLanguage      KFontEnglishLanguage = 82
	KFontTajikiLanguage       KFontEnglishLanguage = 55
	KFontTamilLanguage        KFontEnglishLanguage = 74
	KFontTatarLanguage        KFontEnglishLanguage = 135
	KFontTeluguLanguage       KFontEnglishLanguage = 75
	KFontThaiLanguage         KFontEnglishLanguage = 22
	KFontTibetanLanguage      KFontEnglishLanguage = 63
	KFontTigrinyaLanguage     KFontEnglishLanguage = 86
	KFontTradChineseLanguage  KFontEnglishLanguage = 19
	KFontTurkishLanguage      KFontEnglishLanguage = 17
	KFontTurkmenLanguage      KFontEnglishLanguage = 56
	KFontUighurLanguage       KFontEnglishLanguage = 136
	KFontUkrainianLanguage    KFontEnglishLanguage = 45
	KFontUrduLanguage         KFontEnglishLanguage = 20
	KFontUzbekLanguage        KFontEnglishLanguage = 47
	KFontVietnameseLanguage   KFontEnglishLanguage = 80
	KFontWelshLanguage        KFontEnglishLanguage = 128
	KFontYiddishLanguage      KFontEnglishLanguage = 41
)

func (e KFontEnglishLanguage) String() string {
	switch e {
	case KFontAlbanianLanguage:
		return "KFontAlbanianLanguage"
	case KFontAmharicLanguage:
		return "KFontAmharicLanguage"
	case KFontArabicLanguage:
		return "KFontArabicLanguage"
	case KFontArmenianLanguage:
		return "KFontArmenianLanguage"
	case KFontAssameseLanguage:
		return "KFontAssameseLanguage"
	case KFontAymaraLanguage:
		return "KFontAymaraLanguage"
	case KFontAzerbaijanArLanguage:
		return "KFontAzerbaijanArLanguage"
	case KFontAzerbaijaniLanguage:
		return "KFontAzerbaijaniLanguage"
	case KFontBasqueLanguage:
		return "KFontBasqueLanguage"
	case KFontBengaliLanguage:
		return "KFontBengaliLanguage"
	case KFontBulgarianLanguage:
		return "KFontBulgarianLanguage"
	case KFontBurmeseLanguage:
		return "KFontBurmeseLanguage"
	case KFontByelorussianLanguage:
		return "KFontByelorussianLanguage"
	case KFontCatalanLanguage:
		return "KFontCatalanLanguage"
	case KFontChewaLanguage:
		return "KFontChewaLanguage"
	case KFontCroatianLanguage:
		return "KFontCroatianLanguage"
	case KFontCzechLanguage:
		return "KFontCzechLanguage"
	case KFontDanishLanguage:
		return "KFontDanishLanguage"
	case KFontDutchLanguage:
		return "KFontDutchLanguage"
	case KFontDzongkhaLanguage:
		return "KFontDzongkhaLanguage"
	case KFontEnglishLanguageValue:
		return "KFontEnglishLanguageValue"
	case KFontEsperantoLanguage:
		return "KFontEsperantoLanguage"
	case KFontEstonianLanguage:
		return "KFontEstonianLanguage"
	case KFontFaeroeseLanguage:
		return "KFontFaeroeseLanguage"
	case KFontFarsiLanguage:
		return "KFontFarsiLanguage"
	case KFontFinnishLanguage:
		return "KFontFinnishLanguage"
	case KFontFlemishLanguage:
		return "KFontFlemishLanguage"
	case KFontFrenchLanguage:
		return "KFontFrenchLanguage"
	case KFontGallaLanguage:
		return "KFontGallaLanguage"
	case KFontGeorgianLanguage:
		return "KFontGeorgianLanguage"
	case KFontGermanLanguage:
		return "KFontGermanLanguage"
	case KFontGreekLanguage:
		return "KFontGreekLanguage"
	case KFontGuaraniLanguage:
		return "KFontGuaraniLanguage"
	case KFontGujaratiLanguage:
		return "KFontGujaratiLanguage"
	case KFontHebrewLanguage:
		return "KFontHebrewLanguage"
	case KFontHindiLanguage:
		return "KFontHindiLanguage"
	case KFontHungarianLanguage:
		return "KFontHungarianLanguage"
	case KFontIcelandicLanguage:
		return "KFontIcelandicLanguage"
	case KFontIndonesianLanguage:
		return "KFontIndonesianLanguage"
	case KFontIrishLanguage:
		return "KFontIrishLanguage"
	case KFontItalianLanguage:
		return "KFontItalianLanguage"
	case KFontJapaneseLanguage:
		return "KFontJapaneseLanguage"
	case KFontJavaneseRomLanguage:
		return "KFontJavaneseRomLanguage"
	case KFontKannadaLanguage:
		return "KFontKannadaLanguage"
	case KFontKashmiriLanguage:
		return "KFontKashmiriLanguage"
	case KFontKazakhLanguage:
		return "KFontKazakhLanguage"
	case KFontKhmerLanguage:
		return "KFontKhmerLanguage"
	case KFontKirghizLanguage:
		return "KFontKirghizLanguage"
	case KFontKoreanLanguage:
		return "KFontKoreanLanguage"
	case KFontKurdishLanguage:
		return "KFontKurdishLanguage"
	case KFontLaoLanguage:
		return "KFontLaoLanguage"
	case KFontLappishLanguage:
		return "KFontLappishLanguage"
	case KFontLatinLanguage:
		return "KFontLatinLanguage"
	case KFontLatvianLanguage:
		return "KFontLatvianLanguage"
	case KFontLithuanianLanguage:
		return "KFontLithuanianLanguage"
	case KFontMacedonianLanguage:
		return "KFontMacedonianLanguage"
	case KFontMalagasyLanguage:
		return "KFontMalagasyLanguage"
	case KFontMalayArabicLanguage:
		return "KFontMalayArabicLanguage"
	case KFontMalayRomanLanguage:
		return "KFontMalayRomanLanguage"
	case KFontMalayalamLanguage:
		return "KFontMalayalamLanguage"
	case KFontMalteseLanguage:
		return "KFontMalteseLanguage"
	case KFontMarathiLanguage:
		return "KFontMarathiLanguage"
	case KFontMoldavianLanguage:
		return "KFontMoldavianLanguage"
	case KFontMongolianCyrLanguage:
		return "KFontMongolianCyrLanguage"
	case KFontMongolianLanguage:
		return "KFontMongolianLanguage"
	case KFontNepaliLanguage:
		return "KFontNepaliLanguage"
	case KFontNorwegianLanguage:
		return "KFontNorwegianLanguage"
	case KFontOriyaLanguage:
		return "KFontOriyaLanguage"
	case KFontPashtoLanguage:
		return "KFontPashtoLanguage"
	case KFontPolishLanguage:
		return "KFontPolishLanguage"
	case KFontPortugueseLanguage:
		return "KFontPortugueseLanguage"
	case KFontPunjabiLanguage:
		return "KFontPunjabiLanguage"
	case KFontQuechuaLanguage:
		return "KFontQuechuaLanguage"
	case KFontRomanianLanguage:
		return "KFontRomanianLanguage"
	case KFontRuandaLanguage:
		return "KFontRuandaLanguage"
	case KFontRundiLanguage:
		return "KFontRundiLanguage"
	case KFontRussianLanguage:
		return "KFontRussianLanguage"
	case KFontSanskritLanguage:
		return "KFontSanskritLanguage"
	case KFontSerbianLanguage:
		return "KFontSerbianLanguage"
	case KFontSimpChineseLanguage:
		return "KFontSimpChineseLanguage"
	case KFontSindhiLanguage:
		return "KFontSindhiLanguage"
	case KFontSinhaleseLanguage:
		return "KFontSinhaleseLanguage"
	case KFontSlovakLanguage:
		return "KFontSlovakLanguage"
	case KFontSlovenianLanguage:
		return "KFontSlovenianLanguage"
	case KFontSomaliLanguage:
		return "KFontSomaliLanguage"
	case KFontSpanishLanguage:
		return "KFontSpanishLanguage"
	case KFontSundaneseRomLanguage:
		return "KFontSundaneseRomLanguage"
	case KFontSwahiliLanguage:
		return "KFontSwahiliLanguage"
	case KFontSwedishLanguage:
		return "KFontSwedishLanguage"
	case KFontTagalogLanguage:
		return "KFontTagalogLanguage"
	case KFontTajikiLanguage:
		return "KFontTajikiLanguage"
	case KFontTamilLanguage:
		return "KFontTamilLanguage"
	case KFontTatarLanguage:
		return "KFontTatarLanguage"
	case KFontTeluguLanguage:
		return "KFontTeluguLanguage"
	case KFontThaiLanguage:
		return "KFontThaiLanguage"
	case KFontTibetanLanguage:
		return "KFontTibetanLanguage"
	case KFontTigrinyaLanguage:
		return "KFontTigrinyaLanguage"
	case KFontTradChineseLanguage:
		return "KFontTradChineseLanguage"
	case KFontTurkishLanguage:
		return "KFontTurkishLanguage"
	case KFontTurkmenLanguage:
		return "KFontTurkmenLanguage"
	case KFontUighurLanguage:
		return "KFontUighurLanguage"
	case KFontUkrainianLanguage:
		return "KFontUkrainianLanguage"
	case KFontUrduLanguage:
		return "KFontUrduLanguage"
	case KFontUzbekLanguage:
		return "KFontUzbekLanguage"
	case KFontVietnameseLanguage:
		return "KFontVietnameseLanguage"
	case KFontWelshLanguage:
		return "KFontWelshLanguage"
	case KFontYiddishLanguage:
		return "KFontYiddishLanguage"
	default:
		return fmt.Sprintf("KFontEnglishLanguage(%d)", e)
	}
}

type KFontMicrosoft uint32

const (
	KFontMicrosoftStandardScript KFontMicrosoft = 1
	KFontMicrosoftSymbolScript   KFontMicrosoft = 0
	KFontMicrosoftUCS4Script     KFontMicrosoft = 10
)

func (e KFontMicrosoft) String() string {
	switch e {
	case KFontMicrosoftStandardScript:
		return "KFontMicrosoftStandardScript"
	case KFontMicrosoftSymbolScript:
		return "KFontMicrosoftSymbolScript"
	case KFontMicrosoftUCS4Script:
		return "KFontMicrosoftUCS4Script"
	default:
		return fmt.Sprintf("KFontMicrosoft(%d)", e)
	}
}

type KFontNo uint32

const (
	KFontNoLanguageCode KFontNo = 4294967295
	KFontNoPlatformCode KFontNo = 4294967295
	KFontNoScriptCode   KFontNo = 4294967295
)

func (e KFontNo) String() string {
	switch e {
	case KFontNoLanguageCode:
		return "KFontNoLanguageCode"
	default:
		return fmt.Sprintf("KFontNo(%d)", e)
	}
}

const KFontNoNameCode uint32 = 4294967295

type KFontRomanScript uint32

const (
	KFontAmharicScript            KFontRomanScript = 28
	KFontArabicScript             KFontRomanScript = 4
	KFontArmenianScript           KFontRomanScript = 24
	KFontBengaliScript            KFontRomanScript = 13
	KFontBurmeseScript            KFontRomanScript = 19
	KFontChineseScript            KFontRomanScript = 2
	KFontCyrillicScript           KFontRomanScript = 7
	KFontDevanagariScript         KFontRomanScript = 9
	KFontEastEuropeanRomanScript  KFontRomanScript = 29
	KFontEthiopicScript           KFontRomanScript = 28
	KFontExtendedArabicScript     KFontRomanScript = 31
	KFontGeezScript               KFontRomanScript = 28
	KFontGeorgianScript           KFontRomanScript = 23
	KFontGreekScript              KFontRomanScript = 6
	KFontGujaratiScript           KFontRomanScript = 11
	KFontGurmukhiScript           KFontRomanScript = 10
	KFontHebrewScript             KFontRomanScript = 5
	KFontJapaneseScript           KFontRomanScript = 1
	KFontKannadaScript            KFontRomanScript = 16
	KFontKhmerScript              KFontRomanScript = 20
	KFontKoreanScript             KFontRomanScript = 3
	KFontLaotianScript            KFontRomanScript = 22
	KFontMalayalamScript          KFontRomanScript = 17
	KFontMongolianScript          KFontRomanScript = 27
	KFontOriyaScript              KFontRomanScript = 12
	KFontRSymbolScript            KFontRomanScript = 8
	KFontRomanScriptValue         KFontRomanScript = 0
	KFontRussian                  KFontRomanScript = 7
	KFontSimpleChineseScript      KFontRomanScript = 25
	KFontSindhiScript             KFontRomanScript = 31
	KFontSinhaleseScript          KFontRomanScript = 18
	KFontSlavicScript             KFontRomanScript = 29
	KFontTamilScript              KFontRomanScript = 14
	KFontTeluguScript             KFontRomanScript = 15
	KFontThaiScript               KFontRomanScript = 21
	KFontTibetanScript            KFontRomanScript = 26
	KFontTraditionalChineseScript KFontRomanScript = 2
	KFontUninterpretedScript      KFontRomanScript = 32
	KFontVietnameseScript         KFontRomanScript = 30
)

func (e KFontRomanScript) String() string {
	switch e {
	case KFontAmharicScript:
		return "KFontAmharicScript"
	case KFontArabicScript:
		return "KFontArabicScript"
	case KFontArmenianScript:
		return "KFontArmenianScript"
	case KFontBengaliScript:
		return "KFontBengaliScript"
	case KFontBurmeseScript:
		return "KFontBurmeseScript"
	case KFontChineseScript:
		return "KFontChineseScript"
	case KFontCyrillicScript:
		return "KFontCyrillicScript"
	case KFontDevanagariScript:
		return "KFontDevanagariScript"
	case KFontEastEuropeanRomanScript:
		return "KFontEastEuropeanRomanScript"
	case KFontExtendedArabicScript:
		return "KFontExtendedArabicScript"
	case KFontGeorgianScript:
		return "KFontGeorgianScript"
	case KFontGreekScript:
		return "KFontGreekScript"
	case KFontGujaratiScript:
		return "KFontGujaratiScript"
	case KFontGurmukhiScript:
		return "KFontGurmukhiScript"
	case KFontHebrewScript:
		return "KFontHebrewScript"
	case KFontJapaneseScript:
		return "KFontJapaneseScript"
	case KFontKannadaScript:
		return "KFontKannadaScript"
	case KFontKhmerScript:
		return "KFontKhmerScript"
	case KFontKoreanScript:
		return "KFontKoreanScript"
	case KFontLaotianScript:
		return "KFontLaotianScript"
	case KFontMalayalamScript:
		return "KFontMalayalamScript"
	case KFontMongolianScript:
		return "KFontMongolianScript"
	case KFontOriyaScript:
		return "KFontOriyaScript"
	case KFontRSymbolScript:
		return "KFontRSymbolScript"
	case KFontRomanScriptValue:
		return "KFontRomanScriptValue"
	case KFontSimpleChineseScript:
		return "KFontSimpleChineseScript"
	case KFontSinhaleseScript:
		return "KFontSinhaleseScript"
	case KFontTamilScript:
		return "KFontTamilScript"
	case KFontTeluguScript:
		return "KFontTeluguScript"
	case KFontThaiScript:
		return "KFontThaiScript"
	case KFontTibetanScript:
		return "KFontTibetanScript"
	case KFontUninterpretedScript:
		return "KFontUninterpretedScript"
	case KFontVietnameseScript:
		return "KFontVietnameseScript"
	default:
		return fmt.Sprintf("KFontRomanScript(%d)", e)
	}
}

type KFontUnicodeDefaultSemantics uint32

const (
	KFontISO10646_1993Semantics                KFontUnicodeDefaultSemantics = 2
	KFontUnicodeDefaultSemanticsValue          KFontUnicodeDefaultSemantics = 0
	KFontUnicodeV1_1Semantics                  KFontUnicodeDefaultSemantics = 1
	KFontUnicodeV2_0BMPOnlySemantics           KFontUnicodeDefaultSemantics = 3
	KFontUnicodeV2_0FullCoverageSemantics      KFontUnicodeDefaultSemantics = 4
	KFontUnicodeV4_0VariationSequenceSemantics KFontUnicodeDefaultSemantics = 5
	KFontUnicode_FullRepertoire                KFontUnicodeDefaultSemantics = 6
)

func (e KFontUnicodeDefaultSemantics) String() string {
	switch e {
	case KFontISO10646_1993Semantics:
		return "KFontISO10646_1993Semantics"
	case KFontUnicodeDefaultSemanticsValue:
		return "KFontUnicodeDefaultSemanticsValue"
	case KFontUnicodeV1_1Semantics:
		return "KFontUnicodeV1_1Semantics"
	case KFontUnicodeV2_0BMPOnlySemantics:
		return "KFontUnicodeV2_0BMPOnlySemantics"
	case KFontUnicodeV2_0FullCoverageSemantics:
		return "KFontUnicodeV2_0FullCoverageSemantics"
	case KFontUnicodeV4_0VariationSequenceSemantics:
		return "KFontUnicodeV4_0VariationSequenceSemantics"
	case KFontUnicode_FullRepertoire:
		return "KFontUnicode_FullRepertoire"
	default:
		return fmt.Sprintf("KFontUnicodeDefaultSemantics(%d)", e)
	}
}

type KFontUnicodePlatform uint32

const (
	KFontCustomPlatform       KFontUnicodePlatform = 4
	KFontMacintoshPlatform    KFontUnicodePlatform = 1
	KFontMicrosoftPlatform    KFontUnicodePlatform = 3
	KFontReservedPlatform     KFontUnicodePlatform = 2
	KFontUnicodePlatformValue KFontUnicodePlatform = 0
)

func (e KFontUnicodePlatform) String() string {
	switch e {
	case KFontCustomPlatform:
		return "KFontCustomPlatform"
	case KFontMacintoshPlatform:
		return "KFontMacintoshPlatform"
	case KFontMicrosoftPlatform:
		return "KFontMicrosoftPlatform"
	case KFontReservedPlatform:
		return "KFontReservedPlatform"
	case KFontUnicodePlatformValue:
		return "KFontUnicodePlatformValue"
	default:
		return fmt.Sprintf("KFontUnicodePlatform(%d)", e)
	}
}

type KFullWidthIdeographsSelector uint32

const (
	KFullWidthIdeographsSelectorValue KFullWidthIdeographsSelector = 0
	KHalfWidthIdeographsSelector      KFullWidthIdeographsSelector = 2
	KProportionalIdeographsSelector   KFullWidthIdeographsSelector = 1
)

func (e KFullWidthIdeographsSelector) String() string {
	switch e {
	case KFullWidthIdeographsSelectorValue:
		return "KFullWidthIdeographsSelectorValue"
	case KHalfWidthIdeographsSelector:
		return "KHalfWidthIdeographsSelector"
	case KProportionalIdeographsSelector:
		return "KProportionalIdeographsSelector"
	default:
		return fmt.Sprintf("KFullWidthIdeographsSelector(%d)", e)
	}
}

type KFullWidthKanaSelector uint32

const (
	KFullWidthKanaSelectorValue KFullWidthKanaSelector = 0
	KProportionalKanaSelector   KFullWidthKanaSelector = 1
)

func (e KFullWidthKanaSelector) String() string {
	switch e {
	case KFullWidthKanaSelectorValue:
		return "KFullWidthKanaSelectorValue"
	case KProportionalKanaSelector:
		return "KProportionalKanaSelector"
	default:
		return fmt.Sprintf("KFullWidthKanaSelector(%d)", e)
	}
}

type KHalfWidthCJKRomanSelector uint32

const (
	KDefaultCJKRomanSelector        KHalfWidthCJKRomanSelector = 2
	KFullWidthCJKRomanSelector      KHalfWidthCJKRomanSelector = 3
	KHalfWidthCJKRomanSelectorValue KHalfWidthCJKRomanSelector = 0
	KProportionalCJKRomanSelector   KHalfWidthCJKRomanSelector = 1
)

func (e KHalfWidthCJKRomanSelector) String() string {
	switch e {
	case KDefaultCJKRomanSelector:
		return "KDefaultCJKRomanSelector"
	case KFullWidthCJKRomanSelector:
		return "KFullWidthCJKRomanSelector"
	case KHalfWidthCJKRomanSelectorValue:
		return "KHalfWidthCJKRomanSelectorValue"
	case KProportionalCJKRomanSelector:
		return "KProportionalCJKRomanSelector"
	default:
		return fmt.Sprintf("KHalfWidthCJKRomanSelector(%d)", e)
	}
}

type KHyphenToMinusOnSelector uint32

const (
	KAsteriskToMultiplyOffSelector  KHyphenToMinusOnSelector = 3
	KAsteriskToMultiplyOnSelector   KHyphenToMinusOnSelector = 2
	KExponentsOffSelector           KHyphenToMinusOnSelector = 9
	KExponentsOnSelector            KHyphenToMinusOnSelector = 8
	KHyphenToMinusOffSelector       KHyphenToMinusOnSelector = 1
	KHyphenToMinusOnSelectorValue   KHyphenToMinusOnSelector = 0
	KInequalityLigaturesOffSelector KHyphenToMinusOnSelector = 7
	KInequalityLigaturesOnSelector  KHyphenToMinusOnSelector = 6
	KMathematicalGreekOffSelector   KHyphenToMinusOnSelector = 11
	KMathematicalGreekOnSelector    KHyphenToMinusOnSelector = 10
	KSlashToDivideOffSelector       KHyphenToMinusOnSelector = 5
	KSlashToDivideOnSelector        KHyphenToMinusOnSelector = 4
)

func (e KHyphenToMinusOnSelector) String() string {
	switch e {
	case KAsteriskToMultiplyOffSelector:
		return "KAsteriskToMultiplyOffSelector"
	case KAsteriskToMultiplyOnSelector:
		return "KAsteriskToMultiplyOnSelector"
	case KExponentsOffSelector:
		return "KExponentsOffSelector"
	case KExponentsOnSelector:
		return "KExponentsOnSelector"
	case KHyphenToMinusOffSelector:
		return "KHyphenToMinusOffSelector"
	case KHyphenToMinusOnSelectorValue:
		return "KHyphenToMinusOnSelectorValue"
	case KInequalityLigaturesOffSelector:
		return "KInequalityLigaturesOffSelector"
	case KInequalityLigaturesOnSelector:
		return "KInequalityLigaturesOnSelector"
	case KMathematicalGreekOffSelector:
		return "KMathematicalGreekOffSelector"
	case KMathematicalGreekOnSelector:
		return "KMathematicalGreekOnSelector"
	case KSlashToDivideOffSelector:
		return "KSlashToDivideOffSelector"
	case KSlashToDivideOnSelector:
		return "KSlashToDivideOnSelector"
	default:
		return fmt.Sprintf("KHyphenToMinusOnSelector(%d)", e)
	}
}

type KHyphensToEmDashOnSelector uint32

const (
	KFormInterrobangOffSelector     KHyphensToEmDashOnSelector = 7
	KFormInterrobangOnSelector      KHyphensToEmDashOnSelector = 6
	KHyphenToEnDashOffSelector      KHyphensToEmDashOnSelector = 3
	KHyphenToEnDashOnSelector       KHyphensToEmDashOnSelector = 2
	KHyphensToEmDashOffSelector     KHyphensToEmDashOnSelector = 1
	KHyphensToEmDashOnSelectorValue KHyphensToEmDashOnSelector = 0
	KPeriodsToEllipsisOffSelector   KHyphensToEmDashOnSelector = 11
	KPeriodsToEllipsisOnSelector    KHyphensToEmDashOnSelector = 10
	KSlashedZeroOffSelector         KHyphensToEmDashOnSelector = 5
	KSlashedZeroOnSelector          KHyphensToEmDashOnSelector = 4
	KSmartQuotesOffSelector         KHyphensToEmDashOnSelector = 9
	KSmartQuotesOnSelector          KHyphensToEmDashOnSelector = 8
)

func (e KHyphensToEmDashOnSelector) String() string {
	switch e {
	case KFormInterrobangOffSelector:
		return "KFormInterrobangOffSelector"
	case KFormInterrobangOnSelector:
		return "KFormInterrobangOnSelector"
	case KHyphenToEnDashOffSelector:
		return "KHyphenToEnDashOffSelector"
	case KHyphenToEnDashOnSelector:
		return "KHyphenToEnDashOnSelector"
	case KHyphensToEmDashOffSelector:
		return "KHyphensToEmDashOffSelector"
	case KHyphensToEmDashOnSelectorValue:
		return "KHyphensToEmDashOnSelectorValue"
	case KPeriodsToEllipsisOffSelector:
		return "KPeriodsToEllipsisOffSelector"
	case KPeriodsToEllipsisOnSelector:
		return "KPeriodsToEllipsisOnSelector"
	case KSlashedZeroOffSelector:
		return "KSlashedZeroOffSelector"
	case KSlashedZeroOnSelector:
		return "KSlashedZeroOnSelector"
	case KSmartQuotesOffSelector:
		return "KSmartQuotesOffSelector"
	case KSmartQuotesOnSelector:
		return "KSmartQuotesOnSelector"
	default:
		return fmt.Sprintf("KHyphensToEmDashOnSelector(%d)", e)
	}
}

type KJUS uint32

const (
	KJUSTCurrentVersion           KJUS = 0x10000
	KJUSTStandardFormat           KJUS = 0
	KJUSTTag                      KJUS = 0x6a757374
	KJUSTnoGlyphcode              KJUS = 0xffff
	KJUSTpcConditionalAddAction   KJUS = 2
	KJUSTpcDecompositionAction    KJUS = 0
	KJUSTpcDuctilityAction        KJUS = 4
	KJUSTpcGlyphRepeatAddAction   KJUS = 5
	KJUSTpcGlyphStretchAction     KJUS = 3
	KJUSTpcUnconditionalAddAction KJUS = 1
)

func (e KJUS) String() string {
	switch e {
	case KJUSTCurrentVersion:
		return "KJUSTCurrentVersion"
	case KJUSTStandardFormat:
		return "KJUSTStandardFormat"
	case KJUSTTag:
		return "KJUSTTag"
	case KJUSTnoGlyphcode:
		return "KJUSTnoGlyphcode"
	case KJUSTpcConditionalAddAction:
		return "KJUSTpcConditionalAddAction"
	case KJUSTpcDuctilityAction:
		return "KJUSTpcDuctilityAction"
	case KJUSTpcGlyphRepeatAddAction:
		return "KJUSTpcGlyphRepeatAddAction"
	case KJUSTpcGlyphStretchAction:
		return "KJUSTpcGlyphStretchAction"
	case KJUSTpcUnconditionalAddAction:
		return "KJUSTpcUnconditionalAddAction"
	default:
		return fmt.Sprintf("KJUS(%d)", e)
	}
}

type KJUSTKashidaPriority uint32

const (
	KJUSTKashidaPriorityValue KJUSTKashidaPriority = 0
	KJUSTLetterPriority       KJUSTKashidaPriority = 2
	KJUSTNullPriority         KJUSTKashidaPriority = 3
	KJUSTPriorityCount        KJUSTKashidaPriority = 4
	KJUSTSpacePriority        KJUSTKashidaPriority = 1
)

func (e KJUSTKashidaPriority) String() string {
	switch e {
	case KJUSTKashidaPriorityValue:
		return "KJUSTKashidaPriorityValue"
	case KJUSTLetterPriority:
		return "KJUSTLetterPriority"
	case KJUSTNullPriority:
		return "KJUSTNullPriority"
	case KJUSTPriorityCount:
		return "KJUSTPriorityCount"
	case KJUSTSpacePriority:
		return "KJUSTSpacePriority"
	default:
		return fmt.Sprintf("KJUSTKashidaPriority(%d)", e)
	}
}

type KJUSTOverridePriority uint32

const (
	KJUSTOverrideLimits        KJUSTOverridePriority = 0x4000
	KJUSTOverridePriorityValue KJUSTOverridePriority = 0x8000
	KJUSTOverrideUnlimited     KJUSTOverridePriority = 0x2000
	KJUSTPriorityMask          KJUSTOverridePriority = 0x3
	KJUSTUnlimited             KJUSTOverridePriority = 0x1000
)

func (e KJUSTOverridePriority) String() string {
	switch e {
	case KJUSTOverrideLimits:
		return "KJUSTOverrideLimits"
	case KJUSTOverridePriorityValue:
		return "KJUSTOverridePriorityValue"
	case KJUSTOverrideUnlimited:
		return "KJUSTOverrideUnlimited"
	case KJUSTPriorityMask:
		return "KJUSTPriorityMask"
	case KJUSTUnlimited:
		return "KJUSTUnlimited"
	default:
		return fmt.Sprintf("KJUSTOverridePriority(%d)", e)
	}
}

type KKERNLineStart uint32

const (
	KKERNCrossStreamResetNote KKERNLineStart = 2
	KKERNLineEndKerning       KKERNLineStart = 0x2
	KKERNLineStartValue       KKERNLineStart = 0x1
	KKERNNoCrossKerning       KKERNLineStart = 0x4
	KKERNNoStakeNote          KKERNLineStart = 1
	KKERNNotApplied           KKERNLineStart = 0x1
	KKERNNotesRequested       KKERNLineStart = 0x8
)

func (e KKERNLineStart) String() string {
	switch e {
	case KKERNCrossStreamResetNote:
		return "KKERNCrossStreamResetNote"
	case KKERNLineStartValue:
		return "KKERNLineStartValue"
	case KKERNNoCrossKerning:
		return "KKERNNoCrossKerning"
	case KKERNNotesRequested:
		return "KKERNNotesRequested"
	default:
		return fmt.Sprintf("KKERNLineStart(%d)", e)
	}
}

type KKERNOrderedList uint32

const (
	KKERNIndexArray       KKERNOrderedList = 3
	KKERNOrderedListValue KKERNOrderedList = 0
	KKERNSimpleArray      KKERNOrderedList = 2
	KKERNStateTable       KKERNOrderedList = 1
)

func (e KKERNOrderedList) String() string {
	switch e {
	case KKERNIndexArray:
		return "KKERNIndexArray"
	case KKERNOrderedListValue:
		return "KKERNOrderedListValue"
	case KKERNSimpleArray:
		return "KKERNSimpleArray"
	case KKERNStateTable:
		return "KKERNStateTable"
	default:
		return fmt.Sprintf("KKERNOrderedList(%d)", e)
	}
}

type KKERNTag uint32

const (
	KKERNCrossStream      KKERNTag = 0x4000
	KKERNCurrentVersion   KKERNTag = 0x10000
	KKERNFormatMask       KKERNTag = 0xff
	KKERNResetCrossStream KKERNTag = 0x8000
	KKERNTagValue         KKERNTag = 0x6b65726e
	KKERNUnusedBits       KKERNTag = 0x1f00
	KKERNVariation        KKERNTag = 0x2000
	KKERNVertical         KKERNTag = 0x8000
)

func (e KKERNTag) String() string {
	switch e {
	case KKERNCrossStream:
		return "KKERNCrossStream"
	case KKERNCurrentVersion:
		return "KKERNCurrentVersion"
	case KKERNFormatMask:
		return "KKERNFormatMask"
	case KKERNResetCrossStream:
		return "KKERNResetCrossStream"
	case KKERNTagValue:
		return "KKERNTagValue"
	case KKERNUnusedBits:
		return "KKERNUnusedBits"
	case KKERNVariation:
		return "KKERNVariation"
	default:
		return fmt.Sprintf("KKERNTag(%d)", e)
	}
}

type KKERXActionTypeMask uint32

const (
	KKERXActionOffsetMask        KKERXActionTypeMask = 0xffffff
	KKERXActionTypeAnchorPoints  KKERXActionTypeMask = 1073741824
	KKERXActionTypeControlPoints KKERXActionTypeMask = 0
	KKERXActionTypeCoordinates   KKERXActionTypeMask = 2147483648
	KKERXActionTypeMaskValue     KKERXActionTypeMask = 3221225472
	KKERXUnusedFlags             KKERXActionTypeMask = 0x3f000000
)

func (e KKERXActionTypeMask) String() string {
	switch e {
	case KKERXActionOffsetMask:
		return "KKERXActionOffsetMask"
	case KKERXActionTypeAnchorPoints:
		return "KKERXActionTypeAnchorPoints"
	case KKERXActionTypeControlPoints:
		return "KKERXActionTypeControlPoints"
	case KKERXActionTypeCoordinates:
		return "KKERXActionTypeCoordinates"
	case KKERXActionTypeMaskValue:
		return "KKERXActionTypeMaskValue"
	case KKERXUnusedFlags:
		return "KKERXUnusedFlags"
	default:
		return fmt.Sprintf("KKERXActionTypeMask(%d)", e)
	}
}

type KKERXLineStart uint32

const (
	KKERXCrossStreamResetNote KKERXLineStart = 2
	KKERXLineEndKerning       KKERXLineStart = 0x2
	KKERXLineStartValue       KKERXLineStart = 0x1
	KKERXNoCrossKerning       KKERXLineStart = 0x4
	KKERXNoStakeNote          KKERXLineStart = 1
	KKERXNotApplied           KKERXLineStart = 0x1
	KKERXNotesRequested       KKERXLineStart = 0x8
)

func (e KKERXLineStart) String() string {
	switch e {
	case KKERXCrossStreamResetNote:
		return "KKERXCrossStreamResetNote"
	case KKERXLineStartValue:
		return "KKERXLineStartValue"
	case KKERXNoCrossKerning:
		return "KKERXNoCrossKerning"
	case KKERXNotesRequested:
		return "KKERXNotesRequested"
	default:
		return fmt.Sprintf("KKERXLineStart(%d)", e)
	}
}

type KKERXOrderedList uint32

const (
	KKERXControlPoint     KKERXOrderedList = 4
	KKERXIndexArray       KKERXOrderedList = 6
	KKERXOrderedListValue KKERXOrderedList = 0
	KKERXSimpleArray      KKERXOrderedList = 2
	KKERXStateTable       KKERXOrderedList = 1
)

func (e KKERXOrderedList) String() string {
	switch e {
	case KKERXControlPoint:
		return "KKERXControlPoint"
	case KKERXIndexArray:
		return "KKERXIndexArray"
	case KKERXOrderedListValue:
		return "KKERXOrderedListValue"
	case KKERXSimpleArray:
		return "KKERXSimpleArray"
	case KKERXStateTable:
		return "KKERXStateTable"
	default:
		return fmt.Sprintf("KKERXOrderedList(%d)", e)
	}
}

type KKERXTag int32

const (
	KKERXCrossStream      KKERXTag = 0x40000000
	KKERXCurrentVersion   KKERXTag = 0x20000
	KKERXDescending       KKERXTag = 0x10000000
	KKERXFormatMask       KKERXTag = 0xff
	KKERXResetCrossStream KKERXTag = 0x8000
	KKERXTagValue         KKERXTag = 0x6b657278
	KKERXUnusedBits       KKERXTag = 0xfffff00
	KKERXVariation        KKERXTag = 0x20000000
	KKERXVertical         KKERXTag = -2147483648
)

func (e KKERXTag) String() string {
	switch e {
	case KKERXCrossStream:
		return "KKERXCrossStream"
	case KKERXCurrentVersion:
		return "KKERXCurrentVersion"
	case KKERXDescending:
		return "KKERXDescending"
	case KKERXFormatMask:
		return "KKERXFormatMask"
	case KKERXResetCrossStream:
		return "KKERXResetCrossStream"
	case KKERXTagValue:
		return "KKERXTagValue"
	case KKERXUnusedBits:
		return "KKERXUnusedBits"
	case KKERXVariation:
		return "KKERXVariation"
	case KKERXVertical:
		return "KKERXVertical"
	default:
		return fmt.Sprintf("KKERXTag(%d)", e)
	}
}

const KKERXValuesAreLong uint32 = 0x1

type KLCAR uint32

const (
	KLCARCtlPointFormat KLCAR = 1
	KLCARCurrentVersion KLCAR = 0x10000
	KLCARLinearFormat   KLCAR = 0
	KLCARTag            KLCAR = 0x6c636172
)

func (e KLCAR) String() string {
	switch e {
	case KLCARCtlPointFormat:
		return "KLCARCtlPointFormat"
	case KLCARCurrentVersion:
		return "KLCARCurrentVersion"
	case KLCARLinearFormat:
		return "KLCARLinearFormat"
	case KLCARTag:
		return "KLCARTag"
	default:
		return fmt.Sprintf("KLCAR(%d)", e)
	}
}

const KLTAGCurrentVersion uint32 = 1

type KLinguisticRearrangement uint32

const (
	KLinguisticRearrangementOffSelector KLinguisticRearrangement = 1
	KLinguisticRearrangementOnSelector  KLinguisticRearrangement = 0
)

func (e KLinguisticRearrangement) String() string {
	switch e {
	case KLinguisticRearrangementOffSelector:
		return "KLinguisticRearrangementOffSelector"
	case KLinguisticRearrangementOnSelector:
		return "KLinguisticRearrangementOnSelector"
	default:
		return fmt.Sprintf("KLinguisticRearrangement(%d)", e)
	}
}

type KLowerCaseNumbersSelector uint32

const (
	KLowerCaseNumbersSelectorValue KLowerCaseNumbersSelector = 0
	KUpperCaseNumbersSelector      KLowerCaseNumbersSelector = 1
)

func (e KLowerCaseNumbersSelector) String() string {
	switch e {
	case KLowerCaseNumbersSelectorValue:
		return "KLowerCaseNumbersSelectorValue"
	case KUpperCaseNumbersSelector:
		return "KUpperCaseNumbersSelector"
	default:
		return fmt.Sprintf("KLowerCaseNumbersSelector(%d)", e)
	}
}

type KMOR int32

const (
	KMORTContextualType          KMOR = 1
	KMORTCoverDescending         KMOR = 0x4000
	KMORTCoverIgnoreVertical     KMOR = 0x2000
	KMORTCoverTypeMask           KMOR = 0xf
	KMORTCoverVertical           KMOR = 0x8000
	KMORTCurrInsertBefore        KMOR = 0x800
	KMORTCurrInsertCountMask     KMOR = 0x3e0
	KMORTCurrInsertCountShift    KMOR = 5
	KMORTCurrInsertKashidaLike   KMOR = 0x2000
	KMORTCurrJustTableCountMask  KMOR = 0x7f
	KMORTCurrJustTableCountShift KMOR = 0
	KMORTCurrentVersion          KMOR = 0x10000
	KMORTDoInsertionsBefore      KMOR = 0x80
	KMORTInsertionType           KMOR = 5
	KMORTInsertionsCountMask     KMOR = 0x3f
	KMORTIsSplitVowelPiece       KMOR = 0x40
	KMORTLigFormOffsetMask       KMOR = 0x3fffffff
	KMORTLigFormOffsetShift      KMOR = 2
	KMORTLigLastAction           KMOR = -2147483648
	KMORTLigStoreLigature        KMOR = 0x40000000
	KMORTLigatureType            KMOR = 2
	KMORTMarkInsertBefore        KMOR = 0x400
	KMORTMarkInsertCountMask     KMOR = 0x1f
	KMORTMarkInsertCountShift    KMOR = 0
	KMORTMarkInsertKashidaLike   KMOR = 0x1000
	KMORTMarkJustTableCountMask  KMOR = 0x3f80
	KMORTMarkJustTableCountShift KMOR = 7
	KMORTRearrangementType       KMOR = 0
	KMORTSwashType               KMOR = 4
	KMORTTag                     KMOR = 0x6d6f7274
	KMORTraCDx                   KMOR = 6
	KMORTraCDxA                  KMOR = 8
	KMORTraCDxAB                 KMOR = 12
	KMORTraCDxBA                 KMOR = 13
	KMORTraDCx                   KMOR = 7
	KMORTraDCxA                  KMOR = 9
	KMORTraDCxAB                 KMOR = 14
	KMORTraDCxBA                 KMOR = 15
	KMORTraDx                    KMOR = 2
	KMORTraDxA                   KMOR = 3
	KMORTraDxAB                  KMOR = 10
	KMORTraDxBA                  KMOR = 11
	KMORTraNoAction              KMOR = 0
	KMORTraxA                    KMOR = 1
	KMORTraxAB                   KMOR = 4
	KMORTraxBA                   KMOR = 5
)

func (e KMOR) String() string {
	switch e {
	case KMORTContextualType:
		return "KMORTContextualType"
	case KMORTCoverDescending:
		return "KMORTCoverDescending"
	case KMORTCoverIgnoreVertical:
		return "KMORTCoverIgnoreVertical"
	case KMORTCoverTypeMask:
		return "KMORTCoverTypeMask"
	case KMORTCoverVertical:
		return "KMORTCoverVertical"
	case KMORTCurrInsertBefore:
		return "KMORTCurrInsertBefore"
	case KMORTCurrInsertCountMask:
		return "KMORTCurrInsertCountMask"
	case KMORTCurrInsertCountShift:
		return "KMORTCurrInsertCountShift"
	case KMORTCurrJustTableCountMask:
		return "KMORTCurrJustTableCountMask"
	case KMORTCurrJustTableCountShift:
		return "KMORTCurrJustTableCountShift"
	case KMORTCurrentVersion:
		return "KMORTCurrentVersion"
	case KMORTDoInsertionsBefore:
		return "KMORTDoInsertionsBefore"
	case KMORTInsertionsCountMask:
		return "KMORTInsertionsCountMask"
	case KMORTIsSplitVowelPiece:
		return "KMORTIsSplitVowelPiece"
	case KMORTLigFormOffsetMask:
		return "KMORTLigFormOffsetMask"
	case KMORTLigFormOffsetShift:
		return "KMORTLigFormOffsetShift"
	case KMORTLigLastAction:
		return "KMORTLigLastAction"
	case KMORTLigStoreLigature:
		return "KMORTLigStoreLigature"
	case KMORTMarkInsertBefore:
		return "KMORTMarkInsertBefore"
	case KMORTMarkInsertCountMask:
		return "KMORTMarkInsertCountMask"
	case KMORTMarkInsertKashidaLike:
		return "KMORTMarkInsertKashidaLike"
	case KMORTMarkJustTableCountMask:
		return "KMORTMarkJustTableCountMask"
	case KMORTMarkJustTableCountShift:
		return "KMORTMarkJustTableCountShift"
	case KMORTSwashType:
		return "KMORTSwashType"
	case KMORTTag:
		return "KMORTTag"
	case KMORTraCDx:
		return "KMORTraCDx"
	case KMORTraCDxA:
		return "KMORTraCDxA"
	case KMORTraCDxAB:
		return "KMORTraCDxAB"
	case KMORTraCDxBA:
		return "KMORTraCDxBA"
	case KMORTraDCxA:
		return "KMORTraDCxA"
	case KMORTraDCxAB:
		return "KMORTraDCxAB"
	case KMORTraDxA:
		return "KMORTraDxA"
	case KMORTraDxAB:
		return "KMORTraDxAB"
	case KMORTraDxBA:
		return "KMORTraDxBA"
	default:
		return fmt.Sprintf("KMOR(%d)", e)
	}
}

type KMORX int32

const (
	KMORXCoverDescending     KMORX = 0x40000000
	KMORXCoverIgnoreVertical KMORX = 0x20000000
	KMORXCoverLogicalOrder   KMORX = 0x10000000
	KMORXCoverTypeMask       KMORX = 0xff
	KMORXCoverVertical       KMORX = -2147483648
	KMORXCurrentVersion      KMORX = 0x20000
	KMORXTag                 KMORX = 0x6d6f7278
)

func (e KMORX) String() string {
	switch e {
	case KMORXCoverDescending:
		return "KMORXCoverDescending"
	case KMORXCoverIgnoreVertical:
		return "KMORXCoverIgnoreVertical"
	case KMORXCoverLogicalOrder:
		return "KMORXCoverLogicalOrder"
	case KMORXCoverTypeMask:
		return "KMORXCoverTypeMask"
	case KMORXCoverVertical:
		return "KMORXCoverVertical"
	case KMORXCurrentVersion:
		return "KMORXCurrentVersion"
	case KMORXTag:
		return "KMORXTag"
	default:
		return fmt.Sprintf("KMORX(%d)", e)
	}
}

type KMonospacedNumbersSelector uint32

const (
	KMonospacedNumbersSelectorValue KMonospacedNumbersSelector = 0
	KProportionalNumbersSelector    KMonospacedNumbersSelector = 1
	KQuarterWidthNumbersSelector    KMonospacedNumbersSelector = 3
	KThirdWidthNumbersSelector      KMonospacedNumbersSelector = 2
)

func (e KMonospacedNumbersSelector) String() string {
	switch e {
	case KMonospacedNumbersSelectorValue:
		return "KMonospacedNumbersSelectorValue"
	case KProportionalNumbersSelector:
		return "KProportionalNumbersSelector"
	case KQuarterWidthNumbersSelector:
		return "KQuarterWidthNumbersSelector"
	case KThirdWidthNumbersSelector:
		return "KThirdWidthNumbersSelector"
	default:
		return fmt.Sprintf("KMonospacedNumbersSelector(%d)", e)
	}
}

const KNoAlternatesSelector uint32 = 0

type KNoAnnotationSelector uint32

const (
	KBoxAnnotationSelector                KNoAnnotationSelector = 1
	KCircleAnnotationSelector             KNoAnnotationSelector = 3
	KDiamondAnnotationSelector            KNoAnnotationSelector = 8
	KInvertedBoxAnnotationSelector        KNoAnnotationSelector = 9
	KInvertedCircleAnnotationSelector     KNoAnnotationSelector = 4
	KInvertedRoundedBoxAnnotationSelector KNoAnnotationSelector = 10
	KNoAnnotationSelectorValue            KNoAnnotationSelector = 0
	KParenthesisAnnotationSelector        KNoAnnotationSelector = 5
	KPeriodAnnotationSelector             KNoAnnotationSelector = 6
	KRomanNumeralAnnotationSelector       KNoAnnotationSelector = 7
	KRoundedBoxAnnotationSelector         KNoAnnotationSelector = 2
)

func (e KNoAnnotationSelector) String() string {
	switch e {
	case KBoxAnnotationSelector:
		return "KBoxAnnotationSelector"
	case KCircleAnnotationSelector:
		return "KCircleAnnotationSelector"
	case KDiamondAnnotationSelector:
		return "KDiamondAnnotationSelector"
	case KInvertedBoxAnnotationSelector:
		return "KInvertedBoxAnnotationSelector"
	case KInvertedCircleAnnotationSelector:
		return "KInvertedCircleAnnotationSelector"
	case KInvertedRoundedBoxAnnotationSelector:
		return "KInvertedRoundedBoxAnnotationSelector"
	case KNoAnnotationSelectorValue:
		return "KNoAnnotationSelectorValue"
	case KParenthesisAnnotationSelector:
		return "KParenthesisAnnotationSelector"
	case KPeriodAnnotationSelector:
		return "KPeriodAnnotationSelector"
	case KRomanNumeralAnnotationSelector:
		return "KRomanNumeralAnnotationSelector"
	case KRoundedBoxAnnotationSelector:
		return "KRoundedBoxAnnotationSelector"
	default:
		return fmt.Sprintf("KNoAnnotationSelector(%d)", e)
	}
}

type KNoCJKItalicRomanSelector uint32

const (
	KCJKItalicRomanOffSelector     KNoCJKItalicRomanSelector = 3
	KCJKItalicRomanOnSelector      KNoCJKItalicRomanSelector = 2
	KCJKItalicRomanSelector        KNoCJKItalicRomanSelector = 1
	KNoCJKItalicRomanSelectorValue KNoCJKItalicRomanSelector = 0
)

func (e KNoCJKItalicRomanSelector) String() string {
	switch e {
	case KCJKItalicRomanOffSelector:
		return "KCJKItalicRomanOffSelector"
	case KCJKItalicRomanOnSelector:
		return "KCJKItalicRomanOnSelector"
	case KCJKItalicRomanSelector:
		return "KCJKItalicRomanSelector"
	case KNoCJKItalicRomanSelectorValue:
		return "KNoCJKItalicRomanSelectorValue"
	default:
		return fmt.Sprintf("KNoCJKItalicRomanSelector(%d)", e)
	}
}

type KNoCJKSymbolAlternativesSelector uint32

const (
	KCJKSymbolAltFiveSelector             KNoCJKSymbolAlternativesSelector = 5
	KCJKSymbolAltFourSelector             KNoCJKSymbolAlternativesSelector = 4
	KCJKSymbolAltOneSelector              KNoCJKSymbolAlternativesSelector = 1
	KCJKSymbolAltThreeSelector            KNoCJKSymbolAlternativesSelector = 3
	KCJKSymbolAltTwoSelector              KNoCJKSymbolAlternativesSelector = 2
	KNoCJKSymbolAlternativesSelectorValue KNoCJKSymbolAlternativesSelector = 0
)

func (e KNoCJKSymbolAlternativesSelector) String() string {
	switch e {
	case KCJKSymbolAltFiveSelector:
		return "KCJKSymbolAltFiveSelector"
	case KCJKSymbolAltFourSelector:
		return "KCJKSymbolAltFourSelector"
	case KCJKSymbolAltOneSelector:
		return "KCJKSymbolAltOneSelector"
	case KCJKSymbolAltThreeSelector:
		return "KCJKSymbolAltThreeSelector"
	case KCJKSymbolAltTwoSelector:
		return "KCJKSymbolAltTwoSelector"
	case KNoCJKSymbolAlternativesSelectorValue:
		return "KNoCJKSymbolAlternativesSelectorValue"
	default:
		return fmt.Sprintf("KNoCJKSymbolAlternativesSelector(%d)", e)
	}
}

type KNoFractionsSelector uint32

const (
	KDiagonalFractionsSelector KNoFractionsSelector = 2
	KNoFractionsSelectorValue  KNoFractionsSelector = 0
	KVerticalFractionsSelector KNoFractionsSelector = 1
)

func (e KNoFractionsSelector) String() string {
	switch e {
	case KDiagonalFractionsSelector:
		return "KDiagonalFractionsSelector"
	case KNoFractionsSelectorValue:
		return "KNoFractionsSelectorValue"
	case KVerticalFractionsSelector:
		return "KVerticalFractionsSelector"
	default:
		return fmt.Sprintf("KNoFractionsSelector(%d)", e)
	}
}

type KNoIdeographicAlternativesSelector uint32

const (
	KIdeographicAltFiveSelector             KNoIdeographicAlternativesSelector = 5
	KIdeographicAltFourSelector             KNoIdeographicAlternativesSelector = 4
	KIdeographicAltOneSelector              KNoIdeographicAlternativesSelector = 1
	KIdeographicAltThreeSelector            KNoIdeographicAlternativesSelector = 3
	KIdeographicAltTwoSelector              KNoIdeographicAlternativesSelector = 2
	KNoIdeographicAlternativesSelectorValue KNoIdeographicAlternativesSelector = 0
)

func (e KNoIdeographicAlternativesSelector) String() string {
	switch e {
	case KIdeographicAltFiveSelector:
		return "KIdeographicAltFiveSelector"
	case KIdeographicAltFourSelector:
		return "KIdeographicAltFourSelector"
	case KIdeographicAltOneSelector:
		return "KIdeographicAltOneSelector"
	case KIdeographicAltThreeSelector:
		return "KIdeographicAltThreeSelector"
	case KIdeographicAltTwoSelector:
		return "KIdeographicAltTwoSelector"
	case KNoIdeographicAlternativesSelectorValue:
		return "KNoIdeographicAlternativesSelectorValue"
	default:
		return fmt.Sprintf("KNoIdeographicAlternativesSelector(%d)", e)
	}
}

type KNoOrnamentsSelector uint32

const (
	KDecorativeBordersSelector    KNoOrnamentsSelector = 4
	KDingbatsSelector             KNoOrnamentsSelector = 1
	KFleuronsSelector             KNoOrnamentsSelector = 3
	KInternationalSymbolsSelector KNoOrnamentsSelector = 5
	KMathSymbolsSelector          KNoOrnamentsSelector = 6
	KNoOrnamentsSelectorValue     KNoOrnamentsSelector = 0
	KPiCharactersSelector         KNoOrnamentsSelector = 2
)

func (e KNoOrnamentsSelector) String() string {
	switch e {
	case KDecorativeBordersSelector:
		return "KDecorativeBordersSelector"
	case KDingbatsSelector:
		return "KDingbatsSelector"
	case KFleuronsSelector:
		return "KFleuronsSelector"
	case KInternationalSymbolsSelector:
		return "KInternationalSymbolsSelector"
	case KMathSymbolsSelector:
		return "KMathSymbolsSelector"
	case KNoOrnamentsSelectorValue:
		return "KNoOrnamentsSelectorValue"
	case KPiCharactersSelector:
		return "KPiCharactersSelector"
	default:
		return fmt.Sprintf("KNoOrnamentsSelector(%d)", e)
	}
}

type KNoRubyKanaSelector uint32

const (
	KNoRubyKanaSelectorValue KNoRubyKanaSelector = 0
	KRubyKanaOffSelector     KNoRubyKanaSelector = 3
	KRubyKanaOnSelector      KNoRubyKanaSelector = 2
	KRubyKanaSelector        KNoRubyKanaSelector = 1
)

func (e KNoRubyKanaSelector) String() string {
	switch e {
	case KNoRubyKanaSelectorValue:
		return "KNoRubyKanaSelectorValue"
	case KRubyKanaOffSelector:
		return "KRubyKanaOffSelector"
	case KRubyKanaOnSelector:
		return "KRubyKanaOnSelector"
	case KRubyKanaSelector:
		return "KRubyKanaSelector"
	default:
		return fmt.Sprintf("KNoRubyKanaSelector(%d)", e)
	}
}

type KNoStyleOptionsSelector uint32

const (
	KDisplayTextSelector         KNoStyleOptionsSelector = 1
	KEngravedTextSelector        KNoStyleOptionsSelector = 2
	KIlluminatedCapsSelector     KNoStyleOptionsSelector = 3
	KNoStyleOptionsSelectorValue KNoStyleOptionsSelector = 0
	KTallCapsSelector            KNoStyleOptionsSelector = 5
	KTitlingCapsSelector         KNoStyleOptionsSelector = 4
)

func (e KNoStyleOptionsSelector) String() string {
	switch e {
	case KDisplayTextSelector:
		return "KDisplayTextSelector"
	case KEngravedTextSelector:
		return "KEngravedTextSelector"
	case KIlluminatedCapsSelector:
		return "KIlluminatedCapsSelector"
	case KNoStyleOptionsSelectorValue:
		return "KNoStyleOptionsSelectorValue"
	case KTallCapsSelector:
		return "KTallCapsSelector"
	case KTitlingCapsSelector:
		return "KTitlingCapsSelector"
	default:
		return fmt.Sprintf("KNoStyleOptionsSelector(%d)", e)
	}
}

type KNoStylisticAlternatesSelector uint32

const (
	KNoStylisticAlternatesSelectorValue KNoStylisticAlternatesSelector = 0
	KStylisticAltEightOffSelector       KNoStylisticAlternatesSelector = 17
	KStylisticAltEightOnSelector        KNoStylisticAlternatesSelector = 16
	KStylisticAltEighteenOffSelector    KNoStylisticAlternatesSelector = 37
	KStylisticAltEighteenOnSelector     KNoStylisticAlternatesSelector = 36
	KStylisticAltElevenOffSelector      KNoStylisticAlternatesSelector = 23
	KStylisticAltElevenOnSelector       KNoStylisticAlternatesSelector = 22
	KStylisticAltFifteenOffSelector     KNoStylisticAlternatesSelector = 31
	KStylisticAltFifteenOnSelector      KNoStylisticAlternatesSelector = 30
	KStylisticAltFiveOffSelector        KNoStylisticAlternatesSelector = 11
	KStylisticAltFiveOnSelector         KNoStylisticAlternatesSelector = 10
	KStylisticAltFourOffSelector        KNoStylisticAlternatesSelector = 9
	KStylisticAltFourOnSelector         KNoStylisticAlternatesSelector = 8
	KStylisticAltFourteenOffSelector    KNoStylisticAlternatesSelector = 29
	KStylisticAltFourteenOnSelector     KNoStylisticAlternatesSelector = 28
	KStylisticAltNineOffSelector        KNoStylisticAlternatesSelector = 19
	KStylisticAltNineOnSelector         KNoStylisticAlternatesSelector = 18
	KStylisticAltNineteenOffSelector    KNoStylisticAlternatesSelector = 39
	KStylisticAltNineteenOnSelector     KNoStylisticAlternatesSelector = 38
	KStylisticAltOneOffSelector         KNoStylisticAlternatesSelector = 3
	KStylisticAltOneOnSelector          KNoStylisticAlternatesSelector = 2
	KStylisticAltSevenOffSelector       KNoStylisticAlternatesSelector = 15
	KStylisticAltSevenOnSelector        KNoStylisticAlternatesSelector = 14
	KStylisticAltSeventeenOffSelector   KNoStylisticAlternatesSelector = 35
	KStylisticAltSeventeenOnSelector    KNoStylisticAlternatesSelector = 34
	KStylisticAltSixOffSelector         KNoStylisticAlternatesSelector = 13
	KStylisticAltSixOnSelector          KNoStylisticAlternatesSelector = 12
	KStylisticAltSixteenOffSelector     KNoStylisticAlternatesSelector = 33
	KStylisticAltSixteenOnSelector      KNoStylisticAlternatesSelector = 32
	KStylisticAltTenOffSelector         KNoStylisticAlternatesSelector = 21
	KStylisticAltTenOnSelector          KNoStylisticAlternatesSelector = 20
	KStylisticAltThirteenOffSelector    KNoStylisticAlternatesSelector = 27
	KStylisticAltThirteenOnSelector     KNoStylisticAlternatesSelector = 26
	KStylisticAltThreeOffSelector       KNoStylisticAlternatesSelector = 7
	KStylisticAltThreeOnSelector        KNoStylisticAlternatesSelector = 6
	KStylisticAltTwelveOffSelector      KNoStylisticAlternatesSelector = 25
	KStylisticAltTwelveOnSelector       KNoStylisticAlternatesSelector = 24
	KStylisticAltTwentyOffSelector      KNoStylisticAlternatesSelector = 41
	KStylisticAltTwentyOnSelector       KNoStylisticAlternatesSelector = 40
	KStylisticAltTwoOffSelector         KNoStylisticAlternatesSelector = 5
	KStylisticAltTwoOnSelector          KNoStylisticAlternatesSelector = 4
)

func (e KNoStylisticAlternatesSelector) String() string {
	switch e {
	case KNoStylisticAlternatesSelectorValue:
		return "KNoStylisticAlternatesSelectorValue"
	case KStylisticAltEightOffSelector:
		return "KStylisticAltEightOffSelector"
	case KStylisticAltEightOnSelector:
		return "KStylisticAltEightOnSelector"
	case KStylisticAltEighteenOffSelector:
		return "KStylisticAltEighteenOffSelector"
	case KStylisticAltEighteenOnSelector:
		return "KStylisticAltEighteenOnSelector"
	case KStylisticAltElevenOffSelector:
		return "KStylisticAltElevenOffSelector"
	case KStylisticAltElevenOnSelector:
		return "KStylisticAltElevenOnSelector"
	case KStylisticAltFifteenOffSelector:
		return "KStylisticAltFifteenOffSelector"
	case KStylisticAltFifteenOnSelector:
		return "KStylisticAltFifteenOnSelector"
	case KStylisticAltFiveOffSelector:
		return "KStylisticAltFiveOffSelector"
	case KStylisticAltFiveOnSelector:
		return "KStylisticAltFiveOnSelector"
	case KStylisticAltFourOffSelector:
		return "KStylisticAltFourOffSelector"
	case KStylisticAltFourOnSelector:
		return "KStylisticAltFourOnSelector"
	case KStylisticAltFourteenOffSelector:
		return "KStylisticAltFourteenOffSelector"
	case KStylisticAltFourteenOnSelector:
		return "KStylisticAltFourteenOnSelector"
	case KStylisticAltNineOffSelector:
		return "KStylisticAltNineOffSelector"
	case KStylisticAltNineOnSelector:
		return "KStylisticAltNineOnSelector"
	case KStylisticAltNineteenOffSelector:
		return "KStylisticAltNineteenOffSelector"
	case KStylisticAltNineteenOnSelector:
		return "KStylisticAltNineteenOnSelector"
	case KStylisticAltOneOffSelector:
		return "KStylisticAltOneOffSelector"
	case KStylisticAltOneOnSelector:
		return "KStylisticAltOneOnSelector"
	case KStylisticAltSevenOffSelector:
		return "KStylisticAltSevenOffSelector"
	case KStylisticAltSevenOnSelector:
		return "KStylisticAltSevenOnSelector"
	case KStylisticAltSeventeenOffSelector:
		return "KStylisticAltSeventeenOffSelector"
	case KStylisticAltSeventeenOnSelector:
		return "KStylisticAltSeventeenOnSelector"
	case KStylisticAltSixOffSelector:
		return "KStylisticAltSixOffSelector"
	case KStylisticAltSixOnSelector:
		return "KStylisticAltSixOnSelector"
	case KStylisticAltSixteenOffSelector:
		return "KStylisticAltSixteenOffSelector"
	case KStylisticAltSixteenOnSelector:
		return "KStylisticAltSixteenOnSelector"
	case KStylisticAltTenOffSelector:
		return "KStylisticAltTenOffSelector"
	case KStylisticAltTenOnSelector:
		return "KStylisticAltTenOnSelector"
	case KStylisticAltThirteenOffSelector:
		return "KStylisticAltThirteenOffSelector"
	case KStylisticAltThirteenOnSelector:
		return "KStylisticAltThirteenOnSelector"
	case KStylisticAltThreeOffSelector:
		return "KStylisticAltThreeOffSelector"
	case KStylisticAltThreeOnSelector:
		return "KStylisticAltThreeOnSelector"
	case KStylisticAltTwelveOffSelector:
		return "KStylisticAltTwelveOffSelector"
	case KStylisticAltTwelveOnSelector:
		return "KStylisticAltTwelveOnSelector"
	case KStylisticAltTwentyOffSelector:
		return "KStylisticAltTwentyOffSelector"
	case KStylisticAltTwentyOnSelector:
		return "KStylisticAltTwentyOnSelector"
	case KStylisticAltTwoOffSelector:
		return "KStylisticAltTwoOffSelector"
	case KStylisticAltTwoOnSelector:
		return "KStylisticAltTwoOnSelector"
	default:
		return fmt.Sprintf("KNoStylisticAlternatesSelector(%d)", e)
	}
}

type KNoTransliterationSelector uint32

const (
	KHanjaToHangulAltOneSelector    KNoTransliterationSelector = 7
	KHanjaToHangulAltThreeSelector  KNoTransliterationSelector = 9
	KHanjaToHangulAltTwoSelector    KNoTransliterationSelector = 8
	KHanjaToHangulSelector          KNoTransliterationSelector = 1
	KHiraganaToKatakanaSelector     KNoTransliterationSelector = 2
	KKanaToRomanizationSelector     KNoTransliterationSelector = 4
	KKatakanaToHiraganaSelector     KNoTransliterationSelector = 3
	KNoTransliterationSelectorValue KNoTransliterationSelector = 0
	KRomanizationToHiraganaSelector KNoTransliterationSelector = 5
	KRomanizationToKatakanaSelector KNoTransliterationSelector = 6
)

func (e KNoTransliterationSelector) String() string {
	switch e {
	case KHanjaToHangulAltOneSelector:
		return "KHanjaToHangulAltOneSelector"
	case KHanjaToHangulAltThreeSelector:
		return "KHanjaToHangulAltThreeSelector"
	case KHanjaToHangulAltTwoSelector:
		return "KHanjaToHangulAltTwoSelector"
	case KHanjaToHangulSelector:
		return "KHanjaToHangulSelector"
	case KHiraganaToKatakanaSelector:
		return "KHiraganaToKatakanaSelector"
	case KKanaToRomanizationSelector:
		return "KKanaToRomanizationSelector"
	case KKatakanaToHiraganaSelector:
		return "KKatakanaToHiraganaSelector"
	case KNoTransliterationSelectorValue:
		return "KNoTransliterationSelectorValue"
	case KRomanizationToHiraganaSelector:
		return "KRomanizationToHiraganaSelector"
	case KRomanizationToKatakanaSelector:
		return "KRomanizationToKatakanaSelector"
	default:
		return fmt.Sprintf("KNoTransliterationSelector(%d)", e)
	}
}

type KNormalPositionSelector uint32

const (
	KInferiorsSelector           KNormalPositionSelector = 2
	KNormalPositionSelectorValue KNormalPositionSelector = 0
	KOrdinalsSelector            KNormalPositionSelector = 3
	KScientificInferiorsSelector KNormalPositionSelector = 4
	KSuperiorsSelector           KNormalPositionSelector = 1
)

func (e KNormalPositionSelector) String() string {
	switch e {
	case KInferiorsSelector:
		return "KInferiorsSelector"
	case KNormalPositionSelectorValue:
		return "KNormalPositionSelectorValue"
	case KOrdinalsSelector:
		return "KOrdinalsSelector"
	case KScientificInferiorsSelector:
		return "KScientificInferiorsSelector"
	case KSuperiorsSelector:
		return "KSuperiorsSelector"
	default:
		return fmt.Sprintf("KNormalPositionSelector(%d)", e)
	}
}

type KOPBD uint32

const (
	KOPBDControlPointFormat KOPBD = 1
	KOPBDCurrentVersion     KOPBD = 0x10000
	KOPBDDistanceFormat     KOPBD = 0
	KOPBDTag                KOPBD = 0x6f706264
)

func (e KOPBD) String() string {
	switch e {
	case KOPBDControlPointFormat:
		return "KOPBDControlPointFormat"
	case KOPBDCurrentVersion:
		return "KOPBDCurrentVersion"
	case KOPBDDistanceFormat:
		return "KOPBDDistanceFormat"
	case KOPBDTag:
		return "KOPBDTag"
	default:
		return fmt.Sprintf("KOPBD(%d)", e)
	}
}

type KPROPLDirectionClass uint32

const (
	KPROPALDirectionClass     KPROPLDirectionClass = 2
	KPROPANDirectionClass     KPROPLDirectionClass = 6
	KPROPBNDirectionClass     KPROPLDirectionClass = 19
	KPROPCSDirectionClass     KPROPLDirectionClass = 7
	KPROPENDirectionClass     KPROPLDirectionClass = 3
	KPROPESDirectionClass     KPROPLDirectionClass = 4
	KPROPETDirectionClass     KPROPLDirectionClass = 5
	KPROPLDirectionClassValue KPROPLDirectionClass = 0
	KPROPLREDirectionClass    KPROPLDirectionClass = 13
	KPROPLRODirectionClass    KPROPLDirectionClass = 14
	KPROPNSMDirectionClass    KPROPLDirectionClass = 18
	KPROPNumDirectionClasses  KPROPLDirectionClass = 20
	KPROPONDirectionClass     KPROPLDirectionClass = 11
	KPROPPDFDirectionClass    KPROPLDirectionClass = 17
	KPROPPSDirectionClass     KPROPLDirectionClass = 8
	KPROPRDirectionClass      KPROPLDirectionClass = 1
	KPROPRLEDirectionClass    KPROPLDirectionClass = 15
	KPROPRLODirectionClass    KPROPLDirectionClass = 16
	KPROPSDirectionClass      KPROPLDirectionClass = 9
	KPROPSENDirectionClass    KPROPLDirectionClass = 12
	KPROPWSDirectionClass     KPROPLDirectionClass = 10
)

func (e KPROPLDirectionClass) String() string {
	switch e {
	case KPROPALDirectionClass:
		return "KPROPALDirectionClass"
	case KPROPANDirectionClass:
		return "KPROPANDirectionClass"
	case KPROPBNDirectionClass:
		return "KPROPBNDirectionClass"
	case KPROPCSDirectionClass:
		return "KPROPCSDirectionClass"
	case KPROPENDirectionClass:
		return "KPROPENDirectionClass"
	case KPROPESDirectionClass:
		return "KPROPESDirectionClass"
	case KPROPETDirectionClass:
		return "KPROPETDirectionClass"
	case KPROPLDirectionClassValue:
		return "KPROPLDirectionClassValue"
	case KPROPLREDirectionClass:
		return "KPROPLREDirectionClass"
	case KPROPLRODirectionClass:
		return "KPROPLRODirectionClass"
	case KPROPNSMDirectionClass:
		return "KPROPNSMDirectionClass"
	case KPROPNumDirectionClasses:
		return "KPROPNumDirectionClasses"
	case KPROPONDirectionClass:
		return "KPROPONDirectionClass"
	case KPROPPDFDirectionClass:
		return "KPROPPDFDirectionClass"
	case KPROPPSDirectionClass:
		return "KPROPPSDirectionClass"
	case KPROPRDirectionClass:
		return "KPROPRDirectionClass"
	case KPROPRLEDirectionClass:
		return "KPROPRLEDirectionClass"
	case KPROPRLODirectionClass:
		return "KPROPRLODirectionClass"
	case KPROPSDirectionClass:
		return "KPROPSDirectionClass"
	case KPROPSENDirectionClass:
		return "KPROPSENDirectionClass"
	case KPROPWSDirectionClass:
		return "KPROPWSDirectionClass"
	default:
		return fmt.Sprintf("KPROPLDirectionClass(%d)", e)
	}
}

type KPROPTag uint32

const (
	KPROPCanHangLTMask    KPROPTag = 0x4000
	KPROPCanHangRBMask    KPROPTag = 0x2000
	KPROPCurrentVersion   KPROPTag = 0x30000
	KPROPDirectionMask    KPROPTag = 0x1f
	KPROPIsFloaterMask    KPROPTag = 0x8000
	KPROPPairOffsetMask   KPROPTag = 0xf00
	KPROPPairOffsetShift  KPROPTag = 8
	KPROPPairOffsetSign   KPROPTag = 7
	KPROPRightConnectMask KPROPTag = 0x80
	KPROPTagValue         KPROPTag = 0x70726f70
	KPROPUseRLPairMask    KPROPTag = 0x1000
	KPROPZeroReserved     KPROPTag = 0x60
)

func (e KPROPTag) String() string {
	switch e {
	case KPROPCanHangLTMask:
		return "KPROPCanHangLTMask"
	case KPROPCanHangRBMask:
		return "KPROPCanHangRBMask"
	case KPROPCurrentVersion:
		return "KPROPCurrentVersion"
	case KPROPDirectionMask:
		return "KPROPDirectionMask"
	case KPROPIsFloaterMask:
		return "KPROPIsFloaterMask"
	case KPROPPairOffsetMask:
		return "KPROPPairOffsetMask"
	case KPROPPairOffsetShift:
		return "KPROPPairOffsetShift"
	case KPROPPairOffsetSign:
		return "KPROPPairOffsetSign"
	case KPROPRightConnectMask:
		return "KPROPRightConnectMask"
	case KPROPTagValue:
		return "KPROPTagValue"
	case KPROPUseRLPairMask:
		return "KPROPUseRLPairMask"
	case KPROPZeroReserved:
		return "KPROPZeroReserved"
	default:
		return fmt.Sprintf("KPROPTag(%d)", e)
	}
}

type KPreventOverlap uint32

const (
	KPreventOverlapOffSelector KPreventOverlap = 1
	KPreventOverlapOnSelector  KPreventOverlap = 0
)

func (e KPreventOverlap) String() string {
	switch e {
	case KPreventOverlapOffSelector:
		return "KPreventOverlapOffSelector"
	case KPreventOverlapOnSelector:
		return "KPreventOverlapOnSelector"
	default:
		return fmt.Sprintf("KPreventOverlap(%d)", e)
	}
}

type KProportionalTextSelector uint32

const (
	KAltHalfWidthTextSelector      KProportionalTextSelector = 6
	KAltProportionalTextSelector   KProportionalTextSelector = 5
	KHalfWidthTextSelector         KProportionalTextSelector = 2
	KMonospacedTextSelector        KProportionalTextSelector = 1
	KProportionalTextSelectorValue KProportionalTextSelector = 0
	KQuarterWidthTextSelector      KProportionalTextSelector = 4
	KThirdWidthTextSelector        KProportionalTextSelector = 3
)

func (e KProportionalTextSelector) String() string {
	switch e {
	case KAltHalfWidthTextSelector:
		return "KAltHalfWidthTextSelector"
	case KAltProportionalTextSelector:
		return "KAltProportionalTextSelector"
	case KHalfWidthTextSelector:
		return "KHalfWidthTextSelector"
	case KMonospacedTextSelector:
		return "KMonospacedTextSelector"
	case KProportionalTextSelectorValue:
		return "KProportionalTextSelectorValue"
	case KQuarterWidthTextSelector:
		return "KQuarterWidthTextSelector"
	case KThirdWidthTextSelector:
		return "KThirdWidthTextSelector"
	default:
		return fmt.Sprintf("KProportionalTextSelector(%d)", e)
	}
}

type KRequiredLigaturesOnSelector uint32

const (
	KAbbrevSquaredLigaturesOffSelector KRequiredLigaturesOnSelector = 15
	KAbbrevSquaredLigaturesOnSelector  KRequiredLigaturesOnSelector = 14
	KCommonLigaturesOffSelector        KRequiredLigaturesOnSelector = 3
	KCommonLigaturesOnSelector         KRequiredLigaturesOnSelector = 2
	KContextualLigaturesOffSelector    KRequiredLigaturesOnSelector = 19
	KContextualLigaturesOnSelector     KRequiredLigaturesOnSelector = 18
	KDiphthongLigaturesOffSelector     KRequiredLigaturesOnSelector = 11
	KDiphthongLigaturesOnSelector      KRequiredLigaturesOnSelector = 10
	KHistoricalLigaturesOffSelector    KRequiredLigaturesOnSelector = 21
	KHistoricalLigaturesOnSelector     KRequiredLigaturesOnSelector = 20
	KLogosOffSelector                  KRequiredLigaturesOnSelector = 7
	KLogosOnSelector                   KRequiredLigaturesOnSelector = 6
	KRareLigaturesOffSelector          KRequiredLigaturesOnSelector = 5
	KRareLigaturesOnSelector           KRequiredLigaturesOnSelector = 4
	KRebusPicturesOffSelector          KRequiredLigaturesOnSelector = 9
	KRebusPicturesOnSelector           KRequiredLigaturesOnSelector = 8
	KRequiredLigaturesOffSelector      KRequiredLigaturesOnSelector = 1
	KRequiredLigaturesOnSelectorValue  KRequiredLigaturesOnSelector = 0
	KSquaredLigaturesOffSelector       KRequiredLigaturesOnSelector = 13
	KSquaredLigaturesOnSelector        KRequiredLigaturesOnSelector = 12
	KSymbolLigaturesOffSelector        KRequiredLigaturesOnSelector = 17
	KSymbolLigaturesOnSelector         KRequiredLigaturesOnSelector = 16
)

func (e KRequiredLigaturesOnSelector) String() string {
	switch e {
	case KAbbrevSquaredLigaturesOffSelector:
		return "KAbbrevSquaredLigaturesOffSelector"
	case KAbbrevSquaredLigaturesOnSelector:
		return "KAbbrevSquaredLigaturesOnSelector"
	case KCommonLigaturesOffSelector:
		return "KCommonLigaturesOffSelector"
	case KCommonLigaturesOnSelector:
		return "KCommonLigaturesOnSelector"
	case KContextualLigaturesOffSelector:
		return "KContextualLigaturesOffSelector"
	case KContextualLigaturesOnSelector:
		return "KContextualLigaturesOnSelector"
	case KDiphthongLigaturesOffSelector:
		return "KDiphthongLigaturesOffSelector"
	case KDiphthongLigaturesOnSelector:
		return "KDiphthongLigaturesOnSelector"
	case KHistoricalLigaturesOffSelector:
		return "KHistoricalLigaturesOffSelector"
	case KHistoricalLigaturesOnSelector:
		return "KHistoricalLigaturesOnSelector"
	case KLogosOffSelector:
		return "KLogosOffSelector"
	case KLogosOnSelector:
		return "KLogosOnSelector"
	case KRareLigaturesOffSelector:
		return "KRareLigaturesOffSelector"
	case KRareLigaturesOnSelector:
		return "KRareLigaturesOnSelector"
	case KRebusPicturesOffSelector:
		return "KRebusPicturesOffSelector"
	case KRebusPicturesOnSelector:
		return "KRebusPicturesOnSelector"
	case KRequiredLigaturesOffSelector:
		return "KRequiredLigaturesOffSelector"
	case KRequiredLigaturesOnSelectorValue:
		return "KRequiredLigaturesOnSelectorValue"
	case KSquaredLigaturesOffSelector:
		return "KSquaredLigaturesOffSelector"
	case KSquaredLigaturesOnSelector:
		return "KSquaredLigaturesOnSelector"
	case KSymbolLigaturesOffSelector:
		return "KSymbolLigaturesOffSelector"
	case KSymbolLigaturesOnSelector:
		return "KSymbolLigaturesOnSelector"
	default:
		return fmt.Sprintf("KRequiredLigaturesOnSelector(%d)", e)
	}
}

type KSFNTLookup uint32

const (
	KSFNTLookupSegmentArray  KSFNTLookup = 4
	KSFNTLookupSegmentSingle KSFNTLookup = 2
	KSFNTLookupSimpleArray   KSFNTLookup = 0
	KSFNTLookupSingleTable   KSFNTLookup = 6
	KSFNTLookupTrimmedArray  KSFNTLookup = 8
	KSFNTLookupVector        KSFNTLookup = 10
)

func (e KSFNTLookup) String() string {
	switch e {
	case KSFNTLookupSegmentArray:
		return "KSFNTLookupSegmentArray"
	case KSFNTLookupSegmentSingle:
		return "KSFNTLookupSegmentSingle"
	case KSFNTLookupSimpleArray:
		return "KSFNTLookupSimpleArray"
	case KSFNTLookupSingleTable:
		return "KSFNTLookupSingleTable"
	case KSFNTLookupTrimmedArray:
		return "KSFNTLookupTrimmedArray"
	case KSFNTLookupVector:
		return "KSFNTLookupVector"
	default:
		return fmt.Sprintf("KSFNTLookup(%d)", e)
	}
}

type KST uint32

const (
	KSTClassDeletedGlyph KST = 2
	KSTClassEndOfLine    KST = 3
	KSTClassEndOfText    KST = 0
	KSTClassOutOfBounds  KST = 1
	KSTLigActionMask     KST = 0x3fff
	KSTMarkEnd           KST = 0x2000
	KSTNoAdvance         KST = 0x4000
	KSTRearrVerbMask     KST = 0xf
	KSTSetMark           KST = 0x8000
)

func (e KST) String() string {
	switch e {
	case KSTClassDeletedGlyph:
		return "KSTClassDeletedGlyph"
	case KSTClassEndOfLine:
		return "KSTClassEndOfLine"
	case KSTClassEndOfText:
		return "KSTClassEndOfText"
	case KSTClassOutOfBounds:
		return "KSTClassOutOfBounds"
	case KSTLigActionMask:
		return "KSTLigActionMask"
	case KSTMarkEnd:
		return "KSTMarkEnd"
	case KSTNoAdvance:
		return "KSTNoAdvance"
	case KSTRearrVerbMask:
		return "KSTRearrVerbMask"
	case KSTSetMark:
		return "KSTSetMark"
	default:
		return fmt.Sprintf("KST(%d)", e)
	}
}

const KSTKCrossStreamReset uint32 = 0x2000

const KSTXHasLigAction uint32 = 0x2000

type KShowDiacriticsSelector uint32

const (
	KDecomposeDiacriticsSelector KShowDiacriticsSelector = 2
	KHideDiacriticsSelector      KShowDiacriticsSelector = 1
	KShowDiacriticsSelectorValue KShowDiacriticsSelector = 0
)

func (e KShowDiacriticsSelector) String() string {
	switch e {
	case KDecomposeDiacriticsSelector:
		return "KDecomposeDiacriticsSelector"
	case KHideDiacriticsSelector:
		return "KHideDiacriticsSelector"
	case KShowDiacriticsSelectorValue:
		return "KShowDiacriticsSelectorValue"
	default:
		return fmt.Sprintf("KShowDiacriticsSelector(%d)", e)
	}
}

type KSubstituteVerticalForms uint32

const (
	KSubstituteVerticalFormsOffSelector KSubstituteVerticalForms = 1
	KSubstituteVerticalFormsOnSelector  KSubstituteVerticalForms = 0
)

func (e KSubstituteVerticalForms) String() string {
	switch e {
	case KSubstituteVerticalFormsOffSelector:
		return "KSubstituteVerticalFormsOffSelector"
	case KSubstituteVerticalFormsOnSelector:
		return "KSubstituteVerticalFormsOnSelector"
	default:
		return fmt.Sprintf("KSubstituteVerticalForms(%d)", e)
	}
}

type KTRAK uint32

const (
	KTRAKCurrentVersion KTRAK = 0x10000
	KTRAKTag            KTRAK = 0x7472616b
	KTRAKUniformFormat  KTRAK = 0
)

func (e KTRAK) String() string {
	switch e {
	case KTRAKCurrentVersion:
		return "KTRAKCurrentVersion"
	case KTRAKTag:
		return "KTRAKTag"
	case KTRAKUniformFormat:
		return "KTRAKUniformFormat"
	default:
		return fmt.Sprintf("KTRAK(%d)", e)
	}
}

type KTraditionalCharactersSelector uint32

const (
	KExpertCharactersSelector           KTraditionalCharactersSelector = 10
	KHojoCharactersSelector             KTraditionalCharactersSelector = 12
	KJIS1978CharactersSelector          KTraditionalCharactersSelector = 2
	KJIS1983CharactersSelector          KTraditionalCharactersSelector = 3
	KJIS1990CharactersSelector          KTraditionalCharactersSelector = 4
	KJIS2004CharactersSelector          KTraditionalCharactersSelector = 11
	KNLCCharactersSelector              KTraditionalCharactersSelector = 13
	KSimplifiedCharactersSelector       KTraditionalCharactersSelector = 1
	KTraditionalAltFiveSelector         KTraditionalCharactersSelector = 9
	KTraditionalAltFourSelector         KTraditionalCharactersSelector = 8
	KTraditionalAltOneSelector          KTraditionalCharactersSelector = 5
	KTraditionalAltThreeSelector        KTraditionalCharactersSelector = 7
	KTraditionalAltTwoSelector          KTraditionalCharactersSelector = 6
	KTraditionalCharactersSelectorValue KTraditionalCharactersSelector = 0
	KTraditionalNamesCharactersSelector KTraditionalCharactersSelector = 14
)

func (e KTraditionalCharactersSelector) String() string {
	switch e {
	case KExpertCharactersSelector:
		return "KExpertCharactersSelector"
	case KHojoCharactersSelector:
		return "KHojoCharactersSelector"
	case KJIS1978CharactersSelector:
		return "KJIS1978CharactersSelector"
	case KJIS1983CharactersSelector:
		return "KJIS1983CharactersSelector"
	case KJIS1990CharactersSelector:
		return "KJIS1990CharactersSelector"
	case KJIS2004CharactersSelector:
		return "KJIS2004CharactersSelector"
	case KNLCCharactersSelector:
		return "KNLCCharactersSelector"
	case KSimplifiedCharactersSelector:
		return "KSimplifiedCharactersSelector"
	case KTraditionalAltFiveSelector:
		return "KTraditionalAltFiveSelector"
	case KTraditionalAltFourSelector:
		return "KTraditionalAltFourSelector"
	case KTraditionalAltOneSelector:
		return "KTraditionalAltOneSelector"
	case KTraditionalAltThreeSelector:
		return "KTraditionalAltThreeSelector"
	case KTraditionalAltTwoSelector:
		return "KTraditionalAltTwoSelector"
	case KTraditionalCharactersSelectorValue:
		return "KTraditionalCharactersSelectorValue"
	case KTraditionalNamesCharactersSelector:
		return "KTraditionalNamesCharactersSelector"
	default:
		return fmt.Sprintf("KTraditionalCharactersSelector(%d)", e)
	}
}

type KUnconnectedSelector uint32

const (
	KCursiveSelector            KUnconnectedSelector = 2
	KPartiallyConnectedSelector KUnconnectedSelector = 1
	KUnconnectedSelectorValue   KUnconnectedSelector = 0
)

func (e KUnconnectedSelector) String() string {
	switch e {
	case KCursiveSelector:
		return "KCursiveSelector"
	case KPartiallyConnectedSelector:
		return "KPartiallyConnectedSelector"
	case KUnconnectedSelectorValue:
		return "KUnconnectedSelectorValue"
	default:
		return fmt.Sprintf("KUnconnectedSelector(%d)", e)
	}
}

type KUpperAndLowerCaseSelector uint32

const (
	KAllCapsSelector                 KUpperAndLowerCaseSelector = 1
	KAllLowerCaseSelector            KUpperAndLowerCaseSelector = 2
	KInitialCapsAndSmallCapsSelector KUpperAndLowerCaseSelector = 5
	KInitialCapsSelector             KUpperAndLowerCaseSelector = 4
	KSmallCapsSelector               KUpperAndLowerCaseSelector = 3
	KUpperAndLowerCaseSelectorValue  KUpperAndLowerCaseSelector = 0
)

func (e KUpperAndLowerCaseSelector) String() string {
	switch e {
	case KAllCapsSelector:
		return "KAllCapsSelector"
	case KAllLowerCaseSelector:
		return "KAllLowerCaseSelector"
	case KInitialCapsAndSmallCapsSelector:
		return "KInitialCapsAndSmallCapsSelector"
	case KInitialCapsSelector:
		return "KInitialCapsSelector"
	case KSmallCapsSelector:
		return "KSmallCapsSelector"
	case KUpperAndLowerCaseSelectorValue:
		return "KUpperAndLowerCaseSelectorValue"
	default:
		return fmt.Sprintf("KUpperAndLowerCaseSelector(%d)", e)
	}
}

type KWordInitialSwashesOnSelector uint32

const (
	KLineFinalSwashesOffSelector       KWordInitialSwashesOnSelector = 7
	KLineFinalSwashesOnSelector        KWordInitialSwashesOnSelector = 6
	KLineInitialSwashesOffSelector     KWordInitialSwashesOnSelector = 5
	KLineInitialSwashesOnSelector      KWordInitialSwashesOnSelector = 4
	KNonFinalSwashesOffSelector        KWordInitialSwashesOnSelector = 9
	KNonFinalSwashesOnSelector         KWordInitialSwashesOnSelector = 8
	KWordFinalSwashesOffSelector       KWordInitialSwashesOnSelector = 3
	KWordFinalSwashesOnSelector        KWordInitialSwashesOnSelector = 2
	KWordInitialSwashesOffSelector     KWordInitialSwashesOnSelector = 1
	KWordInitialSwashesOnSelectorValue KWordInitialSwashesOnSelector = 0
)

func (e KWordInitialSwashesOnSelector) String() string {
	switch e {
	case KLineFinalSwashesOffSelector:
		return "KLineFinalSwashesOffSelector"
	case KLineFinalSwashesOnSelector:
		return "KLineFinalSwashesOnSelector"
	case KLineInitialSwashesOffSelector:
		return "KLineInitialSwashesOffSelector"
	case KLineInitialSwashesOnSelector:
		return "KLineInitialSwashesOnSelector"
	case KNonFinalSwashesOffSelector:
		return "KNonFinalSwashesOffSelector"
	case KNonFinalSwashesOnSelector:
		return "KNonFinalSwashesOnSelector"
	case KWordFinalSwashesOffSelector:
		return "KWordFinalSwashesOffSelector"
	case KWordFinalSwashesOnSelector:
		return "KWordFinalSwashesOnSelector"
	case KWordInitialSwashesOffSelector:
		return "KWordInitialSwashesOffSelector"
	case KWordInitialSwashesOnSelectorValue:
		return "KWordInitialSwashesOnSelectorValue"
	default:
		return fmt.Sprintf("KWordInitialSwashesOnSelector(%d)", e)
	}
}

const NameFontTableTag uint32 = 0x6e616d65

const NonGlyphID uint32 = 65535

const Os2FontTableTag uint32 = 0x4f532f32

type SizeofSfntcmapencoding uint32

const (
	Sizeof_sfntCMapEncoding SizeofSfntcmapencoding = 8
)

func (e SizeofSfntcmapencoding) String() string {
	switch e {
	case Sizeof_sfntCMapEncoding:
		return "Sizeof_sfntCMapEncoding"
	default:
		return fmt.Sprintf("SizeofSfntcmapencoding(%d)", e)
	}
}

type SizeofSfntcmapextendedsubheader uint32

const (
	Sizeof_sfntCMapExtendedSubHeader SizeofSfntcmapextendedsubheader = 12
)

func (e SizeofSfntcmapextendedsubheader) String() string {
	switch e {
	case Sizeof_sfntCMapExtendedSubHeader:
		return "Sizeof_sfntCMapExtendedSubHeader"
	default:
		return fmt.Sprintf("SizeofSfntcmapextendedsubheader(%d)", e)
	}
}

type SizeofSfntcmapheader uint32

const (
	Sizeof_sfntCMapHeader SizeofSfntcmapheader = 4
)

func (e SizeofSfntcmapheader) String() string {
	switch e {
	case Sizeof_sfntCMapHeader:
		return "Sizeof_sfntCMapHeader"
	default:
		return fmt.Sprintf("SizeofSfntcmapheader(%d)", e)
	}
}

type SizeofSfntcmapsubheader uint32

const (
	Sizeof_sfntCMapSubHeader SizeofSfntcmapsubheader = 6
)

func (e SizeofSfntcmapsubheader) String() string {
	switch e {
	case Sizeof_sfntCMapSubHeader:
		return "Sizeof_sfntCMapSubHeader"
	default:
		return fmt.Sprintf("SizeofSfntcmapsubheader(%d)", e)
	}
}

type SizeofSfntdescriptorheader uint32

const (
	Sizeof_sfntDescriptorHeader SizeofSfntdescriptorheader = 8
)

func (e SizeofSfntdescriptorheader) String() string {
	switch e {
	case Sizeof_sfntDescriptorHeader:
		return "Sizeof_sfntDescriptorHeader"
	default:
		return fmt.Sprintf("SizeofSfntdescriptorheader(%d)", e)
	}
}

type SizeofSfntdirectory uint32

const (
	Sizeof_sfntDirectory SizeofSfntdirectory = 12
)

func (e SizeofSfntdirectory) String() string {
	switch e {
	case Sizeof_sfntDirectory:
		return "Sizeof_sfntDirectory"
	default:
		return fmt.Sprintf("SizeofSfntdirectory(%d)", e)
	}
}

type SizeofSfntinstance uint32

const (
	Sizeof_sfntInstance SizeofSfntinstance = 4
)

func (e SizeofSfntinstance) String() string {
	switch e {
	case Sizeof_sfntInstance:
		return "Sizeof_sfntInstance"
	default:
		return fmt.Sprintf("SizeofSfntinstance(%d)", e)
	}
}

type SizeofSfntnameheader uint32

const (
	Sizeof_sfntNameHeader SizeofSfntnameheader = 6
)

func (e SizeofSfntnameheader) String() string {
	switch e {
	case Sizeof_sfntNameHeader:
		return "Sizeof_sfntNameHeader"
	default:
		return fmt.Sprintf("SizeofSfntnameheader(%d)", e)
	}
}

type SizeofSfntnamerecord uint32

const (
	Sizeof_sfntNameRecord SizeofSfntnamerecord = 12
)

func (e SizeofSfntnamerecord) String() string {
	switch e {
	case Sizeof_sfntNameRecord:
		return "Sizeof_sfntNameRecord"
	default:
		return fmt.Sprintf("SizeofSfntnamerecord(%d)", e)
	}
}

type SizeofSfntvariationaxis uint32

const (
	Sizeof_sfntVariationAxis SizeofSfntvariationaxis = 20
)

func (e SizeofSfntvariationaxis) String() string {
	switch e {
	case Sizeof_sfntVariationAxis:
		return "Sizeof_sfntVariationAxis"
	default:
		return fmt.Sprintf("SizeofSfntvariationaxis(%d)", e)
	}
}

type SizeofSfntvariationheader uint32

const (
	Sizeof_sfntVariationHeader SizeofSfntvariationheader = 16
)

func (e SizeofSfntvariationheader) String() string {
	switch e {
	case Sizeof_sfntVariationHeader:
		return "Sizeof_sfntVariationHeader"
	default:
		return fmt.Sprintf("SizeofSfntvariationheader(%d)", e)
	}
}

const VariationFontTableTag uint32 = 0x66766172
