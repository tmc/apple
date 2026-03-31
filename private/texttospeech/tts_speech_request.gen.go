// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/avfaudio"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSSpeechRequest] class.
var (
	_TTSSpeechRequestClass     TTSSpeechRequestClass
	_TTSSpeechRequestClassOnce sync.Once
)

func getTTSSpeechRequestClass() TTSSpeechRequestClass {
	_TTSSpeechRequestClassOnce.Do(func() {
		_TTSSpeechRequestClass = TTSSpeechRequestClass{class: objc.GetClass("TTSSpeechRequest")}
	})
	return _TTSSpeechRequestClass
}

// GetTTSSpeechRequestClass returns the class object for TTSSpeechRequest.
func GetTTSSpeechRequestClass() TTSSpeechRequestClass {
	return getTTSSpeechRequestClass()
}

type TTSSpeechRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSSpeechRequestClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSSpeechRequestClass) Alloc() TTSSpeechRequest {
	rv := objc.Send[TTSSpeechRequest](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TTSSpeechRequest.AudioDeviceId]
//   - [TTSSpeechRequest.SetAudioDeviceId]
//   - [TTSSpeechRequest.AudioQueueFlags]
//   - [TTSSpeechRequest.SetAudioQueueFlags]
//   - [TTSSpeechRequest.AudioSessionID]
//   - [TTSSpeechRequest.SetAudioSessionID]
//   - [TTSSpeechRequest.AudioSessionIDIsValid]
//   - [TTSSpeechRequest.SetAudioSessionIDIsValid]
//   - [TTSSpeechRequest.Channels]
//   - [TTSSpeechRequest.SetChannels]
//   - [TTSSpeechRequest.ClientContext]
//   - [TTSSpeechRequest.SetClientContext]
//   - [TTSSpeechRequest.DispatchTime]
//   - [TTSSpeechRequest.SetDispatchTime]
//   - [TTSSpeechRequest.EncodeWithCoder]
//   - [TTSSpeechRequest.Gender]
//   - [TTSSpeechRequest.SetGender]
//   - [TTSSpeechRequest.HandledTime]
//   - [TTSSpeechRequest.SetHandledTime]
//   - [TTSSpeechRequest.IgnoreSubstitutions]
//   - [TTSSpeechRequest.SetIgnoreSubstitutions]
//   - [TTSSpeechRequest.JobIdentifier]
//   - [TTSSpeechRequest.SetJobIdentifier]
//   - [TTSSpeechRequest.LanguageCode]
//   - [TTSSpeechRequest.SetLanguageCode]
//   - [TTSSpeechRequest.Latency]
//   - [TTSSpeechRequest.SetLatency]
//   - [TTSSpeechRequest.MaintainsInput]
//   - [TTSSpeechRequest.SetMaintainsInput]
//   - [TTSSpeechRequest.Pitch]
//   - [TTSSpeechRequest.SetPitch]
//   - [TTSSpeechRequest.Rate]
//   - [TTSSpeechRequest.SetRate]
//   - [TTSSpeechRequest.SentSpeechDone]
//   - [TTSSpeechRequest.SetSentSpeechDone]
//   - [TTSSpeechRequest.SetAudioBufferCallback]
//   - [TTSSpeechRequest.SetLatencyCallback]
//   - [TTSSpeechRequest.SpeechStringType]
//   - [TTSSpeechRequest.SetSpeechStringType]
//   - [TTSSpeechRequest.SynthesisProviderVoice]
//   - [TTSSpeechRequest.SetSynthesisProviderVoice]
//   - [TTSSpeechRequest.SynthesizeSilently]
//   - [TTSSpeechRequest.SetSynthesizeSilently]
//   - [TTSSpeechRequest.Text]
//   - [TTSSpeechRequest.SetText]
//   - [TTSSpeechRequest.Voice]
//   - [TTSSpeechRequest.SetVoice]
//   - [TTSSpeechRequest.VoiceSettings]
//   - [TTSSpeechRequest.SetVoiceSettings]
//   - [TTSSpeechRequest.Volume]
//   - [TTSSpeechRequest.SetVolume]
//   - [TTSSpeechRequest.Voucher]
//   - [TTSSpeechRequest.SetVoucher]
//   - [TTSSpeechRequest.InitWithCoder]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest
type TTSSpeechRequest struct {
	objectivec.Object
}

// TTSSpeechRequestFromID constructs a [TTSSpeechRequest] from an objc.ID.
func TTSSpeechRequestFromID(id objc.ID) TTSSpeechRequest {
	return TTSSpeechRequest{objectivec.Object{ID: id}}
}

// Ensure TTSSpeechRequest implements ITTSSpeechRequest.
var _ ITTSSpeechRequest = TTSSpeechRequest{}

// An interface definition for the [TTSSpeechRequest] class.
//
// # Methods
//
//   - [ITTSSpeechRequest.AudioDeviceId]
//   - [ITTSSpeechRequest.SetAudioDeviceId]
//   - [ITTSSpeechRequest.AudioQueueFlags]
//   - [ITTSSpeechRequest.SetAudioQueueFlags]
//   - [ITTSSpeechRequest.AudioSessionID]
//   - [ITTSSpeechRequest.SetAudioSessionID]
//   - [ITTSSpeechRequest.AudioSessionIDIsValid]
//   - [ITTSSpeechRequest.SetAudioSessionIDIsValid]
//   - [ITTSSpeechRequest.Channels]
//   - [ITTSSpeechRequest.SetChannels]
//   - [ITTSSpeechRequest.ClientContext]
//   - [ITTSSpeechRequest.SetClientContext]
//   - [ITTSSpeechRequest.DispatchTime]
//   - [ITTSSpeechRequest.SetDispatchTime]
//   - [ITTSSpeechRequest.EncodeWithCoder]
//   - [ITTSSpeechRequest.Gender]
//   - [ITTSSpeechRequest.SetGender]
//   - [ITTSSpeechRequest.HandledTime]
//   - [ITTSSpeechRequest.SetHandledTime]
//   - [ITTSSpeechRequest.IgnoreSubstitutions]
//   - [ITTSSpeechRequest.SetIgnoreSubstitutions]
//   - [ITTSSpeechRequest.JobIdentifier]
//   - [ITTSSpeechRequest.SetJobIdentifier]
//   - [ITTSSpeechRequest.LanguageCode]
//   - [ITTSSpeechRequest.SetLanguageCode]
//   - [ITTSSpeechRequest.Latency]
//   - [ITTSSpeechRequest.SetLatency]
//   - [ITTSSpeechRequest.MaintainsInput]
//   - [ITTSSpeechRequest.SetMaintainsInput]
//   - [ITTSSpeechRequest.Pitch]
//   - [ITTSSpeechRequest.SetPitch]
//   - [ITTSSpeechRequest.Rate]
//   - [ITTSSpeechRequest.SetRate]
//   - [ITTSSpeechRequest.SentSpeechDone]
//   - [ITTSSpeechRequest.SetSentSpeechDone]
//   - [ITTSSpeechRequest.SetAudioBufferCallback]
//   - [ITTSSpeechRequest.SetLatencyCallback]
//   - [ITTSSpeechRequest.SpeechStringType]
//   - [ITTSSpeechRequest.SetSpeechStringType]
//   - [ITTSSpeechRequest.SynthesisProviderVoice]
//   - [ITTSSpeechRequest.SetSynthesisProviderVoice]
//   - [ITTSSpeechRequest.SynthesizeSilently]
//   - [ITTSSpeechRequest.SetSynthesizeSilently]
//   - [ITTSSpeechRequest.Text]
//   - [ITTSSpeechRequest.SetText]
//   - [ITTSSpeechRequest.Voice]
//   - [ITTSSpeechRequest.SetVoice]
//   - [ITTSSpeechRequest.VoiceSettings]
//   - [ITTSSpeechRequest.SetVoiceSettings]
//   - [ITTSSpeechRequest.Volume]
//   - [ITTSSpeechRequest.SetVolume]
//   - [ITTSSpeechRequest.Voucher]
//   - [ITTSSpeechRequest.SetVoucher]
//   - [ITTSSpeechRequest.InitWithCoder]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest
type ITTSSpeechRequest interface {
	objectivec.IObject

	// Topic: Methods

	AudioDeviceId() uint32
	SetAudioDeviceId(value uint32)
	AudioQueueFlags() uint32
	SetAudioQueueFlags(value uint32)
	AudioSessionID() uint32
	SetAudioSessionID(value uint32)
	AudioSessionIDIsValid() bool
	SetAudioSessionIDIsValid(value bool)
	Channels() foundation.INSArray
	SetChannels(value foundation.INSArray)
	ClientContext() unsafe.Pointer
	SetClientContext(value unsafe.Pointer)
	DispatchTime() float64
	SetDispatchTime(value float64)
	EncodeWithCoder(coder foundation.INSCoder)
	Gender() int64
	SetGender(value int64)
	HandledTime() float64
	SetHandledTime(value float64)
	IgnoreSubstitutions() bool
	SetIgnoreSubstitutions(value bool)
	JobIdentifier() string
	SetJobIdentifier(value string)
	LanguageCode() string
	SetLanguageCode(value string)
	Latency() float64
	SetLatency(value float64)
	MaintainsInput() bool
	SetMaintainsInput(value bool)
	Pitch() float64
	SetPitch(value float64)
	Rate() float64
	SetRate(value float64)
	SentSpeechDone() bool
	SetSentSpeechDone(value bool)
	SetAudioBufferCallback(callback VoidHandler)
	SetLatencyCallback(callback VoidHandler)
	SpeechStringType() uint64
	SetSpeechStringType(value uint64)
	SynthesisProviderVoice() avfaudio.AVSpeechSynthesisProviderVoice
	SetSynthesisProviderVoice(value avfaudio.AVSpeechSynthesisProviderVoice)
	SynthesizeSilently() bool
	SetSynthesizeSilently(value bool)
	Text() string
	SetText(value string)
	Voice() ITTSSpeechVoice
	SetVoice(value ITTSSpeechVoice)
	VoiceSettings() foundation.INSDictionary
	SetVoiceSettings(value foundation.INSDictionary)
	Volume() float64
	SetVolume(value float64)
	Voucher() objectivec.Object
	SetVoucher(value objectivec.Object)
	InitWithCoder(coder foundation.INSCoder) TTSSpeechRequest
}

// Init initializes the instance.
func (t TTSSpeechRequest) Init() TTSSpeechRequest {
	rv := objc.Send[TTSSpeechRequest](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSSpeechRequest) Autorelease() TTSSpeechRequest {
	rv := objc.Send[TTSSpeechRequest](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSSpeechRequest creates a new TTSSpeechRequest instance.
func NewTTSSpeechRequest() TTSSpeechRequest {
	class := getTTSSpeechRequestClass()
	rv := objc.Send[TTSSpeechRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/initWithCoder:
func NewTTSSpeechRequestWithCoder(coder objectivec.IObject) TTSSpeechRequest {
	instance := getTTSSpeechRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return TTSSpeechRequestFromID(rv)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/encodeWithCoder:
func (t TTSSpeechRequest) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/setAudioBufferCallback:
func (t TTSSpeechRequest) SetAudioBufferCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.Send[objc.ID](t.ID, objc.Sel("setAudioBufferCallback:"), _block0)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/setLatencyCallback:
func (t TTSSpeechRequest) SetLatencyCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.Send[objc.ID](t.ID, objc.Sel("setLatencyCallback:"), _block0)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/initWithCoder:
func (t TTSSpeechRequest) InitWithCoder(coder foundation.INSCoder) TTSSpeechRequest {
	rv := objc.Send[TTSSpeechRequest](t.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/supportsSecureCoding
func (_TTSSpeechRequestClass TTSSpeechRequestClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_TTSSpeechRequestClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/audioDeviceId
func (t TTSSpeechRequest) AudioDeviceId() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("audioDeviceId"))
	return rv
}
func (t TTSSpeechRequest) SetAudioDeviceId(value uint32) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioDeviceId:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/audioQueueFlags
func (t TTSSpeechRequest) AudioQueueFlags() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("audioQueueFlags"))
	return rv
}
func (t TTSSpeechRequest) SetAudioQueueFlags(value uint32) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioQueueFlags:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/audioSessionID
func (t TTSSpeechRequest) AudioSessionID() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("audioSessionID"))
	return rv
}
func (t TTSSpeechRequest) SetAudioSessionID(value uint32) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioSessionID:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/audioSessionIDIsValid
func (t TTSSpeechRequest) AudioSessionIDIsValid() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("audioSessionIDIsValid"))
	return rv
}
func (t TTSSpeechRequest) SetAudioSessionIDIsValid(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioSessionIDIsValid:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/channels
func (t TTSSpeechRequest) Channels() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("channels"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSSpeechRequest) SetChannels(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setChannels:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/clientContext
func (t TTSSpeechRequest) ClientContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t.ID, objc.Sel("clientContext"))
	return rv
}
func (t TTSSpeechRequest) SetClientContext(value unsafe.Pointer) {
	objc.Send[struct{}](t.ID, objc.Sel("setClientContext:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/dispatchTime
func (t TTSSpeechRequest) DispatchTime() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("dispatchTime"))
	return rv
}
func (t TTSSpeechRequest) SetDispatchTime(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setDispatchTime:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/gender
func (t TTSSpeechRequest) Gender() int64 {
	rv := objc.Send[int64](t.ID, objc.Sel("gender"))
	return rv
}
func (t TTSSpeechRequest) SetGender(value int64) {
	objc.Send[struct{}](t.ID, objc.Sel("setGender:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/handledTime
func (t TTSSpeechRequest) HandledTime() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("handledTime"))
	return rv
}
func (t TTSSpeechRequest) SetHandledTime(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setHandledTime:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/ignoreSubstitutions
func (t TTSSpeechRequest) IgnoreSubstitutions() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("ignoreSubstitutions"))
	return rv
}
func (t TTSSpeechRequest) SetIgnoreSubstitutions(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setIgnoreSubstitutions:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/jobIdentifier
func (t TTSSpeechRequest) JobIdentifier() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("jobIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechRequest) SetJobIdentifier(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setJobIdentifier:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/languageCode
func (t TTSSpeechRequest) LanguageCode() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("languageCode"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechRequest) SetLanguageCode(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setLanguageCode:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/latency
func (t TTSSpeechRequest) Latency() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("latency"))
	return rv
}
func (t TTSSpeechRequest) SetLatency(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setLatency:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/maintainsInput
func (t TTSSpeechRequest) MaintainsInput() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("maintainsInput"))
	return rv
}
func (t TTSSpeechRequest) SetMaintainsInput(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setMaintainsInput:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/pitch
func (t TTSSpeechRequest) Pitch() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("pitch"))
	return rv
}
func (t TTSSpeechRequest) SetPitch(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setPitch:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/rate
func (t TTSSpeechRequest) Rate() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("rate"))
	return rv
}
func (t TTSSpeechRequest) SetRate(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setRate:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/sentSpeechDone
func (t TTSSpeechRequest) SentSpeechDone() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("sentSpeechDone"))
	return rv
}
func (t TTSSpeechRequest) SetSentSpeechDone(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setSentSpeechDone:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/speechStringType
func (t TTSSpeechRequest) SpeechStringType() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("speechStringType"))
	return rv
}
func (t TTSSpeechRequest) SetSpeechStringType(value uint64) {
	objc.Send[struct{}](t.ID, objc.Sel("setSpeechStringType:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/synthesisProviderVoice
func (t TTSSpeechRequest) SynthesisProviderVoice() avfaudio.AVSpeechSynthesisProviderVoice {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("synthesisProviderVoice"))
	return avfaudio.AVSpeechSynthesisProviderVoiceFromID(objc.ID(rv))
}
func (t TTSSpeechRequest) SetSynthesisProviderVoice(value avfaudio.AVSpeechSynthesisProviderVoice) {
	objc.Send[struct{}](t.ID, objc.Sel("setSynthesisProviderVoice:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/synthesizeSilently
func (t TTSSpeechRequest) SynthesizeSilently() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("synthesizeSilently"))
	return rv
}
func (t TTSSpeechRequest) SetSynthesizeSilently(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setSynthesizeSilently:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/text
func (t TTSSpeechRequest) Text() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("text"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechRequest) SetText(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setText:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/voice
func (t TTSSpeechRequest) Voice() ITTSSpeechVoice {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("voice"))
	return TTSSpeechVoiceFromID(objc.ID(rv))
}
func (t TTSSpeechRequest) SetVoice(value ITTSSpeechVoice) {
	objc.Send[struct{}](t.ID, objc.Sel("setVoice:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/voiceSettings
func (t TTSSpeechRequest) VoiceSettings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("voiceSettings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t TTSSpeechRequest) SetVoiceSettings(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setVoiceSettings:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/volume
func (t TTSSpeechRequest) Volume() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("volume"))
	return rv
}
func (t TTSSpeechRequest) SetVolume(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setVolume:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechRequest/voucher
func (t TTSSpeechRequest) Voucher() objectivec.Object {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("voucher"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (t TTSSpeechRequest) SetVoucher(value objectivec.Object) {
	objc.Send[struct{}](t.ID, objc.Sel("setVoucher:"), value)
}

// SetAudioBufferCallbackSync is a synchronous wrapper around [TTSSpeechRequest.SetAudioBufferCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSSpeechRequest) SetAudioBufferCallbackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.SetAudioBufferCallback(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetLatencyCallbackSync is a synchronous wrapper around [TTSSpeechRequest.SetLatencyCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSSpeechRequest) SetLatencyCallbackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.SetLatencyCallback(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
