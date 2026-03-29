// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSSpeechSynthesizer] class.
var (
	_TTSSpeechSynthesizerClass     TTSSpeechSynthesizerClass
	_TTSSpeechSynthesizerClassOnce sync.Once
)

func getTTSSpeechSynthesizerClass() TTSSpeechSynthesizerClass {
	_TTSSpeechSynthesizerClassOnce.Do(func() {
		_TTSSpeechSynthesizerClass = TTSSpeechSynthesizerClass{class: objc.GetClass("TTSSpeechSynthesizer")}
	})
	return _TTSSpeechSynthesizerClass
}

// GetTTSSpeechSynthesizerClass returns the class object for TTSSpeechSynthesizer.
func GetTTSSpeechSynthesizerClass() TTSSpeechSynthesizerClass {
	return getTTSSpeechSynthesizerClass()
}

type TTSSpeechSynthesizerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSSpeechSynthesizerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSSpeechSynthesizerClass) Alloc() TTSSpeechSynthesizer {
	rv := objc.Send[TTSSpeechSynthesizer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSSpeechSynthesizer._continueSpeakingRequestWithError]
//   - [TTSSpeechSynthesizer._makeRequestForVoiceAndLanguageCode]
//   - [TTSSpeechSynthesizer._mediaServicesDied]
//   - [TTSSpeechSynthesizer._pauseSpeakingRequestAtNextBoundarySynchronouslyError]
//   - [TTSSpeechSynthesizer._preprocessTextLanguageCode]
//   - [TTSSpeechSynthesizer._processMarkerForRequest]
//   - [TTSSpeechSynthesizer._resolveVoiceForLanguage]
//   - [TTSSpeechSynthesizer._setDelegate]
//   - [TTSSpeechSynthesizer._startSpeakingStringOrSSMLStringWithLanguageCodeJobIdRequestError]
//   - [TTSSpeechSynthesizer._stopSpeakingRequestAtNextBoundarySynchronouslyError]
//   - [TTSSpeechSynthesizer.AudioDeviceId]
//   - [TTSSpeechSynthesizer.SetAudioDeviceId]
//   - [TTSSpeechSynthesizer.AudioEffects]
//   - [TTSSpeechSynthesizer.SetAudioEffects]
//   - [TTSSpeechSynthesizer.AudioQueueFlags]
//   - [TTSSpeechSynthesizer.SetAudioQueueFlags]
//   - [TTSSpeechSynthesizer.BundleIdentifier]
//   - [TTSSpeechSynthesizer.SetBundleIdentifier]
//   - [TTSSpeechSynthesizer.ContinueSpeakingRequestWithError]
//   - [TTSSpeechSynthesizer.ContinueSpeakingWithError]
//   - [TTSSpeechSynthesizer.CoreSynth]
//   - [TTSSpeechSynthesizer.SetCoreSynth]
//   - [TTSSpeechSynthesizer.Delegate]
//   - [TTSSpeechSynthesizer.SetDelegate]
//   - [TTSSpeechSynthesizer.DelegateTargetQueue]
//   - [TTSSpeechSynthesizer.SetDelegateTargetQueue]
//   - [TTSSpeechSynthesizer.Footprint]
//   - [TTSSpeechSynthesizer.GetPerVoiceSettings]
//   - [TTSSpeechSynthesizer.IgnoreSubstitutions]
//   - [TTSSpeechSynthesizer.SetIgnoreSubstitutions]
//   - [TTSSpeechSynthesizer.IsSpeaking]
//   - [TTSSpeechSynthesizer.MaximumRate]
//   - [TTSSpeechSynthesizer.MinimumRate]
//   - [TTSSpeechSynthesizer.NormalizedRate]
//   - [TTSSpeechSynthesizer.SetNormalizedRate]
//   - [TTSSpeechSynthesizer.OutputChannels]
//   - [TTSSpeechSynthesizer.SetOutputChannels]
//   - [TTSSpeechSynthesizer.PauseSpeakingAtNextBoundaryError]
//   - [TTSSpeechSynthesizer.PauseSpeakingAtNextBoundarySynchronouslyError]
//   - [TTSSpeechSynthesizer.PauseSpeakingRequestAtNextBoundaryError]
//   - [TTSSpeechSynthesizer.PauseSpeakingRequestAtNextBoundarySynchronouslyError]
//   - [TTSSpeechSynthesizer.PerVoiceSettings]
//   - [TTSSpeechSynthesizer.SetPerVoiceSettings]
//   - [TTSSpeechSynthesizer.PhonemeSubstitutions]
//   - [TTSSpeechSynthesizer.SetPhonemeSubstitutions]
//   - [TTSSpeechSynthesizer.Pitch]
//   - [TTSSpeechSynthesizer.SetPitch]
//   - [TTSSpeechSynthesizer.Rate]
//   - [TTSSpeechSynthesizer.SetRate]
//   - [TTSSpeechSynthesizer.RequestClientIdentifier]
//   - [TTSSpeechSynthesizer.SetRequestClientIdentifier]
//   - [TTSSpeechSynthesizer.ResolvedVoiceIdentifier]
//   - [TTSSpeechSynthesizer.ResolvedVoiceIdentifierForLanguageCode]
//   - [TTSSpeechSynthesizer.SetAudioBufferCallback]
//   - [TTSSpeechSynthesizer.SetFootprint]
//   - [TTSSpeechSynthesizer.SkipLuthorRules]
//   - [TTSSpeechSynthesizer.SetSkipLuthorRules]
//   - [TTSSpeechSynthesizer.SpeakingRequestClientContext]
//   - [TTSSpeechSynthesizer.SetSpeakingRequestClientContext]
//   - [TTSSpeechSynthesizer.SpeechRequestDidStopWithSuccessPhonemesSpokenError]
//   - [TTSSpeechSynthesizer.SpeechRequestWithMarker]
//   - [TTSSpeechSynthesizer.SpeechRequestDidContinue]
//   - [TTSSpeechSynthesizer.SpeechRequestDidPause]
//   - [TTSSpeechSynthesizer.SpeechRequestDidStart]
//   - [TTSSpeechSynthesizer.SpeechSource]
//   - [TTSSpeechSynthesizer.SetSpeechSource]
//   - [TTSSpeechSynthesizer.SpeechString]
//   - [TTSSpeechSynthesizer.StartSpeakingSSMLWithLanguageCodeJobIdentifierRequestError]
//   - [TTSSpeechSynthesizer.StartSpeakingSSMLWithLanguageCodeRequestError]
//   - [TTSSpeechSynthesizer.StartSpeakingStringError]
//   - [TTSSpeechSynthesizer.StartSpeakingStringRequestError]
//   - [TTSSpeechSynthesizer.StartSpeakingStringToURLWithLanguageCodeError]
//   - [TTSSpeechSynthesizer.StartSpeakingStringToURLWithLanguageCodeRequestError]
//   - [TTSSpeechSynthesizer.StartSpeakingStringWithLanguageCodeError]
//   - [TTSSpeechSynthesizer.StartSpeakingStringWithLanguageCodeJobIdentifierRequestError]
//   - [TTSSpeechSynthesizer.StartSpeakingStringWithLanguageCodeRequestError]
//   - [TTSSpeechSynthesizer.StopSpeakingAtNextBoundaryError]
//   - [TTSSpeechSynthesizer.StopSpeakingAtNextBoundarySynchronouslyError]
//   - [TTSSpeechSynthesizer.StopSpeakingRequestAtNextBoundaryError]
//   - [TTSSpeechSynthesizer.StopSpeakingRequestAtNextBoundarySynchronouslyError]
//   - [TTSSpeechSynthesizer.SupportsAccurateWordCallbacks]
//   - [TTSSpeechSynthesizer.SetSupportsAccurateWordCallbacks]
//   - [TTSSpeechSynthesizer.SynthesizeSilently]
//   - [TTSSpeechSynthesizer.SetSynthesizeSilently]
//   - [TTSSpeechSynthesizer.TestingLastRuleConversion]
//   - [TTSSpeechSynthesizer.TestingSetLastRuleConversionReplacement]
//   - [TTSSpeechSynthesizer.UpdateCoreSynthSubstitutions]
//   - [TTSSpeechSynthesizer.UseSpecificAudioSession]
//   - [TTSSpeechSynthesizer.UserSubstitutions]
//   - [TTSSpeechSynthesizer.SetUserSubstitutions]
//   - [TTSSpeechSynthesizer.VoiceIdentifier]
//   - [TTSSpeechSynthesizer.SetVoiceIdentifier]
//   - [TTSSpeechSynthesizer.VoiceResolver]
//   - [TTSSpeechSynthesizer.Volume]
//   - [TTSSpeechSynthesizer.SetVolume]
//   - [TTSSpeechSynthesizer.Voucher]
//   - [TTSSpeechSynthesizer.SetVoucher]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer
type TTSSpeechSynthesizer struct {
	objectivec.Object
}

// TTSSpeechSynthesizerFromID constructs a [TTSSpeechSynthesizer] from an objc.ID.
func TTSSpeechSynthesizerFromID(id objc.ID) TTSSpeechSynthesizer {
	return TTSSpeechSynthesizer{objectivec.Object{ID: id}}
}
// Ensure TTSSpeechSynthesizer implements ITTSSpeechSynthesizer.
var _ ITTSSpeechSynthesizer = TTSSpeechSynthesizer{}

// An interface definition for the [TTSSpeechSynthesizer] class.
//
// # Methods
//
//   - [ITTSSpeechSynthesizer._continueSpeakingRequestWithError]
//   - [ITTSSpeechSynthesizer._makeRequestForVoiceAndLanguageCode]
//   - [ITTSSpeechSynthesizer._mediaServicesDied]
//   - [ITTSSpeechSynthesizer._pauseSpeakingRequestAtNextBoundarySynchronouslyError]
//   - [ITTSSpeechSynthesizer._preprocessTextLanguageCode]
//   - [ITTSSpeechSynthesizer._processMarkerForRequest]
//   - [ITTSSpeechSynthesizer._resolveVoiceForLanguage]
//   - [ITTSSpeechSynthesizer._setDelegate]
//   - [ITTSSpeechSynthesizer._startSpeakingStringOrSSMLStringWithLanguageCodeJobIdRequestError]
//   - [ITTSSpeechSynthesizer._stopSpeakingRequestAtNextBoundarySynchronouslyError]
//   - [ITTSSpeechSynthesizer.AudioDeviceId]
//   - [ITTSSpeechSynthesizer.SetAudioDeviceId]
//   - [ITTSSpeechSynthesizer.AudioEffects]
//   - [ITTSSpeechSynthesizer.SetAudioEffects]
//   - [ITTSSpeechSynthesizer.AudioQueueFlags]
//   - [ITTSSpeechSynthesizer.SetAudioQueueFlags]
//   - [ITTSSpeechSynthesizer.BundleIdentifier]
//   - [ITTSSpeechSynthesizer.SetBundleIdentifier]
//   - [ITTSSpeechSynthesizer.ContinueSpeakingRequestWithError]
//   - [ITTSSpeechSynthesizer.ContinueSpeakingWithError]
//   - [ITTSSpeechSynthesizer.CoreSynth]
//   - [ITTSSpeechSynthesizer.SetCoreSynth]
//   - [ITTSSpeechSynthesizer.Delegate]
//   - [ITTSSpeechSynthesizer.SetDelegate]
//   - [ITTSSpeechSynthesizer.DelegateTargetQueue]
//   - [ITTSSpeechSynthesizer.SetDelegateTargetQueue]
//   - [ITTSSpeechSynthesizer.Footprint]
//   - [ITTSSpeechSynthesizer.GetPerVoiceSettings]
//   - [ITTSSpeechSynthesizer.IgnoreSubstitutions]
//   - [ITTSSpeechSynthesizer.SetIgnoreSubstitutions]
//   - [ITTSSpeechSynthesizer.IsSpeaking]
//   - [ITTSSpeechSynthesizer.MaximumRate]
//   - [ITTSSpeechSynthesizer.MinimumRate]
//   - [ITTSSpeechSynthesizer.NormalizedRate]
//   - [ITTSSpeechSynthesizer.SetNormalizedRate]
//   - [ITTSSpeechSynthesizer.OutputChannels]
//   - [ITTSSpeechSynthesizer.SetOutputChannels]
//   - [ITTSSpeechSynthesizer.PauseSpeakingAtNextBoundaryError]
//   - [ITTSSpeechSynthesizer.PauseSpeakingAtNextBoundarySynchronouslyError]
//   - [ITTSSpeechSynthesizer.PauseSpeakingRequestAtNextBoundaryError]
//   - [ITTSSpeechSynthesizer.PauseSpeakingRequestAtNextBoundarySynchronouslyError]
//   - [ITTSSpeechSynthesizer.PerVoiceSettings]
//   - [ITTSSpeechSynthesizer.SetPerVoiceSettings]
//   - [ITTSSpeechSynthesizer.PhonemeSubstitutions]
//   - [ITTSSpeechSynthesizer.SetPhonemeSubstitutions]
//   - [ITTSSpeechSynthesizer.Pitch]
//   - [ITTSSpeechSynthesizer.SetPitch]
//   - [ITTSSpeechSynthesizer.Rate]
//   - [ITTSSpeechSynthesizer.SetRate]
//   - [ITTSSpeechSynthesizer.RequestClientIdentifier]
//   - [ITTSSpeechSynthesizer.SetRequestClientIdentifier]
//   - [ITTSSpeechSynthesizer.ResolvedVoiceIdentifier]
//   - [ITTSSpeechSynthesizer.ResolvedVoiceIdentifierForLanguageCode]
//   - [ITTSSpeechSynthesizer.SetAudioBufferCallback]
//   - [ITTSSpeechSynthesizer.SetFootprint]
//   - [ITTSSpeechSynthesizer.SkipLuthorRules]
//   - [ITTSSpeechSynthesizer.SetSkipLuthorRules]
//   - [ITTSSpeechSynthesizer.SpeakingRequestClientContext]
//   - [ITTSSpeechSynthesizer.SetSpeakingRequestClientContext]
//   - [ITTSSpeechSynthesizer.SpeechRequestDidStopWithSuccessPhonemesSpokenError]
//   - [ITTSSpeechSynthesizer.SpeechRequestWithMarker]
//   - [ITTSSpeechSynthesizer.SpeechRequestDidContinue]
//   - [ITTSSpeechSynthesizer.SpeechRequestDidPause]
//   - [ITTSSpeechSynthesizer.SpeechRequestDidStart]
//   - [ITTSSpeechSynthesizer.SpeechSource]
//   - [ITTSSpeechSynthesizer.SetSpeechSource]
//   - [ITTSSpeechSynthesizer.SpeechString]
//   - [ITTSSpeechSynthesizer.StartSpeakingSSMLWithLanguageCodeJobIdentifierRequestError]
//   - [ITTSSpeechSynthesizer.StartSpeakingSSMLWithLanguageCodeRequestError]
//   - [ITTSSpeechSynthesizer.StartSpeakingStringError]
//   - [ITTSSpeechSynthesizer.StartSpeakingStringRequestError]
//   - [ITTSSpeechSynthesizer.StartSpeakingStringToURLWithLanguageCodeError]
//   - [ITTSSpeechSynthesizer.StartSpeakingStringToURLWithLanguageCodeRequestError]
//   - [ITTSSpeechSynthesizer.StartSpeakingStringWithLanguageCodeError]
//   - [ITTSSpeechSynthesizer.StartSpeakingStringWithLanguageCodeJobIdentifierRequestError]
//   - [ITTSSpeechSynthesizer.StartSpeakingStringWithLanguageCodeRequestError]
//   - [ITTSSpeechSynthesizer.StopSpeakingAtNextBoundaryError]
//   - [ITTSSpeechSynthesizer.StopSpeakingAtNextBoundarySynchronouslyError]
//   - [ITTSSpeechSynthesizer.StopSpeakingRequestAtNextBoundaryError]
//   - [ITTSSpeechSynthesizer.StopSpeakingRequestAtNextBoundarySynchronouslyError]
//   - [ITTSSpeechSynthesizer.SupportsAccurateWordCallbacks]
//   - [ITTSSpeechSynthesizer.SetSupportsAccurateWordCallbacks]
//   - [ITTSSpeechSynthesizer.SynthesizeSilently]
//   - [ITTSSpeechSynthesizer.SetSynthesizeSilently]
//   - [ITTSSpeechSynthesizer.TestingLastRuleConversion]
//   - [ITTSSpeechSynthesizer.TestingSetLastRuleConversionReplacement]
//   - [ITTSSpeechSynthesizer.UpdateCoreSynthSubstitutions]
//   - [ITTSSpeechSynthesizer.UseSpecificAudioSession]
//   - [ITTSSpeechSynthesizer.UserSubstitutions]
//   - [ITTSSpeechSynthesizer.SetUserSubstitutions]
//   - [ITTSSpeechSynthesizer.VoiceIdentifier]
//   - [ITTSSpeechSynthesizer.SetVoiceIdentifier]
//   - [ITTSSpeechSynthesizer.VoiceResolver]
//   - [ITTSSpeechSynthesizer.Volume]
//   - [ITTSSpeechSynthesizer.SetVolume]
//   - [ITTSSpeechSynthesizer.Voucher]
//   - [ITTSSpeechSynthesizer.SetVoucher]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer
type ITTSSpeechSynthesizer interface {
	objectivec.IObject

	// Topic: Methods

	_continueSpeakingRequestWithError(request objectivec.IObject) (bool, error)
	_makeRequestForVoiceAndLanguageCode(voice objectivec.IObject, code objectivec.IObject) objectivec.IObject
	_mediaServicesDied()
	_pauseSpeakingRequestAtNextBoundarySynchronouslyError(request objectivec.IObject, boundary int64, synchronously bool) (bool, error)
	_preprocessTextLanguageCode(text objectivec.IObject, code objectivec.IObject) objectivec.IObject
	_processMarkerForRequest(marker objectivec.IObject, request objectivec.IObject) objectivec.IObject
	_resolveVoiceForLanguage(language objectivec.IObject) objectivec.IObject
	_setDelegate(delegate objectivec.IObject)
	_startSpeakingStringOrSSMLStringWithLanguageCodeJobIdRequestError(string_ objectivec.IObject, sSMLString objectivec.IObject, code objectivec.IObject, id objectivec.IObject, request []objectivec.IObject) (bool, error)
	_stopSpeakingRequestAtNextBoundarySynchronouslyError(request objectivec.IObject, boundary int64, synchronously bool) (bool, error)
	AudioDeviceId() uint32
	SetAudioDeviceId(value uint32)
	AudioEffects() foundation.INSArray
	SetAudioEffects(value foundation.INSArray)
	AudioQueueFlags() uint32
	SetAudioQueueFlags(value uint32)
	BundleIdentifier() string
	SetBundleIdentifier(value string)
	ContinueSpeakingRequestWithError(request objectivec.IObject) (bool, error)
	ContinueSpeakingWithError() (bool, error)
	CoreSynth() unsafe.Pointer
	SetCoreSynth(value unsafe.Pointer)
	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)
	DelegateTargetQueue() objectivec.Object
	SetDelegateTargetQueue(value objectivec.Object)
	Footprint() int64
	GetPerVoiceSettings() objectivec.IObject
	IgnoreSubstitutions() bool
	SetIgnoreSubstitutions(value bool)
	IsSpeaking() bool
	MaximumRate() float32
	MinimumRate() float32
	NormalizedRate() float32
	SetNormalizedRate(value float32)
	OutputChannels() foundation.INSArray
	SetOutputChannels(value foundation.INSArray)
	PauseSpeakingAtNextBoundaryError(boundary int64) (bool, error)
	PauseSpeakingAtNextBoundarySynchronouslyError(boundary int64, synchronously bool) (bool, error)
	PauseSpeakingRequestAtNextBoundaryError(request objectivec.IObject, boundary int64) (bool, error)
	PauseSpeakingRequestAtNextBoundarySynchronouslyError(request objectivec.IObject, boundary int64, synchronously bool) (bool, error)
	PerVoiceSettings() foundation.INSDictionary
	SetPerVoiceSettings(value foundation.INSDictionary)
	PhonemeSubstitutions() foundation.INSArray
	SetPhonemeSubstitutions(value foundation.INSArray)
	Pitch() float32
	SetPitch(value float32)
	Rate() float32
	SetRate(value float32)
	RequestClientIdentifier() uint64
	SetRequestClientIdentifier(value uint64)
	ResolvedVoiceIdentifier() string
	ResolvedVoiceIdentifierForLanguageCode(code objectivec.IObject) objectivec.IObject
	SetAudioBufferCallback(callback VoidHandler)
	SetFootprint(footprint int64)
	SkipLuthorRules() bool
	SetSkipLuthorRules(value bool)
	SpeakingRequestClientContext() unsafe.Pointer
	SetSpeakingRequestClientContext(value unsafe.Pointer)
	SpeechRequestDidStopWithSuccessPhonemesSpokenError(request objectivec.IObject, success bool, spoken objectivec.IObject, error_ objectivec.IObject)
	SpeechRequestWithMarker(request objectivec.IObject, marker objectivec.IObject)
	SpeechRequestDidContinue(continue_ objectivec.IObject)
	SpeechRequestDidPause(pause objectivec.IObject)
	SpeechRequestDidStart(start objectivec.IObject)
	SpeechSource() string
	SetSpeechSource(value string)
	SpeechString() objectivec.IObject
	StartSpeakingSSMLWithLanguageCodeJobIdentifierRequestError(ssml objectivec.IObject, code objectivec.IObject, identifier objectivec.IObject, request []objectivec.IObject) (bool, error)
	StartSpeakingSSMLWithLanguageCodeRequestError(ssml objectivec.IObject, code objectivec.IObject, request []objectivec.IObject) (bool, error)
	StartSpeakingStringError(string_ objectivec.IObject) (bool, error)
	StartSpeakingStringRequestError(string_ objectivec.IObject, request []objectivec.IObject) (bool, error)
	StartSpeakingStringToURLWithLanguageCodeError(string_ objectivec.IObject, url foundation.INSURL, code objectivec.IObject) (bool, error)
	StartSpeakingStringToURLWithLanguageCodeRequestError(string_ objectivec.IObject, url foundation.INSURL, code objectivec.IObject, request []objectivec.IObject) (bool, error)
	StartSpeakingStringWithLanguageCodeError(string_ objectivec.IObject, code objectivec.IObject) (bool, error)
	StartSpeakingStringWithLanguageCodeJobIdentifierRequestError(string_ objectivec.IObject, code objectivec.IObject, identifier objectivec.IObject, request []objectivec.IObject) (bool, error)
	StartSpeakingStringWithLanguageCodeRequestError(string_ objectivec.IObject, code objectivec.IObject, request []objectivec.IObject) (bool, error)
	StopSpeakingAtNextBoundaryError(boundary int64) (bool, error)
	StopSpeakingAtNextBoundarySynchronouslyError(boundary int64, synchronously bool) (bool, error)
	StopSpeakingRequestAtNextBoundaryError(request objectivec.IObject, boundary int64) (bool, error)
	StopSpeakingRequestAtNextBoundarySynchronouslyError(request objectivec.IObject, boundary int64, synchronously bool) (bool, error)
	SupportsAccurateWordCallbacks() bool
	SetSupportsAccurateWordCallbacks(value bool)
	SynthesizeSilently() bool
	SetSynthesizeSilently(value bool)
	TestingLastRuleConversion() objectivec.IObject
	TestingSetLastRuleConversionReplacement(conversion objectivec.IObject, replacement objectivec.IObject)
	UpdateCoreSynthSubstitutions()
	UseSpecificAudioSession(session uint32)
	UserSubstitutions() foundation.INSArray
	SetUserSubstitutions(value foundation.INSArray)
	VoiceIdentifier() string
	SetVoiceIdentifier(value string)
	VoiceResolver() objectivec.IObject
	Volume() float32
	SetVolume(value float32)
	Voucher() objectivec.Object
	SetVoucher(value objectivec.Object)
}

// Init initializes the instance.
func (t TTSSpeechSynthesizer) Init() TTSSpeechSynthesizer {
	rv := objc.Send[TTSSpeechSynthesizer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSSpeechSynthesizer) Autorelease() TTSSpeechSynthesizer {
	rv := objc.Send[TTSSpeechSynthesizer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSSpeechSynthesizer creates a new TTSSpeechSynthesizer instance.
func NewTTSSpeechSynthesizer() TTSSpeechSynthesizer {
	class := getTTSSpeechSynthesizerClass()
	rv := objc.Send[TTSSpeechSynthesizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/_continueSpeakingRequest:withError:
func (t TTSSpeechSynthesizer) _continueSpeakingRequestWithError(request objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("_continueSpeakingRequest:withError:"), request, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_continueSpeakingRequest:withError: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/_makeRequestForVoice:andLanguageCode:
func (t TTSSpeechSynthesizer) _makeRequestForVoiceAndLanguageCode(voice objectivec.IObject, code objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("_makeRequestForVoice:andLanguageCode:"), voice, code)
	return objectivec.Object{ID: rv}
}

// MakeRequestForVoiceAndLanguageCode is an exported wrapper for the private method _makeRequestForVoiceAndLanguageCode.
func (t TTSSpeechSynthesizer) MakeRequestForVoiceAndLanguageCode(voice objectivec.IObject, code objectivec.IObject) objectivec.IObject {
	return t._makeRequestForVoiceAndLanguageCode(voice, code)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/_mediaServicesDied
func (t TTSSpeechSynthesizer) _mediaServicesDied() {
	objc.Send[objc.ID](t.ID, objc.Sel("_mediaServicesDied"))
}

// MediaServicesDied is an exported wrapper for the private method _mediaServicesDied.
func (t TTSSpeechSynthesizer) MediaServicesDied() {
	t._mediaServicesDied()
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/_pauseSpeakingRequest:atNextBoundary:synchronously:error:
func (t TTSSpeechSynthesizer) _pauseSpeakingRequestAtNextBoundarySynchronouslyError(request objectivec.IObject, boundary int64, synchronously bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("_pauseSpeakingRequest:atNextBoundary:synchronously:error:"), request, boundary, synchronously, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_pauseSpeakingRequest:atNextBoundary:synchronously:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/_preprocessText:languageCode:
func (t TTSSpeechSynthesizer) _preprocessTextLanguageCode(text objectivec.IObject, code objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("_preprocessText:languageCode:"), text, code)
	return objectivec.Object{ID: rv}
}

// PreprocessTextLanguageCode is an exported wrapper for the private method _preprocessTextLanguageCode.
func (t TTSSpeechSynthesizer) PreprocessTextLanguageCode(text objectivec.IObject, code objectivec.IObject) objectivec.IObject {
	return t._preprocessTextLanguageCode(text, code)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/_processMarker:forRequest:
func (t TTSSpeechSynthesizer) _processMarkerForRequest(marker objectivec.IObject, request objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("_processMarker:forRequest:"), marker, request)
	return objectivec.Object{ID: rv}
}

// ProcessMarkerForRequest is an exported wrapper for the private method _processMarkerForRequest.
func (t TTSSpeechSynthesizer) ProcessMarkerForRequest(marker objectivec.IObject, request objectivec.IObject) objectivec.IObject {
	return t._processMarkerForRequest(marker, request)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/_resolveVoiceForLanguage:
func (t TTSSpeechSynthesizer) _resolveVoiceForLanguage(language objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("_resolveVoiceForLanguage:"), language)
	return objectivec.Object{ID: rv}
}

// ResolveVoiceForLanguage is an exported wrapper for the private method _resolveVoiceForLanguage.
func (t TTSSpeechSynthesizer) ResolveVoiceForLanguage(language objectivec.IObject) objectivec.IObject {
	return t._resolveVoiceForLanguage(language)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/_setDelegate:
func (t TTSSpeechSynthesizer) _setDelegate(delegate objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("_setDelegate:"), delegate)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/_startSpeakingString:orSSMLString:withLanguageCode:jobId:request:error:
func (t TTSSpeechSynthesizer) _startSpeakingStringOrSSMLStringWithLanguageCodeJobIdRequestError(string_ objectivec.IObject, sSMLString objectivec.IObject, code objectivec.IObject, id objectivec.IObject, request []objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("_startSpeakingString:orSSMLString:withLanguageCode:jobId:request:error:"), string_, sSMLString, code, id, objectivec.IObjectSliceToNSArray(request), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_startSpeakingString:orSSMLString:withLanguageCode:jobId:request:error: returned NO with nil NSError")
	}
	return rv, nil

}

// StartSpeakingStringOrSSMLStringWithLanguageCodeJobIdRequestError is an exported wrapper for the private method _startSpeakingStringOrSSMLStringWithLanguageCodeJobIdRequestError.
func (t TTSSpeechSynthesizer) StartSpeakingStringOrSSMLStringWithLanguageCodeJobIdRequestError(string_ objectivec.IObject, sSMLString objectivec.IObject, code objectivec.IObject, id objectivec.IObject, request []objectivec.IObject) (bool, error) {
	return t._startSpeakingStringOrSSMLStringWithLanguageCodeJobIdRequestError(string_, sSMLString, code, id, request)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/_stopSpeakingRequest:atNextBoundary:synchronously:error:
func (t TTSSpeechSynthesizer) _stopSpeakingRequestAtNextBoundarySynchronouslyError(request objectivec.IObject, boundary int64, synchronously bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("_stopSpeakingRequest:atNextBoundary:synchronously:error:"), request, boundary, synchronously, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_stopSpeakingRequest:atNextBoundary:synchronously:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/continueSpeakingRequest:withError:
func (t TTSSpeechSynthesizer) ContinueSpeakingRequestWithError(request objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("continueSpeakingRequest:withError:"), request, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("continueSpeakingRequest:withError: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/continueSpeakingWithError:
func (t TTSSpeechSynthesizer) ContinueSpeakingWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("continueSpeakingWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("continueSpeakingWithError: returned NO with nil NSError")
	}
	return rv, nil

}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/footprint
func (t TTSSpeechSynthesizer) Footprint() int64 {
	rv := objc.Send[int64](t.ID, objc.Sel("footprint"))
	return rv
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/getPerVoiceSettings
func (t TTSSpeechSynthesizer) GetPerVoiceSettings() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("getPerVoiceSettings"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/isSpeaking
func (t TTSSpeechSynthesizer) IsSpeaking() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isSpeaking"))
	return rv
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/maximumRate
func (t TTSSpeechSynthesizer) MaximumRate() float32 {
	rv := objc.Send[float32](t.ID, objc.Sel("maximumRate"))
	return rv
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/minimumRate
func (t TTSSpeechSynthesizer) MinimumRate() float32 {
	rv := objc.Send[float32](t.ID, objc.Sel("minimumRate"))
	return rv
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/pauseSpeakingAtNextBoundary:error:
func (t TTSSpeechSynthesizer) PauseSpeakingAtNextBoundaryError(boundary int64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("pauseSpeakingAtNextBoundary:error:"), boundary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("pauseSpeakingAtNextBoundary:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/pauseSpeakingAtNextBoundary:synchronously:error:
func (t TTSSpeechSynthesizer) PauseSpeakingAtNextBoundarySynchronouslyError(boundary int64, synchronously bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("pauseSpeakingAtNextBoundary:synchronously:error:"), boundary, synchronously, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("pauseSpeakingAtNextBoundary:synchronously:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/pauseSpeakingRequest:atNextBoundary:error:
func (t TTSSpeechSynthesizer) PauseSpeakingRequestAtNextBoundaryError(request objectivec.IObject, boundary int64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("pauseSpeakingRequest:atNextBoundary:error:"), request, boundary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("pauseSpeakingRequest:atNextBoundary:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/pauseSpeakingRequest:atNextBoundary:synchronously:error:
func (t TTSSpeechSynthesizer) PauseSpeakingRequestAtNextBoundarySynchronouslyError(request objectivec.IObject, boundary int64, synchronously bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("pauseSpeakingRequest:atNextBoundary:synchronously:error:"), request, boundary, synchronously, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("pauseSpeakingRequest:atNextBoundary:synchronously:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/resolvedVoiceIdentifierForLanguageCode:
func (t TTSSpeechSynthesizer) ResolvedVoiceIdentifierForLanguageCode(code objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("resolvedVoiceIdentifierForLanguageCode:"), code)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/setAudioBufferCallback:
func (t TTSSpeechSynthesizer) SetAudioBufferCallback(callback VoidHandler) {
_block0, _ := NewVoidBlock(callback)
	objc.Send[objc.ID](t.ID, objc.Sel("setAudioBufferCallback:"), _block0)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/setFootprint:
func (t TTSSpeechSynthesizer) SetFootprint(footprint int64) {
	objc.Send[objc.ID](t.ID, objc.Sel("setFootprint:"), footprint)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/speechRequest:didStopWithSuccess:phonemesSpoken:error:
func (t TTSSpeechSynthesizer) SpeechRequestDidStopWithSuccessPhonemesSpokenError(request objectivec.IObject, success bool, spoken objectivec.IObject, error_ objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("speechRequest:didStopWithSuccess:phonemesSpoken:error:"), request, success, spoken, error_)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/speechRequest:withMarker:
func (t TTSSpeechSynthesizer) SpeechRequestWithMarker(request objectivec.IObject, marker objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("speechRequest:withMarker:"), request, marker)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/speechRequestDidContinue:
func (t TTSSpeechSynthesizer) SpeechRequestDidContinue(continue_ objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("speechRequestDidContinue:"), continue_)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/speechRequestDidPause:
func (t TTSSpeechSynthesizer) SpeechRequestDidPause(pause objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("speechRequestDidPause:"), pause)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/speechRequestDidStart:
func (t TTSSpeechSynthesizer) SpeechRequestDidStart(start objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("speechRequestDidStart:"), start)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/speechString
func (t TTSSpeechSynthesizer) SpeechString() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("speechString"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/startSpeakingSSML:withLanguageCode:jobIdentifier:request:error:
func (t TTSSpeechSynthesizer) StartSpeakingSSMLWithLanguageCodeJobIdentifierRequestError(ssml objectivec.IObject, code objectivec.IObject, identifier objectivec.IObject, request []objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("startSpeakingSSML:withLanguageCode:jobIdentifier:request:error:"), ssml, code, identifier, objectivec.IObjectSliceToNSArray(request), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startSpeakingSSML:withLanguageCode:jobIdentifier:request:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/startSpeakingSSML:withLanguageCode:request:error:
func (t TTSSpeechSynthesizer) StartSpeakingSSMLWithLanguageCodeRequestError(ssml objectivec.IObject, code objectivec.IObject, request []objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("startSpeakingSSML:withLanguageCode:request:error:"), ssml, code, objectivec.IObjectSliceToNSArray(request), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startSpeakingSSML:withLanguageCode:request:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/startSpeakingString:error:
func (t TTSSpeechSynthesizer) StartSpeakingStringError(string_ objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("startSpeakingString:error:"), string_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startSpeakingString:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/startSpeakingString:request:error:
func (t TTSSpeechSynthesizer) StartSpeakingStringRequestError(string_ objectivec.IObject, request []objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("startSpeakingString:request:error:"), string_, objectivec.IObjectSliceToNSArray(request), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startSpeakingString:request:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/startSpeakingString:toURL:withLanguageCode:error:
func (t TTSSpeechSynthesizer) StartSpeakingStringToURLWithLanguageCodeError(string_ objectivec.IObject, url foundation.INSURL, code objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("startSpeakingString:toURL:withLanguageCode:error:"), string_, url, code, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startSpeakingString:toURL:withLanguageCode:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/startSpeakingString:toURL:withLanguageCode:request:error:
func (t TTSSpeechSynthesizer) StartSpeakingStringToURLWithLanguageCodeRequestError(string_ objectivec.IObject, url foundation.INSURL, code objectivec.IObject, request []objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("startSpeakingString:toURL:withLanguageCode:request:error:"), string_, url, code, objectivec.IObjectSliceToNSArray(request), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startSpeakingString:toURL:withLanguageCode:request:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/startSpeakingString:withLanguageCode:error:
func (t TTSSpeechSynthesizer) StartSpeakingStringWithLanguageCodeError(string_ objectivec.IObject, code objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("startSpeakingString:withLanguageCode:error:"), string_, code, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startSpeakingString:withLanguageCode:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/startSpeakingString:withLanguageCode:jobIdentifier:request:error:
func (t TTSSpeechSynthesizer) StartSpeakingStringWithLanguageCodeJobIdentifierRequestError(string_ objectivec.IObject, code objectivec.IObject, identifier objectivec.IObject, request []objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("startSpeakingString:withLanguageCode:jobIdentifier:request:error:"), string_, code, identifier, objectivec.IObjectSliceToNSArray(request), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startSpeakingString:withLanguageCode:jobIdentifier:request:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/startSpeakingString:withLanguageCode:request:error:
func (t TTSSpeechSynthesizer) StartSpeakingStringWithLanguageCodeRequestError(string_ objectivec.IObject, code objectivec.IObject, request []objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("startSpeakingString:withLanguageCode:request:error:"), string_, code, objectivec.IObjectSliceToNSArray(request), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startSpeakingString:withLanguageCode:request:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/stopSpeakingAtNextBoundary:error:
func (t TTSSpeechSynthesizer) StopSpeakingAtNextBoundaryError(boundary int64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("stopSpeakingAtNextBoundary:error:"), boundary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("stopSpeakingAtNextBoundary:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/stopSpeakingAtNextBoundary:synchronously:error:
func (t TTSSpeechSynthesizer) StopSpeakingAtNextBoundarySynchronouslyError(boundary int64, synchronously bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("stopSpeakingAtNextBoundary:synchronously:error:"), boundary, synchronously, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("stopSpeakingAtNextBoundary:synchronously:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/stopSpeakingRequest:atNextBoundary:error:
func (t TTSSpeechSynthesizer) StopSpeakingRequestAtNextBoundaryError(request objectivec.IObject, boundary int64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("stopSpeakingRequest:atNextBoundary:error:"), request, boundary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("stopSpeakingRequest:atNextBoundary:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/stopSpeakingRequest:atNextBoundary:synchronously:error:
func (t TTSSpeechSynthesizer) StopSpeakingRequestAtNextBoundarySynchronouslyError(request objectivec.IObject, boundary int64, synchronously bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("stopSpeakingRequest:atNextBoundary:synchronously:error:"), request, boundary, synchronously, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("stopSpeakingRequest:atNextBoundary:synchronously:error: returned NO with nil NSError")
	}
	return rv, nil

}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/testingLastRuleConversion
func (t TTSSpeechSynthesizer) TestingLastRuleConversion() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("testingLastRuleConversion"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/testingSetLastRuleConversion:replacement:
func (t TTSSpeechSynthesizer) TestingSetLastRuleConversionReplacement(conversion objectivec.IObject, replacement objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("testingSetLastRuleConversion:replacement:"), conversion, replacement)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/updateCoreSynthSubstitutions
func (t TTSSpeechSynthesizer) UpdateCoreSynthSubstitutions() {
	objc.Send[objc.ID](t.ID, objc.Sel("updateCoreSynthSubstitutions"))
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/useSpecificAudioSession:
func (t TTSSpeechSynthesizer) UseSpecificAudioSession(session uint32) {
	objc.Send[objc.ID](t.ID, objc.Sel("useSpecificAudioSession:"), session)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/voiceResolver
func (t TTSSpeechSynthesizer) VoiceResolver() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("voiceResolver"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/_speechVoiceForIdentifier:language:footprint:
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) _speechVoiceForIdentifierLanguageFootprint(identifier objectivec.IObject, language objectivec.IObject, footprint int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("_speechVoiceForIdentifier:language:footprint:"), identifier, language, footprint)
	return objectivec.Object{ID: rv}
}

// SpeechVoiceForIdentifierLanguageFootprint is an exported wrapper for the private method _speechVoiceForIdentifierLanguageFootprint.
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) SpeechVoiceForIdentifierLanguageFootprint(identifier objectivec.IObject, language objectivec.IObject, footprint int64) objectivec.IObject {
	return _TTSSpeechSynthesizerClass._speechVoiceForIdentifierLanguageFootprint(identifier, language, footprint)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/audioFileSettingsForVoice:
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) AudioFileSettingsForVoice(voice objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("audioFileSettingsForVoice:"), voice)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/availableLanguageCodes
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) AvailableLanguageCodes() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("availableLanguageCodes"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/combinedProsodyMarkupForIdentifier:string:rate:pitch:volume:
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) CombinedProsodyMarkupForIdentifierStringRatePitchVolume(identifier objectivec.IObject, string_ objectivec.IObject, rate objectivec.IObject, pitch objectivec.IObject, volume objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("combinedProsodyMarkupForIdentifier:string:rate:pitch:volume:"), identifier, string_, rate, pitch, volume)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/employSpeechMarkupForType:identifier:withLanguage:
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) EmploySpeechMarkupForTypeIdentifierWithLanguage(type_ int64, identifier objectivec.IObject, language objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("employSpeechMarkupForType:identifier:withLanguage:"), type_, identifier, language)
	return rv
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/genericMarkMarkupForIdentifier:name:
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) GenericMarkMarkupForIdentifierName(identifier objectivec.IObject, name objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("genericMarkMarkupForIdentifier:name:"), identifier, name)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/isSystemVoice:
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) IsSystemVoice(voice objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("isSystemVoice:"), voice)
	return rv
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/remapVoiceIdentifier:
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) RemapVoiceIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("remapVoiceIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/speechMarkupStringForType:forIdentifier:string:
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) SpeechMarkupStringForTypeForIdentifierString(type_ int64, identifier objectivec.IObject, string_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("speechMarkupStringForType:forIdentifier:string:"), type_, identifier, string_)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/supportedIPAPhonemeLanguages
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) SupportedIPAPhonemeLanguages() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("supportedIPAPhonemeLanguages"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/unavailableVoiceIdentifiers
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) UnavailableVoiceIdentifiers() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("unavailableVoiceIdentifiers"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/voiceForIdentifier:
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) VoiceForIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("voiceForIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/audioDeviceId
func (t TTSSpeechSynthesizer) AudioDeviceId() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("audioDeviceId"))
	return rv
}
func (t TTSSpeechSynthesizer) SetAudioDeviceId(value uint32) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioDeviceId:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/audioEffects
func (t TTSSpeechSynthesizer) AudioEffects() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("audioEffects"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSSpeechSynthesizer) SetAudioEffects(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioEffects:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/audioQueueFlags
func (t TTSSpeechSynthesizer) AudioQueueFlags() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("audioQueueFlags"))
	return rv
}
func (t TTSSpeechSynthesizer) SetAudioQueueFlags(value uint32) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioQueueFlags:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/bundleIdentifier
func (t TTSSpeechSynthesizer) BundleIdentifier() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("bundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechSynthesizer) SetBundleIdentifier(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setBundleIdentifier:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/coreSynth
func (t TTSSpeechSynthesizer) CoreSynth() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t.ID, objc.Sel("coreSynth"))
	return rv
}
func (t TTSSpeechSynthesizer) SetCoreSynth(value unsafe.Pointer) {
	objc.Send[struct{}](t.ID, objc.Sel("setCoreSynth:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/delegate
func (t TTSSpeechSynthesizer) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (t TTSSpeechSynthesizer) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setDelegate:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/delegateTargetQueue
func (t TTSSpeechSynthesizer) DelegateTargetQueue() objectivec.Object {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("delegateTargetQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (t TTSSpeechSynthesizer) SetDelegateTargetQueue(value objectivec.Object) {
	objc.Send[struct{}](t.ID, objc.Sel("setDelegateTargetQueue:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/ignoreSubstitutions
func (t TTSSpeechSynthesizer) IgnoreSubstitutions() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("ignoreSubstitutions"))
	return rv
}
func (t TTSSpeechSynthesizer) SetIgnoreSubstitutions(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setIgnoreSubstitutions:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/normalizedRate
func (t TTSSpeechSynthesizer) NormalizedRate() float32 {
	rv := objc.Send[float32](t.ID, objc.Sel("normalizedRate"))
	return rv
}
func (t TTSSpeechSynthesizer) SetNormalizedRate(value float32) {
	objc.Send[struct{}](t.ID, objc.Sel("setNormalizedRate:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/outputChannels
func (t TTSSpeechSynthesizer) OutputChannels() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("outputChannels"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSSpeechSynthesizer) SetOutputChannels(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setOutputChannels:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/perVoiceSettings
func (t TTSSpeechSynthesizer) PerVoiceSettings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("perVoiceSettings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t TTSSpeechSynthesizer) SetPerVoiceSettings(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setPerVoiceSettings:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/phonemeSubstitutions
func (t TTSSpeechSynthesizer) PhonemeSubstitutions() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("phonemeSubstitutions"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSSpeechSynthesizer) SetPhonemeSubstitutions(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setPhonemeSubstitutions:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/pitch
func (t TTSSpeechSynthesizer) Pitch() float32 {
	rv := objc.Send[float32](t.ID, objc.Sel("pitch"))
	return rv
}
func (t TTSSpeechSynthesizer) SetPitch(value float32) {
	objc.Send[struct{}](t.ID, objc.Sel("setPitch:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/rate
func (t TTSSpeechSynthesizer) Rate() float32 {
	rv := objc.Send[float32](t.ID, objc.Sel("rate"))
	return rv
}
func (t TTSSpeechSynthesizer) SetRate(value float32) {
	objc.Send[struct{}](t.ID, objc.Sel("setRate:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/requestClientIdentifier
func (t TTSSpeechSynthesizer) RequestClientIdentifier() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("requestClientIdentifier"))
	return rv
}
func (t TTSSpeechSynthesizer) SetRequestClientIdentifier(value uint64) {
	objc.Send[struct{}](t.ID, objc.Sel("setRequestClientIdentifier:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/resolvedVoiceIdentifier
func (t TTSSpeechSynthesizer) ResolvedVoiceIdentifier() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("resolvedVoiceIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/skipLuthorRules
func (t TTSSpeechSynthesizer) SkipLuthorRules() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("skipLuthorRules"))
	return rv
}
func (t TTSSpeechSynthesizer) SetSkipLuthorRules(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setSkipLuthorRules:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/speakingRequestClientContext
func (t TTSSpeechSynthesizer) SpeakingRequestClientContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t.ID, objc.Sel("speakingRequestClientContext"))
	return rv
}
func (t TTSSpeechSynthesizer) SetSpeakingRequestClientContext(value unsafe.Pointer) {
	objc.Send[struct{}](t.ID, objc.Sel("setSpeakingRequestClientContext:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/speechSource
func (t TTSSpeechSynthesizer) SpeechSource() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("speechSource"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechSynthesizer) SetSpeechSource(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setSpeechSource:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/supportsAccurateWordCallbacks
func (t TTSSpeechSynthesizer) SupportsAccurateWordCallbacks() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("supportsAccurateWordCallbacks"))
	return rv
}
func (t TTSSpeechSynthesizer) SetSupportsAccurateWordCallbacks(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setSupportsAccurateWordCallbacks:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/synthesizeSilently
func (t TTSSpeechSynthesizer) SynthesizeSilently() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("synthesizeSilently"))
	return rv
}
func (t TTSSpeechSynthesizer) SetSynthesizeSilently(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setSynthesizeSilently:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/userSubstitutions
func (t TTSSpeechSynthesizer) UserSubstitutions() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("userSubstitutions"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSSpeechSynthesizer) SetUserSubstitutions(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setUserSubstitutions:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/voiceIdentifier
func (t TTSSpeechSynthesizer) VoiceIdentifier() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("voiceIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechSynthesizer) SetVoiceIdentifier(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setVoiceIdentifier:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/volume
func (t TTSSpeechSynthesizer) Volume() float32 {
	rv := objc.Send[float32](t.ID, objc.Sel("volume"))
	return rv
}
func (t TTSSpeechSynthesizer) SetVolume(value float32) {
	objc.Send[struct{}](t.ID, objc.Sel("setVolume:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizer/voucher
func (t TTSSpeechSynthesizer) Voucher() objectivec.Object {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("voucher"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (t TTSSpeechSynthesizer) SetVoucher(value objectivec.Object) {
	objc.Send[struct{}](t.ID, objc.Sel("setVoucher:"), value)
}

// SetAudioBufferCallbackSync is a synchronous wrapper around [TTSSpeechSynthesizer.SetAudioBufferCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSSpeechSynthesizer) SetAudioBufferCallbackSync(ctx context.Context) error {
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

