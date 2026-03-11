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
		UTTypes.AHAP = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAIFF"); err == nil && ptr != 0 {
		UTTypes.AIFF = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeARReferenceObject"); err == nil && ptr != 0 {
		UTTypes.ARReferenceObject = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAVI"); err == nil && ptr != 0 {
		UTTypes.AVI = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAliasFile"); err == nil && ptr != 0 {
		UTTypes.AliasFile = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAppleArchive"); err == nil && ptr != 0 {
		UTTypes.AppleArchive = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAppleProtectedMPEG4Audio"); err == nil && ptr != 0 {
		UTTypes.AppleProtectedMPEG4Audio = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAppleProtectedMPEG4Video"); err == nil && ptr != 0 {
		UTTypes.AppleProtectedMPEG4Video = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAppleScript"); err == nil && ptr != 0 {
		UTTypes.AppleScript = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeApplication"); err == nil && ptr != 0 {
		UTTypes.Application = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeApplicationBundle"); err == nil && ptr != 0 {
		UTTypes.ApplicationBundle = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeApplicationExtension"); err == nil && ptr != 0 {
		UTTypes.ApplicationExtension = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeArchive"); err == nil && ptr != 0 {
		UTTypes.Archive = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAssemblyLanguageSource"); err == nil && ptr != 0 {
		UTTypes.AssemblyLanguageSource = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAudio"); err == nil && ptr != 0 {
		UTTypes.Audio = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeAudiovisualContent"); err == nil && ptr != 0 {
		UTTypes.AudiovisualContent = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeBMP"); err == nil && ptr != 0 {
		UTTypes.BMP = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeBZ2"); err == nil && ptr != 0 {
		UTTypes.BZ2 = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeBinaryPropertyList"); err == nil && ptr != 0 {
		UTTypes.BinaryPropertyList = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeBookmark"); err == nil && ptr != 0 {
		UTTypes.Bookmark = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeBundle"); err == nil && ptr != 0 {
		UTTypes.Bundle = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeCHeader"); err == nil && ptr != 0 {
		UTTypes.CHeader = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeCPlusPlusHeader"); err == nil && ptr != 0 {
		UTTypes.CPlusPlusHeader = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeCPlusPlusSource"); err == nil && ptr != 0 {
		UTTypes.CPlusPlusSource = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeCSS"); err == nil && ptr != 0 {
		UTTypes.CSS = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeCSource"); err == nil && ptr != 0 {
		UTTypes.CSource = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeCalendarEvent"); err == nil && ptr != 0 {
		UTTypes.CalendarEvent = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeCommaSeparatedText"); err == nil && ptr != 0 {
		UTTypes.CommaSeparatedText = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeCompositeContent"); err == nil && ptr != 0 {
		UTTypes.CompositeContent = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeContact"); err == nil && ptr != 0 {
		UTTypes.Contact = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeContent"); err == nil && ptr != 0 {
		UTTypes.Content = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeDNG"); err == nil && ptr != 0 {
		UTTypes.DNG = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeData"); err == nil && ptr != 0 {
		UTTypes.Data = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeDatabase"); err == nil && ptr != 0 {
		UTTypes.Database = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeDelimitedText"); err == nil && ptr != 0 {
		UTTypes.DelimitedText = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeDirectory"); err == nil && ptr != 0 {
		UTTypes.Directory = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeDiskImage"); err == nil && ptr != 0 {
		UTTypes.DiskImage = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeEPUB"); err == nil && ptr != 0 {
		UTTypes.EPUB = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeEXE"); err == nil && ptr != 0 {
		UTTypes.EXE = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeEXR"); err == nil && ptr != 0 {
		UTTypes.EXR = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeEmailMessage"); err == nil && ptr != 0 {
		UTTypes.EmailMessage = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeExecutable"); err == nil && ptr != 0 {
		UTTypes.Executable = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeFileURL"); err == nil && ptr != 0 {
		UTTypes.FileURL = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeFlatRTFD"); err == nil && ptr != 0 {
		UTTypes.FlatRTFD = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeFolder"); err == nil && ptr != 0 {
		UTTypes.Folder = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeFont"); err == nil && ptr != 0 {
		UTTypes.Font = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeFramework"); err == nil && ptr != 0 {
		UTTypes.Framework = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeGIF"); err == nil && ptr != 0 {
		UTTypes.GIF = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeGZIP"); err == nil && ptr != 0 {
		UTTypes.GZIP = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeGeoJSON"); err == nil && ptr != 0 {
		UTTypes.GeoJSON = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeHEIC"); err == nil && ptr != 0 {
		UTTypes.HEIC = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeHEICS"); err == nil && ptr != 0 {
		UTTypes.HEICS = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeHEIF"); err == nil && ptr != 0 {
		UTTypes.HEIF = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeHTML"); err == nil && ptr != 0 {
		UTTypes.HTML = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeICNS"); err == nil && ptr != 0 {
		UTTypes.ICNS = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeICO"); err == nil && ptr != 0 {
		UTTypes.ICO = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeImage"); err == nil && ptr != 0 {
		UTTypes.Image = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeInternetLocation"); err == nil && ptr != 0 {
		UTTypes.InternetLocation = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeInternetShortcut"); err == nil && ptr != 0 {
		UTTypes.InternetShortcut = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeItem"); err == nil && ptr != 0 {
		UTTypes.Item = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeJPEG"); err == nil && ptr != 0 {
		UTTypes.JPEG = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeJPEGXL"); err == nil && ptr != 0 {
		UTTypes.JPEGXL = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeJSON"); err == nil && ptr != 0 {
		UTTypes.JSON = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeJavaScript"); err == nil && ptr != 0 {
		UTTypes.JavaScript = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeLinkPresentationMetadata"); err == nil && ptr != 0 {
		UTTypes.LinkPresentationMetadata = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeLivePhoto"); err == nil && ptr != 0 {
		UTTypes.LivePhoto = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeLog"); err == nil && ptr != 0 {
		UTTypes.Log = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeM3UPlaylist"); err == nil && ptr != 0 {
		UTTypes.M3UPlaylist = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMIDI"); err == nil && ptr != 0 {
		UTTypes.MIDI = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMP3"); err == nil && ptr != 0 {
		UTTypes.MP3 = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMPEG"); err == nil && ptr != 0 {
		UTTypes.MPEG = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMPEG2TransportStream"); err == nil && ptr != 0 {
		UTTypes.MPEG2TransportStream = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMPEG2Video"); err == nil && ptr != 0 {
		UTTypes.MPEG2Video = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMPEG4Audio"); err == nil && ptr != 0 {
		UTTypes.MPEG4Audio = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMPEG4Movie"); err == nil && ptr != 0 {
		UTTypes.MPEG4Movie = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMakefile"); err == nil && ptr != 0 {
		UTTypes.Makefile = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMessage"); err == nil && ptr != 0 {
		UTTypes.Message = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMountPoint"); err == nil && ptr != 0 {
		UTTypes.MountPoint = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeMovie"); err == nil && ptr != 0 {
		UTTypes.Movie = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeOSAScript"); err == nil && ptr != 0 {
		UTTypes.OSAScript = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeOSAScriptBundle"); err == nil && ptr != 0 {
		UTTypes.OSAScriptBundle = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeObjectiveCPlusPlusSource"); err == nil && ptr != 0 {
		UTTypes.ObjectiveCPlusPlusSource = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeObjectiveCSource"); err == nil && ptr != 0 {
		UTTypes.ObjectiveCSource = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePDF"); err == nil && ptr != 0 {
		UTTypes.PDF = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePHPScript"); err == nil && ptr != 0 {
		UTTypes.PHPScript = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePKCS12"); err == nil && ptr != 0 {
		UTTypes.PKCS12 = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePNG"); err == nil && ptr != 0 {
		UTTypes.PNG = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePackage"); err == nil && ptr != 0 {
		UTTypes.Package = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePerlScript"); err == nil && ptr != 0 {
		UTTypes.PerlScript = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePlainText"); err == nil && ptr != 0 {
		UTTypes.PlainText = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePlaylist"); err == nil && ptr != 0 {
		UTTypes.Playlist = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePluginBundle"); err == nil && ptr != 0 {
		UTTypes.PluginBundle = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePresentation"); err == nil && ptr != 0 {
		UTTypes.Presentation = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePropertyList"); err == nil && ptr != 0 {
		UTTypes.PropertyList = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypePythonScript"); err == nil && ptr != 0 {
		UTTypes.PythonScript = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeQuickLookGenerator"); err == nil && ptr != 0 {
		UTTypes.QuickLookGenerator = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeQuickTimeMovie"); err == nil && ptr != 0 {
		UTTypes.QuickTimeMovie = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeRAWImage"); err == nil && ptr != 0 {
		UTTypes.RAWImage = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeRTF"); err == nil && ptr != 0 {
		UTTypes.RTF = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeRTFD"); err == nil && ptr != 0 {
		UTTypes.RTFD = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeRealityFile"); err == nil && ptr != 0 {
		UTTypes.RealityFile = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeResolvable"); err == nil && ptr != 0 {
		UTTypes.Resolvable = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeRubyScript"); err == nil && ptr != 0 {
		UTTypes.RubyScript = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeSVG"); err == nil && ptr != 0 {
		UTTypes.SVG = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeSceneKitScene"); err == nil && ptr != 0 {
		UTTypes.SceneKitScene = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeScript"); err == nil && ptr != 0 {
		UTTypes.Script = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeShellScript"); err == nil && ptr != 0 {
		UTTypes.ShellScript = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeSourceCode"); err == nil && ptr != 0 {
		UTTypes.SourceCode = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeSpotlightImporter"); err == nil && ptr != 0 {
		UTTypes.SpotlightImporter = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeSpreadsheet"); err == nil && ptr != 0 {
		UTTypes.Spreadsheet = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeSwiftSource"); err == nil && ptr != 0 {
		UTTypes.SwiftSource = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeSymbolicLink"); err == nil && ptr != 0 {
		UTTypes.SymbolicLink = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeSystemPreferencesPane"); err == nil && ptr != 0 {
		UTTypes.SystemPreferencesPane = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeTIFF"); err == nil && ptr != 0 {
		UTTypes.TIFF = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeTabSeparatedText"); err == nil && ptr != 0 {
		UTTypes.TabSeparatedText = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeTarArchive"); err == nil && ptr != 0 {
		UTTypes.TarArchive = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeText"); err == nil && ptr != 0 {
		UTTypes.Text = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeToDoItem"); err == nil && ptr != 0 {
		UTTypes.ToDoItem = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeURL"); err == nil && ptr != 0 {
		UTTypes.URL = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeURLBookmarkData"); err == nil && ptr != 0 {
		UTTypes.URLBookmarkData = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeUSD"); err == nil && ptr != 0 {
		UTTypes.USD = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeUSDZ"); err == nil && ptr != 0 {
		UTTypes.USDZ = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeUTF16ExternalPlainText"); err == nil && ptr != 0 {
		UTTypes.UTF16ExternalPlainText = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeUTF16PlainText"); err == nil && ptr != 0 {
		UTTypes.UTF16PlainText = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeUTF8PlainText"); err == nil && ptr != 0 {
		UTTypes.UTF8PlainText = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeUTF8TabSeparatedText"); err == nil && ptr != 0 {
		UTTypes.UTF8TabSeparatedText = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeUnixExecutable"); err == nil && ptr != 0 {
		UTTypes.UnixExecutable = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeVCard"); err == nil && ptr != 0 {
		UTTypes.VCard = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeVideo"); err == nil && ptr != 0 {
		UTTypes.Video = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeVolume"); err == nil && ptr != 0 {
		UTTypes.Volume = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeWAV"); err == nil && ptr != 0 {
		UTTypes.WAV = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeWebArchive"); err == nil && ptr != 0 {
		UTTypes.WebArchive = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeWebP"); err == nil && ptr != 0 {
		UTTypes.WebP = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeX509Certificate"); err == nil && ptr != 0 {
		UTTypes.X509Certificate = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeXML"); err == nil && ptr != 0 {
		UTTypes.XML = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeXMLPropertyList"); err == nil && ptr != 0 {
		UTTypes.XMLPropertyList = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeXPCService"); err == nil && ptr != 0 {
		UTTypes.XPCService = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeYAML"); err == nil && ptr != 0 {
		UTTypes.YAML = *(*UTType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UTTypeZIP"); err == nil && ptr != 0 {
		UTTypes.ZIP = *(*UTType)(unsafe.Pointer(ptr))
	}

}

// UTTypes provides typed accessors for [UTType] constants.
var UTTypes struct {
	AHAP UTType
	// AIFF: A type that represents data in AIFF audio format.
	AIFF UTType
	// ARReferenceObject: A type that represents an augmented reality reference object.
	ARReferenceObject UTType
	// AVI: A type that represents data in AVI movie format.
	AVI UTType
	// AliasFile: A type that represents an alias file.
	AliasFile UTType
	// AppleArchive: A type that represents an Apple archive of files and directories.
	AppleArchive UTType
	// AppleProtectedMPEG4Audio: A type that represents data in Apple-protected MPEG-4 format.
	AppleProtectedMPEG4Audio UTType
	// AppleProtectedMPEG4Video: A type that represents data in Apple-protected MPEG-4 format.
	AppleProtectedMPEG4Video UTType
	// AppleScript: A type that represents an AppleScript text-based script.
	AppleScript UTType
	// Application: A base type that represents a macOS, iOS, iPadOS, watchOS, and tvOS app.
	Application UTType
	// ApplicationBundle: A type that represents a bundled app.
	ApplicationBundle UTType
	// ApplicationExtension: A type that represents an app extension.
	ApplicationExtension UTType
	// Archive: A base type that represents an archive of files and directories.
	Archive UTType
	// AssemblyLanguageSource: A type that represents assembly language source code.
	AssemblyLanguageSource UTType
	// Audio: A type that represents audio that doesn’t contain video.
	Audio UTType
	// AudiovisualContent: A base type that represents data that contains video content that may or may not also include audio.
	AudiovisualContent UTType
	// BMP: A type that represents a Windows bitmap image.
	BMP UTType
	// BZ2: A type that represents a bzip2 archive.
	BZ2 UTType
	// BinaryPropertyList: A type that represents a binary property list.
	BinaryPropertyList UTType
	// Bookmark: A base type that represents bookmark data.
	Bookmark UTType
	// Bundle: A base type that represents a directory that conforms to one of the bundle layouts.
	Bundle UTType
	// CHeader: A type that represents a C header file.
	CHeader UTType
	// CPlusPlusHeader: A type that represents a C++ header file.
	CPlusPlusHeader UTType
	// CPlusPlusSource: A type that represents a C++ source code file.
	CPlusPlusSource UTType
	CSS UTType
	// CSource: A type that represents a C source code file.
	CSource UTType
	// CalendarEvent: A base type that represents a calendar event.
	CalendarEvent UTType
	// CommaSeparatedText: A type that represents text containing comma-separated values.
	CommaSeparatedText UTType
	// CompositeContent: A base type that represents a content format supporting mixed embedded content.
	CompositeContent UTType
	// Contact: A base type that represents contact information.
	Contact UTType
	// Content: A base type that represents anything containing user-viewable content.
	Content UTType
	DNG UTType
	// Data: A base type that represents any sort of byte stream, including files and in-memory data.
	Data UTType
	// Database: A base type that represents a database store.
	Database UTType
	// DelimitedText: A base type that represents text containing delimited values.
	DelimitedText UTType
	// Directory: A type that represents a file system directory, including packages and folders.
	Directory UTType
	// DiskImage: A type that represents a data item that’s mountable as a volume.
	DiskImage UTType
	// EPUB: A type that represents data in the electronic publication (EPUB) format.
	EPUB UTType
	// EXE: A type that represents a Windows executable.
	EXE UTType
	EXR UTType
	// EmailMessage: A type that represents an email message.
	EmailMessage UTType
	// Executable: A type that represents an executable.
	Executable UTType
	// FileURL: A type that represents a URL to a file in the file system.
	FileURL UTType
	// FlatRTFD: A type that represents flattened Rich Text Format Directory documents.
	FlatRTFD UTType
	// Folder: A type that represents a user-browsable directory.
	Folder UTType
	// Font: A base type that represents a font.
	Font UTType
	// Framework: A type that represents an Apple framework bundle.
	Framework UTType
	// GIF: A type that represents a GIF image.
	GIF UTType
	// GZIP: A type that represents a GNU zip archive.
	GZIP UTType
	GeoJSON UTType
	// HEIC: A type that represents High Efficiency Image Coding images.
	HEIC UTType
	HEICS UTType
	// HEIF: A type that represents High Efficiency Image File Format images.
	HEIF UTType
	// HTML: A type that represents any version of HTML.
	HTML UTType
	// ICNS: A type that represents Apple icon data.
	ICNS UTType
	// ICO: A type that represents Windows icon data.
	ICO UTType
	// Image: A base type that represents image data.
	Image UTType
	// InternetLocation: A base type that represents an Apple internet location file.
	InternetLocation UTType
	// InternetShortcut: A type that represents a Microsoft internet shortcut file.
	InternetShortcut UTType
	// Item: A generic base type for most objects, such as files or directories.
	Item UTType
	// JPEG: A type that represents a JPEG image.
	JPEG UTType
	JPEGXL UTType
	// JSON: A type that represents JavaScript Object Notation (JSON) data.
	JSON UTType
	// JavaScript: A type that represents JavaScript source code.
	JavaScript UTType
	LinkPresentationMetadata UTType
	// LivePhoto: A type that represents Live Photos.
	LivePhoto UTType
	// Log: A base type that represents console log data.
	Log UTType
	// M3UPlaylist: A type that represents an M3U or M3U8 playlist.
	M3UPlaylist UTType
	// MIDI: A type that represents data in MIDI audio format.
	MIDI UTType
	// MP3: A type that represents MP3 audio.
	MP3 UTType
	// MPEG: A type that represents an MPEG-1 or MPEG-2 movie.
	MPEG UTType
	// MPEG2TransportStream: A type that represents data in MPEG-2 transport stream movie format.
	MPEG2TransportStream UTType
	// MPEG2Video: A type that represents an MPEG-2 video.
	MPEG2Video UTType
	// MPEG4Audio: A type that represents an MPEG-4 audio layer file.
	MPEG4Audio UTType
	// MPEG4Movie: A type that represents an MPEG-4 movie.
	MPEG4Movie UTType
	// Makefile: A type that represents a Makefile.
	Makefile UTType
	// Message: A base type that represents a message.
	Message UTType
	// MountPoint: A type that represents a volume mount point.
	MountPoint UTType
	// Movie: A base type representing media formats that may contain both video and audio.
	Movie UTType
	// OSAScript: A type that represents an Open Scripting Architecture binary script.
	OSAScript UTType
	// OSAScriptBundle: A type that represents an Open Scripting Architecture script bundle.
	OSAScriptBundle UTType
	// ObjectiveCPlusPlusSource: A type that represents an Objective-C++ source code file.
	ObjectiveCPlusPlusSource UTType
	// ObjectiveCSource: A type that represents an Objective-C source code file.
	ObjectiveCSource UTType
	// PDF: A type that represents Adobe Portable Document Format (PDF) documents.
	PDF UTType
	// PHPScript: A type that represents a PHP script.
	PHPScript UTType
	// PKCS12: A type that represents Public Key Cryptography Standard (PKCS) 12 data.
	PKCS12 UTType
	// PNG: A type that represents a PNG image.
	PNG UTType
	// Package: A base type that represents a packaged directory.
	Package UTType
	// PerlScript: A type that represents a Perl script.
	PerlScript UTType
	// PlainText: A type that represents text with no markup and an unspecified encoding.
	PlainText UTType
	// Playlist: A base type that represents a playlist.
	Playlist UTType
	// PluginBundle: A base type that represents a bundle-based plug-in.
	PluginBundle UTType
	// Presentation: A base type that represents a presentation document.
	Presentation UTType
	// PropertyList: A base type that represents a property list.
	PropertyList UTType
	// PythonScript: A type that represents a Python script.
	PythonScript UTType
	// QuickLookGenerator: A type that represents a QuickLook preview generator bundle.
	QuickLookGenerator UTType
	// QuickTimeMovie: A type that represents a QuickTime movie.
	QuickTimeMovie UTType
	// RAWImage: A base type that represents a raw image format that you use in digital photography.
	RAWImage UTType
	// RTF: A type that represents Rich Text Format data.
	RTF UTType
	// RTFD: A type that represents Rich Text Format Directory documents.
	RTFD UTType
	// RealityFile: A type that represents a Reality Composer file.
	RealityFile UTType
	// Resolvable: A base type that represents a resolvable reference, including symbolic links and aliases.
	Resolvable UTType
	// RubyScript: A type that represents a Ruby script.
	RubyScript UTType
	// SVG: A type that represents a scalable vector graphics (SVG) image.
	SVG UTType
	// SceneKitScene: A type that represents a SceneKit serialized scene.
	SceneKitScene UTType
	// Script: A base type that represents any scripting language source.
	Script UTType
	// ShellScript: A base type that represents a shell script.
	ShellScript UTType
	// SourceCode: A base type that represents source code of any programming language.
	SourceCode UTType
	// SpotlightImporter: A type that represents a Spotlight metadata importer bundle.
	SpotlightImporter UTType
	// Spreadsheet: A base type that represents a spreadsheet document.
	Spreadsheet UTType
	// SwiftSource: A type that represents a Swift source code file.
	SwiftSource UTType
	// SymbolicLink: A type that represents a symbolic link.
	SymbolicLink UTType
	// SystemPreferencesPane: A type that represents a System Preferences pane.
	SystemPreferencesPane UTType
	// TIFF: A type that represents a TIFF image.
	TIFF UTType
	// TabSeparatedText: A type that represents text containing tab-separated values.
	TabSeparatedText UTType
	TarArchive UTType
	// Text: A base type that represents all text-encoded data, including text with markup.
	Text UTType
	// ToDoItem: A type that represents a to-do item.
	ToDoItem UTType
	// URL: A type that represents a URL.
	URL UTType
	// URLBookmarkData: A type that represents a URL bookmark.
	URLBookmarkData UTType
	// USD: A type that represents Universal Scene Description content.
	USD UTType
	// USDZ: A type that represents Universal Scene Description Package content.
	USDZ UTType
	// UTF16ExternalPlainText: A type that represents plain text encoded as UTF-16 with an optional BOM.
	UTF16ExternalPlainText UTType
	// UTF16PlainText: A type that represents plain text encoded as UTF-16 in native byte order with an optional bill of materials.
	UTF16PlainText UTType
	// UTF8PlainText: A type that represents plain text encoded as UTF-8.
	UTF8PlainText UTType
	// UTF8TabSeparatedText: A type that represents UTF-8–encoded text containing tab-separated values.
	UTF8TabSeparatedText UTType
	// UnixExecutable: A type that represents a UNIX executable.
	UnixExecutable UTType
	// VCard: A type that represents a vCard file.
	VCard UTType
	// Video: A type that represents video that doesn’t contain audio.
	Video UTType
	// Volume: A type that represents the root folder of a volume or mount point.
	Volume UTType
	// WAV: A type that represents data in Microsoft Waveform Audio File Format.
	WAV UTType
	// WebArchive: A type that represents WebKit web archive data.
	WebArchive UTType
	// WebP: A type that represents a WebP image.
	WebP UTType
	// X509Certificate: A type that represents an X.509 certificate.
	X509Certificate UTType
	// XML: A type that represents generic XML data.
	XML UTType
	// XMLPropertyList: A type that represents an XML property list.
	XMLPropertyList UTType
	// XPCService: A type that represents an XPC service bundle.
	XPCService UTType
	// YAML: A type that represents Yet Another Markup Language data.
	YAML UTType
	// ZIP: A type that represents a zip archive.
	ZIP UTType
}

