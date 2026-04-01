// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSpeechSynthesizer] class.
var (
	_AVSpeechSynthesizerClass     AVSpeechSynthesizerClass
	_AVSpeechSynthesizerClassOnce sync.Once
)

func getAVSpeechSynthesizerClass() AVSpeechSynthesizerClass {
	_AVSpeechSynthesizerClassOnce.Do(func() {
		_AVSpeechSynthesizerClass = AVSpeechSynthesizerClass{class: objc.GetClass("AVSpeechSynthesizer")}
	})
	return _AVSpeechSynthesizerClass
}

// GetAVSpeechSynthesizerClass returns the class object for AVSpeechSynthesizer.
func GetAVSpeechSynthesizerClass() AVSpeechSynthesizerClass {
	return getAVSpeechSynthesizerClass()
}

type AVSpeechSynthesizerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSpeechSynthesizerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSpeechSynthesizerClass) Alloc() AVSpeechSynthesizer {
	rv := objc.Send[AVSpeechSynthesizer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that produces synthesized speech from text utterances and enables
// monitoring or controlling of ongoing speech.
//
// # Overview
//
// To speak some text, create an [AVSpeechUtterance] instance that contains
// the text and pass it to [AVSpeechSynthesizer.SpeakUtterance] on a speech synthesizer instance.
// You can optionally also retrieve an [AVSpeechSynthesisVoice] and set it on
// the utterance’s [AVSpeechSynthesizer.Voice] property to have the speech synthesizer use that
// voice when speaking the utterance’s text.
//
// The speech synthesizer maintains a queue of utterances that it speaks. If
// the synthesizer isn’t speaking, calling [AVSpeechSynthesizer.SpeakUtterance] begins speaking
// that utterance either immediately or after pausing for its
// [AVSpeechSynthesizer.PreUtteranceDelay], if necessary. If the synthesizer is speaking, the
// synthesizer adds utterances to a queue and speaks them in the order it
// receives them.
//
// After speech begins, you can use the synthesizer object to pause or stop
// speech. After pausing, you can resume the speech from its paused point or
// stop the speech entirely and remove all remaining utterances in the queue.
//
// You can monitor the speech synthesizer by examining its [AVSpeechSynthesizer.Speaking] and
// [AVSpeechSynthesizer.Paused] properties, or by setting a delegate that conforms to
// [AVSpeechSynthesizerDelegate]. The delegate receives significant events as
// they occur during speech synthesis.
//
// An [AVSpeechSynthesizer] also controls the route where the speech plays.
// For more information, see Directing speech output.
//
// # Controlling speech
//
//   - [AVSpeechSynthesizer.SpeakUtterance]: Adds the utterance you specify to the speech synthesizer’s queue.
//   - [AVSpeechSynthesizer.ContinueSpeaking]: Resumes speech from its paused point.
//   - [AVSpeechSynthesizer.PauseSpeakingAtBoundary]: Pauses speech at the boundary you specify.
//   - [AVSpeechSynthesizer.StopSpeakingAtBoundary]: Stops speech at the boundary you specify.
//
// # Inspecting a speech synthesizer
//
//   - [AVSpeechSynthesizer.Speaking]: A Boolean value that indicates whether the speech synthesizer is speaking or is in a paused state and has utterances to speak.
//   - [AVSpeechSynthesizer.Paused]: A Boolean value that indicates whether a speech synthesizer is in a paused state.
//
// # Managing the delegate
//
//   - [AVSpeechSynthesizer.Delegate]: The delegate object for the speech synthesizer.
//   - [AVSpeechSynthesizer.SetDelegate]
//
// # Directing speech output
//
//   - [AVSpeechSynthesizer.WriteUtteranceToBufferCallback]: Generates speech for the utterance and invokes the callback with the audio buffer.
//   - [AVSpeechSynthesizer.WriteUtteranceToBufferCallbackToMarkerCallback]: Generates audio buffers and associated metadata for storage or further speech synthesis processing.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer
type AVSpeechSynthesizer struct {
	objectivec.Object
}

// AVSpeechSynthesizerFromID constructs a [AVSpeechSynthesizer] from an objc.ID.
//
// An object that produces synthesized speech from text utterances and enables
// monitoring or controlling of ongoing speech.
func AVSpeechSynthesizerFromID(id objc.ID) AVSpeechSynthesizer {
	return AVSpeechSynthesizer{objectivec.Object{ID: id}}
}

// NOTE: AVSpeechSynthesizer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVSpeechSynthesizer] class.
//
// # Controlling speech
//
//   - [IAVSpeechSynthesizer.SpeakUtterance]: Adds the utterance you specify to the speech synthesizer’s queue.
//   - [IAVSpeechSynthesizer.ContinueSpeaking]: Resumes speech from its paused point.
//   - [IAVSpeechSynthesizer.PauseSpeakingAtBoundary]: Pauses speech at the boundary you specify.
//   - [IAVSpeechSynthesizer.StopSpeakingAtBoundary]: Stops speech at the boundary you specify.
//
// # Inspecting a speech synthesizer
//
//   - [IAVSpeechSynthesizer.Speaking]: A Boolean value that indicates whether the speech synthesizer is speaking or is in a paused state and has utterances to speak.
//   - [IAVSpeechSynthesizer.Paused]: A Boolean value that indicates whether a speech synthesizer is in a paused state.
//
// # Managing the delegate
//
//   - [IAVSpeechSynthesizer.Delegate]: The delegate object for the speech synthesizer.
//   - [IAVSpeechSynthesizer.SetDelegate]
//
// # Directing speech output
//
//   - [IAVSpeechSynthesizer.WriteUtteranceToBufferCallback]: Generates speech for the utterance and invokes the callback with the audio buffer.
//   - [IAVSpeechSynthesizer.WriteUtteranceToBufferCallbackToMarkerCallback]: Generates audio buffers and associated metadata for storage or further speech synthesis processing.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer
type IAVSpeechSynthesizer interface {
	objectivec.IObject

	// Topic: Controlling speech

	// Adds the utterance you specify to the speech synthesizer’s queue.
	SpeakUtterance(utterance IAVSpeechUtterance)
	// Resumes speech from its paused point.
	ContinueSpeaking() bool
	// Pauses speech at the boundary you specify.
	PauseSpeakingAtBoundary(boundary AVSpeechBoundary) bool
	// Stops speech at the boundary you specify.
	StopSpeakingAtBoundary(boundary AVSpeechBoundary) bool

	// Topic: Inspecting a speech synthesizer

	// A Boolean value that indicates whether the speech synthesizer is speaking or is in a paused state and has utterances to speak.
	Speaking() bool
	// A Boolean value that indicates whether a speech synthesizer is in a paused state.
	Paused() bool

	// Topic: Managing the delegate

	// The delegate object for the speech synthesizer.
	Delegate() AVSpeechSynthesizerDelegate
	SetDelegate(value AVSpeechSynthesizerDelegate)

	// Topic: Directing speech output

	// Generates speech for the utterance and invokes the callback with the audio buffer.
	WriteUtteranceToBufferCallback(utterance IAVSpeechUtterance, bufferCallback AVSpeechSynthesizerBufferCallback)
	// Generates audio buffers and associated metadata for storage or further speech synthesis processing.
	WriteUtteranceToBufferCallbackToMarkerCallback(utterance IAVSpeechUtterance, bufferCallback AVSpeechSynthesizerBufferCallback, markerCallback AVSpeechSynthesizerMarkerCallback)

	// The amount of time the speech synthesizer pauses before speaking the utterance.
	PreUtteranceDelay() float64
	SetPreUtteranceDelay(value float64)
	// The voice the speech synthesizer uses when speaking the utterance.
	Voice() IAVSpeechSynthesisVoice
	SetVoice(value IAVSpeechSynthesisVoice)
}

// Init initializes the instance.
func (s AVSpeechSynthesizer) Init() AVSpeechSynthesizer {
	rv := objc.Send[AVSpeechSynthesizer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSpeechSynthesizer) Autorelease() AVSpeechSynthesizer {
	rv := objc.Send[AVSpeechSynthesizer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSpeechSynthesizer creates a new AVSpeechSynthesizer instance.
func NewAVSpeechSynthesizer() AVSpeechSynthesizer {
	class := getAVSpeechSynthesizerClass()
	rv := objc.Send[AVSpeechSynthesizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Adds the utterance you specify to the speech synthesizer’s queue.
//
// utterance: An [AVSpeechUtterance] instance that contains text to speak.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/speak(_:)
func (s AVSpeechSynthesizer) SpeakUtterance(utterance IAVSpeechUtterance) {
	objc.Send[objc.ID](s.ID, objc.Sel("speakUtterance:"), utterance)
}

// Resumes speech from its paused point.
//
// # Return Value
//
// true if speech resumes; otherwise, false.
//
// # Discussion
//
// This method only has an effect if the speech synthesizer is in a paused
// state.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/continueSpeaking()
func (s AVSpeechSynthesizer) ContinueSpeaking() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("continueSpeaking"))
	return rv
}

// Pauses speech at the boundary you specify.
//
// boundary: An enumeration that describes whether to pause speech immediately or only
// after the synthesizer finishes speaking the current word.
//
// # Return Value
//
// true if speech pauses; otherwise, false.
//
// # Discussion
//
// The `boundary` parameter also affects how the speech synthesizer resumes
// speaking text after a pause and call to [ContinueSpeaking]. If the boundary
// is [AVSpeechBoundaryImmediate], speech resumes from the exact point where
// it pauses, even if that point occurs in the middle of speaking a word. If
// the boundary is [AVSpeechBoundaryWord], speech resumes from the word that
// follows the last spoken word where it pauses.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/pauseSpeaking(at:)
func (s AVSpeechSynthesizer) PauseSpeakingAtBoundary(boundary AVSpeechBoundary) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("pauseSpeakingAtBoundary:"), boundary)
	return rv
}

// Stops speech at the boundary you specify.
//
// boundary: An enumeration that describes whether to stop speech immediately or only
// after the synthesizer finishes speaking the current word.
//
// # Return Value
//
// true if speech stops; otherwise, false.
//
// # Discussion
//
// Unlike pausing a speech synthesizer, which can resume after a pause,
// stopping the synthesizer immediately cancels speech and removes all
// unspoken utterances from the synthesizer’s queue.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/stopSpeaking(at:)
func (s AVSpeechSynthesizer) StopSpeakingAtBoundary(boundary AVSpeechBoundary) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("stopSpeakingAtBoundary:"), boundary)
	return rv
}

// Generates speech for the utterance and invokes the callback with the audio
// buffer.
//
// utterance: The utterance for synthesizing speech.
//
// bufferCallback: The system calls this closure with the generated audio buffer.
//
// # Discussion
//
// Call this method to receive audio buffers to store or further process
// synthesized speech.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/write(_:toBufferCallback:)
func (s AVSpeechSynthesizer) WriteUtteranceToBufferCallback(utterance IAVSpeechUtterance, bufferCallback AVSpeechSynthesizerBufferCallback) {
	objc.Send[objc.ID](s.ID, objc.Sel("writeUtterance:toBufferCallback:"), utterance, bufferCallback)
}

// Generates audio buffers and associated metadata for storage or further
// speech synthesis processing.
//
// utterance: A utterance for a synthesizer to speak.
//
// bufferCallback: A callback that the system invokes with the synthesized audio data.
//
// markerCallback: A callback that the system invokes with marker information.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/write(_:toBufferCallback:toMarkerCallback:)
func (s AVSpeechSynthesizer) WriteUtteranceToBufferCallbackToMarkerCallback(utterance IAVSpeechUtterance, bufferCallback AVSpeechSynthesizerBufferCallback, markerCallback AVSpeechSynthesizerMarkerCallback) {
	objc.Send[objc.ID](s.ID, objc.Sel("writeUtterance:toBufferCallback:toMarkerCallback:"), utterance, bufferCallback, markerCallback)
}

// Prompts the user to authorize your app to use personal voices.
//
// handler: A completion handler that the system calls after the user responds to a
// request to authorize use of personal voices, which receives the
// authorization status as an argument.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/requestPersonalVoiceAuthorization(completionHandler:)
func (_AVSpeechSynthesizerClass AVSpeechSynthesizerClass) RequestPersonalVoiceAuthorizationWithCompletionHandler(handler AVSpeechSynthesisPersonalVoiceAuthorizationStatusHandler) {
	_block0, _ := NewAVSpeechSynthesisPersonalVoiceAuthorizationStatusBlock(handler)
	objc.Send[objc.ID](objc.ID(_AVSpeechSynthesizerClass.class), objc.Sel("requestPersonalVoiceAuthorizationWithCompletionHandler:"), _block0)
}

// A Boolean value that indicates whether the speech synthesizer is speaking
// or is in a paused state and has utterances to speak.
//
// # Discussion
//
// If `true`, the synthesizer is speaking or is in a paused state with
// utterances in its queue. If `false`, the synthesizer isn’t speaking and
// it doesn’t have any utterances in its queue.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/isSpeaking
func (s AVSpeechSynthesizer) Speaking() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isSpeaking"))
	return rv
}

// A Boolean value that indicates whether a speech synthesizer is in a paused
// state.
//
// # Discussion
//
// If `true`, the speech synthesizer is in a paused state after beginning to
// speak an utterance; otherwise, `false`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/isPaused
func (s AVSpeechSynthesizer) Paused() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isPaused"))
	return rv
}

// The delegate object for the speech synthesizer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/delegate
func (s AVSpeechSynthesizer) Delegate() AVSpeechSynthesizerDelegate {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return AVSpeechSynthesizerDelegateObjectFromID(rv)
}
func (s AVSpeechSynthesizer) SetDelegate(value AVSpeechSynthesizerDelegate) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}

// The amount of time the speech synthesizer pauses before speaking the
// utterance.
//
// See: https://developer.apple.com/documentation/avfaudio/avspeechutterance/preutterancedelay
func (s AVSpeechSynthesizer) PreUtteranceDelay() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("preUtteranceDelay"))
	return rv
}
func (s AVSpeechSynthesizer) SetPreUtteranceDelay(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setPreUtteranceDelay:"), value)
}

// The voice the speech synthesizer uses when speaking the utterance.
//
// See: https://developer.apple.com/documentation/avfaudio/avspeechutterance/voice
func (s AVSpeechSynthesizer) Voice() IAVSpeechSynthesisVoice {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("voice"))
	return AVSpeechSynthesisVoiceFromID(objc.ID(rv))
}
func (s AVSpeechSynthesizer) SetVoice(value IAVSpeechSynthesisVoice) {
	objc.Send[struct{}](s.ID, objc.Sel("setVoice:"), value)
}

// Your app’s authorization to use personal voices.
//
// # Discussion
//
// The user can grant or deny your app’s request to use personal voices when
// they’re initially prompted, and change the authorization in the Settings
// app. Additionally, the framework denies the request if the device doesn’t
// support using personal voices.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/personalVoiceAuthorizationStatus-swift.type.property
func (_AVSpeechSynthesizerClass AVSpeechSynthesizerClass) PersonalVoiceAuthorizationStatus() AVSpeechSynthesisPersonalVoiceAuthorizationStatus {
	rv := objc.Send[AVSpeechSynthesisPersonalVoiceAuthorizationStatus](objc.ID(_AVSpeechSynthesizerClass.class), objc.Sel("personalVoiceAuthorizationStatus"))
	return AVSpeechSynthesisPersonalVoiceAuthorizationStatus(rv)
}

// A notification that indicates a change in available voices for speech
// synthesis.
//
// See: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizer/availablevoicesdidchangenotification
func (_AVSpeechSynthesizerClass AVSpeechSynthesizerClass) AvailableVoicesDidChangeNotification() foundation.NSString {
	rv := objc.Send[objc.ID](objc.ID(_AVSpeechSynthesizerClass.class), objc.Sel("AVSpeechSynthesisAvailableVoicesDidChangeNotification"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// RequestPersonalVoiceAuthorization is a synchronous wrapper around [AVSpeechSynthesizer.RequestPersonalVoiceAuthorizationWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (sc AVSpeechSynthesizerClass) RequestPersonalVoiceAuthorization(ctx context.Context) (AVSpeechSynthesisPersonalVoiceAuthorizationStatus, error) {
	done := make(chan AVSpeechSynthesisPersonalVoiceAuthorizationStatus, 1)
	sc.RequestPersonalVoiceAuthorizationWithCompletionHandler(func(val AVSpeechSynthesisPersonalVoiceAuthorizationStatus) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return *new(AVSpeechSynthesisPersonalVoiceAuthorizationStatus), ctx.Err()
	}
}
