// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
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

// A distinct voice for use in speech synthesis.
//
// # Overview
//
// The primary factors that distinguish a voice in speech synthesis are
// language, locale, and quality. Create an instance of
// [AVSpeechSynthesisVoice] to select a voice that’s appropriate for the
// text and the language, and set it as the value of the [AVSpeechSynthesisVoice.Voice] property on
// an [AVSpeechUtterance] instance. The voice may optionally reflect a local
// variant of the language, such as Australian or South African English. For a
// complete list of supported languages, see [Languages Supported by
// VoiceOver].
//
// # Obtaining voices
//
//   - [AVSpeechSynthesisVoice.AVSpeechSynthesisVoiceIdentifierAlex]: The voice that the system identifies as Alex.
//
// # Inspecting voices
//
//   - [AVSpeechSynthesisVoice.Identifier]: The unique identifier of a voice.
//   - [AVSpeechSynthesisVoice.Name]: The name of a voice.
//   - [AVSpeechSynthesisVoice.Quality]: The speech quality of a voice.
//   - [AVSpeechSynthesisVoice.Gender]: The gender for a voice.
//   - [AVSpeechSynthesisVoice.VoiceTraits]: The traits of a voice.
//   - [AVSpeechSynthesisVoice.AudioFileSettings]: A dictionary that contains audio file settings.
//
// # Working with language codes
//
//   - [AVSpeechSynthesisVoice.Language]: A BCP 47 code that contains the voice’s language and locale.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice
//
// [Languages Supported by VoiceOver]: https://support.apple.com/en-us/HT206175
type AVSpeechSynthesisVoice struct {
	objectivec.Object
}

// AVSpeechSynthesisVoiceFromID constructs a [AVSpeechSynthesisVoice] from an objc.ID.
//
// A distinct voice for use in speech synthesis.
func AVSpeechSynthesisVoiceFromID(id objc.ID) AVSpeechSynthesisVoice {
	return AVSpeechSynthesisVoice{objectivec.Object{ID: id}}
}

// NOTE: AVSpeechSynthesisVoice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVSpeechSynthesisVoice] class.
//
// # Obtaining voices
//
//   - [IAVSpeechSynthesisVoice.AVSpeechSynthesisVoiceIdentifierAlex]: The voice that the system identifies as Alex.
//
// # Inspecting voices
//
//   - [IAVSpeechSynthesisVoice.Identifier]: The unique identifier of a voice.
//   - [IAVSpeechSynthesisVoice.Name]: The name of a voice.
//   - [IAVSpeechSynthesisVoice.Quality]: The speech quality of a voice.
//   - [IAVSpeechSynthesisVoice.Gender]: The gender for a voice.
//   - [IAVSpeechSynthesisVoice.VoiceTraits]: The traits of a voice.
//   - [IAVSpeechSynthesisVoice.AudioFileSettings]: A dictionary that contains audio file settings.
//
// # Working with language codes
//
//   - [IAVSpeechSynthesisVoice.Language]: A BCP 47 code that contains the voice’s language and locale.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice
type IAVSpeechSynthesisVoice interface {
	objectivec.IObject

	// Topic: Obtaining voices

	// The voice that the system identifies as Alex.
	AVSpeechSynthesisVoiceIdentifierAlex() string

	// Topic: Inspecting voices

	// The unique identifier of a voice.
	Identifier() string
	// The name of a voice.
	Name() string
	// The speech quality of a voice.
	Quality() AVSpeechSynthesisVoiceQuality
	// The gender for a voice.
	Gender() AVSpeechSynthesisVoiceGender
	// The traits of a voice.
	VoiceTraits() AVSpeechSynthesisVoiceTraits
	// A dictionary that contains audio file settings.
	AudioFileSettings() foundation.INSDictionary

	// Topic: Working with language codes

	// A BCP 47 code that contains the voice’s language and locale.
	Language() string

	// The voice the speech synthesizer uses when speaking the utterance.
	Voice() IAVSpeechSynthesisVoice
	SetVoice(value IAVSpeechSynthesisVoice)
	EncodeWithCoder(coder foundation.INSCoder)
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

// Retrieves a voice for the identifier you specify.
//
// identifier: The unique identifier for a voice.
//
// # Return Value
//
// A voice for the specified identifier if the identifier is valid and the
// voice is available on the device; otherwise, `nil`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/init(identifier:)
func NewSpeechSynthesisVoiceWithIdentifier(identifier string) AVSpeechSynthesisVoice {
	rv := objc.Send[objc.ID](objc.ID(getAVSpeechSynthesisVoiceClass().class), objc.Sel("voiceWithIdentifier:"), objc.String(identifier))
	return AVSpeechSynthesisVoiceFromID(rv)
}

// Retrieves a voice for the BCP 47 code language code you specify.
//
// # Return Value
//
// A voice for the specified language and locale code if the code is valid;
// otherwise, `nil`.
//
// # Discussion
//
// - languageCode: A BCP 47 code that identifies the language and locale for a
// voice.
//
// # Discussion
//
// Pass `nil` for `languageCode` to receive the default voice for the
// system’s language and region.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/init(language:)
func NewSpeechSynthesisVoiceWithLanguage(languageCode string) AVSpeechSynthesisVoice {
	rv := objc.Send[objc.ID](objc.ID(getAVSpeechSynthesisVoiceClass().class), objc.Sel("voiceWithLanguage:"), objc.String(languageCode))
	return AVSpeechSynthesisVoiceFromID(rv)
}

func (s AVSpeechSynthesisVoice) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Retrieves all available voices on the device.
//
// # Return Value
//
// An array of voices.
//
// # Discussion
//
// Use the [Language] property to identify each voice by its language and
// locale.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/speechVoices()
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) SpeechVoices() []AVSpeechSynthesisVoice {
	rv := objc.Send[[]objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("speechVoices"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVSpeechSynthesisVoice {
		return AVSpeechSynthesisVoiceFromID(id)
	})
}

// Returns the language and locale code for the user’s current locale.
//
// # Return Value
//
// A string that contains the BCP 47 language and locale code for the user’s
// current locale.
//
// # Discussion
//
// This code reflects the user’s language and region preferences in the
// Settings app.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/currentLanguageCode()
func (_AVSpeechSynthesisVoiceClass AVSpeechSynthesisVoiceClass) CurrentLanguageCode() string {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisVoiceClass.class), objc.Sel("currentLanguageCode"))
	return foundation.NSStringFromID(rv).String()
}

// The voice that the system identifies as Alex.
//
// See: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisvoiceidentifieralex
func (s AVSpeechSynthesisVoice) AVSpeechSynthesisVoiceIdentifierAlex() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("AVSpeechSynthesisVoiceIdentifierAlex"))
	return foundation.NSStringFromID(rv).String()
}

// The unique identifier of a voice.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/identifier
func (s AVSpeechSynthesisVoice) Identifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}

// The name of a voice.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/name
func (s AVSpeechSynthesisVoice) Name() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The speech quality of a voice.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/quality
func (s AVSpeechSynthesisVoice) Quality() AVSpeechSynthesisVoiceQuality {
	rv := objc.Send[AVSpeechSynthesisVoiceQuality](s.ID, objc.Sel("quality"))
	return AVSpeechSynthesisVoiceQuality(rv)
}

// The gender for a voice.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/gender
func (s AVSpeechSynthesisVoice) Gender() AVSpeechSynthesisVoiceGender {
	rv := objc.Send[AVSpeechSynthesisVoiceGender](s.ID, objc.Sel("gender"))
	return AVSpeechSynthesisVoiceGender(rv)
}

// The traits of a voice.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/voiceTraits
func (s AVSpeechSynthesisVoice) VoiceTraits() AVSpeechSynthesisVoiceTraits {
	rv := objc.Send[AVSpeechSynthesisVoiceTraits](s.ID, objc.Sel("voiceTraits"))
	return AVSpeechSynthesisVoiceTraits(rv)
}

// A dictionary that contains audio file settings.
//
// # Discussion
//
// If you want to generate speech and save it as an audio file to share or
// play later, use this dictionary to create an [AVAudioFile] instance and
// pass it as the `settings` parameter.
//
// You can determine the [AVAudioCommonFormat] and interleaved properties of a
// voice from this dictionary. The format of this dictionary matches the data
// that [AVSpeechSynthesizerBufferCallback] provides for the same voice.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/audioFileSettings
//
// [AVAudioCommonFormat]: https://developer.apple.com/documentation/AVFAudio/AVAudioCommonFormat
func (s AVSpeechSynthesisVoice) AudioFileSettings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("audioFileSettings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// A BCP 47 code that contains the voice’s language and locale.
//
// # Discussion
//
// The language of a voice controls the conversion of text to spoken phonemes.
// For best results, ensure that the language of an utterance’s text matches
// the voice for the utterance. The locale of a voice reflects regional
// variations in pronunciation or accent. For example, a voice with a language
// code of `en-US` speaks English text with a North American accent, and a
// language code of `en-AU` speaks English text with an Australian accent.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/language
func (s AVSpeechSynthesisVoice) Language() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("language"))
	return foundation.NSStringFromID(rv).String()
}

// The voice the speech synthesizer uses when speaking the utterance.
//
// See: https://developer.apple.com/documentation/avfaudio/avspeechutterance/voice
func (s AVSpeechSynthesisVoice) Voice() IAVSpeechSynthesisVoice {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("voice"))
	return AVSpeechSynthesisVoiceFromID(objc.ID(rv))
}
func (s AVSpeechSynthesisVoice) SetVoice(value IAVSpeechSynthesisVoice) {
	objc.Send[struct{}](s.ID, objc.Sel("setVoice:"), value)
}
