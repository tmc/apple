// Code generated from Apple documentation. DO NOT EDIT.

package colorsync

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// See: https://developer.apple.com/documentation/ColorSync/kCMMApplyTransformProcName
	KCMMApplyTransformProcName string
	// See: https://developer.apple.com/documentation/ColorSync/kCMMCreateTransformPropertyProcName
	KCMMCreateTransformPropertyProcName string
	// See: https://developer.apple.com/documentation/ColorSync/kCMMInitializeLinkProfileProcName
	KCMMInitializeLinkProfileProcName string
	// See: https://developer.apple.com/documentation/ColorSync/kCMMInitializeTransformProcName
	KCMMInitializeTransformProcName string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncACESCGLinearProfile
	KColorSyncACESCGLinearProfile string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncAdobeRGB1998Profile
	KColorSyncAdobeRGB1998Profile string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncBestQuality
	KColorSyncBestQuality string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncBlackPointCompensation
	KColorSyncBlackPointCompensation string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncCameraDeviceClass
	KColorSyncCameraDeviceClass string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncConversion1DLut
	KColorSyncConversion1DLut string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncConversion3DLut
	KColorSyncConversion3DLut string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncConversionBPC
	KColorSyncConversionBPC string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncConversionChannelID
	KColorSyncConversionChannelID string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncConversionGridPoints
	KColorSyncConversionGridPoints string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncConversionInpChan
	KColorSyncConversionInpChan string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncConversionMatrix
	KColorSyncConversionMatrix string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncConversionNDLut
	KColorSyncConversionNDLut string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncConversionOutChan
	KColorSyncConversionOutChan string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncConversionParamCurve0
	KColorSyncConversionParamCurve0 string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncConversionParamCurve1
	KColorSyncConversionParamCurve1 string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncConversionParamCurve2
	KColorSyncConversionParamCurve2 string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncConversionParamCurve3
	KColorSyncConversionParamCurve3 string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncConversionParamCurve4
	KColorSyncConversionParamCurve4 string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncConvertQuality
	KColorSyncConvertQuality string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncConvertUseExtendedRange
	KColorSyncConvertUseExtendedRange string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncCustomProfiles
	KColorSyncCustomProfiles string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDCIP3Profile
	KColorSyncDCIP3Profile string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDeviceClass
	KColorSyncDeviceClass string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDeviceDefaultProfileID
	KColorSyncDeviceDefaultProfileID string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDeviceDescription
	KColorSyncDeviceDescription string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDeviceDescriptions
	KColorSyncDeviceDescriptions string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDeviceHostScope
	KColorSyncDeviceHostScope string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDeviceID
	KColorSyncDeviceID string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDeviceModeDescription
	KColorSyncDeviceModeDescription string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDeviceModeDescriptions
	KColorSyncDeviceModeDescriptions string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDeviceProfileID
	KColorSyncDeviceProfileID string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDeviceProfileIsCurrent
	KColorSyncDeviceProfileIsCurrent string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDeviceProfileIsDefault
	KColorSyncDeviceProfileIsDefault string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDeviceProfileIsFactory
	KColorSyncDeviceProfileIsFactory string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDeviceProfileURL
	KColorSyncDeviceProfileURL string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDeviceProfilesNotification
	KColorSyncDeviceProfilesNotification string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDeviceRegisteredNotification
	KColorSyncDeviceRegisteredNotification string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDeviceUnregisteredNotification
	KColorSyncDeviceUnregisteredNotification string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDeviceUserScope
	KColorSyncDeviceUserScope string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDisplayDeviceClass
	KColorSyncDisplayDeviceClass string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDisplayDeviceProfilesNotification
	KColorSyncDisplayDeviceProfilesNotification string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDisplayP3Profile
	KColorSyncDisplayP3Profile string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDoNotSubstituteProfiles
	KColorSyncDoNotSubstituteProfiles string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncDraftQuality
	KColorSyncDraftQuality string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncExtendedRange
	KColorSyncExtendedRange string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncFactoryProfiles
	KColorSyncFactoryProfiles string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncFixedPointRange
	KColorSyncFixedPointRange string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncGenericCMYKProfile
	KColorSyncGenericCMYKProfile string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncGenericGrayGamma22Profile
	KColorSyncGenericGrayGamma22Profile string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncGenericGrayProfile
	KColorSyncGenericGrayProfile string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncGenericLabProfile
	KColorSyncGenericLabProfile string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncGenericRGBProfile
	KColorSyncGenericRGBProfile string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncGenericXYZProfile
	KColorSyncGenericXYZProfile string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncHDRDerivative
	KColorSyncHDRDerivative string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncHLGDerivative
	KColorSyncHLGDerivative string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncITUR2020Profile
	KColorSyncITUR2020Profile string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncITUR709Profile
	KColorSyncITUR709Profile string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncNormalQuality
	KColorSyncNormalQuality string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncPQDerivative
	KColorSyncPQDerivative string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncPreferredCMM
	KColorSyncPreferredCMM string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncPrinterDeviceClass
	KColorSyncPrinterDeviceClass string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncProfile
	KColorSyncProfile string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncProfileCacheSeed
	KColorSyncProfileCacheSeed string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncProfileClass
	KColorSyncProfileClass string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncProfileColorSpace
	KColorSyncProfileColorSpace string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncProfileComputerDomain
	KColorSyncProfileComputerDomain string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncProfileDescription
	KColorSyncProfileDescription string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncProfileHeader
	KColorSyncProfileHeader string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncProfileHostScope
	KColorSyncProfileHostScope string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncProfileIsValid
	KColorSyncProfileIsValid string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncProfileMD5Digest
	KColorSyncProfileMD5Digest string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncProfilePCS
	KColorSyncProfilePCS string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncProfileRepositoryChangeNotification
	KColorSyncProfileRepositoryChangeNotification string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncProfileURL
	KColorSyncProfileURL string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncProfileUserDomain
	KColorSyncProfileUserDomain string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncProfileUserScope
	KColorSyncProfileUserScope string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncROMMRGBProfile
	KColorSyncROMMRGBProfile string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncRegistrationUpdateWindowServer
	KColorSyncRegistrationUpdateWindowServer string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncRenderingIntent
	KColorSyncRenderingIntent string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncRenderingIntentAbsolute
	KColorSyncRenderingIntentAbsolute string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncRenderingIntentPerceptual
	KColorSyncRenderingIntentPerceptual string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncRenderingIntentRelative
	KColorSyncRenderingIntentRelative string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncRenderingIntentSaturation
	KColorSyncRenderingIntentSaturation string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncRenderingIntentUseProfileHeader
	KColorSyncRenderingIntentUseProfileHeader string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSRGBProfile
	KColorSyncSRGBProfile string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncScannerDeviceClass
	KColorSyncScannerDeviceClass string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigAToB0Tag
	KColorSyncSigAToB0Tag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigAToB1Tag
	KColorSyncSigAToB1Tag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigAToB2Tag
	KColorSyncSigAToB2Tag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigAbstractClass
	KColorSyncSigAbstractClass string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigBToA0Tag
	KColorSyncSigBToA0Tag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigBToA1Tag
	KColorSyncSigBToA1Tag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigBToA2Tag
	KColorSyncSigBToA2Tag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigBlueColorantTag
	KColorSyncSigBlueColorantTag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigBlueTRCTag
	KColorSyncSigBlueTRCTag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigCmykData
	KColorSyncSigCmykData string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigColorSpaceClass
	KColorSyncSigColorSpaceClass string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigCopyrightTag
	KColorSyncSigCopyrightTag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigDeviceMfgDescTag
	KColorSyncSigDeviceMfgDescTag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigDeviceModelDescTag
	KColorSyncSigDeviceModelDescTag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigDisplayClass
	KColorSyncSigDisplayClass string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigGamutTag
	KColorSyncSigGamutTag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigGrayData
	KColorSyncSigGrayData string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigGrayTRCTag
	KColorSyncSigGrayTRCTag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigGreenColorantTag
	KColorSyncSigGreenColorantTag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigGreenTRCTag
	KColorSyncSigGreenTRCTag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigInputClass
	KColorSyncSigInputClass string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigLabData
	KColorSyncSigLabData string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigLinkClass
	KColorSyncSigLinkClass string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigMediaBlackPointTag
	KColorSyncSigMediaBlackPointTag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigMediaWhitePointTag
	KColorSyncSigMediaWhitePointTag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigNamedColor2Tag
	KColorSyncSigNamedColor2Tag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigNamedColorClass
	KColorSyncSigNamedColorClass string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigOutputClass
	KColorSyncSigOutputClass string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigPreview0Tag
	KColorSyncSigPreview0Tag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigPreview1Tag
	KColorSyncSigPreview1Tag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigPreview2Tag
	KColorSyncSigPreview2Tag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigProfileDescriptionTag
	KColorSyncSigProfileDescriptionTag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigProfileSequenceDescTag
	KColorSyncSigProfileSequenceDescTag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigRedColorantTag
	KColorSyncSigRedColorantTag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigRedTRCTag
	KColorSyncSigRedTRCTag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigRgbData
	KColorSyncSigRgbData string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigTechnologyTag
	KColorSyncSigTechnologyTag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigViewingCondDescTag
	KColorSyncSigViewingCondDescTag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigViewingConditionsTag
	KColorSyncSigViewingConditionsTag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncSigXYZData
	KColorSyncSigXYZData string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncTransformCodeFragmentMD5
	KColorSyncTransformCodeFragmentMD5 string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncTransformCodeFragmentType
	KColorSyncTransformCodeFragmentType string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncTransformCreator
	KColorSyncTransformCreator string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncTransformDeviceToDevice
	KColorSyncTransformDeviceToDevice string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncTransformDeviceToPCS
	KColorSyncTransformDeviceToPCS string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncTransformDstSpace
	KColorSyncTransformDstSpace string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncTransformFullConversionData
	KColorSyncTransformFullConversionData string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncTransformGamutCheck
	KColorSyncTransformGamutCheck string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncTransformInfo
	KColorSyncTransformInfo string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncTransformPCSToDevice
	KColorSyncTransformPCSToDevice string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncTransformPCSToPCS
	KColorSyncTransformPCSToPCS string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncTransformParametricConversionData
	KColorSyncTransformParametricConversionData string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncTransformProfileSequnce
	KColorSyncTransformProfileSequnce string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncTransformSimplifiedConversionData
	KColorSyncTransformSimplifiedConversionData string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncTransformSrcSpace
	KColorSyncTransformSrcSpace string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncTransformTag
	KColorSyncTransformTag string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncTransformUseITU709OETF
	KColorSyncTransformUseITU709OETF string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncWaitForCacheReply
	KColorSyncWaitForCacheReply string
	// See: https://developer.apple.com/documentation/ColorSync/kColorSyncWebSafeColorsProfile
	KColorSyncWebSafeColorsProfile string
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMApplyTransformProcName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMApplyTransformProcName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMCreateTransformPropertyProcName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMCreateTransformPropertyProcName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMInitializeLinkProfileProcName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMInitializeLinkProfileProcName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMMInitializeTransformProcName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMMInitializeTransformProcName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncACESCGLinearProfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncACESCGLinearProfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncAdobeRGB1998Profile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncAdobeRGB1998Profile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncBestQuality"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncBestQuality = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncBlackPointCompensation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncBlackPointCompensation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncCameraDeviceClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncCameraDeviceClass = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncConversion1DLut"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncConversion1DLut = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncConversion3DLut"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncConversion3DLut = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncConversionBPC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncConversionBPC = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncConversionChannelID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncConversionChannelID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncConversionGridPoints"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncConversionGridPoints = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncConversionInpChan"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncConversionInpChan = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncConversionMatrix"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncConversionMatrix = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncConversionNDLut"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncConversionNDLut = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncConversionOutChan"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncConversionOutChan = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncConversionParamCurve0"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncConversionParamCurve0 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncConversionParamCurve1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncConversionParamCurve1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncConversionParamCurve2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncConversionParamCurve2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncConversionParamCurve3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncConversionParamCurve3 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncConversionParamCurve4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncConversionParamCurve4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncConvertQuality"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncConvertQuality = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncConvertUseExtendedRange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncConvertUseExtendedRange = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncCustomProfiles"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncCustomProfiles = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDCIP3Profile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDCIP3Profile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDeviceClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDeviceClass = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDeviceDefaultProfileID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDeviceDefaultProfileID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDeviceDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDeviceDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDeviceDescriptions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDeviceDescriptions = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDeviceHostScope"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDeviceHostScope = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDeviceID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDeviceID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDeviceModeDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDeviceModeDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDeviceModeDescriptions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDeviceModeDescriptions = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDeviceProfileID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDeviceProfileID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDeviceProfileIsCurrent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDeviceProfileIsCurrent = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDeviceProfileIsDefault"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDeviceProfileIsDefault = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDeviceProfileIsFactory"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDeviceProfileIsFactory = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDeviceProfileURL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDeviceProfileURL = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDeviceProfilesNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDeviceProfilesNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDeviceRegisteredNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDeviceRegisteredNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDeviceUnregisteredNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDeviceUnregisteredNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDeviceUserScope"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDeviceUserScope = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDisplayDeviceClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDisplayDeviceClass = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDisplayDeviceProfilesNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDisplayDeviceProfilesNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDisplayP3Profile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDisplayP3Profile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDoNotSubstituteProfiles"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDoNotSubstituteProfiles = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncDraftQuality"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncDraftQuality = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncExtendedRange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncExtendedRange = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncFactoryProfiles"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncFactoryProfiles = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncFixedPointRange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncFixedPointRange = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncGenericCMYKProfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncGenericCMYKProfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncGenericGrayGamma22Profile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncGenericGrayGamma22Profile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncGenericGrayProfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncGenericGrayProfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncGenericLabProfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncGenericLabProfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncGenericRGBProfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncGenericRGBProfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncGenericXYZProfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncGenericXYZProfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncHDRDerivative"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncHDRDerivative = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncHLGDerivative"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncHLGDerivative = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncITUR2020Profile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncITUR2020Profile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncITUR709Profile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncITUR709Profile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncNormalQuality"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncNormalQuality = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncPQDerivative"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncPQDerivative = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncPreferredCMM"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncPreferredCMM = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncPrinterDeviceClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncPrinterDeviceClass = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncProfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncProfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncProfileCacheSeed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncProfileCacheSeed = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncProfileClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncProfileClass = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncProfileColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncProfileColorSpace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncProfileComputerDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncProfileComputerDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncProfileDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncProfileDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncProfileHeader"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncProfileHeader = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncProfileHostScope"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncProfileHostScope = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncProfileIsValid"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncProfileIsValid = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncProfileMD5Digest"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncProfileMD5Digest = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncProfilePCS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncProfilePCS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncProfileRepositoryChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncProfileRepositoryChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncProfileURL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncProfileURL = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncProfileUserDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncProfileUserDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncProfileUserScope"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncProfileUserScope = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncROMMRGBProfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncROMMRGBProfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncRegistrationUpdateWindowServer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncRegistrationUpdateWindowServer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncRenderingIntent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncRenderingIntent = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncRenderingIntentAbsolute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncRenderingIntentAbsolute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncRenderingIntentPerceptual"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncRenderingIntentPerceptual = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncRenderingIntentRelative"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncRenderingIntentRelative = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncRenderingIntentSaturation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncRenderingIntentSaturation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncRenderingIntentUseProfileHeader"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncRenderingIntentUseProfileHeader = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSRGBProfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSRGBProfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncScannerDeviceClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncScannerDeviceClass = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigAToB0Tag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigAToB0Tag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigAToB1Tag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigAToB1Tag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigAToB2Tag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigAToB2Tag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigAbstractClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigAbstractClass = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigBToA0Tag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigBToA0Tag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigBToA1Tag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigBToA1Tag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigBToA2Tag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigBToA2Tag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigBlueColorantTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigBlueColorantTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigBlueTRCTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigBlueTRCTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigCmykData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigCmykData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigColorSpaceClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigColorSpaceClass = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigCopyrightTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigCopyrightTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigDeviceMfgDescTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigDeviceMfgDescTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigDeviceModelDescTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigDeviceModelDescTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigDisplayClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigDisplayClass = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigGamutTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigGamutTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigGrayData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigGrayData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigGrayTRCTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigGrayTRCTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigGreenColorantTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigGreenColorantTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigGreenTRCTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigGreenTRCTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigInputClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigInputClass = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigLabData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigLabData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigLinkClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigLinkClass = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigMediaBlackPointTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigMediaBlackPointTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigMediaWhitePointTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigMediaWhitePointTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigNamedColor2Tag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigNamedColor2Tag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigNamedColorClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigNamedColorClass = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigOutputClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigOutputClass = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigPreview0Tag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigPreview0Tag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigPreview1Tag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigPreview1Tag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigPreview2Tag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigPreview2Tag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigProfileDescriptionTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigProfileDescriptionTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigProfileSequenceDescTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigProfileSequenceDescTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigRedColorantTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigRedColorantTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigRedTRCTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigRedTRCTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigRgbData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigRgbData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigTechnologyTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigTechnologyTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigViewingCondDescTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigViewingCondDescTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigViewingConditionsTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigViewingConditionsTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncSigXYZData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncSigXYZData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncTransformCodeFragmentMD5"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncTransformCodeFragmentMD5 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncTransformCodeFragmentType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncTransformCodeFragmentType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncTransformCreator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncTransformCreator = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncTransformDeviceToDevice"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncTransformDeviceToDevice = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncTransformDeviceToPCS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncTransformDeviceToPCS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncTransformDstSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncTransformDstSpace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncTransformFullConversionData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncTransformFullConversionData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncTransformGamutCheck"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncTransformGamutCheck = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncTransformInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncTransformInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncTransformPCSToDevice"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncTransformPCSToDevice = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncTransformPCSToPCS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncTransformPCSToPCS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncTransformParametricConversionData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncTransformParametricConversionData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncTransformProfileSequnce"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncTransformProfileSequnce = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncTransformSimplifiedConversionData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncTransformSimplifiedConversionData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncTransformSrcSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncTransformSrcSpace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncTransformTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncTransformTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncTransformUseITU709OETF"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncTransformUseITU709OETF = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncWaitForCacheReply"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncWaitForCacheReply = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kColorSyncWebSafeColorsProfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KColorSyncWebSafeColorsProfile = objc.GoString(cstr)
			}
		}
	}

}
