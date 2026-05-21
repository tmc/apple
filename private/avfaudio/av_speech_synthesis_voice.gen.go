// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
func (a AVSpeechSynthesisVoice) Init() AVSpeechSynthesisVoice {
	rv := objc.Send[AVSpeechSynthesisVoice](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVSpeechSynthesisVoice) Autorelease() AVSpeechSynthesisVoice {
	rv := objc.Send[AVSpeechSynthesisVoice](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSpeechSynthesisVoice creates a new AVSpeechSynthesisVoice instance.
func NewAVSpeechSynthesisVoice() AVSpeechSynthesisVoice {
	class := getAVSpeechSynthesisVoiceClass()
	rv := objc.Send[AVSpeechSynthesisVoice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSpeechSynthesisVoiceWithLanguage(language objectivec.IObject) AVSpeechSynthesisVoice {
	instance := getAVSpeechSynthesisVoiceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLanguage:"), language)
	return AVSpeechSynthesisVoiceFromID(rv)
}

func (a AVSpeechSynthesisVoice) AssetSize() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("assetSize"))
	return rv
}
func (a AVSpeechSynthesisVoice) BackupName() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("backupName"))
	return objectivec.Object{ID: rv}
}
func (a AVSpeechSynthesisVoice) CanBeDownloaded() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("canBeDownloaded"))
	return rv
}
func (a AVSpeechSynthesisVoice) IsDefault() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isDefault"))
	return rv
}
func (a AVSpeechSynthesisVoice) IsInstalled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isInstalled"))
	return rv
}
func (a AVSpeechSynthesisVoice) IsSystemVoice() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isSystemVoice"))
	return rv
}
func (a AVSpeechSynthesisVoice) SetAssetSize(size int64) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAssetSize:"), size)
}
func (a AVSpeechSynthesisVoice) SetBackupName(name objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setBackupName:"), name)
}
func (a AVSpeechSynthesisVoice) SetCanBeDownloaded(downloaded bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setCanBeDownloaded:"), downloaded)
}
func (a AVSpeechSynthesisVoice) SetIsDefault(default_ bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setIsDefault:"), default_)
}
func (a AVSpeechSynthesisVoice) SetIsInstalled(installed bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setIsInstalled:"), installed)
}
func (a AVSpeechSynthesisVoice) SiriDisplayName() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("siriDisplayName"))
	return objectivec.Object{ID: rv}
}
func (a AVSpeechSynthesisVoice) InitWithLanguage(language objectivec.IObject) AVSpeechSynthesisVoice {
	rv := objc.Send[AVSpeechSynthesisVoice](a.ID, objc.Sel("initWithLanguage:"), language)
	return rv
}

func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) _speechVoicesIncludingSiri() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_speechVoicesIncludingSiri"))
	return objectivec.Object{ID: rv}
}

// SpeechVoicesIncludingSiri is an exported wrapper for the private method _speechVoicesIncludingSiri.
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) SpeechVoicesIncludingSiri() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_speechVoicesIncludingSiri")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_speechVoicesIncludingSiri"}
		return nil, err
	}
	return _AVSpeechSynthesisVoiceClass._speechVoicesIncludingSiri(), nil
}

// CanSpeechVoicesIncludingSiri reports whether the receiver responds to the private selector _speechVoicesIncludingSiri.
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) CanSpeechVoicesIncludingSiri() bool {
	return objc.RespondsToSelector(objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_speechVoicesIncludingSiri"))
}
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) _speechVoicesIncludingSiriWithSiri(siri bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_speechVoicesIncludingSiri:"), siri)
	return objectivec.Object{ID: rv}
}

// SpeechVoicesIncludingSiriWithSiri is an exported wrapper for the private method _speechVoicesIncludingSiriWithSiri.
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) SpeechVoicesIncludingSiriWithSiri(siri bool) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_speechVoicesIncludingSiri:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_speechVoicesIncludingSiri:"}
		return nil, err
	}
	return _AVSpeechSynthesisVoiceClass._speechVoicesIncludingSiriWithSiri(siri), nil
}

// CanSpeechVoicesIncludingSiriWithSiri reports whether the receiver responds to the private selector _speechVoicesIncludingSiri:.
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) CanSpeechVoicesIncludingSiriWithSiri() bool {
	return objc.RespondsToSelector(objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_speechVoicesIncludingSiri:"))
}
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) _speechVoicesIncludingSiriAndSuperCompact() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_speechVoicesIncludingSiriAndSuperCompact"))
	return objectivec.Object{ID: rv}
}

// SpeechVoicesIncludingSiriAndSuperCompact is an exported wrapper for the private method _speechVoicesIncludingSiriAndSuperCompact.
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) SpeechVoicesIncludingSiriAndSuperCompact() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_speechVoicesIncludingSiriAndSuperCompact")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_speechVoicesIncludingSiriAndSuperCompact"}
		return nil, err
	}
	return _AVSpeechSynthesisVoiceClass._speechVoicesIncludingSiriAndSuperCompact(), nil
}

// CanSpeechVoicesIncludingSiriAndSuperCompact reports whether the receiver responds to the private selector _speechVoicesIncludingSiriAndSuperCompact.
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) CanSpeechVoicesIncludingSiriAndSuperCompact() bool {
	return objc.RespondsToSelector(objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_speechVoicesIncludingSiriAndSuperCompact"))
}
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) _speechVoicesIncludingSiriAndSuperCompactWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_speechVoicesIncludingSiriAndSuperCompactWithCompletionHandler:"), _block0)
}

// SpeechVoicesIncludingSiriAndSuperCompactWithCompletionHandler is an exported wrapper for the private method _speechVoicesIncludingSiriAndSuperCompactWithCompletionHandler.
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) SpeechVoicesIncludingSiriAndSuperCompactWithCompletionHandler(handler ErrorHandler) error {
	if !objc.RespondsToSelector(objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_speechVoicesIncludingSiriAndSuperCompactWithCompletionHandler:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_speechVoicesIncludingSiriAndSuperCompactWithCompletionHandler:"}
		return err
	}
	_AVSpeechSynthesisVoiceClass._speechVoicesIncludingSiriAndSuperCompactWithCompletionHandler(handler)
	return nil
}

// CanSpeechVoicesIncludingSiriAndSuperCompactWithCompletionHandler reports whether the receiver responds to the private selector _speechVoicesIncludingSiriAndSuperCompactWithCompletionHandler:.
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) CanSpeechVoicesIncludingSiriAndSuperCompactWithCompletionHandler() bool {
	return objc.RespondsToSelector(objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_speechVoicesIncludingSiriAndSuperCompactWithCompletionHandler:"))
}
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) _speechVoicesIncludingSiriWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_speechVoicesIncludingSiriWithCompletionHandler:"), _block0)
}

// SpeechVoicesIncludingSiriWithCompletionHandler is an exported wrapper for the private method _speechVoicesIncludingSiriWithCompletionHandler.
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) SpeechVoicesIncludingSiriWithCompletionHandler(handler ErrorHandler) error {
	if !objc.RespondsToSelector(objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_speechVoicesIncludingSiriWithCompletionHandler:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_speechVoicesIncludingSiriWithCompletionHandler:"}
		return err
	}
	_AVSpeechSynthesisVoiceClass._speechVoicesIncludingSiriWithCompletionHandler(handler)
	return nil
}

// CanSpeechVoicesIncludingSiriWithCompletionHandler reports whether the receiver responds to the private selector _speechVoicesIncludingSiriWithCompletionHandler:.
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) CanSpeechVoicesIncludingSiriWithCompletionHandler() bool {
	return objc.RespondsToSelector(objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_speechVoicesIncludingSiriWithCompletionHandler:"))
}
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) _voiceFromInternalVoiceListWithIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_voiceFromInternalVoiceListWithIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}

// VoiceFromInternalVoiceListWithIdentifier is an exported wrapper for the private method _voiceFromInternalVoiceListWithIdentifier.
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) VoiceFromInternalVoiceListWithIdentifier(identifier objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_voiceFromInternalVoiceListWithIdentifier:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_voiceFromInternalVoiceListWithIdentifier:"}
		return nil, err
	}
	return _AVSpeechSynthesisVoiceClass._voiceFromInternalVoiceListWithIdentifier(identifier), nil
}

// CanVoiceFromInternalVoiceListWithIdentifier reports whether the receiver responds to the private selector _voiceFromInternalVoiceListWithIdentifier:.
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) CanVoiceFromInternalVoiceListWithIdentifier() bool {
	return objc.RespondsToSelector(objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("_voiceFromInternalVoiceListWithIdentifier:"))
}
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) SpeechVoicesIncludingSuperCompact() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("speechVoicesIncludingSuperCompact"))
	return objectivec.Object{ID: rv}
}
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) SpeechVoicesIncludingSuperCompactWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("speechVoicesIncludingSuperCompactWithCompletionHandler:"), _block0)
}
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) SpeechVoicesWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("speechVoicesWithCompletionHandler:"), _block0)
}
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) Test_setInternalSpeechVoices(voices objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("test_setInternalSpeechVoices:"), voices)
}
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) Test_speechVoices() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("test_speechVoices"))
	return objectivec.Object{ID: rv}
}
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) VoiceWithIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("voiceWithIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) VoiceWithLanguage(language objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("voiceWithLanguage:"), language)
	return objectivec.Object{ID: rv}
}

func (a AVSpeechSynthesisVoice) CoreVoiceIsSiriVoice() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("coreVoiceIsSiriVoice"))
	return rv
}
func (a AVSpeechSynthesisVoice) CoreVoiceIsSystemVoice() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("coreVoiceIsSystemVoice"))
	return rv
}
func (a AVSpeechSynthesisVoice) CoreVoiceLocalizedName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("coreVoiceLocalizedName"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVSpeechSynthesisVoice) CoreVoiceWrapper() objectivec.Object {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("coreVoiceWrapper"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (a AVSpeechSynthesisVoice) SetCoreVoiceWrapper(value objectivec.Object) {
	objc.Send[struct{}](a.ID, objc.Sel("setCoreVoiceWrapper:"), value)
}
func (a AVSpeechSynthesisVoice) Gender() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("gender"))
	return rv
}
func (a AVSpeechSynthesisVoice) SetGender(value int64) {
	objc.Send[struct{}](a.ID, objc.Sel("setGender:"), value)
}
func (a AVSpeechSynthesisVoice) Identifier() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVSpeechSynthesisVoice) SetIdentifier(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setIdentifier:"), objc.String(value))
}
func (a AVSpeechSynthesisVoice) IsNoveltyVoice() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isNoveltyVoice"))
	return rv
}
func (a AVSpeechSynthesisVoice) SetIsNoveltyVoice(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setIsNoveltyVoice:"), value)
}
func (a AVSpeechSynthesisVoice) IsSiriVoice() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isSiriVoice"))
	return rv
}
func (a AVSpeechSynthesisVoice) Language() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("language"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVSpeechSynthesisVoice) SetLanguage(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setLanguage:"), objc.String(value))
}
func (a AVSpeechSynthesisVoice) Name() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVSpeechSynthesisVoice) SetName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setName:"), objc.String(value))
}
func (a AVSpeechSynthesisVoice) NameWithoutQuality() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("nameWithoutQuality"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVSpeechSynthesisVoice) NonLocalizedName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("nonLocalizedName"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVSpeechSynthesisVoice) NonLocalizedNameWithoutQuality() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("nonLocalizedNameWithoutQuality"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVSpeechSynthesisVoice) Quality() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("quality"))
	return rv
}
func (a AVSpeechSynthesisVoice) SetQuality(value int64) {
	objc.Send[struct{}](a.ID, objc.Sel("setQuality:"), value)
}
func (a AVSpeechSynthesisVoice) SynthesisProviderVoice() IAVSpeechSynthesisProviderVoice {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("synthesisProviderVoice"))
	return AVSpeechSynthesisProviderVoiceFromID(objc.ID(rv))
}
func (a AVSpeechSynthesisVoice) SetSynthesisProviderVoice(value IAVSpeechSynthesisProviderVoice) {
	objc.Send[struct{}](a.ID, objc.Sel("setSynthesisProviderVoice:"), value)
}
func (a AVSpeechSynthesisVoice) VoiceTraits() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("voiceTraits"))
	return rv
}
func (a AVSpeechSynthesisVoice) SetVoiceTraits(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setVoiceTraits:"), value)
}

// _speechVoicesIncludingSiriAndSuperCompactSync is a synchronous wrapper around [AVSpeechSynthesisVoice._speechVoicesIncludingSiriAndSuperCompactWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (ac AVSpeechSynthesisVoiceClass) _speechVoicesIncludingSiriAndSuperCompactSync(ctx context.Context) error {
	done := make(chan error, 1)
	ac._speechVoicesIncludingSiriAndSuperCompactWithCompletionHandler(func(err error) {
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
func (ac AVSpeechSynthesisVoiceClass) _speechVoicesIncludingSiriSync(ctx context.Context) error {
	done := make(chan error, 1)
	ac._speechVoicesIncludingSiriWithCompletionHandler(func(err error) {
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
func (ac AVSpeechSynthesisVoiceClass) SpeechVoicesIncludingSuperCompactSync(ctx context.Context) error {
	done := make(chan error, 1)
	ac.SpeechVoicesIncludingSuperCompactWithCompletionHandler(func(err error) {
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
func (ac AVSpeechSynthesisVoiceClass) SpeechVoices(ctx context.Context) error {
	done := make(chan error, 1)
	ac.SpeechVoicesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
