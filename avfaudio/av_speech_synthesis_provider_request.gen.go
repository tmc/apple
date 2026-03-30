// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSpeechSynthesisProviderRequest] class.
var (
	_AVSpeechSynthesisProviderRequestClass     AVSpeechSynthesisProviderRequestClass
	_AVSpeechSynthesisProviderRequestClassOnce sync.Once
)

func getAVSpeechSynthesisProviderRequestClass() AVSpeechSynthesisProviderRequestClass {
	_AVSpeechSynthesisProviderRequestClassOnce.Do(func() {
		_AVSpeechSynthesisProviderRequestClass = AVSpeechSynthesisProviderRequestClass{class: objc.GetClass("AVSpeechSynthesisProviderRequest")}
	})
	return _AVSpeechSynthesisProviderRequestClass
}

// GetAVSpeechSynthesisProviderRequestClass returns the class object for AVSpeechSynthesisProviderRequest.
func GetAVSpeechSynthesisProviderRequestClass() AVSpeechSynthesisProviderRequestClass {
	return getAVSpeechSynthesisProviderRequestClass()
}

type AVSpeechSynthesisProviderRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSpeechSynthesisProviderRequestClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSpeechSynthesisProviderRequestClass) Alloc() AVSpeechSynthesisProviderRequest {
	rv := objc.Send[AVSpeechSynthesisProviderRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the text to synthesize and the voice to use.
//
// # Creating a request
//
//   - [AVSpeechSynthesisProviderRequest.InitWithSSMLRepresentationVoice]: Creates a request with a voice and a description.
//
// # Inspecting a request
//
//   - [AVSpeechSynthesisProviderRequest.SsmlRepresentation]: The description of the text to synthesize.
//   - [AVSpeechSynthesisProviderRequest.Voice]: The voice to use in the speech request.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest
type AVSpeechSynthesisProviderRequest struct {
	objectivec.Object
}

// AVSpeechSynthesisProviderRequestFromID constructs a [AVSpeechSynthesisProviderRequest] from an objc.ID.
//
// An object that represents the text to synthesize and the voice to use.
func AVSpeechSynthesisProviderRequestFromID(id objc.ID) AVSpeechSynthesisProviderRequest {
	return AVSpeechSynthesisProviderRequest{objectivec.Object{ID: id}}
}

// NOTE: AVSpeechSynthesisProviderRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVSpeechSynthesisProviderRequest] class.
//
// # Creating a request
//
//   - [IAVSpeechSynthesisProviderRequest.InitWithSSMLRepresentationVoice]: Creates a request with a voice and a description.
//
// # Inspecting a request
//
//   - [IAVSpeechSynthesisProviderRequest.SsmlRepresentation]: The description of the text to synthesize.
//   - [IAVSpeechSynthesisProviderRequest.Voice]: The voice to use in the speech request.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest
type IAVSpeechSynthesisProviderRequest interface {
	objectivec.IObject

	// Topic: Creating a request

	// Creates a request with a voice and a description.
	InitWithSSMLRepresentationVoice(text string, voice IAVSpeechSynthesisProviderVoice) AVSpeechSynthesisProviderRequest

	// Topic: Inspecting a request

	// The description of the text to synthesize.
	SsmlRepresentation() string
	// The voice to use in the speech request.
	Voice() IAVSpeechSynthesisProviderVoice

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s AVSpeechSynthesisProviderRequest) Init() AVSpeechSynthesisProviderRequest {
	rv := objc.Send[AVSpeechSynthesisProviderRequest](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSpeechSynthesisProviderRequest) Autorelease() AVSpeechSynthesisProviderRequest {
	rv := objc.Send[AVSpeechSynthesisProviderRequest](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSpeechSynthesisProviderRequest creates a new AVSpeechSynthesisProviderRequest instance.
func NewAVSpeechSynthesisProviderRequest() AVSpeechSynthesisProviderRequest {
	class := getAVSpeechSynthesisProviderRequestClass()
	rv := objc.Send[AVSpeechSynthesisProviderRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a request with a voice and a description.
//
// text: The description of the text to synthesize.
//
// voice: The voice to use in the speech request.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest/init(ssmlRepresentation:voice:)
func NewSpeechSynthesisProviderRequestWithSSMLRepresentationVoice(text string, voice IAVSpeechSynthesisProviderVoice) AVSpeechSynthesisProviderRequest {
	instance := getAVSpeechSynthesisProviderRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSSMLRepresentation:voice:"), objc.String(text), voice)
	return AVSpeechSynthesisProviderRequestFromID(rv)
}

// Creates a request with a voice and a description.
//
// text: The description of the text to synthesize.
//
// voice: The voice to use in the speech request.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest/init(ssmlRepresentation:voice:)
func (s AVSpeechSynthesisProviderRequest) InitWithSSMLRepresentationVoice(text string, voice IAVSpeechSynthesisProviderVoice) AVSpeechSynthesisProviderRequest {
	rv := objc.Send[AVSpeechSynthesisProviderRequest](s.ID, objc.Sel("initWithSSMLRepresentation:voice:"), objc.String(text), voice)
	return rv
}
func (s AVSpeechSynthesisProviderRequest) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The description of the text to synthesize.
//
// # Discussion
//
// The Speech Synthesis Markup Language describes the speech synthesis
// attributes for the customization of pitch, rate, intonation, and more.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest/ssmlRepresentation
func (s AVSpeechSynthesisProviderRequest) SsmlRepresentation() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("ssmlRepresentation"))
	return foundation.NSStringFromID(rv).String()
}

// The voice to use in the speech request.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest/voice
func (s AVSpeechSynthesisProviderRequest) Voice() IAVSpeechSynthesisProviderVoice {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("voice"))
	return AVSpeechSynthesisProviderVoiceFromID(objc.ID(rv))
}
