// Code generated from Apple documentation. DO NOT EDIT.

package uniformtypeidentifiers

import (
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var UTTagClassFilenameExtension string

var UTTagClassMIMEType string

var UTType3DContent UTType

var uTTypeAHAP UTType

var uTTypeAIFF UTType

var uTTypeARReferenceObject UTType

var uTTypeAVI UTType

var uTTypeAliasFile UTType

var uTTypeAppleArchive UTType

var uTTypeAppleProtectedMPEG4Audio UTType

var uTTypeAppleProtectedMPEG4Video UTType

var uTTypeAppleScript UTType

var uTTypeApplication UTType

var uTTypeApplicationBundle UTType

var uTTypeApplicationExtension UTType

var uTTypeArchive UTType

var uTTypeAssemblyLanguageSource UTType

var uTTypeAudio UTType

var uTTypeAudiovisualContent UTType

var uTTypeBMP UTType

var uTTypeBZ2 UTType

var uTTypeBinaryPropertyList UTType

var uTTypeBookmark UTType

var uTTypeBundle UTType

var uTTypeCHeader UTType

var uTTypeCPlusPlusHeader UTType

var uTTypeCPlusPlusSource UTType

var uTTypeCSS UTType

var uTTypeCSource UTType

var uTTypeCalendarEvent UTType

var uTTypeCommaSeparatedText UTType

var uTTypeCompositeContent UTType

var uTTypeContact UTType

var uTTypeContent UTType

var uTTypeDNG UTType

var uTTypeData UTType

var uTTypeDatabase UTType

var uTTypeDelimitedText UTType

var uTTypeDirectory UTType

var uTTypeDiskImage UTType

var uTTypeEPUB UTType

var uTTypeEXE UTType

var uTTypeEXR UTType

var uTTypeEmailMessage UTType

var uTTypeExecutable UTType

var uTTypeFileURL UTType

var uTTypeFlatRTFD UTType

var uTTypeFolder UTType

var uTTypeFont UTType

var uTTypeFramework UTType

var uTTypeGIF UTType

var uTTypeGZIP UTType

var uTTypeGeoJSON UTType

var uTTypeHEIC UTType

var uTTypeHEICS UTType

var uTTypeHEIF UTType

var uTTypeHTML UTType

var uTTypeICNS UTType

var uTTypeICO UTType

var uTTypeImage UTType

var uTTypeInternetLocation UTType

var uTTypeInternetShortcut UTType

var uTTypeItem UTType

var uTTypeJPEG UTType

var uTTypeJPEGXL UTType

var uTTypeJSON UTType

var uTTypeJavaScript UTType

var uTTypeLinkPresentationMetadata UTType

var uTTypeLivePhoto UTType

var uTTypeLog UTType

var uTTypeM3UPlaylist UTType

var uTTypeMIDI UTType

var uTTypeMP3 UTType

var uTTypeMPEG UTType

var uTTypeMPEG2TransportStream UTType

var uTTypeMPEG2Video UTType

var uTTypeMPEG4Audio UTType

var uTTypeMPEG4Movie UTType

var uTTypeMakefile UTType

var uTTypeMessage UTType

var uTTypeMountPoint UTType

var uTTypeMovie UTType

var uTTypeOSAScript UTType

var uTTypeOSAScriptBundle UTType

var uTTypeObjectiveCPlusPlusSource UTType

var uTTypeObjectiveCSource UTType

var uTTypePDF UTType

var uTTypePHPScript UTType

var uTTypePKCS12 UTType

var uTTypePNG UTType

var uTTypePackage UTType

var uTTypePerlScript UTType

var uTTypePlainText UTType

var uTTypePlaylist UTType

var uTTypePluginBundle UTType

var uTTypePresentation UTType

var uTTypePropertyList UTType

var uTTypePythonScript UTType

var uTTypeQuickLookGenerator UTType

var uTTypeQuickTimeMovie UTType

var uTTypeRAWImage UTType

var uTTypeRTF UTType

var uTTypeRTFD UTType

var uTTypeRealityFile UTType

var uTTypeResolvable UTType

var uTTypeRubyScript UTType

var uTTypeSVG UTType

var uTTypeSceneKitScene UTType

var uTTypeScript UTType

var uTTypeShellScript UTType

var uTTypeSourceCode UTType

var uTTypeSpotlightImporter UTType

var uTTypeSpreadsheet UTType

var uTTypeSwiftSource UTType

var uTTypeSymbolicLink UTType

var uTTypeSystemPreferencesPane UTType

var uTTypeTIFF UTType

var uTTypeTabSeparatedText UTType

var uTTypeTarArchive UTType

var uTTypeText UTType

var uTTypeToDoItem UTType

var uTTypeURL UTType

var uTTypeURLBookmarkData UTType

var uTTypeUSD UTType

var uTTypeUSDZ UTType

var uTTypeUTF16ExternalPlainText UTType

var uTTypeUTF16PlainText UTType

var uTTypeUTF8PlainText UTType

var uTTypeUTF8TabSeparatedText UTType

var uTTypeUnixExecutable UTType

var uTTypeVCard UTType

var uTTypeVideo UTType

var uTTypeVolume UTType

var uTTypeWAV UTType

var uTTypeWebArchive UTType

var uTTypeWebP UTType

var uTTypeX509Certificate UTType

var uTTypeXML UTType

var uTTypeXMLPropertyList UTType

var uTTypeXPCService UTType

var uTTypeYAML UTType

var uTTypeZIP UTType

func init() {
	if frameworkHandle == 0 {
		return
	}


	if ptr, err := purego.Dlsym(frameworkHandle, "UTTagClassFilenameExtension"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UTTagClassFilenameExtension = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTagClassMIMEType"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UTTagClassMIMEType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTType3DContent"); err == nil && ptr != 0 {
		UTType3DContent = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAHAP"); err == nil && ptr != 0 {
		uTTypeAHAP = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAIFF"); err == nil && ptr != 0 {
		uTTypeAIFF = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeARReferenceObject"); err == nil && ptr != 0 {
		uTTypeARReferenceObject = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAVI"); err == nil && ptr != 0 {
		uTTypeAVI = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAliasFile"); err == nil && ptr != 0 {
		uTTypeAliasFile = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAppleArchive"); err == nil && ptr != 0 {
		uTTypeAppleArchive = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAppleProtectedMPEG4Audio"); err == nil && ptr != 0 {
		uTTypeAppleProtectedMPEG4Audio = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAppleProtectedMPEG4Video"); err == nil && ptr != 0 {
		uTTypeAppleProtectedMPEG4Video = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAppleScript"); err == nil && ptr != 0 {
		uTTypeAppleScript = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeApplication"); err == nil && ptr != 0 {
		uTTypeApplication = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeApplicationBundle"); err == nil && ptr != 0 {
		uTTypeApplicationBundle = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeApplicationExtension"); err == nil && ptr != 0 {
		uTTypeApplicationExtension = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeArchive"); err == nil && ptr != 0 {
		uTTypeArchive = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAssemblyLanguageSource"); err == nil && ptr != 0 {
		uTTypeAssemblyLanguageSource = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAudio"); err == nil && ptr != 0 {
		uTTypeAudio = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAudiovisualContent"); err == nil && ptr != 0 {
		uTTypeAudiovisualContent = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeBMP"); err == nil && ptr != 0 {
		uTTypeBMP = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeBZ2"); err == nil && ptr != 0 {
		uTTypeBZ2 = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeBinaryPropertyList"); err == nil && ptr != 0 {
		uTTypeBinaryPropertyList = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeBookmark"); err == nil && ptr != 0 {
		uTTypeBookmark = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeBundle"); err == nil && ptr != 0 {
		uTTypeBundle = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeCHeader"); err == nil && ptr != 0 {
		uTTypeCHeader = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeCPlusPlusHeader"); err == nil && ptr != 0 {
		uTTypeCPlusPlusHeader = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeCPlusPlusSource"); err == nil && ptr != 0 {
		uTTypeCPlusPlusSource = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeCSS"); err == nil && ptr != 0 {
		uTTypeCSS = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeCSource"); err == nil && ptr != 0 {
		uTTypeCSource = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeCalendarEvent"); err == nil && ptr != 0 {
		uTTypeCalendarEvent = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeCommaSeparatedText"); err == nil && ptr != 0 {
		uTTypeCommaSeparatedText = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeCompositeContent"); err == nil && ptr != 0 {
		uTTypeCompositeContent = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeContact"); err == nil && ptr != 0 {
		uTTypeContact = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeContent"); err == nil && ptr != 0 {
		uTTypeContent = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeDNG"); err == nil && ptr != 0 {
		uTTypeDNG = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeData"); err == nil && ptr != 0 {
		uTTypeData = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeDatabase"); err == nil && ptr != 0 {
		uTTypeDatabase = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeDelimitedText"); err == nil && ptr != 0 {
		uTTypeDelimitedText = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeDirectory"); err == nil && ptr != 0 {
		uTTypeDirectory = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeDiskImage"); err == nil && ptr != 0 {
		uTTypeDiskImage = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeEPUB"); err == nil && ptr != 0 {
		uTTypeEPUB = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeEXE"); err == nil && ptr != 0 {
		uTTypeEXE = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeEXR"); err == nil && ptr != 0 {
		uTTypeEXR = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeEmailMessage"); err == nil && ptr != 0 {
		uTTypeEmailMessage = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeExecutable"); err == nil && ptr != 0 {
		uTTypeExecutable = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeFileURL"); err == nil && ptr != 0 {
		uTTypeFileURL = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeFlatRTFD"); err == nil && ptr != 0 {
		uTTypeFlatRTFD = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeFolder"); err == nil && ptr != 0 {
		uTTypeFolder = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeFont"); err == nil && ptr != 0 {
		uTTypeFont = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeFramework"); err == nil && ptr != 0 {
		uTTypeFramework = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeGIF"); err == nil && ptr != 0 {
		uTTypeGIF = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeGZIP"); err == nil && ptr != 0 {
		uTTypeGZIP = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeGeoJSON"); err == nil && ptr != 0 {
		uTTypeGeoJSON = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeHEIC"); err == nil && ptr != 0 {
		uTTypeHEIC = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeHEICS"); err == nil && ptr != 0 {
		uTTypeHEICS = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeHEIF"); err == nil && ptr != 0 {
		uTTypeHEIF = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeHTML"); err == nil && ptr != 0 {
		uTTypeHTML = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeICNS"); err == nil && ptr != 0 {
		uTTypeICNS = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeICO"); err == nil && ptr != 0 {
		uTTypeICO = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeImage"); err == nil && ptr != 0 {
		uTTypeImage = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeInternetLocation"); err == nil && ptr != 0 {
		uTTypeInternetLocation = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeInternetShortcut"); err == nil && ptr != 0 {
		uTTypeInternetShortcut = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeItem"); err == nil && ptr != 0 {
		uTTypeItem = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeJPEG"); err == nil && ptr != 0 {
		uTTypeJPEG = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeJPEGXL"); err == nil && ptr != 0 {
		uTTypeJPEGXL = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeJSON"); err == nil && ptr != 0 {
		uTTypeJSON = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeJavaScript"); err == nil && ptr != 0 {
		uTTypeJavaScript = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeLinkPresentationMetadata"); err == nil && ptr != 0 {
		uTTypeLinkPresentationMetadata = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeLivePhoto"); err == nil && ptr != 0 {
		uTTypeLivePhoto = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeLog"); err == nil && ptr != 0 {
		uTTypeLog = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeM3UPlaylist"); err == nil && ptr != 0 {
		uTTypeM3UPlaylist = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMIDI"); err == nil && ptr != 0 {
		uTTypeMIDI = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMP3"); err == nil && ptr != 0 {
		uTTypeMP3 = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMPEG"); err == nil && ptr != 0 {
		uTTypeMPEG = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMPEG2TransportStream"); err == nil && ptr != 0 {
		uTTypeMPEG2TransportStream = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMPEG2Video"); err == nil && ptr != 0 {
		uTTypeMPEG2Video = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMPEG4Audio"); err == nil && ptr != 0 {
		uTTypeMPEG4Audio = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMPEG4Movie"); err == nil && ptr != 0 {
		uTTypeMPEG4Movie = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMakefile"); err == nil && ptr != 0 {
		uTTypeMakefile = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMessage"); err == nil && ptr != 0 {
		uTTypeMessage = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMountPoint"); err == nil && ptr != 0 {
		uTTypeMountPoint = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMovie"); err == nil && ptr != 0 {
		uTTypeMovie = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeOSAScript"); err == nil && ptr != 0 {
		uTTypeOSAScript = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeOSAScriptBundle"); err == nil && ptr != 0 {
		uTTypeOSAScriptBundle = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeObjectiveCPlusPlusSource"); err == nil && ptr != 0 {
		uTTypeObjectiveCPlusPlusSource = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeObjectiveCSource"); err == nil && ptr != 0 {
		uTTypeObjectiveCSource = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePDF"); err == nil && ptr != 0 {
		uTTypePDF = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePHPScript"); err == nil && ptr != 0 {
		uTTypePHPScript = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePKCS12"); err == nil && ptr != 0 {
		uTTypePKCS12 = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePNG"); err == nil && ptr != 0 {
		uTTypePNG = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePackage"); err == nil && ptr != 0 {
		uTTypePackage = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePerlScript"); err == nil && ptr != 0 {
		uTTypePerlScript = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePlainText"); err == nil && ptr != 0 {
		uTTypePlainText = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePlaylist"); err == nil && ptr != 0 {
		uTTypePlaylist = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePluginBundle"); err == nil && ptr != 0 {
		uTTypePluginBundle = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePresentation"); err == nil && ptr != 0 {
		uTTypePresentation = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePropertyList"); err == nil && ptr != 0 {
		uTTypePropertyList = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePythonScript"); err == nil && ptr != 0 {
		uTTypePythonScript = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeQuickLookGenerator"); err == nil && ptr != 0 {
		uTTypeQuickLookGenerator = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeQuickTimeMovie"); err == nil && ptr != 0 {
		uTTypeQuickTimeMovie = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeRAWImage"); err == nil && ptr != 0 {
		uTTypeRAWImage = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeRTF"); err == nil && ptr != 0 {
		uTTypeRTF = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeRTFD"); err == nil && ptr != 0 {
		uTTypeRTFD = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeRealityFile"); err == nil && ptr != 0 {
		uTTypeRealityFile = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeResolvable"); err == nil && ptr != 0 {
		uTTypeResolvable = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeRubyScript"); err == nil && ptr != 0 {
		uTTypeRubyScript = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeSVG"); err == nil && ptr != 0 {
		uTTypeSVG = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeSceneKitScene"); err == nil && ptr != 0 {
		uTTypeSceneKitScene = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeScript"); err == nil && ptr != 0 {
		uTTypeScript = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeShellScript"); err == nil && ptr != 0 {
		uTTypeShellScript = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeSourceCode"); err == nil && ptr != 0 {
		uTTypeSourceCode = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeSpotlightImporter"); err == nil && ptr != 0 {
		uTTypeSpotlightImporter = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeSpreadsheet"); err == nil && ptr != 0 {
		uTTypeSpreadsheet = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeSwiftSource"); err == nil && ptr != 0 {
		uTTypeSwiftSource = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeSymbolicLink"); err == nil && ptr != 0 {
		uTTypeSymbolicLink = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeSystemPreferencesPane"); err == nil && ptr != 0 {
		uTTypeSystemPreferencesPane = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeTIFF"); err == nil && ptr != 0 {
		uTTypeTIFF = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeTabSeparatedText"); err == nil && ptr != 0 {
		uTTypeTabSeparatedText = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeTarArchive"); err == nil && ptr != 0 {
		uTTypeTarArchive = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeText"); err == nil && ptr != 0 {
		uTTypeText = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeToDoItem"); err == nil && ptr != 0 {
		uTTypeToDoItem = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeURL"); err == nil && ptr != 0 {
		uTTypeURL = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeURLBookmarkData"); err == nil && ptr != 0 {
		uTTypeURLBookmarkData = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeUSD"); err == nil && ptr != 0 {
		uTTypeUSD = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeUSDZ"); err == nil && ptr != 0 {
		uTTypeUSDZ = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeUTF16ExternalPlainText"); err == nil && ptr != 0 {
		uTTypeUTF16ExternalPlainText = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeUTF16PlainText"); err == nil && ptr != 0 {
		uTTypeUTF16PlainText = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeUTF8PlainText"); err == nil && ptr != 0 {
		uTTypeUTF8PlainText = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeUTF8TabSeparatedText"); err == nil && ptr != 0 {
		uTTypeUTF8TabSeparatedText = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeUnixExecutable"); err == nil && ptr != 0 {
		uTTypeUnixExecutable = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeVCard"); err == nil && ptr != 0 {
		uTTypeVCard = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeVideo"); err == nil && ptr != 0 {
		uTTypeVideo = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeVolume"); err == nil && ptr != 0 {
		uTTypeVolume = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeWAV"); err == nil && ptr != 0 {
		uTTypeWAV = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeWebArchive"); err == nil && ptr != 0 {
		uTTypeWebArchive = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeWebP"); err == nil && ptr != 0 {
		uTTypeWebP = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeX509Certificate"); err == nil && ptr != 0 {
		uTTypeX509Certificate = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeXML"); err == nil && ptr != 0 {
		uTTypeXML = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeXMLPropertyList"); err == nil && ptr != 0 {
		uTTypeXMLPropertyList = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeXPCService"); err == nil && ptr != 0 {
		uTTypeXPCService = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeYAML"); err == nil && ptr != 0 {
		uTTypeYAML = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeZIP"); err == nil && ptr != 0 {
		uTTypeZIP = *(*UTType)(unsafe.Pointer(ptr))
	}

}

type UTTypeValues struct{}

// UTTypes provides typed accessors for [UTType] constants.
var UTTypes UTTypeValues

func (UTTypeValues) AHAP() UTType { return uTTypeAHAP }

// AIFF returns A type that represents data in AIFF audio format.
func (UTTypeValues) AIFF() UTType { return uTTypeAIFF }

// ARReferenceObject returns A type that represents an augmented reality reference object.
func (UTTypeValues) ARReferenceObject() UTType { return uTTypeARReferenceObject }

// AVI returns A type that represents data in AVI movie format.
func (UTTypeValues) AVI() UTType { return uTTypeAVI }

// AliasFile returns A type that represents an alias file.
func (UTTypeValues) AliasFile() UTType { return uTTypeAliasFile }

// AppleArchive returns A type that represents an Apple archive of files and directories.
func (UTTypeValues) AppleArchive() UTType { return uTTypeAppleArchive }

// AppleProtectedMPEG4Audio returns A type that represents data in Apple-protected MPEG-4 format.
func (UTTypeValues) AppleProtectedMPEG4Audio() UTType { return uTTypeAppleProtectedMPEG4Audio }

// AppleProtectedMPEG4Video returns A type that represents data in Apple-protected MPEG-4 format.
func (UTTypeValues) AppleProtectedMPEG4Video() UTType { return uTTypeAppleProtectedMPEG4Video }

// AppleScript returns A type that represents an AppleScript text-based script.
func (UTTypeValues) AppleScript() UTType { return uTTypeAppleScript }

// Application returns A base type that represents a macOS, iOS, iPadOS, watchOS, and tvOS app.
func (UTTypeValues) Application() UTType { return uTTypeApplication }

// ApplicationBundle returns A type that represents a bundled app.
func (UTTypeValues) ApplicationBundle() UTType { return uTTypeApplicationBundle }

// ApplicationExtension returns A type that represents an app extension.
func (UTTypeValues) ApplicationExtension() UTType { return uTTypeApplicationExtension }

// Archive returns A base type that represents an archive of files and directories.
func (UTTypeValues) Archive() UTType { return uTTypeArchive }

// AssemblyLanguageSource returns A type that represents assembly language source code.
func (UTTypeValues) AssemblyLanguageSource() UTType { return uTTypeAssemblyLanguageSource }

// Audio returns A type that represents audio that doesn’t contain video.
func (UTTypeValues) Audio() UTType { return uTTypeAudio }

// AudiovisualContent returns A base type that represents data that contains video content that may or may not also include audio.
func (UTTypeValues) AudiovisualContent() UTType { return uTTypeAudiovisualContent }

// BMP returns A type that represents a Windows bitmap image.
func (UTTypeValues) BMP() UTType { return uTTypeBMP }

// BZ2 returns A type that represents a bzip2 archive.
func (UTTypeValues) BZ2() UTType { return uTTypeBZ2 }

// BinaryPropertyList returns A type that represents a binary property list.
func (UTTypeValues) BinaryPropertyList() UTType { return uTTypeBinaryPropertyList }

// Bookmark returns A base type that represents bookmark data.
func (UTTypeValues) Bookmark() UTType { return uTTypeBookmark }

// Bundle returns A base type that represents a directory that conforms to one of the bundle layouts.
func (UTTypeValues) Bundle() UTType { return uTTypeBundle }

// CHeader returns A type that represents a C header file.
func (UTTypeValues) CHeader() UTType { return uTTypeCHeader }

// CPlusPlusHeader returns A type that represents a C++ header file.
func (UTTypeValues) CPlusPlusHeader() UTType { return uTTypeCPlusPlusHeader }

// CPlusPlusSource returns A type that represents a C++ source code file.
func (UTTypeValues) CPlusPlusSource() UTType { return uTTypeCPlusPlusSource }

func (UTTypeValues) CSS() UTType { return uTTypeCSS }

// CSource returns A type that represents a C source code file.
func (UTTypeValues) CSource() UTType { return uTTypeCSource }

// CalendarEvent returns A base type that represents a calendar event.
func (UTTypeValues) CalendarEvent() UTType { return uTTypeCalendarEvent }

// CommaSeparatedText returns A type that represents text containing comma-separated values.
func (UTTypeValues) CommaSeparatedText() UTType { return uTTypeCommaSeparatedText }

// CompositeContent returns A base type that represents a content format supporting mixed embedded content.
func (UTTypeValues) CompositeContent() UTType { return uTTypeCompositeContent }

// Contact returns A base type that represents contact information.
func (UTTypeValues) Contact() UTType { return uTTypeContact }

// Content returns A base type that represents anything containing user-viewable content.
func (UTTypeValues) Content() UTType { return uTTypeContent }

func (UTTypeValues) DNG() UTType { return uTTypeDNG }

// Data returns A base type that represents any sort of byte stream, including files and in-memory data.
func (UTTypeValues) Data() UTType { return uTTypeData }

// Database returns A base type that represents a database store.
func (UTTypeValues) Database() UTType { return uTTypeDatabase }

// DelimitedText returns A base type that represents text containing delimited values.
func (UTTypeValues) DelimitedText() UTType { return uTTypeDelimitedText }

// Directory returns A type that represents a file system directory, including packages and folders.
func (UTTypeValues) Directory() UTType { return uTTypeDirectory }

// DiskImage returns A type that represents a data item that’s mountable as a volume.
func (UTTypeValues) DiskImage() UTType { return uTTypeDiskImage }

// EPUB returns A type that represents data in the electronic publication (EPUB) format.
func (UTTypeValues) EPUB() UTType { return uTTypeEPUB }

// EXE returns A type that represents a Windows executable.
func (UTTypeValues) EXE() UTType { return uTTypeEXE }

func (UTTypeValues) EXR() UTType { return uTTypeEXR }

// EmailMessage returns A type that represents an email message.
func (UTTypeValues) EmailMessage() UTType { return uTTypeEmailMessage }

// Executable returns A type that represents an executable.
func (UTTypeValues) Executable() UTType { return uTTypeExecutable }

// FileURL returns A type that represents a URL to a file in the file system.
func (UTTypeValues) FileURL() UTType { return uTTypeFileURL }

// FlatRTFD returns A type that represents flattened Rich Text Format Directory documents.
func (UTTypeValues) FlatRTFD() UTType { return uTTypeFlatRTFD }

// Folder returns A type that represents a user-browsable directory.
func (UTTypeValues) Folder() UTType { return uTTypeFolder }

// Font returns A base type that represents a font.
func (UTTypeValues) Font() UTType { return uTTypeFont }

// Framework returns A type that represents an Apple framework bundle.
func (UTTypeValues) Framework() UTType { return uTTypeFramework }

// GIF returns A type that represents a GIF image.
func (UTTypeValues) GIF() UTType { return uTTypeGIF }

// GZIP returns A type that represents a GNU zip archive.
func (UTTypeValues) GZIP() UTType { return uTTypeGZIP }

func (UTTypeValues) GeoJSON() UTType { return uTTypeGeoJSON }

// HEIC returns A type that represents High Efficiency Image Coding images.
func (UTTypeValues) HEIC() UTType { return uTTypeHEIC }

func (UTTypeValues) HEICS() UTType { return uTTypeHEICS }

// HEIF returns A type that represents High Efficiency Image File Format images.
func (UTTypeValues) HEIF() UTType { return uTTypeHEIF }

// HTML returns A type that represents any version of HTML.
func (UTTypeValues) HTML() UTType { return uTTypeHTML }

// ICNS returns A type that represents Apple icon data.
func (UTTypeValues) ICNS() UTType { return uTTypeICNS }

// ICO returns A type that represents Windows icon data.
func (UTTypeValues) ICO() UTType { return uTTypeICO }

// Image returns A base type that represents image data.
func (UTTypeValues) Image() UTType { return uTTypeImage }

// InternetLocation returns A base type that represents an Apple internet location file.
func (UTTypeValues) InternetLocation() UTType { return uTTypeInternetLocation }

// InternetShortcut returns A type that represents a Microsoft internet shortcut file.
func (UTTypeValues) InternetShortcut() UTType { return uTTypeInternetShortcut }

// Item returns A generic base type for most objects, such as files or directories.
func (UTTypeValues) Item() UTType { return uTTypeItem }

// JPEG returns A type that represents a JPEG image.
func (UTTypeValues) JPEG() UTType { return uTTypeJPEG }

func (UTTypeValues) JPEGXL() UTType { return uTTypeJPEGXL }

// JSON returns A type that represents JavaScript Object Notation (JSON) data.
func (UTTypeValues) JSON() UTType { return uTTypeJSON }

// JavaScript returns A type that represents JavaScript source code.
func (UTTypeValues) JavaScript() UTType { return uTTypeJavaScript }

func (UTTypeValues) LinkPresentationMetadata() UTType { return uTTypeLinkPresentationMetadata }

// LivePhoto returns A type that represents Live Photos.
func (UTTypeValues) LivePhoto() UTType { return uTTypeLivePhoto }

// Log returns A base type that represents console log data.
func (UTTypeValues) Log() UTType { return uTTypeLog }

// M3UPlaylist returns A type that represents an M3U or M3U8 playlist.
func (UTTypeValues) M3UPlaylist() UTType { return uTTypeM3UPlaylist }

// MIDI returns A type that represents data in MIDI audio format.
func (UTTypeValues) MIDI() UTType { return uTTypeMIDI }

// MP3 returns A type that represents MP3 audio.
func (UTTypeValues) MP3() UTType { return uTTypeMP3 }

// MPEG returns A type that represents an MPEG-1 or MPEG-2 movie.
func (UTTypeValues) MPEG() UTType { return uTTypeMPEG }

// MPEG2TransportStream returns A type that represents data in MPEG-2 transport stream movie format.
func (UTTypeValues) MPEG2TransportStream() UTType { return uTTypeMPEG2TransportStream }

// MPEG2Video returns A type that represents an MPEG-2 video.
func (UTTypeValues) MPEG2Video() UTType { return uTTypeMPEG2Video }

// MPEG4Audio returns A type that represents an MPEG-4 audio layer file.
func (UTTypeValues) MPEG4Audio() UTType { return uTTypeMPEG4Audio }

// MPEG4Movie returns A type that represents an MPEG-4 movie.
func (UTTypeValues) MPEG4Movie() UTType { return uTTypeMPEG4Movie }

// Makefile returns A type that represents a Makefile.
func (UTTypeValues) Makefile() UTType { return uTTypeMakefile }

// Message returns A base type that represents a message.
func (UTTypeValues) Message() UTType { return uTTypeMessage }

// MountPoint returns A type that represents a volume mount point.
func (UTTypeValues) MountPoint() UTType { return uTTypeMountPoint }

// Movie returns A base type representing media formats that may contain both video and audio.
func (UTTypeValues) Movie() UTType { return uTTypeMovie }

// OSAScript returns A type that represents an Open Scripting Architecture binary script.
func (UTTypeValues) OSAScript() UTType { return uTTypeOSAScript }

// OSAScriptBundle returns A type that represents an Open Scripting Architecture script bundle.
func (UTTypeValues) OSAScriptBundle() UTType { return uTTypeOSAScriptBundle }

// ObjectiveCPlusPlusSource returns A type that represents an Objective-C++ source code file.
func (UTTypeValues) ObjectiveCPlusPlusSource() UTType { return uTTypeObjectiveCPlusPlusSource }

// ObjectiveCSource returns A type that represents an Objective-C source code file.
func (UTTypeValues) ObjectiveCSource() UTType { return uTTypeObjectiveCSource }

// PDF returns A type that represents Adobe Portable Document Format (PDF) documents.
func (UTTypeValues) PDF() UTType { return uTTypePDF }

// PHPScript returns A type that represents a PHP script.
func (UTTypeValues) PHPScript() UTType { return uTTypePHPScript }

// PKCS12 returns A type that represents Public Key Cryptography Standard (PKCS) 12 data.
func (UTTypeValues) PKCS12() UTType { return uTTypePKCS12 }

// PNG returns A type that represents a PNG image.
func (UTTypeValues) PNG() UTType { return uTTypePNG }

// Package returns A base type that represents a packaged directory.
func (UTTypeValues) Package() UTType { return uTTypePackage }

// PerlScript returns A type that represents a Perl script.
func (UTTypeValues) PerlScript() UTType { return uTTypePerlScript }

// PlainText returns A type that represents text with no markup and an unspecified encoding.
func (UTTypeValues) PlainText() UTType { return uTTypePlainText }

// Playlist returns A base type that represents a playlist.
func (UTTypeValues) Playlist() UTType { return uTTypePlaylist }

// PluginBundle returns A base type that represents a bundle-based plug-in.
func (UTTypeValues) PluginBundle() UTType { return uTTypePluginBundle }

// Presentation returns A base type that represents a presentation document.
func (UTTypeValues) Presentation() UTType { return uTTypePresentation }

// PropertyList returns A base type that represents a property list.
func (UTTypeValues) PropertyList() UTType { return uTTypePropertyList }

// PythonScript returns A type that represents a Python script.
func (UTTypeValues) PythonScript() UTType { return uTTypePythonScript }

// QuickLookGenerator returns A type that represents a QuickLook preview generator bundle.
func (UTTypeValues) QuickLookGenerator() UTType { return uTTypeQuickLookGenerator }

// QuickTimeMovie returns A type that represents a QuickTime movie.
func (UTTypeValues) QuickTimeMovie() UTType { return uTTypeQuickTimeMovie }

// RAWImage returns A base type that represents a raw image format that you use in digital photography.
func (UTTypeValues) RAWImage() UTType { return uTTypeRAWImage }

// RTF returns A type that represents Rich Text Format data.
func (UTTypeValues) RTF() UTType { return uTTypeRTF }

// RTFD returns A type that represents Rich Text Format Directory documents.
func (UTTypeValues) RTFD() UTType { return uTTypeRTFD }

// RealityFile returns A type that represents a Reality Composer file.
func (UTTypeValues) RealityFile() UTType { return uTTypeRealityFile }

// Resolvable returns A base type that represents a resolvable reference, including symbolic links and aliases.
func (UTTypeValues) Resolvable() UTType { return uTTypeResolvable }

// RubyScript returns A type that represents a Ruby script.
func (UTTypeValues) RubyScript() UTType { return uTTypeRubyScript }

// SVG returns A type that represents a scalable vector graphics (SVG) image.
func (UTTypeValues) SVG() UTType { return uTTypeSVG }

// SceneKitScene returns A type that represents a SceneKit serialized scene.
func (UTTypeValues) SceneKitScene() UTType { return uTTypeSceneKitScene }

// Script returns A base type that represents any scripting language source.
func (UTTypeValues) Script() UTType { return uTTypeScript }

// ShellScript returns A base type that represents a shell script.
func (UTTypeValues) ShellScript() UTType { return uTTypeShellScript }

// SourceCode returns A base type that represents source code of any programming language.
func (UTTypeValues) SourceCode() UTType { return uTTypeSourceCode }

// SpotlightImporter returns A type that represents a Spotlight metadata importer bundle.
func (UTTypeValues) SpotlightImporter() UTType { return uTTypeSpotlightImporter }

// Spreadsheet returns A base type that represents a spreadsheet document.
func (UTTypeValues) Spreadsheet() UTType { return uTTypeSpreadsheet }

// SwiftSource returns A type that represents a Swift source code file.
func (UTTypeValues) SwiftSource() UTType { return uTTypeSwiftSource }

// SymbolicLink returns A type that represents a symbolic link.
func (UTTypeValues) SymbolicLink() UTType { return uTTypeSymbolicLink }

// SystemPreferencesPane returns A type that represents a System Preferences pane.
func (UTTypeValues) SystemPreferencesPane() UTType { return uTTypeSystemPreferencesPane }

// TIFF returns A type that represents a TIFF image.
func (UTTypeValues) TIFF() UTType { return uTTypeTIFF }

// TabSeparatedText returns A type that represents text containing tab-separated values.
func (UTTypeValues) TabSeparatedText() UTType { return uTTypeTabSeparatedText }

func (UTTypeValues) TarArchive() UTType { return uTTypeTarArchive }

// Text returns A base type that represents all text-encoded data, including text with markup.
func (UTTypeValues) Text() UTType { return uTTypeText }

// ToDoItem returns A type that represents a to-do item.
func (UTTypeValues) ToDoItem() UTType { return uTTypeToDoItem }

// URL returns A type that represents a URL.
func (UTTypeValues) URL() UTType { return uTTypeURL }

// URLBookmarkData returns A type that represents a URL bookmark.
func (UTTypeValues) URLBookmarkData() UTType { return uTTypeURLBookmarkData }

// USD returns A type that represents Universal Scene Description content.
func (UTTypeValues) USD() UTType { return uTTypeUSD }

// USDZ returns A type that represents Universal Scene Description Package content.
func (UTTypeValues) USDZ() UTType { return uTTypeUSDZ }

// UTF16ExternalPlainText returns A type that represents plain text encoded as UTF-16 with an optional BOM.
func (UTTypeValues) UTF16ExternalPlainText() UTType { return uTTypeUTF16ExternalPlainText }

// UTF16PlainText returns A type that represents plain text encoded as UTF-16 in native byte order with an optional bill of materials.
func (UTTypeValues) UTF16PlainText() UTType { return uTTypeUTF16PlainText }

// UTF8PlainText returns A type that represents plain text encoded as UTF-8.
func (UTTypeValues) UTF8PlainText() UTType { return uTTypeUTF8PlainText }

// UTF8TabSeparatedText returns A type that represents UTF-8–encoded text containing tab-separated values.
func (UTTypeValues) UTF8TabSeparatedText() UTType { return uTTypeUTF8TabSeparatedText }

// UnixExecutable returns A type that represents a UNIX executable.
func (UTTypeValues) UnixExecutable() UTType { return uTTypeUnixExecutable }

// VCard returns A type that represents a vCard file.
func (UTTypeValues) VCard() UTType { return uTTypeVCard }

// Video returns A type that represents video that doesn’t contain audio.
func (UTTypeValues) Video() UTType { return uTTypeVideo }

// Volume returns A type that represents the root folder of a volume or mount point.
func (UTTypeValues) Volume() UTType { return uTTypeVolume }

// WAV returns A type that represents data in Microsoft Waveform Audio File Format.
func (UTTypeValues) WAV() UTType { return uTTypeWAV }

// WebArchive returns A type that represents WebKit web archive data.
func (UTTypeValues) WebArchive() UTType { return uTTypeWebArchive }

// WebP returns A type that represents a WebP image.
func (UTTypeValues) WebP() UTType { return uTTypeWebP }

// X509Certificate returns A type that represents an X.509 certificate.
func (UTTypeValues) X509Certificate() UTType { return uTTypeX509Certificate }

// XML returns A type that represents generic XML data.
func (UTTypeValues) XML() UTType { return uTTypeXML }

// XMLPropertyList returns A type that represents an XML property list.
func (UTTypeValues) XMLPropertyList() UTType { return uTTypeXMLPropertyList }

// XPCService returns A type that represents an XPC service bundle.
func (UTTypeValues) XPCService() UTType { return uTTypeXPCService }

// YAML returns A type that represents Yet Another Markup Language data.
func (UTTypeValues) YAML() UTType { return uTTypeYAML }

// ZIP returns A type that represents a zip archive.
func (UTTypeValues) ZIP() UTType { return uTTypeZIP }


