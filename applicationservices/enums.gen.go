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
	StopOnError AXCopyMultipleAttributeOptions = 0x1
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
type AXError int32

const ()

// See: https://developer.apple.com/documentation/applicationservices/axmenuitemmodifiers
type AXMenuItemModifiers uint32

const (
	Control   AXMenuItemModifiers = 4
	NoCommand AXMenuItemModifiers = 8
	Option    AXMenuItemModifiers = 2
	Shift     AXMenuItemModifiers = 1
)

func (e AXMenuItemModifiers) String() string {
	switch e {
	case Control:
		return "Control"
	case NoCommand:
		return "NoCommand"
	case Option:
		return "Option"
	case Shift:
		return "Shift"
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

type At uint32

const (
	AtAbsoluteCenter   At = 5
	AtBottom           At = 3
	AtBottomLeft       At = 11
	AtBottomRight      At = 15
	AtCenterBottom     At = 7
	AtCenterLeft       At = 9
	AtCenterRight      At = 13
	AtCenterTop        At = 6
	AtHorizontalCenter At = 4
	AtLeft             At = 8
	AtNone             At = 0
	AtRight            At = 12
	AtTop              At = 2
	AtTopLeft          At = 10
	AtTopRight         At = 14
	AtVerticalCenter   At = 1
)

func (e At) String() string {
	switch e {
	case AtAbsoluteCenter:
		return "AtAbsoluteCenter"
	case AtBottom:
		return "AtBottom"
	case AtBottomLeft:
		return "AtBottomLeft"
	case AtBottomRight:
		return "AtBottomRight"
	case AtCenterBottom:
		return "AtCenterBottom"
	case AtCenterLeft:
		return "AtCenterLeft"
	case AtCenterRight:
		return "AtCenterRight"
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
	case AtTop:
		return "AtTop"
	case AtTopLeft:
		return "AtTopLeft"
	case AtTopRight:
		return "AtTopRight"
	case AtVerticalCenter:
		return "AtVerticalCenter"
	default:
		return fmt.Sprintf("At(%d)", e)
	}
}

type BadPasteboardSyncErr int32

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

type BadTranslationRef int32

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

type Cdev int32

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

type CmAToB0Tag uint32

const (
	CmAToB0TagValue            CmAToB0Tag = 'A'<<24 | '2'<<16 | 'B'<<8 | '0' // 'A2B0'
	CmAToB1Tag                 CmAToB0Tag = 'A'<<24 | '2'<<16 | 'B'<<8 | '1' // 'A2B1'
	CmAToB2Tag                 CmAToB0Tag = 'A'<<24 | '2'<<16 | 'B'<<8 | '2' // 'A2B2'
	CmBToA0Tag                 CmAToB0Tag = 'B'<<24 | '2'<<16 | 'A'<<8 | '0' // 'B2A0'
	CmBToA1Tag                 CmAToB0Tag = 'B'<<24 | '2'<<16 | 'A'<<8 | '1' // 'B2A1'
	CmBToA2Tag                 CmAToB0Tag = 'B'<<24 | '2'<<16 | 'A'<<8 | '2' // 'B2A2'
	CmBlueColorantTag          CmAToB0Tag = 'b'<<24 | 'X'<<16 | 'Y'<<8 | 'Z' // 'bXYZ'
	CmBlueTRCTag               CmAToB0Tag = 'b'<<24 | 'T'<<16 | 'R'<<8 | 'C' // 'bTRC'
	CmCalibrationDateTimeTag   CmAToB0Tag = 'c'<<24 | 'a'<<16 | 'l'<<8 | 't' // 'calt'
	CmCharTargetTag            CmAToB0Tag = 't'<<24 | 'a'<<16 | 'r'<<8 | 'g' // 'targ'
	CmChromaticAdaptationTag   CmAToB0Tag = 'c'<<24 | 'h'<<16 | 'a'<<8 | 'd' // 'chad'
	CmCopyrightTag             CmAToB0Tag = 'c'<<24 | 'p'<<16 | 'r'<<8 | 't' // 'cprt'
	CmDeviceMfgDescTag         CmAToB0Tag = 'd'<<24 | 'm'<<16 | 'n'<<8 | 'd' // 'dmnd'
	CmDeviceModelDescTag       CmAToB0Tag = 'd'<<24 | 'm'<<16 | 'd'<<8 | 'd' // 'dmdd'
	CmGamutTag                 CmAToB0Tag = 'g'<<24 | 'a'<<16 | 'm'<<8 | 't' // 'gamt'
	CmGrayTRCTag               CmAToB0Tag = 'k'<<24 | 'T'<<16 | 'R'<<8 | 'C' // 'kTRC'
	CmGreenColorantTag         CmAToB0Tag = 'g'<<24 | 'X'<<16 | 'Y'<<8 | 'Z' // 'gXYZ'
	CmGreenTRCTag              CmAToB0Tag = 'g'<<24 | 'T'<<16 | 'R'<<8 | 'C' // 'gTRC'
	CmLuminanceTag             CmAToB0Tag = 'l'<<24 | 'u'<<16 | 'm'<<8 | 'i' // 'lumi'
	CmMeasurementTag           CmAToB0Tag = 'm'<<24 | 'e'<<16 | 'a'<<8 | 's' // 'meas'
	CmMediaBlackPointTag       CmAToB0Tag = 'b'<<24 | 'k'<<16 | 'p'<<8 | 't' // 'bkpt'
	CmMediaWhitePointTag       CmAToB0Tag = 'w'<<24 | 't'<<16 | 'p'<<8 | 't' // 'wtpt'
	CmNamedColor2Tag           CmAToB0Tag = 'n'<<24 | 'c'<<16 | 'l'<<8 | '2' // 'ncl2'
	CmNamedColorTag            CmAToB0Tag = 'n'<<24 | 'c'<<16 | 'o'<<8 | 'l' // 'ncol'
	CmPS2CRD0Tag               CmAToB0Tag = 'p'<<24 | 's'<<16 | 'd'<<8 | '0' // 'psd0'
	CmPS2CRD1Tag               CmAToB0Tag = 'p'<<24 | 's'<<16 | 'd'<<8 | '1' // 'psd1'
	CmPS2CRD2Tag               CmAToB0Tag = 'p'<<24 | 's'<<16 | 'd'<<8 | '2' // 'psd2'
	CmPS2CRD3Tag               CmAToB0Tag = 'p'<<24 | 's'<<16 | 'd'<<8 | '3' // 'psd3'
	CmPS2CSATag                CmAToB0Tag = 'p'<<24 | 's'<<16 | '2'<<8 | 's' // 'ps2s'
	CmPS2RenderingIntentTag    CmAToB0Tag = 'p'<<24 | 's'<<16 | '2'<<8 | 'i' // 'ps2i'
	CmPreview0Tag              CmAToB0Tag = 'p'<<24 | 'r'<<16 | 'e'<<8 | '0' // 'pre0'
	CmPreview1Tag              CmAToB0Tag = 'p'<<24 | 'r'<<16 | 'e'<<8 | '1' // 'pre1'
	CmPreview2Tag              CmAToB0Tag = 'p'<<24 | 'r'<<16 | 'e'<<8 | '2' // 'pre2'
	CmProfileDescriptionTag    CmAToB0Tag = 'd'<<24 | 'e'<<16 | 's'<<8 | 'c' // 'desc'
	CmProfileSequenceDescTag   CmAToB0Tag = 'p'<<24 | 's'<<16 | 'e'<<8 | 'q' // 'pseq'
	CmRedColorantTag           CmAToB0Tag = 'r'<<24 | 'X'<<16 | 'Y'<<8 | 'Z' // 'rXYZ'
	CmRedTRCTag                CmAToB0Tag = 'r'<<24 | 'T'<<16 | 'R'<<8 | 'C' // 'rTRC'
	CmScreeningDescTag         CmAToB0Tag = 's'<<24 | 'c'<<16 | 'r'<<8 | 'd' // 'scrd'
	CmScreeningTag             CmAToB0Tag = 's'<<24 | 'c'<<16 | 'r'<<8 | 'n' // 'scrn'
	CmTechnologyTag            CmAToB0Tag = 't'<<24 | 'e'<<16 | 'c'<<8 | 'h' // 'tech'
	CmUcrBgTag                 CmAToB0Tag = 'b'<<24 | 'f'<<16 | 'd'<<8 | ' ' // 'bfd '
	CmViewingConditionsDescTag CmAToB0Tag = 'v'<<24 | 'u'<<16 | 'e'<<8 | 'd' // 'vued'
	CmViewingConditionsTag     CmAToB0Tag = 'v'<<24 | 'i'<<16 | 'e'<<8 | 'w' // 'view'
)

func (e CmAToB0Tag) String() string {
	switch e {
	case CmAToB0TagValue:
		return "CmAToB0TagValue"
	case CmAToB1Tag:
		return "CmAToB1Tag"
	case CmAToB2Tag:
		return "CmAToB2Tag"
	case CmBToA0Tag:
		return "CmBToA0Tag"
	case CmBToA1Tag:
		return "CmBToA1Tag"
	case CmBToA2Tag:
		return "CmBToA2Tag"
	case CmBlueColorantTag:
		return "CmBlueColorantTag"
	case CmBlueTRCTag:
		return "CmBlueTRCTag"
	case CmCalibrationDateTimeTag:
		return "CmCalibrationDateTimeTag"
	case CmCharTargetTag:
		return "CmCharTargetTag"
	case CmChromaticAdaptationTag:
		return "CmChromaticAdaptationTag"
	case CmCopyrightTag:
		return "CmCopyrightTag"
	case CmDeviceMfgDescTag:
		return "CmDeviceMfgDescTag"
	case CmDeviceModelDescTag:
		return "CmDeviceModelDescTag"
	case CmGamutTag:
		return "CmGamutTag"
	case CmGrayTRCTag:
		return "CmGrayTRCTag"
	case CmGreenColorantTag:
		return "CmGreenColorantTag"
	case CmGreenTRCTag:
		return "CmGreenTRCTag"
	case CmLuminanceTag:
		return "CmLuminanceTag"
	case CmMeasurementTag:
		return "CmMeasurementTag"
	case CmMediaBlackPointTag:
		return "CmMediaBlackPointTag"
	case CmMediaWhitePointTag:
		return "CmMediaWhitePointTag"
	case CmNamedColor2Tag:
		return "CmNamedColor2Tag"
	case CmNamedColorTag:
		return "CmNamedColorTag"
	case CmPS2CRD0Tag:
		return "CmPS2CRD0Tag"
	case CmPS2CRD1Tag:
		return "CmPS2CRD1Tag"
	case CmPS2CRD2Tag:
		return "CmPS2CRD2Tag"
	case CmPS2CRD3Tag:
		return "CmPS2CRD3Tag"
	case CmPS2CSATag:
		return "CmPS2CSATag"
	case CmPS2RenderingIntentTag:
		return "CmPS2RenderingIntentTag"
	case CmPreview0Tag:
		return "CmPreview0Tag"
	case CmPreview1Tag:
		return "CmPreview1Tag"
	case CmPreview2Tag:
		return "CmPreview2Tag"
	case CmProfileDescriptionTag:
		return "CmProfileDescriptionTag"
	case CmProfileSequenceDescTag:
		return "CmProfileSequenceDescTag"
	case CmRedColorantTag:
		return "CmRedColorantTag"
	case CmRedTRCTag:
		return "CmRedTRCTag"
	case CmScreeningDescTag:
		return "CmScreeningDescTag"
	case CmScreeningTag:
		return "CmScreeningTag"
	case CmTechnologyTag:
		return "CmTechnologyTag"
	case CmUcrBgTag:
		return "CmUcrBgTag"
	case CmViewingConditionsDescTag:
		return "CmViewingConditionsDescTag"
	case CmViewingConditionsTag:
		return "CmViewingConditionsTag"
	default:
		return fmt.Sprintf("CmAToB0Tag(%d)", e)
	}
}

type CmAsciiData uint32

const (
	// CmAsciiDataValue: ASCII data.
	CmAsciiDataValue CmAsciiData = 0
	// CmBinaryData: Binary data.
	CmBinaryData CmAsciiData = 1
)

func (e CmAsciiData) String() string {
	switch e {
	case CmAsciiDataValue:
		return "CmAsciiDataValue"
	case CmBinaryData:
		return "CmBinaryData"
	default:
		return fmt.Sprintf("CmAsciiData(%d)", e)
	}
}

type CmBlackPoint uint32

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

type CmCS1 uint32

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

type CmColorSpace int32

const (
	CmColorSpaceAlphaMask         CmColorSpace = 0x80
	CmColorSpaceEncodingMask      CmColorSpace = 0xf0000
	CmColorSpacePackingMask       CmColorSpace = 0xff00
	CmColorSpacePremulAlphaMask   CmColorSpace = 0x40
	CmColorSpaceReservedMask      CmColorSpace = -1048576
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

type CmCurrent uint32

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

type CmDefault uint32

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

type CmDevice uint32

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

type CmDeviceDBNotFoundErr int32

const (
	// CmDeviceAlreadyRegistered: Device already registered; returned by a CM device integration routine.
	CmDeviceAlreadyRegistered CmDeviceDBNotFoundErr = -4228
	// CmDeviceDBNotFoundErrValue: Preferences not found or loaded; returned by a CM device integration routine.
	CmDeviceDBNotFoundErrValue CmDeviceDBNotFoundErr = -4227
	// CmDeviceNotRegistered: Device not found; returned by a CM device integration routine.
	CmDeviceNotRegistered CmDeviceDBNotFoundErr = -4229
	// CmDeviceProfilesNotFound: Profiles not found; returned by a CM device integration routine.
	CmDeviceProfilesNotFound CmDeviceDBNotFoundErr = -4230
	// CmInternalCFErr: CoreFoundation failure; returned by a CM device integration routine.
	CmInternalCFErr   CmDeviceDBNotFoundErr = -4231
	CmPrefsSynchError CmDeviceDBNotFoundErr = -4232
)

func (e CmDeviceDBNotFoundErr) String() string {
	switch e {
	case CmDeviceAlreadyRegistered:
		return "CmDeviceAlreadyRegistered"
	case CmDeviceDBNotFoundErrValue:
		return "CmDeviceDBNotFoundErrValue"
	case CmDeviceNotRegistered:
		return "CmDeviceNotRegistered"
	case CmDeviceProfilesNotFound:
		return "CmDeviceProfilesNotFound"
	case CmInternalCFErr:
		return "CmInternalCFErr"
	case CmPrefsSynchError:
		return "CmPrefsSynchError"
	default:
		return fmt.Sprintf("CmDeviceDBNotFoundErr(%d)", e)
	}
}

type CmDeviceState int32

const (
	CmDeviceStateAppleRsvdBits  CmDeviceState = -16711681
	CmDeviceStateBusy           CmDeviceState = 0x2
	CmDeviceStateDefault        CmDeviceState = 0
	CmDeviceStateDeviceRsvdBits CmDeviceState = 0xff0000
	CmDeviceStateForceNotify    CmDeviceState = -2147483648
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

type CmEmbedded uint32

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

type CmFlare0 uint32

const (
	CmFlare0Value CmFlare0 = 0
	CmFlare100    CmFlare0 = 0x1
)

func (e CmFlare0) String() string {
	switch e {
	case CmFlare0Value:
		return "CmFlare0Value"
	case CmFlare100:
		return "CmFlare100"
	default:
		return fmt.Sprintf("CmFlare0(%d)", e)
	}
}

type CmGeometryUnknown uint32

const (
	CmGeometry045or450     CmGeometryUnknown = 0x1
	CmGeometry0dord0       CmGeometryUnknown = 0x2
	CmGeometryUnknownValue CmGeometryUnknown = 0
)

func (e CmGeometryUnknown) String() string {
	switch e {
	case CmGeometry045or450:
		return "CmGeometry045or450"
	case CmGeometry0dord0:
		return "CmGeometry0dord0"
	case CmGeometryUnknownValue:
		return "CmGeometryUnknownValue"
	default:
		return fmt.Sprintf("CmGeometryUnknown(%d)", e)
	}
}

type CmGray8Space uint32

const (
	CmARGB32PmulSpace CmGray8Space = 6337
	// CmARGB32Space: # Discussion
	CmARGB32Space      CmGray8Space = 6273
	CmARGB64LPmulSpace CmGray8Space = 31425
	CmARGB64LSpace     CmGray8Space = 31361
	CmARGB64PmulSpace  CmGray8Space = 15041
	CmARGB64Space      CmGray8Space = 14977
	// CmCMYK32Space: A CMYK color space composed of cyan, magenta, yellow, and black components whose values are packed with 8 bits of storage per component.
	CmCMYK32Space  CmGray8Space = 2050
	CmCMYK64LSpace CmGray8Space = 27138
	// CmCMYK64Space: A CMYK color space composed of cyan, magenta, yellow, and black components whose values are packed with 16 bits of storage per component.
	CmCMYK64Space CmGray8Space = 10754
	// CmGamutResult1Space: # Discussion
	CmGamutResult1Space CmGray8Space = 2828
	CmGray16LSpace      CmGray8Space = 16394
	// CmGray16Space: A luminance color space with a single 16-bit component, gray.
	CmGray16Space       CmGray8Space = 10
	CmGray8SpaceValue   CmGray8Space = 10250
	CmGrayA16PmulSpace  CmGray8Space = 8394
	CmGrayA16Space      CmGray8Space = 8330
	CmGrayA32LPmulSpace CmGray8Space = 16586
	CmGrayA32LSpace     CmGray8Space = 16522
	CmGrayA32PmulSpace  CmGray8Space = 202
	// CmGrayA32Space: A luminance color space with two components, a gray component followed by an alpha channel component.
	CmGrayA32Space CmGray8Space = 138
	// CmHLS32Space: An HLS color space composed of hue, lightness, and saturation components whose values are packed with 10 bits of storage per component.
	CmHLS32Space CmGray8Space = 2564
	// CmHSV32Space: An HSV color space composed of hue, saturation, and value components whose values are packed with 10 bits of storage per component.
	CmHSV32Space CmGray8Space = 2563
	// CmLAB24Space: # Discussion
	CmLAB24Space CmGray8Space = 8456
	// CmLAB32Space: # Discussion
	CmLAB32Space  CmGray8Space = 2568
	CmLAB48LSpace CmGray8Space = 26888
	// CmLAB48Space: # Discussion
	CmLAB48Space CmGray8Space = 10504
	// CmLUV32Space: An L*u*v* color space composed of L*, u*, and v* components whose values are packed with 10 bits per component.
	CmLUV32Space CmGray8Space = 2567
	// CmMCEight8Space: An eight-channel multichannel (HiFi) data color space, whose values are packed with 8 bits per component.
	CmMCEight8Space CmGray8Space = 9492
	// CmMCFive8Space: A five-channel multichannel (HiFi) data color space, whose values are packed with 8 bits per component.
	CmMCFive8Space CmGray8Space = 8721
	// CmMCSeven8Space: A seven-channel multichannel (HiFi) data color space, whose values are packed with 8 bits per component.
	CmMCSeven8Space CmGray8Space = 9235
	// CmMCSix8Space: A six-channel multichannel (HiFi) data color space, whose values are packed with 8 bits per component.
	CmMCSix8Space          CmGray8Space = 8978
	CmNamedIndexed32LSpace CmGray8Space = 26384
	// CmNamedIndexed32Space: A color space where each color is stored as a single 32-bit value, specifying an index into a named color space.
	CmNamedIndexed32Space CmGray8Space = 10000
	CmRGB16LSpace         CmGray8Space = 17665
	// CmRGB16Space: An RGB color space composed of red, green, and blue components whose values are packed with 5 bits of storage per component.
	CmRGB16Space CmGray8Space = 1281
	// CmRGB24Space: An RGB color space composed of red, green, and blue components whose values are packed with 8 bits of storage per component.
	CmRGB24Space CmGray8Space = 8449
	// CmRGB32Space: An RGB color space composed of red, green, and blue components whose values are packed with 8 bits of storage per component.
	CmRGB32Space  CmGray8Space = 2049
	CmRGB48LSpace CmGray8Space = 26881
	// CmRGB48Space: An RGB color space composed of red, green, and blue components whose values are packed with 16 bits of storage per component.
	CmRGB48Space      CmGray8Space = 10497
	CmRGB565LSpace    CmGray8Space = 17921
	CmRGB565Space     CmGray8Space = 1537
	CmRGBA32PmulSpace CmGray8Space = 2241
	// CmRGBA32Space: An RGB color space composed of red, green, and blue color value components, followed by an alpha channel component.
	CmRGBA32Space      CmGray8Space = 2177
	CmRGBA64LPmulSpace CmGray8Space = 27329
	CmRGBA64LSpace     CmGray8Space = 27265
	CmRGBA64PmulSpace  CmGray8Space = 10945
	CmRGBA64Space      CmGray8Space = 10881
	CmXYZ24Space       CmGray8Space = 8454
	// CmXYZ32Space: An XYZ color space composed of X, Y, and Z components whose values are packed with 10 bits per component.
	CmXYZ32Space  CmGray8Space = 2566
	CmXYZ48LSpace CmGray8Space = 26886
	CmXYZ48Space  CmGray8Space = 10502
	// CmYXY32Space: A Yxy color space composed of Y, x, and y components whose values are packed with 10 bits of storage per component.
	CmYXY32Space CmGray8Space = 2565
)

func (e CmGray8Space) String() string {
	switch e {
	case CmARGB32PmulSpace:
		return "CmARGB32PmulSpace"
	case CmARGB32Space:
		return "CmARGB32Space"
	case CmARGB64LPmulSpace:
		return "CmARGB64LPmulSpace"
	case CmARGB64LSpace:
		return "CmARGB64LSpace"
	case CmARGB64PmulSpace:
		return "CmARGB64PmulSpace"
	case CmARGB64Space:
		return "CmARGB64Space"
	case CmCMYK32Space:
		return "CmCMYK32Space"
	case CmCMYK64LSpace:
		return "CmCMYK64LSpace"
	case CmCMYK64Space:
		return "CmCMYK64Space"
	case CmGamutResult1Space:
		return "CmGamutResult1Space"
	case CmGray16LSpace:
		return "CmGray16LSpace"
	case CmGray16Space:
		return "CmGray16Space"
	case CmGray8SpaceValue:
		return "CmGray8SpaceValue"
	case CmGrayA16PmulSpace:
		return "CmGrayA16PmulSpace"
	case CmGrayA16Space:
		return "CmGrayA16Space"
	case CmGrayA32LPmulSpace:
		return "CmGrayA32LPmulSpace"
	case CmGrayA32LSpace:
		return "CmGrayA32LSpace"
	case CmGrayA32PmulSpace:
		return "CmGrayA32PmulSpace"
	case CmGrayA32Space:
		return "CmGrayA32Space"
	case CmHLS32Space:
		return "CmHLS32Space"
	case CmHSV32Space:
		return "CmHSV32Space"
	case CmLAB24Space:
		return "CmLAB24Space"
	case CmLAB32Space:
		return "CmLAB32Space"
	case CmLAB48LSpace:
		return "CmLAB48LSpace"
	case CmLAB48Space:
		return "CmLAB48Space"
	case CmLUV32Space:
		return "CmLUV32Space"
	case CmMCEight8Space:
		return "CmMCEight8Space"
	case CmMCFive8Space:
		return "CmMCFive8Space"
	case CmMCSeven8Space:
		return "CmMCSeven8Space"
	case CmMCSix8Space:
		return "CmMCSix8Space"
	case CmNamedIndexed32LSpace:
		return "CmNamedIndexed32LSpace"
	case CmNamedIndexed32Space:
		return "CmNamedIndexed32Space"
	case CmRGB16LSpace:
		return "CmRGB16LSpace"
	case CmRGB16Space:
		return "CmRGB16Space"
	case CmRGB24Space:
		return "CmRGB24Space"
	case CmRGB32Space:
		return "CmRGB32Space"
	case CmRGB48LSpace:
		return "CmRGB48LSpace"
	case CmRGB48Space:
		return "CmRGB48Space"
	case CmRGB565LSpace:
		return "CmRGB565LSpace"
	case CmRGB565Space:
		return "CmRGB565Space"
	case CmRGBA32PmulSpace:
		return "CmRGBA32PmulSpace"
	case CmRGBA32Space:
		return "CmRGBA32Space"
	case CmRGBA64LPmulSpace:
		return "CmRGBA64LPmulSpace"
	case CmRGBA64LSpace:
		return "CmRGBA64LSpace"
	case CmRGBA64PmulSpace:
		return "CmRGBA64PmulSpace"
	case CmRGBA64Space:
		return "CmRGBA64Space"
	case CmXYZ24Space:
		return "CmXYZ24Space"
	case CmXYZ32Space:
		return "CmXYZ32Space"
	case CmXYZ48LSpace:
		return "CmXYZ48LSpace"
	case CmXYZ48Space:
		return "CmXYZ48Space"
	case CmYXY32Space:
		return "CmYXY32Space"
	default:
		return fmt.Sprintf("CmGray8Space(%d)", e)
	}
}

type CmICCProfileVersion4 uint32

const (
	CmCS1ProfileVersion       CmICCProfileVersion4 = 0x100
	CmCS2ProfileVersion       CmICCProfileVersion4 = 33554432
	CmICCProfileVersion2      CmICCProfileVersion4 = 0x2000000
	CmICCProfileVersion21     CmICCProfileVersion4 = 0x2100000
	CmICCProfileVersion4Value CmICCProfileVersion4 = 0x4000000
)

func (e CmICCProfileVersion4) String() string {
	switch e {
	case CmCS1ProfileVersion:
		return "CmCS1ProfileVersion"
	case CmCS2ProfileVersion:
		return "CmCS2ProfileVersion"
	case CmICCProfileVersion21:
		return "CmICCProfileVersion21"
	case CmICCProfileVersion4Value:
		return "CmICCProfileVersion4Value"
	default:
		return fmt.Sprintf("CmICCProfileVersion4(%d)", e)
	}
}

type CmICCReservedFlagsMask int32

const (
	CmBlackPointCompensationMask CmICCReservedFlagsMask = 0x4
	// CmCMSReservedFlagsMask: # Discussion
	CmCMSReservedFlagsMask CmICCReservedFlagsMask = -65536
	// CmEmbeddedMask: This mask provides access to bit 0 of the `flags` field, which specifies whether the profile is embedded.
	CmEmbeddedMask CmICCReservedFlagsMask = 0x1
	// CmEmbeddedUseMask: # Discussion
	CmEmbeddedUseMask CmICCReservedFlagsMask = 0x2
	// CmGamutCheckingMask: # Discussion
	CmGamutCheckingMask CmICCReservedFlagsMask = 0x80000
	// CmICCReservedFlagsMaskValue: # Discussion
	CmICCReservedFlagsMaskValue CmICCReservedFlagsMask = 0xffff
	// CmInterpolationMask: # Discussion
	CmInterpolationMask CmICCReservedFlagsMask = 0x40000
	// CmQualityMask: # Discussion
	CmQualityMask CmICCReservedFlagsMask = 0x30000
)

func (e CmICCReservedFlagsMask) String() string {
	switch e {
	case CmBlackPointCompensationMask:
		return "CmBlackPointCompensationMask"
	case CmCMSReservedFlagsMask:
		return "CmCMSReservedFlagsMask"
	case CmEmbeddedMask:
		return "CmEmbeddedMask"
	case CmEmbeddedUseMask:
		return "CmEmbeddedUseMask"
	case CmGamutCheckingMask:
		return "CmGamutCheckingMask"
	case CmICCReservedFlagsMaskValue:
		return "CmICCReservedFlagsMaskValue"
	case CmInterpolationMask:
		return "CmInterpolationMask"
	case CmQualityMask:
		return "CmQualityMask"
	default:
		return fmt.Sprintf("CmICCReservedFlagsMask(%d)", e)
	}
}

type CmIlluminant uint32

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

type CmInputClass uint32

const (
	// CmAbstractClass: An abstract profile.
	CmAbstractClass CmInputClass = 'a'<<24 | 'b'<<16 | 's'<<8 | 't' // 'abst'
	// CmColorSpaceClass: A color space profile.
	CmColorSpaceClass CmInputClass = 's'<<24 | 'p'<<16 | 'a'<<8 | 'c' // 'spac'
	// CmDisplayClass: A display device profile defined for a monitor.
	CmDisplayClass CmInputClass = 'm'<<24 | 'n'<<16 | 't'<<8 | 'r' // 'mntr'
	// CmInputClassValue: An input device profile defined for a scanner.
	CmInputClassValue CmInputClass = 's'<<24 | 'c'<<16 | 'n'<<8 | 'r' // 'scnr'
	// CmLinkClass: A device link profile.
	CmLinkClass CmInputClass = 'l'<<24 | 'i'<<16 | 'n'<<8 | 'k' // 'link'
	// CmNamedColorClass: A named color space profile.
	CmNamedColorClass CmInputClass = 'n'<<24 | 'm'<<16 | 'c'<<8 | 'l' // 'nmcl'
	// CmOutputClass: An output device profile defined for a printer.
	CmOutputClass CmInputClass = 'p'<<24 | 'r'<<16 | 't'<<8 | 'r' // 'prtr'
)

func (e CmInputClass) String() string {
	switch e {
	case CmAbstractClass:
		return "CmAbstractClass"
	case CmColorSpaceClass:
		return "CmColorSpaceClass"
	case CmDisplayClass:
		return "CmDisplayClass"
	case CmInputClassValue:
		return "CmInputClassValue"
	case CmLinkClass:
		return "CmLinkClass"
	case CmNamedColorClass:
		return "CmNamedColorClass"
	case CmOutputClass:
		return "CmOutputClass"
	default:
		return fmt.Sprintf("CmInputClass(%d)", e)
	}
}

type CmInputUse uint32

const (
	CmDisplayUse    CmInputUse = 'd'<<24 | 'p'<<16 | 'l'<<8 | 'y' // 'dply'
	CmInputUseValue CmInputUse = 'i'<<24 | 'n'<<16 | 'p'<<8 | 't' // 'inpt'
	CmOutputUse     CmInputUse = 'o'<<24 | 'u'<<16 | 't'<<8 | 'p' // 'outp'
	CmProofUse      CmInputUse = 'p'<<24 | 'r'<<16 | 'u'<<8 | 'f' // 'pruf'
)

func (e CmInputUse) String() string {
	switch e {
	case CmDisplayUse:
		return "CmDisplayUse"
	case CmInputUseValue:
		return "CmInputUseValue"
	case CmOutputUse:
		return "CmOutputUse"
	case CmProofUse:
		return "CmProofUse"
	default:
		return fmt.Sprintf("CmInputUse(%d)", e)
	}
}

type CmIterate uint32

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

type CmMacintosh uint32

const (
	CmMacintoshValue  CmMacintosh = 'A'<<24 | 'P'<<16 | 'P'<<8 | 'L' // 'APPL'
	CmMicrosoft       CmMacintosh = 'M'<<24 | 'S'<<16 | 'F'<<8 | 'T' // 'MSFT'
	CmSiliconGraphics CmMacintosh = 'S'<<24 | 'G'<<16 | 'I'<<8 | ' ' // 'SGI '
	CmSolaris         CmMacintosh = 'S'<<24 | 'U'<<16 | 'N'<<8 | 'W' // 'SUNW'
	CmTaligent        CmMacintosh = 'T'<<24 | 'G'<<16 | 'N'<<8 | 'T' // 'TGNT'
)

func (e CmMacintosh) String() string {
	switch e {
	case CmMacintoshValue:
		return "CmMacintoshValue"
	case CmMicrosoft:
		return "CmMicrosoft"
	case CmSiliconGraphics:
		return "CmSiliconGraphics"
	case CmSolaris:
		return "CmSolaris"
	case CmTaligent:
		return "CmTaligent"
	default:
		return fmt.Sprintf("CmMacintosh(%d)", e)
	}
}

type CmMagic uint32

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

type CmNoColorPacking uint32

const (
	Cm16_8ColorPacking CmNoColorPacking = 0x2000
	// Cm24_8ColorPacking: The color values for three 8-bit color channels are stored in consecutive bytes, for a total of 24 bits.
	Cm24_8ColorPacking CmNoColorPacking = 0x2100
	// Cm32_16ColorPacking: The color values for two 16-bit color channels are stored in a 32-bit word.
	Cm32_16ColorPacking CmNoColorPacking = 0x2600
	// Cm32_32ColorPacking: The color value for a 32-bit color channel is stored in a 32-bit word.
	Cm32_32ColorPacking CmNoColorPacking = 0x2700
	// Cm32_8ColorPacking: The color values for four 8-bit color channels are stored in consecutive bytes, for a total of 32 bits.
	Cm32_8ColorPacking CmNoColorPacking = 2048
	// Cm40_8ColorPacking: The color values for five 8-bit color channels are stored in consecutive bytes, for a total of 40 bits.
	Cm40_8ColorPacking CmNoColorPacking = 0x2200
	// Cm48_16ColorPacking: The color values for three 16-bit color channels are stored in 48 consecutive bits.
	Cm48_16ColorPacking CmNoColorPacking = 0x2900
	// Cm48_8ColorPacking: The color values for six 8-bit color channels are stored in consecutive bytes, for a total of 48 bits.
	Cm48_8ColorPacking CmNoColorPacking = 0x2300
	// Cm56_8ColorPacking: The color values for seven 8-bit color channels are stored in consecutive bytes, for a total of 56 bits.
	Cm56_8ColorPacking CmNoColorPacking = 0x2400
	// Cm64_16ColorPacking: The color values for four 16-bit color channels are stored in 64 consecutive bits.
	Cm64_16ColorPacking CmNoColorPacking = 0x2a00
	// Cm64_8ColorPacking: The color values for eight 8-bit color channels are stored in consecutive bytes, for a total of 64 bits.
	Cm64_8ColorPacking CmNoColorPacking = 0x2500
	Cm8_8ColorPacking  CmNoColorPacking = 0x2800
	// CmAlphaFirstPacking: An alpha channel is added to the color value as its first component.
	CmAlphaFirstPacking   CmNoColorPacking = 0x1000
	CmAlphaLastPacking    CmNoColorPacking = 0
	CmLittleEndianPacking CmNoColorPacking = 0x4000
	// CmLong10ColorPacking: The color values for three 10-bit color channels are stored consecutively in a 32-bit long, with the two highest order bits unused.
	CmLong10ColorPacking CmNoColorPacking = 0xa00
	// CmLong8ColorPacking: # Discussion
	CmLong8ColorPacking CmNoColorPacking = 0x800
	// CmNoColorPackingValue: This constant is not used for ColorSync bitmaps.
	CmNoColorPackingValue CmNoColorPacking = 0
	// CmOneBitDirectPacking: One bit is used as the pixel format.
	CmOneBitDirectPacking   CmNoColorPacking = 0xb00
	CmReverseChannelPacking CmNoColorPacking = 0x8000
	CmWord565ColorPacking   CmNoColorPacking = 0x600
	// CmWord5ColorPacking: The color values for three 5-bit color channels are stored consecutively in 16-bits, with the highest order bit unused.
	CmWord5ColorPacking CmNoColorPacking = 0x500
)

func (e CmNoColorPacking) String() string {
	switch e {
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
	case Cm40_8ColorPacking:
		return "Cm40_8ColorPacking"
	case Cm48_16ColorPacking:
		return "Cm48_16ColorPacking"
	case Cm48_8ColorPacking:
		return "Cm48_8ColorPacking"
	case Cm56_8ColorPacking:
		return "Cm56_8ColorPacking"
	case Cm64_16ColorPacking:
		return "Cm64_16ColorPacking"
	case Cm64_8ColorPacking:
		return "Cm64_8ColorPacking"
	case Cm8_8ColorPacking:
		return "Cm8_8ColorPacking"
	case CmAlphaFirstPacking:
		return "CmAlphaFirstPacking"
	case CmAlphaLastPacking:
		return "CmAlphaLastPacking"
	case CmLittleEndianPacking:
		return "CmLittleEndianPacking"
	case CmLong10ColorPacking:
		return "CmLong10ColorPacking"
	case CmOneBitDirectPacking:
		return "CmOneBitDirectPacking"
	case CmReverseChannelPacking:
		return "CmReverseChannelPacking"
	case CmWord565ColorPacking:
		return "CmWord565ColorPacking"
	case CmWord5ColorPacking:
		return "CmWord5ColorPacking"
	default:
		return fmt.Sprintf("CmNoColorPacking(%d)", e)
	}
}

type CmNoProfileBase uint32

const (
	CmBufferBasedProfile CmNoProfileBase = 6
	// CmNoProfileBaseValue: The profile is temporary.
	CmNoProfileBaseValue CmNoProfileBase = 0
	CmPathBasedProfile   CmNoProfileBase = 5
)

func (e CmNoProfileBase) String() string {
	switch e {
	case CmBufferBasedProfile:
		return "CmBufferBasedProfile"
	case CmNoProfileBaseValue:
		return "CmNoProfileBaseValue"
	case CmPathBasedProfile:
		return "CmPathBasedProfile"
	default:
		return fmt.Sprintf("CmNoProfileBase(%d)", e)
	}
}

type CmNoSpace uint32

const (
	// CmAlphaPmulSpace: A premultiplied alpha channel component is added to the color value.
	CmAlphaPmulSpace CmNoSpace = 0x40
	// CmAlphaSpace: An alpha channel component is added to the color value.
	CmAlphaSpace CmNoSpace = 0x80
	// CmCMYKSpace: A CMYK color space composed of cyan, magenta, yellow, and black.
	CmCMYKSpace CmNoSpace = 0x2
	// CmGamutResultSpace: # Discussion
	CmGamutResultSpace CmNoSpace = 0xc
	CmGrayAPmulSpace   CmNoSpace = 202
	// CmGrayASpace: A luminance color space with two components, a gray component followed by an alpha channel component.
	CmGrayASpace CmNoSpace = 138
	// CmGraySpace: A luminance color space with a single component, gray.
	CmGraySpace CmNoSpace = 0xa
	// CmHLSSpace: An HLS color space composed of hue, lightness, and saturation components.
	CmHLSSpace CmNoSpace = 0x4
	// CmHSVSpace: An HSV color space composed of hue, saturation, and value components.
	CmHSVSpace CmNoSpace = 0x3
	// CmLABSpace: An L*a*b* color space composed of L*, a*, b* components.
	CmLABSpace CmNoSpace = 0x8
	// CmLUVSpace: An L*u*v* color space composed of L*, u*, and v* components.
	CmLUVSpace CmNoSpace = 0x7
	// CmMCEightSpace: An eight-channel multichannel (HiFi) data color space.
	CmMCEightSpace CmNoSpace = 0x14
	// CmMCFiveSpace: A five-channel multichannel (HiFi) data color space.
	CmMCFiveSpace CmNoSpace = 0x11
	// CmMCSevenSpace: A seven-channel multichannel (HiFi) data color space.
	CmMCSevenSpace CmNoSpace = 0x13
	// CmMCSixSpace: A six-channel multichannel (HiFi) data color space.
	CmMCSixSpace CmNoSpace = 0x12
	// CmNamedIndexedSpace: A named indexed color space.
	CmNamedIndexedSpace CmNoSpace = 0x10
	// CmNoSpaceValue: The ColorSync Manager does not use this constant.
	CmNoSpaceValue  CmNoSpace = 0
	CmRGBAPmulSpace CmNoSpace = 193
	// CmRGBASpace: # Discussion
	CmRGBASpace CmNoSpace = 129
	// CmRGBSpace: An RGB color space composed of red, green, and blue components.
	CmRGBSpace CmNoSpace = 0x1
	// CmReservedSpace1: This field is reserved for use by QuickDraw GX.
	CmReservedSpace1 CmNoSpace = 0x9
	// CmReservedSpace2: This field is reserved for use by QuickDraw GX.
	CmReservedSpace2 CmNoSpace = 0xb
	// CmXYZSpace: An XYZ color space composed of X, Y, and Z components.
	CmXYZSpace CmNoSpace = 0x6
	// CmYXYSpace: A Yxy color space composed of Y, x, and y components.
	CmYXYSpace CmNoSpace = 0x5
)

func (e CmNoSpace) String() string {
	switch e {
	case CmAlphaPmulSpace:
		return "CmAlphaPmulSpace"
	case CmAlphaSpace:
		return "CmAlphaSpace"
	case CmCMYKSpace:
		return "CmCMYKSpace"
	case CmGamutResultSpace:
		return "CmGamutResultSpace"
	case CmGrayAPmulSpace:
		return "CmGrayAPmulSpace"
	case CmGrayASpace:
		return "CmGrayASpace"
	case CmGraySpace:
		return "CmGraySpace"
	case CmHLSSpace:
		return "CmHLSSpace"
	case CmHSVSpace:
		return "CmHSVSpace"
	case CmLABSpace:
		return "CmLABSpace"
	case CmLUVSpace:
		return "CmLUVSpace"
	case CmMCEightSpace:
		return "CmMCEightSpace"
	case CmMCFiveSpace:
		return "CmMCFiveSpace"
	case CmMCSevenSpace:
		return "CmMCSevenSpace"
	case CmMCSixSpace:
		return "CmMCSixSpace"
	case CmNamedIndexedSpace:
		return "CmNamedIndexedSpace"
	case CmNoSpaceValue:
		return "CmNoSpaceValue"
	case CmRGBAPmulSpace:
		return "CmRGBAPmulSpace"
	case CmRGBASpace:
		return "CmRGBASpace"
	case CmRGBSpace:
		return "CmRGBSpace"
	case CmReservedSpace1:
		return "CmReservedSpace1"
	case CmReservedSpace2:
		return "CmReservedSpace2"
	case CmXYZSpace:
		return "CmXYZSpace"
	case CmYXYSpace:
		return "CmYXYSpace"
	default:
		return fmt.Sprintf("CmNoSpace(%d)", e)
	}
}

type CmNormalMode uint32

const (
	// CmBestMode: Best mode indicates that the CMM should maximize resource usage to ensure the highest possible quality.
	CmBestMode CmNormalMode = 2
	// CmDraftMode: Draft mode indicates that the CMM should sacrifice quality, if necessary, to minimize resource requirements.
	CmDraftMode CmNormalMode = 1
	// CmNormalModeValue: This is the default setting.
	CmNormalModeValue CmNormalMode = 0
)

func (e CmNormalMode) String() string {
	switch e {
	case CmBestMode:
		return "CmBestMode"
	case CmDraftMode:
		return "CmDraftMode"
	case CmNormalModeValue:
		return "CmNormalModeValue"
	default:
		return fmt.Sprintf("CmNormalMode(%d)", e)
	}
}

type CmNumHeader uint32

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

type CmOpenReadAccess uint32

const (
	// CmAbortWriteAccess: Cancel the current write attempt.
	CmAbortWriteAccess CmOpenReadAccess = 7
	// CmBeginAccess: Begin the process of procedural access.
	CmBeginAccess CmOpenReadAccess = 8
	// CmCloseAccess: Close the profile for reading or writing.
	CmCloseAccess CmOpenReadAccess = 5
	// CmCreateNewAccess: Create a new data stream for the profile.
	CmCreateNewAccess CmOpenReadAccess = 6
	// CmEndAccess: End the process of procedural access.
	CmEndAccess           CmOpenReadAccess = 9
	CmOpenReadAccessValue CmOpenReadAccess = 1
	// CmOpenWriteAccess: Open the profile for writing.
	CmOpenWriteAccess CmOpenReadAccess = 2
	// CmReadAccess: Read the number of bytes specified by the `size` parameter.
	CmReadAccess CmOpenReadAccess = 3
	// CmWriteAccess: Write the number of bytes specified by the `size` parameter.
	CmWriteAccess CmOpenReadAccess = 4
)

func (e CmOpenReadAccess) String() string {
	switch e {
	case CmAbortWriteAccess:
		return "CmAbortWriteAccess"
	case CmBeginAccess:
		return "CmBeginAccess"
	case CmCloseAccess:
		return "CmCloseAccess"
	case CmCreateNewAccess:
		return "CmCreateNewAccess"
	case CmEndAccess:
		return "CmEndAccess"
	case CmOpenReadAccessValue:
		return "CmOpenReadAccessValue"
	case CmOpenWriteAccess:
		return "CmOpenWriteAccess"
	case CmReadAccess:
		return "CmReadAccess"
	case CmWriteAccess:
		return "CmWriteAccess"
	default:
		return fmt.Sprintf("CmOpenReadAccess(%d)", e)
	}
}

type CmOpenReadSpool uint32

const (
	// CmCloseSpool: Directs the function to complete the data transfer.
	CmCloseSpool CmOpenReadSpool = 5
	// CmOpenReadSpoolValue: Directs the function to begin the process of reading data.
	CmOpenReadSpoolValue CmOpenReadSpool = 1
	// CmOpenWriteSpool: Directs the function to begin the process of writing data.
	CmOpenWriteSpool CmOpenReadSpool = 2
	// CmReadSpool: Directs the function to read the number of bytes specified by the [CMFlattenProcPtr] function’s `size` parameter.
	CmReadSpool CmOpenReadSpool = 3
	// CmWriteSpool: Directs the function to write the number of bytes specified by the [CMFlattenProcPtr] function’s `size` parameter.
	CmWriteSpool CmOpenReadSpool = 4
)

func (e CmOpenReadSpool) String() string {
	switch e {
	case CmCloseSpool:
		return "CmCloseSpool"
	case CmOpenReadSpoolValue:
		return "CmOpenReadSpoolValue"
	case CmOpenWriteSpool:
		return "CmOpenWriteSpool"
	case CmReadSpool:
		return "CmReadSpool"
	case CmWriteSpool:
		return "CmWriteSpool"
	default:
		return fmt.Sprintf("CmOpenReadSpool(%d)", e)
	}
}

type CmOriginalProfileLocationSize uint32

const (
	CmCurrentProfileLocationSize       CmOriginalProfileLocationSize = 1032
	CmOriginalProfileLocationSizeValue CmOriginalProfileLocationSize = 72
)

func (e CmOriginalProfileLocationSize) String() string {
	switch e {
	case CmCurrentProfileLocationSize:
		return "CmCurrentProfileLocationSize"
	case CmOriginalProfileLocationSizeValue:
		return "CmOriginalProfileLocationSizeValue"
	default:
		return fmt.Sprintf("CmOriginalProfileLocationSize(%d)", e)
	}
}

type CmP uint32

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

type CmPS2CRDVMSizeTag uint32

const (
	CmMakeAndModelTag         CmPS2CRDVMSizeTag = 'm'<<24 | 'm'<<16 | 'o'<<8 | 'd' // 'mmod'
	CmNativeDisplayInfoTag    CmPS2CRDVMSizeTag = 'n'<<24 | 'd'<<16 | 'i'<<8 | 'n' // 'ndin'
	CmPS2CRDVMSizeTagValue    CmPS2CRDVMSizeTag = 'p'<<24 | 's'<<16 | 'v'<<8 | 'm' // 'psvm'
	CmProfileDescriptionMLTag CmPS2CRDVMSizeTag = 'd'<<24 | 's'<<16 | 'c'<<8 | 'm' // 'dscm'
	// CmVideoCardGammaTag: # Discussion
	CmVideoCardGammaTag CmPS2CRDVMSizeTag = 'v'<<24 | 'c'<<16 | 'g'<<8 | 't' // 'vcgt'
)

func (e CmPS2CRDVMSizeTag) String() string {
	switch e {
	case CmMakeAndModelTag:
		return "CmMakeAndModelTag"
	case CmNativeDisplayInfoTag:
		return "CmNativeDisplayInfoTag"
	case CmPS2CRDVMSizeTagValue:
		return "CmPS2CRDVMSizeTagValue"
	case CmProfileDescriptionMLTag:
		return "CmProfileDescriptionMLTag"
	case CmVideoCardGammaTag:
		return "CmVideoCardGammaTag"
	default:
		return fmt.Sprintf("CmPS2CRDVMSizeTag(%d)", e)
	}
}

type CmParametric uint32

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

type CmPerceptual uint32

const (
	// CmAbsoluteColorimetric: This approach is based on a device-independent color space in which the result is an idealized print viewed on an ideal type of paper having a large dynamic range and color gamut.
	CmAbsoluteColorimetric CmPerceptual = 3
	// CmPerceptualValue: All the colors of a given gamut can be scaled to fit within another gamut.
	CmPerceptualValue CmPerceptual = 0
	// CmRelativeColorimetric: The colors that fall within the gamuts of both devices are left unchanged.
	CmRelativeColorimetric CmPerceptual = 1
	// CmSaturation: The relative saturation of colors is maintained from gamut to gamut.
	CmSaturation CmPerceptual = 2
)

func (e CmPerceptual) String() string {
	switch e {
	case CmAbsoluteColorimetric:
		return "CmAbsoluteColorimetric"
	case CmPerceptualValue:
		return "CmPerceptualValue"
	case CmRelativeColorimetric:
		return "CmRelativeColorimetric"
	case CmSaturation:
		return "CmSaturation"
	default:
		return fmt.Sprintf("CmPerceptual(%d)", e)
	}
}

type CmProfileIterateData uint32

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

type CmProfileMajorVersionMask int32

const (
	CmCurrentProfileMajorVersion   CmProfileMajorVersionMask = 0x2000000
	CmProfileMajorVersionMaskValue CmProfileMajorVersionMask = -16777216
)

func (e CmProfileMajorVersionMask) String() string {
	switch e {
	case CmCurrentProfileMajorVersion:
		return "CmCurrentProfileMajorVersion"
	case CmProfileMajorVersionMaskValue:
		return "CmProfileMajorVersionMaskValue"
	default:
		return fmt.Sprintf("CmProfileMajorVersionMask(%d)", e)
	}
}

type CmPrtrDefaultScreens uint32

const (
	// CmLinesPer: Lines per unit; can have an associated value of `0` for lines per centimeter or `1` for lines per inch.
	CmLinesPer CmPrtrDefaultScreens = 1
	// CmPrtrDefaultScreensValue: Use printer default screens; can have an associated value of `0` for `false` or `1` for `true`.
	CmPrtrDefaultScreensValue CmPrtrDefaultScreens = 0
)

func (e CmPrtrDefaultScreens) String() string {
	switch e {
	case CmLinesPer:
		return "CmLinesPer"
	case CmPrtrDefaultScreensValue:
		return "CmPrtrDefaultScreensValue"
	default:
		return fmt.Sprintf("CmPrtrDefaultScreens(%d)", e)
	}
}

type CmReflective uint32

const (
	// CmGlossy: If the bit 1 of the associated mask is `0` then glossy; if `1` then matte.
	CmGlossy CmReflective = 1
	// CmReflectiveValue: If the bit 0 of the associated mask is `0` then reflective media; if `1` then transparency media.
	CmReflectiveValue CmReflective = 0
)

func (e CmReflective) String() string {
	switch e {
	case CmGlossy:
		return "CmGlossy"
	case CmReflectiveValue:
		return "CmReflectiveValue"
	default:
		return fmt.Sprintf("CmReflective(%d)", e)
	}
}

type CmReflectiveTransparentMask uint32

const (
	// CmGlossyMatteMask: # Discussion
	CmGlossyMatteMask CmReflectiveTransparentMask = 0x2
	// CmReflectiveTransparentMaskValue: # Discussion
	CmReflectiveTransparentMaskValue CmReflectiveTransparentMask = 0x1
)

func (e CmReflectiveTransparentMask) String() string {
	switch e {
	case CmGlossyMatteMask:
		return "CmGlossyMatteMask"
	case CmReflectiveTransparentMaskValue:
		return "CmReflectiveTransparentMaskValue"
	default:
		return fmt.Sprintf("CmReflectiveTransparentMask(%d)", e)
	}
}

type CmSRGB16Channel uint32

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

type CmScannerDeviceClass uint32

const (
	CmCameraDeviceClass       CmScannerDeviceClass = 'c'<<24 | 'm'<<16 | 'r'<<8 | 'a' // 'cmra'
	CmDisplayDeviceClass      CmScannerDeviceClass = 'm'<<24 | 'n'<<16 | 't'<<8 | 'r' // 'mntr'
	CmPrinterDeviceClass      CmScannerDeviceClass = 'p'<<24 | 'r'<<16 | 't'<<8 | 'r' // 'prtr'
	CmProofDeviceClass        CmScannerDeviceClass = 'p'<<24 | 'r'<<16 | 'u'<<8 | 'f' // 'pruf'
	CmScannerDeviceClassValue CmScannerDeviceClass = 's'<<24 | 'c'<<16 | 'n'<<8 | 'r' // 'scnr'
)

func (e CmScannerDeviceClass) String() string {
	switch e {
	case CmCameraDeviceClass:
		return "CmCameraDeviceClass"
	case CmDisplayDeviceClass:
		return "CmDisplayDeviceClass"
	case CmPrinterDeviceClass:
		return "CmPrinterDeviceClass"
	case CmProofDeviceClass:
		return "CmProofDeviceClass"
	case CmScannerDeviceClassValue:
		return "CmScannerDeviceClassValue"
	default:
		return fmt.Sprintf("CmScannerDeviceClass(%d)", e)
	}
}

type CmSigCrdInfoType uint32

const (
	CmSigCrdInfoTypeValue        CmSigCrdInfoType = 'c'<<24 | 'r'<<16 | 'd'<<8 | 'i' // 'crdi'
	CmSigCurveType               CmSigCrdInfoType = 'c'<<24 | 'u'<<16 | 'r'<<8 | 'v' // 'curv'
	CmSigDataType                CmSigCrdInfoType = 'd'<<24 | 'a'<<16 | 't'<<8 | 'a' // 'data'
	CmSigDateTimeType            CmSigCrdInfoType = 'd'<<24 | 't'<<16 | 'i'<<8 | 'm' // 'dtim'
	CmSigLut16Type               CmSigCrdInfoType = 'm'<<24 | 'f'<<16 | 't'<<8 | '2' // 'mft2'
	CmSigLut8Type                CmSigCrdInfoType = 'm'<<24 | 'f'<<16 | 't'<<8 | '1' // 'mft1'
	CmSigMeasurementType         CmSigCrdInfoType = 'm'<<24 | 'e'<<16 | 'a'<<8 | 's' // 'meas'
	CmSigMultiFunctA2BType       CmSigCrdInfoType = 'm'<<24 | 'A'<<16 | 'B'<<8 | ' ' // 'mAB '
	CmSigMultiFunctB2AType       CmSigCrdInfoType = 'm'<<24 | 'B'<<16 | 'A'<<8 | ' ' // 'mBA '
	CmSigNamedColor2Type         CmSigCrdInfoType = 'n'<<24 | 'c'<<16 | 'l'<<8 | '2' // 'ncl2'
	CmSigNamedColorType          CmSigCrdInfoType = 'n'<<24 | 'c'<<16 | 'o'<<8 | 'l' // 'ncol'
	CmSigParametricCurveType     CmSigCrdInfoType = 'p'<<24 | 'a'<<16 | 'r'<<8 | 'a' // 'para'
	CmSigProfileDescriptionType  CmSigCrdInfoType = 'd'<<24 | 'e'<<16 | 's'<<8 | 'c' // 'desc'
	CmSigProfileSequenceDescType CmSigCrdInfoType = 'p'<<24 | 's'<<16 | 'e'<<8 | 'q' // 'pseq'
	CmSigS15Fixed16Type          CmSigCrdInfoType = 's'<<24 | 'f'<<16 | '3'<<8 | '2' // 'sf32'
	CmSigScreeningType           CmSigCrdInfoType = 's'<<24 | 'c'<<16 | 'r'<<8 | 'n' // 'scrn'
	CmSigSignatureType           CmSigCrdInfoType = 's'<<24 | 'i'<<16 | 'g'<<8 | ' ' // 'sig '
	CmSigTextType                CmSigCrdInfoType = 't'<<24 | 'e'<<16 | 'x'<<8 | 't' // 'text'
	CmSigU16Fixed16Type          CmSigCrdInfoType = 'u'<<24 | 'f'<<16 | '3'<<8 | '2' // 'uf32'
	CmSigU1Fixed15Type           CmSigCrdInfoType = 'u'<<24 | 'f'<<16 | '1'<<8 | '6' // 'uf16'
	CmSigUInt16Type              CmSigCrdInfoType = 'u'<<24 | 'i'<<16 | '1'<<8 | '6' // 'ui16'
	CmSigUInt32Type              CmSigCrdInfoType = 'u'<<24 | 'i'<<16 | '3'<<8 | '2' // 'ui32'
	CmSigUInt64Type              CmSigCrdInfoType = 'u'<<24 | 'i'<<16 | '6'<<8 | '4' // 'ui64'
	CmSigUInt8Type               CmSigCrdInfoType = 'u'<<24 | 'i'<<16 | '0'<<8 | '8' // 'ui08'
	CmSigUcrBgType               CmSigCrdInfoType = 'b'<<24 | 'f'<<16 | 'd'<<8 | ' ' // 'bfd '
	CmSigUnicodeTextType         CmSigCrdInfoType = 'u'<<24 | 't'<<16 | 'x'<<8 | 't' // 'utxt'
	CmSigViewingConditionsType   CmSigCrdInfoType = 'v'<<24 | 'i'<<16 | 'e'<<8 | 'w' // 'view'
	CmSigXYZType                 CmSigCrdInfoType = 'X'<<24 | 'Y'<<16 | 'Z'<<8 | ' ' // 'XYZ '
)

func (e CmSigCrdInfoType) String() string {
	switch e {
	case CmSigCrdInfoTypeValue:
		return "CmSigCrdInfoTypeValue"
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
	case CmSigMeasurementType:
		return "CmSigMeasurementType"
	case CmSigMultiFunctA2BType:
		return "CmSigMultiFunctA2BType"
	case CmSigMultiFunctB2AType:
		return "CmSigMultiFunctB2AType"
	case CmSigNamedColor2Type:
		return "CmSigNamedColor2Type"
	case CmSigNamedColorType:
		return "CmSigNamedColorType"
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
	case CmSigViewingConditionsType:
		return "CmSigViewingConditionsType"
	case CmSigXYZType:
		return "CmSigXYZType"
	default:
		return fmt.Sprintf("CmSigCrdInfoType(%d)", e)
	}
}

type CmSigPS2CRDVMSizeType uint32

const (
	CmSigMakeAndModelType          CmSigPS2CRDVMSizeType = 'm'<<24 | 'm'<<16 | 'o'<<8 | 'd' // 'mmod'
	CmSigMultiLocalizedUniCodeType CmSigPS2CRDVMSizeType = 'm'<<24 | 'l'<<16 | 'u'<<8 | 'c' // 'mluc'
	CmSigNativeDisplayInfoType     CmSigPS2CRDVMSizeType = 'n'<<24 | 'd'<<16 | 'i'<<8 | 'n' // 'ndin'
	CmSigPS2CRDVMSizeTypeValue     CmSigPS2CRDVMSizeType = 'p'<<24 | 's'<<16 | 'v'<<8 | 'm' // 'psvm'
	// CmSigVideoCardGammaType: # Discussion
	CmSigVideoCardGammaType CmSigPS2CRDVMSizeType = 'v'<<24 | 'c'<<16 | 'g'<<8 | 't' // 'vcgt'
)

func (e CmSigPS2CRDVMSizeType) String() string {
	switch e {
	case CmSigMakeAndModelType:
		return "CmSigMakeAndModelType"
	case CmSigMultiLocalizedUniCodeType:
		return "CmSigMultiLocalizedUniCodeType"
	case CmSigNativeDisplayInfoType:
		return "CmSigNativeDisplayInfoType"
	case CmSigPS2CRDVMSizeTypeValue:
		return "CmSigPS2CRDVMSizeTypeValue"
	case CmSigVideoCardGammaType:
		return "CmSigVideoCardGammaType"
	default:
		return fmt.Sprintf("CmSigPS2CRDVMSizeType(%d)", e)
	}
}

type CmSpotFunction uint32

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

type CmStdobsUnknown uint32

const (
	CmStdobs1931TwoDegrees CmStdobsUnknown = 0x1
	CmStdobs1964TenDegrees CmStdobsUnknown = 0x2
	CmStdobsUnknownValue   CmStdobsUnknown = 0
)

func (e CmStdobsUnknown) String() string {
	switch e {
	case CmStdobs1931TwoDegrees:
		return "CmStdobs1931TwoDegrees"
	case CmStdobs1964TenDegrees:
		return "CmStdobs1964TenDegrees"
	case CmStdobsUnknownValue:
		return "CmStdobsUnknownValue"
	default:
		return fmt.Sprintf("CmStdobsUnknown(%d)", e)
	}
}

type CmTechnology uint32

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

type CmUseDefaultChromaticAdaptation uint32

const (
	CmBradfordChromaticAdaptation        CmUseDefaultChromaticAdaptation = 3
	CmLinearChromaticAdaptation          CmUseDefaultChromaticAdaptation = 1
	CmUseDefaultChromaticAdaptationValue CmUseDefaultChromaticAdaptation = 0
	CmVonKriesChromaticAdaptation        CmUseDefaultChromaticAdaptation = 2
)

func (e CmUseDefaultChromaticAdaptation) String() string {
	switch e {
	case CmBradfordChromaticAdaptation:
		return "CmBradfordChromaticAdaptation"
	case CmLinearChromaticAdaptation:
		return "CmLinearChromaticAdaptation"
	case CmUseDefaultChromaticAdaptationValue:
		return "CmUseDefaultChromaticAdaptationValue"
	case CmVonKriesChromaticAdaptation:
		return "CmVonKriesChromaticAdaptation"
	default:
		return fmt.Sprintf("CmUseDefaultChromaticAdaptation(%d)", e)
	}
}

type CmVideoCardGamma uint32

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

type CmXYZData uint32

const (
	Cm10CLRData CmXYZData = 'A'<<24 | 'C'<<16 | 'L'<<8 | 'R' // 'ACLR'
	Cm11CLRData CmXYZData = 'B'<<24 | 'C'<<16 | 'L'<<8 | 'R' // 'BCLR'
	Cm12CLRData CmXYZData = 'C'<<24 | 'C'<<16 | 'L'<<8 | 'R' // 'CCLR'
	Cm13CLRData CmXYZData = 'D'<<24 | 'C'<<16 | 'L'<<8 | 'R' // 'DCLR'
	Cm14CLRData CmXYZData = 'E'<<24 | 'C'<<16 | 'L'<<8 | 'R' // 'ECLR'
	Cm15CLRData CmXYZData = 'F'<<24 | 'C'<<16 | 'L'<<8 | 'R' // 'FCLR'
	Cm3CLRData  CmXYZData = '3'<<24 | 'C'<<16 | 'L'<<8 | 'R' // '3CLR'
	Cm4CLRData  CmXYZData = '4'<<24 | 'C'<<16 | 'L'<<8 | 'R' // '4CLR'
	Cm5CLRData  CmXYZData = '5'<<24 | 'C'<<16 | 'L'<<8 | 'R' // '5CLR'
	Cm6CLRData  CmXYZData = '6'<<24 | 'C'<<16 | 'L'<<8 | 'R' // '6CLR'
	Cm7CLRData  CmXYZData = '7'<<24 | 'C'<<16 | 'L'<<8 | 'R' // '7CLR'
	Cm8CLRData  CmXYZData = '8'<<24 | 'C'<<16 | 'L'<<8 | 'R' // '8CLR'
	Cm9CLRData  CmXYZData = '9'<<24 | 'C'<<16 | 'L'<<8 | 'R' // '9CLR'
	// CmCMYData: The CMY data color space.
	CmCMYData CmXYZData = 'C'<<24 | 'M'<<16 | 'Y'<<8 | ' ' // 'CMY '
	// CmCMYKData: The CMYK data color space.
	CmCMYKData CmXYZData = 'C'<<24 | 'M'<<16 | 'Y'<<8 | 'K' // 'CMYK'
	// CmGrayData: The Gray data color space.
	CmGrayData CmXYZData = 'G'<<24 | 'R'<<16 | 'A'<<8 | 'Y' // 'GRAY'
	// CmHLSData: The HLS data color space.
	CmHLSData CmXYZData = 'H'<<24 | 'L'<<16 | 'S'<<8 | ' ' // 'HLS '
	// CmHSVData: The HSV data color space.
	CmHSVData CmXYZData = 'H'<<24 | 'S'<<16 | 'V'<<8 | ' ' // 'HSV '
	// CmLabData: The L*a*b* data color space.
	CmLabData CmXYZData = 'L'<<24 | 'a'<<16 | 'b'<<8 | ' ' // 'Lab '
	// CmLuvData: The L*u*v* data color space.
	CmLuvData CmXYZData = 'L'<<24 | 'u'<<16 | 'v'<<8 | ' ' // 'Luv '
	// CmMCH5Data: The five-channel multichannel (HiFi) data color space.
	CmMCH5Data CmXYZData = 'M'<<24 | 'C'<<16 | 'H'<<8 | '5' // 'MCH5'
	// CmMCH6Data: The six-channel multichannel (HiFi) data color space.
	CmMCH6Data CmXYZData = 'M'<<24 | 'C'<<16 | 'H'<<8 | '6' // 'MCH6'
	// CmMCH7Data: The seven-channel multichannel (HiFi) data color space.
	CmMCH7Data CmXYZData = 'M'<<24 | 'C'<<16 | 'H'<<8 | '7' // 'MCH7'
	// CmMCH8Data: The eight-channel multichannel (HiFi) data color space.
	CmMCH8Data  CmXYZData = 'M'<<24 | 'C'<<16 | 'H'<<8 | '8' // 'MCH8'
	CmNamedData CmXYZData = 'N'<<24 | 'A'<<16 | 'M'<<8 | 'E' // 'NAME'
	// CmRGBData: The RGB data color space.
	CmRGBData  CmXYZData = 'R'<<24 | 'G'<<16 | 'B'<<8 | ' ' // 'RGB '
	CmSRGBData CmXYZData = 's'<<24 | 'R'<<16 | 'G'<<8 | 'B' // 'sRGB'
	// CmXYZDataValue: The XYZ data color space.
	CmXYZDataValue CmXYZData = 'X'<<24 | 'Y'<<16 | 'Z'<<8 | ' ' // 'XYZ '
	CmYCbCrData    CmXYZData = 'Y'<<24 | 'C'<<16 | 'b'<<8 | 'r' // 'YCbr'
	// CmYxyData: The Yxy data color space.
	CmYxyData CmXYZData = 'Y'<<24 | 'x'<<16 | 'y'<<8 | ' ' // 'Yxy '
)

func (e CmXYZData) String() string {
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
	case Cm3CLRData:
		return "Cm3CLRData"
	case Cm4CLRData:
		return "Cm4CLRData"
	case Cm5CLRData:
		return "Cm5CLRData"
	case Cm6CLRData:
		return "Cm6CLRData"
	case Cm7CLRData:
		return "Cm7CLRData"
	case Cm8CLRData:
		return "Cm8CLRData"
	case Cm9CLRData:
		return "Cm9CLRData"
	case CmCMYData:
		return "CmCMYData"
	case CmCMYKData:
		return "CmCMYKData"
	case CmGrayData:
		return "CmGrayData"
	case CmHLSData:
		return "CmHLSData"
	case CmHSVData:
		return "CmHSVData"
	case CmLabData:
		return "CmLabData"
	case CmLuvData:
		return "CmLuvData"
	case CmMCH5Data:
		return "CmMCH5Data"
	case CmMCH6Data:
		return "CmMCH6Data"
	case CmMCH7Data:
		return "CmMCH7Data"
	case CmMCH8Data:
		return "CmMCH8Data"
	case CmNamedData:
		return "CmNamedData"
	case CmRGBData:
		return "CmRGBData"
	case CmSRGBData:
		return "CmSRGBData"
	case CmXYZDataValue:
		return "CmXYZDataValue"
	case CmYCbCrData:
		return "CmYCbCrData"
	case CmYxyData:
		return "CmYxyData"
	default:
		return fmt.Sprintf("CmXYZData(%d)", e)
	}
}

type CsMax uint32

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

type Extended uint32

const (
	ExtendedBlock    Extended = 0x4c43
	ExtendedBlockLen Extended = 40
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

type IcNoPerm uint32

const (
	IcNoPermValue   IcNoPerm = 0
	IcReadOnlyPerm  IcNoPerm = 1
	IcReadWritePerm IcNoPerm = 2
)

func (e IcNoPerm) String() string {
	switch e {
	case IcNoPermValue:
		return "IcNoPermValue"
	case IcReadOnlyPerm:
		return "IcReadOnlyPerm"
	case IcReadWritePerm:
		return "IcReadWritePerm"
	default:
		return fmt.Sprintf("IcNoPerm(%d)", e)
	}
}

type IcPrefNotFoundErr int32

const (
	IcConfigInappropriateErr IcPrefNotFoundErr = -675
	IcConfigNotFoundErr      IcPrefNotFoundErr = -674
	IcInternalErr            IcPrefNotFoundErr = -669
	IcNoMoreWritersErr       IcPrefNotFoundErr = -671
	IcNoURLErr               IcPrefNotFoundErr = -673
	IcNothingToOverrideErr   IcPrefNotFoundErr = -672
	IcPermErr                IcPrefNotFoundErr = -667
	IcPrefDataErr            IcPrefNotFoundErr = -668
	IcPrefNotFoundErrValue   IcPrefNotFoundErr = -666
	IcProfileNotFoundErr     IcPrefNotFoundErr = -676
	IcTooManyProfilesErr     IcPrefNotFoundErr = -677
	IcTruncatedErr           IcPrefNotFoundErr = -670
)

func (e IcPrefNotFoundErr) String() string {
	switch e {
	case IcConfigInappropriateErr:
		return "IcConfigInappropriateErr"
	case IcConfigNotFoundErr:
		return "IcConfigNotFoundErr"
	case IcInternalErr:
		return "IcInternalErr"
	case IcNoMoreWritersErr:
		return "IcNoMoreWritersErr"
	case IcNoURLErr:
		return "IcNoURLErr"
	case IcNothingToOverrideErr:
		return "IcNothingToOverrideErr"
	case IcPermErr:
		return "IcPermErr"
	case IcPrefDataErr:
		return "IcPrefDataErr"
	case IcPrefNotFoundErrValue:
		return "IcPrefNotFoundErrValue"
	case IcProfileNotFoundErr:
		return "IcProfileNotFoundErr"
	case IcTooManyProfilesErr:
		return "IcTooManyProfilesErr"
	case IcTruncatedErr:
		return "IcTruncatedErr"
	default:
		return fmt.Sprintf("IcPrefNotFoundErr(%d)", e)
	}
}

type InitDev uint32

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

type K1MonochromePixelFormat uint32

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

type KATSCubicCurveType uint32

const (
	KATSCubicCurveTypeValue KATSCubicCurveType = 0x1
	KATSOtherCurveType      KATSCubicCurveType = 0x3
	KATSQuadCurveType       KATSCubicCurveType = 0x2
)

func (e KATSCubicCurveType) String() string {
	switch e {
	case KATSCubicCurveTypeValue:
		return "KATSCubicCurveTypeValue"
	case KATSOtherCurveType:
		return "KATSOtherCurveType"
	case KATSQuadCurveType:
		return "KATSQuadCurveType"
	default:
		return fmt.Sprintf("KATSCubicCurveType(%d)", e)
	}
}

type KATSDeleted uint32

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

type KATSFlatDataUstl uint32

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

type KATSFlattenedFontSpecifierRawName uint32

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

type KATSFontAutoActivation uint32

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

type KATSFontContext uint32

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

type KATSFontFilterCurrent uint32

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

type KATSFontFormat uint32

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

type KATSGenerationUnspecified uint32

const (
	KATSFontContainerRefUnspecified KATSGenerationUnspecified = 0
	KATSFontFamilyRefUnspecified    KATSGenerationUnspecified = 0
	KATSFontRefUnspecified          KATSGenerationUnspecified = 0
	KATSGenerationUnspecifiedValue  KATSGenerationUnspecified = 0
)

func (e KATSGenerationUnspecified) String() string {
	switch e {
	case KATSFontContainerRefUnspecified:
		return "KATSFontContainerRefUnspecified"
	default:
		return fmt.Sprintf("KATSGenerationUnspecified(%d)", e)
	}
}

type KATSGlyphInfo uint32

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

type KATSItalicQDSkew uint32

const (
	KATSBoldQDStretch     KATSItalicQDSkew = 98304
	KATSItalicQDSkewValue KATSItalicQDSkew = 16384
	KATSRadiansFactor     KATSItalicQDSkew = 1144
)

func (e KATSItalicQDSkew) String() string {
	switch e {
	case KATSBoldQDStretch:
		return "KATSBoldQDStretch"
	case KATSItalicQDSkewValue:
		return "KATSItalicQDSkewValue"
	case KATSRadiansFactor:
		return "KATSRadiansFactor"
	default:
		return fmt.Sprintf("KATSItalicQDSkew(%d)", e)
	}
}

type KATSIterationCompleted int32

const (
	KATSInvalidFontAccess          KATSIterationCompleted = -982
	KATSInvalidFontContainerAccess KATSIterationCompleted = -985
	KATSInvalidFontFamilyAccess    KATSIterationCompleted = -981
	KATSInvalidFontTableAccess     KATSIterationCompleted = -984
	KATSInvalidGlyphAccess         KATSIterationCompleted = -986
	KATSIterationCompletedValue    KATSIterationCompleted = -980
	KATSIterationScopeModified     KATSIterationCompleted = -983
)

func (e KATSIterationCompleted) String() string {
	switch e {
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
	case KATSIterationCompletedValue:
		return "KATSIterationCompletedValue"
	case KATSIterationScopeModified:
		return "KATSIterationScopeModified"
	default:
		return fmt.Sprintf("KATSIterationCompleted(%d)", e)
	}
}

type KATSLine uint32

const (
	KATSLineAppleReserved                 KATSLine = 0xfce00000
	KATSLineApplyAntiAliasing             KATSLine = 0x800
	KATSLineBreakToNearestCharacter       KATSLine = 0x2000000
	KATSLineDisableAllBaselineAdjustments KATSLine = 0x80000
	KATSLineDisableAllGlyphMorphing       KATSLine = 0x20000
	KATSLineDisableAllJustification       KATSLine = 0x10000
	KATSLineDisableAllKerningAdjustments  KATSLine = 0x40000
	KATSLineDisableAllLayoutOperations    KATSLine = 2031616
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
	case KATSLineDisableAllLayoutOperations:
		return "KATSLineDisableAllLayoutOperations"
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

type KATSOptionFlagsActivateDisabled uint32

const (
	KATSOptionFlagsActivateDisabledValue KATSOptionFlagsActivateDisabled = 32
	KATSOptionFlagsDoNotNotify           KATSOptionFlagsActivateDisabled = 128
	KATSOptionFlagsProcessSubdirectories KATSOptionFlagsActivateDisabled = 64
	KATSOptionFlagsRecordPersistently    KATSOptionFlagsActivateDisabled = 262144
)

func (e KATSOptionFlagsActivateDisabled) String() string {
	switch e {
	case KATSOptionFlagsActivateDisabledValue:
		return "KATSOptionFlagsActivateDisabledValue"
	case KATSOptionFlagsDoNotNotify:
		return "KATSOptionFlagsDoNotNotify"
	case KATSOptionFlagsProcessSubdirectories:
		return "KATSOptionFlagsProcessSubdirectories"
	case KATSOptionFlagsRecordPersistently:
		return "KATSOptionFlagsRecordPersistently"
	default:
		return fmt.Sprintf("KATSOptionFlagsActivateDisabled(%d)", e)
	}
}

type KATSOptionFlagsDefault uint32

const (
	KATSOptionFlagsComposeFontPostScriptName KATSOptionFlagsDefault = 1
	KATSOptionFlagsDefaultValue              KATSOptionFlagsDefault = 0
	KATSOptionFlagsUseDataFork               KATSOptionFlagsDefault = 768
	KATSOptionFlagsUseDataForkAsResourceFork KATSOptionFlagsDefault = 256
	KATSOptionFlagsUseResourceFork           KATSOptionFlagsDefault = 512
)

func (e KATSOptionFlagsDefault) String() string {
	switch e {
	case KATSOptionFlagsComposeFontPostScriptName:
		return "KATSOptionFlagsComposeFontPostScriptName"
	case KATSOptionFlagsDefaultValue:
		return "KATSOptionFlagsDefaultValue"
	case KATSOptionFlagsUseDataFork:
		return "KATSOptionFlagsUseDataFork"
	case KATSOptionFlagsUseDataForkAsResourceFork:
		return "KATSOptionFlagsUseDataForkAsResourceFork"
	case KATSOptionFlagsUseResourceFork:
		return "KATSOptionFlagsUseResourceFork"
	default:
		return fmt.Sprintf("KATSOptionFlagsDefault(%d)", e)
	}
}

type KATSOptionFlagsIterateByPrecedenceMask uint32

const (
	KATSOptionFlagsDefaultScope                 KATSOptionFlagsIterateByPrecedenceMask = 0
	KATSOptionFlagsIncludeDisabledMask          KATSOptionFlagsIterateByPrecedenceMask = 128
	KATSOptionFlagsIterateByPrecedenceMaskValue KATSOptionFlagsIterateByPrecedenceMask = 32
	KATSOptionFlagsIterationScopeMask           KATSOptionFlagsIterateByPrecedenceMask = 28672
	KATSOptionFlagsRestrictedScope              KATSOptionFlagsIterateByPrecedenceMask = 8192
	KATSOptionFlagsUnRestrictedScope            KATSOptionFlagsIterateByPrecedenceMask = 4096
)

func (e KATSOptionFlagsIterateByPrecedenceMask) String() string {
	switch e {
	case KATSOptionFlagsDefaultScope:
		return "KATSOptionFlagsDefaultScope"
	case KATSOptionFlagsIncludeDisabledMask:
		return "KATSOptionFlagsIncludeDisabledMask"
	case KATSOptionFlagsIterateByPrecedenceMaskValue:
		return "KATSOptionFlagsIterateByPrecedenceMaskValue"
	case KATSOptionFlagsIterationScopeMask:
		return "KATSOptionFlagsIterationScopeMask"
	case KATSOptionFlagsRestrictedScope:
		return "KATSOptionFlagsRestrictedScope"
	case KATSOptionFlagsUnRestrictedScope:
		return "KATSOptionFlagsUnRestrictedScope"
	default:
		return fmt.Sprintf("KATSOptionFlagsIterateByPrecedenceMask(%d)", e)
	}
}

type KATSStyle uint32

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

type KATSUBackground uint32

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

type KATSUBy uint32

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

type KATSUDataStreamUnicodeStyled uint32

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

type KATSUDefaultFontFallbacks uint32

const (
	KATSUDefaultFontFallbacksValue    KATSUDefaultFontFallbacks = 0
	KATSULastResortOnlyFallback       KATSUDefaultFontFallbacks = 1
	KATSUSequentialFallbacksExclusive KATSUDefaultFontFallbacks = 3
	KATSUSequentialFallbacksPreferred KATSUDefaultFontFallbacks = 2
)

func (e KATSUDefaultFontFallbacks) String() string {
	switch e {
	case KATSUDefaultFontFallbacksValue:
		return "KATSUDefaultFontFallbacksValue"
	case KATSULastResortOnlyFallback:
		return "KATSULastResortOnlyFallback"
	case KATSUSequentialFallbacksExclusive:
		return "KATSUSequentialFallbacksExclusive"
	case KATSUSequentialFallbacksPreferred:
		return "KATSUSequentialFallbacksPreferred"
	default:
		return fmt.Sprintf("KATSUDefaultFontFallbacks(%d)", e)
	}
}

type KATSUDirectData uint32

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

type KATSUFlattenOptionNoOptions uint32

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

type KATSUFromTextBeginning uint32

const (
	KATSUFromFollowingLayout    KATSUFromTextBeginning = 0xfffffffd
	KATSUFromPreviousLayout     KATSUFromTextBeginning = 0xfffffffe
	KATSUFromTextBeginningValue KATSUFromTextBeginning = 0xffffffff
	KATSUToTextEnd              KATSUFromTextBeginning = 0xffffffff
)

func (e KATSUFromTextBeginning) String() string {
	switch e {
	case KATSUFromFollowingLayout:
		return "KATSUFromFollowingLayout"
	case KATSUFromPreviousLayout:
		return "KATSUFromPreviousLayout"
	case KATSUFromTextBeginningValue:
		return "KATSUFromTextBeginningValue"
	default:
		return fmt.Sprintf("KATSUFromTextBeginning(%d)", e)
	}
}

type KATSUInvalidFontI uint32

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

type KATSULayoutOperation uint32

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

type KATSULayoutOperationCallbackStatus uint32

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

type KATSULeftTab uint32

const (
	KATSUCenterTab      KATSULeftTab = 1
	KATSUDecimalTab     KATSULeftTab = 3
	KATSULeftTabValue   KATSULeftTab = 0
	KATSUNumberTabTypes KATSULeftTab = 4
	KATSURightTab       KATSULeftTab = 2
)

func (e KATSULeftTab) String() string {
	switch e {
	case KATSUCenterTab:
		return "KATSUCenterTab"
	case KATSUDecimalTab:
		return "KATSUDecimalTab"
	case KATSULeftTabValue:
		return "KATSULeftTabValue"
	case KATSUNumberTabTypes:
		return "KATSUNumberTabTypes"
	case KATSURightTab:
		return "KATSURightTab"
	default:
		return fmt.Sprintf("KATSULeftTab(%d)", e)
	}
}

type KATSULeftToRightBaseDirection uint32

const (
	KATSULeftToRightBaseDirectionValue KATSULeftToRightBaseDirection = 0
	KATSURightToLeftBaseDirection      KATSULeftToRightBaseDirection = 1
)

func (e KATSULeftToRightBaseDirection) String() string {
	switch e {
	case KATSULeftToRightBaseDirectionValue:
		return "KATSULeftToRightBaseDirectionValue"
	case KATSURightToLeftBaseDirection:
		return "KATSURightToLeftBaseDirection"
	default:
		return fmt.Sprintf("KATSULeftToRightBaseDirection(%d)", e)
	}
}

type KATSULineWidthTag uint32

const (
	KATSUAfterWithStreamShiftTag          KATSULineWidthTag = 268
	KATSUAscentTag                        KATSULineWidthTag = 284
	KATSUBaselineClassTag                 KATSULineWidthTag = 274
	KATSUBeforeWithStreamShiftTag         KATSULineWidthTag = 267
	KATSUCGContextTag                     KATSULineWidthTag = 32767
	KATSUColorTag                         KATSULineWidthTag = 263
	KATSUCrossStreamShiftTag              KATSULineWidthTag = 269
	KATSUDecompositionFactorTag           KATSULineWidthTag = 273
	KATSUDescentTag                       KATSULineWidthTag = 285
	KATSUFontMatrixTag                    KATSULineWidthTag = 289
	KATSUFontTag                          KATSULineWidthTag = 261
	KATSUForceHangingTag                  KATSULineWidthTag = 280
	KATSUGlyphSelectorTag                 KATSULineWidthTag = 287
	KATSUHangingInhibitFactorTag          KATSULineWidthTag = 271
	KATSUImposeWidthTag                   KATSULineWidthTag = 266
	KATSUKerningInhibitFactorTag          KATSULineWidthTag = 272
	KATSULangRegionTag                    KATSULineWidthTag = 264
	KATSULanguageTag                      KATSULineWidthTag = 264
	KATSULayoutOperationOverrideTag       KATSULineWidthTag = 15
	KATSULeadingTag                       KATSULineWidthTag = 286
	KATSULineAscentTag                    KATSULineWidthTag = 8
	KATSULineBaselineValuesTag            KATSULineWidthTag = 6
	KATSULineDecimalTabCharacterTag       KATSULineWidthTag = 14
	KATSULineDescentTag                   KATSULineWidthTag = 9
	KATSULineDirectionTag                 KATSULineWidthTag = 3
	KATSULineFlushFactorTag               KATSULineWidthTag = 5
	KATSULineFontFallbacksTag             KATSULineWidthTag = 13
	KATSULineHighlightCGColorTag          KATSULineWidthTag = 17
	KATSULineJustificationFactorTag       KATSULineWidthTag = 4
	KATSULineLangRegionTag                KATSULineWidthTag = 10
	KATSULineLanguageTag                  KATSULineWidthTag = 10
	KATSULineLayoutOptionsTag             KATSULineWidthTag = 7
	KATSULineRotationTag                  KATSULineWidthTag = 2
	KATSULineTextLocatorTag               KATSULineWidthTag = 11
	KATSULineTruncationTag                KATSULineWidthTag = 12
	KATSULineWidthTagValue                KATSULineWidthTag = 1
	KATSUMaxATSUITagValue                 KATSULineWidthTag = 65535
	KATSUMaxLineTag                       KATSULineWidthTag = 18
	KATSUMaxStyleTag                      KATSULineWidthTag = 299
	KATSUNoCaretAngleTag                  KATSULineWidthTag = 277
	KATSUNoLigatureSplitTag               KATSULineWidthTag = 276
	KATSUNoOpticalAlignmentTag            KATSULineWidthTag = 279
	KATSUNoSpecialJustificationTag        KATSULineWidthTag = 281
	KATSUPriorityJustOverrideTag          KATSULineWidthTag = 275
	KATSUQDBoldfaceTag                    KATSULineWidthTag = 256
	KATSUQDCondensedTag                   KATSULineWidthTag = 259
	KATSUQDExtendedTag                    KATSULineWidthTag = 260
	KATSUQDItalicTag                      KATSULineWidthTag = 257
	KATSUQDUnderlineTag                   KATSULineWidthTag = 258
	KATSURGBAlphaColorTag                 KATSULineWidthTag = 288
	KATSUSizeTag                          KATSULineWidthTag = 262
	KATSUStyleDropShadowBlurOptionTag     KATSULineWidthTag = 296
	KATSUStyleDropShadowColorOptionTag    KATSULineWidthTag = 298
	KATSUStyleDropShadowOffsetOptionTag   KATSULineWidthTag = 297
	KATSUStyleDropShadowTag               KATSULineWidthTag = 295
	KATSUStyleRenderingOptionsTag         KATSULineWidthTag = 283
	KATSUStyleStrikeThroughColorOptionTag KATSULineWidthTag = 294
	KATSUStyleStrikeThroughCountOptionTag KATSULineWidthTag = 293
	KATSUStyleStrikeThroughTag            KATSULineWidthTag = 292
	KATSUStyleTextLocatorTag              KATSULineWidthTag = 282
	KATSUStyleUnderlineColorOptionTag     KATSULineWidthTag = 291
	KATSUStyleUnderlineCountOptionTag     KATSULineWidthTag = 290
	KATSUSuppressCrossKerningTag          KATSULineWidthTag = 278
	KATSUTrackingTag                      KATSULineWidthTag = 270
	KATSUVerticalCharacterTag             KATSULineWidthTag = 265
)

func (e KATSULineWidthTag) String() string {
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
	case KATSUColorTag:
		return "KATSUColorTag"
	case KATSUCrossStreamShiftTag:
		return "KATSUCrossStreamShiftTag"
	case KATSUDecompositionFactorTag:
		return "KATSUDecompositionFactorTag"
	case KATSUDescentTag:
		return "KATSUDescentTag"
	case KATSUFontMatrixTag:
		return "KATSUFontMatrixTag"
	case KATSUFontTag:
		return "KATSUFontTag"
	case KATSUForceHangingTag:
		return "KATSUForceHangingTag"
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
	case KATSULineDirectionTag:
		return "KATSULineDirectionTag"
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
	case KATSULineWidthTagValue:
		return "KATSULineWidthTagValue"
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
		return fmt.Sprintf("KATSULineWidthTag(%d)", e)
	}
}

type KATSUNo uint32

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

type KATSUStrongly uint32

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

type KATSUStyleSingleLineCount uint32

const (
	KATSUStyleDoubleLineCount      KATSUStyleSingleLineCount = 2
	KATSUStyleSingleLineCountValue KATSUStyleSingleLineCount = 1
)

func (e KATSUStyleSingleLineCount) String() string {
	switch e {
	case KATSUStyleDoubleLineCount:
		return "KATSUStyleDoubleLineCount"
	case KATSUStyleSingleLineCountValue:
		return "KATSUStyleSingleLineCountValue"
	default:
		return fmt.Sprintf("KATSUStyleSingleLineCount(%d)", e)
	}
}

type KATSUStyleUnequal uint32

const (
	KATSUStyleContainedBy  KATSUStyleUnequal = 3
	KATSUStyleContains     KATSUStyleUnequal = 1
	KATSUStyleEquals       KATSUStyleUnequal = 2
	KATSUStyleUnequalValue KATSUStyleUnequal = 0
)

func (e KATSUStyleUnequal) String() string {
	switch e {
	case KATSUStyleContainedBy:
		return "KATSUStyleContainedBy"
	case KATSUStyleContains:
		return "KATSUStyleContains"
	case KATSUStyleEquals:
		return "KATSUStyleEquals"
	case KATSUStyleUnequalValue:
		return "KATSUStyleUnequalValue"
	default:
		return fmt.Sprintf("KATSUStyleUnequal(%d)", e)
	}
}

type KATSUTruncateNone uint32

const (
	KATSUTruncFeatNoSquishing      KATSUTruncateNone = 0x8
	KATSUTruncateEnd               KATSUTruncateNone = 2
	KATSUTruncateMiddle            KATSUTruncateNone = 3
	KATSUTruncateNoneValue         KATSUTruncateNone = 0
	KATSUTruncateSpecificationMask KATSUTruncateNone = 0x7
	KATSUTruncateStart             KATSUTruncateNone = 1
)

func (e KATSUTruncateNone) String() string {
	switch e {
	case KATSUTruncFeatNoSquishing:
		return "KATSUTruncFeatNoSquishing"
	case KATSUTruncateEnd:
		return "KATSUTruncateEnd"
	case KATSUTruncateMiddle:
		return "KATSUTruncateMiddle"
	case KATSUTruncateNoneValue:
		return "KATSUTruncateNoneValue"
	case KATSUTruncateSpecificationMask:
		return "KATSUTruncateSpecificationMask"
	case KATSUTruncateStart:
		return "KATSUTruncateStart"
	default:
		return fmt.Sprintf("KATSUTruncateNone(%d)", e)
	}
}

type KATSUUnFlattenOptionNoOptions uint32

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

type KATSUUseGrafPortPenLoc uint32

const (
	KATSUClearAll               KATSUUseGrafPortPenLoc = 0xffffffff
	KATSUUseGrafPortPenLocValue KATSUUseGrafPortPenLoc = 0xffffffff
)

func (e KATSUUseGrafPortPenLoc) String() string {
	switch e {
	case KATSUClearAll:
		return "KATSUClearAll"
	default:
		return fmt.Sprintf("KATSUUseGrafPortPenLoc(%d)", e)
	}
}

type KATSUUseLineControl uint32

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

type KATSUse uint32

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

type KATSUseGlyphAdvance int32

const (
	KATSNoTracking           KATSUseGlyphAdvance = -2147483648
	KATSUseGlyphAdvanceValue KATSUseGlyphAdvance = 0x7fffffff
	KATSUseLineHeight        KATSUseGlyphAdvance = 0x7fffffff
)

func (e KATSUseGlyphAdvance) String() string {
	switch e {
	case KATSNoTracking:
		return "KATSNoTracking"
	case KATSUseGlyphAdvanceValue:
		return "KATSUseGlyphAdvanceValue"
	default:
		return fmt.Sprintf("KATSUseGlyphAdvance(%d)", e)
	}
}

type KAlign uint32

const (
	KAlignAbsoluteCenter   KAlign = 5
	KAlignBottom           KAlign = 0x3
	KAlignBottomLeft       KAlign = 11
	KAlignBottomRight      KAlign = 15
	KAlignCenterBottom     KAlign = 7
	KAlignCenterLeft       KAlign = 9
	KAlignCenterRight      KAlign = 13
	KAlignCenterTop        KAlign = 6
	KAlignHorizontalCenter KAlign = 0x4
	KAlignLeft             KAlign = 0x8
	KAlignNone             KAlign = 0
	KAlignRight            KAlign = 0xc
	KAlignTop              KAlign = 0x2
	KAlignTopLeft          KAlign = 10
	KAlignTopRight         KAlign = 14
	KAlignVerticalCenter   KAlign = 0x1
)

func (e KAlign) String() string {
	switch e {
	case KAlignAbsoluteCenter:
		return "KAlignAbsoluteCenter"
	case KAlignBottom:
		return "KAlignBottom"
	case KAlignBottomLeft:
		return "KAlignBottomLeft"
	case KAlignBottomRight:
		return "KAlignBottomRight"
	case KAlignCenterBottom:
		return "KAlignCenterBottom"
	case KAlignCenterLeft:
		return "KAlignCenterLeft"
	case KAlignCenterRight:
		return "KAlignCenterRight"
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
	case KAlignTop:
		return "KAlignTop"
	case KAlignTopLeft:
		return "KAlignTopLeft"
	case KAlignTopRight:
		return "KAlignTopRight"
	case KAlignVerticalCenter:
		return "KAlignVerticalCenter"
	default:
		return fmt.Sprintf("KAlign(%d)", e)
	}
}

type KAllPPDDomains uint32

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

type KAudioUnit uint32

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

type KDefaultCMM uint32

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

type KFMCurrentFilter uint32

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

type KFMFontTechnologyFilterSelector uint32

const (
	KFMFontCallbackFilterSelector        KFMFontTechnologyFilterSelector = 5
	KFMFontContainerFilterSelector       KFMFontTechnologyFilterSelector = 2
	KFMFontDirectoryFilterSelector       KFMFontTechnologyFilterSelector = 6
	KFMFontFamilyCallbackFilterSelector  KFMFontTechnologyFilterSelector = 4
	KFMFontFileRefFilterSelector         KFMFontTechnologyFilterSelector = 10
	KFMFontTechnologyFilterSelectorValue KFMFontTechnologyFilterSelector = 1
	KFMGenerationFilterSelector          KFMFontTechnologyFilterSelector = 3
)

func (e KFMFontTechnologyFilterSelector) String() string {
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
	case KFMFontTechnologyFilterSelectorValue:
		return "KFMFontTechnologyFilterSelectorValue"
	case KFMGenerationFilterSelector:
		return "KFMGenerationFilterSelector"
	default:
		return fmt.Sprintf("KFMFontTechnologyFilterSelector(%d)", e)
	}
}

type KFMTrueTypeFontTechnology uint32

const (
	KFMPostScriptFontTechnology    KFMTrueTypeFontTechnology = 't'<<24 | 'y'<<16 | 'p'<<8 | '1' // 'typ1'
	KFMTrueTypeFontTechnologyValue KFMTrueTypeFontTechnology = 't'<<24 | 'r'<<16 | 'u'<<8 | 'e' // 'true'
)

func (e KFMTrueTypeFontTechnology) String() string {
	switch e {
	case KFMPostScriptFontTechnology:
		return "KFMPostScriptFontTechnology"
	case KFMTrueTypeFontTechnologyValue:
		return "KFMTrueTypeFontTechnologyValue"
	default:
		return fmt.Sprintf("KFMTrueTypeFontTechnology(%d)", e)
	}
}

type KGlyphCollection uint32

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

type KHIShapeEnumerate uint32

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

type KHIShapeParseFrom uint32

const (
	KHIShapeParseFromBottom      KHIShapeParseFrom = 1
	KHIShapeParseFromBottomRight KHIShapeParseFrom = 3
	KHIShapeParseFromLeft        KHIShapeParseFrom = 0
	KHIShapeParseFromRight       KHIShapeParseFrom = 2
	KHIShapeParseFromTop         KHIShapeParseFrom = 0
	KHIShapeParseFromTopLeft     KHIShapeParseFrom = 0
)

func (e KHIShapeParseFrom) String() string {
	switch e {
	case KHIShapeParseFromBottom:
		return "KHIShapeParseFromBottom"
	case KHIShapeParseFromBottomRight:
		return "KHIShapeParseFromBottomRight"
	case KHIShapeParseFromLeft:
		return "KHIShapeParseFromLeft"
	case KHIShapeParseFromRight:
		return "KHIShapeParseFromRight"
	default:
		return fmt.Sprintf("KHIShapeParseFrom(%d)", e)
	}
}

type KICAttrLockedBit uint32

const (
	KICAttrLockedBitValue KICAttrLockedBit = 0
	KICAttrVolatileBit    KICAttrLockedBit = 1
)

func (e KICAttrLockedBit) String() string {
	switch e {
	case KICAttrLockedBitValue:
		return "KICAttrLockedBitValue"
	case KICAttrVolatileBit:
		return "KICAttrVolatileBit"
	default:
		return fmt.Sprintf("KICAttrLockedBit(%d)", e)
	}
}

type KICAttrNoChange uint32

const (
	KICAttrLockedMask    KICAttrNoChange = 0x1
	KICAttrNoChangeValue KICAttrNoChange = 0xffffffff
	KICAttrVolatileMask  KICAttrNoChange = 0x2
)

func (e KICAttrNoChange) String() string {
	switch e {
	case KICAttrLockedMask:
		return "KICAttrLockedMask"
	case KICAttrNoChangeValue:
		return "KICAttrNoChangeValue"
	case KICAttrVolatileMask:
		return "KICAttrVolatileMask"
	default:
		return fmt.Sprintf("KICAttrNoChange(%d)", e)
	}
}

type KICComponentInterface uint32

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

type KICComponentVersion uint32

const (
	KICComponentVersionValue KICComponentVersion = 0
	KICNumVersion            KICComponentVersion = 1
)

func (e KICComponentVersion) String() string {
	switch e {
	case KICComponentVersionValue:
		return "KICComponentVersionValue"
	case KICNumVersion:
		return "KICNumVersion"
	default:
		return fmt.Sprintf("KICComponentVersion(%d)", e)
	}
}

type KICEditPreferenceEventClass uint32

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

type KICFileSpecHeader uint32

const (
	KICFileSpecHeaderSize KICFileSpecHeader = 106
)

func (e KICFileSpecHeader) String() string {
	switch e {
	case KICFileSpecHeaderSize:
		return "KICFileSpecHeaderSize"
	default:
		return fmt.Sprintf("KICFileSpecHeader(%d)", e)
	}
}

type KICFileType uint32

const (
	KICCreator       KICFileType = 'I'<<24 | 'C'<<16 | 'A'<<8 | 'p' // 'ICAp'
	KICFileTypeValue KICFileType = 'I'<<24 | 'C'<<16 | 'A'<<8 | 'p' // 'ICAp'
)

func (e KICFileType) String() string {
	switch e {
	case KICCreator:
		return "KICCreator"
	default:
		return fmt.Sprintf("KICFileType(%d)", e)
	}
}

type KICMapBinaryBit uint32

const (
	KICMapBinaryBitValue  KICMapBinaryBit = 0
	KICMapDataForkBit     KICMapBinaryBit = 2
	KICMapNotIncomingBit  KICMapBinaryBit = 4
	KICMapNotOutgoingBit  KICMapBinaryBit = 5
	KICMapPostBit         KICMapBinaryBit = 3
	KICMapResourceForkBit KICMapBinaryBit = 1
)

func (e KICMapBinaryBit) String() string {
	switch e {
	case KICMapBinaryBitValue:
		return "KICMapBinaryBitValue"
	case KICMapDataForkBit:
		return "KICMapDataForkBit"
	case KICMapNotIncomingBit:
		return "KICMapNotIncomingBit"
	case KICMapNotOutgoingBit:
		return "KICMapNotOutgoingBit"
	case KICMapPostBit:
		return "KICMapPostBit"
	case KICMapResourceForkBit:
		return "KICMapResourceForkBit"
	default:
		return fmt.Sprintf("KICMapBinaryBit(%d)", e)
	}
}

type KICMapBinaryMask uint32

const (
	KICMapBinaryMaskValue  KICMapBinaryMask = 0x1
	KICMapDataForkMask     KICMapBinaryMask = 0x4
	KICMapNotIncomingMask  KICMapBinaryMask = 0x10
	KICMapNotOutgoingMask  KICMapBinaryMask = 0x20
	KICMapPostMask         KICMapBinaryMask = 0x8
	KICMapResourceForkMask KICMapBinaryMask = 0x2
)

func (e KICMapBinaryMask) String() string {
	switch e {
	case KICMapBinaryMaskValue:
		return "KICMapBinaryMaskValue"
	case KICMapDataForkMask:
		return "KICMapDataForkMask"
	case KICMapNotIncomingMask:
		return "KICMapNotIncomingMask"
	case KICMapNotOutgoingMask:
		return "KICMapNotOutgoingMask"
	case KICMapPostMask:
		return "KICMapPostMask"
	case KICMapResourceForkMask:
		return "KICMapResourceForkMask"
	default:
		return fmt.Sprintf("KICMapBinaryMask(%d)", e)
	}
}

type KICMapFixed uint32

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

type KICNilProfileI uint32

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

const KICNoUserInteractionBit uint32 = 0

const KICNoUserInteractionMask uint32 = 0x1

type KICServicesTCPBit uint32

const (
	KICServicesTCPBitValue KICServicesTCPBit = 0
	KICServicesUDPBit      KICServicesTCPBit = 1
)

func (e KICServicesTCPBit) String() string {
	switch e {
	case KICServicesTCPBitValue:
		return "KICServicesTCPBitValue"
	case KICServicesUDPBit:
		return "KICServicesUDPBit"
	default:
		return fmt.Sprintf("KICServicesTCPBit(%d)", e)
	}
}

type KICServicesTCPMask uint32

const (
	KICServicesTCPMaskValue KICServicesTCPMask = 0x1
	KICServicesUDPMask      KICServicesTCPMask = 0x2
)

func (e KICServicesTCPMask) String() string {
	switch e {
	case KICServicesTCPMaskValue:
		return "KICServicesTCPMaskValue"
	case KICServicesUDPMask:
		return "KICServicesUDPMask"
	default:
		return fmt.Sprintf("KICServicesTCPMask(%d)", e)
	}
}

type KImmediate int32

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

type KInternetEventClass uint32

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

type KInvalid int32

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

type KInvertHighlighting uint32

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

type KNeuter int16

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

type KNoConstraint uint32

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

type KNoEndingProsody int32

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

type KNoProcess uint32

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

type KNoTransform int32

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
	KUseProfileIntent KNoTransform = -1
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

type KPMAllocationFailure int32

const (
	KPMAllocationFailureValue   KPMAllocationFailure = -108
	KPMCVMSymbolNotFound        KPMAllocationFailure = -9662
	KPMCloseFailed              KPMAllocationFailure = -9785
	KPMCreateMessageFailed      KPMAllocationFailure = -9620
	KPMDeleteSubTicketFailed    KPMAllocationFailure = -9585
	KPMDocumentNotFound         KPMAllocationFailure = -9644
	KPMDontSwitchPDEError       KPMAllocationFailure = -9531
	KPMEditRequestFailed        KPMAllocationFailure = -9544
	KPMFeatureNotInstalled      KPMAllocationFailure = -9533
	KPMFileOrDirOperationFailed KPMAllocationFailure = -9634
	KPMFontNameTooLong          KPMAllocationFailure = -9704
	KPMFontNotFound             KPMAllocationFailure = -9703
	KPMGeneralCGError           KPMAllocationFailure = -9705
	KPMIOAttrNotAvailable       KPMAllocationFailure = -9787
	KPMIOMSymbolNotFound        KPMAllocationFailure = -9661
	KPMInternalError            KPMAllocationFailure = -30870
	// KPMInvalidAllocator: The specified memory allocator is invalid.
	KPMInvalidAllocator  KPMAllocationFailure = -30890
	KPMInvalidCVMContext KPMAllocationFailure = -9665
	// KPMInvalidCalibrationTarget: The dictionary specifying a printer calibration target is invalid.
	KPMInvalidCalibrationTarget KPMAllocationFailure = -30898
	// KPMInvalidConnection: The printer connection type is invalid.
	KPMInvalidConnection KPMAllocationFailure = -30887
	// KPMInvalidFileType: The file type is invalid.
	KPMInvalidFileType   KPMAllocationFailure = -30895
	KPMInvalidIOMContext KPMAllocationFailure = -9664
	// KPMInvalidIndex: An array index is invalid.
	KPMInvalidIndex KPMAllocationFailure = -30882
	// KPMInvalidItem: The item being added to a ticket is invalid.
	KPMInvalidItem  KPMAllocationFailure = -30892
	KPMInvalidJobID KPMAllocationFailure = -9666
	// KPMInvalidJobTemplate: An internal error occurred while creating a job template.
	KPMInvalidJobTemplate KPMAllocationFailure = -30885
	// KPMInvalidKey: The key in a ticket, job template, or dictionary is invalid.
	KPMInvalidKey        KPMAllocationFailure = -30888
	KPMInvalidLookupSpec KPMAllocationFailure = -9542
	// KPMInvalidObject: The object is invalid.
	KPMInvalidObject     KPMAllocationFailure = -30896
	KPMInvalidPBMRef     KPMAllocationFailure = -9540
	KPMInvalidPDEContext KPMAllocationFailure = -9530
	KPMInvalidPMContext  KPMAllocationFailure = -9663
	// KPMInvalidPaper: Your application passed an invalid paper object.
	KPMInvalidPaper          KPMAllocationFailure = -30897
	KPMInvalidPrinterAddress KPMAllocationFailure = -9780
	// KPMInvalidPrinterInfo: The printer information is invalid.
	KPMInvalidPrinterInfo KPMAllocationFailure = -30886
	// KPMInvalidReply: A remote server or client sent an invalid reply.
	KPMInvalidReply     KPMAllocationFailure = -30894
	KPMInvalidState     KPMAllocationFailure = -9706
	KPMInvalidSubTicket KPMAllocationFailure = -9584
	// KPMInvalidTicket: The job ticket is invalid.
	KPMInvalidTicket KPMAllocationFailure = -30891
	// KPMInvalidType: The data type in a ticket, job template, or dictionary is not the expected type.
	KPMInvalidType KPMAllocationFailure = -30893
	// KPMInvalidValue: The value in a ticket, job template, or dictionary is missing.
	KPMInvalidValue                                   KPMAllocationFailure = -30889
	KPMItemIsLocked                                   KPMAllocationFailure = -9586
	KPMJobBusy                                        KPMAllocationFailure = -9642
	KPMJobCanceled                                    KPMAllocationFailure = -9643
	KPMJobGetTicketBadFormatError                     KPMAllocationFailure = -9672
	KPMJobGetTicketReadError                          KPMAllocationFailure = -9673
	KPMJobManagerAborted                              KPMAllocationFailure = -9671
	KPMJobNotFound                                    KPMAllocationFailure = -9641
	KPMJobStreamEndError                              KPMAllocationFailure = -9670
	KPMJobStreamOpenFailed                            KPMAllocationFailure = -9668
	KPMJobStreamReadFailed                            KPMAllocationFailure = -9669
	KPMKeyNotFound                                    KPMAllocationFailure = -9589
	KPMKeyNotUnique                                   KPMAllocationFailure = -9590
	KPMKeyOrValueNotFound                             KPMAllocationFailure = -9623
	KPMLastErrorCodeToMakeMaintenanceOfThisListEasier KPMAllocationFailure = -9799
	KPMMessagingError                                 KPMAllocationFailure = -9624
	KPMNoDefaultItem                                  KPMAllocationFailure = -9500
	KPMNoDefaultSettings                              KPMAllocationFailure = -9501
	KPMNoPrinterJobID                                 KPMAllocationFailure = -9667
	KPMNoSelectedPrinters                             KPMAllocationFailure = -9541
	KPMOpenFailed                                     KPMAllocationFailure = -9781
	KPMPMSymbolNotFound                               KPMAllocationFailure = -9660
	KPMPermissionError                                KPMAllocationFailure = -9636
	KPMPluginNotFound                                 KPMAllocationFailure = -9701
	KPMPluginRegisterationFailed                      KPMAllocationFailure = -9702
	KPMPrBrowserNoUI                                  KPMAllocationFailure = -9545
	KPMQueueAlreadyExists                             KPMAllocationFailure = -9639
	KPMQueueJobFailed                                 KPMAllocationFailure = -9640
	KPMQueueNotFound                                  KPMAllocationFailure = -9638
	KPMReadFailed                                     KPMAllocationFailure = -9782
	KPMReadGotZeroData                                KPMAllocationFailure = -9788
	KPMServerAlreadyRunning                           KPMAllocationFailure = -9631
	KPMServerAttributeRestricted                      KPMAllocationFailure = -9633
	KPMServerCommunicationFailed                      KPMAllocationFailure = -9621
	KPMServerNotFound                                 KPMAllocationFailure = -9630
	KPMServerSuspended                                KPMAllocationFailure = -9632
	KPMStatusFailed                                   KPMAllocationFailure = -9784
	// KPMStringConversionFailure: An internal error occurred while converting a string.
	KPMStringConversionFailure KPMAllocationFailure = -30883
	KPMSubTicketNotFound       KPMAllocationFailure = -9583
	KPMSyncRequestFailed       KPMAllocationFailure = -9543
	KPMTemplateIsLocked        KPMAllocationFailure = -9588
	KPMTicketIsLocked          KPMAllocationFailure = -9587
	KPMTicketTypeNotFound      KPMAllocationFailure = -9580
	KPMUnableToFindProcess     KPMAllocationFailure = -9532
	KPMUnexpectedImagingError  KPMAllocationFailure = -9707
	KPMUnknownDataType         KPMAllocationFailure = -9591
	KPMUnknownMessage          KPMAllocationFailure = -9637
	KPMUnsupportedConnection   KPMAllocationFailure = -9786
	KPMUpdateTicketFailed      KPMAllocationFailure = -9581
	KPMUserOrGroupNotFound     KPMAllocationFailure = -9635
	KPMValidateTicketFailed    KPMAllocationFailure = -9582
	KPMWriteFailed             KPMAllocationFailure = -9783
	// KPMXMLParseError: An error occurred while parsing XML data.
	KPMXMLParseError KPMAllocationFailure = -30884
)

func (e KPMAllocationFailure) String() string {
	switch e {
	case KPMAllocationFailureValue:
		return "KPMAllocationFailureValue"
	case KPMCVMSymbolNotFound:
		return "KPMCVMSymbolNotFound"
	case KPMCloseFailed:
		return "KPMCloseFailed"
	case KPMCreateMessageFailed:
		return "KPMCreateMessageFailed"
	case KPMDeleteSubTicketFailed:
		return "KPMDeleteSubTicketFailed"
	case KPMDocumentNotFound:
		return "KPMDocumentNotFound"
	case KPMDontSwitchPDEError:
		return "KPMDontSwitchPDEError"
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
	case KPMIOAttrNotAvailable:
		return "KPMIOAttrNotAvailable"
	case KPMIOMSymbolNotFound:
		return "KPMIOMSymbolNotFound"
	case KPMInternalError:
		return "KPMInternalError"
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
	case KPMInvalidPaper:
		return "KPMInvalidPaper"
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
	case KPMNoDefaultSettings:
		return "KPMNoDefaultSettings"
	case KPMNoPrinterJobID:
		return "KPMNoPrinterJobID"
	case KPMNoSelectedPrinters:
		return "KPMNoSelectedPrinters"
	case KPMOpenFailed:
		return "KPMOpenFailed"
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
	case KPMWriteFailed:
		return "KPMWriteFailed"
	case KPMXMLParseError:
		return "KPMXMLParseError"
	default:
		return fmt.Sprintf("KPMAllocationFailure(%d)", e)
	}
}

type KPMBorder uint32

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

const KPMCancel uint32 = 0x80

type KPMCoverPage uint32

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

type KPMDestination uint32

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

type KPMDuplexNone uint32

const (
	// KPMDuplexNoTumble: # Discussion
	KPMDuplexNoTumble KPMDuplexNone = 0x2
	// KPMDuplexNoneValue: # Discussion
	KPMDuplexNoneValue KPMDuplexNone = 0x1
	// KPMDuplexTumble: # Discussion
	KPMDuplexTumble KPMDuplexNone = 0x3
	// KPMSimplexTumble: # Discussion
	KPMSimplexTumble KPMDuplexNone = 0x4
)

func (e KPMDuplexNone) String() string {
	switch e {
	case KPMDuplexNoTumble:
		return "KPMDuplexNoTumble"
	case KPMDuplexNoneValue:
		return "KPMDuplexNoneValue"
	case KPMDuplexTumble:
		return "KPMDuplexTumble"
	case KPMSimplexTumble:
		return "KPMSimplexTumble"
	default:
		return fmt.Sprintf("KPMDuplexNone(%d)", e)
	}
}

type KPMHideInlineItems uint32

const (
	KPMHideInlineItemsValue             KPMHideInlineItems = 0
	KPMShowDefaultInlineItems           KPMHideInlineItems = 32768
	KPMShowInlineCopies                 KPMHideInlineItems = 1
	KPMShowInlineOrientation            KPMHideInlineItems = 8
	KPMShowInlinePageRange              KPMHideInlineItems = 2
	KPMShowInlinePageRangeWithSelection KPMHideInlineItems = 64
	KPMShowInlinePaperSize              KPMHideInlineItems = 4
	KPMShowInlineScale                  KPMHideInlineItems = 128
	KPMShowPageAttributesPDE            KPMHideInlineItems = 256
)

func (e KPMHideInlineItems) String() string {
	switch e {
	case KPMHideInlineItemsValue:
		return "KPMHideInlineItemsValue"
	case KPMShowDefaultInlineItems:
		return "KPMShowDefaultInlineItems"
	case KPMShowInlineCopies:
		return "KPMShowInlineCopies"
	case KPMShowInlineOrientation:
		return "KPMShowInlineOrientation"
	case KPMShowInlinePageRange:
		return "KPMShowInlinePageRange"
	case KPMShowInlinePageRangeWithSelection:
		return "KPMShowInlinePageRangeWithSelection"
	case KPMShowInlinePaperSize:
		return "KPMShowInlinePaperSize"
	case KPMShowInlineScale:
		return "KPMShowInlineScale"
	case KPMShowPageAttributesPDE:
		return "KPMShowPageAttributesPDE"
	default:
		return fmt.Sprintf("KPMHideInlineItems(%d)", e)
	}
}

type KPMInvalidPrintSession int32

const (
	// KPMInvalidPreset: Your application passed an invalid preset object.
	KPMInvalidPreset KPMInvalidPrintSession = -30899
	// KPMInvalidPrintSessionValue: Your application passed an invalid printing session object.
	KPMInvalidPrintSessionValue KPMInvalidPrintSession = -30879
	// KPMInvalidPrinter: Your application passed an invalid printer object.
	KPMInvalidPrinter KPMInvalidPrintSession = -30880
	// KPMObjectInUse: The specified object is in use.
	KPMObjectInUse KPMInvalidPrintSession = -30881
)

func (e KPMInvalidPrintSession) String() string {
	switch e {
	case KPMInvalidPreset:
		return "KPMInvalidPreset"
	case KPMInvalidPrintSessionValue:
		return "KPMInvalidPrintSessionValue"
	case KPMInvalidPrinter:
		return "KPMInvalidPrinter"
	case KPMObjectInUse:
		return "KPMObjectInUse"
	default:
		return fmt.Sprintf("KPMInvalidPrintSession(%d)", e)
	}
}

type KPMLayout uint32

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

type KPMNoError int32

const (
	// KPMGeneralError: An unspecified error occurred.
	KPMGeneralError KPMNoError = -30870
	// KPMInvalidPageFormat: Your application passed an invalid page format object.
	KPMInvalidPageFormat KPMNoError = -30876
	KPMInvalidParameter  KPMNoError = -50
	// KPMInvalidPrintSettings: Your application passed an invalid print settings object.
	KPMInvalidPrintSettings KPMNoError = -30875
	// KPMNoDefaultPrinter: The user has not specified a default printer.
	KPMNoDefaultPrinter KPMNoError = -30872
	KPMNoErrorValue     KPMNoError = 0
	// KPMNoSuchEntry: There is no entry to match your application’s request.
	KPMNoSuchEntry KPMNoError = -30874
	// KPMNotImplemented: The function is not implemented.
	KPMNotImplemented KPMNoError = -30873
	// KPMOutOfScope: Your application called this function out of sequence with other printing functions.
	KPMOutOfScope KPMNoError = -30871
	// KPMValueOutOfRange: Your application passed an out-of-range value.
	KPMValueOutOfRange KPMNoError = -30877
)

func (e KPMNoError) String() string {
	switch e {
	case KPMGeneralError:
		return "KPMGeneralError"
	case KPMInvalidPageFormat:
		return "KPMInvalidPageFormat"
	case KPMInvalidParameter:
		return "KPMInvalidParameter"
	case KPMInvalidPrintSettings:
		return "KPMInvalidPrintSettings"
	case KPMNoDefaultPrinter:
		return "KPMNoDefaultPrinter"
	case KPMNoErrorValue:
		return "KPMNoErrorValue"
	case KPMNoSuchEntry:
		return "KPMNoSuchEntry"
	case KPMNotImplemented:
		return "KPMNotImplemented"
	case KPMOutOfScope:
		return "KPMOutOfScope"
	case KPMValueOutOfRange:
		return "KPMValueOutOfRange"
	default:
		return fmt.Sprintf("KPMNoError(%d)", e)
	}
}

type KPMPaperType uint32

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

type KPMPortrait uint32

const (
	// KPMLandscape: # Discussion
	KPMLandscape KPMPortrait = 2
	// KPMPortraitValue: Specifies portrait (vertical) page orientation.
	KPMPortraitValue KPMPortrait = 1
	// KPMReverseLandscape: # Discussion
	KPMReverseLandscape KPMPortrait = 4
	// KPMReversePortrait: # Discussion
	KPMReversePortrait KPMPortrait = 3
)

func (e KPMPortrait) String() string {
	switch e {
	case KPMLandscape:
		return "KPMLandscape"
	case KPMPortraitValue:
		return "KPMPortraitValue"
	case KPMReverseLandscape:
		return "KPMReverseLandscape"
	case KPMReversePortrait:
		return "KPMReversePortrait"
	default:
		return fmt.Sprintf("KPMPortrait(%d)", e)
	}
}

type KPMPrintAll int32

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

type KPMPrinter uint32

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

type KPMQuality uint32

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

type KPMScaling uint32

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

type KPMUnknownColorSpaceModel uint32

const (
	KPMCMYKColorSpaceModel         KPMUnknownColorSpaceModel = 3
	KPMDevNColorSpaceModel         KPMUnknownColorSpaceModel = 4
	KPMGrayColorSpaceModel         KPMUnknownColorSpaceModel = 1
	KPMRGBColorSpaceModel          KPMUnknownColorSpaceModel = 2
	KPMUnknownColorSpaceModelValue KPMUnknownColorSpaceModel = 0
)

func (e KPMUnknownColorSpaceModel) String() string {
	switch e {
	case KPMCMYKColorSpaceModel:
		return "KPMCMYKColorSpaceModel"
	case KPMDevNColorSpaceModel:
		return "KPMDevNColorSpaceModel"
	case KPMGrayColorSpaceModel:
		return "KPMGrayColorSpaceModel"
	case KPMRGBColorSpaceModel:
		return "KPMRGBColorSpaceModel"
	case KPMUnknownColorSpaceModelValue:
		return "KPMUnknownColorSpaceModelValue"
	default:
		return fmt.Sprintf("KPMUnknownColorSpaceModel(%d)", e)
	}
}

const KPMUnlocked uint32 = 0

type KPlotIconRef uint32

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

type KProcessDictionaryIncludeAllInformation int32

const (
	KProcessDictionaryIncludeAllInformationMask KProcessDictionaryIncludeAllInformation = -1
)

func (e KProcessDictionaryIncludeAllInformation) String() string {
	switch e {
	case KProcessDictionaryIncludeAllInformationMask:
		return "KProcessDictionaryIncludeAllInformationMask"
	default:
		return fmt.Sprintf("KProcessDictionaryIncludeAllInformation(%d)", e)
	}
}

type KProcessTransformTo uint32

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

type KQuit uint32

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

type KSelector uint32

const (
	KSelectorAll1BitData      KSelector = 16843009
	KSelectorAll32BitData     KSelector = 134219784
	KSelectorAll4BitData      KSelector = 33686018
	KSelectorAll8BitData      KSelector = 67372036
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
	case KSelectorLarge1Bit:
		return "KSelectorLarge1Bit"
	case KSelectorLarge32Bit:
		return "KSelectorLarge32Bit"
	case KSelectorLarge4Bit:
		return "KSelectorLarge4Bit"
	case KSelectorLarge8Bit:
		return "KSelectorLarge8Bit"
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

type KSetFrontProcess uint32

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

type KSpeech int32

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

type KTextToSpeech uint32

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

type KTransform uint32

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
	KTransformSelectedDisabled KTransform = 16385
	KTransformSelectedOffline  KTransform = 16386
	KTransformSelectedOpen     KTransform = 16387
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
	case KTransformSelectedDisabled:
		return "KTransformSelectedDisabled"
	case KTransformSelectedOffline:
		return "KTransformSelectedOffline"
	case KTransformSelectedOpen:
		return "KTransformSelectedOpen"
	default:
		return fmt.Sprintf("KTransform(%d)", e)
	}
}

type KTranslation uint32

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

type KUAZoomFocusType uint32

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

type Launch uint32

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

type ModeReserved uint32

const (
	Mode32BitCompatible        ModeReserved = 0x80
	ModeCanBackground          ModeReserved = 0x1000
	ModeControlPanel           ModeReserved = 0x80000
	ModeDeskAccessory          ModeReserved = 0x20000
	ModeDisplayManagerAware    ModeReserved = 0x4
	ModeDoesActivateOnFGSwitch ModeReserved = 0x800
	ModeGetAppDiedMsg          ModeReserved = 0x100
	ModeGetFrontClicks         ModeReserved = 0x200
	ModeHighLevelEventAware    ModeReserved = 0x40
	ModeLaunchDontSwitch       ModeReserved = 0x40000
	ModeLocalAndRemoteHLEvents ModeReserved = 0x20
	ModeMultiLaunch            ModeReserved = 0x10000
	ModeNeedSuspendResume      ModeReserved = 0x4000
	ModeOnlyBackground         ModeReserved = 0x400
	ModeReservedValue          ModeReserved = 0x1000000
	ModeStationeryAware        ModeReserved = 0x10
	ModeUseTextEditServices    ModeReserved = 0x8
)

func (e ModeReserved) String() string {
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
	case ModeLocalAndRemoteHLEvents:
		return "ModeLocalAndRemoteHLEvents"
	case ModeMultiLaunch:
		return "ModeMultiLaunch"
	case ModeNeedSuspendResume:
		return "ModeNeedSuspendResume"
	case ModeOnlyBackground:
		return "ModeOnlyBackground"
	case ModeReservedValue:
		return "ModeReservedValue"
	case ModeStationeryAware:
		return "ModeStationeryAware"
	case ModeUseTextEditServices:
		return "ModeUseTextEditServices"
	default:
		return fmt.Sprintf("ModeReserved(%d)", e)
	}
}

type ModeText uint32

const (
	// ModeLiteral: When the speech channel is in text-processing mode, indicates that characters and digits are spoken literally (for example, “cat” is spoken as “C-A-T” and “12” is spoken as "one, two").
	ModeLiteral ModeText = 'L'<<24 | 'T'<<16 | 'R'<<8 | 'L' // 'LTRL'
	// ModeNormal: When the speech channel is in text-processing mode, indicates that the synthesizer should process characters as expected and assemble digits into numbers.
	ModeNormal ModeText = 'N'<<24 | 'O'<<16 | 'R'<<8 | 'M' // 'NORM'
	// ModePhonemes: Used with soInputMode to indicate that the speech channel is in phoneme-processing mode.
	ModePhonemes ModeText = 'P'<<24 | 'H'<<16 | 'O'<<8 | 'N' // 'PHON'
	// ModeTextValue: Used with soInputMode to indicate that the speech channel is in text-processing mode.
	ModeTextValue ModeText = 'T'<<24 | 'E'<<16 | 'X'<<8 | 'T' // 'TEXT'
	ModeTune      ModeText = 'T'<<24 | 'U'<<16 | 'N'<<8 | 'E' // 'TUNE'
)

func (e ModeText) String() string {
	switch e {
	case ModeLiteral:
		return "ModeLiteral"
	case ModeNormal:
		return "ModeNormal"
	case ModePhonemes:
		return "ModePhonemes"
	case ModeTextValue:
		return "ModeTextValue"
	case ModeTune:
		return "ModeTune"
	default:
		return fmt.Sprintf("ModeText(%d)", e)
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
	NotSaved         PasteboardFlavorFlags = 4
	Promised         PasteboardFlavorFlags = 512
	RequestOnly      PasteboardFlavorFlags = 8
	SenderOnly       PasteboardFlavorFlags = 1
	SenderTranslated PasteboardFlavorFlags = 2
	SystemTranslated PasteboardFlavorFlags = 256
)

func (e PasteboardFlavorFlags) String() string {
	switch e {
	case NotSaved:
		return "NotSaved"
	case Promised:
		return "Promised"
	case RequestOnly:
		return "RequestOnly"
	case SenderOnly:
		return "SenderOnly"
	case SenderTranslated:
		return "SenderTranslated"
	case SystemTranslated:
		return "SystemTranslated"
	default:
		return fmt.Sprintf("PasteboardFlavorFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/applicationservices/pasteboardstandardlocation
type PasteboardStandardLocation uint32

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
	ClientIsOwner PasteboardSyncFlags = 2
	Modified      PasteboardSyncFlags = 1
)

func (e PasteboardSyncFlags) String() string {
	switch e {
	case ClientIsOwner:
		return "ClientIsOwner"
	case Modified:
		return "Modified"
	default:
		return fmt.Sprintf("PasteboardSyncFlags(%d)", e)
	}
}

type So uint32

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

type SoVoice uint32

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

type Sv uint32

const (
	SvAll1BitData      Sv = 16843009
	SvAll4BitData      Sv = 33686018
	SvAll8BitData      Sv = 67372036
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
	case SvLarge1Bit:
		return "SvLarge1Bit"
	case SvLarge4Bit:
		return "SvLarge4Bit"
	case SvLarge8Bit:
		return "SvLarge8Bit"
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

type Tt uint32

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
	TtSelectedDisabled Tt = 16385
	TtSelectedOffline  Tt = 16386
	TtSelectedOpen     Tt = 16387
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
	case TtSelectedDisabled:
		return "TtSelectedDisabled"
	case TtSelectedOffline:
		return "TtSelectedOffline"
	case TtSelectedOpen:
		return "TtSelectedOpen"
	default:
		return fmt.Sprintf("Tt(%d)", e)
	}
}

// IconRef is an alias for referenced enum type KPlotIconRef.
type IconRef = KPlotIconRef
