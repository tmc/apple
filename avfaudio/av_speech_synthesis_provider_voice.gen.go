// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSpeechSynthesisProviderVoice] class.
var (
	_AVSpeechSynthesisProviderVoiceClass     AVSpeechSynthesisProviderVoiceClass
	_AVSpeechSynthesisProviderVoiceClassOnce sync.Once
)

func getAVSpeechSynthesisProviderVoiceClass() AVSpeechSynthesisProviderVoiceClass {
	_AVSpeechSynthesisProviderVoiceClassOnce.Do(func() {
		_AVSpeechSynthesisProviderVoiceClass = AVSpeechSynthesisProviderVoiceClass{class: objc.GetClass("AVSpeechSynthesisProviderVoice")}
	})
	return _AVSpeechSynthesisProviderVoiceClass
}

// GetAVSpeechSynthesisProviderVoiceClass returns the class object for AVSpeechSynthesisProviderVoice.
func GetAVSpeechSynthesisProviderVoiceClass() AVSpeechSynthesisProviderVoiceClass {
	return getAVSpeechSynthesisProviderVoiceClass()
}

type AVSpeechSynthesisProviderVoiceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSpeechSynthesisProviderVoiceClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSpeechSynthesisProviderVoiceClass) Alloc() AVSpeechSynthesisProviderVoice {
	rv := objc.Send[AVSpeechSynthesisProviderVoice](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a voice that an audio unit provides to its host.
//
// # Overview
// 
// This is a voice that an [AVSpeechSynthesisProviderAudioUnit] provides to
// the system, distinct from [AVSpeechSynthesisVoice]. Use [AVSpeechSynthesisProviderVoice.SpeechVoices] to
// access the underlying [AVSpeechSynthesisVoice] in the voice quality
// [SpeechSynthesisVoiceQualityEnhanced].
//
// # Creating a voice
//
//   - [AVSpeechSynthesisProviderVoice.InitWithNameIdentifierPrimaryLanguagesSupportedLanguages]: Creates a voice with a name, an identifier, and language information.
//
// # Inspecting a voice
//
//   - [AVSpeechSynthesisProviderVoice.Age]: The age of the voice, in years.
//   - [AVSpeechSynthesisProviderVoice.SetAge]
//   - [AVSpeechSynthesisProviderVoice.Gender]: The gender of the voice.
//   - [AVSpeechSynthesisProviderVoice.SetGender]
//   - [AVSpeechSynthesisProviderVoice.Identifier]: The unique identifier for the voice.
//   - [AVSpeechSynthesisProviderVoice.Name]: The localized name of the voice.
//   - [AVSpeechSynthesisProviderVoice.PrimaryLanguages]: A list of BCP 47 codes that identify the languages the synthesizer uses.
//   - [AVSpeechSynthesisProviderVoice.SupportedLanguages]: A list of BCP 47 codes that identify the languages a voice supports.
//   - [AVSpeechSynthesisProviderVoice.Version]: The version of the voice.
//   - [AVSpeechSynthesisProviderVoice.SetVersion]
//   - [AVSpeechSynthesisProviderVoice.VoiceSize]: The size of the voice package on disk, in bytes.
//   - [AVSpeechSynthesisProviderVoice.SetVoiceSize]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice
type AVSpeechSynthesisProviderVoice struct {
	objectivec.Object
}

// AVSpeechSynthesisProviderVoiceFromID constructs a [AVSpeechSynthesisProviderVoice] from an objc.ID.
//
// An object that represents a voice that an audio unit provides to its host.
func AVSpeechSynthesisProviderVoiceFromID(id objc.ID) AVSpeechSynthesisProviderVoice {
	return AVSpeechSynthesisProviderVoice{objectivec.Object{ID: id}}
}
// NOTE: AVSpeechSynthesisProviderVoice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVSpeechSynthesisProviderVoice] class.
//
// # Creating a voice
//
//   - [IAVSpeechSynthesisProviderVoice.InitWithNameIdentifierPrimaryLanguagesSupportedLanguages]: Creates a voice with a name, an identifier, and language information.
//
// # Inspecting a voice
//
//   - [IAVSpeechSynthesisProviderVoice.Age]: The age of the voice, in years.
//   - [IAVSpeechSynthesisProviderVoice.SetAge]
//   - [IAVSpeechSynthesisProviderVoice.Gender]: The gender of the voice.
//   - [IAVSpeechSynthesisProviderVoice.SetGender]
//   - [IAVSpeechSynthesisProviderVoice.Identifier]: The unique identifier for the voice.
//   - [IAVSpeechSynthesisProviderVoice.Name]: The localized name of the voice.
//   - [IAVSpeechSynthesisProviderVoice.PrimaryLanguages]: A list of BCP 47 codes that identify the languages the synthesizer uses.
//   - [IAVSpeechSynthesisProviderVoice.SupportedLanguages]: A list of BCP 47 codes that identify the languages a voice supports.
//   - [IAVSpeechSynthesisProviderVoice.Version]: The version of the voice.
//   - [IAVSpeechSynthesisProviderVoice.SetVersion]
//   - [IAVSpeechSynthesisProviderVoice.VoiceSize]: The size of the voice package on disk, in bytes.
//   - [IAVSpeechSynthesisProviderVoice.SetVoiceSize]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice
type IAVSpeechSynthesisProviderVoice interface {
	objectivec.IObject

	// Topic: Creating a voice

	// Creates a voice with a name, an identifier, and language information.
	InitWithNameIdentifierPrimaryLanguagesSupportedLanguages(name string, identifier string, primaryLanguages []string, supportedLanguages []string) AVSpeechSynthesisProviderVoice

	// Topic: Inspecting a voice

	// The age of the voice, in years.
	Age() int
	SetAge(value int)
	// The gender of the voice.
	Gender() AVSpeechSynthesisVoiceGender
	SetGender(value AVSpeechSynthesisVoiceGender)
	// The unique identifier for the voice.
	Identifier() string
	// The localized name of the voice.
	Name() string
	// A list of BCP 47 codes that identify the languages the synthesizer uses.
	PrimaryLanguages() []string
	// A list of BCP 47 codes that identify the languages a voice supports.
	SupportedLanguages() []string
	// The version of the voice.
	Version() string
	SetVersion(value string)
	// The size of the voice package on disk, in bytes.
	VoiceSize() int64
	SetVoiceSize(value int64)

	// A list of voices the audio unit provides to the system.
	SpeechVoices() IAVSpeechSynthesisProviderVoice
	SetSpeechVoices(value IAVSpeechSynthesisProviderVoice)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s AVSpeechSynthesisProviderVoice) Init() AVSpeechSynthesisProviderVoice {
	rv := objc.Send[AVSpeechSynthesisProviderVoice](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSpeechSynthesisProviderVoice) Autorelease() AVSpeechSynthesisProviderVoice {
	rv := objc.Send[AVSpeechSynthesisProviderVoice](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSpeechSynthesisProviderVoice creates a new AVSpeechSynthesisProviderVoice instance.
func NewAVSpeechSynthesisProviderVoice() AVSpeechSynthesisProviderVoice {
	class := getAVSpeechSynthesisProviderVoiceClass()
	rv := objc.Send[AVSpeechSynthesisProviderVoice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a voice with a name, an identifier, and language information.
//
// name: The localized name of the voice.
//
// identifier: The unique identifier for the voice.
//
// primaryLanguages: A list of BCP 47 codes that identify the primary languages.
//
// supportedLanguages: A list of BCP 47 codes that identify the languages the voice supports.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/init(name:identifier:primaryLanguages:supportedLanguages:)
func NewSpeechSynthesisProviderVoiceWithNameIdentifierPrimaryLanguagesSupportedLanguages(name string, identifier string, primaryLanguages []string, supportedLanguages []string) AVSpeechSynthesisProviderVoice {
	instance := getAVSpeechSynthesisProviderVoiceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:identifier:primaryLanguages:supportedLanguages:"), objc.String(name), objc.String(identifier), objectivec.StringSliceToNSArray(primaryLanguages), objectivec.StringSliceToNSArray(supportedLanguages))
	return AVSpeechSynthesisProviderVoiceFromID(rv)
}

// Creates a voice with a name, an identifier, and language information.
//
// name: The localized name of the voice.
//
// identifier: The unique identifier for the voice.
//
// primaryLanguages: A list of BCP 47 codes that identify the primary languages.
//
// supportedLanguages: A list of BCP 47 codes that identify the languages the voice supports.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/init(name:identifier:primaryLanguages:supportedLanguages:)
func (s AVSpeechSynthesisProviderVoice) InitWithNameIdentifierPrimaryLanguagesSupportedLanguages(name string, identifier string, primaryLanguages []string, supportedLanguages []string) AVSpeechSynthesisProviderVoice {
	rv := objc.Send[AVSpeechSynthesisProviderVoice](s.ID, objc.Sel("initWithName:identifier:primaryLanguages:supportedLanguages:"), objc.String(name), objc.String(identifier), objectivec.StringSliceToNSArray(primaryLanguages), objectivec.StringSliceToNSArray(supportedLanguages))
	return rv
}
func (s AVSpeechSynthesisProviderVoice) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Updates the voices your app provides to the system.
//
// # Discussion
// 
// Use this method to inform the system when you add or remove voices.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/updateSpeechVoices()
func (_AVSpeechSynthesisProviderVoiceClass AVSpeechSynthesisProviderVoiceClass) UpdateSpeechVoices() {
	objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisProviderVoiceClass.class), objc.Sel("updateSpeechVoices"))
}

// The age of the voice, in years.
//
// # Discussion
// 
// The system treats this value as a personality trait and defaults to `0`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/age
func (s AVSpeechSynthesisProviderVoice) Age() int {
	rv := objc.Send[int](s.ID, objc.Sel("age"))
	return rv
}
func (s AVSpeechSynthesisProviderVoice) SetAge(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setAge:"), value)
}
// The gender of the voice.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/gender
func (s AVSpeechSynthesisProviderVoice) Gender() AVSpeechSynthesisVoiceGender {
	rv := objc.Send[AVSpeechSynthesisVoiceGender](s.ID, objc.Sel("gender"))
	return AVSpeechSynthesisVoiceGender(rv)
}
func (s AVSpeechSynthesisProviderVoice) SetGender(value AVSpeechSynthesisVoiceGender) {
	objc.Send[struct{}](s.ID, objc.Sel("setGender:"), value)
}
// The unique identifier for the voice.
//
// # Discussion
// 
// Use reverse domain notation to format the identifier. The behavior is
// undefined unless all voices within an extension have a unique identifier.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/identifier
func (s AVSpeechSynthesisProviderVoice) Identifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
// The localized name of the voice.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/name
func (s AVSpeechSynthesisProviderVoice) Name() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
// A list of BCP 47 codes that identify the languages the synthesizer uses.
//
// # Discussion
// 
// These languages are what a voice primarily supports. For example, if the
// primary language is `zh-CN —` with no additional [SupportedLanguages] —
// the system may switch voices to speak a phrase that contains other
// languages. Changing voices depends on user preferences and what
// accessibility feature is using the voice.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/primaryLanguages
func (s AVSpeechSynthesisProviderVoice) PrimaryLanguages() []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("primaryLanguages"))
	return objc.ConvertSliceToStrings(rv)
}
// A list of BCP 47 codes that identify the languages a voice supports.
//
// # Discussion
// 
// These languages are what a voice supports — when given a multi-language
// phrase — without the need to switch voice. For example, if the primary
// language is `zh-CN`, and this value contains `zh-CN` and `en-US`, a
// synthesizer that receives a phrase with both languages would speak the
// entire phrase.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/supportedLanguages
func (s AVSpeechSynthesisProviderVoice) SupportedLanguages() []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("supportedLanguages"))
	return objc.ConvertSliceToStrings(rv)
}
// The version of the voice.
//
// # Discussion
// 
// This value is for your own tracking and doesn’t impact the behavior of
// the system.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/version
func (s AVSpeechSynthesisProviderVoice) Version() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("version"))
	return foundation.NSStringFromID(rv).String()
}
func (s AVSpeechSynthesisProviderVoice) SetVersion(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setVersion:"), objc.String(value))
}
// The size of the voice package on disk, in bytes.
//
// # Discussion
// 
// This value defaults to `0`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/voiceSize
func (s AVSpeechSynthesisProviderVoice) VoiceSize() int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("voiceSize"))
	return rv
}
func (s AVSpeechSynthesisProviderVoice) SetVoiceSize(value int64) {
	objc.Send[struct{}](s.ID, objc.Sel("setVoiceSize:"), value)
}
// A list of voices the audio unit provides to the system.
//
// See: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovideraudiounit/speechvoices
func (s AVSpeechSynthesisProviderVoice) SpeechVoices() IAVSpeechSynthesisProviderVoice {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("speechVoices"))
	return AVSpeechSynthesisProviderVoiceFromID(objc.ID(rv))
}
func (s AVSpeechSynthesisProviderVoice) SetSpeechVoices(value IAVSpeechSynthesisProviderVoice) {
	objc.Send[struct{}](s.ID, objc.Sel("setSpeechVoices:"), value)
}

