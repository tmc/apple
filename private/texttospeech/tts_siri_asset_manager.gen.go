// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSSiriAssetManager] class.
var (
	_TTSSiriAssetManagerClass     TTSSiriAssetManagerClass
	_TTSSiriAssetManagerClassOnce sync.Once
)

func getTTSSiriAssetManagerClass() TTSSiriAssetManagerClass {
	_TTSSiriAssetManagerClassOnce.Do(func() {
		_TTSSiriAssetManagerClass = TTSSiriAssetManagerClass{class: objc.GetClass("TTSSiriAssetManager")}
	})
	return _TTSSiriAssetManagerClass
}

// GetTTSSiriAssetManagerClass returns the class object for TTSSiriAssetManager.
func GetTTSSiriAssetManagerClass() TTSSiriAssetManagerClass {
	return getTTSSiriAssetManagerClass()
}

type TTSSiriAssetManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSSiriAssetManagerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSSiriAssetManagerClass) Alloc() TTSSiriAssetManager {
	rv := objc.Send[TTSSiriAssetManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager
type TTSSiriAssetManager struct {
	objectivec.Object
}

// TTSSiriAssetManagerFromID constructs a [TTSSiriAssetManager] from an objc.ID.
func TTSSiriAssetManagerFromID(id objc.ID) TTSSiriAssetManager {
	return TTSSiriAssetManager{objectivec.Object{ID: id}}
}

// Ensure TTSSiriAssetManager implements ITTSSiriAssetManager.
var _ ITTSSiriAssetManager = TTSSiriAssetManager{}

// An interface definition for the [TTSSiriAssetManager] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager
type ITTSSiriAssetManager interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TTSSiriAssetManager) Init() TTSSiriAssetManager {
	rv := objc.Send[TTSSiriAssetManager](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSSiriAssetManager) Autorelease() TTSSiriAssetManager {
	rv := objc.Send[TTSSiriAssetManager](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSSiriAssetManager creates a new TTSSiriAssetManager instance.
func NewTTSSiriAssetManager() TTSSiriAssetManager {
	class := getTTSSiriAssetManagerClass()
	rv := objc.Send[TTSSiriAssetManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/_assetFilterForLanguage:gender:footprint:voiceName:voiceType:locallyAvailable:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) _assetFilterForLanguageGenderFootprintVoiceNameVoiceTypeLocallyAvailable(language objectivec.IObject, gender int64, footprint int64, name objectivec.IObject, type_ int64, available bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("_assetFilterForLanguage:gender:footprint:voiceName:voiceType:locallyAvailable:"), language, gender, footprint, name, type_, available)
	return objectivec.Object{ID: rv}
}

// AssetFilterForLanguageGenderFootprintVoiceNameVoiceTypeLocallyAvailable is an exported wrapper for the private method _assetFilterForLanguageGenderFootprintVoiceNameVoiceTypeLocallyAvailable.
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) AssetFilterForLanguageGenderFootprintVoiceNameVoiceTypeLocallyAvailable(language objectivec.IObject, gender int64, footprint int64, name objectivec.IObject, type_ int64, available bool) objectivec.IObject {
	return _TTSSiriAssetManagerClass._assetFilterForLanguageGenderFootprintVoiceNameVoiceTypeLocallyAvailable(language, gender, footprint, name, type_, available)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/_assetTechnologyForVoiceType:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) _assetTechnologyForVoiceType(type_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("_assetTechnologyForVoiceType:"), type_)
	return objectivec.Object{ID: rv}
}

// AssetTechnologyForVoiceType is an exported wrapper for the private method _assetTechnologyForVoiceType.
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) AssetTechnologyForVoiceType(type_ int64) objectivec.IObject {
	return _TTSSiriAssetManagerClass._assetTechnologyForVoiceType(type_)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/_assetTypesForVoiceType:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) _assetTypesForVoiceType(type_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("_assetTypesForVoiceType:"), type_)
	return objectivec.Object{ID: rv}
}

// AssetTypesForVoiceType is an exported wrapper for the private method _assetTypesForVoiceType.
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) AssetTypesForVoiceType(type_ int64) objectivec.IObject {
	return _TTSSiriAssetManagerClass._assetTypesForVoiceType(type_)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/_assetsForLanguage:voiceType:installedOnly:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) _assetsForLanguageVoiceTypeInstalledOnly(language objectivec.IObject, type_ int64, only bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("_assetsForLanguage:voiceType:installedOnly:"), language, type_, only)
	return objectivec.Object{ID: rv}
}

// AssetsForLanguageVoiceTypeInstalledOnly is an exported wrapper for the private method _assetsForLanguageVoiceTypeInstalledOnly.
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) AssetsForLanguageVoiceTypeInstalledOnly(language objectivec.IObject, type_ int64, only bool) objectivec.IObject {
	return _TTSSiriAssetManagerClass._assetsForLanguageVoiceTypeInstalledOnly(language, type_, only)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/_footprintForQuality:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) _footprintForQuality(quality objectivec.IObject) int64 {
	rv := objc.Send[int64](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("_footprintForQuality:"), quality)
	return rv
}

// FootprintForQuality is an exported wrapper for the private method _footprintForQuality.
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) FootprintForQuality(quality objectivec.IObject) int64 {
	return _TTSSiriAssetManagerClass._footprintForQuality(quality)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/_footprintForType:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) _footprintForType(type_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("_footprintForType:"), type_)
	return objectivec.Object{ID: rv}
}

// FootprintForType is an exported wrapper for the private method _footprintForType.
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) FootprintForType(type_ int64) objectivec.IObject {
	return _TTSSiriAssetManagerClass._footprintForType(type_)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/_siriAssetForLanguage:gender:footprint:voiceName:voiceType:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) _siriAssetForLanguageGenderFootprintVoiceNameVoiceType(language objectivec.IObject, gender int64, footprint int64, name objectivec.IObject, type_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("_siriAssetForLanguage:gender:footprint:voiceName:voiceType:"), language, gender, footprint, name, type_)
	return objectivec.Object{ID: rv}
}

// SiriAssetForLanguageGenderFootprintVoiceNameVoiceType is an exported wrapper for the private method _siriAssetForLanguageGenderFootprintVoiceNameVoiceType.
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) SiriAssetForLanguageGenderFootprintVoiceNameVoiceType(language objectivec.IObject, gender int64, footprint int64, name objectivec.IObject, type_ int64) objectivec.IObject {
	return _TTSSiriAssetManagerClass._siriAssetForLanguageGenderFootprintVoiceNameVoiceType(language, gender, footprint, name, type_)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/_siriAssetForLanguage:gender:footprint:voiceName:voiceType:locallyAvailable:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) _siriAssetForLanguageGenderFootprintVoiceNameVoiceTypeLocallyAvailable(language objectivec.IObject, gender int64, footprint int64, name objectivec.IObject, type_ int64, available bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("_siriAssetForLanguage:gender:footprint:voiceName:voiceType:locallyAvailable:"), language, gender, footprint, name, type_, available)
	return objectivec.Object{ID: rv}
}

// SiriAssetForLanguageGenderFootprintVoiceNameVoiceTypeLocallyAvailable is an exported wrapper for the private method _siriAssetForLanguageGenderFootprintVoiceNameVoiceTypeLocallyAvailable.
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) SiriAssetForLanguageGenderFootprintVoiceNameVoiceTypeLocallyAvailable(language objectivec.IObject, gender int64, footprint int64, name objectivec.IObject, type_ int64, available bool) objectivec.IObject {
	return _TTSSiriAssetManagerClass._siriAssetForLanguageGenderFootprintVoiceNameVoiceTypeLocallyAvailable(language, gender, footprint, name, type_, available)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/_voiceTypeForAssetTechnology:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) _voiceTypeForAssetTechnology(technology objectivec.IObject) int64 {
	rv := objc.Send[int64](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("_voiceTypeForAssetTechnology:"), technology)
	return rv
}

// VoiceTypeForAssetTechnology is an exported wrapper for the private method _voiceTypeForAssetTechnology.
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) VoiceTypeForAssetTechnology(technology objectivec.IObject) int64 {
	return _TTSSiriAssetManagerClass._voiceTypeForAssetTechnology(technology)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/assetForLanguage:gender:footprint:voiceName:voiceType:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) AssetForLanguageGenderFootprintVoiceNameVoiceType(language objectivec.IObject, gender int64, footprint int64, name objectivec.IObject, type_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("assetForLanguage:gender:footprint:voiceName:voiceType:"), language, gender, footprint, name, type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/assetsForLanguage:voiceType:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) AssetsForLanguageVoiceType(language objectivec.IObject, type_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("assetsForLanguage:voiceType:"), language, type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/convertTTSLanguageCodeToSiriLanguageCode:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) ConvertTTSLanguageCodeToSiriLanguageCode(code objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("convertTTSLanguageCodeToSiriLanguageCode:"), code)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/deprecatedVoicesMap
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) DeprecatedVoicesMap() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("deprecatedVoicesMap"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/downloadAsset:progressHandler:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) DownloadAssetProgressHandler(asset objectivec.IObject, handler VoidHandler) {
	_block1, _ := NewVoidBlock(handler)
	objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("downloadAsset:progressHandler:"), asset, _block1)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/downloadVoiceResourceForLanguage:completion:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) DownloadVoiceResourceForLanguageCompletion(language objectivec.IObject, completion VoidHandler) {
	_block1, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("downloadVoiceResourceForLanguage:completion:"), language, _block1)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/installedAssetForLanguage:gender:footprint:voiceName:voiceType:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) InstalledAssetForLanguageGenderFootprintVoiceNameVoiceType(language objectivec.IObject, gender int64, footprint int64, name objectivec.IObject, type_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("installedAssetForLanguage:gender:footprint:voiceName:voiceType:"), language, gender, footprint, name, type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/installedAssetsForLanguage:voiceType:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) InstalledAssetsForLanguageVoiceType(language objectivec.IObject, type_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("installedAssetsForLanguage:voiceType:"), language, type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/installedVoiceResourceForLanguage:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) InstalledVoiceResourceForLanguage(language objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("installedVoiceResourceForLanguage:"), language)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/purgeAsset:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) PurgeAsset(asset objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("purgeAsset:"), asset)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/spaceCheck:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) SpaceCheck(check objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("spaceCheck:"), check)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/stopDownload:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) StopDownload(download objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("stopDownload:"), download)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/ttsAssetFromVoiceAsset:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) TtsAssetFromVoiceAsset(asset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("ttsAssetFromVoiceAsset:"), asset)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/voiceAssetFromTTSAsset:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) VoiceAssetFromTTSAsset(tTSAsset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("voiceAssetFromTTSAsset:"), tTSAsset)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/voiceIdentifierForAsset:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) VoiceIdentifierForAsset(asset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("voiceIdentifierForAsset:"), asset)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/voiceIdentifierForType:footprint:language:name:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) VoiceIdentifierForTypeFootprintLanguageName(type_ int64, footprint int64, language objectivec.IObject, name objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("voiceIdentifierForType:footprint:language:name:"), type_, footprint, language, name)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSiriAssetManager/voiceResourceForLanguage:voiceType:
func (_TTSSiriAssetManagerClass TTSSiriAssetManagerClass) VoiceResourceForLanguageVoiceType(language objectivec.IObject, type_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSiriAssetManagerClass.class), objc.Sel("voiceResourceForLanguage:voiceType:"), language, type_)
	return objectivec.Object{ID: rv}
}

// DownloadAssetProgressHandlerSync is a synchronous wrapper around [TTSSiriAssetManager.DownloadAssetProgressHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (tc TTSSiriAssetManagerClass) DownloadAssetProgressHandlerSync(ctx context.Context, asset objectivec.IObject) error {
	done := make(chan struct{}, 1)
	tc.DownloadAssetProgressHandler(asset, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// DownloadVoiceResourceForLanguageCompletionSync is a synchronous wrapper around [TTSSiriAssetManager.DownloadVoiceResourceForLanguageCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (tc TTSSiriAssetManagerClass) DownloadVoiceResourceForLanguageCompletionSync(ctx context.Context, language objectivec.IObject) error {
	done := make(chan struct{}, 1)
	tc.DownloadVoiceResourceForLanguageCompletion(language, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
