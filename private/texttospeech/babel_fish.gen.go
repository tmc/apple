// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [BabelFish] class.
var (
	_BabelFishClass     BabelFishClass
	_BabelFishClassOnce sync.Once
)

func getBabelFishClass() BabelFishClass {
	_BabelFishClassOnce.Do(func() {
		_BabelFishClass = BabelFishClass{class: objc.GetClass("BabelFish")}
	})
	return _BabelFishClass
}

// GetBabelFishClass returns the class object for BabelFish.
func GetBabelFishClass() BabelFishClass {
	return getBabelFishClass()
}

type BabelFishClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (bc BabelFishClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BabelFishClass) Alloc() BabelFish {
	rv := objc.SendIfResponds[BabelFish](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [BabelFish._characterRangesForLocaleCode]
//   - [BabelFish._disambiguationKeyForVoice]
//   - [BabelFish._disambiguationMatrixForVoices]
//   - [BabelFish._ensureCachedVoicesLoaded]
//   - [BabelFish._monitorAudioDevices]
//   - [BabelFish._voiceForIdentifier]
//   - [BabelFish.AllVoices]
//   - [BabelFish.CachedVoices]
//   - [BabelFish.SetCachedVoices]
//   - [BabelFish.CharacterRangesForVoice]
//   - [BabelFish.DevicesChanged]
//   - [BabelFish.DisambiguationMap]
//   - [BabelFish.SetDisambiguationMap]
//   - [BabelFish.ExemplarCharacterRangesForLocaleString]
//   - [BabelFish.GetChannel]
//   - [BabelFish.HasSiriEntitlement]
//   - [BabelFish.KillChannel]
//   - [BabelFish.LanguageToCharacterRangesCache]
//   - [BabelFish.SetLanguageToCharacterRangesCache]
//   - [BabelFish.LocaleNameCache]
//   - [BabelFish.SetLocaleNameCache]
//   - [BabelFish.NameForLocale]
//   - [BabelFish.NameForVoice]
//   - [BabelFish.NewChannel]
//   - [BabelFish.NextChannelIdentifier]
//   - [BabelFish.PerformForChannelBlock]
//   - [BabelFish.ResourceCacheDidReceiveUpdate]
//   - [BabelFish.SpeechBusy]
//   - [BabelFish.SpeechChannelMap]
//   - [BabelFish.SetSpeechChannelMap]
//   - [BabelFish.VoiceAccessQueue]
//   - [BabelFish.SetVoiceAccessQueue]
//   - [BabelFish.VoiceForIdentifier]
//   - [BabelFish.VoiceForLocale]
//   - [BabelFish.VoiceForLocaleAllowedIdentifiers]
//   - [BabelFish.VoiceForSpecId]
//   - [BabelFish.VoicesBySpec]
//   - [BabelFish.SetVoicesBySpec]
//   - [BabelFish.DebugDescription]
//   - [BabelFish.Description]
//   - [BabelFish.Hash]
//   - [BabelFish.Superclass]
type BabelFish struct {
	objectivec.Object
}

// BabelFishFromID constructs a [BabelFish] from an objc.ID.
func BabelFishFromID(id objc.ID) BabelFish {
	return BabelFish{objectivec.Object{ID: id}}
}

// Ensure BabelFish implements IBabelFish.
var _ IBabelFish = BabelFish{}

// An interface definition for the [BabelFish] class.
//
// # Methods
//
//   - [IBabelFish._characterRangesForLocaleCode]
//   - [IBabelFish._disambiguationKeyForVoice]
//   - [IBabelFish._disambiguationMatrixForVoices]
//   - [IBabelFish._ensureCachedVoicesLoaded]
//   - [IBabelFish._monitorAudioDevices]
//   - [IBabelFish._voiceForIdentifier]
//   - [IBabelFish.AllVoices]
//   - [IBabelFish.CachedVoices]
//   - [IBabelFish.SetCachedVoices]
//   - [IBabelFish.CharacterRangesForVoice]
//   - [IBabelFish.DevicesChanged]
//   - [IBabelFish.DisambiguationMap]
//   - [IBabelFish.SetDisambiguationMap]
//   - [IBabelFish.ExemplarCharacterRangesForLocaleString]
//   - [IBabelFish.GetChannel]
//   - [IBabelFish.HasSiriEntitlement]
//   - [IBabelFish.KillChannel]
//   - [IBabelFish.LanguageToCharacterRangesCache]
//   - [IBabelFish.SetLanguageToCharacterRangesCache]
//   - [IBabelFish.LocaleNameCache]
//   - [IBabelFish.SetLocaleNameCache]
//   - [IBabelFish.NameForLocale]
//   - [IBabelFish.NameForVoice]
//   - [IBabelFish.NewChannel]
//   - [IBabelFish.NextChannelIdentifier]
//   - [IBabelFish.PerformForChannelBlock]
//   - [IBabelFish.ResourceCacheDidReceiveUpdate]
//   - [IBabelFish.SpeechBusy]
//   - [IBabelFish.SpeechChannelMap]
//   - [IBabelFish.SetSpeechChannelMap]
//   - [IBabelFish.VoiceAccessQueue]
//   - [IBabelFish.SetVoiceAccessQueue]
//   - [IBabelFish.VoiceForIdentifier]
//   - [IBabelFish.VoiceForLocale]
//   - [IBabelFish.VoiceForLocaleAllowedIdentifiers]
//   - [IBabelFish.VoiceForSpecId]
//   - [IBabelFish.VoicesBySpec]
//   - [IBabelFish.SetVoicesBySpec]
//   - [IBabelFish.DebugDescription]
//   - [IBabelFish.Description]
//   - [IBabelFish.Hash]
//   - [IBabelFish.Superclass]
type IBabelFish interface {
	objectivec.IObject

	// Topic: Methods

	_characterRangesForLocaleCode(code objectivec.IObject) objectivec.IObject
	_disambiguationKeyForVoice(voice objectivec.IObject) objectivec.IObject
	_disambiguationMatrixForVoices(voices objectivec.IObject) objectivec.IObject
	_ensureCachedVoicesLoaded()
	_monitorAudioDevices()
	_voiceForIdentifier(identifier objectivec.IObject) objectivec.IObject
	AllVoices() objectivec.IObject
	CachedVoices() foundation.INSArray
	SetCachedVoices(value foundation.INSArray)
	CharacterRangesForVoice(voice objectivec.IObject) objectivec.IObject
	DevicesChanged()
	DisambiguationMap() foundation.INSDictionary
	SetDisambiguationMap(value foundation.INSDictionary)
	ExemplarCharacterRangesForLocaleString(string_ objectivec.IObject) objectivec.IObject
	GetChannel(channel int64) objectivec.IObject
	HasSiriEntitlement() bool
	KillChannel(channel int64) int16
	LanguageToCharacterRangesCache() foundation.INSDictionary
	SetLanguageToCharacterRangesCache(value foundation.INSDictionary)
	LocaleNameCache() foundation.INSDictionary
	SetLocaleNameCache(value foundation.INSDictionary)
	NameForLocale(locale objectivec.IObject) objectivec.IObject
	NameForVoice(voice objectivec.IObject) objectivec.IObject
	NewChannel(channel uint32) objectivec.IObject
	NextChannelIdentifier() uint64
	PerformForChannelBlock(channel int64, block VoidHandler) int16
	ResourceCacheDidReceiveUpdate()
	SpeechBusy() bool
	SpeechChannelMap() foundation.INSDictionary
	SetSpeechChannelMap(value foundation.INSDictionary)
	VoiceAccessQueue() objectivec.Object
	SetVoiceAccessQueue(value objectivec.Object)
	VoiceForIdentifier(identifier objectivec.IObject) objectivec.IObject
	VoiceForLocale(locale objectivec.IObject) objectivec.IObject
	VoiceForLocaleAllowedIdentifiers(locale objectivec.IObject, identifiers objectivec.IObject) objectivec.IObject
	VoiceForSpecId(id uint32) objectivec.IObject
	VoicesBySpec() foundation.INSDictionary
	SetVoicesBySpec(value foundation.INSDictionary)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (b BabelFish) Init() BabelFish {
	rv := objc.SendIfResponds[BabelFish](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BabelFish) Autorelease() BabelFish {
	rv := objc.SendIfResponds[BabelFish](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBabelFish creates a new BabelFish instance.
func NewBabelFish() BabelFish {
	class := getBabelFishClass()
	rv := objc.SendIfResponds[BabelFish](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (b BabelFish) _characterRangesForLocaleCode(code objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("_characterRangesForLocaleCode:"), code)
	return objectivec.Object{ID: rv}
}

// CharacterRangesForLocaleCode is an exported wrapper for the private method _characterRangesForLocaleCode.
func (b BabelFish) CharacterRangesForLocaleCode(code objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(b.ID, objc.Sel("_characterRangesForLocaleCode:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_characterRangesForLocaleCode:"}
		return nil, err
	}
	return b._characterRangesForLocaleCode(code), nil
}

// CanCharacterRangesForLocaleCode reports whether the receiver responds to the private selector _characterRangesForLocaleCode:.
func (b BabelFish) CanCharacterRangesForLocaleCode() bool {
	return objc.RespondsToSelector(b.ID, objc.Sel("_characterRangesForLocaleCode:"))
}
func (b BabelFish) _disambiguationKeyForVoice(voice objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("_disambiguationKeyForVoice:"), voice)
	return objectivec.Object{ID: rv}
}

// DisambiguationKeyForVoice is an exported wrapper for the private method _disambiguationKeyForVoice.
func (b BabelFish) DisambiguationKeyForVoice(voice objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(b.ID, objc.Sel("_disambiguationKeyForVoice:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_disambiguationKeyForVoice:"}
		return nil, err
	}
	return b._disambiguationKeyForVoice(voice), nil
}

// CanDisambiguationKeyForVoice reports whether the receiver responds to the private selector _disambiguationKeyForVoice:.
func (b BabelFish) CanDisambiguationKeyForVoice() bool {
	return objc.RespondsToSelector(b.ID, objc.Sel("_disambiguationKeyForVoice:"))
}
func (b BabelFish) _disambiguationMatrixForVoices(voices objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("_disambiguationMatrixForVoices:"), voices)
	return objectivec.Object{ID: rv}
}

// DisambiguationMatrixForVoices is an exported wrapper for the private method _disambiguationMatrixForVoices.
func (b BabelFish) DisambiguationMatrixForVoices(voices objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(b.ID, objc.Sel("_disambiguationMatrixForVoices:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_disambiguationMatrixForVoices:"}
		return nil, err
	}
	return b._disambiguationMatrixForVoices(voices), nil
}

// CanDisambiguationMatrixForVoices reports whether the receiver responds to the private selector _disambiguationMatrixForVoices:.
func (b BabelFish) CanDisambiguationMatrixForVoices() bool {
	return objc.RespondsToSelector(b.ID, objc.Sel("_disambiguationMatrixForVoices:"))
}
func (b BabelFish) _ensureCachedVoicesLoaded() {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("_ensureCachedVoicesLoaded"))
}

// EnsureCachedVoicesLoaded is an exported wrapper for the private method _ensureCachedVoicesLoaded.
func (b BabelFish) EnsureCachedVoicesLoaded() error {
	if !objc.RespondsToSelector(b.ID, objc.Sel("_ensureCachedVoicesLoaded")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_ensureCachedVoicesLoaded"}
		return err
	}
	b._ensureCachedVoicesLoaded()
	return nil
}

// CanEnsureCachedVoicesLoaded reports whether the receiver responds to the private selector _ensureCachedVoicesLoaded.
func (b BabelFish) CanEnsureCachedVoicesLoaded() bool {
	return objc.RespondsToSelector(b.ID, objc.Sel("_ensureCachedVoicesLoaded"))
}
func (b BabelFish) _monitorAudioDevices() {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("_monitorAudioDevices"))
}

// MonitorAudioDevices is an exported wrapper for the private method _monitorAudioDevices.
func (b BabelFish) MonitorAudioDevices() error {
	if !objc.RespondsToSelector(b.ID, objc.Sel("_monitorAudioDevices")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_monitorAudioDevices"}
		return err
	}
	b._monitorAudioDevices()
	return nil
}

// CanMonitorAudioDevices reports whether the receiver responds to the private selector _monitorAudioDevices.
func (b BabelFish) CanMonitorAudioDevices() bool {
	return objc.RespondsToSelector(b.ID, objc.Sel("_monitorAudioDevices"))
}
func (b BabelFish) _voiceForIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("_voiceForIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}
func (b BabelFish) AllVoices() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("allVoices"))
	return objectivec.Object{ID: rv}
}
func (b BabelFish) CharacterRangesForVoice(voice objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("characterRangesForVoice:"), voice)
	return objectivec.Object{ID: rv}
}
func (b BabelFish) DevicesChanged() {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("devicesChanged"))
}
func (b BabelFish) ExemplarCharacterRangesForLocaleString(string_ objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("exemplarCharacterRangesForLocaleString:"), string_)
	return objectivec.Object{ID: rv}
}
func (b BabelFish) GetChannel(channel int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("getChannel:"), channel)
	return objectivec.Object{ID: rv}
}
func (b BabelFish) HasSiriEntitlement() bool {
	rv := objc.SendIfResponds[bool](b.ID, objc.Sel("hasSiriEntitlement"))
	return rv
}
func (b BabelFish) KillChannel(channel int64) int16 {
	rv := objc.SendIfResponds[int16](b.ID, objc.Sel("killChannel:"), channel)
	return rv
}
func (b BabelFish) NameForLocale(locale objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("nameForLocale:"), locale)
	return objectivec.Object{ID: rv}
}
func (b BabelFish) NameForVoice(voice objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("nameForVoice:"), voice)
	return objectivec.Object{ID: rv}
}
func (b BabelFish) NewChannel(channel uint32) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("newChannel:"), channel)
	return objectivec.Object{ID: rv}
}
func (b BabelFish) NextChannelIdentifier() uint64 {
	rv := objc.SendIfResponds[uint64](b.ID, objc.Sel("nextChannelIdentifier"))
	return rv
}
func (b BabelFish) PerformForChannelBlock(channel int64, block VoidHandler) int16 {
	_block1, _ := NewVoidBlock(block)
	rv := objc.SendIfResponds[int16](b.ID, objc.Sel("performForChannel:block:"), channel, _block1)
	return rv
}
func (b BabelFish) ResourceCacheDidReceiveUpdate() {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("resourceCacheDidReceiveUpdate"))
}
func (b BabelFish) SpeechBusy() bool {
	rv := objc.SendIfResponds[bool](b.ID, objc.Sel("speechBusy"))
	return rv
}
func (b BabelFish) VoiceForIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("voiceForIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}
func (b BabelFish) VoiceForLocale(locale objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("voiceForLocale:"), locale)
	return objectivec.Object{ID: rv}
}
func (b BabelFish) VoiceForLocaleAllowedIdentifiers(locale objectivec.IObject, identifiers objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("voiceForLocale:allowedIdentifiers:"), locale, identifiers)
	return objectivec.Object{ID: rv}
}
func (b BabelFish) VoiceForSpecId(id uint32) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("voiceForSpecId:"), id)
	return objectivec.Object{ID: rv}
}

func (_BabelFishClass BabelFishClass) AxLanguageManager() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_BabelFishClass.class), objc.Sel("axLanguageManager"))
	return objectivec.Object{ID: rv}
}
func (_BabelFishClass BabelFishClass) AxSettings() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_BabelFishClass.class), objc.Sel("axSettings"))
	return objectivec.Object{ID: rv}
}
func (_BabelFishClass BabelFishClass) LocaleUtilities() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_BabelFishClass.class), objc.Sel("localeUtilities"))
	return objectivec.Object{ID: rv}
}
func (_BabelFishClass BabelFishClass) ResourceManager() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_BabelFishClass.class), objc.Sel("resourceManager"))
	return objectivec.Object{ID: rv}
}
func (_BabelFishClass BabelFishClass) SharedInstance() BabelFish {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_BabelFishClass.class), objc.Sel("sharedInstance"))
	return BabelFishFromID(rv)
}
func (_BabelFishClass BabelFishClass) SpecIdForTTSVoice(tTSVoice objectivec.IObject) uint32 {
	rv := objc.SendIfResponds[uint32](objc.ID(_BabelFishClass.class), objc.Sel("specIdForTTSVoice:"), tTSVoice)
	return rv
}
func (_BabelFishClass BabelFishClass) StringKeyForVoiceSpecId(id uint32) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_BabelFishClass.class), objc.Sel("stringKeyForVoiceSpecId:"), id)
	return objectivec.Object{ID: rv}
}

func (b BabelFish) CachedVoices() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("cachedVoices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (b BabelFish) SetCachedVoices(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setCachedVoices:"), value)
}
func (b BabelFish) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (b BabelFish) Description() string {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (b BabelFish) DisambiguationMap() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("disambiguationMap"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (b BabelFish) SetDisambiguationMap(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setDisambiguationMap:"), value)
}
func (b BabelFish) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](b.ID, objc.Sel("hash"))
	return rv
}
func (b BabelFish) LanguageToCharacterRangesCache() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("languageToCharacterRangesCache"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (b BabelFish) SetLanguageToCharacterRangesCache(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setLanguageToCharacterRangesCache:"), value)
}
func (b BabelFish) LocaleNameCache() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("localeNameCache"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (b BabelFish) SetLocaleNameCache(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setLocaleNameCache:"), value)
}
func (b BabelFish) SpeechChannelMap() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("speechChannelMap"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (b BabelFish) SetSpeechChannelMap(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setSpeechChannelMap:"), value)
}
func (b BabelFish) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](b.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (b BabelFish) VoiceAccessQueue() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("voiceAccessQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (b BabelFish) SetVoiceAccessQueue(value objectivec.Object) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setVoiceAccessQueue:"), value)
}
func (b BabelFish) VoicesBySpec() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("voicesBySpec"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (b BabelFish) SetVoicesBySpec(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setVoicesBySpec:"), value)
}

// PerformForChannelBlockSync is a synchronous wrapper around [BabelFish.PerformForChannelBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (b BabelFish) PerformForChannelBlockSync(ctx context.Context, channel int64) error {
	done := make(chan struct{}, 1)
	b.PerformForChannelBlock(channel, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
