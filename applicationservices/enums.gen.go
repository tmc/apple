// Code generated from Apple documentation for ApplicationServices. DO NOT EDIT.

package applicationservices

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/applicationservices/atsfontfilterselector
type ATSFontFilterSelector uint32

const (
	KATSFileReferenceFilterSelector                 ATSFontFilterSelector = 10
	KATSFontFilterSelectorFontApplierFunction       ATSFontFilterSelector = 9
	KATSFontFilterSelectorFontFamily                ATSFontFilterSelector = 7
	KATSFontFilterSelectorFontFamilyApplierFunction ATSFontFilterSelector = 8
	KATSFontFilterSelectorGeneration                ATSFontFilterSelector = 3
	KATSFontFilterSelectorUnspecified               ATSFontFilterSelector = 0
)

func (e ATSFontFilterSelector) String() string {
	switch e {
	case KATSFileReferenceFilterSelector:
		return "KATSFileReferenceFilterSelector"
	case KATSFontFilterSelectorFontApplierFunction:
		return "KATSFontFilterSelectorFontApplierFunction"
	case KATSFontFilterSelectorFontFamily:
		return "KATSFontFilterSelectorFontFamily"
	case KATSFontFilterSelectorFontFamilyApplierFunction:
		return "KATSFontFilterSelectorFontFamilyApplierFunction"
	case KATSFontFilterSelectorGeneration:
		return "KATSFontFilterSelectorGeneration"
	case KATSFontFilterSelectorUnspecified:
		return "KATSFontFilterSelectorUnspecified"
	default:
		return fmt.Sprintf("ATSFontFilterSelector(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/applicationservices/atsfontnotifyaction
type ATSFontNotifyAction uint32

const (
	KATSFontNotifyActionDirectoriesChanged ATSFontNotifyAction = 2
	KATSFontNotifyActionFontsChanged       ATSFontNotifyAction = 1
)

func (e ATSFontNotifyAction) String() string {
	switch e {
	case KATSFontNotifyActionDirectoriesChanged:
		return "KATSFontNotifyActionDirectoriesChanged"
	case KATSFontNotifyActionFontsChanged:
		return "KATSFontNotifyActionFontsChanged"
	default:
		return fmt.Sprintf("ATSFontNotifyAction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/applicationservices/atsfontnotifyoption
type ATSFontNotifyOption uint32

const (
	KATSFontNotifyOptionDefault               ATSFontNotifyOption = 0
	KATSFontNotifyOptionReceiveWhileSuspended ATSFontNotifyOption = 1
)

func (e ATSFontNotifyOption) String() string {
	switch e {
	case KATSFontNotifyOptionDefault:
		return "KATSFontNotifyOptionDefault"
	case KATSFontNotifyOptionReceiveWhileSuspended:
		return "KATSFontNotifyOptionReceiveWhileSuspended"
	default:
		return fmt.Sprintf("ATSFontNotifyOption(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/applicationservices/atsfontquerymessageid
type ATSFontQueryMessageID uint32

const (
	KATSQueryActivateFontMessage ATSFontQueryMessageID = 'a'<<24 | 't'<<16 | 's'<<8 | 'a' // 'atsa'
)

func (e ATSFontQueryMessageID) String() string {
	switch e {
	case KATSQueryActivateFontMessage:
		return "KATSQueryActivateFontMessage"
	default:
		return fmt.Sprintf("ATSFontQueryMessageID(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/applicationservices/axcopymultipleattributeoptions
type AXCopyMultipleAttributeOptions uint32

const (
	StopOnError AXCopyMultipleAttributeOptions = 0
)

func (e AXCopyMultipleAttributeOptions) String() string {
	switch e {
	case StopOnError:
		return "StopOnError"
	default:
		return fmt.Sprintf("AXCopyMultipleAttributeOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/applicationservices/axerror
type AXError int

const ()

// See: https://developer.apple.com/documentation/applicationservices/axmenuitemmodifiers
type AXMenuItemModifiers uint32

const (
	Control   AXMenuItemModifiers = 0
	NoCommand AXMenuItemModifiers = 0
	Option    AXMenuItemModifiers = 0
	Shift     AXMenuItemModifiers = 0
)

func (e AXMenuItemModifiers) String() string {
	switch e {
	case Control:
		return "Control"
	default:
		return fmt.Sprintf("AXMenuItemModifiers(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/applicationservices/axpriority
type AXPriority int

const ()

// See: https://developer.apple.com/documentation/applicationservices/axunderlinestyle
type AXUnderlineStyle uint32

const ()

// See: https://developer.apple.com/documentation/applicationservices/axvaluetype
type AXValueType uint32

const ()

type At uint

const (
	AtAbsoluteCenter   At = 1
	AtBottom           At = 3
	AtBottomLeft       At = 3
	AtBottomRight      At = 3
	AtCenterBottom     At = 3
	AtCenterLeft       At = 1
	AtCenterRight      At = 1
	AtCenterTop        At = 2
	AtHorizontalCenter At = 4
	AtLeft             At = 8
	AtNone             At = 0
	AtRight            At = 12
	AtTop              At = 2
	AtTopLeft          At = 2
	AtTopRight         At = 2
	AtVerticalCenter   At = 1
)

func (e At) String() string {
	switch e {
	case AtAbsoluteCenter:
		return "AtAbsoluteCenter"
	case AtBottom:
		return "AtBottom"
	case AtCenterTop:
		return "AtCenterTop"
	case AtHorizontalCenter:
		return "AtHorizontalCenter"
	case AtLeft:
		return "AtLeft"
	case AtNone:
		return "AtNone"
	case AtRight:
		return "AtRight"
	default:
		return fmt.Sprintf("At(%d)", e)
	}
}

type BadPasteboardSyncErr int

const (
	BadPasteboardFlavorErr       BadPasteboardSyncErr = -25133
	BadPasteboardIndexErr        BadPasteboardSyncErr = -25131
	BadPasteboardItemErr         BadPasteboardSyncErr = -25132
	BadPasteboardSyncErrValue    BadPasteboardSyncErr = -25130
	DuplicatePasteboardFlavorErr BadPasteboardSyncErr = -25134
	NoPasteboardPromiseKeeperErr BadPasteboardSyncErr = -25136
	NotPasteboardOwnerErr        BadPasteboardSyncErr = -25135
)

func (e BadPasteboardSyncErr) String() string {
	switch e {
	case BadPasteboardFlavorErr:
		return "BadPasteboardFlavorErr"
	case BadPasteboardIndexErr:
		return "BadPasteboardIndexErr"
	case BadPasteboardItemErr:
		return "BadPasteboardItemErr"
	case BadPasteboardSyncErrValue:
		return "BadPasteboardSyncErrValue"
	case DuplicatePasteboardFlavorErr:
		return "DuplicatePasteboardFlavorErr"
	case NoPasteboardPromiseKeeperErr:
		return "NoPasteboardPromiseKeeperErr"
	case NotPasteboardOwnerErr:
		return "NotPasteboardOwnerErr"
	default:
		return fmt.Sprintf("BadPasteboardSyncErr(%d)", e)
	}
}

type BadTranslationRef int

const (
	BadTranslationRefErr BadTranslationRef = -3031
)

func (e BadTranslationRef) String() string {
	switch e {
	case BadTranslationRefErr:
		return "BadTranslationRefErr"
	default:
		return fmt.Sprintf("BadTranslationRef(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/applicationservices/cmfloatbitmapflags
type CMFloatBitmapFlags uint32

const (
	KCMFloatBitmapFlagsAlpha        CMFloatBitmapFlags = 1
	KCMFloatBitmapFlagsAlphaPremul  CMFloatBitmapFlags = 2
	KCMFloatBitmapFlagsNone         CMFloatBitmapFlags = 0
	KCMFloatBitmapFlagsRangeClipped CMFloatBitmapFlags = 4
)

func (e CMFloatBitmapFlags) String() string {
	switch e {
	case KCMFloatBitmapFlagsAlpha:
		return "KCMFloatBitmapFlagsAlpha"
	case KCMFloatBitmapFlagsAlphaPremul:
		return "KCMFloatBitmapFlagsAlphaPremul"
	case KCMFloatBitmapFlagsNone:
		return "KCMFloatBitmapFlagsNone"
	case KCMFloatBitmapFlagsRangeClipped:
		return "KCMFloatBitmapFlagsRangeClipped"
	default:
		return fmt.Sprintf("CMFloatBitmapFlags(%d)", e)
	}
}

type Cdev int

const (
	CdevGenErr Cdev = -1
	CdevMemErr Cdev = 0
	CdevResErr Cdev = 1
	CdevUnset  Cdev = 3
)

func (e Cdev) String() string {
	switch e {
	case CdevGenErr:
		return "CdevGenErr"
	case CdevMemErr:
		return "CdevMemErr"
	case CdevResErr:
		return "CdevResErr"
	case CdevUnset:
		return "CdevUnset"
	default:
		return fmt.Sprintf("Cdev(%d)", e)
	}
}

type Cm int

const (
	Cm10CLRData        Cm = 'A'<<24 | 'C'<<16 | 'L'<<8 | 'R' // 'ACLR'
	Cm11CLRData        Cm = 'B'<<24 | 'C'<<16 | 'L'<<8 | 'R' // 'BCLR'
	Cm12CLRData        Cm = 'C'<<24 | 'C'<<16 | 'L'<<8 | 'R' // 'CCLR'
	Cm13CLRData        Cm = 'D'<<24 | 'C'<<16 | 'L'<<8 | 'R' // 'DCLR'
	Cm14CLRData        Cm = 'E'<<24 | 'C'<<16 | 'L'<<8 | 'R' // 'ECLR'
	Cm15CLRData        Cm = 'F'<<24 | 'C'<<16 | 'L'<<8 | 'R' // 'FCLR'
	Cm16_8ColorPacking Cm = 0x2000
	// Cm24_8ColorPacking: The color values for three 8-bit color channels are stored in consecutive bytes, for a total of 24 bits.
	Cm24_8ColorPacking Cm = 0x2100
	// Cm32_16ColorPacking: The color values for two 16-bit color channels are stored in a 32-bit word.
	Cm32_16ColorPacking Cm = 0x2600
	// Cm32_32ColorPacking: The color value for a 32-bit color channel is stored in a 32-bit word.
	Cm32_32ColorPacking Cm = 0x2700
	// Cm32_8ColorPacking: The color values for four 8-bit color channels are stored in consecutive bytes, for a total of 32 bits.
	Cm32_8ColorPacking Cm = 2048
	Cm3CLRData         Cm = '3'<<24 | 'C'<<16 | 'L'<<8 | 'R' // '3CLR'
	// Cm40_8ColorPacking: The color values for five 8-bit color channels are stored in consecutive bytes, for a total of 40 bits.
	Cm40_8ColorPacking Cm = 0x2200
	// Cm48_16ColorPacking: The color values for three 16-bit color channels are stored in 48 consecutive bits.
	Cm48_16ColorPacking Cm = 0x2900
	// Cm48_8ColorPacking: The color values for six 8-bit color channels are stored in consecutive bytes, for a total of 48 bits.
	Cm48_8ColorPacking Cm = 0x2300
	Cm4CLRData         Cm = '4'<<24 | 'C'<<16 | 'L'<<8 | 'R' // '4CLR'
	// Cm56_8ColorPacking: The color values for seven 8-bit color channels are stored in consecutive bytes, for a total of 56 bits.
	Cm56_8ColorPacking Cm = 0x2400
	Cm5CLRData         Cm = '5'<<24 | 'C'<<16 | 'L'<<8 | 'R' // '5CLR'
	// Cm64_16ColorPacking: The color values for four 16-bit color channels are stored in 64 consecutive bits.
	Cm64_16ColorPacking Cm = 0x2a00
	// Cm64_8ColorPacking: The color values for eight 8-bit color channels are stored in consecutive bytes, for a total of 64 bits.
	Cm64_8ColorPacking Cm = 0x2500
	Cm6CLRData         Cm = '6'<<24 | 'C'<<16 | 'L'<<8 | 'R' // '6CLR'
	Cm7CLRData         Cm = '7'<<24 | 'C'<<16 | 'L'<<8 | 'R' // '7CLR'
	Cm8CLRData         Cm = '8'<<24 | 'C'<<16 | 'L'<<8 | 'R' // '8CLR'
	Cm8_8ColorPacking  Cm = 0x2800
	Cm9CLRData         Cm = '9'<<24 | 'C'<<16 | 'L'<<8 | 'R' // '9CLR'
	CmARGB32PmulSpace  Cm = 1
	// CmARGB32Space: # Discussion
	CmARGB32Space      Cm = 1
	CmARGB64LPmulSpace Cm = 1
	CmARGB64LSpace     Cm = 1
	CmARGB64PmulSpace  Cm = 1
	CmARGB64Space      Cm = 1
	CmAToB0Tag         Cm = 'A'<<24 | '2'<<16 | 'B'<<8 | '0' // 'A2B0'
	CmAToB1Tag         Cm = 'A'<<24 | '2'<<16 | 'B'<<8 | '1' // 'A2B1'
	CmAToB2Tag         Cm = 'A'<<24 | '2'<<16 | 'B'<<8 | '2' // 'A2B2'
	// CmAbortWriteAccess: Cancel the current write attempt.
	CmAbortWriteAccess Cm = 7
	// CmAbsoluteColorimetric: This approach is based on a device-independent color space in which the result is an idealized print viewed on an ideal type of paper having a large dynamic range and color gamut.
	CmAbsoluteColorimetric Cm = 3
	// CmAbstractClass: An abstract profile.
	CmAbstractClass Cm = 'a'<<24 | 'b'<<16 | 's'<<8 | 't' // 'abst'
	// CmAlphaFirstPacking: An alpha channel is added to the color value as its first component.
	CmAlphaFirstPacking Cm = 0x1000
	CmAlphaLastPacking  Cm = 0
	// CmAlphaPmulSpace: A premultiplied alpha channel component is added to the color value.
	CmAlphaPmulSpace Cm = 0x40
	// CmAlphaSpace: An alpha channel component is added to the color value.
	CmAlphaSpace Cm = 0x80
	// CmAsciiData: ASCII data.
	CmAsciiData Cm = 0
	CmBToA0Tag  Cm = 'B'<<24 | '2'<<16 | 'A'<<8 | '0' // 'B2A0'
	CmBToA1Tag  Cm = 'B'<<24 | '2'<<16 | 'A'<<8 | '1' // 'B2A1'
	CmBToA2Tag  Cm = 'B'<<24 | '2'<<16 | 'A'<<8 | '2' // 'B2A2'
	// CmBeginAccess: Begin the process of procedural access.
	CmBeginAccess Cm = 8
	// CmBestMode: Best mode indicates that the CMM should maximize resource usage to ensure the highest possible quality.
	CmBestMode Cm = 2
	// CmBinaryData: Binary data.
	CmBinaryData                  Cm = 1
	CmBlackPointCompensationMask  Cm = 0x4
	CmBlueColorantTag             Cm = 'b'<<24 | 'X'<<16 | 'Y'<<8 | 'Z' // 'bXYZ'
	CmBlueTRCTag                  Cm = 'b'<<24 | 'T'<<16 | 'R'<<8 | 'C' // 'bTRC'
	CmBradfordChromaticAdaptation Cm = 3
	CmBufferBasedProfile          Cm = 6
	// CmCMSReservedFlagsMask: # Discussion
	CmCMSReservedFlagsMask Cm = 0xffff0000
	// CmCMYData: The CMY data color space.
	CmCMYData Cm = 'C'<<24 | 'M'<<16 | 'Y'<<8 | ' ' // 'CMY '
	// CmCMYK32Space: A CMYK color space composed of cyan, magenta, yellow, and black components whose values are packed with 8 bits of storage per component.
	CmCMYK32Space  Cm = 2
	CmCMYK64LSpace Cm = 2
	// CmCMYK64Space: A CMYK color space composed of cyan, magenta, yellow, and black components whose values are packed with 16 bits of storage per component.
	CmCMYK64Space Cm = 2
	// CmCMYKData: The CMYK data color space.
	CmCMYKData Cm = 'C'<<24 | 'M'<<16 | 'Y'<<8 | 'K' // 'CMYK'
	// CmCMYKSpace: A CMYK color space composed of cyan, magenta, yellow, and black.
	CmCMYKSpace              Cm = 0x2
	CmCS1ProfileVersion      Cm = 0x100
	CmCS2ProfileVersion      Cm = 33554432
	CmCalibrationDateTimeTag Cm = 'c'<<24 | 'a'<<16 | 'l'<<8 | 't' // 'calt'
	CmCameraDeviceClass      Cm = 'c'<<24 | 'm'<<16 | 'r'<<8 | 'a' // 'cmra'
	CmCharTargetTag          Cm = 't'<<24 | 'a'<<16 | 'r'<<8 | 'g' // 'targ'
	CmChromaticAdaptationTag Cm = 'c'<<24 | 'h'<<16 | 'a'<<8 | 'd' // 'chad'
	// CmCloseAccess: Close the profile for reading or writing.
	CmCloseAccess Cm = 5
	// CmCloseSpool: Directs the function to complete the data transfer.
	CmCloseSpool Cm = 5
	// CmColorSpaceClass: A color space profile.
	CmColorSpaceClass Cm = 's'<<24 | 'p'<<16 | 'a'<<8 | 'c' // 'spac'
	CmCopyrightTag    Cm = 'c'<<24 | 'p'<<16 | 'r'<<8 | 't' // 'cprt'
	// CmCreateNewAccess: Create a new data stream for the profile.
	CmCreateNewAccess            Cm = 6
	CmCurrentProfileLocationSize Cm = 0
	CmCurrentProfileMajorVersion Cm = 0x2000000
	// CmDeviceAlreadyRegistered: Device already registered; returned by a CM device integration routine.
	CmDeviceAlreadyRegistered Cm = -4228
	// CmDeviceDBNotFoundErr: Preferences not found or loaded; returned by a CM device integration routine.
	CmDeviceDBNotFoundErr Cm = -4227
	CmDeviceMfgDescTag    Cm = 'd'<<24 | 'm'<<16 | 'n'<<8 | 'd' // 'dmnd'
	CmDeviceModelDescTag  Cm = 'd'<<24 | 'm'<<16 | 'd'<<8 | 'd' // 'dmdd'
	// CmDeviceNotRegistered: Device not found; returned by a CM device integration routine.
	CmDeviceNotRegistered Cm = -4229
	// CmDeviceProfilesNotFound: Profiles not found; returned by a CM device integration routine.
	CmDeviceProfilesNotFound Cm = -4230
	// CmDisplayClass: A display device profile defined for a monitor.
	CmDisplayClass       Cm = 'm'<<24 | 'n'<<16 | 't'<<8 | 'r' // 'mntr'
	CmDisplayDeviceClass Cm = 'm'<<24 | 'n'<<16 | 't'<<8 | 'r' // 'mntr'
	CmDisplayUse         Cm = 'd'<<24 | 'p'<<16 | 'l'<<8 | 'y' // 'dply'
	// CmDraftMode: Draft mode indicates that the CMM should sacrifice quality, if necessary, to minimize resource requirements.
	CmDraftMode Cm = 1
	// CmEmbeddedMask: This mask provides access to bit 0 of the `flags` field, which specifies whether the profile is embedded.
	CmEmbeddedMask Cm = 0x1
	// CmEmbeddedUseMask: # Discussion
	CmEmbeddedUseMask Cm = 0x2
	// CmEndAccess: End the process of procedural access.
	CmEndAccess Cm = 9
	CmFlare0    Cm = 0
	CmFlare100  Cm = 0x1
	// CmGamutCheckingMask: # Discussion
	CmGamutCheckingMask Cm = 0x80000
	// CmGamutResult1Space: # Discussion
	CmGamutResult1Space Cm = 2816
	// CmGamutResultSpace: # Discussion
	CmGamutResultSpace Cm = 0xc
	CmGamutTag         Cm = 'g'<<24 | 'a'<<16 | 'm'<<8 | 't' // 'gamt'
	CmGeometry045or450 Cm = 0x1
	CmGeometry0dord0   Cm = 0x2
	CmGeometryUnknown  Cm = 0
	// CmGlossy: If the bit 1 of the associated mask is `0` then glossy; if `1` then matte.
	CmGlossy Cm = 1
	// CmGlossyMatteMask: # Discussion
	CmGlossyMatteMask Cm = 0x2
	CmGray16LSpace    Cm = 10
	// CmGray16Space: A luminance color space with a single 16-bit component, gray.
	CmGray16Space       Cm = 10
	CmGray8Space        Cm = 10
	CmGrayA16PmulSpace  Cm = 10
	CmGrayA16Space      Cm = 10
	CmGrayA32LPmulSpace Cm = 10
	CmGrayA32LSpace     Cm = 10
	CmGrayA32PmulSpace  Cm = 10
	// CmGrayA32Space: A luminance color space with two components, a gray component followed by an alpha channel component.
	CmGrayA32Space   Cm = 10
	CmGrayAPmulSpace Cm = 10
	// CmGrayASpace: A luminance color space with two components, a gray component followed by an alpha channel component.
	CmGrayASpace Cm = 10
	// CmGrayData: The Gray data color space.
	CmGrayData Cm = 'G'<<24 | 'R'<<16 | 'A'<<8 | 'Y' // 'GRAY'
	// CmGraySpace: A luminance color space with a single component, gray.
	CmGraySpace        Cm = 0xa
	CmGrayTRCTag       Cm = 'k'<<24 | 'T'<<16 | 'R'<<8 | 'C' // 'kTRC'
	CmGreenColorantTag Cm = 'g'<<24 | 'X'<<16 | 'Y'<<8 | 'Z' // 'gXYZ'
	CmGreenTRCTag      Cm = 'g'<<24 | 'T'<<16 | 'R'<<8 | 'C' // 'gTRC'
	// CmHLS32Space: An HLS color space composed of hue, lightness, and saturation components whose values are packed with 10 bits of storage per component.
	CmHLS32Space Cm = 4
	// CmHLSData: The HLS data color space.
	CmHLSData Cm = 'H'<<24 | 'L'<<16 | 'S'<<8 | ' ' // 'HLS '
	// CmHLSSpace: An HLS color space composed of hue, lightness, and saturation components.
	CmHLSSpace Cm = 0x4
	// CmHSV32Space: An HSV color space composed of hue, saturation, and value components whose values are packed with 10 bits of storage per component.
	CmHSV32Space Cm = 3
	// CmHSVData: The HSV data color space.
	CmHSVData Cm = 'H'<<24 | 'S'<<16 | 'V'<<8 | ' ' // 'HSV '
	// CmHSVSpace: An HSV color space composed of hue, saturation, and value components.
	CmHSVSpace            Cm = 0x3
	CmICCProfileVersion2  Cm = 0x2000000
	CmICCProfileVersion21 Cm = 0x2100000
	CmICCProfileVersion4  Cm = 0x4000000
	// CmICCReservedFlagsMask: # Discussion
	CmICCReservedFlagsMask Cm = 0xffff
	// CmInputClass: An input device profile defined for a scanner.
	CmInputClass Cm = 's'<<24 | 'c'<<16 | 'n'<<8 | 'r' // 'scnr'
	CmInputUse   Cm = 'i'<<24 | 'n'<<16 | 'p'<<8 | 't' // 'inpt'
	// CmInternalCFErr: CoreFoundation failure; returned by a CM device integration routine.
	CmInternalCFErr Cm = -4231
	// CmInterpolationMask: # Discussion
	CmInterpolationMask Cm = 0x40000
	// CmLAB24Space: # Discussion
	CmLAB24Space Cm = 8
	// CmLAB32Space: # Discussion
	CmLAB32Space  Cm = 8
	CmLAB48LSpace Cm = 8
	// CmLAB48Space: # Discussion
	CmLAB48Space Cm = 8
	// CmLABSpace: An L*a*b* color space composed of L*, a*, b* components.
	CmLABSpace Cm = 0x8
	// CmLUV32Space: An L*u*v* color space composed of L*, u*, and v* components whose values are packed with 10 bits per component.
	CmLUV32Space Cm = 7
	// CmLUVSpace: An L*u*v* color space composed of L*, u*, and v* components.
	CmLUVSpace Cm = 0x7
	// CmLabData: The L*a*b* data color space.
	CmLabData                   Cm = 'L'<<24 | 'a'<<16 | 'b'<<8 | ' ' // 'Lab '
	CmLinearChromaticAdaptation Cm = 1
	// CmLinesPer: Lines per unit; can have an associated value of `0` for lines per centimeter or `1` for lines per inch.
	CmLinesPer Cm = 1
	// CmLinkClass: A device link profile.
	CmLinkClass           Cm = 'l'<<24 | 'i'<<16 | 'n'<<8 | 'k' // 'link'
	CmLittleEndianPacking Cm = 0x4000
	// CmLong10ColorPacking: The color values for three 10-bit color channels are stored consecutively in a 32-bit long, with the two highest order bits unused.
	CmLong10ColorPacking Cm = 0xa00
	// CmLong8ColorPacking: # Discussion
	CmLong8ColorPacking Cm = 0x800
	CmLuminanceTag      Cm = 'l'<<24 | 'u'<<16 | 'm'<<8 | 'i' // 'lumi'
	// CmLuvData: The L*u*v* data color space.
	CmLuvData Cm = 'L'<<24 | 'u'<<16 | 'v'<<8 | ' ' // 'Luv '
	// CmMCEight8Space: An eight-channel multichannel (HiFi) data color space, whose values are packed with 8 bits per component.
	CmMCEight8Space Cm = 9472
	// CmMCEightSpace: An eight-channel multichannel (HiFi) data color space.
	CmMCEightSpace Cm = 0x14
	// CmMCFive8Space: A five-channel multichannel (HiFi) data color space, whose values are packed with 8 bits per component.
	CmMCFive8Space Cm = 8704
	// CmMCFiveSpace: A five-channel multichannel (HiFi) data color space.
	CmMCFiveSpace Cm = 0x11
	// CmMCH5Data: The five-channel multichannel (HiFi) data color space.
	CmMCH5Data Cm = 'M'<<24 | 'C'<<16 | 'H'<<8 | '5' // 'MCH5'
	// CmMCH6Data: The six-channel multichannel (HiFi) data color space.
	CmMCH6Data Cm = 'M'<<24 | 'C'<<16 | 'H'<<8 | '6' // 'MCH6'
	// CmMCH7Data: The seven-channel multichannel (HiFi) data color space.
	CmMCH7Data Cm = 'M'<<24 | 'C'<<16 | 'H'<<8 | '7' // 'MCH7'
	// CmMCH8Data: The eight-channel multichannel (HiFi) data color space.
	CmMCH8Data Cm = 'M'<<24 | 'C'<<16 | 'H'<<8 | '8' // 'MCH8'
	// CmMCSeven8Space: A seven-channel multichannel (HiFi) data color space, whose values are packed with 8 bits per component.
	CmMCSeven8Space Cm = 9216
	// CmMCSevenSpace: A seven-channel multichannel (HiFi) data color space.
	CmMCSevenSpace Cm = 0x13
	// CmMCSix8Space: A six-channel multichannel (HiFi) data color space, whose values are packed with 8 bits per component.
	CmMCSix8Space Cm = 8960
	// CmMCSixSpace: A six-channel multichannel (HiFi) data color space.
	CmMCSixSpace         Cm = 0x12
	CmMacintosh          Cm = 'A'<<24 | 'P'<<16 | 'P'<<8 | 'L' // 'APPL'
	CmMakeAndModelTag    Cm = 'm'<<24 | 'm'<<16 | 'o'<<8 | 'd' // 'mmod'
	CmMeasurementTag     Cm = 'm'<<24 | 'e'<<16 | 'a'<<8 | 's' // 'meas'
	CmMediaBlackPointTag Cm = 'b'<<24 | 'k'<<16 | 'p'<<8 | 't' // 'bkpt'
	CmMediaWhitePointTag Cm = 'w'<<24 | 't'<<16 | 'p'<<8 | 't' // 'wtpt'
	CmMicrosoft          Cm = 'M'<<24 | 'S'<<16 | 'F'<<8 | 'T' // 'MSFT'
	CmNamedColor2Tag     Cm = 'n'<<24 | 'c'<<16 | 'l'<<8 | '2' // 'ncl2'
	// CmNamedColorClass: A named color space profile.
	CmNamedColorClass      Cm = 'n'<<24 | 'm'<<16 | 'c'<<8 | 'l' // 'nmcl'
	CmNamedColorTag        Cm = 'n'<<24 | 'c'<<16 | 'o'<<8 | 'l' // 'ncol'
	CmNamedData            Cm = 'N'<<24 | 'A'<<16 | 'M'<<8 | 'E' // 'NAME'
	CmNamedIndexed32LSpace Cm = 9984
	// CmNamedIndexed32Space: A color space where each color is stored as a single 32-bit value, specifying an index into a named color space.
	CmNamedIndexed32Space Cm = 9984
	// CmNamedIndexedSpace: A named indexed color space.
	CmNamedIndexedSpace    Cm = 0x10
	CmNativeDisplayInfoTag Cm = 'n'<<24 | 'd'<<16 | 'i'<<8 | 'n' // 'ndin'
	// CmNoColorPacking: This constant is not used for ColorSync bitmaps.
	CmNoColorPacking Cm = 0
	// CmNoProfileBase: The profile is temporary.
	CmNoProfileBase Cm = 0
	// CmNoSpace: The ColorSync Manager does not use this constant.
	CmNoSpace Cm = 0
	// CmNormalMode: This is the default setting.
	CmNormalMode Cm = 0
	// CmOneBitDirectPacking: One bit is used as the pixel format.
	CmOneBitDirectPacking Cm = 0xb00
	CmOpenReadAccess      Cm = 1
	// CmOpenReadSpool: Directs the function to begin the process of reading data.
	CmOpenReadSpool Cm = 1
	// CmOpenWriteAccess: Open the profile for writing.
	CmOpenWriteAccess Cm = 2
	// CmOpenWriteSpool: Directs the function to begin the process of writing data.
	CmOpenWriteSpool              Cm = 2
	CmOriginalProfileLocationSize Cm = 72
	// CmOutputClass: An output device profile defined for a printer.
	CmOutputClass           Cm = 'p'<<24 | 'r'<<16 | 't'<<8 | 'r' // 'prtr'
	CmOutputUse             Cm = 'o'<<24 | 'u'<<16 | 't'<<8 | 'p' // 'outp'
	CmPS2CRD0Tag            Cm = 'p'<<24 | 's'<<16 | 'd'<<8 | '0' // 'psd0'
	CmPS2CRD1Tag            Cm = 'p'<<24 | 's'<<16 | 'd'<<8 | '1' // 'psd1'
	CmPS2CRD2Tag            Cm = 'p'<<24 | 's'<<16 | 'd'<<8 | '2' // 'psd2'
	CmPS2CRD3Tag            Cm = 'p'<<24 | 's'<<16 | 'd'<<8 | '3' // 'psd3'
	CmPS2CRDVMSizeTag       Cm = 'p'<<24 | 's'<<16 | 'v'<<8 | 'm' // 'psvm'
	CmPS2CSATag             Cm = 'p'<<24 | 's'<<16 | '2'<<8 | 's' // 'ps2s'
	CmPS2RenderingIntentTag Cm = 'p'<<24 | 's'<<16 | '2'<<8 | 'i' // 'ps2i'
	CmPathBasedProfile      Cm = 5
	// CmPerceptual: All the colors of a given gamut can be scaled to fit within another gamut.
	CmPerceptual              Cm = 0
	CmPrefsSynchError         Cm = -4232
	CmPreview0Tag             Cm = 'p'<<24 | 'r'<<16 | 'e'<<8 | '0' // 'pre0'
	CmPreview1Tag             Cm = 'p'<<24 | 'r'<<16 | 'e'<<8 | '1' // 'pre1'
	CmPreview2Tag             Cm = 'p'<<24 | 'r'<<16 | 'e'<<8 | '2' // 'pre2'
	CmPrinterDeviceClass      Cm = 'p'<<24 | 'r'<<16 | 't'<<8 | 'r' // 'prtr'
	CmProfileDescriptionMLTag Cm = 'd'<<24 | 's'<<16 | 'c'<<8 | 'm' // 'dscm'
	CmProfileDescriptionTag   Cm = 'd'<<24 | 'e'<<16 | 's'<<8 | 'c' // 'desc'
	CmProfileMajorVersionMask Cm = 0xff000000
	CmProfileSequenceDescTag  Cm = 'p'<<24 | 's'<<16 | 'e'<<8 | 'q' // 'pseq'
	CmProofDeviceClass        Cm = 'p'<<24 | 'r'<<16 | 'u'<<8 | 'f' // 'pruf'
	CmProofUse                Cm = 'p'<<24 | 'r'<<16 | 'u'<<8 | 'f' // 'pruf'
	// CmPrtrDefaultScreens: Use printer default screens; can have an associated value of `0` for `false` or `1` for `true`.
	CmPrtrDefaultScreens Cm = 0
	// CmQualityMask: # Discussion
	CmQualityMask Cm = 0x30000
	CmRGB16LSpace Cm = 1
	// CmRGB16Space: An RGB color space composed of red, green, and blue components whose values are packed with 5 bits of storage per component.
	CmRGB16Space Cm = 1
	// CmRGB24Space: An RGB color space composed of red, green, and blue components whose values are packed with 8 bits of storage per component.
	CmRGB24Space Cm = 1
	// CmRGB32Space: An RGB color space composed of red, green, and blue components whose values are packed with 8 bits of storage per component.
	CmRGB32Space  Cm = 1
	CmRGB48LSpace Cm = 1
	// CmRGB48Space: An RGB color space composed of red, green, and blue components whose values are packed with 16 bits of storage per component.
	CmRGB48Space      Cm = 1
	CmRGB565LSpace    Cm = 1
	CmRGB565Space     Cm = 1
	CmRGBA32PmulSpace Cm = 1
	// CmRGBA32Space: An RGB color space composed of red, green, and blue color value components, followed by an alpha channel component.
	CmRGBA32Space      Cm = 1
	CmRGBA64LPmulSpace Cm = 1
	CmRGBA64LSpace     Cm = 1
	CmRGBA64PmulSpace  Cm = 1
	CmRGBA64Space      Cm = 1
	CmRGBAPmulSpace    Cm = 1
	// CmRGBASpace: # Discussion
	CmRGBASpace Cm = 1
	// CmRGBData: The RGB data color space.
	CmRGBData Cm = 'R'<<24 | 'G'<<16 | 'B'<<8 | ' ' // 'RGB '
	// CmRGBSpace: An RGB color space composed of red, green, and blue components.
	CmRGBSpace Cm = 0x1
	// CmReadAccess: Read the number of bytes specified by the `size` parameter.
	CmReadAccess Cm = 3
	// CmReadSpool: Directs the function to read the number of bytes specified by the [CMFlattenProcPtr] function’s `size` parameter.
	CmReadSpool      Cm = 3
	CmRedColorantTag Cm = 'r'<<24 | 'X'<<16 | 'Y'<<8 | 'Z' // 'rXYZ'
	CmRedTRCTag      Cm = 'r'<<24 | 'T'<<16 | 'R'<<8 | 'C' // 'rTRC'
	// CmReflective: If the bit 0 of the associated mask is `0` then reflective media; if `1` then transparency media.
	CmReflective Cm = 0
	// CmReflectiveTransparentMask: # Discussion
	CmReflectiveTransparentMask Cm = 0x1
	// CmRelativeColorimetric: The colors that fall within the gamuts of both devices are left unchanged.
	CmRelativeColorimetric Cm = 1
	// CmReservedSpace1: This field is reserved for use by QuickDraw GX.
	CmReservedSpace1 Cm = 0x9
	// CmReservedSpace2: This field is reserved for use by QuickDraw GX.
	CmReservedSpace2        Cm = 0xb
	CmReverseChannelPacking Cm = 0x8000
	CmSRGBData              Cm = 's'<<24 | 'R'<<16 | 'G'<<8 | 'B' // 'sRGB'
	// CmSaturation: The relative saturation of colors is maintained from gamut to gamut.
	CmSaturation                    Cm = 2
	CmScannerDeviceClass            Cm = 's'<<24 | 'c'<<16 | 'n'<<8 | 'r' // 'scnr'
	CmScreeningDescTag              Cm = 's'<<24 | 'c'<<16 | 'r'<<8 | 'd' // 'scrd'
	CmScreeningTag                  Cm = 's'<<24 | 'c'<<16 | 'r'<<8 | 'n' // 'scrn'
	CmSiliconGraphics               Cm = 'S'<<24 | 'G'<<16 | 'I'<<8 | ' ' // 'SGI '
	CmSolaris                       Cm = 'S'<<24 | 'U'<<16 | 'N'<<8 | 'W' // 'SUNW'
	CmStdobs1931TwoDegrees          Cm = 0x1
	CmStdobs1964TenDegrees          Cm = 0x2
	CmStdobsUnknown                 Cm = 0
	CmTaligent                      Cm = 'T'<<24 | 'G'<<16 | 'N'<<8 | 'T' // 'TGNT'
	CmTechnologyTag                 Cm = 't'<<24 | 'e'<<16 | 'c'<<8 | 'h' // 'tech'
	CmUcrBgTag                      Cm = 'b'<<24 | 'f'<<16 | 'd'<<8 | ' ' // 'bfd '
	CmUseDefaultChromaticAdaptation Cm = 0
	// CmVideoCardGammaTag: # Discussion
	CmVideoCardGammaTag           Cm = 'v'<<24 | 'c'<<16 | 'g'<<8 | 't' // 'vcgt'
	CmViewingConditionsDescTag    Cm = 'v'<<24 | 'u'<<16 | 'e'<<8 | 'd' // 'vued'
	CmViewingConditionsTag        Cm = 'v'<<24 | 'i'<<16 | 'e'<<8 | 'w' // 'view'
	CmVonKriesChromaticAdaptation Cm = 2
	CmWord565ColorPacking         Cm = 0x600
	// CmWord5ColorPacking: The color values for three 5-bit color channels are stored consecutively in 16-bits, with the highest order bit unused.
	CmWord5ColorPacking Cm = 0x500
	// CmWriteAccess: Write the number of bytes specified by the `size` parameter.
	CmWriteAccess Cm = 4
	// CmWriteSpool: Directs the function to write the number of bytes specified by the [CMFlattenProcPtr] function’s `size` parameter.
	CmWriteSpool Cm = 4
	CmXYZ24Space Cm = 6
	// CmXYZ32Space: An XYZ color space composed of X, Y, and Z components whose values are packed with 10 bits per component.
	CmXYZ32Space  Cm = 6
	CmXYZ48LSpace Cm = 6
	CmXYZ48Space  Cm = 6
	// CmXYZData: The XYZ data color space.
	CmXYZData Cm = 'X'<<24 | 'Y'<<16 | 'Z'<<8 | ' ' // 'XYZ '
	// CmXYZSpace: An XYZ color space composed of X, Y, and Z components.
	CmXYZSpace  Cm = 0x6
	CmYCbCrData Cm = 'Y'<<24 | 'C'<<16 | 'b'<<8 | 'r' // 'YCbr'
	// CmYXY32Space: A Yxy color space composed of Y, x, and y components whose values are packed with 10 bits of storage per component.
	CmYXY32Space Cm = 5
	// CmYXYSpace: A Yxy color space composed of Y, x, and y components.
	CmYXYSpace Cm = 0x5
	// CmYxyData: The Yxy data color space.
	CmYxyData Cm = 'Y'<<24 | 'x'<<16 | 'y'<<8 | ' ' // 'Yxy '
)

func (e Cm) String() string {
	switch e {
	case Cm10CLRData:
		return "Cm10CLRData"
	case Cm11CLRData:
		return "Cm11CLRData"
	case Cm12CLRData:
		return "Cm12CLRData"
	case Cm13CLRData:
		return "Cm13CLRData"
	case Cm14CLRData:
		return "Cm14CLRData"
	case Cm15CLRData:
		return "Cm15CLRData"
	case Cm16_8ColorPacking:
		return "Cm16_8ColorPacking"
	case Cm24_8ColorPacking:
		return "Cm24_8ColorPacking"
	case Cm32_16ColorPacking:
		return "Cm32_16ColorPacking"
	case Cm32_32ColorPacking:
		return "Cm32_32ColorPacking"
	case Cm32_8ColorPacking:
		return "Cm32_8ColorPacking"
	case Cm3CLRData:
		return "Cm3CLRData"
	case Cm40_8ColorPacking:
		return "Cm40_8ColorPacking"
	case Cm48_16ColorPacking:
		return "Cm48_16ColorPacking"
	case Cm48_8ColorPacking:
		return "Cm48_8ColorPacking"
	case Cm4CLRData:
		return "Cm4CLRData"
	case Cm56_8ColorPacking:
		return "Cm56_8ColorPacking"
	case Cm5CLRData:
		return "Cm5CLRData"
	case Cm64_16ColorPacking:
		return "Cm64_16ColorPacking"
	case Cm64_8ColorPacking:
		return "Cm64_8ColorPacking"
	case Cm6CLRData:
		return "Cm6CLRData"
	case Cm7CLRData:
		return "Cm7CLRData"
	case Cm8CLRData:
		return "Cm8CLRData"
	case Cm8_8ColorPacking:
		return "Cm8_8ColorPacking"
	case Cm9CLRData:
		return "Cm9CLRData"
	case CmARGB32PmulSpace:
		return "CmARGB32PmulSpace"
	case CmAToB0Tag:
		return "CmAToB0Tag"
	case CmAToB1Tag:
		return "CmAToB1Tag"
	case CmAToB2Tag:
		return "CmAToB2Tag"
	case CmAbortWriteAccess:
		return "CmAbortWriteAccess"
	case CmAbsoluteColorimetric:
		return "CmAbsoluteColorimetric"
	case CmAbstractClass:
		return "CmAbstractClass"
	case CmAlphaFirstPacking:
		return "CmAlphaFirstPacking"
	case CmAlphaLastPacking:
		return "CmAlphaLastPacking"
	case CmAlphaPmulSpace:
		return "CmAlphaPmulSpace"
	case CmAlphaSpace:
		return "CmAlphaSpace"
	case CmBToA0Tag:
		return "CmBToA0Tag"
	case CmBToA1Tag:
		return "CmBToA1Tag"
	case CmBToA2Tag:
		return "CmBToA2Tag"
	case CmBeginAccess:
		return "CmBeginAccess"
	case CmBestMode:
		return "CmBestMode"
	case CmBlackPointCompensationMask:
		return "CmBlackPointCompensationMask"
	case CmBlueColorantTag:
		return "CmBlueColorantTag"
	case CmBlueTRCTag:
		return "CmBlueTRCTag"
	case CmBufferBasedProfile:
		return "CmBufferBasedProfile"
	case CmCMSReservedFlagsMask:
		return "CmCMSReservedFlagsMask"
	case CmCMYData:
		return "CmCMYData"
	case CmCMYKData:
		return "CmCMYKData"
	case CmCS1ProfileVersion:
		return "CmCS1ProfileVersion"
	case CmCS2ProfileVersion:
		return "CmCS2ProfileVersion"
	case CmCalibrationDateTimeTag:
		return "CmCalibrationDateTimeTag"
	case CmCameraDeviceClass:
		return "CmCameraDeviceClass"
	case CmCharTargetTag:
		return "CmCharTargetTag"
	case CmChromaticAdaptationTag:
		return "CmChromaticAdaptationTag"
	case CmCloseAccess:
		return "CmCloseAccess"
	case CmColorSpaceClass:
		return "CmColorSpaceClass"
	case CmCopyrightTag:
		return "CmCopyrightTag"
	case CmDeviceAlreadyRegistered:
		return "CmDeviceAlreadyRegistered"
	case CmDeviceDBNotFoundErr:
		return "CmDeviceDBNotFoundErr"
	case CmDeviceMfgDescTag:
		return "CmDeviceMfgDescTag"
	case CmDeviceModelDescTag:
		return "CmDeviceModelDescTag"
	case CmDeviceNotRegistered:
		return "CmDeviceNotRegistered"
	case CmDeviceProfilesNotFound:
		return "CmDeviceProfilesNotFound"
	case CmDisplayClass:
		return "CmDisplayClass"
	case CmDisplayUse:
		return "CmDisplayUse"
	case CmEndAccess:
		return "CmEndAccess"
	case CmGamutCheckingMask:
		return "CmGamutCheckingMask"
	case CmGamutResult1Space:
		return "CmGamutResult1Space"
	case CmGamutResultSpace:
		return "CmGamutResultSpace"
	case CmGamutTag:
		return "CmGamutTag"
	case CmGray16LSpace:
		return "CmGray16LSpace"
	case CmGrayData:
		return "CmGrayData"
	case CmGrayTRCTag:
		return "CmGrayTRCTag"
	case CmGreenColorantTag:
		return "CmGreenColorantTag"
	case CmGreenTRCTag:
		return "CmGreenTRCTag"
	case CmHLSData:
		return "CmHLSData"
	case CmHSVData:
		return "CmHSVData"
	case CmICCProfileVersion21:
		return "CmICCProfileVersion21"
	case CmICCProfileVersion4:
		return "CmICCProfileVersion4"
	case CmICCReservedFlagsMask:
		return "CmICCReservedFlagsMask"
	case CmInputClass:
		return "CmInputClass"
	case CmInputUse:
		return "CmInputUse"
	case CmInternalCFErr:
		return "CmInternalCFErr"
	case CmInterpolationMask:
		return "CmInterpolationMask"
	case CmLabData:
		return "CmLabData"
	case CmLinkClass:
		return "CmLinkClass"
	case CmLittleEndianPacking:
		return "CmLittleEndianPacking"
	case CmLong10ColorPacking:
		return "CmLong10ColorPacking"
	case CmLuminanceTag:
		return "CmLuminanceTag"
	case CmLuvData:
		return "CmLuvData"
	case CmMCEightSpace:
		return "CmMCEightSpace"
	case CmMCFiveSpace:
		return "CmMCFiveSpace"
	case CmMCH5Data:
		return "CmMCH5Data"
	case CmMCH6Data:
		return "CmMCH6Data"
	case CmMCH7Data:
		return "CmMCH7Data"
	case CmMCH8Data:
		return "CmMCH8Data"
	case CmMCSevenSpace:
		return "CmMCSevenSpace"
	case CmMCSixSpace:
		return "CmMCSixSpace"
	case CmMacintosh:
		return "CmMacintosh"
	case CmMakeAndModelTag:
		return "CmMakeAndModelTag"
	case CmMeasurementTag:
		return "CmMeasurementTag"
	case CmMediaBlackPointTag:
		return "CmMediaBlackPointTag"
	case CmMediaWhitePointTag:
		return "CmMediaWhitePointTag"
	case CmMicrosoft:
		return "CmMicrosoft"
	case CmNamedColor2Tag:
		return "CmNamedColor2Tag"
	case CmNamedColorClass:
		return "CmNamedColorClass"
	case CmNamedColorTag:
		return "CmNamedColorTag"
	case CmNamedData:
		return "CmNamedData"
	case CmNamedIndexedSpace:
		return "CmNamedIndexedSpace"
	case CmNativeDisplayInfoTag:
		return "CmNativeDisplayInfoTag"
	case CmOriginalProfileLocationSize:
		return "CmOriginalProfileLocationSize"
	case CmOutputClass:
		return "CmOutputClass"
	case CmOutputUse:
		return "CmOutputUse"
	case CmPS2CRD0Tag:
		return "CmPS2CRD0Tag"
	case CmPS2CRD1Tag:
		return "CmPS2CRD1Tag"
	case CmPS2CRD2Tag:
		return "CmPS2CRD2Tag"
	case CmPS2CRD3Tag:
		return "CmPS2CRD3Tag"
	case CmPS2CRDVMSizeTag:
		return "CmPS2CRDVMSizeTag"
	case CmPS2CSATag:
		return "CmPS2CSATag"
	case CmPS2RenderingIntentTag:
		return "CmPS2RenderingIntentTag"
	case CmPrefsSynchError:
		return "CmPrefsSynchError"
	case CmPreview0Tag:
		return "CmPreview0Tag"
	case CmPreview1Tag:
		return "CmPreview1Tag"
	case CmPreview2Tag:
		return "CmPreview2Tag"
	case CmProfileDescriptionMLTag:
		return "CmProfileDescriptionMLTag"
	case CmProfileDescriptionTag:
		return "CmProfileDescriptionTag"
	case CmProfileMajorVersionMask:
		return "CmProfileMajorVersionMask"
	case CmProfileSequenceDescTag:
		return "CmProfileSequenceDescTag"
	case CmProofDeviceClass:
		return "CmProofDeviceClass"
	case CmQualityMask:
		return "CmQualityMask"
	case CmRGBData:
		return "CmRGBData"
	case CmRedColorantTag:
		return "CmRedColorantTag"
	case CmRedTRCTag:
		return "CmRedTRCTag"
	case CmReservedSpace2:
		return "CmReservedSpace2"
	case CmReverseChannelPacking:
		return "CmReverseChannelPacking"
	case CmSRGBData:
		return "CmSRGBData"
	case CmScreeningDescTag:
		return "CmScreeningDescTag"
	case CmScreeningTag:
		return "CmScreeningTag"
	case CmSiliconGraphics:
		return "CmSiliconGraphics"
	case CmSolaris:
		return "CmSolaris"
	case CmTaligent:
		return "CmTaligent"
	case CmTechnologyTag:
		return "CmTechnologyTag"
	case CmUcrBgTag:
		return "CmUcrBgTag"
	case CmVideoCardGammaTag:
		return "CmVideoCardGammaTag"
	case CmViewingConditionsDescTag:
		return "CmViewingConditionsDescTag"
	case CmViewingConditionsTag:
		return "CmViewingConditionsTag"
	case CmWord565ColorPacking:
		return "CmWord565ColorPacking"
	case CmWord5ColorPacking:
		return "CmWord5ColorPacking"
	case CmXYZData:
		return "CmXYZData"
	case CmYCbCrData:
		return "CmYCbCrData"
	case CmYxyData:
		return "CmYxyData"
	default:
		return fmt.Sprintf("Cm(%d)", e)
	}
}

type CmBlackPoint uint

const (
	CmBlackPointCompensation CmBlackPoint = 1
)

func (e CmBlackPoint) String() string {
	switch e {
	case CmBlackPointCompensation:
		return "CmBlackPointCompensation"
	default:
		return fmt.Sprintf("CmBlackPoint(%d)", e)
	}
}

type CmCS1 uint

const (
	// CmCS1ChromTag: The tag signature for the profile chromaticities tag whose element data specifies the XYZ chromaticities for the six primary and secondary colors (red, green, blue, cyan, magenta, and yellow).
	CmCS1ChromTag CmCS1 = 'c'<<24 | 'h'<<16 | 'r'<<8 | 'm' // 'chrm'
	// CmCS1CustTag: Private data for a custom CMM.
	CmCS1CustTag CmCS1 = 'c'<<24 | 'u'<<16 | 's'<<8 | 't' // 'cust'
	// CmCS1NameTag: The tag signature for the profile name string.
	CmCS1NameTag CmCS1 = 'n'<<24 | 'a'<<16 | 'm'<<8 | 'e' // 'name'
	// CmCS1TRCTag: The tag signature for profile tonal response curve data for the associated device.
	CmCS1TRCTag CmCS1 = 't'<<24 | 'r'<<16 | 'c'<<8 | ' ' // 'trc '
)

func (e CmCS1) String() string {
	switch e {
	case CmCS1ChromTag:
		return "CmCS1ChromTag"
	case CmCS1CustTag:
		return "CmCS1CustTag"
	case CmCS1NameTag:
		return "CmCS1NameTag"
	case CmCS1TRCTag:
		return "CmCS1TRCTag"
	default:
		return fmt.Sprintf("CmCS1(%d)", e)
	}
}

type CmColorSpace uint

const (
	CmColorSpaceAlphaMask         CmColorSpace = 0x80
	CmColorSpaceEncodingMask      CmColorSpace = 0xf0000
	CmColorSpacePackingMask       CmColorSpace = 0xff00
	CmColorSpacePremulAlphaMask   CmColorSpace = 0x40
	CmColorSpaceReservedMask      CmColorSpace = 0xfff00000
	CmColorSpaceSpaceAndAlphaMask CmColorSpace = 0xff
	CmColorSpaceSpaceMask         CmColorSpace = 0x3f
)

func (e CmColorSpace) String() string {
	switch e {
	case CmColorSpaceAlphaMask:
		return "CmColorSpaceAlphaMask"
	case CmColorSpaceEncodingMask:
		return "CmColorSpaceEncodingMask"
	case CmColorSpacePackingMask:
		return "CmColorSpacePackingMask"
	case CmColorSpacePremulAlphaMask:
		return "CmColorSpacePremulAlphaMask"
	case CmColorSpaceReservedMask:
		return "CmColorSpaceReservedMask"
	case CmColorSpaceSpaceAndAlphaMask:
		return "CmColorSpaceSpaceAndAlphaMask"
	case CmColorSpaceSpaceMask:
		return "CmColorSpaceSpaceMask"
	default:
		return fmt.Sprintf("CmColorSpace(%d)", e)
	}
}

type CmCurrent uint

const (
	CmCurrentDeviceInfoVersion  CmCurrent = 65536
	CmCurrentProfileInfoVersion CmCurrent = 65536
)

func (e CmCurrent) String() string {
	switch e {
	case CmCurrentDeviceInfoVersion:
		return "CmCurrentDeviceInfoVersion"
	default:
		return fmt.Sprintf("CmCurrent(%d)", e)
	}
}

type CmDefault uint

const (
	CmDefaultDeviceID  CmDefault = 0
	CmDefaultProfileID CmDefault = 0
)

func (e CmDefault) String() string {
	switch e {
	case CmDefaultDeviceID:
		return "CmDefaultDeviceID"
	default:
		return fmt.Sprintf("CmDefault(%d)", e)
	}
}

type CmDevice uint

const (
	CmDeviceInfoVersion1        CmDevice = 0x10000
	CmDeviceProfileInfoVersion1 CmDevice = 0x10000
	CmDeviceProfileInfoVersion2 CmDevice = 0x20000
)

func (e CmDevice) String() string {
	switch e {
	case CmDeviceInfoVersion1:
		return "CmDeviceInfoVersion1"
	case CmDeviceProfileInfoVersion2:
		return "CmDeviceProfileInfoVersion2"
	default:
		return fmt.Sprintf("CmDevice(%d)", e)
	}
}

type CmDeviceState uint

const (
	CmDeviceStateAppleRsvdBits  CmDeviceState = 0xff00ffff
	CmDeviceStateBusy           CmDeviceState = 0x2
	CmDeviceStateDefault        CmDeviceState = 0
	CmDeviceStateDeviceRsvdBits CmDeviceState = 0xff0000
	CmDeviceStateForceNotify    CmDeviceState = 0x80000000
	CmDeviceStateOffline        CmDeviceState = 0x1
)

func (e CmDeviceState) String() string {
	switch e {
	case CmDeviceStateAppleRsvdBits:
		return "CmDeviceStateAppleRsvdBits"
	case CmDeviceStateBusy:
		return "CmDeviceStateBusy"
	case CmDeviceStateDefault:
		return "CmDeviceStateDefault"
	case CmDeviceStateDeviceRsvdBits:
		return "CmDeviceStateDeviceRsvdBits"
	case CmDeviceStateForceNotify:
		return "CmDeviceStateForceNotify"
	case CmDeviceStateOffline:
		return "CmDeviceStateOffline"
	default:
		return fmt.Sprintf("CmDeviceState(%d)", e)
	}
}

type CmEmbedded uint

const (
	// CmEmbeddedProfile: 0 is not embedded profile, 1 is embedded profile
	CmEmbeddedProfile CmEmbedded = 0
	// CmEmbeddedUse: 0 is to use anywhere, 1 is to use as embedded profile only
	CmEmbeddedUse CmEmbedded = 1
)

func (e CmEmbedded) String() string {
	switch e {
	case CmEmbeddedProfile:
		return "CmEmbeddedProfile"
	case CmEmbeddedUse:
		return "CmEmbeddedUse"
	default:
		return fmt.Sprintf("CmEmbedded(%d)", e)
	}
}

type CmIlluminant uint

const (
	CmIlluminantA         CmIlluminant = 0x6
	CmIlluminantD50       CmIlluminant = 0x1
	CmIlluminantD55       CmIlluminant = 0x5
	CmIlluminantD65       CmIlluminant = 0x2
	CmIlluminantD93       CmIlluminant = 0x3
	CmIlluminantEquiPower CmIlluminant = 0x7
	CmIlluminantF2        CmIlluminant = 0x4
	CmIlluminantF8        CmIlluminant = 0x8
	CmIlluminantUnknown   CmIlluminant = 0
)

func (e CmIlluminant) String() string {
	switch e {
	case CmIlluminantA:
		return "CmIlluminantA"
	case CmIlluminantD50:
		return "CmIlluminantD50"
	case CmIlluminantD55:
		return "CmIlluminantD55"
	case CmIlluminantD65:
		return "CmIlluminantD65"
	case CmIlluminantD93:
		return "CmIlluminantD93"
	case CmIlluminantEquiPower:
		return "CmIlluminantEquiPower"
	case CmIlluminantF2:
		return "CmIlluminantF2"
	case CmIlluminantF8:
		return "CmIlluminantF8"
	case CmIlluminantUnknown:
		return "CmIlluminantUnknown"
	default:
		return fmt.Sprintf("CmIlluminant(%d)", e)
	}
}

type CmIterate uint

const (
	// CmIterateAllDeviceProfiles: Iterate all profiles, without replacement.
	CmIterateAllDeviceProfiles CmIterate = 0x4
	// CmIterateCurrentDeviceProfiles: # Discussion
	CmIterateCurrentDeviceProfiles CmIterate = 0x3
	// CmIterateCustomDeviceProfiles: # Discussion
	CmIterateCustomDeviceProfiles CmIterate = 0x2
	CmIterateDeviceProfilesMask   CmIterate = 0xf
	// CmIterateFactoryDeviceProfiles: Iterate profiles registered through the routine [CMSetDeviceFactoryProfiles].
	CmIterateFactoryDeviceProfiles CmIterate = 0x1
)

func (e CmIterate) String() string {
	switch e {
	case CmIterateAllDeviceProfiles:
		return "CmIterateAllDeviceProfiles"
	case CmIterateCurrentDeviceProfiles:
		return "CmIterateCurrentDeviceProfiles"
	case CmIterateCustomDeviceProfiles:
		return "CmIterateCustomDeviceProfiles"
	case CmIterateDeviceProfilesMask:
		return "CmIterateDeviceProfilesMask"
	case CmIterateFactoryDeviceProfiles:
		return "CmIterateFactoryDeviceProfiles"
	default:
		return fmt.Sprintf("CmIterate(%d)", e)
	}
}

type CmMagic uint

const (
	CmMagicNumber CmMagic = 'a'<<24 | 'c'<<16 | 's'<<8 | 'p' // 'acsp'
)

func (e CmMagic) String() string {
	switch e {
	case CmMagicNumber:
		return "CmMagicNumber"
	default:
		return fmt.Sprintf("CmMagic(%d)", e)
	}
}

type CmNumHeader uint

const (
	CmNumHeaderElements CmNumHeader = 10
)

func (e CmNumHeader) String() string {
	switch e {
	case CmNumHeaderElements:
		return "CmNumHeaderElements"
	default:
		return fmt.Sprintf("CmNumHeader(%d)", e)
	}
}

type CmP uint

const (
	// CmPS7bit: The data is 7-bit safe—therefore the data could be in 7-bit ASCII encoding or in ASCII base-85 encoding.
	CmPS7bit CmP = 1
	// CmPS8bit: The data is 8-bit safe—therefore the data could be in 7-bit or 8-bit ASCII encoding.
	CmPS8bit CmP = 2
)

func (e CmP) String() string {
	switch e {
	case CmPS7bit:
		return "CmPS7bit"
	case CmPS8bit:
		return "CmPS8bit"
	default:
		return fmt.Sprintf("CmP(%d)", e)
	}
}

type CmParametric uint

const (
	// CmParametricType0: Y = X^gamma
	CmParametricType0 CmParametric = 0
	// CmParametricType1: Y = (aX+b)^gamma     [X>=-b/a],  Y = 0    [X<-b/a]
	CmParametricType1 CmParametric = 1
	// CmParametricType2: Y = (aX+b)^gamma + c [X>=-b/a],  Y = c    [X<-b/a]
	CmParametricType2 CmParametric = 2
	// CmParametricType3: Y = (aX+b)^gamma     [X>=d],     Y = cX   [X<d]
	CmParametricType3 CmParametric = 3
	// CmParametricType4: Y = (aX+b)^gamma + e [X>=d],     Y = cX+f [X<d]
	CmParametricType4 CmParametric = 4
)

func (e CmParametric) String() string {
	switch e {
	case CmParametricType0:
		return "CmParametricType0"
	case CmParametricType1:
		return "CmParametricType1"
	case CmParametricType2:
		return "CmParametricType2"
	case CmParametricType3:
		return "CmParametricType3"
	case CmParametricType4:
		return "CmParametricType4"
	default:
		return fmt.Sprintf("CmParametric(%d)", e)
	}
}

type CmProfileIterateData uint

const (
	CmProfileIterateDataVersion1 CmProfileIterateData = 0x10000
	// CmProfileIterateDataVersion2: Added makeAndModel
	CmProfileIterateDataVersion2 CmProfileIterateData = 0x20000
	// CmProfileIterateDataVersion3: Added MD5 digest
	CmProfileIterateDataVersion3 CmProfileIterateData = 0x30000
	CmProfileIterateDataVersion4 CmProfileIterateData = 0x40000
)

func (e CmProfileIterateData) String() string {
	switch e {
	case CmProfileIterateDataVersion1:
		return "CmProfileIterateDataVersion1"
	case CmProfileIterateDataVersion2:
		return "CmProfileIterateDataVersion2"
	case CmProfileIterateDataVersion3:
		return "CmProfileIterateDataVersion3"
	case CmProfileIterateDataVersion4:
		return "CmProfileIterateDataVersion4"
	default:
		return fmt.Sprintf("CmProfileIterateData(%d)", e)
	}
}

type CmSRGB16Channel uint

const (
	// CmSRGB16ChannelEncoding: Used for sRGB64 encoding ( ±3.12 format)
	CmSRGB16ChannelEncoding CmSRGB16Channel = 0x10000
)

func (e CmSRGB16Channel) String() string {
	switch e {
	case CmSRGB16ChannelEncoding:
		return "CmSRGB16ChannelEncoding"
	default:
		return fmt.Sprintf("CmSRGB16Channel(%d)", e)
	}
}

type CmSig uint

const (
	CmSigCrdInfoType               CmSig = 'c'<<24 | 'r'<<16 | 'd'<<8 | 'i' // 'crdi'
	CmSigCurveType                 CmSig = 'c'<<24 | 'u'<<16 | 'r'<<8 | 'v' // 'curv'
	CmSigDataType                  CmSig = 'd'<<24 | 'a'<<16 | 't'<<8 | 'a' // 'data'
	CmSigDateTimeType              CmSig = 'd'<<24 | 't'<<16 | 'i'<<8 | 'm' // 'dtim'
	CmSigLut16Type                 CmSig = 'm'<<24 | 'f'<<16 | 't'<<8 | '2' // 'mft2'
	CmSigLut8Type                  CmSig = 'm'<<24 | 'f'<<16 | 't'<<8 | '1' // 'mft1'
	CmSigMakeAndModelType          CmSig = 'm'<<24 | 'm'<<16 | 'o'<<8 | 'd' // 'mmod'
	CmSigMeasurementType           CmSig = 'm'<<24 | 'e'<<16 | 'a'<<8 | 's' // 'meas'
	CmSigMultiFunctA2BType         CmSig = 'm'<<24 | 'A'<<16 | 'B'<<8 | ' ' // 'mAB '
	CmSigMultiFunctB2AType         CmSig = 'm'<<24 | 'B'<<16 | 'A'<<8 | ' ' // 'mBA '
	CmSigMultiLocalizedUniCodeType CmSig = 'm'<<24 | 'l'<<16 | 'u'<<8 | 'c' // 'mluc'
	CmSigNamedColor2Type           CmSig = 'n'<<24 | 'c'<<16 | 'l'<<8 | '2' // 'ncl2'
	CmSigNamedColorType            CmSig = 'n'<<24 | 'c'<<16 | 'o'<<8 | 'l' // 'ncol'
	CmSigNativeDisplayInfoType     CmSig = 'n'<<24 | 'd'<<16 | 'i'<<8 | 'n' // 'ndin'
	CmSigPS2CRDVMSizeType          CmSig = 'p'<<24 | 's'<<16 | 'v'<<8 | 'm' // 'psvm'
	CmSigParametricCurveType       CmSig = 'p'<<24 | 'a'<<16 | 'r'<<8 | 'a' // 'para'
	CmSigProfileDescriptionType    CmSig = 'd'<<24 | 'e'<<16 | 's'<<8 | 'c' // 'desc'
	CmSigProfileSequenceDescType   CmSig = 'p'<<24 | 's'<<16 | 'e'<<8 | 'q' // 'pseq'
	CmSigS15Fixed16Type            CmSig = 's'<<24 | 'f'<<16 | '3'<<8 | '2' // 'sf32'
	CmSigScreeningType             CmSig = 's'<<24 | 'c'<<16 | 'r'<<8 | 'n' // 'scrn'
	CmSigSignatureType             CmSig = 's'<<24 | 'i'<<16 | 'g'<<8 | ' ' // 'sig '
	CmSigTextType                  CmSig = 't'<<24 | 'e'<<16 | 'x'<<8 | 't' // 'text'
	CmSigU16Fixed16Type            CmSig = 'u'<<24 | 'f'<<16 | '3'<<8 | '2' // 'uf32'
	CmSigU1Fixed15Type             CmSig = 'u'<<24 | 'f'<<16 | '1'<<8 | '6' // 'uf16'
	CmSigUInt16Type                CmSig = 'u'<<24 | 'i'<<16 | '1'<<8 | '6' // 'ui16'
	CmSigUInt32Type                CmSig = 'u'<<24 | 'i'<<16 | '3'<<8 | '2' // 'ui32'
	CmSigUInt64Type                CmSig = 'u'<<24 | 'i'<<16 | '6'<<8 | '4' // 'ui64'
	CmSigUInt8Type                 CmSig = 'u'<<24 | 'i'<<16 | '0'<<8 | '8' // 'ui08'
	CmSigUcrBgType                 CmSig = 'b'<<24 | 'f'<<16 | 'd'<<8 | ' ' // 'bfd '
	CmSigUnicodeTextType           CmSig = 'u'<<24 | 't'<<16 | 'x'<<8 | 't' // 'utxt'
	// CmSigVideoCardGammaType: # Discussion
	CmSigVideoCardGammaType    CmSig = 'v'<<24 | 'c'<<16 | 'g'<<8 | 't' // 'vcgt'
	CmSigViewingConditionsType CmSig = 'v'<<24 | 'i'<<16 | 'e'<<8 | 'w' // 'view'
	CmSigXYZType               CmSig = 'X'<<24 | 'Y'<<16 | 'Z'<<8 | ' ' // 'XYZ '
)

func (e CmSig) String() string {
	switch e {
	case CmSigCrdInfoType:
		return "CmSigCrdInfoType"
	case CmSigCurveType:
		return "CmSigCurveType"
	case CmSigDataType:
		return "CmSigDataType"
	case CmSigDateTimeType:
		return "CmSigDateTimeType"
	case CmSigLut16Type:
		return "CmSigLut16Type"
	case CmSigLut8Type:
		return "CmSigLut8Type"
	case CmSigMakeAndModelType:
		return "CmSigMakeAndModelType"
	case CmSigMeasurementType:
		return "CmSigMeasurementType"
	case CmSigMultiFunctA2BType:
		return "CmSigMultiFunctA2BType"
	case CmSigMultiFunctB2AType:
		return "CmSigMultiFunctB2AType"
	case CmSigMultiLocalizedUniCodeType:
		return "CmSigMultiLocalizedUniCodeType"
	case CmSigNamedColor2Type:
		return "CmSigNamedColor2Type"
	case CmSigNamedColorType:
		return "CmSigNamedColorType"
	case CmSigNativeDisplayInfoType:
		return "CmSigNativeDisplayInfoType"
	case CmSigPS2CRDVMSizeType:
		return "CmSigPS2CRDVMSizeType"
	case CmSigParametricCurveType:
		return "CmSigParametricCurveType"
	case CmSigProfileDescriptionType:
		return "CmSigProfileDescriptionType"
	case CmSigProfileSequenceDescType:
		return "CmSigProfileSequenceDescType"
	case CmSigS15Fixed16Type:
		return "CmSigS15Fixed16Type"
	case CmSigScreeningType:
		return "CmSigScreeningType"
	case CmSigSignatureType:
		return "CmSigSignatureType"
	case CmSigTextType:
		return "CmSigTextType"
	case CmSigU16Fixed16Type:
		return "CmSigU16Fixed16Type"
	case CmSigU1Fixed15Type:
		return "CmSigU1Fixed15Type"
	case CmSigUInt16Type:
		return "CmSigUInt16Type"
	case CmSigUInt32Type:
		return "CmSigUInt32Type"
	case CmSigUInt64Type:
		return "CmSigUInt64Type"
	case CmSigUInt8Type:
		return "CmSigUInt8Type"
	case CmSigUcrBgType:
		return "CmSigUcrBgType"
	case CmSigUnicodeTextType:
		return "CmSigUnicodeTextType"
	case CmSigVideoCardGammaType:
		return "CmSigVideoCardGammaType"
	case CmSigViewingConditionsType:
		return "CmSigViewingConditionsType"
	case CmSigXYZType:
		return "CmSigXYZType"
	default:
		return fmt.Sprintf("CmSig(%d)", e)
	}
}

type CmSpotFunction uint

const (
	CmSpotFunctionCross   CmSpotFunction = 7
	CmSpotFunctionDefault CmSpotFunction = 1
	CmSpotFunctionDiamond CmSpotFunction = 3
	CmSpotFunctionEllipse CmSpotFunction = 4
	CmSpotFunctionLine    CmSpotFunction = 5
	CmSpotFunctionRound   CmSpotFunction = 2
	CmSpotFunctionSquare  CmSpotFunction = 6
	CmSpotFunctionUnknown CmSpotFunction = 0
)

func (e CmSpotFunction) String() string {
	switch e {
	case CmSpotFunctionCross:
		return "CmSpotFunctionCross"
	case CmSpotFunctionDefault:
		return "CmSpotFunctionDefault"
	case CmSpotFunctionDiamond:
		return "CmSpotFunctionDiamond"
	case CmSpotFunctionEllipse:
		return "CmSpotFunctionEllipse"
	case CmSpotFunctionLine:
		return "CmSpotFunctionLine"
	case CmSpotFunctionRound:
		return "CmSpotFunctionRound"
	case CmSpotFunctionSquare:
		return "CmSpotFunctionSquare"
	case CmSpotFunctionUnknown:
		return "CmSpotFunctionUnknown"
	default:
		return fmt.Sprintf("CmSpotFunction(%d)", e)
	}
}

type CmTechnology uint

const (
	CmTechnologyAMDisplay                  CmTechnology = 'A'<<24 | 'M'<<16 | 'D'<<8 | ' ' // 'AMD '
	CmTechnologyCRTDisplay                 CmTechnology = 'C'<<24 | 'R'<<16 | 'T'<<8 | ' ' // 'CRT '
	CmTechnologyDigitalCamera              CmTechnology = 'd'<<24 | 'c'<<16 | 'a'<<8 | 'm' // 'dcam'
	CmTechnologyDyeSublimationPrinter      CmTechnology = 'd'<<24 | 's'<<16 | 'u'<<8 | 'b' // 'dsub'
	CmTechnologyElectrophotographicPrinter CmTechnology = 'e'<<24 | 'p'<<16 | 'h'<<8 | 'o' // 'epho'
	CmTechnologyElectrostaticPrinter       CmTechnology = 'e'<<24 | 's'<<16 | 't'<<8 | 'a' // 'esta'
	CmTechnologyFilmScanner                CmTechnology = 'f'<<24 | 's'<<16 | 'c'<<8 | 'n' // 'fscn'
	CmTechnologyFilmWriter                 CmTechnology = 'f'<<24 | 'p'<<16 | 'r'<<8 | 'n' // 'fprn'
	CmTechnologyFlexography                CmTechnology = 'f'<<24 | 'l'<<16 | 'e'<<8 | 'x' // 'flex'
	CmTechnologyGravure                    CmTechnology = 'g'<<24 | 'r'<<16 | 'a'<<8 | 'v' // 'grav'
	CmTechnologyInkJetPrinter              CmTechnology = 'i'<<24 | 'j'<<16 | 'e'<<8 | 't' // 'ijet'
	CmTechnologyOffsetLithography          CmTechnology = 'o'<<24 | 'f'<<16 | 'f'<<8 | 's' // 'offs'
	CmTechnologyPMDisplay                  CmTechnology = 'P'<<24 | 'M'<<16 | 'D'<<8 | ' ' // 'PMD '
	CmTechnologyPhotoCD                    CmTechnology = 'K'<<24 | 'P'<<16 | 'C'<<8 | 'D' // 'KPCD'
	CmTechnologyPhotoImageSetter           CmTechnology = 'i'<<24 | 'm'<<16 | 'g'<<8 | 's' // 'imgs'
	CmTechnologyPhotographicPaperPrinter   CmTechnology = 'r'<<24 | 'p'<<16 | 'h'<<8 | 'o' // 'rpho'
	CmTechnologyProjectionTelevision       CmTechnology = 'p'<<24 | 'j'<<16 | 't'<<8 | 'v' // 'pjtv'
	CmTechnologyReflectiveScanner          CmTechnology = 'r'<<24 | 's'<<16 | 'c'<<8 | 'n' // 'rscn'
	CmTechnologySilkscreen                 CmTechnology = 's'<<24 | 'i'<<16 | 'l'<<8 | 'k' // 'silk'
	CmTechnologyThermalWaxPrinter          CmTechnology = 't'<<24 | 'w'<<16 | 'a'<<8 | 'x' // 'twax'
	CmTechnologyVideoCamera                CmTechnology = 'v'<<24 | 'i'<<16 | 'd'<<8 | 'c' // 'vidc'
	CmTechnologyVideoMonitor               CmTechnology = 'v'<<24 | 'i'<<16 | 'd'<<8 | 'm' // 'vidm'
)

func (e CmTechnology) String() string {
	switch e {
	case CmTechnologyAMDisplay:
		return "CmTechnologyAMDisplay"
	case CmTechnologyCRTDisplay:
		return "CmTechnologyCRTDisplay"
	case CmTechnologyDigitalCamera:
		return "CmTechnologyDigitalCamera"
	case CmTechnologyDyeSublimationPrinter:
		return "CmTechnologyDyeSublimationPrinter"
	case CmTechnologyElectrophotographicPrinter:
		return "CmTechnologyElectrophotographicPrinter"
	case CmTechnologyElectrostaticPrinter:
		return "CmTechnologyElectrostaticPrinter"
	case CmTechnologyFilmScanner:
		return "CmTechnologyFilmScanner"
	case CmTechnologyFilmWriter:
		return "CmTechnologyFilmWriter"
	case CmTechnologyFlexography:
		return "CmTechnologyFlexography"
	case CmTechnologyGravure:
		return "CmTechnologyGravure"
	case CmTechnologyInkJetPrinter:
		return "CmTechnologyInkJetPrinter"
	case CmTechnologyOffsetLithography:
		return "CmTechnologyOffsetLithography"
	case CmTechnologyPMDisplay:
		return "CmTechnologyPMDisplay"
	case CmTechnologyPhotoCD:
		return "CmTechnologyPhotoCD"
	case CmTechnologyPhotoImageSetter:
		return "CmTechnologyPhotoImageSetter"
	case CmTechnologyPhotographicPaperPrinter:
		return "CmTechnologyPhotographicPaperPrinter"
	case CmTechnologyProjectionTelevision:
		return "CmTechnologyProjectionTelevision"
	case CmTechnologyReflectiveScanner:
		return "CmTechnologyReflectiveScanner"
	case CmTechnologySilkscreen:
		return "CmTechnologySilkscreen"
	case CmTechnologyThermalWaxPrinter:
		return "CmTechnologyThermalWaxPrinter"
	case CmTechnologyVideoCamera:
		return "CmTechnologyVideoCamera"
	case CmTechnologyVideoMonitor:
		return "CmTechnologyVideoMonitor"
	default:
		return fmt.Sprintf("CmTechnology(%d)", e)
	}
}

type CmVideoCardGamma uint

const (
	// CmVideoCardGammaFormulaType: The video card gamma tag data is stored as a formula.
	CmVideoCardGammaFormulaType CmVideoCardGamma = 1
	// CmVideoCardGammaTableType: The video card gamma data is stored in a table format.
	CmVideoCardGammaTableType CmVideoCardGamma = 0
)

func (e CmVideoCardGamma) String() string {
	switch e {
	case CmVideoCardGammaFormulaType:
		return "CmVideoCardGammaFormulaType"
	case CmVideoCardGammaTableType:
		return "CmVideoCardGammaTableType"
	default:
		return fmt.Sprintf("CmVideoCardGamma(%d)", e)
	}
}

type CsMax uint

const (
	CS_MAX_PATH CsMax = 1024
)

func (e CsMax) String() string {
	switch e {
	case CS_MAX_PATH:
		return "CS_MAX_PATH"
	default:
		return fmt.Sprintf("CsMax(%d)", e)
	}
}

type Extended uint

const (
	ExtendedBlock    Extended = 0x4c43
	ExtendedBlockLen Extended = 0
)

func (e Extended) String() string {
	switch e {
	case ExtendedBlock:
		return "ExtendedBlock"
	case ExtendedBlockLen:
		return "ExtendedBlockLen"
	default:
		return fmt.Sprintf("Extended(%d)", e)
	}
}

type Ic int

const (
	IcConfigInappropriateErr Ic = -675
	IcConfigNotFoundErr      Ic = -674
	IcInternalErr            Ic = -669
	IcNoMoreWritersErr       Ic = -671
	IcNoPerm                 Ic = 0
	IcNoURLErr               Ic = -673
	IcNothingToOverrideErr   Ic = -672
	IcPermErr                Ic = -667
	IcPrefDataErr            Ic = -668
	IcPrefNotFoundErr        Ic = -666
	IcProfileNotFoundErr     Ic = -676
	IcReadOnlyPerm           Ic = 1
	IcReadWritePerm          Ic = 2
	IcTooManyProfilesErr     Ic = -677
	IcTruncatedErr           Ic = -670
)

func (e Ic) String() string {
	switch e {
	case IcConfigInappropriateErr:
		return "IcConfigInappropriateErr"
	case IcConfigNotFoundErr:
		return "IcConfigNotFoundErr"
	case IcInternalErr:
		return "IcInternalErr"
	case IcNoMoreWritersErr:
		return "IcNoMoreWritersErr"
	case IcNoPerm:
		return "IcNoPerm"
	case IcNoURLErr:
		return "IcNoURLErr"
	case IcNothingToOverrideErr:
		return "IcNothingToOverrideErr"
	case IcPermErr:
		return "IcPermErr"
	case IcPrefDataErr:
		return "IcPrefDataErr"
	case IcPrefNotFoundErr:
		return "IcPrefNotFoundErr"
	case IcProfileNotFoundErr:
		return "IcProfileNotFoundErr"
	case IcReadOnlyPerm:
		return "IcReadOnlyPerm"
	case IcReadWritePerm:
		return "IcReadWritePerm"
	case IcTooManyProfilesErr:
		return "IcTooManyProfilesErr"
	case IcTruncatedErr:
		return "IcTruncatedErr"
	default:
		return fmt.Sprintf("Ic(%d)", e)
	}
}

type InitDev uint

const (
	ActivDev     InitDev = 5
	ClearDev     InitDev = 13
	CloseDev     InitDev = 2
	CopyDev      InitDev = 11
	CursorDev    InitDev = 14
	CutDev       InitDev = 10
	DeactivDev   InitDev = 6
	HitDev       InitDev = 1
	InitDevValue InitDev = 0
	KeyEvtDev    InitDev = 7
	MacDev       InitDev = 8
	NulDev       InitDev = 3
	PasteDev     InitDev = 12
	UndoDev      InitDev = 9
	UpdateDev    InitDev = 4
)

func (e InitDev) String() string {
	switch e {
	case ActivDev:
		return "ActivDev"
	case ClearDev:
		return "ClearDev"
	case CloseDev:
		return "CloseDev"
	case CopyDev:
		return "CopyDev"
	case CursorDev:
		return "CursorDev"
	case CutDev:
		return "CutDev"
	case DeactivDev:
		return "DeactivDev"
	case HitDev:
		return "HitDev"
	case InitDevValue:
		return "InitDevValue"
	case KeyEvtDev:
		return "KeyEvtDev"
	case MacDev:
		return "MacDev"
	case NulDev:
		return "NulDev"
	case PasteDev:
		return "PasteDev"
	case UndoDev:
		return "UndoDev"
	case UpdateDev:
		return "UpdateDev"
	default:
		return fmt.Sprintf("InitDev(%d)", e)
	}
}

type K1MonochromePixelFormat uint

const (
	K16BE555PixelFormat          K1MonochromePixelFormat = 0x10
	K16BE565PixelFormat          K1MonochromePixelFormat = 'B'<<24 | '5'<<16 | '6'<<8 | '5' // 'B565'
	K16LE5551PixelFormat         K1MonochromePixelFormat = '5'<<24 | '5'<<16 | '5'<<8 | '1' // '5551'
	K16LE555PixelFormat          K1MonochromePixelFormat = 'L'<<24 | '5'<<16 | '5'<<8 | '5' // 'L555'
	K16LE565PixelFormat          K1MonochromePixelFormat = 'L'<<24 | '5'<<16 | '6'<<8 | '5' // 'L565'
	K1IndexedGrayPixelFormat     K1MonochromePixelFormat = 0x21
	K1MonochromePixelFormatValue K1MonochromePixelFormat = 0x1
	K24BGRPixelFormat            K1MonochromePixelFormat = '2'<<24 | '4'<<16 | 'B'<<8 | 'G' // '24BG'
	K24RGBPixelFormat            K1MonochromePixelFormat = 0x18
	K2IndexedGrayPixelFormat     K1MonochromePixelFormat = 0x22
	K2IndexedPixelFormat         K1MonochromePixelFormat = 0x2
	K2vuyPixelFormat             K1MonochromePixelFormat = '2'<<24 | 'v'<<16 | 'u'<<8 | 'y' // '2vuy'
	K32ABGRPixelFormat           K1MonochromePixelFormat = 'A'<<24 | 'B'<<16 | 'G'<<8 | 'R' // 'ABGR'
	K32ARGBPixelFormat           K1MonochromePixelFormat = 0x20
	K32BGRAPixelFormat           K1MonochromePixelFormat = 'B'<<24 | 'G'<<16 | 'R'<<8 | 'A' // 'BGRA'
	K32RGBAPixelFormat           K1MonochromePixelFormat = 'R'<<24 | 'G'<<16 | 'B'<<8 | 'A' // 'RGBA'
	K4IndexedGrayPixelFormat     K1MonochromePixelFormat = 0x24
	K4IndexedPixelFormat         K1MonochromePixelFormat = 0x4
	K8IndexedGrayPixelFormat     K1MonochromePixelFormat = 0x28
	K8IndexedPixelFormat         K1MonochromePixelFormat = 0x8
	KUYVY422PixelFormat          K1MonochromePixelFormat = 'U'<<24 | 'Y'<<16 | 'V'<<8 | 'Y' // 'UYVY'
	KYUV211PixelFormat           K1MonochromePixelFormat = 'Y'<<24 | '2'<<16 | '1'<<8 | '1' // 'Y211'
	KYUV411PixelFormat           K1MonochromePixelFormat = 'Y'<<24 | '4'<<16 | '1'<<8 | '1' // 'Y411'
	KYUVSPixelFormat             K1MonochromePixelFormat = 'y'<<24 | 'u'<<16 | 'v'<<8 | 's' // 'yuvs'
	KYUVUPixelFormat             K1MonochromePixelFormat = 'y'<<24 | 'u'<<16 | 'v'<<8 | 'u' // 'yuvu'
	KYVU9PixelFormat             K1MonochromePixelFormat = 'Y'<<24 | 'V'<<16 | 'U'<<8 | '9' // 'YVU9'
	KYVYU422PixelFormat          K1MonochromePixelFormat = 'Y'<<24 | 'V'<<16 | 'Y'<<8 | 'U' // 'YVYU'
)

func (e K1MonochromePixelFormat) String() string {
	switch e {
	case K16BE555PixelFormat:
		return "K16BE555PixelFormat"
	case K16BE565PixelFormat:
		return "K16BE565PixelFormat"
	case K16LE5551PixelFormat:
		return "K16LE5551PixelFormat"
	case K16LE555PixelFormat:
		return "K16LE555PixelFormat"
	case K16LE565PixelFormat:
		return "K16LE565PixelFormat"
	case K1IndexedGrayPixelFormat:
		return "K1IndexedGrayPixelFormat"
	case K1MonochromePixelFormatValue:
		return "K1MonochromePixelFormatValue"
	case K24BGRPixelFormat:
		return "K24BGRPixelFormat"
	case K24RGBPixelFormat:
		return "K24RGBPixelFormat"
	case K2IndexedGrayPixelFormat:
		return "K2IndexedGrayPixelFormat"
	case K2IndexedPixelFormat:
		return "K2IndexedPixelFormat"
	case K2vuyPixelFormat:
		return "K2vuyPixelFormat"
	case K32ABGRPixelFormat:
		return "K32ABGRPixelFormat"
	case K32ARGBPixelFormat:
		return "K32ARGBPixelFormat"
	case K32BGRAPixelFormat:
		return "K32BGRAPixelFormat"
	case K32RGBAPixelFormat:
		return "K32RGBAPixelFormat"
	case K4IndexedGrayPixelFormat:
		return "K4IndexedGrayPixelFormat"
	case K4IndexedPixelFormat:
		return "K4IndexedPixelFormat"
	case K8IndexedGrayPixelFormat:
		return "K8IndexedGrayPixelFormat"
	case K8IndexedPixelFormat:
		return "K8IndexedPixelFormat"
	case KUYVY422PixelFormat:
		return "KUYVY422PixelFormat"
	case KYUV211PixelFormat:
		return "KYUV211PixelFormat"
	case KYUV411PixelFormat:
		return "KYUV411PixelFormat"
	case KYUVSPixelFormat:
		return "KYUVSPixelFormat"
	case KYUVUPixelFormat:
		return "KYUVUPixelFormat"
	case KYVU9PixelFormat:
		return "KYVU9PixelFormat"
	case KYVYU422PixelFormat:
		return "KYVYU422PixelFormat"
	default:
		return fmt.Sprintf("K1MonochromePixelFormat(%d)", e)
	}
}

type KATS int

const (
	KATSBoldQDStretch               KATS = 65536
	KATSCubicCurveType              KATS = 0x1
	KATSFontContainerRefUnspecified KATS = 0
	KATSFontFamilyRefUnspecified    KATS = 0
	KATSFontRefUnspecified          KATS = 0
	KATSGenerationUnspecified       KATS = 0
	KATSInvalidFontAccess           KATS = -982
	KATSInvalidFontContainerAccess  KATS = -985
	KATSInvalidFontFamilyAccess     KATS = -981
	KATSInvalidFontTableAccess      KATS = -984
	KATSInvalidGlyphAccess          KATS = -986
	KATSItalicQDSkew                KATS = 65536
	KATSIterationCompleted          KATS = -980
	KATSIterationScopeModified      KATS = -983
	KATSNoTracking                  KATS = 0x80000000
	KATSOtherCurveType              KATS = 0x3
	KATSQuadCurveType               KATS = 0x2
	KATSRadiansFactor               KATS = 1144
	KATSUseGlyphAdvance             KATS = 0x7fffffff
	KATSUseLineHeight               KATS = 0x7fffffff
)

func (e KATS) String() string {
	switch e {
	case KATSBoldQDStretch:
		return "KATSBoldQDStretch"
	case KATSCubicCurveType:
		return "KATSCubicCurveType"
	case KATSFontContainerRefUnspecified:
		return "KATSFontContainerRefUnspecified"
	case KATSInvalidFontAccess:
		return "KATSInvalidFontAccess"
	case KATSInvalidFontContainerAccess:
		return "KATSInvalidFontContainerAccess"
	case KATSInvalidFontFamilyAccess:
		return "KATSInvalidFontFamilyAccess"
	case KATSInvalidFontTableAccess:
		return "KATSInvalidFontTableAccess"
	case KATSInvalidGlyphAccess:
		return "KATSInvalidGlyphAccess"
	case KATSIterationCompleted:
		return "KATSIterationCompleted"
	case KATSIterationScopeModified:
		return "KATSIterationScopeModified"
	case KATSNoTracking:
		return "KATSNoTracking"
	case KATSOtherCurveType:
		return "KATSOtherCurveType"
	case KATSQuadCurveType:
		return "KATSQuadCurveType"
	case KATSRadiansFactor:
		return "KATSRadiansFactor"
	case KATSUseGlyphAdvance:
		return "KATSUseGlyphAdvance"
	default:
		return fmt.Sprintf("KATS(%d)", e)
	}
}

type KATSDeleted uint

const (
	KATSDeletedGlyphcode KATSDeleted = 0xffff
)

func (e KATSDeleted) String() string {
	switch e {
	case KATSDeletedGlyphcode:
		return "KATSDeletedGlyphcode"
	default:
		return fmt.Sprintf("KATSDeleted(%d)", e)
	}
}

type KATSFlatDataUstl uint

const (
	KATSFlatDataUstlCurrentVersion KATSFlatDataUstl = 2
	KATSFlatDataUstlVersion0       KATSFlatDataUstl = 0
	KATSFlatDataUstlVersion1       KATSFlatDataUstl = 1
	KATSFlatDataUstlVersion2       KATSFlatDataUstl = 2
)

func (e KATSFlatDataUstl) String() string {
	switch e {
	case KATSFlatDataUstlCurrentVersion:
		return "KATSFlatDataUstlCurrentVersion"
	case KATSFlatDataUstlVersion0:
		return "KATSFlatDataUstlVersion0"
	case KATSFlatDataUstlVersion1:
		return "KATSFlatDataUstlVersion1"
	default:
		return fmt.Sprintf("KATSFlatDataUstl(%d)", e)
	}
}

type KATSFlattenedFontSpecifierRawName uint

const (
	KATSFlattenedFontSpecifierRawNameData KATSFlattenedFontSpecifierRawName = 'n'<<24 | 'a'<<16 | 'm'<<8 | 'd' // 'namd'
)

func (e KATSFlattenedFontSpecifierRawName) String() string {
	switch e {
	case KATSFlattenedFontSpecifierRawNameData:
		return "KATSFlattenedFontSpecifierRawNameData"
	default:
		return fmt.Sprintf("KATSFlattenedFontSpecifierRawName(%d)", e)
	}
}

type KATSFontAutoActivation uint

const (
	KATSFontAutoActivationAsk      KATSFontAutoActivation = 4
	KATSFontAutoActivationDefault  KATSFontAutoActivation = 0
	KATSFontAutoActivationDisabled KATSFontAutoActivation = 1
	KATSFontAutoActivationEnabled  KATSFontAutoActivation = 2
)

func (e KATSFontAutoActivation) String() string {
	switch e {
	case KATSFontAutoActivationAsk:
		return "KATSFontAutoActivationAsk"
	case KATSFontAutoActivationDefault:
		return "KATSFontAutoActivationDefault"
	case KATSFontAutoActivationDisabled:
		return "KATSFontAutoActivationDisabled"
	case KATSFontAutoActivationEnabled:
		return "KATSFontAutoActivationEnabled"
	default:
		return fmt.Sprintf("KATSFontAutoActivation(%d)", e)
	}
}

type KATSFontContext uint

const (
	KATSFontContextGlobal      KATSFontContext = 1
	KATSFontContextLocal       KATSFontContext = 2
	KATSFontContextUnspecified KATSFontContext = 0
)

func (e KATSFontContext) String() string {
	switch e {
	case KATSFontContextGlobal:
		return "KATSFontContextGlobal"
	case KATSFontContextLocal:
		return "KATSFontContextLocal"
	case KATSFontContextUnspecified:
		return "KATSFontContextUnspecified"
	default:
		return fmt.Sprintf("KATSFontContext(%d)", e)
	}
}

type KATSFontFilterCurrent uint

const (
	KATSFontFilterCurrentVersion KATSFontFilterCurrent = 0
)

func (e KATSFontFilterCurrent) String() string {
	switch e {
	case KATSFontFilterCurrentVersion:
		return "KATSFontFilterCurrentVersion"
	default:
		return fmt.Sprintf("KATSFontFilterCurrent(%d)", e)
	}
}

type KATSFontFormat uint

const (
	KATSFontFormatUnspecified KATSFontFormat = 0
)

func (e KATSFontFormat) String() string {
	switch e {
	case KATSFontFormatUnspecified:
		return "KATSFontFormatUnspecified"
	default:
		return fmt.Sprintf("KATSFontFormat(%d)", e)
	}
}

type KATSGlyphInfo uint

const (
	KATSGlyphInfoAppleReserved   KATSGlyphInfo = 0x1ffbffe8
	KATSGlyphInfoByteSizeMask    KATSGlyphInfo = 0x7
	KATSGlyphInfoHasImposedWidth KATSGlyphInfo = 0x10
	KATSGlyphInfoIsAttachment    KATSGlyphInfo = 0x80000000
	KATSGlyphInfoIsLTHanger      KATSGlyphInfo = 0x40000000
	KATSGlyphInfoIsRBHanger      KATSGlyphInfo = 0x20000000
	KATSGlyphInfoIsWhiteSpace    KATSGlyphInfo = 0x40000
	KATSGlyphInfoTerminatorGlyph KATSGlyphInfo = 0x80000
)

func (e KATSGlyphInfo) String() string {
	switch e {
	case KATSGlyphInfoAppleReserved:
		return "KATSGlyphInfoAppleReserved"
	case KATSGlyphInfoByteSizeMask:
		return "KATSGlyphInfoByteSizeMask"
	case KATSGlyphInfoHasImposedWidth:
		return "KATSGlyphInfoHasImposedWidth"
	case KATSGlyphInfoIsAttachment:
		return "KATSGlyphInfoIsAttachment"
	case KATSGlyphInfoIsLTHanger:
		return "KATSGlyphInfoIsLTHanger"
	case KATSGlyphInfoIsRBHanger:
		return "KATSGlyphInfoIsRBHanger"
	case KATSGlyphInfoIsWhiteSpace:
		return "KATSGlyphInfoIsWhiteSpace"
	case KATSGlyphInfoTerminatorGlyph:
		return "KATSGlyphInfoTerminatorGlyph"
	default:
		return fmt.Sprintf("KATSGlyphInfo(%d)", e)
	}
}

type KATSLine uint

const (
	KATSLineAppleReserved                 KATSLine = 0xfce00000
	KATSLineApplyAntiAliasing             KATSLine = 0x800
	KATSLineBreakToNearestCharacter       KATSLine = 0x2000000
	KATSLineDisableAllBaselineAdjustments KATSLine = 0x80000
	KATSLineDisableAllGlyphMorphing       KATSLine = 0x20000
	KATSLineDisableAllJustification       KATSLine = 0x10000
	KATSLineDisableAllKerningAdjustments  KATSLine = 0x40000
	KATSLineDisableAllLayoutOperations    KATSLine = 65536
	KATSLineDisableAllTrackingAdjustments KATSLine = 0x100000
	KATSLineDisableAutoAdjustDisplayPos   KATSLine = 0x4000
	KATSLineDisableNegativeJustification  KATSLine = 0x2000
	KATSLineFillOutToWidth                KATSLine = 0x100
	KATSLineFractDisable                  KATSLine = 0x40
	KATSLineHasNoHangers                  KATSLine = 0x2
	KATSLineHasNoOpticalAlignment         KATSLine = 0x4
	KATSLineIgnoreFontLeading             KATSLine = 0x400
	KATSLineImposeNoAngleForEnds          KATSLine = 0x80
	KATSLineIsDisplayOnly                 KATSLine = 0x1
	KATSLineKeepSpacesOutOfMargin         KATSLine = 0x8
	KATSLineLastNoJustification           KATSLine = 0x20
	KATSLineNoAntiAliasing                KATSLine = 0x1000
	KATSLineNoLayoutOptions               KATSLine = 0
	KATSLineNoSpecialJustification        KATSLine = 0x10
	KATSLineTabAdjustEnabled              KATSLine = 0x200
	KATSLineUseDeviceMetrics              KATSLine = 0x1000000
	KATSLineUseQDRendering                KATSLine = 0x8000
)

func (e KATSLine) String() string {
	switch e {
	case KATSLineAppleReserved:
		return "KATSLineAppleReserved"
	case KATSLineApplyAntiAliasing:
		return "KATSLineApplyAntiAliasing"
	case KATSLineBreakToNearestCharacter:
		return "KATSLineBreakToNearestCharacter"
	case KATSLineDisableAllBaselineAdjustments:
		return "KATSLineDisableAllBaselineAdjustments"
	case KATSLineDisableAllGlyphMorphing:
		return "KATSLineDisableAllGlyphMorphing"
	case KATSLineDisableAllJustification:
		return "KATSLineDisableAllJustification"
	case KATSLineDisableAllKerningAdjustments:
		return "KATSLineDisableAllKerningAdjustments"
	case KATSLineDisableAllTrackingAdjustments:
		return "KATSLineDisableAllTrackingAdjustments"
	case KATSLineDisableAutoAdjustDisplayPos:
		return "KATSLineDisableAutoAdjustDisplayPos"
	case KATSLineDisableNegativeJustification:
		return "KATSLineDisableNegativeJustification"
	case KATSLineFillOutToWidth:
		return "KATSLineFillOutToWidth"
	case KATSLineFractDisable:
		return "KATSLineFractDisable"
	case KATSLineHasNoHangers:
		return "KATSLineHasNoHangers"
	case KATSLineHasNoOpticalAlignment:
		return "KATSLineHasNoOpticalAlignment"
	case KATSLineIgnoreFontLeading:
		return "KATSLineIgnoreFontLeading"
	case KATSLineImposeNoAngleForEnds:
		return "KATSLineImposeNoAngleForEnds"
	case KATSLineIsDisplayOnly:
		return "KATSLineIsDisplayOnly"
	case KATSLineKeepSpacesOutOfMargin:
		return "KATSLineKeepSpacesOutOfMargin"
	case KATSLineLastNoJustification:
		return "KATSLineLastNoJustification"
	case KATSLineNoAntiAliasing:
		return "KATSLineNoAntiAliasing"
	case KATSLineNoLayoutOptions:
		return "KATSLineNoLayoutOptions"
	case KATSLineNoSpecialJustification:
		return "KATSLineNoSpecialJustification"
	case KATSLineTabAdjustEnabled:
		return "KATSLineTabAdjustEnabled"
	case KATSLineUseDeviceMetrics:
		return "KATSLineUseDeviceMetrics"
	case KATSLineUseQDRendering:
		return "KATSLineUseQDRendering"
	default:
		return fmt.Sprintf("KATSLine(%d)", e)
	}
}

type KATSOptionFlags uint

const (
	KATSOptionFlagsActivateDisabled          KATSOptionFlags = 0x1
	KATSOptionFlagsComposeFontPostScriptName KATSOptionFlags = 1
	KATSOptionFlagsDefault                   KATSOptionFlags = 0
	KATSOptionFlagsDefaultScope              KATSOptionFlags = 0
	KATSOptionFlagsDoNotNotify               KATSOptionFlags = 0x1
	KATSOptionFlagsIncludeDisabledMask       KATSOptionFlags = 0x1
	KATSOptionFlagsIterateByPrecedenceMask   KATSOptionFlags = 0x1
	KATSOptionFlagsIterationScopeMask        KATSOptionFlags = 0x7
	KATSOptionFlagsProcessSubdirectories     KATSOptionFlags = 0x1
	KATSOptionFlagsRecordPersistently        KATSOptionFlags = 0x1
	KATSOptionFlagsRestrictedScope           KATSOptionFlags = 0x2
	KATSOptionFlagsUnRestrictedScope         KATSOptionFlags = 0x1
	KATSOptionFlagsUseDataFork               KATSOptionFlags = 768
	KATSOptionFlagsUseDataForkAsResourceFork KATSOptionFlags = 256
	KATSOptionFlagsUseResourceFork           KATSOptionFlags = 512
)

func (e KATSOptionFlags) String() string {
	switch e {
	case KATSOptionFlagsActivateDisabled:
		return "KATSOptionFlagsActivateDisabled"
	case KATSOptionFlagsDefault:
		return "KATSOptionFlagsDefault"
	case KATSOptionFlagsIterationScopeMask:
		return "KATSOptionFlagsIterationScopeMask"
	case KATSOptionFlagsRestrictedScope:
		return "KATSOptionFlagsRestrictedScope"
	case KATSOptionFlagsUseDataFork:
		return "KATSOptionFlagsUseDataFork"
	case KATSOptionFlagsUseDataForkAsResourceFork:
		return "KATSOptionFlagsUseDataForkAsResourceFork"
	case KATSOptionFlagsUseResourceFork:
		return "KATSOptionFlagsUseResourceFork"
	default:
		return fmt.Sprintf("KATSOptionFlags(%d)", e)
	}
}

type KATSStyle uint

const (
	KATSStyleAppleReserved     KATSStyle = 0xfffffff8
	KATSStyleApplyAntiAliasing KATSStyle = 0x2
	KATSStyleApplyHints        KATSStyle = 0
	KATSStyleNoAntiAliasing    KATSStyle = 0x4
	KATSStyleNoHinting         KATSStyle = 0x1
	KATSStyleNoOptions         KATSStyle = 0
)

func (e KATSStyle) String() string {
	switch e {
	case KATSStyleAppleReserved:
		return "KATSStyleAppleReserved"
	case KATSStyleApplyAntiAliasing:
		return "KATSStyleApplyAntiAliasing"
	case KATSStyleApplyHints:
		return "KATSStyleApplyHints"
	case KATSStyleNoAntiAliasing:
		return "KATSStyleNoAntiAliasing"
	case KATSStyleNoHinting:
		return "KATSStyleNoHinting"
	default:
		return fmt.Sprintf("KATSStyle(%d)", e)
	}
}

type KATSU uint

const (
	KATSUAfterWithStreamShiftTag          KATSU = 268
	KATSUAscentTag                        KATSU = 284
	KATSUBaselineClassTag                 KATSU = 274
	KATSUBeforeWithStreamShiftTag         KATSU = 267
	KATSUCGContextTag                     KATSU = 32767
	KATSUCenterTab                        KATSU = 1
	KATSUClearAll                         KATSU = 0xffffffff
	KATSUColorTag                         KATSU = 263
	KATSUCrossStreamShiftTag              KATSU = 269
	KATSUDecimalTab                       KATSU = 3
	KATSUDecompositionFactorTag           KATSU = 273
	KATSUDefaultFontFallbacks             KATSU = 0
	KATSUDescentTag                       KATSU = 285
	KATSUFontMatrixTag                    KATSU = 289
	KATSUFontTag                          KATSU = 261
	KATSUForceHangingTag                  KATSU = 280
	KATSUFromFollowingLayout              KATSU = 0xfffffffd
	KATSUFromPreviousLayout               KATSU = 0xfffffffe
	KATSUFromTextBeginning                KATSU = 0xffffffff
	KATSUGlyphSelectorTag                 KATSU = 287
	KATSUHangingInhibitFactorTag          KATSU = 271
	KATSUImposeWidthTag                   KATSU = 266
	KATSUKerningInhibitFactorTag          KATSU = 272
	KATSULangRegionTag                    KATSU = 264
	KATSULanguageTag                      KATSU = 264
	KATSULastResortOnlyFallback           KATSU = 1
	KATSULayoutOperationOverrideTag       KATSU = 15
	KATSULeadingTag                       KATSU = 286
	KATSULeftTab                          KATSU = 0
	KATSULeftToRightBaseDirection         KATSU = 0
	KATSULineAscentTag                    KATSU = 8
	KATSULineBaselineValuesTag            KATSU = 6
	KATSULineDecimalTabCharacterTag       KATSU = 14
	KATSULineDescentTag                   KATSU = 9
	KATSULineDirectionTag                 KATSU = 3
	KATSULineFlushFactorTag               KATSU = 5
	KATSULineFontFallbacksTag             KATSU = 13
	KATSULineHighlightCGColorTag          KATSU = 17
	KATSULineJustificationFactorTag       KATSU = 4
	KATSULineLangRegionTag                KATSU = 10
	KATSULineLanguageTag                  KATSU = 10
	KATSULineLayoutOptionsTag             KATSU = 7
	KATSULineRotationTag                  KATSU = 2
	KATSULineTextLocatorTag               KATSU = 11
	KATSULineTruncationTag                KATSU = 12
	KATSULineWidthTag                     KATSU = 1
	KATSUMaxATSUITagValue                 KATSU = 65535
	KATSUMaxLineTag                       KATSU = 18
	KATSUMaxStyleTag                      KATSU = 299
	KATSUNoCaretAngleTag                  KATSU = 277
	KATSUNoLigatureSplitTag               KATSU = 276
	KATSUNoOpticalAlignmentTag            KATSU = 279
	KATSUNoSpecialJustificationTag        KATSU = 281
	KATSUNumberTabTypes                   KATSU = 4
	KATSUPriorityJustOverrideTag          KATSU = 275
	KATSUQDBoldfaceTag                    KATSU = 256
	KATSUQDCondensedTag                   KATSU = 259
	KATSUQDExtendedTag                    KATSU = 260
	KATSUQDItalicTag                      KATSU = 257
	KATSUQDUnderlineTag                   KATSU = 258
	KATSURGBAlphaColorTag                 KATSU = 288
	KATSURightTab                         KATSU = 2
	KATSURightToLeftBaseDirection         KATSU = 1
	KATSUSequentialFallbacksExclusive     KATSU = 3
	KATSUSequentialFallbacksPreferred     KATSU = 2
	KATSUSizeTag                          KATSU = 262
	KATSUStyleDropShadowBlurOptionTag     KATSU = 296
	KATSUStyleDropShadowColorOptionTag    KATSU = 298
	KATSUStyleDropShadowOffsetOptionTag   KATSU = 297
	KATSUStyleDropShadowTag               KATSU = 295
	KATSUStyleRenderingOptionsTag         KATSU = 283
	KATSUStyleStrikeThroughColorOptionTag KATSU = 294
	KATSUStyleStrikeThroughCountOptionTag KATSU = 293
	KATSUStyleStrikeThroughTag            KATSU = 292
	KATSUStyleTextLocatorTag              KATSU = 282
	KATSUStyleUnderlineColorOptionTag     KATSU = 291
	KATSUStyleUnderlineCountOptionTag     KATSU = 290
	KATSUSuppressCrossKerningTag          KATSU = 278
	KATSUToTextEnd                        KATSU = 0xffffffff
	KATSUTrackingTag                      KATSU = 270
	KATSUTruncFeatNoSquishing             KATSU = 0x8
	KATSUTruncateEnd                      KATSU = 2
	KATSUTruncateMiddle                   KATSU = 3
	KATSUTruncateNone                     KATSU = 0
	KATSUTruncateSpecificationMask        KATSU = 0x7
	KATSUTruncateStart                    KATSU = 1
	KATSUUseGrafPortPenLoc                KATSU = 0xffffffff
	KATSUVerticalCharacterTag             KATSU = 265
)

func (e KATSU) String() string {
	switch e {
	case KATSUAfterWithStreamShiftTag:
		return "KATSUAfterWithStreamShiftTag"
	case KATSUAscentTag:
		return "KATSUAscentTag"
	case KATSUBaselineClassTag:
		return "KATSUBaselineClassTag"
	case KATSUBeforeWithStreamShiftTag:
		return "KATSUBeforeWithStreamShiftTag"
	case KATSUCGContextTag:
		return "KATSUCGContextTag"
	case KATSUCenterTab:
		return "KATSUCenterTab"
	case KATSUClearAll:
		return "KATSUClearAll"
	case KATSUColorTag:
		return "KATSUColorTag"
	case KATSUCrossStreamShiftTag:
		return "KATSUCrossStreamShiftTag"
	case KATSUDecimalTab:
		return "KATSUDecimalTab"
	case KATSUDecompositionFactorTag:
		return "KATSUDecompositionFactorTag"
	case KATSUDefaultFontFallbacks:
		return "KATSUDefaultFontFallbacks"
	case KATSUDescentTag:
		return "KATSUDescentTag"
	case KATSUFontMatrixTag:
		return "KATSUFontMatrixTag"
	case KATSUFontTag:
		return "KATSUFontTag"
	case KATSUForceHangingTag:
		return "KATSUForceHangingTag"
	case KATSUFromFollowingLayout:
		return "KATSUFromFollowingLayout"
	case KATSUFromPreviousLayout:
		return "KATSUFromPreviousLayout"
	case KATSUGlyphSelectorTag:
		return "KATSUGlyphSelectorTag"
	case KATSUHangingInhibitFactorTag:
		return "KATSUHangingInhibitFactorTag"
	case KATSUImposeWidthTag:
		return "KATSUImposeWidthTag"
	case KATSUKerningInhibitFactorTag:
		return "KATSUKerningInhibitFactorTag"
	case KATSULangRegionTag:
		return "KATSULangRegionTag"
	case KATSULayoutOperationOverrideTag:
		return "KATSULayoutOperationOverrideTag"
	case KATSULeadingTag:
		return "KATSULeadingTag"
	case KATSULineAscentTag:
		return "KATSULineAscentTag"
	case KATSULineBaselineValuesTag:
		return "KATSULineBaselineValuesTag"
	case KATSULineDecimalTabCharacterTag:
		return "KATSULineDecimalTabCharacterTag"
	case KATSULineDescentTag:
		return "KATSULineDescentTag"
	case KATSULineFlushFactorTag:
		return "KATSULineFlushFactorTag"
	case KATSULineFontFallbacksTag:
		return "KATSULineFontFallbacksTag"
	case KATSULineHighlightCGColorTag:
		return "KATSULineHighlightCGColorTag"
	case KATSULineJustificationFactorTag:
		return "KATSULineJustificationFactorTag"
	case KATSULineLangRegionTag:
		return "KATSULineLangRegionTag"
	case KATSULineLayoutOptionsTag:
		return "KATSULineLayoutOptionsTag"
	case KATSULineRotationTag:
		return "KATSULineRotationTag"
	case KATSULineTextLocatorTag:
		return "KATSULineTextLocatorTag"
	case KATSULineTruncationTag:
		return "KATSULineTruncationTag"
	case KATSUMaxATSUITagValue:
		return "KATSUMaxATSUITagValue"
	case KATSUMaxLineTag:
		return "KATSUMaxLineTag"
	case KATSUMaxStyleTag:
		return "KATSUMaxStyleTag"
	case KATSUNoCaretAngleTag:
		return "KATSUNoCaretAngleTag"
	case KATSUNoLigatureSplitTag:
		return "KATSUNoLigatureSplitTag"
	case KATSUNoOpticalAlignmentTag:
		return "KATSUNoOpticalAlignmentTag"
	case KATSUNoSpecialJustificationTag:
		return "KATSUNoSpecialJustificationTag"
	case KATSUPriorityJustOverrideTag:
		return "KATSUPriorityJustOverrideTag"
	case KATSUQDBoldfaceTag:
		return "KATSUQDBoldfaceTag"
	case KATSUQDCondensedTag:
		return "KATSUQDCondensedTag"
	case KATSUQDExtendedTag:
		return "KATSUQDExtendedTag"
	case KATSUQDItalicTag:
		return "KATSUQDItalicTag"
	case KATSUQDUnderlineTag:
		return "KATSUQDUnderlineTag"
	case KATSURGBAlphaColorTag:
		return "KATSURGBAlphaColorTag"
	case KATSUSizeTag:
		return "KATSUSizeTag"
	case KATSUStyleDropShadowBlurOptionTag:
		return "KATSUStyleDropShadowBlurOptionTag"
	case KATSUStyleDropShadowColorOptionTag:
		return "KATSUStyleDropShadowColorOptionTag"
	case KATSUStyleDropShadowOffsetOptionTag:
		return "KATSUStyleDropShadowOffsetOptionTag"
	case KATSUStyleDropShadowTag:
		return "KATSUStyleDropShadowTag"
	case KATSUStyleRenderingOptionsTag:
		return "KATSUStyleRenderingOptionsTag"
	case KATSUStyleStrikeThroughColorOptionTag:
		return "KATSUStyleStrikeThroughColorOptionTag"
	case KATSUStyleStrikeThroughCountOptionTag:
		return "KATSUStyleStrikeThroughCountOptionTag"
	case KATSUStyleStrikeThroughTag:
		return "KATSUStyleStrikeThroughTag"
	case KATSUStyleTextLocatorTag:
		return "KATSUStyleTextLocatorTag"
	case KATSUStyleUnderlineColorOptionTag:
		return "KATSUStyleUnderlineColorOptionTag"
	case KATSUStyleUnderlineCountOptionTag:
		return "KATSUStyleUnderlineCountOptionTag"
	case KATSUSuppressCrossKerningTag:
		return "KATSUSuppressCrossKerningTag"
	case KATSUTrackingTag:
		return "KATSUTrackingTag"
	case KATSUVerticalCharacterTag:
		return "KATSUVerticalCharacterTag"
	default:
		return fmt.Sprintf("KATSU(%d)", e)
	}
}

type KATSUBackground uint

const (
	KATSUBackgroundCallback KATSUBackground = 1
	KATSUBackgroundColor    KATSUBackground = 0
)

func (e KATSUBackground) String() string {
	switch e {
	case KATSUBackgroundCallback:
		return "KATSUBackgroundCallback"
	case KATSUBackgroundColor:
		return "KATSUBackgroundColor"
	default:
		return fmt.Sprintf("KATSUBackground(%d)", e)
	}
}

type KATSUBy uint

const (
	KATSUByCharacter          KATSUBy = 0
	KATSUByCharacterCluster   KATSUBy = 3
	KATSUByCluster            KATSUBy = 1
	KATSUByTypographicCluster KATSUBy = 1
	KATSUByWord               KATSUBy = 2
)

func (e KATSUBy) String() string {
	switch e {
	case KATSUByCharacter:
		return "KATSUByCharacter"
	case KATSUByCharacterCluster:
		return "KATSUByCharacterCluster"
	case KATSUByCluster:
		return "KATSUByCluster"
	case KATSUByWord:
		return "KATSUByWord"
	default:
		return fmt.Sprintf("KATSUBy(%d)", e)
	}
}

type KATSUDataStreamUnicodeStyled uint

const (
	KATSUDataStreamUnicodeStyledText KATSUDataStreamUnicodeStyled = 'u'<<24 | 's'<<16 | 't'<<8 | 'l' // 'ustl'
)

func (e KATSUDataStreamUnicodeStyled) String() string {
	switch e {
	case KATSUDataStreamUnicodeStyledText:
		return "KATSUDataStreamUnicodeStyledText"
	default:
		return fmt.Sprintf("KATSUDataStreamUnicodeStyled(%d)", e)
	}
}

type KATSUDirectData uint

const (
	KATSUDirectDataAdvanceDeltaFixedArray               KATSUDirectData = 0
	KATSUDirectDataBaselineDeltaFixedArray              KATSUDirectData = 1
	KATSUDirectDataDeviceDeltaSInt16Array               KATSUDirectData = 2
	KATSUDirectDataLayoutRecordATSLayoutRecordCurrent   KATSUDirectData = 100
	KATSUDirectDataLayoutRecordATSLayoutRecordVersion1  KATSUDirectData = 100
	KATSUDirectDataStyleIndexUInt16Array                KATSUDirectData = 3
	KATSUDirectDataStyleSettingATSUStyleSettingRefArray KATSUDirectData = 4
)

func (e KATSUDirectData) String() string {
	switch e {
	case KATSUDirectDataAdvanceDeltaFixedArray:
		return "KATSUDirectDataAdvanceDeltaFixedArray"
	case KATSUDirectDataBaselineDeltaFixedArray:
		return "KATSUDirectDataBaselineDeltaFixedArray"
	case KATSUDirectDataDeviceDeltaSInt16Array:
		return "KATSUDirectDataDeviceDeltaSInt16Array"
	case KATSUDirectDataLayoutRecordATSLayoutRecordCurrent:
		return "KATSUDirectDataLayoutRecordATSLayoutRecordCurrent"
	case KATSUDirectDataStyleIndexUInt16Array:
		return "KATSUDirectDataStyleIndexUInt16Array"
	case KATSUDirectDataStyleSettingATSUStyleSettingRefArray:
		return "KATSUDirectDataStyleSettingATSUStyleSettingRefArray"
	default:
		return fmt.Sprintf("KATSUDirectData(%d)", e)
	}
}

type KATSUFlattenOptionNoOptions uint

const (
	KATSUFlattenOptionNoOptionsMask KATSUFlattenOptionNoOptions = 0
)

func (e KATSUFlattenOptionNoOptions) String() string {
	switch e {
	case KATSUFlattenOptionNoOptionsMask:
		return "KATSUFlattenOptionNoOptionsMask"
	default:
		return fmt.Sprintf("KATSUFlattenOptionNoOptions(%d)", e)
	}
}

type KATSUInvalidFontI uint

const (
	KATSUInvalidFontID KATSUInvalidFontI = 0
)

func (e KATSUInvalidFontI) String() string {
	switch e {
	case KATSUInvalidFontID:
		return "KATSUInvalidFontID"
	default:
		return fmt.Sprintf("KATSUInvalidFontI(%d)", e)
	}
}

type KATSULayoutOperation uint

const (
	KATSULayoutOperationAppleReserved        KATSULayoutOperation = 0xffffffc0
	KATSULayoutOperationBaselineAdjustment   KATSULayoutOperation = 0x8
	KATSULayoutOperationJustification        KATSULayoutOperation = 0x1
	KATSULayoutOperationKerningAdjustment    KATSULayoutOperation = 0x4
	KATSULayoutOperationMorph                KATSULayoutOperation = 0x2
	KATSULayoutOperationNone                 KATSULayoutOperation = 0
	KATSULayoutOperationPostLayoutAdjustment KATSULayoutOperation = 0x20
	KATSULayoutOperationTrackingAdjustment   KATSULayoutOperation = 0x10
)

func (e KATSULayoutOperation) String() string {
	switch e {
	case KATSULayoutOperationAppleReserved:
		return "KATSULayoutOperationAppleReserved"
	case KATSULayoutOperationBaselineAdjustment:
		return "KATSULayoutOperationBaselineAdjustment"
	case KATSULayoutOperationJustification:
		return "KATSULayoutOperationJustification"
	case KATSULayoutOperationKerningAdjustment:
		return "KATSULayoutOperationKerningAdjustment"
	case KATSULayoutOperationMorph:
		return "KATSULayoutOperationMorph"
	case KATSULayoutOperationNone:
		return "KATSULayoutOperationNone"
	case KATSULayoutOperationPostLayoutAdjustment:
		return "KATSULayoutOperationPostLayoutAdjustment"
	case KATSULayoutOperationTrackingAdjustment:
		return "KATSULayoutOperationTrackingAdjustment"
	default:
		return fmt.Sprintf("KATSULayoutOperation(%d)", e)
	}
}

type KATSULayoutOperationCallbackStatus int

const (
	KATSULayoutOperationCallbackStatusContinue KATSULayoutOperationCallbackStatus = 0x1
	KATSULayoutOperationCallbackStatusHandled  KATSULayoutOperationCallbackStatus = 0
)

func (e KATSULayoutOperationCallbackStatus) String() string {
	switch e {
	case KATSULayoutOperationCallbackStatusContinue:
		return "KATSULayoutOperationCallbackStatusContinue"
	case KATSULayoutOperationCallbackStatusHandled:
		return "KATSULayoutOperationCallbackStatusHandled"
	default:
		return fmt.Sprintf("KATSULayoutOperationCallbackStatus(%d)", e)
	}
}

type KATSUNo uint

const (
	KATSUNoSelector KATSUNo = 0xffff
)

func (e KATSUNo) String() string {
	switch e {
	case KATSUNoSelector:
		return "KATSUNoSelector"
	default:
		return fmt.Sprintf("KATSUNo(%d)", e)
	}
}

type KATSUStrongly uint

const (
	KATSUStronglyHorizontal KATSUStrongly = 0
	KATSUStronglyVertical   KATSUStrongly = 1
)

func (e KATSUStrongly) String() string {
	switch e {
	case KATSUStronglyHorizontal:
		return "KATSUStronglyHorizontal"
	case KATSUStronglyVertical:
		return "KATSUStronglyVertical"
	default:
		return fmt.Sprintf("KATSUStrongly(%d)", e)
	}
}

type KATSUStyle uint

const (
	KATSUStyleContainedBy     KATSUStyle = 3
	KATSUStyleContains        KATSUStyle = 1
	KATSUStyleDoubleLineCount KATSUStyle = 2
	KATSUStyleEquals          KATSUStyle = 2
	KATSUStyleSingleLineCount KATSUStyle = 1
	KATSUStyleUnequal         KATSUStyle = 0
)

func (e KATSUStyle) String() string {
	switch e {
	case KATSUStyleContainedBy:
		return "KATSUStyleContainedBy"
	case KATSUStyleContains:
		return "KATSUStyleContains"
	case KATSUStyleDoubleLineCount:
		return "KATSUStyleDoubleLineCount"
	case KATSUStyleUnequal:
		return "KATSUStyleUnequal"
	default:
		return fmt.Sprintf("KATSUStyle(%d)", e)
	}
}

type KATSUUnFlattenOptionNoOptions uint

const (
	KATSUUnFlattenOptionNoOptionsMask KATSUUnFlattenOptionNoOptions = 0
)

func (e KATSUUnFlattenOptionNoOptions) String() string {
	switch e {
	case KATSUUnFlattenOptionNoOptionsMask:
		return "KATSUUnFlattenOptionNoOptionsMask"
	default:
		return fmt.Sprintf("KATSUUnFlattenOptionNoOptions(%d)", e)
	}
}

type KATSUUseLineControl uint

const (
	KATSUUseLineControlWidth KATSUUseLineControl = 0x7fffffff
)

func (e KATSUUseLineControl) String() string {
	switch e {
	case KATSUUseLineControlWidth:
		return "KATSUUseLineControlWidth"
	default:
		return fmt.Sprintf("KATSUUseLineControl(%d)", e)
	}
}

type KATSUse uint

const (
	KATSUseCaretOrigins      KATSUse = 0
	KATSUseDeviceOrigins     KATSUse = 1
	KATSUseFractionalOrigins KATSUse = 2
	KATSUseOriginFlags       KATSUse = 3
)

func (e KATSUse) String() string {
	switch e {
	case KATSUseCaretOrigins:
		return "KATSUseCaretOrigins"
	case KATSUseDeviceOrigins:
		return "KATSUseDeviceOrigins"
	case KATSUseFractionalOrigins:
		return "KATSUseFractionalOrigins"
	case KATSUseOriginFlags:
		return "KATSUseOriginFlags"
	default:
		return fmt.Sprintf("KATSUse(%d)", e)
	}
}

type KAlign uint

const (
	KAlignAbsoluteCenter   KAlign = 1
	KAlignBottom           KAlign = 0x3
	KAlignBottomLeft       KAlign = 3
	KAlignBottomRight      KAlign = 3
	KAlignCenterBottom     KAlign = 3
	KAlignCenterLeft       KAlign = 1
	KAlignCenterRight      KAlign = 1
	KAlignCenterTop        KAlign = 2
	KAlignHorizontalCenter KAlign = 0x4
	KAlignLeft             KAlign = 0x8
	KAlignNone             KAlign = 0
	KAlignRight            KAlign = 0xc
	KAlignTop              KAlign = 0x2
	KAlignTopLeft          KAlign = 2
	KAlignTopRight         KAlign = 2
	KAlignVerticalCenter   KAlign = 0x1
)

func (e KAlign) String() string {
	switch e {
	case KAlignAbsoluteCenter:
		return "KAlignAbsoluteCenter"
	case KAlignBottom:
		return "KAlignBottom"
	case KAlignCenterTop:
		return "KAlignCenterTop"
	case KAlignHorizontalCenter:
		return "KAlignHorizontalCenter"
	case KAlignLeft:
		return "KAlignLeft"
	case KAlignNone:
		return "KAlignNone"
	case KAlignRight:
		return "KAlignRight"
	default:
		return fmt.Sprintf("KAlign(%d)", e)
	}
}

type KAllPPDDomains uint

const (
	// KAllPPDDomainsValue: # Discussion
	KAllPPDDomainsValue KAllPPDDomains = 1
	// KCUPSPPDDomain: # Discussion
	KCUPSPPDDomain KAllPPDDomains = 6
	// KLocalPPDDomain: # Discussion
	KLocalPPDDomain KAllPPDDomains = 3
	// KNetworkPPDDomain: # Discussion
	KNetworkPPDDomain KAllPPDDomains = 4
	// KSystemPPDDomain: # Discussion
	KSystemPPDDomain KAllPPDDomains = 2
	// KUserPPDDomain: # Discussion
	KUserPPDDomain KAllPPDDomains = 5
)

func (e KAllPPDDomains) String() string {
	switch e {
	case KAllPPDDomainsValue:
		return "KAllPPDDomainsValue"
	case KCUPSPPDDomain:
		return "KCUPSPPDDomain"
	case KLocalPPDDomain:
		return "KLocalPPDDomain"
	case KNetworkPPDDomain:
		return "KNetworkPPDDomain"
	case KSystemPPDDomain:
		return "KSystemPPDDomain"
	case KUserPPDDomain:
		return "KUserPPDDomain"
	default:
		return fmt.Sprintf("KAllPPDDomains(%d)", e)
	}
}

type KAudioUnit uint

const (
	// KAudioUnitProperty_SpeechChannel: The speech channel property in the speech synthesis audio unit.
	KAudioUnitProperty_SpeechChannel KAudioUnit = 3331
	// KAudioUnitProperty_Voice: The voice property in the speech synthesis audio unit.
	KAudioUnitProperty_Voice KAudioUnit = 3330
	// KAudioUnitSubType_SpeechSynthesis: The speech synthesis component subtype used in the creation of a speech synthesis audio unit.
	KAudioUnitSubType_SpeechSynthesis KAudioUnit = 't'<<24 | 't'<<16 | 's'<<8 | 'p' // 'ttsp'
)

func (e KAudioUnit) String() string {
	switch e {
	case KAudioUnitProperty_SpeechChannel:
		return "KAudioUnitProperty_SpeechChannel"
	case KAudioUnitProperty_Voice:
		return "KAudioUnitProperty_Voice"
	case KAudioUnitSubType_SpeechSynthesis:
		return "KAudioUnitSubType_SpeechSynthesis"
	default:
		return fmt.Sprintf("KAudioUnit(%d)", e)
	}
}

type KDefaultCMM uint

const (
	// KDefaultCMMSignature: Signature for the default CMM supplied with the ColorSync Manager.
	KDefaultCMMSignature KDefaultCMM = 'a'<<24 | 'p'<<16 | 'p'<<8 | 'l' // 'appl'
)

func (e KDefaultCMM) String() string {
	switch e {
	case KDefaultCMMSignature:
		return "KDefaultCMMSignature"
	default:
		return fmt.Sprintf("KDefaultCMM(%d)", e)
	}
}

type KFM uint

const (
	KFMFontCallbackFilterSelector       KFM = 5
	KFMFontContainerFilterSelector      KFM = 2
	KFMFontDirectoryFilterSelector      KFM = 6
	KFMFontFamilyCallbackFilterSelector KFM = 4
	KFMFontFileRefFilterSelector        KFM = 10
	KFMFontTechnologyFilterSelector     KFM = 1
	KFMGenerationFilterSelector         KFM = 3
	KFMPostScriptFontTechnology         KFM = 't'<<24 | 'y'<<16 | 'p'<<8 | '1' // 'typ1'
	KFMTrueTypeFontTechnology           KFM = 't'<<24 | 'r'<<16 | 'u'<<8 | 'e' // 'true'
)

func (e KFM) String() string {
	switch e {
	case KFMFontCallbackFilterSelector:
		return "KFMFontCallbackFilterSelector"
	case KFMFontContainerFilterSelector:
		return "KFMFontContainerFilterSelector"
	case KFMFontDirectoryFilterSelector:
		return "KFMFontDirectoryFilterSelector"
	case KFMFontFamilyCallbackFilterSelector:
		return "KFMFontFamilyCallbackFilterSelector"
	case KFMFontFileRefFilterSelector:
		return "KFMFontFileRefFilterSelector"
	case KFMFontTechnologyFilterSelector:
		return "KFMFontTechnologyFilterSelector"
	case KFMGenerationFilterSelector:
		return "KFMGenerationFilterSelector"
	case KFMPostScriptFontTechnology:
		return "KFMPostScriptFontTechnology"
	case KFMTrueTypeFontTechnology:
		return "KFMTrueTypeFontTechnology"
	default:
		return fmt.Sprintf("KFM(%d)", e)
	}
}

type KFMCurrentFilter uint

const (
	KFMCurrentFilterFormat KFMCurrentFilter = 0
)

func (e KFMCurrentFilter) String() string {
	switch e {
	case KFMCurrentFilterFormat:
		return "KFMCurrentFilterFormat"
	default:
		return fmt.Sprintf("KFMCurrentFilter(%d)", e)
	}
}

type KGlyphCollection uint

const (
	KGlyphCollectionAdobeCNS1   KGlyphCollection = 1
	KGlyphCollectionAdobeGB1    KGlyphCollection = 2
	KGlyphCollectionAdobeJapan1 KGlyphCollection = 3
	KGlyphCollectionAdobeJapan2 KGlyphCollection = 4
	KGlyphCollectionAdobeKorea1 KGlyphCollection = 5
	KGlyphCollectionGID         KGlyphCollection = 0
	KGlyphCollectionUnspecified KGlyphCollection = 0xff
)

func (e KGlyphCollection) String() string {
	switch e {
	case KGlyphCollectionAdobeCNS1:
		return "KGlyphCollectionAdobeCNS1"
	case KGlyphCollectionAdobeGB1:
		return "KGlyphCollectionAdobeGB1"
	case KGlyphCollectionAdobeJapan1:
		return "KGlyphCollectionAdobeJapan1"
	case KGlyphCollectionAdobeJapan2:
		return "KGlyphCollectionAdobeJapan2"
	case KGlyphCollectionAdobeKorea1:
		return "KGlyphCollectionAdobeKorea1"
	case KGlyphCollectionGID:
		return "KGlyphCollectionGID"
	case KGlyphCollectionUnspecified:
		return "KGlyphCollectionUnspecified"
	default:
		return fmt.Sprintf("KGlyphCollection(%d)", e)
	}
}

type KHIShapeEnumerate uint

const (
	KHIShapeEnumerateInit      KHIShapeEnumerate = 1
	KHIShapeEnumerateRect      KHIShapeEnumerate = 2
	KHIShapeEnumerateTerminate KHIShapeEnumerate = 3
)

func (e KHIShapeEnumerate) String() string {
	switch e {
	case KHIShapeEnumerateInit:
		return "KHIShapeEnumerateInit"
	case KHIShapeEnumerateRect:
		return "KHIShapeEnumerateRect"
	case KHIShapeEnumerateTerminate:
		return "KHIShapeEnumerateTerminate"
	default:
		return fmt.Sprintf("KHIShapeEnumerate(%d)", e)
	}
}

type KHIShapeParseFrom uint

const (
	KHIShapeParseFromBottom      KHIShapeParseFrom = 1
	KHIShapeParseFromBottomRight KHIShapeParseFrom = 1
	KHIShapeParseFromLeft        KHIShapeParseFrom = 0
	KHIShapeParseFromRight       KHIShapeParseFrom = 2
	KHIShapeParseFromTop         KHIShapeParseFrom = 0
	KHIShapeParseFromTopLeft     KHIShapeParseFrom = 0
)

func (e KHIShapeParseFrom) String() string {
	switch e {
	case KHIShapeParseFromBottom:
		return "KHIShapeParseFromBottom"
	case KHIShapeParseFromLeft:
		return "KHIShapeParseFromLeft"
	case KHIShapeParseFromRight:
		return "KHIShapeParseFromRight"
	default:
		return fmt.Sprintf("KHIShapeParseFrom(%d)", e)
	}
}

type KIC uint

const (
	KICComponentVersion KIC = 0
	KICCreator          KIC = 'I'<<24 | 'C'<<16 | 'A'<<8 | 'p' // 'ICAp'
	KICFileType         KIC = 'I'<<24 | 'C'<<16 | 'A'<<8 | 'p' // 'ICAp'
	KICNumVersion       KIC = 1
)

func (e KIC) String() string {
	switch e {
	case KICComponentVersion:
		return "KICComponentVersion"
	case KICCreator:
		return "KICCreator"
	case KICNumVersion:
		return "KICNumVersion"
	default:
		return fmt.Sprintf("KIC(%d)", e)
	}
}

type KICAttr uint

const (
	KICAttrLockedBit    KICAttr = 0
	KICAttrLockedMask   KICAttr = 0x1
	KICAttrNoChange     KICAttr = 0xffffffff
	KICAttrVolatileBit  KICAttr = 1
	KICAttrVolatileMask KICAttr = 0x2
)

func (e KICAttr) String() string {
	switch e {
	case KICAttrLockedBit:
		return "KICAttrLockedBit"
	case KICAttrLockedMask:
		return "KICAttrLockedMask"
	case KICAttrNoChange:
		return "KICAttrNoChange"
	case KICAttrVolatileMask:
		return "KICAttrVolatileMask"
	default:
		return fmt.Sprintf("KICAttr(%d)", e)
	}
}

type KICComponentInterface uint

const (
	KICComponentInterfaceVersion  KICComponentInterface = 262144
	KICComponentInterfaceVersion0 KICComponentInterface = 0
	KICComponentInterfaceVersion1 KICComponentInterface = 0x10000
	KICComponentInterfaceVersion2 KICComponentInterface = 0x20000
	KICComponentInterfaceVersion3 KICComponentInterface = 0x30000
	KICComponentInterfaceVersion4 KICComponentInterface = 0x40000
)

func (e KICComponentInterface) String() string {
	switch e {
	case KICComponentInterfaceVersion:
		return "KICComponentInterfaceVersion"
	case KICComponentInterfaceVersion0:
		return "KICComponentInterfaceVersion0"
	case KICComponentInterfaceVersion1:
		return "KICComponentInterfaceVersion1"
	case KICComponentInterfaceVersion2:
		return "KICComponentInterfaceVersion2"
	case KICComponentInterfaceVersion3:
		return "KICComponentInterfaceVersion3"
	default:
		return fmt.Sprintf("KICComponentInterface(%d)", e)
	}
}

type KICEditPreferenceEventClass uint

const (
	KICEditPreferenceEvent           KICEditPreferenceEventClass = 'I'<<24 | 'C'<<16 | 'A'<<8 | 'p' // 'ICAp'
	KICEditPreferenceEventClassValue KICEditPreferenceEventClass = 'I'<<24 | 'C'<<16 | 'A'<<8 | 'p' // 'ICAp'
	KeyICEditPreferenceDestination   KICEditPreferenceEventClass = 'd'<<24 | 'e'<<16 | 's'<<8 | 't' // 'dest'
)

func (e KICEditPreferenceEventClass) String() string {
	switch e {
	case KICEditPreferenceEvent:
		return "KICEditPreferenceEvent"
	case KeyICEditPreferenceDestination:
		return "KeyICEditPreferenceDestination"
	default:
		return fmt.Sprintf("KICEditPreferenceEventClass(%d)", e)
	}
}

type KICFileSpecHeader uint

const (
	KICFileSpecHeaderSize KICFileSpecHeader = 0
)

func (e KICFileSpecHeader) String() string {
	switch e {
	case KICFileSpecHeaderSize:
		return "KICFileSpecHeaderSize"
	default:
		return fmt.Sprintf("KICFileSpecHeader(%d)", e)
	}
}

type KICMap uint

const (
	KICMapBinaryBit        KICMap = 0
	KICMapBinaryMask       KICMap = 0x1
	KICMapDataForkBit      KICMap = 2
	KICMapDataForkMask     KICMap = 0x4
	KICMapNotIncomingBit   KICMap = 4
	KICMapNotIncomingMask  KICMap = 0x10
	KICMapNotOutgoingBit   KICMap = 5
	KICMapNotOutgoingMask  KICMap = 0x20
	KICMapPostBit          KICMap = 3
	KICMapPostMask         KICMap = 0x8
	KICMapResourceForkBit  KICMap = 1
	KICMapResourceForkMask KICMap = 0x2
)

func (e KICMap) String() string {
	switch e {
	case KICMapBinaryBit:
		return "KICMapBinaryBit"
	case KICMapBinaryMask:
		return "KICMapBinaryMask"
	case KICMapDataForkBit:
		return "KICMapDataForkBit"
	case KICMapDataForkMask:
		return "KICMapDataForkMask"
	case KICMapNotIncomingMask:
		return "KICMapNotIncomingMask"
	case KICMapNotOutgoingBit:
		return "KICMapNotOutgoingBit"
	case KICMapNotOutgoingMask:
		return "KICMapNotOutgoingMask"
	case KICMapPostBit:
		return "KICMapPostBit"
	case KICMapPostMask:
		return "KICMapPostMask"
	default:
		return fmt.Sprintf("KICMap(%d)", e)
	}
}

type KICMapFixed uint

const (
	KICMapFixedLength KICMapFixed = 22
)

func (e KICMapFixed) String() string {
	switch e {
	case KICMapFixedLength:
		return "KICMapFixedLength"
	default:
		return fmt.Sprintf("KICMapFixed(%d)", e)
	}
}

type KICNilProfileI uint

const (
	KICNilProfileID KICNilProfileI = 0
)

func (e KICNilProfileI) String() string {
	switch e {
	case KICNilProfileID:
		return "KICNilProfileID"
	default:
		return fmt.Sprintf("KICNilProfileI(%d)", e)
	}
}

type KICNoUserInteraction uint

const (
	KICNoUserInteractionBit  KICNoUserInteraction = 0
	KICNoUserInteractionMask KICNoUserInteraction = 0x1
)

func (e KICNoUserInteraction) String() string {
	switch e {
	case KICNoUserInteractionBit:
		return "KICNoUserInteractionBit"
	case KICNoUserInteractionMask:
		return "KICNoUserInteractionMask"
	default:
		return fmt.Sprintf("KICNoUserInteraction(%d)", e)
	}
}

type KICServices uint

const (
	KICServicesTCPBit  KICServices = 0
	KICServicesTCPMask KICServices = 0x1
	KICServicesUDPBit  KICServices = 1
	KICServicesUDPMask KICServices = 0x2
)

func (e KICServices) String() string {
	switch e {
	case KICServicesTCPBit:
		return "KICServicesTCPBit"
	case KICServicesTCPMask:
		return "KICServicesTCPMask"
	case KICServicesUDPMask:
		return "KICServicesUDPMask"
	default:
		return fmt.Sprintf("KICServices(%d)", e)
	}
}

type KImmediate uint

const (
	// KEndOfSentence: Speech should be paused or stopped at the end of the sentence.
	KEndOfSentence KImmediate = 2
	// KEndOfWord: Speech should be paused or stopped at the endof the word.
	KEndOfWord KImmediate = 1
	// KImmediateValue: Speech should be paused or stopped immediately.
	KImmediateValue KImmediate = 0
)

func (e KImmediate) String() string {
	switch e {
	case KEndOfSentence:
		return "KEndOfSentence"
	case KEndOfWord:
		return "KEndOfWord"
	case KImmediateValue:
		return "KImmediateValue"
	default:
		return fmt.Sprintf("KImmediate(%d)", e)
	}
}

type KInternetEventClass uint

const (
	KAEFetchURL              KInternetEventClass = 'F'<<24 | 'U'<<16 | 'R'<<8 | 'L' // 'FURL'
	KAEGetURL                KInternetEventClass = 'G'<<24 | 'U'<<16 | 'R'<<8 | 'L' // 'GURL'
	KInternetEventClassValue KInternetEventClass = 'G'<<24 | 'U'<<16 | 'R'<<8 | 'L' // 'GURL'
	KeyAEAttaching           KInternetEventClass = 'A'<<24 | 't'<<16 | 'c'<<8 | 'h' // 'Atch'
)

func (e KInternetEventClass) String() string {
	switch e {
	case KAEFetchURL:
		return "KAEFetchURL"
	case KAEGetURL:
		return "KAEGetURL"
	case KeyAEAttaching:
		return "KeyAEAttaching"
	default:
		return fmt.Sprintf("KInternetEventClass(%d)", e)
	}
}

type KInvalid int

const (
	KInvalidFont       KInvalid = 0
	KInvalidFontFamily KInvalid = -1
	KInvalidGeneration KInvalid = 0
)

func (e KInvalid) String() string {
	switch e {
	case KInvalidFont:
		return "KInvalidFont"
	case KInvalidFontFamily:
		return "KInvalidFontFamily"
	default:
		return fmt.Sprintf("KInvalid(%d)", e)
	}
}

type KInvertHighlighting uint

const (
	KInvertHighlightingValue KInvertHighlighting = 0
	KRedrawHighlighting      KInvertHighlighting = 1
)

func (e KInvertHighlighting) String() string {
	switch e {
	case KInvertHighlightingValue:
		return "KInvertHighlightingValue"
	case KRedrawHighlighting:
		return "KRedrawHighlighting"
	default:
		return fmt.Sprintf("KInvertHighlighting(%d)", e)
	}
}

type KNeuter uint

const (
	// KFemale: Female voice.
	KFemale KNeuter = 2
	// KMale: Male voice.
	KMale KNeuter = 1
	// KNeuterValue: Neuter voice.
	KNeuterValue KNeuter = 0
)

func (e KNeuter) String() string {
	switch e {
	case KFemale:
		return "KFemale"
	case KMale:
		return "KMale"
	case KNeuterValue:
		return "KNeuterValue"
	default:
		return fmt.Sprintf("KNeuter(%d)", e)
	}
}

type KNoConstraint uint

const (
	KHorizontalConstraint KNoConstraint = 2
	KNoConstraintValue    KNoConstraint = 0
	KVerticalConstraint   KNoConstraint = 1
)

func (e KNoConstraint) String() string {
	switch e {
	case KHorizontalConstraint:
		return "KHorizontalConstraint"
	case KNoConstraintValue:
		return "KNoConstraintValue"
	case KVerticalConstraint:
		return "KVerticalConstraint"
	default:
		return fmt.Sprintf("KNoConstraint(%d)", e)
	}
}

type KNoEndingProsody uint

const (
	// KNoEndingProsodyValue: # Discussion
	KNoEndingProsodyValue KNoEndingProsody = 1
	// KNoSpeechInterrupt: # Discussion
	KNoSpeechInterrupt KNoEndingProsody = 2
	// KPreflightThenPause: # Discussion
	KPreflightThenPause KNoEndingProsody = 4
)

func (e KNoEndingProsody) String() string {
	switch e {
	case KNoEndingProsodyValue:
		return "KNoEndingProsodyValue"
	case KNoSpeechInterrupt:
		return "KNoSpeechInterrupt"
	case KPreflightThenPause:
		return "KPreflightThenPause"
	default:
		return fmt.Sprintf("KNoEndingProsody(%d)", e)
	}
}

type KNoProcess uint

const (
	KCurrentProcess KNoProcess = 2
	KNoProcessValue KNoProcess = 0
	KSystemProcess  KNoProcess = 1
)

func (e KNoProcess) String() string {
	switch e {
	case KCurrentProcess:
		return "KCurrentProcess"
	case KNoProcessValue:
		return "KNoProcessValue"
	case KSystemProcess:
		return "KSystemProcess"
	default:
		return fmt.Sprintf("KNoProcess(%d)", e)
	}
}

type KNoTransform uint

const (
	// KDeviceToPCS: Device Dependent to Device Independent
	KDeviceToPCS KNoTransform = 1
	// KNoTransformValue: Not used.
	KNoTransformValue KNoTransform = 0
	// KPCSToDevice: Device Independent to Device Dependent
	KPCSToDevice KNoTransform = 2
	// KPCSToPCS: Independent, through device's gamut
	KPCSToPCS KNoTransform = 3
	// KUseAtoB: Use 'A2B*' tag from this profile or equivalent
	KUseAtoB KNoTransform = 1
	// KUseBtoA: Use 'B2A*' tag from this profile or equivalent
	KUseBtoA KNoTransform = 2
	// KUseBtoB: Use 'pre*' tag from this profile or equivalent
	KUseBtoB KNoTransform = 3
	// KUseProfileIntent: For renderingIntent in NCMConcatProfileSpec
	KUseProfileIntent KNoTransform = 0xffffffff
)

func (e KNoTransform) String() string {
	switch e {
	case KDeviceToPCS:
		return "KDeviceToPCS"
	case KNoTransformValue:
		return "KNoTransformValue"
	case KPCSToDevice:
		return "KPCSToDevice"
	case KPCSToPCS:
		return "KPCSToPCS"
	case KUseProfileIntent:
		return "KUseProfileIntent"
	default:
		return fmt.Sprintf("KNoTransform(%d)", e)
	}
}

type KPM int

const (
	KPMAllocationFailure   KPM = -108
	KPMCMYKColorSpaceModel KPM = 3
	KPMCVMSymbolNotFound   KPM = -9662
	// KPMCancel: Specifies that the user clicked the Cancel button in a Print or Page Setup dialog.
	KPMCancel                KPM = 0x80
	KPMCloseFailed           KPM = -9785
	KPMCreateMessageFailed   KPM = -9620
	KPMDeleteSubTicketFailed KPM = -9585
	KPMDevNColorSpaceModel   KPM = 4
	KPMDocumentNotFound      KPM = -9644
	KPMDontSwitchPDEError    KPM = -9531
	// KPMDuplexNoTumble: # Discussion
	KPMDuplexNoTumble KPM = 0x2
	// KPMDuplexNone: # Discussion
	KPMDuplexNone KPM = 0x1
	// KPMDuplexTumble: # Discussion
	KPMDuplexTumble             KPM = 0x3
	KPMEditRequestFailed        KPM = -9544
	KPMFeatureNotInstalled      KPM = -9533
	KPMFileOrDirOperationFailed KPM = -9634
	KPMFontNameTooLong          KPM = -9704
	KPMFontNotFound             KPM = -9703
	KPMGeneralCGError           KPM = -9705
	// KPMGeneralError: An unspecified error occurred.
	KPMGeneralError        KPM = -30870
	KPMGrayColorSpaceModel KPM = 1
	KPMHideInlineItems     KPM = 0
	KPMIOAttrNotAvailable  KPM = -9787
	KPMIOMSymbolNotFound   KPM = -9661
	KPMInternalError       KPM = -30870
	// KPMInvalidAllocator: The specified memory allocator is invalid.
	KPMInvalidAllocator  KPM = -30890
	KPMInvalidCVMContext KPM = -9665
	// KPMInvalidCalibrationTarget: The dictionary specifying a printer calibration target is invalid.
	KPMInvalidCalibrationTarget KPM = -30898
	// KPMInvalidConnection: The printer connection type is invalid.
	KPMInvalidConnection KPM = -30887
	// KPMInvalidFileType: The file type is invalid.
	KPMInvalidFileType   KPM = -30895
	KPMInvalidIOMContext KPM = -9664
	// KPMInvalidIndex: An array index is invalid.
	KPMInvalidIndex KPM = -30882
	// KPMInvalidItem: The item being added to a ticket is invalid.
	KPMInvalidItem  KPM = -30892
	KPMInvalidJobID KPM = -9666
	// KPMInvalidJobTemplate: An internal error occurred while creating a job template.
	KPMInvalidJobTemplate KPM = -30885
	// KPMInvalidKey: The key in a ticket, job template, or dictionary is invalid.
	KPMInvalidKey        KPM = -30888
	KPMInvalidLookupSpec KPM = -9542
	// KPMInvalidObject: The object is invalid.
	KPMInvalidObject     KPM = -30896
	KPMInvalidPBMRef     KPM = -9540
	KPMInvalidPDEContext KPM = -9530
	KPMInvalidPMContext  KPM = -9663
	// KPMInvalidPageFormat: Your application passed an invalid page format object.
	KPMInvalidPageFormat KPM = -30876
	// KPMInvalidPaper: Your application passed an invalid paper object.
	KPMInvalidPaper     KPM = -30897
	KPMInvalidParameter KPM = -50
	// KPMInvalidPreset: Your application passed an invalid preset object.
	KPMInvalidPreset KPM = -30899
	// KPMInvalidPrintSession: Your application passed an invalid printing session object.
	KPMInvalidPrintSession KPM = -30879
	// KPMInvalidPrintSettings: Your application passed an invalid print settings object.
	KPMInvalidPrintSettings KPM = -30875
	// KPMInvalidPrinter: Your application passed an invalid printer object.
	KPMInvalidPrinter        KPM = -30880
	KPMInvalidPrinterAddress KPM = -9780
	// KPMInvalidPrinterInfo: The printer information is invalid.
	KPMInvalidPrinterInfo KPM = -30886
	// KPMInvalidReply: A remote server or client sent an invalid reply.
	KPMInvalidReply     KPM = -30894
	KPMInvalidState     KPM = -9706
	KPMInvalidSubTicket KPM = -9584
	// KPMInvalidTicket: The job ticket is invalid.
	KPMInvalidTicket KPM = -30891
	// KPMInvalidType: The data type in a ticket, job template, or dictionary is not the expected type.
	KPMInvalidType KPM = -30893
	// KPMInvalidValue: The value in a ticket, job template, or dictionary is missing.
	KPMInvalidValue               KPM = -30889
	KPMItemIsLocked               KPM = -9586
	KPMJobBusy                    KPM = -9642
	KPMJobCanceled                KPM = -9643
	KPMJobGetTicketBadFormatError KPM = -9672
	KPMJobGetTicketReadError      KPM = -9673
	KPMJobManagerAborted          KPM = -9671
	KPMJobNotFound                KPM = -9641
	KPMJobStreamEndError          KPM = -9670
	KPMJobStreamOpenFailed        KPM = -9668
	KPMJobStreamReadFailed        KPM = -9669
	KPMKeyNotFound                KPM = -9589
	KPMKeyNotUnique               KPM = -9590
	KPMKeyOrValueNotFound         KPM = -9623
	// KPMLandscape: # Discussion
	KPMLandscape                                      KPM = 2
	KPMLastErrorCodeToMakeMaintenanceOfThisListEasier KPM = -9799
	KPMMessagingError                                 KPM = -9624
	KPMNoDefaultItem                                  KPM = -9500
	// KPMNoDefaultPrinter: The user has not specified a default printer.
	KPMNoDefaultPrinter   KPM = -30872
	KPMNoDefaultSettings  KPM = -9501
	KPMNoError            KPM = 0
	KPMNoPrinterJobID     KPM = -9667
	KPMNoSelectedPrinters KPM = -9541
	// KPMNoSuchEntry: There is no entry to match your application’s request.
	KPMNoSuchEntry KPM = -30874
	// KPMNotImplemented: The function is not implemented.
	KPMNotImplemented KPM = -30873
	// KPMObjectInUse: The specified object is in use.
	KPMObjectInUse KPM = -30881
	KPMOpenFailed  KPM = -9781
	// KPMOutOfScope: Your application called this function out of sequence with other printing functions.
	KPMOutOfScope                KPM = -30871
	KPMPMSymbolNotFound          KPM = -9660
	KPMPermissionError           KPM = -9636
	KPMPluginNotFound            KPM = -9701
	KPMPluginRegisterationFailed KPM = -9702
	// KPMPortrait: Specifies portrait (vertical) page orientation.
	KPMPortrait           KPM = 1
	KPMPrBrowserNoUI      KPM = -9545
	KPMQueueAlreadyExists KPM = -9639
	KPMQueueJobFailed     KPM = -9640
	KPMQueueNotFound      KPM = -9638
	KPMRGBColorSpaceModel KPM = 2
	KPMReadFailed         KPM = -9782
	KPMReadGotZeroData    KPM = -9788
	// KPMReverseLandscape: # Discussion
	KPMReverseLandscape KPM = 4
	// KPMReversePortrait: # Discussion
	KPMReversePortrait                  KPM = 3
	KPMServerAlreadyRunning             KPM = -9631
	KPMServerAttributeRestricted        KPM = -9633
	KPMServerCommunicationFailed        KPM = -9621
	KPMServerNotFound                   KPM = -9630
	KPMServerSuspended                  KPM = -9632
	KPMShowDefaultInlineItems           KPM = 32768
	KPMShowInlineCopies                 KPM = 1
	KPMShowInlineOrientation            KPM = 8
	KPMShowInlinePageRange              KPM = 2
	KPMShowInlinePageRangeWithSelection KPM = 64
	KPMShowInlinePaperSize              KPM = 4
	KPMShowInlineScale                  KPM = 128
	KPMShowPageAttributesPDE            KPM = 256
	// KPMSimplexTumble: # Discussion
	KPMSimplexTumble KPM = 0x4
	KPMStatusFailed  KPM = -9784
	// KPMStringConversionFailure: An internal error occurred while converting a string.
	KPMStringConversionFailure KPM = -30883
	KPMSubTicketNotFound       KPM = -9583
	KPMSyncRequestFailed       KPM = -9543
	KPMTemplateIsLocked        KPM = -9588
	KPMTicketIsLocked          KPM = -9587
	KPMTicketTypeNotFound      KPM = -9580
	KPMUnableToFindProcess     KPM = -9532
	KPMUnexpectedImagingError  KPM = -9707
	KPMUnknownColorSpaceModel  KPM = 0
	KPMUnknownDataType         KPM = -9591
	KPMUnknownMessage          KPM = -9637
	KPMUnlocked                KPM = 0
	KPMUnsupportedConnection   KPM = -9786
	KPMUpdateTicketFailed      KPM = -9581
	KPMUserOrGroupNotFound     KPM = -9635
	KPMValidateTicketFailed    KPM = -9582
	// KPMValueOutOfRange: Your application passed an out-of-range value.
	KPMValueOutOfRange KPM = -30877
	KPMWriteFailed     KPM = -9783
	// KPMXMLParseError: An error occurred while parsing XML data.
	KPMXMLParseError KPM = -30884
)

func (e KPM) String() string {
	switch e {
	case KPMAllocationFailure:
		return "KPMAllocationFailure"
	case KPMCMYKColorSpaceModel:
		return "KPMCMYKColorSpaceModel"
	case KPMCVMSymbolNotFound:
		return "KPMCVMSymbolNotFound"
	case KPMCancel:
		return "KPMCancel"
	case KPMCloseFailed:
		return "KPMCloseFailed"
	case KPMCreateMessageFailed:
		return "KPMCreateMessageFailed"
	case KPMDeleteSubTicketFailed:
		return "KPMDeleteSubTicketFailed"
	case KPMDevNColorSpaceModel:
		return "KPMDevNColorSpaceModel"
	case KPMDocumentNotFound:
		return "KPMDocumentNotFound"
	case KPMDontSwitchPDEError:
		return "KPMDontSwitchPDEError"
	case KPMDuplexNoTumble:
		return "KPMDuplexNoTumble"
	case KPMDuplexNone:
		return "KPMDuplexNone"
	case KPMEditRequestFailed:
		return "KPMEditRequestFailed"
	case KPMFeatureNotInstalled:
		return "KPMFeatureNotInstalled"
	case KPMFileOrDirOperationFailed:
		return "KPMFileOrDirOperationFailed"
	case KPMFontNameTooLong:
		return "KPMFontNameTooLong"
	case KPMFontNotFound:
		return "KPMFontNotFound"
	case KPMGeneralCGError:
		return "KPMGeneralCGError"
	case KPMGeneralError:
		return "KPMGeneralError"
	case KPMHideInlineItems:
		return "KPMHideInlineItems"
	case KPMIOAttrNotAvailable:
		return "KPMIOAttrNotAvailable"
	case KPMIOMSymbolNotFound:
		return "KPMIOMSymbolNotFound"
	case KPMInvalidAllocator:
		return "KPMInvalidAllocator"
	case KPMInvalidCVMContext:
		return "KPMInvalidCVMContext"
	case KPMInvalidCalibrationTarget:
		return "KPMInvalidCalibrationTarget"
	case KPMInvalidConnection:
		return "KPMInvalidConnection"
	case KPMInvalidFileType:
		return "KPMInvalidFileType"
	case KPMInvalidIOMContext:
		return "KPMInvalidIOMContext"
	case KPMInvalidIndex:
		return "KPMInvalidIndex"
	case KPMInvalidItem:
		return "KPMInvalidItem"
	case KPMInvalidJobID:
		return "KPMInvalidJobID"
	case KPMInvalidJobTemplate:
		return "KPMInvalidJobTemplate"
	case KPMInvalidKey:
		return "KPMInvalidKey"
	case KPMInvalidLookupSpec:
		return "KPMInvalidLookupSpec"
	case KPMInvalidObject:
		return "KPMInvalidObject"
	case KPMInvalidPBMRef:
		return "KPMInvalidPBMRef"
	case KPMInvalidPDEContext:
		return "KPMInvalidPDEContext"
	case KPMInvalidPMContext:
		return "KPMInvalidPMContext"
	case KPMInvalidPageFormat:
		return "KPMInvalidPageFormat"
	case KPMInvalidPaper:
		return "KPMInvalidPaper"
	case KPMInvalidParameter:
		return "KPMInvalidParameter"
	case KPMInvalidPreset:
		return "KPMInvalidPreset"
	case KPMInvalidPrintSession:
		return "KPMInvalidPrintSession"
	case KPMInvalidPrintSettings:
		return "KPMInvalidPrintSettings"
	case KPMInvalidPrinter:
		return "KPMInvalidPrinter"
	case KPMInvalidPrinterAddress:
		return "KPMInvalidPrinterAddress"
	case KPMInvalidPrinterInfo:
		return "KPMInvalidPrinterInfo"
	case KPMInvalidReply:
		return "KPMInvalidReply"
	case KPMInvalidState:
		return "KPMInvalidState"
	case KPMInvalidSubTicket:
		return "KPMInvalidSubTicket"
	case KPMInvalidTicket:
		return "KPMInvalidTicket"
	case KPMInvalidType:
		return "KPMInvalidType"
	case KPMInvalidValue:
		return "KPMInvalidValue"
	case KPMItemIsLocked:
		return "KPMItemIsLocked"
	case KPMJobBusy:
		return "KPMJobBusy"
	case KPMJobCanceled:
		return "KPMJobCanceled"
	case KPMJobGetTicketBadFormatError:
		return "KPMJobGetTicketBadFormatError"
	case KPMJobGetTicketReadError:
		return "KPMJobGetTicketReadError"
	case KPMJobManagerAborted:
		return "KPMJobManagerAborted"
	case KPMJobNotFound:
		return "KPMJobNotFound"
	case KPMJobStreamEndError:
		return "KPMJobStreamEndError"
	case KPMJobStreamOpenFailed:
		return "KPMJobStreamOpenFailed"
	case KPMJobStreamReadFailed:
		return "KPMJobStreamReadFailed"
	case KPMKeyNotFound:
		return "KPMKeyNotFound"
	case KPMKeyNotUnique:
		return "KPMKeyNotUnique"
	case KPMKeyOrValueNotFound:
		return "KPMKeyOrValueNotFound"
	case KPMLastErrorCodeToMakeMaintenanceOfThisListEasier:
		return "KPMLastErrorCodeToMakeMaintenanceOfThisListEasier"
	case KPMMessagingError:
		return "KPMMessagingError"
	case KPMNoDefaultItem:
		return "KPMNoDefaultItem"
	case KPMNoDefaultPrinter:
		return "KPMNoDefaultPrinter"
	case KPMNoDefaultSettings:
		return "KPMNoDefaultSettings"
	case KPMNoPrinterJobID:
		return "KPMNoPrinterJobID"
	case KPMNoSelectedPrinters:
		return "KPMNoSelectedPrinters"
	case KPMNoSuchEntry:
		return "KPMNoSuchEntry"
	case KPMNotImplemented:
		return "KPMNotImplemented"
	case KPMObjectInUse:
		return "KPMObjectInUse"
	case KPMOpenFailed:
		return "KPMOpenFailed"
	case KPMOutOfScope:
		return "KPMOutOfScope"
	case KPMPMSymbolNotFound:
		return "KPMPMSymbolNotFound"
	case KPMPermissionError:
		return "KPMPermissionError"
	case KPMPluginNotFound:
		return "KPMPluginNotFound"
	case KPMPluginRegisterationFailed:
		return "KPMPluginRegisterationFailed"
	case KPMPrBrowserNoUI:
		return "KPMPrBrowserNoUI"
	case KPMQueueAlreadyExists:
		return "KPMQueueAlreadyExists"
	case KPMQueueJobFailed:
		return "KPMQueueJobFailed"
	case KPMQueueNotFound:
		return "KPMQueueNotFound"
	case KPMReadFailed:
		return "KPMReadFailed"
	case KPMReadGotZeroData:
		return "KPMReadGotZeroData"
	case KPMServerAlreadyRunning:
		return "KPMServerAlreadyRunning"
	case KPMServerAttributeRestricted:
		return "KPMServerAttributeRestricted"
	case KPMServerCommunicationFailed:
		return "KPMServerCommunicationFailed"
	case KPMServerNotFound:
		return "KPMServerNotFound"
	case KPMServerSuspended:
		return "KPMServerSuspended"
	case KPMShowDefaultInlineItems:
		return "KPMShowDefaultInlineItems"
	case KPMShowInlineOrientation:
		return "KPMShowInlineOrientation"
	case KPMShowInlinePageRangeWithSelection:
		return "KPMShowInlinePageRangeWithSelection"
	case KPMShowPageAttributesPDE:
		return "KPMShowPageAttributesPDE"
	case KPMStatusFailed:
		return "KPMStatusFailed"
	case KPMStringConversionFailure:
		return "KPMStringConversionFailure"
	case KPMSubTicketNotFound:
		return "KPMSubTicketNotFound"
	case KPMSyncRequestFailed:
		return "KPMSyncRequestFailed"
	case KPMTemplateIsLocked:
		return "KPMTemplateIsLocked"
	case KPMTicketIsLocked:
		return "KPMTicketIsLocked"
	case KPMTicketTypeNotFound:
		return "KPMTicketTypeNotFound"
	case KPMUnableToFindProcess:
		return "KPMUnableToFindProcess"
	case KPMUnexpectedImagingError:
		return "KPMUnexpectedImagingError"
	case KPMUnknownDataType:
		return "KPMUnknownDataType"
	case KPMUnknownMessage:
		return "KPMUnknownMessage"
	case KPMUnsupportedConnection:
		return "KPMUnsupportedConnection"
	case KPMUpdateTicketFailed:
		return "KPMUpdateTicketFailed"
	case KPMUserOrGroupNotFound:
		return "KPMUserOrGroupNotFound"
	case KPMValidateTicketFailed:
		return "KPMValidateTicketFailed"
	case KPMValueOutOfRange:
		return "KPMValueOutOfRange"
	case KPMWriteFailed:
		return "KPMWriteFailed"
	case KPMXMLParseError:
		return "KPMXMLParseError"
	default:
		return fmt.Sprintf("KPM(%d)", e)
	}
}

type KPMBorder uint

const (
	KPMBorderDoubleHairline  KPMBorder = 2
	KPMBorderDoubleThickline KPMBorder = 4
	KPMBorderSingleHairline  KPMBorder = 1
	KPMBorderSingleThickline KPMBorder = 3
)

func (e KPMBorder) String() string {
	switch e {
	case KPMBorderDoubleHairline:
		return "KPMBorderDoubleHairline"
	case KPMBorderDoubleThickline:
		return "KPMBorderDoubleThickline"
	case KPMBorderSingleHairline:
		return "KPMBorderSingleHairline"
	case KPMBorderSingleThickline:
		return "KPMBorderSingleThickline"
	default:
		return fmt.Sprintf("KPMBorder(%d)", e)
	}
}

type KPMCoverPage uint

const (
	KPMCoverPageAfter  KPMCoverPage = 3
	KPMCoverPageBefore KPMCoverPage = 2
	KPMCoverPageNone   KPMCoverPage = 1
)

func (e KPMCoverPage) String() string {
	switch e {
	case KPMCoverPageAfter:
		return "KPMCoverPageAfter"
	case KPMCoverPageBefore:
		return "KPMCoverPageBefore"
	case KPMCoverPageNone:
		return "KPMCoverPageNone"
	default:
		return fmt.Sprintf("KPMCoverPage(%d)", e)
	}
}

type KPMDestination uint

const (
	// KPMDestinationFax: Specifies output to a fax.
	KPMDestinationFax KPMDestination = 3
	// KPMDestinationFile: Specifies output to a file.
	KPMDestinationFile KPMDestination = 2
	// KPMDestinationInvalid: Specifies the destination is invalid.
	KPMDestinationInvalid KPMDestination = 0
	// KPMDestinationPreview: Specifies output to print preview.
	KPMDestinationPreview KPMDestination = 4
	// KPMDestinationPrinter: Specifies output to a printer.
	KPMDestinationPrinter KPMDestination = 1
	// KPMDestinationProcessPDF: Specifies output to a PDF workflow option.
	KPMDestinationProcessPDF KPMDestination = 5
)

func (e KPMDestination) String() string {
	switch e {
	case KPMDestinationFax:
		return "KPMDestinationFax"
	case KPMDestinationFile:
		return "KPMDestinationFile"
	case KPMDestinationInvalid:
		return "KPMDestinationInvalid"
	case KPMDestinationPreview:
		return "KPMDestinationPreview"
	case KPMDestinationPrinter:
		return "KPMDestinationPrinter"
	case KPMDestinationProcessPDF:
		return "KPMDestinationProcessPDF"
	default:
		return fmt.Sprintf("KPMDestination(%d)", e)
	}
}

type KPMLayout uint

const (
	KPMLayoutBottomTopLeftRight KPMLayout = 7
	KPMLayoutBottomTopRightLeft KPMLayout = 8
	KPMLayoutLeftRightBottomTop KPMLayout = 2
	KPMLayoutLeftRightTopBottom KPMLayout = 1
	KPMLayoutRightLeftBottomTop KPMLayout = 4
	KPMLayoutRightLeftTopBottom KPMLayout = 3
	KPMLayoutTopBottomLeftRight KPMLayout = 5
	KPMLayoutTopBottomRightLeft KPMLayout = 6
)

func (e KPMLayout) String() string {
	switch e {
	case KPMLayoutBottomTopLeftRight:
		return "KPMLayoutBottomTopLeftRight"
	case KPMLayoutBottomTopRightLeft:
		return "KPMLayoutBottomTopRightLeft"
	case KPMLayoutLeftRightBottomTop:
		return "KPMLayoutLeftRightBottomTop"
	case KPMLayoutLeftRightTopBottom:
		return "KPMLayoutLeftRightTopBottom"
	case KPMLayoutRightLeftBottomTop:
		return "KPMLayoutRightLeftBottomTop"
	case KPMLayoutRightLeftTopBottom:
		return "KPMLayoutRightLeftTopBottom"
	case KPMLayoutTopBottomLeftRight:
		return "KPMLayoutTopBottomLeftRight"
	case KPMLayoutTopBottomRightLeft:
		return "KPMLayoutTopBottomRightLeft"
	default:
		return fmt.Sprintf("KPMLayout(%d)", e)
	}
}

type KPMPaperType uint

const (
	KPMPaperTypeCoated       KPMPaperType = 0x2
	KPMPaperTypeGlossy       KPMPaperType = 0x4
	KPMPaperTypePlain        KPMPaperType = 0x1
	KPMPaperTypePremium      KPMPaperType = 0x3
	KPMPaperTypeTShirt       KPMPaperType = 0x6
	KPMPaperTypeTransparency KPMPaperType = 0x5
	KPMPaperTypeUnknown      KPMPaperType = 0
)

func (e KPMPaperType) String() string {
	switch e {
	case KPMPaperTypeCoated:
		return "KPMPaperTypeCoated"
	case KPMPaperTypeGlossy:
		return "KPMPaperTypeGlossy"
	case KPMPaperTypePlain:
		return "KPMPaperTypePlain"
	case KPMPaperTypePremium:
		return "KPMPaperTypePremium"
	case KPMPaperTypeTShirt:
		return "KPMPaperTypeTShirt"
	case KPMPaperTypeTransparency:
		return "KPMPaperTypeTransparency"
	case KPMPaperTypeUnknown:
		return "KPMPaperTypeUnknown"
	default:
		return fmt.Sprintf("KPMPaperType(%d)", e)
	}
}

type KPMPrintAll int

const (
	// KPMPrintAllPages: # Discussion
	KPMPrintAllPages KPMPrintAll = -1
)

func (e KPMPrintAll) String() string {
	switch e {
	case KPMPrintAllPages:
		return "KPMPrintAllPages"
	default:
		return fmt.Sprintf("KPMPrintAll(%d)", e)
	}
}

type KPMPrinter uint

const (
	// KPMPrinterIdle: Specifies the idle state.
	KPMPrinterIdle KPMPrinter = 3
	// KPMPrinterProcessing: Specifies the processing state.
	KPMPrinterProcessing KPMPrinter = 4
	// KPMPrinterStopped: Specifies the stopped state.
	KPMPrinterStopped KPMPrinter = 5
)

func (e KPMPrinter) String() string {
	switch e {
	case KPMPrinterIdle:
		return "KPMPrinterIdle"
	case KPMPrinterProcessing:
		return "KPMPrinterProcessing"
	case KPMPrinterStopped:
		return "KPMPrinterStopped"
	default:
		return fmt.Sprintf("KPMPrinter(%d)", e)
	}
}

type KPMQuality uint

const (
	// KPMQualityBest: Specifies to get the best print quality for all objects and photos on a page.
	KPMQualityBest KPMQuality = 0xd
	// KPMQualityDraft: Specifies to print at the highest speed, with the amount of ink used as a secondary consideration.
	KPMQualityDraft KPMQuality = 0x4
	// KPMQualityHighest: Specifies to use the highest print quality available to the printer.
	KPMQualityHighest KPMQuality = 0xf
	// KPMQualityInkSaver: Specifies to use a mode that saves ink, even if it slows printing.
	KPMQualityInkSaver KPMQuality = 0x1
	// KPMQualityLowest: Specifies to use the lowest print quality available to the printer.
	KPMQualityLowest KPMQuality = 0
	// KPMQualityNormal: Specifies a general usage mode that balances quality and speed.
	KPMQualityNormal KPMQuality = 0x8
	// KPMQualityPhoto: Specifies to optimize the quality of photos on the page, with speed not a concern.
	KPMQualityPhoto KPMQuality = 0xb
)

func (e KPMQuality) String() string {
	switch e {
	case KPMQualityBest:
		return "KPMQualityBest"
	case KPMQualityDraft:
		return "KPMQualityDraft"
	case KPMQualityHighest:
		return "KPMQualityHighest"
	case KPMQualityInkSaver:
		return "KPMQualityInkSaver"
	case KPMQualityLowest:
		return "KPMQualityLowest"
	case KPMQualityNormal:
		return "KPMQualityNormal"
	case KPMQualityPhoto:
		return "KPMQualityPhoto"
	default:
		return fmt.Sprintf("KPMQuality(%d)", e)
	}
}

type KPMScaling uint

const (
	KPMScalingCenterOnImgArea KPMScaling = 6
	KPMScalingCenterOnPaper   KPMScaling = 5
	KPMScalingPinBottomLeft   KPMScaling = 3
	KPMScalingPinBottomRight  KPMScaling = 4
	KPMScalingPinTopLeft      KPMScaling = 1
	KPMScalingPinTopRight     KPMScaling = 2
)

func (e KPMScaling) String() string {
	switch e {
	case KPMScalingCenterOnImgArea:
		return "KPMScalingCenterOnImgArea"
	case KPMScalingCenterOnPaper:
		return "KPMScalingCenterOnPaper"
	case KPMScalingPinBottomLeft:
		return "KPMScalingPinBottomLeft"
	case KPMScalingPinBottomRight:
		return "KPMScalingPinBottomRight"
	case KPMScalingPinTopLeft:
		return "KPMScalingPinTopLeft"
	case KPMScalingPinTopRight:
		return "KPMScalingPinTopRight"
	default:
		return fmt.Sprintf("KPMScaling(%d)", e)
	}
}

type KPlotIconRef uint

const (
	KPlotIconRefNoImage     KPlotIconRef = 2
	KPlotIconRefNoMask      KPlotIconRef = 4
	KPlotIconRefNormalFlags KPlotIconRef = 0
)

func (e KPlotIconRef) String() string {
	switch e {
	case KPlotIconRefNoImage:
		return "KPlotIconRefNoImage"
	case KPlotIconRefNoMask:
		return "KPlotIconRefNoMask"
	case KPlotIconRefNormalFlags:
		return "KPlotIconRefNormalFlags"
	default:
		return fmt.Sprintf("KPlotIconRef(%d)", e)
	}
}

type KProcessDictionaryIncludeAllInformation uint

const (
	KProcessDictionaryIncludeAllInformationMask KProcessDictionaryIncludeAllInformation = 0xffffffff
)

func (e KProcessDictionaryIncludeAllInformation) String() string {
	switch e {
	case KProcessDictionaryIncludeAllInformationMask:
		return "KProcessDictionaryIncludeAllInformationMask"
	default:
		return fmt.Sprintf("KProcessDictionaryIncludeAllInformation(%d)", e)
	}
}

type KProcessTransformTo uint

const (
	KProcessTransformToBackgroundApplication KProcessTransformTo = 2
	KProcessTransformToForegroundApplication KProcessTransformTo = 1
	KProcessTransformToUIElementApplication  KProcessTransformTo = 4
)

func (e KProcessTransformTo) String() string {
	switch e {
	case KProcessTransformToBackgroundApplication:
		return "KProcessTransformToBackgroundApplication"
	case KProcessTransformToForegroundApplication:
		return "KProcessTransformToForegroundApplication"
	case KProcessTransformToUIElementApplication:
		return "KProcessTransformToUIElementApplication"
	default:
		return fmt.Sprintf("KProcessTransformTo(%d)", e)
	}
}

type KQuit uint

const (
	KQuitAtNormalTimeMask             KQuit = 2
	KQuitBeforeFBAsQuitMask           KQuit = 4
	KQuitBeforeNormalTimeMask         KQuit = 1
	KQuitBeforeShellQuitsMask         KQuit = 8
	KQuitBeforeTerminatorAppQuitsMask KQuit = 16
	KQuitNeverMask                    KQuit = 32
	KQuitNotQuitDuringInstallMask     KQuit = 0x100
	KQuitNotQuitDuringLogoutMask      KQuit = 0x200
	KQuitOptionsMask                  KQuit = 0x7f
)

func (e KQuit) String() string {
	switch e {
	case KQuitAtNormalTimeMask:
		return "KQuitAtNormalTimeMask"
	case KQuitBeforeFBAsQuitMask:
		return "KQuitBeforeFBAsQuitMask"
	case KQuitBeforeNormalTimeMask:
		return "KQuitBeforeNormalTimeMask"
	case KQuitBeforeShellQuitsMask:
		return "KQuitBeforeShellQuitsMask"
	case KQuitBeforeTerminatorAppQuitsMask:
		return "KQuitBeforeTerminatorAppQuitsMask"
	case KQuitNeverMask:
		return "KQuitNeverMask"
	case KQuitNotQuitDuringInstallMask:
		return "KQuitNotQuitDuringInstallMask"
	case KQuitNotQuitDuringLogoutMask:
		return "KQuitNotQuitDuringLogoutMask"
	case KQuitOptionsMask:
		return "KQuitOptionsMask"
	default:
		return fmt.Sprintf("KQuit(%d)", e)
	}
}

type KSelector uint

const (
	KSelectorAll1BitData      KSelector = 1
	KSelectorAll32BitData     KSelector = 8
	KSelectorAll4BitData      KSelector = 2
	KSelectorAll8BitData      KSelector = 4
	KSelectorAllAvailableData KSelector = 0xffffffff
	KSelectorAllHugeData      KSelector = 0xff000000
	KSelectorAllLargeData     KSelector = 0xff
	KSelectorAllMiniData      KSelector = 0xff0000
	KSelectorAllSmallData     KSelector = 0xff00
	KSelectorHuge1Bit         KSelector = 0x1000000
	KSelectorHuge32Bit        KSelector = 0x8000000
	KSelectorHuge4Bit         KSelector = 0x2000000
	KSelectorHuge8Bit         KSelector = 0x4000000
	KSelectorHuge8BitMask     KSelector = 0x10000000
	KSelectorLarge1Bit        KSelector = 0x1
	KSelectorLarge32Bit       KSelector = 0x8
	KSelectorLarge4Bit        KSelector = 0x2
	KSelectorLarge8Bit        KSelector = 0x4
	KSelectorLarge8BitMask    KSelector = 0x10
	KSelectorMini1Bit         KSelector = 0x10000
	KSelectorMini4Bit         KSelector = 0x20000
	KSelectorMini8Bit         KSelector = 0x40000
	KSelectorSmall1Bit        KSelector = 0x100
	KSelectorSmall32Bit       KSelector = 0x800
	KSelectorSmall4Bit        KSelector = 0x200
	KSelectorSmall8Bit        KSelector = 0x400
	KSelectorSmall8BitMask    KSelector = 0x1000
)

func (e KSelector) String() string {
	switch e {
	case KSelectorAll1BitData:
		return "KSelectorAll1BitData"
	case KSelectorAll32BitData:
		return "KSelectorAll32BitData"
	case KSelectorAll4BitData:
		return "KSelectorAll4BitData"
	case KSelectorAll8BitData:
		return "KSelectorAll8BitData"
	case KSelectorAllAvailableData:
		return "KSelectorAllAvailableData"
	case KSelectorAllHugeData:
		return "KSelectorAllHugeData"
	case KSelectorAllLargeData:
		return "KSelectorAllLargeData"
	case KSelectorAllMiniData:
		return "KSelectorAllMiniData"
	case KSelectorAllSmallData:
		return "KSelectorAllSmallData"
	case KSelectorHuge1Bit:
		return "KSelectorHuge1Bit"
	case KSelectorHuge32Bit:
		return "KSelectorHuge32Bit"
	case KSelectorHuge4Bit:
		return "KSelectorHuge4Bit"
	case KSelectorHuge8Bit:
		return "KSelectorHuge8Bit"
	case KSelectorHuge8BitMask:
		return "KSelectorHuge8BitMask"
	case KSelectorLarge8BitMask:
		return "KSelectorLarge8BitMask"
	case KSelectorMini1Bit:
		return "KSelectorMini1Bit"
	case KSelectorMini4Bit:
		return "KSelectorMini4Bit"
	case KSelectorMini8Bit:
		return "KSelectorMini8Bit"
	case KSelectorSmall1Bit:
		return "KSelectorSmall1Bit"
	case KSelectorSmall32Bit:
		return "KSelectorSmall32Bit"
	case KSelectorSmall4Bit:
		return "KSelectorSmall4Bit"
	case KSelectorSmall8Bit:
		return "KSelectorSmall8Bit"
	case KSelectorSmall8BitMask:
		return "KSelectorSmall8BitMask"
	default:
		return fmt.Sprintf("KSelector(%d)", e)
	}
}

type KSetFrontProcess uint

const (
	KSetFrontProcessCausedByUser    KSetFrontProcess = 2
	KSetFrontProcessFrontWindowOnly KSetFrontProcess = 1
)

func (e KSetFrontProcess) String() string {
	switch e {
	case KSetFrontProcessCausedByUser:
		return "KSetFrontProcessCausedByUser"
	case KSetFrontProcessFrontWindowOnly:
		return "KSetFrontProcessFrontWindowOnly"
	default:
		return fmt.Sprintf("KSetFrontProcess(%d)", e)
	}
}

type KSpeech uint

const (
	// KSpeechGenerateTune: # Discussion
	KSpeechGenerateTune KSpeech = 1
	// KSpeechRelativeDuration: # Discussion
	KSpeechRelativeDuration KSpeech = 4
	// KSpeechRelativePitch: # Discussion
	KSpeechRelativePitch KSpeech = 2
	// KSpeechShowSyllables: # Discussion
	KSpeechShowSyllables KSpeech = 8
)

func (e KSpeech) String() string {
	switch e {
	case KSpeechGenerateTune:
		return "KSpeechGenerateTune"
	case KSpeechRelativeDuration:
		return "KSpeechRelativeDuration"
	case KSpeechRelativePitch:
		return "KSpeechRelativePitch"
	case KSpeechShowSyllables:
		return "KSpeechShowSyllables"
	default:
		return fmt.Sprintf("KSpeech(%d)", e)
	}
}

type KTextToSpeech uint

const (
	// KTextToSpeechSynthType: The type of a synthesizer component.
	KTextToSpeechSynthType KTextToSpeech = 't'<<24 | 't'<<16 | 's'<<8 | 'c' // 'ttsc'
	// KTextToSpeechVoiceBundleType: The type of a voice bundle file.
	KTextToSpeechVoiceBundleType KTextToSpeech = 't'<<24 | 't'<<16 | 'v'<<8 | 'b' // 'ttvb'
	// KTextToSpeechVoiceFileType: The type of a voice file.
	KTextToSpeechVoiceFileType KTextToSpeech = 't'<<24 | 't'<<16 | 'v'<<8 | 'f' // 'ttvf'
	// KTextToSpeechVoiceType: The type of a voice resource.
	KTextToSpeechVoiceType KTextToSpeech = 't'<<24 | 't'<<16 | 'v'<<8 | 'd' // 'ttvd'
)

func (e KTextToSpeech) String() string {
	switch e {
	case KTextToSpeechSynthType:
		return "KTextToSpeechSynthType"
	case KTextToSpeechVoiceBundleType:
		return "KTextToSpeechVoiceBundleType"
	case KTextToSpeechVoiceFileType:
		return "KTextToSpeechVoiceFileType"
	case KTextToSpeechVoiceType:
		return "KTextToSpeechVoiceType"
	default:
		return fmt.Sprintf("KTextToSpeech(%d)", e)
	}
}

type KTransform uint

const (
	KTransformDisabled         KTransform = 0x1
	KTransformLabel1           KTransform = 0x100
	KTransformLabel2           KTransform = 0x200
	KTransformLabel3           KTransform = 0x300
	KTransformLabel4           KTransform = 0x400
	KTransformLabel5           KTransform = 0x500
	KTransformLabel6           KTransform = 0x600
	KTransformLabel7           KTransform = 0x700
	KTransformNone             KTransform = 0
	KTransformOffline          KTransform = 0x2
	KTransformOpen             KTransform = 0x3
	KTransformSelected         KTransform = 0x4000
	KTransformSelectedDisabled KTransform = 16384
	KTransformSelectedOffline  KTransform = 16384
	KTransformSelectedOpen     KTransform = 16384
)

func (e KTransform) String() string {
	switch e {
	case KTransformDisabled:
		return "KTransformDisabled"
	case KTransformLabel1:
		return "KTransformLabel1"
	case KTransformLabel2:
		return "KTransformLabel2"
	case KTransformLabel3:
		return "KTransformLabel3"
	case KTransformLabel4:
		return "KTransformLabel4"
	case KTransformLabel5:
		return "KTransformLabel5"
	case KTransformLabel6:
		return "KTransformLabel6"
	case KTransformLabel7:
		return "KTransformLabel7"
	case KTransformNone:
		return "KTransformNone"
	case KTransformOffline:
		return "KTransformOffline"
	case KTransformOpen:
		return "KTransformOpen"
	case KTransformSelected:
		return "KTransformSelected"
	default:
		return fmt.Sprintf("KTransform(%d)", e)
	}
}

type KTranslation uint

const (
	KTranslationDataTranslation KTranslation = 1
	KTranslationFileTranslation KTranslation = 2
)

func (e KTranslation) String() string {
	switch e {
	case KTranslationDataTranslation:
		return "KTranslationDataTranslation"
	case KTranslationFileTranslation:
		return "KTranslationFileTranslation"
	default:
		return fmt.Sprintf("KTranslation(%d)", e)
	}
}

type KUAZoomFocusType uint

const (
	// KUAZoomFocusTypeInsertionPoint: # Discussion
	KUAZoomFocusTypeInsertionPoint KUAZoomFocusType = 1
	// KUAZoomFocusTypeOther: # Discussion
	KUAZoomFocusTypeOther KUAZoomFocusType = 0
)

func (e KUAZoomFocusType) String() string {
	switch e {
	case KUAZoomFocusTypeInsertionPoint:
		return "KUAZoomFocusTypeInsertionPoint"
	case KUAZoomFocusTypeOther:
		return "KUAZoomFocusTypeOther"
	default:
		return fmt.Sprintf("KUAZoomFocusType(%d)", e)
	}
}

type Launch uint

const (
	LaunchAllow24Bit    Launch = 0x100
	LaunchContinue      Launch = 0x4000
	LaunchDontSwitch    Launch = 0x200
	LaunchInhibitDaemon Launch = 0x80
	LaunchNoFileFlags   Launch = 0x800
	LaunchUseMinimum    Launch = 0x400
)

func (e Launch) String() string {
	switch e {
	case LaunchAllow24Bit:
		return "LaunchAllow24Bit"
	case LaunchContinue:
		return "LaunchContinue"
	case LaunchDontSwitch:
		return "LaunchDontSwitch"
	case LaunchInhibitDaemon:
		return "LaunchInhibitDaemon"
	case LaunchNoFileFlags:
		return "LaunchNoFileFlags"
	case LaunchUseMinimum:
		return "LaunchUseMinimum"
	default:
		return fmt.Sprintf("Launch(%d)", e)
	}
}

type Mode uint

const (
	Mode32BitCompatible        Mode = 0x80
	ModeCanBackground          Mode = 0x1000
	ModeControlPanel           Mode = 0x80000
	ModeDeskAccessory          Mode = 0x20000
	ModeDisplayManagerAware    Mode = 0x4
	ModeDoesActivateOnFGSwitch Mode = 0x800
	ModeGetAppDiedMsg          Mode = 0x100
	ModeGetFrontClicks         Mode = 0x200
	ModeHighLevelEventAware    Mode = 0x40
	ModeLaunchDontSwitch       Mode = 0x40000
	// ModeLiteral: When the speech channel is in text-processing mode, indicates that characters and digits are spoken literally (for example, “cat” is spoken as “C-A-T” and “12” is spoken as "one, two").
	ModeLiteral                Mode = 'L'<<24 | 'T'<<16 | 'R'<<8 | 'L' // 'LTRL'
	ModeLocalAndRemoteHLEvents Mode = 0x20
	ModeMultiLaunch            Mode = 0x10000
	ModeNeedSuspendResume      Mode = 0x4000
	// ModeNormal: When the speech channel is in text-processing mode, indicates that the synthesizer should process characters as expected and assemble digits into numbers.
	ModeNormal         Mode = 'N'<<24 | 'O'<<16 | 'R'<<8 | 'M' // 'NORM'
	ModeOnlyBackground Mode = 0x400
	// ModePhonemes: Used with soInputMode to indicate that the speech channel is in phoneme-processing mode.
	ModePhonemes        Mode = 'P'<<24 | 'H'<<16 | 'O'<<8 | 'N' // 'PHON'
	ModeReserved        Mode = 0x1000000
	ModeStationeryAware Mode = 0x10
	// ModeText: Used with soInputMode to indicate that the speech channel is in text-processing mode.
	ModeText                Mode = 'T'<<24 | 'E'<<16 | 'X'<<8 | 'T' // 'TEXT'
	ModeTune                Mode = 'T'<<24 | 'U'<<16 | 'N'<<8 | 'E' // 'TUNE'
	ModeUseTextEditServices Mode = 0x8
)

func (e Mode) String() string {
	switch e {
	case Mode32BitCompatible:
		return "Mode32BitCompatible"
	case ModeCanBackground:
		return "ModeCanBackground"
	case ModeControlPanel:
		return "ModeControlPanel"
	case ModeDeskAccessory:
		return "ModeDeskAccessory"
	case ModeDisplayManagerAware:
		return "ModeDisplayManagerAware"
	case ModeDoesActivateOnFGSwitch:
		return "ModeDoesActivateOnFGSwitch"
	case ModeGetAppDiedMsg:
		return "ModeGetAppDiedMsg"
	case ModeGetFrontClicks:
		return "ModeGetFrontClicks"
	case ModeHighLevelEventAware:
		return "ModeHighLevelEventAware"
	case ModeLaunchDontSwitch:
		return "ModeLaunchDontSwitch"
	case ModeLiteral:
		return "ModeLiteral"
	case ModeLocalAndRemoteHLEvents:
		return "ModeLocalAndRemoteHLEvents"
	case ModeMultiLaunch:
		return "ModeMultiLaunch"
	case ModeNeedSuspendResume:
		return "ModeNeedSuspendResume"
	case ModeNormal:
		return "ModeNormal"
	case ModeOnlyBackground:
		return "ModeOnlyBackground"
	case ModePhonemes:
		return "ModePhonemes"
	case ModeReserved:
		return "ModeReserved"
	case ModeStationeryAware:
		return "ModeStationeryAware"
	case ModeText:
		return "ModeText"
	case ModeTune:
		return "ModeTune"
	case ModeUseTextEditServices:
		return "ModeUseTextEditServices"
	default:
		return fmt.Sprintf("Mode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/applicationservices/pmdataformat
type PMDataFormat uint32

const (
	// KPMDataFormatXMLCompressed: # Discussion
	KPMDataFormatXMLCompressed PMDataFormat = 2
	// KPMDataFormatXMLDefault: # Discussion
	KPMDataFormatXMLDefault PMDataFormat = 0
	// KPMDataFormatXMLMinimal: # Discussion
	KPMDataFormatXMLMinimal PMDataFormat = 1
)

func (e PMDataFormat) String() string {
	switch e {
	case KPMDataFormatXMLCompressed:
		return "KPMDataFormatXMLCompressed"
	case KPMDataFormatXMLDefault:
		return "KPMDataFormatXMLDefault"
	case KPMDataFormatXMLMinimal:
		return "KPMDataFormatXMLMinimal"
	default:
		return fmt.Sprintf("PMDataFormat(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/applicationservices/pmpagetopapermappingtype
type PMPageToPaperMappingType uint32

const (
	KPMPageToPaperMappingNone       PMPageToPaperMappingType = 1
	KPMPageToPaperMappingScaleToFit PMPageToPaperMappingType = 2
)

func (e PMPageToPaperMappingType) String() string {
	switch e {
	case KPMPageToPaperMappingNone:
		return "KPMPageToPaperMappingNone"
	case KPMPageToPaperMappingScaleToFit:
		return "KPMPageToPaperMappingScaleToFit"
	default:
		return fmt.Sprintf("PMPageToPaperMappingType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/applicationservices/pasteboardflavorflags
type PasteboardFlavorFlags uint32

const (
	NotSaved         PasteboardFlavorFlags = 0
	Promised         PasteboardFlavorFlags = 0
	RequestOnly      PasteboardFlavorFlags = 0
	SenderOnly       PasteboardFlavorFlags = 0
	SenderTranslated PasteboardFlavorFlags = 0
	SystemTranslated PasteboardFlavorFlags = 0
)

func (e PasteboardFlavorFlags) String() string {
	switch e {
	case NotSaved:
		return "NotSaved"
	default:
		return fmt.Sprintf("PasteboardFlavorFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/applicationservices/pasteboardstandardlocation
type PasteboardStandardLocation int

const (
	KPasteboardStandardLocationTrash   PasteboardStandardLocation = 't'<<24 | 'r'<<16 | 's'<<8 | 'h' // 'trsh'
	KPasteboardStandardLocationUnknown PasteboardStandardLocation = 'u'<<24 | 'n'<<16 | 'k'<<8 | 'n' // 'unkn'
)

func (e PasteboardStandardLocation) String() string {
	switch e {
	case KPasteboardStandardLocationTrash:
		return "KPasteboardStandardLocationTrash"
	case KPasteboardStandardLocationUnknown:
		return "KPasteboardStandardLocationUnknown"
	default:
		return fmt.Sprintf("PasteboardStandardLocation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/applicationservices/pasteboardsyncflags
type PasteboardSyncFlags uint32

const (
	ClientIsOwner PasteboardSyncFlags = 0
	Modified      PasteboardSyncFlags = 0
)

func (e PasteboardSyncFlags) String() string {
	switch e {
	case ClientIsOwner:
		return "ClientIsOwner"
	default:
		return fmt.Sprintf("PasteboardSyncFlags(%d)", e)
	}
}

type So uint

const (
	// SoCharacterMode: # Discussion
	SoCharacterMode So = 'c'<<24 | 'h'<<16 | 'a'<<8 | 'r' // 'char'
	// SoCommandDelimiter: # Discussion
	SoCommandDelimiter So = 'd'<<24 | 'l'<<16 | 'i'<<8 | 'm' // 'dlim'
	// SoCurrentA5: # Discussion
	SoCurrentA5 So = 'm'<<24 | 'y'<<16 | 'A'<<8 | '5' // 'myA5'
	// SoCurrentVoice: # Discussion
	SoCurrentVoice So = 'c'<<24 | 'v'<<16 | 'o'<<8 | 'x' // 'cvox'
	// SoErrorCallBack: # Discussion
	SoErrorCallBack So = 'e'<<24 | 'r'<<16 | 'c'<<8 | 'b' // 'ercb'
	// SoErrors: # Discussion
	SoErrors So = 'e'<<24 | 'r'<<16 | 'r'<<8 | 'o' // 'erro'
	// SoInputMode: # Discussion
	SoInputMode So = 'i'<<24 | 'n'<<16 | 'p'<<8 | 't' // 'inpt'
	// SoNumberMode: # Discussion
	SoNumberMode So = 'n'<<24 | 'm'<<16 | 'b'<<8 | 'r' // 'nmbr'
	// SoOutputToAudioDevice: # Discussion
	SoOutputToAudioDevice So = 'o'<<24 | 'p'<<16 | 'a'<<8 | 'd' // 'opad'
	// SoOutputToExtAudioFile: Pass an ExtAudioFileRef in the `speechInfo` parameter to write to this file, or [NULL] to generate sound.
	SoOutputToExtAudioFile So = 'o'<<24 | 'p'<<16 | 'a'<<8 | 'x' // 'opax'
	// SoOutputToFileWithCFURL: Pass a [CFURLRef] in the `speechInfo` parameter to write to this file, or [NULL] to generate sound.
	SoOutputToFileWithCFURL So = 'o'<<24 | 'p'<<16 | 'a'<<8 | 'f' // 'opaf'
	// SoPhonemeCallBack: # Discussion
	SoPhonemeCallBack So = 'p'<<24 | 'h'<<16 | 'c'<<8 | 'b' // 'phcb'
	// SoPhonemeOptions: Get or set options for the generation of phonetic output.
	SoPhonemeOptions So = 'p'<<24 | 'o'<<16 | 'p'<<8 | 't' // 'popt'
	// SoPhonemeSymbols: # Discussion
	SoPhonemeSymbols So = 'p'<<24 | 'h'<<16 | 's'<<8 | 'y' // 'phsy'
	// SoPitchBase: # Discussion
	SoPitchBase So = 'p'<<24 | 'b'<<16 | 'a'<<8 | 's' // 'pbas'
	// SoPitchMod: # Discussion
	SoPitchMod So = 'p'<<24 | 'm'<<16 | 'o'<<8 | 'd' // 'pmod'
	// SoRate: # Discussion
	SoRate So = 'r'<<24 | 'a'<<16 | 't'<<8 | 'e' // 'rate'
	// SoRecentSync: # Discussion
	SoRecentSync So = 's'<<24 | 'y'<<16 | 'n'<<8 | 'c' // 'sync'
	// SoRefCon: # Discussion
	SoRefCon So = 'r'<<24 | 'e'<<16 | 'f'<<8 | 'c' // 'refc'
	// SoReset: # Discussion
	SoReset So = 'r'<<24 | 's'<<16 | 'e'<<8 | 't' // 'rset'
	// SoSoundOutput: Get or set the speech channel’s current outputchannel.
	SoSoundOutput So = 's'<<24 | 'n'<<16 | 'd'<<8 | 'o' // 'sndo'
	// SoSpeechDoneCallBack: # Discussion
	SoSpeechDoneCallBack So = 's'<<24 | 'd'<<16 | 'c'<<8 | 'b' // 'sdcb'
	// SoStatus: # Discussion
	SoStatus So = 's'<<24 | 't'<<16 | 'a'<<8 | 't' // 'stat'
	// SoSyncCallBack: # Discussion
	SoSyncCallBack So = 's'<<24 | 'y'<<16 | 'c'<<8 | 'b' // 'sycb'
	// SoSynthExtension: # Discussion
	SoSynthExtension So = 'x'<<24 | 't'<<16 | 'n'<<8 | 'd' // 'xtnd'
	// SoSynthType: # Discussion
	SoSynthType So = 'v'<<24 | 'e'<<16 | 'r'<<8 | 's' // 'vers'
	// SoTextDoneCallBack: # Discussion
	SoTextDoneCallBack So = 't'<<24 | 'd'<<16 | 'c'<<8 | 'b' // 'tdcb'
	// SoVolume: # Discussion
	SoVolume So = 'v'<<24 | 'o'<<16 | 'l'<<8 | 'm' // 'volm'
	// SoWordCallBack: # Discussion
	SoWordCallBack So = 'w'<<24 | 'd'<<16 | 'c'<<8 | 'b' // 'wdcb'
)

func (e So) String() string {
	switch e {
	case SoCharacterMode:
		return "SoCharacterMode"
	case SoCommandDelimiter:
		return "SoCommandDelimiter"
	case SoCurrentA5:
		return "SoCurrentA5"
	case SoCurrentVoice:
		return "SoCurrentVoice"
	case SoErrorCallBack:
		return "SoErrorCallBack"
	case SoErrors:
		return "SoErrors"
	case SoInputMode:
		return "SoInputMode"
	case SoNumberMode:
		return "SoNumberMode"
	case SoOutputToAudioDevice:
		return "SoOutputToAudioDevice"
	case SoOutputToExtAudioFile:
		return "SoOutputToExtAudioFile"
	case SoOutputToFileWithCFURL:
		return "SoOutputToFileWithCFURL"
	case SoPhonemeCallBack:
		return "SoPhonemeCallBack"
	case SoPhonemeOptions:
		return "SoPhonemeOptions"
	case SoPhonemeSymbols:
		return "SoPhonemeSymbols"
	case SoPitchBase:
		return "SoPitchBase"
	case SoPitchMod:
		return "SoPitchMod"
	case SoRate:
		return "SoRate"
	case SoRecentSync:
		return "SoRecentSync"
	case SoRefCon:
		return "SoRefCon"
	case SoReset:
		return "SoReset"
	case SoSoundOutput:
		return "SoSoundOutput"
	case SoSpeechDoneCallBack:
		return "SoSpeechDoneCallBack"
	case SoStatus:
		return "SoStatus"
	case SoSyncCallBack:
		return "SoSyncCallBack"
	case SoSynthExtension:
		return "SoSynthExtension"
	case SoSynthType:
		return "SoSynthType"
	case SoTextDoneCallBack:
		return "SoTextDoneCallBack"
	case SoVolume:
		return "SoVolume"
	case SoWordCallBack:
		return "SoWordCallBack"
	default:
		return fmt.Sprintf("So(%d)", e)
	}
}

type SoVoice uint

const (
	// SoVoiceDescription: Get basic voice information.
	SoVoiceDescription SoVoice = 'i'<<24 | 'n'<<16 | 'f'<<8 | 'o' // 'info'
	// SoVoiceFile: Get voice file reference information.
	SoVoiceFile SoVoice = 'f'<<24 | 'r'<<16 | 'e'<<8 | 'f' // 'fref'
)

func (e SoVoice) String() string {
	switch e {
	case SoVoiceDescription:
		return "SoVoiceDescription"
	case SoVoiceFile:
		return "SoVoiceFile"
	default:
		return fmt.Sprintf("SoVoice(%d)", e)
	}
}

type Sv uint

const (
	SvAll1BitData      Sv = 1
	SvAll4BitData      Sv = 2
	SvAll8BitData      Sv = 4
	SvAllAvailableData Sv = 4294967295
	SvAllLargeData     Sv = 255
	SvAllMiniData      Sv = 16711680
	SvAllSmallData     Sv = 65280
	SvLarge1Bit        Sv = 1
	SvLarge4Bit        Sv = 2
	SvLarge8Bit        Sv = 4
	SvMini1Bit         Sv = 65536
	SvMini4Bit         Sv = 131072
	SvMini8Bit         Sv = 262144
	SvSmall1Bit        Sv = 256
	SvSmall4Bit        Sv = 512
	SvSmall8Bit        Sv = 1024
)

func (e Sv) String() string {
	switch e {
	case SvAll1BitData:
		return "SvAll1BitData"
	case SvAll4BitData:
		return "SvAll4BitData"
	case SvAll8BitData:
		return "SvAll8BitData"
	case SvAllAvailableData:
		return "SvAllAvailableData"
	case SvAllLargeData:
		return "SvAllLargeData"
	case SvAllMiniData:
		return "SvAllMiniData"
	case SvAllSmallData:
		return "SvAllSmallData"
	case SvMini1Bit:
		return "SvMini1Bit"
	case SvMini4Bit:
		return "SvMini4Bit"
	case SvMini8Bit:
		return "SvMini8Bit"
	case SvSmall1Bit:
		return "SvSmall1Bit"
	case SvSmall4Bit:
		return "SvSmall4Bit"
	case SvSmall8Bit:
		return "SvSmall8Bit"
	default:
		return fmt.Sprintf("Sv(%d)", e)
	}
}

type Tt uint

const (
	TtDisabled         Tt = 1
	TtLabel1           Tt = 256
	TtLabel2           Tt = 512
	TtLabel3           Tt = 768
	TtLabel4           Tt = 1024
	TtLabel5           Tt = 1280
	TtLabel6           Tt = 1536
	TtLabel7           Tt = 1792
	TtNone             Tt = 0
	TtOffline          Tt = 2
	TtOpen             Tt = 3
	TtSelected         Tt = 16384
	TtSelectedDisabled Tt = 16384
	TtSelectedOffline  Tt = 16384
	TtSelectedOpen     Tt = 16384
)

func (e Tt) String() string {
	switch e {
	case TtDisabled:
		return "TtDisabled"
	case TtLabel1:
		return "TtLabel1"
	case TtLabel2:
		return "TtLabel2"
	case TtLabel3:
		return "TtLabel3"
	case TtLabel4:
		return "TtLabel4"
	case TtLabel5:
		return "TtLabel5"
	case TtLabel6:
		return "TtLabel6"
	case TtLabel7:
		return "TtLabel7"
	case TtNone:
		return "TtNone"
	case TtOffline:
		return "TtOffline"
	case TtOpen:
		return "TtOpen"
	case TtSelected:
		return "TtSelected"
	default:
		return fmt.Sprintf("Tt(%d)", e)
	}
}
