// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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

//
// # Methods
//
//   - [AVSpeechSynthesisProviderRequest.JobIdentifier]
//   - [AVSpeechSynthesisProviderRequest.SetJobIdentifier]
//   - [AVSpeechSynthesisProviderRequest.InitWithCoder]
//   - [AVSpeechSynthesisProviderRequest.SsmlRepresentation]
//   - [AVSpeechSynthesisProviderRequest.SetSsmlRepresentation]
//   - [AVSpeechSynthesisProviderRequest.Voice]
//   - [AVSpeechSynthesisProviderRequest.SetVoice]
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest
type AVSpeechSynthesisProviderRequest struct {
	objectivec.Object
}

// AVSpeechSynthesisProviderRequestFromID constructs a [AVSpeechSynthesisProviderRequest] from an objc.ID.
func AVSpeechSynthesisProviderRequestFromID(id objc.ID) AVSpeechSynthesisProviderRequest {
	return AVSpeechSynthesisProviderRequest{objectivec.Object{ID: id}}
}
// Ensure AVSpeechSynthesisProviderRequest implements IAVSpeechSynthesisProviderRequest.
var _ IAVSpeechSynthesisProviderRequest = AVSpeechSynthesisProviderRequest{}

// An interface definition for the [AVSpeechSynthesisProviderRequest] class.
//
// # Methods
//
//   - [IAVSpeechSynthesisProviderRequest.JobIdentifier]
//   - [IAVSpeechSynthesisProviderRequest.SetJobIdentifier]
//   - [IAVSpeechSynthesisProviderRequest.InitWithCoder]
//   - [IAVSpeechSynthesisProviderRequest.SsmlRepresentation]
//   - [IAVSpeechSynthesisProviderRequest.SetSsmlRepresentation]
//   - [IAVSpeechSynthesisProviderRequest.Voice]
//   - [IAVSpeechSynthesisProviderRequest.SetVoice]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest
type IAVSpeechSynthesisProviderRequest interface {
	objectivec.IObject

	// Topic: Methods

	JobIdentifier() string
	SetJobIdentifier(value string)
	InitWithCoder(coder foundation.INSCoder) AVSpeechSynthesisProviderRequest
	SsmlRepresentation() string
	SetSsmlRepresentation(value string)
	Voice() IAVSpeechSynthesisProviderVoice
	SetVoice(value IAVSpeechSynthesisProviderVoice)
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

//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest/initWithCoder:
func NewSpeechSynthesisProviderRequestWithCoder(coder objectivec.IObject) AVSpeechSynthesisProviderRequest {
	instance := getAVSpeechSynthesisProviderRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AVSpeechSynthesisProviderRequestFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest/initWithCoder:
func (s AVSpeechSynthesisProviderRequest) InitWithCoder(coder foundation.INSCoder) AVSpeechSynthesisProviderRequest {
	rv := objc.Send[AVSpeechSynthesisProviderRequest](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest/supportsSecureCoding
func (_AVSpeechSynthesisProviderRequestClass AVSpeechSynthesisProviderRequestClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_AVSpeechSynthesisProviderRequestClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest/jobIdentifier
func (s AVSpeechSynthesisProviderRequest) JobIdentifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("jobIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (s AVSpeechSynthesisProviderRequest) SetJobIdentifier(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setJobIdentifier:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest/ssmlRepresentation
func (s AVSpeechSynthesisProviderRequest) SsmlRepresentation() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("ssmlRepresentation"))
	return foundation.NSStringFromID(rv).String()
}
func (s AVSpeechSynthesisProviderRequest) SetSsmlRepresentation(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setSsmlRepresentation:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderRequest/voice
func (s AVSpeechSynthesisProviderRequest) Voice() IAVSpeechSynthesisProviderVoice {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("voice"))
	return AVSpeechSynthesisProviderVoiceFromID(objc.ID(rv))
}
func (s AVSpeechSynthesisProviderRequest) SetVoice(value IAVSpeechSynthesisProviderVoice) {
	objc.Send[struct{}](s.ID, objc.Sel("setVoice:"), value)
}

