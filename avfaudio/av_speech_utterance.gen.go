// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSpeechUtterance] class.
var (
	_AVSpeechUtteranceClass     AVSpeechUtteranceClass
	_AVSpeechUtteranceClassOnce sync.Once
)

func getAVSpeechUtteranceClass() AVSpeechUtteranceClass {
	_AVSpeechUtteranceClassOnce.Do(func() {
		_AVSpeechUtteranceClass = AVSpeechUtteranceClass{class: objc.GetClass("AVSpeechUtterance")}
	})
	return _AVSpeechUtteranceClass
}

// GetAVSpeechUtteranceClass returns the class object for AVSpeechUtterance.
func GetAVSpeechUtteranceClass() AVSpeechUtteranceClass {
	return getAVSpeechUtteranceClass()
}

type AVSpeechUtteranceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSpeechUtteranceClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSpeechUtteranceClass) Alloc() AVSpeechUtterance {
	rv := objc.Send[AVSpeechUtterance](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that encapsulates the text for speech synthesis and parameters
// that affect the speech.
//
// # Overview
// 
// An [AVSpeechUtterance] is the basic unit of speech synthesis.
// 
// To synthesize speech, create an [AVSpeechUtterance] instance with text you
// want a speech synthesizer to speak. Optionally, change the [AVSpeechUtterance.Voice],
// [AVSpeechUtterance.PitchMultiplier], [AVSpeechUtterance.Volume], [AVSpeechUtterance.Rate], [AVSpeechUtterance.PreUtteranceDelay], or
// [AVSpeechUtterance.PostUtteranceDelay] parameters for the utterance. Pass the utterance to an
// instance of [AVSpeechSynthesizer] to begin speech, or enqueue the utterance
// to speak later if the synthesizer is already speaking.
// 
// Split a body of text into multiple utterances if you want to apply
// different speech parameters. For example, you can emphasize a sentence by
// increasing the pitch and decreasing the rate of that utterance relative to
// others, or you can introduce pauses between sentences by putting each into
// an utterance with a leading or trailing delay.
// 
// Set and use the [AVSpeechSynthesizerDelegate] to receive notifications when
// the synthesizer starts or finishes speaking an utterance. Create an
// utterance for each meaningful unit in a body of text if you want to receive
// notifications as its speech progresses.
//
// # Creating an utterance
//
//   - [AVSpeechUtterance.InitWithString]: Creates an utterance with the text string that you specify for the speech synthesizer to speak.
//   - [AVSpeechUtterance.InitWithAttributedString]: Creates an utterance with the attributed text string that you specify for the speech synthesizer to speak.
//   - [AVSpeechUtterance.AVSpeechSynthesisIPANotationAttribute]: A string that contains International Phonetic Alphabet (IPA) symbols the speech synthesizer uses to control pronunciation of certain words or phrases.
//   - [AVSpeechUtterance.InitWithSSMLRepresentation]: Creates a speech utterance with an Speech Synthesis Markup Language (SSML) string.
//
// # Configuring an utterance
//
//   - [AVSpeechUtterance.Voice]: The voice the speech synthesizer uses when speaking the utterance.
//   - [AVSpeechUtterance.SetVoice]
//   - [AVSpeechUtterance.PitchMultiplier]: The baseline pitch the speech synthesizer uses when speaking the utterance.
//   - [AVSpeechUtterance.SetPitchMultiplier]
//   - [AVSpeechUtterance.Volume]: The volume the speech synthesizer uses when speaking the utterance.
//   - [AVSpeechUtterance.SetVolume]
//   - [AVSpeechUtterance.PrefersAssistiveTechnologySettings]: A Boolean that specifies whether assistive technology settings take precedence over the property values of this utterance.
//   - [AVSpeechUtterance.SetPrefersAssistiveTechnologySettings]
//
// # Configuring utterance timing
//
//   - [AVSpeechUtterance.Rate]: The rate the speech synthesizer uses when speaking the utterance.
//   - [AVSpeechUtterance.SetRate]
//   - [AVSpeechUtterance.AVSpeechUtteranceMinimumSpeechRate]: The minimum rate the speech synthesizer uses when speaking an utterance.
//   - [AVSpeechUtterance.AVSpeechUtteranceMaximumSpeechRate]: The maximum rate the speech synthesizer uses when speaking an utterance.
//   - [AVSpeechUtterance.AVSpeechUtteranceDefaultSpeechRate]: The default rate the speech synthesizer uses when speaking an utterance.
//   - [AVSpeechUtterance.PreUtteranceDelay]: The amount of time the speech synthesizer pauses before speaking the utterance.
//   - [AVSpeechUtterance.SetPreUtteranceDelay]
//   - [AVSpeechUtterance.PostUtteranceDelay]: The amount of time the speech synthesizer pauses after speaking an utterance before handling the next utterance in the queue.
//   - [AVSpeechUtterance.SetPostUtteranceDelay]
//
// # Inspecting utterance text
//
//   - [AVSpeechUtterance.SpeechString]: A string that contains the text for speech synthesis.
//   - [AVSpeechUtterance.AttributedSpeechString]: An attributed string that contains the text for speech synthesis.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance
type AVSpeechUtterance struct {
	objectivec.Object
}

// AVSpeechUtteranceFromID constructs a [AVSpeechUtterance] from an objc.ID.
//
// An object that encapsulates the text for speech synthesis and parameters
// that affect the speech.
func AVSpeechUtteranceFromID(id objc.ID) AVSpeechUtterance {
	return AVSpeechUtterance{objectivec.Object{ID: id}}
}
// NOTE: AVSpeechUtterance adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVSpeechUtterance] class.
//
// # Creating an utterance
//
//   - [IAVSpeechUtterance.InitWithString]: Creates an utterance with the text string that you specify for the speech synthesizer to speak.
//   - [IAVSpeechUtterance.InitWithAttributedString]: Creates an utterance with the attributed text string that you specify for the speech synthesizer to speak.
//   - [IAVSpeechUtterance.AVSpeechSynthesisIPANotationAttribute]: A string that contains International Phonetic Alphabet (IPA) symbols the speech synthesizer uses to control pronunciation of certain words or phrases.
//   - [IAVSpeechUtterance.InitWithSSMLRepresentation]: Creates a speech utterance with an Speech Synthesis Markup Language (SSML) string.
//
// # Configuring an utterance
//
//   - [IAVSpeechUtterance.Voice]: The voice the speech synthesizer uses when speaking the utterance.
//   - [IAVSpeechUtterance.SetVoice]
//   - [IAVSpeechUtterance.PitchMultiplier]: The baseline pitch the speech synthesizer uses when speaking the utterance.
//   - [IAVSpeechUtterance.SetPitchMultiplier]
//   - [IAVSpeechUtterance.Volume]: The volume the speech synthesizer uses when speaking the utterance.
//   - [IAVSpeechUtterance.SetVolume]
//   - [IAVSpeechUtterance.PrefersAssistiveTechnologySettings]: A Boolean that specifies whether assistive technology settings take precedence over the property values of this utterance.
//   - [IAVSpeechUtterance.SetPrefersAssistiveTechnologySettings]
//
// # Configuring utterance timing
//
//   - [IAVSpeechUtterance.Rate]: The rate the speech synthesizer uses when speaking the utterance.
//   - [IAVSpeechUtterance.SetRate]
//   - [IAVSpeechUtterance.AVSpeechUtteranceMinimumSpeechRate]: The minimum rate the speech synthesizer uses when speaking an utterance.
//   - [IAVSpeechUtterance.AVSpeechUtteranceMaximumSpeechRate]: The maximum rate the speech synthesizer uses when speaking an utterance.
//   - [IAVSpeechUtterance.AVSpeechUtteranceDefaultSpeechRate]: The default rate the speech synthesizer uses when speaking an utterance.
//   - [IAVSpeechUtterance.PreUtteranceDelay]: The amount of time the speech synthesizer pauses before speaking the utterance.
//   - [IAVSpeechUtterance.SetPreUtteranceDelay]
//   - [IAVSpeechUtterance.PostUtteranceDelay]: The amount of time the speech synthesizer pauses after speaking an utterance before handling the next utterance in the queue.
//   - [IAVSpeechUtterance.SetPostUtteranceDelay]
//
// # Inspecting utterance text
//
//   - [IAVSpeechUtterance.SpeechString]: A string that contains the text for speech synthesis.
//   - [IAVSpeechUtterance.AttributedSpeechString]: An attributed string that contains the text for speech synthesis.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance
type IAVSpeechUtterance interface {
	objectivec.IObject

	// Topic: Creating an utterance

	// Creates an utterance with the text string that you specify for the speech synthesizer to speak.
	InitWithString(string_ string) AVSpeechUtterance
	// Creates an utterance with the attributed text string that you specify for the speech synthesizer to speak.
	InitWithAttributedString(string_ foundation.NSAttributedString) AVSpeechUtterance
	// A string that contains International Phonetic Alphabet (IPA) symbols the speech synthesizer uses to control pronunciation of certain words or phrases.
	AVSpeechSynthesisIPANotationAttribute() string
	// Creates a speech utterance with an Speech Synthesis Markup Language (SSML) string.
	InitWithSSMLRepresentation(string_ string) AVSpeechUtterance

	// Topic: Configuring an utterance

	// The voice the speech synthesizer uses when speaking the utterance.
	Voice() IAVSpeechSynthesisVoice
	SetVoice(value IAVSpeechSynthesisVoice)
	// The baseline pitch the speech synthesizer uses when speaking the utterance.
	PitchMultiplier() float32
	SetPitchMultiplier(value float32)
	// The volume the speech synthesizer uses when speaking the utterance.
	Volume() float32
	SetVolume(value float32)
	// A Boolean that specifies whether assistive technology settings take precedence over the property values of this utterance.
	PrefersAssistiveTechnologySettings() bool
	SetPrefersAssistiveTechnologySettings(value bool)

	// Topic: Configuring utterance timing

	// The rate the speech synthesizer uses when speaking the utterance.
	Rate() float32
	SetRate(value float32)
	// The minimum rate the speech synthesizer uses when speaking an utterance.
	AVSpeechUtteranceMinimumSpeechRate() float32
	// The maximum rate the speech synthesizer uses when speaking an utterance.
	AVSpeechUtteranceMaximumSpeechRate() float32
	// The default rate the speech synthesizer uses when speaking an utterance.
	AVSpeechUtteranceDefaultSpeechRate() float32
	// The amount of time the speech synthesizer pauses before speaking the utterance.
	PreUtteranceDelay() float64
	SetPreUtteranceDelay(value float64)
	// The amount of time the speech synthesizer pauses after speaking an utterance before handling the next utterance in the queue.
	PostUtteranceDelay() float64
	SetPostUtteranceDelay(value float64)

	// Topic: Inspecting utterance text

	// A string that contains the text for speech synthesis.
	SpeechString() string
	// An attributed string that contains the text for speech synthesis.
	AttributedSpeechString() foundation.NSAttributedString

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s AVSpeechUtterance) Init() AVSpeechUtterance {
	rv := objc.Send[AVSpeechUtterance](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSpeechUtterance) Autorelease() AVSpeechUtterance {
	rv := objc.Send[AVSpeechUtterance](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSpeechUtterance creates a new AVSpeechUtterance instance.
func NewAVSpeechUtterance() AVSpeechUtterance {
	class := getAVSpeechUtteranceClass()
	rv := objc.Send[AVSpeechUtterance](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an utterance with the attributed text string that you specify for
// the speech synthesizer to speak.
//
// string: A string that contains the text to speak.
//
// # Discussion
// 
// To speak the text, pass the utterance to an instance of
// [AVSpeechSynthesizer].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/init(attributedString:)
func NewSpeechUtteranceWithAttributedString(string_ foundation.NSAttributedString) AVSpeechUtterance {
	instance := getAVSpeechUtteranceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAttributedString:"), string_)
	return AVSpeechUtteranceFromID(rv)
}

// Creates a speech utterance with an Speech Synthesis Markup Language (SSML)
// string.
//
// string: A string to speak that contains valid SSML markup. The initializer returns
// `nil` if you pass an invalid SSML string.
//
// # Discussion
// 
// If using SSML to request voices that fall under certain attributes, the
// system may split a single utterance into multiple parts and send each to an
// appropriate synthesizer.
// 
// If no voice matches the properties, the utterance uses the voice set in its
// [Voice] property. If you don’t specify a voice, the system uses its
// default voice.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/init(ssmlRepresentation:)
func NewSpeechUtteranceWithSSMLRepresentation(string_ string) AVSpeechUtterance {
	instance := getAVSpeechUtteranceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSSMLRepresentation:"), objc.String(string_))
	return AVSpeechUtteranceFromID(rv)
}

// Creates an utterance with the text string that you specify for the speech
// synthesizer to speak.
//
// string: A string that contains the text to speak.
//
// # Return Value
// 
// An [AVSpeechUtterance] object that can speak the specified text.
//
// # Discussion
// 
// To speak the text, pass the utterance to an instance of
// [AVSpeechSynthesizer].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/init(string:)
func NewSpeechUtteranceWithString(string_ string) AVSpeechUtterance {
	instance := getAVSpeechUtteranceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:"), objc.String(string_))
	return AVSpeechUtteranceFromID(rv)
}

// Creates an utterance with the text string that you specify for the speech
// synthesizer to speak.
//
// string: A string that contains the text to speak.
//
// # Return Value
// 
// An [AVSpeechUtterance] object that can speak the specified text.
//
// # Discussion
// 
// To speak the text, pass the utterance to an instance of
// [AVSpeechSynthesizer].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/init(string:)
func (s AVSpeechUtterance) InitWithString(string_ string) AVSpeechUtterance {
	rv := objc.Send[AVSpeechUtterance](s.ID, objc.Sel("initWithString:"), objc.String(string_))
	return rv
}
// Creates an utterance with the attributed text string that you specify for
// the speech synthesizer to speak.
//
// string: A string that contains the text to speak.
//
// # Discussion
// 
// To speak the text, pass the utterance to an instance of
// [AVSpeechSynthesizer].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/init(attributedString:)
func (s AVSpeechUtterance) InitWithAttributedString(string_ foundation.NSAttributedString) AVSpeechUtterance {
	rv := objc.Send[AVSpeechUtterance](s.ID, objc.Sel("initWithAttributedString:"), string_)
	return rv
}
// Creates a speech utterance with an Speech Synthesis Markup Language (SSML)
// string.
//
// string: A string to speak that contains valid SSML markup. The initializer returns
// `nil` if you pass an invalid SSML string.
//
// # Discussion
// 
// If using SSML to request voices that fall under certain attributes, the
// system may split a single utterance into multiple parts and send each to an
// appropriate synthesizer.
// 
// If no voice matches the properties, the utterance uses the voice set in its
// [Voice] property. If you don’t specify a voice, the system uses its
// default voice.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/init(ssmlRepresentation:)
func (s AVSpeechUtterance) InitWithSSMLRepresentation(string_ string) AVSpeechUtterance {
	rv := objc.Send[AVSpeechUtterance](s.ID, objc.Sel("initWithSSMLRepresentation:"), objc.String(string_))
	return rv
}
func (s AVSpeechUtterance) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates an utterance with the attributed text string that you specify for
// the speech synthesizer to speak.
//
// string: A string that contains the text to speak.
//
// # Return Value
// 
// An [AVSpeechUtterance] object that can speak the specified text.
//
// # Discussion
// 
// To speak the text, pass the utterance to an instance of
// [AVSpeechSynthesizer].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/speechUtteranceWithAttributedString:
func (_AVSpeechUtteranceClass AVSpeechUtteranceClass) SpeechUtteranceWithAttributedString(string_ foundation.NSAttributedString) AVSpeechUtterance {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechUtteranceClass.class), objc.Sel("speechUtteranceWithAttributedString:"), string_)
	return AVSpeechUtteranceFromID(rv)
}
// Returns a new speech utterance with an Speech Synthesis Markup Language
// (SSML) string.
//
// string: A string to speak that contains valid SSML markup. The initializer returns
// `nil` if you pass an invalid SSML string.
//
// # Return Value
// 
// A new speech utterance, or [nil] if the SSML string is invalid.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// # Discussion
// 
// If using SSML to request voices that fall under certain attributes, the
// system may split a single utterance into multiple parts and send each to an
// appropriate synthesizer.
// 
// If no voice matches the properties, the utterance uses the voice set in its
// [Voice] property. If you don’t specify a voice, the system uses its
// default voice.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/speechUtteranceWithSSMLRepresentation:
func (_AVSpeechUtteranceClass AVSpeechUtteranceClass) SpeechUtteranceWithSSMLRepresentation(string_ string) AVSpeechUtterance {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechUtteranceClass.class), objc.Sel("speechUtteranceWithSSMLRepresentation:"), objc.String(string_))
	return AVSpeechUtteranceFromID(rv)
}
// Creates an utterance with the text string that you specify for the speech
// synthesizer to speak.
//
// string: A string that contains the text to speak.
//
// # Return Value
// 
// An [AVSpeechUtterance] object that can speak the specified text.
//
// # Discussion
// 
// To speak the text, pass the utterance to an instance of
// [AVSpeechSynthesizer].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/speechUtteranceWithString:
func (_AVSpeechUtteranceClass AVSpeechUtteranceClass) SpeechUtteranceWithString(string_ string) AVSpeechUtterance {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechUtteranceClass.class), objc.Sel("speechUtteranceWithString:"), objc.String(string_))
	return AVSpeechUtteranceFromID(rv)
}

// A string that contains International Phonetic Alphabet (IPA) symbols the
// speech synthesizer uses to control pronunciation of certain words or
// phrases.
//
// See: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisipanotationattribute
func (s AVSpeechUtterance) AVSpeechSynthesisIPANotationAttribute() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("AVSpeechSynthesisIPANotationAttribute"))
	return foundation.NSStringFromID(rv).String()
}
// The voice the speech synthesizer uses when speaking the utterance.
//
// # Discussion
// 
// If you don’t specify a voice, the speech synthesizer uses the system’s
// default voice to speak the utterance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/voice
func (s AVSpeechUtterance) Voice() IAVSpeechSynthesisVoice {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("voice"))
	return AVSpeechSynthesisVoiceFromID(objc.ID(rv))
}
func (s AVSpeechUtterance) SetVoice(value IAVSpeechSynthesisVoice) {
	objc.Send[struct{}](s.ID, objc.Sel("setVoice:"), value)
}
// The baseline pitch the speech synthesizer uses when speaking the utterance.
//
// # Discussion
// 
// Before enqueing the utterance, set this property to a value within the
// range of `0.5` for lower pitch to `2.0` for higher pitch. The default value
// is `1.0`. Setting this after enqueing the utterance has no effect.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/pitchMultiplier
func (s AVSpeechUtterance) PitchMultiplier() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("pitchMultiplier"))
	return rv
}
func (s AVSpeechUtterance) SetPitchMultiplier(value float32) {
	objc.Send[struct{}](s.ID, objc.Sel("setPitchMultiplier:"), value)
}
// The volume the speech synthesizer uses when speaking the utterance.
//
// # Discussion
// 
// Before enqueing the utterance, set this property to a value within the
// range of `0.0` for silent to `1.0` for loudest volume. The default value is
// `1.0`. Setting this after enqueing the utterance has no effect.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/volume
func (s AVSpeechUtterance) Volume() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("volume"))
	return rv
}
func (s AVSpeechUtterance) SetVolume(value float32) {
	objc.Send[struct{}](s.ID, objc.Sel("setVolume:"), value)
}
// A Boolean that specifies whether assistive technology settings take
// precedence over the property values of this utterance.
//
// # Discussion
// 
// If this property is `true`, but no assistive technology, such as VoiceOver,
// is on, the speech synthesizer uses the utterance property values.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/prefersAssistiveTechnologySettings
func (s AVSpeechUtterance) PrefersAssistiveTechnologySettings() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("prefersAssistiveTechnologySettings"))
	return rv
}
func (s AVSpeechUtterance) SetPrefersAssistiveTechnologySettings(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setPrefersAssistiveTechnologySettings:"), value)
}
// The rate the speech synthesizer uses when speaking the utterance.
//
// # Discussion
// 
// The speech rate is a decimal representation within the range of
// [AVSpeechUtteranceMinimumSpeechRate] and
// [AVSpeechUtteranceMaximumSpeechRate]. Lower values correspond to slower
// speech, and higher values correspond to faster speech. The default value is
// [AVSpeechUtteranceDefaultSpeechRate]. Set this property before enqueing the
// utterance because setting it afterward has no effect.
//
// [AVSpeechUtteranceDefaultSpeechRate]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtteranceDefaultSpeechRate
// [AVSpeechUtteranceMaximumSpeechRate]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtteranceMaximumSpeechRate
// [AVSpeechUtteranceMinimumSpeechRate]: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtteranceMinimumSpeechRate
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/rate
func (s AVSpeechUtterance) Rate() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("rate"))
	return rv
}
func (s AVSpeechUtterance) SetRate(value float32) {
	objc.Send[struct{}](s.ID, objc.Sel("setRate:"), value)
}
// The minimum rate the speech synthesizer uses when speaking an utterance.
//
// See: https://developer.apple.com/documentation/avfaudio/avspeechutteranceminimumspeechrate
func (s AVSpeechUtterance) AVSpeechUtteranceMinimumSpeechRate() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("AVSpeechUtteranceMinimumSpeechRate"))
	return rv
}
// The maximum rate the speech synthesizer uses when speaking an utterance.
//
// See: https://developer.apple.com/documentation/avfaudio/avspeechutterancemaximumspeechrate
func (s AVSpeechUtterance) AVSpeechUtteranceMaximumSpeechRate() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("AVSpeechUtteranceMaximumSpeechRate"))
	return rv
}
// The default rate the speech synthesizer uses when speaking an utterance.
//
// See: https://developer.apple.com/documentation/avfaudio/avspeechutterancedefaultspeechrate
func (s AVSpeechUtterance) AVSpeechUtteranceDefaultSpeechRate() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("AVSpeechUtteranceDefaultSpeechRate"))
	return rv
}
// The amount of time the speech synthesizer pauses before speaking the
// utterance.
//
// # Discussion
// 
// When multiple utterances exist in the queue, the speech synthesizer pauses
// a minimum amount of time equal to the sum of the current utterance’s
// [PostUtteranceDelay] and the next utterance’s `preUtteranceDelay`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/preUtteranceDelay
func (s AVSpeechUtterance) PreUtteranceDelay() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("preUtteranceDelay"))
	return rv
}
func (s AVSpeechUtterance) SetPreUtteranceDelay(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setPreUtteranceDelay:"), value)
}
// The amount of time the speech synthesizer pauses after speaking an
// utterance before handling the next utterance in the queue.
//
// # Discussion
// 
// When multiple utterances exist in the queue, the speech synthesizer pauses
// a minimum amount of time equal to the sum of the current utterance’s
// `postUtteranceDelay` and the next utterance’s [PreUtteranceDelay].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/postUtteranceDelay
func (s AVSpeechUtterance) PostUtteranceDelay() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("postUtteranceDelay"))
	return rv
}
func (s AVSpeechUtterance) SetPostUtteranceDelay(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setPostUtteranceDelay:"), value)
}
// A string that contains the text for speech synthesis.
//
// # Discussion
// 
// You can’t change an utterance’s text after initializaiton. If you want
// the speech synthesizer to speak different text, create a new utterance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/speechString
func (s AVSpeechUtterance) SpeechString() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("speechString"))
	return foundation.NSStringFromID(rv).String()
}
// An attributed string that contains the text for speech synthesis.
//
// # Discussion
// 
// You can’t change an utterance’s text after initializaiton. If you want
// the speech synthesizer to speak different text, create a new utterance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/attributedSpeechString
func (s AVSpeechUtterance) AttributedSpeechString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("attributedSpeechString"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}

