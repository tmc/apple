// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSpeechSynthesisVoice] class.
var (
	_AVSpeechSynthesisVoiceClass     AVSpeechSynthesisVoiceClass
	_AVSpeechSynthesisVoiceClassOnce sync.Once
)

func getAVSpeechSynthesisVoiceClass() AVSpeechSynthesisVoiceClass {
	_AVSpeechSynthesisVoiceClassOnce.Do(func() {
		_AVSpeechSynthesisVoiceClass = AVSpeechSynthesisVoiceClass{class: objc.GetClass("AVSpeechSynthesisVoice")}
	})
	return _AVSpeechSynthesisVoiceClass
}

// GetAVSpeechSynthesisVoiceClass returns the class object for AVSpeechSynthesisVoice.
func GetAVSpeechSynthesisVoiceClass() AVSpeechSynthesisVoiceClass {
	return getAVSpeechSynthesisVoiceClass()
}

type AVSpeechSynthesisVoiceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSpeechSynthesisVoiceClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSpeechSynthesisVoiceClass) Alloc() AVSpeechSynthesisVoice {
	rv := objc.Send[AVSpeechSynthesisVoice](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVSpeechSynthesisVoice.AssetSize]
//   - [AVSpeechSynthesisVoice.BackupName]
//   - [AVSpeechSynthesisVoice.CanBeDownloaded]
//   - [AVSpeechSynthesisVoice.CoreVoiceIsSiriVoice]
//   - [AVSpeechSynthesisVoice.CoreVoiceIsSystemVoice]
//   - [AVSpeechSynthesisVoice.CoreVoiceLocalizedName]
//   - [AVSpeechSynthesisVoice.CoreVoiceWrapper]
//   - [AVSpeechSynthesisVoice.SetCoreVoiceWrapper]
//   - [AVSpeechSynthesisVoice.IsDefault]
//   - [AVSpeechSynthesisVoice.IsInstalled]
//   - [AVSpeechSynthesisVoice.IsNoveltyVoice]
//   - [AVSpeechSynthesisVoice.SetIsNoveltyVoice]
//   - [AVSpeechSynthesisVoice.IsSiriVoice]
//   - [AVSpeechSynthesisVoice.IsSystemVoice]
//   - [AVSpeechSynthesisVoice.NameWithoutQuality]
//   - [AVSpeechSynthesisVoice.NonLocalizedName]
//   - [AVSpeechSynthesisVoice.NonLocalizedNameWithoutQuality]
//   - [AVSpeechSynthesisVoice.SetAssetSize]
//   - [AVSpeechSynthesisVoice.SetBackupName]
//   - [AVSpeechSynthesisVoice.SetCanBeDownloaded]
//   - [AVSpeechSynthesisVoice.SetIsDefault]
//   - [AVSpeechSynthesisVoice.SetIsInstalled]
//   - [AVSpeechSynthesisVoice.SiriDisplayName]
//   - [AVSpeechSynthesisVoice.SynthesisProviderVoice]
//   - [AVSpeechSynthesisVoice.SetSynthesisProviderVoice]
//   - [AVSpeechSynthesisVoice.InitWithCoder]
//   - [AVSpeechSynthesisVoice.InitWithLanguage]
//   - [AVSpeechSynthesisVoice.Gender]
//   - [AVSpeechSynthesisVoice.SetGender]
//   - [AVSpeechSynthesisVoice.Identifier]
//   - [AVSpeechSynthesisVoice.SetIdentifier]
//   - [AVSpeechSynthesisVoice.Language]
//   - [AVSpeechSynthesisVoice.SetLanguage]
//   - [AVSpeechSynthesisVoice.Name]
//   - [AVSpeechSynthesisVoice.SetName]
//   - [AVSpeechSynthesisVoice.Quality]
//   - [AVSpeechSynthesisVoice.SetQuality]
//   - [AVSpeechSynthesisVoice.VoiceTraits]
//   - [AVSpeechSynthesisVoice.SetVoiceTraits]
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice
type AVSpeechSynthesisVoice struct {
	objectivec.Object
}

// AVSpeechSynthesisVoiceFromID constructs a [AVSpeechSynthesisVoice] from an objc.ID.
func AVSpeechSynthesisVoiceFromID(id objc.ID) AVSpeechSynthesisVoice {
	return AVSpeechSynthesisVoice{objectivec.Object{ID: id}}
}
// Ensure AVSpeechSynthesisVoice implements IAVSpeechSynthesisVoice.
var _ IAVSpeechSynthesisVoice = AVSpeechSynthesisVoice{}

// An interface definition for the [AVSpeechSynthesisVoice] class.
//
// # Methods
//
//   - [IAVSpeechSynthesisVoice.AssetSize]
//   - [IAVSpeechSynthesisVoice.BackupName]
//   - [IAVSpeechSynthesisVoice.CanBeDownloaded]
//   - [IAVSpeechSynthesisVoice.CoreVoiceIsSiriVoice]
//   - [IAVSpeechSynthesisVoice.CoreVoiceIsSystemVoice]
//   - [IAVSpeechSynthesisVoice.CoreVoiceLocalizedName]
//   - [IAVSpeechSynthesisVoice.CoreVoiceWrapper]
//   - [IAVSpeechSynthesisVoice.SetCoreVoiceWrapper]
//   - [IAVSpeechSynthesisVoice.IsDefault]
//   - [IAVSpeechSynthesisVoice.IsInstalled]
//   - [IAVSpeechSynthesisVoice.IsNoveltyVoice]
//   - [IAVSpeechSynthesisVoice.SetIsNoveltyVoice]
//   - [IAVSpeechSynthesisVoice.IsSiriVoice]
//   - [IAVSpeechSynthesisVoice.IsSystemVoice]
//   - [IAVSpeechSynthesisVoice.NameWithoutQuality]
//   - [IAVSpeechSynthesisVoice.NonLocalizedName]
//   - [IAVSpeechSynthesisVoice.NonLocalizedNameWithoutQuality]
//   - [IAVSpeechSynthesisVoice.SetAssetSize]
//   - [IAVSpeechSynthesisVoice.SetBackupName]
//   - [IAVSpeechSynthesisVoice.SetCanBeDownloaded]
//   - [IAVSpeechSynthesisVoice.SetIsDefault]
//   - [IAVSpeechSynthesisVoice.SetIsInstalled]
//   - [IAVSpeechSynthesisVoice.SiriDisplayName]
//   - [IAVSpeechSynthesisVoice.SynthesisProviderVoice]
//   - [IAVSpeechSynthesisVoice.SetSynthesisProviderVoice]
//   - [IAVSpeechSynthesisVoice.InitWithCoder]
//   - [IAVSpeechSynthesisVoice.InitWithLanguage]
//   - [IAVSpeechSynthesisVoice.Gender]
//   - [IAVSpeechSynthesisVoice.SetGender]
//   - [IAVSpeechSynthesisVoice.Identifier]
//   - [IAVSpeechSynthesisVoice.SetIdentifier]
//   - [IAVSpeechSynthesisVoice.Language]
//   - [IAVSpeechSynthesisVoice.SetLanguage]
//   - [IAVSpeechSynthesisVoice.Name]
//   - [IAVSpeechSynthesisVoice.SetName]
//   - [IAVSpeechSynthesisVoice.Quality]
//   - [IAVSpeechSynthesisVoice.SetQuality]
//   - [IAVSpeechSynthesisVoice.VoiceTraits]
//   - [IAVSpeechSynthesisVoice.SetVoiceTraits]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice
type IAVSpeechSynthesisVoice interface {
	objectivec.IObject

	// Topic: Methods

	AssetSize() int64
	BackupName() objectivec.IObject
	CanBeDownloaded() bool
	CoreVoiceIsSiriVoice() bool
	CoreVoiceIsSystemVoice() bool
	CoreVoiceLocalizedName() string
	CoreVoiceWrapper() objectivec.Object
	SetCoreVoiceWrapper(value objectivec.Object)
	IsDefault() bool
	IsInstalled() bool
	IsNoveltyVoice() bool
	SetIsNoveltyVoice(value bool)
	IsSiriVoice() bool
	IsSystemVoice() bool
	NameWithoutQuality() string
	NonLocalizedName() string
	NonLocalizedNameWithoutQuality() string
	SetAssetSize(size int64)
	SetBackupName(name objectivec.IObject)
	SetCanBeDownloaded(downloaded bool)
	SetIsDefault(default_ bool)
	SetIsInstalled(installed bool)
	SiriDisplayName() objectivec.IObject
	SynthesisProviderVoice() IAVSpeechSynthesisProviderVoice
	SetSynthesisProviderVoice(value IAVSpeechSynthesisProviderVoice)
	InitWithCoder(coder foundation.INSCoder) AVSpeechSynthesisVoice
	InitWithLanguage(language objectivec.IObject) AVSpeechSynthesisVoice
	Gender() int64
	SetGender(value int64)
	Identifier() string
	SetIdentifier(value string)
	Language() string
	SetLanguage(value string)
	Name() string
	SetName(value string)
	Quality() int64
	SetQuality(value int64)
	VoiceTraits() uint64
	SetVoiceTraits(value uint64)
}

// Init initializes the instance.
func (s AVSpeechSynthesisVoice) Init() AVSpeechSynthesisVoice {
	rv := objc.Send[AVSpeechSynthesisVoice](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSpeechSynthesisVoice) Autorelease() AVSpeechSynthesisVoice {
	rv := objc.Send[AVSpeechSynthesisVoice](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSpeechSynthesisVoice creates a new AVSpeechSynthesisVoice instance.
func NewAVSpeechSynthesisVoice() AVSpeechSynthesisVoice {
	class := getAVSpeechSynthesisVoiceClass()
	rv := objc.Send[AVSpeechSynthesisVoice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/initWithCoder:
func NewSpeechSynthesisVoiceWithCoder(coder objectivec.IObject) AVSpeechSynthesisVoice {
	instance := getAVSpeechSynthesisVoiceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AVSpeechSynthesisVoiceFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/initWithLanguage:
func NewSpeechSynthesisVoiceWithLanguage(language objectivec.IObject) AVSpeechSynthesisVoice {
	instance := getAVSpeechSynthesisVoiceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLanguage:"), language)
	return AVSpeechSynthesisVoiceFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/assetSize
func (s AVSpeechSynthesisVoice) AssetSize() int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("assetSize"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/backupName
func (s AVSpeechSynthesisVoice) BackupName() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("backupName"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/canBeDownloaded
func (s AVSpeechSynthesisVoice) CanBeDownloaded() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("canBeDownloaded"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/isDefault
func (s AVSpeechSynthesisVoice) IsDefault() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isDefault"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/isInstalled
func (s AVSpeechSynthesisVoice) IsInstalled() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isInstalled"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/isSystemVoice
func (s AVSpeechSynthesisVoice) IsSystemVoice() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isSystemVoice"))
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/setAssetSize:
func (s AVSpeechSynthesisVoice) SetAssetSize(size int64) {
	objc.Send[objc.ID](s.ID, objc.Sel("setAssetSize:"), size)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/setBackupName:
func (s AVSpeechSynthesisVoice) SetBackupName(name objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("setBackupName:"), name)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/setCanBeDownloaded:
func (s AVSpeechSynthesisVoice) SetCanBeDownloaded(downloaded bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("setCanBeDownloaded:"), downloaded)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/setIsDefault:
func (s AVSpeechSynthesisVoice) SetIsDefault(default_ bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("setIsDefault:"), default_)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/setIsInstalled:
func (s AVSpeechSynthesisVoice) SetIsInstalled(installed bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("setIsInstalled:"), installed)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/setNonLocalizedNameWithoutQuality:
func (s AVSpeechSynthesisVoice) SetNonLocalizedNameWithoutQuality(quality objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("setNonLocalizedNameWithoutQuality:"), quality)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/siriDisplayName
func (s AVSpeechSynthesisVoice) SiriDisplayName() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("siriDisplayName"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/initWithCoder:
func (s AVSpeechSynthesisVoice) InitWithCoder(coder foundation.INSCoder) AVSpeechSynthesisVoice {
	rv := objc.Send[AVSpeechSynthesisVoice](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/initWithLanguage:
func (s AVSpeechSynthesisVoice) InitWithLanguage(language objectivec.IObject) AVSpeechSynthesisVoice {
	rv := objc.Send[AVSpeechSynthesisVoice](s.ID, objc.Sel("initWithLanguage:"), language)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/_speechVoicesIncludingSiri
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) _speechVoicesIncludingSiri() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_speechVoicesIncludingSiri"))
	return objectivec.Object{ID: rv}
}

// SpeechVoicesIncludingSiri is an exported wrapper for the private method _speechVoicesIncludingSiri.
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) SpeechVoicesIncludingSiri() objectivec.IObject {
	return _AVSpeechSynthesisVoiceClass._speechVoicesIncludingSiri()
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/_speechVoicesIncludingSiri:
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) _speechVoicesIncludingSiriWithSiri(siri bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_speechVoicesIncludingSiri:"), siri)
	return objectivec.Object{ID: rv}
}

// SpeechVoicesIncludingSiriWithSiri is an exported wrapper for the private method _speechVoicesIncludingSiriWithSiri.
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) SpeechVoicesIncludingSiriWithSiri(siri bool) objectivec.IObject {
	return _AVSpeechSynthesisVoiceClass._speechVoicesIncludingSiriWithSiri(siri)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/_speechVoicesIncludingSiriAndSuperCompact
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) _speechVoicesIncludingSiriAndSuperCompact() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_speechVoicesIncludingSiriAndSuperCompact"))
	return objectivec.Object{ID: rv}
}

// SpeechVoicesIncludingSiriAndSuperCompact is an exported wrapper for the private method _speechVoicesIncludingSiriAndSuperCompact.
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) SpeechVoicesIncludingSiriAndSuperCompact() objectivec.IObject {
	return _AVSpeechSynthesisVoiceClass._speechVoicesIncludingSiriAndSuperCompact()
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/_speechVoicesIncludingSiriAndSuperCompactWithCompletionHandler:
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) _speechVoicesIncludingSiriAndSuperCompactWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_speechVoicesIncludingSiriAndSuperCompactWithCompletionHandler:"), _block0)
}

// SpeechVoicesIncludingSiriAndSuperCompactWithCompletionHandler is an exported wrapper for the private method _speechVoicesIncludingSiriAndSuperCompactWithCompletionHandler.
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) SpeechVoicesIncludingSiriAndSuperCompactWithCompletionHandler(handler ErrorHandler) {
	_AVSpeechSynthesisVoiceClass._speechVoicesIncludingSiriAndSuperCompactWithCompletionHandler(handler)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/_speechVoicesIncludingSiriWithCompletionHandler:
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) _speechVoicesIncludingSiriWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_speechVoicesIncludingSiriWithCompletionHandler:"), _block0)
}

// SpeechVoicesIncludingSiriWithCompletionHandler is an exported wrapper for the private method _speechVoicesIncludingSiriWithCompletionHandler.
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) SpeechVoicesIncludingSiriWithCompletionHandler(handler ErrorHandler) {
	_AVSpeechSynthesisVoiceClass._speechVoicesIncludingSiriWithCompletionHandler(handler)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/_voiceFromInternalVoiceListWithIdentifier:
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) _voiceFromInternalVoiceListWithIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_voiceFromInternalVoiceListWithIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}

// VoiceFromInternalVoiceListWithIdentifier is an exported wrapper for the private method _voiceFromInternalVoiceListWithIdentifier.
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) VoiceFromInternalVoiceListWithIdentifier(identifier objectivec.IObject) objectivec.IObject {
	return _AVSpeechSynthesisVoiceClass._voiceFromInternalVoiceListWithIdentifier(identifier)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/speechVoicesIncludingSuperCompact
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) SpeechVoicesIncludingSuperCompact() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("speechVoicesIncludingSuperCompact"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/speechVoicesIncludingSuperCompactWithCompletionHandler:
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) SpeechVoicesIncludingSuperCompactWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("speechVoicesIncludingSuperCompactWithCompletionHandler:"), _block0)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/speechVoicesWithCompletionHandler:
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) SpeechVoicesWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("speechVoicesWithCompletionHandler:"), _block0)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/supportsSecureCoding
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/test_setInternalSpeechVoices:
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) Test_setInternalSpeechVoices(voices objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("test_setInternalSpeechVoices:"), voices)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/test_speechVoices
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) Test_speechVoices() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("test_speechVoices"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/voiceWithIdentifier:
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) VoiceWithIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("voiceWithIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/voiceWithLanguage:
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) VoiceWithLanguage(language objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("voiceWithLanguage:"), language)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/coreVoiceIsSiriVoice
func (s AVSpeechSynthesisVoice) CoreVoiceIsSiriVoice() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("coreVoiceIsSiriVoice"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/coreVoiceIsSystemVoice
func (s AVSpeechSynthesisVoice) CoreVoiceIsSystemVoice() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("coreVoiceIsSystemVoice"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/coreVoiceLocalizedName
func (s AVSpeechSynthesisVoice) CoreVoiceLocalizedName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("coreVoiceLocalizedName"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/coreVoiceWrapper
func (s AVSpeechSynthesisVoice) CoreVoiceWrapper() objectivec.Object {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("coreVoiceWrapper"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (s AVSpeechSynthesisVoice) SetCoreVoiceWrapper(value objectivec.Object) {
	objc.Send[struct{}](s.ID, objc.Sel("setCoreVoiceWrapper:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/gender
func (s AVSpeechSynthesisVoice) Gender() int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("gender"))
	return rv
}
func (s AVSpeechSynthesisVoice) SetGender(value int64) {
	objc.Send[struct{}](s.ID, objc.Sel("setGender:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/identifier
func (s AVSpeechSynthesisVoice) Identifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
func (s AVSpeechSynthesisVoice) SetIdentifier(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setIdentifier:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/isNoveltyVoice
func (s AVSpeechSynthesisVoice) IsNoveltyVoice() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isNoveltyVoice"))
	return rv
}
func (s AVSpeechSynthesisVoice) SetIsNoveltyVoice(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setIsNoveltyVoice:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/isSiriVoice
func (s AVSpeechSynthesisVoice) IsSiriVoice() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isSiriVoice"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/language
func (s AVSpeechSynthesisVoice) Language() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("language"))
	return foundation.NSStringFromID(rv).String()
}
func (s AVSpeechSynthesisVoice) SetLanguage(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setLanguage:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/name
func (s AVSpeechSynthesisVoice) Name() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (s AVSpeechSynthesisVoice) SetName(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/nameWithoutQuality
func (s AVSpeechSynthesisVoice) NameWithoutQuality() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("nameWithoutQuality"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/nonLocalizedName
func (s AVSpeechSynthesisVoice) NonLocalizedName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("nonLocalizedName"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/nonLocalizedNameWithoutQuality
func (s AVSpeechSynthesisVoice) NonLocalizedNameWithoutQuality() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("nonLocalizedNameWithoutQuality"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/quality
func (s AVSpeechSynthesisVoice) Quality() int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("quality"))
	return rv
}
func (s AVSpeechSynthesisVoice) SetQuality(value int64) {
	objc.Send[struct{}](s.ID, objc.Sel("setQuality:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/synthesisProviderVoice
func (s AVSpeechSynthesisVoice) SynthesisProviderVoice() IAVSpeechSynthesisProviderVoice {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("synthesisProviderVoice"))
	return AVSpeechSynthesisProviderVoiceFromID(objc.ID(rv))
}
func (s AVSpeechSynthesisVoice) SetSynthesisProviderVoice(value IAVSpeechSynthesisProviderVoice) {
	objc.Send[struct{}](s.ID, objc.Sel("setSynthesisProviderVoice:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/voiceTraits
func (s AVSpeechSynthesisVoice) VoiceTraits() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("voiceTraits"))
	return rv
}
func (s AVSpeechSynthesisVoice) SetVoiceTraits(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setVoiceTraits:"), value)
}

// _speechVoicesIncludingSiriAndSuperCompactSync is a synchronous wrapper around [AVSpeechSynthesisVoice._speechVoicesIncludingSiriAndSuperCompactWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (sc AVSpeechSynthesisVoiceClass) _speechVoicesIncludingSiriAndSuperCompactSync(ctx context.Context) error {
	done := make(chan error, 1)
	sc._speechVoicesIncludingSiriAndSuperCompactWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// _speechVoicesIncludingSiriSync is a synchronous wrapper around [AVSpeechSynthesisVoice._speechVoicesIncludingSiriWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (sc AVSpeechSynthesisVoiceClass) _speechVoicesIncludingSiriSync(ctx context.Context) error {
	done := make(chan error, 1)
	sc._speechVoicesIncludingSiriWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SpeechVoicesIncludingSuperCompactSync is a synchronous wrapper around [AVSpeechSynthesisVoice.SpeechVoicesIncludingSuperCompactWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (sc AVSpeechSynthesisVoiceClass) SpeechVoicesIncludingSuperCompactSync(ctx context.Context) error {
	done := make(chan error, 1)
	sc.SpeechVoicesIncludingSuperCompactWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SpeechVoices is a synchronous wrapper around [AVSpeechSynthesisVoice.SpeechVoicesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (sc AVSpeechSynthesisVoiceClass) SpeechVoices(ctx context.Context) error {
	done := make(chan error, 1)
	sc.SpeechVoicesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

