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
	rv := objc.Send[TTSFirstPartyAudioUnit](objc.ID(tc.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSFirstPartyAudioUnit
type TTSFirstPartyAudioUnit struct {
	objectivec.Object
}

// TTSFirstPartyAudioUnitFromID constructs a [TTSFirstPartyAudioUnit] from an objc.ID.
func TTSFirstPartyAudioUnitFromID(id objc.ID) TTSFirstPartyAudioUnit {
	return TTSFirstPartyAudioUnit{objectivec.Object{ID: id}}
}

// NOTE: TTSFirstPartyAudioUnit struct embeds objectivec.Object (parent type unavailable) but
// ITTSFirstPartyAudioUnit embeds the parent interface; skip compile-time assertion.

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
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSFirstPartyAudioUnit
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
	rv := objc.Send[TTSFirstPartyAudioUnit](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSFirstPartyAudioUnit) Autorelease() TTSFirstPartyAudioUnit {
	rv := objc.Send[TTSFirstPartyAudioUnit](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSFirstPartyAudioUnit creates a new TTSFirstPartyAudioUnit instance.
func NewTTSFirstPartyAudioUnit() TTSFirstPartyAudioUnit {
	class := getTTSFirstPartyAudioUnitClass()
	rv := objc.Send[TTSFirstPartyAudioUnit](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSFirstPartyAudioUnit/defaultSettingsForVoice:
func (t TTSFirstPartyAudioUnit) DefaultSettingsForVoice(voice objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("defaultSettingsForVoice:"), voice)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSFirstPartyAudioUnit/echo:
func (t TTSFirstPartyAudioUnit) Echo(echo objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("echo:"), echo)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSFirstPartyAudioUnit/messageChannelFor:
func (t TTSFirstPartyAudioUnit) MessageChannelFor(for_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("messageChannelFor:"), for_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSFirstPartyAudioUnit/prewarmWithVoice:
func (t TTSFirstPartyAudioUnit) PrewarmWithVoice(voice objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("prewarmWithVoice:"), voice)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSFirstPartyAudioUnit/requireFirstUnlockForVoiceLoad
func (t TTSFirstPartyAudioUnit) RequireFirstUnlockForVoiceLoad() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("requireFirstUnlockForVoiceLoad"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSFirstPartyAudioUnit/voicesExternallyManaged
func (t TTSFirstPartyAudioUnit) VoicesExternallyManaged() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("voicesExternallyManaged"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSFirstPartyAudioUnit/registerInProcess
func (_TTSFirstPartyAudioUnitClass TTSFirstPartyAudioUnitClass) RegisterInProcess() {
	objc.Send[objc.ID](objc.ID(_TTSFirstPartyAudioUnitClass.class), objc.Sel("registerInProcess"))
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSFirstPartyAudioUnit/shouldLogSensitiveSpeech
func (_TTSFirstPartyAudioUnitClass TTSFirstPartyAudioUnitClass) ShouldLogSensitiveSpeech() bool {
	rv := objc.Send[bool](objc.ID(_TTSFirstPartyAudioUnitClass.class), objc.Sel("shouldLogSensitiveSpeech"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSFirstPartyAudioUnit/channel
func (t TTSFirstPartyAudioUnit) Channel() ITTSAUMessagingAU {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("channel"))
	return TTSAUMessagingAUFromID(objc.ID(rv))
}
func (t TTSFirstPartyAudioUnit) SetChannel(value ITTSAUMessagingAU) {
	objc.Send[struct{}](t.ID, objc.Sel("setChannel:"), value)
}
