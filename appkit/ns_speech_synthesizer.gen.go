// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSpeechSynthesizer] class.
var (
	_NSSpeechSynthesizerClass     NSSpeechSynthesizerClass
	_NSSpeechSynthesizerClassOnce sync.Once
)

func getNSSpeechSynthesizerClass() NSSpeechSynthesizerClass {
	_NSSpeechSynthesizerClassOnce.Do(func() {
		_NSSpeechSynthesizerClass = NSSpeechSynthesizerClass{class: objc.GetClass("NSSpeechSynthesizer")}
	})
	return _NSSpeechSynthesizerClass
}

// GetNSSpeechSynthesizerClass returns the class object for NSSpeechSynthesizer.
func GetNSSpeechSynthesizerClass() NSSpeechSynthesizerClass {
	return getNSSpeechSynthesizerClass()
}

type NSSpeechSynthesizerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSpeechSynthesizerClass) Alloc() NSSpeechSynthesizer {
	rv := objc.Send[NSSpeechSynthesizer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The Cocoa interface to speech synthesis in macOS.
//
// # Overview
// 
// Speech synthesis, also called text-to-speech (TTS), parses text and
// converts it into audible speech. It offers a concurrent feedback mode that
// can be used in concert with or in place of traditional visual and aural
// notifications. For example, your application can use a speech synthesizer
// object to “pronounce” the text of important alert dialogs. Synthesized
// speech has several advantages. It can provide urgent information to users
// without forcing them to shift attention from their current task. And
// because speech doesn’t rely on visual elements for meaning, it is a
// crucial technology for users with vision or attention disabilities.
// 
// In addition, synthesized speech can help save system resources. Because
// sound samples can take up large amounts of room on disk, using text in
// place of sampled sound is extremely efficient, and so a multimedia
// application might use an [NSSpeechSynthesizer] object to provide a
// narration of a QuickTime movie instead of including sampled-sound data on a
// movie track.
// 
// When you create an [NSSpeechSynthesizer] instance using the default
// initializer (`init`), the class uses the selected in System Preferences >
// Speech. Alternatively, you can select a specific voice for an
// [NSSpeechSynthesizer] instance by initializing it with [NSSpeechSynthesizer.InitWithVoice]. To
// begin synthesis, send either [NSSpeechSynthesizer.StartSpeakingString] or
// [NSSpeechSynthesizer.StartSpeakingStringToURL] to the instance. The former generates speech
// through the system’s default sound output device; the latter saves the
// generated speech to a file. If you wish to be notified when the current
// speech concludes, set the [NSSpeechSynthesizer.Delegate] property and implement the delegate
// method [SpeechSynthesizerDidFinishSpeaking].
// 
// Speech synthesis is just one of the macOS speech technologies. The speech
// recognizer technology allows applications to “listen to” text spoken in
// U.S. English; the [NSSpeechRecognizer] class is the Cocoa interface to this
// technology. Both technologies provide benefits for all users, and are
// particularly useful to those users who have difficulties seeing the screen
// or using the mouse and keyboard.
// 
// # Speech Feedback Window
// 
// The speech feedback window ([Figure 1]) displays the text recognized from
// the user’s speech and the text from which an [NSSpeechSynthesizer] object
// synthesizes speech. Using the feedback window makes spoken exchange more
// natural and helps the user understand the synthesized speech.
// 
// [media-1965715]
// 
// For example, your application may use an [NSSpeechRecognizer] object to
// listen for the command “Play some music.” When it recognizes this
// command, your application might then respond by speaking “Which
// artist?” using a speech synthesizer.
// 
// When [NSSpeechSynthesizer.UsesFeedbackWindow] is [true], the speech synthesizer uses the
// feedback window if its visible, which the user specifies in System
// Preferences > Speech.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Customizing the Speech Synthesizer Behavior
//
//   - [NSSpeechSynthesizer.Delegate]: The synthesizer’s delegate.
//   - [NSSpeechSynthesizer.SetDelegate]
//
// # Configuring Speech Synthesizers
//
//   - [NSSpeechSynthesizer.UsesFeedbackWindow]: Indicates whether the receiver uses the speech feedback window.
//   - [NSSpeechSynthesizer.SetUsesFeedbackWindow]
//   - [NSSpeechSynthesizer.Rate]: The synthesizer’s speaking rate (words per minute).
//   - [NSSpeechSynthesizer.SetRate]
//   - [NSSpeechSynthesizer.Volume]: The synthesizer’s speaking volume.
//   - [NSSpeechSynthesizer.SetVolume]
//
// # Synthesizing Speech
//
//   - [NSSpeechSynthesizer.Speaking]: Indicates whether the receiver is currently generating synthesized speech.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer
type NSSpeechSynthesizer struct {
	objectivec.Object
}

// NSSpeechSynthesizerFromID constructs a [NSSpeechSynthesizer] from an objc.ID.
//
// The Cocoa interface to speech synthesis in macOS.
func NSSpeechSynthesizerFromID(id objc.ID) NSSpeechSynthesizer {
	return NSSpeechSynthesizer{objectivec.Object{ID: id}}
}
// NOTE: NSSpeechSynthesizer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSpeechSynthesizer] class.
//
// # Customizing the Speech Synthesizer Behavior
//
//   - [INSSpeechSynthesizer.Delegate]: The synthesizer’s delegate.
//   - [INSSpeechSynthesizer.SetDelegate]
//
// # Configuring Speech Synthesizers
//
//   - [INSSpeechSynthesizer.UsesFeedbackWindow]: Indicates whether the receiver uses the speech feedback window.
//   - [INSSpeechSynthesizer.SetUsesFeedbackWindow]
//   - [INSSpeechSynthesizer.Rate]: The synthesizer’s speaking rate (words per minute).
//   - [INSSpeechSynthesizer.SetRate]
//   - [INSSpeechSynthesizer.Volume]: The synthesizer’s speaking volume.
//   - [INSSpeechSynthesizer.SetVolume]
//
// # Synthesizing Speech
//
//   - [INSSpeechSynthesizer.Speaking]: Indicates whether the receiver is currently generating synthesized speech.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer
type INSSpeechSynthesizer interface {
	objectivec.IObject

	// Topic: Customizing the Speech Synthesizer Behavior

	// The synthesizer’s delegate.
	Delegate() NSSpeechSynthesizerDelegate
	SetDelegate(value NSSpeechSynthesizerDelegate)

	// Topic: Configuring Speech Synthesizers

	// Indicates whether the receiver uses the speech feedback window.
	UsesFeedbackWindow() bool
	SetUsesFeedbackWindow(value bool)
	// The synthesizer’s speaking rate (words per minute).
	Rate() float32
	SetRate(value float32)
	// The synthesizer’s speaking volume.
	Volume() float32
	SetVolume(value float32)

	// Topic: Synthesizing Speech

	// Indicates whether the receiver is currently generating synthesized speech.
	Speaking() bool

	// The perceived gender of the voice. The supported values are listed in 
	Gender() NSVoiceAttributeKey
}

// Init initializes the instance.
func (s NSSpeechSynthesizer) Init() NSSpeechSynthesizer {
	rv := objc.Send[NSSpeechSynthesizer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSpeechSynthesizer) Autorelease() NSSpeechSynthesizer {
	rv := objc.Send[NSSpeechSynthesizer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSpeechSynthesizer creates a new NSSpeechSynthesizer instance.
func NewNSSpeechSynthesizer() NSSpeechSynthesizer {
	class := getNSSpeechSynthesizerClass()
	rv := objc.Send[NSSpeechSynthesizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes the receiver with a voice.
//
// voice: Identifier of the voice to set as the current voice. When `nil`, the
// default voice is used. Passing in a specific voice means the initial
// speaking rate is determined by the synthesizer’s default speaking rate;
// passing `nil` means the speaking rate is automatically set to the rate the
// user specifies in Speech preferences.
//
// # Return Value
// 
// Initialized speech synthesizer or `nil` when the voice identified by
// `voiceIdentifier` is not available or when there’s an allocation error.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/init(voice:)
func NewSpeechSynthesizerWithVoice(voice NSSpeechSynthesizerVoiceName) NSSpeechSynthesizer {
	instance := getNSSpeechSynthesizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVoice:"), objc.String(string(voice)))
	return NSSpeechSynthesizerFromID(rv)
}

// The synthesizer’s delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/delegate
func (s NSSpeechSynthesizer) Delegate() NSSpeechSynthesizerDelegate {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return NSSpeechSynthesizerDelegateObjectFromID(rv)
}
func (s NSSpeechSynthesizer) SetDelegate(value NSSpeechSynthesizerDelegate) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}

// Indicates whether the receiver uses the speech feedback window.
//
// # Discussion
// 
// [true] when the receiver uses the speech feedback window, [false]
// otherwise.
// 
// See the class description for details on the [UsesFeedbackWindow]
// attribute.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/usesFeedbackWindow
func (s NSSpeechSynthesizer) UsesFeedbackWindow() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("usesFeedbackWindow"))
	return rv
}
func (s NSSpeechSynthesizer) SetUsesFeedbackWindow(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setUsesFeedbackWindow:"), value)
}

// The synthesizer’s speaking rate (words per minute).
//
// # Discussion
// 
// The range of supported rates is not predefined by the Speech Synthesis
// framework; but the synthesizer may only respond to a limited range of
// speech rates. Average human speech occurs at a rate of 180 to 220 words per
// minute.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/rate
func (s NSSpeechSynthesizer) Rate() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("rate"))
	return rv
}
func (s NSSpeechSynthesizer) SetRate(value float32) {
	objc.Send[struct{}](s.ID, objc.Sel("setRate:"), value)
}

// The synthesizer’s speaking volume.
//
// # Discussion
// 
// Volumes are expressed in floating-point units ranging from 0.0 through 1.0.
// A value of 0.0 corresponds to silence, and a value of 1.0 corresponds to
// the maximum possible volume. Volume units lie on a scale that is linear
// with amplitude or voltage. A doubling of perceived loudness corresponds to
// a doubling of the volume. Setting a value outside this range is undefined.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/volume
func (s NSSpeechSynthesizer) Volume() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("volume"))
	return rv
}
func (s NSSpeechSynthesizer) SetVolume(value float32) {
	objc.Send[struct{}](s.ID, objc.Sel("setVolume:"), value)
}

// Indicates whether the receiver is currently generating synthesized speech.
//
// # Discussion
// 
// [true] when the receiver is generating synthesized speech, [false]
// otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/isSpeaking
func (s NSSpeechSynthesizer) Speaking() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isSpeaking"))
	return rv
}

// The perceived gender of the voice. The supported values are listed in
//
// See: https://developer.apple.com/documentation/appkit/nsspeechsynthesizer/voiceattributekey/gender
func (s NSSpeechSynthesizer) Gender() NSVoiceAttributeKey {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("NSVoiceGender"))
	return NSVoiceAttributeKey(foundation.NSStringFromID(rv).String())
}

// Provides the identifiers of the voices available on the system.
//
// # Return Value
// 
// Array of strings representing the identifiers of each voice available on
// the system.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/availableVoices
func (_NSSpeechSynthesizerClass NSSpeechSynthesizerClass) AvailableVoices() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSSpeechSynthesizerClass.class), objc.Sel("availableVoices"))
	return objc.ConvertSliceToStrings(rv)
}

// Provides the identifier of the default voice.
//
// # Return Value
// 
// Identifier of the default voice.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/defaultVoice
func (_NSSpeechSynthesizerClass NSSpeechSynthesizerClass) DefaultVoice() NSSpeechSynthesizerVoiceName {
	rv := objc.Send[objc.ID](objc.ID(_NSSpeechSynthesizerClass.class), objc.Sel("defaultVoice"))
	return NSSpeechSynthesizerVoiceName(foundation.NSStringFromID(rv).String())
}

// A Boolean value indicating whether any application is currently speaking
// through the sound output device.
//
// # Discussion
// 
// This property is [true] when another application is producing speech
// through the sound output device. You usually invoke this method to prevent
// your application from speaking over speech being generated by another
// application or system component.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/isAnyApplicationSpeaking
func (_NSSpeechSynthesizerClass NSSpeechSynthesizerClass) AnyApplicationSpeaking() bool {
	rv := objc.Send[bool](objc.ID(_NSSpeechSynthesizerClass.class), objc.Sel("isAnyApplicationSpeaking"))
	return rv
}

