// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/avfaudio"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSAXResource] class.
var (
	_TTSAXResourceClass     TTSAXResourceClass
	_TTSAXResourceClassOnce sync.Once
)

func getTTSAXResourceClass() TTSAXResourceClass {
	_TTSAXResourceClassOnce.Do(func() {
		_TTSAXResourceClass = TTSAXResourceClass{class: objc.GetClass("TTSAXResource")}
	})
	return _TTSAXResourceClass
}

// GetTTSAXResourceClass returns the class object for TTSAXResource.
func GetTTSAXResourceClass() TTSAXResourceClass {
	return getTTSAXResourceClass()
}

type TTSAXResourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSAXResourceClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSAXResourceClass) Alloc() TTSAXResource {
	rv := objc.SendIfResponds[TTSAXResource](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TTSAXResource.AssetSize]
//   - [TTSAXResource.CanBeDownloaded]
//   - [TTSAXResource.SetCanBeDownloaded]
//   - [TTSAXResource.ContentPath]
//   - [TTSAXResource.SetContentPath]
//   - [TTSAXResource.Footprint]
//   - [TTSAXResource.SetFootprint]
//   - [TTSAXResource.Gender]
//   - [TTSAXResource.SetGender]
//   - [TTSAXResource.Identifier]
//   - [TTSAXResource.SetIdentifier]
//   - [TTSAXResource.IsDefault]
//   - [TTSAXResource.SetIsDefault]
//   - [TTSAXResource.IsInstalled]
//   - [TTSAXResource.IsNoveltyVoice]
//   - [TTSAXResource.SetIsNoveltyVoice]
//   - [TTSAXResource.IsPersonalVoice]
//   - [TTSAXResource.SetIsPersonalVoice]
//   - [TTSAXResource.IsSystemVoice]
//   - [TTSAXResource.SetIsSystemVoice]
//   - [TTSAXResource.Language]
//   - [TTSAXResource.SetLanguage]
//   - [TTSAXResource.Languages]
//   - [TTSAXResource.SetLanguages]
//   - [TTSAXResource.LocalizedName]
//   - [TTSAXResource.SetLocalizedName]
//   - [TTSAXResource.LocalizedNameForLanguageWithLanguage]
//   - [TTSAXResource.LocalizedNameWithFootprint]
//   - [TTSAXResource.SetLocalizedNameWithFootprint]
//   - [TTSAXResource.Name]
//   - [TTSAXResource.SetName]
//   - [TTSAXResource.NameAndFootprintForLanguage]
//   - [TTSAXResource.PrimaryLanguage]
//   - [TTSAXResource.SpeechVoice]
//   - [TTSAXResource.Subtype]
//   - [TTSAXResource.SetSubtype]
//   - [TTSAXResource.SynthesisProviderVoice]
//   - [TTSAXResource.SetSynthesisProviderVoice]
//   - [TTSAXResource.Type]
//   - [TTSAXResource.SetType]
//   - [TTSAXResource.VoiceId]
//   - [TTSAXResource.SetVoiceId]
//   - [TTSAXResource.VoiceType]
//   - [TTSAXResource.SetVoiceType]
type TTSAXResource struct {
	objectivec.Object
}

// TTSAXResourceFromID constructs a [TTSAXResource] from an objc.ID.
func TTSAXResourceFromID(id objc.ID) TTSAXResource {
	return TTSAXResource{objectivec.Object{ID: id}}
}

// Ensure TTSAXResource implements ITTSAXResource.
var _ ITTSAXResource = TTSAXResource{}

// An interface definition for the [TTSAXResource] class.
//
// # Methods
//
//   - [ITTSAXResource.AssetSize]
//   - [ITTSAXResource.CanBeDownloaded]
//   - [ITTSAXResource.SetCanBeDownloaded]
//   - [ITTSAXResource.ContentPath]
//   - [ITTSAXResource.SetContentPath]
//   - [ITTSAXResource.Footprint]
//   - [ITTSAXResource.SetFootprint]
//   - [ITTSAXResource.Gender]
//   - [ITTSAXResource.SetGender]
//   - [ITTSAXResource.Identifier]
//   - [ITTSAXResource.SetIdentifier]
//   - [ITTSAXResource.IsDefault]
//   - [ITTSAXResource.SetIsDefault]
//   - [ITTSAXResource.IsInstalled]
//   - [ITTSAXResource.IsNoveltyVoice]
//   - [ITTSAXResource.SetIsNoveltyVoice]
//   - [ITTSAXResource.IsPersonalVoice]
//   - [ITTSAXResource.SetIsPersonalVoice]
//   - [ITTSAXResource.IsSystemVoice]
//   - [ITTSAXResource.SetIsSystemVoice]
//   - [ITTSAXResource.Language]
//   - [ITTSAXResource.SetLanguage]
//   - [ITTSAXResource.Languages]
//   - [ITTSAXResource.SetLanguages]
//   - [ITTSAXResource.LocalizedName]
//   - [ITTSAXResource.SetLocalizedName]
//   - [ITTSAXResource.LocalizedNameForLanguageWithLanguage]
//   - [ITTSAXResource.LocalizedNameWithFootprint]
//   - [ITTSAXResource.SetLocalizedNameWithFootprint]
//   - [ITTSAXResource.Name]
//   - [ITTSAXResource.SetName]
//   - [ITTSAXResource.NameAndFootprintForLanguage]
//   - [ITTSAXResource.PrimaryLanguage]
//   - [ITTSAXResource.SpeechVoice]
//   - [ITTSAXResource.Subtype]
//   - [ITTSAXResource.SetSubtype]
//   - [ITTSAXResource.SynthesisProviderVoice]
//   - [ITTSAXResource.SetSynthesisProviderVoice]
//   - [ITTSAXResource.Type]
//   - [ITTSAXResource.SetType]
//   - [ITTSAXResource.VoiceId]
//   - [ITTSAXResource.SetVoiceId]
//   - [ITTSAXResource.VoiceType]
//   - [ITTSAXResource.SetVoiceType]
type ITTSAXResource interface {
	objectivec.IObject

	// Topic: Methods

	AssetSize() int64
	CanBeDownloaded() bool
	SetCanBeDownloaded(value bool)
	ContentPath() string
	SetContentPath(value string)
	Footprint() int64
	SetFootprint(value int64)
	Gender() int64
	SetGender(value int64)
	Identifier() string
	SetIdentifier(value string)
	IsDefault() bool
	SetIsDefault(value bool)
	IsInstalled() bool
	IsNoveltyVoice() bool
	SetIsNoveltyVoice(value bool)
	IsPersonalVoice() bool
	SetIsPersonalVoice(value bool)
	IsSystemVoice() bool
	SetIsSystemVoice(value bool)
	Language() string
	SetLanguage(value string)
	Languages() foundation.INSArray
	SetLanguages(value foundation.INSArray)
	LocalizedName() string
	SetLocalizedName(value string)
	LocalizedNameForLanguageWithLanguage(language objectivec.IObject) objectivec.IObject
	LocalizedNameWithFootprint() string
	SetLocalizedNameWithFootprint(value string)
	Name() string
	SetName(value string)
	NameAndFootprintForLanguage(language objectivec.IObject) objectivec.IObject
	PrimaryLanguage() objectivec.IObject
	SpeechVoice() objectivec.IObject
	Subtype() uint64
	SetSubtype(value uint64)
	SynthesisProviderVoice() avfaudio.AVSpeechSynthesisProviderVoice
	SetSynthesisProviderVoice(value avfaudio.AVSpeechSynthesisProviderVoice)
	Type() uint64
	SetType(value uint64)
	VoiceId() string
	SetVoiceId(value string)
	VoiceType() int64
	SetVoiceType(value int64)
}

// Init initializes the instance.
func (t TTSAXResource) Init() TTSAXResource {
	rv := objc.SendIfResponds[TTSAXResource](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSAXResource) Autorelease() TTSAXResource {
	rv := objc.SendIfResponds[TTSAXResource](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSAXResource creates a new TTSAXResource instance.
func NewTTSAXResource() TTSAXResource {
	class := getTTSAXResourceClass()
	rv := objc.SendIfResponds[TTSAXResource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (t TTSAXResource) AssetSize() int64 {
	rv := objc.SendIfResponds[int64](t.ID, objc.Sel("assetSize"))
	return rv
}
func (t TTSAXResource) IsInstalled() bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("isInstalled"))
	return rv
}
func (t TTSAXResource) LocalizedNameForLanguageWithLanguage(language objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("localizedNameForLanguage:"), language)
	return objectivec.Object{ID: rv}
}
func (t TTSAXResource) NameAndFootprintForLanguage(language objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("nameAndFootprintForLanguage:"), language)
	return objectivec.Object{ID: rv}
}
func (t TTSAXResource) PrimaryLanguage() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("primaryLanguage"))
	return objectivec.Object{ID: rv}
}
func (t TTSAXResource) SpeechVoice() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("speechVoice"))
	return objectivec.Object{ID: rv}
}

func (_TTSAXResourceClass TTSAXResourceClass) LocalizedNameWithForLanguage(name objectivec.IObject, language objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_TTSAXResourceClass.class), objc.Sel("localizedName:forLanguage:"), name, language)
	return objectivec.Object{ID: rv}
}

func (t TTSAXResource) CanBeDownloaded() bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("canBeDownloaded"))
	return rv
}
func (t TTSAXResource) SetCanBeDownloaded(value bool) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setCanBeDownloaded:"), value)
}
func (t TTSAXResource) ContentPath() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("contentPath"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSAXResource) SetContentPath(value string) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setContentPath:"), objc.String(value))
}
func (t TTSAXResource) Footprint() int64 {
	rv := objc.SendIfResponds[int64](t.ID, objc.Sel("footprint"))
	return rv
}
func (t TTSAXResource) SetFootprint(value int64) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setFootprint:"), value)
}
func (t TTSAXResource) Gender() int64 {
	rv := objc.SendIfResponds[int64](t.ID, objc.Sel("gender"))
	return rv
}
func (t TTSAXResource) SetGender(value int64) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setGender:"), value)
}
func (t TTSAXResource) Identifier() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSAXResource) SetIdentifier(value string) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setIdentifier:"), objc.String(value))
}
func (t TTSAXResource) IsDefault() bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("isDefault"))
	return rv
}
func (t TTSAXResource) SetIsDefault(value bool) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setIsDefault:"), value)
}
func (t TTSAXResource) IsNoveltyVoice() bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("isNoveltyVoice"))
	return rv
}
func (t TTSAXResource) SetIsNoveltyVoice(value bool) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setIsNoveltyVoice:"), value)
}
func (t TTSAXResource) IsPersonalVoice() bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("isPersonalVoice"))
	return rv
}
func (t TTSAXResource) SetIsPersonalVoice(value bool) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setIsPersonalVoice:"), value)
}
func (t TTSAXResource) IsSystemVoice() bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("isSystemVoice"))
	return rv
}
func (t TTSAXResource) SetIsSystemVoice(value bool) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setIsSystemVoice:"), value)
}
func (t TTSAXResource) Language() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("language"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSAXResource) SetLanguage(value string) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setLanguage:"), objc.String(value))
}
func (t TTSAXResource) Languages() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("languages"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSAXResource) SetLanguages(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setLanguages:"), value)
}
func (t TTSAXResource) LocalizedName() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("localizedName"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSAXResource) SetLocalizedName(value string) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setLocalizedName:"), objc.String(value))
}
func (t TTSAXResource) LocalizedNameWithFootprint() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("localizedNameWithFootprint"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSAXResource) SetLocalizedNameWithFootprint(value string) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setLocalizedNameWithFootprint:"), objc.String(value))
}
func (t TTSAXResource) Name() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSAXResource) SetName(value string) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setName:"), objc.String(value))
}
func (t TTSAXResource) Subtype() uint64 {
	rv := objc.SendIfResponds[uint64](t.ID, objc.Sel("subtype"))
	return rv
}
func (t TTSAXResource) SetSubtype(value uint64) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setSubtype:"), value)
}
func (t TTSAXResource) SynthesisProviderVoice() avfaudio.AVSpeechSynthesisProviderVoice {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("synthesisProviderVoice"))
	return avfaudio.AVSpeechSynthesisProviderVoiceFromID(objc.ID(rv))
}
func (t TTSAXResource) SetSynthesisProviderVoice(value avfaudio.AVSpeechSynthesisProviderVoice) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setSynthesisProviderVoice:"), value)
}
func (t TTSAXResource) Type() uint64 {
	rv := objc.SendIfResponds[uint64](t.ID, objc.Sel("type"))
	return rv
}
func (t TTSAXResource) SetType(value uint64) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setType:"), value)
}
func (t TTSAXResource) VoiceId() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("voiceId"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSAXResource) SetVoiceId(value string) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setVoiceId:"), objc.String(value))
}
func (t TTSAXResource) VoiceType() int64 {
	rv := objc.SendIfResponds[int64](t.ID, objc.Sel("voiceType"))
	return rv
}
func (t TTSAXResource) SetVoiceType(value int64) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setVoiceType:"), value)
}
