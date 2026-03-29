// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSpeechSynthesisProviderAudioUnit] class.
var (
	_AVSpeechSynthesisProviderAudioUnitClass     AVSpeechSynthesisProviderAudioUnitClass
	_AVSpeechSynthesisProviderAudioUnitClassOnce sync.Once
)

func getAVSpeechSynthesisProviderAudioUnitClass() AVSpeechSynthesisProviderAudioUnitClass {
	_AVSpeechSynthesisProviderAudioUnitClassOnce.Do(func() {
		_AVSpeechSynthesisProviderAudioUnitClass = AVSpeechSynthesisProviderAudioUnitClass{class: objc.GetClass("AVSpeechSynthesisProviderAudioUnit")}
	})
	return _AVSpeechSynthesisProviderAudioUnitClass
}

// GetAVSpeechSynthesisProviderAudioUnitClass returns the class object for AVSpeechSynthesisProviderAudioUnit.
func GetAVSpeechSynthesisProviderAudioUnitClass() AVSpeechSynthesisProviderAudioUnitClass {
	return getAVSpeechSynthesisProviderAudioUnitClass()
}

type AVSpeechSynthesisProviderAudioUnitClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSpeechSynthesisProviderAudioUnitClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSpeechSynthesisProviderAudioUnitClass) Alloc() AVSpeechSynthesisProviderAudioUnit {
	rv := objc.Send[AVSpeechSynthesisProviderAudioUnit](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that generates speech from text.
//
// # Overview
// 
// Use a speech synthesizer audio unit to generate audio buffers that contain
// speech for a given voice and speech markup. The audio unit receives an
// [AVSpeechSynthesisProviderRequest] as input, and extracts audio buffers
// through the render block.
// 
// Use [AVSpeechSynthesisProviderAudioUnit.SpeechSynthesisOutputMetadataBlock] to provide metadata as an array of
// [AVSpeechSynthesisMarker].
// 
// The system scans and loads voices for audio unit extensions of this type,
// and the voices it provides are available for use in [AVSpeechSynthesizer]
// and accessibility technologies like VoiceOver and Speak Screen.
//
// # Rendering speech
//
//   - [AVSpeechSynthesisProviderAudioUnit.SynthesizeSpeechRequest]: Sets the text to synthesize and the voice to use.
//
// # Supplying metadata
//
//   - [AVSpeechSynthesisProviderAudioUnit.SpeechSynthesisOutputMetadataBlock]: A block that subclasses use to send marker information to the host.
//   - [AVSpeechSynthesisProviderAudioUnit.SetSpeechSynthesisOutputMetadataBlock]
//
// # Getting and setting voices
//
//   - [AVSpeechSynthesisProviderAudioUnit.SpeechVoices]: A list of voices the audio unit provides to the system.
//   - [AVSpeechSynthesisProviderAudioUnit.SetSpeechVoices]
//
// # Cancelling a request
//
//   - [AVSpeechSynthesisProviderAudioUnit.CancelSpeechRequest]: Informs the audio unit to discard the speech request.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderAudioUnit
type AVSpeechSynthesisProviderAudioUnit struct {
	objectivec.Object
}

// AVSpeechSynthesisProviderAudioUnitFromID constructs a [AVSpeechSynthesisProviderAudioUnit] from an objc.ID.
//
// An object that generates speech from text.
func AVSpeechSynthesisProviderAudioUnitFromID(id objc.ID) AVSpeechSynthesisProviderAudioUnit {
	return AVSpeechSynthesisProviderAudioUnit{objectivec.Object{ID: id}}
}
// NOTE: AVSpeechSynthesisProviderAudioUnit adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVSpeechSynthesisProviderAudioUnit] class.
//
// # Rendering speech
//
//   - [IAVSpeechSynthesisProviderAudioUnit.SynthesizeSpeechRequest]: Sets the text to synthesize and the voice to use.
//
// # Supplying metadata
//
//   - [IAVSpeechSynthesisProviderAudioUnit.SpeechSynthesisOutputMetadataBlock]: A block that subclasses use to send marker information to the host.
//   - [IAVSpeechSynthesisProviderAudioUnit.SetSpeechSynthesisOutputMetadataBlock]
//
// # Getting and setting voices
//
//   - [IAVSpeechSynthesisProviderAudioUnit.SpeechVoices]: A list of voices the audio unit provides to the system.
//   - [IAVSpeechSynthesisProviderAudioUnit.SetSpeechVoices]
//
// # Cancelling a request
//
//   - [IAVSpeechSynthesisProviderAudioUnit.CancelSpeechRequest]: Informs the audio unit to discard the speech request.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderAudioUnit
type IAVSpeechSynthesisProviderAudioUnit interface {
	objectivec.IObject

	// Topic: Rendering speech

	// Sets the text to synthesize and the voice to use.
	SynthesizeSpeechRequest(speechRequest IAVSpeechSynthesisProviderRequest)

	// Topic: Supplying metadata

	// A block that subclasses use to send marker information to the host.
	SpeechSynthesisOutputMetadataBlock() AVSpeechSynthesisProviderOutputBlock
	SetSpeechSynthesisOutputMetadataBlock(value AVSpeechSynthesisProviderOutputBlock)

	// Topic: Getting and setting voices

	// A list of voices the audio unit provides to the system.
	SpeechVoices() []AVSpeechSynthesisProviderVoice
	SetSpeechVoices(value []AVSpeechSynthesisProviderVoice)

	// Topic: Cancelling a request

	// Informs the audio unit to discard the speech request.
	CancelSpeechRequest()
}

// Init initializes the instance.
func (s AVSpeechSynthesisProviderAudioUnit) Init() AVSpeechSynthesisProviderAudioUnit {
	rv := objc.Send[AVSpeechSynthesisProviderAudioUnit](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSpeechSynthesisProviderAudioUnit) Autorelease() AVSpeechSynthesisProviderAudioUnit {
	rv := objc.Send[AVSpeechSynthesisProviderAudioUnit](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSpeechSynthesisProviderAudioUnit creates a new AVSpeechSynthesisProviderAudioUnit instance.
func NewAVSpeechSynthesisProviderAudioUnit() AVSpeechSynthesisProviderAudioUnit {
	class := getAVSpeechSynthesisProviderAudioUnitClass()
	rv := objc.Send[AVSpeechSynthesisProviderAudioUnit](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Sets the text to synthesize and the voice to use.
//
// speechRequest: A speech request to synthesize.
//
// # Discussion
// 
// When the synthesizer finishes generating audio buffers for the speech
// request, use [AUInternalRenderBlock] to report
// [offlineUnitRenderAction_Complete].
//
// [AUInternalRenderBlock]: https://developer.apple.com/documentation/AudioToolbox/AUInternalRenderBlock
// [offlineUnitRenderAction_Complete]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRenderActionFlags/offlineUnitRenderAction_Complete
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderAudioUnit/synthesizeSpeechRequest(_:)
func (s AVSpeechSynthesisProviderAudioUnit) SynthesizeSpeechRequest(speechRequest IAVSpeechSynthesisProviderRequest) {
	objc.Send[objc.ID](s.ID, objc.Sel("synthesizeSpeechRequest:"), speechRequest)
}
// Informs the audio unit to discard the speech request.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderAudioUnit/cancelSpeechRequest()
func (s AVSpeechSynthesisProviderAudioUnit) CancelSpeechRequest() {
	objc.Send[objc.ID](s.ID, objc.Sel("cancelSpeechRequest"))
}

// A block that subclasses use to send marker information to the host.
//
// # Discussion
// 
// A host sets this block to retrieve metadata for a request.
// 
// A synthesizer calls this method when it produces data relevant to the audio
// buffers it’s sending back to a host. In some cases, the system may delay
// speech output until it delivers these markers. For example, word
// highlighting depends on marker data from synthesizers to time what word to
// highlight. The array of markers can reference audio buffers that the system
// delivers at a later time.
// 
// There may be cases where a subclass doesn’t have marker data until it
// completes extra audio processing. If marker data changes, this block
// replaces that audio buffer range’s marker data.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderAudioUnit/speechSynthesisOutputMetadataBlock
func (s AVSpeechSynthesisProviderAudioUnit) SpeechSynthesisOutputMetadataBlock() AVSpeechSynthesisProviderOutputBlock {
	rv := objc.Send[AVSpeechSynthesisProviderOutputBlock](s.ID, objc.Sel("speechSynthesisOutputMetadataBlock"))
	return AVSpeechSynthesisProviderOutputBlock(rv)
}
func (s AVSpeechSynthesisProviderAudioUnit) SetSpeechSynthesisOutputMetadataBlock(value AVSpeechSynthesisProviderOutputBlock) {
	objc.Send[struct{}](s.ID, objc.Sel("setSpeechSynthesisOutputMetadataBlock:"), value)
}
// A list of voices the audio unit provides to the system.
//
// # Discussion
// 
// The list of voices that a user selects through Settings. Speech synthesizer
// audio unit extensions must provide this list. Override the getter to
// perform complex fetches that provide a dynamic list of voices.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderAudioUnit/speechVoices
func (s AVSpeechSynthesisProviderAudioUnit) SpeechVoices() []AVSpeechSynthesisProviderVoice {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("speechVoices"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVSpeechSynthesisProviderVoice {
		return AVSpeechSynthesisProviderVoiceFromID(id)
	})
}
func (s AVSpeechSynthesisProviderAudioUnit) SetSpeechVoices(value []AVSpeechSynthesisProviderVoice) {
	objc.Send[struct{}](s.ID, objc.Sel("setSpeechVoices:"), objectivec.IObjectSliceToNSArray(value))
}

