// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSFirstPartyAudioUnit] class.
var (
	_TTSFirstPartyAudioUnitClass     TTSFirstPartyAudioUnitClass
	_TTSFirstPartyAudioUnitClassOnce sync.Once
)

func getTTSFirstPartyAudioUnitClass() TTSFirstPartyAudioUnitClass {
	_TTSFirstPartyAudioUnitClassOnce.Do(func() {
		_TTSFirstPartyAudioUnitClass = TTSFirstPartyAudioUnitClass{class: objc.GetClass("TTSFirstPartyAudioUnit")}
	})
	return _TTSFirstPartyAudioUnitClass
}

// GetTTSFirstPartyAudioUnitClass returns the class object for TTSFirstPartyAudioUnit.
func GetTTSFirstPartyAudioUnitClass() TTSFirstPartyAudioUnitClass {
	return getTTSFirstPartyAudioUnitClass()
}

type TTSFirstPartyAudioUnitClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSFirstPartyAudioUnitClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSFirstPartyAudioUnitClass) Alloc() TTSFirstPartyAudioUnit {
	rv := objc.SendIfResponds[TTSFirstPartyAudioUnit](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TTSFirstPartyAudioUnit.Channel]
//   - [TTSFirstPartyAudioUnit.SetChannel]
//   - [TTSFirstPartyAudioUnit.DefaultSettingsForVoice]
//   - [TTSFirstPartyAudioUnit.Echo]
//   - [TTSFirstPartyAudioUnit.MessageChannelFor]
//   - [TTSFirstPartyAudioUnit.PrewarmWithVoice]
//   - [TTSFirstPartyAudioUnit.RequireFirstUnlockForVoiceLoad]
//   - [TTSFirstPartyAudioUnit.VoicesExternallyManaged]
type TTSFirstPartyAudioUnit struct {
	objectivec.Object
}

// TTSFirstPartyAudioUnitFromID constructs a [TTSFirstPartyAudioUnit] from an objc.ID.
func TTSFirstPartyAudioUnitFromID(id objc.ID) TTSFirstPartyAudioUnit {
	return TTSFirstPartyAudioUnit{objectivec.Object{ID: id}}
}

// NOTE: TTSFirstPartyAudioUnit embeds objectivec.Object because the parent type is
// unavailable, but ITTSFirstPartyAudioUnit embeds IAVSpeechSynthesisProviderAudioUnit, which that fallback
// cannot satisfy; skip compile-time assertion.

// An interface definition for the [TTSFirstPartyAudioUnit] class.
//
// # Methods
//
//   - [ITTSFirstPartyAudioUnit.Channel]
//   - [ITTSFirstPartyAudioUnit.SetChannel]
//   - [ITTSFirstPartyAudioUnit.DefaultSettingsForVoice]
//   - [ITTSFirstPartyAudioUnit.Echo]
//   - [ITTSFirstPartyAudioUnit.MessageChannelFor]
//   - [ITTSFirstPartyAudioUnit.PrewarmWithVoice]
//   - [ITTSFirstPartyAudioUnit.RequireFirstUnlockForVoiceLoad]
//   - [ITTSFirstPartyAudioUnit.VoicesExternallyManaged]
type ITTSFirstPartyAudioUnit interface {
	IAVSpeechSynthesisProviderAudioUnit

	// Topic: Methods

	Channel() ITTSAUMessagingAU
	SetChannel(value ITTSAUMessagingAU)
	DefaultSettingsForVoice(voice objectivec.IObject) objectivec.IObject
	Echo(echo objectivec.IObject) objectivec.IObject
	MessageChannelFor(for_ objectivec.IObject) objectivec.IObject
	PrewarmWithVoice(voice objectivec.IObject)
	RequireFirstUnlockForVoiceLoad() objectivec.IObject
	VoicesExternallyManaged() objectivec.IObject
}

// Init initializes the instance.
func (t TTSFirstPartyAudioUnit) Init() TTSFirstPartyAudioUnit {
	rv := objc.SendIfResponds[TTSFirstPartyAudioUnit](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSFirstPartyAudioUnit) Autorelease() TTSFirstPartyAudioUnit {
	rv := objc.SendIfResponds[TTSFirstPartyAudioUnit](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSFirstPartyAudioUnit creates a new TTSFirstPartyAudioUnit instance.
func NewTTSFirstPartyAudioUnit() TTSFirstPartyAudioUnit {
	class := getTTSFirstPartyAudioUnitClass()
	rv := objc.SendIfResponds[TTSFirstPartyAudioUnit](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (t TTSFirstPartyAudioUnit) DefaultSettingsForVoice(voice objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("defaultSettingsForVoice:"), voice)
	return objectivec.Object{ID: rv}
}
func (t TTSFirstPartyAudioUnit) Echo(echo objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("echo:"), echo)
	return objectivec.Object{ID: rv}
}
func (t TTSFirstPartyAudioUnit) MessageChannelFor(for_ objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("messageChannelFor:"), for_)
	return objectivec.Object{ID: rv}
}
func (t TTSFirstPartyAudioUnit) PrewarmWithVoice(voice objectivec.IObject) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("prewarmWithVoice:"), voice)
}
func (t TTSFirstPartyAudioUnit) RequireFirstUnlockForVoiceLoad() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("requireFirstUnlockForVoiceLoad"))
	return objectivec.Object{ID: rv}
}
func (t TTSFirstPartyAudioUnit) VoicesExternallyManaged() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("voicesExternallyManaged"))
	return objectivec.Object{ID: rv}
}

func (_TTSFirstPartyAudioUnitClass TTSFirstPartyAudioUnitClass) RegisterInProcess() {
	objc.SendIfResponds[objc.ID](objc.ID(_TTSFirstPartyAudioUnitClass.class), objc.Sel("registerInProcess"))
}
func (_TTSFirstPartyAudioUnitClass TTSFirstPartyAudioUnitClass) ShouldLogSensitiveSpeech() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_TTSFirstPartyAudioUnitClass.class), objc.Sel("shouldLogSensitiveSpeech"))
	return rv
}

func (t TTSFirstPartyAudioUnit) Channel() ITTSAUMessagingAU {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("channel"))
	return TTSAUMessagingAUFromID(objc.ID(rv))
}
func (t TTSFirstPartyAudioUnit) SetChannel(value ITTSAUMessagingAU) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setChannel:"), value)
}
